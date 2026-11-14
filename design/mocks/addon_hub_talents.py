"""The talent marks on the client's whole talent window, for the store gallery (design/listing/gen_captures.py).

    python3 design/mocks/addon_hub_talents.py

Writes design/reviews/2026-11-14-store-listing/round-2/talent-marks-board.png at device scale 2. The L10 board
(addon_hub_l10.py l10-talents) draws only the strip and two nodes; this one draws the 1.60 client's
three-tree window as the beta capture shows it (web/public/addon/talent-marks-trees.jpg: Primary and
Secondary tabs, Search, Unspent Talents, the trees side by side on their art, Apply Changes at the foot)
with the marks drawn to the built code's own rules:

- Tree, talents, icons, prerequisites and art: data/builds/<Data.lua's build>/talents/hunter.json, trees/, icons/.
- The plan: that build's addon-builds.json, the Beast Mastery leveling order (Data.lua's default for the spec).
- Badges: surfaces/TalentMarks.lua. Progress.cells over the plan, a plate only on a node not yet taken as
  planned, its text the plan's next run there (badgeText, L.talentBadgeRange), its look by badgeLook: the
  next node solid gold with dark digits; a planned node whose run starts beyond Progress.reachablePoints
  on the inset plate, edge at ALPHA.unreachable, digits at ALPHA.dimDigits. Sizes from Theme.SIZES.
- The ring: TalentGlow.ring() through Theme.showGlow, a 3 px gold outline 3 to 6 px outside the button with
  a 1 px dark edge at ALPHA.ringEdge.
- The strip: TalentMarks.stripSegments with Locale's talentStrip* strings, on the background at ALPHA.strip,
  bottom left, centred on Apply Changes.

The character is made up: a level 33 hunter following the default Beast Mastery leveling build, 22 points
placed as planned, one in Deflection the plan never wanted (the strip's off-plan count, Progress.build's
offPlan) and one to spend; the next talent is untouched (0/1), so its badge is the single point 23
(a partly taken one names its whole run, "21–22" on 1/2, which is right but reads oddly on the hero plate).
"""
import json
import os
import re

import addon_hub as hub
import addon_ux as ux
from addon_builds import L, esc, lua_format
from addon_ux import ADDON, C, ROOT, icon_uri

DATA_BUILD = re.search(r'build = "([\d.]+)"', (ROOT / 'addon/ForeverSixty/Data.lua').read_text()).group(1)
BUILD = ROOT / 'data/builds' / DATA_BUILD
OUT = ROOT / 'design/reviews/2026-11-14-store-listing/round-2'
CLASS, SPEC, KIND = 'hunter', 'hunter-beast-mastery', 'leveling'
LEVEL, ON_PLAN = 33, 22  # level 33 allows 24 points
OFF_PLAN = {(3, 1, 3): 1}  # Deflection, which the plan never takes

# Theme.SIZES and Theme.ALPHA, as the addon draws the marks.
BADGE, BADGE_TEXT, BADGE_OFFSET, BADGE_PAD = 16, 12, 4, 6
STRIP, STRIP_TEXT, STRIP_INSET, STRIP_PAD, WINDOW_BORDER = 26, 13, 8, 12, 8
GLOW_INSET, GLOW_THICKNESS, GLOW_EDGE = 3, 3, 1
A_PLATE, A_UNREACHABLE, A_DIM_DIGITS, A_STRIP, A_RING_EDGE = 0.9, 0.5, 0.6, 0.9, 0.6

# The window's layout, in CSS px.
NODE, COL_STEP, TIER_STEP, GRID_LEFT, GRID_TOP = 38, 60, 56, 41, 62
PANEL_W, PANEL_H, PANEL_GAP = 296, 470, 6
POINTS_PER_TIER = 5


def rgba(hex_color, alpha):
    h = hex_color.lstrip('#')
    return f'rgba({int(h[0:2], 16)},{int(h[2:4], 16)},{int(h[4:6], 16)},{alpha})'


def plan():
    trees = json.loads((BUILD / f'talents/{CLASS}.json').read_text())['trees']
    order = [tuple(p) for p in json.loads((BUILD / 'addon-builds.json').read_text())[SPEC][KIND]['order']]
    return trees, order


def ranks_of(order):
    ranks = dict(OFF_PLAN)
    for key in order[:ON_PLAN]:
        ranks[key] = ranks.get(key, 0) + 1
    return ranks


def cells(order, ranks):
    """Progress.cells: every planned cell once, with its runs of order indexes; `next` is the plan's next point."""
    have_left = dict(ranks)
    point = None
    for key in order:
        if have_left.get(key, 0) > 0:
            have_left[key] -= 1
        elif point is None:
            point = key
    out = {}
    for index, key in enumerate(order, 1):
        cell = out.setdefault(key, {'want': 0, 'runs': [], 'have': ranks.get(key, 0), 'next': key == point})
        cell['want'] += 1
        if cell['runs'] and cell['runs'][-1][1] == index - 1:
            cell['runs'][-1][1] = index
        else:
            cell['runs'].append([index, index])
    return out, point


def next_run(cell):
    """Progress.nextRun, else the last run (TalentMarks badgeRun)."""
    counted = 0
    for run in cell['runs']:
        counted += run[1] - run[0] + 1
        if cell['have'] < counted:
            return run
    return cell['runs'][-1]


def badge_html(cell, reach):
    """TalentMarks.badgeText and badgeLook; None for a node taken as planned (no plate)."""
    if not cell['next'] and cell['have'] >= cell['want']:
        return None
    run = next_run(cell)
    text = str(run[1]) if run[0] == run[1] else lua_format(L('talentBadgeRange'), run[0], run[1])
    if cell['next']:
        plate, edge, digits, shadow = ADDON['gold'], ADDON['gold'], ADDON['bg'], 'none'
    elif run[0] > reach:
        plate, edge = rgba(ADDON['inset'], A_PLATE), rgba(ADDON['gold'], A_UNREACHABLE)
        digits, shadow = rgba(ADDON['gold'], A_DIM_DIGITS), '1px -1px 0 #000'
    else:
        plate, edge, digits, shadow = rgba(ADDON['bg'], A_PLATE), ADDON['gold'], ADDON['gold'], '1px -1px 0 #000'
    return (f'<div class="bdg" style="background:{plate};box-shadow:inset 0 0 0 1px {edge};color:{digits};'
            f'text-shadow:{shadow}">{esc(text)}</div>')


def ring_html():
    out = GLOW_INSET + GLOW_THICKNESS
    return (f'<div style="position:absolute;inset:-{out}px;box-shadow:inset 0 0 0 {GLOW_THICKNESS}px {ADDON["gold"]},'
            f'0 0 0 {GLOW_EDGE}px {rgba("#000000", A_RING_EDGE)}"></div>')


def node_xy(talent):
    return GRID_LEFT + talent['column'] * COL_STEP, GRID_TOP + talent['tier'] * TIER_STEP


def connector(src, dst, lit):
    """The client's prerequisite line: down from the prerequisite, across when the columns differ."""
    (sx, sy), (dx, dy) = node_xy(src), node_xy(dst)
    colour = '#d9b54a' if lit else '#6a6a6a'
    mid = NODE // 2 - 2
    if sx == dx:
        return f'<div class="ln" style="left:{sx + mid}px;top:{sy + NODE}px;width:4px;height:{dy - sy - NODE}px;background:{colour}"></div>'
    left, width = min(sx, dx) + mid, abs(dx - sx)
    return (f'<div class="ln" style="left:{left}px;top:{sy + mid}px;width:{width + 4}px;height:4px;background:{colour}"></div>'
            f'<div class="ln" style="left:{dx + mid}px;top:{sy + mid}px;width:4px;height:{dy - sy - mid}px;background:{colour}"></div>')


def node_html(talent, rank, open_tier, planned, reach):
    x, y = node_xy(talent)
    available = talent['tier'] <= open_tier
    state = 'max' if rank >= talent['max_rank'] else 'part' if rank > 0 else 'open' if available else 'locked'
    rank_colour = {'max': C['gold'], 'part': C['green'], 'open': C['green'], 'locked': '#8a8a8a'}[state]
    marks = ''
    if planned is not None:
        badge = badge_html(planned, reach)
        marks = (ring_html() if planned['next'] else '') + (badge or '')
    return (f'<div class="node {state}" style="left:{x}px;top:{y}px">'
            f'<img src="{icon_uri(talent["icon"])}" alt="">'
            f'<span class="rk" style="color:{rank_colour}">{rank}/{talent["max_rank"]}</span>{marks}</div>')


def tree_html(tab, tree, ranks, planned, reach):
    spent = sum(n for (t, _, _), n in ranks.items() if t == tab)
    open_tier = spent // POINTS_PER_TIER
    by_id = {t['id']: t for t in tree['talents']}
    rank = {t['id']: ranks.get((tab, t['tier'] + 1, t['column'] + 1), 0) for t in tree['talents']}
    lines = ''.join(connector(by_id[t['prereq_talent_id']], t, rank[t['prereq_talent_id']] >= by_id[t['prereq_talent_id']]['max_rank'])
                    for t in tree['talents'] if t['prereq_talent_id'])
    nodes = ''.join(node_html(t, rank[t['id']], open_tier, planned.get((tab, t['tier'] + 1, t['column'] + 1)), reach)
                    for t in tree['talents'])
    art = ux.b64((BUILD / 'trees' / f'{tree["background"]}.webp').read_bytes(), 'image/webp')
    return (f'<div class="tree"><div class="art" style="background-image:url({art})"></div>'
            f'<div class="th"><img src="{icon_uri(tree["icon"])}" alt=""><span>{esc(tree["name"])}</span></div>'
            f'{lines}{nodes}</div>')


def strip_text(order, ranks, cells_by_key, point, trees):
    """TalentMarks.stripSegments with a point waiting: "Spend 1 now: <talent> · <matched> of <allowed> · N off plan",
    the off-plan segment in the warning colour; returned as HTML."""
    spent = sum(ranks.values())
    unspent = LEVEL - 9 - spent
    tab, tier, column = point
    name = next(t['name'] for t in trees[tab - 1]['talents'] if (t['tier'] + 1, t['column'] + 1) == (tier, column))
    matched = sum(min(ranks.get(k, 0), c['want']) for k, c in cells_by_key.items())
    allowed = min(len(order), LEVEL - 9)
    lead = lua_format(L('talentStripSpendOn'), unspent, name)
    count = lua_format(L('talentStripCount'), min(matched, allowed), allowed)
    segments = [esc(lead), esc(count)]
    off_plan = spent - matched
    if off_plan > 0:
        segments.append(f'<span style="color:{ADDON["warning"]}">{esc(lua_format(L("talentStripOffPlan"), off_plan))}</span>')
    return esc(L('talentStripSeparator')).join(segments), unspent


CSS = f'''
.tw {{ width: {3 * PANEL_W + 2 * PANEL_GAP + 2 * WINDOW_BORDER + 4}px; padding: 0 {WINDOW_BORDER}px {WINDOW_BORDER}px; }}
.tw .portrait {{ background-image: url({ux.CREST}); }}
/* The talent window's own frame: the beta capture's bronze and orange rim and title. */
.frame.tw {{ border-color: #9a6424; box-shadow: inset 0 0 0 1px #000, inset 0 0 0 3px #5e3a12, 0 0 0 1px #2a1606,
  0 8px 28px rgba(0,0,0,.7); background: linear-gradient(#17130e, #0e0c09); }}
.frame.tw .title {{ background: linear-gradient(#3a2610, #1c1208); border-bottom-color: #7a4f1c; }}
.caret {{ display: inline-block; margin-left: 8px; border: 5px solid transparent; border-top-color: currentColor;
  transform: translateY(3px); }}
.lock {{ display: inline-block; margin-left: 8px; width: 9px; height: 7px; background: #7c7c7c; border-radius: 1px;
  position: relative; top: 2px; }}
.lock::before {{ content: ''; position: absolute; left: 1px; top: -6px; width: 5px; height: 6px;
  border: 1px solid #7c7c7c; border-bottom: 0; border-radius: 4px 4px 0 0; }}
.tbar {{ height: 38px; display: flex; align-items: center; gap: 6px; padding: 0 2px 0 52px; }}
.tab {{ height: 24px; display: inline-flex; align-items: center; padding: 0 16px; font-size: 13px; border-radius: 3px; }}
.tab.on {{ color: {C['gold']}; border: 1px solid #8a6d2a; background: linear-gradient(#4a3b16, #241c09); }}
.tab.off {{ color: #7c7c7c; border: 1px solid #3a3a3a; background: linear-gradient(#232323, #141414); }}
.search {{ margin-left: auto; width: 170px; height: 22px; border: 1px solid #4a4a4a; border-radius: 3px;
  background: #07080c; color: #6f6f6f; font-size: 12px; display: flex; align-items: center; padding: 0 8px; }}
.unspent {{ display: flex; align-items: center; gap: 8px; margin-left: 14px; font-size: 13px; color: {C['gold']}; }}
.unspent b {{ min-width: 26px; height: 24px; display: inline-flex; align-items: center; justify-content: center;
  font-weight: 400; font-size: 16px; color: #fff; border: 1px solid #6b5626; background: #07080c; }}
.trees {{ display: flex; gap: {PANEL_GAP}px; }}
.tree {{ position: relative; width: {PANEL_W}px; height: {PANEL_H}px; overflow: hidden;
  box-shadow: inset 0 0 0 1px #000, inset 0 0 0 2px #5a3f1a; border-radius: 2px; }}
/* The client lights its tree art far brighter than the raw texture reads (the beta capture). */
.tree .art {{ position: absolute; inset: 0; background-size: cover; background-position: center;
  filter: brightness(2.1) saturate(1.5); }}
.th {{ position: absolute; left: 0; right: 0; top: 10px; display: flex; align-items: center; justify-content: center;
  gap: 8px; font-size: 15px; color: {C['gold']}; text-shadow: 1px -1px 0 #000; }}
.th img {{ width: 28px; height: 28px; border-radius: 50%; box-shadow: 0 0 0 2px #b8913f, 0 0 0 3px #000; }}
.ln {{ position: absolute; box-shadow: 0 0 0 1px rgba(0,0,0,.7); }}
.node {{ position: absolute; width: {NODE}px; height: {NODE}px; }}
.node img {{ width: 100%; height: 100%; display: block; border-radius: 3px; box-shadow: 0 0 0 1px #000, 0 0 0 2px #5a5a5a; }}
.node.max img {{ box-shadow: 0 0 0 1px #000, 0 0 0 2px {C['gold']}; }}
.node.part img, .node.open img {{ box-shadow: 0 0 0 1px #000, 0 0 0 2px #2fbf3a; }}
.node.locked img {{ filter: grayscale(1) brightness(.55); }}
.rk {{ position: absolute; right: -6px; bottom: -7px; font-size: 11px; line-height: 13px; padding: 0 3px;
  background: #07080c; border: 1px solid #3a3a3a; border-radius: 2px; }}
.bdg {{ position: absolute; top: -{BADGE_OFFSET}px; right: -{BADGE_OFFSET}px; height: {BADGE}px; min-width: {BADGE}px;
  padding: 0 {BADGE_PAD // 2}px; box-sizing: border-box; display: flex; align-items: center; justify-content: center;
  font-family: Site, sans-serif; font-size: {BADGE_TEXT}px; line-height: {BADGE}px; white-space: nowrap; }}
.foot {{ position: relative; height: 40px; display: flex; align-items: center; justify-content: center; }}
.apply {{ height: 22px; padding: 0 22px; display: inline-flex; align-items: center; font-size: 12.5px; color: #8a8a8a;
  border-radius: 3px; border: 1px solid #4b3a20; background: linear-gradient(#3a1410, #1e0806); }}
.tstrip {{ position: absolute; left: {STRIP_INSET}px; top: {(40 - STRIP) // 2}px; height: {STRIP}px;
  display: flex; align-items: center; padding: 0 {STRIP_PAD}px; background: {rgba(ADDON['bg'], A_STRIP)};
  font-family: Site, sans-serif; font-size: {STRIP_TEXT}px; color: {ADDON['body']}; white-space: nowrap; }}
'''


def window():
    trees, order = plan()
    ranks = ranks_of(order)
    planned, point = cells(order, ranks)
    reach = LEVEL - 9  # Progress.reachablePoints: points spent plus points waiting
    strip, unspent = strip_text(order, ranks, planned, point, trees)
    body = ''.join(tree_html(tab, tree, ranks, planned, reach) for tab, tree in enumerate(trees, 1))
    return (f'<div class="frame tw"><div class="portrait"></div><div class="title">Talents</div><span class="x">×</span>'
            f'<div class="tbar"><span class="tab on">Primary <i class="caret"></i></span><span class="tab off">Secondary <i class="lock"></i></span>'
            f'<span class="search">Search</span><span class="unspent">Unspent Talents <b>{unspent}</b></span></div>'
            f'<div class="trees">{body}</div>'
            f'<div class="foot"><div class="tstrip"><span>{strip}</span></div><span class="apply">Apply Changes</span></div></div>')


def board_talents():
    html = (f'<style>{CSS}</style><div class="note"><b>A level 33 hunter on Beast Mastery leveling, one point to '
            f'spend, one off plan.</b> The ring and the gold plate on the next talent, the plan\'s later points numbered and dimmed '
            f'past what the level allows, taken nodes bare, the strip at the foot.</div>{window()}')
    return hub.board2('<b>Store plate · Talent marks</b> · the whole window', html)


def main():
    from playwright.sync_api import sync_playwright
    OUT.mkdir(parents=True, exist_ok=True)
    with sync_playwright() as pw:
        exe = os.environ.get('PLAYWRIGHT_CHROMIUM')
        browser = pw.chromium.launch(executable_path=exe) if exe else pw.chromium.launch()
        page = browser.new_context(device_scale_factor=2, viewport={'width': 1800, 'height': 1200}).new_page()
        page.set_content(board_talents())
        page.evaluate('document.fonts.ready')
        page.wait_for_timeout(150)
        page.locator('#board').screenshot(path=str(OUT / 'talent-marks-board.png'))
        print(OUT / 'talent-marks-board.png')
        browser.close()


if __name__ == '__main__':
    main()
