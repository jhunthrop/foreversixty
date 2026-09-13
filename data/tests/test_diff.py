import json
from pathlib import Path

from pipeline.diff import diff_builds, diff_entities


def test_diff_entities_reports_added_removed_changed():
    before = [{"id": 1, "name": "A", "x": 1}, {"id": 2, "name": "B", "x": 2}]
    after = [{"id": 2, "name": "B", "x": 3}, {"id": 3, "name": "C", "x": 0}]
    d = diff_entities(before, after)
    assert d["added"] == [{"id": 3, "name": "C", "x": 0}]
    assert d["removed"] == [{"id": 1, "name": "A", "x": 1}]
    assert d["changed"] == [{"id": 2, "before": before[1], "after": after[0], "fields": ["x"]}]


def test_diff_builds_writes_file(tmp_path: Path):
    for build, zones in [("a", [{"id": 1, "name": "Old"}]), ("b", [{"id": 1, "name": "New"}])]:
        (tmp_path / build).mkdir()
        (tmp_path / build / "zones.json").write_text(json.dumps(zones))
    out = diff_builds("a", "b", root=tmp_path, out=tmp_path / "diffs")
    payload = json.loads(out.read_text())
    assert payload["from"] == "a" and payload["to"] == "b"
    assert payload["entities"]["zones"]["changed"][0]["fields"] == ["name"]
