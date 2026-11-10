"""Web-layout mock: the tier list (/tiers, /tiers/tank, /tiers/healer), composed ONLY from ws_prims primitives.

Look: the directions mock g-tiers (round 5); rules: design/specs/2026-11-06-tiers.md. Every figure is read at render time from data/builds/<active build>/bis/<spec>.json
(band 60, preset "raid", the chosen faction) and ranked by the shipped rules of web/src/lib/tiers/ (tier-list.ts, tier-rules.ts):
DPS sorted on set_dps with S/A/B/C/D at 5/10/20/30% behind the top and a 1% tie margin; tanks on dtps, lower is better;
healers on set_dps (= effective HPS) with ties wherever the error (2 standard errors of the difference) cannot order a pair
at the tuned profile and both tuning bounds. Nothing here is typed in except the sample signed-in character's name.

  FS_DATA_ROOT=<a checkout whose active build is the live one> python3 design/mocks/ws_tiers.py
  node design/mocks/ws_shoot.mjs web-tiers 390,1440,1920,2000
FS_DATA_ROOT defaults to this repo; the 2026-11-06 renders used an origin/main worktree (build 1.60.1.70790, the live build).
"""
import json
import math
import os
import sys
from datetime import datetime
from pathlib import Path

from mocklib import CLASS_COLOR
from ws_lib import REGION
from ws_prims import (Button, ColumnHeads, Frame, Link, Page, PageHeader, Plate, RankGroup, RankList, RankRow, Skeleton,
                      build_page, control_bar, crest_img, seg, write)

DATA_ROOT = Path(os.environ.get('FS_DATA_ROOT', Path(__file__).resolve().parents[2]))
BUILD = json.loads((DATA_ROOT / 'web/src/data/active-build.json').read_text())['build']
BIS = DATA_ROOT / 'data/builds' / BUILD / 'bis'
CATALOG = json.loads((DATA_ROOT / 'data/curated/specs.json').read_text())
CLASS_NAME = {c['slug']: c['name'] for c in json.loads((DATA_ROOT / 'data/builds' / BUILD / 'classes.json').read_text())}

# The shipped rules (web/src/lib/tiers/tier-rules.ts, tier-copy.ts), mirrored, not re-decided.
TIERS = (('S', 0), ('A', 5), ('B', 10), ('C', 20), ('D', 30))
HEALER_Z = 2
DPS_FIGHT_SECONDS = 180
TANK_BOSS_LEVEL = 63
ROLE_LABEL = {'dps': 'DPS', 'tank': 'Tank', 'healer': 'Healer'}
FACTION_LABEL = {'alliance': 'Alliance', 'horde': 'Horde'}
YOU = ('Thandrel', 'warrior', 'Fury Warrior &middot; 60')   # sample signed-in character
YOU_SPEC = 'warrior-fury'


# ------------------------------------------------------------------------------------------------ data
def band_of(f: dict, faction: str) -> dict | None:
    return next((b for b in f['bands'] if int(b['band']) == 60 and b['preset'] == 'raid' and b['faction'] == faction), None)


def load(faction: str) -> tuple[list[dict], dict]:
    rows, files = [], []
    for s in CATALOG:
        p = BIS / f"{s['spec']}.json"
        if not p.exists():
            continue
        f = json.loads(p.read_text())
        b = band_of(f, faction)
        if b is None:
            continue
        files.append(f)
        rows.append({'spec': s['spec'], 'cls': s['class_slug'], 'name': s['name'], 'role': b.get('role') or 'dps',
                     'race': ' '.join(w.capitalize() for w in b['race'].split('-')), 'set': b['set_dps'],
                     'm': b.get('metrics') or {}, 'low': bool(b.get('weights_low_confidence'))})
    newest = max(files, key=lambda f: f['generated_at'])
    heal = next((f['heal_profile']['duration_sec'] for f in files if f.get('heal_profile')), None)
    meta = {'updated': datetime.strptime(newest['generated_at'][:10], '%Y-%m-%d'), 'preset': files[0]['presets']['raid']['label'], 'heal_sec': heal}
    return rows, meta


def list_label(role: str, faction: str) -> str:
    """The Plate's accessible name (tier-copy.ts listLabel)."""
    return f'{ROLE_LABEL[role]} tier list, {FACTION_LABEL[faction]}'


def heal_points(r: dict):
    m = r['m']
    if 'hps_error' not in m or not m.get('hps_at_bounds'):
        return None
    return [(r['set'], m['hps_error'])] + [(x['hps'], x['hps_error']) for x in m['hps_at_bounds']]


def ranked(rows: list[dict], role: str) -> tuple[list[dict], bool]:
    pool = [r for r in rows if r['role'] == role]
    tank = role == 'tank'
    metric = (lambda r: r['m']['dtps']) if tank else (lambda r: r['set'])
    pool.sort(key=lambda r: (metric(r) if tank else -metric(r), r['name']))
    top = metric(pool[0])
    is_ranked = role != 'healer' or all(heal_points(r) is not None for r in pool)
    out = []
    for i, r in enumerate(pool):
        v = metric(r)
        gap, frac = ((v / top - 1) * 100, top / v) if tank else ((1 - v / top) * 100, v / top)
        out.append({**r, 'metric': v, 'gap': 0 if i == 0 else gap, 'frac': frac, 'rank': i + 1, 'tie': None})
    # A tied healer shares its rank ("=2"); the DPS list's 1% ties are no longer drawn (round 5), so only healers are marked.
    if is_ranked and role == 'healer':
        def tied(a, b):
            pa, pb = heal_points(a), heal_points(b)
            if pa is None or pb is None or len(pa) != len(pb):
                return True
            return not all(x[0] - y[0] > HEALER_Z * math.hypot(x[1], y[1]) for x, y in zip(pa, pb))
        above = [i > 0 and tied(out[i - 1], out[i]) for i in range(len(out))]
        for i, r in enumerate(out):
            if above[i]:
                r['rank'] = out[i - 1]['rank']
            below = i + 1 < len(out) and above[i + 1]
            r['tie'] = 'both' if above[i] and below else ('above' if above[i] else ('below' if below else None))
    if role == 'dps':
        for r in out:
            g = round(r['gap'] * 1e6) / 1e6
            r['tier'] = [t for t, f in TIERS if g >= f][-1]
    if tank:
        best_eh = max(r['m']['effective_health'] for r in out)
        best_tps = max(r['m']['tps'] for r in out)
        for r in out:
            r['best_eh'] = r['m']['effective_health'] == best_eh
            r['best_tps'] = r['m']['tps'] == best_tps
    return out, is_ranked


def fnum(v: float, role: str) -> str:
    return f'{v:,.0f}' if role == 'tank' else f'{v:,.1f}'


def full_name(r: dict) -> str:
    return f"{r['name']} {CLASS_NAME[r['cls']]}"


# ------------------------------------------------------------------------------------------------ pieces
# Round 5 (directions mock g-tiers, design/mocks/directions/pages.py `tiers`): the title band with the fight's facts, the role and
# faction tabs, then a split page: the compact list on the Plate (S to D on the damage list, the name the row's one link, the
# accent bar from 90% of the lowest spec, the figure under a head that names its unit, "vs #1") and the spec panel in the rail.
TAB_LABEL = {'dps': 'Damage', 'tank': 'Tanks', 'healer': 'Healers'}
UNIT = {'dps': 'DPS', 'healer': 'HPS', 'tank': 'taken / s'}
BAR_FLOOR = 0.9  # tier-row-view.ts barFraction: the bar starts at 90% of the lowest spec


def bar_frac(r: dict, rows: list[dict]) -> float:
    floor = min(x['frac'] for x in rows) * BAR_FLOOR
    return (r['frac'] - floor) / (1 - floor)


def row_html(r: dict, rows: list[dict], ranked_: bool, you: bool = False, state: str = '') -> str:
    role = r['role']
    if role == 'tank':
        figs = [(fnum(r['metric'], role), '', '', r['rank'] == 1),
                (fnum(r['m']['effective_health'], role), '', 'Effective health', r['best_eh']),
                (fnum(r['m']['tps'], role), '', 'Threat', r['best_tps'])]
        lead = r['rank'] == 1
        gap = '' if lead else f"+{r['gap']:.1f}%"
    else:
        figs = [(fnum(r['metric'], role), '', '', False)]
        lead = ranked_ and r['rank'] == 1 and r['gap'] == 0
        if not ranked_:
            gap = f"&minus;{r['gap']:.1f}%" if r['gap'] else '0.0%'
        else:
            gap = '' if lead else f"&minus;{r['gap']:.1f}%"
    rank = '' if not ranked_ else (f"={r['rank']}" if role == 'healer' and r['tie'] else str(r['rank']))
    tie = healer_tie(r, rows) if ranked_ else ''
    return RankRow(rank, crest_img(r['cls'], 36), CLASS_COLOR[r['cls']], full_name(r), '', bar_frac(r, rows), figs, gap, [],
                   gap_top=lead, low=' &middot; '.join(x for x in (tie, 'Stat weights less certain' if r['low'] else '') if x), selected=you, state=state, accent=True)


def healer_tie(r: dict, rows: list[dict]) -> str:
    """tier-copy.ts healerTie: a tied healer names its neighbours in full ("&asymp; tie with Restoration Shaman")."""
    if r['role'] != 'healer' or not r['tie']:
        return ''
    i = rows.index(r)
    mates = [rows[j] for j in {'above': [i - 1], 'below': [i + 1], 'both': [i - 1, i + 1]}[r['tie']]]
    names = [full_name(m) for m in mates]
    joined = names[0] if len(names) == 1 else f"{', '.join(names[:-1])} and {names[-1]}"
    return f'&asymp; tie with {joined}'


def heads(role: str) -> str:
    if role == 'tank':
        return ColumnHeads([('#', True), ('', False), ('Spec', False), ('', False), ('Taken / s', True), ('Health', True),
                            ('Threat / s', True), ('vs #1', True), ('', False)])
    return ColumnHeads([('#', True), ('', False), ('Spec', False), ('', False), ('DPS' if role == 'dps' else 'HPS', True),
                        ('vs #1', True), ('', False)])


def the_list(rows: list[dict], role: str, ranked_: bool, you_spec: str = '', hover_spec: str = '') -> str:
    def one(r):
        return row_html(r, rows, ranked_, r['spec'] == you_spec, 'hover' if r['spec'] == hover_spec else '')
    fig_px = 96 if role == 'tank' else 88
    if role == 'dps':
        body = ''.join(RankGroup(letter, '', '', '', [one(r) for r in rows if r['tier'] == letter])
                       for letter, _ in TIERS if any(r['tier'] == letter for r in rows))
        return RankList(heads(role), body, 1, grouped=True, compact=True, fig_px=fig_px)
    return RankList(heads(role), ''.join(one(r) for r in rows), 3 if role == 'tank' else 1, compact=True, fig_px=fig_px)


def fight_length(seconds: int) -> str:
    return f'{seconds // 60} min' if seconds % 60 == 0 else f'{seconds} s'


def meta_line(role: str, meta: dict) -> str:
    """tier-copy.ts metaLine, then the Updated stamp."""
    fight = ([f'level {TANK_BOSS_LEVEL}'] if role == 'tank'
             else [fight_length(meta['heal_sec'])] if role == 'healer'
             else [f'level {TANK_BOSS_LEVEL}', fight_length(DPS_FIGHT_SECONDS)])
    return ' &middot; '.join(['Best in slot, raid buffs', *fight, f'Updated {meta["updated"].strftime("%-d %b")}'])


def bar(role: str, faction: str, open_m: bool = False) -> str:
    roles = seg([(TAB_LABEL[r], r == role, None, '') for r in ('dps', 'tank', 'healer')])
    facs = seg([(FACTION_LABEL[f], f == faction, None, '') for f in ('alliance', 'horde')])
    return control_bar([('Role', roles), ('Faction', facs)], f'{TAB_LABEL[role]} &middot; {FACTION_LABEL[faction]}', open_m=open_m,
                       labels_from=1920)


def card(role: str, lists: dict, signed_in: bool, state: str = 'ready') -> str:
    """The spec panel in the rail (TierSpecCard.astro). state: ready | loading | unknown."""
    if not signed_in:   # a quiet link, no panel (TierSpecCard.astro)
        return f'<p class="tier-quiet">{Link("Sign in to mark your spec &rsaquo;")}</p>'
    if state == 'loading':
        return f'<div class="sk">{Skeleton(180, 28)}{Skeleton(220, 24)}{Skeleton(230, 44)}</div>'
    if state == 'unknown':
        return (f'<h2>Your spec</h2><p class="cl">Thandrel&rsquo;s spec is not known yet.</p>'
                f'<div class="ca">{Button("Choose a character")}</div>')
    me_role = next(rl for rl, lst in lists.items() if any(r['spec'] == YOU_SPEC for r in lst[0]))
    me = next(r for r in lists[me_role][0] if r['spec'] == YOU_SPEC)
    if me_role != role:
        return (f'<h2>{full_name(me)}</h2><p class="cl">{Link(f"On the {TAB_LABEL[me_role]} list &rsaquo;")}</p>'
                f'<div class="ca">{Button("Best in slot", True)}{Button("Guide")}</div>')
    # The row is marked on the list beside it: no standing line repeats its tier, place and figure.
    return (f'<div class="cb"><h2 style="color:var(--class-{YOU[1]})">{full_name(me)}</h2>'
            f'<div class="ca">{Button("Best in slot", True)}{Button("Guide")}</div></div>')


PAGE_CSS = '''<style>
/* page-level layout only (no surfaces): the spec panel's type and rhythm (TierSpecCard.astro) */
.frame h2{font:400 22px/28px var(--font-title);letter-spacing:-.015em}
.frame .cl{margin-top:6px;font:600 15px/24px var(--font-body);color:var(--strong)}
.frame .ca{display:flex;flex-wrap:wrap;gap:12px;margin-top:16px}.frame .btn{height:44px;display:inline-flex;align-items:center}
.frame .sk,.frame .cb{display:flex;flex-direction:column;gap:6px;min-height:114px}.frame .cb .ca{margin-top:10px}
.tier-quiet{margin:0;font:400 15px/24px var(--font-body)}
.tiers-end{padding-bottom:var(--section-gap)}
.statesheet{display:grid;gap:32px}.statesheet h2{margin-bottom:10px}
</style>'''


def tiers_page(role: str, faction: str = 'alliance', signed_in: bool = False, open_m: bool = False, answer_state: str = 'ready',
               hover_spec: str = '') -> str:
    data = {f: load(f) for f in ('alliance', 'horde')}
    rows, meta = data[faction]
    lists = {r: ranked(rows, r) for r in ('dps', 'tank', 'healer')}
    lst, ranked_ = lists[role]
    header = PageHeader('Tier list at 60', meta=[meta_line(role, meta)])
    plate = Plate(the_list(lst, role, ranked_, YOU_SPEC if signed_in else '', hover_spec), list_label(role, faction))
    inner = card(role, lists, signed_in, answer_state)
    rail = (Frame(inner, False, REGION['head_answer']) if signed_in
            else f'<div data-region="{REGION["head_answer"]}">{inner}</div>')
    body = header + bar(role, faction, open_m) + f'<div class="tiers-end">{Page("split", plate, rail)}</div>' + PAGE_CSS
    note = (f'Tier list mock (directions mock g-tiers, round 5), built only from ws_prims. Data: build {BUILD}, band 60, preset raid, '
            f'{faction}; ranked by the shipped tier rules. Signed-in character "Thandrel" is a sample.')
    return build_page(f'{ROLE_LABEL[role]} tier list', body, 'Tools', YOU if signed_in else None, note)


def states_page() -> str:
    """Every state the build must show, on one sheet (rendered at 1440 and 390)."""
    rows, _ = load('alliance')
    lists = {r: ranked(rows, r) for r in ('dps', 'tank', 'healer')}
    dps, _ = lists['dps']
    pick = lambda spec: next(r for r in dps if r['spec'] == spec)
    low = next((r for r in dps if r['low']), None)
    sample = [('Default', row_html(dps[1], dps, True)),
              ('Hover (the lift, the rail slice in the class colour, the bar 18% brighter)', row_html(dps[1], dps, True, state='hover')),
              ('Your spec (the accent tint and a cyan rail slice)', row_html(pick(YOU_SPEC), dps, True, you=True))]
    if low:
        sample.append(('Stat weights less certain', row_html(low, dps, True)))
    heal = lists['healer'][0]
    blocks = ''.join(f'<section><h2 class="t-h3">{t}</h2>{Plate(RankList(heads("dps"), h, 1, compact=True, fig_px=88))}</section>' for t, h in sample)
    blocks += (f'<section><h2 class="t-h3">Healer list published without the measured error: no ranks</h2>'
               f'{Plate(RankList(heads("healer"), row_html(heal[1], heal, False), 1, compact=True, fig_px=88))}</section>')
    blocks += f'<section><h2 class="t-h3">No data for this build</h2>{Plate("<p class=\'empty\'>No tier list is published for this build yet.</p>")}</section>'
    blocks += f'<section><h2 class="t-h3">A role with no simmed spec</h2>{Plate("<p class=\'empty\'>No Tank specs are simmed yet.</p>")}</section>'
    for t, st, role, signed in (('Spec panel: signed out', 'ready', 'dps', False), ('Spec panel: loading (reserved at its final size)', 'loading', 'dps', True),
                                ('Spec panel: spec not known', 'unknown', 'dps', True),
                                ('Spec panel: your spec is on another role&rsquo;s list (tank page, Fury Warrior)', 'ready', 'tank', True)):
        blocks += f'<section><h2 class="t-h3">{t}</h2><div style="max-width:340px">{Frame(card(role, lists, signed, st), True)}</div></section>'
    body = '<div class="pc" style="padding-top:24px"><h1 class="t-title">Tier list: states</h1></div>' + Page('ledger', f'<div class="statesheet">{blocks}</div>') + PAGE_CSS
    return build_page('Tier list states', body, 'Tools', YOU, 'Tier list states sheet; data as web-tiers.')


PAGES = {
    'web-tiers': lambda: tiers_page('dps'),
    'web-tiers-you': lambda: tiers_page('dps', signed_in=True),
    'web-tiers-horde': lambda: tiers_page('dps', 'horde'),
    'web-tiers-healer': lambda: tiers_page('healer'),
    'web-tiers-tank': lambda: tiers_page('tank'),
    'web-tiers-tank-you': lambda: tiers_page('tank', signed_in=True),
    'web-tiers-bar': lambda: tiers_page('dps', open_m=True),
    'web-tiers-states': states_page,
}


def compare() -> None:
    """Before/after at the same scale against the live captures of 2026-11-06."""
    import ws_compose  # Pillow: only the compare step needs it; the parity job's python3 has the standard library alone
    ws_compose.BEFORE = Path(__file__).resolve().parents[2] / 'design/reviews/2026-11-06-tiers/before'
    for w, sc in ((1440, .6), (1920, .5), (2000, .5), (390, .9)):
        print(ws_compose.compose('tiers', 'tiers', 'web-tiers', w, sc, max_h=3200))
    for role in ('healer', 'tank'):
        for w, sc in ((1440, .6), (2000, .5), (390, .9)):
            print(ws_compose.compose(f'tiers-{role}', f'tiers-{role}', f'web-tiers-{role}', w, sc, max_h=3200))


if __name__ == '__main__':
    if sys.argv[1:] == ['compare']:
        compare()
    else:
        for name, fn in PAGES.items():
            print(write(f'{name}.html', fn()))
