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
    book_ranks = json.loads((DATA / "curated" / "book-ranks.json").read_text())
    books = {b["id"] for b in book_ranks["books"]}
    for spec in SPECS:
        doc = json.loads(render_spec(spec, ACTIVE, DATA))
        for key in ("downranks",):
            for entry in doc.get(key) or []:
                top = int(entry["verified"].rsplit("#", 1)[1])
                assert top not in books, (spec, key, entry.get("family") or entry.get("name"), top)


def test_no_rule_names_an_ahnqiraj_book_rank_a_trainer_rank_could_name():
    """Every list's stamped top rank is the one the coach names in a reading (an aura's
    uptime fact carries its rank; F1 falls back to the last id). A book rank is that top only
    when the family has no trainer rank: the coach told hunters "Aspect of the Hawk (Rank
    7)", the Ahn'Qiraj book, when a level-60 trainer teaches rank 6."""
    book_ranks = json.loads((DATA / "curated" / "book-ranks.json").read_text())
    books = {b["id"] for b in book_ranks["books"]}
    for spec in SPECS:
        doc = json.loads(render_spec(spec, ACTIVE, DATA))
        for key in ("core", "auras", "defensives", "efficiency"):
            for entry in doc.get(key) or []:
                top = int(entry["verified"].rsplit("#", 1)[1])
                name = entry.get("family") or entry.get("name")
                if any(i not in books for i in entry["spell_ids"]):
                    assert top not in books, (spec, key, name, top)


def _apl_cast_ids(spec: str) -> set[int]:
    ids: set[int] = set()

    def walk(value: object) -> None:
        if isinstance(value, dict):
            cast = value.get("castSpell")
            if isinstance(cast, dict) and isinstance(cast.get("spellId"), dict):
                ids.add(cast["spellId"]["spellId"])
            for inner in value.values():
                walk(inner)
        elif isinstance(value, list):
            for inner in value:
                walk(inner)

    walk(json.loads((DATA / "curated" / "apl" / f"{spec}.json").read_text()))
    return ids


#: Core cooldowns the rotation fires through autocastOtherCooldowns, never a castSpell line.
AUTOCAST_CORE = {"Blade Flurry"}


def test_every_judged_cooldown_and_dot_is_one_the_rotation_casts():
    """Round 11: Affliction's rules still judged Bane of Doom's cooldown and uptime and
    Wrack's uptime after the rotation moved to Bane of Agony and dropped Wrack, so a player
    following the site lost points for it. A core cooldown or a DoT the coach judges must be
    cast by the spec's own APL, at some rank."""
    for spec in SPECS:
        doc = json.loads(render_spec(spec, ACTIVE, DATA))
        cast = _apl_cast_ids(spec)
        for entry in doc.get("core") or []:
            if entry["family"] in AUTOCAST_CORE:
                continue
            assert set(entry["spell_ids"]) & cast, (spec, "core", entry["family"])
        for aura in doc.get("auras") or []:
            if aura["kind"] == "dot":
                assert set(aura["spell_ids"]) & cast, (spec, "dot", aura["name"])


def test_demoralizing_shout_is_covered_by_a_druids_roar():
    # protection.md leaves the slot to a druid's Roar when one holds it: its time is the slot
    # filled, never a gap in the warrior's uptime (round 26).
    doc = json.loads(render_spec("warrior-protection", ACTIVE, DATA))
    shout = next(a for a in doc["auras"] if a["name"] == "Demoralizing Shout")
    druid = DATA / "builds" / ACTIVE / "spellconst" / "druid.json"
    roar = json.loads(druid.read_text(encoding="utf-8"))["spells"]
    assert shout["covered_by"]
    assert all(roar[str(i)]["name"] == "Demoralizing Roar" for i in shout["covered_by"])
    for spec in ("warrior-fury", "warrior-arms"):
        auras = json.loads(render_spec(spec, ACTIVE, DATA))["auras"]
        assert all("covered_by" not in a for a in auras), spec
