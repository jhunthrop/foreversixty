"""`pipeline.blizzard_cdn` against a small fake CDN: version service, configs,
encoding, root, one archive and its index."""

import hashlib
from pathlib import Path

import httpx
import pytest

from pipeline.blizzard_cdn import (
    USER_AGENT,
    BuildConfigs,
    CdnClient,
    CdnMissing,
    ClientBuild,
    open_build,
    prune,
    resolve_build,
)
from tests.test_tact import archive_index, encoding_file, root_file

VERSION = "1.60.1.70959"
FILE_ID = 1349477
CONTENT = b"WDC5 pretend table bytes"


def md5(data: bytes) -> bytes:
    return hashlib.md5(data).digest()


def blte(data: bytes) -> bytes:
    return b"BLTE\0\0\0\0N" + data


class FakeCdn:
    def __init__(self) -> None:
        file_blte = blte(CONTENT)
        self.file_ekey = md5(file_blte)
        root = root_file([(0x2, 0, [(FILE_ID, md5(CONTENT))])])
        self.root_blte = blte(root)
        root_ckey, self.root_ekey = md5(root), md5(self.root_blte)
        encoding = encoding_file({md5(CONTENT): self.file_ekey, root_ckey: self.root_ekey})
        self.encoding_blte = blte(encoding)
        self.encoding_ekey = md5(self.encoding_blte)
        self.archive = b"padding!" + file_blte
        self.archive_name = "ab" * 16
        self.index = archive_index({self.file_ekey: (8, len(file_blte))})
        self.build_config = (
            f"root = {root_ckey.hex()}\nencoding = {md5(encoding).hex()} "
            f"{self.encoding_ekey.hex()}\nbuild-name = WOW-70959patch1.60.1_Retail\n"
        ).encode()
        self.cdn_config = f"archives = {self.archive_name}\n".encode()
        self.requests: list[str] = []
        self.agents: set[str] = set()

    def files(self) -> dict[str, bytes]:
        def at(kind: str, key: str) -> str:
            return f"/tpr/wow/{kind}/{key[:2]}/{key[2:4]}/{key}"

        return {
            at("config", "bc" * 16): self.build_config,
            at("config", "cc" * 16): self.cdn_config,
            at("data", self.encoding_ekey.hex()): self.encoding_blte,
            at("data", self.root_ekey.hex()): self.root_blte,
            at("data", self.archive_name) + ".index": self.index,
        }

    def handler(self, request: httpx.Request) -> httpx.Response:
        self.requests.append(request.url.path)
        self.agents.add(request.headers.get("user-agent", ""))
        path = request.url.path
        if path == "/wowf/versions":
            return httpx.Response(
                200,
                text="Region!STRING:0|BuildConfig!HEX:16|CDNConfig!HEX:16|VersionsName!String:0\n"
                f"us|{'bc' * 16}|{'cc' * 16}|{VERSION}\n",
            )
        if path == "/wowf/cdns":
            return httpx.Response(
                200,
                text="Name!STRING:0|Path!STRING:0|Hosts!STRING:0|Servers!STRING:0\n"
                "us|tpr/wow|cdn.test|https://cdn.test/?maxhosts=4\n",
            )
        if path.endswith(self.archive_name):
            start, end = request.headers["range"].removeprefix("bytes=").split("-")
            return httpx.Response(206, content=self.archive[int(start) : int(end) + 1])
        if path.endswith("WoW.txt"):
            return httpx.Response(200, text="")
        body = self.files().get(path)
        return httpx.Response(200, content=body) if body is not None else httpx.Response(404)


@pytest.fixture
def cdn(tmp_path: Path) -> tuple[FakeCdn, CdnClient]:
    fake = FakeCdn()
    http = httpx.Client(
        transport=httpx.MockTransport(fake.handler), headers={"User-Agent": USER_AGENT}
    )
    return fake, CdnClient(cache=tmp_path / "cache", http=http, min_interval=0.0)


def test_a_file_is_read_by_id_through_root_encoding_and_an_archive_range(cdn):
    fake, client = cdn
    build = open_build(VERSION, client)
    assert build.read(FILE_ID) == CONTENT
    assert fake.agents == {USER_AGENT}
    with pytest.raises(CdnMissing):
        build.read(42)


def test_a_second_run_of_the_build_downloads_nothing_but_the_version_table(cdn, tmp_path):
    fake, client = cdn
    open_build(VERSION, client).read(FILE_ID)
    fake.requests.clear()
    again = CdnClient(cache=tmp_path / "cache", http=client.http, min_interval=0.0)
    assert open_build(VERSION, again).read(FILE_ID) == CONTENT
    assert fake.requests == ["/wowf/versions"]


def test_a_build_the_service_no_longer_names_comes_from_the_pins(cdn):
    _, client = cdn
    pins = {"1.60.1.70938": {"build_config": "aa", "cdn_config": "bb"}}
    assert resolve_build(client, "1.60.1.70938", pins) == BuildConfigs("1.60.1.70938", "aa", "bb")
    with pytest.raises(SystemExit, match="no pin"):
        resolve_build(client, "1.60.1.1", pins)
    assert resolve_build(client, None, pins).version == VERSION


def test_a_build_config_for_another_build_is_refused(cdn):
    _, client = cdn
    with pytest.raises(SystemExit, match="not 1.60.1.70938"):
        ClientBuild(client, BuildConfigs("1.60.1.70938", "bc" * 16, "cc" * 16))


def test_prune_keeps_only_the_named_builds(tmp_path: Path):
    for name in ("keep", "old"):
        (tmp_path / "builds" / name).mkdir(parents=True)
    assert prune(tmp_path, ["keep"]) == ["old"]
    assert [p.name for p in (tmp_path / "builds").iterdir()] == ["keep"]
