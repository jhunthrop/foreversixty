"""Mock boards for the addon source filter (design/specs/2026-11-13-addon-source-filter.md, section 11).

    python3 design/mocks/addon_sources.py                    # every board
    python3 design/mocks/addon_sources.py panel-sources-leveling

Writes design/reviews/2026-11-13-addon-sources/<board>.png through addon_ux.py's Playwright render
(device scale 2), on addon_ux.py's own plates, tooltips and tokens.

Sources, so nothing is invented:
- Picks, by_kind entries, standings, places and dps_delta: data/builds/1.60.1.70790/bis/hunter-beast-mastery.json,
  band 20 bare alliance (Bow Jackzon, level 24) and band 60 raid alliance.
- Item names, qualities and icons: data/builds/1.60.1.70790/items/hunter.json and icons/*.webp.
- Re-picks: resolve() below, the twin of web/src/lib/bis/source-filter.ts filteredBand and doll-view.ts
  cappedSlot (the leveling default keeps standings up to Friendly; a pair's second slot steers clear of the
  first slot's item through not_with, then next).
- Counts: slots of the band whose pick kinds, by_kind keys or alternatives' kinds include the kind
  (doll-view.ts sourceCounts).
- Every DPS figure (Upgrades rows, Next lines, the Shift set-aside gap) is priced the way the addon prices it
  in game: the two items' stats from items/hunter.json times the band's dps_per_point (the bis file's
  `weights`). The worn item a gain is counted against is named per slot by worn_for(): the best item for the
  slot by that pricing that a hunter of the band's character wears (quality up to uncommon, a required level
  ten under the character's, cloth, leather or no armour below 40) that the upgrade beats by 0.3 DPS or
  more. Weapons (main hand, ranged) are left off
  the Upgrades rows: their own DPS is outside the stat weights, so the stat pricing would understate them.
- Drop chances: the published pick's own `chance`; a re-picked by_kind item carries none.
"""
import sys

import addon_ux as ux
from addon_ux import ADDON, C, QUALITY, BIS, ITEMS, img, l, medallion, page, tooltip

OUT = ux.ROOT / 'design/reviews/2026-11-13-addon-sources'
ux.OUT = OUT

KINDS = ['dungeon', 'raid', 'world', 'quest', 'crafted', 'rep', 'vendor', 'pvp']
LABEL = {'dungeon': 'Dungeon', 'raid': 'Raid', 'world': 'World / BoE', 'quest': 'Quest', 'crafted': 'Crafted',
         'rep': 'Reputation', 'vendor': 'Vendor', 'pvp': 'PvP'}
FILTER_KIND = {'dungeon': 'dungeon', 'raid': 'raid', 'world': 'world', 'world_drop': 'world', 'zone': 'world',
               'quest': 'quest', 'crafted': 'crafted', 'rep': 'rep', 'vendor': 'vendor', 'pvp': 'pvp'}
KIND_WORD = {'quest': 'Quest', 'vendor': 'Vendor', 'crafted': 'Crafted', 'rep': 'Reputation', 'pvp': 'PvP',
             'world': 'World', 'world_drop': 'World drop'}
STANDING = {'neutral': 0, 'friendly': 1, 'honored': 2, 'revered': 3, 'exalted': 4}
PAIR_FIRST = {'finger2': 'finger1', 'trinket2': 'trinket1', 'off_hand': 'main_hand'}
SLOT_WORD = {'head': 'Head', 'neck': 'Neck', 'shoulder': 'Shoulder', 'back': 'Back', 'chest': 'Chest',
             'wrist': 'Wrist', 'hands': 'Hands', 'waist': 'Waist', 'legs': 'Legs', 'feet': 'Feet',
             'finger1': 'Finger', 'finger2': 'Finger', 'trinket1': 'Trinket', 'trinket2': 'Trinket',
             'main_hand': 'Main Hand', 'off_hand': 'Off Hand', 'ranged': 'Ranged'}
LEVELING_DEFAULT = [k for k in KINDS if k != 'rep']
GOLD, MUTED, BODY = ADDON['gold'], ADDON['muted'], ADDON['body']


def band(level, preset):
    return next(b for b in BIS['bands'] if b['band'] == level and b['preset'] == preset and b['faction'] == 'alliance')


B20, B60 = band(20, 'bare'), band(60, 'raid')


def above_cap(standing):
    return STANDING.get(standing or '', 99) > STANDING['friendly']


def pick_kinds(row):
    return row.get('kinds') or ([FILTER_KIND[row['source_kind']]] if row.get('source_kind') in FILTER_KIND else [])


def capped_row(row):
    """doll-view.ts cappedSlot: reputation above Friendly taken out of the pick and its by_kind."""
    def capped(kind):
        st = (row.get('kind_standings') or {}).get(kind) or (row.get('standing') if kind == 'rep' else None)
        return above_cap(st) if kind == 'rep' else (st is not None and above_cap(st))

    by_kind = {}
    for kind, entry in (row.get('by_kind') or {}).items():
        node = entry
        while node is not None and (above_cap(node.get('standing')) if kind == 'rep'
                                    else node.get('standing') is not None and above_cap(node.get('standing'))):
            node = node.get('next')
        if node is not None:
            by_kind[kind] = node
    kinds = pick_kinds(row)
    rest = [k for k in kinds if not capped(k)]
    out = dict(row, by_kind=by_kind)
    if len(rest) < len(kinds):
        out = dict(out, kinds=rest) if rest else dict(out, item_id=0)
    return out


def resolve(rows, on, rep_cap=False):
    """source-filter.ts filteredBand: slot -> (item_id, source, source_kind, dps_delta, standing) or None."""
    allowed = set(on) | ({'rep'} if rep_cap else set())
    out = {}
    for raw in rows:
        if 'item_id' not in raw:
            continue
        row = capped_row(raw) if rep_cap else raw
        slot = row['slot']
        mate = out.get(PAIR_FIRST.get(slot)) if slot in PAIR_FIRST else None
        mate_id = mate['item_id'] if mate else None
        blocked = mate_id is not None and mate_id in (row.get('not_with') or [])
        if row.get('item_id') and not blocked and any(k in allowed for k in pick_kinds(row)):
            out[slot] = {'item_id': row['item_id'], 'source': row.get('source', ''), 'kind': row['source_kind'],
                         'delta': 0, 'standing': row.get('standing'), 'repicked': False,
                         'chance': row.get('chance')}
            continue
        best = None
        for kind in KINDS:
            entry = (row.get('by_kind') or {}).get(kind)
            if kind not in allowed or entry is None:
                continue
            if mate_id is not None and mate_id in (entry.get('not_with') or []):
                entry = entry.get('next')
                if entry is None or mate_id in (entry.get('not_with') or []):
                    continue
            if best is None or entry['rank'] < best[0]['rank']:
                best = (entry, kind)
        out[slot] = None if best is None else {
            'item_id': best[0]['item_id'], 'source': best[0].get('source', ''),
            'kind': best[0].get('source_kind') or best[1], 'delta': best[0].get('dps_delta') or 0,
            'standing': best[0].get('standing'), 'repicked': True}
    return out


def counts(rows):
    tally = dict.fromkeys(KINDS, 0)
    for row in rows:
        if 'item_id' not in row:
            continue
        kinds = set(pick_kinds(row)) | set(row.get('by_kind') or {})
        kinds |= {FILTER_KIND[a['source_kind']] for a in row.get('alternatives', []) if a.get('source_kind') in FILTER_KIND}
        for kind in kinds:
            tally[kind] += 1
    return tally


def words(on, level):
    if len(on) == len(KINDS):
        return 'All sources'
    if level < 60 and sorted(on) == sorted(LEVELING_DEFAULT):
        return 'No rep above Friendly'
    if len(on) == 1:
        return f'{LABEL[on[0]]} only'
    if not on:
        return 'No sources'
    return f'{len(on)} of {len(KINDS)} sources'


def source_text(pick):
    """ItemTooltip.sourceParts: the instance and the boss, else the place and the kind's word; a standing
    ends a reputation or vendor source."""
    place, kind = pick['source'], pick['kind']
    if pick.get('standing'):
        return f'{place} · {pick["standing"].capitalize()}'
    if kind in ('dungeon', 'raid') and ': ' in place:
        return ' · '.join(place.split(': ', 1))
    if kind == 'crafted':
        return f'Crafted · {place}'
    word = KIND_WORD.get(kind)
    return place if word is None or word.lower() == place.lower() else f'{place} · {word}'


def detail_text(slot, pick):
    chance = pick.get('chance')
    pct = f' · {max(1, round(chance * 100))}%' if chance else ''
    return f'{SLOT_WORD[slot]} · {source_text(pick)}{pct}'


ITEM_SLOT = {'finger1': 'finger', 'finger2': 'finger', 'trinket1': 'trinket', 'trinket2': 'trinket'}


def dpp(band_row):
    return {w['stat']: w['dps_per_point'] for w in band_row['weights']}


def score(item_id, weights):
    return sum(v * weights.get(k, 0) for k, v in ITEMS[item_id]['stats'].items())


def worn_for(slot, band_row, level, below):
    """The named worn item a gain is counted against: the best-priced item for the slot at quality up to
    uncommon and a required level ten under `level` (cloth, leather or no armour below 40) that `below`
    (the item shown as the upgrade) still beats by at least 0.3 DPS."""
    weights = dpp(band_row)
    armour = (None, 'cloth', 'leather') if level < 40 else (None, 'cloth', 'leather', 'mail')
    pool = [i for i in ITEMS.values() if i.get('slot') == ITEM_SLOT.get(slot, slot) and i['quality'] <= 2
            and 1 <= (i.get('required_level') or 0) <= level - 10 and i.get('armor_type') in armour and i['stats']
            and score(i['id'], weights) <= score(below, weights) - 0.3]
    return max(pool, key=lambda i: (score(i['id'], weights), -i['id']))['id'] if pool else None


def gain(item_id, worn_id, band_row):
    weights = dpp(band_row)
    return round(score(item_id, weights) - score(worn_id, weights), 1)


def iname(item_id):
    it = ITEMS[item_id]
    return f'<span style="color:{QUALITY[it["quality"]]}">{it["name"]}</span>'


CSS = f'''<style>
.tt, .tt .t {{ font-family: Site, sans-serif; }}
.srow .n.zero {{ color: {MUTED}; }}
.foot {{ height: 18px; display: flex; align-items: center; gap: 0; font-size: 12px; color: {MUTED};
  font-family: Site, sans-serif; margin-top: 2px; white-space: nowrap; }}
.foot .chg {{ color: {GOLD}; cursor: pointer; }}
.foot .chg.hover {{ text-decoration: underline; }}
.foot .state.hover {{ color: {BODY}; }}
.fly {{ position: absolute; width: 232px; background: {ADDON['raised']}; border: 1px solid {ADDON['line']};
  box-shadow: 0 8px 28px rgba(0,0,0,.7); font-family: Site, sans-serif; color: {BODY}; padding: 12px 14px 12px; }}
.fly .hd {{ display: flex; align-items: center; justify-content: space-between; height: 22px; margin-bottom: 6px; }}
.fly .hd .t {{ font-family: Title, serif; color: {ADDON['cream']}; font-size: 15px; }}
.fly .xx {{ width: 20px; height: 20px; border: 1px solid {ADDON['line']}; display: flex; align-items: center;
  justify-content: center; font-size: 14px; color: {BODY}; }}
.srow {{ height: 22px; display: flex; align-items: center; gap: 8px; font-size: 13px; padding: 0 2px; }}
.srow.hover {{ background: {ADDON['hover']}; }}
.srow.focus {{ outline: 1px solid {GOLD}; outline-offset: -1px; }}
.srow .n {{ flex: 1; }}
.srow .c {{ color: {MUTED}; font-variant-numeric: tabular-nums; font-size: 12px; }}
.box {{ width: 14px; height: 14px; flex: none; display: inline-flex; align-items: center; justify-content: center; }}
.box.on {{ background: {GOLD}; }}
.box.off {{ border: 1px solid {ADDON['line']}; }}
.fly .ft {{ font-size: 12px; color: {MUTED}; margin-top: 8px; line-height: 15px; }}
.fly .df {{ font-size: 12px; color: {MUTED}; margin-top: 8px; }}
.screen {{ position: relative; }}
.edge {{ position: absolute; top: 0; bottom: 0; width: 2px; background: rgba(229,185,85,.35); }}
.edge-cap {{ position: absolute; top: -18px; font-family: -apple-system, Arial, sans-serif; font-size: 11px; color: #9a9484; }}
.set {{ background: {ADDON['raised']}; border: 1px solid {ADDON['line']}; font-family: Site, sans-serif; color: {BODY};
  padding: 16px 18px 18px; width: 560px; }}
.set .title {{ font-family: Title, serif; color: {GOLD}; font-size: 20px; margin-bottom: 14px; }}
.set .gh {{ display: flex; justify-content: space-between; align-items: baseline; color: {GOLD}; font-size: 14px; }}
.set .rl {{ height: 1px; background: {ADDON['line']}; margin: 6px 0 8px; }}
.set .hint {{ font-size: 13px; color: {MUTED}; margin-bottom: 6px; }}
.set .sw {{ display: flex; align-items: flex-start; gap: 10px; padding: 6px 0; font-size: 13.5px; }}
.set .sw .h {{ font-size: 13px; color: {MUTED}; margin-top: 3px; }}
.set .sw .c {{ margin-left: auto; color: {MUTED}; font-size: 12px; font-variant-numeric: tabular-nums; }}
.set .dim {{ opacity: .5; }}
</style>'''

TICK = ('<svg width="12" height="12" viewBox="0 0 12 12"><path d="M2 6.2 L4.8 9 L10 2.8" fill="none" '
        f'stroke="{ADDON["bg"]}" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/></svg>')


def box(on):
    return f'<span class="box {"on" if on else "off"}">{TICK if on else ""}</span>'


def upgrade_row(item_id, gain, detail):
    it = ITEMS[item_id]
    return (f'<div class="irow">{img(it["icon"], 30)}<div class="tx"><div class="n"><span>{iname(item_id)}</span>'
            f'<span style="color:{C["green"]}">+{gain} DPS</span></div><div class="s">{detail}</div></div></div>')


def foot(state, hover_change=False, hover_state=False):
    return (f'<div class="foot"><span class="state{" hover" if hover_state else ""}">{state}</span>'
            f'<span>&nbsp;·&nbsp;</span><span class="chg{" hover" if hover_change else ""}">Change</span></div>')


def talents():
    return (f'<div class="sect">Talents<span class="r">12 of 15 points</span></div><div class="irow">'
            f'{img("ability_searingarrow", 30)}<div class="tx" style="overflow:hidden"><div class="n" style="overflow:hidden;'
            f'text-overflow:ellipsis;display:block">Level 25: Lethal Attacks</div><div class="s">Beast Mastery leveling</div></div>'
            f'<span class="sbtn" style="{ux.SMALL_BTN}">Change</span>'
            f'<span class="sbtn" style="{ux.SMALL_BTN};margin-left:6px">Talents</span></div>')


def upgrades(rows, state, rest=None, **hover):
    body = '<div class="sect">Upgrades<span class="r">2 h ago</span></div>' + ''.join(rows)
    if rest:
        body += f'<div class="small" style="margin:2px 0 4px">{rest}</div>'
    return body + foot(state, **hover)


WEAPONS = ('main_hand', 'off_hand', 'ranged')


def priced_rows(band_row, level, picks, slots):
    """(gain, slot, pick) for `slots`, each priced against worn_for's item, best gain first, gains above 0."""
    priced = []
    for slot in slots:
        pick = picks.get(slot)
        if not pick or slot in WEAPONS:
            continue
        worn = worn_for(slot, band_row, level, pick['item_id'])
        figure = gain(pick['item_id'], worn, band_row) if worn else 0
        if figure > 0:
            priced.append((figure, slot, pick))
    priced.sort(key=lambda row: -row[0])
    return priced


def as_rows(priced):
    return [upgrade_row(p['item_id'], f'{g:.1f}', detail_text(s, p)) for g, s, p in priced]


def rows_for(band_row, level, picks, slots, limit=3):
    return as_rows(priced_rows(band_row, level, picks, slots)[:limit])


def leveling_rows(on, rep_cap=True, limit=3):
    """Up to `limit` Upgrades rows under a choice: the re-picked slots first, then the kept world drops
    (band 20 re-picks two armour slots under the leveling default)."""
    picks = resolve(B20['slots'], on, rep_cap)
    repicked = [s for s, p in picks.items() if p and p['repicked'] and s not in WEAPONS]
    priced = priced_rows(B20, 24, picks, repicked)[:limit]
    kept = [s for s, p in picks.items() if p and not p['repicked'] and p['kind'] == 'world_drop']
    priced += priced_rows(B20, 24, picks, kept)[:limit - len(priced)]
    # The panel lists the rows by gain (Gear.upgrades' order).
    return as_rows(sorted(priced, key=lambda row: -row[0]))


def bow_panel(sections, width=340):
    return ux.panel('Bow Jackzon', 'Level 24 Hunter · Employed Millennials', sections, width=width,
                    score=ux.score_fact(1240, 'Rare'))


def board_panel_leveling():
    p = bow_panel([upgrades(leveling_rows(LEVELING_DEFAULT), 'No rep above Friendly'), talents()])
    return page(CSS + '<p class="cap"><b>Leveling default</b> · Bow Jackzon, level 24 Beast Mastery, band 20: '
                'Reputation above Friendly is off, so the Honored neck and ring are re-picked</p>'
                f'<div style="padding:8px 0 0 8px">{p}</div>')


def board_panel_all():
    picks = resolve(B60['slots'], KINDS)
    rows = rows_for(B60, 60, picks, list(picks))
    p = ux.panel('Bow Jackzon', 'Level 60 Hunter · Employed Millennials', [upgrades(rows, 'All sources')],
                 score=ux.score_fact(4210, 'Epic'))
    return page(CSS + '<p class="cap"><b>Level 60</b> · every kind on by default: the line still shows, '
                'so the switch is findable; a finished build on plan shows no Talents region</p>'
                f'<div style="padding:8px 0 0 8px">{p}</div>')


def board_panel_empty():
    p = bow_panel([upgrades([], 'No sources', rest='No upgrades from these sources'), talents()])
    return page(CSS + '<p class="cap"><b>No sources</b> · every kind turned off: the Upgrades region says so '
                'where the rows were; "2 h ago" is the sim\'s age, which still holds</p>'
                f'<div style="padding:8px 0 0 8px">{p}</div>')


def flyout(on, saved, focus=None, hover=None, rep_tip=False):
    tally = counts(B20['slots'])
    rows = ''.join(
        f'<div class="srow{" focus" if k == focus else ""}{" hover" if k == hover else ""}">{box(k in on)}'
        f'<span class="n{" zero" if tally[k] == 0 else ""}">{LABEL[k]}</span><span class="c">{tally[k]}</span></div>'
        for k in KINDS)
    default = '<div class="df">Default</div>' if saved else ''
    return f'<div class="hd"><span class="t">Sources</span><span class="xx">×</span></div>{rows}{default}'


def with_flyout(panel_html, fly_html, left, screen_w=None, extra='', extra_w=0, height=600):
    """The panel at (x, 0) with the flyout 4 px beside it: on its right, or on its left at the screen edge."""
    if left:
        x = screen_w - 342
        fly = f'<div class="fly" style="left:{x - 4 - 232}px;top:0">{fly_html}</div>'
        edge = (f'<div class="edge" style="left:{screen_w}px"></div>'
                f'<div class="edge-cap" style="left:{screen_w - 120}px">screen right edge</div>')
        return (f'<div class="screen" style="width:{screen_w + 4}px;height:{height}px;margin-top:18px">'
                f'<div style="position:absolute;left:{x}px;top:0">{panel_html}</div>{fly}{edge}</div>')
    return (f'<div class="screen" style="width:{342 + 4 + 232 + extra_w}px;height:{height}px">'
            f'<div style="position:absolute;left:0;top:0">{panel_html}</div>'
            f'<div class="fly" style="left:{342 + 4}px;top:0">{fly_html}</div>{extra}</div>')


def board_flyout_default():
    p = bow_panel([upgrades(leveling_rows(LEVELING_DEFAULT), 'No rep above Friendly'), talents()])
    tip = tooltip([l('Standings up to Friendly still count.', color=C['grey'])])
    # The row's tooltip beside the hovered Reputation row (6th row: header 28 + 5 rows of 22, padding 12).
    extra = f'<div style="position:absolute;left:{342 + 4 + 232 + 6}px;top:{12 + 28 + 5 * 22 - 4}px">{tip}</div>'
    body = with_flyout(p, flyout(LEVELING_DEFAULT, saved=False, hover='rep'), left=False, extra=extra,
                       extra_w=270, height=392)
    return page(CSS + '<p class="cap"><b>Sources, leveling default</b> · nothing saved, so no Default; the '
                'Reputation row hovered</p>' + body)


def board_flyout_chosen():
    on = ['dungeon']
    p = bow_panel([upgrades(leveling_rows(on, rep_cap=False), 'Dungeon only'), talents()])
    body = with_flyout(p, flyout(on, saved=True, focus='quest'), left=True, screen_w=640, height=392)
    return page(CSS + '<p class="cap"><b>Sources, Dungeon only</b> · a saved choice shows Default; the panel at '
                'the screen\'s right edge puts the flyout on its left; Quest has keyboard focus</p>' + body)


def board_state_tooltip():
    p = bow_panel([upgrades(leveling_rows(LEVELING_DEFAULT), 'No rep above Friendly', hover_state=True), talents()])
    tip = tooltip([f'<div class="t" style="color:#fff">Sources</div>',
                   l('Off: Reputation above Friendly', color=C['grey'])], width=240)
    return page(CSS + '<p class="cap"><b>Hover on the state words</b> · what is off, in the flyout\'s words</p>'
                f'<div style="display:flex;gap:6px;align-items:flex-end;padding:8px 0 0 8px">{p}'
                f'<div style="margin-bottom:96px">{tip}</div></div>')


def board_settings():
    tally = counts(B20['slots'])
    on = ['dungeon', 'world', 'quest', 'crafted', 'vendor', 'pvp']
    rows = ''.join(f'<div class="sw">{box(k in on)}<span style="{"color:" + MUTED if tally[k] == 0 else ""}">'
                   f'{LABEL[k]}</span><span class="c">{tally[k]}</span></div>'
                   for k in KINDS)
    prior = ''.join(f'<div class="sw dim">{box(True)}<div><div>{label}</div><div class="h">{hint}</div></div></div>'
                    for label, hint in [('Item tooltips', 'Our lines under the client\'s on every item.'),
                                        ('Score badges in chat', 'A guildmate\'s score beside their name.')])
    page_html = (f'<div class="set"><div class="title">Forever Sixty</div>{prior}'
                 f'<div style="height:14px"></div><div class="gh"><span>Gear sources</span>'
                 f'<span style="color:{MUTED};font-size:12px">Default</span></div><div class="rl"></div>'
                 f'{rows}</div>')
    return page(CSS + '<p class="cap"><b>Settings</b> · the Gear sources group after the tooltip switches '
                '(dimmed here); a saved choice, 6 of 8 sources, shows Default</p>' + page_html)


STAT_WORD = {'agility': 'Agility', 'stamina': 'Stamina', 'strength': 'Strength', 'intellect': 'Intellect',
             'spirit': 'Spirit'}


def worn_tooltip(item_id, slot_word):
    """The client's own lines for a worn item: name, Soulbound, slot, armour, base stats, level."""
    it = ITEMS[item_id]
    lines = [f'<div class="t" style="color:{QUALITY[it["quality"]]}">{it["name"]}</div>', l('Soulbound'),
             l(slot_word, (it.get('armor_type') or '').capitalize() or None)]
    if it.get('armor'):
        lines.append(l(f'{it["armor"]} Armor'))
    lines += [l(f'+{v} {STAT_WORD[k]}') for k, v in it['stats'].items() if k in STAT_WORD]
    if it.get('required_level'):
        lines.append(l(f'Requires Level {it["required_level"]}'))
    return lines


def next_line(item_id, figure):
    return l(f'Next upgrade: {iname(item_id)} <span style="color:{C["green"]}">+{figure:.1f} DPS</span>')


def board_set_aside():
    picks = resolve(B20['slots'], LEVELING_DEFAULT, rep_cap=True)
    neck = picks['neck']
    best = next(row for row in B20['slots'] if row['slot'] == 'neck')['item_id']
    worn = worn_for('neck', B20, 24, neck['item_id'])
    gap = gain(best, neck['item_id'], B20)
    lines = worn_tooltip(worn, 'Neck') + [
        next_line(neck['item_id'], gain(neck['item_id'], worn, B20)),
        l(detail_text('neck', neck).split(' · ', 1)[1], color=C['grey']),
        l(f'All sources: {iname(best)} +{gap:.1f} DPS', color=C['grey']),
        l('Forever Sixty · Beast Mastery estimate', color='#8a8a8a')]
    return page(CSS + '<p class="cap" style="max-width:640px;line-height:18px"><b>Shift held</b> · the worn neck under '
                'the leveling default: Next upgrade is the re-picked quest neck; the Shift line names the Honored Sentinel\'s Medallion '
                'set aside and the gap, priced from the two items\' stats with the band\'s weights</p>'
                f'<div style="padding-top:6px">{tooltip(lines, width=380)}</div>')


def board_rep_standing():
    picks = resolve(B20['slots'], KINDS)
    ring = picks['finger2']
    worn = worn_for('finger2', B20, 24, ring['item_id'])
    lines = worn_tooltip(worn, 'Finger') + [
        next_line(ring['item_id'], gain(ring['item_id'], worn, B20)),
        l(source_text(ring), color=C['grey'])]
    return page(CSS + '<p class="cap"><b>Reputation source</b> · Reputation on (all sources): hovering a worn ring, '
                'the muted line under Next upgrade is that upgrade\'s source, its faction and the standing it asks</p>'
                f'<div style="padding-top:6px">{tooltip(lines, width=380)}</div>')


PANE_ORDER = ['head', 'neck', 'shoulder', 'back', 'chest', 'wrist', 'hands', 'waist', 'legs', 'feet',
              'finger1', 'finger2', 'trinket1', 'trinket2', 'main_hand', 'off_hand', 'ranged']


def kind_best(row, kind):
    """The best item `kind` gives a slot: the pick where it is that kind's, else the kind's by_kind entry,
    else an alternative of that kind."""
    if kind in pick_kinds(row):
        return row['item_id']
    entry = (row.get('by_kind') or {}).get(kind)
    if entry:
        return entry['item_id']
    return next((a['item_id'] for a in row.get('alternatives', []) if FILTER_KIND.get(a['source_kind']) == kind), None)


def board_row_hover():
    rows = {row['slot']: row for row in B20['slots'] if 'item_id' in row}
    lines = [f'<div class="t" style="color:#fff">World / BoE</div>']
    for slot in PANE_ORDER:
        best = kind_best(rows[slot], 'world') if slot in rows else None
        line = l(f'{SLOT_WORD[slot]}: {iname(best)}', color=C['grey']) if best else None
        if line and line not in lines:  # both rings naming the same unique ring read once
            lines.append(line)
    p = bow_panel([upgrades(leveling_rows(LEVELING_DEFAULT), 'No rep above Friendly'), talents()])
    extra = f'<div style="position:absolute;left:{342 + 4 + 232 + 6}px;top:{12 + 28 + 2 * 22 - 4}px">{tooltip(lines)}</div>'
    body = with_flyout(p, flyout(LEVELING_DEFAULT, saved=False, hover='world'), left=False, extra=extra,
                       extra_w=300, height=392)
    return page(CSS + '<p class="cap"><b>Row hover</b> · World / BoE: the best item that kind gives each slot at '
                'band 20, in character pane order</p>' + body)


def board_no_world():
    on = [k for k in LEVELING_DEFAULT if k != 'world']
    picks = resolve(B20['slots'], on, rep_cap=True)
    away = [s for s, p in picks.items() if p and p['repicked']]
    p = bow_panel([upgrades(rows_for(B20, 24, picks, away), words(on, 24)), talents()])
    return page(CSS + '<p class="cap"><b>No world drops</b> · level 24, World / BoE turned off on top of the leveling '
                'default: the world drop slots re-pick from crafts and vendors, the Honored neck re-picked to a quest neck</p>'
                f'<div style="padding:8px 0 0 8px">{p}</div>')


BOARDS = {
    'panel-sources-leveling': board_panel_leveling,
    'panel-sources-all': board_panel_all,
    'panel-sources-empty': board_panel_empty,
    'flyout-sources-default': board_flyout_default,
    'flyout-sources-chosen': board_flyout_chosen,
    'panel-sources-state-tooltip': board_state_tooltip,
    'settings-gear-sources': board_settings,
    'tooltip-set-aside-shift': board_set_aside,
    'tooltip-rep-standing': board_rep_standing,
    'flyout-sources-row-hover': board_row_hover,
    'panel-sources-no-world': board_no_world,
}


if __name__ == '__main__':
    ux.BOARDS = BOARDS
    ux.main(sys.argv[1:] or list(BOARDS))
