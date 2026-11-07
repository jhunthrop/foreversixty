"""Writes the guild parity fixtures: web/tests/e2e/fixtures/guild-parity/*.json, the API answers the build is served
(web/tests/e2e/support/guild-parity/) for the data ws_guild2.py draws the mocks from (guild_data.py).

  python3 design/mocks/guild_fixtures.py
"""
import json
import re
from pathlib import Path

OUT = Path(__file__).resolve().parents[2] / 'web/tests/e2e/fixtures/guild-parity'


PRINT_WIDTH = 110  # web/.prettierrc printWidth: the fixtures pass `prettier --check` as written
_SCALAR_LIST = re.compile(r'^( *)("[^"\n]*": )?\[\n((?:\1  (?:"(?:[^"\\\n]|\\.)*"|-?[\d.eE+-]+|true|false|null),?\n)+)\1\]', re.M)


def _one_line(m: re.Match) -> str:
    indent, key, body = m.group(1), m.group(2) or '', m.group(3)
    line = f'{indent}{key}[{", ".join(x.strip().rstrip(",") for x in body.splitlines())}]'
    return line if len(line) + 1 <= PRINT_WIDTH else m.group(0)


def write(name: str, data: object) -> Path:
    """JSON as prettier prints it: two-space indent, a list of scalars on one line when it fits."""
    OUT.mkdir(parents=True, exist_ok=True)
    path = OUT / name
    path.write_text(_SCALAR_LIST.sub(_one_line, json.dumps(data, indent=2)) + '\n')
    return path


def write_raid_planner() -> list[Path]:
    """Lane B: Thursday's Onyxia's Lair planner as its officer (Dunmaro) reads it (guild_raid.py)."""
    import guild_raid
    return [write('raid-me.json', guild_raid.me()), write('raid-guild.json', guild_raid.guild_page()),
            write('raid-home.json', guild_raid.guild_home()), write('raid-onyxia.json', guild_raid.raid_view())]


# One writer per page group; each returns the files it wrote.
WRITERS: list = [write_raid_planner]

if __name__ == '__main__':
    for writer in WRITERS:
        for path in writer():
            print(path)
