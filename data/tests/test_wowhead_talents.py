import json
from pathlib import Path

import httpx
import pytest

from pipeline import wowhead_talents
from pipeline.wowhead_talents import WowheadTalentsError, fetch_wowhead_talents, unwrap_talents

PAYLOAD = {"trees": {"161": {"id": 161}}, "talents": {"161": {"1": {"id": 1}}}}
ENVELOPE = (
    'WH.setPageData("wow.talents.classicplus", '
    + json.dumps(PAYLOAD)
    + ');\nWH.setPageData("other", {"x": {"y": 1}});\nvar tail = {"trees": 1};\n'
)


def test_unwrap_takes_the_first_object_and_ignores_trailing_statements():
    assert unwrap_talents(ENVELOPE) == PAYLOAD


def test_unwrap_refuses_a_body_without_an_object():
    with pytest.raises(WowheadTalentsError, match="no JSON object"):
        unwrap_talents("WH.setPageData('x');")


def test_unwrap_refuses_a_payload_without_talents():
    with pytest.raises(WowheadTalentsError, match="lacks talents"):
        unwrap_talents('WH.setPageData("x", {"trees": {"1": {}}});')


def _client(handler) -> httpx.Client:
    return httpx.Client(transport=httpx.MockTransport(handler))


def test_fetch_writes_the_payload_with_provenance(tmp_path: Path):
    seen: list[httpx.Request] = []

    def handler(request: httpx.Request) -> httpx.Response:
        seen.append(request)
        return httpx.Response(200, text=ENVELOPE)

    path = fetch_wowhead_talents("1.60.1.1", root=tmp_path, client=_client(handler))
    assert path == tmp_path / "1.60.1.1" / "raw" / "wowhead-talents.json"
    saved = json.loads(path.read_text())
    meta = saved.pop("_meta")
    assert saved == PAYLOAD
    assert meta["url"].startswith("https://nether.wowhead.com/forever/data/talents-classic")
    assert "dv=100" in meta["url"]
    assert meta["fetched_at"]
    assert len(seen) == 1


def test_fetch_retries_a_server_error_then_succeeds(
    tmp_path: Path, monkeypatch: pytest.MonkeyPatch
):
    monkeypatch.setattr(wowhead_talents.time, "sleep", lambda _seconds: None)
    answers = iter([httpx.Response(503), httpx.Response(200, text=ENVELOPE)])

    path = fetch_wowhead_talents("b", root=tmp_path, client=_client(lambda _request: next(answers)))
    assert json.loads(path.read_text())["trees"] == PAYLOAD["trees"]


def test_fetch_gives_up_after_the_attempt_limit_and_writes_nothing(
    tmp_path: Path, monkeypatch: pytest.MonkeyPatch
):
    monkeypatch.setattr(wowhead_talents.time, "sleep", lambda _seconds: None)
    with pytest.raises(httpx.HTTPStatusError):
        fetch_wowhead_talents("b", root=tmp_path, client=_client(lambda _r: httpx.Response(503)))
    assert not (tmp_path / "b").exists()
