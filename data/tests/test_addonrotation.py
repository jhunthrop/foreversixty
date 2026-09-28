import json
import re
from pathlib import Path

import pytest

from pipeline.addonrotation import (
    CONDITION_MAX_CHARS,
    LEVEL_BANDS,
    AddonRotationError,
    _one_line,
    build_rotation,
    build_rotations,
)

BUILD = "1.60.1.70009"

SPELLRANKS = {
    "build": BUILD,
    "classes": {
        "warrior": {
            "Battle Shout": [
                {"id": 6673, "level": 1, "rank": 1},
                {"id": 5242, "level": 12, "rank": 2},
                {"id": 25289, "level": 60, "rank": 7},
            ],
            "Bloodrage": [{"id": 2687, "level": 1, "rank": 0}],
            "Bloodthirst": [
                {"id": 23881, "level": 40, "rank": 1},
                {"id": 23894, "level": 60, "rank": 4},
            ],
            "Whirlwind": [{"id": 1680, "level": 36, "rank": 0}],
        }
    },
}


def _write_apl(curated_dir: Path, spec: str, priority_list: list[dict]) -> None:
    apl_dir = curated_dir / "apl"
    apl_dir.mkdir(parents=True, exist_ok=True)
    (apl_dir / f"{spec}.json").write_text(
        json.dumps({"spec": spec, "rotation": {"priorityList": priority_list}}),
        encoding="utf-8",
    )


def _bands(spec: str, curated_dir: Path) -> dict:
    """spec's rotation, keyed by level, against the test module's SPELLRANKS."""
    return {band.level: band.lines for band in build_rotation(spec, curated_dir, SPELLRANKS)}


def _cast(spell_id: int, rank: int | None = None, notes: str = "") -> dict:
    spell = {"spellId": spell_id} if rank is None else {"spellId": spell_id, "rank": rank}
    return {"notes": notes, "action": {"castSpell": {"spellId": spell}}}


def test_level_bands_match_the_ladders_own_rungs():
    """sim/request/ladder.go's ladderLevels is the one source; this module's
    LEVEL_BANDS is a copy (Go and Python do not import each other), and this
    test is what keeps the copy honest."""
    source = Path("../sim/request/ladder.go").read_text(encoding="utf-8")
    match = re.search(r"var ladderLevels = \[\]int\{([^}]*)\}", source)
    assert match, "ladderLevels literal not found in ladder.go"
    levels = [int(part.strip()) for part in match.group(1).split(",")]
    assert LEVEL_BANDS == levels


def test_a_line_resolves_to_the_highest_rank_learned_at_each_band(tmp_path):
    _write_apl(tmp_path, "warrior-fury", [
        _cast(2687, notes="Bloodrage on cooldown."),
        _cast(25289, rank=7, notes="Battle Shout up."),
    ])
    bands = _bands("warrior-fury", tmp_path)
    assert [line.spell_id for line in bands[10]] == [2687, 6673]
    assert [line.spell_id for line in bands[20]] == [2687, 5242]
    assert [line.spell_id for line in bands[60]] == [2687, 25289]


def test_an_ability_not_learned_yet_is_left_out_of_that_band_only(tmp_path):
    _write_apl(tmp_path, "warrior-fury", [_cast(23894, rank=4, notes="Bloodthirst on cooldown.")])
    bands = _bands("warrior-fury", tmp_path)
    assert bands[30] == []
    assert [line.spell_id for line in bands[40]] == [23881]
    assert [line.spell_id for line in bands[60]] == [23894]


def test_a_strict_sequence_unfolds_into_one_line_per_step_sharing_the_notes(tmp_path):
    _write_apl(tmp_path, "warrior-fury", [
        {
            "notes": "Opener.",
            "action": {"strictSequence": {"actions": [
                {"castSpell": {"spellId": {"spellId": 2687}}},
                {"castSpell": {"spellId": {"spellId": 1680}}},
            ]}},
        },
    ])
    bands = _bands("warrior-fury", tmp_path)
    assert [line.spell_id for line in bands[38]] == [2687, 1680]
    assert bands[38][1].condition == "Opener."
    assert bands[30] == [] or [line.spell_id for line in bands[30]] == [2687]


def test_an_id_spellranks_does_not_track_is_skipped_rather_than_erroring(tmp_path):
    _write_apl(tmp_path, "warrior-fury", [_cast(999999, notes="A talent proc, not in spellranks.")])
    bands = _bands("warrior-fury", tmp_path)
    assert all(band == [] for band in bands.values())


def test_actions_with_no_spell_id_name_no_line(tmp_path):
    _write_apl(tmp_path, "warrior-fury", [
        {"action": {"autocastOtherCooldowns": {}}},
        {"notes": "Seal.", "action": {"condition": {"cmp": {}}, "castPaladinPrimarySeal": {}}},
    ])
    bands = build_rotation("warrior-fury", tmp_path, SPELLRANKS)
    assert all(band.lines == [] for band in bands)


def test_every_band_from_level_bands_is_present_in_order(tmp_path):
    _write_apl(tmp_path, "warrior-fury", [])
    bands = build_rotation("warrior-fury", tmp_path, SPELLRANKS)
    assert [band.level for band in bands] == LEVEL_BANDS


def test_missing_apl_refuses(tmp_path):
    with pytest.raises(AddonRotationError, match="warrior-fury"):
        build_rotation("warrior-fury", tmp_path, SPELLRANKS)


def test_build_rotations_covers_every_curated_spec(tmp_path):
    (tmp_path / BUILD).mkdir()
    (tmp_path / BUILD / "spellranks.json").write_text(json.dumps(SPELLRANKS), encoding="utf-8")
    rotations = build_rotations(tmp_path, BUILD, curated_dir=Path("curated"))
    assert "warrior-fury" in rotations
    assert len(rotations) == len(list(Path("curated/apl").glob("*.json")))


def test_missing_spellranks_refuses(tmp_path):
    (tmp_path / BUILD).mkdir()
    with pytest.raises(AddonRotationError, match="spellranks.json"):
        build_rotations(tmp_path, BUILD, curated_dir=Path("curated"))


def test_one_line_keeps_a_single_sentence_whole():
    assert _one_line("Frostbolt is the whole rotation.") == "Frostbolt is the whole rotation."


def test_one_line_cuts_at_the_first_sentence_boundary():
    assert _one_line("First bit. Second bit that a player never needs to read.") == "First bit."


def test_one_line_caps_a_run_on_sentence_with_no_boundary():
    long_sentence = "a" * (CONDITION_MAX_CHARS + 50)
    result = _one_line(long_sentence)
    assert len(result) == CONDITION_MAX_CHARS
    assert result.endswith("…")


def test_a_line_carries_only_the_first_sentence_of_a_long_curated_note(tmp_path):
    long_note = (
        "Short reason. A paragraph of engine-debugging narration that a "
        "player opening the rotation card has no use for and should never see."
    )
    _write_apl(tmp_path, "warrior-fury", [_cast(2687, notes=long_note)])
    bands = _bands("warrior-fury", tmp_path)
    assert bands[10][0].condition == "Short reason."


def test_the_real_committed_data_produces_a_rotation_for_every_spec():
    """End-to-end against the real curated files and the real build's
    spellranks.json, the same inputs `make` regenerates Data.lua from."""
    rotations = build_rotations(Path("builds"), BUILD)
    fury = {band.level: band.lines for band in rotations["warrior-fury"]}
    names_at_60 = {line.name for line in fury[60]}
    assert "Bloodthirst" in names_at_60
    assert "Heroic Strike" in names_at_60
    # Bloodthirst is not learned until 40; a real leveling player's rotation
    # card must not show it earlier.
    names_at_30 = {line.name for line in fury[30]}
    assert "Bloodthirst" not in names_at_30
