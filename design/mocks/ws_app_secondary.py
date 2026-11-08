"""The secondary screens of the flows-first app (design/specs/2026-11-08-app-flows.md section 2): Guides, Patch notes,
Sims, the guild's Roster, Calendar and Raid history, and the account screens. Each restyles what the app already does,
in the flows primitives (ws_app_flows.py), and never grows into a copy of the site: where the site has more, one row at
the end opens it ("Full guide on the site"). The words are the app's (apps/shared/copy/en.json); the data is the app's
contract fixtures (docs/contracts/fixtures/app), on the flows scenario's clock (Fri 6 Nov 2026) and character (Thandrel).

Removed states, no longer a native state: guide-you (the reader carries no per-character answer; Your guide only opens the
same reader) and sim-saved (a saved sim is a Your sims row that opens the one result screen).
"""
from __future__ import annotations

import json
from datetime import datetime, timedelta, timezone
from functools import cache

from mocklib import CLASS_COLOR, ROOT
from ws_app_flows import DOLL, ONY, RAID, ROLES, WHEN, page, signup_buttons
from ws_app_frame import FACTION_MARK, app_page
from ws_prims import (Button, Disclosure, Figure, Figures, Meter, PageHeader, Row, Rows, Section, Skeleton, Switch, Tag, Tally,
                      Countdown, crest_img, icon)

FIXTURES = ROOT / 'docs/contracts/fixtures/app'
WARRIOR = CLASS_COLOR['warrior']
CHICAGO = timezone(timedelta(hours=-6))  # the guild's time zone in December (CST)
DOT = ' &middot; '
NO_ICON = '__none__'  # an icon id with no file: the bordered empty tile the app draws for a spell with no icon

SECONDARY_CSS = '''
.lbl.first{margin-top:12px}
.glance{margin-top:20px;padding:16px;border:1px solid var(--line-warm);border-radius:8px;background:var(--card-top,#131824)}
.glance .lbl{margin:0 0 6px}.glance .lbl+.t-lead{margin-bottom:14px}
.t-lead.role{font:700 18px/24px var(--font-body);color:var(--strong)}
.nchips{display:flex;flex-wrap:wrap;gap:8px;list-style:none;margin:0;padding:0}
.nchips li{display:inline-flex;align-items:center;gap:6px;min-height:32px;padding:0 10px;border:1px solid var(--line);border-radius:16px;
 font:600 14px/18px var(--font-body);color:var(--strong)}
.nchips b{font:500 12px/16px var(--font-mono);color:var(--gold)}
.prose{padding-top:8px;font-size:15px;line-height:22px;color:var(--text)}
.blist{margin:0;padding-left:18px;font-size:15px;line-height:22px;color:var(--text)}.blist li+li{margin-top:4px}
.why{margin-top:4px;color:var(--text)}
.mono{font:500 13px/18px var(--font-mono);color:var(--muted)}
.dateb{display:inline-flex;align-items:center;justify-content:center;min-width:56px;min-height:48px;padding:4px 6px;border:1px solid var(--line);
 border-radius:4px;font:500 13px/16px var(--font-mono);color:var(--strong);text-align:center}
.fmark{width:24px;height:24px}
.ggrid{display:grid;grid-template-columns:repeat(6,40px);gap:8px;padding-top:12px}
.answer{margin-top:20px}.answer .fig .v{font-size:52px;line-height:56px}
.state{display:block;margin-top:2px;font-size:13px;line-height:18px;color:var(--muted)}.state .on{color:var(--gold)}.state .bad{color:var(--wipe)}
.prog{margin-top:8px}
.unread .ttl::after{content:"";display:inline-block;width:6px;height:6px;margin-left:8px;border-radius:50%;background:var(--gold);vertical-align:middle}
.ticks{list-style:none;margin:0;padding:8px 0 0}.ticks li{display:flex;gap:8px;padding:6px 0;font-size:15px;line-height:22px;color:var(--text)}
.ticks li::before{content:"\\2713";color:var(--kill);font-weight:700}
.version{margin-top:16px;text-align:center}
.skelcol{display:flex;flex-direction:column;gap:10px}
'''


def screen(body: str, title: str) -> str:
    return page(body + f'<style>{SECONDARY_CSS}</style>', title)


@cache
def fixture(name: str) -> dict:
    return json.loads((FIXTURES / f'{name}.json').read_text())


def data(name: str) -> dict:
    """The `data` of an API envelope, or the whole file for a CDN file."""
    body = fixture(name)
    return body.get('data', body) if isinstance(body, dict) else body


def site_row(title: str, sub: str) -> str:
    """The one row that hands the rest to the site (the in-app browser sheet)."""
    return f'<div data-region="footer-note">{Rows([Row(f"{title} &rsaquo;", sub)], open_=True)}</div>'


def footer(text: str, extra: str = '') -> str:
    return f'<div data-region="footer-note"><p class="t-small" style="margin-top:16px">{text}</p>{extra}</div>'


def error_block(line: str, sub: str = '') -> str:
    row = Row(line, sub, trail=Button('Try again', small=True), link=False)
    return f'<div class="blk" data-region="section-error">{Rows([row], open_=True)}</div>'


def skeleton_block(blocks: list[tuple[int, int]], region: str = 'section-loading') -> str:
    inner = ''.join(f'<div style="margin-top:12px">{Skeleton(w, h)}</div>' for w, h in blocks)
    return f'<div class="blk" data-region="{region}">{inner}</div>'


def chip_strip(labels: list[str], current: str) -> str:
    return '<div class="chiprow" data-region="controls">' + ''.join(Button(x, primary=(x == current), small=True) for x in labels) + '</div>'


def numbered(items: list[str]) -> str:
    return '<ol class="nchips">' + ''.join(f'<li><b>{i}</b>{x}</li>' for i, x in enumerate(items, 1)) + '</ol>'


def bullets(label: str, lines: list[str]) -> str:
    return f'<p class="lbl">{label}</p><ul class="blist">' + ''.join(f'<li>{x}</li>' for x in lines) + '</ul>' if lines else ''


def short_date(iso: str) -> str:
    d = datetime.fromisoformat(iso[:10])
    return f'{d.day} {d:%b}'


def long_date(iso: str) -> str:
    d = datetime.fromisoformat(iso[:10])
    return f'{d.day} {d:%b} {d.year}'


# ------------------------------------------------------------------------------------------------ Guides (F9)
GUIDE = 'cdn.guides.warrior-protection'
SECTIONS = ['Overview', 'Talents', 'Rotation', 'Stats', 'Gear', 'Races']
ROLE = {'dps': 'DPS', 'healer': 'Healer', 'tank': 'Tank'}
SITE_GUIDE = 'Stat weights, the talent grid, enchants, professions, leveling and what changed from 1.12'


def spec_title(g: dict) -> str:
    return f'{g["spec_name"]} {g["class_name"]}'


def guide_head(g: dict) -> str:
    races = ', '.join(r['name'] for r in g['recommended_races'])
    meta = [f'Updated {long_date(g["updated"])}'] + ([f'Recommended races: {races}'] if races else [])
    return PageHeader(spec_title(g), CLASS_COLOR[g['class']], icon(g['icon'], 56), meta)


def glance(g: dict) -> str:
    """The answer first (the one warm card): the role and the stat priority as numbered chips."""
    role = f'<p class="lbl">Role</p><p class="t-lead role">{g["role_line"]}</p>' if g.get('role_line') else ''
    return f'<div class="glance" data-region="summary">{role}<p class="lbl">Stat priority</p>{numbered(g["stat_priority"])}</div>'


def overview(g: dict) -> str:
    o = g['overview']
    body = (f'<p class="prose">{g["summary"]}</p>' + bullets('Strengths', o['strengths']) + bullets('Weaknesses', o['weaknesses'])
            + bullets('In a raid', o['raid']))
    return Section('Overview', body)


def ranks(n: int | None) -> str:
    return '' if n is None else f'{n} rank' if n == 1 else f'{n} ranks'


def talents(g: dict) -> str:
    t = g['talents']
    rows = [Row(f'{k["name"]} <span class="mono">{ranks(k.get("max_rank"))}</span>', k.get('text') or '', icon(k['icon'], 40), link=False,
                foot=f'<p class="t-small why">{k["why"]}</p>' if k.get('why') else '') for k in t['key']]
    builds = [Row(label, '', trail=f'<span class="mono">{" / ".join(str(x) for x in b["totals"])}</span>'
                  + (' <span class="t-small">Level 60</span>' if key == 'raid_build' else ''), link=False, slim=True)
              for key, label in (('raid_build', 'Raid build'), ('leveling_build', 'Leveling build')) if (b := t.get(key))]
    planner = [Row('Open this build in the planner &rsaquo;', '')] if any((t.get(k) or {}).get('code') for k in ('raid_build', 'leveling_build')) else []
    return Section('Talents', Rows(rows, open_=True) + Rows(builds + planner, open_=True))


def spell_rows(spells: list[dict]) -> str:
    def lead(s: dict) -> str:
        return icon(s['icon'] or NO_ICON, 40) if s.get('spell_id') else ''
    return Rows([Row(f'{s["name"]}' + (f' <span class="mono">Rank {s["rank"]}</span>' if s.get('rank') else ''), s['clause'], lead(s),
                     rank=str(i), link=False) for i, s in enumerate(spells, 1)], open_=True)


def rotation(g: dict) -> str:
    r = g['rotation']
    groups = ''.join(f'<p class="lbl">{label}</p>' + spell_rows(r[key]) for key, label in (('opener', 'Opener'), ('priority', 'Priority'),
                                                                                         ('cooldowns', 'Cooldowns')) if r.get(key))
    return Section('Rotation', groups + '<p class="t-small" style="margin-top:8px">Ranks are the ones you have at 60.</p>')


def stats(g: dict) -> str:
    cap = f'<p class="prose">{g["hit_cap"]}</p>' if g.get('hit_cap') else ''
    return Section('Stats', f'<div style="padding-top:12px">{numbered(g["stat_priority"])}</div>{cap}')


def gear_link(g: dict) -> str:
    gear = g.get('gear')
    if not gear:
        return ''
    return Section('Gear', Rows([Row(f'Best in slot for {g["spec_name"]}{DOT}{gear["band"]}{DOT}{gear["preset_label"]} &rsaquo;', '')], open_=True))


def races(g: dict) -> str:
    rows = [Row(r['name'], '<br>'.join(f'{x["name"]}: {x["effect"]}' for x in r['racials']),
                f'<img class="fmark" src="{FACTION_MARK.format(faction=r["faction"])}" alt="{r["faction"].capitalize()}">',
                Tag('Recommended', 'kill') if r['recommended'] else '', link=False) for r in g['races']]
    return Section('Races', Rows(rows, open_=True))


def sources(g: dict) -> str:
    rows = Rows([Row(s['label'], '', link=False, slim=True) for s in g['sources']], open_=True)
    return f'<div class="blk" data-region="section-sources">{Disclosure("Sources", rows, meta=str(len(g["sources"])))}</div>'


def guide(state: str):
    """F9.2: the head band, the section chips, the at-a-glance card and the guide's sections, then the site row."""
    def b(platform: str) -> str:
        g = data(GUIDE)
        name = spec_title(g)
        if state == 'loading':
            body = (skeleton_block([(200, 56), (240, 12)], 'head') + skeleton_block([(60, 10), (180, 16), (330, 72)], 'summary'))
            return app_page(platform, screen(body, name), 'home', back='Warrior')
        if state == 'error':
            return app_page(platform, screen(error_block('This guide did not load.'), name), 'home', back='Warrior', title=name)
        body = (guide_head(g) + chip_strip(SECTIONS, 'Overview') + glance(g) + overview(g) + talents(g) + rotation(g) + stats(g)
                + gear_link(g) + races(g) + sources(g) + site_row('Full guide on the site', SITE_GUIDE))
        return app_page(platform, screen(body, name), 'home', back='Warrior', title=name, share=True)
    return b


def guides_class(platform: str) -> str:
    """F9.6: a class's spec rows, each Saved or Not saved (the scenario keeps Thandrel's class on the phone)."""
    cls = next(c for c in data('cdn.guides.index')['classes'] if c['class'] == 'warrior')
    saved = {'warrior-fury', 'warrior-protection'}
    rows = [Row(s['spec_name'], f'{ROLE[s["role"]]}{DOT}Updated {short_date(s["updated"])}', icon(s['icon'], 40),
                Tag('Saved', 'kill') if s['spec'] in saved else Tag('Not saved', 'muted')) for s in cls['specs']]
    body = PageHeader(cls['class_name'], WARRIOR, crest_img('warrior', 56)) + Section('Specs', Rows(rows, open_=True))
    return app_page(platform, screen(body, cls['class_name']), 'home', back='Guides', title=cls['class_name'])


# ------------------------------------------------------------------------------------------------ Patch notes (F9.5)
TODAY = '2026-11-06'


def date_block(text: str) -> str:
    return f'<span class="dateb">{text}</span>'


def patch_notes(platform: str) -> str:
    """Key dates still ahead, then the changelog newest first (each opens in the reader), then the site row."""
    dates = [d for d in data('cdn.dates') if (d.get('endIso') or d['iso']) >= TODAY]
    date_rows = [Row(d['value'], d.get('note', ''), date_block(d['key']), link=False) for d in dates]
    changes = [Row(e['title'], '', date_block(short_date(e['date'])), Tag('Game' if e['kind'] == 'game' else 'Site', 'muted'))
               for e in data('cdn.changelog')['entries']]
    body = (PageHeader('Patch notes') + Section('Key dates', Rows(date_rows, open_=True)) + Section('What changed', Rows(changes, open_=True))
            + site_row('Full patch notes on the site', 'Every change with its sources, the game and the site'))
    return app_page(platform, screen(body, 'Patch notes'), 'home', back='Home', title='Patch notes')


# ------------------------------------------------------------------------------------------------ Account (F8, F10)
BATTLETAG = 'Thandrel#1604'
PREMIUM_LINE = 'Premium &middot; renews 27 Oct 2027'
FEATURES = ['Coach: after every logged raid, what to fix first. It names the casts you missed, the buffs you let drop and the consumables you skipped.',
            'Server sims keep running, and the site tells you when they finish.',
            'See if you are improving: your Forever Score and every boss parse on a chart over time, with the fights behind each point.',
            'A message in your own Discord channel when a server sim finishes or a log is parsed.',
            'A Supporter mark beside your name on your characters and in your guild&rsquo;s roster.']


def plan_row(premium: bool) -> str:
    """Free in a link region (the US): See plans opens foreversixty.gg/premium; Premium: the plan line, opening Plan."""
    if premium:
        return Row('Plan', PREMIUM_LINE, trail=Tag('Founder'))
    return Row('Plan', 'Free', trail='<a class="link" href="#">See plans &#8599;</a>', link=False)


def account_rows(premium: bool) -> str:
    rows = [plan_row(premium),
            Row('Character', f'<span style="color:{WARRIOR}">Thandrel</span>{DOT}US{DOT}PvP', crest_img('warrior', 32), link=False),
            Row('Notifications', 'Raids, Gear and sims, Logs on'),
            Row('Devices', 'This phone and 2 more'),
            Row('Offline data', '26.6 MB on this phone'),
            Row('Refresh from Battle.net', 'Updated 2 h ago', trail='<a class="link" href="#">Refresh</a>', link=False),
            Row('Manage on the site &rsaquo;', 'Discord, your name on the site, guild seats')]
    return Section('Account', Rows(rows, open_=True))


def about_rows() -> str:
    return Section('About', Rows([Row('Privacy &#8599;', ''), Row('Terms &#8599;', ''), Row('About Forever Sixty', '')], open_=True))


def account(state: str):
    """Settings, the menu's last Account row: the plan, the character, notifications, devices, offline data, the Battle.net
    refresh, one row for what only the site manages, then sign out and delete."""
    def b(platform: str) -> str:
        head = PageHeader('Settings', meta=[BATTLETAG, 'Signed in with Battle.net'])
        if state == 'loading':
            body = PageHeader('Settings') + skeleton_block([(361, 56)] * 6)
        elif state == 'error':
            body = PageHeader('Settings') + error_block('We could not load your account', 'Check your connection')
        else:
            body = (head + account_rows(state == 'premium') + about_rows()
                    + '<div class="vstack">' + Button('Sign out') + Button('Delete account', tone='danger') + '</div>'
                    + '<p class="t-small version">Version 0.1.0 (14) &middot; data 1.60.1.70790</p>')
        return app_page(platform, screen(body, 'Settings'), 'home', back='Home', title='Settings')
    return b


def plan(platform: str) -> str:
    """F8.4 Plan, a Premium player in a link region (the US): the plan line, what it includes, Manage billing on the site.
    A free player never reaches it: their Plan row is See plans (US) or plain Free (everywhere else)."""
    body = (PageHeader('Premium', meta=[f'{PREMIUM_LINE}{DOT}$40 a year', Tag('Founder')])
            + '<p class="t-small" style="margin-top:8px">Founder price, kept while your plan stays active.</p>'
            + Section('What your plan includes', '<ul class="ticks">' + ''.join(f'<li>{x}</li>' for x in FEATURES) + '</ul>')
            + '<div class="vstack">' + Button('Manage billing on foreversixty.gg &#8599;') + '</div>'
            + footer('Raid sign-ups, best in slot, talents and your character are free for everyone, in the app and on the web.'))
    return app_page(platform, screen(body, 'Plan'), 'home', back='Settings', title='Plan')


# ------------------------------------------------------------------------------------------------ Guild: Roster, Calendar, Raid history (F16)
RANK = {'leader': 'Guild Master', 'officer': 'Officer', 'member': 'Member'}
ROSTER_ROWS = 10


def roster_row(m: dict) -> str:
    att = m['attendance']
    trail = (f'<span class="t-small">{att["present"]} of {att["nights"]} nights</span>'
             + Meter(att['present'] / att['nights'], 'kill', f'{att["present"]} of {att["nights"]} nights'))
    level = (m.get('addon') or {}).get('level', 60)
    return Row(m['name'], f'{m["spec_name"]}{DOT}{level}{DOT}{RANK[m["rank"]]}', icon(m['spec_icon'], 40), trail, CLASS_COLOR[m['class']])


def guild_roster(platform: str) -> str:
    """F16.1: the counts, the role chips and the sort, the members with their attendance."""
    home = data('guild.home.roster')
    c = home['counts']
    counts = ('<div class="blk" data-region="summary">'
              + Figures([(str(c['on_roster']), 'on the roster'), (str(c['in_game_no_account']), 'in game, no account'), (str(c['waiting']), 'waiting')])
              + '</div>')
    chips = chip_strip(['All', 'Tanks', 'Healers', 'DPS', 'Sort: Rank'], 'All')
    body = (PageHeader('Roster', meta=['Olympus XXVII']) + counts + chips
            + Section('Members', Rows([roster_row(m) for m in home['roster'][:ROSTER_ROWS]], open_=True), meta=f'{c["on_roster"]} on the roster')
            + footer('Attendance is the last 8 raid nights with a log. Tap a raider for their page.'))
    return app_page(platform, screen(body, 'Roster'), 'raid', back='Raid', title='Roster')


PLANNED = [  # weekday block, zone, when, count line, your answer (None: not signed up)
    ('Wed 6 Jan', RAID, '8:00 pm CST', '21 of 40 signed up', None),
    ('Sat 9 Jan', 'Barrow Deeps', '8:00 pm CST', '9 of 20 signed up', 'In &middot; DPS'),
]


def guild_calendar(platform: str) -> str:
    """F16.4: the planned raids, then the phone's calendar (the switch, which raids, the alert) and the link to share."""
    rows = [Row(zone, f'{when}{DOT}{count}', date_block(day), Tag(you, 'kill') if you else Tag('Not signed up')) for day, zone, when, count, you in PLANNED]
    sync = Rows([Row('Raids in my calendar', 'A calendar called Forever Sixty raids. The app keeps it up to date.', trail=Switch('Raids in my calendar'), link=False),
                 Row('Which raids', 'Only ones I am In or Late for'),
                 Row('Alert', '1 hour before invites')], open_=True)
    share = (Rows([Row('Calendar link', 'For Google Calendar on the web, Outlook and other people. Anyone with the link sees your guild&rsquo;s raid times.',
                       link=False)], open_=True)
             + '<div class="sticky">' + Button('Copy link') + Button('Reset link') + '</div>')
    body = (PageHeader('Calendar', meta=['Olympus XXVII']) + Section('Planned', Rows(rows, open_=True)) + Section('Calendar', sync)
            + Section('Share', share)
            + footer('Your sign-up shows in the event title: &ldquo;Onyxia&rsquo;s Lair &middot; In&rdquo;. Changing it in the app changes the event.'))
    return app_page(platform, screen(body, 'Calendar'), 'raid', back='Raid', title='Calendar')


def guild_raid(platform: str) -> str:
    """F4.2 the raid page, a member: the head, the countdown with the times, your sign-up, who is coming, the officers' notes."""
    times = '<p class="t-small">Invites 7:45 pm &middot; Pull 8:00 pm &middot; Ends 11:00 pm, guild time</p>'
    cd = Countdown(f'Olympus XXVII{DOT}{RAID} tonight', [('02', 'hours'), ('33', 'min')], times, 'guild')
    groups = Rows([Row('Tanks 2', 'Grunvald, Brakka', link=False, slim=True),
                   Row('Healers 5', 'Mirethil, Selka, Hollowmend, Aenwyn, Brightwen', link=False, slim=True),
                   Row('Damage 11', 'Kaelra, Thalrin, Varro and 8 more', link=False, slim=True),
                   Row('Late 3', 'Brightwen, Ysolde, Torvak', link=False, slim=True),
                   Row('Not signed up 4', 'Mirethil, Selka, Varro, Ysolde', link=False, slim=True)], open_=True)
    body = (PageHeader(RAID, 'var(--strong)', icon(ONY, 56), ['Olympus XXVII', WHEN]) + f'<div class="blk" data-region="countdowns">{cd}</div>'
            + Section('Your sign-up', '<p class="t-small me" style="margin:12px 0 8px">You have not answered, as Thandrel, Damage</p>' + signup_buttons())
            + Section('Who is coming', f'<div class="blk">{Tally(ROLES)}</div>' + groups)
            + Section('Notes', '<p class="prose">Bring fire resistance gear and flasks.</p>'))
    return app_page(platform, screen(body, RAID), 'raid', back='Calendar', title=RAID, share=True)


def night_facts(r: dict) -> tuple[str, str, str]:
    """The date block, `5 bosses · 5 kills · 1 wipe` and `10 raiders · 1 h 52 m` (a dungeon counts players)."""
    at = datetime.fromisoformat(r['created_at'].replace('Z', '+00:00')).astimezone(CHICAGO)
    bosses = len({f['name'] for f in r.get('fights', [])}) or r['kill_count']
    wipes = 'no wipes' if not r['wipe_count'] else f'{r["wipe_count"]} wipe' + ('s' if r['wipe_count'] > 1 else '')
    fights = f'{bosses} boss{"es" if bosses != 1 else ""}{DOT}{r["kill_count"]} kill{"s" if r["kill_count"] != 1 else ""}{DOT}{wipes}'
    h, m = divmod(r['duration_ms'] // 60000, 60)
    size = f'{r["raiders"]} {"raiders" if r["raiders"] > 5 else "players"}{DOT}{f"{h} h {m} m" if h else f"{m} m"}'
    return f'{at:%a}<br>{at.day} {at:%b}', fights, size


def night_row(r: dict) -> str:
    block, fights, size = night_facts(r)
    there = any(p['character_key'] == 'us/pvp/obnoxious-yell' for p in r.get('present', []))
    return Row(r['zone'], f'{fights}<br>{size}', date_block(block), foot=Tag('You were there', 'kill') if there else '')


def guild_history(platform: str) -> str:
    """F16.3 Raid history: the guild's logged nights, newest first; a night opens your raid night or its report."""
    rows = [night_row(r) for r in data('guild.raids.history')['rows'][:8]]
    body = (PageHeader('Raid history', meta=['Olympus XXVII']) + Section('Logged raid nights', Rows(rows, open_=True))
            + '<div class="vstack">' + Button('Load more') + '</div>'
            + footer('A night opens your raid night when you were in it, otherwise its report.'))
    return app_page(platform, screen(body, 'Raid history'), 'raid', back='Raid', title='Raid history')


# ------------------------------------------------------------------------------------------------ Sims (F5)
SIM_SITE = ('Run a new sim on the site', 'Gear, talents and consumables are set on foreversixty.gg/sim')
SAVED = [('Raid night set', 'Fury Warrior &middot; Sim', '1,012.3 DPS', 'Saved from the browser &middot; 4 Nov'),
         ('Lionheart Helm', 'Fury Warrior &middot; Top Gear', 'Lionheart Helm is +14.5 DPS', 'Saved from the browser &middot; 2 Nov'),
         ('Fury Warrior', 'Sim', '936.1 DPS', 'Ran on the server &middot; 29 Oct')]


def sim_row(title: str, meta: str, state: str, lead: str = 'ability_warrior_innerrage', trail: str = '', unread: bool = False, foot: str = '') -> str:
    row = Row(title, f'{meta}<span class="state">{state}</span>', icon(lead, 40), trail, link=not trail, foot=foot)
    return row.replace('class="row', 'class="unread row', 1) if unread else row


def mine_section() -> str:
    rows = [sim_row(t, f'{m}{DOT}<b>{h}</b>', o) for t, m, h, o in SAVED]
    return Section('Your sims', Rows(rows, open_=True))


def server_section(running: bool) -> str:
    """Server runs: done rows open the result; a run in flight shows its state (a bulk run its combinations); a failed
    run's one control is Run again."""
    done = data('sims.list')['rows'][0]
    rows = [sim_row(done['title'], f'Fury Warrior{DOT}Sim{DOT}<b>{done["headline"]}</b>', 'Finished 2 h ago &middot; Not opened yet', unread=True)]
    if running:
        rows += [sim_row('Fury Warrior', 'Top Gear', '<span class="on">Running</span> &middot; takes about 2 minutes &middot; Started 1 min ago',
                         foot=f'<div class="prog">{Meter(40 / 64, label="40 of 64 combinations")}<span class="t-small">40 of 64 combinations</span></div>'),
                 sim_row('Fury Warrior', 'Talent compare', 'Starting &middot; Started just now')]
    else:
        rows += [sim_row('Fury Warrior', 'Top Gear', '<span class="on">Running</span> &middot; takes about 2 minutes &middot; Started 1 min ago'),
                 sim_row('Fury Warrior', 'Sim', '<span class="bad">Failed: the engine stopped</span> &middot; Ended 1 d ago',
                         trail=Button('Run again', small=True))]
    return Section('Server runs', Rows(rows, open_=True))


def sims(state: str):
    """F5.1. free: a free player's saved sims; premium: the server runs over them; running: runs in flight."""
    def b(platform: str) -> str:
        premium = state != 'free'
        note = ('Start a sim on foreversixty.gg/sim. Server runs keep going after you close the app.' if premium
                else 'Start a sim on foreversixty.gg/sim and it shows here.')
        body = (PageHeader('Sims') + (server_section(state == 'running') if premium else '') + mine_section()
                + footer(note, Rows([Row(f'{SIM_SITE[0]} &rsaquo;', SIM_SITE[1])], open_=True)))
        return app_page(platform, screen(body, 'Sims'), 'home', back='Home', title='Sims')
    return b


def sim_result(platform: str) -> str:
    """F5.3: the answer, the error line, the gear it ran, and the full breakdown on the site."""
    done = data('sims.list')['rows'][0]
    grid = '<div class="ggrid">' + ''.join(icon(ic, 40, f'q{q}') for _, _, ic, q, _, _ in DOLL) + '</div>'
    body = (PageHeader(done['title'], meta=[f'<span style="color:{WARRIOR}">Fury Warrior</span>{DOT}level 60'])
            + f'<div class="answer" data-region="summary">{Figure("1,042.5", unit="DPS", size="big")}'
            + f'<p class="t-small">&plusmn; 6.0{DOT}3,000 iterations{DOT}engine {done["engine_version"]}</p></div>'
            + Section('The gear it ran', grid)
            + site_row('Full breakdown on foreversixty.gg', 'Abilities, the timeline and the distribution'))
    return app_page(platform, screen(body, done['title']), 'home', back='Sims', title=done['title'], share=True)


SCREENS = {
    'guides-class': guides_class,
    'guide': guide('in'),
    'guide-loading': guide('loading'),
    'guide-error': guide('error'),
    'changes': patch_notes,
    'sim': sims('free'),
    'sim-you': sims('premium'),
    'sim-running': sims('running'),
    'sim-result': sim_result,
    'guild-raids': guild_calendar,
    'guild-raid': guild_raid,
    'guild-roster': guild_roster,
    'guild-loot': guild_history,
    'account': account('free'),
    'account-premium': account('premium'),
    'account-loading': account('loading'),
    'account-error': account('error'),
    'premium': plan,
}
