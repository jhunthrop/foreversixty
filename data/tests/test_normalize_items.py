from pathlib import Path

from pipeline.csvio import read_csv
from pipeline.normalize import write_json
from pipeline.normalize.items import normalize_items

HERE = Path(__file__).parent


def test_items_match_golden(tmp_path: Path):
    items = normalize_items(
        read_csv(HERE / "fixtures/ItemSparse.csv"),
        read_csv(HERE / "fixtures/Item.csv"),
    )
    out = tmp_path / "items.json"
    write_json(items, out)
    assert out.read_text() == (HERE / "golden/items.json").read_text()


def test_item_without_item_row_is_skipped():
    sparse = [
        {
            "ID": "1",
            "Display_lang": "Orphan",
            "OverallQualityID": "1",
            "ItemLevel": "1",
            "RequiredLevel": "1",
            "InventoryType": "0",
        }
    ]
    assert normalize_items(sparse, []) == []
