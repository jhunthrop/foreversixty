#!/usr/bin/env python3
"""Mocks for design/specs/2026-10-31-paperdoll.md (fix round 1).

Every item, icon, quality, item level, required level and source line comes from the active build
(web/src/data/active-build.json): data/builds/<build>/items/warrior.json and loot.json. Sets are
launch-legal: each item has an ungated dungeon, crafted or world-drop source, none of the 1.10
summoned-boss items (22300-22450), no Balzaphon drop and no late recipe (the 2026-10-30 BiS phase
audit). Gains are a real comparison: the item's stats times the spec's scale factors from
bis/<spec>.json (strength = 1.00; a percent stat converted with the weight's own rating factor).
Run from the repo root:  python3 design/mocks/2026-10-31-paperdoll/gen.py
"""
import html as h
import json
import pathlib
import subprocess

ROOT = pathlib.Path(__file__).resolve().parents[3]
HERE = pathlib.Path(__file__).resolve().parent
BUILD = json.loads((ROOT / "web/src/data/active-build.json").read_text())["build"]
DATA = ROOT / "data/builds" / BUILD
ICONS = DATA / "icons"
CRESTS = ROOT / "web/public/icons/hd/crests"
FACTION = ROOT / "web/public/icons/hd/faction"
FONTS = ROOT / "apps/shared/fonts"
OUT = HERE / "html"

ITEMS = {i["id"]: i for i in json.loads((DATA / "items/warrior.json").read_text())["items"]}
NAMES = json.loads((DATA / "itemnames.json").read_text())["names"]
LOOT = json.loads((DATA / "loot.json").read_text())

LEFT = ["head", "neck", "shoulder", "back", "chest", "shirt", "tabard", "wrist"]
RIGHT = ["hands", "waist", "legs", "feet", "finger1", "finger2", "trinket1", "trinket2"]
ARMOUR = [s for s in LEFT + RIGHT if s not in ("shirt", "tabard")]
LABELS = {"head": "Head", "neck": "Neck", "shoulder": "Shoulder", "back": "Back", "chest": "Chest",
          "shirt": "Shirt", "tabard": "Tabard", "wrist": "Wrist", "hands": "Hands", "waist": "Waist",
          "legs": "Legs", "feet": "Feet", "finger1": "Finger 1", "finger2": "Finger 2",
          "trinket1": "Trinket 1", "trinket2": "Trinket 2", "main_hand": "Main hand",
          "off_hand": "Off hand", "ranged": "Ranged"}
SILHOUETTE = {"head": "head", "neck": "neck", "shoulder": "shoulder", "back": "chest", "chest": "chest",
              "shirt": "shirt", "tabard": "tabard", "wrist": "wrists", "hands": "hands", "waist": "waist",
              "legs": "legs", "feet": "feet", "finger1": "finger", "finger2": "finger",
              "trinket1": "trinket", "trinket2": "trinket", "main_hand": "mainhand",
              "off_hand": "offhand", "ranged": "ranged"}
QUALITY = {0: "poor", 1: "common", 2: "uncommon", 3: "rare", 4: "epic", 5: "legendary"}
ITEM_SLOT = {"finger1": "finger", "finger2": "finger", "trinket1": "trinket", "trinket2": "trinket"}
LATE_RECIPES = {22385, 21278}           # audit `late-recipe`
SUMMONED_1_10 = range(22300, 22451)    # audit `t05`: Valthalak, Sothos and Jarien, Mor Grayhoof and kin


def uri(p):
    return pathlib.Path(p).as_uri()


# ---- sources: ungated dungeon, crafted and world-drop sources only ------------------------------
SOURCES = {}
for src in LOOT["sources"]:
    if src.get("opens") or src["kind"] not in ("dungeon", "crafted", "world_drop"):
        continue
    for boss in src.get("bosses") or []:
        if boss.get("name") == "Balzaphon":
            continue
        for iid in boss.get("items", []):
            chance = (boss.get("item_chances") or {}).get(str(iid))
            SOURCES.setdefault(iid, []).append((0, f'{src["name"]} — {boss["name"].strip()}', chance))
    for iid in src.get("items") or []:
        if src["kind"] == "crafted":
            SOURCES.setdefault(iid, []).append((1, f'Crafted: {src["name"]}', None))
        else:
            SOURCES.setdefault(iid, []).append(
                (1, f'World drop (level {src.get("level_min")} to {src.get("level_max")})', None))


def legal(iid):
    return (iid in SOURCES and iid not in LATE_RECIPES and iid not in SUMMONED_1_10
            and ITEMS[iid]["required_level"] <= 60 and ITEMS[iid]["quality"] >= 2)


def weights(spec):
    d = json.loads((DATA / f"bis/{spec}.json").read_text())
    b = next(x for x in d["bands"] if x["band"] == 60 and x["preset"] == "raid" and x["faction"] == "alliance")
    return {w["stat"]: (w["scale_factor"], w.get("rating_factor")) for w in b["weights"]}


def score(iid, w):
    it = ITEMS[iid]
    total = 0.0
    for stat, value in it["stats"].items():
        if stat not in w:
            continue
        factor, rating = w[stat]
        if stat in it.get("percent_stats", {}) and rating:
            value = it["percent_stats"][stat] * rating
        total += value * factor
    return total


def best(slot, w, taken=(), max_ilvl=999, two_hand=None):
    key = ITEM_SLOT.get(slot, slot)
    pool = [i for i, it in ITEMS.items() if it["slot"] == key and legal(i) and i not in taken
            and it["item_level"] <= max_ilvl and (two_hand is None or it["two_hand"] == two_hand)]
    if slot in ("main_hand", "off_hand", "ranged"):
        return max(pool, key=lambda i: (ITEMS[i]["dps"], i))
    return max(pool, key=lambda i: (score(i, w), i))


def full_set(w, max_ilvl=999, two_hander=False):
    gear = {}
    for s in ARMOUR:
        pair = {"finger2": "finger1", "trinket2": "trinket1"}.get(s)
        gear[s] = best(s, w, taken=(gear[pair],) if pair else (), max_ilvl=max_ilvl)
    if two_hander:
        gear["main_hand"] = best("main_hand", w, max_ilvl=max_ilvl, two_hand=True)
        gear["off_hand"] = None
    else:
        gear["main_hand"] = best("main_hand", w, max_ilvl=max_ilvl, two_hand=False)
        gear["off_hand"] = best("main_hand", w, taken=(gear["main_hand"],), max_ilvl=max_ilvl, two_hand=False)
    gear["ranged"] = best("ranged", w, max_ilvl=max_ilvl)
    return gear


# ---- the sets ------------------------------------------------------------------------------------
FURY_W, ARMS_W = weights("warrior-fury"), weights("warrior-arms")
FURY = full_set(FURY_W)
ARMS_PICK = full_set(ARMS_W, two_hander=True)
ARMS = full_set(ARMS_W, max_ilvl=58, two_hander=True)
ARMS["neck"] = ARMS_PICK["neck"]                              # one slot already at its best
ARMS["main_hand"] = 12790                                     # Arcanite Champion: crafted, two-handed
assert ITEMS[12790]["two_hand"] and legal(12790)
UNKNOWN = {"chest": 288892, "waist": 288887, "trinket2": 288893}
for slot, iid in UNKNOWN.items():
    assert str(iid) not in NAMES and iid not in ITEMS, iid
    ARMS[slot] = ("unknown", iid)
ARMS["ranged"] = None

# The longest launch-legal names a warrior can wear, by slot, for the worst-case board.
LONGEST = {}
for i, it in ITEMS.items():
    if legal(i) and it["quality"] >= 3:
        k = it["slot"]
        if k not in LONGEST or len(it["name"]) > len(ITEMS[LONGEST[k]]["name"]):
            LONGEST[k] = i
WORST = {s: LONGEST[ITEM_SLOT.get(s, s)] for s in ARMOUR}
WORST["finger2"] = max((i for i, it in ITEMS.items() if it["slot"] == "finger" and legal(i) and it["quality"] >= 3
                        and i != WORST["finger1"]), key=lambda i: len(ITEMS[i]["name"]))
WORST["trinket2"] = max((i for i, it in ITEMS.items() if it["slot"] == "trinket" and legal(i) and it["quality"] >= 3
                         and i != WORST["trinket1"]), key=lambda i: len(ITEMS[i]["name"]))
WORST["main_hand"] = max((i for i, it in ITEMS.items() if it["slot"] == "main_hand" and legal(i) and it["quality"] >= 3
                          and not it["two_hand"]), key=lambda i: len(ITEMS[i]["name"]))
WORST["off_hand"] = max((i for i, it in ITEMS.items() if it["slot"] in ("off_hand", "main_hand") and legal(i)
                         and it["quality"] >= 3 and not it["two_hand"] and i != WORST["main_hand"]),
                        key=lambda i: len(ITEMS[i]["name"]))
WORST["ranged"] = LONGEST["ranged"]


def gains(worn, pick, w):
    """slot -> ('gain', pts, pick) | ('best',) | ('unscored', pick), from a real stat comparison."""
    out = {}
    for s in ARMOUR:
        p, v = pick.get(s), worn.get(s)
        if p is None:
            continue
        if isinstance(v, tuple):
            out[s] = ("unscored", p)
        elif v == p:
            out[s] = ("best",)
        elif v is None or score(p, w) - score(v, w) > 0.05:
            out[s] = ("gain", score(p, w) - (score(v, w) if v else 0), p)
    return out


ARMS_GAINS = gains(ARMS, ARMS_PICK, ARMS_W)


# ---- markup --------------------------------------------------------------------------------------
CSS = (HERE / "paperdoll.css").read_text().replace("{FONTS}", uri(FONTS))
USER_GLYPH = ('<svg viewBox="0 0 24 24" width="14" height="14" fill="none" stroke="currentColor" stroke-width="2" '
              'stroke-linecap="round"><circle cx="12" cy="8" r="4"/><path d="M4 21c1.5-4 4.5-6 8-6s6.5 2 8 6"/></svg>')


def icon(iid):
    return uri(ICONS / (ITEMS[iid]["icon"] + ".webp"))


def socket(slot, state, iid=None):
    if state == "item":
        return f'<span class="sock" style="--q:var(--q-{QUALITY[ITEMS[iid]["quality"]]})"><img src="{icon(iid)}" alt=""></span>'
    if state == "unknown":
        return f'<span class="sock unknown"><img src="{uri(ICONS / "inv_misc_questionmark.webp")}" alt=""></span>'
    if state == "loading":
        return '<span class="sock loading"></span>'
    if state == "error":
        return '<span class="sock bare"></span>'
    return f'<span class="sock empty"><img class="sil" src="{uri(HERE / "slots" / (SILHOUETTE[slot] + ".jpg"))}" alt=""></span>'


def fmt(pts):
    return f"+{pts:.1f}"


def detail(g):
    if g is None:
        return ""
    if g[0] == "best":
        return '<span class="det"><span class="best">Best in slot</span></span>'
    pick = g[-1]
    q = QUALITY[ITEMS[pick]["quality"]]
    tail = f'<b class="gain">{fmt(g[1])}</b>' if g[0] == "gain" else '<span class="unscored">Not scored</span>'
    return (f'<span class="det"><img class="pick-ico" style="--q:var(--q-{q})" src="{icon(pick)}" alt="">'
            f'<span class="pick q-{q}">{h.escape(ITEMS[pick]["name"])}</span>{tail}</span>')


def cell(slot, gear, side, rows, g=None, force=None, extra=""):
    v = gear.get(slot)
    pos = f'style="--r:{rows[0]};--rp:{rows[1]}"'
    cls = f"slot {side} {extra}".strip()
    det = detail(g)
    if force == "loading":
        return f'<div {pos} class="{cls} is-static">{socket(slot, "loading")}<span class="txt"><span class="bar"></span></span></div>'
    if force == "error" and isinstance(v, int):
        return (f'<button {pos} class="{cls}">{socket(slot, "error")}<span class="txt"><span class="name plain">'
                f'{h.escape(ITEMS[v]["name"])}</span></span></button>')
    if slot == "off_hand" and v is None and isinstance(gear.get("main_hand"), int) and ITEMS[gear["main_hand"]]["two_hand"]:
        return (f'<div {pos} class="{cls} is-static">{socket(slot, "empty")}<span class="txt">'
                f'<span class="name muted">Two-hand in main hand</span></span></div>')
    if isinstance(v, tuple):
        return (f'<div {pos} class="{cls} is-static" aria-label="{LABELS[slot]}: Item {v[1]}, not in our item data yet">'
                f'{socket(slot, "unknown")}<span class="txt"><span class="name id">{LABELS[slot]} · Item {v[1]}</span>{det}</span></div>')
    if v is None:
        return (f'<div {pos} class="{cls} is-static">{socket(slot, "empty")}<span class="txt">'
                f'<span class="name muted">{LABELS[slot]}</span>{det}</span></div>')
    q = QUALITY[ITEMS[v]["quality"]]
    return (f'<button {pos} class="{cls}" data-item="{v}" aria-label="{LABELS[slot]}: {h.escape(ITEMS[v]["name"])}">'
            f'{socket(slot, "item", v)}<span class="txt"><span class="name q-{q}">{h.escape(ITEMS[v]["name"])}</span>{det}</span></button>')


def upgrade_row(slot, g, rp):
    """Phone only: the pick on its own full-width, tappable row under its slot's pair."""
    if g is None or g[0] == "best":
        return ""
    pick = g[-1]
    q = QUALITY[ITEMS[pick]["quality"]]
    tail = f'<b class="gain">{fmt(g[1])}</b>' if g[0] == "gain" else '<span class="unscored">Not scored</span>'
    return (f'<button class="up" style="--rp:{rp}" data-item="{pick}" aria-label="{LABELS[slot]} upgrade: '
            f'{h.escape(ITEMS[pick]["name"])}"><span class="up-slot">{LABELS[slot]}</span>'
            f'<span class="up-sock" style="--q:var(--q-{q})"><img src="{icon(pick)}" alt=""></span>'
            f'<span class="up-name q-{q}">{h.escape(ITEMS[pick]["name"])}</span>{tail}</button>')


def item_level(gear):
    """Shown only when every equipped slot except shirt and tabard is known; else None."""
    counted = [v for s, v in gear.items() if s not in ("shirt", "tabard") and v is not None]
    if not counted or any(isinstance(v, tuple) for v in counted):
        return None
    return f"{sum(ITEMS[v]['item_level'] for v in counted) / len(counted):.1f}"


def centre(who, gear, force=None):
    ilvl = item_level(gear)
    if force == "loading":
        lvl = '<span class="ilvl"><span class="micro">Item level</span><span class="figure skel">&nbsp;</span></span>'
    elif ilvl is None:
        lvl = ""
    else:
        lvl = f'<span class="ilvl"><span class="micro">Item level</span><span class="figure">{ilvl}</span></span>'
    return f'''<div class="centre">
  <img class="crest" src="{uri(CRESTS / (who["class"] + ".webp"))}" alt="" style="--c:var(--class-{who["class"]})">
  <div class="who">
    <span class="cname" style="color:var(--class-{who["class"]})">{who["name"]}</span>
    <span class="desc"><img class="fac" src="{uri(FACTION / (who["faction"] + "-logo-64.webp"))}" alt="">{who["race"]} {who["spec"]} {who["klass"]}</span>
    <span class="lvl">Level {who["level"]}</span>{lvl}
    <span class="pill">{who["source"]}</span>
  </div>
</div>'''


def paperdoll(who, gear, g=None, force=None):
    g = g or {}
    left = [s for s in LEFT if s in gear]
    parts, rp = [centre(who, gear, force)], 2
    for i in range(len(RIGHT)):
        l = left[i] if i < len(left) else None
        r = RIGHT[i]
        if l:
            parts.append(cell(l, gear, "l", (i + 1, rp), g.get(l), force))
        parts.append(cell(r, gear, "r", (i + 1, rp), g.get(r), force))
        rp += 1
        for s in (l, r):
            if s and g.get(s) and g[s][0] != "best" and force is None:
                parts.append(upgrade_row(s, g[s], rp))
                rp += 1
    parts.append(cell("main_hand", gear, "l mh", (9, rp), g.get("main_hand"), force))
    parts.append(cell("off_hand", gear, "r oh", (9, rp), g.get("off_hand"), force))
    rp += 1
    parts.append(cell("ranged", gear, "l rg", (10, rp), g.get("ranged"), force))
    rp += 1
    if g.get("ranged") and force is None:
        parts.append(upgrade_row("ranged", g["ranged"], rp))
    unknown = sum(1 for v in gear.values() if isinstance(v, tuple))
    note = ""
    if force == "error":
        note = '<p class="pd-note">Item icons and colours did not load. Reload the page to try again.</p>'
    elif unknown == 1:
        note = ('<p class="pd-note">1 item is not in our item data yet. It is new in the beta, and it gets a name '
                'and icon when the client lists it. Item level shows once every item is known.</p>')
    elif unknown:
        note = (f'<p class="pd-note">{unknown} items are not in our item data yet. They are new in the beta, and '
                'they get names and icons when the client lists them. Item level shows once every item is known.</p>')
    dcls = " has-detail" if g and force is None else ""
    return f'<div class="pd"><div class="stage{dcls}"><div class="pd-grid">{"".join(parts)}</div></div>{note}</div>'


def card(title, body, aside=""):
    return (f'<section class="card"><header class="chead"><span class="tile">{USER_GLYPH}</span>'
            f'<h2 class="ctitle">{title}</h2><span class="aside">{aside}</span></header><div class="cbody">{body}</div></section>')


def page(body, width, label, extra=""):
    return (f'<!doctype html><html lang="en"><head><meta charset="utf-8"><style>{CSS}</style></head>'
            f'<body style="width:{width}px"><div class="board-label">{label}</div><main class="col">{body}</main>{extra}</body></html>')


def write(name, html):
    OUT.mkdir(exist_ok=True)
    (OUT / f"{name}.html").write_text(html)


# ---- the shared tooltip, as ItemTooltip.svelte renders lib/items/tooltip.ts's model ----------
STAT_ORDER = ["strength", "agility", "stamina", "intellect", "spirit", "crit", "hit", "spell_power", "healing",
              "attack_power", "ranged_attack_power", "defense", "dodge", "parry", "block", "block_value", "mp5"]
STAT_LABEL = {"strength": "Strength", "agility": "Agility", "stamina": "Stamina", "intellect": "Intellect",
              "spirit": "Spirit", "crit": "Crit", "hit": "Hit", "spell_power": "Spell power", "healing": "Healing",
              "attack_power": "Attack power", "ranged_attack_power": "Ranged attack power", "defense": "Defense",
              "dodge": "Dodge", "parry": "Parry", "block": "Block", "block_value": "Block value", "mp5": "Mp5"}
SLOT_WORD = {"finger": "Finger", "trinket": "Trinket", "main_hand": "Main hand", "off_hand": "Off hand"}


def tooltip(iid, tid):
    it = ITEMS[iid]
    q = QUALITY[it["quality"]]
    head = [f'<span class="tt-name q-{q}">{h.escape(it["name"])}</span>',
            f'<span class="tt-small">Item Level {it["item_level"]}</span>']
    if it["unique"]:
        head.append('<span class="tt-small">Unique</span>')
    if it.get("client_unconfirmed"):
        head.append('<span class="tt-small">Stats from the original 1.12 tables; not yet seen in Forever\'s client</span>')
    body = [f'<div class="tt-head"><img src="{icon(iid)}" alt=""><div class="tt-title">{"".join(head)}</div></div>']
    slot = SLOT_WORD.get(it["slot"], LABELS.get(it["slot"], it["slot"].title()))
    typ = ("Two-Handed Weapon" if it["two_hand"] else "Ranged Weapon" if it["slot"] == "ranged" else "One-Handed Weapon") \
        if it["damage_max"] > 0 else ""
    body.append(f'<div class="tt-row"><span>{slot}</span><span>{typ}</span></div>')
    if it["armor"]:
        body.append(f'<span class="tt-12">{it["armor"]} Armor</span>')
    if it["damage_max"] > 0:
        body.append(f'<div class="tt-col tt-12"><span>{round(it["damage_min"])} - {round(it["damage_max"])} Damage</span>'
                    f'<span>Speed {it["speed"]:.2f}</span><span class="tt-muted">({it["dps"]:.1f} damage per second)</span></div>')
    stats = [f'+{it["stats"][k]} {STAT_LABEL[k]}' for k in STAT_ORDER if it["stats"].get(k)]
    if stats:
        body.append('<div class="tt-col tt-12 tt-green">' + "".join(f"<span>{s}</span>" for s in stats) + "</div>")
    if it["effect_text"]:
        body.append(f'<p class="tt-12 tt-green">{h.escape(it["effect_text"])}</p>')
    body.append(f'<span class="tt-small">Requires Level {it["required_level"]}</span>')
    srcs = sorted(SOURCES.get(iid, []), key=lambda x: (x[0], -(x[2] or -1)))
    lines = [t + ("" if c is None else " (<1%)" if c < 1 else f" ({round(c)}%)") for _, t, c in srcs][:3]
    body.append('<div class="tt-src">' + "".join(f"<span>{h.escape(x)}</span>" for x in lines) + "</div>")
    return f'<div class="tt" role="tooltip" id="{tid}">{"".join(body)}</div>'


# The placement rule, as a script: below the cell by 8 px, flip above when it does not fit;
# align to the cell's outer edge (start for the left column, end for the right); on a phone the
# panel is the viewport width minus 32 px; always kept 8 px inside the viewport.
PLACE = """<script>
document.fonts.ready.then(() => {
for (const tt of document.querySelectorAll('.tt[data-anchor]')) {
  const a = document.querySelector(tt.dataset.anchor); a.classList.add('is-hover');
  const mode = document.body.dataset.scroll;
  if (mode === 'bottom') window.scrollTo(0, a.getBoundingClientRect().bottom + scrollY - (innerHeight - 120));
  if (mode === 'top') window.scrollTo(0, a.getBoundingClientRect().top + scrollY - 260);
  const vw = document.documentElement.clientWidth, vh = window.innerHeight;
  const r = a.getBoundingClientRect(); const w = tt.getBoundingClientRect().width; const th = tt.offsetHeight;
  const end = a.classList.contains('r') && !a.classList.contains('up');
  let left = end ? r.right - w : r.left; left = Math.max(8, Math.min(left, vw - 8 - w));
  const below = r.bottom + 8 + th <= vh - 8;
  const top = below ? r.bottom + 8 : r.top - 8 - th;
  tt.style.left = (left + scrollX) + 'px'; tt.style.top = (top + scrollY) + 'px';
  tt.dataset.side = below ? 'below' : 'above';
}
});
</script>"""


def tip_page(name, width, label, who, gear, anchor_slot, scroll=None, g=None, height=800, use_up=False, lead=0):
    # `lead` stands in for page content above the card, so the cell can sit near the viewport bottom.
    gearcard = (f'<div style="height:{lead}px"></div>' if lead else "") + card("Character", paperdoll(who, gear, g))
    v = gear[anchor_slot] if not use_up else g[anchor_slot][-1]
    sel = f'.up[data-item="{v}"]' if use_up else f'.slot[data-item="{v}"]'
    tt = tooltip(v, "tt").replace('class="tt"', f'class="tt" data-anchor=\'{sel}\'', 1)
    # A spacer keeps the page scrollable, so the capture can put the cell where the rule is tested.
    spacer = '<div style="height:900px"></div>'
    write(name, page(gearcard, width, label, extra=spacer + tt + PLACE).replace(
        "<body ", f'<body data-scroll="{scroll or ""}" '))


def boards():
    fury_who = {"name": "Brakka", "class": "warrior", "klass": "Warrior", "race": "Human", "spec": "Fury",
                "level": 60, "faction": "alliance", "source": "Addon export, 2 h ago"}
    arms_who = fury_who | {"spec": "Arms", "source": "Addon export, 12 min ago"}
    sim_aside = '<button class="btn-sec">Change source</button>'
    total = sum(x[1] for x in ARMS_GAINS.values() if x[0] == "gain")
    n = sum(1 for x in ARMS_GAINS.values() if x[0] == "gain")
    char_aside = f'<span class="meta-gold">+{total:.1f} points across {n} slots</span>'
    fury = paperdoll(fury_who, FURY)
    arms = paperdoll(arms_who, ARMS, ARMS_GAINS)
    for w in (1280, 1024, 2000, 390, 360):
        write(f"raid-{w}", page(card("Character", fury, sim_aside), w, f"Sim character card · launch-legal level-60 Fury set · {w} px"))
    for w in (1280, 390, 360):
        write(f"beta-{w}", page(card("Gear", arms, char_aside), w,
                                f"Character page gear · 3 beta items, Arcanite Champion (two-hand), empty ranged, real gains · {w} px"))
    worst = paperdoll(fury_who | {"source": "Build from planner"}, WORST)
    for w in (1024, 390, 360):
        write(f"longest-{w}", page(card("Character", worst, sim_aside), w, f"Worst case: the longest launch-legal names · {w} px"))
    empty = {s: None for s in ARMOUR} | {"main_hand": FURY["main_hand"], "off_hand": None, "ranged": None,
                                          "head": FURY["head"], "trinket1": FURY["trinket1"]}
    write("empty-1280", page(card("Character", paperdoll(fury_who | {"source": "Build from planner"}, empty), sim_aside),
                             1280, "A planner build with 3 of 17 slots filled · 1280 px"))
    for w in (1280, 390):
        write(f"loading-{w}", page(card("Character", paperdoll(fury_who, FURY, force="loading"), sim_aside), w, f"Loading · {w} px"))
    write("states-1280", page(states(), 1280, "Slot states (desktop cell; phone sizes in the spec)"))
    tip_page("tip-left-1280", 1280, "Tooltip, left column: hangs from the cell's left edge, 8 px below", fury_who, FURY, "shoulder")
    tip_page("tip-right-1280", 1280, "Tooltip, right column: hangs from the cell's right edge, 8 px below", fury_who, FURY, "legs")
    tip_page("tip-flip-1280", 1280, "Tooltip near the viewport bottom: flips above the cell", fury_who, FURY, "trinket1", scroll="bottom", lead=500)
    tip_page("tip-phone-390", 390, "Phone: the panel is the viewport minus 32 px, below the tapped cell", fury_who, FURY,
             "hands", height=844)
    tip_page("tip-phoneflip-390", 390, "Phone near the bottom: flips above", fury_who, FURY, "trinket2", scroll="bottom", height=844)
    tip_page("tip-phoneup-390", 390, "Phone: tapping an upgrade row opens the pick's tooltip", arms_who, ARMS, "legs",
             g=ARMS_GAINS, use_up=True, height=844, scroll="top")


def states():
    rows = []

    def row(title, inner, note):
        rows.append(f'<div class="st"><div class="st-h">{title}</div><div class="st-c">{inner}</div><div class="st-n">{note}</div></div>')

    def one(slot, gear=FURY, side="l", g=None, force=None, extra=""):
        return f'<div class="pd one"><div class="pd-grid solo">{cell(slot, gear, side, (1, 1), g, force, extra)}</div></div>'

    gain_slot = next(s for s, x in ARMS_GAINS.items() if x[0] == "gain")
    best_slot = next(s for s, x in ARMS_GAINS.items() if x[0] == "best")
    unk_slot = next(s for s, x in ARMS_GAINS.items() if x[0] == "unscored")
    row("Default", one("head"), "Socket 44, 2 px quality border; name 14/18 600 in the quality text colour, 2-line clamp with an ellipsis.")
    row("Hover", one("head", extra="is-hover"), "Row-hover ground on the 8 px-radius cell; icon brightness 1.12; the tooltip opens.")
    row("Focus-visible", one("head", extra="is-focus"), "2 px gold outline, offset 2 px; the tooltip opens on focus, Esc closes it.")
    row("Active / tapped", one("head", extra="is-active"), "Gold 10% ground; a tap toggles the tooltip.")
    row("Right column", one("legs", side="r"), "Mirrored: socket on the outer edge, text right-aligned toward the crest.")
    row("With a gain", one(gain_slot, ARMS, g=ARMS_GAINS[gain_slot]), "Character page, desktop: line 2 is the pick (16 px icon, quality colour, its own tooltip) and the gain, gold 13/600, never cut.")
    row("Already the best", one(best_slot, ARMS, g=ARMS_GAINS[best_slot]), "The pick is the worn item: \"Best in slot\" in muted. Never a gain.")
    row("Empty, with a gain", one("back", {"back": None}, g=("gain", score(ARMS_PICK["back"], ARMS_W), ARMS_PICK["back"])), "Empty socket and its slot word; the gain (the whole pick's score) still sits on its line.")
    row("Not in our item data yet", one(unk_slot, ARMS, g=ARMS_GAINS[unk_slot]), "Question mark at 40%, half saturation, 1 px line border; \"Slot · Item id\" 13 px tabular muted. A pick shows \"Not scored\".")
    row("Empty", one("ranged", {"ranged": None}), "1 px dashed line socket, the client's slot texture at 45%, the slot word in muted. Not focusable.")
    row("Two-hand in main hand", one("off_hand", ARMS, side="r"), "The client draws the plain off-hand slot texture (PaperDollFrame.lua, classic_era 1.15.9). We add the words.")
    row("Loading", one("head", force="loading"), "Shimmer socket and a 12 px bar at 60%; the row keeps its height.")
    row("Items file failed", one("head", force="error"), "Plain socket; the name in body colour, no quality claimed; still a tooltip trigger.")
    return f'<div class="states">{"".join(rows)}</div>'


def render():
    shots = sorted(OUT.glob("*.html"))
    subprocess.run(["node", str(HERE / "shoot.mjs"), *map(str, shots)], check=True, cwd=ROOT / "web")


if __name__ == "__main__":
    print("build", BUILD)
    print("fury", {s: ITEMS[v]["name"] if v else None for s, v in FURY.items()})
    print("arms", {s: (ITEMS[v]["name"] if isinstance(v, int) else v) for s, v in ARMS.items()})
    print("gains", {s: (x[0], round(x[1], 1) if x[0] == "gain" else "", ITEMS[x[-1]]["name"] if x[0] != "best" else "") for s, x in ARMS_GAINS.items()})
    print("worst", {s: ITEMS[v]["name"] for s, v in WORST.items()})
    for old in HERE.glob("*.png"):
        old.unlink()
    for old in OUT.glob("*.html") if OUT.exists() else []:
        old.unlink()
    boards()
    render()
