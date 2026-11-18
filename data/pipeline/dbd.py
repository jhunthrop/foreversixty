"""WoWDBDefs (`.dbd`) table definitions: a DB2's column names, types and order.

A DB2 file carries bit widths and compression but no names; the community's
definitions (https://github.com/wowdev/WoWDBDefs, the source wago.tools and
DBCD read) name each column per layout hash and build. `read_definition`
picks the block for one file: the block listing the exact build, else the
first (newest) block listing the file's layout hash.

`DBD_COMMIT` pins the definitions, so two runs of one build read the same
column names; bump it when a new build's layout is missing (the fetch fails
naming the table and layout hash).
"""

from __future__ import annotations

import re
from dataclasses import dataclass

import httpx

#: wowdev/WoWDBDefs master on 2026-11-17, which lists every 1.60.1 build to 70959.
DBD_COMMIT = "5ebf2124996b538d6784584e69db6fe8d6bd8d06"
DBD_URL = "https://raw.githubusercontent.com/wowdev/WoWDBDefs/{commit}/definitions/{table}.dbd"

_ENTRY = re.compile(
    r"^(?:\$(?P<annotations>[^$]*)\$)?(?P<name>[A-Za-z_0-9]+)"
    r"(?:<(?P<size>u?\d+)>)?(?:\[(?P<count>\d+)\])?"
)


class DbdError(ValueError):
    """No definition fits this table's file."""


@dataclass(frozen=True)
class Field:
    """One entry of a definition's layout, in the file's column order."""

    name: str
    kind: str  # "int", "float", "string" or "locstring"
    bits: int  # 8/16/32/64 for an int, 32 otherwise
    signed: bool
    count: int  # array length, 1 for a scalar
    is_id: bool
    inline: bool  # False: not stored in the record (an id list or a relation)
    relation: bool

    def column_names(self) -> list[str]:
        if self.count == 1:
            return [self.name]
        return [f"{self.name}_{i}" for i in range(self.count)]


def _strip_comment(line: str) -> str:
    return line.split("//", 1)[0].strip()


def _column_types(lines: list[str]) -> dict[str, str]:
    types: dict[str, str] = {}
    for line in lines:
        line = _strip_comment(line)
        if not line:
            continue
        kind, _, name = line.partition(" ")
        types[name.rstrip("?")] = kind.split("<", 1)[0]
    return types


def _builds_match(line: str, build: str) -> bool:
    target = tuple(int(p) for p in build.split("."))
    for item in line.removeprefix("BUILD").split(","):
        item = item.strip()
        if "-" in item:
            low, high = (tuple(int(p) for p in b.split(".")) for b in item.split("-"))
            if low <= target <= high:
                return True
        elif item == build:
            return True
    return False


def _parse_field(line: str, types: dict[str, str]) -> Field:
    match = _ENTRY.match(line)
    if match is None:
        raise DbdError(f"cannot read layout entry {line!r}")
    annotations = set((match.group("annotations") or "").split(","))
    name = match.group("name")
    kind = types.get(name)
    if kind is None:
        raise DbdError(f"layout entry {name!r} has no COLUMNS type")
    size = match.group("size") or "32"
    return Field(
        name=name,
        kind=kind,
        bits=int(size.lstrip("u")),
        signed=not size.startswith("u"),
        count=int(match.group("count") or 1),
        is_id="id" in annotations,
        inline="noninline" not in annotations,
        relation="relation" in annotations,
    )


def read_definition(text: str, layout_hash: int, build: str) -> list[Field]:
    """The layout entries for a file with `layout_hash` in `build`."""
    blocks = [b.strip().splitlines() for b in re.split(r"\n\s*\n", text.strip())]
    types = _column_types(blocks[0][1:])
    layout = f"{layout_hash:08X}"
    by_build: list[str] | None = None
    by_layout: list[str] | None = None
    for block in blocks[1:]:
        heads = [line for line in block if line.startswith(("LAYOUT", "BUILD", "COMMENT"))]
        layouts = {
            h.strip() for line in heads if line.startswith("LAYOUT") for h in line[6:].split(",")
        }
        if layouts and layout not in layouts:
            continue
        if any(line.startswith("BUILD") and _builds_match(line, build) for line in heads):
            by_build = block
            break
        if layouts and by_layout is None:
            by_layout = block
    chosen = by_build or by_layout
    if chosen is None:
        raise DbdError(f"no definition for layout {layout} in build {build}")
    entries = [
        _strip_comment(line)
        for line in chosen
        if not line.startswith(("LAYOUT", "BUILD", "COMMENT")) and _strip_comment(line)
    ]
    return [_parse_field(line, types) for line in entries]


def fetch_definition(table: str, http: httpx.Client, commit: str = DBD_COMMIT) -> str:
    response = http.get(DBD_URL.format(commit=commit, table=table))
    response.raise_for_status()
    return response.text
