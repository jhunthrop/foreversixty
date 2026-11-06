"""Web-layout mocks: /logs, /reports (the list), /rankings, /rankings/<slug> and /ratings, composed ONLY from ws_prims.
The report itself (/reports/<id>, every tab, /reports/sample) is ws_report.py.

Spec: design/specs/2026-11-07-logs-reports.md. Data and copy, read at render time, never retyped:
  - design/mocks/data/logs-facts.json (node design/mocks/logs_facts.mjs): the page copy modules (lib/reports, lib/report,
    lib/rankings, lib/score, lib/rating) and the e2e rankings stubs (tests/e2e/support/rankings-fixture.ts);
  - design/mocks/data/logs-reports-stub.json: the /v1/reports?mine=1 rows reports-page.spec.ts serves (gap G8);
  - web/public/sample-report/sample/meta.json: the sample report row;
  - web/src/pages/ratings.astro: the ratings page's own prose, weight table and sample score (parsed from the source).
SAMPLE (typed, marked): none beyond what the stubs carry; the signed-in character on /rankings is the stub's own
"Obnoxious Yell" (Fury Warrior), the signed-in account on /logs and /reports is the stub's "Fixture#1234".

  node design/mocks/logs_facts.mjs && python3 design/mocks/ws_logs.py
  for p in $(python3 design/mocks/ws_logs.py --list); do node design/mocks/ws_shoot.mjs $p 390,1440,1920,2000; done
"""
import json
import re
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
FACTS = json.loads((MOCKDATA / 'logs-facts.json').read_text())
COPY = FACTS['copy']
MINE = json.loads((MOCKDATA / 'logs-reports-stub.json').read_text())
SAMPLE_META = json.loads((ROOT / 'web/public/sample-report/sample/meta.json').read_text())
RATINGS_SRC = (ROOT / 'web/src/pages/ratings.astro').read_text()
ACCOUNT = ('Fixture#1234', None, 'No characters yet')
RANK_YOU = ('Obnoxious Yell', 'warrior', 'Fury Warrior &middot; 60')
UPDATED = datetime.strptime(FACTS['board_you']['updated_at'][:10], '%Y-%m-%d')
MONTHS = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'June', 'July', 'Aug', 'Sept', 'Oct', 'Nov', 'Dec']


def Button(label: str, primary: bool = False, small: bool = False) -> str:
    return _Button(label, primary, small).replace('<a ', '<button type="button" ', 1).replace(' href="#"', '', 1).replace('</a>', '</button>')


def day(d: datetime) -> str:
    return f'{d.day} {MONTHS[d.month - 1]}'


def ago(hours: float) -> str:
    if hours < 24:
        return 'today' if hours >= 1 else 'just now'
    days = int(hours // 24)
    return f'{days} day{"s" if days != 1 else ""} ago'


PAGE_CSS = '''<style>
.ans-who{display:flex;align-items:center;gap:8px 12px;flex-wrap:wrap}.ans-who b{font:700 18px/22px var(--font-body)}.ans-links{display:flex;gap:16px}
@media(min-width:900px){.head .ans{min-width:360px}.ans-who{justify-content:flex-end}.ans-links{justify-content:flex-end}}
.prose p{font-size:15px;line-height:24px;max-width:72ch;margin-top:12px}.prose p.muted{color:var(--muted)}
.up{display:flex;flex-direction:column;gap:12px;padding:18px 20px}.up.open{padding:16px 0 0}.up p{font-size:15px;line-height:24px;max-width:72ch}.up .t-small{max-width:72ch}
.up-act{display:flex;gap:16px;align-items:center;flex-wrap:wrap}
code{font:13px/20px ui-monospace,Menlo,monospace;color:var(--strong)}
.legend{display:flex;flex-wrap:wrap;gap:4px 14px;align-items:baseline;font-size:13px;color:var(--muted)}.legend b{font:700 13px/20px var(--font-body)}
.pager{display:flex;gap:16px;align-items:center;justify-content:flex-end;padding:12px 16px;border-top:1px solid var(--line-soft);font-size:14px}
.pinned{border-bottom:1px solid var(--line)}
.score-fig .v{color:var(--c)}
.tiers .row .ttl{color:var(--c)}
.parts .st-r .f.p{color:var(--strong)}
.empty .act{display:flex;gap:16px;align-items:center;margin-top:12px;flex-wrap:wrap}
</style>'''


# ------------------------------------------------------------------------------------------------ shared pieces
def report_row(r: dict) -> str:
    sub = ' &middot; '.join([escape(r['zone']), f'{r["boss_count"]} bosses', f'{r["fight_count"]} fights', 'Not credited'])
    status = '' if r['status'] in ('', 'complete') else Tag(r['status'].capitalize(), 'muted')
    return Row(escape(r['title']), sub, trail=f'{status}<span class="t-small">{ago(r["hours_ago"])}</span>')


def sample_row() -> str:
    d = datetime.strptime(SAMPLE_META['created_at'][:10], '%Y-%m-%d')
    return Row(escape(SAMPLE_META['title']), f'{escape(SAMPLE_META["zone"])} &middot; {len(SAMPLE_META["fights"])} fights',
               trail=f'{Tag(COPY["reportCopy"]["sampleBadge"], "muted")}<span class="t-small">{day(d)} {d.year}</span>')


def recent_section() -> str:
    rc = COPY['recentReportsCopy']
    body = (f'<p class="empty">{escape(rc["empty"])}</p>' if not MINE['recent_public'] else '') + Rows([sample_row()], open_=True)
    return Section(rc['heading'], body, action=Link('All public reports &rsaquo;'))


def upload_body(state: str) -> str:
    """state: out | in. Upload.svelte's own words; the drop zone is the one dashed surface."""
    intro = ('<p>A whole <code>WoWCombatLog.txt</code>. Your report appears as soon as the parse starts, usually within a minute or two; '
             'the fights fill in as they are read.</p><p class="t-small">The game writes it to its <code>Logs</code> folder, beside <code>Interface</code> '
             'and <code>WTF</code>. Up to 4 GB. Type <code>/combatlog</code> again to stop logging first, so the file is whole.</p>')
    if state == 'out':
        zone = DropZone('<span>Drop <code>WoWCombatLog.txt</code> here, or</span>', '<span class="link muted">choose a file</span>', off=True)
        act = (f'<p>Uploading needs a signed-in account; the report is filed under it.</p><div class="up-act">{Button(COPY["myReportsCopy"]["signInAction"], primary=True)}'
               f'{Link(COPY["yourReportsPageCopy"]["sampleAction"] + " &rsaquo;")}</div>')
        return f'<div class="up">{intro}{zone}{act}</div>'
    zone = DropZone('<span>Drop <code>WoWCombatLog.txt</code> here, or</span>', '<a class="link" href="#">choose a file</a>')
    return f'<div class="up{" open" if state == "open" else ""}">{intro}{zone}</div>'


def companion_section(signed_in: bool) -> str:
    lc = COPY['logsCompanionCopy']
    body = f'<div class="up open"><p>{escape(lc["pointer"])}</p><div class="up-act">{Button(lc["pointerLink"])}</div>'
    if signed_in:
        body += f'<span class="t-label">{escape(lc["pairLabel"])}</span><p class="t-small">{escape(COPY["pairingCopy"]["companionFirst"])}</p><div class="up-act">{Button("Get a pairing code", small=True)}</div>'
    return Section('The companion', body + '</div>', meta='Live, while you raid')


# ------------------------------------------------------------------------------------------------ /logs
def logs_page(state: str) -> str:
    """state: out | empty | some."""
    bc = COPY['logsBandCopy']
    answer = ''
    if state == 'some':
        r = MINE['rows'][0]
        answer = (f'<span class="t-label">Your latest report &middot; {ago(r["hours_ago"])}</span>'
                  f'{Figure(escape(r["title"]), size="big")}<span class="t-small">{escape(r["zone"])} &middot; {r["boss_count"]} bosses &middot; {r["fight_count"]} fights</span>'
                  f'<span class="ans-links">{Link("Open it &rsaquo;")}</span>')
    head = PageHeader(escape(bc['h1']), meta=[escape(bc['jobLine'])], answer=answer)
    if state == 'out':
        main = Section('Upload a log', Plate(upload_body('out'), 'Upload a log'), meta='For a night already logged') + recent_section()
        rail = companion_section(False) + Section('Your reports', f'<p class="empty">{escape(COPY["myReportsCopy"]["signedOut"])}</p>')
    elif state == 'empty':
        main = (Section('Upload a log', Plate(upload_body('in'), 'Upload a log'), meta='For a night already logged')
                + Section('Your reports', f'<p class="empty">{escape(COPY["myReportsCopy"]["empty"])}</p>') + recent_section())
        rail = companion_section(True)
    else:
        main = (Section('Your reports', Plate(Rows([report_row(r) for r in MINE['rows']]), 'Your reports'), action=Link('All your reports &rsaquo;'),
                        meta=f'{len(MINE["rows"])} reports, newest first')
                + Section('Upload a log', upload_body('open'), meta='For a night already logged') + recent_section())
        rail = companion_section(True)
    body = head + Page('split', main, rail) + PAGE_CSS
    who = ACCOUNT if state != 'out' else None
    return build_page('Logs', body, 'Logs', who, f'/logs mock, {state}. Data: logs-reports-stub.json (reports-page.spec.ts rows), the sample meta, lib/reports copy.')


# ------------------------------------------------------------------------------------------------ /reports
def reports_page(state: str) -> str:
    """state: out | empty | list | failed."""
    yc, mc = COPY['yourReportsPageCopy'], COPY['myReportsCopy']
    meta = [escape(yc['job'])] + ([f'{len(MINE["rows"])} reports'] if state == 'list' else [])
    head = PageHeader(escape(yc['title']), meta=meta, actions=Button(mc['uploadAction']) if state in ('list', 'empty') else '')
    if state == 'out':
        inner = f'<div class="empty"><p>{escape(yc["signedOut"])}</p><div class="act">{Button(mc["signInAction"], primary=True)}{Link(yc["sampleAction"] + " &rsaquo;")}</div></div>'
    elif state == 'empty':
        inner = f'<div class="empty"><p>{escape(yc["empty"])} {escape(mc["empty"].split(". ", 1)[1])}</p><div class="act">{Button(mc["uploadAction"], primary=True)}{Link(yc["sampleAction"] + " &rsaquo;")}</div></div>'
    elif state == 'failed':
        inner = f'<div class="empty" role="alert"><p>{escape(mc["failed"])}</p><div class="act">{Button("Try again", small=True)}</div></div>'
    else:
        inner = Rows([report_row(r) for r in MINE['rows']])
    body = head + Page('ledger', Plate(inner, 'Your reports')) + PAGE_CSS
    return build_page('Your reports', body, 'Logs', ACCOUNT if state != 'out' else None, f'/reports mock, {state}. Data: logs-reports-stub.json.')


# ------------------------------------------------------------------------------------------------ rankings
def pct_colour(rank: int, total: int) -> str:
    p = min(99, 100 - (rank - 1) / max(total, 1) * 100)
    for floor, key in ((99, 'pink'), (95, 'orange'), (75, 'purple'), (50, 'blue'), (25, 'green')):
        if p >= floor:
            return PARSE[key]
    return PARSE['grey']


def legend() -> str:
    steps = [('99+', 'pink'), ('95+', 'orange'), ('75+', 'purple'), ('50+', 'blue'), ('25+', 'green'), (COPY['rankingsCopy']['legendUnderLast'], 'grey')]
    return '<p class="legend">' + escape(COPY['rankingsCopy']['legendLead']) + ''.join(f' <b style="color:{PARSE[k]}">{t}</b>' for t, k in steps) + '</p>'


def ranking_row(row: dict, total: int, newest: datetime, multi: bool, you: bool = False) -> str:
    p = row['player']
    sl = p['class'].lower()
    days = (newest - datetime.strptime(row['fought_at'][:10], '%Y-%m-%d')).days + 1
    guild = row.get('guild')
    sub = (f'{escape(guild["name"])} &middot; ' if guild else '') + f'{p["spec"]} {p["class"]} &middot; {row["size"]}-man'
    top = 900 - 12   # the board's first value (the stub's rank-1 row); the bar is value over it
    rank = f'<span style="color:{pct_colour(row["rank"], total)}">{row["rank"]}</span>'
    secs = row['duration_ms'] // 1000
    figs = [(f'{row["value"]:,}', 'DPS', '', False), (f'{row["execution_score"] * 100:.0f}%', '', 'Executed', False), (f'{secs // 60}:{secs % 60:02d}', '', 'Length', False)]
    return RankRow(rank, crest_img(sl, 32), CLASS_COLOR[sl], escape(p['name']), sub, row['value'] / top, figs, f'{days} day{"s" if days != 1 else ""} ago',
                   [('Report', True), ('Build', False)], you='You' if you else '', selected=you)


def board(data: dict, own_key: str = '', off_page: bool = False) -> str:
    rows = data['rows']
    total = data['total']
    newest = max(datetime.strptime(r['fought_at'][:10], '%Y-%m-%d') for r in rows)
    multi = len({r['player']['class'] for r in rows}) > 1
    heads = ColumnHeads([('#', True), ('', False), ('Character', False), ('Share of the top', False), ('DPS', True), ('Executed', True), ('Length', True), ('Date', True), ('', False)])
    body = ''
    if off_page and data.get('character_rank'):
        cr = data['character_rank']
        body += f'<div class="pinned" aria-label="{escape(COPY["rankingsCopy"]["ownBest"])}">' + ranking_row({**cr['row'], 'rank': cr['rank']}, total, newest, multi, you=True) + '</div>'
    body += ''.join(ranking_row(r, total, newest, multi, you=(r['player']['key'] == own_key and not off_page)) for r in rows)
    pages = -(-total // data['per_page'])
    pager = f'<div class="pager"><span class="t-small">Page 1 of {pages}</span><span class="muted">Previous</span><a class="link" href="#">Next</a></div>' if pages > 1 else ''
    return Plate(RankList(heads, body, 3, links3=True) + pager, 'Rankings')


def rankings_bar(encounter: str, cls: str, spec: str, open_m: bool = False) -> str:
    rc = COPY['rankingsCopy']
    groups = [(rc['encounterLabel'], sel(encounter)), (rc['classLabel'], sel(cls)), (rc['specLabel'], sel(spec)), (rc['rulesetLabel'], sel('Normal' if cls != rc['everyClass'] else rc['allRulesets'])),
              (rc['metricLabel'], seg([('Damage', True, None, ''), ('Healing', False, None, '')])),
              (rc['boardLabel'], seg([('Characters', True, None, ''), ('Guilds', False, None, '')])),
              ('Filters', '<span class="sel" role="button" aria-haspopup="dialog">Filters</span>')]
    return control_bar(groups, f'{encounter} &middot; {spec if spec != rc["everySpec"] else cls}', open_m=open_m, label='Ranking choices', labels_from=1920)


def rankings_page(state: str) -> str:
    """state: index | empty | board | you | off | unranked | bar."""
    rc, ec = COPY['rankingsCopy'], COPY['encounterPickerCopy']
    upd = f'Updated {day(UPDATED)}'
    tier = Notice(f'<span>{escape(rc["tierQuestion"])}</span> {Link(rc["tierLink"] + " &rsaquo;")}', tone='neutral')
    if state in ('index', 'empty'):
        head = PageHeader(escape(rc['h1']), meta=[escape(rc['jobSignedOut']), upd])
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
    body = head + rankings_bar(enc, cls, spec, open_m=state == 'bar') + f'<div data-region="summary">{tier}</div>' + Page('ledger', plate + f'<div class="section" data-region="footer-note">{legend()}</div>')
    return build_page(f'{enc} rankings', body + PAGE_CSS, 'Logs', RANK_YOU if you else None, f'/rankings/<slug> mock, {state}. Data: rankings-fixture.ts boardFor().')


# ------------------------------------------------------------------------------------------------ /ratings
def astro_paras(card_title: str) -> list[str]:
    """The <p> texts of one Card in ratings.astro, tags other than <strong> stripped, whitespace collapsed."""
    i = RATINGS_SRC.index(f'title="{card_title}"')
    j = RATINGS_SRC.find('</Card>', i)
    out = []
    for m in re.finditer(r'<p[^>]*>(.*?)</p>', RATINGS_SRC[i:j], re.S):
        t = re.sub(r'\s+', ' ', m.group(1)).strip()
        if '{' in t:
            continue
        out.append(t)
    return out


def astro_testid_para(testid: str) -> str:
    m = re.search(r'data-testid="' + testid + r'">(.*?)</p>', RATINGS_SRC, re.S)
    return re.sub(r'\s+', ' ', m.group(1)).strip()


def ratings_page() -> str:
    sc, rc = COPY['scoreCopy'], COPY['ratingCopy']
    fs = FACTS['fns']['score']
    sample = {k: int(v) for k, v in re.findall(r'(\w+): (\d+)', re.search(r'components: \{([^}]*)\}', RATINGS_SRC).group(1))}
    score = int(re.search(r'score: (\d+),', RATINGS_SRC).group(1))
    weights = re.findall(r"\{ component: '(\w+)', dps: (\d+), healer: (\d+), tank: (\d+) \}", RATINGS_SRC)
    trend = [int(x) for x in re.search(r'const sampleTrend = \[([\d, ]+)\]', RATINGS_SRC).group(1).split(',')]
    maxes = FACTS['score']['component_max']
    total = sum(maxes.values())
    tiers = FACTS['score']['tiers']
    tcol = {'common': '#ffffff', 'uncommon': '#1eff00', 'rare': '#3d94f0', 'epic': '#b866f5', 'legendary': '#ff8000'}
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
            f'<div class="fig"><span class="t-label">Weakest part</span><span class="t-small" style="font-size:15px;color:var(--text)">{escape(rc["weakestPart"] if isinstance(rc["weakestPart"], str) else "Activity")}</span></div></div>')
    rfig = rfig.replace('[fn]', 'Weakest part of a fight: Activity.')
    s2 = Section(sc['explainerRatingHeading'], f'<p class="t-lead">{escape(sc["explainerRatingBridge"])} {escape(sc["explainerPreparationShared"])}</p>{rfig}'
                 f'<div class="prose"><p class="muted">{astro_testid_para("ratings-visibility")}</p>' + ''.join(f'<p>{p}</p>' for p in astro_paras('One number first, then six parts')) + '</div>',
                 meta=escape(sc['explainerRatingCaption']))
    comp = Rows([Row(p.split('</strong>')[0].replace('<strong>', '').rstrip('.'), p.split('</strong>', 1)[1].strip(), link=False) for p in astro_paras('What each component measures')], open_=True)
    s3 = Section('What each component measures', comp)
    wrows = [StatRow('', f'<b>{c}</b>', '', [d, h, t], primary_phone=f'{d} / {h} / {t}', bar=False) for c, d, h, t in weights]
    s4 = Section('Weights, by role', StatTable([('', False), ('Component', False), ('DPS', True), ('Healer', True), ('Tank', True)], wrows, 3, bar=False, open_=True)
                 + f'<div class="prose"><p class="muted">{astro_testid_para("ratings-roles")}</p><p class="muted">{astro_testid_para("ratings-adjustable")}</p></div>')
    s5 = Section('Ranked against your spec, or absolute', '<div class="prose">' + ''.join(f'<p>{p}</p>' for p in astro_paras('Ranked against your spec, or absolute')) + '</div>')
    s6 = Section('The cap', f'<div class="prose"><p>{escape(rc["cappedRule"])}</p></div>')
    trows = [Row(escape(t['tier'].capitalize()), escape(fs['tierShort'][t['tier']]), link=False, title_color=tcol[t['tier']], selected=t['tier'] == tier) for t in tiers]
    rail = (Section(sc['tierColumn'] + 's', Rows(trows, open_=True) + f'<p class="t-small" style="margin-top:10px">{escape(sc["tiersNote"])}</p>')
            + Section(sc['explainerLogsHeading'], f'<p class="t-small" style="margin-top:10px">{escape(sc["explainerLogsBody"])}</p><p style="margin-top:12px">{Button(sc["explainerLogsButton"])}</p>'))
    body = head + Page('split', s1 + s2 + s3 + s4 + s5 + s6, rail) + PAGE_CSS
    return build_page(sc['explainerPageTitle'], body, 'Logs', None, '/ratings mock. Data: ratings.astro sample score, trend, weights and prose; lib/score and lib/rating copy.')


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
