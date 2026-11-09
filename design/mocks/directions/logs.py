"""/logs and /reports in the directions system (plan 2026-11-09-redesign, batch 1).

  python3 design/mocks/directions/logs.py   -> html/b1-logs.html, html/b1-reports.html

Data: web/src/fixtures/reports (production's one public report, the 17 Sept beta night, as the signed-in player's list and
the public feed) and the sample's meta, through design/mocks/ws_logs.py. The production row carries no zone or title; it
is the same log as the sample (its first 42 pulls), so the row takes the sample's zone, Zephras Isle.
One row component for both pages (ux-designer, batch 1): report_row().
"""
import sys
from datetime import datetime
from html import escape
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

import ws_logs as W  # noqa: E402
from common import write  # noqa: E402
from prims import PRIMS_CSS, panel, row_list, tag  # noqa: E402
from style import footer, header, page, titlebar  # noqa: E402

VIEWER = ('Thandrel', 'warrior', 60)
FIRST_PULL = {r['id']: r['first_pull'] for r in W.MINE['rows']}  # the public feed's row has no started_at; the same report's own does
SAME_LOG_ZONE = {'z2opt2x2gs5u': W.SAMPLE_META['zone']}  # production's report is the sample's log: same zone
MONTHS = ('Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sept', 'Oct', 'Nov', 'Dec')
COMPANION_URL = '#'

LOGS_CSS = PRIMS_CSS + '''
.rp{padding-top:6px}
.drop{border:1px dashed rgba(243,238,226,.4);border-radius:var(--r-sm);padding:26px 20px;text-align:center;color:var(--text)}
.drop b{display:block;color:#fff;font-size:16px;margin-bottom:4px}.drop span{font-size:13px;color:var(--muted)}
.drop code{font:13px ui-monospace,Menlo,monospace;color:#fff}
.rl .zn{display:inline-flex;gap:8px;align-items:center}
'''


def full_date(iso_day: str) -> str:
    d = datetime.strptime(iso_day[:10], '%Y-%m-%d')
    return f'{d.day} {MONTHS[d.month - 1]} {d.year}'


def uploaded(iso_day: str) -> str:
    """Relative inside two weeks, then the date with its year (wow-player, batch 1: "20 Oct" had no year)."""
    days = (W.NOW - datetime.strptime(iso_day[:10], '%Y-%m-%d')).days
    if days < 1:
        return 'today'
    return f'{days} day{"s" if days != 1 else ""} ago' if days < 14 else full_date(iso_day)


def report_row(title: str, zone: str, pulls: int, bosses: int, kills: int, when: str, sample: bool = False) -> tuple[str, str, str, str]:
    """The one report row: title as a link, zone and what it holds, the upload date at the right."""
    held = f'{kills} of {bosses} bosses killed' if bosses else 'Trash only'
    sub = ' &middot; '.join(x for x in (escape(zone), f'{pulls} pulls', held) if x)
    trail = tag('Sample') if sample else f'<span class="muted" style="font-size:13px">Uploaded {when}</span>'
    return ('', f'<a href="#">{escape(title)}</a>', sub, trail)


def row_for(r: dict) -> tuple[str, str, str, str]:
    zone = r['zone'] or SAME_LOG_ZONE.get(r['id'], '')
    first_pull = FIRST_PULL.get(r['id'], r['first_pull'])
    title = r['title'] or f'{zone or "Report"}, {full_date(first_pull)}'
    return report_row(title, '' if title.startswith(zone) else zone, r['fight_count'], r['boss_count'], r['kill_count'], uploaded(r['created_at']))


def sample_row() -> tuple[str, str, str, str]:
    sm = W.SAMPLE_META
    return report_row(sm['title'], sm['zone'], len(sm['fights']), 0, 0, '', sample=True)


def logs_page() -> str:
    lc = W.COPY['logsCompanionCopy']
    upload = ('<div class="drop"><b>Drop <code>WoWCombatLog.txt</code> here</b><span>From the game&rsquo;s Logs folder, up to 4 GB. '
              '<a href="#">Choose a file</a></span></div>')
    mine = [row_for(r) for r in W.MINE['rows']]
    # the public feed leaves out the viewer's own reports, which sit just above (reviewers, round 2: listed twice)
    recent = [row_for(r) for r in W.MINE['recent'] if r['id'] not in FIRST_PULL] + [sample_row()]
    main = (panel('Upload a log', upload, framed=True, ornate=True)
            + panel('Your reports', row_list(mine), action='<a href="#">All &rsaquo;</a>')
            + panel(W.COPY['recentReportsCopy']['heading'], row_list(recent)))
    # LogsCompanion.svelte signed in: the pointer and its button, then the pairing block (PairingBlock.svelte)
    rail = panel('The companion', f'<p>{escape(lc["pointer"])}</p><p style="margin-top:12px"><a class="btn" href="{COMPANION_URL}">'
                 f'{escape(lc["pointerLink"])}</a></p><p style="margin-top:18px;color:#fff;font-weight:600">{escape(lc["pairLabel"])}</p>'
                 '<p style="margin-top:8px"><a class="btn2" href="#">Get a pairing code</a></p>')
    body = (titlebar(None, W.COPY['logsBandCopy']['h1'], '', top=header(*VIEWER, 'Logs'))
            + f'<main class="w rp"><div class="twocol"><div>{main}</div><aside>{rail}</aside></div></main>' + footer())
    return page('Forever Sixty, logs', LOGS_CSS, body)


def reports_page() -> str:
    yc = W.COPY['yourReportsPageCopy']
    rows = [row_for(r) for r in W.MINE['rows']]
    body = (titlebar(None, yc['title'], '', right='<a class="btn" href="#">Upload a log</a>', top=header(*VIEWER, 'Logs'))
            + f'<main class="w rp">{panel("", row_list(rows))}</main>' + footer())
    return page('Forever Sixty, your reports', LOGS_CSS, body)


def build() -> None:
    write('b1-logs', logs_page())
    write('b1-reports', reports_page())


if __name__ == '__main__':
    build()
