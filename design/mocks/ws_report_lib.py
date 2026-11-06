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
FIX = ROOT / 'web/src/fixtures/report'
SAMPLE_DIR = ROOT / 'web/public/sample-report/sample'
MOCKDATA = Path(__file__).resolve().parent / 'data'
FACTS = json.loads((MOCKDATA / 'logs-facts.json').read_text())
COPY = FACTS['copy']
RATING = json.loads((MOCKDATA / 'logs-rating-fight3.json').read_text())
COACH = json.loads((ROOT / 'web/src/fixtures/coach/fury-model.json').read_text())
SPELL_NAME = {s['id']: s['name'] for s in json.loads((B / 'spells.json').read_text())}
SPELL_ICON = json.loads((B / 'spellicons.json').read_text())
POWER = {0: 'Mana', 1: 'Rage', 2: 'Focus', 3: 'Energy'}
ONYXIA = {'zone': "Onyxia's Lair", 'boss': 'Onyxia', 'size': 40, 'duration_ms': 232000}   # coach-stubs.ts ONYXIA
VIEWER = 'Player-4184-000000A1'   # SAMPLE: the signed-in owner of Baelgrim (coach-stubs.ts premium viewer)

# The report's navigation (spec 6.3, decision R3): one TabStrip. Below 1760 the first 7, from 1760 the first 13, the rest in More.
TABS = ['Summary', 'Damage done', 'Healing', 'Damage taken', 'Deaths', 'Mechanics', 'Timeline', 'Threat', 'Casts', 'Buffs',
        'Debuffs', 'Rating', 'Coach', 'Interrupts', 'Dispels', 'Resources', 'Events', 'Queries', 'Compare']
TAB_MD, TAB_XL = 7, 13
_ic: dict[str, str] = {}


# ------------------------------------------------------------------------------------------------ data
class Report:
    def __init__(self, base: Path, onyxia: bool = False):
        self.base = base
        self.meta = json.loads((base / 'meta.json').read_text())
        self.onyxia = onyxia
        self._s: dict[int, dict] = {}
        self.realms = {}
        for f in self.meta['fights']:
            pass

    def summary(self, i: int) -> dict:
        if i not in self._s:
            s = json.loads((self.base / f'fights/{i}/summary.json').read_text())
            if self.onyxia and i == 3:   # coach-stubs.ts onyxiaSummary: the boss named, the warriors Fury, 3:52
                s = json.loads(json.dumps(s).replace('Warden Kelthas', ONYXIA['boss']))
                s['duration_ms'] = ONYXIA['duration_ms']
                for r in s['roster'] + s['combatants']:
                    if r['guid'] in ('Player-4184-000000A1', 'Player-4184-000000A3'):
                        r['class'], r['spec'] = 'Warrior', 'Fury'
            self._s[i] = s
        return self._s[i]

    def fight(self, i: int) -> dict:
        f = dict(next(x for x in self.meta['fights'] if x['index'] == i))
        if self.onyxia and i == 3:
            f.update(name=ONYXIA['boss'], size=ONYXIA['size'], duration_ms=ONYXIA['duration_ms'], zone=ONYXIA['zone'])
        return f

    @property
    def title(self) -> str:
        t = self.meta['title']
        return t.replace('Sanguine Depths', ONYXIA['zone']) if self.onyxia else t

    @property
    def zone(self) -> str:
        return ONYXIA['zone'] if self.onyxia else self.meta['zone']


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


def mmss(ms: float) -> str:
    s = ms / 1000
    return f'{int(s // 60)}:{int(s % 60):02d}' if s >= 60 or s == int(s) else f'{s:.1f}s'


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


def friendlies(R: Report, s: dict, source: str) -> set[str]:
    if source == 'Your group':   # lib/report/group.ts: the pull's own `group` (engine 0.6.0), else the report's
        f = R.fight(s['fight_index']) if 'fight_index' in s else {}
        return set(f.get('group') or R.meta.get('group') or [])
    return {r['guid'] for r in s.get('roster', [])}


def group_dps(R: Report, s: dict, source: str) -> float:
    fr = friendlies(R, s, source)
    return sum(a['total'] for a in s['damage_done'] if a['guid'] in fr) / (s['duration_ms'] / 1000)


# ------------------------------------------------------------------------------------------------ frame
def Button(label: str, primary: bool = False, small: bool = False) -> str:
    return _Button(label, primary, small).replace('<a ', '<button type="button" ', 1).replace(' href="#"', '', 1).replace('</a>', '</button>')


def head(R: Report, fi: int, tab: str, source: str, viewer: bool = False, sample: bool = False, tabs: bool = True) -> str:
    f, s = R.fight(fi), R.summary(fi)
    date = f['start'][:10]
    y, m, d = date.split('-')
    months = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'June', 'July', 'Aug', 'Sept', 'Oct', 'Nov', 'Dec']
    nf = len(R.meta['fights'])
    tag = Tag(COPY['reportCopy']['sampleBadge'], 'muted') if sample else Tag(COPY['reportCopy']['notCreditedTag'], 'muted')
    meta = [f'{escape(R.zone)} &middot; {int(d)} {months[int(m) - 1]} {y} &middot; {nf} fights', tag]
    if sample:
        meta.append(escape(COPY['reportCopy']['sampleNote']))
    deaths = sum(1 for d in s['deaths'] if d['guid'] in friendlies(R, s, source))
    label = fight_label(f)
    ros = roster(s)
    if viewer and VIEWER in ros:
        me = ros[VIEWER]
        fig = Figure(n1(me['dps']), unit='DPS', tone='gold', size='big')
        cap = f'{pname(me)} &middot; {escape(me.get("spec", ""))} {escape(me.get("class", ""))} &middot; {n0(me["damage_done"])} damage'
    else:
        fig = Figure(n0(group_dps(R, s, source)), unit='DPS', size='big')
        cap = f'{"Your group" if source == "Your group" else "Group"} damage &middot; {deaths} death{"" if deaths == 1 else "s"}'
    ans = (f'<span class="t-label">{label}</span>{fig}<span class="t-small">{cap}</span>'
           f'<span class="ans-act">{Button("Sim this fight", small=True)}{Button("Copy link", small=True)}</span>')
    strip = TabStripMore([t for t in TABS if t != 'Coach' or viewer], tab, TAB_MD, TAB_XL - (0 if viewer else 1)) if tabs else ''
    return PageHeader(escape(R.title), meta=meta, answer=ans, tabs=strip)


def fight_label(f: dict) -> str:
    if f['kind'] == 'trash':   # "Trash · 12:43", never "Trash · Trash"
        return f'{f["name"]} &middot; {mmss(f["duration_ms"])}'
    return f'{f["name"]} &middot; {outcome(f)} &middot; {mmss(f["duration_ms"])}'


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
    f, s = R.fight(fi), R.summary(fi)
    fr = friendlies(R, s, source)
    phases = ' &middot; '.join(f'{escape(p["name"])} {mmss(p["start_ms"])}&ndash;{mmss(p["end_ms"])}' for p in s.get('phases', []))
    items = [('Pull', escape(f['name'])), ('Result', outcome(f)), ('Length', mmss(f['duration_ms'])),
             ('Players', str(len(fr))), ('Deaths', str(sum(1 for d in s['deaths'] if d['guid'] in fr)))]
    if f['kind'] == 'trash':
        items.append(('Mobs died', str(f.get('mob_deaths', 0))))
    if phases and not (R.onyxia and fi == 3):   # the Onyxia re-read keeps the fixture's 60 s phases: not shown against 3:52
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
    series = [(split_name(a['name'])[0], colour(ros.get(a['guid'])), a['series'], f'<span class="num">{a["total"] / dur:,.0f}</span> {unit}') for a in rows]
    total = sum(a['total'] for a in s[key] if a['guid'] in fr) / dur
    title = {'damage': 'Damage over the fight', 'healing': 'Healing over the fight', 'taken': 'Damage taken over the fight'}[metric]
    meta = f'Group {total:,.0f} {unit} over the whole fight'
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


PARSES_MISSING = ('The rankings could not be reached, so the Parse column is unknown; every other figure is from the log. '
                  '<a class="link" href="#">Try again</a>')


