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

from ws_prims import Button, Figure, LaneBlock, Page, PageHeader, Prose, ReadingPanel, ReadingSection, Toc, build_page, write, REGION

HERE = Path(__file__).resolve().parent
F = json.loads(Path(os.environ.get('FS_CREATORS_FACTS') or HERE / 'data' / 'creators-page-facts.json').read_text())
PR = F['program']

CSS = '''<style>
/* page-level layout only (no surfaces): the lead's inset on the panel, Apply under its sentence */
.cr-lead{padding:22px 24px 0}.cr-apply{padding-top:14px}.cr-ans{display:flex;flex-wrap:wrap;align-items:center;gap:10px 20px}.cr-lane{display:flex;flex-direction:column;gap:20px}.cr-cap{display:block;font-size:12px;line-height:16px;color:var(--muted)}
main{padding-bottom:40px}@media(min-width:1024px){main{padding-bottom:64px}}
@media(max-width:767px){.cr-lead{padding:18px 16px 0}.cr-apply .btn{display:flex;width:100%}}
</style>'''


def e(text: str) -> str:
    return escape(str(text), quote=False)


def paragraphs(lines: list[str]) -> str:
    return Prose(''.join(f'<p>{e(line)}</p>' for line in lines))


def facts(rows: list[dict]) -> str:
    """web ui/Facts: each figure's label and its value with the unit."""
    return '<dl class="kv">' + ''.join(f'<dt>{e(f["label"])}</dt><dd>{e(f["value"])}</dd>' for f in rows) + '</dl>'


def lane(section_id: str, first: bool) -> str:
    ln = next((x for x in PR['lanes'] if x['section'] == section_id), None)
    if ln is None:
        return ''
    html = LaneBlock(e(ln['label']), facts(ln['figures']))
    if first:
        html += LaneBlock(e(PR['examplesLabel']), facts(F['examples']), e(PR['examplesNote']))
    return f'<div class="cr-lane">{html}</div>'


def creators_page() -> str:
    a = PR['answer']
    # The built Figure's caption: its one 12/16 muted line under the value.
    answer = Figure(e(a['value']), unit=e(a['unit']), tone='gold')[:-len('</div>')] + f'<span class="cr-cap">{e(a["line"])}</span></div>'
    head = PageHeader(e(PR['pageTitle']), answer=f'<div class="cr-ans">{answer}{Button(e(PR["apply"]), primary=True)}</div>')
    toc = Toc(e(PR['contentsLabel']), [(e(s['title']), s['id']) for s in PR['sections']], PR['sections'][0]['id'])
    lead = f'<div class="cr-lead" data-region="{REGION["intro"]}">{paragraphs([PR["lead"]])}</div>'
    body = (lead + ''.join(ReadingSection(e(s['title']), paragraphs(s['paragraphs']), lane(s['id'], i == 0), s['id'])
                           for i, s in enumerate(PR['sections']))
            + ReadingSection(e(PR['applySection']['title']),
                             paragraphs(PR['applySection']['paragraphs']) + f'<div class="cr-apply">{Button(e(PR["applyAction"]), primary=True)}</div>',
                             '', PR['applySection']['id']))
    html = head + Page('reading', ReadingPanel(body), toc, toc_from=1024, toc_label=e(PR['contentsLabel'])) + CSS
    note = '/creators mock: creators-page-facts.json (lib/creators/copy.ts, model.ts). Apply is the support email until the server side lands.'
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
