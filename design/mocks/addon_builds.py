"""Acceptance boards for the Builds pane and the panel's Talents row
(design/specs/2026-11-13-addon-builds-pane.md section 1.5).

    python3 design/mocks/addon_builds.py            # every board
    python3 design/mocks/addon_builds.py builds-hunter-42

Writes design/reviews/2026-11-13-addon-builds/<board>.png at device scale 2, each frame at the addon's
real size: the Builds dialog 420 wide (Theme.SIZES.dialogWidth, 16 px padding, 12 px text), the panel
340 wide (Panel.WIDTH). The layout replays surfaces/BuildsPane.lua's (BuildsPane.LAYOUT) and
surfaces/Panel.lua's drawTalents line for line.

Sources, so nothing is invented:
- Every string: addon/ForeverSixty/Locale.lua, read here (the L.* keys and the specPlaystyle table).
- Builds, their orders and their whys: data/builds/1.60.1.70790/addon-builds.json (sim/cmd/addon-builds,
  from the guides), the split from those orders over data/builds/1.60.1.70790/talents/<class>.json.
- Spec icons: web/src/data/specs.json "icon" (the planner's and Data.lua's specIcons), drawn from
  data/builds/1.60.1.70790/icons; talent icons from talents/<class>.json.
- The crest: addon/ForeverSixty/media/crests/<class>.tga under ring.tga tinted in the class colour
  (Theme.buildCrest and Theme.paintCrest), never a square class icon.
- The level 42 hunter's figures (27 of 33 points, Intimidation next, one point waiting) are the owner's
  capture of 2026-11-12; the next talent on every other board is the first point of the followed order
  the character has not taken, from the same orders.
"""
import collections
import json
import os
import re
import sys

import addon_ux as ux
from PIL import ImageFont

from addon_ux import ADDON, BUILD, ROOT, icon_uri, page, tga_uri

OUT = ROOT / 'design/reviews/2026-11-13-addon-builds'
LOCALE = (ROOT / 'addon/ForeverSixty/Locale.lua').read_text()
BUILDS = json.loads((BUILD / 'addon-builds.json').read_text())
SPECS = {s['spec']: s for s in json.loads((ROOT / 'web/src/data/specs.json').read_text())}
CLASS_HEX = {'hunter': '#aad372', 'paladin': '#f48cba', 'druid': '#ff7c0a', 'mage': '#3fc7eb'}
TREE_ORDER = {'druid': ['balance', 'feral', 'feral-bear', 'restoration']}

# BuildsPane.LAYOUT and the dialog's sizes (Theme.SIZES), in px.
W, PAD, TEXT, LINE = 420, 16, 12, 15
G = {'heading': 24, 'crest': 20, 'specLine': 18, 'specIcon': 16, 'indent': 22, 'rowGap': 6, 'blockGap': 8,
     'lineGap': 1, 'afterHeading': 4, 'playstyleGap': 6, 'groundLead': 6, 'nextIcon': 14, 'nextIconGap': 4,
     'siteCrest': 20, 'moreGap': 6}


# Keys the Builds window's header retired (spec 2026-11-13-addon-hub 3.2), kept for this older board set.
RETIRED = {'dialogOwnBuilds': '%s builds', 'dialogLevelPoints': 'Level %d'}


def L(key: str) -> str:
    """A Locale.lua string by key, SITE substituted."""
    m = re.search(rf'(?:^L\.|^\t){key} = (.+)$', LOCALE, re.M)
    if not m and key in RETIRED:
        return RETIRED[key]
    if not m:
        raise KeyError(key)
    expr = m.group(1).rstrip(',')
    parts = re.findall(r'"((?:[^"\\]|\\.)*)"|(SITE)', expr)
    return ''.join('foreversixty.gg' if site else lua_unescape(lit) for lit, site in parts)


def lua_unescape(lit: str) -> str:
    """A Lua string literal's decimal byte escapes (\\194\\183 is the middle dot) decoded as the client does."""
    raw = re.sub(rb'\\(\d{1,3})', lambda m: bytes([int(m.group(1))]), lit.encode('utf-8'))
    return raw.decode('utf-8')


def playstyle(spec: str) -> str:
    return re.search(rf'\["{re.escape(spec)}"\] = "([^"]*)"', LOCALE).group(1)


def lua_format(fmt: str, *args) -> str:
    return re.sub(r'%[ds]', lambda _m, it=iter(args): str(next(it)), fmt)


def trees(cls):
    return json.loads((BUILD / f'talents/{cls}.json').read_text())['trees']


def split(cls, order):
    counts = collections.Counter(p[0] for p in order)
    names = [t['name'] for t in trees(cls)]
    tabs = sorted(counts, key=lambda t: (-counts[t], t))
    return L('dialogSplitSeparator').join(lua_format(L('dialogSplitPart'), names[t - 1], counts[t]) for t in tabs)


def esc(text):
    return text.replace('&', '&amp;').replace('<', '&lt;').replace('>', '&gt;')


CLIENT_GOLD = '#ffd100'  # Theme.HEX.clientGold (forever.tooltipYellow), the client's gold
FONTS = ROOT / 'addon/ForeverSixty/media/fonts'
MEASURE = {400: ImageFont.truetype(str(FONTS / 'OpenSans-Regular.ttf'), TEXT),
           600: ImageFont.truetype(str(FONTS / 'OpenSans-SemiBold.ttf'), TEXT)}


def measure(text, weight=400):
    return MEASURE[weight].getlength(text)


CSS = f'''
.dlg {{ position: relative; width: {W}px; background: {ADDON['raised']}; border: 1px solid {ADDON['line']};
  box-shadow: 0 8px 28px rgba(0,0,0,.7); font-family: Site, sans-serif; font-size: {TEXT}px; line-height: {LINE}px;
  color: {ADDON['body']}; padding: 0 {PAD}px 12px; }}
.dlg .title {{ height: 28px; display: flex; align-items: center; font-family: Title, serif; font-size: 15px;
  color: {ADDON['cream']}; }}
.dlg .x {{ position: absolute; top: 4px; right: 6px; width: 20px; height: 20px; border: 1px solid {ADDON['line']};
  display: flex; align-items: center; justify-content: center; font-size: 14px; }}
.tabs {{ display: flex; gap: 8px; height: 24px; margin-bottom: 4px; }}
.tab {{ min-width: 96px; padding: 0 8px; white-space: nowrap; display: flex; align-items: center;
  justify-content: center; color: {ADDON['muted']}; border-bottom: 2px solid transparent; }}
.tab.on {{ color: {ADDON['gold']}; border-color: {ADDON['gold']}; }}
.m {{ color: {ADDON['muted']}; }}
.one {{ white-space: nowrap; overflow: hidden; text-overflow: ellipsis; }}
.head {{ height: {G['heading']}px; display: flex; align-items: center; gap: 8px; margin-bottom: {G['afterHeading']}px; }}
.head .h {{ flex: 1; min-width: 0; }}
.crest {{ position: relative; width: {G['crest']}px; height: {G['crest']}px; flex: none; display: block; }}
.crest img {{ position: absolute; inset: 0; width: 100%; height: 100%; }}
.spec {{ height: {G['specLine']}px; display: flex; align-items: center; }}
.spec img {{ width: {G['specIcon']}px; height: {G['specIcon']}px; box-shadow: 0 0 0 1px #000;
  margin-right: {G['indent'] - G['specIcon']}px; }}
.spec .n {{ font-weight: 600; color: {ADDON['strong']}; }}
.spec .p {{ margin-left: {G['playstyleGap']}px; color: {ADDON['muted']}; white-space: nowrap; }}
.brow {{ position: relative; margin: 0 0 {G['rowGap']}px {G['indent']}px; }}
.brow::before {{ content: ''; position: absolute; left: -{G['groundLead']}px; top: 0; bottom: 0; right: 0; z-index: 0; }}
.brow::after {{ content: ''; position: absolute; left: -{G['groundLead']}px; top: 0; bottom: 0; width: 2px; }}
.brow > div {{ position: relative; z-index: 1; height: {LINE}px; margin-bottom: {G['lineGap']}px; }}
.brow .l1 {{ display: flex; justify-content: space-between; gap: 8px; }}
.brow .fol {{ color: {ADDON['gold']}; }}
.brow.followed::before {{ background: rgba(229, 185, 85, .08); }}
.brow.followed::after {{ background: {ADDON['gold']}; }}
.brow.hover::before, .brow.cursor::before {{ background: {ADDON['hover']}; }}
.brow.cursor::after {{ background: {ADDON['body']}; }}
.brow.pressed::before {{ background: {ADDON['inset']}; }}
.brow .nx {{ display: flex; align-items: center; gap: {G['nextIconGap']}px; }}
.brow .nx img {{ width: {G['nextIcon']}px; height: {G['nextIcon']}px; box-shadow: 0 0 0 1px #000; }}
.sb {{ position: absolute; right: 0; top: 0; height: 20px; padding: 0 7px; display: inline-flex; align-items: center;
  font-weight: 600; color: {ADDON['strong']}; background: {ADDON['blue']}; border: 1px solid {ADDON['blue_edge']}; }}
.sb.hov {{ background: {ADDON['blue_edge']}; }}
.sb.held {{ background: {ADDON['blue_edge']}; padding-top: 2px; }}
.site {{ margin: 10px 0 6px; color: {ADDON['body']}; }}
.box {{ height: 42px; border: 1px solid {ADDON['line']}; background: {ADDON['inset']}; padding: 5px 8px;
  color: {ADDON['muted']}; }}
.srow {{ display: flex; gap: 8px; margin-bottom: 4px; }}
.srow .tx {{ flex: 1; min-width: 0; }}
.more {{ margin: 0 0 {G['moreGap']}px {G['siteCrest'] + 8}px; color: {ADDON['muted']}; }}
.foot {{ display: flex; justify-content: flex-end; margin-top: 8px; color: {ADDON['muted']}; opacity: .75;
  height: {LINE}px; }}
.tip {{ background: rgba(8, 10, 22, .95); border: 1px solid #9095a3; border-radius: 4px; padding: 7px 10px 8px;
  width: 300px; font-family: Site, sans-serif; font-size: 12px; line-height: 15px; color: {ADDON['body']}; }}
.tip .t {{ color: #fff; font-size: 14px; line-height: 18px; margin-bottom: 1px; }}
.note2 {{ font-family: -apple-system, 'Helvetica Neue', Arial, sans-serif; font-size: 12px; color: #9a9484;
  margin: 0 0 10px; max-width: 900px; }}
'''


def crest(cls, size=G['crest']):
    ring_uri = tga_uri(ROOT / 'addon/ForeverSixty/media/crests/ring.tga')
    face = ROOT / f'addon/ForeverSixty/media/crests/{cls}.tga'
    return (f'<span class="crest" style="width:{size}px;height:{size}px"><img src="{tga_uri(face)}">'
            f'<span style="position:absolute;inset:0;background:{CLASS_HEX[cls]};'
            f'-webkit-mask:url({ring_uri}) center/100% no-repeat;mask:url({ring_uri}) center/100% no-repeat"></span></span>')


def build_rows(spec, level, followed=None, nxt=None):
    """BuildsPane's specRows: the leveling build, the raid build, or one row for both."""
    cls = spec.split('-')[0]
    rows = []
    for kind in ('leveling', 'raid'):
        b = BUILDS[spec][kind]
        if rows and rows[0]['order'] == b['order']:
            rows[0].update(kind='both', label=L('dialogKindBoth'))
            continue
        rows.append({'spec': spec, 'kind': kind, 'order': b['order'], 'why': b.get('why'),
                     'label': L('dialogKindRaid') if kind == 'raid' else L('dialogKindLeveling'),
                     'detail': split(cls, b['order'])})
    for r in rows:
        r['following'] = followed is not None and (followed == r['kind'] or r['kind'] == 'both')
        if r['following']:
            r['next'] = nxt
    return rows


def row_html(r, state=''):
    """BuildsPane's drawRow: label and Following or Follow, the split, the next point, the why."""
    button = {'hov': ' hov', 'held': ' held'}.get(r.get('button'), '')
    right = (f'<span class="fol">{L("dialogFollowingMark")}</span>' if r['following']
             else f'<span class="sb{button}">{L("dialogFollow")}</span>')
    lines = [f'<div class="l1"><span class="one">{esc(r["label"])}</span>{right}</div>']
    if r.get('failed'):
        lines.append(f'<div class="one" style="color:{ADDON["warning"]}">{L("dialogOwnFailed")}</div>')
    else:
        room = 'padding-right:52px' if not r['following'] else ''
        lines.append(f'<div class="one m" style="{room}">{esc(r["detail"])}</div>')
    if r.get('next'):
        html, icon = r['next']
        lines.append(f'<div class="nx">{ux.img(icon, G["nextIcon"])}<span class="one">{html}</span></div>')
    if r.get('why'):
        lines.append(f'<div class="one m">{esc(r["why"])}</div>')
    if r['following'] and not state:
        state = 'followed'
    return f'<div class="brow {state}">{"".join(lines)}</div>'


def spec_block(spec, rows, states=None):
    """BuildsPane's drawSpecLine: the icon, the name strong, the playstyle muted when it fits whole."""
    s = SPECS[spec]
    icon = f'<img src="{icon_uri(s["icon"])}">'
    tagline = playstyle(spec)
    room = W - 2 * PAD - G['indent'] - measure(s['name'], 600) - G['playstyleGap']
    tag = f'<span class="p">{esc(tagline)}</span>' if measure(tagline) <= room else ''
    line = f'<div class="spec">{icon}<span class="n">{s["name"]}</span>{tag}</div>'
    return line + ''.join(row_html(r, (states or {}).get(i, '')) for i, r in enumerate(rows))


def own_pane(cls, heading, note_html, blocks):
    note_span = f'<span class="m" style="white-space:nowrap">{note_html}</span>' if note_html else ''
    out = f'<div class="head">{crest(cls)}<span class="h one">{heading}</span>{note_span}</div>'
    return out + f'<div style="height:{G["blockGap"]}px"></div>'.join(blocks)


def site_rows(rows, more):
    """The site rows: the 20 px crest, the name and its points, the split, the character and the day."""
    out = ''
    for name, detail, target, cls in rows:
        out += (f'<div class="srow">{crest(cls, G["siteCrest"])}<div class="tx">'
                f'<div class="one">{esc(name)}</div><div class="one m">{esc(detail)}</div>'
                f'<div class="one m">{esc(target)}</div></div></div>')
    if more:
        out += f'<div class="more">{lua_format(L("dialogInboxMore"), more)}</div>'
    return out


def dialog(own, site='', hint=None):
    """The Builds tab: with the class's own builds neither the companion line nor Tab to switch shows."""
    box = f'<div class="box">{esc(L("dialogPaste"))}</div>'
    return (f'<div class="dlg"><div class="title">{L("addonName")}</div><div class="x">×</div>'
            f'<div class="tabs"><div class="tab on">{L("dialogTabPaste")}</div><div class="tab">{L("dialogCopy")}</div>'
            f'</div>{own}<div class="site">{L("dialogSiteHeading")}</div>{site}{box}'
            + (f'<div class="foot"><span>{hint}</span></div>' if hint else '') + '</div>')


def board(caption, html):
    return page(f'<style>{CSS}</style><p class="note2">{caption}</p>{html}')


def note(level, matched=None, allowed=None, off=0):
    parts = [lua_format(L('dialogLevelPoints'), level)]
    if allowed:
        parts.append(lua_format(L('panelOnPlan'), matched, allowed))
    if off:
        parts.append(f'<span style="color:{ADDON["warning"]}">{lua_format(L("panelOffPlanCount"), off)}</span>')
    return L('dialogSeparator').join(parts)


def talent_icon(cls, name):
    for t in trees(cls):
        for n in t['talents']:
            if n['name'] == name:
                return n['icon']
    raise KeyError(name)


def next_line(cls, name, unspent=0, at=None):
    """Follow.nextPointLine: "Intimidation · Spend 1 now" (the spend in the client gold) or "Level 43: ..."."""
    if unspent > 0:
        html = f'{name}{L("exportTreeSeparator")}<span style="color:{CLIENT_GOLD}">{lua_format(L("panelSpendNow"), unspent)}</span>'
    else:
        html = lua_format(L('panelTalentAt'), at, name)
    return html, talent_icon(cls, name)


def class_blocks(cls, level, followed_spec=None, kind=None, nxt=None, specs=None, states=None):
    order = specs or [s for s in SPECS if s.startswith(cls + '-')]
    return [spec_block(spec, build_rows(spec, level, kind if spec == followed_spec else None, nxt),
                       (states or {}).get(spec)) for spec in order]


def board_hunter_42():
    blocks = class_blocks('hunter', 42, 'hunter-beast-mastery', 'leveling', next_line('hunter', 'Intimidation', 1),
                          states={'hunter-marksmanship': {1: 'cursor'}})
    own = own_pane('hunter', lua_format(L('dialogOwnBuilds'), 'Hunter'), note(42, 27, 33, 5), blocks)
    return board("Builds tab · level 42 Hunter following Beast Mastery's leveling build, one point to spend and five "
                 "off plan (the owner's case, 2026-11-12); the cursor on Marksmanship's raid build · 420 wide",
                 dialog(own, hint=L('dialogEnterFollow')))


def board_hunter_60_raid():
    blocks = class_blocks('hunter', 60, 'hunter-marksmanship', 'raid', None)
    own = own_pane('hunter', lua_format(L('dialogOwnBuilds'), 'Hunter'), note(60, 51, 51), blocks)
    return board('Builds tab · level 60 Hunter following the Marksmanship raid build, every point placed', dialog(own))


def board_paladin_13():
    question = 'Are you Holy or Retribution?'
    splitp = L('exportTreeSeparator').join([lua_format(L('pathSplit'), 2, 'Holy'),
                                            lua_format(L('pathSplit'), 2, 'Retribution')])
    heading = f'{question}{L("exportTreeSeparator")}<span class="m">{splitp}</span>'
    own = own_pane('paladin', heading, None, class_blocks('paladin', 13))
    return board('Builds tab · level 13 Paladin, points split 2 and 2, nothing followed: the question heads the list',
                 dialog(own))


def board_druid_24():
    blocks = class_blocks('druid', 24, 'druid-feral', 'leveling', next_line('druid', 'Feral Instinct', 0, 25),
                          specs=['druid-balance', 'druid-feral', 'druid-feral-bear', 'druid-restoration'])
    own = own_pane('druid', lua_format(L('dialogOwnBuilds'), 'Druid'), note(24, 15, 15), blocks)
    site = site_rows([(f'Moonkin for Molten Core{L("dialogSeparator")}51 points', 'Druid · Balance 31 / Restoration 20',
                       f'{lua_format(L("dialogForCharacter"), "Leafcall")}{L("dialogSeparator")}12 Nov', 'druid')], 3)
    return board('Builds tab · level 24 Druid following Feral, four builds waiting from the site: six own rows leave '
                 'room for one site row and "and 3 more"', dialog(own, site=site))


def board_row_states():
    base = build_rows('hunter-marksmanship', 42)[0]
    cases = [('Default', '', {}), ('Hover (the row and its Follow)', 'hover', {'button': 'hov'}),
             ('Cursor (Up, Down)', 'cursor', {}), ('Held (the row and its Follow)', 'pressed', {'button': 'held'}),
             ('Following', '', {'following': True, 'next': next_line('hunter', 'Lethal Attacks', 1)}),
             ('Did not load', '', {'failed': True})]
    cells = [f'<div style="width:{W - 2 * PAD}px"><p class="note2" style="margin:0 0 4px">{caption}</p>'
             f'{row_html(dict(base, **extra), state)}</div>' for caption, state, extra in cases]
    tip = (f'<div class="tip"><div class="t">{esc(base["label"])}</div><div class="m">{esc(base["detail"])}</div>'
           f'<div>{esc(base["why"])}</div><div class="m">{L("dialogClickFollow")}</div></div>')
    grid = ('<div style="display:grid;grid-template-columns:repeat(3,auto);gap:14px 30px;padding-left:6px">'
            + ''.join(cells) + '</div>')
    return board("A build row's states (Marksmanship's leveling build), and its hover tooltip",
                 f'<div class="dlg" style="width:auto;padding:14px 16px">{grid}<p class="note2" '
                 f'style="margin:16px 0 4px">Hover tooltip</p>{tip}</div>')


def panel_talents(after=True, name='Beast Mastery build'):
    note_r = (f'{lua_format(L("panelOnPlan"), 27, 33)}{L("exportTreeSeparator")}'
              f'<span style="color:{ADDON["warning"]}">{lua_format(L("panelOffPlanCount"), 5)}</span>')
    icon = ux.img(talent_icon('hunter', 'Intimidation'), 28)
    btns = (f'<span class="sbtn" style="height:20px;padding:0 7px">{L("panelChange")}</span>'
            f'<span class="sbtn" style="height:20px;padding:0 7px">{L("panelTalentsButton")}</span>')
    title, _ = next_line('hunter', 'Intimidation', 1)
    if after:
        row = (f'<div style="display:flex;gap:8px;height:38px">{icon}<div style="flex:1;min-width:0">'
               f'<div class="one" style="height:14px;line-height:14px;margin-bottom:4px">{title}</div>'
               f'<div style="display:flex;align-items:center;gap:6px;height:20px">'
               f'<span class="one" style="flex:1;color:{ADDON["muted"]};font-size:{ux.SMALL["size"] + 1}px">{name}</span>{btns}'
               f'</div></div></div>')
    else:
        row = (f'<div style="display:flex;gap:8px;height:34px">{icon}<div style="flex:1;min-width:0">'
               f'<div>Intimidation</div><div class="one" style="color:{ADDON["muted"]};max-width:145px">Beast Mastery '
               f'build{L("exportTreeSeparator")}<span style="color:{CLIENT_GOLD}">Spen…</span></div></div>'
               f'<span class="sbtn" style="height:20px;width:72px;margin-top:4px"></span>'
               f'<span class="sbtn" style="height:20px;padding:0 7px;margin-top:4px">{L("panelTalentsButton")}</span></div>')
    sect = f'<div class="sect">{L("panelTalents")}<span class="r">{note_r}</span></div>{row}'
    return ux.panel('Frickin Laserbeams', 'Level 42 Hunter', [sect])


def board_panel_talents():
    return board("The panel's Talents row at 340 wide, level 42 Beast Mastery, one point waiting: the owner's capture "
                 "replayed (left), the fix following the leveling build, the level's own kind, named for its spec "
                 "(middle), and following the raid build at 42, named in full (right)",
                 f'<div class="row"><div>{panel_talents(False)}</div><div>{panel_talents(True)}</div>'
                 f'<div>{panel_talents(True, "Beast Mastery raid build")}</div></div>')


BOARDS = {
    'builds-hunter-42': board_hunter_42,
    'builds-hunter-60-raid': board_hunter_60_raid,
    'builds-paladin-13-ask': board_paladin_13,
    'builds-druid-24-site': board_druid_24,
    'builds-row-states': board_row_states,
    'panel-talents-42': board_panel_talents,
}


def main(names):
    from playwright.sync_api import sync_playwright
    OUT.mkdir(parents=True, exist_ok=True)
    with sync_playwright() as pw:
        # PLAYWRIGHT_CHROMIUM: a headless shell when the pip package's own build is not downloaded.
        exe = os.environ.get('PLAYWRIGHT_CHROMIUM')
        browser = pw.chromium.launch(executable_path=exe) if exe else pw.chromium.launch()
        ctx = browser.new_context(device_scale_factor=2, viewport={'width': 1400, 'height': 1000})
        pg = ctx.new_page()
        for name in names:
            pg.set_content(BOARDS[name]())
            pg.evaluate('document.fonts.ready')
            pg.wait_for_timeout(150)
            pg.locator('#board').screenshot(path=str(OUT / f'{name}.png'))
            print(OUT / f'{name}.png')
        browser.close()


if __name__ == '__main__':
    main(sys.argv[1:] or list(BOARDS))
