"""Review boards for lane L3, the Gear tab as built (design/specs/2026-11-13-addon-hub.md 4.3).

    python3 design/mocks/addon_hub_l3.py              # every board
    python3 design/mocks/addon_hub_l3.py gear-hunter-60

Writes design/reviews/2026-11-13-addon-hub/l3/<board>.png at device scale 2, the window at its real 740 px width,
drawn to the built tab's own rules (addon/ForeverSixty/surfaces/GearTab.lua and GearTabView.lua), not the
approved board's:

- The picks are the shipped addon's own: addon/ForeverSixty/Data.lua `bis[spec][band][faction]`, decoded as
  Gear.decodePick does, re-picked under the source choice as Sources.resolve does. Data.lua carries one list per
  band (the bare character's: pipeline addonbis keeps no other preset), so the Conditions row of hub-gear.png has
  nothing to choose and the tab leaves it out.
- Item names, qualities, icons and stats: data/builds/1.60.1.70938/items/<class>.json.
- Gains: Progress.slotGain's rule ported: each stat over its rating per percent (Rules.RATING_PER_PERCENT) times
  the shipped weight (Data.lua weights), the ranged weapon's own DPS at the band's measured weapon term
  (Gear.weaponTerm), times the band's DPS per point (Data.lua dpsPerPoint).
- Stat weights: StatWeights' block, the Rotation tab's own (addon_stat_weights.py): the shipped weights per rating
  point over the top stat, attack power once.
- Source lines: ItemTooltip.sourceParts' rule; the drop chance only where Data.lua carries one.
- Source counts: Sources.counts over the band's published picks.
- Conditions (the gear-hit lane): at a band with a Raid-ready list (Data.lua `bisRaid`) the row shows, Raid-ready by
  default, and the doll, Upgrades and counts read that list (GearTab.withCondition); its hover is Data.lua
  `bisConditions`. The hit lines under the weights are GearTab.hitCap's over Data.lua `hitToCap`.

Every character here is made up: Velisandre, a level 60 Night Elf hunter, and Brannoc, a level 34 Orc warrior.
"""
import json
import os
import re
import sys

import addon_hub as hub
import addon_stat_weights as sw
from addon_hub import ADDON, CLIENT_GOLD, QUALITY, esc, section, window
import addon_ux as ux
from addon_ux import ROOT, icon_uri

OUT = ROOT / 'design/reviews/2026-11-13-addon-hub/l3'
BUILD = ROOT / 'data/builds/1.60.1.70938'
DATA_LUA = (ROOT / 'addon/ForeverSixty/Data.lua').read_text()

RATING_PER_PERCENT = {'hit': 10, 'crit': 14, 'expertise': 10, 'dodge': 12, 'parry': 15, 'block': 5, 'defense': 1,
                      'haste': 10}
KINDS = [('dungeon', 'Dungeon'), ('raid', 'Raid'), ('world', 'World / BoE'), ('quest', 'Quest'),
         ('crafted', 'Crafted'), ('rep', 'Reputation'), ('vendor', 'Vendor'), ('pvp', 'PvP')]
KIND_OF_CODE = {'D': 'dungeon', 'A': 'raid', 'W': 'world', 'B': 'world', 'Q': 'quest', 'C': 'crafted', 'R': 'rep',
                'V': 'vendor', 'P': 'pvp'}
KIND_NAME = {'Q': 'Quest', 'V': 'Vendor', 'D': 'Dungeon', 'C': 'Crafted', 'R': 'Reputation', 'P': 'PvP', 'W': 'World',
             'A': 'Raid', 'B': 'World Drop'}
STANDING = {'n': 'Neutral', 'f': 'Friendly', 'h': 'Honored', 'r': 'Revered', 'e': 'Exalted'}
STANDING_RANK = {'n': 0, 'f': 1, 'h': 2, 'r': 3, 'e': 4}
SLOT_WORD = {'head': 'Head', 'neck': 'Neck', 'shoulder': 'Shoulder', 'back': 'Back', 'chest': 'Chest', 'wrist': 'Wrist',
             'hands': 'Hands', 'waist': 'Waist', 'legs': 'Legs', 'feet': 'Feet', 'finger1': 'Finger',
             'finger2': 'Finger', 'trinket1': 'Trinket', 'trinket2': 'Trinket', 'main_hand': 'Main Hand',
             'off_hand': 'Off Hand', 'ranged': 'Ranged'}
RIGHT_ROOM = 258
LEFT = ['head', 'neck', 'shoulder', 'back', 'chest', 'wrist', 'main_hand', 'off_hand', 'ranged']
RIGHT = ['hands', 'waist', 'legs', 'feet', 'finger1', 'finger2', 'trinket1', 'trinket2']
PAIR_FIRST = {'finger2': 'finger1', 'trinket2': 'trinket1', 'off_hand': 'main_hand'}


def lua_block(start, key):
    i = DATA_LUA.index(key, DATA_LUA.index(start))
    depth, j = 0, DATA_LUA.index('{', i)
    for k in range(j, len(DATA_LUA)):
        depth += {'{': 1, '}': -1}.get(DATA_LUA[k], 0)
        if depth == 0:
            return DATA_LUA[j:k + 1]
    raise ValueError(key)


PLACES = re.findall(r'"((?:[^"\\]|\\.)*)"', lua_block('\tbisSources = {', 'bisSources'))


def picks_of(spec, band, faction, table='bis'):
    spec_block = lua_block(f'\t{table} = {{', f'["{spec}"] = {{')
    band_block = spec_block[spec_block.index(f'[{band}] = {{'):]
    fac = band_block[band_block.index(f'{faction} = {{'):]
    fac = fac[:fac.index('}')]
    return dict(re.findall(r'(\w+) = "([^"]*)"', fac))


def has_raid(spec, band, faction):
    try:
        return bool(picks_of(spec, band, faction, 'bisRaid'))
    except ValueError:
        return False


def cap_percent(value):
    return f'{value:g}' if value == int(value) else f'{value:.1f}'


def hit_type(spec):
    """Rules.hitType: melee, a hunter at range, or a caster."""
    melee_specs = {'paladin-protection', 'paladin-retribution', 'shaman-enhancement', 'hunter-survival'}
    cls = spec.split('-')[0]
    if cls in ('warrior', 'rogue') or spec in melee_specs or spec.startswith('druid-feral'):
        return 'melee'
    return 'ranged' if cls == 'hunter' else 'spell'


def hit_lines(spec, band, faction, condition):
    """GearTab.hitCap's lines over Data.lua's hitToCap and its hitCaps pool."""
    pool = re.findall(r'^\t(\{.*\}),$', lua_block('local hitCaps = {', 'hitCaps'), re.M)
    try:
        block = lua_block('\thitToCap = {', f'["{spec}"] = {{')
    except ValueError:
        return []
    preset = block[block.index(f'{condition} = {{'):] if f'{condition} = {{' in block else ''
    preset = preset[:preset.index('\n\t\t\t},')] if preset else ''
    found = re.search(rf'\[{band}\] = \{{[^}}]*?{faction} = hitCaps\[(\d+)\]', preset)
    if not found:
        return []
    entry = pool[int(found.group(1)) - 1]
    parts = {k: (float(a), float(b)) for k, a, b in re.findall(r'(hit|white|spell) = \{ ([\d.]+), ([\d.]+) \}', entry)}
    school = re.search(r'school = \{ ([\d.]+), "([^"]+)" \}', entry)
    if 'hit' in parts and hit_type(spec) == 'ranged':
        # A ranged hunter's priority casts nothing that rolls spell hit (GearTab.hitCap).
        parts.pop('spell', None)
    weapon_word = 'Ranged hit' if hit_type(spec) == 'ranged' else 'Melee hit'
    lines = []
    for key, word in (('hit', weapon_word), ('white', 'White swings'), ('spell', 'Spell hit')):
        if key in parts:
            to, cap = parts[key]
            lines.append(f'{word}: at the {cap_percent(cap)}% cap' if to <= 0
                         else f'{word}: {cap_percent(to)}% to the {cap_percent(cap)}% cap')
    if school and 'spell' in parts:
        lines.append(f'{school.group(2)} spells: {cap_percent(float(school.group(1)))}% to the '
                     f'{cap_percent(parts["spell"][1])}% cap')
    return lines


def conditions_of(spec):
    block = lua_block('\tbisConditions = {', 'bisConditions')
    line = re.search(rf'\["{re.escape(spec)}"\] = \{{ (.*) \}},', block)
    return dict(re.findall(r'(\w+) = "((?:[^"\\]|\\.)*)"', line.group(1)))


def weights_of(spec):
    line = re.search(rf'\["{re.escape(spec)}"\] = \{{([^}}]*)\}}', lua_block('\tweights = {', 'weights'))
    return {k: float(v) for k, v in re.findall(r'\["(\w+)"\] = ([\d.]+)', line.group(1))}


def per_point(spec, band):
    block = lua_block('\tdpsPerPoint = {', f'["{spec}"] = {{')
    found = re.search(rf'\[{band}\] = \{{ reference = "(\w+)", dps = ([\d.]+), weaponDps = ([\d.]+)', block)
    return {'reference': found.group(1), 'dps': float(found.group(2)), 'weaponDps': float(found.group(3))}


def decode(packed):
    """Gear.decodePick: the pick, its alternatives, the tail's head and by-kind entries."""
    main, _, tail = packed.partition('||') if packed.count('|') > 3 else (packed, '', '')
    fields = packed.split('|')
    pick = {'item_id': int(fields[0]), 'kind': fields[1], 'source': fields[2] or None,
            'chance': float(fields[3]) if len(fields) > 3 and fields[3] else None, 'alternatives': [],
            'head': None, 'byKind': []}
    alts = fields[4:]
    k = 0
    while k + 2 < len(alts) and alts[k]:
        pick['alternatives'].append({'item_id': int(alts[k]), 'kind': alts[k + 1], 'source': alts[k + 2]})
        k += 3
    tail = packed.split('||', 1)[1].split('||', 1)[1] if packed.count('||') >= 2 else ''
    for entry in [e for e in tail.split(';') if e]:
        if entry.startswith('*'):
            body, _, rest = entry[1:].partition(',')
            kinds = re.findall(r'([A-Z])([a-z]?)', body.replace('!', ''))
            pick['head'] = {'kinds': [(KIND_OF_CODE[c], s or None) for c, s in kinds],
                            'notWith': [int(x) for x in rest.split('+')] if rest else []}
            if pick['head']['kinds'] and pick['head']['kinds'][0][0] == KIND_OF_CODE.get(pick['kind']):
                pick['standing'] = pick['head']['kinds'][0][1]
            continue
        m = re.match(r'([A-Z])([a-z]?)(!?)(\d+),(\d+)(?:,([\d+]+))?$', entry)
        code, standing, _, item_id, place, not_with = m.groups()
        pick['byKind'].append({'item_id': int(item_id), 'kind': code, 'filterKind': KIND_OF_CODE[code],
                               'standing': standing or None, 'source': PLACES[int(place) - 1],
                               'notWith': [int(x) for x in not_with.split('+')] if not_with else []})
    return pick


def allowed(on, rep_cap, kind, standing):
    if kind not in on:
        return False
    if rep_cap is None:
        return True
    if kind == 'rep':
        return standing is not None and STANDING_RANK[standing] <= STANDING_RANK[rep_cap]
    return standing is None or STANDING_RANK[standing] <= STANDING_RANK[rep_cap]


def resolve(pick, on, rep_cap, mate):
    """Sources.resolve: the pick while one of its kinds is on, else the best by-kind entry allowed."""
    kinds = pick['head']['kinds'] if pick['head'] and pick['head']['kinds'] else [(KIND_OF_CODE[pick['kind']], None)]
    blocked = mate is not None and pick['head'] and mate in pick['head']['notWith']
    if any(allowed(on, rep_cap, k, s) for k, s in kinds) and not blocked:
        return pick
    for entry in pick['byKind']:
        if allowed(on, rep_cap, entry['filterKind'], entry['standing']) and mate not in entry['notWith']:
            return {'item_id': entry['item_id'], 'kind': entry['kind'], 'source': entry['source'],
                    'standing': entry['standing'], 'chance': None}
    return None


def counts(picks):
    out = {k: 0 for k, _ in KINDS}
    for pick in picks.values():
        found = {k for k, _ in (pick['head']['kinds'] if pick['head'] and pick['head']['kinds'] else
                                [(KIND_OF_CODE[pick['kind']], None)])}
        found |= {e['filterKind'] for e in pick['byKind']}
        found |= {KIND_OF_CODE[a['kind']] for a in pick['alternatives']}
        for k in found:
            out[k] += 1
    return out


class Character:
    def __init__(self, cls, spec, band, faction, worn, bags, unenchanted, on=None, rep_cap=None, condition=None):
        self.items = {it['id']: it for it in json.loads((BUILD / f'items/{cls}.json').read_text())['items']}
        self.cls, self.spec, self.band, self.faction = cls, spec, band, faction
        self.worn, self.bags, self.unenchanted = worn, bags, unenchanted
        self.on = on or [k for k, _ in KINDS]
        self.rep_cap = rep_cap
        self.weights = weights_of(spec)
        self.pp = per_point(spec, band)
        self.raid = has_raid(spec, band, faction)
        self.condition = (condition or 'raid') if self.raid else 'bare'
        table = 'bisRaid' if self.condition == 'raid' else 'bis'
        self.packed = {slot: decode(p) for slot, p in picks_of(spec, band, faction, table).items()}

    def score(self, item_id, slot):
        it = self.items[item_id]
        total = sum(v / RATING_PER_PERCENT.get(k, 1) * self.weights.get(k, 0) for k, v in it['stats'].items())
        weapon_slot = 'ranged' if self.pp['reference'] == 'ranged_attack_power' else 'main_hand'
        if slot == weapon_slot and it['dps']:
            total += it['dps'] * self.pp['weaponDps'] / self.pp['dps']
        return total

    def picks(self):
        out = {}
        for slot in LEFT + RIGHT:
            pick = self.packed.get(slot)
            if pick is None:
                continue
            mate = out.get(PAIR_FIRST.get(slot))
            out[slot] = resolve(pick, self.on, self.rep_cap, mate['item_id'] if mate else None)
        return {s: p for s, p in out.items() if p is not None}

    def doll(self):
        picks, cells = self.picks(), {}
        two_hand = self.worn.get('main_hand') and self.items[self.worn['main_hand']].get('two_hand')
        for slot in LEFT + RIGHT:
            if slot == 'off_hand' and two_hand:
                continue
            pick, worn = picks.get(slot), self.worn.get(slot)
            cell = {'slot': slot, 'worn': worn, 'pick': pick, 'noEnchant': slot in self.unenchanted and worn}
            if pick:
                cell['holds'] = worn == pick['item_id']
                base = self.score(worn, slot) if worn else 0
                cell['gain'] = 0 if cell['holds'] else round((self.score(pick['item_id'], slot) - base) * self.pp['dps'], 1)
                cell['inBags'] = not cell['holds'] and pick['item_id'] in self.bags
            cells[slot] = cell
        return cells


# --- drawing, to GearTabView's layout -------------------------------------------------------------------------

def item_icon(ch, item_id, size):
    it = ch.items[item_id]
    return (f'<img class="ic" src="{icon_uri(it["icon"])}" style="width:{size}px;height:{size}px;'
            f'box-shadow:0 0 0 1px {QUALITY[it["quality"]]}">')


def slot_icon(size):
    return f'<span class="ic" style="width:{size}px;height:{size}px;background:{ADDON["inset"]};display:block"></span>'


def name_of(ch, item_id, room=None):
    it = ch.items[item_id]
    name = it['name'] if room is None else word_cut(it['name'], lambda t: len(t) * 6.0 <= room) or it['name']
    return f'<span style="color:{QUALITY[it["quality"]]}">{esc(name)}</span>'


def status(cell):
    parts = []
    if cell.get('holds'):
        parts.append('<span class="m">Best in slot</span>')
    elif cell.get('gain', 0) > 0:
        gain = f'<span class="ok">+{cell["gain"]:.1f} DPS</span>'
        if cell.get('inBags'):
            gain += ' <span class="cg">in your bags</span>'
        parts.append(gain)
    elif cell.get('pick'):
        parts.append('<span class="m">Better than the list</span>')
    if cell['noEnchant']:
        parts.append('<span class="wa">No enchant</span>')
    if not cell['worn'] and not parts:
        parts.append('<span class="m">Empty</span>')
    return '<span class="m"> · </span>'.join(parts)


def doll_cell(ch, cell, mirrored):
    icon = item_icon(ch, cell['worn'], 28) if cell['worn'] else slot_icon(28)
    name = name_of(ch, cell['worn'], 157) if cell['worn'] else f'<span class="m">{SLOT_WORD[cell["slot"]]}</span>'
    return (f'<div class="cell{" r" if mirrored else ""}">{icon}<div class="tx"><div class="one">{name}</div>'
            f'<div class="one">{status(cell)}</div></div></div>')


def source_parts(pick):
    source, kind = pick.get('source'), pick['kind']
    if kind == 'C':
        return 'Crafted', source
    if kind in ('D', 'A') and source and ': ' in source:
        return tuple(source.split(': ', 1))
    standing = STANDING.get(pick.get('standing'))
    if standing:
        return source, standing
    if kind in ('W', 'V') or KIND_NAME[kind].lower() == source.lower():
        return source, None
    return source, KIND_NAME[kind]


def word_cut(text, fits):
    if fits(text):
        return text
    words = text.split(' ')
    for count in range(len(words) - 1, 0, -1):
        cut = ' '.join(words[:count]).rstrip(' ·,:') + '…'
        if fits(cut):
            return cut
    return None


def where_text(cell, room=None, per_char=6.0):
    """GearTab.whereLine: whole; else the slot gives way; else the chance; else the place's last words."""
    slot = SLOT_WORD[cell['slot']]
    if cell.get('inBags'):
        return f'{slot} · <span class="cg">In your bags</span>'
    pick = cell['pick']
    first, second = source_parts(pick)
    chance = f'{max(1, int(pick["chance"] * 100 + 0.5))}% drop' if pick.get('chance') else None

    def text(place, with_chance):
        return ' · '.join(p for p in (place, second, chance if with_chance else None) if p)

    def fits(t):
        return room is None or len(t) * per_char <= room

    whole = f'{slot} · {text(first, True)}'
    if fits(whole):
        return esc(whole)
    for with_chance in ((True, False) if chance else (False,)):
        if fits(text(first, with_chance)):
            return esc(text(first, with_chance))
    tail = f' · {second}' if second else ''
    place = word_cut(first, lambda c: fits(c + tail))
    return esc(place + tail) if place else None


def where(ch, cell):
    return f'<div class="m one">{where_text(cell, RIGHT_ROOM) or SLOT_WORD[cell["slot"]]}</div>'


def more_where(ch, cell):
    room = 300 - len(ch.items[cell['pick']['item_id']]['name'] + ' · ') * 6.0
    w = where_text(cell, room)
    return f'<span class="m"> · {w}</span>' if w else ''


def heading(title, note=''):
    r = f'<span class="r">{note}</span>' if note else ''
    return (f'<div class="sh" style="font-weight:400"><span style="font-family:Title,serif;font-size:14px;'
            f'color:{ADDON["cream"]}">{title}</span>{r}</div>')


def upgrades_block(ch, cells):
    found = [c for c in cells.values() if c.get('pick') and not c.get('holds') and c.get('gain', 0) > 0]
    total = sum(c['gain'] for c in found)
    found.sort(key=lambda c: -c['gain'])
    rows = ''.join(
        f'<div class="up">{item_icon(ch, c["pick"]["item_id"], 24)}<div class="tx"><div class="row">'
        f'<span class="one" style="flex:1">{name_of(ch, c["pick"]["item_id"])}</span>'
        f'<span class="ok">+{c["gain"]:.1f} DPS</span></div>{where(ch, c)}</div></div>'
        for c in found[:4])
    # GearTabView: no total without an upgrade, and the empty line in its place.
    note = f"+{total:.1f} DPS in all" if found else ''
    rows = rows or '<div class="m">Nothing on this list beats what you wear.</div>'
    return f'<div class="sec">{heading("Upgrades", note)}{rows}</div>', found[4:]


def weights_block(ch, sim):
    """StatWeights' block (the Rotation tab's own, addon_stat_weights) with GearTab.hitCap's lines under it."""
    m = sw.model(ch.weights, ch.pp['dps'], sw.L('weightsSim') if sim else '', ch.pp['reference'])
    head, body = sw.html(m)
    hit = ''.join(f'<div class="m" style="height:18px">{line}</div>'
                  for line in hit_lines(ch.spec, ch.band, ch.faction, ch.condition))
    hit = f'<div style="margin-top:6px">{hit}</div>' if hit else ''
    return f'<div class="sec">{head}{body}{hit}</div>'


def checkbox(on):
    if on:
        return (f'<span style="width:16px;height:16px;background:{ADDON["gold"]};display:inline-flex;align-items:center;'
                f'justify-content:center;color:{ADDON["bg"]};font-size:12px;font-weight:700;flex:none">✓</span>')
    return f'<span style="width:14px;height:14px;border:1px solid {ADDON["muted"]};display:inline-block;flex:none"></span>'


def sources_block(ch, words, off, saved):
    c = counts(ch.packed)
    on = set(ch.on)
    rows = ''
    for kind, label in KINDS:
        checked = kind in on and not (kind == 'rep' and ch.rep_cap)
        colour = ADDON['muted'] if c[kind] == 0 else ADDON['body']
        rows += (f'<div class="k">{checkbox(checked)}<span style="color:{colour}">{label}</span>'
                 f'<span class="c">{c[kind]}</span></div>')
    foot = ''
    off = None if ch.rep_cap else off
    if off or saved:
        link = '<span class="link">Default</span>' if saved else ''
        foot = f'<div class="m" style="margin-top:4px">{off}{" · " if off and saved else ""}{link}</div>'
    return f'<div class="sec">{heading("Sources", "" if words == "All sources" else words)}<div class="kinds">{rows}</div>{foot}</div>'


def board(ch, caption, title, identity, band_on, words, off=None, saved=False, sim=False, tips=()):
    cells = ch.doll()
    bands = ''.join(f'<span class="btn{" sel" if b == band_on else ""}" style="width:32px;justify-content:center;'
                    f'padding:0">{b}</span>' for b in (20, 30, 40, 50, 60))
    left_cells = [cells[s] for s in LEFT if s in cells]
    right_cells = [cells[s] for s in RIGHT if s in cells]
    conditions = ''
    if ch.raid:
        conditions = ''.join(f'<span class="btn{" sel" if c == ch.condition else ""}">{w}</span>'
                             for c, w in (('raid', 'Raid-ready'), ('bare', 'Bare')))
        conditions = (f'<div class="bands" style="margin-bottom:12px"><span class="m" style="width:62px">Conditions'
                      f'</span>{conditions}</div>')
    left = (f'<div class="bands" style="margin-bottom:{10 if ch.raid else 12}px"><span class="m" style="width:62px">'
            f'Level</span>{bands}</div>{conditions}'
            f'<div class="doll"><div class="dcol" style="width:191px">{"".join(doll_cell(ch, c, False) for c in left_cells)}</div>'
            f'<div class="dcol" style="width:191px">{"".join(doll_cell(ch, c, True) for c in right_cells)}</div></div>')
    up_html, rest = upgrades_block(ch, cells)
    unench = [SLOT_WORD[s] for s in LEFT + RIGHT if s in cells and cells[s]['noEnchant']]
    if unench:
        left += f'<div class="wa" style="margin-top:4px">No enchant: {", ".join(unench)}</div>'
    if rest:
        more = ''.join(
            f'<div class="row" style="height:18px">{item_icon(ch, c["pick"]["item_id"], 16)}'
            f'<span class="one" style="flex:1">{name_of(ch, c["pick"]["item_id"])}{more_where(ch, c)}</span>'
            f'<span class="ok">+{c["gain"]:.1f} DPS</span></div>' for c in rest[:6])
        left += f'<div class="sec" style="margin-top:20px">{heading("More upgrades")}{more}</div>'
    right = up_html + weights_block(ch, sim) + sources_block(ch, words, off, saved)
    scored = [c for c in cells.values() if c.get('pick')]
    best = sum(1 for c in scored if c.get('holds') or c.get('gain', 0) <= 0)
    missing = sum(1 for c in cells.values() if c['noEnchant'])
    figures = [f'{best} of {len(scored)} best in slot']
    if missing:
        figures.append('1 enchant missing' if missing == 1 else f'{missing} enchants missing')
    html = window('Gear', ' · '.join([identity] + figures), left, right, ('Esc to close', 'Shift-click an item to link it'),
                  cls=ch.cls, title=title, dots=('Gear',))
    return hub.board2(caption, html + (f'<div class="row" style="margin-top:18px">{"".join(tips)}</div>' if tips else ''))


def hover(title, groups, tag):
    """A text hover as Theme.showLines draws it: the title, then muted headings over wrapped lines."""
    lines = [f'<div class="t">{title}</div>']
    for heading_text, body in groups:
        if heading_text:
            lines.append(f'<div class="gap"></div><div style="color:{ADDON["muted"]}">{heading_text}</div>')
        lines.append(f'<div style="white-space:normal">{body}</div>')
    return ux.tooltip(lines, tag=tag, width=300)


def conditions_tips(ch):
    entry = conditions_of(ch.spec)
    return [
        hover(entry['label'], [('Buffs', entry['buffs']), ('Debuffs', entry['debuffs']),
                               ('Consumables', entry['consumes'])], 'Hover: Raid-ready'),
        hover('Bare', [(None, 'Bare character: no raid buffs, debuffs or consumables.')], 'Hover: Bare'),
        hover('Hit cap', [(None, 'Hit is worth its full weight up to the cap and nothing past it. An item\'s '
                                 '"Equip: Improves your chance to hit by 1.0%." counts 1% toward it, 10 hit rating is '
                                 '1%. Counted with this list worn.')], 'Hover: the hit lines'),
    ]


def velisandre(condition='bare'):
    # Level 60 Night Elf hunter: the bare list's picks but for a few older pieces, the list's gloves in her bags.
    worn = {s: p['item_id'] for s, p in Character('hunter', 'hunter-beast-mastery', 60, 'alliance', {}, set(),
                                                     set(), condition='bare').picks().items()}
    worn.update({'back': 22269, 'neck': 22403, 'wrist': 13211, 'hands': 13255, 'feet': 14616, 'finger2': 13098,
                 'main_hand': 13052})
    return Character('hunter', 'hunter-beast-mastery', 60, 'alliance', worn, {271948}, {'chest', 'feet'},
                     condition=condition)


def torvald(condition):
    # Level 60 Dwarf warrior (Fury, dual wield): wears the bare list, its legs without an enchant.
    worn = {s: p['item_id'] for s, p in Character('warrior', 'warrior-fury', 60, 'alliance', {}, set(), set(),
                                                     condition='bare').picks().items()}
    return Character('warrior', 'warrior-fury', 60, 'alliance', worn, set(), {'legs'}, condition=condition)


def brannoc():
    # Level 34 Orc warrior (Fury): two one-handers, levelling gear; the leveling default (no rep above Friendly).
    ch = Character('warrior', 'warrior-fury', 30, 'horde', {}, set(), {'legs'},
                   on=[k for k, _ in KINDS if k != 'rep'], rep_cap='f')
    picks = ch.picks()
    worn = {s: p['item_id'] for s, p in picks.items() if s in ('chest', 'waist', 'finger1', 'neck', 'ranged')}
    alts = {s: ch.packed[s]['alternatives'][0]['item_id'] for s in ('head', 'shoulder', 'legs', 'feet', 'wrist', 'back')
            if ch.packed[s]['alternatives'] and ch.packed[s]['alternatives'][0]['item_id'] in ch.items}
    worn.update(alts)
    worn['main_hand'], worn['off_hand'] = 7752, 13033
    ch.worn = {s: i for s, i in worn.items() if i}
    return ch


BOARDS = {
    'gear-hunter-60-raid': lambda: board(
        velisandre('raid'), '<b>Gear, level 60, Raid-ready (the gear-hit lane)</b> · the Conditions row opens on '
        'Raid-ready: the doll, Upgrades and the counts read the list ranked under raid buffs and consumables; under '
        'the weights the hit cap of that list. The three hovers below are the row\'s and the hit lines\'.',
        'Velisandre', 'Level 60 Hunter', 60, 'All sources', tips=conditions_tips(velisandre('raid'))),
    'gear-warrior-60-bare': lambda: board(
        torvald('bare'), '<b>Gear, level 60, Bare (the gear-hit lane)</b> · a dual-wield Fury warrior who wears the '
        'bare list, Bare chosen: every slot best in slot, the hit lines for specials and white swings.',
        'Torvald', 'Level 60 Warrior', 60, 'All sources'),
    'gear-hunter-60': lambda: board(
        velisandre('bare'), '<b>Gear, level 60, as built</b> · the doll at her band against what she wears, the list\'s '
        'gloves in her bags; Upgrades with the drop chance only where the data has one; the shipped weights; '
        'All sources.', 'Velisandre', 'Level 60 Hunter', 60, 'All sources'),
    'gear-warrior-34': lambda: board(
        brannoc(), '<b>Gear, level 34, as built</b> · two one-handers, so the off hand shows; no gloves or trinkets yet; band 30; '
        'the leveling default keeps reputation up to Friendly, so the rep ring re-picks to its quest ring.',
        'Brannoc', 'Level 34 Warrior', 30, 'No rep above Friendly', 'Off: Reputation above Friendly'),
}


def main(names):
    from playwright.sync_api import sync_playwright
    OUT.mkdir(parents=True, exist_ok=True)
    with sync_playwright() as pw:
        exe = os.environ.get('PLAYWRIGHT_CHROMIUM')
        browser = pw.chromium.launch(executable_path=exe) if exe else pw.chromium.launch()
        page = browser.new_context(device_scale_factor=2, viewport={'width': 1400, 'height': 1000}).new_page()
        for name in names:
            page.set_content(BOARDS[name]())
            page.evaluate('document.fonts.ready')
            page.wait_for_timeout(150)
            page.locator('#board').screenshot(path=str(OUT / f'{name}.png'))
            print(OUT / f'{name}.png')
        browser.close()


if __name__ == '__main__':
    main(sys.argv[1:] or list(BOARDS))
