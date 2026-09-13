import hashlib
import json
from pathlib import Path

from pipeline.manifest import read_manifest, verify, write_manifest


def test_manifest_lists_files_with_sha256(tmp_path: Path):
    (tmp_path / "zones.json").write_text("[]\n")
    m = write_manifest(
        tmp_path,
        build="1.15.7.61582",
        product="wow_classic_era",
        fetched_at="2026-09-14T00:00:00Z",
    )
    assert m["build"] == "1.15.7.61582"
    assert m["files"] == {"zones.json": hashlib.sha256(b"[]\n").hexdigest()}
    assert json.loads((tmp_path / "manifest.json").read_text()) == m
    assert read_manifest(tmp_path) == m


def test_verify_reports_changed_files(tmp_path: Path):
    (tmp_path / "zones.json").write_text("[]\n")
    write_manifest(tmp_path, "b", "p", "t")
    (tmp_path / "zones.json").write_text("[1]\n")
    assert verify(tmp_path) == ["zones.json"]
