import json
import warnings
from pathlib import Path

import pytest

from pipeline.curated import CuratedError
from pipeline.loot.overlay import OverlayError, apply_overlays, load_overlays
from pipeline.models import LootBoss, LootFile, LootSource

HERE = Path(__file__).parent
FIXTURE = HERE / "fixtures/loot/curated"
CURATED = Path("curated/loot")


def base() -> LootFile:
    return LootFile(
        sources=[
            LootSource(
                id="raid:molten-core",
                kind="raid",
                name="Molten Core",
                zone_id=2717,
                bosses=[LootBoss(id="raid:molten-core:900", name="Big Boss", npc_id=900,
                                 items=[100])],
                trash=[101],
            ),
            LootSource(
                id="dungeon:the-deadmines",
                kind="dungeon",
                name="The Deadmines",
                zone_id=1581,
                bosses=[LootBoss(id="dungeon:the-deadmines:902", name="Dungeon Boss",
                                 npc_id=902, items=[103])],
            ),
        ]
    )


def write(tmp_path: Path, name: str, payload: dict) -> Path:
    (tmp_path / name).write_text(json.dumps(payload), encoding="utf-8")
    return tmp_path


def test_a_missing_overlay_directory_is_not_an_error(tmp_path):
    assert load_overlays(tmp_path / "nope") == []


def test_the_fixture_overlays_add_patch_and_remove():
    result = apply_overlays(base(), load_overlays(FIXTURE))
    by_id = {source.id: source for source in result.sources}
    assert set(by_id) == {"raid:molten-core", "raid:barrow-deeps"}
    assert by_id["raid:molten-core"].opens == "later"
    assert by_id["raid:barrow-deeps"].items == [100]


def test_a_replace_only_touches_the_keys_it_names():
    result = apply_overlays(base(), load_overlays(FIXTURE))
    molten = next(s for s in result.sources if s.id == "raid:molten-core")
    assert molten.name == "Molten Core"
    assert molten.zone_id == 2717
    assert molten.trash == [101]
    assert [b.items for b in molten.bosses] == [[100]]


def test_the_result_stays_sorted_by_kind_then_id():
    result = apply_overlays(base(), load_overlays(FIXTURE))
    assert [s.id for s in result.sources] == ["raid:barrow-deeps", "raid:molten-core"]


def test_an_overlay_with_no_source_is_refused(tmp_path):
    write(tmp_path, "a.json", {"notes": "x", "add": []})
    with pytest.raises(CuratedError):
        load_overlays(tmp_path)


def test_an_overlay_source_kind_outside_the_vocabulary_is_refused(tmp_path):
    write(tmp_path, "a.json", {
        "sources": [{"label": "l", "url": "u", "kind": "hearsay"}], "notes": "x"})
    with pytest.raises(CuratedError, match="hearsay"):
        load_overlays(tmp_path)


def test_adding_a_source_that_already_exists_is_refused(tmp_path):
    write(tmp_path, "a.json", {
        "sources": [{"label": "l", "url": "https://x.invalid", "kind": "site"}],
        "notes": "x",
        "add": [{"id": "raid:molten-core", "kind": "raid", "name": "Again"}],
    })
    with pytest.raises(OverlayError, match="already"):
        apply_overlays(base(), load_overlays(tmp_path))


def test_replacing_or_removing_an_unknown_source_is_refused(tmp_path):
    for payload in (
        {"replace": [{"id": "raid:nope", "opens": "later"}]},
        {"remove": ["raid:nope"]},
    ):
        write(tmp_path, "a.json", {
            "sources": [{"label": "l", "url": "https://x.invalid", "kind": "site"}],
            "notes": "x",
        } | payload)
        with pytest.raises(OverlayError, match="raid:nope"):
            apply_overlays(base(), load_overlays(tmp_path))


def test_a_replace_may_restate_a_sources_item_lists(tmp_path):
    # Contract 10.4 has no per-item removal, so this is how a wrong item
    # is corrected: restate the list it is in.
    write(tmp_path, "a.json", {
        "sources": [{"label": "l", "url": "https://x.invalid", "kind": "site"}],
        "notes": "x",
        "replace": [{"id": "raid:molten-core", "trash": []}],
    })
    result = apply_overlays(base(), load_overlays(tmp_path))
    molten = next(s for s in result.sources if s.id == "raid:molten-core")
    assert molten.trash == []
    assert [b.items for b in molten.bosses] == [[100]]


def test_a_replace_may_not_change_a_sources_kind(tmp_path):
    write(tmp_path, "a.json", {
        "sources": [{"label": "l", "url": "https://x.invalid", "kind": "site"}],
        "notes": "x",
        "replace": [{"id": "raid:molten-core", "kind": "dungeon"}],
    })
    with pytest.raises(OverlayError, match="kind"):
        apply_overlays(base(), load_overlays(tmp_path))


def test_a_replace_touching_bosses_yields_real_loot_boss_instances(tmp_path):
    # model_copy(update=...) assigns a patch's dumped dict verbatim, without
    # re-validating; a `replace` that names `bosses` must still come back
    # out as `LootBoss` instances, not raw dicts, since `LootSourcePatch`
    # explicitly advertises `bosses` as a replaceable field.
    write(tmp_path, "a.json", {
        "sources": [{"label": "l", "url": "https://x.invalid", "kind": "site"}],
        "notes": "x",
        "replace": [{
            "id": "raid:molten-core",
            "bosses": [
                {"id": "raid:molten-core:900", "name": "New Boss", "npc_id": 900,
                 "items": [100, 200]}
            ],
        }],
    })
    result = apply_overlays(base(), load_overlays(tmp_path))
    molten = next(s for s in result.sources if s.id == "raid:molten-core")
    assert len(molten.bosses) == 1
    boss = molten.bosses[0]
    assert isinstance(boss, LootBoss)
    assert boss.name == "New Boss"
    assert boss.npc_id == 900
    assert boss.items == [100, 200]
    with warnings.catch_warnings():
        warnings.simplefilter("error")
        molten.model_dump()


def test_the_real_curated_overlays_all_load_and_apply_cleanly():
    """The committed overlays must be loadable on their own terms; that
    they match the committed loot.json is tests/test_loot_build.py's job."""
    for path, document in load_overlays(CURATED):
        assert document.sources, path
        assert document.notes.strip(), path


#: The raid sources the generator emits for build 1.60.1.69893 after
#: contract 10.4's build filter. Onyxia's Lair is absent -- all sixteen
#: of her fork-database drops are gone from the client -- which is why
#: the overlay adds her rather than patching her.
GENERATED_RAID_IDS = [
    "raid:ahnqiraj",
    "raid:blackwing-lair",
    "raid:molten-core",
    "raid:naxxramas",
    "raid:ruins-of-ahnqiraj",
    "raid:zulgurub",
]
#: api/internal/phase/phase.go's names, plus contract 10.4's sentinel.
PHASES = {"pre-beta", "beta", "launch", "raids-1"}
OPENS_LATER = "later"


def raid_phase_overlay():
    for path, document in load_overlays(CURATED):
        if path.name == "forever-raid-phases.json":
            return document
    raise AssertionError("no forever-raid-phases.json under curated/loot")


def test_the_six_generated_raids_are_gated_as_unreleased_without_a_date():
    document = raid_phase_overlay()
    assert sorted(patch.id for patch in document.replace) == GENERATED_RAID_IDS
    assert {patch.opens for patch in document.replace} == {OPENS_LATER}


def test_onyxia_is_added_with_the_announced_date_and_no_items():
    """The one raid the phase calendar has a date for, and the one the
    build filter removed entirely. An empty item list is the honest
    statement: the raid is announced, its loot table is not known."""
    document = raid_phase_overlay()
    assert [source.id for source in document.add] == ["raid:onyxias-lair"]
    onyxia = document.add[0]
    assert (onyxia.kind, onyxia.zone_id, onyxia.opens) == ("raid", 2159, "raids-1")
    assert onyxia.items == []


def test_the_raid_phase_overlay_removes_nothing():
    assert raid_phase_overlay().remove == []


def test_the_notes_state_the_per_raid_gap():
    """Contract 10.4: 'the first overlay records the gap per raid in its
    notes'. Spot-check the two extremes rather than the whole table."""
    notes = raid_phase_overlay().notes
    for fragment in ("Molten Core", "10 of 139", "Zul'Gurub", "2 of 98"):
        assert fragment in notes


def test_every_phase_an_overlay_names_is_a_real_phase_or_the_sentinel():
    """An `opens` the API does not know is a filter that matches nothing."""
    allowed = PHASES | {OPENS_LATER}
    for path, document in load_overlays(CURATED):
        for patch in document.replace:
            assert patch.opens is None or patch.opens in allowed, (path, patch.id)
        for source in document.add:
            assert source.opens is None or source.opens in allowed, (path, source.id)
