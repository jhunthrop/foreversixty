"""Blizzard's TACT/CASC file formats, read from bytes: no network here.

`pipeline.blizzard_cdn` downloads; this module only parses. The layouts are
wowdev.wiki's (TACT, BLTE, Encoding, Root, CDN archive indexes), cross-checked
against wowdev/TACTSharp, which reads the same files.

* `parse_config`: a build or CDN config (`key = value value ...`).
* `blte_decode`: a BLTE stream ('N' plain and 'Z' zlib chunks; an 'E' chunk
  whose key we do not have decodes to zeros, the way TACTSharp and wago.tools
  leave a section they cannot decrypt).
* `EncodingFile`: content key -> encoding key, by the page index, so only the
  pages asked about are parsed.
* `root_content_keys`: file data id -> content key from the WoW root (MFST).
* `ArchiveIndex`: encoding key -> (offset, size) inside one CDN archive.
"""

from __future__ import annotations

import bisect
import struct
import zlib
from collections.abc import Mapping
from dataclasses import dataclass

BLTE_MAGIC = b"BLTE"
ENCODING_MAGIC = b"EN"
ROOT_MAGIC = b"TSFM"
#: wowdev.wiki's root locale flag for enUS, and the set meaning "every WoW locale".
LOCALE_ENUS = 0x2
LOCALE_ALL_WOW = 0x1F3F6
#: Root content flags: low-violence variants are skipped, and a block flagged
#: NO_NAMES carries no name hashes.
CONTENT_LOW_VIOLENCE = 0x80
CONTENT_NO_NAMES = 0x10000000
INDEX_FOOTER_SIZE = 20


class TactError(ValueError):
    """A TACT file is not what its format says it should be."""


def parse_config(text: str) -> dict[str, list[str]]:
    """A build or CDN config: each `key = a b c` line as key -> [a, b, c]."""
    config: dict[str, list[str]] = {}
    for line in text.splitlines():
        line = line.strip()
        if not line or line.startswith("#") or "=" not in line:
            continue
        key, _, value = line.partition("=")
        config[key.strip()] = value.split()
    return config


def parse_psv(text: str) -> list[dict[str, str]]:
    """Blizzard's version-service table (`Name!TYPE:n|...` header, `|` rows)."""
    rows: list[dict[str, str]] = []
    header: list[str] | None = None
    for line in text.splitlines():
        if not line.strip() or line.startswith("##"):
            continue
        cells = line.split("|")
        if header is None:
            header = [cell.split("!", 1)[0] for cell in cells]
            continue
        rows.append(dict(zip(header, cells, strict=False)))
    return rows


_MASK32 = 0xFFFFFFFF
_SIGMA16 = struct.unpack("<4I", b"expand 16-byte k")


def _rotl(value: int, shift: int) -> int:
    return ((value << shift) | (value >> (32 - shift))) & _MASK32


def _salsa20_block(state: list[int]) -> bytes:
    x = list(state)
    for _ in range(10):
        for a, b, c, d in (
            (0, 4, 8, 12), (5, 9, 13, 1), (10, 14, 2, 6), (15, 3, 7, 11),
            (0, 1, 2, 3), (5, 6, 7, 4), (10, 11, 8, 9), (15, 12, 13, 14),
        ):  # fmt: skip
            x[b] ^= _rotl((x[a] + x[d]) & _MASK32, 7)
            x[c] ^= _rotl((x[b] + x[a]) & _MASK32, 9)
            x[d] ^= _rotl((x[c] + x[b]) & _MASK32, 13)
            x[a] ^= _rotl((x[d] + x[c]) & _MASK32, 18)
    return struct.pack("<16I", *((x[i] + state[i]) & _MASK32 for i in range(16)))


def salsa20(key: bytes, nonce: bytes, data: bytes) -> bytes:
    """Salsa20/20 with a 16-byte key (TACT's 'S' encryption), encrypt = decrypt."""
    k = struct.unpack("<4I", key)
    n = struct.unpack("<2I", nonce)
    out = bytearray(len(data))
    for block in range(-(-len(data) // 64)):
        state = [
            _SIGMA16[0], *k, _SIGMA16[1], *n, block & _MASK32, block >> 32,
            _SIGMA16[2], *k, _SIGMA16[3],
        ]  # fmt: skip
        stream = _salsa20_block(state)
        start = block * 64
        chunk = data[start : start + 64]
        out[start : start + len(chunk)] = bytes(a ^ b for a, b in zip(chunk, stream, strict=False))
    return bytes(out)


def _decrypt_chunk(body: bytes, index: int, keys: Mapping[str, bytes]) -> bytes | None:
    """An 'E' chunk's inner chunk, or None when its key is not known."""
    name_size = body[0]
    key_name = f"{struct.unpack_from('<Q', body, 1)[0]:016X}"
    iv_size = body[1 + name_size]
    iv = bytearray(body[2 + name_size : 2 + name_size + iv_size].ljust(8, b"\0"))
    kind = body[2 + name_size + iv_size : 3 + name_size + iv_size]
    key = keys.get(key_name)
    if key is None:
        return None
    if kind != b"S":
        raise TactError(f"BLTE encryption {kind!r} is not supported")
    for i in range(4):
        iv[i] ^= (index >> (8 * i)) & 0xFF
    return salsa20(key, bytes(iv), body[3 + name_size + iv_size :])


def _decode_chunk(chunk: bytes, size: int, index: int, keys: Mapping[str, bytes]) -> bytes:
    mode, body = chunk[:1], chunk[1:]
    if mode == b"N":
        return body
    if mode == b"Z":
        return zlib.decompress(body)
    if mode == b"E":
        inner = _decrypt_chunk(body, index, keys)
        # A key nobody has published (unreleased content) reads as zeros, as in DBCD.
        return bytes(size) if inner is None else _decode_chunk(inner, size, index, keys)
    raise TactError(f"unsupported BLTE chunk mode {mode!r}")


def parse_keys(text: str) -> dict[str, bytes]:
    """wowdev/TACTKeys' `WoW.txt`: `KEYNAME KEY` per line, both hex."""
    keys: dict[str, bytes] = {}
    for line in text.splitlines():
        parts = line.split()
        if len(parts) >= 2 and len(parts[0]) == 16 and len(parts[1]) == 32:
            keys[parts[0].upper()] = bytes.fromhex(parts[1])
    return keys


def blte_decode(data: bytes, keys: Mapping[str, bytes] | None = None) -> bytes:
    """The plain bytes of a BLTE stream; `keys` decrypt the chunks they can."""
    keys = keys or {}
    if data[:4] != BLTE_MAGIC:
        raise TactError("not a BLTE stream")
    header_size = struct.unpack_from(">I", data, 4)[0]
    if header_size == 0:
        return _decode_chunk(data[8:], 0, 0, keys)
    if data[8] != 0x0F:
        raise TactError(f"unexpected BLTE table format {data[8]:#x}")
    count = int.from_bytes(data[9:12], "big")
    out = bytearray()
    offset = header_size
    for index in range(count):
        comp, decomp = struct.unpack_from(">II", data, 12 + index * 24)
        piece = _decode_chunk(data[offset : offset + comp], decomp, index, keys)
        if len(piece) != decomp:
            raise TactError(f"BLTE chunk {index}: {len(piece)} bytes, expected {decomp}")
        out += piece
        offset += comp
    return bytes(out)


@dataclass(frozen=True)
class EncodingFile:
    """A decoded encoding file, looked up by content key one page at a time."""

    data: bytes
    page_size: int
    page_first_keys: list[bytes]
    pages_offset: int

    @classmethod
    def parse(cls, data: bytes) -> EncodingFile:
        if data[:2] != ENCODING_MAGIC:
            raise TactError("not an encoding file")
        ckey_size, ekey_size = data[3], data[4]
        if (ckey_size, ekey_size) != (16, 16):
            raise TactError(f"encoding key sizes {ckey_size}/{ekey_size}, expected 16/16")
        page_kb = struct.unpack_from(">H", data, 5)[0]
        page_count = struct.unpack_from(">I", data, 9)[0]
        espec_size = struct.unpack_from(">I", data, 18)[0]
        index_offset = 22 + espec_size
        first_keys = [
            data[index_offset + i * 32 : index_offset + i * 32 + 16] for i in range(page_count)
        ]
        return cls(data, page_kb * 1024, first_keys, index_offset + page_count * 32)

    def encoding_key(self, content_key: bytes) -> bytes | None:
        page = bisect.bisect_right(self.page_first_keys, content_key) - 1
        if page < 0:
            return None
        offset = self.pages_offset + page * self.page_size
        end = offset + self.page_size
        while offset + 22 <= end:
            key_count = self.data[offset]
            if key_count == 0:
                return None
            ckey = self.data[offset + 6 : offset + 22]
            if ckey == content_key:
                return self.data[offset + 22 : offset + 38]
            offset += 22 + 16 * key_count
        return None


def root_content_keys(data: bytes, wanted: set[int] | None = None) -> dict[int, bytes]:
    """File data id -> content key for the enUS (or every-locale) blocks of a
    WoW root, the first entry winning, as TACTSharp's normal load does.
    `wanted` limits the answer to those ids."""
    if data[:4] != ROOT_MAGIC:
        raise TactError("not an MFST root file (pre-8.2 roots are not supported)")
    version = 0
    offset = 12
    first, second = struct.unpack_from("<II", data, 4)
    if second in (1, 2):
        version, offset = second, first
    found: dict[int, bytes] = {}
    while offset < len(data):
        count = struct.unpack_from("<I", data, offset)[0]
        offset += 4
        if version == 2:
            locale, flags1, flags2 = struct.unpack_from("<III", data, offset)
            content = flags1 | flags2 | (data[offset + 12] << 17)
            offset += 13
        else:
            content, locale = struct.unpack_from("<II", data, offset)
            offset += 8
        named = not (content & CONTENT_NO_NAMES)
        deltas = struct.unpack_from(f"<{count}i", data, offset)
        ckeys_at = offset + 4 * count
        offset = ckeys_at + count * (16 + (8 if named else 0))
        all_locales = (locale & LOCALE_ALL_WOW) == LOCALE_ALL_WOW
        if (not all_locales and not locale & LOCALE_ENUS) or content & CONTENT_LOW_VIOLENCE:
            continue
        file_id = -1
        for i, delta in enumerate(deltas):
            file_id += delta + 1
            if file_id in found or (wanted is not None and file_id not in wanted):
                continue
            found[file_id] = data[ckeys_at + 16 * i : ckeys_at + 16 * i + 16]
    return found


@dataclass(frozen=True)
class ArchiveIndex:
    """One CDN archive's `.index`: encoding key -> (offset, size) in the archive."""

    data: bytes
    block_size: int
    entry_size: int
    key_size: int
    size_bytes: int
    offset_bytes: int
    last_keys: list[bytes]

    @classmethod
    def parse(cls, data: bytes) -> ArchiveIndex:
        footer = data[-INDEX_FOOTER_SIZE:]
        block_kb, offset_bytes, size_bytes, key_size = footer[3], footer[4], footer[5], footer[6]
        hash_bytes = footer[7]
        count = struct.unpack_from("<I", footer, 8)[0]
        block_size = block_kb * 1024
        entry_size = key_size + size_bytes + offset_bytes
        per_block = block_size // entry_size
        blocks = -(-count // per_block)
        toc = blocks * block_size
        last_keys = [data[toc + i * key_size : toc + (i + 1) * key_size] for i in range(blocks)]
        if toc + blocks * (key_size + hash_bytes) > len(data):
            raise TactError("archive index is shorter than its footer says")
        return cls(data, block_size, entry_size, key_size, size_bytes, offset_bytes, last_keys)

    def find(self, ekey: bytes) -> tuple[int, int] | None:
        """(offset, size) of `ekey` in the archive, or None."""
        key = ekey[: self.key_size]
        block = bisect.bisect_left(self.last_keys, key)
        if block >= len(self.last_keys):
            return None
        start = block * self.block_size
        for at in range(start, start + self.block_size - self.entry_size + 1, self.entry_size):
            entry_key = self.data[at : at + self.key_size]
            if entry_key == key:
                pos = at + self.key_size
                size = int.from_bytes(self.data[pos : pos + self.size_bytes], "big")
                pos += self.size_bytes
                offset = int.from_bytes(self.data[pos : pos + self.offset_bytes], "big")
                return offset, size
            if entry_key > key or not any(entry_key):
                return None
        return None
