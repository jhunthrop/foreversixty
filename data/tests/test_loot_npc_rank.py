# data/tests/test_loot_npc_rank.py
"""Task L1.7: every loot.json boss row carries its creature's classic-db
`creature_template.Rank` (raw/classicdb/sources.json's `npc_ranks`), so the
BiS ranker can tell a rare spawn (Ironspine) from a boss or a trash mob."""

import json
from pathlib import Path

from pipeline.classic_sources import ClassicDbSourceRecord
from pipeline.loot.npc_rank import apply_npc_ranks, rare_spawn_names
from pipeline.models import LootBoss, LootFile, LootSource
from pipeline.normalize import write_document

BUILD = Path(__file__).resolve().parent.parent / "builds" / "1.60.1.70291"


def _document() -> LootFile:
    return LootFile(
        sources=[
            LootSource(
                id="dungeon:scarlet-monastery",
                kind="dungeon",
                name="Scarlet Monastery",
                bosses=[
                    LootBoss(id="d:sm:6489", name="Ironspine", npc_id=6489, items=[7686]),
                    LootBoss(id="d:sm:3975", name="Herod", npc_id=3975, items=[7718]),
                    LootBoss(id="d:sm:0", name="", npc_id=0, items=[1]),
                ],
            ),
            LootSource(
                id="crafted:leatherworking", kind="crafted", name="Leatherworking", items=[7358]
            ),
        ]
    )


def test_a_boss_row_takes_its_npcs_rank_and_a_normal_or_unknown_npc_has_none():
    document = apply_npc_ranks(_document(), {6489: 2, 3975: 3}, {6489})
    bosses = document.sources[0].bosses
    assert [boss.rank for boss in bosses] == [2, 3, None]
    assert [boss.rare_spawn for boss in bosses] == [True, None, None]
    assert document.sources[1].bosses is None


def test_the_rank_is_written_only_where_the_dump_states_one(tmp_path):
    path = tmp_path / "loot.json"
    write_document(apply_npc_ranks(_document(), {6489: 2, 3975: 3}, {6489}), path)
    bosses = json.loads(path.read_text(encoding="utf-8"))["sources"][0]["bosses"]
    assert (bosses[0]["rank"], bosses[0]["rare_spawn"]) == (2, True)
    assert bosses[1]["rank"] == 3 and "rare_spawn" not in bosses[1]
    assert "rank" not in bosses[2] and "rare_spawn" not in bosses[2]


def test_the_real_loot_marks_ironspine_a_rare_spawn_and_vancleef_and_greenskin_not():
    loot = json.loads((BUILD / "loot.json").read_text(encoding="utf-8"))
    bosses = {
        boss["npc_id"]: (boss.get("rank"), boss.get("rare_spawn", False))
        for source in loot["sources"]
        for boss in source.get("bosses") or []
    }
    assert bosses[6489] == (2, True)  # Ironspine, a Scarlet Monastery graveyard rare
    assert bosses[639] == (2, False)  # Edwin VanCleef: ranked rare elite, always there
    assert bosses[647] == (1, False)  # Captain Greenskin, an elite


def _world_document() -> LootFile:
    return LootFile(
        sources=[
            LootSource(id="world:swiftmane", kind="world", name="Swiftmane", items=[285330]),
            LootSource(
                id="world:bluegill", kind="world", name="Bluegill Muckdweller", items=[3022]
            ),
            LootSource(id="world:twin", kind="world", name="Twin", items=[1]),
            LootSource(id="world_drop:21-28", kind="world_drop", name="World drop", items=[3021]),
        ]
    )


def _creature(npc_id: int, name: str) -> ClassicDbSourceRecord:
    return ClassicDbSourceRecord(kind="creature_drop", npc_id=npc_id, name=name)


def test_a_world_source_is_a_rare_spawn_when_every_creature_of_its_name_is_one():
    classic = {
        285330: [_creature(5831, "Swiftmane")],
        3022: [_creature(1028, "Bluegill Muckdweller")],
        # Two creatures share the name "Twin"; only one is a rare spawn.
        1: [_creature(7, "Twin"), _creature(8, "Twin")],
    }
    names = rare_spawn_names(classic, {5831, 7})
    assert names == {"Swiftmane"}
    document = apply_npc_ranks(_world_document(), {}, {5831, 7}, names)
    assert [source.rare_spawn for source in document.sources] == [True, None, None, None]


def test_the_real_loot_marks_swiftmane_a_rare_spawn_and_bluegill_not():
    loot = json.loads((BUILD / "loot.json").read_text(encoding="utf-8"))
    world = {s["id"]: s.get("rare_spawn", False) for s in loot["sources"] if s["kind"] == "world"}
    assert world["world:swiftmane"] is True
    assert world["world:lupos"] is True
    assert world["world:bluegill-muckdweller"] is False
