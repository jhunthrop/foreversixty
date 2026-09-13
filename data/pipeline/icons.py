"""Talent and item icons: names from the client, images from CASC.

`ManifestInterfaceData` maps a file data id to `Interface\\ICONS\\Name.blp`.
The site refers to icons by the lowercase name without the extension, and the
pipeline writes one 64x64 WebP per icon into `builds/<build>/icons/`.
"""

from __future__ import annotations

import io
import json
import logging
import os
from pathlib import Path

import httpx
from PIL import Image

from pipeline.wago import BASE_URL, USER_AGENT

logger = logging.getLogger(__name__)

ICON_SIZE = 64
CACHE_DIR = Path(".icon-cache")
_BLP_SUFFIX = ".blp"

#: The client's own placeholder art, shown for anything the client itself has no
#: icon for. This is a real `ManifestInterfaceData` row shipped in the game data,
#: not a name this pipeline invented: items whose `IconFileDataID` is 0 (or names
#: no file) are emitted pointing here so that every `icon` in the output resolves
#: to an image, rather than to an empty name the site would request as `.webp`.
PLACEHOLDER_ICON = "inv_misc_questionmark"


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


def _atomic_write(path: Path, data: bytes) -> None:
    """Write data to path without ever leaving a partially-written file behind."""
    tmp = path.parent / f"{path.name}.tmp-{os.getpid()}"
    tmp.write_bytes(data)
    os.replace(tmp, path)


def _warn_about_name_collisions(wanted: dict[int, str]) -> None:
    """Two distinct file ids can map to the same lowercase icon name.

    The site addresses icons by name, so exactly one file per name is correct
    by design, but a silent skip is indistinguishable from a substituted
    icon: log which ids collided and which one wins.
    """
    by_name: dict[str, list[int]] = {}
    for file_id, name in wanted.items():
        by_name.setdefault(name, []).append(file_id)
    for name, ids in sorted(by_name.items()):
        if len(ids) > 1:
            logger.warning(
                "icon name %r is claimed by file ids %s; only %s will be written",
                name,
                sorted(ids),
                min(ids),
            )


def download_icons(
    wanted: dict[int, str],
    out_dir: Path,
    cache_dir: Path = CACHE_DIR,
    client: httpx.Client | None = None,
) -> int:
    """Write out_dir/<name>.webp for each file id. Returns the number written.

    Downloads land in cache_dir/<file id>.blp first, so a rerun after the
    output directory is cleared costs nothing, and an icon that is already
    converted is left alone.
    """
    own = client is None
    if client is None:
        client = httpx.Client(base_url=BASE_URL, headers={"User-Agent": USER_AGENT})
    out_dir.mkdir(parents=True, exist_ok=True)
    cache_dir.mkdir(parents=True, exist_ok=True)
    _warn_about_name_collisions(wanted)
    written = 0
    try:
        for file_id, name in sorted(wanted.items()):
            target = out_dir / f"{name}.webp"
            if target.exists():
                continue
            cached = cache_dir / f"{file_id}.blp"
            if not cached.exists():
                response = client.get(f"/api/casc/{file_id}", timeout=60)
                if response.status_code == 404:
                    logger.warning("icon %s (%s) is not in CASC; skipping", file_id, name)
                    continue
                response.raise_for_status()
                _atomic_write(cached, response.content)
            _atomic_write(target, blp_to_webp(cached.read_bytes()))
            written += 1
    finally:
        if own:
            client.close()
    return written


def wanted_icons(
    build_dir: Path,
    misc_rows: list[dict[str, str]],
    item_rows: list[dict[str, str]],
    names: dict[int, str],
) -> dict[int, str]:
    """The icons the already-normalised JSON under build_dir refers to."""
    spell_icons = {
        int(r["SpellID"]): int(r["SpellIconFileDataID"])
        for r in misc_rows
        if r.get("DifficultyID", "0") == "0"
    }
    item_icons = {int(r["ID"]): int(r["IconFileDataID"]) for r in item_rows}
    file_ids: set[int] = set()
    for path in sorted((build_dir / "talents").glob("*.json")):
        payload = json.loads(path.read_text(encoding="utf-8"))
        for tree in payload["trees"]:
            for talent in tree["talents"]:
                file_ids.add(spell_icons.get(talent["ranks"][0]["spell_id"], 0))
    # Items the client gives no icon for are emitted pointing at PLACEHOLDER_ICON
    # (see pipeline.normalize.gear), so its file must be downloaded too. Resolve
    # the id from the client's own table rather than hardcoding it, and pick the
    # lowest on a name collision, the same rule download_icons applies.
    placeholder_id = min(
        (file_id for file_id, name in names.items() if name == PLACEHOLDER_ICON),
        default=0,
    )
    for path in sorted((build_dir / "items").glob("*.json")):
        payload = json.loads(path.read_text(encoding="utf-8"))
        for item in payload["items"]:
            file_id = item_icons.get(item["id"], 0)
            file_ids.add(file_id if file_id in names else placeholder_id)
    file_ids.discard(0)
    missing = sorted(i for i in file_ids if i not in names)
    if missing:
        logger.warning(
            "%d icon file ids are not in ManifestInterfaceData: %s", len(missing), missing[:10]
        )
    return {file_id: names[file_id] for file_id in sorted(file_ids) if file_id in names}


def icons_for_build(
    build: str,
    root: Path = Path("builds"),
    cache_dir: Path = CACHE_DIR,
    client: httpx.Client | None = None,
) -> int:
    from pipeline.csvio import read_csv

    build_dir = root / build
    raw = build_dir / "raw"
    if not raw.exists():
        raise SystemExit(f"no raw data at {raw}; run `python -m pipeline fetch` first")
    names = icon_names(read_csv(raw / "ManifestInterfaceData.csv"))
    wanted = wanted_icons(
        build_dir, read_csv(raw / "SpellMisc.csv"), read_csv(raw / "Item.csv"), names
    )
    written = download_icons(wanted, build_dir / "icons", cache_dir, client)
    print(f"{written} icons written, {len(wanted)} referenced")
    return written
