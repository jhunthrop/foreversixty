"""Web-layout mock 1 (rev 4, 2026-11-06): Best in slot as approved and shipped, the parity gate's source of truth.

Spec: design/specs/2026-11-05-web-layout.md section 5 and the build decisions in 5.13 (option C rows: line 1 slot,
level, icon, name and two short stats, the gain figure; line 2 the full source alone; the gain column 144 px with
what the figure is against on one line; Sources a compact button; group labels from 1920; 360 px inspector from 1760).
Rev 3 (persona fixes, three-line rows, the inline next best) is superseded by 5.13; its renders stay in design/reviews.

The head answer is one figure with its unit, the trust line and the links. Nothing else: no disclosure, no metrics
list, no explanatory note (the holy-priest drift of 2026-11-06 is not the design). Tanks show one figure, damage taken
per second (the tiers' tank sort); an owner decision may change which.

States: design/mocks/web/bis-parity.json. Data: design/mocks/data/bis-facts.json (names, numbers and sources read
from the built page by bis_facts.mjs; the mock decides every block and where it sits). Regions: design/parity.md.
    python3 design/mocks/ws_bis.py  -> design/mocks/web/web-bis-<state>.html
"""
import json
from html import escape
from pathlib import Path

from mocklib import CLASS_COLOR
from ws_lib import REGION, control_bar, crest_img, icon, sel, seg
from ws_prims import Meter, PageHeader, Plate, Section, build_page, write

HERE = Path(__file__).resolve().parent
CONFIG = json.loads((HERE / 'web/bis-parity.json').read_text())
FACTS = json.loads((HERE / 'data/bis-facts.json').read_text())
RARITY = {'poor': '#9d9d9d', 'common': '#ffffff', 'uncommon': '#1eff00', 'rare': '#3d94f0', 'epic': '#b866f5', 'legendary': '#ff8000'}
TANK_FIGURE = ('dtps', 'DTPS')  # the tank answer: one figure (id in the facts, unit shown)

# The page's own rules (the built Paperdoll, BisAnswer, BisWeights and summary measure the same).
BIS_CSS = '''
/* The BiS head keeps a fixed floor from 1024 so the title and the controls hold their y across specs
   (one figure, a healer's metrics, a tank's four): BisSpecPage HEAD_ROW_MIN_HEIGHT, b66cf09b8. */
@media(min-width:1024px){.head{min-height:calc(165px + 36px)}} /* row floor plus the head padding 20/16; the mock .head is the row and is border-box */
.bis-pg{padding-top:20px}
.ans .bis-ans{display:flex;flex-direction:column;align-items:flex-end;gap:4px;text-align:right;max-width:640px}
.bis-ans .fig .v{font-size:36px;line-height:40px}.bis-ans .fig small{font-size:13px}
.bis-ans .cap{font-size:12px;line-height:16px;color:var(--muted)}
.bis-line{font-size:13px;line-height:18px;color:var(--muted)}.bis-links{font-size:14px;font-weight:600}
.cbar[data-labels-from="1920"] .pc{column-gap:12px}
.cbar .seg a{height:40px}.cbar .seg a .sp{width:22px;height:22px;border-radius:4px;background:#10141d;border:1px solid var(--line)}
.src-btn{height:36px;padding:0 10px;font-size:13px;background:none;cursor:pointer}
.bis-sum{display:flex;flex-wrap:wrap;align-items:center;justify-content:space-between;gap:10px 16px;padding:10px 0}
.bis-sum p{font-size:15px;line-height:24px;min-width:0}.bis-sum p a,.bis-sum p button{color:var(--gold);font-weight:600}
.bis-sum p button{background:none;border:0;padding:0;margin-left:6px;font:inherit;font-weight:600;cursor:pointer}
.bis-sum .seg a{height:44px}
.bis-note{color:var(--muted);font-size:13px;line-height:20px;margin-bottom:8px}
.doll-cols{display:grid;grid-template-columns:minmax(0,1fr)}
@media(min-width:1024px){.doll-cols{grid-template-columns:minmax(0,1fr) minmax(0,1fr)}.doll-col+.doll-col{border-left:1px solid var(--line-soft)}
 .doll[data-view="where"] .doll-cols{display:block;columns:2;column-gap:0}}
.doll-grp{break-inside:avoid}.doll-grp-h{display:flex;align-items:baseline;gap:12px;padding:14px 16px 6px}
.doll-grp-h b{font:700 18px/24px var(--font-body);color:var(--strong)}.doll-grp-h span{margin-left:auto;font-size:13px;color:var(--muted)}
.cell{display:grid;grid-template-columns:64px 40px minmax(0,1fr) auto;column-gap:12px;align-items:start;min-height:58px;padding:9px 16px;border-bottom:1px solid var(--line-soft)}
.cell .lb{display:flex;flex-direction:column;padding-top:1px;font:600 13px/18px var(--font-body);color:var(--muted)}
.cell .lb .fl{font-size:12px;line-height:16px;font-weight:700}
.cell .bd{display:flex;flex-direction:column;gap:1px;min-width:0}
.cell .l1{display:flex;align-items:baseline;column-gap:10px;min-width:0}.cell .l1 .nm{flex:none;max-width:100%;font:700 15px/20px var(--font-body)}
.cell .l1 .st{min-width:0;overflow:hidden;text-overflow:ellipsis;white-space:nowrap;font-size:13px;line-height:20px;color:var(--muted)}
.cell .l2{font-size:13px;line-height:18px;color:var(--muted);overflow-wrap:anywhere}
.cell .gn{display:flex;flex-direction:column;align-items:flex-end;width:144px;min-width:0;justify-self:end}
.cell .gn .v{font:700 16px/20px var(--font-body);font-variant-numeric:tabular-nums;color:var(--strong)}.cell .gn .v small{font-size:13px;font-weight:600;color:var(--muted);margin-left:4px}
.cell .gn .bs{max-width:100%;overflow:hidden;text-overflow:ellipsis;white-space:nowrap;font-size:12px;line-height:16px;color:var(--muted);text-align:right}
.cell .ic0{width:40px;height:40px;border-radius:5px;border:1px dashed var(--line-soft)}
.bis-stack{margin-top:32px}
.wt-s{font-size:15px;line-height:24px}
.wt-cols{display:grid;grid-template-columns:repeat(auto-fit,minmax(150px,1fr));gap:16px;margin-top:12px}
.wt{display:flex;flex-direction:column;gap:6px;min-width:0;font-size:14px;line-height:1.5}.wt.faint{opacity:.6}.wt-n{font-size:15px;font-weight:600}.wt .fig.sm .v{font-size:16px;line-height:20px}
.wt-l{font-size:13px;line-height:20px;margin-top:8px;color:var(--muted)}.wt-l a{font-weight:600}
.bis-guide{color:var(--muted);font-size:13px;line-height:20px;margin:var(--section-gap) 0 8px}.bis-guide a{font-weight:600}
.sec-h+.sec-b{padding-top:12px}
'''


def colored(text: str, rarity: str) -> str:
    return f'<span style="color:{RARITY.get(rarity, RARITY["common"])}">{escape(text)}</span>'


def answer(f: dict) -> str:
    a = f['answer']
    if a['figure']:
        fig = a['figure']
        label, value, unit, caption = fig['label'], fig['value'], fig['unit'], fig['caption']
    else:
        tank = next(t for t in a['tank'] if t['id'] == TANK_FIGURE[0])
        label, value, unit, caption = 'This set', tank['value'], TANK_FIGURE[1], ''
    ranking = f' &middot; <a href="#">{escape(a["rankingSim"])}</a>' if a['rankingSim'] else ''
    links = ' &middot; '.join(f'<a href="#">{escape(t)}</a>' for t in (a['simLink'], a['plannerLink']) if t)
    cap = f'<span class="cap">{escape(caption)}</span>' if caption else ''
    return (f'<div class="bis-ans"><div class="fig"><span class="t-label">{escape(label)}</span>'
            f'<span class="v">{escape(value)}<small>{escape(unit)}</small></span>{cap}</div>'
            f'<p class="bis-line">{escape(a["trust"])}{ranking}</p><p class="bis-links">{links}</p></div>')


def controls(f: dict) -> str:
    c = f['controls']
    groups = []
    for g in c['groups']:
        if g['kind'] == 'segmented':
            items = [(escape(o['label']), o['on'], None, '<span class="sp"></span>' if o['icon'] else '') for o in g['options']]
            groups.append((g['label'], seg(items, label=g['label'])))
        else:
            groups.append((g['label'], sel(escape(g['value'] or ''))))
    groups.append(('Sources', f'<button type="button" class="sel src-btn" aria-expanded="false" aria-haspopup="dialog">{escape(c["sources"])}</button>'))
    summary = ' &middot; '.join(escape(o['label']) for g in c['groups'] for o in g['options'] if o['on'])
    return control_bar(groups, summary, label=c['label'], labels_from=1920)


def summary(f: dict) -> str:
    s = f['summary']
    nxt = f'<button type="button">{escape(s["next"])}</button>' if s['next'] else ''
    view = seg([(escape(o['label']), o['on'], None, '') for o in s['view']], label='View')
    return (f'<div class="bis-sum" data-region="{REGION["summary"]}"><p>{escape(s["lead"])} <a href="#">{escape(s["signIn"])}</a> '
            f'{escape(s["after"])}{nxt}</p>{view}</div>')


def notes(f: dict) -> str:
    if not f['notes']:
        return ''
    return f'<div data-region="{REGION["notes"]}">' + ''.join(f'<p class="bis-note">{escape(n)}</p>' for n in f['notes']) + '</div>'


def cell(c: dict) -> str:
    flag = f'<span class="fl">{escape(c["flag"]["text"])}</span>' if c['flag'] else ''
    lead = f'<span class="lb">{escape(c["label"])}{flag}</span>'
    if c['state'] == 'empty':
        return (f'<div class="cell" role="option" aria-disabled="true">{lead}<span class="ic0"></span><span class="bd">'
                f'<span class="l1"><span class="nm muted" style="font-weight:400">{escape(c["name"])}</span></span><span class="l2">{escape(c["line2"])}</span></span><span></span></div>')
    q = {'poor': 'q0', 'uncommon': 'q2', 'rare': 'q3', 'epic': 'q4', 'legendary': 'q5'}.get(c['quality'], '')
    ic = icon(c['icon'], 40, q) if c['icon'] else '<span class="ic0"></span>'
    stats = f'<span class="st">{escape(c["stats"])}</span>' if c['stats'] else ''
    g = c['gain'] or {'value': '', 'unit': ''}
    unit = f'<small>{escape(g["unit"])}</small>' if g['unit'] else ''
    b = c['basis']
    basis = f'<span class="bs">{escape(b["prefix"])} {colored(b["item"], b["quality"]) if b["item"] else ""}</span>' if b else ''
    return (f'<div class="cell" role="option" aria-selected="false">{lead}{ic}<span class="bd"><span class="l1">'
            f'<span class="nm" style="color:{RARITY.get(c["quality"], RARITY["common"])}">{escape(c["name"])}</span>{stats}</span>'
            f'<span class="l2">{escape(c["line2"])}</span></span><span class="gn"><span class="v">{escape(g["value"])}{unit}</span>{basis}</span></div>')


def doll(f: dict) -> str:
    d = f['doll']
    if d['view'] == 'where':
        inner = ''.join(f'<div class="doll-grp" role="group" aria-label="{escape(g["title"])}"><div class="doll-grp-h"><b>{escape(g["title"])}</b>'
                        f'<span>{escape(g["count"])}</span></div>{"".join(cell(c) for c in g["cells"])}</div>' for g in d['groups'])
    else:
        inner = ''.join(f'<div class="doll-col">{"".join(cell(c) for c in d[side])}</div>' for side in ('left', 'right'))
    return (f'<div class="doll" data-view="{d["view"]}" data-region="{REGION["slots"]}">'
            f'<div class="doll-cols" role="listbox" aria-label="{escape(d["label"])}">{inner}</div></div>')


def weight(r: dict) -> str:
    if r['faint']:
        return f'<li class="wt faint"><span class="wt-n">{escape(r["name"])}</span><span class="muted">{escape(r["line"])}</span></li>'
    return (f'<li class="wt"><span class="wt-n">{escape(r["name"])}</span><span class="fig sm"><span class="v">{escape(r["scale"])}</span></span>'
            f'{Meter(r["bar"])}<span class="muted">{escape(r["line"])}</span></li>')


def weights(f: dict) -> str:
    """Always open, one column per stat (spec 5.3 item 6; the 5.10 fold retired on main 2026-11-06)."""
    w = f['weights']
    cols = f'<p class="wt-l">{escape(w["caption"])}</p><ul class="wt-cols" aria-label="Stat weights">{"".join(weight(r) for r in w["rows"])}</ul>' if w['rows'] else ''
    exp = f'<p class="wt-l">{escape(w["expertise"])}</p>' if w['expertise'] else ''
    body = (f'<div class="sec-b"><p class="wt-s">{escape(w["sentence"])}</p>{cols}{exp}'
            f'<p class="wt-l">{escape(w["addon"])} <a href="#">{escape(w["addonLink"])}</a></p></div>')
    return f'<div class="bis-stack">{Section(w["title"], body)}</div>'


def render(f: dict) -> str:
    color = CLASS_COLOR[f['class']]
    head = PageHeader(escape(f['title']), title_color=color, crest=crest_img(f['class'], 56), meta=[escape(f['meta'])], answer=answer(f))
    guide = f'<p class="bis-guide" data-region="{REGION["footer_note"]}">{escape(f["guide"]["lead"])} <a href="#">{escape(f["guide"]["link"])}</a></p>'
    body = (head + controls(f) + f'<div class="pc bis-pg">{summary(f)}{notes(f)}{Plate(doll(f), label='Gear')}{weights(f)}{guide}</div>'
            + f'<style>{BIS_CSS}</style>')
    return build_page(f'{f["title"]} best in slot (mock, {f["id"]})', body, nav='Tools',
                      note='Parity mock (rev 4). Data from design/mocks/data/bis-facts.json; layout from spec 5 and 5.13.')


if __name__ == '__main__':
    for state in CONFIG['states']:
        print(write(state['mock'], render(FACTS[state['id']])))
