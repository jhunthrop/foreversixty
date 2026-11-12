"""Web-layout reference mocks for the creator program (design/specs/2026-11-08-creator-program.md), composed ONLY from ws_prims:
  - web-creators: the public /creators page (directions mock design/mocks/directions/png/b4-creators-*), the parity twin of
    web/src/pages/creators/index.astro: the PageHeader with the bounty as its answer and Apply, the contents rail, one panel of
    ReadingSections with their figures in the lane, Apply last;
  - web-creators-apply (-pending, -approved, -declined): /creators/apply (Page form) empty and after sending;
  - web-admin-creators (-edit, -approve): the owner's /admin/creators (a split Page: the Requests and creators' Rows, the rail form);
  - web-admin-creators-payouts (-confirm, -confirm-short, -payouts-off): /admin/creators/payouts (the month Select, the statements, the Lines rail);
  - web-creators-invite / web-creators-done: /creators/onboarding (Page form) accepting an invite, and back from Stripe;
  - web-creators-stats: the creator's /creators/stats.
Every word and figure is read at render time from design/mocks/data/creators-facts.json, which design/mocks/creators_facts.mjs
writes from web/src/lib/creators/{copy,model}.ts over the e2e fixtures. Nothing is typed in but the layout.

  node design/mocks/creators_facts.mjs && python3 design/mocks/ws_creators.py
  node design/mocks/ws_shoot.mjs web-creators 390,1440,2000
FS_CREATORS_FACTS and FS_MOCK_OUT point the generator elsewhere (a parity spec's test-time render).
"""
import json
import os
import sys
from html import escape
from pathlib import Path

import ws_prims
from ws_prims import (REGION, Button as _Button, Dialog, Figure, LaneBlock, Link, Page, PageForm, PageHeader, Plate, Prose, ReadingPanel,
                      ReadingSection, Row, Rows, RowsSlim, Section, Select, Tag, Toc, assert_no_surface_styles, build_page, write)
from ws_lib import page as _page

HERE = Path(__file__).resolve().parent
F = json.loads(Path(os.environ.get('FS_CREATORS_FACTS') or HERE / 'data' / 'creators-facts.json').read_text())
PG = F['page']
PR = PG['program']
A, P, O, S = F['admin'], F['payouts'], F['onboarding'], F['stats']
R, G, Y = F['requests'], F['program'], F['apply']
OWNER = ('Thandrel', 'warrior', 'Level 60')
# The creator's pages are drawn signed in as the e2e owner fixture (me-60.json), so the chrome matches the build's.
CREATOR = OWNER

CSS = '''<style>
/* page-level layout only (no surfaces) */
.stack{display:flex;flex-direction:column;gap:12px;min-width:0}.acts{display:flex;gap:10px 12px;flex-wrap:wrap;align-items:center}
.figs5{display:flex;flex-wrap:wrap;gap:8px 20px}.figs5>span{width:104px;display:flex;flex-direction:column}.figs5>span.lead{width:220px}.btn[disabled]{opacity:.45;cursor:not-allowed}
.f{display:flex;flex-direction:column;gap:6px;min-width:0}.f .t-small{font-size:12px;line-height:16px}
.ro{font-size:15px;color:var(--strong)}.status{min-height:22px}.page-foot{padding-bottom:48px}
.trail-end{display:flex;gap:10px;align-items:center;flex-wrap:wrap}.amt{font-variant-numeric:tabular-nums}
section:not(.rsec)>.sec-h+*{margin-top:12px!important}fieldset.fs{border:0;padding:0;margin:0}.radio{width:18px;height:18px;accent-color:var(--gold)}.rprose p:last-child{margin-bottom:0}
</style>'''


def e(t) -> str:
    return escape(str(t), quote=False)


def Btn(label: str, primary: bool = False, small: bool = True, disabled: bool = False) -> str:
    """The system Button as the build renders an action (a <button>)."""
    b = _Button(label, primary, small).replace('<a ', '<button type="button" ', 1).replace(' href="#"', '', 1).replace('</a>', '</button>')
    return b.replace('<button ', '<button disabled ', 1) if disabled else b


# The built Tag's tones as the mock library names them (ui/Tag.svelte success is kill, warn has no mock hue: gold).
TONE = {"success": "kill", "warn": "gold", "muted": "muted", "danger": "wipe", "gold": "gold"}


def Input(placeholder: str) -> str:
    return f'<input class="field" type="text" placeholder="{placeholder}">'


def figures(items: list[dict]) -> str:
    return '<span class="figs5">' + ''.join(
        f'<span{" class=\"lead\"" if f["lead"] else ""}>{Figure(e(f["value"]), e(f["label"]), size="sm", tone="gold" if f["lead"] else "")}<span class="t-small">{e(f["note"])}</span></span>' for f in items) + '</span>'


def render(title: str, body: str, who: tuple | None, note: str) -> str:
    assert_no_surface_styles(body)
    return _page(title, body + CSS + f'<style>{ws_prims.PRIM_CSS}</style>', '', who, note, '')


# ================================================================================================ /admin/creators
def creators_plate() -> str:
    c = A['copy']
    rows = [Row(e(r['name']), e(r['sub']), link=False, foot=figures(r['figures']),
                trail=f'<span class="trail-end">{Tag(e(r["status"]), TONE[r["tone"]])}{Btn(e(r["edit"]))}</span>') for r in A['rows']]
    return Section(e(c['tableTitle']), Plate(Rows(rows), e(c['tableTitle'])) + f'<p style="margin-top:12px">{Link(e(c["statementsLink"]) + " &rsaquo;")}</p>',
                   meta=e(A['meta']))


def field(f: dict, readonly: str = '') -> str:
    control = f'<span class="ro">{e(readonly)}</span>' if readonly else Input(e(f['placeholder']))
    return f'<label class="f"><span class="t-label">{e(f["label"])}</span>{control}<span class="t-small">{e(f["hint"])}</span></label>'


def add_form() -> str:
    c = A['copy']
    return Section(e(c['addTitle']), '<div class="stack">' + f'<p class="t-small">{e(c["formJob"])}</p>'
                   + ''.join(field(f) for f in A['fields'])
                   + f'<div class="acts">{Btn(e(c["add"]), primary=True, small=False)}</div><div class="status" aria-live="polite"></div></div>')


def edit_form() -> str:
    c, row = A['copy'], A['rows'][0]
    fs = A['fields']
    fixed = row['sub'].rsplit(' · since', 1)[0]
    body = ('<div class="stack">' + field(fs[0]) + field(fs[1])
            + f'<div class="f"><span class="t-label">{e(fs[2]["label"])} · {e(fs[3]["label"])}</span><span class="ro">{e(fixed)}</span></div>'
            + field(fs[4]) + field(fs[5])
            + f'<div class="acts">{Btn(e(c["save"]), primary=True, small=False)}{Btn(e(c["cancelEdit"]), small=False)}</div>'
            + f'<div class="acts">{Btn(e(c["invite"]))}{Btn(e(c["refresh"]))}</div>'
            + f'<div class="status" aria-live="polite"><p class="t-small">{e(c["inviteCopied"])}</p></div></div>')
    return Section(e(A['editTitle']), body)


def requests_section() -> str:
    """The Requests Section above Your creators: open Rows (no Plate: the page's one Plate is the creators'), the channel as the
    title, platform, size and day under it, the applicant's own words in the foot, Approve and Decline in the trail."""
    rc = R['copy']
    rows = [Row(e(r['title']), e(r['sub']), link=False, foot=f'<span class="t-small">{e(r["about"])}</span>',
                trail=f'<span class="trail-end">{Btn(e(rc["approve"]), primary=True)}{Btn(e(rc["decline"]))}</span>') for r in R['rows']]
    return Section(e(rc['title']), Rows(rows, open_=True), meta=e(R['meta']))


def approve_form() -> str:
    """The rail in Approve mode: the add form prefilled from the request; no invite (the account is already bound)."""
    rc, fs, r = R['copy'], A['fields'], R['rows'][0]
    values = [r['title'].split('/')[-1], 'https://' + r['title'], r['title'].split('/')[-1], r['title'].split('/')[-1].upper(), '2026-12-09', '500']
    body = ('<div class="stack">' + f'<p class="t-small">{e(rc["approveJob"])}</p>'
            + ''.join(f'<label class="f"><span class="t-label">{e(f["label"])}</span><input class="field" type="text" value="{e(v)}">'
                      f'<span class="t-small">{e(f["hint"])}</span></label>' for f, v in zip(fs, values))
            + f'<div class="acts">{Btn(e(A["copy"]["add"]), primary=True, small=False)}{Btn(e(A["copy"]["cancelEdit"]), small=False)}</div></div>')
    return Section(e(f'Approve {r["title"]}'), body)


def admin_creators(rail: str) -> str:
    c = A['copy']
    main = f'<div class="stack" style="gap:32px">{requests_section()}{creators_plate()}</div>'
    body = PageHeader(e(c['pageTitle']), lead=e(c['job'])) + f'<div class="page-foot">{Page("split", main, rail)}</div>'
    return render('Creators · Forever Sixty', body, OWNER, '/admin/creators mock (the owner); data creators-facts.json (copy.ts, model.ts over fixtures/creators).')


# ================================================================================================ /admin/creators/payouts
def payouts(dialog: bool | str = False, off: bool = False) -> str:
    """off: the server's payout switch is off (CREATOR_PAYOUTS): the Notice above the Plate, Approve disabled."""
    c = P['copy']
    def trail(r: dict) -> str:
        act = Btn(e(r['action']), primary=True, disabled=off) if r['action'] else ''
        tag = Tag(e(r["notReady"]), "warn") if r["notReady"] else ''
        return f'<span class="trail-end">{tag}{act}{Btn(e(r["lines"]))}</span>'
    rows = [Row(e(r['name']), e(r['sub']), link=False, selected=(i == 0), trail=trail(r), foot=figures(r['figures'])) for i, r in enumerate(P['rows'])]
    control = f'<span aria-label="{e(c["monthLabel"])}">{Select(e(P["month"]))}</span>'
    notice = f'<p class="notice wipe" style="margin-bottom:16px">{e(P["off"])}</p>' if off else ''
    main = Section(e(P["title"]), notice + Plate(Rows(rows), e(P["month"])), meta=e(P["meta"]), action=control)
    lines = RowsSlim([Row(e(l['text']), link=False, slim=True, trail=f'<span class="amt{" wipe" if l["refund"] else ""}">{e(l["amount"])}</span>') for l in P["lines"]])
    rail = Section(e(P['linesTitle']), lines)
    head = PageHeader(e(c['pageTitle']), lead=e(c['job']), actions=Link('&lsaquo; ' + e(c['back'])))
    body = head + f'<div class="page-foot">{Page("split", main, rail)}</div>'
    if dialog:
        d = P['confirm']
        short = dialog == 'short'
        line = f'<p class="t-small">{e(d["balanceShortLine"] if short else d["balance"])}</p>'
        if short:
            line += f'<p class="t-small wipe" role="alert">{e(d["short"])}</p>'
        body += Dialog(e(d['title']), e(d['body']), line, Btn(e(d['cancel']), small=False) + Btn(e(d['pay']), primary=not short, small=False, disabled=short), 'pay-title')
    return render('Statements · Forever Sixty', body, OWNER, '/admin/creators/payouts mock (the owner, 2026-12-09: November closed); data creators-facts.json.')


# ================================================================================================ /creators/onboarding
def onboarding(state: str) -> str:
    c = O['copy']
    if state == 'invite':
        inner = f'<div class="stack"><p>{e(O["join"])}</p><p class="t-small">{e(c["stripeLead"])}</p><div class="acts">{Btn(e(c["accept"]), primary=True, small=False)}</div></div>'
    else:
        inner = f'<div class="stack"><p>{e(c["doneReady"])}</p><div class="acts">{Link(e(c["statsLink"]) + " &rsaquo;")}</div></div>'
    body = PageHeader(e(c['pageTitle'])) + f'<div class="page-foot">{PageForm(Section(e(c["inviteTitle"] if state == "invite" else c["stripeTitle"]), inner))}</div>'
    return render('Creator payouts · Forever Sixty', body, CREATOR, f'/creators/onboarding mock ({state}); data creators-facts.json.')


# ================================================================================================ /creators/stats
def stats() -> str:
    c = S['copy']
    summary = Plate(f'<div style="padding:16px">{figures(S["figures"])}</div>', e(c['summaryLabel']))
    players = Section(e(c['playersTitle']), RowsSlim([Row(e(w), link=False, slim=True) for w in S['weeks']]), meta=e(c['playersMeta']))
    paid = (RowsSlim([Row(e(p['text']), link=False, slim=True, trail=f'<span class="amt">{e(p["amount"])}</span>') for p in S['payouts']])
            if S['payouts'] else f'<p class="t-small">{e(S["payoutsEmpty"])}</p>')
    dash = f'<div class="acts" style="margin-top:16px">{Btn(e(S["dashboard"]), small=False)}</div>' if S['dashboard'] else ''
    payouts_ = Section(e(c['payoutsTitle']), paid + dash)
    body = PageHeader(e(S['name']), meta=[e(S['meta'])]) + f'<div class="page-foot">{Page("split", f'<div class="stack" style="gap:32px">{summary}{players}</div>', payouts_)}</div>'
    return render('Creator stats · Forever Sixty', body, CREATOR, '/creators/stats mock (the creator); data creators-facts.json.')


# ================================================================================================ /creators (public)
PAGE_CSS = '''<style>
/* the public page's layout only (no surfaces): the lead's inset on the panel, Apply under its sentence */
.cr-lead{padding:22px 24px 0}.cr-apply{padding-top:14px}.cr-ans{display:flex;flex-wrap:wrap;align-items:center;gap:10px 20px}.cr-lane{display:flex;flex-direction:column;gap:20px}.cr-lane dd small{display:block;font-size:13px;font-weight:400;color:var(--muted)}.cr-cap{display:block;font-size:12px;line-height:16px;color:var(--muted)}
main{padding-bottom:40px}@media(min-width:1024px){main{padding-bottom:64px}}
@media(max-width:767px){.cr-lead{padding:18px 16px 0}.cr-apply .btn{display:flex;width:100%}}
</style>'''


def paragraphs(lines: list[str]) -> str:
    return Prose(''.join(f'<p>{e(line)}</p>' for line in lines))


def facts(rows: list[dict]) -> str:
    """web ui/Facts: each figure's label and its value with the unit."""
    sub = lambda f: f'<small>{e(f["sub"])}</small>' if f.get('sub') else ''
    return '<dl class="kv">' + ''.join(f'<dt>{e(f["label"])}</dt><dd>{e(f["value"])}{sub(f)}</dd>' for f in rows) + '</dl>'


def lane(section_id: str, first: bool) -> str:
    ln = next((x for x in PR['lanes'] if x['section'] == section_id), None)
    if ln is None:
        return ''
    html = LaneBlock(e(ln['label']), facts(ln['figures']))
    if first:
        html += LaneBlock(e(PR['examplesLabel']), facts(PG['examples']), e(PR['examplesNote']))
    return f'<div class="cr-lane">{html}</div>'


def program() -> str:
    """The public page: a reading Page, the answer and Apply in the head, the lead, a ReadingSection per part with its figures."""
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
    html = head + Page('reading', ReadingPanel(body), toc, toc_from=1024, toc_label=e(PR['contentsLabel'])) + PAGE_CSS
    note = '/creators mock (public, signed out); data creators-facts.json page (lib/creators/copy.ts, model.ts). Apply goes to /creators/apply.'
    return build_page(PR['pageTitle'], html, '', None, note)


# ================================================================================================ /creators/apply
def apply(state: str) -> str:
    c = Y['copy']
    f = c['fields']
    if state == 'form':
        radios = ''.join(f'<label class="acts" style="min-height:44px"><input class="radio" type="radio" name="p"{" checked" if k == "twitch" else ""}> {e(v)}</label>'
                         for k, v in c['platforms'].items())
        inner = ('<div class="stack">'
                 + field(f['channel'])
                 + f'<fieldset class="f fs"><legend class="t-label">{e(f["platform"]["label"])}</legend><div class="acts">{radios}</div></fieldset>'
                 + field(f['audience'])
                 + f'<label class="f"><span class="t-label">{e(f["about"]["label"])}</span><textarea class="field" rows="5" style="height:118px;padding-top:10px"></textarea>'
                   f'<span class="t-small">{e(Y["counter"])}</span></label>'
                 + field(f['email'])
                 + f'<label class="acts" style="align-items:flex-start;min-height:44px"><input type="checkbox" class="radio"> <span>{e(c["disclose"])}</span></label>'
                 + f'<p class="t-small">{e(c["consent"])}</p>'
                 + f'<div class="acts">{Btn(e(c["submit"]), primary=True, small=False)}</div></div>')
        section = Section(e(c['formTitle']), inner)
    else:
        tag, tone, line = {'pending': (c['pendingTag'], 'muted', Y['pending']), 'approved': (c['approvedTag'], 'kill', c['approved']),
                           'declined': (c['declinedTag'], 'muted', c['declined'])}[state]
        act = f'<div class="acts">{Btn(e(c["setup"]), primary=True, small=False)}</div>' if state == 'approved' else ''
        section = Section(e(c['statusTitle']), f'<div class="stack"><div>{Tag(e(tag), tone)}</div><p>{e(line)}</p>{act}</div>')
    body = PageHeader(e(c['pageTitle'])) + f'<div class="page-foot">{PageForm(section)}</div>'
    return render('Apply · Forever Sixty', body, CREATOR, f'/creators/apply mock ({state}); data creators-facts.json (applyCopy).')


PAGES = {
    'web-admin-creators-approve': lambda: admin_creators(approve_form()),
    'web-creators': program,
    'web-creators-apply': lambda: apply('form'),
    'web-creators-apply-pending': lambda: apply('pending'),
    'web-creators-apply-approved': lambda: apply('approved'),
    'web-creators-apply-declined': lambda: apply('declined'),
    'web-admin-creators': lambda: admin_creators(add_form()),
    'web-admin-creators-edit': lambda: admin_creators(edit_form()),
    'web-admin-creators-payouts': lambda: payouts(),
    'web-admin-creators-confirm': lambda: payouts(dialog=True),
    'web-admin-creators-confirm-short': lambda: payouts(dialog='short'),
    'web-admin-creators-payouts-off': lambda: payouts(off=True),
    'web-creators-invite': lambda: onboarding('invite'),
    'web-creators-done': lambda: onboarding('done'),
    'web-creators-stats': stats,
}

if __name__ == '__main__':
    if sys.argv[1:2] == ['--list']:
        print(' '.join(PAGES))
    else:
        only = sys.argv[1:]
        for name, fn in PAGES.items():
            if not only or name in only:
                print(write(f'{name}.html', fn()))
