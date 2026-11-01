"""The coach rule-file generator: restamps, detects drift, and agrees with the committed files."""

import json
from pathlib import Path

import pytest

from pipeline.coachrules import SPECS, drifted_rule_files, render_spec, write_rule_files

DATA = Path(__file__).resolve().parent.parent
ACTIVE = json.loads((DATA.parent / "web/src/data/active-build.json").read_text())["build"]


def test_every_spec_has_an_apl_and_a_rule_list():
    assert len(SPECS) == 28
    for spec in SPECS:
        assert (DATA / "curated" / "apl" / f"{spec}.json").is_file()


def test_committed_files_match_the_active_build():
    assert drifted_rule_files(ACTIVE, DATA) == []


def test_every_stamp_names_the_requested_build():
    doc = json.loads(render_spec("warrior-fury", ACTIVE, DATA))
    assert doc["build"] == ACTIVE
    stamps = [e["verified"] for e in doc["core"]] + [a["verified"] for a in doc["auras"]]
    assert stamps
    assert all(s.startswith(f"{ACTIVE}/spellconst/warrior.json#") for s in stamps)


def test_write_then_check_is_clean_and_edits_are_drift(tmp_path):
    (tmp_path / "curated" / "coach").mkdir(parents=True)
    (tmp_path / "curated" / "apl").symlink_to(DATA / "curated" / "apl")
    (tmp_path / "curated" / "book-ranks.json").symlink_to(DATA / "curated" / "book-ranks.json")
    (tmp_path / "builds").symlink_to(DATA / "builds")
    paths = write_rule_files(ACTIVE, tmp_path)
    assert len(paths) == len(SPECS)
    assert drifted_rule_files(ACTIVE, tmp_path) == []
    paths[0].write_text("{}\n")
    assert drifted_rule_files(ACTIVE, tmp_path) == [paths[0]]


def test_unknown_spec_is_refused():
    with pytest.raises(KeyError):
        render_spec("warrior-nope", ACTIVE, DATA)


def test_no_downrank_is_judged_against_an_ahnqiraj_book_rank():
    """A Phase 1 player cannot learn a book rank (curated/book-ranks.json), so a downrank
    is judged against the top trainer rank (round 6 healers: Renew 25315, Prayer of
    Healing 25316)."""
    books = {b["id"] for b in json.loads((DATA / "curated" / "book-ranks.json").read_text())["books"]}
    for spec in SPECS:
        doc = json.loads(render_spec(spec, ACTIVE, DATA))
        for key in ("downranks",):
            for entry in doc.get(key) or []:
                top = int(entry["verified"].rsplit("#", 1)[1])
                assert top not in books, (spec, key, entry.get("family") or entry.get("name"), top)
