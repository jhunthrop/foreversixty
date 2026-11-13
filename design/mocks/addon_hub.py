"""Acceptance boards for the Forever Sixty window, the addon's one hub (design/specs/2026-11-13-addon-hub.md).

    python3 design/mocks/addon_hub.py              # every board
    python3 design/mocks/addon_hub.py hub-me hub-chip

Writes design/reviews/2026-11-13-addon-hub/<board>.png at device scale 2, the window at its real 740 px width
(Theme.SIZES: padding 16, windowLeft 388, windowGutter 32, windowRight 288; Hub.LAYOUT: header 56, strip 28,
body padding 12, footer 28), the same frame the builds-window lane draws for the Builds tab.

Every character, guild and raid group here is made up: Velisandre, a level 60 Night Elf hunter in the guild
Thornwake Company, and her raid team. No real player, guild or realm.

Sources, so nothing else is invented:
- Colours and faces: design/tokens/tokens.json through addon_tokens (the addon's Tokens.lua), parse and
  rarity colours included.
- Gear: data/builds/1.60.1.70790/bis/hunter-beast-mastery.json band 60 raid (alliance): every pick, its
  source and kind; Velisandre's worn items are that list's own alternatives and each gain is the list's
  dps_delta for that alternative. Item names, qualities and icons from items/hunter.json and icons/.
- Source counts: the slots with a pick or an alternative of that kind in the same band (Sources.counts).
- Stat weights: the same band's weights (scale factor, DPS per point, error), the SimC convention.
- Rotation: addon/ForeverSixty/Data.lua rotations["hunter-beast-mastery"] level 60, read from the file;
  spell icons from spellicons.json.
- Forever Score: the five parts and their maxima from docs/contracts/fixtures/score-labels.json and
  web/src/lib/score/score.ts (COMPONENT_MAX, TIER_FROM); Velisandre's figures are made up inside them.
- Inbox shapes (last raid, raid assignment, raid sign-up, sim result): docs/contracts/fixtures/inbox-*.json.
- Stat weights: DPS per point normalized to the top stat (the SimC convention), haste per 1% outside it.
- Rep standings on sources: the band's slot `standing`.
- Sign-up words (In, Late, Bench, Out, Not signed up, as DPS) and role names: web/src/lib/guild/raid-copy.ts.
- Score words (`N to Epic (1500)`, `410 of 500`): web/src/lib/score/copy.ts.
- Master loot items: Magmadar's Striker's Mark, Giantstalker's Leggings, Aged Core Leather Gloves (items/hunter.json).
- Dates: Thu 10 Dec 2026 is raid day; the last raid is Tue 8 Dec.
"""
import json
import os
import re
import sys

import addon_ux as ux
from addon_tokens import color
from addon_ux import ADDON, ROOT, b64, icon_uri, page, tga_uri

OUT = ROOT / 'design/reviews/2026-11-13-addon-hub'
BUILD = ROOT / 'data/builds/1.60.1.70790'
W, PAD, LEFT, GUTTER, RIGHT = 740, 16, 388, 32, 288
HEADER, STRIP, FOOTER, BODY_PAD = 56, 28, 28, 12
CLIENT_GOLD = '#ffd100'

ITEMS = {it['id']: it for it in json.loads((BUILD / 'items/hunter.json').read_text())['items']}
BIS = json.loads((BUILD / 'bis/hunter-beast-mastery.json').read_text())
RAID60 = next(b for b in BIS['bands'] if b['band'] == 60 and b['preset'] == 'raid' and b['faction'] == 'alliance')
SPELL_ICONS = json.loads((BUILD / 'spellicons.json').read_text())
DATA_LUA = (ROOT / 'addon/ForeverSixty/Data.lua').read_text()
SCORE_LABELS = json.loads((ROOT / 'docs/contracts/fixtures/score-labels.json').read_text())
COMPONENT_MAX = {'gear': 500, 'experience': 700, 'parses': 600, 'attendance': 200, 'discipline': 100}
TIER_FROM = [('legendary', 2000), ('epic', 1500), ('rare', 1000), ('uncommon', 500), ('common', 0)]

QUALITY = {1: color('rarity.common'), 2: color('rarity.uncommon'), 3: color('rarity.rareText'),
           4: color('rarity.epicText'), 5: color('rarity.legendary')}
TIER_HEX = {'common': QUALITY[1], 'uncommon': QUALITY[2], 'rare': QUALITY[3], 'epic': QUALITY[4],
            'legendary': QUALITY[5]}
PARSE = [(100, color('parse.gold')), (99, color('parse.pinkText')), (95, color('parse.orange')),
         (75, color('parse.purple')), (50, color('parse.blue')), (25, color('parse.green')),
         (0, color('parse.greyText'))]
CLASS_HEX = {c: color(f'class.{c}') for c in
             ('warrior', 'paladin', 'hunter', 'rogue', 'priest', 'shaman', 'mage', 'warlock', 'druid')}
TABS = ['Me', 'Builds', 'Gear', 'Rotation', 'Guild', 'Officer']
SLOT_NAMES = {'head': 'Head', 'neck': 'Neck', 'shoulder': 'Shoulder', 'back': 'Back', 'chest': 'Chest',
              'wrist': 'Wrist', 'hands': 'Hands', 'waist': 'Waist', 'legs': 'Legs', 'feet': 'Feet',
              'finger1': 'Finger', 'finger2': 'Finger', 'trinket1': 'Trinket', 'trinket2': 'Trinket',
              'main_hand': 'Main hand', 'off_hand': 'Off hand', 'ranged': 'Ranged'}
KINDS = [('dungeon', 'Dungeon'), ('raid', 'Raid'), ('world_drop', 'World / BoE'), ('quest', 'Quest'),
         ('crafted', 'Crafted'), ('rep', 'Reputation'), ('vendor', 'Vendor'), ('pvp', 'PvP')]

# Velisandre's worn items: the pick where she has it, else that slot's alternative from the same list.
WORN = {'neck': 22403, 'back': 22330, 'wrist': 13211, 'hands': 13255, 'feet': 14616, 'finger2': 13098,
        'main_hand': 13052}


def esc(text):
    return str(text).replace('&', '&amp;').replace('<', '&lt;').replace('>', '&gt;')


def parse_hex(pct):
    return next(hex_ for floor, hex_ in PARSE if pct >= floor)


def tier_of(score):
    return next(name for name, floor in TIER_FROM if score >= floor)


def crest(cls, size):
    ring = tga_uri(ROOT / 'addon/ForeverSixty/media/crests/ring.tga')
    face = tga_uri(ROOT / f'addon/ForeverSixty/media/crests/{cls}.tga')
    return (f'<span class="crest" style="width:{size}px;height:{size}px"><img src="{face}">'
            f'<span style="position:absolute;inset:0;background:{CLASS_HEX[cls]};-webkit-mask:url({ring}) '
            f'center/100% no-repeat"></span></span>')


def icon(name, size, extra=''):
    return f'<img class="ic" src="{icon_uri(name)}" style="width:{size}px;height:{size}px;{extra}">'


def item_icon(item_id, size):
    it = ITEMS[item_id]
    return icon(it['icon'], size, f'box-shadow:0 0 0 1px #000,0 0 0 2px {QUALITY[it["quality"]]}55')


def item_name(item_id):
    it = ITEMS[item_id]
    return f'<span style="color:{QUALITY[it["quality"]]}">{esc(it["name"])}</span>'


# --- gear from the band's BiS list ------------------------------------------------------------------

def slot_rows():
    rows = {}
    for s in RAID60['slots']:
        if s.get('item_id') is None:
            continue
        worn = WORN.get(s['slot'], s['item_id'])
        alt = next((a for a in s.get('alternatives', []) if a['item_id'] == worn), None)
        gain = -alt['dps_delta'] if alt else 0.0
        rows[s['slot']] = {'slot': s['slot'], 'pick': s['item_id'], 'worn': worn, 'gain': gain,
                           'source': s.get('source'), 'kind': s.get('source_kind')}
    return rows


def kind_counts():
    counts = {k: 0 for k, _ in KINDS}
    for s in RAID60['slots']:
        kinds = {s.get('source_kind')} | {a.get('source_kind') for a in s.get('alternatives', [])}
        for k in counts:
            counts[k] += k in kinds
    return counts


def short_source(source, kind):
    if kind == 'dungeon' and ':' in source:
        return source.split(':')[0]
    return source


# --- the rotation from Data.lua -----------------------------------------------------------------------

def rotation_lines(spec, level):
    table = re.search(r'^local rotationLines = \{\n(.*?)^\}', DATA_LUA, re.M | re.S).group(1)
    lines = [m.groups() for m in re.finditer(
        r'\{ spellId = (\d+), name = "([^"]*)", condition = "([^"]*)"', table)]
    block = re.search(rf'\["{re.escape(spec)}"\] = \{{\n(.*?)\n\t\t\}},', DATA_LUA[DATA_LUA.index('rotations = {'):],
                      re.S).group(1)
    row = re.search(rf'\{{ level = {level}, lines = \{{ ([^}}]*) \}} \}}', block).group(1)
    return [lines[int(i) - 1] for i in re.findall(r'rotationLines\[(\d+)\]', row)]


# --- the frame ------------------------------------------------------------------------------------------

CSS = f'''
.cap {{ max-width: 980px; }}
.win {{ position: relative; width: {W}px; background: {ADDON['raised']}; border: 1px solid {ADDON['line']};
  box-shadow: 0 8px 28px rgba(0,0,0,.7); font-family: Site, sans-serif; font-size: 12px; line-height: 15px;
  color: {ADDON['body']}; }}
.hd {{ height: {HEADER}px; display: flex; align-items: center; gap: 10px; padding: 0 {PAD}px;
  border-bottom: 1px solid {ADDON['line']}; }}
.hd .htt {{ font-family: Title, serif; font-size: 16px; line-height: 20px; color: {ADDON['cream']}; }}
.hd .nt {{ color: {ADDON['muted']}; white-space: nowrap; }}
.crest {{ position: relative; display: block; flex: none; }}
.crest img {{ position: absolute; inset: 0; width: 100%; height: 100%; }}
.x {{ position: absolute; top: 18px; right: 12px; width: 20px; height: 20px; border: 1px solid {ADDON['line']};
  display: flex; align-items: center; justify-content: center; font-size: 14px; }}
.strip {{ height: {STRIP}px; display: flex; gap: 4px; padding: 0 {PAD}px; border-bottom: 1px solid {ADDON['line']}; }}
.strip .tab {{ min-width: 96px; padding: 0 12px; display: flex; align-items: center; justify-content: center;
  color: {ADDON['muted']}; border-bottom: 2px solid transparent; margin-bottom: -1px; white-space: nowrap; }}
.strip .tab.on {{ color: {ADDON['gold']}; border-color: {ADDON['gold']}; }}
.strip .tab .dot {{ width: 6px; height: 6px; border-radius: 50%; background: {CLIENT_GOLD}; margin-left: 6px; }}
.bd {{ display: flex; padding: {BODY_PAD}px {PAD}px; }}
.bd.one {{ display: block; }}
.lc {{ width: {LEFT}px; flex: none; }}
.vr {{ width: 1px; background: {ADDON['line']}; margin: -{BODY_PAD}px {GUTTER // 2 - 1}px -{BODY_PAD}px {GUTTER // 2}px; flex: none; }}
.rc {{ width: {RIGHT}px; flex: none; }}
.ft {{ height: {FOOTER}px; border-top: 1px solid {ADDON['line']}; display: flex; align-items: center;
  justify-content: space-between; padding: 0 {PAD}px; color: {ADDON['muted']}; }}
.ft span {{ opacity: .75; }}
.sec {{ margin-bottom: 20px; }}
.sec:last-child {{ margin-bottom: 0; }}
.sh {{ height: 18px; display: flex; align-items: center; justify-content: space-between; gap: 8px;
  font-weight: 600; color: {ADDON['strong']}; margin-bottom: 4px; }}
.sh .r {{ font-weight: 400; color: {ADDON['muted']}; white-space: nowrap; overflow: hidden; text-overflow: ellipsis; }}
.m {{ color: {ADDON['muted']}; }}
.s {{ color: {ADDON['strong']}; }}
.g {{ color: {ADDON['gold']}; }}
.cg {{ color: {CLIENT_GOLD}; }}
.ok {{ color: {ADDON['success']}; }}
.wa {{ color: {ADDON['warning']}; }}
.one {{ white-space: nowrap; overflow: hidden; text-overflow: ellipsis; }}
.ic {{ box-shadow: 0 0 0 1px #000; flex: none; display: block; }}
.row {{ display: flex; align-items: center; gap: 8px; }}
.btn {{ display: inline-flex; align-items: center; height: 20px; padding: 0 8px; font-weight: 600; flex: none;
  color: {ADDON['strong']}; background: {ADDON['blue']}; border: 1px solid {ADDON['blue_edge']}; white-space: nowrap; }}
.btn.sel {{ background: {ADDON['teal']}; border-color: {ADDON['teal_edge']}; }}
.btn.off {{ background: transparent; border-color: {ADDON['line']}; color: {ADDON['muted']}; font-weight: 400; }}
.link {{ color: {ADDON['gold']}; white-space: nowrap; flex: none; }}
/* Me */
.big {{ display: flex; align-items: baseline; gap: 8px; height: 30px; }}
.big .n {{ font-size: 24px; line-height: 30px; font-weight: 600; }}
.part {{ height: 18px; display: flex; align-items: center; gap: 8px; }}
.part .l {{ width: 78px; color: {ADDON['muted']}; }}
.part .tr {{ position: relative; height: 6px; background: {ADDON['inset']}; }}
.part .tr i {{ position: absolute; left: 0; top: 0; bottom: 0; }}
.part .v {{ width: 34px; text-align: right; font-variant-numeric: tabular-nums; }}
.tbl {{ width: 100%; border-collapse: collapse; }}
.tbl td, .tbl th {{ height: 18px; padding: 0 6px 0 0; white-space: nowrap; text-align: left; font-weight: 400; }}
.tbl th {{ color: {ADDON['muted']}; height: 16px; }}
.tbl .n {{ text-align: right; font-variant-numeric: tabular-nums; }}
.tbl td:last-child, .tbl th:last-child {{ padding-right: 0; }}
.now {{ display: flex; gap: 8px; align-items: flex-start; min-height: 34px; margin-bottom: 4px; }}
.now .tx {{ flex: 1; min-width: 0; }}
/* Gear */
.bands {{ display: flex; gap: 4px; margin-bottom: 10px; align-items: center; }}
.bands .b {{ height: 20px; min-width: 34px; padding: 0 6px; display: flex; align-items: center; justify-content: center;
  border: 1px solid {ADDON['line']}; color: {ADDON['muted']}; }}
.bands .b.on {{ color: {ADDON['strong']}; background: {ADDON['teal']}; border-color: {ADDON['teal_edge']}; font-weight: 600; }}
.doll {{ display: flex; justify-content: space-between; }}
.dcol {{ width: 190px; }}
.cell {{ display: flex; gap: 6px; height: 32px; margin-bottom: 4px; align-items: center; }}
.cell .tx {{ flex: 1; min-width: 0; }}
.cell.r {{ flex-direction: row-reverse; text-align: right; }}
.cell .ic {{ width: 28px; height: 28px; }}
.weap {{ display: flex; gap: 8px; justify-content: center; border-top: 1px solid {ADDON['line']}; padding-top: 8px; margin-top: 4px; }}
.weap .cell {{ width: 186px; }}
.chk {{ display: inline-block; width: 12px; height: 12px; border: 1px solid {ADDON['line']}; background: {ADDON['inset']};
  position: relative; flex: none; }}
.chk.on::after {{ content: ''; position: absolute; left: 3px; top: 0px; width: 4px; height: 8px;
  border: solid {ADDON['gold']}; border-width: 0 2px 2px 0; transform: rotate(45deg); }}
.kinds {{ display: grid; grid-template-columns: 1fr 1fr; column-gap: 16px; }}
.kinds .k {{ height: 20px; display: flex; align-items: center; gap: 6px; }}
.kinds .k .c {{ margin-left: auto; color: {ADDON['muted']}; font-variant-numeric: tabular-nums; }}
.up {{ display: flex; gap: 8px; height: 32px; margin-bottom: 4px; align-items: center; }}
.up .tx {{ flex: 1; min-width: 0; }}
/* Rotation */
.prio {{ display: flex; gap: 8px; min-height: 34px; margin-bottom: 4px; align-items: flex-start; }}
.prio .no {{ width: 12px; color: {ADDON['muted']}; text-align: right; flex: none; padding-top: 2px; }}
.prio .tx {{ flex: 1; min-width: 0; }}
.step {{ display: inline-flex; align-items: center; justify-content: center; width: 20px; height: 18px;
  border: 1px solid {ADDON['line']}; color: {ADDON['body']}; }}
/* Guild */
.signup {{ display: flex; gap: 4px; margin: 6px 0 2px; }}
.signup .btn {{ min-width: 52px; justify-content: center; }}
.job {{ border-left: 2px solid {CLIENT_GOLD}; padding: 2px 0 2px 8px; margin: 8px 0 0; }}
.cls {{ white-space: nowrap; }}
.pill {{ display: inline-block; padding: 0 4px; border: 1px solid {ADDON['line']}; color: {ADDON['muted']}; height: 14px;
  line-height: 13px; font-size: 11px; }}
/* Chip */
.mm {{ position: relative; width: 150px; height: 150px; border-radius: 50%; flex: none;
  background: radial-gradient(circle at 45% 40%, #5b6b3f, #33402a 55%, #1d2617);
  box-shadow: 0 0 0 4px #2b2416, 0 0 0 6px #6b5626, 0 4px 14px #000; }}
.chip {{ display: inline-flex; align-items: center; gap: 8px; height: 28px; padding: 0 10px 0 3px;
  background: {ADDON['raised']}; border: 1px solid {ADDON['line']}; box-shadow: 0 2px 8px rgba(0,0,0,.6);
  font-family: Site, sans-serif; font-size: 12px; color: {ADDON['body']}; white-space: nowrap; }}
.chip .more {{ color: {ADDON['muted']}; }}
.tip {{ background: rgba(8, 10, 22, .95); border: 1px solid #9095a3; border-radius: 4px; padding: 7px 10px 8px;
  font-family: Site, sans-serif; font-size: 12px; line-height: 16px; color: {ADDON['body']}; width: 260px; }}
.tip .t {{ color: #fff; font-size: 13px; margin-bottom: 2px; }}
'''


def window(tab, note, left, right=None, foot=('Esc to close', ''), cls='hunter', title='Velisandre',
           tabs=TABS, dots=()):
    strip = ''.join(f'<div class="tab{" on" if t == tab else ""}">{t}'
                    f'{"<span class=dot></span>" if t in dots else ""}</div>' for t in tabs)
    body = (f'<div class="bd"><div class="lc">{left}</div><div class="vr"></div><div class="rc">{right}</div></div>'
            if right is not None else f'<div class="bd one">{left}</div>')
    return (f'<div class="win"><div class="hd">{crest(cls, 32)}<div><div class="htt">{esc(title)}</div>'
            f'<div class="nt">{note}</div></div></div><div class="x">×</div><div class="strip">{strip}</div>'
            f'{body}<div class="ft"><span>{foot[0]}</span><span>{foot[1]}</span></div></div>')


def section(heading, body, right=''):
    r = f'<span class="r">{right}</span>' if right else ''
    return f'<div class="sec"><div class="sh"><span>{heading}</span>{r}</div>{body}</div>'


def board(caption, html):
    return page(f'<style>{CSS}</style><p class="cap">{caption}</p>{html}')


def note(*parts):
    return ' · '.join(parts)


IDENTITY = 'Level 60 Hunter'
GUILD = 'Thornwake Company'
RANK = 'Officer'
TODAY = 'Thu 10 Dec'
EXTRA_CSS = f'''
.tbl.fix {{ table-layout: fixed; }}
.tbl.fix td {{ overflow: hidden; text-overflow: ellipsis; }}
.live {{ display: inline-block; width: 6px; height: 6px; border-radius: 50%; margin-left: 4px; vertical-align: 1px; }}
.spec {{ width: 14px; height: 14px; vertical-align: -3px; margin-right: 4px; display: inline-block; box-shadow: 0 0 0 1px #000; }}
.edge {{ box-shadow: inset 2px 0 0 {CLIENT_GOLD}; padding-left: 6px; margin-left: -8px; }}
.armed {{ background: {ADDON['teal']}; border-color: {ADDON['teal_edge']}; }}
.warnmark {{ color: {ADDON['warning']}; font-weight: 600; }}
.kinds {{ grid-template-rows: repeat(4, 20px); grid-auto-flow: column; }}
.ctl {{ color: {ADDON['muted']}; margin: 0 4px 0 8px; }}
'''

# --- Velisandre's data ----------------------------------------------------------------------------

SCORE = {'gear': 410, 'experience': 380, 'parses': 228, 'attendance': 120, 'discipline': 46}
LAST_RAID = {'raid': 'Molten Core', 'date': 'Tue 8 Dec', 'bosses': [
    ('Lucifron', 71, 'Downtime', 18, 0, ['Flask'], 0.82),
    ('Magmadar', 64, 'Downtime', 24, 1, ['Flask'], 0.78),
    ('Gehennas', 83, 'Hunter\'s Mark late', 6, 0, [], 0.91),
    ('Garr', 58, 'Downtime', 31, 0, [], 0.74),
    ('Baron Geddon', 77, 'Aimed Shot late', 9, 0, [], 0.88),
    ('Shazzrah', 88, 'Serpent Sting off', 4, 0, [], 0.93),
    ('Sulfuron Harbinger', 69, 'Downtime', 15, 0, [], 0.80),
    ('Golemagg', 92, 'Downtime', 5, 0, [], 0.95),
]}
BAG_UPGRADE = 867  # Gloves of Holy Might: bought on the auction house, not yet worn


def median(values):
    v = sorted(values)
    return v[len(v) // 2] if len(v) % 2 else (v[len(v) // 2 - 1] + v[len(v) // 2]) / 2


def window2(tab, figures, left, right=None, foot=('Esc to close', ''), tabs=TABS, dots=('Gear',)):
    return window(tab, note(IDENTITY, *figures), left, right, foot, tabs=tabs, dots=dots)


def board2(caption, html):
    return page(f'<style>{CSS}{EXTRA_CSS}</style><p class="cap">{caption}</p>{html}')


# --- Me ----------------------------------------------------------------------------------------------

def score_block(stale=False):
    total = sum(SCORE.values())
    tier = tier_of(total)
    nxt = next((name, floor) for name, floor in reversed(TIER_FROM) if floor > total)
    longest = max(COMPONENT_MAX.values())
    missed = sum(len(b[5]) for b in LAST_RAID['bosses'])
    rows = ''
    for comp in SCORE_LABELS['components']:
        key = comp['key']
        fill = min(1, SCORE[key] / COMPONENT_MAX[key])
        q = 'epic' if fill >= .9 else 'rare' if fill >= .75 else 'uncommon' if fill >= .5 else 'common'
        track = round(130 * COMPONENT_MAX[key] / longest)
        why = f'<span class="m" style="margin-left:6px">{missed} flasks missed</span>' if key == 'discipline' else ''
        rows += (f'<div class="part"><span class="l">{comp["label"]}</span><span class="tr" style="width:{track}px">'
                 f'<i style="width:{round(track * fill)}px;background:{TIER_HEX[q]}"></i></span>{why}'
                 f'<span style="flex:1"></span><span class="v" style="width:70px">{SCORE[key]} <span class="m">of {COMPONENT_MAX[key]}</span></span></div>')
    head = (f'<div class="big"><span class="n" style="color:{TIER_HEX[tier]}">{total}</span>'
            f'<span style="color:{TIER_HEX[tier]}">{tier.title()}</span>'
            f'<span class="m">· {nxt[1] - total} to {nxt[0].title()} ({nxt[1]})</span></div>')
    asof = '<span class="wa">as of 2 Dec</span> · 22 fights' if stale else 'as of 9 Dec · 22 fights'
    return section('Forever Score', head + rows, asof)


def last_raid_block():
    rows = '<tr><th>Boss</th><th class="n">Parse</th><th>Biggest loss</th><th class="n">Died</th><th>Consumables</th></tr>'
    for name, pct, loss, secs, deaths, missed, _ in LAST_RAID['bosses']:
        rows += (f'<tr><td class="s">{esc(name)}</td><td class="n" style="color:{parse_hex(pct)}">{pct}</td>'
                 f'<td class="m">{esc(loss)} {secs} s</td><td class="n{" wa" if deaths else " m"}">{deaths}</td>'
                 f'<td class="{"wa" if missed else "m"}">{"No " + ", ".join(missed).lower() if missed else "All up"}</td></tr>')
    bosses = LAST_RAID['bosses']
    best, worst = max(bosses, key=lambda b: b[1]), min(bosses, key=lambda b: b[1])
    med = median([b[1] for b in bosses])
    foot = (f'<div style="margin-top:6px">Median <span style="color:{parse_hex(med)}">{med:g}</span>'
            f' <span class="m">· {sum(b[4] for b in bosses)} death · best {best[0]} {best[1]} · worst {worst[0]} {worst[1]} ·</span> '
            f'<span class="link">Full log</span></div>')
    return section('Last raid', f'<table class="tbl">{rows}</table>{foot}',
                   f'{LAST_RAID["raid"]} · {LAST_RAID["date"]} · 8 of 10 bosses')


def now_row(icon_html, line1, line2, tab):
    return (f'<div class="now">{icon_html}<div class="tx"><div class="one">{line1}</div>'
            f'<div class="m one">{line2}</div></div><span class="link">{tab}</span></div>')


def sim_block():
    rows = slot_rows()
    top = sorted((r for r in rows.values() if r['gain'] > 0), key=lambda r: -r['gain'])[:3]
    gains = ' '.join(f'<span class="row" style="gap:4px;display:inline-flex">{item_icon(r["pick"], 16)}'
                     f'<span class="ok">+{r["gain"]:.1f} DPS</span></span>' for r in top)
    return section('Your sim', '<div><span class="s" style="font-size:14px;font-weight:600">941.2 DPS</span>'
                   f' <span class="m">± 3.1 · Raid-ready list {RAID60["set_dps"]:.1f}</span></div>'
                   f'<div class="row" style="margin-top:6px;gap:10px"><span class="m">Biggest gains</span>{gains}</div>'
                   '<div class="row" style="margin-top:6px"><span class="m" style="flex:1">Mon 7 Dec</span>'
                   '<span class="btn">Sim again</span></div>')


def right_now(guild=True):
    out = ''
    if guild:
        out += now_row(icon('inv_misc_head_dragon_black', 20),
                       'Onyxia\'s Lair <span class="m">· 20:00, in 3 h 12 min</span>',
                       '<span class="ok">In</span> · Phase 2 ranged, west side', 'Guild')
    out += now_row(item_icon(BAG_UPGRADE, 20), f'{item_name(BAG_UPGRADE)} <span class="ok">+11.2 DPS</span>',
                   'Hands · in your bags', 'Gear')
    out += now_row(icon(SPELL_ICONS['2643'], 20), 'Garr: 74% followed, 31 s downtime', 'Your weakest boss on Tuesday',
                   'Rotation')
    out += now_row(icon('ability_hunter_beasttaming', 20), 'Beast Mastery raid build',
                   '<span class="ok">51 of 51 points on plan</span>', 'Builds')
    return section('Right now', out)


def board_me():
    left = score_block() + last_raid_block()
    right = right_now() + sim_block()
    return window2('Me', (GUILD, 'Beast Mastery raid build'), left, right,
                   ('Esc to close', 'Companion synced 14 min ago'), dots=('Gear', 'Guild'))


def board_me_setup():
    steps = [
        ('1', icon('ability_hunter_beasttaming', 20), 'Pick your spec', '', 'Builds'),
        ('2', item_icon(13211, 20), 'Hover an item you wear', 'Slashclaw Bracers, below', 'Gear'),
        ('3', icon('inv_scroll_08', 20),
         'Get the companion app', 'foreversixty.gg/app', ''),
    ]
    rows = ''.join(f'<div class="now"><span class="m" style="width:10px;padding-top:2px">{n}</span>{ic}<div class="tx">'
                   f'<div class="s">{a}</div><div class="{"m" if t else "link"}">{b}</div></div><span class="link">{t}</span></div>'
                   for n, ic, a, b, t in steps)
    it = ITEMS[13211]
    stats = it['stats']
    tip = ('<div class="tip" style="width:288px;margin-top:6px">'
           f'<div class="t" style="color:{QUALITY[it["quality"]]}">{esc(it["name"])}</div><div>Binds when picked up</div>'
           f'<div class="row" style="justify-content:space-between"><span>Wrist</span><span>Mail</span></div>'
           f'<div>{it["armor"]} Armor</div><div>+{stats["agility"]} Agility</div><div>+{stats["stamina"]} Stamina</div>'
           f'<div>+{stats["spirit"]} Spirit</div><div>+{stats["hit"]} Hit</div><div>Requires Level {it["required_level"]}</div>'
           f'<div style="color:#fff;margin-top:2px">Next upgrade: <span style="color:{QUALITY[3]}">Forest Stalker\'s Bracers</span>'
           f' <span style="color:{ADDON["success"]}">+9.0 DPS</span></div></div>')
    right = section('Set up', rows + tip, '3 steps')
    left = score_block()
    return window2('Me', (GUILD, 'no spec chosen'), left, right, ('Esc to close', '/fs tour'),
                   dots=('Builds',))


def board_empty():
    no_pack = (section('Forever Score', '<div class="m">No ratings pack. Forever Sixty Data, on CurseForge or Wago.</div>',
                       '') + last_raid_block())
    w1 = window2('Me', (GUILD, 'Beast Mastery raid build'), no_pack, right_now() + sim_block(),
                 ('Esc to close', 'No ratings pack · companion synced 14 min ago'), dots=('Gear', 'Guild'))
    no_comp = (score_block() + section('Last raid', '<div class="m">No raid yet. The companion app brings it in.</div>'
                                       '<div style="margin-top:4px"><span class="link">foreversixty.gg/app</span></div>'))
    right2 = section('Right now', now_row(item_icon(BAG_UPGRADE, 20), f'{item_name(BAG_UPGRADE)} <span class="ok">+11.2 DPS</span>',
                                          'Hands · in your bags', 'Gear')
                     + now_row(icon('ability_hunter_beasttaming', 20), 'Beast Mastery raid build',
                               '<span class="ok">51 of 51 points on plan</span>', 'Builds'))
    w2 = window2('Me', (GUILD, 'Beast Mastery raid build'), no_comp, right2, ('Esc to close', 'No companion'),
                 dots=('Gear',))
    stale = score_block(stale=True) + last_raid_block()
    right3 = right2
    w3 = window2('Me', ('Beast Mastery raid build',), stale, right3,
                 ('Esc to close', '<span class="wa">Ratings 8 days old</span> · companion synced 14 min ago'),
                 tabs=['Me', 'Builds', 'Gear', 'Rotation'], dots=('Gear',))
    label = '<p class="cap" style="margin:18px 0 8px">{}</p>'
    return (label.format('<b>1</b> · no ratings pack') + w1 + label.format('<b>2</b> · no companion app')
            + w2 + label.format('<b>3</b> · not in a guild, ratings pack 8 days old') + w3)


# --- Gear -------------------------------------------------------------------------------------------------

LEFT_SLOTS = ['head', 'neck', 'shoulder', 'back', 'chest', 'wrist', 'main_hand', 'ranged']
RIGHT_SLOTS = ['hands', 'waist', 'legs', 'feet', 'finger1', 'finger2', 'trinket1', 'trinket2']
NO_ENCHANT = {'chest', 'feet'}


def doll_cell(r, side=''):
    best = r['worn'] == r['pick']
    ench = '<span class="warnmark">No enchant</span>' if r['slot'] in NO_ENCHANT else ''
    if best:
        status = '<span class="m">Best in slot</span>'
    else:
        status = f'<span class="ok">+{r["gain"]:.1f} DPS</span>'
    if r['slot'] == 'hands':
        status += ' <span class="cg">in your bags</span>'
    parts = [status, ench] if side != 'r' else [ench, status]
    line2 = ' · '.join(p for p in parts if p) if side != 'r' else ' · '.join(p for p in parts if p)
    return (f'<div class="cell {side}">{item_icon(r["worn"], 28)}<div class="tx"><div class="one">{item_name(r["worn"])}</div>'
            f'<div class="one">{line2}</div></div></div>')


def source_words(r):
    src = esc(r['source'])
    slot = next(s for s in RAID60['slots'] if s['slot'] == r['slot'])
    if slot.get('standing'):
        src += f' · {slot["standing"].title()}'
    return src


def weights_block():
    weights = [w for w in RAID60['weights'] if w['scale_factor'] > 0 and w['stat'] != 'melee_haste']
    top = max(w['dps_per_point'] for w in weights)
    names = {'hit': 'Hit', 'agility': 'Agility', 'crit': 'Crit', 'intellect': 'Intellect',
             'ranged_attack_power': 'Ranged attack power'}
    rows = '<tr><th>Stat</th><th class="n">Weight</th><th class="n">DPS</th><th class="n">±</th></tr>'
    for w in sorted(weights, key=lambda w: -w['dps_per_point']):
        rows += (f'<tr><td class="s">{names[w["stat"]]}</td><td class="n">{w["dps_per_point"] / top:.2f}</td>'
                 f'<td class="n m">{w["dps_per_point"]:.2f}</td><td class="n m">{w["scale_error"] * w["dps_per_point"] / max(w["scale_factor"], 1e-9):.2f}</td></tr>')
    haste = next(w for w in RAID60['weights'] if w['stat'] == 'melee_haste')
    return section('Stat weights', f'<table class="tbl">{rows}</table><div class="m" style="margin-top:2px">'
                   f'Haste {haste["dps_per_point"]:.2f} DPS per 1%, ± {haste["scale_error"] * haste["dps_per_point"] / haste["scale_factor"]:.2f}</div>',
                   'your sim, per rating point')


def board_gear():
    rows = slot_rows()
    best = sum(1 for r in rows.values() if r['worn'] == r['pick'])
    bands = ''.join(f'<span class="btn{" sel" if b == "60" else ""}" style="min-width:30px;justify-content:center">{b}</span>'
                    for b in ('20', '30', '40', '50', '60'))
    presets = '<span class="btn sel">Raid-ready</span><span class="btn">Bare</span>'
    left = (f'<div class="bands"><span class="m" style="width:66px">Level</span>{bands}</div>'
            f'<div class="bands"><span class="m" style="width:66px">Conditions</span>{presets}</div>'
            f'<div class="doll"><div class="dcol">{"".join(doll_cell(rows[s]) for s in LEFT_SLOTS)}</div>'
            f'<div class="dcol">{"".join(doll_cell(rows[s], "r") for s in RIGHT_SLOTS)}</div></div>')
    ups = sorted((r for r in rows.values() if r['gain'] > 0), key=lambda r: -r['gain'])[:4]
    up_rows = ''
    for r in ups:
        where = '<span class="cg">In your bags</span>' if r['pick'] == BAG_UPGRADE else source_words(r)
        up_rows += (f'<div class="up">{item_icon(r["pick"], 24)}<div class="tx"><div class="row"><span class="one" style="flex:1">'
                    f'{item_name(r["pick"])}</span><span class="ok">+{r["gain"]:.1f} DPS</span></div>'
                    f'<div class="m one">{SLOT_NAMES[r["slot"]]} · {where}</div></div></div>')
    total = sum(r['gain'] for r in rows.values())
    counts = kind_counts()
    on = {k for k, _ in KINDS} - {'pvp'}
    kinds = ''.join(f'<div class="k"><span class="chk{" on" if k in on else ""}"></span><span>{label}</span>'
                    f'<span class="c">{counts[k]}</span></div>' for k, label in KINDS)
    right = (section('Upgrades', up_rows, f'+{total:.1f} DPS in all')
             + weights_block()
             + section('Sources', f'<div class="kinds">{kinds}</div><div class="m" style="margin-top:4px">'
                       f'Off: PvP · <span class="link">Default</span></div>', '7 of 8 sources'))
    return window2('Gear', (f'{best} of {len(rows)} best in slot', f'{len(NO_ENCHANT)} enchants missing'), left, right,
                   ('Esc to close', 'Shift-click an item to link it'), dots=('Gear', 'Guild'))


# --- Rotation ---------------------------------------------------------------------------------------------

def board_rotation():
    lines = rotation_lines('hunter-beast-mastery', 60)
    prio = ''.join(f'<div class="prio"><span class="no">{i}</span>{icon(SPELL_ICONS[sid], 24)}<div class="tx">'
                   f'<div class="s">{esc(name)}</div><div class="m">{esc(cond[0].upper() + cond[1:])}</div></div></div>'
                   for i, (sid, name, cond) in enumerate(lines, 1))
    step = ('<span class="row" style="gap:4px"><span class="btn">‹</span><span class="s" style="min-width:56px;'
            'text-align:center">Level 60</span><span class="btn">›</span></span>')
    left = f'<div class="sec"><div class="sh"><span>Priority</span>{step}</div>{prio}</div>'
    bosses = LAST_RAID['bosses']
    overall = round(100 * sum(b[6] for b in bosses) / len(bosses))
    adh = ''.join(f'<tr><td class="s">{esc(b[0])}</td><td class="n s">{round(b[6] * 100)}%</td>'
                  f'<td class="m">{esc(b[2])} {b[3]} s</td></tr>' for b in sorted(bosses, key=lambda b: b[6]))
    right = section('How you played it',
                    f'<div style="margin-bottom:4px">Followed <span class="s">{overall}%</span>'
                    f' <span class="m">across 8 bosses</span></div>'
                    f'<table class="tbl"><tr><th>Boss</th><th class="n">Followed</th><th>Biggest loss</th></tr>{adh}</table>',
                    'Molten Core · Tue 8 Dec')
    losses = {}
    for b in bosses:
        losses[b[2]] = losses.get(b[2], 0) + b[3]
    lrows = ''.join(f'<tr><td>{esc(k)}</td><td class="n wa">{v} s</td></tr>' for k, v in sorted(losses.items(), key=lambda kv: -kv[1]))
    right += section('Lost across the night', f'<table class="tbl">{lrows}</table>')
    return window2('Rotation', ('Beast Mastery raid build', f'followed {overall}% on Tuesday'), left, right,
                   ('Esc to close', ''), dots=('Gear', 'Guild'))


# --- Guild and Officer data ---------------------------------------------------------------------------------
# name, class, spec, score, attendance %, sign-up, addon ('live', 'off' or None), DPS gap, enchants missing,
# item level, role, has Onyxia Scale Cloak, has Fire Protection Potions

ROSTER = [
    ('Brannoc', 'warrior', 'warrior-protection', 1612, 100, 'In', 'live', 0.0, 0, 66, 'Tank', True, True),
    ('Ysolde', 'priest', 'priest-holy', 1490, 96, 'In', 'live', 0.0, 0, 64, 'Healer', True, True),
    ('Torvald', 'warrior', 'warrior-fury', 1421, 94, 'In', 'live', 21.4, 0, 63, 'Melee', True, True),
    ('Kethra', 'rogue', 'rogue-combat', 1377, 92, 'In', 'live', 14.0, 1, 62, 'Melee', True, True),
    ('Mirelle', 'mage', 'mage-frost', 1302, 90, 'In', 'live', 9.8, 0, 62, 'Ranged', True, True),
    ('Nessaly', 'mage', 'mage-fire', 1268, 88, 'In', 'off', None, None, 61, 'Ranged', True, None),
    ('Dravik', 'warlock', 'warlock-destruction', 1254, 84, 'Late', 'live', 18.6, 0, 61, 'Ranged', True, False),
    ('Velisandre', 'hunter', 'hunter-beast-mastery', 1184, 88, 'In', 'live', 40.4, 2, 61, 'Ranged', True, False),
    ('Holbrand', 'paladin', 'paladin-protection', 1105, 76, 'In', 'live', 0.0, 0, 62, 'Tank', True, True),
    ('Saffa', 'druid', 'druid-restoration', 998, 70, 'In', 'live', 0.0, 2, 58, 'Healer', False, False),
    ('Orwen', 'priest', 'priest-holy', 1050, 82, 'In', None, None, None, 60, 'Healer', True, None),
    ('Tamsin', 'paladin', 'paladin-holy', 1142, 86, 'In', 'live', 0.0, 0, 61, 'Healer', True, True),
    ('Ilvar', 'hunter', 'hunter-marksmanship', 1033, 80, 'Bench', 'live', 26.2, 0, 59, 'Ranged', True, True),
    ('Quenna', 'priest', 'priest-shadow', 861, 52, 'Bench', None, None, None, 55, 'Ranged', False, None),
    ('Garrow', 'rogue', 'rogue-combat', 1188, 90, 'Out', 'live', 12.5, 0, 62, 'Melee', True, True),
    ('Pell', 'warlock', 'warlock-destruction', 920, 61, '', None, None, None, 57, 'Ranged', False, None),
]
SPEC_ICON = {'warrior-protection': 'inv_shield_06', 'warrior-fury': 'ability_warrior_innerrage',
             'priest-holy': 'spell_holy_guardianspirit', 'priest-shadow': 'spell_shadow_shadowwordpain',
             'rogue-combat': 'ability_backstab', 'mage-frost': 'spell_frost_frostbolt02', 'mage-fire': 'spell_fire_firebolt02',
             'warlock-destruction': 'spell_shadow_rainoffire', 'hunter-beast-mastery': 'ability_hunter_beasttaming',
             'hunter-marksmanship': 'ability_marksmanship', 'paladin-protection': 'spell_holy_devotionaura',
             'paladin-holy': 'spell_holy_holybolt', 'druid-restoration': 'spell_nature_healingtouch'}
RAIDERS = 31
NOT_SIGNED = RAIDERS - sum(1 for r in ROSTER if r[5])
WANTED = {'Tank': 4, 'Healer': 10, 'Melee': 12, 'Ranged': 14}
SIGN_CLASS = {'In': 'ok', 'Late': 'cg', 'Bench': 'm', 'Out': 'wa'}


def coming():
    return [r for r in ROSTER if r[5] in ('In', 'Late')]


def signups_line():
    c = {a: sum(1 for r in ROSTER if r[5] == a) for a in ('In', 'Late', 'Bench', 'Out')}
    return (f'<span class="ok">{c["In"]} In</span> · <span class="cg">{c["Late"]} Late</span> · {c["Bench"]} Bench · '
            f'<span class="wa">{c["Out"]} Out</span> · {NOT_SIGNED} not signed up')


def roles_line():
    have = {k: sum(1 for r in coming() if r[10] == k) for k in WANTED}
    return ' · '.join(f'{k}s {have[k]} of {WANTED[k]}' if k in ('Tank', 'Healer') else f'{k} {have[k]} of {WANTED[k]}'
                      for k in WANTED)


def name_cell(name, cls, spec=None, live=None):
    sp = f'<img class="spec" src="{icon_uri(SPEC_ICON[spec])}">' if spec else ''
    dot = ''
    if live == 'live':
        dot = f'<span class="live" style="background:{ADDON["success"]}"></span>'
    elif live == 'off':
        dot = f'<span class="live" style="background:{ADDON["muted"]}"></span>'
    return f'{sp}<span class="cls" style="color:{CLASS_HEX[cls]}">{esc(name)}</span>{dot}'


def names(rows):
    return ', '.join(name_cell(r[0], r[1]) for r in rows)


def standing():
    ranked = sorted(ROSTER, key=lambda r: -r[3])
    return next(i for i, r in enumerate(ranked, 1) if r[0] == 'Velisandre')


# --- Guild ------------------------------------------------------------------------------------------------

def board_guild():
    me = next(r for r in ROSTER if r[0] == 'Velisandre')
    tonight = (f'<div class="row">{icon("inv_misc_head_dragon_black", 32)}<div style="flex:1">'
               f'<div class="s" style="font-size:14px;font-weight:600">Onyxia\'s Lair</div>'
               f'<div class="m">{TODAY}, 20:00 server · in 3 h 12 min · invites 19:45</div></div></div>'
               f'<div style="margin-top:8px">Your sign-up: <span class="ok">In</span> <span class="m">as DPS, sent Tue 19:04</span></div>'
               f'<div class="signup"><span class="btn sel">In</span><span class="btn">Late</span><span class="btn">Bench</span>'
               f'<span class="btn">Out</span></div>'
               f'<div style="margin-top:8px">Your job: <span class="cg">Phase 2 ranged, west side</span></div>'
               f'<div class="m">Fire on the whelps from the west wall when she lifts off. With {name_cell("Mirelle", "mage")}.</div>')
    left = section('Tonight', tonight, 'Raid leader ' + name_cell('Brannoc', 'warrior'))
    prep = ('<table class="tbl">'
            f'<tr><td class="m">Attunement</td><td>{item_name(16309)} <span class="ok">in your bags</span></td></tr>'
            '<tr><td class="m">Fire resistance</td><td>42</td></tr>'
            '<tr><td class="m">Missing</td><td class="wa">Greater Fire Protection Potion</td></tr></table>')
    left += section('Your prep', prep, 'for Onyxia')
    left += section('Sign-ups', f'<div>{signups_line()}</div><div class="m">{roles_line()}</div>',
                    f'{len(coming())} of 40 coming')
    nxt = [(f'{TODAY} 20:00', 'Onyxia\'s Lair', 'In as DPS', 'ok'), ('Sun 13 Dec 19:30', 'Molten Core', 'Not signed up', 'cg'),
           ('Tue 15 Dec 19:30', 'Molten Core', 'Not signed up', 'cg')]
    wrows = ''.join(f'<tr><td class="m">{d}</td><td class="s">{n}</td><td class="{c}">{a}</td></tr>' for d, n, a, c in nxt)
    right = section('Next raids', f'<table class="tbl">{wrows}</table>')
    top = sorted(ROSTER, key=lambda r: -r[3])[:5]
    trows = ''.join(f'<tr><td class="m n">{i}</td><td>{name_cell(r[0], r[1], r[2])}</td>'
                    f'<td class="n" style="color:{TIER_HEX[tier_of(r[3])]}">{r[3]}</td><td class="n m">{r[4]}%</td></tr>'
                    for i, r in enumerate(top, 1))
    rank = standing()
    trows += (f'<tr><td class="m n">{rank}</td><td>{name_cell(me[0], me[1], me[2])}</td>'
              f'<td class="n" style="color:{TIER_HEX[tier_of(me[3])]}">{me[3]}</td><td class="n m">{me[4]}%</td></tr>')
    right += section(GUILD, '<div class="m" style="margin-bottom:4px">Molten Core 10/10 · Onyxia 1/1 · Thu, Sun, Tue</div>'
                     f'<table class="tbl"><tr><th class="n">#</th><th>Raider</th><th class="n">Score</th><th class="n">Att.</th></tr>{trows}</table>',
                     f'you are {rank}th of {RAIDERS}')
    right += section('Officers sent', f'<div class="row">{icon("ability_hunter_beasttaming", 20)}<div style="flex:1">'
                     f'<div>Beast Mastery raid build</div><div class="m">from {name_cell("Brannoc", "warrior")}, Tue · you follow it</div></div></div>')
    return window2('Guild', (GUILD, RANK), left, right, ('Esc to close', ''), dots=('Gear',))


# --- Officer ------------------------------------------------------------------------------------------------

def roster_table():
    order = {'In': 0, 'Late': 1, 'Bench': 2, 'Out': 3, '': 4}
    rows = sorted(ROSTER, key=lambda r: (order[r[5]], -r[3]))
    cols = '<colgroup><col style="width:104px"><col style="width:40px"><col style="width:38px"><col style="width:72px">' \
           '<col style="width:38px"><col style="width:36px"><col style="width:24px"></colgroup>'
    head = ('<tr><th>Raider</th><th class="n">Score</th><th class="n">Att.</th><th>Sign-up</th>'
            '<th class="n">Gap</th><th class="n">Ench.</th><th class="n">iLvl</th></tr>')
    body = ''
    for name, cls, spec, score, att, sign, live, gap, ench, ilvl, *_ in rows:
        gap_cell = '' if gap is None else ('<span class="m">Best</span>' if gap == 0 else f'{gap:.1f}')
        ench_cell = '' if ench is None else (f'<span class="wa">{ench}</span>' if ench else '<span class="m">0</span>')
        body += (f'<tr><td>{name_cell(name, cls, spec, live)}</td>'
                 f'<td class="n" style="color:{TIER_HEX[tier_of(score)]}">{score}</td><td class="n">{att}%</td>'
                 f'<td class="{SIGN_CLASS.get(sign, "m")}">{sign or "Not signed"}</td>'
                 f'<td class="n">{gap_cell}</td><td class="n">{ench_cell}</td><td class="n m">{ilvl}</td></tr>')
    legend = (f'<span class="live" style="background:{ADDON["success"]};margin:0 4px 0 0"></span>addon live '
              f'<span class="live" style="background:{ADDON["muted"]};margin:0 4px 0 8px"></span>sharing off · Gap in DPS')
    return (f'<table class="tbl fix">{cols}{head}{body}</table><div class="m" style="margin-top:4px">'
            f'{len(rows)} of {RAIDERS} shown · <span class="link">Show all</span> · '
            f'<span class="link">Not signed up ({NOT_SIGNED})</span></div><div class="m" style="margin-top:2px">{legend}</div>')


def board_officer():
    left = section('Roster', roster_table(), 'by sign-up, then Score')
    c = coming()
    live = [r for r in c if r[6] == 'live']
    on_build = [r for r in live if r[0] not in ('Dravik',)]
    ench = [r for r in c if r[8]]
    cloak = [r for r in c if not r[11]]
    pots = [r for r in live if r[12] is False]
    ready = ('<table class="tbl">'
             f'<tr><td class="m">Addon live</td><td class="n s">{len(live)} of {len(c)}</td><td>{names([r for r in c if r[6] != "live"])} without</td></tr>'
             f'<tr><td class="m">Raid build</td><td class="n s">{len(on_build)} of {len(live)}</td><td>{name_cell("Dravik", "warlock")} off</td></tr>'
             f'<tr><td class="m">Enchants</td><td class="n wa">{sum(r[8] for r in ench)}</td><td>{", ".join(name_cell(r[0], r[1]) + f" {r[8]}" for r in ench)}</td></tr>'
             f'<tr><td class="m">No fire prot.</td><td class="n wa">{len(pots)}</td><td>{names(pots)}</td></tr>'
             f'<tr><td class="m">Jobs</td><td class="n s">9 of 10</td><td class="wa">Off tank open</td></tr>'
             '</table>')
    right = section('Tonight\'s readiness', ready, 'Onyxia\'s Lair · 20:00')
    right += section('Sign-ups', f'<div>{signups_line()}</div><div class="m">{roles_line()}</div>'
                     '<div class="row" style="margin-top:6px"><span class="m" style="flex:1">3 approvals on the site</span>'
                     f'<span class="btn">Remind {NOT_SIGNED} in chat</span></div>')
    right += section('Push a build', f'<div class="row">{icon("ability_hunter_beasttaming", 20)}<div style="flex:1;min-width:0">'
                     f'<div class="one">Beast Mastery raid build</div><div class="m one">sent Tue · 9 follow it</div></div>'
                     f'<span class="btn">Send to guild</span></div>')
    right += section('Master loot', '<div class="m">Opens here when you loot as master looter. Tuesday: 7 items, '
                     '6 to the biggest gain.</div>')
    return window2('Officer', (GUILD, RANK), left, right,
                   ('Esc to close', 'Click a name to whisper · Shift-click to invite'), dots=('Gear',))


def board_officer_loot():
    item_id = 17069  # Striker's Mark, Magmadar
    rows = [('Velisandre', 'hunter', 'hunter-beast-mastery', 21.1, 12651 if False else 13023, 0, 88, True),
            ('Ilvar', 'hunter', 'hunter-marksmanship', 26.0, 12651, 1, 80, False)]
    head = '<tr><th>Raider</th><th class="n">DPS</th><th>Wears</th><th class="n">Won</th><th class="n">Att.</th><th></th></tr>'
    body = ''
    for i, (name, cls, spec, gain, worn, won, att, listed) in enumerate(sorted(rows, key=lambda r: -r[3])):
        btn = '<span class="btn armed">Give to Ilvar?</span>' if i == 0 else '<span class="btn">Give</span>'
        flag = ' <span class="g">on list</span>' if listed else ''
        body += (f'<tr><td>{name_cell(name, cls, spec, "live")}</td><td class="n ok">+{gain:.1f}</td>'
                 f'<td><span class="row" style="gap:4px">{item_icon(worn, 14)}<span class="one" style="max-width:66px">{item_name(worn)}</span></span></td><td class="n">{won}</td><td class="n m">{att}%</td>'
                 f'<td style="text-align:right">{btn}</td></tr>')
    left = section('Master loot', f'<div class="row" style="margin-bottom:8px">{item_icon(item_id, 32)}<div style="flex:1;min-width:0">'
                   f'<div class="one" style="font-size:13px">{item_name(item_id)}</div><div class="m">Ranged · Magmadar</div></div></div>'
                   f'<table class="tbl">{head}{body}</table>'
                   f'<div class="m" style="margin-top:6px">2 of 2 hunters answered · {name_cell("Torvald", "warrior")}, '
                   f'{name_cell("Kethra", "rogue")} no gain · {name_cell("Orwen", "priest")} no addon</div>',
                   '<span class="link">Roster</span> · item 2 of 3')
    sess = [(16847, 'Ilvar', 'hunter', '+14.8'), (17069, None, None, 'open'), (18823, None, None, 'waiting')]
    srows = ''.join(f'<div class="up">{item_icon(i, 20)}<div class="tx"><div class="one">{item_name(i)}</div>'
                    f'<div class="m">{("to " + name_cell(n, c) + " <span class=ok>" + g + " DPS</span>") if n else g}</div></div></div>'
                    for i, n, c, g in sess)
    right = section('Magmadar', srows, '3 items')
    return window2('Officer', (GUILD, 'Master looter'), left, right,
                   ('Esc for the roster', ''), dots=())


# --- the chip ------------------------------------------------------------------------------------------------

def chip(line, more=0, cls='hunter'):
    m = f' <span class="more">+{more}</span>' if more else ''
    return f'<span class="chip" style="max-width:280px">{crest(cls, 22)}<span class="one">{line}{m}</span></span>'


def board_chip():
    states = [
        ('Raid day, signed up', chip('Onyxia\'s Lair 20:00 · <span class="ok">In</span>', 2), 'Opens Guild'),
        ('Raid day, not signed up', chip('Onyxia\'s Lair 20:00 · <span class="cg">Sign up</span>', 1), 'Opens Guild'),
        ('An upgrade in the bags', chip('<span class="ok">Equip</span> Gloves of Holy Might <span class="ok">+11.2 DPS</span>'), 'Opens Gear'),
        ('Builds from the site', chip('<span class="cg">2 builds</span> from the site'), 'Opens Builds'),
        ('After a raid night', chip('Tuesday: median <span style="color:' + parse_hex(74) + '">74</span>, best Golemagg'), 'Opens Me'),
        ('Master loot', chip('Striker\'s Mark · 2 of 2 answered'), 'Opens Officer'),
        ('A level 42 alt, a point waiting', chip('<span class="cg">Spend 1 point:</span> Intimidation'), 'Opens Builds'),
        ('In combat', chip('<span class="m">After combat</span>'), 'Does not open'),
        ('Nothing waiting', chip('<span class="m">Forever Sixty</span>'), 'Opens the last tab'),
    ]
    rows = ''.join(f'<div class="row" style="height:36px;gap:16px"><span style="width:250px;color:#c9c2b2">{a}</span>'
                   f'<span style="width:290px">{c}</span><span style="color:#9a9484">{d}</span></div>' for a, c, d in states)
    tip = ('<div class="tip"><div class="t">Forever Sixty</div>'
           '<div>Onyxia\'s Lair 20:00 · In</div><div>Equip Gloves of Holy Might +11.2 DPS</div>'
           '<div>2 builds from the site</div>'
           '<div style="margin-top:6px;color:#b8b8b8">Click: open on Guild<br>Right-click: settings<br>Shift-drag: move</div></div>')
    scene = (f'<div style="position:relative;width:440px;height:330px">'
             f'<div class="mm" style="position:absolute;right:0;top:0"></div>'
             f'<div style="position:absolute;right:166px;top:10px">{chip("Onyxia&#39;s Lair 20:00 · <span class=ok>In</span>", 2)}</div>'
             f'<div style="position:absolute;right:166px;top:48px">{tip}</div></div>')
    return f'<div class="row" style="align-items:flex-start;gap:40px"><div>{rows}</div>{scene}</div>'


BOARDS = {
    'hub-me': ('<b>Me</b> · Velisandre, a made-up level 60 Beast Mastery hunter on raid day (Thu 10 Dec). Score and parts '
               'from the ratings pack, last raid and sim from the companion inbox; Right now links each line to its tab.',
               board_me),
    'hub-me-setup': ('<b>Me, first open</b> · replaces the tour: the window opens once on Me with three steps, each a link '
                     'to its tab, and the tooltip line the second step points at.', board_me_setup),
    'hub-gear': ('<b>Gear</b> · the paperdoll at level 60, Raid-ready, against what she wears; the gloves she bought sit in '
                 'her bags. Upgrades, the stat weights from her sim and the source filter, PvP off.', board_gear),
    'hub-rotation': ('<b>Rotation</b> · the level 60 priority from Data.lua and how Tuesday\'s raid went against it.',
                     board_rotation),
    'hub-guild': ('<b>Guild</b> · tonight\'s raid, her sign-up and job, her prep for Onyxia, the guild\'s sign-ups, the '
                  'next raids, the top of the guild with her place, the build the officers sent. Thornwake Company is made up.',
                  board_guild),
    'hub-officer': ('<b>Officer</b> · shown when the rank can edit officer notes: the raiders by sign-up, readiness from '
                    'members running the addon and from inspect, sign-ups, build push, master loot.', board_officer),
    'hub-officer-loot': ('<b>Officer, master loot</b> · opened by looting Magmadar as master looter. Give arms on the first '
                         'click (teal, names the raider) and hands the item out on the second.', board_officer_loot),
    'hub-chip': ('<b>The chip</b> · the launcher beside the minimap: one line, the most urgent thing, and a click opens '
                 'the window on the tab that owns it. Capped at 280 px. Hover lists everything waiting.', board_chip),
    'hub-empty': ('<b>Empty states</b> · each section says what fills it, once, in its own place.', board_empty),
}


def main(names_):
    from playwright.sync_api import sync_playwright
    OUT.mkdir(parents=True, exist_ok=True)
    with sync_playwright() as pw:
        exe = os.environ.get('PLAYWRIGHT_CHROMIUM')
        browser = pw.chromium.launch(executable_path=exe) if exe else pw.chromium.launch()
        ctx = browser.new_context(device_scale_factor=2, viewport={'width': 1400, 'height': 1000})
        pg = ctx.new_page()
        for name in names_:
            caption, build = BOARDS[name]
            pg.set_content(board2(caption, build()))
            pg.evaluate('document.fonts.ready')
            pg.wait_for_timeout(150)
            pg.locator('#board').screenshot(path=str(OUT / f'{name}.png'))
            print(OUT / f'{name}.png')
        browser.close()


if __name__ == '__main__':
    main(sys.argv[1:] or list(BOARDS))
