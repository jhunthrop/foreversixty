"""Writes the guild parity fixtures: web/tests/e2e/fixtures/guild-parity/*.json, the API answers the build is served
(web/tests/e2e/support/guild-parity/) for the data ws_guild2.py draws the mocks from (guild_data.py).

  python3 design/mocks/guild_fixtures.py
"""
import json
from pathlib import Path

OUT = Path(__file__).resolve().parents[2] / 'web/tests/e2e/fixtures/guild-parity'


def write(name: str, data: object) -> Path:
    OUT.mkdir(parents=True, exist_ok=True)
    path = OUT / name
    path.write_text(json.dumps(data, indent=1) + '\n')
    return path


# One writer per page group; each returns the files it wrote.
WRITERS: list = []


def overview_writer() -> list[Path]:
    """Lane A: the member and officer Overviews and the Loot tab (guild_overview_fixtures.py)."""
    from guild_overview_fixtures import overview_fixtures
    return [write(name, data) for name, data in overview_fixtures().items()]


WRITERS.append(overview_writer)

if __name__ == '__main__':
    for writer in WRITERS:
        for path in writer():
            print(path)
