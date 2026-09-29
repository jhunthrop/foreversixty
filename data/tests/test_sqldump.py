# data/tests/test_sqldump.py
"""Shared quote-aware mysqldump row reader (`pipeline.sqldump`), split out
of `pipeline.classic_quest_levels` for `pipeline.classic_sources` to reuse
-- see that module's own doc for why a naive regex over the raw text is
wrong (a quoted comments/Title field can contain a literal comma, paren
or semicolon)."""

from pipeline import sqldump

SAMPLE_SQL = """
CREATE TABLE `widget` (
  `entry` mediumint unsigned NOT NULL DEFAULT '0',
  `item` mediumint unsigned NOT NULL DEFAULT '0',
  `chance` float NOT NULL DEFAULT '100',
  `comments` varchar(300) DEFAULT ''
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;
INSERT INTO `widget` VALUES (1,100,50.5,'Bring, (a bundle) of things'),(2,200,-10,'It\\'s here');
INSERT INTO `widget` VALUES (3,300,0,'second statement');
"""  # noqa: E501


def test_table_columns_reads_declared_order():
    assert sqldump.table_columns(SAMPLE_SQL, "widget") == ["entry", "item", "chance", "comments"]


def test_table_columns_raises_for_an_unknown_table():
    import pytest

    with pytest.raises(ValueError, match="no CREATE TABLE"):
        sqldump.table_columns(SAMPLE_SQL, "nope")


def test_iter_table_rows_is_not_fooled_by_a_comma_and_parens_in_a_quoted_field():
    rows = list(sqldump.iter_table_rows(SAMPLE_SQL, "widget"))
    assert len(rows) == 3  # not more, from a misparsed comments field


def test_iter_table_rows_spans_more_than_one_insert_statement():
    rows = list(sqldump.iter_table_rows(SAMPLE_SQL, "widget"))
    fields = [sqldump.split_fields(row) for row in rows]
    assert [f[0] for f in fields] == ["1", "2", "3"]


def test_split_fields_handles_an_escaped_quote():
    fields = sqldump.split_fields("2,200,-10,'It\\'s here'")
    assert fields == ["2", "200", "-10", "'It\\'s here'"]


def test_unquote_strips_quotes_and_unescapes():
    assert sqldump.unquote("'It\\'s here'") == "It's here"
    assert sqldump.unquote("NULL") is None
    assert sqldump.unquote("123") == "123"


def test_iter_table_records_zips_fields_up_by_column_name():
    records = list(sqldump.iter_table_records(SAMPLE_SQL, "widget"))
    assert records[0] == {
        "entry": "1", "item": "100", "chance": "50.5",
        "comments": "'Bring, (a bundle) of things'",
    }
    assert records[1]["comments"] == "'It\\'s here'"
    assert records[2] == {
        "entry": "3", "item": "300", "chance": "0", "comments": "'second statement'",
    }


def test_iter_table_records_raises_when_a_row_has_the_wrong_field_count():
    bad_sql = (
        "CREATE TABLE `x` (\n  `a` int,\n  `b` int\n) ENGINE=MyISAM;\n"
        "INSERT INTO `x` VALUES (1,2,3);"
    )
    import pytest

    with pytest.raises(ValueError, match="has 3 fields, schema names 2"):
        list(sqldump.iter_table_records(bad_sql, "x"))
