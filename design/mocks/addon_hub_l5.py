"""Lane L5 review renders: the Guild tab as built (addon/ForeverSixty/surfaces/GuildHubTab.lua, GuildRaid.lua).

    python3 design/mocks/addon_hub_l5.py            # every state
    python3 design/mocks/addon_hub_l5.py l5-raid-day

Writes design/reviews/2026-11-13-addon-hub/l5/<state>.png at device scale 1: the window at its real 740 px
width, drawn with addon_hub.py's frame and tokens, but holding exactly what the built tab draws from the
data it has. Where the approved board (hub-guild.png) shows a fact no message carries, the render leaves
it out as the build does: no raid leader and no invite time (raid_signup and raid_roster carry neither),
no per-raid progression in the guild line (the ratings pack's guild record has standing and nights; each
raider's attendance is the table's column). Section headings are in the
title face, as the Builds tab's (Theme.titleFont at windowSection 14).

States: raid day not signed up (the dot on Guild, the answers blue); answered and queued (Late teal, the
footer's Sends when the companion runs); an older companion (no raid_roster: Sign-ups left out); no
companion; no ratings pack. Every character and guild is made up (Velisandre, Thornwake Company).
"""
import os
import sys

import addon_hub as hub
from addon_hub import (ADDON, CLASS_HEX, CLIENT_GOLD, GUILD, ROSTER, SPEC_ICON, TIER_HEX, esc, icon, icon_uri,
                       item_name, section, tier_of, window)

OUT = hub.OUT / 'l5'
TITLE_CSS = f'''
.sh > span:first-child {{ font-family: Title, serif; font-size: 14px; font-weight: 400; color: {ADDON['cream']};
  text-shadow: 1px -1px 0 #000; }}

.prep .ic {{ display: inline-block; vertical-align: -3px; margin-right: 4px; }}
.prep td:first-child {{ width: 132px; color: {ADDON['muted']}; }}
.nx td:nth-child(1) {{ width: 112px; }} .nx td:nth-child(2) {{ width: 98px; }}
'''
RANKS = (1, 2, 3)  # the raider ranks the officers named on the site
TODAY = 'Thu 10 Dec'


def name(n, cls):
    return f'<span class="cls" style="color:{CLASS_HEX[cls]}">{esc(n)}</span>'


def tonight(answer='none', queued=None, accepts=True):
    words = {'none': ('Not signed up', 'cg', ''), 'in': ('In', 'ok', ' <span class="m">as DPS, sent Tue 19:04</span>'),
             'late': ('Late', 'cg', ' <span class="m">as DPS, not sent yet</span>')}
    shown = queued or answer
    word, cls, detail = words[shown]
    btns = ''.join(f'<span class="btn{" sel" if a.lower() == shown else ""}">{a}</span>'
                   for a in ('In', 'Late', 'Bench', 'Out'))
    body = (f'<div class="row">{icon("inv_misc_head_dragon_black", 32)}<div style="flex:1">'
            f'<div class="s" style="font-size:14px;font-weight:600">Onyxia\'s Lair</div>'
            f'<div class="m">{TODAY}, 20:00 local · in 3 h 12 min</div></div></div>'
            f'<div style="margin-top:6px">Your sign-up: <span class="{cls}">{word}</span>{detail}</div>')
    body += f'<div class="signup">{btns}</div>' if accepts else '<div class="m">Change it on the site</div>'
    body += ('<div style="margin-top:4px">Your job: <span class="cg">Phase 2 ranged, west side</span></div>'
             f'<div class="m">Fire on the whelps from the west wall when she lifts off. With {name("Mirelle", "mage")}.</div>')
    return section('Tonight', body)


def prep():
    rows = (f'<tr><td>Attunement</td><td>{hub.item_icon(16309, 14)} {item_name(16309)} '
            '<span class="ok">in your bags</span></td></tr>'
            '<tr><td>Fire resistance</td><td>42</td></tr>'
            '<tr><td>In your bags</td><td>Flask 2 · Food 20</td></tr>'
            '<tr><td>Missing</td><td class="wa">Greater Fire Protection Potion</td></tr>')
    return section('Your prep', f'<table class="tbl prep">{rows}</table>', 'for Onyxia')


def signups():
    line = ('<span class="ok">11 In</span> · <span class="cg">1 Late</span> · 2 Bench · <span class="wa">1 Out</span>'
            ' · 16 not signed up')
    roles = ' · '.join(f'<span class="{"ok" if c >= w else "wa"}">{n} {c} of {w}</span>'
                       for n, c, w in (('Tanks', 2, 2), ('Healers', 4, 8), ('Melee', 2, 10), ('Ranged', 4, 20)))
    return section('Sign-ups', f'<div>{line}</div><div>{roles}</div>', '12 of 40 coming')


def next_raids(first=('Not signed up', 'cg')):
    rows = [(f'{TODAY} 20:00', 'Onyxia\'s Lair', *first, True), ('Sun 13 Dec 19:30', 'Molten Core', 'Not signed up', 'cg', False),
            ('Tue 15 Dec 19:30', 'Molten Core', 'Not signed up', 'cg', False)]
    body = ''.join(f'<tr><td class="{"" if shown else "m"}">{d}</td><td class="s">{n}</td><td class="{c}">{a}</td></tr>'
                   for d, n, a, c, shown in rows)
    return section('Next raids', f'<table class="tbl nx">{body}</table>')


def guild_block(stale=False):
    raiders = [r for r in ROSTER if r[0] != 'Pell']
    ranked = sorted(raiders, key=lambda r: -r[3])
    place = next(i for i, r in enumerate(ranked, 1) if r[0] == 'Velisandre')
    def row(i, r, own=''):
        return (f'<tr{own}><td class="m n" style="width:12px">{i}</td><td>'
                f'<span style="display:inline-block;vertical-align:-3px;margin-right:4px">{hub.crest(r[1], 14)}</span>'
                f'{name(r[0], r[1])}</td><td class="n" style="color:{TIER_HEX[tier_of(r[3])]}">{r[3]}</td>'
                f'<td class="n m">{r[4]}%</td></tr>')
    body = ''.join(row(i, r) for i, r in enumerate(ranked[:5], 1)) + row(place, ranked[place - 1],
                                                                          ' style="border-top:4px solid transparent"')
    line = '3rd of all guilds · ' + ('<span class="wa">Ratings 8 days old</span>' if stale
                                     else '41 raid nights in 90 days')
    table = ('<table class="tbl"><tr><th class="n">#</th><th>Raider</th><th class="n">Score</th><th class="n">Att.</th></tr>'
             f'{body}</table>')
    return section(GUILD, f'<div class="m" style="margin-bottom:4px">{line}</div>{table}',
                   f'you are {place}th of {len(raiders)} raiders')


def sent(following=True):
    tail = ' · you follow it' if following else ''
    button = '' if following else '<span class="btn" style="min-width:64px;justify-content:center">Follow</span>'
    return section('Officers sent', f'<div class="row" style="align-items:flex-start">{icon("ability_hunter_beasttaming", 20)}'
                   f'<div style="flex:1"><div>Beast Mastery raid build</div><div class="m">from {name("Brannoc", "warrior")},'
                   f' Tue{tail}</div></div>{button}</div>')


def hover_tip():
    return ('<div class="tip" style="position:absolute;left:640px;top:200px;width:auto">'
            '<div class="t">Molten Core</div><div style="color:#b8b8b8">22 of 40 coming</div></div>')


def win(left, right, foot='', dots=('Guild',)):
    return window('Guild', 'Level 60 Hunter · Thornwake Company · Officer', left, right, ('Esc to close', foot),
                  dots=dots)


STATES = {
    'l5-raid-day': ('Raid day, not signed up: the Guild dot, the four answers, the job, prep, sign-ups.',
                    lambda: win(tonight() + prep() + signups(), next_raids() + guild_block() + sent())),
    'l5-queued': ('Late pressed: teal, queued for the companion, the footer says when it sends; no dot.',
                  lambda: win(tonight('none', 'late') + prep() + signups(),
                              next_raids(('Late as DPS', 'cg')) + guild_block() + sent(),
                              'Sends when the companion runs', dots=())),
    'l5-old-companion': ('A companion older than raid_roster and without answers: Sign-ups left out, change it on '
                         'the site; the pack is 8 days old.',
                         lambda: win(tonight('in', accepts=False) + prep(),
                                     next_raids(('In as DPS', 'ok')) + guild_block(stale=True) + sent(), dots=())),
    'l5-follow-hover': ('Not following the pushed build: Follow beside it. The mouse on the Sun 13 Dec row: its '
                        'sign-ups in the hover; a click puts Molten Core on top.',
                        lambda: win(tonight() + prep() + signups(),
                                    next_raids() + guild_block() + sent(following=False)).replace(
                                        '<div class="ft">', hover_tip() + '<div class="ft">')),
    'l5-no-companion': ('No companion: Tonight says so; the guild block stays from the ratings pack.',
                        lambda: win(section('Tonight', '<div class="m">No raid from the site yet.</div>'),
                                    guild_block() + sent(), dots=())),
    'l5-no-pack': ('A guild the site has no raids for, and no ratings pack: the top five and your place left out.',
                   lambda: win(section('Tonight', f'<div class="m">No raids from {GUILD} on the site yet.</div>'), sent(),
                               dots=())),
}


def main(names):
    from playwright.sync_api import sync_playwright
    OUT.mkdir(parents=True, exist_ok=True)
    with sync_playwright() as pw:
        exe = os.environ.get('PLAYWRIGHT_CHROMIUM')
        browser = pw.chromium.launch(executable_path=exe) if exe else pw.chromium.launch()
        page = browser.new_context(device_scale_factor=1, viewport={'width': 900, 'height': 900}).new_page()
        for key in names:
            caption, build = STATES[key]
            html = hub.board2(f'<b>Guild, as built</b> · {caption}', build()).replace('</style>', TITLE_CSS + '</style>', 1)
            page.set_content(html)
            page.evaluate('document.fonts.ready')
            page.wait_for_timeout(150)
            page.locator('#board').screenshot(path=str(OUT / f'{key}.png'))
            print(OUT / f'{key}.png')
        browser.close()


if __name__ == '__main__':
    main(sys.argv[1:] or list(STATES))
