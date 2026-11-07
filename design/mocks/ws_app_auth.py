"""Web-layout mock: /app/auth, the page a phone-app sign-in lands on when the app did not catch the link (design/specs/
2026-11-07-bis-landing.md section 7), composed ONLY from ws_prims primitives: the message page of the flat system, as the
build-unavailable page (ws_premium.py b_page). The words are read from web/src/lib/app-links/handoff-copy.ts.

  python3 design/mocks/ws_app_auth.py          (FS_MOCK_OUT points it elsewhere: the parity spec's test-time render)
States: web-app-auth (no code in the address bar: the start-again line), web-app-auth-code (a code: the Open button).
"""
import re
import sys
from html import escape
from pathlib import Path

from mocklib import ROOT
from ws_prims import Button, PageForm, PageHeader, build_page, write

COPY = dict(re.findall(r"^\s+(\w+): ['\"](.+?)['\"],$", (Path(ROOT) / 'web/src/lib/app-links/handoff-copy.ts').read_text(), re.M))


def e(text: str) -> str:
    return escape(text, quote=False)


def page(code: bool) -> str:
    first = Button(e(COPY['button']), primary=True, act=True) if code else f'<p class="t-small">{e(COPY["missing"])}</p>'
    main = f'<div class="aa">{first}<p class="t-small">{e(COPY["notYou"])}</p></div>'
    css = '<style>.aa{display:flex;flex-direction:column;align-items:flex-start;gap:12px}</style>'
    body = PageHeader(e(COPY['heading']), meta=[e(COPY['body'])]) + PageForm(main) + css
    return build_page(e(COPY['title']), body, '', None, '/app/auth mock (design/specs/2026-11-07-bis-landing.md section 7).')


MOCKS = {'web-app-auth': lambda: page(False), 'web-app-auth-code': lambda: page(True)}

if __name__ == '__main__':
    for name in sys.argv[1:] or list(MOCKS):
        print(write(f'{name}.html', MOCKS[name]()))
