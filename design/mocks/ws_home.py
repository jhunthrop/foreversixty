"""Web-layout mock: home (/), every state, composed ONLY from ws_prims primitives. Spec: design/specs/2026-11-06-home.md.

States (design/mocks/web/home-parity.json): out (signed out), 34 (a level 34 leveler, no guild), 60 (a level 60 raider in a
guild with a raid planned), new (signed in, no character), loading (session hinted, /v1/me in flight), error (34, the
upgrades failed). Signed in, every figure is read from web/tests/e2e/fixtures/home/ (home_fixtures.py writes them), the
same files the parity spec routes; signed out, the example is the active build's Marksmanship Hunter BiS file, which the
built page reads too.

  python3 design/mocks/home_fixtures.py && python3 design/mocks/ws_home.py
  node design/mocks/ws_shoot.mjs web-home-out 390,1440,1920,2000
"""
import json
from datetime import datetime
from pathlib import Path

from mocklib import CLASS_COLOR
from ws_prims import (Button, Checklist, Figure, ItemRow, Link, Page, PageHeader, Plate, Row, RowGrid, Rows, Section,
                      Skeleton, Tag, build_page, crest_img, icon, write)
from ws_lib import REGION

ROOT = Path(__file__).resolve().parents[2]
BUILD = json.loads((ROOT / 'web/src/data/active-build.json').read_text())['build']
DATA = ROOT / 'data/builds' / BUILD
FIX = ROOT / 'web/tests/e2e/fixtures/home'
CLASSES = json.loads((ROOT / 'web/src/data/classes.json').read_text())
SPECS = json.loads((ROOT / 'data/curated/specs.json').read_text())
DATES = json.loads((ROOT / 'web/src/data/dates.json').read_text())
SLOT_LABELS = {'head': 'Head', 'neck': 'Neck', 'shoulder': 'Shoulder', 'back': 'Back', 'chest': 'Chest', 'wrist': 'Wrist', 'hands': 'Hands',
               'waist': 'Waist', 'legs': 'Legs', 'feet': 'Feet', 'finger1': 'Finger 1', 'finger2': 'Finger 2', 'trinket1': 'Trinket 1',
               'trinket2': 'Trinket 2', 'main_hand': 'Main hand', 'off_hand': 'Off hand', 'ranged': 'Ranged'}
EXAMPLE = ('hunter-marksmanship', 'hunter', 'Marksmanship Hunter', 'ranged')
UPGRADES_SHOWN = 3
HOME_CSS = '''
.sec-b{padding-top:12px}
.row{min-height:58px}
.ph-meta .t-small+.t-small::before{content:none}
.head-sk{display:flex;flex-direction:column;gap:12px;padding-block:20px 16px;min-height:118px;justify-content:flex-end}
.sum-sk{padding-block:14px 4px}
.plate.pad{padding:16px}
.sk-lines{display:flex;flex-direction:column;gap:12px}
.err{display:flex;flex-wrap:wrap;align-items:center;gap:12px;min-height:44px;font-size:14px;color:var(--muted)}
'''


def J(name: str) -> dict:
    return json.loads((FIX / name).read_text())


def short_date(iso: str) -> str:
    d = datetime.fromisoformat(iso.replace('Z', '+00:00'))
    return f'{d.day} {d.strftime("%b")}'


def sec(title: str, body: str, action: str = '', meta: str = '') -> str:
    return Section(title, f'<div class="sec-b">{body}</div>', Link(f'{action} &rsaquo;') if action else '', meta)


def summary(text: str, action: str = '', tone: str = 'gold') -> str:
    act = f'<a class="link" href="#">{action} &rsaquo;</a>' if action else ''
    return f'<div class="pc"><div data-region="{REGION["summary"]}"><p class="notice {tone}">{text}{act}</p></div></div>'


def item_html(items: dict, item_id: int, size: int = 40) -> tuple[str, int, str]:
    it = items.get(item_id) or {}
    q = it.get('quality', 1)
    return icon(it.get('icon', ''), size, {2: 'q2', 3: 'q3', 4: 'q4', 5: 'q5'}.get(q, '')), q, it.get('name', '')


def source_line(slot: dict) -> str:
    return f"{slot['source']} &middot; crafted" if slot.get('source_kind') == 'crafted' else slot['source']


# ------------------------------------------------------------------------------------------------ signed out
def class_picker() -> str:
    rows = []
    for c in CLASSES:
        specs = ', '.join(s['name'] for s in SPECS if s['class_slug'] == c['slug'])
        rows.append(Row(c['name'], specs, lead=crest_img(c['slug'], 44), title_color=c['color']))
    return sec('Pick your class', RowGrid(rows, 3))


def raids_open() -> str:
    raid = next((d for d in DATES if d.get('role') == 'raids'), None)
    return f"First raids {short_date(raid['iso'])}" if raid else ''


def example_plate() -> str:
    spec, cls, title, slot_name = EXAMPLE
    bis = json.loads((DATA / 'bis' / f'{spec}.json').read_text())
    items = {i['id']: i for i in json.loads((DATA / 'items' / f'{cls}.json').read_text())['items']}
    rows = []
    for b in sorted((b for b in bis['bands'] if b['faction'] == 'alliance'), key=lambda b: (b['band'], b.get('preset') != 'bare')):
        pick = next(s for s in b['slots'] if s['slot'] == slot_name and 'item_id' in s)
        ic, q, _ = item_html(items, pick['item_id'])
        label = f"{b['band']} to {b['band'] + 9}" if b['band'] < 60 else ('60 raid' if b.get('preset') == 'raid' else '60')
        rows.append(ItemRow(pick['item_name'], q, ic, source_line(pick), label, (f"{b['set_dps']:,.1f}", 'DPS', 'the set')))
    first, last = min(b['band'] for b in bis['bands']), max(b['band'] for b in bis['bands'])
    return sec(f'{title}, level {first} to {last}', Plate(''.join(rows), 'Best in slot by level'), 'Best in slot')


def how_to_start() -> str:
    return sec('How to start', Rows([
        Row('Sign in with Battle.net', 'Imports your characters with their level, gear and item level.', rank='1'),
        Row('Install the addon', 'Rates every item in your bags with your spec&rsquo;s weights.', rank='2'),
    ], open_=True))


def state_out() -> str:
    head = PageHeader('Play your class better.', meta=['Pick your class.', raids_open()], actions=Button('Sign in with Battle.net', True))
    return head + Page('ledger', class_picker()) + Page('split', example_plate(), how_to_start())


# ------------------------------------------------------------------------------------------------ signed in
def spec_slug(name: str) -> str:
    return name.lower().replace(' ', '-')


def upgrades_of(char: dict) -> tuple[list[tuple], dict, int, dict]:
    cls = char['class'].lower()
    bis = J(f"bis-{cls}-{spec_slug(char['spec'])}.json")
    band = bis['bands'][0]
    items = {i['id']: i for i in J(f'items-{cls}.json')['items']}
    ups = []
    for s in band['slots']:
        worn = char['build']['gear'].get(s['slot'])
        alt = next((a for a in s.get('alternatives') or [] if a['item_id'] == worn and a.get('dps_delta', 0) < 0), None)
        if alt:
            ups.append((-alt['dps_delta'], s, worn))
    ups.sort(key=lambda u: -u[0])
    return ups, items, band['band'], band


def band_words(band: int) -> str:
    return '60' if band >= 60 else f'{band} to {band + 9}'


def head_in(char: dict, ups: list | None, band: int) -> str:
    cls = char['class'].lower()
    meta = [f"Level {char['level']} {char['spec']} {char['class']}", f"{char['faction'].capitalize()} {char['race']}", f"Item level {char['item_level']}"]
    answer = ''
    if ups is not None:
        total = sum(u[0] for u in ups)
        answer = Figure(f'+{total:,.1f}', f'Upgrades at {band_words(band)}', 'DPS') + Link('Open best in slot &rsaquo;')
    return PageHeader(char['name'], CLASS_COLOR[cls], crest_img(cls, 56), meta, answer)


def upgrades_section(char: dict, ups: list, items: dict, band: int) -> str:
    rows = []
    for gain, s, worn in ups[:UPGRADES_SHOWN]:
        ic, q, _ = item_html(items, s['item_id'])
        worn_name = items.get(worn, {}).get('name', '')
        rows.append(ItemRow(s['item_name'], q, ic, source_line(s), SLOT_LABELS[s['slot']], (f'+{gain:,.1f}', 'DPS', f'vs your {worn_name}')))
    return sec('Your upgrades', Plate(''.join(rows), 'Your upgrades'), f'All {len(ups)} on best in slot', ('Level 60' if band >= 60 else f'Levels {band_words(band)}'))


def talents_to_spend(char: dict) -> int:
    return max(0, char['level'] - 9 - sum(char['build']['talents']['points']))


def talents_section(char: dict) -> str:
    n = talents_to_spend(char)
    if n == 0:
        return ''
    spent = sum(char['build']['talents']['points'])
    return sec('Talents', Rows([Row(f'{n} talent points to spend', f"Level {char['level']} earns {char['level'] - 9}. You have spent {spent}.")], open_=True))


def sim_section(sims: dict) -> str:
    r = sims['rows'][0]
    return sec('Simulator', Rows([Row(r['title'], f"Simmed {short_date(r['created_at'])}", trail=f'<span class="g-sm">{Figure(f"{r["dps"]:,.1f}", "", "DPS", size="sm")}</span>')], open_=True), 'Open the simulator')


def logs_section(reports: dict) -> str:
    if not reports['rows']:
        rows = [Row('Upload a log', 'The companion uploads your logs as you play.')]
    else:
        rows = [Row(r['title'], f"{r['zone']} &middot; {r['boss_count']} bosses &middot; {short_date(r['created_at'])}") for r in reports['rows'][:3]]
    return sec('Your logs', Rows(rows, open_=True), 'All logs')


def setup_section(me: dict, char: dict) -> str:
    return sec('Set up', Checklist([
        ('ok', 'Battle.net import', f"Imported {short_date(me['bnet_imported_at'])}", None),
        ('ok', 'Addon export', f"Exported {short_date(char['build']['captured_at'])}", None),
    ]))


def guild_section(char: dict) -> str:
    g = char['guild']
    att = g['attendance']
    return sec('Guild', Rows([Row(g['name'], f"{g['rank']} &middot; {att['present']} of {att['nights']} raids")], open_=True))


def raid_section(home: dict) -> str:
    r = home['next_raid']
    d = datetime.fromisoformat(r['starts_at'].replace('Z', '+00:00'))
    when = f"Wed {d.day - 1} {d.strftime('%b')} &middot; 8 pm CST"
    counts = f"{r['counts']['in']} in, {r['counts']['late']} late"
    row = Row(r['title'], f'{when} &middot; {counts}', trail=Tag(r['own']['state'].capitalize()))
    needs = Checklist([('warn', n, '', None) for n in home['standing']['needs_before_next_raid']])
    return sec('Next raid', Rows([row], open_=True) + needs, 'Raid page')


def next_step(char: dict, ups: list | None, home: dict | None) -> tuple[str, str]:
    if home and home['next_raid']['own']['state'] != 'in':
        r = home['next_raid']
        return (f"Next: {r['title']} is on Wed 9 Dec and you are {r['own']['state'].capitalize()}. Confirm your sign-up.", 'Open the raid')
    n = talents_to_spend(char)
    top = ups[0][1] if ups else None
    if n:
        then = f", then get {top['item_name']}" if top else ''
        return (f'Next: spend your {n} talent points{then}.', 'Open the planner')
    if top:
        return (f"Next: get {top['item_name']} from {top['source']}.", 'Open best in slot')
    return ('Nothing waiting.', '')


def state_in(level: str, failed: bool = False) -> str:
    me = J(f'me-{level}.json')
    char = me['characters'][0]
    ups, items, band, _ = upgrades_of(char)
    home = J('guild-home-60.json') if 'guild' in char else None
    text, action = next_step(char, None if failed else ups, home)
    if failed:
        plate = Plate('<div class="err" role="alert"><span>We could not load your upgrades.</span><button type="button" class="btn sm">Try again</button></div>', 'Your upgrades')
        upgrades = sec('Your upgrades', plate.replace('class="plate"', 'class="plate pad"'))
    else:
        upgrades = upgrades_section(char, ups, items, band)
    main = (raid_section(home) if home else '') + upgrades + talents_section(char) + sim_section(J(f'sims-{level}.json'))
    rail = (guild_section(char) if home else '') + logs_section(J(f'reports-{level}.json')) + setup_section(me, char)
    return head_in(char, None if failed else ups, band) + summary(text, action) + Page('split', main, rail)


def state_new() -> str:
    head = PageHeader('Play your class better.', meta=['Signed in as Hunthrop#1894'])
    rows = Rows([
        Row('Check again', 'Battle.net lists a character after you play it and log out.', trail='<button type="button" class="btn primary sm">Check again</button>', link=False),
        Row('Paste an addon export', 'Your gear and talents, straight from the game.'),
        Row('Get the addon', 'Rates every item in your bags with your spec&rsquo;s weights.'),
    ], open_=True)
    return (head + summary('No character yet. Battle.net lists a character after you play it and log out.')
            + Page('ledger', sec('Get your character in', rows) + '<div class="gap-sec"></div>' + class_picker()))


def state_loading() -> str:
    head = f'<div class="pc"><div data-region="{REGION["head"]}" class="head-sk" role="status" aria-busy="true">{Skeleton(56, 56, True)}{Skeleton(280, 28)}</div></div>'
    sk = f'<div class="pc"><div data-region="{REGION["summary"]}"><div class="sum-sk">{Skeleton(520, 20)}</div></div></div>'
    lines = '<div class="sk-lines" role="status" aria-busy="true">' + ''.join(Skeleton(900, 42) for _ in range(3)) + '</div>'
    plate = Plate(lines, 'Your upgrades').replace('class="plate"', 'class="plate pad"')
    return head + sk + Page('split', sec('Your upgrades', plate), '')


STATES = {
    'out': (state_out, None),
    '34': (lambda: state_in('34'), ('Felsnap', 'hunter', 'Level 34')),
    '60': (lambda: state_in('60'), ('Thandrel', 'warrior', 'Level 60')),
    'new': (state_new, ('Hunthrop', 'warrior', 'No character')),
    'loading': (state_loading, ('Felsnap', 'hunter', 'Level 34')),
    'error': (lambda: state_in('34', failed=True), ('Felsnap', 'hunter', 'Level 34')),
}


def main() -> None:
    for state, (render, who) in STATES.items():
        body = render() + f'<style>{HOME_CSS}.gap-sec{{height:var(--section-gap)}}</style>'
        write(f'web-home-{state}.html', build_page('Forever Sixty', body, who=who))


if __name__ == '__main__':
    main()
