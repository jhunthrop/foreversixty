import io
import json
from pathlib import Path

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
from tests.cdn_fakes import FakeSource

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


def frame_source(sizes: dict[int, tuple[int, int]], fail: frozenset[int] = frozenset()):
    """Each quadrant's file as a solid PNG of its size; an id in `fail` raises."""

    def png(file_id: int):
        def make() -> bytes:
            if file_id in fail:
                raise RuntimeError("CDN hosts answered [500]")
            buffer = io.BytesIO()
            Image.new("RGB", sizes[file_id], (10, 20, 30)).save(buffer, "PNG")
            return buffer.getvalue()

        return make

    return FakeSource({file_id: png(file_id) for file_id in sizes})


def _sizes_for(ids: dict[str, int]) -> dict[int, tuple[int, int]]:
    return {file_id: QUADRANT_SIZES[quadrant] for quadrant, file_id in ids.items()}


def test_backgrounds_for_build_regenerates_the_webp_but_reuses_the_blp_cache(tmp_path: Path):
    build_dir = tmp_path / "builds" / "1.0.0.1"
    _write_frame_manifest(build_dir / "raw", FRAME_IDS)
    _write_tree_json(build_dir)
    sizes = _sizes_for(FRAME_IDS)
    cache_dir = tmp_path / "cache"

    first = frame_source(sizes)
    written = backgrounds_for_build(
        "1.0.0.1", root=tmp_path / "builds", cache_dir=cache_dir, source=first
    )
    assert written == 1
    assert len(first.calls) == 4  # one read per quadrant, nothing cached yet

    # Do not delete the output this time: the target already exists from the first run,
    # and the second run must still regenerate it from the cached BLPs -- with zero new
    # reads -- rather than skipping because the file is already there. That skip is
    # exactly what let TREATMENT changes go silently stale.
    target = build_dir / "trees" / "testtree.webp"
    before = target.read_bytes()
    second = frame_source(sizes)
    written_again = backgrounds_for_build(
        "1.0.0.1", root=tmp_path / "builds", cache_dir=cache_dir, source=second
    )
    assert written_again == 1
    assert second.calls == []
    assert target.exists()
    assert target.read_bytes() == before  # same BLP input -> byte-identical re-emit


def test_backgrounds_for_build_keys_the_cache_by_build_and_rejects_an_empty_file(
    tmp_path: Path,
):
    """A file data id is stable across builds but its bytes are not, so the cache is
    keyed by build and an empty file is an error, not an empty texture."""
    build_dir = tmp_path / "builds" / "1.0.0.1"
    _write_frame_manifest(build_dir / "raw", FRAME_IDS)
    _write_tree_json(build_dir)
    backgrounds_for_build(
        "1.0.0.1",
        root=tmp_path / "builds",
        cache_dir=tmp_path / "c",
        source=frame_source(_sizes_for(FRAME_IDS)),
    )
    assert (tmp_path / "c" / "1.0.0.1").is_dir(), "the cache is keyed by build as well as id"

    empty = FakeSource({file_id: b"" for file_id in FRAME_IDS.values()})
    with pytest.raises(ArtDataError, match="no bytes"):
        backgrounds_for_build(
            "1.0.0.1", root=tmp_path / "builds", cache_dir=tmp_path / "c2", source=empty
        )


def test_backgrounds_for_build_raises_naming_the_tree_and_missing_quadrant(tmp_path: Path):
    build_dir = tmp_path / "builds" / "1.0.0.1"
    partial = {q: i for q, i in FRAME_IDS.items() if q != "TopRight"}
    _write_frame_manifest(build_dir / "raw", partial)
    _write_tree_json(build_dir)
    with pytest.raises(ArtDataError, match=r"testtree.*TopRight"):
        backgrounds_for_build(
            "1.0.0.1",
            root=tmp_path / "builds",
            cache_dir=tmp_path / "cache",
            source=frame_source(_sizes_for(partial)),
        )


def test_backgrounds_for_build_raises_on_a_cdn_error(tmp_path: Path):
    build_dir = tmp_path / "builds" / "1.0.0.1"
    _write_frame_manifest(build_dir / "raw", FRAME_IDS)
    _write_tree_json(build_dir)
    source = frame_source(_sizes_for(FRAME_IDS), fail=frozenset({FRAME_IDS["TopLeft"]}))
    with pytest.raises(RuntimeError, match="500"):
        backgrounds_for_build(
            "1.0.0.1", root=tmp_path / "builds", cache_dir=tmp_path / "cache", source=source
        )
