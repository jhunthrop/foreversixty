"""Design exploration: three ways to make the signed-in home feel alive (owner, 2026-11-06: "the site still feels dead
and dull ... boring and lifeless"). Mocks only, nothing here is a build contract.

  A  The character's journey: a class and faction hero band, a level 1 to 60 track with what the next band unlocks,
     the upgrades as a visual strip, a sim that ran on import.
  B  The world is moving: a compact character strip, the raid and guild countdowns, a "Forever right now" feed.
  C  Atmosphere: today's information architecture, art-led (the Fury talent-pane art as a faded backdrop), richer
     surfaces, count-up, hover glow and a shimmer on the new upgrade, all off under prefers-reduced-motion.

One character for all three: Go Hogs, level 20 Fury Warrior, Horde, in a guild, Battle.net imported, no logs or sims.
Upgrades, unlocks, dungeons and talents are read from the active build's real Fury Warrior BiS, talent and trainer
files; the worn items are real runners-up so every delta is a real sim figure. The sim figure, ability shares, feed
rows, tier and countdown events are EXAMPLES and the mocks label them so.

  python3 design/mocks/home_alive.py
  node design/mocks/home_alive_shoot.mjs
"""
import json
from datetime import date
from pathlib import Path

from mocklib import CLASS_COLOR, emblem, data_uri
from ws_lib import page, write, icon, crest_img, QCOL

ROOT = Path(__file__).resolve().parents[2]
BUILD = json.loads((ROOT / 'web/src/data/active-build.json').read_text())['build']
DATA = ROOT / 'data/builds' / BUILD
TODAY = date(2026, 11, 6)
RAIDS_OPEN = date(2026, 12, 9)

CLS, SPEC, FACTION = 'warrior', 'fury', 'horde'
CC = CLASS_COLOR[CLS]
HORDE_RED = '#b3241c'
NAME, LEVEL, RACE, ILVL, GUILD, RANK = 'Go Hogs', 20, 'Orc', 21, 'Olympus XXVII', 'Member'
WHO = (NAME, CLS, f'Level {LEVEL}')
BAND, NEXT_BAND = 20, 30
BANDS = (20, 30, 40, 50, 60)

# The worn item in each slot, a real runner-up from the band-20 list, so each upgrade's delta is the sim's own figure.
WORN = {'waist': "Cobrahn's Grasp", 'hands': "Thorbia's Gauntlets", 'chest': 'Mutant Scale Breastplate',
        'legs': "Veteran's Chain Leggings", 'finger1': "Bounty Hunter's Ring", 'off_hand': 'Cruel Barb'}
SLOT = {'waist': 'Waist', 'hands': 'Hands', 'chest': 'Chest', 'legs': 'Legs', 'finger1': 'Finger', 'off_hand': 'Off hand',
        'main_hand': 'Main hand'}
# Example: the share of damage per ability in the auto-sim (plausible for a level 20 dual-wield Fury build).
ABILITIES = [('Melee', 'inv_sword_04', 54), ('Heroic Strike', 'ability_rogue_ambush', 22), ('Rend', 'ability_gouge', 11)]
# The 9 Dec raids (dates.json note). Icons are stand-ins until the raid list carries its own.
RAIDS = [('Barrow Deeps', 'spell_shadow_raisedead'), ('Hyjal Summit', 'spell_nature_timestop'),
         ("Onyxia's Lair", 'inv_misc_head_dragon_black')]
EXAMPLE = '<span class="ex">Example</span>'


# ------------------------------------------------------------------------------------------------ real data
def _load() -> tuple[dict, dict, dict]:
    bis = json.loads((DATA / 'bis' / f'{CLS}-{SPEC}.json').read_text())
    items = {i['id']: i for i in json.loads((DATA / 'items' / f'{CLS}.json').read_text())['items']}
    band = {b['band']: b for b in bis['bands'] if b['faction'] == FACTION and b.get('preset') == 'bare'}
    return band, items, json.loads((DATA / 'talents' / f'{CLS}.json').read_text())


BANDS_DATA, ITEMS, TALENTS = _load()
NAME_TO_ID = {i['name']: i['id'] for i in ITEMS.values()}


def item(item_id: int) -> dict:
    return ITEMS.get(item_id, {})


def upgrades() -> list[dict]:
    """Each upgrade: the band-20 pick over the worn runner-up, with the sim's delta."""
    out = []
    for s in BANDS_DATA[BAND]['slots']:
        worn = WORN.get(s['slot'])
        alt = next((a for a in s.get('alternatives') or [] if a['item_name'] == worn), None)
        if alt:
            it = item(s['item_id'])
            out.append({'slot': SLOT[s['slot']], 'name': s['item_name'], 'icon': it.get('icon', ''), 'q': it.get('quality', 2),
                        'source': s['source'], 'gain': -alt['dps_delta'], 'worn': worn})
    return sorted(out, key=lambda u: -u['gain'])


UPS = upgrades()
GAIN = sum(u['gain'] for u in UPS)
BIS_DPS = BANDS_DATA[BAND]['set_dps']
NOW_DPS = BIS_DPS - GAIN   # example: what the auto-sim would read for the worn set


def talent_unlocks() -> list[tuple[str, str, int]]:
    """Fury talents the band-30 build adds over band 20: (name, icon, ranks)."""
    tree = next(t for t in TALENTS['trees'] if t['name'] == 'Fury')
    a, b = (BANDS_DATA[n]['talents'].split('-')[1] for n in (BAND, NEXT_BAND))
    return [(t['name'], t['icon'], int(b[i]) - int(a[i])) for i, t in enumerate(tree['talents'][:len(b)]) if int(b[i]) > int(a[i])]


def new_picks(n: int = 4) -> list[dict]:
    """Band-30 picks that come from a dungeon (the dungeons a level 30 can now run)."""
    out = []
    for s in BANDS_DATA[NEXT_BAND]['slots']:
        src = s.get('source') or ''
        if s.get('item_id') and any(d in src for d in ('Razorfen', 'Scarlet', 'Shadowfang', 'Blackfathom')):
            it = item(s['item_id'])
            out.append({'name': s['item_name'], 'icon': it.get('icon', ''), 'q': it.get('quality', 2), 'source': src, 'slot': SLOT.get(s['slot'], s['slot'].replace('_', ' ').capitalize())})
    return out[:n]


def dungeons() -> list[str]:
    seen = []
    for p in new_picks(9):
        d = p['source'].split(':')[0]
        if d not in seen:
            seen.append(d)
    return seen


TRAINED = [('Berserker Stance', 'ability_racial_avatar', 30), ('Berserker Rage', 'spell_nature_ancestralguardian', 30),
           ('Execute', 'inv_sword_48', 24)]   # trainables/warrior.json, levels 21 to 30


# ------------------------------------------------------------------------------------------------ shared bits
def qname(name: str, q: int) -> str:
    return f'<span class="iname" style="color:{QCOL.get(q, "#fff")}">{name}</span>'


def qicon(name: str, size: int, q: int, glow: bool = False) -> str:
    style = f'--q:{QCOL.get(q, "#fff")}'
    return f'<span class="qi{" glow" if glow else ""}" style="{style}">{icon(name, size)}</span>'


def mocknote(lines: list[str]) -> str:
    li = ''.join(f'<li>{x}</li>' for x in lines)
    return f'<div class="pc"><aside class="mnote"><b>Mock notes</b><ul>{li}</ul></aside></div>'


def sec_head(title: str, meta: str = '', action: str = '') -> str:
    m = f'<span class="t-small">{meta}</span>' if meta else ''
    a = f'<a class="act" href="#">{action} &rsaquo;</a>' if action else ''
    return f'<div class="sh2"><h2 class="t-h3">{title}</h2>{m}{a}</div>'


def days_to_raids() -> int:
    return (RAIDS_OPEN - TODAY).days


COMMON_CSS = '''
.ex{display:inline-block;margin-left:8px;padding:0 5px;border:1px dashed #6b6250;border-radius:3px;font:700 10px/16px var(--font-body);letter-spacing:.1em;text-transform:uppercase;color:#a99d80;vertical-align:middle}
.sh2{display:flex;align-items:baseline;gap:12px;padding-bottom:10px;border-bottom:1px solid var(--line-soft);margin-bottom:14px}
.main>.sh2:not(:first-child){margin-top:48px}.rail>.sh2:not(:first-child){margin-top:32px}
.sh2 .act{margin-left:auto;font-weight:600;font-size:14px}
.iname{font-weight:700}
.qi{display:inline-flex;border-radius:6px;flex:none;line-height:0}
.qi .ic{border-color:var(--q)}
.qi.glow .ic{box-shadow:0 0 0 1px var(--q),0 0 18px -2px color-mix(in srgb,var(--q) 70%,transparent)}
.mnote{margin-top:56px;padding:14px 0 0;border-top:1px dashed #3a3326;font-size:13px;line-height:20px;color:var(--muted)}
.mnote b{color:#c9b98f;font-size:12px;letter-spacing:.12em;text-transform:uppercase}
.mnote ul{margin-top:6px;display:grid;gap:2px;list-style:disc;padding-left:18px}
.title-xl{font:400 56px/60px var(--font-title);letter-spacing:.01em;margin:0}
.meta{display:flex;flex-wrap:wrap;gap:4px 18px;font-size:15px;color:var(--nav)}
.meta b{color:var(--strong);font-weight:600}
.em{width:22px;height:22px;object-fit:contain;vertical-align:-5px;margin-right:4px}
.big{font:700 44px/44px var(--font-body);font-variant-numeric:tabular-nums;color:var(--strong)}
.big small{font-size:16px;font-weight:600;color:var(--muted);margin-left:4px}
.sbar{height:6px;border-radius:3px;background:#1c2230;overflow:hidden}.sbar>i{display:block;height:100%;border-radius:3px}
.live{display:inline-block;width:8px;height:8px;border-radius:50%;background:var(--kill);box-shadow:0 0 0 0 #7fd48a88;animation:pulse 2s infinite;margin-right:8px;vertical-align:1px}
@keyframes pulse{0%{box-shadow:0 0 0 0 #7fd48a88}70%{box-shadow:0 0 0 8px #7fd48a00}100%{box-shadow:0 0 0 0 #7fd48a00}}
@media(prefers-reduced-motion:reduce){*{animation:none!important;transition:none!important}}
'''


# ================================================================================================ A: the journey
A_CSS = '''
.heroA{position:relative;overflow:hidden;border-bottom:1px solid var(--line-soft);
 background:radial-gradient(600px 340px at 12% 40%,color-mix(in srgb,var(--cc) 26%,transparent),transparent 70%),
  linear-gradient(100deg,color-mix(in srgb,var(--fc) 30%,#07090d) 0%,#0b0d13 55%,#07090d 100%)}
.heroA .wm{position:absolute;right:-60px;top:-40px;width:520px;opacity:.07;pointer-events:none}
.heroA .pc{position:relative;padding-top:36px;padding-bottom:28px}
.heroA .top{display:grid;grid-template-columns:auto minmax(0,1fr) auto;gap:28px;align-items:center}
.heroA .crest{box-shadow:0 0 0 3px var(--cc),0 0 40px -4px color-mix(in srgb,var(--cc) 60%,transparent)}
.heroA .sim{text-align:right}
.heroA .sim .lab{font:700 12px/16px var(--font-body);letter-spacing:.12em;text-transform:uppercase;color:var(--muted)}
.heroA .sim .auto{font-size:13px;color:var(--kill)}
.heroA .sim .vs{display:flex;gap:10px;align-items:center;justify-content:flex-end;margin-top:8px;font-size:13px;color:var(--muted)}
.heroA .sim .vs .sbar{width:180px}
.track{position:relative;margin:34px 0 0;height:64px}
.track .line{position:absolute;left:0;right:0;top:22px;height:4px;border-radius:2px;background:#232a39}
.track .fill{position:absolute;left:0;top:22px;height:4px;border-radius:2px;background:linear-gradient(90deg,color-mix(in srgb,var(--cc) 50%,transparent),var(--cc))}
.track .mk{position:absolute;top:0;transform:translateX(-50%);text-align:center;font:700 13px/16px var(--font-body);color:var(--muted)}
.track .mk i{display:block;width:14px;height:14px;margin:15px auto 8px;border-radius:50%;background:#07090d;border:2px solid #3a4256}
.track .mk.done i{background:var(--cc);border-color:var(--cc)}
.track .mk.next{color:var(--gold)}.track .mk.next i{border-color:var(--gold);box-shadow:0 0 0 4px #e5b95522}
.track .you{position:absolute;top:-4px;transform:translateX(-50%);font:700 12px/16px var(--font-body);color:var(--cc);white-space:nowrap}
.track .start{left:0;transform:none}.track .end{right:0;left:auto;transform:none}
.unlock{display:grid;grid-template-columns:auto repeat(3,minmax(0,1fr));gap:12px 40px;margin-top:6px;padding-top:18px;border-top:1px solid #ffffff12}
.unlock .k{font:700 12px/16px var(--font-body);letter-spacing:.12em;text-transform:uppercase;color:var(--muted);margin-bottom:10px}
.unlock .lead{max-width:220px}
.unlock .lead .n{font:400 30px/34px var(--font-title);color:var(--gold)}
.unlock .lead p{font-size:14px;line-height:21px;color:var(--nav);margin-top:6px}
.ul{display:flex;flex-direction:column;gap:8px}
.ul .r{display:flex;gap:10px;align-items:center;font-size:14px;line-height:18px;min-width:0}
.ul .r .s{font-size:12px;color:var(--muted)}
.ul .r>span:last-child{min-width:0}
.bodyA{display:grid;grid-template-columns:minmax(0,1fr) 340px;gap:var(--gap);padding-top:40px;align-items:start}
@media(min-width:1760px){.bodyA{grid-template-columns:minmax(0,1fr) 400px}}
.bodyA .main{min-width:0}
.strip{display:grid;grid-template-columns:repeat(3,minmax(0,1fr));gap:14px}
.tile{display:grid;grid-template-columns:56px minmax(0,1fr);gap:8px 14px;align-items:start;padding:14px 16px;color:inherit;border-radius:8px;background:linear-gradient(180deg,color-mix(in srgb,var(--q) 9%,#0d111a),#0b0e15);border:1px solid color-mix(in srgb,var(--q) 28%,#1c2230);min-width:0}
.tile:hover{border-color:color-mix(in srgb,var(--q) 60%,#1c2230)}
.tile .iname{font-size:15px;line-height:19px}
.tile .s{font-size:12px;line-height:16px;color:var(--muted)}
.tile>.qi{grid-row:span 2}.tile .d{display:flex;align-items:center;gap:10px}.tile .d .sbar{flex:1}
.tile .d b{font:700 20px/22px var(--font-body);font-variant-numeric:tabular-nums;color:var(--kill)}
.tile .d span{font-size:12px;color:var(--muted)}
.tile .sbar>i{background:var(--kill)}
.simA{display:grid;grid-template-columns:240px minmax(0,1fr);gap:40px;align-items:center}
.ab{display:grid;grid-template-columns:36px 140px minmax(0,1fr) 48px;gap:12px;align-items:center;padding:7px 0;font-size:15px}
.ab .sbar{height:10px}.ab .sbar>i{background:linear-gradient(90deg,color-mix(in srgb,var(--cc) 45%,transparent),var(--cc))}
.ab b{text-align:right;font-variant-numeric:tabular-nums}
.rrow{display:flex;gap:12px;align-items:center;padding:10px 0;border-bottom:1px solid var(--line-soft);font-size:14px;line-height:20px}
.rrow:last-child{border-bottom:0}
.rrow .t{font-weight:700;color:var(--strong)}.rrow .s{font-size:13px;color:var(--muted)}
.ok{width:22px;height:22px;border-radius:50%;background:#7fd48a22;color:var(--kill);display:inline-flex;align-items:center;justify-content:center;font-weight:700;font-size:13px;flex:none}
'''


def track() -> str:
    pct = lambda lv: (lv - 1) / 59 * 100
    marks = []
    for b in BANDS:
        cls = 'done' if b <= LEVEL else ('next' if b == NEXT_BAND else '')
        marks.append(f'<div class="mk {cls}" style="left:{pct(b):.2f}%"><i></i>{b}</div>')
    marks.append('<div class="mk done start" style="left:0"><i style="margin-left:0"></i>1</div>')
    you = f'<div class="you" style="left:{pct(LEVEL):.2f}%">You &middot; {LEVEL}</div>'
    return (f'<div class="track" aria-label="Level {LEVEL} of 60"><div class="line"></div><div class="fill" style="width:{pct(LEVEL):.2f}%"></div>'
            f'{"".join(marks)}{you}</div>')


def unlock_block() -> str:
    tal = ''.join(f'<div class="r">{icon(i, 32)}<span>{n}<br><span class="s">{r} point{"s" if r > 1 else ""}</span></span></div>' for n, i, r in talent_unlocks())
    trn = ''.join(f'<div class="r">{icon(i, 32)}<span>{n}<br><span class="s">Trainer, level {lv}</span></span></div>' for n, i, lv in TRAINED)
    picks = ''.join(f'<div class="r">{qicon(p["icon"], 32, p["q"])}<span>{qname(p["name"], p["q"])}<br><span class="s">{p["source"]}</span></span></div>' for p in new_picks(3))
    lead = (f'<div class="lead"><div class="k">Next band</div><div class="n">Level {NEXT_BAND}</div>'
            f'<p>{NEXT_BAND - LEVEL} levels away. Your list, talents and trainer all change at {NEXT_BAND}.</p>'
            f'<p><a class="link" href="#">Plan level {NEXT_BAND} talents &rsaquo;</a><br><a class="link" href="#">Band {NEXT_BAND} on best in slot &rsaquo;</a></p></div>')
    return (f'<div class="unlock">{lead}<div><div class="k">New talents</div><div class="ul">{tal}</div></div>'
            f'<div><div class="k">Trainer</div><div class="ul">{trn}</div></div>'
            f'<div><div class="k">New picks &middot; {", ".join(dungeons())}</div><div class="ul">{picks}</div></div></div>')


def hero_a() -> str:
    meta = (f'<div class="meta"><span><b>Level {LEVEL} Fury Warrior</b></span><span><img class="em" src="{emblem("horde")}" alt="">Horde {RACE}</span>'
            f'<span>Item level {ILVL}</span><span>&lt;{GUILD}&gt;</span></div>')
    pct = NOW_DPS / BIS_DPS * 100
    sim = (f'<div class="sim"><div class="lab">Your DPS now</div><div class="big">{NOW_DPS:.1f}<small>DPS</small></div>'
           f'<div class="auto">Simmed automatically on import &middot; 6 Nov {EXAMPLE}</div>'
           f'<div class="vs"><span>{pct:.0f}% of best in slot</span><div class="sbar"><i style="width:{pct:.0f}%;background:var(--cc)"></i></div><span class="strong">{BIS_DPS:.1f}</span></div></div>')
    top = (f'<div class="top">{crest_img(CLS, 112)}<div><h1 class="title-xl" style="color:{CC}">{NAME}</h1>{meta}</div>{sim}</div>')
    return (f'<section class="heroA full-bleed" style="--cc:{CC};--fc:{HORDE_RED}"><img class="wm" src="{data_uri(ROOT / "web/public/icons/hd/faction/horde-512.webp")}" alt="">'
            f'<div class="pc">{top}{track()}{unlock_block()}</div></section>')


def upgrade_strip() -> str:
    top = max(u['gain'] for u in UPS)
    tiles = ''.join(
        f'<a class="tile" href="#" style="--q:{QCOL[u["q"]]}">{qicon(u["icon"], 56, u["q"], glow=True)}'
        f'<div>{qname(u["name"], u["q"])}<div class="s">{u["slot"]} &middot; {u["source"]}</div></div>'
        f'<div class="d"><b>+{u["gain"]:.1f}</b><span>DPS</span><div class="sbar"><i style="width:{u["gain"] / top * 100:.0f}%"></i></div></div>'
        f'<div></div><div class="s">vs your {u["worn"]}</div></a>' for u in UPS)
    return sec_head('Your upgrades', f'Levels {BAND} to {BAND + 9} &middot; +{GAIN:.1f} DPS in all', 'All on best in slot') + f'<div class="strip">{tiles}</div>'


def sim_a() -> str:
    rows = ''.join(f'<div class="ab">{icon(i, 36)}<span class="strong">{n}</span><div class="sbar"><i style="width:{p / ABILITIES[0][2] * 100:.0f}%"></i></div><b>{p}%</b></div>' for n, i, p in ABILITIES)
    left = (f'<div><div class="big">{NOW_DPS:.1f}<small>DPS</small></div><p class="t-small" style="margin-top:8px">Your imported gear and talents against a level 22 target, 3 minutes. '
            f'Simmed automatically on import {EXAMPLE}</p></div>')
    return sec_head('Where your damage comes from', 'Top 3 abilities by damage share', 'Open the simulator') + f'<div class="simA" style="--cc:{CC}">{left}<div>{rows}</div></div>'


def rail_a() -> str:
    guild = sec_head('Guild', '', 'Guild page') + (f'<div class="rrow">{crest_img(CLS, 36)}<div><div class="t">{GUILD}</div><div class="s">{RANK} &middot; 14 members online {EXAMPLE}</div></div></div>')
    raids = sec_head('First raids', '9 Dec') + ''.join(f'<div class="rrow">{icon(i, 32)}<div><div class="t">{n}</div><div class="s">Level 60</div></div></div>' for n, i in RAIDS)
    setup = sec_head('Set up') + (f'<div class="rrow"><span class="ok">&#10003;</span><div><div class="t">Battle.net import</div><div class="s">Imported 6 Nov &middot; gear, level, talents</div></div></div>'
                                  f'<div class="rrow"><span class="ok" style="background:#e5b9551c;color:var(--gold)">+</span><div><div class="t">Addon and companion (optional)</div><div class="s">Uploads your logs as you play. <a href="#">Get them &rsaquo;</a></div></div></div>')
    return guild + raids + setup


def page_a() -> str:
    body = (hero_a() + f'<div class="pc"><div class="bodyA"><div class="main">{upgrade_strip()}{sim_a()}</div><aside class="rail">{rail_a()}</aside></div></div>'
            + mocknote([
                'Real today: the character (Battle.net import), the band-20 upgrades and their DPS deltas (BiS file; worn items are real runners-up), '
                f'the band-{NEXT_BAND} talent build, trainer spells and dungeon picks (talents, trainables and BiS files), the 9 Dec raid list (dates.json).',
                'New work: an auto-sim on import (one server sim per character per gear change) that stores DPS and ability shares; '
                'the guild online count (guild roster presence).',
                'The hero is the one focal moment: crest, name in class colour, the sim figure; the journey track and unlock row sit inside the band.',
                'Upgrades become a strip of six tiles; the green bar is the delta against the largest, so the order reads at a glance.',
            ]) + f'<style>{COMMON_CSS}{A_CSS}</style>')
    return page('Forever Sixty', body, who=WHO, note='Home alive, direction A: the journey. Exploration mock.')


# ================================================================================================ B: the world is moving
FEED = [  # (icon, quality or 0, html, kind, minutes ago) -- all examples
    ('inv_axe_23', 3, '<b>Brakka</b> <span class="m">Fury Warrior</span> looted <span class="iname" style="color:#3d94f0">Butcher\'s Cleaver</span> in Shadowfang Keep', 'Your spec', 4),
    ('spell_shadow_raisedead', 0, '<b>Nightfall</b> uploaded a full Shadowfang Keep, 8 bosses in 41 min', 'Logs', 9),
    ('ability_warrior_innerrage', 0, '<b>Fury Warrior, 20 to 29:</b> Grukk climbed 6 places to <b>#3</b> on the rankings', 'Your spec', 17),
    ('inv_chest_leather_08', 3, '<b>Best in slot changed</b> for Fury Warrior 20 to 29: <span class="iname" style="color:#3d94f0">Blackened Defias Armor</span> is the new chest', 'Your spec', 62),
    ('inv_misc_book_09', 0, '<b>Fury Warrior guide</b> updated: the leveling rotation from 20 now opens with Rend', 'Guides', 140),
    ('inv_misc_head_dragon_black', 0, '<b>Olympus XXVII</b> posted the first Onyxia sign-up: 22 in, 3 late', 'Your guild', 205),
    ('spell_shadow_deathpact', 0, '<b>Talent tree hotfix</b> picked up in the planner: Death Wish now costs 10 rage', 'Patch', 300),
]


def ago(m: int) -> str:
    return f'{m} min ago' if m < 60 else f'{m // 60} h ago'


B_CSS = '''
.stripB{border-bottom:1px solid var(--line-soft);background:linear-gradient(90deg,color-mix(in srgb,var(--cc) 10%,transparent),transparent 40%)}
.stripB .pc{display:flex;align-items:center;gap:20px;padding-block:22px 18px}
.stripB .t{font:400 36px/40px var(--font-title);margin:0}
.stripB .stats{margin-left:auto;display:flex;gap:36px}
.stat .v{font:700 26px/28px var(--font-body);font-variant-numeric:tabular-nums;color:var(--strong)}
.stat .v small{font-size:13px;color:var(--muted);margin-left:3px;font-weight:600}
.stat .l{font-size:12px;letter-spacing:.1em;text-transform:uppercase;font-weight:700;color:var(--muted)}
.cds{display:grid;grid-template-columns:minmax(0,1.3fr) minmax(0,1fr);gap:var(--gap);padding-top:28px}
.cd{position:relative;overflow:hidden;padding:20px 24px;border-radius:8px;border:1px solid var(--line);background:#0b0e15}
.cd.raid{background:radial-gradient(500px 200px at 100% 0,#5a1d7a33,transparent 70%),#0b0e15;border-color:#3b2a4d}
.cd.guild{background:radial-gradient(400px 200px at 100% 0,#e5b9551a,transparent 70%),#0b0e15}
.cd .k{font:700 12px/16px var(--font-body);letter-spacing:.12em;text-transform:uppercase;color:var(--muted)}
.cd .row{display:flex;align-items:flex-end;gap:28px;margin-top:10px}
.clock{display:flex;gap:14px;font-variant-numeric:tabular-nums}
.clock div{text-align:center}.clock b{display:block;font:700 40px/42px var(--font-body);color:var(--strong)}
.clock span{font-size:11px;letter-spacing:.12em;text-transform:uppercase;color:var(--muted)}
.cd .raids{display:flex;gap:18px;margin-left:auto}
.cd .raids div{display:flex;flex-direction:column;align-items:center;gap:6px;font-size:13px;font-weight:600;color:var(--nav);text-align:center;width:88px}
.cd .what{margin-left:auto;text-align:right;font-size:14px;color:var(--nav)}
.cd .what b{display:block;font-size:18px;color:var(--strong)}
.bodyB{display:grid;grid-template-columns:minmax(0,1fr) 340px;gap:var(--gap);padding-top:40px;align-items:start}
@media(min-width:1760px){.bodyB{grid-template-columns:minmax(0,1fr) 400px}}
.filt{display:inline-flex;gap:4px;margin-left:16px}
.filt a{padding:3px 10px;border-radius:999px;font-size:13px;font-weight:600;color:var(--nav);border:1px solid var(--line-soft)}
.filt a.on{color:var(--gold);border-color:#e5b95566;background:#e5b95512}
.fr{display:grid;grid-template-columns:40px minmax(0,1fr) auto;gap:14px;align-items:center;padding:12px 0;border-bottom:1px solid var(--line-soft);font-size:15px;line-height:22px}
.fr:hover{background:linear-gradient(90deg,#ffffff06,transparent)}
.fr b{color:var(--strong)}.fr .m{color:var(--muted)}
.fr .when{text-align:right;font-size:13px;color:var(--muted);white-space:nowrap}
.fr .when .kind{display:block;font-size:11px;letter-spacing:.1em;text-transform:uppercase;font-weight:700;color:#7f7864}
.fr.fresh{animation:slidein .6s ease-out}
@keyframes slidein{from{opacity:0;transform:translateY(-6px)}to{opacity:1;transform:none}}
.tier{display:flex;align-items:center;gap:16px}
.tier .L{width:60px;height:60px;border-radius:8px;display:flex;align-items:center;justify-content:center;font:700 34px/1 var(--font-body);color:#17110a;background:#e5b955}
.mv{display:grid;grid-template-columns:28px 32px minmax(0,1fr) auto;gap:10px;align-items:center;padding:8px 0;border-bottom:1px solid var(--line-soft);font-size:14px}
.mv .up{color:var(--kill);font-weight:700;font-size:13px}
.mini{display:flex;gap:10px;align-items:center;padding:8px 0;border-bottom:1px solid var(--line-soft);font-size:14px;line-height:18px}
.mini .s{font-size:12px;color:var(--muted)}.mini .g{margin-left:auto;font-weight:700;color:var(--kill);font-variant-numeric:tabular-nums}
'''


def clock(d: int, h: int, m: int) -> str:
    return f'<div class="clock" data-clock>' + ''.join(f'<div><b>{v:02d}</b><span>{u}</span></div>' for v, u in ((d, 'days'), (h, 'hours'), (m, 'min'))) + '</div>'


def strip_b() -> str:
    stats = (f'<div class="stats"><div class="stat"><div class="v">{NOW_DPS:.1f}<small>DPS</small></div><div class="l">Simmed on import</div></div>'
             f'<div class="stat"><div class="v" style="color:var(--kill)">+{GAIN:.1f}<small>DPS</small></div><div class="l">{len(UPS)} upgrades</div></div>'
             f'<div class="stat"><div class="v">{NEXT_BAND - LEVEL}<small>levels</small></div><div class="l">To band {NEXT_BAND}</div></div></div>')
    meta = f'<div class="meta"><span><b>Level {LEVEL} Fury Warrior</b></span><span><img class="em" src="{emblem("horde")}" alt="">Horde {RACE}</span><span>&lt;{GUILD}&gt;</span></div>'
    return f'<section class="stripB full-bleed" style="--cc:{CC}"><div class="pc">{crest_img(CLS, 64)}<div><h1 class="t" style="color:{CC}">{NAME}</h1>{meta}</div>{stats}</div></section>'


def countdowns() -> str:
    raids = ''.join(f'<div>{icon(i, 48)}<span>{n}</span></div>' for n, i in RAIDS)
    raid = (f'<div class="cd raid"><div class="k"><span class="live"></span>First raids open &middot; Wed 9 Dec</div>'
            f'<div class="row">{clock(days_to_raids(), 4, 12)}<div class="raids">{raids}</div></div></div>')
    guild = (f'<div class="cd guild"><div class="k">{GUILD} &middot; next guild event {EXAMPLE}</div>'
             f'<div class="row">{clock(1, 22, 41)}<div class="what"><b>Dungeon night: Shadowfang Keep</b>Sat 8 Nov, 8 pm CST &middot; 9 in &middot; you: not signed<br><a href="#">Sign up on the raid page &rsaquo;</a></div></div></div>')
    return f'<div class="pc"><div class="cds">{raid}{guild}</div></div>'


def feed() -> str:
    filt = '<span class="filt"><a class="on" href="#">All</a><a href="#">Your spec</a><a href="#">Your guild</a><a href="#">Logs</a></span>'
    rows = ''.join(f'<a class="fr{" fresh" if i == 0 else ""}" href="#" style="color:inherit">{qicon(ic, 40, q) if q else icon(ic, 40)}<div>{html}</div>'
                   f'<div class="when">{ago(m)}<span class="kind">{kind}</span></div></a>' for i, (ic, q, html, kind, m) in enumerate(FEED))
    head = f'<div class="sh2"><h2 class="t-h3"><span class="live"></span>Forever right now</h2>{filt}{EXAMPLE}<a class="act" href="#">Everything &rsaquo;</a></div>'
    return head + rows


def rail_b() -> str:
    tier = sec_head('Your spec this week', '', 'Tier list') + (
        f'<div class="tier"><div class="L">A</div><div><div class="strong" style="font-weight:700;font-size:16px">Fury Warrior at 20 to 29</div>'
        f'<div class="t-small">Up from B last week, #4 of 23 melee {EXAMPLE}</div></div></div>')
    movers = [('3', 'Grukk', '+6'), ('5', 'Brakka', '+4'), ('9', 'Zugzug', '+11')]
    mv = sec_head('Top movers', 'Fury, 20 to 29', 'Rankings') + ''.join(
        f'<div class="mv"><b class="strong">#{r}</b>{crest_img(CLS, 28)}<span style="color:{CC};font-weight:700">{n}</span><span class="up">&#9650; {u}</span></div>' for r, n, u in movers)
    ups = sec_head('Your upgrades', f'+{GAIN:.1f} DPS', 'Best in slot') + ''.join(
        f'<div class="mini">{qicon(u["icon"], 32, u["q"])}<div>{qname(u["name"], u["q"])}<div class="s">{u["slot"]} &middot; {u["source"]}</div></div><span class="g">+{u["gain"]:.1f}</span></div>' for u in UPS[:3])
    return ups + tier + mv


def page_b() -> str:
    body = (strip_b() + countdowns() + f'<div class="pc"><div class="bodyB"><div class="main">{feed()}</div><aside class="rail">{rail_b()}</aside></div></div>'
            + mocknote([
                'Real today: the character, the upgrades and their deltas (BiS file), the 9 Dec raid date and list (dates.json), the guild raid planner '
                '(a raid title, time and sign-up counts), rankings and the tier list exist per spec.',
                'New work: a site activity feed (an events table written by log upload, BiS and guide regen, ranking snapshots and guild raid posts, read '
                'as one list with a spec and guild filter); rank movement needs a weekly rankings snapshot; the tier delta needs last week\'s tier; '
                'the auto-sim on import for the DPS stat; the countdowns tick client-side from ISO times we already have.',
                'Every feed row, mover, tier and the guild event are examples. The live dot pulses; the newest row slides in; both stop under reduced motion.',
                'Risk: the feed is only alive if events arrive; a quiet hour needs a floor (BiS and guide changes from the nightly, at least).',
            ]) + f'<style>{COMMON_CSS}{B_CSS}</style>')
    return page('Forever Sixty', body, who=WHO, note='Home alive, direction B: the world is moving. Exploration mock.')


# ================================================================================================ C: atmosphere
C_CSS = '''
.heroC{position:relative;overflow:hidden;min-height:300px;border-bottom:1px solid #ffffff10}
.heroC .art{position:absolute;inset:-40px 0 -40px 30%;background-image:var(--art);background-size:cover;background-position:center 30%;
 filter:blur(2px) sepia(.5) saturate(1.6) brightness(2.8) contrast(1.1);opacity:.7;mask-image:linear-gradient(90deg,transparent,#000 35%,#000 80%,transparent);-webkit-mask-image:linear-gradient(90deg,transparent,#000 35%,#000 80%,transparent)}
.heroC .tint{position:absolute;inset:0;background:radial-gradient(700px 360px at 18% 55%,color-mix(in srgb,var(--cc) 22%,transparent),transparent 70%),
 linear-gradient(0deg,#07090d 0%,transparent 45%),linear-gradient(90deg,#07090dee 25%,#07090d55 60%,#07090d99 100%)}
.heroC .ember{position:absolute;inset:0;background:radial-gradient(2px 2px at 62% 40%,#ffb15a,transparent),radial-gradient(1.5px 1.5px at 70% 62%,#ff8a3a,transparent),radial-gradient(2px 2px at 81% 30%,#ffc27a,transparent),radial-gradient(1.5px 1.5px at 88% 55%,#ff9a4a,transparent);opacity:.6;animation:drift 9s ease-in-out infinite alternate}
@keyframes drift{to{transform:translateY(-14px);opacity:.25}}
.heroC .pc{position:relative;display:grid;grid-template-columns:minmax(0,1fr) auto;gap:32px;align-items:end;padding-top:56px;padding-bottom:28px}
.heroC .id{display:flex;gap:24px;align-items:center}
.heroC .crest{box-shadow:0 0 0 2px var(--cc),0 0 0 8px #07090d88,0 0 60px -6px color-mix(in srgb,var(--cc) 70%,transparent)}
.heroC .ans{text-align:right}
.heroC .ans .lab{font:700 12px/16px var(--font-body);letter-spacing:.12em;text-transform:uppercase;color:var(--muted)}
.heroC .ans .big{font-size:56px;line-height:56px;color:var(--kill);text-shadow:0 0 30px #7fd48a44}
.sumC{display:flex;gap:14px;align-items:center;margin-top:20px;padding:12px 16px;border-radius:8px;background:linear-gradient(90deg,#e5b95514,transparent 70%);border-left:2px solid var(--gold);color:var(--gold);font-size:15px}
.splitC{display:grid;grid-template-columns:minmax(0,1fr) 340px;gap:var(--gap);padding-top:36px;align-items:start}
@media(min-width:1760px){.splitC{grid-template-columns:minmax(0,1fr) 400px}}
.splitC .main{min-width:0}
.plateC{border-radius:10px;border:1px solid #2a3245;background:linear-gradient(180deg,#111623,#0b0e15);box-shadow:0 1px 0 #ffffff08 inset,0 20px 40px -24px #000}
.ir{position:relative;display:grid;grid-template-columns:80px 52px minmax(0,1fr) auto;gap:16px;align-items:center;padding:14px 20px;border-bottom:1px solid #ffffff0a;overflow:hidden;transition:background .2s}
.ir:last-child{border-bottom:0}
.ir::before{content:"";position:absolute;left:0;top:0;bottom:0;width:3px;background:var(--q);opacity:.7}
.ir:hover{background:linear-gradient(90deg,color-mix(in srgb,var(--q) 10%,transparent),transparent 60%)}
.ir:hover .qi .ic{box-shadow:0 0 0 1px var(--q),0 0 22px color-mix(in srgb,var(--q) 80%,transparent)}
.ir .slot{font-size:13px;color:var(--muted);font-weight:600}
.ir .iname{font-size:18px;line-height:22px}
.ir .s{font-size:13px;color:var(--muted)}
.ir .g{text-align:right}.ir .g b{font:700 22px/24px var(--font-body);font-variant-numeric:tabular-nums;color:var(--kill)}.ir .g small{font-size:12px;color:var(--muted);margin-left:3px}
.ir .g .s{display:block}
.ir.new::after{content:"";position:absolute;inset:0;background:linear-gradient(105deg,transparent 40%,#ffffff14 50%,transparent 60%);background-size:250% 100%;animation:shimmer 3.2s ease-in-out infinite}
@keyframes shimmer{from{background-position:120% 0}to{background-position:-120% 0}}
.newtag{display:inline-block;margin-left:8px;padding:0 6px;border-radius:3px;background:#7fd48a22;color:var(--kill);font:700 10px/16px var(--font-body);letter-spacing:.1em;text-transform:uppercase;vertical-align:2px}
.inv{display:grid;grid-template-columns:52px minmax(0,1fr) auto;gap:16px;align-items:center;padding:16px 20px}
.inv .t{font-weight:700;color:var(--strong);font-size:16px}.inv .s{font-size:13px;color:var(--muted)}
.btnC{display:inline-flex;align-items:center;height:38px;padding:0 16px;border-radius:6px;border:1px solid #e5b95566;background:linear-gradient(180deg,#e5b95522,#e5b9550a);color:var(--gold);font:700 12px/16px var(--font-body);letter-spacing:.1em;text-transform:uppercase}
.rc{display:flex;gap:12px;align-items:center;padding:12px 0;border-bottom:1px solid var(--line-soft);font-size:14px;line-height:20px}
.rc:last-child{border-bottom:0}.rc .t{font-weight:700;color:var(--strong)}.rc .s{font-size:13px;color:var(--muted)}
.ok{width:22px;height:22px;border-radius:50%;background:#7fd48a22;color:var(--kill);display:inline-flex;align-items:center;justify-content:center;font-weight:700;font-size:13px;flex:none}
.motion{display:grid;grid-template-columns:repeat(4,minmax(0,1fr));gap:20px;margin-top:8px}
.motion div{font-size:13px;line-height:19px;color:var(--nav)}.motion b{display:block;color:var(--strong);font-size:14px}
'''

COUNT_UP_JS = '''<script>
(()=>{const r=matchMedia('(prefers-reduced-motion: reduce)').matches;document.querySelectorAll('[data-count]').forEach(el=>{
const to=parseFloat(el.dataset.count),pre=el.dataset.pre||'';if(r){return}const t0=performance.now(),d=900;
const step=t=>{const k=Math.min(1,(t-t0)/d),e=1-Math.pow(1-k,3);el.textContent=pre+(to*e).toFixed(1);if(k<1)requestAnimationFrame(step)};requestAnimationFrame(step)})})();
</script>'''


def hero_c() -> str:
    art = data_uri(DATA / 'trees' / f'{CLS}{SPEC}.webp')
    meta = (f'<div class="meta" style="margin-top:6px"><span><b>Level {LEVEL} Fury Warrior</b></span><span><img class="em" src="{emblem("horde")}" alt="">Horde {RACE}</span>'
            f'<span>Item level {ILVL}</span></div>')
    ans = (f'<div class="ans"><div class="lab">Upgrades at {BAND} to {BAND + 9}</div><div class="big"><span data-count="{GAIN:.1f}" data-pre="+">+{GAIN:.1f}</span><small>DPS</small></div>'
           f'<a class="link" href="#">Open best in slot &rsaquo;</a></div>')
    return (f'<section class="heroC full-bleed" style="--cc:{CC};--art:url({art})"><div class="art"></div><div class="tint"></div><div class="ember"></div>'
            f'<div class="pc"><div><div class="id">{crest_img(CLS, 104)}<div><h1 class="title-xl" style="color:{CC}">{NAME}</h1>{meta}</div></div>'
            f'<div class="sumC">Next: get {UPS[0]["name"]} from {UPS[0]["source"]}. <a class="link" href="#">Open best in slot &rsaquo;</a></div></div>{ans}</div></section>')


def upgrades_c() -> str:
    rows = ''.join(
        f'<a class="ir{" new" if i == 0 else ""}" href="#" style="--q:{QCOL[u["q"]]};color:inherit"><span class="slot">{u["slot"]}</span>{qicon(u["icon"], 52, u["q"], glow=True)}'
        f'<div>{qname(u["name"], u["q"])}{"<span class=newtag>New</span>" if i == 0 else ""}<div class="s">{u["source"]}</div></div>'
        f'<div class="g"><b>+{u["gain"]:.1f}</b><small>DPS</small><span class="s">vs your {u["worn"]}</span></div></a>' for i, u in enumerate(UPS[:3]))
    return sec_head('Your upgrades', f'Levels {BAND} to {BAND + 9}', f'All {len(UPS)} on best in slot') + f'<div class="plateC">{rows}</div>'


def invite(ic: str, title: str, sub: str, btn: str) -> str:
    return f'<div class="plateC"><div class="inv">{icon(ic, 52)}<div><div class="t">{title}</div><div class="s">{sub}</div></div><a class="btnC" href="#">{btn}</a></div></div>'


def rail_c() -> str:
    guild = sec_head('Guild') + f'<div class="rc">{crest_img(CLS, 36)}<div><div class="t">{GUILD}</div><div class="s">{RANK}</div></div></div>'
    logs = sec_head('Your logs', '', 'All logs') + f'<div class="rc">{icon("inv_misc_book_11", 36)}<div><div class="t">Upload a log</div><div class="s">The companion uploads your logs as you play.</div></div></div>'
    setup = sec_head('Set up') + (f'<div class="rc"><span class="ok">&#10003;</span><div><div class="t">Battle.net import</div><div class="s">Imported 6 Nov</div></div></div>'
                                  f'<div class="rc">{icon("inv_misc_bag_10", 36)}<div><div class="t">Addon and companion (optional)</div><div class="s">Rates every item in your bags and uploads your logs. <a href="#">Get them &rsaquo;</a></div></div></div>')
    return guild + logs + setup


def motion_notes() -> str:
    items = [('Count-up on load', 'The answer figure counts from 0 over 0.9 s, ease-out; once per visit.'),
             ('Icon hover glow', 'An item row lights its quality colour along the left and glows the icon frame.'),
             ('New-upgrade shimmer', 'A slow sheen crosses the top upgrade while it is new since your last visit.'),
             ('Embers', 'Four faint embers drift over the art, 9 s loop. Everything is off under prefers-reduced-motion.')]
    return sec_head('Motion', 'Mock note, not page content') + '<div class="motion">' + ''.join(f'<div><b>{t}</b>{d}</div>' for t, d in items) + '</div>'


def page_c() -> str:
    main = (upgrades_c()
            + sec_head('Simulator', '', 'Open the simulator') + invite('ability_warrior_innerrage', 'Run your first sim', 'Your imported gear and talents are ready to sim.', 'Sim now')
            + motion_notes())
    body = (hero_c() + f'<div class="pc"><div class="splitC"><div class="main">{main}</div><aside class="rail">{rail_c()}</aside></div></div>'
            + mocknote([
                'Real today: everything on the page; the information architecture is the shipped home (head, next step, upgrades, simulator, guild, logs, set up).',
                'New work: art and motion only. The backdrop is the Fury talent-pane art from the client (data/builds/*/trees), one per spec, already in the repo; '
                'it is 300 px and must stay blurred and faded, or we trace an HD version as we did the crests.',
                'The "New" shimmer needs the last-visit time per character (a local timestamp is enough).',
                'Weakness: the simulator, logs and set-up blocks are still invitations, so the lower half stays as empty as today.',
            ]) + f'<style>{COMMON_CSS}{C_CSS}</style>' + COUNT_UP_JS)
    return page('Forever Sixty', body, who=WHO, note='Home alive, direction C: atmosphere. Exploration mock.')


def main() -> None:
    for key, render in (('a', page_a), ('b', page_b), ('c', page_c)):
        print(write(f'web-home-alive-{key}.html', render()))


if __name__ == '__main__':
    main()
