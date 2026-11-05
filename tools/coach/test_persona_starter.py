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
    # Round 25: the flask, potion and weapon enchant come from the spec's raid preset, as the food
    # does. Frost opens on a Major Spellblasting Potion, not the mana potion its resource suggested;
    # a rogue brings no flask and drinks a Major Frenzy Potion.
    mage = know("mage-frost", tmp_path)
    assert "Flask of Supreme Power" in mage and "Brilliant Wizard Oil" in mage and "Major Spellblasting Potion" in mage
    assert "Mana Potion" not in mage
    assert "Mighty Rage" not in mage and "Grilled Squid" not in mage and "Sharpening" not in mage
    assert "Major Mana Potion" in know("mage-arcane", tmp_path)
    rogue = know("rogue-combat", tmp_path)
    assert "Wizard Oil" not in rogue and "Flask" not in rogue and "Major Frenzy Potion" in rogue
    cat = know("druid-feral", tmp_path)
    assert "Flask" not in cat and "Elemental Sharpening Stone" in cat and "Major Frenzy Potion" in cat
    assert "Major Frenzy Potion" in know("hunter-survival", tmp_path)
    assert "Chromatic" not in know("paladin-protection", tmp_path)
    assert "Sharpening" in know("warrior-arms", tmp_path)


def test_food_is_the_published_presets_food(tmp_path: Path) -> None:
    # Round 9's raid foods (sim/request/raid-consumes.json, top band): casters and the
    # spell-power Protection paladin Baked Salmon, Strength users and the other tanks Smoked
    # Desert Dumplings, rogues and hunters Poached Sunscale Salmon, healers Sage's Tea.
    for slug, food in [
        ("mage-frost", "Baked Salmon"),
        ("warrior-arms", "Smoked Desert Dumplings"),
        ("paladin-protection", "Baked Salmon"),
        ("druid-feral-bear", "Smoked Desert Dumplings"),
        ("rogue-combat", "Poached Sunscale Salmon"),
        ("hunter-marksmanship", "Poached Sunscale Salmon"),
        ("priest-discipline", "Sage's Tea"),
    ]:
        text = know(slug, tmp_path)
        assert f"food ({food})" in text, (slug, food)
        assert "Tender Wolf Steak" not in text, slug


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


def starter_slugs() -> list[str]:
    return [slug for slug in ps.spec_slugs() if slug not in ps.FULL_SPECS]


def names_in(name: str, text: str) -> bool:
    return bool(re.search(rf"\b{re.escape(name)}\b(?!-)", text))


def test_only_abilities_the_guide_names_positively() -> None:
    # The ruled-out set is read from each guide's rotation section, the single source of truth: a
    # button the guide rules out or makes conditional is never named, unless a recorded correction
    # restates the guide's condition for it.
    for slug in starter_slugs():
        row = ps.spec_row(slug)
        ruled_out = ps.ruled_out_of(row) - set(ps.GUIDE_LINES.get(slug, {}))
        text = priority_text(slug)
        for name in ruled_out:
            assert not names_in(name, text), (slug, name)


@pytest.mark.parametrize(
    ("prose", "names", "ruled_out"),
    [
        # feral.md, marksmanship.md, discipline.md, affliction.md, fire.md, assassination.md.
        ("Rake is not worth its energy here: Rip alone keeps the boss bleeding.", {"Rake", "Rip"}, {"Rake"}),
        ("Serpent Sting is not worth a global in this build either.", {"Serpent Sting"}, {"Serpent Sting"}),
        ("Shadowfiend cannot be learned at launch.", {"Shadowfiend"}, {"Shadowfiend"}),
        ("Wrack stays off your raid bars.", {"Wrack", "Shadow Bolt"}, {"Wrack"}),
        ("You have no Ice Block and no Invisibility, so do not pull.", {"Ice Block", "Invisibility"},
         {"Ice Block", "Invisibility"}),
        ("Rupture and Expose Armor do not make the list.", {"Rupture", "Expose Armor"}, {"Rupture", "Expose Armor"}),
        # A comparison is not a rule-out of the far side: frost.md.
        (
            (
                "Frostfire Bolt is not a Frost filler: its hit is smaller than Frostbolt's and half your Frost "
                "talents do nothing for it."
            ),
            {"Frostfire Bolt", "Frostbolt"},
            {"Frostfire Bolt"},
        ),
        # A rule-out phrase about something else in the clause: protection.md.
        ("Divine Shield and Divine Protection come after it, once Forbearance is gone.",
         {"Divine Shield", "Divine Protection"}, set()),
        # Mana to spare is pacing, not a group or talent condition: holy.md.
        ("Holy Light on the tank below 70%, but only with mana to spare.", {"Holy Light"}, set()),
        # Another spec's lack of a button rules nothing out for this one: affliction.md.
        ("Demonology and Destruction have no Amplify Curse, which is why they run Doom.", {"Amplify Curse"}, set()),
    ],
)
def test_guide_rule_outs_are_read(prose: str, names: set[str], ruled_out: set[str]) -> None:
    from starter_guide import excluded_abilities

    assert excluded_abilities(prose, names) == ruled_out


def test_situational_buttons_the_guide_recommends_are_kept() -> None:
    # arms.md: Heroic Strike above ninety rage; feral.md: Claw when you cannot get behind;
    # protection.md: Lay on Hands for the killing swing and Exorcism on Undead and Demons.
    for slug, name in [
        ("warrior-arms", "Heroic Strike"),
        ("druid-feral", "Claw"),
        ("paladin-protection", "Lay on Hands"),
        ("paladin-protection", "Exorcism"),
    ]:
        assert name in priority_names(slug), (slug, name)
        assert names_in(name, ps.rotation_prose(ps.spec_row(slug))), (slug, name)


def test_a_form_named_as_a_state_is_not_a_button() -> None:
    from starter_guide import guide_order

    names = {"Faerie Fire", "Bear Form", "Dire Bear Form", "Barkskin"}
    prose = "Be in Dire Bear Form. Faerie Fire from range. In Bear Form it is free. You can cast Barkskin in Bear Form."
    assert guide_order(prose, names) == ["Dire Bear Form", "Faerie Fire", "Barkskin"]


def test_markdown_structure_never_leaks_into_a_line() -> None:
    for slug in starter_slugs():
        for line in ps.priority_lines(ps.spec_row(slug)):
            assert "\n" not in line and "#" not in line and ": - " not in line, (slug, line)


def test_a_sentence_that_points_back_is_not_a_reason() -> None:
    from starter_priority import guide_candidates

    prose = "Riptide on the tank. It is your cheapest heal per mana, and it sets up the Chain Heal."
    assert guide_candidates(prose, {"Riptide", "Chain Heal"}, frozenset()) == [("Riptide", "Riptide on the tank.")]


def test_thousands_are_spelled() -> None:
    assert ps.usable_sentence("Blizzard costs 1,400 mana, so it is for big pulls.") == (
        "Blizzard costs one thousand four hundred mana, so it is for big pulls."
    )


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
    # Forever druids cannot learn Tiger's Fury; Rip at five combo points, Shred fills the rest.
    feral = priority_names("druid-feral")
    assert "Tiger's Fury" not in feral
    assert feral.index("Shifting Power") < feral.index("Rip") < feral.index("Shred")
    beast = priority_names("hunter-beast-mastery")
    # Forever's Multi-Shot and Aimed Shot share a cooldown. Round 25: the guide and the card now cast
    # Aimed Shot every time that cooldown comes up and keep Multi-Shot for packs and moving, so Aimed
    # Shot is no longer "only the opener" and follows the burst (beast-mastery.md).
    assert beast.index("Bestial Wrath") < beast.index("Aimed Shot") < beast.index("Multi-Shot") < beast.index("Arcane Shot")


def test_healer_resource_drill_is_true_for_every_healer(tmp_path: Path) -> None:
    for slug in ("paladin-holy", "shaman-restoration", "druid-restoration", "priest-discipline"):
        assert "over time" not in line_of(run_starter(slug, tmp_path), "drill.resource")


def test_survival_never_lists_the_ranged_aspect() -> None:
    assert "Aspect of the Falcon" not in priority_names("hunter-survival")


def test_dropped_talents_are_never_cited() -> None:
    assert "Early Demise" not in priority_text("priest-shadow")
    assert "Tiger's Fury" not in priority_text("druid-feral")


def test_figures_must_be_in_the_guide() -> None:
    assert ps.corroborated_note("Stormstrike on cooldown, on an 8 s cooldown.", "an 8-second cooldown") is not None
    assert ps.corroborated_note("Stormstrike on cooldown, on an 8 s cooldown.", "8 targets") is None
    assert ps.corroborated_note("Arcane Power on cooldown: a flat 30% damage cooldown.", "30 seconds") is None


def test_pyroblast_carries_the_guides_condition() -> None:
    assert any(
        line.startswith("Pyroblast") and "only on three stacks of Heating Up" in line
        for line in ps.priority_lines(ps.spec_row("mage-fire"))
    )


def test_swiftmend_carries_the_guides_gate() -> None:
    """Round 11: the starter said only "Swiftmend." where the guide and the rotation gate it
    on the tank under three quarters health with the mana keeping pace."""
    line = next(line for n, line in ps.priority_entries_of(ps.spec_row("druid-restoration")) if n == "Swiftmend")
    assert "under three quarters health" in line, line


def test_shadow_priority_follows_the_launch_guide() -> None:
    """Round 14 (regcast): the Shadow overrides kept the October rotation search's order
    (Inner Focus into Devouring Plague, Mind Blast only when Mind Flay cannot be cast, Shadow
    Word: Death only below twenty percent) after shadow.md and the rotation moved to Inner
    Focus on Mind Blast, Mind Blast and Shadow Word: Death on cooldown ahead of Mind Flay."""
    entries = dict(ps.priority_entries_of(ps.spec_row("priest-shadow")))
    assert "Mind Blast" in entries["Inner Focus"] and "Devouring Plague" not in entries["Inner Focus"]
    assert entries["Mind Blast"].startswith("Mind Blast on cooldown, ahead of Mind Flay")
    assert entries["Shadow Word: Death"].startswith("Shadow Word: Death on cooldown")
    names = [n for n, _ in ps.priority_entries_of(ps.spec_row("priest-shadow"))]
    assert names.index("Mind Blast") < names.index("Mind Flay")
    assert names.index("Shadow Word: Death") < names.index("Mind Flay")


# Round 5 (wow-player read 4 rulings).

def test_totem_conditional_abilities_carry_their_condition() -> None:
    for slug, name in (("hunter-survival", "Wing Clip"), ("warrior-arms", "Hamstring")):
        line = next(line for n, line in ps.priority_entries_of(ps.spec_row(slug)) if n == name)
        assert "only with a Windfury Totem in the group" in line, (slug, line)


def test_guide_order_statements_win() -> None:
    elemental = priority_names("shaman-elemental")
    # Round 25: elemental.md and the card put the cooldowns and Flame Shock ("ahead of everything but
    # your cooldowns") right behind Mana Spring, ahead of Searing Totem.
    assert elemental[:3] == ["Mana Spring Totem", "Lightning Mastery", "Flame Shock"]
    shadow = priority_names("priest-shadow")
    assert shadow.index("Mind Blast") < shadow.index("Mind Flay")
    prot = priority_names("paladin-protection")
    assert prot.index("Holy Shield") < prot.index("Judgement") < prot.index("Consecration")
    # Round 25: feral.md and the card open with Ravage; Shifting Power still sits ahead of the finishers.
    feral = priority_names("druid-feral")
    assert feral[0] == "Ravage" and feral.index("Shifting Power") < feral.index("Rip")


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
    # The guide page shows the rotation data's priority list above its prose, so a button in that
    # list is in the guide too (Survival's Lacerate and Strider Kick).
    for slug, lines in ps.GUIDE_LINES.items():
        row = ps.spec_row(slug)
        prose = ps.rotation_prose(row)
        shown = ps.rotation_cast_names(row)
        for name, line in lines.items():
            assert line.startswith(name) and not re.search(r"\d", line), (slug, line)
            in_guide = name in prose or name in shown or (name == "Shoot" and "wand" in prose.lower())
            assert in_guide, (slug, name)
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


def test_discipline_names_its_overrides_once() -> None:
    """Round 14: GUIDE_LINES named "priest-discipline" twice, so the second entry (Dark
    Sacrifice) silently replaced the first. The dead Binding Heal line is gone: the starter's
    own line from the rotation ("under ninety percent yourself") is the guide's gate."""
    # Round 25 adds the shield (the guide says "Shield the tank") and the lines that name the level a
    # spell arrives (Penance, Prayer of Healing, Inner Focus); round 26 adds Lightwell, trained at forty.
    assert set(ps.GUIDE_LINES["priest-discipline"]) == {
        "Power Word: Shield", "Penance", "Prayer of Healing", "Inner Focus", "Dark Sacrifice", "Lightwell",
    }
    source = (Path(__file__).parent / "starter_overrides.py").read_text()
    # Only GUIDE_LINES is checked: DRILL_LINES names the spec too, for its own drills (round 25).
    guide_lines = source[source.index("GUIDE_LINES:") : source.index("GUIDE_ORDER:")]
    assert guide_lines.count('"priest-discipline": {') == 1
