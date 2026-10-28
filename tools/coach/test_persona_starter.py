"""Tests for the starter persona generator."""

from __future__ import annotations

import re
import sys
from pathlib import Path

import pytest

TOOLS = Path(__file__).resolve().parent
sys.path.insert(0, str(TOOLS))

import persona_starter as ps  # noqa: E402

HEADS = ["## Who they are", "## What they know", "## Drills", "## Voice"]
MAX_DRILL_CHARS = 240


def run_starter(slug: str, tmp_path: Path) -> str:
    out = tmp_path / f"{slug}.md"
    ps.write_starter(slug, out)
    return out.read_text(encoding="utf-8")


def front_matter(text: str) -> dict:
    block = text.split("---\n")[1]
    fm: dict = {}
    key = ""
    for line in block.splitlines():
        if line.startswith("  - "):
            fm[key].append(line[4:])
        else:
            key, _, value = line.partition(":")
            fm[key] = [] if not value.strip() else value.strip()
    return fm


def strip_front_matter(text: str) -> str:
    return text.split("---\n", 2)[2]


def section(text: str, head: str) -> str:
    start = text.index(head) + len(head)
    rest = text[start:]
    nxt = re.search(r"^## ", rest, re.MULTILINE)
    return rest[: nxt.start()] if nxt else rest


def strip_section(text: str, head: str) -> str:
    return text.replace(head + section(text, head), "")


def strip_voice_rule_lines(text: str) -> str:
    return "\n".join(l for l in text.splitlines() if not l.startswith("- Never tell"))


def test_starter_has_four_sections_in_order(tmp_path: Path) -> None:
    out = run_starter("mage-frost", tmp_path)
    assert [l for l in out.splitlines() if l.startswith("## ")] == HEADS


def test_starter_front_matter(tmp_path: Path) -> None:
    fm = front_matter(run_starter("mage-frost", tmp_path))
    assert fm["status"] == "starter" and fm["reviewed"] == "pending" and fm["role"] == "dps"
    assert "data/curated/apl/mage-frost.json" in fm["sources"]


def test_starter_has_no_digits_outside_drills(tmp_path: Path) -> None:
    out = run_starter("mage-frost", tmp_path)
    assert not re.search(r"\d", strip_section(strip_front_matter(out), "## Drills"))


def test_healer_starter_never_says_heal_more_and_skips_idle_and_adherence(tmp_path: Path) -> None:
    out = run_starter("druid-restoration", tmp_path)
    assert "heal more" not in strip_voice_rule_lines(out).lower()
    assert "drill.idle:" not in out and "drill.adherence:" not in out


def test_voice_section_is_the_shared_text(tmp_path: Path) -> None:
    out = run_starter("rogue-combat", tmp_path)
    assert section(out, "## Voice").strip() == (TOOLS / "voice.md").read_text().strip()


def test_every_starter_is_clean_and_matches_specs_role(tmp_path: Path) -> None:
    slugs = ps.write_all(tmp_path)
    assert len(slugs) == 25 and not ps.FULL_SPECS & set(slugs)
    for row in ps.load_json(ps.SPECS_PATH):
        if row["spec"] in ps.FULL_SPECS:
            continue
        text = (tmp_path / f"{row['spec']}.md").read_text()
        assert front_matter(text)["role"] == row["role"]
        assert "!" not in text
        assert not re.search(r"\d", strip_section(strip_front_matter(text), "## Drills"))
        for line in section(text, "## Drills").splitlines():
            assert len(line) - len("- drill.x: ") <= MAX_DRILL_CHARS, line
        assert len(re.findall(r"^- drill\.", text, re.MULTILINE)) == len(ps.ROLE_CODES[row["role"]])
        if row["role"] == "healer":
            assert "heal more" not in strip_voice_rule_lines(text).lower()


def test_generator_is_deterministic_and_committed_files_match(tmp_path: Path) -> None:
    ps.write_all(tmp_path)
    for slug in ps.spec_slugs():
        if slug in ps.FULL_SPECS:
            continue
        assert (tmp_path / f"{slug}.md").read_text() == (ps.COACH_DIR / f"{slug}.md").read_text(), slug


def test_full_spec_is_refused(tmp_path: Path) -> None:
    with pytest.raises(SystemExit):
        ps.write_starter("warrior-fury", tmp_path / "x.md")


def test_only_the_first_sentence_is_used() -> None:
    note = (
        "Sinister Strike builds when Backstab cannot be cast, in the engine. "
        "Held until 60 energy so a dagger build never spends the energy here."
    )
    assert ps.usable_sentence(note) is None


def test_parentheticals_are_stripped_not_fatal() -> None:
    note = "Wrack (talent, learned at level 40) is an instant six second DoT, recast in its last second."
    assert ps.usable_sentence(note) == "Wrack is an instant six second DoT, recast in its last second."


def test_numbers_become_words() -> None:
    assert ps.usable_sentence("Shoot with your wand once mana is under 35% of the bar.") == (
        "Shoot with your wand once mana is under thirty-five percent of the bar."
    )


def test_dangling_references_are_dropped() -> None:
    assert ps.usable_sentence("A hurt party member gets Rejuvenation, at the same pace.") is None


def test_priority_intro_states_the_line_count(tmp_path: Path) -> None:
    out = run_starter("druid-restoration", tmp_path)
    total, kept = ps.priority_lines("druid-restoration")
    assert f"holds {ps.number_words(total)} lines" in out
    assert kept and all(line[0].isupper() for line in kept)
    assert "Held until" not in run_starter("rogue-combat", tmp_path)
