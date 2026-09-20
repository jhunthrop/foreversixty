"""A per-class item file may only name an id its own build's `ItemSparse` carries.

Every build's `items/<class-slug>.json` is normally generated straight from that
build's own `ItemSparse` (`pipeline.normalize.gear.build_class_items`), which makes
this trivially true. `forever-prebeta` is the one exception: `pipeline.forever`
copies its `items/` from a DIFFERENT build's tables because the Forever build fetches
no items of its own (see that module's docstring), so its ids need filtering against
the build the site actually serves rather than the one they were copied from --
`pipeline.forever._copy_class_items` is where that filter lives. Item 16963, "Helm of
Wrath", is the id that leaked before the filter existed: it has a row in Classic Era's
`ItemSparse` but not in the 1.60 client's, and the simulator refused it outright (see
web/src/lib/sim/sim-items.ts's comment on the same id).

No committed file carries `ItemSparse` directly (raw CSVs are gitignored), so this
uses each build's own `items.json` as the join key -- it is a straight `ItemSparse` x
`Item` join (`pipeline.normalize.items.normalize_items`), so its id set is exactly
what `items/<class-slug>.json` may name.
"""

import json
from pathlib import Path

BUILDS_DIR = Path("builds")
ACTIVE_BUILD_FILE = Path("..") / "web" / "src" / "data" / "active-build.json"


def _item_sparse_ids(build: str) -> set[int]:
    payload = json.loads((BUILDS_DIR / build / "items.json").read_text(encoding="utf-8"))
    return {item["id"] for item in payload}


def _per_class_items(build: str) -> dict[str, list[dict]]:
    items_dir = BUILDS_DIR / build / "items"
    return {
        path.stem: json.loads(path.read_text(encoding="utf-8"))["items"]
        for path in sorted(items_dir.glob("*.json"))
    }


def _builds_with_items() -> list[str]:
    return sorted(
        build_dir.name
        for build_dir in BUILDS_DIR.iterdir()
        if (build_dir / "items").is_dir() and (build_dir / "items.json").is_file()
    )


def test_every_id_in_every_per_class_file_has_an_itemsparse_row():
    """Pinned against every committed build that has both items/ and items.json."""
    builds = _builds_with_items()
    assert builds  # guard: an empty build list would make the loop vacuously true
    for build in builds:
        valid_ids = _item_sparse_ids(build)
        for slug, items in _per_class_items(build).items():
            unknown = sorted({item["id"] for item in items} - valid_ids)
            assert unknown == [], (build, slug, unknown[:10])


def test_helm_of_wrath_does_not_leak_into_forever_prebetas_item_files():
    """The specific id that motivated the filter: real in Classic Era, absent from
    the 1.60 client the simulator actually runs against."""
    active_build = json.loads(ACTIVE_BUILD_FILE.read_text(encoding="utf-8"))["build"]
    assert 16963 not in _item_sparse_ids(active_build)
    for items in _per_class_items("forever-prebeta").values():
        assert all(item["id"] != 16963 for item in items)


#: golden regression pin for pipeline.forever._copy_class_items: how many of
#: 1.15.9.69722's own per-class items have no row in the active build's ItemSparse,
#: per class. A change here means either the source build's items/ changed or the
#: active build's items.json did -- both are real pipeline events, not test noise.
GOLDEN_DROPPED_PER_CLASS = {
    "druid": 4271,
    "hunter": 5470,
    "mage": 2958,
    "paladin": 6122,
    "priest": 2972,
    "rogue": 4076,
    "shaman": 5561,
    "warlock": 2949,
    "warrior": 6593,
}


def test_forever_prebetas_dropped_item_count_is_pinned_per_class():
    era_counts = {slug: len(items) for slug, items in _per_class_items("1.15.9.69722").items()}
    prebeta_counts = {
        slug: len(items) for slug, items in _per_class_items("forever-prebeta").items()
    }
    dropped = {slug: era_counts[slug] - prebeta_counts[slug] for slug in era_counts}
    assert dropped == GOLDEN_DROPPED_PER_CLASS
