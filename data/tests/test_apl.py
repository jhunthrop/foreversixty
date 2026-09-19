import json
from pathlib import Path

import pytest

from pipeline.apl import (
    APL_STATES,
    EMPTY_ROTATION,
    AplError,
    action_ids,
    cast_spell_action_ids,
    load_all,
    load_apl,
    parse_rotation,
)
from pipeline.curated import SOURCE_KINDS
from pipeline.manifest import newest_build
from pipeline.specs import load_specs

APL_DIR = Path("curated/apl")
#: Resolved from the manifest, not named: a build id here would have to be
#: edited every time a client build lands, and the ranks are checked against
#: whatever build is current, which is the point of the check.
SPELLCONST = Path("builds") / newest_build() / "spellconst"
WRITTEN = {"warrior-fury", "mage-frost"}


def documents():
    return load_all(APL_DIR)


def test_there_is_exactly_one_rotation_per_spec():
    assert set(documents()) == {spec.spec for spec in load_specs(Path("curated"))}


def test_every_rotation_parses_as_the_engines_own_aplrotation():
    """json_format.Parse rejects an unknown field, so this is the engine's schema,
    not one we maintain in parallel."""
    type_field = None
    for key, document in documents().items():
        rotation = parse_rotation(document.rotation)
        if type_field is None:
            type_field = rotation.DESCRIPTOR.fields_by_name["type"].enum_type
        assert rotation.type == type_field.values_by_name["TypeAPL"].number, key


def test_a_written_rotation_round_trips_unchanged():
    """A field the engine silently renames or drops would show up here as a
    document that does not survive its own schema."""
    from google.protobuf import json_format

    for key in WRITTEN:
        document = documents()[key]
        back = json_format.MessageToDict(parse_rotation(document.rotation))
        assert back == document.rotation, key


def test_the_two_launch_specs_are_written_and_the_rest_are_not():
    states = {key: document.state for key, document in documents().items()}
    assert {key for key, state in states.items() if state == "written"} == WRITTEN
    assert set(states.values()) <= APL_STATES


def test_a_written_rotation_has_priorities_and_sourced_notes():
    for key in WRITTEN:
        document = documents()[key]
        assert document.rotation["priorityList"], key
        assert document.notes.strip(), key
        assert document.sources, key
        assert all(source.kind in SOURCE_KINDS for source in document.sources)
        assert all(source.url.startswith("https://") for source in document.sources)


def test_an_unwritten_rotation_is_empty_rather_than_wrong():
    for key, document in documents().items():
        if document.state == "unwritten":
            assert document.rotation == EMPTY_ROTATION, key


def test_every_spell_the_rotations_name_exists_with_that_rank():
    """The engine's own checked-in presets are stale on ranks -- their Frostbolt
    is rank 10 where the client's 25304 is Rank 11. Check against the client."""
    if not SPELLCONST.exists():
        pytest.skip("spellconst has not been generated yet (Task 8)")
    by_class = {
        path.stem: json.loads(path.read_text())["spells"] for path in SPELLCONST.glob("*.json")
    }
    checked = 0
    for key, document in documents().items():
        spells = by_class[key.split("-", 1)[0]]
        for spell_id, rank in action_ids(document.rotation):
            assert str(spell_id) in spells, f"{key} names spell {spell_id}, which the build has not"
            assert spells[str(spell_id)]["rank"] == rank, (
                f"{key} names spell {spell_id} at rank {rank}; "
                f"the build says rank {spells[str(spell_id)]['rank']}"
            )
            checked += 1
    assert checked >= 10, "the written rotations should name at least ten spell references"


#: On-next-swing abilities the fork (wowsims-forever's sim/) registers twice
#: under one spell id: a tag-0 direct spell (`SpellFlagNoOnCastComplete`, no
#: GCD -- meant to fire only when a queued swing lands) and an APL-castable
#: queue action at a different tag (`core.ActionID.WithTag`). Untagged
#: casts resolve to tag 0 (`core.ProtoToActionID` ignores rank), so casting a
#: known on-next-swing spell id at any other tag is a free-swing bug, not a
#: rotation choice.
#:
#: Grepped from wowsims-forever's sim/ for the registration pattern (a spell
#: paired with a `WithTag`-suffixed queue spell flagged `SpellFlagAPL`):
#: - sim/warrior/heroic_strike_cleave.go: Heroic Strike and Cleave queue at
#:   tag 1 (`makeQueueSpellsAndAura`).
#: - sim/druid/_maul.go: Maul queues at tag 1 the same way; the file is
#:   underscore-prefixed (Go excludes it from the build) because bear-form
#:   tank druid is not implemented yet in this fork, so no rotation can name
#:   48480 today, but the id is kept here so one never regresses silently
#:   once the spec ships.
#: - sim/hunter/raptor_strike.go: Raptor Strike is the odd one out -- its
#:   tag-1 spell is an internal damage sub-cast the tag-0 ability calls
#:   itself, and the APL-castable queue (`SpellFlagAPL`) is tag 3
#:   (`hunter.RaptorStrike.WithTag(3)`), not tag 1.
ON_NEXT_SWING_QUEUE_TAG = {
    25286: 1,  # Heroic Strike
    20569: 1,  # Cleave
    48480: 1,  # Maul
    2973: 3,  # Raptor Strike rank 1
    14260: 3,  # Raptor Strike rank 2
    14261: 3,  # Raptor Strike rank 3
    14262: 3,  # Raptor Strike rank 4
    14263: 3,  # Raptor Strike rank 5
    14264: 3,  # Raptor Strike rank 6
    14265: 3,  # Raptor Strike rank 7
    14266: 3,  # Raptor Strike rank 8
}


def test_on_next_swing_casts_use_the_engines_queue_tag():
    """Task follow-up, HIGH: warrior-fury cast Heroic Strike (25286) with no
    tag. The untagged id resolved to the tag-0 direct spell and was cast
    every APL iteration as a free swing: 5,664 DPS instead of ~1,759 with the
    queue tag."""
    checked = 0
    for key, document in documents().items():
        for spell_id, _rank, tag in cast_spell_action_ids(document.rotation):
            if spell_id in ON_NEXT_SWING_QUEUE_TAG:
                assert tag == ON_NEXT_SWING_QUEUE_TAG[spell_id], (
                    f"{key} casts on-next-swing spell {spell_id} with tag {tag}; "
                    f"the engine's APL-castable queue is tag {ON_NEXT_SWING_QUEUE_TAG[spell_id]}"
                )
                checked += 1
    assert checked >= 1, "no written rotation names a known on-next-swing spell"


def test_cast_spell_action_ids_ignores_conditions_and_auras():
    """Only a `castSpell` action's own ActionID carries a tag that matters for
    casting; an `auraIsActive` condition referencing the same spell id is not
    itself a cast and must not be mistaken for one."""
    rotation = {
        "priorityList": [
            {
                "action": {
                    "condition": {"auraIsActive": {"auraId": {"spellId": 25286, "rank": 9}}},
                    "castSpell": {"spellId": {"spellId": 25286, "rank": 9, "tag": 1}},
                }
            },
            {"action": {"castSpell": {"spellId": {"itemId": 13446}}}},
        ]
    }
    assert list(cast_spell_action_ids(rotation)) == [(25286, 9, 1)]


#: Every distinct spell id the two written rotations name, mapped to the
#: ability the rotation's own notes claim it is (e.g. "Death Wish on
#: cooldown."). A rank match alone does not catch a wrong id: Death Wish is
#: rankless, so the rank test above passed 12292 -- Sweeping Strikes --
#: silently when the rotation's own notes said "Death Wish". This sidecar is
#: the check that ties each id back to spells.json by name, not just by rank.
EXPECTED_ABILITY_NAMES = {
    "warrior-fury": {
        25289: "Battle Shout",
        2687: "Bloodrage",
        12328: "Death Wish",
        23894: "Bloodthirst",
        1680: "Whirlwind",
        20662: "Execute",
        25286: "Heroic Strike",
    },
    "mage-frost": {
        25304: "Frostbolt",
    },
}


def test_every_named_spell_matches_its_own_label_by_name():
    """Task 10 review, HIGH: warrior-fury's "Death Wish on cooldown" action cast
    12292, which builds/1.60.1.69893/spells.json names Sweeping Strikes; Death
    Wish is 12328. Cross-check every id a written rotation names against the
    build's own spell names, not just its rank."""
    spells_path = Path("builds") / newest_build() / "spells.json"
    if not spells_path.exists():
        pytest.skip("spells.json has not been generated yet")
    by_id = {entry["id"]: entry["name"] for entry in json.loads(spells_path.read_text())}
    for key, expected in EXPECTED_ABILITY_NAMES.items():
        document = documents()[key]
        named_ids = {spell_id for spell_id, _rank in action_ids(document.rotation)}
        assert named_ids == set(expected), (
            f"{key}'s sidecar of expected ability names is out of date with its rotation"
        )
        for spell_id, ability_name in expected.items():
            assert by_id.get(spell_id) == ability_name, (
                f"{key} names spell {spell_id} as {ability_name!r}; "
                f"spells.json calls it {by_id.get(spell_id)!r}"
            )


def test_an_unknown_field_in_a_rotation_is_rejected():
    with pytest.raises(AplError, match="priorityLst"):
        parse_rotation({"type": "TypeAPL", "priorityLst": []})


def test_a_document_whose_spec_does_not_match_its_filename_is_rejected(tmp_path: Path):
    path = tmp_path / "warrior-fury.json"
    path.write_text(
        json.dumps(
            {
                "spec": "mage-frost",
                "state": "unwritten",
                "notes": "",
                "sources": [],
                "rotation": EMPTY_ROTATION,
            }
        )
    )
    with pytest.raises(AplError, match="warrior-fury"):
        load_apl(path)


def test_an_unknown_state_is_rejected(tmp_path: Path):
    path = tmp_path / "warrior-fury.json"
    path.write_text(
        json.dumps(
            {
                "spec": "warrior-fury",
                "state": "drafted",
                "notes": "",
                "sources": [],
                "rotation": EMPTY_ROTATION,
            }
        )
    )
    with pytest.raises(AplError, match="drafted"):
        load_apl(path)


def test_a_written_rotation_that_cites_nothing_is_rejected(tmp_path: Path):
    path = tmp_path / "warrior-fury.json"
    path.write_text(
        json.dumps(
            {
                "spec": "warrior-fury",
                "state": "written",
                "notes": "Bloodthirst on cooldown.",
                "sources": [],
                "rotation": EMPTY_ROTATION,
            }
        )
    )
    with pytest.raises(AplError, match="cites nothing"):
        load_apl(path)


def test_action_ids_finds_nested_spell_references():
    rotation = {
        "priorityList": [
            {
                "action": {
                    "condition": {"auraIsActive": {"auraId": {"spellId": 25289, "rank": 7}}},
                    "castSpell": {"spellId": {"spellId": 23894, "rank": 4}},
                }
            },
            {"action": {"castSpell": {"spellId": {"itemId": 13446}}}},
        ]
    }
    assert set(action_ids(rotation)) == {(25289, 7), (23894, 4)}


def test_action_ids_treats_a_rankless_reference_as_rank_zero():
    assert list(action_ids({"castSpell": {"spellId": {"spellId": 1680}}})) == [(1680, 0)]
