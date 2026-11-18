"""Every row of a WDC5 DB2, named by its WoWDBDefs definition, as CSV cells.

`pipeline.db2` reads one sparse table against a hand-written schema; this reads
any WDC5 table the client ships, the way wowdev/DBCD (the reader behind
wago.tools' CSV exports) does: every compression type (none, bitpacked,
signed, common, pallet, pallet array), sparse tables with inline strings,
id lists, copy tables, the relationship map, and sections left zero-filled
because their TACT key is unknown (skipped, as DBCD skips them).

`table_csv` renders the rows the way the pipeline's raw CSVs have always read:
wago's export shape, so nothing downstream changes. Arrays are `Name_<i>`,
columns in definition order, and a float is printed like wago's PHP renders
it (rounded to 11 decimals, at most 14 significant digits).
"""

from __future__ import annotations

import struct
from collections.abc import Iterator
from dataclasses import dataclass
from decimal import ROUND_HALF_UP, Context, Decimal

from pipeline.db2 import _HEADER_SIZE, _SECTION_SIZE, Db2Error, _BitCursor, _read_header
from pipeline.db2 import _read_sections as read_sections
from pipeline.dbd import Field

FLAG_SPARSE = 0x1
FLAG_SECONDARY_KEY = 0x2
NONE, IMMEDIATE, COMMON, PALLET, PALLET_ARRAY, SIGNED_IMMEDIATE = range(6)
_PADDING = bytes(8)
_ELEVEN_PLACES = Decimal("1e-11")
_WIDE = Context(prec=80)


@dataclass(frozen=True)
class _Column:
    record_offset: int
    size: int
    extra_size: int
    compression: int
    a: int
    b: int
    c: int


def _signed(value: int, bits: int) -> int:
    value &= (1 << bits) - 1
    return value - (1 << bits) if value & (1 << (bits - 1)) else value


def _typed(raw: int, field: Field) -> int | float:
    if field.kind == "float":
        return struct.unpack("<f", struct.pack("<I", raw & 0xFFFFFFFF))[0]
    if field.signed:
        return _signed(raw, field.bits)
    return raw & ((1 << field.bits) - 1)


def format_float(value: float) -> str:
    """A float the way wago's export prints one: PHP's 14 significant digits
    after rounding to 11 decimals half away from zero, as PHP's round() does
    (0.4f is "0.40000000596", 393.510009765625 "393.51000976563")."""
    rounded = Decimal(value).quantize(_ELEVEN_PLACES, rounding=ROUND_HALF_UP, context=_WIDE)
    text = format(float(rounded), ".14G")
    if "E" not in text:
        return text
    mantissa, exponent = text.split("E")
    if "." not in mantissa:
        mantissa += ".0"
    return f"{mantissa}E{exponent[0]}{exponent[1:].lstrip('0') or '0'}"


#: Column names the export prefixes with "_" because they are SQL keywords
#: (Lock's `Index` array reads `_Index_0`...; pipeline/guild_events.py relies on it).
RESERVED_NAMES = frozenset({"Index"})


def _export_name(column: str) -> str:
    base, _, suffix = column.rpartition("_")
    if column in RESERVED_NAMES or (suffix.isdigit() and base in RESERVED_NAMES):
        return f"_{column}"
    return column


def _cell(value: object) -> str:
    if isinstance(value, float):
        return format_float(value)
    return str(value)


class _Reader:
    def __init__(self, data: bytes, fields: list[Field]) -> None:
        self.data = data
        self.fields = fields
        self.header = _read_header(data)
        if self.header.version != 5 and data[:4] == b"WDC5":
            raise Db2Error(f"WDC5 version {self.header.version}")
        self.sections = read_sections(data, self.header)
        inline = [f for f in fields if f.inline]
        if len(inline) != self.header.field_count:
            raise Db2Error(
                f"definition has {len(inline)} inline fields, file has {self.header.field_count}"
            )
        offset = _HEADER_SIZE + _SECTION_SIZE * self.header.section_count
        count = self.header.field_count
        self.field_bits = [struct.unpack_from("<hH", data, offset + 4 * i)[0] for i in range(count)]
        offset += 4 * count
        self.columns = [
            _Column(*struct.unpack_from("<HHIIiii", data, offset + 24 * i)) for i in range(count)
        ]
        offset += 24 * count
        self.pallets: dict[int, list[int]] = {}
        for kind in (PALLET, PALLET_ARRAY):
            for i, col in enumerate(self.columns):
                if col.compression == kind:
                    n = col.extra_size // 4
                    self.pallets[i] = list(struct.unpack_from(f"<{n}I", data, offset))
                    offset += col.extra_size
        self.common: dict[int, dict[int, int]] = {}
        for i, col in enumerate(self.columns):
            if col.compression == COMMON:
                pairs = struct.unpack_from(f"<{col.extra_size // 4}I", data, offset)
                self.common[i] = {
                    _signed(pairs[j], 32): pairs[j + 1] for j in range(0, len(pairs), 2)
                }
                offset += col.extra_size
        strings = bytearray()
        for section in self.sections:
            if not self.sparse:
                start = section.file_offset + section.record_count * self.header.record_size
                strings += data[start : start + section.string_table_size]
        self.strings = bytes(strings)

    @property
    def sparse(self) -> bool:
        return bool(self.header.flags & FLAG_SPARSE)

    def _string_at(self, index: int) -> str:
        if index < 0:
            index = 0
        end = self.strings.index(b"\x00", index)
        return self.strings[index:end].decode("utf-8", errors="replace")

    def _scalar(self, cursor: _BitCursor, fi: int, field: Field, row_id: int, string_base: int):
        col = self.columns[fi]
        if col.compression == NONE:
            bits = 32 - self.field_bits[fi]
            bits = bits if bits > 0 else col.b
            if field.kind in ("string", "locstring"):
                if self.sparse:
                    return cursor.read_cstring()
                position = cursor.bit_pos >> 3
                return self._string_at(string_base + position + _signed(cursor.read_uint(bits), 32))
            return _typed(cursor.read_uint(bits), field)
        if col.compression == IMMEDIATE:
            return _typed(cursor.read_uint(col.b), field)
        if col.compression == SIGNED_IMMEDIATE:
            return _typed(cursor.read_int(col.b) & ((1 << 64) - 1), field)
        if col.compression == COMMON:
            return _typed(self.common[fi].get(row_id, col.a & 0xFFFFFFFF), field)
        if col.compression in (PALLET, PALLET_ARRAY):
            return _typed(self.pallets[fi][cursor.read_uint(col.b)], field)
        raise Db2Error(f"field {field.name}: unknown compression {col.compression}")

    def _array(self, cursor: _BitCursor, fi: int, field: Field, string_base: int) -> list:
        col = self.columns[fi]
        if col.compression == PALLET_ARRAY:
            index = cursor.read_uint(col.b)
            return [_typed(self.pallets[fi][index * col.c + j], field) for j in range(col.c)]
        if col.compression != NONE:
            raise Db2Error(f"array {field.name}: unsupported compression {col.compression}")
        bits = 32 - self.field_bits[fi]
        bits = bits if bits > 0 else col.b
        values: list = []
        for _ in range(col.size // (32 if field.kind != "int" else field.bits)):
            if field.kind in ("string", "locstring"):
                position = cursor.bit_pos >> 3
                raw = _signed(cursor.read_uint(bits), 32)
                values.append(self._string_at(string_base + position + raw))
            else:
                values.append(_typed(cursor.read_uint(bits), field))
        return values

    def rows(self) -> Iterator[dict[str, object]]:
        id_field = next((f for f in self.fields if f.is_id), None)
        if id_field is None:
            raise Db2Error("definition names no $id$ column")
        previous = 0
        copies: dict[int, int] = {}
        by_id: dict[int, dict[str, object]] = {}
        for section in self.sections:
            for row in self._section_rows(section, previous, copies, id_field):
                by_id[int(row[id_field.name])] = row  # type: ignore[call-overload]
                yield row
            previous += section.record_count
        for dest, source in copies.items():
            if source in by_id:
                yield {**by_id[source], id_field.name: dest}

    def _section_rows(
        self, section, previous: int, copies: dict[int, int], id_field: Field
    ) -> Iterator[dict[str, object]]:
        header, data = self.header, self.data
        size = header.record_size
        if self.sparse:
            records_end = section.offset_records_end
            records = data[section.file_offset : records_end]
            offset = records_end
        else:
            records = data[section.file_offset : section.file_offset + section.record_count * size]
            offset = section.file_offset + section.record_count * size + section.string_table_size
        if section.tact_key_hash and not any(records):
            if section.id_list_size or section.copy_table_count:
                empty = struct.unpack_from("<i", data, offset)[0] == 0
            elif section.offset_map_id_count:
                empty = struct.unpack_from("<IH", data, offset)[1] == 0
            else:
                empty = True
            if empty:
                return
        ids = list(struct.unpack_from(f"<{section.id_list_size // 4}i", data, offset))
        offset += section.id_list_size
        if ids and not any(ids):
            ids = list(range(header.min_id + previous, header.min_id + previous + len(ids)))
        for _ in range(section.copy_table_count):
            dest, source = struct.unpack_from("<ii", data, offset)
            offset += 8
            if dest != source:
                copies[dest] = source
        sizes: list[int] = []
        if section.offset_map_id_count:
            n = section.offset_map_id_count
            sizes = [struct.unpack_from("<IH", data, offset + 6 * i)[1] for i in range(n)]
            offset += 6 * n
            if header.flags & FLAG_SECONDARY_KEY:
                ids = list(struct.unpack_from(f"<{n}i", data, offset))
                offset += 4 * n
        relations: dict[int, int] = {}
        if section.relationship_data_size:
            count = struct.unpack_from("<i", data, offset)[0]
            offset += 12
            for i in range(count):
                foreign, index = struct.unpack_from("<ii", data, offset + 8 * i)
                relations[index] = foreign
            offset += 8 * count
        if section.offset_map_id_count and not header.flags & FLAG_SECONDARY_KEY:
            ids = list(struct.unpack_from(f"<{section.offset_map_id_count}i", data, offset))
        position = 0
        for i in range(section.record_count):
            if self.sparse:
                buf = records[position : position + sizes[i]] + _PADDING
                position += sizes[i]
            else:
                buf = records[i * size : (i + 1) * size] + _PADDING
            row_id = ids[i] if ids else -1
            ref_key = row_id if header.flags & FLAG_SECONDARY_KEY else i
            global_index = previous + i
            string_base = (global_index - header.record_count) * size
            yield self._row(buf, row_id, relations.get(ref_key, 0), string_base, id_field)

    def _row(
        self, buf: bytes, row_id: int, ref: int, string_base: int, id_field: Field
    ) -> dict[str, object]:
        if row_id == -1:
            if not id_field.inline:
                raise Db2Error("the id is not in the record and the file has no id list")
            # Common-compressed columns are keyed by the row's id, so it is read first.
            row_id = self._inline_id(buf, id_field)
        cursor = _BitCursor(buf)
        values: dict[str, object] = {}
        fi = 0
        for field in self.fields:
            if not field.inline:
                values[field.name] = ref if field.relation else row_id
                continue
            if field.count > 1:
                for name, value in zip(
                    field.column_names(), self._array(cursor, fi, field, string_base), strict=True
                ):
                    values[name] = value
            else:
                values[field.name] = self._scalar(cursor, fi, field, row_id, string_base)
            fi += 1
        values[id_field.name] = row_id
        return values

    def _inline_id(self, buf: bytes, id_field: Field) -> int:
        fi = self.header.id_index
        col = self.columns[fi]
        if self.sparse or col.compression == COMMON:
            raise Db2Error(f"inline id with compression {col.compression} is not supported")
        cursor = _BitCursor(buf)
        cursor.bit_pos = col.record_offset
        return int(self._scalar(cursor, fi, id_field, -1, 0))  # type: ignore[arg-type]


def read_rows(data: bytes, fields: list[Field]) -> list[dict[str, object]]:
    """Every row of a WDC5 table, keyed by the definition's column names."""
    if data[:4] != b"WDC5":
        raise Db2Error(f"not a WDC5 file (magic {data[:4]!r})")
    header = _read_header(data)
    if header.record_count == 0 or header.section_count == 0:
        return []  # the client ships the table empty (the Dungeon Journal in 1.60)
    return list(_Reader(data, fields).rows())


def layout_hash(data: bytes) -> int:
    return _read_header(data).layout_hash


_QUOTED = (" ", ",", '"', "\n", "\r")


def _csv_text(value: str) -> str:
    """A string cell the way the export writes one: quoted when it holds a space,
    a comma, a quote or a line break, quotes doubled."""
    if any(c in value for c in _QUOTED):
        return '"' + value.replace('"', '""') + '"'
    return value


def table_csv(data: bytes, fields: list[Field]) -> str:
    """The table as CSV text in the raw-export shape (see the module doc)."""
    columns = [name for field in fields for name in field.column_names()]
    strings = {
        name
        for field in fields
        if field.kind in ("string", "locstring")
        for name in field.column_names()
    }
    lines = [",".join(_export_name(name) for name in columns)]
    for row in read_rows(data, fields):
        lines.append(
            ",".join(
                _csv_text(str(row[name])) if name in strings else _cell(row[name])
                for name in columns
            )
        )
    return "\n".join(lines) + "\n"
