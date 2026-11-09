"""The shared drawn primitives (owner, 2026-11-08: "we should use shared components and resources as a first class
citizen"). Every directions surface draws its panels, tables, player cells, meters, facts, row lists and charts from
here, once; a surface that needs something new adds it here first. The inventory and the web component each one
becomes are in docs/superpowers/plans/2026-11-09-redesign.md, "Shared component inventory".
"""
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
.pn{margin-top:22px}.pn>.hd{display:flex;align-items:baseline;gap:12px;padding:16px 20px 10px}
.pn>.hd h2{font-size:21px;line-height:26px}.pn>.hd .m{font-size:13px;color:var(--muted)}.pn>.hd .a{margin-left:auto;font-size:14px;white-space:nowrap}
.pn>.bd{padding:0 20px 16px}.pn>.bd.flush{padding:0 0 6px}
/* data table: tabular figures right-aligned; the first column a player or ability cell */
.tb{width:100%;border-collapse:collapse;font-size:14px}
.tb th{font:600 12px var(--body);color:var(--muted);text-align:left;padding:8px 10px;border-bottom:1px solid rgba(255,255,255,.14);white-space:nowrap}
.tb th.n,.tb td.n{text-align:right;font-variant-numeric:tabular-nums lining-nums;white-space:nowrap}
.tb th.on{color:#fff}
.tb td{padding:7px 10px;border-bottom:1px solid rgba(255,255,255,.07);color:var(--text);vertical-align:middle}
.tb tr:last-child td{border-bottom:0}.tb td.strong{color:#fff;font-weight:600}
.tb tr.tot td{color:#fff;font-weight:600;border-top:1px solid rgba(255,255,255,.14)}
.tb .first{padding-left:20px}.tb .last{padding-right:20px}
.tb tr.me td{background:rgba(14,134,202,.14)}
/* a player (or ability) cell: medallion or icon, the name in class colour, one sub line */
.pc{display:flex;align-items:center;gap:10px;min-width:0}.pc .nm{display:block;font-weight:600;line-height:19px;white-space:nowrap}
.pc .sb{display:block;font-size:12.5px;color:var(--muted);line-height:17px;white-space:nowrap}
/* meter: the share bar inside a table row, in the row's class colour */
.mt{display:block;height:6px;background:rgba(255,255,255,.07);min-width:90px}.mt i{display:block;height:100%}
/* facts: label and value pairs, one per line */
.fx{display:grid;grid-template-columns:auto 1fr;gap:6px 16px;font-size:14px}.fx dt{color:var(--muted)}.fx dd{margin:0;color:#fff;text-align:right}
/* row list: lead, main with a sub line, trail */
.rl li{display:grid;grid-template-columns:auto 1fr auto;gap:10px;align-items:center;padding:8px 0;border-bottom:1px solid rgba(255,255,255,.07)}
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
@media(max-width:700px){.tc .ttip{display:none}.tb th{white-space:normal}}
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
.ir{display:grid;grid-template-columns:36px 1fr auto;gap:10px;align-items:center;padding:7px 0;border-bottom:1px solid rgba(255,255,255,.07)}
.ir:last-child{border-bottom:0}.ir img{width:36px;height:36px;border-radius:var(--r-ic);border:1px solid #000;box-shadow:0 0 0 1px var(--qb,var(--rule))}
.ir .n{display:block;font-weight:600;line-height:19px}.ir .s{display:block;font-size:12.5px;color:var(--muted);line-height:17px}.ir .t{font-size:13px;color:var(--muted);white-space:nowrap}
.twocol{display:grid;grid-template-columns:minmax(0,1fr) 340px;gap:28px;align-items:start}
@media(max-width:1100px){.twocol{grid-template-columns:minmax(0,1fr)}}
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
          hide_phone: set[int] | None = None) -> str:
    """heads: (label, numeric). rows: cells as HTML. The first and last columns take the panel's edge padding;
    hide_phone: column indexes dropped under 700 px."""
    hide_phone = hide_phone or set()

    def cls(i: int, numeric: bool, n: int) -> str:
        c = ['n' if numeric else '', 'first' if i == 0 else '', 'last' if i == n - 1 else '', 'hp' if i in hide_phone else '']
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


def chips(items: list[tuple[str, bool]]) -> str:
    return '<nav class="chips">' + ''.join(f'<a class="{"on" if on else ""}" href="#">{escape(t)}</a>' for t, on in items) + '</nav>'


def repeats(colours: list[str]) -> list[bool]:
    """Two players of one class share a colour: every repeat of a colour is drawn dashed, in the chart and its legend."""
    seen: set[str] = set()
    out = []
    for c in colours:
        out.append(c in seen)
        seen.add(c)
    return out


SHADE_STEPS = (1.0, 1.45, .62, 1.8, .4)  # the class colour, then lighter and darker turns for each further player of it


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


def line_chart(series: list[tuple[str, str, list[float]]], duration_s: float, height: int = 220,
               marks: list[float] | None = None, smooth: int = 10, label: str = '', width: int = CHART_W) -> str:
    """series: (name, colour, per-second values). Gridlines and axis in the page's rule colours, minute ticks in mm:ss,
    deaths as short red ticks on the axis. The values are smoothed over `smooth` seconds, as the live chart does.
    `width` is the drawing's own width: the phone drawing is narrower so its type stays readable when it scales."""
    pad_l, pad_r, pad_t, pad_b = 46, 12, 10, 24
    w, h = width, height
    pw, ph = w - pad_l - pad_r, h - pad_t - pad_b
    sm = [(n, c, _smooth(v, smooth)) for n, c, v in series]
    top = max((max(v) for _, _, v in sm if v), default=1) or 1
    step = 10 ** max(0, len(str(int(top))) - 1)
    ymax = (int(top / step) + 1) * step
    y = lambda v: pad_t + ph - v / ymax * ph
    x = lambda t: pad_l + t / duration_s * pw
    grid = ''.join(f'<line x1="{pad_l}" x2="{w - pad_r}" y1="{y(ymax * i / 4):.1f}" y2="{y(ymax * i / 4):.1f}" stroke="{GRID}"/>'
                   f'<text x="{pad_l - 8}" y="{y(ymax * i / 4) + 4:.1f}" text-anchor="end" fill="{AXIS}" font-size="12">{int(ymax * i / 4):,}</text>'
                   for i in range(5))
    tick_every = (60 if duration_s > 240 else 30) * (3 if width < CHART_W / 2 else 1)
    ticks = ''.join(f'<text x="{x(t):.1f}" y="{h - 6}" text-anchor="middle" fill="{AXIS}" font-size="12">{mmss(t)}</text>'
                    for t in range(0, int(duration_s) + 1, tick_every))
    # a series spans the whole fight whatever its sample count, so its index maps onto the fight's length
    dashed = repeats([c for _, c, _ in sm])
    dash = ' stroke-dasharray="7 4"'
    lines = ''.join(f'<polyline fill="none" stroke="{c}" stroke-width="2" stroke-linejoin="round"{dash if d else ""} points="'
                    + ' '.join(f'{x(i * duration_s / max(len(vals) - 1, 1)):.1f},{y(v):.1f}' for i, v in enumerate(vals)) + '"/>'
                    for (_, c, vals), d in zip(sm, dashed) if vals)
    dm = ''.join(f'<line x1="{x(t):.1f}" x2="{x(t):.1f}" y1="{pad_t + ph - 8}" y2="{pad_t + ph + 2}" stroke="{DEATH}" stroke-width="2.5"/>'
                 for t in marks or [])
    return (f'<svg viewBox="0 0 {w} {h}" role="img" aria-label="{escape(label)}" font-family="Open Sans,Arial,sans-serif">{grid}{ticks}'
            f'<line x1="{pad_l}" x2="{w - pad_r}" y1="{pad_t + ph}" y2="{pad_t + ph}" stroke="{AXIS}" stroke-opacity=".5"/>{lines}{dm}</svg>')


PHONE_CHART_W, PHONE_CHART_H = 380, 200


def responsive_chart(series: list[tuple[str, str, list[float]]], duration_s: float, marks: list[float] | None = None,
                     label: str = '') -> str:
    """The desktop drawing and the phone drawing of one chart; CSS shows the one that fits."""
    return (f'<div class="od">{line_chart(series, duration_s, marks=marks, label=label)}</div>'
            f'<div class="om">{line_chart(series, duration_s, PHONE_CHART_H, marks, label=label, width=PHONE_CHART_W)}</div>')


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
