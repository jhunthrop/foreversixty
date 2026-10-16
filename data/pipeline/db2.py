"""Minimal reader for Blizzard's WDC5 ("DB2") client database format.

Scope: **sparse** tables -- ones whose records sit behind an offset map with
variable-length, inline null-terminated strings (`header.flags & 0x01`, see
`FLAG_OFFSET_MAP`). That is the layout `ItemSparse.db2` uses on
`wow_classic_beta` build `1.60.1.70009` (and every other 1.60.1 build wago.tools
currently serves -- see the module-level finding below). Fixed-size tables
with an external string table (e.g. `Item.db2`, which uses the same magic but
*not* the offset-map flag) are out of scope: `pipeline.wago`'s CSV export of
`Item.csv` is already complete for this build (31,818/31,819 rows -- the one
row difference is `Item.csv`'s own header/footer accounting, not a gap), so
there is nothing this reader would add for that table.

Format reference: https://wowdev.wiki/DB2 (the "WDC5" section) via its API
(`action=parse`, since the wiki disables `action=raw`). The exact column
order/widths for `ItemSparse` on this build come from WoWDBDefs' layout hash
`6FCC3191` (matches this build's `header.layout_hash` exactly):
https://raw.githubusercontent.com/wowdev/WoWDBDefs/master/definitions/ItemSparse.dbd
(`BUILD 1.60.1.69876, 1.60.1.69893, 1.60.1.69913, 1.60.1.69977, 1.60.1.70009,
1.60.1.70058, 1.60.1.70094`). Because `.dbd` files carry no field-storage
compression details of their own, and this reader only needs to support the
compression type `ItemSparse`/`Item` actually use on this build
(`field_compression_none` -- see `night-relics` finding below), it deliberately
does not parse `field_storage_info`/`pallet_data`/`common_data` at all: it
reads `ITEM_SPARSE_SCHEMA`'s declared bit widths sequentially per record,
the same way `field_compression_none` columns are read
(https://github.com/wowdev/DBCD's `WDC5Reader`/`WDC5Row.GetFieldValue`, case
`CompressionType.None`). `read_sparse_rows` raises `Db2Error` up front if a
file needs pallet or common data, so a table this reader can't handle fails
loudly instead of silently decoding wrong values.

**night-relics finding, 2026-09-29, superseded (see accuracy audit,
2026-09-29 evening):** `pipeline/csvio.py`'s `check_item_sparse_completeness`
assumed wago.tools' CSV export was *truncating* `ItemSparse.csv` to
19,226/31,819 rows. It is not truncating anything. This reader decodes
wago.tools' raw CASC download of the same table
(`GET /api/casc/1572924?build=1.60.1.70009`) -- bypassing their CSV renderer
entirely -- and gets the *same* ~19,225 row count. For the two items the
2026-09-29 accuracy audit flagged (Acidproof Cloak, ID 3582; Amulet of the
Dawn, ID 22657), this reader's decode of the raw record bytes matches
wago.tools' CSV exactly, field for field (`ItemLevel` 18/60, `Material` 7,
`Bonding` 1, `InventoryType` 16, and `RequiredLevel` 0 for both -- confirmed
against the raw hex, not just the CSV). `ItemSparse.db2` is byte-identical
across build `1.60.1.70009`, the later `1.60.1.70058`, and CN's
`1.60.1.70094` (same md5). **The client's own shipped `ItemSparse` table is
genuinely ~60% populated for this content wave, and a meaningful share of
the rows that *are* present carry a zeroed `RequiredLevel`** (roughly half,
by sampling section 0 -- some legitimately, e.g. starter gear, but at least
these two do not). No fetch strategy recovers data the client was never
shipped with; see the lane report for the full trail (`scratchpad/day3/...`).
"""

import struct
from dataclasses import dataclass
from pathlib import Path

MAGIC = b"WDC5"

_HEADER_FMT = "<4sI128s9I2H7I"
_HEADER_SIZE = struct.calcsize(_HEADER_FMT)
_SECTION_FMT = "<QIIIIIIII"
_SECTION_SIZE = struct.calcsize(_SECTION_FMT)
_OFFSET_MAP_ENTRY_FMT = "<IH"
_OFFSET_MAP_ENTRY_SIZE = struct.calcsize(_OFFSET_MAP_ENTRY_FMT)

#: `header.flags` bit meaning "records are addressed through an offset map,
#: with variable-length, inline null-terminated strings" -- wowdev.wiki's
#: "Known Flag Meanings" calls this 0x01, "Has offset map".
FLAG_OFFSET_MAP = 0x01


class Db2Error(Exception):
    """Raised when a `.db2` file isn't a WDC5 sparse table this reader supports."""


@dataclass(frozen=True)
class Column:
    """One column of a table's schema (e.g. `ItemSparse.dbd`'s layout for a
    specific `layout_hash`), in on-disk field order.

    `kind` is `"str"` (inline null-terminated string), `"i"` (signed
    integer), `"u"` (unsigned integer) or `"f"` (32-bit float). `bits` is the
    per-element bit width, ignored for `"str"`. `count` is the array length
    (`1` for a scalar column); an array decodes to a `list`.
    """

    name: str
    kind: str
    count: int = 1
    bits: int = 32


@dataclass(frozen=True)
class Section:
    """One `wdc5_section_header` (wowdev.wiki's WDC5 struct)."""

    tact_key_hash: int
    file_offset: int
    record_count: int
    string_table_size: int
    offset_records_end: int
    id_list_size: int
    relationship_data_size: int
    offset_map_id_count: int
    copy_table_count: int

    @property
    def encrypted(self) -> bool:
        """A TACT-key-encrypted section: its rows are unavailable without the
        decryption key, the same as they are to wago.tools' own CSV export."""
        return self.tact_key_hash != 0


@dataclass(frozen=True)
class Header:
    """`wdc5_db2_header` (wowdev.wiki's WDC5 struct)."""

    version: int
    schema: str
    record_count: int
    field_count: int
    record_size: int
    string_table_size: int
    table_hash: int
    layout_hash: int
    min_id: int
    max_id: int
    locale: int
    flags: int
    id_index: int
    total_field_count: int
    bitpacked_data_offset: int
    lookup_column_count: int
    field_storage_info_size: int
    common_data_size: int
    pallet_data_size: int
    section_count: int

    @property
    def sparse(self) -> bool:
        return bool(self.flags & FLAG_OFFSET_MAP)


def _read_header(data: bytes) -> Header:
    if len(data) < _HEADER_SIZE or data[:4] != MAGIC:
        raise Db2Error(f"not a WDC5 file (magic {data[:4]!r}, want {MAGIC!r})")
    (
        _magic,
        version,
        schema_raw,
        record_count,
        field_count,
        record_size,
        string_table_size,
        table_hash,
        layout_hash,
        min_id,
        max_id,
        locale,
        flags,
        id_index,
        total_field_count,
        bitpacked_data_offset,
        lookup_column_count,
        field_storage_info_size,
        common_data_size,
        pallet_data_size,
        section_count,
    ) = struct.unpack_from(_HEADER_FMT, data, 0)
    return Header(
        version=version,
        schema=schema_raw.split(b"\x00", 1)[0].decode("ascii", errors="replace"),
        record_count=record_count,
        field_count=field_count,
        record_size=record_size,
        string_table_size=string_table_size,
        table_hash=table_hash,
        layout_hash=layout_hash,
        min_id=min_id,
        max_id=max_id,
        locale=locale,
        flags=flags,
        id_index=id_index,
        total_field_count=total_field_count,
        bitpacked_data_offset=bitpacked_data_offset,
        lookup_column_count=lookup_column_count,
        field_storage_info_size=field_storage_info_size,
        common_data_size=common_data_size,
        pallet_data_size=pallet_data_size,
        section_count=section_count,
    )


def _read_sections(data: bytes, header: Header) -> list[Section]:
    sections = []
    offset = _HEADER_SIZE
    for _ in range(header.section_count):
        sections.append(Section(*struct.unpack_from(_SECTION_FMT, data, offset)))
        offset += _SECTION_SIZE
    return sections


class _BitCursor:
    """Reads little-endian, arbitrary-bit-width values sequentially from a
    record's bytes, per wowdev.wiki's bitpacked-field formula: read the bytes
    spanning the requested bits, shift right by the bit's offset within its
    first byte, then mask to width.
    """

    def __init__(self, buf: bytes):
        self._buf = buf
        self.bit_pos = 0

    def read_uint(self, bits: int) -> int:
        byte_start = self.bit_pos // 8
        shift = self.bit_pos % 8
        nbytes = (bits + shift + 7) // 8
        raw = int.from_bytes(self._buf[byte_start : byte_start + nbytes], "little")
        value = (raw >> shift) & ((1 << bits) - 1)
        self.bit_pos += bits
        return value

    def read_int(self, bits: int) -> int:
        value = self.read_uint(bits)
        if value & (1 << (bits - 1)):
            value -= 1 << bits
        return value

    def read_cstring(self) -> str:
        if self.bit_pos % 8:
            raise Db2Error("string field is not byte-aligned; unsupported schema")
        start = self.bit_pos // 8
        end = self._buf.index(b"\x00", start)
        self.bit_pos = (end + 1) * 8
        return self._buf[start:end].decode("utf-8", errors="replace")


def decode_record(record: bytes, schema: list[Column]) -> dict[str, object]:
    """Decode one sparse record's raw bytes into `{column.name: value}`,
    reading `schema`'s columns in on-disk order. An array column decodes to a
    `list` of length `column.count`; a scalar column decodes to a bare value.
    """
    cursor = _BitCursor(record)
    row: dict[str, object] = {}
    for column in schema:
        values = []
        for _ in range(column.count):
            if column.kind == "str":
                values.append(cursor.read_cstring())
            elif column.kind == "f":
                raw = cursor.read_uint(32)
                values.append(struct.unpack("<f", struct.pack("<I", raw))[0])
            elif column.kind == "u":
                values.append(cursor.read_uint(column.bits))
            elif column.kind == "i":
                values.append(cursor.read_int(column.bits))
            else:
                raise Db2Error(f"unknown column kind {column.kind!r} for {column.name!r}")
        row[column.name] = values if column.count > 1 else values[0]
    return row


def read_sparse_rows(
    path: Path, schema: list[Column], *, expected_layout_hash: int | None = None
) -> dict[int, dict[str, object]]:
    """Decode every non-encrypted row of a sparse (offset-map) WDC5 table at
    `path`, keyed by ID.

    `schema` must match the file's on-disk column order exactly -- e.g. the
    WoWDBDefs `.dbd` layout whose hash equals the file's `header.layout_hash`
    (pass it as `expected_layout_hash` to fail fast on a build/layout
    mismatch rather than silently decoding garbage).

    TACT-key-encrypted sections (`Section.encrypted`) are skipped: their rows
    are unavailable without the decryption key, the same as to wago.tools'
    own CSV export.
    """
    data = path.read_bytes()
    header = _read_header(data)
    if not header.sparse:
        raise Db2Error(
            f"{path} has flags={header.flags:#x}; pipeline.db2 only reads the "
            f"offset-map ('sparse') layout ItemSparse uses, not a fixed-record "
            f"table with an external string table"
        )
    if header.common_data_size or header.pallet_data_size:
        raise Db2Error(
            f"{path} uses common_data/pallet_data-compressed columns "
            f"(common_data_size={header.common_data_size}, "
            f"pallet_data_size={header.pallet_data_size}); pipeline.db2 only "
            f"supports field_compression_none columns"
        )
    if expected_layout_hash is not None and header.layout_hash != expected_layout_hash:
        raise Db2Error(
            f"{path} has layout_hash={header.layout_hash:#x}, expected "
            f"{expected_layout_hash:#x}; `schema` is for a different build's "
            f"column layout"
        )
    if header.total_field_count != len(schema):
        raise Db2Error(
            f"{path} has {header.total_field_count} fields, but `schema` has "
            f"{len(schema)} columns"
        )

    rows: dict[int, dict[str, object]] = {}
    for section in _read_sections(data, header):
        if section.encrypted:
            continue
        copy_table_offset = section.offset_records_end + section.id_list_size
        copies = [
            struct.unpack_from("<II", data, copy_table_offset + i * 8)
            for i in range(section.copy_table_count)
        ]
        entry_count = section.offset_map_id_count
        entries_offset = copy_table_offset + section.copy_table_count * 8
        entries = [
            struct.unpack_from(
                _OFFSET_MAP_ENTRY_FMT, data, entries_offset + i * _OFFSET_MAP_ENTRY_SIZE
            )
            for i in range(entry_count)
        ]
        ids_offset = (
            entries_offset
            + entry_count * _OFFSET_MAP_ENTRY_SIZE
            + section.relationship_data_size
        )
        ids = struct.unpack_from(f"<{entry_count}I", data, ids_offset)
        for row_id, (record_offset, record_size) in zip(ids, entries, strict=True):
            record_bytes = data[record_offset : record_offset + record_size]
            rows[row_id] = decode_record(record_bytes, schema)
        # `copy_table`: rows that are byte-identical duplicates of another row,
        # deduplicated on disk. wago.tools' own CSV export expands these back
        # out, so we do too, to match it row for row.
        for new_id, copied_id in copies:
            if copied_id in rows:
                rows[new_id] = rows[copied_id]
    return rows


#: `ItemSparse`'s column layout for `layout_hash` `0x6FCC3191`, valid for
#: `wow_classic_beta` builds 1.60.1.69876 through (at least) 1.60.1.70094 --
#: see WoWDBDefs' `ItemSparse.dbd`. `ID` itself is excluded: the format
#: stores it out-of-line (`header.flags & 0x04`, "non-inline IDs"), as
#: `read_sparse_rows`'s returned dict key, not as a record field.
ITEM_SPARSE_LAYOUT_HASH = 0x6FCC3191
ITEM_SPARSE_SCHEMA: list[Column] = [
    Column("Description_lang", "str"),
    Column("Display3_lang", "str"),
    Column("Display2_lang", "str"),
    Column("Display1_lang", "str"),
    Column("Display_lang", "str"),
    Column("ExpansionID", "i"),
    Column("DmgVariance", "f"),
    Column("LimitCategory", "i"),
    Column("DurationInInventory", "u"),
    Column("QualityModifier", "f"),
    Column("BagFamily", "u"),
    Column("StartQuestID", "i"),
    Column("LanguageID", "i"),
    Column("ItemRange", "f"),
    Column("StatPercentageOfSocket", "f", count=10),
    Column("StatPercentEditor", "i", count=10),
    Column("StatModifier_bonusStat", "i", count=10),
    Column("Stackable", "i"),
    Column("MaxCount", "i"),
    Column("MinReputation", "i"),
    Column("RequiredAbility", "u"),
    Column("AllowableRace", "i", count=2),
    Column("SellPrice", "u"),
    Column("BuyPrice", "u"),
    Column("VendorStackCount", "u"),
    Column("PriceVariance", "f"),
    Column("PriceRandomValue", "f"),
    Column("Flags", "i", count=5),
    Column("OppositeFactionItemID", "i"),
    Column("ModifiedCraftingReagentItemID", "i"),
    Column("ContentTuningID", "i"),
    Column("PlayerLevelToItemLevelCurveID", "i"),
    Column("ItemLevelOffsetCurveID", "i"),
    Column("ItemLevelOffsetItemLevel", "i"),
    Column("ItemSquishEraID", "i"),
    Column("ItemNameDescriptionID", "u", bits=16),
    Column("RequiredTransmogHoliday", "u", bits=16),
    Column("RequiredHoliday", "u", bits=16),
    Column("Gem_properties", "u", bits=16),
    Column("Socket_match_enchantment_ID", "u", bits=16),
    Column("TotemCategoryID", "u", bits=16),
    Column("InstanceBound", "u", bits=16),
    Column("ZoneBound", "u", count=2, bits=16),
    Column("ItemSet", "u", bits=16),
    Column("LockID", "u", bits=16),
    Column("PageID", "u", bits=16),
    Column("ItemDelay", "u", bits=16),
    Column("MinFactionID", "u", bits=16),
    Column("RequiredSkillRank", "u", bits=16),
    Column("RequiredSkill", "u", bits=16),
    Column("ItemLevel", "u", bits=16),
    Column("AllowableClass", "i", bits=16),
    Column("ArtifactID", "u", bits=8),
    Column("SpellWeight", "u", bits=8),
    Column("SpellWeightCategory", "u", bits=8),
    Column("SocketType", "u", count=3, bits=8),
    Column("SheatheType", "u", bits=8),
    Column("Material", "u", bits=8),
    Column("PageMaterialID", "u", bits=8),
    Column("Bonding", "u", bits=8),
    Column("DamageType", "u", bits=8),
    Column("ContainerSlots", "u", bits=8),
    Column("RequiredPVPMedal", "u", bits=8),
    Column("RequiredPVPRank", "i", bits=8),
    Column("RequiredLevel", "i", bits=8),
    Column("InventoryType", "i", bits=8),
    Column("OverallQualityID", "i", bits=8),
    Column("AmmunitionType", "u", bits=8),
]


def rows_to_csv_dicts(rows: dict[int, dict[str, object]]) -> list[dict[str, str]]:
    """Flatten `read_sparse_rows`'s output into the same shape
    `pipeline.csvio.read_csv` returns for `ItemSparse.csv` -- `ID` plus one
    string-valued column per scalar field, arrays expanded to `Name_0`,
    `Name_1`, ... (matching wago.tools' own CSV column naming), so it is a
    drop-in replacement for `normalize_items`/`normalize_gear`.
    """
    out = []
    for row_id, row in sorted(rows.items()):
        flat: dict[str, str] = {"ID": str(row_id)}
        for name, value in row.items():
            if isinstance(value, list):
                for i, element in enumerate(value):
                    flat[f"{name}_{i}"] = str(element)
            else:
                flat[name] = str(value)
        out.append(flat)
    return out
