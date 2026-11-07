"""Guild parity fixtures, Overview states: before raids open, no raid logged, not in a guild, low-level character: the API answers web/tests/e2e/support/guild-parity/states.ts routes,
written by guild_fixtures.py into web/tests/e2e/fixtures/guild-parity/.
"""
from pathlib import Path


def fixtures() -> dict[str, object]:
    """File name to JSON body."""
    return {}


def writer(write) -> list[Path]:
    return [write(name, data) for name, data in fixtures().items()]
