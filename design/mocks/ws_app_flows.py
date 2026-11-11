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
from ws_app_frame import AVATAR, BADGE, DEFAULT, FACTION_MARK, FACTION_RING, app_page, glyph, portrait
from ws_lib import QCOL, region_slug
from ws_prims import (Button, Countdown, Figures, ItemRow, PageHeader, Row, Rows, Section, Skeleton, Switch,
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
MEMBER, OFFICER, ANSWERED = {'guild': 1, 'activity': 3}, {'guild': 6, 'activity': 3}, {'activity': 3}

FLOWS_CSS = '''
.ap{padding:8px 16px 24px}
.r5{display:none}
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


# Round 5 on Home (lanes M1a and M1i; spec 2026-11-09-directions.md Mobile, mocks directions/png/m-<platform>-home.png): the
# hero behind the bar and the head, the blocks as panels on the ground, 16 under the hero, 14 at the sides and 20 apart;
# the raid card titles the zone and puts its answers in one segmented pill on Android, four 44 pt buttons on iOS
# (the closing [data-platform=ios] rules). The words stay the app-flows mock's; only the geometry is round 5's.
HOME_R5_CSS = '''
.ap.home{padding:0 14px 22px}
.ap.home .r5{display:block}
.ap.home .band-r5{margin:0 -14px;width:auto;padding:6px 16px 18px;background:#071d2e}
.ap.home .ph-wrap,.ap.home .head{display:block;padding:0}
.ap.home .ph-title{gap:12px}.ap.home .ph-title>div{flex:1}
.ap.home .t-title{font-size:30px;line-height:34px;letter-spacing:0}
.ap.home .ph-meta{display:block;margin-top:0;line-height:20px}
.ap.home .ph-meta .t-small{font-size:14px;line-height:20px}
.ap.home .ph-meta .t-small+.t-small::before{content:" \\00b7 "}
.ap.home>.blk,.ap.home>section{margin-top:20px}
.ap.home>.band-r5+*{margin-top:16px}
.ap.home .blk{background:#170e09;border-radius:10px;padding:8px 16px}
.ap.home .blk[data-region=section-loading]{background:none;padding:0}
.ap.home .blk[data-region=section-loading]>div:first-child{margin-top:0!important}
.ap.home .blk[data-region=section-loading] .skel{display:block;width:100%!important}
.ap.home .blk[data-region=section-loading]>div:first-child .skel{height:365px!important}
.ap.home .blk .row .trail .btn{min-height:48px}
.ap.home .blk:has(.cdn){padding:18px 18px 16px}
.ap.home .cdn{padding:0;border:0;background:none;border-radius:0}
.ap.home .cdn:has(.clk) .eb,.ap.home .cdn:has(.clk) .as{display:none}
.ap.home .cdn .eb{font:600 15px/22px var(--font-body);color:var(--strong)}
.ap.home .cdn .rw{margin-top:8px;gap:8px}
.ap.home .cdn .as{flex-basis:100%}
.ap.home .cdn .as p+p{margin-top:8px!important}
.ap.home .r5z{font:400 24px/30px var(--font-title);margin:0}
.ap.home .r5p{font-size:13px;line-height:20px;margin:0}
.ap.home .cdn .clk{gap:18px}.ap.home .cdn .clk>span{align-items:flex-start}
.ap.home .cdn .clk b{font:600 36px/38px var(--font-body)}
.ap.home .cdn .clk>span>span{font-size:12px;line-height:16px}
.ap.home .blk .me{margin-top:12px;line-height:20px}
.ap.home .blk .btns{margin-top:8px;gap:0;border:1px solid var(--line);border-radius:24px;overflow:hidden}
.ap.home .blk .btns .btn{height:48px;min-height:48px;border:0;border-radius:0}
.ap.home .blk .me+.chips{margin-top:0}
.ap.home .blk .tally{margin-top:14px}
.ap.home .blk:has(.cdn) .tally+.rows{margin-top:14px;border-top:1px solid var(--line)}
.ap.home .blk:has(.cdn) .rows+.rows{margin-top:4px}
.ap.home .blk:not(:has(.cdn)) .rows{margin-top:0}
.ap.home .blk:not(:has(.cdn)) .btns{margin:4px 0 8px}
.ap.home .blk .vstack{margin:12px 0}
.ap.home .blk>p.t-small:first-child{margin-top:8px}
.ap.home section .sec-h{padding-bottom:0;border-bottom:0;margin-bottom:8px;align-items:center}
.ap.home section .sec-h h2{font:400 21px/26px var(--font-title)}
.ap.home section>.sec-h+*{margin-top:0}
.ap.home section>.rows,.ap.home section>.irow-box{background:#170e09;padding:0 16px}
.ap.home section>.sec-h+*{padding-top:8px;border-radius:10px 10px 0 0}
.ap.home section>:last-child{padding-bottom:8px;border-radius:0 0 10px 10px}
.ap.home section>.sec-h+:last-child{border-radius:10px}
.ap.home section>.irow-box+*{border-top:1px solid var(--line)}
.ap.home .irow{padding-inline:0}
/* a row's lines as the Pixel 9 rounds them (420 dpi: a 20 sp line is 54 px, 20.57 dp; 18 sp is 48 px, 18.29 dp) */
.ap.home section .row .ttl{line-height:20.57px}.ap.home section .row .sub{line-height:18.29px}
/* round 2: iOS answers with its Segmented (brand B1): one connected group, hairline dividers, the chosen one teal, 44 pt */
[data-platform=ios] .ap.home .blk .btns{border-color:#5a5d70;border-radius:6px;background:#170e09}
[data-platform=ios] .ap.home .blk .btns .btn{height:44px;min-height:44px;padding:0 4px;font:600 13px/18px var(--font-body);color:var(--text)}
[data-platform=ios] .ap.home .blk .btns .btn+.btn{border-left:1px solid #5a5d70}
[data-platform=ios] .ap.home .blk .btns .btn.primary{font-weight:700;color:#fff;background:#0f6064;box-shadow:inset 0 0 0 1px #408076}
/* round 2 (review backfill): the status line under the answers; the upgrades' synced stamp; the raid line's tally */
.ap.home .chips .btn.primary{background:#0f6064;border-color:#408076;color:#fff}
.ap.home .btn.primary{background:#0f6064;border-color:#408076;color:#fff;box-shadow:none}
.ap.home .row .trail .btn,.ap.home .vstack .btn:not(.primary){background:#114f72;border-color:#1d6a96;color:#fff}
.ap.home .stl{margin:8px 0 0;font:400 13px/20px var(--font-body);color:var(--kill)}
.ap.home section>.stmp{margin:-8px 0 8px;font:400 14px/20px var(--font-body);color:var(--muted)}
[data-platform=ios] .ap.home section>.stmp{margin-top:-12px}
.ap.home section>.stmp+*{padding-top:8px;border-radius:10px 10px 0 0}
/* round 3: New for you's ages on their own line; the compact tally (iOS 14/20 in the text colour, Android 13 muted) */
.ap.home .row .sub.age{margin-top:0}
[data-platform=ios] .ap.home .tline{font:400 14px/20px var(--font-body);color:var(--text);margin:0 0 8px}
[data-platform=ios] .ap.home .tline .ok{color:#fff}
[data-platform=ios] .ap.home .blk .btns .btn{height:36px;min-height:36px}
[data-platform=ios] .ap.home .blk:not(:has(.cdn)) .btns{margin:0}
.ap.home .tline{margin:0 0 4px;font:400 13px/18.29px var(--font-body);color:var(--muted)}
.ap.home .tline .ok{color:var(--kill)}.ap.home .tline .lo{color:var(--wipe)}
.ap.home .blk:not(:has(.cdn)) .tally{margin:0 0 8px}
[data-platform=ios] .ap.home .blk .row .trail .btn{min-height:44px}
'''


def _page(body: str, title: str, home: bool = False) -> str:
    cls, css = ('ap home', FLOWS_CSS + HOME_R5_CSS) if home else ('ap', FLOWS_CSS)
    return build_page(title, f'<div class="{cls}">{body}</div><style>{css}</style>',
                      note='Native app mock (design/specs/2026-11-08-app-flows.md).')


def _signup_buttons(selected: str = '') -> str:
    return '<div class="btns" role="group" aria-label="Your answer">' + ''.join(
        Button(x, primary=(x == selected), small=True) for x in ('In', 'Late', 'Bench', 'Out')) + '</div>'


def _late_chips(chosen: str = '') -> str:
    """15, 30, 60 min; the minutes already said selected (they stay open after a pick, round 2)."""
    return '<div class="chips" role="group" aria-label="How late">' + ''.join(
        Button(x, primary=(x == chosen), small=True) for x in ('15 min', '30 min', '60 min')) + '</div>'


def _raid_card(officer: bool = False, answered: str = '', late: bool = False) -> str:
    """The next raid (F1): the countdown, who you answer as, the four answers (the current one selected, so a change is one
    tap), who is coming, what is not ready; an officer's open items with their one action. late: Late's minutes are open."""
    cd = Countdown(f'Olympus XXVII &middot; {RAID} tonight', [('02', 'hours'), ('48', 'min')],
                   f'<p class="t-small">{WHEN}</p>', 'guild')
    answered = 'Late' if late else answered
    # round 3 (AS5): once answered the answers are the state; the line names only the character
    me = f'<p class="t-small me">{AS_ME_CAP if answered else "You have not answered, " + AS_ME}</p>'
    # late: Late with 30 minutes just sent, the minutes selected and the status line saying so
    how_late = '<p class="t-small me">How late?</p>' + _late_chips('30 min') + '<p class="stl">Saved.</p>' if late else ''
    ready = Rows([Row('2 things before you pull', 'No flask in your bags; boots not enchanted &middot; from your last log-out, 9 h ago', slim=True)], open_=True)
    # round 5 (HOME_R5_CSS): the card titles the zone and puts the guild before the time
    r5 = f'<h2 class="r5 r5z">{RAID}</h2><p class="r5 r5p">Olympus XXVII &middot; {WHEN}</p>'
    return r5 + cd + me + _signup_buttons(answered) + how_late + Tally(ROLES) + ready + (_needs_you() if officer else '')


def _needs_you() -> str:
    """The officer's open items (F3): the unanswered with the short role named, the applicant; the badge's number."""
    return Rows([Row('5 have not answered, 2 of them healers', 'Healers short by 3', trail=Button('Ping on Discord', small=True), link=False),
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
# Round 2 (review backfill): the head's line names the guild rank; a Switch row ends with the item level the import has.
YOU_LINE = 'Fury Warrior &middot; 60 &middot; Raider'
ITEM_LEVEL = {'Lightvow': 61, 'Kaelor': 44, 'Felsnap': 31, 'Brewmist': 19, 'Ashmantle': 15, 'Pockets': 5}
NEWEST_DEVICE = {'ios': 'iPhone 17', 'android': 'Pixel 9'}
# The tools, in the order the owner ruled (2026-11-08); the ones a signed-out visitor can use carry True.
TOOLS = [('Guides', 'Your spec&rsquo;s guide, then every class', True), ('Tiers', 'Every spec, this week', True),
         ('Patch notes', 'The game and the site', True), ('Sims', 'Your server runs and saved sims', False),
         ('Logs and reports', 'Your reports and the coach', False), ('Planner', 'Your builds and the guide builds', True),
         ('Rankings', 'Your spec and your guild', True)]
ACCOUNT = [('Notifications', 'Gear and sims, guild ladder, tiers off'), ('Plan', 'Premium &middot; Founder price'), ('Devices', '3 signed in'),
           ('Offline data', '26.6 MB on this phone'), ('Settings', 'Sign out, delete account')]


# Round 5 on the profile drawer (lanes M5a, M5i; spec 2026-11-09-directions.md Mobile, mock directions/png/m-android-menu.png):
# You on the painted band (the portrait at 56, the name in the display face at 26 in the class colour, the line, Main, the
# stamp with Refresh), then on the drawer's panel Switch character, Tools as labels alone (owner ruling 2026-11-09) and
# Account, each under a cream display title at 16. 18 at the sides. iOS since lane M5i.
MENU_R5_CSS = """
.drawer.r5menu{padding:0}
.r5menu .you{padding:18px 18px 12px;background:#0c3a52;border-bottom:1px solid #000}
.r5menu .yt{display:flex;align-items:center;gap:16px}
.r5menu .yt h2{margin:0;font:400 26px/32px var(--font-title)}
.r5menu .yt .ln{display:block;font:400 13px/20.57px var(--font-body);color:var(--text)}
.r5menu .yt .tag{display:inline-block;margin-top:4px;padding:1px 8px;border:1px solid var(--gold);border-radius:4px;color:var(--gold);font:600 12px/16px var(--font-body)}
.r5menu .yt .sub{display:block;margin-top:4px;font:400 13px/20.57px var(--font-body);color:var(--text)}
.r5menu .stamp{display:flex;align-items:center;min-height:48px;margin:4px 0 0;font:400 12px/16px var(--font-body);color:var(--text)}
.r5menu .stamp a{font-weight:600;color:var(--gold)}
.r5menu .wide{display:flex;align-items:center;justify-content:center;height:48px;margin:16px 0 4px;border-radius:4px;background:#0f6064;border:1px solid #408076;
  color:#fff;font:700 15px/20px var(--font-body)}
.r5menu .bd{padding:0 18px}
.r5menu section{margin:0}
.r5menu h3{margin:0;padding:14px 0 6px;font:400 16px/22.1px var(--font-title);color:#f3eee2}
.r5menu .row{display:flex;flex-direction:column;justify-content:center;box-sizing:border-box;min-height:48px;padding:7px 0;border-bottom:1px solid var(--line)}
.r5menu .row:last-child{border-bottom:0}.r5menu .row.ch{flex-direction:row;align-items:center;justify-content:flex-start;gap:12px}
.r5menu b.t{display:block;font:700 15px/20.57px var(--font-body);color:#fff}
.r5menu .s{display:block;font:400 13px/18.29px var(--font-body);color:var(--muted)}
.r5menu .fold{display:flex;align-items:center;min-height:48px;font:600 15px/20px var(--font-body);color:var(--gold)}
.r5menu .foot{margin:24px 0;display:flex;flex-direction:column;gap:4px}
/* iOS (lane M5i, m-ios-menu): the band runs under the status bar, You 18 under it; a label row is 44 */
[data-platform=ios] .r5menu .you{padding-top:calc(var(--status) + 18px)}
[data-platform=ios] .r5menu .row:not(:has(.s)){min-height:44px}
"""


def _r5_portrait(faction: str, avatar: str | None, cls: str, px: int) -> str:
    """portrait() in round 5: a character with no Battle.net portrait shows its class crest in the faction ring."""
    if avatar:
        return portrait(faction, avatar, px)
    b = round(px * BADGE)
    crest = crest_img(cls, px).replace('class="crest"', 'class="pf-img"', 1)
    badge = f'<img class="pf-badge" src="{FACTION_MARK.format(faction=faction)}" alt="" style="width:{b}px;height:{b}px">'
    return f'<span class="pf" style="width:{px}px;height:{px}px;--ring:{FACTION_RING[faction]}">{crest}{badge}</span>'


def _switch_line(name: str, line: str) -> str:
    level = ITEM_LEVEL.get(name)
    return f'{line} &middot; item level {level}' if level else line


def _menu_r5(state: str, platform: str) -> str:
    """The drawer's body in round 5 (MENU_R5_CSS); the characters, tools and account rows are menu()'s."""
    def section(title: str, rows: str) -> str:
        return f'<section data-region="{region_slug(title)}"><h3>{title}</h3>{rows}</section>'

    def row(title: str, sub: str = '') -> str:
        return f'<div class="row"><b class="t">{title}</b>' + (f'<span class="s">{sub}</span>' if sub else '') + '</div>'

    if state == 'out':
        you = ('<div class="you"><div data-region="head"><div data-region="head-title"><div class="yt" style="display:block">'
               '<h2 style="color:#f3eee2">Forever Sixty</h2><span class="sub" data-region="head-meta">Sign in and your characters arrive with their gear, '
               'talents and guild.</span></div></div><a class="btn wide" href="#">Sign in with Battle.net</a></div></div>')
        body = section('Tools', ''.join(row(n) for n, _, out in TOOLS if out))
    else:
        me = CHARACTERS[0]
        you = (f'<div class="you"><div data-region="head"><div class="yt" data-region="head-title">{portrait(me[5], me[4], 56)}'
               f'<div style="flex:1"><h2 style="color:{CLASS_COLOR[me[1]]}">{me[0]}</h2><div data-region="head-meta"><span class="ln">{YOU_LINE}</span>'
               '<span class="tag">Main</span></div></div></div>'
               '<p class="stamp">Updated 2 h ago from Battle.net&nbsp;&middot;&nbsp;<a href="#">Refresh</a></p></div></div>')
        others = CHARACTERS[1:] if state != 'one' else []
        chars = [f'<div class="row ch">{_r5_portrait(f, av, c, 40)}<div><b class="t" style="color:{CLASS_COLOR[c]}">{n}</b><span class="s">{_switch_line(n, s)}</span></div></div>'
                 for n, c, s, _, av, f in others]
        if state == 'low':
            chars += [f'<div class="row ch">{_r5_portrait("alliance", None, c, 40)}<div><b class="t" style="color:{CLASS_COLOR[c]}">{n}</b><span class="s">{_switch_line(n, s)}</span></div></div>'
                      for n, c, s in LOW]
        fold = (f'<a class="fold" href="#">{"Hide low-level characters" if state == "low" else "Show low-level characters (2)"}</a>' if others else '')
        body = ((section('Switch character', ''.join(chars) + fold) if others else '')
                + section('Tools', ''.join(row(n) for n, _, _ in TOOLS))
                + section('Account', ''.join(row(n, f'{s} &middot; newest {NEWEST_DEVICE[platform]}' if n == 'Devices' else s) for n, s in ACCOUNT)))
    foot = ('<div class="foot"><span class="t-small">A fan-made player tool. Not affiliated with Blizzard.</span>'
            '<span class="t-small">&copy; 2026 Commish LLC d/b/a Forever Sixty</span>'
            '<span class="t-small">Version 0.1.0 (7) &middot; data 1.60.1.70790</span>'
            '<span class="t-small"><a class="link" href="#">Help</a> &middot; <a class="link" href="#">Privacy</a></span></div>')
    return you + f'<div class="bd">{body}{foot}</div>'


def menu(state: str):
    """state: in (6 characters), one (just the main), low (the low-level fold open), out (signed out)."""
    def b(platform: str) -> str:
        html = build_page('Menu', f'<div class="drawer-wrap"><div class="drawer r5menu" role="dialog" aria-label="Menu">{_menu_r5(state, platform)}</div>'
                          f'<div class="dim" aria-hidden="true"></div></div><style>{FLOWS_CSS}{MENU_R5_CSS}</style>', note='Profile menu mock, round 5.')
        return app_page(platform, html, None, navbar=False, who=None if state == 'out' else 'warrior')
    return b


# ------------------------------------------------------------------------------------------------ Home (section 4)
def _head(leveler: bool = False) -> str:
    """One head on every Home state: the class crest, the name in the class colour and the fact line. The profile button
    above it is the portrait, a different image, so the character is never drawn twice the same way (owner, 2026-11-08)."""
    if leveler:
        return PageHeader('Felsnap', HUNTER, crest_img('hunter', 56), ['Level 34 Beast Mastery Hunter', '6 levels to the 40 list'])
    return PageHeader('Thandrel', WARRIOR, crest_img('warrior', 56), ['Level 60 Fury Warrior', 'Olympus XXVII'])


# Next upgrades (round 2, review backfill): the figures in the section's meta (Home's head carries none, owner ruling), the
# synced stamp (iOS at the top of the panel, Android at its foot), the worn item a row replaces, Best run as a row with its
# drop's icon (iOS titles the place, Android the item). The rows are the gear the Gear tab can reach (Greenhammer waits for
# the first raids), as the Android parity fixture has them.
UPGRADES = {
    'raid': ('12 upgrades &middot; +47.5 DPS', 4, [
        ('Lionheart Helm', 'inv_helmet_36', 'Head &middot; Blacksmithing &middot; replaces Helm of Valor', '+14.5'),
        ('Don Julio&rsquo;s Band', 'inv_jewelry_ring_35', 'Finger &middot; Stormpike Guard, needs revered', '+6.8'),
        ('Timbermaw Tunic', 'inv_chest_fur', 'Chest &middot; Leatherworking', '+5.3')],
        ('Blackrock Spire', 'Blackcrow', 3, 'inv_weapon_crossbow_02', '8%', '+1.9')),
    'leveler': ('6 upgrades &middot; +3.1 DPS', 3, [
        ('Felbough Arm', 'inv_spear_08', 'Main hand &middot; Fel Ancient, Ashenvale', '+1.3'),
        ('Cartilage Shapers', 'inv_gauntlets_18', 'Hands &middot; quest, Lyn the Ignored', '+0.8'),
        ('Ring of Precision', 'inv_jewelry_ring_14', 'Finger &middot; Blackfathom Deeps', '+0.6')],
        ('Blackfathom Deeps', 'Ring of Precision', 3, 'inv_jewelry_ring_14', '', '+0.6')),
}
SYNCED = 'Synced 2 h ago from the addon'


def _best_run(platform: str, place: str, item: str, q: int, ic: str, chance: str, gain: str) -> str:
    """Best run as a row: iOS titles the dungeon (`Best run · Blackcrow · 8%`), Android the item (`Best run · Blackrock Spire · 8%`)."""
    ios = platform == 'ios'
    line = ' &middot; '.join(x for x in (f'Best run &middot; {item if ios else place}', chance) if x)
    row = ItemRow(place if ios else item, 1 if ios else q, icon(ic, 40, f'q{q}'), line, gain=(gain, 'DPS', ''))
    return row.replace(f'color:{QCOL[1]}', 'color:var(--strong)', 1) if ios else row


def _upgrades(platform: str, leveler: bool = False, tonight: bool = True) -> str:
    """Next upgrades (F5): three, each with its gain and where it comes from; Best run as a row (left out on a raid day: the
    raid is tonight)."""
    meta, quality, picks, run = UPGRADES['leveler' if leveler else 'raid']
    rows = ''.join(ItemRow(n, quality, icon(i, 40, f'q{quality}'), src, gain=(g, 'DPS', '')) for n, i, src, g in picks)
    if tonight:
        rows += _best_run(platform, *run)
    # round 3 (row 8): the stamp under the meta line, outside the panel, on both
    html = Section('Next upgrades', rows, action='Gear &rsaquo;', meta=meta)
    return html.replace('</div>', f'</div><p class="stmp">{SYNCED}</p>', 1)


# icon, title, body, when (round 3: Activity's buckets on their own line), unread
ACTIVITY_PREVIEW = {
    'day': [('ability_warrior_innerrage', 'Kaelra passed you in the guild sims', '1,012 DPS, you are #7 of 14', '1 h ago', True),
            ('inv_misc_pocketwatch_01', 'Sim done: 949 DPS &plusmn;6', '+21 vs your last', '3 h ago', True),
            ('inv_misc_book_09', 'Fury Warrior BiS changed', 'A new pick for Hands', '6:02 pm', False)],
    # the day after: Last night shows the loot and the report, so the preview starts below them
    'after': [('ability_warrior_innerrage', 'Kaelra passed you in the guild sims', '1,012 DPS, you are #7 of 14', '8 h ago', True),
              ('inv_misc_pocketwatch_01', 'Sim done: 949 DPS &plusmn;6', '+21 vs your last', '6:02 pm', False),
              ('inv_misc_book_09', 'Fury Warrior BiS changed', 'A new pick for Hands', 'Tue', False)],
    'live': [('inv_misc_head_dragon_black', 'Brightwen is Late, 15 min', 'Onyxia&rsquo;s Lair', '4 min ago', True),
             ('inv_misc_head_dragon_black', 'Selka is In', 'Restoration Druid', '20 min ago', True),
             ('inv_misc_head_dragon_black', 'Healers short for Onyxia tonight', '3 of 8 missing', '3 h ago', False)],
    'leveler': [('inv_misc_book_09', 'Level 34: your 30 to 39 list', 'Beast Mastery Hunter', '2 h ago', True),
                ('inv_misc_pocketwatch_01', 'Sim done: 58.4 DPS &plusmn;1', '+3.1 vs your last', '6:02 pm', False),
                ('inv_misc_book_09', 'Beast Mastery Hunter moved to A', 'Leveling tiers, 30 to 39', 'Sun', False)],
}


def _act_row(ic: str, t: str, s: str, when: str, unread: bool) -> str:
    """A feed row on Home: its body, then its when on a line of its own (round 3, row 3)."""
    sub = f'{s}</div><div class="sub age">{when}'
    return Row(t, sub, _act_lead(ic, t)).replace('class="row', f'class="act-row{" unread" if unread else ""} row', 1)


def _changed(kind: str, unread: int = 3) -> str:
    """New for you (F4, the casual's answer): the three newest Activity rows, unread ones marked; the link carries the
    unread count, the same number as the Activity badge; all of them are one tap away."""
    rows = ACTIVITY_PREVIEW[kind]
    link = f'All activity ({unread}) &rsaquo;' if unread else 'All activity &rsaquo;'
    return Section('New for you', Rows([_act_row(i, t, s, w, u) for i, t, s, w, u in rows], open_=True), action=link)


def _home(blocks, title: str, badges: dict, who: str = 'warrior', avatar=DEFAULT):
    """blocks: the body, or a function of the platform for the blocks the two apps draw apart (round 2)."""
    def b(platform: str) -> str:
        body = blocks(platform) if callable(blocks) else blocks
        return app_page(platform, _page(body, title, home=True), 'home', badges=badges, who=who, avatar=avatar)
    return b


def _tally_line(platform: str) -> str:
    """Who is coming in one line (S6, round 3 row 2): `Healers 5/8 · 3 short`, the short role in wipe red. iOS counts it as
    a row (its parity item), so the mock's iOS line carries the row class `sr`."""
    def part(label: str, have: int, want: int) -> str:
        short = f' &middot; {want - have} short' if have < want else ''
        return f'{label} <span class="{"ok" if have >= want else "lo"}">{have}/{want}{short}</span>'
    cls = 'tline sr' if platform == 'ios' else 'tline'
    return f'<p class="{cls}">' + ' &middot; '.join(part(label, have, want) for _, label, have, want in ROLES) + '</p>'


def _raid_line(platform: str, answered: str = '') -> str:
    """Any other day: the raid as a row, the role tally on one line, the answers (iOS 36 pt). Answered, the line keeps the
    counts and who it is for (round 3, AS5)."""
    sub = f'{COUNTS} &middot; {AS_ME_CAP}' if answered else f'{COUNTS} &middot; you have not answered, {AS_ME}'
    row = Rows([Row(f'{RAID} &middot; {WHEN}', sub, icon(ONY, 40))], open_=True)
    return _blk(row + _tally_line(platform) + _signup_buttons(answered))


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
    return _blk('<p class="t-small">Join a guild to sign up for its raids here.</p>'
                + '<div class="vstack">' + Button('Find a guild') + '</div>', 'section-raid-with-a-guild')


HOME = {
    'home-raid-day': lambda: _home(lambda p: _head() + _blk(_raid_card()) + _upgrades(p, tonight=False) + _changed('day'), 'Home', MEMBER),
    'home-raid-day-late': lambda: _home(lambda p: _head() + _blk(_raid_card(late=True)) + _upgrades(p, tonight=False) + _changed('day'), 'Home', ANSWERED),
    'home-raid-day-officer': lambda: _home(lambda p: _head() + _blk(_raid_card(officer=True)) + _upgrades(p, tonight=False) + _changed('day'), 'Home', OFFICER),
    'home-raid-live': lambda: _home(_head() + _raid_live() + _changed('live'), 'Home', ANSWERED),
    'home-after-raid': lambda: _home(lambda p: _head() + _last_night() + _upgrades(p) + _changed('after'), 'Home', ANSWERED),
    'home-60': lambda: _home(lambda p: _head() + _upgrades(p) + _raid_line(p) + _changed('day'), 'Home', MEMBER),
    # no avatar_url yet (a new character, a failed media fetch): the crest stands in on the profile button
    'home-60-noavatar': lambda: _home(lambda p: _head() + _upgrades(p) + _raid_line(p) + _changed('day'), 'Home', MEMBER, avatar=None),
    'home-60-answered': lambda: _home(lambda p: _head() + _upgrades(p) + _raid_line(p, 'In') + _changed('day'), 'Home', ANSWERED),
    'home-noguild': lambda: _home(lambda p: _head() + _upgrades(p) + _changed('day', 2) + _find_guild(), 'Home', {'activity': 2}),
    'home-34': lambda: _home(lambda p: _head(leveler=True) + _upgrades(p, leveler=True) + _changed('leveler', 1), 'Home', {'activity': 1}, who='hunter'),
    # loading: the head and the blocks at their final size, no shift when they fill
    'home-loading': lambda: _home(_head() + _blk(''.join(f'<div style="margin-top:12px">{Skeleton(361, 56)}</div>' for _ in range(5)), 'section-loading'),
                                  'Home', {}),
    'home-error': lambda: _home(_head() + _blk(Rows([Row('We could not load your home', 'Check your connection', trail=Button('Try again', small=True), link=False)],
                                                    open_=True), 'section-error'), 'Home', {}),
    # signed in, no character yet (F8): the one thing to do
    'home-new': lambda: _home(PageHeader('No characters yet', meta=['Log out in game, then check again.'])
                              + '<div class="vstack">' + Button('Check again', primary=True) + '</div>', 'Home', {}, who=None),
}


# Round 5 on the Raid tab (lanes M2a, M2i; spec 2026-11-09-directions.md Mobile, mocks directions/png/m-android-raid.png
# and m-android-signup.png): the blue band behind the bar, then on the ground the one-line raid head (the zone's icon, the
# raid, its time and the pull, who the answer is for, the answer chip) over the four answers as one 48 dp segmented pill,
# Late's minutes inline under them; Who is coming, Before you pull, Needs you and the guild's pages as panels with display
# titles. iOS since lane M2i, with 44 pt buttons. Line heights are the Pixel 9's rounding
# (420 dpi: 13/18/20/22/24 sp lines are 20.57/18.29/20.57/22.1/24.38 dp). Open Sans sets wider on the Pixel 9 than in
# Chromium, so the two lines it breaks there carry the break here (<br>) and the head's lines hold to one line as they do there.
RAID_R5_CSS = """
.ap.r5raid{padding:16px 14px 22px}
.ap.r5raid>*{margin-top:0}.ap.r5raid>*+*{margin-top:20px}
.r5raid .pn{background:#170e09;border-radius:10px;padding:8px 16px;box-shadow:0 0 0 1px #000}
.r5raid .sh{display:flex;align-items:center;gap:12px;height:26px;margin-bottom:8px}
.r5raid .sh h2{margin:0;font:400 21px/26px var(--font-title);color:#f3eee2}
.r5raid .sh .act{margin-left:auto;font:600 14px/24.38px var(--font-body);color:var(--gold)}
.r5raid .sh .meta{font:400 13px/20.57px var(--font-body);color:var(--muted)}
.r5raid .hd{padding:8px 0}
.r5raid .rl{display:flex;align-items:center;gap:12px}.r5raid .rl img{width:40px;height:40px;flex:none;border-radius:4px}
.r5raid .rl .tx{flex:1;min-width:0}.r5raid .rl .s{white-space:nowrap}
.r5raid b,.r5raid .t{display:block;font:700 15px/20.57px var(--font-body);color:#fff}
.r5raid .s{display:block;font:400 13px/18.29px var(--font-body);color:var(--muted)}
.r5raid .chip{flex:none;padding:3px 10px;border-radius:4px;background:#0f6064;border:1px solid #408076;color:#fff;font:700 13px/18.29px var(--font-body)}
.r5raid .btn{display:flex;align-items:center;justify-content:center;box-sizing:border-box;height:48px;min-height:48px;margin:0;padding:0 16px;
  border:1px solid var(--line);border-radius:4px;background:none;color:#fff;font:600 14px/20px var(--font-body)}
.r5raid .seg{display:grid;grid-template-columns:repeat(4,1fr);margin-top:12px;border:1px solid #5a5d70;border-radius:24px;overflow:hidden}
.r5raid .seg .btn{height:48px;min-height:48px;border:0;border-radius:0}.r5raid .seg .btn+.btn{border-left:1px solid #5a5d70}
.r5raid .seg .btn.on,.r5raid .lm .btn.on{background:#408076}
.r5raid .lm{display:flex;gap:8px;margin-top:8px}.r5raid .lm .btn{border-color:#5a5d70;border-radius:8px}
.r5raid .row{display:block;min-height:0;padding:10px 0;border-bottom:1px solid var(--line)}.r5raid .row:last-child{border-bottom:0}
.r5raid .row .in{display:flex;align-items:center;gap:12px}
.r5raid .row .tr{margin-top:8px;display:flex}
.r5raid .row .lk{font-weight:600;color:var(--gold)}
.r5raid .mk{display:flex;align-items:center;justify-content:center;width:22px;height:22px;flex:none;border-radius:50%;color:#0c0a09;font:700 13px/22px var(--font-body)}
.r5raid .mk.ok{background:#7fd48a}.r5raid .mk.miss{background:#f0736b}.r5raid .mk.warn{background:#ffec71}
.r5raid .ok{color:#7fd48a}.r5raid .lo{color:#f0736b}
.r5raid .stamp{padding:8px 0;font:400 13px/20.57px var(--font-body);color:var(--muted)}
.r5raid .note{margin:8px 0 0;font:400 13px/20.57px var(--font-body);color:var(--text)}
.r5raid .note.strong{color:#fff}
.r5raid .wide{width:100%;margin:12px 0}
.r5raid .wide.primary{background:#0f6064;border-color:#408076}
.r5raid .lvh{display:flex;align-items:center;gap:8px;margin-top:12px}.r5raid .lvh i{width:8px;height:8px;border-radius:50%;background:var(--kill)}
.r5raid .lvh span{font:600 15px/22.1px var(--font-body);color:#fff}
.r5raid .lvh+.note{margin:4px 0 8px}
.r5raid .rule{border-top:1px solid var(--line)}
/* round 2 (review backfill): the status line under the answers, the officers' notes, Who is coming's closing lines */
.r5raid .stl{margin:8px 0 0;font:400 13px/20px var(--font-body);color:var(--kill)}
.r5raid .notes{margin:0;font:400 13px/18.29px var(--font-body);color:var(--text);display:-webkit-box;-webkit-box-orient:vertical;-webkit-line-clamp:2;overflow:hidden}
.r5raid .notes b{font-weight:700;color:#fff;display:inline}
.r5raid .rfoot p{margin:0;padding:8px 0;font:400 13px/18.29px var(--font-body);color:var(--muted)}
.r5raid .rfoot p.warn{color:var(--wipe)}
.r5raid .row .split{font-weight:400;color:var(--muted)}
/* round 3: the notes close the head's card (16 under the answers on iOS, 8 on Android); Android asks How late? over the
   chips; the Covered line (a muted Missing, each short name in wipe red beside its 20 pt spell icon); the officers'
   not-synced line is a row that opens the roster */
.r5raid .pn .notes{margin-top:8px}
.r5raid .hl{margin:8px 0 0;font:400 14px/20px var(--font-body);color:var(--muted)}
.r5raid .cov{display:flex;flex-wrap:wrap;align-items:center;gap:4px;padding:8px 0;font:400 13px/18.29px var(--font-body);color:var(--muted)}
.r5raid .cov .cg{display:inline-flex;align-items:center;gap:4px;color:var(--wipe);white-space:nowrap}
.r5raid .cov .cg img,.r5raid .cov .cg .ci{display:block;width:20px;height:20px;border-radius:4px;background-size:cover}
.r5raid .cov .cg .n{color:var(--muted)}
.r5raid .rfoot .ns{display:flex;align-items:center;gap:4px;min-height:48px;padding:0;border:0;font:400 13px/18.29px var(--font-body);color:var(--text)}
.r5raid .rfoot .ns span{flex:1}.r5raid .rfoot .ns i{font-style:normal;color:var(--muted)}
[data-platform=ios] .r5raid .rfoot .ns{min-height:44px}
/* round 5: buttons from the tokens: the primary teal with its tealEdge and strong text, the secondary blue with blueEdge */
.r5raid .btn{background:#114f72;border-color:#1d6a96}
.r5raid .seg .btn,.r5raid .lm .btn{background:none}
.r5raid .seg .btn.on,.r5raid .lm .btn.on,.r5raid .wide.primary{background:#0f6064;border-color:#408076;color:#fff}
[data-platform=ios] .r5raid .cov{padding:0;gap:4px}
[data-platform=android] .r5raid .sh .act{line-height:18px}
[data-platform=ios] .r5raid .btn{height:44px;min-height:44px}
/* iOS's Segmented (brand B1): one connected group, hairline dividers, the chosen one teal with its edge, 44 pt */
[data-platform=ios] .r5raid .seg{border:1px solid #5a5d70;border-radius:6px;background:#170e09}
[data-platform=ios] .r5raid .seg .btn{height:44px;min-height:44px;padding:0 4px;font:600 13px/18px var(--font-body);color:var(--text)}
[data-platform=ios] .r5raid .seg .btn.on{font-weight:700;color:#fff;box-shadow:inset 0 0 0 1px #408076;background:#0f6064}
[data-platform=ios] .r5raid .pn .notes{margin-top:16px}
[data-platform=ios] .r5raid .rfoot{border-top:1px solid var(--line);padding:10px 0;display:flex;flex-direction:column;gap:8px}
[data-platform=ios] .r5raid .rfoot p{padding:0;line-height:18px}
[data-platform=ios] .r5raid .rfoot p.warn{color:var(--muted)}
"""

# Who is coming (S6, one tally model on both apps): `Healers 5 of 8 · 3 short` in tone, the names in class colour, a Late
# with its minutes (G5); Damage's melee and ranged split (Android on the title line, iOS a line under the names).
R5_COMING = [('Tanks', 2, 2, [('Grunvald', 'warrior'), ('Brakka', 'druid')], 0),
             ('Healers', 5, 8, [('Mirethil', 'priest'), ('Selka', 'druid'), ('Hollowmend', 'priest'), ('Aenwyn', 'paladin'),
                                ('Brightwen<br>(Late&nbsp;&middot;&nbsp;15&nbsp;min)', 'priest')], 0),
             ('Damage', 11, 30, [('Thandrel', 'warrior'), ('Kaelra', 'warrior'), ('Thalrin', 'mage')], 8)]
DAMAGE_SPLIT = '5 melee, 6 ranged'
RAID_INVITES = 'Invites 7:45 pm &middot; 40-man'
RAID_NOTES = ('Bring <b>fire resistance</b> gear and flasks. Whelps go to the left wall, the tanks call the deep breath in voice; '
              'consumables check at 7:50 pm.')
RAID_ANSWERS = 'Bench 1 &middot; Out 3 &middot; 5 not answered'
# The Covered line (round 3, row 4): the gaps nobody brings, then the short ones, by the buff's short name, at most three
COVER_MISSING = [('Fort', 'spell_holy_wordfortitude')]
COVER_SHORT = [('WF', 'spell_nature_windfury', '1 of 2')]
RAID_NOT_SYNCED = 'Not synced this week: Kaelra, Thalrin, Brakka, Mirethil'
R5_READY = [('ok', 'Drakefire Amulet', 'Attuned', None), ('ok', 'Elixir of the Mongoose', '3 in your bags', None),
            ('miss', 'Flask of the Titans', 'None in your bags', 'Auction house'),
            ('warn', 'Boots not enchanted', '', None)]
R5_PAGES = [('Roster', '38 members, 31 at 60'), ('Calendar', 'Onyxia Wed, Barrow Deeps Sat'), ('Loot history', 'Last raid: 4 pieces'),
            ('Guild sims', 'You are #7 of 14 DPS')]


def _r5_section(title: str, body: str, action: str = '', meta: str = '') -> str:
    m = f'<span class="meta">{meta}</span>' if meta else ''
    a = f'<span class="act">{action}</span>' if action else ''
    return f'<section data-region="{region_slug(title)}"><div class="sh"><h2>{title}</h2>{m}{a}</div><div class="pn">{body}</div></section>'


def _r5_row(title: str, sub: str = '', lead: str = '', trail: str = '') -> str:
    sub_html = f'<span class="s">{sub}</span>' if sub else ''
    inner = f'<div class="in">{lead}<div><span class="t">{title}</span>{sub_html}</div></div>' if lead else f'<span class="t">{title}</span>{sub_html}'
    return f'<div class="row">{inner}' + (f'<div class="tr">{trail}</div>' if trail else '') + '</div>'


def _r5_roles(roles: list, platform: str, split: bool = True) -> str:
    """Who is coming's rows. split: Damage's title carries its melee and ranged split (round 4: both apps on one line; the
    night's rows on iOS do not)."""
    out = ''
    for role, n, want, people, more in roles:
        short = want - n
        tally = f'{n} of {want}' + (f' &middot; {short} short' if short > 0 else '')
        dmg = split and role == 'Damage'
        title = f'{role} <span class="{"ok" if short <= 0 else "lo"}">{tally}</span>'
        if dmg:
            title += f'<span class="split"> &middot; {DAMAGE_SPLIT}</span>'
        names = ', '.join(f'<span style="color:{CLASS_COLOR[c]}">{nm}</span>' for nm, c in people) + (f' and {more} more' if more else '')
        out += _r5_row(title, names)
    return out


def _cover_icon(platform: str, name: str) -> str:
    """A buff's 20 pt spell icon: an image iOS counts (parityItem image); Android draws it without a parity item."""
    html = icon(name, 20)
    if platform == 'ios':
        return html
    src = html.split('src="', 1)[1].split('"', 1)[0]
    return f'<svg class="ci" width="20" height="20" viewBox="0 0 20 20" aria-hidden="true"><image href="{src}" width="20" height="20"/></svg>'


def _cover(platform: str) -> str:
    """The Covered line (round 3, row 4): `Missing`, then each gap in wipe red beside its spell icon (iOS sets ` · ` between
    the missing and the short ones), the short ones with their count."""
    gaps = [(n, i, '') for n, i in COVER_MISSING] + list(COVER_SHORT)

    def gap(k: int, name: str, ic: str, count: str) -> str:
        if platform == 'ios':  # the separator before each name after the first, the count muted
            sep = '<span class="n">&middot;&nbsp;</span>' if k else ''
            n = f' <span class="n">{count}</span>' if count else ''
            return f'<span class="cg">{sep}{_cover_icon(platform, ic) if ic else ""}<span>{name}{n}</span></span>'
        sep = '<span class="n"> &middot;</span>' if k < len(gaps) - 1 else ''  # Android: the separator after each name
        return f'<span class="cg">{_cover_icon(platform, ic) if ic else ""}<span>{name}{" " + count if count else ""}{sep}</span></span>'
    return '<div class="cov"><span>Missing</span>' + ''.join(gap(k, *g) for k, g in enumerate(gaps)) + '</div>'


def _r5_coming(platform: str, officer: bool = False) -> str:
    """Who is coming: the roles, then the other answers' counts, what the cover misses (G4) and, for officers, who signed In
    without syncing this week, as a row that opens the roster filtered to them (G1; iOS counts it as a row)."""
    unsynced = ''
    if officer:
        unsynced = (f'<a class="ns sr" href="#"><span>{RAID_NOT_SYNCED}</span><i>&rsaquo;</i></a>' if platform == 'ios'
                    else f'<a class="ns" href="#"><span>{RAID_NOT_SYNCED}</span><i>&rsaquo;</i></a>')
    foot = f'<div class="rfoot"><p>{RAID_ANSWERS}</p>{_cover(platform)}{unsynced}</div>'
    return _r5_section('Who is coming', _r5_roles(R5_COMING, platform) + foot, action='Roster &rsaquo;')


def _r5_head(platform: str, answer: str, late_min: int | None = None, status: str = '') -> str:
    """The raid head: the zone, the time and the pull, the invite time and size, who the answer is for (round 3: no answer
    chip, the answers are the state); the four answers (Android's 48 dp pill, iOS's Segmented), Late's minutes (Android asks
    How late? over them), the status line; the officers' notes, two lines, close the card on both."""
    seg = ''.join(f'<a class="btn{" on" if x == answer else ""}" href="#">{x}</a>' for x in ('In', 'Late', 'Bench', 'Out'))
    ask = '<p class="hl">How late?</p>' if platform == 'android' else ''
    late = (ask + '<div class="lm">' + ''.join(f'<a class="btn{" on" if m == late_min else ""}" href="#">{m} min</a>' for m in (15, 30, 60)) + '</div>'
            if answer == 'Late' else '')
    stl = f'<p class="stl">{status}</p>' if status else ''
    notes = f'<p class="notes">{RAID_NOTES}</p>'
    return (f'<div class="blk-r5" data-region="section-next-raid"><div class="pn"><div class="hd"><div class="rl">{icon(ONY, 40)}'
            f'<div class="tx"><b>{RAID}</b><span class="s">8:00 pm CST &middot; pull in 2 h 48 min</span><span class="s">{RAID_INVITES}</span>'
            f'<span class="s">{AS_ME_CAP}</span></div>'
            f'</div><div class="seg">{seg}</div>{late}{stl}{notes}</div></div></div>')


def _r5_ready(items: list, title: str = 'Before you pull') -> str:
    marks = {'ok': '&#10003;', 'miss': '&times;', 'warn': '!'}
    rows = ''.join(_r5_row(t, sub + (f' &middot; <a class="lk" href="#">{fix}</a>' if fix else ''), f'<span class="mk {k}">{marks[k]}</span>')
                   for k, t, sub, fix in items)
    return _r5_section(title, rows + '<p class="stamp" style="margin:0">From your last log-out, 9 h ago</p>')


def _r5_lone(title: str, lines: list[str], button: str, primary: bool = False) -> str:
    notes = ''.join(f'<p class="note{" strong" if i == 0 and len(lines) > 1 else ""}">{x}</p>' for i, x in enumerate(lines))
    return _r5_section(title, notes + f'<a class="btn wide{" primary" if primary else ""}" href="#">{button}</a>')


def _raid_r5(state: str, platform: str) -> str:
    """The Raid tab's body in round 5 (RAID_R5_CSS); the states and words are raid()'s. Round 2 (review backfill): an
    officer's Needs you sits directly under the head (G3)."""
    pages = _r5_section('Olympus XXVII', ''.join(_r5_row(t, sub) for t, sub in R5_PAGES))
    if state in ('member', 'officer', 'signup'):
        # signup: Late with 30 minutes just sent, so the status line says Saved.
        head = _r5_head(platform, 'Late', 30, 'Saved.') if state == 'signup' else _r5_head(platform, 'In')
        needs = (_r5_section('Needs you', _r5_row('5 have not answered, 2 of them healers', 'Healers short by 3', trail='<a class="btn" href="#">Ping on Discord</a>')
                             + _r5_row('1 applicant', 'Mirelle, Holy Priest 60, 2 min ago'), meta='6 open') if state == 'officer' else '')
        return head + needs + _r5_coming(platform, officer=state == 'officer') + _r5_ready(R5_READY) + pages
    if state == 'noraid':
        return ('<div class="blk-r5" data-region="section-next-raid"><div class="pn"><p class="note" style="margin:0;padding:8px 0;font-size:14px;line-height:20px">'
                'No raid posted yet.</p></div></div>' + pages)
    if state == 'live':
        night = ('<div class="lvh"><i></i><span>Olympus XXVII &middot; Onyxia&rsquo;s Lair in progress</span></div>'
                 '<p class="note">Pulled 1 h 12 min ago &middot; 18 in, 3 late &middot; you are In as<br>Thandrel, Damage</p>'
                 f'<div class="rule"></div>{_r5_roles(R5_COMING, platform, split=platform == "android")}')
        return _r5_section('Raid night', night) + _r5_ready(R5_READY, 'Ready check') + pages
    return _r5_lone('Raid with a guild', [], 'Sign in with Battle.net', primary=True)


def _raid_page_r5(state: str, platform: str) -> str:
    return build_page('Raid', f'<div class="ap r5raid">{_raid_r5(state, platform)}</div><style>{FLOWS_CSS}{RAID_R5_CSS}</style>',
                      note='Native app mock (design/specs/2026-11-08-app-flows.md), round 5.')


# ------------------------------------------------------------------------------------------------ Raid (the tab)
def raid(state: str):
    """The Guild tab's raid states (owner, 2026-11-08; the tab is Guild since 2026-11-10, ws_app_guild.py draws its own
    sections): the next raid first with your answer, then for an officer what needs you, then the guild's raid pages. Live:
    raid night leads. Signed out: sign in. No guild and under 60 are ws_app_guild.py's."""
    def b(platform: str) -> str:
        badges = {'member': ANSWERED, 'signup': ANSWERED, 'officer': OFFICER, 'noraid': ANSWERED, 'live': ANSWERED,
                  'out': {}}[state]
        who = {'out': None}.get(state, 'warrior')
        return app_page(platform, _raid_page_r5(state, platform), 'guild', title='Guild', badges=badges, who=who)
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
    ('Finger 1', 'Don Julio&rsquo;s Band', 'inv_jewelry_ring_35', 4, 'Stormpike Guard', '+6.8'),
    ('Finger 2', 'Cutthroat&rsquo;s Signet', 'inv_jewelry_ring_01', 3, 'Pix Xizzix, vendor', '+0.8'),
    ('Trinket 1', 'Hand of Justice', 'inv_jewelry_talisman_01', 3, 'Blackrock Depths: Emperor Dagran Thaurissan', None),
    ('Trinket 2', 'Molten Heart of the Mountain', 'inv_misc_gem_ruby_01', 3, 'Enchanting', '+3.2'),
    ('Main hand', 'Ironwood Blade', 'inv_sword_07', 4, 'Blacksmithing', '+5.9'),
    ('Off hand', 'Greenhammer', 'inv_mace_1h_raidhyjalc60_d_01', 4, 'Blacksmithing', '+9.7'),
    ('Ranged', 'Blackcrow', 'inv_weapon_crossbow_04', 3, 'Blackrock Spire: Shadow Hunter Vosh&rsquo;gajin, 8%', '+1.9'),
]


# Round 5 on the Gear tab and item sheet (lanes M3a, M3i; spec 2026-11-09-directions.md Paperdoll "Mobile Gear tab", mocks
# directions/png/m-android-gear-frame.png and m-android-item.png): the paperdoll's phone form (the strip with the portrait in
# its gold-ringed medallion, the spec and the set's DPS; the line naming the list with Change; the slots in two columns, each
# its gain or Equipped), then Where to go tonight on a plain panel. Greenhammer unlocks with the first raids, so the off hand is
# the Persuader he wears (the directions mock's GATED). The sheet is the client's tooltip with the gain closing it, then one
# list: where it drops, what it replaces, the runner-up and the guild. iOS since lane M3i. Line heights are the
# Pixel 9's rounding (see RAID_R5_CSS); Open Sans sets wider there, so the lines it breaks carry the break here (<br>).
GEAR_R5_CSS = """
.ap.r5gear{padding:16px 14px 22px}
.ap.r5gear>*{margin-top:0}.ap.r5gear>*+*{margin-top:20px}
.r5gear .pd{background:#170e09;border-radius:10px;box-shadow:0 0 0 1px #000;overflow:hidden}
.r5gear .strip{display:flex;align-items:center;gap:10px;padding:8px 12px;background:#0b2f3b;border-bottom:1px solid #000}
.r5gear .strip .med{width:40px;height:40px;flex:none;border-radius:50%;box-sizing:border-box;border:1.5px solid #b8a24a;padding:1.5px}
.r5gear .strip .med img{display:block;width:100%;height:100%;border-radius:50%}
.r5gear .strip h1{flex:1;margin:0;font:400 19px/23px var(--font-title);color:#f3eee2}
.r5gear .strip b{font:600 17px/23px var(--font-body);color:#fff}.r5gear .strip b small{font-size:11px;color:var(--text)}
.r5gear .lead{margin:0;padding:12px 12px 0;font:400 14px/20px var(--font-body);color:var(--text)}.r5gear .lead span{display:block}
.r5gear .lead a{font-weight:600;color:var(--gold)}
/* round 2 (review backfill): the measured header under the band line (iOS 13/18 muted, outside the controls region;
   Android the note, inside it) */
.r5gear .lead .sum{margin-top:4px;font:400 13px/18px var(--font-body);color:var(--muted)}
[data-platform=android] .r5gear .lead .sum{font-size:14px;line-height:20px}
[data-platform=android] .r5gear .strip h1{font-size:20px;line-height:26px}
[data-platform=android] .r5gear .strip b{font-size:16px;font-weight:700}[data-platform=android] .r5gear .strip b small{font-size:12px}
[data-platform=android] .r5gear .lead [data-region=controls]>span:first-child{display:block;white-space:nowrap;letter-spacing:-.15px}
.r5gear .ls{padding:12px}.r5gear .ls>div{display:grid;grid-template-columns:1fr 1fr;gap:6px 10px}
.r5gear .cell{display:flex;align-items:center;gap:8px;min-height:48px;min-width:0;padding:0;border:0}
.r5gear .sk{width:44px;height:44px;flex:none;box-sizing:border-box;padding:3px;border:1px solid var(--line);border-radius:3px;background:#100a06}
.r5gear .sk img{display:block;width:38px;height:38px}
.r5gear .cell .n{display:block;font:600 13px/16px var(--font-body)}
.r5gear .cell .s{display:block;font:400 12px/16px var(--font-body);color:var(--muted)}.r5gear .cell .g{font-weight:700;color:var(--gain,#32ae88)}
.r5gear .sh{display:flex;align-items:center;height:26px;margin-bottom:8px}
.r5gear .sh h2{margin:0;font:400 21px/26px var(--font-title);color:#f3eee2}
.r5gear .pn{background:#170e09;border-radius:10px;padding:8px 16px;box-shadow:0 0 0 1px #000}
.r5gear .row{display:flex;align-items:center;gap:12px;padding:10px 0;border-bottom:1px solid var(--line)}.r5gear .row:last-child{border-bottom:0}
.r5gear .row img{width:40px;height:40px;flex:none}
.r5gear b.t{display:block;font:700 15px/20.57px var(--font-body);color:#fff}
.r5gear .row .s{display:block;font:400 13px/18.29px var(--font-body);color:var(--muted)}
.sheetp.r5item{padding:0 16px 24px;border-top:0;border-radius:16px 16px 0 0}
.r5item .sheet-grab{margin:9px auto}
.r5item .tip{background:#060914;border:1px solid #3a4256;border-radius:6px;padding:12px 14px;display:flex;flex-direction:column;gap:4px;
  font:400 16px/22px var(--font-body);color:#fff}
.r5item .tip .nm{font-weight:700;font-size:17px;color:#b866f5}.r5item .tip .sp{display:flex;justify-content:space-between}
.r5item .tip .gr{color:#1eff00}.r5item .tip .gn{font-weight:600;color:var(--gold)}.r5item .tip .gap{height:4px;margin:-4px 0 0}
.r5item .lines{margin-top:12px}
.r5item .row{display:flex;align-items:center;gap:12px;min-height:0;padding:10px 0;border-bottom:1px solid var(--line)}.r5item .row:last-child{border-bottom:0}
.r5item .row img{width:40px;height:40px;flex:none}
.r5item b.t{display:block;font:700 15px/20.57px var(--font-body);color:#fff}
.r5item .row .s{display:block;font:400 13px/18.29px var(--font-body);color:var(--muted)}
"""

R5_DOLL = [r if r[1] != 'Greenhammer' else ('Off hand', 'Persuader', 'inv_hammer_08', 4, 'Blacksmithing', None) for r in DOLL]
# place, what it gives, icon, quality, and (round 2) its upgrades and their gains together
R5_WHERE = [('Blackrock Spire', 'Blackcrow &middot; Shadow Hunter Vosh&rsquo;gajin, 8%', 'inv_weapon_crossbow_04', 3, '1 upgrade &middot; +1.9 DPS'),
            ('Blacksmithing', 'Lionheart Helm, Ironwood Blade &middot; ask a blacksmith', 'inv_helmet_36', 4, '2 upgrades &middot; +20.4 DPS'),
            ('Stormpike Guard', 'Don Julio&rsquo;s Band, needs revered', 'inv_jewelry_ring_35', 4, '1 upgrade &middot; +6.8 DPS')]
# The header line (round 2): the measured total, the gains not modelled, when the gear synced; the apps order it apart.
# The header line (round 3): what the apps compute from the fixture (12 upgrades, their gains together, the synced age).
GEAR_SUMMARY = '+47.5 DPS from 12 upgrades &middot; synced 2 h ago'
GEAR_FACTS = '60 raid &middot; Alliance &middot; 5 of 17 slots match the list &middot; <a href="#">Change</a>'


def _r5_gear_row(title: str, sub: str, lead: str, color: str = '#fff') -> str:
    s = f'<span class="s">{sub}</span>' if sub else ''
    return f'<div class="row">{lead}<div><b class="t" style="color:{color}">{title}</b>{s}</div></div>'


def _gear_page_r5(platform: str) -> str:
    """The Gear tab's body in round 5 (GEAR_R5_CSS)."""
    cells = ''.join(f'<div class="cell"><span class="sk">{icon(ic, 38, f"q{q}")}</span><div><span class="n" style="color:{QCOL[q]}">{name}</span>'
                    f'<span class="s">{slot} &middot; ' + ('Equipped' if g is None else f'<span class="g">{g} DPS</span>') + '</span></div></div>'
                    for slot, name, ic, q, _, g in R5_DOLL)
    doll = (f'<div class="pd"><div class="strip"><span class="med"><img src="{AVATAR.format(name="thandrel")}" alt=""></span>'
            f'<h1>Fury Warrior</h1><b>1,004 <small>DPS</small></b></div>'
            + _gear_lead(platform)
            + f'<div class="ls"><div data-region="slots">{cells}</div></div></div>')
    where = ''.join(_r5_gear_row(t, f'{s}<br>{fig}', icon(ic, 40, f'q{q}')) for t, s, ic, q, fig in R5_WHERE)
    body = doll + (f'<section data-region="{region_slug("Where to go")}"><div class="sh"><h2>Where to go</h2></div>'
                   f'<div class="pn">{where}</div></section>')
    return build_page('Gear', f'<div class="ap r5gear">{body}</div><style>{FLOWS_CSS}{GEAR_R5_CSS}</style>',
                      note='Native app mock (design/specs/2026-11-08-app-flows.md), round 5.')


def _gear_lead(platform: str) -> str:
    """The line naming the list with Change, then the measured header (iOS holds the controls region to the line alone)."""
    summary = f'<span class="sum">{GEAR_SUMMARY}</span>'
    if platform == 'ios':
        return f'<p class="lead"><span data-region="controls">{GEAR_FACTS}</span>{summary}</p>'
    return f'<p class="lead"><span data-region="controls"><span>{GEAR_FACTS}</span>{summary}</span></p>'


def _item_sheet_r5(platform: str) -> str:
    """Deathbringer in round 5 (GEAR_R5_CSS): the tooltip in the client's order, the gain closing it, then the list. iOS
    adds the guild's next raid to the drop's zone under the source (needs.next_raid; it opens that raid; Android to follow)."""
    tip = ('<div class="tip" data-region="inspector"><span class="nm">Deathbringer</span><span>Binds when picked up</span>'
           '<span class="sp"><span>One-Hand</span><span>Axe</span></span><span class="sp"><span>114 - 213 Damage</span><span>Speed 2.90</span></span>'
           '<span>(56.4 damage per second)</span><span class="gap"></span><span>Requires Level 60</span>'
           '<span class="gr">Chance on hit: Sends a shadowy bolt at the<br>enemy causing 125 Shadow damage.</span><span class="gn">+12.6 DPS</span></div>')
    next_raid = _r5_gear_row('Next raid: Onyxia&rsquo;s Lair, Wed 8:00 PM', '', '') if platform == 'ios' else ''
    lines = (_r5_gear_row('Onyxia, Onyxia&rsquo;s Lair, 8%', '', '') + next_raid
             + _r5_gear_row('Replaces Ironwood Blade', 'Equipped', icon('inv_sword_07', 40, 'q4'))
             + _r5_gear_row('Vis&rsquo;kag the Bloodletter', 'Runner-up &middot; &minus;1.8 DPS', icon('inv_sword_18', 40, 'q4'), QCOL[4])
             + _r5_gear_row('Grunvald', 'Protection Warrior &middot; needs it, the pick on their list', crest_img('warrior', 40), WARRIOR)
             + _r5_gear_row('Torvak', 'Enhancement Shaman &middot; has it', crest_img('shaman', 40), CLASS_COLOR['shaman']))
    sheet = ('<div class="dim" style="height:120px"></div><div class="sheetp r5item" role="dialog" aria-label="Deathbringer"><span class="sheet-grab"></span>'
             f'{tip}<div class="lines" data-region="section-where-to-get-it">{lines}</div></div>')
    return build_page('Deathbringer', sheet + f'<style>{FLOWS_CSS}{GEAR_R5_CSS}</style>', note='Item sheet mock, round 5.')


def gear(state: str):
    """Gear (owner, 2026-11-08): your gear against your band's list. The summary line names the list (your spec, band and
    faction, preselected; Change opens the picker); then where to go tonight; then every slot (Home holds the top three)."""
    def b(platform: str) -> str:
        if state == 'in':
            return app_page(platform, _gear_page_r5(platform), 'gear', title='Gear', badges=MEMBER)
        if state == 'out':
            body = (PageHeader('Best in slot for your class', meta=['Pick a class, then a spec and a level range.'])
                    + Section('Classes', Rows([Row(c.capitalize(), '', crest_img(c, 40), title_color=CLASS_COLOR[c])
                                               for c in ('warrior', 'paladin', 'hunter', 'rogue', 'priest', 'shaman', 'mage', 'warlock', 'druid')], open_=True)))
            return app_page(platform, _page(body, 'Gear'), 'gear', who=None)
    return b


# ------------------------------------------------------------------------------------------------ Activity (the tab)
ACTIVITY = [
    ('Today', [('ability_warrior_innerrage', 'Kaelra passed you in the guild sims', 'You are #7 of 14', 'Guild ladder', '8 h ago', True),
               ('inv_axe_09', f'You got <span style="color:{QCOL[4]}">Deathbringer</span>', 'Onyxia &middot; main hand, +12.6 DPS', 'Gear and sims', '9 h ago', True),
               (ONY, 'Onyxia&rsquo;s Lair report ready', '87th percentile, your best on Onyxia so far &middot; the coach has 2 notes', 'Logs', '9 h ago', True)]),
    ('Yesterday', [(ONY, 'You&rsquo;re In for Onyxia', AS_ME_CAP + ' &middot; 8:00 pm CST', 'Raids', '6:02 pm', False),
                   ('inv_misc_pocketwatch_01', 'Sim done: 949 DPS &plusmn;6', '+21 vs your last', 'Gear and sims', '2:40 pm', False),
                   (ONY, 'Healers short for Onyxia tonight', '3 of 8 missing', 'Raids', '1:15 pm', False)]),
    ('This week', [('inv_misc_book_09', 'Fury Warrior moved to S', 'Raid tiers, 60', 'Tiers', 'Tue', False),
                   ('inv_misc_book_09', 'Fury Warrior BiS changed', f'A new pick for Hands: <span style="color:{QCOL[4]}">Theramore Gauntlets</span>', 'Gear and sims', 'Tue', False)]),
]
# The row's figure (round 3, row 11): once, at its end, formatted like the rest.
ACTIVITY_FIGURES = {'Kaelra passed you in the guild sims': '1,196 DPS'}
GROUPS = ['All', 'Raids', 'Gear and sims', 'Logs', 'Guild ladder', 'Tiers']
# The rows about a class wear its crest (one crest language, lane M4a): a guildmate's spec icon, a spec's tier move.
ACTIVITY_CRESTS = {'ability_warrior_innerrage': 'warrior', 'Fury Warrior moved to S': 'warrior'}

# Round 5 on the Activity tab (lanes M4a, M4i; spec 2026-11-09-directions.md Mobile, mock directions/png/m-android-activity.png):
# the blue band carries the bar and the group chips (36 faces in 48 dp targets, the first teal), then a panel per day on the
# ground with its display title and Mark all read on the first; each row its lead (the crest for a row about a class), title,
# line and time. iOS since lane M4i. Line heights are the Pixel 9's rounding (see RAID_R5_CSS).
ACTIVITY_R5_CSS = """
.ap.r5act{padding:0 14px 22px}
.r5act .bandc{margin:0 -14px;padding-left:16px;background:#0c3a52;border-bottom:1px solid #000}
.r5act .chips{display:flex;gap:8px;padding:2px 16px 8px 0;overflow:hidden;white-space:nowrap;-webkit-mask-image:linear-gradient(90deg,#000 calc(100% - 32px),transparent);mask-image:linear-gradient(90deg,#000 calc(100% - 32px),transparent)}
.r5act .chips .btn{flex:none;display:flex;align-items:center;box-sizing:border-box;height:36px;min-height:36px;margin:6px 0;padding:0 16px;
  border:1px solid #5a5d70;border-radius:8px;background:none;color:#fff;font:600 14px/20px var(--font-body)}
.r5act .chips .btn.on{background:#0f6064;border-color:#408076}
.ap.r5act>section,.ap.r5act>.blk-r5{margin-top:20px}.ap.r5act>.bandc+*{margin-top:16px}.ap.r5act>.blk-r5:first-child{margin-top:16px}
.r5act .sh{display:flex;align-items:center;height:26px;margin-bottom:8px}
.r5act .sh h2{margin:0;font:400 21px/26px var(--font-title);color:#f3eee2}
.r5act .sh .act{margin-left:auto;font:600 14px/24.38px var(--font-body);color:var(--gold)}
.r5act .pn{background:#170e09;border-radius:10px;padding:8px 16px;box-shadow:0 0 0 1px #000}
.r5act .row{position:relative;display:flex;align-items:center;gap:12px;padding:10px 0;border-bottom:1px solid var(--line)}.r5act .row:last-child{border-bottom:0}
.r5act .row img{width:40px;height:40px;flex:none}
.r5act .row.dot::before{content:"";position:absolute;left:-10px;top:50%;width:6px;height:6px;margin-top:-6px;border-radius:50%;background:var(--gold)}
.r5act b.t{display:block;font:700 15px/20.57px var(--font-body);color:#fff}
.r5act .s{display:block;font:400 13px/18.29px var(--font-body);color:var(--muted)}
.r5act .row .fig{margin-left:auto;flex:none;font:700 15px/20.57px var(--font-body);font-variant-numeric:tabular-nums;color:#fff}
.r5act .more{display:flex;align-items:center;min-height:48px;margin-top:8px;padding:0 16px;font:600 15px/20px var(--font-body);color:var(--gold)}
[data-platform=ios] .r5act .more{min-height:44px;padding:0}
[data-platform=android] .r5act .row .fig{line-height:22px}[data-platform=android] .r5act .sh .act{line-height:18px}
.r5act .note{margin:0;padding:8px 0;font:400 13px/20.57px var(--font-body);color:var(--text)}
.r5act .wide{display:flex;align-items:center;justify-content:center;height:48px;margin:12px 0;border-radius:4px;background:#0f6064;border:1px solid #408076;
  color:#fff;font:700 15px/20px var(--font-body)}
"""


def _act_lead(ic: str, title: str) -> str:
    crest = ACTIVITY_CRESTS.get(ic) or ACTIVITY_CRESTS.get(title)
    return crest_img(crest, 40) if crest else icon(ic, 40)


def _activity_r5(state: str) -> str:
    """The Activity tab's body in round 5 (ACTIVITY_R5_CSS); the rows and words are activity()'s."""
    if state == 'out':
        return ('<div class="blk-r5" data-region="section-activity"><div class="pn"><p class="note">What changed for your characters: loot, sims, '
                'reports, raids.</p><a class="btn wide" href="#">Sign in with Battle.net</a></div></div>')
    if state == 'empty':  # nothing to filter: no chips under the bar
        return '<div class="blk-r5" data-region="section-empty"><div class="pn"><p class="note">Nothing yet. Loot, sims, reports and raids land here.</p></div></div>'
    chips = ''.join(f'<a class="btn{" on" if g == "All" else ""}" href="#">{g}</a>' for g in GROUPS)
    band = f'<div class="bandc"><div class="chips" data-region="controls">{chips}</div></div>'
    days = ''
    for when, rows in ACTIVITY:
        act = '<span class="act">Mark all read</span>' if when == 'Today' else ''
        lines = ''.join(f'<div class="row{" dot" if u else ""}">{_act_lead(i, t)}<div><b class="t">{t}</b><span class="s">{s}</span>'
                        f'<span class="s">{at}</span></div>' + (f'<span class="fig">{ACTIVITY_FIGURES[t]}</span>' if t in ACTIVITY_FIGURES else '')
                        + '</div>' for i, t, s, _, at, u in rows)
        days += f'<section data-region="{region_slug(when)}"><div class="sh"><h2>{when}</h2>{act}</div><div class="pn">{lines}</div></section>'
    return band + days + '<a class="more" href="#">Show older</a>'


def _activity_page_r5(state: str) -> str:
    return build_page('Activity', f'<div class="ap r5act">{_activity_r5(state)}</div><style>{FLOWS_CSS}{ACTIVITY_R5_CSS}</style>',
                      note='Native app mock (design/specs/2026-11-08-app-flows.md), round 5.')


def activity(state: str):
    """Activity (owner, 2026-11-08): every notification event, newest first, read state, the six groups as filters (All
    preselected); every row opens its answer. The push inbox: a player with pushes off sees the same list."""
    def b(platform: str) -> str:
        badges = {'in': {'guild': 0, 'activity': 3}}.get(state, {})
        return app_page(platform, _activity_page_r5(state), 'activity', title='Activity', badges=badges, who=None if state == 'out' else 'warrior')
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
    return app_page(platform, _page(body, RAID), 'guild', back='Guild', title=RAID, share=True, badges=OFFICER)


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
        # Android sets a 28 sp figure on a 38.5 dp line (the Pixel 9's font padding), not the web's 32 (lane M2a, was M0's drift)
        android = '<style>[data-platform=android] .ap .figs .fig .v{line-height:38.5px}</style>'
        return app_page(platform, _page(body + android, 'Mirelle'), 'guild', back='Guild', title='Mirelle', badges=OFFICER)
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
    if platform == 'ios':
        # officer_dms (D16, off by default; built on iOS, Android to follow): channel mentions are the default; on, the
        # officers' Not answered row gets a DM link.
        body += Section('Discord', Rows(['<div class="row setrow"><div class="b"><div class="ttl">Let officers DM me on Discord</div>'
                                         '<div class="sub">When you have not answered a raid; off, they mention you in the channel</div></div>'
                                         + Switch('', False) + '</div>'], open_=True))
    return app_page(platform, _page(body, 'Notifications'), 'home', back='Home', title='Notifications')


# ------------------------------------------------------------------------------------------------ F4, F5: the item sheet
def item_sheet(platform: str) -> str:
    """Deathbringer from Onyxia (data/builds 1.60.1.70790: 8%, item level 75): the client's tooltip lines, where it drops,
    the runner-up from the same boss, and who in the guild has it or needs it (review S1)."""
    return app_page(platform, _item_sheet_r5(platform), 'gear', navbar=False)

SCREENS = {
    'menu': menu('in'),
    'menu-one': menu('one'),
    'menu-low': menu('low'),
    'menu-out': menu('out'),
    **{k: v() for k, v in HOME.items()},
    'raid': raid('member'),
    'raid-signup': raid('signup'),
    'raid-officer': raid('officer'),
    'raid-live': raid('live'),
    'raid-no-raid': raid('noraid'),
    'raid-out': raid('out'),
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
