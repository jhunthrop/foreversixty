"""The report in the directions system (plan 2026-11-09-redesign, batch 1, "logs and reports").

  python3 design/mocks/directions/report.py   -> html/b1-report-<view>.html

Data, never invented:
- Every pull view reads the real 17 Sept 2026 beta night on Zephras Isle (web/public/sample-report/sample, the bundled
  sample, anonymised names; trash only) through design/mocks/ws_report_lib.py, the parity mocks' own reader. Queries runs
  its SQL on the pull's events.parquet with duckdb, as the page runs it in the browser.
- Mechanics needs a boss pull, which no real log has yet (D1): it reads Onyxia's mechanics table
  (logs/engine/mechanics/tables/1084.json) and the Onyxia wipe in the guild report fixture (design/mocks/guild_report_data.py).
- Coach reads web/src/fixtures/coach/fury-model.json; its copy fixes are proposals for api/internal/coach (COACH_COPY).
- The player view's gear: the sample has none (the game writes gear at boss pulls), so the frame is filled from the
  Beast Mastery 20 to 29 list (design/mocks/data/bis-facts.json) to draw the state with gear.
Every element is drawn by prims.py; nothing page-local beyond layout.
"""
import base64
import json
import sys
from html import escape
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))  # design/mocks: the parity mocks' data readers

import duckdb  # noqa: E402
import guild_data as GD  # noqa: E402
import guild_report_data as GR  # noqa: E402
import ws_report as RP  # noqa: E402  melee() and events(): the parity mock's own readers
import ws_report_lib as L  # noqa: E402
import data as D  # noqa: E402
import paperdoll as PD  # noqa: E402
from common import doll_cells, write  # noqa: E402
from pages import two_handed  # noqa: E402
from prims import (PRIMS_CSS, ability_icon, chart_frame, chips, code_block, column_chart, facts, icon_cell, item_row, lanes, line_chart,  # noqa: E402
                   meter, panel, player_cell, responsive_chart, row_list, shades, strip_into_view, table, tag, tick_step)
from style import footer, header, page, titlebar, tooltip  # noqa: E402

R = L.Report()
PULL = L.DEFAULT_PULL
VIEWER = ('Thandrel', 'warrior', 60)
# lib/report/frame.ts REPORT_TABS, in order; below 1760 px the strip shows TABS_MD and the rest sit in More
TABS = ['Summary', 'Damage done', 'Healing', 'Damage taken', 'Deaths', 'Mechanics', 'Timeline']
COACH = json.loads((L.ROOT / 'web/src/fixtures/coach/fury-model.json').read_text())
MECH_TABLE = json.loads((L.ROOT / 'logs/engine/mechanics/tables/1084.json').read_text())
CHART_SERIES = 5  # the chart draws the top five, as the live chart does
TABLE_ROWS = 15  # a long table (265 friendlies took damage) shows its top 15 and "All 265"
FEW_SERIES = 3  # healing and damage taken: the top three by player (reviewers, batch 1); 19 or 265 lines is noise
FOCUS_PLAYER = 'Soren Everhart'  # the player view's player: the pull's top damage
UNIT = {'damage_done': 'DPS', 'healing': 'HPS', 'damage_taken': 'DTPS'}

REPORT_CSS = PRIMS_CSS + PD.CSS + '''
.rp{padding-top:6px}
.heal{color:var(--gain)}
.dh{display:flex;flex-wrap:wrap;align-items:center;gap:4px 12px;padding:14px 20px 8px;border-top:1px solid rgba(255,255,255,.1)}
.dh:first-child{border-top:0}.dh .k{margin-left:auto;font-size:13px;color:var(--muted)}.dfoot{padding:6px 20px 14px;font-size:13px;color:var(--muted)}
.tb tr.kb td{background:rgba(240,115,107,.1)}
.tl .ln{padding:6px 20px 14px}.tl .ln .lb .pc>span{min-width:0}.tl .ln .lb .nm{font-size:13.5px;overflow:hidden;text-overflow:ellipsis}
.pn .chips{padding:0 20px 10px}
.cz{max-width:62ch}.cz h3{font:600 15px var(--body);color:#fff;margin:16px 0 4px}.cz p{font-size:15px;line-height:1.6}.cz .lead{font-size:16px;color:#fff}
.cz .fix{border-left:3px solid var(--gold);padding:2px 0 2px 14px;margin-top:14px}.cz .fix h3{margin-top:6px}
.cz ul.ok li{padding:5px 0 5px 22px;position:relative;font-size:14.5px}
.cz ul.ok li::before{content:"";position:absolute;left:4px;top:13px;width:8px;height:8px;border-radius:50%;background:var(--gain)}
.cz details{margin-top:14px;border-top:1px solid rgba(255,255,255,.1);padding-top:10px;font-size:14px;color:var(--muted)}
.cz summary{cursor:pointer;color:#fff;font-weight:600}
.mech{display:grid;grid-template-columns:24px minmax(0,1fr) auto;gap:4px 12px;padding:12px 20px;border-top:1px solid rgba(255,255,255,.08)}
.mech:first-child{border-top:0}.mech .nm{font-weight:600;color:#fff}.mech .nt{grid-column:2/4;font-size:14px;color:var(--text);max-width:75ch}
.mech .who{grid-column:2/4;display:flex;flex-wrap:wrap;gap:6px 18px;font-size:14px}.mech .ct{font:600 14px var(--body);color:#fff;white-space:nowrap}
.mech.ok .ct{color:var(--gain)}
.q{display:grid;gap:12px}.q .run{display:flex;align-items:center;gap:14px;font-size:13px;color:var(--muted)}
.cols{font:13px/1.9 ui-monospace,Menlo,monospace;color:var(--text)}
.sm{display:grid;grid-template-columns:230px minmax(0,1fr);gap:0 20px;align-items:center;padding:8px 20px;border-top:1px solid rgba(255,255,255,.07)}
.sm:first-of-type{border-top:0}.sm>div:first-child{min-width:0;overflow:hidden}
.dt .tb{table-layout:fixed}.dt th:nth-child(1){width:84px}.dt th:nth-child(3){width:168px}.dt th:nth-child(4){width:90px}
.dt th:nth-child(5){width:110px}
.pv{display:grid;grid-template-columns:auto minmax(0,1fr);gap:24px;align-items:start;padding:4px 20px 18px}
.pv .pd{--pd-max:520px;width:520px}.pv .gl .ir{padding:4px 0}.pv .gl .ir img{width:28px;height:28px}.pv .gl .ir{grid-template-columns:28px 1fr auto}
.pv .acts{display:flex;gap:16px;align-items:center;margin-top:14px;flex-wrap:wrap}
@media(max-width:1100px){.pv{grid-template-columns:1fr}}
@media(max-width:700px){.tl .ln{padding:6px 14px 12px}.tl .ln .lb .medal{display:none}
.dh{padding:12px 14px 6px}.dh .k{margin-left:0;flex-basis:100%}
.mech{padding:12px 14px}.sm{grid-template-columns:1fr;padding:8px 14px}.pv{padding:4px 14px 14px}.pv .pd,.pv .gl{display:none}.pv .pdl{display:block}
.dfoot{padding:6px 14px 12px}.pn .chips{padding:0 14px 10px}.dt th:nth-child(1){width:64px}.dt th:nth-child(4){width:56px}
.dt th:nth-child(5){width:76px}}
'''


# ------------------------------------------------------------------------------------------------ helpers
def who(r: dict | None, name: str = '') -> tuple[str | None, str]:
    """A roster row's class slug and its display name (the realm is dropped; the sample is anonymised)."""
    return (L.slug(r) if r else None), L.split_name((r or {}).get('name') or name)[0]


# What the client itself draws: the auto-attack icon for melee. A spell with no icon in the build (Forever's new spells,
# such as Shock 1259652) takes the icon of another spell of the same name, else its school's base spell; never the
# question mark (wow-player and ux-designer, batch 1).
MELEE_ICON = 'ability_meleedamage'
SCHOOL_ICON = {1: 'ability_meleedamage', 2: 'spell_holy_holybolt', 4: 'spell_fire_fire', 8: 'spell_nature_lightning',
               16: 'spell_frost_frostbolt02', 32: 'spell_shadow_shadowbolt', 64: 'spell_nature_starfall'}
_BY_NAME: dict[str, str] = {}
for _sid, _name in L.SPELL_NAME.items():
    if L.SPELL_ICON.get(str(_sid)):
        _BY_NAME.setdefault(_name, L.SPELL_ICON[str(_sid)])
_SCHOOL = {ab['spell_id']: ab.get('school') for f in (PULL,) for a in R.summary(f)['damage_taken'] for ab in a['abilities']}
# The client's icons where the build's spell table has none for the spell (wow-player, batch 1 round 2): Auto Shot is the
# bow-and-arrow ability_whirlwind in the Classic client. Onyxia's fire breaths are fire school.
CLIENT_ICON = {75: 'ability_whirlwind'}
_SCHOOL.update({18435: 4, 17086: 4})  # Flame Breath, Deep Breath


def icon_name(spell_id: int, name: str = '', fallback: str = '') -> str:
    """fallback: the icon when neither the id nor the name has one and the school is unknown (a boss ability: the boss's own)."""
    if spell_id == 0:
        return MELEE_ICON
    name = name or L.SPELL_NAME.get(spell_id, '')
    school = SCHOOL_ICON.get(_SCHOOL[spell_id], '') if _SCHOOL.get(spell_id) else ''
    found = CLIENT_ICON.get(spell_id) or L.SPELL_ICON.get(str(spell_id)) or _BY_NAME.get(name) or school or fallback or MELEE_ICON
    return found if (L.B / 'icons' / f'{found}.webp').exists() else MELEE_ICON


def spell_icon(spell_id: int, name: str = '', size: int = 24, fallback: str = '') -> str:
    """The spell's icon as the shared ability icon (prims.ability_icon)."""
    uri = 'data:image/webp;base64,' + base64.b64encode((L.B / 'icons' / f'{icon_name(spell_id, name, fallback)}.webp').read_bytes()).decode()
    return ability_icon(uri, size)


def scope(source: str) -> tuple[dict, set[str], dict, float]:
    s = R.summary(PULL)
    return s, L.friendlies(R, s, source), L.roster(s), s['duration_ms'] / 1000


def group_rate(key: str, source: str) -> float:
    s, fr, _, dur = scope(source)
    return sum(a['effective'] for a in s[key] if a['guid'] in fr) / dur


def report_band(tab: str, source: str, right: str, title: str = '', meta: str = '', pull: str = '') -> str:
    """The report's title band. The strip is the live order (frame.ts); a tab past the strip shows as "More: <tab>"."""
    f = R.fight(PULL)
    meta = meta or (f'{escape(R.zone)} &middot; Pull {PULL}, {escape(f["name"])} &middot; {L.mmss(f["duration_ms"])} &middot; '
                    f'{f.get("mob_deaths", 0)} mobs died')
    more = 'More' if tab in TABS else f'More: {tab}'
    strip = [(t, t == tab) for t in TABS] + [(f'{more} &#9662;', tab not in TABS)]
    ctl = [strip, [(f'{pull or f"Pull {PULL}"} &#9662;', False), (f'{source} &#9662;', False)]]
    band = titlebar(None, title or R.title, meta, right=right, ctl=ctl, top=header(*VIEWER, 'Logs'))
    return '<div class="rb">' + band.replace('&amp;#9662;', '&#9662;') + '</div>'


def this_pull(source: str) -> str:
    """Only what the band does not say (reviewers: Pull, Length and Mobs died repeated it): the start and the scope's counts."""
    f = R.fight(PULL)
    s, fr, _, _ = scope(source)
    items = [('Started', L.clock(f['start'])), (f'In {source.lower()}', f'{len(L.in_scope(R, s, source)):,}'),
             ('Deaths', str(sum(1 for d in s['deaths'] if d['guid'] in fr)))]
    return panel('This pull', facts(items))


def series_for(acts: list[dict], ros: dict, dur: float, unit: str) -> tuple[list, list]:
    """Each actor's line in its class colour, repeats of a class stepped lighter or darker (prims.shades)."""
    named = [(who(ros.get(a['guid']), a['name'])[1], L.colour(ros.get(a['guid'])), a) for a in acts]
    cols = shades([c for _, c, _ in named])
    series = [(n, c, a['series']) for (n, _, a), c in zip(named, cols)]
    legend = [(n, c, f'{L.rate(a["effective"] / dur)} {unit}') for (n, _, a), c in zip(named, cols)]
    return series, legend


def metric_chart(source: str, key: str, title: str, count: int = CHART_SERIES) -> str:
    s, fr, ros, dur = scope(source)
    acts = sorted([a for a in s[key] if a['guid'] in fr and a['total'] > 0], key=lambda a: -a['effective'])[:count]
    if not acts:
        return panel(title, '<p class="empty">Nothing in this window.</p>')
    series, legend = series_for(acts, ros, dur, UNIT[key])
    marks = [d['at_ms'] / 1000 for d in s['deaths'] if d['guid'] in fr]
    return panel(title, chart_frame(responsive_chart(series, dur, marks, f'{title}, per player', unit=UNIT[key]), legend))


def figure(key: str, source: str) -> str:
    return f'<b class="num">{L.rate(group_rate(key, source))}</b> group {UNIT[key]}'


def link_chevron() -> str:
    return '<a href="#" aria-label="Abilities">&rsaquo;</a>'


# ------------------------------------------------------------------------------------------------ summary and actor tabs
def players_table(source: str) -> str:
    s, _, _, _ = scope(source)
    ros = sorted(L.in_scope(R, s, source), key=lambda r: (-r['damage_done'], -r['healing_done']))
    top = max((r['damage_done'] for r in ros), default=1) or 1
    rows = [[player_cell(*who(r), r['class']), meter(r['damage_done'] / top, L.colour(r)), L.n0(r['damage_done']), L.rate(r['dps']),
             L.n0(r['healing_done']), L.n0(r['damage_taken']), f'{r["activity_pct"]:.0f}%', str(r['deaths'])] for r in ros]
    heads = [('Player', False), ('', False), ('Damage', True), ('DPS', True), ('Healing', True), ('Taken', True), ('Active', True), ('Deaths', True)]
    return table(heads, rows, sort='Damage', hide_phone={1, 4, 5, 6}, bar_col=1)


def taken_by_ability(source: str) -> str:
    s, fr, _, dur = scope(source)
    abil: dict[str, list] = {}
    for a in s['damage_taken']:
        if a['guid'] in fr:
            for ab in a['abilities']:
                abil.setdefault(ab['name'] or 'Melee', [ab.get('spell_id', 0), 0])[1] += ab.get('effective', ab['total'])
    tot = sum(v for _, v in abil.values()) or 1
    rows = [[icon_cell(spell_icon(sid, k), RP.melee(sid, k)), meter(v / tot, 'var(--rule-warm)'), f'{v / tot * 100:.1f}%', L.rate(v / dur), L.n0(v)]
            for k, (sid, v) in sorted(abil.items(), key=lambda kv: -kv[1][1])]
    return table([('Ability', False), ('', False), ('Share', True), ('Per sec', True), ('Amount', True)], rows, hide_phone={1, 3}, bar_col=1)


def tab_summary() -> tuple[str, str]:
    src = 'Your group'
    main = (panel('Players', players_table(src), flush=True, framed=True)
            + metric_chart(src, 'damage_done', 'Damage over the fight')
            + panel('Damage taken by ability', taken_by_ability(src), flush=True))
    return main, this_pull(src)


def actor_rows(key: str, acts: list[dict], ros: dict, dur: float, total_ms: int) -> list[list[str]]:
    top = acts[0]['effective'] if acts else 1
    rows = []
    for a in acts:
        r = ros.get(a['guid'])
        slug, name = who(r, a['name'])
        over = f'{(a["total"] - a["effective"]) / a["total"] * 100:.0f}%' if a['total'] else '0%'
        sub = (r or {}).get('class', '') + (f'<span class="om"> &middot; {over} overheal</span>' if key == 'healing' else '')
        cells = [player_cell(slug, name, sub), meter(a['effective'] / top, L.colour(r)), L.n0(a['effective']), L.rate(a['effective'] / dur)]
        if key == 'healing':
            cells.append(over)
        cells += [f'{a["effective"] / sum(x["effective"] for x in acts) * 100:.1f}%', f'{a["active_ms"] / total_ms * 100:.0f}%', link_chevron()]
        rows.append(cells)
    return rows


def actor_tab(key: str, title: str, source: str, side_title: str) -> tuple[str, str]:
    s, fr, ros, dur = scope(source)
    acts = sorted([a for a in s[key] if a['guid'] in fr and a['effective'] > 0], key=lambda a: -a['effective'])
    tot = sum(a['effective'] for a in acts) or 1
    rows = actor_rows(key, acts, ros, dur, s['duration_ms'])[:TABLE_ROWS]
    over = [('Overheal', True)] if key == 'healing' else []
    heads = [('Player', False), ('', False), ('Amount', True), (UNIT[key], True), *over, ('Share', True), ('Active', True), ('', False)]
    raw = sum(a['total'] for a in acts) or 1
    total = ['<b class="strong">Total</b>', '', L.n0(tot), L.rate(tot / dur), *([f'{(raw - tot) / raw * 100:.0f}%'] if over else []), '100%', '', '']
    hide = {1, 4, 5} if not over else {1, 4, 5, 6}
    tbl = table(heads, rows + [total], sort='Amount', row_cls=[''] * len(rows) + ['tot'], hide_phone=hide, bar_col=1)
    more = f'<a href="#">All {len(acts)} &rsaquo;</a>' if len(acts) > TABLE_ROWS else ''
    side: dict[str, float] = {}
    for a in acts:
        for t in a.get('targets', []):
            side[t['name']] = side.get(t['name'], 0) + t['total']
    st = sum(side.values()) or 1
    targets = row_list([('', escape(L.split_name(k)[0]), f'{v / st * 100:.1f}%', L.n0(v)) for k, v in sorted(side.items(), key=lambda kv: -kv[1])[:8]])
    count = CHART_SERIES if key == 'damage_done' else FEW_SERIES
    main = panel(title, tbl, action=more, flush=True, framed=True) + metric_chart(source, key, f'{title} over the fight', count)
    if key == 'damage_taken':
        main += per_minute(source) + panel('Damage taken by ability', taken_by_ability(source), flush=True)
    return main, this_pull(source) + panel(side_title, targets)


def per_minute(source: str) -> str:
    """The group's damage taken in each minute: the bars form, so the heavy minutes stand out at a glance."""
    s, fr, _, dur = scope(source)
    mins = int(dur // 60) + 1
    per = [0.0] * mins
    for a in s['damage_taken']:
        if a['guid'] in fr:
            for i, v in enumerate(a['series']):
                per[min(int(i * dur / max(len(a['series']) - 1, 1) // 60), mins - 1)] += v
    bars = [(f'{m}:00', v, 'var(--rule-warm)') for m, v in enumerate(per)]
    body = (f'<div class="od">{column_chart(bars, "Damage", label="Damage taken each minute")}</div>'
            f'<div class="om">{column_chart(bars, "Damage", 230, "Damage taken each minute", 380, every=3)}</div>')
    return panel('Damage taken each minute', f'<div class="cf">{body}</div>')


# ------------------------------------------------------------------------------------------------ deaths
def death_rows(d: dict) -> tuple[list[list[str]], list[str], bool]:
    kb = d['killing_blow']
    events = sorted([('hit', h) for h in d['last']] + [('heal', h) for h in d.get('heals', [])], key=lambda e: e[1]['at_ms'])
    rows, cls = [], []
    for kind, h in events:
        mx = h.get('max_hp') or kb['max_hp']
        killing = kind == 'hit' and h['at_ms'] == kb['at_ms']
        hp = 'dead' if killing else f'{L.n0(h["hp_after"])} / {L.n0(mx)}'
        amt = f'<span class="heal">+{L.n0(h["amount"])}</span>' if kind == 'heal' else f'&minus;{L.n0(h["amount"])}'
        src_name = L.split_name(h['source_name'])[0] if kind == 'heal' else h['source_name']
        name = RP.melee(h['spell_id'], h['spell_name'])
        sub = f'from {escape(src_name)}' + (f' &middot; {L.n0(h.get("overkill", 0))} overkill' if killing else '')
        cell = icon_cell(spell_icon(h['spell_id'], h['spell_name']), name, sub) + (f' {tag("Killing blow", "wipe")}' if killing else '')
        rows.append([f'<span class="num muted">{L.at(h["at_ms"])}</span>', cell,
                     meter(0 if killing else h.get('hp_after', 0) / mx, 'var(--gain)' if kind == 'heal' else '#c0392b'), amt, hp])
        cls.append('kb' if killing else '')
    return rows, cls, any(k == 'heal' for k, _ in events)


def tab_deaths() -> tuple[str, str]:
    src = 'All friendlies'
    s, fr, ros, _ = scope(src)
    blocks = []
    for d in [d for d in s['deaths'] if d['guid'] in fr]:
        r = ros.get(d['guid'])
        slug, name = who(r, d['name'])
        kb = d['killing_blow']
        rows, cls, healed = death_rows(d)
        title = 'Last hits and heals' if healed else 'Last hits'
        tbl = '<div class="dt">' + table([('Time', False), (title, False), ('', False), ('Amount', True), ('Health', True)], rows, row_cls=cls,
                                         hide_phone={2}) + '</div>'
        held = ', '.join(escape(x['name']) for x in d['auras_held'])
        foot = ' '.join(x for x in (f'Holding {held}.' if held else '', '' if healed else 'No heal landed in the last 10 s.') if x)
        blocks.append(f'<div class="dh">{player_cell(slug, name, r["class"] if r else "", 32)}<span class="k"><span class="num">{L.at(d["at_ms"])}</span>'
                      f' &middot; {escape(kb["source_name"])}</span></div>{tbl}' + (f'<p class="dfoot">{foot}</p>' if foot else ''))
    # the damage-taken chart was cut: on this pull it is flat for eleven minutes (ai-slop-reviewer); the recap is the page
    main = panel('Deaths', ''.join(blocks) or '<p class="empty">Nobody died.</p>', flush=True, framed=True)
    return main, this_pull(src)


def death_count(source: str) -> int:
    s, fr, _, _ = scope(source)
    return sum(1 for d in s['deaths'] if d['guid'] in fr)


# ------------------------------------------------------------------------------------------------ mechanics (Onyxia, the fixture wipe)
MECH_PULL = next(p for p in GR.PULLS if p[2] == 1)  # Onyxia #1, the wipe: the pull with the most to fix
CLASS_OF = {row[0]: row[1] for row in GD.R}
BOSS_ICON = dict(D.RAIDS)["Onyxia's Lair"]  # the client's Onyxia icon, for her abilities the build has no icon for


def mech_block(m: dict, deaths: list[tuple[str, str]]) -> str:
    """One mechanic: icon, name, "Tank takes it" where the table says so, its deaths with who died. Every row of this table is
    avoidable, so the kind is not tagged (ai-slop-reviewer: chip only the exceptions)."""
    killed = [n for n, ability in deaths if ability == m['name']]
    people = ''.join(player_cell(CLASS_OF.get(n), n, size=22) for n in killed)
    count = f'{len(killed)} {"death" if len(killed) == 1 else "deaths"}' if killed else '<span class="muted">No deaths</span>'
    note = f'<p class="nt">{escape(m["note"])}</p>' if m.get('note') else ''
    return (f'<div class="mech">{spell_icon(m["spell_id"], m["name"], fallback=BOSS_ICON)}<span><span class="nm">{escape(m["name"])}</span>'
            f'{" " + tag("Tank takes it") if m.get("role") == "tank" else ""}</span><span class="ct">{count}</span>{note}'
            + (f'<div class="who">{people}</div>' if killed else '') + '</div>')


def tab_mechanics() -> tuple[str, str]:
    """The table's mechanics, then any killing ability the table lacks (Deep Breath is not in 1084.json yet), so the rows add
    up to the band's deaths (reviewers, round 2)."""
    _, _, _, _, ms, deaths = MECH_PULL
    known = {m['name'] for m in MECH_TABLE['mechanics']}
    extra = [{'spell_id': GR.SPELL.get(a, 0), 'name': a} for a in dict.fromkeys(a for _, a in deaths) if a not in known]
    rows = sorted(MECH_TABLE['mechanics'] + extra, key=lambda m: -sum(1 for _, a in deaths if a == m['name']))
    main = panel('Mechanics', ''.join(mech_block(m, deaths) for m in rows), flush=True, framed=True)
    return main, ''  # who died sits under each mechanic; a rail list repeated it (ai-slop-reviewer)


# ------------------------------------------------------------------------------------------------ timeline
LANE = {'cast': 'var(--cream)', 'buff': '#7fd48a99', 'debuff': '#f0736b99', 'death': '#f0736b'}
SHOW = ['Casts', 'Damage', 'Heals', 'Auras', 'Deaths', 'Power']
# the filter chips carry each mark's colour, so they are the legend too; damage marks take each player's class colour
SWATCH = {'Casts': LANE['cast'], 'Damage': 'linear-gradient(90deg,#aad372 50%,#ff7c0a 50%)', 'Heals': '#32ae88', 'Auras': LANE['buff'],
          'Deaths': LANE['death'], 'Power': '#3fa7e0'}
LANE_W = 1000  # prims.lanes draws each lane on 0..1000 x 0..30


def lane_body(r: dict, s: dict, evs: list, dur: int) -> str:
    """One player's lane on the shared lane track: casts above its rule, hits on it, auras at the foot, a death full height."""
    w, lane = LANE_W, ''
    for a in [a for a in s['auras'] if a['target_guid'] == r['guid']]:
        for sg in a['segments']:
            lane += (f'<rect x="{sg["start_ms"] / dur * w:.1f}" y="24" width="{max(1, (sg["end_ms"] - sg["start_ms"]) / dur * w):.1f}" height="5" rx="2" '
                     f'fill="{LANE["buff"] if a["type"] == "BUFF" else LANE["debuff"]}"/>')
    for ms, k, src_n, dst, _sid, _sp, _amt in evs:
        if k == 'cast_success' and src_n == r['name']:
            lane += f'<rect x="{ms / dur * w - 1:.1f}" y="1" width="2" height="11" fill="{LANE["cast"]}"/>'
        elif k in ('damage', 'damage_landed') and src_n == r['name']:
            lane += f'<rect x="{ms / dur * w - 1:.1f}" y="12" width="2" height="9" fill="{L.colour(r)}"/>'
        elif k == 'death' and dst == r['name']:
            lane += f'<rect x="{ms / dur * w - 1.5:.1f}" y="0" width="3" height="30" fill="{LANE["death"]}"/>'
    return lane


def tab_timeline() -> tuple[str, str]:
    src = 'Your group'
    s, _, _, _ = scope(src)
    players = sorted(L.in_scope(R, s, src), key=lambda r: -r['damage_done'])
    evs = RP.events(R, PULL, {r['name'] for r in players})
    dur = s['duration_ms']
    rows = [(player_cell(*who(r), size=22), lane_body(r, s, evs, dur)) for r in players]
    show = chips([(t, t in ('Casts', 'Damage', 'Auras', 'Deaths')) for t in SHOW], SWATCH)
    body = f'{show}<div class="tl">{lanes(rows, dur / 1000, tick_step(dur / 1000, phone=False))}</div>'
    return panel('Timeline', body, flush=True, framed=True), this_pull(src)


# ------------------------------------------------------------------------------------------------ the charts: threat and resources


def tab_threat() -> tuple[str, str]:
    src = 'Your group'
    s, fr, ros, dur = scope(src)
    by_enemy: dict[str, float] = {}
    for t in s['threat_by_target']:
        if t['guid'] in fr:
            by_enemy[t['target_guid']] = by_enemy.get(t['target_guid'], 0) + t['threat']
    eg = max(by_enemy, key=by_enemy.get)
    tt = sorted([t for t in s['threat_by_target'] if t['guid'] in fr and t['target_guid'] == eg], key=lambda t: -t['threat'])
    enemy = tt[0]['target_name']
    acts = [{'guid': t['guid'], 'name': t['name'], 'effective': t['threat'], 'series': t['series']} for t in tt]
    series, _ = series_for(acts, ros, dur, '')
    tot, top = sum(t['threat'] for t in tt) or 1, tt[0]['threat']
    legend = [(n, c, f'{L.n0(t["threat"])}') for (n, c, _), t in zip(series, tt)]
    chart = chart_frame(responsive_chart(series, dur, label=f'Threat on {enemy}', unit='Threat'), legend)
    rows = [[player_cell(*who(ros.get(t['guid']), t['name']), ros.get(t['guid'], {}).get('class', '')), meter(t['threat'] / top, L.colour(ros.get(t['guid']))),
             L.n0(t['threat']), f'{t["threat"] / tot * 100:.1f}%'] for t in tt]
    tbl = table([('Player', False), ('', False), ('Threat', True), ('Share', True)], rows, sort='Threat', hide_phone={1}, bar_col=1)
    # "Estimated" carries the caveat; the formula sits behind the link (ai-slop-reviewer: no explainer under the chart)
    main = (panel(enemy, chart, meta=tag('Estimated'), action='<a href="#">What it leaves out</a>', framed=True)
            + panel('By player', tbl, flush=True))
    return main, this_pull(src)  # Taunts shows only when someone taunted; nobody did on this pull


def tab_resources() -> tuple[str, str]:
    src = 'Your group'
    s, fr, ros, dur = scope(src)
    rows = [x for x in s['resources'] if x['guid'] in fr and x['guid'] in ros and x['max']]
    powers = {L.POWER.get(x['power_type'], 'Power') for x in rows}
    single = len(powers) == 1  # one power for everyone: the panel names it once; mixed, each axis names its own
    blocks = ''
    for x in sorted(rows, key=lambda x: -x['spent']):
        r = ros[x['guid']]
        sub = f'Spent {L.n0(x["spent"])} &middot; Empty {L.mmss(x["zero_ms"])}'
        unit = '' if single else L.POWER.get(x['power_type'], '')
        chart = f'<div class="od">{_area(x, r, dur, 130, 600, unit)}</div><div class="om">{_area(x, r, dur, 160, 380, unit)}</div>'
        blocks += f'<div class="sm"><div>{player_cell(*who(r), sub)}</div><div class="cf">{chart}</div></div>'
    return panel('Resources', blocks, meta=next(iter(powers)) if single else '', flush=True, framed=True), this_pull(src)


def _area(x: dict, r: dict, dur: float, height: int, width: int, unit: str) -> str:
    """One player's power bar over the pull."""
    return line_chart([('', L.colour(r), x['series'])], dur, height, smooth=1, width=width, unit=unit,
                      area=True, label=f'{L.POWER.get(x["power_type"], "Power")} over the pull')


# ------------------------------------------------------------------------------------------------ queries
# Proposed starting point: amount less overkill, so the totals match Damage done's (wow-player: 467 here, 450 there)
QUERY = ("SELECT source_name, spell_name, count(*) AS hits,\n       sum(amount - greatest(overkill, 0)) AS total\nFROM read_parquet('events.parquet')\n"
         "WHERE kind = 'damage'\nGROUP BY 1, 2\nORDER BY total DESC")
STARTS = ['Damage by ability', 'Pet damage', 'Every miss', 'Healing and overhealing', 'Healing by target']
COLUMNS = ('time_unix_nano kind event source_name dest_name spell_id spell_name amount overkill overheal absorbed critical miss_type '
           'aura_type power_type')
SHOWN_ROWS = 12


def tab_queries() -> tuple[str, str]:
    src = 'Your group'
    s, _, ros, _ = scope(src)
    p = R.base / f'fights/{PULL}/events.parquet'
    res = duckdb.sql(QUERY.replace("read_parquet('events.parquet')", f"read_parquet('{p}')")).fetchall()
    by_name = {r['name']: r for r in s['roster']}
    rows = [[f'<b style="color:{L.colour(by_name.get(a))}">{escape(L.split_name(a)[0])}</b>', escape(RP.melee(0, b or '')), str(c), L.n0(d)]
            for a, b, c, d in res[:SHOWN_ROWS]]
    run = f'<div class="run"><a class="btn" href="#">Run</a><span>{len(res)} rows &middot; {SHOWN_ROWS} shown</span></div>'
    body = f'<div class="q">{chips([(t, i == 0) for i, t in enumerate(STARTS)])}{code_block(QUERY)}{run}</div>'
    out = table([('source_name', False), ('spell_name', False), ('hits', True), ('total', True)], rows)
    main = panel('Queries', body, framed=True) + panel('Result', out, flush=True)
    cols = panel('Columns', f'<p class="cols">{COLUMNS.replace(" ", " &middot; ")}</p>', action='<a href="#">All 66</a>')
    return main, this_pull(src) + cols


# ------------------------------------------------------------------------------------------------ one player: abilities and gear
GEAR_KEY = "low-20"  # bis-facts.json: Beast Mastery Hunter, 20 to 29


def player_abilities(a: dict, dur: float) -> str:
    top = max(ab['effective'] for ab in a['abilities'])
    rows = []
    for ab in sorted(a['abilities'], key=lambda x: -x['effective']):
        name = RP.melee(ab['spell_id'], ab['name'])
        rows.append([icon_cell(spell_icon(ab['spell_id'], ab['name']), name), meter(ab['effective'] / top, L.colour({'class': a['class']})),
                     L.n0(ab['effective']), L.rate(ab['effective'] / dur), f'{ab["effective"] / a["effective"] * 100:.1f}%', str(ab['hits']),
                     f'{ab["crits"] / max(ab["hits"], 1) * 100:.0f}%', f'{L.n0(ab["effective"] / ab["hits"]) if ab["hits"] else "0"}'])
    heads = [('Ability', False), ('', False), ('Amount', True), ('DPS', True), ('Share', True), ('Hits', True), ('Crit', True), ('Average', True)]
    return table(heads, rows, sort='Amount', hide_phone={1, 4, 7}, bar_col=1)


def gear_at_pull(name: str, dps: str) -> str:
    """The frame in report mode beside the worn list (names in quality colour, as the client's sheet hover gives them); the
    list form alone on a phone. The selected slot's tooltip stands in the frame's centre, as on BiS."""
    facts_ = D.bis(GEAR_KEY)
    cells = doll_cells(facts_)
    sel = 'ranged'
    two = two_handed(facts_, cells)
    tip = tooltip('hunter', cells[sel]['item'])
    doll = PD.frame(cells, sel, two, name, 'Hunter', dps, tip)
    lst = PD.list_form(cells, sel, two, name, dps, f'<div style="margin-top:12px">{tip}</div>')
    worn = [(PD.SLOT_LABEL[n], cells[n]) for n in PD.LIST_SLOTS if cells.get(n, {}).get('state') == 'item']
    rows = ''.join(item_row(D.icon_uri(c['icon']), c['name'], c['quality'], slot) for slot, c in worn)
    acts = '<p class="acts"><a class="btn" href="#">Sim this gear</a><a href="#">Best in slot &rsaquo;</a></p>'
    return f'<div>{doll}{lst}</div><div class="gl">{rows}{acts}</div>'


def tab_player() -> tuple[str, str]:
    s, fr, ros, dur = scope('Your group')
    a = next(x for x in s['damage_done'] if L.split_name(x['name'])[0] == FOCUS_PLAYER)
    r = ros[a['guid']]
    series = [(FOCUS_PLAYER, L.colour(r), a['series'])]
    chart = chart_frame(responsive_chart(series, dur, label=f'{FOCUS_PLAYER}, damage', unit='DPS', area=True), [])
    side: dict[str, float] = {}
    for t in a.get('targets', []):
        side[t['name']] = side.get(t['name'], 0) + t['total']
    st = sum(side.values()) or 1
    targets = row_list([('', escape(L.split_name(k)[0]), f'{v / st * 100:.1f}%', L.n0(v)) for k, v in sorted(side.items(), key=lambda kv: -kv[1])[:8]])
    main = (panel('Abilities', player_abilities(a, dur), flush=True, framed=True)
            + panel('Damage over the pull', chart)
            + panel('Gear at the pull', f'<div class="pv" style="--frame:url({PD.frame_uri()})">'
                    f'{gear_at_pull(FOCUS_PLAYER, L.rate(a["effective"] / dur))}</div>', flush=True))
    return main, this_pull('Your group') + panel('Targets', targets)


# ------------------------------------------------------------------------------------------------ coach
# Proposed copy for api/internal/coach (reviewers, batch 1): level, not head count; units on every percentage; no
# double colon; the sentence that explains Heroic Strike to a Fury warrior and the unfinished potion check are cut.
COACH_COPY = {
    # the opening leads with the measure against the benchmark (wow-player); "alive the whole fight" stays in Doing well only
    'Solid kill: you were alive the whole fight and Battle Shout was up 100% of it. Almost everything you left behind sits in one button.':
        '6.2 Bloodthirsts a minute against 9.1 for the top 25 Fury Warriors. Almost everything you left behind sits in that one button.',
    'Bloodthirst is the button the Fury rotation is built around, and 9 casts': '9 casts',
    'From forty, on trash pulls,': 'From level 40, on trash pulls,',
    'Bloodthirst: 9 casts lost': '9 Bloodthirst casts lost',
    'against 97% for the top 25.': 'against 97% for the top 25 Fury Warriors.',
}
COACH_CUT_WHY = {'f2'}
COACH_SKIP = ('Combat potions',)


def copy(text: str) -> str:
    for old, new in COACH_COPY.items():
        text = text.replace(old, new)
    return text


def tab_coach() -> tuple[str, str]:
    rv = COACH['review']
    fx = rv['fixFirst']
    then = ''.join(f'<h3>{escape(copy(t["title"]))}</h3>'
                   + ('' if t['finding'] in COACH_CUT_WHY else f'<p>{escape(copy(t["why"]))}</p>')
                   + ''.join(f'<p>{escape(x["sentence"])}</p>' for x in t.get('links', [])) for t in rv['thenFix'])
    skipped = [x for x in rv['couldntCheck'] if not x.startswith(COACH_SKIP)]
    body = (f'<div class="cz"><p class="lead">{escape(copy(rv["opening"]))}</p>'
            f'<div class="fix">{tag("Fix first", "gold")}<h3>{escape(copy(fx["title"]))}</h3><p>{escape(copy(fx["why"]))}</p>'
            f'<p style="margin-top:8px;color:#fff">{escape(copy(fx["drillText"]))}</p></div>'
            f'{then}<h3>Doing well</h3><ul class="ok">{"".join(f"<li>{escape(copy(x["text"]))}</li>" for x in rv["doingWell"])}</ul>'
            f'<details><summary>Not checked ({len(skipped)})</summary><ul>{"".join(f"<li>{escape(x)}</li>" for x in skipped)}</ul></details></div>')
    who_line = player_cell('warrior', 'Baelgrim', 'Fury Warrior', 32)
    main = panel('Coach', f'<div style="padding:4px 0 12px">{who_line}</div>{body}', framed=True, ornate=True)
    # the fixture's basis_labels: the top 25 Fury Warriors on Onyxia (40)
    measured = table([('', False), ('You', True), ('Top 25', True)],
                     [['Bloodthirst casts a minute', '6.2', '9.1'], ['Battle Shout uptime', '100%', '97%']])
    return main, panel('Against the top 25', measured, flush=True)


# ------------------------------------------------------------------------------------------------ pages
VIEWS = {
    'summary': ('Summary', 'Your group', tab_summary, 'damage_done'),
    'damage': ('Damage done', 'Your group', lambda: actor_tab('damage_done', 'Damage done', 'Your group', 'Targets'), 'damage_done'),
    'healing': ('Healing', 'All friendlies', lambda: actor_tab('healing', 'Healing', 'All friendlies', 'Healed'), 'healing'),
    'taken': ('Damage taken', 'All friendlies', lambda: actor_tab('damage_taken', 'Damage taken', 'All friendlies', 'Sources'), 'damage_taken'),
    'deaths': ('Deaths', 'All friendlies', tab_deaths, 'damage_taken'),
    'mechanics': ('Mechanics', 'Your raid', tab_mechanics, ''),
    'timeline': ('Timeline', 'Your group', tab_timeline, ''),
    'threat': ('Threat', 'Your group', tab_threat, ''),
    'resources': ('Resources', 'Your group', tab_resources, ''),
    'queries': ('Queries', 'Your group', tab_queries, ''),
    'player': ('Damage done', FOCUS_PLAYER, tab_player, 'damage_done'),
    'coach': ('Coach', 'Baelgrim', tab_coach, ''),
}


def band_for(view: str, tab: str, source: str, key: str) -> str:
    """The band's one figure follows the tab: group DPS, HPS or DTPS on the rate tabs, deaths on Deaths and Mechanics, the
    player's DPS on their view, none where no single figure answers the tab (timeline, threat, power, queries, coach)."""
    if view == 'coach':  # the coach fixture is its own report: Onyxia, a kill, 40 players
        return report_band(tab, source, '', title="Onyxia's Lair", meta='Sample review &middot; Pull 3, Onyxia &middot; Kill &middot; 40 players',
                           pull='Pull 3, Onyxia')
    if view == 'mechanics':  # the guild report fixture's Onyxia night: 02:00 UTC on 9 Dec, the raids' first night
        _, _, _, _, ms, deaths = MECH_PULL
        return report_band(tab, source, f'<b class="num">{len(deaths)}</b> deaths', title="Onyxia's Lair, 9 Dec 2026",
                           meta=f'Pull 2, Onyxia #1 &middot; Wipe &middot; {L.mmss(ms)} &middot; 40 players', pull='Onyxia #1, Wipe')
    if view == 'player':
        a = next(x for x in R.summary(PULL)['damage_done'] if L.split_name(x['name'])[0] == FOCUS_PLAYER)
        return report_band(tab, source, f'<b class="num">{L.rate(a["effective"] / (R.summary(PULL)["duration_ms"] / 1000))}</b> DPS')
    if view == 'deaths':
        n = death_count(source)
        return report_band(tab, source, f'<b class="num">{n}</b> {"death" if n == 1 else "deaths"}')
    return report_band(tab, source, figure(key, source) if key else '')


def report_page(view: str) -> str:
    tab, source, fn, key = VIEWS[view]
    main, rail = fn()
    body = (band_for(view, tab, source, key) + strip_into_view()
            + f'<main class="w rp"><div class="twocol"><div>{main}</div><aside>{rail}</aside></div></main>' + footer())
    return page(f'Forever Sixty, report, {tab}', REPORT_CSS, body)


def build() -> None:
    for view in VIEWS:
        write(f'b1-report-{view}', report_page(view))


if __name__ == '__main__':
    build()
