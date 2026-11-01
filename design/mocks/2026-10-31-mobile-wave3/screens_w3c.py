"""F14 widgets and live surfaces, F15 system integration, F16 guild extras."""
from w3kit import *  # noqa: F403
from w3kit import W3, TAB_CLEAR, buy_prices, money
from kit import SEAL, HORDE as HORDE_SRC
from model import ROSTER, SPEC, upgrades
from screens_more import APP_ICON, ios_lock, ios_note, and_shade, and_note
from screens_w3b import share_sheet

SHA = CLASS_COLOR["shaman"]
WALL = "radial-gradient(520px 420px at 78% 12%,rgba(214,110,40,.32),transparent 70%),radial-gradient(500px 500px at 10% 90%,rgba(63,143,224,.18),transparent 70%),linear-gradient(180deg,#14100b,#06080c)"


def wall(plat: str, inner: str, dock: bool = True, extra: str = "") -> str:
    top = DEVICES[plat]["top"]
    d = ""
    if dock:
        icons = "".join(f"<span style=\"width:58px;height:58px;border-radius:{14 if plat == 'ios' else 29}px;background:{c}\"></span>" for c in ["#2f8f4f", "#2f6fd0", "#8a5ad0", "#d0792f"])
        d = (f"<div class=\"glass\" style=\"position:absolute;left:14px;right:14px;bottom:{34 if plat == 'ios' else 30}px;height:88px;border-radius:{30 if plat == 'ios' else 34}px;display:flex;justify-content:space-around;align-items:center\">{icons}</div>")
    return device(plat, f"<div style=\"position:absolute;inset:0;background:{WALL}\"></div><div style=\"position:absolute;left:16px;right:16px;top:{top + 18}px\">{inner}</div>{d}{extra}", None)


def wg(plat: str, w: int, h: int, inner: str, bg: str = "") -> str:
    r = 22 if plat == "ios" else 28
    return (f"<div style=\"width:{w}px;height:{h}px;border-radius:{r}px;overflow:hidden;position:relative;padding:{14 if plat == 'ios' else 16}px;"
            f"background:{bg or 'linear-gradient(180deg,rgba(229,185,85,.09),rgba(229,185,85,.01) 45%),#10131b'};border:1px solid rgba(229,185,85,.16);box-shadow:0 10px 30px rgba(0,0,0,.5)\">{inner}</div>")


def seal(n: int = 18) -> str:
    return f"<img src=\"{SEAL}\" width=\"{n}\" height=\"{n}\" alt=\"\">"


def small_state(state: str) -> str:
    cd = "<div class=\"mono gold\" style=\"font-size:13px;margin-top:2px\">in 53 min</div>"
    head_ = lambda t="Next raid": f"<div style=\"display:flex;align-items:center;gap:6px\">{seal(16)}<span class=\"micro gold\" style=\"font-size:10px\">{t}</span></div>"  # noqa: E731
    zone = "<div class=\"cond\" style=\"font-size:22px;line-height:24px;margin-top:6px\">Onyxia's Lair</div><div class=\"muted\" style=\"font-size:13px\">Tonight 8:00 pm</div>"
    foot = lambda t: f"<div style=\"position:absolute;left:14px;bottom:12px\">{t}</div>"  # noqa: E731
    if state == "in":
        return head_() + zone + cd + foot(tag("In · DPS", "green"))
    if state == "open":
        return head_() + zone + cd + foot(tag("Not signed up", "gold"))
    if state == "signedout":
        return head_("Forever Sixty") + "<div style=\"font-size:15px;line-height:20px;margin-top:10px;font-weight:600\">Sign in to see your next raid.</div>" + foot("<span class=\"gold\" style=\"font-size:13px;font-weight:600\">Open the app</span>")
    if state == "noguild":
        return head_("Raids") + "<div style=\"font-size:15px;line-height:20px;margin-top:10px;font-weight:600\">No guild yet</div><div class=\"muted\" style=\"font-size:12px;margin-top:4px\">Your raids show here once you are on a roster.</div>"
    if state == "noraid":
        return head_() + "<div style=\"font-size:15px;line-height:20px;margin-top:10px;font-weight:600\">No raid planned</div><div class=\"muted\" style=\"font-size:12px;margin-top:4px\">Last night out: Lower Blackrock Spire, Sun 6 Dec.</div>"
    if state == "locked":
        return head_() + zone + cd + foot(tag("Locked", "muted"))
    if state == "cancelled":
        return head_() + zone.replace("Thu 8:00 pm", "Thu 8:00 pm") + foot(tag("Cancelled", "wipe"))
    if state == "stale":
        return head_() + zone + cd + foot("<span class=\"muted\" style=\"font-size:11px\">Updated 14 h ago</span>")
    return head_() + zone + foot(tag("Choose a spec", "gold"))


def nr_medium(plat: str, answered: str | None = "In") -> str:
    def b(t: str) -> str:
        on = t == answered
        return (f"<span style=\"flex:1;height:44px;border-radius:{10 if plat == 'ios' else 22}px;display:grid;place-items:center;font-weight:700;font-size:14px;"
                f"{'background:rgba(127,212,138,.18);color:var(--strong);box-shadow:inset 0 -2px 0 var(--kill)' if on else 'background:rgba(255,255,255,.07);color:var(--nav)'}\">{t}</span>")
    return (f"<div style=\"display:flex;gap:12px;height:100%\"><div style=\"flex:1;min-width:0;position:relative\">{small_state('open' if not answered else 'in').replace(tag('In · DPS', 'green'), '').replace(tag('Not signed up', 'gold'), '')}</div>"
            f"<div style=\"width:160px;display:flex;flex-direction:column;justify-content:space-between\"><div><div class=\"mono strong\" style=\"font-size:15px\">38 of 40</div><div class=\"muted\" style=\"font-size:12px\">signed up</div></div>"
            f"<div style=\"display:flex;gap:6px\">{b('In')}{b('Late')}{b('Out')}</div></div></div>")


def up_medium(plat: str) -> str:
    rows, _, total = upgrades()
    r = "".join(f"<div style=\"display:flex;align-items:center;gap:10px;margin-top:4px\">{item_icon(x['pick'], 28)}<div style=\"flex:1;min-width:0;font-size:13px;font-weight:700;color:{Q_TEXT[item(x['pick'])['quality']]};white-space:nowrap;overflow:hidden;text-overflow:ellipsis\">{item(x['pick'])['name']}</div>"
                 f"<span class=\"mono strong\" style=\"font-size:13px\">{'+' + str(x['gain']) if x['gain'] is not None else '—'}</span><span class=\"muted\" style=\"font-size:10px;margin-left:-6px\">DPS</span></div>" for x in rows[:3])
    return (f"<div style=\"display:flex;align-items:center;gap:8px\">{crest('shaman', 20)}<span class=\"micro muted\" style=\"font-size:10px\">Upgrades · Lightningbolltz</span></div>{r}"
            f"<div class=\"muted\" style=\"font-size:11px;margin-top:5px\"><span class=\"mono strong\">{len(rows)}</span> upgrades · <span class=\"mono strong\">+{total}</span> DPS</div>")


def up_small() -> str:
    rows, _, _ = upgrades()
    x = rows[0]
    it = item(x["pick"])
    return (f"<div style=\"display:flex;align-items:center;gap:6px\">{crest('shaman', 16)}<span class=\"micro muted\" style=\"font-size:10px\">Next upgrade</span></div>"
            f"<div style=\"margin-top:10px\">{item_icon(x['pick'], 44)}</div><div style=\"font-size:13px;font-weight:700;margin-top:8px;color:{Q_TEXT[it['quality']]};line-height:16px\">{it['name']}</div>"
            f"<div class=\"mono strong\" style=\"font-size:13px;margin-top:2px\">+{x['gain']} <span class=\"muted\">DPS</span></div>")


def home_widgets(plat: str) -> str:
    sm, md = (170, 170) if plat == "ios" else (164, 164)
    W = 370
    dots = "<div style=\"display:flex;justify-content:center;gap:8px;margin-top:20px\">" + "".join(f"<i style=\"width:7px;height:7px;border-radius:4px;background:{'#fff' if i == 0 else 'rgba(255,255,255,.35)'}\"></i>" for i in range(3)) + "</div>"
    inner = (f"<div style=\"display:flex;gap:{W - 2 * sm}px\">{wg(plat, sm, sm, small_state('in'))}{wg(plat, sm, sm, up_small())}</div>"
             f"<div style=\"margin-top:16px\">{wg(plat, W, md, nr_medium(plat))}</div><div style=\"margin-top:16px\">{wg(plat, W, md, up_medium(plat))}</div>")
    return wall(plat, inner + (dots if plat == "ios" else ""))


def widget_states(plat: str) -> str:
    items = [("Signed out", "signedout"), ("No guild", "noguild"), ("No raid planned", "noraid"), ("Sign-ups locked", "locked"), ("Cancelled", "cancelled"), ("Older than 12 h", "stale"), ("Not signed up", "open"), ("No spec saved", "nospec")]
    cells = "".join(f"<div><div class=\"t-micro muted\" style=\"margin-bottom:4px;color:#b9b3a4\">{n}</div>{wg(plat, 170, 150, small_state(s))}</div>" for n, s in items)
    return wall(plat, f"<div style=\"display:grid;grid-template-columns:170px 170px;gap:8px 30px;margin-top:-6px\">{cells}</div>", dock=False)


def lock_widgets() -> str:
    inline = "<div style=\"text-align:center;color:#fff;font:600 15px -apple-system,Barlow;opacity:.95;margin-top:6px\">Onyxia · Tonight 8 pm · In</div>"
    ring = lambda txt, frac, sub: (f"<span style=\"width:64px;height:64px;border-radius:32px;background:rgba(255,255,255,.15);display:grid;place-items:center;position:relative\"><svg width=\"64\" height=\"64\" style=\"position:absolute;inset:0\"><circle cx=\"32\" cy=\"32\" r=\"28\" fill=\"none\" stroke=\"rgba(255,255,255,.25)\" stroke-width=\"5\"/><circle cx=\"32\" cy=\"32\" r=\"28\" fill=\"none\" stroke=\"#fff\" stroke-width=\"5\" stroke-dasharray=\"{175.9 * frac:.0f} 176\" transform=\"rotate(-90 32 32)\" stroke-linecap=\"round\"/></svg><b style=\"color:#fff;font:700 18px/16px -apple-system,Barlow;position:relative;text-align:center\">{txt}<br><span style=\"font-size:10px;font-weight:600;opacity:.85\">{sub}</span></b></span>")  # noqa: E731
    rect = (f"<div style=\"width:170px;height:64px;border-radius:16px;background:rgba(255,255,255,.15);color:#fff;font-family:-apple-system,Barlow;padding:8px 12px\"><div style=\"font-weight:700;font-size:14px\">Onyxia's Lair</div>"
            f"<div style=\"font-size:13px;opacity:.9\">Tonight 8:00 pm</div><div style=\"font-size:13px;opacity:.9\">You: In · DPS</div></div>")
    up = (f"<div style=\"width:170px;height:78px;border-radius:16px;background:rgba(255,255,255,.15);color:#fff;font-family:-apple-system,Barlow;padding:8px 12px\"><div style=\"font-weight:700;font-size:13px\">Next upgrade · Waist</div>"
          f"<div style=\"font-size:13px;opacity:.95\">Knowledge of the Timbermaw</div><div style=\"font-size:13px;opacity:.9\">+5.3 DPS</div></div>")
    clock = ("<div style=\"text-align:center;color:#fff;font-family:-apple-system,Barlow,'SF Pro Display'\"><div style=\"font-size:20px;font-weight:600;opacity:.85\">Thursday 10 December</div>" + inline.replace("margin-top:6px", "margin-top:2px") +
             "<div style=\"font-size:104px;font-weight:700;line-height:1;letter-spacing:-2px\">7:07</div></div>")
    widgets = f"<div style=\"display:flex;gap:10px;justify-content:center;margin-top:10px\">{rect}{ring('38', .95, 'of 40')}</div><div style=\"display:flex;gap:10px;justify-content:center;margin-top:10px\">{up}{ring('53', .15, 'min')}</div>"
    la = live_card()
    body = (f"<div style=\"position:absolute;inset:0;background:{WALL}\"></div><div style=\"position:relative;padding-top:74px\">{clock}{widgets}</div>"
            f"<div style=\"position:absolute;left:10px;right:10px;bottom:120px\">{la}</div>")
    return device("ios", body, None)


def live_card() -> str:
    return (f"<div class=\"glass\" style=\"border-radius:26px;padding:16px;color:#fff;font-family:-apple-system,Barlow,'SF Pro Text'\"><div style=\"display:flex;align-items:center;gap:10px\"><img src=\"{APP_ICON}\" style=\"width:28px;height:28px;border-radius:7px\" alt=\"\">"
            f"<b style=\"font-size:15px;flex:1\">Onyxia's Lair</b><span style=\"font-size:13px;opacity:.8\">Olympus XXVII</span></div>"
            f"<div style=\"display:flex;align-items:baseline;gap:8px;margin-top:10px\"><span style=\"font:700 36px 'JetBrains Mono'\">38:00</span><span style=\"font-size:15px;opacity:.85\">until invites at 7:45 pm</span></div>"
            f"<div style=\"height:6px;border-radius:3px;background:rgba(255,255,255,.2);margin-top:10px\"><i style=\"display:block;height:6px;width:36%;border-radius:3px;background:#e5b955\"></i></div>"
            f"<div style=\"display:flex;justify-content:space-between;margin-top:10px;font-size:14px\"><span>You are In · DPS · Lightningbolltz</span><span style=\"opacity:.85\">38 of 40</span></div></div>")


def island() -> str:
    top = DEVICES["ios"]["top"]
    comp = (f"<div style=\"position:absolute;top:11px;left:50%;transform:translateX(-50%);width:230px;height:37px;border-radius:20px;background:#000;z-index:25;display:flex;align-items:center;justify-content:space-between;padding:0 12px\">"
            f"{seal(18)}<span style=\"font:700 14px 'JetBrains Mono';color:#e5b955\">38m</span></div>")
    exp = (f"<div style=\"position:absolute;top:86px;left:14px;right:14px;height:136px;border-radius:44px;background:#000;z-index:24;padding:20px 22px 0;color:#fff;font-family:-apple-system,Barlow\">"
           f"<div style=\"display:flex;justify-content:space-between;align-items:center\"><div><div style=\"font-weight:700;font-size:16px\">Onyxia's Lair</div><div style=\"font-size:13px;opacity:.75\">Invites 7:45 pm · You are In · DPS</div></div><span style=\"font:700 30px 'JetBrains Mono';color:#e5b955\">38:00</span></div>"
           f"<div style=\"display:flex;gap:10px;margin-top:14px\"><span style=\"flex:1;height:40px;border-radius:20px;background:#1f2430;display:grid;place-items:center;font-weight:600;font-size:14px\">Late</span><span style=\"flex:1;height:40px;border-radius:20px;background:#1f2430;display:grid;place-items:center;font-weight:600;font-size:14px\">Out</span></div></div>")
    sim = (f"<div class=\"glass\" style=\"position:absolute;left:10px;right:10px;top:290px;border-radius:26px;padding:16px;color:#fff;font-family:-apple-system,Barlow\"><div style=\"display:flex;align-items:center;gap:10px\">{tile('flask', 28)}<b style=\"font-size:15px;flex:1\">Elemental · Top Gear</b><span style=\"font-size:13px;opacity:.8\">Running</span></div>"
           f"<div style=\"display:flex;justify-content:space-between;margin-top:10px;font-size:14px\"><span><span class=\"mono\" style=\"font-weight:700\">40</span> of <span class=\"mono\">64</span> combinations</span><span style=\"opacity:.85\">about 2 min left</span></div>"
           f"<div style=\"height:8px;border-radius:4px;background:rgba(255,255,255,.2);margin-top:8px\"><i style=\"display:block;height:8px;width:62%;border-radius:4px;background:#e5b955\"></i></div></div>"
           f"<div class=\"glass\" style=\"position:absolute;left:10px;right:10px;top:420px;border-radius:26px;padding:16px;color:#fff;font-family:-apple-system,Barlow\"><div style=\"display:flex;align-items:center;gap:10px\">{tile('flask', 28)}<b style=\"font-size:15px;flex:1\">Elemental · Top Gear</b><span class=\"kill\" style=\"font-size:13px\">Finished</span></div>"
           f"<div style=\"display:flex;align-items:baseline;gap:8px;margin-top:8px\"><span style=\"font:700 30px 'JetBrains Mono'\">412</span><span style=\"font-size:14px;opacity:.85\">DPS · tap for the result</span></div></div>")
    lab = lambda y, t: f"<div style=\"position:absolute;left:18px;top:{y}px;color:#b9b3a4;font:12px 'JetBrains Mono'\">{t}</div>"  # noqa: E731
    body = f"<div style=\"position:absolute;inset:0;background:{WALL}\"></div>{lab(60, 'Compact (top): seal and minutes')}{lab(68, '') }{exp}{lab(232, 'Expanded: raid starting; Late and Out answer in place')}{lab(262, 'Lock screen: server sim running, then finished')}{lab(232, 'Server sim, running') if False else ''}{sim}"
    return device("ios", body + comp if False else body, None).replace("<div class=\"ios-di\"></div>", comp)


def and_widgets() -> str:
    return home_widgets("and")


def and_ongoing() -> str:
    cards = (and_note("Invites for Onyxia's Lair in 38 minutes", "Pull at 8:00 pm, in 53 minutes. You are In as DPS on Lightningbolltz. 38 of 40 signed up.", actions=["Late", "Out"], top=True)
             + and_note("Elemental · Top Gear", "40 of 64 combinations · about 2 min left", sub="Server sim", attach="")
             .replace("</div></div></div>", "</div></div></div><div style=\"height:6px;border-radius:3px;background:#2a3040;margin:10px 0 0 44px\"><i style=\"display:block;height:6px;width:62%;border-radius:3px;background:#e5b955\"></i></div>", 1)
             + and_note("Elemental · Top Gear finished", "412 DPS · tap for the result", sub="Server sim", bottom=True))
    qs = "".join(f"<span style=\"height:64px;border-radius:32px;background:{'#e5b955' if i == 0 else '#232836'};display:flex;align-items:center;gap:8px;padding:0 14px;color:{'#3a2a08' if i == 0 else '#e9e4d8'};font:500 13px Roboto\">{glyph('calendar', 20) if i == 0 else ''}{'Next raid<br>Thu 8 pm · In' if i == 0 else ''}</span>" for i in range(4))
    return device("and", f"<div style=\"position:absolute;inset:0;background:#0b0d12\"></div><div style=\"position:relative;padding:52px 20px 8px;color:#e9e4d8;font:500 14px Roboto\"><span style=\"font-size:16px\">Thu, Dec 10</span>"
                  f"<div style=\"display:grid;grid-template-columns:1fr 1fr;gap:10px;margin-top:14px\">{qs}</div></div><div style=\"position:relative;display:flex;flex-direction:column;gap:3px;padding:12px 12px 0\">{cards}</div>", None)


# ---------------- F15 ----------------

def siri(plat: str, kind: str = "signup") -> str:
    if kind == "signup":
        ask = "Sign me up for tonight's raid"
        card_ = (f"<div class=\"glass-w\" style=\"border-radius:22px;padding:16px;color:#fff;font-family:-apple-system,Barlow\"><div style=\"display:flex;gap:10px;align-items:center\"><img src=\"{APP_ICON}\" style=\"width:30px;height:30px;border-radius:7px\" alt=\"\"><b style=\"font-size:15px\">Forever Sixty</b></div>"
                 f"<div style=\"font-size:17px;line-height:23px;margin-top:10px\">Sign Lightningbolltz up as In for Onyxia's Lair, Thursday 8 pm?</div><div style=\"display:flex;gap:10px;margin-top:14px\"><span style=\"flex:1;height:46px;border-radius:23px;background:rgba(255,255,255,.14);display:grid;place-items:center;font-weight:600\">Cancel</span><span style=\"flex:1;height:46px;border-radius:23px;background:#e5b955;color:#0b0e14;display:grid;place-items:center;font-weight:700\">Sign up</span></div></div>")
        done = (f"<div class=\"glass-w\" style=\"border-radius:22px;padding:16px;color:#fff;font-family:-apple-system,Barlow;margin-top:10px\"><div style=\"font-size:17px;line-height:23px\"><b>You are In</b> for Onyxia's Lair, Thursday at 8 pm, as DPS on Lightningbolltz. 25 of 40 signed up.</div></div>")
        body = card_ + done
    else:
        ask = "What's my next upgrade"
        it = item(228190)
        body = (f"<div class=\"glass-w\" style=\"border-radius:22px;padding:16px;color:#fff;font-family:-apple-system,Barlow\"><div style=\"display:flex;gap:10px;align-items:center\"><img src=\"{APP_ICON}\" style=\"width:30px;height:30px;border-radius:7px\" alt=\"\"><b style=\"font-size:15px\">Forever Sixty</b></div>"
                f"<div style=\"display:flex;gap:12px;align-items:center;margin-top:12px\">{item_icon(228190, 52)}<div><div style=\"font-weight:700;font-size:17px;color:{Q_TEXT[it['quality']]}\">{it['name']}</div><div style=\"font-size:14px;opacity:.85\">Waist · item level {it['item_level']} · Vendor: Meilosh · {money(buy_prices()[228190])}</div></div></div>"
                f"<div style=\"font-size:15px;margin-top:10px\"><span class=\"mono\" style=\"font-weight:700\">+5.3</span> DPS for Lightningbolltz. Replaces Belt of the Archmage.</div></div>")
    ph = (f"<div style=\"position:absolute;inset:0;background:radial-gradient(420px 360px at 50% 100%,rgba(120,160,255,.45),transparent 70%),radial-gradient(360px 300px at 30% 100%,rgba(190,90,220,.35),transparent 70%),#05060a\"></div>"
          f"<div style=\"position:absolute;left:16px;right:16px;top:90px;color:#fff;font-family:-apple-system,Barlow\"><div style=\"font-size:15px;opacity:.7\">You said</div><div style=\"font-size:24px;font-weight:600;margin:4px 0 22px\">“{ask}”</div>{body}</div>"
          f"<div class=\"siri-glow\" style=\"position:absolute;left:50%;bottom:60px;transform:translateX(-50%);width:84px;height:84px;border-radius:42px;background:rgba(255,255,255,.08);display:grid;place-items:center;color:#fff\">{glyph('siri', 40)}</div>")
    return device("ios", ph, None)


def spotlight() -> str:
    def r(lead: str, t: str, s: str) -> str:
        return f"<div style=\"display:flex;gap:12px;align-items:center;padding:8px 14px\">{lead}<div style=\"flex:1\"><div style=\"font:600 16px -apple-system,Barlow;color:#fff\">{t}</div><div style=\"font:13px -apple-system,Barlow;color:#9a9aa2\">{s}</div></div></div>"
    sec = lambda t: f"<div style=\"font:600 13px -apple-system,Barlow;color:#9a9aa2;padding:14px 14px 4px\">{t}</div>"  # noqa: E731
    body = ("<div style=\"position:absolute;inset:0;background:#0b0c10\"></div><div style=\"position:absolute;top:70px;left:14px;right:14px\">"
            f"<div class=\"glass\" style=\"height:46px;border-radius:23px;display:flex;align-items:center;gap:8px;padding:0 14px;color:#fff;font-family:-apple-system,Barlow\">{glyph('search', 18)}<span style=\"font-size:17px\">lightning</span></div>"
            f"{sec('Forever Sixty')}"
            + r(crest("shaman", 38), "Lightningbolltz", "Elemental Shaman · 60 · Olympus XXVII")
            + r(sp("spell_nature_lightning", 38), "Elemental Shaman guide", "Guide · Updated 24 Sep 2026")
            + r(tile("calendar", 38), "Onyxia's Lair", "Raid · Thu 10 Dec, 8 pm")
            + r(sp("spell_nature_lightningoverload", 38), "Lightning Overload", "Talent · Elemental Shaman")
            + sec("Shortcuts")
            + r(tile("check", 38), "Sign me up for tonight's raid", "Forever Sixty")
            + r(tile("shield", 38), "What's my next upgrade", "Forever Sixty") + "</div>")
    return device("ios", body, None)


def launcher_menu() -> str:
    items = [("calendar", "Next raid"), ("check", "Sign up In"), ("shield", "My upgrades"), ("search", "Search")]
    menu = "".join(f"<div style=\"display:flex;align-items:center;gap:14px;height:56px;padding:0 18px;color:#e9e4d8;font:500 15px Roboto\">{tile(g, 36)}{t}</div>" for g, t in items)
    body = (f"<div style=\"position:absolute;inset:0;background:{WALL}\"></div><div style=\"position:absolute;left:24px;top:330px\"><div style=\"width:260px;border-radius:28px;background:#1d212c;box-shadow:0 12px 40px rgba(0,0,0,.6);overflow:hidden\">{menu}</div>"
            f"<img src=\"{APP_ICON}\" style=\"width:60px;height:60px;border-radius:30px;margin:14px 0 0 8px\" alt=\"\"><div style=\"color:#e9e4d8;font:500 12px Roboto;margin:4px 0 0 10px\">Forever Sixty</div></div>")
    return device("and", body, None)


def assistant() -> str:
    body = (f"<div style=\"position:absolute;inset:0;background:#0b0d12\"></div><div style=\"position:absolute;left:16px;right:16px;bottom:130px;color:#e9e4d8;font-family:Roboto\"><div style=\"font-size:15px;opacity:.7;margin-bottom:6px\">You said</div>"
            f"<div style=\"font-size:22px;margin-bottom:16px\">“Open my upgrades in Forever Sixty”</div><div style=\"border-radius:28px;background:#1a1f2b;padding:16px;display:flex;gap:12px;align-items:center\"><img src=\"{APP_ICON}\" style=\"width:44px;height:44px;border-radius:22px\" alt=\"\"><div><div style=\"font-weight:500\">Opening My upgrades</div><div style=\"font-size:13px;opacity:.7\">Forever Sixty</div></div></div></div>")
    return device("and", body, None)


def launcher_search() -> str:
    def r(lead: str, t: str, s: str) -> str:
        return f"<div style=\"display:flex;gap:12px;align-items:center;padding:10px 18px\">{lead}<div><div style=\"font:500 15px Roboto;color:#e9e4d8\">{t}</div><div style=\"font:13px Roboto;color:#9a9aa2\">{s}</div></div></div>"
    body = (f"<div style=\"position:absolute;inset:0;background:#0b0d12\"></div><div style=\"position:absolute;top:60px;left:12px;right:12px\"><div style=\"height:52px;border-radius:26px;background:#1a1f2b;display:flex;align-items:center;gap:12px;padding:0 18px;color:#e9e4d8;font:16px Roboto\">{glyph('search', 20)}lightning</div>"
            + r(crest("shaman", 40), "Lightningbolltz", "Forever Sixty · character") + r(crest("shaman", 40), "Voltaris", "Forever Sixty · watching") + r(tile("calendar", 40), "Next raid", "Forever Sixty · shortcut") + "</div>")
    return device("and", body, None)


def offline_data(plat: str) -> str:
    packs = [("book", "Guides", "Shaman, Warrior · 2.4 MB", "Updated today"), ("shield", "Best in slot lists", "Shaman · 3.1 MB", "Updated today"), ("tree", "Talents", "Shaman · 0.6 MB", "Updated today"),
             ("search", "Search index", "0.4 MB", "Updated today"), ("grid", "Item icons", "41 MB", "")]
    rows = "".join(f"<div class=\"lrow\" style=\"min-height:64px\">{tile(g)}<div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:600\">{n}</div><div class=\"t-meta muted\">{m}{' · ' + u if u else ''}</div></div></div>" for g, n, m, u in packs)
    tog = ("<div class=\"lrow\" style=\"min-height:64px\"><div class=\"t\"><div class=\"t-body\">Keep my classes ready</div><div class=\"t-meta muted\">Saves the guide, lists and talents for every class you have a character in.</div></div><span class=\"sw on\"></span></div>"
           "<div class=\"lrow\" style=\"min-height:64px\"><div class=\"t\"><div class=\"t-body\">Download on Wi-Fi only</div><div class=\"t-meta muted\">Icons and new guides wait for Wi-Fi.</div></div><span class=\"sw on\"></span></div>")
    inner = (label("On this phone, 47 MB", first=True) + card(rows) + label("Updates") + card(tog)
             + f"<div style=\"display:flex;flex-direction:column;gap:10px;padding:16px\"><span class=\"btn btn-primary full\">{glyph('refresh', 16)}Update now</span><span class=\"btn btn-secondary full\" style=\"color:var(--wipe)\">Remove downloaded data</span></div>"
             + note("Checked today at 9:41 pm. Your characters, sign-ups and sims are not stored here; they load when you are online."))
    return pushed(plat, "You", "Offline data", inner + TAB_CLEAR, "You")


def banner_states(plat: str) -> str:
    def b(inner: str, note_: str) -> str:
        return f"<div class=\"t-micro muted pad\" style=\"margin:16px 0 6px\">{note_}</div>{inner}"
    stale = f"<div class=\"offline t-meta\">{glyph('cloud', 16)}<span>Forever Sixty did not answer. Showing what you had at 9:41 pm.</span> <span class=\"gold\" style=\"font-weight:600;margin-left:auto\">Retry</span></div>"
    newdata = (f"<div class=\"offline t-meta\" style=\"background:rgba(229,185,85,.10);border-bottom-color:rgba(229,185,85,.3)\">{glyph('download', 16)}<span>New game data is ready.</span> <span class=\"gold\" style=\"font-weight:600;margin-left:auto\">Update</span></div>")
    inner = (b(offline_strip(), "No connection") + b(stale, "Connected, but the server did not answer and the cache is older than its limit")
             + b(newdata, "The game data changed while you were reading: nothing moves until you tap Update")
             + b(card("<div class=\"t-body\">This guide is not on your phone yet. Connect once to save it.</div>", margin="0 16px", pad="16px"), "Never saved and offline")
             + note("An offline strip never covers an error: a server error says what failed, in the failing region only."))
    return pushed(plat, "You", "Stale data states", inner, "You")


# ---------------- F16 ----------------

def roster_row(name: str, spec: str, lvl: int, rank: str, att: int, supp: bool = False, wait: bool = False) -> str:
    cls, sname, role, icon = SPEC[spec]
    segs = "".join(f"<i style=\"width:7px;height:12px;border-radius:2px;background:{'var(--kill)' if k < att else 'var(--line)'}\"></i>" for k in range(8))
    return (f"<div class=\"lrow\" style=\"min-height:64px\">{spec_mark(spec, 32)}<div class=\"t\"><div class=\"t-body\" style=\"font-weight:700;color:{CLASS_COLOR[cls]}\">{name}{' <span class=gold>' + glyph('crown', 13) + '</span>' if supp else ''}</div>"
            f"<div class=\"t-meta muted\">{sname} {cls.title()} · {lvl} · {rank}</div></div>"
            + (tag("Waiting", "muted") if wait else f"<div style=\"text-align:right\"><div style=\"display:flex;gap:2px;justify-content:flex-end\">{segs}</div><div class=\"mono t-micro strong\" style=\"margin-top:4px;letter-spacing:0\">{att} of 8</div></div>")
            + f"<span class=\"chev\">{glyph('chev-right', 18)}</span></div>")


def roster(plat: str, state: str = "default") -> str:
    people = [("Kraggor", "warrior-protection", "Guild Master", 8, True), ("Mendara", "priest-holy", "Officer", 8, False), ("Draknar", "warrior-fury", "Officer", 7, False), ("Lightningbolltz", "shaman-elemental", "Raider", 6, True),
              ("Rainspeaker", "shaman-restoration", "Raider", 6, False), ("Skarn", "rogue-combat", "Raider", 5, False), ("Zekra", "rogue-combat", "Raider", 5, False), ("Hexlash", "warlock-affliction", "Raider", 3, False)]
    head_ = card(f"<div style=\"display:flex;justify-content:space-between;padding:14px 16px\"><div><div class=\"cond t-stat\">38</div><div class=\"t-meta muted\">on the roster</div></div><div><div class=\"cond t-stat\">4</div><div class=\"t-meta muted\">in game, no account</div></div><div><div class=\"cond t-stat\">2</div><div class=\"t-meta muted\">waiting</div></div></div>", margin="14px 16px 0")
    chips = f"<div class=\"hscroll\" style=\"margin-top:12px\">{chip('All', True, False)}{chip('Tanks', chev=False)}{chip('Healers', chev=False)}{chip('DPS', chev=False)}{chip('Sort: Rank')}</div>"
    if state == "empty":
        inner = head_ + card("<div class=\"t-body strong\" style=\"font-size:17px\">No raiders match.</div><div class=\"t-meta muted\" style=\"margin-top:4px\">Change the role or clear the search.</div>", margin="14px 16px 0", pad="16px")
    else:
        rows = "".join(roster_row(n, s, 60, r, a, sup) for n, s, r, a, sup in people)
        inner = head_ + chips + card(rows, margin="12px 16px 0") + note("Attendance is the last 8 raid nights with a log. Tap a raider for their page.")
    right = icon_btn(plat, "search", "Search the roster")
    return pushed(plat, "Raids", "Roster", inner + TAB_CLEAR, "Raids", right=right)


def raid_history(plat: str) -> str:
    def r(zone: str, when: str, kills: str, raiders: int, length: str, you: bool = True, pressed: bool = False) -> str:
        return (f"<div class=\"lrow{' pressed' if pressed else ''}\" style=\"min-height:72px;align-items:flex-start;padding:12px 16px\"><div class=\"dateblk\" style=\"width:48px;height:52px\"><span class=\"m\">{when.split()[0]}</span><span class=\"d\" style=\"font-size:20px\">{when.split()[1]}</span><span class=\"m\">{when.split()[2]}</span></div>"
                f"<div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:600\">{zone}</div><div class=\"t-meta muted\">{kills}</div><div class=\"t-meta muted\" style=\"white-space:nowrap\"><span class=\"mono\">{raiders}</span> {'raiders' if raiders > 10 else 'players'} · <span class=\"mono\">{length}</span></div>"
                + (f"<div style=\"margin-top:6px\">{tag('You were there', 'green')}</div>" if you else "") + f"</div><span class=\"chev\">{glyph('chev-right', 18)}</span></div>")
    rows = (r("Onyxia's Lair", "THU 17 DEC", "1 boss · 1 kill · no wipes", 38, "1 h 47 m", pressed=True) + r("Blackrock Depths", "TUE 15 DEC", "4 bosses · 4 kills · no wipes", 5, "1 h 38 m", False)
            + r("Onyxia's Lair", "THU 10 DEC", "1 boss · 1 kill · 2 wipes", 39, "2 h 26 m") + r("Lower Blackrock Spire", "SUN 6 DEC", "3 bosses · 3 kills · no wipes", 5, "1 h 12 m"))
    inner = (label("Logged raid nights", first=True) + card(rows) + f"<div class=\"lrow\" style=\"min-height:52px;margin:0 16px\"><span class=\"t t-body gold\" style=\"font-weight:600\">Load more</span></div>"
             + note("A night opens your raid night when you were in it, otherwise its report."))
    return pushed(plat, "Raids", "Raid history", inner + TAB_CLEAR, "Raids")


def calendar_settings(plat: str, denied: bool = False) -> str:
    rows = ("<div class=\"lrow\" style=\"min-height:64px\"><div class=\"t\"><div class=\"t-body\">Raids in my calendar</div><div class=\"t-meta muted\">A calendar called Forever Sixty raids. The app keeps it up to date.</div></div><span class=\"sw" + ("" if denied else " on") + "\"></span></div>"
            "<div class=\"lrow\" style=\"min-height:56px\"><div class=\"t t-body\">Which raids</div><span class=\"val t-body\">All planned<span class=\"chev\">" + glyph('chev-down', 14) + "</span></span></div>"
            "<div class=\"lrow\" style=\"min-height:56px\"><div class=\"t t-body\">Alert</div><span class=\"val t-body\">1 hour before invites<span class=\"chev\">" + glyph('chev-down', 14) + "</span></span></div>")
    link = (f"<div class=\"lrow\" style=\"min-height:64px\">{tile('link')}<div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:600\">Calendar link</div><div class=\"t-meta muted\">For Google Calendar on the web, Outlook and other people. Anyone with the link sees your guild's raid times.</div></div></div>"
            f"<div style=\"display:flex;gap:10px;padding:0 16px 14px\"><span class=\"btn btn-secondary\" style=\"flex:1\">Copy link</span><span class=\"btn btn-secondary\" style=\"flex:1\">Reset link</span></div>")
    ask = "" if not denied else card("<div class=\"t-body wipe\" role=\"alert\">Calendar access is off for Forever Sixty.</div><div class=\"t-meta muted\" style=\"margin-top:4px\">Turn it on in Settings to add raids to your calendar.</div><div style=\"margin-top:10px\"><span class=\"btn btn-secondary\">Open Settings</span></div>", margin="14px 16px 0", pad="14px 16px")
    inner = ask + label("Calendar", first=True) + card(rows) + label("Share") + card(link) + note("Your sign-up shows in the event title: “Onyxia's Lair · In”. Changing it in the app changes the event.")
    return pushed(plat, "You", "Calendar", inner + TAB_CLEAR, "You")


def cal_event(plat: str) -> str:
    if plat == "ios":
        body = (f"<div style=\"position:absolute;inset:0;background:#000\"></div><div style=\"position:absolute;top:80px;left:0;right:0;color:#fff;font-family:-apple-system,Barlow\"><div style=\"padding:0 20px\"><div style=\"font-size:13px;color:#ff453a;font-weight:600\">Forever Sixty raids</div>"
                f"<div style=\"font:700 28px -apple-system,Barlow;margin-top:4px\">Onyxia's Lair · In</div><div style=\"font-size:17px;margin-top:6px\">Thursday, 10 December</div><div style=\"font-size:17px;opacity:.8\">7:45 pm to 11:00 pm</div></div>"
                f"<div style=\"margin:20px 16px 0;border-radius:14px;background:#1c1c1e\"><div style=\"padding:14px 16px;border-bottom:1px solid #2c2c2e\"><div style=\"font-size:13px;opacity:.6\">Location</div><div style=\"font-size:17px\">Olympus XXVII · Onyxia's Lair</div></div>"
                f"<div style=\"padding:14px 16px;border-bottom:1px solid #2c2c2e\"><div style=\"font-size:13px;opacity:.6\">Alert</div><div style=\"font-size:17px\">1 hour before</div></div><div style=\"padding:14px 16px\"><div style=\"font-size:13px;opacity:.6\">Notes</div><div style=\"font-size:17px\">Pull 8:00 pm. 24 of 40 signed up.<br><span style=\"color:#4aa3ff\">foreversixty.gg/app/raid/…</span></div></div></div></div>")
    else:
        body = (f"<div style=\"position:absolute;inset:0;background:#101319\"></div><div style=\"position:absolute;top:70px;left:0;right:0;color:#e9e4d8;font-family:Roboto;padding:0 20px\"><div style=\"font-size:13px;color:#e5b955\">Forever Sixty raids</div>"
                f"<div style=\"font-size:26px;margin-top:6px\">Onyxia's Lair · In</div><div style=\"font-size:16px;margin-top:8px\">Thursday, 10 December · 7:45 to 11:00 pm</div><div style=\"font-size:16px;opacity:.8;margin-top:16px\">Olympus XXVII · Onyxia's Lair</div>"
                f"<div style=\"font-size:16px;opacity:.8;margin-top:16px\">Notification · 1 hour before</div><div style=\"font-size:16px;opacity:.8;margin-top:16px\">Pull 8:00 pm. 24 of 40 signed up.<br><span style=\"color:#e5b955\">foreversixty.gg/app/raid/…</span></div></div>")
    return device(plat, body, None)
