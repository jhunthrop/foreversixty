"""The content phase calendar, stated once.

Contract 10.4. The four boundaries were written out twice -- as a Go
table in `api/internal/phase/phase.go` and as prose in
`web/src/data/dates.json` -- with a comment asking whoever moved one to
move the other. `curated/phases.json` is now the statement; the Go
package is tested against it (the API lane's test), the web reads the
copy this emits, and `GET /v1/phases` serves it.

`later` is deliberately not in here. It is contract 10.4's sentinel for a
loot source with no known date, and a boundary by that name would make
`phase.At` bracket a moment into it.
"""

from __future__ import annotations

import json
from datetime import datetime
from pathlib import Path

from pipeline.models import PhaseBoundary
from pipeline.normalize import write_records

WEB_PATH = Path("../web/src/data/phases.json")

#: The sentinel a loot source carries when its date is unknown. Named
#: here so the loot validator and this module cannot disagree about it.
OPENS_LATER = "later"


class PhaseError(SystemExit):
    """curated/phases.json says something the pipeline will not publish."""


def _instant(value: str, where: str) -> datetime:
    """Parse an RFC 3339 UTC instant, which is what Go's time.Time reads.

    `Z` and nothing else: an offset would still parse here and then
    compare unequal to the Go table's UTC instants for no visible reason.
    """
    if not value.endswith("Z"):
        raise PhaseError(f"{where} start {value!r} is not RFC 3339 UTC (it must end in Z)")
    try:
        return datetime.fromisoformat(value.replace("Z", "+00:00"))
    except ValueError as error:
        raise PhaseError(f"{where} start {value!r} is not RFC 3339: {error}") from error


def load_phases(curated_dir: Path = Path("curated")) -> list[PhaseBoundary]:
    path = curated_dir / "phases.json"
    if not path.exists():
        raise PhaseError(f"missing {path}")
    rows = [PhaseBoundary(**entry) for entry in json.loads(path.read_text(encoding="utf-8"))]
    if not rows:
        raise PhaseError(f"{path} lists no phase")
    seen: set[str] = set()
    previous: datetime | None = None
    for row in rows:
        if row.name in seen:
            raise PhaseError(f"{path} names phase {row.name!r} twice")
        seen.add(row.name)
        if row.name == OPENS_LATER:
            raise PhaseError(
                f"{path} names a phase {OPENS_LATER!r}; that is the sentinel a loot "
                f"source carries when its date is unknown, not a boundary"
            )
        start = _instant(row.start, f"{path} phase {row.name}")
        if previous is not None and start < previous:
            raise PhaseError(
                f"{path} is out of order: {row.name!r} starts before the phase above it"
            )
        previous = start
    return rows


def write_phases(curated_dir: Path = Path("curated"), web_path: Path = WEB_PATH) -> Path:
    write_records(load_phases(curated_dir), web_path)
    return web_path


def check_phases(curated_dir: Path = Path("curated"), web_path: Path = WEB_PATH) -> bool:
    """True when the emitted copy has drifted from the curated one.

    Compares contents, not timestamps, so it catches an edit made straight
    to the generated file as well as one to the curated list that was
    never re-emitted -- the same gate `pipeline/specs.py`'s
    `check_specs` is.
    """
    import tempfile

    rows = load_phases(curated_dir)
    with tempfile.TemporaryDirectory() as directory:
        expected = Path(directory) / "phases.json"
        write_records(rows, expected)
        wanted = expected.read_text(encoding="utf-8")
    return not web_path.exists() or web_path.read_text(encoding="utf-8") != wanted
