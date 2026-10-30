"""Boards for design/specs/2026-10-30-controls-charts.md. Run from the repo root:
python3 design/mocks/2026-10-30-controls-charts/gen_controls_charts.py && node design/mocks/2026-10-30-controls-charts/render.mjs
Draws on the 2026-10-28 polish parts (direction B, Forge): tokens, crests, glyphs.

One data model feeds every board (fix round 1): a 40-player Horde raid on a mock Ragnaros kill of
5:40, submerged (untargetable) from 3:00 to 4:30. Per-second damage and healing are generated once;
every chip, head figure, tooltip row, share and table row is computed from them, never typed. Threat
is the base-1 model the engine ships (logs/engine/summary/threat.go): 1 per point of damage, 0.5 per
point of effective healing split across the enemies in combat, no stance, form, talent or taunt
modifiers. Every label is the spec's copy."""
import math
import random
import statistics
import sys
from pathlib import Path

HERE = Path(__file__).resolve().parent
sys.path.insert(0, str(HERE.parent / "2026-10-28-polish"))
from parts import CLASS_COLOR, GLYPHS, crest, glyph, uri  # noqa: E402
from theme import head  # noqa: E402

RENDERS = HERE / "renders"
GLYPHS["skull"] = "<path d=\"M12 3a7 7 0 0 0-7 7c0 2.4 1.2 4 2.5 5v3h9v-3c1.3-1 2.5-2.6 2.5-5a7 7 0 0 0-7-7z\"/><circle cx=\"9.5\" cy=\"10.5\" r=\"1.3\"/><circle cx=\"14.5\" cy=\"10.5\" r=\"1.3\"/><path d=\"M10 18v3M14 18v3\"/>"
GLYPHS["target"] = "<circle cx=\"12\" cy=\"12\" r=\"8\"/><circle cx=\"12\" cy=\"12\" r=\"3\"/><path d=\"M12 1v4M12 19v4M1 12h4M19 12h4\"/>"
GLYPHS["pin"] = "<path d=\"M12 21s-6-5.6-6-11a6 6 0 0 1 12 0c0 5.4-6 11-6 11z\"/><circle cx=\"12\" cy=\"10\" r=\"2.2\"/>"


def icon(stem: str) -> str:
    return uri(HERE / "assets" / f"{stem}.webp")


TAUNT, DEATH_WISH, INNERVATE = icon("spell_nature_reincarnation"), icon("spell_shadow_deathpact"), icon("spell_nature_lightning")

CSS = """
.board{display:flex;flex-direction:column;min-height:100vh;background:#07090d}
.dirbar{display:flex;align-items:center;gap:12px;height:34px;padding:0 20px;background:#000;border-bottom:1px solid #1c2230;font-size:12px;color:#9a9484;font-family:"JetBrains Mono",monospace}
.dirbar b{color:#e5b955}
.cap{font-size:12px;color:#9a9484;font-family:"JetBrains Mono",monospace;margin:0 0 10px}
.num{font-variant-numeric:tabular-nums}
.ql{font-size:13px;font-weight:600;color:var(--muted)}
.eyebrow{display:flex;align-items:center;gap:12px;font-size:11px;font-weight:700;letter-spacing:.18em;text-transform:uppercase;color:var(--gold)}
.eyebrow::before{content:"";width:24px;height:1px;background:rgba(229,185,85,.6)}
.stitle{font-family:var(--fh);font-weight:800;font-size:18px;line-height:1;text-transform:uppercase;letter-spacing:.01em;color:var(--strong)}
/* the closed field */
.sel{position:relative;display:inline-flex;align-items:center;gap:8px;height:52px;padding:0 40px 0 12px;border-radius:8px;border:1px solid rgba(229,185,85,.17);background:rgba(255,255,255,.045);color:var(--strong);font-size:15px;font-weight:600;white-space:nowrap;min-width:180px}
.sel.c{height:44px}.sel.d{height:36px;font-size:13px;padding-left:10px;padding-right:34px;min-width:150px}
.sel .chv{position:absolute;right:12px;top:50%;transform:translateY(-50%);color:var(--muted);display:flex}
.sel.d .chv{right:10px}
.sel.hover{border-color:rgba(229,185,85,.45);background:rgba(255,255,255,.06)}.sel.hover .chv{color:var(--strong)}
.sel.focus{outline:2px solid var(--gold);outline-offset:2px}
.sel.open{border-color:var(--gold);background:rgba(255,255,255,.06);box-shadow:0 0 26px rgba(229,185,85,.32)}
.sel.open .chv{transform:translateY(-50%) rotate(180deg);color:var(--strong)}
.sel.act{border-color:var(--gold);background:#101624}
.sel.dis{opacity:.45}
.sel.err{border-color:var(--wipe)}
.sel.ph{color:var(--muted);font-weight:500}
.sel .crest{box-shadow:0 0 0 1.5px color-mix(in srgb,var(--c) 60%,transparent)}
.spin{width:14px;height:14px;border-radius:999px;border:2px solid rgba(229,185,85,.25);border-top-color:var(--gold)}
.errline{font-size:13px;color:var(--wipe);margin-top:6px}
/* the listbox */
.lb{position:absolute;z-index:5;width:max-content;min-width:100%;max-width:420px;border:1px solid rgba(229,185,85,.17);border-radius:12px;background:linear-gradient(180deg, rgba(229,185,85,.075), rgba(229,185,85,.012) 45%), #0d0f15;box-shadow:inset 0 1px 0 rgba(245,210,122,.45),0 16px 38px rgba(0,0,0,.55);padding:4px 0;overflow:hidden}
.lb .scroll{max-height:360px;overflow:hidden;position:relative}
.lb .fade{position:absolute;left:0;right:0;bottom:0;height:24px;background:linear-gradient(180deg,rgba(13,15,21,0),#0d0f15)}
.opt{position:relative;display:flex;align-items:center;gap:10px;min-height:36px;padding:6px 40px 6px 12px;font-size:14px;font-weight:600;color:var(--strong)}
.opt.ic{min-height:40px}
.lb.dense .opt{font-size:13px;min-height:32px}.lb.dense .opt.ic{min-height:36px}
.opt .box{flex:none;width:16px;height:16px;border-radius:4px;border:1px solid #4a4030;display:grid;place-items:center}
.opt .box.on{background:var(--gold);border-color:var(--gold);color:#0b0e14}
.opt .you{font-size:11px;font-weight:700;letter-spacing:.08em;text-transform:uppercase;color:var(--gold);border:1px solid rgba(229,185,85,.38);border-radius:6px;padding:0 6px;height:18px;display:inline-flex;align-items:center;margin-left:6px}
.opt .t{display:flex;flex-direction:column;gap:2px;min-width:0}
.opt .ds{font-size:12px;font-weight:500;color:var(--muted);line-height:1.35;white-space:normal;max-width:330px}
.opt .meta{margin-left:auto;padding-left:16px;font-size:13px;font-weight:500;color:var(--muted)}
.opt.act{background:rgba(229,185,85,.10);box-shadow:inset 2px 0 0 var(--gold)}
.opt .ck{position:absolute;right:12px;top:50%;transform:translateY(-50%);color:var(--gold);display:flex}
.opt.dis{opacity:.45}
.opt mark{background:none;color:var(--gold);font-weight:700}
.grp{padding:10px 12px 4px;font-size:11px;font-weight:700;letter-spacing:.14em;text-transform:uppercase;color:var(--muted)}
.grp.rule{border-top:1px solid var(--line-soft);margin-top:4px}
.lbsearch{display:flex;align-items:center;gap:8px;padding:6px 8px 8px;border-bottom:1px solid var(--line-soft);margin-bottom:4px}
.lbsearch .in{flex:1;display:flex;align-items:center;gap:8px;height:36px;padding:0 10px;border-radius:6px;border:1px solid var(--gold);background:#07090d;color:var(--strong);font-size:14px;outline:2px solid rgba(229,185,85,.35);outline-offset:0}
.lbsearch .in .caret{width:1px;height:16px;background:var(--gold);margin-left:-6px}
.lbsearch .n{font-size:12px;color:var(--muted);white-space:nowrap;padding-right:4px}
.lbempty{padding:12px;font-size:14px;color:var(--muted)}
/* the chart frame */
.frame{background:linear-gradient(180deg, rgba(229,185,85,.075), rgba(229,185,85,.012) 45%), #0d0f15;border:1px solid rgba(229,185,85,.17);border-radius:12px;box-shadow:0 16px 38px rgba(0,0,0,.55);position:relative}
.fhead{display:flex;align-items:center;gap:12px;min-height:48px;padding:10px 16px}
.fhead .meta{font-size:13px;color:var(--muted)}
.fhead .right{margin-left:auto;display:flex;align-items:center;gap:10px}
.legend{display:flex;flex-wrap:wrap;min-height:42px;gap:4px 6px;padding:8px 12px 6px;border-top:1px solid var(--line-soft)}
.chip{display:inline-flex;align-items:center;gap:7px;height:28px;padding:0 8px;border-radius:6px;font-size:13px;font-weight:600;white-space:nowrap}
.chip .sw{width:12px;height:3px;border-radius:2px}
.chip .v{font-weight:500;color:var(--muted)}
.chip.off{opacity:.4}.chip.off .sw{background:none!important;box-shadow:inset 0 0 0 1px currentColor;height:4px}
.chip.hov{background:rgba(229,185,85,.06)}
.chip .only{font-size:12px;font-weight:600;color:var(--gold);margin-left:2px}
.chip .pin{display:flex;color:var(--muted)}
.chip .you{font-size:10px;font-weight:700;letter-spacing:.08em;color:var(--gold)}
.chip.med .sw{background:repeating-linear-gradient(90deg,#e9e4d8 0 4px,transparent 4px 7px)!important;height:2px;opacity:.8}
.chip.more{color:var(--gold)}
.ghost{display:inline-flex;align-items:center;gap:6px;height:36px;padding:0 10px;border-radius:8px;color:var(--gold);font-size:13px;font-weight:700;letter-spacing:.07em;text-transform:uppercase}
.sw2{display:inline-flex;align-items:center;gap:8px;font-size:13px;color:var(--text)}
.sw2 .tr{position:relative;width:36px;height:20px;border-radius:999px;background:#232a3a;border:1px solid #303a50}
.sw2 .tr::after{content:"";position:absolute;top:2px;left:2px;width:14px;height:14px;border-radius:999px;background:#8a8576}
.tip{position:absolute;z-index:6;width:268px;border:1px solid rgba(229,185,85,.17);border-radius:10px;background:#0d0f15;box-shadow:inset 0 1px 0 rgba(245,210,122,.45),0 16px 38px rgba(0,0,0,.6);padding:8px 0 6px;pointer-events:none}
.tip .th{display:flex;align-items:baseline;gap:8px;padding:0 12px 6px;border-bottom:1px solid var(--line-soft);margin-bottom:4px}
.tip .th b{font-size:14px;color:var(--strong)}.tip .th span{font-size:12px;color:var(--muted);margin-left:auto}
.tip .ev{display:flex;align-items:center;gap:8px;padding:3px 12px;font-size:12px;color:var(--text)}
.tip .ev img{width:16px;height:16px;border-radius:3px}
.tip .r{display:flex;align-items:center;gap:8px;height:22px;padding:0 12px;font-size:13px;font-weight:600}
.tip .r .sw{width:3px;height:12px;border-radius:2px}
.tip .r .v{margin-left:auto;color:var(--strong);font-weight:600}
.tip .r .p{width:50px;text-align:right;color:var(--muted);font-weight:500;font-size:12px}
.tip .r.pin{background:rgba(229,185,85,.08)}
.tip .r.dead{color:var(--muted);font-weight:500}
.tip .more{padding:4px 12px 0;font-size:12px;color:var(--muted)}
.tip .ch{display:flex;padding:2px 12px 2px 23px;font-size:11px;color:var(--muted)}.tip .ch span:first-child{margin-left:auto}.tip .ch span:last-child{width:50px;text-align:right}
.tip .ev .t{width:30px;color:var(--muted);font-variant-numeric:tabular-nums}
.tip .ev.death{color:#ffb4b0}
.tip .ev .g{width:16px;height:16px;display:grid;place-items:center;border-radius:3px}
.cnote{font-size:12px;color:var(--muted);line-height:1.5}
.tabs{display:flex;gap:4px;border-bottom:1px solid var(--line-soft)}
.tabs span{padding:10px 12px;font-size:14px;font-weight:600;color:var(--nav);border-bottom:2px solid transparent}
.tabs span.on{color:var(--strong);border-color:var(--gold)}
.tbl{width:100%;border-collapse:collapse;font-size:14px}
.tbl th{font-size:13px;font-weight:600;color:var(--muted);text-align:left;padding:8px 12px}
.tbl td{padding:8px 12px;border-top:1px solid var(--line-soft)}
.tbl .r{text-align:right}
.bar{height:6px;border-radius:3px;background:rgba(255,255,255,.06);position:relative;width:220px}
.bar i{position:absolute;inset:0 auto 0 0;border-radius:3px}
.skel{background:linear-gradient(90deg,rgba(255,255,255,.035),rgba(255,255,255,.07),rgba(255,255,255,.035));border-radius:6px}
.sheet{background:#0d0f15;border:1px solid rgba(229,185,85,.17);border-radius:16px 16px 0 0;box-shadow:inset 0 1px 0 rgba(245,210,122,.45),0 -16px 38px rgba(0,0,0,.6)}
.srow{display:flex;align-items:center;gap:12px;min-height:52px;padding:6px 18px;border-top:1px solid var(--line-soft);font-size:15px;font-weight:600}
.srow .v{margin-left:auto;font-size:14px;font-weight:500;color:var(--muted)}
.fold{position:absolute;left:0;right:0;border-top:2px dashed #ff5a5a}
.fold span{position:absolute;right:6px;top:-18px;font-size:11px;color:#ff5a5a;font-family:"JetBrains Mono",monospace;background:#07090d;padding:0 4px}
.meas{position:absolute;left:-2px;font-size:10px;color:#7fd48a;font-family:"JetBrains Mono",monospace;writing-mode:vertical-rl;transform:rotate(180deg)}
.lbl{font-size:12px;color:var(--muted);font-family:"JetBrains Mono",monospace}
.overlay-note{position:absolute;font-size:11px;font-family:"JetBrains Mono",monospace;color:#7fd48a}
"""

# ===================================================================== the one data model
DURATION = 340
SUBMERGE = (180, 270)
PHASES = [("Ragnaros", 0, 180), ("Submerge", 180, 270), ("Ragnaros", 270, DURATION)]
SONS = 8  # Sons of Flame in combat while Ragnaros is submerged
VIEWER = "Lightningbolltz"
random.seed(60)

# name, class, role, damage per second on Ragnaros (healing per second for a healer)
ROSTER = [
    ("Brokk", "warrior", "tank", 270), ("Thandril", "warrior", "tank", 250),
    ("Garrosk", "warrior", "dps", 700), ("Durgan", "warrior", "dps", 660), ("Kargath", "warrior", "dps", 640),
    ("Rukh", "warrior", "dps", 610), ("Hrothgar", "warrior", "dps", 590), ("Bek", "warrior", "dps", 560),
    ("Velyn", "rogue", "dps", 650), ("Zaxis", "rogue", "dps", 620), ("Shadra", "rogue", "dps", 590), ("Kriv", "rogue", "dps", 540),
    ("Morrigan", "mage", "dps", 640), ("Ysolde", "mage", "dps", 610), ("Faelan", "mage", "dps", 590),
    ("Orrin", "mage", "dps", 560), ("Tamsin", "mage", "dps", 530), ("Quill", "mage", "dps", 500),
    ("Ashka", "warlock", "dps", 630), ("Vexa", "warlock", "dps", 600), ("Malrik", "warlock", "dps", 570),
    ("Sorn", "warlock", "dps", 540), ("Ilyth", "warlock", "dps", 500),
    ("Kael", "hunter", "dps", 480), ("Torva", "hunter", "dps", 450), ("Brenn", "hunter", "dps", 430), ("Hask", "hunter", "dps", 400),
    ("Lightningbolltz", "shaman", "dps", 420),
    ("Zunja", "shaman", "dps", 450), ("Tavi", "shaman", "dps", 410), ("Mokra", "shaman", "healer", 380),
    ("Seraphine", "priest", "healer", 420), ("Elowen", "priest", "healer", 400), ("Aldric", "priest", "healer", 370),
    ("Nimue", "priest", "healer", 350), ("Corwin", "priest", "healer", 320), ("Lysa", "priest", "healer", 300),
    ("Oakhorn", "druid", "healer", 360), ("Fennick", "druid", "healer", 340), ("Bramble", "druid", "healer", 310),
]
assert len(ROSTER) == 40
CLS = {r[0]: r[1] for r in ROSTER}
ROLE = {r[0]: r[2] for r in ROSTER}
NAMES = [r[0] for r in ROSTER]

# Deaths, with the killing blow from the Deaths tab's data (summary.Death.killing_blow).
DEATHS = [
    {"name": "Brokk", "at": 126, "spell": "Melee", "source": "Ragnaros", "amount": 3862, "overkill": 1104},
]
DEAD_AT = {d["name"]: d["at"] for d in DEATHS}
TAUNTS = [(76, "Brokk"), (127, "Thandril")]  # (second, who taunted Ragnaros)
# Inferred from Ragnaros's melee swings: the unit his swings land on (spec §B5.6).
TARGET_CHANGES = [(72, "Thandril"), (76, "Brokk"), (127, "Thandril")]
CASTS = [(6, "Garrosk", DEATH_WISH, "Death Wish"), (276, "Garrosk", DEATH_WISH, "Death Wish"), (100, "Oakhorn", INNERVATE, "Innervate")]


def boss_targetable(t: int) -> bool:
    return not (SUBMERGE[0] <= t < SUBMERGE[1])


def per_second():
    """Damage done to Ragnaros and healing done, per player per second."""
    dmg, heal = {}, {}
    for name, cls, role, rate in ROSTER:
        d, h = [], []
        for t in range(DURATION):
            alive = t < DEAD_AT.get(name, DURATION + 1)
            ramp = min(1.0, (t + 1) / 4)
            noise = 0.55 + random.random() * 0.9
            if role == "healer":
                d.append(0.0)
                h.append(rate * noise * ramp if alive else 0.0)
                continue
            v = rate * noise * ramp if (alive and boss_targetable(t)) else 0.0
            if name == "Garrosk" and (6 <= t < 36 or 276 <= t < 306):
                v *= 1.35  # Death Wish
            d.append(v)
            h.append(0.0)
        dmg[name], heal[name] = d, h
    return dmg, heal


DMG, HEAL = per_second()


def enemies_in_combat(t: int) -> int:
    return 1 if boss_targetable(t) else 1 + SONS


def threat_on_rag():
    """base-1: damage to Ragnaros 1:1; effective healing at 0.5, split across the enemies in combat."""
    out = {}
    for name in NAMES:
        total, line = 0.0, []
        for t in range(DURATION):
            total += DMG[name][t] + 0.5 * HEAL[name][t] / enemies_in_combat(t)
            line.append(total)
        out[name] = line
    return out


THREAT = threat_on_rag()


def window_dps(name: str, w) -> int:
    """The legend's and the table's figure: damage over the window, divided by the window's seconds."""
    return round(sum(DMG[name][w[0]:w[1]]) / (w[1] - w[0]))


def smooth(values, n=5):
    return [sum(values[max(0, t - n + 1):t + 1]) / len(values[max(0, t - n + 1):t + 1]) for t in range(len(values))]


SMOOTH = {n: smooth(DMG[n]) for n in NAMES}
MEDIAN = [statistics.median(SMOOTH[n][t] for n in NAMES if ROLE[n] == "dps") for t in range(DURATION)]


def standing(name: str, t: int) -> float:
    return THREAT[name][min(t, DURATION - 1)]


def share(name: str, t: int) -> float:
    total = sum(standing(n, t) for n in NAMES)
    return standing(name, t) / total * 100


def holder_at(t: int) -> str:
    who = "Brokk"
    for at, name in TARGET_CHANGES:
        if at <= t:
            who = name
    return who


def threat_drawn(w, top_n=8, source=None):
    """Top N by standing at the window's end, plus whoever Ragnaros attacked in the window and whoever
    taunted in it, plus the Source player (spec §B5.3 rule 5)."""
    by = sorted(NAMES, key=lambda n: -standing(n, w[1] - 1))
    pinned = {holder_at(w[0])} | {n for at, n in TARGET_CHANGES if w[0] <= at < w[1]} | {n for at, n in TAUNTS if w[0] <= at < w[1]}
    if source:
        pinned.add(source)
    drawn = by[:top_n] + [n for n in by if n in pinned and n not in by[:top_n]]
    return sorted(drawn, key=lambda n: -standing(n, w[1] - 1)), pinned


def damage_drawn(w, top_n=8):
    by = sorted(NAMES, key=lambda n: -window_dps(n, w))
    pinned = {VIEWER}
    drawn = by[:top_n] + [n for n in by if n in pinned and n not in by[:top_n]]
    return drawn, pinned


def mmss(t: float) -> str:
    t = int(round(t))
    return f"{t // 60}:{t % 60:02d}"


def kfmt(v: float) -> str:
    if v >= 1000 and v % 1000 == 0:
        return f"{v / 1000:.0f}k"
    return f"{v / 1000:.1f}k" if v >= 1000 else f"{v:.0f}"


def nice_step(span: float, width: float) -> int:
    for step in (5, 10, 15, 30, 60, 120, 300):
        if width / (span / step) >= 64:
            return step
    return 600


def nice_max(v: float, ticks=4):
    raw = v / ticks
    mag = 10 ** math.floor(math.log10(raw))
    for m in (1, 1.5, 2, 2.5, 3, 4, 5, 10):
        if raw <= m * mag:
            return m * mag * ticks, [m * mag * i for i in range(ticks + 1)]
    return v, [v]


def colour(name: str) -> str:
    return CLASS_COLOR[CLS[name]]


# Sanity: the head figure is the sum of every player's (rounded) chip, and the threat leader agrees.
FULL = (0, DURATION)
RAID_DPS = sum(window_dps(n, FULL) for n in NAMES)


# ===================================================================== drawing
LANE_H, AXIS_H = 26, 22


def lane_tile(x, top, src=None, glyph_name=None, death=False, count=None):
    out = []
    cy = top / 2
    if death:
        out.append(f"<circle cx=\"{x:.1f}\" cy=\"{cy:.1f}\" r=\"10\" fill=\"#5a1f1f\" stroke=\"#ff5a5a\" stroke-opacity=\".6\"/>")
        out.append(f"<g transform=\"translate({x - 7:.1f},{cy - 7:.1f}) scale(.5833)\" fill=\"none\" stroke=\"#ff5a5a\" stroke-width=\"2.4\" stroke-linecap=\"round\">{GLYPHS['skull']}</g>")
    elif count is not None:
        out.append(f"<circle cx=\"{x:.1f}\" cy=\"{cy:.1f}\" r=\"10\" fill=\"#5a1f1f\" stroke=\"#ff5a5a\"/>")
        out.append(f"<text x=\"{x:.1f}\" y=\"{cy + 4:.1f}\" font-size=\"11\" font-weight=\"700\" fill=\"#ffb4b0\" text-anchor=\"middle\">{count}</text>")
    elif glyph_name:
        out.append(f"<rect x=\"{x - 10:.1f}\" y=\"{cy - 10:.1f}\" width=\"20\" height=\"20\" rx=\"4\" fill=\"#0d111a\" stroke=\"#e9e4d8\" stroke-opacity=\".5\"/>")
        out.append(f"<g transform=\"translate({x - 7:.1f},{cy - 7:.1f}) scale(.5833)\" fill=\"none\" stroke=\"#e9e4d8\" stroke-width=\"2.4\" stroke-linecap=\"round\">{GLYPHS[glyph_name]}</g>")
    else:
        out.append(f"<rect x=\"{x - 10:.1f}\" y=\"{cy - 10:.1f}\" width=\"20\" height=\"20\" rx=\"4\" fill=\"#0d111a\" stroke=\"#e5b955\" stroke-opacity=\".7\"/>")
        out.append(f"<image href=\"{src}\" x=\"{x - 9:.1f}\" y=\"{cy - 9:.1f}\" width=\"18\" height=\"18\"/>")
    return "".join(out)


def chart_svg(w, plot_h, win, series, *, yticks, ymax, markers=(), cursor=None, drag=None, dim=None, gutter=52,
              median=None, deaths=True, wipe=None):
    """series: (name, colour, values per second, dead_at or None). markers: (second, kind, payload)."""
    x0, x1 = win
    right = 12
    pw = w - gutter - right
    top = LANE_H
    h = top + plot_h + AXIS_H
    fx = lambda t: gutter + (t - x0) / (x1 - x0) * pw  # noqa: E731
    fy = lambda v: top + plot_h - (min(v, ymax) / ymax) * plot_h  # noqa: E731
    cid = f"pc{w}{plot_h}{x0}{x1}"
    s = [f"<svg width=\"{w}\" height=\"{h}\" viewBox=\"0 0 {w} {h}\" style=\"display:block;font-family:Barlow,sans-serif\">",
         f"<defs><clipPath id=\"{cid}\"><rect x=\"{gutter}\" y=\"{top}\" width=\"{pw}\" height=\"{plot_h}\"/></clipPath></defs>"]
    for i, (name, a, b) in enumerate(PHASES):
        a2, b2 = max(a, x0), min(b, x1)
        if b2 <= a2:
            continue
        if i % 2 == 1:
            s.append(f"<rect x=\"{fx(a2):.1f}\" y=\"{top}\" width=\"{fx(b2) - fx(a2):.1f}\" height=\"{plot_h}\" fill=\"rgba(229,185,85,.05)\"/>")
        if a > x0:
            s.append(f"<line x1=\"{fx(a):.1f}\" x2=\"{fx(a):.1f}\" y1=\"{top}\" y2=\"{top + plot_h}\" stroke=\"#e5b955\" stroke-opacity=\".7\" stroke-dasharray=\"2 3\"/>")
        lw = len(name) * 8.2 + 8
        s.append(f"<rect x=\"{fx(a2) + 6:.1f}\" y=\"{top + 4}\" width=\"{lw:.0f}\" height=\"16\" rx=\"3\" fill=\"#0d0f15\" fill-opacity=\".9\"/>")
        s.append(f"<text x=\"{fx(a2) + 10:.1f}\" y=\"{top + 16}\" font-size=\"11\" font-weight=\"700\" letter-spacing=\".12em\" fill=\"#9a9484\">{name.upper()}</text>")
    for v in yticks:
        y = fy(v)
        s.append(f"<line x1=\"{gutter}\" x2=\"{w - right}\" y1=\"{y:.1f}\" y2=\"{y:.1f}\" stroke=\"#1c2230\"/>")
        s.append(f"<text x=\"{gutter - 8}\" y=\"{y + 4:.1f}\" font-size=\"11\" fill=\"#9a9484\" text-anchor=\"end\" style=\"font-variant-numeric:tabular-nums\">{kfmt(v) if v else '0'}</text>")
    step = nice_step(x1 - x0, pw)
    t = math.ceil(x0 / step) * step
    while t <= x1 + 0.01:
        x = fx(t)
        s.append(f"<line x1=\"{x:.1f}\" x2=\"{x:.1f}\" y1=\"{top + plot_h}\" y2=\"{top + plot_h + 4}\" stroke=\"#262e40\"/>")
        anchor = "start" if t == x0 else "end" if abs(t - x1) < .01 else "middle"
        s.append(f"<text x=\"{x:.1f}\" y=\"{top + plot_h + 17}\" font-size=\"11\" fill=\"#9a9484\" text-anchor=\"{anchor}\" style=\"font-variant-numeric:tabular-nums\">{mmss(t)}</text>")
        t += step
    s.append(f"<line x1=\"{gutter}\" x2=\"{w - right}\" y1=\"{top + plot_h + .5}\" y2=\"{top + plot_h + .5}\" stroke=\"#262e40\"/>")
    s.append(f"<g clip-path=\"url(#{cid})\">")
    if median is not None:
        pts = [(fx(i), fy(v)) for i, v in enumerate(median) if x0 - 1 <= i <= x1 + 1]
        s.append(f"<path d=\"M{' L'.join(f'{x:.1f},{y:.1f}' for x, y in pts)}\" fill=\"none\" stroke=\"#e9e4d8\" stroke-opacity=\".7\" stroke-width=\"1.5\" stroke-dasharray=\"5 4\"/>")
    for name, col, vals, dead in series:
        end = dead if dead is not None else len(vals)
        pts = [(fx(i), fy(vals[i])) for i in range(len(vals)) if x0 - 1 <= i <= min(x1 + 1, end)]
        if not pts:
            continue
        wid, op = 1.5, 1
        if dim is not None:
            wid, op = (2.5, 1) if name == dim else (1.5, .25)
        s.append(f"<path d=\"M{' L'.join(f'{x:.1f},{y:.1f}' for x, y in pts)}\" fill=\"none\" stroke=\"{col}\" stroke-width=\"{wid}\" stroke-opacity=\"{op}\" stroke-linejoin=\"round\"/>")
        if dead is not None and x0 <= dead <= x1:  # the line stops at the death; a hollow end
            s.append(f"<circle cx=\"{fx(dead):.1f}\" cy=\"{fy(vals[dead]):.1f}\" r=\"3\" fill=\"#07090d\" stroke=\"{col}\" stroke-width=\"1.5\"/>")
    s.append("</g>")
    # markers: deaths first (the plot line), then lane tiles with the collision rule
    placed = []
    for at, kind, payload in sorted(markers, key=lambda m: (m[0], m[1] != "death")):
        if not x0 <= at <= x1:
            continue
        x = fx(at)
        if kind == "death":
            s.append(f"<line x1=\"{x:.1f}\" x2=\"{x:.1f}\" y1=\"{top}\" y2=\"{top + plot_h}\" stroke=\"#ff5a5a\" stroke-width=\"1.5\" stroke-dasharray=\"4 3\"/>")
        elif kind == "target":
            s.append(f"<line x1=\"{x:.1f}\" x2=\"{x:.1f}\" y1=\"{top}\" y2=\"{top + plot_h}\" stroke=\"#e9e4d8\" stroke-opacity=\".35\" stroke-dasharray=\"1 3\"/>")
        else:
            s.append(f"<line x1=\"{x:.1f}\" x2=\"{x:.1f}\" y1=\"{top}\" y2=\"{top + plot_h}\" stroke=\"#e5b955\" stroke-opacity=\".35\"/>")
        lx = x
        while any(abs(lx - p) < 22 for p in placed):
            lx += 22
        if lx != x:
            s.append(f"<line x1=\"{x:.1f}\" x2=\"{lx:.1f}\" y1=\"{top - 1}\" y2=\"{top / 2 + 10:.1f}\" stroke=\"#e5b955\" stroke-opacity=\".6\"/>")
        placed.append(lx)
        s.append(lane_tile(lx, top, src=payload if kind in ("taunt", "cast") else None, glyph_name="target" if kind == "target" else None, death=kind == "death"))
    if wipe is not None:
        at0, at1, n = wipe
        for t in range(at0, at1 + 1):
            s.append(f"<line x1=\"{fx(t):.1f}\" x2=\"{fx(t):.1f}\" y1=\"{top}\" y2=\"{top + plot_h}\" stroke=\"#ff5a5a\" stroke-opacity=\".5\" stroke-dasharray=\"4 3\"/>")
        s.append(lane_tile(fx((at0 + at1) / 2), top, count=n))
    if drag is not None:
        a, b = fx(drag[0]), fx(drag[1])
        s.append(f"<rect x=\"{a:.1f}\" y=\"{top}\" width=\"{b - a:.1f}\" height=\"{plot_h}\" fill=\"rgba(229,185,85,.13)\"/>")
        for x in (a, b):
            s.append(f"<line x1=\"{x:.1f}\" x2=\"{x:.1f}\" y1=\"{top}\" y2=\"{top + plot_h}\" stroke=\"#e5b955\"/>")
        label = f"{mmss(drag[0])} to {mmss(drag[1])} · {drag[1] - drag[0]} s"
        lw = 8 + len(label) * 6.6
        mid = (a + b) / 2
        s.append(f"<rect x=\"{mid - lw / 2:.1f}\" y=\"{top + 26}\" width=\"{lw:.1f}\" height=\"22\" rx=\"6\" fill=\"#0d0f15\" stroke=\"#e5b955\" stroke-opacity=\".6\"/>")
        s.append(f"<text x=\"{mid:.1f}\" y=\"{top + 41}\" font-size=\"12\" font-weight=\"600\" fill=\"#f2eee4\" text-anchor=\"middle\" style=\"font-variant-numeric:tabular-nums\">{label}</text>")
    if cursor is not None:
        x = fx(cursor)
        s.append(f"<line x1=\"{x:.1f}\" x2=\"{x:.1f}\" y1=\"{top}\" y2=\"{top + plot_h}\" stroke=\"#e9e4d8\" stroke-opacity=\".7\" stroke-dasharray=\"3 3\"/>")
        for name, col, vals, dead in series:
            if dead is not None and cursor > dead:
                continue
            s.append(f"<circle cx=\"{x:.1f}\" cy=\"{fy(vals[cursor]):.1f}\" r=\"3.5\" fill=\"{col}\" stroke=\"#07090d\" stroke-width=\"1.5\"/>")
        s.append(f"<rect x=\"{x - 20:.1f}\" y=\"{top + plot_h + 4}\" width=\"40\" height=\"17\" rx=\"4\" fill=\"#e9e4d8\"/>")
        s.append(f"<text x=\"{x:.1f}\" y=\"{top + plot_h + 16.5}\" font-size=\"11\" font-weight=\"700\" fill=\"#07090d\" text-anchor=\"middle\" style=\"font-variant-numeric:tabular-nums\">{mmss(cursor)}</text>")
    s.append("</svg>")
    return "".join(s)


RAID_RATE = [sum(DMG[n][t] for n in NAMES) for t in range(DURATION)]  # the brush's shape on every chart


def overview_svg(w, win, gutter=52, strip_h=36, phone=False):
    right = 12
    pw = w - gutter - right
    h = strip_h
    peak = max(RAID_RATE)
    fx = lambda t: gutter + t / DURATION * pw  # noqa: E731
    pts = [(fx(i), 4 + (h - 8) * (1 - v / peak)) for i, v in enumerate(smooth(RAID_RATE))]
    d = f"M{gutter},{h - 4} " + " ".join(f"L{x:.1f},{y:.1f}" for x, y in pts) + f" L{pts[-1][0]:.1f},{h - 4} Z"
    a, b = fx(win[0]), fx(win[1])
    s = [f"<svg width=\"{w}\" height=\"{h + 16}\" viewBox=\"0 0 {w} {h + 16}\" style=\"display:block;font-family:Barlow,sans-serif\">",
         f"<rect x=\"{gutter}\" y=\"0\" width=\"{pw}\" height=\"{h}\" rx=\"6\" fill=\"rgba(255,255,255,.025)\" stroke=\"#1c2230\"/>",
         f"<path d=\"{d}\" fill=\"rgba(154,148,132,.22)\" stroke=\"#9a9484\" stroke-opacity=\".5\"/>"]
    for _n, start, _e in PHASES:
        if start > 0:
            s.append(f"<line x1=\"{fx(start):.1f}\" x2=\"{fx(start):.1f}\" y1=\"2\" y2=\"{h - 2}\" stroke=\"#e5b955\" stroke-opacity=\".5\" stroke-dasharray=\"2 3\"/>")
    for dth in DEATHS:
        s.append(f"<line x1=\"{fx(dth['at']):.1f}\" x2=\"{fx(dth['at']):.1f}\" y1=\"2\" y2=\"{h - 2}\" stroke=\"#ff5a5a\" stroke-width=\"1.5\"/>")
    if win != FULL:
        s.append(f"<rect x=\"{gutter}\" y=\"0\" width=\"{a - gutter:.1f}\" height=\"{h}\" fill=\"rgba(7,9,13,.6)\"/>")
        s.append(f"<rect x=\"{b:.1f}\" y=\"0\" width=\"{gutter + pw - b:.1f}\" height=\"{h}\" fill=\"rgba(7,9,13,.6)\"/>")
    s.append(f"<rect x=\"{a:.1f}\" y=\".5\" width=\"{b - a:.1f}\" height=\"{h - 1}\" fill=\"rgba(229,185,85,.10)\" stroke=\"#e5b955\" stroke-opacity=\".8\"/>")
    hw = 12 if phone else 8
    for x in (a, b):
        s.append(f"<rect x=\"{x - hw / 2:.1f}\" y=\"{h / 2 - 12}\" width=\"{hw}\" height=\"24\" rx=\"3\" fill=\"#e5b955\"/>")
        s.append(f"<line x1=\"{x - 1:.1f}\" x2=\"{x - 1:.1f}\" y1=\"{h / 2 - 5}\" y2=\"{h / 2 + 5}\" stroke=\"#0b0e14\"/><line x1=\"{x + 1:.1f}\" x2=\"{x + 1:.1f}\" y1=\"{h / 2 - 5}\" y2=\"{h / 2 + 5}\" stroke=\"#0b0e14\"/>")
    for t, anchor in ((0, "start"), (DURATION, "end")):
        if win != FULL and min(abs(fx(t) - a), abs(fx(t) - b)) < 36:
            continue
        s.append(f"<text x=\"{fx(t):.1f}\" y=\"{h + 13}\" font-size=\"11\" fill=\"#9a9484\" text-anchor=\"{anchor}\" style=\"font-variant-numeric:tabular-nums\">{mmss(t)}</text>")
    if win != FULL:
        for t, x in ((win[0], a), (win[1], b)):
            s.append(f"<text x=\"{x:.1f}\" y=\"{h + 13}\" font-size=\"11\" font-weight=\"700\" fill=\"#e5b955\" text-anchor=\"middle\" style=\"font-variant-numeric:tabular-nums\">{mmss(t)}</text>")
    s.append("</svg>")
    return "".join(s)


def chart_markers(threat: bool):
    m = [(d["at"], "death", d) for d in DEATHS]
    if threat:
        m += [(at, "taunt", TAUNT) for at, _n in TAUNTS] + [(at, "target", n) for at, n in TARGET_CHANGES]
    else:
        m += [(at, "cast", src) for at, _n, src, _s in CASTS]
    return m


def death_line(d) -> str:
    over = f" ({d['overkill']:,} overkill)" if d.get("overkill") else ""
    return f"{d['name']} died · {d['spell']} from {d['source']} · {d['amount']:,}{over}"


def events_near(t: int, threat: bool):
    """Every marker within ±2 s of the hovered second, in time order (spec §B5.7)."""
    ev = []
    for d in DEATHS:
        if abs(d["at"] - t) <= 2:
            ev.append((d["at"], "death", death_line(d)))
    if threat:
        ev += [(at, "taunt", f"Taunt · {n} on Ragnaros") for at, n in TAUNTS if abs(at - t) <= 2]
        ev += [(at, "target", f"Ragnaros turned to {n}") for at, n in TARGET_CHANGES if abs(at - t) <= 2]
    else:
        ev += [(at, "cast", f"{sp} · {n}") for at, n, _src, sp in CASTS if abs(at - t) <= 2]
    order = {"death": 0, "target": 1, "taunt": 2, "cast": 3}
    return sorted(ev, key=lambda e: (e[0], order[e[1]]))


def event_html(at, kind, text):
    if kind == "death":
        g = f"<span class=\"g\" style=\"background:#5a1f1f;color:#ff5a5a\">{glyph('skull', 12)}</span>"
    elif kind == "target":
        g = f"<span class=\"g\" style=\"color:#e9e4d8\">{glyph('target', 14)}</span>"
    else:
        g = f"<img src=\"{TAUNT if kind == 'taunt' else DEATH_WISH}\" alt=\"\">"
    return f"<div class=\"ev{' death' if kind == 'death' else ''}\"><span class=\"t\">{mmss(at)}</span>{g}{text}</div>"


# ===================================================================== dropdown parts
CHECK = glyph("check", 16)


def field(label, size="", state="", lead="", ph=False, width=None, chevron=True):
    cls = f"sel {size} {state} {'ph' if ph else ''}"
    style = f" style=\"width:{width}px\"" if width else ""
    tail = "<span class=\"chv\"><span class=\"spin\"></span></span>" if state == "loading" else (
        f"<span class=\"chv\">{glyph('chev-down', 16 if size != 'd' else 14)}</span>" if chevron else "")
    return f"<span class=\"{cls}\"{style}>{lead}<span style=\"overflow:hidden;text-overflow:ellipsis\">{label}</span>{tail}</span>"


def opt(label, *, desc="", meta="", lead="", sel=False, act=False, dis=False, colour=None, box=None, you=False):
    cls = "opt" + (" ic" if lead else "") + (" act" if act else "") + (" dis" if dis else "")
    name = f"<span style=\"color:{colour}\">{label}</span>" if colour else f"<span>{label}</span>"
    if you:
        name += "<span class=\"you\">You</span>"
    body = f"<span class=\"t\"><span style=\"display:flex;align-items:center\">{name}</span>{f'<span class=\"ds\">{desc}</span>' if desc else ''}</span>"
    m = f"<span class=\"meta num\">{meta}</span>" if meta else ""
    ck = f"<span class=\"ck\">{CHECK}</span>" if sel else ""
    bx = "" if box is None else f"<span class=\"box{' on' if box else ''}\">{glyph('check', 12) if box else ''}</span>"
    return f"<div class=\"{cls}\" role=\"option\">{bx}{lead}{body}{m}{ck}</div>"


def listbox(inner, *, top=None, bottom=None, left=0, right=None, width=None, search="", fade=False, dense=False, foot=""):
    pos = f"top:{top}px" if top is not None else f"bottom:{bottom}px"
    pos += f";right:{right}px" if right is not None else f";left:{left}px"
    style = pos + (f";width:{width}px;min-width:0" if width else "")
    fd = "<div class=\"fade\"></div>" if fade else ""
    return f"<div class=\"lb{' dense' if dense else ''}\" role=\"listbox\" style=\"{style}\">{search}<div class=\"scroll\">{inner}{fd}</div>{foot}</div>"


def labelled(label, inner):
    return f"<div style=\"display:flex;flex-direction:column;gap:6px\"><span class=\"ql\">{label}</span>{inner}</div>"


def search_row(text, n):
    return ("<div class=\"lbsearch\"><span class=\"in\"><span style=\"color:var(--muted);display:flex\">" + glyph("search", 16) +
            f"</span>{text}<span class=\"caret\"></span></span><span class=\"n num\">{n}</span></div>")


FIGHT_STYLES = [
    ("Patchwerk", "One target, standing still for the whole fight, with an execute phase below 25% health."),
    ("Execute heavy", "One target, standing still, with a wider execute phase: below 35% health instead of 25%."),
    ("Light movement", "One target, forced out of melee range for 5 seconds every 45 seconds; no parse has measured the cost of that yet."),
    ("Heavy movement", "One target, forced out of melee range for 5 seconds every 20 seconds; no parse has measured the cost of that yet."),
    ("Cleave, 2 targets", "Two targets, both present and taking damage for the whole fight, sharing the same 25% execute threshold as Patchwerk."),
    ("Cleave, 3 targets", "Three targets, all present and taking damage for the whole fight, sharing the same 25% execute threshold as Patchwerk."),
]


def board(body, label, width=None):
    style = f"width:{width}px" if width else ""
    title = "Fix round 1" if width and width < 400 else "Controls and charts · 2026-10-30 · fix round 1"
    return head("B", CSS) + f"<div class=\"board\" style=\"{style}\"><div class=\"dirbar\"><b>{title}</b><span>{label}</span></div>{body}</div></body></html>"


def write(name, html):
    (RENDERS / f"{name}.html").write_text(html)


def page(inner, pad="32px 48px 48px"):
    return f"<div style=\"max-width:1344px;margin:0 auto;padding:{pad};width:100%\">{inner}</div>"


def you_lead():
    return f"<span style=\"width:20px;display:flex;justify-content:center;color:var(--muted)\">{glyph('user', 16)}</span>"


def dd_plain():
    rows = "".join(opt(n, desc=d, sel=(i == 0), act=(i == 1)) for i, (n, d) in enumerate(FIGHT_STYLES))
    return labelled("Fight style", f"<div style=\"position:relative;width:260px\">{field('Patchwerk', 'c', 'open', width=260)}"
                    f"{listbox(rows, top=50, width=400, fade=True)}</div>")


def dd_icons():
    by_role = {"tank": "Tanks", "healer": "Healers", "dps": "Damage"}
    inner = opt("All friendlies", lead=you_lead()) + opt("All enemies", lead=you_lead())
    for role in ("tank", "healer", "dps"):
        inner += f"<div class=\"grp rule\">{by_role[role]}</div>"
        pool = ["Garrosk", "Velyn", "Morrigan", "Ysolde", "Ashka", "Kael", VIEWER] if role == "dps" else [n for n in NAMES if ROLE[n] == role][:6]
        for n in pool:
            inner += opt(n, lead=crest(CLS[n], 20), colour=colour(n), sel=(n == "Morrigan"), act=(n == "Ysolde"), you=(n == VIEWER))
    scrolled = f"<div style=\"margin-top:-400px\">{inner}</div><div class=\"fade\" style=\"top:0;bottom:auto;transform:scaleY(-1)\"></div>"
    closed = field(f"<span style=\"color:{colour('Morrigan')}\">Morrigan</span>", "d", "open", lead=crest("mage", 20), width=200)
    return (f"<div style=\"display:flex;align-items:center;gap:8px\"><span class=\"ql\">Source</span>"
            f"<div style=\"position:relative;width:200px\">{closed}{listbox(scrolled, search=search_row('', '40 players'), top=42, width=260, fade=True, dense=True)}</div></div>"
            "<p class=\"lbl\" style=\"margin-top:430px\">Dense list (13 px rows) under a dense field. 40 players, so it is searchable; focus is in the filter. "
            "Opened scrolled to the selection (Morrigan). Ysolde is the keyboard-active row.</p>")


def dd_search():
    mk = lambda s: s.replace("Black", "<mark>Black</mark>")  # noqa: E731
    inner = ("<div class=\"grp\">Raids</div>" + opt(mk("Blackwing Lair"), act=True) +
             "<div class=\"grp rule\">Dungeons</div>" + opt(mk("Blackrock Depths")) + opt("Lower " + mk("Blackrock Spire")) + opt("Upper " + mk("Blackrock Spire")))
    return labelled("Source", f"<div style=\"position:relative;width:260px\">{field('Anywhere', 'c', 'open', width=260)}"
                    f"{listbox(search_row('black', '4 matches') + inner, top=50, width=300)}</div>"
                    "<p class=\"lbl\" style=\"margin-top:300px\">Focus is in the filter input (the combobox, aria-activedescendant = Blackwing Lair). "
                    "The trigger keeps aria-expanded and aria-controls.</p>")


def dd_search_empty():
    return labelled("Source", f"<div style=\"position:relative;width:260px\">{field('Anywhere', 'c', 'open', width=260)}"
                    f"{listbox(search_row('naxx', '0 matches') + '<div class=lbempty>Nothing matches “naxx”.</div>', top=50, width=300)}</div>")


def build_dropdowns_open():
    cols = [("Plain, with descriptions · sim Fight style (44 px field)", dd_plain()),
            ("With icons, searchable · report Source (dense 36 px field, 13 px rows)", dd_icons()),
            ("Searchable · sim Find an item → Source", dd_search())]
    inner = "".join(f"<div style=\"flex:1;min-width:0\"><p class=\"cap\">{c}</p><div class=\"card\" style=\"padding:20px;min-height:610px;overflow:visible\">{b}</div></div>" for c, b in cols)
    flip = (f"<div style=\"position:absolute;left:340px;top:150px\">{field('Every enemy', 'd', 'open', width=180)}"
            f"{listbox(opt('Every enemy', sel=True) + opt('Ragnaros', act=True) + opt('Son of Flame', meta='×8'), bottom=40, width=220, dense=True)}</div>")
    right_slot = (f"<div style=\"position:absolute;right:20px;top:24px\">{field('Top 8 of 40', 'd', 'open', width=170)}"
                  f"{listbox(opt('Raid median', box=True) + opt('Garrosk', box=True, lead=crest('warrior', 20), colour=colour('Garrosk')), top=42, right=0, width=240, dense=True)}</div>")
    second = (f"<div style=\"display:flex;gap:24px;margin-top:28px\"><div style=\"flex:1\"><p class=\"cap\">Searchable, no match</p><div class=\"card\" style=\"padding:20px;height:240px\">{dd_search_empty()}</div></div>"
              "<div style=\"flex:2\"><p class=\"cap\">Placement: top layer (popover), 4 px from the field, field width to 420 px, max 360 px tall. "
              "Flips above when under 240 px fits below. A field in a right slot opens bottom-end (right edges aligned).</p>"
              "<div class=\"card\" style=\"padding:20px;height:240px;position:relative;overflow:hidden\"><div style=\"position:absolute;left:20px;right:20px;bottom:12px;height:36px;border:1px dashed #262e40;border-radius:8px;display:flex;align-items:center;padding:0 12px\" class=\"lbl\">viewport bottom edge, 40 px below the field</div>"
              f"{flip}<div class=\"lbl\" style=\"position:absolute;left:30px;top:60px;width:290px\">flipped: the list opens above the field, same 4 px gap →</div>{right_slot}</div></div></div>")
    write("dropdowns-open-1440", board(page(f"<div style=\"display:flex;gap:24px\">{inner}</div>{second}"), "the shared dropdown, open, fine pointer"))


def build_dropdown_states():
    states = [("Default", "", False), ("Hover", "hover", False), ("Focus-visible", "focus", False), ("Open", "open", False),
              ("Active (pressed)", "act", False), ("Disabled", "dis", False), ("Error", "err", False), ("Loading", "loading", False),
              ("Empty", "dis", True), ("Placeholder", "", True)]

    def row(size, w):
        cells = ""
        for name, st, ph in states:
            label = {"Loading": "Loading…", "Empty": "No enemies", "Placeholder": "Choose a spec"}.get(name, "Patchwerk")
            extra = "<div class=\"errline\">Pick a fight style.</div>" if name == "Error" else ""
            f = field(label, size, st, ph=ph or name == "Loading", width=w, chevron=name != "Empty")
            cells += f"<div style=\"display:flex;flex-direction:column;gap:8px\"><span class=\"lbl\">{name}</span>{f}{extra}</div>"
        return cells
    sizes = [("Default · 52 px · forms", "", 200), ("Compact · 44 px · filter bars, card heads, every field on a coarse pointer", "c", 200),
             ("Dense · 36 px · report and sim tool bars, fine pointer from 768 px", "d", 170)]
    grid = "".join(f"<p class=\"cap\" style=\"margin-top:22px\">{c}</p><div style=\"display:grid;grid-template-columns:repeat(5,auto);gap:20px 24px;justify-content:start\">{row(s, w)}</div>" for c, s, w in sizes)
    note = ("<p class=\"cap\" style=\"margin-top:22px\">Hover, open and pressed are the /account trigger's (.csel-btn): hover border gold 45% and a --color-strong chevron; "
            "open a gold border, the --glow-gold glow and the chevron turned in --color-strong; pressed a gold border on #101624. The radius stays 8 px: a field is an input.</p>")
    with_icon = ("<p class=\"cap\" style=\"margin-top:28px\">The closed field with a lead icon (20 px crest, spec or spell icon; class-coloured name)</p><div style=\"display:flex;gap:24px\">"
                 + field(f"<span style=\"color:{colour('Morrigan')}\">Morrigan</span>", "c", "", lead=crest("mage", 20), width=220)
                 + field(f"<span style=\"color:{colour('Brokk')}\">Brokk</span>", "d", "hover", lead=crest("warrior", 20), width=180) + "</div>")
    anatomy = ("<p class=\"cap\" style=\"margin-top:28px\">Option rows: plain 36 px (dense 32) · icon 40 (dense 36) · description 52+ · group head 28 · selected check · active · disabled · meta · multi-select box</p>"
               "<div style=\"display:flex;gap:24px;align-items:flex-start\"><div class=\"lb\" style=\"position:relative;width:420px;min-width:0\"><div class=\"scroll\">"
               + opt("Every enemy", sel=True) + opt("Ragnaros", act=True) + opt("Son of Flame", meta="×8")
               + "<div class=\"grp rule\">Damage</div>" + opt("Garrosk", lead=crest("warrior", 20), colour=colour("Garrosk"))
               + opt("Execute heavy", desc="One target, standing still, with a wider execute phase: below 35% health instead of 25%.")
               + opt("Dungeon pull", dis=True) + "</div></div>"
               "<div class=\"lb dense\" style=\"position:relative;width:300px;min-width:0\"><div class=\"scroll\">"
               + opt("Raid median", box=True) + opt(VIEWER, box=True, lead=crest("shaman", 20), colour=colour(VIEWER), you=True)
               + opt("Kael", box=False, lead=crest("hunter", 20), colour=colour("Kael"), act=True) + "</div></div></div>")
    write("dropdown-states-1440", board(page(grid + note + with_icon + anatomy), "the closed field: every state, three sizes; the option rows"))


# ===================================================================== the report around the charts
TABS = ["Damage done", "Healing", "Damage taken", "Deaths", "Casts", "Buffs", "Threat", "More"]


def report_head(tab, extra=""):
    tabs = "".join(f"<span class=\"{'on' if t == tab else ''}\">{t}{' ▾' if t == 'More' else ''}</span>" for t in TABS)
    return (f"<span class=\"eyebrow\">Report</span><div class=\"h1\" style=\"font-size:44px;margin-top:8px\">Molten Core</div>"
            f"<p style=\"font-size:13px;color:var(--muted);margin-top:8px\">Ragnaros · Kill · {mmss(DURATION)} · 40 players · mock data</p>"
            f"<div style=\"display:flex;align-items:center;gap:16px;margin-top:20px\"><div class=\"tabs\" style=\"flex:1\">{tabs}</div>{extra}"
            f"<span class=\"ql\">Source</span>{field('All friendlies', 'd', width=170)}</div>")


def chip(name, value, *, off=False, hov=False, only=False, pinned=False):
    col = colour(name)
    cls = "chip" + (" off" if off else "") + (" hov" if hov else "")
    o = "<span class=\"only\">Only</span>" if only else ""
    pin = f"<span class=\"pin\" title=\"Always drawn\">{glyph('pin', 12)}</span>" if pinned else ""
    you = "<span class=\"you\">YOU</span>" if name == VIEWER else ""
    return f"<span class=\"{cls}\" style=\"color:{col}\"><span class=\"sw\" style=\"background:{col}\"></span>{pin}{name}{you}<span class=\"v num\">{value}</span>{o}</span>"


def median_chip(value):
    return f"<span class=\"chip med\" style=\"color:#e9e4d8\"><span class=\"sw\"></span>Raid median<span class=\"v num\">{value}</span></span>"


THREAT_WIN = (90, 160)
CAVEAT = ("An estimate: 1 threat per point of damage and 0.5 per point of healing, split across the enemies in combat. "
          "No stance, form, talent or taunt modifiers yet. <a class=\"ilink\">What it leaves out</a>")
CAVEAT_PHONE = "Damage and healing only. <a class=\"ilink\">What it leaves out</a>"


def threat_tooltip(t, drawn):
    alive = sorted([n for n in drawn if n not in DEAD_AT or DEAD_AT[n] >= t], key=lambda n: -standing(n, t))
    dead = [n for n in drawn if n in DEAD_AT and DEAD_AT[n] < t]
    rows = "".join(f"<div class=\"r\" style=\"color:{colour(n)}\"><span class=\"sw\" style=\"background:{colour(n)}\"></span>{n}"
                   f"<span class=\"v num\">{standing(n, t):,.0f}</span><span class=\"p num\">{share(n, t):.1f}%</span></div>" for n in alive[:12])
    rows += "".join(f"<div class=\"r dead\"><span class=\"sw\" style=\"background:{colour(n)};opacity:.4\"></span>{n}<span class=\"v num\" style=\"color:var(--muted);font-weight:500\">Died {mmss(DEAD_AT[n])}</span><span class=\"p\"></span></div>" for n in dead)
    ev = "".join(event_html(*e) for e in events_near(t, True))
    return (f"<div class=\"th\"><b class=\"num\">{mmss(t)}</b><span>Ragnaros phase · estimated threat</span></div>{ev}"
            f"<div class=\"ch\"><span>Estimated threat</span><span>Share</span></div>{rows}")


def threat_frame(width, *, win=THREAT_WIN, cursor=127):
    inner_w = width - 2
    drawn, pinned = threat_drawn(win)
    gutter = 52
    plot_h = 300
    top_val = max(standing(n, win[1] - 1) for n in drawn)
    ymax, ticks = nice_max(top_val)
    series = [(n, colour(n), THREAT[n], DEAD_AT.get(n)) for n in drawn]
    svg = chart_svg(inner_w, plot_h, win, series, yticks=ticks, ymax=ymax, markers=chart_markers(True), cursor=cursor, gutter=gutter)
    ov = overview_svg(inner_w, win, gutter=gutter)
    chips = "".join(chip(n, kfmt(standing(n, win[1] - 1)), pinned=n in pinned) for n in drawn)
    chips += f"<span class=\"ghost\" style=\"height:28px\">+{len(NAMES) - len(drawn)} more</span>"
    right = (f"<span class=\"ql\">Players</span>{field(f'Top 8 of {len(NAMES)}', 'd', width=150)}"
             f"<span class=\"ql\">Enemy</span>{field('Ragnaros', 'd', width=150)}<span class=\"ghost\">{glyph('x', 14)}Whole fight</span>")
    head_ = (f"<div class=\"fhead\"><span class=\"stitle\">Estimated threat on Ragnaros</span><span class=\"meta\">Cumulative · {mmss(win[0])} to {mmss(win[1])}</span>"
             f"<span class=\"right\">{right}</span></div>")
    pw = inner_w - gutter - 12
    gx = gutter + (cursor - win[0]) / (win[1] - win[0]) * pw
    left_side = (cursor - win[0]) / (win[1] - win[0]) > 0.45
    pos = f"left:{gx - 16 - 300:.0f}px" if left_side else f"left:{gx + 16:.0f}px"
    tip = f"<div class=\"tip\" style=\"{pos};top:{LANE_H}px;width:300px\">{threat_tooltip(cursor, drawn)}</div>"
    foot = f"<div style=\"border-top:1px solid var(--line-soft);padding:8px 16px;font-size:12px;color:var(--muted)\">{CAVEAT}</div>"
    return (f"<div class=\"frame\" style=\"width:{width}px\">{head_}<div class=\"legend\">{chips}</div>"
            f"<div style=\"position:relative\">{svg}{tip}</div><div style=\"padding:6px 0 10px\">{ov}</div>{foot}</div>"), drawn


def threat_table(win=THREAT_WIN, rows_shown=6):
    t = win[1] - 1
    by = sorted(NAMES, key=lambda n: -standing(n, t))
    top = standing(by[0], t)
    rows = "".join(f"<tr><td style=\"color:{colour(n)};font-weight:600\"><span style=\"display:inline-flex;align-items:center;gap:10px\">{crest(CLS[n], 24)}{n}</span></td>"
                   f"<td><div class=\"bar\"><i style=\"width:{standing(n, t) / top * 100:.0f}%;background:{colour(n)}\"></i></div></td>"
                   f"<td class=\"r num\">{standing(n, t):,.0f}</td><td class=\"r num\" style=\"color:var(--muted)\">{share(n, t):.1f}%</td></tr>" for n in by[:rows_shown])
    return ("<div class=\"card\" style=\"margin-top:22px\"><div class=\"fhead\"><span class=\"stitle\">Estimated threat</span>"
            f"<span class=\"meta\">On Ragnaros · standing at {mmss(win[1])} · {len(NAMES)} players</span></div>"
            f"<table class=\"tbl\"><tr><th>Player</th><th>Estimated threat</th><th class=\"r\">Standing</th><th class=\"r\">Share of estimated threat</th></tr>{rows}</table>"
            f"<p class=\"lbl\" style=\"padding:8px 12px 12px\">{len(NAMES) - rows_shown} more rows. Leader {by[0]} {standing(by[0], t):,.0f} = the legend's first chip; shares are of all {len(NAMES)} players' estimated threat.</p></div>")


def build_threat(name):
    frame, _drawn = threat_frame(1248)
    body = report_head("Threat") + f"<div style=\"margin-top:16px\">{frame}</div>" + threat_table()
    write(name, board(page(body), "threat: estimated, the tank pinned, target changes, events within ±2 s with the killing blow, zoomed 1:30 to 2:40"))


def damage_frame(width, *, win=FULL, drag=(48, 80), dim="Morrigan", menu=""):
    inner_w = width - 2
    drawn, pinned = damage_drawn(win)
    series = [(n, colour(n), SMOOTH[n], DEAD_AT.get(n)) for n in drawn]
    peak = max(max(SMOOTH[n][win[0]:win[1]]) for n in drawn)
    ymax, ticks = nice_max(peak)
    svg = chart_svg(inner_w, 220, win, series, yticks=ticks, ymax=ymax, markers=chart_markers(False), drag=drag, dim=dim, median=MEDIAN)
    ov = overview_svg(inner_w, win)
    med = round(statistics.median(window_dps(n, win) for n in NAMES if ROLE[n] == "dps"))
    chips = median_chip(f"{med:,}") + "".join(chip(n, f"{window_dps(n, win):,}", hov=(n == dim), only=(n == dim), pinned=n in pinned) for n in drawn)
    chips += f"<span class=\"ghost\" style=\"height:28px\">+{len(NAMES) - len(drawn)} more</span>"
    right = (f"<span class=\"ql\">Players</span>{field(f'Top 8 of {len(NAMES)}', 'd', 'open' if menu else '', width=150)}"
             f"<span class=\"ql\">Window</span>{field('Whole fight', 'd', width=150)}<span class=\"sw2\"><span class=\"tr\"></span>Log scale</span>")
    raid = sum(window_dps(n, win) for n in NAMES)
    head_ = (f"<div class=\"fhead\" style=\"position:relative\"><span class=\"stitle\">Damage done</span><span class=\"meta num\">Raid {raid:,} DPS · on Ragnaros · 5 s average</span>"
             f"<span class=\"right\">{right}</span>{menu}</div>")
    return (f"<div class=\"frame\" style=\"width:{width}px;overflow:visible\">{head_}<div class=\"legend\">{chips}</div>"
            f"<div style=\"position:relative\">{svg}</div><div style=\"padding:6px 0 16px\">{ov}</div></div>"), raid


def damage_table(win=FULL, rows_shown=4):
    by = sorted(NAMES, key=lambda n: -window_dps(n, win))
    top = window_dps(by[0], win)
    rows = "".join(f"<tr><td style=\"color:{colour(n)};font-weight:600\"><span style=\"display:inline-flex;align-items:center;gap:10px\">{crest(CLS[n], 24)}{n}</span></td>"
                   f"<td><div class=\"bar\"><i style=\"width:{window_dps(n, win) / top * 100:.0f}%;background:{colour(n)}\"></i></div></td>"
                   f"<td class=\"r num\">{sum(DMG[n][win[0]:win[1]]):,.0f}</td><td class=\"r num\">{window_dps(n, win):,}</td></tr>" for n in by[:rows_shown])
    return ("<div class=\"card\" style=\"margin-top:22px\"><div class=\"fhead\"><span class=\"stitle\">Damage done</span>"
            f"<span class=\"meta\">On Ragnaros · {len(NAMES)} players · {len(DEATHS)} death · {mmss(win[1] - win[0])}</span></div>"
            f"<table class=\"tbl\"><tr><th>Player</th><th>Amount</th><th class=\"r\">Total</th><th class=\"r\">DPS</th></tr>{rows}</table></div>")


def build_damage():
    frame, raid = damage_frame(1248)
    total = sum(window_dps(n, FULL) for n in NAMES)
    ann = (f"<div class=\"lbl\" style=\"margin-top:10px\">Reconciles: the {len(NAMES)} players' DPS figures sum to {total:,} = the head's Raid {raid:,} DPS. "
           f"Submerge (3:00 to 4:30): Ragnaros cannot be damaged, the lines sit at 0. Drawn: top 8 of {len(NAMES)} plus {VIEWER} (you, pinned) and the raid median (dashed). "
           "Legend hover on Morrigan with Only; a drag in progress 0:48 to 1:20.</div>")
    write("damage-timeline-1440", board(page(report_head("Damage done") + f"<div style=\"margin-top:16px\">{frame}</div>{ann}{damage_table()}"),
                                        "damage done on Ragnaros: top 8 of 40, you pinned, raid median, cast and death lane, drag-to-zoom"))


def choose_players_menu():
    drawn, _p = damage_drawn(FULL)
    rows = ("<div class=\"grp\">Compare</div>" + opt("Raid median", box=True, desc=f"Median DPS of the {sum(1 for n in NAMES if ROLE[n] == 'dps')} damage dealers")
            + opt("Shaman median", box=False, desc=f"Median DPS of the {sum(1 for n in NAMES if CLS[n] == 'shaman' and ROLE[n] == 'dps')} damage-dealing shamans (your class)"))
    for role, label in (("tank", "Tanks"), ("dps", "Damage")):
        rows += f"<div class=\"grp rule\">{label}</div>"
        names = sorted([n for n in NAMES if ROLE[n] == role], key=lambda n: -window_dps(n, FULL))
        for n in names[:7 if role == "dps" else 2]:
            rows += opt(n, box=n in drawn, lead=crest(CLS[n], 20), colour=colour(n), meta=f"{window_dps(n, FULL):,}", you=n == VIEWER, act=n == "Brokk")
    foot = ("<div style=\"display:flex;align-items:center;gap:8px;border-top:1px solid var(--line-soft);padding:6px 8px\">"
            f"<span class=\"ghost\" style=\"height:32px\">Top 8</span><span class=\"lbl\" style=\"margin-left:auto;padding-right:6px\">{len(drawn)} drawn</span></div>")
    return listbox(search_row("", f"{len(NAMES)} players") + rows, top=46, right=340, width=320, fade=True, dense=True, foot=foot)


def build_choose_players():
    frame, _raid = damage_frame(1248, drag=None, dim=None, menu=choose_players_menu())
    note = ("<p class=\"lbl\" style=\"margin-top:10px\">Players: a searchable multi-select (the shared listbox with checkboxes). Space or click toggles a row and the list stays open; "
            "Enter or Escape closes it. The Compare group holds the raid median (on by default) and your class's median. The DPS meta is the window's. "
            "Top 8 resets the default. The list opens bottom-end because the field is in the head's right slot.</p>")
    write("choose-players-1440", board(page(report_head("Damage done") + f"<div style=\"margin-top:16px\">{frame}</div>{note}"), "Choose players: searchable multi-select, open"))


# ===================================================================== states
SMALL_W = 612
SMALL_PLOT = 200
# One reserved height for every state (spec §B11.2): head 48 + rule 1 + legend 42 + lane 26 + plot + axis 22 + brush 6+52+16 + border 2.
SMALL_H = 48 + 1 + 42 + LANE_H + SMALL_PLOT + AXIS_H + 6 + 52 + 16 + 2


def small_frame(title, meta, legend, body, brush, *, right=""):
    return (f"<div class=\"frame\" style=\"height:{SMALL_H}px;overflow:visible\"><div class=\"fhead\"><span class=\"stitle\">{title}</span><span class=\"meta\">{meta}</span>"
            f"<span class=\"right\">{right}</span></div><div class=\"legend\">{legend}</div>{body}<div style=\"padding:6px 0 16px\">{brush}</div></div>")


def head_controls(sel="Whole fight", menu_open=False):
    return (f"<span class=\"ql\">Window</span>{field(sel, 'd', 'open' if menu_open else '', width=150)}<span class=\"sw2\"><span class=\"tr\"></span>Log scale</span>")


def build_chart_states():
    w = SMALL_W
    brush_full = overview_svg(w - 2, FULL)
    skel_chips = "".join(f"<span class=\"skel\" style=\"width:{x}px;height:20px;align-self:center\"></span>" for x in (96, 110, 88, 102, 92))
    skel_plot = f"<div style=\"position:relative;height:{LANE_H + SMALL_PLOT + AXIS_H}px\"><div class=\"skel\" style=\"position:absolute;left:52px;right:12px;top:{LANE_H}px;height:{SMALL_PLOT}px\"></div></div>"
    loading = small_frame("Damage done", "Loading the chart", skel_chips, skel_plot,
                          f"<div style=\"margin:0 12px 0 52px;height:52px\"><div class=\"skel\" style=\"height:36px\"></div></div>", right=head_controls())
    empty_win = (200, 250)  # inside Submerge: Ragnaros takes no damage
    axes = chart_svg(w - 2, SMALL_PLOT, empty_win, [], yticks=[0, 400, 800, 1200], ymax=1200)
    empty_body = (f"<div style=\"position:relative\">{axes}<div style=\"position:absolute;left:52px;right:12px;top:{LANE_H}px;height:{SMALL_PLOT}px;display:grid;place-items:center;font-size:14px;color:var(--muted)\">No damage in this window.</div></div>")
    empty = small_frame("Damage done", "Submerge · no damage", "", empty_body, overview_svg(w - 2, empty_win), right=head_controls("3:20 to 4:10"))
    err_body = f"<div style=\"height:{LANE_H + SMALL_PLOT + AXIS_H}px;display:flex;flex-direction:column;align-items:center;justify-content:center;gap:14px\"><span style=\"font-size:14px;color:var(--muted)\">The chart didn't load.</span><span class=\"btn btn-secondary sm\">Try again</span></div>"
    error = small_frame("Damage done", "", "", err_body, f"<div style=\"height:52px\"></div>", right=head_controls())
    # a menu over the plot: bottom-end under the right-slot field, top layer, the chart stands down
    inner = opt("Whole fight", sel=True) + opt("First 30s") + opt("Last 30s", act=True) + "".join(
        opt(f"{p} · {mmss(a)} to {mmss(b)}") for p, a, b in PHASES) + opt("20s before Brokk died · 2:06")
    drawn, pinned = damage_drawn(FULL)
    plot = chart_svg(w - 2, SMALL_PLOT, FULL, [(n, colour(n), SMOOTH[n], DEAD_AT.get(n)) for n in drawn], yticks=[0, 400, 800, 1200], ymax=1200,
                     markers=chart_markers(False), median=MEDIAN)
    menu = listbox(inner, top=46, right=126, width=260, dense=True)
    over = (f"<div class=\"frame\" style=\"height:{SMALL_H}px;overflow:visible\"><div class=\"fhead\" style=\"position:relative\"><span class=\"stitle\">Damage done</span><span class=\"meta\">Window menu open</span>"
            f"<span class=\"right\">{head_controls(menu_open=True)}</span>{menu}</div><div class=\"legend\">{median_chip('')}{''.join(chip(n, '', pinned=n in pinned) for n in drawn[:4])}</div>"
            f"<div style=\"position:relative\">{plot}</div><div style=\"padding:6px 0 16px\">{brush_full}</div></div>")
    # a wipe: 14 deaths in 5 s collapse into one counted circle; its tooltip lists them
    wipe_names = [n for n in NAMES if ROLE[n] != "tank"][:14]
    wipe_series = []
    for i, n in enumerate(drawn[:6]):
        vals = [v if t < 302 + (i % 5) else 0 for t, v in enumerate(SMOOTH[n])]
        wipe_series.append((n, colour(n), vals, None))
    wplot = chart_svg(w - 2, SMALL_PLOT, (270, 320), wipe_series, yticks=[0, 400, 800, 1200], ymax=1200, wipe=(302, 306, 14))
    wtip = ("<div class=\"tip\" style=\"left:250px;top:30px;width:300px\"><div class=\"th\"><b class=\"num\">5:02 to 5:06</b><span>14 deaths</span></div>"
            + "".join(f"<div class=\"ev death\"><span class=\"t\">{mmss(302 + i * 4 // 14)}</span><span class=\"g\" style=\"background:#5a1f1f;color:#ff5a5a\">{glyph('skull', 12)}</span>{n} · Magma Blast</div>" for i, n in enumerate(wipe_names[:6]))
            + "<div class=\"more\">+8 more · open on the Deaths tab</div></div>")
    wipe = small_frame("Damage done", "A mock wipe pull", "", f"<div style=\"position:relative\">{wplot}{wtip}</div>", overview_svg(w - 2, (270, 320)), right=head_controls("4:30 to 5:20"))
    cap = lambda t: f"<p class=\"cap\">{t}</p>"  # noqa: E731
    grid = (f"<div style=\"display:grid;grid-template-columns:{w}px {w}px;gap:28px 24px\">"
            f"<div>{cap(f'Loading: every state is the same {SMALL_H} px; skeleton chips, plot and brush; no spinner')}{loading}</div>"
            f"<div>{cap('Empty: axes, legend row and brush stay; one centred line; the head says why')}{empty}</div>"
            f"<div>{cap('Error: the chart chunk failed; the tables below still work')}{error}</div>"
            f"<div>{cap('A menu over the plot: bottom-end under the right-slot field, top layer; no tooltip or crosshair while it is open')}{over}</div>"
            f"<div>{cap('A wipe: 14 deaths in 5 s collapse into one counted marker; its tooltip lists them and links to Deaths')}{wipe}</div></div>")
    write("chart-states-1440", board(page(grid), "chart states: loading, empty, error, a menu over the plot, a wipe"))


# ===================================================================== phone, 360 x 640
PH_W = 360
PH_FRAME = 324  # 360 minus the 18 px gutters
# The fold stack after a tab change on a phone (spec §B9.1), each measured or fixed:
PH = {"scroll margin": 8, "Source row": 44, "row gap": 8, "Table row": 44, "gap to chart": 12,
      "chart head": 48, "Enemy row": 52, "legend": 37, "lane": LANE_H, "plot": 160, "x axis": AXIS_H,
      "event line": 28, "brush": 62, "caveat": 28, "frame borders": 2}
PH_FRAME_H = sum(v for k, v in PH.items() if k not in ("scroll margin", "Source row", "row gap", "Table row", "gap to chart"))
PH_BOTTOM = sum(PH.values())
assert PH_BOTTOM <= 640, PH_BOTTOM


def phone_threat(cursor=127):
    win = THREAT_WIN
    drawn, pinned = threat_drawn(win)
    inner_w = PH_FRAME - 2
    top_val = max(standing(n, win[1] - 1) for n in drawn)
    ymax, ticks = nice_max(top_val)
    series = [(n, colour(n), THREAT[n], DEAD_AT.get(n)) for n in drawn]
    svg = chart_svg(inner_w, PH["plot"], win, series, yticks=ticks, ymax=ymax, markers=chart_markers(True), cursor=cursor, gutter=40)
    ov = overview_svg(inner_w, win, gutter=40, strip_h=32, phone=True)
    # The legend is one row: the reserved time slot, then chips in drawn order while they fit, then +N more.
    shown = drawn[:2]
    chips = "".join(chip(n, kfmt(standing(n, cursor))) for n in shown)
    legend = (f"<div class=\"legend\" style=\"min-height:0;height:37px;flex-wrap:nowrap;overflow:hidden;padding:4px 8px\">"
              f"{chips}<span class=\"chip more\" style=\"margin-left:auto\">+{len(drawn) - len(shown)}</span></div>")
    ev = events_near(cursor, True)
    first = ev[0]
    d0 = DEATHS[0]
    phone_text = f"{d0['name']} died · {d0['spell']} · {d0['amount']:,}" if first[1] == "death" else first[2]
    event = (f"<div style=\"display:flex;align-items:center;gap:8px;height:{PH['event line']}px;padding:0 12px;font-size:12px;color:#ffb4b0;white-space:nowrap;overflow:hidden\">"
             f"<span style=\"color:#ff5a5a;display:flex\">{glyph('skull', 14)}</span><span style=\"overflow:hidden;text-overflow:ellipsis\">{phone_text}</span>"
             f"<span style=\"margin-left:auto;color:var(--muted)\">+{len(ev) - 1}</span>{glyph('chev-right', 14)}</div>")
    head_ = f"<div class=\"fhead\" style=\"min-height:{PH['chart head']}px\"><span class=\"stitle\">Estimated threat</span><span class=\"meta\">{mmss(win[0])} to {mmss(win[1])}</span></div>"
    enemy = f"<div style=\"padding:0 12px 8px\">{field('Ragnaros', 'c', width=inner_w - 24)}</div>"
    foot = f"<div style=\"height:{PH['caveat']}px;display:flex;align-items:center;border-top:1px solid var(--line-soft);padding:0 12px;font-size:12px;color:var(--muted);white-space:nowrap\">{CAVEAT_PHONE}</div>"
    return (f"<div class=\"frame\" style=\"width:{PH_FRAME}px;height:{PH_FRAME_H}px;overflow:hidden\">{head_}{enemy}{legend}{svg}{event}"
            f"<div style=\"padding:6px 0 8px\">{ov}</div>{foot}</div>")


def build_phone():
    rows = (f"<div style=\"height:{PH['scroll margin']}px\"></div>"
            f"<div style=\"display:flex;align-items:center;gap:8px;height:44px\"><span class=\"ql\" style=\"width:48px\">Source</span>{field('All friendlies', 'c', width=PH_FRAME - 56)}</div>"
            f"<div style=\"height:{PH['row gap']}px\"></div>"
            f"<div style=\"display:flex;align-items:center;gap:8px;height:44px\"><span class=\"ql\" style=\"width:48px\">Table</span>{field('Threat', 'c', width=PH_FRAME - 56)}</div>"
            f"<div style=\"height:{PH['gap to chart']}px\"></div>{phone_threat()}"
            f"<div class=\"card\" style=\"margin-top:12px;padding:12px;font-size:13px;color:var(--muted)\">Estimated threat table: first row</div>")
    view = (f"<div style=\"position:relative;width:{PH_W}px;height:720px;overflow:hidden;padding:0 18px;background:#07090d\">{rows}"
            f"<div class=\"fold\" style=\"top:640px\"><span>640 px fold</span></div>"
            f"<div class=\"fold\" style=\"top:{PH_BOTTOM}px;border-top-color:#7fd48a\"><span style=\"color:#7fd48a;right:auto;left:6px;top:2px\">chart ends {PH_BOTTOM} px</span></div></div>")
    table = "".join(f"<tr><td style=\"padding:2px 0\">{k}</td><td style=\"text-align:right\">{v}</td></tr>" for k, v in PH.items())
    meas = (f"<div style=\"padding:12px 18px 24px\"><p class=\"lbl\">The viewport after a tab change (the page scrolls the Source row to the top). "
            f"Heights in px, the fold stack:</p><table class=\"lbl\" style=\"width:100%;margin-top:6px\">{table}"
            f"<tr><td style=\"padding-top:4px;color:#e5b955\">total</td><td style=\"text-align:right;color:#e5b955\">{PH_BOTTOM}</td></tr></table></div>")
    write("controls-charts-phone-360", board(view + meas, f"360 × 640 · {PH_BOTTOM} px", width=PH_W))


def build_legend_sheet():
    drawn, pinned = threat_drawn(THREAT_WIN)
    t = THREAT_WIN[1] - 1
    rows = (f"<div class=\"srow\" style=\"border-top:0\"><span class=\"in\" style=\"flex:1;display:flex;align-items:center;gap:8px;height:40px;padding:0 10px;border-radius:6px;border:1px solid var(--line);background:#07090d;color:var(--muted);font-size:14px;font-weight:500\">{glyph('search', 16)}Filter players</span></div>")
    rows += "<div class=\"grp\" style=\"padding:12px 18px 4px\">Drawn</div>"
    for n in drawn:
        on = True
        pin = f"<span style=\"color:var(--muted);display:flex\">{glyph('pin', 14)}</span>" if n in pinned else ""
        rows += (f"<div class=\"srow\">{crest(CLS[n], 24)}<span style=\"color:{colour(n)}\">{n}</span>{pin}<span class=\"v num\">{kfmt(standing(n, t))}</span>"
                 f"<span class=\"switch{' on' if on else ''}\"></span></div>")
    rows += "<div class=\"grp\" style=\"padding:12px 18px 4px\">Not drawn</div>"
    rest = [n for n in sorted(NAMES, key=lambda n: -standing(n, t)) if n not in drawn][:3]
    for n in rest:
        rows += f"<div class=\"srow\">{crest(CLS[n], 24)}<span style=\"color:{colour(n)}\">{n}</span><span class=\"v num\">{kfmt(standing(n, t))}</span><span class=\"switch\"></span></div>"
    sheet = (f"<div class=\"sheet\" style=\"position:absolute;left:0;right:0;bottom:0;height:520px;overflow:hidden\">"
             "<div style=\"display:flex;justify-content:center;padding:8px 0 4px\"><span style=\"width:36px;height:4px;border-radius:2px;background:#3a3326\"></span></div>"
             f"<div style=\"display:flex;align-items:center;gap:12px;padding:6px 18px 10px\"><span class=\"stitle\">Players</span><span class=\"meta\" style=\"font-size:13px;color:var(--muted)\">{len(drawn)} of {len(NAMES)} drawn</span>"
             "<span style=\"margin-left:auto;display:flex;gap:4px\"><span class=\"ghost\">Top 8</span><span class=\"ghost\">Done</span></span></div>"
             f"{rows}</div>")
    view = (f"<div style=\"position:relative;width:{PH_W}px;height:640px;overflow:hidden;background:#07090d\"><div style=\"padding:0 18px;opacity:.35\">"
            f"<div style=\"height:116px\"></div>{phone_threat()}</div><div style=\"position:absolute;inset:0;background:rgba(0,0,0,.55)\"></div>{sheet}</div>")
    note = ("<div style=\"padding:12px 18px 24px\"><p class=\"lbl\">Tapping +N in the legend, or the head's Players field on a coarse pointer, opens this sheet "
            "(the account selector's phone sheet recipe). Each row: crest, name, the value at the window's end, a Switch. Pinned rows "
            "(the enemy's targets and taunters in the window) carry the pin and cannot be switched off. Done or a swipe down closes it.</p></div>")
    write("legend-sheet-phone-360", board(view + note, "360 · legend sheet", width=PH_W))


def main():
    RENDERS.mkdir(exist_ok=True)
    build_dropdowns_open()
    build_dropdown_states()
    build_threat("threat-chart-1440")
    build_threat("threat-chart-2000")
    build_damage()
    build_choose_players()
    build_chart_states()
    build_phone()
    build_legend_sheet()
    print("phone fold stack", PH_BOTTOM, "frame", PH_FRAME_H, "small frame", SMALL_H, "raid dps", RAID_DPS)


if __name__ == "__main__":
    main()
