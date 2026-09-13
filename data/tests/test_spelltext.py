from pathlib import Path

from pipeline.csvio import read_csv
from pipeline.spelltext import load_spell_text

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
