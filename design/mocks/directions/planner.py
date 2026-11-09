"""Mock batch 2: the talent planner in the directions system (plan 2026-11-09-redesign).

Data, never invented: design/mocks/data/planner-facts.json, which design/mocks/planner_facts.mjs writes by running the
planner's own models (lib/planner/*) on the active build: every talent, rank, description, point order row and copy line.
The shown state is the facts' `guide` state, the published Survival Hunter level 60 build (Survival fights in melee in
Forever). Drawn by prims.py: the talent grid, point order, panels, fields and the client tooltip.
"""
import json
from html import escape
from pathlib import Path

import data as D
from common import write
from prims import PRIMS_CSS, field, panel, point_order, talent_trees
from style import footer, header, page, titlebar

FACTS = json.loads((Path(__file__).resolve().parent.parent / 'data/planner-facts.json').read_text())
STATE = FACTS['states']['guide']
P, C = FACTS['copy']['planner'], FACTS['copy']
VIEWER = ('Felsnap', 'hunter', 34)  # the facts' signed-in character: Send this build goes to Felsnap
HOVER = 'Lethal Attacks'  # the hovered talent: the mock's one open tooltip

PLANNER_CSS = PRIMS_CSS + '''
.pl{padding-top:6px}.tfoot{display:flex;justify-content:space-between;align-items:center;padding-top:14px;font-size:14px}
.tip .gold{color:#ffd100}.shl{display:block;font-size:12.5px;color:var(--muted);margin-bottom:6px}
@media(max-width:700px){.pl aside{order:-1}.pl .twocol{display:flex;flex-direction:column}}
'''


def talent_tip(c: dict) -> str:
    """The client's talent tooltip: name, rank, this rank's text in gold, then the next rank's."""
    now = f'<div class="gold" style="margin-top:4px">{escape(c["now"])}</div>' if c['now'] else ''
    nxt = f'<div style="margin-top:8px">Next rank:</div><div class="gold">{escape(c["next"])}</div>' if c['next'] else ''
    return f'<div class="tip"><div class="tl"><span>{escape(c["name"])}</span></div><div>Rank {c["rank"]}/{c["max"]}</div>{now}{nxt}</div>'


def trees_html() -> str:
    trees, hover = [], None
    for i, t in enumerate(STATE['trees']):
        cells = []
        for c in t['cells']:
            cells.append((c['tier'], c['column'], D.icon_uri(c['icon']), c['rank'], c['max']))
            if c['name'] == HOVER:
                hover = (i, c['tier'], c['column'], talent_tip(c))
        trees.append((t['name'], t['points'], cells))
    return talent_trees(trees, hover)


def order_html() -> str:
    pts = STATE['points']
    rows = []
    for r, lo in enumerate(range(10, 61, 10)):
        socks = []
        for lvl in range(lo, min(lo + 9, 60) + 1):
            i = lvl - 10
            socks.append((D.icon_uri(pts[i]['icon']), pts[i]['rank'], pts[i]['tip']) if i < len(pts) else None)
        rows.append((STATE['order_rows'][r], socks))
    return point_order(rows)


def compare_html() -> str:
    st = STATE['compare']['status']
    line = f'<span class="num strong">{st["count"]}</span> {escape(C["band"]["differsSuffix"])}'
    return f'<p>{line}</p><p style="margin-top:10px"><a href="#">{escape(C["band"]["seeWhatDiffers"])} &rsaquo;</a></p>'


def in_game_html() -> str:
    send = FACTS['send']['idle']
    return (f'<p><a class="btn" href="#">{escape(send["label"])}</a></p>'
            f'<p style="margin-top:12px"><a class="btn2" href="#">{escape(C["addon"]["copyForGame"])}</a></p>')


def share_html() -> str:
    return (f'<span class="shl">{escape(P["shareTitleLabel"])}</span>'
            + field(P['shareTitlePlaceholder'], f'<a class="btn2" href="#">{escape(P["share"])}</a>', P['shareTitleLabel']))


def planner_page() -> str:
    classes = [c['name'] if isinstance(c, dict) else c for c in FACTS['classes']]
    cls_name = STATE['className']
    band = titlebar(STATE['cls'], STATE['title'], escape(STATE['meta']), right=f'<b class="num">{STATE["spent"]}</b> of {STATE["available"]} points',
                    ctl=[[(n, n == cls_name) for n in classes], [(f'{STATE["race"]} &#9662;', False), ('Level 60 &#9662;', False)]],
                    top=header(*VIEWER, 'Planner')).replace('&amp;#9662;', '&#9662;')
    foot = f'<div class="tfoot"><span></span><a class="muted" style="font-weight:400" href="#">{escape(P["reset"])}</a></div>'  # the band carries the points
    main = panel('Talents', trees_html() + foot, framed=True) + panel(P['pointOrderTitle'], order_html())
    rail = (panel(C['band']['sharedHeading'], compare_html()) + panel(P['inGameTitle'], in_game_html())
            + panel(P['shareTitle'], share_html()))
    body = band + f'<main class="w pl"><div class="twocol"><div>{main}</div><aside>{rail}</aside></div></main>' + footer()
    return page('Forever Sixty, talent planner', PLANNER_CSS, body)


def build() -> None:
    write('b2-planner', planner_page())


if __name__ == '__main__':
    build()
