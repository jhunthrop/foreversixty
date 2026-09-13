"""Talent and item icons: names from the client, images from CASC.

`ManifestInterfaceData` maps a file data id to `Interface\\ICONS\\Name.blp`.
The site refers to icons by the lowercase name without the extension, and the
pipeline writes one 64x64 WebP per icon into `builds/<build>/icons/`.
"""

from __future__ import annotations

import io
import logging
from pathlib import Path

from PIL import Image

logger = logging.getLogger(__name__)

ICON_SIZE = 64
CACHE_DIR = Path(".icon-cache")
_BLP_SUFFIX = ".blp"


def icon_names(manifest_rows: list[dict[str, str]]) -> dict[int, str]:
    """File data id -> lowercase icon name with no extension."""
    names: dict[int, str] = {}
    for row in manifest_rows:
        file_name = row["FileName"]
        if not file_name.lower().endswith(_BLP_SUFFIX):
            continue
        names[int(row["ID"])] = file_name[: -len(_BLP_SUFFIX)].lower()
    return names


def blp_to_webp(blp: bytes, size: int = ICON_SIZE) -> bytes:
    with Image.open(io.BytesIO(blp)) as image:
        square = image.convert("RGBA").resize((size, size), Image.Resampling.LANCZOS)
    buffer = io.BytesIO()
    square.save(buffer, "WEBP", quality=90, method=6)
    return buffer.getvalue()
