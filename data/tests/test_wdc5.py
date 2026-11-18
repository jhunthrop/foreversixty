"""The CDN tables decode to exactly what wago.tools' CSV export served.

The fixtures are 1.60.1.70938 DB2 files from Blizzard's CDN for tables that did
not change since 1.60.1.70532, and wago's 70532 export of each: every
compression the reader handles (immediate, signed, common, pallet, pallet
array), copy tables, the relationship map, strings and floats.
"""

from pathlib import Path

import pytest

from pipeline.dbd import read_definition
from pipeline.tact import blte_decode, parse_keys
from pipeline.wdc5 import format_float, layout_hash, read_rows, table_csv

FIXTURES = Path(__file__).parent / "fixtures" / "cdn"
BUILD = "1.60.1.70938"
GOLDEN = (
    "ArmorLocation",
    "ItemArmorQuality",
    "SpellCastTimes",
    "TalentTab",
    "SkillLineXTraitTree",
    "TraitCurrency",
    "SpellRange",
)


def fields_for(table: str, data: bytes):
    text = (FIXTURES / f"{table}.dbd").read_text(encoding="utf-8")
    return read_definition(text, layout_hash(data), BUILD)


@pytest.mark.parametrize("table", GOLDEN)
def test_a_table_decodes_byte_for_byte_to_wagos_export(table: str):
    data = (FIXTURES / f"{table}.db2").read_bytes()
    expected = (FIXTURES / f"{table}.wago-70532.csv").read_text(encoding="utf-8")
    assert table_csv(data, fields_for(table, data)) == expected


def test_a_section_under_a_released_key_decrypts_and_an_unreleased_one_is_skipped():
    """Map ships one section under D2C028CB7AC14D0B (released, in wowdev/TACTKeys)
    and three rows under keys nobody has published."""
    blte = (FIXTURES / "Map.blte").read_bytes()
    keys = parse_keys((FIXTURES / "keys.txt").read_text())
    without = blte_decode(blte)
    with_key = blte_decode(blte, keys)
    rows_without = read_rows(without, fields_for("Map", without))
    rows_with = read_rows(with_key, fields_for("Map", with_key))
    assert len(rows_with) == len(rows_without) + 5
    assert len({row["ID"] for row in rows_with}) == len(rows_with)


@pytest.mark.parametrize(
    ("value", "text"),
    [
        (0.4000000059604645, "0.40000000596"),
        (393.510009765625, "393.51000976563"),
        (-467.757965087890625, "-467.75796508789"),
        (1.0, "1"),
        (0.0, "0"),
        (-1.0000000150474662e29, "-1.0000000150475E+29"),
        (1e20, "1.0E+20"),
        (1e-05, "1.0E-5"),
    ],
)
def test_floats_print_the_way_the_export_printed_them(value: float, text: str):
    assert format_float(value) == text


def test_an_empty_table_has_no_rows():
    data = (FIXTURES / "ArmorLocation.db2").read_bytes()
    empty = bytearray(data[:204])
    empty[0x88:0x8C] = bytes(4)  # record_count
    assert read_rows(bytes(empty), fields_for("ArmorLocation", data)) == []
