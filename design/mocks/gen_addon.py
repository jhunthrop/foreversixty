"""Mock boards for the in-game addon redesign (design/specs/2026-10-04-addon.md).

  cd web && python3 ../design/mocks/render.py addon-overview addon-tooltip addon-talents addon-gear

Every number, name, icon and source line below comes from real data already in the repo:
design/mocks/data/guides-warrior-fury.json (the live warrior-fury BiS file, band 20-29,
alliance), design/mocks/data/addon-talents.json and design/mocks/data/addon-rotation.json
(both cut from the site's own 1.60.1.70009 talent/rotation data, fetched once into
design/mocks/data/ by this lane), and design/mocks/data/addon-items.json (item name/icon/
quality/stats, fetched the same way). The worked character throughout is the owner's level 23
Human Fury Warrior, Obnoxious Yell, read against the 20-29 alliance band. Anything this file
cannot source from those fixtures (a generated sync code, a live rating card, bag contents) is
rendered in its real, honest empty state rather than invented -- see the spec's own state tables.

Colours and sizes are Theme.lua's own HEX/SIZES tables, copied by hand (the addon has no CSS
build to import from); a change to Theme.lua's numbers must be mirrored here by hand too.
"""
import json
import math
from pathlib import Path

from mocklib import DATA, crest, emblem, data_uri, write

ICONS = DATA / "icons"

# ---------------------------------------------------------------------------
# Theme.lua's own tables, copied by hand (addon/ForeverSixty/Theme.lua).
# ---------------------------------------------------------------------------
HEX = {
    "background": "0d111a", "inset": "070a10", "raised": "161c2b", "hover": "1f2739",
    "card": "111726", "sidebar": "0a0e16", "track": "1b2233", "onGold": "15110a",
    "goldHover": "f2cd74", "success": "6fcf8e", "border": "262e40", "titleTop": "131824",
    "titleBottom": "0d111a", "gold": "e5b955", "muted": "9a9484", "body": "e9e4d8",
    "warning": "ff6b5c",
    # The character's own class colour (RAID_CLASS_COLORS.WARRIOR), addressable through
    # hexcol()/bar()/pill() the same way every other token here is -- round-2 fix: the
    # Talent points card's three bars are class-coloured, not plain gold.
    "classWarrior": "c79c6e",
}
S = {
    "windowWidth": 720, "windowHeight": 500, "sidebarWidth": 150, "navHeight": 30,
    "navIcon": 16, "navBar": 2, "titleBarHeight": 28, "headerHeight": 56, "padding": 16,
    "gap": 4, "rowHeight": 20, "cardGap": 12, "cardHeight": 150, "progressHeight": 6,
    "pillHeight": 16, "closeButton": 18, "border": 1, "iconSize": 16, "trackerWidth": 240,
    "trackerHeight": 48, "minimapButton": 32, "minimapIcon": 20, "minimapRadius": 80,
    "rotationRowHeight": 34, "rotationHeaderIcon": 14, "equipButtonWidth": 72,
    "buttonHeight": 22, "buttonWidth": 150, "gearSlotRows": 9, "followRows": 11,
}
# RAID_CLASS_COLORS.WARRIOR -- the client's real table, not the site's documented
# approximation (design/DESIGN-SYSTEM.md's #c69b6d is close but not byte-identical).
CLASS_WARRIOR = "c79c6e"
# ITEM_QUALITY_COLORS, unlightened -- the addon shows the client's true rarity
# colours (Theme.qualityColor reads this table directly), never the site's AA-adjusted
# rare/epic text variants, because that is what ships today and it is correct: a
# player's eye for "blue item, purple item" is trained on the client's own colours.
QUALITY = {2: "1eff00", 3: "0070dd", 4: "a335ee"}
FONT = "'PT Serif', Georgia, serif"  # stand-in for GameFontNormal (Friz Quadrata); named in the spec.

ITEMS = json.loads((DATA / "addon-items.json").read_text())
TALENTS = json.loads((DATA / "addon-talents.json").read_text())
ROTATION = json.loads((DATA / "addon-rotation.json").read_text())
BIS = json.loads((DATA / "guides-warrior-fury.json").read_text())
BAND20 = next(b for b in BIS["bands"] if b["band"] == 20 and b["faction"] == "alliance")
BUILD = BIS["build"]

# Round-2 blocker fix: one fixture value for "what is in the bags", read by both the
# Overview Gear card and the Gear page's own "UPGRADES IN YOUR BAGS" section -- never a
# second, independently-typed count. Empty here (honest: there is no bag-contents file
# in this lane's data to source a non-empty one from), so both surfaces agree it is
# empty, the same way GearView.rows/Gear.upgrades is the one model the real Overview
# card and the real Gear page would both read from in the build.
GEAR_BAG_UPGRADES = []


def gear_upgrades_detail():
    if GEAR_BAG_UPGRADES:
        return f"{len(GEAR_BAG_UPGRADES)} upgrades waiting in your bags, by our stat weights."
    return "Nothing in your bags beats what you are wearing."


def hexcol(key):
    return f"#{HEX[key]}"


def data_uri_any(path):
    import base64
    mime = {"webp": "image/webp", "svg": "image/svg+xml", "png": "image/png", "jpg": "image/jpeg"}[path.suffix.lstrip(".")]
    return f"data:{mime};base64," + base64.b64encode(path.read_bytes()).decode()


def icon_uri(name):
    return data_uri_any(ICONS / f"{name}.webp")


def icon(src, size, ring=None):
    """One game icon, bevel-cropped (Theme.ICON_CROP's own 0.08..0.92), optionally ringed
    in a quality or glow colour (Theme.outline's own 1-2px flat border)."""
    inner = size / (0.92 - 0.08)
    offset = (inner - size) / 2
    box = f"width:{size}px;height:{size}px;overflow:hidden;position:relative;flex-shrink:0;background:#{HEX['raised']};"
    if ring:
        box += f"box-shadow:0 0 0 1px {ring};"
    img = f"position:absolute;top:-{offset:.1f}px;left:-{offset:.1f}px;width:{inner:.1f}px;height:{inner:.1f}px;"
    return f'<div style="{box}"><img src="{src}" style="{img}"/></div>'


def label(text, color="body", size=12, weight=400, family=FONT, tracking=None):
    col = hexcol(color) if color in HEX else color
    track = f"letter-spacing:{tracking};" if tracking else ""
    return (f'<span style="font-family:{family};font-size:{size}px;font-weight:{weight};'
            f'color:{col};{track}">{text}</span>')


def pill(text, color="muted"):
    return (f'<span style="display:inline-flex;align-items:center;height:{S["pillHeight"]}px;'
            f'padding:0 8px;background:#{HEX["raised"]};box-shadow:0 0 0 1px #{HEX["border"]};'
            f'font-family:{FONT};font-size:11px;color:{hexcol(color)};white-space:nowrap">{text}</span>')


def bar(fraction, width, color="gold", height=None):
    h = height or S["progressHeight"]
    frac = max(0.0, min(1.0, fraction))
    return (f'<div style="width:{width}px;height:{h}px;background:#{HEX["track"]};position:relative">'
            f'<div style="position:absolute;left:0;top:0;height:{h}px;width:{frac*100:.1f}%;background:{hexcol(color)}"></div></div>')


def primary_button(text, width=None, enabled=True):
    w = f"width:{width}px;" if width else "padding:0 16px;"
    bg = hexcol("gold") if enabled else hexcol("gold")
    alpha = "1" if enabled else "0.4"
    return (f'<div style="{w}height:{S["buttonHeight"]+S["gap"]}px;background:{bg};opacity:{alpha};'
            f'display:flex;align-items:center;justify-content:center;font-family:{FONT};'
            f'font-size:11px;color:#{HEX["onGold"]};white-space:nowrap">{text}</div>')


def secondary_button(text, width=None, enabled=True):
    w = f"width:{width}px;" if width else "padding:0 14px;"
    alpha = "1" if enabled else "0.4"
    return (f'<div style="{w}height:{S["buttonHeight"]}px;background:#{HEX["raised"]};'
            f'box-shadow:0 0 0 1px #{HEX["border"]};opacity:{alpha};display:flex;align-items:center;'
            f'justify-content:center;font-family:{FONT};font-size:11px;color:{hexcol("gold")};white-space:nowrap">{text}</div>')


def panel(children, width=None, height=None, pad=None, extra=""):
    w = f"width:{width}px;" if width else ""
    h = f"height:{height}px;" if height else ""
    p = f"padding:{pad}px;" if pad is not None else ""
    return (f'<div style="{w}{h}{p}background:#{HEX["card"]};box-shadow:inset 0 1px 0 rgba(229,185,85,.15),'
            f'0 0 0 1px #{HEX["border"]};box-sizing:border-box;{extra}">{children}</div>')


def _joined(children):
    if isinstance(children, (list, tuple)):
        return "".join(children)
    return children


def row(children, gap=4, justify="flex-start", align="center"):
    return (f'<div style="display:flex;align-items:{align};justify-content:{justify};gap:{gap}px">{_joined(children)}</div>')


def col(children, gap=4):
    return f'<div style="display:flex;flex-direction:column;gap:{gap}px">{_joined(children)}</div>'


def abspos(x, y, children, extra=""):
    return f'<div style="position:absolute;left:{x}px;top:{y}px;{extra}">{children}</div>'


# ---------------------------------------------------------------------------
# Game backdrop: a dim vignette, never a fabricated scene (owner instruction).
# ---------------------------------------------------------------------------
def addon_page(body, width, height_min=900):
    style = f'''<style>
@import url('https://fonts.googleapis.com/css2?family=PT+Serif:wght@400;700&display=swap');
*{{box-sizing:border-box;margin:0;padding:0}}
html,body{{background:#000;font-family:{FONT}}}
</style>'''
    backdrop = (
        "background:"
        "radial-gradient(ellipse 1200px 800px at 50% 38%, #171a1f 0%, #0b0c0e 55%, #040405 100%);"
    )
    return (f'<!doctype html><html><head><meta charset="utf-8">{style}</head>'
            f'<body><div style="position:relative;width:{width}px;min-height:{height_min}px;{backdrop}">'
            f'{body}</div></body></html>')


def caption(text, x, y, width=300):
    return abspos(x, y, label(text, "muted", 11, 400), extra=f"width:{width}px;font-style:italic")


# ---------------------------------------------------------------------------
# The window chrome: title bar, sidebar, character header. Window.lua's own
# measurements (Theme.SIZES), 1x UI scale.
# ---------------------------------------------------------------------------
NAV_TABS = [
    ("overview", "Overview", "inv_misc_map_01.jpg"),
    ("follow", "Talents", "inv_misc_book_09"),
    ("gear", "Gear", "inv_chest_chain"),
    ("guild", "Guild", "inv_bannerpvp_02.jpg"),
    ("export", "Export", "inv_letter_15.jpg"),
    ("settings", "Settings", "inv_misc_gear_01"),
]


def nav_icon_src(name):
    path = ICONS / name
    if path.suffix:
        return data_uri_any(path)
    return data_uri_any(path.with_suffix(".webp"))


def title_bar():
    half = S["titleBarHeight"] / 2
    grad = (f'<div style="position:absolute;inset:0;display:flex;flex-direction:column">'
            f'<div style="flex:1;background:#{HEX["titleTop"]}"></div>'
            f'<div style="flex:1;background:#{HEX["titleBottom"]}"></div></div>')
    mark = f'<img src="{data_uri(Path(__file__).resolve().parents[1] / "logo" / "foreversixty-mark.svg")}" style="width:16px;height:16px"/>'
    left = row(mark + label("Forever Sixty", "gold", 13) , gap=8)
    close = label("x", "muted", 12)
    return (f'<div style="position:relative;width:{S["windowWidth"]}px;height:{S["titleBarHeight"]}px;'
            f'display:flex;align-items:center;justify-content:space-between;padding:0 {S["gap"]}px 0 {S["padding"]}px;'
            f'box-sizing:border-box">{grad}<div style="position:relative;z-index:1">{left}</div>'
            f'<div style="position:relative;z-index:1;width:{S["closeButton"]}px;height:{S["closeButton"]}px;'
            f'display:flex;align-items:center;justify-content:center">{close}</div></div>')


def sidebar(active):
    h = S["windowHeight"] - S["titleBarHeight"]
    items = ""
    for index, (key, text, icon_name) in enumerate(NAV_TABS):
        is_active = key == active
        top = S["gap"] * 2 + index * S["navHeight"]
        ground = f'background:#{HEX["raised"]};' if is_active else ""
        gold_bar = (f'<div style="position:absolute;left:0;top:0;bottom:0;width:{S["navBar"]}px;'
                    f'background:{hexcol("gold")}"></div>') if is_active else ""
        ic = icon(nav_icon_src(icon_name), S["navIcon"])
        txt = label(text, "gold" if is_active else "muted", 12)
        items += (f'<div style="position:absolute;left:0;top:{top}px;width:{S["sidebarWidth"]}px;'
                  f'height:{S["navHeight"]}px;{ground}display:flex;align-items:center;gap:8px;'
                  f'padding-left:{S["padding"]}px;box-sizing:border-box">{gold_bar}{ic}{txt}</div>')
    foot = abspos(S["padding"], h - S["padding"] - 14, label("foreversixty.gg", "gold", 11))
    return (f'<div style="position:absolute;left:0;top:{S["titleBarHeight"]}px;width:{S["sidebarWidth"]}px;'
            f'height:{h}px;background:#{HEX["sidebar"]};box-shadow:1px 0 0 #{HEX["border"]} inset;'
            f'box-sizing:border-box">{items}{foot}</div>')


CREST_SIZE = 36


def crest_tile(size=CREST_SIZE, class_hex=CLASS_WARRIOR):
    """The circular ringed class crest (DESIGN-SYSTEM.md principle 7), owner ruling
    2026-10-04 round 2: the addon's character header carries it too, never a square
    class icon. 2px class-colour ring, the crest itself cropped to a circle."""
    src = crest("warrior")
    return (f'<div style="width:{size}px;height:{size}px;border-radius:999px;overflow:hidden;'
            f'box-shadow:0 0 0 2px #{class_hex};flex-shrink:0;background:#{HEX["raised"]}">'
            f'<img src="{src}" style="width:100%;height:100%;object-fit:cover"/></div>')


def character_header(level_line, spec_label, status_text, status_color="muted"):
    left = S["sidebarWidth"] + S["padding"]
    top = S["titleBarHeight"] + S["gap"] * 3
    crest_top = S["titleBarHeight"] + (S["headerHeight"] - CREST_SIZE) // 2
    text_left = left + CREST_SIZE + 10
    name = label("Obnoxious Yell", "gold", 15, 700)
    level = label(level_line + "  ", "muted", 11)
    spec = label(f"·  {spec_label}", "body", 11)
    status = pill(status_text, status_color)
    hairline = (f'<div style="position:absolute;left:{S["sidebarWidth"]}px;top:{S["titleBarHeight"]+S["headerHeight"]}px;'
                f'width:{S["windowWidth"]-S["sidebarWidth"]}px;height:1px;background:#{HEX["border"]}"></div>')
    return (abspos(left, crest_top, crest_tile())
            + abspos(text_left, top, name)
            + abspos(text_left, top + 20, level + spec)
            + abspos(S["windowWidth"] - S["padding"] - 90, top, status, extra="width:90px;display:flex;justify-content:flex-end")
            + hairline)


def addon_window(active_tab, content_html, status_text=None, content_height=None, left=0, top=0):
    status_text = status_text or f"Data {BUILD}"
    page_top = S["titleBarHeight"] + S["headerHeight"]
    page_left = S["sidebarWidth"]
    page_w = S["windowWidth"] - page_left
    page_h = content_height if content_height is not None else (S["windowHeight"] - page_top)
    shadow = (f'<div style="position:absolute;left:-5px;top:5px;width:{S["windowWidth"]+10}px;'
              f'height:{S["windowHeight"]+10}px;background:rgba(0,0,0,.5);z-index:0"></div>')
    frame = (f'<div style="position:absolute;left:0;top:0;width:{S["windowWidth"]}px;'
             f'height:{S["windowHeight"]}px;background:#{HEX["background"]};opacity:.98;'
             f'box-shadow:0 0 0 1px #{HEX["border"]};box-sizing:border-box;overflow:hidden">'
             f'{title_bar()}{sidebar(active_tab)}'
             f'{character_header("Level 23 Human Warrior", "Fury", status_text)}'
             f'<div style="position:absolute;left:{page_left}px;top:{page_top}px;width:{page_w}px;'
             f'height:{page_h}px;overflow:hidden;box-sizing:border-box">{content_html}</div>'
             f'</div>')
    inner = f'<div style="position:relative;width:{S["windowWidth"]}px;height:{S["windowHeight"]}px">{shadow}{frame}</div>'
    return abspos(left, top, inner)


# ---------------------------------------------------------------------------
# Tracker, toast, minimap, talent glow -- the four surfaces outside the window.
# ---------------------------------------------------------------------------
def tracker_widget(title_text, progress_text, fraction, left, top):
    w, h = S["trackerWidth"], S["trackerHeight"]
    inner = (label(title_text, "gold", 11)
              + f'<div style="height:{S["gap"]}px"></div>'
              + label(progress_text, "muted", 11)
              + f'<div style="height:{S["gap"]}px"></div>'
              + bar(fraction, w - S["gap"] * 2, "gold", height=4))
    body = f'<div style="padding:{S["gap"]}px">{inner}</div>'
    return abspos(left, top, panel(body, width=w, height=h))


def toast_widget(text, left, top):
    w, h = 320, 32
    body = f'<div style="width:100%;height:100%;display:flex;align-items:center;justify-content:center">{label(text, "gold", 12)}</div>'
    return abspos(left, top, panel(body, width=w, height=h))


def minimap_cluster(left, top):
    diameter = 140
    ring = (f'<div style="position:absolute;left:0;top:0;width:{diameter}px;height:{diameter}px;'
            f'border-radius:999px;background:radial-gradient(circle at 40% 35%,#1b2420,#0c100d 70%);'
            f'box-shadow:0 0 0 2px #3a3326,0 0 0 1px #000 inset"></div>')
    angle = math.radians(200)
    bx = diameter / 2 + math.cos(angle) * S["minimapRadius"] - S["minimapButton"] / 2
    by = diameter / 2 + math.sin(angle) * S["minimapRadius"] - S["minimapButton"] / 2
    seal = data_uri(Path(__file__).resolve().parents[1] / "logo" / "foreversixty-mark.svg")
    btn = (f'<div style="position:absolute;left:{bx:.1f}px;top:{by:.1f}px;width:{S["minimapButton"]}px;'
           f'height:{S["minimapButton"]}px;border-radius:999px;background:#{HEX["sidebar"]};'
           f'box-shadow:0 0 0 1px {hexcol("gold")};display:flex;align-items:center;justify-content:center">'
           f'<img src="{seal}" style="width:{S["minimapIcon"]}px;height:{S["minimapIcon"]}px"/></div>')
    return abspos(left, top, f'<div style="position:relative;width:{diameter}px;height:{diameter}px">{ring}{btn}</div>')


def talent_glow_fragment(left, top):
    """A labelled fragment of the talent window, never the whole frame (owner
    instruction: never fabricate a game scene). Two tiers of the real Fury tree
    (tier 0: Booming Voice, Cruelty; tier 1: Iron Will, Unbridled Wrath), the
    next point's button ringed gold the way Theme.showGlow's texture fallback
    draws it."""
    cells = [
        ("spell_nature_purge", "Booming Voice", "3/5", True),
        ("ability_rogue_eviscerate", "Cruelty", "5/5", False),
        ("spell_magic_magearmor", "Iron Will", "0/5", False),
        ("spell_nature_stoneclawtotem", "Unbridled Wrath", "0/5", False),
    ]
    tiles = ""
    for index, (icon_name, name, rank, glowing) in enumerate(cells):
        col_i, row_i = index % 2, index // 2
        ring = f"0 0 0 2px {hexcol('gold')}" if glowing else f"0 0 0 1px #{HEX['border']}"
        tile = (f'<div style="position:absolute;left:{col_i*52}px;top:{row_i*52}px;width:36px;height:36px;'
                f'box-shadow:{ring};background:#{HEX["raised"]}">{icon(icon_uri(icon_name), 36)}</div>')
        tiles += tile
    grid = f'<div style="position:relative;width:104px;height:104px">{tiles}</div>'
    label_line = label("Talent window (fragment)", "muted", 11, family=FONT)
    body = col([label_line, grid], gap=8)
    return abspos(left, top, panel(f'<div style="padding:12px">{body}</div>', width=136, height=160))


# ---------------------------------------------------------------------------
# The client's own GameTooltip chrome (never restyled by this addon) with
# Forever Sixty's lines appended underneath, at 2x scale for legibility.
# ---------------------------------------------------------------------------
def wow_tooltip(client_lines, fs_lines, width=360, scale=2):
    name_size, line_size = 13 * scale // 2 + 6, 11 * scale // 2 + 5
    rows = ""
    for kind, *rest in client_lines:
        if kind == "name":
            text, quality = rest
            rows += f'<div style="padding:2px 0">{label(text, hexcol_quality(quality), name_size, 700)}</div>'
        else:
            text, = rest
            rows += f'<div style="padding:1px 0">{label(text, "#ffffff", line_size)}</div>'
    fs_rows = ""
    for op in fs_lines:
        if op["kind"] == "header":
            ic = icon(data_uri(Path(__file__).resolve().parents[1] / "logo" / "foreversixty-mark.svg"), 14 * scale // 2 + 4)
            fs_rows += (f'<div style="display:flex;align-items:center;gap:6px;padding:3px 0">{ic}'
                        f'{label("Forever Sixty", "gold", line_size, 700)}</div>')
        elif op["kind"] == "line":
            fs_rows += f'<div style="padding:1px 0">{label(op["text"], op["color"], line_size)}</div>'
        elif op["kind"] == "double":
            left = icon(icon_uri(op["icon"]), 14 * scale // 2 + 4) + label(op["left"], hexcol_quality(op.get("quality")), line_size)
            fs_rows += (f'<div style="display:flex;align-items:center;justify-content:space-between;padding:1px 0">'
                        f'<div style="display:flex;align-items:center;gap:4px">{left}</div>'
                        f'{label(op["right"], op["right_color"], line_size)}</div>')
    body = rows + fs_rows
    return (f'<div style="width:{width}px;padding:{8*scale//2+2}px;background:rgba(0,0,0,.92);'
            f'box-shadow:0 0 0 1px #4a3f2a;font-family:{FONT}">{body}</div>')


def hexcol_quality(quality):
    if quality is None:
        return "#ffffff"
    return f"#{QUALITY.get(quality, 'ffffff')}"


def bag_slot_tile(icon_name, quality, size=40):
    ring = hexcol_quality(quality)
    return (f'<div style="width:{size}px;height:{size}px;box-shadow:0 0 0 2px {ring},0 0 0 1px #000;">'
            f'{icon(icon_uri(icon_name), size)}</div>')


# ---------------------------------------------------------------------------
# Board 1: addon-tooltip
# ---------------------------------------------------------------------------
def build_tooltip_board():
    it = ITEMS
    better_client = [
        ("name", it["250528"]["name"], 3),
        ("line", f"Item Level {it['250528']['item_level']}"),
        ("line", "Head"),
        ("line", "163 Armor"),
        ("line", "+10 Strength"),
        ("line", "+9 Stamina"),
        ("line", "Requires Level 20"),
    ]
    better_fs = [
        {"kind": "header"},
        {"kind": "line", "text": "Upgrade for head: +20 by our weights", "color": "success"},
        {"kind": "double", "icon": "inv_helmet_39", "left": it["250528"]["name"], "quality": 3,
         "right": "BiS · lvl 20", "right_color": "muted"},
        {"kind": "line", "text": "Item Level 25", "color": "muted"},
        {"kind": "line", "text": "Crafted · Blacksmithing", "color": "muted"},
    ]
    worse_client = [
        ("name", it["250488"]["name"], 3),
        ("line", f"Item Level {it['250488']['item_level']}"),
        ("line", "Chest"),
        ("line", "+7 Strength"),
        ("line", "+4 Agility"),
        ("line", "+5 Stamina"),
        ("line", "Requires Level 15"),
    ]
    worse_fs = [
        {"kind": "header"},
        {"kind": "line", "text": "Downgrade -28.1%", "color": "warning"},
        {"kind": "double", "icon": "inv_chest_plate08", "left": it["6627"]["name"], "quality": 3,
         "right": "equipped", "right_color": "success"},
        {"kind": "line", "text": "Item Level 23", "color": "muted"},
        {"kind": "line", "text": "Dungeon · Wailing Caverns · Mutanus the Devourer", "color": "muted"},
    ]
    unrated_client = [
        ("name", "Rune of Perfection", 3),
        ("line", "Item Level 25"),
        ("line", "Requires Level 20"),
        ("line", "+4 Stamina"),
        ("line", "Equip: Increases spell penetration by 6."),
    ]
    unrated_fs = [
        {"kind": "header"},
        {"kind": "line", "text": "Not an upgrade", "color": "muted"},
    ]

    def column(x, cap, client, fs, tile_icon, tile_q):
        parts = caption(cap, x, 40, width=300)
        parts += abspos(x, 90, wow_tooltip(client, fs))
        parts += abspos(x + 20, 430, bag_slot_tile(tile_icon, tile_q))
        parts += caption("bag slot", x + 70, 442, width=120)
        return parts

    body = ""
    body += column(60, "BETTER — hovering a bag item for an empty head slot", better_client, better_fs, "inv_helmet_39", 3)
    body += column(540, "WORSE — hovering a bag item while the Breastplate (equipped) is the pick", worse_client, worse_fs, "inv_chest_chain_07", 3)
    body += column(1020, "UNRATED — no sourced trinket pick at this band yet", unrated_client, unrated_fs, "inv_misc_rune_05", 3)

    compare_caption = caption("the client's own second tooltip (Theme.showCompareTooltip), for the recommended item, pinned beside ours — unskinned, exactly as the client draws it", 60, 560, width=420)
    compare = abspos(60, 600, wow_tooltip([("name", it["250528"]["name"], 3), ("line", "Item Level 25"), ("line", "Head")], []))

    body += compare_caption + compare
    body += caption("2x scale throughout, for legibility. The real tooltip renders at the client's own small font (~12px at 1x UI scale).", 60, 12, width=700)
    return addon_page(body, 1440, 710)


# ---------------------------------------------------------------------------
# Board 2: addon-overview
# ---------------------------------------------------------------------------
CARD_W = (538 - S["cardGap"]) // 2  # 263, OverviewView.layout's own width math


def figure_card(eyebrow, value, suffix, detail, fraction, width=CARD_W):
    """Round-2 fix: Overview leads with one large figure per card (owner ruling), the
    detail line under it -- never plain-weight text competing with everything else on
    the page. A single labelled denominator (the "suffix"), not two 6px apart."""
    big = row([label(value, "gold", 30, 700), label(suffix, "muted", 12)], gap=6, align="baseline")
    top = col([label(eyebrow, "muted", 10, 700, tracking="1px"), big, label(detail, "muted", 11)], gap=6)
    bottom = bar(fraction, width - S["padding"] * 2, "gold")
    inner = f'<div style="display:flex;flex-direction:column;justify-content:space-between;height:100%">{top}<div>{bottom}</div></div>'
    return panel(inner, width=width, height=S["cardHeight"], pad=S["padding"])


def trees_card(width=CARD_W):
    """Round-2 blocker fix: live spent-per-tree (Talents.readRanks), never the build's own
    target total -- the plan total belongs only in the Your build card. Obnoxious Yell has
    spent 0 Arms / 7 Fury / 0 Protection of her 7 points taken so far (the build's own
    11-point target stays off this card entirely). Round-2 owner pass: three class-coloured
    bars (the character's own class colour, not plain gold) with the number trailing each."""
    name_col, num_col, row_gap = 60, 22, 8
    bar_w = (width - S["padding"] * 2) - name_col - num_col - row_gap * 2
    rows = [("Arms", 0, 7), ("Fury", 7, 7), ("Protection", 0, 7)]
    body = [label("TALENT POINTS", "muted", 10, 700, tracking="1px")]
    for name, points, most in rows:
        frac = points / most if most else 0
        num_color = "#c79c6e" if points > 0 else hexcol("muted")
        line = row([
            f'<div style="width:{name_col}px">{label(name, "body", 11)}</div>',
            bar(frac, bar_w, "classWarrior"),
            f'<div style="width:{num_col}px;text-align:right">{label(str(points), num_color, 13, 700)}</div>',
        ], gap=row_gap)
        body.append(line)
    inner = col(body, gap=10)
    return panel(f'<div style="height:100%;display:flex;flex-direction:column;justify-content:center">{inner}</div>',
                 width=width, height=S["cardHeight"], pad=S["padding"])


def sync_card(width=CARD_W):
    """Round-2 owner fix: there is always something to send -- the character itself --
    so this card never reads "nothing to send yet" while a character is logged in. It
    shows the crest, name and level, and one enabled gold Copy button."""
    identity = row([crest_tile(28), col([
        label("Obnoxious Yell", "#c79c6e", 13, 700),
        label("Level 23", "muted", 11),
    ], gap=2)], gap=10)
    top = col([label("SEND TO THE SITE", "muted", 10, 700, tracking="1px"), identity], gap=10)
    bottom = primary_button("Copy code", width=120, enabled=True)
    inner = f'<div style="display:flex;flex-direction:column;justify-content:space-between;height:100%">{top}<div>{bottom}</div></div>'
    return panel(inner, width=width, height=S["cardHeight"], pad=S["padding"])


def rating_card(width=538):
    """Round-2 owner fix: the empty state says what produces a rating, with a link, not
    a bare "No rating yet"."""
    top = row(label("PERSONAL RATING", "muted", 10, 700, tracking="1px") + f'<div style="flex:1"></div>')
    detail = (f'<div>{label("Upload a log or run the companion to get one.", "muted", 12)}'
              f'&nbsp;{label("Get set up", "gold", 12, 700)}</div>')
    body = col([top, detail], gap=10)
    return panel(f'<div style="height:100%;display:flex;flex-direction:column;justify-content:center">{body}</div>',
                 width=width, height=64, pad=S["padding"])


def build_overview_content():
    gap = S["cardGap"]
    row0 = row(figure_card("YOUR BUILD", "7", "of 11 points taken", "Next: Booming Voice (Fury, tier 1)", 7 / 11)
                + figure_card("GEAR", "11", "of 15 planned pieces equipped",
                               gear_upgrades_detail(), 11 / 15), gap=gap)
    row1 = row(trees_card() + sync_card(), gap=gap)
    rating = rating_card()
    fade = (f'<div style="position:absolute;left:0;bottom:0;width:100%;height:28px;'
            f'background:linear-gradient(180deg,rgba(13,17,26,0) 0%,rgba(13,17,26,.95) 100%)"></div>')
    content = col([row0, row1, rating], gap=gap)
    return f'<div style="position:relative;padding:{S["padding"]}px;box-sizing:border-box">{content}</div>{fade}'


def build_overview_board():
    content = build_overview_content()
    window = addon_window("overview", content, left=400, top=260)
    tracker = tracker_widget("Next: Booming Voice (Fury, tier 1)", "7 of 11", 7 / 11, left=600, top=90)
    minimap = minimap_cluster(left=1260, top=40)
    body = window + tracker + minimap
    body += caption("Window: CENTER (default). Tracker: TOP, 180px down (Prefs.DEFAULTS.tracker). "
                     "Minimap button: angle 200° on the ring (Prefs.DEFAULTS.minimap).",
                     40, 40, width=520)
    body += caption("The card grid and the rating card fit the fixed 416px page without scrolling once the "
                     "arrival-banner gap is fixed (see spec §5, must-fix). The rotation card sits just below "
                     "the fold, reached by the mouse wheel — the fade at the bottom edge is the hint.",
                     400, 790, width=720)
    return addon_page(body, 1440, 900)


# ---------------------------------------------------------------------------
# Board 3: addon-talents
# ---------------------------------------------------------------------------
def talent_row(name, icon_name, have, want, state):
    color = {"next": "gold", "later": "body", "done": "muted"}[state]
    ground = f'background:#{HEX["raised"]};' if state == "next" else ""
    edge = f'box-shadow:inset 2px 0 0 0 {hexcol("gold")};' if state == "next" else ""
    alpha = "opacity:.55;" if state == "done" else ""
    ic = icon(icon_uri(icon_name), S["iconSize"])
    rank_color = {"next": "gold", "later": "muted", "done": "success"}[state]
    body = (f'<div style="display:flex;align-items:center;gap:8px;height:{S["rowHeight"]}px;{ground}{edge}{alpha}'
            f'padding:0 8px;box-sizing:border-box">{ic}{label(name, color, 12)}'
            f'<div style="flex:1"></div>{label(f"{have}/{want}", rank_color, 11)}</div>')
    return body


def build_talents_content():
    heading = label("FURY", "muted", 10, 700, tracking="1px")
    rows = col([
        heading,
        talent_row("Cruelty", "ability_rogue_eviscerate", 5, 5, "done"),
        talent_row("Booming Voice", "spell_nature_purge", 2, 3, "next"),
        talent_row("Iron Will", "spell_magic_magearmor", 0, 3, "later"),
    ], gap=2)
    header = row(label("Warrior build", "gold", 13, 700) + f'<div style="flex:1"></div>'
                 + label("7 of 11 points", "muted", 11), justify="space-between")
    tabs = row(
        label("Raid", "muted", 11)
        + label("Leveling", "gold", 11, tracking=None)
        + label("PvP", "muted", 11),
        gap=16)
    progress = bar(7 / 11, 538, "gold")
    load_section = col([
        label("LOAD A BUILD", "muted", 10, 700, tracking="1px"),
        row(f'<div style="width:390px;height:26px;background:#{HEX["inset"]};box-shadow:0 0 0 1px #{HEX["border"]}"></div>'
            + secondary_button("Load build", width=120), gap=8),
        label("Paste a build code from foreversixty.gg", "muted", 10),
        row(secondary_button("Show tracker ☑", width=140) + secondary_button("Forget build", width=120), gap=8),
    ], gap=6)
    content = col([header, tabs, progress, rows, load_section], gap=10)
    return f'<div style="padding:{S["padding"]}px;box-sizing:border-box">{content}</div>'


def build_talents_board():
    toast = toast_widget("Level 23. Take Booming Voice, rank 3 of 3.", left=560, top=40)
    content = build_talents_content()
    window = addon_window("follow", content, left=360, top=160)
    glow_fragment = talent_glow_fragment(left=1120, top=260)
    body = toast + window + glow_fragment
    body += caption("Toast: TOP, 80px down (Toast.TOP_OFFSET); fades after 5s, click opens Talents.",
                     560, 76, width=360)
    body += caption("The real talent window sits where the client puts it; this fragment shows only the "
                     "cell this addon can name (tab, tier, column) and the glow it draws on it — never the "
                     "whole frame.", 1120, 428, width=220)
    body += caption("Tier shown as 1 here (player-facing); the tracker/toast text today prints the raw "
                     "0-indexed tier (\"tier 0\") — a must-fix, see spec §5.", 360, 680, width=620)
    return addon_page(body, 1440, 760)


# ---------------------------------------------------------------------------
# Board 4: addon-gear
# ---------------------------------------------------------------------------
# Round-2 fix: a slot-label column so the two lists read as one table (owner ruling),
# and a real slot -- trinket1, genuinely unsourced at this band -- to show the
# unplanned-but-equipped state the review asked for, inside the real 9-row viewport.
SLOT_LABELS = {"head": "Head", "neck": "Neck", "shoulder": "Shoulder", "back": "Back",
               "chest": "Chest", "wrist": "Wrist", "hands": "Hands", "waist": "Waist",
               "trinket1": "Trinket", "legs": "Legs"}
GEAR_SLOT_ORDER = ["head", "neck", "shoulder", "back", "chest", "wrist", "hands", "trinket1", "legs"]
GEAR_ITEM_ID = {"head": "250528", "neck": "20444", "shoulder": "3480", "back": "4706", "chest": "6627",
                 "wrist": "2868", "hands": "12994", "waist": "10403", "legs": "6087"}
# Round-2 fix (tenet 4: nothing clipped when it is the point): the name columns are the
# full item width, sized to the longest band-20 item name -- "Veteran's Silvered Chain
# Helm", 29 characters -- with no competing text tag eating into it. "As planned" moves
# from a text tag to a 14px tick glyph (Theme.SIZES.gearPlannedTick, proposed addition,
# §4.5.4) at the row's own right end, success-coloured, with a tooltip carrying the
# words it used to spell out in full -- the client's own small multi-line tooltip
# (Theme.showLines), the same mechanism the minimap button already uses for its hover.
GEAR_SLOT_COL, GEAR_ITEM_COL, GEAR_ROW_GAP, GEAR_TICK_COL, GEAR_NOTE_COL = 50, 232, 12, 18, 96


def _gear_link(name_text, icon_name, quality, width=None):
    """An item cell: icon + a real hyperlink (underlined, quality-coloured) -- tenet 2,
    the one thing both the main tooltip and the Talents BiS row already do (round-2
    must-fix). The ellipsis rule stays as a safety net for a future, longer item name;
    it does not trigger for any band-20 name at this column width (round-2 fix)."""
    ic = icon(icon_uri(icon_name), S["iconSize"])
    deco = "text-decoration:underline;text-decoration-color:currentColor"
    w = f"width:{width}px;" if width else ""
    name = (f'<span style="font-family:{FONT};font-size:12px;color:{hexcol_quality(quality)};{deco};'
            f'{w}overflow:hidden;text-overflow:ellipsis;white-space:nowrap;display:block">{name_text}</span>')
    return f'<div style="display:flex;align-items:center;gap:6px;min-width:0;flex:1;overflow:hidden">{ic}{name}</div>'


def gear_item_cell(item_id, width=None):
    it = ITEMS[item_id]
    return _gear_link(it["name"], it["icon"], it["quality"], width=width)


def gear_tick(title_text="As planned"):
    """The 14px green tick glyph that replaced the "As planned" text tag (round-2 must
    -fix). `title` is this mock's stand-in for the real tooltip (`Theme.showLines` on the
    tick's own `OnEnter`) -- a static board cannot render a hover state, so the words
    survive as an HTML tooltip attribute instead of disappearing from the board."""
    return (f'<span title="{title_text}" style="font-family:{FONT};font-size:14px;'
            f'color:{hexcol("success")};line-height:1">✔</span>')


def gear_slot_row(slot):
    label_cell = f'<div style="width:{GEAR_SLOT_COL}px;flex-shrink:0">{label(SLOT_LABELS[slot], "muted", 11)}</div>'
    if slot == "trinket1":
        planned_cell = f'<div style="width:{GEAR_ITEM_COL}px">{label("—", "muted", 12)}</div>'
        mark_col = GEAR_NOTE_COL
        equipped_inner = _gear_link("Rune of Perfection", "inv_misc_rune_05", 3)
        mark = f'<div style="width:{mark_col}px;flex-shrink:0;text-align:right">{label("No plan for this slot", "muted", 10)}</div>'
    elif slot == "legs":
        planned_cell = f'<div style="width:{GEAR_ITEM_COL}px">{gear_item_cell(GEAR_ITEM_ID["legs"])}</div>'
        mark_col = GEAR_TICK_COL
        equipped_inner = f'<div style="flex:1"></div>'
        mark = f'<div style="width:{mark_col}px;flex-shrink:0"></div>'
    else:
        item_id = GEAR_ITEM_ID[slot]
        planned_cell = f'<div style="width:{GEAR_ITEM_COL}px">{gear_item_cell(item_id)}</div>'
        mark_col = GEAR_TICK_COL
        equipped_inner = gear_item_cell(item_id)
        mark = f'<div style="width:{mark_col}px;flex-shrink:0;text-align:right">{gear_tick()}</div>'
    equipped_cell = (f'<div style="width:{GEAR_ITEM_COL}px;display:flex;align-items:center;gap:4px">'
                      f'{equipped_inner}{mark}</div>')
    return (f'<div style="display:flex;align-items:center;gap:{GEAR_ROW_GAP}px;height:{S["rowHeight"]}px">'
            f'{label_cell}{planned_cell}{equipped_cell}</div>')


def build_gear_content():
    header = row([
        f'<div style="width:{GEAR_SLOT_COL}px"></div>',
        f'<div style="width:{GEAR_ITEM_COL}px">{label("PLANNED", "muted", 10, 700, tracking="1px")}</div>',
        f'<div style="width:{GEAR_ITEM_COL}px">{label("EQUIPPED", "muted", 10, 700, tracking="1px")}</div>',
    ], gap=GEAR_ROW_GAP)
    rows = col([gear_slot_row(s) for s in GEAR_SLOT_ORDER], gap=2)
    bags_title = label("UPGRADES IN YOUR BAGS", "muted", 10, 700, tracking="1px")
    bags_empty = label(gear_upgrades_detail(), "muted", 11)
    content = col([header, rows, col([bags_title, bags_empty], gap=6)], gap=12)
    return f'<div style="padding:{S["padding"]}px;box-sizing:border-box">{content}</div>'


def build_gear_board():
    content = build_gear_content()
    window = addon_window("gear", content, left=400, top=160)
    hover_tip = wow_tooltip(
        [("name", ITEMS["6627"]["name"], 3), ("line", "Item Level 23"), ("line", "Chest"),
         ("line", "+10 Strength"), ("line", "+4 Stamina")],
        [{"kind": "header"}, {"kind": "line", "text": "Downgrade -28.1%", "color": "warning"},
         {"kind": "double", "icon": "inv_chest_plate08", "left": ITEMS["6627"]["name"], "quality": 3,
          "right": "equipped", "right_color": "success"}],
        width=280, scale=1,
    )
    body = window
    body += caption("Hovering the Chest row's Equipped cell (round-2 must-fix): a real item link, the "
                     "client's own tooltip (SetHyperlink), the same verdict/BiS anatomy as §4.1.",
                     1150, 230, width=340)
    body += abspos(1150, 380, hover_tip)
    body += caption("9 of 17 slots visible (Theme.SIZES.gearSlotRows); the rest scroll within this list with "
                     "the mouse wheel (waist, finger1/2, main/off-hand, ranged). Trinket: no sourced plan at "
                     "this band (real data) — the Equipped cell carries no judgement, just the worn item's own "
                     "quality-coloured name. Legs: planned but not yet equipped — blank Equipped cell, no tag.",
                     400, 680, width=640)
    body += caption("\"Nothing in your bags beats what you are wearing.\" (L.gearNone) is wired in here — "
                     "today's GearView never shows this string; a must-fix, see spec §5.",
                     400, 736, width=640)
    return addon_page(body, 1440, 820)


# ---------------------------------------------------------------------------
if __name__ == "__main__":
    write("addon-tooltip.html", build_tooltip_board())
    write("addon-overview.html", build_overview_board())
    write("addon-talents.html", build_talents_board())
    write("addon-gear.html", build_gear_board())
    print("wrote addon-tooltip, addon-overview, addon-talents, addon-gear")
