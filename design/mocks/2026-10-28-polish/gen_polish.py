"""Boards for design/specs/2026-10-28-polish.md. Run from the repo root:
python3 design/mocks/2026-10-28-polish/gen_polish.py && node design/mocks/2026-10-28-polish/render.mjs
Writes HTML boards into renders/; render.mjs screenshots each at its width (name rules there).
Direction A (Frost) and B (Forge) draw the same three surfaces; B is the recommended direction,
so the header, footer, premium, account and component boards use it."""
import sys
from pathlib import Path

sys.path.insert(0, str(Path(__file__).resolve().parent))
from parts import (ALTS, CLASS_COLOR, CLASSES, FOOTER_CSS, HEADER_CSS, ME, TICK, account_menu, account_menu_body, asset,  # noqa: E402
                   build_icon, crest, discord_mark, footer, glyph, header, header_phone, uri, HERE)
from theme import DIRECTIONS, head  # noqa: E402

RENDERS = HERE / "renders"

PAGE_CSS = HEADER_CSS + FOOTER_CSS + """
.board{display:flex;flex-direction:column;min-height:100vh;background:#07090d}
.dirbar{display:flex;align-items:center;gap:12px;height:34px;padding:0 20px;background:#000;border-bottom:1px solid #1c2230;font-size:12px;color:#9a9484;font-family:"JetBrains Mono",monospace}
.dirbar b{color:#e5b955;font-weight:700}
.band{background:var(--band);border-bottom:1px solid rgba(255,255,255,.05)}
.eyebrow{display:flex;align-items:center;gap:10px}.eyebrow i{width:28px;height:1px;background:var(--gold)}
.lede{font-size:16px;color:var(--text);line-height:1.5;max-width:600px}
.phone .lede{font-size:15px}
.crestrow{display:flex;gap:14px}
.crestrow .c{display:flex;flex-direction:column;align-items:center;gap:6px;width:68px}
.crestrow .c span{font-size:11px;font-weight:700;letter-spacing:.08em;text-transform:uppercase}
.phone .crestrow{display:grid;grid-template-columns:repeat(5,1fr);gap:12px 6px}
.phone .crestrow .c{width:auto}\n.price-note{min-height:58px}.phone .price-note{min-height:0}
.row .vsub{display:none}.phone .row .vsub{display:block}.phone .row .val{display:none}
.phone .renew{display:none}.renew-sub{display:none}.phone .renew-sub{display:block}
.cap{padding:14px 16px 8px;font-size:12px;color:#9a9484;font-family:\"JetBrains Mono\",monospace;background:#000}
"""


def dirbar(key: str, what: str) -> str:
    return f"<div class=\"dirbar\"><b>Direction {key} · {DIRECTIONS[key]["name"]}</b><span>{what}</span></div>"


def board(key: str, body: str, phone: bool = False, label: str = "", extra_css: str = "") -> str:
    width = 390 if phone else None
    style = f"width:{width}px" if width else ""
    bar = dirbar(key, label) if label else ""
    return head(key, PAGE_CSS + extra_css) + f"<div class=\"board{" phone" if phone else ""}\" style=\"{style}\">{bar}{body}</div></body></html>"


def crest_row(size: int) -> str:
    return "<div class=\"crestrow\">" + "".join(
        f"<span class=\"c\">{crest(c, size)}<span style=\"color:{CLASS_COLOR[c]}\">{c.title()}</span></span>" for c in CLASSES) + "</div>"


def launch_card(phone: bool) -> str:
    tiles = "".join(f"<div class=\"stat\" style=\"flex:1\"><span class=\"n mono\" style=\"font-family:var(--fh)\">{n}</span><span class=\"micro\">{u}</span></div>"
                    for n, u in [("07", "Days"), ("02", "Hours"), ("41", "Minutes")])
    icons = "".join(f"<img class=\"icon\" src=\"{build_icon(s)}\" alt=\"\" width=\"36\" height=\"36\">" for s in ["inv_weapon_bow_08", "inv_helmet_15", "inv_jewelry_ring_13"])
    return (f"<div class=\"card\" style=\"padding:20px;display:flex;flex-direction:column;gap:16px{";width:100%" if phone else ""}\">"
            f"<div style=\"display:flex;justify-content:space-between;align-items:center\"><span class=\"micro gold\" style=\"display:flex;align-items:center;gap:8px\"><span class=\"dot\" style=\"color:var(--gold)\"></span>Launch · Nov 4 · 15:00 PST</span></div>"
            f"<div style=\"display:flex;gap:10px\">{tiles}</div>"
            f"<div style=\"border-top:1px solid var(--line-soft);padding-top:14px;display:flex;align-items:center;gap:12px\">{crest("hunter", 44)}"
            f"<span style=\"flex:1;min-width:0\"><span class=\"micro\">Example best in slot</span><div style=\"font-size:15px;font-weight:700;color:{CLASS_COLOR["hunter"]}\">Marksmanship Hunter</div>"
            f"<div style=\"font-size:13px;color:var(--muted)\">Leveling best in slot at every level range</div></span></div>"
            f"<div style=\"display:flex;align-items:center;gap:8px\">{icons}<span class=\"alink\" style=\"margin-left:auto\">All 28 specs{glyph("chev-right", 16)}</span></div></div>")


def lookup(phone: bool) -> str:
    sel = lambda v: f"<span class=\"select\" style=\"{"flex:1" if phone else "width:140px"}\">{v}<span style=\"color:var(--muted);display:flex\">{glyph("chev-down", 16)}</span></span>"
    field = f"<span class=\"field\" style=\"{"" if phone else "flex:1"}\"><span style=\"color:var(--gold);display:flex\">{glyph("search", 18)}</span><span style=\"flex:1\">Character name</span><span class=\"kbd\">/</span></span>"
    if phone:
        controls = f"{field}<div style=\"display:flex;gap:10px\">{sel("US")}{sel("Normal")}</div><span class=\"btn btn-primary full\">Look up</span>"
        return f"<div style=\"display:flex;flex-direction:column;gap:10px\"><span class=\"h3\">Look up a character</span><span style=\"font-size:14px;color:var(--muted)\">See any character’s Forever Score.</span>{controls}</div>"
    return (f"<div style=\"display:flex;flex-direction:column;gap:12px\"><div style=\"display:flex;align-items:baseline;gap:14px\"><span class=\"h3\">Look up a character</span>"
            f"<span style=\"font-size:14px;color:var(--muted)\">See any character’s Forever Score.</span></div>"
            f"<div style=\"display:flex;gap:10px\">{field}{sel("US")}{sel("Normal")}<span class=\"btn btn-primary\" style=\"height:52px;padding:0 26px\">Look up</span></div></div>")


def home(key: str, phone: bool) -> str:
    hero_text = (f"<div style=\"display:flex;flex-direction:column;gap:{18 if phone else 22}px;min-width:0\">"
                 f"<span class=\"eyebrow micro gold\"><i></i>World of Warcraft: Forever</span>"
                 f"<h1 class=\"h1\" style=\"{"font-size:var(--h1-phone)" if phone else ""}\">Play your class <span class=\"grad\">better.</span></h1>"
                 f"<p class=\"lede\">Pick your class for its best-in-slot list, talents and rotation, or sign in and your own characters arrive with their gear, talents and guild.</p>"
                 f"{crest_row(52 if phone else 60)}"
                 f"<div style=\"display:flex;gap:12px;{"flex-direction:column" if phone else ""}\"><span class=\"btn btn-primary{" full" if phone else ""}\">Sign in with Battle.net</span>"
                 f"<span class=\"btn btn-secondary{" full" if phone else ""}\">{glyph("paste", 16)}Paste an addon export</span></div></div>")
    if phone:
        hero = f"<div class=\"page-col\" style=\"padding-top:28px;padding-bottom:28px;display:flex;flex-direction:column;gap:24px\">{hero_text}{launch_card(True)}{lookup(True)}</div>"
        top = header_phone("out")
    else:
        hero = (f"<div class=\"page-col\" style=\"padding-top:56px;padding-bottom:48px;display:flex;flex-direction:column;gap:44px\">"
                f"<div style=\"display:grid;grid-template-columns:minmax(0,1.45fr) minmax(0,1fr);gap:56px;align-items:center\">{hero_text}{launch_card(False)}</div>{lookup(False)}</div>")
        top = header(state="out")
    return board(key, f"<div class=\"band\">{top}{hero}</div><div style=\"height:120px\"></div>", phone, "Home hero and header, signed out")



def value_row(icon: str, label: str, value: str, link: bool = True, muted: bool = False) -> str:
    chev = f"<span class=\"chev\" style=\"display:flex\">{glyph("chev-right", 16)}</span>" if link else "<span style=\"width:16px\"></span>"
    tone = " muted" if muted else ""
    return (f"<div class=\"row{" link" if link else ""}\"><span class=\"itile sm\">{glyph(icon, 16)}</span>"
            f"<span class=\"t\"><span class=\"lab\">{label}</span><div class=\"sub vsub\">{value}</div></span>"
            f"<span class=\"val{tone}\">{value}</span>{chev}</div>")


def plan_card() -> str:
    return (f"<section class=\"card\" id=\"plan\"><div class=\"card-head\"><span class=\"itile\">{glyph("crown", 18)}</span>"
            f"<span class=\"grow\"><span class=\"micro\">Plan</span><div style=\"display:flex;align-items:center;gap:10px\"><span class=\"h2\" style=\"font-size:24px\">Premium</span><span class=\"tag tag-gold\">Founder</span></div><div class=\"renew-sub\" style=\"font-size:13px;color:var(--muted)\">Renews 27 Oct 2027</div></span>"
            f"<span class=\"renew\" style=\"text-align:right;font-size:13px;color:var(--muted)\">Renews<br><span style=\"color:var(--strong);font-weight:600\">27 Oct 2027</span></span></div>"
            f"{value_row("server", "Server sims", "On · 3 runs this month")}"
            f"{value_row("chart", "History charts", "Open on your main character")}"
            f"{value_row("bell", "Discord notifications", "On · 2 events")}"
            f"{value_row("star", "Supporter mark", "Shown on your characters")}"
            f"<div style=\"display:flex;gap:10px;padding:16px 20px;border-top:1px solid var(--line-soft)\"><span class=\"btn btn-secondary\">{glyph("card", 16)}Manage billing</span></div></section>")


def switch_row(label: str, on: bool, sub: str = "", nested: bool = False) -> str:
    sub_html = f"<div class=\"sub\">{sub}</div>" if sub else ""
    return (f"<div class=\"row{" nested" if nested else ""}\"><span class=\"t\"><span class=\"lab\" style=\"font-weight:{500 if nested else 600}\">{label}</span>{sub_html}</span>"
            f"<span class=\"switch{" on" if on else ""}\" role=\"switch\"></span></div>")


def discord_card(phone: bool = False) -> str:
    open_btn = f"<span class=\"btn btn-secondary sm\" style=\"text-transform:none;letter-spacing:0;font-size:13px\">Open in Discord{glyph("external", 14)}</span>"
    dest = (f"<div class=\"row\"><span class=\"itile sm\">{glyph("hash", 16)}</span><span class=\"t\"><span class=\"lab\">Connected to a Discord channel</span>"
            f"<div class=\"sub mono\" style=\"font-size:12px\">Channel 1198…4568</div></span>{"" if phone else open_btn}</div>")
    phone_open = f"<div style=\"padding:0 20px 12px 70px\">{open_btn}</div>" if phone else ""
    return (f"<section class=\"card\" id=\"discord\"><div class=\"card-head\"><span class=\"itile discord\">{discord_mark(18)}</span>"
            f"<span class=\"grow\"><span class=\"micro\">Notifications</span><div class=\"h2\" style=\"font-size:24px\">Discord</div></span>"
            f"<span class=\"tag tag-green\"><span class=\"dot\"></span>Connected</span></div>"
            f"{dest}{phone_open}"
            f"<div style=\"padding:14px 20px 4px;border-top:1px solid var(--line-soft)\"><span class=\"micro\">Post when</span></div>"
            f"<div style=\"margin-top:4px\">{switch_row("A server sim finishes", True).replace("class=\"row\"", "class=\"row\" style=\"border-top:0\"")}"
            f"{switch_row("A combat log you upload finishes parsing", True)}"
            f"{switch_row("Include private logs", False, "Off, private logs never post. On, everyone in the channel sees the log’s name and “Only you”.", True)}</div>"
            f"<div style=\"display:flex;flex-wrap:wrap;gap:10px;padding:16px 20px;border-top:1px solid var(--line-soft)\"><span class=\"btn btn-secondary\">{glyph("send", 16)}Send test message</span>"
            f"<span class=\"btn btn-danger\">Remove</span></div>"
            f"<div class=\"row link\"><span class=\"itile sm\">{glyph("refresh", 16)}</span><span class=\"t\"><span class=\"lab\">Connect a different channel</span>"
            f"<div class=\"sub\">Connecting again replaces this channel and removes its webhook from Discord.</div></span><span class=\"chev\" style=\"display:flex\">{glyph("chev-right", 16)}</span></div></section>")


def current_character() -> str:
    return (f"<section class=\"card\" style=\"padding:22px;display:flex;flex-direction:column;gap:18px\"><span class=\"micro\">Current character</span>"
            f"<div style=\"display:flex;align-items:center;gap:16px\">{crest("shaman", 64)}<span style=\"flex:1;min-width:0\">"
            f"<div style=\"display:flex;align-items:center;gap:10px;flex-wrap:wrap\"><span class=\"h2\" style=\"color:{CLASS_COLOR["shaman"]}\">Lightningbolltz</span><span class=\"tag tag-gold\">{glyph("star", 12)}Supporter</span></div>"
            f"<div style=\"font-size:14px;color:var(--muted);margin-top:4px\">Orc Shaman · Level 18 · Forever (Normal US)</div></span></div>"
            f"<div style=\"display:flex;align-items:center;gap:12px;flex-wrap:wrap;padding:14px;border-radius:var(--rc);background:rgba(255,255,255,.03);border:1px dashed var(--line)\">"
            f"<span style=\"flex:1;min-width:200px;font-size:14px\">No export yet. Paste one from the addon to see gear, talents and your score.</span>"
            f"<span class=\"btn btn-primary sm\">Paste an export</span></div></section>")


def characters_card() -> str:
    rows = ""
    for name, cls, desc, cur in ALTS[:2]:
        tags = "<span class=\"tag tag-gold\">Current</span><span class=\"tag tag-muted\">Main</span>" if cur else ""
        rows += (f"<div class=\"row link\">{crest(cls, 40)}<span class=\"t\"><span class=\"lab\" style=\"color:{CLASS_COLOR[cls]}\">{name}</span><div class=\"sub\">{desc}</div></span>"
                 f"{tags}<span class=\"chev\" style=\"display:flex\">{glyph("chev-right", 16)}</span></div>")
    return (f"<section class=\"card\"><div class=\"card-head\"><span class=\"itile\">{glyph("user", 18)}</span><span class=\"grow h3\">Characters</span><span style=\"color:var(--gold);font-size:14px;font-weight:600\">Manage</span></div>"
            f"<div class=\"row\" style=\"min-height:44px\"><span class=\"t\" style=\"font-size:13px;color:var(--muted)\">Last raid · no log in the last 14 days</span></div>{rows}</section>")


def account_head(phone: bool) -> str:
    who = (f"<div style=\"display:flex;align-items:center;gap:12px\"><span class=\"itile\">{glyph("user", 18)}</span><span><span style=\"font-weight:700;color:var(--strong)\">Hunthrop#1894</span>"
           f"<div style=\"font-size:13px;color:var(--muted)\">Signed in with Battle.net</div></span><span class=\"btn btn-secondary sm\" style=\"margin-left:{"auto" if phone else "8px"}\">{glyph("logout", 14)}Sign out</span></div>")
    return (f"<div style=\"display:flex;{"flex-direction:column;gap:16px" if phone else "align-items:flex-end;justify-content:space-between"};padding:{"24px 0 20px" if phone else "40px 0 28px"}\">"
            f"<div><span class=\"micro gold eyebrow\"><i></i>Account</span><h1 class=\"h1\" style=\"font-size:{"34px" if phone else "44px"};margin-top:10px\">Your account</h1></div>{who}</div>")


def account(key: str, phone: bool) -> str:
    if phone:
        body = (f"{header_phone("in")}<div class=\"page-col\">{account_head(True)}<div style=\"display:flex;flex-direction:column;gap:16px;padding-bottom:40px\">"
                f"{current_character()}{plan_card()}{discord_card(True)}</div></div>")
    else:
        body = (f"{header(state="in")}<div class=\"page-col\">{account_head(False)}"
                f"<div style=\"display:grid;grid-template-columns:minmax(0,1fr) 440px;gap:24px;align-items:start;padding-bottom:64px\">"
                f"<div style=\"display:flex;flex-direction:column;gap:20px\">{current_character()}{characters_card()}</div>"
                f"<div style=\"display:flex;flex-direction:column;gap:20px\">{plan_card()}{discord_card()}</div></div></div>")
    return board(key, body, phone, "Account page, Plan and Discord blocks (Premium, connected channel)")



FREE_FEATURES = ["Addon export and import", "BiS lists for every spec and level band", "Talent planner", "Browser sims", "Logs, live logging and parses", "Rankings and the tier list", "Class guides"]
PREMIUM_FEATURES = [
    "Coach: after every logged raid, what to fix first — the casts you missed, the buffs you let drop, the consumables you skipped — learned from the engine and from the top logs on the site.",
    "Try up to 5,000 gear, talent or drop combinations in one sim and find your best DPS. Your browser stops at 400.",
    "Close the tab. Server sims keep running, and the site tells you when they finish.",
    "Runs go up to 100,000 iterations instead of 30,000, so close upgrades come out clear.",
    "See if you are improving: your Forever Score and every boss parse on a chart over time, with the fights behind each point.",
    "A message in your own Discord channel when a server sim finishes or a log is parsed.",
    "A Supporter mark beside your name on your characters and in your guild’s roster.",
]


def ticks(lines: list[str]) -> str:
    return "<ul style=\"list-style:none;margin:0;padding:0\">" + "".join(
        f"<li style=\"display:flex;gap:10px;padding:8px 0;border-top:1px solid var(--line-soft);font-size:14px;line-height:1.45\">{TICK}<span>{line}</span></li>" for line in lines) + "</ul>"


def price(figure: str, unit: str, size: int) -> str:
    return (f"<div style=\"display:flex;align-items:baseline;gap:6px\"><span style=\"font-family:var(--fh);font-weight:var(--fhw);font-size:{size}px;line-height:1;color:var(--strong)\">{figure}</span>"
            f"<span style=\"font-size:14px;color:var(--muted)\">{unit}</span></div>")


def free_card() -> str:
    return (f"<article class=\"card quiet\" style=\"padding:24px;display:flex;flex-direction:column;gap:14px\"><div><span class=\"h3\" style=\"font-size:22px\">Free</span>"
            f"<div style=\"font-size:14px;color:var(--muted)\">For every player</div></div>{price("$0", "", 52)}"
            f"<div class=\"price-note\" style=\"font-size:13px;color:var(--muted)\">No card, no trial, no end date.</div>"
            f"<div style=\"min-height:44px;display:flex;align-items:center;font-size:14px;color:var(--text)\">Sign in with Battle.net from the top of any page.</div>"
            f"{ticks(FREE_FEATURES)}<div style=\"font-size:14px;font-weight:700;color:var(--strong)\">This list never shrinks.</div></article>")


def premium_card() -> str:
    return (f"<article class=\"card\" style=\"padding:24px;display:flex;flex-direction:column;gap:14px;border-color:rgba(229,185,85,.45);box-shadow:inset 0 1px 0 var(--hl),0 0 0 1px rgba(229,185,85,.12),0 18px 48px rgba(229,185,85,.10),var(--shadow)\">"
            f"<div style=\"display:flex;justify-content:space-between;align-items:flex-start\"><div><span class=\"h3\" style=\"font-size:22px;color:var(--gold)\">Premium</span>"
            f"<div style=\"font-size:14px;color:var(--muted)\">For one player</div></div><span class=\"tag tag-gold\">Founder</span></div>"
            f"{price("$40", "/year", 52)}<div style=\"font-size:13px;line-height:1.5;color:var(--muted)\" class=\"price-note\">$50 after 6 Jan 2027<br>$3.33 a month, billed yearly<br><span style=\"color:var(--gold)\">Founder price, kept while you stay subscribed.</span></div>"
            f"<span class=\"btn btn-primary full\">Subscribe · $40 a year</span>"
            f"<div style=\"font-size:13px;color:var(--text)\">Free stays complete. Premium pays for the servers and keeps the site going.</div>"
            f"{ticks(PREMIUM_FEATURES)}<div style=\"font-size:13px;color:var(--muted)\">Sims cover damage specs; healing and tanking specs are not simulated yet. The coach reviews every spec; damage specs get the most findings.</div></article>")


def coach_item(icon: str | None, title: str, why: str) -> str:
    img = f"<img class=\"icon\" src=\"{asset(icon)}\" alt=\"\" width=\"24\" height=\"24\">" if icon else ""
    return (f"<div style=\"padding:12px 0;border-top:1px solid var(--line-soft)\"><div style=\"display:flex;align-items:center;gap:10px;font-size:15px;font-weight:600;color:var(--strong)\">{img}{title}</div>"
            f"<div style=\"font-size:13px;color:var(--text);margin-top:4px\">{why}</div></div>")


def coach_card() -> str:
    return (f"<article class=\"card quiet\" style=\"padding:24px;display:flex;flex-direction:column;gap:14px\" aria-label=\"A sample coach review\">"
            f"<div style=\"display:flex;align-items:center;gap:10px\"><span class=\"tag tag-muted\">Sample</span><span class=\"micro\">What the coach says after a raid</span></div>"
            f"<div style=\"display:flex;align-items:center;gap:12px\">{crest("warrior", 44)}<span><span class=\"h3\">Your Fury coach</span>"
            f"<div style=\"font-size:13px;color:var(--muted)\">Fury Warrior · 40-player kill</div></span></div>"
            f"<p style=\"font-size:14px;line-height:1.5;color:var(--text)\">Solid kill: you were alive the whole fight and Battle Shout was up 100% of it. Almost everything you left behind sits in one button.</p>"
            f"<div style=\"border-left:2px solid var(--gold);padding:2px 0 2px 16px;display:flex;flex-direction:column;gap:8px\"><span class=\"micro gold\">Fix first</span>"
            f"<div style=\"display:flex;align-items:center;gap:12px\"><img class=\"icon\" src=\"{asset("spell_nature_bloodlust")}\" alt=\"\" width=\"36\" height=\"36\" style=\"--q:#a8762a;box-shadow:var(--glow)\">"
            f"<span style=\"font-size:19px;font-weight:700;line-height:1.25;color:var(--strong)\">Bloodthirst: 9 casts lost</span></div>"
            f"<p style=\"font-size:14px;line-height:1.5\">Bloodthirst is the button the Fury rotation is built around, and 9 casts went unused while you had the rage, longest at 1:12–1:25. At 2:44–2:55 you sat on a full bar instead, so that one gap cost you twice.</p></div>"
            f"<div><span class=\"micro\">Then</span>{coach_item(None, "No flask at the pull", "A flask is the cheapest damage you will ever buy; bring one to every 40-player boss.")}"
            f"{coach_item("ability_rogue_ambush", "Rage full for 14 seconds", "When the bar is full, Heroic Strike turns rage into damage instead of losing it.")}</div>"
            f"<div><span class=\"micro\" style=\"color:var(--kill)\">Doing well</span><div style=\"display:flex;gap:10px;align-items:flex-start;margin-top:8px;font-size:14px\">{TICK}"
            f"<img class=\"icon\" src=\"{asset("ability_warrior_battleshout")}\" alt=\"\" width=\"20\" height=\"20\"><span>Battle Shout was up 100% of the fight, against 97% for the top 25.</span></div></div>"
            f"<div style=\"margin-top:auto;padding-top:12px;border-top:1px solid var(--line-soft);font-size:13px;color:var(--muted)\">A sample review. Yours is written from your own log after every raid you upload.</div></article>")


def plans_band(phone: bool) -> str:
    seg = f"<span class=\"seg\" style=\"{"display:flex" if phone else ""}\"><span class=\"on\" style=\"{"flex:1;justify-content:center" if phone else ""}\">Yearly</span><span style=\"{"flex:1;justify-content:center" if phone else ""}\">Monthly</span></span>"
    text = (f"<div style=\"display:flex;flex-direction:column;gap:12px;min-width:0\"><span class=\"micro gold eyebrow\"><i></i>Plans</span>"
            f"<h1 class=\"h1\" style=\"font-size:{"32px" if phone else "44px"}\">What is free, and what the plans add</h1>"
            f"<p class=\"lede\" style=\"max-width:720px\">Everything a player needs is free, and the free list never shrinks. Premium and the guild plans pay for the site. There are no ads for anyone.</p>"
            f"<div style=\"display:flex;gap:10px;align-items:flex-start;font-size:13px\"><span class=\"tag tag-gold\">Founder</span><span>Founder prices until 6 January 2027. Buy by then and the price stays for as long as you stay subscribed.</span></div>"
            f"{"" if phone else "<div style=\"font-size:14px\">Running a guild? <span style=\"color:var(--gold);font-weight:600\">Guild and Raid guild are further down</span></div>"}</div>")
    inner = f"<div style=\"display:flex;flex-direction:column;gap:18px\">{text}{seg}</div>" if phone else f"<div style=\"display:grid;grid-template-columns:minmax(0,1fr) auto;gap:24px;align-items:end\">{text}{seg}</div>"
    return f"<div class=\"band\"><div class=\"page-col\" style=\"padding-top:{28 if phone else 44}px;padding-bottom:{24 if phone else 36}px\">{inner}</div></div>"


def premium(key: str, phone: bool) -> str:
    grid = "display:flex;flex-direction:column;gap:16px" if phone else "display:grid;grid-template-columns:repeat(3,minmax(0,1fr));gap:20px;align-items:stretch"
    top = header_phone("out") if phone else header(state="out")
    body = (f"{top}{plans_band(phone)}<div class=\"page-col\" style=\"padding-top:24px;padding-bottom:40px\"><div style=\"{grid}\">{free_card()}{premium_card()}{coach_card()}</div>"
            f"<div style=\"margin-top:40px\"><span class=\"h2\">For guild leaders</span><p style=\"margin-top:8px;font-size:15px\">One officer pays, and the whole guild gets raids, sign-ups and attendance. Members never pay to sign up.</p>"
            f"</div></div>")
    return board(key, body, phone, "/premium plans row: Free, Premium, and a sample coach review in place of the example")



def faint_page() -> str:
    """A dimmed home hero behind the open menu, so the popover reads in place."""
    return (f"<div class=\"page-col\" style=\"padding-top:56px;opacity:.35;filter:saturate(.6)\"><span class=\"micro gold\">World of Warcraft: Forever</span>"
            f"<div class=\"h1\" style=\"margin:18px 0\">Play your class better.</div><p class=\"lede\">Pick your class for its best-in-slot list, talents and rotation.</p></div>")


def scene(inner: str, height: int, label: str) -> str:
    return (f"<div class=\"cap\">{label}</div><div style=\"position:relative;height:{height}px;overflow:hidden;background:var(--band);border-bottom:1px solid #1c2230\">{inner}</div>")


def header_board(width: int) -> str:
    right = "right:calc(max(0px, (100% - 1344px) / 2) + 48px)"
    s1 = scene(header(current="Tools", state="in") + faint_page(), 200, "1 · Signed in, closed. The right end is the account control: crest, name in class colour, spec and level, chevron.")
    s2 = scene(header(state="in", menu=account_menu("in", right)) + faint_page(), 860, "2 · Signed in, open: account head, your characters (switch), paste, then profile, account, plan, sign out.")
    s3 = scene(header(state="out", menu=account_menu("out", right)) + faint_page(), 520, "3 · Signed out, open: Sign in with Battle.net is the one filled button; email link and paste beneath.")
    return board("B", s1 + s2 + s3, False, f"Header and account menu at {width}")


def header_phone_board() -> str:
    sheet_in = (f"<div style=\"position:absolute;inset:0;background:rgba(0,0,0,.55)\"></div><div class=\"card sheet\" style=\"position:absolute;left:0;right:0;bottom:0;max-height:88%;overflow:hidden\">"
                f"<div class=\"grab\"><span></span></div>{account_menu_body("in")}</div>")
    sheet_out = (f"<div style=\"position:absolute;inset:0;background:rgba(0,0,0,.55)\"></div><div class=\"card sheet\" style=\"position:absolute;left:0;right:0;bottom:0\">"
                 f"<div class=\"grab\"><span></span></div>{account_menu_body("out")}</div>")
    doors = "".join(f"<div class=\"row link\" style=\"min-height:52px\"><span class=\"t lab\" style=\"text-transform:uppercase;letter-spacing:.1em;font-size:14px\">{d}</span><span class=\"chev\" style=\"display:flex\">{glyph("chev-down", 16)}</span></div>" for d in ["Tools", "Logs", "Guild", "Guides"])
    menu = (f"<div class=\"card\" style=\"position:absolute;left:0;right:0;top:56px;border-radius:0;padding-bottom:8px\">"
            f"<div style=\"padding:12px 18px\"><span class=\"field\" style=\"height:48px\"><span style=\"color:var(--gold);display:flex\">{glyph("search", 18)}</span>Character name</span></div>"
            f"{doors}<div class=\"row link\" style=\"min-height:52px\"><span class=\"t lab\" style=\"text-transform:uppercase;letter-spacing:.1em;font-size:14px\">Get set up</span></div></div>")
    s1 = scene(header_phone("in") + sheet_in, 844, "1 · Crest tapped: the account sheet")
    s2 = scene(header_phone("out") + sheet_out, 844, "2 · Signed out: Sign in tapped")
    s3 = scene(header_phone("in") + menu, 700, "3 · Menu: search, the four doors, Get set up. No Discord row.")
    return board("B", s1 + s2 + s3, True, "Header and account sheet at 390")


def footer_board(phone: bool) -> str:
    spacer = "<div class=\"page-col\" style=\"padding-top:40px;padding-bottom:60px;color:var(--muted);font-size:13px\">… page content …</div>"
    return board("B", spacer + footer(), phone, "Footer")



def compare_board() -> str:
    before = uri(HERE / "captures/account-signedin-1440-full.png")
    crop = (f"<div style=\"width:396px;height:816px;overflow:hidden;border-radius:6px;border:1px solid #262e40;position:relative\">"
            f"<img src=\"{before}\" alt=\"\" style=\"position:absolute;left:-948px;top:-675px;width:1440px\"></div>")
    after = f"<div style=\"width:440px;display:flex;flex-direction:column;gap:20px\">{plan_card()}{discord_card()}</div>"
    body = (f"<div style=\"display:flex;gap:64px;padding:40px 48px;align-items:flex-start;justify-content:center\">"
            f"<div><div class=\"micro\" style=\"margin-bottom:12px\">Before · production, 28 Oct</div>{crop}</div>"
            f"<div><div class=\"micro gold\" style=\"margin-bottom:12px\">After · settings rows, switches, button hierarchy</div>{after}</div></div>")
    return board("B", body, False, "Account Plan and Discord blocks, before and after")


def state_cell(label: str, html: str) -> str:
    return f"<div style=\"display:flex;flex-direction:column;gap:8px;align-items:flex-start\"><span class=\"note\">{label}</span>{html}</div>"


def components_board() -> str:
    spinner = "<span style=\"width:14px;height:14px;border-radius:999px;border:2px solid currentColor;border-right-color:transparent\"></span>"
    btn_rows = ""
    for kind, text in [("primary", "Subscribe · $40 a year"), ("secondary", "Manage billing"), ("ghost", "See plans"), ("danger", "Remove")]:
        cells = [state_cell("default", f"<span class=\"btn btn-{kind}\">{text}</span>"),
                 state_cell("hover", f"<span class=\"btn btn-{kind}\" style=\"filter:brightness(1.12)\">{text}</span>"),
                 state_cell("focus-visible", f"<span class=\"btn btn-{kind}\" style=\"outline:2px solid var(--gold);outline-offset:2px\">{text}</span>"),
                 state_cell("active", f"<span class=\"btn btn-{kind}\" style=\"filter:brightness(.92);transform:translateY(1px)\">{text}</span>"),
                 state_cell("disabled", f"<span class=\"btn btn-{kind}\" style=\"opacity:.45\">{text}</span>"),
                 state_cell("loading", f"<span class=\"btn btn-{kind}\" style=\"opacity:.85\">{spinner}{text}</span>")]
        btn_rows += f"<div style=\"display:grid;grid-template-columns:repeat(6,auto);gap:24px;justify-content:start\">{"".join(cells)}</div>"
    switches = "".join(state_cell(lab, f"<span class=\"switch {cls}\"></span>") for lab, cls in
                       [("off", ""), ("on", "on"), ("focus-visible", "on focus"), ("disabled", "dis"), ("busy (saving)", "on dis")])
    tags = "".join([f"<span class=\"tag tag-gold\">Founder</span>", f"<span class=\"tag tag-green\"><span class=\"dot\"></span>Connected</span>", f"<span class=\"tag tag-muted\">Sample</span>", f"<span class=\"tag tag-muted\">Main</span>"])
    links = (f"<span class=\"alink\">All 28 specs{glyph("chev-right", 16)}</span><span style=\"font-size:15px\">An inline link in prose reads like <span class=\"ilink\">the refund policy</span>, with a soft gold underline.</span>")
    tiles = "".join(f"<div class=\"stat\" style=\"width:150px\"><span class=\"n\">{n}</span><span class=\"micro\">{u}</span></div>" for n, u in [("07", "Days"), ("3", "Runs this month"), ("88", "Forever Score")])
    rows = (f"<div class=\"card\" style=\"width:460px\">{value_row("server", "Server sims", "On · 3 runs this month")}"
            f"<div class=\"row link hover\"><span class=\"itile sm\">{glyph("chart", 16)}</span><span class=\"t lab\">History charts (hover)</span><span class=\"val\">Open on your main character</span><span class=\"chev\" style=\"display:flex\">{glyph("chev-right", 16)}</span></div>"
            f"{switch_row("A server sim finishes", True)}{switch_row("Include private logs", False, "Off, private logs never post.", True)}"
            f"{value_row("bell", "Discord notifications", "Off", False, True)}</div>")
    def section(title: str, inner: str) -> str:
        return f"<section style=\"display:flex;flex-direction:column;gap:16px;padding:28px 0;border-top:1px solid #1c2230\"><span class=\"h3\">{title}</span>{inner}</section>"
    body = (f"<div class=\"page-col\" style=\"padding-top:12px;padding-bottom:48px\">"
            f"{section("Buttons", btn_rows)}{section("Switch", f"<div style=\"display:flex;gap:32px\">{switches}</div>")}"
            f"{section("Settings rows", rows)}{section("Tags", f"<div style=\"display:flex;gap:10px\">{tags}</div>")}"
            f"{section("Links", f"<div style=\"display:flex;flex-direction:column;gap:6px\">{links}</div>")}"
            f"{section("Stat tiles", f"<div style=\"display:flex;gap:12px\">{tiles}</div>")}"
            f"{section("Card head", f"<div class=\"card\" style=\"width:460px\"><div class=\"card-head\"><span class=\"itile\">{glyph("crown", 18)}</span><span class=\"grow\"><span class=\"micro\">Plan</span><div class=\"h2\" style=\"font-size:24px\">Premium</div></span><span class=\"tag tag-gold\">Founder</span></div></div>")}</div>")
    return board("B", body, False, "Component sheet: every state")


BOARDS = {
    "polish-header-1440": lambda: header_board(1440),
    "polish-header-2000": lambda: header_board(2000),
    "polish-header-phone": header_phone_board,
    "polish-footer-1440": lambda: footer_board(False),
    "polish-footer-2000": lambda: footer_board(False),
    "polish-footer-phone": lambda: footer_board(True),
    "polish-account-compare": compare_board,
    "polish-components": components_board,
}
for _key in DIRECTIONS:
    BOARDS[f"dir{_key}-home-1440"] = (lambda k: lambda: home(k, False))(_key)
    BOARDS[f"dir{_key}-home-phone"] = (lambda k: lambda: home(k, True))(_key)
    BOARDS[f"dir{_key}-account-1440"] = (lambda k: lambda: account(k, False))(_key)
    BOARDS[f"dir{_key}-account-phone"] = (lambda k: lambda: account(k, True))(_key)
    BOARDS[f"dir{_key}-premium-1440"] = (lambda k: lambda: premium(k, False))(_key)
    BOARDS[f"dir{_key}-premium-phone"] = (lambda k: lambda: premium(k, True))(_key)


def main() -> None:
    RENDERS.mkdir(exist_ok=True)
    names = sys.argv[1:] or list(BOARDS)
    for name in names:
        (RENDERS / f"{name}.html").write_text(BOARDS[name]())
        print(name)


if __name__ == "__main__":
    main()
