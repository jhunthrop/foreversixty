"""Mock boards for the addon UX pass (design/specs/2026-11-08-addon-ux.md).

    python3 design/mocks/addon_ux.py            # every board
    python3 design/mocks/addon_ux.py panel-raider-60 welcome

Writes design/reviews/2026-11-08-addon-ux/<board>.png, rendered by Playwright (python) at
device scale 2. Each board is a standalone HTML page built here, so the PNG is reproducible.

Sources, so nothing is invented:
- Item names, qualities, stats, weapon damage and icons: data/builds/1.60.1.70790/items/hunter.json
  and data/builds/1.60.1.70790/icons/*.webp (the client's own 64 px icons).
- Picks, sources, drop chances and alternatives: data/builds/1.60.1.70790/bis/hunter-beast-mastery.json
  (band 20 bare and band 60 raid, alliance).
- Talent and spell icons: data/builds/1.60.1.70790/talents/hunter.json and spellicons.json.
- Rotation lines: addon/ForeverSixty/Data.lua rotations["hunter-beast-mastery"] (bands 10, 22, 60).
- The "before" boards replay today's line order from addon/ForeverSixty/surfaces/ItemTooltip.lua,
  PlayerTooltip.lua, RewardMarks.lua and Chip.lua with the copy in Locale.lua, and match the
  owner's beta captures in web/public/addon/*.jpg.
- The raid job: docs/contracts/fixtures/inbox-raid-assignment.json's shape with an Onyxia slot
  from api/internal/guilds/raid_assign_templates.go.
- The DPS figures on Bow Jackzon (+3.5, +2.2, +0.8) are the capture's; the Tunic of Westfall
  figure is derived here from band-20 dps_per_point (see tunic_gain); the level 60 gains are
  illustrative layout figures (the build reads Progress.pickGain) and are labelled so in the spec.

The app QR code is drawn by segno (pip install segno) at error correction H, the same encoding the
addon's build step bakes into its texture.

Font: the client draws every tooltip and frame in Friz Quadrata, which ships only inside the
client's archive. Marcellus (apps/shared/fonts, OFL) is the closest flared glyphic face in the
repo and stands in for it.
"""
import base64
import io
import json
import os
import sys
from pathlib import Path

from PIL import Image

from addon_tokens import FONT_FILES, color, role

ROOT = Path(__file__).resolve().parents[2]
BUILD = ROOT / 'data/builds/1.60.1.70790'
ICONS = BUILD / 'icons'
# MOCK_OUT: write the boards elsewhere (the BA1 gate writes design/reviews/2026-11-10-ba1/captures).
OUT = Path(os.environ['MOCK_OUT']) if os.environ.get('MOCK_OUT') else ROOT / 'design/reviews/2026-11-08-addon-ux'
FONT = ROOT / 'apps/shared/fonts/Marcellus-Regular.ttf'

ITEMS = {it['id']: it for it in json.loads((BUILD / 'items/hunter.json').read_text())['items']}
BIS = json.loads((BUILD / 'bis/hunter-beast-mastery.json').read_text())


def band(level, preset='bare', faction='alliance'):
    return next(b for b in BIS['bands'] if b['band'] == level and b['preset'] == preset and b['faction'] == faction)


BAND20 = band(20)
DPP20 = {w['stat']: w['dps_per_point'] for w in BAND20['weights']}

# --- Blizzard's own font colours (FrameXML's GlobalColors), and the item quality colours ----------
C = {
    'green': '#1aff1a',   # GREEN_FONT_COLOR: "do it"
    'gold': '#ffd100',    # NORMAL_FONT_COLOR: yours, best in slot
    'white': '#ffffff',   # HIGHLIGHT_FONT_COLOR: information
    'grey': '#b8b8b8',    # Theme.tooltipMuted: no / details (GRAY_FONT_COLOR 808080 fails contrast here)
    'red': '#ff2020',     # RED_FONT_COLOR: the client's own refusals
    'blue': '#0070dd',
}
# Item names in the site's rarity colours (Theme.QUALITY_HEX: rare and epic in the text tokens), and the
# addon's own colours, all from design/tokens/tokens.json (addon_tokens), as Tokens.lua carries them.
QUALITY = {0: color('rarity.poor'), 1: color('rarity.common'), 2: color('rarity.uncommon'),
           3: color('rarity.rareText'), 4: color('rarity.epicText'), 5: color('rarity.legendary')}
TIER = {'Common': QUALITY[1], 'Uncommon': QUALITY[2], 'Rare': QUALITY[3], 'Epic': QUALITY[4], 'Legendary': QUALITY[5]}
ADDON = {'gold': color('base.gold'), 'muted': color('base.muted'), 'body': color('base.text'),
         'success': color('outcome.kill'), 'tmuted': color('addon.tooltipMuted'), 'warning': color('outcome.wipe'),
         'bg': color('base.bg'), 'raised': color('base.raised'), 'hover': color('base.hover'),
         'line': color('base.line'), 'cream': color('forever.cream'), 'strong': color('base.strong'),
         'blue': color('forever.blue'), 'blue_edge': color('forever.blueEdge'), 'teal': color('forever.teal'),
         'teal_edge': color('forever.tealEdge'), 'inset': color('addon.inset')}
BODY_FACES = {w: f for w, f in FONT_FILES['body'].items()}


def b64(data: bytes, mime: str) -> str:
    return f'data:{mime};base64,' + base64.b64encode(data).decode()


def icon_uri(name: str) -> str:
    return b64((ICONS / f'{name}.webp').read_bytes(), 'image/webp')


def tga_uri(path: Path) -> str:
    buf = io.BytesIO()
    Image.open(path).save(buf, 'PNG')
    return b64(buf.getvalue(), 'image/png')


CREST = b64((ROOT / 'web/public/icons/hd/crests/hunter.webp').read_bytes(), 'image/webp')
FONT_URI = b64(FONT.read_bytes(), 'font/ttf')


def item(item_id):
    return ITEMS[item_id]


def iname(item_id, name=None):
    it = item(item_id)
    return f'<span style="color:{QUALITY[it["quality"]]}">{name or it["name"]}</span>'


def img(name, size, extra=''):
    return f'<img class="ic" src="{icon_uri(name)}" style="width:{size}px;height:{size}px;{extra}">'


# The client's bags-greenarrow atlas, redrawn: a green up arrow with a dark rim.
ARROW = ('<svg class="arrow" viewBox="0 0 12 12" width="{s}" height="{s}"><path d="M6 0.6 L11.2 6.4 L8 6.4 L8 11.4 '
         'L4 11.4 L4 6.4 L0.8 6.4 Z" fill="#35e04a" stroke="#06300b" stroke-width="1.1" stroke-linejoin="round"/>'
         '<path d="M6 2.2 L9.2 5.6 L6.9 5.6 L6.9 10.3" fill="none" stroke="#b6ffbf" stroke-width="0.8" opacity=".7"/></svg>')


def arrow(size=12):
    return ARROW.format(s=size)


COIN = '<span class="coin" style="background:radial-gradient(circle at 35% 35%,{a},{b})"></span>'
SILVER = COIN.format(a='#f4f4f4', b='#8c8c8c')
COPPER = COIN.format(a='#f1a46b', b='#8a4a1f')
GOLDC = COIN.format(a='#ffe27a', b='#a07812')

SITE_FONTS = ''.join(f"@font-face {{{{ font-family: Site; font-weight: {w}; src: url({b64(f.read_bytes(), 'font/ttf')}); }}}}"
                     for w, f in BODY_FACES.items())
TITLE_URI = b64(FONT_FILES['title'][400].read_bytes(), 'font/ttf')
BODY, SMALL = role('body'), role('small')

CSS = f'''
@font-face {{ font-family: Friz; src: url({FONT_URI}); }}
@font-face {{ font-family: Title; src: url({TITLE_URI}); }}
{SITE_FONTS.replace('{{', '{').replace('}}', '}')}
* {{ box-sizing: border-box; }}
body {{ margin: 0; background: #000; }}
#board {{ display: inline-block; padding: 28px; font-family: Friz, serif; color: #fff;
  background: radial-gradient(ellipse at 30% 20%, #23343a 0%, #121b22 45%, #090d12 100%); }}
.cap {{ font-family: -apple-system, 'Helvetica Neue', Arial, sans-serif; font-size: 13px; color: #c9c2b2;
  letter-spacing: .02em; margin: 0 0 10px; }}
.cap b {{ color: #e5b955; font-weight: 600; }}
.row {{ display: flex; gap: 22px; align-items: flex-start; }}
.col {{ display: flex; flex-direction: column; gap: 14px; }}
.tt {{ position: relative; background: rgba(8, 10, 22, .93); border-radius: 5px; padding: 7px 10px 8px;
  border: 1px solid #9095a3; box-shadow: inset 0 0 0 1px #20232f, 0 0 0 1px #05060a, 0 4px 14px rgba(0,0,0,.6);
  font-size: 13.5px; line-height: 17px; min-width: 200px; }}
.tt .t {{ font-size: 15.5px; line-height: 20px; }}
.tt .l {{ display: flex; justify-content: space-between; gap: 24px; white-space: nowrap; }}
.tt .gap {{ height: 9px; }}
.tt .tag {{ position: absolute; top: -19px; left: 2px; font-size: 12px; color: {C['grey']}; }}
.coin {{ display: inline-block; width: 10px; height: 10px; border-radius: 50%; margin: 0 3px 0 2px;
  vertical-align: -1px; box-shadow: 0 0 0 1px rgba(0,0,0,.6); }}
.arrow {{ vertical-align: -1px; margin-right: 3px; }}
.ic {{ border-radius: 3px; box-shadow: 0 0 0 1px #000, 0 0 0 2px #3b3f4a; vertical-align: middle; }}
.hl {{ outline: 1.5px dashed rgba(229,185,85,.75); outline-offset: 2px; border-radius: 2px; }}
.note {{ font-family: -apple-system, 'Helvetica Neue', Arial, sans-serif; font-size: 12px; color: #9a9484;
  max-width: 330px; line-height: 16px; }}
.note b {{ color: #e9e4d8; font-weight: 600; }}
/* The 1.60 client's dark frame (Combined Backpack, the character pane): bronze rim, portrait ring. */
.frame {{ position: relative; background: linear-gradient(#141826, #0c0f18 40%, #0a0d15); border-radius: 6px;
  border: 2px solid #4a3e26; box-shadow: inset 0 0 0 1px #000, inset 0 0 0 2px #2a2416, 0 8px 28px rgba(0,0,0,.7);
  color: #fff; }}
.frame .title {{ height: 30px; display: flex; align-items: center; justify-content: center; color: {C['gold']};
  font-size: 15px; background: linear-gradient(#1d1a12, #11100b); border-bottom: 1px solid #3b311d;
  border-radius: 4px 4px 0 0; }}
.frame .x {{ position: absolute; top: 5px; right: 6px; width: 20px; height: 20px; border-radius: 3px;
  background: linear-gradient(#8a1c12, #4e0b06); border: 1px solid #c7a24a; color: #ffd9a0; font-size: 13px;
  display: flex; align-items: center; justify-content: center; font-family: Arial, sans-serif; font-weight: bold; }}
.frame .portrait {{ position: absolute; top: -14px; left: -14px; width: 58px; height: 58px; border-radius: 50%;
  background: #000 center/100% no-repeat; box-shadow: 0 0 0 3px #b8913f, 0 0 0 4px #3b2d10, 0 3px 8px #000; }}
.frame .body {{ padding: 12px 14px 14px; }}
.sect {{ color: {C['gold']}; font-size: 13.5px; margin: 2px 0 6px; display: flex; justify-content: space-between;
  align-items: baseline; }}
.sect .r {{ color: {C['grey']}; font-size: 12px; }}
.rule {{ height: 1px; background: linear-gradient(90deg, transparent, #3d3624 15%, #3d3624 85%, transparent);
  margin: 10px 0 9px; }}
.irow {{ display: flex; gap: 9px; align-items: center; margin: 0 0 7px; }}
.irow .tx {{ flex: 1; min-width: 0; }}
.irow .n {{ font-size: 13.5px; line-height: 16px; display: flex; justify-content: space-between; gap: 10px;
  white-space: nowrap; }}
.irow .s {{ font-size: 12px; line-height: 15px; color: {C['grey']}; white-space: nowrap; overflow: hidden;
  text-overflow: ellipsis; }}
.dim {{ opacity: .45; }}
.btn {{ display: inline-flex; align-items: center; justify-content: center; height: 22px; padding: 0 12px;
  font-size: 12.5px; color: {C['gold']}; border-radius: 3px; border: 1px solid #6b5626;
  background: linear-gradient(#5a1610, #2e0906); box-shadow: inset 0 1px 0 rgba(255,220,150,.15); white-space: nowrap; }}
.btn.grey {{ color: #e8e8e8; border-color: #4b4f5c; background: linear-gradient(#2b2f3b, #171a22); }}
.small {{ font-size: 12px; color: {C['grey']}; }}
/* The addon's own frames (Widgets.umberFrame, titleBand, clientButton) as the site's plates: the raised
   ground in one hairline, Open Sans text, Marcellus titles in cream, the site's button recipes. */
.plate {{ position: relative; background: {ADDON['raised']}; border: 1px solid {ADDON['line']};
  box-shadow: 0 8px 28px rgba(0,0,0,.7); color: {ADDON['body']}; font-family: Site, sans-serif;
  font-size: {BODY['size'] + 1}px; }}
.plate .band {{ height: 64px; display: flex; align-items: center; gap: 12px; padding: 0 42px 0 12px;
  border-bottom: 1px solid {ADDON['line']}; }}
.plate .band .t {{ font-family: Title, serif; font-size: {role('title')['size']}px; color: {ADDON['cream']};
  line-height: 22px; text-shadow: 1px 1px 0 #000; }}
.plate .band .m {{ font-size: {SMALL['size'] + 1}px; color: {ADDON['muted']}; margin-top: 3px; white-space: nowrap;
  overflow: hidden; text-overflow: ellipsis; }}
.plate .band .num {{ font-family: Site, sans-serif; font-weight: 600; font-size: {role('large')['size'] + 1}px; }}
.plate .band .unit {{ font-size: 12px; color: {ADDON['muted']}; }}
.plate .x {{ position: absolute; top: 6px; right: 8px; width: 20px; height: 20px; color: {ADDON['body']};
  border: 1px solid {ADDON['line']}; box-sizing: border-box;
  font-size: 14px; display: flex; align-items: center; justify-content: center; font-family: Site, sans-serif; }}
.plate .body {{ padding: 12px 14px 14px; }}
.plate .sect {{ font-family: Title, serif; color: {ADDON['cream']}; font-size: {role('heading')['size']}px; }}
.plate .sect .r {{ font-family: Site, sans-serif; color: {ADDON['muted']}; }}
.plate .rule {{ background: {ADDON['line']}; }}
.plate .small, .plate .irow .s {{ color: {ADDON['muted']}; font-size: {SMALL['size'] + 1}px; }}
.sbtn {{ display: inline-flex; align-items: center; justify-content: center; height: 22px; padding: 0 12px;
  font-family: Site, sans-serif; font-weight: 600; font-size: 12.5px; color: {ADDON['strong']};
  background: {ADDON['blue']}; border: 1px solid {ADDON['blue_edge']}; white-space: nowrap; }}
.sbtn.primary {{ background: {ADDON['teal']}; border-color: {ADDON['teal_edge']}; }}
'''


def page(body: str) -> str:
    return f'<!doctype html><html><head><meta charset="utf-8"><style>{CSS}</style></head><body>' \
           f'<div id="board">{body}</div></body></html>'


def l(left, right=None, color='#fff', rcolor=None, cls=''):
    style = f'color:{color}'
    if right is None:
        return f'<div class="l {cls}" style="{style}"><span>{left}</span></div>'
    return f'<div class="l {cls}" style="{style}"><span>{left}</span><span style="color:{rcolor or color}">{right}</span></div>'


def tooltip(lines, tag=None, width=None, extra_style=''):
    w = f'width:{width}px;' if width else ''
    t = f'<div class="tag">{tag}</div>' if tag else ''
    return f'<div class="tt" style="{w}{extra_style}">{t}{"".join(lines)}</div>'


# --- The client's own lines for the items on the boards (items/hunter.json + the capture) ----------
def short_ash_bow():
    return [
        f'<div class="t" style="color:{QUALITY[2]}">Short Ash Bow</div>', l('Soulbound'),
        l('Ranged', 'Bow'), l('12 - 23 Damage', 'Speed 1.90'), l('(9.2 damage per second)'), l('+2 Spirit'),
        l('Durability 55 / 55'), l('Requires Level 18'), l(f'Sell Price: 16{SILVER}10{COPPER}'),
    ]


def ranger_bow(bound='Binds when equipped'):
    return [
        f'<div class="t" style="color:{QUALITY[3]}">Ranger Bow</div>', l(bound),
        l('Ranged', 'Bow'), l('23 - 45 Damage', 'Speed 2.70'), l('(12.6 damage per second)'), l('+1 Agility'),
        l('Requires Level 20'), l(f'Sell Price: 24{SILVER}21{COPPER}'),
    ]


def tunic_of_westfall():
    return [
        f'<div class="t" style="color:{QUALITY[2]}">Tunic of Westfall</div>', l('Binds when picked up'),
        l('Chest', 'Leather'), l('92 Armor'), l('+11 Agility'), l('+5 Stamina'), l('Durability 85 / 85'),
        l('Requires Level 14'), l(f'Sell Price: 7{SILVER}63{COPPER}'),
    ]


def tunic_gain():
    """Tunic of Westfall over a worn Fine Leather Tunic (4 Agi, 3 Sta), band-20 BM dps_per_point."""
    tunic, worn = item(2041)['stats'], item(4243)['stats']
    return round(sum((tunic.get(s, 0) - worn.get(s, 0)) * DPP20.get(s, 0) for s in set(tunic) | set(worn)), 1)


def ring_gain():
    """Protector's Band over a worn Ring of Scorn, band-20 BM dps_per_point."""
    new, worn = item(20439)['stats'], item(3235)['stats']
    return round(sum((new.get(s, 0) - worn.get(s, 0)) * DPP20.get(s, 0) for s in set(new) | set(worn)), 1)


# =================================================================================================
# Item tooltip
# =================================================================================================
def board_item_before():
    m = ADDON['tmuted']
    lines = short_ash_bow() + [
        l('Press F6 to submit an issue for this Item', color='#3fa9ff'),
        '<div class="gap"></div>',
        l(f'<span style="color:{C["gold"]}">Forever Sixty</span><span style="color:{m}"> · Beast Mastery · Estimate</span>',
          'Equipped', rcolor=m),
        l(f'<span style="color:{m}">Next upgrade:</span> {iname(3021)} <span style="color:{ADDON["success"]}">+3.5 DPS</span>'),
        l('(+1 Agi, +3.4 weapon DPS)', color=m),
        l('Planned for your build', color=ADDON['gold']),
        l('World drop', color=m),
        l('Also good:', color=m),
        l(f'&nbsp;&nbsp;&nbsp;&nbsp;{iname(13136)} <span style="color:{m}">+2.2</span>', color=m),
        l('&nbsp;&nbsp;&nbsp;&nbsp;(needs bullets)', color=m),
        l(f'&nbsp;&nbsp;&nbsp;&nbsp;{iname(285279)} <span style="color:{m}">+0.8</span>', color=m),
    ]
    compare = tooltip([l('Next upgrade', color=m)] + ranger_bow(), width=262)
    body = (f'<p class="cap"><b>Before</b> · item tooltip today: hovering the worn Short Ash Bow, Bow Jackzon, '
            f'level 24 Beast Mastery (web/public/addon/item-tooltip-ranged.jpg)</p>'
            f'<div class="row">{tooltip(lines, width=330)}{compare}</div>'
            f'<p class="note" style="margin-top:12px;max-width:640px"><b>Eleven lines after a blank row</b>, a second '
            f'header, five colours, the gain split from its breakdown by a plan line, and a source that belongs to '
            f'the next upgrade printed under a line about the worn bow.</p>')
    return page(body)


def verdict_next(item_id, gain, source):
    return [
        l(f'Next upgrade: {iname(item_id)} <span style="color:{C["green"]}">+{gain} DPS</span>'),
        l(source, color=C['grey']),
    ]


def shift_block_bow():
    g = C['grey']
    return [
        # Stat changes only: the weapon's own DPS beside the +3.5 headline read as the same number twice
        # (owner, 2026-11-08); the type in the game's words, never what it implies.
        l('+1 Agi', color=g),
        l(f'Also good: {iname(13136)} +2.2 (gun)', color=g),
        l(f'<span style="visibility:hidden">Also good: </span>{iname(285279)} +0.8 (thrown)', color=g),
        l('In your build\'s gear list', color=g),
        l('Forever Sixty · Beast Mastery estimate', color='#8a8a8a'),
    ]


def board_item_after():
    lines = short_ash_bow() + verdict_next(3021, '3.5', 'World drop')
    compare = tooltip(ranger_bow(), tag='Next upgrade', width=262)
    body = (f'<p class="cap"><b>After</b> · the same hover: one answer, one supporting line, no header, no gap</p>'
            f'<div class="row" style="padding-top:18px">{tooltip(lines, width=290)}{compare}</div>')
    return page(body)


def board_item_after_shift():
    lines = short_ash_bow() + verdict_next(3021, '3.5', 'World drop') + shift_block_bow()
    compare = tooltip(ranger_bow(), tag='Next upgrade', width=262)
    body = (f'<p class="cap"><b>After, Shift held</b> · the client\'s compare and our details arrive together</p>'
            # No fixed width: the client's tooltip grows to its widest line, so no line crosses the border.
            f'<div class="row" style="padding-top:18px">{tooltip(lines)}{compare}</div>')
    return page(body)


def mini(title_html, lines, caption, width=270):
    """A cut-down tooltip: the item's name, an ellipsis for the client's lines, then ours."""
    head = [f'<div class="t">{title_html}</div>', l('…', color='#777')]
    return (f'<div class="col" style="gap:6px;width:{width}px"><div class="note">{caption}</div>'
            f'{tooltip(head + lines, width=width)}</div>')


def board_item_states():
    g, gold, green, wht = C['grey'], C['gold'], C['green'], C['white']
    up = arrow()
    cards = [
        mini(iname(3021), [l(f'{up}Upgrade +3.5 DPS', color=green), l('Your best in slot until level 30', color=gold)],
             '<b>1 · Upgrade, and the pick</b> · Ranger Bow in the bags'),
        mini(iname(13136), [l(f'{up}Upgrade +2.2 DPS', color=green),
                            l(f'Better: {iname(3021)} +3.5 · World drop', color=g)],
             '<b>2 · Upgrade, not the best</b> · what\'s better, in one line'),
        mini(iname(3021), [l('Your best in slot', color=gold),
                           l(f'At level 30: {iname(273029, "Golemsight Long Gun")}', color=g)],
             '<b>3 · Worn, and it is the pick</b>'),
        mini(iname(3039), [l(f'{up}In your bags: {iname(3021)} <span style="color:{green}">+3.5 DPS</span>')],
             '<b>4 · Worn, a bag item beats it</b> · the slot carries the arrow'),
        mini(iname(285279), [l('Sidegrade', color=g)], '<b>5 · Sidegrade</b> · +0.8 DPS, under the 1.0 floor; the size sits behind Shift'),
        mini(iname(13136), [l('Downgrade −1.3 DPS', color=g)], '<b>6 · Downgrade</b>'),
        mini(iname(6087), [l('Not for you until level 40', color=g)],
             '<b>7 · A Forever rule</b> · hunters wear mail at 40; the client\'s red line says why'),
        mini(iname(2042), [l('Not for your build', color=g)], '<b>8 · Nothing your spec wants</b>'),
        mini(iname(3021), [l(f'{up}Upgrade at level 20: +3.5 DPS', color=green)],
             '<b>9 · Level too low</b> · at level 18, from the auction house'),
        mini(f'<span style="color:{QUALITY[2]}">Sturdy Quiver</span>', [], '<b>10 · Ammo, bags, anything not gear</b> · silence'),
        mini(f'<span style="color:{QUALITY[1]}">Retrieving item information</span>', [],
             '<b>11 · An item the client has not sent</b> · silence, redrawn when it arrives'),
        mini(iname(20439), [l(f'{up}Upgrade +{ring_gain()} DPS over {iname(3235)}', color=green)],
             '<b>12 · Two slots (rings, trinkets, one-handers)</b> · names the one it replaces'),
    ]
    rows = ''.join(f'<div class="row" style="margin-bottom:22px">{"".join(cards[i:i + 4])}</div>'
                   for i in range(0, len(cards), 4))
    body = (f'<p class="cap"><b>After</b> · every verdict, one grammar: green acts, gold is yours, white informs, '
            f'grey says no. Shift adds the details to each.</p>{rows}')
    return page(body)


# =================================================================================================
# Quest reward
# =================================================================================================
QUEST_REWARDS = [2041, 6087, 2042]


def quest_frame(labels, arrows=()):
    cells = ''
    for idx, rid in enumerate(QUEST_REWARDS):
        it = item(rid)
        mark = (f'<div style="position:absolute;top:-3px;left:30px">{arrow(14)}</div>' if rid in arrows else '')
        label = labels.get(rid, '')
        cells += (f'<div style="position:relative;display:flex;gap:6px;align-items:center;width:168px;height:44px;'
                  f'padding:3px;background:rgba(60,40,15,.18);border:1px solid rgba(80,55,20,.45);border-radius:3px">'
                  f'{img(it["icon"], 36)}{mark}<div style="font-size:12.5px;line-height:14px;color:#2a1a08">'
                  f'{it["name"]}{label}</div></div>')
    return (f'<div class="frame" style="width:400px"><div class="title">The Defias Brotherhood</div>'
            f'<div style="margin:10px;padding:14px 14px 16px;border-radius:3px;color:#2a1a08;'
            f'background:radial-gradient(ellipse at 40% 30%,#efe0bd,#d7c193 70%,#c4ab7a);">'
            f'<div style="font-size:16px;margin-bottom:6px">Rewards</div>'
            f'<div style="font-size:12.5px;margin-bottom:10px">You will be able to choose one of these rewards:</div>'
            f'<div style="display:grid;grid-template-columns:1fr 1fr;gap:8px">{cells}</div>'
            f'<div style="font-size:12.5px;margin:12px 0 4px">You will also receive:</div>'
            f'<div style="font-size:12.5px">1{GOLDC} 75{SILVER}</div></div>'
            f'<div style="display:flex;justify-content:space-between;padding:0 10px 10px">'
            f'<span class="btn">Complete Quest</span><span class="btn grey">Cancel</span></div></div>')


def board_quest_before():
    m = ADDON['tmuted']
    labels = {2041: f'<div style="color:#8a5a00;font-size:11.5px;margin-top:2px">Best for Beast Mastery</div>'}
    lines = tunic_of_westfall() + [
        '<div class="gap"></div>',
        l(f'<span style="color:{C["gold"]}">Forever Sixty</span><span style="color:{m}"> · Beast Mastery · Estimate</span>'),
        l(f'Upgrade +{tunic_gain()} DPS over Fine Leather Tunic', color=ADDON['success']),
        l('(+7 Agi)', color=m),
        l('Planned for Beast Mastery at 20', color=ADDON['gold']),
        l('The Defias Brotherhood · Quest', color=m),
    ]
    body = (f'<p class="cap"><b>Before</b> · the quest reward today: a gold label on the pick, the item tooltip\'s '
            f'whole block on hover; "Vendor" on every reward when none is of use</p>'
            f'<div class="row">{quest_frame(labels)}{tooltip(lines, width=330)}</div>')
    return page(body)


def board_quest_after():
    gold, green = C['gold'], C['green']
    labels = {2041: f'<div style="color:#7a5200;font-size:11.5px;margin-top:2px">Best in slot</div>'}
    lines = tunic_of_westfall() + [
        l(f'{arrow()}Upgrade +{tunic_gain()} DPS', color=green),
        l('Your best in slot until level 30', color=gold),
    ]
    none_labels = {2042: '<div style="color:#5b4a2a;font-size:11.5px;margin-top:2px">Sells for most</div>'}
    body = (f'<p class="cap"><b>After</b> · the green arrow marks the upgrade (the bags\' own mark), the gold label '
            f'names the pick; the hover is two lines</p>'
            f'<div class="row">{quest_frame(labels, arrows=(2041,))}{tooltip(lines, width=250)}</div>'
            f'<p class="cap" style="margin-top:22px"><b>After, no reward is of use</b> · one muted label on the '
            f'reward that sells for the most, never "Vendor" on all three</p>'
            f'<div class="row">{quest_frame(none_labels)}<div class="note">Shown for a character none of the three '
            f'helps (the same quest at level 30, wearing better). The sell price is the client\'s own (GetItemInfo), '
            f'so it is exact.</div></div>')
    return page(body)


# =================================================================================================
# Player tooltip
# =================================================================================================
def player_lines():
    return [f'<div class="t" style="color:#aad372">Bow Jackzon</div>', l('Employed Millennials'),
            l('Level 24 High Order Skyborne (Player)'), l('Hunter'), l('Alliance'), l('PvP')]


def board_player_before():
    lines = player_lines() + [
        l('Forever Score 1240 · Rare · as of 13 Oct', color=TIER['Rare']),
        l('Shift: score details · target them, then /fs link', color=ADDON['tmuted']),
        '<div class="gap"></div>', l('&lt;Right click for Frame Settings&gt;', color=C['green']),
    ]
    body = (f'<p class="cap"><b>Before</b> · player tooltip today (web/public/addon/player-tooltip-score.jpg)</p>'
            f'{tooltip(lines, width=380)}')
    return page(body)


def board_player_after():
    g = C['grey']
    base = player_lines()
    plain = base + [l('Forever Score 1240', color=TIER['Rare'])]
    shift = base + [
        l('Forever Score 1240', color=TIER['Rare']),
        l('Gear', '310', color=g, rcolor='#fff'), l('Experience', '220', color=g, rcolor='#fff'),
        l('Parses', '410', color=g, rcolor='#fff'), l('Attendance', '180', color=g, rcolor='#fff'),
        l('Preparation (flasks, food, enchants)', '120', color=g, rcolor='#fff'),
    ]
    stale = base[:1] + [l('…', color='#777'), l('Forever Score 1240 (13 Oct)', color=TIER['Rare'])]
    body = (f'<p class="cap"><b>After</b> · one line in the tier\'s colour (the colour is the tier), Shift for the '
            f'parts; a score over a week old carries its date</p>'
            f'<div class="row">'
            f'<div class="col"><div class="note"><b>Default</b></div>{tooltip(plain, width=330)}</div>'
            f'<div class="col"><div class="note"><b>Shift held</b> (part values are layout figures; the record sums '
            f'to 1240)</div>{tooltip(shift, width=330)}</div>'
            f'<div class="col"><div class="note"><b>Score older than 7 days</b></div>{tooltip(stale, width=250)}'
            f'<div class="note">No record or no rated fights, or "Gear tips" off: no line at all.</div></div></div>')
    return page(body)


# =================================================================================================
# Chip
# =================================================================================================
def minimap(chip_html, size=180):
    return (f'<div style="position:relative;width:{size + 280}px;height:{size + 20}px">'
            f'<div style="position:absolute;right:0;top:0;width:{size}px;height:{size}px;border-radius:50%;'
            f'background:radial-gradient(circle at 45% 40%,#4b6a3a,#2a4429 55%,#16261a);'
            f'box-shadow:0 0 0 5px #8b6a35,0 0 0 7px #2d220f,0 4px 12px #000"></div>'
            f'<div style="position:absolute;right:{size + 10}px;top:12px">{chip_html}</div></div>')


def chip(text, ring=0.0, alpha=1.0, icon_alpha=1.0):
    glow = (f'box-shadow:0 0 0 2px rgba(255,209,0,{ring}),0 0 {int(10 * ring) + 2}px {int(4 * ring)}px '
            f'rgba(255,209,0,{ring * .7}),0 0 0 1px #000;' if ring else 'box-shadow:0 0 0 1px #000;')
    label = (f'<span style="font-family:Site,sans-serif;font-size:12.5px;color:{ADDON["body"]};padding:0 10px 0 6px;'
             f'white-space:nowrap">{text}</span>'
             if text else '')
    return (f'<div style="display:inline-flex;align-items:center;height:30px;opacity:{alpha};border-radius:15px;'
            f'background:{ADDON["bg"]};border:1px solid {ADDON["line"]};{glow}">'
            f'<div style="width:28px;height:28px;border-radius:50%;background:url({CREST}) center/100%;'
            f'opacity:{icon_alpha}"></div>{label}</div>')


def board_chip():
    text = 'Next talent at 25: Lethal Attacks'
    frames = ''.join(
        f'<div class="col" style="gap:4px"><div class="note">{t}</div>{minimap(chip(text, ring=r), 120)}</div>'
        for t, r in (('0.0 s', .3), ('0.8 s', 1.0), ('1.6 s; five cycles, then still', .3)))
    hover = tooltip([
        l('Click: open Forever Sixty', color='#8a8a8a'),
        l('Right-click: hide until level 25', color='#8a8a8a'),
    ], width=250)
    raid = chip('Tonight 8:00 pm: Onyxia\'s Lair')
    before = minimap(chip('', icon_alpha=.4), 140)
    body = (f'<p class="cap"><b>Before</b> · the idle chip today: a class icon at 40% beside the minimap, no words '
            f'(the owner missed it)</p>{before}'
            f'<p class="cap" style="margin-top:18px"><b>After</b> · the chip always says one thing in words. Until it is '
            f'clicked or hovered once, it pulses a gold ring for five cycles (8 s) after each login, for at most three '
            f'logins, never in combat</p>'
            f'<div class="row">{frames}</div>'
            f'<p class="cap" style="margin-top:18px"><b>After, hover</b> · the two things a click does; the line itself '
            f'repeats only when the chip had to shorten it</p>'
            f'<div class="row">{minimap(chip(text), 120)}{hover}</div>'
            f'<p class="cap" style="margin-top:18px"><b>After, raid day at 60</b> · the raid outranks every other state '
            f'from 18 hours before the pull</p>{minimap(raid, 120)}')
    return page(body)


# =================================================================================================
# Welcome: a callout docked to the chip, not a window in the middle of the screen
# =================================================================================================
def bag_slot_illustration():
    return (f'<div style="position:relative;width:40px;height:40px;border-radius:4px;background:#1b1e26;'
            f'box-shadow:inset 0 0 0 1px #000,0 0 0 1px #4a4f5c;flex:none">'
            f'<img src="{icon_uri("inv_weapon_bow_11")}" style="position:absolute;inset:3px;width:34px;height:34px;'
            f'border-radius:3px;box-shadow:0 0 0 1px {QUALITY[3]}">'
            f'<div style="position:absolute;top:-4px;right:-5px">{arrow(15)}</div></div>')


def talent_illustration():
    return (f'<div style="position:relative;width:40px;height:40px;flex:none">'
            f'<img src="{icon_uri("ability_searingarrow")}" style="position:absolute;inset:4px;width:32px;height:32px;'
            f'border-radius:4px;box-shadow:0 0 0 2px #ffd100,0 0 8px 2px rgba(255,209,0,.55)">'
            f'<div style="position:absolute;right:-6px;bottom:-5px;font-size:10.5px;padding:0 4px;border-radius:7px;'
            f'background:#ffd100;color:#1a1200">13</div></div>')


# =================================================================================================
# Loot roll
# =================================================================================================
def roll_frame(verdict_html, marked=True):
    it = item(17069)
    mark = f'<div style="position:absolute;top:-5px;right:-6px">{arrow(15)}</div>' if marked else ''
    buttons = ''.join(f'<span class="btn grey" style="padding:0 7px">{b}</span>' for b in ('Need', 'Greed', 'Pass'))
    return (f'<div class="frame" style="width:400px;padding:8px 10px;display:flex;gap:10px;align-items:center">'
            f'<div style="position:relative">{img(it["icon"], 40)}{mark}</div>'
            f'<div style="flex:1"><div style="color:{QUALITY[4]};font-size:14px">{it["name"]}</div>'
            f'<div style="font-size:12.5px;white-space:nowrap">{verdict_html}</div></div>'
            f'<div style="display:flex;gap:4px">{buttons}</div></div>')


def board_loot_roll():
    up = roll_frame(f'<span style="color:{C["green"]}">+5.8 DPS</span>')
    bis = roll_frame(f'<span style="color:{C["green"]}">+5.8 DPS</span> · '
                     f'<span style="color:{C["gold"]}">Best in slot</span>')
    no = roll_frame(f'<span style="color:{C["grey"]}">Not for you</span>', marked=False)
    body = (f'<p class="cap"><b>After</b> · the group loot roll frame: the same verdict in a few words beside the item, '
            f'no source (it is dropping in front of you)</p>'
            f'<div class="col"><div class="note"><b>An upgrade</b> · Striker\'s Mark from Magmadar; the gain is a layout '
            f'figure</div>{up}<div class="note"><b>An upgrade that is the pick</b></div>{bis}'
            f'<div class="note"><b>Not usable, or nothing your spec wants</b></div>{no}</div>')
    return page(body)


# =================================================================================================
# The addon panel
# =================================================================================================
SMALL_BTN = 'height:20px;padding:0 8px;font-size:12px'
EQUIP = f'<span class="sbtn" style="{SMALL_BTN}">Equip</span>'


def upgrade_row(icon, name_html, gain, source, bag=False):
    lead = f'<span style="color:{C["green"]}">{arrow()}In your bags</span> · ' if bag else ''
    return (f'<div class="irow">{img(icon, 30)}<div class="tx"><div class="n"><span>{name_html}</span>'
            f'<span style="color:{C["green"]}">+{gain} DPS</span></div><div class="s">{lead}{source}</div></div>'
            f'{EQUIP if bag else ""}</div>')


def spell_row(icon, name, cond, dim=False):
    return (f'<div class="irow {"dim" if dim else ""}" style="margin-bottom:5px">{img(icon, 22)}<div class="tx">'
            f'<div class="n" style="font-size:13px"><span>{name}</span></div>'
            f'<div class="s">{cond}</div></div></div>')


def panel(name, fact, sections, width=340, score=''):
    """Panel.lua on Widgets.umberFrame and titleBand: the medallion, the name in the title face over the
    muted fact line, the Forever Score on the right in its tier colour, then the sections."""
    body = '<div class="rule"></div>'.join(sections)
    return (f'<div class="plate" style="width:{width}px"><div class="band">{medallion()}'
            f'<div style="flex:1;min-width:0"><div style="display:flex;justify-content:space-between;'
            f'align-items:baseline;gap:8px"><div class="t">{name}</div>{score}</div><div class="m">{fact}</div></div>'
            f'</div><div class="x">×</div><div class="body">{body}</div></div>')


def score_fact(score, tier):
    return (f'<div style="white-space:nowrap"><span class="num" style="color:{TIER[tier]}">{score:,}</span>'
            f'<span class="unit"> score</span></div>')


def leveller_panel():
    upgrades = ('<div class="sect">Upgrades</div>'
                + upgrade_row('inv_weapon_bow_11', iname(3021), '3.5', 'Ranged · World drop')
                + upgrade_row('inv_belt_26', iname(10403), '1.1', 'Waist · The Deadmines · Captain Greenskin · 30%')
                + '<div class="small">Your other slots are best for your level.</div>')
    talents = (f'<div class="sect">Talents<span class="r">12 of 15 on plan · <span style="color:{ADDON['warning']}">3 off plan'
               f'</span></span></div><div class="irow">{img("ability_searingarrow", 30)}<div class="tx"><div class="n">'
               f'<span>Level 25: Lethal Attacks</span><span class="small">1 of 5</span></div>'
               f'<div class="s">Beast Mastery leveling · <span style="color:{C["gold"]}">Change</span></div></div>'
               f'<span class="sbtn" style="{SMALL_BTN}">Talents</span></div>')
    rotation = ('<div class="sect">Rotation at 24<span class="r">hover for the full rule</span></div>'
                + spell_row('ability_hunter_snipershot', "Hunter's Mark", 'Whenever the target is not marked')
                + spell_row('inv_spear_07', 'Aimed Shot', 'When its cooldown, shared with Multi-Shot, is up')
                + spell_row('ability_upgrademoonglaive', 'Multi-Shot', 'On cooldown, above a quarter mana')
                + spell_row('ability_hunter_quickshot', 'Serpent Sting', 'Under 1.5 seconds left')
                + spell_row('ability_impalingbolt', 'Arcane Shot', 'To fill a gap')
                + spell_row('ability_hunter_runningshot', 'Rapid Fire', 'At level 26', dim=True))
    return panel('Bow Jackzon', 'Level 24 Hunter · Employed Millennials', [upgrades, talents, rotation],
                 score=score_fact(1240, 'Rare'))


def board_panel_leveller():
    p = leveller_panel()
    body = (f'<p class="cap"><b>Levelling</b> · Bow Jackzon, level 24 Beast Mastery, guild without a planned raid; '
            f'the rotation is Data.lua\'s band 22, every line</p><div style="padding:16px 0 0 16px">{p}</div>')
    return page(body)


def answer_buttons():
    return ''.join(f'<span class="sbtn" style="{SMALL_BTN}">{a}</span>' for a in ('In', 'Late', 'Bench', 'Out'))


def tonight(answer_html, buttons, sub='as Ashvale, Damage'):
    row = (f'<div class="sect">Tonight<span class="r">8:00 pm · in 3 h 12 min</span></div>'
           f'<div class="irow">{img("inv_misc_head_dragon_black", 30)}<div class="tx"><div class="n">'
           f'<span>Onyxia\'s Lair</span>{answer_html}</div><div class="s">{sub}</div></div></div>')
    if buttons:
        row += f'<div style="display:flex;gap:6px;margin:0 0 8px 39px">{answer_buttons()}</div>'
    return (row + f'<div style="margin-left:39px;font-size:13px">Your job: <span style="color:{C["gold"]}">Phase 2 '
            f'ranged, west side</span></div><div class="small" style="margin-left:39px">with Thornhide</div>')


def board_panel_raider():
    upgrades = (f'<div class="sect">Upgrades<span class="r">your sim · 2 days ago · '
                f'<span class="sbtn" style="height:18px;padding:0 6px;font-size:11.5px">Sim again</span></span></div>'
                + upgrade_row('inv_weapon_bow_11', iname(13023), '6.4', 'Ranged · World drop')
                + upgrade_row('inv_gauntlets_11', iname(279257), '4.1', 'Hands · Leatherworking')
                + upgrade_row('inv_pants_03', iname(237819), '3.0', 'Legs · Illiyana Moonblaze · Vendor'))
    rotation = '<div class="sect" style="margin-bottom:0">Rotation at 60<span class="r">show</span></div>'
    fact, score = 'Level 60 Hunter · Employed Millennials', score_fact(1840, 'Epic')
    p1 = panel('Ashvale', fact, [tonight(f'<span style="color:{C["gold"]}">Not answered</span>', True),
                                 upgrades, rotation], score=score)
    pq = panel('Ashvale', fact, [tonight('<span style="color:#ffffff">In, not sent yet</span>', False,
                                         sub='as Ashvale, Damage<br>Sends when the companion app runs'),
                                 upgrades, rotation], score=score)
    p2 = panel('Ashvale', fact, [tonight(f'<span style="color:{C["green"]}">You\'re In</span>', False),
                                 upgrades, rotation], score=score)
    body = (f'<p class="cap"><b>Raid day</b> · a level 60 Beast Mastery hunter in a guild. Left: not answered. '
            f'Middle: In pressed, queued in the companion outbox, buttons gone. Right: the site confirmed it. '
            f'Gains at 60 are layout figures.</p>'
            f'<div class="row" style="padding:16px 0 0 16px;gap:36px">{p1}{pq}{p2}</div>')
    return page(body)


def board_panel_fresh():
    specs = [('ability_hunter_beasttaming', 'Beast Mastery', 'Your pet fights, Bestial Wrath makes it deadly'),
             ('ability_marksmanship', 'Marksmanship', 'Bows and guns, Barrage and Sniper Shot'),
             ('ability_hunter_swiftstrike', 'Survival', 'Melee hunter, traps and Lacerating Strikes')]
    path = ('<div class="sect">Pick your path<span class="r">you can change it later</span></div>'
            + ''.join(f'<div class="irow">{img(i, 30)}<div class="tx"><div class="n"><span>{n}</span></div>'
                      f'<div class="s" style="white-space:normal">{t}</div></div>'
                      f'<span class="sbtn" style="{SMALL_BTN}">Pick</span></div>'
                      for i, n, t in specs)
            + '<div class="small">Loads that path\'s leveling build. Your first talent point comes at level 10.</div>')
    rotation = ('<div class="sect">Rotation at 6</div>'
                + spell_row('ability_hunter_snipershot', "Hunter's Mark", 'Whenever the target is not marked')
                + spell_row('ability_hunter_quickshot', 'Serpent Sting', 'Under 1.5 seconds left')
                + spell_row('ability_impalingbolt', 'Arcane Shot', 'To fill a gap'))
    gear = ('<div class="sect" style="margin-bottom:2px">Upgrades</div>'
            '<div class="small">Gear picks start at level 20. Until then, the green arrow on your bags marks '
            'anything that beats what you wear.</div>')
    p = panel('Bow Jackzon', 'Level 6 Hunter', [path, rotation, gear])
    body = (f'<p class="cap"><b>Fresh install</b> · level 6, no build picked, no companion, no Data addon, no guild'
            f'</p><div style="padding:16px 0 0 16px">{p}</div>')
    return page(body)


# =================================================================================================
# First-run onboarding: six short pages on one loop (site, addon, companion, guild)
# =================================================================================================
PAGES = 5
OB_W, OB_H = 640, 450
MARK = b64((ROOT / 'design/logo/foreversixty-mark.svg').read_bytes(), 'image/svg+xml')
BANNER = b64((ROOT / 'design/mocks/data/icons/inv_bannerpvp_02.jpg').read_bytes(), 'image/jpeg')
ADDON_ICON = tga_uri(ROOT / 'addon/ForeverSixty/icon.png')
# Onboarding.lua coming(): the plain word in client gold, 6 px after the line.
COMING = f'<span style="font-size:12px;color:{C["gold"]};margin-left:6px">coming</span>'


# The tour wears the panel's chrome (design/specs/2026-11-08-addon-ux.md 4.2 Look; Widgets.umberFrame,
# Widgets.titleBand, Widgets.clientButton): the umber ground, the title band with the bronze-ring
# medallion and faction notch, the page heading in bone in the title face over a muted
# "Welcome to Forever Sixty", "2 of 5" on the right, the client's red X and red buttons.
# Marcellus stands in for the client's MORPHEUS title face here too (mixed case, like the build).
U = {'top': ADDON['raised'], 'ground': ADDON['raised'], 'edge': ADDON['line'], 'line': ADDON['line'],
     'bone': ADDON['cream'], 'bone_dim': ADDON['muted'], 'ink': ADDON['body'], 'muted': ADDON['muted'],
     'brass': ADDON['gold'], 'alliance': color('faction.allianceBar'), 'horde': color('faction.hordeBar')}
TITLE_FONT = b64((ROOT / 'apps/shared/fonts/Marcellus-Regular.ttf').read_bytes(), 'font/ttf')
RING = tga_uri(ROOT / 'addon/ForeverSixty/media/medallion-ring.tga')
NOTCH = tga_uri(ROOT / 'addon/ForeverSixty/media/medallion-notch.tga')
# Onboarding.ARROW_PARTS: { x, y, width, height } in a 16 x 10 box pointing right.
ARROW_PARTS = [(0, 4, 10, 2), (10, 0, 2, 10), (12, 2, 2, 6), (14, 4, 2, 2)]


def medallion(faction='alliance', size=44):
    inset = 6
    return (f'<div style="position:relative;width:{size}px;height:{size}px;flex:none">'
            f'<div style="position:absolute;inset:{inset}px;border-radius:50%;background:#000 url({CREST}) center/100%">'
            f'</div><img src="{RING}" style="position:absolute;inset:0;width:100%;height:100%">'
            f'<div style="position:absolute;inset:0;background:{U[faction]};-webkit-mask:url({NOTCH}) center/100%">'
            f'</div></div>')


def flow_arrow(direction):
    """Onboarding.arrowParts: the right arrow, mirrored for left, turned for down; brass rectangles."""
    parts = []
    for x, y, w, h in ARROW_PARTS:
        if direction == 'left':
            x = 16 - x - w
        elif direction == 'down':
            x, y, w, h = y, x, h, w
        parts.append(f'<div style="position:absolute;left:{x}px;top:{y}px;width:{w}px;height:{h}px;'
                     f'background:{U["brass"]}"></div>')
    box = (10, 16) if direction == 'down' else (16, 10)
    return (f'<div style="position:relative;width:{box[0]}px;height:{box[1]}px;flex:none;margin:12px 3px 0">'
            f'{"".join(parts)}</div>')


def onboarding(page, heading, content, next_label='Next', extra_right='', total=PAGES, faction='alliance'):
    left_extra = ('<span class="sbtn" style="padding:0 10px;position:relative">Tooltips only</span>'
                  if page == 1 else '')
    skip = '' if page == total else '<span class="sbtn" style="padding:0 14px">Skip tour</span>'
    back = '' if page == 1 else '<span class="sbtn">Back</span>'
    # Onboarding.frameHeight: the page is as tall as its content (24 over the footer, 14 under it),
    # 240 to 450.
    footer = (f'<div style="margin:24px 16px 0;padding-bottom:14px;display:flex;align-items:center;'
              f'justify-content:space-between"><span style="display:flex;gap:8px;align-items:center">'
              f'{skip}{left_extra}</span><span style="display:flex;gap:8px">{back}{extra_right}'
              f'<span class="sbtn primary" style="padding:0 20px">{next_label}</span></span></div>')
    band = (f'<div style="height:64px;display:flex;align-items:center;gap:12px;padding:0 42px 0 12px;'
            f'border-bottom:1px solid {U["line"]}">{medallion(faction)}'
            f'<div style="flex:1;min-width:0"><div style="font-family:Title,serif;font-size:19px;color:{U["bone"]};'
            f'text-shadow:1px 1px 0 #000;line-height:22px">{heading}</div>'
            f'<div style="font-size:12px;color:{U["muted"]};margin-top:3px">Welcome to Forever Sixty</div></div>'
            f'<div style="white-space:nowrap;align-self:flex-start;margin-top:12px"><span style="font-family:Site,'
            f'sans-serif;font-weight:600;font-size:17px;'
            f'color:{U["ink"]}">{page}</span><span style="font-size:12px;color:{U["bone_dim"]}"> of {total}'
            f'</span></div></div>')
    return (f'<style>@font-face {{ font-family: Title; src: url({TITLE_FONT}); }}</style>'
            f'<div style="position:relative;width:{OB_W}px;min-height:240px;max-height:{OB_H}px;'
            f'box-sizing:border-box;color:{U["ink"]};font-family:Site,sans-serif;'
            f'background:{U["ground"]};border:1px solid {U["edge"]};box-shadow:0 8px 28px rgba(0,0,0,.7)">'
            f'{band}<div style="position:absolute;top:6px;right:8px;width:20px;height:20px;color:{U["ink"]};'
            f'border:1px solid {U["line"]};box-sizing:border-box;'
            f'font-size:14px;display:flex;align-items:center;justify-content:center">×</div>'
            f'<div style="padding:16px 26px 0 26px">{content}</div>{footer}</div>')


def ob_board(caption, frame):
    return page(f'<p class="cap">{caption}</p><div style="padding:18px 0 0 18px">{frame}</div>')


def node_icon(uri, size=40, round_=False):
    radius = '50%' if round_ else '4px'
    return (f'<div style="width:{size}px;height:{size}px;border-radius:{radius};background:#000 url({uri}) center/100%;'
            f'box-shadow:0 0 0 1px #000,0 0 0 2px {U["edge"]};flex:none"></div>')


def loop_cell(icon, title, text):
    return (f'<div style="width:172px;display:flex;gap:9px;align-items:flex-start">{icon}'
            f'<div><div style="color:{U["ink"]};font-size:13.5px;line-height:16px">{title}</div>'
            f'<div style="font-size:11.5px;color:{U["muted"]};line-height:14px">{text}</div></div></div>')


def board_onboarding_1(guild=True):
    top = [loop_cell(node_icon(CREST, 36, round_=True), 'Your character', 'Imported from Battle.net when you sign in.'),
           loop_cell(node_icon(ADDON_ICON, 36), 'foreversixty.gg', 'Best in slot for your level, the sim on your gear, '
                     'guides, tiers.'),
           loop_cell(node_icon(icon_uri('inv_misc_bag_08'), 36), 'In game', 'The arrow on your bags, the talent '
                     'ring, your panel.')]
    guild_cell = loop_cell(node_icon(BANNER, 36), 'Your guild', 'Raids from the site, in your Discord and in game.')
    bottom = [guild_cell if guild else '<div style="width:172px"></div>',
              loop_cell(node_icon(icon_uri('inv_misc_book_09'), 36), 'Back on the site', 'Your fights rated, the '
                        'coach, your Forever Score.'),
              loop_cell(node_icon(icon_uri('ability_warrior_battleshout'), 36), 'Your fights', 'The companion app '
                        'uploads each one.')]
    row1 = f'<div style="display:flex;align-items:flex-start;gap:6px">{top[0]}{flow_arrow("right")}{top[1]}' \
           f'{flow_arrow("right")}{top[2]}</div>'
    down = (f'<div style="display:flex;justify-content:flex-end;padding-right:150px;margin:6px 0">'
            f'{flow_arrow("down")}</div>')
    first_arrow = flow_arrow('left') if guild else '<div style="width:22px"></div>'
    row2 = f'<div style="display:flex;align-items:flex-start;gap:6px">{bottom[0]}{first_arrow}{bottom[1]}' \
           f'{flow_arrow("left")}{bottom[2]}</div>'
    diagram = f'<div style="width:586px;padding-top:6px">{row1}{down}{row2}</div>'
    return ob_board('<b>Onboarding, page 1 of 5</b> (4 for a character with no guild) · how the pieces connect. '
                    'Skip tour (and Escape) ends it for good; Tooltips only turns the setting on and ends it',
                    onboarding(1, 'How Forever Sixty works', diagram, total=PAGES if guild else PAGES - 1))


def board_onboarding_1_hover():
    return board_onboarding_1().replace(
        'Onboarding, page 1 of 5</b>', 'Onboarding, page 1 of 5, hovering Tooltips only</b>').replace(
        'Tooltips only</span>', 'Tooltips only' + HOVER_HINT + '</span>', 1)


HOVER_HINT = ('<div class="tt" style="position:absolute;left:-30px;bottom:30px;min-width:0;width:300px;'
              'font-size:12.5px;line-height:16px;text-align:left;color:#fff"><div style="color:#ffd100;font-size:14px;'
              'line-height:18px">Tooltips only</div><div>Turns off everything but the tooltip lines.</div>'
              '<div>Ends the tour. /fs help opens it again.</div></div>')


def board_onboarding_1_solo():
    return board_onboarding_1(guild=False).replace(
        'Onboarding, page 1 of 5</b> (4 for a character with no guild)',
        'Onboarding, page 1 of 4, a character with no guild</b> (no guild node, no guild clause)')


def icon_row(icon, title, text, source=None, extra=''):
    """drawRows in Onboarding.lua: an optional muted source right-aligned in 160 px, a 30 px icon, the
    title in ink and its line muted under it."""
    src = (f'<div style="width:160px;flex:none;text-align:right;font-size:12px;color:{U["muted"]};'
           f'line-height:15px;margin-right:16px">{source}</div>' if source else '')
    return (f'<div style="display:flex;align-items:flex-start;margin-bottom:8px;min-height:44px">{src}'
            f'{node_icon(icon, 30)}<div style="margin-left:10px"><div style="color:{U["ink"]};font-size:13.5px;'
            f'line-height:17px">{title}</div><div style="font-size:12px;color:{U["muted"]};line-height:15px;'
            f'margin-top:1px">{text}</div>{extra}</div></div>')


ARROW_ICON = tga_uri(ROOT / 'addon/ForeverSixty/media/arrow.tga')


def tooltip_card(name, quality, lines):
    """Onboarding.lua exampleCard: the client's tooltip ground (#080a16) and edge (#9095a3), the item's
    name in its quality colour, then the lines the addon adds."""
    body = ''.join(f'<div style="color:{color};font-size:12px;line-height:15px">{text}</div>'
                   for text, color in lines)
    return (f'<div style="display:inline-block;margin-top:6px;padding:7px 10px 7px;background:rgba(8,10,22,.94);'
            f'border:1px solid #9095a3"><div style="color:{QUALITY[quality]};font-size:13.5px;line-height:18px">'
            f'{name}</div>{body}</div>')


def example_card():
    """TourExample for Bow Jackzon (level 6, no build, no upgrade): the Beast Mastery band-20 main hand
    from the shipped BiS data, never a hand-typed item, so a nightly refresh redraws it."""
    pick = next(s for s in BAND20['slots'] if s['slot'] == 'main_hand')
    place, boss = pick['source'].split(': ', 1)
    source = f'{place} · {boss} · {round(pick["chance"] * 100)}% drop' if pick.get('chance') else pick['source']
    return tooltip_card(pick['item_name'], ITEMS[pick['item_id']]['quality'],
                        [(f'Best in slot at level {BAND20["band"]}', C['gold']), (source, '#b8b8b8')])


def board_onboarding_2():
    rows = (icon_row(ARROW_ICON, 'The green arrow', 'On a bag item, a quest reward or a loot roll that beats what '
                     'you wear.', 'BiS pages')
            # Level 6 Beast Mastery, alliance, no upgrade yet: Data.lua bis band 20 main_hand, item 5200.
            + icon_row(icon_uri('inv_weapon_bow_02'), 'One line on every item', '+2.2 DPS, Best in slot or Not for '
                       'you. Shift shows the rest.', 'Spec weights or your sim',
                       example_card())
            + icon_row(icon_uri('ability_searingarrow'), 'The gold talent ring', "Your next point, in that build's "
                       'order.', 'Class guide build')
            + icon_row(icon_uri('inv_spear_07'), 'Rotation in your panel', 'The buttons to press at your level, and '
                       'what changes next.', 'Class guide rotation'))
    return ob_board('<b>Page 2</b> · what the site gives you in game; each row names where it comes from',
                    onboarding(2, 'The site, in your game', rows))


def board_onboarding_3():
    rows = (icon_row(icon_uri('ability_warrior_battleshout'), 'The companion uploads your fights',
                     'The companion reads your combat log and sends each fight as it ends. Turn on Advanced Combat '
                     'Logging once (System, Network), then type /combatlog when you play.')
            + icon_row(icon_uri('inv_misc_book_09'), 'The site rates them and coaches you',
                       'Each fight gets a report, a percentile and the one thing that cost you most, like '
                       '"biggest loss: standing in fire".')
            + icon_row(ADDON_ICON, 'Your Forever Score', 'Gear, experience, parses, attendance and preparation, in '
                       'one number, on your tooltip for players with Forever Sixty and Forever Sixty Data. '
                       'Settings can hide it from others.'))
    return ob_board('<b>Page 3</b> · what the game gives back', onboarding(3, 'Your game, back to the site', rows))


def answer_place(name, live):
    tag = '' if live else COMING
    color = U['ink'] if live else U['muted']
    return f'<div style="font-size:13px;color:{color};margin-bottom:5px">{name}{tag}</div>'


def board_onboarding_4():
    chip_line = f'<div style="font-size:13.5px;color:{U["ink"]}">Tonight 8:00 pm: Onyxia\'s Lair</div>'
    content = (f'<div style="font-size:13.5px;color:{U["ink"]};margin:0 0 14px">Your officers post a raid on the site '
               f'(and in your Discord, if your guild has linked it). On raid day, it is here:</div>'
               f'<div style="display:flex;gap:28px;align-items:flex-start">'
               f'<div style="width:300px">{chip_line}<div style="font-size:12px;color:{U["muted"]};margin:8px 0 14px">The line by your '
               f'minimap, on raid day.</div>'
               f'<div style="font-size:13px">Job: <span style="color:{C["gold"]}">Whelps, west cave</span>'
               f'</div><div style="font-size:12px;color:{U["muted"]};margin-top:2px">In your panel, and in chat when you zone in.</div></div>'
               f'<div><div style="color:{U["ink"]};font-size:14px;margin-bottom:6px">Answer In, Late, Bench or Out</div>'
               f'{answer_place("On foreversixty.gg", True)}{answer_place("In Discord, on the raid post", True)}'
               f'<div style="font-size:12px;color:{U["muted"]};margin:2px 0 0">In game and on the phone app{COMING}</div>'
               f'</div></div>')
    return ob_board('<b>Page 4, only for a character in a guild</b> · your guild, everywhere. One muted "coming" '
                    'line for what is not live (in-game answers need the sign-up contract; the apps are not in the '
                    'stores yet)', onboarding(4, 'Your guild, everywhere', content))


def url_box(url):
    return (f'<div style="height:22px;padding:0 9px;display:inline-flex;align-items:center;font-size:13px;color:{U["ink"]};'
            f'background:{ADDON["inset"]};border:1px solid {U["edge"]};min-width:200px">{url}</div>')


APP_URL = 'https://foreversixty.gg/app'


QR_ASSET = ROOT / 'design/assets/app-qr.png'


def qr_uri():
    """The phone-app QR: the web lane's design/assets/app-qr.png once it exists; until then segno at
    error correction H with a 4-module quiet zone, black on white (the same encoding)."""
    if QR_ASSET.exists():
        return b64(QR_ASSET.read_bytes(), 'image/png')
    import segno
    buf = io.BytesIO()
    segno.make(APP_URL, error='h').save(buf, kind='svg', scale=4, border=4, dark='#000', light='#fff')
    return b64(buf.getvalue(), 'image/svg+xml')


def qr_block():
    """Onboarding.QR_SIZE (48) while the apps are in testing, labelled Join the app beta."""
    return (f'<div style="width:120px;flex:none;text-align:center;padding-top:4px">'
            f'<img src="{qr_uri()}" style="width:48px;height:48px">'
            f'<div style="font-size:12px;color:{U["ink"]};margin-top:4px">Join the app beta</div>'
            f'<div style="font-size:12px;color:{U["muted"]};margin-top:1px">foreversixty.gg/app</div></div>')


SPECS = [('Beast Mastery', 'Your pet fights, Bestial Wrath'),
         ('Marksmanship', 'Bows and guns, Barrage and Sniper Shot'),
         ('Survival', 'Melee hunter, traps and Lacerating Strikes')]


def path_tiles(build_word='leveling'):
    """drawSetup with no build: Pick your path, the note, one red client button per spec (188 wide, 196
    apart) with its playstyle muted under it."""
    tiles = ''.join(f'<div style="width:188px"><span class="sbtn" style="width:188px">{name}</span>'
                    f'<div style="font-size:12px;color:{U["muted"]};line-height:15px;margin-top:6px">{text}</div>'
                    f'</div>' for name, text in SPECS)
    return (f'<div style="font-size:13.5px;color:{U["ink"]};margin-bottom:6px">Pick your path '
            f'<span style="font-size:12px;color:{U["muted"]};margin-left:6px">loads its {build_word} build</span></div>'
            f'<div style="display:flex;gap:8px;margin-bottom:18px">{tiles}</div>')


def connect_rows(seen=None):
    if seen is not None:
        return (f'<div style="font-size:13.5px;color:{C["green"]};margin-bottom:2px">The companion has run</div>'
                f'<div style="font-size:12px;color:{U["muted"]};margin-bottom:10px">The site saw Bow Jackzon {seen}.'
                f'</div>')
    rows = [
        ('Sign in on the site', 'Sign in with Battle.net. Upgrades, talents and rotation work without it.',
         url_box('foreversixty.gg')),
        ('Get the companion for your PC', 'Uploads your fights for reports and the coach, and keeps your builds in '
         'step with the site.', url_box('foreversixty.gg/setup')),
    ]
    return ''.join(f'<div style="margin-bottom:10px"><div style="color:{U["ink"]};font-size:13.5px">{t}</div>'
                   f'<div style="font-size:12px;color:{U["muted"]};margin:1px 0 4px">{d}</div>{u}</div>'
                   for t, d, u in rows)


LINKS_LINE = f'<div style="font-size:12px;color:{U["muted"]}">Click a box, then Ctrl+C.</div>'
DONE_LINE = f'<div style="font-size:12px;color:{U["muted"]};margin-top:4px">/fs help opens this again.</div>'


def setup_page(top, connect, page_no, total, boxes=True):
    content = (f'{top}<div style="display:flex;gap:20px"><div style="flex:1">{connect}'
               f'{LINKS_LINE if boxes else ""}{DONE_LINE}</div>{qr_block()}</div>')
    return onboarding(page_no, 'Set it up', content, next_label='Done', total=total,
                      extra_right='<span class="sbtn">Open my panel</span>')


def board_onboarding_5():
    return ob_board('<b>Last page</b> · set it up. Level 6, no build, no guild, so this is page 4 of 4: one button '
                    'per spec with its playstyle, nothing pre-picked, and Done closes the tour',
                    setup_page(path_tiles(), connect_rows(), 4, 4))


def board_onboarding_5_connected():
    top = (f'<div style="display:flex;justify-content:space-between;align-items:flex-start;margin-bottom:24px">'
           f'<div><span style="font-size:13.5px;color:{U["ink"]}">Your path: Beast Mastery leveling</span>'
           f'<span style="font-size:12px;color:{U["muted"]};margin-left:10px">15 of 15 points</span>'
           f'<div style="font-size:12px;color:{U["muted"]};margin-top:3px">Next point at level 25: Lethal Attacks'
           f'</div></div><span class="sbtn">Change</span></div>')
    return ob_board('<b>Last page, a character already set up</b> · level 24, in a guild (page 5 of 5), following a '
                    'build, with an ack in the inbox. Past 3 days without one, the line turns to the warning colour: '
                    '"The site has not seen Bow Jackzon in 4 days. Is the companion running?"',
                    setup_page(top, connect_rows('2 hours ago'), 5, 5, boxes=False))


def board_onboarding_5_sixty():
    return ob_board('<b>Last page at level 60, no build yet</b> · the same buttons load the raid build',
                    setup_page(path_tiles(build_word='raid'), connect_rows(), 5, 5))


MINIMAP_ICON = tga_uri(ROOT / 'addon/ForeverSixty/media/minimap.tga')


def board_minimap_button():
    """Minimap.lua: the flat seal (media/minimap, 32 px texture drawn at Theme.SIZES.minimapIcon 20) in the
    client's own round tracking border on the minimap's rim, with the chip beside the minimap."""
    button = (f'<div style="position:absolute;left:-8px;top:96px;width:32px;height:32px;border-radius:50%;'
              f'background:#000;box-shadow:0 0 0 3px #8b6a35,0 0 0 4px #2d220f,0 2px 6px #000;display:flex;'
              f'align-items:center;justify-content:center"><img src="{MINIMAP_ICON}" style="width:20px;height:20px">'
              f'</div>')
    mm = minimap(chip('Next talent at 25: Lethal Attacks'), 180).replace(
        'box-shadow:0 0 0 5px #8b6a35,0 0 0 7px #2d220f,0 4px 12px #000"></div>',
        'box-shadow:0 0 0 5px #8b6a35,0 0 0 7px #2d220f,0 4px 12px #000">' + button + '</div>', 1)
    return page('<p class="cap"><b>Minimap button and chip</b> · the flat seal at 20 px on the minimap rim; the chip '
                f'on the window ground in one hairline</p><div style="padding:8px 0 0 8px">{mm}</div>')


def board_button_states():
    """Widgets.clientButton and Widgets.closeButton in each state (Theme.CONTROL_STATES.siteButton and
    .close): the secondary at rest and under the cursor (its ground takes its edge colour), the selected
    primary likewise, a held button's label one pixel down, and the close X at rest and hovered."""
    def btn(label, ground, edge, drop=False):
        return (f'<span class="sbtn" style="background:{ground};border-color:{edge};'
                f'{"padding-top:2px" if drop else ""}">{label}</span>')

    def x(hover):
        ground = ADDON['hover'] if hover else 'transparent'
        ink = ADDON['strong'] if hover else ADDON['body']
        return (f'<span style="display:inline-flex;width:20px;height:20px;align-items:center;justify-content:center;'
                f'border:1px solid {ADDON["line"]};background:{ground};color:{ink};font-size:14px">×</span>')

    def cell(caption, html):
        return f'<div class="col" style="gap:6px;align-items:flex-start"><div class="note">{caption}</div>{html}</div>'

    cells = [
        cell('Secondary, rest', btn('Late', ADDON['blue'], ADDON['blue_edge'])),
        cell('Secondary, hover', btn('Late', ADDON['blue_edge'], ADDON['blue_edge'])),
        cell('Secondary, held', btn('Late', ADDON['blue_edge'], ADDON['blue_edge'], drop=True)),
        cell('Primary or selected, rest', btn('Next', ADDON['teal'], ADDON['teal_edge'])),
        cell('Primary or selected, hover', btn('Next', ADDON['teal_edge'], ADDON['teal_edge'])),
        cell('Close, rest', x(False)),
        cell('Close, hover', x(True)),
    ]
    return page('<p class="cap"><b>Button states</b> · the site\'s recipes on the addon\'s plate '
                '(Theme.CONTROL_STATES.siteButton, .close)</p>'
                f'<div class="plate" style="padding:18px 20px;display:flex;gap:26px">{"".join(cells)}</div>')


BOARDS = {
    'button-states': board_button_states,
    'minimap-button': board_minimap_button,
    'item-tooltip-before': board_item_before,
    'item-tooltip-after': board_item_after,
    'item-tooltip-after-shift': board_item_after_shift,
    'item-tooltip-states': board_item_states,
    'quest-reward-before': board_quest_before,
    'quest-reward-after': board_quest_after,
    'loot-roll-after': board_loot_roll,
    'player-tooltip-before': board_player_before,
    'player-tooltip-after': board_player_after,
    'chip-pulse': board_chip,
    'panel-leveller-24': board_panel_leveller,
    'panel-raider-60': board_panel_raider,
    'panel-fresh-install': board_panel_fresh,
    'onboarding-1-loop': board_onboarding_1,
    'onboarding-1-hover': board_onboarding_1_hover,
    'onboarding-1-loop-solo': board_onboarding_1_solo,
    'onboarding-2-site-in-game': board_onboarding_2,
    'onboarding-3-game-to-site': board_onboarding_3,
    'onboarding-4-guild': board_onboarding_4,
    'onboarding-5-set-up': board_onboarding_5,
    'onboarding-5-set-up-done': board_onboarding_5_connected,
    'onboarding-5-set-up-60': board_onboarding_5_sixty,
}


def main(names):
    from playwright.sync_api import sync_playwright
    OUT.mkdir(parents=True, exist_ok=True)
    with sync_playwright() as pw:
        # PLAYWRIGHT_CHROMIUM: a headless shell when the pip package's own build is not downloaded.
        exe = os.environ.get('PLAYWRIGHT_CHROMIUM')
        browser = pw.chromium.launch(executable_path=exe) if exe else pw.chromium.launch()
        ctx = browser.new_context(device_scale_factor=2, viewport={'width': 1800, 'height': 1000})
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
