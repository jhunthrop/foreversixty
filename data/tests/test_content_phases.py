"""`pipeline.loot.content_phases`: the curated content-phase gate.

`curated/content-phases.json` keys every source that opened after the
original Classic Phase 1 by a stable game id. These tests pin how each key
reaches loot.json, that a duplicate quest entry can no longer leak around a
gate, and that an id missing from a night's fetch keeps its gate.
"""

from __future__ import annotations

import json
from pathlib import Path

import pytest

from pipeline.loot.content_phases import (
    ContentPhaseError,
    apply_content_phases,
    load_content_phases,
    merge_duplicate_quests,
)
from pipeline.models import LootBoss, LootFile, LootSource, QuestSource

CURATED = Path("curated")
PHASES = [
    {"name": "pre-beta", "start": "0001-01-01T00:00:00Z"},
    {"name": "launch", "start": "2026-11-04T23:00:00Z"},
    {"name": "raids-1", "start": "2026-12-09T00:00:00Z"},
]
PROVENANCE = [{"label": "l", "url": "https://x.invalid", "kind": "site"}]


def _write(tmp_path: Path, gates: list[dict]) -> Path:
    (tmp_path / "phases.json").write_text(json.dumps(PHASES), encoding="utf-8")
    document = {"sources": PROVENANCE, "notes": "why", "gates": gates}
    (tmp_path / "content-phases.json").write_text(json.dumps(document), encoding="utf-8")
    return tmp_path


def _gate(key: str = "g", phase: str = "later", date: str = "unannounced", **keys) -> dict:
    return {
        "key": key,
        "label": key,
        "original_patch": "1.9",
        "phase": phase,
        "date": date,
        "confidence": "inferred",
        "why": "because",
        **keys,
    }


def _quest(quest_id: int, name: str = "q", **extra) -> QuestSource:
    return QuestSource(
        quest_id=quest_id, name=name, faction="both", min_level=60, level=60,
        level_source="classic-db", **extra,
    )  # fmt: skip


def _apply(tmp_path: Path, gates: list[dict], document: LootFile) -> tuple[LootFile, list[str]]:
    return apply_content_phases(document, load_content_phases(_write(tmp_path, gates)))


def _by_id(document: LootFile) -> dict[str, LootSource]:
    return {source.id: source for source in document.sources}


# -- the curated file itself ------------------------------------------------


def test_the_curated_file_loads_and_every_gate_names_a_known_phase():
    phases = load_content_phases(CURATED)
    assert phases.gates, "content-phases.json names no gate"
    for gate in phases.gates:
        assert gate.phase in {"pre-beta", "beta", "launch", "raids-1", "later"}, gate.key


def test_the_go_hand_lists_now_live_in_the_curated_file():
    """data.go's raidLockedQuestOpens, repFactionRaidPhaseOpens and
    worldBossSources, and overlay.py's WORLD_BOSS_NPC_NAMES, moved here."""
    index = load_content_phases(CURATED).index()
    for quest_id in (8756, 8789, 8790):
        assert quest_id in index.quests
    for faction_id in (609, 910, 270):
        assert faction_id in index.factions
    for npc_id in (12397, 6109, 14889, 14888, 14890, 14887):
        assert npc_id in index.npcs
    assert index.npcs[15192].opens == "later"  # Anachronos, gate deleted in 3592e40a5


@pytest.mark.parametrize(
    ("gate", "message"),
    [
        (_gate(phase="phase-9"), "phase-9"),
        (_gate(phase="later", date="2027-01-01"), "unannounced"),
        (_gate(phase="raids-1", date="2026-12-10"), "2026-12-09"),
        (_gate(), "names no id"),
    ],
)
def test_a_malformed_gate_fails_fast(tmp_path, gate, message):
    with pytest.raises(ContentPhaseError, match=message):
        load_content_phases(_write(tmp_path, [gate]))


def test_one_id_in_two_gates_fails_fast(tmp_path):
    gates = [
        _gate("a", npcs=[{"id": 1, "name": "One"}]),
        _gate("b", npcs=[{"id": 1, "name": "One"}]),
    ]
    with pytest.raises(ContentPhaseError, match="npc 1"):
        load_content_phases(_write(tmp_path, gates))


# -- each key reaches loot.json -----------------------------------------------


def test_a_vendor_and_a_world_npc_are_gated_by_npc(tmp_path):
    document = LootFile(
        sources=[
            LootSource(
                id="vendor:15192", kind="vendor", name="Anachronos", npc_id=15192, items=[1]
            ),
            LootSource(
                id="world:prince-skaldrenox", kind="world", name="Prince Skaldrenox", items=[2]
            ),
            LootSource(id="world:aged-kodo", kind="world", name="Aged Kodo", items=[3]),
        ]
    )
    gates = [
        _gate(
            npcs=[{"id": 15192, "name": "Anachronos"}, {"id": 15203, "name": "Prince Skaldrenox"}]
        )
    ]
    out, missing = _apply(tmp_path, gates, document)
    by_id = _by_id(out)
    assert by_id["vendor:15192"].opens == "later"
    assert by_id["world:prince-skaldrenox"].opens == "later"
    assert by_id["world:aged-kodo"].opens is None
    assert missing == []


def test_a_gated_boss_inside_an_open_dungeon_gets_its_own_source(tmp_path):
    """Balzaphon no longer inherits Stratholme's open status."""
    document = LootFile(
        sources=[
            LootSource(
                id="dungeon:stratholme",
                kind="dungeon",
                name="Stratholme",
                zone_id=2017,
                bosses=[
                    LootBoss(
                        id="dungeon:stratholme:14684", name="Balzaphon", npc_id=14684, items=[23125]
                    ),
                    LootBoss(
                        id="dungeon:stratholme:10440",
                        name="Baron Rivendare",
                        npc_id=10440,
                        items=[13335],
                    ),
                ],
            )  # fmt: skip
        ]
    )
    gates = [_gate("scourge-invasion", npcs=[{"id": 14684, "name": "Balzaphon"}])]
    out, _ = _apply(tmp_path, gates, document)
    by_id = _by_id(out)
    assert by_id["dungeon:stratholme"].opens is None
    assert [boss.name for boss in by_id["dungeon:stratholme"].bosses] == ["Baron Rivendare"]
    invasion = by_id["dungeon:stratholme:scourge-invasion"]
    assert (invasion.kind, invasion.name, invasion.zone_id, invasion.opens) == (
        "dungeon", "Stratholme", 2017, "later",
    )  # fmt: skip
    assert [boss.items for boss in invasion.bosses] == [[23125]]


def test_zone_kind_and_faction_gates(tmp_path):
    document = LootFile(
        sources=[
            LootSource(
                id="dungeon:dire-maul", kind="dungeon", name="Dire Maul", zone_id=2557, bosses=[]
            ),
            LootSource(id="zone:1377", kind="zone", name="Silithus", zone_id=1377, items=[20682]),
            LootSource(id="pvp:rank-12:horde", kind="pvp", name="Rank 12", items=[19595]),
            LootSource(
                id="rep:silverwing-sentinels:exalted",
                kind="rep",
                name="Silverwing Sentinels",
                faction_id=889,
                standing="exalted",
                items=[19595],
            ),
            LootSource(
                id="rep:timbermaw-hold:honored",
                kind="rep",
                name="Timbermaw Hold",
                faction_id=576,
                standing="honored",
                items=[9],
            ),
        ]  # fmt: skip
    )
    gates = [
        _gate("dm", zones=[{"id": 2557, "name": "Dire Maul"}, {"id": 1377, "name": "Silithus"}]),
        _gate("honor", kinds=["pvp"]),
        _gate("bg", factions=[{"id": 889, "name": "Silverwing Sentinels"}]),
    ]
    out, _ = _apply(tmp_path, gates, document)
    by_id = _by_id(out)
    assert by_id["dungeon:dire-maul"].opens == "later"
    assert by_id["zone:1377"].opens == "later"
    assert by_id["pvp:rank-12:horde"].opens == "later"
    assert by_id["rep:silverwing-sentinels:exalted"].opens == "later"
    assert by_id["rep:timbermaw-hold:honored"].opens is None


def test_a_faction_gate_from_a_standing_up(tmp_path):
    document = LootFile(
        sources=[
            LootSource(
                id="rep:x:friendly",
                kind="rep",
                name="X",
                faction_id=5,
                standing="friendly",
                items=[1],
            ),
            LootSource(
                id="rep:x:exalted",
                kind="rep",
                name="X",
                faction_id=5,
                standing="exalted",
                items=[2],
            ),
        ],
        quests={"3": [_quest(10, required_rep_faction=5, required_rep_standing="revered")]},
    )
    gates = [_gate(factions=[{"id": 5, "name": "X", "from_standing": "revered"}])]
    out, _ = _apply(tmp_path, gates, document)
    by_id = _by_id(out)
    assert by_id["rep:x:friendly"].opens is None
    assert by_id["rep:x:exalted"].opens == "later"
    assert out.quests["3"][0].opens == "later"


def test_quests_are_gated_by_id_and_by_the_faction_they_need(tmp_path):
    document = LootFile(
        sources=[],
        quests={
            "22678": [_quest(9222, "Epic Armaments of Battle - Friend of the Dawn")],
            "19601": [
                _quest(
                    8104,
                    "The Jewel of Kajaro",
                    required_rep_faction=270,
                    required_rep_standing="revered",
                )
            ],
            "1": [_quest(100, "An Ordinary Quest")],
        },  # fmt: skip
    )
    gates = [
        _gate("ad", quests=[{"id": 9222, "name": "Epic Armaments"}]),
        _gate("zg", factions=[{"id": 270, "name": "Zandalar Tribe"}]),
    ]
    out, _ = _apply(tmp_path, gates, document)
    assert out.quests["22678"][0].opens == "later"
    assert out.quests["19601"][0].opens == "later"
    assert out.quests["1"][0].opens is None


def test_an_item_gate_splits_the_item_out_of_every_flat_source(tmp_path):
    document = LootFile(
        sources=[
            LootSource(
                id="crafted:leatherworking", kind="crafted", name="Leatherworking",
                profession="leatherworking", items=[21278, 8],
            )
        ],
        quests={"21278": [_quest(1)]},
    )  # fmt: skip
    out, _ = _apply(
        tmp_path,
        [_gate("late-recipes", items=[{"id": 21278, "name": "Stormshroud Gloves"}])],
        document,
    )
    by_id = _by_id(out)
    assert by_id["crafted:leatherworking"].items == [8]
    sibling = by_id["crafted:leatherworking:late-recipes"]
    assert (sibling.items, sibling.opens, sibling.profession) == (
        [21278],
        "later",
        "leatherworking",
    )
    assert out.quests["21278"][0].opens == "later"


def test_a_gate_dated_at_launch_opens_its_source(tmp_path):
    """A confirmed date flips a gate with no code change."""
    document = LootFile(
        sources=[
            LootSource(id="vendor:15192", kind="vendor", name="Anachronos", npc_id=15192, items=[1])
        ]
    )
    gates = [_gate(phase="launch", date="2026-11-04", npcs=[{"id": 15192, "name": "Anachronos"}])]
    out, _ = _apply(tmp_path, gates, document)
    assert _by_id(out)["vendor:15192"].opens is None


def test_an_existing_stricter_gate_is_kept(tmp_path):
    document = LootFile(
        sources=[
            LootSource(
                id="world:lord-kazzak", kind="world", name="Lord Kazzak", items=[1], opens="later"
            )
        ]
    )
    gates = [_gate(phase="raids-1", date="2026-12-09", npcs=[{"id": 12397, "name": "Lord Kazzak"}])]
    out, _ = _apply(tmp_path, gates, document)
    assert _by_id(out)["world:lord-kazzak"].opens == "later"


# -- duplicate quest entries ----------------------------------------------------


def test_duplicate_quest_entries_merge_so_the_gated_copy_wins():
    """Treasure of the Timeless One: one copy states Brood of Nozdormu and
    `opens`, the other is bare. One ungated way is enough to reach an item,
    so the bare copy used to defeat the gate."""
    document = LootFile(
        sources=[],
        quests={
            "21520": [
                _quest(8745, "Treasure of the Timeless One"),
                _quest(
                    8745, "Treasure of the Timeless One", opens="later",
                    required_rep_faction=910, required_rep_standing="neutral",
                ),
                _quest(9000, "Another Quest"),
            ]
        },
    )  # fmt: skip
    merged = merge_duplicate_quests(document).quests["21520"]
    assert [entry.quest_id for entry in merged] == [8745, 9000]
    assert (merged[0].opens, merged[0].required_rep_faction, merged[0].required_rep_standing) == (
        "later", 910, "neutral",
    )  # fmt: skip


def test_a_bare_duplicate_takes_the_faction_gate_of_its_twin(tmp_path):
    """The Eye of Zuldazar: only one copy names Zandalar Tribe."""
    document = LootFile(
        sources=[],
        quests={
            "19593": [
                _quest(8051),
                _quest(8051, required_rep_faction=270, required_rep_standing="revered"),
            ]
        },
    )  # fmt: skip
    out, _ = _apply(tmp_path, [_gate(factions=[{"id": 270, "name": "Zandalar Tribe"}])], document)
    assert [(entry.quest_id, entry.opens) for entry in out.quests["19593"]] == [(8051, "later")]


def test_a_quest_reward_inherits_the_rep_gate_of_its_item():
    """Earthstrike: quest 8573 states no reputation, but the same item is a
    Cenarion Circle exalted reward, so the quest inherits that faction."""
    document = LootFile(
        sources=[
            LootSource(
                id="rep:cenarion-circle:exalted", kind="rep", name="Cenarion Circle",
                faction_id=609, standing="exalted", items=[21180],
            )
        ],
        quests={"21180": [_quest(8573, "Champion's Battlegear")]},
    )  # fmt: skip
    out, _ = apply_content_phases(document, load_content_phases(CURATED))
    entry = out.quests["21180"][0]
    assert (entry.required_rep_faction, entry.required_rep_standing, entry.opens) == (
        609, "exalted", "later",
    )  # fmt: skip


# -- an id missing from tonight's fetch -----------------------------------------


def test_a_source_missing_from_the_fetch_keeps_its_gate(tmp_path, caplog):
    """Commit 3592e40a5 deleted the Anachronos gate when the vendor vanished
    from one fetch, and the vendor came back ungated. A missing id is now a
    warning: the gate stays in the curated file and applies the night the
    source returns."""
    gates = [
        _gate("aq", npcs=[{"id": 15192, "name": "Anachronos"}], quests=[{"id": 8745, "name": "T"}])
    ]
    curated = _write(tmp_path, gates)
    before = (curated / "content-phases.json").read_text(encoding="utf-8")
    phases = load_content_phases(curated)

    empty_night = LootFile(
        sources=[LootSource(id="world:aged-kodo", kind="world", name="Aged Kodo", items=[3])]
    )
    out, missing = apply_content_phases(empty_night, phases)
    assert out == empty_night
    assert missing == ["aq: npc 15192 Anachronos", "aq: quest 8745 T"]
    assert "npc 15192 Anachronos" in caplog.text
    assert (curated / "content-phases.json").read_text(encoding="utf-8") == before

    returned = LootFile(
        sources=[
            LootSource(
                id="vendor:15192", kind="vendor", name="Anachronos", npc_id=15192, items=[234032]
            )
        ]
    )
    out, missing = apply_content_phases(returned, phases)
    assert _by_id(out)["vendor:15192"].opens == "later"
    assert missing == ["aq: quest 8745 T"]


# -- launch: the banned list against the active build -----------------------

BANNED = Path("../sim/cmd/leveling-bis/testdata/launch-banned.json")
ACTIVE_BUILD = Path("../web/src/data/active-build.json")
#: The ranker never names a `zone` source as an item's source
#: (sim/cmd/leveling-bis/band.go's sourceKindPriority), so a zone listing
#: alone does not put an item on a list.
UNPICKABLE_KINDS = frozenset({"quest", "zone"})


def _open_ways(document: LootFile, item_ids: set[int]) -> set[int]:
    """The ids in `item_ids` an ungated, rankable source or quest gives."""
    found: set[int] = set()
    for source in document.sources:
        if source.kind in UNPICKABLE_KINDS or source.opens:
            continue
        items = {*(source.items or []), *(source.trash or [])}
        items |= {item for boss in source.bosses or [] for item in boss.items}
        found |= items & item_ids
    for item_id, entries in document.quests.items():
        if int(item_id) in item_ids and any(not entry.opens for entry in entries):
            found.add(int(item_id))
    return found


def _banned_while_closed(phases) -> dict[int, str]:
    """Banned item id -> its gate, for every gate still closed at launch: a
    gate the owner opens unbans its items with no edit here."""
    rows = json.loads(BANNED.read_text(encoding="utf-8"))["items"]
    gates = {gate.key: gate for gate in phases.gates}
    assert {row["gate"] for row in rows} <= set(gates), "launch-banned.json names an unknown gate"
    closed = {key for key, gate in gates.items() if phases.closed(gate.phase)}
    assert closed <= {row["gate"] for row in rows}, "every closed gate needs a banned example"
    return {row["id"]: row["gate"] for row in rows if row["gate"] in closed}


def test_no_banned_item_has_an_open_way_in_the_active_build():
    phases = load_content_phases(CURATED)
    banned = _banned_while_closed(phases)
    build = json.loads(ACTIVE_BUILD.read_text(encoding="utf-8"))["build"]
    document = LootFile(
        **json.loads((Path("builds") / build / "loot.json").read_text(encoding="utf-8"))
    )
    gated, missing = apply_content_phases(document, phases)
    leaks = _open_ways(gated, set(banned))
    assert not leaks, {item_id: banned[item_id] for item_id in sorted(leaks)}
    assert missing == [], missing
