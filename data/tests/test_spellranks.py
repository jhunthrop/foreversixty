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

FIXTURES = Path(__file__).parent / "fixtures" / "spellconst"


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
