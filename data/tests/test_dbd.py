import pytest

from pipeline.dbd import DbdError, read_definition

DBD = """COLUMNS
int ID
locstring Name_lang // the name
int<Map::ID> MapID
float Pos?
int Flags

LAYOUT 0000AAAA
BUILD 1.60.1.70938, 1.60.1.70959
$noninline,id$ID<32>
Name_lang
Flags<u16>[2]
$noninline,relation$MapID<32>

LAYOUT 0000AAAA, 0000BBBB
BUILD 1.15.0.1-1.15.9.99999
$id$ID<32>
Name_lang
Pos[3]
"""


def test_the_block_listing_the_build_wins():
    fields = read_definition(DBD, 0xAAAA, "1.60.1.70959")
    assert [f.name for f in fields] == ["ID", "Name_lang", "Flags", "MapID"]
    id_, name, flags, map_id = fields
    assert (id_.is_id, id_.inline) == (True, False)
    assert name.kind == "locstring"
    assert (flags.bits, flags.signed, flags.count) == (16, False, 2)
    assert flags.column_names() == ["Flags_0", "Flags_1"]
    assert (map_id.relation, map_id.inline, map_id.kind) == (True, False, "int")


def test_a_build_range_matches_and_an_unlisted_build_falls_back_to_the_layout():
    assert [f.name for f in read_definition(DBD, 0xBBBB, "1.15.2.5")] == ["ID", "Name_lang", "Pos"]
    # 1.60.1.71000 is listed nowhere: the first block with the layout hash answers.
    assert read_definition(DBD, 0xAAAA, "1.60.1.71000")[2].name == "Flags"


def test_a_layout_no_block_lists_is_an_error():
    with pytest.raises(DbdError, match="0000CCCC"):
        read_definition(DBD, 0xCCCC, "1.60.1.70959")
