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
from ws_prims import (Button as _Button, DropZone, Facts, Figure, Link, Meter, MeterClass, Notice, Page, PageHeader, Plate, RankList,
                      RankRow, ColumnHeads, Row, Rows, Section, Spark, StatRow, StatTable, TabStrip, Tag, build_page, control_bar, crest_img,
                      seg, sel, write)

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
MONTHS = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'June', 'July', 'Aug', 'Sept', 'Oct', 'Nov', 'Dec']


# A Button that navigates (sign in, the companion's step, the upload) is a link, as the build's is; one that acts on the
# page (a pairing code, Try again) is a button.
LinkButton = _Button


def Button(label: str, primary: bool = False, small: bool = False) -> str:
    return _Button(label, primary, small).replace('<a ', '<button type="button" ', 1).replace(' href="#"', '', 1).replace('</a>', '</button>')


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


FROM_LINE = 'From the game&rsquo;s Logs folder, up to 4 GB.'   # Upload.svelte FROM_LINE


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
        body += f'<span class="t-label">{escape(lc["pairLabel"])}</span><p class="t-small">{escape(COPY["pairingCopy"]["companionFirst"])}</p><div class="up-act">{Button("Get a pairing code", small=True)}</div>'
    return Section('The companion', body + '</div>', panel=True)


# ------------------------------------------------------------------------------------------------ /logs
def logs_page(state: str) -> str:
    """state: out | empty | some (directions b1-logs): the title band, Upload a log (the lead panel), Your reports, Recent public
    reports; the companion in the rail."""
    head = PageHeader(escape(COPY['logsBandCopy']['h1']))
    signed_in = state != 'out'
    main = Section('Upload a log', upload_body('out' if state == 'out' else 'in'), panel=True, frame='lead')
    if state == 'empty':
        main += Section('Your reports', f'<div class="empty"><p>{escape(COPY["myReportsCopy"]["empty"])}</p></div>', panel=True)
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
        inner = f'<div class="empty"><p>{escape(yc["empty"])}</p><div class="act">{LinkButton(mc["uploadAction"], primary=True)}{Link(yc["sampleAction"] + " &rsaquo;")}</div></div>'
    elif state == 'failed':
        inner = f'<div class="empty" role="alert"><p>{escape(mc["failed"])}</p><div class="act">{Button("Try again", small=True)}</div></div>'
    else:
        inner = Rows([report_row(r) for r in MINE['rows']])
    body = head + Page('ledger', Plate(inner, 'Your reports')) + PAGE_CSS
    return build_page('Your reports', body, 'Logs', ACCOUNT if state != 'out' else None, f'/reports mock, {state}. Data: web/src/fixtures/reports (production z2opt2x2gs5u).')


# ------------------------------------------------------------------------------------------------ rankings
# MUST 3 as ruled: the Warcraft Logs parse scale exactly (tokens.css --color-parse-*; grey is the -text variant, as
# percentileToken returns it). The percentile is floored, never rounded: rank 1 of 212 is 100 (gold), rank 2 is 99.53 and
# rank 3 99.06 (pink), rank 4 98.58 (orange). The shipped rankColor clamps at 99 and percentileToken rounds: bug B4.
PARSE_SCALE = ((100, '#e5cc80', '100'), (99, '#e268a8', '99+'), (95, '#ff8000', '95+'), (75, '#a335ee', '75+'), (50, '#0070ff', '50+'), (25, '#1eff00', '25+'))
PARSE_GREY = '#8c8c8c'


def rank_percentile(rank: int, total: int) -> float:
    return 100 - (rank - 1) / max(total, 1) * 100   # lib/rankings/row-view.ts rankPercentile


def pct_colour(rank: int, total: int) -> str:
    p = int(rank_percentile(rank, total))
    return next((c for floor, c, _ in PARSE_SCALE if p >= floor), PARSE_GREY)


def legend() -> str:
    steps = [(t, c) for _, c, t in PARSE_SCALE] + [(COPY['rankingsCopy']['legendUnderLast'], PARSE_GREY)]
    return '<p class="legend">' + escape(COPY['rankingsCopy']['legendLead']) + ''.join(f' <b style="color:{c}">{t}</b>' for t, c in steps) + '</p>'


def ranking_row(row: dict, total: int, newest: datetime, multi: bool, you: bool = False) -> str:
    p = row['player']
    sl = p['class'].lower()
    days = (newest - datetime.strptime(row['fought_at'][:10], '%Y-%m-%d')).days + 1
    guild = row.get('guild')
    sub = (f'{escape(guild["name"])} &middot; ' if guild else '') + f'{p["spec"]} {p["class"]} &middot; {row["size"]}-man'
    top = 900 - 12   # the board's first value (the stub's rank-1 row); the bar is value over it
    rank = f'<span style="color:{pct_colour(row["rank"], total)}">{row["rank"]}</span>'
    secs = row['duration_ms'] // 1000
    ilvl = str(row['item_level']) if row.get('item_level') else '&ndash;'   # G12: no board row carries an item level yet
    # a scored kill's Executed opens compare mode on that fight (shipped; lib/sim/execution.ts executionHref)
    figs = [(f'{row["value"]:,}', 'DPS', '', False), (ilvl, '', 'Item level', False), (f'<a href="#">{row["execution_score"] * 100:.0f}%</a>', '', 'Executed', False),
            (f'{secs // 60}:{secs % 60:02d}', '', 'Length', False)]
    rr = RankRow(rank, crest_img(sl, 32) if multi else '', CLASS_COLOR[sl], escape(p['name']), sub, row['value'] / top, figs, f'{days} day{"s" if days != 1 else ""} ago',
                 [('Report', True), ('Build', False)], you='You' if you else '', selected=you)
    # an unknown figure keeps its desktop column (a dash under its head) and is left out of the phone row's figure line
    return rr.replace('<span class="fg"><span class="lbl">Item level</span>&ndash;', '<span class="fg unk"><span class="lbl">Item level</span>&ndash;', 1)


def board(data: dict, own_key: str = '', off_page: bool = False) -> str:
    rows = data['rows']
    total = data['total']
    newest = max(datetime.strptime(r['fought_at'][:10], '%Y-%m-%d') for r in rows)
    multi = len({r['player']['class'] for r in rows}) > 1
    heads = ColumnHeads([('#', True)] + ([('', False)] if multi else []) + [('Character', False), ('Share of the top', False), ('DPS', True), ('Item level', True),
                         ('Executed', True), ('Length', True), ('Date', True), ('', False)])
    body = ''
    if off_page and data.get('character_rank'):
        cr = data['character_rank']
        body += f'<div class="pinned" aria-label="{escape(COPY["rankingsCopy"]["ownBest"])}">' + ranking_row({**cr['row'], 'rank': cr['rank']}, total, newest, multi, you=True) + '</div>'
    body += ''.join(ranking_row(r, total, newest, multi, you=(r['player']['key'] == own_key and not off_page)) for r in rows)
    pages = -(-total // data['per_page'])
    pager = f'<div class="pager"><span class="t-small">Page 1 of {pages}</span><span class="muted">Previous</span><a class="link" href="#">Next</a></div>' if pages > 1 else ''
    rl = RankList(heads, body, 4, links3=True)
    if not multi:   # K3 NO (ruled): a single-class board draws no crest; the class-coloured name carries the class
        rl = rl.replace('<div class="rl links3"', '<div class="rl links3 nocrest"', 1)
    return Plate(rl + pager, 'Rankings')


def rankings_bar(encounter: str, cls: str, spec: str, open_m: bool = False) -> str:
    rc = COPY['rankingsCopy']
    groups = [(rc['encounterLabel'], sel(encounter)), (rc['classLabel'], sel(cls)), (rc['specLabel'], sel(spec)), (rc['rulesetLabel'], sel('Normal' if cls != rc['everyClass'] else rc['allRulesets'])),
              (rc['metricLabel'], seg([('Damage', True, None, ''), ('Healing', False, None, '')])),
              (rc['boardLabel'], seg([('Characters', True, None, ''), ('Guilds', False, None, '')])),
              ('Filters', '<button type="button" class="sel" aria-expanded="false" aria-haspopup="dialog">Filters</button>')]   # the BiS Sources pattern
    return control_bar(groups, f'{encounter} &middot; {spec if spec != rc["everySpec"] else cls}', open_m=open_m, label='Ranking choices', labels_from=1920)


def rankings_page(state: str) -> str:
    """state: index | empty | board | you | off | unranked | bar."""
    rc, ec = COPY['rankingsCopy'], COPY['encounterPickerCopy']
    upd = f'Updated {day(UPDATED)}'
    tier = Notice(f'<span>{escape(rc["tierQuestion"])}</span> {Link(rc["tierLink"] + " &rsaquo;")}', tone='neutral')
    if state in ('index', 'empty'):
        head = PageHeader(escape(rc['h1']), meta=[escape(rc['jobSignedOut'])])   # the index reads no board, so no Updated stamp
        if state == 'index':
            rows = [Row(escape(e['name']), 'Characters and guilds, every class and spec', trail='<span class="link">Open the board &rsaquo;</span>') for e in FACTS['encounters']]
            plate = Section(ec['heading'], Plate(Rows(rows), ec['heading']), meta=f'{len(rows)} encounters ranked')
        else:
            plate = Section(ec['heading'], Plate(f'<div class="empty"><p>{escape(ec["noneYet"])}</p><p>{escape(ec["tryGuildsProgress"])}</p><div class="act">{Link(ec["openGuildsBoard"] + " &rsaquo;")}</div></div>', ec['heading']))
        body = head + f'<div data-region="summary">{tier}</div>' + Page('ledger', plate + f'<div class="section">{legend()}</div>')
        return build_page('Rankings', body + PAGE_CSS, 'Logs', None, f'/rankings mock, {state}. Data: rankings-fixture.ts encounters.')
    you = state in ('you', 'off', 'unranked')
    if you:
        data = FACTS['board_you_off_page' if state == 'off' else 'board_you']
        enc, cls, spec = 'Onyxia', 'Warrior', 'Fury'
        cr = data.get('character_rank')
        if state == 'unranked':
            ans = (f'<span class="t-label">{escape(rc["answerLabel"])}</span><span class="ans-who">{crest_img("warrior", 40)}<b style="color:{CLASS_COLOR["warrior"]}">Obnoxious Yell</b></span>'
                   f'<span class="t-small">{escape(FACTS["fns"]["answerUnranked"])}</span><span class="ans-links">{Link(COPY["rankingsEmptyCopy"]["action"] + " &rsaquo;")}</span>')
        else:
            n = cr['rank']
            ordn = f'{n}{"th" if 10 <= n % 100 <= 20 else {1: "st", 2: "nd", 3: "rd"}.get(n % 10, "th")}'
            ans = (f'<span class="t-label">{escape(rc["answerLabel"])}</span><span class="ans-who">{crest_img("warrior", 40)}<b style="color:{CLASS_COLOR["warrior"]}">Obnoxious Yell</b>'
                   f'{Figure(ordn, unit=escape(FACTS["fns"]["answerUnit"]), tone="gold", size="big")}</span>'
                   f'<span class="t-small">{escape(FACTS["fns"]["answerNote"])}</span><span class="ans-links">{Link(rc["action"].replace(" →", "") + " &rsaquo;")}</span>')
        head = PageHeader(enc, meta=[f'{rc["h1"]} &middot; Fury Warriors &middot; Normal ruleset', upd], answer=ans)
        plate = Section('Fury Warriors', board(data, FACTS['own_key'], off_page=state == 'off'), meta=f'{data["total"]} ranked, each character&rsquo;s best kill')
        if state == 'unranked':
            data = dict(data, character_rank=None)
            plate = Section('Fury Warriors', board(dict(data, rows=[r for r in data['rows'] if r['player']['key'] != FACTS['own_key']])), meta=f'{data["total"]} ranked, each character&rsquo;s best kill')
    else:
        data = FACTS['board_signed_out']
        enc, cls, spec = 'Warden Kelthas', rc['everyClass'], rc['everySpec']
        head = PageHeader(enc, meta=[f'{rc["h1"]} &middot; {rc["everyClassAndSpec"]} &middot; {rc["allRulesets"]}', upd])
        plate = Section(rc['everyClassAndSpec'], board(data), meta=f'{data["total"]} ranked')
    body = head + rankings_bar(enc, cls, spec, open_m=state == 'bar') + f'<div data-region="summary">{tier}</div>' + Page('ledger', plate + f'<div class="section"><div data-region="footer-note">{legend()}</div></div>')
    return build_page(f'{enc} rankings', body + PAGE_CSS, 'Logs', RANK_YOU if you else None, f'/rankings/<slug> mock, {state}. Data: rankings-fixture.ts boardFor().')


# ------------------------------------------------------------------------------------------------ /ratings
def ratings_page() -> str:
    sc, rc = COPY['scoreCopy'], COPY['ratingCopy']
    fs = FACTS['fns']['score']
    rp = FACTS['ratings_page']   # lib/score/ratings-page.ts: the sample, the weights and the page's prose
    sample = rp['sample']['components']
    score = rp['sample']['score']
    weights = [(w['component'], str(w['dps']), str(w['healer']), str(w['tank'])) for w in rp['weights']]
    trend = rp['trend']
    pc = rp['copy']
    maxes = FACTS['score']['component_max']
    total = sum(maxes.values())
    tiers = FACTS['score']['tiers']
    tcol = {'common': '#ffffff', 'uncommon': '#1eff00', 'rare': '#0070dd', 'epic': '#a335ee', 'legendary': '#ff8000'}   # --color-rarity-*, exactly
    tier = next(t for t, f in reversed(list(zip(['common', 'uncommon', 'rare', 'epic', 'legendary'], [0, 500, 1000, 1500, 2000]))) if score >= f)
    labels = {c['key']: c for c in sc['components']}
    gaps = sorted(((maxes[k] - v, k) for k, v in sample.items()), reverse=True)
    ans = (f'<span class="t-label">{escape(fs["samplePill"])} &middot; {escape(sc["explainerHeading"])}</span>'
           f'<div class="fig big score-fig" style="--c:{tcol[tier]}"><span class="v">{score:,}<small>of {total:,} &middot; {tier.capitalize()}</small></span></div>'
           f'<span class="t-small">{escape(sc["mostToGainLead"])} {labels[gaps[0][1]]["label"]}{escape(fs["mostToGainRest"].replace("88", str(gaps[0][0])))}</span>')
    head = PageHeader(escape(sc['explainerPageTitle']), meta=[escape(sc['explainerScoreCaption']), escape(sc['explainerRatingCaption'])], answer=ans,
                      tabs=TabStrip([sc['explainerHeading'], sc['explainerRatingHeading']], sc['explainerHeading']))
    prow = [StatRow('', f'<b>{escape(labels[k]["label"])}</b>', escape(labels[k]['moves']), [f'{v:,}', escape(fs['upTo'][k])], Meter(v / maxes[k]),
                    primary_phone=f'{v} {escape(fs["upTo"][k])}') for k, v in sample.items()]
    parts = StatTable([('', False), ('Part', False), ('Of its top', False), ('Points', True), ('Top', True)], prow, 2)
    s1 = Section(sc['explainerHeading'], f'<p class="t-lead">{escape(sc["explainerLead"])}</p><div class="parts" style="margin-top:16px">{Plate(parts, sc["explainerPartsHeading"])}</div>'
                 f'<div class="prose"><p>{escape(sc["explainerWhat"])}</p><p class="muted">{escape(sc["explainerLaunch"])}</p></div>', meta=escape(sc['explainerScoreCaption']))
    rfig = (f'<div class="figs" style="--c:3;margin-top:16px">{Figure(str(trend[-1]), "Latest fight", rc["outOf100"], "gold", "big")}'
            f'<div class="fig"><span class="t-label">Last {len(trend)} fights</span>{Spark(trend, "#e5b955", total=100)}<span class="t-small">Up {trend[-1] - trend[0]} since the first</span></div>'
            f'<div class="fig"><span class="t-label">Weakest part</span><span class="t-small" style="font-size:15px;color:var(--text)">Weakest part of a fight: {escape(rp["weakest"])}.</span></div></div>')
    s2 = Section(sc['explainerRatingHeading'], f'<p class="t-lead">{escape(sc["explainerRatingBridge"])} {escape(sc["explainerPreparationShared"])}</p>{rfig}'
                 f'<div class="prose"><p class="muted">{escape(pc["visibility"])}</p>' + ''.join(f'<p>{escape(p)}</p>' for p in pc['oneNumberFirst']) + '</div>',
                 meta=escape(sc['explainerRatingCaption']))
    comp = Rows([Row(escape(c['part']), escape(c['body']), link=False) for c in pc['components']], open_=True)
    s3 = Section('What each component measures', comp)
    wrows = [StatRow('', f'<b>{c}</b>', '', [d, h, t], primary_phone=f'{d} / {h} / {t}', bar=False) for c, d, h, t in weights]
    s4 = Section('Weights, by role', StatTable([('', False), ('Component', False), ('DPS', True), ('Healer', True), ('Tank', True)], wrows, 3, bar=False, open_=True)
                 + f'<div class="prose"><p class="muted">{escape(pc["roles"])}</p><p class="muted">{escape(pc["adjustable"])}</p></div>')
    s5 = Section('Ranked against your spec, or absolute', '<div class="prose">' + ''.join(f'<p>{escape(p)}</p>' for p in pc['rankedOrAbsolute']) + '</div>')
    s6 = Section('The cap', f'<div class="prose"><p>{escape(rc["cappedRule"])}</p></div>')
    trows = [Row(escape(t['tier'].capitalize()), escape(fs['tierShort'][t['tier']]), link=False, title_color=tcol[t['tier']], selected=t['tier'] == tier,
                 trail='<span class="t-small">The sample</span>' if t['tier'] == tier else '').replace('<div class="row', f'<div style="--c:{tcol[t["tier"]]}" class="row', 1)
             for t in tiers]
    rail = (Section(sc['tierColumn'] + 's', f'<div class="tiers">{Rows(trows, open_=True)}</div>' + f'<p class="t-small" style="margin-top:10px">{escape(sc["tiersNote"])}</p>')
            + Section(sc['explainerLogsHeading'], f'<p class="t-small" style="margin-top:10px">{escape(sc["explainerLogsBody"])}</p><p style="margin-top:12px">{LinkButton(sc["explainerLogsButton"])}</p>'))
    body = head + Page('split', s1 + s2 + s3 + s4 + s5 + s6, rail) + PAGE_CSS
    return build_page(sc['explainerPageTitle'], body, 'Logs', None, '/ratings mock. Data: lib/score/ratings-page.ts (sample score, trend, weights, prose); lib/score and lib/rating copy.')


PAGES = {
    'web-logs': lambda: logs_page('out'),
    'web-logs-empty': lambda: logs_page('empty'),
    'web-logs-you': lambda: logs_page('some'),
    'web-reports': lambda: reports_page('out'),
    'web-reports-empty': lambda: reports_page('empty'),
    'web-reports-list': lambda: reports_page('list'),
    'web-reports-failed': lambda: reports_page('failed'),
    'web-rankings': lambda: rankings_page('index'),
    'web-rankings-empty': lambda: rankings_page('empty'),
    'web-rankings-board': lambda: rankings_page('board'),
    'web-rankings-you': lambda: rankings_page('you'),
    'web-rankings-off': lambda: rankings_page('off'),
    'web-rankings-unranked': lambda: rankings_page('unranked'),
    'web-rankings-bar': lambda: rankings_page('bar'),
    'web-ratings': ratings_page,
}

if __name__ == '__main__':
    if sys.argv[1:] == ['--list']:
        print(' '.join(PAGES))
    else:
        only = sys.argv[1:]
        for name, fn in PAGES.items():
            if not only or name in only:
                print(write(f'{name}.html', fn()))
