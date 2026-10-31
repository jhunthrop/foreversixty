"""F5 sims, F6 talents, F7 notifications (lock screens, settings, the in-context ask), F8 You and Premium."""
from kit import (DATA, HERE, ME, ROLES, back_btn, content_top, device, glyph, item_icon, offline_strip, sheet, spec_icon, top_bar, uri)
from model import SLOT_ORDER, slots

APP_ICON = uri(HERE.parents[1] / "logo/export/foreversixty-icon-64.png")


def tile(g: str, size: int = 36) -> str:
    return f"<span style=\"flex:none;width:{size}px;height:{size}px;border-radius:8px;display:grid;place-items:center;background:var(--tile);color:var(--tile-ink)\">{glyph(g, size // 2)}</span>"


def pushed(plat: str, back: str, title: str, inner: str, tab: str, overlay: str = "", xl: bool = False, right: str = "") -> str:
    top = content_top(plat)
    left = back_btn(plat, back) + (f"<span class=\"cond\" style=\"font-size:22px\">{title}</span>" if plat == "and" else "")
    head_ = "" if plat == "and" else f"<div class=\"cond t-title pad\" style=\"margin-top:4px\">{title}</div>"
    return device(plat, top_bar(plat, f"<span style=\"display:flex;align-items:center;gap:4px\">{left}</span>", right) + f"<div style=\"padding-top:{top}px\">{head_}{inner}</div>", tab, overlay=overlay, xl=xl)


# ---------- F5 sims (revision for today's API, 2026-10-31) ----------
# Rows follow the web's ServerRunRow: title, meta "{Spec} · {Kind}" (kind alone when untitled), a state line
# with the age at right (dropping under the state when they do not fit), an 8 pt bar only for a bulk run
# (combos_total > 0), `Run again` only on a failed row.

ELE = "spell_nature_lightning"
WAR = "ability_warrior_innerrage"
HEADLINE = "+2 DPS from Spritecaster Cape"
BASE_BACK = 19121     # Deep Woodlands Cloak: the cloak the run wore; the result swaps in Spritecaster Cape
CHANGED_SLOT = "back"
LEGEND = "Gold corner: the slot this result changes"
FOOT_PREMIUM = "Start a sim on foreversixty.gg/sim. Server runs keep going after you close the app."
FOOT_FREE = "Start a sim on foreversixty.gg/sim and it shows here."
OVERFLOW = "<div class=\"t-meta muted\" style=\"padding:8px 16px\">Older finished runs are in Your sims below.</div>"
TAB_CLEAR = "<div style=\"height:120px\"></div>"


def hit(plat: str) -> int:
    return 44 if plat == "ios" else 48


def _dot() -> str:
    return "<span style=\"display:inline-block;width:8px;height:8px;border-radius:4px;background:var(--gold)\"></span>"


def _state_line(left: str, right: str, xl: bool) -> str:
    # State never wraps first: the age drops under it when both do not fit (flex-wrap), always under at XL.
    left_style = "" if xl else "white-space:nowrap"
    wrap = "flex-direction:column;gap:2px" if xl else "flex-wrap:wrap;justify-content:space-between;column-gap:12px"
    return (f"<div class=\"t-meta\" style=\"display:flex;align-items:baseline;{wrap};margin-top:4px\"><span style=\"{left_style}\">{left}</span>"
            f"<span class=\"muted\" style=\"{'' if xl else 'white-space:nowrap'}\">{right}</span></div>")


def _btn(label: str, plat: str, *, disabled: bool = False, xl: bool = False) -> str:
    fs = f"font-size:{ROLES[f'{plat}-xl']['meta']}px;white-space:normal;padding:10px 18px;" if xl else ""
    return f"<span class=\"btn btn-secondary full\" style=\"margin-top:8px;min-height:{hit(plat)}px;{fs}{'opacity:.45' if disabled else ''}\">{label}</span>"


def _line(text: str, tone: str = "wipe", extra: str = "") -> str:
    return f"<div class=\"t-meta {tone}\" role=\"alert\" style=\"margin-top:6px\">{text}{extra}</div>"


def _plans(plat: str) -> str:
    return f"<span class=\"gold\" style=\"text-decoration:underline;display:inline-flex;min-height:{hit(plat)}px;align-items:center;margin-left:4px\">See plans</span>"


def _srow(title: str, meta: str, state: str, age: str, *, unseen: bool = False, bar: tuple[int, int] | None = None, action: str = "", link: bool = False, xl: bool = False) -> str:
    dot = f"<span aria-label=\"Not opened yet\" style=\"margin-right:8px;display:inline-block;vertical-align:middle\">{_dot()}</span>" if unseen else ""
    barhtml = ""
    if bar:
        done, total = bar
        barhtml = (f"<div class=\"bar8\" style=\"margin-top:8px\"><i style=\"width:{done * 100 // total}%\"></i></div>"
                   f"<div class=\"t-micro muted mono\" style=\"margin-top:4px\">{done} of {total} combinations</div>")
    chev = f"<span class=\"chev\" style=\"align-self:center\">{glyph('chev-right', 16)}</span>" if link else ""
    return (f"<div class=\"lrow\" style=\"align-items:flex-start;min-height:{88 if not link else 72}px;padding:12px 16px\">{spec_icon(ELE, 36)}"
            f"<div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:600\">{dot}{title}</div><div class=\"t-meta muted\">{meta}</div>"
            f"{_state_line(state, age, xl)}{barhtml}{action}</div>{chev}</div>")


RUN_PLAIN = lambda xl=False: _srow("Council pull", "Elemental · Sim", "<span class=\"gold\">Running</span> · takes about 4 minutes", "Started 1 min ago", xl=xl)
RUN_BULK = lambda xl=False: _srow("Elemental", "Top Gear", "<span class=\"gold\">Running</span> · takes about 4 minutes", "Started 3 min ago", bar=(40, 64), xl=xl)
DONE_UNSEEN = lambda xl=False: _srow("MC full buffs", "Elemental · Sim", "<span class=\"mono strong\">412 DPS</span>", "Finished 2 h ago", unseen=True, link=True, xl=xl)


def _failed(plat: str, xl: bool = False, action: str | None = None) -> str:
    return _srow("Elemental", "Stat weights", "<span class=\"wipe\">Failed: the engine stopped</span>", "Ended yesterday",
                 action=_btn("Run again", plat, xl=xl) if action is None else action, xl=xl)


def _srows(plat: str, xl: bool = False) -> str:
    return RUN_BULK(xl) + RUN_PLAIN(xl) + DONE_UNSEEN(xl) + _failed(plat, xl)


def _mine_row(ic: str, title: str, headline: str, meta: str, *, unseen: bool = False, mono: bool = True) -> str:
    dot = f"<span style=\"margin-right:8px;display:inline-block;vertical-align:middle\">{_dot()}</span>" if unseen else ""
    return (f"<div class=\"lrow\" style=\"align-items:flex-start;min-height:72px;padding:12px 16px\">{spec_icon(ic, 32)}<div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:600\">{dot}{title}</div>"
            f"<div class=\"t-body strong{' mono' if mono else ''}\">{headline}</div><div class=\"t-meta muted\">{meta}</div></div>"
            f"<span class=\"chev\" style=\"align-self:center\">{glyph('chev-right', 16)}</span></div>")


def _mine(count: int = 3) -> str:
    rows = [_mine_row(ELE, "MC full buffs", "412 DPS", "Ran on the server · 31 Oct", unseen=True),
            _mine_row(ELE, "Elemental", HEADLINE, "Saved from the browser · 28 Oct", mono=False),
            _mine_row(WAR, "Fury 40 to 49", "88 DPS", "Saved from the browser · 22 Oct")]
    return "".join(rows[:count])


def _section(label: str, inner: str, first: bool = False) -> str:
    return f"<div class=\"t-micro micro pad\" style=\"margin-top:{12 if first else 22}px\">{label}</div><div class=\"card quiet\" style=\"margin:8px 16px 0\">{inner}</div>"


def _skel_row(h: int, icon: int) -> str:
    return (f"<div class=\"lrow\" style=\"min-height:{h}px;padding:12px 16px;align-items:flex-start\"><span class=\"skel\" style=\"flex:none;width:{icon}px;height:{icon}px\"></span>"
            f"<div class=\"t\" style=\"display:flex;flex-direction:column;gap:8px\"><span class=\"skel\" style=\"height:15px;width:180px;max-width:100%\"></span><span class=\"skel\" style=\"height:12px;width:120px\"></span>"
            f"<span class=\"skel\" style=\"height:14px;width:140px\"></span></div></div>")


def _failed_block(msg: str) -> str:
    return (f"<div style=\"padding:16px\"><div class=\"t-body wipe\" role=\"alert\">{msg}</div>"
            f"<span class=\"btn btn-secondary\" style=\"margin-top:12px\">{glyph('refresh', 16)}Try again</span></div>")


def _spinner_row() -> str:
    return ("<div style=\"height:56px;display:grid;place-items:center\"><span style=\"width:22px;height:22px;border-radius:11px;"
            "border:3px solid var(--line);border-top-color:var(--gold)\"></span></div>")


def _msg(text: str) -> str:
    return f"<div style=\"padding:18px 16px\"><div class=\"t-body\">{text}</div></div>"


def _screen(plat: str, title: str, back: str, sections: str, *, strip: str = "", foot: str = "", xl: bool = False, height: int | None = None, tab: bool = True) -> str:
    """A pushed screen: iOS large title then the offline strip, Android title in the bar then the strip."""
    top = content_top(plat)
    head_ = "" if plat == "and" else f"<div class=\"cond t-title pad\" style=\"margin-top:4px\">{title}</div>"
    foot_html = f"<div class=\"t-meta muted pad\" style=\"margin-top:12px\">{foot}</div>" if foot else ""
    left = back_btn(plat, back) + (f"<span class=\"cond\" style=\"font-size:22px\">{title}</span>" if plat == "and" else "")
    bar = top_bar(plat, f"<span style=\"display:flex;align-items:center;gap:4px\">{left}</span>", "")
    return device(plat, bar + f"<div style=\"padding-top:{top}px\">{head_}{strip}{sections}{foot_html}{TAB_CLEAR}</div>", "Home", xl=xl, height=height)


def sims(plat: str, state: str = "premium", xl: bool = False) -> str:
    strip = offline_strip() if state == "offline" else ""
    server: str | None
    height = None
    foot = FOOT_PREMIUM
    if state == "premium":
        server, mine, height = _srows(plat, xl) + OVERFLOW, _mine(2), (None if xl else 1160)
        if xl:
            server = _failed(plat, True) + RUN_PLAIN(True)
    elif state == "free":
        server, mine, foot = None, _mine(), FOOT_FREE
    elif state == "lapsed":
        failed = _failed(plat, action=_btn("Run again", plat) + _line("Server sims are part of Premium.", "wipe", _plans(plat)))
        old = _srow("MC full buffs", "Elemental · Sim", "<span class=\"mono strong\">412 DPS</span>", "Finished 12 Oct", link=True)
        server, mine, foot = failed + old, _mine(), FOOT_FREE
    elif state == "loading":
        server, mine, foot = "".join(_skel_row(88, 36) for _ in range(3)), "".join(_skel_row(72, 32) for _ in range(3)), ""
    elif state == "empty":
        server, mine, foot = _msg("No server runs yet. Choose Run on our servers when you start a sim, and it keeps going after you close the app."), _msg("No sims yet. Start one on foreversixty.gg/sim and it shows here."), ""
    elif state == "error":
        server, mine = _failed_block("Your server runs did not load."), _mine()
    elif state == "paging":
        server, mine = None, _mine(3) + _spinner_row()
        foot = ""
    elif state == "pagefail":
        server, mine = None, _mine(3) + _failed_block("Your sims did not load.")
        foot = ""
    else:  # offline: cached rows, polling stopped, Run again disabled with its sentence
        server = (_srow("Council pull", "Elemental · Sim", "<span class=\"mono strong\">388 DPS</span>", "Finished 2 h ago", link=True)
                  + _failed(plat, action=_btn("Run again", plat, disabled=True) + _line("You are offline. Run again needs a connection.", "muted")))
        mine = _mine()
    sections = (_section("Server runs", server, first=True) if server is not None else "") + _section("Your sims", mine, first=server is None)
    return _screen(plat, "Sims", "Home", sections, strip=strip, foot=foot, xl=xl, height=height)


def rerun(plat: str, state: str) -> str:
    """One failed Server runs row in each Run again state."""
    strip = offline_strip() if state == "offline" else ""
    action = {
        "starting": _btn("Starting…", plat, disabled=True),
        "started": _btn("Started", plat, disabled=True),
        "failed": _btn("Run again", plat) + _line("Could not start it again. Try again."),
        "lapsed": _btn("Run again", plat) + _line("Server sims are part of Premium.", "wipe", _plans(plat)),
        "lapsed-plain": _btn("Run again", plat) + _line("Server sims are part of Premium."),
        "offline": _btn("Run again", plat, disabled=True) + _line("You are offline. Run again needs a connection.", "muted"),
    }[state]
    return _screen(plat, "Sims", "Home", _section("Server runs", _failed(plat, action=action), first=True), strip=strip, height=640)


def _gear_icons(kind: str, count: int | None, unknown: bool) -> str:
    keys = SLOT_ORDER if count is None else SLOT_ORDER[:count]
    out = []
    for k in keys:
        item_id = BASE_BACK if (k == CHANGED_SLOT and kind != "run") else slots()[k]["item_id"]
        icon = item_icon(item_id, 40)
        if unknown and k == "trinket2":
            icon = ("<span style=\"display:grid;place-items:center;width:40px;height:40px;border-radius:6px;border:1.5px solid var(--line);background:#11151f;"
                    "color:var(--muted)\" class=\"t-body mono\">?</span>")
        if k == CHANGED_SLOT and kind != "run":
            icon = (f"<span style=\"position:relative;display:block;width:40px;height:40px\">{icon}<span aria-label=\"Changed by this result\" style=\"position:absolute;top:0;right:0;width:0;height:0;"
                    "border-top:12px solid var(--gold);border-left:12px solid transparent;border-top-right-radius:6px\"></span></span>")
        out.append(icon)
    return "".join(out)


def sim_result(plat: str, kind: str = "run", xl: bool = False, state: str = "ok") -> str:
    fig = 64 if (xl and plat == "ios") else 60 if xl else 52
    gh_fs = ROLES[f"{plat}-xl"]["meta"] if xl else 13
    title, kind_label = ("MC full buffs", "Sim") if kind == "run" else ("Elemental", "Top Gear")
    if state == "loading":
        inner = ("<div class=\"pad\"><span class=\"skel\" style=\"display:block;height:20px;width:160px;margin-top:8px\"></span><span class=\"skel\" style=\"display:block;height:12px;width:120px;margin-top:10px\"></span>"
                 "<span class=\"skel\" style=\"display:block;height:52px;width:200px;margin-top:16px\"></span><div style=\"display:grid;grid-template-columns:repeat(6,40px);gap:8px;margin-top:40px\">"
                 + "".join("<span class=\"skel\" style=\"display:block;width:40px;height:40px\"></span>" for _ in range(12)) + "</div></div>")
        return _result_screen(plat, "Sim", inner, xl)
    if state == "failed":
        return _result_screen(plat, "Sim", "<div class=\"pad\">" + _failed_block("That sim did not load.") + "</div>", xl)
    if kind == "run":
        answer = (f"<div style=\"display:flex;align-items:baseline;gap:8px;flex-wrap:wrap;margin-top:16px\"><span class=\"cond\" style=\"font-size:{fig}px;line-height:1\">412.3</span><span class=\"t-body muted\">DPS</span></div>")
    else:
        answer = f"<div class=\"cond t-stat\" style=\"margin-top:16px;text-transform:none\">{HEADLINE}</div>"
    head_ = "" if plat == "and" else f"<div class=\"cond t-title\" style=\"margin-top:4px\">{title}</div>"
    gear = ""
    if state != "nogear":
        legend = f"<div class=\"t-meta muted\" style=\"margin-top:8px\">{LEGEND}</div>" if kind != "run" else ""
        gear = (f"<div class=\"t-micro micro\" style=\"margin-top:20px\">The gear it ran</div><div style=\"display:grid;grid-template-columns:repeat(6,40px);gap:8px;margin-top:8px\">"
                f"{_gear_icons(kind, None, state == 'unknown')}</div>{legend}")
    inner = (f"<div class=\"pad\">{head_}<div class=\"t-micro micro\" style=\"margin-top:4px\">{kind_label}</div>"
             f"<div class=\"t-meta\" style=\"margin-top:4px\"><span style=\"color:var(--c-shaman,#2f7bdb)\">Elemental Shaman</span><span class=\"muted\"> · level 60</span></div>{answer}"
             f"<div class=\"t-meta muted mono\" style=\"margin-top:6px\">± 1.8 · 10,000 iterations · engine 1f542def3</div>{gear}"
             f"<span class=\"btn btn-ghost\" style=\"margin-top:12px;padding:0;min-height:{hit(plat)}px;font-size:{gh_fs}px;white-space:normal;text-align:left;display:flex;flex-wrap:wrap;justify-content:flex-start;min-width:0;max-width:100%;overflow-wrap:anywhere\">Full breakdown on foreversixty.gg{glyph('external', 14)}</span></div>")
    return _result_screen(plat, title, inner, xl, strip=offline_strip() if state == "offline" else "", height=1500 if xl else None)


def _result_screen(plat: str, title: str, inner: str, xl: bool, strip: str = "", height: int | None = None) -> str:
    top = content_top(plat)
    left = back_btn(plat, "Sims") + (f"<span class=\"cond\" style=\"font-size:22px\">{title}</span>" if plat == "and" else "")
    bar = top_bar(plat, f"<span style=\"display:flex;align-items:center;gap:4px\">{left}</span>", "")
    return device(plat, bar + f"<div style=\"padding-top:{top}px\">{strip}{inner}{TAB_CLEAR}</div>", "Home", xl=xl, height=height)


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
    feats = ["Server sims that keep running after you close the app", "Coach notes on your raid logs", "Score and parse history charts", "Sim results and logs kept for a year"]
    rows = "".join(f"<div style=\"display:flex;gap:10px;padding:6px 0\" class=\"t-body\"><span class=\"kill\" style=\"margin-top:3px\">{glyph('check', 16)}</span>{f}</div>" for f in feats)
    manage = ("<span class=\"btn btn-secondary full\" style=\"margin-top:16px\">Manage billing on foreversixty.gg" + glyph('external', 14) + "</span>" if link_region
              else "<div class=\"t-meta muted\" style=\"margin-top:16px\">Billing is managed on the web.</div>")
    inner = (f"<div class=\"card\" style=\"margin:10px 16px 0;padding:18px\"><div style=\"display:flex;justify-content:space-between;align-items:center\"><span class=\"cond t-h2\">Premium</span><span class=\"tag tag-gold\">Founder</span></div>"
             f"<div class=\"t-body\" style=\"margin-top:6px\">Renews 27 Oct 2027 · $50.00 a year</div><div class=\"t-meta muted\">Founder price, kept while your plan stays active.</div>"
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

