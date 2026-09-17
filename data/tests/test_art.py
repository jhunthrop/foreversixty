import io
import json
from pathlib import Path

import httpx
import pytest
from PIL import Image

from pipeline.art import (
    BACKGROUND_SIZE,
    DEAD_MARGIN_THRESHOLD,
    MAX_DEAD_MARGIN_FRACTION,
    QUADRANT_SIZES,
    TREATMENT,
    ArtDataError,
    background_webp,
    backgrounds_for_build,
    compose_background,
    crop_dead_margin,
    process_pixel,
    talent_frame_ids,
)

COLOURS = {
    "TopLeft": (200, 40, 40),
    "TopRight": (40, 200, 40),
    "BottomLeft": (40, 40, 200),
    "BottomRight": (200, 200, 40),
}


def quadrant_bytes(sizes=None) -> dict[str, bytes]:
    sizes = sizes or QUADRANT_SIZES
    out = {}
    for name, colour in COLOURS.items():
        buffer = io.BytesIO()
        Image.new("RGB", sizes[name], colour).save(buffer, "PNG")
        out[name] = buffer.getvalue()
    return out


def test_the_four_quadrants_compose_into_one_320x384_image():
    image = compose_background(quadrant_bytes())
    assert image.size == BACKGROUND_SIZE
    assert image.getpixel((10, 10)) == COLOURS["TopLeft"]
    assert image.getpixel((300, 10)) == COLOURS["TopRight"]
    assert image.getpixel((10, 300)) == COLOURS["BottomLeft"]
    assert image.getpixel((300, 300)) == COLOURS["BottomRight"]


def test_a_quadrant_of_the_wrong_size_is_refused():
    sizes = {**QUADRANT_SIZES, "TopRight": (32, 256)}
    with pytest.raises(ArtDataError, match="TopRight"):
        compose_background(quadrant_bytes(sizes))


def test_a_missing_quadrant_is_refused():
    quadrants = quadrant_bytes()
    del quadrants["BottomRight"]
    with pytest.raises(ArtDataError, match="BottomRight"):
        compose_background(quadrants)


def test_a_black_bottom_and_right_margin_is_cropped_to_the_art():
    # The client's real quadrants pad their drawn art with solid near-black
    # beyond it; a stitched panel with the same shape (art in the top-left,
    # dead margin along the bottom and right) must crop to just the art.
    image = Image.new("RGB", BACKGROUND_SIZE, (0, 0, 0))
    art = Image.new("RGB", (300, 331), (120, 90, 60))
    image.paste(art, (0, 0))
    cropped = crop_dead_margin(image)
    assert cropped.size == (300, 331)
    assert cropped.getpixel((299, 330)) == (120, 90, 60)


def test_an_all_black_image_is_refused():
    image = Image.new("RGB", BACKGROUND_SIZE, (0, 0, 0))
    with pytest.raises(ArtDataError, match="dead-margin threshold"):
        crop_dead_margin(image)


def test_a_margin_free_image_is_unchanged():
    image = Image.new("RGB", BACKGROUND_SIZE, (120, 90, 60))
    cropped = crop_dead_margin(image)
    assert cropped.size == BACKGROUND_SIZE
    assert cropped.tobytes() == image.tobytes()


def test_a_pixel_at_the_threshold_itself_is_still_dead_margin():
    # DEAD_MARGIN_THRESHOLD is a strict "greater than" cutoff: a pixel whose
    # luma lands exactly on it is still margin, not art, so the crop doesn't
    # keep a sliver of near-black noise at its edge.
    image = Image.new("RGB", BACKGROUND_SIZE, (0, 0, 0))
    art = Image.new("RGB", (300, 331), (120, 90, 60))
    image.paste(art, (0, 0))
    margin_colour = (DEAD_MARGIN_THRESHOLD,) * 3
    image.paste(Image.new("RGB", (20, 331), margin_colour), (300, 0))
    cropped = crop_dead_margin(image)
    assert cropped.size == (300, 331)


def test_a_crop_over_the_max_fraction_is_refused():
    # Almost the whole panel is below the threshold -- the tiny corner of
    # "art" left is a broken-source shape, not a real dead margin, and must
    # fail loudly rather than emit a sliver.
    width, height = BACKGROUND_SIZE
    image = Image.new("RGB", BACKGROUND_SIZE, (0, 0, 0))
    sliver = max(1, round(width * (1 - MAX_DEAD_MARGIN_FRACTION) / 2))
    image.paste(Image.new("RGB", (sliver, sliver), (120, 90, 60)), (0, 0))
    with pytest.raises(ArtDataError, match="budget"):
        crop_dead_margin(image)


def test_the_treatment_desaturates_darkens_and_tints_toward_the_palette():
    before = (200, 40, 40)
    after = process_pixel(before)
    # Less saturated: the spread between the channels shrinks.
    assert (max(after) - min(after)) < (max(before) - min(before)) / 2
    # Darker: nothing survives at its old brightness.
    assert sum(after) < sum(before) / 2
    # And moved substantially toward --color-raised, without being crushed all
    # the way to it: tint_strength is deliberately low enough that a talent
    # tree's own art stays recognisable behind the grid (fix round 1 --
    # tint_strength=0.55 read as near-black on warriorarms and druidbalance).
    before_distance = sum(abs(c - t) for c, t in zip(before, TREATMENT.tint, strict=True))
    after_distance = sum(abs(c - t) for c, t in zip(after, TREATMENT.tint, strict=True))
    assert after_distance < before_distance * 0.5


def test_the_processed_image_matches_the_treatment_pixel_for_pixel():
    data = background_webp(quadrant_bytes())
    with Image.open(io.BytesIO(data)) as image:
        assert image.format == "WEBP"
        assert image.size == BACKGROUND_SIZE
        rgb = image.convert("RGB")
        for point, name in (
            ((10, 10), "TopLeft"),
            ((300, 10), "TopRight"),
            ((10, 300), "BottomLeft"),
            ((300, 300), "BottomRight"),
        ):
            want = process_pixel(COLOURS[name])
            got = rgb.getpixel(point)
            # Pillow's own blend and enhance round in C; WebP at quality 90 is
            # lossy. Two levels per channel is well inside both and still
            # pins the treatment to process_pixel, which is the definition.
            assert all(abs(g - w) <= 2 for g, w in zip(got, want, strict=True)), (name, got, want)


def test_talent_frame_ids_pairs_each_background_with_its_four_quadrants():
    rows = [
        {
            "ID": "136983",
            "FilePath": "Interface\\TALENTFRAME\\",
            "FileName": "WarriorArms-TopLeft.blp",
        },
        {
            "ID": "136984",
            "FilePath": "Interface\\TALENTFRAME\\",
            "FileName": "WarriorArms-TopRight.blp",
        },
        {
            "ID": "136981",
            "FilePath": "Interface\\TALENTFRAME\\",
            "FileName": "WarriorArms-BottomLeft.blp",
        },
        {
            "ID": "136982",
            "FilePath": "Interface\\TALENTFRAME\\",
            "FileName": "WarriorArms-BottomRight.blp",
        },
        {
            "ID": "136900",
            "FilePath": "Interface\\TALENTFRAME\\",
            "FileName": "MageArcane-TopLeft.blp",
        },
        {
            "ID": "132154",
            "FilePath": "Interface\\ICONS\\",
            "FileName": "Ability_GolemThunderClap.blp",
        },
    ]
    assert talent_frame_ids(rows) == {
        "warriorarms": {
            "TopLeft": 136983,
            "TopRight": 136984,
            "BottomLeft": 136981,
            "BottomRight": 136982,
        },
        "magearcane": {"TopLeft": 136900},
    }


def test_talent_frame_ids_matches_the_quadrant_suffix_case_insensitively():
    # 1.60.1.69893's own ManifestInterfaceData mixes cases within one build: most
    # tabs are "DruidBalance-BottomLeft.blp" (mixed case) but all of Paladin
    # Combat's are "PALADINCOMBAT-BOTTOMLEFT.BLP" (all caps). Both must resolve to
    # the same canonical quadrant name and merge under the same lowercased
    # background, or a build like that silently drops half its trees.
    rows = [
        {
            "ID": "136912",
            "FilePath": "Interface\\TALENTFRAME\\",
            "FileName": "PaladinCombat-BottomLeft.blp",
        },
        {
            "ID": "136913",
            "FilePath": "Interface\\TALENTFRAME\\",
            "FileName": "PALADINCOMBAT-BOTTOMRIGHT.BLP",
        },
    ]
    assert talent_frame_ids(rows) == {
        "paladincombat": {"BottomLeft": 136912, "BottomRight": 136913},
    }


#: Quadrant name -> the file data id it carries in the fixtures below.
FRAME_IDS = {"TopLeft": 2001, "TopRight": 2002, "BottomLeft": 2003, "BottomRight": 2004}


def _write_frame_manifest(raw_dir: Path, ids: dict[str, int], background: str = "TestTree") -> None:
    raw_dir.mkdir(parents=True, exist_ok=True)
    lines = ["ID,FilePath,FileName"]
    for quadrant, file_id in ids.items():
        lines.append(f"{file_id},Interface\\TALENTFRAME\\,{background}-{quadrant}.blp")
    (raw_dir / "ManifestInterfaceData.csv").write_text("\n".join(lines) + "\n", encoding="utf-8")


def _write_tree_json(build_dir: Path, background: str = "testtree") -> None:
    talents_dir = build_dir / "talents"
    talents_dir.mkdir(parents=True, exist_ok=True)
    (talents_dir / "warrior.json").write_text(
        json.dumps({"trees": [{"background": background}]}), encoding="utf-8"
    )


def frame_transport(
    calls: list[str],
    sizes: dict[int, tuple[int, int]],
    fail: frozenset[int] = frozenset(),
) -> httpx.MockTransport:
    def handler(request: httpx.Request) -> httpx.Response:
        calls.append(request.url.path)
        file_id = int(request.url.path.rsplit("/", 1)[-1])
        if file_id in fail:
            return httpx.Response(500, text="boom")
        buffer = io.BytesIO()
        Image.new("RGB", sizes[file_id], (10, 20, 30)).save(buffer, "PNG")
        return httpx.Response(200, content=buffer.getvalue())

    return httpx.MockTransport(handler)


def _sizes_for(ids: dict[str, int]) -> dict[int, tuple[int, int]]:
    return {file_id: QUADRANT_SIZES[quadrant] for quadrant, file_id in ids.items()}


def test_backgrounds_for_build_regenerates_the_webp_but_reuses_the_blp_cache(tmp_path: Path):
    build_dir = tmp_path / "builds" / "1.0.0.1"
    _write_frame_manifest(build_dir / "raw", FRAME_IDS)
    _write_tree_json(build_dir)
    sizes = _sizes_for(FRAME_IDS)
    cache_dir = tmp_path / "cache"

    first_calls: list[str] = []
    client = httpx.Client(transport=frame_transport(first_calls, sizes), base_url="https://x")
    written = backgrounds_for_build(
        "1.0.0.1", root=tmp_path / "builds", cache_dir=cache_dir, client=client
    )
    assert written == 1
    assert len(first_calls) == 4  # one fetch per quadrant, nothing cached yet

    # Do not delete the output this time: the target already exists from the first run,
    # and the second run must still regenerate it from the cached BLPs -- with zero new
    # HTTP calls -- rather than skipping because the file is already there. That skip is
    # exactly what let TREATMENT changes go silently stale; a client that would fail this
    # assertion the moment it is asked anything proves the cache, not the network, served it.
    target = build_dir / "trees" / "testtree.webp"
    before = target.read_bytes()
    second_calls: list[str] = []
    client_2 = httpx.Client(transport=frame_transport(second_calls, sizes), base_url="https://x")
    written_again = backgrounds_for_build(
        "1.0.0.1", root=tmp_path / "builds", cache_dir=cache_dir, client=client_2
    )
    assert written_again == 1
    assert second_calls == []
    assert target.exists()
    assert target.read_bytes() == before  # same BLP input -> byte-identical re-emit


def test_backgrounds_for_build_raises_naming_the_tree_and_missing_quadrant(tmp_path: Path):
    build_dir = tmp_path / "builds" / "1.0.0.1"
    partial = {q: i for q, i in FRAME_IDS.items() if q != "TopRight"}
    _write_frame_manifest(build_dir / "raw", partial)
    _write_tree_json(build_dir)
    client = httpx.Client(transport=frame_transport([], _sizes_for(partial)), base_url="https://x")
    with pytest.raises(ArtDataError, match=r"testtree.*TopRight"):
        backgrounds_for_build(
            "1.0.0.1", root=tmp_path / "builds", cache_dir=tmp_path / "cache", client=client
        )


def test_backgrounds_for_build_raises_on_an_http_error(tmp_path: Path):
    build_dir = tmp_path / "builds" / "1.0.0.1"
    _write_frame_manifest(build_dir / "raw", FRAME_IDS)
    _write_tree_json(build_dir)
    client = httpx.Client(
        transport=frame_transport(
            [], _sizes_for(FRAME_IDS), fail=frozenset({FRAME_IDS["TopLeft"]})
        ),
        base_url="https://x",
    )
    with pytest.raises(httpx.HTTPStatusError):
        backgrounds_for_build(
            "1.0.0.1", root=tmp_path / "builds", cache_dir=tmp_path / "cache", client=client
        )
