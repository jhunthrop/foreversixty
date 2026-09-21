"""Turn a WoW spell description into readable text.

The client stores descriptions with `$`-tokens that the game resolves at
runtime. We resolve only the tokens whose value is in the tables we fetch:

  $s<n> / $S<n>        the displayed value of effect <n> (1-based)
  $o<n> / $O<n>        that value times the number of periodic ticks
  $t<n> / $T<n>        effect <n>'s tick period, in seconds
  $d   / $D            the spell's duration
  $m<n> / $M<n>        the minimum of effect <n>'s displayed value
  $/<divisor>;s<n>     any of the above, divided first
  $<spell id><token>   the same token read off a different spell
  ${<arithmetic>}      worked out, once every token inside it is a number;
                       a trailing `.N` is the client's "N decimals"
  $l<one>:<many>;      the word that agrees with the number before it

Everything else stays in the string exactly as the client stored it:
`$h` (proc chance), `$a<n>` (radius), `$?x[..][..]` (conditionals), and any
`${..}` that still holds a token with no value here (`$rap`, `$<mult>`). The
planner shows the raw token rather than a number nobody can source.
"""

from __future__ import annotations

import ast
import operator
import re
from collections.abc import Callable, Mapping
from dataclasses import dataclass, field, replace

from pipeline.csvio import populated


class SpellTextError(ValueError):
    """A spell effect row states something this module will not guess at."""


TOKEN = re.compile(
    r"\$"
    r"(?:/(?P<divisor>\d+);)?"
    r"(?P<ref>\d+)?"
    r"(?P<kind>[sSoOtTdDmM])"
    r"(?P<index>[1-9])?"
)

#: `${300/10}`, `${500/-1000}.1`: the braces, then the client's optional decimals suffix. The
#: suffix is a dot and ONE digit not followed by another, so "by ${30/-10}." keeps its full
#: stop and "${8*2.5}%" is untouched.
BRACES = re.compile(r"\$\{(?P<expr>[^{}]*)\}(?:\.(?P<places>\d)(?!\d))?")
_ARITHMETIC = re.compile(r"[\d\s+\-*/().]+")
#: `$lpoint:points;` -- and the number it has to agree with, somewhere before it.
PLURAL = re.compile(
    r"(?P<number>\d+(?:\.\d+)?)(?P<between>[^$\d]*)"
    r"\$[lL](?P<one>[^:;]+):(?P<many>[^;]+);"
)

_BINARY: dict[type[ast.operator], Callable[[float, float], float]] = {
    ast.Add: operator.add,
    ast.Sub: operator.sub,
    ast.Mult: operator.mul,
    ast.Div: operator.truediv,
}

_MS_PER_SECOND = 1000
_MS_PER_MINUTE = 60_000


@dataclass(frozen=True)
class Effect:
    base_points: int
    die_sides: int
    period_ms: int


@dataclass(frozen=True)
class SpellRow:
    description: str
    duration_ms: int | None
    icon_file_id: int
    effects: dict[int, Effect] = field(default_factory=dict)


def _format_number(value: float) -> str:
    if value == int(value):
        return str(int(value))
    return f"{value:.3f}".rstrip("0").rstrip(".")


def _evaluate(node: ast.expr) -> float:
    """Four-function arithmetic over number literals, and nothing else."""
    if isinstance(node, ast.Constant) and isinstance(node.value, (int, float)):
        return float(node.value)
    if isinstance(node, ast.UnaryOp) and isinstance(node.op, (ast.USub, ast.UAdd)):
        value = _evaluate(node.operand)
        return -value if isinstance(node.op, ast.USub) else value
    if isinstance(node, ast.BinOp) and type(node.op) in _BINARY:
        return _BINARY[type(node.op)](_evaluate(node.left), _evaluate(node.right))
    raise ValueError(f"not arithmetic: {ast.dump(node)}")


def _work_out(match: re.Match[str]) -> str:
    """One `${..}` as its value, or exactly as it stood when it is not all numbers.

    The sign is dropped, as it is for every other token here: the client divides a
    negative amount by a negative number to show it positive, and the amounts that reach
    this point have already lost theirs.
    """
    expr = match["expr"]
    if not _ARITHMETIC.fullmatch(expr):
        return match[0]
    try:
        value = abs(_evaluate(ast.parse(expr.strip(), mode="eval").body))
    except (SyntaxError, ValueError, ZeroDivisionError):
        return match[0]
    if match["places"] is None:
        return _format_number(value)
    return _format_number(round(value, int(match["places"])))


def _agree(match: re.Match[str]) -> str:
    word = match["one"] if float(match["number"]) == 1 else match["many"]
    return f"{match['number']}{match['between']}{word}"


def _bounds(effect: Effect) -> tuple[int, int]:
    """The low and high displayed values for one effect."""
    if effect.die_sides > 1:
        return effect.base_points + 1, effect.base_points + effect.die_sides
    single = effect.base_points + effect.die_sides
    return single, single


def _render(low: float, high: float, divisor: int) -> str:
    first, second = sorted((abs(low) / divisor, abs(high) / divisor))
    if first == second:
        return _format_number(first)
    return f"{_format_number(first)} to {_format_number(second)}"


def _render_duration(duration_ms: int) -> str:
    if duration_ms >= _MS_PER_MINUTE and duration_ms % _MS_PER_MINUTE == 0:
        return f"{duration_ms // _MS_PER_MINUTE} min"
    return f"{_format_number(duration_ms / _MS_PER_SECOND)} sec"


def _overridden(effects: dict[int, Effect], overrides: Mapping[int, int]) -> dict[int, Effect]:
    """The overridden effects, keeping each one's tick period.

    Die sides go to zero: a curve names one value per rank, not a spread, so
    `$s` and `$m` must both render exactly that number.
    """
    out: dict[int, Effect] = {}
    for index, value in overrides.items():
        period = effects[index].period_ms if index in effects else 0
        out[index] = Effect(base_points=value, die_sides=0, period_ms=period)
    return out


class SpellText:
    def __init__(self, spells: dict[int, SpellRow]) -> None:
        self._spells = spells

    def describe(self, spell_id: int, overrides: Mapping[int, int] | None = None) -> str:
        """The spell's description with its `$`-tokens resolved.

        `overrides` maps a 0-based effect index to the value that effect
        displays, which is how one rank of a trait talent gets its own
        sentence: the 1.60 client stores one description per talent and puts
        the per-rank numbers on a curve (see `pipeline/curves.py`). Only the
        described spell's own effects are overridden -- a `$<id>s1` token
        still reads the referenced spell as the client stores it.
        """
        row = self._spells.get(spell_id)
        if row is None:
            return ""
        if overrides:
            row = replace(row, effects={**row.effects, **_overridden(row.effects, overrides)})
        resolved = TOKEN.sub(lambda match: self._substitute(row, match), row.description)
        return PLURAL.sub(_agree, BRACES.sub(_work_out, resolved))

    def icon_file_id(self, spell_id: int) -> int:
        row = self._spells.get(spell_id)
        return row.icon_file_id if row else 0

    def _substitute(self, row: SpellRow, match: re.Match[str]) -> str:
        target = row
        if match["ref"]:
            referenced = self._spells.get(int(match["ref"]))
            if referenced is None:
                return match[0]
            target = referenced
        divisor = int(match["divisor"]) if match["divisor"] else 1
        kind = match["kind"].lower()
        if kind == "d":
            if target.duration_ms is None:
                return match[0]
            if match["divisor"]:
                return _render(target.duration_ms, target.duration_ms, divisor)
            return _render_duration(target.duration_ms)
        effect = target.effects.get(int(match["index"] or 1) - 1)
        if effect is None:
            return match[0]
        if kind == "t":
            if effect.period_ms <= 0:
                return match[0]
            return _render(effect.period_ms, effect.period_ms, divisor * _MS_PER_SECOND)
        low, high = _bounds(effect)
        if kind == "o":
            if target.duration_ms is None or effect.period_ms <= 0:
                return match[0]
            ticks = target.duration_ms // effect.period_ms
            low, high = low * ticks, high * ticks
        if kind == "m":
            high = low
        return _render(low, high, divisor)


def _base_points(row: dict[str, str]) -> int:
    """The effect's base points: Classic Era exports the integer column, the 1.60 client
    (Forever beta) the modern float column; either way the formulas want a whole number.
    The 1.60 client also has no EffectDieSides (its spread is the Variance column, which
    the descriptions do not use yet), so a missing die-sides column reads as no spread.

    Both builds' SpellEffect rows carry both column headers, but only one is ever
    populated: Era's own EffectBasePoints is always populated and its EffectBasePointsF
    is always "0" (unused padding), while build 1.60.1.69893's EffectBasePoints is
    always empty and EffectBasePointsF carries the real value -- confirmed against
    every row of both builds' own SpellEffect tables (see data/README.md). Reading
    the float column whenever it merely exists, as an earlier version of this
    function did, silently zeroed every Era description's numbers (see
    test_effect_base_points_read_the_float_column_when_the_build_has_it, which this
    still satisfies: the integer column is genuinely absent from that test's beta
    fixture, not merely empty). Preferring the int column when both are populated is
    the same rule `pipeline/normalize/item_curves.py`'s `_budget` uses for
    RandPropPoints; if the two ever populated the same row with different numbers,
    that would mean the "only one is ever real" premise this whole function rests on
    is false, so this raises rather than silently pick a side.

    Era's EffectBasePointsF padding is the literal string "0", which `populated`
    (correctly) still reports as present -- a column being non-empty is not the same
    as it carrying real data. Comparing that padding against a real, nonzero
    EffectBasePoints made this raise on the vast majority of Era's own SpellEffect
    rows rather than only on a genuine disagreement, so a literal "0" float column is
    treated as Era's padding (not populated, for comparison purposes) whenever the int
    column is also populated. Any other float value is still compared and still raises
    on a mismatch -- see test_base_points_columns_that_disagree_are_an_error_not_a_guess.
    """
    raw, float_raw = populated(row, "EffectBasePoints"), populated(row, "EffectBasePointsF")
    if raw is not None and float_raw is not None and float_raw != "0":
        int_value, float_value = int(raw), round(float(float_raw))
        if int_value != float_value:
            raise SpellTextError(
                f"spell {row.get('SpellID')} effect {row.get('EffectIndex')} has both "
                f"EffectBasePoints ({int_value}) and EffectBasePointsF ({float_value}) "
                "populated and disagreeing"
            )
        return int_value
    if raw is not None:
        return int(raw)
    return round(float(row["EffectBasePointsF"]))


def effect_amount(row: dict[str, str]) -> int:
    """One SpellEffect row's final amount, on either client schema.

    Classic Era states an amount one short and puts the last point in
    EffectDieSides -- "Increase Spell Dam 85" is 84 + 1, "Attack Power 24" is
    23 + 1. The 1.60 client (Forever beta) has no EffectDieSides column at all
    and its EffectBasePointsF is already final: spell 9103, the "+8 Strength"
    enchant's spell, reads 8 where Era read 7. Adding a missing die-sides
    column as zero makes one expression right on both.

    Public so `pipeline/simdb/equip.py` reads amounts the same way the
    description renderer does; the per-build column choice lives in
    `_base_points` and nowhere else.
    """
    return _base_points(row) + int(row.get("EffectDieSides") or 0)


def load_spell_text(
    spell_rows: list[dict[str, str]],
    misc_rows: list[dict[str, str]],
    effect_rows: list[dict[str, str]],
    duration_rows: list[dict[str, str]],
) -> SpellText:
    durations = {int(r["ID"]): int(r["Duration"]) for r in duration_rows}
    effects: dict[int, dict[int, Effect]] = {}
    for row in effect_rows:
        if row.get("DifficultyID", "0") != "0":
            continue
        effects.setdefault(int(row["SpellID"]), {})[int(row["EffectIndex"])] = Effect(
            base_points=_base_points(row),
            die_sides=int(row.get("EffectDieSides") or 0),
            period_ms=int(row["EffectAuraPeriod"]),
        )
    misc: dict[int, tuple[int | None, int]] = {}
    for row in misc_rows:
        if row.get("DifficultyID", "0") != "0":
            continue
        duration = durations.get(int(row["DurationIndex"]))
        if duration is not None and duration < 0:
            duration = None
        misc[int(row["SpellID"])] = (duration, int(row["SpellIconFileDataID"]))
    spells: dict[int, SpellRow] = {}
    for row in spell_rows:
        spell_id = int(row["ID"])
        duration, icon_file_id = misc.get(spell_id, (None, 0))
        spells[spell_id] = SpellRow(
            description=row["Description_lang"],
            duration_ms=duration,
            icon_file_id=icon_file_id,
            effects=effects.get(spell_id, {}),
        )
    return SpellText(spells)
