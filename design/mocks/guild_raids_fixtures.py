"""Guild parity fixtures, Raids tab (list, plan form, log match) and the raid page (planner, assignments): the API answers web/tests/e2e/support/guild-parity/raids.ts routes,
written by guild_fixtures.py into web/tests/e2e/fixtures/guild-parity/.
"""
from pathlib import Path


def fixtures() -> dict[str, object]:
    """File name to JSON body."""
    return {}


def writer(write) -> list[Path]:
    return [write(name, data) for name, data in fixtures().items()]
