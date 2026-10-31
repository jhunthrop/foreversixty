"""F9 guides and tier lists, F10 reports, F11 rankings and character pages."""
from w3kit import *  # noqa: F403
from w3kit import W3, CLASSES, TAB_CLEAR
from model import ROSTER, SPEC, WORN, slots
from screens_more import ranks_of
from kit import DATA

SHA = CLASS_COLOR["shaman"]


def hscroll(items: list[str], active: int = 0) -> str:
    return "<div class=\"hscroll\" style=\"margin-top:12px\">" + "".join(
        f"<span class=\"t-meta\" style=\"display:inline-flex;align-items:center;min-height:36px;padding:0 14px;border-radius:999px;border:1px solid {'var(--gold)' if i == active else 'var(--line)'};"
        f"{'background:rgba(229,185,85,.12);color:var(--gold);font-weight:700' if i == active else 'color:var(--nav)'};white-space:nowrap\">{t}</span>" for i, t in enumerate(items)) + "</div>"


# ---------------- F9 ----------------

def guides_root(plat: str, strip: str = "") -> str:
    your = (f"<div style=\"display:flex;gap:14px;align-items:center;padding:16px\">{sp('spell_nature_lightning', 56)}<div style=\"flex:1;min-width:0\">"
            f"<div class=\"t-micro micro gold\">Your guide</div><div class=\"cond t-sec\" style=\"font-size:24px;color:{SHA}\">Elemental Shaman</div>"
            f"<div class=\"t-meta muted\">DPS · Updated 24 Sep 2026</div></div><span class=\"chev\">{glyph('chev-right', 20)}</span></div>"
            f"<div style=\"padding:0 16px 14px;display:flex;gap:8px\">{tag(glyph('check', 12) + 'Saved on this phone', 'green')}</div>")
    grid = "".join(f"<div style=\"display:flex;flex-direction:column;align-items:center;gap:6px;padding:12px 0\">{crest(c, 44)}"
                   f"<span class=\"t-meta\" style=\"font-weight:600;color:{CLASS_COLOR[c]}\">{c.title()}</span></div>" for c in CLASSES)
    inner = (seg_n(["Guides", "Tier lists", "What changed"], 0) + card(your, quiet=False, margin="14px 16px 0")
             + label("All classes") + card(f"<div style=\"display:grid;grid-template-columns:repeat(3,1fr)\">{grid}</div>")
             + note("Saved on this phone: Shaman, Warrior · 2.4 MB. Guides update by themselves when the game data does."))
    return root(plat, "Guides", inner + TAB_CLEAR, "Guides", strip=strip)


def at_a_glance() -> str:
    stats = ["Intellect", "Spell power", "Hit", "Nature power", "Critical strike", "Spell haste", "Spell penetration"]
    chips = "".join(f"<span class=\"numchip t-meta\"><span class=\"n\">{i + 1}</span>{s}</span>" for i, s in enumerate(stats))
    return card(f"<div class=\"t-micro micro\">Role</div><div class=\"t-body strong\" style=\"margin-top:2px\">Ranged spell damage</div>"
                f"<div class=\"t-micro micro\" style=\"margin-top:14px\">Stat priority</div><div style=\"display:flex;flex-wrap:wrap;gap:6px;margin-top:8px\">{chips}</div>",
                quiet=False, margin="14px 16px 0", pad="16px")


def guide_head(plat: str) -> str:
    return (f"<div style=\"background:var(--band);padding:8px 16px 16px;border-bottom:1px solid var(--line-soft)\"><div style=\"display:flex;gap:14px;align-items:center\">{sp('spell_nature_lightning', 56)}"
            f"<div><div class=\"cond t-h2\" style=\"color:{SHA}\">Elemental Shaman</div><div class=\"t-meta muted\">Updated 24 Sep 2026 · Recommended races: Dwarf, Orc</div></div></div></div>")


def guide_page(plat: str, strip: str = "") -> str:
    over = ("<div class=\"t-micro micro pad\" style=\"margin-top:20px\">Overview</div>" + card(
        "<div class=\"t-body\">Elemental is Shaman's ranged spellcasting spec, dealing damage with Lightning Bolt, Chain Lightning, Flame Shock and Lava Burst, backed by fire totems.</div>"
        "<div class=\"t-micro micro\" style=\"margin-top:12px\">Strengths</div><div class=\"t-body\" style=\"margin-top:4px\">Forever's baseline cast-time cut to Lightning Bolt and Chain Lightning<br>Chain Lightning hits every target for close to single-target damage</div>"
        "<div class=\"t-micro micro\" style=\"margin-top:12px\">Weaknesses</div><div class=\"t-body\" style=\"margin-top:4px\">Slower leveling than Enhancement<br>Hard casts leave it idle while moving</div>", pad="14px 16px"))
    body = (guide_head(plat) + hscroll(["Overview", "Talents", "Rotation", "Stats", "Gear", "Races"], 0) + at_a_glance() + over)
    top = content_top(plat)
    left = back_btn(plat, "Guides") + (f"<span class=\"cond\" style=\"font-size:22px\">Elemental Shaman</span>" if plat == "and" else "")
    bar = top_bar(plat, f"<span style=\"display:flex;align-items:center;gap:4px\">{left}</span>", icon_btn(plat, "share", "Share"))
    return device(plat, bar + f"<div style=\"padding-top:{top}px\">{strip}{body}{TAB_CLEAR}</div>", "Guides")


KEY_TALENTS = [("Concussion", "Raises the damage of Lightning Bolt, Chain Lightning and Earth Shock."), ("Call of Flame", "Boosts Lava Burst and Fire Nova alongside Flame Shock."),
               ("Elemental Fury", "Raises critical strike damage on the spec's spells and totems."), ("Lightning Overload", "A chance to cast a second, weaker bolt for free."),
               ("Elemental Alacrity", "Cuts the cast time of Lightning Bolt, Chain Lightning and Lava Burst."), ("Lava Burst", "The capstone nuke; hits harder against Flame Shock.")]
ROT = [("spell_nature_manaregentotem", "Mana Spring Totem", "Mana Spring Totem", "Down first. It lasts five minutes."), ("spell_shaman_lavaburst", "Lava Burst", "Lava Burst", "On every cooldown, ahead of the totems."),
       ("spell_fire_searingtotem", "Searing Totem", "Searing Totem", "One target."), ("spell_fire_selfdestruct", "Magma Totem", "Magma Totem", "Two or more targets."),
       ("spell_nature_chainlightning", "Chain Lightning", "Chain Lightning", "Two or more targets."), ("spell_nature_lightning", "Lightning Bolt", "Lightning Bolt", "Fills every free global.")]


def guide_mid(plat: str) -> str:
    tal = {t["name"]: t for t in DATA["trees"][0]["talents"]}
    rows = "".join(f"<div class=\"lrow\" style=\"min-height:64px\">{sp(tal[n]['icon'], 40)}<div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:600\">{n} <span class=\"mono muted t-meta\">{tal[n]['max_rank']} {'rank' if tal[n]['max_rank'] == 1 else 'ranks'}</span></div><div class=\"t-meta muted\">{why}</div></div></div>" for n, why in KEY_TALENTS)
    g = ranks_of(DATA["bands"]["raid"]["talents"])
    tot = " / ".join(str(sum(t)) for t in g)
    build = (f"<div class=\"lrow\" style=\"min-height:60px\">{tile('tree')}<div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:600\">Raid build</div><div class=\"t-meta muted\"><span class=\"mono\">{tot}</span> · Level 60</div></div><span class=\"t-meta gold\" style=\"font-weight:600\">Compare</span><span class=\"chev\">{glyph('chev-right', 18)}</span></div>"
             f"<div class=\"lrow\" style=\"min-height:52px\"><span class=\"t t-body gold\" style=\"font-weight:600\">Open this build in the planner</span><span class=\"chev gold\">{glyph('chev-right', 18)}</span></div>")
    ranks = W3["ranks"]
    rot = "".join(f"<div class=\"lrow\" style=\"min-height:60px\"><span class=\"mono muted t-meta\" style=\"width:14px\">{i + 1}</span>{sp(ic, 40)}<div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:600\">{n} <span class=\"mono muted t-meta\">Rank {ranks[n]['max_rank']}</span></div><div class=\"t-meta muted\">{cond}</div></div></div>" for i, (ic, n, _, cond) in enumerate(ROT))
    inner = (hscroll(["Overview", "Talents", "Rotation", "Stats", "Gear", "Races"], 1) + label("Talents", extra="margin-top:18px") + card(rows + build)
             + label("Rotation") + card(rot) + note("Ranks are the ones you have at 60. Lava Burst is new in Forever."))
    top = content_top(plat)
    left = back_btn(plat, "Guides") + (f"<span class=\"cond\" style=\"font-size:22px\">Elemental Shaman</span>" if plat == "and" else "")
    bar = top_bar(plat, f"<span style=\"display:flex;align-items:center;gap:4px\">{left}</span>", icon_btn(plat, "share", "Share"))
    return device(plat, bar + f"<div style=\"padding-top:{top}px\">{inner}{TAB_CLEAR}</div>", "Guides")


TIER_COL = {"S": "#ff8000", "A": "#b866f5", "B": "#3d94f0", "C": "#1eff00", "D": "#9a9484"}
ORD = lambda n: f"{n}{'th' if 10 <= n % 100 <= 20 else {1: 'st', 2: 'nd', 3: 'rd'}.get(n % 10, 'th')}"  # noqa: E731


def tier_row(r: dict, mine: bool = False) -> str:
    cls, spec = r["class"], r["slug"]
    ties = "<span class=\"muted\" title=\"tie\">≈ </span>" if r["tie"] else ""
    return (f"<div class=\"lrow\" style=\"min-height:68px;gap:10px;{'background:var(--row);box-shadow:inset 3px 0 0 var(--gold)' if mine else ''}\"><span class=\"mono muted t-meta\" style=\"width:20px;text-align:right\">{r['rank']}</span>{spec_mark(spec, 36)}"
            f"<div class=\"t\"><div class=\"t-body\" style=\"font-weight:700;color:{CLASS_COLOR[cls]}\">{r['name']} {cls.title()}{' ' + tag('Yours') if mine else ''}</div>"
            f"<div class=\"t-meta muted\">{r['race'].title()}</div><div class=\"bar8\" style=\"margin-top:6px;width:100%;max-width:150px\"><i style=\"width:{r['fraction'] * 100:.0f}%;background:{TIER_COL[r['tier']]}\"></i></div></div>"
            f"<div style=\"text-align:right;flex:none\"><span class=\"tierchip\" style=\"color:{TIER_COL[r['tier']]}\">{r['tier']}</span><div class=\"mono t-meta strong\" style=\"margin-top:4px\">{ties}{'−' + str(r['gap']) + '%' if r['gap'] else 'Top'}</div></div></div>")


def tier_list(plat: str) -> str:
    rows = W3["tiers"]
    me = next(r for r in rows if r["slug"] == "shaman-elemental")
    top = rows[0]
    call = card(f"<div style=\"display:flex;gap:12px;align-items:center;padding:14px 16px\">{spec_mark('shaman-elemental', 40)}<div style=\"flex:1\"><div class=\"t-micro micro gold\">Your spec</div>"
                f"<div class=\"t-body strong\" style=\"font-weight:700\">Elemental Shaman is {ORD(me['rank'])} of {len(rows)} DPS specs.</div><div class=\"t-meta muted\">{me['gap']}% behind {top['name']} {top['class'].title()}.</div></div></div>", quiet=False, margin="14px 16px 0")
    chips = f"<div class=\"hscroll\" style=\"margin-top:12px\">{chip('DPS', True)}{chip('Tank', chev=False)}{chip('Healer', chev=False)}{chip(f'<img src={HORDE} width=22 height=22 alt=>Horde')}</div>"
    inner = (seg_n(["Guides", "Tier lists", "What changed"], 1) + chips + call
             + note("Ranked at best in slot with the Phase 1 raid buffs. Each spec is simmed as its best Horde race, named under it.")
             + card("".join(tier_row(r, r["slug"] == "shaman-elemental") for r in rows[:6])) + TAB_CLEAR)
    return root(plat, "Guides", inner, "Guides")


def whats_new(plat: str) -> str:
    dates = W3["dates"]
    drows = ""
    for d in dates[1:]:
        run = f"<div class=\"t-meta gold\">{d['running']}</div>" if d.get("running") else ""
        drows += (f"<div class=\"lrow\" style=\"min-height:60px\"><div class=\"dateblk\" style=\"width:56px;height:48px\"><span class=\"d\" style=\"font-size:15px\">{d['key']}</span></div>"
                  f"<div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:600\">{d['value']}</div><div class=\"t-meta muted\">{d.get('note', '')}</div>{run}</div></div>")
    ch = "".join(f"<div class=\"lrow\" style=\"min-height:64px;align-items:flex-start;padding:12px 16px\"><span class=\"mono muted t-meta\" style=\"width:44px\">{c['date'][5:].replace('-', '/')}</span><div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:600\">{c['title']}</div></div>{tag('Site', 'muted')}</div>" for c in W3["changelog"][:5])
    inner = seg_n(["Guides", "Tier lists", "What changed"], 2) + label("Key dates", first=True) + card(drows) + label("What changed") + card(ch) + TAB_CLEAR
    return root(plat, "Guides", inner, "Guides")


def class_page(plat: str) -> str:
    sp_rows = [("spell_nature_lightning", "Elemental", "DPS", True, "Updated 24 Sep"), ("spell_nature_lightningshield", "Enhancement", "DPS", False, ""), ("spell_nature_magicimmunity", "Restoration", "Healer", True, "Updated 24 Sep")]
    rows = "".join(f"<div class=\"lrow\" style=\"min-height:68px\">{sp(ic, 40)}<div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:600\">{n}</div><div class=\"t-meta muted\">{r}{' · ' + u if u else ''}</div></div>"
                   + (tag(glyph('check', 12) + 'Saved', 'green') if saved else f"<span class=\"t-meta nav\" style=\"display:flex;align-items:center;gap:6px;color:var(--nav)\">{glyph('cloud-down', 18)}Not saved</span>") + f"<span class=\"chev\">{glyph('chev-right', 18)}</span></div>" for ic, n, r, saved, u in sp_rows)
    top = content_top(plat)
    left = back_btn(plat, "Guides") + (f"<span class=\"cond\" style=\"font-size:22px\">Shaman</span>" if plat == "and" else "")
    head_ = "" if plat == "and" else f"<div class=\"cond t-title pad\" style=\"margin-top:4px;color:{SHA}\">Shaman</div>"
    inner = (offline_strip() + head_ + card(rows, margin="14px 16px 0") + note("Enhancement is not on your phone yet. Connect once to save it.", "color:var(--nav)"))
    return device(plat, top_bar(plat, f"<span style=\"display:flex;align-items:center;gap:4px\">{left}</span>") + f"<div style=\"padding-top:{top}px\">{inner}</div>", "Guides")


def guide_loading(plat: str) -> str:
    inner = (seg_n(["Guides", "Tier lists", "What changed"], 0) + "<div class=\"card quiet\" style=\"margin:14px 16px 0;padding:16px\">" + skel("40%", 14, 0) + skel("70%", 22) + skel("50%", 14) + "</div>"
             + "<div class=\"card quiet\" style=\"margin:22px 16px 0;padding:16px\">" + skel("100%", 120, 0) + "</div>" + note("Loading the guides…"))
    return root(plat, "Guides", inner, "Guides")


# ---------------- F10 ----------------

def rep_row(zone: str, meta: str, status: str = "", vis: str = "users", chev: bool = True, cls: str = "shaman", pressed: bool = False) -> str:
    return (f"<div class=\"lrow{' pressed' if pressed else ''}\" style=\"min-height:68px;gap:12px\">{crest(cls, 36)}<div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:600\">{zone}</div><div class=\"t-meta muted\">{meta}</div></div>"
            f"{status}<span class=\"muted\" style=\"display:flex\">{glyph(vis, 16)}</span><span class=\"chev\">{glyph('chev-right', 18)}</span></div>")


def reports_list(plat: str, state: str = "default") -> str:
    live = "<span class=\"tag tag-green\"><span class=\"dot\" style=\"background:var(--kill);width:7px;height:7px\"></span>Live</span>"
    if state == "empty":
        inner = card("<div class=\"t-body strong\" style=\"font-size:17px\">No reports yet. Upload a log, or let the companion send your next raid.</div>"
                     "<div style=\"margin-top:14px\"><span class=\"btn btn-secondary\">Get the companion</span></div>", margin="14px 16px 0", pad="18px 16px")
    elif state == "loading":
        inner = label("This week", first=True) + card("".join(f"<div class=\"lrow\" style=\"min-height:68px\"><span class=\"skel\" style=\"width:36px;height:36px;border-radius:18px\"></span><div class=\"t\">{skel('60%', 15, 0)}{skel('40%', 12, 8)}</div></div>" for _ in range(4)))
    elif state == "error":
        inner = card(fail("Your reports did not load."), margin="14px 16px 0")
    else:
        week = (rep_row("Onyxia's Lair", "Updated 20 s ago · 1 boss · 1 kill", live, "users", pressed=True) + rep_row("Onyxia's Lair", "Thu 10 Dec · 1 boss · 1 kill", "", "users")
                + rep_row("Onyxia's Lair", "Thu 3 Dec · 1 boss · 1 kill", "", "users")
                + rep_row("Blackrock Depths", "Tue 1 Dec · 4 bosses · 4 kills", "<span class=\"tag tag-gold\">Processing</span>", "lock"))
        old = rep_row("Lower Blackrock Spire", "Sun 29 Nov · 3 bosses · 3 kills", "", "lock") + rep_row("Stratholme", "Sun 22 Nov · could not be read", "<span class=\"tag tag-wipe\">Failed</span>", "lock", cls="shaman")
        inner = label("This week", first=True) + card(week) + label("Earlier") + card(old) + note("Newest first. Logs come from the companion on your PC; the app reads them, it does not upload.")
    return pushed(plat, "Home", "Your reports", inner + TAB_CLEAR, "Home", right=icon_btn(plat, "refresh", "Refresh"))


FIGHTS = [("wipe", "Onyxia", "Wipe · 1:12 · boss at 61%", ""), ("wipe", "Onyxia", "Wipe · 2:41 · boss at 22%", ""), ("kill", "Onyxia", "Kill · 3:12 · 40 players · 1 death", "")]


def fights_card() -> str:
    rows = "".join(f"<div class=\"lrow\" style=\"min-height:60px\"><span class=\"dot\" style=\"background:{'var(--kill)' if k == 'kill' else 'var(--wipe)'}\"></span><div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:600\">{n}</div>"
                   f"<div class=\"t-meta muted\">{m}</div></div><span class=\"chev\">{glyph('chev-right', 18)}</span></div>" for k, n, m, _ in FIGHTS)
    rows += (f"<div class=\"lrow\" style=\"min-height:60px\"><span class=\"dot\" style=\"background:var(--muted)\"></span><div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:600\">Trash · Onyxia's Lair</div>"
             f"<div class=\"t-meta muted\">9 pulls · 14:20</div></div><span class=\"chev\">{glyph('chev-right', 18)}</span></div>")
    return rows


def parses_card() -> str:
    return card(f"<div style=\"padding:14px 16px 4px\"><div class=\"t-micro micro gold\">Your parses</div></div>"
                f"<div class=\"lrow\" style=\"min-height:72px;border-top:0\">{crest('shaman', 36)}<div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:700\">Onyxia</div><div class=\"t-meta muted\">Kill 3:12 · Lightningbolltz, Elemental</div></div>"
                f"<div style=\"text-align:right\"><div class=\"mono t-sec strong\" style=\"font-size:20px\">412 <span class=\"t-meta muted\">DPS</span></div><div class=\"t-meta\">{pct(78)} <span class=\"muted\">percentile</span></div></div></div>", quiet=False, margin="14px 16px 0")


def report_head(plat: str, live: bool = False) -> str:
    return (f"<div class=\"pad\"><div style=\"display:flex;align-items:center;gap:8px;margin-top:4px\"><span class=\"t-micro micro gold\">Report</span>{tag('Guild', 'muted')}</div>"
            f"<div class=\"t-meta muted\" style=\"margin-top:4px\">Thursday 3 December 2026 · logged by Lightningbolltz</div></div>")


def coach_card() -> str:
    return card(f"<div style=\"padding:14px 16px\"><div style=\"display:flex;justify-content:space-between\"><span class=\"cond t-sec\">Your Elemental coach</span>{tag('Premium', 'gold')}</div>"
                f"<div class=\"t-micro micro\" style=\"margin-top:12px\">Fix first</div><div class=\"t-body\" style=\"margin-top:2px\">Mana Spring Totem was down for 41 seconds of the kill. Drop it before the pull.</div>"
                f"<div class=\"t-micro micro\" style=\"margin-top:12px\">Practice</div><div class=\"t-body\" style=\"margin-top:2px\">Cast Lava Burst on cooldown. 3 of 11 casts waited more than 2 seconds.</div>"
                f"<div class=\"t-micro micro\" style=\"margin-top:12px\">Doing well</div><div class=\"t-body\" style=\"margin-top:2px\">You never stood in the breath.</div>"
                f"<div class=\"t-meta muted\" style=\"margin-top:12px\">Reviewed 3 Dec · Based on: this log, the Elemental guide.</div></div>", margin="8px 16px 0")


def report_page(plat: str, scroll: int = 0, overlay: str = "", state: str = "default") -> str:
    top = content_top(plat)
    left = back_btn(plat, "Reports") + (f"<span class=\"cond\" style=\"font-size:22px\">Onyxia's Lair</span>" if plat == "and" else "")
    bar = top_bar(plat, f"<span style=\"display:flex;align-items:center;gap:4px\">{left}</span>", icon_btn(plat, "external", "Open on foreversixty.gg"))
    title = "" if plat == "and" else "<div class=\"cond t-title pad\" style=\"margin-top:4px\">Onyxia's Lair</div>"
    body = (title + report_head(plat) + parses_card() + label("Fights") .replace("</div>", "<span style=\"float:right\" class=\"mono\">3 boss pulls</span></div>", 1)
            + card(fights_card()) + label("Coach") + coach_card()
            + note("Meters, casts and deaths for every fight are on foreversixty.gg. Tap Open on foreversixty.gg."))
    cover = f"<div style=\"position:absolute;top:0;left:0;right:0;height:{DEVICES[plat]['top']}px;background:var(--bg);z-index:11\"></div>"
    return device(plat, cover + bar + f"<div style=\"padding-top:{top}px\"><div style=\"transform:translateY(-{scroll}px)\">{body}{TAB_CLEAR}</div></div>", "Home", overlay=overlay)


def meter_rows() -> str:
    names = [n for n, _, _ in ROSTER["DPS"]]
    vals = [596, 571, 548, 533, 497, 412, 401, 388, 372, 355, 341]
    order = [n for n in names if n != "Lightningbolltz"][:5] + ["Lightningbolltz"] + [n for n in names if n != "Lightningbolltz"][5:10]
    out = ""
    for i, n in enumerate(order):
        spec = dict((a, b) for a, b, _ in ROSTER["DPS"])[n]
        cls = SPEC[spec][0]
        me = n == "Lightningbolltz"
        out += (f"<div class=\"meter\" style=\"{'background:var(--row);box-shadow:inset 3px 0 0 var(--gold)' if me else ''}\"><span class=\"fill\" style=\"width:{vals[i] / vals[0] * 100:.0f}%;background:{CLASS_COLOR[cls]}\"></span>"
                f"<span class=\"mono muted t-meta\" style=\"width:18px\">{i + 1}</span>{spec_mark(spec, 28)}<span class=\"t-body\" style=\"font-weight:{700 if me else 600};color:{CLASS_COLOR[cls]};flex:1\">{n}</span>"
                f"<span class=\"mono t-body strong\">{vals[i]}</span><span class=\"t-micro muted\" style=\"width:26px\">DPS</span></div>")
    return out


def fight_sheet(plat: str) -> str:
    inner = (f"<div style=\"padding:0 16px\"><div style=\"display:flex;align-items:center;gap:10px\"><span class=\"dot\" style=\"background:var(--kill)\"></span><span class=\"t-body strong\" style=\"font-weight:700\">Onyxia</span>"
             f"<span class=\"t-meta muted\">Kill · 3:12 · 40 players</span></div></div>{seg_n(['Damage', 'Healing', 'Damage taken'], 0)}"
             f"<div style=\"margin-top:8px\">{meter_rows()}</div><div class=\"lrow\" style=\"min-height:48px;border-top:1px solid var(--line-soft)\"><span class=\"t t-body gold\" style=\"font-weight:600\">Show all 40</span></div>")
    return sheet(plat, inner, 700, "Damage done")


def report_states(plat: str) -> str:
    def s(t: str, b: str, tone: str = "") -> str:
        return card(f"<div class=\"t-body {tone}\" style=\"font-size:17px;font-weight:600\">{t}</div><div class=\"t-meta muted\" style=\"margin-top:4px\">{b}</div>", margin="10px 16px 0", pad="14px 16px")
    inner = (label("Live", first=True) + s("Live. Updated 20 s ago.", "The fight in progress shows when it ends.", "kill") + label("Processing") + s("This report is still being read.", "It opens here when it is ready. You do not need to stay.")
             + label("Failed") + s("This report could not be read.", "Upload the log again from the companion.", "wipe") + label("Not yours to see") + s("This report is not available to you.", "It is private, or it was removed.")
             + label("Offline") + card(offline_strip() .replace("padding:6px 16px", "padding:6px 16px"), margin="10px 16px 0"))
    return pushed(plat, "Reports", "States", inner, "Home")


# ---------------- F11 ----------------

NAMES = ["Voltaris", "Stormwrath", "Zulthrax", "Crackleheim", "Thundergaze", "Ashbolt", "Hexwhirl", "Surgekin", "Overcharge", "Galewind", "Bolthammer", "Emberstorm", "Skyrend", "Lightningbolltz"]
GUILDS = ["Hordebreakers", "Nightfall Raiders", "Olympus XXVII", "Iron Pact", "Blood & Thunder", "Redridge Rejects", "Skullsplitters"]


def rank_row(i: int, name: str, own: bool = False, pin: bool = False) -> str:
    val = 548 - int(i ** 0.8 * 5)
    pcol = parse_color(100 - (i / 212 * 100)) if i < 200 else parse_color(15)
    n = i + 1
    return (f"<div class=\"lrow\" style=\"min-height:68px;gap:12px;{'background:var(--row);box-shadow:inset 3px 0 0 var(--gold);' if own else ''}\"><span class=\"mono t-body\" style=\"width:30px;text-align:right;font-weight:700;color:{pcol}\">{n}</span>"
            f"<div class=\"t\"><div class=\"t-body\" style=\"font-weight:700;color:{SHA}\">{name}{' ' + tag('You') if own else ''}</div><div class=\"t-meta muted\">{GUILDS[i % len(GUILDS)]} · 40 · 3:{12 + i % 40:02d} · {3 + i} d ago</div></div>"
            f"<div style=\"text-align:right\"><div class=\"mono t-body strong\">{val}</div><div class=\"t-micro muted\">DPS</div></div><span class=\"chev\">{glyph('chev-right', 18)}</span></div>")


def rankings(plat: str, state: str = "default") -> str:
    ans = card(f"<div style=\"padding:14px 16px\"><div class=\"t-micro micro\">Your rank</div><div style=\"display:flex;align-items:baseline;gap:8px;margin-top:2px\"><span class=\"cond t-stat\">14th</span>"
               f"<span class=\"t-body muted\">of 212 Elemental Shaman characters</span></div><div class=\"t-meta muted\" style=\"margin-top:2px\">Your best kill on Onyxia, Normal, Phase 1.</div></div>", quiet=False, margin="14px 16px 0")
    summ = (f"<div style=\"display:flex;align-items:center;gap:10px;padding:12px 16px 0\"><div class=\"t-meta nav\" style=\"flex:1;color:var(--nav)\">Onyxia · Normal · All regions · Phase 1 · Damage</div>"
            f"<span class=\"btn btn-secondary\" style=\"min-height:44px\">{glyph('sliders', 16)}Filters</span></div>")
    if state == "pinned":
        ans = ans.replace("14th", "204th").replace("of 212", "of 212")
        pin = (f"<div class=\"t-micro micro pad\" style=\"margin-top:16px\">Your best kill</div>" + card(rank_row(203, "Lightningbolltz", own=True), margin="6px 16px 0")
               + "<div style=\"height:1px;background:var(--line-soft);margin:14px 16px 0\"></div>")
        rows = card("".join(rank_row(i, NAMES[i]) for i in range(5)), margin="14px 16px 0")
        inner = ans + summ + pin + rows
    elif state == "empty":
        inner = ans.replace("14th", "—").replace("of 212 Elemental Shaman characters", "Not ranked yet on Onyxia") + summ + card("<div class=\"t-body\">No ranked fights yet for this filter. Rankings fill in as reports are uploaded.</div>", margin="14px 16px 0", pad="18px 16px")
    else:
        rows = "".join(rank_row(i, NAMES[i], own=NAMES[i] == "Lightningbolltz") for i in range(7))
        jump = f"<div class=\"pad\" style=\"margin-top:10px\"><span class=\"t-meta gold\" style=\"font-weight:600;display:inline-flex;min-height:44px;align-items:center\">Jump to your row, 14th</span></div>"
        inner = ans + summ + jump + card(rows, margin="4px 16px 0")
    return pushed(plat, "Home", "Rankings", inner + TAB_CLEAR, "Home")


def filters_sheet(plat: str) -> str:
    def row_(title: str, opts: list[str], on: int) -> str:
        c = "".join(f"<span class=\"t-meta\" style=\"display:inline-flex;align-items:center;min-height:40px;padding:0 12px;border-radius:999px;border:1px solid {'var(--gold)' if i == on else 'var(--line)'};{'background:rgba(229,185,85,.12);color:var(--gold);font-weight:700' if i == on else 'color:var(--nav)'}\">{o}</span>" for i, o in enumerate(opts))
        return f"<div class=\"t-micro micro\" style=\"margin-top:14px\">{title}</div><div style=\"display:flex;flex-wrap:wrap;gap:8px;margin-top:6px\">{c}</div>"
    inner = (f"<div style=\"padding:0 16px\"><div class=\"t-micro micro\">Encounter</div><div class=\"lrow\" style=\"border:1px solid var(--line);border-radius:8px;margin-top:6px;min-height:52px\"><span class=\"t t-body strong\" style=\"font-weight:600\">Onyxia</span><span class=\"chev\">{glyph('chev-down', 16)}</span></div>"
             + row_("Measure", ["Damage", "Healing", "Damage taken", "Execution"], 0) + row_("Ruleset", ["Normal", "PvP", "RP", "Hardcore", "Every ruleset"], 0)
             + row_("Region", ["All", "US", "EU", "KR", "TW", "CN"], 0) + row_("Phase", ["Phase 1", "Launch", "Beta"], 0) + row_("Faction", ["Both", "Alliance", "Horde"], 0)
             + "<div class=\"lrow\" style=\"padding:0;margin-top:10px\"><span class=\"t t-body\">Today only</span><span class=\"sw\"></span></div>"
             + "<div style=\"display:flex;gap:10px;margin-top:10px\"><span class=\"btn btn-secondary\">Reset filters</span></div></div>")
    return rankings(plat, "default").replace("</div></div></div></div></div></div>", "</div></div></div></div></div></div>") if False else _over(plat, rankings(plat), inner, "Filters", 780)


def _over(plat: str, base: str, inner: str, title: str, h: int) -> str:
    """Re-draws base with a sheet on top: the base device is rebuilt by the caller when it needs the scrim, so the sheet is appended inside it."""
    return base.replace("<div class=\"ios-home\"></div>", sheet(plat, inner, h, title) + "<div class=\"ios-home\"></div>").replace("<div class=\"and-gest\"></div>", sheet(plat, inner, h, title) + "<div class=\"and-gest\"></div>")


def score_card() -> str:
    comps = [("Gear", 412, 500), ("Experience", 620, 700), ("Parses", 540, 600), ("Attendance", 180, 200), ("Discipline", 78, 100)]
    bars = "".join(f"<div style=\"display:flex;align-items:center;gap:10px;margin-top:8px\"><span class=\"t-meta muted\" style=\"width:86px\">{n}</span><div class=\"bar8\" style=\"flex:1\"><i style=\"width:{v / m * 100:.0f}%;background:var(--tile-ink)\"></i></div><span class=\"mono t-meta strong\" style=\"width:34px;text-align:right\">{v}</span></div>" for n, v, m in comps)
    spark = "<svg width=\"96\" height=\"34\" viewBox=\"0 0 96 34\" fill=\"none\" stroke=\"#b866f5\" stroke-width=\"2\" stroke-linecap=\"round\"><path d=\"M2 30 L18 26 L34 27 L50 18 L66 14 L82 9 L94 5\"/></svg>"
    return card(f"<div style=\"padding:14px 16px\"><div style=\"display:flex;justify-content:space-between;align-items:flex-start\"><div><div class=\"t-micro micro\">Forever Score</div><div style=\"display:flex;align-items:baseline;gap:8px\"><span class=\"cond t-stat\" style=\"font-size:36px\">1,830</span>"
                f"<span class=\"tag\" style=\"color:#b866f5;border-color:#b866f5;background:rgba(163,53,238,.12)\">Epic</span></div><div class=\"t-meta muted\">As of 4 Nov</div></div>{spark}</div>{bars}"
                f"<div class=\"t-meta gold\" style=\"margin-top:12px;font-weight:600\">How is this calculated?</div></div>", quiet=False, margin="14px 16px 0")


def char_head() -> str:
    return (f"<div class=\"pad\" style=\"display:flex;gap:14px;align-items:center;margin-top:4px\">{crest('shaman', 64)}<div style=\"flex:1;min-width:0\"><div class=\"cond t-h2\" style=\"color:{SHA}\">Voltaris</div>"
            f"<div class=\"t-meta muted\">Orc Elemental Shaman · Level 60</div><div class=\"t-meta muted\">Hordebreakers · US Normal</div></div></div>")


def bests_card() -> str:
    rows = [("Onyxia", "Normal · 3 d ago", "541", 97), ("Barrow Deeps", "Normal · 9 d ago", "498", 81)]
    r = "".join(f"<div class=\"lrow\" style=\"min-height:64px\"><div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:600\">{e}</div><div class=\"t-meta muted\">{m}</div></div><div style=\"text-align:right\"><div class=\"mono t-body strong\">{v} <span class=\"t-meta muted\">DPS</span></div><div class=\"t-meta\">{pct(p)} <span class=\"muted\">percentile</span></div></div></div>" for e, m, v, p in rows)
    return label("Bests") + card(r)


def char_page(plat: str, part: int = 1, state: str = "default") -> str:
    right = icon_btn(plat, "star-o", "Watch") + icon_btn(plat, "share", "Share")
    top = content_top(plat)
    left = back_btn(plat, "Rankings") + (f"<span class=\"cond\" style=\"font-size:22px\">Voltaris</span>" if plat == "and" else "")
    bar = top_bar(plat, f"<span style=\"display:flex;align-items:center;gap:4px\">{left}</span>", right)
    if state == "empty":
        inner = (char_head() + card("<div class=\"t-body strong\" style=\"font-size:17px\">No logs yet.</div><div class=\"t-meta muted\" style=\"margin-top:4px\">Voltaris has no rated fight in a public report, so there is no score to show.</div>", margin="16px 16px 0", pad="16px")
                 + label("Gear") + card("<div class=\"t-body\">Voltaris has not sent gear to Forever Sixty.</div>", pad="16px"))
    elif state == "notfound":
        inner = card("<div class=\"t-body strong\" style=\"font-size:17px\">We have no character called Voltxris.</div><div class=\"t-meta muted\" style=\"margin-top:4px\">Check the spelling, the region and the ruleset.</div><div style=\"display:flex;gap:10px;margin-top:12px\"><span class=\"btn btn-secondary\">Search again</span></div>", margin="16px 16px 0", pad="16px")
    elif part == 1:
        inner = char_head() + score_card() + bests_card() + label("Recent fights") + card(
            "".join(f"<div class=\"lrow\" style=\"min-height:56px\"><span class=\"dot\" style=\"background:var(--kill)\"></span><div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:600\">{e}</div><div class=\"t-meta muted\">Kill · {d}</div></div><span class=\"mono t-body strong\">{v} <span class=\"t-meta muted\">DPS</span></span></div>" for e, d, v in [("Onyxia", "3 d ago", "541"), ("Barrow Deeps", "9 d ago", "498")]))
    else:
        inner = (char_head() + label("Gear") + card(paperdoll(), pad="14px 8px 10px")
                 + card(f"<div class=\"lrow\" style=\"min-height:52px\"><span class=\"t t-body strong\" style=\"font-weight:600\">Gear gap</span><span class=\"mono t-body strong\">38 points</span></div>", margin="8px 16px 0")
                 + label("Builds seen") + card(f"<div class=\"lrow\" style=\"min-height:60px\">{tile('tree')}<div class=\"t\"><div class=\"t-body strong mono\" style=\"font-weight:700\">30 / 3 / 18</div><div class=\"t-meta muted\">Elemental · first seen 29 Oct</div></div><span class=\"t-meta gold\" style=\"font-weight:600\">Open in planner</span></div>"))
    return device(plat, bar + f"<div style=\"padding-top:{top}px\">{inner}{TAB_CLEAR}</div>", "Home")


PD_LEFT = ["head", "neck", "shoulder", "back", "chest", "wrist"]
PD_RIGHT = ["hands", "waist", "legs", "feet", "finger1", "finger2", "trinket1", "trinket2"]
LAB = {"head": "Head", "neck": "Neck", "shoulder": "Shoulder", "back": "Back", "chest": "Chest", "wrist": "Wrist", "hands": "Hands", "waist": "Waist", "legs": "Legs", "feet": "Feet",
       "finger1": "Ring", "finger2": "Ring", "trinket1": "Trinket", "trinket2": "Trinket", "main_hand": "Main hand", "off_hand": "Off hand", "ranged": "Ranged"}


def pd_cell(key: str, side: str, gap: bool) -> str:
    s = slots().get(key)
    if not s:
        return ""
    iid = s["item_id"]
    it = item(iid)
    dot = "<i style=\"position:absolute;right:-3px;top:-3px;width:10px;height:10px;border-radius:5px;background:var(--gold);border:2px solid #0d111a\"></i>" if gap else ""
    ic = f"<span style=\"position:relative;flex:none\">{item_icon(iid, 40)}{dot}</span>"
    txt = (f"<div style=\"min-width:0;text-align:{'left' if side == 'l' else 'right'}\"><div style=\"font-size:12px;line-height:14px;font-weight:600;color:{Q_TEXT[it['quality']]}\">{it['name']}</div>"
           f"<div class=\"t-micro muted\" style=\"letter-spacing:0;text-transform:none;font-weight:400\">{LAB[key]} · ilvl {it['item_level']}</div></div>")
    return f"<div style=\"display:flex;gap:8px;align-items:center;min-height:50px;flex-direction:{'row' if side == 'l' else 'row-reverse'}\">{ic}{txt}</div>"


def paperdoll(gaps: tuple = ("back", "waist")) -> str:
    l = "".join(pd_cell(k, "l", k in gaps) for k in PD_LEFT)
    r = "".join(pd_cell(k, "r", k in gaps) for k in PD_RIGHT)
    w = "".join(pd_cell(k, "l", False) for k in ("main_hand", "off_hand", "ranged") if k in slots())
    return (f"<div style=\"display:grid;grid-template-columns:1fr 1fr;gap:4px 8px\"><div>{l}</div><div>{r}</div></div>"
            f"<div style=\"display:grid;grid-template-columns:1fr 1fr 1fr;gap:4px;margin-top:8px;padding-top:8px;border-top:1px solid var(--line-soft)\">{w}</div>"
            "<div class=\"t-meta muted\" style=\"padding:6px 8px 0\">Gold dot: the BiS list has a better item for this slot.</div>")


# ---------------- changes to v1 root screens ----------------

def home_rows(plat: str) -> str:
    """Home scrolled to Next for {name}: the v1 rows plus Logs and Rankings, and the search button on the bar."""
    import screens_home as H
    top = content_top(plat)

    def r(ic: str, title: str, sub: str, trail: str = "") -> str:
        return (f"<div class=\"lrow\">{ic}<div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:600\">{title}</div><div class=\"t-meta muted\">{sub}</div></div>{trail}<span class=\"chev\">{glyph('chev-right', 18)}</span></div>")
    rows = (r(sp("spell_nature_lightning", 36), "Talents", "Your build is the Elemental 60 guide build", tag("Matches", "green"))
            + r(tile("flask"), "Sims", "Your Elemental sim finished: <span class=\"mono strong\">412 DPS</span>")
            + r(tile("calendar"), "Next raid", "Onyxia's Lair · Thursday 8 pm", tag("In · DPS", "green"))
            + r(tile("doc"), "Logs", "Your last raid: Onyxia's Lair, 1 kill", "<span class=\"tag tag-gold\">New</span>")
            + r(tile("trophy"), "Rankings", "14th of 212 Elemental Shaman characters on Onyxia"))
    body = (f"<div style=\"padding-top:{top}px\"><div style=\"transform:translateY(-250px)\">{H.hero()}{H.upgrades_card()}"
            f"<div class=\"t-micro micro pad\" style=\"margin-top:22px\">Next for Lightningbolltz</div><div class=\"card quiet\" style=\"margin:8px 16px 0\">{rows}</div></div></div>")
    cover = f"<div style=\"position:absolute;top:0;left:0;right:0;height:{DEVICES[plat]['top']}px;background:var(--bg);z-index:11\"></div>"
    return device(plat, cover + root_bar(plat) + body, "Home")


def raids_rows(plat: str) -> str:
    """The Raids tab with the two new rows under the guild head and the search button on the bar."""
    import screens_raids as R
    top = content_top(plat)
    rows = (f"<div class=\"lrow\">{tile('users')}<div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:600\">Roster</div><div class=\"t-meta muted\">38 on the roster · attendance</div></div><span class=\"chev\">{glyph('chev-right', 18)}</span></div>"
            f"<div class=\"lrow\">{tile('doc')}<div class=\"t\"><div class=\"t-body strong\" style=\"font-weight:600\">Raid history</div><div class=\"t-meta muted\">Logged raid nights</div></div><span class=\"chev\">{glyph('chev-right', 18)}</span></div>")
    body = f"<div style=\"padding-top:{top}px\">{R.guild_head()}<div class=\"card quiet\" style=\"margin:12px 16px 0\">{rows}</div>{R.next_raid_card()}</div>"
    return device(plat, root_bar(plat) + body, "Raids")
