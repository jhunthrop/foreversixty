"""Boards for design/specs/2026-10-29-premium-entry.md. Run from the repo root:
python3 design/mocks/2026-10-29-premium-entry/gen_premium_entry.py && node design/mocks/2026-10-29-premium-entry/render.mjs
Draws on the 2026-10-28 polish parts (direction B, Forge): the same header, account panel and
tokens, plus the Premium nav tag, the free Plan row's Upgrade, the sim's server-run link and the
/premium#server-sims landing. Every string is the spec's."""
import sys
from pathlib import Path

HERE = Path(__file__).resolve().parent
sys.path.insert(0, str(HERE.parent / "2026-10-28-polish"))
from parts import (CLASS_COLOR, HEADER_CSS, ME, PASTE_ROW, SEAL, char_rows, crest, glyph, header_phone)  # noqa: E402
from theme import head  # noqa: E402

RENDERS = HERE / "renders"
DOORS = ["Tools", "Logs", "Guild", "Guides"]

CSS = HEADER_CSS + """
.board{display:flex;flex-direction:column;min-height:100vh;background:#07090d}
.dirbar{display:flex;align-items:center;gap:12px;height:34px;padding:0 20px;background:#000;border-bottom:1px solid #1c2230;font-size:12px;color:#9a9484;font-family:"JetBrains Mono",monospace}
.dirbar b{color:#e5b955}
.cap{padding:14px 16px 8px;font-size:12px;color:#9a9484;font-family:"JetBrains Mono",monospace;background:#000}
.band{background:var(--band)}
.pill{display:inline-flex;align-items:center;height:22px;padding:0 8px;border-radius:var(--rpill);font-size:11px;font-weight:700;letter-spacing:.08em;text-transform:uppercase;border:1px solid;white-space:nowrap}
.pill-site{color:var(--gold);background:rgba(229,185,85,.14);border-color:rgba(229,185,85,.35)}
.pill-site.hover{background:rgba(229,185,85,.22);border-color:rgba(229,185,85,.6);color:var(--gold-hover)}
.pill-site.cur{background:rgba(229,185,85,.22);border-color:var(--gold)}
.pill-site.act{background:rgba(229,185,85,.30)}
.door.prem{padding:0 4px}
.door.prem.focus{outline:2px solid var(--gold);outline-offset:2px}
.upg{display:inline-flex;align-items:center;gap:4px;color:var(--gold);font-size:13px;font-weight:600}
.mbody .row{padding:8px 18px;min-height:52px}
.states{display:flex;gap:28px;align-items:center;padding:22px 48px}
.states .s{display:flex;flex-direction:column;gap:10px;align-items:flex-start}
.feat{display:flex;gap:10px;padding:6px 8px;border-radius:var(--rc);font-size:14px;line-height:1.5;color:var(--text)}
.feat.hl{background:rgba(229,185,85,.08);box-shadow:inset 2px 0 0 var(--gold)}
.feat svg{flex:none;margin-top:4px;color:var(--gold)}
.runbtn{display:inline-flex;align-items:center;gap:4px;height:36px;padding:0 16px;border:1px solid #3a3226;border-radius:var(--rc);color:var(--nav);font-size:12px;font-weight:700;letter-spacing:.12em;text-transform:uppercase}
.selectish{display:inline-flex;align-items:center;gap:8px;height:36px;padding:0 12px;border:1px solid var(--card-b);border-radius:var(--rc);font-size:13px;color:var(--strong)}
"""

TICK = "<svg width=\"14\" height=\"14\" viewBox=\"0 0 14 14\" aria-hidden=\"true\"><path d=\"M2.5 7.5l3 3 6-7\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"/></svg>"


def board(body: str, label: str, phone: bool = False) -> str:
    style = "width:390px" if phone else ""
    return head("B", CSS) + f"<div class=\"board\" style=\"{style}\"><div class=\"dirbar\"><b>Premium entry · 2026-10-29</b><span>{label}</span></div>{body}</div></body></html>"


def scene(inner: str, height: int, label: str) -> str:
    return f"<div class=\"cap\">{label}</div><div style=\"position:relative;height:{height}px;overflow:hidden;background:var(--band);border-bottom:1px solid #1c2230\">{inner}</div>"


def faint() -> str:
    return ("<div class=\"page-col\" style=\"padding-top:56px;opacity:.35;filter:saturate(.6)\"><span class=\"micro gold\">World of Warcraft: Forever</span>"
            "<div class=\"h1\" style=\"margin:18px 0\">Play your class better.</div><p style=\"font-size:16px\">Pick your class for its best-in-slot list, talents and rotation.</p></div>")


def tag(state: str = "") -> str:
    return f"<span class=\"pill pill-site {state}\">Premium</span>"


def header(state: str, premium_tag: bool, menu: str = "", tag_state: str = "", focus: bool = False) -> str:
    doors = "".join(f"<span class=\"door\">{d}{glyph("chev-down", 12)}</span>" for d in DOORS)
    doors += "<span class=\"door\">Get set up</span>"
    if premium_tag:
        doors += f"<span class=\"door prem{" focus" if focus else ""}\">{tag(tag_state)}</span>"
    opened = " open" if menu else ""
    if state == "in":
        ctl = (f"<span class=\"acct{opened}\"><span>{crest("shaman", 36)}</span>"
               f"<span><span class=\"nm\" style=\"color:{CLASS_COLOR["shaman"]}\">{ME["name"]}</span><br><span class=\"ds\">{ME["desc"]}</span></span>"
               f"<span class=\"chev\" style=\"display:flex\">{glyph("chev-down", 16)}</span></span>")
    else:
        ctl = (f"<span class=\"signin\"><span class=\"av\">{glyph("user", 18)}</span>"
               "<span><span style=\"font-size:14px;font-weight:700;color:var(--strong)\">Sign in</span><br>"
               "<span style=\"font-size:12px;color:var(--muted)\">or paste an export</span></span></span>")
    return (f"<header class=\"hdr\"><div class=\"page-col in\"><span class=\"wordmark\"><img src=\"{SEAL}\" alt=\"\" width=\"32\" height=\"32\">Forever Sixty</span>"
            f"<nav class=\"doors\" style=\"align-items:center\">{doors}</nav><span class=\"right\"><span class=\"hbtn\">{glyph("search", 18)}<span class=\"kbd\">/</span></span>{ctl}</span></div>{menu}</header>")


def plan_row(free: bool) -> str:
    if free:
        sub, trail = "Free", f"<span class=\"upg\">Upgrade{glyph("chev-right", 16)}</span>"
    else:
        sub = "Premium · $50/yr</div><div class=\"sub\" style=\"margin-top:0\">Renews 27 Oct 2027"
        trail = f"<span class=\"tag tag-gold\">Founder</span><span class=\"chev\" style=\"display:flex\">{glyph("chev-right", 16)}</span>"
    return (f"<div class=\"row link\" style=\"border-top:0\"><span class=\"itile sm\">{glyph("crown", 16)}</span><span class=\"t\"><span class=\"lab\">Plan</span>"
            f"<div class=\"sub\">{sub}</div></span>{trail}</div>")


def simple_row(icon: str, lab: str, sub: str, quiet: bool = False) -> str:
    sub_html = f"<div class=\"sub\">{sub}</div>" if sub else ""
    chev = "" if quiet else f"<span class=\"chev\" style=\"display:flex\">{glyph("chev-right", 16)}</span>"
    tone = " style=\"color:var(--muted)\"" if quiet else ""
    return f"<div class=\"row link\" style=\"border-top:0\"><span class=\"itile sm\"{tone}>{glyph(icon, 16)}</span><span class=\"t\"><span class=\"lab\">{lab}</span>{sub_html}</span>{chev}</div>"


def panel_body(free: bool) -> str:
    head_tag = "" if free else "<span class=\"tag tag-gold\"><span class=\"dot\"></span>Premium</span>"
    return (f"<div class=\"mbody\"><div class=\"who\"><span class=\"itile\">{glyph("user", 18)}</span><span style=\"flex:1;min-width:0\"><span style=\"font-size:15px;font-weight:700;color:var(--strong)\">{ME["battletag"]}</span>"
            f"<div style=\"font-size:13px;color:var(--muted)\">Signed in with Battle.net</div></span>{head_tag}</div>"
            "<div class=\"lbl\" style=\"border-top:1px solid var(--line-soft)\"><span class=\"micro\">Your characters</span><span style=\"color:var(--gold);font-size:13px;font-weight:600\">Manage</span></div>"
            f"{char_rows()}{PASTE_ROW.format(icon=glyph("paste", 16))}<div style=\"border-top:1px solid var(--line-soft);margin-top:6px;padding-top:6px\">"
            f"{simple_row("user", "Character page", "Score, gear and fights")}{simple_row("sliders", "Account", "Devices, Discord, notifications")}"
            f"{plan_row(free)}{simple_row("logout", "Sign out", "", True)}</div></div>")


def panel(free: bool, style: str) -> str:
    return f"<div class=\"menu card\" style=\"{style}\">{panel_body(free)}</div>"


def header_board(width: int) -> str:
    right = "right:calc(max(0px, (100% - 1344px) / 2) + 48px)"
    s1 = scene(header("out", True) + faint(), 190, "1 · Signed out: the Premium tag after Get set up. No dot, no panel, no chevron.")
    s2 = scene(header("in", True) + faint(), 190, "2 · Signed in, free: the same tag.")
    s3 = scene(header("in", False) + faint(), 190, "3 · Premium from any source (own, seat, grant, fallback): no tag. Hidden before first paint.")
    states = "".join(f"<span class=\"s\"><span class=\"note\">{lab}</span><span class=\"door prem{" focus" if st == "focus" else ""}\" style=\"display:inline-flex;align-items:center;height:40px\">{tag(st if st != "focus" else "")}</span></span>"
                     for lab, st in [("Default", ""), ("Hover", "hover"), ("Focus-visible", "focus"), ("Active", "act"), ("Current (/premium)", "cur")])
    s4 = f"<div class=\"cap\">4 · The tag's states</div><div class=\"states\" style=\"background:var(--bg)\">{states}</div>"
    s5 = scene(header("in", True, panel(True, right)) + faint(), 820, "5 · Free, account control open: the Plan row reads Free with Upgrade, one link to /premium#premium.")
    return board(s1 + s2 + s3 + s4 + s5, f"Header with the Premium tag at {width}")


def menu_sheet(premium_row: bool) -> str:
    doors = "".join(f"<div class=\"row link\" style=\"min-height:52px\"><span class=\"t lab\" style=\"text-transform:uppercase;letter-spacing:.1em;font-size:14px\">{d}</span><span class=\"chev\" style=\"display:flex\">{glyph("chev-down", 16)}</span></div>" for d in DOORS)
    doors += "<div class=\"row link\" style=\"min-height:52px\"><span class=\"t lab\" style=\"text-transform:uppercase;letter-spacing:.1em;font-size:14px\">Get set up</span></div>"
    if premium_row:
        doors += f"<div class=\"row link\" style=\"min-height:52px\"><span class=\"t\">{tag()}</span><span class=\"chev\" style=\"display:flex\">{glyph("chev-right", 16)}</span></div>"
    return (f"<div class=\"card\" style=\"position:absolute;left:0;right:0;top:56px;border-radius:0;padding-bottom:8px\">"
            f"<div style=\"padding:12px 18px\"><span class=\"field\" style=\"height:48px\"><span style=\"color:var(--gold);display:flex\">{glyph("search", 18)}</span>Character name</span></div>{doors}</div>")


def phone_board() -> str:
    sheet = (f"<div style=\"position:absolute;inset:0;background:rgba(0,0,0,.55)\"></div><div class=\"card sheet\" style=\"position:absolute;left:0;right:0;bottom:0;max-height:88%;overflow:hidden\">"
             f"<div class=\"grab\"><span></span></div>{panel_body(True)}</div>")
    s1 = scene(header_phone("in") + menu_sheet(True), 560, "1 · Free: Menu open, Premium is the last row")
    s2 = scene(header_phone("in") + menu_sheet(False), 500, "2 · Premium: no Premium row")
    s3 = scene(header_phone("in") + sheet, 844, "3 · Free: account sheet, Plan · Free · Upgrade")
    return board(s1 + s2 + s3, "Phone menu and account sheet at 390", True)


def account_board() -> str:
    left = panel(True, "position:relative;right:auto;top:auto")
    right = panel(False, "position:relative;right:auto;top:auto")
    body = (f"<div style=\"display:flex;gap:48px;justify-content:center;padding:32px 0 48px;background:var(--band)\">"
            f"<div><div class=\"note\" style=\"margin-bottom:10px\">Free: Upgrade on the Plan row, no head tag</div>{left}</div>"
            f"<div><div class=\"note\" style=\"margin-bottom:10px\">Premium (own plan): unchanged</div>{right}</div></div>")
    return board(body, "Account panel, free beside Premium, at 1440")


FEATURES = [
    ("coach", "Coach: after every logged raid, what to fix first. It names the casts you missed, the buffs you let drop and the consumables you skipped."),
    ("server-sims", "Try up to 5,000 gear, talent or drop combinations in one sim and find your best DPS. Your browser stops at 400."),
    ("server-sims", "Close the tab. Server sims keep running, and the site tells you when they finish."),
    ("server-sims", "Runs go up to 100,000 iterations instead of 30,000, so close upgrades come out clear."),
    ("history", "See if you are improving: your Forever Score and every boss parse on a chart over time, with the fights behind each point."),
    ("notifications", "A message in your own Discord channel when a server sim finishes or a log is parsed."),
    (None, "A Supporter mark beside your name on your characters and in your guild's roster."),
]


def run_control(free: bool) -> str:
    serv = (f"<span class=\"runbtn\">Run on our servers <span class=\"pill pill-site\" style=\"margin-left:4px\">Premium</span></span>" if free
            else "<span class=\"runbtn\">Run on our servers</span>")
    return (f"<div class=\"card\" style=\"padding:16px 20px;display:flex;align-items:center;gap:12px;flex-wrap:wrap\">"
            f"<span class=\"btn btn-primary sm\">Run sim</span><span class=\"selectish\">Precision · Standard{glyph("chev-down", 14)}</span>{serv}</div>")


def upsell_board() -> str:
    feats = "".join(f"<div class=\"feat{" hl" if a == "server-sims" else ""}\">{TICK}<span>{t}</span></div>" for a, t in FEATURES)
    card = (f"<div class=\"card\" style=\"width:440px;padding:22px 22px 18px\"><div class=\"h2\">Premium</div><div class=\"note\" style=\"margin:4px 0 6px\">For one player</div>"
            f"<div style=\"font-family:var(--fh);font-weight:800;font-size:40px;color:var(--strong)\">$50<span style=\"font-size:16px;color:var(--muted)\"> /year</span></div>"
            f"<span class=\"btn btn-primary full\" style=\"margin:12px 0 14px\">Subscribe · $50</span>{feats}</div>")
    s1 = (f"<div class=\"cap\">1 · /sim, signed in. Free: the payer's button as a link with the Premium tag, to /premium#server-sims. Premium: today's button.</div>"
          f"<div style=\"display:grid;grid-template-columns:1fr 1fr;gap:24px;padding:28px 48px;background:var(--band)\"><div><div class=\"note\" style=\"margin-bottom:8px\">Free</div>{run_control(True)}</div>"
          f"<div><div class=\"note\" style=\"margin-bottom:8px\">Premium</div>{run_control(False)}</div></div>")
    s2 = (f"<div class=\"cap\">2 · Landing on /premium#server-sims: the three server-sim lines tinted, a 2 px gold rule, no reflow, no script.</div>"
          f"<div style=\"padding:28px 48px 48px;background:var(--band)\">{card}</div>")
    return board(s1 + s2, "Contextual upsell: sim run control and its /premium landing, at 1440")


BOARDS = {
    "premium-entry-header-1440": lambda: header_board(1440),
    "premium-entry-header-2000": lambda: header_board(2000),
    "premium-entry-phone": phone_board,
    "premium-entry-account-1440": account_board,
    "premium-entry-upsell-1440": upsell_board,
}


def main() -> None:
    RENDERS.mkdir(exist_ok=True)
    for name in sys.argv[1:] or list(BOARDS):
        (RENDERS / f"{name}.html").write_text(BOARDS[name]())
        print(name)


if __name__ == "__main__":
    main()
