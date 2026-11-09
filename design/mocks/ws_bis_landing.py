"""Web-layout mock: the best-in-slot landing (/bis), composed ONLY from ws_prims primitives.

Spec: design/specs/2026-11-07-bis-landing.md. Every word, class, spec, range caption, count and button label is read from
design/mocks/data/bis-landing-facts.json, which design/mocks/bis_landing_facts.mjs writes by running the landing's own models
(lib/bis/landing-*.ts) on the active build and the account fixtures. Nothing is typed in but the layout.

  node design/mocks/bis_landing_facts.mjs && python3 design/mocks/ws_bis_landing.py
  python3 design/mocks/ws_bis_landing.py [mock ...]              (every mock, or just those named; --list names them)
  node design/mocks/ws_shoot.mjs <mock> 390,1440,1920,2000
FS_BIS_LANDING_FACTS and FS_MOCK_OUT point the generator elsewhere (the parity spec's test-time render).
"""
import json
import os
import sys
from html import escape
from pathlib import Path

from mocklib import CLASS_COLOR
from ws_lib import page as chrome_page
from ws_prims import (PRIM_CSS, ClassSpecPicker, assert_no_surface_styles, Button, Figure, Link, Meter, Page, PageHeader, Plate, Row, RowsSlim, Skeleton, Tag, build_page,
                      control_bar, crest_img, icon, seg, write, REGION)

HERE = Path(__file__).resolve().parent
F = json.loads(Path(os.environ.get('FS_BIS_LANDING_FACTS') or HERE / 'data' / 'bis-landing-facts.json').read_text())
C = F['copy']


def e(text: str) -> str:
    return escape(str(text), quote=False)


def rsaquo(text: str) -> str:
    return f'{e(text)} &rsaquo;'


# ------------------------------------------------------------------------------------------------ the head's answer
def answer(acct: dict | None, loading: bool = False) -> str:
    """ui/CharacterAnswer with its figure: the character, the count Figure over a Meter (or the sentence in its place), the one
    sentence-case Button. loading: the reserve (the session is hinted, /v1/me in flight)."""
    lab = f'<span class="t-label">{e(C["answerLabel"])}</span>'
    if loading:
        return (f'<div class="ans-box bl-ans" aria-busy="true">{lab}<span class="ans-row">{Skeleton(40, 40, True)}<span class="ans-who">'
                f'{Skeleton(150, 22)}{Skeleton(190, 16)}</span></span><span class="bl-fig">{Skeleton(160, 40)}{Skeleton(220, 8)}</span>'
                f'{Skeleton(240, 44)}</div>')
    a = acct['answer']
    if a['state'] == 'none':
        return f'<div class="ans-box bl-ans">{lab}<span class="t-small">{e(a["message"])}</span><span class="ans-links">{Link(rsaquo(a["link"]["label"]))}</span></div>'
    if a['state'] == 'error':
        return f'<div class="ans-box bl-ans">{lab}<span class="t-small" role="alert">{e(a["message"])}</span>{Button(e(a["retry"]), act=True, tone="ghost")}</div>'
    ch = a['character']
    color = CLASS_COLOR.get(ch['classSlug'], '#f2eee4')
    who = (f'<span class="ans-row">{crest_img(ch["classSlug"], 40)}<span class="ans-who"><b class="ans-name" style="color:{color}">{e(ch["name"])}</b>'
           f'<span class="t-small">{e(ch["line"])}</span></span></span>')
    body = a['body']
    if body['kind'] == 'count':
        fig = Figure(str(body['worn']), unit=e(body['unit'])).replace('</span></div>', f'</span><span class="fcap">{e(body["caption"])}</span></div>', 1)
        note = f'<span class="t-small">{e(body["note"])}</span>' if body.get('note') else ''
        mid = f'<span class="bl-fig">{fig}{Meter(body["bar"])}{note}</span>'
    elif body['kind'] == 'failed':
        mid = f'<span class="t-small" role="alert">{e(body["text"])}</span>{Button(e(body["retry"]), act=True, tone="ghost")}'
    else:
        mid = f'<span class="t-small">{e(body["text"])}</span>'
    link = Link(rsaquo(body['link']['label'])) if body['kind'] == 'sentence' and body.get('link') else ''
    return f'<div class="ans-box bl-ans">{lab}{who}{mid}<span class="ans-links">{link}{Button(e(a["action"]["label"]), primary=True)}</span></div>'


# ------------------------------------------------------------------------------------------------ the page
def band_bar(selected: int) -> str:
    items = [(e(b['label']), b['band'] == selected, None, '') for b in F['bands']]
    label = next(b['label'] for b in F['bands'] if b['band'] == selected)
    end = f'<span class="t-small">{e(C["tierQuestion"])}</span>{Link(rsaquo(C["tierLink"]))}'
    return control_bar([(e(C['bandGroupLabel']), seg(items, label=e(C['bandGroupLabel'])))], e(label), end, label=e(C['bandGroupLabel']))


def summary(selected: int) -> str:
    text = next(b['summary'] for b in F['bands'] if b['band'] == selected)
    return f'<p class="t-small bl-sum" data-region="{REGION["summary"]}">{e(text)}</p>'


X_GLYPH = ('<svg width="16" height="16" viewBox="0 0 16 16" fill="none" stroke="currentColor" stroke-width="1.75" stroke-linecap="round">'
           '<path d="M3.5 3.5l9 9M12.5 3.5l-9 9"/></svg>')


def finished_line() -> str:
    """The finished server sim (components/character/FinishedSimLine.svelte, as shipped): the sentence and Open result as one
    link on a line with a 2 px gold edge (no box), a dismiss."""
    f = F['finished']
    return (f'<div class="pc"><div class="bl-done" role="status"><a href="#"><span>{e(f["prefix"])}<span class="num strong">{e(f["figure"])}</span></span>'
            f'<span class="gold link">{e(f["open"])}</span></a><button type="button" class="bl-x" aria-label="{e(f["dismiss"])}">'
            f'<span aria-hidden="true">{X_GLYPH}</span></button></div></div>')


def classes(yours: str = '') -> str:
    """The class and spec picker /guides shares (ws_prims.ClassSpecPicker)."""
    picks = [(c['slug'], e(c['name']), c['color'], [(s['key'], e(s['name']), icon(s['icon'], 28)) for s in c['specs']]) for c in F['classes']]
    return ClassSpecPicker(picks, e(C['plateLabel']), yours=yours, yours_label=e(C['yours']))


def landing(state: str) -> str:
    """state: out | out-<band> | 24 | 34 | 60 | new | nogear | sim | loading | error."""
    signed_out = state.startswith('out')
    acct = None if signed_out or state == 'loading' else F['accounts'][state]
    selected = int(state.split('-')[1]) if state.startswith('out-') else (acct or {}).get('band') or 20
    ans = '' if signed_out else answer(acct, loading=state == 'loading')
    meta = [f'Updated {e(F["updated"])}'] if F['updated'] else None
    head = PageHeader(e(C['title']), '', '', meta, ans)
    yours = ''
    if acct and acct['answer']['state'] == 'ready':
        me = F['me'].get(state)
        ch = me['characters'][0] if me else None
        yours = f"{ch['class'].lower()}-{ch['spec'].lower().replace(' ', '-')}" if ch else ''
    done = finished_line() if state == 'sim' else ''
    body = head + done + band_bar(selected) + Page('ledger', summary(selected) + f'<div id="bis-classes">{classes(yours)}</div>') + PAGE_CSS
    who = None
    if acct and acct['answer']['state'] == 'ready':
        ch = acct['answer']['character']
        who = (ch['name'], ch['classSlug'], ch['line'].split(' · ')[0])
    note = f'BiS landing mock (design/specs/2026-11-07-bis-landing.md); data bis-landing-facts.json, build {F["build"]}.'
    if state == 'loading':   # the session is hinted and /v1/me is in flight: the header's pending chip
        assert_no_surface_styles(body)
        return chrome_page('Best in slot', body + f'<style>{PRIM_CSS}</style>', 'BiS', None, note, chrome_state='pending')
    return build_page('Best in slot', body, 'BiS', who, note)


PAGE_CSS = '''<style>
/* page-level layout only (no surfaces): the answer's inner rows, the summary line, the finished-sim line */
.ans-row{display:flex;align-items:center;gap:12px}.ans-who{display:flex;flex-direction:column;min-width:0}.ans-name{font:700 18px/22px var(--font-body)}
.ans-links{display:flex;gap:4px 16px;flex-wrap:wrap;padding-top:4px}
.bl-fig{display:flex;flex-direction:column;gap:6px;min-width:0;width:100%;max-width:380px}.fcap{font-size:12px;line-height:16px;color:var(--muted)}
@media(min-width:900px){.ans-row{flex-direction:row-reverse}.ans-who{align-items:flex-end}.ans-links{justify-content:flex-end}.bl-fig{align-items:flex-end}
.bl-fig .meter{width:100%}.bl-fig .fig{align-items:flex-end}}
.ans-box.bl-ans{min-height:224px}
.bl-sum{padding:0 0 16px}
.bl-done{display:flex;align-items:center;justify-content:space-between;gap:12px;min-height:44px;padding:0 0 0 22px;margin-bottom:12px;font-size:13px}
.bl-done{box-shadow:inset 2px 0 0 var(--gold)}
.bl-done a{display:inline-flex;flex-wrap:wrap;gap:0 6px;align-items:center;min-height:44px;color:var(--text)}
.bl-x{display:inline-flex;align-items:center;justify-content:center;width:44px;height:44px;border:0;background:none;color:var(--muted);cursor:pointer}
main{padding-bottom:40px}  /* the page's own bottom space (index.astro .bis-landing) */
@media(max-width:767px){.cbar .pc{padding-block:0;min-height:0}}   /* the phone bar is its 48 px line, as the built ControlBar */
</style>'''

MOCKS = {
    'web-bis-landing': lambda: landing('out'),
    **{f'web-bis-landing-out-{b}': (lambda b=b: landing(f'out-{b}')) for b in (30, 40, 50, 60)},
    'web-bis-landing-24': lambda: landing('24'),
    'web-bis-landing-34': lambda: landing('34'),
    'web-bis-landing-60': lambda: landing('60'),
    'web-bis-landing-new': lambda: landing('new'),
    'web-bis-landing-nogear': lambda: landing('nogear'),
    'web-bis-landing-sim': lambda: landing('sim'),
    'web-bis-landing-loading': lambda: landing('loading'),
    'web-bis-landing-error': lambda: landing('error'),
}


# A server-sim account never sees the header's Premium door (Header.astro: data-premium-nav="hide").
SERVER_SIMS = {'web-bis-landing-sim'}


def main(argv: list[str]) -> None:
    if argv[:1] == ['--list']:
        print('\n'.join(MOCKS))
        return
    for name in argv or list(MOCKS):
        path = write(f'{name}.html', MOCKS[name]())
        if name in SERVER_SIMS:
            path.write_text(path.read_text().replace('<html lang="en"', '<html lang="en" data-premium-nav="hide"', 1))
        print(path)


if __name__ == '__main__':
    main(sys.argv[1:])
