"""Icons for the wowhead supplement (Part 2 of docs/superpowers/specs/2026-09-27-
wowhead-item-supplement-design.md): a supplement item names its own icon the same
way the client does, and that name either already resolves through
ManifestInterfaceData (the existing CASC path) or does not, in which case
wowhead's own hosted copy is the fallback.
"""

import json
import struct
from pathlib import Path

import httpx
from PIL import Image

from pipeline import wowhead_items as wh
from pipeline.icons import (
    download_zamimg_icons,
    file_id_by_name,
    icon_names,
    icons_for_build,
    supplement_icon_names,
)

HERE = Path(__file__).parent
WOWHEAD_FIXTURE = HERE / "fixtures" / "wowhead-gear-planner.js"


def write_item_sparse(raw: Path, ids: list[int]) -> None:
    lines = ["ID"] + [str(i) for i in ids]
    (raw / "ItemSparse.csv").write_text("\n".join(lines) + "\n", encoding="utf-8")


def write_classes(build_dir: Path) -> None:
    (build_dir / "classes.json").write_text(
        json.dumps([{"id": 1, "name": "Warrior", "slug": "warrior"}])
    )


def make_jpeg(colour: tuple[int, int, int] = (200, 100, 50)) -> bytes:
    import io

    image = Image.new("RGB", (8, 8), colour)
    buffer = io.BytesIO()
    image.save(buffer, "JPEG")
    return buffer.getvalue()


def make_blp2(colour: tuple[int, int, int, int] = (12, 34, 56, 255)) -> bytes:
    red, green, blue, alpha = colour
    palette = struct.pack("<4B", blue, green, red, 255) + bytes(255 * 4)
    pixels = bytes(64)
    alpha_channel = bytes([alpha]) * 64
    body = pixels + alpha_channel
    header = b"BLP2" + struct.pack("<i", 1) + struct.pack("<4b", 1, 8, 0, 1)
    header += struct.pack("<II", 8, 8)
    header += struct.pack("<16I", 148 + 1024, *([0] * 15))
    header += struct.pack("<16I", len(body), *([0] * 15))
    return header + palette + body


def test_file_id_by_name_reverses_icon_names_case_insensitively():
    rows = [
        {"ID": "1", "FilePath": "Interface\\ICONS\\", "FileName": "INV_Misc_Cape_10.blp"},
        {"ID": "2", "FilePath": "Interface\\ICONS\\", "FileName": "inv_sword_36.blp"},
    ]
    assert file_id_by_name(rows) == {"inv_misc_cape_10": 1, "inv_sword_36": 2}


def test_file_id_by_name_is_the_exact_reverse_of_icon_names():
    rows = [
        {
            "ID": "132154",
            "FilePath": "Interface\\ICONS\\",
            "FileName": "Ability_GolemThunderClap.blp",
        },
    ]
    names = icon_names(rows)
    reverse = file_id_by_name(rows)
    assert reverse == {name: file_id for file_id, name in names.items()}


def test_file_id_by_name_keeps_the_lowest_id_for_a_shared_name():
    rows = [
        {"ID": "50", "FilePath": "Interface\\ICONS\\", "FileName": "shared.blp"},
        {"ID": "12", "FilePath": "Interface\\ICONS\\", "FileName": "shared.blp"},
    ]
    assert file_id_by_name(rows) == {"shared": 12}


def test_supplement_icon_names_is_empty_without_a_payload(tmp_path: Path):
    raw = tmp_path / "raw"
    raw.mkdir()
    assert supplement_icon_names(raw) == set()


def test_supplement_icon_names_reads_the_payload_through_the_planners_own_gates(
    tmp_path: Path,
):
    raw = tmp_path / "raw"
    raw.mkdir()
    (raw / wh.RAW_FILE).write_text(WOWHEAD_FIXTURE.read_text(encoding="utf-8"), encoding="utf-8")
    write_item_sparse(raw, [11726])  # the belt is "the client's" for this test
    # The ring is quality 1, the QA robe is junk-named, the chest needs level
    # 70, and the belt is now the client's -- only the sword and the cloak
    # name a supplement icon.
    assert supplement_icon_names(raw) == {"inv_sword_36", "inv_misc_cape_10"}


def test_supplement_icon_names_excludes_an_id_the_client_already_has(tmp_path: Path):
    raw = tmp_path / "raw"
    raw.mkdir()
    (raw / wh.RAW_FILE).write_text(WOWHEAD_FIXTURE.read_text(encoding="utf-8"), encoding="utf-8")
    write_item_sparse(raw, [11726, 271218])  # the sword is now the client's too
    assert supplement_icon_names(raw) == {"inv_misc_cape_10"}


def test_download_zamimg_icons_writes_a_webp_and_caches_the_jpeg(tmp_path: Path):
    calls: list[str] = []

    def handler(request: httpx.Request) -> httpx.Response:
        calls.append(str(request.url))
        return httpx.Response(200, content=make_jpeg())

    client = httpx.Client(transport=httpx.MockTransport(handler))
    written = download_zamimg_icons(
        {"inv_sword_36"}, tmp_path / "icons", cache_dir=tmp_path / "cache", client=client
    )
    assert written == 1
    assert (tmp_path / "icons" / "inv_sword_36.webp").exists()
    assert (tmp_path / "cache" / "zamimg" / "inv_sword_36.jpg").exists()
    with Image.open(tmp_path / "icons" / "inv_sword_36.webp") as img:
        assert img.size == (64, 64)
    assert calls == ["https://wow.zamimg.com/images/wow/icons/large/inv_sword_36.jpg"]


def test_download_zamimg_icons_reuses_its_cache_without_a_second_fetch(tmp_path: Path):
    calls: list[str] = []

    def handler(request: httpx.Request) -> httpx.Response:
        calls.append(str(request.url))
        return httpx.Response(200, content=make_jpeg())

    client = httpx.Client(transport=httpx.MockTransport(handler))
    kwargs = {"cache_dir": tmp_path / "cache", "client": client}
    download_zamimg_icons({"a"}, tmp_path / "one", **kwargs)
    download_zamimg_icons({"a"}, tmp_path / "two", **kwargs)
    assert calls == ["https://wow.zamimg.com/images/wow/icons/large/a.jpg"]
    assert (tmp_path / "two" / "a.webp").exists()


def test_download_zamimg_icons_skips_an_icon_it_does_not_have(tmp_path: Path, caplog):
    client = httpx.Client(
        transport=httpx.MockTransport(lambda request: httpx.Response(404, content=b""))
    )
    with caplog.at_level("WARNING"):
        written = download_zamimg_icons(
            {"nope"}, tmp_path / "icons", cache_dir=tmp_path / "cache", client=client
        )
    assert written == 0
    assert not (tmp_path / "icons" / "nope.webp").exists()
    assert "nope" in caplog.text


def test_download_zamimg_icons_skips_work_already_done(tmp_path: Path):
    calls: list[str] = []

    def handler(request: httpx.Request) -> httpx.Response:
        calls.append(str(request.url))
        return httpx.Response(200, content=make_jpeg())

    client = httpx.Client(transport=httpx.MockTransport(handler))
    kwargs = {"cache_dir": tmp_path / "cache", "client": client}
    assert download_zamimg_icons({"a"}, tmp_path / "icons", **kwargs) == 1
    assert download_zamimg_icons({"a"}, tmp_path / "icons", **kwargs) == 0
    assert calls == ["https://wow.zamimg.com/images/wow/icons/large/a.jpg"]


def test_icons_for_build_routes_a_client_owned_supplement_icon_through_casc(
    tmp_path: Path,
):
    """The cloak's icon (inv_misc_cape_10) happens to be in this build's own
    ManifestInterfaceData (under some other file id); it must be fetched from
    CASC, not from zamimg."""
    build_dir = tmp_path / "builds" / "1.0.0.1"
    raw = build_dir / "raw"
    raw.mkdir(parents=True)
    (raw / "ManifestInterfaceData.csv").write_text(
        "ID,FilePath,FileName\n1,Interface\\ICONS\\,inv_misc_cape_10.blp\n",
        encoding="utf-8",
    )
    (raw / wh.RAW_FILE).write_text(WOWHEAD_FIXTURE.read_text(encoding="utf-8"), encoding="utf-8")
    write_item_sparse(raw, [11726])  # belt is the client's; sword and cloak are not
    write_classes(build_dir)
    (build_dir / "talents").mkdir()
    (build_dir / "items").mkdir()
    # rotation_icon_names (icons_for_build's own addition, day3 data-
    # followups-11 lane) needs a spellranks.json to read even though this
    # test has nothing to do with rotations; an empty-but-well-formed one,
    # with curated_dir pointing at a directory with no apl/ folder (so the
    # REAL data/curated/apl specs this test's default cwd would otherwise
    # pick up never enter it), resolves to "no rotation icons" rather than
    # a missing-file refusal -- the same stub
    # test_icons_for_build_downloads_what_the_build_refers_to uses.
    (build_dir / "spellranks.json").write_text(json.dumps({"build": "1.0.0.1", "classes": {}}))

    casc_calls: list[str] = []
    zamimg_calls: list[str] = []

    def handler(request: httpx.Request) -> httpx.Response:
        if request.url.host == "wago.tools":
            casc_calls.append(request.url.path)
            return httpx.Response(200, content=make_blp2())
        zamimg_calls.append(request.url.path)
        return httpx.Response(200, content=make_jpeg())

    client = httpx.Client(transport=httpx.MockTransport(handler), base_url="https://wago.tools")
    written = icons_for_build(
        "1.0.0.1",
        root=tmp_path / "builds",
        cache_dir=tmp_path / "cache",
        client=client,
        curated_dir=tmp_path / "curated",
    )
    assert written == 2
    assert casc_calls == ["/api/casc/1"]  # the cloak, through CASC
    assert zamimg_calls == ["/images/wow/icons/large/inv_sword_36.jpg"]  # the sword, via zamimg
    assert (build_dir / "icons" / "inv_misc_cape_10.webp").exists()
    assert (build_dir / "icons" / "inv_sword_36.webp").exists()


def test_icons_for_build_is_unaffected_by_a_build_with_no_wowhead_payload(tmp_path: Path):
    build_dir = tmp_path / "builds" / "1.0.0.1"
    raw = build_dir / "raw"
    raw.mkdir(parents=True)
    (raw / "ManifestInterfaceData.csv").write_text(
        "ID,FilePath,FileName\n1,Interface\\ICONS\\,inv_misc_cape_10.blp\n",
        encoding="utf-8",
    )
    write_item_sparse(raw, [11726])
    write_classes(build_dir)
    (build_dir / "talents").mkdir()
    (build_dir / "items").mkdir()
    # Same stub as the test above, for the same reason: rotation_icon_names
    # needs a spellranks.json to read, and curated_dir must not be the real
    # data/curated/ this test's default cwd would otherwise resolve to.
    (build_dir / "spellranks.json").write_text(json.dumps({"build": "1.0.0.1", "classes": {}}))

    client = httpx.Client(
        transport=httpx.MockTransport(lambda request: httpx.Response(200, content=make_blp2())),
        base_url="https://wago.tools",
    )
    written = icons_for_build(
        "1.0.0.1",
        root=tmp_path / "builds",
        cache_dir=tmp_path / "cache",
        client=client,
        curated_dir=tmp_path / "curated",
    )
    assert written == 0  # nothing referenced by the (empty) talents/items JSON
    assert not (build_dir / "icons").exists() or list((build_dir / "icons").iterdir()) == []
