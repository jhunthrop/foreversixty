"""The parts of the client's description grammar that are not one effect's number.

`pipeline.spelltext` substitutes the per-effect tokens (`$s1`, `$d`, `$o1` ...).
This module owns the rest of the grammar the 1.60 client's Spell.Description_lang
and SpellDescriptionVariables use, each checked against Wowhead's Forever tooltips
(nether.wowhead.com/tooltip/spell/<id>?dataEnv=16) on 2026-11-10:

  $?s123[yes][no]       a condition on the reader. A page has no reader, so a test
  $?s1[a]?s2[b][else]   on a spell, aura or class (`s`, `a`, `c`, `pc`) is false and
  $?!s123[yes][no]      `!` makes it true; the chain reads on to its last branch, which
                        may stand after whitespace (Chronostatic Preservation 436516:
                        "[...]\\r\\n\\r\\n[Fuses ...]" reads "Fuses ...").
  $?$PL>=60[89][6]      a test on the reader's level, read at level 60, the level
                        Wowhead's Forever tooltips are written for (Blade Dance 400012
                        prints `${$PL*5}` as 60 * 5).
  $<name>               a SpellDescriptionVariables value, itself written in this grammar
                        (Frostbolt's `$<frostdamage>` is 1.0 for a mage with no Ice Shards).
  ${expr}               arithmetic, rounded to a whole number unless a `.N` suffix asks for
                        N decimals (Drink 430: `${($s1/5)*$d}` with 42 and 18 sec is 151;
                        431: 104 and 21 sec is 437). A stat the reader would bring
                        (`$AP`, `$SPI`, `$rap`) stays a word in the formula, as Wowhead
                        writes it: Dark Sacrifice 1277324 "(440 + Spirit) Mana".
  $l<one>:<many>;       the word that agrees with the nearest number before it, each one
                        on its own (Nature's Grasp 16689: "next 1 melee attack that
                        strikes").
"""

from __future__ import annotations

import ast
import math
import operator
import re
from collections.abc import Callable, Mapping
from dataclasses import dataclass

#: The level a description's `$PL` is read at: Forever's level cap, the level Wowhead's
#: Forever tooltips print (400012, 401955, 436516).
READER_LEVEL = 60

#: A stat the reader brings to a formula -> the words Wowhead's Forever tooltips print for it.
STAT_WORDS: dict[str, str] = {
    "AP": "Attack Power",  # Wyvern Strike 458436
    "apm": "Attack Power",  # Revenge 6572: ($<APRatio>*$apm) prints "(0.25 * 1) * Attack Power"
    "RAP": "Ranged Attack Power",
    "rap": "Ranged Attack Power",  # Summon Hawk 1293241
    "SP": "Spell Power",  # Seal of Righteousness 20154
    "sp": "Spell Power",
    "SPI": "Spirit",  # Dark Sacrifice 1277324
    "SPFR": "Frost spell power",  # Frost Ward 6143
    "SPH": "Spell Power",  # Seal of Righteousness 20154 prints its $SPH as "Spell Power"
    "SPS": "Shadow spell power",  # Demonic Brand 1293695
    # The other schools, named the way Wowhead names Frost and Shadow.
    "SPFI": "Fire spell power",
    "SPA": "Arcane spell power",
    "SPN": "Nature spell power",
    "bh": "Healing",  # Prayer of Mending 401859
    # Death Mark 1316709: Wowhead prints the bare token; the spell's own sentence names it
    # "Critical Strike chance".
    "mcrit": "Critical Strike chance",
}

_LEVEL_TEST = re.compile(r"\$PL(?P<op>>=|<=|==|!=|>|<|=)(?P<value>\d+)")
_COMPARE: dict[str, Callable[[int, int], bool]] = {
    ">=": operator.ge,
    "<=": operator.le,
    "==": operator.eq,
    "=": operator.eq,
    "!=": operator.ne,
    ">": operator.gt,
    "<": operator.lt,
}


def _matching_bracket(text: str, start: int) -> int:
    """The index of the `]` closing the `[` at `start`, or -1."""
    depth = 0
    for index in range(start, len(text)):
        if text[index] == "[":
            depth += 1
        elif text[index] == "]":
            depth -= 1
            if depth == 0:
                return index
    return -1


def _atom_holds(atom: str) -> bool:
    negate = atom.startswith("!")
    body = atom.lstrip("!")
    level = _LEVEL_TEST.fullmatch(body)
    holds = bool(level) and _COMPARE[level["op"]](READER_LEVEL, int(level["value"]))
    return holds != negate


def condition_holds(condition: str) -> bool:
    """One `$?` test, read for a level-60 reader who has no talent, aura or other spell."""
    condition = condition.strip("()")
    return any(
        all(_atom_holds(atom.strip("()")) for atom in alternative.split("&"))
        for alternative in condition.split("|")
    )


@dataclass(frozen=True)
class _Chain:
    end: int
    text: str


def _read_chain(text: str, start: int) -> _Chain | None:
    """The branch a `$?` chain at `start` shows, and where the chain ends."""
    cursor = start
    chosen: str | None = None
    while text.startswith("$?", cursor) or (cursor > start and text.startswith("?", cursor)):
        cursor += 2 if text.startswith("$?", cursor) else 1
        open_at = text.find("[", cursor)
        if open_at < 0:
            return None
        close_at = _matching_bracket(text, open_at)
        if close_at < 0:
            return None
        if chosen is None and condition_holds(text[cursor:open_at]):
            chosen = text[open_at + 1 : close_at]
        cursor = close_at + 1
    stripped = len(text) - len(text[cursor:].lstrip())
    if stripped < len(text) and text[stripped] == "[":
        close_at = _matching_bracket(text, stripped)
        if close_at >= 0:
            if chosen is None:
                chosen = text[stripped + 1 : close_at]
            cursor = close_at + 1
    return _Chain(cursor, chosen or "")


def resolve_conditions(text: str) -> str:
    """Every `$?` chain replaced by the branch a level-60 reader with nothing else sees."""
    out: list[str] = []
    cursor = 0
    while (start := text.find("$?", cursor)) >= 0:
        chain = _read_chain(text, start)
        if chain is None:
            break
        out.append(text[cursor:start])
        out.append(resolve_conditions(chain.text))
        cursor = chain.end
    out.append(text[cursor:])
    return "".join(out)


#: `$name=` where a variable starts: at the start, after a line break, or straight after the
#: previous value's closing brace or bracket (row 985 runs its variables together).
_VARIABLE_START = re.compile(r"(?:^|(?<=[\n}\]]))\$(?P<name>\w+)=", re.MULTILINE)


def parse_variables(definitions: str) -> dict[str, str]:
    """One SpellDescriptionVariables.Variables cell as name -> definition."""
    starts = list(_VARIABLE_START.finditer(definitions))
    out: dict[str, str] = {}
    for current, following in zip(starts, [*starts[1:], None], strict=True):
        end = following.start() if following else len(definitions)
        out[current["name"]] = definitions[current.end() : end].strip()
    return out


VARIABLE = re.compile(r"\$<(?P<name>\w+)>")


def drop_unknown_variables(text: str) -> str:
    """A `$<name>` outside any `${..}` that no variable defines, as nothing.

    The client prints an undefined variable as nothing: Earth Shield 974's "heal the
    shielded target for $<heal>." has no SpellXDescriptionVariables row, and Wowhead's
    Forever tooltip prints "heal the shielded target for ." (2026-11-10). One inside a
    formula keeps the formula raw instead (`work_out` cannot read it).
    """
    out: list[str] = []
    cursor = 0
    while (start := text.find("${", cursor)) >= 0:
        end = matching_brace(text, start + 1)
        if end < 0:
            break
        out.append(VARIABLE.sub("", text[cursor:start]))
        out.append(text[start : end + 1])
        cursor = end + 1
    out.append(VARIABLE.sub("", text[cursor:]))
    return "".join(out)


#: `|cFFFF2020Requires Cat Form|r`: the client's colour escapes. This text-only pipeline
#: keeps the words and drops the colour, as it drops `$@spellicon`.
COLOUR_ESCAPE = re.compile(r"\|[cC][0-9A-Fa-f]{8}|\|[rR]")

#: Two spaces inside a sentence: the client's own slip (Mind Flay 15407 "over $d  and
#: slowing", Enrage 5229 "but  reduces") or what an empty value leaves. A page shows one
#: space either way, so the text carries one. Two after a full stop is the client's own
#: house style ("Lasts 30 sec.  Must remain seated") and a line's leading indent is the
#: combo-point tables' layout ("   1 point  : 14 sec"); both stay.
MID_SENTENCE_SPACES = re.compile(r"(?<=[^\s.!?:;)\]\"'])  +(?=[^\s(:])")


def matching_brace(text: str, start: int) -> int:
    """The index of the `}` closing the `{` at `start`, or -1."""
    depth = 0
    for index in range(start, len(text)):
        if text[index] == "{":
            depth += 1
        elif text[index] == "}":
            depth -= 1
            if depth == 0:
                return index
    return -1


def flatten_braces(expression: str) -> str:
    """`${..}` nested inside a formula, as plain parentheses."""
    while (start := expression.find("${")) >= 0:
        end = matching_brace(expression, start + 1)
        if end < 0:
            return expression
        expression = f"{expression[:start]}({expression[start + 2 : end]}){expression[end + 1 :]}"
    return expression


# --- arithmetic ---------------------------------------------------------------------------

_BINARY: dict[type[ast.operator], tuple[Callable[[float, float], float], str, int]] = {
    ast.Add: (operator.add, "+", 1),
    ast.Sub: (operator.sub, "-", 1),
    ast.Mult: (operator.mul, "*", 2),
    ast.Div: (operator.truediv, "/", 2),
}


@dataclass(frozen=True)
class Formula:
    """Arithmetic that still holds a stat the reader brings, ready to print."""

    text: str
    precedence: int


Value = float | Formula


def format_number(value: float) -> str:
    if value == int(value):
        return str(int(value))
    return f"{value:.3f}".rstrip("0").rstrip(".")


def _operand(value: Value, precedence: int, right: bool) -> str:
    if isinstance(value, Formula):
        tighter = value.precedence > precedence or (value.precedence == precedence and not right)
        return value.text if tighter else f"({value.text})"
    return format_number(value)


def _identity(op: type[ast.operator], left: Value, right: Value) -> Value | None:
    """`x * 1`, `x / 1`, `x + 0`, `x - 0`, `1 * x`, `0 + x` as `x`: a talent multiplier the
    reader does not have (`$<mult>` is 1.0 without the talent) says nothing in a formula."""
    if right == 1.0 and op in (ast.Mult, ast.Div) or right == 0.0 and op in (ast.Add, ast.Sub):
        return left
    if left == 1.0 and op is ast.Mult or left == 0.0 and op is ast.Add:
        return right
    return None


#: The client's formula functions. A comparison on a stat the reader brings is read with
#: every stat at nothing, as Wowhead's Forever tooltips read it: Demonic Brand 1293695's
#: `$cond($gt($SPFI,$SPS),$SPFI,$SPS)` prints "Shadow spell power".
FUNCTIONS: dict[str, Callable[..., float]] = {
    "gt": lambda a, b: float(a > b),
    "lt": lambda a, b: float(a < b),
    "ge": lambda a, b: float(a >= b),
    "le": lambda a, b: float(a <= b),
    "eq": lambda a, b: float(a == b),
    "min": min,
    "max": max,
}


def _call(node: ast.Call, stats: Mapping[str, str], zero: bool) -> Value:
    name = node.func.id if isinstance(node.func, ast.Name) else ""
    if node.keywords or (name not in FUNCTIONS and name != "cond"):
        raise ValueError(f"not a formula function: {ast.dump(node)}")
    at_zero = [_fold(arg, stats, zero=True) for arg in node.args]
    if name == "cond":
        if len(node.args) != 3:
            raise ValueError("cond takes three arguments")
        return _fold(node.args[1] if at_zero[0] else node.args[2], stats, zero)
    if name in ("min", "max"):
        chosen = at_zero.index(FUNCTIONS[name](*at_zero))
        return _fold(node.args[chosen], stats, zero)
    return FUNCTIONS[name](*at_zero)


def _fold(node: ast.expr, stats: Mapping[str, str], zero: bool = False) -> Value:
    """`node` worked out; a stat stays a word, or counts as nothing when `zero`."""
    if isinstance(node, ast.Constant) and isinstance(node.value, (int, float)):
        return float(node.value)
    if isinstance(node, ast.Name) and node.id in stats:
        return 0.0 if zero else Formula(stats[node.id], 3)
    if isinstance(node, ast.Call):
        return _call(node, stats, zero)
    if isinstance(node, ast.UnaryOp) and isinstance(node.op, (ast.USub, ast.UAdd)):
        value = _fold(node.operand, stats, zero)
        if isinstance(node.op, ast.UAdd):
            return value
        return -value if isinstance(value, float) else Formula(f"-{_operand(value, 3, False)}", 3)
    if isinstance(node, ast.BinOp) and type(node.op) in _BINARY:
        apply, symbol, precedence = _BINARY[type(node.op)]
        left, right = _fold(node.left, stats, zero), _fold(node.right, stats, zero)
        if isinstance(left, float) and isinstance(right, float):
            return apply(left, right)
        if (kept := _identity(type(node.op), left, right)) is not None:
            return kept
        text = f"{_operand(left, precedence, False)} {symbol} {_operand(right, precedence, True)}"
        return Formula(text, precedence)
    raise ValueError(f"not arithmetic: {ast.dump(node)}")


_ARITHMETIC = re.compile(r"[\w\s+\-*/().,]+")


def work_out(expression: str, places: int | None, stats: Mapping[str, str]) -> str | None:
    """A worked `${..}` as the client prints it, or None when it is not arithmetic.

    A number is rounded half up to `places` decimals (a whole number by default) and its
    sign dropped, as for every other token here: the client divides a negative amount by
    a negative number to show it positive. A formula that still holds a stat is printed
    in brackets with the stat as a word.
    """
    if not _ARITHMETIC.fullmatch(expression) or not expression.strip():
        return None
    try:
        value = _fold(ast.parse(expression.strip(), mode="eval").body, stats)
    except (SyntaxError, ValueError, ZeroDivisionError):
        return None
    if isinstance(value, Formula):
        return f"({value.text})"
    scale = 10 ** (places or 0)
    rounded = math.floor(abs(value) * scale + 0.5) / scale
    # A value that rounds away to nothing keeps its own digits, as Wowhead's Forever
    # tooltips print it: set bonus 1226378's `${$m1/-1000}` of 20 reads "0.0200 sec".
    return format_number(rounded if rounded or not value else abs(value))


# --- plurals ------------------------------------------------------------------------------

_PLURAL = re.compile(r"\$[lL](?P<one>[^:;$]+):(?P<many>[^;$]+);")
_NUMBER = re.compile(r"\d+(?:\.\d+)?")


def agree_plurals(text: str) -> str:
    """Each `$l<one>:<many>;` as the word that agrees with the nearest number before it."""
    out: list[str] = []
    cursor = 0
    for match in _PLURAL.finditer(text):
        before = "".join(out) + text[cursor : match.start()]
        numbers = _NUMBER.findall(before)
        out.append(text[cursor : match.start()])
        if numbers:
            out.append(match["one"] if float(numbers[-1]) == 1 else match["many"])
        else:
            out.append(match[0])
        cursor = match.end()
    out.append(text[cursor:])
    return "".join(out)
