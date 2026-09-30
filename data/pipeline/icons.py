"""Talent and item icons: names from the client, images from CASC.

`ManifestInterfaceData` maps a file data id to `Interface\\ICONS\\Name.blp`.
The site refers to icons by the lowercase name without the extension, and the
pipeline writes one 64x64 WebP per icon into `builds/<build>/icons/`.

A wowhead supplement item (`pipeline/wowhead_items.py`, the ids the client's
`ItemSparse` lacks) names its own icon the same way -- a lowercase name with
no extension -- but the client has no file data id for one it never shipped.
`file_id_by_name` is the reverse of `icon_names`, for looking a supplement
icon's name back up to a CASC id when the client happens to carry the art
under a different item; a name that resolves nowhere falls back to wowhead's
own hosted copy (`download_zamimg_icons`).
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

#: Wowhead's own hosted icon art, for a supplement item's icon name the
#: client's ManifestInterfaceData does not carry at all -- the same shape of
#: gap the supplement itself exists to fill (`pipeline/wowhead_items.py`'s
#: module docstring). Not versioned by client build: unlike a CASC file data
#: id, this is a name-addressed URL wowhead itself hosts, so the same name
#: fetches the same art regardless of which client build asked for it.
ZAMIMG_ICON_URL = "https://wow.zamimg.com/images/wow/icons/large/{name}.jpg"
_ZAMIMG_CACHE_SUBDIR = "zamimg"

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


def resolve_icon_name(
    base_icon: str,
    item_id: int,
    fork_icons: dict[int, str],
    wowhead_icons: dict[int, str],
) -> tuple[str, str]:
    """`base_icon` (whatever `resolve_icon` already produced from the client's
    own tables), with a fallback chain applied when it is still the
    placeholder. Returns `(icon_name, origin)`; `origin` is one of "client",
    "fork" or "wowhead", for a caller that logs or counts where each item's
    icon actually came from.

    771 of the 9,391 real items on build 1.60.1.70009 carry
    `IconFileDataID` 0 in the client's own `Item` table (night-icons finding,
    2026-09-29) -- not a join this pipeline is missing (`ManifestInterfaceData`
    resolves every nonzero id it is asked for), but Blizzard genuinely stating
    no icon for an item it shipped as a Forever hotfix (768 of the 771 are
    Forever-new ids, `>= 200_000`). Two other sources still know a real icon
    for the same item: the engine fork's own `assets/database/db.json`
    (`pipeline.forkdb`, curated from AtlasLoot), tried first as the more
    deliberately-curated of the two; then wowhead's Forever gear-planner
    payload (`pipeline.wowhead_items`), which names one for all 771 on that
    build and is tried last precisely because it is the least curated -- a
    wowhead-only fallback for an item the client and the fork both agree has
    no icon is trusted over neither disagreeing at all, but the fork's own
    say-so beats it when the fork has one.
    """
    if base_icon not in (PLACEHOLDER_ICON, "0", ""):
        return base_icon, "client"
    fork_icon = fork_icons.get(item_id)
    if fork_icon and fork_icon != PLACEHOLDER_ICON:
        return fork_icon, "fork"
    wowhead_icon = wowhead_icons.get(item_id)
    # The planner payload writes a literal "0" for an item wowhead has no
    # icon for either (Rotmender's set, "DNT"/"Template Item" test rows on
    # 1.60.1.70009); that is no icon, not an icon named "0".
    if wowhead_icon and wowhead_icon not in (PLACEHOLDER_ICON, "0"):
        return wowhead_icon, "wowhead"
    # Nothing named a real icon: a stored "0"/"" is normalised to the one
    # placeholder the site knows how to draw around.
    return (PLACEHOLDER_ICON if base_icon in ("0", "") else base_icon), "client"


def icon_names(manifest_rows: list[dict[str, str]]) -> dict[int, str]:
    """File data id -> lowercase icon name with no extension."""
    names: dict[int, str] = {}
    for row in manifest_rows:
        file_name = row["FileName"]
        if not file_name.lower().endswith(_BLP_SUFFIX):
            continue
        names[int(row["ID"])] = file_name[: -len(_BLP_SUFFIX)].lower()
    return names


def file_id_by_name(manifest_rows: list[dict[str, str]]) -> dict[str, int]:
    """Lowercase icon name -> file data id, the reverse of `icon_names`.

    A supplement item names its icon the way `icon_names` already normalises
    the client's own names (lowercase, no extension), so looking one up here
    is case-insensitive by construction rather than by a separate compare.
    Two file ids can share a lowercase name; the lower id wins, the same rule
    `wanted_icons`/`download_icons` apply to a collision.
    """
    by_name: dict[str, list[int]] = {}
    for file_id, name in icon_names(manifest_rows).items():
        by_name.setdefault(name, []).append(file_id)
    return {name: min(ids) for name, ids in by_name.items()}


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
    from pipeline.casc import CascMissing, CascMissingReason, fetch_casc_file

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
                # CascMissing covers both a 404 and an empty body; branch on
                # its structured `reason` rather than its message text, so a
                # reword of either message in casc.py can't silently misroute
                # this. The wording below is this module's own, kept
                # byte-for-byte what it was before the migration because
                # tests/test_icons.py matches on the empty-body phrasing
                # specifically.
                if error.reason is CascMissingReason.EMPTY:
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


def download_zamimg_icons(
    names: set[str],
    out_dir: Path,
    cache_dir: Path = CACHE_DIR,
    client: httpx.Client | None = None,
) -> int:
    """Write out_dir/<name>.webp for a supplement icon name that has no CASC
    file id at all -- the client never shipped it (`file_id_by_name` found no
    entry). Returns the number written.

    Cached at cache_dir/zamimg/<name>.jpg, unversioned: unlike a CASC file
    data id the same name always fetches the same wowhead-hosted image, so a
    rerun for a later build reuses the same cached bytes rather than
    refetching them.
    """
    own = client is None
    if client is None:
        client = httpx.Client(headers={"User-Agent": USER_AGENT})
    out_dir.mkdir(parents=True, exist_ok=True)
    zamimg_cache = cache_dir / _ZAMIMG_CACHE_SUBDIR
    zamimg_cache.mkdir(parents=True, exist_ok=True)
    written = 0
    try:
        for name in sorted(names):
            target = out_dir / f"{name}.webp"
            if target.exists():
                continue
            cached = zamimg_cache / f"{name}.jpg"
            if cached.exists():
                jpeg = cached.read_bytes()
            else:
                response = client.get(ZAMIMG_ICON_URL.format(name=name))
                if response.status_code != 200 or not response.content:
                    logger.warning(
                        "icon %s is not on zamimg (status %s); skipping",
                        name,
                        response.status_code,
                    )
                    continue
                jpeg = response.content
                _atomic_write(cached, jpeg)
            _atomic_write(target, blp_to_webp(jpeg))
            written += 1
    finally:
        if own:
            client.close()
    return written


def class_icon_names(build_dir: Path) -> set[str]:
    """classicon_<slug> for every class build_dir/classes.json lists.

    The report view shows one of these beside every character's name, but no
    talent or item ever references one, so _referenced_names cannot derive
    them the way it derives everything else: they have to come straight from
    the class list. This is the one source for that list; forever.py's
    pre-beta path uses it too rather than keeping its own copy of the class
    slugs, which would otherwise be free to drift from classes.json.
    """
    payload = json.loads((build_dir / "classes.json").read_text(encoding="utf-8"))
    return {f"classicon_{klass['slug']}" for klass in payload}


def _referenced_names(build_dir: Path, extra: frozenset[str] = frozenset()) -> set[str]:
    """Every icon name the already-emitted JSON under build_dir refers to.

    The normalizers already resolved each talent's and each item's icon (and
    already substituted PLACEHOLDER_ICON where the client had nothing), so the
    emitted `icon` is the single source of truth for what has to be
    downloaded there. Class icons are the exception: nothing in talents or
    items references them, so they are added separately from classes.json
    (see class_icon_names). Each talent TREE's own icon (the talent tab's
    icon, day3 data-followups-11 lane) is read here too, alongside its
    talents -- it lives on the same `talents/<class>.json` row.

    `extra` folds in an icon set this function cannot derive from the
    already-emitted JSON alone -- a rotation line's icon
    (`rotation_icon_names`) and a curated spec's own tab icon
    (`spec_icon_names`), both day3 data-followups-11 lane additions.
    """
    names: set[str] = set()
    for path in sorted((build_dir / "talents").glob("*.json")):
        payload = json.loads(path.read_text(encoding="utf-8"))
        for tree in payload["trees"]:
            names.add(tree["icon"])
            for talent in tree["talents"]:
                names.add(talent["icon"])
    for path in sorted((build_dir / "items").glob("*.json")):
        payload = json.loads(path.read_text(encoding="utf-8"))
        for item in payload["items"]:
            names.add(item["icon"])
    names |= class_icon_names(build_dir)
    names |= extra
    return names


def rotation_icon_names(build_dir: Path, curated_dir: Path = Path("curated")) -> set[str]:
    """Every icon name `pipeline.addonrotation.build_rotations` resolves for
    this build's rotation lines.

    Unlike a talent's or an item's icon, a rotation line's icon is not sitting
    in an already-emitted file under build_dir at the point `icons_for_build`
    runs -- `addon-data.json` (the file that would carry it) is written by a
    LATER pipeline step (`.github/workflows/data.yml` runs `icons` before
    `addon-data`, so this step's own output cannot be a dependency of this
    one), so it is recomputed here the same way `build_addon_data` itself
    will, rather than read back off a file.
    """
    from pipeline.addonrotation import build_rotations

    rotations = build_rotations(build_dir.parent, build_dir.name, curated_dir=curated_dir)
    return {line.icon for bands in rotations.values() for band in bands for line in band.lines}


def spec_icon_names(curated_dir: Path = Path("curated")) -> set[str]:
    """Every icon name `curated/specs.json` itself names (day3 data-
    followups-11 lane): each spec's own talent-tab icon, hand-curated
    there from the same client `TalentTab.SpellIconID` a build's own
    `talents/<class>.json` tree icon resolves (see specs.json's own
    comment) -- named again here, rather than only relying on it already
    being `_referenced_names`' talent-tree set, so a future curated value
    that drifts from the build still gets its icon downloaded rather than
    silently 404ing on the site.
    """
    path = curated_dir / "specs.json"
    if not path.exists():
        return set()
    return {entry["icon"] for entry in json.loads(path.read_text(encoding="utf-8"))}


def wanted_icons(
    build_dir: Path, names: dict[int, str], extra: frozenset[str] = frozenset()
) -> dict[int, str]:
    """File id -> icon name for every icon the emitted JSON under build_dir
    refers to, plus `extra` (see `_referenced_names`)."""
    referenced = _referenced_names(build_dir, extra)
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


def supplement_icon_names(raw: Path) -> set[str]:
    """Icon names the wowhead supplement's own items name (`w.icon`), for a
    build's raw/ directory. Empty when the build never fetched wowhead's
    payload -- `fetch-wowhead` is a workflow step, not a hard requirement, so
    a build's icons must still resolve without it.

    Local imports: `pipeline.wowhead_items` and `pipeline.normalize.gear`
    both import from this module (`resolve_icon`), so importing either at
    module scope here would re-enter icons.py while it is still initialising.
    """
    from pipeline import wowhead_items as wh
    from pipeline.csvio import read_csv
    from pipeline.normalize.gear import int_column

    path = raw / wh.RAW_FILE
    if not path.exists():
        return set()
    client_ids = {int_column(row, "ID") for row in read_csv(raw / "ItemSparse.csv")}
    supplement = wh.supplement(wh.load_items(path), client_ids)
    return {item.icon for item in supplement}


def icons_for_build(
    build: str,
    root: Path = Path("builds"),
    cache_dir: Path = CACHE_DIR,
    client: httpx.Client | None = None,
    curated_dir: Path = Path("curated"),
) -> int:
    from pipeline.csvio import read_csv

    build_dir = root / build
    raw = build_dir / "raw"
    if not raw.exists():
        raise SystemExit(f"no raw data at {raw}; run `python -m pipeline fetch` first")
    manifest_rows = read_csv(raw / "ManifestInterfaceData.csv")
    names = icon_names(manifest_rows)
    extra = rotation_icon_names(build_dir, curated_dir) | spec_icon_names(curated_dir)
    wanted = wanted_icons(build_dir, names, frozenset(extra))

    # A supplement icon the client happens to carry (under whatever item it
    # was originally shipped for) goes through the same CASC path as every
    # other icon; a name the client has nothing for at all falls back to
    # wowhead's own hosted copy. Both are additive to what wanted_icons
    # already found by reading the emitted JSON, since the wowhead items
    # (`normalize`'s side of this feature) may not have landed in
    # items/<class>.json in every build this runs against.
    by_name = file_id_by_name(manifest_rows)
    supplement_names = supplement_icon_names(raw)
    wanted = {
        **wanted,
        **{by_name[name]: name for name in supplement_names if name in by_name},
    }
    zamimg_names = {name for name in supplement_names if name not in by_name}

    written = download_icons(wanted, build_dir / "icons", cache_dir, client, version=build)
    written += download_zamimg_icons(zamimg_names, build_dir / "icons", cache_dir, client)
    print(f"{written} icons written, {len(wanted) + len(zamimg_names)} referenced")
    return written
