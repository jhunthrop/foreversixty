"""Web-layout parity mock for /ratings (lane P5, plan 2026-11-09-redesign), composed ONLY from ws_prims.

The page explains the Forever Score and the Performance Rating (components/score/RatingsPage.svelte). Round 5 sets it in the
logs lane's panel form (the directions logs mocks' prims.panel; ws_prims.Section(panel=True)): the Forever Score is the
page's framed panel (the parts as a table, under the sample's note), the Performance Rating a plain panel (three figures),
then the six components with their weights by role and what each measures, and how a fight is judged as Rows (lane
h-report-tail, directions mock report_tail.py b1-ratings); the rail holds the tiers (in the game's item quality colours, the
sample's with the points to the next) and how logs get here.

Data and copy, read at render time, never retyped: design/mocks/data/rankings-facts.json (node design/mocks/rankings_facts.mjs):
lib/score/ratings-page.ts (the sample, the trend, the weights and the page's prose), lib/score and lib/rating copy.

  node design/mocks/rankings_facts.mjs && python3 design/mocks/ws_ratings.py
"""
import json
import os
import sys
from html import escape
from pathlib import Path

from ws_prims import (Button, Figure, Meter, Page, PageHeader, Row, Rows, Section, StatRow, StatTable, TabStrip, build_page, write)

MOCKDATA = Path(__file__).resolve().parent / 'data'
FACTS = json.loads(Path(os.environ.get('FS_RANKINGS_FACTS') or MOCKDATA / 'rankings-facts.json').read_text())
SC, RC = FACTS['copy']['scoreCopy'], FACTS['copy']['ratingCopy']
R = FACTS['ratings']
TIER_HEX = {'common': '#ffffff', 'uncommon': '#1eff00', 'rare': '#0070dd', 'epic': '#a335ee', 'legendary': '#ff8000'}   # --color-rarity-*

PAGE_CSS = '''<style>
.st-r .nm{line-height:20px}
.parts{margin-top:0}
.figs3{display:grid;grid-template-columns:repeat(3,minmax(0,1fr));gap:20px 24px}.figs3>div{display:flex;flex-direction:column;gap:2px;min-width:0}
.tier{font-size:19px;line-height:24px}.tiers .row.cur{background:color-mix(in srgb,var(--c) 12%,transparent);box-shadow:inset 3px 0 0 var(--c)}
</style>'''


def head() -> str:
    score = R['sample']['score']
    tier = R['sample']['tier']
    total = sum(R['component_max'].values())
    labels = {c['key']: c['label'] for c in SC['components']}
    gaps = sorted(((R['component_max'][k] - v, k) for k, v in R['sample']['components'].items()), reverse=True)
    key = gaps[0][1]
    # The answer is a Figure (big, in the tier's colour) and one short line under it (RatingsPage.svelte), no spacer column.
    fig = Figure(f'{score:,}', f'{escape(SC["samplePill"])} &middot; {escape(SC["explainerHeading"])}',
                 f'of {total:,} &middot; {tier.capitalize()}', size='big').replace('<span class="v">', f'<span class="v" style="color:{TIER_HEX[tier]}">', 1)
    ans = fig + f'<span class="t-small">{escape(SC["mostToGainLead"])} {labels[key]}{escape(R["most_to_gain_rest"][key])}</span>'
    return PageHeader(escape(SC['explainerPageTitle']), meta=[escape(SC['explainerScoreCaption']), escape(SC['explainerRatingCaption'])],
                      answer=ans, tabs=TabStrip([SC['explainerHeading'], SC['explainerRatingHeading']], SC['explainerHeading']))


def score_panel() -> str:
    parts = R['copy']['parts']
    rows = [StatRow('', f'<b>{escape(c["label"])}</b>', escape(parts[c['key']]),
                    [f'{R["sample"]["components"][c["key"]]:,}', f'{R["component_max"][c["key"]]:,}'],
                    Meter(R['sample']['components'][c['key']] / R['component_max'][c['key']]),
                    primary_phone=f'of {R["component_max"][c["key"]]}') for c in SC['components']]
    table = StatTable([('', False), ('Part', False), ('Of its top', False), ('Points', True), ('Top', True)], rows, 2, open_=True)
    body = f'<div class="parts">{table}</div>' + f'<p class="t-small muted" style="padding-top:12px">{escape(SC["explainerLaunch"])}</p>'
    return Section(escape(SC['explainerHeading']), body, meta=escape(R['copy']['sampleNote']), panel=True, frame='framed')


def rating_panel() -> str:
    trend = R['trend']
    strong = 'font:600 15px/20px var(--font-body);color:var(--strong)'
    figs = (f'<div class="figs3">{Figure(str(trend[-1]), "Latest fight", escape(RC["outOf100"]), tone="gold", size="big")}'
            f'<div><span class="t-label">Last {len(trend)} fights</span><span style="{strong}">{" &middot; ".join(map(str, trend))}</span>'
            f'<span class="t-small">Up {trend[-1] - trend[0]} since the first</span></div>'
            f'<div><span class="t-label">Weakest part</span><span style="{strong}">{escape(R["weakest"])}</span></div></div>')
    body = figs + f'<p class="t-small muted" style="padding-top:16px">{escape(R["copy"]["visibility"])} {escape(RC["cappedRule"])}</p>'
    return Section(escape(SC['explainerRatingHeading']), body, panel=True)


def reference_panels() -> str:
    """The six components with their weights by role and what each measures (one table), then how a fight is judged as Rows."""
    pc = R['copy']
    measures = {c['part']: c['body'] for c in pc['components']}
    wrows = [StatRow('', f'<b>{w["component"]}</b>', escape(measures.get(w['component'], '')), [str(w['dps']), str(w['healer']), str(w['tank'])],
                     primary_phone=f'Healer {w["healer"]} &middot; Tank {w["tank"]}', bar=False) for w in R['weights']]
    weights = StatTable([('', False), ('Component', False), ('DPS', True), ('Healer', True), ('Tank', True)], wrows, 3, bar=False)
    rules = Rows([Row(escape(r['title']), escape(r['body']), link=False) for r in pc['rules']], open_=True)
    return (Section('Weights, by role', weights, meta=escape(pc['adjustable']), panel=True, flush=True)
            + Section('How a fight is judged', rules, panel=True))


def to_next() -> str:
    """The sample's row in the Tiers rail: the points to the next tier up ("170 to Legendary")."""
    score = R['sample']['score']
    nxt = next((t for t, start in R['tier_from'].items() if start > score), None)
    return f'{R["tier_from"][nxt] - score:,} to {nxt.capitalize()}' if nxt else ''


def rail() -> str:
    tier = R['sample']['tier']
    nxt = to_next()
    sample = f'<span class="t-small" style="text-align:right">The sample{f"<br><span>{nxt}</span>" if nxt else ""}</span>'
    rows = [Row(f'<span class="tier" style="color:{TIER_HEX[t["tier"]]}">{t["tier"].capitalize()}</span>', escape(t['range']), link=False,
                selected=t['tier'] == tier, trail=sample if t['tier'] == tier else '')
            .replace('<div class="row', f'<div style="--c:{TIER_HEX[t["tier"]]}" class="row', 1) for t in R['tiers']]
    tiers = Section(f'{SC["tierColumn"]}s', f'<div class="tiers">{Rows(rows, open_=True)}</div>', panel=True)
    logs = Section(escape(SC['explainerLogsHeading']), f'<p class="t-small">{escape(SC["explainerLogsBody"])}</p><p style="margin-top:12px">{Button(escape(SC["explainerLogsButton"]))}</p>',
                   panel=True)
    return tiers + logs


def ratings_page() -> str:
    body = head() + Page('split', score_panel() + rating_panel() + reference_panels(), rail(), panels=True) + PAGE_CSS
    return build_page(SC['explainerPageTitle'], body, 'Logs', None, '/ratings mock. Data: lib/score/ratings-page.ts; lib/score and lib/rating copy.')


PAGES = {'web-ratings': ratings_page}

if __name__ == '__main__':
    if sys.argv[1:] == ['--list']:
        print(' '.join(PAGES))
    else:
        for name, fn in PAGES.items():
            print(write(f'{name}.html', fn()))
