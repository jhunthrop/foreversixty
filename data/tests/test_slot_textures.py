"""pipeline.slot_textures: the client's empty paperdoll slot textures."""

import io

import pytest
from PIL import Image

from pipeline.slot_textures import HD_SIZE, SLOT_TEXTURES, hd_webp, texture_ids


def _rows(suffixes: list[str]) -> list[dict[str, str]]:
    return [
        {
            "ID": str(1000 + i),
            "FilePath": "Interface\\PaperDoll\\",
            "FileName": f"UI-PaperDoll-Slot-{s}.blp",
        }
        for i, s in enumerate(suffixes)
    ]


def test_back_uses_the_chest_texture_as_the_client_does() -> None:
    assert SLOT_TEXTURES["back"] == "Chest"
    assert SLOT_TEXTURES["off_hand"] == "SecondaryHand"


def test_texture_ids_maps_every_suffix_case_insensitively() -> None:
    suffixes = sorted(set(SLOT_TEXTURES.values()))
    rows = _rows([s.upper() for s in suffixes]) + [
        {"ID": "1", "FilePath": "Interface\\Other\\", "FileName": "UI-PaperDoll-Slot-Head.blp"}
    ]
    ids = texture_ids(rows)
    assert set(ids) == set(suffixes)
    assert 1 not in ids.values()


def test_texture_ids_names_a_missing_texture() -> None:
    with pytest.raises(ValueError, match="Wrists"):
        texture_ids(_rows([s for s in set(SLOT_TEXTURES.values()) if s != "Wrists"]))


def test_hd_webp_resamples_to_the_hd_size() -> None:
    source = io.BytesIO()
    Image.new("RGBA", (64, 64), (10, 20, 30, 255)).save(source, "PNG")
    out = Image.open(io.BytesIO(hd_webp(source.getvalue())))
    assert out.size == (HD_SIZE, HD_SIZE)
