import hashlib
import struct
import zlib

import pytest

from pipeline.tact import (
    ArchiveIndex,
    EncodingFile,
    TactError,
    blte_decode,
    parse_config,
    parse_keys,
    parse_psv,
    root_content_keys,
    salsa20,
)

KEY_NAME = "D2C028CB7AC14D0B"
KEY = bytes(range(16))


def blte(chunks: list[bytes]) -> bytes:
    """A multi-chunk BLTE stream; each chunk is mode byte + body, decoded size = body
    for 'N' (the tests only need the table to agree with what each chunk decodes to)."""
    sizes = [len(_plain(c)) for c in chunks]
    table = b"".join(
        struct.pack(">II", len(c), size) + hashlib.md5(c).digest()
        for c, size in zip(chunks, sizes, strict=True)
    )
    header_size = 12 + len(table)
    head = b"BLTE" + struct.pack(">I", header_size) + b"\x0f" + len(chunks).to_bytes(3, "big")
    return head + table + b"".join(chunks)


def _plain(chunk: bytes) -> bytes:
    if chunk[:1] == b"N":
        return chunk[1:]
    if chunk[:1] == b"Z":
        return zlib.decompress(chunk[1:])
    return _PLAIN_OF[chunk]


_PLAIN_OF: dict[bytes, bytes] = {}


def encrypted_chunk(inner: bytes, index: int, plain: bytes) -> bytes:
    name = struct.pack("<Q", int(KEY_NAME, 16))
    iv = b"\x01\x02\x03\x04"
    nonce = bytearray(iv.ljust(8, b"\0"))
    for i in range(4):
        nonce[i] ^= (index >> (8 * i)) & 0xFF
    body = b"\x08" + name + b"\x04" + iv + b"S" + salsa20(KEY, bytes(nonce), inner)
    chunk = b"E" + body
    _PLAIN_OF[chunk] = plain
    return chunk


def test_salsa20_matches_the_ecrypt_vector():
    """Set 1, vector 0 of the eSTREAM Salsa20/20 128-bit test vectors."""
    stream = salsa20(bytes.fromhex("80" + "00" * 15), bytes(8), bytes(64))
    assert stream.hex().upper() == (
        "4DFA5E481DA23EA09A31022050859936DA52FCEE218005164F267CB65F5CFD7F"
        "2B4F97E0FF16924A52DF269515110A07F9E460BC65EF95DA58F740B7D1DBB0AA"
    )


def test_blte_decodes_plain_zlib_and_encrypted_chunks():
    secret = b"a section behind a released key"
    data = blte(
        [b"Nhello ", b"Z" + zlib.compress(b"world "), encrypted_chunk(b"N" + secret, 2, secret)]
    )
    assert blte_decode(data, {KEY_NAME: KEY}) == b"hello world " + secret


def test_an_unknown_key_reads_as_zeros_the_size_of_the_chunk():
    secret = b"unreleased"
    data = blte([b"Nab", encrypted_chunk(b"N" + secret, 1, secret)])
    assert blte_decode(data) == b"ab" + bytes(len(secret))


def test_a_single_chunk_stream_and_a_bad_magic():
    assert blte_decode(b"BLTE\0\0\0\0Nplain") == b"plain"
    with pytest.raises(TactError):
        blte_decode(b"NOPE")


def test_configs_and_the_version_table_parse():
    config = parse_config("# Build Configuration\n\nroot = aa\nencoding = bb cc\n")
    assert config == {"root": ["aa"], "encoding": ["bb", "cc"]}
    rows = parse_psv(
        "Region!STRING:0|BuildConfig!HEX:16|VersionsName!String:0\n## seqn = 1\n"
        "us|ad88|1.60.1.70959\n"
    )
    assert rows == [{"Region": "us", "BuildConfig": "ad88", "VersionsName": "1.60.1.70959"}]
    assert parse_keys(f"{KEY_NAME} {KEY.hex()}\nbad line\n") == {KEY_NAME: KEY}


def encoding_file(pairs: dict[bytes, bytes], page_size: int = 1024) -> bytes:
    entries = sorted(pairs.items())
    pages: list[bytes] = []
    first: list[bytes] = []
    page = b""
    for ckey, ekey in entries:
        entry = b"\x01" + (10).to_bytes(5, "big") + ckey + ekey
        if len(page) + len(entry) > page_size:
            pages.append(page.ljust(page_size, b"\0"))
            page = b""
        if not page:
            first.append(ckey)
        page += entry
    pages.append(page.ljust(page_size, b"\0"))
    head = b"EN\x01\x10\x10" + struct.pack(">HHII", page_size // 1024, 4, len(pages), 0)
    head += b"\0" + struct.pack(">I", 0)
    index = b"".join(k + bytes(16) for k in first)
    return head + index + b"".join(pages)


def test_the_encoding_file_maps_content_keys_to_encoding_keys_across_pages():
    pairs = {bytes([i]) * 16: bytes([255 - i]) * 16 for i in range(1, 60)}
    encoding = EncodingFile.parse(encoding_file(pairs))
    assert len(encoding.page_first_keys) > 1
    for ckey, ekey in pairs.items():
        assert encoding.encoding_key(ckey) == ekey
    assert encoding.encoding_key(bytes(16)) is None
    assert encoding.encoding_key(b"\x05" * 15 + b"\x06") is None


def archive_index(entries: dict[bytes, tuple[int, int]]) -> bytes:
    rows = sorted(entries.items())
    block = b"".join(k + struct.pack(">II", size, off) for k, (off, size) in rows)
    block = block.ljust(4096, b"\0")
    toc = rows[-1][0] + bytes(8)
    footer = bytes(8) + bytes([1, 0, 0, 4, 4, 4, 16, 8]) + struct.pack("<I", len(rows)) + bytes(8)
    return block + toc + footer


def test_an_archive_index_finds_an_encoding_key():
    entries = {bytes([i]) * 16: (i * 100, i + 1) for i in range(1, 30)}
    index = ArchiveIndex.parse(archive_index(entries))
    assert index.find(b"\x07" * 16) == (700, 8)
    assert index.find(b"\xff" * 16) is None
    assert index.find(b"\x07" * 15 + b"\x08") is None


def root_file(blocks: list[tuple[int, int, list[tuple[int, bytes]]]]) -> bytes:
    out = b"TSFM" + struct.pack("<IIII", 24, 2, 0, 0) + bytes(4)
    for locale, content, files in blocks:
        out += struct.pack("<IIII", len(files), locale, content, 0) + b"\0"
        previous = -1
        for file_id, _ in files:
            out += struct.pack("<i", file_id - previous - 1)
            previous = file_id
        out += b"".join(ckey for _, ckey in files)
        if not content & 0x10000000:
            out += bytes(8 * len(files))
    return out


def test_the_root_keeps_the_enus_blocks_and_skips_low_violence_and_other_locales():
    root = root_file(
        [
            (0x2, 0, [(10, b"a" * 16), (12, b"b" * 16)]),
            (0x10, 0, [(14, b"x" * 16)]),  # frFR only
            (0x2, 0x80, [(16, b"y" * 16)]),  # low violence
            (0x2, 0x10000000, [(20, b"c" * 16)]),  # no name hashes
        ]
    )
    assert root_content_keys(root) == {10: b"a" * 16, 12: b"b" * 16, 20: b"c" * 16}
    assert root_content_keys(root, {12}) == {12: b"b" * 16}
    with pytest.raises(TactError):
        root_content_keys(b"\0" * 16)
