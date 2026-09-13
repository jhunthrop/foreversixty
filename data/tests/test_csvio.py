from pathlib import Path

from pipeline.csvio import read_csv

FIXTURES = Path(__file__).parent / "fixtures"


def test_read_csv_returns_rows_keyed_by_header():
    rows = read_csv(FIXTURES / "Tiny.csv")
    assert rows == [
        {"ID": "1", "Name_lang": "Alpha", "Flags": "0"},
        {"ID": "2", "Name_lang": "Beta, with comma", "Flags": "4"},
    ]
