"""Web-layout parity mocks for /rankings and /rankings/<slug> (lane P5, plan 2026-11-09-redesign), composed ONLY from ws_prims.

The look is the directions mock (design/mocks/directions/rankings.py, png/b1-rankings*.png): the compact title band, the
board tabs and filters on the band, the board in the page's lead panel (the percentile legend, then the compact ranked list
with crests, class colours and the percentile column) and the rail beside it (Your rank framed, the tier list pointer).
This generator carries the built structure and its regions (Rankings.svelte: Page split with panels, Section panel,
RankList compact).

Data and copy, read at render time, never retyped: design/mocks/data/rankings-facts.json (node design/mocks/rankings_facts.mjs):
the e2e rankings stubs (web/tests/e2e/support/rankings-fixture.ts) read through the page's own view helpers
(lib/rankings/rank-row-view.ts, guild-row-view.ts, own-best.ts, own-rank.ts) and lib/rankings/copy.ts.

  node design/mocks/rankings_facts.mjs && python3 design/mocks/ws_rankings.py
"""
import json
import os
import re
import sys
from html import escape
from pathlib import Path

from mocklib import CLASS_COLOR, PARSE
from ws_prims import (Button, Facts, Figure, Link, Page, PageHeader, Plate, RankList, RankRow, ColumnHeads, Row, Rows, Section,
                      build_page, control_bar, crest_img, seg, sel, write)

MOCKDATA = Path(__file__).resolve().parent / 'data'
FACTS = json.loads(Path(os.environ.get('FS_RANKINGS_FACTS') or MOCKDATA / 'rankings-facts.json').read_text())
COPY = FACTS['copy']
RC, EC, EMPTY = COPY['rankingsCopy'], COPY['encounterPickerCopy'], COPY['rankingsEmptyCopy']
YOU = ('Obnoxious Yell', 'warrior', 'Fury Warrior &middot; 60')
UPDATED = 'Updated 13 Dec'   # the stubs' updated_at, formatDate (and the stamp's reserve before a board answers)
PARSE_HEX = {**PARSE, 'gold': '#e5cc80'}   # tokens.css --color-parse-*


def token_hex(css: str) -> str:
    """'var(--color-parse-gold)' -> its hex, as the built page resolves it."""
    m = re.fullmatch(r'var\(--color-parse-([a-z]+)\)', css or '')
    return PARSE_HEX[m.group(1)] if m else css


PAGE_CSS = '''<style>
.legend{display:flex;flex-wrap:wrap;gap:4px 14px;align-items:baseline;padding:0 20px 12px;font-size:13px;color:var(--muted)}
.legend b{font:700 13px/20px var(--font-body)}
.pager{display:flex;gap:16px;align-items:center;justify-content:flex-end;padding:12px 20px 16px;border-top:1px solid var(--line-soft);font-size:14px}
.pin{border-bottom:1px solid var(--line)}
.empty{padding:4px 20px 20px}.empty .es{display:flex;flex-direction:column;gap:4px;padding:4px 0}.empty .ln{color:var(--strong);font-weight:600}
.empty .fl{font-size:14px;line-height:20px;color:var(--muted)}.empty .act{display:flex;gap:8px 20px;align-items:center;flex-wrap:wrap;margin-top:12px}
.else{margin-top:16px;padding-top:16px;border-top:1px solid var(--line-soft)}.else .t-label{display:block;padding:12px 20px 0}
.encs{padding:0 20px 8px}.ptr{display:flex;flex-wrap:wrap;gap:4px 12px;align-items:baseline;padding:16px}
.you-n{display:flex;flex-direction:column;gap:4px;min-height:88px}.you-n .kv{margin-top:12px}.you-n p.go{margin-top:12px}

.pin-l{display:block;padding:0 20px 4px}.lrows{padding-bottom:8px}.enc-b{border-top:1px solid var(--line-soft)}.enc-b .encs{padding-bottom:0}.encs .rows{margin:0}.gold{color:var(--gold)}.ruled{border-top:1px solid var(--line-soft)}
</style>'''


def legend() -> str:
    steps = ''.join(f'<b style="color:{token_hex(s["color"])}">{escape(s["label"])}</b>' for s in FACTS['legend'])
    return f'<p class="legend"><span>{escape(RC["legendLead"])}</span>{steps}</p>'


def figure_cells(figures: list[dict]) -> list[tuple]:
    """The row's figures as the wide row draws them: a phone-only one (the Date under the figure) is not a column."""
    return [figure_cell(f) for f in figures if not f.get('phoneOnly')]


def figure_cell(f: dict) -> tuple:
    """A built RankFigure as the mock RankRow's (value, unit, phone label, best); a figure with an href is a link."""
    v = escape(f['value'])
    return (f'<a href="#">{v}</a>' if f.get('href') else v, '', escape(f['label']), False)


def board_row(v: dict, you: bool = False) -> str:
    rank = f'<span style="color:{token_hex(v["rankColor"])}">{v["rank"]}</span>'
    return RankRow(rank, crest_img(v['classSlug'], 28), CLASS_COLOR[v['classSlug']], escape(v['name']), escape(v['sub']), v['fraction'],
                   figure_cells(v['figures']), escape(v['gap']['text']), [(escape(lk['label']), False) for lk in v.get('links', [])],
                   you='You' if you else '', selected=you)


def board_list(b: dict, own_key: str = '') -> str:
    heads = ColumnHeads([('#', False), ('', False), ('Character', False), ('', False), ('DPS', True), (RC['percentileLabel'], True),
                         ('Executed', True), ('Length', True), ('Date', True), ('', False)])
    pin = ''
    if b['pinned']:   # captioned, its own list above the board (Rankings.svelte)
        pin = (f'<p class="t-label pin-l">{escape(RC["ownBest"])}</p>'
               + RankList(heads, board_row(b['pinned'], you=True), 4, compact=True, fig_px=72))
    rows = ''.join(board_row(r, you=r['key'] == own_key and not b['pinned']) for r in b['rows'])
    pages = -(-b['total'] // b['per_page'])
    pager = (f'<div class="pager"><span class="t-small">Page 1 of {pages}</span><span class="muted">Previous</span>{Link("Next")}</div>'
             if pages > 1 else '')
    board = RankList('' if b['pinned'] else heads, rows, 4, compact=True, fig_px=72)   # one column head serves both lists
    return legend() + pin + f'<div class="ruled">{board}</div>' + pager


def bar(groups: list[tuple[str, str]], summary: str) -> str:
    return control_bar(groups, summary, label='Ranking choices', labels_from=1920)


def board_tabs(on: str) -> tuple[str, str]:
    return (RC['boardLabel'], seg([('Characters', on == 'character', None, ''), ('Guilds', on == 'guild', None, '')]))


def character_bar(encounter: str, cls: str, spec: str, ruleset: str, summary: str) -> str:
    filters = '<button type="button" class="sel" aria-expanded="false" aria-haspopup="dialog">Filters</button>'
    return bar([board_tabs('character'), (RC['encounterLabel'], sel(encounter)), (RC['classLabel'], sel(cls)), (RC['specLabel'], sel(spec)),
                (RC['rulesetLabel'], sel(ruleset)), (RC['metricLabel'], seg([('Damage', True, None, ''), ('Healing', False, None, '')])),
                ('Filters', filters)], summary)


def pointer() -> str:
    """The tier list is a link in the band now (Rankings.svelte): the rail carries no pointer card."""
    return ''


def companion() -> str:
    """Nothing ranked where the visitor looks: the rail says what fills it (scoreCopy's companion lines)."""
    sc = COPY['scoreCopy']
    return Section(escape(sc['explainerLogsHeading']), f'<p class="t-small">{escape(sc["explainerLogsBody"])}</p>'
                   f'<p style="margin-top:12px">{Button(escape(sc["explainerLogsButton"]))}</p>', panel=True)


def your_rank(answer: dict, unranked: bool = False) -> str:
    if unranked:
        body = (f'<div class="you-n"><p>{escape(answer["text"])}</p><p class="go" style="margin-top:8px">'
                f'{Link(escape(EMPTY["action"]) + " &rsaquo;")}</p></div>')
    else:
        fig = answer['figures'][0]
        big = Figure(escape(fig['value']), unit=escape(fig['unit']), size='big').replace('class="v"', f'class="v" style="color:{token_hex(fig["color"])}"', 1)
        # The best kill's figures are the marked row's: the answer and its next step only (Rankings.svelte).
        body = (f'<div class="you-n">{big}'
                f'<p class="go">{Link(escape(RC["parsesLink"]) + " &rsaquo;")}</p></div>')
    return Section(RC['answerLabel'], body, panel=True)


def empty_block(line: str, fill: str, widen: bool = False, elsewhere: str = '') -> str:
    w = Link(escape(EMPTY['widen']) + ' &rsaquo;') if widen else ''
    tail = f'<div class="else">{elsewhere}</div>' if elsewhere else ''
    # RankingsEmpty.svelte: the padded wrapper, ui/EmptyState's `fill` shape (py-1), and the bosses ranked elsewhere
    # as its children, behind a hairline inside the same padding.
    return (f'<div class="empty"><div class="es"><p class="ln">{escape(line)}</p><p class="fl">{escape(fill)}</p>'
            f'<div class="act">{Button(escape(EMPTY["action"]), primary=True, small=True)}{w}</div>{tail}</div></div>')


def encounter_rows(encs: list[dict]) -> str:
    rows = [Row(escape(e['name']), escape(FACTS['fns']['where'][e['slug']]),
                trail=f'<span class="num">{escape(FACTS["fns"]["ranked"][e["slug"]])}</span>') for e in encs]
    return f'<div class="encs">{Rows(rows)}</div>'


def head(meta: str) -> str:
    return PageHeader(escape(RC['h1']), meta=[meta, Link(escape(RC['tierLink']) + ' &rsaquo;')], lead=escape(RC['lead']))


def page(title: str, html: str, who: tuple | None, note: str) -> str:
    return build_page(title, html + PAGE_CSS, 'Logs', who, note)


HIDDEN_STAMP = f'<span style="visibility:hidden">{UPDATED}</span>'   # the stamp's reserve: no character board answered


def landing_heads(first: str, amount: str) -> str:
    return ColumnHeads([('#', False), ('', False), (first, False), ('', False), (amount, True), (RC['percentileLabel'], True), ('Date', True),
                        ('', False)])


def landing_rows(rows: list[dict], heads: str) -> str:
    """RankingsLanding's compact list: two figures (the value, the percentile)."""
    # The landing's rank stays plain (no legend explains a colour there): the percentile carries it (lib/rankings/landing.ts).
    body = ''.join(RankRow(escape(v['rank']), crest_img(v['classSlug'], 28),
                           CLASS_COLOR[v['classSlug']], escape(v['name']), escape(v['sub']), v['fraction'],
                           [(f'<a href="#">{escape(f["value"])}</a>' if f.get('href') else escape(f['value']), escape(f.get('unit') or ''),
                             escape(f['label']), False) for f in v['figures']], escape(v['gap']['text']), []) for v in rows)
    return f'<div class="lrows">{RankList(heads, body, 2, compact=True, fig_px=72)}</div>'


def landing() -> str:
    """The landing: each encounter (gold name, zone and size, kills) over its top three; then the spec leaders, damage and healing."""
    encs = FACTS['encounters']
    blocks = ''.join(f'<div class="enc-b"><div class="encs">{Rows([Row(f"<span class=gold>{escape(e["name"])}</span>", escape(FACTS["fns"]["where"][e["slug"]]), trail=f"<span class=num>{escape(FACTS["fns"]["ranked"][e["slug"]])} <span class=gold>&rsaquo;</span></span>")])}</div>'
                     f'{landing_rows(FACTS["landing"]["top"][e["slug"]], landing_heads("Character", "DPS") if i == 0 else "")}</div>' for i, e in enumerate(encs))
    picker = Section(escape(EC['heading']), blocks, panel=True, frame='lead', flush=True)
    leader = next((e for e in encs if (e['size'] or 0) >= 10), encs[0])
    lead = FACTS['landing']['leaders']
    body = (landing_rows(lead['damage'], landing_heads(escape(FACTS['fns']['damageHead']), 'DPS'))
            + landing_rows(lead['healing'], landing_heads(escape(FACTS['fns']['healingHead']), 'HPS')))
    leaders = Section(escape(FACTS['fns']['leadersHeading']), body, meta=escape(leader['name']), panel=True, frame='plain', flush=True)
    return picker + leaders


def index_page(empty: bool) -> str:
    if empty:
        main = Section(escape(EC['heading']), empty_block(EC['noneYet'], EC['fill']), panel=True, frame='lead', flush=True)
        meta = HIDDEN_STAMP
    else:
        main, meta = landing(), UPDATED   # the landing's boards answer, so the band carries the stamp
    rail = (companion() if empty else '') + pointer()
    html = head(meta) + bar([board_tabs('character')], '') + Page('split', main, rail, panels=True)
    return page('Rankings', html, None, f'/rankings mock, {"empty" if empty else "index"}. Data: rankings-fixture.ts ranked encounters.')


def board_page(state: str) -> str:
    """state: board | empty | you | off | unranked."""
    if state in ('board', 'empty'):
        enc, cohort = 'Warden Kelthas', RC['everyClassAndSpec'] if state == 'board' else 'Fury Warriors'
        b = FACTS['board_signed_out']
        if state == 'board':
            body = board_list(b)
            sec = Section(escape(cohort), body, meta=f'{b["total"]} ranked', panel=True, frame='lead', flush=True)
            cbar = character_bar(enc, RC['everyClass'], RC['everySpec'], RC['allRulesets'], f'{enc} &middot; {RC["everyClass"]}')
        else:
            onyxia = [e for e in FACTS['encounters'] if e['slug'] == 'onyxia']
            elsewhere = f'<p class="t-label">{escape(EMPTY["elsewhere"])}</p>' + encounter_rows(onyxia)
            body = empty_block(FACTS['fns']['boardEmpty'], EMPTY['boardFill'], widen=True, elsewhere=elsewhere)
            sec = Section(escape(cohort), body, panel=True, frame='lead', flush=True)
            cbar = character_bar(enc, 'Warrior', 'Fury', RC['allRulesets'], f'{enc} &middot; Fury')
        rail = (companion() if state == 'empty' else '') + pointer()
        html = head(UPDATED) + cbar + Page('split', sec, rail, panels=True)
        return page(f'{enc} rankings', html, None, f'/rankings/<slug> mock, {state}. Data: rankings-fixture.ts boardFor().')
    key = {'you': 'board_you', 'off': 'board_you_off_page', 'unranked': 'board_you_unranked'}[state]
    b = FACTS[key]
    widen = Link(escape(RC['widen']) + ' &rsaquo;')
    # Your rank says "14th of 212 Fury Warriors" beside the board: the head's count is left to it (Rankings.svelte).
    meta = '' if state in ('you', 'off') else f'{b["total"]} ranked'
    sec = Section('Fury Warriors', board_list(b, FACTS['own_key']), action=widen, meta=meta,
                  panel=True, frame='lead', flush=True)
    answer = FACTS[{'you': 'answer_you', 'off': 'answer_off', 'unranked': 'answer_unranked'}[state]]
    rail = your_rank(answer, unranked=state == 'unranked')
    html = head(UPDATED) + character_bar('Onyxia', 'Warrior', 'Fury', 'Normal', 'Onyxia &middot; Fury') + Page('split', sec, rail, panels=True)
    return page('Onyxia rankings', html, YOU, f'/rankings mock, {state}. Data: rankings-fixture.ts boardFor().')


def guild_page() -> str:
    heads = ColumnHeads([('#', False), ('', False), ('Guild', False), ('', False), ('Kill time', True), ('Date', True), ('', False)])
    rows = ''.join(RankRow(g['rank'], '', 'var(--accent)', escape(g['name']), escape(g['sub']), g['fraction'],
                           [figure_cell(f) for f in g['figures']], escape(g['gap']['text']), [])
                   for g in FACTS['guilds'])
    rl = RankList(heads, rows, 1, compact=True)
    sec = Section(escape(RC['guildTitle']['speed']), rl, panel=True, frame='lead', flush=True)
    cbar = bar([board_tabs('guild'), (RC['encounterLabel'], sel('Warden Kelthas')), ('Kind', sel('Speed'))], 'Warden Kelthas &middot; Fastest kills')
    html = head(HIDDEN_STAMP) + cbar + Page('split', sec, pointer(), panels=True)
    return page('Warden Kelthas rankings', html, None, '/rankings/<slug> mock, guilds (Speed). Data: rankings-fixture.ts guildSpeedRows().')


PAGES = {
    'web-rankings': lambda: index_page(False),
    'web-rankings-empty': lambda: index_page(True),
    'web-rankings-board': lambda: board_page('board'),
    'web-rankings-board-empty': lambda: board_page('empty'),
    'web-rankings-guilds': guild_page,
    'web-rankings-you': lambda: board_page('you'),
    'web-rankings-off': lambda: board_page('off'),
    'web-rankings-unranked': lambda: board_page('unranked'),
}

if __name__ == '__main__':
    if sys.argv[1:] == ['--list']:
        print(' '.join(PAGES))
    else:
        only = sys.argv[1:]
        for name, fn in PAGES.items():
            if not only or name in only:
                print(write(f'{name}.html', fn()))
