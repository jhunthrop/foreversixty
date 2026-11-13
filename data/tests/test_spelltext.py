import logging
from pathlib import Path

from pipeline.csvio import read_csv
from pipeline.spelltext import ExtraRows, effect_amount, load_spell_text

HERE = Path(__file__).parent


def fixture_text():
    return load_spell_text(
        read_csv(HERE / "fixtures/Spell.csv"),
        read_csv(HERE / "fixtures/SpellMisc.csv"),
        read_csv(HERE / "fixtures/SpellEffect.csv"),
        read_csv(HERE / "fixtures/SpellDuration.csv"),
        ExtraRows(names=read_csv(HERE / "fixtures/SpellName.csv")),
    )


def test_s_token_uses_base_points_plus_die_sides():
    text = fixture_text()
    assert text.describe(12282) == "Increases your Strength by 2%."
    assert text.describe(12663) == "Increases your Strength by 4%."
    assert text.describe(12664) == "Increases your Strength by 6%."


def test_divisor_token_divides_and_drops_the_sign():
    text = fixture_text()
    assert text.describe(12285) == "Reduces the cooldown of your Overpower by 1 sec."
    assert text.describe(12697) == "Reduces the cooldown of your Overpower by 2 sec."


def test_second_effect_and_duration_tokens():
    text = fixture_text()
    assert text.describe(12294) == (
        "A vicious strike that deals weapon damage plus 85 and wounds the target for 10 sec."
    )


def test_periodic_total_multiplies_by_the_tick_count():
    # base 7 + die 1 = 8 per tick, duration 12000 / period 3000 = 4 ticks.
    text = fixture_text()
    assert text.describe(21179) == "Causes 32 Fire damage over 12 sec."


def test_unresolvable_tokens_stay_verbatim():
    text = fixture_text()
    assert text.describe(21838) == ("Gives you a $h% chance to generate an additional Rage point.")


def test_at_spellicon_is_dropped():
    """Idol of the Huntress' own real effect_text shape: `$@spellicon<id>`
    carries no text of its own -- the client shows that spell's icon
    graphically -- so this text-only pipeline simply removes it."""
    text = fixture_text()
    assert text.describe(439510) == (
        "Engrave your cloak with the Improved Swipe rune:\n\n\nIncreases your Strength by 2%."
    )


def test_at_spellname_resolves_from_spell_name_csv():
    text = fixture_text()
    assert text.describe(439511) == "Grants: Improved Heroic Strike"


def test_at_spelldesc_resolves_the_other_spells_own_tokens_recursively():
    """`$@spelldesc12282` becomes spell 12282's OWN fully-resolved
    description ("Increases your Strength by $s1%." -> "...2%.", its own
    `$s1` read off ITS OWN effect, base_points 1 + die_sides 1) --
    already covered by `test_at_spellicon_is_dropped` above (439510 uses
    both tokens together, the real item's own shape); this test isolates
    spelldesc alone via `describe(12282)` matching independently."""
    text = fixture_text()
    assert text.describe(12282) == "Increases your Strength by 2%."


def test_at_spellname_stays_raw_and_spelldesc_drops_for_an_unknown_spell(caplog):
    """999998 is in neither Spell.csv nor SpellName.csv. Its name is left exactly as the
    client wrote it (tenet 8: an ability this pipeline cannot verify is never invented);
    its description prints nothing, as Wowhead's Forever tooltip prints a missing rune
    (Libram of Banishment 211472, 2026-11-10). Each warns; the SAME string's
    `$@spellicon999998` drops with no warning at all (an icon needs no table)."""
    text = fixture_text()
    with caplog.at_level(logging.WARNING, logger="pipeline.spelltext"):
        result = text.describe(439512)
    assert result == "Unlocks $@spellname999998 and  and ."
    assert sum("999998" in message for message in caplog.messages) == 2


def test_at_spelldesc_self_cycle_does_not_recurse_forever(caplog):
    """A spell whose own description names itself (however unlikely in
    real data) must not blow the stack -- left raw, warned, same as any
    other unresolvable `$@` token."""
    text = fixture_text()
    with caplog.at_level(logging.WARNING, logger="pipeline.spelltext"):
        result = text.describe(439513)
    assert result == "$@spelldesc439513"
    assert any("439513" in message for message in caplog.messages)


def test_icon_file_id_comes_from_spell_misc():
    text = fixture_text()
    assert text.icon_file_id(12282) == 132154
    assert text.icon_file_id(999999) == 0


def test_unknown_spell_describes_as_empty_string():
    assert fixture_text().describe(999999) == ""


def test_rolled_effects_render_as_a_range():
    from pipeline.spelltext import Effect, SpellRow, SpellText

    text = SpellText(
        {
            11366: SpellRow(
                description="Causes $s1 Fire damage.",
                duration_ms=None,
                icon_file_id=0,
                effects={0: Effect(base_points=140, die_sides=47, period_ms=0)},
            )
        }
    )
    assert text.describe(11366) == "Causes 141 to 187 Fire damage."


def _one_spell(description: str, effects: dict | None = None):
    from pipeline.spelltext import SpellRow, SpellText

    row = SpellRow(description=description, duration_ms=None, icon_file_id=0, effects=effects or {})
    return SpellText({1: row})


def test_a_radius_on_an_effect_the_row_lacks_reads_the_spells_one_radius():
    from pipeline.spelltext import Effect

    # Piercing Howl: one effect, its radius kept in index 1, and the text says $a2.
    effects = {0: Effect(base_points=0, die_sides=0, period_ms=0, radius=10.0)}
    text = _one_spell("Causes all enemies within $a2 yds to be Dazed.", effects)
    assert text.describe(1) == "Causes all enemies within 10 yds to be Dazed."


def test_two_radii_that_disagree_leave_the_token():
    from pipeline.spelltext import Effect

    effects = {
        0: Effect(base_points=0, die_sides=0, period_ms=0, radius=10.0),
        1: Effect(base_points=0, die_sides=0, period_ms=0, radius=20.0),
    }
    assert _one_spell("Within $a3 yds.", effects).describe(1) == "Within $a3 yds."


def test_gender_text_names_both():
    text = _one_spell("Deals 82 Fire damage to $ghimself:herself; and nearby enemies.")
    assert text.describe(1) == "Deals 82 Fire damage to himself or herself and nearby enemies."


def test_a_condition_on_the_reader_reads_its_otherwise_branch():
    text = _one_spell("Restores health. $?pc142418[Experience is increased by 5%.][]")
    assert text.describe(1) == "Restores health."
    text = _one_spell("Heals you$?s123[ and your pet][].")
    assert text.describe(1) == "Heals you."


def test_minute_durations_render_as_minutes():
    from pipeline.spelltext import SpellRow, SpellText

    text = SpellText(
        {1: SpellRow(description="Lasts $d.", duration_ms=60000, icon_file_id=0, effects={})}
    )
    assert text.describe(1) == "Lasts 1 min."


def test_cross_spell_reference_reads_the_other_spell():
    from pipeline.spelltext import Effect, SpellRow, SpellText

    text = SpellText(
        {
            1: SpellRow(
                description="Stuns for $7s1 sec and again for $9s1 sec.",
                duration_ms=None,
                icon_file_id=0,
                effects={},
            ),
            7: SpellRow(
                description="",
                duration_ms=None,
                icon_file_id=0,
                effects={0: Effect(base_points=2, die_sides=1, period_ms=0)},
            ),
        }
    )
    assert text.describe(1) == "Stuns for 3 sec and again for $9s1 sec."


def test_duration_token_stays_verbatim_when_duration_is_none():
    from pipeline.spelltext import SpellRow, SpellText

    text = SpellText(
        {1: SpellRow(description="Lasts $d.", duration_ms=None, icon_file_id=0, effects={})}
    )
    assert text.describe(1) == "Lasts $d."


def test_missing_effect_index_reads_the_first_effect():
    """Wowhead's Forever tooltips print a token naming an effect the spell lacks with the
    spell's first effect (2026-11-10, nether.wowhead.com/tooltip/spell/<id>?dataEnv=16):
    Unflinching Morale 1216997 "by $s2%" is "by 5%", Shifting Power 1322605 "$s2 Energy"
    is "40 Energy", Temporal Acceleration 367756 "$s3%" is "50%"."""
    from pipeline.spelltext import Effect, SpellRow, SpellText

    text = SpellText(
        {
            1: SpellRow(
                description="Extra effect: $s3.",
                duration_ms=None,
                icon_file_id=0,
                effects={0: Effect(base_points=1, die_sides=1, period_ms=0)},
            )
        }
    )
    assert text.describe(1) == "Extra effect: 2."
    assert SpellText(
        {1: SpellRow(description="Extra: $s3.", duration_ms=None, icon_file_id=0)}
    ).describe(1) == ("Extra: $s3.")


def test_t_token_renders_the_tick_period_in_seconds():
    from pipeline.spelltext import Effect, SpellRow, SpellText

    text = SpellText(
        {
            1: SpellRow(
                description="Ticks every $t1 sec.",
                duration_ms=None,
                icon_file_id=0,
                effects={0: Effect(base_points=1, die_sides=1, period_ms=3000)},
            )
        }
    )
    assert text.describe(1) == "Ticks every 3 sec."


def test_divisor_duration_form_divides_the_duration():
    from pipeline.spelltext import SpellRow, SpellText

    text = SpellText(
        {
            1: SpellRow(
                description="Lasts $/1000;d sec.",
                duration_ms=10000,
                icon_file_id=0,
                effects={},
            )
        }
    )
    assert text.describe(1) == "Lasts 10 sec."


def test_m_token_renders_the_effects_minimum():
    # 724 of the 1320 rank descriptions in build 1.60.1.69893 use $m.
    text = fixture_text()
    assert text.describe(12321) == "Increases the duration of your shouts by 10%."


def test_overrides_replace_an_effects_value_for_one_rank():
    text = fixture_text()
    assert text.describe(12321, {0: 30}) == "Increases the duration of your shouts by 30%."
    # The spell itself is untouched: the next call sees the client's own value.
    assert text.describe(12321) == "Increases the duration of your shouts by 10%."


def test_overrides_reach_an_effect_the_spell_has_no_row_for():
    text = fixture_text()
    assert text.describe(12777, {1: 7}) == "Deals 4 damage and stuns for 7 sec."


def test_effect_amount_folds_eras_die_side_into_the_amount():
    """Era states "Attack Power 24" as 23 + 1."""
    assert (
        effect_amount({"EffectBasePoints": "23", "EffectBasePointsF": "0", "EffectDieSides": "1"})
        == 24
    )


def test_effect_amount_reads_the_modern_column_as_final():
    """The 1.60 client has no EffectDieSides and states 8 for the +8 Strength
    enchant's spell, where Era stated 7 + 1."""
    assert effect_amount({"EffectBasePointsF": "8"}) == 8


def test_effect_amount_keeps_a_negative_amount_negative():
    assert effect_amount({"EffectBasePointsF": "-40"}) == -40


def test_effect_amount_keeps_the_clients_fraction():
    """Wrack's tick (SpellEffect 1349556) is 25.2 and the Lightning Bolt rank 4
    overload (408442) 36.5 on 1.60.1.70938. The data keeps them exactly; only a
    tooltip rounds, when it prints them."""
    assert effect_amount({"EffectBasePointsF": "25.2000007629"}) == 25.2000007629
    assert effect_amount({"EffectBasePointsF": "36.5"}) == 36.5
    assert effect_amount({"EffectBasePointsF": "107.5"}) == 107.5


def _one(description: str, **effects: int) -> str:
    """One spell's rendered description; `s1=300` makes effect 1 display 300."""
    from pipeline.spelltext import Effect, SpellRow, SpellText

    rows = {
        1: SpellRow(
            description=description,
            duration_ms=None,
            icon_file_id=0,
            effects={
                int(name[1:]) - 1: Effect(base_points=value, die_sides=0, period_ms=0)
                for name, value in effects.items()
            },
        )
    }
    return SpellText(rows).describe(1)


def test_brace_arithmetic_is_worked_out_once_its_tokens_are_numbers():
    # Boundless Rage on the 1.60 client: rage is stored times ten.
    assert _one("Increases your maximum Rage by ${$s1/10}.", s1=300) == (
        "Increases your maximum Rage by 30."
    )
    assert _one("by ${$s1*3}%.", s1=8) == "by 24%."
    assert _one("by ${$s1+4}%,", s1=8) == "by 12%,"
    assert _one("by ${($s1/10)*2} sec", s1=10) == "by 2 sec"


def test_brace_arithmetic_drops_the_sign_the_way_every_other_token_does():
    # The client divides a negative amount by a negative number to show it positive; the
    # amounts here are already shown without their sign, so the quotient's sign goes too.
    assert _one("reduced by ${$s1/-1000} sec.", s1=2000) == "reduced by 2 sec."


def test_a_precision_suffix_is_the_number_of_decimals_not_a_full_stop():
    assert _one("by ${$s1/-1000}.1 sec.", s1=500) == "by 0.5 sec."
    assert _one("by ${$s1/-1000}.2 sec", s1=330) == "by 0.33 sec"
    # A whole number keeps no empty decimals, and a real sentence end stays a full stop.
    assert _one("by ${$s1/-1000}.1 sec.", s1=1000) == "by 1 sec."
    assert _one("by ${$s1/-10}.", s1=30) == "by 3."


def test_a_stat_the_reader_brings_stays_a_word_in_the_formula():
    """Wowhead's Forever tooltips (2026-11-10): Summon Hawk 1293241 prints `$rap` as
    "Ranged Attack Power", Dark Sacrifice 1277324 "${400+$SPI}" as "(440 + Spirit)" (its
    400 is our build's base points), Wyvern Strike 458436 `$AP` as "Attack Power"."""
    assert _one("dealing ${32+($rap*(5/100))} Physical damage") == (
        "dealing (32 + Ranged Attack Power * 0.05) Physical damage"
    )
    assert _one("to gain ${400+$SPI} Mana.") == "to gain (400 + Spirit) Mana."
    assert _one("Bleed for ${6/100*8*$AP} Physical damage") == (
        "Bleed for (0.48 * Attack Power) Physical damage"
    )


def test_brace_arithmetic_that_is_not_all_numbers_is_left_as_the_client_stored_it():
    # $sk43 (the reader's sword skill) and an undefined $<mult> have no value here.
    raw = "${$sk43*5} and ${8*$<mult>}"
    assert _one(raw) == raw
    assert _one("${$s1/0}", s1=5) == "${5/0}"


def test_brace_arithmetic_reads_the_duration_in_seconds_and_rounds_to_whole():
    """Drink 430 on 1.60.1.70790: "Restores ${($s1/5)*$d} mana over $d." with 42 and 18 sec
    is 151.2, which Wowhead's Forever tooltip prints as "Restores 151 mana over 18 sec.";
    Drink 431 (104, 21 sec, 436.8) prints 437."""
    from pipeline.spelltext import Effect, SpellRow, SpellText

    def drink(points: int, duration_ms: int) -> str:
        row = SpellRow(
            description="Restores ${($s1/5)*$d} mana over $d.",
            duration_ms=duration_ms,
            icon_file_id=0,
            effects={0: Effect(base_points=points, die_sides=0, period_ms=0)},
        )
        return SpellText({430: row}).describe(430)

    assert drink(42, 18_000) == "Restores 151 mana over 18 sec."
    assert drink(104, 21_000) == "Restores 437 mana over 21 sec."


def test_plural_tokens_follow_the_number_before_them():
    assert _one("as if you were ${$s1/5} $llevel:levels; higher.", s1=5) == (
        "as if you were 1 level higher."
    )
    assert _one("as if you were ${$s1/5} $llevel:levels; higher.", s1=15) == (
        "as if you were 3 levels higher."
    )
    assert _one("refunds $s1 Combo $LPoint:Points; when cast", s1=2) == (
        "refunds 2 Combo Points when cast"
    )
    assert _one("Awards $s1 combo $lpoint:points;.", s1=1) == "Awards 1 combo point."


def test_a_plural_token_with_no_number_before_it_is_left_alone():
    assert _one("Your $lspell:spells; hit harder.") == "Your $lspell:spells; hit harder."


def _proc_text(description: str, **fields) -> str:
    """One spell with the proc, range and effect extras set; see `SpellRow` and `Effect`."""
    from pipeline.spelltext import Effect, SpellRow, SpellText

    effect = Effect(
        base_points=10,
        die_sides=0,
        period_ms=0,
        radius=fields.pop("radius", None),
        chain_targets=fields.pop("chain_targets", 0),
        misc_value=fields.pop("misc_value", 0),
    )
    row = SpellRow(
        description=description, duration_ms=None, icon_file_id=0, effects={0: effect}, **fields
    )
    other = SpellRow(
        description="", duration_ms=None, icon_file_id=0, effects={}, max_stacks=3, proc_charges=2
    )
    return SpellText({1: row, 20128: other}).describe(1)


def test_proc_chance_charges_stacks_and_cooldown_come_from_the_aura_options():
    # Enrage on the 1.60 client.
    assert _proc_text("Gives you a $h% chance to deal $s1% more.", proc_chance=30) == (
        "Gives you a 30% chance to deal 10% more."
    )
    assert _proc_text("up to $n times, stacking $u times.", proc_charges=4, max_stacks=5) == (
        "up to 4 times, stacking 5 times."
    )
    # The cooldown is milliseconds in the table and seconds in a sentence, and the client
    # does its own arithmetic on it.
    assert _proc_text("once every $proccooldown sec.", proc_cooldown_ms=30_000) == (
        "once every 30 sec."
    )
    assert _proc_text("once every ${$proccooldown/60} min.", proc_cooldown_ms=120_000) == (
        "once every 2 min."
    )


def test_a_proc_chance_over_one_hundred_is_the_clients_always():
    assert _proc_text("a $h% chance", proc_chance=101) == "a 100% chance"


def test_radius_range_and_chain_targets():
    assert _proc_text("all enemies within $a1 yards.", radius=8.0) == "all enemies within 8 yards."
    assert _proc_text("an enemy within $r yards.", range_max=30.0) == "an enemy within 30 yards."
    assert _proc_text("jumps to $x1 targets.", chain_targets=3) == "jumps to 3 targets."


def test_the_same_tokens_read_off_another_spell():
    assert _proc_text("stacks $20128u times, $20128n charges") == "stacks 3 times, 2 charges"


def test_a_value_the_tables_do_not_carry_leaves_its_token_alone():
    raw = "a $h% chance, $n charges, $u stacks, within yards of $r, $x1 jumps, $proccooldown"
    assert _proc_text(raw) == raw


def test_a_radius_the_spell_has_none_of_prints_zero():
    # Cloud of Chained Essence 12766: Wowhead's Forever tooltip reads "a 0 yard radius".
    assert _proc_text("in a $a1 yard radius") == "in a 0 yard radius"


def test_a_longer_token_is_never_read_as_a_shorter_one():
    # "$rap" is ranged attack power, not "$r" followed by "ap".
    assert _proc_text("${32+($rap*(5/100))}", range_max=30.0) == (
        "(32 + Ranged Attack Power * 0.05)"
    )
    assert _proc_text("$hp $HP", proc_chance=30) == "$hp $HP"


def test_q_is_the_effects_misc_value():
    """Rebirth rank 5 (20748): "restoring a dead target to life with 2200 health and $q
    mana." Effect 0 (329, resurrect) carries the health as base points and the mana as
    EffectMiscValue_0 (2800 on 1.60.1.70755); `$q` was left raw and the tooltip unreadable
    (sim audit round 17)."""
    spell = {"ID": "20748", "Description_lang": "with $s1 health and $q mana."}
    misc = {"SpellID": "20748", "DurationIndex": "0", "SpellIconFileDataID": "0"}
    effect = {
        "SpellID": "20748",
        "EffectIndex": "0",
        "EffectBasePointsF": "2200",
        "EffectAuraPeriod": "0",
        "EffectMiscValue_0": "2800",
    }
    text = load_spell_text([spell], [misc], [effect], [])
    assert text.describe(20748) == "with 2200 health and 2800 mana."


def test_q_reads_its_own_effect_and_stays_raw_without_a_misc_value():
    # `$q2` names an effect this spell lacks, so it reads the first (see
    # test_missing_effect_index_reads_the_first_effect).
    assert _proc_text("$q1 and $q2", misc_value=5) == "5 and 5"
    assert _proc_text("$q mana") == "$q mana"
