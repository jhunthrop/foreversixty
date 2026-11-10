"""Web-layout mock: the public /creators page (directions mock design/mocks/directions/png/b4-creators-*), composed ONLY from
ws_prims primitives, the parity twin of web/src/pages/creators/index.astro.

Every word and figure is read from design/mocks/data/creators-page-facts.json, which design/mocks/creators_facts.mjs writes from
web/src/lib/creators/copy.ts (the program's terms, owner-approved on the creator-program branch, and the apply path). Nothing is
typed in but the layout. The program's server side is not on main, so Apply is the support email.

  node design/mocks/creators_facts.mjs && python3 design/mocks/ws_creators.py
  node design/mocks/ws_shoot.mjs web-creators 390,1440,2000
FS_CREATORS_FACTS and FS_MOCK_OUT point the generator elsewhere (a parity spec's test-time render).
"""
import json
import os
import sys
from html import escape
from pathlib import Path

from ws_prims import Button, LaneBlock, Page, PageHeader, Prose, ReadingPanel, ReadingSection, Toc, build_page, write, REGION

HERE = Path(__file__).resolve().parent
F = json.loads(Path(os.environ.get('FS_CREATORS_FACTS') or HERE / 'data' / 'creators-page-facts.json').read_text())
PR = F['program']

CSS = '''<style>
/* page-level layout only (no surfaces): the lead's inset on the panel, Apply under its sentence */
.cr-lead{padding:22px 24px 0}.cr-apply{padding-top:14px}
main{padding-bottom:40px}@media(min-width:1024px){main{padding-bottom:64px}}
@media(max-width:767px){.cr-lead{padding:18px 16px 0}.cr-apply .btn{display:flex;width:100%}}
</style>'''


def e(text: str) -> str:
    return escape(str(text), quote=False)


def paragraphs(lines: list[str]) -> str:
    return Prose(''.join(f'<p>{e(line)}</p>' for line in lines))


def terms() -> str:
    """The terms at a glance (web ui/Facts in a LaneBlock): each figure with its unit."""
    rows = ''.join(f'<dt>{e(f["label"])}</dt><dd>{e(f["value"])}</dd>' for f in PR['figures'])
    return LaneBlock(e(PR['glanceLabel']), f'<dl class="kv">{rows}</dl>')


def creators_page() -> str:
    apply_btn = Button(e(PR['apply']), primary=True)
    head = PageHeader(e(PR['pageTitle']), actions=apply_btn)
    first, *rest = PR['sections']
    sections = [*PR['sections'], PR['applySection']]
    toc = Toc(e(PR['contentsLabel']), [(e(s['title']), s['id']) for s in sections], sections[0]['id'])
    lead = f'<div class="cr-lead" data-region="{REGION["intro"]}">{paragraphs([PR["lead"]])}</div>'
    body = (lead + ReadingSection(e(first['title']), paragraphs(first['paragraphs']), terms(), first['id'])
            + ''.join(ReadingSection(e(s['title']), paragraphs(s['paragraphs']), '', s['id']) for s in rest)
            + ReadingSection(e(PR['applySection']['title']),
                             paragraphs(PR['applySection']['paragraphs']) + f'<div class="cr-apply">{Button(e(PR["applyAction"]), primary=True)}</div>',
                             '', PR['applySection']['id']))
    html = head + Page('reading', ReadingPanel(body), toc, toc_from=1024, toc_label=e(PR['contentsLabel'])) + CSS
    note = '/creators mock: creators-page-facts.json (lib/creators/copy.ts). Apply is the support email until the program\'s server side lands.'
    return build_page(PR['pageTitle'], html, '', None, note)


PAGES = {'web-creators': creators_page}

if __name__ == '__main__':
    if sys.argv[1:2] == ['--list']:
        print(' '.join(PAGES))
    else:
        only = sys.argv[1:]
        bad = [n for n in only if n not in PAGES]
        if bad:
            sys.exit(f'ws_creators.py: no such mock: {" ".join(bad)} (see --list)')
        for name, fn in PAGES.items():
            if not only or name in only:
                print(write(f'{name}.html', fn()))
