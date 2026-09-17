import io

import pytest
from PIL import Image

from pipeline.art import (
    BACKGROUND_SIZE,
    QUADRANT_SIZES,
    TREATMENT,
    ArtDataError,
    background_webp,
    compose_background,
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


def test_the_treatment_desaturates_darkens_and_tints_toward_the_palette():
    before = (200, 40, 40)
    after = process_pixel(before)
    # Less saturated: the spread between the channels shrinks.
    assert (max(after) - min(after)) < (max(before) - min(before)) / 2
    # Darker: nothing survives at its old brightness.
    assert sum(after) < sum(before) / 2
    # And pulled most of the way to --color-raised.
    for channel, tint in zip(after, TREATMENT.tint, strict=True):
        assert abs(channel - tint) <= 30


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
