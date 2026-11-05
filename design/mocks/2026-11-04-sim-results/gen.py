#!/usr/bin/env python3
"""Mocks for design/specs/2026-11-04-sim-results.md: the saved sim result view (run kind).

Every figure, name and icon comes from the real run (design/mocks/data/simview/fury-result.json) and
the active build; see data.py for the resolution order. Regions marked ILLUSTRATIVE on a board draw
data the engine does not return yet (spec section 11): their SHAPE is the spec, their values are
derived from the run (buff bars from cast times and spell durations, the target bar from Deep Wounds
ticks) or, for the histogram and rage sources, modelled on the run's own totals.
Run from the repo root:  python3 design/mocks/2026-11-04-sim-results/gen.py
"""
import html as h
import json
import pathlib
import subprocess
import sys

HERE = pathlib.Path(__file__).resolve().parent
sys.path.insert(0, str(HERE))
import data as D          # noqa: E402
import timeline as T      # noqa: E402

OUT = HERE / "html"
CSS = (HERE / "simview.css").read_text().replace("{FONTS}", (D.ROOT / "apps/shared/fonts").as_uri())
CREST = (D.ROOT / "web/public/icons/hd/crests" / f"{D.CLASS}.webp").as_uri()
GUIDE_TALENTS = json.loads((D.DATA / "bis/warrior-fury.json").read_text())
QUALITY_TEXT = {0: "var(--q0)", 1: "var(--q1)", 2: "var(--q2)", 3: "var(--q3t)", 4: "var(--q4t)", 5: "var(--q5)"}
QUALITY = {0: "var(--q0)", 1: "var(--q1)", 2: "var(--q2)", 3: "var(--q3)", 4: "var(--q4)", 5: "var(--q5)"}
E = h.escape

GLYPH = {
    "bars": '<path d="M5 20V10M12 20V4M19 20v-7"/>',
    "clock": '<circle cx="12" cy="12" r="8"/><path d="M12 8v4l3 2"/>',
    "shield": '<path d="M12 3l7 3v6c0 4-3 7-7 9-4-2-7-5-7-9V6z"/>',
    "tree": '<path d="M12 21v-6M7 15h10M12 3l5 6H7z"/><path d="M12 9l5 6H7z"/>',
    "bolt": '<path d="M13 3L5 13h6l-1 8 8-10h-6z"/>',
    "list": '<path d="M9 6h11M9 12h11M9 18h11M4 6h1M4 12h1M4 18h1"/>',
    "gear": '<circle cx="12" cy="12" r="3"/><path d="M12 3v3M12 18v3M3 12h3M18 12h3M5.6 5.6l2.1 2.1M16.3 16.3l2.1 2.1M5.6 18.4l2.1-2.1M16.3 7.7l2.1-2.1"/>',
    "spark": '<path d="M12 3v4M12 17v4M3 12h4M17 12h4M6 6l3 3M15 15l3 3M6 18l3-3M15 9l3-3"/>',
    "alert": '<path d="M12 4l9 16H3z"/><path d="M12 10v4M12 17v.5"/>',
    "link": '<path d="M10 14a4 4 0 006 0l3-3a4 4 0 00-6-6l-1 1"/><path d="M14 10a4 4 0 00-6 0l-3 3a4 4 0 006 6l1-1"/>',
    "copy": '<rect x="8" y="8" width="12" height="12" rx="2"/><path d="M16 8V5a1 1 0 00-1-1H5a1 1 0 00-1 1v10a1 1 0 001 1h3"/>',
    "down": '<path d="M12 4v11M7 10l5 5 5-5M5 20h14"/>',
    "play": '<path d="M7 4l12 8-12 8z"/>',
}


def glyph(name):
    return f'<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">{GLYPH[name]}</svg>'


def icon(name, cls="", alt=""):
    if not name:
        return f'<span class="ic {cls}" style="background:#14171f"></span>'
    return f'<img class="ic {cls}" src="{D.icon_path(name).as_uri()}" alt="{E(alt)}">'


def num(v, d=0):
    return f"{v:,.{d}f}"


def pct(v, d=1):
    return f"{v * 100:.{d}f}%"


def chead(tile, label, title, meta="", anchor=""):
    return (f'<div class="chead" id="{anchor}"><div class="itile">{glyph(tile)}</div><div><div class="clabel">{E(label)}</div>'
            f'<h2 class="ctitle">{E(title)}</h2></div><div class="cmeta">{meta}</div></div>')


def crest(size_cls=""):
    return (f'<div class="crest {size_cls}" style="--ring:var(--warrior)"><div class="disc" style="background-image:url({CREST})"></div>'
            f'<img src="{CREST}" alt="Warrior"></div>')


DPS = D.RUN["dps"]
ENC = D.REQ["encounter"]
APM_CASTS = sum(n for k, n in D.CASTS.items() if not k.startswith("other:") and "/" not in k)
APM = APM_CASTS / (D.DUR_MS / 60000)
SPLIT = "/".join(str(sum(int(c) for c in tree)) for tree in D.REQ["character"]["talents"].split("-"))


def site():
    nav = "".join(f'<span class="navl{" on" if n == "TOOLS" else ""}">{n}</span>' for n in ["TOOLS", "LOGS", "GUILD", "GUIDES", "GET SET UP"])
    return (f'<div class="site"><div class="wrap"><div class="inner"><div class="wm"><span class="seal">LX</span>FOREVER SIXTY</div>{nav}</div></div></div>'
            f'<div class="wrap"><div class="inner simtabs"><span class="on">QUICK SIM</span><span>TOP GEAR</span><span>DROPTIMIZER</span><span>TALENTS</span><span>WEIGHTS</span></div></div>')


def band():
    fight = f'{ENC["duration_sec"] // 60}:{ENC["duration_sec"] % 60:02d} ± {round(ENC["variation"] * 100)}%'
    facts = [("Fight", fight), ("Targets", str(ENC["targets"])), ("Buffs", "Raid"), ("Actions per minute", f"{APM:.1f}"),
             ("Range", f'{num(DPS["min"])} to {num(DPS["max"])}'), ("Runs", num(D.RUN["iterations_run"])), ("Saved", "18 Sep 2026")]
    f = "".join(f'<div class="fact"><div class="k">{k}</div><div class="v numeric">{v}</div></div>' for k, v in facts)
    return f'''<header class="band"><div class="wrap"><div class="inner">
  <div class="band-top">
    <div class="who">{crest()}<div style="min-width:0">
      <div class="eyebrow">Saved sim</div>
      <h1 class="h1" style="color:var(--warrior)">Fury Warrior</h1>
      <div class="desc"><b>Orc Fury Warrior</b> · Level 60 · {SPLIT} talents</div>
      <div class="nextline">Next: find the upgrade that adds the most. <a class="alink" href="#">Run Top Gear</a></div>
    </div></div>
    <div class="answer">
      <div class="fig"><span class="n numeric">{DPS["mean"]:.1f}</span><span class="u">DPS</span></div>
      <div class="err numeric">± <b>{DPS["error"]:.1f}</b> · {DPS["error"] / DPS["mean"] * 100:.2f}% error</div>
      <div class="acts"><span class="btn p">{glyph("play")}Run this yourself</span><span class="btn s">{glyph("link")}Copy link</span></div>
    </div>
  </div>
  <div class="facts">{f}</div>
</div></div></header>'''


SECTIONS = [("damage", "Damage"), ("timeline", "Timeline"), ("gear", "Gear"), ("talents", "Talents"),
            ("resource", "Rage"), ("buffs", "Buffs"), ("rotation", "Rotation"), ("settings", "Settings")]


def snav(on="damage", mini=True):
    links = "".join(f'<a class="{"on" if k == on else ""}" href="#{k}">{v}</a>' for k, v in SECTIONS)
    m = f'<span class="mini numeric">{DPS["mean"]:.1f}<span>DPS</span></span>' if mini else ""
    return f'<nav class="snav"><div class="wrap"><div class="inner">{links}{m}</div></div></nav>'


def notes():
    reck = D.CONST["1719"]
    rows = [f'''<div class="note-row">{icon("ability_criticalstrike", "", "Recklessness")}<p>Recklessness has a {reck["cooldown_ms"] // 60000} minute cooldown, longer than this fight, so the result counts one use every pull. In a raid it is back only every {reck["cooldown_ms"] // 60000} minutes.</p></div>''']
    return f'''<section class="card quiet notes">{chead("alert", "This run", "Worth knowing", "1 note")}<div class="cbody">{"".join(rows)}</div></section>'''


def casts_bit(r):
    if r["casts"] is None:
        return ""
    return f'<i>{r["casts"]}</i> {"swings" if r["white"] else "casts"} · '


def ability_table():
    rows = D.ability_rows()
    peak = rows[0]["share"]
    body = []
    for r in rows:
        bits = []
        if r["periodic"]:
            bits.append(f'<i>{r["ticks"]}</i> ticks · tick crit <i>{pct(r["crit"])}</i>')
        else:
            split = " · ".join(f"{k} <i>{pct(v)}</i>" for k, v in r["split"])
            if split:
                bits.append(split)
        if r["key"] == "other:attack/1":
            bits.append(f'<i>{D.CASTS["spell:11567"]}</i> swings became Heroic Strike')
        sub = " · ".join(bits)
        ph = f'<div class="sub ph-only">{casts_bit(r)}<i>{num(r["avg"])}</i> avg · crit <i>{pct(r["crit"])}</i></div>'
        rank = f' <span class="rank">Rank {r["rank"]}</span>' if r["rank"] else ""
        label = r["name"]
        if r["white"]:
            w = D.weapon("main_hand" if r["key"].endswith("/1") else "off_hand")
            label = f'{r["name"]} <span class="rank">· <span style="color:{QUALITY_TEXT[w["quality"]]}">{E(w["name"])}</span></span>'
        body.append(f'''<tr><td><div class="ab">{icon(r["icon"], "", r["name"])}<div><div class="nm">{label}{rank}</div><div class="sub">{sub}</div>{ph}</div></div></td>
<td class="numeric"><span class="dps">{r["dps"]:.1f}</span><div class="sub ph-only dim">{pct(r["share"])}</div></td>
<td><div class="share"><div class="sbar"><i style="width:{r["share"] / peak * 100:.1f}%"></i></div><span class="pct numeric">{pct(r["share"])}</span></div></td>
<td class="numeric">{"" if r["casts"] is None else r["casts"]}<span class="dim">{"–" if r["casts"] is None else ""}</span></td>
<td class="numeric">{num(r["avg"])}</td><td class="numeric">{pct(r["crit"])}</td>
<td class="numeric">{pct(DW_UP) if r["periodic"] else '<span class="dim">–</span>'}</td></tr>''')
    body.append(f'<tr class="total"><td>Total, {len(rows)} sources</td><td class="numeric"><span class="dps">{sum(r["dps"] for r in rows):.1f}</span></td><td class="numeric" style="text-align:right">100%</td><td></td><td colspan="3" class="numeric">{num(D.TOTAL)} damage a fight</td></tr>')
    head = '<tr><th>Ability</th><th class="sorted">DPS</th><th>Share</th><th>Casts</th><th>Average</th><th>Crit</th><th>Uptime</th></tr>'
    return f'''<section class="card" id="damage">{chead("bars", "This run", "Damage by ability", f'<span class="numeric">{DPS["mean"]:.1f} DPS over {T.fmt(D.DUR_MS)[:-2]} on average</span>')}
<div class="cbody"><table class="at">{head}{"".join(body)}</table></div></section>'''


def spread():
    counts, lo, hi, width = D.histogram()
    peak = max(counts)
    cum, total = 0, sum(counts)
    p16 = p84 = None
    for i, c in enumerate(counts):
        cum += c
        if p16 is None and cum >= total * .16:
            p16 = lo + width * i
        if p84 is None and cum >= total * .84:
            p84 = lo + width * (i + 1)
    bars = "".join(f'<i class="{"in" if p16 <= lo + width * (i + .5) <= p84 else ""}" style="height:{max(1, c / peak * 100):.1f}%"></i>' for i, c in enumerate(counts))
    mx = (DPS["mean"] - lo) / (hi - lo) * 100
    tiles = [("Mean", f'{DPS["mean"]:.1f}'), ("Error", f'± {DPS["error"]:.1f}'), ("Spread", f'± {DPS["stddev"]:.1f}'),
             ("Lowest", num(DPS["min"])), ("Highest", num(DPS["max"])), ("Runs", num(D.RUN["iterations_run"]))]
    t = "".join(f'<div class="st"><div class="v numeric">{v}</div><div class="k">{k}</div></div>' for k, v in tiles)
    return f'''<section class="card" id="spread">{chead("spark", "3,000 runs", "DPS spread", '<span class="illu">Illustrative bins</span>')}
<div class="cbody hist"><div class="hbars" style="margin-top:22px">{bars}<span class="mean" style="left:{mx:.1f}%"></span><span class="meanlab numeric" style="left:{mx:.1f}%">{DPS["mean"]:.1f}</span></div>
<div class="haxis numeric"><span>{num(lo)}</span><span>DPS per run</span><span>{num(hi)}</span></div>
<p class="hcap">Two runs in three land between <b class="numeric">{num(p16)}</b> and <b class="numeric">{num(p84)}</b> DPS. Your pulls will spread the same way.</p>
<div class="hstats">{t}</div></div></section>'''


DW_SEGS = [[a, b + 3000] for a, b in D.merged(D.sample_times("spell:12867"), 3100)]
DW_UP = sum(min(b, D.SAMPLE_END) - a for a, b in DW_SEGS) / D.SAMPLE_END


def timeline(plot_w, label_w, mode="whole", swings=False, crosshair=None, extra=""):
    if mode == "whole":
        t0, t1, step = -3000, D.SAMPLE_END + 1000, 15000
        if plot_w < 600:
            step = 30000
        w0 = w1 = None
    else:
        t0, t1, step = 150000, 180000, 5000
        w0, w1 = t0, t1
    svg, hgt = T.plot(plot_w, t0, t1, swings, step, crosshair)
    br = T.brush(plot_w, -3000, D.SAMPLE_END + 1000, w0, w1)
    seg = f'<span class="seg"><span class="{"on" if mode == "whole" else ""}">Whole fight</span><span class="{"on" if mode != "whole" else ""}">Close up</span></span>'
    sw = f'<span class="sw"><span class="tr{" on" if swings else ""}"></span>Swings and ticks</span>'
    legend = ('<span class="legend"><span><i class="lg-t"></i>Damage</span><span><i class="lg-c"></i>Cooldown</span>'
              '<span><i class="lg-b"></i>On you</span><span><i class="lg-d"></i>On the target</span><span><i class="lg-r"></i>Rage</span></span>')
    meta = f'<span class="numeric">The median run · {T.fmt(D.SAMPLE_END)[:-2]} long</span>'
    return f'''<section class="card" id="timeline">{chead("clock", "One fight", "Timeline", meta)}
<div class="tl-tools">{seg}{sw}{legend}</div>
<div class="tl"><div class="tl-grid"><div class="tl-lab">{T.labels(swings)}</div>
<div class="tl-plot" style="overflow:hidden">{svg}{extra}</div></div>
<div class="tl-grid tl-brush"><div class="tl-lab"></div><div>{br}</div></div></div>
{cast_list(t0, t1, mode)}</section>'''


def cast_list(t0, t1, mode):
    rows = [x for x in D.SAMPLE if not x["action"].startswith("other:") and x["action"] not in ("spell:12867", "spell:11567")]
    inview = [x for x in rows if t0 <= x["at_ms"] <= t1]
    show = inview[:5]
    body = "".join(f'<div class="cl-row"><span class="tm numeric">{T.fmt(x["at_ms"])}</span>{icon(D.icon_of(x["action"]), "s24")}<span>{E(D.name_of(x["action"]))}{" queued" if x["action"].endswith("/1") else ""}</span><span class="rg numeric">Rage {x["resources"]["rage"]}</span></div>' for x in show)
    label = "Casts in view" if mode != "whole" else "Casts"
    return f'''<div class="castlist"><div class="foot" style="border-top:0"><span class="ql">{label}: {len(inview)} of {len(rows)}. Swings and Deep Wounds ticks are on the lanes above.</span><span class="btn s sm" style="margin-left:auto">Show all {len(inview)}</span></div>{body}</div>'''


def paperdoll():
    gear = {g["slot"]: g for g in D.REQ["character"]["gear"]}
    labels = {"head": "Head", "neck": "Neck", "shoulder": "Shoulder", "back": "Back", "chest": "Chest", "wrist": "Wrist",
              "hands": "Hands", "waist": "Waist", "legs": "Legs", "feet": "Feet", "finger1": "Finger 1", "finger2": "Finger 2",
              "trinket1": "Trinket 1", "trinket2": "Trinket 2", "main_hand": "Main hand", "off_hand": "Off hand", "ranged": "Ranged"}

    def slot(s, right=False):
        g = gear[s]
        it = D.ITEMS[g["item_id"]]
        e = D.ENCHANTS.get(g.get("enchant"))
        en = ""
        if e:
            nm = e["name"].split(" - ", 1)[-1]
            en = f'<div class="e"><svg viewBox="0 0 24 24" width="12" height="12" fill="none" stroke="currentColor" stroke-width="2.4" stroke-linecap="round"><path d="M12 3v6M12 15v6M3 12h6M15 12h6"/></svg> {E(nm)}</div>'
        return (f'<div class="slot{" r" if right else ""}"><div class="sock" style="--q:{QUALITY[it["quality"]]}"><img src="{D.icon_path(it["icon"]).as_uri()}" alt=""></div>'
                f'<div class="t"><div class="n" style="--qt:{QUALITY_TEXT[it["quality"]]}">{E(it["name"])}</div>{en}<div class="l numeric">{labels[s]} · {it["item_level"]}</div></div></div>')
    left = "".join(slot(s) for s in ["head", "neck", "shoulder", "back", "chest", "wrist"])
    right = "".join(slot(s, True) for s in ["hands", "waist", "legs", "feet", "finger1", "finger2", "trinket1", "trinket2"])
    weap = "".join(slot(s) for s in ["main_hand", "off_hand", "ranged"])
    ilvls = [D.ITEMS[g["item_id"]]["item_level"] for g in D.REQ["character"]["gear"]]
    mid = f'<div class="pd-mid">{crest()}<div><div class="nm">Orc Fury Warrior</div><div class="ql">Average item level</div><div class="ilv numeric">{sum(ilvls) / len(ilvls):.1f}</div></div></div>'
    return f'''<section class="card" id="gear">{chead("shield", "This run's input", "Gear", '<span class="numeric">17 of 17 slots · 11 enchants</span>')}
<div class="cbody pad"><div class="stage"><div class="pd"><div class="pd-col">{left}</div>{mid}<div class="pd-col">{right}</div></div><div class="pd-weap">{weap}</div></div></div>
<div class="foot"><span class="note">No set bonus is active: no two pieces share a set.</span><a class="alink" style="margin-left:auto" href="#">Upgrades for this character on Best in slot</a></div></section>'''


def talents(string, title_meta):
    trees = D.talent_trees(string)
    cols = []
    for t in trees:
        cells = []
        for tier in range(t["tiers"]):
            for col in range(4):
                c = t["cells"].get((tier, col))
                if not c:
                    cells.append('<div class="tc empty"></div>')
                    continue
                tal, r = c
                cls = "off" if r == 0 else ("full" if r == tal["max_rank"] else "part")
                badge = f'<b class="numeric">{r}/{tal["max_rank"]}</b>' if r else ""
                cells.append(f'<div class="tc {cls}" title="{E(tal["name"])}"><img src="{D.icon_path(tal["icon"]).as_uri()}" alt="{E(tal["name"])}">{badge}</div>')
        bg = (D.DATA / "trees" / f'{t["bg"]}.webp').as_uri() if t["bg"] else ""
        alt = "" if t["points"] == max(x["points"] for x in trees) else " alt"
        cols.append(f'<div class="tree{alt}" style="background-image:url({bg})"><h3>{E(t["name"])}<span class="numeric">{t["points"]}</span></h3><div class="tgrid">{"".join(cells)}</div></div>')
    split = " / ".join(str(t["points"]) for t in trees)
    top = max(x["points"] for x in trees)
    tabs = "".join(f'<span class="{"on" if t["points"] == top else ""}">{E(t["name"])} <b class="numeric">{t["points"]}</b></span>' for t in trees)
    return f'''<section class="card" id="talents">{chead("tree", "This run's input", "Talents", title_meta.format(split=split))}
<div class="cbody pad"><div class="ttabs ph-flex">{tabs}</div><div class="trees">{"".join(cols)}</div></div>
<div class="foot"><span class="note">Read only. Hover or focus a talent for its tooltip.</span><a class="alink" style="margin-left:auto" href="#">Open in planner</a></div></section>'''


def resources():
    g, sp, wa = (D.SUMMARY["resources"][0][k] for k in ("gained", "spent", "wasted"))
    tiles = [("Gained", num(g)), ("Spent", num(sp)), ("Lost at the cap", f"{wa} · {wa / g * 100:.1f}%")]
    t = "".join(f'<div class="st"><div class="v numeric">{v}</div><div class="k">{k}</div></div>' for k, v in tiles)
    # ILLUSTRATIVE: per-source rows are new data (spec 11); the split is modelled to sum to the run's totals.
    gained = [("other:attack/1", 512), ("other:attack/2", 431), ("spell:2687", 80), (None, 86)]
    spent = [("spell:20662", 408), ("spell:1680", 350), ("spell:11567", 336), ("spell:12328", 10)]

    def rows(lst, total):
        peak = max(v for _, v in lst)
        out = []
        for k, v in lst:
            nm = D.name_of(k) if k else "Unbridled Wrath"
            ic = D.icon_of(k) if k else "spell_nature_stoneclawtotem"
            out.append(f'<div class="lrow nob">{icon(ic, "s24")}<span>{E(nm)}</span><div class="bar"><i style="width:{v / peak * 100:.0f}%"></i></div><span class="v numeric">{v}</span></div>')
        return "".join(out)
    return f'''<section class="card" id="resource">{chead("bolt", "Average fight", "Rage", '<span class="illu">Illustrative sources</span>')}
<div class="cbody"><div class="tiles3">{t}</div><div class="two"><div><div class="subh"><span>Gained from</span><span class="numeric">{num(g)}</span></div>{rows(gained, g)}</div>
<div><div class="subh"><span>Spent on</span><span class="numeric">{num(sp)}</span></div>{rows(spent, sp)}</div></div></div></section>'''


def auras():
    perm, rows = D.aura_rows()
    you = "".join(f'<div class="lrow">{icon(r["icon"], "s24", r["name"])}<span>{E(r["name"])}</span><div class="bar b"><i style="width:{r["up"] * 100:.1f}%"></i></div><span class="v numeric">{pct(r["up"])}</span><span class="c numeric">{r["n"]}×</span></div>' for r in rows)
    up = "".join(f'<span>{icon(r["icon"], "s20", r["name"])}{E(r["name"])}</span>' for r in perm)
    debuff_keys = [b for b in D.REQ["character"]["buffs"] if b in TARGET_DEBUFFS]
    tgt = f'<div class="lrow">{icon("ability_backstab", "s24")}<span>Deep Wounds</span><div class="bar d"><i style="width:{DW_UP * 100:.1f}%"></i></div><span class="v numeric">{pct(DW_UP)}</span><span class="c numeric">{len(DW_SEGS)}×</span></div>'
    tup = "".join(f'<span>{icon(D.SIMBUFFS[b]["icon"], "s20")}{E(D.SIMBUFFS[b]["name"])}</span>' for b in debuff_keys[:8])
    return f'''<section class="card" id="buffs">{chead("spark", "Average fight", "Buffs and debuffs", "Uptime, then times applied")}
<div class="cbody two"><div><div class="subh"><span>On you</span><span>{len(rows)} that come and go</span></div>{you}
<div class="ql" style="padding-top:14px">Up all fight</div><div class="uprow">{up}</div></div>
<div><div class="subh"><span>On the target</span><span class="illu">Illustrative</span></div>{tgt}
<div class="ql" style="padding-top:14px">Up all fight, from the raid</div><div class="uprow">{tup}<span class="ql">and {len(debuff_keys) - 8} more</span></div></div></div></section>'''


TARGET_DEBUFFS = {"curse_of_elements", "curse_of_recklessness", "curse_of_shadow", "curse_of_weakness", "demoralizing_roar",
                  "demoralizing_shout", "expose_armor", "faerie_fire", "improved_scorch", "improved_shadow_bolt", "insect_swarm",
                  "judgement_of_light", "judgement_of_the_crusader", "judgement_of_wisdom", "scorpid_sting", "shadow_weaving",
                  "stormstrike", "sunder_armor", "thunder_clap", "winters_chill"}


def rank_html(sid):
    return f' <span class="rank">Rank {sid["rank"]}</span>' if sid.get("rank") else ""


def rotation():
    pre, lines = D.apl_lines()
    cast_keys = {int(k.split(":")[1].split("/")[0]) for k, n in D.CASTS.items() if k.startswith("spell:") and n > 0}
    covered = set()
    out = ['<div class="rhead">Before the pull</div>']
    for sid, at in pre:
        key = f'spell:{sid["spellId"]}'
        out.append(f'<div class="rl"><span class="ix numeric" style="white-space:nowrap">{E(at.replace("s", "\u00a0s"))}</span>{icon(D.icon_of(key), "s28")}<div><div class="nm">{E(D.name_of(key))}{rank_html(sid)}</div></div><span></span></div>')
    out.append('<div class="rhead">Then, from the top at every decision</div>')
    i = 0
    for sid, note in lines:
        i += 1
        if sid is None:
            others = [k for k in D.CASTS if k.startswith("spell:") and int(k.split(":")[1].split("/")[0]) not in {l[0]["spellId"] for l in lines if l[0]} and "/" not in k]
            names = ", ".join(D.name_of(k) for k in others)
            out.append(f'<div class="rl"><span class="ix numeric">{i}</span>{icon(D.icon_of(others[0]) if others else None, "s28")}<div><div class="nm">Other cooldowns as they come up</div><div class="w">This run: {E(names)}.</div></div><span class="ct numeric">{sum(D.CASTS[k] for k in others)} casts</span></div>')
            continue
        key = f'spell:{sid["spellId"]}'
        covered.add(sid["spellId"])
        n = D.CASTS.get(key, 0)
        rank = f' <span class="rank">Rank {sid["rank"]}</span>' if sid.get("rank") else ""
        off = n == 0
        ct = "Not cast in this run" if off else f"{n} cast{'' if n == 1 else 's'}"
        out.append(f'<div class="rl{" off" if off else ""}"><span class="ix numeric">{i}</span>{icon(D.icon_of(key), "s28")}<div><div class="nm">{E(D.name_of(key))}{rank}<span class="ct-ph numeric">{ct}</span></div><div class="w">{E(note)}</div></div><span class="ct numeric">{ct}</span></div>')
    return f'''<section class="card" id="rotation">{chead("list", "Fury, at 60", "Rotation", f"{len(lines)} lines · casts are per fight")}<div class="cbody">{"".join(out)}</div>
<div class="foot"><span class="note">The Fury rotation the engine follows. A line with no cast in this run is dimmed.</span><a class="alink" style="margin-left:auto" href="#">Read the Fury guide</a></div></section>'''


def settings():
    kv = [("Fight length", f'3:00, ± 20% (2:24 to 3:36)'), ("Targets", "1"), ("Execute phase", "Last 25%"),
          ("Race", "Orc"), ("Runs", "3,000"), ("Ran on", "Our server"), ("Engine", '<span class="mono">d5c7088</span>')]
    rows = "".join(f'<div class="kv"><span class="k">{k}</span><span class="v numeric">{v}</span></div>' for k, v in kv)
    buffs = [b for b in D.REQ["character"]["buffs"] if b not in TARGET_DEBUFFS]
    debs = [b for b in D.REQ["character"]["buffs"] if b in TARGET_DEBUFFS]

    def chips(lst):
        return "".join(f'<span class="chip{" more" if i >= 6 else ""}"><img src="{D.icon_path(D.SIMBUFFS[b]["icon"]).as_uri()}" alt="">{E(D.SIMBUFFS[b]["name"])}</span>' for i, b in enumerate(lst) if b in D.SIMBUFFS) + f'<span class="btn s sm ph-flex">Show all {len(lst)}</span>'
    return f'''<section class="card" id="settings">{chead("gear", "This run's input", "Settings", "")}
<div class="cbody two"><div>{rows}</div><div><div class="subh"><span>On you</span><span class="numeric">{len(buffs)}</span></div><div class="chips">{chips(buffs)}</div>
<div class="subh"><span>On the target</span><span class="numeric">{len(debs)}</span></div><div class="chips">{chips(debs)}</div>
<div class="subh"><span>Consumables</span><span class="numeric">0</span></div><div class="chips"><span class="ql">None in this run.</span></div></div></div>
<div class="foot acts-foot"><span class="btn s sm">{glyph("copy")}Copy for the game</span><span class="btn s sm">{glyph("copy")}Copy JSON</span><span class="btn s sm">{glyph("down")}Download JSON</span><span class="note">The JSON is this whole result, request included.</span></div></section>'''


def page(width, label):
    plot_w = (min(width, 1440) - 96 - 2 - 200) if width >= 768 else 868
    tl = timeline(plot_w, 200)
    body = f'''{site()}{band()}{snav()}<main class="wrap"><div class="inner page">
{notes()}
<div class="grid-dmg">{ability_table()}{spread()}</div>
{tl}{paperdoll()}{talents(GUIDE_TALENTS_STR, '<span class="numeric">{split}</span>')}
<div>{resources()}</div>{auras()}{rotation()}{settings()}
</div></main>'''
    return doc(body, label, width)


def doc(body, label, width):
    return f'<!doctype html><html><head><meta charset="utf-8"><meta name="viewport" content="width={width}"><style>{CSS}</style></head><body><div class="board-label">{label}</div>{body}</body></html>'


GUIDE_TALENTS_STR = next(iter(json.dumps(GUIDE_TALENTS)[i:] for i in [0]))  # replaced below


def all_talents(o, out):
    """Every talents string in the published Fury BiS file; the level-60 band holds the most points."""
    if isinstance(o, dict):
        for k, v in o.items():
            if k == "talents" and isinstance(v, str) and v.count("-") == 2:
                out.append(v)
            else:
                all_talents(v, out)
    elif isinstance(o, list):
        for v in o:
            all_talents(v, out)
    return out


GUIDE_TALENTS_STR = max(all_talents(GUIDE_TALENTS, []), key=lambda t: sum(int(c) for c in t if c.isdigit()))
ILLU = ("<b>ILLUSTRATIVE</b> regions draw data the engine does not return yet (spec 11): DPS spread bins, rage sources, "
        "the target column, and the timeline's On you / On the target bars (derived here from cast times, spell durations and Deep Wounds ticks). "
        "The <b>Saved</b> date stands in for created_at (spec 11.2). <b>Talents</b> show the published Fury build: this run's stored string does not fit build 70790's trees (spec 4.7, open question 1).")


def closeup(width):
    if width >= 768:
        plot_w = min(width, 1440) - 96 - 2 - 200
    else:
        plot_w = width - 36 - 2 - 44
    t0, t1 = 150000, 180000
    X = lambda t: (t - t0) / (t1 - t0) * plot_w
    ch = 170000
    before = [x for x in D.SAMPLE if x["at_ms"] <= ch and not x["action"].startswith("other:") and x["action"] != "spell:12867"]
    last = before[-5:][::-1]
    rage = [v for t, v in D.rage_points() if t <= ch][-1]
    rec = "".join(f'<div class="r">{icon(D.icon_of(x["action"]), "s20")}<span>{E(D.name_of(x["action"]))}{" queued" if x["action"].endswith("/1") else ""}</span><span class="when numeric">{(ch - x["at_ms"]) / 1000:.1f} s ago</span></div>' for x in last)
    rx = X(ch) + 12 if X(ch) + 300 < plot_w else X(ch) - 292
    readout = (f'<div class="readout" style="left:{rx:.0f}px;top:40px"><div class="t numeric">{T.fmt(ch)}</div>'
               f'<div class="ql">Rage <b style="color:var(--strong)">{rage}</b> · On you: Recklessness</div><hr>{rec}</div>')
    ex = D.sample_times("spell:20662")
    tick = ex[2]
    nxt = ex[3]
    lane_y = T.AXIS + T.GROUP + T.CAST_KEYS.index("spell:20662") * T.LANE
    tip = (f'<div class="tip" style="left:{max(4, X(tick) - 260):.0f}px;top:{lane_y + 30}px"><div class="h">{icon("inv_sword_48", "s24")}Execute <span class="rank">Rank 5</span></div>'
           f'<div class="numeric">{T.fmt(tick)} · rage after <b>{[x for x in D.SAMPLE if x["at_ms"] == tick and x["action"] == "spell:20662"][0]["resources"]["rage"]}</b></div>'
           f'<div class="numeric dim">Next Execute {(nxt - tick) / 1000:.1f} s later</div><div class="dim" style="margin-top:4px">Hit or crit shows here once the run records it.</div></div>')
    if width < 768:
        tip = ""
        readout = (f'<div class="readout" style="position:static;width:auto;margin:12px 16px"><div class="t numeric">{T.fmt(ch)}</div>'
                   f'<div class="ql">Rage <b style="color:var(--strong)">{rage}</b> · On you: Recklessness</div><hr>{rec}</div>')
        body = timeline(plot_w, 44, mode="close", swings=True, crosshair=ch) + readout
    else:
        body = timeline(plot_w, 200, mode="close", swings=True, crosshair=ch, extra=readout + tip)
    label = (f"<b>TIMELINE, CLOSE UP</b> · {width}px · 30 s window (2:30 to 3:00) chosen on the brush · Swings and ticks on · "
             "crosshair at 2:50.0 with its readout · Execute tick tooltip open (desktop)")
    return doc(f'<main class="wrap"><div class="inner page">{body}</div></main>', label, width)


def states(width):
    sk = lambda w, hh, extra="": f'<div class="sk" style="width:{w};height:{hh}px;{extra}"></div>'
    loading = f'''<section class="card" style="height:520px">{chead("bars", "This run", "Damage by ability")}<div class="cbody pad" style="display:flex;flex-direction:column;gap:14px">{"".join(sk("100%", 34) for _ in range(9))}</div></section>'''
    tl_loading = f'''<section class="card">{chead("clock", "One fight", "Timeline", '<span class="numeric">The median run · 3:34.5 long</span>')}<div class="cbody" style="height:{T.AXIS + sum(T.height_of(r) for r in T.lanes(False)) + 56 + 44 + 52}px;padding:16px 20px">{sk("100%", 36)}<div style="height:16px"></div>{sk("100%", 380)}</div></section>'''
    old_tl = f'''<section class="card">{chead("clock", "One fight", "Timeline", '<span class="illu">Excerpt: where the On you and On the target groups sit</span>')}<div class="cbody pad"><p class="ql" style="margin:0">Buff and debuff bars need a run from 5 Nov 2026 or later. The casts and rage below are from this one. <a class="alink" href="#">Run this yourself</a></p></div></section>'''
    old_spread = f'''<section class="card">{chead("spark", "3,000 runs", "DPS spread")}<div class="cbody hist"><div style="position:relative;height:8px;background:var(--line-soft);border-radius:4px;margin-top:28px"><span style="position:absolute;left:30%;right:30%;top:0;bottom:0;background:rgba(229,185,85,.35);border-radius:4px"></span><span style="position:absolute;left:45.6%;top:-6px;width:2px;height:20px;background:var(--gold-light)"></span></div><div class="haxis numeric"><span>311</span><span>mean 402.3</span><span>511</span></div><p class="hcap">This run kept the range, not the shape. The shaded part is one spread either side of the mean.</p></div></section>'''
    old_res = f'''<section class="card">{chead("bolt", "Average fight", "Rage")}<div class="cbody"><div class="tiles3" style="padding-bottom:16px"><div class="st"><div class="v numeric">1,109</div><div class="k">Gained</div></div><div class="st"><div class="v numeric">1,104</div><div class="k">Spent</div></div><div class="st"><div class="v numeric">14 · 1.3%</div><div class="k">Lost at the cap</div></div></div><p class="ql" style="padding:0 20px 16px;margin:0">Where the rage came from and went needs a newer run.</p></div></section>'''
    old_tal = f'''<section class="card">{chead("tree", "This run's input", "Talents", '<span class="numeric">20 / 31 / 0</span>')}<div class="cbody pad"><p style="margin:0">These talents were saved on an older talent layout. The point split above is as saved; the trees have changed since, so they are not drawn.</p></div></section>'''
    err = f'''<section class="card">{chead("alert", "Saved sim", "This sim did not load")}<div class="cbody pad"><p style="margin:0 0 12px">The server did not answer. Your link is fine; try again in a moment.</p><span class="btn s">Try again</span></div></section>'''
    label = f"<b>STATES</b> · {width}px · loading skeletons at ready height · an older saved run (no new data) · talents on an older layout · load failed"
    grid = f'<div class="grid-dmg">{loading}<div style="display:flex;flex-direction:column;gap:16px">{old_spread}{err}</div></div>{tl_loading}{old_tl}{old_res}{old_tal}'
    return doc(f'<main class="wrap"><div class="inner page">{grid}</div></main>', label, width)


def main():
    OUT.mkdir(exist_ok=True)
    files = []
    for w in (1440, 390, 1024, 2000):
        p = OUT / f"result-{w}.html"
        p.write_text(page(w, f"<b>SAVED SIM, RUN KIND</b> · {w}px · real run: Fury warrior, 3000 runs, engine d5c708834 · {ILLU}"))
        files.append(p)
    for w in (1440, 390):
        p = OUT / f"timeline-{w}.html"
        p.write_text(closeup(w))
        files.append(p)
    p = OUT / "states-1440.html"
    p.write_text(states(1440))
    files.append(p)
    subprocess.run(["node", str(HERE / "shoot.mjs"), *map(str, files)], cwd=D.ROOT / "web", check=True)
    shrink(HERE.glob("*.png"))


def shrink(pngs):
    """The boards are tall 2x captures; a 256-colour palette keeps them reviewable and small in git."""
    try:
        from PIL import Image
    except ImportError:
        return
    for png in pngs:
        Image.open(png).convert("RGB").quantize(colors=256, method=Image.Quantize.FASTOCTREE, dither=Image.Dither.NONE).save(png, optimize=True)


if __name__ == "__main__":
    main()
