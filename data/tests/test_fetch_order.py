"""data.yml's fetch job runs each pipeline stage after the stages whose output it reads.

The client watcher's fetch of 1.60.1.70888 failed for two days at `loot`: its client
enchant table reads the rating factors from `gametables/combatratings.txt`, and
`gametables` ran after it (runs from 10 Nov, e.g. 43212317094). A new build directory
has none of a stage's inputs, so a stage that runs too early fails only on a new
build, never on a re-fetch. This reads the order off the workflow itself.
"""

from __future__ import annotations

import re
from pathlib import Path

import pytest

WORKFLOW = Path(__file__).resolve().parents[2] / ".github" / "workflows" / "data.yml"
STAGE = re.compile(r"^\s*- run: uv run python -m pipeline ([a-z-]+)", re.MULTILINE)

#: (producer, consumer): the consumer reads a file the producer writes.
READS_AFTER = [
    ("normalize", "gametables"),  # gametables.py refreshes normalize's manifest.json
    ("gametables", "loot"),  # loot/client_enchants.py: load_rating_factors
    ("gametables", "simdb"),  # simdb/__init__.py: load_rating_factors
    ("simconst", "icons"),
    ("levels", "icons"),
    ("loot", "simdb"),
    ("trainables", "spelltips"),
    ("loot", "spelltips"),
]


def fetch_job_stages() -> list[str]:
    text = WORKFLOW.read_text()
    fetch = text[text.index("\n  fetch:\n") :]
    return STAGE.findall(fetch)


@pytest.mark.parametrize(("producer", "consumer"), READS_AFTER)
def test_a_stage_runs_after_the_stage_it_reads(producer: str, consumer: str) -> None:
    stages = fetch_job_stages()
    assert producer in stages and consumer in stages, stages
    assert stages.index(producer) < stages.index(consumer), (
        f"`{consumer}` reads what `{producer}` writes, but the fetch job runs it first: {stages}"
    )
