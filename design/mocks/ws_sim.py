"""Web-layout mock: the simulator (/sim landing, a live result on /sim, a saved run /sim/<id>), composed ONLY from ws_prims.

Spec: design/specs/2026-11-06-sim.md. The run is web/src/fixtures/sim/result-full.json, the e2e fixture behind /sim/simfullfixt1:
a real forever-sim run of the published warrior-fury level 60 raid set (10,000 runs). Every name, icon, rank, item, talent,
rotation line, buff and figure is read at render time from that file and from data/builds/<active build>/ and data/curated/
(spells.json, spelltips, spellicons.json, items/<class>.json, talents/<class>.json, simbuffs.json, apl/<spec>.json,
engine-spell-names.json, extra-spell-icons.json). Figures follow the shipped rules (report/ability-table.ts, sim/warnings.ts,
sim/gear-card.ts). Typed in, and marked SAMPLE: the signed-in character's name "Thandrel", the last sims, the live estimate and runs done
mid-run, the engine-refusal reason and the "Your sims" rows, because no public account exists to read them from. Every
error band is computed from the fixture's own spread (band()).

  FS_DATA_ROOT=<a checkout at the live build> python3 design/mocks/ws_sim.py
  for p in web-sim web-sim-you web-sim-running web-sim-result web-sim-saved web-sim-bar web-sim-more web-sim-states; do
    node design/mocks/ws_shoot.mjs $p 390,1440,1920,2000; done
  python3 design/mocks/ws_sim.py compare
"""
import base64
import decimal
import json
import os
import re
import sys
from pathlib import Path

from mocklib import CLASS_COLOR
from ws_prims import (Button as _Button, Facts, Field, Figure, Figures, FormRow, Histogram, IconList, Lanes, LanesLegend, Link, Meter, Notice,
                      Page, PageHeader, Plate, Row, RowGrid, Rows, Section, Sheet, Skeleton, StatRow, StatTable, TabStrip, Trees, build_page, talent_ranks_from_strings, talent_tree_dicts,
                      control_bar, crest_img, seg, sel, write, ItemTip, PaperdollFrame)

ROOT = Path(__file__).resolve().parents[2]
DATA_ROOT = Path(os.environ.get('FS_DATA_ROOT', ROOT))
BUILD = json.loads((DATA_ROOT / 'web/src/data/active-build.json').read_text())['build']
B = DATA_ROOT / 'data/builds' / BUILD
CUR = DATA_ROOT / 'data/curated'
RES = json.loads((DATA_ROOT / 'web/src/fixtures/sim/result-full.json').read_text())
REQ, CHAR, SUM = RES['request'], RES['request']['character'], RES['summary']


def copy_string(module: str, key: str) -> str:
    """One plain string from a web copy module (`key: '...'`, the string on its line or the next), so a lead reads the
    shipped words: `copy_string('lib/sim/page-copy.ts', 'lead')`."""
    text = (DATA_ROOT / 'web/src' / module).read_text()
    found = re.search(rf"^\s+{key}:\s*'((?:[^'\\]|\\.)*)',$", text, re.M)
    if found is None:
        raise SystemExit(f'ws_sim: no string {key} in web/src/{module}')
    return found.group(1).replace("\\'", "'")
CLS = CHAR['class']
SPEC = REQ['spec']

SPELL_NAME = {s['id']: s['name'] for s in json.loads((B / 'spells.json').read_text())}
ENGINE_NAME = {int(k): v['name'] for k, v in json.loads((CUR / 'engine-spell-names.json').read_text())['spells'].items()}
TIPS = {int(k): v for k, v in json.loads((B / f'spelltips/{CLS}.json').read_text())['spells'].items()}
SPELL_ICON = json.loads((B / 'spellicons.json').read_text())
ICON_ROWS = {int(k): v for k, v in json.loads((CUR / 'extra-spell-icons.json').read_text())['icon_rows'].items()}
ITEMS = {i['id']: i for i in json.loads((B / f'items/{CLS}.json').read_text())['items']}
BUFFS = json.loads((B / 'simbuffs.json').read_text())['entries']
CONSUMES = {c['id']: c for c in json.loads((B / 'simconsumes.json').read_text())}
TALENTS = json.loads((B / f'talents/{CLS}.json').read_text())['trees']
APL = json.loads((CUR / f'apl/{SPEC}.json').read_text())
SPECS = json.loads((CUR / 'specs.json').read_text())
CLASSES = json.loads((B / 'classes.json').read_text())
CLASS_NAME = {c['slug']: c['name'] for c in CLASSES}
SPEC_ROW = next(s for s in SPECS if s['spec'] == SPEC)
SPEC_LABEL = f"{SPEC_ROW['name']} {CLASS_NAME[CLS]}"
RACIAL_ICON = json.loads((B / 'racialicons.json').read_text())
TALENT_ICON = {t['name']: t['icon'] for tr in TALENTS for t in tr['talents']}
WEIGHTS = next(w['weights'] for w in json.loads((B / 'stat-weights.json').read_text()) if w['spec'] == SPEC)
ITEM_BY_SPELL = {249469: 1249110}  # Frozen Heart of the Mountain and its use spell (engine-spell-names.json says so)
QUALITY = {0: '#9d9d9d', 1: '#ffffff', 2: '#1eff00', 3: '#3d94f0', 4: '#b866f5', 5: '#ff8000'}  # AA-safe text colours

YOU = ('Thandrel', CLS, f'{SPEC_LABEL} &middot; {CHAR["level"]}')     # SAMPLE signed-in character
LAST_SIM = (899.6, '3 Nov', 3000)                                       # SAMPLE: the previous saved run (DPS, date, runs)
LAST_SIM_CLOSE = (903.5, '3 Nov', 3000)                                 # SAMPLE: a previous run inside the noise (states sheet)
LIVE = (905.3, 1357)                                                    # SAMPLE: the live estimate and runs done mid-run
HANDS = {'other:attack/1': ('Main hand', 'main_hand'), 'other:attack/2': ('Off hand', 'off_hand'), 'other:attack/3': ('Extra attacks', 'main_hand')}
OFF_HAND_SPELLS = {462891}   # Whirlwind's off-hand strike, as lib/report/ability-table.ts and lib/sim/timeline.ts (Forever's
                             # Whirlwind hits with both weapons: spelltips/warrior.json 1680)
# SAMPLE rows of GET /v1/sims?mine=1, the three the parity gate stubs (web/tests/e2e/support/sim-parity.ts), drawn as SimYourSims
# draws a run: the title, then the spec and the day, the DPS grouped to a whole number
SERVER_AGES = ['3 days ago', '5 days ago', '8 days ago']   # SAMPLE: the gate's runs finished 3, 1 Nov and 29 Oct, seen on 6 Nov
HISTORY = [('Raid buffs', SPEC_ROW['name'], '3 Nov', '900'), ('No buffs', SPEC_ROW['name'], '1 Nov', '402'), ('Raid buffs', SPEC_ROW['name'], '29 Oct', '881')]

_icache: dict[str, str] = {}


# Round 5 (lane P6): every section on the simulator is a Section panel, the heading inside its Plate (web ui/Section `panel`),
# and the split page sets them as panels (ui/Page `panels`). The built components pass `panel` on every Section here.
_Section = Section
_Page = Page


def Section(*args, **kw) -> str:
    kw.setdefault('panel', True)
    return _Section(*args, **kw)


def Page(layout: str, main: str = '', rail: str = '', sticky_rail: bool = False) -> str:
    return _Page(layout, main, rail, sticky_rail, panels=layout == 'split')

# ------------------------------------------------------------------------------------------------ data helpers
def Button(label: str, primary: bool = False, small: bool = False) -> str:
    """The system Button as the build renders it, a <button> (ws_prims draws an <a>): the parity gate counts roles. Its case is the
    system's (decision 2026-11-06: no page-level case override)."""
    return _Button(label, primary, small).replace('<a ', '<button type="button" ', 1).replace(' href="#"', '', 1).replace('</a>', '</button>')


def Form(label: str, fields: list[tuple[str, str, bool]], submit: str) -> str:
    """ws_prims.FormRow as the build renders it (ui/FormRow, ui/Field): a named <form> whose fields are text inputs."""
    return FormRow(fields, submit).replace('<div class="formrow">', f'<form class="formrow" aria-label="{label}">', 1)[:-len('</div>')] + '</form>'


def Input(placeholder: str, label: str, grow: bool = False) -> str:
    """ws_prims.Field as a real <input> (the build's ui/Field): the parity gate counts a textbox."""
    return f'<input class="field{" grow" if grow else ""}" placeholder="{placeholder}" aria-label="{label}">'


def gicon(name, size: int = 28, q=None) -> str:
    """A real game icon from the build's icon set; the dark socket (never an invented icon) when the data has none."""
    cls = f'ic q{q}' if q else 'ic'
    p = B / 'icons' / f'{name}.webp' if name else None
    if p is None or not p.exists():
        return f'<span class="{cls}" style="width:{size}px;height:{size}px"></span>'
    if name not in _icache:
        _icache[name] = 'data:image/webp;base64,' + base64.b64encode(p.read_bytes()).decode()
    return f'<img class="{cls}" src="{_icache[name]}" alt="" width="{size}" height="{size}" style="width:{size}px;height:{size}px">'


def parse(key: str) -> tuple[str, int, int]:
    kind, rest = key.split(':', 1)
    ident, _, tag = rest.partition('/')
    return kind, int(ident) if ident.isdigit() else 0, int(tag or 0)


def spell_name(sid: int) -> str:
    return ENGINE_NAME.get(sid) or SPELL_NAME.get(sid) or 'An unnamed spell'


def spell_icon(sid: int):
    for k in (sid, ICON_ROWS.get(sid)):
        if k and str(k) in SPELL_ICON:
            return SPELL_ICON[str(k)]
    item = next((i for i, s in ITEM_BY_SPELL.items() if s == sid), None)
    if item and item in ITEMS:
        return ITEMS[item]['icon']
    by_name = next((e['icon'] for e in BUFFS.values() if e['name'] == spell_name(sid)), None)
    return by_name or RACIAL_ICON.get(spell_name(sid)) or TALENT_ICON.get(spell_name(sid))


def item_of(iid: int) -> dict:
    if iid in ITEMS:
        return ITEMS[iid]
    c = CONSUMES.get(iid, {})
    icon = next((e['icon'] for e in BUFFS.values() if e['name'] == c.get('name')), None)
    return {'name': c.get('name', f'Item {iid}'), 'icon': icon, 'quality': c.get('quality', 1)}


def act_name(key: str) -> str:
    if key == 'other:refund':
        return 'Refunded on a miss, dodge or parry'   # the shipped label (lib/sim/action-names)
    kind, ident, _ = parse(key)
    return item_of(ident)['name'] if kind == 'item' else spell_name(ident)


def display_name(key: str) -> str:
    """The row and lane name: the action's name, with the shipped "(off hand)" suffix on an off-hand strike."""
    if key in HANDS:
        return HANDS[key][0]
    kind, ident, _ = parse(key)
    return f'{spell_name(ident)} (off hand)' if kind == 'spell' and ident in OFF_HAND_SPELLS else act_name(key)


def lane_icon(key: str, size: int) -> str:
    """A swing lane shows its weapon; every other lane its action's icon."""
    if key in HANDS:
        it = gear_item(HANDS[key][1])
        return gicon(it['icon'], size, it['quality'])
    return act_icon(key, size)


def act_icon(key: str, size: int = 28) -> str:
    kind, ident, _ = parse(key)
    if kind == 'other':
        return gicon(None, size)
    if kind == 'item':
        it = item_of(ident)
        return gicon(it.get('icon'), size, it.get('quality'))
    return gicon(spell_icon(ident), size)


def rank_of(sid: int) -> int:
    return (TIPS.get(sid) or {}).get('rank') or 0


def fnum(v: float, d: int = 1) -> str:
    return f'{v:,.{d}f}'


def clock(sec: float) -> str:
    s = int(sec)
    return f'{s // 60}:{s % 60:02d}'


def gear_item(slot: str):
    g = next((x for x in CHAR['gear'] if x['slot'] == slot), None)
    return ITEMS.get(g['item_id']) if g else None


DUR = SUM['duration_ms'] / 1000
CASTS = {c['spell_name']: c['succeeded'] for c in SUM['casts']}
AURAS = {a['name']: a for a in SUM['auras']}
MEAN, ERR, SIGMA = RES['dps']['mean'], RES['dps']['error'], RES['dps']['stddev']
RUNS = RES['iterations_run']
Z95 = 1.96                            # ruling S2: one "±" everywhere, the 95% confidence band (lib/sim/estimate.ts)


def band(runs: int) -> float:
    """The 95% band of a mean over `runs` runs of this set: 1.96 x stddev / sqrt(runs)."""
    return Z95 * SIGMA / runs ** .5


BAND = Z95 * ERR
FIGURE = fnum(MEAN)
ERR_LINE = f"<b>&plusmn; {BAND:.1f} DPS</b>"   # the 95% band, in DPS, unlabelled (2026-11-09 head contract)
PRECISION = {500: 'Fast', 3000: 'Normal', 10000: 'High'}[REQ['iterations']]          # lib/sim/precision.ts PRECISION_ITERATIONS
PRECISION_LABEL = f"{PRECISION}, {REQ['iterations']:,} runs"
LANE_LABEL = ''   # the one head contract (2026-11-09): a finished result names no lane; the label's row is held
DAMAGE_KEYS = {a['name'] for a in SUM['damage_done'][0]['abilities'] if a['total'] > 0}
DAMAGE_SCHOOLS = {'physical' if (a.get('school') or 1) & 1 else 'other' for a in SUM['damage_done'][0]['abilities'] if a['total'] > 0}
SPLIT = '/'.join(str(sum(int(c) for c in part)) for part in CHAR['talents'].split('-'))
RACE = CHAR['race'].replace('-', ' ').title()
FIGHT = REQ['encounter']
FIGHT_LEN = f"{clock(FIGHT['duration_sec'])} &plusmn; {FIGHT['variation'] * 100:.0f}%"


# ------------------------------------------------------------------------------------------------ the ability table (the Plate)
TABLE_HEADS = [('', False), ('Ability', False), ('Share of damage', False), ('DPS', True), ('Share', True), ('Casts', True),
               ('Avg hit', True), ('Crit', True)]
N_FIGS = len(TABLE_HEADS) - 3
SORTABLE = tuple(t for t, _ in TABLE_HEADS[3:])   # the shipped table sorts by every figure column (heads are buttons)
MULTI = re.compile(r'up to (\d+) enemies')   # a tooltip that hits several enemies earns "hits a cast" on its row


def ability_rows(state: str = '') -> tuple[list[str], str]:
    rows = sorted((a for a in SUM['damage_done'][0]['abilities'] if a['total'] > 0), key=lambda a: -a['total'])
    total = sum(a['total'] for a in rows)
    top = rows[0]['total']
    hs = CASTS.get('spell:11567', 0)
    out = []
    for i, a in enumerate(rows):
        kind, ident, tag = parse(a['name'])
        m = a.get('misses') or {}
        attempts = a['hits'] + m.get('MISS', 0) + m.get('DODGE', 0) + m.get('PARRY', 0)
        periodic = a['ticks'] > 0
        if kind == 'other':
            wpn = gear_item({1: 'main_hand', 2: 'off_hand', 3: 'main_hand'}[tag])
            label = {1: 'Main hand', 2: 'Off hand', 3: 'Extra attacks'}[tag]
            icon = gicon(wpn['icon'], 28, wpn['quality'])
            name = f'<b>{label}</b> <span class="wpn" style="color:{QUALITY[wpn["quality"]]}">&middot; {wpn["name"]}</span>'
        else:
            icon = act_icon(a['name'])
            nm = display_name(a['name'])
            rank = rank_of(ident)
            name = f'<b>{nm}</b>' + (f'<span class="rk">Rank {rank}</span>' if rank else '')
        if periodic:
            casts, avg, crit = '&ndash;', a['total'] / a['ticks'], a['crits'] / a['ticks'] * 100
            deb = AURAS.get(a['name'])
            sub = [f'{a["ticks"]} ticks', f'tick crit {crit:.1f}%'] + ([f"up {deb['uptime_ms'] / SUM['duration_ms'] * 100:.1f}% of the fight"] if deb else [])
        else:
            casts = attempts if kind == 'other' else (CASTS.get(a['name']) or attempts)   # a proc strike with no cast row counts its attempts (ability-table.ts)
            avg = a['total'] / a['hits'] if a['hits'] else 0
            crit = a['crits'] / a['hits'] * 100 if a['hits'] else 0
            sub = [f'{lab} {fixed1(m[k] / attempts * 100)}%' for k, lab in (('MISS', 'Miss'), ('DODGE', 'Dodge'), ('PARRY', 'Parry'), ('GLANCING', 'Glance'), ('BLOCK', 'Block')) if m.get(k)]
            if kind == 'other' and tag == 1 and hs:
                sub.append(f'{hs} swings became Heroic Strike')
            if kind == 'other' and tag == 3:
                src = ' &middot; '.join(f"{BUFFS[s['source']]['name'] if s['source'] in BUFFS else s['source']} {s['count']:.0f}" for s in RES.get('extra_attack_sources', []))
                sub.append(f'From {src}')
            cast_n = CASTS.get(a['name'])
            if kind == 'spell' and cast_n and MULTI.search((TIPS.get(ident) or {}).get('description', '')):
                sub.append(f'{attempts / cast_n:.1f} hits a cast, {REQ["encounter"]["targets"]} target')
        phone = [x for x in (f'{casts} {"swings" if kind == "other" else "casts"}' if casts != '&ndash;' else '', f'{avg:,.0f} avg hit', f'crit {crit:.1f}%') if x]
        sub_html = ' &middot; '.join(f'<span>{s}</span>' for s in sub)
        sub_html += ''.join(f' <span class="m">{"&middot; " if (sub or j) else ""}{p}</span>' for j, p in enumerate(phone))
        share = f'{a["total"] / total * 100:.1f}%'
        out.append(StatRow(icon, name, sub_html, [fnum(a['total'] / DUR), share, str(casts), f'{avg:,.0f}', f'{crit:.1f}%'],
                           Meter(a['total'] / top), share, state='hover' if state == 'hover' and i == 2 else ''))
    out.append(StatRow('', f'Total, {len(rows)} sources &middot; {total:,.0f} damage a fight', '', [fnum(total / DUR), '100%', '', '', ''], total=True))
    return out, f'{fnum(MEAN)} DPS over {clock(DUR)} on average'


def ability_table(state: str = '') -> str:
    return StatTable(TABLE_HEADS, ability_rows(state)[0], N_FIGS, sort='DPS', sortable=SORTABLE)


def table_skeleton(n: int = 11) -> str:
    row = StatRow(Skeleton(28, 28), Skeleton(160, 16), Skeleton(200, 12), [Skeleton(56, 16)] + [Skeleton(44, 14)] * (N_FIGS - 1), Skeleton(120, 8))
    return StatTable(TABLE_HEADS, [row] * n, N_FIGS, sort='DPS', sortable=SORTABLE)


# ------------------------------------------------------------------------------------------------ the rail
def spread_section() -> str:
    h = RES['dps_histogram']
    counts, start, bw = h['counts'], h['start'], h['bin_dps']
    total, acc, lo, hi = sum(counts), 0, None, None
    for i, c in enumerate(counts):
        acc += c
        if lo is None and acc >= total * .16:
            lo = start + i * bw
        if hi is None and acc >= total * .84:
            hi = start + (i + 1) * bw
    d = RES['dps']
    hist = Histogram(counts, start, bw, MEAN, lo, hi, f'DPS of each of {RES["iterations_run"]:,} runs, from {d["min"]:,.0f} to {d["max"]:,.0f}, most near {MEAN:,.0f}',
                     (f'{d["min"]:,.0f}', 'DPS per run', f'{d["max"]:,.0f}'))
    figs = Figures([(fnum(MEAN), 'Mean', '', '', 'sm'), (f'&plusmn; {BAND:.1f}', 'Error', '', '', 'sm'), (f'{d["stddev"]:.1f}', 'Std dev', '', '', 'sm'),
                    (f'{d["min"]:,.0f}', 'Lowest', '', '', 'sm'), (f'{d["max"]:,.0f}', 'Highest', '', '', 'sm'), (f'{RES["iterations_run"]:,}', 'Runs', '', '', 'sm')])
    return Section('DPS spread', f'<div class="rail-body">{hist}{figs}{benchmark_row()}</div>')


# The run beside its spec's published best-in-slot set (web lib/sim/bis-benchmark.ts): the band at the run's level, faction and
# preset in data/builds/<build>/bis/<spec>.json, drawn only when the run fought the file's own `run` block. The gap carries both
# runs' 95% band combined; at 10 or more it is whole (lib/sim/combos.ts gainLabel).
FACTION = {'troll': 'horde', 'orc': 'horde', 'tauren': 'horde', 'undead': 'horde', 'human': 'alliance', 'dwarf': 'alliance',
           'gnome': 'alliance', 'night-elf': 'alliance'}


def benchmark_row() -> str:
    f = json.loads((B / 'bis' / f'{SPEC}.json').read_text())
    run, enc = f['run'], REQ['encounter']
    if (enc['duration_sec'], enc['variation'], enc['targets'], enc['execute_ratio']) != (run['duration_sec'], run['variation'], run['targets'], run['execute_ratio']):
        return ''
    band_row = next(b for b in f['bands'] if b['band'] == CHAR['level'] and b['faction'] == FACTION[CHAR['race']] and b.get('preset', 'bare') == REQ['preset'])
    dps, err = round(band_row['set_dps'], 2), round(band_row.get('set_dps_error', 0), 2)
    gap, margin = MEAN - dps, Z95 * (ERR ** 2 + err ** 2) ** .5
    mag = round(abs(gap), 1)
    gap_txt = f'{round(mag):,}' if mag >= 10 or mag == 0 else f'{mag:.1f}'
    # "1,004.0 · 9.9% under": the set's DPS and the run against it, once (DpsDistribution).
    trail = f'<span class="num"><b>{fnum(dps)}</b> <span class="t-small">&middot; {abs(gap) / dps * 100:.1f}% {"under" if gap < 0 else "over"}</span></span>'
    return Rows([Row(f'Level {band_row["band"]} {REQ["preset"]} BiS', '', trail=trail)], open_=True)


# "No effect on your damage": the shipped rule, lib/sim/buff-effect.ts hasNoDamageEffect, read from the data (ruling M4).
# A buff is dimmed when nothing it changes can reach this spec's damage: it works through no weapon swing, raises no school
# this run dealt damage in (school_damage against the run's schools), and every stat it adds (simbuffs.json `stats`, measured
# on a level 60 sheet) weighs nothing for the spec (stat-weights.json) - target armor and melee haste count only when the run
# dealt physical damage. An entry with no effect fields says nothing, so nothing is dimmed. Consumables are never judged.
CASTER_STATS = {'spell_crit', 'spell_damage'}
BUFF_BY_NAME = {e['name']: e for e in BUFFS.values()}


def stat_counts(stat: str) -> bool:
    if stat in ('target_armor', 'melee_haste'):
        return 'physical' in DAMAGE_SCHOOLS
    if stat in CASTER_STATS:
        return WEIGHTS.get('spell_power', 0) > 0
    return WEIGHTS.get(stat, 0) > 0


def no_damage_effect(e: dict) -> bool:
    if e.get('target') is None or e.get('melee_effect'):
        return False
    if set(e.get('school_damage') or []) & DAMAGE_SCHOOLS:
        return False
    return not any(stat_counts(st) for st in (e.get('stats') or {}))


def buff_items(keys: list[str]) -> list[tuple[str, str, bool]]:
    out = []
    for k in keys:
        base, _, grade = k.partition(':')
        e = BUFFS.get(base)
        if not e:
            continue
        out.append((gicon(e['icon'], 22), e['name'] + (', improved' if grade == 'improved' else ''), no_damage_effect(e)))
    return sorted(out, key=lambda x: x[2])


def consume_items() -> list[tuple[str, str, bool]]:
    """Every consumable the request carries, one entry each (an imbue's simbuffs key is the whole `hand_imbue:item` id)."""
    out = []
    for k in CHAR['consumes']:
        hand = k.partition(':')[0] if ':' in k else ''
        e = BUFFS.get(k) or BUFFS.get(k.partition(':')[2])
        assert e, f'consumable {k} has no simbuffs entry'
        out.append((gicon(e['icon'], 22), e['name'] + {'main_hand_imbue': ' (main hand)', 'off_hand_imbue': ' (off hand)'}.get(hand, ''), False))
    return out


def fixed1(x: float) -> str:
    """JavaScript's toFixed(1), which the build formats with: half away from zero on the double's exact value (6.25 -> 6.3)."""
    return str(decimal.Decimal(x).quantize(decimal.Decimal('0.1'), rounding=decimal.ROUND_HALF_UP))


def apm() -> float:
    return sum(v for k, v in CASTS.items() if not k.startswith('other:') and '/' not in k) / (DUR / 60)


def raid_keys() -> list[str]:
    presets = json.loads((CUR / 'presets.json').read_text())
    for p in presets.get('presets', []) if isinstance(presets.get('presets'), list) else []:
        if p.get('id') == 'raid' and p.get('buffs'):
            return p['buffs']
    return CHAR['buffs']


def this_run_section() -> str:
    ts = RES.get('target_stats') or {}
    lo, hi = FIGHT['duration_sec'] * (1 - FIGHT['variation']), FIGHT['duration_sec'] * (1 + FIGHT['variation'])
    facts = Facts([('Fight length', f"{clock(FIGHT['duration_sec'])}, &plusmn; {FIGHT['variation'] * 100:.0f}%<small>{clock(lo)} to {clock(hi)}</small>"),
                   ('Target', f"Level {ts.get('level', 63)} boss<small>{ts['armor']:,} armor, {ts['armor_after_debuffs']:,} after debuffs</small>" if ts else 'Level 63 boss'),
                   ('Execute phase', f"Last {FIGHT['execute_ratio'] * 100:.0f}% of the fight"),
                   ('Race', RACE), ('Runs', f"{RUNS:,}"), ('Actions per minute', f'{apm():.1f}'),
                   ('Engine', f'<span class="muted">{RES["engine_version"][:7]}</span>')])
    keys = raid_keys()
    on_you = [k for k in keys if not (BUFFS.get(k.split(':')[0]) or {}).get('target')]
    on_tgt = [k for k in keys if (BUFFS.get(k.split(':')[0]) or {}).get('target')]
    cons = consume_items()
    you_l, tgt_l = buff_items(on_you), buff_items(on_tgt)
    lists = (IconList(you_l, f'On you &middot; {len(you_l)}') + IconList(tgt_l, f'On the target &middot; {len(tgt_l)}')
             + IconList(cons, f'Consumables &middot; {len(cons)}'))
    acts = f'<div class="rail-acts">{Button("Copy for the game", small=True)}{Button("Copy JSON", small=True)}{Button("Download JSON", small=True)}</div>'
    return Section('This run', f'<div class="rail-body flow">{facts}{lists}{acts}</div>')


def rot_title(key: str, name: str, rank: int = 0) -> str:
    """The built rotation title (ActionMark): the 24 px icon, then the name and its rank."""
    rk = f'<span class="rk-s">Rank {rank}</span>' if rank else ''   # glued, as the built ActionMark: the rank never takes a line alone
    icon = act_icon(key, 24) if key else gicon('inv_misc_pocketwatch_01', 24)   # Other cooldowns: OTHER_COOLDOWNS_ICON
    return f'<span class="am">{icon}<span>{name}{rk}</span></span>'


def rotation_section(ran: bool = True) -> str:
    """The shipped rows (lib/sim/rotation-card.ts): a prepull line by its time, a priority line by its number; the icon sits
    in the title; Other cooldowns carries the first one's icon and their summed casts."""
    rot = APL['rotation']
    notes = {n['step']: n['notes'] for n in APL.get('level_notes', []) if n['from_level'] <= CHAR['level']}
    sample_pre = {e['action'] for e in RES['sample'] if e['at_ms'] < 0}
    pre = []
    for p in rot['prepullActions']:
        sid = p['action']['castSpell']['spellId']['spellId']
        key = f'spell:{sid}'
        cast = key in sample_pre
        r = Row(rot_title(key, spell_name(sid), rank_of(sid)), '', lead=f'<span class="pre-t num">{p["doAtValue"]["const"]["val"].replace("s", " s")}</span>',
                trail=(f'<span class="cnt">{"Cast" if cast else "Not cast in this run"}</span>' if ran else ''), link=False)
        if cast or not ran:   # after a run, a step it never cast is left out (SimRotationSection)
            pre.append(r)
    named = {f"spell:{(p['action'].get('castSpell') or {}).get('spellId', {}).get('spellId')}" for p in rot['priorityList']}
    named |= {f"spell:{p['action']['castSpell']['spellId']['spellId']}" for p in rot['prepullActions']}
    other = [k for k, v in CASTS.items() if v and not k.startswith('other:') and '/' not in k and k not in named]
    lines = []
    for i, p in enumerate(rot['priorityList']):
        a = p['action']
        if 'autocastOtherCooldowns' in a:
            total = round(sum(CASTS[k] for k in other))
            sub = f'This run: {", ".join(act_name(k) for k in other)}.' if ran else ''
            cnt = (f'{total} cast{"s" if total != 1 else ""}' if total else 'Not cast in this run')
            r = Row(rot_title(other[0] if ran and other else '', 'Other cooldowns as they come up'), sub, rank=str(len(lines) + 1),
                    trail=(f'<span class="cnt">{cnt}</span>' if ran else ''), link=False)
            if total or not ran:
                lines.append(r)
            continue
        sid = a['castSpell']['spellId']['spellId']
        key = f'spell:{sid}'
        n = CASTS.get(key, 0)
        cnt = f'{n} cast{"s" if n != 1 else ""}' if n else 'Not cast in this run'
        note = notes.get(i, p.get('notes', ''))
        if re.search(r'^(?:on cooldown|whenever)\b|\bon cooldown\b', note, re.I):   # rotation-card.ts obviousNote
            note = ''
        r = Row(rot_title(key, spell_name(sid), rank_of(sid)), note,
                rank=str(len(lines) + 1), trail=(f'<span class="cnt">{cnt}</span>' if ran else ''), link=False)
        if n or not ran:
            lines.append(r)
    body = (f'<div class="rot"><h3 class="t-label rot-l">Before the pull</h3>{Rows(pre, open_=True)}<h3 class="t-label rot-l">Then, from the top at every decision</h3>'
            f'{Rows(lines, open_=True)}</div>')
    return Section('Rotation', body, Link(f'{SPEC_ROW["name"]} guide &rsaquo;'))


def history_section(state: str = 'ready', server: bool = False) -> str:
    """Your sims. `server`: a payer's list, where server runs are rows of the same list (no panel of their own): a titled run
    names its spec and age (SimYourSims + ServerRunRow); the gate's three runs are server runs."""
    if state == 'empty':
        return Section('Your sims', '<p class="empty-open">Nothing saved yet. Run a sim and press Save.</p>')
    if state == 'failed':
        return Section('Your sims', f'<p class="empty-open">Your sims did not load. {Link("Try again")}</p>')
    if server:
        rows = [Row(t, f'<time>{a}</time>', trail=(Figure(f'{v} DPS', size='sm') if v else ''))
                for (t, _, _, v), a in zip(HISTORY, SERVER_AGES)]
        return Section('Your sims', Rows(rows, open_=True), Link('All &rsaquo;'))
    rows = [Row(t, d if s == SPEC_LABEL else f'{s} &middot; {d}', trail=(Figure(v, unit='DPS', size='sm') if v else '')) for t, s, d, v in HISTORY]
    return Section('Your sims', Rows(rows, open_=True), Link('All &rsaquo;'))


# ------------------------------------------------------------------------------------------------ main sections below the Plate
def timeline_section(swings_and_ticks: bool = False) -> str:
    """The median run on one time axis. With Swings and ticks Off (the default) the Casts group holds casts only: no white
    swing lane and no periodic tick lane (an action whose events carry outcome "tick", Deep Wounds here)."""
    T0, T1 = -2500, RES['sample_duration_ms']
    X = lambda ms: (ms - T0) / (T1 - T0) * 1000
    ev = RES['sample']
    ticking = {e['action'] for e in ev if e.get('outcome') == 'tick'}
    cast_keys = []
    for e in ev:
        k = e['action']
        if k in cast_keys or '/' in k and not k.startswith('other:'):
            continue
        if not swings_and_ticks and (k.startswith('other:') or k in ticking):
            continue
        if not swings_and_ticks and CASTS and k not in CASTS and k in DAMAGE_KEYS:
            continue   # a hit no one pressed (the off-hand Whirlwind) lands on Swings and ticks (lib/sim/timeline-lanes.ts placeEvent)
        cast_keys.append(k)
    whole = lambda a: len(a['segments']) == 1 and a['segments'][0][0] <= 0 and a['segments'][0][1] >= T1 * .99
    auras = [a for a in RES['sample_auras'] if a['on'] == 'player' and not whole(a) and a['action'] not in ('spell:11567/1', 'spell:11585')]
    tgt = [a for a in RES['sample_auras'] if a['on'] == 'target' and not whole(a)]
    exe = next((e['at_ms'] for e in ev if e['action'] == 'spell:20662'), None)
    LANE, HEAD, RESH = 24, 30, 72
    rows, labels, y = [], [], 0

    def head(t):
        nonlocal y
        labels.append(f'<span class="g" style="height:{HEAD}px">{t}</span>')
        y += HEAD

    def lane(key, body):
        nonlocal y
        labels.append(f'<span style="height:{LANE}px">{lane_icon(key, 18)}<span class="t">{display_name(key)}</span></span>')
        rows.append(f'<line x1="0" x2="1000" y1="{y + LANE - .5}" y2="{y + LANE - .5}" stroke="#1c2230" vector-effect="non-scaling-stroke"/>' + body(y))
        y += LANE

    head('Casts')
    for k in cast_keys:
        def body(y0, k=k):
            out = ''
            for e in ev:
                if e['action'] != k:
                    continue
                crit = e.get('outcome') == 'crit'
                miss = e.get('outcome') in ('miss', 'dodge', 'parry')
                h = 18 if crit else 14
                fill = '#f5d27a' if crit else ('none' if miss else ('#c79c6e' if (e.get('amount') or 0) > 0 else '#e9e4d8'))
                stroke = ' stroke="#9a9484" vector-effect="non-scaling-stroke"' if miss else ''
                out += f'<rect x="{X(e["at_ms"]) - 1.5:.1f}" y="{y0 + (LANE - h) / 2}" width="3" height="{h}" fill="{fill}"{stroke}/>'
            return out
        lane(k, body)
    for title, group, col in (('On you', auras, '#e9e4d859'), ('On the target', tgt, '#8a6a2ac7')):
        if group:
            head(title)
            for a in group:
                lane(a['action'], lambda y0, a=a, col=col: ''.join(f'<rect x="{X(s):.1f}" y="{y0 + 7}" width="{max(X(e) - X(s), 1.5):.1f}" height="10" rx="3" fill="{col}"/>' for s, e in a['segments']))
    pts, last = [], 0
    for e in ev:
        r = (e.get('resources') or {}).get('rage')
        if r is None:
            continue
        pts += [(X(e['at_ms']), last), (X(e['at_ms']), r)]
        last = r
    pts.append((1000, last))
    peak = max(p[1] for p in pts)
    path = ' '.join(f'{x:.1f},{y + RESH - r / 100 * (RESH - 8):.1f}' for x, r in pts)
    labels.append(f'<span class="res" style="height:{RESH}px"><b>Rage</b><small>peak {peak:.0f} of 100</small></span>')
    rows.append(f'<polyline points="{path}" fill="none" stroke="#e5b955" stroke-width="1.5" vector-effect="non-scaling-stroke"/>'
                f'<line x1="0" x2="1000" y1="{y + 8}" y2="{y + 8}" stroke="#3a3326" stroke-dasharray="4 3" vector-effect="non-scaling-stroke"/>')
    y += RESH
    secs = range(0, int(T1 / 1000) + 1, 15)
    grid = ''.join(f'<line x1="{X(t * 1000):.1f}" x2="{X(t * 1000):.1f}" y1="0" y2="{y}" stroke="#1c2230" vector-effect="non-scaling-stroke"/>' for t in secs)
    wash = f'<rect x="{X(exe):.1f}" y="0" width="{1000 - X(exe):.1f}" height="{y}" fill="#d66e2812"/>' if exe else ''
    pull = f'<line x1="{X(0):.1f}" x2="{X(0):.1f}" y1="0" y2="{y}" stroke="#e5b955" stroke-dasharray="3 3" vector-effect="non-scaling-stroke"/>'
    svg = f'<svg viewBox="0 0 1000 {y}" preserveAspectRatio="none" aria-hidden="true" style="height:{y}px">{wash}{grid}{pull}{"".join(rows)}</svg>'
    ticks = ''.join(f'<span class="{"pull" if t == 0 else ""}" style="left:{X(t * 1000) / 10:.2f}%">{"Pull" if t == 0 else clock(t)}</span>' for t in secs)
    washlab = f'<span class="wash" style="left:{X(exe) / 10:.2f}%">EXECUTE PHASE</span>' if exe else ''
    tl = (f'<div class="tl"><div class="lbl"><span style="height:30px"></span>{"".join(labels)}</div>'
          f'<div class="plot" role="application" aria-label="Fight timeline, {clock(T1 / 1000)} long. Arrow keys move between casts.">'
          f'<div class="axis">{ticks}</div>{washlab}{svg}</div></div>')
    # the shipped legend's swatches at each mark's size (SimTimeline: a 3 x 12 cast, a 4 x 16 crit, 16 x 8 auras, a 16 x 2 line)
    legend = [('#c79c6e', 'Damage', 3, 12), ('#f5d27a', 'Crit', 4, 16), ('#9a9484', 'Missed', 3, 12), ('#e9e4d8', 'Cooldown', 3, 12),
              ('#e9e4d859', 'On you', 16, 8), ('#8a6a2ac7', 'On the target', 16, 8), ('#e5b955', 'Rage', 16, 2)]
    tools = (f'<div class="tl-tools">{seg([("Whole fight", True, None, ""), ("Close up", False, None, "")], "sm")}'
             f'<span class="tl-sw"><span class="t-small">Swings and ticks</span>{seg([("Off", not swings_and_ticks, None, ""), ("On", swings_and_ticks, None, "")], "sm")}</span>'
             f'{LanesLegend(legend)}</div>')
    brush = '<div class="tl-brush" aria-hidden="true"><span></span><span class="tr"></span></div>'   # the shipped TimelineBrush strip
    n = sum(1 for e in ev if e['action'] in cast_keys and e['action'] not in HANDS and e['action'] not in ticking and parse(e['action'])[1] not in OFF_HAND_SPELLS)
    foot = (f'<p class="t-small tl-foot">{n} casts in this run. '
            + '<button type="button" class="tl-list" aria-expanded="false">List every cast &rsaquo;</button></p><p class="sr" aria-live="polite"></p>')
    return Section('Timeline', tools + Lanes(tl) + brush + foot)


DOLL_LEFT = ['head', 'neck', 'shoulder', 'back', 'chest', 'wrist', 'main_hand', 'off_hand', 'ranged']
DOLL_RIGHT = ['hands', 'waist', 'legs', 'feet', 'finger1', 'finger2', 'trinket1', 'trinket2']
SLOT_LABEL = {'head': 'Head', 'neck': 'Neck', 'shoulder': 'Shoulder', 'back': 'Back', 'chest': 'Chest', 'wrist': 'Wrist', 'main_hand': 'Main hand',
              'off_hand': 'Off hand', 'ranged': 'Ranged', 'hands': 'Hands', 'waist': 'Waist', 'legs': 'Legs', 'feet': 'Feet', 'finger1': 'Finger 1',
              'finger2': 'Finger 2', 'trinket1': 'Trinket 1', 'trinket2': 'Trinket 2'}


def enchant_names() -> dict[int, str]:
    """The client's name per enchant id (lib/sim/enchant-names.ts): a row without `display` draws no line."""
    path = B / 'enchants.json'
    rows = json.loads(path.read_text()) if path.exists() else []
    return {r['id']: r['display'] for r in rows if r.get('display')}


RAIL_BAND, RAIL_FACTION, RAIL_PRESETS = 60, 'alliance', ('raid', 'bare')   # lib/sim/bis-set-rail.ts, lib/bis/presets.ts selectBand


def rail_gear() -> list[dict]:
    """The landing's character: the rail's published set (the level 60 alliance band, raid preset first) with its enchants, the
    set the "Sim a best-in-slot set" link's code carries (lib/bis/sim-link.ts bandGearSlots)."""
    bands = [b for b in json.loads((B / 'bis' / f'{SPEC}.json').read_text())['bands']
             if b['band'] == RAIL_BAND and b['faction'] == RAIL_FACTION]
    band = next((b for p in RAIL_PRESETS for b in bands if b.get('preset', 'bare') == p), bands[0])
    return [{'slot': s['slot'], 'item_id': s['item_id'], 'enchant': s.get('enchant_id')} for s in band['slots'] if s['item_id'] > 0]


# The frame opens on the first worn item in this order (lib/ui/paperdoll.ts OPEN_ORDER: left column, weapons, right column).
DOLL_OPEN_ORDER = ['head', 'neck', 'shoulder', 'back', 'chest', 'shirt', 'tabard', 'wrist', 'main_hand', 'off_hand', 'ranged',
                   'hands', 'waist', 'legs', 'feet', 'finger1', 'finger2', 'trinket1', 'trinket2']
QUALITY_NAME = {0: 'poor', 1: 'common', 2: 'uncommon', 3: 'rare', 4: 'epic', 5: 'legendary'}


def doll(gear: list[dict] | None = None) -> str:
    """The sim's Gear (SimGearSection on character/GearDoll): the shared PaperdollFrame in `worn` mode, as BiS, the character
    page and the reports draw it (owner ruling 2026-11-11: one gear look), capped at the character page's 640 px. Headless
    (the page's head names the character); the title bar is the spec and the level; the frame opens on the first worn item,
    its client tooltip in the centre. Under 423 px the frame's own list form: `facts` rows "Slot · Item level N" with the
    client enchant green under it (lib/gear/doll-cells.ts; no gain without an upgrade). In a Plate on the landing, open under
    a Section on a result. `gear`: the worn rows, the run's character by default."""
    names = enchant_names()
    slots = []
    for row in CHAR['gear'] if gear is None else gear:
        it = ITEMS.get(row['item_id'])
        if not it:
            continue
        enchant = names.get(row.get('enchant') or 0)
        slots.append({'slot': row['slot'], 'label': f"{SLOT_LABEL[row['slot']]}: {it['name']}", 'covered': False, 'selected': False,
                      'icon': it['icon'], 'quality': it['quality'], 'name': it['name'], 'line2': f"Item level {it['item_level']}",
                      **({'enchant': enchant} if enchant else {})})
    by = {s['slot']: s for s in slots}
    first = next((by[n] for n in DOLL_OPEN_ORDER if n in by), None)
    centre = ''
    if first:
        first['selected'] = True
        centre = ItemTip(first['name'], QUALITY_NAME.get(first['quality'], 'common'),
                         [first['name'], first['line2']] + ([first['enchant']] if first.get('enchant') else []), [], '')
    frame = PaperdollFrame(slots, SPEC_LABEL, f"Level {CHAR['level']}", None, centre, 'Gear by slot', heading='h2', head=False, facts=True)
    return f'<div style="width:100%;max-width:640px;margin:0 auto">{frame}</div>'


def ilvl() -> float:
    lv = [gear_item(s)['item_level'] for s in DOLL_LEFT + DOLL_RIGHT if gear_item(s)]
    return sum(lv) / len(lv)




def gear_section(with_stats: bool = True) -> str:
    fs = RES['final_stats']
    miss = {'sp': [0, 0], 'w': [0, 0]}
    for a in SUM['damage_done'][0]['abilities']:
        m = a.get('misses') or {}
        att = a['hits'] + m.get('MISS', 0) + m.get('DODGE', 0) + m.get('PARRY', 0)
        if a['ticks'] or not att:
            continue
        b = miss['w' if a['name'].startswith('other:') else 'sp']
        b[0] += m.get('MISS', 0)
        b[1] += att
    stats = Figures([(f"{fs['attack_power']:,}", 'Attack power', '', '', 'sm'), (f"{fs['melee_hit_pct']:.1f}%", 'Hit', '', '', 'sm'),
                     (f"{fs['melee_crit_pct_vs_target']:.1f}%", 'Crit on the boss', '', '', 'sm'),
                     (f"{fs['weapon_skill']['main_hand']} &middot; {fs['weapon_skill']['off_hand']}", 'Weapon skill', '', '', 'sm'),
                     (f"{fs['melee_haste_pct']:.0f}%", 'Haste from gear', '', '', 'sm')], cols=2)
    measured = (f'<p class="t-small">Missed in this run: {miss["sp"][0] / miss["sp"][1] * 100:.0f}% of specials '
                f'&middot; {miss["w"][0] / miss["w"][1] * 100:.1f}% of white swings.</p>')
    # SimGearSection: the frame beside its stats from a 760 px panel, centred alone; the link is the action in every state.
    aside = f'<div class="gear-aside">{stats}{measured}</div>' if with_stats else ''
    body = f'<div class="gear-c"><div class="gear-grid{" pair" if aside else ""}">{doll()}{aside}</div></div>'
    return Section('Gear', body, Link('Upgrades on Best in slot &rsaquo;'))


def talents_section() -> str:
    trees = talent_tree_dicts(CLS, talent_ranks_from_strings(CLS, CHAR['talents'].split('-')))
    return Section('Talents', Trees(trees), Link('Open in planner &rsaquo;'))


def rage_section() -> str:
    r, src = SUM['resources'][0], RES['resource_sources'][0]
    figs = Figures([(f"{r['gained']:,}", 'Gained', '', '', 'sm'), (f"{r['spent']:,}", 'Spent', '', '', 'sm'),
                    (f"{r['wasted']:,} &middot; {r['wasted'] / r['gained'] * 100:.1f}%", 'Lost at the cap', '', '', 'sm')])

    def lst(rows, title):
        top = max(x['amount'] for x in rows)
        out = []
        for x in sorted(rows, key=lambda x: -x['amount']):
            k = x['action']
            nm = {'other:attack/1': 'Main hand', 'other:attack/2': 'Off hand', 'other:attack/3': 'Extra attacks'}.get(k) or act_name(k)
            ic = gicon(gear_item('off_hand' if k.endswith('/2') else 'main_hand')['icon'], 24) if k.startswith('other:') else act_icon(k, 24)
            lost = f'<span class="wipe">&minus;{x["wasted"]:.1f} lost</span>' if x.get('wasted') else ''
            out.append(StatRow(ic, f'<b>{nm}</b>', lost, [f'{x["amount"]:,.0f}'], Meter(x['amount'] / top)))
        return f'<div><h3 class="t-label rot-l">{title} &middot; {sum(x["amount"] for x in rows):,.0f}</h3>{StatTable([("", False)] * 3 + [("", True)], out, 1, open_=True)}</div>'
    return Section('Rage', figs + f'<div class="two">{lst(src["gained"], "Gained from")}{lst(src["spent"], "Spent on")}</div>')


def suppressed(key: str, requested: list[str]) -> bool:
    """A buff of an exclusive group that a lower rank in the same request holds off (lib/sim/aura-card.ts `isSuppressed`)."""
    own = (BUFFS.get(key.partition(':')[0]) or {}).get('exclusive_with')
    if not own:
        return False
    return any((t := (BUFFS.get(o.partition(':')[0]) or {}).get('exclusive_with')) and t['group'] == own['group'] and t['rank'] < own['rank']
               for o in requested)


def buffs_section() -> str:
    """The two columns as lib/sim/aura-card.ts files them: the engine's tracks, then every buff the request named that has no
    track (up all fight, build rulings R1 and R2 of spec 2026-11-04), the no-effect entries dimmed and last."""
    requested = CHAR['buffs']
    tracked = {act_name(a['name']) for a in SUM['auras']}
    untracked = [k for k in requested if (BUFFS.get(k.partition(':')[0]) or {}).get('name') not in tracked and not suppressed(k, requested)]

    def col(kind, title):
        rows, upall = [], []
        for a in SUM['auras']:
            if a['type'] != kind or a['name'] in ('spell:11567/1', 'spell:11585'):
                continue
            up = a['uptime_ms'] / SUM['duration_ms']
            if up >= .999:
                entry = BUFF_BY_NAME.get(act_name(a['name']))
                upall.append((act_icon(a['name'], 20), act_name(a['name']) + (f' &times; {a["max_stacks"]}' if (a.get('max_stacks') or 0) > 1 else ''),
                              bool(entry) and no_damage_effect(entry)))
            else:
                rows.append((up, StatRow(act_icon(a['name'], 24), f'<b>{display_name(a["name"])}</b>', '', [f'{up * 100:.1f}%', f'{a["applications"]:.0f}&times;'], Meter(up))))
        on_target = kind == 'DEBUFF'
        upall += [(gicon(e['icon'], 20), e['name'] + (', improved' if k.endswith(':improved') else ''), no_damage_effect(e))
                  for k in untracked if (e := BUFFS.get(k.partition(':')[0])) and bool(e.get('target')) == on_target]
        upall.sort(key=lambda x: x[2])
        rows.sort(key=lambda x: -x[0])
        t = StatTable([('', False)] * 3 + [('', True), ('', True)], [r for _, r in rows], 2, open_=True) if rows else ''
        come = '1 that comes and goes' if len(rows) == 1 else f'{len(rows)} that come and go'
        return f'<div><h3 class="t-label rot-l">{title} &middot; {come}</h3>{t}{IconList(upall, "Up all fight") if upall else ""}</div>'
    return Section('Buffs and debuffs', f'<div class="two">{col("BUFF", "On you")}{col("DEBUFF", "On the target")}</div>')


# ------------------------------------------------------------------------------------------------ head, bar, pages
TOOLS = ['Quick sim', 'Top gear', 'Droptimizer', 'Talents', 'Weights', 'Spec support']


def run_bar(state: str = 'ready', open_m: bool = False) -> str:
    """The run line. `state` picks the one primary: ready (Run sim), done or any failed/stopped state (Run again: the one retry),
    running (Stop), engine (Loading engine...), server (Running on our servers...), disabled (a spec the sim does not run)."""
    btn = {'ready': Button('Run sim', primary=True), 'done': Button('Run again', primary=True), 'retry': Button('Try again', primary=True),
           'running': Button('Stop'), 'engine': Button('Loading engine&hellip;'), 'server': Button('Running on our servers&hellip;'),
           'disabled': '<button type="button" class="btn primary is-off" aria-disabled="true">Run sim</button>'}[state]
    off = state in ('running', 'engine', 'server')
    buffs = seg([('Raid', True, None, ''), ('None', False, None, ''), ('Custom', False, None, '')])
    summary = (f'<span class="sum-t"><b>{FIGHT_STYLE} {clock(FIGHT["duration_sec"])}</b> &middot; Raid</span>'
               f'<span class="run-m">{btn}</span>')
    more = '<button type="button" class="link">More settings</button>'   # it opens the sheet: a button, as the build
    end = f'<span class="bar-end">{more}{btn}</span>'
    bar = control_bar([('Fight', sel(FIGHT_STYLE)), ('Length', sel(clock(FIGHT['duration_sec']))), ('Targets', sel(str(FIGHT['targets']))), ('Buffs', buffs),
                       ('Precision', sel(PRECISION_LABEL))], summary, end=end, open_m=open_m)
    bar = bar.replace('</div></details>', f'<span class="more-m">{more}</span></div></details>', 1)
    return bar.replace('<div class="bar-d">', '<div class="bar-d is-run">', 1) if off else bar


FIGHT_STYLE = 'Patchwerk'   # the request's profile is empty: the shipped styleLabel of the default style


def you_header(answer: str, actions: str = '') -> str:
    # The parity gate's character: a linked addon export (`who=`) with no capture date, the default source: no source line.
    meta = [f'{RACE} {SPEC_LABEL} &middot; Level {CHAR["level"]} &middot; {SPLIT} talents &middot; item level {ilvl():.1f}']
    return PageHeader('Thandrel', CLASS_COLOR[CLS], crest=crest_img(CLS, 56), meta=meta, answer=answer, actions=actions,
                      tabs=TabStrip(TOOLS, 'Quick sim'))


def since_last(last: tuple = LAST_SIM) -> str:
    """The change since the character's last sim, carrying the error of the gap: the two runs' 95% bands combined (root of the
    sum of squares). A gap inside that band reads "within the noise", muted, never green or red."""
    d = MEAN - last[0]
    b = (BAND ** 2 + band(last[2]) ** 2) ** .5
    sign = '+' if d >= 0 else '&minus;'
    if abs(d) <= b:
        return f'<span class="fig-cap">Within the noise of your last sim ({last[1]}): {sign}{abs(d):.1f} &plusmn; {b:.1f}</span>'
    return f'<span class="fig-cap {"kill" if d > 0 else "wipe"}">{sign}{abs(d):.1f} &plusmn; {b:.1f} DPS since your last sim ({last[1]})</span>'


def answer_result(signed_in: bool = True, last: tuple | None = None) -> str:
    """The live result. `last` draws the change since the character's last sim: hidden on the shipped page (gaps G1 and G15,
    no character key and no error on a saved-sim row), so only the states sheet passes it."""
    return (f'<span class="t-label">{LANE_LABEL}</span>{Figure(FIGURE, unit="DPS", tone="gold", size="big")}'
            f'<span class="fig-cap">{ERR_LINE}</span>{since_last(last) if signed_in and last else ""}')


def result_actions(signed_in: bool = True) -> str:
    """The result's actions, beside the answer on its last line (PageHeader `actions`)."""
    save = Button('Save this sim', small=True) if signed_in else Button('Sign in to save', small=True)
    return f'{save}{Link("Find upgrades in Top Gear &rsaquo;")}'


def answer_none() -> str:
    """Signed in, the character loaded, not run: the shipped answer while G1 is open (no saved sim can be matched to a character)."""
    return '<span class="t-label"></span>'


def answer_last() -> str:
    return (f'<span class="t-label">Your last sim &middot; {LAST_SIM[1]}</span>{Figure(fnum(LAST_SIM[0]), unit="DPS", size="big")}'
            f'<span class="fig-cap">Raid buffs, {FIGHT_STYLE} {clock(FIGHT["duration_sec"])}. Run sim for today&rsquo;s number.</span><span class="ans-acts">{Link("Open it &rsaquo;")}</span>')


def answer_running() -> str:
    est, done = LIVE
    return (f'<span class="t-label">Running in your browser</span>{Figure(f"{est:,.0f}", unit="DPS", tone="gold", size="big")}'
            f'<span class="fig-cap"><b>&plusmn; {band(done):.1f} DPS</b> &middot; {done:,} of {REQ["iterations"]:,} runs</span><span class="ans-meter">{Meter(done / REQ["iterations"])}</span>')


def answer_stopped() -> str:
    est, done = LIVE
    return (f'<span class="t-label">Stopped at {done:,} of {REQ["iterations"]:,} runs</span>{Figure(f"about {est:,.0f}", unit="DPS", tone="muted", size="big")}'
            f'<span class="fig-cap">Stopped at {done:,} runs: about {est:,.0f} &plusmn; {band(done):.1f}. A stopped run is not saved.</span>')


def answer_failed(label: str, reason: str, fix: str = '') -> str:
    """A failed run: what failed, why, and what the player does about it. The retry is the bar's one primary, never a second button here."""
    return f'<span class="t-label">{label}</span><span class="fig-cap wipe">{reason}</span>' + (f'<span class="fig-cap">{fix}</span>' if fix else '')


def result_page_body(saved: bool) -> str:
    main = Section('Damage by ability', ability_table(), frame='framed', flush=True) + timeline_section() + gear_section() + talents_section() + rage_section() + buffs_section()
    rail = spread_section() + ('' if saved else history_section(server=True)) + this_run_section() + rotation_section()
    return Page('split', main, rail)


PAGE_CSS = '''<style>
/* the page's own bottom space (sim.astro .sim-page, 40 px) plus Base.astro's [data-page-body] gap (40 px, 64 px from 1024) */
main{padding-bottom:80px}@media(min-width:1024px){main{padding-bottom:104px}}
/* page-level layout only: no border, background, radius or shadow (the guard reads inline styles; this block keeps to layout too) */
/* the answer slot keeps the result's size from the first state to the last (no layout shift): label, figure, error, change, actions */
.head .ans .t-label{min-height:16px}.head .ans .fig-cap{min-height:24px;display:flex;align-items:center;justify-content:inherit}
.statesheet .head .ans{min-height:0}
.fig-cap{font-size:13px;line-height:20px;color:var(--muted)}.fig-cap b{color:var(--text);font-weight:600}.fig-cap.kill{color:var(--kill);font-weight:600}.fig-cap.wipe{color:var(--wipe)}
.ans-acts{display:flex;gap:16px;align-items:center;flex-wrap:wrap;margin-top:6px}@media(min-width:900px){.ans-acts{justify-content:flex-end}}
.ans-meter{display:block;width:100%;max-width:360px;margin-top:6px}
.wk{display:inline-flex;gap:10px;align-items:center;color:var(--text)}
.rail-body{display:flex;flex-direction:column;gap:12px;padding-top:12px}.rail-cap{margin:0}.rail-acts{display:flex;flex-wrap:wrap;gap:8px;margin-top:8px}
.rail .figs{grid-template-columns:repeat(3,minmax(0,1fr))!important;gap:14px 16px}
.rot-l{padding:16px 0 4px}.dim{opacity:.5}.pre-t{width:40px;font-size:13px;color:var(--muted)}.cnt{font-size:13px;color:var(--muted);font-weight:600;white-space:nowrap}
.rk-s{font-size:13px;color:var(--muted);font-weight:400}
.st-h .st-sort{all:unset;cursor:pointer}
.tl-list{all:unset;cursor:pointer;color:var(--gold);font-weight:600}.sr{position:absolute;width:1px;height:1px;overflow:hidden;clip:rect(0 0 0 0)}
.lanes ul.lg{list-style:none;margin:0}
.two{display:grid;gap:var(--gap);grid-template-columns:minmax(0,1fr);padding-top:8px}@media(min-width:1024px){.two{grid-template-columns:repeat(2,minmax(0,1fr))}}
.layout-split .main>section>.figs{padding-top:12px}
.tl-tools{display:flex;gap:10px 16px;align-items:center;flex-wrap:wrap;min-height:56px;padding:10px 0;margin-top:12px}.tl-sw{display:inline-flex;gap:12px;align-items:center}.tl-tools .lg{display:flex;flex-wrap:wrap;align-items:center;gap:4px 16px;font-size:13px;line-height:20px;color:var(--muted);list-style:none;margin:0 0 0 auto;padding:0}.tl-tools .lg li{display:inline-flex;align-items:center;gap:6px}.tl-tools .lg i{display:inline-block;border-radius:1px;background:var(--k)}.tl-brush{display:grid;grid-template-columns:200px minmax(0,1fr);height:44px;align-items:center}.tl-brush .tr{height:20px;border-radius:4px;background:#ffffff08}.tl-foot{padding-top:10px}
.gear-c{container-type:inline-size}.gear-grid{display:grid;gap:20px}.gear-aside{display:flex;flex-direction:column;gap:8px;min-width:0}
@container (min-width:760px){.gear-grid.pair{grid-template-columns:minmax(0,640px) minmax(240px,1fr);align-items:start}}
.empty-open{padding:12px 0;color:var(--muted);font-size:14px}
.btn.is-off{opacity:.45;pointer-events:none}
button.link{background:none;border:0;padding:0;cursor:pointer;font:600 14px/20px var(--font-body);color:var(--gold)}
input.field{font:400 14px/20px var(--font-body);color:var(--strong)}input.field::placeholder{color:var(--muted)}
.bar-m summary{gap:8px}.bar-m summary>span:first-child{display:contents}.bar-m summary .sum-t{flex:1;min-width:0;overflow:hidden;text-overflow:ellipsis;white-space:nowrap}.bar-m summary .chg{order:2;flex:none;margin-left:0}.bar-m summary .run-m{order:3;flex:none}.bar-m summary .run-m .btn{height:36px;padding:0 12px;font-size:12px}
.paste{display:flex;flex-direction:column;gap:12px;padding:16px}.paste-hint{margin-top:12px}
.ledger-panels{display:flex;flex-direction:column;gap:22px}
.ways-fill .row{grid-template-columns:180px minmax(0,1fr)}.ways-fill .row .trail{justify-self:stretch;min-width:0}.ways-fill .row .trail>form{flex:1}
.way{display:flex;flex-direction:column;gap:10px;padding:16px 0}
.cbar .grp>.t-label{display:none}@media(min-width:1920px){.cbar .grp>.t-label{display:block}}
@media(max-width:1023px){.pc .layout-split{display:flex;flex-direction:column;align-items:stretch;gap:var(--section-gap)}.pc .layout-split section{min-width:0;max-width:100%}.pc .layout-split>.main,.pc .layout-split>.rail{display:contents}[data-region=section-damage-by-ability]{order:1}[data-region=section-dps-spread]{order:2}[data-region=section-timeline]{order:3}[data-region=section-gear]{order:4}[data-region=section-talents]{order:5}[data-region=section-rage]{order:6}[data-region=section-buffs-and-debuffs]{order:7}[data-region=section-this-run]{order:8}[data-region=section-rotation]{order:9}[data-region=section-your-sims]{order:10}}
.statesheet{display:grid;gap:40px;grid-template-columns:minmax(0,1fr)}.statesheet>section{min-width:0}.statesheet>section>h2{margin-bottom:8px}
.cls-links{display:flex;gap:4px 16px;flex-wrap:wrap;padding-top:2px}.cls-links .link{font-size:14px;line-height:18px}
/* round 5 panels (ui/Section `panel`): the body starts right under the head row, so the open-section top gaps above go */
.pn-b>.rail-body,.pn-b>.two{padding-top:0}.pn-b>.figs+.two{padding-top:8px}.pn-b>.tl-tools{margin-top:0}.rot .rows .row{border-bottom:0}.rk-s{white-space:nowrap}.am .rk-s{margin-left:4px}
.formrow label.f>.t-label:empty{display:none}
.rgrid .cls-links{flex-wrap:nowrap}.rgrid .cls-links .link{white-space:nowrap;flex:none}
.bar-end{display:flex;gap:20px;align-items:center}.bar-end .link{font-size:14px}
.bar-d.is-run .grp,.bar-d.is-run .bar-end .link{opacity:.5;pointer-events:none}
.fig.muted .v{color:var(--muted)}
/* the built Figure sm (ui/Figure.svelte, also BiS weights and the guild pages): 16/20 */
body .fig.sm .v{font-size:16px;line-height:20px}
/* the built histogram's axis: 8 px under the plot, 13/20 (DpsDistribution) */
body .hist .ax{padding-top:8px;line-height:20px}
/* the read-only trees (ws_prims TalentTreesAll) shrink in their column */
.tt-trees{min-width:0}
/* the built Section's body starts 12 px under its hairline (ui/Section.svelte); a Plate already has it, open blocks take it here */
.sec-h+.doll,.sec-h+div>.tt-trees{margin-top:12px}.sec-h+.two{padding-top:12px}
/* the built Section head wraps its action under a long meta (gap-y 4 px); a Row is --row-h, 58 px (spec 2026-11-05 P4) */
body .sec-h{flex-wrap:wrap;row-gap:4px}body .sec-h .act{line-height:21px}body .row{min-height:58px}body .sec-h+.rows{margin-top:12px}
/* Rotation as the built rows draw it: the icon in the title (ActionMark) */
.am{display:flex;align-items:center;gap:10px;min-width:0}.am>span:last-child{min-width:0}
/* This run spaces its blocks as the built primitives do (IconList's own 16 px label margin), not by the rail's 12 px gap */
.rail-body.flow{gap:0}.rail-body.flow>.rail-cap{padding-top:12px}.rail-body.flow>.rail-acts{margin-top:20px}
/* spec 4.7: the row's name line is 15/20 (the body's 24 px strut made it 25 and the row 60) */
.st-r .nm{line-height:20px}
.sheet .kv dd{display:flex;flex-direction:column;align-items:flex-end;gap:4px}
</style>'''


def regions(html: str) -> str:
    """The one region ws_prims leaves to the page: `summary` (Worth knowing, the Notice lines under the bar). Every other region
    (head, head-title, head-meta, head-answer, controls, plate, slots, section-<slug>) is emitted by the primitives themselves,
    once, as the built ui/ components emit it (design/parity.md)."""
    return html.replace('<div class="pc"><p class="notice', '<div class="pc"><p data-region="summary" class="notice', 1)


NOTE = (f'Simulator mock (design/specs/2026-11-06-sim.md), built only from ws_prims. Run: web/src/fixtures/sim/result-full.json '
        f'(the published warrior-fury 60 raid set, 10,000 runs); names, icons, items, talents and rotation from build {BUILD}. '
        'SAMPLE: the signed-in name Thandrel, the last sims, the live estimate, the engine-refusal reason and the Your sims rows.')


def finish(title: str, body: str, who) -> str:
    return regions(build_page(title, body + PAGE_CSS, 'Tools', who, NOTE))


# ---- 1. /sim, signed out, no character
def landing_out() -> str:
    head = PageHeader('Simulator', tabs=TabStrip(TOOLS, 'Quick sim'), lead=copy_string('lib/sim/page-copy.ts', 'lead'))
    paste = (Form('Addon export or build code', [('', Input('FS1:&hellip;', 'Addon export or build code', grow=True), True)], Button('Run', primary=True))
             + f'<p class="t-small paste-hint">In game: /fs copy, then paste here. {Link("Get set up &rsaquo;")}</p>')
    # ui/Row `fillTrail`: a 180 px label column, the control filling the rest (the field grows, Load beside it)
    ways = '<div class="ways-fill">' + Rows([Row('Your characters', trail=Button('Sign in with Battle.net'), link=False),
                 Row('A planner build', trail=Form('A planner build', [('', Input('foreversixty.gg/b/&hellip;', 'A planner build', grow=True), True)], Button('Load')), link=False),
                 Row('A logged fight', trail=Form('A logged fight', [('', Input('foreversixty.gg/reports/&hellip;?fight=2', 'A logged fight', grow=True), True)], Button('Load')), link=False)], open_=True) + '</div>'
    rows = []
    for c in CLASSES:
        specs = [s for s in SPECS if s['class_slug'] == c['slug'] and s['role'] == 'dps']
        if specs:
            links = ''.join(f'<a class="link" href="#">{s["name"]} &rsaquo;</a>' for s in specs)
            rows.append(Row(c['name'], f'<span class="cls-links">{links}</span>', lead=crest_img(c['slug'], 32), title_color=CLASS_COLOR[c['slug']], link=False))
    main = (Section('Sim your character', paste, panel=True) + Section('Other ways in', ways, panel=True)
            + Section('Sim a best-in-slot set', RowGrid(rows, 3), panel=True))
    return finish('Simulator', head + Ledger(main), None)


def Ledger(main: str) -> str:
    """ui/Page `panels` with the ledger layout: one column of Section panels, 28 px under the band, 22 px apart."""
    return f'<div class="pc" style="padding-top:28px"><div class="ledger-panels">{main}</div></div>'


# ---- 2. /sim, signed in, the header's character loaded, not run yet
def landing_you() -> str:
    # Your sims under the talents, so the columns end together (SimPreRun); the rail is the rotation
    main = Section('Gear', doll(rail_gear()), Link('Upgrades on Best in slot &rsaquo;')) + talents_section() + history_section()
    return finish('Simulator', you_header(answer_none()) + run_bar() + Page('split', main, rotation_section(ran=False)), YOU)


# ---- 3. running in the browser
def running() -> str:
    main = Section('Damage by ability', table_skeleton(), meta='Draws when the run finishes', frame='framed', flush=True) + gear_section(with_stats=False) + talents_section() + history_section()
    return finish('Simulator, running', you_header(answer_running()) + run_bar('running') + Page('split', main, rotation_section(ran=False)), YOU)


# ---- 4. the live result on /sim
def result_live() -> str:
    return finish('Simulator, result', you_header(answer_result(), result_actions()) + run_bar('done') + result_page_body(saved=False), YOU)


# ---- 5. /sim/<id>, a saved run, signed out
def result_saved() -> str:
    meta = [f'{RACE} &middot; Level {CHAR["level"]} &middot; {SPLIT} talents &middot; {FIGHT_LEN} fight, raid buffs, {len(CHAR["consumes"])} consumables', 'Saved 4 Nov']   # the current year goes unsaid (2026-11-09 head contract)
    ans = f'<span class="t-label"></span>{Figure(FIGURE, unit="DPS", tone="gold", size="big")}<span class="fig-cap">{ERR_LINE}</span>'
    head = PageHeader(SPEC_LABEL, CLASS_COLOR[CLS], crest=crest_img(CLS, 56), meta=meta, answer=ans,
                      actions=f'{_Button("Open in the sim", primary=True)}{Button("Copy link")}')
    return finish(f'{SPEC_LABEL}, saved sim', head + result_page_body(saved=True), None)


# ---- 6. phone: the bar's choices open
def bar_open() -> str:
    main = Section('Gear', doll())
    return finish('Simulator, choices', you_header(answer_none()) + run_bar(open_m=True) + Page('split', main, history_section()), YOU)


# ---- 7. every other state on one sheet
def states() -> str:
    def blk(t, html):
        return f'<section class="st-blk"><h2 class="t-h3">{t}</h2>{html}</section>'
    hdr = lambda ans, title='Thandrel', cls=CLS, meta=f'{RACE} {SPEC_LABEL} &middot; Level 60': PageHeader(title, crest=crest_img(cls, 56), meta=[meta], answer=ans)
    sheet = [
        blk('No sim yet for this character', hdr('<span class="t-label"></span>') + run_bar()),
        blk('Engine loading (the first run on this device)', hdr(f'<span class="t-label">Loading the engine</span><span class="fig-cap">The engine is about 4 MB. It loads once and is cached after that.</span><span class="ans-meter">{Meter(.35)}</span>') + run_bar('engine')),
        blk('On our servers: queued (Premium)', hdr(f'<span class="t-label">On our servers</span><span class="fig-cap"><b>Starting</b></span><span class="ans-meter">{Meter(0, "muted")}</span>') + run_bar('server')),
        blk('On our servers: running', hdr(f'<span class="t-label">On our servers</span><span class="fig-cap"><b>Running</b> &middot; takes about 2 minutes</span><span class="ans-meter">{Meter(.4)}</span>') + run_bar('server')),
        blk('On our servers: stopped before it finished (job reason "stopped before completion")',
            hdr(answer_failed('On our servers &middot; did not finish', 'Our server stopped this run before it finished.', f'Press Run again. If it stops again, choose Normal precision ({3000:,} runs): it finishes sooner.')) + run_bar('done')),
        blk('On our servers: failed for any other reason (job reason "the run could not be completed")',
            hdr(answer_failed('On our servers &middot; failed', 'Our server could not finish this run.', 'Press Run again. If it fails twice, run it in your browser: choose a precision under High.')) + run_bar('done')),
        blk('In your browser: the engine did not download', hdr(answer_failed('In your browser &middot; engine not loaded', 'The engine did not download.', 'Check your connection, then press Try again.')) + run_bar('retry')),
        blk('In your browser: the engine refused the character (SAMPLE reason: the engine names the id it cannot read)',
            hdr(answer_failed('In your browser &middot; failed', 'The engine could not run this character: Finger 1 holds item 999999, which the engine does not know.', 'Paste a fresh addon export, or swap that item in Top gear, then press Run again.')) + run_bar('done')),
        blk('Stopped by the player (the partial number stays, greyed)', hdr(answer_stopped()) + run_bar('done')),
        blk('Result, signed out (Save needs an account)', hdr(answer_result(signed_in=False))),
        blk('Later (G1, G15): the last sim on the landing', hdr(answer_last()) + run_bar()),
        blk('Later (G1, G15): the change since the last sim', hdr(answer_result(last=LAST_SIM))),
        blk('Later (G1, G15): the change since the last sim inside the noise', hdr(answer_result(last=LAST_SIM_CLOSE))),
        blk('A spec the sim does not run (Holy Priest)', hdr('<span class="t-label">Not simulated yet</span><span class="fig-cap">Holy Priest is not simulated yet; there is nothing to run.</span><span class="ans-acts"><a class="link" href="#">Spec support &rsaquo;</a></span>', 'Lightwell', 'priest', 'Human Holy Priest &middot; Level 60') + run_bar('disabled')),
        blk('Signed in, no characters yet', PageHeader('Simulator', meta=['Your character&rsquo;s DPS, simmed in your browser.']) + Notice('<span>' + 'No characters yet: they are on their way from Battle.net. Paste an export below meanwhile, or check <a class="link" href="#">your account.</a>' + '</span>', tone='neutral')),
        blk('The character could not be read', Notice('<span>' + 'That character could not be read. <a class="link" href="#">Try again</a>' + '</span>', tone='neutral')),
        blk('No build recorded for the character', Notice('<span>' + 'No build yet for Thandrel. <a class="link" href="#">Paste the addon export</a>, or refresh from Battle.net on your <a class="link" href="#">account page</a>.' + '</span>', tone='neutral')),
        blk('A paste that does not decode', Plate('<div class="paste">' + FormRow([('Addon export or build code', Field('FS1:1.60.1.70790:warrior:troll:3032&hellip;', grow=True), True)], '<span class="btn primary is-off">Run</span>') + '<p class="t-small wipe">That does not look like an addon export or a build code.</p></div>')),
        blk('Saved sim: loading (prerendered skeleton at the final geometry)', PageHeader('&nbsp;', meta=[Skeleton(260, 14)], answer=Skeleton(220, 60) + Skeleton(160, 14)) + Page('split', Section('Damage by ability', table_skeleton(6), frame='framed', flush=True), Section('DPS spread', Skeleton(300, 168)))),
        blk('Saved sim: no such id', Plate('<p class="empty">No sim with that id. <a class="link" href="#">Open the simulator &rsaquo;</a></p>')),
        blk('Saved sim: the server did not answer', Plate(f'<p class="empty">This sim did not load. The server did not answer. Your link is fine; try again in a moment. {Link("Try again")}</p>')),
        blk('Saved sim: an older engine', Notice('<span>' + 'This run used an older engine. A run today may give a different number. <a class="link" href="#">Open in the sim &rsaquo;</a>' + '</span>', tone='neutral')),
        blk('Older run: no histogram kept', Section('DPS spread', f'<div class="rail-body">{Meter(.5, "muted")}<p class="t-small">This run kept the range, not the shape. The shaded part is one spread either side of the mean.</p></div>')),
        blk('Older run: no recorded fight', Section('Timeline', '<p class="empty-open">This run has no recorded fight to draw.</p>')),
        blk('No gear on file', Section('Gear', '<p class="empty-open">This run had no gear on file.</p>')),
        blk('Talents saved on an older layout', Section('Talents', '<p class="empty-open">These talents were saved on an older talent layout. The point split above is as saved; the trees have changed since, so they are not drawn.</p>', meta='20 / 31 / 0')),
        blk('A lazy section failed', Section('Timeline', f'<p class="empty-open">This part of the page didn&rsquo;t load. {Link("Try again")}</p>')),
        blk('Your sims: empty, then failed', history_section('empty') + history_section('failed')),
        blk('Ability table: row hover', Plate(StatTable(TABLE_HEADS, ability_rows('hover')[0][:4], N_FIGS, sort='DPS'))),
        blk('Timeline: Swings and ticks on', timeline_section(swings_and_ticks=True)),
        blk('Buffs and debuffs at a narrow column (the primitive keeps every figure inside its column)', f'<div style="max-width:420px">{buffs_section()}</div>'),
    ]
    body = '<div class="pc" style="padding-top:24px"><h1 class="t-title">Simulator: states</h1></div>' + Page('ledger', f'<div class="statesheet">{"".join(sheet)}</div>')
    return finish('Simulator states', body, YOU)


# ---- 8. More settings, the overlay sheet (ruling S7), over the live result
def more_settings() -> str:
    ts = RES['target_stats']
    target = sel(f"Level {ts['level']} boss")
    rows = Facts([('Execute phase', f"{seg([('On', True, None, ''), ('Off', False, None, '')], 'sm')}<small>Last {FIGHT['execute_ratio'] * 100:.0f}% of the fight</small>"),
                  ('Target', f"{target}<small>{ts['armor']:,} armor</small>"),
                  ('Rotation', sel(f"The {SPEC_ROW['name']} rotation"))])
    body = (rows + IconList(consume_items(), f'Consumables &middot; {len(CHAR["consumes"])}') + f'<p class="t-small rail-cap">{Link("Change consumables &rsaquo;")}</p>'
            + f'<h3 class="t-label il-l">Request</h3><p class="t-small rail-cap">The whole request this run sends, as JSON. {Link("Show it &rsaquo;")}</p>')
    sheet = Sheet('More settings', body, Button('Back to the preset') + Button('Done', primary=True))
    return finish('Simulator, more settings', you_header(answer_result(), result_actions()) + run_bar('done') + result_page_body(saved=False) + sheet, YOU)


PAGES = {'web-sim': landing_out, 'web-sim-you': landing_you, 'web-sim-running': running, 'web-sim-result': result_live,
         'web-sim-saved': result_saved, 'web-sim-bar': bar_open, 'web-sim-more': more_settings, 'web-sim-states': states}


def compare() -> None:
    """Before/after at the same scale against the 2026-11-06 captures (production for /sim, the fixture build for /sim/<id>).
    Pillow only here: the parity spec renders the mocks with the standard library alone."""
    import ws_compose
    ws_compose.BEFORE = ROOT / 'design/reviews/2026-11-06-sim/before'
    for w, sc in ((1440, .5), (2000, .45), (390, .8)):
        print(ws_compose.compose('sim-landing', 'sim-landing', 'web-sim', w, sc, max_h=3600))
        print(ws_compose.compose('sim-result', 'sim-live', 'web-sim-result', w, sc, max_h=8000))
        print(ws_compose.compose('sim-saved', 'sim-saved', 'web-sim-saved', w, sc, max_h=8000))


if __name__ == '__main__':
    if sys.argv[1:] == ['compare']:
        compare()
    else:
        for name, fn in PAGES.items():
            print(write(f'{name}.html', fn()))
