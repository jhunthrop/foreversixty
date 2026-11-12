"""Guild mocks, the Overview states the persona round left out (GUILD-ALL lane, 2026-11-07), in the home B "alive" direction:
before raids open as an officer at level 23 (the owner's screenshot), a low-level member before the raids open, no raid planned
or logged once the raids are open, and a signed-in visitor in no guild. Composed only from ws_prims primitives; the data is
guild_states_fixtures.py's, which guild_fixtures.py writes for the build. Registered in ws_guild2.all_pages().

Before raids open is the first raids' Countdown, then an open Section with the guild's level spread as a Tally, the professions
line and the viewer's one next step (a sentence, a sentence-case Button, an officer's invite link), then the guild's activity
as FeedRows. No Card, no caps title, no boxed tiles: the Overview's one Plate is the next raid, and these states have none.
"""
from datetime import datetime

import guild_overview_data as O
import guild_overview_fixtures as F
import guild_states_fixtures as S
from guild_data import ALLIANCE, C, GUILD, TABS
import guild_data as gd
from ws_guild2 import _view, ch, hdr, pad
from ws_lib import REGION
from ws_prims import (EmptyState, Button, Countdown, Emblem, FeedRow, Link, Notice, Page, PageHeader, Row, Rows, Section, TabStrip, Tally, build_page,
                      crest_img, icon)

RAIDS = [('Barrow Deeps', 'spell_shadow_raisedead'), ('Hyjal Summit', 'spell_nature_timestop'), ("Onyxia's Lair", 'inv_misc_head_dragon_black')]
RAIDS_OPEN = '2026-12-09T00:00:00Z'
# lib/guild/professions.ts: the primary professions, the only ones the line names.
PRIMARY = {'alchemy': 'Alchemy', 'blacksmithing': 'Blacksmithing', 'enchanting': 'Enchanting', 'engineering': 'Engineering', 'herbalism': 'Herbalism',
           'leatherworking': 'Leatherworking', 'mining': 'Mining', 'skinning': 'Skinning', 'tailoring': 'Tailoring'}
# The raid list beside the clock (RaidsCountdown.svelte), and the step row (GuildWeekOneCard.svelte).
STATES_CSS = '''
.raids{display:flex;justify-content:flex-end;gap:18px;margin:0;padding:0;list-style:none}
.raids li{display:flex;flex-direction:column;align-items:center;gap:6px;width:88px;font:600 13px/16px var(--font-body);text-align:center;color:var(--nav)}
.w1{display:flex;flex-direction:column;gap:16px}.w1 .spread{display:flex;flex-direction:column;gap:8px}
.w1 .step{display:flex;flex-wrap:wrap;align-items:center;justify-content:space-between;gap:12px 24px}.w1 .step p{margin:0;font-size:15px;line-height:24px}
.w1 .step .go{display:flex;flex-wrap:wrap;align-items:center;gap:16px}.w1 .step .go .link{display:inline-flex;align-items:center;min-height:44px}
.w1 .profs{margin:0;font-size:14px;line-height:20px;color:var(--muted)}.w1 .profs b{color:var(--strong);font-weight:600}
.lvl{padding-top:12px;font-size:14px;line-height:20px;color:var(--text)}
.muted-line{margin:0;padding-top:12px;font-size:14px;line-height:20px;color:var(--muted)}.plate .muted-line{padding-top:0;font-size:13px}
'''


def at(iso: str) -> datetime:
    return datetime.fromisoformat(iso.replace('Z', '+00:00'))


def clock_to(target: str, now: str) -> list[tuple[str, str]]:
    secs = int((at(target) - at(now)).total_seconds())
    return [(f'{secs // 86400:02d}', 'days'), (f'{secs % 86400 // 3600:02d}', 'hours'), (f'{secs % 3600 // 60:02d}', 'min')]


def raids_countdown(now: str) -> str:
    """RaidsCountdown: the clock to the opening while ahead, `Raids open now` and no clock once open."""
    raids = '<ul class="raids" aria-label="First raids">' + ''.join(f'<li>{icon(i, 40)}<span>{n}</span></li>' for n, i in RAIDS) + '</ul>'
    if at(now) >= at(RAIDS_OPEN):
        panel = Countdown('Raids open now', None, raids, 'raid', live=True)
    else:
        panel = Countdown('First raids open &middot; Wed 9 Dec', clock_to(RAIDS_OPEN, now), raids, 'raid', live=True)
    return f'<div data-region="{REGION["countdowns"]}">{panel}</div>'


def spread(levels: dict) -> dict:
    out = {'at60': 0, 'fifties': 0, 'forties': 0, 'under40': 0, 'none': 0}
    for lv in levels.values():
        if lv is None:
            out['none'] += 1
        elif lv >= 60:
            out['at60'] += 1
        elif lv >= 50:
            out['fifties'] += 1
        elif lv >= 40:
            out['forties'] += 1
        elif lv >= 10:
            out['under40'] += 1
    return out


def level_tally(levels: dict) -> str:
    s = spread(levels)
    n = s['at60'] + s['fifties'] + s['forties'] + s['under40']
    tally = Tally([('', 'At 60', s['at60'], n, '', '' if s['at60'] >= n else 'warn'), ('', '50 to 59', s['fifties'], n, '', 'none'), ('', '40 to 49', s['forties'], n, '', 'none'),
                   ('', 'Under 40', s['under40'], n, '', 'none')])
    none = f'<p class="t-small" style="margin:0;color:var(--muted)">{s["none"]} with no level on file yet.</p>' if s['none'] else ''
    return f'<div class="spread">{tally}{none}</div>'


def professions(levels: dict) -> str:
    holders: dict[str, list[str]] = {}
    for name, profs in S.PROFESSIONS.items():
        if (levels.get(name) or 0) < 10:
            continue
        for p in profs:
            if p in PRIMARY:
                holders.setdefault(PRIMARY[p], []).append(name)
    facts = []
    for prof in sorted(holders):
        names = sorted(holders[prof])
        more = len(names) - 3
        facts.append(f"{prof}: {', '.join(names[:3])}{f' +{more}' if more > 0 else ''}")
    return f'<p class="profs"><b>Professions</b> {" &middot; ".join(facts)}</p>' if facts else ''


def step(sentence: str, go: str) -> str:
    return f'<div class="step"><p>{sentence}</p><div class="go">{go}</div></div>'


def week_one(title: str, levels: dict, step_html: str, meta: str = '') -> str:
    body = f'<div class="w1">{level_tally(levels)}{professions(levels)}{step_html}</div>'
    return Section(title, body, meta=meta, panel=True)


def ago(iso: str, now: str) -> str:
    m = int((at(now) - at(iso)).total_seconds() // 60)
    if m < 60:
        return f'{m} min ago'
    return f'{m // 60} h ago' if m < 1440 else f'{m // 1440} d ago'


def activity(feed: dict, now: str) -> str:
    """The guild's activity (GuildActivity.svelte): the Your guild rows of /v1/feed, crest led."""
    rows = ''
    for r in feed['rows']:
        name, rest = r['segments'][0], r['segments'][1]['text']
        rows += FeedRow(crest_img(r['crest'], 40), f'<b style="color:{C[name["class"]]}">{name["text"]}</b>{rest}', ago(r['at'], now), 'Your guild')
    return Section('Guild activity', f'<div>{rows}</div>', panel=True)


def last_raid(progression: str, logged: bool) -> str:
    if logged:
        n = O.last_raid_counts()
        row = Row(f'{O.ZONE_NAME} &middot; raid night', f"{F.day_label(O.LAST_RAID_AT)} &middot; 1 boss ({O.BOSS}) &middot; {n['kills']} kill &middot; {n['wipes']} wipe",
                  trail='<span class="gold" style="font-weight:600">Report &rsaquo;</span>')
        body = Rows([row], open_=True)
    else:
        body = EmptyState('No raid logged yet.', 'Get the companion')  # LastRaid.svelte: the companion fills it
    return Section('Last raid', body, meta=progression, panel=True)



def head(officer: bool, tabs: list[str] | None = None, raiders: int | None = gd.RAIDERS) -> str:
    """The guild's head as Guild.svelte draws it (ws_guild2.hdr): the faction crest, the one band line, the tabs."""
    return hdr(officer=officer, tabs=tabs, raiders=raiders)


def leveling_line(name: str, level: int) -> str:
    return f'<div class="pc"><p class="lvl">{name} is level {level}. Rank and readiness checks start at 60.</p></div>'


def page(title: str, body: str, who: tuple) -> str:
    return build_page(title, _view(body) + f'<style>{STATES_CSS}</style>', 'Guild', who)


def before_raids():
    """An officer at level 23 while nothing is logged (the owner's screenshot): nothing on the calendar, so Plan an event."""
    levels = S.PRE_LEVELS
    go = Link('Copy the invite link &rsaquo;') + Button('Plan an event', primary=True, small=True, act=True)
    main = (raids_countdown(S.PRE_NOW) + week_one('Before raids open', levels, step('Nothing on the calendar yet. Post a dungeon run or a guild event.', go))
            + activity(S.pre_feed(), S.PRE_NOW))
    rail = ''  # an officer's rail is the cover and attendance (batch 3): nothing before a raid is planned
    body = head(True) + leveling_line(O.OFFICER, S.LOW_LEVEL) + Page('split', main, rail, sticky_rail=True, panels=True)
    return page('Guild, before raids open', body, ('Dunmaro', 'hunter', 'Beast Mastery Hunter &middot; 23'))


def low_level():
    """A level 23 member before the raids open: the level spread, then how far to 60 and the leveling guide."""
    name, levels = S.LOW_MEMBER, S.PRE_LEVELS
    cls = ch(name)[1]
    go = Button('Level as Subtlety', primary=True, small=True)
    sentence = f'<span style="color:{C[cls]};font-weight:600">{name}</span> is level {S.LOW_LEVEL}, {60 - S.LOW_LEVEL} levels to 60.'
    main = raids_countdown(S.PRE_NOW) + week_one('Before raids open', levels, step(sentence, go)) + activity(S.pre_feed(), S.PRE_NOW)
    rail = last_raid('No boss down this tier yet.', False)
    body = head(False) + leveling_line(name, S.LOW_LEVEL) + Page('split', main, rail, sticky_rail=True, panels=True)
    return page('Guild, low-level member', body, (name, cls, 'Subtlety Rogue &middot; 23'))


def no_raid():
    """The raids are open, nothing is planned and nothing is logged: the member's honest Overview."""
    name, levels = S.NO_RAID_MEMBER, S.OPEN_LEVELS
    cls = ch(name)[1]
    go = Button('Open your best in slot', primary=True, small=True)
    sentence = f'<span style="color:{C[cls]};font-weight:600">{name}</span> is level 60. Check your gear before the first raid.'
    main = (raids_countdown(S.OPEN_NOW) + week_one('Your first raid', levels, step(sentence, go), meta='Your first logged night starts the progression board.')
            + activity(S.open_feed(), S.OPEN_NOW))
    rail = last_raid('No boss down this tier yet.', False)
    body = head(False) + Page('split', main, rail, sticky_rail=True, panels=True)
    return page('Guild, no raid yet', body, (name, cls, 'Holy Priest &middot; 60'))


def public():
    """A signed-in visitor in no guild: the recruiting line, the last raid, and the roster closed to them."""
    notice = Notice(f'{GUILD} is recruiting.', Link('See what they look for &rsaquo;'))
    main = notice + last_raid('1 of 1 named bosses down: Onyxia.', True)
    gate = Section('Roster', '<p class="muted-line">Roster is for members of this guild. None of your characters is on it.</p>')
    body = head(True, tabs=['Overview', 'Raids', 'Progression'], raiders=None)
    body += Page('split', main, '', sticky_rail=False, panels=True) + f'<div class="pc" style="padding-top:61px">{gate}</div>'
    return page('Guild, visitor', body, ('Felsnap', 'warrior', 'Fury Warrior &middot; 34'))


PAGES: dict = {'web-guild-before-raids': before_raids, 'web-guild-low-level': low_level, 'web-guild-no-raid': no_raid, 'web-guild-public': public}
