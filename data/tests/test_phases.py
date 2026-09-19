import json
from pathlib import Path

import pytest

from pipeline.phases import PhaseError, check_phases, load_phases, write_phases

CURATED = Path("curated")
WEB = Path("../web/src/data/phases.json")

#: api/internal/phase/phase.go's table, transcribed. The Go side is
#: tested against the file this produces; this is the other direction, so
#: that an edit to the curated file that nobody meant shows up here too.
BOUNDARIES = [
    ("pre-beta", "0001-01-01T00:00:00Z"),
    ("beta", "2026-09-17T00:00:00Z"),
    ("launch", "2026-11-04T23:00:00Z"),
    ("raids-1", "2026-12-09T00:00:00Z"),
]


def test_the_curated_calendar_is_the_go_tables_boundaries_in_order():
    assert [(p.name, p.start) for p in load_phases(CURATED)] == BOUNDARIES


def test_the_sentinel_is_not_a_boundary():
    """`later` means 'unreleased, no date' on a loot source (contract
    10.4). It is not a phase: phase.At would bracket into it."""
    assert "later" not in {p.name for p in load_phases(CURATED)}


def test_the_emitted_web_copy_matches_the_curated_one():
    assert not check_phases(CURATED, WEB)
    emitted = json.loads(WEB.read_text(encoding="utf-8"))
    assert [(row["name"], row["start"]) for row in emitted] == BOUNDARIES


def test_a_start_that_is_not_rfc3339_utc_is_refused(tmp_path):
    (tmp_path / "phases.json").write_text(
        json.dumps([{"name": "beta", "start": "2026-09-17"}]), encoding="utf-8"
    )
    with pytest.raises(PhaseError, match="RFC 3339"):
        load_phases(tmp_path)


def test_boundaries_out_of_order_are_refused(tmp_path):
    (tmp_path / "phases.json").write_text(
        json.dumps(
            [
                {"name": "launch", "start": "2026-11-04T23:00:00Z"},
                {"name": "beta", "start": "2026-09-17T00:00:00Z"},
            ]
        ),
        encoding="utf-8",
    )
    with pytest.raises(PhaseError, match="order"):
        load_phases(tmp_path)


def test_a_duplicate_phase_name_is_refused(tmp_path):
    (tmp_path / "phases.json").write_text(
        json.dumps(
            [
                {"name": "beta", "start": "2026-09-17T00:00:00Z"},
                {"name": "beta", "start": "2026-11-04T23:00:00Z"},
            ]
        ),
        encoding="utf-8",
    )
    with pytest.raises(PhaseError, match="twice"):
        load_phases(tmp_path)


def test_write_phases_is_deterministic(tmp_path):
    out = tmp_path / "phases.json"
    write_phases(CURATED, out)
    once = out.read_bytes()
    write_phases(CURATED, out)
    assert out.read_bytes() == once
