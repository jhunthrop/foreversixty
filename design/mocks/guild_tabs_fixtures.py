"""Guild parity fixtures, Roster, Progression and Readiness tabs: the API answers web/tests/e2e/support/guild-parity/tabs.ts routes,
written by guild_fixtures.py into web/tests/e2e/fixtures/guild-parity/.

Progression and Readiness read the Overview's own answers (overview-progression.json, overview-readiness*.json). The Roster state's
home is the officer's Overview home with the three raiders waiting for approval (guild_overview_data.APPROVALS) moved off the
verified roster onto the waiting list, so the tab shows both its Sections.
"""
from pathlib import Path

import guild_overview_data as D
import guild_overview_fixtures as F

HOME = 'tabs-home-officer.json'


def pending_row(name: str) -> dict:
    """A waiting character: its roster row, unverified, which the officer may turn away."""
    return {**F.roster_row(D.row(name)), 'verified': False, 'may_remove': True, 'attendance': None}


def roster_home() -> dict:
    home = F.home(D.OFFICER, True)
    return {**home, 'roster': [r for r in home['roster'] if r['name'] not in D.APPROVALS], 'pending': [pending_row(n) for n in D.APPROVALS]}


def fixtures() -> dict[str, object]:
    """File name to JSON body."""
    return {HOME: roster_home()}


def writer(write) -> list[Path]:
    return [write(name, data) for name, data in fixtures().items()]
