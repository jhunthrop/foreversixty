"""F12 universal search and the item page, F13 the talent planner."""
from w3kit import *  # noqa: F403
from w3kit import W3, TAB_CLEAR
from screens_home import tooltip_panel
from screens_more import ranks_of, tree_grid, LEVEL_POINTS, _icon
from model import slots
from kit import DATA

SHA = CLASS_COLOR["shaman"]


def kb(plat: str) -> str:
    keys = "".join(f"<span style=\"height:{42 if plat == 'ios' else 46}px;border-radius:{6 if plat == 'ios' else 8}px;background:{'#3a3d46' if plat == 'ios' else '#2a2f3c'}\"></span>" for _ in range(10))
    rows = "".join(f"<div style=\"display:grid;grid-template-columns:repeat(10,1fr);gap:5px;padding:0 4px\">{keys}</div>" for _ in range(3))
    return (f"<div style=\"position:absolute;left:0;right:0;bottom:0;height:{300 if plat == 'ios' else 290}px;background:{'#1c1d22' if plat == 'ios' else '#161a22'};z-index:14;padding-top:12px;display:flex;flex-direction:column;gap:11px\">{rows}"
            f"<div style=\"margin:0 4px;height:44px;border-radius:8px;background:{'#3a3d46' if plat == 'ios' else '#2a2f3c'}\"></div></div>")


def search_bar(plat: str, q: str = "", placeholder: str = "Search items, characters, guides, talents") -> str:
    top = DEVICES[plat]["top"]
    txt = f"<span class=\"t-body strong\">{q}</span>" if q else f"<span class=\"t-body muted\">{placeholder}</span>"
    clear = f"<span class=\"muted\" style=\"display:flex\">{glyph('x', 16)}</span>" if q else ""
    if plat == "ios":
        return (f"<div style=\"position:absolute;top:{top + 2}px;left:0;right:0;height:56px;display:flex;align-items:center;gap:10px;padding:0 16px;z-index:12;background:rgba(7,9,13,.96)\">"
                f"<span class=\"glass\" style=\"flex:1;height:44px;border-radius:22px;display:flex;align-items:center;gap:8px;padding:0 14px\"><span class=\"muted\" style=\"display:flex\">{glyph('search', 18)}</span>{txt}<span style=\"flex:1\"></span>{clear}</span>"
                f"<span class=\"t-body gold\" style=\"font-weight:600\">Cancel</span></div>")
    return (f"<div style=\"position:absolute;top:{top + 4}px;left:12px;right:12px;height:56px;display:flex;align-items:center;gap:12px;padding:0 16px;border-radius:28px;z-index:12;background:#1a1f2b\">"
            f"<span style=\"display:flex\">{glyph('arrow-left', 22)}</span>{txt}<span style=\"flex:1\"></span>{clear}</div>")


def scopes(active: int) -> str:
    return hs(["All", "Items", "Characters", "Guides", "Talents"], active)


def hs(items: list[str], active: int) -> str:
    return "<div class=\"hscroll\" style=\"margin-top:8px\">" + "".join(
        f"<span class=\"t-meta\" style=\"display:inline-flex;align-items:center;min-height:36px;padding:0 14px;border-radius:999px;border:1px solid {'var(--gold)' if i == active else 'var(--line)'};"
        f"{'background:rgba(229,185,85,.12);color:var(--gold);font-weight:700' if i == active else 'color:var(--nav)'}\">{t}</span>" for i, t in enumerate(items)) + "</div>"


def search_screen(plat: str, q: str, scope: int, body: str, keyboard: bool = True) -> str:
    top = DEVICES[plat]["top"] + 62
    return device(plat, search_bar(plat, q) + f"<div style=\"padding-top:{top}px\">{scopes(scope)}{body}</div>" + (kb(plat) if keyboard else ""), None)


def crow(name: str, cls: str, meta: str, you: bool = False, watch: bool = False) -> str:
    star = f"<span class=\"{'gold' if watch else 'muted'}\" style=\"display:flex;min-width:44px;justify-content:center\">{glyph('star-f' if watch else 'star-o', 20)}</span>"
    return (f"<div class=\"lrow\" style=\"min-height:60px\">{crest(cls, 36)}<div class=\"t\"><div class=\"t-body\" style=\"font-weight:700;color:{CLASS_COLOR[cls]}\">{name}{' ' + tag('You') if you else ''}</div><div class=\"t-meta muted\">{meta}</div></div>{star}</div>")


def search_empty(plat: str) -> str:
    watch = crow("Voltaris", "shaman", "Elemental Shaman · 60 · Hordebreakers", watch=True) + crow("Kraggor", "warrior", "Protection Warrior · 60 · Olympus XXVII", watch=True)
    rec = (f"<div class=\"lrow\" style=\"min-height:52px\">{item_icon(11623, 32)}<div class=\"t\">{qname(11623)}<div class=\"t-micro muted\">Item</div></div></div>"
           f"<div class=\"lrow\" style=\"min-height:52px\">{sp('spell_nature_lightning', 32)}<div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:600\">Elemental Shaman</div><div class=\"t-micro muted\">Guide</div></div></div>")
    jump = "".join(f"<div class=\"lrow\" style=\"min-height:52px\">{tile(g, 32)}<span class=\"t t-body strong\" style=\"font-weight:600\">{t}</span><span class=\"chev\">{glyph('chev-right', 18)}</span></div>" for g, t in [("shield", "Best in slot for Lightningbolltz"), ("book", "Elemental Shaman guide"), ("trophy", "Rankings")])
    body = label("For Lightningbolltz", first=True) + card(jump) + label("Watching") + card(watch) + label("Recent", extra="display:flex;justify-content:space-between") .replace("</div>", "<span class=\"gold\" style=\"text-transform:none;letter-spacing:0;font-size:13px\">Clear</span></div>", 1) + card(rec)
    return search_screen(plat, "", 0, body, keyboard=False)


def search_all(plat: str) -> str:
    tal = [t for t in DATA["trees"][0]["talents"] if t["name"].startswith("Elemental")][:3]
    trows = "".join(f"<div class=\"lrow\" style=\"min-height:56px\">{sp(t['icon'], 36)}<div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:600\">{t['name']}</div><div class=\"t-meta muted\">Elemental Shaman · Elemental · {t['max_rank']} {'rank' if t['max_rank'] == 1 else 'ranks'}</div></div></div>" for t in tal)
    g = f"<div class=\"lrow\" style=\"min-height:60px\">{sp('spell_nature_lightning', 36)}<div class=\"t\"><div class=\"t-body\" style=\"font-weight:700;color:{SHA}\">Elemental Shaman</div><div class=\"t-meta muted\">Guide · DPS</div></div><span class=\"chev\">{glyph('chev-right', 18)}</span></div>"
    ch = crow("Elementalx", "mage", "Mage · Level 58 · US Normal") + crow("Elemental", "druid", "Druid · Level 44 · EU Normal")
    body = label("Guides", first=True) + card(g) + label("Talents") + card(trows) + label("Characters") + card(ch)
    return search_screen(plat, "ele", 0, body)


def search_items(plat: str) -> str:
    ids = [16697, 16695, 16690, 16693, 16692]
    rows = "".join(f"<div class=\"lrow\" style=\"min-height:64px\">{item_icon(i, 40)}<div class=\"t\">{qname(i)}<div class=\"t-meta muted\">{item(i)['slot'].title()} · item level {item(i)['item_level']} · requires level {item(i)['required_level']}</div></div></div>" for i in ids)
    body = label("Items", first=True) + card(rows) + note("Items from the game's own data. Tap one for its tooltip, where it drops and which best in slot lists use it.")
    return search_screen(plat, "devout", 1, body)


def search_none(plat: str) -> str:
    body = card("<div class=\"t-body strong\" style=\"font-size:17px\">Nothing matches “xqz”.</div><div class=\"t-meta muted\" style=\"margin-top:4px\">Check the spelling. Characters need at least two letters and a connection.</div>", margin="14px 16px 0", pad="16px")
    body += card(offline_strip().replace("Offline. Showing what you had at Thu 9:41 pm.", "Offline. Items, guides and talents still search; characters need a connection."), margin="10px 16px 0")
    return search_screen(plat, "xqz", 0, body)


CAPE = 11623


def item_head() -> str:
    it = item(CAPE)
    return (f"<div class=\"pad\" style=\"display:flex;gap:14px;align-items:center;margin-top:4px\">{item_icon(CAPE, 64)}<div style=\"min-width:0\"><div class=\"cond t-h2\" style=\"color:{Q_TEXT[it['quality']]};text-transform:none\">{it['name']}</div>"
            f"<div class=\"t-meta muted\">{Q_WORD[it['quality']]} · Back · item level {it['item_level']}</div></div></div>")


def item_page(plat: str, part: int = 1) -> str:
    top = content_top(plat)
    left = back_btn(plat, "Search") + (f"<span class=\"cond\" style=\"font-size:22px\">Spritecaster Cape</span>" if plat == "and" else "")
    bar = top_bar(plat, f"<span style=\"display:flex;align-items:center;gap:4px\">{left}</span>", "")
    mine = card(f"<div style=\"padding:12px 16px;display:flex;gap:10px;align-items:center\">{crest('shaman', 36)}<div style=\"flex:1\"><div class=\"t-body strong\" style=\"font-weight:700\">Your Back pick</div><div class=\"t-meta muted\">Elemental · 60 · Phase 1 raid · Horde</div></div>{tag('Yours')}</div>", quiet=False, margin="14px 16px 0")
    tip = f"<div style=\"padding:0 16px;margin-top:14px\">{tooltip_panel(CAPE, 'Back')}</div>"
    drops = W3["cape"]["drops"]
    drow = "".join(f"<div class=\"lrow\" style=\"min-height:60px\">{tile('swords', 36)}<div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:600\">{d['boss']}</div><div class=\"t-meta muted\">{d['zone']} · Dungeon</div></div></div>" for d in drops)
    if part == 1:
        inner = item_head() + mine + tip + label("Where it comes from") + card(drow) + note("Drop chances show where the game's data has them.")
    else:
        uses = W3["cape"]["uses"]
        def urow(u: dict) -> str:
            place = "Pick" if u["place"] == 0 else f"Runner-up {u['place']}"
            band = f"{u['band']} · Phase 1 raid" if u["preset"] == "raid" else f"{u['band']}"
            return (f"<div class=\"lrow\" style=\"min-height:60px\">{spec_mark(u['spec'], 36)}<div class=\"t\"><div class=\"t-body\" style=\"font-weight:700;color:{CLASS_COLOR[u['class']]}\">{u['name']} {u['class'].title()}</div>"
                    f"<div class=\"t-meta muted\">{band}</div></div>{tag(place, 'green' if u['place'] == 0 else 'muted')}<span class=\"chev\">{glyph('chev-right', 18)}</span></div>")
        shown = uses[:6]
        inner = (hs(["Item", "Source", "Best in slot for"], 2) + label("Best in slot for", extra="margin-top:18px") + card("".join(urow(u) for u in shown))
                 + card(f"<div class=\"lrow\" style=\"min-height:52px\"><span class=\"t t-body gold\" style=\"font-weight:600\">{len(uses) - len(shown)} more specs</span><span class=\"chev gold\">{glyph('chev-down', 18)}</span></div>", margin="8px 16px 0")
                 + note("Each row opens that list. The best placement of each spec is shown."))
    return device(plat, bar + f"<div style=\"padding-top:{top}px\">{inner}{TAB_CLEAR}</div>", None)


# ---------------- F13 ----------------

def build_ranks(kind: str = "guide") -> list[list[int]]:
    g = ranks_of(DATA["bands"]["raid"]["talents"])
    r = [list(t) for t in g]
    if kind == "mid":
        r[0][15] = 0            # Elemental Mastery
        r[0][14] = 3            # Elemental Fury 5 -> 3
    if kind == "low":
        for i, t in enumerate(DATA["trees"][0]["talents"]):
            if t["tier"] >= 2 or t["name"] in ("Reverberation", "Call of Flame"):
                r[0][i] = 0
    return r


def stepper(ic: str, name: str, rank: int, mx: int, text: str, tone: str = "", locked: bool = False) -> str:
    btn = lambda g, dis: (f"<span style=\"width:48px;height:48px;border-radius:24px;display:grid;place-items:center;border:1px solid var(--card-b);background:rgba(255,255,255,.05);{'opacity:.4' if dis else ''}\">{glyph(g, 22)}</span>")  # noqa: E731
    return (f"<div style=\"position:absolute;left:0;right:0;bottom:92px;z-index:16;background:#10141d;border-top:1px solid var(--line);padding:10px 16px;display:flex;gap:12px;align-items:center;height:96px\">{sp(ic, 44)}"
            f"<div style=\"flex:1;min-width:0\"><div class=\"t-body strong\" style=\"font-weight:700\">{name} <span class=\"mono t-meta muted\">{rank}/{mx}</span></div><div class=\"t-meta {tone or 'muted'}\" style=\"line-height:16px\">{text}</div></div>"
            f"{btn('minus', rank == 0)}{btn('plus', locked or rank == mx)}</div>")


def action_bar(primary: str = "Share", extra: str = "") -> str:
    return (f"<div style=\"position:absolute;left:0;right:0;bottom:0;z-index:16;background:#0d111a;border-top:1px solid var(--line-soft);padding:10px 16px 34px;display:flex;gap:10px\">"
            f"<span class=\"btn btn-secondary\" style=\"flex:1\">Reset</span><span class=\"btn btn-primary\" style=\"flex:2\">{glyph('share', 16)}{primary}</span></div>")


def editor(plat: str, kind: str = "mid", overlay: str = "") -> str:
    ranks = build_ranks(kind)
    totals = [sum(t) for t in ranks]
    spent = sum(totals)
    left = LEVEL_POINTS - spent
    top = content_top(plat)
    names = [t["name"] for t in DATA["trees"]]
    tabs = "".join(f"<span class=\"t-meta\" style=\"flex:1;text-align:center;padding:9px 0;{'color:var(--strong);font-weight:700;box-shadow:inset 0 -2px 0 var(--gold)' if i == 0 else 'color:var(--nav)'}\">{n} <span class=\"mono\">{p}</span></span>" for i, (n, p) in enumerate(zip(names, totals)))
    chips = (f"<div style=\"display:flex;align-items:center;gap:8px;padding:6px 16px 0\">{crest('shaman', 30)}{chip('<b>Orc</b>')}{chip('Level 60')}<span style=\"flex:1\"></span>"
             f"<span class=\"mono t-body\" style=\"text-align:right;line-height:16px\"><b class=\"{'gold' if left else 'kill'}\" style=\"font-size:20px\">{left}</b> <span class=\"muted t-meta\">left</span></span></div>")
    talent = {"mid": ("Elemental Fury", "spell_fire_volcano", 3, 5, "Increases the critical strike damage bonus of your Searing and Magma Totems and your Fire, Frost, and Nature spells by 60%.", "", False),
              "low": ("Lightning Overload", "spell_nature_lightningoverload", 0, 3, f"Needs 20 points in Elemental. You have {totals[0]}.", "wipe", True),
              "guide": ("Lava Burst", "spell_shaman_lavaburst", 1, 1, "You hurl molten lava at the target, dealing 114 to 145 Fire damage. If your Flame Shock is on the target, Lava Burst deals 20% increased damage.", "", False)}[kind]
    differs = set()
    body = (chips + f"<div style=\"display:flex;margin:10px 16px 0;border-bottom:1px solid var(--line)\">{tabs}</div>" + tree_grid(ranks[0], differs))
    bar = top_bar(plat, f"<span style=\"display:flex;align-items:center;gap:4px\">{back_btn(plat, 'Talents')}" + ("<span class=\"cond\" style=\"font-size:22px\">Plan a build</span>" if plat == "and" else "") + "</span>",
                  icon_btn(plat, "undo", "Undo"))
    ins = stepper(talent[1], talent[0], talent[2], talent[3], talent[4], talent[5], talent[6])
    return device(plat, bar + f"<div style=\"padding-top:{top}px\">{body}</div>" + ins + action_bar(), None, overlay=overlay)


def confirm_sheet(plat: str) -> str:
    inner = (f"<div style=\"padding:0 20px\"><div class=\"t-body\">Sharing posts the following to a link anyone can open. Once created, that link cannot be edited:</div>"
             f"<div class=\"t-body\" style=\"margin-top:10px;line-height:26px\">· Class and race<br>· Talent order<br>· Title: “Elemental raid”</div>"
             f"<div class=\"t-micro micro\" style=\"margin-top:16px\">Title (optional)</div><div style=\"margin-top:6px;display:flex;align-items:center;min-height:48px;padding:0 14px;border-radius:8px;border:1px solid var(--gold);background:rgba(0,0,0,.35)\" class=\"t-body\">Elemental raid<span style=\"flex:1\"></span><span class=\"mono muted t-meta\">14/60</span></div>"
             f"<div style=\"display:flex;flex-direction:column;gap:10px;margin-top:18px\"><span class=\"btn btn-primary full\">Share anyway</span><span class=\"btn btn-secondary full\">Cancel</span></div>"
             f"<div class=\"t-meta muted\" style=\"margin-top:12px\">Signed in, it also appears in Your builds.</div></div>")
    return sheet(plat, inner, 520, "Share this build")


def share_sheet(plat: str, title: str = "Orc Shaman · Elemental 30 / Enhancement 3 / Restoration 18", url: str = "foreversixty.gg/b/k7q2m9xa", note_: str = "") -> str:
    targets = "".join(f"<div style=\"display:flex;flex-direction:column;align-items:center;gap:6px;width:72px\"><span style=\"width:56px;height:56px;border-radius:{14 if plat == 'ios' else 28}px;background:{c}\"></span><span style=\"font-size:12px;color:#e9e4d8\">{n}</span></div>" for n, c in [("Messages", "#2f8f4f"), ("Discord", "#5865f2"), ("Mail", "#2f6fd0"), ("Copy link", "#3a3f4c")])
    inner = (f"<div style=\"padding:0 16px\"><div style=\"display:flex;gap:12px;align-items:center;padding:12px;border-radius:14px;background:rgba(255,255,255,.05)\"><div style=\"width:84px;height:44px;border-radius:6px;background:var(--band);border:1px solid var(--card-b);display:grid;place-items:center\">{glyph('tree', 22)}</div>"
             f"<div style=\"min-width:0\"><div class=\"t-meta strong\" style=\"font-weight:600;line-height:16px\">{title}</div><div class=\"t-micro muted\" style=\"letter-spacing:0;text-transform:none\">{url}</div></div></div>"
             f"<div style=\"display:flex;gap:6px;justify-content:space-between;margin-top:18px\">{targets}</div>{note_}</div>")
    return sheet(plat, inner, 360, "")


def viewer(plat: str) -> str:
    g = build_ranks("guide")
    tot = " / ".join(str(sum(t)) for t in g)
    top = content_top(plat)
    left = back_btn(plat, "Home") + (f"<span class=\"cond\" style=\"font-size:22px\">Shared build</span>" if plat == "and" else "")
    bar = top_bar(plat, f"<span style=\"display:flex;align-items:center;gap:4px\">{left}</span>", icon_btn(plat, "share", "Share"))
    head_ = "" if plat == "and" else "<div class=\"cond t-title pad\" style=\"margin-top:4px\">Shared build</div>"
    tabs = "".join(f"<span class=\"t-meta\" style=\"flex:1;text-align:center;padding:9px 0;{'color:var(--strong);font-weight:700;box-shadow:inset 0 -2px 0 var(--gold)' if i == 0 else 'color:var(--nav)'}\">{t['name']} <span class=\"mono\">{sum(g[i])}</span></span>" for i, t in enumerate(DATA["trees"]))
    inner = (f"<div class=\"pad\" style=\"display:flex;gap:12px;align-items:center\">{crest('shaman', 44)}<div><div class=\"t-body strong\" style=\"font-weight:700\">Orc Shaman <span class=\"mono\">{tot}</span></div><div class=\"t-meta muted\">Elemental raid · Saved 4 Nov · 12 views</div></div></div>"
             f"<div style=\"display:flex;gap:10px;padding:12px 16px 0\"><span class=\"btn btn-primary\" style=\"flex:1\">Edit a copy</span><span class=\"btn btn-secondary\" style=\"flex:1\">Compare with mine</span></div>"
             f"<div style=\"display:flex;margin:12px 16px 0;border-bottom:1px solid var(--line)\">{tabs}</div>{tree_grid(g[0], set())}"
             f"<div class=\"t-meta muted pad\" style=\"margin-top:8px\">Opened from foreversixty.gg/b/k7q2m9xa. The data is the game build this was made on.</div>")
    return device(plat, bar + f"<div style=\"padding-top:{top}px\">{head_}{inner}</div>", "Home")


def my_builds(plat: str) -> str:
    def r(title: str, meta: str, views: str, pressed: bool = False) -> str:
        return f"<div class=\"lrow{' pressed' if pressed else ''}\" style=\"min-height:68px\">{crest('shaman', 36)}<div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:600\">{title}</div><div class=\"t-meta muted\">{meta}</div></div><span class=\"mono t-meta muted\">{views} views</span><span class=\"chev\">{glyph('chev-right', 18)}</span></div>"
    rows = r("Elemental raid", "Orc Shaman · 30 / 3 / 18 · 4 Nov", "12") + r("Orc Shaman 31 / 20 / 0", "Orc Shaman · 29 Oct", "3") + r("Orc Warrior 31 / 20 / 0", "Orc Warrior · 21 Oct", "0")
    sw = ("<div style=\"position:absolute;right:0;top:0;bottom:0;width:104px;background:var(--wipe);display:grid;place-items:center;color:#0b0e14;font-weight:700\" class=\"t-meta\">Remove</div>")
    inner = card(rows, margin="14px 16px 0") + note("Remove only takes a build off your list. The link keeps working for anyone who has it.") + label("Plan a new build") + card(f"<div class=\"lrow\" style=\"min-height:56px\">{tile('plus')}<span class=\"t t-body strong\" style=\"font-weight:600\">Shaman, level 60</span><span class=\"chev\">{glyph('chev-right', 18)}</span></div>")
    return pushed(plat, "Talents", "Your builds", inner + TAB_CLEAR, "Home")
