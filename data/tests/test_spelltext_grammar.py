"""The client's description grammar, each case read off a real 1.60.1.70790 spell.

Every expected string is what Wowhead's Forever tooltip prints for that spell
(nether.wowhead.com/tooltip/spell/<id>?dataEnv=16, read 2026-11-10), unless the test
says otherwise. Effect values are the build's own SpellEffect rows.
"""

from pathlib import Path

import pytest

from pipeline.spelltext import (
    Effect,
    ExtraRows,
    SpellRow,
    SpellText,
    SpellTextError,
    load_spell_text,
)
from pipeline.spelltext_gaps import load_spelltext_gaps
from pipeline.spelltext_grammar import (
    agree_plurals,
    condition_holds,
    parse_variables,
    resolve_conditions,
    work_out,
)


def _effect(points: int, **fields) -> Effect:
    return Effect(base_points=points, die_sides=0, period_ms=0, **fields)


def _spell(description: str, *effects: Effect, **fields) -> str:
    row = SpellRow(
        description=description,
        duration_ms=fields.pop("duration_ms", None),
        icon_file_id=0,
        effects=dict(enumerate(effects)),
        **fields,
    )
    return SpellText({1: row}).describe(1)


# --- the reported bug --------------------------------------------------------------------


def test_booming_voice_reads_the_clients_own_words():
    """The site's "Rage cost by of your Shouts" is the client's text, not a dropped value.
    Booming Voice 12321's Spell.Description_lang in 1.60.1.70790's hotfix push is
    "Increases the area of effect of your Shouts by $s1% and reduces the Rage cost by of
    your Shouts and Howls by $s2%." and Wowhead's Forever tooltip prints rank 5 as
    "Increases the area of effect of your Shouts by 50% and reduces the Rage cost by of
    your Shouts and Howls by 25%."; both of its values render."""
    template = (
        "Increases the area of effect of your Shouts by $s1% and reduces the Rage cost by of "
        "your Shouts and Howls by $s2%."
    )
    text = SpellText({12321: SpellRow(template, None, 0, {0: _effect(10), 1: _effect(5)})})
    assert text.describe(12321, {0: 50, 1: 25}) == (
        "Increases the area of effect of your Shouts by 50% and reduces the Rage cost by of "
        "your Shouts and Howls by 25%."
    )


# --- conditions ---------------------------------------------------------------------------


def test_a_chain_of_conditions_reads_its_last_branch():
    # Dormant Heart of the Mountain 249473's item effect.
    text = "Your casts of $?s2060[Greater Heal]?s5185[Healing Touch][Greater Heal or Healing Touch]"
    assert resolve_conditions(text) == "Your casts of Greater Heal or Healing Touch"


def test_the_last_branch_may_stand_after_a_line_break():
    # Chronostatic Preservation 436516 prints only its "Fuses ..." branch.
    text = "$?a443369[Heals a friendly target.]\r\n\r\n[Fuses Arcane, Fire, and Frost magic.]"
    assert resolve_conditions(text) == "Fuses Arcane, Fire, and Frost magic."


def test_a_negated_condition_holds_for_a_reader_without_the_spell():
    assert condition_holds("!s12345")
    assert not condition_holds("s12345")
    assert resolve_conditions("$?!s12345[yes][no]") == "yes"


def test_a_level_test_reads_level_sixty():
    # SpellDescriptionVariables 840's $buff.
    text = "$?$PL>=60[89]$?$PL>=56[74]$?$PL>52[60][6]"
    assert resolve_conditions(text) == "89"
    assert condition_holds("$PL>42")
    assert not condition_holds("$PL<20")


def test_nested_brackets_inside_a_branch_stay_whole():
    assert resolve_conditions("$?s1[a [b] c][d]e") == "de"
    assert resolve_conditions("$?!s1[a [b] c][d]e") == "a [b] ce"


# --- description variables ----------------------------------------------------------------


def test_variables_parse_one_per_line_and_run_together():
    cell = "$mult1=$?s17815[${1.05}][${1.0}]\n$mult=$?s17836[${1.25}][${$<mult1>}]"
    assert parse_variables(cell) == {
        "mult1": "$?s17815[${1.05}][${1.0}]",
        "mult": "$?s17836[${1.25}][${$<mult1>}]",
    }
    # Row 985 writes its variables with no line break between them.
    assert parse_variables("$sk=${$sk43+$skm43}$sc=${$PL*5}") == {
        "sk": "${$sk43+$skm43}",
        "sc": "${$PL*5}",
    }
    assert parse_variables("$1h=${0.0765*$SPH}") == {"1h": "${0.0765*$SPH}"}


def test_frostbolt_reads_its_damage_multiplier_from_its_variables():
    """Frostbolt 116: "${$m2$*$<frostdamage>} to ${$M2*$<frostdamage>}" (the `$*` is the
    client's own typo). Effect 2 is 19 with variance 0.105 and 0.2 a level over levels 4
    to 8 (Wowhead: "(18.8 * 1) to (20.8 * 1)"); `$<frostdamage>` is 1.0 without Ice Shards,
    so a sentence reads 19 to 21."""
    variables = parse_variables(
        "$frostdamage1=$?s11151[${1.02}][${1.0}]\n"
        "$frostdamage2=$?s12952[${1.04}][${$<frostdamage1>}]\n"
        "$frostdamage=$?s12953[${1.06}][${$<frostdamage2>}]"
    )
    frost = _effect(19, variance=0.10526315868, level_bonus=0.2 * 4)
    text = _spell(
        "causing ${$m2$*$<frostdamage>} to ${$M2*$<frostdamage>}  Frost damage",
        _effect(-40),
        frost,
        variables=variables,
    )
    # The client's own two spaces before "Frost" print as the one a page shows.
    assert text == "causing 19 to 21 Frost damage"


def test_a_level_formula_variable_is_read_at_level_sixty():
    """Shadow Word: Death 401955: "${$<damagepower>*$m1/100} to ${$<damagepower>*$M1/100}",
    $damagepower = 9.456667 + 0.635108*$PL + 0.039063*$PL*$PL, m1 346 with variance 0.15.
    Wowhead: "[(9.456667 + 0.635108 * 60 + 0.039063 * 60 * 60) * 320 / 100] to [... * 372
    / 100]", which is 602.2 to 700.1."""
    variables = parse_variables("$damagepower=${9.456667+0.635108*$PL+0.039063*$PL*$PL}")
    text = _spell(
        "inflicts ${$<damagepower>*$m1/100} to ${$<damagepower>*$M1/100} Shadow damage",
        _effect(346, variance=0.15028901398),
        variables=variables,
    )
    assert text == "inflicts 602 to 700 Shadow damage"


def test_an_undefined_variable_prints_nothing():
    # Earth Shield 974 has no variables: "heal the shielded target for ."
    assert _spell("heal the shielded target for $<heal>.") == "heal the shielded target for ."


# --- formulas ---------------------------------------------------------------------------


def test_combo_point_formulas_read_points_per_combo_point():
    """Rip 1079: "${6*($m1+$b1)} damage" with m1 4 and EffectPointsPerResource 5.5 prints
    "57 damage" for 1 point and 189 for 5."""
    rip = _effect(4, points_per_resource=5.5)
    assert _spell("1 point  : ${6*($m1+$b1)} damage", rip) == "1 point  : 57 damage"
    assert _spell("5 points: ${6*($m1+5*$b1)} damage", rip) == "5 points: 189 damage"


def test_a_negative_amount_keeps_its_sign_inside_a_formula():
    # Expose Armor 8647: "${$b*-1} Armor" with -90 a combo point prints "90 Armor".
    assert _spell("${$b*-1} Armor", _effect(0, points_per_resource=-90)) == "90 Armor"


def test_a_talent_multiplier_of_one_drops_out_of_a_formula():
    # Prayer of Mending 401859: Wowhead's "[(172 + (Healing * 0.42899999)) * (1 * 1)]".
    assert work_out("(172+(bh*0.429))*(1*1)", None, {"bh": "Healing"}) == (
        "(172 + Healing * 0.429)"
    )


def test_a_formula_keeps_its_brackets_where_precedence_needs_them():
    words = {"AP": "Attack Power", "mcrit": "Critical Strike chance"}
    assert work_out("(63+102+0.04*AP)*(mcrit/100+1)", None, words) == (
        "((165 + 0.04 * Attack Power) * (Critical Strike chance / 100 + 1))"
    )
    assert work_out("100-(AP-5)", None, words) == "(100 - (Attack Power - 5))"


def test_the_readers_level_is_sixty():
    # Blade Dance 400012: "beyond ${$PL*5}" prints "(60 * 5)"; worked out, 300.
    assert _spell("defense skill beyond ${$PL*5}.") == "defense skill beyond 300."
    assert _spell("Level $PL target:") == "Level 60 target:"


# --- one effect's facts -------------------------------------------------------------------


def test_amplitude_chain_amplitude_and_target_level():
    # Mana Shield 1463: "Drains $e mana per damage absorbed." with EffectAmplitude 2.
    assert _spell("Drains $e mana per damage.", _effect(120, amplitude=2)) == (
        "Drains 2 mana per damage."
    )
    # Thornspike 1293183: "${$f*100}%" with EffectChainAmplitude 0.5.
    assert _spell("reduces the damage by ${$f*100}%.", _effect(45, chain_amplitude=0.5)) == (
        "reduces the damage by 50%."
    )
    # Mind Soothe 453: "level $v or lower" with MaxTargetLevel 40.
    assert _spell("targets level $v or lower.", _effect(-10), max_target_level=40) == (
        "targets level 40 or lower."
    )


def test_home_is_the_hearthstone_location():
    # Astral Recall 556.
    assert _spell("back to $z.") == "back to Hearthstone Location."


def test_variance_spreads_a_sentence_value_into_a_range():
    # Mana Burn 2691 prints "308 to 312" for 30 with variance 0.133 and 280 from its levels.
    assert _spell(
        "consumes up to $s1 mana.", _effect(30, variance=0.13333334, level_bonus=280)
    ) == ("consumes up to 308 to 312 mana.")


def test_a_level_gain_is_capped_at_the_spells_max_level():
    """Immolate 348 effect 1: 8 and 0.7 a level, levels 1 to 5, prints 10.8 on Wowhead;
    a sentence prints it whole."""
    spell = {"ID": "348", "Description_lang": "Burns the enemy for $s1 Fire damage."}
    misc = {"SpellID": "348", "DurationIndex": "0", "SpellIconFileDataID": "0"}
    effect = {
        "SpellID": "348",
        "EffectIndex": "0",
        "EffectBasePointsF": "8",
        "EffectAuraPeriod": "0",
        "EffectRealPointsPerLevel": "0.69999998808",
    }
    levels = {"SpellID": "348", "MaxLevel": "5", "BaseLevel": "1", "SpellLevel": "1"}
    text = load_spell_text([spell], [misc], [effect], [], ExtraRows(levels=[levels]))
    assert text.describe(348) == "Burns the enemy for 11 Fire damage."


def test_a_fractional_amount_is_kept_and_rounded_only_when_printed():
    """Wrack's tick 1349556 on 1.60.1.70938: 25.2 at level 40 plus 1.75 a level to 60 is
    60.2 a tick, six ticks 361.2. The game keeps the float and rounds what it prints:
    60 a tick, and a formula over it works on 60.2 (361), where a base rounded
    first (25) worked on 60 (360)."""
    spell = {"ID": "1316697", "Description_lang": "$s1 a sec, ${$s1*6} in all."}
    misc = {"SpellID": "1316697", "DurationIndex": "1", "SpellIconFileDataID": "0"}
    effect = {
        "SpellID": "1316697",
        "EffectIndex": "0",
        "EffectBasePointsF": "25.2000007629",
        "EffectAuraPeriod": "1000",
        "EffectRealPointsPerLevel": "1.75",
    }
    levels = {"SpellID": "1316697", "MaxLevel": "60", "BaseLevel": "40", "SpellLevel": "40"}
    durations = [{"ID": "1", "Duration": "6000"}]
    text = load_spell_text([spell], [misc], [effect], durations, ExtraRows(levels=[levels]))
    assert text.describe(1316697) == "60 a sec, 361 in all."


def test_a_spell_with_no_text_of_its_own_still_lends_its_duration():
    # Berserking 20554's "$26635d" reads spell 26635's duration; 26635 has no Spell row.
    spell = {"ID": "20554", "Description_lang": "for $26635d."}
    misc = [
        {"SpellID": "20554", "DurationIndex": "0", "SpellIconFileDataID": "0"},
        {"SpellID": "26635", "DurationIndex": "9", "SpellIconFileDataID": "0"},
    ]
    durations = [{"ID": "9", "Duration": "15000"}]
    assert load_spell_text([spell], misc, [], durations).describe(20554) == "for 15 sec."


# --- other spells and markup --------------------------------------------------------------


def test_a_spell_tooltip_reference_reads_its_name_then_its_text():
    # Feral Charge 1238122 ends with Feral Charge (Cat)'s tooltip.
    rows = {
        1238122: SpellRow("Charge an enemy.\r\n\r\n$@spelltooltip414924", None, 0),
        414924: SpellRow("Leap behind an enemy.", None, 0),
    }
    text = SpellText(rows, {414924: "Feral Charge (Cat)"})
    assert text.describe(1238122) == (
        "Charge an enemy.\r\n\r\nFeral Charge (Cat)\r\nLeap behind an enemy."
    )


def test_colour_escapes_keep_their_words():
    assert _spell("|CFFFF2020Requires Bear Form|R\r\nCharge an enemy.") == (
        "Requires Bear Form\r\nCharge an enemy."
    )


def test_each_plural_agrees_with_its_own_nearest_number():
    # Nature's Grasp 16689: "The next $n melee $Lattack:attacks; that $Lstrikes:strike;".
    assert agree_plurals("The next 1 melee $Lattack:attacks; that $Lstrikes:strike; the") == (
        "The next 1 melee attack that strikes the"
    )
    # Cone of Cold 120: "The first 4 targets hit by this cone $Lis:are; slowed".
    assert agree_plurals("The first 4 targets hit by this cone $Lis:are; slowed") == (
        "The first 4 targets hit by this cone are slowed"
    )


def test_an_enchant_reads_its_enchantments_amounts_and_item_level():
    """Heavy Armor Kit 2833 applies enchantment 8484 (EffectPointsMin 24 and 3, ItemLevelMin
    15); Classic Wowhead prints "by 24. Only usable on items level 15 and above." (the
    Forever one prints 0 for every `$ec` token)."""
    spell = {
        "ID": "2833",
        "Description_lang": "Stamina by $ec2 and armor by $ec1. Only usable on items level "
        "$ecim and above.",
    }
    effect = {
        "SpellID": "2833",
        "EffectIndex": "0",
        "Effect": "53",
        "EffectBasePointsF": "0",
        "EffectAuraPeriod": "0",
        "EffectMiscValue_0": "8484",
    }
    enchantment = {
        "ID": "8484",
        "EffectPointsMin_0": "24",
        "EffectPointsMin_1": "3",
        "EffectPointsMin_2": "0",
        "ItemLevelMin": "15",
    }
    text = load_spell_text([spell], [], [effect], [], ExtraRows(item_enchantments=[enchantment]))
    assert text.describe(2833) == (
        "Stamina by 3 and armor by 24. Only usable on items level 15 and above."
    )


def test_a_multiplier_token_multiplies_first():
    # Execute 5308: "$*10;F1 additional damage" with a chain amplitude of 0.3 prints 3.
    execute = _effect(125, chain_amplitude=0.30000001192)
    assert _spell("into $*10;F1 additional damage.", execute) == "into 3 additional damage."


def test_a_spell_that_lasts_until_cancelled():
    # Find Treasure 2481 (SpellDuration -1): "Lasts until canceled."
    spell = {"ID": "2481", "Description_lang": "Lasts $d."}
    misc = {"SpellID": "2481", "DurationIndex": "21", "SpellIconFileDataID": "0"}
    text = load_spell_text([spell], [misc], [], [{"ID": "21", "Duration": "-1"}])
    assert text.describe(2481) == "Lasts until canceled."


def test_a_periodic_total_without_a_tick_period_divides_by_five_point_two_seconds():
    # Food 433 (17, 18 sec) prints 59 and Drink 1135 (332, 30 sec) 1915 on Wowhead.
    assert _spell("Restores $o1 health over $d.", _effect(17), duration_ms=18_000) == (
        "Restores 59 health over 18 sec."
    )
    assert _spell("Restores $o1 mana over $d.", _effect(332), duration_ms=30_000) == (
        "Restores 1915 mana over 30 sec."
    )


def test_a_stat_comparison_reads_every_stat_at_nothing():
    """Demonic Brand 1293695 prints `$cond($gt($SPFI,$SPS),$SPFI,$SPS)` as "Shadow spell
    power" and `$pl` as 60: "[(((60 - 26) * 1.5) + 14 + (0.078 * ((Shadow spell power))))
    * (1 * 1 * 1)]"."""
    template = "deal ${((($pl-26)*1.5)+14+(0.078*($cond($gt(($SPFI),($SPS)),($SPFI),($SPS)))))}"
    assert _spell(template) == "deal (65 + 0.078 * Shadow spell power)"
    # Divine Aegis 431622: "Shields up to ${20*$pl} damage" prints "(20 * 60)".
    assert _spell("Shields up to ${20*$pl} damage") == "Shields up to 1200 damage"


def test_a_value_that_rounds_away_keeps_its_digits():
    # Set bonus 1226378: `${$m1/-1000}` of 20 prints "0.0200 sec" on Wowhead.
    assert _spell("reduced by ${$m1/-1000} sec", _effect(20)) == "reduced by 0.02 sec"


def test_a_curated_gap_fills_a_token_no_client_table_carries():
    # Berserking 20554 reads spell 26635, which no 1.60.1.70790 table has; Wowhead's Forever
    # tooltip prints "for 15 sec" (curated/spelltext-gaps.json).
    spell = {"ID": "20554", "Description_lang": "speed by 10% for $26635d."}
    text = load_spell_text([spell], [], [], [], None, {20554: {"$26635d": "15 sec"}})
    assert text.describe(20554) == "speed by 10% for 15 sec."


def test_a_curated_gap_the_template_no_longer_holds_is_an_error():
    spell = {"ID": "20554", "Description_lang": "speed by 10% for 15 sec."}
    with pytest.raises(SpellTextError, match="20554"):
        load_spell_text([spell], [], [], [], None, {20554: {"$26635d": "15 sec"}})


def test_the_curated_gaps_file_loads_with_every_entry_sourced():
    gaps = load_spelltext_gaps(Path("curated"))
    assert gaps[20554] == {"$26635d": "15 sec"}
