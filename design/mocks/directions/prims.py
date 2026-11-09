"""The shared drawn primitives (owner, 2026-11-08: "we should use shared components and resources as a first class
citizen"). Every directions surface draws its panels, tables, player cells, meters, facts, row lists and charts from
here, once; a surface that needs something new adds it here first. The inventory and the web component each one
becomes are in docs/superpowers/plans/2026-11-09-redesign.md, "Shared component inventory".
"""
import math
from html import escape

import data as D
from style import medallion

# The chart palette: the page's own tokens, so a chart reads as part of the panel, not a pasted widget.
GRID = 'rgba(255,255,255,.08)'
AXIS = 'rgba(255,255,255,.5)'
DEATH = '#f0736b'
CHART_W = 1000  # the SVG's own coordinate width; it scales to the panel

PRIMS_CSS = '''
/* panel: the one card with a heading row (title, one meta line, one action at the right) */
.pn{margin-top:22px;overflow:clip}.pn>.hd{display:flex;align-items:baseline;gap:12px;padding:16px 20px 10px}
.pn>.hd h2{font-size:21px;line-height:26px}.pn>.hd .m{font-size:13px;color:var(--muted)}.pn>.hd .a{margin-left:auto;font-size:14px;white-space:nowrap}
.pn>.bd{padding:0 20px 16px}.pn>.bd.flush{padding:0 0 6px}
/* data table: tabular figures right-aligned; the first column a player or ability cell; the header sticks (the panel clips,
   not hides, so it sticks to the viewport) */
.tb{width:100%;border-collapse:collapse;font-size:14px}
.tb th{position:sticky;top:0;z-index:2;background:var(--card);font:600 12px var(--body);color:var(--muted);text-align:left;padding:8px 10px;border-bottom:1px solid rgba(255,255,255,.14);white-space:nowrap}
.tb th.n,.tb td.n{text-align:right;font-variant-numeric:tabular-nums lining-nums;white-space:nowrap}
.tb th.on{color:#fff}
.tb td{padding:7px 10px;border-bottom:1px solid rgba(255,255,255,.07);color:var(--text);vertical-align:middle}
.tb tr:last-child td{border-bottom:0}.tb td.strong{color:#fff;font-weight:600}
.tb tr.tot td{color:#fff;font-weight:600;border-top:1px solid rgba(255,255,255,.14)}
.tb .first{padding-left:20px}.tb .last{padding-right:20px}
.tb tr.me td{background:rgba(14,134,202,.14)}.tb tr.me td.first{box-shadow:inset 3px 0 0 var(--blue)}
/* a table with a meter column (bar_col): the meter and the name column take one width on every tab, so bars never jump */
.tb td.mc{width:168px}.tb td.mc .mt{width:148px}.tb th.first:has(+ th.mc),.tb td.first:has(+ td.mc){width:250px}
/* a player (or ability) cell: medallion or icon, the name in class colour, one sub line */
.pc{display:flex;align-items:center;gap:10px;min-width:0}.pc .nm{display:block;font-weight:600;line-height:19px;white-space:nowrap}
.pc .sb{display:block;font-size:12.5px;color:var(--muted);line-height:17px;white-space:nowrap}
/* meter: the share bar inside a table row, in the row's class colour */
.mt{display:block;height:6px;background:rgba(255,255,255,.07);min-width:90px}.mt i{display:block;height:100%}
/* facts: label and value pairs, one per line */
.fx{display:grid;grid-template-columns:auto 1fr;gap:6px 16px;font-size:14px}.fx dt{color:var(--muted)}.fx dd{margin:0;color:#fff;text-align:right}
/* row list: lead, main with a sub line, trail */
.rl li{display:grid;grid-template-columns:auto minmax(0,1fr) auto;gap:10px;align-items:center;padding:8px 0;border-bottom:1px solid rgba(255,255,255,.07)}
.rl li:last-child{border-bottom:0}.rl .mn>b{display:block;font-weight:600;line-height:19px}.rl .mn>span{display:block;font-size:13px;color:var(--muted)}
.rl .tr{text-align:right;font-variant-numeric:tabular-nums;color:#fff;white-space:nowrap}
/* chart frame: the modern line chart in the page's palette; legend chips carry the figure and its unit */
.cf svg{display:block;width:100%;height:auto}.cf .lg{display:flex;flex-wrap:wrap;gap:6px 18px;padding:10px 0 0;font-size:13px}
.cf .lg span{display:flex;align-items:center;gap:7px;color:var(--text)}.cf .lg i{width:12px;height:3px;border-radius:2px}
.cf .lg b{color:#fff;font-weight:600;font-variant-numeric:tabular-nums}
.empty{color:var(--muted);font-size:14px;padding:6px 0}.pn>.hd0{height:10px}
/* chips: a panel's own filter row (what a chart or lane view shows), toggles in the control style */
.chips{display:flex;flex-wrap:wrap;gap:8px;padding:0 0 12px}.chips a{font:600 13px var(--body);color:var(--text);padding:4px 11px;
border:1px solid rgba(255,255,255,.18);border-radius:var(--r-sm)}.chips a.on{color:#fff;border-color:var(--blue);background:rgba(0,147,178,.16)}
.chips a i{display:inline-block;width:10px;height:10px;border-radius:2px;margin-right:7px;vertical-align:-1px}
/* code: SQL and file names, on the panel's own black */
.code{margin:0;padding:12px 14px;background:#0b0705;border:1px solid rgba(255,255,255,.12);border-radius:var(--r-sm);
font:13px/1.6 ui-monospace,Menlo,monospace;color:#fff;white-space:pre-wrap;overflow-x:auto}
/* the rank colour ladder's one-line legend */
.plg{display:flex;flex-wrap:wrap;gap:4px 12px;font:700 13px var(--body)}
.om{display:none}@media(max-width:700px){.od{display:none}.om{display:block}}
/* ability icon: a spell or item icon at row size, the client's thin black edge */
.ai{display:inline-block;flex:none;border-radius:var(--r-ic);border:1px solid #000;box-shadow:0 0 0 1px var(--rule);vertical-align:middle}
/* settings bar: the run settings under the band, each a value with its picker, the one action at the end */
.setb{display:flex;flex-wrap:wrap;align-items:center;gap:8px 22px;padding:10px 18px;margin-top:22px}.setb .f{display:flex;flex-direction:column;font-size:12px;color:var(--muted)}
.setb .f b{font:600 15px var(--body);color:#fff}.setb .f b::after{content:" \\25BE";color:var(--muted);font-size:11px}.setb .go{margin-left:auto}
/* field: one input with its button, the client's dark edit box */
.fld{display:flex;gap:10px}.fld input{flex:1;min-width:0;font:15px var(--body);color:#fff;background:#0b0c10;border:1px solid var(--rule);border-radius:var(--r-sm);padding:10px 12px}
.fld input::placeholder{color:var(--muted)}
/* histogram: runs per DPS bin, the mean in gold, the middle two thirds shaded */
.hg svg{display:block;width:100%;height:auto}.hg .ax{display:flex;justify-content:space-between;font-size:12px;color:var(--muted);padding-top:4px}
/* lane timeline: one lane per row label, events as ticks, auras as bars, one time axis */
.ln{display:grid;grid-template-columns:var(--lw,170px) minmax(0,1fr);gap:0 12px}.ln .lb{display:flex;align-items:center;gap:8px;height:30px;min-width:0;font-size:13.5px;font-weight:600;color:#fff;overflow:hidden;white-space:nowrap;text-overflow:ellipsis}
.ln svg{display:block;width:100%;height:30px}.ln .ax{grid-column:2;position:relative;height:20px;font-size:12px;color:var(--muted)}
.ln .ax span{position:absolute;top:2px;transform:translateX(-50%)}.ln .ax span:first-child{transform:none}
@media(max-width:700px){.setb{gap:8px 16px;padding:10px 12px}.setb .go{margin-left:0;width:100%}.setb .go .btn{width:100%;text-align:center}.ln{--lw:104px}}
/* talent trees: the client's grid, four columns by seven tiers; a spent talent lit with its rank in the corner, an
   unspent one dimmed; the tree name and its points over each */
.tt{display:grid;grid-template-columns:repeat(3,minmax(0,1fr));gap:18px}.tt h3{font:400 18px var(--title);color:var(--cream);margin-bottom:8px}
.tt h3 span{font:600 14px var(--body);color:var(--muted);margin-left:6px}
.tg{display:grid;grid-template-columns:repeat(4,40px);grid-auto-rows:40px;gap:10px 12px;padding:12px;background:rgba(0,0,0,.28);border-radius:var(--r-sm);width:max-content}
.tc{position:relative;width:40px;height:40px}.tc img{width:40px;height:40px;border-radius:var(--r-ic);border:1px solid #000;box-shadow:0 0 0 1px var(--rule)}
.tc.off img{filter:grayscale(1) brightness(.45)}.tc.on img{box-shadow:0 0 0 1px #1eff00}.tc.max img{box-shadow:0 0 0 1px var(--gold)}
.tc b{position:absolute;right:-6px;bottom:-6px;font:700 11px/1 var(--body);color:#1eff00;background:#000;border-radius:3px;padding:2px 3px}
.tc.max b{color:var(--gold)}.tc.hv img{outline:2px solid #fff;outline-offset:1px}
.tc .ttip{position:absolute;left:52px;top:-4px;z-index:3;width:280px}
@media(max-width:1100px){.tt{grid-template-columns:1fr}}
/* phone: the hover tooltip becomes the tap sheet (not drawn here); table heads may wrap rather than push a column off */
@media(max-width:700px){.tc .ttip{display:none}.tb th{white-space:normal}
/* a row whose trail is actions (Award, Accept and Decline) puts them under the row on a phone; a figure trail stays */
.rl li:has(.tr .btn,.tr .btn2){grid-template-columns:auto 1fr}.rl li:has(.tr .btn,.tr .btn2) .tr{grid-column:1/-1;text-align:left}}
/* point order: a row per ten levels, one socket per level with the talent taken there and its rank */
.po li{display:grid;grid-template-columns:96px 1fr;gap:10px;align-items:center;padding:6px 0;border-bottom:1px solid rgba(255,255,255,.07)}
.po li:last-child{border-bottom:0}.po .lv{font-size:12.5px;color:var(--muted)}.po .sk{display:flex;flex-wrap:wrap;gap:4px}
.po .s{position:relative;width:28px;height:28px}.po .s img{width:28px;height:28px;border-radius:var(--r-ic);border:1px solid #000;box-shadow:0 0 0 1px var(--rule)}
.po .s b{position:absolute;right:-3px;bottom:-4px;font:700 10px/1 var(--body);color:#1eff00;background:#000;border-radius:2px;padding:1px 2px}
.po .s.e{border:1px dashed rgba(255,255,255,.2);border-radius:var(--r-ic)}
/* reading section: a heading, the prose at a reading measure, and its side block (the facts that prose is about) */
.rs{display:grid;grid-template-columns:minmax(0,1fr) minmax(0,380px);gap:28px;padding:22px 24px;align-items:start}
.rs+.rs{border-top:1px solid rgba(255,255,255,.1)}.rs h2{font-size:24px;margin-bottom:10px;grid-column:1/-1}
.prose{max-width:72ch;font-size:16px;line-height:1.65}.prose p+p{margin-top:12px}.prose b,.prose strong{color:#fff}
.prose ul{list-style:none}.prose li{padding:3px 0 3px 18px;position:relative}.prose li::before{content:"";position:absolute;left:3px;top:13px;width:6px;height:6px;border-radius:50%;background:var(--rule-warm)}
.side{border-left:1px solid rgba(255,255,255,.12);padding:2px 0 2px 18px;overflow-x:auto}.side h3{font:600 13px var(--body);color:var(--muted);margin-bottom:6px;text-transform:none}
.side+.side{margin-top:14px}
@media(max-width:1100px){.rs{grid-template-columns:minmax(0,1fr);padding:18px 16px}}
/* table of contents: the page's sections, the current one marked */
.toc a{display:block;padding:6px 0 6px 12px;color:var(--text);font-weight:400;border-left:2px solid rgba(255,255,255,.12)}.toc a.on{color:#fff;font-weight:600;border-left-color:var(--blue)}
/* tag: a small state word (Open, Awarded, Officer, Kill); one shape, a tone per meaning */
.tag{display:inline-block;font:600 12px/1 var(--body);padding:4px 7px;border-radius:var(--r-ic);border:1px solid rgba(255,255,255,.22);color:var(--text);white-space:nowrap}
.tag.gold{color:var(--gold);border-color:rgba(255,236,113,.45)}.tag.kill{color:var(--gain);border-color:rgba(50,174,136,.55)}
.tag.wipe{color:#f0736b;border-color:rgba(240,115,107,.5)}
/* answers: a raid sign-up's In / Late / Bench / Out, the chosen one lit */
.ans{display:inline-flex;border:1px solid var(--rule);border-radius:var(--r-sm);overflow:hidden}.ans a{font:600 13px var(--body);color:var(--text);padding:6px 12px}
.ans a+a{border-left:1px solid var(--rule)}.ans a.on{background:var(--teal);color:#fff}
/* need grid: what a raid needs against what it has; a cell is ok, short or missing */
.ng td.c{text-align:center;font-weight:700;width:56px}.ng .ok{color:var(--gain)}.ng .warn{color:var(--gold)}.ng .miss{color:#f0736b}.ng .none{color:var(--muted)}
/* item row: icon with its quality edge, the name in quality colour, one sub line, one trail */
.ir{display:grid;grid-template-columns:36px minmax(0,1fr) auto;gap:10px;align-items:center;padding:7px 0;border-bottom:1px solid rgba(255,255,255,.07)}
.ir:last-child{border-bottom:0}.ir img{width:36px;height:36px;border-radius:var(--r-ic);border:1px solid #000;box-shadow:0 0 0 1px var(--qb,var(--rule))}
.ir .n{display:block;font-weight:600;line-height:19px}.ir .s{display:block;font-size:12.5px;color:var(--muted);line-height:17px}.ir .t{font-size:13px;color:var(--muted);white-space:nowrap}
.twocol{display:grid;grid-template-columns:minmax(0,1fr) 340px;gap:28px;align-items:start}
@media(max-width:1100px){.twocol{grid-template-columns:minmax(0,1fr)}.twocol.ans>aside{order:-1}}
/* a report or rankings band (.rb) on a phone: its tab group and its pickers scroll as one strip, opening on the selected
   tab (strip_into_view), so the band stays three short lines */
@media(max-width:700px){.rb .ctl{flex-wrap:nowrap;overflow-x:auto;scrollbar-width:none;-webkit-mask-image:linear-gradient(90deg,#000 85%,transparent);
mask-image:linear-gradient(90deg,#000 85%,transparent)}.rb .ctl .g{flex:none;overflow:visible;max-width:none;-webkit-mask-image:none;mask-image:none}
.pc .sb{white-space:normal}.pc .sb .om{display:inline}}
@media(max-width:700px){.pn>.hd{padding:14px 14px 8px}.pn>.bd{padding:0 14px 14px}.tb .first{padding-left:14px}.tb .last{padding-right:14px}
.tb .hp{display:none}.cf .lg{font-size:12px}}
'''


def panel(title: str, body: str, meta: str = '', action: str = '', framed: bool = False, ornate: bool = False,
          flush: bool = False) -> str:
    """The one panel. `flush` lets a table run edge to edge; `framed` and `ornate` are for the page's featured panels
    only (two or three a page)."""
    cls = 'card pn' + (' framed' if framed else '') + (' ornate' if ornate else '')
    # no title: the band already names what the panel holds (a list page's one panel), so the panel starts at its body
    hd = (f'<div class="hd"><h2>{escape(title)}</h2>{f"<span class=m>{meta}</span>" if meta else ""}'
          f'{f"<span class=a>{action}</span>" if action else ""}</div>') if title else '<div class="hd0"></div>'
    return f'<section class="{cls}">{hd}<div class="bd{" flush" if flush else ""}">{body}</div></section>'


def player_cell(slug: str | None, name: str, sub: str = '', size: int = 28) -> str:
    colour = D.CLASS_COLOR.get(slug or '', '#fff')
    med = medallion(slug, size) if slug else ''
    return (f'<span class="pc">{med}<span><span class="nm" style="color:{colour}">{escape(name)}</span>'
            f'{f"<span class=sb>{sub}</span>" if sub else ""}</span></span>')


def icon_cell(icon_html: str, name: str, sub: str = '') -> str:
    return (f'<span class="pc">{icon_html}<span><span class="nm" style="color:#fff">{escape(name)}</span>'
            f'{f"<span class=sb>{sub}</span>" if sub else ""}</span></span>')


def meter(frac: float, colour: str = 'var(--blue)') -> str:
    return f'<span class="mt"><i style="width:{max(0.0, min(frac, 1.0)) * 100:.1f}%;background:{colour}"></i></span>'


def table(heads: list[tuple[str, bool]], rows: list[list[str]], sort: str = '', row_cls: list[str] | None = None,
          hide_phone: set[int] | None = None, bar_col: int | None = None) -> str:
    """heads: (label, numeric). rows: cells as HTML. The first and last columns take the panel's edge padding;
    hide_phone: column indexes dropped under 700 px; bar_col: the meter column, one fixed width on every table."""
    hide_phone = hide_phone or set()

    def cls(i: int, numeric: bool, n: int) -> str:
        c = ['n' if numeric else '', 'first' if i == 0 else '', 'last' if i == n - 1 else '', 'hp' if i in hide_phone else '',
             'mc' if i == bar_col else '']
        return ' '.join(x for x in c if x)

    n = len(heads)
    th = ''.join(f'<th class="{cls(i, num, n)}{" on" if label == sort and label else ""}">{label}</th>' for i, (label, num) in enumerate(heads))
    trs = []
    for k, r in enumerate(rows):
        rc = (row_cls or [''] * len(rows))[k]
        trs.append(f'<tr class="{rc}">' + ''.join(f'<td class="{cls(i, heads[i][1], n)}">{c}</td>' for i, c in enumerate(r)) + '</tr>')
    return f'<table class="tb"><thead><tr>{th}</tr></thead><tbody>{"".join(trs)}</tbody></table>'


def facts(items: list[tuple[str, str]]) -> str:
    return '<dl class="fx">' + ''.join(f'<dt>{escape(k)}</dt><dd class="num">{v}</dd>' for k, v in items) + '</dl>'


def row_list(rows: list[tuple[str, str, str, str]]) -> str:
    """(lead, main, sub, trail), all HTML."""
    return '<ul class="rl">' + ''.join(f'<li>{lead}<span class="mn"><b>{main}</b>{f"<span>{sub}</span>" if sub else ""}</span>'
                                       f'<span class="tr">{trail}</span></li>' for lead, main, sub, trail in rows) + '</ul>'


def mmss(seconds: float) -> str:
    s = int(round(seconds))
    return f'{s // 60}:{s % 60:02d}'


def _smooth(values: list[float], k: int) -> list[float]:
    if k <= 1:
        return values
    out, acc = [], 0.0
    for i, v in enumerate(values):
        acc += v
        if i >= k:
            acc -= values[i - k]
        out.append(acc / min(i + 1, k))
    return out


def chips(items: list[tuple[str, bool]], swatches: dict[str, str] | None = None) -> str:
    """A filter row. swatches: a chip's mark colour, so the filter is its own legend (no separate key to repeat it)."""
    sw = swatches or {}
    return '<nav class="chips">' + ''.join(f'<a class="{"on" if on else ""}" href="#">'
                                           + (f"<i style='background:{sw[t]}'></i>" if t in sw else '') + f'{escape(t)}</a>' for t, on in items) + '</nav>'


def repeats(colours: list[str]) -> list[bool]:
    """Two players of one class share a colour: every repeat of a colour is drawn dashed, in the chart and its legend."""
    seen: set[str] = set()
    out = []
    for c in colours:
        out.append(c in seen)
        seen.add(c)
    return out


SHADE_STEPS = (1.0, 1.7, .5, 1.85, .35)  # the class colour, then lighter and darker turns for each further player of it


def shades(colours: list[str]) -> list[str]:
    """Same-class players stay in their class hue but each repeat steps lighter or darker, so lines and chips separate
    (with the dash) without leaving the class colour players read by."""
    count: dict[str, int] = {}
    out = []
    for c in colours:
        k = count.get(c, 0)
        count[c] = k + 1
        out.append(_scale(c, SHADE_STEPS[k % len(SHADE_STEPS)]))
    return out


def _scale(hex_colour: str, k: float) -> str:
    if not hex_colour.startswith('#') or len(hex_colour) != 7 or k == 1.0:
        return hex_colour
    rgb = [int(hex_colour[i:i + 2], 16) for i in (1, 3, 5)]
    rgb = [round(v + (255 - v) * (k - 1)) if k > 1 else round(v * k) for v in rgb]
    return '#' + ''.join(f'{max(0, min(255, v)):02x}' for v in rgb)


# ------------------------------------------------------------------------------------------------ charts
# One axis for every chart (reviewers, batch 1: "one tick scheme for all charts"): round minute ticks from TICK_STEPS,
# the y unit above the value labels, gridlines in GRID, crisp 1 px rules. Line, area and column charts all draw on it.
TICK_STEPS = (15, 30, 60, 120, 180, 300, 600, 900)
MAX_TICKS, MAX_TICKS_PHONE = 8, 4
PAD = {'l': 46, 'r': 14, 't': 24, 'b': 24}
NICE = (1, 2, 4, 8, 10)  # each top divides into four round quarters (0.25, 0.5, 1, 2, 2.5 of the magnitude)


def tick_step(duration_s: float, phone: bool) -> int:
    cap = MAX_TICKS_PHONE if phone else MAX_TICKS
    return next((s for s in TICK_STEPS if duration_s / s <= cap), TICK_STEPS[-1])


def nice_max(top: float) -> float:
    """The y axis's top: the next NICE step above the data, so the four gridlines read as round numbers."""
    top = top or 1
    mag = 10 ** math.floor(math.log10(top))
    return next(m * mag for m in NICE if top <= m * mag)


def _fmt(v: float) -> str:
    return f'{v:,.0f}' if v >= 10 or v == int(v) else f'{v:.2f}'.rstrip('0')


class Axis:
    """The drawing box of one chart: x is seconds into the fight (or a column index), y the series' own unit."""

    def __init__(self, width: int, height: int, duration_s: float, ymax: float, unit: str):
        self.w, self.h, self.dur, self.ymax, self.unit = width, height, duration_s, ymax, unit
        self.pw, self.ph = width - PAD['l'] - PAD['r'], height - PAD['t'] - PAD['b']
        self.base = PAD['t'] + self.ph
        self.phone = width < CHART_W / 2

    def x(self, t: float) -> float:
        return PAD['l'] + t / self.dur * self.pw

    def y(self, v: float) -> float:
        return self.base - min(v, self.ymax) / self.ymax * self.ph

    def frame(self, time_axis: bool = True) -> str:
        grid = ''.join(f'<line x1="{PAD["l"]}" x2="{self.w - PAD["r"]}" y1="{self.y(self.ymax * i / 4):.1f}" y2="{self.y(self.ymax * i / 4):.1f}" '
                       f'stroke="{GRID}" shape-rendering="crispEdges"/><text x="{PAD["l"] - 8}" y="{self.y(self.ymax * i / 4) + 4:.1f}" '
                       f'text-anchor="end" fill="{AXIS}" font-size="12">{_fmt(self.ymax * i / 4)}</text>' for i in range(1, 5))
        unit = (f'<text x="{PAD["l"] - 8}" y="{self.base + 4:.1f}" text-anchor="end" fill="{AXIS}" font-size="12">0</text>'
                f'<text x="{PAD["l"]}" y="11" fill="{AXIS}" font-size="12" font-weight="600">{escape(self.unit)}</text>')
        ticks = ''
        if time_axis:
            step = tick_step(self.dur, self.phone)
            ticks = ''.join(f'<text x="{self.x(t):.1f}" y="{self.h - 6}" text-anchor="{"start" if t == 0 else "middle"}" fill="{AXIS}" '
                            f'font-size="12">{mmss(t)}</text>' for t in range(0, int(self.dur) + 1, step))
        base = (f'<line x1="{PAD["l"]}" x2="{self.w - PAD["r"]}" y1="{self.base}" y2="{self.base}" stroke="{AXIS}" stroke-opacity=".5" '
                f'shape-rendering="crispEdges"/>')
        return grid + unit + ticks + base

    def svg(self, inner: str, label: str) -> str:
        return (f'<svg viewBox="0 0 {self.w} {self.h}" role="img" aria-label="{escape(label)}" font-family="Open Sans,Arial,sans-serif">'
                f'{inner}</svg>')


def _points(ax: Axis, vals: list[float]) -> list[tuple[float, float]]:
    # a series spans the whole fight whatever its sample count, so its index maps onto the fight's length
    return [(ax.x(i * ax.dur / max(len(vals) - 1, 1)), ax.y(v)) for i, v in enumerate(vals)]


def _death_marks(ax: Axis, marks: list[float] | None) -> str:
    return ''.join(f'<line x1="{ax.x(t):.1f}" x2="{ax.x(t):.1f}" y1="{ax.base - 8}" y2="{ax.base + 2}" stroke="{DEATH}" stroke-width="2.5"/>'
                   for t in marks or [])


def line_chart(series: list[tuple[str, str, list[float]]], duration_s: float, height: int = 240, marks: list[float] | None = None,
               smooth: int = 10, label: str = '', width: int = CHART_W, unit: str = '', area: bool = False,
               ymax: float | None = None) -> str:
    """series: (name, colour, per-second values), colours already through shades(). Deaths are short red ticks on the axis.
    The values are smoothed over `smooth` seconds, as the live chart does. `area` fills under each line at a low alpha
    (one-series charts: resources, a single player's output). `width` is the drawing's own width: the phone drawing is
    narrower so its type stays readable when it scales."""
    sm = [(n, c, _smooth(v, smooth)) for n, c, v in series]
    ax = Axis(width, height, duration_s, ymax or nice_max(max((max(v) for _, _, v in sm if v), default=1)), unit)
    dashed = repeats([c for _, c, _ in sm])
    out = ''
    for (_, c, vals), d in zip(sm, dashed):
        if not vals:
            continue
        pts = _points(ax, vals)
        path = ' '.join(f'{px:.1f},{py:.1f}' for px, py in pts)
        if area:
            out += f'<polygon fill="{c}" fill-opacity=".16" points="{pts[0][0]:.1f},{ax.base} {path} {pts[-1][0]:.1f},{ax.base}"/>'
        dash = ' stroke-dasharray="7 4"' if d else ''
        out += f'<polyline fill="none" stroke="{c}" stroke-width="{1.5 if area else 2}" stroke-linejoin="round"{dash} points="{path}"/>'
    return ax.svg(ax.frame() + out + _death_marks(ax, marks), label)


def column_chart(bars: list[tuple[str, float, str]], unit: str, height: int = 200, label: str = '', width: int = CHART_W,
                 every: int = 1) -> str:
    """bars: (x label, value, colour). Square-topped columns with a 2 px gap; a label under every `every`th column."""
    n = len(bars) or 1
    ax = Axis(width, height, 1, nice_max(max((v for _, v, _ in bars), default=1)), unit)
    slot = ax.pw / n
    cols = ''.join(f'<rect x="{PAD["l"] + i * slot + 1:.1f}" y="{ax.y(v):.1f}" width="{max(slot - 2, 1):.1f}" height="{ax.base - ax.y(v):.1f}" '
                   f'fill="{c}" shape-rendering="crispEdges"/>' for i, (_, v, c) in enumerate(bars))
    labels = ''.join(f'<text x="{PAD["l"] + (i + .5) * slot:.1f}" y="{height - 6}" text-anchor="middle" fill="{AXIS}" font-size="12">'
                     f'{escape(t)}</text>' for i, (t, _, _) in enumerate(bars) if i % every == 0)
    return ax.svg(ax.frame(time_axis=False) + cols + labels, label)


PHONE_CHART_W, PHONE_CHART_H = 380, 230  # renders about 215 px tall on a 390 phone, its 12 px type at about 11 px


def responsive_chart(series: list[tuple[str, str, list[float]]], duration_s: float, marks: list[float] | None = None,
                     label: str = '', unit: str = '', area: bool = False, height: int = 240, ymax: float | None = None) -> str:
    """The desktop drawing and the phone drawing of one chart; CSS shows the one that fits."""
    return (f'<div class="od">{line_chart(series, duration_s, height, marks, label=label, unit=unit, area=area, ymax=ymax)}</div>'
            f'<div class="om">{line_chart(series, duration_s, PHONE_CHART_H, marks, label=label, width=PHONE_CHART_W, unit=unit, area=area, ymax=ymax)}</div>')


def chart_frame(chart_svg: str, legend: list[tuple[str, str, str]]) -> str:
    """The one chart frame: the chart, then a legend whose chips carry each series' figure with its unit."""
    dashed = repeats([c for _, c, _ in legend])

    def chip(c: str, d: bool) -> str:
        return f'repeating-linear-gradient(90deg,{c} 0 4px,transparent 4px 6px)' if d else c

    lg = ''.join(f'<span><i style="background:{chip(c, d)}"></i>{escape(n)} <b>{fig}</b></span>' for (n, c, fig), d in zip(legend, dashed))
    return f'<div class="cf">{chart_svg}<div class="lg">{lg}</div></div>'


def talent_trees(trees: list[tuple[str, int, list[tuple[int, int, str, int, int]]]],
                 hover: tuple[int, int, int, str] | None = None) -> str:
    """trees: (name, points, cells (tier, column, icon src, rank, max rank)). The client's grid; the planner edits it, the
    sim and guides show it read only. hover: (tree index, tier, column, the client tooltip HTML) drawn beside that cell."""
    out = []
    for i, (name, points, cells) in enumerate(trees):
        tcs = []
        for tier, col, src, rank, mx in cells:
            state = 'max' if rank and rank == mx else ('on' if rank else 'off')
            badge = f'<b>{rank}/{mx}</b>' if rank else ''
            hv = hover is not None and hover[:3] == (i, tier, col)
            tip = f'<span class="ttip">{hover[3]}</span>' if hv else ''
            tcs.append(f'<span class="tc {state}{" hv" if hv else ""}" style="grid-row:{tier + 1};grid-column:{col + 1}"><img src="{src}" alt="">{badge}{tip}</span>')
        out.append(f'<div><h3>{escape(name)}<span class="num">{points}</span></h3><div class="tg">{"".join(tcs)}</div></div>')
    return f'<div class="tt">{"".join(out)}</div>'


def point_order(rows: list[tuple[str, list[tuple[str, int, str] | None]]]) -> str:
    """rows: (label, sockets); a socket is (icon src, rank, tooltip text) or None for a level with no point yet."""
    def sock(s: tuple[str, int, str] | None) -> str:
        if s is None:
            return '<span class="s e"></span>'
        src, rank, tip = s
        return f'<span class="s" title="{escape(tip)}"><img src="{src}" alt=""><b>{rank}</b></span>'

    return '<ul class="po">' + ''.join(f'<li><span class="lv">{escape(lb)}</span><span class="sk">{"".join(sock(s) for s in ss)}</span></li>'
                                       for lb, ss in rows) + '</ul>'


def reading_section(title: str, prose_html: str, sides: list[tuple[str, str]], anchor: str = '') -> str:
    """sides: (heading, body HTML) blocks beside the prose."""
    side = ''.join(f'<div class="side">{f"<h3>{escape(h)}</h3>" if h else ""}{b}</div>' for h, b in sides)
    wide = '' if sides else ' style="grid-column:1/-1"'
    return (f'<section class="rs"{f" id={anchor}" if anchor else ""}><h2>{escape(title)}</h2><div class="prose"{wide}>{prose_html}</div>'
            f'{f"<div>{side}</div>" if sides else ""}</section>')


def toc(items: list[tuple[str, str]], current: str) -> str:
    return '<nav class="toc">' + ''.join(f'<a class="{"on" if a == current else ""}" href="#{a}">{escape(t)}</a>' for t, a in items) + '</nav>'


def item_row(icon_src: str, name: str, quality: str, sub: str = '', trail: str = '') -> str:
    """The one item row: the icon's edge and the name both in the item's quality colour (the page's lifted colours)."""
    q = f'var(--q-{quality})'
    img = f'<img src="{icon_src}" alt="" style="--qb:{q}">' if icon_src else '<span class="ai" style="width:36px;height:36px;background:#141414"></span>'
    return (f'<div class="ir">{img}<span><span class="n" style="color:{q}">{escape(name)}</span>{f"<span class=s>{sub}</span>" if sub else ""}</span>'
            f'<span class="t">{trail}</span></div>')


def tag(text: str, tone: str = '') -> str:
    return f'<span class="tag {tone}">{escape(text)}</span>'


def answers(choices: list[str], current: str) -> str:
    return '<span class="ans" role="radiogroup">' + ''.join(f'<a class="{"on" if c.lower() == current.lower() else ""}" href="#" role="radio">{escape(c)}</a>'
                                                            for c in choices) + '</span>'


def need_grid(heads: list[str], rows: list[tuple[str, str, list[tuple[str, str]]]]) -> str:
    """rows: (label HTML, sub, cells (tone, text)); tone one of ok, warn, miss, none. Drawn as the shared table."""
    body = [[f'<span class="pc"><span><span class="nm" style="color:#fff">{lb}</span>{f"<span class=sb>{sub}</span>" if sub else ""}</span></span>']
            + [f'<span class="{tone}">{text}</span>' for tone, text in cells] for lb, sub, cells in rows]
    return '<div class="ng">' + table([('', False)] + [(h, True) for h in heads], body).replace('<td class="n', '<td class="c n') + '</div>'


def votes(yes: int, no: int, voters: int) -> str:
    return f'{tag(f"{yes} yes", "kill")} {tag(f"{no} no", "wipe" if no else "")} <span class="muted" style="font-size:12.5px">of {voters}</span>'


def ability_icon(src: str, size: int = 24) -> str:
    return f'<img class="ai" src="{src}" alt="" width="{size}" height="{size}">'


def settings_bar(fields: list[tuple[str, str]], action: str = '') -> str:
    """(label, value) pickers in a row, the one action at the end (the sim's run bar, the tools' settings)."""
    fs = ''.join(f'<span class="f">{escape(k)}<b>{v}</b></span>' for k, v in fields)
    return f'<div class="card setb">{fs}{f"<span class=go>{action}</span>" if action else ""}</div>'


def field(placeholder: str, button: str, label: str = '') -> str:
    return f'<div class="fld"><input type="text" placeholder="{escape(placeholder)}" aria-label="{escape(label or placeholder)}">{button}</div>'


def histogram(counts: list[int], start: float, bin_w: float, mean: float, lo: float, hi: float, label: str = '') -> str:
    """Runs per bin as bars; the middle two thirds (lo to hi) shaded; the mean as a gold rule; ends labelled."""
    w, h = 400, 120
    n = len(counts)
    top = max(counts) or 1
    bw = w / n

    def xv(v: float) -> float:
        return (v - start) / (n * bin_w) * w

    shade = f'<rect x="{xv(lo):.1f}" y="0" width="{xv(hi) - xv(lo):.1f}" height="{h}" fill="rgba(0,147,178,.12)"/>'
    bars = ''.join(f'<rect x="{i * bw + .5:.1f}" y="{h - c / top * (h - 4):.1f}" width="{bw - 1:.1f}" height="{c / top * (h - 4):.1f}" fill="var(--blue)"/>'
                   for i, c in enumerate(counts) if c)
    rule = f'<line x1="{xv(mean):.1f}" x2="{xv(mean):.1f}" y1="0" y2="{h}" stroke="var(--gold)" stroke-width="2"/>'
    return (f'<div class="hg"><svg viewBox="0 0 {w} {h}" role="img" aria-label="{escape(label)}">{shade}{bars}{rule}</svg>'
            f'<div class="ax"><span>{start:,.0f}</span><span>{start + n * bin_w:,.0f}</span></div></div>')


def lanes(rows: list[tuple[str, str]], duration_s: float, ticks: int = 60) -> str:
    """rows: (label HTML, SVG body drawn on a 0..1000 x 0..30 lane). One shared minute axis under them."""
    labels = ''.join(f'<div class="lb">{lb}</div>' for lb, _ in rows)
    svgs = ''.join(f'<svg viewBox="0 0 1000 30" preserveAspectRatio="none"><rect x="0" y="13" width="1000" height="4" fill="rgba(255,255,255,.06)"/>{body}</svg>'
                   for _, body in rows)
    # each tick sits at its own time on the lane's scale, not spread evenly
    axis = ''.join(f'<span style="left:{t / duration_s * 100:.2f}%">{mmss(t)}</span>' for t in range(0, int(duration_s) + 1, ticks))
    return f'<div class="ln"><div>{labels}</div><div>{svgs}</div><div class="ax">{axis}</div></div>'


# ------------------------------------------------------------------------------------------------ small marks
PARSE = {'gold': '#e5cc80', 'pink': '#e268a8', 'orange': '#ff8000', 'purple': '#a335ee', 'blue': '#0070ff', 'green': '#1eff00', 'grey': '#8c8c8c'}
PARSE_STEPS = ((100, 'gold'), (99, 'pink'), (95, 'orange'), (75, 'purple'), (50, 'blue'), (25, 'green'), (0, 'grey'))
PARSE_LEGEND = (('100', 'gold'), ('99+', 'pink'), ('95+', 'orange'), ('75+', 'purple'), ('50+', 'blue'), ('25+', 'green'), ('under 25', 'grey'))


def parse_colour(pct: float) -> str:
    """The rank colour ladder (lib/report/format.ts percentileToken): gold 100, pink 99+, orange 95+, ... grey under 25."""
    return PARSE[next(name for floor, name in PARSE_STEPS if pct >= floor)]


def parse_legend(lead: str = '') -> str:
    lead_html = f'<span style="color:var(--muted);font-weight:400">{escape(lead)}</span>' if lead else ''
    return '<p class="plg">' + lead_html + ''.join(f'<b style="color:{PARSE[c]}">{t}</b>' for t, c in PARSE_LEGEND) + '</p>'


def strip_into_view() -> str:
    """On a phone the band's control strip scrolls; the selected tab starts in view (a tab past the fold would be hidden)."""
    return ('<script>if(innerWidth<700)for(const n of document.querySelectorAll(".rb .ctl")){const a=n.querySelector("a.on");'
            'if(a)n.scrollLeft=Math.max(0,a.offsetLeft-n.offsetLeft)}</script>')


def code_block(text: str) -> str:
    return f'<pre class="code">{escape(text)}</pre>'
