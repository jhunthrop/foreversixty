"""Web-layout parity mock for /ratings (lane P5, plan 2026-11-09-redesign), composed ONLY from ws_prims.

The page explains the Forever Score and the Performance Rating (components/score/RatingsPage.svelte). Round 5 sets it in the
logs lane's panel form (the directions logs mocks' prims.panel; ws_prims.Section(panel=True)): the Forever Score is the
page's lead panel, the Performance Rating framed, the reference sections plain panels; the rail holds the tiers (in the
game's item quality colours, so on their panel) and how logs get here.

Data and copy, read at render time, never retyped: design/mocks/data/rankings-facts.json (node design/mocks/rankings_facts.mjs):
lib/score/ratings-page.ts (the sample, the trend, the weights and the page's prose), lib/score and lib/rating copy.

  node design/mocks/rankings_facts.mjs && python3 design/mocks/ws_ratings.py
"""
import json
import os
import sys
from html import escape
from pathlib import Path

from ws_prims import (Button, Figure, Meter, Page, PageHeader, Row, Rows, Section, Spark, StatRow, StatTable, TabStrip, build_page, write)

MOCKDATA = Path(__file__).resolve().parent / 'data'
FACTS = json.loads(Path(os.environ.get('FS_RANKINGS_FACTS') or MOCKDATA / 'rankings-facts.json').read_text())
SC, RC = FACTS['copy']['scoreCopy'], FACTS['copy']['ratingCopy']
R = FACTS['ratings']
TIER_HEX = {'common': '#ffffff', 'uncommon': '#1eff00', 'rare': '#0070dd', 'epic': '#a335ee', 'legendary': '#ff8000'}   # --color-rarity-*

PAGE_CSS = '''<style>
.prose{display:flex;flex-direction:column;gap:12px;padding-top:12px}.prose:first-child{padding-top:0}.prose p{font-size:15px;line-height:24px;max-width:72ch}.prose p.muted{color:var(--muted)}
.st-r .nm{line-height:20px}
.parts{margin-top:0}
/* the built open Section sets its body 12 px under the hairline (ui/Section.svelte pt-3) */
section>.sec-h{padding-bottom:10px;border-bottom:1px solid rgba(255,255,255,.14)}section>.sec-h+*,section>.sec-h+.rows{margin-top:12px}section>.sec-h+.prose{padding-top:0}
.figs3{display:grid;grid-template-columns:repeat(3,minmax(0,1fr));gap:20px 24px}.figs3>div{display:flex;flex-direction:column;gap:2px;min-width:0}
.tier{font-size:19px;line-height:24px}.tiers .row.cur{background:color-mix(in srgb,var(--c) 12%,transparent);box-shadow:inset 3px 0 0 var(--c)}
</style>'''


def prose(paras: list[str], muted: list[bool] | None = None) -> str:
    muted = muted or []
    return '<div class="prose">' + ''.join(f'<p class="{"muted" if i < len(muted) and muted[i] else ""}">{escape(p)}</p>'
                                           for i, p in enumerate(paras)) + '</div>'


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
    rows = [StatRow('', f'<b>{escape(c["label"])}</b>', escape(c['moves']),
                    [f'{R["sample"]["components"][c["key"]]:,}', f'{R["component_max"][c["key"]]:,}'],
                    Meter(R['sample']['components'][c['key']] / R['component_max'][c['key']]),
                    primary_phone=f'{R["sample"]["components"][c["key"]]} of {R["component_max"][c["key"]]}') for c in SC['components']]
    table = StatTable([('', False), ('Part', False), ('Of its top', False), ('Points', True), ('Top', True)], rows, 2)
    body = f'<div class="parts">{table}</div>' + prose([SC['explainerLaunch']], [True])
    return Section(escape(SC['explainerHeading']), body, panel=True, frame='lead')


def rating_panel() -> str:
    trend = R['trend']
    figs = (f'<div class="figs3"><div><span class="t-label">Latest fight</span><span class="v" style="font:700 36px/40px var(--font-body);color:var(--gold)">'
            f'{trend[-1]}<small class="muted" style="font-size:13px;margin-left:4px">{escape(RC["outOf100"])}</small></span></div>'
            f'<div><span class="t-label">Last {len(trend)} fights</span>{Spark(trend, "#e5b955", total=100)}'
            f'<span class="t-small">Up {trend[-1] - trend[0]} since the first</span></div>'
            f'<div><span class="t-label">Weakest part</span><span style="font:600 15px/20px var(--font-body);color:var(--strong)">{escape(R["weakest"])}</span></div></div>')
    pc = R['copy']
    body = figs + prose([pc['visibility'], *pc['oneNumberFirst']], [True])
    return Section(escape(SC['explainerRatingHeading']), body, panel=True, frame='framed')


def reference_panels() -> str:
    pc = R['copy']
    comp = Rows([Row(escape(c['part']), escape(c['body']), link=False) for c in pc['components']], open_=True)
    wrows = [StatRow('', f'<b>{w["component"]}</b>', '', [str(w['dps']), str(w['healer']), str(w['tank'])],
                     primary_phone=f'{w["dps"]} / {w["healer"]} / {w["tank"]}', bar=False) for w in R['weights']]
    weights = StatTable([('', False), ('Component', False), ('DPS', True), ('Healer', True), ('Tank', True)], wrows, 3, bar=False, open_=True)
    # the lists boxed under their heading on the ground, the prose open (ai-slop round 1: seven identical panels read as a
    # template; ux round 2: small text on the bare ground)
    return (Section('What each component measures', comp, boxed=True)
            + Section('Weights, by role', weights + prose([pc['roles'], pc['adjustable']], [True, True]), boxed=True)
            + Section('Ranked against your spec, or absolute', prose(pc['rankedOrAbsolute']))
            + Section('The cap', prose([RC['cappedRule']])))


def rail() -> str:
    tier = R['sample']['tier']
    rows = [Row(f'<span class="tier" style="color:{TIER_HEX[t["tier"]]}">{t["tier"].capitalize()}</span>', escape(t['range']), link=False,
                selected=t['tier'] == tier, trail='<span class="t-small">The sample</span>' if t['tier'] == tier else '')
            .replace('<div class="row', f'<div style="--c:{TIER_HEX[t["tier"]]}" class="row', 1) for t in R['tiers']]
    tiers = Section(f'{SC["tierColumn"]}s', f'<div class="tiers">{Rows(rows, open_=True)}</div><p class="t-small" style="margin-top:10px">{escape(SC["tiersNote"])}</p>',
                    panel=True, frame='framed')
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
