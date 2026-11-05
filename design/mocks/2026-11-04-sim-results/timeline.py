"""The timeline region as SVG: lanes, ticks, bars, the rage line, the axis, the brush (spec 3.7)."""
import html as h
import data as D

LANE, GROUP, AXIS, RAGE = 24, 30, 30, 72
DMG = "#e8c96a"     # --color-school-physical: a damaging cast
CD = "#e9e4d8"      # --color-text: a cast that deals no damage
CRIT = "#fbe7a1"    # --gold-light
MUTED = "#9a9484"
BUFF = "rgba(233,228,216,.35)"
DEBUFF = "rgba(168,118,42,.78)"
EXEC = "rgba(214,110,40,.07)"


def fmt(ms):
    s = ms / 1000
    sign = "-" if s < 0 else ""
    s = abs(s)
    return f"{sign}{int(s // 60)}:{s % 60:04.1f}"


def clock(ms):
    s = round(ms / 1000)
    sign = "-" if s < 0 else ""
    return f"{sign}{abs(s) // 60}:{abs(s) % 60:02d}"


CAST_KEYS = []
for x in D.SAMPLE:
    k = x["action"]
    if k.startswith("other:") or k == "spell:12867" or k == "spell:11567/1":
        continue
    if k not in CAST_KEYS:
        CAST_KEYS.append(k)
DAMAGING = {r["key"] for r in D.ability_rows()}
SWING_KEYS = ["other:attack/1", "other:attack/2", "other:attack/3", "spell:12867"]
BUFF_KEYS = [a["action"] for a in D.SAMPLE_AURAS if a["on"] == "player"]
DEBUFF_KEYS = [a["action"] for a in D.SAMPLE_AURAS if a["on"] == "target"]
EXEC_AT = D.sample_times("spell:20662")[0]


def lane_name(kind, k):
    if kind == "swing" and k == "spell:12867":
        return "Deep Wounds ticks"
    return D.name_of(k)


def lanes(swings):
    out = [("group", "Casts", len(CAST_KEYS))] + [("cast", k) for k in CAST_KEYS]
    if swings:
        out += [("group", "Swings and ticks", len(SWING_KEYS))] + [("swing", k) for k in SWING_KEYS]
    out += [("group", "On you", len(BUFF_KEYS))] + [("buff", k) for k in BUFF_KEYS]
    out += [("group", "On the target", len(DEBUFF_KEYS))] + [("debuff", k) for k in DEBUFF_KEYS]
    out += [("rage", None)]
    return out


def height_of(row):
    return {"group": GROUP, "rage": RAGE}.get(row[0], LANE)


def icon_img(k, name):
    ic = D.icon_of(k)
    if not ic:
        return f'<span class="ic" style="width:18px;height:18px;background:#14171f" aria-label="{h.escape(name)}"></span>'
    return f'<img class="ic" src="{D.icon_path(ic).as_uri()}" alt="{h.escape(name)}">'


def labels(swings):
    parts = ['<div class="tl-axis"><span>Time</span></div>']
    for row in lanes(swings):
        if row[0] == "group":
            parts.append(f'<div class="tl-grp"><span>{h.escape(row[1])}</span></div>')
        elif row[0] == "rage":
            peak = max(v for _, v in D.rage_points())
            parts.append(f'<div class="tl-rage"><b>Rage</b><span>peak {peak} of 100</span></div>')
        else:
            nm = lane_name(row[0], row[1])
            parts.append(f'<div class="tl-row" title="{h.escape(nm)}">{icon_img(row[1], nm)}<span>{h.escape(nm)}</span></div>')
    return "".join(parts)


def tick(x, y, hh, ev, kind, col):
    """A cast tick, shaped by its outcome when the run records one (spec 3.7)."""
    outcome = ev.get("outcome")
    if kind == "swing":
        w, th = 2, 10
    else:
        w, th = 3, 14
    if kind == "swing" and ev["action"] == "spell:12867":
        outcome = None                       # periodic ticks draw plain at every outcome
    if outcome == "crit":
        w, th, col = w + 1, th + 4, CRIT
    if outcome in ("miss", "dodge", "parry"):
        return f'<rect x="{x - w / 2:.1f}" y="{y + (hh - th) / 2:.1f}" width="{w}" height="{th}" rx="1" fill="none" stroke="{MUTED}" stroke-width="1"/>'
    op = ".6" if outcome == "glance" else (".8" if kind == "swing" else "1")
    return f'<rect x="{x - w / 2:.1f}" y="{y + (hh - th) / 2:.1f}" width="{w}" height="{th}" rx="1" fill="{col}" opacity="{op}"/>'


def plot(width, t0, t1, swings, step, crosshair=None, close=False, inlabels=False):
    span = t1 - t0
    X = lambda t: (t - t0) / span * width
    rows = lanes(swings)
    total_h = AXIS + sum(height_of(r) for r in rows)
    s = [f'<svg width="{width}" height="{total_h}" viewBox="0 0 {width} {total_h}" xmlns="http://www.w3.org/2000/svg" font-family="Barlow" font-size="13">']
    ex = max(EXEC_AT, t0)
    if ex < t1:
        s.append(f'<rect x="{X(ex):.1f}" y="{AXIS}" width="{width - X(ex):.1f}" height="{total_h - AXIS}" fill="{EXEC}"/>')
        s.append(f'<text x="{X(ex) + 6:.1f}" y="{AXIS + 14}" fill="#d66e28" font-weight="700" font-size="11" letter-spacing="1.4">EXECUTE PHASE</text>')
    t = (t0 // step + 1) * step
    while t < t1:
        x = X(t)
        s.append(f'<line x1="{x:.1f}" x2="{x:.1f}" y1="{AXIS - 6}" y2="{total_h}" stroke="#1c2230"/>')
        if t != 0:
            s.append(f'<text x="{x:.1f}" y="{AXIS - 10}" fill="{MUTED}" text-anchor="middle" font-weight="600">{clock(t)}</text>')
        t += step
    if t0 < 0 < t1:
        s.append(f'<line x1="{X(0):.1f}" x2="{X(0):.1f}" y1="{AXIS - 6}" y2="{total_h}" stroke="#4a4030" stroke-dasharray="3 3"/>')
        s.append(f'<text x="{X(0) + 4:.1f}" y="{AXIS - 10}" fill="#e5b955" font-weight="600">Pull</text>')
    y = AXIS
    for row in rows:
        hh = height_of(row)
        kind = row[0]
        if kind in ("cast", "swing"):
            k = row[1]
            if k == "spell:11567" and close:
                for tt in D.sample_times("spell:11567/1"):
                    if t0 <= tt <= t1:
                        s.append(f'<rect x="{X(tt) - 1.5:.1f}" y="{y + 5}" width="3" height="14" rx="1" fill="none" stroke="{DMG}" stroke-width="1"/>')
            col = DMG if k in DAMAGING else CD
            for ev in D.sample_events(k):
                if t0 <= ev["at_ms"] <= t1:
                    s.append(tick(X(ev["at_ms"]), y, hh, ev, kind, col))
        elif kind in ("buff", "debuff"):
            col = BUFF if kind == "buff" else DEBUFF
            for a, b in D.segments(row[1], "player" if kind == "buff" else "target"):
                a, b = max(a, t0), min(b, D.SAMPLE_END, t1)
                if b > a:
                    s.append(f'<rect x="{X(a):.1f}" y="{y + 7}" width="{max(2, X(b) - X(a)):.1f}" height="10" rx="3" fill="{col}"/>')
        elif kind == "rage":
            pts = D.rage_points()
            top, bot = y + 8, y + hh - 8
            Y = lambda v: bot - v / 100 * (bot - top)
            s.append(f'<line x1="0" x2="{width}" y1="{Y(100):.1f}" y2="{Y(100):.1f}" stroke="#3a3326" stroke-dasharray="2 4"/>')
            path, prev = [], None
            for tt, v in pts:
                if tt < t0 or tt > t1:
                    prev = v if tt < t0 else prev
                    continue
                if not path:
                    path.append(f"M{X(max(t0, tt)):.1f},{Y(prev if prev is not None else v):.1f}")
                path.append(f"H{X(tt):.1f}V{Y(v):.1f}")
                prev = v
            path.append(f"H{X(min(t1, D.SAMPLE_END)):.1f}")
            d = "".join(path)
            s.append(f'<path d="{d}V{bot}H{X(max(t0, pts[0][0])):.1f}Z" fill="rgba(229,185,85,.12)"/>')
            s.append(f'<path d="{d}" fill="none" stroke="#e5b955" stroke-width="1.5"/>')
        if inlabels and kind in ("cast", "swing", "buff", "debuff"):
            nm = lane_name(kind, row[1])
            s.append(f'<rect x="2" y="{y + 5}" width="{len(nm) * 5.6 + 8:.0f}" height="14" rx="3" fill="#0d0f15" opacity=".88"/>'
                     f'<text x="6" y="{y + 16}" fill="{MUTED}" font-size="11">{h.escape(nm)}</text>')
        y += hh
    if crosshair is not None:
        x = X(crosshair)
        s.append(f'<line x1="{x:.1f}" x2="{x:.1f}" y1="{AXIS - 4}" y2="{total_h}" stroke="#f5d27a" stroke-width="1"/>')
        s.append(f'<rect x="{x - 26:.1f}" y="2" width="52" height="20" rx="4" fill="#f5d27a"/><text x="{x:.1f}" y="16" text-anchor="middle" fill="#1a1206" font-weight="700">{clock(crosshair)}</text>')
    s.append("</svg>")
    return "".join(s), total_h


def brush(width, t0, t1, w0, w1):
    span = t1 - t0
    X = lambda t: (t - t0) / span * width
    d = "M" + " ".join(f"{X(t):.1f},{36 - v / 100 * 24:.1f}" for t, v in D.rage_points() if t0 <= t <= t1)
    s = [f'<svg width="{width}" height="44" xmlns="http://www.w3.org/2000/svg">', f'<path d="{d}" fill="none" stroke="#5d5849" stroke-width="1"/>']
    if w0 is not None:
        s.append(f'<rect x="{X(w0):.1f}" y="4" width="{X(w1) - X(w0):.1f}" height="36" rx="4" fill="rgba(229,185,85,.10)" stroke="#e5b955"/>')
        for xx in (X(w0), X(w1)):
            s.append(f'<rect x="{xx - 3:.1f}" y="12" width="6" height="20" rx="2" fill="#e5b955"/>')
    s.append("</svg>")
    return "".join(s)
