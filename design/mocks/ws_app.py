"""Native app mocks for the mobile redesign (design/specs/2026-11-07-mobile-redesign.md), iOS 393 x 852 and Android 412 x 915.

    python3 design/mocks/ws_app.py [filter]        -> design/mocks/app/<platform>-<screen>.html
    node design/mocks/ws_app_shoot.mjs [filter]    -> design/reviews/2026-11-07-mobile/<platform>-<screen>.png

Two kinds of screen (spec section 4):
- WEB: a screen that matches a web page is that page's committed phone mock (design/mocks/web/<mock>.html, the website's
  390 px layout) inside the app frame (ws_app_frame.py). The phone deltas the persona rulings add (home order R2, the
  level 34 home R4, the feed filters R5) are CSS over the same markup, so the web mock stays the one source of structure
  and data. A web mock that changes changes the app mock on the next run.
- APP: app-only screens (the profile menu and the flows screens in ws_app_flows.py, devices, deletion, offline data, update required, raid night, widgets, Live Activity,
  search, the signed-out gate) composed from ws_prims primitives alone, in the same system.
Data: the home and guild fixtures (web/tests/e2e/fixtures/home/*.json): Thandrel (Fury Warrior 60, Olympus XXVII) and
Felsnap (Beast Mastery Hunter 34, no guild). The clock is ws_home.NOW (Fri 6 Nov 2026, 18:00 UTC).
"""
from __future__ import annotations

import json
import sys
from html import escape

from mocklib import ROOT, CLASS_COLOR
from ws_app_frame import DEVICES, app_page, from_web, glyph, write_app
import ws_app_flows as flows
import ws_app_guild as guild_tab
from ws_home import NOW, RAIDS, clock_to, day_text, raids_open
from ws_prims import (Button, Checklist, Countdown, Figure, Notice, PageHeader, Rows, Row, Section, Switch, Tag, Meter, build_page,
                      crest_img, icon)

FIX = ROOT / 'web/tests/e2e/fixtures/home'


def fx(name: str) -> dict:
    return json.loads((FIX / f'{name}.json').read_text())


ME60, ME34, HOME60 = fx('me-60'), fx('me-34'), fx('home-60')
GUILD60 = fx('guild-home-60')
NEXT_RAID = GUILD60['next_raid']
WARRIOR = CLASS_COLOR['warrior']

# ------------------------------------------------------------------------------------------------ phone deltas (CSS)
# The signed-out home (R3) is the web home mock; the signed-in Home states are app-composed in ws_app_flows.py.
# R5: three filter segments full width under the feed title; Logs moves behind the filter button at the end.
FEED_FILTERS = '''
[data-region="section-forever-right-now"] .sec-h{flex-wrap:wrap;row-gap:10px}
[data-region="section-forever-right-now"] .sec-h .act{flex-basis:100%;margin-left:0;display:flex;gap:8px}
[data-region="section-forever-right-now"] .seg{flex:1;display:flex}
[data-region="section-forever-right-now"] .seg a{flex:1;justify-content:center;height:44px;padding:0 8px}
[data-region="section-forever-right-now"] .seg a:nth-child(4){display:none}
[data-region="section-forever-right-now"] .sec-h .act::after{content:"";flex:none;width:44px;height:44px;border:1px solid var(--line);border-radius:4px;
 background:center/20px no-repeat url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24' fill='none' stroke='%23b9b3a4' stroke-width='1.8' stroke-linecap='round' stroke-linejoin='round'%3E%3Cpath d='M4 5h16l-6 8v6l-4-2v-4z'/%3E%3C/svg%3E")}
'''
# R3: signed out, How to start becomes one footer link (the rail section is hidden; the link is the app's footer row).
HOME_OUT = '''
[data-region="section-how-to-start"]{display:none}
'''


def _out_footer() -> str:
    return '<div class="pc" style="order:9;padding-top:24px"><a class="link" href="#">How to start &rsaquo;</a></div>'


def _ptr() -> str:
    return '<div class="app-ptr" aria-hidden="true"><span></span></div>'


def web_screen(mock: str, tab: str | None, back: str = '', title: str = '', share: bool = False, css: str = '', inject: str = '',
               ptr: bool = False, who: str | None = 'warrior', badges: dict | None = None):
    def build(platform: str) -> str:
        html = from_web(mock, platform, tab, back, title, share, extra_css=css, who=who, badges=badges)
        if inject or ptr:
            html = html.replace('<main>', '<main>' + (_ptr() if ptr else ''), 1).replace('</main>', inject + '</main>', 1)
        return html
    build.web = True
    return build


# ------------------------------------------------------------------------------------------------ app-only screens
APP_CSS2 = '''
.ap{padding:8px 16px 24px}
.ap .t-title{font-size:32px;line-height:38px}
.ap section{margin-top:28px}
.sheet-grab{display:block;width:36px;height:5px;border-radius:3px;background:var(--line);margin:8px auto 4px}
.scrim{background:#000000a6;height:160px}
.sheetp{background:var(--raised);border-top-left-radius:12px;border-top-right-radius:12px;border-top:1px solid var(--line);padding:0 16px 24px}
.sheetp h2{font:700 18px/24px var(--font-body);color:var(--strong);padding:12px 0 4px}
.rowbtn{display:flex;align-items:center;justify-content:center;min-height:44px;font-weight:600}
.stack{display:flex;flex-direction:column;gap:12px;margin-top:20px}.stack .btn{width:100%;height:48px}
.wgrid{display:grid;grid-template-columns:repeat(2,minmax(0,1fr));gap:16px}
.wd{background:var(--raised);border:1px solid var(--line);border-radius:22px;padding:14px;aspect-ratio:1;display:flex;flex-direction:column;gap:6px;min-width:0}
.wd.m{grid-column:1/-1;aspect-ratio:auto;height:170px;flex-direction:row;gap:16px}
.wd.m>div{flex:1;display:flex;flex-direction:column;gap:6px;min-width:0}.wd.m>div+div{border-left:1px solid var(--line-soft);padding-left:16px}
.wd .k{font:700 12px/16px var(--font-body);color:var(--muted)}
.wd .big{font:700 30px/34px var(--font-body);font-variant-numeric:tabular-nums;color:var(--strong)}
.wd .nm{font:700 15px/20px var(--font-body);color:var(--strong)}
.wd .s{font-size:13px;line-height:18px;color:var(--muted)}
.wd .g{color:var(--gold)}
.lock{background:linear-gradient(180deg,#17202f,#07090d);border-radius:24px;padding:20px 16px;display:flex;flex-direction:column;gap:14px}
.lock .tm{font:600 64px/1 var(--font-body);text-align:center;color:var(--strong)}
.lock .row3{display:flex;gap:10px;justify-content:center}
.lw{display:flex;align-items:center;gap:8px;padding:8px 12px;border-radius:14px;background:#ffffff1f;color:var(--strong);font:600 14px/18px var(--font-body)}
.la{background:#0d111ae6;border-radius:20px;padding:14px 16px;display:flex;flex-direction:column;gap:10px}
.la.bare{padding:12px 0 0;background:none}
.la .top{display:flex;align-items:center;gap:12px}.la .top .nm{font:700 16px/20px var(--font-body);color:var(--strong)}
.la .clk{margin-left:auto;font:700 24px/28px var(--font-body);font-variant-numeric:tabular-nums;color:var(--gold)}
.la .btns{display:grid;grid-template-columns:repeat(4,1fr);gap:8px}.la .btns .btn{height:44px;padding:0}
.island{display:flex;align-items:center;justify-content:space-between;gap:10px;width:250px;height:37px;margin:0 auto;padding:0 14px;border-radius:20px;background:#000;color:var(--strong);font:700 14px/1 var(--font-body)}
.island .g{color:var(--gold);font-variant-numeric:tabular-nums}
.island.x{width:100%;height:auto;border-radius:36px;padding:16px 20px;flex-direction:column;align-items:stretch;gap:10px}
.notif{background:#1b2230;border-radius:20px;padding:14px 16px;display:flex;flex-direction:column;gap:8px}
.notif .top{display:flex;gap:8px;align-items:center;font-size:13px;color:var(--muted)}
.notif .nm{font:700 15px/20px var(--font-body);color:var(--strong)}
.lbl{font:700 12px/16px var(--font-body);color:var(--muted);margin:22px 0 8px}
.searchf{display:flex;align-items:center;gap:8px;height:44px;padding:0 12px;border:1px solid var(--line);border-radius:4px;background:#0a0d14;color:var(--muted)}
.seal{display:block;width:84px;height:84px;margin:40px auto 8px}
.center{text-align:center}.center .t-small{max-width:300px;margin:8px auto 0}
'''


def page(body: str, title: str) -> str:
    return build_page(title, f'<div class="ap">{body}</div><style>{APP_CSS2}</style>', note='Native app mock (design/specs/2026-11-07-mobile-redesign.md).')


def title(text: str, color: str = '', crest: str = '', meta: list[str] | None = None, answer: str = '') -> str:
    return PageHeader(text, color, crest, meta, answer)


def thandrel_crest(px: int = 40) -> str:
    return crest_img('warrior', px)


def devices(platform: str) -> str:
    phone = 'iPhone 16 Pro' if platform == 'ios' else 'Pixel 9'
    rows = [Row(f'{phone} &middot; This device', 'Signed in 30 Oct &middot; Last used just now', trail=Tag('This device', 'muted'), link=False),
            Row('Companion on Windows', 'Signed in 24 Oct &middot; Last used 12 min ago', trail=Button('Revoke', small=True, tone='danger'), link=False),
            Row('iPad Air', 'Signed in 2 Nov &middot; Last used 3 d ago', trail=Button('Revoke', small=True, tone='danger'), link=False)]
    body = (title('Devices') + Section('Signed in', Rows(rows, open_=True))
            + '<div class="stack">' + Button('Sign out everywhere else', tone='danger') + '</div>'
            + '<p class="t-small" style="margin-top:16px">Revoke signs a device out at once. To sign this one out, use Sign out in Settings.</p>')
    return app_page(platform, page(body, 'Devices'), 'home', back='Home', title='Devices')


DELETE_STATES = {
    'free': ('Delete your account', ['Your characters, sims, builds and logs you uploaded', 'Your guild sign-ups and your place on guild rosters',
                                     'Your devices: every phone and the companion are signed out'], ''),
    'premium': ('Delete your account', ['Your characters, sims, builds and logs you uploaded', 'Your guild sign-ups and your place on guild rosters',
                                        'Your devices: every phone and the companion are signed out'],
                'Your Premium ends when the account is deleted. Cancel the subscription in the App Store or Google Play first, or you are billed again.'),
}


def delete(state: str):
    def b(platform: str) -> str:
        if state == 'done':
            body = (f'<div class="center"><img class="seal" src="../../../web/public/seal-400.png" alt="">'
                    + title('Account deleted') + '<p class="t-small">Everything is removed within 30 days. You can sign in with Battle.net again any time to start fresh.</p>'
                    + '<div class="stack">' + Button('Done', primary=True) + '</div></div>')
            return app_page(platform, page(body, 'Account deleted'), None, navbar=False)
        if state == 'failed':
            n = Notice('That did not go through. Nothing was deleted. Try again.', tone='neutral')
            body = title('Delete your account') + n + '<div class="stack">' + Button('Try again', primary=True) + Button('Cancel') + '</div>'
            return app_page(platform, page(body, 'Delete your account'), 'home', back='Home')
        head, lose, warn = DELETE_STATES[state]
        rows = Rows([Row(x, link=False, slim=True) for x in lose], open_=True)
        body = (title(head) + '<p class="t-lead" style="margin-top:8px">This removes your characters, sims and guild sign-ups. It cannot be undone.</p>'
                + Section('What goes', rows)
                + (Notice(warn) if warn else '')
                + Section('When', '<p class="t-small" style="padding-top:12px">At once from the app and the site. Backups clear within 30 days.</p>')
                + '<div class="stack">' + Button('Delete my account', tone='danger') + Button('Cancel') + '</div>')
        return app_page(platform, page(body, head), 'home', back='Home', title=head)
    return b


def offline(state: str):
    def b(platform: str) -> str:
        rows = [('Guides', 'Fury Warrior, Beast Mastery Hunter &middot; 4.2 MB &middot; Updated today'),
                ('Best in slot lists', 'All specs &middot; 1.1 MB &middot; Updated today'), ('Talents', '9 classes &middot; 0.6 MB &middot; Updated yesterday'),
                ('Search index', '2.3 MB &middot; Updated today'), ('Item icons', '18.4 MB &middot; Updated today')]
        strip = Notice('Offline. Showing what you had at 4:12 pm. Sign-ups need a connection.', tone='neutral') if state == 'offline' else ''
        sec_body = Rows([Row(n, s, link=False) for n, s in rows], open_=True) if state != 'empty' else (
            '<p class="t-small" style="padding-top:12px">Nothing is saved on this phone yet. Guides and lists save here the first time you open them.</p>')
        body = (title('Offline data') + strip + Section('On this phone, 26.6 MB' if state != 'empty' else 'On this phone', sec_body)
                + Section('Updates', Rows([Row('Keep my classes ready', 'Saves the guide, lists and talents for every class you have a character in.', trail=Switch('Keep my classes ready'), link=False),
                                          Row('Download on Wi-Fi only', 'Icons and new guides wait for Wi-Fi.', trail=Switch('Download on Wi-Fi only'), link=False)], open_=True))
                + '<div class="stack">' + Button('Update now') + Button('Remove downloaded data', tone='danger') + '</div>'
                + '<p class="t-small" style="margin-top:16px">Checked today at 4:12 pm. Your characters, sign-ups and sims are not stored here; they load when you are online.</p>')
        return app_page(platform, page(body, 'Offline data'), 'home', back='Home', title='Offline data')
    return b


def update_required(platform: str) -> str:
    store = 'App Store' if platform == 'ios' else 'Google Play'
    body = (f'<div class="center"><img class="seal" src="../../../web/public/seal-400.png" alt="">' + title('Update Forever Sixty')
            + '<p class="t-small">This version is too old. Update to keep signing up and simming. Your characters and sign-ups are safe.</p>'
            + '<div class="stack">' + Button(f'Update in {store}', primary=True) + '</div></div>')
    return app_page(platform, page(body, 'Update Forever Sixty'), None, navbar=False)


def expired(platform: str) -> str:
    body = (f'<div class="center"><img class="seal" src="../../../web/public/seal-400.png" alt="">' + title('Sign in again')
            + '<p class="t-small">Your Battle.net link expired. Sign in again to bring your characters back.</p>'
            + '<div class="stack">' + Button('Sign in with Battle.net', primary=True) + '</div></div>')
    return app_page(platform, page(body, 'Sign in again'), 'home', navbar=False)


def raid_night(state: str):
    """R10 and R13: the raid page on raid night. live: starts in 42 min, you are In; signup: not answered yet."""
    def b(platform: str) -> str:
        clock = [('00', 'hours'), ('42', 'min')]
        you = Tag('You are In', 'kill') if state == 'live' else Tag('Not signed up', 'muted')
        cd = Countdown("Olympus XXVII &middot; Onyxia's Lair tonight", clock, f'<p class="t-small">Wed 6 Jan, 8:00 pm CST &middot; 18 in, 3 late</p>{you}', 'guild', live=True)
        buttons = ''.join(Button(x, primary=(x == 'In' and state == 'live'), small=True) for x in ('In', 'Late', 'Bench', 'Out'))
        counts = Rows([Row('Tanks 2 of 2', 'Grunvald, Brakka', link=False, slim=True),
                       Row('Healers 5 of 8', 'Mirethil, Selka, Hollowmend, Aenwyn, Brightwen &middot; 3 short', link=False, slim=True),
                       Row('Damage 11 of 30', 'Thandrel, Kaelra, Thalrin and 8 more', link=False, slim=True)], open_=True)
        ready = Checklist([('ok', 'Drakefire Amulet', 'Attuned to Onyxia&rsquo;s Lair', None),
                           ('ok', 'Elixir of the Mongoose', '3 in your bags, from the addon&rsquo;s last export', None),
                           ('miss', 'Flask of the Titans', 'None in your bags', 'Auction house'),
                           ('warn', 'Boots not enchanted', '', None)])
        body = (title("Onyxia's Lair", 'var(--strong)', icon('inv_misc_head_dragon_black', 56), ['Olympus XXVII', 'Wed 6 Jan, 8:00 pm CST']) + cd
                + Section('Your sign-up', f'<div class="la bare"><div class="btns">{buttons}</div></div>')
                + Section('Who is coming', counts, 'Roster &rsaquo;')
                + Section('Before you pull', ready + '<p class="t-small" style="margin-top:8px">From your last log-out, 9 h ago</p>'))
        return app_page(platform, page(body, "Onyxia's Lair"), 'guild', back='Guild', title="Onyxia's Lair", share=True)
    return b


def widgets(state: str):
    """R13. guild: Thandrel in Olympus XXVII before its first raid; noguild: Felsnap at 34 (first raids); open: raids are open."""
    def b(platform: str) -> str:
        at = raids_open()
        if state == 'guild':
            k, big, nm, s = "Olympus XXVII", '33d 8h', "Onyxia's Lair", 'Wed 9 Dec, 8:00 pm CST &middot; You are In'
            up = ('+14.5 DPS', 'next upgrade', 'Lionheart Helm, Blacksmithing')
        elif state == 'noguild':
            d = clock_to(at)
            k, big, nm, s = 'First raids', f'{int(d[0][0])}d {int(d[1][0])}h', day_text(at), 'Barrow Deeps, Hyjal Summit, Onyxia&rsquo;s Lair'
            up = ('+1.3 DPS', 'next upgrade', 'Felbough Arm, Ashenvale')
        else:
            k, big, nm, s = 'Raids open now', 'Open', "Barrow Deeps, Hyjal Summit, Onyxia's Lair", 'No raid planned yet'
            up = ('+14.5 DPS', 'next upgrade', 'Lionheart Helm, Blacksmithing')
        small = f'<div class="wd"><span class="k">{k}</span><span class="big">{big}</span><span class="nm">{nm}</span><span class="s">{s}</span></div>'
        small_up = f'<div class="wd"><span class="k">Upgrades</span><span class="big g">{up[0]}</span><span class="s">{up[1]}</span></div>'
        medium = (f'<div class="wd m"><div><span class="k">{k}</span><span class="big">{big}</span><span class="nm">{nm}</span><span class="s">{s}</span></div>'
                  f'<div><span class="k">Upgrades</span><span class="big g">{up[0]}</span><span class="s">{up[1]}</span><span class="s">{up[2]}</span></div></div>')
        lock = (f'<div class="lock"><span class="tm">9:41</span><div class="row3"><span class="lw">{glyph("guild", 18)}{big} &middot; {nm}</span>'
                f'<span class="lw">{glyph("bis", 18)}{up[0]}</span></div></div>')
        label = 'Home screen widgets' if platform == 'ios' else 'Widgets'
        body = (title('Widgets') + f'<p class="lbl">{label}: small</p><div class="wgrid">{small}{small_up}</div>'
                + f'<p class="lbl">Medium</p><div class="wgrid">{medium}</div><p class="lbl">Lock screen</p>{lock}')
        return app_page(platform, page(body, 'Widgets'), None, navbar=False)
    return b


def live_activity(platform: str) -> str:
    """R13: raid night. iOS: lock screen Live Activity and the Dynamic Island; Android: the ongoing notification."""
    buttons = ''.join(Button(x, primary=(x == 'In'), small=True) for x in ('In', 'Late', 'Bench', 'Out'))
    ony = icon('inv_misc_head_dragon_black', 36)
    if platform == 'ios':
        body = (title('Live Activity') + '<p class="lbl">Lock screen</p>'
                + f'<div class="la"><div class="top">{ony}<span><span class="nm">Onyxia&rsquo;s Lair</span><br><span class="t-small">Olympus XXVII &middot; You are In</span></span><span class="clk">42:00</span></div>'
                + '<span class="t-small">Starts 8:00 pm CST &middot; 18 in, 3 late &middot; 2 tanks, 5 heals</span></div>'
                + '<p class="lbl">Raid done</p><div class="la"><div class="top">' + ony + '<span><span class="nm">Onyxia&rsquo;s Lair done</span><br>'
                + '<span class="t-small">1 kill &middot; you got Deathbringer &middot; report ready</span></span></div></div>'
                + '<p class="lbl">Dynamic Island, compact</p><div class="island"><span>Ony</span><span class="g">42:00</span></div>'
                + f'<p class="lbl">Dynamic Island, expanded</p><div class="island x"><div class="top" style="display:flex;gap:10px;align-items:center">{ony}<span>Onyxia&rsquo;s Lair</span><span class="g" style="margin-left:auto">42:00</span></div>'
                + f'<div class="la bare"><div class="btns">{buttons}</div></div></div>'
                + '<p class="lbl">Sim running</p><div class="la"><div class="top"><span class="nm">Sim: Fury Warrior, 60 raid</span><span class="clk" style="font-size:16px">64%</span></div>'
                + Meter(0.64) + '</div>')
    else:
        body = (title('Ongoing notifications') + '<p class="lbl">Raid night</p>'
                + f'<div class="notif"><div class="top">{glyph("guild", 16)}Forever Sixty &middot; 42 min</div><div class="top">{ony}<span><span class="nm">Onyxia&rsquo;s Lair starts in 42 min</span><br><span class="t-small">Olympus XXVII &middot; You are In &middot; 18 in, 3 late</span></span></div>'
                + f'<div class="la bare"><div class="btns">{buttons}</div></div></div>'
                + '<p class="lbl">Sim running</p><div class="notif"><div class="top">Forever Sixty</div><span class="nm">Sim: Fury Warrior, 60 raid</span>' + Meter(0.64) + '<span class="t-small">64%</span></div>')
    return app_page(platform, page(body, 'Live Activity'), None, navbar=False)


def search(state: str):
    def b(platform: str) -> str:
        field = f'<div class="searchf">{glyph("more", 18)}<span>{"Thal" if state == "results" else "Characters, items, spells"}</span></div>'
        if state == 'results':
            rows = Section('Characters', Rows([Row('Thandrel', 'Level 60 Fury Warrior &middot; Olympus XXVII', thandrel_crest(40), title_color=WARRIOR),
                                              Row('Thalrin', 'Level 60 Fury Warrior &middot; Olympus XXVII', crest_img('warrior', 40), title_color=WARRIOR)]))
            rows += Section('Items', Rows([Row('Thalassian Pauldrons', 'Shoulder &middot; Molten Core', icon('inv_shoulder_15', 40, 'q4'), title_color='#b866f5')]))
        else:
            rows = '<p class="t-small" style="margin-top:20px">Find a character by name, or an item to see where it drops and who wants it.</p>'
        body = field + rows
        return app_page(platform, page(body, 'Search'), None, navbar=False)
    return b


# ------------------------------------------------------------------------------------------------ the catalogue
# id: (builder(platform) -> html). Order follows the build order (spec section 7).
def W(mock, tab, **kw):
    return web_screen(mock, tab, **kw)


SCREENS: dict = {
    # Home (tab Home)
    'home-out': W('web-home-out', 'home', css=HOME_OUT + FEED_FILTERS, inject=_out_footer(), who=None),
    'expired': expired,
    'search': search('empty'),
    'search-results': search('results'),
    # Character (pushed from Home's head or You)
    'character-owner-34': W('web-character-owner-34', 'home', back='Home', title='Felsnap', share=True),
    'character-owner-60': W('web-character-owner-60-premium', 'home', back='Home', title='Thandrel', share=True),
    'character-visitor-60': W('web-character-visitor-60', 'home', back='Search', title='Elyra Duskvale', share=True),
    'character-no-score': W('web-character-no-score', 'home', back='Home'),
    'character-not-found': W('web-character-not-found', 'home', back='Search'),
    'character-loading': W('web-character-loading', 'home', back='Home'),
    'character-error': W('web-character-error', 'home', back='Home'),
    # BiS (tab BiS): the landing answers for your character; the spec list is the paperdoll page
    'bis-spec-60': W('web-bis-dps-60', 'gear', back='Gear', title='Marksmanship', share=True),
    'bis-spec-healer-60': W('web-bis-healer-60', 'gear', back='Gear', title='Holy', share=True),
    'bis-spec-20': W('web-bis-low-20', 'gear', back='Gear', title='Beast Mastery', share=True),
    # Guides (tab Guides): Guide | Tiers | What changed segments
    'guides': W('web-guides', 'home', back='Home', title='Guides'),
    'guides-you': W('web-guides-you', 'home', back='Home', title='Guides'),
    'guides-class': W('web-guides-hunter', 'home', back='Guides', title='Hunter'),
    'guide': W('web-guides-hunter-survival', 'home', back='Hunter', title='Survival', share=True),
    'guide-you': W('web-guides-hunter-survival-you', 'home', back='Hunter', title='Survival', share=True),
    'guide-loading': W('web-guides-hunter-survival-loading', 'home', back='Hunter'),
    'guide-error': W('web-guides-hunter-survival-error', 'home', back='Hunter'),
    'tiers': W('web-tiers', 'home', back='Home', title='Tiers'),
    'tiers-you': W('web-tiers-you', 'home', back='Home', title='Tiers'),
    'changes': W('web-changelog-game', 'home', back='Home', title='What changed'),
    # Sims, talents and the planner (from Home's Next rows)
    'sim': W('web-sim', 'home', back='Home'),
    'sim-you': W('web-sim-you', 'home', back='Home'),
    'sim-running': W('web-sim-running', 'home', back='Sims'),
    'sim-result': W('web-sim-result', 'home', back='Sims', title='Sim result', share=True),
    'sim-saved': W('web-sim-saved', 'home', back='Sims'),
    'planner-you': W('web-planner-you', 'home', back='Home', title='Talents', share=True),
    'planner-guide': W('web-planner-guide', 'home', back='Home', title='Talents', share=True),
    'planner-loading': W('web-planner-loading', 'home', back='Home'),
    'planner-error': W('web-planner-error', 'home', back='Home'),
    'build-shared': W('web-b', 'home', back='Home', share=True),
    # Logs, reports and rankings (You > Logs; Home's Next rows)
    'reports': W('web-reports-list', 'home', back='Home'),
    'reports-empty': W('web-reports-empty', 'home', back='Home'),
    'reports-failed': W('web-reports-failed', 'home', back='Home'),
    'report': W('web-report-summary', 'home', back='Logs', share=True),
    'report-player': W('web-report-player', 'home', back='Report', share=True),
    'report-coach': W('web-report-coach', 'home', back='Report'),
    'report-processing': W('web-report-processing', 'home', back='Logs'),
    'report-loading': W('web-report-loading', 'home', back='Logs'),
    'report-error': W('web-report-error', 'home', back='Logs'),
    'rankings': W('web-rankings', 'home', back='Home'),
    'rankings-you': W('web-rankings-you', 'home', back='Home'),
    'rankings-empty': W('web-rankings-empty', 'home', back='Home'),
    # Guild (tab Guild) and raids
    'guild-raids': W('web-guild-raids', 'guild', back='Guild'),
    'guild-raid': W('web-guild-raid', 'guild', back='Guild', share=True),
    'guild-roster': W('web-guild-roster', 'guild', back='Guild'),
    'guild-loot': W('web-guild-loot', 'guild', back='Guild'),
    'raid-night-live': raid_night('live'),
    'raid-night-signup': raid_night('signup'),
    # You, account, premium, system
    'account': W('web-account', 'home', back='Home'),
    'account-premium': W('web-account-premium', 'home', back='Home'),
    'account-loading': W('web-account-loading', 'home', back='Home'),
    'account-error': W('web-account-error', 'home', back='Home'),
    'premium': W('web-premium-free', 'home', back='Plan'),
    'devices': devices,
    'delete-free': delete('free'),
    'delete-premium': delete('premium'),
    'delete-failed': delete('failed'),
    'delete-done': delete('done'),
    'offline-data': offline('ready'),
    'offline-data-empty': offline('empty'),
    'offline-strip': offline('offline'),
    'update-required': update_required,
    'widgets-guild': widgets('guild'),
    'widgets-noguild': widgets('noguild'),
    'widgets-open': widgets('open'),
    'live-activity': live_activity,
    **flows.SCREENS,
    **guild_tab.SCREENS,
}


def pack(folder: str = 'design/reviews/2026-11-07-mobile') -> None:
    """Shrinks the captures to a 256-colour palette (about a third of the size; the gate reads region files, not pixels)."""
    from PIL import Image
    shots = ROOT / folder
    for p in sorted(shots.glob('*.png')):
        im = Image.open(p)
        if im.mode != 'P':
            im.convert('RGB').quantize(256, method=Image.Quantize.MEDIANCUT, dither=Image.Dither.NONE).save(p, optimize=True)


def write_ignore() -> None:
    """The web-derived mocks are rebuilt from design/mocks/web on every run, so only the app-only HTML is committed."""
    names = sorted(f'{p}-{sid}.html' for sid, b in SCREENS.items() if getattr(b, 'web', False) for p in DEVICES)
    (ROOT / 'design/mocks/app/.gitignore').write_text('# generated by design/mocks/ws_app.py: web-derived screens, rebuilt from design/mocks/web\n' + '\n'.join(names) + '\n')


def main(filters: list[str]) -> None:
    if filters[:1] == ['pack']:
        pack(*filters[1:2])
        return
    write_ignore()
    n = 0
    for sid, build in SCREENS.items():
        if filters and not any(f in sid for f in filters):
            continue
        for platform in DEVICES:
            write_app(f'{platform}-{sid}.html', build(platform))
            n += 1
    print(f'{n} app mocks in design/mocks/app')


if __name__ == '__main__':
    main(sys.argv[1:])
