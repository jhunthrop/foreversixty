"""Device chrome and shared atoms for the 2026-10-30 mobile app boards (design/specs/2026-10-30-mobile-apps.md).
iPhone 17 is drawn at 402 x 874 pt, Pixel 9 at 412 x 915 dp. Every colour, radius and type size is a
web/src/styles/tokens.css value; the large-text boards swap the type roles for iOS AX5 and Android 200 %.
Run gen_mobile_apps.py, not this file."""
import base64
import json
import sys
from pathlib import Path

HERE = Path(__file__).resolve().parent
sys.path.insert(0, str(HERE.parent / "2026-10-28-polish"))
from parts import CLASS_COLOR, GLYPHS, SEAL, crest, uri  # noqa: E402

DATA = json.loads((HERE / "data.json").read_text())
HORDE = uri(HERE.parents[2] / "web/public/icons/hd/faction/horde-logo-64.webp")
ALLIANCE = uri(HERE.parents[2] / "web/public/icons/hd/faction/alliance-logo-64.webp")

DEVICES = {
    "ios": {"w": 402, "h": 874, "top": 62, "bottom": 34, "radius": 55, "label": "iPhone 17 · 402 x 874 pt"},
    "and": {"w": 412, "h": 915, "top": 40, "bottom": 24, "radius": 40, "label": "Pixel 9 · 412 x 915 dp"},
}

# The game's quality colours: borders use the game hue, text uses the AA `-text` variant (tokens.css).
Q_BORDER = {0: "#9d9d9d", 1: "#ffffff", 2: "#1eff00", 3: "#0070dd", 4: "#a335ee", 5: "#ff8000"}
Q_TEXT = {0: "#9d9d9d", 1: "#ffffff", 2: "#1eff00", 3: "#3d94f0", 4: "#b866f5", 5: "#ff8000"}
Q_WORD = ["Poor", "Common", "Uncommon", "Rare", "Epic", "Legendary"]

EXTRA_GLYPHS = {
    "house": "<path d=\"M3 11 12 4l9 7\"/><path d=\"M5 10v10h14V10\"/><path d=\"M10 20v-6h4v6\"/>",
    "calendar": "<rect x=\"3\" y=\"5\" width=\"18\" height=\"16\" rx=\"2\"/><path d=\"M3 10h18M8 3v4M16 3v4\"/>",
    "shield": "<path d=\"M12 3 4 6v6c0 4.5 3.4 8.2 8 9 4.6-.8 8-4.5 8-9V6z\"/>",
    "person": "<circle cx=\"12\" cy=\"8\" r=\"4\"/><path d=\"M4 21c0-4 3.6-7 8-7s8 3 8 7\"/>",
    "back": "<path d=\"m15 5-7 7 7 7\"/>",
    "arrow-left": "<path d=\"M19 12H5M11 18l-6-6 6-6\"/>",
    "more": "<circle cx=\"5\" cy=\"12\" r=\"1.4\"/><circle cx=\"12\" cy=\"12\" r=\"1.4\"/><circle cx=\"19\" cy=\"12\" r=\"1.4\"/>",
    "more-v": "<circle cx=\"12\" cy=\"5\" r=\"1.4\"/><circle cx=\"12\" cy=\"12\" r=\"1.4\"/><circle cx=\"12\" cy=\"19\" r=\"1.4\"/>",
    "wifi-off": "<path d=\"M2 8.5a15 15 0 0 1 5-3M22 8.5A15 15 0 0 0 12 5M5 12a10 10 0 0 1 4-2.3M19 12a10 10 0 0 0-2-1.4M8.5 15.5a5 5 0 0 1 7 0M12 19h.01M3 3l18 18\"/>",
    "clock": "<circle cx=\"12\" cy=\"12\" r=\"9\"/><path d=\"M12 7v5l3 2\"/>",
    "moon": "<path d=\"M20 14.5A8 8 0 1 1 9.5 4a6.5 6.5 0 0 0 10.5 10.5z\"/>",
    "cal-plus": "<rect x=\"3\" y=\"5\" width=\"18\" height=\"16\" rx=\"2\"/><path d=\"M3 10h18M8 3v4M16 3v4M12 13v5M9.5 15.5h5\"/>",
    "share": "<path d=\"M12 3v12M7 8l5-5 5 5M5 13v7h14v-7\"/>",
    "globe": "<circle cx=\"12\" cy=\"12\" r=\"9\"/><path d=\"M3 12h18M12 3c3 3.5 3 14.5 0 18M12 3c-3 3.5-3 14.5 0 18\"/>",
    "devices": "<rect x=\"7\" y=\"3\" width=\"10\" height=\"18\" rx=\"2\"/><path d=\"M11 18h2\"/>",
    "lock": "<rect x=\"5\" y=\"11\" width=\"14\" height=\"10\" rx=\"2\"/><path d=\"M8 11V8a4 4 0 0 1 8 0v3\"/>",
    "swap": "<path d=\"M7 4 3 8l4 4M3 8h14M17 20l4-4-4-4M21 16H7\"/>",
    "flask": "<path d=\"M9 3h6M10 3v6L4.5 19a1.5 1.5 0 0 0 1.3 2h12.4a1.5 1.5 0 0 0 1.3-2L14 9V3\"/><path d=\"M7 15h10\"/>",
    "tree": "<circle cx=\"6\" cy=\"6\" r=\"2.2\"/><circle cx=\"18\" cy=\"6\" r=\"2.2\"/><circle cx=\"12\" cy=\"18\" r=\"2.2\"/><path d=\"M6 8.2v2.3a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V8.2M12 12.5v3.3\"/>",
    "doc": "<path d=\"M6 3h9l4 4v14H6z\"/><path d=\"M14 3v5h5M9 13h7M9 17h5\"/>",
    "play": "<path d=\"M7 4v16l13-8z\"/>",
    "info": "<circle cx=\"12\" cy=\"12\" r=\"9\"/><path d=\"M12 11v6M12 7.5h.01\"/>",
}


def glyph(name: str, size: int = 18, sw: float = 2) -> str:
    body = EXTRA_GLYPHS.get(name) or GLYPHS[name]
    return (f"<svg width=\"{size}\" height=\"{size}\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"{sw}\" "
            f"stroke-linecap=\"round\" stroke-linejoin=\"round\" aria-hidden=\"true\">{body}</svg>")


def asset(stem: str) -> str:
    return "data:image/webp;base64," + base64.b64encode((HERE / "assets" / f"{stem}.webp").read_bytes()).decode()


def item(item_id: int) -> dict:
    return DATA["items"][str(item_id)]


def item_icon(item_id: int, size: int = 44) -> str:
    it = item(item_id)
    return (f"<img class=\"ii\" src=\"{asset(it['icon'])}\" alt=\"\" style=\"width:{size}px;height:{size}px;--q:{Q_BORDER[it['quality']]}\">")


def spec_icon(stem: str, size: int = 24) -> str:
    return f"<img class=\"si\" src=\"{asset(stem)}\" alt=\"\" style=\"width:{size}px;height:{size}px\">"


def qname(item_id: int, cls: str = "t-body") -> str:
    it = item(item_id)
    return f"<span class=\"{cls} qn\" style=\"color:{Q_TEXT[it['quality']]}\">{it['name']}</span>"


ROLES = {
    "default": {"title": 34, "h2": 26, "stat": 28, "sec": 18, "name": 17, "body": 15, "meta": 13, "micro": 11, "seg": 14},
    # iOS AX5 (Apple's Dynamic Type table: largeTitle 60, title1 58, headline 53, subheadline 49, footnote 44, caption2 40).
    "ios-xl": {"title": 60, "h2": 56, "stat": 58, "sec": 53, "name": 53, "body": 49, "meta": 44, "micro": 40, "seg": 44},
    # Android 14 font scale 2.0, non-linear: small sizes double, large sizes grow less.
    "and-xl": {"title": 50, "h2": 42, "stat": 44, "sec": 34, "name": 32, "body": 30, "meta": 26, "micro": 22, "seg": 28},
}


def role_css(scale: str) -> str:
    r = ROLES[scale]
    lh = lambda px: round(px * (1.3 if px < 30 else 1.12))  # noqa: E731
    return (f".t-title{{font-size:{r['title']}px;line-height:{lh(r['title'])}px}}.t-h2{{font-size:{r['h2']}px;line-height:{lh(r['h2'])}px}}"
            f".t-stat{{font-size:{r['stat']}px;line-height:{lh(r['stat'])}px}}.t-sec{{font-size:{r['sec']}px;line-height:{lh(r['sec'])}px}}"
            f".t-name{{font-size:{r['name']}px;line-height:{lh(r['name'])}px}}.t-body{{font-size:{r['body']}px;line-height:{round(r['body'] * 1.47)}px}}"
            f".t-meta{{font-size:{r['meta']}px;line-height:{round(r['meta'] * 1.38)}px}}.t-micro{{font-size:{r['micro']}px;line-height:{round(r['micro'] * 1.36)}px}}"
            f".t-seg{{font-size:{r['seg']}px}}")


CSS = """
*{box-sizing:border-box}html,body{margin:0;background:#000;color:#e9e4d8;font-family:"Barlow","Helvetica Neue",Arial,sans-serif;-webkit-font-smoothing:antialiased}
:root{--bg:#07090d;--raised:#0d111a;--card-top:#131824;--line:#262e40;--line-soft:#1c2230;--line-warm:#3a3326;--text:#e9e4d8;--strong:#f2eee4;--muted:#9a9484;--nav:#b9b3a4;
--gold:#e5b955;--gold-hover:#f5d27a;--gold-deep:#a8762a;--kill:#7fd48a;--wipe:#f0736b;--card:linear-gradient(180deg,rgba(229,185,85,.075),rgba(229,185,85,.012) 45%),#0d0f15;
--card-b:rgba(229,185,85,.17);--hl:rgba(245,210,122,.45);--tile:rgba(229,185,85,.12);--tile-ink:#f5d27a;--row:rgba(229,185,85,.06);
--band:radial-gradient(700px 300px at 72% -10%,rgba(214,110,40,.28),transparent 70%),linear-gradient(180deg,#14100b,#07090d)}
.board{display:inline-flex;flex-direction:column;background:#000;padding:0 0 36px}
.bar{display:flex;gap:16px;align-items:center;height:40px;padding:0 28px;border-bottom:1px solid #1c2230;font:12px "JetBrains Mono",monospace;color:#9a9484;white-space:nowrap}
.bar b{color:#e5b955}
.row-of{display:flex;gap:36px;padding:24px 28px 0;align-items:flex-start}
.cell{display:flex;flex-direction:column;gap:10px}
.cap{font:12px/1.4 "JetBrains Mono",monospace;color:#9a9484;max-width:412px;min-height:34px}
.cap b{color:#e9e4d8;font-weight:700}
.dev{position:relative;overflow:hidden;background:var(--bg);flex:none;box-shadow:0 0 0 9px #16181d,0 0 0 10px #2b2f38}
.scroll{position:absolute;left:0;right:0;overflow:hidden}
.mono{font-family:"JetBrains Mono",Menlo,monospace;font-variant-numeric:tabular-nums}
.cond{font-family:"Barlow Condensed","Barlow",sans-serif;font-weight:800;text-transform:uppercase;letter-spacing:.01em;color:var(--strong)}
.micro{font-weight:700;letter-spacing:.14em;text-transform:uppercase;color:var(--muted)}
.gold{color:var(--gold)}.muted{color:var(--muted)}.strong{color:var(--strong)}.kill{color:var(--kill)}.wipe{color:var(--wipe)}
.wordmark{font-family:"Cinzel",serif;font-weight:700;letter-spacing:.10em;text-transform:uppercase;color:var(--strong)}
.pad{padding:0 16px}
.card{background:var(--card);border:1px solid var(--card-b);border-radius:12px;box-shadow:inset 0 1px 0 var(--hl),0 16px 38px rgba(0,0,0,.55)}
.card.quiet{box-shadow:none;background:var(--raised);border-color:var(--line)}
.ii{flex:none;border-radius:6px;border:1.5px solid var(--q);object-fit:cover;box-shadow:inset 0 0 0 1px rgba(0,0,0,.6),0 0 10px color-mix(in srgb,var(--q) 22%,transparent)}
.si{flex:none;border-radius:6px;border:1px solid #3a4256;object-fit:cover}
.crest{flex:none;border-radius:999px;object-fit:cover;box-shadow:0 0 0 2px color-mix(in srgb,var(--c) 60%,transparent),0 0 14px color-mix(in srgb,var(--c) 25%,transparent);background:#0d111a}
.lrow{display:flex;align-items:center;gap:12px;min-height:56px;padding:10px 16px;border-top:1px solid var(--line-soft)}
.lrow:first-child{border-top:0}
.lrow .t{flex:1;min-width:0}
.chev{color:var(--muted);display:flex;flex:none}
.tag{display:inline-flex;align-items:center;gap:6px;height:22px;padding:0 8px;border-radius:6px;font-size:11px;font-weight:700;letter-spacing:.08em;text-transform:uppercase;white-space:nowrap;border:1px solid}
.tag-gold{color:var(--gold);background:rgba(229,185,85,.12);border-color:rgba(229,185,85,.38)}
.tag-green{color:var(--kill);background:rgba(127,212,138,.10);border-color:rgba(127,212,138,.32)}
.tag-muted{color:var(--muted);background:rgba(154,148,132,.10);border-color:rgba(154,148,132,.30)}
.tag-wipe{color:var(--wipe);background:rgba(240,115,107,.14);border-color:rgba(240,115,107,.35)}
.btn{display:inline-flex;align-items:center;justify-content:center;gap:8px;min-height:48px;padding:0 18px;border-radius:8px;font-size:13px;font-weight:700;letter-spacing:.07em;text-transform:uppercase;white-space:nowrap;border:1px solid transparent}
.btn-primary{background:linear-gradient(180deg,#f0cc6c,#c99a3a);color:#0b0e14;box-shadow:inset 0 0 0 1px #a8762a,0 6px 20px rgba(229,185,85,.22)}
.btn-secondary{background:rgba(255,255,255,.045);border-color:var(--card-b);color:var(--strong)}
.btn-ghost{color:var(--gold);padding:0 8px}
.btn.full{width:100%}
.bar8{position:relative;height:8px;border-radius:4px;background:var(--line-soft);overflow:hidden;flex:none}
.bar8 i{position:absolute;left:0;top:0;bottom:0;background:var(--gold);border-radius:4px}
.seg4{display:grid;grid-template-columns:repeat(4,1fr);border:1px solid var(--line);border-radius:8px;overflow:hidden;background:var(--raised)}
.seg4 span{display:flex;align-items:center;justify-content:center;min-height:48px;font-weight:600;color:var(--muted);border-left:1px solid var(--line)}
.seg4 span:first-child{border-left:0}
.seg4 span.on-go{background:rgba(127,212,138,.15);color:var(--strong);box-shadow:inset 0 -2px 0 var(--kill)}
.seg4 span.on-no{background:var(--line-warm);color:var(--strong);box-shadow:inset 0 -2px 0 var(--muted)}
.seg4 span.pend{outline:1.5px dashed var(--gold);outline-offset:-4px;color:var(--strong)}
.seg4 span.focus{outline:2px solid var(--gold);outline-offset:-3px}
.seg4.dis{opacity:.5}
.seg4.xl{grid-template-columns:1fr 1fr}.seg4.xl span{border-left:0;border-top:1px solid var(--line)}.seg4.xl span:nth-child(2n){border-left:1px solid var(--line)}.seg4.xl span:nth-child(-n+2){border-top:0}
.dateblk{flex:none;display:flex;flex-direction:column;align-items:center;justify-content:center;width:52px;height:56px;border:1px solid var(--line);border-radius:8px;background:rgba(255,255,255,.02)}
.dateblk .d{font:700 22px/1 "JetBrains Mono",monospace;color:var(--strong)}
.dateblk .m{font-size:10px;font-weight:700;letter-spacing:.14em;color:var(--muted)}
.skel{background:linear-gradient(90deg,#131824,#1a2030,#131824);border-radius:6px}
.offline{display:flex;align-items:center;gap:8px;min-height:36px;padding:6px 16px;background:#11151f;border-bottom:1px solid var(--line);color:var(--nav)}
.sheet{position:absolute;left:0;right:0;bottom:0;background:#10141d;border-top:1px solid var(--line)}
.scrim{position:absolute;inset:0;background:rgba(0,0,0,.55)}
.grab{display:flex;justify-content:center;padding:8px 0 4px}.grab i{width:36px;height:5px;border-radius:3px;background:#4a5268}
.and .grab i{width:32px;height:4px;border-radius:2px}
.tip{background:#0a0d14;border:1px solid #3a4256;border-radius:8px;padding:12px 14px}
.tip .ln{color:#f2eee4}.tip .gr{color:#1eff00}
/* iOS chrome */
.ios-sb{position:absolute;top:0;left:0;right:0;height:62px;display:flex;align-items:center;justify-content:space-between;padding:6px 34px 0 46px;font:600 17px -apple-system,"SF Pro Text","Helvetica Neue",sans-serif;color:#fff;z-index:20}
.ios-di{position:absolute;top:11px;left:50%;transform:translateX(-50%);width:126px;height:37px;border-radius:20px;background:#000;z-index:21}
.glass{background:rgba(30,34,44,.62);backdrop-filter:blur(18px);border:1px solid rgba(255,255,255,.10);box-shadow:inset 0 1px 0 rgba(255,255,255,.08),0 8px 24px rgba(0,0,0,.4)}
.ios-tab{position:absolute;left:20px;right:20px;bottom:22px;height:64px;border-radius:32px;display:grid;grid-template-columns:repeat(4,1fr);align-items:center;padding:0 6px;z-index:15}
.ios-tab .it{display:flex;flex-direction:column;align-items:center;gap:2px;font:600 10px -apple-system,"SF Pro Text",sans-serif;color:#c9c4b8;height:54px;justify-content:center;border-radius:27px}
.ios-tab .it.on{color:var(--gold);background:rgba(229,185,85,.12)}
.ios-home{position:absolute;bottom:8px;left:50%;transform:translateX(-50%);width:140px;height:5px;border-radius:3px;background:#f2eee4;z-index:22}
.ios-btn{display:inline-flex;align-items:center;justify-content:center;height:44px;min-width:44px;border-radius:22px;color:var(--strong);padding:0 4px}
/* Android chrome */
.and-sb{position:absolute;top:0;left:0;right:0;height:40px;display:flex;align-items:center;justify-content:space-between;padding:0 22px 0 26px;font:500 14px "Roboto","Helvetica Neue",sans-serif;color:#e9e4d8;z-index:20}
.and-cam{position:absolute;top:13px;left:50%;transform:translateX(-50%);width:16px;height:16px;border-radius:50%;background:#000;box-shadow:0 0 0 2px #111;z-index:21}
.and-nav{position:absolute;left:0;right:0;bottom:0;height:104px;background:#0d111a;border-top:1px solid var(--line-soft);display:grid;grid-template-columns:repeat(4,1fr);padding:12px 8px 0;z-index:15}
.and-nav .it{display:flex;flex-direction:column;align-items:center;gap:4px;font:500 12px "Roboto",sans-serif;color:#c9c4b8}
.and-nav .ind{width:64px;height:32px;border-radius:16px;display:flex;align-items:center;justify-content:center}
.and-nav .it.on{color:var(--strong)}.and-nav .it.on .ind{background:rgba(229,185,85,.18);color:var(--gold)}
.and-gest{position:absolute;bottom:8px;left:50%;transform:translateX(-50%);width:108px;height:4px;border-radius:2px;background:#c9c4b8;z-index:22}
.and-bar{display:flex;align-items:center;gap:4px;height:64px;padding:0 4px 0 4px}
.and-icon{display:inline-flex;align-items:center;justify-content:center;width:48px;height:48px;border-radius:24px;color:var(--strong)}
.selector{display:inline-flex;align-items:center;gap:10px;min-height:44px;padding:4px 14px 4px 4px;border-radius:999px;max-width:100%}
.selector .nm{font-weight:700}.selector .ds{color:var(--muted)}
.focusring{outline:2px solid var(--gold);outline-offset:2px}
.pressed{background:rgba(229,185,85,.10)!important}
.toast{position:absolute;left:16px;right:16px;display:flex;align-items:center;gap:10px;padding:12px 14px;border-radius:12px;background:#1a1f2b;border:1px solid var(--line);color:var(--strong);z-index:16}
.sw{flex:none;position:relative;width:51px;height:31px;border-radius:16px;background:#2a3142}
.sw::after{content:"";position:absolute;top:2px;left:2px;width:27px;height:27px;border-radius:50%;background:#e9e4d8}
.sw.on{background:var(--gold)}.sw.on::after{left:22px;background:#fff}
.and .sw{width:52px;height:32px;border:2px solid #6b6658;background:#1a1f2b}.and .sw::after{top:6px;left:6px;width:16px;height:16px;background:#9a9484}
.and .sw.on{border-color:var(--gold);background:var(--gold)}.and .sw.on::after{top:2px;left:22px;width:24px;height:24px;background:#3a2a08}
.val{color:var(--muted);display:inline-flex;align-items:center;gap:4px;white-space:nowrap}
"""

FONTS = ("https://fonts.googleapis.com/css2?family=Barlow+Condensed:wght@700;800&family=Barlow:wght@400;500;600;700"
         "&family=JetBrains+Mono:wght@500;700&family=Cinzel:wght@700&family=Roboto:wght@500&display=swap")


def head(scale: str = "default") -> str:
    return (f"<!doctype html><html lang=\"en\"><head><meta charset=\"utf-8\"><title>Forever Sixty mobile mock</title>"
            f"<link rel=\"stylesheet\" href=\"{FONTS}\"><style>{CSS}{role_css('default')}"
            f".xl-ios{{}}</style></head><body>")


def scoped_roles(scale: str, scope: str) -> str:
    """The large-text type roles, scoped to one device so a board can hold both sizes."""
    css = role_css(scale)
    return "<style>" + css.replace(".t-", f".{scope} .t-") + "</style>"


def ios_status(dark_text: bool = False) -> str:
    sig = ("<svg width=\"18\" height=\"12\" viewBox=\"0 0 18 12\" fill=\"#fff\"><rect x=\"0\" y=\"8\" width=\"3\" height=\"4\" rx=\"1\"/><rect x=\"5\" y=\"5.5\" width=\"3\" height=\"6.5\" rx=\"1\"/>"
           "<rect x=\"10\" y=\"3\" width=\"3\" height=\"9\" rx=\"1\"/><rect x=\"15\" y=\"0\" width=\"3\" height=\"12\" rx=\"1\"/></svg>")
    wifi = "<svg width=\"16\" height=\"12\" viewBox=\"0 0 16 12\" fill=\"#fff\"><path d=\"M8 12 5.6 9.4a3.4 3.4 0 0 1 4.8 0zM3.4 7.2a6.5 6.5 0 0 1 9.2 0l1.5-1.6a8.7 8.7 0 0 0-12.2 0zM0 3.9 1.5 5.4a11.4 11.4 0 0 1 13 0L16 3.9a13.6 13.6 0 0 0-16 0z\"/></svg>"
    bat = "<svg width=\"27\" height=\"13\" viewBox=\"0 0 27 13\"><rect x=\".5\" y=\".5\" width=\"23\" height=\"12\" rx=\"3.5\" fill=\"none\" stroke=\"#fff\" stroke-opacity=\".4\"/><rect x=\"2\" y=\"2\" width=\"17\" height=\"9\" rx=\"2\" fill=\"#fff\"/><rect x=\"24.5\" y=\"4.5\" width=\"1.5\" height=\"4\" rx=\".7\" fill=\"#fff\" fill-opacity=\".5\"/></svg>"
    return f"<div class=\"ios-sb\"><span>9:41</span><span style=\"display:flex;gap:6px;align-items:center\">{sig}{wifi}{bat}</span></div><div class=\"ios-di\"></div>"


def and_status() -> str:
    icons = ("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"#e9e4d8\"><path d=\"M8 13 1 5.5a10 10 0 0 1 14 0z\"/></svg>"
             "<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"#e9e4d8\"><path d=\"M14 2v12H2z\"/></svg>"
             "<svg width=\"10\" height=\"16\" viewBox=\"0 0 10 16\" fill=\"#e9e4d8\"><rect x=\"3\" y=\"0\" width=\"4\" height=\"2\" rx=\"1\"/><rect x=\"0\" y=\"2\" width=\"10\" height=\"14\" rx=\"2\"/></svg>")
    return f"<div class=\"and-sb\"><span>9:41</span><span style=\"display:flex;gap:6px;align-items:center\">{icons}</span></div><div class=\"and-cam\"></div>"


TABS = [("house", "Home"), ("calendar", "Raids"), ("shield", "BiS"), ("person", "You")]


def ios_tabbar(active: str) -> str:
    items = "".join(f"<div class=\"it{' on' if label == active else ''}\">{glyph(g, 24, 1.9)}<span>{label}</span></div>" for g, label in TABS)
    return f"<div class=\"ios-tab glass\">{items}</div>"


def and_navbar(active: str) -> str:
    items = "".join(f"<div class=\"it{' on' if label == active else ''}\"><span class=\"ind\">{glyph(g, 24, 1.9)}</span><span>{label}</span></div>" for g, label in TABS)
    return f"<div class=\"and-nav\">{items}</div>"


ME = {"name": "Lightningbolltz", "cls": "shaman", "desc": "Orc Elemental Shaman · Level 60", "short": "Elemental · 60",
      "guild": "Olympus XXVII", "battletag": "Hunthrop#1894"}
ALTS = [("Lightningbolltz", "shaman", "Orc Elemental Shaman · Level 60", "Olympus XXVII"),
        ("Vorgash", "warrior", "Orc Protection Warrior · Level 60", "Olympus XXVII"),
        ("Grimtusk", "warrior", "Orc Fury Warrior · Level 41", "Olympus XXVII"),
        ("Sablewhisper", "priest", "Undead Shadow Priest · Level 23", "No guild")]


def selector(plat: str, xl: bool = False, open_: bool = False, name: str = ME["name"], cls: str = ME["cls"], desc: str = ME["short"]) -> str:
    """The one way to switch (memory rule): crest, the full name in its class colour, the chevron. At large
    text the description line drops; the name never truncates and wraps rather than clipping."""
    ring = "border:1px solid var(--gold);box-shadow:0 0 26px rgba(229,185,85,.32);" if open_ else ""
    cls_glass = "glass" if plat == "ios" else ""
    bg = "" if plat == "ios" else "background:rgba(255,255,255,.04);border:1px solid var(--card-b);"
    # The bar caps the selector's type (iOS AX1 headline 28 pt, Android 1.3x = 22 sp) so the full name always fits.
    cap = ("font-size:28px;line-height:32px;" if plat == "ios" else "font-size:22px;line-height:28px;") if xl else ""
    ds = "" if xl else f"<span class=\"ds t-micro\" style=\"display:block;letter-spacing:.04em\">{desc}</span>"
    return (f"<span class=\"selector {cls_glass}\" style=\"{bg}{ring}\">{crest(cls, 34)}<span style=\"min-width:0\">"
            f"<span class=\"nm{'' if xl else ' t-name'}\" style=\"display:block;color:{CLASS_COLOR[cls]};{cap}\">{name}</span>{ds}</span>"
            f"<span class=\"chev\" style=\"color:var(--nav)\">{glyph('chev-down', 16)}</span></span>")


def top_bar(plat: str, left: str, right: str = "", height: int = 60) -> str:
    """The navigation bar under the status bar: iOS glass items floating over the content, Android a
    64 dp top app bar on --bg."""
    top = DEVICES[plat]["top"]
    if plat == "ios":
        return (f"<div style=\"position:absolute;top:0;left:0;right:0;height:{top}px;z-index:12;background:rgba(7,9,13,.96)\"></div>"
                f"<div style=\"position:absolute;top:{top}px;left:0;right:0;height:{height}px;display:flex;align-items:center;justify-content:space-between;"
                f"padding:0 16px;gap:10px;z-index:12;background:linear-gradient(180deg,rgba(7,9,13,.96),rgba(7,9,13,.6) 80%,transparent)\">{left}<span style=\"display:flex;gap:8px\">{right}</span></div>")
    return (f"<div class=\"and-bar\" style=\"position:absolute;top:{top}px;left:0;right:0;height:{max(height, 64)}px;justify-content:space-between;padding:0 4px 0 12px;"
            f"z-index:12;background:var(--bg)\">{left}<span style=\"display:flex\">{right}</span></div>")


def icon_btn(plat: str, g: str, label: str = "") -> str:
    if plat == "ios":
        return f"<span class=\"ios-btn glass\" aria-label=\"{label}\">{glyph(g, 20)}</span>"
    return f"<span class=\"and-icon\" aria-label=\"{label}\">{glyph(g, 22)}</span>"


def back_btn(plat: str, label: str) -> str:
    if plat == "ios":
        return f"<span class=\"ios-btn glass\" style=\"padding:0 14px 0 8px;gap:2px;font:600 15px 'Barlow'\">{glyph('back', 22)}{label}</span>"
    return f"<span class=\"and-icon\">{glyph('arrow-left', 24)}</span>"


def device(plat: str, body: str, tab: str | None = "Home", overlay: str = "", xl: bool = False, height: int | None = None) -> str:
    d = DEVICES[plat]
    h = height or d["h"]
    scope = f"xl-{plat}" if xl else ""
    roles = scoped_roles(f"{plat}-xl", scope) if xl else ""
    chrome = ios_status() if plat == "ios" else and_status()
    nav = ""
    if tab:
        nav = ios_tabbar(tab) if plat == "ios" else and_navbar(tab)
    home = "<div class=\"ios-home\"></div>" if plat == "ios" else "<div class=\"and-gest\"></div>"
    return (f"{roles}<div class=\"dev {plat} {scope}\" style=\"width:{d['w']}px;height:{h}px;border-radius:{d['radius']}px\">"
            f"<div class=\"scroll\" style=\"top:0;bottom:0\">{body}</div>{nav}{overlay}{chrome}{home}</div>")


def content_top(plat: str, bar: int = 60) -> int:
    return DEVICES[plat]["top"] + (bar if plat == "ios" else max(bar, 64))


def offline_strip(stamp: str = "Thu 9:41 pm") -> str:
    return (f"<div class=\"offline t-meta\">{glyph('wifi-off', 16)}<span>Offline. Showing what you had at {stamp}.</span></div>")


def sheet(plat: str, inner: str, height: int, title: str = "") -> str:
    rad = "38px 38px 0 0" if plat == "ios" else "28px 28px 0 0"
    head_ = f"<div class=\"cond t-sec\" style=\"padding:6px 20px 10px\">{title}</div>" if title else ""
    return (f"<div class=\"scrim\"></div><div class=\"sheet\" style=\"height:{height}px;border-radius:{rad};z-index:30\"><div class=\"grab\"><i></i></div>{head_}{inner}</div>")


def cell(cap: str, dev: str) -> str:
    return f"<div class=\"cell\"><div class=\"cap\">{cap}</div>{dev}</div>"


def board(title: str, label: str, cells: list[str]) -> str:
    return (head() + f"<div class=\"board\"><div class=\"bar\"><b>Forever Sixty apps · 2026-10-30</b><span>{title}</span><span>{label}</span></div>"
            f"<div class=\"row-of\">{''.join(cells)}</div></div></body></html>")
