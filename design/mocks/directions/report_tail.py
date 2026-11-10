"""The report's remaining views in the directions system (batch H lane h-report-tail): Events, Dispels, Interrupts, the
whole night (fight=all) and /ratings. Queries and Timeline keep their batch 1 mocks (report.py, b1-report-queries and
b1-report-timeline).

  python3 design/mocks/directions/report_tail.py   -> html/b1-report-{events,dispels,interrupts,night}.html, html/b1-ratings.html

Data, never invented: every pull view reads the bundled sample (web/public/sample-report/sample, the 8 Nov 2026 night in
Excavation Site: Wetlands, names changed) at the pull the live page opens on (Saltspine #1, pull 11); Events reads the
pull's events.parquet through duckdb as the page's stream does; the night folds the sample's six boss pulls as
lib/report/night.ts does. /ratings reads lib/score/ratings-page.ts and lib/score/copy.ts's sample figures (typed here).
Every element is drawn by prims.py.
"""
import sys
from html import escape
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

import report as RB  # noqa: E402  the batch 1 report mock: band, scope, this_pull, icons
import ws_report as RP  # noqa: E402
import ws_report_lib as L  # noqa: E402
from common import write  # noqa: E402
from prims import chips, facts, icon_cell, meter, panel, player_cell, row_list, table  # noqa: E402
from style import footer, header, page, titlebar  # noqa: E402

R, PULL, who, spell_icon = RB.R, RB.PULL, RB.who, RB.spell_icon
EVENT_ROWS = 40   # a page of lines; the rest behind Show more
SHOW = ['Casts', 'Damage', 'Heals', 'Auras', 'Deaths']
TAIL_CSS = RB.REPORT_CSS + '''
.evh{display:flex;flex-wrap:wrap;align-items:center;gap:8px 16px;padding:0 20px 10px}.evh .chips{padding:0}
.evh .fld{margin-left:auto;width:240px}.evh .fld input{padding:6px 10px;font-size:14px}
.ev .pc .nm{font-weight:400;color:var(--text)!important}.ev b{font-weight:600}
.more{padding:12px 20px;display:flex;align-items:center;gap:14px;font-size:13px;color:var(--muted)}
.foot1{padding:10px 20px 14px;font-size:13px;color:var(--muted);border-top:1px solid rgba(255,255,255,.07)}
.foot1 a{color:var(--gold)}
.parts .tb td.first{width:44%}
@media(max-width:700px){.evh{padding:0 14px 10px}.evh .fld{margin-left:0;width:100%}.more,.foot1{padding-left:14px;padding-right:14px}}
'''


def names_html(sentence: str) -> str:
    return sentence.replace('<span style', '<b style').replace('</span>', '</b>')


def tab_events() -> tuple[str, str]:
    src = 'Your group'
    s, _, _, _ = RB.scope(src)
    scope = L.in_scope(R, s, src)
    by_name = {r['name']: r for r in s['roster']}
    evs = [e for e in RP.events(R, PULL, {r['name'] for r in scope}) if e[1] in ('cast_success', 'aura_applied', 'aura_removed', 'death')]
    rows = [[L.at(ms), f'<span class="pc">{spell_icon(sid, sp, 20) if sid else ""}<span class="ev">{names_html(RP.ev_sentence(by_name, k, a, b, sid, sp, amt))}</span></span>',
             L.n0(amt) if amt else '&ndash;'] for ms, k, a, b, sid, sp, amt in evs[:EVENT_ROWS]]
    head = (f'<div class="evh">{chips([(t, True) for t in SHOW])}'
            '<div class="fld"><input placeholder="Find a name, spell or amount"></div></div>')
    more = len(evs) - EVENT_ROWS
    foot = (f'<div class="more"><a class="btn2" href="#">Show 100 more</a><span>{more:,} more</span></div>'
            '<p class="foot1">Hits and heals: only those before a death. <a href="#">Load every hit and heal</a></p>')
    body = head + table([('Time', False), ('Event', False), ('Amount', True)], rows) + foot
    return panel('Events', body, meta=f'{len(evs):,} in this window', flush=True), RB.this_pull(src)


def kick_rows(key: str) -> list[list[str]]:
    s = R.summary(PULL)
    ros = L.roster(s)
    verb = 'Interrupted' if key == 'interrupts' else 'Removed'
    def cell(x: dict) -> str:
        name = (f'<b style="color:{L.colour(ros.get(x["source_guid"]))}">{escape(L.split_name(x["source_name"])[0])}</b>'
                f' &middot; {escape(x["spell_name"])}')
        sub = f'{verb} {escape(x["extra_spell_name"])} on {escape(L.split_name(x["target_name"])[0])}'
        return (f'<span class="pc">{spell_icon(x["spell_id"], x["spell_name"])}<span><span class="nm">{name}</span>'
                f'<span class="sb">{sub}</span></span></span>')
    return [[cell(x), str(x['count'])] for x in s[key]]


def enemy_casts() -> str:
    """The rail on Interrupts: what the enemies cast this pull, by spell, most first (summary.casts, not players)."""
    s = R.summary(PULL)
    players = {r['guid'] for r in s['roster']}
    by: dict[int, list] = {}
    for c in s['casts']:
        if c['guid'] not in players and not c.get('owner_guid'):
            e = by.setdefault(c['spell_id'], [c['spell_name'], set(), 0])
            e[1].add(L.split_name(c['name'])[0])
            e[2] += c['succeeded']
    top = sorted(by.items(), key=lambda kv: -kv[1][2])[:8]
    rows = [(spell_icon(sid, n), escape(n), escape(', '.join(sorted(w))), f'<b class="num">{k}</b>') for sid, (n, w, k) in top]
    return panel('Enemy casts', row_list(rows), meta='this pull')


def group_debuffs() -> str:
    """The rail on Dispels: the debuffs on the group this pull, by spell, with how often each landed (summary.auras)."""
    s = R.summary(PULL)
    players = {r['guid'] for r in s['roster']}
    by: dict[str, list] = {}
    for a in s['auras']:
        if a['type'] == 'DEBUFF' and a['target_guid'] in players:
            e = by.setdefault(a['name'], [a['spell_id'], set(), 0])
            e[1].add(a['target_guid'])
            e[2] += len(a['segments'])
    top = sorted(by.items(), key=lambda kv: -kv[1][2])[:8]
    rows = [(spell_icon(sid, n), escape(n), f'on {len(g)} {"player" if len(g) == 1 else "players"}', f'<b class="num">{k}&times;</b>')
            for n, (sid, g, k) in top]
    return panel('Debuffs on your group', row_list(rows), action='<a href="#">Debuffs &rsaquo;</a>')


def tab_dispels() -> tuple[str, str]:
    rows = kick_rows('dispels')
    body = table([('Who, with what', False), ('Count', True)], rows)
    return (panel('Dispels', body, meta=f'{sum(int(r[1]) for r in rows)} this pull', action='<a href="#">Copy CSV</a>', flush=True),
            RB.this_pull('Your group') + group_debuffs())


def tab_interrupts() -> tuple[str, str]:
    rows = kick_rows('interrupts')
    body = table([('Who, with what', False), ('Count', True)], rows) if rows else '<p class="empty">Nothing interrupted this pull.</p>'
    return panel('Interrupts', body, meta=f'{sum(int(r[1]) for r in rows)} this pull'), RB.this_pull('Your group') + enemy_casts()


# ------------------------------------------------------------------------------------------------ the whole night
def night() -> tuple[str, str, str]:
    bosses = [f for f in R.meta['fights'] if f['kind'] == 'encounter']
    players: dict[str, dict] = {}
    by_boss: dict[str, dict] = {}
    for f in bosses:
        s = R.summary(f['index'])
        group = set(f.get('group') or f['players'])
        for r in s['roster']:
            if r['guid'] not in group:
                continue
            p = players.setdefault(r['guid'], {'r': r, 'fights': 0, 'dmg': 0, 'heal': 0, 'taken': 0, 'active': 0, 'ms': 0, 'deaths': 0})
            p['fights'] += 1
            p['dmg'] += r['damage_done']
            p['heal'] += r['healing_done']
            p['taken'] += r['damage_taken']
            p['active'] += r['active_ms']
            p['ms'] += s['duration_ms']
            p['deaths'] += r['deaths']
        b = by_boss.setdefault(f['name'], {'pulls': 0, 'kills': 0, 'ms': 0, 'deaths': 0, 'best': None})
        b['pulls'] += 1
        b['kills'] += 1 if f['kill'] else 0
        b['ms'] += f['duration_ms']
        b['deaths'] += f.get('group_deaths', 0)
        if f['kill'] and (b['best'] is None or f['duration_ms'] < b['best']):
            b['best'] = f['duration_ms']
    ps = sorted(players.values(), key=lambda p: -p['dmg'])
    prow = [[player_cell(*who(p['r']), p['r']['class']), str(p['fights']), L.n0(p['dmg']),
             L.rate(p['dmg'] / p['ms'] * 1000), L.n0(p['heal']), L.n0(p['taken']), f'{p["active"] / p["ms"] * 100:.0f}%', str(p['deaths'])] for p in ps]
    heads = [('Player', False), ('Pulls', True), ('Damage', True), ('DPS', True), ('Healing', True), ('Taken', True), ('Active', True), ('Deaths', True)]
    brow = [[icon_cell('', n, f'Best kill {L.mmss(b["best"])}' if b['best'] else 'No kill'),
             str(b['pulls']), str(b['kills']), str(b['pulls'] - b['kills']), L.mmss(b['ms']), str(b['deaths'])] for n, b in by_boss.items()]
    kills = sum(b['kills'] for b in by_boss.values())
    time_ms = sum(f['duration_ms'] for f in bosses)
    deaths = sum(b['deaths'] for b in by_boss.values())
    main = (panel('Players', table(heads, prow, sort='Damage', hide_phone={4, 5, 6}), meta='across every boss pull', flush=True, framed=True)
            + panel('Bosses', table([('Boss', False), ('Pulls', True), ('Kills', True), ('Wipes', True), ('Time', True), ('Deaths', True)], brow,
                                    hide_phone={3, 4}), flush=True))
    rail = panel('The night', facts([('Boss pulls', str(len(bosses))), ('Kills', str(kills)), ('Wipes', str(len(bosses) - kills)),
                                     ('In combat', L.mmss(time_ms)), ('Deaths', str(deaths))]))
    meta = f'{escape(R.zone)} &middot; All boss pulls &middot; {len(bosses)} pulls &middot; {kills} kills &middot; {L.mmss(time_ms)}'
    return main, rail, meta


def night_page() -> str:
    main, rail, meta = night()
    strip = [(t, t == 'Summary') for t in RB.TABS] + [('More &#9662;', False)]
    ctl = [strip, [('All boss pulls &#9662;', False), ('Your group &#9662;', False)]]
    band = '<div class="rb">' + titlebar(None, R.title, meta, ctl=ctl, top=header(*RB.VIEWER, 'Logs')).replace('&amp;#9662;', '&#9662;') + '</div>'
    body = band + f'<main class="w rp"><div class="twocol"><div>{main}</div><aside>{rail}</aside></div></main>' + footer()
    return page('Forever Sixty, report, the whole night', TAIL_CSS, body)


# ------------------------------------------------------------------------------------------------ /ratings
SCORE = 1830
PARTS = [('Gear', 412, 500, 'Your gear against your spec’s best set for your level'), ('Experience', 620, 700, '25 a dungeon cleared, 40 a raid boss killed'),
         ('Parses', 540, 600, 'Your best boss parses against your spec'), ('Attendance', 180, 200, 'Your guild’s raid nights in the last 90 days'),
         ('Preparation', 78, 100, 'Flask, food and enchants on every logged pull')]
TIERS = [('Common', 'below 500', '#9d9d9d'), ('Uncommon', '500 to 999', '#1eff00'), ('Rare', '1000 to 1499', '#0070dd'), ('Epic', '1500 to 1999', '#a335ee'),
         ('Legendary', '2000 and up', '#ff8000')]
WEIGHTS = [('Output', 35, 30, 10, 'Damage or effective healing against what the fight allowed'), ('Survival', 15, 10, 35, 'Avoidable deaths and damage'),
           ('Mechanics', 20, 20, 20, 'Interrupts, dispels and debuffs carried'), ('Utility', 15, 20, 20, 'Uptime of your class’s buffs and debuffs'),
           ('Preparation', 10, 10, 10, 'Flask, food, oil and world buffs at the pull'), ('Activity', 5, 10, 5, 'Time spent casting or swinging')]


def ratings_page() -> str:
    parts = table([('Part', False), ('', False), ('Points', True), ('Top', True)],
                  [[f'<span class="pc"><span><span class="nm" style="color:#fff">{p}</span><span class="sb">{d}</span></span></span>', meter(v / t), str(v), str(t)]
                   for p, v, t, d in PARTS], hide_phone={1}, bar_col=1)
    weights = table([('Component', False), ('DPS', True), ('Healer', True), ('Tank', True)],
                    [[f'<span class="pc"><span><span class="nm" style="color:#fff">{c}</span><span class="sb">{d}</span></span></span>', str(a), str(b), str(t)]
                     for c, a, b, t, d in WEIGHTS])
    rating = ('<div class="rt">' + facts([('Latest fight', '<b style="color:var(--gold);font-size:22px">65</b> of 100'), ('Last 6 fights', 'Up 25'),
                                          ('Weakest part', 'Activity')]) + '</div>'
              '<p class="foot1" style="padding:12px 0 0">Public reports only. An early avoidable death caps a fight at 40.</p>')
    rules = row_list([('', 'Ranked', '20 or more kills of the encounter, spec and kill speed: against your spec', ''),
                      ('', 'Absolute', 'Fewer: against the encounter’s own numbers', ''),
                      ('', 'Wipes', 'Survival and Mechanics only', '')])
    main = (panel('Forever Score', parts, meta='Sample', flush=True, ornate=True) + panel('Performance Rating', rating, framed=True)
            + panel('Weights, by role', weights, meta='Guild officers can change these', flush=True) + panel('How a fight is judged', rules))
    tiers = row_list([('', f'<span style="color:{c}">{n}</span>', r, 'Sample' if n == 'Epic' else '') for n, r, c in TIERS])
    rail = panel('Tiers', tiers, framed=True) + panel('Get your logs here', '<p class="empty">The companion uploads your combat log as you play.</p>'
                                                      '<a class="btn" href="#">Get the companion</a>')
    right = '<b class="num" style="color:#a335ee">1,830</b> of 2,100 &middot; Epic'
    band = titlebar(None, 'Forever Score and ratings', 'Most to gain: Gear, 88 below its top', right=right,
                    ctl=[[('Forever Score', True), ('Performance Rating', False)]], top=header(*RB.VIEWER, ''))
    body = band + f'<main class="w rp"><div class="twocol"><div>{main}</div><aside>{rail}</aside></div></main>' + footer()
    return page('Forever Sixty, ratings', TAIL_CSS, body)


def report_page(tab: str, fn) -> str:
    main, rail = fn()
    body = RB.report_band(tab, 'Your group', '') + f'<main class="w rp"><div class="twocol"><div>{main}</div><aside>{rail}</aside></div></main>' + footer()
    return page(f'Forever Sixty, report, {tab}', TAIL_CSS, body)


def build() -> None:
    write('b1-report-events', report_page('Events', tab_events))
    write('b1-report-dispels', report_page('Dispels', tab_dispels))
    write('b1-report-interrupts', report_page('Interrupts', tab_interrupts))
    write('b1-report-night', night_page())
    write('b1-ratings', ratings_page())


if __name__ == '__main__':
    build()
