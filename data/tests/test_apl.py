import json
from pathlib import Path

import pytest

from pipeline.apl import (
    APL_STATES,
    EMPTY_ROTATION,
    AplError,
    action_ids,
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
