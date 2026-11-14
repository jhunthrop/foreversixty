"""Lane L8 review renders: Me's first open as built (addon/ForeverSixty/surfaces/Onboarding.lua, TourExample.lua,
MeFrame.lua), the window that replaces the 640 tour.

    python3 design/mocks/addon_hub_l8.py            # every state
    python3 design/mocks/addon_hub_l8.py l8-first-open

Writes design/reviews/2026-11-13-addon-hub/l8/<state>.png at device scale 2, the window at its real 740 px width.
As the L2 boards (addon_hub_me.py), every word and figure in the tab comes from the addon's own model: Set up's
rows from Onboarding.steps and the rest from MeTab.model, through design/mocks/addon_me_model.lua, drawn on
Onboarding.LAYOUT and MeFrame.LAYOUT. The tooltip under the steps is the client's own GameTooltip of the worn item
in game (TourExample shows it with SetInventoryItem); here it is drawn from the build's item data in the client's
words (the hit line as the client prints it, design/specs/2026-11-13-addon-hub.md 13, L9 data checks), with the
one line the item hook adds for that item: Next upgrade, the pick Gear's top upgrade names.

Made up, as on the acceptance boards: Velisandre, a level 60 hunter in Thornwake Company.
"""
import re
import sys

import addon_builds_window as bw
import addon_hub as hub
import addon_hub_me as me
from addon_hub import ADDON, ROOT, esc, icon, item_icon

OUT = ROOT / 'design/reviews/2026-11-13-addon-hub/l8'
SOURCE = (ROOT / 'addon/ForeverSixty/surfaces/Onboarding.lua').read_text()
STEP = {k: int(v) for k, v in re.findall(r'(\w+) = (\d+)', SOURCE[SOURCE.index('Onboarding.LAYOUT = {'):
                                                                SOURCE.index('}', SOURCE.index('Onboarding.LAYOUT = {'))])}
G = me.LAYOUT
SPECS = [{'slug': 'hunter-beast-mastery', 'name': 'Beast Mastery'}, {'slug': 'hunter-marksmanship', 'name': 'Marksmanship'},
         {'slug': 'hunter-survival', 'name': 'Survival'}]
# TourExample.choose: the worn item Gear's top upgrade replaces (Trueaim Gauntlets, the list's pick Gloves of Holy
# Might), the same order the Gear tab's Upgrades use (hub.slot_rows by gain).
TOP = sorted((r for r in hub.slot_rows().values() if r['gain'] > 0), key=lambda r: -r['gain'])[0]
WORN = hub.WORN[TOP['slot']]
EXAMPLE = {'slot': TOP['slot'], 'slotId': 10, 'itemId': WORN, 'name': hub.ITEMS[WORN]['name']}
SETUP = {'specs': SPECS, 'classToken': 'HUNTER', 'example': EXAMPLE, 'gearOpen': True}

CSS = f'''
.su {{ position: relative; height: {STEP['row']}px; }}
.su .no {{ position: absolute; left: 0; top: {(G['nowIcon'] + 4 - G['line']) // 2}px; color: {ADDON['muted']}; }}
.su .tk {{ position: absolute; left: 0; top: {(G['nowIcon'] + 4 - STEP['tick']) // 2}px; width: {STEP['tick']}px;
  height: {STEP['tick']}px; }}
.su .ni {{ position: absolute; left: {STEP['number']}px; top: 2px; line-height: 0; }}
.su .tx {{ position: absolute; left: {STEP['number'] + G['nowIcon'] + G['nowGap']}px;
  width: {288 - STEP['number'] - G['nowIcon'] - G['nowGap'] - G['nowLink']}px; }}
.su .tx div {{ white-space: nowrap; overflow: hidden; text-overflow: ellipsis; }}
.su .tx div + div {{ width: {288 - STEP['number'] - G['nowIcon'] - G['nowGap']}px; }}
.su .lk {{ position: absolute; right: 0; top: 0; color: {ADDON['gold']}; }}
.su.done .tx div:first-child {{ color: {ADDON['muted']}; }}
.su.done .ni img, .su.done .ni span {{ filter: grayscale(1); }}
.gt {{ margin-top: {STEP['tooltipGap']}px; width: 288px; box-sizing: border-box; background: rgba(8, 10, 22, .95);
  border: 1px solid #9095a3; border-radius: 4px; padding: 7px 10px 8px; font-family: Site, sans-serif; font-size: 12px;
  line-height: 16px; color: #fff; }}
.gt .t {{ font-size: 13px; margin-bottom: 2px; }}
.gt .two {{ display: flex; justify-content: space-between; }}
.gt .g {{ color: #1eff00; }}
'''

TICK = ('<svg class="tk" viewBox="0 0 12 12"><path d="M1.5 6.5 L4.5 9.5 L10.5 2.5" fill="none" stroke="#3fd23f" '
        'stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"/></svg>')


def STATES():
    base = dict(me.BASE, leads=[me.LEADS[k] for k in ('builds', 'gear', 'rotation', 'guild')])
    first = dict(base, companion=False, lastRaid=None, sim=None, syncedAgo=None, raidName=None, buildName=None,
                 spec=None, setupInput=SETUP)
    done = dict(base, setupInput=dict(SETUP, build={}, buildName='Beast Mastery raid build',
                                      specIcon='ability_hunter_beasttaming', hovered=True, companion=True))
    return {'first-open': {k: v for k, v in first.items() if v is not None}, 'steps-done': done}


def crest_html(size):
    return hub.crest('hunter', size)


def step_icon(row):
    if row.get('crest'):
        return crest_html(G['nowIcon'])
    if row.get('itemId'):
        return item_icon(row['itemId'], G['nowIcon'])
    name = row.get('icon') or ''
    if name.startswith('Interface'):
        name = name.split('\\')[-1].lower()
    return icon(name, G['nowIcon'], 'box-shadow:0 0 0 1px #000')


def tooltip_html():
    it = hub.ITEMS[WORN]
    pick = hub.ITEMS[TOP['pick']]
    hit = it['percent_stats'].get('hit')
    stats = ''.join(f'<div>+{v} {k.title()}</div>' for k, v in it['stats'].items() if k != 'hit' or not hit)
    equip = f'<div class="g">Equip: Improves your chance to hit by {hit}%.</div>' if hit else ''
    return (f'<div class="gt"><div class="t" style="color:{hub.QUALITY[it["quality"]]}">{esc(it["name"])}</div>'
            f'<div>Binds when picked up</div><div class="two"><span>{hub.SLOT_NAMES[it["slot"]]}</span>'
            f'<span>{it["armor_type"].title()}</span></div><div>{it["armor"]} Armor</div>{stats}{equip}'
            f'<div>Requires Level {it["required_level"]}</div>'
            f'<div>Next upgrade: <span style="color:{hub.QUALITY[pick["quality"]]}">{esc(pick["name"])}</span> '
            f'<span class="g">+{TOP["gain"]:.1f} DPS</span></div></div>')


def setup_html(model):
    out = me.section('Set up', model['note'])
    for row in model['rows']:
        mark = TICK if row['done'] else f'<span class="no">{row["number"]}</span>'
        under = (f'<div style="color:{ADDON["gold"]}">{esc(row["link"])}</div>' if row.get('link') else
                 f'<div class="m">{esc(row["sub"])}</div>' if row.get('sub') else '')
        tab = {'builds': 'Builds', 'gear': 'Gear'}.get(row.get('tab'), '')
        out += (f'<div class="su{" done" if row["done"] else ""}">{mark}<span class="ni">{step_icon(row)}</span>'
                f'<div class="tx"><div>{esc(row["line"])}</div>{under}</div><span class="lk">{tab}</span></div>')
    if model.get('example'):
        out += tooltip_html()
    return out + f'<div class="gap"></div>'


def board(caption, model, dots):
    figures = me.SEP.join(['Level 60 Hunter'] + model['figures'])
    left = me.score_html(model.get('score')) + me.raid_html(model.get('lastRaid'))
    right = setup_html(model['setup']) + me.sim_html(model.get('sim'))
    win = bw.window('hunter', 'Velisandre', esc(figures), f'<div class="me">{left}</div>',
                    f'<div class="me">{right}</div>', 'Esc to close', me.wow(model.get('footer') or ''),
                    strip=('Me', dots))
    return bw.ux.page(f'<style>{bw.CSS}{me.CSS}{CSS}</style><div title="{esc(caption)}">{win}</div>')


BOARDS = {
    'l8-first-open': ('first-open', 'The first login of a new tour version: Me with Set up, nothing done yet.',
                      ('Builds', 'Gear')),
    'l8-steps-done': ('steps-done', 'Set up again from /fs tour with every step done: ticked and muted.',
                      ('Gear', 'Guild')),
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
            key, caption, dots = BOARDS[name]
            page.set_content(board(caption, data[key], dots))
            page.evaluate('document.fonts.ready')
            page.wait_for_timeout(150)
            page.locator('#board').screenshot(path=str(OUT / f'{name}.png'))
            print(OUT / f'{name}.png')
        browser.close()


if __name__ == '__main__':
    main(sys.argv[1:] or list(BOARDS))
