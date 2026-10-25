"""Character history charts (design/specs/2026-10-24-character-history.md).

  python3 design/mocks/gen_character_history.py   -> renders/character-history-states.png

One sheet at 1440 wide holding 1440 boards and 390 frames: the premium owner's History section (score chart with
tier bands, dated points and a gap; per-boss parse chart with the picker; the fights behind each), every chart
and list state, the free owner's sentence and the visitor's sparkline, and the account row. Charts are inline SVG
drawn in pixel units from the example points below (the same geometry rules the spec fixes). The CHARACTER, every
score, parse, date and fight are EXAMPLES. Tier names and colours are docs/contracts/fixtures/score-labels.json's;
the copy mirrors the spec's section 6.
"""
import datetime
import os
import subprocess
import tempfile
from pathlib import Path

import gen_character_header as CH
import gen_nav as N
from mocklib import WEB, RENDERS, GOLD, MUTED, TEXT, BODY, BORDER, SOFT, RAISED, nav_phone

LINE_WARM = '#4a4030'
SCORE_MAX = 2100
TIER_FROM = [('common', 0), ('uncommon', 500), ('rare', 1000), ('epic', 1500), ('legendary', 2000)]
TIER_LABEL = dict(common='Common', uncommon='Uncommon', rare='Rare', epic='Epic', legendary='Legendary')
TIER_COLOR = dict(common='#ffffff', uncommon='#1eff00', rare='#0070dd', epic='#a335ee', legendary='#ff8000')
TIER_TEXT = dict(common='#ffffff', uncommon='#1eff00', rare='#3d94f0', epic='#b866f5', legendary='#ff8000')
PARSE_VAR = [(100, '#e5cc80'), (99, '#e268a8'), (95, '#ff8000'), (75, '#a335ee'), (50, '#0070ff'), (25, '#1eff00'), (0, '#8c8c8c')]
MONO = "'JetBrains Mono','SF Mono',Menlo,monospace"
BODYF = "'Barlow','Helvetica Neue',Arial,sans-serif"
MONTHS = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec']
GAP_LABEL_MIN = 72     # px: the label sits inside the band from this width, else in a line under the chart
UPSELL = 'History charts draw your score and parses over time.'


def tier_of(score):
    return [t for t, f in TIER_FROM if score >= f][-1]


def parse_color(p):
    return [c for lim, c in PARSE_VAR if round(p) >= lim][0]


def d(m, day, y=2026):
    return datetime.date(y, m, day)


def short(date, year=False):
    return f'{date.day} {MONTHS[date.month - 1]}' + (f' {date.year}' if year else '')


# ------------------------------------------------------------------ EXAMPLE data (API row shape, spec section 10)
SCORE = [  # date, score, visibility, fights that day
    (d(10, 12), 740, 'public', 3), (d(10, 15), 790, 'public', 2), (d(10, 19), 880, 'public', 4), (d(10, 26), 1010, 'public', 3),
    (d(11, 17), 1180, 'public', 2), (d(11, 18), 1240, 'private', 1), (d(11, 21), 1320, 'public', 4), (d(11, 24), 1390, 'public', 3),
    (d(11, 28), 1470, 'public', 5), (d(12, 1), 1530, 'public', 3), (d(12, 5), 1610, 'public', 6), (d(12, 8), 1700, 'public', 4),
    (d(12, 9), 1760, 'guild', 3), (d(12, 11), 1830, 'public', 2),
]
GAPS = [(d(10, 26), d(11, 17), 'No rated fights')]
ENC = {
    'Golemagg the Incinerator': [(d(11, 21), 54), (d(11, 28), 61), (d(12, 3), 58), (d(12, 8), 77), (d(12, 11), 88)],
    'Baron Geddon': [(d(11, 24), 49), (d(12, 1), 63), (d(12, 6), 71), (d(12, 9), 79)],
    'Garr': [(d(11, 19), 38), (d(11, 26), 52), (d(12, 2), 66), (d(12, 5), 70), (d(12, 8), 77)],
    'Magmadar': [(d(11, 17), 44), (d(11, 21), 57), (d(11, 28), 60), (d(12, 1), 74), (d(12, 4), 79), (d(12, 5), 83)],
    'Lucifron': [(d(11, 17), 51), (d(11, 18), 56), (d(11, 24), 63), (d(11, 28), 74), (d(12, 1), 81), (d(12, 3), 80)],
    'Gehennas': [(d(11, 28), 69)],
}
ORDER = sorted(ENC, key=lambda k: ENC[k][-1][0], reverse=True)
# the fights behind the score chart, newest first: boss, parse, DPS, date, execution, visibility
FIGHTS = [
    ('Golemagg the Incinerator', 88, 1402, d(12, 11), 94, 'public'), ('Sulfuron Harbinger', 66, 1318, d(12, 11), 87, 'public'),
    ('Baron Geddon', 79, 1380, d(12, 9), 91, 'guild'), ('Garr', 77, 1344, d(12, 8), 90, 'public'),
    ('Golemagg the Incinerator', 77, 1291, d(12, 8), 88, 'public'), ('Magmadar', 83, 1410, d(12, 5), 92, 'public'),
    ('Lucifron', 81, 1366, d(12, 1), 89, 'private'), ('Lucifron', 74, 1302, d(11, 28), 85, 'public'),
]
NARROW = [(d(6, 1), 520, 'public', 2), (d(6, 12), 610, 'public', 3), (d(7, 3), 700, 'public', 2), (d(7, 20), 820, 'public', 4), (d(8, 9), 860, 'public', 2),
          (d(9, 20), 1010, 'public', 3), (d(10, 5), 1090, 'public', 2), (d(11, 2), 1210, 'public', 4), (d(12, 1), 1400, 'public', 3), (d(12, 11), 1470, 'public', 2)]

CSS_RAW = '''<style>
.hs-pill{display:inline-flex;align-items:center;gap:8px;flex-shrink:0;height:44px;padding:0 14px;border:1px solid BORDER_;border-radius:4px;color:BODY_;font-size:12px;font-weight:700;letter-spacing:.06em;text-transform:uppercase;white-space:nowrap;background:transparent}
.hs-pill b{font-family:MONO_;font-weight:500;font-size:11px;color:MUTED_;letter-spacing:0}
.hs-pill.is-sel{border-color:GOLD_;color:TEXT_;background:#e5b95514}
.hs-pill.is-hover{border-color:LINEWARM_;color:TEXT_}
.hs-pill.is-active{background:#ffffff0d}
.hs-pill.is-focus{outline:2px solid GOLD_;outline-offset:2px}
.hs-step{width:44px;height:44px;display:grid;place-items:center;border:1px solid BORDER_;border-radius:4px;color:BODY_;background:transparent;font-size:16px}
.hs-step.is-hover{border-color:GOLD_;color:#f5d27a}
.hs-step.is-focus{outline:2px solid GOLD_;outline-offset:2px}
.hs-step.is-off{color:#5d5a50;border-color:SOFT_}
.hs-plot.is-focus{outline:2px solid GOLD_;outline-offset:2px;border-radius:4px}
.hs-row{display:grid;align-items:center;gap:12px;min-height:44px;padding:8px;border-bottom:1px solid SOFT_;font-size:14px}
.hs-row.is-sel{background:#e5b95514;box-shadow:inset 2px 0 0 GOLD_}
.hs-row a{color:GOLD_}
.hs-more{display:inline-flex;align-items:center;min-height:44px;padding:0 4px;color:TEXT_;font-size:13px;font-weight:600;text-decoration:underline;text-underline-offset:3px}
.hs-more.is-hover{color:#f5d27a}
.hs-fade{mask-image:linear-gradient(90deg,#000 calc(100% - 28px),transparent);-webkit-mask-image:linear-gradient(90deg,#000 calc(100% - 28px),transparent)}
</style>'''
CSS = (CSS_RAW.replace('BORDER_', BORDER).replace('BODY_', BODY).replace('MONO_', MONO).replace('MUTED_', MUTED).replace('GOLD_', GOLD)
       .replace('TEXT_', TEXT).replace('LINEWARM_', LINE_WARM).replace('SOFT_', SOFT))


def cap(t):
    return f'<span class="label" style="color:{GOLD}">{t}</span>'


def note(t):
    return f'<span style="font-size:12px;color:{MUTED};line-height:1.45">{t}</span>'


def cell(title, html, n='', w=None):
    size = f'width:{w}px;flex-shrink:0' if w else ''
    return f'<div style="display:flex;flex-direction:column;gap:8px;{size}">{cap(title)}{html}{note(n) if n else ""}</div>'


def h2(t):
    return f'<h2 class="display" style="font-size:18px;font-weight:700;letter-spacing:.1em;text-transform:uppercase;color:{TEXT}">{t}</h2>'


def h3(t):
    return f'<h3 style="font-size:15px;font-weight:600;color:{TEXT}">{t}</h3>'


# ------------------------------------------------------------------ chart geometry (spec 3.4)
def score_domain(vals):
    lo = (int(min(vals) - 60) // 100) * 100
    hi = -(-int(max(vals) + 60) // 100) * 100
    if hi - lo < 400:
        mid = (hi + lo) / 2
        lo, hi = int(mid - 200), int(mid + 200)
    return max(0, lo), min(SCORE_MAX, hi)


def parse_domain(vals):
    m = min(vals)
    lo = 0 if m < 30 else 25 if m < 55 else 50 if m < 80 else 75
    return lo, 100


def y_ticks_score(lo, hi):
    span = hi - lo
    step = 100 if span <= 600 else 200 if span <= 1200 else 500
    return list(range(((lo + step - 1) // step) * step, hi + 1, step))


def x_ticks(first, last, w_plot, phone):
    days = (last - first).days or 1
    n = 3 if phone else min(6, max(2, int(w_plot // 96)))
    return [first + datetime.timedelta(days=round(days * i / (n - 1))) for i in range(n)]


def chart(kind, pts, W, H, sel, hollow=(), gaps=(), preview=None, focus=False, phone=False):
    """kind 'score' | 'parse'. pts: [(date, value)]. Returns (the plot box, the under-chart gap lines)."""
    L, R, T, B = (36 if phone else 40), 14, 12, 26
    pw, ph = W - L - R, H - T - B
    vals = [v for _, v in pts]
    lo, hi = score_domain(vals) if kind == 'score' else parse_domain(vals)
    first, last = pts[0][0], pts[-1][0]
    span = max(1, (last - first).days)
    X = lambda dt: L + pw * (dt - first).days / span
    Y = lambda v: T + ph * (1 - (v - lo) / (hi - lo))
    o = []
    tier_lines = [f for _, f in TIER_FROM]
    if kind == 'score':
        for i, (t, f) in enumerate(TIER_FROM):
            top = TIER_FROM[i + 1][1] if i + 1 < len(TIER_FROM) else SCORE_MAX
            a, b = max(f, lo), min(top, hi)
            if a >= b:
                continue
            op = .04 if t == 'common' else .08
            o.append(f'<rect x="{L}" y="{Y(b):.1f}" width="{pw}" height="{Y(a) - Y(b):.1f}" fill="{TIER_COLOR[t]}" fill-opacity="{op}"/>')
            if Y(a) - Y(b) >= 22:
                o.append(f'<text x="{L + 8}" y="{Y(a) - 6:.1f}" font-family="{BODYF}" font-size="11" font-weight="600" fill="{TIER_TEXT[t]}" fill-opacity=".85">{TIER_LABEL[t]}</text>')
        ticks = y_ticks_score(lo, hi)
        for f in [f for f in tier_lines if lo < f < hi]:
            o.append(f'<line x1="{L}" x2="{L + pw}" y1="{Y(f):.1f}" y2="{Y(f):.1f}" stroke="{BORDER}"/>')
    else:
        ticks = [v for v in (0, 25, 50, 75, 100) if lo <= v <= hi]
    for v in ticks:
        if not (kind == 'score' and v in tier_lines):
            o.append(f'<line x1="{L}" x2="{L + pw}" y1="{Y(v):.1f}" y2="{Y(v):.1f}" stroke="{SOFT}"/>')
        o.append(f'<text x="{L - 8}" y="{Y(v) + 4:.1f}" text-anchor="end" font-family="{MONO}" font-size="11" fill="{MUTED}">{v}</text>')
    gap_line = []
    for g0, g1, why in gaps:
        x0, x1 = X(g0), X(g1)
        o.append(f'<rect x="{x0:.1f}" y="{T}" width="{x1 - x0:.1f}" height="{ph}" fill="#9a9484" fill-opacity=".08"/>')
        o.append(f'<line x1="{x0:.1f}" x2="{x0:.1f}" y1="{T}" y2="{T + ph}" stroke="{MUTED}" stroke-opacity=".4" stroke-dasharray="3 3"/>'
                 f'<line x1="{x1:.1f}" x2="{x1:.1f}" y1="{T}" y2="{T + ph}" stroke="{MUTED}" stroke-opacity=".4" stroke-dasharray="3 3"/>')
        if x1 - x0 >= GAP_LABEL_MIN:
            words = why.split(' ')
            mid = (len(words) + 1) // 2
            lines = [' '.join(words[:mid]), ' '.join(words[mid:])] if len(why) * 6 > x1 - x0 - 8 else [why]
            for k, ln in enumerate(lines):
                o.append(f'<text x="{(x0 + x1) / 2:.1f}" y="{T + 18 + 14 * k}" text-anchor="middle" font-family="{BODYF}" font-size="11" fill="{MUTED}">{ln}</text>')
        else:
            gap_line.append(f'Gap {short(g0)} to {short(g1)}: {why}')
    cur = preview if preview is not None else sel
    guide = BORDER if preview is None else LINE_WARM
    o.append(f'<line x1="{X(pts[cur][0]):.1f}" x2="{X(pts[cur][0]):.1f}" y1="{T}" y2="{T + ph}" stroke="{guide}"/>')
    runs, run = [], [0]
    gap_pairs = set((a, b) for a, b, _ in gaps)
    for i in range(1, len(pts)):
        if (pts[i - 1][0], pts[i][0]) in gap_pairs:
            runs.append(run)
            run = []
        run.append(i)
    runs.append(run)
    for r in runs:
        if len(r) > 1:
            path = ' '.join(('M' if k == 0 else 'L') + f'{X(pts[i][0]):.1f},{Y(pts[i][1]):.1f}' for k, i in enumerate(r))
            o.append(f'<path d="{path}" fill="none" stroke="{GOLD}" stroke-width="2" stroke-linejoin="round" stroke-linecap="round"/>')
    dense = len(pts) > 40
    for i, (dt, v) in enumerate(pts):
        if i == sel or (dense and i != len(pts) - 1):
            continue
        if i in hollow:
            o.append(f'<circle cx="{X(dt):.1f}" cy="{Y(v):.1f}" r="3.5" fill="{RAISED}" stroke="{GOLD}" stroke-width="1.5"/>')
        else:
            o.append(f'<circle cx="{X(dt):.1f}" cy="{Y(v):.1f}" r="3" fill="{GOLD}"/>')
    sv = pts[sel][1]
    col = TIER_COLOR[tier_of(sv)] if kind == 'score' else parse_color(sv)
    fill = RAISED if sel in hollow else col
    o.append(f'<circle cx="{X(pts[sel][0]):.1f}" cy="{Y(sv):.1f}" r="5" fill="{fill}" stroke="{TEXT}" stroke-width="2"/>')
    for k, tk in enumerate(x_ticks(first, last, pw, phone)):
        anchor = 'start' if k == 0 else 'end' if tk == last else 'middle'
        x = X(tk) if anchor == 'middle' else (L if anchor == 'start' else L + pw)
        o.append(f'<text x="{x:.1f}" y="{H - 6}" text-anchor="{anchor}" font-family="{MONO}" font-size="11" fill="{MUTED}">{short(tk)}</text>')
    ring = ' is-focus' if focus else ''
    svg = f'<svg width="{W}" height="{H}" role="img" style="display:block;overflow:visible">{"".join(o)}</svg>'
    return f'<div class="hs-plot{ring}" style="width:{W}px;height:{H}px">{svg}</div>', gap_line


# ------------------------------------------------------------------ pieces of the section
def readout(kind, pts, sel, step_modes=('', '')):
    dt, v = pts[sel]
    if kind == 'score':
        t = tier_of(v)
        big = f'<span class="mono" style="font-size:28px;line-height:1;font-weight:600;color:{TIER_TEXT[t]}">{v}</span>'
        word = f'<span style="font-size:15px;font-weight:600;color:{TIER_TEXT[t]}">{TIER_LABEL[t]}</span>'
        fights = next((f for dd, _, _, f in SCORE if dd == dt), 1)
        noun = 'fight' if fights == 1 else 'fights'
        sub = f'{short(dt, True)} &middot; {fights} {noun}'
    else:
        big = f'<span class="mono" style="font-size:28px;line-height:1;font-weight:600;color:{parse_color(v)}">{v}</span>'
        word = f'<span style="font-size:15px;color:{MUTED}">parse</span>'
        sub = f'{short(dt, True)} &middot; kill {sel + 1} of {len(pts)}'
    prev_c = 'is-off' if sel == 0 else ''
    next_c = 'is-off' if sel == len(pts) - 1 else ''
    pm, nm = step_modes
    pmc = 'is-' + pm if pm else ''
    nmc = 'is-' + nm if nm else ''
    steps = (f'<span style="display:flex;gap:8px;flex-shrink:0"><span class="hs-step {prev_c} {pmc}">&lsaquo;</span>'
             f'<span class="hs-step {next_c} {nmc}">&rsaquo;</span></span>')
    return (f'<div style="display:flex;align-items:center;justify-content:space-between;gap:12px;min-height:56px">'
            f'<span style="display:flex;flex-direction:column;gap:6px;min-width:0"><span style="display:flex;align-items:baseline;gap:10px">{big}{word}</span>'
            f'<span class="mono" style="font-size:12px;color:{MUTED}">{sub}</span></span>{steps}</div>')


def change_caption(pts):
    ch = pts[-1][1] - pts[0][1]
    word = f'Up {ch}' if ch > 0 else f'Down {-ch}' if ch < 0 else 'No change'
    return f'{word} since {short(pts[0][0])}'


def fight_rows(rows, phone, lead='boss', sel_dates=()):
    grid = 'minmax(0,1fr) auto auto auto' if phone else 'minmax(0,1fr) 64px 88px 112px 72px'
    first_head = 'Boss' if lead == 'boss' else 'Date'
    if phone:
        heads = [first_head, 'Parse %', 'DPS', 'Execution']
    else:
        heads = [first_head, 'Parse %', 'DPS', 'Date' if lead == 'boss' else '', 'Execution']
    right = (1, 2, len(heads) - 1)
    head = ''.join(f'<span style="{"text-align:right" if i in right else ""}">{c}</span>' for i, c in enumerate(heads))
    out = [f'<div class="label" style="display:grid;grid-template-columns:{grid};gap:12px;min-height:36px;align-items:end;padding:0 8px 8px;border-bottom:1px solid {SOFT}">{head}</div>']
    for boss, p, dps, dt, ex, vis in rows:
        sel = ' is-sel' if dt in sel_dates else ''
        visw = f'<span style="color:{MUTED};font-size:12px;margin-left:8px">{vis.capitalize()}</span>' if vis != 'public' else ''
        lead_style = f'font-family:{MONO};font-size:13px;color:{MUTED};' if lead == 'date' else ''
        lead_txt = boss if lead == 'boss' else short(dt, True)
        c1 = f'<span style="display:flex;align-items:baseline;min-width:0"><a style="white-space:nowrap;overflow:hidden;text-overflow:ellipsis;{lead_style}">{lead_txt}</a>{visw}</span>'
        c2 = f'<span class="mono" style="text-align:right;font-size:13px;color:{parse_color(p)}">{p}</span>'
        c3 = f'<a class="mono" style="text-align:right;font-size:14px">{dps:,}</a>'
        c4 = f'<span class="mono" style="font-size:13px;color:{MUTED}">{short(dt, True)}</span>' if lead == 'boss' else '<span></span>'
        c5 = f'<a class="mono" style="text-align:right;font-size:13px;font-weight:600">{ex}%</a>'
        cols = c1 + c2 + c3 + c5 if phone else c1 + c2 + c3 + c4 + c5
        out.append(f'<div class="hs-row{sel}" style="grid-template-columns:{grid}">{cols}</div>')
    return ''.join(out)


def show_all(n, word, mode=''):
    m = f' is-{mode}' if mode else ''
    return f'<span class="hs-more{m}">Show all {n} {word}</span>'


def picker(selected, modes=None):
    modes = modes or dict()
    pills = []
    for name in ORDER:
        cls = 'hs-pill' + (' is-sel' if name == selected else '') + (f' is-{modes[name]}' if name in modes else '')
        pills.append(f'<span class="{cls}">{name.split(" the ")[0]}<b>{len(ENC[name])}</b></span>')
    return f'<div class="hs-fade" style="display:flex;gap:8px;overflow:hidden;min-height:44px">{"".join(pills)}</div>'


def skeleton_box(H):
    return f'<span class="sk" style="display:block;width:100%;height:{H}px;border-radius:4px"></span>'


def msg(text, minh=44):
    return f'<p style="margin:0;display:flex;align-items:center;min-height:{minh}px;font-size:14px;color:{MUTED};line-height:1.4">{text}</p>'


def retry(text):
    return (f'<div style="display:flex;align-items:center;flex-wrap:wrap;gap:12px;min-height:44px;font-size:14px;color:{MUTED}">{text}'
            f'<span style="display:inline-flex;align-items:center;min-height:44px;padding:0 12px;border:1px solid #3a3326;border-radius:4px;font-size:13px;font-weight:600;color:{TEXT}">Try again</span></div>')


# ------------------------------------------------------------------ the two columns
def score_col(W, H, phone, sel=None, preview=None, focus=False, data=None, gaps=None, step=('', ''), more_mode=''):
    data = data or SCORE
    gaps = GAPS if gaps is None else gaps
    pts = [(a, b) for a, b, _, _ in data]
    sel = len(pts) - 1 if sel is None else sel
    hollow = [i for i, (_, _, vis, _) in enumerate(data) if vis != 'public']
    cap_txt = "One point a day: the day's last rated fight." + (' Rings are private or guild reports.' if hollow else '')
    svg, gl = chart('score', pts, W, H, sel, hollow, gaps, preview, focus, phone)
    gl_html = ''.join(f'<span style="font-size:12px;color:{MUTED}">{g}</span>' for g in gl)
    shown = preview if preview is not None else sel
    body = (readout('score', pts, shown, step) + svg
            + f'<span style="font-size:12px;color:{TEXT}">{change_caption(pts)}</span>' + gl_html)
    lst = (f'<div style="display:flex;flex-direction:column;gap:2px;margin-top:6px"><span class="label">Fights behind the score</span>'
           f'{fight_rows(FIGHTS, phone, "boss", (pts[sel][0],))}{show_all(31, "fights", more_mode)}</div>')
    return (f'<div style="display:flex;flex-direction:column;gap:12px;min-width:0;flex:1">{h3("Score over time")}'
            f'<span style="font-size:12px;color:{MUTED};line-height:1.45;min-height:44px;display:flex;align-items:center">{cap_txt}</span>{body}{lst}</div>')


def parse_col(W, H, phone, name=None, sel=None, focus=False, picker_modes=None, body=None, list_html=None, pick=True):
    name = name or ORDER[0]
    pts = list(ENC[name])
    sel = len(pts) - 1 if sel is None else sel
    if body is None:
        svg, _ = chart('parse', pts, W, H, sel, (), (), None, focus, phone)
        body = readout('parse', pts, sel) + svg + f'<span style="font-size:12px;color:{TEXT}">{change_caption(pts)}</span>'
        rows = [(name, p, 1300 + p * 3, dt, 70 + (p // 4), 'public') for dt, p in reversed(ENC[name])]
        list_html = (f'<div style="display:flex;flex-direction:column;gap:2px;margin-top:6px"><span class="label">Kills of {name.split(" the ")[0]}</span>'
                     f'{fight_rows(rows, phone, "date", (pts[sel][0],))}</div>')
    ctrl = picker(name, picker_modes) if pick else ''
    return f'<div style="display:flex;flex-direction:column;gap:12px;min-width:0;flex:1">{h3("Parse by boss")}{ctrl}{body}{list_html or ""}</div>'


def panel_shell(inner, width):
    pad = 16 if width < 700 else 20
    return f'<div style="width:{width}px;max-width:100%;border:1px solid {BORDER};border-radius:6px;background:{RAISED};padding:{pad}px">{inner}</div>'


def section(width, inner, caption='Includes your private and guild reports. Only you see this section.'):
    return (f'<section style="display:flex;flex-direction:column;gap:8px;width:{width}px">{h2("History")}'
            f'<p style="margin:0;font-size:13px;color:{MUTED}">{caption}</p>{panel_shell(inner, width)}</section>')


def two_cols(a, b, width):
    if width >= 1200:
        return f'<div style="display:flex;gap:32px;align-items:flex-start">{a}{b}</div>'
    return f'<div style="display:flex;flex-direction:column;gap:28px">{a}{b}</div>'


# ------------------------------------------------------------------ stubs for the neighbours (the shipped card, abbreviated)
def spark(W):
    pts = [v for _, v, _, _ in SCORE]
    n = len(pts)
    lo, hi = 600, 1900
    y = lambda v: 6 + 28 * (1 - (v - lo) / (hi - lo))
    path = ' '.join(('M' if i == 0 else 'L') + f'{W * i / (n - 1):.1f},{y(v):.1f}' for i, v in enumerate(pts))
    svg = (f'<svg width="{W}" height="40" style="display:block;overflow:visible"><path d="{path}" fill="none" stroke="{GOLD}" stroke-width="2" stroke-linejoin="round"/>'
           f'<circle cx="{W}" cy="{y(pts[-1]):.1f}" r="4" fill="{TIER_COLOR["epic"]}"/></svg>')
    return (f'<div style="display:flex;flex-direction:column;gap:4px;width:{W}px">{svg}<div style="display:flex;justify-content:space-between;font-size:12px">'
            f'<span style="color:{TEXT}">Up 1090 since 12 Oct</span><span class="mono" style="color:{MUTED}">11 Dec</span></div></div>')


def score_card_stub(width, phone):
    size = 40 if phone else 48
    big = f'<span class="mono" style="font-size:{size}px;font-weight:600;line-height:1;color:{TIER_COLOR["epic"]}">1830</span>'
    head = (f'<div style="display:flex;flex-direction:column;gap:6px"><span style="display:flex;align-items:baseline;gap:8px">{big}'
            f'<span class="mono" style="color:{MUTED};font-size:14px">of 2100</span><span style="font-size:15px;font-weight:600;color:{TIER_TEXT["epic"]}">Epic</span></span>'
            f'<span style="font-size:13px;color:{MUTED}">170 to Legendary (2000)</span></div>')
    strip = ''.join(f'<span style="flex:1;height:6px;background:{TIER_COLOR[t]};opacity:{1 if t == "epic" else .35}"></span>' for t, _ in TIER_FROM)
    if phone:
        inner = (f'{head}<div style="display:flex;gap:2px">{strip}</div>{spark(width - 32)}'
                 f'<span style="font-size:12px;color:{MUTED}">The five part bars and Most to gain follow (shipped, not redrawn).</span>')
    else:
        col = (width - 40 - 32) // 2
        inner = (f'<div style="display:flex;gap:32px"><div style="flex:1;display:flex;flex-direction:column;gap:16px">{head}'
                 f'<span style="font-size:12px;color:{MUTED}">The five part bars and Most to gain (shipped, not redrawn).</span></div>'
                 f'<div style="flex:1;display:flex;flex-direction:column;gap:12px">{spark(col)}<div style="display:flex;gap:2px">{strip}</div></div></div>')
    pad = 16 if phone else 20
    return (f'<div style="width:{width}px;border:1px solid {BORDER};border-radius:6px;background:{RAISED};padding:{pad}px;'
            f'display:flex;flex-direction:column;gap:16px;box-shadow:inset 0 1px 0 #e5b95559">{inner}</div>')


def score_block(width, phone, tail=''):
    return (f'<section style="display:flex;flex-direction:column;gap:8px;width:{width}px">{h2("Forever Score")}'
            f'<p style="margin:0;font-size:13px;color:{MUTED}">Your whole record, as of 11 Dec.</p>{score_card_stub(width, phone)}{tail}</section>')


def upsell(phone):
    mh = 44 if phone else 0
    link = (f'<a style="color:{TEXT};text-decoration:underline;text-underline-offset:2px;display:inline-flex;align-items:center;min-height:{mh}px">See plans</a>')
    return f'<p style="margin:0;font-size:13px;color:{MUTED};line-height:1.45;display:flex;flex-wrap:wrap;gap:0 6px;align-items:center;min-height:44px">{UPSELL} {link}</p>'


# ------------------------------------------------------------------ boards
W_DESK, H_DESK = 588, 294
W_PH, H_PH = 320, 213


def desktop_full():
    return section(1248, two_cols(score_col(W_DESK, H_DESK, False), parse_col(W_DESK, H_DESK, False), 1248))


def phone_full():
    return section(354, two_cols(score_col(W_PH, H_PH, True), parse_col(W_PH, H_PH, True), 320))


def frame(width, inner):
    phone = width < 700
    bar = nav_phone(True, dict([('class', 'mage')])) if phone else ''
    gut = CH.gutter(width)
    pad = f'padding:22px {gut}px 28px;' if phone else f'padding:22px 48px 28px;max-width:{CH.PAGE_MAX}px;margin:0 auto;'
    gap = 22 if phone else 32
    return (f'<div style="width:{width}px;border:1px solid {SOFT};background:var(--bg);flex-shrink:0">{bar}'
            f'<div style="{pad}display:flex;flex-direction:column;gap:{gap}px">{inner}</div></div>')


def states_desktop():
    mid = 8
    c = lambda t, h, n='', w=628: cell(t, h, n, w)
    items = [
        c('Score chart: hover preview on 28 Nov (readout previews, the list stays on the selection)',
          panel_shell(score_col(W_DESK, H_DESK, False, preview=mid), 628), 'The previewed point shows in the readout and a warm guide line; leaving restores the selected point.'),
        c('Score chart: plot focus-visible, step button focus, Show all hover',
          panel_shell(score_col(W_DESK, H_DESK, False, sel=mid, focus=True, step=('focus', ''), more_mode='hover'), 628),
          'Arrow keys move the selection; the list highlight follows. Gold 2 px outline, 2 px offset.'),
    ]
    one_pt = (f'{h3("Score over time")}<div style="display:flex;flex-direction:column;gap:10px;margin-top:12px">'
              f'<span style="display:flex;align-items:baseline;gap:10px"><span class="mono" style="font-size:28px;font-weight:600;color:{TIER_TEXT["epic"]}">1830</span>'
              f'<span style="font-size:15px;font-weight:600;color:{TIER_TEXT["epic"]}">Epic</span></span>'
              f'<span style="font-size:13px;color:{TEXT}">First score recorded 11 Dec. A line needs a second point.</span></div>')
    items.append(c('Score: one point (no line, no empty box)', panel_shell(one_pt, 628),
                   'The chart box is not drawn; the section collapses once from its reserved height (accepted, spec 3.11).'))
    fl_data = [row for row in SCORE if row[0] < d(10, 26) or row[0] >= d(11, 21)]
    items.append(c('Gap with another reason: Report flagged',
                   panel_shell(score_col(W_DESK, H_DESK, False, data=fl_data, gaps=[(d(10, 26), d(11, 21), 'Report flagged')]), 628),
                   'The owner sees their own flagged report named. A flagged report never counted toward the score.'))
    items.append(c('Narrow gap (20 days in a 193-day window): the label moves under the chart',
                   panel_shell(score_col(W_DESK, H_DESK, False, data=NARROW, gaps=[(d(7, 20), d(8, 9), 'Report flagged')]), 628),
                   'The band is 55 px, under the 72 px floor, so it carries no inside label and the reason is one muted 12 px line.'))
    items.append(c('Parse: another boss; pills at hover and active',
                   panel_shell(parse_col(W_DESK, H_DESK, False, name='Magmadar', picker_modes=dict(Garr='hover', Lucifron='active')), 628),
                   'Selected pill: gold border, 8 percent gold ground. Hover: warm border. The row scrolls sideways and fades at the right edge.'))
    items.append(c('Parse: picker focus-visible, middle point selected',
                   panel_shell(parse_col(W_DESK, H_DESK, False, name='Baron Geddon', sel=1, picker_modes=dict(Garr='focus')), 628),
                   'Left and Right move between pills (roving tabindex); Enter or Space selects.'))
    one_row = fight_rows([('Gehennas', 69, 1355, d(11, 28), 86, 'public')], False, 'date')
    onekill = (msg('One kill so far. A line needs two.')
               + f'<div style="display:flex;flex-direction:column;gap:2px;margin-top:6px"><span class="label">Kills of Gehennas</span>{one_row}</div>')
    items.append(c('Parse: a boss with one kill (a point list, not a chart)', panel_shell(parse_col(W_DESK, H_DESK, False, name='Gehennas', body=onekill, list_html=''), 628),
                   'Every boss with a kill is in the picker; one kill draws no chart.'))
    nokill = msg('No kills on record yet. Parse charts start with your first kill.')
    items.append(c('Parse: no kills on record', panel_shell(parse_col(W_DESK, H_DESK, False, body=nokill, list_html='', pick=False), 628),
                   'No picker, no chart, no list.'))
    one_col = (f'<div style="flex:1;display:flex;flex-direction:column;gap:12px">{CH.skel(140, 15)}<span style="min-height:44px;display:block"></span>'
               f'{CH.skel(120, 28)}{skeleton_box(H_DESK)}</div>')
    ld = f'<div style="display:flex;gap:32px">{one_col}{one_col}</div>'
    items.append(c('Loading (the reserved box, last height)', panel_shell(ld, 1000),
                   'Skeleton titles, readout and chart box at the reserved size. sr-only: Reading your history...', 1000))
    items.append(c('Error (the sparkline above is untouched)', panel_shell(retry('Your history did not load.'), 628), 'LoadError with Try again; it re-reads in place.'))
    items.append(c('Empty: a score, but no fight carries a history point', panel_shell(msg('No history yet. Each rated fight you upload adds a point.'), 628), 'EmptyState, no action.'))
    return '<div style="display:flex;flex-wrap:wrap;gap:24px;align-items:flex-start">' + ''.join(items) + '</div>'


def states_phone():
    gap = [(d(7, 20), d(8, 9), 'No rated fights')]
    items = [
        cell('390: score chart, 28 Nov selected, gap label inside', panel_shell(score_col(W_PH, H_PH, True, sel=8), 354), '', 354),
        cell('390: parse chart, Baron Geddon', panel_shell(parse_col(W_PH, H_PH, True, name='Baron Geddon'), 354), '', 354),
        cell('390: narrow gap, label under the chart', panel_shell(score_col(W_PH, H_PH, True, data=NARROW, gaps=gap), 354), '', 354),
    ]
    return '<div style="display:flex;gap:24px;align-items:flex-start">' + ''.join(items) + '</div>'


def free_visitor():
    free_d = frame(1344, score_block(1248, False, upsell(False)))
    vis_d = frame(1344, score_block(1248, False))
    free_p = frame(390, score_block(354, True, upsell(True)))
    vis_p = frame(390, score_block(354, True))
    desk = (f'<div style="display:flex;flex-direction:column;gap:22px">{cell("Free owner, 1440: the card sparkline, then one sentence", free_d)}'
            f'{cell("Visitor, 1440: the card alone, nothing added", vis_d)}</div>')
    phone = (f'<div style="display:flex;gap:24px;align-items:flex-start">{cell("Free owner, 390", free_p, "", 392)}'
             f'{cell("Visitor, 390", vis_p, "", 392)}</div>')
    return desk + phone


def account_row(label, value, link=False, muted=False):
    color = MUTED if muted else TEXT
    deco = 'text-decoration:underline;' if link else ''
    return (f'<div style="display:grid;grid-template-columns:10rem minmax(0,1fr);column-gap:12px;align-items:baseline;min-height:28px;font-size:14px">'
            f'<span style="color:{MUTED}">{label}</span><span style="color:{color};{deco}">{value}</span></div>')


def account_card(value, link):
    rows = (account_row('Server sims', 'On · 4 runs this month', True) + account_row('History charts', value, link)
            + account_row('Discord notifications', 'Ships December', muted=True))
    return (f'<div style="width:395px;border:1px solid {BORDER};border-radius:6px;background:{RAISED};padding:22px 24px;display:flex;flex-direction:column;gap:6px">{rows}</div>')


def account_rows():
    return ('<div style="display:flex;gap:24px;align-items:flex-start">' + cell('Account, Premium', account_card('Open on your main character', True))
            + cell('Account, Premium, no character yet', account_card('On', False)) + cell('Account, free', account_card('Sparkline only', False)) + '</div>')


def sheet():
    gut = CH.gutter(1440)
    maxw = CH.PAGE_MAX
    page_d = frame(1440, score_block(1248, False) + desktop_full())
    page_p = frame(390, score_block(354, True) + phone_full())
    wrap = lambda b: f'<div style="max-width:{maxw}px;margin:0 auto;padding:36px {gut}px;display:flex;flex-direction:column;gap:24px">{b}</div>'
    title = (f'<h1 class="display" style="font-size:22px;color:{TEXT}">Character history: every state</h1>'
             f'{h2("Premium owner, 1440 (the History section under the Forever Score card)")}')
    phone_cell = cell('390, the section stacked (score, then parse)', page_p, '', 392)
    second = (f'{h2("Premium owner, 390")}<div style="display:flex;gap:24px;align-items:flex-start">{phone_cell}</div>'
              f'{h2("Chart and list states (desktop columns)")}{states_desktop()}{h2("Chart states (390)")}{states_phone()}'
              f'{h2("Free owner and visitor")}{free_visitor()}{h2("Account page: the History row")}{account_rows()}')
    body = wrap(title) + page_d + wrap(second)
    return N.doc(CH.CSS + CSS + body, 1440)


def shoot(html, name, width):
    scratch = Path(tempfile.mkdtemp(prefix='gen_hist_'))
    src = scratch / f'{name}.html'
    src.write_text(html)
    RENDERS.mkdir(exist_ok=True)
    out = RENDERS / f'{name}.png'
    local = WEB / 'node_modules/.bin/playwright'
    override = os.environ.get('FS_PLAYWRIGHT')
    runner = [override] if override else [str(local)] if local.exists() else ['npx', 'playwright']
    size = f'--viewport-size={width},900'
    subprocess.run(runner + ['screenshot', '--browser=chromium', size, '--full-page', '--wait-for-timeout=3000', src.as_uri(), str(out)],
                   check=True, capture_output=True, cwd=WEB)
    return out


if __name__ == '__main__':
    print(shoot(sheet(), 'character-history-states', 1440))
