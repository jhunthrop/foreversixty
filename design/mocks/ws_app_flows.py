"""The flows-first app (design/specs/2026-11-08-app-flows.md): four tabs that are the app's jobs (Home, Raid, Gear, Activity),
the profile menu that holds the tools and the account, and the screens the flows land on. Every element here has a line
in the spec's purpose audit (appendix A); an element with no line does not belong in a mock. ws_app.py registers the
screens; the frame is ws_app_frame.py.

One scenario runs through every screen: Thandrel, Fury Warrior 60 and the main, of Olympus XXVII (a member; the officer
states say so); the guild's raid is Onyxia's Lair, Wed 6 Jan, 8:00 pm CST, four weeks after the first raids opened; 18 in and 3 late,
2 of 2 tanks, 5 of 8 healers, 11 of 30 damage. Gear is the Fury Warrior 60 raid list and Onyxia's loot from data/builds
1.60.1.70790; the leveler is Felsnap, Beast Mastery Hunter 34, on the Beast Mastery 30 list.
"""
from __future__ import annotations

from mocklib import CLASS_COLOR
from ws_app_frame import DEFAULT, app_page, glyph, portrait
from ws_prims import (Button, Checklist, Countdown, Figures, ItemRow, PageHeader, Row, Rows, Section, Skeleton, Switch,
                      Tag, Tally, build_page, crest_img, icon)

WARRIOR, HUNTER = CLASS_COLOR['warrior'], CLASS_COLOR['hunter']
ONY = 'inv_misc_head_dragon_black'
RAID = 'Onyxia&rsquo;s Lair'
WHEN = 'Wed 6 Jan, 8:00 pm CST'
AS_ME = 'as Thandrel, Damage'
AS_ME_CAP = 'As Thandrel, Damage'
COUNTS = '18 in, 3 late'
ROLES = [('', 'Tanks', 2, 2), ('', 'Healers', 5, 8), ('', 'Damage', 11, 30)]
DEATHBRINGER = ('Deathbringer', 'inv_axe_09')
# The tab badges (section 3): the member's unanswered raid, the officer's open items, the unread activity.
MEMBER, OFFICER, ANSWERED = {'raid': 1, 'activity': 3}, {'raid': 6, 'activity': 3}, {'activity': 3}

FLOWS_CSS = '''
.ap{padding:8px 16px 24px}
.ap .pc{padding-inline:0!important}
.ap .t-title{font-size:32px;line-height:38px}
.ap section{margin-top:28px}
.center{text-align:center}.center .t-small{max-width:300px;margin:8px auto 0}
.center .eyebrow{margin:0 auto -12px}.center .t-title{text-wrap:balance}
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
.setrow{display:flex!important;align-items:center;gap:12px}.setrow .b{flex:1;min-width:0}.setrow .swt{flex:none}
/* the drawer (section 3): 85% of the width over the dimmed tab, the tab bar covered */
.drawer-wrap{display:grid;grid-template-columns:85% 15%;min-height:var(--dev-h)}
.drawer{background:var(--raised);border-right:1px solid var(--line);padding:calc(var(--status) + 12px) 16px 24px;min-width:0}
.drawer section{margin-top:24px}.drawer section .sec-h{margin-bottom:4px}
.drawer .stamp{margin-top:6px}
.drawer .foot{margin-top:28px;display:flex;flex-direction:column;gap:6px}
.dim{background:#000000a6}
.rowbtn{display:flex;align-items:center;min-height:44px;color:var(--gold);font:600 15px/20px var(--font-body)}
.blk{margin-top:24px}
.blk .btns{margin-top:12px}
.blk .tally{margin-top:16px;grid-template-columns:repeat(3,minmax(0,1fr))}
.blk .me{margin-top:8px}
.blk .rows{margin-top:8px}
/* Activity: unread rows carry a gold dot and a strong title */
.act-row{position:relative}.act-row.unread::before{content:"";position:absolute;left:-10px;top:24px;width:6px;height:6px;border-radius:50%;background:var(--gold)}
.act-row .when{display:block;font-size:12px;line-height:16px;color:var(--muted);margin-top:2px}
.chiprow{display:flex;gap:8px;overflow-x:auto;padding:4px 0 8px;scrollbar-width:none}
.chiprow .btn{flex:none}
.chiprow{-webkit-mask-image:linear-gradient(90deg,#000 85%,transparent);mask-image:linear-gradient(90deg,#000 85%,transparent)}
.stampl{margin-top:8px}
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
.sheetp{background:var(--raised);border-radius:12px 12px 0 0;border-top:1px solid var(--line);padding:0 16px 24px}
.sheet-grab{display:block;width:36px;height:5px;border-radius:3px;background:var(--line);margin:8px auto 4px}
'''


def _page(body: str, title: str) -> str:
    return build_page(title, f'<div class="ap">{body}</div><style>{FLOWS_CSS}</style>',
                      note='Native app mock (design/specs/2026-11-08-app-flows.md).')


def _signup_buttons(selected: str = '') -> str:
    return '<div class="btns" role="group" aria-label="Your answer">' + ''.join(
        Button(x, primary=(x == selected), small=True) for x in ('In', 'Late', 'Bench', 'Out')) + '</div>'


def _late_chips() -> str:
    return '<div class="chips" role="group" aria-label="How late">' + ''.join(Button(x, small=True) for x in ('15 min', '30 min', '60 min')) + '</div>'


def _raid_card(officer: bool = False, answered: str = '', late: bool = False) -> str:
    """The next raid (F1): the countdown, who you answer as, the four answers (the current one selected, so a change is one
    tap), who is coming, what is not ready; an officer's open items with their one action. late: Late's minutes are open."""
    cd = Countdown(f'Olympus XXVII &middot; {RAID} tonight', [('02', 'hours'), ('48', 'min')],
                   f'<p class="t-small">{WHEN} &middot; answer by 6:00 pm CST</p>', 'guild')
    answered = 'Late' if late else answered
    me = f'<p class="t-small me">{"You are " + answered if answered else "You have not answered"}, {AS_ME}</p>'
    how_late = '<p class="t-small me">How late?</p>' + _late_chips() if late else ''
    ready = Rows([Row('2 things before you pull', 'No flask in your bags; boots not enchanted &middot; from your last log-out, 9 h ago', slim=True)], open_=True)
    return cd + me + _signup_buttons(answered) + how_late + Tally(ROLES) + ready + (_needs_you() if officer else '')


def _needs_you() -> str:
    """The officer's open items (F3): the unanswered with the short role named, the applicant; the badge's number."""
    return Rows([Row('5 have not answered, 2 of them healers', 'Healers short by 3', trail=Button('Preview mention', small=True), link=False),
                 Row('1 applicant', 'Mirelle, Holy Priest 60, 2 min ago')], open_=True)


def _blk(inner: str, region: str = 'countdowns') -> str:
    return f'<div class="blk" data-region="{region}">{inner}</div>'


# ------------------------------------------------------------------------------------------------ the profile menu
CHARACTERS = [  # name, class, line, tag, Battle.net avatar (None: no avatar_url yet, the crest stands in), faction
    ('Thandrel', 'warrior', 'Fury Warrior &middot; 60 &middot; Olympus XXVII', 'Main', 'thandrel', 'alliance'),
    ('Lightvow', 'paladin', 'Holy Paladin &middot; 60 &middot; Olympus XXVII', '', 'lightvow', 'alliance'),
    ('Kaelor', 'mage', 'Frost Mage &middot; 47', '', 'kaelor', 'horde'),
    ('Felsnap', 'hunter', 'Beast Mastery Hunter &middot; 34', '', 'felsnap', 'alliance'),
    ('Brewmist', 'shaman', 'Enhancement Shaman &middot; 22', '', 'brewmist', 'horde'),
    ('Ashmantle', 'warlock', 'Affliction Warlock &middot; 18', '', None, 'alliance'),
]
LOW = [('Pockets', 'rogue', 'Rogue &middot; 6'), ('Bankalt', 'mage', 'Mage &middot; 1')]
# The tools, in the order the owner ruled (2026-11-08); the ones a signed-out visitor can use carry True.
TOOLS = [('Guides', 'Your spec&rsquo;s guide, then every class', True), ('Tiers', 'Every spec, this week', True),
         ('Patch notes', 'The game and the site', True), ('Sims', 'Your server runs and saved sims', False),
         ('Logs and reports', 'Your reports and the coach', False), ('Planner', 'Your builds and the guide builds', True),
         ('Rankings', 'Your spec and your guild', True)]
ACCOUNT = [('Notifications', 'Raids, gear and sims, logs on'), ('Plan', 'Premium &middot; Founder price'), ('Devices', '3 signed in'),
           ('Offline data', '26.6 MB on this phone'), ('Settings', 'Sign out, delete account')]


def menu(state: str):
    """state: in (6 characters), one (just the main), low (the low-level fold open), out (signed out)."""
    def b(platform: str) -> str:
        if state == 'out':
            you = (PageHeader('Forever Sixty', meta=['Sign in and your characters arrive with their gear, talents and guild.'])
                   + '<div class="vstack">' + Button('Sign in with Battle.net', primary=True) + '</div>')
            body = you + Section('Tools', Rows([Row(n, s) for n, s, out in TOOLS if out], open_=True))
        else:
            me = CHARACTERS[0]
            you = (PageHeader(me[0], WARRIOR, portrait(me[5], me[4], 56), [me[2], Tag('Main')])
                   + '<p class="t-small stamp">Updated 2 h ago from Battle.net &middot; <a class="link" href="#">Refresh</a></p>')
            others = CHARACTERS[1:] if state != 'one' else []
            rows = [Row(n, s, portrait(f, av, 40), '', CLASS_COLOR[c]) for n, c, s, _, av, f in others]
            if state == 'low':
                rows += [Row(n, s, crest_img(c, 40), '', CLASS_COLOR[c]) for n, c, s in LOW]
            fold = (f'<a class="rowbtn" href="#">{"Hide low-level characters" if state == "low" else "Show low-level characters (2)"}</a>'
                    if state != 'one' else '')
            chars = Section('Switch character', Rows(rows, open_=True) + fold) if others else ''
            body = (you + chars + Section('Tools', Rows([Row(n, s) for n, s, _ in TOOLS], open_=True))
                    + Section('Account', Rows([Row(n, s) for n, s in ACCOUNT], open_=True)))
        foot = ('<div class="foot"><span class="t-small">Version 0.1.0 (7) &middot; data 1.60.1.70790</span>'
                '<span class="t-small"><a class="link" href="#">Help</a> &middot; <a class="link" href="#">Privacy</a></span></div>')
        html = build_page('Menu', f'<div class="drawer-wrap"><div class="drawer" role="dialog" aria-label="Menu">{body}{foot}</div>'
                          f'<div class="dim" aria-hidden="true"></div></div><style>{FLOWS_CSS}</style>', note='Profile menu mock.')
        return app_page(platform, html, None, navbar=False, who=None if state == 'out' else 'warrior')
    return b


# ------------------------------------------------------------------------------------------------ Home (section 4)
def _head(leveler: bool = False) -> str:
    """One head on every Home state: the class crest, the name in the class colour and the fact line. The profile button
    above it is the portrait, a different image, so the character is never drawn twice the same way (owner, 2026-11-08)."""
    if leveler:
        return PageHeader('Felsnap', HUNTER, crest_img('hunter', 56), ['Level 34 Beast Mastery Hunter', '6 levels to band 40'])
    return PageHeader('Thandrel', WARRIOR, crest_img('warrior', 56), ['Level 60 Fury Warrior', 'Olympus XXVII'])


def _upgrades(leveler: bool = False, tonight: bool = True) -> str:
    """Next upgrades (F5): three, each with its gain and where it comes from; one line names the best dungeon run with its
    gain and chance (left out on a raid day: the raid is tonight)."""
    if leveler:
        rows = [ItemRow('Felbough Arm', 3, icon('inv_spear_08', 40, 'q3'), 'Main hand &middot; Fel Ancient, Ashenvale', gain=('+1.3', 'DPS', '')),
                ItemRow('Cartilage Shapers', 3, icon('inv_gauntlets_18', 40, 'q3'), 'Hands &middot; quest, Lyn the Ignored', gain=('+0.8', 'DPS', '')),
                ItemRow('Ring of Precision', 3, icon('inv_jewelry_ring_14', 40, 'q3'), 'Finger &middot; Blackfathom Deeps', gain=('+0.6', 'DPS', ''))]
        go = 'Best run: Blackfathom Deeps, Ring of Precision +0.6 DPS'
    else:
        rows = [ItemRow('Lionheart Helm', 4, icon('inv_helmet_36', 40, 'q4'), 'Head &middot; Blacksmithing', gain=('+14.5', 'DPS', '')),
                ItemRow('Greenhammer', 4, icon('inv_mace_1h_raidhyjalc60_d_01', 40, 'q4'), 'Off hand &middot; Blacksmithing', gain=('+9.7', 'DPS', '')),
                ItemRow('Don Julio&rsquo;s Band', 4, icon('inv_jewelry_ring_35', 40, 'q4'), 'Finger &middot; Stormpike Guard, needs revered', gain=('+6.8', 'DPS', ''))]
        go = 'Best run: Blackrock Spire, Blackcrow +1.9 DPS, 8%'
    line = Rows([Row(go, '', slim=True)], open_=True) if tonight else ''
    return Section('Next upgrades', ''.join(rows) + line, action='Gear &rsaquo;')


ACTIVITY_PREVIEW = {
    'day': [('ability_warrior_innerrage', 'Kaelra passed you in the guild sims', '1,012.3 DPS, you are #7 of 14 &middot; 1 h ago', True),
            ('inv_misc_pocketwatch_01', 'Sim done: 948.6 DPS &plusmn;6', '+21 vs your last &middot; 3 h ago', True),
            ('inv_misc_book_09', 'Fury Warrior BiS changed', 'A new pick for Hands &middot; yesterday', False)],
    # the day after: Last night shows the loot and the report, so the preview starts below them
    'after': [('ability_warrior_innerrage', 'Kaelra passed you in the guild sims', '1,012.3 DPS, you are #7 of 14 &middot; 8 h ago', True),
              ('inv_misc_pocketwatch_01', 'Sim done: 948.6 DPS &plusmn;6', '+21 vs your last &middot; yesterday', False),
              ('inv_misc_book_09', 'Fury Warrior BiS changed', 'A new pick for Hands &middot; Tue', False)],
    'live': [('inv_misc_head_dragon_black', 'Brightwen is Late, 15 min', 'Onyxia&rsquo;s Lair &middot; 4 min ago', True),
             ('inv_misc_head_dragon_black', 'Selka is In', 'Restoration Druid &middot; 20 min ago', True),
             ('inv_misc_head_dragon_black', 'Healers short for Onyxia tonight', '3 of 8 missing &middot; 3 h ago', False)],
    'leveler': [('inv_misc_book_09', 'Level 34: your 30 to 39 list', 'Beast Mastery Hunter &middot; 2 h ago', True),
                ('inv_misc_pocketwatch_01', 'Sim done: 58.4 DPS &plusmn;1', '+3.1 vs your last &middot; yesterday', False),
                ('inv_misc_book_09', 'Beast Mastery Hunter moved to A', 'Leveling tiers, 30 to 39 &middot; 2 d ago', False)],
}


def _act_row(ic: str, t: str, s: str, unread: bool) -> str:
    return Row(t, s, icon(ic, 40)).replace('class="row', f'class="act-row{" unread" if unread else ""} row', 1)


def _changed(kind: str, unread: int = 3) -> str:
    """New for you (F4, the casual's answer): the three newest Activity rows, unread ones marked; the link carries the
    unread count, the same number as the Activity badge; all of them are one tap away."""
    rows = ACTIVITY_PREVIEW[kind]
    link = f'All activity ({unread}) &rsaquo;' if unread else 'All activity &rsaquo;'
    return Section('New for you', Rows([_act_row(i, t, s, u) for i, t, s, u in rows], open_=True), action=link)


def _home(blocks: str, title: str, badges: dict, who: str = 'warrior', avatar=DEFAULT):
    def b(platform: str) -> str:
        return app_page(platform, _page(blocks, title), 'home', badges=badges, who=who, avatar=avatar)
    return b


def _raid_line(answered: str = '') -> str:
    sub = f'You are {answered}, {AS_ME}' if answered else f'{COUNTS} &middot; you have not answered, {AS_ME}'
    row = Rows([Row(f'{RAID} &middot; {WHEN}', sub, icon(ONY, 40))], open_=True)
    return _blk(row + _signup_buttons(answered))


def _raid_live() -> str:
    return _blk(Countdown(f'Olympus XXVII &middot; {RAID} in progress', None,
                          f'<p class="t-small">Pulled 1 h 12 min ago &middot; {COUNTS} &middot; you are In {AS_ME}</p>'
                          '<p style="margin-top:8px"><a class="link" href="#">Open raid night &rsaquo;</a></p>', 'guild', live=True))


def _last_night() -> str:
    rows = [Row('Deathbringer', 'You got it from Onyxia &middot; main hand, +12.6 DPS over your worn', icon(DEATHBRINGER[1], 40, 'q4'),
                title_color='#b866f5'),
            Row(f'{RAID} report', '87th percentile, your best on Onyxia &middot; the coach has 2 notes', icon(ONY, 40)),
            Row('Sim again with Deathbringer', 'Your sim is from before it dropped', slim=True)]
    return Section('Last night', Rows(rows, open_=True))


def _find_guild() -> str:
    return _blk('<p class="t-small">Raid with a guild that plans its nights here, and answer its raids from the app.</p>'
                + '<div class="vstack">' + Button('Find a guild') + '</div>', 'section-raid-with-a-guild')


HOME = {
    'home-raid-day': lambda: _home(_head() + _blk(_raid_card()) + _upgrades(tonight=False) + _changed('day'), 'Home', MEMBER),
    'home-raid-day-late': lambda: _home(_head() + _blk(_raid_card(late=True)) + _upgrades(tonight=False) + _changed('day'), 'Home', ANSWERED),
    'home-raid-day-officer': lambda: _home(_head() + _blk(_raid_card(officer=True)) + _upgrades(tonight=False) + _changed('day'), 'Home', OFFICER),
    'home-raid-live': lambda: _home(_head() + _raid_live() + _changed('live'), 'Home', ANSWERED),
    'home-after-raid': lambda: _home(_head() + _last_night() + _upgrades() + _changed('after'), 'Home', ANSWERED),
    'home-60': lambda: _home(_head() + _upgrades() + _raid_line() + _changed('day'), 'Home', MEMBER),
    # no avatar_url yet (a new character, a failed media fetch): the crest stands in on the profile button
    'home-60-noavatar': lambda: _home(_head() + _upgrades() + _raid_line() + _changed('day'), 'Home', MEMBER, avatar=None),
    'home-60-answered': lambda: _home(_head() + _upgrades() + _raid_line('In') + _changed('day'), 'Home', ANSWERED),
    'home-noguild': lambda: _home(_head() + _upgrades() + _changed('day', 2) + _find_guild(), 'Home', {'activity': 2}),
    'home-34': lambda: _home(_head(leveler=True) + _upgrades(leveler=True) + _changed('leveler', 1), 'Home', {'activity': 1}, who='hunter'),
    # loading: the head and the blocks at their final size, no shift when they fill
    'home-loading': lambda: _home(_head() + _blk(''.join(f'<div style="margin-top:12px">{Skeleton(361, 56)}</div>' for _ in range(5)), 'section-loading'),
                                  'Home', {}),
    'home-error': lambda: _home(_head() + _blk(Rows([Row('We could not load your home', 'Check your connection', trail=Button('Try again', small=True), link=False)],
                                                    open_=True), 'section-error'), 'Home', {}),
    # signed in, no character yet (F8): the one thing to do
    'home-new': lambda: _home(PageHeader('Your characters are on their way', meta=['Battle.net lists a character after you play it and log out.'])
                              + '<div class="vstack">' + Button('Check again', primary=True) + '</div>', 'Home', {}, who=None),
}


# ------------------------------------------------------------------------------------------------ Raid (the tab)
def raid(state: str):
    """The Raid tab (owner, 2026-11-08): the next raid first with your answer, then for an officer what needs you, then the
    guild's raid pages. Live: raid night leads. No guild: find, join or claim one. Signed out: sign in, then the same."""
    def b(platform: str) -> str:
        pages = Section('Olympus XXVII', Rows([Row('Roster', '38 members, 31 at 60'), Row('Calendar', 'Onyxia Wed, Barrow Deeps Sat'),
                                               Row('Loot history', 'Last raid: 4 pieces'), Row('Guild sims', 'You are #7 of 14 DPS')], open_=True))
        if state in ('member', 'officer'):
            needs = Section('Needs you', _needs_you(), meta='6 open') if state == 'officer' else ''
            body = Section('Next raid', _blk(_raid_card())) + needs + pages
            badges = OFFICER if state == 'officer' else MEMBER
        elif state == 'noraid':
            body = (Section('Next raid', '<p class="t-small" style="padding-top:8px">No raid planned yet. Your officers plan raids on the site; '
                            'a push tells you the moment one is posted.</p>') + pages)
            badges = ANSWERED
        elif state == 'live':
            ready = (Checklist([('ok', 'Drakefire Amulet', 'Attuned', None), ('miss', 'Flask of the Titans', 'None in your bags', 'Auction house'),
                                ('warn', 'Boots not enchanted', 'Enchant Boots: Greater Agility', 'Where to get it')])
                     + '<p class="t-small stampl">From your last log-out, 9 h ago</p>')
            body = (Section('Raid night', _raid_live().replace('<p style="margin-top:8px"><a class="link" href="#">Open raid night &rsaquo;</a></p>', '')
                            + Rows([Row('Tanks 2 of 2', 'Grunvald, Brakka', link=False, slim=True),
                                    Row('Healers 5 of 8', 'Mirethil, Selka, Hollowmend, Aenwyn, Brightwen (Late, 15 min)', link=False, slim=True)], open_=True))
                    + Section('Ready check', ready) + pages)
            badges = ANSWERED
        elif state == 'none':
            body = (PageHeader('Raid with a guild', meta=['Guilds that plan their raids here'])
                    + Section('Guilds recruiting a Fury Warrior', Rows([
                        Row('Olympus XXVII', 'Alliance &middot; Wed and Sat, 8:00 pm CST &middot; needs damage', trail=Button('Apply', small=True), link=False),
                        Row('Dawnbreak', 'Alliance &middot; Tue and Thu, 9:00 pm EST &middot; needs damage', trail=Button('Apply', small=True), link=False)], open_=True))
                    + Section('Your guild is not here?', '<p class="t-small" style="padding-top:8px">Claim it: a guild master or an officer '
                              'confirms with the addon or Battle.net.</p><div class="vstack">' + Button('Claim your guild') + '</div>'))
            badges = {'activity': 2}
        elif state == 'leveler':
            body = (PageHeader('Raids start at 60', meta=['26 levels to go'])
                    + '<p class="t-small" style="margin-top:8px">Onyxia&rsquo;s Lair, Barrow Deeps and Hyjal Summit open at 60. '
                    'Join a guild on the way and its raids show here.</p>'
                    + '<div class="vstack">' + Button('Find a guild') + '</div>')
            return app_page(platform, _page(body, 'Raid'), 'raid', badges={'activity': 1}, who='hunter')
        else:  # out
            body = (PageHeader('Raid with a guild', meta=['Sign up for your guild&rsquo;s raids from your phone.'])
                    + '<div class="vstack">' + Button('Sign in with Battle.net', primary=True) + '</div>')
            badges = {}
        return app_page(platform, _page(body, 'Raid'), 'raid', badges=badges, who=None if state == 'out' else 'warrior')
    return b


# ------------------------------------------------------------------------------------------------ Gear (the tab)
# The Fury Warrior 60 raid list (data/builds/1.60.1.70790/bis/warrior-fury.json), compared with what Thandrel wears.
DOLL = [  # slot label, pick, icon, quality, source, gain or None when worn
    ('Head', 'Lionheart Helm', 'inv_helmet_36', 4, 'Blacksmithing', '+14.5'),
    ('Neck', 'Skibi&rsquo;s Pendant', 'inv_jewelry_amulet_07', 3, 'World drop', None),
    ('Shoulder', 'Theramore Pauldrons', 'inv_shoulder_06', 4, 'Theramore Expeditionary Force', None),
    ('Back', 'Howler&rsquo;s Furs', 'inv_misc_cape_10', 3, 'Pix Xizzix, vendor', '+2.4'),
    ('Chest', 'Timbermaw Tunic', 'inv_chest_fur', 4, 'Leatherworking', '+5.3'),
    ('Wrist', 'Forest Stalker&rsquo;s Bracers', 'inv_bracer_12', 4, 'Silverwing Sentinels', '+0.1'),
    ('Hands', 'Theramore Gauntlets', 'inv_gauntlets_29', 4, 'Theramore Expeditionary Force', '+1.0'),
    ('Waist', 'Valiant&rsquo;s Waistguard', 'inv_belt_33', 4, 'Pix Xizzix, vendor', None),
    ('Legs', 'Sentinel&rsquo;s Plate Legguards', 'inv_pants_plate_20', 4, 'Illiyana Moonblaze, vendor', '+4.7'),
    ('Feet', 'Scalegut Treaders', 'inv_boots_chain_11', 3, 'Leatherworking', '+0.9'),
    ('Finger', 'Don Julio&rsquo;s Band', 'inv_jewelry_ring_35', 4, 'Stormpike Guard', '+6.8'),
    ('Finger', 'Cutthroat&rsquo;s Signet', 'inv_jewelry_ring_01', 3, 'Pix Xizzix, vendor', '+0.8'),
    ('Trinket', 'Hand of Justice', 'inv_jewelry_talisman_01', 3, 'Blackrock Depths: Emperor Dagran Thaurissan', None),
    ('Trinket', 'Molten Heart of the Mountain', 'inv_misc_gem_ruby_01', 3, 'Enchanting', '+3.2'),
    ('Main hand', 'Ironwood Blade', 'inv_sword_07', 4, 'Blacksmithing', '+5.9'),
    ('Off hand', 'Greenhammer', 'inv_mace_1h_raidhyjalc60_d_01', 4, 'Blacksmithing', '+9.7'),
    ('Ranged', 'Blackcrow', 'inv_weapon_crossbow_04', 3, 'Blackrock Spire: Shadow Hunter Vosh&rsquo;gajin, 8%', '+1.9'),
]


def gear(state: str):
    """Gear (owner, 2026-11-08): your gear against your band's list. The summary line names the list (your spec, band and
    faction, preselected; Change opens the picker); then where to go tonight; then every slot (Home holds the top three)."""
    def b(platform: str) -> str:
        if state == 'out':
            body = (PageHeader('Best in slot for your class', meta=['Pick a class, then a spec and a level range.'])
                    + Section('Classes', Rows([Row(c.capitalize(), '', crest_img(c, 40), title_color=CLASS_COLOR[c])
                                               for c in ('warrior', 'paladin', 'hunter', 'rogue', 'priest', 'shaman', 'mage', 'warlock', 'druid')], open_=True)))
            return app_page(platform, _page(body, 'Gear'), 'gear', who=None)
        summary = Rows([Row('Fury Warrior &middot; 60 raid &middot; Alliance', '4 of 17 slots on the list', trail='<a class="link" href="#">Change</a>', link=False)], open_=True)
        where = Rows([Row('Blackrock Spire', 'Blackcrow &middot; Shadow Hunter Vosh&rsquo;gajin, 8%', icon('inv_weapon_crossbow_04', 40, 'q3')),
                      Row('Blacksmithing', 'Lionheart Helm, Greenhammer, Ironwood Blade &middot; ask a blacksmith', icon('inv_helmet_36', 40, 'q4')),
                      Row('Stormpike Guard', 'Don Julio&rsquo;s Band, needs revered', icon('inv_jewelry_ring_35', 40, 'q4'))], open_=True)
        cells = [ItemRow(name, q, icon(ic, 40, f'q{q}'), f'{slot} &middot; {src}', gain=('Worn', '', '') if g is None else (g, 'DPS', ''))
                 for slot, name, ic, q, src, g in DOLL]
        body = ('<div data-region="controls">' + summary + '</div>'
                + Section('Where to go tonight', where)
                + Section('Every slot', '<div data-region="slots">' + ''.join(cells) + '</div>'))
        return app_page(platform, _page(body, 'Gear'), 'gear', badges=MEMBER)
    return b


# ------------------------------------------------------------------------------------------------ Activity (the tab)
ACTIVITY = [
    ('Today', [('ability_warrior_innerrage', 'Kaelra passed you in the guild sims', '1,012.3 DPS, you are #7 of 14', 'Guild ladder', '8 h ago', True),
               ('inv_axe_09', 'You got Deathbringer', 'Onyxia &middot; main hand, +12.6 DPS', 'Gear and sims', '9 h ago', True),
               (ONY, 'Onyxia&rsquo;s Lair report ready', '87th percentile, your best on Onyxia so far &middot; the coach has 2 notes', 'Logs', '9 h ago', True)]),
    ('Yesterday', [(ONY, 'You&rsquo;re In for Onyxia', AS_ME_CAP + ' &middot; 8:00 pm CST', 'Raids', '6:02 pm', False),
                   ('inv_misc_pocketwatch_01', 'Sim done: 948.6 DPS &plusmn;6', '+21 vs your last', 'Gear and sims', '2:40 pm', False),
                   (ONY, 'Healers short for Onyxia tonight', '3 of 8 missing', 'Raids', '1:15 pm', False)]),
    ('This week', [('inv_misc_book_09', 'Fury Warrior moved to S', 'Raid tiers, 60', 'Tiers', 'Tue', False),
                   ('inv_misc_book_09', 'Fury Warrior BiS changed', 'A new pick for Hands: Theramore Gauntlets', 'Gear and sims', 'Tue', False)]),
]
GROUPS = ['All', 'Raids', 'Gear and sims', 'Logs', 'Guild ladder', 'Tiers']


def activity(state: str):
    """Activity (owner, 2026-11-08): every notification event, newest first, read state, the six groups as filters (All
    preselected); every row opens its answer. The push inbox: a player with pushes off sees the same list."""
    def b(platform: str) -> str:
        if state == 'out':
            body = (PageHeader('Activity', meta=['What changed for your characters: loot, sims, reports, raids.'])
                    + '<div class="vstack">' + Button('Sign in with Battle.net', primary=True) + '</div>')
            return app_page(platform, _page(body, 'Activity'), 'activity', who=None)
        chips = '<div class="chiprow" data-region="controls">' + ''.join(Button(g, primary=(g == 'All'), small=True) for g in GROUPS) + '</div>'
        if state == 'empty':
            body = (PageHeader('Activity') + chips
                    + '<p class="t-small" style="margin-top:16px">Nothing yet. Loot, sims, reports and your guild&rsquo;s raids arrive here as they happen.</p>')
            return app_page(platform, _page(body, 'Activity'), 'activity')
        secs = ''
        for when, rows in ACTIVITY:
            secs += Section(when, Rows([_act_row(i, t, f'{s}<span class="when">{at}</span>', u)
                                        for i, t, s, g, at, u in rows], open_=True), action='Mark all read' if when == 'Today' else '')
        body = PageHeader('Activity') + chips + secs
        return app_page(platform, _page(body, 'Activity'), 'activity', badges={'raid': 0, 'activity': 3})
    return b


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
    """F3: healers short by 3, two healers unanswered and one on the bench, so the page says so and mentions the healers
    first; Mark In confirms on the row with Undo; Message opens a Discord direct message."""
    acts = lambda: ('<div class="btns" style="grid-template-columns:repeat(3,1fr)">' + Button('Mark In', small=True) + Button('Bench', small=True)
                    + Button('Discord DM', small=True) + '</div>')
    rows = [Row(n, s, crest_img(c, 40), title_color=CLASS_COLOR[c], link=False,
                foot=('<p class="t-small">Marked In &middot; <a class="link" href="#">Undo</a></p>' if n == 'Grunvald' else acts()))
            for n, c, s in UNANSWERED]
    bench = Rows([Row('Aurelith', 'Holy Paladin &middot; on the bench', crest_img('paladin', 40), title_color=CLASS_COLOR['paladin'], link=False,
                      trail=Button('Mark In', small=True))], open_=True)
    preview = (f'<p class="preview">@Mirethil @Selka: {RAID} is {WHEN} and we need healers. '
               'Answer In, Late, Bench or Out in the app or on the raid page.</p>')
    chips = '<div class="chiprow" data-region="controls">' + ''.join(
        Button(x, primary=(x.startswith('Not answered')), small=True) for x in ('Not answered 4', 'In 19', 'Late 3', 'Bench 2', 'Out 3', 'All 31')) + '</div>'
    body = (PageHeader(RAID, 'var(--strong)', icon(ONY, 56), ['Olympus XXVII', WHEN, 'healers short by 3'])
            + chips
            + '<div class="vstack">' + Button('Mention the 2 healers in Discord', primary=True) + Button('Mention all 4') + '</div>' + preview
            + Section('Not answered', Rows(rows, open_=True), meta='4 of 31, healers first')
            + Section('Healers on the bench', bench + '<p class="t-small stampl">If Mirethil, Selka and Aurelith are In, healers are full.</p>')
            + '<p class="t-small" style="margin-top:16px">Creating and editing raids stays on the site. <a class="link" href="#">Edit on the site &rsaquo;</a></p>')
    return app_page(platform, _page(body, RAID), 'raid', back='Raid', title=RAID, share=True, badges=OFFICER)


def applicant(state: str):
    """F3, review M3: the officer sees the note, the history, attunement, enchants and the gear gap before Accept; after
    Accept the invite is in game, so the page hands over the /ginvite line."""
    def b(platform: str) -> str:
        ready = Rows([Row('Attuned to Onyxia&rsquo;s Lair', 'Drakefire Amulet', link=False, slim=True),
                      Row('Enchants', '6 of 8 slots &middot; cloak and boots missing', link=False, slim=True),
                      Row('Best in slot gap', '9 pieces from the Holy Priest 60 raid list', link=False, slim=True)], open_=True)
        history = Rows([Row('Dawnbreak', 'Officer &middot; 14 raids &middot; left 2 Nov', link=False),
                        Row('Ashen Vow', 'Member &middot; 6 raids &middot; left 18 Oct', link=False)], open_=True)
        head = PageHeader('Mirelle', CLASS_COLOR['priest'], crest_img('priest', 56), ['Holy Priest &middot; 60', 'Alliance Dwarf', 'applied 2 min ago'])
        if state == 'accepted':
            body = (head + '<p class="t-lead" style="margin-top:12px">Accepted into Olympus XXVII. Invite her in game:</p>'
                    + '<div class="codeline"><span>/ginvite Mirelle</span>' + Button('Copy', small=True) + '</div>'
                    + '<p class="t-small" style="margin-top:12px">She shows on the roster once the addon or Battle.net sees her in the guild.</p>')
        else:
            body = (head + Figures([('74', 'Forever Score'), ('612', 'Sim', 'HPS'), ('63', 'Item level')])
                    + Section('Raid ready', ready) + Section('Guild history', history)
                    + Section('Their note', '<p class="t-prose" style="padding-top:8px">Holy main since launch. Wednesday and Sunday nights work for me.</p>')
                    + '<div class="sticky">' + Button('Accept into guild', primary=True) + Button('Decline') + '</div>'
                    + '<p class="t-small" style="margin-top:8px">Her raid role comes from her spec when she answers a raid.</p>')
        return app_page(platform, _page(body, 'Mirelle'), 'raid', back='Raid', title='Mirelle', badges=OFFICER)
    return b


# ------------------------------------------------------------------------------------------------ F8: first run
def pick_main(platform: str) -> str:
    rows = [Row(n, s, crest_img(c, 40), Tag('Main') if i == 0 else '', CLASS_COLOR[c], selected=(i == 0)) for i, (n, c, s, *_) in enumerate(CHARACTERS)]
    body = (PageHeader('Pick your main', meta=['From Battle.net, level 10 and up.'])
            + '<p class="t-small" style="margin-top:8px">Home answers for your main, and your raid sign-ups are for it. '
            'Switch any time from your portrait, top left.</p>'
            + Section('Your characters', Rows(rows, open_=True)) + '<div class="vstack">' + Button('Continue with Thandrel', primary=True) + '</div>')
    return app_page(platform, _page(body, 'Pick your main'), None, navbar=False)


def permission(guild: str):
    """F8 step 3: `You're in` on the small line, the guild alone as the title (coordinator ruling 2026-11-08), so no
    guild name's length decides how the head breaks: a long one wraps over two balanced lines on both engines."""
    def b(platform: str) -> str:
        body = ('<div class="center"><img class="seal" src="../../../web/public/seal-400.png" alt="">'
                + '<p class="t-small eyebrow">You&rsquo;re in</p>' + PageHeader(guild)
                + '<p class="t-small">Get raid sign-ups here. Answer In, Late, Bench or Out from the notification, without opening the app.</p></div>'
                + '<div class="vstack">' + Button('Turn on notifications', primary=True) + Button('Not now') + '</div>')
        return app_page(platform, _page(body, 'Notifications'), None, navbar=False)
    return b


def notification_settings(platform: str) -> str:
    """The groups (F4), each switch on its title row (review S12); Raids and Guild ladder show to a guild member, Officer to an officer."""
    groups = [('Raids', 'Posted, moved, your answer due, raid night', True), ('Gear and sims', 'An upgrade dropped, a sim finished, BiS changed', True),
              ('Logs', 'A report or the coach is ready, a new best', True), ('Guild ladder', 'A guildmate passed you', False),
              ('Tiers', 'Your spec moved this week', False)]  # Officer shows to officers only; Thandrel is a member
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
               'Requires Level 60</div><div style="height:8px"></div>')
    where = Rows([Row('Onyxia, Onyxia&rsquo;s Lair', '8% drop chance &middot; raid &middot; Wed with Olympus XXVII', icon(ONY, 40), link=False)], open_=True)
    runners = Rows([Row('Vis&rsquo;kag the Bloodletter', 'Onyxia &middot; 8% &middot; &minus;1.8 DPS', icon('inv_sword_18', 40, 'q4'),
                        title_color='#b866f5', link=False)], open_=True)
    needs = Rows([Row('Grunvald', 'Protection Warrior &middot; needs it, also on his list', crest_img('warrior', 40), title_color=WARRIOR, link=False),
                  Row('Torvak', 'Enhancement Shaman &middot; has it', crest_img('shaman', 40), title_color=CLASS_COLOR['shaman'], link=False)], open_=True)
    sheet = ('<div class="dim" style="height:120px"></div><div class="sheetp" role="dialog" aria-label="Deathbringer"><span class="sheet-grab"></span>'
             + ItemRow('Deathbringer', 4, icon(DEATHBRINGER[1], 40, 'q4'), 'Main hand', gain=('+12.6', 'DPS', 'over your worn'))
             + tooltip + Section('Where to get it', where) + Section('Runner-up', runners) + Section('Who in the guild needs it', needs) + '</div>')
    html = build_page('Deathbringer', sheet + f'<style>{FLOWS_CSS}</style>', note='Item sheet mock.')
    return app_page(platform, html, 'gear', navbar=False)


SCREENS = {
    'menu': menu('in'),
    'menu-one': menu('one'),
    'menu-low': menu('low'),
    'menu-out': menu('out'),
    **{k: v() for k, v in HOME.items()},
    'raid': raid('member'),
    'raid-officer': raid('officer'),
    'raid-live': raid('live'),
    'raid-none': raid('none'),
    'raid-no-raid': raid('noraid'),
    'raid-out': raid('out'),
    'raid-leveler': raid('leveler'),
    'gear': gear('in'),
    'gear-out': gear('out'),
    'activity': activity('in'),
    'activity-empty': activity('empty'),
    'activity-out': activity('out'),
    'notifications': notifications,
    'notification-settings': notification_settings,
    'raid-unanswered': raid_unanswered,
    'raid-applicant': applicant('review'),
    'raid-applicant-accepted': applicant('accepted'),
    'first-pick-main': pick_main,
    'first-permission': permission('Olympus XXVII'),
    'first-permission-long': permission('Ranchers of Thunderbluff'),
    'item-sheet': item_sheet,
}
