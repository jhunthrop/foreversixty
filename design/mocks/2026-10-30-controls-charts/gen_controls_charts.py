"""Boards for design/specs/2026-10-30-controls-charts.md. Run from the repo root:
python3 design/mocks/2026-10-30-controls-charts/gen_controls_charts.py && node design/mocks/2026-10-30-controls-charts/render.mjs
Draws on the 2026-10-28 polish parts (direction B, Forge): tokens, crests, glyphs. The fight is
mock data shaped like a Ragnaros kill (4:12, submerge at 3:00); every label is the spec's copy."""
import math
import random
import sys
from pathlib import Path

HERE = Path(__file__).resolve().parent
sys.path.insert(0, str(HERE.parent / "2026-10-28-polish"))
from parts import CLASS_COLOR, GLYPHS, crest, glyph, uri  # noqa: E402
from theme import head  # noqa: E402

RENDERS = HERE / "renders"
GLYPHS["skull"] = "<path d=\"M12 3a7 7 0 0 0-7 7c0 2.4 1.2 4 2.5 5v3h9v-3c1.3-1 2.5-2.6 2.5-5a7 7 0 0 0-7-7z\"/><circle cx=\"9.5\" cy=\"10.5\" r=\"1.3\"/><circle cx=\"14.5\" cy=\"10.5\" r=\"1.3\"/><path d=\"M10 18v3M14 18v3\"/>"
GLYPHS["flag"] = "<path d=\"M5 21V4M5 4h11l-2 4 2 4H5\"/>"


def icon(stem: str) -> str:
    return uri(HERE / "assets" / f"{stem}.webp")


TAUNT, DEATH_WISH, INNERVATE = icon("spell_nature_reincarnation"), icon("spell_shadow_deathpact"), icon("spell_nature_lightning")

CSS = """
.board{display:flex;flex-direction:column;min-height:100vh;background:#07090d}
.dirbar{display:flex;align-items:center;gap:12px;height:34px;padding:0 20px;background:#000;border-bottom:1px solid #1c2230;font-size:12px;color:#9a9484;font-family:"JetBrains Mono",monospace}
.dirbar b{color:#e5b955}
.cap{font-size:12px;color:#9a9484;font-family:"JetBrains Mono",monospace;margin:0 0 10px}
.num{font-variant-numeric:tabular-nums}
.ql{font-size:13px;font-weight:600;color:var(--muted)}
.eyebrow{display:flex;align-items:center;gap:12px;font-size:11px;font-weight:700;letter-spacing:.18em;text-transform:uppercase;color:var(--gold)}
.eyebrow::before{content:"";width:24px;height:1px;background:rgba(229,185,85,.6)}
.stitle{font-family:var(--fh);font-weight:800;font-size:18px;line-height:1;text-transform:uppercase;letter-spacing:.01em;color:var(--strong)}
/* the closed field */
.sel{position:relative;display:inline-flex;align-items:center;gap:8px;height:52px;padding:0 40px 0 12px;border-radius:8px;border:1px solid rgba(229,185,85,.17);background:rgba(255,255,255,.045);color:var(--strong);font-size:15px;font-weight:600;white-space:nowrap;min-width:180px}
.sel.c{height:44px}.sel.d{height:36px;font-size:13px;padding-left:10px;padding-right:34px;min-width:150px}
.sel .chv{position:absolute;right:12px;top:50%;transform:translateY(-50%);color:var(--muted);display:flex}
.sel.d .chv{right:10px}
.sel.hover{border-color:rgba(229,185,85,.30);background:rgba(255,255,255,.06)}
.sel.focus{outline:2px solid var(--gold);outline-offset:2px}
.sel.open{border-color:rgba(229,185,85,.55);background:rgba(255,255,255,.06)}
.sel.open .chv{transform:translateY(-50%) rotate(180deg);color:var(--gold)}
.sel.dis{opacity:.45}
.sel.err{border-color:var(--wipe)}
.sel.ph{color:var(--muted);font-weight:500}
.sel .crest{box-shadow:0 0 0 1.5px color-mix(in srgb,var(--c) 60%,transparent)}
.spin{width:14px;height:14px;border-radius:999px;border:2px solid rgba(229,185,85,.25);border-top-color:var(--gold)}
.errline{font-size:13px;color:var(--wipe);margin-top:6px}
/* the listbox */
.lb{position:absolute;z-index:5;width:max-content;min-width:100%;max-width:420px;border:1px solid rgba(229,185,85,.17);border-radius:12px;background:linear-gradient(180deg, rgba(229,185,85,.075), rgba(229,185,85,.012) 45%), #0d0f15;box-shadow:inset 0 1px 0 rgba(245,210,122,.45),0 16px 38px rgba(0,0,0,.55);padding:4px 0;overflow:hidden}
.lb .scroll{max-height:360px;overflow:hidden;position:relative}
.lb .fade{position:absolute;left:0;right:0;bottom:0;height:24px;background:linear-gradient(180deg,rgba(13,15,21,0),#0d0f15)}
.opt{position:relative;display:flex;align-items:center;gap:10px;min-height:36px;padding:6px 40px 6px 12px;font-size:14px;font-weight:600;color:var(--strong)}
.opt.ic{min-height:40px}
.opt .t{display:flex;flex-direction:column;gap:2px;min-width:0}
.opt .ds{font-size:12px;font-weight:500;color:var(--muted);line-height:1.35;white-space:normal;max-width:330px}
.opt .meta{margin-left:auto;padding-left:16px;font-size:13px;font-weight:500;color:var(--muted)}
.opt.act{background:rgba(229,185,85,.10);box-shadow:inset 2px 0 0 var(--gold)}
.opt .ck{position:absolute;right:12px;top:50%;transform:translateY(-50%);color:var(--gold);display:flex}
.opt.dis{opacity:.45}
.opt mark{background:none;color:var(--gold);font-weight:700}
.grp{padding:10px 12px 4px;font-size:11px;font-weight:700;letter-spacing:.14em;text-transform:uppercase;color:var(--muted)}
.grp.rule{border-top:1px solid var(--line-soft);margin-top:4px}
.lbsearch{display:flex;align-items:center;gap:8px;padding:6px 8px 8px;border-bottom:1px solid var(--line-soft);margin-bottom:4px}
.lbsearch .in{flex:1;display:flex;align-items:center;gap:8px;height:36px;padding:0 10px;border-radius:6px;border:1px solid var(--gold);background:#07090d;color:var(--strong);font-size:14px;outline:2px solid rgba(229,185,85,.35);outline-offset:0}
.lbsearch .in .caret{width:1px;height:16px;background:var(--gold);margin-left:-6px}
.lbsearch .n{font-size:12px;color:var(--muted);white-space:nowrap;padding-right:4px}
.lbempty{padding:12px;font-size:14px;color:var(--muted)}
/* the chart frame */
.frame{background:linear-gradient(180deg, rgba(229,185,85,.075), rgba(229,185,85,.012) 45%), #0d0f15;border:1px solid rgba(229,185,85,.17);border-radius:12px;box-shadow:0 16px 38px rgba(0,0,0,.55);position:relative}
.fhead{display:flex;align-items:center;gap:12px;min-height:48px;padding:10px 16px}
.fhead .meta{font-size:13px;color:var(--muted)}
.fhead .right{margin-left:auto;display:flex;align-items:center;gap:10px}
.legend{display:flex;flex-wrap:wrap;min-height:42px;gap:4px 6px;padding:8px 12px 6px;border-top:1px solid var(--line-soft)}
.chip{display:inline-flex;align-items:center;gap:7px;height:28px;padding:0 8px;border-radius:6px;font-size:13px;font-weight:600;white-space:nowrap}
.chip .sw{width:12px;height:3px;border-radius:2px}
.chip .v{font-weight:500;color:var(--muted)}
.chip.off{opacity:.4}.chip.off .sw{background:none!important;box-shadow:inset 0 0 0 1px currentColor;height:4px}
.chip.hov{background:rgba(229,185,85,.06)}
.chip .only{font-size:12px;font-weight:600;color:var(--gold);margin-left:2px}
.chip.more{color:var(--gold)}
.ghost{display:inline-flex;align-items:center;gap:6px;height:36px;padding:0 10px;border-radius:8px;color:var(--gold);font-size:13px;font-weight:700;letter-spacing:.07em;text-transform:uppercase}
.sw2{display:inline-flex;align-items:center;gap:8px;font-size:13px;color:var(--text)}
.sw2 .tr{position:relative;width:36px;height:20px;border-radius:999px;background:#232a3a;border:1px solid #303a50}
.sw2 .tr::after{content:"";position:absolute;top:2px;left:2px;width:14px;height:14px;border-radius:999px;background:#8a8576}
.tip{position:absolute;z-index:6;width:268px;border:1px solid rgba(229,185,85,.17);border-radius:10px;background:#0d0f15;box-shadow:inset 0 1px 0 rgba(245,210,122,.45),0 16px 38px rgba(0,0,0,.6);padding:8px 0 6px;pointer-events:none}
.tip .th{display:flex;align-items:baseline;gap:8px;padding:0 12px 6px;border-bottom:1px solid var(--line-soft);margin-bottom:4px}
.tip .th b{font-size:14px;color:var(--strong)}.tip .th span{font-size:12px;color:var(--muted);margin-left:auto}
.tip .ev{display:flex;align-items:center;gap:8px;padding:3px 12px;font-size:12px;color:var(--text)}
.tip .ev img{width:16px;height:16px;border-radius:3px}
.tip .r{display:flex;align-items:center;gap:8px;height:22px;padding:0 12px;font-size:13px;font-weight:600}
.tip .r .sw{width:3px;height:12px;border-radius:2px}
.tip .r .v{margin-left:auto;color:var(--strong);font-weight:600}
.tip .r .p{width:38px;text-align:right;color:var(--muted);font-weight:500;font-size:12px}
.tip .r.pin{background:rgba(229,185,85,.08)}
.tip .r.dead{color:var(--muted);font-weight:500}
.tip .more{padding:4px 12px 0;font-size:12px;color:var(--muted)}
.cnote{font-size:12px;color:var(--muted);line-height:1.5}
.tabs{display:flex;gap:4px;border-bottom:1px solid var(--line-soft)}
.tabs span{padding:10px 12px;font-size:14px;font-weight:600;color:var(--nav);border-bottom:2px solid transparent}
.tabs span.on{color:var(--strong);border-color:var(--gold)}
.tbl{width:100%;border-collapse:collapse;font-size:14px}
.tbl th{font-size:13px;font-weight:600;color:var(--muted);text-align:left;padding:8px 12px}
.tbl td{padding:8px 12px;border-top:1px solid var(--line-soft)}
.tbl .r{text-align:right}
.bar{height:6px;border-radius:3px;background:rgba(255,255,255,.06);position:relative;width:220px}
.bar i{position:absolute;inset:0 auto 0 0;border-radius:3px}
.skel{background:linear-gradient(90deg,rgba(255,255,255,.035),rgba(255,255,255,.07),rgba(255,255,255,.035));border-radius:6px}
.lbl{font-size:12px;color:var(--muted);font-family:"JetBrains Mono",monospace}
.overlay-note{position:absolute;font-size:11px;font-family:"JetBrains Mono",monospace;color:#7fd48a}
"""

# --------------------------------------------------------------------- the fight (mock data)
DURATION = 252
SUBMERGE = 180
random.seed(60)
PLAYERS = [  # name, class, threat per second, dps, role
    ("Brokk", "warrior", 520, 300, "tank"), ("Thandril", "warrior", 230, 280, "tank"),
    ("Garrosk", "warrior", 340, 780, "dps"), ("Morrigan", "mage", 330, 720, "dps"),
    ("Ashka", "warlock", 320, 690, "dps"), ("Velyn", "rogue", 250, 650, "dps"),
    ("Kael", "hunter", 260, 560, "dps"), ("Lightningbolltz", "shaman", 230, 480, "dps"),
    ("Seraphine", "priest", 140, 40, "healer"), ("Oakhorn", "druid", 120, 30, "healer"),
]
BROKK_DIES, TAUNT_AT = 126, 127


def threat_series():
    out = {}
    for name, _cls, tps, _dps, _role in PLAYERS:
        total, line = 0.0, []
        for t in range(DURATION + 1):
            if name == "Brokk" and t > BROKK_DIES:
                line.append(None)
                continue
            rate = tps
            if name == "Thandril" and t > TAUNT_AT:
                rate = 560
            if t > SUBMERGE:
                rate *= 0.15  # Ragnaros is submerged; the threat goes on the Sons of Flame
            total += rate * (0.75 + random.random() * 0.5)
            line.append(total)
        out[name] = line
    return out


def dps_series():
    out = {}
    for name, _cls, _tps, dps, _role in PLAYERS:
        raw = []
        for t in range(DURATION + 1):
            if name == "Brokk" and t > BROKK_DIES:
                raw.append(0)
                continue
            v = dps * (0.55 + random.random() * 0.9)
            if name == "Garrosk" and (6 <= t <= 36 or 186 <= t <= 216):
                v *= 1.35  # Death Wish
            if SUBMERGE <= t <= SUBMERGE + 9:
                v *= 0.45  # Sons of Flame spawning
            if t < 4:
                v *= t / 4
            raw.append(v)
        smooth = [sum(raw[max(0, t - 4):t + 1]) / len(raw[max(0, t - 4):t + 1]) for t in range(len(raw))]
        out[name] = [None if (name == "Brokk" and t > BROKK_DIES) else v for t, v in enumerate(smooth)]
    return out


THREAT, DPS = threat_series(), dps_series()
CLS = {p[0]: p[1] for p in PLAYERS}


def mmss(t: float) -> str:
    t = int(round(t))
    return f"{t // 60}:{t % 60:02d}"


def kfmt(v: float) -> str:
    if v >= 1000 and v % 1000 == 0:
        return f"{v / 1000:.0f}k"
    return f"{v / 1000:.1f}k" if v >= 1000 else f"{v:.0f}"


def nice_step(span: float, width: float) -> int:
    for step in (5, 10, 15, 30, 60, 120, 300):
        if width / (span / step) >= 64:
            return step
    return 600


# --------------------------------------------------------------------- the chart drawing
def chart_svg(w, plot_h, x0, x1, series, ymax, *, yticks, lane=None, deaths=(), taunts=(), phases=(), cursor=None,
              drag=None, dim=None, area=None, gutter=52, lane_h=26, phone=False):
    """series: list of (name, colour, values per second). Returns an SVG string."""
    right = 12
    pw = w - gutter - right
    top = lane_h
    h = top + plot_h + 22
    fx = lambda t: gutter + (t - x0) / (x1 - x0) * pw  # noqa: E731
    fy = lambda v: top + plot_h - (v / ymax) * plot_h  # noqa: E731
    s = [f"<svg width=\"{w}\" height=\"{h}\" viewBox=\"0 0 {w} {h}\" style=\"display:block;font-family:Barlow,sans-serif\">"]
    s.append(f"<defs><clipPath id=\"pc{w}{plot_h}\"><rect x=\"{gutter}\" y=\"{top}\" width=\"{pw}\" height=\"{plot_h}\"/></clipPath></defs>")
    # phase bands
    for i, (name, a, b) in enumerate(phases):
        a2, b2 = max(a, x0), min(b, x1)
        if b2 <= a2:
            continue
        if i % 2 == 1:
            s.append(f"<rect x=\"{fx(a2):.1f}\" y=\"{top}\" width=\"{fx(b2) - fx(a2):.1f}\" height=\"{plot_h}\" fill=\"rgba(229,185,85,.05)\"/>")
        if a > x0:
            s.append(f"<line x1=\"{fx(a):.1f}\" x2=\"{fx(a):.1f}\" y1=\"{top}\" y2=\"{top + plot_h}\" stroke=\"#e5b955\" stroke-opacity=\".7\" stroke-dasharray=\"2 3\"/>")
        s.append(f"<text x=\"{fx(a2) + 6:.1f}\" y=\"{top + 14}\" font-size=\"11\" font-weight=\"700\" letter-spacing=\".12em\" fill=\"#9a9484\">{name.upper()}</text>")
    # grid and y labels
    for v in yticks:
        y = fy(v)
        s.append(f"<line x1=\"{gutter}\" x2=\"{w - right}\" y1=\"{y:.1f}\" y2=\"{y:.1f}\" stroke=\"#1c2230\"/>")
        s.append(f"<text x=\"{gutter - 8}\" y=\"{y + 4:.1f}\" font-size=\"11\" fill=\"#9a9484\" text-anchor=\"end\" style=\"font-variant-numeric:tabular-nums\">{kfmt(v) if v else '0'}</text>")
    # x ticks
    step = nice_step(x1 - x0, pw)
    t = math.ceil(x0 / step) * step
    while t <= x1 + 0.01:
        x = fx(t)
        s.append(f"<line x1=\"{x:.1f}\" x2=\"{x:.1f}\" y1=\"{top + plot_h}\" y2=\"{top + plot_h + 4}\" stroke=\"#262e40\"/>")
        anchor = "start" if t == x0 else "end" if abs(t - x1) < .01 else "middle"
        s.append(f"<text x=\"{x:.1f}\" y=\"{top + plot_h + 17}\" font-size=\"11\" fill=\"#9a9484\" text-anchor=\"{anchor}\" style=\"font-variant-numeric:tabular-nums\">{mmss(t)}</text>")
        t += step
    s.append(f"<line x1=\"{gutter}\" x2=\"{w - right}\" y1=\"{top + plot_h + .5}\" y2=\"{top + plot_h + .5}\" stroke=\"#262e40\"/>")
    s.append(f"<g clip-path=\"url(#pc{w}{plot_h})\">")
    if area is not None:
        pts = [(fx(i), fy(v)) for i, v in enumerate(area) if v is not None and x0 - 1 <= i <= x1 + 1]
        d = f"M{pts[0][0]:.1f},{top + plot_h} " + " ".join(f"L{x:.1f},{y:.1f}" for x, y in pts) + f" L{pts[-1][0]:.1f},{top + plot_h} Z"
        s.append(f"<path d=\"{d}\" fill=\"rgba(229,185,85,.20)\" stroke=\"#e5b955\" stroke-width=\"1.5\"/>")
    for name, colour, vals in series:
        pts = [(fx(i), fy(v)) for i, v in enumerate(vals) if v is not None and x0 - 1 <= i <= x1 + 1]
        if not pts:
            continue
        d = "M" + " L".join(f"{x:.1f},{y:.1f}" for x, y in pts)
        wid, op = 1.5, 1
        if dim is not None:
            wid, op = (2.5, 1) if name == dim else (1.5, .25)
        s.append(f"<path d=\"{d}\" fill=\"none\" stroke=\"{colour}\" stroke-width=\"{wid}\" stroke-opacity=\"{op}\" stroke-linejoin=\"round\"/>")
    s.append("</g>")
    # deaths: full height dashed red, a skull in the lane
    for at, _who in deaths:
        if not x0 <= at <= x1:
            continue
        x = fx(at)
        s.append(f"<line x1=\"{x:.1f}\" x2=\"{x:.1f}\" y1=\"{top}\" y2=\"{top + plot_h}\" stroke=\"#ff5a5a\" stroke-width=\"1.5\" stroke-dasharray=\"4 3\"/>")
        s.append(f"<circle cx=\"{x:.1f}\" cy=\"{top / 2:.1f}\" r=\"10\" fill=\"#5a1f1f\" stroke=\"#ff5a5a\" stroke-opacity=\".6\"/>")
        s.append(f"<g transform=\"translate({x - 7:.1f},{top / 2 - 7:.1f}) scale(.5833)\" fill=\"none\" stroke=\"#ff5a5a\" stroke-width=\"2.4\" stroke-linecap=\"round\">{GLYPHS['skull']}</g>")
    # lane icons (taunts, casts); a marker within 22 px of a death steps 12 px right of it, with a leader
    death_x = [fx(a) for a, _w in deaths if x0 <= a <= x1]
    for at, src in list(taunts) + list(lane or []):
        if not x0 <= at <= x1:
            continue
        x = fx(at)
        near = [dx for dx in death_x if abs(dx - x) < 22]
        if near:
            x = near[0] + 22
            s.append(f"<line x1=\"{fx(at):.1f}\" x2=\"{x:.1f}\" y1=\"{top - 2}\" y2=\"{top / 2 + 10:.1f}\" stroke=\"#e5b955\" stroke-opacity=\".6\"/>")
        s.append(f"<line x1=\"{fx(at):.1f}\" x2=\"{fx(at):.1f}\" y1=\"{top}\" y2=\"{top + plot_h}\" stroke=\"#e5b955\" stroke-opacity=\".35\"/>")
        s.append(f"<rect x=\"{x - 10:.1f}\" y=\"{top / 2 - 10:.1f}\" width=\"20\" height=\"20\" rx=\"4\" fill=\"#0d111a\" stroke=\"#e5b955\" stroke-opacity=\".7\"/>")
        s.append(f"<image href=\"{src}\" x=\"{x - 9:.1f}\" y=\"{top / 2 - 9:.1f}\" width=\"18\" height=\"18\" clip-path=\"inset(0 round 3px)\"/>")
    # drag selection
    if drag is not None:
        a, b = fx(drag[0]), fx(drag[1])
        s.append(f"<rect x=\"{a:.1f}\" y=\"{top}\" width=\"{b - a:.1f}\" height=\"{plot_h}\" fill=\"rgba(229,185,85,.13)\"/>")
        for x in (a, b):
            s.append(f"<line x1=\"{x:.1f}\" x2=\"{x:.1f}\" y1=\"{top}\" y2=\"{top + plot_h}\" stroke=\"#e5b955\"/>")
        label = f"{mmss(drag[0])} to {mmss(drag[1])} · {drag[1] - drag[0]} s"
        lw = 8 + len(label) * 6.6
        mid = (a + b) / 2
        s.append(f"<rect x=\"{mid - lw / 2:.1f}\" y=\"{top + 8}\" width=\"{lw:.1f}\" height=\"22\" rx=\"6\" fill=\"#0d0f15\" stroke=\"#e5b955\" stroke-opacity=\".6\"/>")
        s.append(f"<text x=\"{mid:.1f}\" y=\"{top + 23}\" font-size=\"12\" font-weight=\"600\" fill=\"#f2eee4\" text-anchor=\"middle\" style=\"font-variant-numeric:tabular-nums\">{label}</text>")
    # crosshair with a dot on every line
    if cursor is not None:
        x = fx(cursor)
        s.append(f"<line x1=\"{x:.1f}\" x2=\"{x:.1f}\" y1=\"{top}\" y2=\"{top + plot_h}\" stroke=\"#e9e4d8\" stroke-opacity=\".7\" stroke-dasharray=\"3 3\"/>")
        for name, colour, vals in series:
            v = vals[int(cursor)] if int(cursor) < len(vals) else None
            if v is not None:
                s.append(f"<circle cx=\"{x:.1f}\" cy=\"{fy(v):.1f}\" r=\"3.5\" fill=\"{colour}\" stroke=\"#07090d\" stroke-width=\"1.5\"/>")
        s.append(f"<rect x=\"{x - 20:.1f}\" y=\"{top + plot_h + 4}\" width=\"40\" height=\"17\" rx=\"4\" fill=\"#e9e4d8\"/>")
        s.append(f"<text x=\"{x:.1f}\" y=\"{top + plot_h + 16.5}\" font-size=\"11\" font-weight=\"700\" fill=\"#07090d\" text-anchor=\"middle\" style=\"font-variant-numeric:tabular-nums\">{mmss(cursor)}</text>")
    s.append("</svg>")
    return "".join(s)


def overview_svg(w, total, window, phases=(), deaths=(), gutter=52, handle_h=36, phone=False):
    """The whole-fight strip under a chart: the raid's total as a muted area, the window as the brush."""
    right = 12
    pw = w - gutter - right
    h = handle_h
    peak = max(v for v in total if v is not None)
    fx = lambda t: gutter + t / DURATION * pw  # noqa: E731
    pts = [(fx(i), 4 + (h - 8) * (1 - v / peak)) for i, v in enumerate(total) if v is not None]
    d = f"M{gutter},{h - 4} " + " ".join(f"L{x:.1f},{y:.1f}" for x, y in pts) + f" L{pts[-1][0]:.1f},{h - 4} Z"
    a, b = fx(window[0]), fx(window[1])
    s = [f"<svg width=\"{w}\" height=\"{h + 18}\" viewBox=\"0 0 {w} {h + 18}\" style=\"display:block;font-family:Barlow,sans-serif\">"]
    s.append(f"<rect x=\"{gutter}\" y=\"0\" width=\"{pw}\" height=\"{h}\" rx=\"6\" fill=\"rgba(255,255,255,.025)\" stroke=\"#1c2230\"/>")
    s.append(f"<path d=\"{d}\" fill=\"rgba(154,148,132,.22)\" stroke=\"#9a9484\" stroke-opacity=\".5\"/>")
    for _n, start, _e in phases:
        if start > 0:
            s.append(f"<line x1=\"{fx(start):.1f}\" x2=\"{fx(start):.1f}\" y1=\"2\" y2=\"{h - 2}\" stroke=\"#e5b955\" stroke-opacity=\".5\" stroke-dasharray=\"2 3\"/>")
    for at, _w in deaths:
        s.append(f"<line x1=\"{fx(at):.1f}\" x2=\"{fx(at):.1f}\" y1=\"2\" y2=\"{h - 2}\" stroke=\"#ff5a5a\" stroke-width=\"1.5\"/>")
    if window != (0, DURATION):
        s.append(f"<rect x=\"{gutter}\" y=\"0\" width=\"{a - gutter:.1f}\" height=\"{h}\" fill=\"rgba(7,9,13,.6)\"/>")
        s.append(f"<rect x=\"{b:.1f}\" y=\"0\" width=\"{gutter + pw - b:.1f}\" height=\"{h}\" fill=\"rgba(7,9,13,.6)\"/>")
    s.append(f"<rect x=\"{a:.1f}\" y=\".5\" width=\"{b - a:.1f}\" height=\"{h - 1}\" fill=\"rgba(229,185,85,.10)\" stroke=\"#e5b955\" stroke-opacity=\".8\"/>")
    hw = 12 if phone else 8
    for x in (a, b):
        s.append(f"<rect x=\"{x - hw / 2:.1f}\" y=\"{h / 2 - 12}\" width=\"{hw}\" height=\"24\" rx=\"3\" fill=\"#e5b955\"/>")
        s.append(f"<line x1=\"{x - 1:.1f}\" x2=\"{x - 1:.1f}\" y1=\"{h / 2 - 5}\" y2=\"{h / 2 + 5}\" stroke=\"#0b0e14\"/><line x1=\"{x + 1:.1f}\" x2=\"{x + 1:.1f}\" y1=\"{h / 2 - 5}\" y2=\"{h / 2 + 5}\" stroke=\"#0b0e14\"/>")
    for t, anchor in ((0, "start"), (DURATION, "end")):
        s.append(f"<text x=\"{fx(t):.1f}\" y=\"{h + 14}\" font-size=\"11\" fill=\"#9a9484\" text-anchor=\"{anchor}\" style=\"font-variant-numeric:tabular-nums\">{mmss(t)}</text>")
    if window != (0, DURATION):
        for t, x in ((window[0], a), (window[1], b)):
            s.append(f"<text x=\"{x:.1f}\" y=\"{h + 14}\" font-size=\"11\" font-weight=\"700\" fill=\"#e5b955\" text-anchor=\"middle\" style=\"font-variant-numeric:tabular-nums\">{mmss(t)}</text>")
    s.append("</svg>")
    return "".join(s)


# --------------------------------------------------------------------- the dropdown parts
CHECK = glyph("check", 16)


def field(label, size="", state="", lead="", ph=False, width=None, chevron=True):
    cls = f"sel {size} {state} {'ph' if ph else ''}"
    style = f" style=\"width:{width}px\"" if width else ""
    tail = "<span class=\"chv\"><span class=\"spin\"></span></span>" if state == "loading" else (
        f"<span class=\"chv\">{glyph('chev-down', 16 if size != 'd' else 14)}</span>" if chevron else "")
    return f"<span class=\"{cls}\"{style}>{lead}<span style=\"overflow:hidden;text-overflow:ellipsis\">{label}</span>{tail}</span>"


def opt(label, *, desc="", meta="", lead="", sel=False, act=False, dis=False, colour=None):
    cls = "opt" + (" ic" if lead else "") + (" act" if act else "") + (" dis" if dis else "")
    name = f"<span style=\"color:{colour}\">{label}</span>" if colour else f"<span>{label}</span>"
    body = f"<span class=\"t\">{name}{f'<span class=\"ds\">{desc}</span>' if desc else ''}</span>"
    m = f"<span class=\"meta num\">{meta}</span>" if meta else ""
    ck = f"<span class=\"ck\">{CHECK}</span>" if sel else ""
    return f"<div class=\"{cls}\" role=\"option\">{lead}{body}{m}{ck}</div>"


def listbox(inner, *, top, width=None, search="", fade=False):
    style = f"top:{top}px;left:0" + (f";width:{width}px" if width else "")
    fd = "<div class=\"fade\"></div>" if fade else ""
    return f"<div class=\"lb\" role=\"listbox\" style=\"{style}\">{search}<div class=\"scroll\">{inner}{fd}</div></div>"


def labelled(label, inner):
    return f"<div style=\"display:flex;flex-direction:column;gap:6px\"><span class=\"ql\">{label}</span>{inner}</div>"


FIGHT_STYLES = [
    ("Patchwerk", "One target, standing still for the whole fight, with an execute phase below 25% health."),
    ("Execute heavy", "One target, standing still, with a wider execute phase: below 35% health instead of 25%."),
    ("Light movement", "One target, forced out of melee range for 5 seconds every 45 seconds; no parse has measured the cost of that yet."),
    ("Heavy movement", "One target, forced out of melee range for 5 seconds every 20 seconds; no parse has measured the cost of that yet."),
    ("Cleave, 2 targets", "Two targets, both present and taking damage for the whole fight, sharing the same 25% execute threshold as Patchwerk."),
    ("Cleave, 3 targets", "Three targets, all present and taking damage for the whole fight, sharing the same 25% execute threshold as Patchwerk."),
]


def dd_plain():
    rows = "".join(opt(n, desc=d, sel=(i == 0), act=(i == 1)) for i, (n, d) in enumerate(FIGHT_STYLES))
    return (labelled("Fight style", f"<div style=\"position:relative;width:260px\">{field('Patchwerk', 'c', 'open', width=260)}"
            f"{listbox(rows, top=50, width=400, fade=True)}</div>"))


def roster_lead(name):
    return crest(CLS[name], 20)


def dd_icons():
    groups = [("", [("All friendlies", None), ("All enemies", None)]),
              ("Tanks", [("Brokk", "warrior"), ("Thandril", "warrior")]),
              ("Healers", [("Seraphine", "priest"), ("Oakhorn", "druid")]),
              ("Damage", [("Garrosk", "warrior"), ("Morrigan", "mage"), ("Ashka", "warlock"), ("Velyn", "rogue"), ("Kael", "hunter")])]
    inner = ""
    for gi, (g, items) in enumerate(groups):
        if g:
            inner += f"<div class=\"grp{' rule' if gi else ''}\">{g}</div>"
        for name, cls in items:
            lead = crest(cls, 20) if cls else f"<span style=\"width:20px;display:flex;justify-content:center;color:var(--muted)\">{glyph('user', 16)}</span>"
            inner += opt(name, lead=lead, colour=CLASS_COLOR[cls] if cls else None, sel=(name == "Morrigan"), act=(name == "Ashka"))
    lead = crest("mage", 20)
    closed = field(f"<span style=\"color:{CLASS_COLOR['mage']}\">Morrigan</span>", "d", "open", lead=lead, width=200)
    scrolled = f"<div style=\"margin-top:-150px\">{inner}</div><div class=\"fade\" style=\"top:0;bottom:auto;transform:scaleY(-1)\"></div>"
    return (f"<div style=\"display:flex;align-items:center;gap:8px\"><span class=\"ql\">Source</span>"
            f"<div style=\"position:relative;width:200px\">{closed}{listbox(scrolled, top=42, width=240, fade=True)}</div></div>"
            "<p class=\"lbl\" style=\"margin-top:400px\">Opened scrolled so the selection (Morrigan) sits in view, fades at both ends.</p>")


def dd_search():
    search = ("<div class=\"lbsearch\"><span class=\"in\"><span style=\"color:var(--muted);display:flex\">" + glyph("search", 16) +
              "</span>black<span class=\"caret\"></span></span><span class=\"n num\">4 matches</span></div>")
    mk = lambda s: s.replace("Black", "<mark>Black</mark>")  # noqa: E731
    inner = ("<div class=\"grp\">Raids</div>" + opt(mk("Blackwing Lair"), act=True) +
             "<div class=\"grp rule\">Dungeons</div>" + opt(mk("Blackrock Depths")) + opt("Lower " + mk("Blackrock Spire")) + opt("Upper " + mk("Blackrock Spire")))
    return labelled("Source", f"<div style=\"position:relative;width:260px\">{field('Anywhere', 'c', 'open', width=260)}"
                    f"{listbox(search + inner, top=50, width=300)}</div>")


def dd_search_empty():
    search = ("<div class=\"lbsearch\"><span class=\"in\"><span style=\"color:var(--muted);display:flex\">" + glyph("search", 16) +
              "</span>naxx<span class=\"caret\"></span></span><span class=\"n num\">0 matches</span></div>")
    return labelled("Source", f"<div style=\"position:relative;width:260px\">{field('Anywhere', 'c', 'open', width=260)}"
                    f"{listbox(search + '<div class=lbempty>Nothing matches “naxx”.</div>', top=50, width=300)}</div>")


def board(body, label, width=None):
    style = f"width:{width}px" if width else ""
    return head("B", CSS) + f"<div class=\"board\" style=\"{style}\"><div class=\"dirbar\"><b>Controls and charts · 2026-10-30</b><span>{label}</span></div>{body}</div></body></html>"


def write(name, html):
    (RENDERS / f"{name}.html").write_text(html)


def page(inner, pad="32px 48px 48px"):
    return f"<div style=\"max-width:1344px;margin:0 auto;padding:{pad};width:100%\">{inner}</div>"


def build_dropdowns_open():
    cols = [("Plain, with descriptions · sim Fight style (44 px compact field)", dd_plain(), 610),
            ("With icons · report Source (dense 36 px field, groups, crests)", dd_icons(), 610),
            ("Searchable · sim Find an item → Source (typeahead filter, groups)", dd_search(), 610)]
    inner = "".join(f"<div style=\"flex:1;min-width:0\"><p class=\"cap\">{c}</p><div class=\"card\" style=\"padding:20px;min-height:{h}px;overflow:visible\">{b}</div></div>" for c, b, h in cols)
    second = (f"<div style=\"display:flex;gap:24px;margin-top:28px\"><div style=\"flex:1\"><p class=\"cap\">Searchable, no match</p><div class=\"card\" style=\"padding:20px;height:240px\">{dd_search_empty()}</div></div>"
              "<div style=\"flex:2\"><p class=\"cap\">Placement: the list lives in the top layer (popover), 4 px under the field, at least the field's width, at most 420 px; "
              "max height 360 px, then it scrolls with a 24 px fade; flips above when less than 240 px fits below; 8 px from every viewport edge.</p>"
              "<div class=\"card\" style=\"padding:20px;height:240px;position:relative;overflow:hidden\"><div style=\"position:absolute;left:20px;right:20px;bottom:12px;height:36px;border:1px dashed #262e40;border-radius:8px;display:flex;align-items:center;padding:0 12px\" class=\"lbl\">viewport bottom edge, 40 px below the field</div>"
              f"<div style=\"position:absolute;left:340px;top:150px\">{field('Every enemy', 'd', 'open', width=180)}"
              f"<div class=\"lb\" style=\"bottom:40px;left:0;width:220px\"><div class=\"scroll\">{opt('Every enemy', sel=True)}{opt('Ragnaros', act=True)}{opt('Son of Flame', meta='×8')}</div></div></div>"
              "<div class=\"lbl\" style=\"position:absolute;left:580px;top:90px\">← flipped: the list opens above the field, same 4 px gap</div></div></div></div>")
    write("dropdowns-open-1440", board(page(f"<div style=\"display:flex;gap:24px\">{inner}</div>{second}"), "the shared dropdown, open, fine pointer"))


def build_dropdown_states():
    states = [("Default", "", False), ("Hover", "hover", False), ("Focus-visible", "focus", False), ("Open", "open", False),
              ("Disabled", "dis", False), ("Error", "err", False), ("Loading", "loading", False), ("Empty", "dis", True), ("Placeholder", "", True)]

    def row(size, w):
        cells = ""
        for name, st, ph in states:
            label = {"Loading": "Loading…", "Empty": "No enemies", "Placeholder": "Choose a spec"}.get(name, "Patchwerk")
            extra = "<div class=\"errline\">Pick a fight style.</div>" if name == "Error" else ""
            f = field(label, size, st, ph=ph or name == "Loading", width=w, chevron=name != "Empty")
            cells += f"<div style=\"display:flex;flex-direction:column;gap:8px\"><span class=\"lbl\">{name}</span>{f}{extra}</div>"
        return cells
    sizes = [("Default field · 52 px · forms (account, guild settings, raid form)", "", 200),
             ("Compact · 44 px · filter bars, card heads, every field on a coarse pointer", "c", 200),
             ("Dense · 36 px · report and sim tool bars, fine pointer at 768 px and up only", "d", 170)]
    grid = "".join(f"<p class=\"cap\" style=\"margin-top:22px\">{c}</p><div style=\"display:grid;grid-template-columns:repeat(5,auto);gap:20px 24px;justify-content:start\">{row(s, w)}</div>" for c, s, w in sizes)
    with_icon = ("<p class=\"cap\" style=\"margin-top:28px\">The closed field with a lead icon (20 px crest, spec icon or spell icon; class-coloured name)</p><div style=\"display:flex;gap:24px\">"
                 + field(f"<span style=\"color:{CLASS_COLOR['mage']}\">Morrigan</span>", "c", "", lead=crest("mage", 20), width=220)
                 + field(f"<span style=\"color:{CLASS_COLOR['warrior']}\">Brokk</span>", "d", "hover", lead=crest("warrior", 20), width=180)
                 + "</div>")
    anatomy = ("<p class=\"cap\" style=\"margin-top:28px\">Option rows: plain 36 px · icon 40 px · with a description 52 px+ · group header 28 px · selected (gold check) · active (10% gold, 2 px gold edge) · disabled (45%) · trailing meta</p>"
               "<div class=\"lb\" style=\"position:relative;width:420px;min-width:0\"><div class=\"scroll\">"
               + opt("Every enemy", sel=True) + opt("Ragnaros", act=True) + opt("Son of Flame", meta="×8")
               + "<div class=\"grp rule\">Damage</div>" + opt("Garrosk", lead=crest("warrior", 20), colour=CLASS_COLOR["warrior"])
               + opt("Execute heavy", desc="One target, standing still, with a wider execute phase: below 35% health instead of 25%.")
               + opt("Dungeon pull", dis=True) + "</div></div>")
    write("dropdown-states-1440", board(page(grid + with_icon + anatomy), "the closed field: every state, three sizes; the option rows"))


# --------------------------------------------------------------------- the report around the charts
PHASES = [("Phase 1", 0, SUBMERGE), ("Submerge", SUBMERGE, DURATION)]
DEATHS = [(BROKK_DIES, "Brokk")]
TABS = ["Damage done", "Healing", "Damage taken", "Deaths", "Casts", "Buffs", "Threat", "More"]


def report_head(tab):
    tabs = "".join(f"<span class=\"{'on' if t == tab else ''}\">{t}{' ▾' if t == 'More' else ''}</span>" for t in TABS)
    return (f"<span class=\"eyebrow\">Report</span><div class=\"h1\" style=\"font-size:44px;margin-top:8px\">Molten Core</div>"
            "<p style=\"font-size:13px;color:var(--muted);margin-top:8px\">Ragnaros · Kill · 4:12 · 26 Oct 2026 · mock data</p>"
            f"<div style=\"display:flex;align-items:center;gap:16px;margin-top:20px\"><div class=\"tabs\" style=\"flex:1\">{tabs}</div>"
            f"<span class=\"ql\">Source</span>{field('All friendlies', 'd', width=170)}</div>")


def chip(name, value, *, off=False, hov=False, only=False):
    colour = CLASS_COLOR[CLS[name]]
    cls = "chip" + (" off" if off else "") + (" hov" if hov else "")
    o = "<span class=\"only\">Only</span>" if only else ""
    return f"<span class=\"{cls}\" style=\"color:{colour}\"><span class=\"sw\" style=\"background:{colour}\"></span>{name}<span class=\"v num\">{value}</span>{o}</span>"


def at(series, t):
    v = series[t]
    return v


def sorted_names(data, t, names):
    def key(n):
        v = data[n][t]
        return -(v if v is not None else -1)
    return sorted(names, key=key)


def threat_tip(t):
    alive = [n for n in CLS if THREAT[n][t] is not None and n not in ("Seraphine", "Oakhorn")]
    names = sorted_names(THREAT, t, alive)
    top = THREAT[names[0]][t]
    rows = "".join(f"<div class=\"r\" style=\"color:{CLASS_COLOR[CLS[n]]}\"><span class=\"sw\" style=\"background:{CLASS_COLOR[CLS[n]]}\"></span>{n}"
                   f"<span class=\"v num\">{THREAT[n][t]:,.0f}</span><span class=\"p num\">{THREAT[n][t] / top * 100:.0f}%</span></div>" for n in names)
    rows += f"<div class=\"r dead\"><span class=\"sw\" style=\"background:{CLASS_COLOR['warrior']};opacity:.4\"></span>Brokk<span class=\"v num\" style=\"color:var(--muted);font-weight:500\">Died 2:06</span><span class=\"p\"></span></div>"
    ev = f"<div class=\"ev\"><img src=\"{TAUNT}\" alt=\"\">Taunt · Thandril on Ragnaros</div>"
    return f"<div class=\"th\"><b class=\"num\">{mmss(t)}</b><span>Phase 1 · threat on Ragnaros</span></div>{ev}{rows}"


def threat_frame(width, *, window=(90, 160), cursor=127, phone=False):
    inner_w = width - 2
    visible = [n for n in CLS if n not in ("Seraphine", "Oakhorn")]
    t_end = window[1]
    order = sorted(visible, key=lambda n: -(THREAT[n][t_end] if THREAT[n][t_end] is not None else THREAT[n][BROKK_DIES]))

    def val(n):
        v = THREAT[n][t_end] if THREAT[n][t_end] is not None else THREAT[n][BROKK_DIES]
        return kfmt(v)
    series = [(n, CLASS_COLOR[CLS[n]], THREAT[n]) for n in visible]
    ymax = 80000
    total = [sum(DPS[n][t] or 0 for n in CLS) for t in range(DURATION + 1)]  # the overview is always the raid's damage per second
    plot_h = 180 if phone else 260
    gutter = 40 if phone else 52
    svg = chart_svg(inner_w, plot_h, window[0], window[1], series, ymax, yticks=[0, 20000, 40000, 60000, 80000],
                    taunts=[(TAUNT_AT, TAUNT), (40, TAUNT)], deaths=DEATHS, phases=PHASES, cursor=cursor, gutter=gutter)
    ov = overview_svg(inner_w, total, window, PHASES, DEATHS, gutter=gutter, phone=phone)
    if phone:
        chips = "".join(chip(n, kfmt(THREAT[n][cursor]) if THREAT[n][cursor] is not None else "died 2:06") for n in order[:6])
        chips += "<span class=\"chip more\">+2 more</span>"
        head_ = ("<div class=\"fhead\" style=\"flex-wrap:wrap\"><span class=\"stitle\">Threat on Ragnaros</span><span class=\"meta\">1:30 to 2:40</span></div>"
                 f"<div style=\"padding:0 12px 10px\">{field('Ragnaros', 'c', width=inner_w - 24)}<p class=\"lbl\" style=\"margin-top:6px\">tap opens the phone's own picker (native select, styled closed field)</p></div>")
        legend = (f"<div class=\"legend\"><span class=\"chip\" style=\"background:#e9e4d8;color:#07090d;height:24px;align-self:center\"><b class=\"num\">2:07</b></span>{chips}</div>")
        event = f"<div style=\"display:flex;align-items:center;gap:8px;height:28px;padding:0 12px;font-size:12px;color:var(--text)\"><img src=\"{TAUNT}\" width=\"16\" height=\"16\" style=\"border-radius:3px\" alt=\"\">Taunt · Thandril on Ragnaros</div>"
        return (f"<div class=\"frame\" style=\"width:{width}px\">{head_}{legend}<div style=\"position:relative\">{svg}"
                "</div>"
                f"{event}<div style=\"padding:4px 0 14px\">{ov}</div></div>")
    chips = "".join(chip(n, val(n)) for n in order) + chip("Seraphine", "6.4k", off=True) + chip("Oakhorn", "5.3k", off=True)
    chips += "<span class=\"ghost\" style=\"height:28px\">Show all</span>"
    enemy = f"<span class=\"ql\">Enemy</span>{field('Ragnaros', 'd', width=170)}"
    head_ = (f"<div class=\"fhead\"><span class=\"stitle\">Threat on Ragnaros</span><span class=\"meta\">Cumulative · 1:30 to 2:40</span>"
             f"<span class=\"right\">{enemy}<span class=\"ghost\">{glyph('x', 14)}Whole fight</span></span></div>")
    gx = gutter + (cursor - window[0]) / (window[1] - window[0]) * (inner_w - gutter - 12)
    tip = f"<div class=\"tip\" style=\"left:{gx + 16:.0f}px;top:4px\">{threat_tip(cursor)}</div>"
    return (f"<div class=\"frame\" style=\"width:{width}px\">{head_}<div class=\"legend\">{chips}</div>"
            f"<div style=\"position:relative\">{svg}{tip}</div><div style=\"padding:6px 0 16px\">{ov}</div></div>")


def threat_table():
    t = 160
    names = sorted_names(THREAT, t, [n for n in CLS if THREAT[n][t] is not None])
    top = THREAT[names[0]][t]
    rows = "".join(f"<tr><td style=\"color:{CLASS_COLOR[CLS[n]]};font-weight:600\"><span style=\"display:inline-flex;align-items:center;gap:10px\">{crest(CLS[n], 24)}{n}</span></td>"
                   f"<td><div class=\"bar\"><i style=\"width:{THREAT[n][t] / top * 100:.0f}%;background:{CLASS_COLOR[CLS[n]]}\"></i></div></td>"
                   f"<td class=\"r num\">{THREAT[n][t]:,.0f}</td><td class=\"r num\" style=\"color:var(--muted)\">{THREAT[n][t] / top * 100:.0f}%</td></tr>" for n in names[:4])
    return ("<div class=\"card\" style=\"margin-top:22px\"><div class=\"fhead\"><span class=\"stitle\">Threat table</span><span class=\"meta\">Table unchanged; it follows the window</span></div>"
            f"<table class=\"tbl\"><tr><th>Player</th><th>Threat</th><th class=\"r\">Standing</th><th class=\"r\">Share</th></tr>{rows}</table></div>")


NOTE = ("One cumulative line per player, on Ragnaros. A taunt in the game puts the taunter on top, and the base threat model does not: "
        "the lines are what damage and healing built, so a taunt mark is the moment the order stopped matching them.")


def build_threat(width_name, page_w):
    content = page_w
    body = (report_head("Threat") + f"<div style=\"margin-top:16px\">{threat_frame(content)}</div>"
            f"<p class=\"cnote\" style=\"margin-top:10px;max-width:900px\">{NOTE}</p>" + threat_table())
    write(width_name, board(page(body), "threat tab: hover tooltip at 2:07, taunt and death markers, zoomed to 1:30 to 2:40, the brush under it"))


def damage_frame(width, *, drag=(48, 80), dim="Morrigan", window=(0, DURATION), states=None):
    inner_w = width - 2
    visible = [n for n in CLS if n not in ("Seraphine", "Oakhorn")]
    order = sorted(visible, key=lambda n: -sum(v or 0 for v in DPS[n]))
    series = [(n, CLASS_COLOR[CLS[n]], DPS[n]) for n in visible]
    total = [sum(DPS[n][t] or 0 for n in CLS) for t in range(DURATION + 1)]
    raid = sum(total) / len(total)
    svg = chart_svg(inner_w, 200, window[0], window[1], series, 1200, yticks=[0, 400, 800, 1200],
                    lane=[(6, DEATH_WISH), (186, DEATH_WISH), (100, INNERVATE)], deaths=DEATHS, phases=PHASES, drag=drag, dim=dim)
    ov = overview_svg(inner_w, total, window, PHASES, DEATHS)
    chips = "".join(chip(n, f"{sum(v or 0 for v in DPS[n]) / (BROKK_DIES if n == 'Brokk' else DURATION):,.0f}", hov=(n == dim), only=(n == dim)) for n in order)
    chips += chip("Seraphine", "40", off=True) + chip("Oakhorn", "30", off=True)
    head_ = (f"<div class=\"fhead\"><span class=\"stitle\">Damage done</span><span class=\"meta num\">Raid {raid:,.0f} DPS · 5 s average</span>"
             f"<span class=\"right\"><span class=\"sw2\"><span class=\"tr\"></span>Log scale</span><span class=\"ql\">Window</span>{field('Whole fight', 'd', width=170)}</span></div>")
    return (f"<div class=\"frame\" style=\"width:{width}px\">{head_}<div class=\"legend\">{chips}</div>"
            f"<div style=\"position:relative\">{svg}</div><div style=\"padding:6px 0 16px\">{ov}</div></div>")


def build_damage():
    ann = ("<div class=\"lbl\" style=\"margin-top:10px\">Legend hover on Morrigan: her line 2.5 px, the rest at 25%; \"Only\" isolates. "
           "Drag in progress 0:48 → 1:20 (no tooltip while dragging); release sets the window, the URL and the tables. "
           "Lane: Death Wish (Garrosk) ×2, Innervate (Oakhorn), Brokk's death. Healers are off by default (outside the top 8).</div>")
    dps = {n: sum(v or 0 for v in DPS[n]) / DURATION for n in ("Garrosk", "Morrigan", "Ashka")}
    rows = "".join(f"<tr><td style=\"color:{CLASS_COLOR[CLS[n]]};font-weight:600\"><span style=\"display:inline-flex;align-items:center;gap:10px\">{crest(CLS[n], 24)}{n}</span></td>"
                   f"<td><div class=\"bar\"><i style=\"width:{d / dps['Garrosk'] * 100:.0f}%;background:{CLASS_COLOR[CLS[n]]}\"></i></div></td><td class=\"r num\">{d:,.0f}</td></tr>"
                   for n, d in dps.items())
    table = ("<div class=\"card\" style=\"margin-top:22px\"><div class=\"fhead\"><span class=\"stitle\">Damage done</span><span class=\"meta\">8 players · 1 death · 4:12</span></div>"
             f"<table class=\"tbl\"><tr><th>Player</th><th>Amount</th><th class=\"r\">DPS</th></tr>{rows}</table></div>")
    write("damage-timeline-1440", board(page(report_head("Damage done") + f"<div style=\"margin-top:16px\">{damage_frame(1248)}</div>{ann}{table}"),
                                        "damage done: one line per player, cast and death lane, legend hover, drag-to-zoom in progress"))


def small_frame(title, meta, body_html, *, right=""):
    return (f"<div class=\"frame\"><div class=\"fhead\"><span class=\"stitle\">{title}</span><span class=\"meta\">{meta}</span><span class=\"right\">{right}</span></div>{body_html}</div>")


def build_chart_states():
    w = 612
    skel_chips = "".join(f"<span class=\"skel\" style=\"width:{x}px;height:20px\"></span>" for x in (96, 110, 88, 102, 92))
    loading = small_frame("Damage done", "Loading the chart", f"<div class=\"legend\">{skel_chips}</div>"
                          "<div style=\"position:relative;height:248px;margin:0 12px 0 52px\"><div class=\"skel\" style=\"position:absolute;inset:26px 0 22px 0\"></div></div>"
                          "<div style=\"height:60px\"></div>", right=field("Whole fight", "d", width=170))
    axes = chart_svg(w - 2, 200, 200, 250, [], 1200, yticks=[0, 400, 800, 1200])
    empty = small_frame("Damage done", "Brokk · 3:20 to 4:10", f"<div class=\"legend\" style=\"min-height:42px\"></div><div style=\"position:relative\">{axes}"
                        "<div style=\"position:absolute;left:52px;right:12px;top:26px;height:200px;display:grid;place-items:center;font-size:14px;color:var(--muted)\">No damage in this window.</div></div><div style=\"height:60px\"></div>",
                        right=field("3:20 to 4:10", "d", width=170))
    error = small_frame("Damage done", "", "<div class=\"legend\" style=\"min-height:42px\"></div><div style=\"height:248px;display:flex;flex-direction:column;align-items:center;justify-content:center;gap:14px\">"
                        "<span style=\"font-size:14px;color:var(--text)\">The chart didn't load.</span><span class=\"btn btn-secondary sm\">Try again</span></div><div style=\"height:60px\"></div>")
    # a menu open over the chart: the list is in the top layer; the chart's tooltip and crosshair stand down
    groups = [("Tanks", [("Brokk", "warrior"), ("Thandril", "warrior")]), ("Damage", [("Garrosk", "warrior"), ("Morrigan", "mage"), ("Ashka", "warlock"), ("Velyn", "rogue")])]
    inner = opt("All friendlies", sel=True, lead=f"<span style=\"width:20px;display:flex;justify-content:center;color:var(--muted)\">{glyph('user', 16)}</span>")
    for gi, (g, items) in enumerate(groups):
        inner += f"<div class=\"grp rule\">{g}</div>" + "".join(opt(n, lead=crest(c, 20), colour=CLASS_COLOR[c], act=(n == "Garrosk")) for n, c in items)
    visible = [n for n in CLS if n not in ("Seraphine", "Oakhorn")]
    plot = chart_svg(w - 2, 200, 0, DURATION, [(n, CLASS_COLOR[CLS[n]], DPS[n]) for n in visible], 1200, yticks=[0, 400, 800, 1200], deaths=DEATHS, phases=PHASES)
    menu = (f"<div style=\"position:absolute;right:16px;top:6px;width:200px;z-index:7\">{field('All friendlies', 'd', 'open', width=200)}{listbox(inner, top=42, width=240, fade=True)}</div>")
    over = (f"<div class=\"frame\" style=\"overflow:visible\"><div class=\"fhead\"><span class=\"stitle\">Damage done</span><span class=\"meta\">Source menu open</span></div>{menu}"
            f"<div class=\"legend\">{''.join(chip(n, '') for n in visible[:5])}</div><div style=\"position:relative\">{plot}"
            "</div><div style=\"height:60px\"></div></div><p class=\"lbl\" style=\"margin-top:8px\">The pointer is over the plot beside the list: no tooltip and no crosshair while any menu is open.</p>")
    grid = (f"<div style=\"display:grid;grid-template-columns:{w}px {w}px;gap:24px\"><div><p class=\"cap\">Loading: the frame holds its final height; skeleton chips and plot, no spinner</p>{loading}</div>"
            f"<div><p class=\"cap\">Empty: axes stay, one centred line</p>{empty}</div>"
            f"<div><p class=\"cap\">Error: the chart chunk failed; the tables below still work</p>{error}</div>"
            f"<div><p class=\"cap\">A menu over the chart: the list is in the top layer, above everything</p>{over}</div></div>")
    write("chart-states-1440", board(page(grid), "chart states: loading, empty, error, and a menu opened over the plot"))


def build_phone():
    tabs = "".join(f"<span class=\"{'on' if t == 'Threat' else ''}\">{t}</span>" for t in ["Deaths", "Casts", "Buffs", "Threat", "More ▾"])
    body = (f"<div style=\"padding:16px 18px 32px\"><span class=\"eyebrow\">Report</span><div class=\"h1\" style=\"font-size:34px;margin-top:8px\">Molten Core</div>"
            "<p style=\"font-size:13px;color:var(--muted);margin-top:6px\">Ragnaros · Kill · 4:12 · mock data</p>"
            f"<div style=\"display:flex;flex-direction:column;gap:6px;margin-top:16px\"><span class=\"ql\">Source</span>{field('All friendlies', 'c', width=354)}</div>"
            f"<div class=\"tabs\" style=\"margin-top:12px;overflow:hidden\">{tabs}</div>"
            f"<div style=\"margin-top:12px\">{threat_frame(354, cursor=127, phone=True)}</div>"
            "<p class=\"cnote\" style=\"margin-top:10px\">Touch and hold the chart to read a moment.</p>"
            "<p class=\"lbl\" style=\"margin-top:6px\">Mock: a finger is holding at 2:07. The legend's values are the values at the finger; the time pill leads the legend; the event line names the taunt.</p>"
            "<div class=\"frame\" style=\"margin-top:16px;padding:12px;display:flex;align-items:center;justify-content:space-between\"><span class=\"num\" style=\"font-size:12px;color:var(--muted)\">Damage done · 4:12 · 1 death</span><span class=\"btn btn-secondary sm\">Show chart</span></div>"
            "</div>")
    write("controls-charts-phone", board(body, "390: native picker, scrub with legend readout", width=390))


def main():
    RENDERS.mkdir(exist_ok=True)
    build_dropdowns_open()
    build_dropdown_states()
    build_threat("threat-chart-1440", 1248)
    build_threat("threat-chart-2000", 1248)
    build_damage()
    build_chart_states()
    build_phone()


if __name__ == "__main__":
    main()
