"""Review boards for Me with no site activity yet (design/specs/2026-11-14-me-empty.md).

    python3 design/mocks/addon_hub_me_empty.py              # every board
    python3 design/mocks/addon_hub_me_empty.py me-empty-l25

Writes design/reviews/2026-11-14-me-empty-and-sources/<board>.png at device scale 2, the window at its real 740 px
width. As the L2 boards (addon_hub_me.py), every word and figure inside the tab comes from the addon's own model:
each state's inputs, the client's facts in `own` as MeTab.readOwn gathers them, go through
design/mocks/addon_me_model.lua into MeTab.model, and the JSON it returns is drawn on MeFrame.LAYOUT. Set up is
Onboarding.steps' own model, drawn by addon_hub_l8.

Made up: Velisandre, a hunter. At 25 she follows the Beast Mastery leveling build and wears band 20 gear whose
gains are the shipped list's own (the pick against the runner-up she wears); at 60 she is fresh, with the Raid-ready
list's gains (addon_hub.slot_rows). The talent named next is a Beast Mastery talent from Data.lua.
"""
import re
import sys

import addon_builds_window as bw
import addon_hub as hub
import addon_hub_l8 as l8
import addon_hub_me as me
from addon_hub import ADDON, ROOT, esc, icon, item_icon

OUT = ROOT / 'design/reviews/2026-11-14-me-empty-and-sources'
G = me.LAYOUT
MONKEY = 'Interface\\Icons\\ability_hunter_aspectofthemonkey'

CSS = f'''
.me .st {{ position: relative; height: {G['partRow']}px; }}
.me .st > * {{ position: absolute; top: {(G['partRow'] - G['line']) // 2}px; white-space: nowrap; }}
.me .st .trk {{ top: {(G['partRow'] - G['track']) // 2}px !important; }}
.me .st .ic16 {{ top: {(G['partRow'] - G['gainIcon']) // 2}px; line-height: 0; }}
.me .ench {{ height: 18px; display: flex; align-items: center; gap: 6px; white-space: nowrap; width: 288px; }}
.me .ench .et {{ min-width: 0; overflow: hidden; text-overflow: ellipsis; }}
.me .site {{ color: {ADDON['muted']}; width: 388px; }}
'''


SLOTS = {'head': 'Head', 'neck': 'Neck', 'shoulder': 'Shoulder', 'back': 'Back', 'chest': 'Chest', 'wrist': 'Wrist',
         'hands': 'Hands', 'waist': 'Waist', 'legs': 'Legs', 'feet': 'Feet', 'finger1': 'Finger', 'finger2': 'Finger',
         'trinket1': 'Trinket', 'trinket2': 'Trinket', 'main_hand': 'Main hand', 'off_hand': 'Off hand',
         'ranged': 'Ranged'}
KIND_WORDS = {'crafted': 'Crafted', 'quest': None, 'world_drop': None, 'vendor': None, 'rep': None}


def where(slot, s):
    """GearTab.sourceWhere's words: the source as the tooltip says it (ItemTooltip.sourceParts), no slot."""
    source, kind, chance = s.get('source') or '', s.get('source_kind'), s.get('chance')
    if kind in ('dungeon', 'raid') and ': ' in source:
        place, boss = source.split(': ', 1)
        text = f'{place} · {boss}' + (f' · {round(chance * 100)}% drop' if chance else '')
    elif kind == 'crafted':
        text = f'Crafted · {source}'
    else:
        text = source + (f' · {round(chance * 100)}% drop' if chance else '')
    return text


def band_rows(band, faction='alliance'):
    """The band's slots with the shipped list's gain of its pick over the last runner-up (what she wears)."""
    spec = next(b for b in hub.BIS['bands'] if b['band'] == band and b['faction'] == faction)
    rows = []
    for s in spec['slots']:
        alts = s.get('alternatives') or []
        if s.get('item_id') and alts:
            rows.append({'itemId': s['item_id'], 'gain': -alts[-1]['dps_delta'], 'where': where(s['slot'], s)})
    return rows


ENCHANT_ICON = 'spell_holy_greaterheal'


def enchant(slot, name):
    return {'slot': slot, 'missing': True, 'plan': {'id': 1, 'name': name, 'icon': ENCHANT_ICON}}


def gear20():
    rows = sorted(band_rows(20), key=lambda r: -r['gain'])[:6]
    return {'band': 20, 'best': 8, 'scored': 14,
            'enchants': [enchant('chest', 'Enchant Chest - Minor Stats'), enchant('wrist', 'Enchant Bracer - Minor Agility')],
            'upgrades': [{'itemId': r['itemId'], 'gain': f'+{r["gain"]:.1f} DPS', 'where': r['where']} for r in rows]}


def gear60():
    band = {s['slot']: s for s in hub.RAID60['slots']}
    rows = sorted((r for r in hub.slot_rows().values() if r['gain'] > 0), key=lambda r: -r['gain'])
    return {'band': 60, 'best': 16 - len(rows), 'scored': 16,
            'enchants': [enchant('back', 'Enchant Cloak - Lesser Agility')],
            'upgrades': [{'itemId': r['pick'], 'gain': f'+{r["gain"]:.1f} DPS', 'where': where(r['slot'], band[r['slot']])}
                         for r in rows]}


BUILD25 = {'name': 'Beast Mastery leveling', 'matched': 16, 'allowed': 16,
           'next': {'name': 'Improved Aspect of the Monkey', 'icon': MONKEY, 'level': 26}}
OWN25 = {'level': 25, 'maxLevel': 60, 'xp': 8140, 'xpMax': 13100, 'rested': 5240, 'build': BUILD25, 'nextBand': 30,
         'gear': gear20(), 'professions': [{'name': 'Skinning', 'rank': 112}, {'name': 'Leatherworking', 'rank': 105}]}
BUILDS_LEAD = {'tab': 'builds', 'label': 'Builds', 'waiting': [],
               'lead': {'icon': 'ability_hunter_beasttaming', 'line': 'Beast Mastery leveling',
                        'sub': '16 of 16 points on plan'}}
QUIET = dict(record=None, lastRaid=None, sim=None, raidName=None, syncedAgo=None, companion=False, guild=None,
             upgrades=[], pack={'available': True, 'generated': '2026-12-09', 'stale': False})
SPEND = {'tab': 'builds', 'label': 'Builds', 'lead': BUILDS_LEAD['lead'],
         'waiting': [{'key': 'spend', 'line': '|cffffd100Spend 1 point|r: Improved Aspect of the Monkey',
                      'short': '|cffffd100Spend 1 point|r'}]}
SETUP5 = dict(l8.SETUP, example=None, gearOpen=False)
_STEP_ICON = l8.step_icon


def step_icon(row):
    """A step with neither item nor icon draws the client's unknown icon, as paintCell does (Theme.UNKNOWN_ICON)."""
    if not (row.get('crest') or row.get('itemId') or row.get('icon')):
        return icon('inv_misc_questionmark', G['nowIcon'], 'box-shadow:0 0 0 1px #000')
    return _STEP_ICON(row)


l8.step_icon = step_icon


def STATES():
    base = dict(me.BASE, **QUIET)
    return {k: {a: b for a, b in v.items() if b is not None} for k, v in {
        'l25': dict(base, level=25, own=OWN25, leads=[BUILDS_LEAD], buildName='Beast Mastery leveling'),
        'l25-point': dict(base, level=26, companion=True, syncedAgo='3 min ago', buildName='Beast Mastery leveling',
                          own=dict(OWN25, level=26, xp=410, xpMax=13700, rested=None,
                                   build=dict(BUILD25, allowed=17, next=dict(BUILD25['next'], level=None, now=True))),
                          leads=[SPEND]),
        'l5': dict(base, level=5, own={'level': 5, 'maxLevel': 60, 'xp': 1040, 'xpMax': 2800, 'rested': 700,
                                       'nextBand': 20}, spec=None, buildName=None,
                   leads=[{'tab': 'builds', 'label': 'Builds', 'waiting': [{'key': 'spec', 'line': 'Pick your spec'}]}]),
        'l5-setup': dict(base, level=5, own={'level': 5, 'maxLevel': 60, 'xp': 1040, 'xpMax': 2800, 'nextBand': 20},
                         spec=None, buildName=None, leads=[], setupInput=SETUP5),
        'l60': dict(base, level=60, companion=True, syncedAgo='14 min ago', guild='Thornwake Company',
                    buildName='Beast Mastery raid build',
                    own={'level': 60, 'maxLevel': 60, 'gear': gear60(),
                         'professions': [{'name': 'Engineering', 'rank': 300}, {'name': 'Skinning', 'rank': 300}],
                         'build': {'name': 'Beast Mastery raid build', 'matched': 51, 'allowed': 51}},
                    leads=[me.LEADS['guild'], me.LEADS['builds']]),
    }.items()}


def standing_html(m):
    if not m:
        return ''
    out = me.section('Where you stand')
    for r in m['rows']:
        x, cells = G['partLabel'], f'<span class="m" style="left:0">{esc(r["label"])}</span>'
        if r.get('fill') is not None:
            fill = round(G['trackMax'] * r['fill'])
            cells += (f'<span class="trk" style="left:{x}px;width:{G["trackMax"]}px"><i style="width:{fill}px;'
                      f'background:{ADDON["body"]}"></i></span>')
            x += G['trackMax'] + G['reasonGap']
        if r.get('icon'):
            name = r['icon'].split('\\')[-1].lower()
            cells += f'<span class="ic16" style="left:{x}px">{icon(name, G["gainIcon"], "box-shadow:0 0 0 1px #000")}</span>'
            x += G['gainIcon'] + G['gainGap']
        color = ADDON['muted'] if r.get('muted') else ADDON['body']
        cells += f'<span style="left:{x}px;color:{color}">{esc(r["value"])}</span>'
        if r.get('note'):
            cells += f'<span class="m" style="right:0">{esc(r["note"])}</span>'
        out += f'<div class="st">{cells}</div>'
    return out + '<div class="gap"></div>'


def short_words(e):
    """GearEnchants.shortWords: "Chest: Minor Stats"."""
    name = re.sub(r'^Enchant [^-]+ - ', '', e['plan']['name'])
    return f'{SLOTS[e["slot"]]}: {name}'


def gear_html(m, width):
    """Each upgrade as a Right now row (MeFrame drawLocalGear, `width` wide), then GearEnchants' rows."""
    if not m:
        return ''
    out = me.section('Gear', m.get('note'))
    out += ''.join(f'<div class="now"><span class="ni">{item_icon(g["itemId"], G["nowIcon"])}</span>'
                   f'<div class="tx" style="width:{width - G["nowIcon"] - G["nowGap"]}px"><div style="width:'
                   f'{width - G["nowIcon"] - G["nowGap"] - G["gainValue"]}px">{hub.item_name(g["itemId"])}</div>'
                   f'<div class="m">{me.wow(g.get("where") or "")}</div></div>'
                   f'<span class="lk ok">{esc(g["gain"])}</span></div>' for g in m['rows'])
    if m['rows']:
        out += f'<div style="height:{4 - (G["nowRow"] - 2 * G["line"])}px"></div>'
    for e in m.get('enchants') or []:
        out += (f'<div class="ench">{icon(e["plan"]["icon"], 16, "box-shadow:0 0 0 1px #000")}'
                f'<span class="et" style="color:{ADDON["warning"]}">{esc(short_words(e))}</span></div>')
    if m.get('empty'):
        out += f'<div class="m" style="margin-top:2px">{esc(m["empty"])}</div>'
    return out + '<div class="gap"></div>'


def site_html(m):
    if not m:
        return ''
    link = f'<div class="link" style="margin-top:4px">{esc(m["link"])}</div>' if m.get('link') else ''
    return f'<div class="site">{esc(m["line"])}</div>{link}'


def board(name, caption, model, level, dots, tabs):
    figures = me.SEP.join([f'Level {level} Hunter'] + model['figures'])
    left = standing_html(model.get('standing'))
    right = (l8.setup_html(model['setup']) if model.get('setup') else me.now_html(model.get('rightNow')))
    # MeFrame.draw: Gear joins the shorter column, the sim counted (heights as the frame lays them out).
    rows, now, sim = len((model.get('standing') or {}).get('rows', [])), len(model.get('rightNow') or []), model.get('sim')
    left_h = 26 + rows * G['partRow'] + G['sectionGap'] if rows else 0
    right_h = (26 + now * G['nowRow'] + G['sectionGap'] - (G['nowRow'] - 2 * G['line']) if now else 0) + (62 if sim else 0)
    if left_h < right_h:
        left += gear_html(model.get('gear'), 388)
    else:
        right += gear_html(model.get('gear'), 288)
    left += site_html(model.get('site'))
    right += me.sim_html(sim)
    saved = bw.TABS
    bw.TABS = tabs
    try:
        win = bw.window('hunter', 'Velisandre', esc(figures), f'<div class="me">{left}</div>',
                        f'<div class="me">{right}</div>', 'Esc to close', me.wow(model.get('footer') or ''),
                        strip=('Me', dots))
    finally:
        bw.TABS = saved
    return bw.ux.page(f'<style>{bw.CSS}{me.CSS}{l8.CSS}{CSS}</style><div title="{esc(caption)}">{win}</div>')


BOARDS = {
    'me-empty-l25': ('l25', 25, 'Level 25, a build followed, no companion: where she stands, band 20 gear, the app.',
                     (), ['Me', 'Builds', 'Gear', 'Rotation']),
    'me-empty-l25-point': ('l25-point', 26, 'Level 26 with a point waiting, the companion running, nothing logged yet.',
                           ('Builds',), ['Me', 'Builds', 'Gear', 'Rotation']),
    'me-empty-l5': ('l5', 5, 'Level 5, no spec yet, no companion.', ('Builds',), ['Me', 'Builds']),
    'me-empty-l60': ('l60', 60, 'Level 60 fresh in a guild, the companion running, no logs yet.', ('Guild',),
                     ['Me', 'Builds', 'Gear', 'Rotation', 'Guild']),
}


def main(names):
    import os
    from playwright.sync_api import sync_playwright
    me.STATES = STATES()
    data = me.models()
    OUT.mkdir(parents=True, exist_ok=True)
    with sync_playwright() as pw:
        exe = os.environ.get('PLAYWRIGHT_CHROMIUM')
        browser = pw.chromium.launch(executable_path=exe) if exe else pw.chromium.launch()
        page = browser.new_context(device_scale_factor=2, viewport={'width': 1400, 'height': 1000}).new_page()
        for name in names:
            key, level, caption, dots, tabs = BOARDS[name]
            page.set_content(board(name, caption, data[key], level, dots, tabs))
            page.evaluate('document.fonts.ready')
            page.wait_for_timeout(150)
            page.locator('#board').screenshot(path=str(OUT / f'{name}.png'))
            print(OUT / f'{name}.png')
        browser.close()


if __name__ == '__main__':
    main(sys.argv[1:] or list(BOARDS))
