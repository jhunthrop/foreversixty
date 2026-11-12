"""The guild tabs batch H redrew in the dense language (design/specs/2026-11-09-guild-dense.md, section 8): Raids, the raid
planner (officer and member), Progression, Applications, Recruit and Settings, drawn from the 40-raider world's API answers
(web/tests/e2e/fixtures/guild-parity/dense-*.json, the files the build reads) with the shared primitives only. They are the
parity references of the `parity` states in design/mocks/web/guild-dense.json (web/tests/e2e/parity-guild.spec.ts).

  python3 design/mocks/ws_guild_dense_tabs.py        -> design/mocks/web/web-guild-dense-{raids,planner,...}.html
"""
from __future__ import annotations

import json
from datetime import datetime, timedelta



from ws_prims import (EmptyState, Select, SignupRow, RowGroup, control_bar, seg, Button, Figures, Link, Page, PageHeader, Row, Rows, Section, Segmented, StatRow, StatTable, Tag, TabStrip,
                      build_page, crest_img, icon, write)
import ws_prims
from guild_data import C
import guild_data as gd
import guild_dense_data as D
from ws_guild2 import ALLIANCE, GUILD_CSS, comp_grid, _FACTS_CSS, _PROSE, _ROLE_CSS, _act, _day, _votes, _view, band_facts, clock_label, guild_time
from guild_fixtures import OUT
from ws_guild_raids import PRIM as RAIDS_PRIM
import ws_guild_settings as WS


def fixture(name: str) -> dict:
    """One of the dense world's API answers, as the build reads it (web/tests/e2e/fixtures/guild-parity)."""
    return json.loads((OUT / name).read_text())

from ws_guild_dense import (at as at_zone, GROUP_WORD, VIEW_MEMBER, VIEW_OFFICER, countdown, day, dense, fig, gap_of, head, legend, name_html,
                            officer_answer, parse_class, ready_counts, strip, tone)
from guild_dense_tabs_fixtures import APPLICATIONS, WANTED

ONY_ICON = 'inv_misc_head_dragon_black'
# ui/StatRow's icon cell is a 28 px block (ui-st-icon), and a dense row's name cell sets 14/18 (stat-table.css data-dense).
ICON_CSS = ('<style>.st-r>span[aria-hidden]>img{display:block}.st.dense .st-r .nm{font-size:14px;line-height:18px}'
            # ui/StatTable keeps its own tracks when dense (stat-table.css), unlike the dense roster mock's wider name floor.
            'div.st.dense{--st-name:minmax(0,2fr);--st-bar:minmax(0,1fr);--st-fig:minmax(max-content,62px)}</style>')
VOTERS = gd.VOTERS


def clock(ms: int) -> str:
    return f'{ms // 60000}:{ms // 1000 % 60:02d}'


def grp(label: str, sub: str, n: int) -> str:
    """A group head row inside a dense table (the roster's role heads)."""
    return StatRow('', f'<b>{label}</b>', sub, [''] * n, bar=False).replace('st-r stc', 'st-r stc grp', 1)


# ---- Raids
def planned_row(raid: dict) -> str:
    """One planned raid as the built PlannedRaidRow draws it from 1024 (ws_guild_raids.planned_row on the dense world's clock)."""
    at = at_zone(raid['starts_at'])
    c = raid['counts']
    signed = c['in'] + c['late']
    left = at - at_zone(D.NOW)
    countdown_words = f'in {gd.countdown_text(left)}'
    word = lambda n, one, many: f'{n} {one if n == 1 else many}'  # noqa: E731
    parts = [word(c['tanks'], 'tank', 'tanks'), word(c['healers'], 'healer', 'healers'), f"{c['melee']} melee", f"{c['ranged']} ranged"]
    late = sum(c['late_groups'].values())
    if late:
        heal = c['late_groups']['healers']
        parts.append(f'{late} late' + (f' ({heal} healer{"" if heal == 1 else "s"})' if heal else ''))
    return (f'<a class="prow" href="#"><span class="dt" aria-hidden="true"><span class="l">{at:%a}</span><span class="d">{at.day}</span>'
            f'<span class="l">{at:%b}</span></span><span><span class="tt">{raid["title"]}</span></span>'
            f'<span><span class="sr-only">{at:%A} {at.day} {at:%B}, </span><span class="tm">{gd.house_clock(at)} CST</span>'
            f'<span class="cd gold">{countdown_words}</span></span>'
            f'<span style="display:flex;flex-direction:column;gap:4px"><span class="ct">{signed}/{raid["size"]}<span class="bar meter"><i style="width:{round(100 * signed / raid["size"])}%"></i></span></span>'
            f'<span class="sm">{" &middot; ".join(parts)}</span></span>'
            f'<span class="chips"><span class="pill">In &middot; DPS</span></span><span class="chev" aria-hidden="true">&rsaquo;</span></a>')


def night_row(f: dict, slowest: int, best: int, before: int | None, loot: int, misses: dict) -> str:
    """One logged night (GuildRaidNights): the date as the row's button, the pulls and the tops under it, the boss time."""
    kill = next(x['duration_ms'] for x in f['fights'] if x['name'] == 'Onyxia' and x['kill'])
    delta = '&ndash;' if before is None else tone(signed_s(round((kill - before) / 1000)), 'kill' if kill < before else 'wipe')
    pulls = sum(1 for x in f['fights'] if x['encounter_id'] is not None)
    tops = (f' &middot; <span>Top HPS <b style="color:{C[f["top_hps"]["class"]]}">{f["top_hps"]["name"]}</b> {f["top_hps"]["value"]:.0f}</span>'
            f' &middot; <span>Top DPS <b style="color:{C[f["top_dps"]["class"]]}">{f["top_dps"]["name"]}</b> {f["top_dps"]["value"]:.0f}</span>')
    sub = f'{pulls} pull{"" if pulls == 1 else "s"}' + (' &middot; <span class="gold">best kill</span>' if kill == best else '') + tops
    night = night_day(at_zone(f['created_at']).strftime('%Y-%m-%d'))
    absent = len(misses['absent'])
    return StatRow(icon(ONY_ICON, 28, 'q5'), f'<button type="button" class="nbtn" aria-expanded="false">{night}</button>', sub,
                   [tone(clock(kill), 'gold') if kill == best else clock(kill), delta, str(f['wipe_count']), str(f['deaths']), str(f['raiders']),
                    tone(str(absent), 'wipe') if absent else str(absent), str(loot)], meter=ws_prims.Meter(kill / slowest, 'gold'))


def night_open(f: dict, misses: dict) -> str:
    """The opened night: its raid page, who was absent and benched, one pull a line."""
    fights = []
    for x in f['fights']:
        dur = clock(x['duration_ms'])
        if x['encounter_id'] is None:
            fights.append(f'<li class="pl"><span class="mute">{x["name"]}</span><span class="mute r">{dur}</span><span class="mute">trash</span></li>')
        else:
            fights.append(f'<li class="pl"><b>{x["name"]} &middot; pull 1</b><span class="mute r">{dur}</span><span class="r">{Tag("Kill", "kill")}</span></li>')
    names = ', '.join(misses['absent'])
    bench = ', '.join(misses['bench'])
    return (f'<div class="nopen"><a class="link" href="#">Open the raid</a>'
            f'<p class="t-small"><span class="mute">Absent</span> <span class="wipe">{names}</span> &middot; <span class="mute">Bench</span> {bench}</p>'
            f'<ul>{"".join(fights)}</ul></div>')


def raids() -> str:
    """The officer's Raids tab (GuildRaids): Planned, then the logged nights as one dense table, the last night open."""
    page = fixture('dense-raids.json')
    grid = fixture('dense-attendance.json')
    loot_nights = {n['report_id']: len(n['decisions']) for n in fixture('dense-loot-nights.json')['nights']}
    misses = {r['id']: {'absent': [x['name'] for x in grid['rows'] if x['cells'][i] == 'absent'],
                        'bench': [x['name'] for x in grid['rows'] if x['cells'][i] == 'bench']} for i, r in enumerate(grid['raids'])}
    rows = page['rows']
    kills = [next(x['duration_ms'] for x in r['fights'] if x['name'] == 'Onyxia' and x['kill']) for r in rows]
    slowest, best = max(kills), min(kills)
    body = []
    for i, r in enumerate(rows):
        before = kills[i + 1] if i + 1 < len(rows) else None
        row = night_row(r, slowest, best, before, loot_nights.get(r['id'], 0), misses[r['raid_id']])
        if i == 0:
            row = row.replace('aria-expanded="false"', 'aria-expanded="true"', 1) + night_open(r, misses[r['raid_id']])
        body.append(row)
    table = dense(StatTable([('', False), ('Night', False), ('Boss time', False), ('Kill time', True), ('vs last', True), ('Wipes', True),
                             ('Deaths', True), ('Raiders', True), ('Absent', True), ('Loot', True)], body, 7))
    meta = f'{len(rows)} nights &middot; {len(rows)} kills &middot; best {clock(best)}'
    planned = ''.join(planned_row(r) for r in page['planned'])
    main = (Section('Planned', f'<div>{planned}</div>', _act('Plan an event', primary=True), meta=str(len(page['planned'])), panel=True)
            + Section('Raid nights', table, meta=meta, panel=True, flush=True))
    ledger = f'<div class="pc" style="padding-top:28px"><div style="display:flex;flex-direction:column;gap:22px">{main}</div></div>'
    return build_page('Guild raids, dense', _view(head('Raids', True) + ledger) + RAIDS_CSS + ICON_CSS + _ROLE_CSS, 'Guild', VIEW_OFFICER)


RAIDS_CSS = (RAIDS_PRIM.replace('minmax(max-content,1.2fr)', 'minmax(340px,1.2fr)')
             + '<style>.nbtn{all:unset;cursor:pointer;font:700 14px/18px var(--font-body);color:var(--strong)}'
             '.nopen{padding:4px 16px 12px 56px;border-bottom:1px solid var(--line-soft)}.nopen>a{display:inline-flex;align-items:center;min-height:36px;font-size:13px}'
             '.nopen p{line-height:20px}.nopen ul{list-style:none;margin:0;padding:0;max-width:560px}'
             '.nopen .pl{display:grid;grid-template-columns:minmax(0,1fr) 48px 56px;align-items:center;gap:12px;min-height:32px;font-size:13px}'
             '.nopen .r{text-align:right;justify-self:end}</style>')


# ---- the raid planner (raid 45)
ROLE_WORD = {'tank': 'Tank', 'healer': 'Healer', 'dps': 'DPS'}
SIGNUP_WORDS = ['In', 'Late', 'Bench', 'Out']
CELL_GLYPH = {'present': '&#10003;', 'bench': 'B', 'absent': '&times;', 'none': ''}
CELL_WORD = {'present': 'Present', 'bench': 'Bench', 'absent': 'Absent', 'none': 'Before joining'}
FACT_GRID = ('display:grid;grid-template-columns:84px 72px 64px 48px 64px;column-gap:12px;align-items:center;text-align:right;'
             'font-variant-numeric:tabular-nums')
PLANNER_CSS = ('<style>.att{display:inline-flex;gap:2px;vertical-align:middle}.att i{width:14px;height:14px;border-radius:2px;font:700 10px/14px var(--font-body);'
               'font-style:normal;text-align:center;color:#07090d}.att i.present{background:var(--kill)}.att i.absent{background:var(--wipe)}'
               '.att i.bench{background:var(--gold)}.att i.none{background:transparent;box-shadow:inset 0 0 0 1px var(--line)}'
               '.att-legend{display:inline-flex;flex-wrap:wrap;gap:4px 8px;font-size:12px;color:var(--muted)}.att-key{display:inline-flex;gap:4px;align-items:center}'
               '.sgr .g.spec{grid-template-columns:minmax(180px,1fr) auto 220px;gap:16px}.sgr .g.spec>input{grid-column:1/-1}'
               '.pl1{display:block;overflow:hidden;text-overflow:ellipsis;white-space:nowrap}</style>')


def strip_html(cells: list[str]) -> str:
    words = ', '.join(CELL_WORD[c] for c in cells)
    return f'<span class="att" role="img" aria-label="{words}" title="{words}">' + ''.join(f'<i class="{c}">{CELL_GLYPH[c]}</i>' for c in cells) + '</span>'


def legend_html() -> str:
    return '<span class="att-legend">' + ''.join(f'<span class="att-key"><span class="att"><i class="{c}" aria-hidden="true">{CELL_GLYPH[c]}</i></span>{CELL_WORD[c]}</span>'
                                                 for c in ('present', 'bench', 'absent', 'none')) + '</span>'


def planner_facts(raid_officer: bool) -> dict:
    """plannerFacts: each roster row's last five nights (officers), item level against the median, the checks it fails and why
    (officers' readiness for the raid, a member's own checklist), rating and 30 nights."""
    home = fixture('dense-home-officer.json' if raid_officer else 'dense-home-member.json')
    roster = home['roster']
    levels = sorted(r['item_level'] for r in roster)
    median = levels[len(levels) // 2]
    readiness = {r['character_key']: r for r in fixture('dense-readiness-raid.json')['rows']} if raid_officer else {}
    nights = {r['account_key']: r['cells'] for r in fixture('dense-attendance.json')['rows']} if raid_officer else {}
    checklist = None if raid_officer else fixture('dense-checklist.json')
    out = {}
    for r in roster:
        why = None
        if r['character_key'] in readiness:
            why = shortfalls(readiness[r['character_key']])
        elif checklist is not None and checklist['character']['character_key'] == r['character_key']:
            why = [w for i in checklist['items'] if i['state'] == 'fail' for w in check_words(i)]
        att = r['attendance_30']
        out[r['character_key']] = {'cells': nights.get(r['account_key']), 'ilvl': r['item_level'], 'delta': r['item_level'] - median,
                                   'why': why, 'rating': r['rating']['overall'], 'pct': round(100 * att['present'] / att['nights'])}
    return out


def shortfalls(row: dict) -> list[str]:
    """readinessShortfalls (attunement asked, upgrades not a shortfall)."""
    out = []
    missing = len(row['enchants']['missing_slots'])
    if missing:
        out.append(f'{missing} enchant{"" if missing == 1 else "s"} missing')
    ff = row['flask_food']
    lacks = (['no flask or elixir'] if not ff['flask_or_elixir'] else []) + (['no food'] if not ff['food'] else [])
    out += lacks
    if row['talent_points_unspent']:
        n = row['talent_points_unspent']
        out.append(f'{n} talent point{"" if n == 1 else "s"} unspent')
    if not row['attuned']:
        out.append('not attuned')
    return out


def check_words(item: dict) -> list[str]:
    if item['key'] == 'enchants':
        return [f"{len(item.get('missing_slots', []))} enchants missing"]
    if item['key'] == 'flask_food':
        return ['no food' if m == 'food' else 'no flask or elixir' for m in item.get('missing', [])]
    if item['key'] == 'attunement':
        return ['not attuned']
    return []


def facts_cells(f: dict) -> str:
    strip = strip_html(f['cells']) if f['cells'] else ''
    ready = ('<span class="mute">&ndash;</span>' if f['why'] is None else '<span class="kill" aria-label="Ready">&#10003;</span>' if not f['why']
             else f'<span class="wipe" title="{", ".join(f["why"])}">{len(f["why"])} short</span>')
    delta = f'&minus;{abs(f["delta"])}' if f['delta'] < 0 else f'+{f["delta"]}'
    return (f'<span style="{FACT_GRID};font-size:14px"><span style="display:flex;justify-content:flex-end">{strip}</span>'
            f'<span><b>{f["ilvl"]}</b><span class="mute" style="margin-left:4px;font-weight:400">{delta}</span></span>{ready}'
            f'<span>{f["rating"]}</span><span>{f["pct"]}%</span></span>')


def signup_line(m: dict, f: dict, officer: bool, sub: str = '') -> str:
    """One PlannerRow: crest, the class-colour name, `Holy Paladin · Healer`, what it is short of, the note; In / Late / Bench / Out
    (officers) or its state Tag, then the figures."""
    line = sub or f"{m['spec']} {m['class'].title()} &middot; {ROLE_WORD[m['role']]}"
    why = f'<span class="wipe" style="display:block">{", ".join(f["why"])}</span>' if f['why'] else ''
    note = f'<span style="display:block;color:var(--text)">{m["note"]}</span>' if m.get('note') else ''
    if officer:
        state = Segmented([(w, w.lower() == m.get('state')) for w in SIGNUP_WORDS], small=True)
    else:
        state = Tag(m['state'].title(), {'in': 'kill', 'late': 'gold', 'bench': 'muted', 'out': 'muted'}[m['state']]) if m.get('state') else ''
    return Row(m["name"], f"<span class=\"pl1\">{line}</span>" + why + note, lead=crest_img(m['class'], 28), title_color=C[m['class']], trail=state + facts_cells(f), link=False,
               slim=True, selected=m.get('own', False))


def signup_groups(raid: dict, facts: dict, officer: bool) -> str:
    rows = raid['signups']
    on = [m for m in rows if m['state'] == 'in']
    groups = [(label, [m for m in on if m['group'] == g]) for label, g in [('Tanks', 'tank'), ('Healers', 'healer'), ('Melee', 'melee'), ('Ranged', 'ranged')]]
    groups += [(label, [m for m in rows if m['state'] == st]) for label, st in [('Late', 'late'), ('Bench', 'bench'), ('Out', 'out')]]
    head = (f'<div style="display:flex;align-items:center;justify-content:space-between;gap:16px;padding-top:4px">{legend_html() if officer else "<span></span>"}'
            f'<span aria-hidden="true" style="{FACT_GRID};font:700 12px/16px var(--font-body);color:var(--muted)">'
            + ''.join(f'<span>{h}</span>' for h in ('Last 5', 'Item level', 'Ready', 'Rating', '30 nights')) + '</span></div>')
    return head + ''.join(RowGroup(label, len(ms), [signup_line(m, facts[m['character_key']], officer) for m in ms])
                          for label, ms in groups if ms or label in ('Tanks', 'Healers', 'Melee', 'Ranged'))


def planner_when(raid: dict) -> str:
    start = at_zone(raid['starts_at'])
    end = start + timedelta(minutes=raid['duration_minutes'])
    clock_word = gd.house_clock
    left = start - at_zone(D.NOW)
    return (f"{start:%A} {start.day} {start:%B} &middot; {clock_word(start)} to {clock_word(end)} CST &middot; "
            f'<span class="gold">in {gd.countdown_text(left)}</span>')


def planner(officer: bool = True) -> str:
    """The raid planner (RaidPlanner) for raid 45 as its officer and as a member: the head, the grouping bar, your sign-up, the
    sign-ups with their figures, the Discord post and who has not answered (officers); the rail's comp, bench and notes."""
    raid = fixture('dense-raid-officer.json' if officer else 'dense-raid-member.json')
    facts = planner_facts(officer)
    n = raid['counts']
    signed = n['in'] + n['late']
    you = raid['viewer']['characters'][0]
    bar = control_bar([('Group by', seg([('Status', False, None, ''), ('Role', True, None, ''), ('Class', False, None, '')], label='Group by'))],
                      'Group by role', label='Sign-up grouping')
    own = '<ul>' + SignupRow(crest_img(you['class'], 32), you['name'], C[you['class']], f"{you['spec']} {you['class'].title()}",
                             Segmented([(w, w.lower() == you['state']) for w in SIGNUP_WORDS]), Select(f"{you['spec']} &middot; {ROLE_WORD[you['role']]}")) + '</ul>'
    signups = Section('Sign-ups', f'<div role="region" aria-label="Sign-ups">{signup_groups(raid, facts, officer)}</div>', meta=f"{signed}/{raid['size']}",
                      panel=True)
    main = Section('Your sign-up', own, named=True, panel=True) + signups
    if officer:
        posted = at_zone(raid['discord_post']['posted_at'])
        main += (f'<p class="t-small" style="display:flex;align-items:center;gap:12px">Posted to #{raid["discord"]["channel_name"]} '
                 f'{posted:%a} {posted.day} {posted:%b}, {posted.hour % 12 or 12} {"am" if posted.hour < 12 else "pm"}. {_act("Repost")}</p>')
        roster = {m['character_key']: m for m in fixture('dense-home-officer.json')['roster']}
        ns = []
        for r in raid['not_signed_raiders']:
            m = roster[r['character_key']]
            sub = f"{m['spec']} {m['class'].title()} &middot; {ROLE_WORD[m['role']]}" + ('' if r['discord_linked'] else ' &middot; No Discord linked')
            ns.append(signup_line({**m, 'state': None}, facts[r['character_key']], True, sub))
        main += Section('Not signed up', Rows(ns, open_=True, list_=True), action=_act('Ping on Discord'), meta=str(len(ns)), panel=True)
    comp = comp_grid(n, raid['comp_target'])
    # CompTally's edit slot is always passed: a member reads its padding with no button in it.
    comp += f'<div style="padding:16px 20px 8px">{_act("Change the target") if officer else ""}</div>'
    rail = Section('Comp against the target', comp, panel=True, flush=True)
    if officer:
        bench = [Row(m['name'], f"{m['spec']} {m['class'].title()} &middot; {ROLE_WORD[m['role']]}", lead=crest_img(m['class'], 24), title_color=C[m['class']],
                     trail=_act('Swap in'), link=False, slim=True) for m in raid['signups'] if m['state'] == 'bench']
        rail += Section('Bench and swaps', Rows(bench, open_=True, list_=True), meta=str(len(bench)), panel=True)
    empty = 'No notes yet. Edit the raid to add some.' if officer else 'No notes for this raid.'
    rail += Section('Notes', EmptyState(empty), named=True, panel=True)  # RaidPlanner.svelte: the shared empty line
    actions = (''.join(_act(x) for x in ['Edit', 'Lock sign-ups', 'Copy the sign-up list']) + Button('Cancel raid', small=True, tone='ghost', act=True)
               if officer else '')
    hdr = PageHeader(raid['title'], ALLIANCE, meta=[Tag('Planned'), planner_when(raid),
                                                    f"<b class=\"strong\">{signed} of {raid['size']} signed up</b> &middot; {n['in']} In &middot; {n['late']} Late &middot; {n['bench']} Bench"],
                     actions=actions, tabs=TabStrip(gd.TABS if officer else [t for t in gd.TABS if t != 'Settings'], 'Raids'),
                     crumb=f'<a href="#">&lsaquo; {gd.GUILD} &middot; Raids</a>')
    body = f'<div style="padding-bottom:40px">{hdr + bar + Page("split", main, rail, sticky_rail=True, panels=True)}</div>{GUILD_CSS}'
    who = VIEW_OFFICER if officer else VIEW_MEMBER
    return build_page('Raid planner, dense', body + PLANNER_CSS + _ROLE_CSS, 'Guild', who)


# ---- Progression
def night_day(date: str) -> str:
    """progressionNightDay: 'Wed 6 Jan', with the year when it is not the clock's."""
    d = datetime.strptime(date, '%Y-%m-%d')
    return f"{d:%a} {d.day} {d:%b}" + ('' if d.year == 2027 else f' {d.year}')


def signed_s(delta: int) -> str:
    return f'{"&minus;" if delta < 0 else "+" if delta > 0 else ""}{abs(delta)} s'


def boss_nights_table(enc: dict) -> str:
    """EncounterProgress's table: every night the boss was pulled, newest first, its kill time on a Meter against the slowest."""
    nights = enc['pulls_by_night']
    kills = [n['kill_ms'] for n in nights if n['kill_ms'] is not None]
    slowest, best = max(kills), min(kills)
    rows, before = [], None
    for i, n in enumerate(nights):
        k = n['kill_ms']
        delta = '&ndash;' if k is None or before is None else tone(signed_s(round((k - before) / 1000)), 'kill' if k < before else 'wipe')
        before = k if k is not None else before
        sub = 'first kill' if i == 0 else 'best kill' if k == best else ''
        rows.append(StatRow('', f'<b>{night_day(n["date"])}</b>', sub,
                            [tone(clock(k), 'gold') if k == best else clock(k), delta, str(n['pulls']), str(n['pulls'] - 1), str(n['deaths'])],
                            meter=ws_prims.Meter(k / slowest, 'gold')))
    rows.reverse()
    return dense(StatTable([('', False), ('Night', False), ('Kill time', False), ('Kill', True), ('vs last', True), ('Pulls', True),
                            ('Wipes', True), ('Deaths', True)], rows, 5))


SPEC_WORD = {('druid', 'Feral'): 'Feral Druid'}


def spec_class(spec: str, cls: str) -> str:
    return f'{spec} {cls.title()}'


def roster_bests(rows: list[dict]) -> str:
    """Roster bests (GuildProgression): the guild's best kill figure per raider, damage then healing, ten and All."""
    ranked = sorted(rows, key=lambda r: (r['metric'], -r['value']))
    shown = [StatRow(crest_img(r['player']['class'], 28), f'<b><a href="#" style="color:{C[r["player"]["class"]]}">{r["player"]["name"]}</a></b>',
                     f'{spec_class(r["player"]["spec"], r["player"]["class"])} &middot; {round(r["execution_score"] * 100)}% of sim',
                     [f'{r["value"]:.0f} {r["metric"].upper()}'], bar=False) for r in ranked[:10]]
    table = dense(StatTable([('', False), ('Raider', False), ('Best', True)], shown, 1, bar=False))
    return table + f'<button type="button" class="link" style="min-height:44px;padding:0 20px">All {len(rows)}</button>'


def best_by_role(enc: dict) -> str:
    word = {'tank': 'Tank', 'healer': 'Healer', 'dps': 'Damage'}
    rows = [StatRow(crest_img(b['class'], 28), f'<b style="color:{C[b["class"]]}">{b["name"]}</b>',
                    f'{word[role]} &middot; {spec_class(b["spec"], b["class"])}', [f'{b["value"]:.0f} {b["metric"].upper()}'], bar=False)
            for role in ('tank', 'healer', 'dps') if (b := enc['best_by_role'].get(role))]
    return dense(StatTable([('', False), ('', False), ('', True)], rows, 1, bar=False))


def progression() -> str:
    """Progression (GuildProgression): each boss's panel and its nights, the roster bests under it; the rail's tier and the best
    per role. Drawn from dense-progression.json and the guild page's roster_best."""
    prog = fixture('dense-progression.json')
    enc = prog['encounters'][0]
    best_night = min((n for n in enc['pulls_by_night'] if n['kill_ms']), key=lambda n: n['kill_ms'])
    first, best = enc['pulls_by_night'][0]['kill_ms'], enc['best_kill_ms']
    meta = (f"{enc['kills']} kills &middot; best {clock(best)}, {night_day(best_night['date'])} &middot; {round((first - best) / 1000)} s faster than "
            f"the first &middot; {enc['deaths_per_pull']:.1f} deaths a pull")
    bests = fixture('dense-page.json')['roster_best']
    main = (Section(enc["name"], boss_nights_table(enc), meta=meta, panel=True, flush=True)
            + Section('Roster bests', roster_bests(bests), panel=True, flush=True))
    tier = prog['tier']
    zone = (f'<ul><li style="display:flex;flex-direction:column;gap:8px;padding-bottom:12px"><span style="display:flex;justify-content:space-between">'
            f'<b>{enc["zone"]}</b><span class="t-small">{tier["down"]} of {tier["named_encounters"]} boss down</span></span>'
            f'{ws_prims.Meter(tier["down"] / tier["named_encounters"], "kill", f"{tier["down"]} of {tier["named_encounters"]} named bosses down")}</li></ul>')
    rail = (Section(tier['name'], zone, meta=f"{tier['pulls_total']} pulls", panel=True)
            + Section(f'Best on {enc["name"]}', best_by_role(enc), panel=True, flush=True))
    return build_page('Guild progression, dense', _view(head('Progression', True) + Page('split', main, rail, panels=True)) + _ROLE_CSS
                      + ICON_CSS, 'Guild', VIEW_OFFICER)


# ---- Applications
GROUP_OF_SPEC = {'paladin-holy': 'Healer', 'warlock-destruction': 'Ranged', 'warrior-protection': 'Tank', 'druid-restoration': 'Healer',
                 'hunter-marksmanship': 'Ranged', 'rogue-combat': 'Melee'}


def applicant_row(a: tuple) -> str:
    (aid, name, cls, spec, spec_name, ilvl, metric, value, pct, days, avail, message, yes, no, mine, status, decided) = a
    g = GROUP_OF_SPEC[spec]
    gap = gap_of(g)
    fills = f'<span class="gold">{GROUP_WORD[g]}, {gap} short</span>' if gap > 0 else f'<span class="mute">{GROUP_WORD[g]}, full</span>'
    actions = Button('Accept', True, True, act=True) + ' ' + Button('Decline', False, True, tone='ghost', act=True)
    sub = f'{spec_name} {fills}<br>' + (message or 'No message') + (f' &middot; Can raid: {avail}' if avail else '')
    return StatRow(crest_img(cls, 28), f'<b style="color:{C[cls]}">{name}</b>', sub,
                   [f'{ilvl} <span class="mute">{ilvl - D.MEDIAN_ILVL:+d}</span>', tone(str(pct), parse_class(pct)) + f' <span class="mute">{value:.0f}</span>',
                    f'{days} d', f'{yes}&ndash;{no} <span class="mute">of {VOTERS}</span>', actions], bar=False, primary_phone=f'{pct} &middot; {days} d')


GROUP_SEAT = {'tank': 'Tank', 'healer': 'Healer', 'melee': 'Melee', 'ranged': 'Ranged'}
APPLICANT_GRID = ('<span style="display:grid;grid-template-columns:72px 76px 56px;column-gap:12px;align-items:center;text-align:right;'
                  'font-size:14px;font-variant-numeric:tabular-nums;white-space:nowrap">')
APPLICANT_ACTIONS = '<span style="display:flex;align-items:center;justify-content:flex-end;gap:8px;width:272px">'
# The percentile's quality colour as a swatch beside the strong figure (ApplicantRow): text in the colour fails contrast.
SWATCH_CSS = ('<style>.pswatch{display:inline-block;width:4px;height:12px;border-radius:2px;margin-right:4px}'
              '.pswatch.p-orange{background:#ff8000}.pswatch.p-purple{background:#a335ee}.pswatch.p-blue{background:#0070ff}'
              '.pswatch.p-green{background:#1eff00}.pswatch.p-grey{background:#8c8c8c}</style>')


def applicant_row(a: dict, short: dict) -> str:
    """One open application (ApplicantRow from 1280): spec and the seat it fills, availability then the message; the figures
    (item level against the median, best percentile and figure, days waiting), the votes, Accept and Decline."""
    c = a['character']
    seat = short.get(c['group'])
    fill = (f' <span class="gold" style="white-space:nowrap">{GROUP_WORD[GROUP_SEAT[c["group"]]]}, {seat} short</span>' if seat
            else f' <span class="mute" style="white-space:nowrap">{GROUP_WORD[GROUP_SEAT[c["group"]]]}, full</span>')
    avail = f"Can raid: {a['availability']} &middot; " if a['availability'] else ''
    sub = f"{c['spec_name']}{fill}<br>{avail}{a['message'] or 'No message.'}"
    parse = c['best_parse']
    pct = round(parse['percentile'])
    days = (datetime.fromisoformat(D.NOW.replace('Z', '+00:00')) - datetime.fromisoformat(a['created_at'].replace('Z', '+00:00'))).days
    figs = (APPLICANT_GRID + f'<span><b>{c["item_level"]}</b> <span class="mute">{c["item_level"] - D.MEDIAN_ILVL:+d}</span></span>'
            f'<span><i class="pswatch {parse_class(pct)}"></i>'
            f'<b>{pct}</b> <span class="mute">{parse["value"]:.0f}</span></span><span>{days} d</span></span>')
    # ui/Votes wraps (flex-wrap): in the 272 px action column its chips and line stack, as the build lays them out.
    votes = _votes(a['votes']['yes'], a['votes']['no'], a['votes']['voters']).replace('display:flex;', 'display:inline-flex;flex-wrap:wrap;flex:0 0 76px;min-width:0;row-gap:4px;', 1)
    acts = APPLICANT_ACTIONS + votes + _act('Accept') + Button('Decline', small=True, tone='ghost', act=True) + '</span>'
    return Row(c['name'], sub, lead=crest_img(c['class'], 28), title_color=C[c['class']], trail=figs + acts, link=False, slim=True)


def applications() -> str:
    """Applications (GuildApplications): the open applications on one figure grid under a head line; the rail's open seats for the
    next raid with who applies for each, the public card editor and the recent decisions. Drawn from dense-applications.json."""
    rows = fixture('dense-applications.json')['rows']
    pending = [a for a in rows if a['status'] == 'pending']
    decided = [a for a in rows if a['status'] != 'pending']
    short = {'tank' if g == 'Tank' else 'healer' if g == 'Healer' else g.lower(): n for g, n in ((g, gap_of(g)) for g in D.GROUPS) if n > 0}
    head = ('<div aria-hidden="true" style="display:flex;justify-content:flex-end;column-gap:12px;padding-bottom:4px">'
            + APPLICANT_GRID.replace('font-size:14px', 'font:700 12px/16px var(--font-body);color:var(--muted)')
            + '<span>Item level</span><span>Best %</span><span>Waiting</span></span><span style="width:272px"></span></div>')
    plate = f'<div role="region" aria-label="Open applications">{head}{Rows([applicant_row(a, short) for a in pending])}</div>'
    main = Section('Open applications', plate, panel=True)
    gaps = Rows([Row(GROUP_WORD[g], f"{D.group_counts(g)['in'] + D.group_counts(g)['late']} of {D.COMP_TARGET[D.GROUP_TARGET[g]]} &middot; "
                     f"{sum(1 for a in pending if a['character']['group'] == g.lower())} applying",
                     trail=f'<span class="wipe" style="font-weight:700">{gap_of(g)} short</span>', link=False, slim=True)
                 for g in D.GROUPS if gap_of(g) > 0], list_=True)
    wanted = ', '.join(w[2] for w in WANTED)
    card = Rows([Row('Looking for', wanted, trail=_act('Edit'), link=False), Row('Raid nights', gd.RECRUIT_NIGHTS, trail=_act('Edit'), link=False)],
                open_=True)
    record = Rows([Row(f"{a['character']['name']} {a['status']}", f"{a['votes']['yes']} yes, {a['votes']['no']} no of {a['votes']['voters']} officers &middot; by {a['decided_by']['name']}",
                       link=False, slim=True) for a in decided], open_=True)
    rail = (Section(f'Short for {day(D.NEXT_RAID_AT)}', gaps, panel=True) + Section('Your public card', card, Link('See it as a visitor'), panel=True)
            + Section('Recent decisions', record, panel=True))
    views = TabStrip(['Members', 'Applications'], 'Applications').replace('strip-tabs ph-tabs', 'strip-tabs', 1)
    hdr = PageHeader(gd.GUILD, '', ws_prims.Emblem('alliance', 56), meta=[band_facts(None)], tabs=TabStrip(gd.TABS, 'Roster') + views)
    return build_page('Guild applications, dense', _view(hdr + Page('split', main, rail, panels=True)) + _ROLE_CSS + SWATCH_CSS, 'Guild',
                      VIEW_OFFICER)


# ---- Recruit (the public card, read by its officer)
def recruit() -> str:
    """The public card read by its officer (GuildRecruit): Looking for, then the open applications line with Review; the rail's
    guild facts, why this guild and the officers. Drawn from the card the build reads (dense-recruit.json)."""
    card = fixture('dense-recruit.json')
    wanted = Rows([Row(w['spec_name'], w['note'], lead=crest_img(w['class'], 30), title_color=C[w['class']], link=False)
                   for w in card['wanted']], list_=True)
    upcoming = ', '.join(f"{_day(at)}, {clock_label(guild_time(at)).rsplit(' ', 1)[0]}" for at in card['raid_nights']['upcoming'])
    t = card['progress']
    facts = [('Region', f'{gd.REGION.upper()} PvP'), ('Faction', gd.FACTION.title()), ('Members', str(card['members'])),
             ('Raids', card['raid_nights']['text']), ('Next raid nights', upcoming),
             ('Progress', f"{t['name']} &middot; {t['down']} of {t['named_encounters']} boss{'' if t['named_encounters'] == 1 else 'es'} down")]
    facts_html = '<dl class="gfacts">' + ''.join(f'<dt>{k}</dt><dd{" style=white-space:nowrap" if k == "Progress" else ""}>{v}</dd>'
                                                  for k, v in facts) + '</dl>'
    rank = {'leader': 'Guild master', 'officer': 'Officer'}
    officers = Rows([Row(o['name'], f"{o['spec']} {o['class'].title()} &middot; {rank[o['rank']]}", lead=crest_img(o['class'], 32),
                         title_color=C[o['class']], link=False) for o in card['officers']], open_=True)
    pending = sum(1 for a in APPLICATIONS if a[15] == 'pending')
    notice = f'<p class="notice neutral">{pending} open applications {Link("Review &rsaquo;")}</p>'
    main = Section('Looking for', wanted, panel=True) + notice
    rail = (Section(gd.GUILD, facts_html, panel=True) + Section('Why this guild', f'<p {_PROSE}>{card["pitch"]}</p>', panel=True)
            + Section('Officers', officers, panel=True))
    hdr = PageHeader(f'Join {gd.GUILD}', '', ws_prims.Emblem(gd.FACTION, 56), meta=[Tag('Recruiting', 'kill')],
                     tabs=TabStrip(gd.TABS, 'Roster'))
    return build_page('Guild recruit, dense', _view(hdr + Page('split', main, rail, panels=True)) + _FACTS_CSS + _ROLE_CSS, 'Guild',
                      VIEW_OFFICER)


# ---- Settings
def seat_rows(view: dict) -> str:
    """The plan's Premium seats (GuildSeats): a held seat movable now, the open seats with Give seat."""
    out = []
    for seat in view['seats']:
        if seat['holder'] is None:
            who = '<span class="sp-empty" aria-hidden="true"></span><span class="t-small" style="font-size:14px">Open seat</span>'
            give = _act('Give seat', small=False).replace('<button ', '<button aria-expanded="false" ', 1)
            out.append(f'<li class="sp-seat"><div class="sp-who">{who}</div><div></div><span>{give}</span></li>')
            continue
        h = seat['holder']
        who = (f'{crest_img(h["class"], 32).replace("<img ", "<img aria-hidden=\"true\" ", 1)}<div><b style="color:{C[h["class"]]}">{h["name"]}</b>'
               f'<div class="t-small" style="font-size:12px">Given {WS.day(seat["given_at"])}</div></div>')
        move = _act('Move seat', small=False).replace('<button ', '<button aria-expanded="false" ', 1)
        out.append(f'<li class="sp-seat"><div class="sp-who">{who}</div>{WS.p("Can move now")}<span>{move}</span></li>')
    return f'<ul class="sp-seats">{"".join(out)}</ul>'


def plan_section(view: dict) -> str:
    """GuildBillingSection: the founder Raid guild billed to the viewer, Manage billing, Who pays, the seats. One frame."""
    plan = view['billing']
    end = datetime.fromisoformat(plan['current_period_end'].replace('Z', '+00:00'))
    renews = f"Renews {end.day - 1} {end:%b %Y} &middot; ${plan['amount_cents'] // 100} a year"
    given = sum(1 for x in view['seats'] if x['holder'] is not None)
    rows = Rows([Row(renews, trail=Tag('Founder'), link=False), Row('Billed to you.', link=False)], open_=True)
    who = Rows([Row('Who pays', trail=_act('Transfer the plan'), link=False)], open_=True)
    seats = (f'<div style="display:flex;flex-direction:column;gap:8px;padding-top:20px"><h3 class="sp-h3">Premium seats '
             f'<span class="t-small" style="font-weight:400;font-size:13px">{given} of {len(view["seats"])} given</span></h3>{seat_rows(view)}</div>')
    return Section('Guild plan', WS.body(rows + f'<div class="sp-btns" style="padding-bottom:12px">{_act("Manage billing")}</div>' + who + seats),
                   meta='Raid guild', panel=True)


def discord_bot_section(settings_view: dict) -> str:
    """GuildDiscordBot, installed: where raids post, the linked members, its three actions."""
    b = settings_view['discord']['bot']
    linked = f"{b['linked_members']} of {b['members']} members have linked Discord. The others get a link when they press a button."
    dest = (f'<div style="display:flex;flex-direction:column;gap:2px">'
            f'{WS.p(f"Posts raids to <b class=\"strong\">#{b["channel_name"]}</b> in <b class=\"strong\">{b["server_name"]}</b>", small=False)}'
            f'{WS.p("With sign-up buttons &middot; <a href=\"#\" style=\"text-decoration:underline\">Open in Discord</a>")}</div>')
    acts = f'<div class="sp-btns" style="padding-top:0">{_act("Send test message", small=False)}{_act("Change channel", small=False)}{_act("Remove", small=False)}</div>'
    inner = f'{dest}{WS.p(linked)}<div style="display:flex;flex-direction:column">{acts}<div style="margin-top:8px;min-height:21px"></div></div>'
    return Section('Discord', WS.body(f'<div style="display:flex;flex-direction:column;gap:12px">{inner}</div>'), action=Tag('Connected', 'kill'),
                   panel=True)


def settings_facts(setup: dict) -> str:
    """The standing facts from the setup read (settings-facts.ts): the officers with the GM marked, invite joins, addon sync."""
    ranks = setup['ranks']
    officers = ', '.join(f"{n} (GM)" if r['rank_index'] == 0 else n for r in ranks['ranks'] if r['rank_index'] <= ranks['officer_max_rank_index']
                         for n in r['names'])
    sync = setup['sync']
    missing = sync['missing']
    shown = ', '.join(missing[:3]) + (f' and {len(missing) - 3} more' if len(missing) > 3 else '')
    rows = [Row(f"{ranks['officers']} officers", officers, link=False, slim=True),
            Row(f"{setup['invite']['joined']} joined by the invite link", link=False, slim=True),
            Row(f"{sync['synced']} of {sync['accounts']} synced in {sync['window_days']} days", f'Not synced: {shown}' if missing else '', link=False, slim=True)]
    return f'<div style="padding-bottom:8px">{Rows(rows, list_=True)}</div>'


def members_section(settings_view: dict, setup: dict) -> str:
    holder = settings_view['claimed_by']
    claimed = (f'<p style="font-size:13px;line-height:20px;margin:0">Claimed by {crest_img(holder["class"], 20).replace("<img ", "<img aria-hidden=\"true\" ", 1)} '
               f'<b style="color:{C[holder["class"]]}">{holder["name"]}</b></p>')
    invite = (f'<div class="sp-col"><label style="display:flex;flex-direction:column;gap:4px;width:100%"><span class="t-small" style="font-size:13px;font-weight:600">Invite link</span>'
              f'<input class="field sp-in" readonly value="https://foreversixty.gg{settings_view["invite"]["url"]}" style="height:40px"></label>{_act("Copy link", small=False)}'
              f'{WS.p(f"Made {WS.day(settings_view["invite"]["rotated_at"])} by <b style=\"color:{C[holder["class"]]}\">{holder["name"]}</b>.")}{_act("Make a new link")}</div>')
    rank = Row('Officer ranks', trail=f'<input class="field sp-in" type="number" aria-label="Officer ranks" value="{settings_view["officer_max_rank_index"]}" style="width:72px">'
               + _act('Save').replace('<button ', '<button disabled ', 1), link=False)
    vis = Row('Default report visibility', trail=f'<span style="display:inline-block;width:160px">{WS._select("Guild only")}</span>', link=False)
    return Section('Members and access', WS.body(claimed + settings_facts(setup) + invite + Rows([rank, vis], open_=True)), panel=True)


def settings() -> str:
    """The officer's Settings tab (GuildSettingsTab), the bot installed: the setup line naming the step left, the plan and Discord;
    Members and access with the standing facts, then Records and crest. Drawn from dense-settings, dense-billing and dense-setup."""
    settings_view, billing, setup = fixture('dense-settings.json'), fixture('dense-billing.json'), fixture('dense-setup.json')
    left = [SETUP_TITLES[x['key']] for x in setup['steps'] if x['state'] == 'todo'][:2]
    line = ' &middot; '.join(['Setup', f"{setup['done']} of {setup['total']} done", *left])
    notice = f'<div class="pc" style="padding-bottom:12px"><p class="notice neutral">{line} {Link("Continue setup &rsaquo;")}</p></div>'
    main = plan_section(billing) + discord_bot_section(settings_view)
    rail = members_section(settings_view, setup) + WS.records_section(crest=True)
    hdr = head('Settings', True) + '<h2 class="sp-sr">Guild settings</h2>'
    return build_page('Guild settings, dense', _view(hdr + notice + Page('split', main, rail, panels=True) + WS.NOTICE_LINE + WS._CSS) + _ROLE_CSS,
                      'Guild', VIEW_OFFICER)


# lib/guild/setup-copy titles.
SETUP_TITLES = {'officer_ranks': 'Officer ranks', 'discord_bot': 'Discord bot', 'recruiting': 'Recruiting', 'raid_schedule': 'Raid schedule',
                'invite_link': 'Guild invite link', 'addon_sync': 'Addon and companion', 'report_visibility': 'Report visibility',
                'linked_role': 'Discord linked role'}


PAGES = {'web-guild-dense-raids': raids, 'web-guild-dense-planner': planner, 'web-guild-dense-raid-member': lambda: planner(False),
         'web-guild-dense-progression': progression, 'web-guild-dense-applications': applications, 'web-guild-dense-recruit': recruit,
         'web-guild-dense-settings': settings}

if __name__ == '__main__':
    import sys

    for n, f in PAGES.items():
        if not sys.argv[1:] or n in sys.argv[1:]:
            print(write(n + '.html', f()))
