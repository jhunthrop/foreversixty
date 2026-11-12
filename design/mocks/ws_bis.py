"""Web-layout mock: Best in slot on the in-game paperdoll (round 5, lane P2, 2026-11-09), the parity gate's source of truth.

Spec: design/specs/2026-11-09-directions.md, Paperdoll and Best in slot; the visual reference is the directions mock
(design/mocks/directions/pages.py `bis_frame`, png/g-bis-frame-*.png). The band holds the nav and the one control bar;
the frame's title bar is the page's h1 (the spec, the band and faction, the set's figure); its centre the selected
slot's client tooltip, ending in what it is worth, then Open in the sim. Beside the frame: a tank's or a healer's set
figures, What to stack (the top weights, what 1.00 is worth, the band's standing notes, the sim's weights page) and
Where to go (the stops). Below: the itinerary on one panel, the stops in route order, a row per piece.

States: design/mocks/web/bis-parity.json. Data: design/mocks/data/bis-facts.json (names, numbers and words read from the
built page by bis_facts.mjs; the mock decides every block and where it sits). Regions: design/parity.md.
    node design/mocks/bis_facts.mjs http://localhost:<port>   (a running `npm run preview -- --port <port>` in web/)
    python3 design/mocks/ws_bis.py  -> design/mocks/web/web-bis-<state>.html
FS_BIS_FACTS and FS_MOCK_OUT point both elsewhere: the parity spec renders its mocks at test time (design/parity.md).
"""
import json
import os
from html import escape
from pathlib import Path

from ws_lib import control_bar, sel, seg
from ws_prims import Facts, ItemRow, ItemTip, PaperdollFrame, Plate, Section, build_page, icon, write

HERE = Path(__file__).resolve().parent
CONFIG = json.loads((HERE / 'web/bis-parity.json').read_text())
# The facts: the committed data/bis-facts.json, or FS_BIS_FACTS (the parity spec's own, read from the build it serves).
FACTS = json.loads(Path(os.environ.get('FS_BIS_FACTS') or HERE / 'data/bis-facts.json').read_text())
QUALITY = {'poor': 0, 'common': 1, 'uncommon': 2, 'rare': 3, 'epic': 4, 'legendary': 5}

# The page's own layout (BisSpecPage, BisSide, BisItinerary measure the same); no surfaces: the primitives draw those.
BIS_CSS = '''
.bis-pg{padding-top:20px}
/* the built ControlBar wraps its groups with gap 12px 20px, the column gap 12 on the BiS bar (labels from 1920) */
.cbar[data-labels-from="1920"] .pc{gap:12px}
.cbar .seg a{height:40px}.cbar .seg a .sp{width:22px;height:22px;border-radius:4px;background:#10141d}
.src-btn{height:40px;padding:0 12px;gap:8px;font-size:14px;background:none;cursor:pointer}
.src-btn::after{width:14px;text-align:center}
.sel.fixed-field{width:var(--field-w);justify-content:space-between}
.bis-top{display:grid;grid-template-columns:minmax(0,1fr);gap:28px;align-items:start}
.bis-doll{width:100%;max-width:640px;justify-self:center}
@media(min-width:1100px){.bis-top{grid-template-columns:minmax(0,640px) minmax(0,1fr);gap:36px}.bis-doll{justify-self:start}}
@media(min-width:1760px){.bis-top{grid-template-columns:minmax(0,760px) minmax(0,560px);justify-content:center;column-gap:64px}.bis-doll{max-width:760px}}
.bis-side{display:flex;flex-direction:column;gap:28px;min-width:0}
.bis-side .sbox{padding:8px 20px}
.bis-side-body{display:flex;flex-direction:column;padding-bottom:8px}
.bis-side-line{margin-top:10px;color:var(--muted);font-size:13px;line-height:20px}
.bis-side-link{margin-top:12px;font-size:14px;font-weight:600}
.bis-stops{flex-direction:row;flex-wrap:wrap;gap:8px}
.bis-stop-chip{display:inline-flex;align-items:center;min-height:36px;padding:0 12px;border:1px solid #82705b;border-radius:999px;color:var(--gold);font-size:14px;font-weight:600;line-height:20px}
.bis-sim-link{font-size:14px;font-weight:600;white-space:nowrap}
.bis-route{margin-top:var(--section-gap)}
.bis-route .plate{padding:16px}
.bis-stop{break-inside:avoid;padding-bottom:18px}@media(min-width:1280px){.bis-stops-cols{columns:2;column-gap:40px}}
.bis-stop h2{display:flex;align-items:baseline;gap:10px;padding:4px 0 8px;font:400 22px/28px var(--font-title);letter-spacing:-.015em;color:#f3eee2}
.bis-stop h2 span{font:600 14px/1 var(--font-body);color:var(--text)}
.irow .tr .g .v{color:#32ae88}
'''


# The bar's one select is the race, a fixed 112 px field (BisSpecPage RACE_FIELD_PX, ControlBar `fieldWidth`): "Night Elf" fits,
# the rare long race name truncates in the field.
RACE_FIELD_PX = 112


def field_select(value: str, width: int) -> str:
    """ws_lib `sel` at a fixed field width (the built ControlBar select with `fieldWidth`), not its content's."""
    return sel(value).replace('class="sel"', f'class="sel fixed-field" style="--field-w:{width}px"', 1)


def sources_button(label: str) -> str:
    """The Sources group's button on the shared ui/FilterPopover (lane E): 40 px, 12 px padding, 14 px type, the chevron 8 px after."""
    return f'<button type="button" class="sel src-btn" aria-expanded="false" aria-haspopup="dialog">{label}</button>'


def controls(f: dict) -> str:
    c = f['controls']
    groups = []
    for g in c['groups']:
        if g['kind'] == 'segmented':
            items = [(escape(o['label']), o['on'], None, '<span class="sp"></span>' if o['icon'] else '') for o in g['options']]
            groups.append((g['label'], seg(items, label=g['label'])))
        else:
            groups.append((g['label'], field_select(escape(g['value'] or ''), RACE_FIELD_PX)))
    groups.append(('Sources', sources_button(escape(c['sources']))))
    summary = ' &middot; '.join(escape(o['label']) for g in c['groups'] for o in g['options'] if o['on'])
    return control_bar(groups, summary, label=c['label'], labels_from=1920)


def centre(f: dict) -> str:
    """The selected slot's client tooltip, ending in what it is worth, then Open in the sim."""
    c = f['centre']
    tip = c['tip']
    selected = next((s for s in f['slots'] if s['selected']), None)
    quality = next((r['quality'] for st in f['itinerary']['stops'] for r in st['rows'] if selected and r['item'] == selected['item']), 'common')
    html = ItemTip(tip['name'], quality, tip['lines'][:-len(tip['tail'])] if tip['tail'] else tip['lines'], tip['tail'], tip['foot']) if tip else ''
    if c['sim']:
        html += f'<a class="bis-sim-link" href="#">{escape(c["sim"])}</a>'
    return html


def side(f: dict) -> str:
    parts = []
    if f['set']:
        s = f['set']
        line = f'<p class="bis-side-line">{escape(s["line"])}</p>' if s['line'] else ''
        parts.append(Section(escape(s['title']), f'<div class="bis-side-body">{Facts([(escape(x["label"]), escape(x["value"])) for x in s["facts"]])}{line}</div>', boxed=True))
    st = f['stack']
    lines = ''.join(f'<p class="bis-side-line">{escape(x)}</p>' for x in st['lines'])
    link = f'<a class="bis-side-link" href="#">{escape(st["link"])}</a>' if st['link'] else ''
    facts = Facts([(escape(x['label']), escape(x['value'])) for x in st['facts']]) if st['facts'] else ''
    parts.append(Section(escape(st['title']), f'<div class="bis-side-body">{facts}{lines}{link}</div>', boxed=True))
    w = f['where']
    if w:
        # web BisSide: one jump chip per stop
        stops = ''.join(f'<a class="bis-stop-chip" href="#">{escape(s["title"])}</a>' for s in w['stops'])
        action = f'<p style="margin-top:10px"><a class="bis-side-link" href="#">{escape(w["action"])}</a></p>' if w['action'] else ''
        parts.append(Section(escape(w['title']), f'<div class="bis-side-body bis-stops">{stops}</div>{action}', boxed=True))
    return f'<div class="bis-side">{"".join(parts)}</div>'


def gain(g: dict | None) -> tuple[str, str, str] | None:
    """The row's figure as the built Figure draws it: its value and unit."""
    return (escape(g['value']), escape(g['unit']), '') if g else None


def itinerary(f: dict) -> str:
    it = f['itinerary']
    stops = '<div class="bis-stops-cols">' + ''.join(
        f'<section class="bis-stop"><h2>{escape(s["title"])}{f"<span>{escape(s["range"])}</span>" if s["range"] else ""}</h2>'
        + ''.join(ItemRow(escape(r['name']), QUALITY.get(r['quality'], 1), icon(r['icon'], 40) if r['icon'] else '<span class="ic"></span>',
                          escape(r['line']), gain=gain(r['gain'])) for r in s['rows'])
        + '</section>' for s in it['stops']) + '</div>'
    return f'<div class="bis-route">{Plate(stops, label=escape(it["label"]))}</div>'


def render(f: dict) -> str:
    figure = (f['figure']['value'], f['figure']['unit']) if f['figure'] else None
    doll = PaperdollFrame(f['slots'], f['title'], f['meta'], figure, centre(f), f['frameLabel'], list_centre=centre(f))
    body = (controls(f) + f'<div class="pc bis-pg"><div class="bis-top"><div class="bis-doll">{doll}</div>{side(f)}</div>{itinerary(f)}</div>'
            + f'<style>{BIS_CSS}</style>')
    return build_page(f'{f["title"]} (mock, {f["id"]})', body, nav='Tools',
                      note='Parity mock (round 5, the paperdoll). Data from design/mocks/data/bis-facts.json; layout from the 2026-11-09 directions spec.')


if __name__ == '__main__':
    for state in CONFIG['states']:
        print(write(state['mock'], render(FACTS[state['id']])))
