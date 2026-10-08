"""A reader for the Lua a WoW client writes into WTF/.../SavedVariables/<Addon>.lua.

The file is a list of global assignments, `Name = <value>`, where a value is
a string, number, boolean, nil or a table constructor. That is all this reads:
no expressions, no function calls, no `local`. Tables whose keys are exactly
1..n come back as lists, every other table as a dict, and an empty table as
an empty list. Comments (`--` and `--[[ ]]`) are skipped.
"""

from __future__ import annotations

import re
from pathlib import Path
from typing import Any

_TOKEN = re.compile(
    r"""
    (?P<space>\s+|--\[(?P<level>=*)\[.*?\](?P=level)\]|--[^\n]*)
  | (?P<number>[-+]?(?:0[xX][0-9a-fA-F]+|(?:\d+\.?\d*|\.\d+)(?:[eE][-+]?\d+)?))
  | (?P<name>[A-Za-z_][A-Za-z0-9_]*)
  | (?P<string>"(?:[^"\\\n]|\\.|\\\n)*"|'(?:[^'\\\n]|\\.|\\\n)*')
  | (?P<punct>[{}\[\]=,;])
    """,
    re.VERBOSE | re.DOTALL,
)

_ESCAPES = {"n": "\n", "t": "\t", "r": "\r", "a": "\a", "b": "\b", "f": "\f", "v": "\v"}
_ESCAPE = re.compile(r"\\(\d{1,3}|.|\n)", re.DOTALL)
_LITERALS = {"true": True, "false": False, "nil": None}


class SavedVariablesError(ValueError):
    """The file is not the subset of Lua a SavedVariables file uses."""


def _unescape(body: str) -> str:
    def replace(match: re.Match[str]) -> str:
        char = match.group(1)
        if char.isdigit():
            return chr(int(char))
        return _ESCAPES.get(char, char)

    return _ESCAPE.sub(replace, body)


def _number(text: str) -> int | float:
    if text.lower().lstrip("+-").startswith("0x"):
        return int(text, 16)
    try:
        return int(text)
    except ValueError:
        return float(text)


def _tokens(text: str) -> list[tuple[str, Any]]:
    tokens: list[tuple[str, Any]] = []
    position = 0
    while position < len(text):
        match = _TOKEN.match(text, position)
        if match is None:
            line = text.count("\n", 0, position) + 1
            raise SavedVariablesError(f"unreadable Lua at line {line}")
        position = match.end()
        kind = match.lastgroup
        if kind == "level":
            continue
        if kind == "space":
            continue
        value = match.group(kind)
        if kind == "string":
            tokens.append(("value", _unescape(value[1:-1])))
        elif kind == "number":
            tokens.append(("value", _number(value)))
        elif kind == "name" and value in _LITERALS:
            tokens.append(("value", _LITERALS[value]))
        else:
            tokens.append((kind, value))
    return tokens


class _Parser:
    def __init__(self, tokens: list[tuple[str, Any]]) -> None:
        self.tokens = tokens
        self.index = 0

    def peek(self) -> tuple[str, Any]:
        return self.tokens[self.index] if self.index < len(self.tokens) else ("end", None)

    def take(self) -> tuple[str, Any]:
        token = self.peek()
        self.index += 1
        return token

    def expect(self, punct: str) -> None:
        kind, value = self.take()
        if (kind, value) != ("punct", punct):
            raise SavedVariablesError(f"expected {punct!r}, found {value!r}")

    def assignments(self) -> dict[str, Any]:
        result: dict[str, Any] = {}
        while self.peek()[0] != "end":
            kind, name = self.take()
            if kind != "name":
                raise SavedVariablesError(f"expected a variable name, found {name!r}")
            self.expect("=")
            result[name] = self.value()
        return result

    def value(self) -> Any:
        kind, token = self.peek()
        if (kind, token) == ("punct", "{"):
            return self.table()
        if kind == "value":
            self.take()
            return token
        raise SavedVariablesError(f"expected a value, found {token!r}")

    def table(self) -> list[Any] | dict[Any, Any]:
        self.expect("{")
        keyed: dict[Any, Any] = {}
        positional = 0
        while self.peek() != ("punct", "}"):
            if self.peek()[0] == "end":
                raise SavedVariablesError("unterminated table")
            positional = self.entry(keyed, positional)
            if self.peek() in (("punct", ","), ("punct", ";")):
                self.take()
        self.expect("}")
        return _as_sequence_if_dense(keyed)

    def entry(self, keyed: dict[Any, Any], positional: int) -> int:
        kind, token = self.peek()
        if (kind, token) == ("punct", "["):
            self.take()
            key = self.value()
            self.expect("]")
            self.expect("=")
            keyed[key] = self.value()
            return positional
        following = self.tokens[self.index + 1] if self.index + 1 < len(self.tokens) else None
        if kind == "name" and following == ("punct", "="):
            self.take()
            self.take()
            keyed[token] = self.value()
            return positional
        positional += 1
        keyed[positional] = self.value()
        return positional


def _as_sequence_if_dense(keyed: dict[Any, Any]) -> list[Any] | dict[Any, Any]:
    size = len(keyed)
    if all(isinstance(key, int) and not isinstance(key, bool) for key in keyed) and all(
        index in keyed for index in range(1, size + 1)
    ):
        return [keyed[index] for index in range(1, size + 1)]
    return keyed


def parse(text: str) -> dict[str, Any]:
    """The globals a SavedVariables file assigns."""
    return _Parser(_tokens(text)).assignments()


def read(path: Path) -> dict[str, Any]:
    """The globals in the SavedVariables file at `path`."""
    return parse(path.read_text(encoding="utf-8"))
