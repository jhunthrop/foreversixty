from pathlib import Path

from pipeline.csvio import read_csv
from pipeline.spelltext import effect_amount, load_spell_text

HERE = Path(__file__).parent


def fixture_text():
    return load_spell_text(
        read_csv(HERE / "fixtures/Spell.csv"),
        read_csv(HERE / "fixtures/SpellMisc.csv"),
        read_csv(HERE / "fixtures/SpellEffect.csv"),
        read_csv(HERE / "fixtures/SpellDuration.csv"),
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
    assert text.describe(21838) == (
        "Gives you a $h% chance to generate an additional Rage point."
    )


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


def test_missing_effect_index_stays_verbatim():
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
    assert text.describe(1) == "Extra effect: $s3."


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
    assert effect_amount(
        {"EffectBasePoints": "23", "EffectBasePointsF": "0", "EffectDieSides": "1"}
    ) == 24


def test_effect_amount_reads_the_modern_column_as_final():
    """The 1.60 client has no EffectDieSides and states 8 for the +8 Strength
    enchant's spell, where Era stated 7 + 1."""
    assert effect_amount({"EffectBasePointsF": "8"}) == 8


def test_effect_amount_keeps_a_negative_amount_negative():
    assert effect_amount({"EffectBasePointsF": "-40"}) == -40


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


def test_brace_arithmetic_that_is_not_all_numbers_is_left_as_the_client_stored_it():
    # $rap and $<mult> have no value in the tables this module reads.
    raw = "${32+($rap*(5/100))} damage and ${(172+($bh*$bc))*$<mult>}"
    assert _one(raw) == raw
    assert _one("${$s1/0}", s1=5) == "${5/0}"


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
