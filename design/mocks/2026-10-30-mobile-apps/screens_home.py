"""F1 sign-in and the switcher, F2 the home answer and the tooltip sheet, F3 best in slot."""
from kit import (ALTS, CLASS_COLOR, HORDE, ME, Q_TEXT, Q_WORD, SEAL, back_btn, content_top, crest, device, glyph, icon_btn,
                 item, item_icon, offline_strip, qname, selector, sheet, spec_icon, top_bar)
from model import SLOT_LABEL, SLOT_ORDER, WORN, band, next_best, slots, source_line, stat_lines, upgrades


def fmt(x: float) -> str:
    return f"{x:+.1f}".replace("-", "−")


# ---------- F2 home ----------

def hero(xl: bool = False) -> str:
    rows, _, total = upgrades()
    figs = (f"<span style=\"display:inline-flex;align-items:baseline;gap:6px\"><span class=\"cond t-stat\">{len(rows)}</span><span class=\"t-body muted\">upgrades</span></span>"
            f"<span style=\"display:inline-flex;align-items:baseline;gap:6px\"><span class=\"cond t-stat\">+{total}</span><span class=\"t-body muted\">DPS</span></span>")
    return (f"<div style=\"background:var(--band);padding:14px 16px 18px;border-bottom:1px solid var(--line-soft)\">"
            f"<div class=\"t-meta muted\">{ME['desc']} · {ME['guild']}</div>"
            f"<div class=\"t-micro micro\" style=\"margin-top:14px\">Upgrades for your level range</div>"
            f"<div style=\"display:flex;flex-wrap:wrap;gap:4px 22px;margin-top:4px\">{figs}</div>"
            f"<div class=\"t-meta muted\" style=\"margin-top:6px\">Best in slot for 60, together · 2 gains not modelled yet</div>"
            f"<div class=\"t-meta\" style=\"margin-top:10px;color:var(--nav);display:flex;gap:6px;align-items:center\">{glyph('refresh', 14)}Synced 2 h ago from the addon</div></div>")


def upgrade_row(r: dict, xl: bool = False, pressed: bool = False) -> str:
    worn = item(r["worn"])
    gain = ""
    if r["gain"] is not None and not xl:
        gain = f"<div style=\"text-align:right;flex:none\"><div class=\"mono t-body strong\">+{r['gain']}</div><div class=\"t-micro muted\">DPS</div></div>"
    gain_line = ("<div class=\"t-meta muted\">Gain not modelled yet</div>" if r["gain"] is None
                 else (f"<div class=\"t-meta strong mono\">+{r['gain']} DPS</div>" if xl else ""))
    size = 56 if xl else 44
    return (f"<div class=\"lrow{' pressed' if pressed else ''}\" style=\"align-items:flex-start;padding:12px 16px\">{item_icon(r['pick'], size)}<div class=\"t\">{qname(r['pick'])}"
            f"<div class=\"t-meta muted\">{SLOT_LABEL[r['slot']]} · {r['source']}</div>"
            f"<div class=\"t-meta muted\">Replaces {worn['name']}, ilvl {worn['item_level']}</div>{gain_line}</div>{gain}</div>")


def upgrades_card(n: int = 5, xl: bool = False, pressed: int = -1) -> str:
    rows, already, _ = upgrades()
    body = "".join(upgrade_row(r, xl, i == pressed) for i, r in enumerate(rows[:n]))
    more = len(rows) - n
    return (f"<div class=\"card\" style=\"margin:14px 16px 0;overflow:hidden\"><div style=\"display:flex;justify-content:space-between;align-items:center;padding:12px 16px 4px\">"
            f"<span class=\"t-micro micro\">Your upgrades</span><span class=\"t-meta muted\">60 · Phase 1 raid</span></div>{body}"
            f"<div class=\"lrow\" style=\"min-height:52px\"><span class=\"t t-body gold\" style=\"font-weight:600\">{more} more upgrades · all 17 slots</span><span class=\"chev gold\">{glyph('chev-right', 18)}</span></div></div>"
            f"<div class=\"t-meta muted pad\" style=\"margin-top:10px\">Already best in slot: {', '.join(already)}.</div>")


def next_rows() -> str:
    rows = [(spec_icon("spell_nature_lightning", 36), "Talents", "<span class=\"tag tag-green\">Optimized</span>", "Matches the Elemental 60 build"),
            (f"<span class=\"itile\" style=\"width:36px;height:36px;border-radius:8px;display:grid;place-items:center;background:var(--tile);color:var(--tile-ink)\">{glyph('flask', 18)}</span>",
             "Sims", "", "Your Elemental sim finished: <span class=\"mono strong\">412 DPS</span>"),
            (f"<span class=\"itile\" style=\"width:36px;height:36px;border-radius:8px;display:grid;place-items:center;background:var(--tile);color:var(--tile-ink)\">{glyph('calendar', 18)}</span>",
             "Next raid", "<span class=\"tag tag-green\">In · DPS</span>", "Onyxia's Lair · Thursday 8 pm")]
    body = "".join(f"<div class=\"lrow\">{ic}<div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:600\">{lab}</div><div class=\"t-meta muted\">{sub}</div></div>{tag}<span class=\"chev\">{glyph('chev-right', 18)}</span></div>" for ic, lab, tag, sub in rows)
    return f"<div class=\"t-micro micro pad\" style=\"margin-top:22px\">Next for {ME['name']}</div><div class=\"card quiet\" style=\"margin:8px 16px 0\">{body}</div>"


def home(plat: str, state: str = "default", xl: bool = False) -> str:
    top = content_top(plat)
    bar = top_bar(plat, selector(plat, xl))
    if state == "loading":
        sk = lambda w, h, m=8: f"<div class=\"skel\" style=\"width:{w};height:{h}px;margin-top:{m}px\"></div>"  # noqa: E731
        body = (f"<div style=\"padding:{top + 14}px 16px 0;background:var(--band);padding-bottom:18px\">{sk('70%', 16, 0)}{sk('48%', 12, 18)}{sk('60%', 34)}{sk('80%', 14)}{sk('44%', 14, 12)}</div>"
                f"<div class=\"card quiet\" style=\"margin:14px 16px 0;padding:12px 16px\">" + "".join(
                    f"<div style=\"display:flex;gap:12px;padding:12px 0\"><div class=\"skel\" style=\"width:44px;height:44px\"></div><div style=\"flex:1\">{sk('70%', 15, 0)}{sk('52%', 12, 8)}{sk('60%', 12, 6)}</div></div>" for _ in range(5)) + "</div>"
                "<div class=\"t-meta muted pad\" style=\"margin-top:10px\" aria-live=\"polite\">Loading your upgrades…</div>")
        return device(plat, bar + body, "Home")
    if state == "nogear":
        body = (f"<div style=\"padding-top:{top}px\"><div style=\"background:var(--band);padding:14px 16px 22px;border-bottom:1px solid var(--line-soft)\">"
                f"<div class=\"t-meta muted\">Orc Fury Warrior · Level 41 · Olympus XXVII</div><div class=\"t-micro micro\" style=\"margin-top:14px\">Upgrades for your level range</div>"
                f"<div class=\"t-body strong\" style=\"margin-top:8px;font-size:17px\">Add the addon to see which slots you can upgrade.</div>"
                f"<div style=\"display:flex;gap:10px;margin-top:16px\"><span class=\"btn btn-primary\">Get the addon</span><span class=\"btn btn-secondary\">Fury 40 to 49 list</span></div>"
                f"<div class=\"t-meta muted\" style=\"margin-top:12px\">The addon sends your gear when you log out. Get the addon opens foreversixty.gg/setup in your browser.</div></div>"
                f"{next_rows_nogear()}</div>")
        return device(plat, top_bar(plat, selector(plat, xl, name="Grimtusk", cls="warrior", desc="Fury · 41")) + body, "Home")
    if state == "error":
        body = (f"<div style=\"padding-top:{top}px\"><div style=\"background:var(--band);padding:14px 16px 22px\"><div class=\"t-meta muted\">{ME['desc']} · {ME['guild']}</div>"
                f"<div class=\"t-micro micro\" style=\"margin-top:14px\">Upgrades for your level range</div>"
                f"<div class=\"t-body\" style=\"margin-top:8px;font-size:17px;color:var(--wipe)\" role=\"alert\">Your upgrades did not load.</div>"
                f"<span class=\"btn btn-secondary\" style=\"margin-top:14px\">{glyph('refresh', 16)}Try again</span></div>{next_rows()}</div>")
        return device(plat, bar + body, "Home")
    strip = offline_strip() if state == "offline" else ""
    body = f"<div style=\"padding-top:{top}px\">{strip}{hero(xl)}{upgrades_card(3 if xl else 5, xl, 0 if state == 'pressed' else -1)}{'' if xl else next_rows()}</div>"
    return device(plat, bar + body, "Home", xl=xl)


def next_rows_nogear() -> str:
    rows = [(spec_icon("ability_warrior_innerrage", 36), "Talents", "Not available yet: no talent export for this character. Open the addon once to send it."),
            (f"<span style=\"width:36px;height:36px;border-radius:8px;display:grid;place-items:center;background:var(--tile);color:var(--tile-ink)\">{glyph('flask', 18)}</span>", "Sims", "No sim yet.")]
    body = "".join(f"<div class=\"lrow\" style=\"align-items:flex-start\">{ic}<div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:600\">{lab}</div><div class=\"t-meta muted\">{sub}</div></div></div>" for ic, lab, sub in rows)
    return f"<div class=\"t-micro micro pad\" style=\"margin-top:22px\">Next for Grimtusk</div><div class=\"card quiet\" style=\"margin:8px 16px 0\">{body}</div>"


def tooltip_sheet(plat: str) -> str:
    pick, worn = 228190, WORN["waist"]
    it = item(pick)
    stats = "".join(f"<div class=\"ln t-body\">{s}</div>" for s in stat_lines(pick))
    alts = "".join(f"<div class=\"lrow\" style=\"padding:10px 0\">{item_icon(i, 36)}<div class=\"t\">{qname(i)}<div class=\"t-meta muted\">{'You wear this · ' if i == worn else ''}{fmt(d)} DPS vs the pick</div></div></div>" for i, _, d in next_best("waist"))
    inner = (f"<div style=\"padding:0 20px\"><div style=\"display:flex;gap:14px;align-items:center\">{item_icon(pick, 56)}<div><div class=\"t-sec\" style=\"font-weight:700;color:{Q_TEXT[it['quality']]}\">{it['name']}</div>"
             f"<div class=\"t-meta muted\">{Q_WORD[it['quality']]} · Waist · item level {it['item_level']}</div></div></div>"
             f"<div class=\"tip\" style=\"margin-top:14px\"><div class=\"ln t-body\">Waist</div><div class=\"ln t-body\">{it['armor']} Armor</div>{stats}"
             f"<div class=\"ln t-body\" style=\"margin-top:6px\">Requires Level {it['required_level']}</div><div class=\"t-body\" style=\"color:#e5cc80\">Item Level {it['item_level']}</div></div>"
             f"<div class=\"t-micro micro\" style=\"margin-top:16px\">Where it comes from</div><div class=\"t-body strong\" style=\"margin-top:4px\">Vendor: Meilosh</div>"
             f"<div class=\"t-micro micro\" style=\"margin-top:16px\">The next best</div>{alts}"
             f"<div class=\"t-meta muted\" style=\"margin-top:6px\">Best in slot for Waist · Elemental 60 · Phase 1 raid</div></div>")
    return sheet(plat, inner, 640)


def home_tooltip(plat: str) -> str:
    top = content_top(plat)
    body = top_bar(plat, selector(plat)) + f"<div style=\"padding-top:{top}px\">{hero()}{upgrades_card()}</div>"
    return device(plat, body, "Home", overlay=tooltip_sheet(plat))


def gear_full(plat: str) -> str:
    top = content_top(plat)
    rows = ""
    for key in SLOT_ORDER[:11]:
        s = slots().get(key)
        worn = WORN.get(key, s["item_id"])
        same = worn == s["item_id"]
        right = (f"<span class=\"tag tag-green\">{glyph('check', 12)}BiS</span>" if same else f"{glyph('chev-right', 14)}{item_icon(s['item_id'], 36)}")
        rows += (f"<div class=\"lrow\" style=\"gap:10px\"><span class=\"t-micro micro\" style=\"width:62px;flex:none\">{SLOT_LABEL[key]}</span>{item_icon(worn, 36)}"
                 f"<div class=\"t\" style=\"min-width:0\">{qname(worn, 't-meta')}<div class=\"t-micro muted\">ilvl {item(worn)['item_level']}</div></div><span style=\"display:flex;align-items:center;gap:6px;color:var(--muted)\">{right}</span></div>")
    title = "<span class=\"cond t-sec\" style=\"font-size:17px\">Your gear</span>"
    bar = top_bar(plat, back_btn(plat, "Home") + (title if plat == "and" else ""), "")
    head_ = "" if plat == "and" else f"<div class=\"cond t-title pad\" style=\"margin-top:6px\">Your gear</div>"
    body = bar + (f"<div style=\"padding-top:{top}px\">{head_}<div class=\"t-meta muted pad\" style=\"margin-top:4px\">What you wear, and the pick for 60 · Phase 1 raid. Tap an item for its tooltip.</div>"
                  f"<div class=\"card quiet\" style=\"margin:12px 16px 0\">{rows}</div></div>")
    return device(plat, body, "Home")


# ---------- F1 sign-in and the switcher ----------

def signed_out(plat: str, overlay: str = "") -> str:
    body = (f"<div style=\"position:absolute;inset:0;background:var(--band)\"></div><div style=\"position:relative;padding:{content_top(plat) + 60}px 24px 0;display:flex;flex-direction:column;align-items:flex-start\">"
            f"<img src=\"{SEAL}\" width=\"64\" height=\"64\" alt=\"\"><div class=\"wordmark\" style=\"font-size:22px;margin-top:16px\">Forever Sixty</div>"
            f"<div class=\"cond t-title\" style=\"margin-top:28px\">Play your class better.</div>"
            f"<div class=\"t-body\" style=\"margin-top:12px;font-size:17px;line-height:25px\">Sign in and your characters arrive with their gear, talents and guild.</div>"
            f"<span class=\"btn btn-primary full\" style=\"margin-top:28px;height:52px\">Sign in with Battle.net</span>"
            f"<span class=\"btn btn-secondary full\" style=\"margin-top:12px;height:52px\">Pick your class</span>"
            f"<div class=\"t-meta muted\" style=\"margin-top:16px\">Pick your class opens best in slot, talents and sims for any spec, no account needed.</div></div>")
    return device(plat, body, "Home", overlay=overlay)


def consent(plat: str) -> str:
    if plat == "ios":
        alert = ("<div class=\"scrim\" style=\"z-index:30\"></div><div style=\"position:absolute;left:50%;top:330px;transform:translateX(-50%);width:300px;border-radius:26px;z-index:31;overflow:hidden\" class=\"glass\">"
                 "<div style=\"padding:20px 18px 16px;text-align:center;font-family:-apple-system,'SF Pro Text',sans-serif;color:#fff\"><div style=\"font-weight:600;font-size:17px\">“Forever Sixty” Wants to Use “foreversixty.gg” to Sign In</div>"
                 "<div style=\"font-size:13px;margin-top:6px;color:#ddd\">This allows the app and website to share information about you.</div></div>"
                 "<div style=\"display:grid;grid-template-columns:1fr 1fr;gap:8px;padding:0 14px 14px\"><span style=\"height:44px;border-radius:22px;background:rgba(255,255,255,.12);display:flex;align-items:center;justify-content:center;color:#fff;font:600 17px -apple-system\">Cancel</span>"
                 "<span style=\"height:44px;border-radius:22px;background:#0a84ff;display:flex;align-items:center;justify-content:center;color:#fff;font:600 17px -apple-system\">Continue</span></div></div>")
        return signed_out(plat, alert)
    tab = ("<div style=\"position:absolute;inset:0;z-index:30;background:#fff\"><div style=\"padding:48px 16px 10px;background:#f1f1f1;display:flex;gap:10px;align-items:center;font:500 14px Roboto;color:#222\">"
           "<span style=\"font-size:20px\">✕</span><span><div style=\"font-weight:600\">Sign in with Battle.net</div><div style=\"color:#666;font-size:12px\">account.battle.net</div></span></div>"
           "<div style=\"padding:60px 28px;font:16px Roboto;color:#222\"><div style=\"font-size:24px;font-weight:600\">BATTLE.NET</div><div style=\"margin-top:28px;border:1px solid #ccc;border-radius:6px;height:48px;padding:12px;color:#888\">Email or phone</div>"
           "<div style=\"margin-top:12px;border:1px solid #ccc;border-radius:6px;height:48px;padding:12px;color:#888\">Password</div><div style=\"margin-top:20px;background:#148eff;color:#fff;border-radius:6px;height:48px;display:flex;align-items:center;justify-content:center\">Log In</div>"
           "<div style=\"margin-top:24px;color:#666;font-size:13px\">Custom Tab drawn by the system browser. Its content is Battle.net's, not ours.</div></div></div>")
    return signed_out(plat, tab)


def importing(plat: str) -> str:
    body = (f"<div style=\"position:absolute;inset:0;background:var(--band)\"></div><div style=\"position:relative;padding:{content_top(plat) + 120}px 24px 0;text-align:center\">"
            f"<img src=\"{SEAL}\" width=\"56\" height=\"56\" alt=\"\"><div class=\"cond t-h2\" style=\"margin-top:20px\">Bringing in your characters</div>"
            f"<div class=\"t-body\" style=\"margin-top:8px\">Reading Battle.net. This takes a few seconds.</div>"
            f"<div class=\"bar8\" style=\"width:180px;margin:24px auto 0\"><i style=\"width:55%\"></i></div></div>")
    return device(plat, body, None)


def signin_error(plat: str) -> str:
    note = ("<div class=\"toast\" role=\"alert\" style=\"top:%dpx;border-color:rgba(240,115,107,.45)\"><span class=\"wipe\">%s</span><span class=\"t-body\">Battle.net could not confirm that sign-in. Try again.</span></div>"
            % (DEVICES_TOP[plat] + 10, glyph('info', 18)))
    return signed_out(plat, note)


DEVICES_TOP = {"ios": 62, "and": 40}


def switcher(plat: str, xl: bool = False) -> str:
    top = content_top(plat)
    under = top_bar(plat, selector(plat, xl, open_=True)) + f"<div style=\"padding-top:{top}px\">{hero(xl)}{upgrades_card(3, xl)}</div>"
    rows = ""
    for i, (name, cls, desc, guild) in enumerate(ALTS):
        mark = f"<span class=\"gold\" style=\"display:flex\">{glyph('check', 20)}</span>" if i == 0 else ""
        stack = "flex-direction:column;align-items:flex-start;" if xl else ""
        rows += (f"<div class=\"lrow\" style=\"min-height:68px;{stack}{'background:var(--row);' if i == 0 else ''}\">{crest(cls, 44)}<div class=\"t\" style=\"overflow-wrap:anywhere\"><div class=\"t-name\" style=\"font-weight:700;color:{CLASS_COLOR[cls]}\">{name}</div>"
                 f"<div class=\"t-meta muted\">{desc}</div><div class=\"t-meta muted\">{guild}</div></div>{mark}</div>")
    inner = (f"<div class=\"t-meta muted\" style=\"padding:0 20px 8px\">{ME['battletag']} · Signed in with Battle.net</div>{rows}"
             f"<div class=\"t-meta muted\" style=\"padding:14px 20px 0\">A character shows here after it logs in on Forever, or after the addon syncs it.</div>")
    return device(plat, under, "Home", overlay=sheet(plat, inner, 600 if xl else 420, "Your characters"), xl=xl)


# ---------- F3 best in slot ----------

def chip(inner: str, on: bool = False) -> str:
    st = "border-color:var(--gold);background:rgba(229,185,85,.10)" if on else ""
    return f"<span style=\"display:inline-flex;align-items:center;gap:6px;min-height:40px;padding:0 12px 0 6px;border-radius:999px;border:1px solid var(--card-b);background:rgba(255,255,255,.04);{st};white-space:nowrap\">{inner}<span class=\"chev\">{glyph('chev-down', 14)}</span></span>"


def bis_row(key: str) -> str:
    s = slots()[key]
    worn = WORN.get(key, s["item_id"])
    nb = next_best(key)
    nxt = f"Next: {nb[0][1]}, {fmt(nb[0][2])} DPS" if nb else "Nothing else ranked"
    if nb and abs(nb[0][2]) < 0.05:
        nxt = f"Next: {nb[0][1]}, same DPS"
    mark = f"<span class=\"tag tag-green\" style=\"flex:none\">{glyph('check', 12)}You wear it</span>" if worn == s["item_id"] else ""
    unm = "<div class=\"t-meta muted\">Gain not modelled yet</div>" if s.get("effect_unmodelled") else ""
    return (f"<div class=\"lrow\" style=\"align-items:flex-start;padding:12px 16px\">{item_icon(s['item_id'], 44)}<div class=\"t\">{qname(s['item_id'])}"
            f"<div class=\"t-meta muted\">{SLOT_LABEL[key]} · {source_line(s)}</div><div class=\"t-meta muted\">{nxt}</div>{unm}</div>{mark}</div>")


def bis(plat: str, state: str = "default", overlay: str = "") -> str:
    top = content_top(plat)
    bar = top_bar(plat, selector(plat), icon_btn(plat, "search", "Search an item"))
    chips = (f"<div style=\"display:flex;gap:8px;padding:10px 16px 0;overflow:hidden\">{chip(spec_icon('spell_nature_lightning', 28) + '<span class=t-meta style=font-weight:600>Elemental Shaman</span>', True)}"
             f"{chip('<span class=t-meta style=font-weight:600;padding-left:6px>60 · Phase 1 raid</span>')}{chip(f'<img src={HORDE} width=24 height=24 alt=>' + '<span class=t-meta style=font-weight:600>Horde</span>')}</div>")
    head_ = (f"<div class=\"pad\" style=\"display:flex;align-items:baseline;gap:10px;margin-top:6px\"><span class=\"cond t-title\">Best in slot</span><span class=\"tag tag-gold\">Yours</span></div>{chips}"
             f"<div class=\"t-meta muted pad\" style=\"margin-top:10px\"><span class=\"mono strong\">{band()['set_dps']:.1f} DPS</span> at best in slot for your level range · Updated 29 Oct 2026</div>")
    strip = offline_strip("Thu 9:41 pm") if state == "offline" else ""
    rows = "".join(bis_row(k) for k in SLOT_ORDER[:7])
    body = bar + f"<div style=\"padding-top:{top}px\">{strip}{head_}<div class=\"card quiet\" style=\"margin:12px 16px 0\">{rows}</div></div>"
    return device(plat, body, "BiS", overlay=overlay)


def bis_picker(plat: str) -> str:
    from kit import CLASS_COLOR as CC
    classes = ["warrior", "paladin", "hunter", "rogue", "priest", "shaman", "mage", "warlock", "druid"]
    grid = "".join(f"<div style=\"display:flex;flex-direction:column;align-items:center;gap:6px;padding:10px 0;border-radius:12px;{'background:rgba(229,185,85,.10);box-shadow:inset 0 0 0 1px var(--gold);' if c == 'shaman' else ''}\">{crest(c, 44)}<span class=\"t-meta\" style=\"font-weight:600;color:{CC[c]}\">{c.title()}</span></div>" for c in classes)
    specs = [("spell_nature_lightning", "Elemental", "DPS", True), ("spell_nature_lightningshield", "Enhancement", "DPS", False), ("spell_nature_magicimmunity", "Restoration", "Healer", False)]
    srow = "".join(f"<span style=\"display:flex;align-items:center;gap:8px;min-height:48px;padding:0 12px 0 6px;border-radius:8px;border:1px solid {'var(--gold)' if on else 'var(--line)'};{'background:rgba(229,185,85,.10)' if on else ''}\">{spec_icon(ic, 32)}<span><div class=\"t-meta strong\" style=\"font-weight:600\">{n}</div><div class=\"t-micro muted\">{r}</div></span></span>" for ic, n, r, on in specs)
    bands = ["20 to 29", "30 to 39", "40 to 49", "50 to 59", "60", "60 · Phase 1 raid"]
    brow = "".join(f"<span class=\"t-meta\" style=\"display:inline-flex;align-items:center;min-height:40px;padding:0 12px;border-radius:999px;border:1px solid {'var(--gold)' if b.endswith('raid') else 'var(--line)'};{'background:rgba(229,185,85,.12);color:var(--gold);font-weight:700' if b.endswith('raid') else 'color:var(--nav)'}\">{b}</span>" for b in bands)
    inner = (f"<div style=\"padding:0 16px\"><div class=\"t-micro micro\">Class</div><div style=\"display:grid;grid-template-columns:repeat(3,1fr);gap:4px;margin-top:6px\">{grid}</div>"
             f"<div class=\"t-micro micro\" style=\"margin-top:14px\">Spec</div><div style=\"display:flex;flex-wrap:wrap;gap:8px;margin-top:6px\">{srow}</div>"
             f"<div class=\"t-micro micro\" style=\"margin-top:14px\">Level range</div><div style=\"display:flex;flex-wrap:wrap;gap:8px;margin-top:6px\">{brow}</div>"
             f"<div class=\"t-micro micro\" style=\"margin-top:14px\">Faction</div><div style=\"display:flex;gap:8px;margin-top:6px\"><span class=\"t-meta\" style=\"display:inline-flex;gap:8px;align-items:center;min-height:44px;padding:0 14px 0 8px;border-radius:999px;border:1px solid var(--line);color:var(--nav)\"><img src=\"{ALLIANCE_SRC()}\" width=26 height=26 alt=\"\">Alliance</span>"
             f"<span class=\"t-meta\" style=\"display:inline-flex;gap:8px;align-items:center;min-height:44px;padding:0 14px 0 8px;border-radius:999px;border:1px solid var(--gold);background:rgba(229,185,85,.12);color:var(--gold);font-weight:700\"><img src=\"{HORDE}\" width=26 height=26 alt=\"\">Horde</span></div>"
             f"<span class=\"btn btn-secondary full\" style=\"margin-top:16px\">Back to Lightningbolltz's list</span></div>")
    return bis(plat, overlay=sheet(plat, inner, 760, "Choose a list"))


def ALLIANCE_SRC() -> str:
    from kit import ALLIANCE
    return ALLIANCE
