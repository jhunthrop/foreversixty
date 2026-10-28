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
MAX_SENTENCE_WORDS = 20


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
    assert fm["status"] == "starter" and fm["reviewed"] == ps.REVIEWED == "wow-player, 2026-10-28" and fm["role"] == "dps"
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
        ("druid-restoration", "Nature's Swiftness"),
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
    assert ps.corroborated_note("Regrowth is the most healing per mana in the kit.", "") is None
    assert ps.corroborated_note("Maul whenever 15 rage is on hand, queued on the swing.", "15 rage") is None
    assert ps.corroborated_note("Primal Bite on cooldown: the biggest threat hit per rage.", "") is None
    assert ps.corroborated_note("Primal Bite on cooldown, ahead of every other button.", "") is not None


def test_seconds_are_spelled() -> None:
    assert ps.usable_sentence("Lacerate again when under 3s remain on the bleed.") == (
        "Lacerate again when under three seconds remain on the bleed."
    )


# Round 4 rulings (P1-P7).

def priority_names(slug: str) -> list[str]:
    return [name for name, _ in ps.priority_entries_of(ps.spec_row(slug))]


def priority_text(slug: str) -> str:
    return "\n".join(ps.priority_lines(ps.spec_row(slug)))


@pytest.mark.parametrize(
    ("slug", "ruled_out"),
    [
        ("druid-balance", ["Moonfire"]),
        ("druid-feral", ["Moonfire", "Wrath", "Claw", "Bear Form", "Rake"]),
        ("druid-feral-bear", ["Feral Charge", "Growl", "Bash"]),
        ("mage-fire", ["Frostfire Bolt"]),
        ("mage-arcane", ["Frostfire Bolt"]),
        ("shaman-elemental", ["Flame Shock", "Earth Shock", "Fire Nova"]),
        ("shaman-enhancement", ["Flame Shock", "Fire Nova"]),
        ("warrior-arms", ["Heroic Strike", "Recklessness"]),
        ("rogue-assassination", ["Sinister Strike", "Venom"]),
        ("paladin-retribution", ["Seal of Command"]),
        ("paladin-protection", ["Lay on Hands", "Exorcism"]),
        ("priest-discipline", ["Shadowfiend", "Dark Sacrifice"]),
        ("hunter-survival", ["Aimed Shot", "Serpent Sting"]),
        ("hunter-marksmanship", ["Serpent Sting"]),
    ],
)
def test_only_abilities_the_guide_names_positively(slug: str, ruled_out: list[str]) -> None:
    text = priority_text(slug)
    for name in ruled_out:
        assert not re.search(rf"\b{re.escape(name)}\b(?!-)", text), (slug, name)


def test_no_changelog_or_rotation_data_talk(tmp_path: Path) -> None:
    ps.write_all(tmp_path)
    banned = re.compile(
        r"this week|now the only|back in the rotation|tracked by slot|fires itself|preset|second Holy Light line|"
        r"share of the fight|raid-ready|new energy tool|new baseline|fraction of a percent",
        re.IGNORECASE,
    )
    for path in tmp_path.glob("*.md"):
        assert not banned.search(section(path.read_text(), "## What they know")), path.stem


def test_each_ability_once_and_every_line_is_about_it() -> None:
    for slug in ps.spec_slugs():
        if slug in ps.FULL_SPECS:
            continue
        entries = ps.priority_entries_of(ps.spec_row(slug))
        names = [name for name, _ in entries]
        assert len(names) == len(set(names)), slug
        for name, line in entries:
            assert name in line, (slug, line)
            assert not (line.split()[0] == "The" and line.rstrip(".").endswith(name)), (slug, line)


def test_order_follows_the_guide() -> None:
    bear = priority_names("druid-feral-bear")
    assert bear.index("Primal Bite") < bear.index("Lacerate") < bear.index("Maul") < bear.index("Barkskin")
    feral = priority_names("druid-feral")
    assert feral.index("Tiger's Fury") < feral.index("Rip") < feral.index("Shred")
    beast = priority_names("hunter-beast-mastery")
    assert beast.index("Bestial Wrath") < beast.index("Aimed Shot") < beast.index("Arcane Shot")


def test_healer_resource_drill_is_true_for_every_healer(tmp_path: Path) -> None:
    for slug in ("paladin-holy", "shaman-restoration", "druid-restoration", "priest-discipline"):
        assert "over time" not in line_of(run_starter(slug, tmp_path), "drill.resource")


def test_dropped_talents_are_never_cited() -> None:
    assert "Early Demise" not in priority_text("priest-shadow")
    assert "Moonfire" not in priority_text("druid-balance")


def test_figures_must_be_in_the_guide() -> None:
    assert ps.corroborated_note("Stormstrike on cooldown, on an 8 s cooldown.", "an 8-second cooldown") is not None
    assert ps.corroborated_note("Stormstrike on cooldown, on an 8 s cooldown.", "8 targets") is None
    assert ps.corroborated_note("Arcane Power on cooldown: a flat 30% damage cooldown.", "30 seconds") is None


def test_pyroblast_carries_the_guides_condition() -> None:
    assert any(line.startswith("Pyroblast only") for line in ps.priority_lines(ps.spec_row("mage-fire")))


# Round 5 (wow-player read 4 rulings).

def test_totem_conditional_abilities_carry_their_condition() -> None:
    for slug, name in (("hunter-survival", "Wing Clip"), ("warrior-arms", "Hamstring")):
        line = next(line for n, line in ps.priority_entries_of(ps.spec_row(slug)) if n == name)
        assert "only with a Windfury Totem in the group" in line, (slug, line)


def test_guide_order_statements_win() -> None:
    elemental = priority_names("shaman-elemental")
    assert elemental[:2] == ["Mana Spring Totem", "Lava Burst"]
    shadow = priority_names("priest-shadow")
    assert shadow.index("Mind Flay") < shadow.index("Mind Blast")
    prot = priority_names("paladin-protection")
    assert prot.index("Holy Shield") < prot.index("Judgement") < prot.index("Consecration")
    assert priority_names("druid-feral")[0] == "Shifting Power"


def test_no_bare_line_repeats_a_named_ability() -> None:
    for slug in ps.spec_slugs():
        if slug in ps.FULL_SPECS:
            continue
        entries = ps.priority_entries_of(ps.spec_row(slug))
        for name, line in entries:
            if line == f"{name}.":
                others = [o for n, o in entries if n != name]
                assert not any(re.search(rf"\b{re.escape(name)}\b", o) for o in others), (slug, name)


def test_no_tool_speak(tmp_path: Path) -> None:
    ps.write_all(tmp_path)
    for path in tmp_path.glob("*.md"):
        text = path.read_text()
        assert "curated rotation" not in text and "reserve floor" not in text, path.stem


def test_every_correction_restates_its_guide() -> None:
    for slug, lines in ps.GUIDE_LINES.items():
        prose = ps.section_after(ps.read_text(ps.guide_path(ps.spec_row(slug))), ps.ROTATION_HEADING)
        for name, line in lines.items():
            assert line.startswith(name) and not re.search(r"\d", line), (slug, line)
            assert name in prose or (name == "Shoot" and "wand" in prose.lower()), (slug, name)
    for slug, order in ps.GUIDE_ORDER.items():
        assert set(order) <= set(priority_names(slug)) | {"Holy Strike"}, slug


def test_guide_lines_are_short() -> None:
    for lines in ps.GUIDE_LINES.values():
        for line in lines.values():
            for sentence in re.split(r"(?<=\.) ", line):
                assert len(sentence.split()) <= MAX_SENTENCE_WORDS, sentence


def test_long_sentences_split_where_the_part_reads_alone() -> None:
    from starter_text import short_sentences

    line = (
        "Bane of Doom is the one curse kept up: it deals its whole damage when the minute is out, "
        "so recast it the moment it fires."
    )
    assert short_sentences(line, "Bane of Doom").startswith("Bane of Doom is the one curse kept up. It deals")
    prefixed = "Wrath: Eclipse, if talented, rewards weaving in a Wrath cast: it shortens your next two Starfire casts a lot."
    assert short_sentences(prefixed, "Wrath").startswith("Wrath: Eclipse")
