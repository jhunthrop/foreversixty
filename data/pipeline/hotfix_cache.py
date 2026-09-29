"""Reader for the WoW client's local hotfix cache, `DBCache.bin` ("XFTH").

Scope: version 9 caches (the format this build's beta client writes; see the
`night-relics`/hotfix-cache finding below). A hotfix is a row Blizzard's
server pushed to the client at runtime, outside the shipped `.db2` files --
the client keeps it in this cache and applies it over the shipped row for the
rest of that session. `pipeline/db2.py`'s reader only ever sees what the
`.db2` file itself carries; this module is the other half of "what does the
client's copy of `ItemSparse`/`Item` actually contain", for the rows the
build's `ItemSparse.db2`/`Item.db2` never got in the first place (or got with
a zeroed `RequiredLevel` -- see `pipeline/csvio.py`'s completeness gate and
`pipeline.db2`'s module doc for the incident this exists to fix).

Format reference: https://wowdev.wiki/DB2 (`DBCache.bin` redirects to
`ADB#DBCache.bin`), fetched via wowdev.wiki's `action=parse` API since
`action=raw` is disabled there. Its version-9 structs:

    struct dbcache_file_header_v5  (headers stayed this shape through v9)
    {
        uint32_t magic;              // 'XFTH'
        uint32_t version;
        uint32_t build_id;
        uint8_t  verification_hash[32];
    };

    struct dbcache_entry_v9
    {
        uint32_t magic;              // 'XFTH' again, once per record
        int32_t  region_id;
        int32_t  index;              // the push id
        uint32_t unique_id;
        uint32_t table_hash;
        uint32_t record_id;
        uint32_t data_size;
        RecordState status;          // uint8: Valid=1, Delete=2, Invalid=3, NotPublic=4
        uint8_t  padd[3];
        uint8_t  data[data_size];    // only ever present when status == Valid
    };

Confirmed against this build's own bytes, not just the wiki text: decoding
`/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad/hotfix/DBCache.bin`
(4,145,561 bytes; header `version=9`, `build=70009`) with the struct above
puts the *second* record's magic exactly at
`header_size + record_header_size + record[0].data_size` -- byte 164, which
is where the raw file's own second `XFTH` sits -- and doing the same all the
way through the file's 66,540 records never lands on a byte that isn't
`XFTH`, so the field order/widths are right, not just wiki-plausible.
`read_records` re-checks this invariant itself on every real file (it raises
`HotfixCacheError` the moment a record's magic doesn't match), so a
differently-shaped v9 file (or a corrupt one) fails loudly instead of
decoding garbage.

Table hashes (WoWDBDefs' table-hash list, cross-checked against this cache:
every record whose `table_hash` equals one of these decodes cleanly with the
matching schema below): `ItemSparse` = 0x919BE54E, `Item` = 0x50238EC2.

`ItemSparse` hotfix records decode with `pipeline.db2.ITEM_SPARSE_SCHEMA` --
confirmed against this cache's own id 720 ("Brawler Gloves", not in this
build's shipped `ItemSparse.csv` at all): `RequiredLevel` 23, `ItemLevel` 28,
a plausible, complete row. `Item` hotfix records decode with `ITEM_SCHEMA`
below, built from WoWDBDefs' `Item.dbd` `LAYOUT 9A2A4834` (the layout hash
this build's own `Item.db2` reports, confirmed by downloading it from
wago.tools' CASC export) -- and cross-checked field for field against this
build's already-complete `Item.csv` for the same id 720: every one of
`ITEM_SCHEMA`'s 16 columns decodes to the exact value `Item.csv` already has.
Both schemas are read with `pipeline.db2.decode_record`, which is a plain
sequential bit-cursor: it does not care whether the *file* those bytes came
from stores this table with palette/common-data compression (`Item.db2`
does; `header.pallet_data_size` is non-zero) -- a hotfix's row bytes are
always the uncompressed, `field_compression_none`-shaped encoding, because a
single pushed row has no palette to share.

This module never writes the client's own cache file; it only reads one the
controller copied from the beta machine (never touch the Windows game
machine directly -- see the lane brief).
"""

from __future__ import annotations

import struct
from dataclasses import dataclass
from enum import IntEnum
from pathlib import Path

from pipeline.db2 import Column

MAGIC = b"XFTH"

_HEADER_FMT = "<4sII32s"
_HEADER_SIZE = struct.calcsize(_HEADER_FMT)
_RECORD_FMT = "<4siiIIIIB3x"
_RECORD_HEADER_SIZE = struct.calcsize(_RECORD_FMT)

#: The only cache version this module decodes -- see the module doc's struct.
SUPPORTED_VERSION = 9

#: WoWDBDefs' table-hash list, cross-checked against this cache (module doc).
ITEM_SPARSE_TABLE_HASH = 0x919BE54E
ITEM_TABLE_HASH = 0x50238EC2


class HotfixCacheError(Exception):
    """Raised when a file isn't a version-9 `DBCache.bin` this reader supports."""


class HotfixStatus(IntEnum):
    """`RecordState` (wowdev.wiki's `ADB#DBCache.bin`). Only `VALID` records
    carry `data`; the rest are zero-length markers over a `record_id`."""

    VALID = 1
    RECORD_REMOVED = 2
    INVALID = 3
    NOT_PUBLIC = 4


@dataclass(frozen=True)
class HotfixHeader:
    version: int
    build: int
    verification_hash: bytes


@dataclass(frozen=True)
class HotfixRecord:
    region_id: int
    push_id: int
    unique_id: int
    table_hash: int
    record_id: int
    status: HotfixStatus
    data: bytes  # b"" unless status is VALID


def _read_header(data: bytes) -> HotfixHeader:
    if len(data) < _HEADER_SIZE or data[:4] != MAGIC:
        raise HotfixCacheError(f"not a DBCache.bin (magic {data[:4]!r}, want {MAGIC!r})")
    _magic, version, build, verification_hash = struct.unpack_from(_HEADER_FMT, data, 0)
    if version != SUPPORTED_VERSION:
        raise HotfixCacheError(
            f"DBCache.bin version {version}; pipeline.hotfix_cache only reads version "
            f"{SUPPORTED_VERSION} (this build's beta client's own format)"
        )
    return HotfixHeader(version=version, build=build, verification_hash=verification_hash)


def read_cache(
    path: Path, *, table_hashes: set[int] | None = None
) -> tuple[HotfixHeader, list[HotfixRecord]]:
    """Parse every record in `path`.

    `table_hashes`, when given, skips storing a record's `data` for a table
    this caller does not want -- the cache is 66,540+ records covering every
    hotfixed table (spells, quests, broadcast text, ...), and only two of
    those are this pipeline's concern. Records for other tables are still
    counted while walking the file (their `data_size` is needed to find the
    next record), just not kept.

    Raises `HotfixCacheError` the instant a record's magic isn't `XFTH`: that
    means either this file isn't a clean version-9 cache, or (module doc)
    this reader's field order/widths are wrong for it -- either way, decoding
    further would silently produce garbage instead of failing.
    """
    data = path.read_bytes()
    header = _read_header(data)
    records: list[HotfixRecord] = []
    offset = _HEADER_SIZE
    while offset < len(data):
        if data[offset : offset + 4] != MAGIC:
            raise HotfixCacheError(
                f"expected a hotfix record's magic {MAGIC!r} at offset {offset}, got "
                f"{data[offset : offset + 4]!r}; the record header's field order/widths "
                f"do not match this file (see the module doc)"
            )
        (
            _magic,
            region_id,
            push_id,
            unique_id,
            table_hash,
            record_id,
            data_size,
            status_raw,
        ) = struct.unpack_from(_RECORD_FMT, data, offset)
        record_end = offset + _RECORD_HEADER_SIZE + data_size
        if table_hashes is None or table_hash in table_hashes:
            records.append(
                HotfixRecord(
                    region_id=region_id,
                    push_id=push_id,
                    unique_id=unique_id,
                    table_hash=table_hash,
                    record_id=record_id,
                    status=HotfixStatus(status_raw),
                    data=data[offset + _RECORD_HEADER_SIZE : record_end],
                )
            )
        offset = record_end
    return header, records


@dataclass(frozen=True)
class TableHotfixes:
    """Every hotfix one table has in a cache, resolved to one verdict per id.

    A `record_id` can appear more than once (an `INVALID` entry superseded by
    a later `VALID` push, for instance); `decode_table_hotfixes` keeps only
    the *last* one in file order, which is this cache's own write order (each
    record is appended as the client receives it), so "last wins" is
    "newest wins".

    `rows` (only `VALID` ids) is keyed by id, each value the schema-decoded
    record (`pipeline.db2.decode_record`'s output shape). `removed_ids` is
    every id whose latest hotfix is `RECORD_REMOVED` -- `pipeline.hotfix_merge`
    drops these from the shipped table entirely. The four counts are *after*
    the last-wins collapse -- one id whose only surviving verdict is
    `INVALID` (every push for it superseded, or never valid to begin with)
    counts once there, not once per record it took to get there -- so
    `valid_count == len(rows)` and `valid_count + removed_count +
    invalid_count + not_public_count` is the table's total *distinct* hotfixed
    id count, which is what the CLI's printed summary means by "records".
    """

    rows: dict[int, dict[str, object]]
    removed_ids: frozenset[int]
    valid_count: int
    removed_count: int
    invalid_count: int
    not_public_count: int


def decode_table_hotfixes(
    header: HotfixHeader, records: list[HotfixRecord], table_hash: int, schema: list[Column]
) -> TableHotfixes:
    """Resolve `records` for one `table_hash` (module doc's "last wins") and
    decode every `VALID` one with `schema` (`pipeline.db2.decode_record`).

    `header` is unused here (validated by the caller before decoding) but
    required so a caller cannot pass records read against a different
    header's build by mistake without at least having one in hand to check.
    """
    del header  # validated by the caller (e.g. the CLI's build check)
    latest: dict[int, HotfixRecord] = {}
    for record in records:
        if record.table_hash == table_hash:
            latest[record.record_id] = record

    from pipeline.db2 import decode_record

    rows: dict[int, dict[str, object]] = {}
    removed_ids: set[int] = set()
    valid = removed = invalid = not_public = 0
    for record_id, record in latest.items():
        if record.status is HotfixStatus.VALID:
            valid += 1
            rows[record_id] = decode_record(record.data, schema)
        elif record.status is HotfixStatus.RECORD_REMOVED:
            removed += 1
            removed_ids.add(record_id)
        elif record.status is HotfixStatus.INVALID:
            invalid += 1
        else:
            not_public += 1
    return TableHotfixes(
        rows=rows,
        removed_ids=frozenset(removed_ids),
        valid_count=valid,
        removed_count=removed,
        invalid_count=invalid,
        not_public_count=not_public,
    )


#: `Item`'s column layout for `layout_hash` 0x9A2A4834 (WoWDBDefs' `Item.dbd`,
#: `LAYOUT 9A2A4834`, valid for `wow_classic_beta` builds 1.60.1.69876 through
#: (at least) 1.60.1.70094 -- the same range `pipeline.db2.ITEM_SPARSE_SCHEMA`
#: covers). Confirmed against this build's real `Item.db2`
#: (`header.layout_hash == 0x9A2A4834`, downloaded via `pipeline.casc`) and,
#: field for field, against a decoded hotfix record (module doc). `ID` is
#: excluded the same way `ITEM_SPARSE_SCHEMA` excludes it: non-inline,
#: supplied by the caller as the dict key, never a record field.
ITEM_LAYOUT_HASH = 0x9A2A4834
ITEM_SCHEMA: list[Column] = [
    Column("ClassID", "i"),
    Column("SubclassID", "u", bits=8),
    Column("Material", "u", bits=8),
    Column("InventoryType", "i", bits=8),
    Column("SheatheType", "u", bits=8),
    Column("ItemPetFoodID", "i"),
    Column("Sound_override_subclassID", "i", bits=8),
    Column("IconFileDataID", "i"),
    Column("ItemGroupSoundsID", "u"),
    Column("ContentTuningID", "i"),
    Column("ModifiedCraftingReagentItemID", "i"),
    Column("AmmunitionType", "u", bits=8),
    Column("CraftingQualityID", "i"),
    Column("ItemSquishEraID", "i"),
    Column("RecraftReagentCountPercentage", "f"),
    Column("OrderSource", "u", bits=8),
]
