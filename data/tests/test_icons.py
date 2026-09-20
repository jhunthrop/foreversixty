import json
import struct
from pathlib import Path

import httpx
from PIL import Image

from pipeline.csvio import read_csv
from pipeline.icons import (
    PLACEHOLDER_ICON,
    blp_to_webp,
    class_icon_names,
    download_icons,
    icon_names,
    icons_for_build,
    wanted_icons,
)

HERE = Path(__file__).parent

#: The classes a real build's classes.json lists (id is irrelevant here; only slug
#: feeds class_icon_names). Mirrors builds/1.60.1.69893/classes.json's slug list.
ALL_CLASS_SLUGS = (
    "warrior",
    "paladin",
    "hunter",
    "rogue",
    "priest",
    "shaman",
    "mage",
    "warlock",
    "druid",
)


def write_classes(tmp_path: Path, slugs: tuple[str, ...] = ("warrior",)) -> None:
    """A minimal classes.json: every wanted_icons/_referenced_names fixture build
    dir needs one now that class icons come from it, same as a real build's does
    (pipeline/normalize writes classes.json before `icons` ever runs)."""
    (tmp_path / "classes.json").write_text(
        json.dumps([{"id": n, "name": slug.title(), "slug": slug} for n, slug in enumerate(slugs)])
    )


def make_blp2(
    width: int = 8,
    height: int = 8,
    colour: tuple[int, int, int, int] = (12, 34, 56, 255),
) -> bytes:
    """A minimal palettised BLP2, the format wago.tools serves from CASC."""
    red, green, blue, alpha = colour
    palette = struct.pack("<4B", blue, green, red, 255) + bytes(255 * 4)
    pixels = bytes(width * height)  # every pixel is palette index 0
    alpha_channel = bytes([alpha]) * (width * height)
    body = pixels + alpha_channel
    header = b"BLP2" + struct.pack("<i", 1) + struct.pack("<4b", 1, 8, 0, 1)
    header += struct.pack("<II", width, height)
    header += struct.pack("<16I", 148 + 1024, *([0] * 15))
    header += struct.pack("<16I", len(body), *([0] * 15))
    return header + palette + body


def test_icon_names_lowercase_and_drop_the_extension():
    names = icon_names(read_csv(HERE / "fixtures/ManifestInterfaceData.csv"))
    assert names[132154] == "ability_golemthunderclap"
    assert names[132759] == "inv_chest_samurai"


def test_icon_names_skip_non_blp_entries():
    names = icon_names([{"ID": "1", "FilePath": "Interface\\FrameXML\\", "FileName": "Fonts.xml"}])
    assert names == {}


def test_blp_to_webp_produces_a_64_square(tmp_path: Path):
    data = blp_to_webp(make_blp2())
    out = tmp_path / "icon.webp"
    out.write_bytes(data)
    with Image.open(out) as img:
        assert img.format == "WEBP"
        assert img.size == (64, 64)
        # Lossy WebP at quality=90 can shift a channel by 1 (quantization), so this
        # compares with a +-1 per-channel tolerance instead of exact equality. Do not
        # tighten this back to `==`: it is not a bug, it is the lossy codec doing its
        # job, and the exact shift is specific to the Pillow/libwebp build in uv.lock.
        pixel = img.convert("RGBA").getpixel((32, 32))
        expected = (12, 34, 56, 255)
        channel_pairs = zip(pixel, expected, strict=True)
        assert all(abs(actual - want) <= 1 for actual, want in channel_pairs)


def test_blp_to_webp_is_deterministic():
    blp = make_blp2()
    assert blp_to_webp(blp) == blp_to_webp(blp)


def blp_transport(calls: list[str], missing: frozenset[int] = frozenset()) -> httpx.MockTransport:
    def handler(request: httpx.Request) -> httpx.Response:
        calls.append(request.url.path)
        file_id = int(request.url.path.rsplit("/", 1)[-1])
        if file_id in missing:
            return httpx.Response(404, json={"errors": "Not found."})
        return httpx.Response(200, content=make_blp2())

    return httpx.MockTransport(handler)


def test_download_icons_writes_one_webp_per_name(tmp_path: Path):
    calls: list[str] = []
    client = httpx.Client(transport=blp_transport(calls), base_url="https://wago.tools")
    written = download_icons(
        {132154: "ability_golemthunderclap", 132156: "ability_hibernation"},
        tmp_path / "icons",
        cache_dir=tmp_path / "cache",
        client=client,
        version="1.0.0.1",
    )
    assert written == 2
    assert sorted(p.name for p in (tmp_path / "icons").iterdir()) == [
        "ability_golemthunderclap.webp",
        "ability_hibernation.webp",
    ]
    assert calls == ["/api/casc/132154", "/api/casc/132156"]


def test_download_icons_skips_work_that_is_already_done(tmp_path: Path):
    calls: list[str] = []
    client = httpx.Client(transport=blp_transport(calls), base_url="https://wago.tools")
    args = ({132154: "ability_golemthunderclap"}, tmp_path / "icons")
    kwargs = {"cache_dir": tmp_path / "cache", "client": client}
    assert download_icons(*args, **kwargs, version="1.0.0.1") == 1
    assert download_icons(*args, **kwargs, version="1.0.0.1") == 0
    assert calls == ["/api/casc/132154"]


def test_download_icons_reuses_the_cache_when_the_output_is_gone(tmp_path: Path):
    calls: list[str] = []
    client = httpx.Client(transport=blp_transport(calls), base_url="https://wago.tools")
    kwargs = {"cache_dir": tmp_path / "cache", "client": client}
    download_icons({132154: "a"}, tmp_path / "one", **kwargs, version="1.0.0.1")
    download_icons({132154: "a"}, tmp_path / "two", **kwargs, version="1.0.0.1")
    assert calls == ["/api/casc/132154"]
    assert (tmp_path / "two" / "a.webp").exists()


def test_download_icons_skips_a_missing_file_id(tmp_path: Path):
    client = httpx.Client(
        transport=blp_transport([], missing=frozenset({132154})), base_url="https://wago.tools"
    )
    written = download_icons(
        {132154: "gone"},
        tmp_path / "icons",
        cache_dir=tmp_path / "cache",
        client=client,
        version="1.0.0.1",
    )
    assert written == 0
    assert not (tmp_path / "icons" / "gone.webp").exists()


def test_wanted_icons_reads_the_emitted_json(tmp_path: Path):
    write_classes(tmp_path)
    (tmp_path / "talents").mkdir(parents=True)
    (tmp_path / "talents" / "warrior.json").write_text(
        json.dumps(
            {
                "build": "1.0.0.1",
                "class_id": 1,
                "class_slug": "warrior",
                "trees": [
                    {
                        "id": 161,
                        "name": "Arms",
                        "position": 0,
                        "talents": [
                            {
                                "id": 124,
                                "name": "Improved Heroic Strike",
                                "icon": "ability_golemthunderclap",
                                "max_rank": 1,
                                "tier": 0,
                                "column": 0,
                                "prereq_talent_id": None,
                                "prereq_rank": None,
                                "ranks": [{"spell_id": 12282, "description": "x"}],
                            }
                        ],
                    }
                ],
            }
        )
    )
    (tmp_path / "items").mkdir()
    (tmp_path / "items" / "warrior.json").write_text(
        json.dumps(
            {
                "build": "1.0.0.1",
                "class_slug": "warrior",
                "items": [{"id": 25, "icon": "inv_sword_04"}],
            }
        )
    )
    names = icon_names(read_csv(HERE / "fixtures/ManifestInterfaceData.csv"))
    wanted = wanted_icons(tmp_path, names)
    assert wanted == {132154: "ability_golemthunderclap", 135274: "inv_sword_04"}


def test_class_icon_names_reads_every_class_slug(tmp_path: Path):
    """No talent or item ever references a class icon, so class_icon_names has to
    read classes.json directly rather than wait for one of those to name it."""
    write_classes(tmp_path, ALL_CLASS_SLUGS)
    assert class_icon_names(tmp_path) == {f"classicon_{slug}" for slug in ALL_CLASS_SLUGS}


def test_wanted_icons_asks_for_a_classicon_per_class(tmp_path: Path):
    """A real build's classes.json lists every playable class; wanted_icons must ask
    CASC for that class's icon regardless of what the talent and item JSON reference,
    or a real build ships with some classes missing their icon (the bug this pins)."""
    write_classes(tmp_path, ALL_CLASS_SLUGS)
    (tmp_path / "talents").mkdir()
    (tmp_path / "items").mkdir()
    class_icon_file_ids = {
        10_000 + n: f"classicon_{slug}" for n, slug in enumerate(ALL_CLASS_SLUGS)
    }
    wanted = wanted_icons(tmp_path, class_icon_file_ids)
    assert set(wanted.values()) == {f"classicon_{slug}" for slug in ALL_CLASS_SLUGS}


def test_icons_for_build_downloads_what_the_build_refers_to(tmp_path: Path, monkeypatch):
    build_dir = tmp_path / "builds" / "1.0.0.1"
    raw = build_dir / "raw"
    raw.mkdir(parents=True)
    manifest = HERE / "fixtures" / "ManifestInterfaceData.csv"
    (raw / manifest.name).write_text(manifest.read_text())
    write_classes(build_dir)
    (build_dir / "talents").mkdir()
    (build_dir / "talents" / "warrior.json").write_text(
        json.dumps(
            {
                "build": "1.0.0.1",
                "class_id": 1,
                "class_slug": "warrior",
                "trees": [
                    {
                        "id": 161,
                        "name": "Arms",
                        "position": 0,
                        "talents": [
                            {
                                "id": 124,
                                "name": "x",
                                "icon": "ability_golemthunderclap",
                                "max_rank": 1,
                                "tier": 0,
                                "column": 0,
                                "prereq_talent_id": None,
                                "prereq_rank": None,
                                "ranks": [{"spell_id": 12282, "description": "x"}],
                            }
                        ],
                    }
                ],
            }
        )
    )
    client = httpx.Client(transport=blp_transport([]), base_url="https://wago.tools")
    written = icons_for_build(
        "1.0.0.1", root=tmp_path / "builds", cache_dir=tmp_path / "cache", client=client
    )
    assert written == 1
    assert (build_dir / "icons" / "ability_golemthunderclap.webp").exists()


def test_download_icons_warns_on_a_name_collision_and_keeps_the_lower_id(tmp_path: Path, caplog):
    calls: list[str] = []
    client = httpx.Client(transport=blp_transport(calls), base_url="https://wago.tools")
    with caplog.at_level("WARNING"):
        written = download_icons(
            {132156: "shared", 132154: "shared"},
            tmp_path / "icons",
            cache_dir=tmp_path / "cache",
            client=client,
            version="1.0.0.1",
        )
    assert written == 1
    assert (tmp_path / "icons" / "shared.webp").exists()
    assert calls == ["/api/casc/132154"]
    assert any(
        "132154" in record.getMessage() and "132156" in record.getMessage()
        for record in caplog.records
    )


def test_download_icons_creates_and_closes_its_own_client(tmp_path: Path, monkeypatch):
    calls: list[str] = []
    closed: list[bool] = []
    real_client_cls = httpx.Client

    class TrackingClient(real_client_cls):
        def close(self):
            closed.append(True)
            super().close()

    def fake_client(*args, **kwargs):
        kwargs["transport"] = blp_transport(calls)
        return TrackingClient(*args, **kwargs)

    monkeypatch.setattr(httpx, "Client", fake_client)
    written = download_icons(
        {132154: "a"}, tmp_path / "icons", cache_dir=tmp_path / "cache", version="1.0.0.1"
    )
    assert written == 1
    assert closed == [True]
    assert calls == ["/api/casc/132154"]


def test_wanted_icons_skips_an_icon_name_absent_from_the_manifest(tmp_path: Path, caplog):
    write_classes(tmp_path)
    (tmp_path / "talents").mkdir(parents=True)
    (tmp_path / "talents" / "warrior.json").write_text(
        json.dumps(
            {
                "build": "1.0.0.1",
                "class_id": 1,
                "class_slug": "warrior",
                "trees": [
                    {
                        "id": 161,
                        "name": "Arms",
                        "position": 0,
                        "talents": [
                            {
                                "id": 124,
                                "name": "x",
                                "icon": "whatever",
                                "max_rank": 1,
                                "tier": 0,
                                "column": 0,
                                "prereq_talent_id": None,
                                "prereq_rank": None,
                                "ranks": [{"spell_id": 99999, "description": "x"}],
                            }
                        ],
                    }
                ],
            }
        )
    )
    (tmp_path / "items").mkdir()
    (tmp_path / "items" / "warrior.json").write_text(
        json.dumps({"build": "1.0.0.1", "class_slug": "warrior", "items": []})
    )
    names = icon_names(read_csv(HERE / "fixtures/ManifestInterfaceData.csv"))
    with caplog.at_level("WARNING"):
        wanted = wanted_icons(tmp_path, names)
    assert "whatever" not in wanted.values()
    assert any("whatever" in record.getMessage() for record in caplog.records)


def write_items(tmp_path: Path, icons: list[str]) -> None:
    write_classes(tmp_path)
    (tmp_path / "items").mkdir(parents=True, exist_ok=True)
    (tmp_path / "items" / "warrior.json").write_text(
        json.dumps(
            {
                "build": "1.0.0.1",
                "class_slug": "warrior",
                "items": [{"id": 25 + n, "icon": icon} for n, icon in enumerate(icons)],
            }
        )
    )


def test_wanted_icons_downloads_the_placeholder_for_an_item_with_no_icon(tmp_path: Path):
    """Items the client has no art for are emitted pointing at PLACEHOLDER_ICON, so its
    file has to be fetched too, or the output would reference a missing image."""
    write_items(tmp_path, ["inv_sword_04", PLACEHOLDER_ICON])
    wanted = wanted_icons(tmp_path, {135274: "inv_sword_04", 134400: PLACEHOLDER_ICON})
    assert wanted == {135274: "inv_sword_04", 134400: PLACEHOLDER_ICON}


def test_wanted_icons_needs_no_placeholder_when_every_item_has_one(tmp_path: Path):
    write_items(tmp_path, ["inv_sword_04"])
    wanted = wanted_icons(tmp_path, {135274: "inv_sword_04", 134400: PLACEHOLDER_ICON})
    assert wanted == {135274: "inv_sword_04"}


def test_wanted_icons_keeps_the_lowest_file_id_for_a_shared_name(tmp_path: Path, caplog):
    """Two file ids can claim one lowercase name; the site addresses icons by name,
    so exactly one file wins, and it is the same one download_icons would write."""
    write_items(tmp_path, ["inv_sword_04"])
    with caplog.at_level("WARNING"):
        wanted = wanted_icons(tmp_path, {135274: "inv_sword_04", 99: "inv_sword_04"})
    assert wanted == {99: "inv_sword_04"}
    assert any("inv_sword_04" in record.getMessage() for record in caplog.records)


def test_download_icons_pins_every_fetch_to_the_version_and_keys_the_cache_by_it(tmp_path: Path):
    """A file data id is stable across builds but its bytes are not, so the version rides
    on every request and the cache cannot hand one build's bytes to another."""
    versions: list[str | None] = []

    def handler(request: httpx.Request) -> httpx.Response:
        versions.append(request.url.params.get("version"))
        return httpx.Response(200, content=make_blp2())

    client = httpx.Client(transport=httpx.MockTransport(handler), base_url="https://wago.tools")
    download_icons(
        {132154: "a"}, tmp_path / "one", cache_dir=tmp_path / "c", client=client, version="1.0.0.1"
    )
    download_icons(
        {132154: "a"}, tmp_path / "two", cache_dir=tmp_path / "c", client=client, version="2.0.0.1"
    )
    assert versions == ["1.0.0.1", "2.0.0.1"]
    assert (tmp_path / "c" / "1.0.0.1" / "132154.blp").exists()
    assert (tmp_path / "c" / "2.0.0.1" / "132154.blp").exists()


def test_download_icons_skips_an_icon_the_version_serves_empty(tmp_path: Path, caplog):
    client = httpx.Client(
        transport=httpx.MockTransport(lambda request: httpx.Response(200, content=b"")),
        base_url="https://wago.tools",
    )
    written = download_icons(
        {132154: "a"},
        tmp_path / "icons",
        cache_dir=tmp_path / "c",
        client=client,
        version="1.0.0.1",
    )
    assert written == 0
    assert not (tmp_path / "icons" / "a.webp").exists()
    assert "empty in CASC at version 1.0.0.1" in caplog.text
