import json
import struct
from pathlib import Path

import httpx
from PIL import Image

from pipeline.csvio import read_csv
from pipeline.icons import blp_to_webp, download_icons, icon_names, icons_for_build, wanted_icons

HERE = Path(__file__).parent


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
    assert download_icons(*args, **kwargs) == 1
    assert download_icons(*args, **kwargs) == 0
    assert calls == ["/api/casc/132154"]


def test_download_icons_reuses_the_cache_when_the_output_is_gone(tmp_path: Path):
    calls: list[str] = []
    client = httpx.Client(transport=blp_transport(calls), base_url="https://wago.tools")
    kwargs = {"cache_dir": tmp_path / "cache", "client": client}
    download_icons({132154: "a"}, tmp_path / "one", **kwargs)
    download_icons({132154: "a"}, tmp_path / "two", **kwargs)
    assert calls == ["/api/casc/132154"]
    assert (tmp_path / "two" / "a.webp").exists()


def test_download_icons_skips_a_missing_file_id(tmp_path: Path):
    client = httpx.Client(
        transport=blp_transport([], missing=frozenset({132154})), base_url="https://wago.tools"
    )
    written = download_icons(
        {132154: "gone"}, tmp_path / "icons", cache_dir=tmp_path / "cache", client=client
    )
    assert written == 0
    assert not (tmp_path / "icons" / "gone.webp").exists()


def test_wanted_icons_reads_the_emitted_json(tmp_path: Path):
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
        json.dumps({"build": "1.0.0.1", "class_slug": "warrior", "items": [{"id": 25}]})
    )
    names = icon_names(read_csv(HERE / "fixtures/ManifestInterfaceData.csv"))
    wanted = wanted_icons(
        tmp_path,
        read_csv(HERE / "fixtures/SpellMisc.csv"),
        [{"ID": "25", "ClassID": "2", "SubclassID": "7", "IconFileDataID": "135274"}],
        names,
    )
    assert wanted == {132154: "ability_golemthunderclap", 135274: "inv_sword_04"}


def test_icons_for_build_downloads_what_the_build_refers_to(tmp_path: Path, monkeypatch):
    build_dir = tmp_path / "builds" / "1.0.0.1"
    raw = build_dir / "raw"
    raw.mkdir(parents=True)
    for table in ("SpellMisc", "ManifestInterfaceData"):
        (raw / f"{table}.csv").write_text((HERE / "fixtures" / f"{table}.csv").read_text())
    (raw / "Item.csv").write_text("ID,ClassID,SubclassID,IconFileDataID\n25,2,7,135274\n")
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


def test_download_icons_warns_on_a_name_collision_and_keeps_the_lower_id(
    tmp_path: Path, caplog
):
    calls: list[str] = []
    client = httpx.Client(transport=blp_transport(calls), base_url="https://wago.tools")
    with caplog.at_level("WARNING"):
        written = download_icons(
            {132156: "shared", 132154: "shared"},
            tmp_path / "icons",
            cache_dir=tmp_path / "cache",
            client=client,
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
    written = download_icons({132154: "a"}, tmp_path / "icons", cache_dir=tmp_path / "cache")
    assert written == 1
    assert closed == [True]
    assert calls == ["/api/casc/132154"]


def test_wanted_icons_skips_a_file_id_absent_from_the_manifest(tmp_path: Path, caplog):
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
    misc_rows = read_csv(HERE / "fixtures/SpellMisc.csv") + [
        {
            "SpellID": "99999",
            "DifficultyID": "0",
            "DurationIndex": "0",
            "SpellIconFileDataID": "999999",
        }
    ]
    with caplog.at_level("WARNING"):
        wanted = wanted_icons(tmp_path, misc_rows, [], names)
    assert 999999 not in wanted
    assert any("999999" in record.getMessage() for record in caplog.records)
