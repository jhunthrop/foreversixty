"""Web-layout mocks: /logs, /reports (the list), /rankings, /rankings/<slug> and /ratings, composed ONLY from ws_prims.
The report itself (/reports/<id>, every tab, /reports/sample) is ws_report.py.

Spec: design/specs/2026-11-07-logs-reports.md. Data and copy, read at render time, never retyped:
  - design/mocks/data/logs-facts.json (node design/mocks/logs_facts.mjs): the page copy modules (lib/reports, lib/report,
    lib/rankings, lib/score, lib/rating) and the e2e rankings stubs (tests/e2e/support/rankings-fixture.ts);
  - web/src/fixtures/reports/{mine,recent}.json: production's one public report (the 17 Sept beta night) as the signed-in
    player's own list and the public feed (G8 folded; decision D1: no retail-shaped stub rows), the files the parity spec serves;
  - web/public/sample-report/sample/meta.json: the sample report row;
  - lib/score/ratings-page.ts (through logs_facts.mjs): the ratings page's own prose, weight table and sample score.
SAMPLE (typed, marked): none beyond what the stubs carry; the signed-in character on /rankings is the stub's own
"Obnoxious Yell" (Fury Warrior), the signed-in account on /logs and /reports is the stub's "Fixture#1234".

  node design/mocks/logs_facts.mjs && python3 design/mocks/ws_logs.py
  for p in $(python3 design/mocks/ws_logs.py --list); do node design/mocks/ws_shoot.mjs $p 390,1440,1920,2000; done
"""
import json
import os
import sys
from datetime import datetime
from html import escape
from pathlib import Path

from mocklib import CLASS_COLOR, PARSE
from ws_prims import (Button as _Button, DropZone, EmptyFill as _EmptyFill, EmptyState, Facts, Figure, Link, LoadError, Meter, MeterClass, Notice, Page,
                      PageHeader, Plate, RankList, RankRow, ColumnHeads, Row, Rows, Section, Spark, StatRow, StatTable, TabStrip, Tag, build_page,
                      control_bar, crest_img, seg, sel, write)

ROOT = Path(__file__).resolve().parents[2]
MOCKDATA = Path(__file__).resolve().parent / 'data'
FACTS = json.loads(Path(os.environ.get('FS_LOGS_FACTS') or MOCKDATA / 'logs-facts.json').read_text())   # FS_LOGS_FACTS: the parity specs' test-time facts
COPY = FACTS['copy']
# The signed-in list and the public feed: web/src/fixtures/reports (G8 folded, decision D1), the files parity-logs.spec.ts serves.
_FIX = ROOT / 'web/src/fixtures/reports'
_MINE = json.loads((_FIX / 'mine.json').read_text())
_RECENT = json.loads((_FIX / 'recent.json').read_text())
MINE = {'rows': [dict(r, first_pull=(r.get('started_at') or r['created_at'])[:10], created_at=r['created_at'][:10]) for r in _MINE['page']['rows']],
        'recent': [dict(r, status='complete', created_at=r['created_at'][:10], first_pull=r['created_at'][:10]) for r in _RECENT['page']['rows']]}
NOW = datetime.strptime(_MINE['now'][:10], '%Y-%m-%d')
SAMPLE_META = json.loads((ROOT / 'web/public/sample-report/sample/meta.json').read_text())
ACCOUNT = ('Fixture#1234', None, 'No characters yet')
RANK_YOU = ('Obnoxious Yell', 'warrior', 'Fury Warrior &middot; 60')
UPDATED = datetime.strptime(FACTS['board_you']['updated_at'][:10], '%Y-%m-%d')
MONTHS = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec']


# A Button that navigates (sign in, the companion's step, the upload) is a link, as the build's is; one that acts on the
# page (a pairing code, Try again) is a button.
LinkButton = _Button


def Button(label: str, primary: bool = False, small: bool = False) -> str:
    return _Button(label, primary, small).replace('<a ', '<button type="button" ', 1).replace(' href="#"', '', 1).replace('</a>', '</button>')


def split_lead(message: str) -> tuple[str, str]:
    """MyReports' designed empty: the first sentence leads, the rest says what fills it (the component's own split rule)."""
    lead, sep, fill = message.partition('. ')
    return (lead + '.', fill) if sep else (message, '')


def EmptyFill(message: str, action: str, secondary: str) -> str:
    """MyReports' designed empty: ws_prims.EmptyFill (ui/EmptyState's `fill` shape) on the component's own first-sentence split,
    the upload as the small primary link button and the sample as the gold link."""
    lead, fill = split_lead(message)
    return _EmptyFill(escape(lead), escape(fill), LinkButton(action, primary=True, small=True), Link(escape(secondary) + ' &rsaquo;'))


def StatePad(inner: str) -> str:
    """MyReports' own 16 px around its empty, error and skeleton states (`<div class="p-4">`), inside whatever panel holds it."""
    return f'<div class="st-pad">{inner}</div>'


def day(d: datetime) -> str:
    return f'{d.day} {MONTHS[d.month - 1]}'


def ago(iso_day: str) -> str:
    """The /reports age rule (spec 5): today, then {n} days ago inside two weeks, then the date."""
    d = datetime.strptime(iso_day, '%Y-%m-%d')
    days = (NOW - d).days
    if days < 1:
        return 'today'
    return f'{days} day{"s" if days != 1 else ""} ago' if days < 14 else day(d)


def report_title(r: dict) -> str:
    """The report's title, else its zone, else the night it logged (production's z2opt2x2gs5u carries neither title nor zone)."""
    d = datetime.strptime(r['first_pull'], '%Y-%m-%d')
    return r['title'] or r['zone'] or f'Report, {day(d)} {d.year}'


PAGE_CSS = '''<style>
/* the built Section's body sits 12 px under its hairline whatever it holds (ui/Section.svelte) */
section>.sec-h+*,section>.sec-h+.rows{margin-top:12px}.empty{line-height:20px}
.sr-only{position:absolute;width:1px;height:1px;padding:0;margin:-1px;overflow:hidden;clip:rect(0,0,0,0);white-space:nowrap;border:0}
.ans-who{display:flex;align-items:center;gap:8px 12px;flex-wrap:wrap}.ans-who b{font:700 18px/22px var(--font-body)}.ans-links{display:flex;gap:16px}
@media(min-width:900px){.head .ans{min-width:360px}.ans-who{justify-content:flex-end}.ans-links{justify-content:flex-end}}
.prose p{font-size:15px;line-height:24px;max-width:72ch;margin-top:12px}.prose p.muted{color:var(--muted)}
.up{display:flex;flex-direction:column;gap:12px;padding:18px 20px}.up.open{padding:4px 0 0}.up p{font-size:15px;line-height:24px;max-width:72ch}.up .t-small{max-width:72ch}
.up-act{display:flex;gap:16px;align-items:center;flex-wrap:wrap}
.up.bare{padding:0}.pn-b>p.empty{padding:8px 0}
code{font:13px/20px ui-monospace,Menlo,monospace;color:var(--strong)}
.legend{display:flex;flex-wrap:wrap;gap:4px 14px;align-items:baseline;font-size:13px;color:var(--muted)}.legend b{font:700 13px/20px var(--font-body)}
.pager{display:flex;gap:16px;align-items:center;justify-content:flex-end;padding:12px 16px;border-top:1px solid var(--line-soft);font-size:14px}
.pinned{border-bottom:1px solid var(--line)}
.fig.score-fig .v{color:var(--c)}
/* MUST 6 as ruled: the tier names in the game's item quality colours exactly (tokens.css --color-rarity-*), at 19 px 700 so
   Rare (#0070dd, 3.9:1) and Epic (#a335ee, 3.6:1) pass AA as large text; the selected tier is tinted in its own colour with a
   3 px rail of it, never the gold selection ground (which hid Epic). */
.tiers .row .ttl{font:700 19px/24px var(--font-body)}
.tiers .rows.open .row{padding-inline:14px}
.tiers .row.cur{background:color-mix(in srgb,var(--c) 12%,transparent);box-shadow:inset 3px 0 0 var(--c)}
/* K3 as ruled: a single-class board has no crest column */
.rl.nocrest .cols{grid-template-columns:28px var(--rr-name) var(--rr-bar) repeat(var(--n),var(--rr-fig)) var(--rr-gap) var(--rr-links)}
.rl.nocrest .rr .ld{display:none}
@media(max-width:1023px){.rl.nocrest .rr.cols{grid-template-columns:22px minmax(0,1fr) auto;grid-template-areas:"rk nm f1" "rk nm gp" ". xx xx" ". br br" ". lk lk"}}
@media(max-width:1023px){.rr .fg.unk{display:none}}
.parts .st-r .f.p{color:var(--strong)}
/* the built StatRow's name line is 20 px (ui/stat-table.css) */
.st-r .nm{line-height:20px}
.empty .act{display:flex;gap:16px;align-items:center;margin-top:12px;flex-wrap:wrap}
/* MyReports' 16 px around its states (the empty, error and fill shapes are ws_prims') */
.st-pad{padding:16px}
</style>'''


# ------------------------------------------------------------------------------------------------ shared pieces
# MUST 4 as ruled: a status Tag only while a report is not finished; a finished report ("complete", "ready") carries none.
UNFINISHED = {'live': 'Live', 'processing': 'Processing', 'uploading': 'Uploading'}


def plural(n: int, one: str, many: str = '') -> str:
    return f'{n} {one if n == 1 else (many or one + "s")}'


def report_row(r: dict) -> str:
    title = report_title(r)
    bits = ([escape(r['zone'])] if r['zone'] and r['zone'] != title else []) + [plural(r['fight_count'], 'pull'),
                                                                              plural(r['boss_count'], 'boss', 'bosses') if r['boss_count'] else 'Trash only']
    status = Tag(UNFINISHED[r['status']], 'muted') if r['status'] in UNFINISHED else ''
    return Row(escape(title), ' &middot; '.join(bits), trail=f'{status}<span class="t-small">{ago(r["created_at"])}</span>')


def sample_row() -> str:
    d = datetime.strptime(SAMPLE_META['created_at'][:10], '%Y-%m-%d')
    return Row(escape(SAMPLE_META['title']), f'{escape(SAMPLE_META["zone"])} &middot; {len(SAMPLE_META["fights"])} fights',
               trail=f'{Tag(COPY["reportCopy"]["sampleBadge"], "muted")}<span class="t-small">{day(d)} {d.year}</span>')


def recent_section(signed_in: bool = False) -> str:
    rc = COPY['recentReportsCopy']
    # lib/reports/recent.ts isPublicListable: the feed lists a report that names a zone or killed a boss (production's one
    # public report has neither, so the feed is its empty line and the sample row). Signed in, the viewer's own reports sit in
    # Your reports above and are left out here (directions b1).
    feed = [r for r in MINE['recent'] if r['kill_count'] > 0 or (r.get('zone') or '').strip()]
    own = {r['id'] for r in MINE['rows']} if signed_in else set()
    shown = [report_row(r) for r in feed if r['id'] not in own]
    body = Rows(shown + [sample_row()], open_=True)   # no public report yet: the sample row alone, no empty line over it
    return Section(rc['heading'], body, panel=True)   # no page lists every public report yet: no action


FROM_LINE = 'Up to 4 GB.'   # Upload.svelte FROM_LINE


def upload_body(state: str) -> str:
    """state: out | in. Upload.svelte's own words (directions b1): the drop zone's two lines, nothing over it."""
    drop = '<b>Drop <code>WoWCombatLog.txt</code> here</b>'
    if state == 'out':
        zone = DropZone(drop, f'<span class="t-small">{FROM_LINE} <span class="link">Choose a file</span></span>', off=True)
        act = (f'<p>Uploading needs a signed-in account; the report is filed under it.</p><div class="up-act">{LinkButton(COPY["myReportsCopy"]["signInAction"], primary=True)}'
               f'{Link(COPY["yourReportsPageCopy"]["sampleAction"] + " &rsaquo;")}</div>')
        return f'<div class="up bare">{zone}{act}</div>'
    # the shipped control: a visually hidden file input and its label (the input is the accessible control)
    zone = DropZone(drop, f'<span class="t-small">{FROM_LINE} <input id="upload-file" class="sr-only" type="file" accept=".txt,text/plain">'
                          '<label class="link" for="upload-file">Choose a file</label></span>')
    return f'<div class="up bare">{zone}</div>'


def companion_section(signed_in: bool) -> str:
    lc = COPY['logsCompanionCopy']
    body = f'<div class="up open"><p>{escape(lc["pointer"])}</p><div class="up-act">{LinkButton(lc["pointerLink"], primary=signed_in)}</div>'
    if signed_in:
        body += f'<span class="t-label">{escape(lc["pairLabel"])}</span><div class="up-act">{Button("Get a pairing code", small=True)}</div>'
    return Section('The companion', body + '</div>', panel=True)


# ------------------------------------------------------------------------------------------------ /logs
def logs_page(state: str) -> str:
    """state: out | empty | some (directions b1-logs): the title band, Upload a log (the lead panel), Your reports, Recent public
    reports; the companion in the rail."""
    head = PageHeader(escape(COPY['logsBandCopy']['h1']))
    signed_in = state != 'out'
    main = Section('Upload a log', upload_body('out' if state == 'out' else 'in'), panel=True, frame='lead')
    if state == 'empty':
        main += Section('Your reports', StatePad(EmptyState(escape(COPY['myReportsCopy']['empty']))), panel=True)   # MyReports' line empty
    elif state == 'some':
        main += Section('Your reports', Rows([report_row(r) for r in MINE['rows']], open_=True), action=Link('All &rsaquo;'), panel=True)
    main += recent_section(signed_in)
    rail = companion_section(signed_in)
    if not signed_in:
        rail += Section('Your reports', f'<p class="empty">{escape(COPY["myReportsCopy"]["signedOut"])}</p>', panel=True)
    body = head + Page('split', main, rail, panels=True) + PAGE_CSS
    who = ACCOUNT if signed_in else None
    return build_page('Logs', body, 'Logs', who, f'/logs mock, {state}. Data: web/src/fixtures/reports (production z2opt2x2gs5u), the sample meta, lib/reports copy.')


# ------------------------------------------------------------------------------------------------ /reports
def reports_page(state: str) -> str:
    """state: out | empty | list | failed (directions b1-reports): the title and, signed in, Upload a log; one panel of rows."""
    yc, mc = COPY['yourReportsPageCopy'], COPY['myReportsCopy']
    head = PageHeader(escape(yc['title']), actions=LinkButton(mc['uploadAction'], primary=True) if state in ('list', 'empty', 'failed') else '')
    if state == 'out':
        inner = f'<div class="empty"><p>{escape(yc["signedOut"])}</p><div class="act">{LinkButton(mc["signInAction"], primary=True)}{Link(yc["sampleAction"] + " &rsaquo;")}</div></div>'
    elif state == 'empty':
        inner = StatePad(EmptyFill(yc['empty'], mc['uploadAction'], yc['sampleAction']))   # MyReports' designed empty, signed in
    elif state == 'failed':
        inner = StatePad(LoadError(escape(mc['failed'])))
    else:
        inner = Rows([report_row(r) for r in MINE['rows']])
    body = head + Page('ledger', Plate(inner, 'Your reports')) + PAGE_CSS
    return build_page('Your reports', body, 'Logs', ACCOUNT if state != 'out' else None, f'/reports mock, {state}. Data: web/src/fixtures/reports (production z2opt2x2gs5u).')


PAGES = {
    'web-logs': lambda: logs_page('out'),
    'web-logs-empty': lambda: logs_page('empty'),
    'web-logs-you': lambda: logs_page('some'),
    'web-reports': lambda: reports_page('out'),
    'web-reports-empty': lambda: reports_page('empty'),
    'web-reports-list': lambda: reports_page('list'),
    'web-reports-failed': lambda: reports_page('failed'),
}

if __name__ == '__main__':
    if sys.argv[1:] == ['--list']:
        print(' '.join(PAGES))
    else:
        only = sys.argv[1:]
        for name, fn in PAGES.items():
            if not only or name in only:
                print(write(f'{name}.html', fn()))
