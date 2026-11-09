"""/rankings in the directions system (plan 2026-11-09-redesign, batch 2 rankings, drawn with the logs batch).

  python3 design/mocks/directions/rankings.py   -> html/b1-rankings.html, b1-rankings-board.html, b1-rankings-guilds.html

Data:
- The character board is the rankings spec's example board (design/mocks/gen_rankings.py, design/specs/2026-10-11-rankings.md:
  16 Fury Warriors on Onyxia, 212 ranked, the signed-in Obnoxious Yell at 14). Its names, guilds, DPS, lengths, executed and
  dates are invented and labelled so there; the encounter, class, spec and the 212 are the e2e stub's
  (web/tests/e2e/support/rankings-fixture.ts). The e2e stub's own rows step by exactly 12 DPS and 1 s, which read as
  generated (reviewers, round 2), so the spec's examples stand in.
- The encounter list is the live one (api.foreversixty.gg/v1/rankings/encounters, 9 Nov: three Excavation Site bosses,
  two ranked kills each) with Onyxia from the stub.
- The guild board's guilds and kill times are the spec's examples too (the board's guilds, their best Onyxia length).
Copy is the page's own (web/src/lib/rankings/copy.ts through design/mocks/data/logs-facts.json).
"""
import json
import sys
from html import escape
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

import gen_rankings as G  # noqa: E402
from common import write  # noqa: E402
from prims import PRIMS_CSS, facts, panel, parse_colour, parse_legend, player_cell, row_list, strip_into_view, table  # noqa: E402
from style import footer, header, page, titlebar  # noqa: E402

HERE = Path(__file__).resolve().parent
COPY = json.loads((HERE.parent / 'data/logs-facts.json').read_text())['copy']['rankingsCopy']
VIEWER = ('Obnoxious Yell', 'warrior', 60)
OWN_GUILD = 'Olympus XXVII'
COHORT = 'Fury Warriors'
ENCOUNTER = G.ENCOUNTER
TOTAL = G.TOTAL
UPDATED_DAY = 13  # the stub's updated_at, 13 Dec 2026; G.DAYS counts back from it
LIVE_ENCOUNTERS = [('Relic Guardian', 'Excavation Site: Wetlands', 5, 2), ('Saltspine', 'Excavation Site: Wetlands', 5, 2),
                   ('Shadetooth', 'Excavation Site: Wetlands', 5, 2)]
BOARD = [r for r in G.BOARD if r[4] == 'US']  # the board's filter is US; the EU rows sit on the EU board
OWN = next(r for r in BOARD if r[1] == VIEWER[0])

RANK_CSS = PRIMS_CSS + '''
.rp{padding-top:6px}.rk{font:700 15px var(--body);font-variant-numeric:tabular-nums}
.tb td .gl{font-size:12.5px;color:var(--muted)}.tb a.rep{font-weight:600}.tb .bd{color:var(--text);font-weight:400}.tb .ng{color:var(--muted)}
.you{display:grid;gap:4px;padding-top:2px}.you .big{font:600 34px/1 var(--body);color:#fff;font-variant-numeric:tabular-nums}
.you .big small{font-size:15px;font-weight:600;color:var(--text)}.you .note{font-size:14px;color:var(--muted)}
.legend{padding:0 20px 10px}.lead{font-size:15px;color:var(--text)}
@media(max-width:700px){.legend{padding:0 14px 8px}.you .big{font-size:28px}}
'''


def pct(rank: int, total: int) -> float:
    return 100 * (1 - (rank - 1) / total)


def rank_cell(rank: int, total: int) -> str:
    return f'<span class="rk" style="color:{parse_colour(pct(rank, total))}">{rank}</span>'


def killed(days_ago: int) -> str:
    """The example's days-ago, folded into the raids' first week (9 to 13 Dec): no kill before Onyxia opens."""
    return f'{UPDATED_DAY - days_ago % 5} Dec'


def band(tab: str, controls: list[str]) -> str:
    ctl = [[('Characters', tab == 'Characters'), ('Guilds', tab == 'Guilds')], [(f'{x} &#9662;', False) for x in controls]]
    band_html = titlebar(None, COPY['h1'], 'Updated 13 Dec 2026', ctl=ctl, top=header(*VIEWER, 'Logs'))
    return '<div class="rb">' + band_html.replace('&amp;#9662;', '&#9662;') + '</div>' + strip_into_view()


def tier_pointer() -> str:
    return panel('', f'<p class="lead">{COPY["tierQuestion"]} <a href="#" style="white-space:nowrap">{COPY["tierLink"]} &rsaquo;</a></p>')


def two_col(main: str, rail: str, answer_first: bool = False) -> str:
    """answer_first: on a phone the rail (the viewer's own rank) comes before the board (ux-designer, round 2)."""
    return f'<main class="w rp"><div class="twocol{" ans" if answer_first else ""}"><div>{main}</div><aside>{rail}</aside></div></main>'


# ------------------------------------------------------------------------------------------------ character board
def board_rows() -> tuple[list[list[str]], list[str]]:
    rows, cls = [], []
    for i, (_, name, slug, guild, _region, _size, value, executed, days, split, length) in enumerate(BOARD):
        rank = i + 1
        sub = f'<span class="om gl">{escape(guild)}</span>' if guild else ''  # the guild column drops on a phone; the name carries it
        rows.append([rank_cell(rank, TOTAL), player_cell(slug, name, sub, 26), escape(guild) if guild else '<span class="ng">No guild</span>',
                     f'{value:,}', f'{executed}%', length, killed(days), f'<a class="bd" href="#">{split}</a>', '<a class="rep" href="#">Report</a>'])
        cls.append('me' if name == VIEWER[0] else '')
    return rows, cls


def your_rank() -> str:
    rank = BOARD.index(OWN) + 1
    _, _, _, _, _, _, value, executed, days, _, length = OWN
    best = facts([('DPS', f'{value:,}'), ('Executed', f'{executed}%'), ('Length', length), ('Killed', killed(days))])
    body = (f'<div class="you"><p class="big" style="color:{parse_colour(pct(rank, TOTAL))}">#{rank} <small>of {TOTAL} {COHORT}</small></p>'
            f'<p class="note">On {ENCOUNTER}, Normal ruleset</p></div>'
            f'<div style="margin-top:14px">{best}</div><p style="margin-top:14px"><a href="#">{COPY["parsesLink"]} &rsaquo;</a></p>')
    return panel(COPY['answerLabel'], body, framed=True)


def board_page() -> str:
    rows, cls = board_rows()
    heads = [('#', True), ('Character', False), ('Guild', False), ('DPS', True), ('Executed', True), ('Length', True), ('Date', True),
             ('Build', False), ('', False)]
    tbl = table(heads, rows, sort='DPS', row_cls=cls, hide_phone={2, 4, 5, 6, 7})
    board = panel(COHORT, f'<div class="legend">{parse_legend("Percentile")}</div>{tbl}', action=f'<a href="#">{COPY["anotherBoard"]}</a>',
                  flush=True, framed=True, ornate=True)
    body = band('Characters', [ENCOUNTER, 'Fury Warrior', 'Damage', 'Normal', 'US']) + two_col(board, your_rank() + tier_pointer(), True) + footer()
    return page('Forever Sixty, rankings, Onyxia', RANK_CSS, body)


# ------------------------------------------------------------------------------------------------ guild board
def guild_rows() -> list[tuple[str, str, str]]:
    """Each guild on the board at its fastest kill (the spec's example lengths), regions as the board gives them."""
    best: dict[str, tuple[str, str, int]] = {}
    for _, _, _, guild, region, _, _, _, days, _, length in G.BOARD:
        if guild and (guild not in best or length < best[guild][1]):
            best[guild] = (region, length, days)
    return sorted(((g, *v) for g, v in best.items()), key=lambda g: g[2])


def guild_page() -> str:
    ordered = guild_rows()
    rows = [[f'<span class="rk" style="color:#fff">{i + 1}</span>', f'<b style="color:#fff">{escape(n)}</b>', region, t, killed(days),
             '<a class="rep" href="#">Report</a>'] for i, (n, region, t, days) in enumerate(ordered)]
    cls = ['me' if n == OWN_GUILD else '' for n, *_ in ordered]
    tbl = table([('#', True), ('Guild', False), ('Region', False), ('Kill time', True), ('Date', True), ('', False)], rows, sort='Kill time',
                row_cls=cls, hide_phone={2, 4})
    board = panel('Fastest kills', tbl, flush=True, framed=True, ornate=True)
    own = next(i for i, g in enumerate(ordered) if g[0] == OWN_GUILD)
    yours = panel('Your guild', f'<div class="you"><p style="color:#fff;font-weight:600">{OWN_GUILD}</p><p class="big">#{own + 1} '
                  f'<small>of {len(ordered)} guilds</small></p><p class="note">{ordered[own][2]} on {ENCOUNTER}</p></div>', framed=True)
    body = band('Guilds', [ENCOUNTER, 'Speed', 'Normal', 'All regions']) + two_col(board, yours, True) + footer()
    return page('Forever Sixty, rankings, guilds', RANK_CSS, body)


# ------------------------------------------------------------------------------------------------ encounter picker
def picker_page() -> str:
    enc = [(ENCOUNTER, "Onyxia's Lair", 40, TOTAL)] + LIVE_ENCOUNTERS
    rows = row_list([('', f'<a href="#">{escape(n)}</a>', f'{escape(zone)} &middot; {size}-player', f'{kills:,} ranked')
                     for n, zone, size, kills in enc])
    main = panel(COPY['pickerHeading'], rows, framed=True, ornate=True)
    rail = panel(COPY['answerLabel'], f'<p><b class="num" style="color:{parse_colour(pct(BOARD.index(OWN) + 1, TOTAL))}">#{BOARD.index(OWN) + 1}</b> '
                                      f'<span class="muted">of {TOTAL} {COHORT} on {ENCOUNTER}</span></p>'
                                      )
    body = band('Characters', []) + two_col(main, rail) + footer()
    return page('Forever Sixty, rankings', RANK_CSS, body)


def build() -> None:
    write('b1-rankings', picker_page())
    write('b1-rankings-board', board_page())
    write('b1-rankings-guilds', guild_page())


if __name__ == '__main__':
    build()
