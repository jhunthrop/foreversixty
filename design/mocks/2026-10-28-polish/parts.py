"""Shared parts of the 2026-10-28 polish boards: real assets (crests, seal, the build icons
fetched from production into assets/, the Discord mark read from DiscordMark.astro), the new
header with its account menu, the new footer, and inline SVG glyphs. Every string is the
spec copy (design/specs/2026-10-28-polish.md); nothing is invented."""
import base64
import re
from pathlib import Path

HERE = Path(__file__).resolve().parent
ROOT = HERE.parents[2]
WEB = ROOT / "web"

CLASS_COLOR = {"warrior": "#c69b6d", "paladin": "#f48cba", "hunter": "#aad372", "rogue": "#fff468", "priest": "#ffffff",
               "shaman": "#3f8fe0", "mage": "#3fc7eb", "warlock": "#8788ee", "druid": "#ff7c0a"}
CLASSES = list(CLASS_COLOR)


def uri(path: Path) -> str:
    mime = {"webp": "image/webp", "svg": "image/svg+xml", "png": "image/png"}[path.suffix.lstrip(".")]
    return f"data:{mime};base64," + base64.b64encode(path.read_bytes()).decode()


def crest_src(slug: str) -> str:
    return uri(WEB / "public/icons/hd/crests" / f"{slug}.webp")


def crest(slug: str, size: int) -> str:
    return f"<img class=\"crest\" src=\"{crest_src(slug)}\" alt=\"\" style=\"width:{size}px;height:{size}px;--c:{CLASS_COLOR[slug]}\">"


def asset(stem: str) -> str:
    return uri(HERE / "assets" / f"{stem}.webp")


def build_icon(stem: str) -> str:
    return uri(WEB / "public/data/1.60.1.70658/icons" / f"{stem}.webp")


SEAL = uri(ROOT / "design/logo/foreversixty-mark.svg")
HORDE = uri(WEB / "public/icons/hd/faction/horde-logo-64.webp")
_DISCORD_PATH = re.search(r"<path d=\"([^\"]+)\"", (WEB / "src/components/DiscordMark.astro").read_text()).group(1)


def discord_mark(size: int = 18) -> str:
    return f"<svg width=\"{size}\" height=\"{size}\" viewBox=\"0 0 24 24\" fill=\"currentColor\" aria-hidden=\"true\"><path d=\"{_DISCORD_PATH}\"/></svg>"


GLYPHS = {
    "search": "<circle cx=\"11\" cy=\"11\" r=\"7\"/><path d=\"m20 20-3.5-3.5\"/>",
    "chev-down": "<path d=\"m6 9 6 6 6-6\"/>",
    "chev-right": "<path d=\"m9 6 6 6-6 6\"/>",
    "menu": "<path d=\"M4 7h16M4 12h16M4 17h16\"/>",
    "user": "<circle cx=\"12\" cy=\"8\" r=\"4\"/><path d=\"M4 21c0-4 3.6-7 8-7s8 3 8 7\"/>",
    "crown": "<path d=\"M3 8l4 4 5-7 5 7 4-4-2 11H5z\"/>",
    "sliders": "<path d=\"M4 6h10M18 6h2M4 12h4M12 12h8M4 18h12\"/><circle cx=\"16\" cy=\"6\" r=\"2\"/><circle cx=\"10\" cy=\"12\" r=\"2\"/><circle cx=\"18\" cy=\"18\" r=\"2\"/>",
    "logout": "<path d=\"M15 4h3a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2h-3M10 17l-5-5 5-5M5 12h11\"/>",
    "paste": "<rect x=\"6\" y=\"4\" width=\"12\" height=\"17\" rx=\"2\"/><path d=\"M9 4h6v3H9zM9 12h6M9 16h4\"/>",
    "server": "<rect x=\"3\" y=\"4\" width=\"18\" height=\"7\" rx=\"2\"/><rect x=\"3\" y=\"13\" width=\"18\" height=\"7\" rx=\"2\"/><path d=\"M7 7.5h.01M7 16.5h.01\"/>",
    "chart": "<path d=\"M4 19V5M4 19h16M8 15l4-5 3 3 5-6\"/>",
    "bell": "<path d=\"M6 16V11a6 6 0 1 1 12 0v5l2 2H4zM10 21h4\"/>",
    "star": "<path d=\"m12 3 2.7 5.6 6.1.9-4.4 4.3 1 6.1L12 17l-5.4 2.9 1-6.1L3.2 9.5l6.1-.9z\"/>",
    "external": "<path d=\"M14 4h6v6M20 4l-9 9M18 14v5a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1h5\"/>",
    "send": "<path d=\"M21 3 10 14M21 3l-7 18-4-7-7-4z\"/>",
    "refresh": "<path d=\"M20 11a8 8 0 1 0-2.3 5.7M20 4v7h-7\"/>",
    "download": "<path d=\"M12 4v11M7 11l5 5 5-5M5 20h14\"/>",
    "check": "<path d=\"M5 12.5l4.5 4.5L19 7\"/>",
    "card": "<rect x=\"3\" y=\"5\" width=\"18\" height=\"14\" rx=\"2\"/><path d=\"M3 10h18\"/>",
    "repo": "<path d=\"M9 19c-4 1.3-4-2-6-2.5M15 21v-3.5c0-1 .1-1.4-.5-2 2.8-.3 5.5-1.4 5.5-6a4.6 4.6 0 0 0-1.3-3.2 4.2 4.2 0 0 0-.1-3.2s-1-.3-3.4 1.3a11.7 11.7 0 0 0-6.2 0C6.6 2.8 5.6 3.1 5.6 3.1a4.2 4.2 0 0 0-.1 3.2A4.6 4.6 0 0 0 4.2 9.5c0 4.6 2.7 5.7 5.5 6-.6.6-.6 1.2-.5 2V21\"/>",
    "x": "<path d=\"M6 6l12 12M18 6 6 18\"/>",
    "hash": "<path d=\"M5 9h14M5 15h14M10 4 8 20M16 4l-2 16\"/>",
}


def glyph(name: str, size: int = 18) -> str:
    return (f"<svg width=\"{size}\" height=\"{size}\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" "
            f"stroke-linecap=\"round\" stroke-linejoin=\"round\" aria-hidden=\"true\">{GLYPHS[name]}</svg>")


TICK = "<svg class=\"tick\" width=\"14\" height=\"14\" viewBox=\"0 0 14 14\" aria-hidden=\"true\"><path d=\"M2.5 7.5l3 3 6-7\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"/></svg>"

ME = {"name": "Lightningbolltz", "cls": "shaman", "desc": "Shaman · 18", "battletag": "Hunthrop#1894"}
ALTS = [("Lightningbolltz", "shaman", "Orc Elemental Shaman · Level 18", True),
        ("Grimtusk", "warrior", "Orc Fury Warrior · Level 12", False),
        ("Sablewhisper", "priest", "Undead Shadow Priest · Level 9", False)]
DOORS = ["Tools", "Logs", "Guild", "Guides"]


HEADER_CSS = """
.hdr{position:relative;z-index:5;height:64px;border-bottom:1px solid rgba(255,255,255,.06);background:rgba(7,9,13,.82);backdrop-filter:blur(10px)}
.hdr .in{height:64px;display:flex;align-items:center;gap:28px}
.hdr .doors{display:flex;gap:6px;margin:0 auto}
.hdr .door{display:inline-flex;align-items:center;gap:6px;height:40px;padding:0 12px;border-radius:var(--rc);font-size:13px;font-weight:700;letter-spacing:.10em;text-transform:uppercase;color:var(--nav)}
.hdr .door.cur{color:var(--strong);background:rgba(255,255,255,.05)}
.hdr .right{display:flex;align-items:center;gap:10px}
.hbtn{display:inline-flex;align-items:center;justify-content:center;gap:8px;height:44px;min-width:44px;padding:0 10px;border-radius:var(--rc);border:1px solid var(--card-b);background:rgba(255,255,255,.03);color:var(--gold)}
.acct{display:inline-flex;align-items:center;gap:10px;height:48px;padding:0 12px 0 5px;border-radius:999px;border:1px solid var(--card-b);background:rgba(255,255,255,.035)}
.acct.open{border-color:var(--gold);box-shadow:var(--glow)}
.acct .nm{font-size:14px;font-weight:700;line-height:1.15}.acct .ds{font-size:12px;color:var(--muted);line-height:1.2}
.signin{display:inline-flex;align-items:center;gap:10px;height:48px;padding:0 16px 0 5px;border-radius:999px;border:1px solid rgba(229,185,85,.45);background:rgba(229,185,85,.08)}
.signin.open{border-color:var(--gold);box-shadow:var(--glow)}
.signin .av{width:36px;height:36px;border-radius:999px;display:grid;place-items:center;background:rgba(229,185,85,.14);color:var(--gold)}
.menu{position:absolute;right:48px;top:70px;width:380px;z-index:9;padding-bottom:6px}
.mbody .who{display:flex;align-items:center;gap:12px;padding:16px 18px 14px}
.mbody .row{padding:8px 18px;min-height:52px}
.mbody .lbl{display:flex;align-items:center;justify-content:space-between;padding:12px 18px 4px}
.cur-mark{color:var(--gold);display:flex}
.sheet{border-radius:var(--rp) var(--rp) 0 0}
.sheet .grab{display:flex;justify-content:center;padding:8px 0 0}.sheet .grab span{width:40px;height:4px;border-radius:2px;background:#3a4256}
"""


def header(current: str = "", state: str = "in", menu: str = "") -> str:
    doors = "".join(f"<span class=\"door{" cur" if d == current else ""}\">{d}{glyph("chev-down", 12)}</span>" for d in DOORS)
    doors += "<span class=\"door\">Get set up</span>"
    opened = " open" if menu else ""
    if state == "in":
        ctl = (f"<span class=\"acct{opened}\"><span>{crest("shaman", 36)}</span>"
               f"<span><span class=\"nm\" style=\"color:{CLASS_COLOR["shaman"]}\">{ME["name"]}</span><br><span class=\"ds\">{ME["desc"]}</span></span>"
               f"<span class=\"chev\" style=\"display:flex\">{glyph("chev-down", 16)}</span></span>")
    else:
        ctl = (f"<span class=\"signin{opened}\"><span class=\"av\">{glyph("user", 18)}</span>"
               f"<span><span style=\"font-size:14px;font-weight:700;color:var(--strong)\">Sign in</span><br>"
               f"<span style=\"font-size:12px;color:var(--muted)\">or paste an export</span></span></span>")
    return (f"<header class=\"hdr\"><div class=\"page-col in\"><span class=\"wordmark\"><img src=\"{SEAL}\" alt=\"\" width=\"32\" height=\"32\">Forever Sixty</span>"
            f"<nav class=\"doors\">{doors}</nav><span class=\"right\"><span class=\"hbtn\">{glyph("search", 18)}<span class=\"kbd\">/</span></span>{ctl}</span></div>{menu}</header>")


def header_phone(state: str = "in") -> str:
    if state == "in":
        ctl = f"<span class=\"hbtn\" style=\"border-radius:999px;padding:0;width:44px;border-color:{CLASS_COLOR["shaman"]}88\">{crest("shaman", 34)}</span>"
    else:
        ctl = f"<span class=\"hbtn\" aria-label=\"Sign in\" style=\"border-radius:999px;width:44px;padding:0;border-color:rgba(229,185,85,.45);background:rgba(229,185,85,.10)\">{glyph("user", 18)}</span>"
    return (f"<header class=\"hdr\" style=\"height:56px\"><div class=\"page-col in\" style=\"height:56px;gap:8px\"><span class=\"wordmark\" style=\"font-size:14px;flex:1\"><img src=\"{SEAL}\" alt=\"\" width=\"28\" height=\"28\">Forever Sixty</span>"
            f"<span class=\"hbtn\" style=\"border:0;background:none;color:var(--nav)\">{glyph("search", 20)}</span>{ctl}<span class=\"hbtn\" style=\"color:var(--strong)\">{glyph("menu", 20)}</span></div></header>")


def char_rows() -> str:
    out = ""
    for name, cls, desc, cur in ALTS:
        mark = f"<span class=\"cur-mark\">{glyph("check", 18)}</span>" if cur else ""
        out += (f"<div class=\"row link{" hover" if cur else ""}\" style=\"border-top:0\">{crest(cls, 36)}<span class=\"t\"><span class=\"lab\" style=\"color:{CLASS_COLOR[cls]}\">{name}</span>"
                f"<div class=\"sub\">{desc}</div></span>{mark}</div>")
    return out


ACCOUNT_ROWS = [("user", "Your profile", "Lightningbolltz’s score, gear and fights", ""),
                ("sliders", "Account", "Characters, devices, Discord", ""),
                ("crown", "Plan", "Premium · renews 27 Oct 2027", "<span class=\"tag tag-gold\">Founder</span>"),
                ("logout", "Sign out", "", "")]


def account_rows() -> str:
    out = ""
    for icon, lab, sub, extra in ACCOUNT_ROWS:
        sub_html = f"<div class=\"sub\">{sub}</div>" if sub else ""
        chev = "" if icon == "logout" else f"<span class=\"chev\" style=\"display:flex\">{glyph("chev-right", 16)}</span>"
        tone = " style=\"color:var(--muted)\"" if icon == "logout" else ""
        out += f"<div class=\"row link\" style=\"border-top:0\"><span class=\"itile sm\"{tone}>{glyph(icon, 16)}</span><span class=\"t\"><span class=\"lab\">{lab}</span>{sub_html}</span>{extra}{chev}</div>"
    return out


PASTE_ROW = "<div class=\"row link\" style=\"border-top:0\"><span class=\"itile sm\">{icon}</span><span class=\"t\"><span class=\"lab\">Paste an export</span><div class=\"sub\">From the addon. Adds a character to this list.</div></span></div>"


def account_menu_body(state: str = "in") -> str:
    return "<div class=\"mbody\">" + _menu_body(state) + "</div>"


def _menu_body(state: str) -> str:
    paste = PASTE_ROW.format(icon=glyph("paste", 16))
    if state == "in":
        return (f"<div class=\"who\"><span class=\"itile\">{glyph("user", 18)}</span><span style=\"flex:1;min-width:0\"><span style=\"font-size:15px;font-weight:700;color:var(--strong)\">{ME["battletag"]}</span>"
                f"<div style=\"font-size:13px;color:var(--muted)\">Signed in with Battle.net</div></span><span class=\"tag tag-gold\"><span class=\"dot\"></span>Premium</span></div>"
                f"<div class=\"lbl\" style=\"border-top:1px solid var(--line-soft)\"><span class=\"micro\">Your characters</span><span style=\"color:var(--gold);font-size:13px;font-weight:600\">Manage</span></div>"
                f"{char_rows()}{paste}<div style=\"border-top:1px solid var(--line-soft);margin-top:6px;padding-top:6px\">{account_rows()}</div>")
    return (f"<div class=\"who\" style=\"flex-direction:column;align-items:stretch;gap:12px\"><span class=\"h3\">Add your character</span>"
            f"<span style=\"font-size:14px;color:var(--text)\">Sign in and your characters arrive with their gear, talents and guild.</span>"
            f"<span class=\"btn btn-primary full\">Sign in with Battle.net</span><span class=\"btn btn-secondary full\">Use an email link</span></div>"
            f"<div style=\"border-top:1px solid var(--line-soft);padding-top:6px\">{paste}"
            f"<div class=\"row link\" style=\"border-top:0\"><span class=\"itile sm\">{glyph("download", 16)}</span><span class=\"t\"><span class=\"lab\">Get the addon</span><div class=\"sub\">Export a character from the game.</div></span></div></div>")


def account_menu(state: str = "in", style: str = "") -> str:
    return f"<div class=\"menu card\" style=\"{style}\">{account_menu_body(state)}</div>"


FOOTER_CSS = """
.ftr{margin-top:auto;border-top:1px solid rgba(255,255,255,.06);background:linear-gradient(180deg,rgba(255,255,255,.025),transparent 70%)}
.ftr .top{display:grid;grid-template-columns:minmax(0,1.7fr) repeat(4,minmax(0,1fr));gap:32px;padding:48px 0 36px}
.ftr .brand p{font-size:15px;color:var(--text);margin-top:14px;max-width:340px;line-height:1.5}
.ftr .col h4{margin:0 0 8px;font-size:11px;font-weight:700;letter-spacing:.14em;text-transform:uppercase;color:var(--muted)}
.ftr .col a{display:flex;align-items:center;min-height:32px;font-size:14px;color:var(--nav)}
.ftr .community{display:flex;flex-wrap:wrap;gap:10px;margin-top:20px}
.ftr .legal{display:flex;justify-content:space-between;gap:24px;padding:18px 0 28px;border-top:1px solid var(--line-soft);font-size:12px;color:var(--muted);line-height:1.5}
.phone .ftr .top{grid-template-columns:1fr 1fr;gap:28px 16px;padding:36px 0 20px}
.phone .ftr .brand{grid-column:1/-1}
.phone .ftr .col a{min-height:44px}
.ftr .legal-links{display:flex;gap:4px;white-space:nowrap}.ftr .legal-links a{display:inline-flex;align-items:center;min-height:32px;padding:0 8px;font-size:13px;color:var(--nav)}
.phone .ftr .legal{flex-direction:column-reverse;gap:8px}.phone .ftr .legal-links a{min-height:44px;padding:0 12px 0 0}
"""

FOOTER_COLUMNS = [
    ("Play", ["Get set up", "Guides", "Best in slot", "Talent planner", "Simulator"]),
    ("Logs", ["Upload a log", "Your reports", "Rankings", "Tier list"]),
    ("Guilds", ["Find a guild", "Claim your guild", "Guild plans"]),
    ("Forever Sixty", ["About", "Premium", "Changelog", "Sources"]),
]


def footer() -> str:
    cols = "".join(f"<div class=\"col\"><h4>{h}</h4>" + "".join(f"<a>{label}</a>" for label in links) + "</div>" for h, links in FOOTER_COLUMNS)
    return (f"<footer class=\"ftr\"><div class=\"page-col\"><div class=\"top\"><div class=\"brand\"><span class=\"wordmark\"><img src=\"{SEAL}\" alt=\"\" width=\"30\" height=\"30\">Forever Sixty</span>"
            f"<p>Play your class better: gear, talents, sims and logs for World of Warcraft: Forever.</p>"
            f"<div class=\"community\"><span class=\"btn btn-secondary sm\" style=\"text-transform:none;letter-spacing:0;font-size:14px;font-weight:600\"><span style=\"color:#9aa4ff;display:flex\">{discord_mark(16)}</span>Join the Discord</span>"
            f"<span class=\"btn btn-secondary sm\" style=\"text-transform:none;letter-spacing:0;font-size:14px;font-weight:600\">{glyph("repo", 16)}Contribute</span></div></div>{cols}</div>"
            f"<div class=\"legal\"><span>Not affiliated with or endorsed by Blizzard Entertainment. World of Warcraft and Warcraft are trademarks of Blizzard Entertainment, Inc.<br>"
            f"Forever Sixty is operated by Commish LLC. © 2026 Commish LLC.</span>"
            f"<span class=\"legal-links\"><a>Terms</a><a>Privacy</a><a>Refunds</a></span></div></div></footer>")
