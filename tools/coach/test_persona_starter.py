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


def line_of(text: str, prefix: str) -> str:
    return next(l for l in text.splitlines() if l.startswith(f"- {prefix}"))


def know(slug: str, tmp_path: Path) -> str:
    return section(run_starter(slug, tmp_path), "## What they know")


def test_priority_lines_open_with_the_button(tmp_path: Path) -> None:
    for slug, ability in [
        ("mage-arcane", "Arcane Blast"),
        ("rogue-combat", "Sinister Strike"),
        ("hunter-survival", "Raptor Strike"),
        ("druid-restoration", "Swiftmend"),
    ]:
        lines = ps.priority_lines(ps.spec_row(slug))
        assert any(line.startswith(ability) for line in lines), (slug, lines)
        assert all(line[0].isupper() for line in lines)
    assert "Held until" not in run_starter("rogue-combat", tmp_path)


def test_survival_hunter_is_a_melee_priority(tmp_path: Path) -> None:
    for button in ("Raptor Strike", "Mongoose Bite", "Lacerate", "Strider Kick"):
        assert button in know("hunter-survival", tmp_path)


def test_consumables_follow_the_spec_not_the_role(tmp_path: Path) -> None:
    mage = know("mage-frost", tmp_path)
    assert "Flask of Supreme Power" in mage and "Brilliant Wizard Oil" in mage and "Major Mana Potion" in mage
    assert "Mighty Rage" not in mage and "Grilled Squid" not in mage and "Sharpening" not in mage
    rogue = know("rogue-combat", tmp_path)
    assert "Wizard Oil" not in rogue and "Sharpening" not in rogue and "Potion" not in rogue
    assert "Chromatic" not in know("paladin-protection", tmp_path)
    assert "Sharpening" in know("warrior-arms", tmp_path)


def test_resource_drill_matches_the_resource(tmp_path: Path) -> None:
    assert "cap" not in line_of(run_starter("mage-frost", tmp_path), "drill.resource")
    assert "cap" in line_of(run_starter("rogue-combat", tmp_path), "drill.resource")
    assert "mana" in line_of(run_starter("shaman-restoration", tmp_path), "drill.resource")


def test_findings_use_their_titles_not_ids(tmp_path: Path) -> None:
    mage = know("mage-frost", tmp_path)
    assert "- Core abilities under-used:" in mage and "core_casts" not in mage


def test_healer_core_casts_drill_is_true_for_every_healer(tmp_path: Path) -> None:
    text = line_of(run_starter("shaman-restoration", tmp_path), "drill.core_casts")
    assert "pair" not in text and "most expensive heal" not in text


def test_claims_the_guide_does_not_make_are_dropped() -> None:
    assert ps.corroborated_note("Regrowth is the most healing per mana in the kit.", set()) is None
    assert ps.corroborated_note("Maul whenever fifteen rage is on hand, queued on the swing.", {"15"}) is None
    assert ps.corroborated_note("Primal Bite on cooldown: the biggest threat hit per rage.", set()) is not None


def test_seconds_are_spelled() -> None:
    assert ps.usable_sentence("Lacerate again when under 3s remain on the bleed.") == (
        "Lacerate again when under three seconds remain on the bleed."
    )
