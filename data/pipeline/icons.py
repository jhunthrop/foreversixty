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


def resolve_icon(file_id: int, names: dict[int, str], owner: str) -> str:
    """The icon name for a file data id, or the client's placeholder art.

    A file data id of 0 means the client itself has no icon for the thing, and a
    nonzero id that names no file in `ManifestInterfaceData` is the same story
    from the other side: either way the client gave us nothing to resolve.
    Emitting "" there would put `icons/.webp` in the output and 404 in the site,
    so both branches point at PLACEHOLDER_ICON and say so in the log. `owner`
    names the row that wanted the icon, e.g. "item 16866 (Helm of Might)".
    """
    name = names.get(file_id)
    if name is not None:
        return name
    logger.warning(
        "%s has no icon in the client (file data id %s); using placeholder %r",
        owner,
        file_id,
        PLACEHOLDER_ICON,
    )
    return PLACEHOLDER_ICON


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
    *,
    version: str,
) -> int:
    """Write out_dir/<name>.webp for each file id. Returns the number written.

    Every fetch pins `version` (a client build): a file data id is stable across
    builds but the bytes behind it are not, and wago's bare route serves whichever
    build is its current default. Downloads land in cache_dir/<version>/<file id>.blp
    first, so a rerun after the output directory is cleared costs nothing, and an
    icon that is already converted is left alone.
    """
    # Imported here rather than at module scope: pipeline/casc.py imports
    # CACHE_DIR and _atomic_write from this module, so a top-level import
    # would make the two modules import each other at load time.
    from pipeline.casc import CascMissing, fetch_casc_file

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
            try:
                blp = fetch_casc_file(
                    file_id, version, cache_dir=cache_dir, client=client, suffix=".blp"
                )
            except CascMissing as error:
                # CascMissing covers both a 404 and an empty body; the wording
                # below is this module's own, kept byte-for-byte what it was
                # before the migration because tests/test_icons.py matches on
                # the empty-body phrasing specifically.
                if "empty" in str(error):
                    logger.warning(
                        "icon %s (%s) is empty in CASC at version %s; skipping",
                        file_id,
                        name,
                        version,
                    )
                else:
                    logger.warning(
                        "icon %s (%s) is not in CASC at version %s; skipping",
                        file_id,
                        name,
                        version,
                    )
                continue
            _atomic_write(target, blp_to_webp(blp))
            written += 1
    finally:
        if own:
            client.close()
    return written


def _referenced_names(build_dir: Path) -> set[str]:
    """Every icon name the already-emitted JSON under build_dir refers to.

    The normalizers already resolved each talent's and each item's icon (and
    already substituted PLACEHOLDER_ICON where the client had nothing), so the
    emitted `icon` is the single source of truth for what has to be downloaded.
    Re-deriving it from the client tables here would be a second copy of that
    rule, free to drift from the one the site actually reads.
    """
    names: set[str] = set()
    for path in sorted((build_dir / "talents").glob("*.json")):
        payload = json.loads(path.read_text(encoding="utf-8"))
        for tree in payload["trees"]:
            for talent in tree["talents"]:
                names.add(talent["icon"])
    for path in sorted((build_dir / "items").glob("*.json")):
        payload = json.loads(path.read_text(encoding="utf-8"))
        for item in payload["items"]:
            names.add(item["icon"])
    return names


def wanted_icons(build_dir: Path, names: dict[int, str]) -> dict[int, str]:
    """File id -> icon name for every icon the emitted JSON under build_dir refers to."""
    referenced = _referenced_names(build_dir)
    candidates = {file_id: name for file_id, name in names.items() if name in referenced}
    _warn_about_name_collisions(candidates)
    by_name: dict[str, list[int]] = {}
    for file_id, name in candidates.items():
        by_name.setdefault(name, []).append(file_id)
    missing = sorted(referenced - set(by_name))
    if missing:
        logger.warning(
            "%d icon names are not in ManifestInterfaceData: %s", len(missing), missing[:10]
        )
    # Lowest file id wins a shared name, the same rule download_icons applies.
    wanted = {min(ids): name for name, ids in by_name.items()}
    return {file_id: wanted[file_id] for file_id in sorted(wanted)}


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
    wanted = wanted_icons(build_dir, names)
    written = download_icons(wanted, build_dir / "icons", cache_dir, client, version=build)
    print(f"{written} icons written, {len(wanted)} referenced")
    return written
