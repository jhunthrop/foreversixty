"""spellranks.json: per class, spell name -> its player-castable rank chain
(docs/superpowers/specs/2026-09-27-level-aware-sim-design.md, design point 4).

tests/fixtures/spellconst/rogue.json carries the reference caveat's own example --
Sinister Strike, five rank-0 ids at level 20 (14873, 15581, 15667, 19472, 1213441)
beside its real ranks 1-8 (1752..11294) -- plus three shapes the filter has to get
right: Kick (one id, no ranks, but learned at level 12: kept, for the level gate),
Stealth (one id, no ranks, level 0: dropped, nothing to gate or rewrite) and
Improved Sap (two ids, both rank 0, both level 0: an unranked ability the client
duplicates -- kept, because there is more than one id, even though nothing here
is a "rank"). tests/fixtures/spellconst/warrior.json adds a second class with a
plain two-rank chain (Rend) and its own single-id, level-gated spell (Charge), so
class separation and the "family" field never leak across files.
"""

import json
from pathlib import Path

import pytest

from pipeline import spellranks
from pipeline.models import SpellConstant
from pipeline.spellranks import _spell_ranks_for_class, count_copy_names, is_copy_name

FIXTURES = Path(__file__).parent / "fixtures" / "spellconst"


def _spell(
    name: str, *, rank: int, level: int, cost: int = 0, cast_time_ms: int = 0, gcd_ms: int = 0
) -> SpellConstant:
    """A spell constant with only the fields the chain builder reads set; the rest zero."""
    return SpellConstant(
        name=name,
        rank=rank,
        spell_level=level,
        cost=cost,
        cost_type=0,
        cast_time_ms=cast_time_ms,
        gcd_ms=gcd_ms,
        cooldown_ms=0,
        category_cooldown_ms=0,
        duration_ms=0,
        school_mask=0,
        family_mask=[0, 0, 0, 0],
        effects=[],
    )


def _result():
    records = spellranks.load_class_spell_constants(FIXTURES)
    return spellranks.build_spell_ranks("9.9.9.9", records)


def test_shape_is_build_then_classes_then_name_to_rank_list():
    result = _result()
    assert result.build == "9.9.9.9"
    assert set(result.classes) == {"rogue", "warrior"}


def test_sinister_strikes_rank_zero_duplicates_never_join_the_chain():
    chain = _result().classes["rogue"]["Sinister Strike"]
    assert [rank.id for rank in chain] == [1752, 1757, 1758, 1759, 1760, 8621, 11293, 11294]
    assert [rank.rank for rank in chain] == list(range(1, 9))
    assert [rank.level for rank in chain] == [1, 6, 14, 22, 30, 38, 46, 54]
    dupes = {14873, 15581, 15667, 19472, 1213441}
    assert dupes.isdisjoint({rank.id for rank in chain})


def test_cost_less_gcd_less_copies_leave_a_ranked_chain():
    """The 1.60 client keeps a copy of Lightning Bolt beside every rank (408439
    beside 403...) and Magma Totem's pulse beside its totem (10579 beside
    10585): same name, same rank, no cost, no cast time, no GCD. The chain
    keeps the castable ids so a rewrite never names one the engine lacks."""
    spells = {
        "403": _spell("Lightning Bolt", rank=1, level=1, cost=15, cast_time_ms=1500, gcd_ms=1500),
        "408439": _spell("Lightning Bolt", rank=1, level=1),
        "529": _spell("Lightning Bolt", rank=2, level=8, cost=30, cast_time_ms=2000, gcd_ms=1500),
        "408440": _spell("Lightning Bolt", rank=2, level=8),
        "10585": _spell("Magma Totem", rank=2, level=36, cost=360, gcd_ms=1000),
        "10579": _spell("Magma Totem", rank=2, level=36),
    }
    ranks = _spell_ranks_for_class(spells)
    assert [r.id for r in ranks["Lightning Bolt"]] == [403, 529]
    assert [r.id for r in ranks["Magma Totem"]] == [10585]


def test_an_uncastable_ranked_copy_yields_to_the_castable_unranked_spell():
    """Stormstrike: the 1.60 client carries 410156 at rank 1, level 1 with no
    cost, cast time or GCD (a Season of Discovery copy) beside the player's
    own 17364 at rank 0, level 40 with a GCD. Grounding Totem's totem 8178
    sits at rank 1 beside the castable 8177 at rank 0. The chain is the
    castable unranked id, so the APL's 17364 resolves and gates at 40."""
    spells = {
        "17364": _spell("Stormstrike", rank=0, level=40, gcd_ms=1500),
        "410156": _spell("Stormstrike", rank=1, level=1),
        "8177": _spell("Grounding Totem", rank=0, level=30, gcd_ms=1000),
        "8178": _spell("Grounding Totem", rank=1, level=30),
    }
    ranks = _spell_ranks_for_class(spells)
    assert [(r.id, r.rank, r.level) for r in ranks["Stormstrike"]] == [(17364, 0, 40)]
    assert [(r.id, r.rank, r.level) for r in ranks["Grounding Totem"]] == [(8177, 0, 30)]


def test_an_unranked_copy_learned_before_the_castable_spell_is_dropped():
    """Water Shield: the player's 408510 (level 20, a GCD) sits beside its
    proc 408511 (level 20) and a level-1 copy 409941, neither of which costs,
    casts or triggers the GCD. Left in, the copy made the spell learned from
    level 1 and the addon showed a level-10 Restoration shaman a Water Shield
    line. The proc stays: it is learned with the spell, and the API's
    adherence reads a combat log's proc id as Water Shield through it."""
    spells = {
        "408510": _spell("Water Shield", rank=0, level=20, gcd_ms=1500),
        "408511": _spell("Water Shield", rank=0, level=20),
        "409941": _spell("Water Shield", rank=0, level=1),
    }
    ranks = _spell_ranks_for_class(spells)
    assert [(r.id, r.level) for r in ranks["Water Shield"]] == [(408510, 20), (408511, 20)]


def test_an_unranked_effect_id_learned_with_the_spell_is_kept():
    """Whirlwind's 462891 is learned at 36 with the castable 1680: it says
    nothing earlier about when the spell is learned, and the API maps a log's
    462891 hit to Whirlwind through it."""
    spells = {
        "1680": _spell("Whirlwind", rank=0, level=36, cost=250, gcd_ms=1500),
        "462891": _spell("Whirlwind", rank=0, level=36),
    }
    assert [r.id for r in _spell_ranks_for_class(spells)["Whirlwind"]] == [1680, 462891]


def test_a_ranked_chain_with_no_castable_id_keeps_every_id():
    spells = {
        "1": _spell("Odd Aura", rank=1, level=10),
        "2": _spell("Odd Aura", rank=2, level=20),
    }
    assert [r.id for r in _spell_ranks_for_class(spells)["Odd Aura"]] == [1, 2]


def test_a_single_id_unranked_spell_with_a_learn_level_is_kept():
    """Kick has no ranks but is learned at 12; sim/request still needs that
    level to gate the action even with nothing to rewrite it to."""
    chain = _result().classes["rogue"]["Kick"]
    assert [(rank.id, rank.rank, rank.level) for rank in chain] == [(1766, 0, 12)]


def test_a_single_id_unranked_spell_at_level_zero_is_dropped():
    """Stealth: one id, rank 0, level 0 -- neither a rewrite nor a level gate
    needs it, so it never reaches the file."""
    assert "Stealth" not in _result().classes["rogue"]


def test_an_unranked_duplicate_pair_is_kept_sorted_by_id_not_collapsed():
    """Improved Sap: two ids, both rank 0, both level 0 -- there is no ranked
    chain (rank >= 1) here at all, so both entries are kept as given rather
    than one being picked as canonical or the pair being dropped outright."""
    chain = _result().classes["rogue"]["Improved Sap"]
    assert [(rank.id, rank.rank, rank.level) for rank in chain] == [(2070, 0, 0), (2071, 0, 0)]


def test_is_copy_name_matches_only_the_prefix():
    assert is_copy_name("Copy of Mortal Strike") is True
    assert is_copy_name("Copy of Frostbolt") is True
    assert is_copy_name("Mortal Strike") is False
    # A name that merely contains the words, not at the start, is a real
    # (if oddly named) player ability and must survive.
    assert is_copy_name("Blueprint: Copy of a Key") is False


def test_a_copy_of_name_never_forms_its_own_chain():
    """Copy of Deadly Poison IV (25348) is the ladder's own real example
    (see sim/request/testdata/ladder/rogue-*.golden.md's "learned but
    unused" sections before this filter existed). It has a nonzero learn
    level -- the one inclusion rule a lone, unranked id can pass on its
    own -- so this proves the name filter runs before that rule, not that
    the level happens to be zero."""
    spells = {
        "25348": _spell("Copy of Deadly Poison IV", rank=0, level=30),
    }
    assert _spell_ranks_for_class(spells) == {}


def test_the_rogue_fixtures_copy_of_entry_is_absent_from_the_result():
    assert "Copy of Deadly Poison IV" not in _result().classes["rogue"]


def test_count_copy_names_counts_across_every_class():
    assert count_copy_names(spellranks.load_class_spell_constants(FIXTURES)) == 1


def test_count_copy_names_is_zero_when_none_are_present():
    records = [
        record
        for record in spellranks.load_class_spell_constants(FIXTURES)
        if record.class_slug == "warrior"
    ]
    assert count_copy_names(records) == 0


def test_classes_do_not_leak_into_each_other():
    result = _result()
    assert set(result.classes["warrior"]) == {"Rend", "Charge"}
    assert [rank.id for rank in result.classes["warrior"]["Rend"]] == [772, 6546]
    assert "Sinister Strike" not in result.classes["warrior"]


def test_output_is_deterministic_and_sort_keys_when_written(tmp_path: Path):
    first = _result().model_dump()
    second = _result().model_dump()
    assert first == second

    build_dir = tmp_path / "builds" / "9.9.9.9"
    (build_dir / spellranks.SPELLCONST).mkdir(parents=True)
    for path in FIXTURES.glob("*.json"):
        (build_dir / spellranks.SPELLCONST / path.name).write_text(path.read_text())
    manifest = {"build": "9.9.9.9", "product": "test", "fetched_at": "t", "files": {}}
    (build_dir / "manifest.json").write_text(json.dumps(manifest))

    out = spellranks.write_spell_ranks("9.9.9.9", root=tmp_path / "builds")
    text = out.read_text(encoding="utf-8")
    written = json.loads(text)
    assert list(written.keys()) == sorted(written.keys())
    assert list(written["classes"].keys()) == sorted(written["classes"].keys())
    assert list(written["classes"]["rogue"].keys()) == sorted(written["classes"]["rogue"].keys())


def test_write_spell_ranks_raises_without_spellconst(tmp_path: Path):
    build_dir = tmp_path / "builds" / "1.0.0.1"
    build_dir.mkdir(parents=True)
    with pytest.raises(SystemExit, match="spellconst"):
        spellranks.write_spell_ranks("1.0.0.1", root=tmp_path / "builds")


CURATED = Path(__file__).parent.parent / "curated"
BUILDS = Path(__file__).parent.parent / "builds"

#: Book ranks the fork's core.IncludeAQ gates today (sim/core/config.go) -- a floor,
#: never a census: the curated file may name more, never fewer.
ENGINE_GATED_BOOK_IDS = {
    25307, 25311, 25309, 25306, 25304, 31016, 25300, 25347, 25295, 25296, 25286, 25288,
    31018, 25298, 25289, 25291, 25361, 25359, 25290, 25345,
}  # fmt: skip


def test_the_curated_book_list_names_every_engine_gated_rank():
    assert ENGINE_GATED_BOOK_IDS <= spellranks.load_book_ids(CURATED)


def test_book_ids_are_marked_and_every_other_rank_is_not():
    spells = {
        "1": _spell("Odd Bolt", rank=9, level=58, cost=1),
        "2": _spell("Odd Bolt", rank=10, level=60, cost=1),
    }
    ranks = _spell_ranks_for_class(spells, frozenset({2}))["Odd Bolt"]
    assert [(r.id, r.book) for r in ranks] == [(1, False), (2, True)]


def test_the_book_flag_is_written_only_when_true():
    spells = {
        "1": _spell("Odd Bolt", rank=9, level=58, cost=1),
        "2": _spell("Odd Bolt", rank=10, level=60, cost=1),
    }
    file = spellranks.SpellRanksFile(
        build="9.9.9.9", classes={"mage": _spell_ranks_for_class(spells, frozenset({2}))}
    )
    rows = file.model_dump(exclude_defaults=True)["classes"]["mage"]["Odd Bolt"]
    assert rows == [
        {"id": 1, "rank": 9, "level": 58},
        {"id": 2, "rank": 10, "level": 60, "book": True},
    ]


def test_named_book_ids_are_flagged_level_sixty_ranks_in_the_active_build():
    """Floor contract: these ids stay flagged in the active build's own table. A named id
    the table does not carry is not a failure (the build may genuinely lack it)."""
    named = {1: 25307, 2: 25304, 3: 25295, 4: 25289}
    active = json.loads((BUILDS.parent.parent / "web/src/data/active-build.json").read_text())
    for path in [BUILDS / active["build"] / "spellranks.json"]:
        rows = {
            row["id"]: row
            for classes in json.loads(path.read_text())["classes"].values()
            for ranks in classes.values()
            for row in ranks
        }
        for book_id in named.values():
            if book_id in rows:
                assert rows[book_id].get("book") is True, (
                    f"{path.parent.name}: {book_id} not flagged"
                )
                assert rows[book_id]["level"] == 60


# --- inferior ranks -------------------------------------------------------------------------

INFERIOR_IDS: set[int] = set()
# 1.60.1.70707 made Aspect of the Hawk rank 6 (14322, 55 ranged attack power against rank 5's
# 47) and Trueshot Aura rank 5 (20906) the strongest ranks again, so neither is inferior.
# Hunter's Mark rank 4 (14325) is 143 against rank 3's 98 since the 23 October hotfix and in
# the 70707 SpellEffect, and it is the id SkillLineAbility teaches (sim-accuracy round 5).
NO_LONGER_INFERIOR = {14322, 20906, 14325}


def test_the_hawk_and_trueshot_top_ranks_are_not_inferior():
    assert not NO_LONGER_INFERIOR & spellranks.load_inferior_ids(CURATED)


def test_the_curated_inferior_list_names_every_swept_rank():
    assert INFERIOR_IDS <= spellranks.load_inferior_ids(CURATED)


def test_inferior_ids_are_marked_and_every_other_rank_is_not():
    spells = {
        "1": _spell("Odd Aspect", rank=5, level=48, cost=1),
        "2": _spell("Odd Aspect", rank=6, level=58, cost=1),
    }
    ranks = _spell_ranks_for_class(spells, inferior_ids=frozenset({2}))["Odd Aspect"]
    assert [(r.id, r.inferior) for r in ranks] == [(1, False), (2, True)]


def test_the_inferior_flag_is_written_only_when_true():
    spells = {
        "1": _spell("Odd Aspect", rank=5, level=48, cost=1),
        "2": _spell("Odd Aspect", rank=6, level=58, cost=1),
    }
    file = spellranks.SpellRanksFile(
        build="9.9.9.9",
        classes={"hunter": _spell_ranks_for_class(spells, inferior_ids=frozenset({2}))},
    )
    rows = file.model_dump(exclude_defaults=True)["classes"]["hunter"]["Odd Aspect"]
    assert rows == [
        {"id": 1, "rank": 5, "level": 48},
        {"id": 2, "rank": 6, "level": 58, "inferior": True},
    ]


def test_the_active_builds_hawk_rank_6_is_not_inferior():
    """1.60.1.70707 made Aspect of the Hawk rank 6 (55) stronger than rank 5 (47) again."""
    active = json.loads((BUILDS.parent.parent / "web/src/data/active-build.json").read_text())
    path = BUILDS / active["build"] / "spellranks.json"
    hawk = json.loads(path.read_text())["classes"]["hunter"]["Aspect of the Hawk"]
    rows = {row["id"]: row for row in hawk}
    assert not rows[14322].get("inferior")
    assert not rows[14321].get("inferior")


# Sim round 22 platform (shaman X2): since client 70753 the spell constants carry
# Stormstrike 17364 at its client spell level 40, and the cards, the ladder and the
# addon read that as its learn level. Stormstrike is a tier-3 talent: a shaman takes
# it at 25 (tier t opens at 10 + 5t) and the engine casts it from there. A spell a
# talent teaches is learned when the talent can first be taken, and a trainer rank
# above it is learned no earlier than that.
def test_a_talent_taught_spell_is_learned_at_its_tiers_level():
    spells = {
        "17364": _spell("Stormstrike", rank=0, level=40, gcd_ms=1500),
        "5570": _spell("Insect Swarm", rank=1, level=20, cost=45, gcd_ms=1500),
        "24974": _spell("Insect Swarm", rank=2, level=22, cost=60, gcd_ms=1500),
        "24975": _spell("Insect Swarm", rank=3, level=30, cost=85, gcd_ms=1500),
        "403": _spell("Lightning Bolt", rank=1, level=1, cost=15, cast_time_ms=1500),
    }
    talent_levels = {17364: 25, 5570: 25}
    ranks = _spell_ranks_for_class(spells, talent_levels=talent_levels)
    assert [(r.id, r.level) for r in ranks["Stormstrike"]] == [(17364, 25)]
    assert [(r.id, r.level) for r in ranks["Insect Swarm"]] == [
        (5570, 25),
        (24974, 25),
        (24975, 30),
    ]
    assert [(r.id, r.level) for r in ranks["Lightning Bolt"]] == [(403, 1)]


def test_a_talent_level_never_adds_a_name_the_client_levels_leave_out():
    """Shadowform 15473 is a single level-0 row: nothing to rewrite or gate, so the file
    leaves it out, and its talent's level does not bring it in."""
    spells = {"15473": _spell("Shadowform", rank=0, level=0, gcd_ms=1500)}
    assert "Shadowform" not in _spell_ranks_for_class(spells, talent_levels={15473: 40})


def test_talent_levels_open_each_tier_five_levels_apart(tmp_path: Path):
    talents = {
        "trees": [
            {
                "name": "Enhancement",
                "talents": [
                    {
                        "name": "Stormstrike",
                        "tier": 3,
                        "spell_id": 17364,
                        "ranks": [{"spell_id": 17364}],
                    },
                    {
                        "name": "Ancestral Knowledge",
                        "tier": 0,
                        "spell_id": 17485,
                        "ranks": [{"spell_id": 17485}, {"spell_id": 17486}],
                    },
                ],
            }
        ]
    }
    (tmp_path / "shaman.json").write_text(json.dumps(talents), encoding="utf-8")
    from pipeline.talent_levels import load_talent_spell_levels

    assert load_talent_spell_levels(tmp_path, "shaman") == {17364: 25, 17485: 10, 17486: 10}
    assert load_talent_spell_levels(tmp_path, "mage") == {}
