#!/usr/bin/env python3
"""Mocks for design/specs/2026-10-31-paperdoll.md.

Real items from the active build (web/src/data/active-build.json): names, icons, qualities and
item levels come from data/builds/<build>/items/warrior.json and bis/. Gains on the character
board are illustrative. Run from the repo root:  python3 design/mocks/2026-10-31-paperdoll/gen.py
It writes html/*.html and renders every board with Playwright (web/node_modules).
"""
import json
import pathlib
import subprocess
import html as h

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

LEFT = ["head", "neck", "shoulder", "back", "chest", "shirt", "tabard", "wrist"]
RIGHT = ["hands", "waist", "legs", "feet", "finger1", "finger2", "trinket1", "trinket2"]
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


def band(spec, level, preset):
    d = json.loads((DATA / f"bis/{spec}.json").read_text())
    b = next(x for x in d["bands"] if x["band"] == level and x["preset"] == preset and x["faction"] == "alliance")
    return {s["slot"]: s["item_id"] for s in b["slots"] if s.get("item_id")}


def uri(p):
    return pathlib.Path(p).as_uri()


CSS = (HERE / "paperdoll.css").read_text().replace("{FONTS}", uri(FONTS))


def socket(slot, item=None, state="item", ghost=None):
    if state == "item":
        q = QUALITY[item["quality"]]
        return f'<span class="sock" style="--q:var(--q-{q})"><img src="{uri(ICONS / (item["icon"] + ".webp"))}" alt=""></span>'
    if state == "unknown":
        return f'<span class="sock unknown"><img src="{uri(ICONS / "inv_misc_questionmark.webp")}" alt=""></span>'
    if state == "twohand":
        return f'<span class="sock empty"><img class="ghost" src="{uri(ICONS / (ghost["icon"] + ".webp"))}" alt=""></span>'
    if state == "loading":
        return '<span class="sock loading"></span>'
    if state == "error":
        return '<span class="sock bare"></span>'
    return f'<span class="sock empty"><img class="sil" src="{uri(HERE / "slots" / (SILHOUETTE[slot] + ".jpg"))}" alt=""></span>'


def cell(slot, gear, side, extra="", detail=None, force=None):
    """One slot. gear[slot] is an item id, None (empty) or ('unknown', id)."""
    v = gear.get(slot)
    cls = f"slot {side} {extra}".strip()
    det = f'<span class="det">{detail}</span>' if detail else ""
    if force == "loading":
        return f'<div class="{cls} is-static">{socket(slot, state="loading")}<span class="txt"><span class="bar"></span></span></div>'
    if force == "error" and isinstance(v, int):
        it = ITEMS[v]
        return f'<button class="{cls}">{socket(slot, state="error")}<span class="txt"><span class="name plain">{h.escape(it["name"])}</span>{det}</span></button>'
    if slot == "off_hand" and v is None and isinstance(gear.get("main_hand"), int) and ITEMS[gear["main_hand"]]["two_hand"]:
        return f'<div class="{cls} is-static">{socket(slot, state="twohand", ghost=ITEMS[gear["main_hand"]])}<span class="txt"><span class="name muted">Two-hand in main hand</span></span></div>'
    if isinstance(v, tuple):
        return (f'<div class="{cls} is-static" aria-label="{LABELS[slot]}: Item {v[1]}, not in our item data yet">'
                f'{socket(slot, state="unknown")}<span class="txt"><span class="name id">Item {v[1]}</span>{det}</span></div>')
    if v is None:
        return f'<div class="{cls} is-static">{socket(slot, state="empty")}<span class="txt"><span class="name muted">{LABELS[slot]}</span>{det}</span></div>'
    it = ITEMS[v]
    q = QUALITY[it["quality"]]
    return (f'<button class="{cls}" aria-label="{LABELS[slot]}: {h.escape(it["name"])}">{socket(slot, it)}'
            f'<span class="txt"><span class="name q-{q}">{h.escape(it["name"])}</span>{det}</span></button>')


def avg_ilvl(gear):
    lv = [ITEMS[v]["item_level"] for v in gear.values() if isinstance(v, int)]
    return f"{sum(lv) / len(lv):.1f}", len(lv), sum(1 for v in gear.values() if v is not None)


def centre(who, gear, force=None):
    if force == "loading":
        figure = '<span class="figure skel">&nbsp;</span>'
    else:
        a, known, worn = avg_ilvl(gear)
        figure = f'<span class="figure">{a}</span>'
        if known < worn:
            figure += f'<span class="partial">{known} of {worn} items</span>'
    return f'''<div class="centre">
  <img class="crest" src="{uri(CRESTS / (who["class"] + ".webp"))}" alt="" style="--c:var(--class-{who["class"]})">
  <div class="who">
    <span class="cname" style="color:var(--class-{who["class"]})">{who["name"]}</span>
    <span class="desc"><img class="fac" src="{uri(FACTION / (who["faction"] + "-logo-64.webp"))}" alt="">{who["race"]} {who["spec"]} {who["klass"]}</span>
    <span class="lvl">Level {who["level"]}</span>
    <span class="ilvl"><span class="micro">Item level</span>{figure}</span>
    <span class="pill">{who["source"]}</span>
  </div>
</div>'''


def paperdoll(who, gear, details=None, force=None, overrides=None):
    details = details or {}
    overrides = overrides or {}
    parts = [centre(who, gear, force)]
    left = [s for s in LEFT if s in gear]
    for i, s in enumerate(left):
        parts.append(cell(s, gear, "l", detail=details.get(s), force=force).replace('class="slot', f'style="--r:{i + 1}" class="slot', 1))
    for i, s in enumerate(RIGHT):
        parts.append(cell(s, gear, "r", extra=overrides.get(s, ""), detail=details.get(s), force=force).replace('class="slot', f'style="--r:{i + 1}" class="slot', 1))
    parts.append(cell("main_hand", gear, "l mh", detail=details.get("main_hand"), force=force))
    parts.append(cell("off_hand", gear, "r oh", detail=details.get("off_hand"), force=force))
    parts.append(cell("ranged", gear, "l rg", detail=details.get("ranged"), force=force))
    unknown = sum(1 for v in gear.values() if isinstance(v, tuple))
    foot = ""
    if unknown:
        foot = (f'<p class="pd-note">{unknown} {"item is" if unknown == 1 else "items are"} not in our item data yet. '
                'They are new in the beta, and they get names and icons when the client lists them.</p>')
    if force == "error":
        foot = '<p class="pd-note">Item icons and colours did not load. Reload the page to try again.</p>'
    return f'<div class="pd"><div class="pd-frame"><div class="pd-grid">{"".join(parts)}</div>{foot}</div></div>'


def card(title, body, aside=""):
    return f'''<section class="card"><header class="chead"><span class="tile">{TILE}</span>
<h2 class="ctitle">{title}</h2><span class="aside">{aside}</span></header><div class="cbody">{body}</div></section>'''


TILE = '<svg viewBox="0 0 24 24" width="14" height="14" fill="none" stroke="currentColor" stroke-width="2"><path d="M12 3l7 3v5c0 5-3.5 8-7 10-3.5-2-7-5-7-10V6z"/></svg>'


def page(body, width, label):
    return f'''<!doctype html><html><head><meta charset="utf-8"><style>{CSS}</style></head>
<body style="width:{width}px"><div class="board-label">{label}</div><main class="col">{body}</main></body></html>'''


def write(name, html):
    OUT.mkdir(exist_ok=True)
    (OUT / f"{name}.html").write_text(html)


FURY = band("warrior-fury", 60, "raid")
FURY_WHO = {"name": "Brakka", "class": "warrior", "klass": "Warrior", "race": "Human", "spec": "Fury",
            "level": 60, "faction": "alliance", "source": "Addon export, 2 h ago"}

ARMS = band("warrior-arms", 60, "bare") | band("warrior-arms", 50, "bare")
ARMS_RAID = band("warrior-arms", 60, "raid")
ARMS["neck"] = ARMS_RAID["neck"]  # one slot already at its best
# The beta case: three worn items the build has never listed, a two-hander, an empty ranged.
for slot, iid in (("chest", 288892), ("waist", 288887), ("trinket2", 288893)):
    assert str(iid) not in NAMES, iid
    ARMS[slot] = ("unknown", iid)
ARMS["off_hand"] = None
ARMS["ranged"] = None
if not ITEMS[ARMS["main_hand"]]["two_hand"]:
    ARMS["main_hand"] = ARMS_RAID["main_hand"]
assert ITEMS[ARMS["main_hand"]]["two_hand"]
ARMS_WHO = {"name": "Brakka", "class": "warrior", "klass": "Warrior", "race": "Human", "spec": "Arms",
            "level": 60, "faction": "alliance", "source": "Addon export, 12 min ago"}


def gain(pts, slot):
    it = ITEMS[ARMS_RAID[slot]]
    return f'<b class="gain">+{pts} pts</b> · <span class="q-{QUALITY[it["quality"]]}">{h.escape(it["name"])}</span>'


# Gains go on the slots whose worn (pre-raid) item differs from the raid pick; the figures are
# illustrative, the items are real.
GAIN_PTS = ["18.4", "14.2", "9.6", "6.1", "4.8"]
GAIN_SLOTS = [s for s in LEFT + RIGHT + ["main_hand"]
              if isinstance(ARMS.get(s), int) and ARMS_RAID.get(s) not in (None, ARMS[s])][:len(GAIN_PTS)]
GAINS = {s: gain(p, s) for s, p in zip(GAIN_SLOTS, GAIN_PTS)}
GAINS["ranged"] = '<b class="gain">+4.0 pts</b> · <span class="q-epic">' + h.escape(ITEMS[22656]["name"]) + "</span>"
BEST = next(s for s in LEFT + RIGHT if isinstance(ARMS.get(s), int) and ARMS_RAID.get(s) == ARMS[s])
GAINS[BEST] = '<span class="best">Your best for level 60</span>'
TOTAL = sum(float(p) for p in GAIN_PTS[:len(GAIN_SLOTS)]) + 4.0

SIM_ASIDE = '<button class="btn-sec">Change source</button>'


def boards():
    fury = paperdoll(FURY_WHO, FURY)
    arms = paperdoll(ARMS_WHO, ARMS, details=GAINS)
    char_aside = f'<span class="meta">+{TOTAL:.1f} points across {len(GAIN_SLOTS) + 1} slots</span>'
    for w in (1280, 1024, 2000):
        write(f"raid-{w}", page(card("Character", fury, SIM_ASIDE), w, f"Sim character card · full raid set · {w} px"))
    write("beta-1280", page(card("Gear", arms, char_aside), 1280,
                            "Character page gear · beta-unknown items, two-hander, empty ranged, gains · 1280 px"))
    write("raid-390", page(card("Character", fury, SIM_ASIDE), 390, "Sim · raid set · 390 px"))
    write("beta-390", page(card("Gear", arms, char_aside), 390, "Character · beta case · 390 px"))
    write("raid-360", page(card("Character", fury, SIM_ASIDE), 360, "Sim · raid set · 360 px"))
    empty = {s: None for s in LEFT if s not in ("shirt", "tabard")} | {s: None for s in RIGHT}
    empty |= {"main_hand": None, "off_hand": None, "ranged": None}
    empty["main_hand"] = FURY["main_hand"]
    empty["head"] = FURY["head"]
    empty["trinket1"] = FURY["trinket1"]
    write("empty-1280", page(card("Character", paperdoll(FURY_WHO | {"source": "Build from planner"}, empty), SIM_ASIDE),
                             1280, "Mostly empty set (a fresh planner build) · 1280 px"))
    write("loading-1280", page(card("Character", paperdoll(FURY_WHO, FURY, force="loading"), SIM_ASIDE), 1280,
                               "Loading (items file pending) · 1280 px"))
    write("loading-390", page(card("Character", paperdoll(FURY_WHO, FURY, force="loading"), SIM_ASIDE), 390,
                              "Loading · 390 px"))
    write("states", page(states(), 1280, "Slot states, every one (desktop row 52 px; phone row 48 px)"))


def tooltip(iid):
    it = ITEMS[iid]
    q = QUALITY[it["quality"]]
    lines = [f'<div class="tt-name q-{q}">{h.escape(it["name"])}</div>', f'<div class="tt-ilvl">Item Level {it["item_level"]}</div>',
             '<div>Binds when picked up</div>', '<div class="tt-row"><span>Head</span><span>Plate</span></div>',
             f'<div>{it["armor"]} Armor</div>']
    names = {"strength": "Strength", "stamina": "Stamina", "agility": "Agility"}
    for k, v in it["stats"].items():
        if k in names:
            lines.append(f"<div>+{v} {names[k]}</div>")
    lines.append(f'<div>Requires Level {it["required_level"]}</div>')
    if "crit" in it["stats"]:
        lines.append(f'<div class="tt-green">Equip: Improves your chance to get a critical strike by {it["stats"]["crit"] // 14}%.</div>')
    if "hit" in it["stats"]:
        lines.append(f'<div class="tt-green">Equip: Improves your chance to hit by {it["stats"]["hit"] // 10}%.</div>')
    lines.append('<div class="tt-src">Blacksmithing</div>')
    return f'<div class="tt">{"".join(lines)}</div>'


def states():
    g = dict(FURY)
    rows = []

    def row(title, inner, note=""):
        rows.append(f'<div class="st"><div class="st-h">{title}</div><div class="st-c">{inner}</div><div class="st-n">{note}</div></div>')

    one = lambda s, gear=g, **k: f'<div class="pd one"><div class="pd-grid solo">{cell(s, gear, k.pop("side", "l"), **k)}</div></div>'
    row("Default", one("head"), "Socket 44, 2 px quality border, name 14/18 600 in the quality text colour.")
    row("Hover", one("head", extra="is-hover"), "--row-hover ground on the 8 px-radius cell; tooltip opens (below).")
    row("Focus-visible", one("head", extra="is-focus"), "2 px --color-gold outline, offset 2 px; tooltip opens on focus.")
    row("Active / tapped", one("head", extra="is-active"), "rgba(229,185,85,.10) ground; a tap toggles the tooltip.")
    row("Right column", one("trinket1", side="r"), "Mirrored: socket on the outer edge, text right-aligned toward the crest.")
    row("With a gain", one(GAIN_SLOTS[0], gear=ARMS, detail=GAINS[GAIN_SLOTS[0]]), "Character page: the slot's line 2. Gain gold 600, pick in its quality colour, its own tooltip.")
    row("Empty", one("ranged", gear=ARMS), "1 px dashed --color-line socket, the client's slot silhouette at 45%, label in --muted. Not focusable.")
    row("Not in our item data yet", one("chest", gear=ARMS), "The game's question-mark icon at 55%, 1 px --color-line border, \"Item 288892\" 13 px tabular --muted. One note under the panel says why.")
    row("Two-hand in main hand", one("off_hand", gear=ARMS, side="r"), "Dashed socket holding the two-hander's icon greyscale at 30%, as the client does.")
    row("Loading", one("head", force="loading"), "Shimmer socket and a 12 px bar at 60%; the row keeps its 52 px, nothing moves.")
    row("Items file failed", one("head", force="error"), "Plain socket, name in --text (no quality claimed), still a tooltip trigger.")
    return f'<div class="states">{"".join(rows)}</div><div class="tt-demo">{one("head", extra="is-hover")}{tooltip(FURY["head"])}</div>'


def render():
    shots = sorted(OUT.glob("*.html"))
    script = HERE / "shoot.mjs"
    subprocess.run(["node", str(script), *map(str, shots)], check=True, cwd=ROOT / "web")


if __name__ == "__main__":
    print("build", BUILD)
    print("gains on", GAIN_SLOTS, "best", BEST)
    boards()
    render()
