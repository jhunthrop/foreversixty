"""Data and frame helpers for ws_report.py (the report mocks): the fixture readers, names, icons, the page head, the control
bar, the chart and the rail's This pull. Split out to keep each generator under 800 lines."""
import base64
import json
import os
import statistics
import sys
from html import escape
from pathlib import Path

from mocklib import CLASS_COLOR, PARSE
from ws_prims import (Button as _Button, Disclosure, EmptyCrest, Facts, Figure, Lanes, LanesLegend, LineChart, Link, Menu, Meter,
                      MeterClass, Notice, Page, PageHeader, Paperdoll, Plate, Row, Rows, Section, Skeleton, Spark, StatRow,
                      StatTable, TabStripMore, Tag, build_page, control_bar, crest_img, seg, sel, write)

ROOT = Path(__file__).resolve().parents[2]
DATA_ROOT = Path(os.environ.get('FS_DATA_ROOT', ROOT))
BUILD = json.loads((DATA_ROOT / 'web/src/data/active-build.json').read_text())['build']
B = DATA_ROOT / 'data/builds' / BUILD
# D1 (ruled 2026-11-07): every report mock reads a REAL Forever log. The only one reachable is the 17 Sept 2026 beta night on
# Zephras Isle (client 1.60.1, 79 trash pulls, anonymised names), committed as the bundled sample and served at
# /sample-report/sample; production's one public report (z2opt2x2gs5u) is the same log's first 42 pulls. It has no boss pull.
SAMPLE_DIR = ROOT / 'web/public/sample-report/sample'
MOCKDATA = Path(__file__).resolve().parent / 'data'
FACTS = json.loads(Path(os.environ.get('FS_LOGS_FACTS') or MOCKDATA / 'logs-facts.json').read_text())   # FS_LOGS_FACTS: the parity specs' test-time facts
COPY = FACTS['copy']
_COACH_GO = (ROOT / 'api/internal/coach/copy.go').read_text()
COACH_TRASH = _COACH_GO.split('NotReviewableTrash = "', 1)[1].split('"', 1)[0]   # api/internal/coach/copy.go, verbatim
SPELL_NAME = {s['id']: s['name'] for s in json.loads((B / 'spells.json').read_text())}
SPELL_ICON = json.loads((B / 'spellicons.json').read_text())
POWER = {0: 'Mana', 1: 'Rage', 2: 'Focus', 3: 'Energy'}
DEFAULT_PULL = 6       # the sample's default pull as the live page opens it (Trash, 12:43, 177 mobs died)
EMPTY_PULL = 13        # 4 s, nobody of the group in it: the pull with no data
COMPARE_PULL = 3       # the other pull the whole group is in
VIEWER = 'Player-4184-00000038'   # SAMPLE (typed, marked): a signed-in owner of Soren Everhart, for the Coach tab only
# The parse ladder, exactly lib/report/format.ts percentileToken over web/src/styles/tokens.css --color-parse-*.
PARSE_TOKENS = {'gold': '#e5cc80', 'pink': '#e268a8', 'orange': '#ff8000', 'purple': '#a335ee', 'blue': '#0070ff', 'green': '#1eff00', 'grey': '#8c8c8c'}

# The report's navigation (spec 6.3, decision R3): one TabStrip. Below 1760 the first 7, from 1760 the first 13, the rest in More.
TABS = ['Summary', 'Damage done', 'Healing', 'Damage taken', 'Deaths', 'Mechanics', 'Timeline', 'Threat', 'Casts', 'Buffs',
        'Debuffs', 'Rating', 'Coach', 'Interrupts', 'Dispels', 'Resources', 'Events', 'Queries', 'Compare']
TAB_MD, TAB_XL = 7, 13
_ic: dict[str, str] = {}


def parse_colour(p: float) -> str:
    """percentileToken: grey under 25, green 25, blue 50, purple 75, orange 95, pink 99, gold at 100 (floored, so 99.5 is pink)."""
    p = int(p)
    for floor, key in ((100, 'gold'), (99, 'pink'), (95, 'orange'), (75, 'purple'), (50, 'blue'), (25, 'green')):
        if p >= floor:
            return PARSE_TOKENS[key]
    return PARSE_TOKENS['grey']


# ------------------------------------------------------------------------------------------------ data
class Report:
    def __init__(self, base: Path = SAMPLE_DIR):
        self.base = base
        self.meta = json.loads((base / 'meta.json').read_text())
        self._s: dict[int, dict] = {}

    def summary(self, i: int) -> dict:
        if i not in self._s:
            self._s[i] = json.loads((self.base / f'fights/{i}/summary.json').read_text())
        return self._s[i]

    def fight(self, i: int) -> dict:
        return next(x for x in self.meta['fights'] if x['index'] == i)

    @property
    def title(self) -> str:
        return self.meta['title']

    @property
    def zone(self) -> str:
        return self.meta['zone']


def split_name(full: str) -> tuple[str, str]:
    """lib/report/names splitUnitName: "Baelgrim-Nightslayer" -> Baelgrim, Nightslayer; the sample's
    "Soren-Everhart-ClassicBetaPvP-" -> "Soren Everhart" (the anonymiser's first and last name)."""
    parts = [p for p in full.split('-') if p]
    if len(parts) >= 3:
        return f'{parts[0]} {parts[1]}', parts[2]
    return (parts[0], parts[1]) if len(parts) == 2 else (full, '')


def roster(s: dict) -> dict[str, dict]:
    return {r['guid']: r for r in s.get('roster', [])}


def slug(r: dict | None) -> str | None:
    c = (r or {}).get('class')
    return c.lower() if c and c.lower() in CLASS_COLOR else None


def crest(r: dict | None, px: int) -> str:
    sl = slug(r)
    return crest_img(sl, px) if sl else EmptyCrest(px)


def colour(r: dict | None) -> str:
    sl = slug(r)
    return CLASS_COLOR[sl] if sl else 'var(--strong)'


def pname(r: dict | None, name: str = '') -> str:
    n = split_name(name or (r or {}).get('name', ''))[0]
    return f'<b style="color:{colour(r)}">{escape(n)}</b>'


def realm_tag(full: str, home: str) -> str:
    realm = split_name(full)[1]
    return realm if realm and realm != home and not realm.startswith('ClassicBeta') else ''


def spell(sid: int, logged: str = '') -> str:
    """The build's name for a spell (spells.json), the log's own name when the build has none; "Spell #n" stays as logged."""
    if logged and not logged.startswith('Spell #'):
        return logged
    return SPELL_NAME.get(sid) or logged or f'Spell #{sid}'


def sicon(sid: int, size: int = 28) -> str:
    name = SPELL_ICON.get(str(sid))
    p = B / 'icons' / f'{name}.webp' if name else None
    if p is None or not p.exists():
        return f'<span class="ic" style="width:{size}px;height:{size}px"></span>'
    if name not in _ic:
        _ic[name] = 'data:image/webp;base64,' + base64.b64encode(p.read_bytes()).decode()
    return f'<img class="ic" src="{_ic[name]}" alt="" width="{size}" height="{size}" style="width:{size}px;height:{size}px">'


def n0(v: float) -> str:
    return f'{v:,.0f}'


def n1(v: float) -> str:
    return f'{v:,.1f}'


def rate(v: float) -> str:
    """A per-second figure: one decimal under 100 (a levelling pull's 1.8 DPS), whole above (1,374 DPS)."""
    return f'{v:,.1f}' if v < 100 else f'{v:,.0f}'


def mmss(ms: float) -> str:
    s = int(ms // 1000)
    return f'{s // 60}:{s % 60:02d}'


def at(ms: float) -> str:
    s = ms / 1000
    return f'{int(s // 60)}:{s % 60:04.1f}'


def clock(iso: str) -> str:
    """The pull's start as a time of day. The page shows the viewer's own clock; the mocks render UTC, as the captures did."""
    from datetime import datetime, timezone
    return datetime.fromisoformat(iso.replace('Z', '+00:00')).astimezone(timezone.utc).strftime('%H:%M:%S')


def outcome(f: dict) -> str:
    if f['kind'] == 'trash':
        return 'Trash'
    return 'Kill' if f.get('kill') else 'Wipe'


def plural(n: int, one: str, many: str = '') -> str:
    return f'{n:,} {one if n == 1 else (many or one + "s")}'


SOURCES = ('All friendlies', 'Your group', 'Everyone in range')


def friendlies(R: Report, s: dict, source: str) -> set[str]:
    if source == 'Your group':   # lib/report/group.ts: the pull's own `group` (engine 0.6.0), else the report's
        f = R.fight(s['fight_index'])
        return set(f.get('group') or R.meta.get('group') or [])
    return {r['guid'] for r in s.get('roster', [])}


def in_scope(R: Report, s: dict, source: str) -> list[dict]:
    fr = friendlies(R, s, source)
    return [r for r in s['roster'] if r['guid'] in fr]


def group_dps(R: Report, s: dict, source: str) -> float:
    fr = friendlies(R, s, source)
    return sum(a['effective'] for a in s['damage_done'] if a['guid'] in fr) / (s['duration_ms'] / 1000)


# ------------------------------------------------------------------------------------------------ frame
def Button(label: str, primary: bool = False, small: bool = False) -> str:
    return _Button(label, primary, small).replace('<a ', '<button type="button" ', 1).replace(' href="#"', '', 1).replace('</a>', '</button>')


def answer(R: Report, fi: int, source: str, viewer: str = '') -> str:
    """Decision R1 as ruled: a trash pull leads with its mobs and length ("177 mobs · 12:43"), never a DPS figure. A boss pull
    (none in the real log: blocked on D1) leads with the viewer's DPS and parse when their character is in it, else the group's DPS."""
    f, s = R.fight(fi), R.summary(fi)
    nd = sum(1 for d in s['deaths'] if d['guid'] in friendlies(R, s, source))
    scope = 'Your group' if source == 'Your group' else ('Everyone in range' if source == 'Everyone in range' else 'Group')
    n = len(in_scope(R, s, source))
    if f['kind'] == 'trash':
        label = f'Trash pull &middot; started {clock(f["start"])}'
        fig = Figure(f'{f.get("mob_deaths", 0):,}', unit=f'{"mob" if f.get("mob_deaths") == 1 else "mobs"} &middot; {mmss(f["duration_ms"])}', size='big')
        cap = f'{scope} &middot; {plural(n, "player")} &middot; {"no deaths" if nd == 0 else plural(nd, "death")}'
    else:   # spec 6.1; not rendered from real data until a real boss pull exists (D1)
        label = f'{escape(f["name"])} #{f.get("attempt", 1)} &middot; {"Kill" if f.get("kill") else "Wipe"} &middot; {mmss(f["duration_ms"])}'
        fig = Figure(n0(group_dps(R, s, source)), unit='DPS', size='big')
        cap = f'{scope} damage &middot; {plural(nd, "death")}'
    acts = (f'<span class="ans-act">{Button("Sim this fight", small=True)}{Button("Copy link", small=True)}</span>'
            f'<span class="ans-more"><button type="button" class="btn sm" aria-haspopup="menu" aria-expanded="false" aria-label="Report actions">&middot;&middot;&middot;</button></span>')
    return f'<span class="t-label">{label}</span>{fig}<span class="t-small">{cap}</span>{acts}'


def head(R: Report, fi: int, tab: str, source: str, viewer: str = '', sample: bool = False, tabs: bool = True) -> str:
    f = R.fight(fi)
    y, m, d = f['start'][:10].split('-')
    months = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'June', 'July', 'Aug', 'Sept', 'Oct', 'Nov', 'Dec']
    nf = len(R.meta['fights'])
    tag = Tag(COPY['reportCopy']['sampleBadge'], 'muted') if sample else Tag(COPY['reportCopy']['notCreditedTag'], 'muted')
    meta = [f'{escape(R.zone)} &middot; {int(d)} {months[int(m) - 1]} {y}<span class="ph-fights"> &middot; {nf} fights</span>', tag]
    if sample:
        meta.append(f'<span class="ph-note">{escape(COPY["reportCopy"]["sampleNote"])}</span>')
    strip = TabStripMore([t for t in TABS if t != 'Coach' or viewer], tab, TAB_MD, TAB_XL - (0 if viewer else 1)) if tabs else ''
    return PageHeader(escape(R.title), meta=meta, answer=answer(R, fi, source, viewer), tabs=strip)


def fight_label(f: dict) -> str:
    if f['kind'] == 'trash':   # "Trash · 12:43", never "Trash · Trash"
        return f'{f["name"]} &middot; {mmss(f["duration_ms"])}'
    return f'{f["name"]} #{f.get("attempt", 1)} &middot; {outcome(f)} &middot; {mmss(f["duration_ms"])}'


def fight_menu(R: Report, fi: int) -> str:
    """The Fight menu (R2 as ruled): Boss pulls numbered per boss ("{boss} #2 · Wipe 34%"); trash rows that killed nothing sit
    behind a toggle; "1 mob died", never "1 mobs died". The real log has no boss pull, so its Boss pulls group is the empty line."""
    fights = R.meta['fights']
    bosses = [(fight_label(f), f'{clock(f["start"])} &middot; {plural(f.get("deaths", 0), "death")}', '', '', f['index'] == fi) for f in fights if f['kind'] == 'encounter']
    killed = [f for f in fights if f['kind'] == 'trash' and f.get('mob_deaths', 0) > 0]
    idle = [f for f in fights if f['kind'] == 'trash' and not f.get('mob_deaths', 0)]
    trash = [(fight_label(f), f'{clock(f["start"])} &middot; {plural(f["mob_deaths"], "mob")} died', '', '', f['index'] == fi) for f in killed]
    boss_group = ([('All boss pulls', plural(len(bosses), 'pull'), '', '', False)] + bosses) if bosses else []
    menu = Menu('Fights', [('Boss pulls', boss_group), (f'Trash &middot; {plural(len(killed), "pull")}', trash)])
    if not bosses:
        menu = menu.replace('<div class="mg" role="presentation">Boss pulls</div>',
                            '<div class="mg" role="presentation">Boss pulls</div><p class="mi-empty">No boss pulls in this report.</p>', 1)
    toggle = (f'<label class="mi-toggle"><input type="checkbox"> Show {plural(len(idle), "pull")} that killed nothing</label>' if idle else '')
    assert menu.endswith('</div>')
    return '<div class="pc menu-at">' + menu[:-len('</div>')] + toggle + '</div></div>'


TABLE_FILTERS = {'Damage done', 'Damage taken', 'Healing'}


def bar(R: Report, fi: int, tab: str, source: str, open_m: bool = False, compare: str = '') -> str:
    f = R.fight(fi)
    groups = [('Fight', sel(fight_label(f))), ('Source', sel(source)), ('Window', sel('Whole fight'))]
    if tab in TABLE_FILTERS:
        groups += [('Target', sel('Every target')), ('Ability', sel('Every ability')),
                   ('Filters', '<span class="sel" role="button" aria-haspopup="dialog">Filters</span>')]
    if tab == 'Threat':
        groups.append(('Enemy', sel('Every enemy')))
    if tab == 'Compare':
        groups = groups[:2] + [('Compare with', sel(compare)), ('Metric', sel('DPS'))]
    short = f'{escape(f["name"])} {mmss(f["duration_ms"])} &middot; {source}'
    return control_bar(groups, short, open_m=open_m, label='Report choices', labels_from=1920)


def this_pull(R: Report, fi: int, source: str = 'All friendlies') -> str:
    """Decision R7 as ruled: a trash pull reads Mobs died, never "Result: Trash"; Result and Phases are a boss pull's."""
    f, s = R.fight(fi), R.summary(fi)
    fr = friendlies(R, s, source)
    items = [('Pull', escape(f['name']))]
    if f['kind'] != 'trash':
        items.append(('Result', outcome(f)))
    items += [('Length', mmss(f['duration_ms']))]
    if f['kind'] == 'trash':
        items.append(('Mobs died', f'{f.get("mob_deaths", 0):,}'))
    items += [('Players', f'{len(in_scope(R, s, source)):,}'), ('Deaths', str(sum(1 for d in s['deaths'] if d['guid'] in fr)))]
    phases = ' &middot; '.join(f'{escape(p["name"])} {mmss(p["start_ms"])}&ndash;{mmss(p["end_ms"])}' for p in s.get('phases', []))
    if phases:
        items.append(('Phases', phases))
    return Section('This pull', Facts(items), meta=f'Started {clock(f["start"])}')


def chart(R: Report, fi: int, metric: str, source: str) -> str:
    s = R.summary(fi)
    fr = friendlies(R, s, source)
    ros = roster(s)
    key = {'damage': 'damage_done', 'healing': 'healing', 'taken': 'damage_taken'}[metric]
    rows = sorted([a for a in s[key] if a['guid'] in fr and a['total'] > 0], key=lambda a: -a['total'])[:5]
    dur = s['duration_ms'] / 1000
    unit = {'damage': 'DPS', 'healing': 'HPS', 'taken': 'taken / s'}[metric]
    series = [(split_name(a['name'])[0], colour(ros.get(a['guid'])), a['series'], f'<span class="num">{rate(a["effective"] / dur)}</span> {unit}') for a in rows]
    total = sum(a['effective'] for a in s[key] if a['guid'] in fr) / dur
    title = {'damage': 'Damage over the fight', 'healing': 'Healing over the fight', 'taken': 'Damage taken over the fight'}[metric]
    meta = f'{"Your group" if source == "Your group" else "Group"} {rate(total)} {unit} over the whole fight'
    if not rows:
        return Section(title, f'<p class="empty">Nothing in this window.</p>', meta=meta)
    phases = [(p['name'], p['start_ms'] / 1000, p['end_ms'] / 1000) for p in s.get('phases', [])]
    marks = [(d['at_ms'] / 1000, '#f0736b') for d in s['deaths'] if d['guid'] in fr]
    lc = LineChart(series, dur, 180, phases, marks, label=f'{title}, per player', step_ms=1000)
    phone = Disclosure(f'<span class="t-small">{meta}</span>', LineChart(series, dur, 140, phases, marks, label=title, step_ms=1000))
    return Section(title, f'<div class="lc-only-d">{lc}</div><div class="lc-only-m">{phone}</div>', meta=f'<span class="lc-only-d">{meta}</span>')


def notices(lines: list[str]) -> str:
    if not lines:
        return ''
    return '<div data-region="summary">' + ''.join(Notice(l, tone='neutral') for l in lines) + '</div>'


# Boss pulls only. The phone reads one line (ruled SHOULD): the short form; 768 px and up the full sentence.
PARSES_MISSING = ('<span class="lc-only-d">The rankings could not be reached, so the Parse column is unknown; every other figure is from the log.</span>'
                  '<span class="lc-only-m">Parses unknown: the rankings did not answer.</span> <a class="link" href="#">Try again</a>')


