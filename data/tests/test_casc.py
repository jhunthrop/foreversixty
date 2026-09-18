from pathlib import Path

import httpx
import pytest

from pipeline.casc import CascError, CascMissing, fetch_casc_file

BUILD = "9.9.9.9"


def transport(
    calls: list[str],
    body: bytes = b"payload",
    status: int = 200,
    headers: dict[str, str] | None = None,
) -> httpx.MockTransport:
    def handler(request: httpx.Request) -> httpx.Response:
        calls.append(str(request.url))
        return httpx.Response(status, content=body, headers=headers or {})

    return httpx.MockTransport(handler)


def client_for(transport_: httpx.MockTransport) -> httpx.Client:
    return httpx.Client(transport=transport_, base_url="https://wago.tools")


def test_every_request_pins_the_build(tmp_path: Path):
    """The same file data id is different content in different builds --
    combatratings.txt is 12,144 bytes on 1.60.1.69893 and 31,614 on Classic
    Era -- and a bare request silently serves whatever wago defaults to."""
    calls: list[str] = []
    fetch_casc_file(
        1391669, BUILD, cache_dir=tmp_path, client=client_for(transport(calls))
    )
    assert calls == [f"https://wago.tools/api/casc/1391669?version={BUILD}"]


def test_the_cache_is_keyed_by_build_as_well_as_by_id(tmp_path: Path):
    calls: list[str] = []
    client = client_for(transport(calls))
    fetch_casc_file(42, "1.0.0.1", cache_dir=tmp_path, client=client, suffix=".blp")
    fetch_casc_file(42, "2.0.0.2", cache_dir=tmp_path, client=client, suffix=".blp")
    assert (tmp_path / "1.0.0.1" / "42.blp").exists()
    assert (tmp_path / "2.0.0.2" / "42.blp").exists()
    assert len(calls) == 2  # two builds, two fetches: the bytes are not the same file


def test_a_second_read_of_the_same_build_serves_from_the_cache(tmp_path: Path):
    calls: list[str] = []
    client = client_for(transport(calls))
    first = fetch_casc_file(42, BUILD, cache_dir=tmp_path, client=client)
    second = fetch_casc_file(42, BUILD, cache_dir=tmp_path, client=client)
    assert first == second == b"payload"
    assert len(calls) == 1


def test_an_unknown_build_is_an_error_naming_the_version(tmp_path: Path):
    """wago answers 400 {"message":"Version not found"} for a build it has
    never seen, which is the typo case worth naming."""
    client = client_for(transport([], body=b'{"message":"Version not found"}', status=400))
    with pytest.raises(CascError, match="9.9.9.9"):
        fetch_casc_file(1391669, BUILD, cache_dir=tmp_path, client=client)


def test_an_empty_body_is_a_missing_file_not_an_empty_one(tmp_path: Path):
    """A file data id the build does not ship answers 200 with nothing at all
    -- six of the seven tables the engine's parser names do exactly that on
    this lineage. Writing a zero-byte file would hand the engine an empty stat
    curve and no error."""
    client = client_for(transport([], body=b""))
    with pytest.raises(CascMissing, match="1391669"):
        fetch_casc_file(1391669, BUILD, cache_dir=tmp_path, client=client)
    assert not (tmp_path / BUILD).exists() or not list((tmp_path / BUILD).iterdir())


def test_a_404_is_the_same_missing_file(tmp_path: Path):
    """Two ways of saying the build does not have it, one exception, so a
    caller that tolerates one tolerates both: icons warn and skip, game tables
    and tree art do not."""
    client = client_for(transport([], body=b'{"errors":"Not found."}', status=404))
    with pytest.raises(CascMissing, match="1391669"):
        fetch_casc_file(1391669, BUILD, cache_dir=tmp_path, client=client)


def test_a_redirect_is_an_error(tmp_path: Path):
    client = client_for(
        transport([], body=b"", status=302, headers={"Location": "https://example.invalid/x"})
    )
    with pytest.raises(CascError):
        fetch_casc_file(1391669, BUILD, cache_dir=tmp_path, client=client)


def test_any_other_failure_stays_an_httpx_error(tmp_path: Path):
    """tests/test_art.py expects httpx.HTTPStatusError from a 500 and that
    test does not change, so the helper raises for status rather than
    wrapping everything."""
    client = client_for(transport([], body=b"boom", status=500))
    with pytest.raises(httpx.HTTPStatusError):
        fetch_casc_file(1391669, BUILD, cache_dir=tmp_path, client=client)
