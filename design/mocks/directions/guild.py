"""Mock batch 3: the guild in the directions system (plan 2026-11-09-redesign, "Every player-facing surface").

Data, never invented: the guild mocks' own data modules in design/mocks (guild_data, guild_overview_data and
guild_overview_fixtures, guild_tabs_fixtures, guild_raid), the values the guild parity specs serve the build: Olympus
XXVII, 24 raiders, Onyxia's Lair on Thu 10 Dec. Helpers that only format that data (times, the checklist's words, the
sign-up lines) come from design/mocks/ws_guild2.py. Every element is drawn by prims.py.
"""
import re
import sys
from html import escape
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))  # design/mocks: the guild data modules

import guild_data as gd  # noqa: E402
import guild_overview_data as O  # noqa: E402
import guild_overview_fixtures as F  # noqa: E402
import guild_raid  # noqa: E402
import guild_tabs_fixtures as T  # noqa: E402
import ws_guild2 as G  # noqa: E402
from mocklib import emblem  # noqa: E402
import data as D  # noqa: E402
from common import write  # noqa: E402
from prims import (PRIMS_CSS, answers, facts, item_row, meter, need_grid, panel, player_cell, row_list, table, tag,  # noqa: E402
                   votes)
from style import footer, header, page, titlebar  # noqa: E402

OFFICER_VIEW = (O.OFFICER, 'hunter', 60)
SIGNUP = ['In', 'Late', 'Bench', 'Out']
SRC = re.compile(r'src="([^"]+)"')
QUALITY_NAME = {0: 'poor', 1: 'common', 2: 'uncommon', 3: 'rare', 4: 'epic', 5: 'legendary'}

GUILD_CSS = PRIMS_CSS + '''
.gd{padding-top:4px}.strip{display:flex;flex-wrap:wrap;align-items:center;gap:10px 18px;padding:12px 18px;margin-top:22px}
.strip p{font-size:14.5px}.strip .end{margin-left:auto;display:flex;gap:14px;align-items:center}
.emb{width:52px;height:52px;flex:none}.grp{font:600 12.5px var(--body);color:var(--muted);padding:12px 0 2px}
.grp b{color:#fff}.figs{display:flex;gap:28px;flex-wrap:wrap}.figs div{display:flex;flex-direction:column}.figs b{font:600 26px var(--body);color:#fff}
.figs span{font-size:12.5px;color:var(--muted)}.chk li{display:grid;grid-template-columns:18px 1fr auto;gap:10px;padding:7px 0;border-bottom:1px solid rgba(255,255,255,.07)}
.chk li:last-child{border-bottom:0}.chk i{width:10px;height:10px;border-radius:50%;margin-top:6px}.chk i.ok{background:var(--gain)}.chk i.warn{background:var(--gold)}
.chk b{color:#fff;display:block}.chk span{font-size:13px;color:var(--muted)}
@media(max-width:700px){.strip .end{margin-left:0}}
'''


def src(html: str) -> str:
    m = SRC.search(html)
    return m.group(1) if m else ''


def emb() -> str:
    return f'<img class="emb" src="{emblem(gd.FACTION)}" alt="">'


def band(active: str, viewer: tuple = OFFICER_VIEW, title: str = gd.GUILD, meta: str = '') -> str:
    tabs = [[(t, t == active) for t in gd.TABS]]
    return titlebar(None, title, meta or ' &middot; '.join(G.head_facts()), media=emb(), ctl=tabs, top=header(*viewer, 'Guild'))


def raid_strip(viewer: str) -> str:
    """The guild's next raid under the band on every tab: the sentence, the viewer's sign-up, the raid's link."""
    nxt = F.next_raid(viewer)
    at = G.guild_time(O.NEXT_RAID_AT)
    _, _, _, role, _, state = G.ch(viewer)
    n = nxt['counts']['in'] + nxt['counts']['late']
    line = (f'<p><b class="strong">{escape(O.ZONE_NAME)}</b> &middot; {at:%A} {G.clock_label(at)} &middot; '
            f'<span class="num">{n}</span> of {nxt["size"]} signed up</p>')
    return (f'<div class="card strip">{line}<span class="end">{answers(SIGNUP, state)}'
            f'<a href="#">Open the raid &rsaquo;</a></span></div>')


def crest_row(name: str, sub: str = '', size: int = 30) -> str:
    _, cls, *_ = G.ch(name)
    return player_cell(cls, name, sub, size)


def guild_page(title: str, body: str, viewer: tuple = OFFICER_VIEW) -> str:
    return page(f'Forever Sixty, {title}', GUILD_CSS, body + footer())


# ------------------------------------------------------------------------------------------------ Overview (officer)
def overview() -> str:
    a = F.officer_actions()
    weekday = G.guild_time(O.NEXT_RAID_AT).strftime('%A')
    act = row_list([
        ('', f'{a["approvals"]["count"]} raiders waiting for approval', ', '.join(a['approvals']['names']), '<a class="btn2" href="#">Review</a>'),
        ('', f'{a["applications"]["count"]} applications', ', '.join(a['applications']['names']), '<a class="btn2" href="#">Open</a>'),
        ('', f'{a["not_signed"]["count"]} not signed up for {weekday}', ', '.join(a['not_signed']['names']), '<a class="btn2" href="#">Ping on Discord</a>'),
        ('', f'{a["drops_to_award"]["count"]} drops to award', escape(a['drops_to_award']['label']), '<a class="btn2" href="#">Open loot</a>')])
    view = F.raid_view(O.OFFICER, True)
    c = view['counts']
    figs = (f'<div class="figs"><div><b class="num" style="color:var(--gain)">{c["in"]}</b><span>In</span></div><div><b class="num" style="color:var(--gold)">{c["late"]}</b><span>Late</span></div>'
            f'<div><b class="num">{c["bench"]}</b><span>Bench</span></div><div><b class="num" style="color:#f0736b">{len(view["not_signed_raiders"])}</b><span>Not signed</span></div></div>')
    at = G.guild_time(O.NEXT_RAID_AT)
    raid = (f'{row_list([(f"<img class=ai src={D.icon_uri(G.RAID_ICON)} width=40 height=40 alt=>", escape(O.ZONE_NAME), f"{G.day_label(at)} &middot; {G.clock_label(at)}", f"<b class=gold>{G.countdown_label()}</b>")])}'
            f'<div style="margin-top:12px">{figs}</div>')
    comp = comp_table()
    cover = cover_grids()
    last = F.recent_attendance()
    main = (panel('Action needed', act, framed=True) + panel('Next raid', raid, action='<a href="#">Open the planner &rsaquo;</a>')
            + panel('Comp against the target', comp, action='<a href="#">Edit target &rsaquo;</a>', flush=True))
    rail = panel('Buff and debuff cover', cover, flush=True) + panel(f'Attendance, last {len(last["raids"])} raids', attendance_table(last), flush=True)
    body = band('Overview') + f'<main class="w gd"><div class="twocol"><div>{main}</div><aside>{rail}</aside></div></main>'
    return guild_page('guild', body)


def comp_grid(counts: dict, target: dict) -> str:
    """The one comp grid (overview and planner): each role's target against who is In; Late raiders turn a short role gold."""
    late = counts['late_groups']
    rows = []
    for role in ('tanks', 'healers', 'melee', 'ranged'):
        have, need = counts[role], target[role]
        tone = 'ok' if have >= need else ('warn' if have + late[role] >= need else 'miss')
        rows.append((role.title(), f'{late[role]} late' if late[role] else '', [('none', str(need)), (tone, str(have))]))
    return need_grid(['Need', 'In'], rows)


def comp_table() -> str:
    return comp_grid(F.raid_view(O.MEMBER, False)['counts'], O.COMP_TARGET)


def cover_grids() -> str:
    rows = []
    for r in F.cover_rows():
        signed = r['providers_in'] + r['providers_late']
        tone = 'miss' if signed == 0 else ('ok' if signed - r['likely'] >= r['need'] else 'warn')
        rows.append((f'<img class="ai" src="{D.icon_uri(r["icon"])}" width="22" height="22" alt=""> {escape(r["name"])}', '',
                     [('none', str(r['need'])), (tone, str(signed))]))
    return need_grid(['Need', 'In'], rows)


def attendance_table(page_data: dict) -> str:
    mark = {'present': ('ok', '&#10003;'), 'absent': ('miss', '&times;'), 'bench': ('warn', 'B'), 'none': ('none', '&mdash;')}
    cols = [G.guild_time(r['starts_at']).strftime('%a') for r in page_data['raids']]
    rows = [(f'<span style="color:{D.CLASS_COLOR.get(G.ch(r["name"])[1], "#fff")}">{escape(r["name"])}</span>', 'slipping' if r['slipping'] else '',
             [mark[x] for x in r['cells']]) for r in sorted(page_data['rows'], key=lambda r: (not r['slipping'], r['name']))]
    return need_grid(cols, rows)


# ------------------------------------------------------------------------------------------------ Roster
def roster() -> str:
    home = T.roster_home()
    waiting = row_list([(player_cell(r['class'], r['name'], f'{r["spec"]} {r["class"].title()}', 32), '', '',
                         '<a class="btn" href="#">Approve</a> <a class="btn2" href="#">Decline</a>') for r in home['pending']])
    rows = []
    for r in sorted(home['roster'], key=lambda r: (r['name'] != O.OFFICER, -r['item_level'])):
        you = r['name'] == O.OFFICER
        rows.append([player_cell(r['class'], r['name'] + (' (you)' if you else ''), f'{r["spec"]} {r["class"].title()}', 28),
                     tag('Officer', 'gold') if r['rank'] == 'officer' else '<span class="muted">Member</span>',
                     str(r['item_level']), f'{r["attendance"]["present"]} of {r["attendance"]["nights"]}'])
    tbl = table([('Raider', False), ('Rank', False), ('Item level', True), ('Raids', True)], rows, sort='Item level',
                row_cls=['me' if r['name'] == O.OFFICER else '' for r in sorted(home['roster'], key=lambda r: (r['name'] != O.OFFICER, -r['item_level']))],
                hide_phone={1})
    main = (panel('Waiting for approval', waiting, action='<a class="btn2" href="#">Approve all</a>', framed=True)
            + panel('Roster', tbl, flush=True))
    body = band('Roster') + f'<main class="w gd">{raid_strip(O.OFFICER)}{main}</main>'
    return guild_page('guild roster', body)


# ------------------------------------------------------------------------------------------------ Raid planner
def raid_planner() -> str:
    raid = guild_raid.raid_view()
    n = raid['counts']
    by = {m['name']: m for m in guild_raid.MEMBERS}
    you = by[guild_raid.VIEWER]
    order = {'tank': 0, 'healer': 1, 'dps': 2}
    rows = sorted((guild_raid.member(s['name']) for s in raid['signups']), key=lambda m: (order[m['role']], G.CLASS_ORDER.index(m['class']), m['name']))
    groups = ''
    for label, pick in [('Tanks', lambda m: m['state'] == 'in' and m['group'] == 'tank'), ('Healers', lambda m: m['state'] == 'in' and m['group'] == 'healer'),
                        ('Melee', lambda m: m['state'] == 'in' and m['group'] == 'melee'), ('Ranged', lambda m: m['state'] == 'in' and m['group'] == 'ranged'),
                        ('Late', lambda m: m['state'] == 'late'), ('Bench', lambda m: m['state'] == 'bench')]:
        ms = [m for m in rows if pick(m)]
        groups += (f'<p class="grp"><b>{label}</b> {len(ms)}</p>'
                   + row_list([(player_cell(m['class'], m['name'], G.planner_sub(m), 28), '', '', answers(SIGNUP, m['state'] or '')) for m in ms]))
    comp = comp_grid(n, raid['comp_target'])
    enc = raid['assignments']['encounters'][0]
    assign = row_list([('', escape(s['label']),
                        f'{len(s["names"])} of {s["count"]} &middot; {escape(s["purpose"])}<br>'
                        + (' '.join(tag(x['name'], 'gold' if x['own'] else '') for x in s['names']) or 'Open'), '') for s in enc['slots']])
    own = row_list([(player_cell(you['class'], you['name'], f'{you["spec"]} {you["class"].title()}', 32), '', '', answers(SIGNUP, you['state']))])
    signed = n['in'] + n['late']
    main = (panel('Your sign-up', own) + panel('Sign-ups', groups, meta=f'<span class="num">{signed}</span> of {raid["size"]}', framed=True))
    rail = (panel('Comp against the target', comp, flush=True) + panel(f'{enc["name"]} assignments', assign)
            + panel('Notes', f'<div class="prose" style="font-size:14.5px">{raid["notes_html"]}</div>'))
    at = G.guild_time(raid['starts_at'])
    meta = f'{at:%A} {at.day} {at:%B} &middot; {G.clock_label(at)} &middot; <span class="num">{signed}</span> of {raid["size"]} signed up'
    bandh = band('Raids', title=raid['title'], meta=meta)
    body = bandh + f'<main class="w gd"><div class="twocol"><div>{main}</div><aside>{rail}</aside></div></main>'
    return guild_page('raid planner', body, (guild_raid.VIEWER, you['class'], 60))


# ------------------------------------------------------------------------------------------------ Progression and readiness
def progression() -> str:
    pg = F.progression()
    t = pg['tier']
    main = ''
    for e in (x for x in pg['encounters'] if x['pulls'] > 0):
        nights = row_list([('', f'{G.day_label(G.guild_time(n["date"] + "T20:00:00Z"))}', f'{n["pulls"]} pulls', tag('Kill', 'kill') if n['killed'] else '')
                           for n in e['pulls_by_night']])
        main += panel(e['name'], nights + f'<p class="muted" style="font-size:13px;margin-top:8px">{e["deaths_per_pull"]:.1f} deaths a pull</p>',
                      meta=escape(e['zone']), framed=True)
    tier = panel('This tier', f'<div class="figs"><div><b class="num">{t["down"]} of {t["named_encounters"]}</b><span>Named bosses down</span></div></div>'
                 + f'<div style="margin-top:10px">{meter(t["down"] / t["named_encounters"], "var(--gain)")}</div>', meta=escape(t['name']))
    zones = row_list([('', escape(z['zone']), f'{z["pulls"]} pulls', '') for z in pg['unnamed'] if z['pulls'] > 0])
    rail = tier + panel('No encounter list yet', zones)
    body = band('Progression') + f'<main class="w gd">{raid_strip(O.OFFICER)}<div class="twocol"><div>{main}</div><aside>{rail}</aside></div></main>'
    return guild_page('guild progression', body)


def readiness() -> str:
    rows = []
    for r in F.readiness(True)['rows']:
        gap = r['gear_gap']
        rows.append([player_cell(r['class'], r['name'] + (' (you)' if r['name'] == O.OFFICER else ''), f'{r["spec"]} {r["class"].title()}', 26),
                     f'{gap["upgrades"]} upgrades' if gap['upgrades'] else '<span class="muted">None</span>',
                     '<span style="color:var(--gain)">All</span>' if not r['enchants']['missing_slots'] else f'<span style="color:#f0736b">{len(r["enchants"]["missing_slots"])} missing</span>',
                     '<span style="color:var(--gain)">Stocked</span>' if r['consumables']['state'] == 'stocked' else '<span style="color:#f0736b">Short</span>',
                     str(r['talent_points_unspent']) if r['talent_points_unspent'] else '<span class="muted">0</span>', str(r['item_level']),
                     '<a class="btn2" href="#">Nudge</a>'])
    tbl = table([('Raider', False), ('Gear', True), ('Enchants', True), ('Consumables', True), ('Unspent', True), ('Item level', True), ('', True)],
                rows, hide_phone={2, 3, 4, 6})
    nxt = G.guild_time(O.NEXT_RAID_AT)
    body = (band('Readiness') + f'<main class="w gd">{raid_strip(O.OFFICER)}'
            + panel('Readiness', tbl, meta=f'Signed In or Late for {escape(O.ZONE_NAME)}, {G.day_label(nxt)}', flush=True, framed=True) + '</main>')
    return guild_page('guild readiness', body)


# ------------------------------------------------------------------------------------------------ Loot council
def loot() -> str:
    night = F.loot_nights()['nights'][0]
    table_items = {i['item_id']: i for i in F.loot_table()['items']}
    drops = []
    for d in night['drops']:
        item = table_items[d['item_id']]
        users = len(item['candidates']) if d['kind'] == 'gear' else len(item['eligible'])
        tail = f'Awarded to {escape(d["award"]["name"])}' if d['award'] else f'{users} candidates'
        drops.append(item_row(D.icon_uri(d['icon']), d['name'], QUALITY_NAME[d['quality']], tail,
                              tag('Open', 'gold') if d['state'] == 'open' else tag('Awarded', 'kill')))
    drop = F.council_drop()
    cands = []
    for c in drop['candidates']:
        won = f'{c["items_won"]} item{"" if c["items_won"] == 1 else "s"} this tier'
        sub = f'+{c["gain_dps"]} DPS over {escape(c["replaces"])} &middot; {c["attendance"]["present"]} of {c["attendance"]["nights"]} nights &middot; {won}'
        foot = f'{tag(c["ms_os"].upper(), "gold")} {tag("BiS for them", "gold") if c["bis_for_them"] else ""} {votes(c["votes"]["yes"], c["votes"]["no"], c["votes"]["voters"])}'
        cands.append(('', player_cell(c['class'], c['name'], '', 30), f'{sub}<br>{foot}', '<a class="btn" href="#">Award</a>'))
    record = ''.join(item_row(D.icon_uri(d['icon']), f'{d["item_name"]} to {d["name"]}', QUALITY_NAME[d['quality']],
                              f'{d["votes_yes"]} yes, {d["votes_no"]} no of {d["votes_officers"]} &middot; by {escape(d["awarded_by"])}',
                              G.day_label(G.guild_time(d['at']))) for d in F.decisions())
    open_n = sum(1 for d in night['drops'] if d['state'] == 'open')
    main = panel(f'{O.BOSS} drops', ''.join(drops), meta=f'{G.day_label(G.guild_time(night["at"]))} &middot; {open_n} open')
    rail = panel(f'Candidates for {drop["name"]}', row_list(cands), framed=True) + panel('Decision record', record, action='<a href="#">Export &rsaquo;</a>')
    body = band('Loot') + f'<main class="w gd">{raid_strip(O.OFFICER)}<div class="twocol" style="grid-template-columns:minmax(0,1fr) minmax(0,1.15fr)"><div>{main}</div><aside>{rail}</aside></div></main>'
    return guild_page('loot council', body)


# ------------------------------------------------------------------------------------------------ Recruiting
def applications() -> str:
    apps = []
    for _i, name, cls, spec, il, parse, msg, avail, y, n, _mine, _sent in gd.APPLICANTS:
        sub = f'{spec} {cls.title()} &middot; item level {il} &middot; top {parse[0].upper()} {round(parse[1])}'
        line = escape(msg) + (f' &middot; can raid {escape(avail)}' if avail else '')
        apps.append(('', player_cell(cls, name, sub, 34), f'{line}<br>{votes(y, n, gd.VOTERS)}',
                     '<a class="btn" href="#">Accept</a> <a class="btn2" href="#">Decline</a>'))
    card = facts([('Looking for', escape(G._wanted_summary())), ('Raid nights', escape(gd.RECRUIT_NIGHTS))])
    decided = row_list([('', f'{escape(name)} accepted', f'{y} yes of {gd.VOTERS} &middot; by {escape(by)}', '') for _i, name, _c, _s, y, by, _r, _d in gd.DECIDED])
    main = panel('Open applications', row_list(apps), framed=True)
    rail = panel('Your public card', card, action='<a href="#">See it as a visitor &rsaquo;</a>') + panel('Recent decisions', decided)
    body = band('Roster', meta='Applications') + f'<main class="w gd"><div class="twocol"><div>{main}</div><aside>{rail}</aside></div></main>'
    return guild_page('applications', body)


def recruitment() -> str:
    wanted = row_list([(player_cell(c, n, '', 30), '', escape(note) if note else '', '') for c, _k, n, note in gd.RECRUIT_WANTED])
    up = ', '.join(f'{G._day(at)}, {G._clock(at)}' for at in gd.RECRUIT_UPCOMING)
    card = facts([('Region', f'{gd.REGION.upper()} PvP'), ('Faction', gd.FACTION.title()), ('Members', str(gd.RAIDERS)),
                  ('Raids', escape(gd.RECRUIT_NIGHTS)), ('Next nights', escape(up)), ('Progress', '1 of 1 named bosses down')])
    officers = row_list([(crest_row(nm, f'{G._spec_class(nm)} &middot; Officer'), '', '', '') for nm in gd.OFFICERS])
    main = (panel('Looking for', wanted, framed=True)
            + panel('Apply', f'<p>Sign in to apply to {escape(gd.GUILD)} with one of your characters.</p><p style="margin-top:12px"><a class="btn" href="#">Sign in with Battle.net</a></p>'))
    rail = panel(gd.GUILD, card) + panel('Why this guild', f'<div class="prose" style="font-size:14.5px">{escape(gd.RECRUIT_PITCH)}</div>') + panel('Officers', officers)
    bandh = titlebar(None, f'Join {gd.GUILD}', f'{tag("Recruiting", "kill")}', media=emb(), top=header('Thandrel', 'warrior', 60, 'Guild'))
    body = bandh + f'<main class="w gd"><div class="twocol"><div>{main}</div><aside>{rail}</aside></div></main>'
    return guild_page('recruitment', body)


# ------------------------------------------------------------------------------------------------ Settings
def settings() -> str:
    """The officer's Settings tab with the Forever Sixty bot installed: the plan and its seats, Discord, members and
    access, records and crest. Data: guild_settings_fixtures (the plan, seats, bot, invite, officer rank)."""
    import guild_settings_fixtures as SF
    from ws_guild_settings import day, renews_line
    view = SF.billing()
    seats = []
    for seat in view['seats']:
        h = seat['holder']
        if h is None:
            seats.append(('', '<span class="muted">Not given</span>', '', '<a class="btn2" href="#">Give seat</a>'))
        else:
            seats.append(('', player_cell(h['class'], h['name'], f'Given {day(seat["given_at"])}', 30), f'Can move from {day(seat["movable_at"])}', ''))
    given = sum(1 for s in view['seats'] if s['holder'] is not None)
    plan = (f'{row_list([("", renews_line(), "Billed to you", tag("Founder", "gold"))])}'
            f'<p style="margin:10px 0 4px"><a class="btn2" href="#">Manage billing</a> &nbsp;<a href="#">Transfer the plan &rsaquo;</a></p>'
            f'<p class="grp"><b>Premium seats</b> {given} of {len(view["seats"])} given</p>{row_list(seats)}')
    b = SF.settings_bot()['discord']['bot']
    discord = (f'<p>Raids post to <b class="strong">#{escape(b["channel_name"])}</b> in <b class="strong">{escape(b["server_name"])}</b>, with sign-up buttons.</p>'
               f'<p class="muted" style="font-size:13.5px;margin-top:6px">{b["linked_members"]} of {b["members"]} members linked</p>'
               f'<p style="margin-top:12px"><a class="btn2" href="#">Send test message</a> <a class="btn2" href="#">Change channel</a> <a href="#" class="muted" style="margin-left:10px">Remove</a></p>')
    s = SF.settings()
    holder = s['claimed_by']
    members = facts([('Claimed by', f'<span style="color:{D.CLASS_COLOR[holder["class"]]}">{escape(holder["name"])}</span>'),
                     ('Officer ranks', f'0 to {s["officer_max_rank_index"]}'), ('Report visibility', 'Guild only')])
    invite = (f'<p class="grp"><b>Invite link</b></p><div class="fld"><input type="text" readonly value="https://foreversixty.gg{SF.INVITE_URL}" aria-label="Invite link">'
              f'<a class="btn2" href="#">Copy</a></div><p class="muted" style="font-size:13px;margin-top:6px">Joins as a verified member &middot; '
              f'made {day(SF.INVITE_MADE)} &middot; <a href="#">New link</a></p>')
    records = row_list([('', '<a href="#" style="color:#fff">Attendance</a>', 'CSV', '&rsaquo;'), ('', '<a href="#" style="color:#fff">Loot</a>', 'CSV', '&rsaquo;'),
                        (f'<img src="{emblem(gd.FACTION)}" width="36" height="36" alt="">', 'Guild crest', 'Your faction&rsquo;s logo', '<a class="btn2" href="#">Choose an image</a>')])
    main = panel('Guild plan', plan, meta='Raid guild', framed=True) + panel('Discord', discord, action=tag('Connected', 'kill'))
    rail = panel('Members and access', members + invite) + panel('Records and crest', records)
    body = band('Settings') + f'<main class="w gd"><div class="twocol"><div>{main}</div><aside>{rail}</aside></div></main>'
    return guild_page('guild settings', body)


def build() -> None:
    write('b3-guild-settings', settings())
    write('b3-guild-overview', overview())
    write('b3-guild-roster', roster())
    write('b3-guild-raid', raid_planner())
    write('b3-guild-progression', progression())
    write('b3-guild-readiness', readiness())
    write('b3-guild-loot', loot())
    write('b3-guild-applications', applications())
    write('b3-guild-recruit', recruitment())


if __name__ == '__main__':
    build()
