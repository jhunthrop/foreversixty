"""Turn a WoW spell description into readable text.

The client stores descriptions with `$`-tokens that the game resolves at
runtime. We resolve the tokens whose value is in the tables we fetch:

  $s<n> / $S<n>        the displayed value of effect <n> (1-based)
  $m<n> / $M<n>        the minimum / maximum of effect <n>'s displayed value
  $o<n> / $O<n>        that value times the number of periodic ticks
  $t<n> / $T<n>        effect <n>'s tick period, in seconds
  $d   / $D            the spell's duration
  $a<n>                effect <n>'s radius, in yards
  $x<n>                how many targets effect <n> chains to
  $q<n>                effect <n>'s misc value (Rebirth's mana: EffectMiscValue_0)
  $b<n>                effect <n>'s points per combo point (Rip 1079: 5.5)
  $bc<n>               effect <n>'s bonus coefficient
  $e<n>                effect <n>'s amplitude (Mana Shield 1463: 2 mana per damage)
  $f<n>                effect <n>'s chain amplitude (Thornspike 1293183: 0.5 per jump)
  $r                   the spell's maximum range, in yards
  $h                   the chance to proc, as a percentage
  $n / $u              the proc's charges / how high the aura stacks
  $i                   how many targets the spell hits at most
  $v                   the highest target level it affects (Mind Soothe 453: 40)
  $ec<n> / $ecim       the enchantment the spell applies: its n-th amount and the lowest
                       item level it takes (Heavy Armor Kit 2833: 24 armor, items level
                       15; Classic Wowhead prints these, the Forever one prints 0)
  $z                   the reader's home, which the client names as Wowhead does
  $proccooldown        the proc's internal cooldown, in seconds
  $/<divisor>;s<n>     any of the above, divided first
  $*<factor>;f<n>      any of the above, multiplied first (Execute 5308: `$*10;F1` with
                       a chain amplitude of 0.3 prints "3 additional damage")
  $<spell id><token>   the same token read off a different spell
  $@spellicon<id>      dropped -- the client shows this spell's own icon
                       graphically; this text-only pipeline has nothing to
                       put in its place
  $@spellname<id>      that spell's own name (SpellName.csv)
  $@spelldesc<id>      that spell's own description, fully resolved --
                       recursively, the same way `describe` resolves any
                       other spell (quest-gates lane, 2026-09-29: an item's
                       relic/rune "engrave" effect text names the ability
                       it grants this way, e.g. Idol of the Huntress's
                       "$@spellicon439510\n$@spelldesc439510")
  $@spelltooltip<id>   that spell's name, then its description on the next line
                       (Feral Charge 1238122 shows Feral Charge (Cat)'s tooltip)

A `$s<n>` naming an effect the spell does not have reads its first effect, as
Wowhead's Forever tooltips print it (2026-11-10: Unflinching Morale 1216997
"$s2%" is 5%, Shifting Power 1322605 "$s2 Energy" is 40, Temporal Acceleration
367756 "$s3%" is 50%, Charge 22120 "plus $s3" is "plus 0").

Conditions, description variables (`$<name>`), `${..}` arithmetic and plurals are
`pipeline.spelltext_grammar`'s. What stays in the string exactly as the client
stored it is a token whose table has no row or a zero for this spell, and any
`${..}` that still holds a token with no value here. The planner shows the raw
token rather than a number nobody can source. A `$@spellname`/`$@spelldesc`
naming a spell this build's own tables have no row for is the one exception
logged rather than silently left raw (`_substitute_at`'s own doc) -- tenet 8's
"nothing we cannot verify is shown as fact" applies to a whole missing ability
the same way it applies to a missing number.
"""

from __future__ import annotations

import logging
import math
import re
from collections.abc import Callable, Mapping
from dataclasses import dataclass, field, replace

from pipeline.csvio import populated
from pipeline.spelltext_grammar import (
    COLOUR_ESCAPE,
    FUNCTIONS,
    MID_SENTENCE_SPACES,
    READER_LEVEL,
    STAT_WORDS,
    VARIABLE,
    agree_plurals,
    drop_unknown_variables,
    flatten_braces,
    format_number,
    matching_brace,
    parse_variables,
    resolve_conditions,
    work_out,
)

logger = logging.getLogger(__name__)


#: Which grammar the shipped descriptions were rendered with. normalize records it in the
#: build's manifest.json under `SPELLTEXT_MANIFEST_KEY`; bump it with any change that moves
#: shipped text, so `tests/test_spelltext_contract.py` holds a build to the grammar it was
#: rendered with and never fails the committed files before the nightly re-renders them.
#: 2: the full grammar of 2026-11-10 (variables, chained conditions, formulas, variance).
SPELLTEXT_RENDERER = 2
SPELLTEXT_MANIFEST_KEY = "spelltext_renderer"


class SpellTextError(ValueError):
    """A spell effect row states something this module will not guess at."""


TOKEN = re.compile(
    r"\$"
    r"(?:/(?P<divisor>\d+);|\*(?P<multiplier>\d+);)?"
    r"(?P<ref>\d+)?"
    r"(?P<kind>proccooldown|ecim|ec|bc|[sSoOtTdDmMaAxXqQrRhHnNuUiIbBeEfFvVzZ])"
    r"(?P<index>[1-9])?"
    # "$rap" is ranged attack power and "$hp" is health: a longer word is never a short
    # token with letters after it.
    r"(?![A-Za-z])"
)

#: The client writes "always" as a chance over one hundred.
_CERTAIN = 100

#: `.N` after a `${..}`: the client's "N decimals". ONE digit not followed by another, so
#: "by ${30/-10}." keeps its full stop and "${8*2.5}%" is untouched.
_PLACES = re.compile(r"\.(?P<places>\d)(?!\d)")
#: A stat the reader brings, or the reader's level, inside a `${..}`.
_READER_TOKEN = re.compile(r"\$(?P<name>[A-Za-z]+)(?![\w])")
#: The reader's level, both spellings the client uses (Divine Aegis 431622 writes `$pl`).
_LEVEL_TOKENS = frozenset({"PL", "pl"})
#: The same, in a sentence ("Level $PL target:").
_LEVEL_IN_TEXT = re.compile(r"\$(?:PL|pl)(?![A-Za-z\d])")
#: The client's own typo `${$m2$*$<frostdamage>}` (Frostbolt 116): a `$` before an operator.
_STRAY_DOLLAR = re.compile(r"\$(?=[\s*+\-/()])")
#: `$@spellicon439510`, `$@spellname439510`, `$@spelldesc439510` -- a
#: whole other spell referenced by id, never one of this spell's own
#: effects (`TOKEN`, above, never matches `@`: it is a different
#: vocabulary, resolved in its own pass -- `SpellText.describe`'s own
#: doc for the ordering this requires).
AT_TOKEN = re.compile(r"\$@(?P<kind>spellicon|spellname|spelldesc|spelltooltip)(?P<id>\d+)")
#: `$ghimself:herself;` -- the client picks by the reader's gender; a page has no reader, so both.
GENDER = re.compile(r"\$[gG](?P<male>[^:;$]+):(?P<female>[^;$]+);")
#: `$z`: the client prints the reader's bind point; Wowhead's Forever tooltip for Astral
#: Recall (556) prints these words.
_HOME = "Hearthstone Location"

_MS_PER_SECOND = 1000
_MS_PER_MINUTE = 60_000
#: A spell that lasts until its reader ends it (SpellDuration -1): Find Treasure 2481's
#: "Lasts $d." prints "Lasts until canceled." on Wowhead's Forever tooltip.
_UNTIL_CANCELLED = "until canceled"
#: `$o` on a periodic heal or mana return the client stores no tick period for. Wowhead's
#: Forever tooltips divide its duration by 5.2 sec, on every food and drink checked on
#: 2026-11-10: Food 433 (17, 18 sec) 59, 434 (58, 21 sec) 234, 435 (115, 24 sec) 531,
#: 1127 (162, 27 sec) 841; Drink 1135 (332, 30 sec) 1915, 10250 (800, 30 sec) 4615.
_UNTIMED_PERIOD_MS = 5200


@dataclass(frozen=True)
class Effect:
    base_points: float
    die_sides: int
    period_ms: int
    #: Yards, from SpellRadius through EffectRadiusIndex_0 (else _1). None: no radius.
    radius: float | None = None
    #: EffectChainTargets. Zero: the effect does not chain.
    chain_targets: int = 0
    #: EffectMiscValue_0, the `$q` token: a resurrect's mana (Rebirth), and so on.
    misc_value: int = 0
    #: EffectPointsPerResource, the `$b` token: what one more combo point adds.
    points_per_resource: float = 0
    #: EffectBonusCoefficient, `$bc`.
    bonus_coefficient: float = 0
    #: EffectAmplitude, `$e`.
    amplitude: float = 0
    #: EffectChainAmplitude, `$f`: what each jump keeps.
    chain_amplitude: float = 0
    #: The 1.60 client's Variance column: the roll spans base * (1 - v/2) to base * (1 + v/2).
    variance: float = 0
    #: EffectRealPointsPerLevel times the levels the reader has over the spell's BaseLevel,
    #: capped at its MaxLevel (`_level_bonus`).
    level_bonus: float = 0


@dataclass(frozen=True)
class SpellRow:
    description: str
    duration_ms: int | None
    icon_file_id: int
    effects: dict[int, Effect] = field(default_factory=dict)
    #: SpellAuraOptions, all zero or None where the spell has no row: the chance to proc (a
    #: percentage), its charges, how high the aura stacks, and its internal cooldown.
    proc_chance: int | None = None
    proc_charges: int = 0
    max_stacks: int = 0
    proc_cooldown_ms: int = 0
    #: Yards, from SpellRange through SpellMisc.RangeIndex.
    range_max: float | None = None
    #: SpellTargetRestrictions.MaxTargets, the `$i` "up to N enemies"; 0 without a cap.
    max_targets: int = 0
    #: SpellTargetRestrictions.MaxTargetLevel, the `$v`; 0 without a cap.
    max_target_level: int = 0
    #: The SpellDuration row says -1: the spell lasts until cancelled (`_UNTIL_CANCELLED`).
    until_cancelled: bool = False
    #: The enchantment an enchant-item effect (53, 54) applies: SpellItemEnchantment's
    #: EffectPointsMin_0..2 (`$ec<n>`) and ItemLevelMin (`$ecim`).
    enchant_points: tuple[int, ...] = ()
    enchant_min_item_level: int = 0
    #: SpellDescriptionVariables through SpellXDescriptionVariables: `$<name>` -> definition.
    variables: Mapping[str, str] = field(default_factory=dict)


def _half_up(value: float) -> float:
    return float(math.floor(value + 0.5))


def _bounds(effect: Effect, whole: bool) -> tuple[float, float]:
    """The low and high displayed values for one effect.

    Era rolls its die (`EffectDieSides`); the 1.60 client spreads the base by its Variance
    and adds the per-level gain, the variance taken on the base before that gain. Checked
    against Wowhead's Forever tooltips on 2026-11-10: Frostbolt 116 (19, variance 0.105,
    0.2 a level over levels 4 to 8) reads 18.8 to 20.8; Revenge 6572 (22, 0.154) 20.3077
    to 23.6923; Immolate 348 (8, 0.7 a level over 1 to 5) 10.8; Eviscerate 2098 (3, 1.333)
    1 to 5. A sentence prints whole numbers (`whole`); a formula keeps the fractions.
    """
    if effect.die_sides > 1:
        return effect.base_points + 1, effect.base_points + effect.die_sides
    single = effect.base_points + effect.die_sides + effect.level_bonus
    spread = abs(effect.base_points) * effect.variance / 2
    low, high = single - spread, single + spread
    return (_half_up(low), _half_up(high)) if whole else (low, high)


def _render(low: float, high: float, divisor: int) -> str:
    first, second = sorted((abs(low) / divisor, abs(high) / divisor))
    if first == second:
        return format_number(first)
    return f"{format_number(first)} to {format_number(second)}"


def _render_duration(duration_ms: int) -> str:
    if duration_ms >= _MS_PER_MINUTE and duration_ms % _MS_PER_MINUTE == 0:
        return f"{duration_ms // _MS_PER_MINUTE} min"
    return f"{format_number(duration_ms / _MS_PER_SECOND)} sec"


def _overridden(effects: dict[int, Effect], overrides: Mapping[int, float]) -> dict[int, Effect]:
    """The overridden effects, keeping each one's tick period.

    Die sides, variance and the per-level gain go to zero: a curve names one value per
    rank, not a spread, so `$s` and `$m` must both render exactly that number.
    """
    out: dict[int, Effect] = {}
    for index, value in overrides.items():
        known = effects.get(index, Effect(base_points=0, die_sides=0, period_ms=0))
        out[index] = replace(known, base_points=value, die_sides=0, variance=0, level_bonus=0)
    return out


#: Tokens that are one fact about the spell, or about one of its effects. A missing or zero
#: fact returns a falsy value, and the token is then left exactly as the client stored it.
_SPELL_FACTS: dict[str, Callable[[SpellRow], float | None]] = {
    "h": lambda row: None if row.proc_chance is None else min(row.proc_chance, _CERTAIN),
    "n": lambda row: row.proc_charges,
    "u": lambda row: row.max_stacks,
    "i": lambda row: row.max_targets,
    "v": lambda row: row.max_target_level,
    "ecim": lambda row: row.enchant_min_item_level,
    "r": lambda row: row.range_max,
    "proccooldown": lambda row: row.proc_cooldown_ms / _MS_PER_SECOND,
}
_EFFECT_FACTS: dict[str, Callable[[Effect], float | None]] = {
    "x": lambda effect: effect.chain_targets,
    "q": lambda effect: effect.misc_value,
    "b": lambda effect: effect.points_per_resource,
    "bc": lambda effect: effect.bonus_coefficient,
    "e": lambda effect: effect.amplitude,
    "f": lambda effect: effect.chain_amplitude,
}


def _radius_of(row: SpellRow, index: int) -> float | None:
    """`$a<n>`: effect n's radius. A Classic row whose text names an effect it does not have
    (Piercing Howl's `$a2` on a one-effect spell) reads as the spell's one radius when every
    effect that has a radius agrees, which is the number the client prints."""
    effect = row.effects.get(index)
    if effect is not None and effect.radius:
        return effect.radius
    radii = {e.radius for e in row.effects.values() if e.radius}
    return radii.pop() if len(radii) == 1 else None


def _effect_of(row: SpellRow, index: int) -> Effect | None:
    """Effect <index>, or the spell's first effect when it has no such effect (module doc)."""
    return row.effects.get(index) or row.effects.get(0)


@dataclass(frozen=True)
class _Amount:
    """One token's value: a low and high number, or a sentence-ready word."""

    low: float
    high: float
    #: Duration tokens print in their own unit ("18 sec") outside a formula.
    duration_ms: int | None = None


def _amount(
    row: SpellRow, kind: str, upper: bool, index: int, whole: bool = True
) -> _Amount | None:
    """What one token on `row` stands for, before any divisor; None when the tables say nothing."""
    if kind in _SPELL_FACTS:
        value = _SPELL_FACTS[kind](row)
        return _Amount(value, value) if value else None
    if kind == "ec":
        points = row.enchant_points[index] if index < len(row.enchant_points) else 0
        return _Amount(points, points) if points else None
    if kind == "d":
        if row.duration_ms is None:
            return None
        return _Amount(row.duration_ms, row.duration_ms, duration_ms=row.duration_ms)
    if kind == "a":
        radius = _radius_of(row, index)
        if radius is None and not any(e.radius for e in row.effects.values()):
            # Wowhead's Forever tooltip prints a radius the spell has none of as 0: Cloud
            # of Chained Essence 12766 "a 0 yard radius".
            radius = 0.0
        return _Amount(radius, radius) if radius is not None else None
    effect = _effect_of(row, index)
    if effect is None:
        return None
    if kind in _EFFECT_FACTS:
        value = _EFFECT_FACTS[kind](effect)
        return _Amount(value, value) if value else None
    if kind == "t":
        period = effect.period_ms / _MS_PER_SECOND
        return _Amount(period, period) if period > 0 else None
    low, high = _bounds(effect, whole)
    if kind == "o":
        if row.duration_ms is None:
            return None
        if effect.period_ms > 0:
            ticks: float = row.duration_ms // effect.period_ms
        else:
            ticks = row.duration_ms / _UNTIMED_PERIOD_MS
        low, high = low * ticks, high * ticks
        if whole:
            low, high = _half_up(low), _half_up(high)
    if kind == "m":
        low, high = (high, high) if upper else (low, low)
    return _Amount(low, high)


def _scale(match: re.Match[str]) -> tuple[int, int]:
    """A token's `$*N;` multiplier and `$/N;` divisor."""
    multiplier = int(match["multiplier"]) if match["multiplier"] else 1
    return multiplier, int(match["divisor"]) if match["divisor"] else 1


def _scaled(match: re.Match[str]) -> bool:
    return bool(match["multiplier"] or match["divisor"])


def _with_reader_names(expression: str) -> str:
    """The reader's level as `READER_LEVEL`, a stat the reader brings as its bare name."""

    def reader(match: re.Match[str]) -> str:
        name = match["name"]
        if name in _LEVEL_TOKENS:
            return str(READER_LEVEL)
        return name if name in STAT_WORDS or name in FUNCTIONS or name == "cond" else match[0]

    return _READER_TOKEN.sub(reader, expression)


class SpellText:
    def __init__(self, spells: dict[int, SpellRow], names: Mapping[int, str] | None = None) -> None:
        self._spells = spells
        #: SpellName.csv's own id -> Name_lang, for `$@spellname<id>`
        #: only -- `_SPELL_FACTS`/`_EFFECT_FACTS` never need a name, only
        #: a spell's own numbers, so nothing else reads this.
        self._names = names or {}

    def describe(
        self,
        spell_id: int,
        overrides: Mapping[int, float] | None = None,
        *,
        _chain: frozenset[int] = frozenset(),
    ) -> str:
        """The spell's description with its `$`-tokens resolved.

        `overrides` maps a 0-based effect index to the value that effect
        displays, which is how one rank of a trait talent gets its own
        sentence: the 1.60 client stores one description per talent and puts
        the per-rank numbers on a curve (see `pipeline/curves.py`). Only the
        described spell's own effects are overridden -- a `$<id>s1` token
        still reads the referenced spell as the client stores it.

        The passes run in the client's own order of binding: description
        variables are written in, then conditions pick their branch, then
        `${..}` is worked out with its tokens as plain numbers, then the
        remaining tokens print in their sentence form.

        `$@spelldesc<id>` (a WHOLE OTHER spell, not one of this spell's own
        effects) is resolved LAST, after every other token pass -- its own
        replacement text is itself a fully-resolved `describe` call, and
        must never be re-scanned by `TOKEN` in THIS spell's own context (a
        bare `$s1` inside it means the REFERENCED spell's first effect, not
        this one's). `_chain` is this call's own private cycle guard (a
        `$@spelldesc` loop, however unlikely, must not recurse forever) and
        is never passed by a caller outside this module.
        """
        row = self._spells.get(spell_id)
        if row is None:
            return ""
        if overrides:
            row = replace(row, effects={**row.effects, **_overridden(row.effects, overrides)})
        resolved = resolve_conditions(self._write_in_variables(row, row.description))
        resolved = self._work_out_braces(row, resolved)
        resolved = TOKEN.sub(lambda match: self._substitute(row, match), resolved)
        resolved = _LEVEL_IN_TEXT.sub(str(READER_LEVEL), resolved)
        resolved = agree_plurals(resolved)
        resolved = GENDER.sub(lambda m: f"{m['male']} or {m['female']}", resolved)
        resolved = COLOUR_ESCAPE.sub("", drop_unknown_variables(resolved))
        resolved = MID_SENTENCE_SPACES.sub(" ", resolved).rstrip()
        chain = _chain | {spell_id}
        return AT_TOKEN.sub(lambda match: self._substitute_at(match, chain), resolved)

    def icon_file_id(self, spell_id: int) -> int:
        row = self._spells.get(spell_id)
        return row.icon_file_id if row else 0

    def _write_in_variables(
        self, row: SpellRow, text: str, seen: frozenset[str] = frozenset()
    ) -> str:
        """Each `$<name>` as its own definition, worked out first, as a `${..}`."""

        def value(match: re.Match[str]) -> str:
            name = match["name"]
            definition = row.variables.get(name)
            if definition is None or name in seen:
                return match[0]
            written = resolve_conditions(self._write_in_variables(row, definition, seen | {name}))
            written = written.strip()
            if written.startswith("${") and matching_brace(written, 1) == len(written) - 1:
                return written
            return f"${{{written}}}"

        return VARIABLE.sub(value, text)

    def _work_out_braces(self, row: SpellRow, text: str) -> str:
        """Every top-level `${..}`, worked out with its tokens as numbers."""
        out: list[str] = []
        cursor = 0
        while (start := text.find("${", cursor)) >= 0:
            end = matching_brace(text, start + 1)
            if end < 0:
                break
            places = _PLACES.match(text, end + 1)
            expression = self._as_numbers(row, flatten_braces(text[start + 2 : end]))
            worked = work_out(
                _with_reader_names(expression),
                int(places["places"]) if places else None,
                STAT_WORDS,
            )
            out.append(text[cursor:start])
            if worked is None:
                out.append(f"${{{expression}}}")
                cursor = end + 1
            else:
                out.append(worked)
                cursor = places.end() if places else end + 1
        out.append(text[cursor:])
        return "".join(out)

    def _as_numbers(self, row: SpellRow, expression: str) -> str:
        """A formula's tokens as plain numbers (a duration in seconds)."""
        expression = _STRAY_DOLLAR.sub("", expression)
        return TOKEN.sub(lambda match: self._number(row, match), expression)

    def _target(self, row: SpellRow, match: re.Match[str]) -> SpellRow | None:
        if not match["ref"]:
            return row
        return self._spells.get(int(match["ref"]))

    def _number(self, row: SpellRow, match: re.Match[str]) -> str:
        """One token inside a formula: its signed value (a duration in seconds)."""
        target = self._target(row, match)
        kind = match["kind"]
        if target is None or kind.lower() == "z":
            return match[0]
        index = int(match["index"] or 1) - 1
        amount = _amount(target, kind.lower(), kind == "M", index, whole=False)
        if amount is None:
            return match[0]
        # `$s` in a formula is the roll's middle; `$m` and `$M` already chose their end.
        value = (amount.low + amount.high) / 2
        if amount.duration_ms is not None and not _scaled(match):
            value = amount.duration_ms / _MS_PER_SECOND
        multiplier, divisor = _scale(match)
        number = format_number(value * multiplier / divisor)
        return number if value >= 0 else f"({number})"

    def _substitute_at(self, match: re.Match[str], chain: frozenset[int]) -> str:
        """One `$@spellicon`/`$@spellname`/`$@spelldesc`/`$@spelltooltip<id>` token --
        `describe`'s own doc for why this runs after every other
        substitution, never before."""
        kind = match["kind"]
        spell_id = int(match["id"])
        if kind == "spellicon":
            # The client renders this spell's own icon graphically; this
            # text-only description has no image to put in its place, so
            # the token is simply dropped (module doc).
            return ""
        if kind == "spellname":
            name = self._names.get(spell_id)
            if not name:
                logger.warning(
                    "spelltext: $@spellname%d names a spell SpellName.csv has no row "
                    "for; leaving the raw token",
                    spell_id,
                )
                return match[0]
            return name
        # spelldesc and spelltooltip.
        if spell_id in chain:
            logger.warning(
                "spelltext: $@%s%d cycles back to a spell already being "
                "described (%s); leaving the raw token",
                kind,
                spell_id,
                sorted(chain),
            )
            return match[0]
        if spell_id not in self._spells:
            # Wowhead's Forever tooltip prints nothing for a spell the client lacks: Libram
            # of Banishment 211472's "Engrave your pants with the Aura Mastery rune:" stops
            # there (its rune 407624 is a Season of Discovery spell).
            logger.warning(
                "spelltext: $@%s%d names a spell this build has no row for; dropping it",
                kind,
                spell_id,
            )
            return ""
        description = self.describe(spell_id, _chain=chain)
        name = self._names.get(spell_id)
        if kind == "spelltooltip" and name:
            return f"{name}\r\n{description}"
        return description

    def _substitute(self, row: SpellRow, match: re.Match[str]) -> str:
        target = self._target(row, match)
        if target is None:
            return match[0]
        kind = match["kind"]
        if kind.lower() == "z":
            return _HOME
        if kind.lower() == "d" and target.until_cancelled:
            return _UNTIL_CANCELLED
        amount = _amount(target, kind.lower(), kind == "M", int(match["index"] or 1) - 1)
        if amount is None:
            return match[0]
        if amount.duration_ms is not None and not _scaled(match):
            return _render_duration(amount.duration_ms)
        multiplier, divisor = _scale(match)
        return _render(amount.low * multiplier, amount.high * multiplier, divisor)


def _base_points(row: dict[str, str]) -> float:
    """The effect's base points exactly as the client states them: Classic Era exports
    the integer column, the 1.60 client (Forever beta) the modern float column, and a
    float keeps its fraction (Wrack's tick 1349556 is 25.2, the Lightning Bolt rank 4
    overload 408442 is 36.5). Nothing here rounds: the game keeps the float and rounds
    only when a tooltip prints it, which `_bounds`, `_amount` and `work_out` do at
    display. The 1.60 client also has no EffectDieSides (its spread is the Variance
    column), so a missing die-sides column reads as no spread.

    Both builds' SpellEffect rows carry both column headers, but only one is ever
    populated: Era's own EffectBasePoints is always populated and its EffectBasePointsF
    is always "0" (unused padding), while build 1.60.1.69893's EffectBasePoints is
    always empty and EffectBasePointsF carries the real value -- confirmed against
    every row of both builds' own SpellEffect tables (see data/README.md). Reading
    the float column whenever it merely exists, as an earlier version of this
    function did, silently zeroed every Era description's numbers (see
    test_either_base_points_column_lands_on_the_same_number). Preferring the int
    column when both are populated is the same rule
    `pipeline/normalize/item_curves.py`'s `_budget` uses for RandPropPoints. When both
    are populated and differ (first seen 2026-09-30 on spell 370100, where wowhead's
    live-client value matches the integer column) the integer column wins and the row
    is logged -- see the inline note below.

    Era's EffectBasePointsF padding is the literal string "0", which `populated`
    (correctly) still reports as present -- a column being non-empty is not the same
    as it carrying real data. Comparing that padding against a real, nonzero
    EffectBasePoints made this raise on the vast majority of Era's own SpellEffect
    rows rather than only on a genuine disagreement, so a literal "0" float column is
    treated as Era's padding (not populated, for comparison purposes) whenever the int
    column is also populated. Any other float value is still compared, and a mismatch
    is logged with the integer column winning -- see
    test_base_points_columns_that_disagree_prefer_the_integer_column.
    """
    raw, float_raw = populated(row, "EffectBasePoints"), populated(row, "EffectBasePointsF")
    if raw is not None and float_raw is not None and float_raw != "0":
        int_value, float_value = int(raw), float(float_raw)
        if int_value != float_value:
            # 2026-09-30: the "only one column is ever real" premise broke on
            # wago.tools' fresh 1.60.1.70009 export -- spell 370100 (Illusionary
            # Rot) effect 0 carries EffectBasePoints 1 and EffectBasePointsF 4.
            # wowhead's live-client rendering of that effect is "2%": the Era
            # convention (integer column plus its one die side), not the float.
            # So the integer column wins a disagreement, the same precedence
            # `pipeline/normalize/item_curves.py`'s `_budget` already applies,
            # and every such row is logged so the count is visible in the run.
            logger.warning(
                "spell %s effect %s: EffectBasePoints %d and EffectBasePointsF %s disagree; "
                "using the integer column (wowhead corroborates it for spell 370100)",
                row.get("SpellID"),
                row.get("EffectIndex"),
                int_value,
                float_raw,
            )
        return float(int_value)
    if raw is not None:
        return float(int(raw))
    return float(row["EffectBasePointsF"])


def effect_amount(row: dict[str, str]) -> float:
    """One SpellEffect row's final amount, on either client schema, exactly as stated.

    Classic Era states an amount one short and puts the last point in
    EffectDieSides -- "Increase Spell Dam 85" is 84 + 1, "Attack Power 24" is
    23 + 1. The 1.60 client (Forever beta) has no EffectDieSides column at all
    and its EffectBasePointsF is already final: spell 9103, the "+8 Strength"
    enchant's spell, reads 8 where Era read 7, and Wrack's tick 1349556 reads
    25.2. Adding a missing die-sides column as zero makes one expression right
    on both.

    Public so `pipeline/simdb/equip.py` and `pipeline/simconst.py` read amounts
    the same way the description renderer does; the per-build column choice lives
    in `_base_points` and nowhere else, and no reader rounds (display rounds).
    """
    return _base_points(row) + int(row.get("EffectDieSides") or 0)


@dataclass(frozen=True)
class ExtraRows:
    """The tables behind the proc, radius, range, target, variable and `$@spellname` tokens.

    Optional as a group: a build fetched before they were on the list normalizes exactly as
    it did, with those tokens left raw (`$@spellname`/`$@spelldesc` warn and stay raw too --
    `SpellText._substitute_at`'s own doc -- `$@spelldesc` needs no table here at all, since
    it reads straight off the same `spells` dict `load_spell_text` already builds).
    """

    aura_options: list[dict[str, str]] = field(default_factory=list)
    radius: list[dict[str, str]] = field(default_factory=list)
    range: list[dict[str, str]] = field(default_factory=list)
    #: `SpellName.csv`'s own `ID, Name_lang` rows, for `$@spellname<id>`.
    names: list[dict[str, str]] = field(default_factory=list)
    #: `SpellTargetRestrictions.csv`, for `$i` and `$v`.
    target_restrictions: list[dict[str, str]] = field(default_factory=list)
    #: `SpellDescriptionVariables.csv` (ID, Variables), for `$<name>`.
    description_variables: list[dict[str, str]] = field(default_factory=list)
    #: `SpellXDescriptionVariables.csv` (SpellID -> SpellDescriptionVariablesID).
    spell_description_variables: list[dict[str, str]] = field(default_factory=list)
    #: `SpellLevels.csv`, for the per-level gain an effect adds (`_level_bonus`).
    levels: list[dict[str, str]] = field(default_factory=list)
    #: `SpellItemEnchantment.csv`, for `$ec<n>` and `$ecim`.
    item_enchantments: list[dict[str, str]] = field(default_factory=list)


def _base_difficulty(rows: list[dict[str, str]]) -> list[dict[str, str]]:
    return [row for row in rows if row.get("DifficultyID", "0") in ("0", "")]


def _aura_options(rows: list[dict[str, str]]) -> dict[int, dict[str, int]]:
    return {
        int(row["SpellID"]): {
            "proc_chance": int(row["ProcChance"]),
            "proc_charges": int(row["ProcCharges"]),
            "max_stacks": int(row["CumulativeAura"]),
            "proc_cooldown_ms": int(row["ProcCategoryRecovery"]),
        }
        for row in _base_difficulty(rows)
    }


def _target_restrictions(rows: list[dict[str, str]]) -> dict[int, dict[str, int]]:
    return {
        int(row["SpellID"]): {
            "max_targets": int(row.get("MaxTargets") or 0),
            "max_target_level": int(row.get("MaxTargetLevel") or 0),
        }
        for row in _base_difficulty(rows)
    }


def _variables(extra: ExtraRows) -> dict[int, dict[str, str]]:
    """Spell id -> its description variables, `$<name>` -> definition."""
    tables = {
        int(row["ID"]): parse_variables(row["Variables"]) for row in extra.description_variables
    }
    out: dict[int, dict[str, str]] = {}
    for row in extra.spell_description_variables:
        table = tables.get(int(row["SpellDescriptionVariablesID"]))
        if table:
            out.setdefault(int(row["SpellID"]), {}).update(table)
    return out


#: SPELL_EFFECT_ENCHANT_ITEM and SPELL_EFFECT_ENCHANT_ITEM_TEMPORARY.
_ENCHANT_EFFECTS = frozenset({"53", "54"})
_ENCHANT_AMOUNTS = ("EffectPointsMin_0", "EffectPointsMin_1", "EffectPointsMin_2")


def _enchants(
    effect_rows: list[dict[str, str]], enchant_rows: list[dict[str, str]]
) -> dict[int, dict[str, object]]:
    """Spell id -> the `SpellRow` fields of the enchantment its enchant-item effect applies."""
    enchantments = {int(row["ID"]): row for row in enchant_rows}
    out: dict[int, dict[str, object]] = {}
    for row in _base_difficulty(effect_rows):
        if row.get("Effect") not in _ENCHANT_EFFECTS:
            continue
        enchantment = enchantments.get(int(row.get("EffectMiscValue_0") or 0))
        if enchantment is not None:
            out.setdefault(
                int(row["SpellID"]),
                {
                    "enchant_points": tuple(int(enchantment[c] or 0) for c in _ENCHANT_AMOUNTS),
                    "enchant_min_item_level": int(enchantment.get("ItemLevelMin") or 0),
                },
            )
    return out


def _float(row: dict[str, str], column: str) -> float:
    return float(row.get(column) or 0)


def _levels_gained(rows: list[dict[str, str]]) -> dict[int, int]:
    """Spell id -> the levels a level-60 reader has over its BaseLevel, capped at MaxLevel.

    The client adds EffectRealPointsPerLevel for each one. Subtracting BaseLevel, not
    SpellLevel, and capping at MaxLevel is what Wowhead's Forever tooltips print: Frostbolt
    116 (levels 4 to 8) gains 4 levels, Cone of Cold 120 (26 to 31) 5, Dark Sacrifice
    1277324 (20 to 28) 8.
    """
    out: dict[int, int] = {}
    for row in _base_difficulty(rows):
        max_level = int(row.get("MaxLevel") or 0)
        level = min(READER_LEVEL, max_level) if max_level > 0 else READER_LEVEL
        out[int(row["SpellID"])] = max(level - int(row.get("BaseLevel") or 0), 0)
    return out


def _effects(
    rows: list[dict[str, str]], radii: dict[int, float], levels_gained: dict[int, int]
) -> dict[int, dict[int, Effect]]:
    effects: dict[int, dict[int, Effect]] = {}
    for row in _base_difficulty(rows):
        gained = levels_gained.get(int(row["SpellID"]), 0)
        effects.setdefault(int(row["SpellID"]), {})[int(row["EffectIndex"] or 0)] = Effect(
            base_points=_base_points(row),
            die_sides=int(row.get("EffectDieSides") or 0),
            period_ms=int(row["EffectAuraPeriod"] or 0),
            # Classic rows often leave index 0 unset and keep the radius in index 1
            # (Demoralizing Shout, Thunder Clap): either one is the effect's radius.
            radius=radii.get(int(row.get("EffectRadiusIndex_0") or 0))
            or radii.get(int(row.get("EffectRadiusIndex_1") or 0)),
            chain_targets=int(row.get("EffectChainTargets") or 0),
            misc_value=int(row.get("EffectMiscValue_0") or 0),
            points_per_resource=_float(row, "EffectPointsPerResource"),
            bonus_coefficient=_float(row, "EffectBonusCoefficient"),
            amplitude=_float(row, "EffectAmplitude"),
            chain_amplitude=_float(row, "EffectChainAmplitude"),
            variance=_float(row, "Variance"),
            level_bonus=_float(row, "EffectRealPointsPerLevel") * gained,
        )
    return effects


@dataclass(frozen=True)
class _Misc:
    """The SpellMisc facts a description reads."""

    duration_ms: int | None = None
    until_cancelled: bool = False
    icon_file_id: int = 0
    range_max: float | None = None


def _misc(
    rows: list[dict[str, str]], durations: dict[int, int], ranges: dict[int, float]
) -> dict[int, _Misc]:
    misc: dict[int, _Misc] = {}
    for row in _base_difficulty(rows):
        duration = durations.get(int(row["DurationIndex"] or 0))
        misc[int(row["SpellID"])] = _Misc(
            duration_ms=duration if duration is None or duration >= 0 else None,
            until_cancelled=duration is not None and duration < 0,
            icon_file_id=int(row["SpellIconFileDataID"] or 0),
            range_max=ranges.get(int(row.get("RangeIndex") or 0)),
        )
    return misc


def _fill_gaps(
    descriptions: dict[int, str], gaps: Mapping[int, Mapping[str, str]]
) -> dict[int, str]:
    """`curated/spelltext-gaps.json`'s values written into their templates
    (`pipeline.spelltext_gaps`); a gap the template no longer holds is an error."""
    out = dict(descriptions)
    for spell_id, tokens in gaps.items():
        for token, value in tokens.items():
            if token not in out.get(spell_id, ""):
                raise SpellTextError(
                    f"spell {spell_id} no longer reads {token!r}; drop it from "
                    "curated/spelltext-gaps.json"
                )
            out[spell_id] = re.sub(
                re.escape(token) + r"(?![A-Za-z\d])", value.replace("\\", r"\\"), out[spell_id]
            )
    return out


def load_spell_text(
    spell_rows: list[dict[str, str]],
    misc_rows: list[dict[str, str]],
    effect_rows: list[dict[str, str]],
    duration_rows: list[dict[str, str]],
    extra: ExtraRows | None = None,
    gaps: Mapping[int, Mapping[str, str]] | None = None,
) -> SpellText:
    """Every spell the tables know. A spell with SpellMisc or SpellEffect rows but no Spell
    row (no text of its own) is still there, with an empty description, so another spell's
    `$<id>d` can read its duration."""
    extra = extra or ExtraRows()
    radii = {int(r["ID"]): float(r["Radius"]) for r in extra.radius}
    ranges = {int(r["ID"]): float(r["RangeMax_0"]) for r in extra.range}
    durations = {int(r["ID"]): int(r["Duration"]) for r in duration_rows}
    aura_options = _aura_options(extra.aura_options)
    targets = _target_restrictions(extra.target_restrictions)
    variables = _variables(extra)
    enchants = _enchants(effect_rows, extra.item_enchantments)
    effects = _effects(effect_rows, radii, _levels_gained(extra.levels))
    misc = _misc(misc_rows, durations, ranges)
    descriptions = _fill_gaps(
        {int(row["ID"]): row["Description_lang"] for row in spell_rows}, gaps or {}
    )
    spells: dict[int, SpellRow] = {}
    for spell_id in descriptions.keys() | misc.keys() | effects.keys():
        facts = misc.get(spell_id, _Misc())
        spells[spell_id] = SpellRow(
            description=descriptions.get(spell_id, ""),
            duration_ms=facts.duration_ms,
            icon_file_id=facts.icon_file_id,
            effects=effects.get(spell_id, {}),
            range_max=facts.range_max,
            until_cancelled=facts.until_cancelled,
            variables=variables.get(spell_id, {}),
            **enchants.get(spell_id, {}),
            **targets.get(spell_id, {}),
            **aura_options.get(spell_id, {}),
        )
    names = {int(r["ID"]): r["Name_lang"] for r in extra.names}
    return SpellText(spells, names)
