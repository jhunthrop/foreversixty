"""Mock batch 3: the guild in the directions system (plan 2026-11-09-redesign, "Every player-facing surface").

Data, never invented: the guild mocks' own data modules in design/mocks (guild_data, guild_overview_data and
guild_overview_fixtures, guild_tabs_fixtures, guild_raid, guild_settings_fixtures), the values the guild parity specs serve
the build: Olympus XXVII, 24 raiders, Onyxia's Lair on Thu 10 Dec. Helpers that only format that data (times, the sign-up
lines) come from design/mocks/ws_guild2.py. Every element is drawn by prims.py.

Review rulings (batch 3): one band line for every tab; the next-raid strip on the Roster only; one primary button (teal)
and one secondary (blue), declines as quiet links; sign-up rows show their state, the viewer's own row its answers;
readiness lists the raiders short of something first and folds the rest; the cover lists its gaps first.
"""
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
from prims import (PRIMS_CSS, answers, facts, field, item_row, meter, need_grid, panel, player_cell, row_list, table, tag,  # noqa: E402
                   votes)
from style import footer, header, page, titlebar  # noqa: E402

OFFICER_VIEW = (O.OFFICER, 'hunter', 60)
SIGNUP = ['In', 'Late', 'Bench', 'Out']
STATE_TONE = {'in': 'kill', 'late': 'gold', 'bench': '', 'out': 'wipe'}
QUALITY_NAME = {0: 'poor', 1: 'common', 2: 'uncommon', 3: 'rare', 4: 'epic', 5: 'legendary'}
COVER_SHOWN = 6  # the cover's gaps and its thinnest lines; the rest behind "Show all"

GUILD_CSS = PRIMS_CSS + '''
.gd{padding-top:4px}.strip{display:flex;flex-wrap:wrap;align-items:center;gap:10px 18px;padding:12px 18px;margin-top:22px}
.strip p{font-size:14.5px}.strip .end{margin-left:auto;display:flex;gap:14px;align-items:center}
.emb{width:52px;height:52px;flex:none}.grp{font:600 12.5px var(--body);color:var(--muted);padding:12px 0 2px}
.grp b{color:#fff}.figs{display:flex;gap:28px;flex-wrap:wrap}.figs div{display:flex;flex-direction:column}.figs b{font:600 26px var(--body);color:#fff}
.figs span{font-size:12.5px;color:var(--muted)}.key{font-size:12.5px;color:var(--muted);padding:8px 20px 12px}
.quiet{color:var(--muted);font-weight:400;margin-left:12px}.warn{color:var(--gold)}
.lbl{display:block;font-size:12.5px;color:var(--muted);margin:14px 0 6px}
@media(max-width:700px){.strip .end{margin-left:0}}
'''


def emb() -> str:
    return f'<img class="emb" src="{emblem(gd.FACTION)}" alt="">'


def band(active: str, viewer: tuple = OFFICER_VIEW, title: str = gd.GUILD, meta: str = '', sub: list[tuple[str, bool]] | None = None) -> str:
    """One band for every tab: the faction emblem, the name, one short line, the tabs (and a sub-view row where a tab has one)."""
    line = meta or f'{gd.REGION.upper()} PvP &middot; {gd.RAIDERS} raiders'
    ctl = [[(t, t == active) for t in gd.TABS]] + ([sub] if sub else [])
    return titlebar(None, title, line, media=emb(), ctl=ctl, top=header(*viewer, 'Guild'))


def raid_strip(viewer: str) -> str:
    """The next raid and the viewer's answer, on the Roster (the tab a raider opens before a raid night)."""
    nxt = F.next_raid(viewer)
    at = G.guild_time(O.NEXT_RAID_AT)
    state = G.ch(viewer)[5]
    n = nxt['counts']['in'] + nxt['counts']['late']
    line = (f'<p><b class="strong">{escape(O.ZONE_NAME)}</b> &middot; {at:%a} {at.day} {at:%b}, {G.clock_label(at)} &middot; '
            f'<span class="num">{n}</span> of {nxt["size"]} signed up</p>')
    return f'<div class="card strip">{line}<span class="end">{answers(SIGNUP, state)}<a href="#">Open the raid &rsaquo;</a></span></div>'


def quiet(label: str) -> str:
    return f'<a class="quiet" href="#">{escape(label)}</a>'


def guild_page(title: str, body: str) -> str:
    return page(f'Forever Sixty, {title}', GUILD_CSS, body + footer())


def twocol(main: str, rail: str) -> str:
    return f'<div class="twocol"><div>{main}</div><aside>{rail}</aside></div>'


# ------------------------------------------------------------------------------------------------ shared grids
def comp_grid(counts: dict, target: dict) -> str:
    """The one comp grid (overview and planner): each role's target, who is In, and how many short."""
    late = counts['late_groups']
    rows = []
    for role in ('tanks', 'healers', 'melee', 'ranged'):
        have, need = counts[role], target[role]
        tone = 'ok' if have >= need else ('warn' if have + late[role] >= need else 'miss')
        short = need - have
        rows.append((role.title(), f'{late[role]} late' if late[role] else '',
                     [('none', str(need)), (tone, str(have)), (tone if short > 0 else 'none', str(short) if short > 0 else '&mdash;')]))
    return need_grid(['Need', 'In', 'Short'], rows)


def cover_grid(rows_src: list[dict], limit: int | None = COVER_SHOWN) -> str:
    """Buff and debuff cover, the gaps first: what the raid needs against who is signed up to bring it."""
    rows = []
    for r in rows_src:
        signed = r['providers_in'] + r['providers_late']
        tone = 'miss' if signed == 0 else ('ok' if signed - r['likely'] >= r['need'] else 'warn')
        rows.append(({'miss': 0, 'warn': 1, 'ok': 2}[tone], signed,
                     (f'<img class="ai" src="{D.icon_uri(r["icon"])}" width="22" height="22" alt=""> {escape(r["name"])}', '',
                      [('none', str(r['need'])), (tone, str(signed))])))
    rows.sort(key=lambda x: (x[0], x[1]))
    shown = [r[2] for r in rows][:limit]
    return need_grid(['Need', 'In'], shown)


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
    at = G.guild_time(O.NEXT_RAID_AT)
    figs = (f'<div class="figs"><div><b class="num" style="color:var(--gain)">{c["in"]}</b><span>In</span></div><div><b class="num" style="color:var(--gold)">{c["late"]}</b><span>Late</span></div>'
            f'<div><b class="num">{c["bench"]}</b><span>Bench</span></div><div><b class="num" style="color:#f0736b">{len(view["not_signed_raiders"])}</b><span>Not signed</span></div></div>')
    raid = (row_list([(f'<img class="ai" src="{D.icon_uri(G.RAID_ICON)}" width="40" height="40" alt="">', escape(O.ZONE_NAME),
                       f'{G.day_label(at)} &middot; {G.clock_label(at)}', f'<b class="gold">{G.countdown_label()}</b>')])
            + f'<div style="margin-top:12px">{figs}</div>')
    cover_rows = F.cover_rows()
    last = F.recent_attendance()
    main = (panel('Action needed', act, framed=True) + panel('Next raid', raid, action='<a href="#">Open the planner &rsaquo;</a>')
            + panel('Comp against the target', comp_grid(F.raid_view(O.MEMBER, False)['counts'], O.COMP_TARGET), action='<a href="#">Edit target &rsaquo;</a>', flush=True))
    rail = (panel('Buff and debuff cover', cover_grid(cover_rows), action=f'<a href="#">All {len(cover_rows)} &rsaquo;</a>', flush=True)
            + panel(f'Attendance, last {len(last["raids"])} raids', attendance_grid(last), action='<a href="#">Roster &rsaquo;</a>', flush=True))
    body = band('Overview') + f'<main class="w gd">{twocol(main, rail)}</main>'
    return guild_page('guild', body)


def attendance_grid(page_data: dict) -> str:
    mark = {'present': ('ok', '&#10003;'), 'absent': ('miss', '&times;'), 'bench': ('warn', 'B'), 'none': ('none', '&mdash;')}
    cols = [f"{G.guild_time(r['starts_at']):%a} {G.guild_time(r['starts_at']).day}" for r in page_data['raids']]
    rows = [(f'<span style="color:{D.CLASS_COLOR.get(G.ch(r["name"])[1], "#fff")}">{escape(r["name"])}</span>', 'slipping' if r['slipping'] else '',
             [mark[x] for x in r['cells']]) for r in sorted(page_data['rows'], key=lambda r: (not r['slipping'], r['name']))]
    return need_grid(cols, rows) + '<p class="key">&#10003; came &middot; &times; missed &middot; B benched</p>'


# ------------------------------------------------------------------------------------------------ Roster
def roster_rows(home: dict) -> list[dict]:
    return sorted(home['roster'], key=lambda r: (r['name'] != O.OFFICER, -r['item_level']))


def roster() -> str:
    home = T.roster_home()
    waiting = row_list([(player_cell(r['class'], r['name'], f'{r["spec"]} {r["class"].title()}', 32), '', '',
                         f'<a class="btn2" href="#">Approve</a>{quiet("Decline")}') for r in home['pending']])
    rows = []
    for r in roster_rows(home):
        rank = ' &middot; Officer' if r['rank'] == 'officer' else ''
        rows.append([player_cell(r['class'], r['name'] + (' (you)' if r['name'] == O.OFFICER else ''), f'{r["spec"]} {r["class"].title()}{rank}', 28),
                     str(r['item_level']), f'{r["attendance"]["present"]} of {r["attendance"]["nights"]}'])
    tbl = table([('Raider', False), ('Item level', True), ('Raids', True)], rows,
                row_cls=['me' if r['name'] == O.OFFICER else '' for r in roster_rows(home)])
    main = (panel('Waiting for approval', waiting, action='<a class="btn" href="#">Approve all</a>', framed=True) + panel('Roster', tbl, flush=True))
    body = band('Roster', sub=[('Members', True), ('Applications', False)]) + f'<main class="w gd">{raid_strip(O.OFFICER)}{main}</main>'
    return guild_page('guild roster', body)


# ------------------------------------------------------------------------------------------------ Raid planner
def raid_planner() -> str:
    raid = guild_raid.raid_view()
    n = raid['counts']
    by = {m['name']: m for m in guild_raid.MEMBERS}
    you = by[guild_raid.VIEWER]
    order = {'tank': 0, 'healer': 1, 'dps': 2}
    rows = sorted((guild_raid.member(s['name']) for s in raid['signups']), key=lambda m: (order[m['role']], G.CLASS_ORDER.index(m['class']), m['name']))

    def row(m: dict) -> tuple:
        mine = m['name'] == you['name']
        trail = answers(SIGNUP, m['state']) if mine else tag(m['state'].title(), STATE_TONE[m['state']])
        return (player_cell(m['class'], m['name'] + (' (you)' if mine else ''), f'{m["spec"]} {m["class"].title()}', 28), '', '', trail)

    groups = ''
    for label, pick in [('Tanks', lambda m: m['state'] == 'in' and m['group'] == 'tank'), ('Healers', lambda m: m['state'] == 'in' and m['group'] == 'healer'),
                        ('Melee', lambda m: m['state'] == 'in' and m['group'] == 'melee'), ('Ranged', lambda m: m['state'] == 'in' and m['group'] == 'ranged'),
                        ('Late', lambda m: m['state'] == 'late'), ('Bench', lambda m: m['state'] == 'bench')]:
        ms = [m for m in rows if pick(m)]
        groups += f'<p class="grp"><b>{label}</b> {len(ms)}</p>' + row_list([row(m) for m in ms])
    enc = raid['assignments']['encounters'][0]
    assign = row_list([('', escape(s['label']),
                        f'{len(s["names"])} of {s["count"]} &middot; {escape(s["purpose"])}<br>'
                        + (' '.join(tag(x['name'], 'gold' if x['own'] else '') for x in s['names']) or '<span class="warn">Open</span>'), '') for s in enc['slots']])
    signed = n['in'] + n['late']
    main = panel('Sign-ups', groups, meta=f'<span class="num">{signed}</span> of {raid["size"]}', framed=True)
    rail = (panel('Comp against the target', comp_grid(n, raid['comp_target']), flush=True) + panel(f'{enc["name"]} assignments', assign)
            + panel('Notes', f'<div class="prose" style="font-size:14.5px">{raid["notes_html"]}</div>'))
    at = G.guild_time(raid['starts_at'])
    meta = f'{at:%a} {at.day} {at:%b}, {G.clock_label(at)} &middot; <span class="num">{signed}</span> of {raid["size"]} signed up'
    body = band('Raids', title=raid['title'], meta=meta) + f'<main class="w gd">{twocol(main, rail)}</main>'
    return guild_page('raid planner', body)


# ------------------------------------------------------------------------------------------------ Progression and readiness
def progression() -> str:
    """One panel: each named boss with its nights, pulls, wipes and kill; the tier's count in the panel line."""
    pg = F.progression()
    t = pg['tier']
    rows = []
    for e in (x for x in pg['encounters'] if x['pulls'] > 0):
        nights = e['pulls_by_night']
        kills = sum(1 for nt in nights if nt['killed'])
        first = next((nt for nt in nights if nt['killed']), None)
        rows.append((f'<img class="ai" src="{D.icon_uri(G.RAID_ICON)}" width="36" height="36" alt="">', escape(e['name']),
                     f'{e["pulls"]} pulls &middot; {e["pulls"] - kills} wipes &middot; {e["deaths_per_pull"]:.1f} deaths a pull'
                     + (f' &middot; first kill {G.day_label(G.guild_time(first["date"] + "T20:00:00Z"))}' if first else ''),
                     tag('Killed', 'kill') if kills else tag('Progress', 'gold')))
    bar = f'<div style="margin:6px 0 4px">{meter(t["down"] / t["named_encounters"], "var(--gain)")}</div>'
    body = band('Progression') + (f'<main class="w gd">{panel(escape(O.ZONE_NAME), bar + row_list(rows), meta=f"{t["down"]} of {t["named_encounters"]} bosses down", framed=True)}</main>')
    return guild_page('guild progression', body)


def short_of(r: dict) -> list[str]:
    out = []
    if r['gear_gap']['upgrades']:
        out.append(f'{r["gear_gap"]["upgrades"]} upgrades')
    if r['enchants']['missing_slots']:
        out.append(f'{len(r["enchants"]["missing_slots"])} enchants missing')
    if r['consumables']['state'] != 'stocked':
        out.append('consumables short')
    if r['talent_points_unspent']:
        out.append(f'{r["talent_points_unspent"]} talent points unspent')
    return out


def readiness() -> str:
    rows = F.readiness(True)['rows']
    short = [r for r in rows if short_of(r)]
    ready = [r for r in rows if not short_of(r)]
    listed = row_list([('', player_cell(r['class'], r['name'], f'{r["spec"]} {r["class"].title()} &middot; item level {r["item_level"]}', 30),
                        '<span class="warn">' + ' &middot; '.join(short_of(r)) + '</span>', '<a class="btn2" href="#">Nudge</a>') for r in short])
    names = ', '.join(escape(r['name']) for r in ready)
    nxt = G.guild_time(O.NEXT_RAID_AT)
    body = (band('Readiness') + '<main class="w gd">'
            + panel('Short of something', listed or '<p class="empty">Nobody.</p>', meta=f'In or Late for {G.day_label(nxt)}', framed=True)
            + panel(f'{len(ready)} ready', f'<p style="font-size:14.5px">{names}</p>') + '</main>')
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
        row = item_row(D.icon_uri(d['icon']), d['name'], QUALITY_NAME[d['quality']], tail,
                       tag('Open', 'gold') if d['state'] == 'open' else tag('Awarded', 'kill'))
        if d['item_id'] == O.SELECTED_DROP:  # the drop the candidates panel is about
            row = row.replace('<div class="ir">', '<div class="ir" style="background:rgba(14,134,202,.14);box-shadow:-10px 0 0 rgba(14,134,202,.14),10px 0 0 rgba(14,134,202,.14)">', 1)
        drops.append(row)
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
    body = (band('Loot') + '<main class="w gd"><div class="twocol" style="grid-template-columns:minmax(0,1fr) minmax(0,1.15fr)">'
            f'<div>{main}</div><aside>{rail}</aside></div></main>')
    return guild_page('loot council', body)


# ------------------------------------------------------------------------------------------------ Recruiting
def applications() -> str:
    apps = []
    for _i, name, cls, spec, il, parse, msg, avail, y, n, _mine, _sent in gd.APPLICANTS:
        sub = f'{spec} {cls.title()} &middot; item level {il} &middot; top {parse[0].upper()} {round(parse[1])}'
        line = escape(msg) + (f' &middot; can raid {escape(avail)}' if avail else '')
        apps.append(('', player_cell(cls, name, sub, 34), f'{line}<br>{votes(y, n, gd.VOTERS)}',
                     f'<a class="btn" href="#">Accept</a>{quiet("Decline")}'))
    card = facts([('Looking for', escape(G._wanted_summary())), ('Raid nights', escape(gd.RECRUIT_NIGHTS))])
    decided = row_list([('', f'{escape(name)} accepted', f'{y} yes of {gd.VOTERS} &middot; by {escape(by)}', '') for _i, name, _c, _s, y, by, _r, _d in gd.DECIDED])
    main = panel('Open applications', row_list(apps), framed=True)
    rail = panel('Your public card', card, action='<a href="#">See it as a visitor &rsaquo;</a>') + panel('Recent decisions', decided)
    body = band('Roster', sub=[('Members', False), ('Applications', True)]) + f'<main class="w gd">{twocol(main, rail)}</main>'
    return guild_page('applications', body)


def recruitment() -> str:
    """The public card, signed out: a visitor's view, Sign in to apply."""
    wanted = row_list([(player_cell(c, n, '', 30), '', escape(note) if note else '', '') for c, _k, n, note in gd.RECRUIT_WANTED])
    up = ', '.join(f'{G._day(at)}, {G._clock(at)}' for at in gd.RECRUIT_UPCOMING)
    card = facts([('Region', f'{gd.REGION.upper()} PvP'), ('Faction', gd.FACTION.title()), ('Members', str(gd.RAIDERS)),
                  ('Raids', escape(gd.RECRUIT_NIGHTS)), ('Next nights', escape(up))])
    officers = row_list([(player_cell(G.ch(nm)[1], nm, G._spec_class(nm), 30), '', '', '') for nm in gd.OFFICERS])
    main = panel('Looking for', wanted, framed=True) + panel('Apply', '<a class="btn" href="#">Sign in with Battle.net to apply</a>')
    rail = panel(gd.GUILD, card) + panel('Why this guild', f'<div class="prose" style="font-size:14.5px">{escape(gd.RECRUIT_PITCH)}</div>') + panel('Officers', officers)
    bandh = titlebar(None, f'Join {gd.GUILD}', tag('Recruiting', 'kill'), media=emb(), top=header('', 'warrior', None, 'Guild'))
    body = bandh + f'<main class="w gd">{twocol(main, rail)}</main>'
    return guild_page('recruitment', body)


# ------------------------------------------------------------------------------------------------ Settings
def settings() -> str:
    """The officer's Settings tab with the Forever Sixty bot installed. Data: guild_settings_fixtures."""
    import guild_settings_fixtures as SF
    from ws_guild_settings import day, renews_line
    view = SF.billing()
    seats = []
    for seat in view['seats']:
        h = seat['holder']
        if h is None:
            seats.append(('', '<span class="muted">Open seat</span>', '', '<a class="btn2" href="#">Give seat</a>'))
        else:
            seats.append(('', player_cell(h['class'], h['name'], f'Given {day(seat["given_at"])} &middot; can move from {day(seat["movable_at"])}', 30), '', ''))
    given = sum(1 for s in view['seats'] if s['holder'] is not None)
    plan = (row_list([('', renews_line(), '', tag('Founder', 'gold'))])
            + f'<p style="margin:10px 0 4px"><a class="btn2" href="#">Manage billing</a>{quiet("Transfer the plan")}</p>'
            + f'<p class="grp"><b>Premium seats</b> {given} of {len(view["seats"])} given</p>{row_list(seats)}')
    b = SF.settings_bot()['discord']['bot']
    discord = (f'<p>Raids post to <b class="strong">#{escape(b["channel_name"])}</b> in <b class="strong">{escape(b["server_name"])}</b>, with sign-up buttons.</p>'
               f'<p class="muted" style="font-size:13.5px;margin-top:6px">{b["linked_members"]} of {b["members"]} members linked</p>'
               f'<p style="margin-top:12px"><a class="btn2" href="#">Send test message</a> <a class="btn2" href="#">Change channel</a>{quiet("Remove")}</p>')
    s = SF.settings()
    holder = s['claimed_by']
    ranks = ' and '.join(['Guild Master'] + [f'rank {i}' for i in range(1, s['officer_max_rank_index'] + 1)])
    members = facts([('Claimed by', f'<span style="color:{D.CLASS_COLOR[holder["class"]]}">{escape(holder["name"])}</span>'),
                     ('Officers', ranks), ('Report visibility', 'Guild only')])
    invite = (f'<span class="lbl">Invite link</span>{field(f"foreversixty.gg{SF.INVITE_URL}", "<a class=btn2 href=#>Copy</a>", "Invite link")}'
              f'<p class="muted" style="font-size:13px;margin-top:6px">Made {day(SF.INVITE_MADE)} {quiet("New link")}</p>')
    crest = (f'<span class="lbl">Guild crest</span><div style="display:flex;align-items:center;gap:12px">'
             f'<img src="{emblem(gd.FACTION)}" width="40" height="40" alt=""><a class="btn2" href="#">Choose an image</a></div>')
    records = row_list([('', '<a href="#" style="color:#fff">Attendance</a>', 'CSV', '&rsaquo;'), ('', '<a href="#" style="color:#fff">Loot</a>', 'CSV', '&rsaquo;')]) + crest
    main = panel('Guild plan', plan, meta='Raid guild', framed=True) + panel('Discord', discord, action=tag('Connected', 'kill'))
    rail = panel('Members and access', members + invite) + panel('Records and crest', records)
    body = band('Settings') + f'<main class="w gd">{twocol(main, rail)}</main>'
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
