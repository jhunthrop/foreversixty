import struct
from pathlib import Path

import pytest

from pipeline.db2 import (
    ITEM_SPARSE_LAYOUT_HASH,
    ITEM_SPARSE_SCHEMA,
    Column,
    Db2Error,
    decode_record,
    read_sparse_rows,
    rows_to_csv_dicts,
)

_HEADER_FMT = "<4sI128s9I2H7I"
_SECTION_FMT = "<QIIIIIIII"

#: A tiny schema exercising every `Column.kind`/array shape `decode_record`
#: supports, without needing 68 real `ItemSparse` columns.
TINY_SCHEMA = [
    Column("Name_lang", "str"),
    Column("Level", "i", bits=8),
    Column("ItemLevel", "u", bits=16),
    Column("Weight", "f"),
    Column("Flags", "u", count=3, bits=8),
]


def _encode_record(values: dict[str, object]) -> bytes:
    """The inverse of `decode_record`, for `TINY_SCHEMA` only -- builds the
    on-disk bytes for one sparse record."""
    out = b""
    for column in TINY_SCHEMA:
        raw = values[column.name]
        elements = raw if column.count > 1 else [raw]
        for element in elements:
            if column.kind == "str":
                out += element.encode("utf-8") + b"\x00"
            elif column.kind == "f":
                out += struct.pack("<f", element)
            else:
                signed = column.kind == "i"
                out += int(element).to_bytes(column.bits // 8, "little", signed=signed)
    return out


def _build_sparse_file(
    records: dict[int, dict[str, object]],
    *,
    copies: dict[int, int] | None = None,
    encrypted_section: bool = False,
    schema: list[Column] = TINY_SCHEMA,
    layout_hash: int = 0xABCDEF01,
    flags: int = 0x01 | 0x04,
    common_data_size: int = 0,
    pallet_data_size: int = 0,
) -> bytes:
    """Builds a minimal, valid-enough WDC5 sparse file for `read_sparse_rows`
    to decode: header + section headers + one section's records/copy-table/
    offset-map, laid out at whatever offset the section header points to.
    `read_sparse_rows` never reads `field_structure`/`field_storage_info`/
    `pallet_data`/`common_data` (it trusts the caller's `schema` instead), so
    this fixture omits them entirely.
    """
    copies = copies or {}
    record_ids = sorted(records)
    record_bytes = {rid: _encode_record(records[rid]) for rid in record_ids}

    header_size = struct.calcsize(_HEADER_FMT)
    section_count = 2 if encrypted_section else 1
    section_size = struct.calcsize(_SECTION_FMT)
    section0_offset = header_size + section_count * section_size

    variable_record_data = b""
    entries = []
    for rid in record_ids:
        entries.append((section0_offset + len(variable_record_data), len(record_bytes[rid])))
        variable_record_data += record_bytes[rid]
    offset_records_end = section0_offset + len(variable_record_data)

    copy_table = b"".join(
        struct.pack("<II", new_id, copied_id) for new_id, copied_id in copies.items()
    )
    offset_map = b"".join(struct.pack("<IH", offset, size) for offset, size in entries)
    offset_map_id_list = struct.pack(f"<{len(record_ids)}I", *record_ids)

    section0 = struct.pack(
        _SECTION_FMT,
        0,  # tact_key_hash (unencrypted)
        section0_offset,  # file_offset
        len(record_ids),  # record_count
        0,  # string_table_size (inline strings -- unused)
        offset_records_end,
        0,  # id_list_size (read_sparse_rows doesn't use id_list)
        0,  # relationship_data_size
        len(record_ids),  # offset_map_id_count
        len(copies),  # copy_table_count
    )
    sections = section0
    if encrypted_section:
        sections += struct.pack(_SECTION_FMT, 0xDEAD, 0, 1, 0, 0, 0, 0, 0, 0)

    header = struct.pack(
        _HEADER_FMT,
        b"WDC5",
        5,  # versionNum
        b"test-fixture".ljust(128, b"\x00"),
        len(record_ids) + (1 if encrypted_section else 0),  # record_count
        len(schema),  # field_count
        0,  # record_size (unused for sparse tables)
        0,  # string_table_size
        0xC0FFEE,  # table_hash
        layout_hash,
        min(record_ids),
        max(record_ids),
        1,  # locale
        flags,
        0,  # id_index
        len(schema),  # total_field_count
        0,  # bitpacked_data_offset
        0,  # lookup_column_count
        0,  # field_storage_info_size
        common_data_size,
        pallet_data_size,
        section_count,
    )

    return (
        header
        + sections
        + variable_record_data
        + copy_table
        + offset_map
        + offset_map_id_list
    )


def _write(tmp_path: Path, data: bytes) -> Path:
    path = tmp_path / "Test.db2"
    path.write_bytes(data)
    return path


def test_decode_record_reads_strings_ints_floats_and_arrays():
    record = _encode_record(
        {
            "Name_lang": "Acidproof Cloak",
            "Level": 9,
            "ItemLevel": 18,
            "Weight": 0.5,
            "Flags": [1, 0, 255],
        }
    )
    row = decode_record(record, TINY_SCHEMA)
    assert row == {
        "Name_lang": "Acidproof Cloak",
        "Level": 9,
        "ItemLevel": 18,
        "Weight": 0.5,
        "Flags": [1, 0, 255],
    }


def test_decode_record_reads_negative_signed_field():
    record = _encode_record(
        {"Name_lang": "", "Level": -1, "ItemLevel": 0, "Weight": 0.0, "Flags": [0, 0, 0]}
    )
    row = decode_record(record, TINY_SCHEMA)
    assert row["Level"] == -1


def test_read_sparse_rows_decodes_every_unencrypted_record(tmp_path):
    records = {
        3582: {
            "Name_lang": "Acidproof Cloak",
            "Level": 9,
            "ItemLevel": 18,
            "Weight": 1.0,
            "Flags": [0, 0, 0],
        },
        22657: {
            "Name_lang": "Amulet of the Dawn",
            "Level": 55,
            "ItemLevel": 60,
            "Weight": 2.5,
            "Flags": [1, 2, 3],
        },
    }
    data = _build_sparse_file(records)
    rows = read_sparse_rows(_write(tmp_path, data), TINY_SCHEMA)

    assert set(rows) == {3582, 22657}
    assert rows[3582]["Name_lang"] == "Acidproof Cloak"
    assert rows[3582]["Level"] == 9
    assert rows[22657]["ItemLevel"] == 60
    assert rows[22657]["Flags"] == [1, 2, 3]


def test_read_sparse_rows_expands_copy_table_rows(tmp_path):
    original = {
        "Name_lang": "Original",
        "Level": 5,
        "ItemLevel": 10,
        "Weight": 1.0,
        "Flags": [0, 0, 0],
    }
    records = {1: original}
    data = _build_sparse_file(records, copies={2: 1, 3: 1})
    rows = read_sparse_rows(_write(tmp_path, data), TINY_SCHEMA)

    assert set(rows) == {1, 2, 3}
    assert rows[2] == rows[1]
    assert rows[3] == rows[1]


def test_read_sparse_rows_skips_encrypted_sections(tmp_path):
    records = {
        1: {"Name_lang": "Visible", "Level": 1, "ItemLevel": 1, "Weight": 0.0, "Flags": [0, 0, 0]}
    }
    data = _build_sparse_file(records, encrypted_section=True)
    rows = read_sparse_rows(_write(tmp_path, data), TINY_SCHEMA)

    # header.record_count includes the encrypted section's 1 row, but it's
    # unreadable without a TACT key, so it must not appear in the result.
    assert set(rows) == {1}


def test_read_sparse_rows_validates_layout_hash(tmp_path):
    records = {1: {"Name_lang": "X", "Level": 1, "ItemLevel": 1, "Weight": 0.0, "Flags": [0, 0, 0]}}
    data = _build_sparse_file(records, layout_hash=0x11111111)
    path = _write(tmp_path, data)

    with pytest.raises(Db2Error, match="layout_hash"):
        read_sparse_rows(path, TINY_SCHEMA, expected_layout_hash=0x22222222)


def test_read_sparse_rows_rejects_schema_field_count_mismatch(tmp_path):
    records = {1: {"Name_lang": "X", "Level": 1, "ItemLevel": 1, "Weight": 0.0, "Flags": [0, 0, 0]}}
    data = _build_sparse_file(records)
    path = _write(tmp_path, data)

    with pytest.raises(Db2Error, match="fields"):
        read_sparse_rows(path, TINY_SCHEMA[:-1])


def test_read_sparse_rows_rejects_non_sparse_flags(tmp_path):
    records = {1: {"Name_lang": "X", "Level": 1, "ItemLevel": 1, "Weight": 0.0, "Flags": [0, 0, 0]}}
    data = _build_sparse_file(records, flags=0x04)  # no offset-map bit
    path = _write(tmp_path, data)

    with pytest.raises(Db2Error, match="offset-map"):
        read_sparse_rows(path, TINY_SCHEMA)


def test_read_sparse_rows_rejects_common_or_pallet_data(tmp_path):
    records = {1: {"Name_lang": "X", "Level": 1, "ItemLevel": 1, "Weight": 0.0, "Flags": [0, 0, 0]}}
    data = _build_sparse_file(records, common_data_size=8)
    path = _write(tmp_path, data)

    with pytest.raises(Db2Error, match="common_data"):
        read_sparse_rows(path, TINY_SCHEMA)


def test_decode_record_rejects_unaligned_string_field():
    unaligned_schema = [Column("Flag", "u", bits=4), Column("Name_lang", "str")]
    with pytest.raises(Db2Error, match="byte-aligned"):
        decode_record(b"\x00\x00", unaligned_schema)


def test_decode_record_rejects_unknown_column_kind():
    bad_schema = [Column("Mystery", "money")]
    with pytest.raises(Db2Error, match="unknown column kind"):
        decode_record(b"\x00", bad_schema)


def test_read_sparse_rows_rejects_bad_magic(tmp_path):
    path = tmp_path / "NotADb2.db2"
    path.write_bytes(b"NOPE" + b"\x00" * 200)

    with pytest.raises(Db2Error, match="WDC5"):
        read_sparse_rows(path, TINY_SCHEMA)


def test_rows_to_csv_dicts_flattens_arrays_and_stringifies():
    rows = {
        1: {"Name_lang": "Alpha", "Level": 5, "ItemLevel": 10, "Weight": 1.5, "Flags": [0, 1, 2]},
    }
    csv_dicts = rows_to_csv_dicts(rows)
    assert csv_dicts == [
        {
            "ID": "1",
            "Name_lang": "Alpha",
            "Level": "5",
            "ItemLevel": "10",
            "Weight": "1.5",
            "Flags_0": "0",
            "Flags_1": "1",
            "Flags_2": "2",
        }
    ]


def test_rows_to_csv_dicts_sorts_by_id():
    rows = {
        3: {"Name_lang": "C", "Level": 0, "ItemLevel": 0, "Weight": 0.0, "Flags": [0, 0, 0]},
        1: {"Name_lang": "A", "Level": 0, "ItemLevel": 0, "Weight": 0.0, "Flags": [0, 0, 0]},
    }
    csv_dicts = rows_to_csv_dicts(rows)
    assert [d["ID"] for d in csv_dicts] == ["1", "3"]


def test_item_sparse_schema_matches_build_1_60_1_70009_layout_hash():
    # This is the layout hash `pipeline.db2`'s module docstring cites from
    # WoWDBDefs' `ItemSparse.dbd` for build 1.60.1.70009 (and every other
    # `wow_classic_beta` 1.60.1 build wago.tools currently serves). It is
    # only a constant-value regression test -- the real cross-check against
    # the build's actual raw `ItemSparse.db2` is manual (see the lane
    # report), since that file is too large to fix into this repo.
    assert ITEM_SPARSE_LAYOUT_HASH == 0x6FCC3191
    assert len(ITEM_SPARSE_SCHEMA) == 68
    assert ITEM_SPARSE_SCHEMA[0].name == "Description_lang"
    assert ITEM_SPARSE_SCHEMA[-1].name == "AmmunitionType"
