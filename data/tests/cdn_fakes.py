"""A `pipeline.casc.FileSource` for tests: client files by file data id, no network."""

from __future__ import annotations

from collections.abc import Callable, Mapping

from pipeline.blizzard_cdn import CdnMissing


class FakeSource:
    """Serves `files` by id (bytes, or a callable raising what a CDN read would);
    an id it lacks is missing from the build. `calls` records every id read."""

    def __init__(self, files: Mapping[int, bytes | Callable[[], bytes]] | None = None) -> None:
        self.files = dict(files or {})
        self.calls: list[int] = []

    def read(self, file_id: int) -> bytes:
        self.calls.append(file_id)
        if file_id not in self.files:
            raise CdnMissing(f"file data id {file_id} is not in the fake build")
        value = self.files[file_id]
        return value() if callable(value) else value
