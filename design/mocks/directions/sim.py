"""Mock batch 2: the simulator in the directions system (plan 2026-11-09-redesign, "Every player-facing surface").

Data, never invented: web/src/fixtures/sim/result-full.json, a real forever-sim run of the published warrior-fury level 60
raid set (10,000 runs), read through design/mocks/ws_sim.py, the parity mock's own reader (names, icons, items, ranks,
casts, auras, the histogram, rage sources, final stats). Typed in and marked as the parity mock marks them: the signed-in
name Thandrel and the Your sims rows. Every element is drawn by prims.py and paperdoll.py.
"""
import re
import sys
from html import escape
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))  # design/mocks: the parity mocks' data readers

import ws_sim as S  # noqa: E402
import data as D  # noqa: E402
import paperdoll as PD  # noqa: E402
from common import write  # noqa: E402
from prims import (PRIMS_CSS, ability_icon, chips, facts, field, histogram, lanes, meter, panel, row_list, settings_bar, talent_trees,  # noqa: E402
                   table)
from style import PAGE_QUALITY, footer, header, medallion, page, titlebar  # noqa: E402

VIEWER = ('Thandrel', 'warrior', 60)
QUALITY_NAME = {0: 'poor', 1: 'common', 2: 'uncommon', 3: 'rare', 4: 'epic', 5: 'legendary'}
SRC = re.compile(r'src="([^"]+)"')
LANE_ABILITIES = 6  # the timeline's lanes: the most-cast abilities, as the live lanes order them
LANE_KEY = '<div class="key"><span><i></i>Cast</span><span><i class="b"></i>Buff up</span></div>'

SIM_CSS = PRIMS_CSS + PD.CSS + '''
.sim{padding-top:2px}.dollwrap{display:flex;justify-content:center}.dollwrap>.pdl{width:100%}.gearwrap{display:grid;grid-template-columns:minmax(0,560px) minmax(0,1fr);gap:24px;align-items:start}
.pd .ctr .fx{width:100%;max-width:300px;background:rgba(6,9,20,.82);border-radius:4px;padding:10px 12px;font-size:13px}
.two{display:grid;grid-template-columns:1fr 1fr;gap:20px}.two h3{font:600 14px var(--body);color:#fff;margin-bottom:4px}
.rk{color:var(--muted);font-size:12.5px;font-weight:400;margin-left:6px}
.cls{display:flex;flex-wrap:wrap;gap:4px 14px;font-size:14px}
@media(max-width:1100px){.gearwrap{grid-template-columns:1fr}}
@media(max-width:700px){.two{grid-template-columns:1fr}}
.simgear .pdl{display:block}.up{font-size:12.5px;color:var(--muted);margin:2px 0 6px}.ups{display:flex;flex-wrap:wrap;gap:6px;margin-bottom:8px}
.bis{display:grid;grid-template-columns:repeat(3,minmax(0,1fr));gap:0 24px}@media(max-width:1100px){.bis{grid-template-columns:minmax(0,1fr)}}.ways td.first{width:190px;color:#fff;font-weight:600}
@media(max-width:700px){.ways tr{display:grid}.ways td{border:0;padding:4px 14px}.ways tr+tr td.first{padding-top:14px}}
.key{display:flex;gap:16px;font-size:13px;color:var(--muted);padding-bottom:8px}.key span{display:flex;align-items:center;gap:6px}
.key i{display:inline-block;width:3px;height:12px;background:#f3eee2}.key i.b{width:12px;height:6px;background:#7fd48a99;border-radius:2px}
'''


def icon_of(html: str) -> str:
    m = SRC.search(html)
    return ability_icon(m.group(1)) if m else ability_icon('')


def action_icon(action: str) -> str:
    """A melee swing shows its weapon, as the ability table does; everything else its spell or item icon."""
    if action.startswith('other:attack/'):
        wpn = S.gear_item({1: 'main_hand', 2: 'off_hand', 3: 'main_hand'}[int(action.rsplit('/', 1)[1])])
        return icon_of(S.gicon(wpn['icon'], 24, wpn['quality']))
    return icon_of(S.act_icon(action))


def fnum(v: float) -> str:
    return f'{v:,.1f}'


# ------------------------------------------------------------------------------------------------ the result
def band_right() -> str:
    return f'<b class="num">{fnum(S.MEAN)}</b> DPS<br><span class="muted" style="font-size:13px">&plusmn; {S.BAND:.1f} (95%)</span>'


def run_settings(action: str) -> str:
    return settings_bar([('Fight', S.FIGHT_STYLE), ('Length', S.FIGHT_LEN), ('Targets', str(S.FIGHT['targets'])), ('Buffs', 'Raid'),
                         ('Precision', S.PRECISION_LABEL)], action)


def ability_table() -> str:
    rows = sorted((a for a in S.SUM['damage_done'][0]['abilities'] if a['total'] > 0), key=lambda a: -a['total'])
    total = sum(a['total'] for a in rows)
    top = rows[0]['total']
    out = []
    for a in rows:
        kind, ident, tag = S.parse(a['name'])
        m = a.get('misses') or {}
        attempts = a['hits'] + m.get('MISS', 0) + m.get('DODGE', 0) + m.get('PARRY', 0)
        if kind == 'other':
            wpn = S.gear_item({1: 'main_hand', 2: 'off_hand', 3: 'main_hand'}[tag])
            label = {1: 'Main hand', 2: 'Off hand', 3: 'Extra attacks'}[tag]
            icon = icon_of(S.gicon(wpn['icon'], 24, wpn['quality']))
            name = (f'<span class="nm" style="color:#fff">{label}</span><span class="sb" style="color:{PAGE_QUALITY[QUALITY_NAME[wpn["quality"]]]}">'
                    f'{escape(wpn["name"])}</span>')
        else:
            icon = icon_of(S.act_icon(a['name']))
            rank = S.rank_of(ident)
            name = f'<span class="nm" style="color:#fff">{escape(S.display_name(a["name"]))}{f"<span class=rk>Rank {rank}</span>" if rank else ""}</span>'
        periodic = a['ticks'] > 0
        casts = '&ndash;' if periodic else str(attempts if kind == 'other' else (S.CASTS.get(a['name']) or attempts))
        hits = a['ticks'] if periodic else a['hits']
        avg = a['total'] / hits if hits else 0
        crit = a['crits'] / hits * 100 if hits else 0
        out.append([f'<span class="pc">{icon}<span>{name}</span></span>', meter(a['total'] / top), fnum(a['total'] / S.DUR),
                    f'{a["total"] / total * 100:.1f}%', casts, f'{avg:,.0f}', f'{crit:.1f}%'])
    out.append(['<b class="strong">Total</b>', '', fnum(total / S.DUR), '100%', '', '', ''])
    heads = [('Ability', False), ('', False), ('DPS', True), ('Share', True), ('Casts', True), ('Avg hit', True), ('Crit', True)]
    return table(heads, out, sort='DPS', row_cls=[''] * (len(out) - 1) + ['tot'], hide_phone={1, 4, 5})


def timeline() -> str:
    """The median run's casts, one lane per ability (the most-cast first), its self buffs as bars under them."""
    dur = S.RES['sample_duration_ms']
    casts: dict[str, list[int]] = {}
    first: dict[str, str] = {}
    for e in S.RES['sample']:
        if e['action'].startswith('spell:') and e['at_ms'] >= 0:
            name = S.display_name(e['action'])  # one lane per ability, whatever rank or id it was cast at
            casts.setdefault(name, []).append(e['at_ms'])
            first.setdefault(name, e['action'])
    rows = []
    for name, times in sorted(casts.items(), key=lambda kv: -len(kv[1]))[:LANE_ABILITIES]:
        body = ''.join(f'<rect x="{t / dur * 1000 - 1:.1f}" y="6" width="2.4" height="18" fill="#f3eee2"/>' for t in times)
        rows.append((f'{action_icon(first[name])}<span>{escape(name)}</span>', body))
    for aura in [a for a in S.RES['sample_auras'] if a['on'] == 'self'][:2]:
        body = ''.join(f'<rect x="{a / dur * 1000:.1f}" y="10" width="{max(1, (b - a) / dur * 1000):.1f}" height="10" rx="2" fill="#7fd48a99"/>' for a, b in aura['segments'])
        rows.append((f'{icon_of(S.act_icon(aura["action"]))}<span>{escape(S.display_name(aura["action"]))}</span>', body))
    return lanes(rows, dur / 1000)


def gear_cells() -> dict[str, dict]:
    cells = {}
    for g in S.CHAR['gear']:
        it = S.ITEMS.get(g['item_id'])
        if it:
            cells[g['slot']] = {'slot': g['slot'], 'state': 'item', 'icon': it['icon'], 'quality': QUALITY_NAME[it['quality']], 'name': it['name']}
    return cells


def gear_panel() -> str:
    fs = S.RES['final_stats']
    stats = facts([('Attack power', f'{fs["attack_power"]:,}'), ('Hit', f'{fs["melee_hit_pct"]:.1f}%'), ('Crit', f'{fs["melee_crit_pct_vs_target"]:.1f}%'),
                   ('Haste', f'{fs["melee_haste_pct"]:.0f}%'), ('Weapon skill', str(fs['weapon_skill']['main_hand'])),
                   ('Target armor after debuffs', f'{S.RES["target_stats"]["armor_after_debuffs"]:,}')])
    # One job per page (owner): the BiS page owns the gear frame; the sim lists what it ran, with the sheet it ran at
    listed = PD.list_form(gear_cells(), '', False, S.SPEC_LABEL, '', f'<div style="margin-top:14px">{stats}</div>', head=False)
    return panel('Gear', f'<div class="simgear" style="--frame:url({PD.frame_uri()})">{listed}</div>',
                 action='<a href="#">Upgrades on Best in slot &rsaquo;</a>')


def talents_panel() -> str:
    trees = []
    for tr, part in zip(sorted(S.TALENTS, key=lambda t: t['position']), S.CHAR['talents'].split('-')):
        tal = sorted(tr['talents'], key=lambda t: (t['tier'], t['column']))
        cells = [(t['tier'], t['column'], SRC.search(S.gicon(t['icon'], 40)).group(1), int(part[i]) if i < len(part) else 0, t['max_rank'])
                 for i, t in enumerate(tal)]
        trees.append((tr['name'], sum(int(c) for c in part), cells))
    main = max(trees, key=lambda t: t[1])  # the spec's tree; the planner owns the whole build
    split = facts([(name, str(points)) for name, points, _ in trees])
    body = f'<div style="display:flex;gap:28px;align-items:flex-start;flex-wrap:wrap">{talent_trees([main])}<div style="min-width:220px">{split}</div></div>'
    return panel('Talents', body, action='<a href="#">Open in planner &rsaquo;</a>')


def buffs_panel() -> str:
    """Each aura's uptime over the fight; on you, then on the target."""
    dur = S.SUM['duration_ms']

    def col(kind: str) -> str:
        # an aura up the whole fight is one icon in a row (its bar says nothing); only the ones that come and go get a bar
        auras = sorted((a for a in S.SUM['auras'] if a['type'] == kind), key=lambda a: -a['uptime_ms'])
        full = [a for a in auras if a['uptime_ms'] >= dur * .995]
        part = [a for a in auras if a['uptime_ms'] < dur * .995][:6]
        icons = ''.join(f'<span title="{escape(S.display_name(a["name"]))}">{action_icon(a["name"])}</span>' for a in full)
        whole = f'<p class="up">Up the whole fight</p><div class="ups">{icons}</div>' if full else ''
        bars = row_list([(action_icon(a['name']), escape(S.display_name(a['name'])), meter(a['uptime_ms'] / dur, 'var(--gain)'),
                          f'{a["uptime_ms"] / dur * 100:.0f}%') for a in part])
        return whole + bars

    return panel('Buffs and debuffs', f'<div class="two"><div><h3>On you</h3>{col("BUFF")}</div><div><h3>On the target</h3>{col("DEBUFF")}</div></div>')


def rage_panel() -> str:
    src = next(r for r in S.RES['resource_sources'] if r['resource'] == 'rage')

    def rows(key: str) -> str:
        items = sorted(src[key], key=lambda x: -x['amount'])[:6]
        return row_list([(action_icon(x['action']), escape(S.display_name(x['action'])), '', f'{x["amount"]:,.0f} rage') for x in items])

    return panel('Rage', f'<div class="two"><div><h3>Gained from</h3>{rows("gained")}</div><div><h3>Spent on</h3>{rows("spent")}</div></div>', )


def spread_panel() -> str:
    h = S.RES['dps_histogram']
    counts, start, bw = h['counts'], h['start'], h['bin_dps']
    total, acc, lo, hi = sum(counts), 0, None, None
    for i, c in enumerate(counts):
        acc += c
        lo = lo if lo is not None or acc < total * .16 else start + i * bw
        hi = hi if hi is not None or acc < total * .84 else start + (i + 1) * bw
    d = S.RES['dps']
    figs = facts([('Mean', f'{fnum(S.MEAN)} DPS'), ('Error', f'&plusmn; {S.BAND:.1f}'), ('Std dev', f'{d["stddev"]:.1f}'),
                  ('Lowest', f'{d["min"]:,.0f}'), ('Highest', f'{d["max"]:,.0f}'), ('Runs', f'{S.RUNS:,}')])
    hist = histogram(counts, start, bw, S.MEAN, lo, hi, f'DPS of each of {S.RUNS:,} runs')
    return panel('DPS spread', f'{hist}<p class="muted" style="font-size:13px;margin:8px 0 12px">Two runs in three: {lo:,.0f} to {hi:,.0f}</p>{figs}')


def history_panel() -> str:
    rows = [(medallion('warrior', 28), f'{escape(b)} &middot; {escape(spec)}', escape(day), f'{dps} DPS') for b, spec, day, dps in S.HISTORY]
    return panel('Your sims', row_list(rows), action='<a href="#">All &rsaquo;</a>')


def result_page() -> str:
    meta = f'{S.RACE} {S.SPEC_LABEL} &middot; 60 &middot; {S.SPLIT} talents &middot; average item level {S.ilvl():.0f}'
    band = titlebar('warrior', VIEWER[0], meta, right=band_right(), title_colour=D.CLASS_COLOR['warrior'],
                    ctl=[[(t, t == 'Quick sim') for t in S.TOOLS]], top=header(*VIEWER, 'Sim'))
    main = (panel('Damage by ability', ability_table(), flush=True, framed=True)
            + panel('Timeline', LANE_KEY + timeline()) + gear_panel() + talents_panel() + rage_panel() + buffs_panel())
    rail = spread_panel() + history_panel()
    body = (band + f'<main class="w sim">{run_settings("<a class=btn href=#>Run again</a>")}'
            f'<div class="twocol"><div>{main}</div><aside>{rail}</aside></div></main>' + footer())
    return page('Forever Sixty, sim', SIM_CSS, body)


# ------------------------------------------------------------------------------------------------ the landing
def landing_page() -> str:
    paste = (f'{field("FS1:…", "<a class=btn href=#>Run</a>", "Addon export or build code")}'
             f'<p class="muted" style="font-size:13.5px;margin-top:10px">In game: <b class="strong">/fs copy</b>, then paste here.</p>')
    ways = '<div class="ways">' + table([('', False), ('', False)], [
        ['Your characters', '<a class="btn2" href="#">Sign in with Battle.net</a>'],
        ['A planner build', field('foreversixty.gg/b/…', '<a class=btn2 href=#>Load</a>')],
        ['A logged fight', field('foreversixty.gg/reports/…', '<a class=btn2 href=#>Load</a>')]]).replace('<thead><tr><th class="first"></th><th class="last"></th></tr></thead>', '') + '</div>'
    rows = []
    for c in S.CLASSES:
        specs = [s for s in S.SPECS if s['class_slug'] == c['slug'] and s['role'] == 'dps']
        if specs:
            links = chips([(s['name'], False) for s in specs])
            rows.append((medallion(c['slug'], 30), f'<span style="color:{D.CLASS_COLOR[c["slug"]]}">{escape(c["name"])}</span>', links, ''))
    main = (panel('Sim your character', paste, framed=True) + panel('Other ways in', ways, flush=True)
            + panel('Sim a best-in-slot set', '<div class="bis">' + ''.join(row_list([r]) for r in rows) + '</div>'))
    band = titlebar(None, 'Simulator', '', ctl=[[(t, t == 'Quick sim') for t in S.TOOLS]], top=header(*VIEWER, 'Sim'))
    body = band + f'<main class="w sim">{main}</main>' + footer()
    return page('Forever Sixty, simulator', SIM_CSS, body)


def build() -> None:
    write('b2-sim', landing_page())
    write('b2-sim-result', result_page())


if __name__ == '__main__':
    build()
