"""The gallery's Logs section: python3 design/mocks/directions/logs_gallery.py [gallery dir]

gallery.py includes section() when it rebuilds the whole page. Run on its own, this module copies the logs PNGs into an
existing gallery (default ~/Desktop/forever-design-directions) and puts the section in place between its markers,
leaving every other section as it is. Verdicts come from logs_reviews.json.
"""
import json
import shutil
import sys
from html import escape
from pathlib import Path

HERE = Path(__file__).resolve().parent
OUT = Path(sys.argv[1]) if sys.argv[1:] else Path.home() / 'Desktop' / 'forever-design-directions'
START, END = '<!-- logs:start -->', '<!-- logs:end -->'
PAGES = [('Logs', 'b1-logs'), ('Your reports', 'b1-reports'), ('Report: Summary', 'b1-report-summary'),
         ('Report: Damage done', 'b1-report-damage'), ('Report: Healing', 'b1-report-healing'), ('Report: Damage taken', 'b1-report-taken'),
         ('Report: Deaths', 'b1-report-deaths'), ('Report: Mechanics (Onyxia, fixture)', 'b1-report-mechanics'),
         ('Report: Timeline', 'b1-report-timeline'), ('Report: Threat chart', 'b1-report-threat'), ('Report: Resources charts', 'b1-report-resources'),
         ('Report: Queries', 'b1-report-queries'), ('Report: one player, abilities and gear at the pull', 'b1-report-player'),
         ('Report: Coach (fixture)', 'b1-report-coach'), ('Rankings: encounters', 'b1-rankings'), ('Rankings: character board', 'b1-rankings-board'),
         ('Rankings: guild board', 'b1-rankings-guilds')]
NOTE = ('Batch 1 of the redesign: every logs, reports and rankings surface in the round-5 look, from the shared primitives '
        '(design/mocks/directions/prims.py: panel, table, player cell, meter, facts, row list, chips, tag, code, and one chart axis for line, '
        'area and column charts). Report pulls are the real 17 Sept beta night (trash only). Mechanics reads Onyxia\'s mechanics table and the '
        'guild report fixture\'s wipe; Coach is the Fury fixture; the player view\'s gear is the Beast Mastery 20 to 29 list as a stand-in (the '
        'sample has no gear); the rankings board is the e2e stub (Fury Warriors on Onyxia, 212 ranked), its guild board the rankings spec\'s examples.')


def fig(src: str, cap: str) -> str:
    return f'<figure><a href="{src}"><img src="{src}" alt="{escape(cap)}" loading="lazy"></a><figcaption>{escape(cap)}</figcaption></figure>'


def verdicts() -> str:
    path = HERE / 'logs_reviews.json'
    if not path.exists():
        return ''
    rv = json.loads(path.read_text())
    cells = ''.join(f'<div><b>{escape(who)}</b><span class="{escape(v["tone"])}">{escape(v["verdict"])}</span><p>{escape(v["note"])}</p></div>'
                    for who, v in rv.items())
    return f'<h3>Reviewers</h3><div class="rv">{cells}</div>'


def section() -> str:
    rows = ''.join(f'<h3>{escape(name)}</h3><div class="two" style="grid-template-columns:3fr 1fr">{fig(f"png/{key}-1440.png", "1440")}'
                   f'{fig(f"png/{key}-390.png", "390")}</div>' for name, key in PAGES)
    return f'{START}<h2 id="logs">Logs</h2><p>{escape(NOTE)}</p>{verdicts()}{rows}{END}'


def install(out: Path) -> Path:
    index = out / 'index.html'
    html = index.read_text()
    for _, key in PAGES:
        for w in (1440, 390):
            shutil.copy2(HERE / 'png' / f'{key}-{w}.png', out / 'png' / f'{key}-{w}.png')
    if START in html:
        head, rest = html.split(START, 1)
        html = head + section() + rest.split(END, 1)[1]
    else:
        html = html.replace('</main>', section() + '</main>', 1)
    index.write_text(html)
    return index


if __name__ == '__main__':
    print(install(OUT))
