"""Shared, quote-aware helpers for reading a mysqldump's own `INSERT INTO
`table` VALUES (...),(...),...;` statements -- the same char-by-char
quote/paren tracking `pipeline.classic_quest_levels` proved out for
`quest_template` (see that module's own doc for why a naive regex over
the raw text is wrong: 180,935 false matches on the real 73MB dump for
what turned out to be 4,245 real quests, because a row's own free-form
text fields can contain literal commas, parens and semicolons inside a
quoted string).

`pipeline.classic_sources` (src-classicdb lane, 2026-09-29) is the
second consumer -- it needs this same row-boundary tracking for ten
more cmangos/classic-db tables, several of them with a text/comments
column of their own, so this is split out here rather than duplicated a
second time. `classic_quest_levels` keeps its own private copy (a
`_ROW_PREFIX` regex over the row's own leading integers is all it ever
needed, no full field split) -- refactoring it onto this module is a
follow-up, not done here, to avoid touching an already-shipped, tested
parser for a lane that does not need to.
"""

from __future__ import annotations

import re
from collections.abc import Iterator

_COLUMN_RE = re.compile(r"^\s*`(\w+)`", re.MULTILINE)


def _statement_span(text: str, marker: str, marker_start: int) -> str:
    """The full "(...),(...),..." value list of one INSERT statement
    starting at `marker_start` (the index `marker` was found at),
    stopping at the statement's own top-level `;` -- one outside any
    quoted string and at paren depth 0."""
    i = marker_start + len(marker)
    n = len(text)
    while i < n and text[i] in " \n\t":
        i += 1
    depth = 0
    in_string = False
    j = i
    while j < n:
        c = text[j]
        if in_string:
            if c == "\\":
                j += 2
                continue
            if c == "'":
                in_string = False
        else:
            if c == "'":
                in_string = True
            elif c == "(":
                depth += 1
            elif c == ")":
                depth -= 1
            elif c == ";" and depth == 0:
                return text[i:j]
        j += 1
    raise ValueError(f"{marker} statement has no top-level terminator")


def _split_rows(statement: str) -> Iterator[str]:
    """Every "(...)" row's own inner content in one INSERT statement's
    value list, quote-aware."""
    depth = 0
    in_string = False
    row_start = 0
    n = len(statement)
    j = 0
    while j < n:
        c = statement[j]
        if in_string:
            if c == "\\":
                j += 2
                continue
            if c == "'":
                in_string = False
        else:
            if c == "'":
                in_string = True
            elif c == "(":
                if depth == 0:
                    row_start = j + 1
                depth += 1
            elif c == ")":
                depth -= 1
                if depth == 0:
                    yield statement[row_start:j]
        j += 1


def table_columns(sql_text: str, table: str) -> list[str]:
    """`table`'s own column names, in declared order, read straight off
    its "CREATE TABLE `table` (...)" block. This is what lets a caller
    zip a row's `split_fields` result up by NAME (`dict(zip(table_columns
    (...), split_fields(row)))`) instead of hand-counted positional
    indices -- a 130-column table like `quest_template` is exactly the
    case a miscounted index silently reads the wrong column for, so this
    is the one place the column order is trusted, straight from the
    dump's own schema rather than retyped by a person.
    """
    marker = f"CREATE TABLE `{table}` ("
    start = sql_text.find(marker)
    if start == -1:
        raise ValueError(f"no CREATE TABLE statement for `{table}`")
    end = sql_text.find(") ENGINE", start)
    if end == -1:
        raise ValueError(f"CREATE TABLE `{table}` has no closing ) ENGINE=...")
    return _COLUMN_RE.findall(sql_text[start:end])


def iter_table_rows(sql_text: str, table: str) -> Iterator[str]:
    """Every row (its raw, un-split "a,b,'c,d'" inner text) from every
    "INSERT INTO `table` VALUES ...;" statement anywhere in a
    mysqldump's text -- a real dump splits a large table across more
    than one INSERT statement, same as `classic_quest_levels`' own
    `parse_quest_template` handles for `quest_template`.
    """
    marker = f"INSERT INTO `{table}` VALUES"
    start = sql_text.find(marker)
    while start != -1:
        statement = _statement_span(sql_text, marker, start)
        yield from _split_rows(statement)
        start = sql_text.find(marker, start + len(marker))


def split_fields(row: str) -> list[str]:
    """One row's own top-level fields, quote-aware: a bare token
    (`123`, `-45`, `NULL`) or a still-single-quoted string
    (`'it\\'s here'`) exactly as written -- `unquote` turns the latter
    into the real string. Splitting only at top-level commas (outside a
    quoted string, backslash-escapes tracked) is what
    `classic_quest_levels`' own doc explains is required at all.
    """
    fields: list[str] = []
    depth = 0
    in_string = False
    start = 0
    n = len(row)
    j = 0
    while j < n:
        c = row[j]
        if in_string:
            if c == "\\":
                j += 2
                continue
            if c == "'":
                in_string = False
        else:
            if c == "'":
                in_string = True
            elif c == "(":
                depth += 1
            elif c == ")":
                depth -= 1
            elif c == "," and depth == 0:
                fields.append(row[start:j])
                start = j + 1
        j += 1
    fields.append(row[start:])
    return fields


def iter_table_records(sql_text: str, table: str) -> Iterator[dict[str, str]]:
    """`iter_table_rows` + `split_fields`, zipped up against
    `table_columns` -- one dict per row, column name -> the field's own
    raw (still-quoted-if-a-string) text. The convenience form every
    `pipeline.classic_sources` table reader uses; a caller that wants a
    real Python value calls `int()`/`float()` or `unquote()` on the
    field it asked for, same as reading any other row-shaped input in
    this pipeline.
    """
    columns = table_columns(sql_text, table)
    for row in iter_table_rows(sql_text, table):
        fields = split_fields(row)
        if len(fields) != len(columns):
            raise ValueError(
                f"`{table}` row has {len(fields)} fields, schema names {len(columns)}: "
                f"{row[:80]!r}"
            )
        yield dict(zip(columns, fields, strict=True))


#: mysqldump's own backslash-escape set (MySQL string literal escape
#: sequences it actually emits) -- a single left-to-right pass over
#: these, not sequential `.replace()` calls, which would misfire the
#: moment one escape's OUTPUT contains another escape's own trigger
#: character (`\\'` unescaping to `'` first, for instance, would then
#: falsely feed a later `\\"`-style pass). `\"` is the one real dump rows
#: were found to need (a creature name like `\"Pretty Boy\" Duncan`,
#: src-classicdb lane, 2026-09-29) -- MySQL escapes a literal `"` even
#: inside a single-quoted string.
_ESCAPES = {
    "'": "'", '"': '"', "\\": "\\", "0": "\0",
    "b": "\b", "n": "\n", "r": "\r", "t": "\t", "Z": "\x1a",
}


def _unescape(text: str) -> str:
    out: list[str] = []
    i = 0
    n = len(text)
    while i < n:
        c = text[i]
        if c == "\\" and i + 1 < n and text[i + 1] in _ESCAPES:
            out.append(_ESCAPES[text[i + 1]])
            i += 2
        else:
            out.append(c)
            i += 1
    return "".join(out)


def unquote(field: str) -> str | None:
    """A `split_fields` string field's real value: `None` for SQL
    `NULL`, otherwise the quotes stripped and every mysqldump backslash
    escape (`_ESCAPES`) resolved in one left-to-right pass. Passing a
    bare numeric field is a caller error (use `int()`/`float()` on those
    directly) but harmless -- it round-trips unchanged."""
    field = field.strip()
    if field == "NULL":
        return None
    if len(field) >= 2 and field[0] == "'" and field[-1] == "'":
        return _unescape(field[1:-1])
    return field
