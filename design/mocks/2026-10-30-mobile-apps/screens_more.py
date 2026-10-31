"""F5 sims, F6 talents, F7 notifications (lock screens, settings, the in-context ask), F8 You and Premium."""
from kit import (DATA, HERE, ME, back_btn, content_top, device, glyph, item_icon, sheet, spec_icon, top_bar, uri)
from model import SLOT_ORDER, slots

APP_ICON = uri(HERE.parents[1] / "logo/export/foreversixty-icon-64.png")


def tile(g: str, size: int = 36) -> str:
    return f"<span style=\"flex:none;width:{size}px;height:{size}px;border-radius:8px;display:grid;place-items:center;background:var(--tile);color:var(--tile-ink)\">{glyph(g, size // 2)}</span>"


def pushed(plat: str, back: str, title: str, inner: str, tab: str, overlay: str = "", xl: bool = False, right: str = "") -> str:
    top = content_top(plat)
    left = back_btn(plat, back) + (f"<span class=\"cond\" style=\"font-size:22px\">{title}</span>" if plat == "and" else "")
    head_ = "" if plat == "and" else f"<div class=\"cond t-title pad\" style=\"margin-top:4px\">{title}</div>"
    return device(plat, top_bar(plat, f"<span style=\"display:flex;align-items:center;gap:4px\">{left}</span>", right) + f"<div style=\"padding-top:{top}px\">{head_}{inner}</div>", tab, overlay=overlay, xl=xl)


# ---------- F5 sims ----------

def sims(plat: str, premium: bool = True) -> str:
    running = (f"<div class=\"lrow\" style=\"align-items:flex-start\">{spec_icon('spell_nature_lightning', 36)}<div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:600\">Elemental · Top gear, 64 combinations</div>"
               f"<div class=\"t-meta muted\"><span class=\"gold\">Running</span> · takes about 4 minutes</div><div class=\"bar8\" style=\"margin-top:8px\"><i style=\"width:62%\"></i></div>"
               f"<div class=\"t-micro muted mono\" style=\"margin-top:4px\">40 of 64 combinations</div></div></div>")
    done = (f"<div class=\"lrow\">{spec_icon('spell_nature_lightning', 36)}<div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:600;display:flex;gap:8px;align-items:center\"><span style=\"width:8px;height:8px;border-radius:4px;background:var(--gold)\" aria-label=\"Not opened yet\"></span>Elemental · Phase 1 raid set</div>"
            f"<div class=\"t-meta muted\">Finished 2 h ago</div></div><span class=\"mono t-body strong\">412 DPS</span><span class=\"chev\">{glyph('chev-right', 16)}</span></div>")
    failed = (f"<div class=\"lrow\" style=\"align-items:flex-start\">{spec_icon('spell_nature_lightning', 36)}<div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:600\">Elemental · Stat weights</div>"
              f"<div class=\"t-meta wipe\">Failed: the engine stopped before the first result.</div><div class=\"t-meta muted\">Ended yesterday</div></div><span class=\"btn btn-secondary\" style=\"min-height:40px;padding:0 12px\">Run again</span></div>")
    mine = "".join(f"<div class=\"lrow\">{spec_icon(ic, 32)}<div class=\"t\"><div class=\"t-body\" style=\"font-weight:600;color:var(--strong)\">{h}</div><div class=\"t-meta muted\">{d}</div></div><span class=\"mono t-body strong\">{dps}</span></div>"
                   for ic, h, d, dps in [("spell_nature_lightning", "Elemental · Bare 60", "Saved from the browser · 26 Oct", "149 DPS"),
                                         ("ability_warrior_innerrage", "Grimtusk · Fury 40 to 49", "Saved from the browser · 22 Oct", "88 DPS")])
    if premium:
        inner = (f"<div class=\"t-micro micro pad\" style=\"margin-top:12px\">Server runs</div><div class=\"card quiet\" style=\"margin:8px 16px 0\">{running}{done}{failed}</div>"
                 f"<div class=\"t-micro micro pad\" style=\"margin-top:22px\">Your sims</div><div class=\"card quiet\" style=\"margin:8px 16px 0\">{mine}</div>"
                 f"<div class=\"t-meta muted pad\" style=\"margin-top:12px\">Start a sim on foreversixty.gg/sim. Server runs keep going after you close the app.</div>")
    else:
        inner = (f"<div class=\"t-micro micro pad\" style=\"margin-top:12px\">Your sims</div><div class=\"card quiet\" style=\"margin:8px 16px 0\">{mine}</div>"
                 f"<div class=\"t-meta muted pad\" style=\"margin-top:12px\">Start a sim on foreversixty.gg/sim and it shows here.</div>")
    return pushed(plat, "Home", "Sims", inner, "Home")


def sim_result(plat: str, premium: bool = True, link_region: bool = True) -> str:
    icons = "".join(item_icon(slots()[k]["item_id"], 40) for k in SLOT_ORDER)
    action = ""
    if premium:
        action = "<span class=\"btn btn-secondary full\" style=\"margin-top:16px\">Run again on the server</span>"
    elif link_region:
        action = "<div class=\"t-meta muted\" style=\"margin-top:16px\">Server sims are part of Premium. <span class=\"gold\" style=\"text-decoration:underline\">See plans</span></div>"
    inner = (f"<div class=\"pad\"><div class=\"t-meta muted\" style=\"margin-top:4px\">Elemental Shaman · level 60 · Phase 1 raid buffs · Finished 2 h ago</div>"
             f"<div style=\"display:flex;align-items:baseline;gap:8px;margin-top:16px\"><span class=\"cond\" style=\"font-size:52px;line-height:1\">412.3</span><span class=\"t-body muted\">DPS</span></div>"
             f"<div class=\"t-meta muted mono\">± 1.8 · 10,000 iterations · engine 1.60.1.70532</div>"
             f"<div class=\"t-meta\" style=\"margin-top:8px\"><span class=\"mono strong\">453.6 DPS</span> at best in slot for your level range</div>"
             f"<div class=\"t-micro micro\" style=\"margin-top:20px\">The gear it ran</div><div style=\"display:grid;grid-template-columns:repeat(6,1fr);gap:8px;margin-top:8px\">{icons}</div>"
             f"{action}<span class=\"btn btn-ghost\" style=\"margin-top:8px;padding:0\">Full breakdown on foreversixty.gg{glyph('external', 14)}</span></div>")
    return pushed(plat, "Sims", "Phase 1 raid set", inner, "Home")


# ---------- F6 talents ----------
# The planner's rules (web/src/lib/planner): a tier opens at 5 points per tier in its tree; a cell's border is
# green when it holds points, gold when it can take one, the line colour when locked (face at 40 %); a
# prerequisite line runs edge to edge, gold once met, over a 5 px --bg halo. The count is band-compare.ts's
# diffCount: your points spent where the guide spends fewer, and "Matches" only when no cell differs.

ICON, PX, PY = 46, 92, 64
LEVEL_POINTS = 51


def ranks_of(code: str) -> list[list[int]]:
    trees = code.split("-")
    return [[int(c, 36) for c in tree] + [0] * (len(DATA["trees"][i]["talents"]) - len(tree)) for i, tree in enumerate(trees)]


def your_build(guide: list[list[int]], state: str) -> list[list[int]]:
    if state == "same":
        return [list(t) for t in guide]
    yours = [list(t) for t in guide]
    yours[0][3] += 2  # Reverberation 0 -> 2, where the guide spends none
    yours[0][5] -= 2  # Call of Flame 3 -> 1; nothing depends on it, so the build stays legal
    return yours


def diff_count(yours: list[list[int]], guide: list[list[int]]) -> tuple[int, set[tuple[int, int]]]:
    count, cells = 0, set()
    for t, (ys, gs) in enumerate(zip(yours, guide)):
        for i, (y, g) in enumerate(zip(ys, gs)):
            count += max(y - g, 0)
            if y != g:
                cells.add((t, i))
    return count, cells


def tree_grid(ranks: list[int], differs: set[int]) -> str:
    talents = DATA["trees"][0]["talents"]
    spent = sum(ranks)
    by_id = {t["id"]: (i, t) for i, t in enumerate(talents)}
    left = (402 - (3 * PX + ICON)) // 2
    pos = lambda t: (left + t["column"] * PX, t["tier"] * PY)  # noqa: E731
    lines = ""
    for i, t in enumerate(talents):
        if t["prereq_talent_id"] is None:
            continue
        j, pre = by_id[t["prereq_talent_id"]]
        (ax, ay), (bx, by) = pos(pre), pos(t)
        h = ICON / 2
        if pre["column"] == t["column"]:
            d = f"M {ax + h} {ay + ICON} V {by}"
        else:
            side = 1 if t["column"] > pre["column"] else -1
            d = f"M {ax + h} {ay + ICON} V {by + h} H {bx + h - side * h}"
        met = ranks[j] >= (t["prereq_rank"] or 0)
        lines += (f"<path d=\"{d}\" fill=\"none\" stroke=\"#07090d\" stroke-width=\"5\" stroke-linecap=\"round\"/>"
                  f"<path d=\"{d}\" fill=\"none\" stroke=\"{'#e5b955' if met else '#262e40'}\" stroke-width=\"2\" stroke-linecap=\"round\"/>")
    cells = ""
    for i, t in enumerate(talents):
        rank = ranks[i]
        prereq_ok = t["prereq_talent_id"] is None or ranks[by_id[t["prereq_talent_id"]][0]] >= (t["prereq_rank"] or 0)
        available = spent >= t["tier"] * 5 and prereq_ok
        if rank:
            border, pill, face = "#7fd48a", ("#7fd48a" if rank >= t["max_rank"] else "#e9e4d8"), ""
        elif available:
            border, pill, face = "#e5b955", "#e5b955", ""
        else:
            border, pill, face = "#262e40", "#9a9484", "opacity:.4;"
        ring = "outline:2px dashed #f0736b;outline-offset:3px;" if i in differs else ""
        x, y = pos(t)
        cells += (f"<span style=\"position:absolute;left:{x}px;top:{y}px\"><img src=\"{_icon(t['icon'])}\" alt=\"\" style=\"display:block;width:{ICON}px;height:{ICON}px;border-radius:8px;border:2px solid {border};{face}{ring}\">"
                  f"<span class=\"mono\" style=\"position:absolute;right:-7px;bottom:-7px;font-size:11px;padding:1px 5px;border-radius:6px;background:#07090d;border:1px solid {border};color:{pill}\">{rank}/{t['max_rank']}</span></span>")
    tiers = max(t["tier"] for t in talents) + 1
    height = (tiers - 1) * PY + ICON + 10
    return (f"<div style=\"position:relative;height:{height}px;margin-top:18px\"><svg width=\"402\" height=\"{height}\" style=\"position:absolute;left:0;top:0\" aria-hidden=\"true\">{lines}</svg>{cells}</div>")


def _icon(stem: str) -> str:
    from kit import asset
    return asset(stem)


def talents(plat: str, state: str = "same") -> str:
    guide = ranks_of(DATA["bands"]["raid"]["talents"])
    yours = your_build(guide, state)
    count, cells = diff_count(yours, guide)
    totals = [sum(t) for t in yours]
    spent = sum(totals)
    unspent = LEVEL_POINTS - spent
    if not cells:
        status = "<span class=\"tag tag-green\">Matches</span><span class=\"t-meta muted\">Your build is the Elemental 60 guide build.</span>"
    else:
        status = f"<span class=\"tag tag-gold\">Differs</span><span class=\"t-meta muted\"><span class=\"mono strong\">{count}</span> of your {spent} points differ from the 60 build.</span>"
    unspent_line = f"<div class=\"t-meta muted pad\">{unspent} points unspent.</div>" if unspent > 0 else ""
    seg = ("<div class=\"seg4\" style=\"grid-template-columns:1fr 1fr;margin:14px 16px 0\"><span class=\"t-seg\">Guide build</span><span class=\"t-seg on-no\" style=\"box-shadow:inset 0 -2px 0 var(--gold);background:rgba(229,185,85,.12)\">Your build</span></div>")
    names = [t["name"] for t in DATA["trees"]]
    tabs = "".join(f"<span class=\"t-meta\" style=\"flex:1;text-align:center;padding:10px 0;{'color:var(--strong);font-weight:700;box-shadow:inset 0 -2px 0 var(--gold)' if i == 0 else 'color:var(--nav)'}\">{n} <span class=\"mono\">{p}</span></span>"
                   for i, (n, p) in enumerate(zip(names, totals)))
    legend = "" if not cells else "<div class=\"t-meta muted pad\" style=\"margin-top:6px\">Dashed: your rank differs from the guide. Tap a talent for both ranks and its text.</div>"
    differs = {i for t, i in cells if t == 0}
    inner = (f"<div class=\"pad\" style=\"display:flex;flex-wrap:wrap;align-items:center;gap:8px;margin-top:6px\">{status}</div>{unspent_line}{seg}"
             f"<div style=\"display:flex;margin:12px 16px 0;border-bottom:1px solid var(--line)\">{tabs}</div>{tree_grid(yours[0], differs)}{legend}"
             f"<div class=\"t-meta muted pad\" style=\"margin-top:8px;padding-bottom:120px\">Change talents in the game; the addon sends them on your next logout.</div>")
    return pushed(plat, "Home", "Talents", inner, "Home")


# ---------- F7 notifications ----------

def ios_lock(cards: str, label: str = "") -> str:
    body = ("<div style=\"position:absolute;inset:0;background:radial-gradient(500px 420px at 70% 15%,rgba(214,110,40,.35),transparent 70%),linear-gradient(180deg,#1b130c,#07090d)\"></div>"
            "<div style=\"position:relative;text-align:center;padding-top:78px;color:#fff;font-family:-apple-system,'SF Pro Display',sans-serif\"><div style=\"font-size:20px;font-weight:600;opacity:.85\">Thursday 10 December</div>"
            "<div style=\"font-size:104px;font-weight:700;line-height:1;margin-top:4px;letter-spacing:-2px\">7:00</div></div>"
            f"<div style=\"position:absolute;left:10px;right:10px;bottom:120px;display:flex;flex-direction:column;gap:8px\">{cards}</div>")
    return device("ios", body, None)


def ios_note(title: str, body: str, when: str = "now", sens: bool = False, big: bool = False, attach: str = "", sub: str = "Olympus XXVII") -> str:
    """iOS draws the subtitle (the guild) on its own bold line under the title."""
    head_ = f"<span style=\"font-size:13px;opacity:.75;white-space:nowrap\">{'Time Sensitive · ' if sens else ''}{when}</span>"
    att = f"<img src=\"{attach}\" style=\"width:38px;height:38px;border-radius:8px;flex:none\" alt=\"\">" if attach else ""
    sub_html = f"<div style=\"font-size:15px;font-weight:600;opacity:.9\">{sub}</div>" if sub else ""
    return (f"<div class=\"glass\" style=\"border-radius:24px;padding:12px 14px;display:flex;gap:10px;align-items:flex-start;color:#fff;font-family:-apple-system,'SF Pro Text',sans-serif;{'padding:16px' if big else ''}\">"
            f"<img src=\"{APP_ICON}\" style=\"width:38px;height:38px;border-radius:9px;flex:none\" alt=\"\"><div style=\"flex:1;min-width:0\"><div style=\"display:flex;justify-content:space-between;gap:8px\"><b style=\"font-size:15px\">{title}</b>{head_}</div>"
            f"{sub_html}<div style=\"font-size:15px;line-height:20px;margin-top:1px\">{body}</div></div>{att}</div>")


def ios_actions(names: list[str]) -> str:
    rows = "".join(f"<div style=\"min-height:50px;display:flex;align-items:center;padding:0 18px;border-top:{'0' if i == 0 else '1px solid rgba(255,255,255,.1)'};font:400 17px -apple-system;color:#fff\">{n}</div>" for i, n in enumerate(names))
    return f"<div class=\"glass\" style=\"border-radius:22px;overflow:hidden;margin-top:8px\">{rows}</div>"


def and_shade(cards: str) -> str:
    qs = ("<div style=\"padding:52px 20px 8px;color:#e9e4d8;font:500 14px Roboto\"><div style=\"display:flex;justify-content:space-between\"><span style=\"font-size:16px\">Thu, Dec 10</span></div>"
          "<div style=\"display:grid;grid-template-columns:repeat(4,1fr);gap:10px;margin-top:14px\">" + "".join("<span style=\"height:52px;border-radius:26px;background:#232836\"></span>" for _ in range(4)) + "</div></div>")
    body = f"<div style=\"position:absolute;inset:0;background:#0b0d12\"></div><div style=\"position:relative\">{qs}<div style=\"display:flex;flex-direction:column;gap:3px;padding:12px 12px 0\">{cards}</div></div>"
    return device("and", body, None)


def and_note(title: str, body: str, sub: str = "Olympus XXVII", actions: list[str] | None = None, top: bool = False, bottom: bool = False, attach: str = "") -> str:
    r = f"{'24px 24px' if top else '6px 6px'} {'24px 24px' if bottom else '6px 6px'}"
    acts = ""
    if actions:
        acts = "<div style=\"display:flex;gap:8px;margin-top:12px;padding-left:44px\">" + "".join(f"<span style=\"min-height:40px;padding:0 16px;border-radius:20px;background:#2a3040;display:flex;align-items:center;font:500 14px Roboto;color:#f5d27a\">{a}</span>" for a in actions) + "</div>"
    small = f"<span style=\"width:32px;height:32px;border-radius:16px;background:#e5b955;display:grid;place-items:center;flex:none\"><img src=\"{APP_ICON}\" style=\"width:24px;height:24px;border-radius:12px;filter:grayscale(1) brightness(.2)\" alt=\"\"></span>"
    att = f"<img src=\"{attach}\" style=\"width:44px;height:44px;border-radius:8px;flex:none\" alt=\"\">" if attach else ""
    return (f"<div style=\"background:#1a1f2b;border-radius:{r};padding:14px 16px;font-family:Roboto,sans-serif;color:#e9e4d8\"><div style=\"display:flex;gap:12px;align-items:flex-start\">{small}"
            f"<div style=\"flex:1;min-width:0\"><div style=\"font-size:12px;color:#b9b3a4\">Forever Sixty · {sub} · now</div><div style=\"font-size:16px;font-weight:500;margin-top:2px\">{title}</div>"
            f"<div style=\"font-size:14px;color:#c9c4b8;margin-top:2px;line-height:20px\">{body}</div></div>{att}</div>{acts}</div>")


STARTING = ("Invites for Onyxia's Lair in 1 hour", "Invites 7:45 pm, pull 8:00 pm. You are In as DPS on Lightningbolltz. 38 of 40 signed up.")
POSTED = ("Barrow Deeps · Saturday 7 pm", "Sign-ups are open, 11 signed up so far. Sign up Lightningbolltz?")
NIGHT = ("Your raid night: Onyxia's Lair", "1 kill, no wipes. Loot: Sapphiron Drape. A new best on Onyxia. Coach notes are ready.")
MOVED_IN = ("You are In for Onyxia's Lair tonight", "Kraggor moved you from the Bench. Invites 7:45 pm, pull 8:00 pm, as DPS on Lightningbolltz.")
MOVED_BENCH = ("You are on the Bench for Onyxia's Lair tonight", "Kraggor moved you from In. An officer may still call you in before the pull at 8:00 pm.")


def lock_board_ios() -> list[str]:
    from kit import asset
    a = ios_lock(ios_note(*STARTING, sens=True) + ios_note(*POSTED, when="12m") + ios_note("Your Elemental sim finished", "412 DPS · Phase 1 raid set. Open the result.", when="2h", sub=""))
    b = ios_lock(ios_note(*POSTED, when="12m", big=True) + ios_actions(["In", "Late", "Out"]))
    c = ios_lock(ios_note("You are In for Barrow Deeps", "Saturday 7 pm, as DPS on Lightningbolltz. 12 signed up.", when="now")
                 + ios_note(*NIGHT, when="7d", attach=asset("inv_misc_cape_16")))
    return [a, b, c]


def lock_board_and() -> list[str]:
    from kit import asset
    a = and_shade(and_note(*POSTED, actions=["In", "Late", "Out"], top=True) + and_note(*STARTING)
                  + and_note("Your Elemental sim finished", "412 DPS · Phase 1 raid set. Open the result.", sub="Sims", bottom=True))
    b = and_shade(and_note("You are In for Barrow Deeps", "Saturday 7 pm, as DPS on Lightningbolltz. 12 signed up.", top=True)
                  + and_note("Pick a spec for Barrow Deeps", "Vorgash has no spec saved for Olympus XXVII. Open to choose one; nothing is signed up yet.", actions=["Choose a spec", "Out"])
                  + and_note(*NIGHT, bottom=True, attach=asset("inv_misc_cape_16")))
    return [a, b]


def moved_boards(plat: str) -> list[str]:
    """raid.roster_moved: time sensitive on iOS, the high Raid starting channel on Android; it breaks quiet hours."""
    if plat == "ios":
        a = ios_lock(ios_note(*MOVED_IN, sens=True))
        b = ios_lock(ios_note(*MOVED_BENCH, sens=True))
        c = ios_lock(ios_note("You are off the roster for Onyxia's Lair", "Kraggor removed Lightningbolltz from tonight's raid. Ask an officer if that is a mistake.", sens=True)
                     + ios_note("You are In for Onyxia's Lair", "Kraggor added you to the roster for Thursday 8 pm, as DPS on Lightningbolltz.", when="2d", sens=True))
        return [a, b, c]
    a = and_shade(and_note(*MOVED_IN, top=True, bottom=True))
    b = and_shade(and_note(*MOVED_BENCH, top=True) + and_note("You are off the roster for Onyxia's Lair", "Kraggor removed Lightningbolltz from tonight's raid. Ask an officer if that is a mistake.", bottom=True))
    return [a, b]


CHANNEL_ROWS = [
    ("Raids", [("Sign-ups open", "Push", ""), ("Not signed up, a day before", "Push", ""), ("Time moved or cancelled", "Push and Discord", ""),
               ("Roster set", "Push", ""), ("Moved on or off the roster", "Push", "Comes through quiet hours"), ("Raid starting", "Push", "1 hour before invites, set by your guild")]),
    ("After the raid", [("Your raid night", "Push", "Report, loot, bests and coach notes, in one"), ("Loot, as it is awarded", "Off", "")]),
    ("Your characters", [("First sync", "Push", ""), ("Approved to a guild roster", "Push", ""), ("New level range", "Push", ""), ("Best in slot changes", "Off", "One digest a week")]),
    ("Sims", [("Server run finished", "Push", "")]),
    ("Officers", [("Characters waiting for approval", "Push", ""), ("A raid is short a day before", "Push", ""), ("Someone drops to Out inside a day", "Push", "")]),
    ("Account and game", [("Premium seat from your guild", "Push", ""), ("Payment failed", "Push", ""), ("Phase and raid dates", "Push", "")]),
]


def notif_settings(plat: str, denied: bool = False, picker: bool = False, xl: bool = False, sections: int = 6) -> str:
    banner = ""
    if denied:
        banner = (f"<div class=\"card quiet\" style=\"margin:10px 16px 0;padding:14px;display:flex;gap:12px\">{tile('bell')}<div style=\"flex:1\"><div class=\"t-body strong\" style=\"font-weight:600\">Notifications are off for Forever Sixty</div>"
                  f"<div class=\"t-meta muted\">Turn them on in {'Settings' if plat == 'ios' else 'Android settings'} to get raid reminders here.</div><span class=\"btn btn-secondary\" style=\"margin-top:10px\">Open settings</span></div></div>")
    out = banner
    for title, rows in CHANNEL_ROWS[:sections]:
        body = ""
        for i, (lab, val, sub) in enumerate(rows):
            sub_html = f"<div class=\"t-meta muted\">{sub}</div>" if sub else ""
            value = f"<span class=\"val t-meta\">{val}{glyph('chev-down' if plat == 'and' else 'swap', 14)}</span>"
            body += f"<div class=\"lrow\" style=\"{'background:var(--row)' if picker and title == 'Raids' and i == 2 else ''}\"><div class=\"t\"><div class=\"t-body strong\">{lab}</div>{sub_html}</div>{value}</div>"
        out += f"<div class=\"t-micro micro pad\" style=\"margin-top:18px\">{title}</div><div class=\"card quiet\" style=\"margin:6px 16px 0\">{body}</div>"
    out += ("<div class=\"t-micro micro pad\" style=\"margin-top:18px\">Quiet hours</div><div class=\"card quiet\" style=\"margin:6px 16px 0\">"
            "<div class=\"lrow\"><div class=\"t t-body strong\">Quiet hours</div><span class=\"sw on\"></span></div>"
            "<div class=\"lrow\"><div class=\"t t-body strong\">From</div><span class=\"val t-body mono\">11:00 pm</span></div><div class=\"lrow\"><div class=\"t t-body strong\">To</div><span class=\"val t-body mono\">8:00 am</span></div></div>"
            "<div class=\"t-meta muted pad\" style=\"margin-top:8px\">In your phone's time zone, Central (Chicago). Raid starting reminders still come through.</div>"
            "<div class=\"t-meta muted pad\" style=\"margin-top:12px;padding-bottom:140px\">Discord messages come as a DM from the Forever Sixty bot to hunthrop.</div>")
    overlay = ""
    if picker:
        opts = ["Push", "Discord", "Push and Discord", "Off"]
        if plat == "ios":
            rows_ = "".join(f"<div style=\"display:flex;gap:10px;align-items:center;min-height:46px;padding:0 16px;border-top:{'0' if i == 0 else '1px solid rgba(255,255,255,.08)'};font:400 17px -apple-system;color:#fff\"><span style=\"width:18px\">{glyph('check', 16) if o == 'Push and Discord' else ''}</span>{o}</div>" for i, o in enumerate(opts))
            overlay = f"<div class=\"glass\" style=\"position:absolute;right:16px;top:420px;width:230px;border-radius:20px;overflow:hidden;z-index:30\">{rows_}</div>"
        else:
            rows_ = "".join(f"<div style=\"display:flex;gap:12px;align-items:center;min-height:48px;padding:0 16px;font:400 16px Roboto;color:#e9e4d8;{'background:rgba(229,185,85,.12)' if o == 'Push and Discord' else ''}\">{o}</div>" for o in opts)
            overlay = f"<div style=\"position:absolute;right:12px;top:392px;width:220px;border-radius:12px;background:#1a1f2b;box-shadow:0 8px 24px rgba(0,0,0,.6);padding:8px 0;z-index:30\">{rows_}</div>"
    return pushed(plat, "You", "Notifications", out, "You", overlay=overlay, xl=xl)


# ---------- F8 You and Premium ----------

def you(plat: str, plan: str = "free-link") -> str:
    tag = "<span class=\"tag tag-gold\">Premium</span>" if plan == "premium" else ""
    plan_val = {"free-link": "<span class=\"val t-meta\">Free</span><span class=\"t-meta gold\" style=\"font-weight:600\">See plans</span><span class=\"chev gold\">" + glyph('external', 14) + "</span>",
                "free-other": "<span class=\"val t-meta\">Free</span>",
                "premium": "<span class=\"tag tag-gold\">Founder</span><span class=\"chev\">" + glyph('chev-right', 16) + "</span>"}[plan]
    plan_sub = "<div class=\"t-meta muted\">Premium · renews 27 Oct 2027</div>" if plan == "premium" else ""
    rows = [("crown", "Plan", plan_sub, plan_val),
            ("bell", "Notifications", "<div class=\"t-meta muted\">Raids, logs, sims and dates</div>", f"<span class=\"chev\">{glyph('chev-right', 16)}</span>"),
            ("user", "Characters", "<div class=\"t-meta muted\">3 on this account</div>", f"<span class=\"chev\">{glyph('chev-right', 16)}</span>"),
            ("devices", "Devices", "<div class=\"t-meta muted\">This " + ("iPhone" if plat == "ios" else "phone") + " and 2 more</div>", f"<span class=\"chev\">{glyph('chev-right', 16)}</span>"),
            ("hash", "Discord", "<div class=\"t-meta muted\">Linked as hunthrop</div>", f"<span class=\"chev\">{glyph('chev-right', 16)}</span>")]
    body = "".join(f"<div class=\"lrow\">{tile(g, 32)}<div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:600\">{lab}</div>{sub}</div>{val}</div>" for g, lab, sub, val in rows)
    more = "".join(f"<div class=\"lrow\"><div class=\"t t-body\" style=\"color:var(--strong)\">{t}</div><span class=\"chev\">{glyph('external' if ext else 'chev-right', 14)}</span></div>"
                   for t, ext in [("Help on the Discord server", True), ("Privacy", True), ("Terms", True), ("About Forever Sixty", False)])
    inner = (f"<div class=\"pad\" style=\"display:flex;gap:12px;align-items:center;margin-top:8px\">{tile('user', 44)}<div style=\"flex:1\"><div class=\"t-name strong\" style=\"font-weight:700\">{ME['battletag']}</div><div class=\"t-meta muted\">Signed in with Battle.net</div></div>{tag}</div>"
             f"<div class=\"card quiet\" style=\"margin:16px 16px 0\">{body}</div><div class=\"card quiet\" style=\"margin:16px 16px 0\">{more}</div>"
             f"<div class=\"pad\" style=\"margin-top:16px\"><span class=\"btn btn-secondary full\" style=\"color:var(--nav)\">Sign out</span></div><div class=\"t-micro muted pad\" style=\"margin-top:10px\">Version 1.0 (1) · data 1.60.1.70532</div>")
    top = content_top(plat)
    title = f"<div class=\"cond t-title pad\" style=\"margin-top:4px\">You</div>" if plat == "ios" else ""
    bar = top_bar(plat, "<span class=\"cond\" style=\"font-size:22px;padding-left:4px\">You</span>" if plat == "and" else "<span></span>")
    return device(plat, bar + f"<div style=\"padding-top:{top}px\">{title}{inner}</div>", "You")


def plan_detail(plat: str, link_region: bool = True) -> str:
    rows = "".join(f"<div style=\"display:flex;gap:10px;padding:6px 0\" class=\"t-body\"><span class=\"kill\" style=\"margin-top:3px\">{glyph('check', 16)}</span>{f}</div>" for f in FEATS)
    manage = ("<span class=\"btn btn-secondary full\" style=\"margin-top:16px;min-height:44px;height:44px;text-transform:none;letter-spacing:0;font:600 14px Barlow,sans-serif;color:var(--nav)\">Manage billing on foreversixty.gg" + glyph('external', 14) + "</span>" if link_region
              else "<div class=\"t-meta muted\" style=\"margin-top:16px\">Billing is managed on foreversixty.gg.</div>")
    inner = (f"<div class=\"card\" style=\"margin:10px 16px 0;padding:18px\"><div style=\"display:flex;justify-content:space-between;align-items:center\"><span class=\"cond t-h2\">Premium</span><span class=\"tag tag-gold\">Founder</span></div>"
             f"<div class=\"t-body\" style=\"margin-top:6px\">Premium · renews 27 Oct 2027 · $40 a year</div><div class=\"t-meta muted\">Founder price, kept while your plan stays active.</div>"
             f"<div style=\"height:1px;background:var(--line-soft);margin:14px 0\"></div><div class=\"t-micro micro\">What your plan includes</div><div style=\"margin-top:6px\">{rows}</div>{manage}</div>"
             f"<div class=\"t-meta muted pad\" style=\"margin-top:12px\">Raid sign-ups, best in slot, talents and your character are free for everyone, in the app and on the web.</div>")
    return pushed(plat, "You", "Plan", inner, "You")


def devices(plat: str, confirm: bool = False) -> str:
    me = "This iPhone" if plat == "ios" else "This phone"
    rows = [(me, "iPhone 17 · signed in 30 Oct" if plat == "ios" else "Pixel 9 · signed in 30 Oct", ""),
            ("Pixel 9" if plat == "ios" else "iPhone 17", "Last used 2 h ago", "Revoke"), ("Companion on Windows", "Last synced today, 6:12 pm", "Revoke")]
    body = "".join(f"<div class=\"lrow\">{tile('devices', 32)}<div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:600\">{n}</div><div class=\"t-meta muted\">{s}</div></div>"
                   + (f"<span class=\"btn btn-secondary\" style=\"min-height:40px;padding:0 12px;color:var(--wipe);border-color:rgba(240,115,107,.32)\">{a}</span>" if a else "") + "</div>" for n, s, a in rows)
    inner = f"<div class=\"card quiet\" style=\"margin:10px 16px 0\">{body}</div><div class=\"t-meta muted pad\" style=\"margin-top:10px\">Revoke signs a device out at once. To sign this one out, use Sign out on You.</div>"
    overlay = ""
    if confirm:
        other = "Pixel 9" if plat == "ios" else "iPhone 17"
        if plat == "ios":
            overlay = ("<div class=\"scrim\" style=\"z-index:30\"></div><div class=\"glass\" style=\"position:absolute;left:50%;top:330px;transform:translateX(-50%);width:300px;border-radius:26px;z-index:31;overflow:hidden;font-family:-apple-system;color:#fff;text-align:center\">"
                       f"<div style=\"padding:20px 18px 14px\"><div style=\"font-weight:600;font-size:17px\">Revoke {other}?</div><div style=\"font-size:13px;margin-top:6px;color:#ddd\">It is signed out of Forever Sixty at once.</div></div>"
                       "<div style=\"display:grid;grid-template-columns:1fr 1fr;gap:8px;padding:0 14px 14px\"><span style=\"height:44px;border-radius:22px;background:rgba(255,255,255,.12);display:flex;align-items:center;justify-content:center;font:600 17px -apple-system\">Cancel</span>"
                       "<span style=\"height:44px;border-radius:22px;background:rgba(255,69,58,.9);display:flex;align-items:center;justify-content:center;font:600 17px -apple-system\">Revoke</span></div></div>")
        else:
            overlay = ("<div class=\"scrim\" style=\"z-index:30\"></div><div style=\"position:absolute;left:24px;right:24px;top:330px;border-radius:28px;background:#1d2230;padding:24px;z-index:31;font-family:Roboto;color:#e9e4d8\">"
                       f"<div style=\"font-size:24px\">Revoke {other}?</div><div style=\"font-size:14px;margin-top:14px;color:#c9c4b8\">It is signed out of Forever Sixty at once.</div>"
                       "<div style=\"display:flex;justify-content:flex-end;gap:8px;margin-top:22px;font:500 14px Roboto\"><span style=\"padding:10px 12px;color:#f5d27a\">Cancel</span><span style=\"padding:10px 12px;color:#f0736b\">Revoke</span></div></div>")
    return pushed(plat, "You", "Devices", inner, "You", overlay=overlay)



# ---------- F8 build boards (spec 5.8.1) ----------

FEATS = ["Coach: after every logged raid, what to fix first. It names the casts you missed, the buffs you let drop and the consumables you skipped.",
         "Server sims keep running, and the site tells you when they finish.",
         "See if you are improving: your Forever Score and every boss parse on a chart over time, with the fights behind each point.",
         "A message in your own Discord channel when a server sim finishes or a log is parsed.",
         "A Supporter mark beside your name on your characters and in your guild's roster."]


def _chev(plat: str = "ios") -> str:
    return f"<span class=\"chev\">{glyph('chev-right', 16)}</span>"


def _plan_row(state: str, xl: bool) -> str:
    """The Plan row. state: free-link, free-other, own."""
    sub = {"free-link": "Free", "free-other": "Free", "own": "Premium · renews 27 Oct 2027"}[state]
    founder = "<span class=\"tag tag-gold\">Founder</span>" if state == "own" else ""
    see = f"<span class=\"t-meta gold\" style=\"font-weight:600;display:inline-flex;gap:4px;align-items:center\">See plans{glyph('external', 14)}</span>"
    if xl:
        lines = f"<div class=\"t-body strong\" style=\"font-weight:600\">Plan</div><div class=\"t-meta muted\">{sub}</div>"
        lines += (f"<div style=\"margin-top:6px\">{founder}</div>" if founder else "")
        lines += (f"<div style=\"margin-top:6px\">{see}</div>" if state == "free-link" else "")
        trail = _chev() if state == "own" else (f"<span class=\"chev gold\">{glyph('external', 14)}</span>" if False else "")
        return f"<div class=\"lrow\" style=\"align-items:flex-start\">{tile('crown', 32)}<div class=\"t\">{lines}</div>{trail}</div>"
    if state == "own":
        return (f"<div class=\"lrow\">{tile('crown', 32)}<div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:600\">Plan</div><div class=\"t-meta muted\">{sub}</div></div>{founder}{_chev()}</div>")
    trail = f"{see}" if state == "free-link" else ""
    return (f"<div class=\"lrow\">{tile('crown', 32)}<div class=\"t t-body strong\" style=\"font-weight:600\">Plan</div><span class=\"val t-meta\">Free</span>{trail}</div>")


def you_build(plat: str, state: str, xl: bool = False, premium_tag: bool = False) -> str:
    rows = [("bell", "Notifications", "Raids, logs, sims and dates"), ("user", "Characters", "3 on this account"),
            ("devices", "Devices", "This " + ("iPhone" if plat == "ios" else "phone") + " and 2 more"), ("hash", "Discord", "Linked as hunthrop")]
    body = _plan_row(state, xl) + "".join(
        f"<div class=\"lrow\" style=\"{'align-items:flex-start' if xl else ''}\">{tile(g, 32)}<div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:600\">{lab}</div><div class=\"t-meta muted\">{sub}</div></div>{_chev()}</div>"
        for g, lab, sub in rows)
    more = "".join(f"<div class=\"lrow\"><div class=\"t t-body\" style=\"color:var(--strong)\">{t}</div><span class=\"chev\">{glyph('external' if ext else 'chev-right', 14)}</span></div>"
                   for t, ext in [("Help on the Discord server", True), ("Privacy", True), ("Terms", True), ("About Forever Sixty", False)])
    tag = "<span class=\"tag tag-gold\">Premium</span>" if premium_tag else ""
    stack = "flex-direction:column;align-items:flex-start;" if xl else "align-items:center;"
    head_ = (f"<div class=\"pad\" style=\"display:flex;gap:12px;{stack}margin-top:8px\">{tile('user', 44)}<div style=\"flex:1;min-width:0;max-width:100%\">"
             f"<div class=\"t-name strong\" style=\"font-weight:700;overflow-wrap:anywhere\">{ME['battletag']}</div><div class=\"t-meta muted\">Signed in with Battle.net</div></div>{tag}</div>")
    inner = (head_ + f"<div class=\"card quiet\" style=\"margin:16px 16px 0\">{body}</div><div class=\"card quiet\" style=\"margin:16px 16px 0\">{more}</div>"
             "<div class=\"pad\" style=\"margin-top:16px\"><span class=\"btn btn-secondary full\" style=\"color:var(--nav)\">Sign out</span></div>")
    top = content_top(plat)
    title = "<div class=\"cond t-title pad\" style=\"margin-top:4px\">You</div>" if plat == "ios" else ""
    bar = top_bar(plat, "<span class=\"cond\" style=\"font-size:22px;padding-left:4px\">You</span>" if plat == "and" else "<span></span>")
    return device(plat, bar + f"<div style=\"padding-top:{top}px\">{title}{inner}</div>", "You", xl=xl)


def plan_build(plat: str, source: str, link: bool) -> str:
    """F8.4. source: own or seat."""
    own = source == "own"
    tag = "<span class=\"tag tag-gold\">Founder</span>" if own else ""
    line = "Premium · renews 27 Oct 2027 · $40 a year" if own else "Premium · a seat from OLYMPUS XXVII"
    second = ("Founder price, kept while your plan stays active." if own
              else "Given by Hunthrop#1894. It lasts while the guild's plan does; nothing is billed to you.")
    rows = "".join(f"<div style=\"display:flex;gap:10px;padding:6px 0\" class=\"t-body\"><span class=\"kill\" style=\"margin-top:3px\">{glyph('check', 16)}</span>{f}</div>" for f in FEATS)
    if not own:
        control = ""
    elif link:
        control = "<span class=\"btn btn-secondary full\" style=\"margin-top:16px;min-height:44px;height:44px;text-transform:none;letter-spacing:0;font:600 14px Barlow,sans-serif;color:var(--nav)\">Manage billing on foreversixty.gg" + glyph('external', 14) + "</span>"
    else:
        control = "<div class=\"t-meta muted\" style=\"margin-top:16px\">Billing is managed on foreversixty.gg.</div>"
    inner = (f"<div class=\"card\" style=\"margin:10px 16px 0;padding:18px\"><div style=\"display:flex;justify-content:space-between;align-items:center\"><span class=\"cond t-h2\">Premium</span>{tag}</div>"
             f"<div class=\"t-body\" style=\"margin-top:6px\">{line}</div><div class=\"t-meta muted\">{second}</div>"
             f"<div style=\"height:1px;background:var(--line-soft);margin:14px 0\"></div><div class=\"t-micro micro\">What your plan includes</div><div style=\"margin-top:6px\">{rows}</div>{control}</div>"
             "<div class=\"t-meta muted pad\" style=\"margin-top:12px\">Raid sign-ups, best in slot, talents and your character are free for everyone, in the app and on the web.</div>")
    return pushed(plat, "You", "Plan", inner, "You")
