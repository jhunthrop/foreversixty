import hashlib
import json
from pathlib import Path

import pytest

from pipeline.manifest import newest_build, read_manifest, refresh_manifest, verify, write_manifest


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


def test_manifest_keys_are_paths_relative_to_the_build_dir(tmp_path: Path):
    (tmp_path / "talents").mkdir()
    (tmp_path / "talents" / "warrior.json").write_text("{}\n")
    (tmp_path / "icons").mkdir()
    (tmp_path / "icons" / "a.webp").write_bytes(b"webp")
    (tmp_path / "raw").mkdir()
    (tmp_path / "raw" / "Talent.csv").write_text("ID\n")
    (tmp_path / "classes.json").write_text("[]\n")
    m = write_manifest(tmp_path, build="b", product="p", fetched_at="t")
    assert sorted(m["files"]) == ["classes.json", "icons/a.webp", "talents/warrior.json"]
    assert verify(tmp_path) == []


def test_verify_reports_a_missing_nested_file(tmp_path: Path):
    (tmp_path / "talents").mkdir()
    (tmp_path / "talents" / "warrior.json").write_text("{}\n")
    write_manifest(tmp_path, "b", "p", "t")
    (tmp_path / "talents" / "warrior.json").unlink()
    assert verify(tmp_path) == ["talents/warrior.json"]


def test_refresh_manifest_rehashes_without_losing_provenance(tmp_path: Path):
    (tmp_path / "items.json").write_text("[]\n")
    write_manifest(
        tmp_path, build="1.2.3.4", product="wow_classic_beta", fetched_at="2026-01-01T00:00:00Z"
    )
    (tmp_path / "simdb.bin").write_bytes(b"\x08\x01")
    refreshed = refresh_manifest(tmp_path)
    assert refreshed["build"] == "1.2.3.4"
    assert refreshed["product"] == "wow_classic_beta"
    assert refreshed["fetched_at"] == "2026-01-01T00:00:00Z"
    assert "simdb.bin" in refreshed["files"]
    assert verify(tmp_path) == []


def test_refresh_manifest_needs_a_manifest_to_refresh(tmp_path: Path):
    """simdb and simconst write into a directory normalize has already filled.
    Running one against a directory normalize has never touched is a mistake
    worth naming, not a manifest invented from nothing."""
    with pytest.raises(SystemExit, match="normalize"):
        refresh_manifest(tmp_path)


def test_newest_build_is_the_one_with_the_newest_fetch(tmp_path: Path):
    """A test that means "the build the site runs on" says so through this,
    not through a build id it would have to be edited to rename."""
    for build, fetched_at in (
        ("1.15.9.69722", "2026-09-17T20:30:42.623936Z"),
        ("1.60.1.69893", "2026-09-17T20:31:10.266098Z"),
    ):
        directory = tmp_path / build
        directory.mkdir()
        (directory / "items.json").write_text("[]\n")
        write_manifest(
            directory, build=build, product="wow_classic_beta", fetched_at=fetched_at
        )
    assert newest_build(tmp_path) == "1.60.1.69893"


def test_newest_build_skips_a_build_that_was_never_fetched(tmp_path: Path):
    """`forever-prebeta` is regenerated from a Wowhead snapshot rather than
    fetched from a client, so its manifest has no fetched_at and it is not a
    client build. Its directory sorts last alphabetically but would win a
    naive max() on a null, which is exactly the bug this forbids."""
    fetched = tmp_path / "1.60.1.69893"
    fetched.mkdir()
    (fetched / "items.json").write_text("[]\n")
    write_manifest(
        fetched,
        build="1.60.1.69893",
        product="wow_classic_beta",
        fetched_at="2026-09-17T20:31:10.266098Z",
    )
    synthetic = tmp_path / "forever-prebeta"
    synthetic.mkdir()
    # Written by hand, because the real builds/forever-prebeta/manifest.json is
    # hand-written too: it carries `build`, `product`, `kind`, `derived_from`
    # and `provenance`, and no `fetched_at` key at all.
    (synthetic / "manifest.json").write_text(
        json.dumps({"build": "forever-prebeta", "product": "wow_forever_prebeta"}) + "\n"
    )
    assert newest_build(tmp_path) == "1.60.1.69893"


def test_newest_build_with_nothing_fetched_is_a_clear_error(tmp_path: Path):
    with pytest.raises(SystemExit, match="fetch"):
        newest_build(tmp_path)
