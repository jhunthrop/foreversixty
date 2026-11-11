"""The Guild tab (design/specs/2026-11-10-guild-tab.md): the Raid tab became Guild, so a player below 60 still has a tab
that is about his guild. Four states: below 60 in a guild, at 60 (the raid first, unchanged), officer, no guild. Composed
from the Raid tab's round-5 parts in ws_app_flows.py; ws_app.py registers the screens.

The world is ws_app_flows's (Olympus XXVII, Onyxia's Lair Wed 6 Jan, 8:00 pm CST), with Felsnap at 36 in the guild for the
below-60 state. Every figure is one the guild home already sends (members[].level and seen_at, reports[], summary.tier,
counts, viewer.officer_tools) or the setup checklist's done and total.
"""
from __future__ import annotations

from mocklib import CLASS_COLOR
import ws_app_flows as F
from ws_app_frame import app_page
from ws_prims import Emblem, build_page, crest_img, icon

GUILD = 'Olympus XXVII'
MEMBERS, AT_CAP, CLOSE = 38, 31, 4
LEVELER_LEVEL = 36

# Recently on: the newest seen_at first (members[]: name, class, level, spec_name, seen).
RECENT = [('Grunvald', 'warrior', 60, 'Protection Warrior', 'Seen just now'), ('Mirethil', 'priest', 60, 'Holy Priest', 'Seen 12 min ago'),
          ('Kaelra', 'warrior', 60, 'Fury Warrior', 'Seen 40 min ago'), ('Thornwick', 'druid', 57, 'Feral Druid', 'Seen 1 h ago'),
          ('Selka', 'druid', 60, 'Restoration Druid', 'Seen 2 h ago')]
CLOSE_ROWS = [('Thornwick', 'druid', 57, 'Feral Druid'), ('Vexa', 'warlock', 56, 'Affliction Warlock'),
              ('Pell', 'mage', 55, 'Frost Mage'), ('Orrin', 'paladin', 55, 'Holy Paladin')]
# reports[]: title, zone, kills of fights, duration, day.
LOGS = [('UBRS, Sunday', 'Blackrock Spire', 7, 8, '1 h 32 min', 'Sun 3 Jan'),
        ('Scholomance, Saturday', 'Scholomance', 9, 9, '1 h 08 min', 'Sat 2 Jan'),
        ('Onyxia, Wednesday', 'Onyxia&rsquo;s Lair', 1, 1, '41 min', 'Wed 30 Dec')]
SETUP_DONE, SETUP_TOTAL, APPLICANTS = 4, 8, 1

GUILD_CSS = '''
.r5raid .gh{display:flex;align-items:center;gap:12px;padding:8px 0}
.r5raid .gh .crest{flex:none}
.r5raid .gh h1{margin:0;font:400 22px/26px var(--font-title);color:#f3eee2}
.r5raid .gh .s{white-space:normal}
.r5raid .spread{display:flex;gap:0;margin:4px 0 8px;border-top:1px solid var(--line);padding-top:10px}
.r5raid .spread>div{flex:1;min-width:0}
.r5raid .spread b{font:700 20px/24px var(--font-body);font-variant-numeric:tabular-nums}
.r5raid .spread .s{font-size:12px;line-height:16px}
.r5raid .spread .you b{color:var(--gold)}
.r5raid .seg.off .btn{color:var(--muted);opacity:.45}
.r5raid .gate{margin:0;font:400 13px/20px var(--font-body);color:var(--muted)}
.ap.r5raid>*+.gate{margin-top:20px}
.r5raid .row .in .crest{flex:none}
.r5raid .row .in>div{flex:1;min-width:0}
.r5raid .row .in .when{flex:none;font:400 13px/18.29px var(--font-body);color:var(--muted);text-align:right}
.r5raid .row .in .lv{flex:none;font:700 15px/20.57px var(--font-body);font-variant-numeric:tabular-nums;color:#fff}
.r5raid .grp{margin:0;padding:10px 0 2px;font:600 13px/18px var(--font-body);color:var(--muted)}
.r5raid .pn>.grp:first-child{padding-top:6px}
.r5raid .row .in .cnt{flex:none;font:600 14px/20px var(--font-body);color:var(--gold)}
.r5raid .kill{color:#7fd48a}
[data-region^=section-guilds-recruiting] .sh{height:auto;min-height:26px}
.r5raid .row .tr .when{font:600 13px/18.29px var(--font-body);color:var(--muted)}
'''


def _crest_row(name: str, cls: str, sub: str, trail: str) -> str:
    lead = crest_img(cls, 32)
    return (f'<div class="row"><div class="in">{lead}<div><span class="t" style="color:{CLASS_COLOR[cls]}">{name}</span>'
            f'<span class="s">{sub}</span></div>{trail}</div></div>')


def _head(leveler: bool) -> str:
    """The guild head: the guild's crest (the faction mark, it has none), the name, members and faction; below 60 the
    level spread under it."""
    head = (f'<div class="gh">{Emblem("alliance", 48)}<div><h1>{GUILD}</h1>'
            f'<span class="s">{MEMBERS} members &middot; Alliance &middot; US PvP</span></div></div>')
    if leveler:
        head += ('<div class="spread"><div><b>' + str(AT_CAP) + '</b><span class="s">at 60</span></div>'
                 f'<div><b>{CLOSE}</b><span class="s">55 to 59</span></div>'
                 f'<div class="you"><b>{LEVELER_LEVEL}</b><span class="s">you</span></div></div>')
    return f'<div class="blk-r5" data-region="section-guild"><div class="pn">{head}</div></div>'


def _raid_view(platform: str) -> str:
    """The next raid below 60 (coordinator ruling 2026-11-10: they sign up): the raid head as at 60, one quiet line under it,
    then Who is coming."""
    head = F._r5_head(platform, 'In').replace('Thandrel, Damage', 'Felsnap, Damage')
    gate = f'<p class="gate">Raids open at 60 &middot; you&rsquo;re {LEVELER_LEVEL}</p>'
    return head + gate + F._r5_coming(platform)


def _roster(leveler: bool) -> str:
    rows = '<p class="grp">Recently on</p>' + ''.join(
        _crest_row(n, c, f'{lv} &middot; {spec}', f'<span class="when">{seen}</span>') for n, c, lv, spec, seen in RECENT)
    if leveler:
        rows += '<p class="grp">Close to 60</p>' + ''.join(
            _crest_row(n, c, spec, f'<span class="lv">{lv}</span>') for n, c, lv, spec in CLOSE_ROWS)
    return F._r5_section('Roster', rows, action='Roster &rsaquo;', meta=f'{AT_CAP} of {MEMBERS} at 60')


def _logs() -> str:
    # the day sits on the right, as the roster's seen line does
    rows =''.join(f'<div class="row"><div class="in"><div><span class="t">{t}</span><span class="s">{zone} &middot; '
                   f'<span class="{"kill" if k == n else ""}">{k}/{n} kills</span> &middot; {dur}</span></div>'
                   f'<span class="when">{day}</span></div></div>' for t, zone, k, n, dur, day in LOGS)
    return F._r5_section('Recent logs', rows)


def _progression() -> str:
    """The tier's line (summary.tier: its name, down of named, pulls), opening the site's Progression tab."""
    return F._r5_section('Progression', F._r5_row('First tier', '<span class="kill">1 of 1 down</span> &middot; 6 pulls this tier'))


def _loot() -> str:
    return F._r5_section('Loot', F._r5_row('Last raid: 4 pieces'), action='Loot history &rsaquo;')


def _officer() -> str:
    def row(title: str, sub: str, count: str = '') -> str:
        c = f'<span class="cnt">{count}</span>' if count else ''
        return f'<div class="row"><div class="in"><div><span class="t">{title}</span><span class="s">{sub}</span></div>{c}</div></div>'
    rows = (row('Plan a raid', 'Onyxia Wed, Barrow Deeps Sat')
            + row('Applications', 'Mirelle, Holy Priest 60', f'{APPLICANTS} waiting')
            + row('Setup', 'Discord bot, raid schedule and addon sync to go', f'{SETUP_DONE} of {SETUP_TOTAL} done'))
    return F._r5_section('Officer', rows)


def _pages(leveler: bool) -> str:
    """The pages the sections above do not already open: Roster has its section's link, Loot history too at 60; Guild sims
    is a 60's page."""
    shown = {'Calendar', 'Loot history'} if leveler else {'Calendar', 'Guild sims'}
    pages = [p for p in F.R5_PAGES if p[0] in shown]
    return F._r5_section('Pages', ''.join(F._r5_row(t, sub) for t, sub in pages))


def _none(leveler: bool) -> str:
    applied = '<span class="when">Applied Mon</span>'
    apply = '<a class="btn" href="#">Apply</a>'
    spec = 'Beast Mastery Hunter' if leveler else 'Fury Warrior'
    rec = F._r5_section(f'Guilds recruiting a {spec}',
                        F._r5_row(GUILD, 'Alliance &middot; Wed and Sat, 8:00 pm CST &middot; needs damage', trail=applied)
                        + F._r5_row('Dawnbreak', 'Alliance &middot; Tue and Thu, 9:00 pm EST &middot; needs damage', trail=apply))
    find = F._r5_section('Find your guild', F._r5_row('Search guilds on the site', 'By name or realm')
                         + F._r5_row('Claim your guild', 'A guild master or an officer confirms with the addon or Battle.net.'))
    gate = f'<p class="gate">Raids open at 60 &middot; you&rsquo;re {LEVELER_LEVEL}</p>' if leveler else ''
    return gate + rec + find


# The raid's assignments (RaidView.assignments, Onyxia's template in api/internal/guilds/raid_assign_templates.go).
WHELPS_PURPOSE = 'Pick up the west cave&rsquo;s whelps in phases 2 and 3 and kill them away from the raid.'
ASSIGNED = [('Main tank', 1, [('Grunvald', 'warrior')]), ('Off tank (phase 3, fear)', 2, [('Brakka', 'druid')]),
            ('Main tank healers', 4, [('Mirethil', 'priest'), ('Selka', 'druid'), ('Hollowmend', 'priest'), ('Aenwyn', 'paladin')]),
            ('Fear cover', 2, [('Mirethil', 'priest')]),
            ('Whelps, west cave', 3, [('Thandrel', 'warrior'), ('Kaelra', 'warrior'), ('Thalrin', 'mage')])]


def _your_job() -> str:
    """Your job (spec 2.3a): the slot with its partners, the boss and the slot's purpose; the web's title."""
    return F._r5_section('Your job', F._r5_row('Whelps, west cave with Kaelra and Thalrin', f'Onyxia &middot; {WHELPS_PURPOSE}'))


def _assignments() -> str:
    """An officer's read-only Assignments (spec 2.3a): each filled slot, its names (yours in gold), how full; Open in planner."""
    def names(people: list) -> str:
        return ', '.join(f'<span style="color:{"var(--gold)" if n == "Thandrel" else CLASS_COLOR[c]}">{n}</span>' for n, c in people)
    rows = '<p class="grp">Onyxia assignments</p>' + ''.join(
        f'<div class="row"><div class="in"><div><span class="t">{label}</span><span class="s">{names(people)}</span></div>'
        f'<span class="when">{len(people)} of {count}</span></div></div>' for label, count, people in ASSIGNED)
    rows += '<p class="stamp" style="margin:0">Assigning is on the site for now.</p>'
    return F._r5_section('Assignments', rows, action='Open in planner &rsaquo;')


def _body(state: str, platform: str) -> str:
    if state == 'leveler':
        return _head(True) + _raid_view(platform) + _roster(True) + _logs() + _progression() + _pages(True)
    if state == 'none':
        return _none(False)
    if state == 'none-leveler':
        return _none(True)
    officer = state == 'officer'
    head = F._r5_head(platform, 'In')
    needs = (F._r5_section('Needs you', F._r5_row('5 have not answered, 2 of them healers', 'Healers short by 3',
                                                  trail='<a class="btn" href="#">Preview mention</a>')
                           + F._r5_row('1 applicant', 'Mirelle, Holy Priest 60, 2 min ago'), meta='6 open') if officer else '')
    return (head + _your_job() + needs + (_officer() + _assignments() if officer else '') + F._r5_coming(platform, officer=officer)
            + F._r5_ready(F.R5_READY)
            + _head(False) + _roster(False) + _loot() + _logs() + _progression() + _pages(False))


def guild(state: str):
    def b(platform: str) -> str:
        badges = {'member': F.ANSWERED, 'officer': F.OFFICER, 'leveler': {'activity': 1}, 'none': {'activity': 2},
                  'none-leveler': {'activity': 1}}[state]
        who = 'hunter' if state in ('leveler', 'none-leveler') else 'warrior'
        html = build_page('Guild', f'<div class="ap r5raid">{_body(state, platform)}</div><style>{F.FLOWS_CSS}{F.RAID_R5_CSS}{GUILD_CSS}</style>',
                          note='Native app mock (design/specs/2026-11-10-guild-tab.md).')
        return app_page(platform, html, 'guild', title='Guild', badges=badges, who=who)
    return b


SCREENS = {
    'guild-tab': guild('member'),
    'guild-tab-leveler': guild('leveler'),
    'guild-tab-officer': guild('officer'),
    'guild-tab-none': guild('none'),
    'guild-tab-none-leveler': guild('none-leveler'),
}
