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
                bosses=[
                    LootBoss(id="raid:molten-core:900", name="Big Boss", npc_id=900, items=[100])
                ],
                trash=[101],
            ),
            LootSource(
                id="dungeon:the-deadmines",
                kind="dungeon",
                name="The Deadmines",
                zone_id=1581,
                bosses=[
                    LootBoss(
                        id="dungeon:the-deadmines:902", name="Dungeon Boss", npc_id=902, items=[103]
                    )
                ],
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
    write(
        tmp_path,
        "a.json",
        {"sources": [{"label": "l", "url": "u", "kind": "hearsay"}], "notes": "x"},
    )
    with pytest.raises(CuratedError, match="hearsay"):
        load_overlays(tmp_path)


def test_adding_a_source_that_already_exists_is_refused(tmp_path):
    write(
        tmp_path,
        "a.json",
        {
            "sources": [{"label": "l", "url": "https://x.invalid", "kind": "site"}],
            "notes": "x",
            "add": [{"id": "raid:molten-core", "kind": "raid", "name": "Again"}],
        },
    )
    with pytest.raises(OverlayError, match="already"):
        apply_overlays(base(), load_overlays(tmp_path))


def test_replacing_or_removing_an_unknown_source_is_refused(tmp_path):
    for payload in (
        {"replace": [{"id": "raid:nope", "opens": "later"}]},
        {"remove": ["raid:nope"]},
    ):
        write(
            tmp_path,
            "a.json",
            {
                "sources": [{"label": "l", "url": "https://x.invalid", "kind": "site"}],
                "notes": "x",
            }
            | payload,
        )
        with pytest.raises(OverlayError, match="raid:nope"):
            apply_overlays(base(), load_overlays(tmp_path))


def test_a_replace_may_restate_a_sources_item_lists(tmp_path):
    # Contract 10.4 has no per-item removal, so this is how a wrong item
    # is corrected: restate the list it is in.
    write(
        tmp_path,
        "a.json",
        {
            "sources": [{"label": "l", "url": "https://x.invalid", "kind": "site"}],
            "notes": "x",
            "replace": [{"id": "raid:molten-core", "trash": []}],
        },
    )
    result = apply_overlays(base(), load_overlays(tmp_path))
    molten = next(s for s in result.sources if s.id == "raid:molten-core")
    assert molten.trash == []
    assert [b.items for b in molten.bosses] == [[100]]


def test_a_replace_may_not_change_a_sources_kind(tmp_path):
    write(
        tmp_path,
        "a.json",
        {
            "sources": [{"label": "l", "url": "https://x.invalid", "kind": "site"}],
            "notes": "x",
            "replace": [{"id": "raid:molten-core", "kind": "dungeon"}],
        },
    )
    with pytest.raises(OverlayError, match="kind"):
        apply_overlays(base(), load_overlays(tmp_path))


def test_a_replace_touching_bosses_yields_real_loot_boss_instances(tmp_path):
    # model_copy(update=...) assigns a patch's dumped dict verbatim, without
    # re-validating; a `replace` that names `bosses` must still come back
    # out as `LootBoss` instances, not raw dicts, since `LootSourcePatch`
    # explicitly advertises `bosses` as a replaceable field.
    write(
        tmp_path,
        "a.json",
        {
            "sources": [{"label": "l", "url": "https://x.invalid", "kind": "site"}],
            "notes": "x",
            "replace": [
                {
                    "id": "raid:molten-core",
                    "bosses": [
                        {
                            "id": "raid:molten-core:900",
                            "name": "New Boss",
                            "npc_id": 900,
                            "items": [100, 200],
                        }
                    ],
                }
            ],
        },
    )
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
#: contract 10.4's build filter. Onyxia's Lair's own sixteen fork-database
#: drops are all gone from the client, but classic-db's own creature_loot_
#: template for Onyxia (npc 10184) still names a real table once she stops
#: being stranded on `world:onyxia` (src-classicdb-fixes lane, 2026-09-29:
#: `fork_instance_npc_zones`' fallback for a scripted boss with no static
#: spawn row) -- so the generator emits `raid:onyxias-lair` on its own now,
#: and every one of the seven raids, Onyxia included, is patched rather
#: than added.
GENERATED_RAID_IDS = [
    "raid:ahnqiraj",
    "raid:blackwing-lair",
    "raid:molten-core",
    "raid:naxxramas",
    "raid:onyxias-lair",
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


def test_the_seven_generated_raids_are_all_patched_not_added():
    """Every raid the generator itself emits -- Onyxia's Lair included,
    src-classicdb-fixes lane, 2026-09-29 -- gets only its phase patched
    here, never re-added: `apply_overlays.add` is an error on a source
    that already exists."""
    document = raid_phase_overlay()
    assert sorted(patch.id for patch in document.replace) == GENERATED_RAID_IDS
    assert document.add == []


def test_six_raids_are_gated_as_unreleased_and_onyxia_carries_the_announced_date():
    document = raid_phase_overlay()
    by_id = {patch.id: patch for patch in document.replace}
    assert by_id["raid:onyxias-lair"].opens == "raids-1"
    unreleased = {
        patch.id: patch.opens for patch in document.replace if patch.id != "raid:onyxias-lair"
    }
    assert set(unreleased.values()) == {OPENS_LATER}
    assert len(unreleased) == 6


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


def test_an_undated_raid_source_defaults_to_opens_later():
    from pipeline.loot.overlay import RAID_DEFAULT_OPENS, apply_overlays
    from pipeline.models import LootFile, LootSource

    document = LootFile(
        sources=[
            LootSource(id="raid:scarlet-enclave", kind="raid", name="Scarlet Enclave", items=[1]),
            LootSource(id="dungeon:the-deadmines", kind="dungeon", name="The Deadmines", items=[2]),
        ],
        quests={},
        factions={},
    )
    out = apply_overlays(document, [])
    by_id = {s.id: s for s in out.sources}
    assert by_id["raid:scarlet-enclave"].opens == RAID_DEFAULT_OPENS
    assert by_id["dungeon:the-deadmines"].opens is None


def _world_boss_document() -> LootFile:
    """A synthetic build carrying all six named world bosses, plus one
    ordinary `world:` source that must NOT be swept up by the same gate
    -- most `world:` buckets ARE launch-farmable, only the six
    WORLD_BOSS_SOURCE_IDS names are held back for the first raid tier."""
    from pipeline.loot.overlay import WORLD_BOSS_SOURCE_IDS

    return LootFile(
        sources=[
            LootSource(
                id=source_id, kind="world", name=source_id.removeprefix("world:"), items=[100 + i]
            )
            for i, source_id in enumerate(sorted(WORLD_BOSS_SOURCE_IDS))
        ]
        + [
            LootSource(
                id="world:some-farmable-mob", kind="world", name="Some Farmable Mob", items=[200]
            ),
        ],
        quests={},
        factions={},
    )


def test_the_world_boss_ids_are_exactly_the_six_named_bosses():
    from pipeline.loot.overlay import WORLD_BOSS_NPC_NAMES, WORLD_BOSS_SOURCE_IDS

    assert set(WORLD_BOSS_NPC_NAMES) == {12397, 6109, 14889, 14888, 14890, 14887}
    assert WORLD_BOSS_SOURCE_IDS == {
        "world:lord-kazzak",
        "world:azuregos",
        "world:emeriss",
        "world:lethon",
        "world:taerar",
        "world:ysondre",
    }


def test_all_six_named_world_bosses_default_to_the_raids_1_phase():
    from pipeline.loot.overlay import (
        RAID_DEFAULT_OPENS,
        WORLD_BOSS_DEFAULT_OPENS,
        WORLD_BOSS_SOURCE_IDS,
    )

    out = apply_overlays(_world_boss_document(), [])
    by_id = {s.id: s for s in out.sources}
    for source_id in WORLD_BOSS_SOURCE_IDS:
        assert by_id[source_id].opens == WORLD_BOSS_DEFAULT_OPENS, source_id
    # Not the raid sentinel: a raid with no announced date at all gates to
    # "later"; these six gate to the KNOWN phase the first raids open in.
    assert WORLD_BOSS_DEFAULT_OPENS != RAID_DEFAULT_OPENS


def test_an_ordinary_world_source_is_left_ungated():
    out = apply_overlays(_world_boss_document(), [])
    by_id = {s.id: s for s in out.sources}
    assert by_id["world:some-farmable-mob"].opens is None


def test_a_curated_overlay_still_wins_over_the_world_boss_default(tmp_path):
    """`replace` runs before the default-fill loop, so a curated fact (a
    world boss server-first killed before the raid tier, say) is not
    clobbered back to WORLD_BOSS_DEFAULT_OPENS."""
    write(
        tmp_path,
        "a.json",
        {
            "sources": [{"label": "l", "url": "https://x.invalid", "kind": "site"}],
            "notes": "x",
            "replace": [{"id": "world:lord-kazzak", "opens": "launch"}],
        },
    )
    out = apply_overlays(_world_boss_document(), load_overlays(tmp_path))
    by_id = {s.id: s for s in out.sources}
    assert by_id["world:lord-kazzak"].opens == "launch"
