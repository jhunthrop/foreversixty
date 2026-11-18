from pathlib import Path

import pytest

from pipeline.casc import CascMissing, CascMissingReason, fetch_casc_file
from tests.cdn_fakes import FakeSource

BUILD = "9.9.9.9"


def test_the_cache_is_keyed_by_build_as_well_as_by_id(tmp_path: Path):
    """The same file data id is different content in different builds --
    combatratings.txt is 12,144 bytes on 1.60.1.69893 and 31,614 on Classic Era."""
    source = FakeSource({42: b"payload"})
    fetch_casc_file(42, "1.0.0.1", cache_dir=tmp_path, source=source, suffix=".blp")
    fetch_casc_file(42, "2.0.0.2", cache_dir=tmp_path, source=source, suffix=".blp")
    assert (tmp_path / "1.0.0.1" / "42.blp").exists()
    assert (tmp_path / "2.0.0.2" / "42.blp").exists()
    assert source.calls == [42, 42]  # two builds, two reads: the bytes are not the same file


def test_a_second_read_of_the_same_build_serves_from_the_cache(tmp_path: Path):
    source = FakeSource({42: b"payload"})
    first = fetch_casc_file(42, BUILD, cache_dir=tmp_path, source=source)
    second = fetch_casc_file(42, BUILD, cache_dir=tmp_path, source=source)
    assert first == second == b"payload"
    assert source.calls == [42]


def test_an_empty_file_is_a_missing_file_not_an_empty_one(tmp_path: Path):
    """Writing a zero-byte file would hand the engine an empty stat curve and no error."""
    with pytest.raises(CascMissing, match="1391669") as excinfo:
        fetch_casc_file(1391669, BUILD, cache_dir=tmp_path, source=FakeSource({1391669: b""}))
    assert excinfo.value.reason is CascMissingReason.EMPTY
    assert not (tmp_path / BUILD).exists() or not list((tmp_path / BUILD).iterdir())


def test_a_file_the_build_lacks_is_the_same_missing_file(tmp_path: Path):
    """Two ways of saying the build does not have it, one exception, so a
    caller that tolerates one tolerates both: icons warn and skip, game tables
    and tree art do not."""
    with pytest.raises(CascMissing, match="1391669") as excinfo:
        fetch_casc_file(1391669, BUILD, cache_dir=tmp_path, source=FakeSource())
    assert excinfo.value.reason is CascMissingReason.NOT_FOUND


def test_a_cdn_failure_propagates(tmp_path: Path):
    """Anything but a missing file is not swallowed: the run fails."""

    def boom() -> bytes:
        raise RuntimeError("CDN hosts answered [500]")

    with pytest.raises(RuntimeError, match="500"):
        fetch_casc_file(7, BUILD, cache_dir=tmp_path, source=FakeSource({7: boom}))
