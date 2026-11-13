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
- The set-aside gap is the entry's own dps_delta from the bis file. The Upgrades rows' gains are
  illustrative layout figures (the build reads Progress.pickGain against what is worn).
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
                         'delta': 0, 'standing': row.get('standing'), 'repicked': False}
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


def iname(item_id):
    it = ITEMS[item_id]
    return f'<span style="color:{QUALITY[it["quality"]]}">{it["name"]}</span>'


CSS = f'''<style>
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


def leveling_rows(on, rep_cap=True, limit=3):
    """Up to `limit` Upgrades rows under a choice (gains illustrative): the re-picked slots first, then the
    ranged slot's pick when it still stands (band 20 re-picks two slots under the leveling default)."""
    picks = resolve(B20['slots'], on, rep_cap)
    slots = [s for s, p in picks.items() if p and p['repicked']]
    if len(slots) < limit and picks.get('ranged') and 'ranged' not in slots:
        slots.append('ranged')
    rows = []
    for slot, gain in zip(slots[:limit], ['2.4', '1.6', '0.9']):
        pick = picks[slot]
        rows.append(upgrade_row(pick['item_id'], gain, f'{SLOT_WORD[slot]} · {source_text(pick)}'))
    return rows


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
    order = [s for s in ('head', 'chest', 'legs') if picks.get(s)]
    rows = [upgrade_row(picks[s]['item_id'], g, f'{SLOT_WORD[s]} · {source_text(picks[s])}')
            for s, g in zip(order, ['6.8', '4.1', '2.7'])]
    p = ux.panel('Bow Jackzon', 'Level 60 Hunter · Employed Millennials', [upgrades(rows, 'All sources')],
                 score=ux.score_fact(4210, 'Epic'))
    return page(CSS + '<p class="cap"><b>Level 60</b> · every kind on by default: the line still shows, '
                f'so the switch is findable</p><div style="padding:8px 0 0 8px">{p}</div>')


def board_panel_empty():
    p = bow_panel([upgrades([], 'No sources', rest='No upgrades from these sources'), talents()])
    return page(CSS + '<p class="cap"><b>No sources</b> · every kind turned off: the Upgrades region says so '
                f'where the rows were</p><div style="padding:8px 0 0 8px">{p}</div>')


def flyout(on, saved, focus=None, hover=None, rep_tip=False):
    tally = counts(B20['slots'])
    rows = ''.join(
        f'<div class="srow{" focus" if k == focus else ""}{" hover" if k == hover else ""}">{box(k in on)}'
        f'<span class="n">{LABEL[k]}</span><span class="c">{tally[k]}</span></div>' for k in KINDS)
    default = '<div class="df">Default</div>' if saved else ''
    return (f'<div class="hd"><span class="t">Sources</span><span class="xx">×</span></div>{rows}{default}'
            '<div class="ft">Counts are gear slots with an option from that source.</div>')


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
    rows = ''.join(f'<div class="sw">{box(k in on)}<span>{LABEL[k]}</span><span class="c">{tally[k]}</span></div>'
                   for k in KINDS)
    prior = ''.join(f'<div class="sw dim">{box(True)}<div><div>{label}</div><div class="h">{hint}</div></div></div>'
                    for label, hint in [('Item tooltips', 'Our lines under the client\'s on every item.'),
                                        ('Score badges in chat', 'A guildmate\'s score beside their name.')])
    page_html = (f'<div class="set"><div class="title">Forever Sixty</div>{prior}'
                 f'<div style="height:14px"></div><div class="gh"><span>Gear sources</span>'
                 f'<span style="color:{MUTED};font-size:12px">Default</span></div><div class="rl"></div>'
                 f'<div class="hint">Upgrades are picked from the sources you leave on.</div>{rows}</div>')
    return page(CSS + '<p class="cap"><b>Settings</b> · the Gear sources group after the tooltip switches '
                '(dimmed here); a saved choice, 6 of 8 sources, shows Default</p>' + page_html)


def bag_tooltip(item_id, slot_line, stats):
    it = ITEMS[item_id]
    return ([f'<div class="t" style="color:{QUALITY[it["quality"]]}">{it["name"]}</div>', l('Soulbound'),
             l(slot_line)] + [l(s) for s in stats] + [l(f'Requires Level {it.get("required_level", 1)}')])


def board_set_aside():
    picks = resolve(B20['slots'], LEVELING_DEFAULT, rep_cap=True)
    neck = picks['neck']
    best = next(s for s in B20['slots'] if s['slot'] == 'neck')
    gap = abs(neck['delta'])
    lines = bag_tooltip(281265, 'Neck', ['+3 Spirit']) + [
             l(f'Next upgrade: {iname(neck["item_id"])} <span style="color:{C["green"]}">+2.4 DPS</span>'),
             l(source_text(neck), color=C['grey']),
             l(f'Best with all sources: {iname(best["item_id"])}, {gap:.1f} DPS more', color=C['grey']),
             l('Forever Sixty · Beast Mastery estimate', color='#8a8a8a')]
    return page(CSS + '<p class="cap"><b>Shift held</b> · the worn neck under the leveling default: the set-aside '
                'line names the Honored pick and its gap (dps_delta from the bis file)</p>'
                f'<div style="padding-top:6px">{tooltip(lines, width=380)}</div>')


def board_rep_standing():
    picks = resolve(B20['slots'], KINDS)
    ring = picks['finger2']
    lines = [f'<div class="t" style="color:{QUALITY[2]}">Volcanic Rock Ring</div>', l('Binds when equipped'),
             l('Finger'), l('+1 Stamina'), l('+3 Intellect'), l('Requires Level 16'),
             l(f'Next upgrade: {iname(ring["item_id"])} <span style="color:{C["green"]}">+1.6 DPS</span>'),
             l(source_text(ring), color=C['grey'])]
    return page(CSS + '<p class="cap"><b>Reputation source</b> · all sources on: the pick\'s source line ends with '
                'the standing it asks</p>' f'<div style="padding-top:6px">{tooltip(lines, width=380)}</div>')


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
}


if __name__ == '__main__':
    ux.BOARDS = BOARDS
    ux.main(sys.argv[1:] or list(BOARDS))
