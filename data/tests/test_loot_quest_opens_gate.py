# data/tests/test_loot_quest_opens_gate.py
"""Quest-gates lane, 2026-09-29: `pipeline.loot.sources.
apply_quest_opens_gate` gates a `QuestSource.opens` whenever the quest's
own classic-db turn-in item(s) are themselves only obtainable from a
source `opens` is already set on -- see that function's own doc, and the
lane's report, for the real-dump shape this pins (Onyxia's Lair's "Head
of Onyxia" gating "For All To See"/"Celebrating Good Times"; Ruins of
Ahn'Qiraj's "Head of Ossirian" gating "The Fall of Ossirian").

Everything here is built by hand rather than from a fixture engine dir --
`apply_quest_opens_gate` only ever reads `LootFile.sources`/`.quests` and
a `quest_id -> turn-in item ids` map, so a real fork/classic-db/wowhead
merge is not needed to exercise it.
"""

from pipeline.classic_sources import ClassicDbQuestInfo, ClassicDbSourceRecord
from pipeline.loot.sources import apply_quest_opens_gate
from pipeline.models import LootBoss, LootFile, LootSource, QuestSource

HEAD_OF_ONYXIA = 18422
ONYXIA_TOOTH_PENDANT = 18404
HEAD_OF_OSSIRIAN = 21220
CHARM_OF_THE_SHIFTING_SANDS = 21504
ORDINARY_TURN_IN = 9001  # a farmed/vendored item, never gated


def _quest_source(quest_id: int, name: str) -> QuestSource:
    return QuestSource(
        quest_id=quest_id, name=name, faction="both", min_level=60, level=60,
        level_source="classic-db",
    )  # fmt: skip


def _classic_sources(turn_ins: dict[int, list[int]]) -> dict[int, list[ClassicDbSourceRecord]]:
    """One `quest_reward` record per quest id, its own `turn_in_item_ids`
    already resolved (the shape `pipeline.classic_sources._parse_quest_
    rewards` produces) -- keyed under a throwaway reward item id, the
    same way `quest_turn_in_items_from_classic_sources` reads it (it
    scans every record regardless of which item id it is filed under)."""
    return {
        quest_id: [
            ClassicDbSourceRecord(
                kind="quest_reward", name=f"quest {quest_id}",
                quest=ClassicDbQuestInfo(
                    quest_id=quest_id, min_level=60, level=60, faction="both",
                    turn_in_item_ids=items,
                ),
            )
        ]
        for quest_id, items in turn_ins.items()
    }  # fmt: skip


def test_a_quest_whose_only_turn_in_item_is_a_raid_boss_drop_is_gated():
    document = LootFile(
        sources=[
            LootSource(
                id="raid:onyxias-lair", kind="raid", name="Onyxia's Lair", opens="raids-1",
                bosses=[LootBoss(id="onyxia", name="Onyxia", npc_id=1, items=[HEAD_OF_ONYXIA])],
            ),
        ],
        quests={
            str(ONYXIA_TOOTH_PENDANT): [_quest_source(7491, "For All To See")],
        },
    )
    classic_sources = _classic_sources({7491: [HEAD_OF_ONYXIA]})
    gated = apply_quest_opens_gate(document, classic_sources)
    assert gated.quests[str(ONYXIA_TOOTH_PENDANT)][0].opens == "raids-1"


def test_a_quest_whose_turn_in_item_also_has_an_ungated_way_stays_open():
    """A turn-in item's raid drop is not its only source -- a vendor also
    sells it -- so completing the quest never actually requires the
    raid; tenet 8's own "never invented" rule means this must stay open,
    not gated on a coincidence."""
    document = LootFile(
        sources=[
            LootSource(
                id="raid:onyxias-lair", kind="raid", name="Onyxia's Lair", opens="raids-1",
                bosses=[LootBoss(id="onyxia", name="Onyxia", npc_id=1, items=[HEAD_OF_ONYXIA])],
            ),
            LootSource(
                id="vendor:1", kind="vendor", name="A Vendor", npc_id=1, items=[HEAD_OF_ONYXIA],
            ),
        ],
        quests={str(ONYXIA_TOOTH_PENDANT): [_quest_source(7491, "For All To See")]},
    )
    classic_sources = _classic_sources({7491: [HEAD_OF_ONYXIA]})
    gated = apply_quest_opens_gate(document, classic_sources)
    assert gated.quests[str(ONYXIA_TOOTH_PENDANT)][0].opens is None


def test_a_quest_needing_an_ordinary_farmed_turn_in_item_stays_open():
    """The overwhelming majority case: `ReqItemId` names a plain world
    drop this build's document has no `opens` on at all -- never gated."""
    document = LootFile(
        sources=[
            LootSource(
                id="world:some-mob", kind="world", name="Some Mob", items=[ORDINARY_TURN_IN],
            ),
        ],
        quests={str(ONYXIA_TOOTH_PENDANT): [_quest_source(7491, "For All To See")]},
    )
    classic_sources = _classic_sources({7491: [ORDINARY_TURN_IN]})
    gated = apply_quest_opens_gate(document, classic_sources)
    assert gated.quests[str(ONYXIA_TOOTH_PENDANT)][0].opens is None


def test_a_turn_in_item_that_is_itself_a_gated_quest_reward_gates_recursively():
    """Item-level recursion: quest B's own turn-in item (`INTERMEDIATE`)
    names NO direct source in `document.sources` at all -- the only way
    to get it is quest A, which is itself gated on a raid boss drop
    directly. Quest B must inherit quest A's gate without a second
    classic-db pass ("or another such quest, recursively", this
    function's own doc)."""
    intermediate = 30001  # quest A's own reward; quest B's own turn-in item
    document = LootFile(
        sources=[
            LootSource(
                id="raid:ruins-of-ahnqiraj", kind="raid", name="Ruins of Ahn'Qiraj",
                opens="later",
                bosses=[
                    LootBoss(id="ossirian", name="Ossirian", npc_id=2, items=[HEAD_OF_OSSIRIAN]),
                ],
            ),
        ],
        quests={
            str(intermediate): [_quest_source(9001, "Quest A")],
            str(CHARM_OF_THE_SHIFTING_SANDS): [_quest_source(8791, "Quest B")],
        },
    )
    classic_sources = _classic_sources(
        {9001: [HEAD_OF_OSSIRIAN], 8791: [intermediate]}
    )
    gated = apply_quest_opens_gate(document, classic_sources)
    assert gated.quests[str(intermediate)][0].opens == "later"
    assert gated.quests[str(CHARM_OF_THE_SHIFTING_SANDS)][0].opens == "later"


def test_a_quest_with_no_classic_db_turn_in_coverage_is_left_alone():
    """A Forever-new quest, or one classic-db itself states no turn-in
    item for -- absent from the map entirely, distinct from an empty
    list -- is never gated, same "unverifiable, not invented" rule."""
    document = LootFile(
        sources=[
            LootSource(
                id="raid:onyxias-lair", kind="raid", name="Onyxia's Lair", opens="raids-1",
                bosses=[LootBoss(id="onyxia", name="Onyxia", npc_id=1, items=[HEAD_OF_ONYXIA])],
            ),
        ],
        quests={str(ONYXIA_TOOTH_PENDANT): [_quest_source(7491, "For All To See")]},
    )
    gated = apply_quest_opens_gate(document, classic_sources={})
    assert gated.quests[str(ONYXIA_TOOTH_PENDANT)][0].opens is None


def test_the_flat_quest_bucket_source_never_ungates_a_per_quest_entry():
    """Every quest-sourced item also sits in the flat, compatibility
    `LootSource(id="quest")` bucket, which carries no per-quest `opens`
    of its own -- `apply_quest_opens_gate` must skip that kind entirely
    (same as `sim/cmd/leveling-bis/data.go`'s `loadLootIndex` already
    does) rather than reading its bare `opens=None` as "an ungated way
    exists" and silently defeating the whole mechanism."""
    document = LootFile(
        sources=[
            LootSource(
                id="raid:onyxias-lair", kind="raid", name="Onyxia's Lair", opens="raids-1",
                bosses=[LootBoss(id="onyxia", name="Onyxia", npc_id=1, items=[HEAD_OF_ONYXIA])],
            ),
            LootSource(id="quest", kind="quest", name="Quests", items=[ONYXIA_TOOTH_PENDANT]),
        ],
        quests={str(ONYXIA_TOOTH_PENDANT): [_quest_source(7491, "For All To See")]},
    )
    classic_sources = _classic_sources({7491: [HEAD_OF_ONYXIA]})
    gated = apply_quest_opens_gate(document, classic_sources)
    assert gated.quests[str(ONYXIA_TOOTH_PENDANT)][0].opens == "raids-1"


def test_two_faction_mirrored_quests_for_the_same_item_are_gated_independently():
    """Same shape `QuestSource`'s own doc already relies on (Hammerbone,
    quest 914): one item, two per-faction QuestSource entries -- both
    must end up gated, not just whichever happens first."""
    document = LootFile(
        sources=[
            LootSource(
                id="raid:onyxias-lair", kind="raid", name="Onyxia's Lair", opens="raids-1",
                bosses=[LootBoss(id="onyxia", name="Onyxia", npc_id=1, items=[HEAD_OF_ONYXIA])],
            ),
        ],
        quests={
            str(ONYXIA_TOOTH_PENDANT): [
                _quest_source(7491, "For All To See"),
                _quest_source(7496, "Celebrating Good Times"),
            ],
        },
    )
    classic_sources = _classic_sources({7491: [HEAD_OF_ONYXIA], 7496: [HEAD_OF_ONYXIA]})
    gated = apply_quest_opens_gate(document, classic_sources)
    assert [entry.opens for entry in gated.quests[str(ONYXIA_TOOTH_PENDANT)]] == [
        "raids-1", "raids-1",
    ]
