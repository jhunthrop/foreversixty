"""Web-layout mock: the simulator's tool pages (/sim/gear, /sim/talents, /sim/drops, /sim/weights, /sim/specs), composed ONLY
from ws_prims, in the quick sim's grammar (design/specs/2026-11-06-sim.md, design/mocks/ws_sim.py): the PageHeader with the tool
tabs, the run bar (the quick sim's ControlBar), the character line, then the split page with the page's one Plate.

The character is the tool parity state's (web/tests/e2e/support/sim-parity.ts `TOOL_CHARACTER`): a level 60 Orc Fury Warrior
entered by hand, wearing Lionheart Helm (12640) and Arcanite Reaper (12784). Names, icons, item levels, the consumables Top gear
offers, the stats the weights page weighs and the Droptimizer's sources are read at render time from data/builds/<active build>/
(items/warrior.json, loot.json, stat-weights.json) and sim/request/raid-consumes.json (the raid set's consumables), the files the
built pages read. Item search shows its first 100 hits (lib/sim/item-search.ts SEARCH_LIMIT) at a fixed height; the mock draws
the class's 100 highest-level items there.

  FS_DATA_ROOT=<checkout> python3 design/mocks/ws_simtools.py
  for p in web-simtools-gear web-simtools-gear-empty web-simtools-talents web-simtools-drops web-simtools-weights web-simtools-specs; do
    node design/mocks/ws_shoot.mjs $p 390,1440,2000; done
"""
import json
from datetime import datetime, timezone
import os
import sys
from pathlib import Path

from mocklib import CLASS_COLOR
from ws_prims import Facts, Field, FormRow, Notice, Page, PageHeader, Plate, Row, Rows, Section, TabStrip, TalentTreesAll, build_page, talent_ranks_from_strings, talent_tree_dicts, write
import ws_sim as S

ROOT = Path(__file__).resolve().parents[2]
DATA_ROOT = Path(os.environ.get('FS_DATA_ROOT', ROOT))
WEB = DATA_ROOT / 'web/src'
B = S.B
ITEMS = S.ITEMS
TOOLS = S.TOOLS
SPEC = 'warrior-fury'
LEVEL = 60
WORN = {'head': 12640, 'main_hand': 12784}
SPECS = json.loads((DATA_ROOT / 'data/curated/specs.json').read_text())
CLASS_NAME = S.CLASS_NAME
QUALITY = S.QUALITY

TITLES = {'gear': 'Top gear', 'talents': 'Talents', 'drops': 'Droptimizer', 'weights': 'Weights', 'specs': 'Spec support'}


# Round 5 (lane P6): every block is a Section panel (web ui/Section `panel`; tools/ToolSection), the page's primary object the
# framed one, on a Page with `panels` (22 px apart, 28 px under the band).
_Section = Section
_Page = Page


def Section(*args, **kw) -> str:
    kw.setdefault('panel', True)
    return _Section(*args, **kw)


def Page(layout: str, main: str = '', rail: str = '', sticky_rail: bool = False) -> str:
    html = _Page(layout, main, rail, sticky_rail, panels=layout == 'split')
    return html.replace('style="padding-top:20px"', 'style="padding-top:28px"', 1) if layout == 'ledger' else html

def ledger(*blocks: str) -> str:
    """The ledger page with the built Page's section gap between its blocks (the summary line included)."""
    return Page('ledger', f'<div class="col">{"".join(blocks)}</div>')


def split(main: list[str], rail: list[str]) -> str:
    return Page('split', f'<div class="col">{"".join(main)}</div>', f'<div class="col rail-col">{"".join(rail)}</div>')


# The tool heads' leads (SimToolHead `lead`): only /sim/weights carries one (seo-fix 2026-11-10).
LEADS = {'weights': S.copy_string('lib/sim/copy.ts', 'weightsLead')}


def head(tool: str) -> str:
    return PageHeader(TITLES[tool], tabs=TabStrip(TOOLS, TITLES[tool]), lead=LEADS.get(tool, ''))


def item_name(iid: int) -> str:
    it = ITEMS[iid]
    return f'<b style="color:{QUALITY[it["quality"]]}">{it["name"]}</b>'


def btn(label: str, primary: bool = False, expanded: bool | None = None) -> str:
    exp = '' if expanded is None else f' aria-expanded="{"true" if expanded else "false"}"'
    return f'<button type="button" class="btn sm{" primary" if primary else ""}"{exp}>{label}</button>'


def check(label: str = '', checked: bool = False, disabled: bool = False) -> str:
    aria = f' aria-label="{label}"' if label else ''
    return f'<input type="checkbox" class="cb"{aria}{" checked" if checked else ""}{" disabled" if disabled else ""}>'


# ---- the run bar and its standing lines (BulkRunBar: SimRunBar with the tool's words)
def run_bar(tool: str) -> str:
    bar = S.run_bar('disabled' if tool != 'weights' else 'ready')
    # Nothing ticked: the count reads muted and the stage ladder waits for a count (BulkRunBar `hasCount`).
    # A weights run says nothing under the bar (BulkRunBar).
    reason = {'gear': 'Tick a candidate to try.', 'talents': 'Pick builds to compare.', 'drops': 'Tick a source.'}.get(tool, '')
    line = f'<span class="t-small">{reason}</span>' if reason else ''
    return bar + f'<div class="pc run-notes"><p>{line}</p></div>'


# ---- the character line (ToolCharacterLine, region `summary`)
def character_line() -> str:
    level = sum(ITEMS[i]['item_level'] for i in WORN.values()) / len(WORN)
    return (f'<div class="tool-who" data-region="summary"><p class="who-line">{S.crest_img("warrior", 24)}<span>Orc Fury Warrior &middot; Level {LEVEL}</span>'
            f'<span class="t-small num">ilvl {level:.1f}</span>'
            f'<span class="t-small num">0/20/0</span>'
            f'<span class="who-acts"><a class="link" href="#">Open in planner &rsaquo;</a><button type="button" class="btn sm ghost">Change source</button></span></p></div>')


# ---- Top gear: the candidates (SlotGrid, CandidateRows)
def candidate(slot_label: str, iid: int) -> str:
    it = ITEMS[iid]
    lead = check(it['name']) + S.gicon(it['icon'], 32, it['quality'])
    trail = f'<span class="t-small num">{it["item_level"]}</span><button type="button" class="btn sm ghost" aria-expanded="false">Modify</button>'
    title = f'<button type="button" class="hover-name">{item_name(iid)}</button>'
    return (f'<div class="cand-group"><div class="cand-h"><h3 class="t-h4">{slot_label}</h3><label class="t-small lock">{check()}Lock</label></div>'
            + Rows([Row(title, '', lead=lead, trail=trail, link=False)], open_=True) + '</div>')


def candidates() -> str:
    return Section('Candidates', '<div class="cands">' + candidate('Head', WORN['head']) + candidate('Main hand', WORN['main_hand']) + '</div>')


def item_search() -> str:
    # The first page of hits, then a button for the next (ItemSearch PAGE): no box that cuts a row in half.
    usable = sorted((i for i in ITEMS.values() if i.get('item_level')), key=lambda i: (-i['item_level'], i['name']))[:8]
    rows = ''.join(f'<div class="hit">{S.gicon(i["icon"], 28, i["quality"])}<span class="hit-n">{item_name(i["id"])}</span>'
                   f'<span class="t-small num">{i["item_level"]}</span><button type="button" class="btn sm ghost">Add</button></div>' for i in usable)
    filters = (f'<div class="search-f"><input type="search" class="fld grow" placeholder="Name" aria-label="Name">'
               f'<label class="fl"><span class="t-label">Minimum item level</span><input type="number" class="fld w24" value="0"></label>'
               f'<label class="fl"><span class="t-label">Slot</span><select class="fld w160"><option>Any slot</option></select></label>'
               f'<span class="fl"><span class="t-label">Source</span><button type="button" class="fld w240 sel" aria-expanded="false">Anywhere</button></span>'
               f'<span class="sw-l"><button type="button" role="switch" aria-checked="true" class="sw"></button>Only items this character can equip</span></div>')
    return Section('Find an item', filters + f'<div class="hits">{rows}</div><div class="hits-more"><button type="button" class="btn sm ghost">8 more</button><p class="t-small trunc">Showing 100 of 9,026. Narrow the search.</p></div>')


def consumables() -> str:
    table = json.loads((DATA_ROOT / 'sim/request/raid-consumes.json').read_text())['specs'][SPEC]
    ids = list(dict.fromkeys([r for r in table if r['from_level'] <= LEVEL][-1]['consumes']))
    rows = ''.join(f'<li><label class="opt">{check()}{x.replace("_", " ")}</label></li>' for x in ids)
    return Section('Try each', f'<ul class="opts ce cols3">{rows}</ul>')


def whole_sets() -> str:
    return Section('Whole sets', ''
                   '<div class="sets-f"><label class="fl w160"><input class="fld" placeholder="Name" aria-label="Name"></label>'
                   '<label class="fl grow"><input class="fld" placeholder="Paste an export string" aria-label="Paste an export string"></label>'
                   '<button type="button" class="btn">Add set</button></div>')


def guide_split() -> str:
    f = json.loads((B / 'bis' / f'{SPEC}.json').read_text())
    band = next(b for b in f['bands'] if b['band'] == LEVEL and b['faction'] == 'horde' and b.get('preset') == 'raid')
    return '/'.join(str(sum(int(c) for c in tree)) for tree in band['talents'].split('-'))


# The character's own talents (TOOL_CHARACTER's `0/5530515/0`), one digit string per tree.
OWN_TALENTS = ['0', '5530515', '0']


def guide_talents() -> list[str]:
    f = json.loads((B / 'bis' / f'{SPEC}.json').read_text())
    band = next(b for b in f['bands'] if b['band'] == LEVEL and b['faction'] == 'horde' and b.get('preset') == 'raid')
    return band['talents'].split('-')


def loadout_trees(per_tree: list[str]) -> str:
    """web TalentCandidates: each build's trees at the rank its string gives them, in the site's read-only trees (view)."""
    return TalentTreesAll(talent_tree_dicts('warrior', talent_ranks_from_strings('warrior', per_tree)), 'view')


def loadout_ranks(per_tree: list[str]) -> list[tuple[str, int]]:
    file = json.loads((B / 'talents' / 'warrior.json').read_text())
    out = []
    for i, tree in enumerate(sorted(file['trees'], key=lambda t: t['position'])):
        digits = per_tree[i] if i < len(per_tree) else ''
        out += [(x['name'], min(x['max_rank'], int(digits[p]) if p < len(digits) else 0)) for p, x in enumerate(tree['talents'])]
    return out


def loadout_diff(talents: list[str], other: list[str]) -> list[tuple[str, int]]:
    """web lib/sim/talent-loadouts.ts loadoutDiff: each talent whose rank differs, gains first, then by size."""
    changes = [(n, r - o) for (n, r), (_, o) in zip(loadout_ranks(talents), loadout_ranks(other)) if r != o]
    return sorted(changes, key=lambda c: (-(c[1] > 0), -abs(c[1])))


def talent_builds(plated: bool = False) -> str:
    # The published set's own talents at the character's band (web lib/sim/talent-loadouts.ts guideLoadout)
    own_split = '/'.join(str(sum(int(c) for c in t)) for t in OWN_TALENTS)
    own = f'<label class="opt">{check()}<span class="grow-n">Your current build</span>' + (f'<span class="t-small num">{own_split}</span>' if plated else '') + '</label>'
    guide = f'<label class="opt">{check()}<span class="grow-n">Guide build, level {LEVEL} &middot; {guide_split()}</span></label>'
    if plated:
        # /sim/talents: each build's three trees beside its tick (TalentCandidates `plate`)
        def diff_groups(changes: list[tuple[str, int]]) -> str:
            out = ''
            for title, rows, tone in (('Gains', [c for c in changes if c[1] > 0], 'kill'), ('Drops', [c for c in changes if c[1] < 0], 'wipe')):
                if rows:
                    lis = ''.join(f'<li class="{tone}">{"+" if d > 0 else "&minus;"}{abs(d)} <span>{n}</span></li>' for n, d in rows)
                    out += f'<div class="tb-g"><h4 class="t-label">{title}</h4><ul class="tb-d">{lis}</ul></div>'
            return out
        own = (f'<div class="tb-row">{own}<div class="tb-t">{loadout_trees(OWN_TALENTS)}'
               f'<div class="tb-gs">{diff_groups(loadout_diff(OWN_TALENTS, guide_talents()))}</div></div></div>')
        changes = loadout_diff(guide_talents(), OWN_TALENTS)
        groups = ''
        for title, rows, tone in (('Gains', [c for c in changes if c[1] > 0], 'kill'), ('Drops', [c for c in changes if c[1] < 0], 'wipe')):
            if rows:
                lis = ''.join(f'<li class="{tone}">{"+" if d > 0 else "&minus;"}{abs(d)} <span>{n}</span></li>' for n, d in rows)
                groups += f'<div class="tb-g"><h4 class="t-label">{title}</h4><ul class="tb-d">{lis}</ul></div>'
        guide = f'<div class="tb-row">{guide}<div class="tb-t">{loadout_trees(guide_talents())}<div class="tb-gs">{groups}</div></div></div>'
    body = (f'<div class="tb">{own}{guide}'
            f'<h3 class="t-h4">Your saved builds</h3>'
            f'<p class="t-small">Saved builds did not load.</p><div class="bt">{btn("Add a build")}</div></div>')
    return Section('Talent builds', body)


def gear() -> str:
    # The other ways to add a candidate stack under the grid; the rail holds the item search alone (TopGear.svelte).
    return finish('Top gear', head('gear') + run_bar('gear')
                  + split([character_line(), candidates(), consumables(), whole_sets(), talent_builds()], [item_search()]))


def talents() -> str:
    return finish('Talents', head('talents') + run_bar('talents') + ledger(character_line(), talent_builds(plated=True)))


PHASES = json.loads((WEB / 'data/phases.json').read_text())


def is_open(source: dict) -> bool:
    """lib/sim/phase.ts hasOpened, at the moment the mock is drawn (the build reads the same clock)."""
    name = source.get('opens')
    if not name:
        return True
    if name == 'later':
        return False
    row = next((p for p in PHASES if p['name'] == name), None)
    if row is None:
        return False
    start = row.get('start') or ''
    if not start or start.startswith('0001'):
        return True
    return datetime.now(timezone.utc) >= datetime.fromisoformat(start.replace('Z', '+00:00'))


# ---- Droptimizer: the sources (SourcePicker), raids and dungeons open, a dungeon's bosses folded
def drops() -> str:
    loot = json.loads((B / 'loot.json').read_text())['sources']
    kinds = [('raid', 'Raids'), ('dungeon', 'Dungeons'), ('crafted', 'Crafted'), ('rep', 'Reputation'), ('pvp', 'PvP')]
    labels = [('raid', 'Raids'), ('dungeon', 'Dungeons'), ('world', 'World bosses'), ('world_drop', 'World drops'), ('zone', 'Zone drops'),
              ('vendor', 'Vendors'), ('crafted', 'Crafted'), ('rep', 'Reputation'), ('pvp', 'PvP'), ('quest', 'Quests')]   # lib/sim/loot.ts LOOT_KINDS
    on = {k for k, _ in kinds}
    present = {s['kind'] for s in loot}
    toggles = ''.join(f'<label class="opt">{check(checked=k in on)}{label}</label>' for k, label in labels if k in present)
    toggles += f'<label class="opt">{check()}Show unreleased content</label>'
    hide_empty = {'dungeon', 'crafted', 'rep', 'pvp', 'world_drop', 'zone', 'vendor'}   # lib/sim/source-rows.ts showsSource

    def usable(src: dict) -> list[dict]:
        ids = list(src.get('items', [])) + [i for b in src.get('bosses') or [] for i in b.get('items', [])] + list(src.get('trash', []))
        return [ITEMS[i] for i in dict.fromkeys(ids) if i in ITEMS]

    groups = []
    for kind, label in kinds:
        srcs = [s for s in loot if s['kind'] == kind]
        shown = [s for s in srcs if is_open(s)]
        if kind == 'pvp':
            # The character's own side (lib/sim/source-rows.ts pvpForFaction): an Orc is Horde.
            shown = [s for s in shown if s.get('faction') in (None, 'horde')]
        shown = [s for s in shown if usable(s) or kind not in hide_empty]
        if not srcs:
            continue
        if not shown:
            # The gated kind's heading, then one line per dated source and the rest named once (SourcePicker).
            undated = ', '.join(s['name'] for s in srcs)
            opened_n = sum(1 for x in srcs if is_open(x))
            groups.append(f'<div class="src-g"><details class="disc"><summary>{label} <span class="t-small">{opened_n} of {len(srcs)} open</span></summary></details></div>')
            continue
        rows = ''
        if kind == 'rep':
            # One row per faction, its standings as inline ticks (lib/sim/source-rows.ts groupByFaction).
            factions: dict[str, list[dict]] = {}
            for s in shown:
                factions.setdefault(s['name'], []).append(s)
            for name, standings in factions.items():
                ticks = ''.join(f'<label class="opt sm">{check()}{(s.get("standing") or "").title()} <span class="t-small num">{len(usable(s))}</span></label>'
                                for s in standings)
                rows += f'<div class="src"><b class="rep-n">{name}</b><span class="rep-t">{ticks}</span></div>'
        else:
            for s in shown:
                bosses = s.get('bosses') or []
                items = usable(s)
                best = max(items, key=lambda it: it.get('item_level', 0)) if items else None
                icon = S.gicon(best['icon'], 24, best['quality']) if best else ''
                reward = f' <span class="t-small">{best["name"]}</span>' if kind == 'pvp' and best else ''
                count = (f'<button type="button" class="t-small bcount" aria-expanded="false">{len(bosses)} {"boss" if len(bosses) == 1 else "bosses"}</button>'
                         if kind == 'dungeon' and bosses else '')
                ilvl = f'<span class="pl">ilvl {best["item_level"]}</span>' if best else ''
                rows += (f'<div class="src"><div class="src-l"><label class="opt grow-n">{check()}{icon}<span class="grow-n"><b>{s["name"]}</b>{reward}</span></label>'
                         f'{count}<span class="t-small num">{len(items)} items{ilvl}</span></div>')
                if kind != 'dungeon' and bosses:
                    rows += ''.join(f'<label class="opt indent">{check()}<span class="grow-n">{x.get("name") or s["name"]}</span></label>' for x in bosses)
                rows += '</div>'
        note = ''   # nothing records the fixture's professions: no line (SourcePicker)
        # Each kind is a Disclosure titled with its ticked count, open when the list is short (OPEN_UP_TO).
        opened = ' open' if len(shown) <= 8 else ''
        # The kind's unopened sources, one deduped line under its rows (SourcePicker `later`, source-rows gatedLine)
        later = sorted({s['name'] for s in srcs if not is_open(s)}) if kind == 'raid' else []
        later_p = f'<p class="t-small later">First raids: {", ".join(later)}</p>' if later else ''
        inner = f'<div class="disc-b">{note}<div class="src-cols">{rows}</div>{later_p}</div>' if opened else ''
        groups.append(f'<div class="src-g"><details class="disc"{opened}><summary>{label} <span class="t-small">{len(shown)} {"source" if len(shown) == 1 else "sources"}</span></summary>{inner}</details></div>')
    body = f'<div><div class="kinds" role="group"><span class="t-small">Show</span>{toggles}</div>{"".join(groups)}</div>'
    return finish('Droptimizer', head('drops') + run_bar('drops')
                  + ledger(character_line(), Section('Sources', body)))


# ---- Weights: the caution, the stats to weigh in the rail; the scale factors arrive with a run
def weights() -> str:
    row = next(w for w in json.loads((B / 'stat-weights.json').read_text()) if w['spec'] == SPEC)
    stats = row.get('stats') or list(row['weights'])
    picks = ''.join(f'<li><label class="opt">{check(checked=True, disabled=i == 0)}{s}</label></li>' for i, s in enumerate(stats))
    # The stats to weigh as one row of ticks over the table (StatWeights.svelte); the reference is its locked tick.
    pick = Section('Stats to weigh', f'<ul class="opts inl">{picks}</ul>')
    return finish('Weights', head('weights') + run_bar('weights') + ledger(character_line(), pick, published_weights()))


# Before a run, the weights the ranker published for the character's band (web lib/sim/published-weights.ts): the level 60
# Horde (an Orc) raid band of data/builds/<build>/bis/<spec>.json, as the built WeightsScaleTable reads one
# (lib/bis/scale-view.ts): per-point rows only, each scale over the top significant per-point row, sorted; haste in the caption.
STAT_LABEL = {'attack_power': 'Attack power', 'strength': 'Strength', 'agility': 'Agility', 'crit': 'Crit', 'hit': 'Hit'}


def reading(w: dict) -> str:
    if w.get('insignificant') or abs(w['weight']) <= w.get('error', 0):
        return 'not-significant'
    return 'no-gain' if w['weight'] < 0 else 'value'


def published_weights() -> str:
    f = json.loads((B / 'bis' / f'{SPEC}.json').read_text())
    band = next(b for b in f['bands'] if b['band'] == LEVEL and b['faction'] == 'horde' and b.get('preset') == 'raid')
    haste = [w for w in band['weights'] if w['stat'] in ('melee_haste', 'spell_haste')]
    per_point = [w for w in band['weights'] if w['stat'] not in ('melee_haste', 'spell_haste', 'expertise')]
    divisor = max(w['scale_factor'] for w in per_point if reading(w) == 'value')
    rows = sorted(per_point, key=lambda w: (reading(w) != 'value', -w['scale_factor']))
    trs = []
    for w in rows:
        label = STAT_LABEL[w['stat']]
        name = f'<b>{label} rating</b>' if w.get('unit') == 'rating' else f'<b>{label}</b>'
        sub = f'{w["rating_factor"]} {label} rating = 1% {label}' if 'rating_factor' in w else ''
        sf = w['scale_factor'] / divisor
        err = w['scale_error'] / divisor
        err_txt = '&lt; 0.01' if err < .005 else f'&plusmn;{err:.2f}'
        trs.append(S.StatRow('', name, sub, [f'{sf:.2f}', f'{w["dps_per_point"]:.3f} DPS', err_txt], primary_phone=f'{w["dps_per_point"]:.3f} DPS &middot; {err_txt} error',
                             meter=S.Meter(min(1, max(.04, sf)))))
    top = STAT_LABEL[next(w for w in rows if reading(w) == 'value')['stat']]
    cap = ''   # the column heads say the normalisation; the caption keeps haste per 1% (owner rule)
    if haste:
        hs = band['haste_scale_factor'] / divisor
        cap += f'Haste: {hs:.2f} per 1%' + ('' if band.get('haste_on_items', True) else ', not in the table because no item at this level range has it')
    table = S.StatTable([('', False), ('Stat', False), ('', False), ('Scale', True), ('Per point', True), ('Error', True)], trs, 3, open_=True)
    body = f'<div class="wt-pub">{f"<p class=t-small>{cap}</p>" if cap else ""}{table}</div>'
    return Section('Published weights', body, action=S.Link(f'Level {LEVEL} raid set &rsaquo;'))


# ---- a tool with no character (SourceSwitcher): the addon export on the Plate, the other ways in open
def gear_empty() -> str:
    """SourceSwitcher, in the landing's own grammar (ws_sim.landing_out's Form, Input and Button)."""
    paste = ('<div class="paste-b">' + S.Form('From the addon', [('', S.Input('FS1:&hellip;', 'From the addon', grow=True), True)], S.Button('Load', primary=True))
                  + '<p class="t-small">Paste the export string from the Forever Sixty addon, or let the companion push it. <a class="link" href="#">Get the addon &rsaquo;</a></p></div>')
    form = lambda label, ph: S.Form(label, [('', S.Input(ph, label), False)], S.Button('Load'))
    ways = Rows([Row('From a build', 'Paste a planner link, saved or not, or a saved build id.', trail=form('From a build', 'foreversixty.gg/b/&hellip;'), link=False),
                 Row('From a logged fight', 'Paste a report link, or open a fight from a report and choose Sim this fight.', trail=form('From a logged fight', 'foreversixty.gg/reports/&hellip;?fight=2'), link=False),
                 Row('Your characters', 'Sign in with Battle.net to find your characters, with their gear and talents ready to sim.', trail=S.Button('Sign in with Battle.net'), link=False)], open_=True)
    return finish('Top gear', head('gear') + ledger(Section('From the addon', paste, frame='framed'), Section('Other ways in', ways)))


# ---- Spec support (SpecGrid): every damage spec on the Plate, "Not yet" until a parse is measured; healers and tanks open below
def specs() -> str:
    notes = json.loads((DATA_ROOT / 'data/curated/rotation-notes.json').read_text()) if (DATA_ROOT / 'data/curated/rotation-notes.json').exists() else {}
    top = top_raid_sets()
    rows = []
    best = max((v[0] for v in top.values()), default=0)
    damage = sorted((s for s in SPECS if s['role'] == 'dps'), key=lambda s: top.get(s['spec'], (0, 10 ** 6))[1])
    for s in damage:
        label = f"{s['name']} {CLASS_NAME[s['class_slug']]}"
        # SpecCard: the spec's name is the toggle for its rotation steps; no parses, so no Tag; the raid BiS set at 60,
        # its place and its bar, the unit once in the section's meta
        title = label   # the rotation reads on /sim: no toggle on this list
        fig = top.get(s['spec'])
        trail = (f'<span class="spec-fig"><span class="t-small num rk">#{fig[1]}</span><span class="sbar">{S.Meter(fig[0] / best)}</span>'
                 f'<span class="num dps">{fig[0]:,.1f}</span></span>') if fig else ''
        rows.append(Row(title, '', lead=S.crest_img(s['class_slug'], 28), trail=trail, title_color=CLASS_COLOR[s['class_slug']], link=False, slim=True))
    # UnsimulatedSpecCard: one wrapping line of links, each the crest and the spec in its class colour, to its guide
    others = ''.join(f'<a class="un-l" href="#" style="color:{CLASS_COLOR[s["class_slug"]]}">{S.crest_img(s["class_slug"], 20)}{s["name"]} {CLASS_NAME[s["class_slug"]]}<span class="t-small">{s["role"].title()}</span></a>'
                     for s in SPECS if s['role'] != 'dps')
    meta = ''
    heads = ('<div class="spec-heads" aria-hidden="true"><span class="rk">Rank</span><span class="sbar">Raid BiS at 60</span>'
             '<span class="dps">DPS</span></div>') if top else ''
    under = f'<div class="un"><h3 class="un-h">Healers and tanks</h3><div class="un-g">{others}</div></div>'
    return finish('Spec support', head('specs') + ledger(Section('Damage specs', heads + Rows(rows) + under, meta=meta, flush=True)))


def top_raid_sets() -> dict[str, tuple[float, int]]:
    """web lib/sim/bis-benchmark.ts topRaidSets over bis-benchmark-load.ts's bands: each damage spec's best level 60 raid set DPS
    (either faction), ranked high to low."""
    best = {}
    for s in SPECS:
        path = B / 'bis' / f"{s['spec']}.json"
        if s['role'] != 'dps' or not path.exists():
            continue
        bands = [b for b in json.loads(path.read_text())['bands']
                 if (b.get('role') or 'dps') == 'dps' and not b.get('set_dps_partial') and b.get('set_dps', 0) > 0
                 and b['band'] == LEVEL and (b.get('preset') or 'bare') == 'raid']
        if bands:
            best[s['spec']] = max(b['set_dps'] for b in bands)
    ranked = sorted(best.items(), key=lambda kv: -kv[1])
    return {spec: (dps, i + 1) for i, (spec, dps) in enumerate(ranked)}


TOOL_CSS = '''<style>
.col>p{margin:0}.col{display:flex;flex-direction:column;gap:22px;min-width:0}.rail-col{gap:22px}.col>.col{gap:22px}.paste-b{display:flex;flex-direction:column;gap:12px}
.fld{box-sizing:border-box;height:40px;padding:0 12px;font:600 14px/20px var(--font-body);color:var(--text);background:transparent;border:1px solid var(--line);border-radius:4px}
.fld.grow,.fl.grow{flex:1;min-width:220px}.fl.grow .fld{width:100%}.w24{width:96px}.w160,.fl.w160 .fld{width:160px}.w240{width:240px;text-align:left}
.tb{display:flex;flex-direction:column;align-items:flex-start}.tb .t-h4{padding-top:12px;margin:0;font-size:15px;line-height:20px}.tb p{padding-top:4px;margin:0}.tb .bt{padding-top:12px}.rail-col .tb{padding-top:11px}
.hits{margin-top:8px}.trunc{margin:0;font-size:12px}.hits-more{display:flex;flex-wrap:wrap;align-items:center;gap:0 12px;padding-top:8px}
.run-notes p{display:flex;flex-wrap:wrap;gap:0 12px;align-items:baseline;margin:0;padding-top:12px;font-size:14px;line-height:20px}
.tool-who{display:flex;flex-direction:column;gap:4px;padding:14px 0 4px}
.who-line{display:flex;flex-wrap:wrap;align-items:center;gap:8px 12px;margin:0;font-size:15px;line-height:24px}
.who-acts{margin-left:auto;display:flex;align-items:center;gap:16px}
.cands{display:flex;flex-direction:column;gap:20px}.cand-h{display:flex;flex-wrap:wrap;align-items:center;justify-content:space-between;gap:8px}
.lock,.opt{display:flex;align-items:center;gap:8px;min-height:44px;font-size:14px}.cb{width:20px;height:20px;margin:0;accent-color:var(--gold)}
.hover-name{all:unset;cursor:pointer}.num{font-variant-numeric:tabular-nums}
.search-f{display:flex;flex-wrap:wrap;align-items:flex-end;gap:12px}.search-f .fl{display:flex;flex-direction:column;gap:6px}.sw-l{display:flex;align-items:center;gap:8px;font-size:14px}
.hit{box-sizing:border-box;display:flex;align-items:center;gap:12px;min-height:44px;padding:6px 0;border-bottom:1px solid var(--line-soft,var(--line))}.hit:last-child{border-bottom:0}.hit-n{flex:1;min-width:0;font-size:15px;line-height:20px}
.wbody{padding-top:10px}.wbody p{margin:0}.opts{list-style:none;margin:0;padding:0}.opts.ce{padding-top:16px}.opts.two{display:grid;grid-template-columns:1fr 1fr;gap:0 16px}
.sets-f{display:flex;flex-wrap:wrap;align-items:flex-end;gap:12px;padding-top:19px}.pad{padding:16px}.kinds{display:flex;flex-wrap:wrap;gap:0 20px}
.src-g{display:flex;flex-direction:column;padding-top:8px}.src{display:flex;flex-direction:column}.indent{padding-left:28px}.grow-n{flex:1}.src p.indent{margin:0;line-height:19.5px}.src .disc>summary{display:flex;align-items:center;height:44px;box-sizing:border-box}.src-g .t-h4{margin:0;font-size:15px;line-height:20px}
.pn-b>.search-f{padding-bottom:16px}.spec-t{display:flex;flex-wrap:wrap;align-items:center;column-gap:16px}.spec-t .disc{font-weight:400}.rail-col .pn-b>.tb{padding-top:0}.pn-b>.opts.ce{padding-top:0}.pn-b .sets-f{padding-top:8px}.pn-b>.cands{padding-top:12px}
.alink{display:inline-flex;align-items:center;min-height:44px}.wt-pub{display:flex;flex-direction:column;gap:8px}.wt-pub p{margin:0;font-size:13px}.gated{display:flex;flex-direction:column;gap:4px;margin:0;padding-left:12px;list-style:none}.gated li{line-height:19.5px}.gb{padding-top:4px}
.wt-pub .st .stc{grid-template-columns:0 var(--st-name) var(--st-bar) repeat(var(--n),var(--st-fig))}@media(max-width:1023px){.wt-pub .st .stc{grid-template-columns:0 minmax(0,1fr) auto}}
.opts.inl{display:flex;flex-wrap:wrap;gap:0 24px}.opts.cols3{display:grid;grid-template-columns:1fr;gap:0 24px}@media(min-width:768px){.opts.cols3{grid-template-columns:1fr 1fr}}@media(min-width:1280px){.opts.cols3{grid-template-columns:1fr 1fr 1fr}}
.src-cols{display:grid;grid-template-columns:1fr;column-gap:32px}@media(min-width:768px){.src-cols{grid-template-columns:1fr 1fr}}@media(min-width:1280px){.src-cols{grid-template-columns:1fr 1fr 1fr}}.rep-n{padding-top:8px;font-size:14px;line-height:20px}.rep-t{display:flex;flex-wrap:wrap;gap:0 16px}.opt.sm{font-size:13px}
.tb-row{display:grid;gap:4px 24px;padding:12px 0;border-bottom:1px solid var(--line-soft,var(--line));align-self:stretch}.tb-row:last-of-type{border-bottom:0}@media(min-width:768px){.tb-row{grid-template-columns:260px minmax(0,1fr);align-items:start}}
.src-l{display:flex;align-items:center;gap:8px;min-height:44px}.src-l .opt{min-width:0}.bcount{all:unset;cursor:pointer;min-height:44px;display:inline-flex;align-items:center}.pl{padding-left:8px}.later{margin:0;padding-top:4px;font-size:13px}.gl2{margin:0;padding:4px 0;font-size:13px;line-height:20px}.gl1{display:flex;flex-wrap:wrap;align-items:center;gap:0 8px;min-height:44px;margin:0;font-size:14px;line-height:24px}.gl1 b{font-size:15px}.disc-b{padding:4px 0 8px}
.tb-t{display:grid;align-items:start;gap:12px 32px;min-width:0}@media(min-width:1024px){.tb-t{grid-template-columns:600px minmax(0,1fr)}}.tb-t>div{min-width:0}.tb-d{display:flex;flex-direction:column;margin:0;padding:0;list-style:none;font-size:13px;line-height:20px}.tb-gs{display:flex;flex-wrap:wrap;gap:8px 32px}.tb-g h4{margin:0;padding-bottom:4px}.tb-d .kill{color:var(--kill)}.tb-d .wipe{color:var(--wipe)}.tb-d span{color:var(--text)}
.spec-fig{display:flex;align-items:center;gap:12px}.spec-fig .rk{width:48px;text-align:right}.spec-heads{display:flex;justify-content:flex-end;align-items:center;gap:12px;padding:4px 16px 0;font:700 12px/16px var(--font-body);color:var(--muted)}.spec-heads .rk{width:48px;text-align:right}.spec-heads .dps{width:64px;text-align:right}.spec-fig .dps{width:64px;text-align:right;font-weight:700}.sbar{display:none;width:240px}@media(min-width:768px){.sbar{display:block}}.spec-n>summary{color:inherit;font-size:15px;font-weight:700;white-space:nowrap}.un{padding:8px 16px;border-top:1px solid var(--line-soft,var(--line))}.un-h{margin:0;padding-top:8px;font:600 15px/20px var(--font-body);color:var(--cream)}.un-g{display:grid;grid-template-columns:1fr;column-gap:20px}@media(min-width:768px){.un-g{grid-template-columns:repeat(4,1fr)}}.un-l{display:flex;align-items:center;gap:8px;min-height:44px;font-size:15px;font-weight:700;text-decoration:none;white-space:nowrap}.un-l .t-small{margin-left:auto}@media(min-width:768px){.un-l{display:inline-flex}.un-l .t-small{margin-left:0}}</style>'''


def finish(title: str, body: str) -> str:
    return build_page(title, body + S.PAGE_CSS + TOOL_CSS, 'Tools', None,
                      'Simulator tools mock (design/specs/2026-11-06-sim.md grammar), built only from ws_prims. Character: a level 60 Orc Fury Warrior entered by hand '
                      f'(Lionheart Helm, Arcanite Reaper); items, sources, consumables and stats from build {S.BUILD}.')


PAGES = {'web-simtools-gear': gear, 'web-simtools-gear-empty': gear_empty, 'web-simtools-talents': talents, 'web-simtools-drops': drops,
         'web-simtools-weights': weights, 'web-simtools-specs': specs}

if __name__ == '__main__':
    for name, fn in PAGES.items():
        if sys.argv[1:] and name not in sys.argv[1:]:
            continue
        print(write(f'{name}.html', fn()))
