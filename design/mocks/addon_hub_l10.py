"""Review boards for lane L10, the marks in the world and Settings as built (design/specs/2026-11-13-addon-hub.md 5, 7).

    python3 design/mocks/addon_hub_l10.py              # every board
    python3 design/mocks/addon_hub_l10.py l10-settings

Writes design/reviews/2026-11-13-addon-hub/l10/<board>.png at device scale 2, every frame at its real size, drawn to
the built code's own rules, not the approved boards':

- Quest reward and roll marks: addon/ForeverSixty/surfaces/RewardMarks.lua. The band's pick says the item tooltip's
  own gold line (ItemTooltip.bestText: "Best for 20 to 29", "Best at level 60") in the client gold #ffd100 with the
  client's 1 px shadow; "Sells for most" in the tooltip grey. Items and stats: data/builds/1.60.1.70790/items/hunter.json
  through design/mocks/addon_ux.py; the Tunic of Westfall gain is addon_ux.tunic_gain (band-20 DPS per point).
- Player tooltip: surfaces/PlayerTooltip.lua. The date joins the score line once the record is stale by
  Score.isStale (3 days, Score.STALE_AFTER_SECONDS), the rule the chat badge and the Me tab already use.
- Paper-doll slot and bag item: surfaces/SlotMarks.lua and ItemTooltip.withAltClick. The slot's hover leads with the bag
  fact (ItemTooltip.bagSection); under Shift the attribution, then the marks' one wording "Alt-click: open Gear"
  (L.markAltGear). Velisandre is the Gear lane's (addon_hub_l3.velisandre, Raid-ready): the doll shows what she wears,
  the bag item is her list's Hands pick and the gain is that tab's, so the marks and the Gear tab agree. Her other bag
  items are a raid night's consumables, icons from simbuffs.json.
- Talent strip: surfaces/TalentMarks.lua; its hover is the one line L.markClickBuilds, "Click: open Builds", dim.
- Settings: surfaces/Settings.lua's layout (padding 16, title 24, line 18, row 44 at least, the 20 px client button,
  foot pad 12) inside the client's Settings canvas, 600 wide; the page scrolls past 560, drawn here whole.

Every character here is made up: Aldrenna, a level 21 Night Elf hunter; Velisandre, a level 60 Night Elf hunter;
Kestrin and Orlaith, two players she hovers.
"""
import os
import sys

import addon_hub as hub
import addon_hub_l3 as gear_lane
import addon_ux as ux
from addon_hub import ADDON, CLIENT_GOLD, QUALITY, esc
from addon_ux import C, ROOT, arrow, img, l, tooltip

OUT = ROOT / 'design/reviews/2026-11-13-addon-hub/l10'
DIM = '#8a8a8a'  # Theme.tooltipDim, the attribution and the hint lines
GREY = C['grey']
SHADOW = 'text-shadow: 1px -1px 0 #000, 1px 1px 0 #000;'

CSS = f'''
.cap {{ max-width: 1000px; }}
.qf {{ width: 384px; }}
.qbody {{ margin: 10px; padding: 14px; border-radius: 3px; color: #2a1a08;
  background: radial-gradient(ellipse at 40% 30%, #efe0bd, #d7c193 70%, #c4ab7a); }}
.rw {{ display: grid; grid-template-columns: 1fr 1fr; gap: 6px 8px; }}
.rb {{ position: relative; display: flex; align-items: center; gap: 4px; height: 42px; }}
.rb .nf {{ flex: 1; height: 38px; padding: 3px 6px; border-radius: 2px; background: linear-gradient(#3b3326, #24201a);
  box-shadow: inset 0 0 0 1px #57492f; display: flex; flex-direction: column; justify-content: center; }}
.rb .nm {{ font-size: 12px; line-height: 14px; color: #fff; }}
.rb .lb {{ font-family: Site, sans-serif; font-size: 10px; line-height: 13px; {SHADOW} white-space: nowrap; }}
.rb .mk {{ position: absolute; left: 26px; top: -2px; }}
.rf {{ position: relative; width: 340px; padding: 8px 10px 12px; display: flex; gap: 10px; align-items: center; }}
.rf .rbtn {{ width: 24px; height: 24px; border-radius: 50%; display: inline-flex; align-items: center;
  justify-content: center; color: #f0d27a; font-size: 15px; background: radial-gradient(#3a3020, #15120c);
  box-shadow: 0 0 0 1px #6b5626; }}
.rf .timer {{ position: absolute; left: 10px; right: 10px; bottom: 4px; height: 4px; background: #2a1d0c; }}
.rf .timer i {{ display: block; width: 62%; height: 100%; background: linear-gradient(#e0b23a, #a87a16); }}
.rf .lb {{ font-family: Site, sans-serif; font-size: 10px; line-height: 13px; white-space: nowrap; {SHADOW} }}
.pane {{ width: max-content; padding: 10px; display: flex; flex-direction: column; gap: 6px; }}
.slot {{ position: relative; width: 40px; height: 40px; }}
.slot .mk {{ position: absolute; top: -4px; right: -5px; }}
.bags {{ display: grid; grid-template-columns: repeat(4, 40px); gap: 4px; padding: 10px; }}
.slot .ct {{ position: absolute; right: 2px; bottom: 1px; font-size: 12px; line-height: 12px; color: #fff;
  text-shadow: 1px 1px 0 #000, -1px -1px 0 #000, 1px -1px 0 #000, -1px 1px 0 #000; }}
.half {{ font-family: Site, sans-serif; font-size: 13px; color: {ADDON['muted']}; }}
.tw {{ position: relative; width: 384px; height: 150px; padding: 0; overflow: visible;
  background: linear-gradient(#1b1d24, #101218); }}
.tstrip {{ position: absolute; left: 12px; bottom: 10px; height: 26px; display: flex; align-items: center;
  padding: 0 10px; background: rgba(16, 10, 6, .82); font-family: Site, sans-serif; font-size: 13px;
  color: {ADDON['body']}; white-space: nowrap; outline: 1px dashed rgba(229,185,85,.0); }}
.apply {{ position: absolute; right: 12px; bottom: 12px; }}
/* The client's Settings window, its AddOns list and our canvas */
.sw {{ display: flex; width: 812px; }}
.sw .side {{ width: 196px; padding: 10px 8px; border-right: 1px solid #3b311d; font-size: 13px; color: #d6d6d6; }}
.sw .side div {{ height: 22px; display: flex; align-items: center; padding-left: 8px; }}
.sw .side .on {{ background: rgba(255, 209, 0, .14); color: {CLIENT_GOLD}; }}
.cv {{ width: 600px; padding: 8px; }}
.pg {{ background: {ADDON['raised']}; border: 1px solid {ADDON['line']}; font-family: Site, sans-serif; font-size: 12px;
  line-height: 15px; color: {ADDON['body']}; padding: 16px 16px 12px; }}
.pg .ptt {{ font-family: Title, serif; font-size: 16px; line-height: 24px; height: 24px; color: {ADDON['cream']};
  text-shadow: 1px -1px 0 #000; }}
.pg .wn, .pg .ln.wn {{ color: {ADDON['warning']}; }}
.pg .phd {{ font-family: Title, serif; font-size: 14px; height: 18px; line-height: 18px; color: {ADDON['cream']};
  text-shadow: 1px -1px 0 #000; }}
.pg .ln {{ min-height: 18px; line-height: 18px; color: {ADDON['muted']}; }}
.pg .ln.wn {{ color: {ADDON['warning']}; }}
.pg .ln b {{ color: {ADDON['body']}; font-weight: 400; }}
.pg .sw1 {{ display: flex; gap: 8px; min-height: 44px; padding: 4px 0 0; }}
.pg .box {{ width: 16px; height: 16px; flex: none; border: 1px solid {ADDON['muted']}; position: relative; }}
.pg .box.on {{ background: {ADDON['gold']}; border-color: {ADDON['gold']}; }}
.pg .box.on::after {{ content: ''; position: absolute; left: 5px; top: 1px; width: 4px; height: 9px;
  border: solid #1a1200; border-width: 0 2px 2px 0; transform: rotate(45deg); }}
.pg .lbl {{ color: {ADDON['body']}; }}
.pg .hint {{ color: {ADDON['muted']}; width: 528px; margin-top: 4px; padding-bottom: 10px; }}
.pg .rs {{ display: flex; align-items: center; gap: 8px; height: 20px; }}
.pg .cb {{ width: 150px; height: 20px; display: flex; align-items: center; justify-content: center; font-weight: 600;
  color: {ADDON['strong']}; background: {ADDON['blue']}; border: 1px solid {ADDON['blue_edge']}; flex: none; }}
'''


def board(caption, html):
    return hub.board2(f'{caption}', f'<style>{CSS}</style>{html}')


def col(*parts, gap=14):
    return f'<div class="col" style="gap:{gap}px">{"".join(parts)}</div>'


def row(*parts, gap=22):
    return f'<div class="row" style="gap:{gap}px;align-items:flex-start">{"".join(parts)}</div>'


def note(text):
    return f'<div class="note">{text}</div>'


def dim(text):
    return l(text, color=DIM)


# --- Quest reward and roll -------------------------------------------------------------------------------

def reward(item_id, label='', marked=False, label_color=CLIENT_GOLD, unusable=False):
    it = ux.item(item_id)
    tint = 'filter:sepia(1) saturate(6) hue-rotate(-50deg) brightness(.8)' if unusable else ''
    mark = f'<div class="mk">{arrow(14)}</div>' if marked else ''
    lab = f'<div class="lb" style="color:{label_color}">{esc(label)}</div>' if label else ''
    return (f'<div class="rb">{img(it["icon"], 36, tint)}{mark}<div class="nf"><div class="nm">{esc(it["name"])}</div>'
            f'{lab}</div></div>')


def quest_frame(cells):
    return (f'<div class="frame qf"><div class="title">The Defias Brotherhood</div><div class="qbody">'
            f'<div style="font-size:16px;margin-bottom:6px">Rewards</div>'
            f'<div style="font-size:12.5px;margin-bottom:10px">You will be able to choose one of these rewards:</div>'
            f'<div class="rw">{"".join(cells)}</div>'
            f'<div style="font-size:12.5px;margin:12px 0 4px">You will also receive:</div>'
            f'<div style="font-size:12.5px">1{ux.GOLDC} 75{ux.SILVER}</div></div>'
            f'<div style="display:flex;justify-content:space-between;padding:0 10px 10px">'
            f'<span style="display:inline-flex;align-items:center;height:22px;padding:0 12px;font-size:12.5px;color:#ffd100;border-radius:3px;border:1px solid #6b5626;background:linear-gradient(#8a1c12,#4e0b06)">Complete Quest</span><span style="display:inline-flex;align-items:center;height:22px;padding:0 12px;font-size:12.5px;color:#ffd100;border-radius:3px;border:1px solid #6b5626;background:linear-gradient(#8a1c12,#4e0b06)">Cancel</span></div></div>')


WARRIOR = {i['id']: i for i in __import__('json').loads(
    (ux.BUILD / 'items/warrior.json').read_text())['items']}


def roll_frame(item_id, line_html, marked=True, items=None):
    it = (items or ux.ITEMS)[item_id]
    mark = f'<div style="position:absolute;top:-5px;right:-6px">{arrow(15)}</div>' if marked else ''
    buttons = ''.join(f'<span class="rbtn" title="{t}">{g}</span>' for g, t in
                      (('&#9856;', 'Need'), ('&#9673;', 'Greed'), ('&#10005;', 'Pass')))
    timer = '<div class="timer"><i></i></div>'
    return (f'<div class="frame rf">{timer}<div style="position:relative">{img(it["icon"], 36)}{mark}</div>'
            f'<div style="flex:1;min-width:0"><div style="color:{QUALITY[it["quality"]]};font-size:13px">'
            f'{esc(it["name"])}</div><div class="lb">{line_html}</div></div>'
            f'<div style="display:flex;gap:6px">{buttons}</div></div>')


def velisandre():
    return gear_lane.velisandre('raid')


def gain_over_worn(ch, item_id, slot):
    """The Gear lane's gain rule (Progress.slotGain) for an item against what the character wears there."""
    return round((ch.score(item_id, slot) - ch.score(ch.worn[slot], slot)) * ch.pp['dps'], 1)


ROLL_ITEM, ROLL_NOT_FOR_YOU = 17069, 16867  # Striker's Mark and Legplates of Might, both Molten Core


def board_rewards(labels=False):
    """With labels, each half names whose moment it is: Aldrenna leveling at 21, Velisandre at 60."""
    gain = ux.tunic_gain()
    roll_gain = gain_over_worn(velisandre(), ROLL_ITEM, 'ranged')
    best = 'Best for 20 to 29'
    quest = quest_frame([reward(2041, best, marked=True), reward(6087, unusable=True), reward(2042)])
    hover = tooltip(ux.tunic_of_westfall() + [
        l(f'{arrow()}Upgrade +{gain} DPS', color=C['green']),
        l(best, color=CLIENT_GOLD),
    ], width=250)
    none = quest_frame([reward(2041), reward(6087, unusable=True), reward(2042, 'Sells for most', label_color=GREY)])
    roll_pick = roll_frame(ROLL_ITEM, f'<span style="color:{C["green"]}">+{roll_gain} DPS</span>'
                                  f'<span style="color:{GREY}"> · </span>'
                                  f'<span style="color:{CLIENT_GOLD}">Best at level 60</span>')
    roll_up = roll_frame(ROLL_ITEM, f'<span style="color:{C["green"]}">+{roll_gain} DPS</span>')
    roll_no = roll_frame(ROLL_NOT_FOR_YOU, f'<span style="color:{GREY}">Not for you</span>', marked=False, items=WARRIOR)
    half = (lambda text: f'<div class="half">{text}</div>') if labels else (lambda text: '')
    html = row(
        col(half('Aldrenna, level 21'),
            note('<b>Aldrenna, level 21 hunter, hands in The Defias Brotherhood.</b> The arrow sits on the reward '
                 'with the biggest gain; the band\'s pick names itself in the item tooltip\'s own gold line, '
                 'word for word, so the label and the hover never disagree.'),
            row(quest, hover, gap=12),
            note('<b>Nothing here helps her</b> (the same quest, wearing better): one grey label on the reward that '
                 'sells for the most.'),
            none),
        col(half('Velisandre, level 60'),
            note('<b>Velisandre, level 60, at a group loot roll.</b> One short line beside the item: the gain, '
                 'then the pick\'s own words in the client gold.'),
            roll_pick, note('An upgrade that is not the pick'), roll_up, note('Plate a hunter cannot wear'),
            roll_no),
    )
    return board('<b>L10 · Quest reward and roll marks</b> · one wording with the item tooltip, one gold (the '
                 'client\'s #ffd100), the client\'s shadow under it', html)


# --- Player tooltip -------------------------------------------------------------------------------------

def player(name, cls, guild, race, score_line, tier_hex):
    return tooltip([
        f'<div class="t" style="color:{hub.CLASS_HEX[cls]}">{name}</div>', l(f'&lt;{guild}&gt;'),
        l(f'Level 60 {race} {cls.capitalize()} (Player)'), l('Alliance'),
        l(score_line, color=tier_hex),
    ], width=270)


def board_player():
    fresh = player('Kestrin', 'hunter', 'Thornwake Company', 'Night Elf', 'Forever Score 1240', QUALITY[3])
    stale = player('Orlaith', 'priest', 'Ashen Covenant', 'Human', 'Forever Score 1563 (6 Dec)', GREY)
    shift = tooltip([
        f'<div class="t" style="color:{hub.CLASS_HEX["hunter"]}">Kestrin</div>', l('&lt;Thornwake Company&gt;'),
        l('Level 60 Night Elf Hunter (Player)'), l('Alliance'),
        l('Forever Score 1240', color=QUALITY[3]),
        l('Gear', '410', color=GREY, rcolor='#fff'), l('Experience', '352', color=GREY, rcolor='#fff'),
        l('Parses', '268', color=GREY, rcolor='#fff'), l('Attendance', '150', color=GREY, rcolor='#fff'),
        l('Preparation', '60', color=GREY, rcolor='#fff'),
    ], width=270)
    html = row(
        col(note('<b>A record from today</b>: the score alone, in its tier\'s colour.'), fresh),
        col(note('<b>A record 4 days old</b> (Thu 10 Dec): past the 3 days every surface calls stale, so the line '
                 'carries its date and greys, as the chat badge greys and the Me tab warns, from the same rule.'), stale),
        col(note('<b>Shift held</b>: the five parts.'), shift),
    )
    return board('<b>L10 · Player tooltip</b> · one staleness rule, Score.isStale (3 days)', html)


# --- Paper doll and bags --------------------------------------------------------------------------------

# The rest of her bags on raid night, (icon, stack): the client's hearthstone, then simbuffs.json's consumables.
BAG_STOCK = [
    ('inv_misc_rune_01', 1),    # Hearthstone
    ('inv_potion_54', 5),       # Major Healing Potion
    ('inv_potion_76', 5),       # Major Mana Potion
    ('inv_potion_32', 4),       # Elixir of the Mongoose
    ('inv_potion_24', 2),       # Greater Fire Protection Potion, for Onyxia
    ('inv_potion_45', 3),       # Mageblood Potion
    ('inv_misc_fish_13', 10),   # Grilled Squid
    ('inv_drink_04', 5),        # Rumsey Rum Black Label
    ('inv_misc_food_47', 20),   # Tender Wolf Steak, her pet's
    ('inv_misc_fish_20', 8),    # Smoked Sagefish
    ('inv_stone_04', 1),        # Major Healthstone
]
BAG_SLOTS = 16
PANE_SLOTS = ('chest', 'hands', 'ranged')
EMPTY_SLOT = '<div class="slot" style="background:#0d0f14;box-shadow:inset 0 0 0 1px #2a2d36"></div>'


def slot_cell(item_id, marked):
    it = ux.item(item_id)
    mark = f'<div class="mk">{arrow(14)}</div>' if marked else ''
    return (f'<div class="slot">{img(it["icon"], 40, f"box-shadow:0 0 0 1px #000,0 0 0 2px {QUALITY[it["quality"]]}")}'
            f'{mark}</div>')


def stock_cell(icon, count):
    stack = f'<div class="ct">{count}</div>' if count > 1 else ''
    return f'<div class="slot">{img(icon, 40, "box-shadow:0 0 0 1px #000")}{stack}</div>'


# items/<class>.json carries no durability: the client's figure for each item drawn here.
DURABILITY = {13255: 40, 867: 40}
STAT_WORDS = {'agility': 'Agility', 'strength': 'Strength', 'stamina': 'Stamina', 'intellect': 'Intellect',
              'spirit': 'Spirit'}


def equip_lines(it):
    """The client's green Equip lines for the item's ratings and attack power; its own effect as Use."""
    stats, out = it['stats'], []
    if stats.get('hit'):
        out.append(f'Equip: Improves your chance to hit by {stats["hit"] / gear_lane.RATING_PER_PERCENT["hit"]:g}%.')
    if stats.get('crit'):
        out.append('Equip: Improves your chance to get a critical strike by '
                   f'{stats["crit"] / gear_lane.RATING_PER_PERCENT["crit"]:g}%.')
    if stats.get('attack_power'):
        out.append(f'Equip: +{stats["attack_power"]} Attack Power.')
    if it['effect_text']:
        out.append(f'Use: {it["effect_text"]}')
    return [l(esc(text), color=C['green']) for text in out]


def item_lines(it):
    """The client's tooltip head for a piece of armour from items/<class>.json."""
    bind = 'Binds when picked up' if it['bonding'] == 'pickup' else 'Binds when equipped'
    primaries = [l(f'+{it["stats"][k]} {word}') for k, word in STAT_WORDS.items() if it['stats'].get(k)]
    return [f'<div class="t" style="color:{QUALITY[it["quality"]]}">{esc(it["name"])}</div>', l(bind),
            l(gear_lane.SLOT_WORD[it['slot']], it['armor_type'].capitalize()), l(f'{it["armor"]} Armor'),
            *primaries, l(f'Durability {DURABILITY[it["id"]]} / {DURABILITY[it["id"]]}'),
            l(f'Requires Level {it["required_level"]}'), *equip_lines(it)]


def board_gear():
    ch = velisandre()
    hands = ch.doll()['hands']
    bag, worn, gain = hands['pick']['item_id'], hands['worn'], hands['gain']
    g, w = ux.item(bag), ux.item(worn)
    attribution = 'Forever Sixty · Beast Mastery, your sim'
    bag_line = (f'{arrow()}In your bags: <span style="color:{QUALITY[g["quality"]]}">{esc(g["name"])}</span> '
                f'<span style="color:{C["green"]}">+{gain} DPS</span>')
    slot_tip = tooltip(item_lines(w) + ['<div class="gap"></div>', l(bag_line), dim(attribution),
                                        dim('Alt-click: open Gear')], width=330)
    bag_tip = tooltip(item_lines(g) + ['<div class="gap"></div>',
                                       l(f'{arrow()}Upgrade +{gain} DPS over {esc(w["name"])}', color=C['green']),
                                       dim(attribution), dim('Alt-click: open Gear')], width=330)
    pane = (f'<div class="frame pane"><div class="small" style="color:{CLIENT_GOLD}">Character</div>'
            + ''.join(slot_cell(ch.worn[s], s == 'hands') for s in PANE_SLOTS) + '</div>')
    cells = [slot_cell(bag, True)] + [stock_cell(icon, count) for icon, count in BAG_STOCK]
    bags = f'<div class="frame bags">{"".join(cells)}{EMPTY_SLOT * (BAG_SLOTS - len(cells))}</div>'
    html = col(
        note(f'<b>Velisandre bought {esc(g["name"])} on the auction house.</b> The arrow is on the gloves in her bags '
             'and on her Hands slot. Shift on either hover ends with the marks\' one wording; an Alt-click on either '
             'opens the Gear tab on Hands, after the client\'s own click. In combat, or on an unmarked slot, Alt-click '
             'stays the client\'s.'),
        row(col(note('The Hands slot, Shift held'), row(pane, slot_tip, gap=12)),
            col(note('The bag item, Shift held'), row(bags, bag_tip, gap=12))),
    )
    return board('<b>L10 · Paper-doll and bag marks</b> · each leads to Gear: "Alt-click: open Gear"', html)


# --- Talent strip ---------------------------------------------------------------------------------------

def board_talents():
    strip = 'Next at 22: Lethal Attacks · 12 of 12'
    hover = tooltip([dim('Click: open Builds')], extra_style='position:absolute;left:12px;bottom:40px;min-width:0;'
                    'width:auto;font-family:Site,sans-serif;font-size:12px')
    node2 = (f'<div style="position:relative;width:40px;height:40px;flex:none">'
             f'{img("spell_frost_wizardmark", 32, "position:absolute;left:4px;top:4px")}'
             f'<div style="position:absolute;right:-6px;bottom:-5px;font-size:10.5px;padding:0 4px;border-radius:7px;'
             f'background:#ffd100;color:#1a1200">14</div></div>')
    window = (f'<div style="position:relative"><div class="frame tw"><div class="title">Talents</div>'
              f'<div style="display:flex;gap:26px;padding:16px 40px">{ux.talent_illustration()}{node2}</div>'
              f'<div class="tstrip">{strip}</div>'
              f'<span class="apply" style="display:inline-flex;align-items:center;height:22px;padding:0 12px;font-size:12.5px;color:#ffd100;border-radius:3px;border:1px solid #6b5626;background:linear-gradient(#8a1c12,#4e0b06)">Apply Changes</span></div>'
              f'<div style="position:absolute;left:0;top:150px;width:384px">{hover.replace("bottom:40px", "top:4px")}</div>'
              f'<div style="height:44px"></div></div>')
    html = col(note('<b>Aldrenna\'s talent window, the cursor on the strip.</b> The strip names her next point; '
                    'a click on it opens the Builds tab, and its hover says so in the same words as every mark '
                    'that opens a tab. Nothing of the client\'s window is touched; in combat the click does '
                    'nothing and the hover says nothing.'), window)
    return board('<b>L10 · Talent strip</b> · leads to Builds: "Click: open Builds"', html)


# --- Settings -------------------------------------------------------------------------------------------

COMMANDS = ['/fs: open the Forever Sixty window', '/fs copy: copy my character for the site',
            '/fs me, builds, gear, rotation, guild, officer: open that tab', "/fs link: copy a player's site page",
            '/fs tour: see the tour again', '/fs help: all the commands']
SECTIONS = [
    ('Tooltips and marks', [
        (False, 'Tooltips only', 'Turns off the line and every mark.'),
        (True, 'Gear tips on item tooltips',
         'Hold Shift for the details. Off also hides player scores and the quest and roll marks.'),
        (True, 'Marks on your bags and character pane',
         'Green arrow: beats what you wear. Gold corner: an upgrade to chase.'),
        (True, 'Marks on the talent window',
         'Numbers your build in order and rings the next talent.'),
        (True, 'Score badges in chat',
         'In LFG, Trade, guild chat and whispers.'),
    ]),
    ('The Forever Sixty line', [
        (True, 'Show the Forever Sixty line', 'Shift-drag moves it.'),
        (False, 'Hide the Forever Sixty line in combat',
         'Off: it fades to 25% instead.'),
        (False, 'Show the minimap button',
         'Left-click pastes or copies a build; right-click opens this page. Off: type /fs options or use '
         'Options &gt; AddOns.'),
    ]),
    ('Guild and the site', [
        (True, 'Share my build with my guild',
         'Guildmates see your level, spec, build and gear gap.'),
        (True, 'Show my Forever Score to others',
         "Applies after the companion's next upload."),
        (True, 'Save my character for the companion app',
         'Written at logout for the companion app.'),
        (False, 'Ask the group about rare drops too', 'As master looter, rares as well as epics.'),
    ]),
]


def command(line):
    token, words = line.split(': ', 1)
    return f'<div class="ln"><b>{esc(token)}:</b> {esc(words)}</div>'


def switch(on, label, hint):
    return (f'<div class="sw1"><div class="box{" on" if on else ""}"></div><div><div class="lbl">{label}</div>'
            f'<div class="hint">{hint}</div></div></div>')


def settings_page(warning=None):
    parts = ['<div class="ptt">Forever Sixty</div>']
    if warning:
        parts.append(f'<div class="ln wn" style="white-space:normal">{warning}</div>')
    parts.append('<div style="height:8px"></div><div class="phd">Commands</div><div style="height:4px"></div>')
    parts += [command(c) for c in COMMANDS]
    for heading, rows in SECTIONS:
        parts.append(f'<div style="height:16px"></div><div class="phd">{heading}</div>')
        parts += [switch(*r) for r in rows]
    parts.append('<div style="height:16px"></div><div class="rs"><div class="cb">Reset positions</div>'
                 f'<span style="color:{ADDON["muted"]}">Puts the Forever Sixty line, the minimap button and the '
                 'Forever Sixty window back in their default places.</span></div><div style="height:16px"></div>')
    ratings = ('Forever Scores: <span class="wn">too old (made 5 Dec)</span>' if warning
               else 'Forever Scores: current')
    parts.append('<div class="ln">Addon data: current for this patch</div>'
                 f'<div class="ln">{ratings}</div>')
    side = ''.join(f'<div class="{"on" if n == "Forever Sixty" else ""}">{n}</div>'
                   for n in ('Game', 'AddOns', 'Bartender4', 'Details!', 'Forever Sixty', 'Questie'))
    return (f'<div class="frame"><div class="title">Options</div><div class="sw"><div class="side">{side}</div>'
            f'<div class="cv"><div class="pg">{"".join(parts)}</div></div></div></div>')


def board_settings():
    html = row(
        col(note('<b>The page as the client shows it</b>, drawn whole: at 560 tall the canvas shows the commands '
                 'and the first section, and the wheel scrolls to the rest.'), settings_page()),
        col(note('<b>Forever Scores too old</b>: the whole sentence once, under the title in the warning colour; '
                 'the foot keeps only the state.'), settings_page('Forever Scores: too old (made 5 Dec). Get the new Forever Sixty '
                                                'Data at foreversixty.gg.')),
    )
    return board('<b>L10 · Settings</b> · the window\'s type (title face, cream headings, 12 px text, no rules), '
                 'switches only, Reset on the window\'s blue button', html)


BOARDS = {
    'l10-rewards': board_rewards,
    'l10-player': board_player,
    'l10-gear-marks': board_gear,
    'l10-talents': board_talents,
    'l10-settings': board_settings,
}


def main(names):
    from playwright.sync_api import sync_playwright
    OUT.mkdir(parents=True, exist_ok=True)
    with sync_playwright() as pw:
        exe = os.environ.get('PLAYWRIGHT_CHROMIUM')
        browser = pw.chromium.launch(executable_path=exe) if exe else pw.chromium.launch()
        ctx = browser.new_context(device_scale_factor=2, viewport={'width': 1800, 'height': 1200})
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
