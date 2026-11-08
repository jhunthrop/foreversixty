"""The flows-first app screens (design/specs/2026-11-08-app-flows.md section 5.1): the profile menu, the time-aware Home
states, the Guild tab's first screen, the raid notification and its confirmation, the officer's Not answered raid page,
the applicant, the first run and the item sheet with the guild's needs. ws_app.py registers them; the frame is
ws_app_frame.py.

One scenario runs through every screen (review round 1, M1): Thandrel, Fury Warrior 60 and the main, an officer of
Olympus XXVII; the guild's raid is Onyxia's Lair, Wed 9 Dec, 8:00 pm CST, the day the first raids open; 18 in and 3 late,
2 of 2 tanks, 5 of 8 healers, 11 of 30 damage. The loot is Onyxia's, from data/builds (Deathbringer, 8%, inv_axe_09).

The Home states are the web home mock (design/mocks/web/web-home-60.html and -34) with the order of section 4 applied
as CSS and the raid blocks injected; the rest are composed from ws_prims alone.
"""
from __future__ import annotations

import re

from mocklib import CLASS_COLOR
from ws_app_frame import app_page, from_web, glyph
from ws_prims import (Button, Checklist, Countdown, Figures, ItemRow, PageHeader, Row, Rows, Section, Segmented, Switch,
                      Tag, Tally, build_page, crest_img, icon)

WARRIOR = CLASS_COLOR['warrior']
ONY = 'inv_misc_head_dragon_black'
RAID = 'Onyxia&rsquo;s Lair'
WHEN = 'Wed 9 Dec, 8:00 pm CST'
AS_ME = 'as Thandrel, Damage'
AS_ME_CAP = 'As Thandrel, Damage'
COUNTS = '18 in, 3 late'
ROLES = [('', 'Tanks', 2, 2), ('', 'Healers', 5, 8), ('', 'Damage', 11, 30)]
DEATHBRINGER = ('Deathbringer', 'inv_axe_09')

FLOWS_CSS = '''
.ap{padding:8px 16px 24px}
.ap .t-title{font-size:32px;line-height:38px}
.ap section{margin-top:28px}
.center{text-align:center}.center .t-small{max-width:300px;margin:8px auto 0}
.seal{display:block;width:84px;height:84px;margin:40px auto 8px}
.lbl{margin:24px 0 8px;font:700 12px/16px var(--font-body);color:var(--muted)}
.btns{display:grid;grid-template-columns:repeat(4,1fr);gap:8px}.btns .btn{justify-content:center}
.chips{display:flex;gap:8px}.chips .btn{flex:1;justify-content:center}
.vstack{display:flex;flex-direction:column;gap:12px;margin-top:20px}.vstack .btn{justify-content:center}
.sticky{display:grid;grid-template-columns:1fr 1fr;gap:12px;padding:12px 0 4px;background:var(--bg);border-top:1px solid var(--line-soft);margin-top:24px}
.sticky .btn{justify-content:center}
.codeline{display:flex;align-items:center;gap:8px;margin-top:8px;padding:10px 12px;border:1px solid var(--line);border-radius:4px;font:500 14px/20px var(--font-mono)}
.codeline span{flex:1}
.preview{margin-top:12px;padding:10px 12px;border-left:3px solid var(--line);color:var(--text);font-size:14px;line-height:20px}
/* the drawer (flows spec section 3): 85% of the width over the dimmed tab, the tab bar covered (round 1, N1) */
.drawer-wrap{display:grid;grid-template-columns:85% 15%;min-height:var(--dev-h)}
.drawer{background:var(--raised);border-right:1px solid var(--line);padding:calc(var(--status) + 12px) 16px 24px;min-width:0}
.drawer .ph{margin-top:0}
.drawer section{margin-top:24px}.drawer section .sec-h{margin-bottom:4px}
.drawer .stamp{margin-top:6px}
.drawer .foot{margin-top:28px;display:flex;flex-direction:column;gap:6px}
.dim{background:#000000a6}
.setrow{display:flex!important;align-items:center;gap:12px}.setrow .b{flex:1;min-width:0}.setrow .swt{flex:none}
.rowbtn{display:flex;align-items:center;min-height:44px;color:var(--gold);font:600 15px/20px var(--font-body)}
/* a notification on the lock screen */
.lockbg{background:linear-gradient(180deg,#17202f,#07090d);border-radius:24px;padding:16px 12px;display:flex;flex-direction:column;gap:12px}
.nt{background:#1b2230;border-radius:20px;padding:12px 14px;display:flex;flex-direction:column;gap:8px}
[data-platform=android] .nt{border-radius:16px;background:#202734}
.nt .top{display:flex;align-items:center;gap:8px;font:600 13px/16px var(--font-body);color:var(--muted)}
.nt .top .app{display:inline-flex;width:20px;height:20px;border-radius:5px;background:var(--bg);color:var(--gold);align-items:center;justify-content:center}
.nt .nm{font:700 15px/20px var(--font-body);color:var(--strong)}
.nt .ln{font:400 15px/20px var(--font-body);color:var(--text)}
.nt .acts{display:grid;grid-template-columns:repeat(4,1fr);border-top:1px solid #ffffff1f;margin:4px -14px -12px}
.nt .acts a{display:flex;align-items:center;justify-content:center;min-height:44px;color:var(--strong);font:600 15px/20px var(--font-body)}
.nt .acts a+a{border-left:1px solid #ffffff1f}
[data-platform=android] .nt .acts{display:flex;gap:16px;border:0;margin:0}
[data-platform=android] .nt .acts a{min-height:48px;color:var(--gold);justify-content:flex-start}[data-platform=android] .nt .acts a+a{border:0}
/* the injected Home blocks */
.fl-blk{padding:24px 16px 0}
.fl-blk .btns{margin-top:12px}
.fl-blk .tally{margin-top:16px;grid-template-columns:repeat(3,minmax(0,1fr))}
.fl-blk .me{margin-top:8px}
.fl-blk .rows{margin-top:8px}
'''


def _page(body: str, title: str) -> str:
    return build_page(title, f'<div class="ap">{body}</div><style>{FLOWS_CSS}</style>',
                      note='Native app mock (design/specs/2026-11-08-app-flows.md).')


def _signup_buttons(selected: str = '') -> str:
    return '<div class="btns" role="group" aria-label="Your sign-up">' + ''.join(
        Button(x, primary=(x == selected), small=True) for x in ('In', 'Late', 'Bench', 'Out')) + '</div>'


def _late_chips() -> str:
    return '<div class="chips" role="group" aria-label="How late">' + ''.join(Button(x, small=True) for x in ('15 min', '30 min', '60 min')) + '</div>'


def _raid_card(officer: bool = False) -> str:
    """Raid day (section 4, F1): the countdown, who you answer as, the four answers, who is coming, what is not ready."""
    cd = Countdown(f'Olympus XXVII &middot; {RAID} tonight', [('02', 'hours'), ('48', 'min')],
                   f'<p class="t-small">{WHEN} &middot; answer by 6:00 pm CST</p>', 'guild')
    me = f'<p class="t-small me">You have not answered, {AS_ME}</p>'
    ready = Rows([Row('2 things before you pull', 'No flask in your bags; boots not enchanted', link=True, slim=True)], open_=True)
    short = (Rows([Row('5 have not answered', 'Healers short by 3', trail=Button('Mention in Discord', small=True), link=False)], open_=True)
             if officer else '')
    return cd + me + _signup_buttons() + Tally(ROLES) + ready + short


# ------------------------------------------------------------------------------------------------ the profile menu
CHARACTERS = [
    ('Thandrel', 'warrior', 'Fury Warrior &middot; 60 &middot; Olympus XXVII', 'Main'),
    ('Lightvow', 'paladin', 'Holy Paladin &middot; 60 &middot; Olympus XXVII', ''),
    ('Kaelor', 'mage', 'Frost Mage &middot; 47', ''),
    ('Felsnap', 'hunter', 'Beast Mastery Hunter &middot; 34', ''),
    ('Brewmist', 'shaman', 'Enhancement Shaman &middot; 22', ''),
    ('Ashmantle', 'warlock', 'Affliction Warlock &middot; 18', ''),
]
LOW = [('Pockets', 'rogue', 'Rogue &middot; 6'), ('Bankalt', 'mage', 'Mage &middot; 1')]
TOOLS = [('Your sims', 'Your server runs and saved sims'), ('Logs and reports', 'Your reports and the coach'),
         ('Talent planner', 'Your builds and the guide builds'), ('Rankings', 'Your spec and your guild'),
         ('Tiers', 'Every spec, this week'), ('What changed', 'The game and the site')]
ACCOUNT = [('Notifications', 'Raids on, gear and sims on'), ('Plan', 'Premium &middot; Founder price'), ('Devices', '3 signed in'),
           ('Offline data', '26.6 MB on this phone'), ('Settings', 'Sign out, delete account')]


def menu(state: str):
    """state: in (6 characters), one (just the main), low (the low-level fold open), out (signed out)."""
    def b(platform: str) -> str:
        if state == 'out':
            you = (PageHeader('Forever Sixty', meta=['Sign in and your characters arrive with their gear, talents and guild.'])
                   + '<div class="vstack">' + Button('Sign in with Battle.net', primary=True) + '</div>')
            tools = Section('Tools', Rows([Row(n, s) for n, s in TOOLS if n in ('Talent planner', 'Tiers', 'What changed')], open_=True))
            body = you + tools
        else:
            me = CHARACTERS[0]
            you = (PageHeader(me[0], WARRIOR, crest_img(me[1], 56), [me[2], Tag('Main')])
                   + '<p class="t-small stamp">Updated 2 h ago from Battle.net &middot; <a class="link" href="#">Refresh</a></p>')
            others = CHARACTERS[1:] if state != 'one' else []
            rows = [Row(n, s, crest_img(c, 40), Tag(t) if t else '', CLASS_COLOR[c]) for n, c, s, t in others]
            if state == 'low':
                rows += [Row(n, s, crest_img(c, 40), '', CLASS_COLOR[c]) for n, c, s in LOW]
            fold = (f'<a class="rowbtn" href="#">{"Hide low-level characters" if state == "low" else "Show low-level characters (2)"}</a>'
                    if state != 'one' else '')
            chars = (Section('Your characters', Rows(rows, open_=True) + fold) if others else
                     Section('Your characters', '<p class="t-small" style="padding-top:8px">Thandrel is your only character. '
                             'New characters arrive here after you play them and log out.</p>'))
            body = (you + chars + Section('Tools', Rows([Row(n, s) for n, s in TOOLS], open_=True))
                    + Section('Account', Rows([Row(n, s) for n, s in ACCOUNT], open_=True)))
        foot = ('<div class="foot"><span class="t-small">Version 0.1.0 (7) &middot; data 1.60.1.70658</span>'
                '<span class="t-small"><a class="link" href="#">Help</a> &middot; <a class="link" href="#">Privacy</a></span></div>')
        html = build_page('Menu', f'<div class="drawer-wrap"><div class="drawer" role="dialog" aria-label="Menu">{body}{foot}</div>'
                          f'<div class="dim" aria-hidden="true"></div></div><style>{FLOWS_CSS}</style>', note='Profile menu mock.')
        return app_page(platform, html, None, navbar=False, who=None if state == 'out' else 'warrior')
    return b


# ------------------------------------------------------------------------------------------------ Home, time-aware
# Section 4 of the flows spec, over the web home mock: the split's wrappers dissolve so each section is one flex item of
# <main>, ordered as /v1/home's `sections`. A section a state leaves out is hidden; an injected block carries its order.
_FLAT = FLOWS_CSS + '''
main{display:flex;flex-direction:column}
main>.pc:has(>.layout-split),main .layout-split,main .layout-split>.main,main .layout-split>.rail{display:contents}
main section[data-region^="section-"]{padding-inline:16px;margin-top:32px;order:9}
main>.pc:has([data-region="countdowns"]){display:none}
'''
# "The head drops to one line when a raid card leads" (section 4): the answer figures and the fact line leave.
_HEAD_LINE = '[data-region="head-answer"],[data-region="head-meta"]{display:none!important}\n'
_HIDE = '{display:none!important}\n'


def _order(region: str, n: int) -> str:
    return f'main section[data-region="{region}"]{{order:{n}}}\n'


def _hide(*regions: str) -> str:
    return ''.join(f'main section[data-region="{r}"]' + _HIDE for r in regions)


UP, FEED, GSIMS = 'section-your-upgrades', 'section-forever-right-now', 'section-guild-sims'
WEEK, MOVERS = 'section-your-spec-this-week', 'section-top-3-for-fury-warrior-60'
# The feed's segments in the web mock, and the two a state opens on instead of All (section 4, review S7, S10).
_SEG_ALL = ('aria-checked="true" class="on">All</a><a href="#" role="radio" aria-checked="false" class="">Your spec</a>'
            '<a href="#" role="radio" aria-checked="false" class="">Your guild</a>')
SEG_GUILD = (_SEG_ALL, 'aria-checked="false" class="">All</a><a href="#" role="radio" aria-checked="false" class="">Your spec</a>'
                       '<a href="#" role="radio" aria-checked="true" class="on">Your guild</a>')
# A leveler without a guild has no Your guild segment: All, Your spec, Logs.
SEG_SPEC = ('aria-checked="true" class="on">All</a><a href="#" role="radio" aria-checked="false" class="">Your spec</a>',
            'aria-checked="false" class="">All</a><a href="#" role="radio" aria-checked="true" class="on">Your spec</a>')


_FROW = re.compile(r'<a class="frow" href="#">.*?<span class="k">([^<]*)</span></span></a>', re.S)


def _filter_feed(html: str, tags: set[str]) -> str:
    """The feed after a filter: the rows whose tag the filter keeps (review round 2: a selected filter must filter)."""
    return _FROW.sub(lambda m: m.group(0) if m.group(1) in tags else '', html)


def _block(order: int, inner: str, region: str = 'countdowns') -> str:
    return f'<div class="fl-blk" style="order:{order}" data-region="{region}">{inner}</div>'


def _raid_line(order: int) -> str:
    """Any other day in a guild: the next raid as one line, its four answers right there (review F1 OVER)."""
    row = Rows([Row(f'{RAID} &middot; {WHEN}', f'Next raid &middot; {COUNTS} &middot; you have not answered', icon(ONY, 40))], open_=True)
    return _block(order, row + _signup_buttons())


def _raid_live(order: int) -> str:
    cd = Countdown(f'Olympus XXVII &middot; {RAID} in progress', None,
                   f'<p class="t-small">Pulled 1 h 12 min ago &middot; {COUNTS} &middot; you are In {AS_ME}</p>'
                   '<p style="margin-top:8px"><a class="link" href="#">Open raid night &rsaquo;</a></p>', 'guild', live=True)
    return _block(order, cd)


def _after_raid(order: int) -> str:
    rows = [Row('Deathbringer', 'You got it &middot; Onyxia &middot; main hand, +12.6 DPS over your worn', icon(DEATHBRINGER[1], 40, 'q4'),
                title_color='#b866f5'),
            Row(f'{RAID} report ready', '87th percentile &middot; the coach has 2 notes', icon(ONY, 40)),
            Row('Your best on Onyxia', '94th percentile, a new best', icon(ONY, 40)),
            Row('Sim again with Deathbringer', 'Your sim is from before it dropped', link=True, slim=True)]
    return f'<div class="pc" style="order:{order}">' + Section('Last night', Rows(rows, open_=True)) + '</div>'


HOME_STATES = {
    # in a guild, raid day before the raid: Next raid (card) · Next upgrade · Feed · Guild sims
    'raid-day': (_FLAT + _HEAD_LINE + _order(UP, 3) + _order(FEED, 4) + _order(GSIMS, 5), lambda: _block(2, _raid_card()), 1, None),
    # the officer's raid day: the same card with the open count (review S3)
    'raid-day-officer': (_FLAT + _HEAD_LINE + _order(UP, 3) + _order(FEED, 4) + _order(GSIMS, 5), lambda: _block(2, _raid_card(officer=True)), 6, None),
    # raid in progress: Raid night card · Feed (Your guild); the rest waits
    'raid-live': (_FLAT + _HEAD_LINE + _order(FEED, 3) + _hide(UP, GSIMS, WEEK, MOVERS), lambda: _raid_live(2), 0, SEG_GUILD),
    # the day after: Raid summary · Next upgrade · Feed
    'after-raid': (_FLAT + _order(UP, 3) + _order(FEED, 4) + _order(GSIMS, 5), lambda: _after_raid(2), 0, None),
    # any other day in a guild: Next upgrade · Next raid (one line, its answers) · Feed · Guild sims
    'normal': (_FLAT + _order(UP, 2) + _order(FEED, 4) + _order(GSIMS, 5),
               lambda: _raid_line(3), 1, None),
    # no guild at 60: Next upgrade · Feed · Your spec this week · First raids (one line) · Find a guild (review S10)
    'noguild': (_FLAT + _order(UP, 2) + _order(FEED, 3) + _order(WEEK, 4) + _hide(GSIMS),
                lambda: (_block(5, Rows([Row('First raids &middot; Wed 9 Dec', 'In 33 days &middot; Barrow Deeps, Hyjal Summit, Onyxia&rsquo;s Lair',
                                              icon(ONY, 40))], open_=True))
                         + _block(6, '<p class="t-small">Raid with a guild that plans its nights here, and answer its raids from the app.</p>'
                                  + '<div class="vstack">' + Button('Find a guild') + '</div>', 'section-find-a-guild')), 0, None),
}
# The web home mock at 60 is a guild member; the no-guild state takes the guild out of the head's facts.
NO_GUILD_SUBS = [('&lt;Olympus XXVII&gt;', '')]


def normal_block() -> str:
    """Any other day in a guild (section 4): the instance line and the next raid with its answers."""
    return HOME_STATES['normal'][1]()


def home_state(state: str, feed_filters_css: str):
    css, block, badge, seg = HOME_STATES[state]

    def b(platform: str) -> str:
        html = from_web('web-home-60', platform, 'home', extra_css=css + feed_filters_css, badge=badge)
        if state == 'noguild':
            for old, new in NO_GUILD_SUBS:
                html = html.replace(old, new)
        if seg:
            html = _filter_feed(html.replace(*seg), {'Your guild'})
        return html.replace('</main>', block() + '</main>', 1)
    b.web = True
    return b


def leveling(feed_filters_css: str, base_css: str):
    """Leveling (section 4, review M4 and S10): head with levels to the next band, upgrades, the instance line, the feed on
    Your spec at the band."""
    def b(platform: str) -> str:
        html = from_web('web-home-34', platform, 'home', extra_css=base_css + feed_filters_css + FLOWS_CSS, who='hunter')
        assert SEG_SPEC[0] in html, 'web-home-34 changed its feed segments'
        html = html.replace(*SEG_SPEC)
        return _filter_feed(html, {'Your spec'})
    b.web = True
    return b


# ------------------------------------------------------------------------------------------------ the Guild tab's first screen
def guild_home(officer: bool):
    """R10 and F1/F3: the next raid first with your answers, then (officer) what needs you, then the guild's pages."""
    def b(platform: str) -> str:
        needs = (Section('Needs you', Rows([Row(f'5 have not answered for {RAID}', 'Healers short by 3', trail=Button('Mention', small=True), link=False),
                                            Row('1 applicant', 'Mirelle, Holy Priest 60, applied 2 min ago')], open_=True), meta='6 open')
                 if officer else '')
        pages = Section('The guild', Rows([Row('Raids', 'Onyxia&rsquo;s Lair Wed, then Barrow Deeps Sat'), Row('Roster', '38 members, 26 at 60'),
                                           Row('Guild sims', 'You are #7 of 14 DPS'), Row('Loot', 'Last night: 4 pieces')], open_=True))
        body = (PageHeader('Olympus XXVII', meta=['Alliance &middot; Normal', 'Officer' if officer else 'Member'])
                + Section('Next raid', _raid_card()) + needs + pages)
        return app_page(platform, _page(body, 'Olympus XXVII'), 'guild', badge=6 if officer else 1)
    return b


# ------------------------------------------------------------------------------------------------ F1, F3: notifications
def notifications(platform: str) -> str:
    """F1: the raid push, long-pressed with the four answers, Late's minutes, and the quiet confirmation; F2: the raid-night
    end state; F3: the applicant push (Review, not Accept, on the lock screen: review M3)."""
    app = f'<span class="app">{glyph("guild", 14)}</span>'
    acts = '<div class="acts">' + ''.join(f'<a href="#">{x}</a>' for x in ('In', 'Late', 'Bench', 'Out')) + '</div>'
    raid = (f'<div class="nt"><div class="top">{app}Forever Sixty &middot; now</div>'
            f'<span class="nm">{RAID} &middot; {WHEN}</span>'
            f'<span class="ln">Are you in, {AS_ME}? {COUNTS}. Answer by 6:00 pm CST.</span>')
    press = 'long-pressed' if platform == 'ios' else 'expanded'
    confirm = lambda nm, ln: (f'<div class="lockbg"><div class="nt"><div class="top">{app}Forever Sixty &middot; now</div>'
                              f'<span class="nm">{nm}</span><span class="ln">{ln}</span></div></div>')
    applicant = (f'<div class="nt"><div class="top">{app}Forever Sixty &middot; 2 min</div><span class="nm">New applicant: Mirelle, Holy Priest 60</span>'
                 '<span class="ln">Forever Score 74 &middot; attuned to Onyxia &middot; healers are 3 short</span>'
                 '<div class="acts" style="grid-template-columns:1fr 1fr"><a href="#">Review</a><a href="#">Decline</a></div></div>')
    ended = (f'<div class="nt"><div class="top">{app}Forever Sixty &middot; 11:46 pm</div><span class="nm">{RAID} done &middot; 1 kill</span>'
             '<span class="ln">You got Deathbringer. The report is ready: 87th percentile.</span></div>')
    body = (PageHeader('Notifications')
            + f'<p class="lbl">Raid, {press} on the lock screen</p><div class="lockbg">' + raid + acts + '</div></div>'
            + f'<p class="lbl">After Late, the minutes ({"a notification content extension" if platform == "ios" else "the expanded notification"})</p>'
            + '<div class="lockbg">' + raid + '<span class="ln">How late?</span>' + _late_chips() + '</div></div>'
            + '<p class="lbl">The answer, sent without opening the app</p>' + confirm('You&rsquo;re In for Onyxia, 8:00 pm CST', f'{AS_ME_CAP} &middot; 19 in, 3 late.')
            + confirm('You&rsquo;re Late (30 min) for Onyxia', f'{AS_ME_CAP} &middot; arriving 8:30 pm CST.')
            + '<p class="lbl">Raid night ends</p><div class="lockbg">' + ended + '</div>'
            + '<p class="lbl">Officer: an applicant</p><div class="lockbg">' + applicant + '</div>')
    return app_page(platform, _page(body, 'Notifications'), None, navbar=False)


# ------------------------------------------------------------------------------------------------ F3: the officer's raid page
# The short role first (review S8): healers, then the rest.
UNANSWERED = [('Mirethil', 'priest', 'Holy Priest &middot; was In last week'), ('Selka', 'druid', 'Restoration Druid &middot; was Late last week'),
              ('Grunvald', 'warrior', 'Protection Warrior &middot; was In last week'), ('Varro', 'rogue', 'Combat Rogue &middot; missed last week'),
              ('Ysolde', 'mage', 'Fire Mage &middot; new this week')]


def raid_unanswered(platform: str) -> str:
    rows = [Row(n, s, crest_img(c, 40), title_color=CLASS_COLOR[c], link=False,
                foot='<div class="btns" style="grid-template-columns:repeat(3,1fr)">' + Button('Mark In', small=True) + Button('Bench', small=True)
                + Button('Message', small=True) + '</div>') for n, c, s in UNANSWERED]
    preview = (f'<p class="preview">@Mirethil @Selka @Grunvald @Varro @Ysolde: {RAID} is {WHEN}. '
               'Answer In, Late, Bench or Out in the app or on the raid page.</p>')
    body = (PageHeader(RAID, 'var(--strong)', icon(ONY, 56), ['Olympus XXVII', WHEN, 'healers short by 3'])
            + '<div class="pc" data-region="controls">' + Segmented([('All 26', False), ('Not answered 5', True), ('In 18', False), ('Late 3', False)], small=True) + '</div>'
            + '<div class="vstack">' + Button('Mention in Discord', primary=True) + '</div>' + preview
            + Section('Not answered', Rows(rows, open_=True), meta='5 of 26, healers first')
            + '<p class="t-small" style="margin-top:16px">Creating and editing raids stays on the site. <a class="link" href="#">Edit on the site &rsaquo;</a></p>')
    return app_page(platform, _page(body, RAID), 'guild', back='Guild', title=RAID, share=True, badge=6)


def applicant(state: str):
    """F3, review M3: the officer sees the note, the history, attunement, enchants and the gear gap before Accept; after
    Accept the invite is in game, so the page hands over the /ginvite line."""
    def b(platform: str) -> str:
        ready = Rows([Row('Attuned to Onyxia&rsquo;s Lair', 'Drakefire Amulet', link=False, slim=True),
                      Row('Enchants', '6 of 8 slots', link=False, slim=True),
                      Row('Best in slot gap', '9 pieces from the Holy Priest 60 raid list', link=False, slim=True)], open_=True)
        history = Rows([Row('Dawnbreak', 'Officer &middot; 14 raids &middot; left 2 Nov', link=False),
                        Row('Ashen Vow', 'Member &middot; 6 raids &middot; left 18 Oct', link=False)], open_=True)
        head = PageHeader('Mirelle', CLASS_COLOR['priest'], crest_img('priest', 56), ['Holy Priest &middot; 60', 'Alliance Dwarf', 'applied 2 min ago'])
        if state == 'accepted':
            body = (head + '<p class="t-lead" style="margin-top:12px">Accepted. Invite them in game:</p>'
                    + '<div class="codeline"><span>/ginvite Mirelle</span>' + Button('Copy', small=True) + '</div>'
                    + '<p class="t-small" style="margin-top:12px">She shows on the roster once the addon or Battle.net sees her in the guild.</p>')
        else:
            body = (head + Figures([('74', 'Forever Score'), ('612', 'Sim', 'HPS'), ('63', 'Item level')])
                    + Section('Raid ready', ready) + Section('Guild history', history)
                    + Section('Their note', '<p class="t-prose" style="padding-top:8px">Holy main since launch. Wednesday and Sunday nights work for me.</p>')
                    + '<div class="sticky">' + Button('Accept', primary=True) + Button('Decline') + '</div>')
        return app_page(platform, _page(body, 'Mirelle'), 'guild', back='Guild', title='Mirelle', badge=6)
    return b


# ------------------------------------------------------------------------------------------------ F8: first run
def pick_main(platform: str) -> str:
    rows = [Row(n, s, crest_img(c, 40), Tag('Main') if i == 0 else '', CLASS_COLOR[c], selected=(i == 0)) for i, (n, c, s, _) in enumerate(CHARACTERS)]
    body = (PageHeader('Pick your main', meta=['From Battle.net, level 10 and up.'])
            + '<p class="t-small" style="margin-top:8px">Home answers for your main, and your raid sign-ups are for it. '
            'Switch any time from your crest, top left.</p>'
            + Section('Your characters', Rows(rows, open_=True)) + '<div class="vstack">' + Button('Continue with Thandrel', primary=True) + '</div>')
    return app_page(platform, _page(body, 'Pick your main'), None, navbar=False)


def permission(platform: str) -> str:
    body = ('<div class="center"><img class="seal" src="../../../web/public/seal-400.png" alt="">' + PageHeader('You&rsquo;re in Olympus XXVII')
            + '<p class="t-small">Get raid sign-ups here. Answer In, Late, Bench or Out from the notification, without opening the app.</p></div>'
            + '<div class="vstack">' + Button('Turn on notifications', primary=True) + Button('Not now') + '</div>')
    return app_page(platform, _page(body, 'Notifications'), None, navbar=False)


def notification_settings(platform: str) -> str:
    """Six groups (F4), each switch on its title row (review S12); Raids and Officer show only to a guild member."""
    groups = [('Raids', 'Posted, moved, your answer due, raid night', True), ('Your gear and sims', 'An upgrade dropped, a sim finished, BiS changed', True),
              ('Logs', 'A report or the coach is ready, a new best', True), ('Guild ladder', 'A guildmate passed you', False),
              ('Tier lists', 'Your spec moved this week', False), ('Officer', 'Short on a role, applicants, late changes', True)]
    # The switch stays on the title row at every width (a Row's trail drops under its body below 600).
    rows = ['<div class="row setrow"><div class="b"><div class="ttl">' + n + '</div><div class="sub">' + s + '</div></div>'
            + Switch('', on) + '</div>' for n, s, on in groups]
    nudge = Rows([Row('Raid-day reminder', 'If you have not answered: 3:00 pm your time', trail='<a class="link" href="#">Change</a>', link=False)], open_=True)
    body = PageHeader('Notifications') + Section('Send me', Rows(rows, open_=True)) + Section('When', nudge)
    return app_page(platform, _page(body, 'Notifications'), 'home', back='Home', title='Notifications')


# ------------------------------------------------------------------------------------------------ F4, F5: the item sheet
def item_sheet(platform: str) -> str:
    """Deathbringer from Onyxia (data/builds 1.60.1.70790: 8%, item level 75): the client's tooltip lines, where it drops,
    the runner-up from the same boss, and who in the guild has it or needs it (review S1)."""
    tooltip = ('<div class="t-small" style="margin-top:8px;line-height:20px">Binds when picked up<br>One-Hand &middot; Axe<br>'
               '114 - 213 Damage &middot; Speed 2.90<br>(56.4 damage per second)<br>'
               '<span style="color:var(--kill)">Chance on hit: Sends a shadowy bolt at the enemy causing 125 Shadow damage.</span><br>'
               'Requires Level 60</div>')
    where = Rows([Row('Onyxia, Onyxia&rsquo;s Lair', '8% drop chance &middot; raid &middot; Wed with Olympus XXVII', icon(ONY, 40), link=False)], open_=True)
    runners = Rows([Row('Vis&rsquo;kag the Bloodletter', 'Onyxia &middot; 8% &middot; &minus;1.8 DPS', icon('inv_sword_18', 40, 'q4'),
                        title_color='#b866f5', link=False)], open_=True)
    needs = Rows([Row('Grunvald', 'Protection Warrior &middot; needs it, also on his list', crest_img('warrior', 40), title_color=WARRIOR, link=False),
                  Row('Torvak', 'Enhancement Shaman &middot; has it', crest_img('shaman', 40), title_color=CLASS_COLOR['shaman'], link=False)], open_=True)
    sheet = ('<div class="dim" style="height:120px"></div><div class="sheetp" role="dialog" aria-label="Deathbringer"><span class="sheet-grab"></span>'
             + ItemRow('Deathbringer', 4, icon(DEATHBRINGER[1], 40, 'q4'), 'Main hand &middot; Item level 75', gain=('+12.6', 'DPS', 'over your worn'))
             + tooltip + Section('Where to get it', where) + Section('Runner-up', runners) + Section('Who in the guild needs it', needs) + '</div>')
    html = build_page('Deathbringer', sheet + f'<style>{FLOWS_CSS}.sheetp{{background:var(--raised);border-radius:12px 12px 0 0;'
                      'border-top:1px solid var(--line);padding:0 16px 24px}.sheet-grab{display:block;width:36px;height:5px;border-radius:3px;'
                      'background:var(--line);margin:8px auto 4px}</style>', note='Item sheet mock.')
    return app_page(platform, html, 'bis', navbar=False)


SCREENS = {
    'menu': menu('in'),
    'menu-one': menu('one'),
    'menu-low': menu('low'),
    'menu-out': menu('out'),
    'guild-member': guild_home(False),
    'guild-officer': guild_home(True),
    'notifications': notifications,
    'notification-settings': notification_settings,
    'guild-raid-unanswered': raid_unanswered,
    'guild-applicant': applicant('review'),
    'guild-applicant-accepted': applicant('accepted'),
    'first-pick-main': pick_main,
    'first-permission': permission,
    'item-sheet-guild': item_sheet,
}
