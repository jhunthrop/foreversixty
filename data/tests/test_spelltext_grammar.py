"""The client's description grammar, each case read off a real 1.60.1.70790 spell.

Every expected string is what Wowhead's Forever tooltip prints for that spell
(nether.wowhead.com/tooltip/spell/<id>?dataEnv=16, read 2026-11-10), unless the test
says otherwise. Effect values are the build's own SpellEffect rows.
"""

from pipeline.spelltext import Effect, ExtraRows, SpellRow, SpellText, load_spell_text
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
    assert text == "causing 19 to 21  Frost damage"


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
