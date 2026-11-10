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
# D1 (ruled 2026-11-07): every report mock reads a REAL Forever log: the bundled sample served at /sample-report/sample,
# since 2026-11-09 the 8 Nov 2026 dungeon night in Excavation Site: Wetlands (104 pulls, six boss kills), cut by
# api/cmd/samplecut with every name changed.
SAMPLE_DIR = ROOT / 'web/public/sample-report/sample'
MOCKDATA = Path(__file__).resolve().parent / 'data'
FACTS = json.loads(Path(os.environ.get('FS_LOGS_FACTS') or MOCKDATA / 'logs-facts.json').read_text())   # FS_LOGS_FACTS: the parity specs' test-time facts
COPY = FACTS['copy']
_COACH_GO = (ROOT / 'api/internal/coach/copy.go').read_text()
COACH_TRASH = _COACH_GO.split('NotReviewableTrash = "', 1)[1].split('"', 1)[0]   # api/internal/coach/copy.go, verbatim
SPELL_NAME = {s['id']: s['name'] for s in json.loads((B / 'spells.json').read_text())}
SPELL_ICON = json.loads((B / 'spellicons.json').read_text())
POWER = {0: 'Mana', 1: 'Rage', 2: 'Focus', 3: 'Energy'}
# FS_REPORT_SAMPLE=1 (the parity spec's test-time render): every report mock is drawn as /reports/sample, the page the
# spec serves on this same log (its Sample report Tag and note in place of Not credited).
AS_SAMPLE = os.environ.get('FS_REPORT_SAMPLE') == '1'
DEFAULT_PULL = 11      # the sample's default pull as the live page opens it (Saltspine #1, a kill, 0:39)
DEATHS_PULL = 20       # the trash pull where the whole group died
COMPARE_PULL = 40      # the second Saltspine kill, the whole group in it
VIEWER = 'Player-4184-00000010'   # SAMPLE (typed, marked): a signed-in owner of Fenwick Silverlode, for the Coach tab only
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
    "Fenwick-Silverlode-ClassicBetaPvP-" -> "Fenwick Silverlode" (the anonymiser's first and last name)."""
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


# The report's icon chain (lib/guild/report-spell-icons.ts spellIconStem): melee, the client's override, the id table, a
# same-named spell's icon, the school's base spell, a plain rune. Never an empty socket.
FALLBACKS = json.loads((ROOT / 'web/src/data/spell-icon-fallbacks.json').read_text())
ICON_BY_NAME: dict[str, str] = {}
for _sp in json.loads((B / 'spells.json').read_text()):
    if str(_sp['id']) in SPELL_ICON:
        ICON_BY_NAME.setdefault(_sp['name'], SPELL_ICON[str(_sp['id'])])
MELEE_NAMES = {'', 'Melee', 'Auto Attack'}


def _real(stem: str | None) -> str | None:
    return stem if stem and (B / 'icons' / f'{stem}.webp').exists() else None


def icon_stem(sid: int, name: str = '', school: int | None = None) -> str:
    if sid == 0 and name in MELEE_NAMES:
        return FALLBACKS['melee']
    by_school = FALLBACKS['schools'].get(str(school & -school)) if school and school > 0 else None
    return (_real(FALLBACKS['client'].get(str(sid))) or _real(SPELL_ICON.get(str(sid)))
            or _real(ICON_BY_NAME.get(name) if name else None) or by_school or (FALLBACKS['melee'] if sid == 0 else FALLBACKS['spell']))


def role_icon(role: str | None, size: int = 14) -> str:
    """The Players row's role icon (lib/report/role-icon.ts): the game's own icon, from the build's files."""
    stem = FALLBACKS['roles'].get(role or '')
    if stem is None:
        return ''
    label = {'tank': 'Tank', 'healer': 'Healer', 'dps': 'Damage'}[role]
    return _icon_img(stem, size).replace('class="ic" src', 'class="role-ic" src', 1).replace('alt=""', f'alt="{label}" title="{label}"', 1)


def _icon_img(name: str, size: int) -> str:
    p = B / 'icons' / f'{name}.webp'
    if name not in _ic:
        _ic[name] = 'data:image/webp;base64,' + base64.b64encode(p.read_bytes()).decode() if p.exists() else ''
    return f'<img class="ic" src="{_ic[name]}" alt="" width="{size}" height="{size}" style="width:{size}px;height:{size}px">'


def sicon(sid: int, size: int = 28, name: str = '', school: int | None = None) -> str:
    return _icon_img(icon_stem(sid, name, school), size)



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


RATE_TABS = {'Summary': ('damage_done', 'DPS'), 'Damage done': ('damage_done', 'DPS'), 'Healing': ('healing_done', 'HPS'),
             'Damage taken': ('damage_taken', 'DTPS')}


def figure(R: Report, fi: int, tab: str, source: str, player: str = '') -> str:
    """The band's one figure follows the tab (directions b1, lib/report/frame.ts reportFigure): group DPS, HPS or DTPS on the
    rate tabs, the player's own on their view, deaths on Deaths and Mechanics, none elsewhere."""
    s = R.summary(fi)
    fr = friendlies(R, s, source)
    if tab == 'Deaths' or (tab == 'Mechanics' and R.fight(fi)['kind'] != 'trash'):
        n = sum(1 for d in s['deaths'] if d['guid'] in fr)
        return Figure(f'{n:,}', unit='death' if n == 1 else 'deaths')
    if tab not in RATE_TABS:
        return ''
    key, unit = RATE_TABS[tab]
    rows = [r for r in in_scope(R, s, source) if not player or r['guid'] == player]
    total = sum(r[key] for r in rows) / (s['duration_ms'] / 1000)
    return Figure(rate(total), unit=unit if player else f'group {unit}')


def head_line(R: Report, fi: int, source: str) -> str:
    """"Zephras Isle · Pull 6, Trash · 12:43 · 177 mobs died" (lib/report/frame.ts reportMeta)."""
    f = R.fight(fi)
    if f['kind'] == 'trash':
        mobs = f.get('mob_deaths', 0)
        return f'{escape(R.zone)} &middot; Pull {fi}, {escape(f["name"])} &middot; {mmss(f["duration_ms"])} &middot; {plural(mobs, "mob")} died'
    n = len(in_scope(R, R.summary(fi), source))
    return (f'{escape(R.zone)} &middot; Pull {fi}, {escape(f["name"])} #{f.get("attempt", 1)} &middot; {outcome(f)} &middot; '
            f'{mmss(f["duration_ms"])} &middot; {plural(n, "player")}')


def head(R: Report, fi: int, tab: str, source: str, viewer: str = '', sample: bool = False, tabs: bool = True, player: str = '') -> str:
    """The title band (directions b1): the title, one meta line, the figure that follows the tab, the tab strip."""
    sample = sample or AS_SAMPLE
    rc = COPY['reportCopy']
    meta = [head_line(R, fi, source), Tag(rc['sampleBadge' if sample else 'notCreditedTag'], 'muted')]
    # the sample's why (launch polish) is the Sample report tag's title in ReportHead: no line of its own
    strip = TabStripMore([t for t in TABS if t != 'Coach' or viewer], tab, TAB_MD, TAB_XL - (0 if viewer else 1)) if tabs else ''
    return PageHeader(escape(R.title), meta=meta, answer=figure(R, fi, tab, source, player), tabs=strip)


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
    # Each zone's trash folds into one entry first (lib/report/frame.ts fightOptions): the night view over its pulls.
    zones: dict[str, int] = {}
    for f in fights:
        if f['kind'] == 'trash' and f.get('zone'):
            zones[f['zone']] = zones.get(f['zone'], 0) + 1
    trash = [(f'All trash in {escape(z)}', plural(n, 'pull'), '', '', False) for z, n in zones.items() if n > 1] + trash
    boss_group = ([('All boss pulls', plural(len(bosses), 'pull'), '', '', False)] + bosses) if bosses else []
    menu = Menu('Fights', [('Boss pulls', boss_group), (f'Trash &middot; {plural(len(killed), "pull")}', trash)])
    if not bosses:
        menu = menu.replace('<div class="mg" role="presentation">Boss pulls</div>',
                            '<div class="mg" role="presentation">Boss pulls</div><p class="mi-empty">No boss pulls in this report.</p>', 1)
    toggle = (f'<label class="mi-toggle"><input type="checkbox"> Show {plural(len(idle), "pull")} that killed nothing</label>' if idle else '')
    assert menu.endswith('</div>')
    return '<div class="pc menu-at">' + menu[:-len('</div>')] + toggle + '</div></div>'


TABLE_FILTERS = {'Damage done', 'Damage taken', 'Healing'}
SEARCHABLE_OVER = 15   # ReportBar.svelte: a Source list longer than this gets the filter field


def bar(R: Report, fi: int, tab: str, source: str, open_m: bool = False, compare: str = '') -> str:
    f = R.fight(fi)
    # The Source list is the scope's players; past 15 it gains a filter field, and the shipped Select draws a
    # button trigger with aria-expanded for it in place of the native combobox (lib/ui/listbox.ts).
    many = len(in_scope(R, R.summary(fi), source if source in SOURCES else 'Your group')) > SEARCHABLE_OVER
    src = f'<button type="button" class="sel" aria-expanded="false" aria-label="Source">{escape(source)}</button>' if many else sel(source)
    groups = [('Fight', sel(fight_label(f))), ('Source', src), ('Window', sel('Whole fight'))]
    if tab in TABLE_FILTERS:
        # The Ability list (report/ReportBar.svelte abilityOptions over the tab's scope): past 15 it gains a filter field,
        # and the shipped Select draws its button trigger, as the Source list does.
        s = R.summary(fi)
        fr = friendlies(R, s, source if source in SOURCES else 'Your group')
        key = {'Damage done': 'damage_done', 'Healing': 'healing', 'Damage taken': 'damage_taken'}.get(tab, 'damage_done')
        spells = {ab['spell_id'] for a in s[key] if a['guid'] in fr for ab in a['abilities']}
        ability = ('<button type="button" class="sel" aria-expanded="false" aria-label="Ability">Every ability</button>'
                   if len(spells) + 1 > SEARCHABLE_OVER else sel('Every ability'))
        groups += [('Target', sel('Every target')), ('Ability', ability),
                   ('Filters', '<button type="button" class="sel" aria-haspopup="dialog" aria-expanded="false">Filters</button>')]
    if tab == 'Threat':
        groups.append(('Enemy', sel('Every enemy')))
    if tab == 'Compare':
        groups = groups[:2] + [('Compare with', sel(compare)), ('Metric', sel('DPS'))]
    short = f'{escape(f["name"])} {mmss(f["duration_ms"])} &middot; {source}'
    return control_bar(groups, short, open_m=open_m, label='Report choices', labels_from=1920)


def this_pull(R: Report, fi: int, source: str = 'All friendlies', note: str = '') -> str:
    """This pull (directions b1, report/ThisPull.svelte): only what the band does not say, then one muted line (`note`: on the
    Summary of a pull with no rating, the honest state), then Sim this fight and Copy link."""
    f, s = R.fight(fi), R.summary(fi)
    items = [('Started', clock(f['start']))]   # the scope's count and deaths are the band's and the tables' (h-audit)
    items += [(escape(p['name']), f'{mmss(p["start_ms"])}&ndash;{mmss(p["end_ms"])}') for p in s.get('phases', [])]
    line = f'<p class="t-small muted pull-note">{escape(note)}</p>' if note else ''
    acts = f'<div class="pull-acts">{_Button("Sim this fight", small=True)}{Button("Copy link", small=True)}</div>'
    return Section('This pull', Facts(items) + line + acts, panel=True)


def chart(R: Report, fi: int, metric: str, source: str) -> str:
    s = R.summary(fi)
    fr = friendlies(R, s, source)
    ros = roster(s)
    key = {'damage': 'damage_done', 'healing': 'healing', 'taken': 'damage_taken'}[metric]
    rows = sorted([a for a in s[key] if a['guid'] in fr and a['total'] > 0], key=lambda a: -a['total'])[:5]
    dur = s['duration_ms'] / 1000
    unit = {'damage': 'DPS', 'healing': 'HPS', 'taken': 'taken / s'}[metric]
    # At rest a legend chip names its line only: the table beside the chart carries the figures (report/TimeChart.svelte).
    series = [(split_name(a['name'])[0], colour(ros.get(a['guid'])), a['series'], '') for a in rows]
    total = sum(a['effective'] for a in s[key] if a['guid'] in fr) / dur
    title = {'damage': 'Damage over the fight', 'healing': 'Healing over the fight', 'taken': 'Damage taken over the fight'}[metric]
    meta = f'{"Your group" if source == "Your group" else "Group"} {rate(total)} {unit}'
    if not rows:
        return Section(title, f'<p class="empty">Nothing in this window.</p>', meta=meta, panel=True)
    phases = [(p['name'], p['start_ms'] / 1000, p['end_ms'] / 1000) for p in s.get('phases', [])]
    marks = [(d['at_ms'] / 1000, '#f0736b') for d in s['deaths'] if d['guid'] in fr]
    lc = LineChart(series, dur, 180, phases, marks, label=f'{title}, per player', step_ms=1000)
    # The whole fight's figure is the band's answer: no head meta; the phone's fold reads Show chart, the figure after it.
    phone = Disclosure(f'Show chart <span class="t-small muted">{meta}</span>', LineChart(series, dur, 140, phases, marks, label=title, step_ms=1000))
    return Section(title, f'<div class="lc-only-d">{lc}</div><div class="lc-only-m">{phone}</div>', panel=True)


def notices(lines: list[str]) -> str:
    if not lines:
        return ''
    # The region is the page column's (the build's notices sit in the column under the full-bleed bar).
    inner = ''.join(Notice(l, tone='neutral').removeprefix('<div class="pc">').removesuffix('</div>') for l in lines)
    return f'<div class="pc"><div data-region="summary">{inner}</div></div>'


# Boss pulls only: one short line at every width, and Try again.
PARSES_MISSING = 'Parses didn&rsquo;t load. <a class="link" href="#">Try again</a>'


