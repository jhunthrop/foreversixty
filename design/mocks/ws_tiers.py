"""Web-layout mock: the tier list (/tiers, /tiers/tank, /tiers/healer), composed ONLY from ws_prims primitives.

Spec: design/specs/2026-11-06-tiers.md. Every figure is read at render time from data/builds/<active build>/bis/<spec>.json
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

from mocklib import CLASS_COLOR, emblem
from ws_prims import (ColumnHeads, Link, Notice, Page, PageHeader, Plate, RankGroup, RankList, RankRow, Section,
                      Skeleton, build_page, control_bar, crest_img, seg, write)
import ws_compose

DATA_ROOT = Path(os.environ.get('FS_DATA_ROOT', Path(__file__).resolve().parents[2]))
BUILD = json.loads((DATA_ROOT / 'web/src/data/active-build.json').read_text())['build']
BIS = DATA_ROOT / 'data/builds' / BUILD / 'bis'
CATALOG = json.loads((DATA_ROOT / 'data/curated/specs.json').read_text())
CLASS_NAME = {c['slug']: c['name'] for c in json.loads((DATA_ROOT / 'data/builds' / BUILD / 'classes.json').read_text())}
SWING = json.loads((DATA_ROOT / 'data/curated/tank-encounter.json').read_text())['boss']['swing_speed_sec']['value']

# The shipped rules (web/src/lib/tiers/tier-rules.ts, tier-copy.ts), mirrored, not re-decided.
TIERS = (('S', 0), ('A', 5), ('B', 10), ('C', 20), ('D', 30))
TIER_COLOUR = {'S': '#ff8000', 'A': '#b866f5', 'B': '#3d94f0', 'C': '#1eff00', 'D': '#9d9d9d'}
TIE_PCT = 1.0
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


def tie_mates(out: list[dict], i: int, role: str) -> list[int]:
    """The rows a row ties with (lib/tiers/tier-list.ts): on DPS every other row within the margin of it, never one only
    chained through a third; on healer the neighbours its mark names; tanks none."""
    if role == 'tank' or out[i]['tie'] is None:
        return []
    if role == 'healer':
        return {'above': [i - 1], 'below': [i + 1], 'both': [i - 1, i + 1]}[out[i]['tie']]
    def within(lower: float, higher: float) -> bool:
        return abs(1 - lower / higher) * 100 <= TIE_PCT
    m = out[i]['metric']
    return [j for j, o in enumerate(out) if j != i and within(*((m, o['metric']) if j < i else (o['metric'], m)))]


def join_names(names: list[str]) -> str:
    return names[0] if len(names) == 1 else f"{', '.join(names[:-1])} and {names[-1]}"


def tie_text(r: dict) -> str:
    """'≈ tie with Arms', a partner across a letter line with its letter ('Beast Mastery (B)'); tier-copy.ts tieText."""
    if not r['mates']:
        return ''
    def label(o: dict) -> str:
        letter = f" ({o['tier']})" if r.get('tier') and o.get('tier') and o['tier'] != r['tier'] else ''
        return (o['name'] if r['role'] == 'dps' else full_name(o)) + letter
    return f"&asymp; tie with {join_names([label(o) for o in r['mates']])}"


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
    if not is_ranked or tank:
        pass
    elif role == 'healer':
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
    else:
        within = lambda a, b: abs(1 - a / b) * 100 <= TIE_PCT
        for i, r in enumerate(out):
            a = i > 0 and within(r['metric'], out[i - 1]['metric'])
            b = i + 1 < len(out) and within(out[i + 1]['metric'], r['metric'])
            r['tie'] = 'both' if a and b else ('above' if a else ('below' if b else None))
    if role == 'dps':
        for r in out:
            g = round(r['gap'] * 1e6) / 1e6
            r['tier'] = [t for t, f in TIERS if g >= f][-1]
    for i, r in enumerate(out):
        r['mates'] = [out[j] for j in tie_mates(out, i, role)] if is_ranked else []
    if tank:
        best_eh = max(r['m']['effective_health'] for r in out)
        best_tps = max(r['m']['tps'] for r in out)
        for r in out:
            r['best_eh'] = r['m']['effective_health'] == best_eh
            r['best_tps'] = r['m']['tps'] == best_tps
    return out, is_ranked


def ordinal(n: int) -> str:
    return f"{n}{'th' if 10 <= n % 100 <= 20 else {1: 'st', 2: 'nd', 3: 'rd'}.get(n % 10, 'th')}"


def fnum(v: float, role: str) -> str:
    return f'{v:,.0f}' if role == 'tank' else f'{v:,.1f}'


def full_name(r: dict) -> str:
    return f"{r['name']} {CLASS_NAME[r['cls']]}"


# ------------------------------------------------------------------------------------------------ pieces
LINKS4 = [('BiS', True), ('Guide', False), ('Planner', False), ('Sim', False)]
LINKS3 = LINKS4[:3]


def row_html(r: dict, ranked_: bool, you: bool = False, state: str = '') -> str:
    role = r['role']
    if role == 'tank':
        figs = [(fnum(r['metric'], role), 'taken / s', '', r['rank'] == 1),
                (fnum(r['m']['effective_health'], role), 'HP', 'Effective health', r['best_eh']),
                (fnum(r['m']['tps'], role), 'threat / s', 'Threat', r['best_tps'])]
        gap = 'Least' if r['rank'] == 1 else f"+{r['gap']:.1f}%"
    else:
        figs = [(fnum(r['metric'], role), 'DPS' if role == 'dps' else 'HPS', '', False)]
        if not ranked_:
            gap = f"&minus;{r['gap']:.1f}%" if r['gap'] else '0.0%'
        else:
            gap = 'Top' if r['rank'] == 1 and r['gap'] == 0 else f"&minus;{r['gap']:.1f}%"
    rank = '' if not ranked_ else (f"={r['rank']}" if role == 'healer' and r['tie'] else str(r['rank']))
    tie = tie_text(r)
    return RankRow(rank, crest_img(r['cls'], 36), CLASS_COLOR[r['cls']], r['name'], f"{CLASS_NAME[r['cls']]} &middot; {r['race']}",
                   r['frac'], figs, gap, LINKS4 if role == 'dps' else LINKS3, gap_top=gap in ('Top', 'Least'), tie=tie,
                   low='Stat weights less certain' if r['low'] else '', you='Your spec' if you else '', selected=you, state=state)


def heads(role: str) -> str:
    if role == 'tank':
        return ColumnHeads([('#', True), ('', False), ('Spec', False), ('Share of the least taken', False), ('Taken per second', True),
                            ('Effective health', True), ('Threat per second', True), ('More taken', True), ('', False)])
    bar = "Share of the top spec&rsquo;s damage" if role == 'dps' else "Share of the top healer&rsquo;s healing"
    return ColumnHeads([('#', True), ('', False), ('Spec', False), (bar, False), ('DPS' if role == 'dps' else 'HPS', True), ('Gap', True), ('', False)])


def band_text(letter: str) -> str:
    i = [t for t, _ in TIERS].index(letter)
    start = TIERS[i][1]
    end = TIERS[i + 1][1] if i + 1 < len(TIERS) else None
    if start == 0:
        return f'Within {end}% of the top'
    return f'{start}% or more behind' if end is None else f'{start} to {end}% behind'


def the_list(rows: list[dict], role: str, ranked_: bool, you_spec: str = '', hover_spec: str = '') -> str:
    def one(r):
        return row_html(r, ranked_, r['spec'] == you_spec, 'hover' if r['spec'] == hover_spec else '')
    if role == 'dps':
        body = ''
        for letter, _ in TIERS:
            g = [r for r in rows if r['tier'] == letter]
            if g:
                body += RankGroup(letter, TIER_COLOUR[letter], band_text(letter), f"{len(g)} spec{'s' if len(g) != 1 else ''}", [one(r) for r in g])
        return RankList(heads(role), body, 1, grouped=True)
    return RankList(heads(role), ''.join(one(r) for r in rows), 3 if role == 'tank' else 1, links3=True, tank=role == 'tank')


def caption(role: str, meta: dict, signed_in: bool) -> str:
    text = {
        'dps': f'Damage per second, one target for {DPS_FIGHT_SECONDS} seconds. The bar is the share of the top spec&rsquo;s damage; &asymp;&nbsp;tie is within {TIE_PCT:.0f}%.',
        'healer': f'Effective healing per second over {meta["heal_sec"]} seconds of a stand-in Phase 1 fight. &asymp;&nbsp;tie: the sim cannot order the pair.',
        'tank': f'Damage taken per second against a stand-in level {TANK_BOSS_LEVEL} boss, lower is better. The best of each column is green.',
    }[role]
    sign = '' if signed_in else ' <span><a class="link" href="#">Sign in</a> to see your spec marked.</span>'
    return Notice(f'<span>{text}</span>{sign}', tone='neutral')


def notes(role: str, faction: str, meta: dict, ranked_: bool) -> str:
    fl = FACTION_LABEL[faction]
    gear = (f'Every spec is simmed in its best {meta["preset"]} gear, with enchants a fresh level 60 can buy, raid buffs and consumables, '
            f'as the guide&rsquo;s recommended {fl} race, named under it.')
    fresh = 'At a fresh 60 the order differs: <a class="link" href="#">see the leveling bands</a>.'
    if role == 'dps':
        cols = [('What the number is', f'Damage per second from our own sim against one target for {DPS_FIGHT_SECONDS} seconds. {gear}'),
                ('What it leaves out', f'Cleave, add fights, movement and what a spec brings the raid are not counted, so a spec can sit lower here than it plays in your raid. '
                 f'For how specs do in real raid logs, see <a class="link" href="#">the log rankings</a>. {fresh}'),
                ('Letters and ties', f'S is within 5% of the top, A within 10%, B within 20%, C within 30%, D beyond. &asymp;&nbsp;tie names the specs within {TIE_PCT:.0f}% of a spec; a tie across a line names the other letter, and the two are a tie.')]
    elif role == 'healer':
        order = ('&asymp;&nbsp;tie means the sim cannot order the pair: their order changes with the heal profile&rsquo;s tuning, or their gap is within the measured error. '
                 'Tied healers share a rank: =2, =2, then 4.') if ranked_ else 'Ranks appear once the next ranking publishes the measured error.'
        cols = [('What the number is', f'Effective healing per second over {meta["heal_sec"]} seconds against a stand-in Phase 1 fight: tank hits and raid-wide pulses, not a named boss. Overhealing is not counted. {gear}'),
                ('What it leaves out', f'Mana and raid utility are not ranked here. No level-60 log is behind this order yet: it rests on the sim&rsquo;s heal model until launch logs exist. {fresh}'),
                ('Ties and ranks', f'Gap to the top is less healing in this fight, not a verdict on the class. {order}')]
    else:
        cols = [('What the number is', 'Damage taken per second, lower is better: it is the sim&rsquo;s outcome against the boss, where effective health is only the hit-point pool going in. The best of each column is green.'),
                ('The boss', f'A stand-in level {TANK_BOSS_LEVEL} target with one melee swing every {SWING:g} seconds, not a named boss.'),
                ('The gear', f'{gear} {fresh}')]
    body = '<div class="notes3">' + ''.join(f'<div><h3 class="t-label">{h}</h3><p class="nt">{t}</p></div>' for h, t in cols) + '</div>'
    return Section('How to read this', body, Link('How we check the sim &rsaquo;'))


def bar(role: str, faction: str, counts: dict, open_m: bool = False) -> str:
    roles = seg([(ROLE_LABEL[r], r == role, counts[r], '') for r in ('dps', 'tank', 'healer')])
    em = lambda f: f'<img class="crest" src="{emblem(f)}" alt="" style="width:20px;height:20px;--c:{"#6fb1ff" if f == "alliance" else "#ff6b5c"}">'
    facs = seg([(FACTION_LABEL[f], f == faction, None, em(f)) for f in ('alliance', 'horde')])
    roles_m = seg([(ROLE_LABEL[r], r == role, counts[r], '') for r in ('dps', 'tank', 'healer')])
    return control_bar([('Role', roles, roles_m), ('Faction', facs, facs)], f'{ROLE_LABEL[role]} &middot; {FACTION_LABEL[faction]}', open_m=open_m)


def answer(rows: list[dict], role: str, all_lists: dict, state: str = 'ready') -> str:
    """The signed-in answer in the PageHeader's answer slot. state: ready | loading | unknown."""
    if state == 'loading':
        return (f'<span class="t-label">Your spec</span><span class="ans-line">{Skeleton(32, 32, True)}{Skeleton(160, 20)}</span>'
                f'{Skeleton(120, 40)}{Skeleton(220, 16)}')
    if state == 'unknown':
        return ('<span class="t-label">Your spec</span><span class="t-small">Thandrel&rsquo;s spec is not known yet.</span>'
                + Link('Choose a character &rsaquo;'))
    me_role = next(rl for rl, lst in all_lists.items() if any(r['spec'] == YOU_SPEC for r in lst[0]))
    lst, _ = all_lists[me_role]
    me = next(r for r in lst if r['spec'] == YOU_SPEC)
    if me_role != role:
        # spec 4.6: line 2 is the crest and name over the Small line, the ready state's 40 px crest and two lines
        return (f'<span class="t-label">Your spec</span>'
                f'<span class="ans-row">{crest_img(me["cls"], 40)}<span class="ans-who"><b class="ans-name" style="color:{CLASS_COLOR[me["cls"]]}">{full_name(me)}</b>'
                f'<span class="t-small">{full_name(me)} is on the {ROLE_LABEL[me_role]} list.</span></span></span>'
                + Link('See where it stands &rsaquo;'))
    top = lst[0]
    tier = f'Tier {me["tier"]} &middot; ' if me.get('tier') else ''
    detail = 'Level with the top spec.' if me['rank'] == 1 else f'{tier}{me["gap"]:.1f}% behind {full_name(top)}'
    fig = f'<div class="fig big gold"><span class="v">{ordinal(me["rank"])}<small>of {len(lst)}</small></span></div>'
    return (f'<span class="t-label">Your spec &middot; {ROLE_LABEL[role]}</span>'
            f'<span class="ans-row">{crest_img(me["cls"], 40)}<span class="ans-who"><b class="ans-name" style="color:{CLASS_COLOR[me["cls"]]}">{full_name(me)}</b>'
            f'<span class="t-small">{detail}</span></span>{fig}</span>'
            f'<span class="ans-links">{Link("Your best in slot &rsaquo;")}{Link("Guide &rsaquo;")}{Link("Sim your character &rsaquo;")}</span>')


PAGE_CSS = '''<style>
/* page-level layout only (no surfaces): the answer column, the notes columns, the caption measure */
.head .ans{min-width:0}.ans-line,.ans-row{display:flex;align-items:center;gap:10px}.ans-row{gap:14px}.ans-who{display:flex;flex-direction:column;min-width:0}.ans-name{font:700 18px/22px var(--font-body)}.ans-row .fig{margin-left:8px}.ans-row .fig .v{white-space:nowrap}.ans-who{flex:1}@media(min-width:900px){.ans-who{flex:none}}
.ans-links{display:flex;gap:16px}
@media(min-width:900px){.head .ans{min-width:420px;min-height:96px;justify-content:flex-end}.ans-line,.ans-row{justify-content:flex-end}.ans-who{text-align:right}}
.notes3{display:grid;gap:12px var(--gap);grid-template-columns:minmax(0,1fr);padding-top:12px}
@media(min-width:1024px){.notes3{grid-template-columns:repeat(3,minmax(0,1fr))}}
.notes3 .nt{margin-top:6px;font-size:15px;line-height:24px;color:var(--text);max-width:62ch}
.tiers-plate{padding-top:12px}
.statesheet{display:grid;gap:32px}.statesheet h2{margin-bottom:10px}
</style>'''


def tiers_page(role: str, faction: str = 'alliance', signed_in: bool = False, open_m: bool = False, answer_state: str = 'ready',
               hover_spec: str = '') -> str:
    data = {f: load(f) for f in ('alliance', 'horde')}
    rows, meta = data[faction]
    lists = {r: ranked(rows, r) for r in ('dps', 'tank', 'healer')}
    counts = {r: len(lists[r][0]) for r in lists}
    lst, ranked_ = lists[role]
    stamp = meta['updated'].strftime('%-d %b')
    header = PageHeader('Tier list', meta=[f'Level 60 &middot; {meta["preset"]} &middot; Updated {stamp}'],
                        answer=f'<div class="ans-box">{answer(rows, role, lists, answer_state)}</div>' if signed_in else '')
    plate = f'<div class="tiers-plate">{Plate(the_list(lst, role, ranked_, YOU_SPEC if signed_in else "", hover_spec), list_label(role, faction))}</div>'
    body = (header + bar(role, faction, counts, open_m) + caption(role, meta, signed_in)
            + Page('ledger', plate + f'<div class="section">{notes(role, faction, meta, ranked_)}</div>') + PAGE_CSS)
    note = (f'Tier list mock (design/specs/2026-11-06-tiers.md), built only from ws_prims. Data: build {BUILD}, band 60, preset raid, '
            f'{faction}; ranked by the shipped tier rules. Signed-in character "Thandrel" is a sample.')
    return build_page(f'{ROLE_LABEL[role]} tier list', body, 'Tools', YOU if signed_in else None, note)


def states_page() -> str:
    """Every state the build must show, on one sheet (rendered at 1440 and 390)."""
    rows, meta = load('alliance')
    lists = {r: ranked(rows, r) for r in ('dps', 'tank', 'healer')}
    dps, _ = lists['dps']
    pick = lambda spec: next(r for r in dps if r['spec'] == spec)
    tied = next(r for r in dps if r['tie'])
    low = next((r for r in dps if r['low']), None)
    sample = [('Default', row_html(dps[1], True)),
              ('Hover (S2 ground, the rail slice in the class colour, the bar 18% brighter)', row_html(dps[1], True, state='hover')),
              ('Focus-visible on the BiS link (2 px gold ring, 2 px offset)', row_html(dps[1], True, state='link:fv')),
              ('Your spec (selected: gold ground, gold rail slice, "Your spec" tag)', row_html(pick(YOU_SPEC), True, you=True)),
              (f'A tie (within {TIE_PCT:.0f}% of the one above)', row_html(tied, True))]
    if low:
        sample.append(('Stat weights less certain', row_html(low, True)))
    heal_rows = [dict(r, rank=i + 1) for i, r in enumerate(lists['healer'][0])]
    unranked = row_html(heal_rows[1], False)
    blocks = ''.join(f'<section><h2 class="t-h3">{t}</h2>{Plate(RankList(heads("dps"), h, 1))}</section>' for t, h in sample)
    blocks += f'<section><h2 class="t-h3">Healer list published without the measured error: no ranks, no ties</h2>{Plate(RankList(heads("healer"), unranked, 1, links3=True))}</section>'
    blocks += f'<section><h2 class="t-h3">No data for this build</h2>{Plate("<p class=\'empty\'>No tier list is published for this build yet.</p>")}</section>'
    blocks += f'<section><h2 class="t-h3">A role with no simmed spec</h2>{Plate("<p class=\'empty\'>No Tank specs are simmed yet.</p>")}</section>'
    for t, st, role in (('Answer: loading (reserved at its final size)', 'loading', 'dps'), ('Answer: spec not known', 'unknown', 'dps'),
                        ('Answer: your spec is on another role&rsquo;s list (tank page, Fury Warrior)', 'ready', 'tank')):
        blocks += f'<section><h2 class="t-h3">{t}</h2>{PageHeader("Tier list", meta=["Level 60"], answer=answer(rows, role, lists, st))}</section>'
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
