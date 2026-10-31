"""F6 talents, read only (design/specs/2026-10-31-mobile-talents.md). One character, two trees: the guide build and
the character's own, compared on the device by band-compare.ts's rule.

The planner's rules (web/src/lib/planner): a tier opens at 5 points per tier in its tree; a cell's border is green
when it holds points, gold when it can take one, the line colour when locked (face at 40 %); a prerequisite line
runs edge to edge, gold once met, over a 5 px --bg halo. The count is diffAgainstBand's diffCount: your points
spent where the guide spends fewer; `Matches` only when no cell differs."""
from kit import DATA, asset, back_btn, content_top, device, glyph, offline_strip, top_bar

ICON, ROW, TIERS, PITCH_MAX, HANG = 48, 70, 7, 96, 6
TREE_H = (TIERS - 1) * ROW + ICON + HANG + 2
BAND = "60"
SPEC = "Elemental"
LEVEL = 60
BORDER = {"maxed": "#7fd48a", "filled": "#7fd48a", "open": "#262e40", "locked": "#262e40"}
PILL = {"maxed": "#7fd48a", "filled": "#e9e4d8", "open": "#9a9484", "locked": "#9a9484"}
OVER, UNDER = "#f0736b", "#e5b955"  # your rank above the guide (respec), below it (take points)
TREE_NAMES = [t["name"] for t in DATA["trees"]]
GUIDE_CODE = DATA["bands"]["raid"]["talents"]  # the planner's own default: raid when the band has it (presets.ts selectBand)

FOOT = ("<style>.tc{position:absolute;display:block}.tc .tile{display:block;width:48px;height:48px;border-radius:8px;background:#131824;border:2px solid var(--b);overflow:hidden}"
        ".tc .tile img{display:block;width:44px;height:44px;border-radius:6px;object-fit:cover}.tc .tile img.dim{opacity:.4}"
        ".tc .bd{position:absolute;right:-6px;bottom:-6px;padding:1px 5px;border-radius:6px;background:#07090d;border:1px solid var(--b);color:var(--p);white-space:nowrap;line-height:1.25}"
        ".tc .dash{position:absolute;inset:-5px;border-radius:12px;border:2px dashed var(--dc);pointer-events:none}"
        ".tc .chip{position:absolute;left:-8px;top:-8px;padding:0 4px;border-radius:6px;background:#07090d;border:1px solid var(--dc);color:var(--dc);font-weight:700;line-height:1.25;z-index:2}"
        ".dot{margin-left:5px;width:8px;height:8px;border-radius:4px;background:#f0736b;display:inline-block}"
        ".tc .selring{position:absolute;inset:-9px;border-radius:15px;border:2px solid #e5b955;pointer-events:none}"
        ".tag.t-micro{height:auto;min-height:22px;padding:2px 9px}.sk{background:linear-gradient(90deg,#131824,#1a2030,#131824);border-radius:6px}"
        ".segx{display:grid;border:1px solid var(--line);overflow:hidden;background:var(--raised)}.segx span{display:flex;align-items:center;justify-content:center;gap:6px;font-weight:600;color:var(--muted);text-align:center;padding:0 8px}"
        ".segx span.on{color:var(--strong);background:rgba(229,185,85,.13)}"
        ".tabx{display:flex}.tabx span{flex:1;display:flex;align-items:center;justify-content:center;gap:6px;color:var(--nav);font-weight:600;white-space:nowrap}"
        ".tabx span.on{color:var(--strong);font-weight:700}"
        ".tabs-stack span{display:flex;justify-content:space-between;align-items:center;gap:12px;padding:10px 14px;border-top:1px solid var(--line-soft);color:var(--nav);font-weight:600}"
        ".tabs-stack span:first-child{border-top:0}.tabs-stack span.on{color:var(--strong);font-weight:700;background:var(--tile);box-shadow:inset 3px 0 0 var(--gold)}"
        ".pin{display:none}.pin.show{display:block}.sbody{overflow:auto;scrollbar-width:none}.sbody::-webkit-scrollbar{display:none}"
        ".dashbox{border:1.5px dashed var(--dc);border-radius:8px;padding:10px 12px}</style>")

JS = ("<script>(()=>{const dev=document.currentScript.previousElementSibling;const go=()=>{const sc=dev.querySelector('.scroll');"
      "const sel=dev.querySelector('[data-keep]');if(sel){const dr=dev.getBoundingClientRect();const r=sel.getBoundingClientRect();const sh=dev.querySelector('.sheet');"
      "const limit=sh?sh.getBoundingClientRect().top-dr.top:dr.height;sc.scrollTop+=(r.bottom-dr.top)-(limit-28);}else if(dev.dataset.scroll==='tree'){const t=dev.querySelector('.tc');const dr=dev.getBoundingClientRect();sc.scrollTop+=t.getBoundingClientRect().top-dr.top-150;}else if(dev.dataset.scroll){sc.scrollTop=parseFloat(dev.dataset.scroll);}"
      "const dr2=dev.getBoundingClientRect();const fb=dev.querySelector('.flowbar');const pn=dev.querySelector('div.pin');"
      "if(fb&&pn&&fb.getBoundingClientRect().top-dr2.top<parseFloat(pn.style.top)){dev.querySelectorAll('div.pin').forEach(e=>e.classList.add('show'));fb.style.visibility='hidden';}"
      "else if(sc.scrollTop>0&&!fb){dev.querySelectorAll('div.pin').forEach(e=>e.classList.add('show'));}"
      "if(sc.scrollTop>60){dev.querySelectorAll('span.pin').forEach(e=>e.classList.add('show'));dev.querySelectorAll('.big').forEach(e=>e.style.visibility='hidden');}"
      "dev.querySelectorAll('[data-end]').forEach(e=>{e.scrollTop=e.scrollHeight});};"
      "document.fonts.ready.then(()=>requestAnimationFrame(go));})();</script>")


def ranks_of(code: str) -> list[list[int]]:
    out = []
    for i, tree in enumerate(code.split("-")):
        talents = DATA["trees"][i]["talents"]
        out.append([min(int(c, 36), talents[j]["max_rank"]) for j, c in enumerate(tree[:len(talents)])] + [0] * max(0, len(talents) - len(tree)))
    return out


GUIDE = ranks_of(GUIDE_CODE)


def yours_differs() -> list[list[int]]:
    y = [list(t) for t in GUIDE]
    y[0][3] += 2   # Reverberation 0 -> 2, where the guide spends none
    y[0][5] -= 2   # Call of Flame 3 -> 1; nothing depends on it
    y[0][7] -= 2   # Elemental Warding 3 -> 1; nothing depends on it
    y[2][1] -= 1   # Tidal Focus (Restoration, a closed tree) 5 -> 4; the tab carries a dot
    # legal: nothing depends on the three reduced talents, so 3 points are unspent
    return y


def diff(yours: list[list[int]], guide: list[list[int]]) -> tuple[int, int, set[tuple[int, int]]]:
    """band-compare.ts diffAgainstBand: diffCount = sum max(yours - guide, 0), missingCount = sum max(guide - yours, 0)."""
    d = m = 0
    cells: set[tuple[int, int]] = set()
    for t, (ys, gs) in enumerate(zip(yours, guide)):
        for i, (y, g) in enumerate(zip(ys, gs)):
            d += max(y - g, 0)
            m += max(g - y, 0)
            if y != g:
                cells.add((t, i))
    return d, m, cells


def state_of(rank: int, mx: int, open_: bool) -> str:
    return "maxed" if rank >= mx else "filled" if rank > 0 else "open" if open_ else "locked"


def pitch(w: int) -> int:
    return min(PITCH_MAX, (w - 32 - ICON) // 3)


def tree_html(w: int, ti: int, ranks: list[int], differs: dict[int, int], sel: int | None = None, xl: bool = False,
              placeholder: bool = False) -> str:
    talents = DATA["trees"][ti]["talents"]
    p = pitch(w)
    left = (w - (3 * p + ICON)) // 2
    pos = lambda t: (left + t["column"] * p, t["tier"] * ROW)  # noqa: E731
    by_id = {t["id"]: (i, t) for i, t in enumerate(talents)}
    spent = sum(ranks)
    h = ICON // 2
    if placeholder:
        pat = [(0, 1), (0, 2), (1, 0), (1, 1), (1, 2), (1, 3), (2, 0), (2, 1), (3, 0), (3, 1), (3, 3), (4, 0), (4, 1), (4, 2)]
        cells = "".join(f"<span class=\"sk\" style=\"position:absolute;left:{left + c * p}px;top:{t * ROW}px;width:48px;height:48px;border-radius:8px\"></span>" for t, c in pat)
        return f"<div style=\"position:relative;height:{TREE_H}px;margin-top:16px\">{cells}</div>"
    lines = ""
    for t in talents:
        if t["prereq_talent_id"] is None:
            continue
        j, pre = by_id[t["prereq_talent_id"]]
        (ax, ay), (bx, by) = pos(pre), pos(t)
        if pre["column"] == t["column"]:
            d = f"M {ax + h} {ay + ICON} V {by}"
        else:
            side = 1 if t["column"] > pre["column"] else -1
            d = f"M {ax + h} {ay + ICON} V {by + h} H {bx + h - side * h}" if pre["tier"] != t["tier"] else f"M {ax + h + side * h} {ay + h} H {bx + h - side * h}"
        met = ranks[j] >= (t["prereq_rank"] or 0)
        col = '#e5b955' if met else '#262e40'
        if pre["column"] == t["column"]:
            tip = f"{bx + h - 4},{by - 7} {bx + h + 4},{by - 7} {bx + h},{by}"
        else:
            ex = bx + h - side * h
            tip = f"{ex - side * 7},{by + h - 4} {ex - side * 7},{by + h + 4} {ex},{by + h}"
        lines += (f"<path d=\"{d}\" fill=\"none\" stroke=\"#07090d\" stroke-width=\"5\" stroke-linecap=\"round\"/>"
                  f"<path d=\"{d}\" fill=\"none\" stroke=\"{'#e5b955' if met else '#262e40'}\" stroke-width=\"2\" stroke-linecap=\"round\"/>")
        lines += f"<polygon points=\"{tip}\" fill=\"{col}\" stroke=\"#07090d\" stroke-width=\"1.5\" stroke-linejoin=\"round\"/>"
    cells = ""
    for i, t in enumerate(talents):
        rank = ranks[i]
        pre_ok = t["prereq_talent_id"] is None or ranks[by_id[t["prereq_talent_id"]][0]] >= (t["prereq_rank"] or 0)
        st = state_of(rank, t["max_rank"], spent >= t["tier"] * 5 and pre_ok)
        x, y = pos(t)
        dv = differs.get(i, 0)
        dc = OVER if dv > 0 else UNDER
        dash = f"<i class=\"dash\"></i><span class=\"chip mono\" style=\"font-size:{14 if xl else 11}px\">{'+' if dv > 0 else '−'}{abs(dv)}</span>" if dv else ""
        ring = "<i class=\"selring\"></i>" if sel == i else ""
        keep = " data-keep=\"1\"" if sel == i else ""
        cells += (f"<span class=\"tc\" style=\"left:{x}px;top:{y}px;--b:{BORDER[st]};--p:{PILL[st]};--dc:{dc}\"{keep}>{ring}{dash}<span class=\"tile\"><img class=\"{'dim' if st == 'locked' else ''}\" src=\"{asset(t['icon'])}\" alt=\"\"></span>"
                  f"<span class=\"bd mono\" style=\"font-size:{14 if xl else 11}px\">{rank}/{t['max_rank']}</span></span>")
    return (f"<div style=\"position:relative;height:{TREE_H}px;margin-top:16px\"><svg width=\"{w}\" height=\"{TREE_H}\" style=\"position:absolute;left:0;top:0\" aria-hidden=\"true\">{lines}</svg>{cells}</div>")


# ---------- the parts of the screen ----------

def tag(kind: str, text: str) -> str:
    return f"<span class=\"tag tag-{kind} t-micro\">{text}</span>"


def status(kind: str, word: str, sentence: str, xl: bool = False, extra: str = "") -> str:
    flow = "flex-direction:column;align-items:flex-start;gap:8px" if xl else "flex-wrap:wrap;align-items:center;gap:6px 10px"
    return (f"<div class=\"pad\" style=\"display:flex;{flow};margin-top:8px;min-height:28px\">{tag(kind, word)}<span class=\"t-body strong\" style=\"flex:{'none' if xl else '1 1 200px'}\">{sentence}</span></div>{extra}")


def n(x: object) -> str:
    return f"<span class=\"mono\">{x}</span>"


def seg2(plat: str, on: int, xl: bool = False) -> str:
    h = 48 if plat == "and" else 44
    rad = "24px" if plat == "and" else "8px"
    names = ["Guide build", "Your build"]
    cells = ""
    for i, name in enumerate(names):
        chk = glyph("check", 16) if plat == "and" and i == on else ""
        cells += f"<span class=\"t-seg{' on' if i == on else ''}\" style=\"min-height:{h}px;{'border-top:1px solid var(--line);' if xl and i else '' if xl else 'border-left:1px solid var(--line);' if i else ''}\"{' ' }>{chk}{name}</span>"
    cols = "1fr" if xl else "1fr 1fr"
    return f"<div class=\"segx\" style=\"grid-template-columns:{cols};border-radius:{rad};margin:12px 16px 0\">{cells}</div>"


def tabs3(plat: str, totals: list[int], active: int, stacked: bool = False, flagged: frozenset = frozenset()) -> str:
    if stacked:
        rows = "".join(f"<span class=\"t-seg{' on' if i == active else ''}\" style=\"min-height:56px\"><span style=\"display:inline;padding:0;border:0;background:none;box-shadow:none\">{nm}</span><span class=\"mono\" style=\"display:inline;padding:0;border:0;background:none;box-shadow:none\">{t}{' <i class=\"dot\"></i>' if i in flagged else ''}</span></span>"
                       for i, (nm, t) in enumerate(zip(TREE_NAMES, totals)))
        return f"<div class=\"tabs-stack\" style=\"margin:12px 16px 0;border:1px solid var(--line);border-radius:8px;overflow:hidden;background:var(--raised)\">{rows}</div>"
    bar = 3 if plat == "and" else 2
    cells = "".join(f"<span class=\"t-seg{' on' if i == active else ''}\" style=\"min-height:{48 if plat == 'and' else 44}px;{'box-shadow:inset 0 -%dpx 0 var(--gold)' % bar if i == active else ''}\">{nm} <b class=\"mono\" style=\"font-weight:700\">{t}</b>{'<i class=\"dot\"></i>' if i in flagged else ''}</span>"
                    for i, (nm, t) in enumerate(zip(TREE_NAMES, totals)))
    return f"<div class=\"tabx\" style=\"margin:8px 16px 0;border-bottom:1px solid var(--line)\">{cells}</div>"


def compare_bar(plat: str, on: int | None, totals: list[int], active: int, xl: bool = False, pinned: bool = False, flagged: frozenset = frozenset()) -> str:
    seg = seg2(plat, on, xl) if on is not None else ""
    top = "" if on is not None else "<div style=\"height:4px\"></div>"
    inner = f"{seg}{top}{tabs3(plat, totals, active, xl, flagged)}"
    if pinned:
        return f"<div class=\"pin\" style=\"position:absolute;left:0;right:0;top:{content_top(plat)}px;z-index:11;background:var(--bg);border-bottom:1px solid var(--line-soft);padding-bottom:8px\">{inner}</div>"
    return f"<div class=\"flowbar\">{inner}</div>"


LEGEND_DIFF = "Dashed: your rank differs from the guide. +n is points over it, −n points under it. Tap a talent for both ranks and its text."
LEGEND_TAP = "Tap a talent for its rank text."
FOOTER = "Change talents in the game; the addon sends them on your next logout."
PLAN = "Plan a build on foreversixty.gg/planner"


def footer_html() -> str:
    return (f"<div class=\"t-meta muted pad\" style=\"margin-top:16px\">{FOOTER}</div>"
            f"<div class=\"pad\" style=\"margin-top:6px;display:flex;align-items:center;gap:6px;min-height:48px\"><span class=\"gold t-body\" style=\"font-weight:600;overflow-wrap:anywhere;min-width:0\">{PLAN}</span><span class=\"gold\" style=\"display:flex\">{glyph('external', 16)}</span></div>")


def screen(plat: str, inner: str, overlay: str = "", xl: bool = False, scroll: float | str | None = None, pin: str = "", w: int | None = None,
           height_pad: int = 130) -> str:
    """The pushed screen. The bar, the pinned compare bar and any sheet are overlays (outside the scrolling body)."""
    top = content_top(plat)
    bar_left = back_btn(plat, "Home")
    inline = "<span class=\"cond\" style=\"position:absolute;top:%dpx;left:0;right:0;text-align:center;font-size:17px;z-index:13;line-height:44px\">Talents</span>" % (DEVICES_TOP[plat] + 8) if plat == "ios" else ""
    bar_title = "<span class=\"cond\" style=\"font-size:22px\">Talents</span>"
    bar = top_bar(plat, f"<span style=\"display:flex;align-items:center;gap:4px\">{bar_left}{bar_title if plat == 'and' else ''}</span>", "")
    big = ""
    body = f"{FOOT}<div style=\"padding-top:{top}px\">{inner.replace('{head}', big)}<div style=\"height:{height_pad}px\"></div></div>"
    mask = f"<div style=\"position:absolute;top:0;left:0;right:0;height:{top if plat == 'ios' else 40}px;background:var(--bg);z-index:11\"></div>"
    html = device(plat, body, "Home", overlay=mask + bar + inline + pin + overlay, xl=xl)
    sc = f" data-scroll=\"{scroll}\"" if scroll is not None else ""
    return html.replace("<div class=\"dev ", f"<div{sc} class=\"dev ", 1) + JS


DEVICES_TOP = {"ios": 62, "and": 40}
WIDTH = {"ios": 402, "and": 412}
XL_TREE_SCROLL = {"ios": "tree", "and": "tree"}


def screen_content(plat: str, state: str = "differs", tab: int = 0, on: int = 1, xl: bool = False, sel: int | None = None,
                   scroll: float | str | None = None, overlay: str = "", offline: bool = False) -> str:
    """One state of the screen. `state`: matches, differs, guide (the Guide build segment), noexport, nobuild, nospec, loading, error."""
    w = WIDTH[plat]
    strip = offline_strip("Thu 9:41 pm") if offline else ""
    pinned = ""
    if state in ("loading",):
        sk = lambda wd, ht=14: f"<span class=\"sk\" style=\"display:inline-block;width:{wd}px;height:{ht}px\"></span>"  # noqa: E731
        inner = (strip + "{head}" + f"<div class=\"pad\" style=\"display:flex;gap:10px;align-items:center;margin-top:8px;min-height:28px\">{sk(72, 22)}{sk(220)}</div>"
                 f"<div class=\"pad\" style=\"margin-top:12px\">{sk(370 if plat == 'ios' else 380, 44 if plat == 'ios' else 48)}</div>"
                 f"<div class=\"pad\" style=\"margin-top:8px;display:flex;gap:12px;justify-content:space-around\">{sk(100, 14)}{sk(100, 14)}{sk(100, 14)}</div>"
                 f"{tree_html(w, 0, [0] * 16, {}, placeholder=True)}"
                 "<div class=\"t-meta muted pad\" style=\"margin-top:16px\">Loading talents…</div>" + footer_html())
        return screen(plat, inner, overlay, xl)
    if state == "error":
        inner = (strip + "{head}" + "<div class=\"pad\" style=\"margin-top:10px\"><div class=\"t-body wipe\" role=\"alert\">Your talents did not load.</div>"
                 "<span class=\"btn btn-secondary\" style=\"margin-top:12px;min-height:48px\">Try again</span></div>" + footer_html())
        return screen(plat, inner, overlay, xl)
    yours = yours_differs() if state in ("differs", "guide") else [list(t) for t in GUIDE]
    has_export = state not in ("noexport",)
    has_guide = state not in ("nobuild", "nospec")
    shown = GUIDE if (state == "noexport" or (has_export and on == 0)) else yours
    if state == "nobuild" or state == "nospec":
        shown = yours
    totals = [sum(t) for t in shown]
    allowed = min(51, LEVEL - 9)
    spent = sum(sum(t) for t in yours)
    unspent = max(0, allowed - spent)
    d, m, cells = diff(yours, GUIDE) if has_guide and has_export else (0, 0, set())
    show_dash = has_guide and has_export and shown is yours
    dset = {i: yours[tab][i] - GUIDE[tab][i] for t, i in cells if t == tab} if show_dash else {}
    flagged = frozenset(t for t, _ in cells) if show_dash else frozenset()
    if state == "matches":
        st = status("green", "Matches", f"Your build is the {SPEC} {BAND} guide build.", xl)
    elif state in ("differs", "guide"):
        st = status("gold", "Differs", f"{n(d)} of your {n(spent)} points differ from the {BAND} guide build.", xl)
    elif state == "noexport":
        st = status("muted", "Guide build", "Not available yet: no talent export for this character. Open the addon once to send it.", xl)
    elif state == "nobuild":
        st = status("muted", "Your build", "No talent build published yet for this spec and level range.", xl)
    else:
        st = status("muted", "Your build", "We do not know this character's spec yet, so there is no guide build to compare. Open the addon once to send it.", xl)
    un = ""
    if has_export and unspent > 0:
        un = f"<div class=\"t-meta muted pad\" style=\"margin-top:4px\">{unspent} points unspent.</div>"
    if has_guide and has_export and d > 0 and m > 0:
        un += f"<div class=\"t-meta muted pad\" style=\"margin-top:2px\">The guide build takes {m} points you have not.</div>"
    seg_on = on if (has_export and has_guide) else None
    bar_html = compare_bar(plat, seg_on, totals, tab, xl, flagged=flagged)
    pinned = "" if xl else compare_bar(plat, seg_on, totals, tab, xl, pinned=True, flagged=flagged)
    legend = LEGEND_DIFF if show_dash and cells else LEGEND_TAP
    tree = tree_html(w, tab, shown[tab], dset, sel, xl)
    inner = (strip + "{head}" + st + un + bar_html + tree + f"<div class=\"t-meta muted pad\" style=\"margin-top:8px\">{legend}</div>" + footer_html())
    return screen(plat, inner, overlay, xl, scroll, pinned, height_pad=420 if sel is not None else 130)


# ---------- the rank sheet ----------

def requires(ti: int, t: dict, ranks: list[int]) -> str:
    """The game's red lines for an unmet requirement, from the drawn build (tier x 5 points in the tree; the prerequisite's rank)."""
    talents = DATA["trees"][ti]["talents"]
    out = []
    if sum(ranks) < 5 * t["tier"]:
        out.append(f"Requires {5 * t['tier']} points in {TREE_NAMES[ti]} Talents")
    if t["prereq_talent_id"] is not None:
        j, pre = next((k, x) for k, x in enumerate(talents) if x["id"] == t["prereq_talent_id"])
        need = t["prereq_rank"] or 0
        if ranks[j] < need:
            out.append(f"Requires {need} point{'' if need == 1 else 's'} in {pre['name']}")
    return "".join(f"<div class=\"t-body wipe\" style=\"margin-top:12px\">{x}</div>" for x in out)


def sheet_html(plat: str, ti: int, idx: int, rank: int, guide_rank: int | None, shown: list[int], xl: bool = False, end: bool = False) -> str:
    t = DATA["trees"][ti]["talents"][idx]
    mx = t["max_rank"]
    st = state_of(rank, mx, False)
    req = requires(ti, t, shown)
    dc = OVER if guide_rank is not None and rank > guide_rank else UNDER
    cur = f"<div class=\"t-body\" style=\"color:var(--text)\">{t['ranks'][rank - 1]}</div>" if rank > 0 else ""
    nxt = (f"<div class=\"t-micro micro\" style=\"margin-top:{12 if rank > 0 else 0}px\">Next rank</div><div class=\"t-body muted\" style=\"margin-top:2px\">{t['ranks'][rank]}</div>" if rank < mx else "")
    gl = (f"<div class=\"dashbox t-body strong\" style=\"margin-top:16px;font-weight:600;--dc:{dc}\">Guide build: rank <span class=\"mono\">{guide_rank}/{mx}</span></div>" if guide_rank is not None and guide_rank != rank else "")
    head_ = (f"<div style=\"display:flex;gap:14px;align-items:center;margin-bottom:14px\"><span style=\"flex:none;width:56px;height:56px;border-radius:8px;background:#131824;border:2px solid {BORDER[st]};overflow:hidden\">"
             f"<img src=\"{asset(t['icon'])}\" alt=\"\" style=\"width:52px;height:52px;border-radius:6px\"></span><div style=\"min-width:0\"><div class=\"t-name strong\" style=\"font-weight:700;overflow-wrap:anywhere\">{t['name']}</div>"
             f"<div class=\"t-meta muted mono\">Rank {rank}/{mx}</div></div></div>")
    inset = 34 if plat == "ios" else 24
    body = (f"<div class=\"sbody\" {'data-end=\"1\"' if end else ''} style=\"padding:0 20px;flex:1;min-height:0\"><div style=\"padding-bottom:{16 + inset}px\">{head_}<div class=\"tip\">{cur}{nxt}{req}</div>{gl}</div></div>")
    rad = "38px 38px 0 0" if plat == "ios" else "28px 28px 0 0"
    cap = DEVICES_TOP[plat] + (12 if plat == "ios" else 8)
    full = {"ios": 874, "and": 915}[plat]
    style = f"height:{full - cap}px;" if xl else ""
    return (f"<div class=\"scrim\" style=\"z-index:29\"></div><div class=\"sheet\" style=\"{style}border-radius:{rad};z-index:30;display:flex;flex-direction:column\"><div class=\"grab\"><i></i></div>"
            f"<div style=\"flex:1;min-height:0;display:flex;flex-direction:column\">{body}</div></div>")


def with_sheet(plat: str, state: str, tab: int, on: int, idx: int, rank: int, guide_rank: int | None, xl: bool = False, end: bool = False) -> str:
    return screen_content(plat, state, tab, on, xl, sel=None if xl else idx, overlay=sheet_html(plat, tab, idx, rank, guide_rank, (GUIDE if on == 0 else yours_differs())[tab], xl, end))


def tree_scrolled(plat: str, state: str, tab: int, on: int, xl: bool, scroll: float) -> str:
    return screen_content(plat, state, tab, on, xl, scroll=scroll)


def named(name: str) -> int:
    return next(i for i, t in enumerate(DATA["trees"][0]["talents"]) if t["name"] == name)


def vectors() -> str:
    y = yours_differs()
    d, m, c = diff(y, GUIDE)
    return f"differs: diffCount={d} missingCount={m} cells={len(c)} spent={sum(map(sum, y))} totals={[sum(t) for t in y]}"
