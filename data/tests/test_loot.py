"""`write_loot_files`: the orchestrator behind `make loot` and CI's `loot` step.

Like `test_normalize_build.py` and `test_simdb.py`, this drives the real
function over a copy of `tests/fixtures/loot/` in `tmp_path` rather than
mocking any of its pieces -- `overlay_dir`, `curated_dir` and `ids_md` are
parameters precisely so a test can redirect them without touching the
shared fixtures `tests/test_loot_sources.py`, `test_loot_gear.py`,
`test_loot_buffs.py` and `test_loot_overlay.py` all read.

`tests/test_loot_sources.py`'s `built()` documents what `build_loot` itself
emits for these fixtures: seven sources --
`raid:molten-core`, `dungeon:the-deadmines`, `world:azuregos`,
`crafted:blacksmithing`, `rep:argent-dawn:exalted`, `pvp:rank-11`, `quest`
-- which is the baseline every assertion below is checked against.
"""

from __future__ import annotations

import json
import shutil
from pathlib import Path

import pytest

from pipeline.loot import ENCHANTS, LOOT, SUFFIXES, write_loot_files
from pipeline.loot.buffs import SIMBUFFS
from pipeline.manifest import read_manifest, verify, write_manifest
from pipeline.models import LootFile, LootSource
from pipeline.normalize import write_document

HERE = Path(__file__).parent
FIXTURES = HERE / "fixtures/loot"
OVERLAY_FIXTURE = FIXTURES / "curated"
BUILD = "1.60.1.69893"

RAW_FILES = ("Map.csv", "ItemSparse.csv", "SpellMisc.csv", "Item.csv",
             "ManifestInterfaceData.csv")
JSON_FILES = ("zones.json", "items.json", "spells.json")

#: What `build_loot` itself emits for these fixtures, before any overlay --
#: see `test_loot_sources.py::test_every_kind_is_emitted_once_and_in_the_contracts_order`.
GENERATED_SOURCE_IDS = [
    "raid:molten-core",
    "dungeon:the-deadmines",
    "world:azuregos",
    "crafted:blacksmithing",
    "rep:argent-dawn:exalted",
    "pvp:rank-11",
    "quest",
]


def prepare_build(tmp_path: Path, *, sparse_rows: str | None = None) -> Path:
    """A `builds/<BUILD>/` directory shaped like one `normalize` produced:
    the three prerequisite JSON files, `raw/`'s CSVs, and a manifest to
    refresh -- `refresh_manifest` refuses a build with no manifest yet."""
    root = tmp_path / "builds"
    build_dir = root / BUILD
    raw = build_dir / "raw"
    raw.mkdir(parents=True)
    for name in RAW_FILES:
        if name == "ItemSparse.csv" and sparse_rows is not None:
            (raw / name).write_text(sparse_rows, encoding="utf-8")
        else:
            shutil.copy(FIXTURES / name, raw / name)
    for name in JSON_FILES:
        shutil.copy(FIXTURES / name, build_dir / name)
    write_manifest(
        build_dir, build=BUILD, product="wow_classic_beta", fetched_at="2026-01-01T00:00:00Z"
    )
    return root


def prepare_curated(tmp_path: Path) -> Path:
    """A `curated/` with `simbuffs.json` under the name `load_overrides`'s
    default parameter looks for -- the shared fixture carries the same
    content as `curated-simbuffs.json`, a name chosen so `test_loot_buffs.py`
    could pass a non-default `name=` to `load_overrides` directly."""
    curated = tmp_path / "curated"
    curated.mkdir()
    shutil.copy(FIXTURES / "curated-simbuffs.json", curated / "simbuffs.json")
    return curated


def run(tmp_path: Path, *, overlay_dir: Path = OVERLAY_FIXTURE, sparse_rows: str | None = None):
    root = prepare_build(tmp_path, sparse_rows=sparse_rows)
    return write_loot_files(
        BUILD,
        engine_dir=FIXTURES,
        root=root,
        overlay_dir=overlay_dir,
        curated_dir=prepare_curated(tmp_path),
        ids_md=FIXTURES / "IDS.md",
    ), root / BUILD


def test_a_missing_raw_directory_is_a_clear_error(tmp_path: Path):
    (tmp_path / "builds" / BUILD).mkdir(parents=True)
    with pytest.raises(SystemExit, match="fetch"):
        write_loot_files(
            BUILD,
            engine_dir=FIXTURES,
            root=tmp_path / "builds",
            overlay_dir=OVERLAY_FIXTURE,
            curated_dir=prepare_curated(tmp_path),
            ids_md=FIXTURES / "IDS.md",
        )


@pytest.mark.parametrize("missing", JSON_FILES)
def test_a_missing_normalize_output_is_a_clear_error(tmp_path: Path, missing: str):
    root = prepare_build(tmp_path)
    (root / BUILD / missing).unlink()
    with pytest.raises(SystemExit, match="normalize"):
        write_loot_files(
            BUILD,
            engine_dir=FIXTURES,
            root=root,
            overlay_dir=OVERLAY_FIXTURE,
            curated_dir=prepare_curated(tmp_path),
            ids_md=FIXTURES / "IDS.md",
        )


def test_check_no_sockets_runs_on_the_loot_path(tmp_path: Path):
    """The same guard `normalize` runs (parity design 4.4): a socketed item
    reaching `loot` must stop the run, not silently get a loot table."""
    sparse = (FIXTURES / "ItemSparse.csv").read_text(encoding="utf-8")
    socketed = sparse.replace(
        "110,Suffixed Sword,0,0,0,0", "110,Suffixed Sword,0,1,0,0"
    )
    assert socketed != sparse
    with pytest.raises(SystemExit, match="gem socket"):
        run(tmp_path, sparse_rows=socketed)


def test_write_loot_files_returns_the_five_paths_it_writes(tmp_path: Path):
    paths, build_dir = run(tmp_path)
    assert paths == [
        build_dir / LOOT,
        build_dir / ENCHANTS,
        build_dir / SUFFIXES,
        build_dir / SIMBUFFS,
        build_dir / "items.json",
    ]
    for path in paths:
        assert path.exists(), path


def test_the_generated_sources_survive_next_to_the_overlays_effects(tmp_path: Path):
    """The build-filter document (see `GENERATED_SOURCE_IDS`) laid over by
    the two fixture overlays: `10-add.json` adds `raid:barrow-deeps`,
    `20-replace-and-remove.json` gates Molten Core and drops the dungeon."""
    _, build_dir = run(tmp_path)
    document = json.loads((build_dir / LOOT).read_text(encoding="utf-8"))
    by_id = {source["id"]: source for source in document["sources"]}
    assert set(by_id) == (set(GENERATED_SOURCE_IDS) | {"raid:barrow-deeps"}) - {
        "dungeon:the-deadmines"
    }
    assert by_id["raid:barrow-deeps"]["items"] == [100]
    assert by_id["raid:molten-core"]["opens"] == "later"
    # The replace only touched `opens`; the rest of Molten Core is untouched.
    assert by_id["raid:molten-core"]["zone_id"] == 2717


def test_a_curated_source_with_an_empty_item_list_survives_because_the_build_filter_ran_first(
    tmp_path: Path,
):
    """`build_loot`'s pruning sweep (`sources = [s for s in sources if
    source_item_ids(s)]`) runs *inside* `build_loot`, before the overlay is
    applied -- see the comment in `pipeline/loot/__init__.py`. An overlay
    that adds a source with a deliberately empty item list (the real
    `forever-raid-phases.json` does this for the announced-but-unlooted
    Onyxia's Lair) must therefore still be in the final document: nothing
    downstream of the overlay re-runs that sweep."""
    overlay_dir = tmp_path / "empty-overlay"
    overlay_dir.mkdir()
    (overlay_dir / "announced.json").write_text(
        json.dumps(
            {
                "sources": [
                    {"label": "Announcement", "url": "https://example.invalid/a", "kind": "site"}
                ],
                "notes": "An announced raid with no known loot table yet.",
                "add": [
                    {"id": "raid:announced-only", "kind": "raid", "name": "Announced Raid",
                     "items": []}
                ],
            }
        ),
        encoding="utf-8",
    )
    _, build_dir = run(tmp_path, overlay_dir=overlay_dir)
    document = json.loads((build_dir / LOOT).read_text(encoding="utf-8"))
    by_id = {source["id"]: source for source in document["sources"]}
    assert by_id["raid:announced-only"]["items"] == []
    # And the ordinary pruning still applies to what build_loot itself
    # generates: nothing with an empty item list from the fork/client join
    # survives, because that sweep runs before this overlay is even loaded.
    assert all(
        source.get("items") or source.get("trash") or source.get("bosses")
        for source in document["sources"]
        if source["id"] != "raid:announced-only"
    )


def test_the_manifest_covers_all_five_writes_after_the_run(tmp_path: Path):
    _, build_dir = run(tmp_path)
    assert verify(build_dir) == []
    files = read_manifest(build_dir)["files"]
    for name in (LOOT, ENCHANTS, SUFFIXES, SIMBUFFS, "items.json"):
        assert name in files


def test_items_json_gains_both_fork_columns(tmp_path: Path):
    _, build_dir = run(tmp_path)
    rows = json.loads((build_dir / "items.json").read_text(encoding="utf-8"))
    by_id = {row["id"]: row for row in rows}
    assert by_id[110]["suffixes"] == [5, 6]
    assert by_id[100]["faction_restriction"] == "alliance_only"
    assert all("suffixes" in row and "faction_restriction" in row for row in rows)


def test_write_document_drops_an_unset_key_entirely_not_as_null(tmp_path: Path):
    """`loot.json`'s sources are a union (contract 6.1): a crafted source
    has no `bosses`. `write_document` must leave the key out of the JSON
    altogether, not write it as `null` -- a file of nulls is a file every
    consumer has to filter."""
    crafted = LootSource(id="crafted:tailoring", kind="crafted", name="Tailoring", items=[1])
    path = tmp_path / "loot.json"
    write_document(LootFile(sources=[crafted]), path)
    text = path.read_text(encoding="utf-8")
    parsed = json.loads(text)
    assert "bosses" not in parsed["sources"][0]
    assert '"bosses"' not in text
