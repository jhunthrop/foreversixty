"""Shared atoms for the Wave 3 boards (design/specs/2026-10-31-mobile-wave3.md). It reuses the v1 kit in
../2026-10-30-mobile-apps and adds the fifth tab (Guides), a few glyphs, and the small helpers every new screen needs.
Run gen_wave3.py, not this file."""
import base64
import json
import sys
from pathlib import Path

HERE = Path(__file__).resolve().parent
V1 = HERE.parent / "2026-10-30-mobile-apps"
sys.path.insert(0, str(V1))
import kit  # noqa: E402
from kit import (CLASS_COLOR, DEVICES, Q_BORDER, Q_TEXT, Q_WORD, content_top, crest, device, glyph, icon_btn, offline_strip, selector, sheet,  # noqa: E402,F401
                 spec_icon, top_bar, back_btn, item, item_icon, qname)

W3 = json.loads((HERE / "w3.json").read_text())
HORDE = kit.HORDE

# The fifth tab. The v1 tab bar is four wide; the Wave 3 boards draw five (spec section 2).
kit.TABS[:] = [("house", "Home"), ("calendar", "Raids"), ("shield", "BiS"), ("book", "Guides"), ("person", "You")]
kit.EXTRA_GLYPHS.update({
    "book": "<path d=\"M5 4h10a3 3 0 0 1 3 3v13H8a3 3 0 0 1-3-3z\"/><path d=\"M5 17a3 3 0 0 1 3-3h10\"/>",
    "trophy": "<path d=\"M8 4h8v5a4 4 0 0 1-8 0zM8 6H5v2a3 3 0 0 0 3 3M16 6h3v2a3 3 0 0 1-3 3M12 13v4M8 20h8M9.5 17h5\"/>",
    "star-o": "<path d=\"m12 3.5 2.6 5.4 5.9.8-4.3 4.1 1 5.9L12 17l-5.2 2.7 1-5.9L3.5 9.7l5.9-.8z\"/>",
    "star-f": "<path fill=\"currentColor\" d=\"m12 3.5 2.6 5.4 5.9.8-4.3 4.1 1 5.9L12 17l-5.2 2.7 1-5.9L3.5 9.7l5.9-.8z\"/>",
    "undo": "<path d=\"M9 7 4 12l5 5M4 12h10a6 6 0 0 1 0 12\" transform=\"translate(0 -4)\"/>",
    "cloud": "<path d=\"M7 18a4 4 0 0 1-.5-8A6 6 0 0 1 18 9a4.5 4.5 0 0 1 0 9z\"/>",
    "cloud-down": "<path d=\"M7 17a4 4 0 0 1-.5-8A6 6 0 0 1 18 8a4.5 4.5 0 0 1 0 9M12 11v6m-3-3 3 3 3-3\"/>",
    "minus": "<path d=\"M5 12h14\"/>",
    "plus": "<path d=\"M12 5v14M5 12h14\"/>",
    "mic": "<rect x=\"9\" y=\"3\" width=\"6\" height=\"11\" rx=\"3\"/><path d=\"M5 11a7 7 0 0 0 14 0M12 18v3\"/>",
    "users": "<circle cx=\"9\" cy=\"8\" r=\"3.5\"/><path d=\"M2.5 20c0-3.5 3-6 6.5-6s6.5 2.5 6.5 6M16 5a3.5 3.5 0 0 1 0 7M18 14c2.4.6 3.5 2.6 3.5 6\"/>",
    "eye": "<path d=\"M2 12s3.6-7 10-7 10 7 10 7-3.6 7-10 7S2 12 2 12z\"/><circle cx=\"12\" cy=\"12\" r=\"3\"/>",
    "swords": "<path d=\"m4 4 9 9M20 4l-9 9M4 20l3-3M20 20l-3-3M6.5 14.5l3 3M14.5 17.5l3-3\"/>",
    "grid": "<rect x=\"4\" y=\"4\" width=\"7\" height=\"7\" rx=\"1.5\"/><rect x=\"13\" y=\"4\" width=\"7\" height=\"7\" rx=\"1.5\"/><rect x=\"4\" y=\"13\" width=\"7\" height=\"7\" rx=\"1.5\"/><rect x=\"13\" y=\"13\" width=\"7\" height=\"7\" rx=\"1.5\"/>",
    "link": "<path d=\"M10 14a4 4 0 0 0 5.6 0l3-3a4 4 0 0 0-5.6-5.6l-1 1M14 10a4 4 0 0 0-5.6 0l-3 3A4 4 0 0 0 11 18.6l1-1\"/>",
    "sparkle": "<path d=\"M12 3v4M12 17v4M3 12h4M17 12h4M6 6l2.5 2.5M15.5 15.5 18 18M18 6l-2.5 2.5M8.5 15.5 6 18\"/>",
    "siri": "<circle cx=\"12\" cy=\"12\" r=\"8\"/><path d=\"M8 12c1.5-3 2.5-3 4 0s2.5 3 4 0\"/>",
})

EXTRA_CSS = """
.ios-tab{grid-template-columns:repeat(5,1fr)!important;left:14px!important;right:14px!important}
.ios-tab .it{font-size:9.5px!important;letter-spacing:-.01em}
.and-nav{grid-template-columns:repeat(5,1fr)!important}
.and-nav .ind{width:56px!important}
.hscroll{display:flex;gap:8px;overflow:hidden;padding:0 16px}
.numchip{display:inline-flex;align-items:center;gap:8px;min-height:36px;padding:0 12px 0 6px;border-radius:999px;border:1px solid var(--line);background:rgba(255,255,255,.03);white-space:nowrap}
.numchip .n{width:22px;height:22px;border-radius:11px;display:grid;place-items:center;background:var(--tile);color:var(--tile-ink);font:700 12px "JetBrains Mono",monospace}
.meter{position:relative;display:flex;align-items:center;gap:10px;padding:6px 16px;min-height:44px}
.meter .fill{position:absolute;left:0;top:3px;bottom:3px;border-radius:0 4px 4px 0;opacity:.22}
.meter>*:not(.fill){position:relative}
.dot{display:inline-block;width:10px;height:10px;border-radius:5px;flex:none}
.tierchip{display:inline-flex;align-items:center;justify-content:center;width:28px;height:28px;border-radius:6px;font:800 16px "Barlow Condensed",sans-serif;border:1.5px solid currentColor;background:rgba(255,255,255,.03)}
.glass-w{background:rgba(34,38,48,.78);backdrop-filter:blur(20px);border:1px solid rgba(255,255,255,.10);box-shadow:inset 0 1px 0 rgba(255,255,255,.08)}
.siri-glow{box-shadow:0 0 0 2px rgba(120,160,255,.35),0 0 40px rgba(150,110,255,.35)}
.postlaunch{display:inline-flex;align-items:center;height:20px;padding:0 7px;border-radius:5px;font-size:10px;font-weight:700;letter-spacing:.08em;text-transform:uppercase;color:#f5d27a;background:rgba(229,185,85,.12);border:1px dashed rgba(229,185,85,.5)}
"""


def board(title: str, label: str, cells: list[str]) -> str:
    html = kit.board(title, label, cells)
    html = html.replace("Forever Sixty apps · 2026-10-30", "Forever Sixty apps · Wave 3 · 2026-10-31")
    return html.replace("</style>", EXTRA_CSS + "</style>", 1)


def cell(cap: str, dev: str) -> str:
    return kit.cell(cap, dev)


def icon_uri(stem: str) -> str:
    p = HERE / "assets" / f"{stem}.webp"
    if not p.exists():
        p = V1 / "assets" / f"{stem}.webp"
    if not p.exists():
        p = HERE / "assets" / "inv_misc_questionmark.webp"
    return "data:image/webp;base64," + base64.b64encode(p.read_bytes()).decode()


def sp(stem: str, size: int = 36, border: str = "#3a4256") -> str:
    """A spell, talent or spec icon: square, the pill radius, a 1 pt border."""
    return f"<img src=\"{icon_uri(stem)}\" alt=\"\" style=\"flex:none;width:{size}px;height:{size}px;border-radius:6px;border:1px solid {border};object-fit:cover\">"


def spec_mark(spec_slug: str, size: int = 36) -> str:
    """The spec's talent-tab icon; a spec whose icon is not in the mock's asset set draws its class crest, never a letter tile."""
    from model import SPEC
    cls = spec_slug.split("-")[0]
    if spec_slug in SPEC and ((HERE / "assets" / f"{SPEC[spec_slug][3]}.webp").exists() or (V1 / "assets" / f"{SPEC[spec_slug][3]}.webp").exists()):
        return sp(SPEC[spec_slug][3], size)
    return crest(cls, size)


def tile(g: str, size: int = 36) -> str:
    return f"<span style=\"flex:none;width:{size}px;height:{size}px;border-radius:8px;display:grid;place-items:center;background:var(--tile);color:var(--tile-ink)\">{glyph(g, size // 2)}</span>"


def tag(text: str, tone: str = "gold") -> str:
    return f"<span class=\"tag tag-{tone}\">{text}</span>"


def label(text: str, first: bool = False, extra: str = "") -> str:
    return f"<div class=\"t-micro micro pad\" style=\"margin-top:{12 if first else 22}px;{extra}\">{text}</div>"


def card(inner: str, quiet: bool = True, margin: str = "8px 16px 0", pad: str = "") -> str:
    return f"<div class=\"card{' quiet' if quiet else ''}\" style=\"margin:{margin};{('padding:' + pad + ';') if pad else ''}\">{inner}</div>"


def row(lead: str, title: str, meta: str = "", trail: str = "", chev: bool = True, pressed: bool = False, minh: int = 56, align: str = "center") -> str:
    ch = f"<span class=\"chev\">{glyph('chev-right', 18)}</span>" if chev else ""
    m = f"<div class=\"t-meta muted\">{meta}</div>" if meta else ""
    return (f"<div class=\"lrow{' pressed' if pressed else ''}\" style=\"min-height:{minh}px;align-items:{align}\">{lead}<div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:600\">{title}</div>{m}</div>{trail}{ch}</div>")


def seg_n(labels: list[str], active: int, margin: str = "12px 16px 0") -> str:
    cols = len(labels)
    parts = "".join(f"<span class=\"t-seg\" style=\"{'background:rgba(229,185,85,.12);color:var(--strong);box-shadow:inset 0 -2px 0 var(--gold);' if i == active else ''}\">{t}</span>" for i, t in enumerate(labels))
    return f"<div class=\"seg4\" style=\"grid-template-columns:repeat({cols},1fr);margin:{margin}\">{parts}</div>"


def chip(inner: str, on: bool = False, chev: bool = True) -> str:
    st = "border-color:var(--gold);background:rgba(229,185,85,.10);" if on else ""
    c = f"<span class=\"chev\">{glyph('chev-down', 14)}</span>" if chev else ""
    return f"<span style=\"display:inline-flex;align-items:center;gap:6px;min-height:40px;padding:0 12px;border-radius:999px;border:1px solid var(--card-b);background:rgba(255,255,255,.04);{st}white-space:nowrap\" class=\"t-meta\">{inner}{c}</span>"


PARSE = [(99, "#e268a8", "pink"), (95, "#ff8000", "orange"), (75, "#a335ee", "purple"), (50, "#0070ff", "blue"), (25, "#1eff00", "green"), (0, "#8c8c8c", "grey")]
PARSE_TEXT = {"pink": "#f08cc0", "orange": "#ff8000", "purple": "#b866f5", "blue": "#3d94f0", "green": "#1eff00", "grey": "#8c8c8c"}


def parse_color(p: float) -> str:
    return next(PARSE_TEXT[name] for floor, _, name in PARSE if p >= floor)


def pct(p: int) -> str:
    return f"<span class=\"mono\" style=\"font-weight:700;color:{parse_color(p)}\">{p}</span>"


def search_btn(plat: str) -> str:
    return icon_btn(plat, "search", "Search")


def root_bar(plat: str) -> str:
    return top_bar(plat, selector(plat), search_btn(plat))


def pushed(plat: str, back: str, title: str, inner: str, tab: str, overlay: str = "", right: str = "", xl: bool = False, height: int | None = None) -> str:
    top = content_top(plat)
    left = back_btn(plat, back) + (f"<span class=\"cond\" style=\"font-size:22px\">{title}</span>" if plat == "and" else "")
    head_ = "" if plat == "and" else f"<div class=\"cond t-title pad\" style=\"margin-top:4px\">{title}</div>"
    bar = top_bar(plat, f"<span style=\"display:flex;align-items:center;gap:4px\">{left}</span>", right)
    return device(plat, bar + f"<div style=\"padding-top:{top}px\">{head_}{inner}</div>", tab, overlay=overlay, xl=xl, height=height)


def root(plat: str, title: str, inner: str, tab: str, overlay: str = "", strip: str = "") -> str:
    top = content_top(plat)
    return device(plat, root_bar(plat) + f"<div style=\"padding-top:{top}px\">{strip}<div class=\"cond t-title pad\" style=\"margin-top:4px\">{title}</div>{inner}</div>", tab, overlay=overlay)


def skel(w: str, h: int, m: int = 8) -> str:
    return f"<div class=\"skel\" style=\"width:{w};height:{h}px;margin-top:{m}px\"></div>"


def fail(msg: str) -> str:
    return (f"<div style=\"padding:16px\"><div class=\"t-body wipe\" role=\"alert\">{msg}</div>"
            f"<span class=\"btn btn-secondary\" style=\"margin-top:12px\">{glyph('refresh', 16)}Try again</span></div>")


def note(text: str, extra: str = "") -> str:
    return f"<div class=\"t-meta muted pad\" style=\"margin-top:10px;{extra}\">{text}</div>"


TAB_CLEAR = "<div style=\"height:130px\"></div>"
CLASSES = ["warrior", "paladin", "hunter", "rogue", "priest", "shaman", "mage", "warlock", "druid"]
