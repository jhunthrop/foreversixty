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
from ws_prims import Facts, Field, FormRow, Notice, Page, PageHeader, Plate, Row, Rows, Section, TabStrip, build_page, write
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

SCOPE = 'The simulator runs damage specs only. Healing and tanking specs are not simulated yet.'
LEADS = {
    'gear': 'Tick the items, enchants, talents and sets you want tried for your damage spec. Every valid combination is simulated and ranked against what you have on.',
    'talents': 'Your damage spec&rsquo;s build against every other build you have, ranked. Gear is locked to what you are wearing, so the only thing that changes is the tree.',
    'drops': 'Pick where you are going. Every item those bosses drop is simulated one at a time against your damage spec&rsquo;s current set, and the upgrades are listed by boss.',
    'weights': 'What one point of each stat is worth for your damage spec, for the addons that ask for a number.',
    'specs': ('The simulator is only worth as much as its numbers. A nightly job is built to sim the top 50 parses for each of the 20 damage specs '
              'it covers and publish the gap here, whatever it is. No real parses exist yet, so every damage spec below still reads "Not yet". '
              'The other 8 specs, healers and tanks, are not simulated yet.'),
}
TITLES = {'gear': 'Top gear', 'talents': 'Talents', 'drops': 'Droptimizer', 'weights': 'Weights', 'specs': 'Spec support'}


def ledger(*blocks: str) -> str:
    """The ledger page with the built Page's section gap between its blocks (the summary line included)."""
    return Page('ledger', f'<div class="col">{"".join(blocks)}</div>')


def split(main: list[str], rail: list[str]) -> str:
    return Page('split', f'<div class="col">{"".join(main)}</div>', f'<div class="col rail-col">{"".join(rail)}</div>')


def head(tool: str) -> str:
    return PageHeader(TITLES[tool], meta=[SCOPE], lead=LEADS[tool], tabs=TabStrip(TOOLS, TITLES[tool]))


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
    count = '' if tool == 'weights' else '<b>Nothing to try yet. Tick an item, a build or a source that fits your character.</b>'
    note = ('One run per stat at the chosen precision: 6 stats, about 60,000 runs.' if tool == 'weights'
            else 'Three stages: everything at 100 iterations, the survivors at 1,000, the finalists at 3,000.')
    return bar + f'<div class="pc run-notes"><p>{count}<span class="t-small">{note}</span></p></div>'


# ---- the character line (ToolCharacterLine, region `summary`)
def character_line() -> str:
    level = sum(ITEMS[i]['item_level'] for i in WORN.values()) / len(WORN)
    return (f'<div class="tool-who" data-region="summary"><p class="who-line">{S.crest_img("warrior", 24)}<span>Orc Fury Warrior &middot; Level {LEVEL}</span>'
            f'<span class="t-small">Item level {level:.1f}</span><span class="t-small">Entered by hand</span>'
            f'<span class="who-acts"><a class="link" href="#">Open in planner &rsaquo;</a>{btn("Change source")}</span></p>'
            f'<p class="t-small">24 talent points. Simmed as a level {LEVEL} with these talents.</p></div>')


# ---- Top gear: the candidates (SlotGrid, CandidateRows)
def candidate(slot_label: str, iid: int) -> str:
    it = ITEMS[iid]
    lead = check(it['name']) + S.gicon(it['icon'], 32, it['quality'])
    trail = f'<span class="t-small num">{it["item_level"]}</span>{btn("Copy and modify", expanded=False)}'
    title = f'<button type="button" class="hover-name">{item_name(iid)}</button>'
    return (f'<div class="cand-group"><div class="cand-h"><h3 class="t-h4">{slot_label}</h3><label class="t-small lock">{check()}Lock to equipped</label></div>'
            + Rows([Row(title, 'Equipped', lead=lead, trail=trail, link=False)], open_=True) + '</div>')


def candidates() -> str:
    return Section('Candidates', Plate('<div class="pad cands">' + candidate('Head', WORN['head']) + candidate('Main hand', WORN['main_hand']) + '</div>'))


def item_search() -> str:
    usable = sorted((i for i in ITEMS.values() if i.get('item_level')), key=lambda i: (-i['item_level'], i['name']))[:100]
    rows = ''.join(f'<div class="hit">{S.gicon(i["icon"], 28, i["quality"])}<span class="hit-n">{item_name(i["id"])}</span>'
                   f'<span class="t-small num">{i["item_level"]}</span>{btn("Add")}</div>' for i in usable)
    filters = (f'<div class="search-f"><input type="search" class="fld grow" placeholder="Name" aria-label="Name">'
               f'<label class="fl"><span class="t-label">Minimum item level</span><input type="number" class="fld w24" value="0"></label>'
               f'<label class="fl"><span class="t-label">Slot</span><select class="fld w160"><option>Any slot</option></select></label>'
               f'<span class="fl"><span class="t-label">Source</span><button type="button" class="fld w240 sel" aria-expanded="false">Anywhere</button></span>'
               f'<span class="sw-l"><button type="button" role="switch" aria-checked="true" class="sw"></button>Only items this character can equip</span></div>')
    return Section('Find an item', filters + f'<div class="hits">{rows}</div><p class="t-small trunc">Showing the first 100.</p>')


def consumables() -> str:
    table = json.loads((DATA_ROOT / 'sim/request/raid-consumes.json').read_text())['specs'][SPEC]
    ids = list(dict.fromkeys([r for r in table if r['from_level'] <= LEVEL][-1]['consumes']))
    rows = ''.join(f'<li><label class="opt">{check()}{x.replace("_", " ")}</label></li>' for x in ids)
    return Section('Try each', f'<p class="t-small">Try each of these as a candidate rather than a setting.</p><ul class="opts">{rows}</ul>')


def whole_sets() -> str:
    return Section('Whole sets', '<p class="t-small">A set replaces every slot at once. Paste a second export string to add one.</p>'
                   '<div class="sets-f"><label class="fl w160"><span class="t-label">Name</span><input class="fld" placeholder="Name"></label>'
                   '<label class="fl grow"><span class="t-label">Paste an export string</span><input class="fld" placeholder="Paste an export string"></label>'
                   '<button type="button" class="btn">Add set</button></div>')


def talent_builds(plated: bool = False) -> str:
    body = (f'<div class="tb"><label class="opt">{check()}Your current build</label><h3 class="t-h4">Your saved builds</h3>'
            f'<p class="t-small">Your saved builds could not be read; the rest of the page still works.</p><div class="bt">{btn("Add a build")}</div></div>')
    return Section('Talent builds', Plate(f'<div class="pad">{body}</div>') if plated else body)


def gear() -> str:
    return finish('Top gear', head('gear') + run_bar('gear') + split([character_line(), candidates()], [item_search(), consumables(), whole_sets(), talent_builds()]))


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
    groups = []
    for kind, label in kinds:
        srcs = [s for s in loot if s['kind'] == kind]
        shown = [s for s in srcs if is_open(s)]
        if not srcs:
            continue
        if not shown:
            lis = ''.join(f'<li class="t-small">{s["name"]}: {"Opens later" if s.get("opens") == "later" else "first raids"}</li>' for s in srcs)
            groups.append(f'<div class="src-g"><h3 class="t-h4">{label}</h3><p class="t-small">Nothing here has opened yet. Here is what is known about each:</p>'
                          f'<ul class="gated">{lis}</ul><div class="gb">{btn("Show unreleased content")}</div></div>')
            continue
        rows = ''
        for s in shown:
            bosses = s.get('bosses') or []
            every = '<span class="t-small">Every boss</span>' if bosses else ''
            rows += f'<div class="src"><label class="opt">{check()}<b class="grow-n">{s["name"]}</b>{every}<span class="t-small num">{len(s.get("items", []))}</span></label>'
            if kind == 'dungeon' and bosses:
                rows += f'<div class="indent"><details class="disc"><summary>Bosses <span class="t-small">{len(bosses)}</span></summary></details></div>'
            elif bosses:
                rows += ''.join(f'<label class="opt indent">{check()}<span class="grow-n">{x.get("name") or s["name"]}</span></label>' for x in bosses)
            if s.get('trash'):
                rows += '<p class="t-small indent">Trash and chests</p>'
            rows += '</div>'
        note = '<p class="t-small">Nothing has recorded this character&rsquo;s professions, so every profession is listed.</p>' if kind == 'crafted' else ''
        groups.append(f'<div class="src-g"><h3 class="t-h4">{label}</h3>{note}{rows}</div>')
    body = f'<div class="pad"><div class="kinds">{toggles}</div>{"".join(groups)}</div>'
    return finish('Droptimizer', head('drops') + run_bar('drops') + ledger(character_line(), Section('Sources', Plate(body)), '<p class="t-small">No source ticked yet.</p>'))


# ---- Weights: the caution, the stats to weigh in the rail; the scale factors arrive with a run
def weights() -> str:
    row = next(w for w in json.loads((B / 'stat-weights.json').read_text()) if w['spec'] == SPEC)
    stats = row.get('stats') or list(row['weights'])
    picks = ''.join(f'<li><label class="opt">{check(checked=True, disabled=i == 0)}{s}</label></li>' for i, s in enumerate(stats))
    caution = ('<p class="notice">A stat weight is a straight-line guess at something that is not a straight line: it holds near the gear you have now and stops '
               'holding as soon as a set bonus, a proc or a hit cap changes. Sim the actual items in Top Gear instead. These are here because addons want them. '
               '<a class="link" href="#">Open Top Gear &rsaquo;</a></p>')
    rail = Section('Stats to weigh', f'<div class="wbody"><p class="t-small">Normalised to: {stats[0]}</p><p class="t-small">Only the stats this sim models for the spec.</p><ul class="opts two">{picks}</ul></div>')
    return finish('Weights', head('weights') + run_bar('weights') + split([character_line(), caution], [rail]))


# ---- a tool with no character (SourceSwitcher): the addon export on the Plate, the other ways in open
def gear_empty() -> str:
    """SourceSwitcher, in the landing's own grammar (ws_sim.landing_out's Form, Input and Button)."""
    paste = Plate('<div class="paste">' + S.Form('From the addon', [('From the addon', S.Input('FS1:&hellip;', 'From the addon', grow=True), True)], S.Button('Load', primary=True))
                  + '<p class="t-small">Paste the export string from the Forever Sixty addon, or let the companion push it. <a class="link" href="#">Get the addon &rsaquo;</a></p></div>',
                  label='From the addon')
    form = lambda label, ph: S.Form(label, [('', S.Input(ph, label), False)], S.Button('Load'))
    ways = Rows([Row('From a build', 'Paste a planner link, saved or not, or a saved build id.', trail=form('From a build', 'foreversixty.gg/b/&hellip;'), link=False),
                 Row('From a logged fight', 'Paste a report link, or open a fight from a report and choose Sim this fight.', trail=form('From a logged fight', 'foreversixty.gg/reports/&hellip;?fight=2'), link=False),
                 Row('Your characters', 'Sign in with Battle.net to find your characters, with their gear and talents ready to sim.', trail=S.Button('Sign in with Battle.net'), link=False)], open_=True)
    return finish('Top gear', head('gear') + ledger(Section('From the addon', paste), Section('Other ways in', ways)))


# ---- Spec support (SpecGrid): every damage spec on the Plate, "Not yet" until a parse is measured; healers and tanks open below
def specs() -> str:
    notes = json.loads((DATA_ROOT / 'data/curated/rotation-notes.json').read_text()) if (DATA_ROOT / 'data/curated/rotation-notes.json').exists() else {}
    rows = []
    for s in [s for s in SPECS if s['role'] == 'dps']:
        label = f"{s['name']} {CLASS_NAME[s['class_slug']]}"
        foot = '<details class="disc"><summary>What it does</summary></details>' if notes.get(s['spec'], True) else ''
        rows.append(Row(label, '<span class="num">No parses yet</span><br>Not yet: no real parse has been checked against this spec.',
                        trail='<span class="tag muted">Not yet</span>', title_color=CLASS_COLOR[s['class_slug']], link=False, foot=foot))
    others = [Row(f"{s['name']} {CLASS_NAME[s['class_slug']]}", 'Not simulated yet: damage specs first; healers and tanks come later',
                  title_color=CLASS_COLOR[s['class_slug']], link=False) for s in SPECS if s['role'] != 'dps']
    return finish('Spec support', head('specs') + ledger(Section('Damage specs', Plate(Rows(rows))), Section('Healers and tanks', Rows(others, open_=True))))


TOOL_CSS = '''<style>
.col>p{margin:0}.col{display:flex;flex-direction:column;gap:48px;min-width:0}.rail-col{gap:32px}.col>.col{gap:48px}
.fld{box-sizing:border-box;height:40px;padding:0 12px;font:600 14px/20px var(--font-body);color:var(--text);background:transparent;border:1px solid var(--line);border-radius:4px}
.fld.grow,.fl.grow{flex:1;min-width:220px}.fl.grow .fld{width:100%}.w24{width:96px}.w160,.fl.w160 .fld{width:160px}.w240{width:240px;text-align:left}
.tb{display:flex;flex-direction:column;align-items:flex-start}.tb .t-h4{padding-top:12px;margin:0;font-size:15px;line-height:20px}.tb p{padding-top:4px;margin:0}.tb .bt{padding-top:12px}.rail-col .tb{padding-top:11px}
.hits{margin-top:8px;max-height:343px;overflow-y:auto}.trunc{padding-top:8px;margin:0;font-size:12px}
.run-notes p{display:flex;flex-wrap:wrap;gap:0 12px;align-items:baseline;margin:0;padding-top:12px;font-size:14px;line-height:20px}
.tool-who{display:flex;flex-direction:column;gap:4px;padding:14px 0 4px}
.who-line{display:flex;flex-wrap:wrap;align-items:center;gap:8px 12px;margin:0;font-size:15px;line-height:24px}
.who-acts{margin-left:auto;display:flex;align-items:center;gap:16px}
.cands{display:flex;flex-direction:column;gap:20px}.cand-h{display:flex;flex-wrap:wrap;align-items:center;justify-content:space-between;gap:8px}
.lock,.opt{display:flex;align-items:center;gap:8px;min-height:44px;font-size:14px}.cb{width:20px;height:20px;margin:0;accent-color:var(--gold)}
.hover-name{all:unset;cursor:pointer}.num{font-variant-numeric:tabular-nums}
.search-f{display:flex;flex-wrap:wrap;align-items:flex-end;gap:12px}.search-f .fl{display:flex;flex-direction:column;gap:6px}.sw-l{display:flex;align-items:center;gap:8px;font-size:14px}
.hit{display:flex;align-items:center;gap:10px;min-height:44px}.hit-n{flex:1;min-width:0}
.wbody{padding-top:10px}.wbody p{margin:0}.opts{list-style:none;margin:0;padding:0}.opts.two{display:grid;grid-template-columns:1fr 1fr;gap:0 16px}
.sets-f{display:flex;flex-wrap:wrap;align-items:flex-end;gap:12px;padding-top:19px}.pad{padding:16px}.kinds{display:flex;flex-wrap:wrap;gap:0 20px}
.src-g{display:flex;flex-direction:column;gap:4px;padding-top:16px}.src{display:flex;flex-direction:column;padding:4px 0}.indent{padding-left:28px}.grow-n{flex:1}.src p.indent{margin:0;line-height:19.5px}.src .disc>summary{display:flex;align-items:center;height:44px;box-sizing:border-box}.src-g .t-h4{margin:0;font-size:15px;line-height:20px}
.gated{display:flex;flex-direction:column;gap:4px;margin:0;padding-left:12px;list-style:none}.gated li{line-height:19.5px}.gb{padding-top:4px}</style>'''


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
