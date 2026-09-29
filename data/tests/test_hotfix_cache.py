import os
import struct
from pathlib import Path

import pytest

from pipeline.db2 import Column
from pipeline.hotfix_cache import (
    ITEM_LAYOUT_HASH,
    ITEM_SCHEMA,
    ITEM_SPARSE_TABLE_HASH,
    ITEM_TABLE_HASH,
    HotfixCacheError,
    HotfixStatus,
    decode_table_hotfixes,
    read_cache,
)

_HEADER_FMT = "<4sII32s"
_RECORD_FMT = "<4siiIIIIB3x"

#: A tiny schema for the synthetic fixtures below -- decoding real
#: ItemSparse/Item hotfix rows is covered by the real-cache test at the
#: bottom, which needs no schema knowledge to just count records.
TINY_SCHEMA: list[Column] = [
    Column("Name_lang", "str"),
    Column("Level", "i", bits=8),
]


def _encode_tiny(name: str, level: int) -> bytes:
    return name.encode("utf-8") + b"\x00" + int(level).to_bytes(1, "little", signed=True)


def _record(
    table_hash: int,
    record_id: int,
    status: HotfixStatus,
    data: bytes = b"",
    *,
    region_id: int = 0,
    push_id: int = 1,
    unique_id: int = 1,
) -> bytes:
    return (
        struct.pack(
            _RECORD_FMT,
            b"XFTH",
            region_id,
            push_id,
            unique_id,
            table_hash,
            record_id,
            len(data),
            int(status),
        )
        + data
    )


def _cache(records: bytes, *, version: int = 9, build: int = 70009) -> bytes:
    return struct.pack(_HEADER_FMT, b"XFTH", version, build, b"\x00" * 32) + records


def _write(tmp_path: Path, data: bytes) -> Path:
    path = tmp_path / "DBCache.bin"
    path.write_bytes(data)
    return path


def test_read_cache_decodes_header_and_every_record(tmp_path: Path):
    valid = _record(0xAAAA, 1, HotfixStatus.VALID, _encode_tiny("Alpha", 9), push_id=5)
    removed = _record(0xBBBB, 2, HotfixStatus.RECORD_REMOVED)
    path = _write(tmp_path, _cache(valid + removed, build=70009))

    header, records = read_cache(path)

    assert header.version == 9
    assert header.build == 70009
    assert [r.table_hash for r in records] == [0xAAAA, 0xBBBB]
    assert records[0].status is HotfixStatus.VALID
    assert records[0].push_id == 5
    assert records[0].data == _encode_tiny("Alpha", 9)
    assert records[1].status is HotfixStatus.RECORD_REMOVED
    assert records[1].data == b""  # module doc: only VALID ever carries data


def test_read_cache_filters_by_table_hash(tmp_path: Path):
    wanted = _record(0xAAAA, 1, HotfixStatus.VALID, _encode_tiny("Alpha", 9))
    other = _record(0xBBBB, 2, HotfixStatus.VALID, _encode_tiny("Beta", 1))
    path = _write(tmp_path, _cache(wanted + other))

    _, records = read_cache(path, table_hashes={0xAAAA})

    assert [r.table_hash for r in records] == [0xAAAA]


def test_read_cache_rejects_an_unsupported_version(tmp_path: Path):
    path = _write(tmp_path, _cache(b"", version=8))
    with pytest.raises(HotfixCacheError, match="version 8"):
        read_cache(path)


def test_read_cache_rejects_a_file_with_the_wrong_magic(tmp_path: Path):
    path = tmp_path / "not-a-cache.bin"
    path.write_bytes(b"NOPE" + b"\x00" * 40)
    with pytest.raises(HotfixCacheError, match="not a DBCache.bin"):
        read_cache(path)


def test_read_cache_raises_when_a_record_desyncs(tmp_path: Path):
    """The exact check the module doc leans on: the second record's magic
    must land exactly where the header's field order/widths say it does. A
    record whose header lies about its own `data_size` desyncs every record
    after it, and that has to fail loudly, not decode whatever bytes it
    lands on as if they were a real record header.
    """
    first = bytearray(_record(0xAAAA, 1, HotfixStatus.VALID, _encode_tiny("Alpha", 9)))
    data_size_offset = struct.calcsize("<4siiIII")  # right before the real data_size field
    first[data_size_offset] = 1  # claim 1 byte of data instead of the real 6
    second = _record(0xBBBB, 2, HotfixStatus.VALID, _encode_tiny("Beta", 1))
    path = _write(tmp_path, _cache(bytes(first) + second))

    with pytest.raises(HotfixCacheError, match="expected a hotfix record's magic"):
        read_cache(path)


def test_decode_table_hotfixes_keeps_only_the_last_write_per_id(tmp_path: Path):
    superseded = _record(0xAAAA, 42, HotfixStatus.INVALID, push_id=1)
    winner = _record(0xAAAA, 42, HotfixStatus.VALID, _encode_tiny("Winner", 5), push_id=2)
    path = _write(tmp_path, _cache(superseded + winner))
    header, records = read_cache(path)

    hotfixes = decode_table_hotfixes(header, records, 0xAAAA, TINY_SCHEMA)

    assert hotfixes.rows == {42: {"Name_lang": "Winner", "Level": 5}}
    assert hotfixes.valid_count == 1
    assert hotfixes.invalid_count == 0  # the INVALID push for id 42 was superseded


def test_decode_table_hotfixes_a_later_removal_wins_over_an_earlier_valid(tmp_path: Path):
    added = _record(0xAAAA, 42, HotfixStatus.VALID, _encode_tiny("Gone", 1), push_id=1)
    removed = _record(0xAAAA, 42, HotfixStatus.RECORD_REMOVED, push_id=2)
    path = _write(tmp_path, _cache(added + removed))
    header, records = read_cache(path)

    hotfixes = decode_table_hotfixes(header, records, 0xAAAA, TINY_SCHEMA)

    assert hotfixes.rows == {}
    assert hotfixes.removed_ids == frozenset({42})
    assert hotfixes.removed_count == 1


def test_decode_table_hotfixes_tallies_every_status(tmp_path: Path):
    records_bytes = (
        _record(0xAAAA, 1, HotfixStatus.VALID, _encode_tiny("A", 1))
        + _record(0xAAAA, 2, HotfixStatus.RECORD_REMOVED)
        + _record(0xAAAA, 3, HotfixStatus.INVALID)
        + _record(0xAAAA, 4, HotfixStatus.NOT_PUBLIC)
    )
    path = _write(tmp_path, _cache(records_bytes))
    header, records = read_cache(path)

    hotfixes = decode_table_hotfixes(header, records, 0xAAAA, TINY_SCHEMA)

    assert (
        hotfixes.valid_count,
        hotfixes.removed_count,
        hotfixes.invalid_count,
        hotfixes.not_public_count,
    ) == (1, 1, 1, 1)


def test_item_schema_is_16_columns_matching_this_builds_field_count():
    # header.total_field_count for the real Item.db2 (layout_hash
    # ITEM_LAYOUT_HASH) is 16 -- see the module doc's cross-check against the
    # downloaded file. A schema edit that drops or adds a column without
    # updating this would decode every later field at the wrong bit offset.
    assert len(ITEM_SCHEMA) == 16
    assert ITEM_LAYOUT_HASH == 0x9A2A4834


@pytest.mark.skipif(
    not os.environ.get("FOREVER_HOTFIX_CACHE"),
    reason="set FOREVER_HOTFIX_CACHE=<path to DBCache.bin> to run against the real client cache",
)
def test_real_cache_decodes_itemsparse_and_item_without_error():
    from pipeline.db2 import ITEM_SPARSE_SCHEMA

    path = Path(os.environ["FOREVER_HOTFIX_CACHE"])
    header, records = read_cache(path, table_hashes={ITEM_SPARSE_TABLE_HASH, ITEM_TABLE_HASH})

    assert header.version == 9
    sparse = decode_table_hotfixes(header, records, ITEM_SPARSE_TABLE_HASH, ITEM_SPARSE_SCHEMA)
    item = decode_table_hotfixes(header, records, ITEM_TABLE_HASH, ITEM_SCHEMA)

    # Every id the client hotfixed ItemSparse for, it also hotfixed Item for
    # (see the lane report) -- both directions, so a schema/table-hash typo
    # that silently decodes zero rows for one table fails this either way.
    assert sparse.valid_count > 0
    assert item.valid_count > 0
    assert set(sparse.rows) <= set(item.rows) or set(item.rows) <= set(sparse.rows)
