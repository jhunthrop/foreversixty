"""Guild parity fixtures, Roster, Progression and Readiness tabs: the API answers web/tests/e2e/support/guild-parity/tabs.ts routes,
written by guild_fixtures.py into web/tests/e2e/fixtures/guild-parity/.

Progression and Readiness read the Overview's own answers (overview-progression.json, overview-readiness*.json). The Roster state's
home is the officer's Overview home: the verified roster and the three characters waiting for approval (guild_overview_data.PENDING),
so the tab shows both its Sections.
"""
from pathlib import Path

import guild_overview_data as D
import guild_overview_fixtures as F

HOME = 'tabs-home-officer.json'


def roster_home() -> dict:
    return F.home(D.OFFICER, True)


def fixtures() -> dict[str, object]:
    """File name to JSON body."""
    return {HOME: roster_home()}


def writer(write) -> list[Path]:
    return [write(name, data) for name, data in fixtures().items()]
