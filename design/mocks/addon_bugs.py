"""Boards for the 2026-11-13 addon bug lane (2.0.43): the Builds tab's Follow, the bag arrow in
Baganator's bags, and long tooltip lines that wrap.

    python3 design/mocks/addon_bugs.py                    # every board
    python3 design/mocks/addon_bugs.py tooltip-long-wrap

Writes design/reviews/2026-11-13-addon-bugs/<board>.png at device scale 2. Nothing is typed here:
- The Builds boards are addon_builds.py's own (BuildsPane.LAYOUT, Locale.lua strings, the guides' builds).
- The bag arrow is SlotMarks.ARROW_PARTS, read from surfaces/SlotMarks.lua, in the success token, the
  same 12 px glyph the paper doll carries; the bag items and their icons are data/builds items/hunter.json.
- The tooltip's longest pick name and longest source are the shipped Data.lua's (worked by the addon's
  own ItemTooltip.nextLine and ItemTooltip.sourceLine over every published pick, see LONGEST below);
  every addon string is from Locale.lua.
"""
import os
import re
import sys
from pathlib import Path

sys.path.insert(0, str(Path(__file__).resolve().parent))
import addon_builds as builds  # noqa: E402
import addon_ux as ux  # noqa: E402
from addon_ux import ADDON, C, QUALITY, icon_uri, iname, l, page, tooltip  # noqa: E402

ROOT = Path(__file__).resolve().parents[2]
OUT = Path(os.environ['MOCK_OUT']) if os.environ.get('MOCK_OUT') else ROOT / 'design/reviews/2026-11-13-addon-bugs'
SLOTMARKS = (ROOT / 'addon/ForeverSixty/surfaces/SlotMarks.lua').read_text()
LOCALE = (ROOT / 'addon/ForeverSixty/Locale.lua').read_text()

# The shipped Data.lua's longest pick name and longest "where it drops" line (1.60.1.70938), as the
# addon prints them: ItemTooltip.nextLine and ItemTooltip.sourceLine over every published pick.
LONGEST = {
    'name': "Knight-Lieutenant's Crackling Leather Spaulders", 'name_quality': 3,
    'source': 'Scarlet Monastery · Scarlet Commander Mograine · 40% drop',
    'source_item': 'Aegis of the Scarlet Commander', 'source_quality': 3,
}


def L(key: str) -> str:
    return builds.L(key)


# --- The arrow, as SlotMarks draws it -----------------------------------------------------------------
def arrow_parts():
    block = re.search(r'SlotMarks\.ARROW_PARTS = \{(.*?)\n\}', SLOTMARKS, re.S).group(1)
    return [tuple(int(v) for v in m) for m in re.findall(r'\{ (\d+), (\d+), (\d+), (\d+) \}', block)]


def glyph(scale=1.0):
    """SlotMarks' 12 x 12 arrow: each part green over a 1 px black outline a sublevel behind."""
    rects = []
    for x, y, w, h in arrow_parts():
        rects.append(f'<rect x="{x - 1}" y="{y - 1}" width="{w + 2}" height="{h + 2}" fill="#000"/>')
    for x, y, w, h in arrow_parts():
        rects.append(f'<rect x="{x}" y="{y}" width="{w}" height="{h}" fill="{ADDON["success"]}"/>')
    size = 12 * scale
    return (f'<svg viewBox="-1 -1 14 14" width="{size + 2 * scale}" height="{size + 2 * scale}" '
            f'shape-rendering="crispEdges">{"".join(rects)}</svg>')


# --- 1. The Builds tab: Follow moves the build ------------------------------------------------------
def builds_board():
    head = builds.lua_format(L('dialogOwnBuilds'), 'Hunter')
    before = builds.own_pane('hunter', head, builds.note(60, 51, 51),
                             builds.class_blocks('hunter', 60, 'hunter-beast-mastery', 'raid', None,
                                                 states={'hunter-marksmanship': {1: 'cursor'}}))
    # Off its old plan after the switch: the note names the level alone, as BuildsPane.note does with no
    # points counted yet.
    after = builds.own_pane('hunter', head, builds.note(60),
                            builds.class_blocks('hunter', 60, 'hunter-marksmanship', 'raid', None))
    caption = ('Builds tab · level 60 Hunter. Left: following the Beast Mastery raid build, the cursor on '
               'Marksmanship\'s Raid build with Follow. Right: the same tab opened again after that Follow, '
               'Following moved to Marksmanship\'s Raid build (2.0.42 kept Beast Mastery: the redraw on the '
               'press hid the button, so the client dropped the click) · 420 wide')
    html = (f'<div style="display:flex;gap:28px;align-items:flex-start">{builds.dialog(before)}'
            f'{builds.dialog(after)}</div>')
    return builds.board(caption, html)


# --- 2. Baganator's bags with the arrow -------------------------------------------------------------
BAG_ITEMS = [
    # (item name in items/hunter.json, upgrade?)
    ('Outrider Leggings', True), ('Sentry\'s Gloves', False),
    ('Forest Leather Bracers', False), ('Scorpashi Breastplate', True), ('Ranger Bow', False),
    ('Short Ash Bow', False), ('Tunic of Westfall', False),
]


def find_item(name):
    for it in ux.ITEMS.values():
        if it['name'] == name and it.get('icon'):
            return it
    return None


def bag_slot(it, upgrade, size=37):
    if it is None:
        return (f'<div style="width:{size}px;height:{size}px;border-radius:3px;background:#14161c;'
                f'box-shadow:inset 0 0 0 1px #000,0 0 0 1px #2b2f39"></div>')
    edge = QUALITY.get(it['quality'], '#9d9d9d')
    mark = (f'<div style="position:absolute;top:1px;right:1px;line-height:0">{glyph()}</div>' if upgrade else '')
    return (f'<div style="position:relative;width:{size}px;height:{size}px;border-radius:3px;background:#14161c;'
            f'box-shadow:inset 0 0 0 1px #000,0 0 0 1px #2b2f39">'
            f'<img src="{icon_uri(it["icon"])}" style="position:absolute;inset:1px;width:{size - 2}px;'
            f'height:{size - 2}px;border-radius:2px;box-shadow:0 0 0 1px {edge}">{mark}</div>')


def bags_board():
    chosen = []
    for name, upgrade in BAG_ITEMS:
        it = find_item(name)
        if it is not None:
            chosen.append((it, upgrade))
    if sum(1 for _, up in chosen if up) == 0:
        raise SystemExit('no upgrade item found in items/hunter.json')
    filler = [it for it in ux.ITEMS.values() if it.get('icon') and it['quality'] in (1, 2)
              and 30 <= (it.get('required_level') or 0) <= 42 and it.get('slot')
              and it['name'] not in {n for n, _ in BAG_ITEMS}][:10]
    cells = [bag_slot(it, up) for it, up in chosen] + [bag_slot(it, False) for it in filler]
    cells += [bag_slot(None, False)] * (24 - len(cells))
    grid = ''.join(cells)
    bag = (f'<div style="width:{8 * 41 + 24}px;background:rgba(14,16,22,.96);border:1px solid #3a3f4b;'
           f'border-radius:4px;box-shadow:0 8px 28px rgba(0,0,0,.7);font-family:Site,sans-serif;color:#e8e8e8">'
           f'<div style="height:28px;display:flex;align-items:center;padding:0 10px;font-size:13px;'
           f'border-bottom:1px solid #2b2f39">Backpack</div>'
           f'<div style="display:grid;grid-template-columns:repeat(8,37px);gap:4px;padding:12px">{grid}</div>'
           f'</div>')
    widget = (f'<div class="note" style="max-width:300px"><b>Baganator, corner widgets:</b> '
              f'"{re.search(r"baganatorArrow = \"([^\"]+)\"", LOCALE).group(1)}" in the top right corner, '
              f'added there on first load. The same arrow as the paper doll: '
              f'<span style="display:inline-block;vertical-align:-2px">{glyph()}</span> SlotMarks.ARROW_PARTS, '
              f'12 px.</div>')
    caption = ('Bags · Baganator\'s backpack for a level 42 Hunter: the green arrow at the top right of the two '
               'items the chip lists as upgrades (2.0.42 drew none in Baganator\'s bags, only on the paper doll)')
    return page(f'<p class="cap">{caption}</p><div class="row">{bag}{widget}</div>')


# --- 3. A long tooltip that wraps ------------------------------------------------------------------
def tooltip_board():
    worn = [l('Brigade Girdle of the Bandit', color=QUALITY[2], cls='t'), l('Binds when equipped'),
            l('Waist', 'Mail'), l('121 Armor'), l('+4 Agility'), l('+9 Stamina'),
            l('Equip: +11 Attack Power.', color=C['green']), l('Requires Level 38')]
    name = f'<span style="color:{QUALITY[LONGEST["name_quality"]]}">{LONGEST["name"]}</span>'
    nxt = f'{L("tooltipNextLabel")} {name} <span style="color:{C["green"]}">+4.2 DPS</span>'
    wrap = 'white-space:normal'
    ours = [f'<div class="l" style="color:#fff;{wrap}"><span>{nxt}</span></div>',
            f'<div class="l" style="color:{C["grey"]};{wrap}"><span>{LONGEST["source"]}</span></div>']
    before = [f'<div class="l" style="color:#fff"><span>{L("tooltipNextLabel")} <span style="color:'
              f'{QUALITY[3]}">Knight-Lieutenant\'s Crac…</span> <span style="color:{C["green"]}">+4.2 DPS</span>'
              f'</span></div>',
              f'<div class="l" style="color:{C["grey"]}"><span>Scarlet M… · Scarlet Commander Mograine</span></div>']
    width = 262
    left = tooltip(worn + before, tag='2.0.42: cut at the client\'s widest row', width=width)
    right = tooltip(worn + ours, tag='2.0.43: whole, wrapped like the client\'s lines', width=width)
    caption = ('Tooltip · the shipped data\'s longest pick name and longest source on a worn belt\'s hover, at the '
               'belt tooltip\'s own width')
    return page(f'<p class="cap">{caption}</p><div class="row" style="margin-top:22px">{left}{right}</div>')


BOARDS = {
    'builds-follow-moves': builds_board,
    'bags-baganator-arrows': bags_board,
    'tooltip-long-wrap': tooltip_board,
}


def main(names):
    from playwright.sync_api import sync_playwright
    OUT.mkdir(parents=True, exist_ok=True)
    with sync_playwright() as pw:
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
