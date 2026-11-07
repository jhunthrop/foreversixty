"""Guild parity fixtures, Settings tab and the guild settings page (billing, Discord, crest, invite link): the API answers web/tests/e2e/support/guild-parity/settings.ts routes,
written by guild_fixtures.py into web/tests/e2e/fixtures/guild-parity/.
"""
from pathlib import Path


def fixtures() -> dict[str, object]:
    """File name to JSON body."""
    return {}


def writer(write) -> list[Path]:
    return [write(name, data) for name, data in fixtures().items()]
