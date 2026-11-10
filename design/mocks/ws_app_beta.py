"""Web-layout mocks for the app beta (2026-11-08), composed ONLY from ws_prims:
  - web-app: foreversixty.gg/app, the permanent link the apps' QR codes open. The PageHeader (the apps, the Coming soon
    Tag, the lead), then a split Page: What the apps do (four lines) and Test the apps (the signup form) in main, On your
    phone (the QR for this page, from 768 only) in the rail;
  - web-app-done: the form replaced by its one status line (a new address and a repeat read the same: no enumeration);
  - web-app-invalid: the address refused, its line under the field; web-app-failed: the try-again line under the button;
  - web-admin-testers: the owner's /admin/testers (PageHeader, a ledger Page: the copy buttons over the Grid of sign-ups).
Data and copy are read at render time from design/mocks/data/app-beta-facts.json (node design/mocks/app_beta_facts.mjs),
the QR included (lib/app-beta/qr.ts draws it for both), never retyped:

  node design/mocks/app_beta_facts.mjs && python3 design/mocks/ws_app_beta.py
FS_APP_BETA_FACTS and FS_MOCK_OUT point the generator elsewhere (the parity spec's test-time render).
"""
import json
import os
import sys
from html import escape
from pathlib import Path

import ws_prims
from ws_prims import Button as _Button, Grid, Page, PageHeader, Section, Segmented, Tag, assert_no_surface_styles, write
from ws_lib import page as _page

HERE = Path(__file__).resolve().parent
F = json.loads(Path(os.environ.get('FS_APP_BETA_FACTS') or HERE / 'data' / 'app-beta-facts.json').read_text())
A, D = F['app'], F['admin']
C = A['copy']
OWNER = ('Thandrel', 'warrior', 'Level 60')

CSS = '''<style>
/* page-level layout only (no surfaces) */
section>.sec-h+*{margin-top:12px}
.feats{margin:0;padding-left:20px;display:flex;flex-direction:column;gap:8px;max-width:62ch;font-size:15px;line-height:24px;list-style:disc}
.feats li::marker{color:var(--gold)}
.beta{display:flex;flex-direction:column;gap:12px;max-width:440px}.beta>p{font-size:15px;line-height:24px;margin:0}
.beta .job{color:var(--muted)}.beta .done{color:var(--strong)}
.beta form{display:flex;flex-direction:column;gap:16px}
.f{display:flex;flex-direction:column;gap:6px;min-width:0}.f .err{color:var(--wipe);font-size:13px;line-height:20px}
input.field.bad{border-color:var(--wipe)}
.f .seg{align-self:flex-start}
/* the built controls (lib/ui/field.ts FIELD_HEIGHT h-11 md:h-10, ui/segmented.css 44 px items) */
.beta .seg a{height:44px}@media(max-width:767px){.beta input.field{height:44px}}
.send{display:flex;flex-direction:column;align-items:flex-start;gap:8px}.send p{margin:0}.send .alert{color:var(--wipe)}
.qr{display:flex;flex-direction:column;align-items:flex-start;gap:12px}.qr p{margin:0}.qr .code{width:200px;line-height:0}
.qr .code svg{width:100%;height:auto}
@media(max-width:767px){.qr-rail{display:none}}
.acts{display:flex;gap:12px;flex-wrap:wrap;align-items:center;padding-top:12px}
.stack{display:flex;flex-direction:column;gap:16px}.page-foot{padding-bottom:48px}
/* the built sm Button is a 44 px target below 1024 (web/src/lib/ui/button.ts BUTTON_SM) */@media(max-width:1023px){.btn.sm{height:44px}}
</style>'''


def e(t) -> str:
    return escape(str(t), quote=False)


def a(t) -> str:
    return escape(str(t), quote=True)


def Btn(label: str, primary: bool = False, small: bool = False, submit: bool = False) -> str:
    """The system Button as the build renders an action (a <button>)."""
    kind = 'submit' if submit else 'button'
    return _Button(label, primary, small).replace('<a ', f'<button type="{kind}" ', 1).replace(' href="#"', '', 1).replace('</a>', '</button>')


def render(title: str, body: str, note: str, who: tuple | None = None) -> str:
    assert_no_surface_styles(body)
    return _page(title, body + CSS + f'<style>{ws_prims.PRIM_CSS}</style>', '', who, note, '')


# ================================================================================================ /app
def email_field(error: str = '') -> str:
    bad = ' bad' if error else ''
    invalid = ' aria-invalid="true" aria-describedby="app-beta-email-error"' if error else ''
    line = f'<span class="err" id="app-beta-email-error">{e(error)}</span>' if error else ''
    return (f'<label class="f" for="app-beta-email"><span class="t-label">{e(C["emailLabel"])}</span>'
            f'<input class="field{bad}" id="app-beta-email" type="email" placeholder="{a(C["emailPlaceholder"])}"{invalid}>{line}</label>')


def signup(state: str) -> str:
    """Test the apps: the job line, then the form, or the one status line once sent."""
    job = f'<p class="job">{e(C["formJob"])}</p>'
    if state == 'done':
        return f'<div class="beta">{job}<p class="done" role="status">{e(A[state])}</p></div>'
    platform = (f'<div class="f"><span class="t-label" aria-hidden="true">{e(C["platformLabel"])}</span>'
                + Segmented([(e(A['platforms']['ios']), True), (e(A['platforms']['android']), False)]).replace(
                    'role="radiogroup"', f'role="radiogroup" aria-label="{a(C["platformLabel"])}"', 1) + '</div>')
    alert = f'<p class="t-small alert" role="alert">{e(C["failed"])}</p>' if state == 'failed' else ''
    send = f'<div class="send">{Btn(e(C["submit"]), primary=True, submit=True)}<p class="t-small">{e(C["consent"])}</p>{alert}</div>'
    form = f'<form novalidate>{email_field(C["invalidEmail"] if state == "invalid" else "")}{platform}{send}</form>'
    return f'<div class="beta">{job}{form}</div>'


def app_page(state: str) -> str:
    feats = '<ul class="feats">' + ''.join(f'<li>{e(line)}</li>' for line in C['features']) + '</ul>'
    main = Section(e(C['featuresTitle']), feats, panel=True, frame='framed') + Section(e(C['formTitle']), signup(state), panel=True)
    qr = (f'<div class="qr"><div class="code" role="img" aria-label="{a(C["qrAlt"])}">{A["qr"]}</div>'
          f'<p class="t-small">{e(C["qrLine"])}</p></div>')
    rail = f'<div class="qr-rail">{Section(e(C["qrTitle"]), qr, panel=True)}</div>'
    head = PageHeader(e(C['heading']), meta=[Tag(e(C['soon']))])
    body = head + Page("split", main, rail, panels=True)
    return render(f'{e(C["pageTitle"])} · Forever Sixty', body, f'/app mock, {state}. Data: app-beta-facts.json (lib/app-beta/copy.ts, qr.ts).')


# ================================================================================================ /admin/testers
def admin_testers() -> str:
    c = D['copy']
    buttons = '<div class="acts">' + ''.join(Btn(e(label), small=True) for label in D['buttons']) + '</div>'
    rows = [(e(r['label']), '', [('none', e(cell['text'])) for cell in r['cells']]) for r in D['rows']]
    grid = Grid([e(col) for col in c['columns']], rows, caption=e(c['caption']), label_header=e(c['emailHeader']))
    main = Section(e(c['tableTitle']), f'<div class="stack">{buttons}{grid}</div>', meta=e(D['meta']))
    body = PageHeader(e(c['pageTitle']), lead=e(c['job'])) + f'<div class="page-foot">{Page("ledger", main)}</div>'
    return render(f'{e(c["pageTitle"])} · Forever Sixty', body,
                  '/admin/testers mock (the owner); data app-beta-facts.json (admin.ts over fixtures/app-beta/testers.json).', OWNER)


APP_STATES = ['idle', 'done', 'invalid', 'failed']
PAGES = {('web-app' if s == 'idle' else f'web-app-{s}'): (lambda s=s: app_page(s)) for s in APP_STATES}
PAGES['web-admin-testers'] = admin_testers

if __name__ == '__main__':
    if sys.argv[1:2] == ['--list']:
        print(' '.join(PAGES))
        sys.exit(0)
    for name in sys.argv[1:] or list(PAGES):
        print(write(f'{name}.html', PAGES[name]()))
