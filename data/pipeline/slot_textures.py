"""The client's empty paperdoll slot textures, as the site's HD slot art.

The Classic character frame draws an empty slot with its own background
texture (`PaperDollItemSlotButton_Update`: `textureName =
self.backgroundTextureName`), one `Interface\\PaperDoll\\UI-PaperDoll-Slot-*`
BLP per slot. The site's paperdoll (design/specs/2026-10-31-paperdoll.md §8)
draws exactly those, as faithful HD reproductions: the client's 64 px pixels,
resampled to 128 px with Lanczos, nothing redrawn or re-tinted.

Back uses `-Chest`, as the client does (`CharacterBackSlot`'s
backgroundTextureName). Run from data/:

    python -m pipeline.slot_textures <build> <out dir>
"""

from __future__ import annotations

import io
import sys
from pathlib import Path

from PIL import Image

from pipeline.casc import fetch_casc_file

#: The site's slot key -> the client texture's suffix after `UI-PaperDoll-Slot-`.
SLOT_TEXTURES: dict[str, str] = {
    "head": "Head",
    "neck": "Neck",
    "shoulder": "Shoulder",
    "back": "Chest",
    "chest": "Chest",
    "shirt": "Shirt",
    "tabard": "Tabard",
    "wrist": "Wrists",
    "hands": "Hands",
    "waist": "Waist",
    "legs": "Legs",
    "feet": "Feet",
    "finger": "Finger",
    "trinket": "Trinket",
    "main_hand": "MainHand",
    "off_hand": "SecondaryHand",
    "ranged": "Ranged",
    "relic": "Relic",
}

#: The client ships these at 64 px; the site draws a socket up to 44 CSS px, so 128 covers 2x.
HD_SIZE = 128
_DIR = "interface\\paperdoll\\"
_PREFIX = "ui-paperdoll-slot-"


def texture_ids(manifest_rows: list[dict[str, str]]) -> dict[str, int]:
    """Texture suffix (as in SLOT_TEXTURES) -> CASC file data id."""
    wanted = {suffix.lower(): suffix for suffix in SLOT_TEXTURES.values()}
    out: dict[str, int] = {}
    for row in manifest_rows:
        name = row["FileName"].lower()
        if (
            row["FilePath"].lower() != _DIR
            or not name.startswith(_PREFIX)
            or not name.endswith(".blp")
        ):
            continue
        suffix = wanted.get(name[len(_PREFIX) : -len(".blp")])
        if suffix is not None:
            out[suffix] = int(row["ID"])
    missing = sorted(set(wanted.values()) - set(out))
    if missing:
        raise ValueError(f"ManifestInterfaceData names no UI-PaperDoll-Slot texture for {missing}")
    return out


def hd_webp(blp: bytes) -> bytes:
    """The texture at HD_SIZE, Lanczos-resampled, as lossless WebP."""
    image = Image.open(io.BytesIO(blp)).convert("RGBA")
    out = io.BytesIO()
    image.resize((HD_SIZE, HD_SIZE), Image.Resampling.LANCZOS).save(out, "WEBP", lossless=True)
    return out.getvalue()


def write_slot_textures(build: str, manifest: Path, out_dir: Path) -> int:
    from pipeline.csvio import read_csv

    ids = texture_ids(read_csv(manifest))
    out_dir.mkdir(parents=True, exist_ok=True)
    for slot, suffix in SLOT_TEXTURES.items():
        (out_dir / f"{slot}.webp").write_bytes(
            hd_webp(fetch_casc_file(ids[suffix], build, suffix=".blp"))
        )
    return len(SLOT_TEXTURES)


if __name__ == "__main__":
    if len(sys.argv) != 4:
        raise SystemExit(
            "usage: python -m pipeline.slot_textures <build> <ManifestInterfaceData.csv> <out dir>"
        )
    written = write_slot_textures(sys.argv[1], Path(sys.argv[2]), Path(sys.argv[3]))
    print(f"{written} slot textures written")
