"""F4 guild and raids: the Raids tab, the raid page, the sign-up, the spec ask and every raid state."""
from kit import (CLASS_COLOR, HORDE, ME, back_btn, content_top, crest, device, glyph, icon_btn, offline_strip, selector, sheet,
                 spec_icon, top_bar)
from model import CLASS_ORDER, NOT_SIGNED, ROSTER, SPEC, class_counts

STATES = ["In", "Late", "Bench", "Out"]


def seg(on: str | None = "In", disabled: bool = False, pending: str | None = None, xl: bool = False, focus: str | None = None) -> str:
    parts = ""
    for s in STATES:
        cls = ""
        if s == on:
            cls = "on-go" if s in ("In", "Late") else "on-no"
        if s == pending:
            cls = "pend"
        if s == focus:
            cls += " focus"
        parts += f"<span class=\"t-seg {cls}\" role=\"radio\" aria-checked=\"{'true' if s == on else 'false'}\">{s}</span>"
    return f"<div class=\"seg4{' dis' if disabled else ''}{' xl' if xl else ''}\" role=\"radiogroup\" aria-label=\"Lightningbolltz's sign-up\">{parts}</div>"


def who(spec: str = "shaman-elemental", name: str = ME["name"], extra: str = "") -> str:
    cls, sname, role, icon = SPEC[spec]
    return (f"<div style=\"display:flex;align-items:center;gap:10px\">{crest(cls, 36)}<div style=\"flex:1;min-width:0\"><div class=\"t-body\" style=\"font-weight:700;color:{CLASS_COLOR[cls]}\">{name}</div>"
            f"<div class=\"t-meta muted\" style=\"display:flex;align-items:center;gap:6px\">{spec_icon(icon, 18)}{sname} · {role}</div></div>{extra}</div>")


def count_block(n: int = 24, size: int | None = 40, muted: bool = False) -> str:
    bar = f"<div class=\"bar8\" style=\"width:120px\"><i style=\"width:{n / size * 100:.0f}%;{'background:var(--muted)' if muted else ''}\"></i></div>" if size else ""
    line = f"{n} of {size} signed up" if size else f"{n} signed up"
    return (f"<div style=\"display:flex;align-items:center;gap:10px;margin-top:12px\">{bar}<span class=\"mono t-body strong\">{line}</span></div>"
            f"<div class=\"t-meta muted\" style=\"margin-top:4px\">3 tanks · 6 healers · 14 DPS · 1 late</div>")


def status_line(text: str, tone: str = "kill") -> str:
    return f"<div class=\"t-meta {tone}\" style=\"min-height:20px;margin-top:8px\" role=\"status\">{text}</div>"


def next_raid_card(state: str = "In", line: str = "Saved.") -> str:
    return (f"<div class=\"card\" style=\"margin:14px 16px 0;padding:16px\"><div style=\"display:flex;gap:12px\"><div class=\"dateblk\"><span class=\"m\">THU</span><span class=\"d\">10</span><span class=\"m\">DEC</span></div>"
            f"<div style=\"flex:1;min-width:0\"><div class=\"t-micro micro gold\">Next raid</div><div class=\"cond t-sec\" style=\"font-size:22px;margin-top:2px\">Onyxia's Lair</div>"
            f"<div class=\"t-meta\">Invites <span class=\"mono\">7:45 pm</span> · Pull <span class=\"mono\">8:00 pm</span></div><div class=\"t-meta\">Invites <span class=\"mono gold\">in 2 days 2 h</span></div></div></div>"
            f"{count_block()}<div style=\"height:1px;background:var(--line-soft);margin:14px 0\"></div><div class=\"t-micro micro\" style=\"margin-bottom:8px\">Your sign-up</div>{who()}"
            f"<div style=\"margin-top:10px\">{seg(state)}</div>{status_line(line)}"
            f"<div class=\"t-body gold\" style=\"font-weight:600;display:flex;align-items:center;gap:4px;min-height:44px\">Who is coming{glyph('chev-right', 16)}</div></div>")


def planned_row(dow: str, day: str, mon: str, title: str, when: str, count: str, chip: str, bar_pct: int | None) -> str:
    bar = f"<div class=\"bar8\" style=\"width:64px\"><i style=\"width:{bar_pct}%\"></i></div>" if bar_pct is not None else ""
    return (f"<div class=\"lrow\" style=\"align-items:flex-start;padding:12px 16px\"><div class=\"dateblk\" style=\"width:48px;height:52px\"><span class=\"m\">{dow}</span><span class=\"d\" style=\"font-size:20px\">{day}</span><span class=\"m\">{mon}</span></div>"
            f"<div class=\"t\"><div style=\"display:flex;justify-content:space-between;gap:8px\"><span class=\"t-body strong\" style=\"font-weight:600\">{title}</span>{chip}</div>"
            f"<div class=\"t-meta muted\">{when}</div><div style=\"display:flex;align-items:center;gap:8px;margin-top:4px\"><span class=\"mono t-meta strong\">{count}</span>{bar}</div></div></div>")


def guild_head() -> str:
    return (f"<div class=\"pad\" style=\"display:flex;align-items:center;gap:12px;margin-top:8px\"><img src=\"{HORDE}\" width=\"40\" height=\"40\" alt=\"\" style=\"border-radius:50%\">"
            f"<div style=\"flex:1\"><div class=\"cond t-sec\" style=\"font-size:22px\">Olympus XXVII</div><div class=\"t-meta muted\">Horde · US Normal · 38 on the roster</div></div></div>")


def raids_tab(plat: str, state: str = "default") -> str:
    top = content_top(plat)
    bar = top_bar(plat, selector(plat))
    if state == "noguild":
        body = (f"<div style=\"padding:{top + 40}px 24px 0\"><span style=\"width:56px;height:56px;border-radius:12px;display:grid;place-items:center;background:var(--tile);color:var(--tile-ink)\">{glyph('calendar', 28)}</span>"
                f"<div class=\"cond t-h2\" style=\"margin-top:18px\">No guild yet</div><div class=\"t-body\" style=\"margin-top:8px;font-size:17px;line-height:25px\">Lightningbolltz is not in a guild on Forever Sixty. When your guild's officers add it, its raids show here and you sign up in one tap.</div>"
                f"<div class=\"t-meta muted\" style=\"margin-top:14px\">Your guild is on Forever Sixty once one officer claims it at foreversixty.gg/guilds.</div></div>")
        return device(plat, bar + body, "Raids")
    if state == "waiting":
        body = (f"<div style=\"padding-top:{top}px\">{guild_head()}<div class=\"card quiet\" style=\"margin:16px 16px 0;padding:16px\">{who(extra='<span class=\"tag tag-muted\">Waiting</span>')}"
                f"<div class=\"t-body\" style=\"margin-top:12px\">You are waiting for approval. Planned raids show once an officer approves you.</div></div></div>")
        return device(plat, bar + body, "Raids")
    if state == "error":
        body = (f"<div style=\"padding-top:{top}px\">{guild_head()}<div class=\"card quiet\" style=\"margin:16px 16px 0;padding:16px\"><div class=\"t-body wipe\" role=\"alert\">Planned raids did not load.</div>"
                f"<span class=\"btn btn-secondary\" style=\"margin-top:12px\">{glyph('refresh', 16)}Try again</span></div></div>")
        return device(plat, bar + body, "Raids")
    if state == "empty":
        body = (f"<div style=\"padding-top:{top}px\">{guild_head()}<div class=\"card quiet\" style=\"margin:16px 16px 0;padding:18px 16px\"><div class=\"t-body strong\" style=\"font-size:17px\">No raid planned yet.</div>"
                f"<div class=\"t-meta muted\" style=\"margin-top:6px\">When an officer plans one, it shows here and we tell you.</div></div></div>")
        return device(plat, bar + body, "Raids")
    strip = offline_strip() if state == "offline" else ""
    card = next_raid_card() if state != "offline" else next_raid_offline()
    rows = (planned_row("SAT", "12", "DEC", "Barrow Deeps", "Saturday 7:00 pm · in 4 days", "11 signed up", "<span class=\"tag tag-gold\">Not signed up</span>", None)
            + planned_row("THU", "17", "DEC", "Onyxia's Lair", "Thu 17 Dec, 8:00 pm", "8/40", "<span class=\"tag tag-gold\">Not signed up</span>", 20))
    body = (f"<div style=\"padding-top:{top}px\">{strip}{guild_head()}{card}<div class=\"t-micro micro pad\" style=\"margin-top:22px\">Also planned <span class=\"mono\">2</span></div>"
            f"<div class=\"card quiet\" style=\"margin:8px 16px 0\">{rows}</div></div>")
    return device(plat, bar + body, "Raids")


def next_raid_offline() -> str:
    return (next_raid_card(line="")
            .replace("class=\"seg4\"", "class=\"seg4 dis\"")
            .replace("<div class=\"t-meta kill\" style=\"min-height:20px;margin-top:8px\" role=\"status\"></div>",
                     "<div class=\"t-meta muted\" style=\"min-height:20px;margin-top:8px\">You are offline. Sign-ups need a connection.</div>"))


def class_strip() -> str:
    counts = class_counts()
    cells = "".join(f"<div style=\"display:flex;flex-direction:column;align-items:center;gap:4px;{'opacity:.4' if counts[c] == 0 else ''}\">{crest(c, 30)}<span class=\"mono t-meta strong\">{counts[c]}</span></div>" for c in CLASS_ORDER)
    return f"<div class=\"t-micro micro pad\" style=\"margin-top:22px\">Classes coming</div><div style=\"display:grid;grid-template-columns:repeat(9,1fr);padding:10px 12px 0\">{cells}</div>"


def state_button(word: str, open_: bool = False) -> str:
    ring = "border-color:var(--gold);box-shadow:0 0 26px rgba(229,185,85,.32);" if open_ else ""
    return (f"<span class=\"t-meta\" style=\"margin-left:auto;flex:none;display:inline-flex;align-items:center;gap:4px;min-height:44px;padding:0 10px 0 12px;border-radius:8px;"
            f"border:1px solid var(--line);background:rgba(255,255,255,.03);color:var(--strong);font-weight:600;{ring}\" aria-label=\"Place for this raider\">{word}{glyph('chev-down', 14)}</span>")


def roster_row(name: str, spec: str, note: str, own: bool = False, control: str = "") -> str:
    cls, sname, _, icon = SPEC[spec]
    note_html = f"<div class=\"t-meta\" style=\"color:var(--text)\">{note}</div>" if note else ""
    return (f"<div style=\"display:flex;gap:10px;align-items:flex-start;padding:8px 16px;min-height:48px;{'background:var(--row);' if own else ''}\">{spec_icon(icon, 28)}"
            f"<div style=\"min-width:0\"><div class=\"t-body\" style=\"font-weight:600;color:{CLASS_COLOR[cls]};line-height:18px\">{name}</div><div class=\"t-meta muted\" style=\"line-height:16px\">{sname}</div>{note_html}</div>{control}</div>")


GROUP_WORD = {"Tanks": "In", "Healers": "In", "DPS": "In", "Late": "Late", "Bench": "Bench", "Out": "Out"}


def roster(officer: bool = False, open_name: str = "") -> str:
    out = "<div class=\"pad\" style=\"display:flex;align-items:baseline;gap:10px;margin-top:26px\"><span class=\"cond t-h2\">Sign-ups</span><span class=\"mono muted t-meta\">24/40</span></div>"
    for group, rows in ROSTER.items():
        n = len(rows)
        out += f"<div class=\"t-micro micro pad\" style=\"margin-top:16px\">{group} <span class=\"mono\">{n}</span></div><div class=\"card quiet\" style=\"margin:6px 16px 0;padding:4px 0\">"
        out += "".join(roster_row(nm, sp, note, nm == ME["name"], state_button(GROUP_WORD[group], nm == open_name) if officer else "") for nm, sp, note in rows) + "</div>"
    names = ", ".join(f"<span style=\"color:{CLASS_COLOR[c]}\">{n}</span>" for n, c in NOT_SIGNED)
    out += f"<div class=\"t-micro micro pad\" style=\"margin-top:16px\">Not signed up <span class=\"mono\">{len(NOT_SIGNED)}</span></div><div class=\"t-meta pad\" style=\"margin-top:6px\">{names}</div>"
    return out


def raid_head(plat: str, chip: str = "<span class=\"tag tag-gold\">Planned</span>", countdown: str = "in 2 days 2 h", muted: bool = False) -> str:
    title = "" if plat == "and" else "<div class=\"cond t-title\" style=\"margin-top:6px\">Onyxia's Lair</div>"
    cd_cls = "muted" if muted else "gold"
    return (f"<div class=\"pad\"><div style=\"display:flex;align-items:center;gap:10px\"><span class=\"t-micro micro gold\">Raid</span>{chip}<span class=\"t-meta muted\">Olympus XXVII</span></div>{title}"
            f"<div class=\"t-body\" style=\"margin-top:6px\">Thursday 10 December</div>"
            f"<div class=\"t-body\"><span class=\"muted\">Invites</span> <span class=\"mono\">7:45 pm</span> · <span class=\"muted\">Pull</span> <span class=\"mono\">8:00 pm</span> · <span class=\"muted\">Ends</span> <span class=\"mono\">11:00 pm</span></div>"
            f"<div class=\"t-body\">Invites <span class=\"mono {cd_cls}\">{countdown}</span></div>"
            f"{count_block(muted=muted)}</div>")


def signup_card(state: str | None = "In", line: str = "Saved.", tone: str = "kill", disabled: bool = False, pending: str | None = None, extra: str = "", xl: bool = False, note: bool = True, second: str = "") -> str:
    change = "<span class=\"t-meta gold\" style=\"font-weight:600;min-height:44px;display:flex;align-items:center\">Change spec</span>"
    note_f = ("<div style=\"margin-top:10px;display:flex;align-items:center;min-height:48px;padding:0 14px;border-radius:8px;border:1px solid var(--card-b);background:rgba(0,0,0,.35)\" class=\"t-body muted\">Note for the officers (optional)</div>" if note else "")
    return (f"<div class=\"card\" style=\"margin:16px 16px 0;padding:16px\"><div class=\"cond t-sec\" style=\"margin-bottom:10px\">Your sign-up</div>{who(extra='' if xl else change)}"
            f"<div style=\"margin-top:12px\">{seg(state, disabled, pending, xl)}</div>{status_line(line, tone)}{note_f}{extra}{second}</div>")


def raid_bar(plat: str) -> str:
    right = icon_btn(plat, "cal-plus", "Add to calendar") + icon_btn(plat, "more" if plat == "ios" else "more-v", "More")
    left = back_btn(plat, "Raids") + ("<span class=\"cond\" style=\"font-size:22px\">Onyxia's Lair</span>" if plat == "and" else "")
    return top_bar(plat, f"<span style=\"display:flex;align-items:center;gap:4px\">{left}</span>", right)


def raid_detail(plat: str, state: str = "default", xl: bool = False, overlay: str = "", height: int | None = None, full: bool = False,
                scroll: int = 0, officer_roster: bool = False) -> str:
    top = content_top(plat)
    head_ = raid_head(plat)
    card = signup_card(xl=xl)
    strip = ""
    if state == "locked":
        head_ = raid_head(plat, "<span class=\"tag tag-muted\">Locked</span>")
        card = signup_card(line="Sign-ups are locked. An officer can still change yours.", tone="muted", disabled=True)
    elif state == "cancelled":
        head_ = raid_head(plat, "<span class=\"tag tag-wipe\">Cancelled</span>", muted=True)
        card = f"<div class=\"card quiet\" style=\"margin:16px 16px 0;padding:16px\"><div class=\"t-body strong\" style=\"font-size:17px\">This raid was cancelled.</div><div class=\"t-meta muted\" style=\"margin-top:6px\">You were In as DPS. Nothing else to do.</div></div>"
    elif state == "unsigned":
        card = signup_card(None, line="")
    elif state == "spec":
        card = signup_card(None, line="", pending="In")
    elif state == "ask":
        ask = ("<div class=\"card quiet\" style=\"margin-top:12px;padding:14px;display:flex;gap:12px;align-items:flex-start;background:#10141d\">"
               f"<span style=\"flex:none;width:36px;height:36px;border-radius:8px;display:grid;place-items:center;background:var(--tile);color:var(--tile-ink)\">{glyph('bell', 18)}</span>"
               "<div style=\"flex:1\"><div class=\"t-body strong\" style=\"font-weight:600\">Get a reminder before raids start?</div><div class=\"t-meta muted\" style=\"margin-top:2px\">One an hour before the start, and one when sign-ups open. Change them any time in You, Notifications.</div>"
               "<div style=\"display:flex;gap:8px;margin-top:10px\"><span class=\"btn btn-secondary\" style=\"border-color:var(--gold);color:var(--gold)\">Turn on reminders</span><span class=\"btn btn-ghost\" style=\"color:var(--nav)\">Not now</span></div></div></div>")
        card = signup_card(extra=ask, note=False)
    elif state == "offline":
        strip = offline_strip()
        card = signup_card(line="You are offline. Sign-ups need a connection.", tone="muted", disabled=True)
    elif state == "failed":
        card = signup_card("In", line="That did not save. Try again.", tone="wipe")
    elif state == "moved":
        vorgash = (f"<div style=\"height:1px;background:var(--line-soft);margin:14px 0\"></div>{who('warrior-protection', 'Vorgash', change_spec())}"
                   f"<div style=\"margin-top:12px\">{seg('Bench')}</div>{status_line('Vorgash moved to Bench.', 'strong')}")
        card = signup_card("In", line="Saved.", note=False, second=vorgash)
    elif state in ("officer", "officer-locked"):
        locked = state == "officer-locked"
        if locked:
            head_ = raid_head(plat, "<span class=\"tag tag-muted\">Locked</span>")
        head_ += officer_row(locked)
        card = signup_card(line="Sign-ups are locked. Members cannot change theirs; officers can." if locked else "Saved.", tone="muted" if locked else "kill", note=False)
    body = f"<div style=\"padding-top:{top}px\"><div style=\"transform:translateY(-{scroll}px)\">{strip}{head_}{card}"
    if officer_roster:
        body += class_strip() + roster(officer=True)
    if full:
        body += class_strip() + ("<div class=\"cond t-h2 pad\" style=\"margin-top:26px\">Notes</div><div class=\"t-body pad\" style=\"margin-top:6px\">Bring <b>Fire Resistance</b> gear for Deep Breath in phase 2 and a stack of <b>Greater Fire Protection Potions</b>. "
                                  "Raid Discord: <span class=\"gold\" style=\"text-decoration:underline\">#raids</span>.</div>") + roster() + "<div style=\"height:130px\"></div>"
    body += "</div></div>"
    return device(plat, raid_bar(plat) + body, "Raids", overlay=overlay, xl=xl, height=height)


def spec_sheet(plat: str) -> str:
    specs = [("spell_nature_lightning", "Elemental", "DPS"), ("spell_nature_lightningshield", "Enhancement", "DPS"), ("spell_nature_magicimmunity", "Restoration", "Healer")]
    rows = "".join(f"<div class=\"lrow\" style=\"min-height:64px\">{spec_icon(ic, 40)}<div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:600\">{n}</div><div class=\"t-meta muted\">{r}</div></div><span class=\"chev\">{glyph('chev-right', 18)}</span></div>" for ic, n, r in specs)
    inner = (f"<div class=\"t-body pad\" style=\"padding:0 20px 10px\">The pick signs <b class=\"strong\">Lightningbolltz</b> up as In for Onyxia's Lair, Thu 10 Dec, and is kept for the next raid.</div>"
             f"<div class=\"card quiet\" style=\"margin:0 16px\">{rows}</div><div class=\"t-meta muted\" style=\"padding:12px 20px\">Close the sheet and nothing is saved.</div>")
    return sheet(plat, inner, 430, "Which spec tonight?")


def raid_spec_ask(plat: str) -> str:
    return raid_detail(plat, "spec", overlay=spec_sheet(plat))


def raid_menu(plat: str) -> str:
    items = [("cal-plus", "Add to calendar"), ("share", "Share the raid"), ("globe", "Open on foreversixty.gg")]
    if plat == "ios":
        rows = "".join(f"<div style=\"display:flex;justify-content:space-between;align-items:center;min-height:48px;padding:0 16px;border-top:1px solid rgba(255,255,255,.08);font:400 17px -apple-system,'SF Pro Text';color:#fff\">{t}{glyph(g, 20)}</div>" for g, t in items)
        menu = f"<div class=\"glass\" style=\"position:absolute;top:118px;right:16px;width:260px;border-radius:22px;overflow:hidden;z-index:30\">{rows.replace('border-top:1px solid rgba(255,255,255,.08);', '', 1)}</div>"
    else:
        rows = "".join(f"<div style=\"display:flex;gap:14px;align-items:center;min-height:48px;padding:0 16px;font:400 16px Roboto;color:#e9e4d8\">{glyph(g, 20)}{t}</div>" for g, t in items)
        menu = f"<div style=\"position:absolute;top:96px;right:8px;width:250px;border-radius:12px;background:#1a1f2b;box-shadow:0 8px 24px rgba(0,0,0,.6);padding:8px 0;z-index:30\">{rows}</div>"
    return raid_detail(plat, overlay=menu)


def change_spec() -> str:
    return "<span class=\"t-meta gold\" style=\"font-weight:600;min-height:44px;display:flex;align-items:center\">Change spec</span>"


def officer_row(locked: bool) -> str:
    lock = "Unlock sign-ups" if locked else "Lock sign-ups"
    note = ("Locked. Everyone signed up was told the roster; moves you make now tell the raider at once." if locked
            else "Locking tells everyone signed up whether they are In or on the Bench.")
    return (f"<div class=\"pad\" style=\"margin-top:14px\"><div class=\"t-micro micro\">Officer</div><div style=\"display:flex;gap:8px;margin-top:6px\">"
            f"<span class=\"btn btn-secondary\">{glyph('lock', 16)}{lock}</span><span class=\"btn btn-secondary\">{glyph('paste', 16)}Copy the list</span></div>"
            f"<div class=\"t-meta muted\" style=\"margin-top:6px\">{note}</div></div>")


def officer_menu(plat: str, top: int) -> str:
    items = [("Move to Bench", False), ("Remove from the raid", True)]
    if plat == "ios":
        rows = "".join(f"<div style=\"display:flex;justify-content:space-between;align-items:center;min-height:48px;padding:0 16px;border-top:{'0' if i == 0 else '1px solid rgba(255,255,255,.08)'};font:400 17px -apple-system,'SF Pro Text';color:{'#ff6961' if red else '#fff'}\">{t}</div>" for i, (t, red) in enumerate(items))
        head_ = "<div style=\"padding:10px 16px 6px;font:400 13px -apple-system;color:#bbb\">Draknar · In as DPS</div>"
        return f"<div class=\"glass\" style=\"position:absolute;top:{top}px;right:16px;width:250px;border-radius:22px;overflow:hidden;z-index:30\">{head_}{rows}</div>"
    rows = "".join(f"<div style=\"display:flex;align-items:center;min-height:48px;padding:0 16px;font:400 16px Roboto;color:{'#f0736b' if red else '#e9e4d8'}\">{t}</div>" for t, red in items)
    return f"<div style=\"position:absolute;top:{top}px;right:12px;width:230px;border-radius:12px;background:#1a1f2b;box-shadow:0 8px 24px rgba(0,0,0,.6);padding:8px 0;z-index:30\">{rows}</div>"


def officer_roster(plat: str, menu: bool = False) -> str:
    top = content_top(plat)
    body = (f"<div style=\"padding-top:{top}px\"><div style=\"transform:translateY(-1040px)\">{raid_head(plat)}{officer_row(False)}{signup_card(note=False)}{class_strip()}"
            f"{roster(officer=True, open_name='Draknar' if menu else '')}</div></div>")
    overlay = officer_menu(plat, 316 if plat == "ios" else 300) if menu else ""
    return device(plat, raid_bar(plat) + body, "Raids", overlay=overlay)


def toast_locked(plat: str) -> str:
    y = 700 if plat == "ios" else 712
    return (f"<div class=\"toast\" role=\"status\" style=\"top:{y}px\"><span class=\"kill\">{glyph('lock', 18)}</span>"
            f"<span class=\"t-body\">Sign-ups locked. 23 In and 2 on the Bench were told.</span></div>")


# ---------- the after-raid summary ----------

def night(plat: str, premium: bool = True) -> str:
    from kit import item_icon, qname
    from screens_more import pushed
    loot = (f"<div class=\"lrow\" style=\"align-items:flex-start;padding:12px 16px\">{item_icon(17078, 44)}<div class=\"t\">{qname(17078)}"
            f"<div class=\"t-meta muted\">Back · Onyxia · awarded by Kraggor</div><div class=\"t-meta muted\">Replaces Arcanoweave Cloak, ilvl 65</div></div></div>")
    best = ("<div class=\"lrow\"><div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:600\">Onyxia</div><div class=\"t-meta muted\">Your best on this fight · 4:12 kill</div></div>"
            "<div style=\"text-align:right\"><div class=\"mono t-body strong\">612 DPS</div><div class=\"t-micro mono\" style=\"color:#a335ee\">82nd percentile</div></div></div>")
    coach = ""
    if premium:
        notes = ["Chain Lightning was off cooldown for 38 s of the 4:12 fight.", "Phase 2: you cast through Deep Breath once and took 3,210 damage.", "Mana ran dry at 3:40; a Major Mana Potion at 2:30 would cover it."]
        coach = ("<div class=\"t-micro micro pad\" style=\"margin-top:22px\">Coach notes</div><div class=\"card quiet\" style=\"margin:8px 16px 0;padding:6px 0\">"
                 + "".join(f"<div class=\"t-body\" style=\"padding:8px 16px;display:flex;gap:10px\"><span class=\"gold mono\">{i + 1}</span>{n}</div>" for i, n in enumerate(notes)) + "</div>")
    inner = (f"<div class=\"t-meta muted pad\" style=\"margin-top:4px\">Onyxia's Lair · Thursday 10 December · Olympus XXVII</div>"
             f"<div class=\"card\" style=\"margin:14px 16px 0;padding:16px;display:flex;gap:22px\"><div><div class=\"cond t-stat\">1</div><div class=\"t-meta muted\">kill</div></div>"
             f"<div><div class=\"cond t-stat\">0</div><div class=\"t-meta muted\">wipes</div></div><div><div class=\"cond t-stat\">1</div><div class=\"t-meta muted\">new best</div></div>"
             f"<div style=\"margin-left:auto;text-align:right\"><div class=\"t-body\" style=\"font-weight:700;color:{CLASS_COLOR['shaman']}\">Lightningbolltz</div><div class=\"t-meta muted\">Elemental · DPS</div></div></div>"
             f"<div class=\"t-micro micro pad\" style=\"margin-top:22px\">Loot awarded to you</div><div class=\"card quiet\" style=\"margin:8px 16px 0\">{loot}</div>"
             f"<div class=\"t-micro micro pad\" style=\"margin-top:22px\">Personal bests</div><div class=\"card quiet\" style=\"margin:8px 16px 0\">{best}</div>{coach}"
             f"<div class=\"pad\"><span class=\"btn btn-ghost\" style=\"padding:0;margin-top:10px\">Full report on foreversixty.gg{glyph('external', 14)}</span></div>")
    return pushed(plat, "Raids", "Your raid night", inner, "Raids")
