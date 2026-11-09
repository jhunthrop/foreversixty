"""Web-layout reference mocks for the creator program (design/specs/2026-11-08-creator-program.md), composed ONLY from ws_prims:
  - web-admin-creators: the owner's /admin/creators (PageHeader, a split Page: the creators' Plate of Rows, the Add a creator rail);
  - web-admin-creators-edit: the same with the rail in edit mode (read-only link and code, invite and refresh);
  - web-admin-creators-payouts: the owner's /admin/creators/payouts (the month Select, the statements' Plate, the Lines rail);
  - web-admin-creators-confirm: the payouts page with the Approve Dialog open;
  - web-creators-invite / web-creators-done: /creators/onboarding (Page form) accepting an invite, and back from Stripe;
  - web-creators-stats: the creator's /creators/stats.
Data and copy are read at render time from design/mocks/data/creators-facts.json (node design/mocks/creators_facts.mjs):

  node design/mocks/creators_facts.mjs && python3 design/mocks/ws_creators.py
FS_CREATORS_FACTS and FS_MOCK_OUT point the generator elsewhere.
"""
import json
import os
import sys
from html import escape
from pathlib import Path

import ws_prims
from ws_prims import (Button as _Button, Dialog, Figure, Link, Page, PageForm, PageHeader, Plate, Prose, ReadingSection, Row, Rows, RowsSlim, Section,
                      Select, Tag,
                      assert_no_surface_styles, write)
from ws_lib import page as _page

HERE = Path(__file__).resolve().parent
F = json.loads(Path(os.environ.get('FS_CREATORS_FACTS') or HERE / 'data' / 'creators-facts.json').read_text())
A, P, O, S = F['admin'], F['payouts'], F['onboarding'], F['stats']
R, G, Y = F['requests'], F['program'], F['apply']
OWNER = ('Thandrel', 'warrior', 'Level 60')
# The creator's pages are drawn signed in as the e2e owner fixture (me-60.json), so the chrome matches the build's.
CREATOR = OWNER

CSS = '''<style>
/* page-level layout only (no surfaces) */
.stack{display:flex;flex-direction:column;gap:12px;min-width:0}.acts{display:flex;gap:10px 12px;flex-wrap:wrap;align-items:center}
.figs5{display:flex;flex-wrap:wrap;gap:8px 20px}.figs5>span{width:104px;display:flex;flex-direction:column}.figs5>span.lead{width:168px}.btn[disabled]{opacity:.45;cursor:not-allowed}
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
def payouts(dialog: bool = False, off: bool = False) -> str:
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
        body += Dialog(e(d['title']), e(d['body']), '', Btn(e(d['cancel']), small=False) + Btn(e(d['pay']), primary=True, small=False), 'pay-title')
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
def program() -> str:
    """The public page: a document (Page reading, no contents rail), its sections from programCopy, Apply in the head and at the end."""
    secs = ''.join(ReadingSection(e(s['title']), Prose(''.join(f'<p>{e(p)}</p>' for p in s['paragraphs'])), anchor=s['id'])
                   for s in G['sections'])
    end = ReadingSection(e(G['applyTitle']), Prose(f'<p>{e(G["applyLead"])}</p>') + f'<div class="acts">{_Button(e(G["apply"]), True, False)}</div>',
                         anchor='apply')
    head = PageHeader(e(G['pageTitle']), lead=e(G['lead']), actions=_Button(e(G['apply']), True, False))
    body = head + f'<div class="page-foot">{Page("reading", secs + end, "", toc_from=0)}</div>'
    return render('Creators · Forever Sixty', body, None, '/creators mock (public, signed out); data creators-facts.json (programCopy).')


# ================================================================================================ /creators/apply
def apply(state: str) -> str:
    c = Y['copy']
    f = c['fields']
    if state == 'form':
        radios = ''.join(f'<label class="acts"><input class="radio" type="radio" name="p"{" checked" if k == "twitch" else ""}> {e(v)}</label>'
                         for k, v in c['platforms'].items())
        inner = ('<div class="stack">'
                 + field(f['channel'])
                 + f'<fieldset class="f fs"><legend class="t-label">{e(f["platform"]["label"])}</legend><div class="acts">{radios}</div></fieldset>'
                 + field(f['audience'])
                 + f'<label class="f"><span class="t-label">{e(f["about"]["label"])}</span><textarea class="field" rows="5" style="height:118px;padding-top:10px"></textarea>'
                   f'<span class="t-small">{e(f["about"]["hint"])} · {e(Y["counter"])}</span></label>'
                 + field(f['email'])
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
