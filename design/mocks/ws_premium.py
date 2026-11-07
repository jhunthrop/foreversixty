"""Web-layout mocks: /premium (the plans), /premium/checkout, /account and the API-rendered shared build /b/:id, composed ONLY
from ws_prims (the /b/:id page reuses ws_planner's own planner blocks, themselves composed from ws_prims).

Spec: design/specs/2026-11-08-premium-account.md. Data and copy are read at render time, never retyped:
  - design/mocks/data/premium-facts.json (node design/mocks/premium_facts.mjs): the plans copy and the plan cards' own model
    (priceCell, buttonLabel, costLine), the billing, checkout, account, Discord and guild copy modules, and the account page's
    models (accountPlanLine, guildRowPlanLine, guild-row-line, characterDescriptor, setupAnswer, deviceStatusLine,
    premiumRowValue, premiumOwnedView, offerStatusLine) evaluated on one account per state at 2026-11-08T10:00Z, UTC;
  - design/mocks/data/planner-facts.json (node design/mocks/planner_facts.mjs) for the shared builds' trees and copy;
  - api/internal/site/handler.go (the API's own not-found and unavailable sentences) and web/src/pages/b-unavailable.astro.
Strings no module carries yet are the PROPOSED dict below (spec section 8.2); nothing else is typed.

  node design/mocks/premium_facts.mjs && python3 design/mocks/ws_premium.py
  for p in $(python3 design/mocks/ws_premium.py --list); do node design/mocks/ws_shoot.mjs $p 390,1440,1920,2000; done
  python3 design/mocks/ws_premium.py compare <before dir>      (before: node design/mocks/shoot_premium_before.mjs <dir>)
FS_PREMIUM_FACTS and FS_MOCK_OUT point the generator elsewhere (a parity spec's test-time render).
"""
import json
import os
import re
import sys
from html import escape
from pathlib import Path

from mocklib import CLASS_COLOR
import ws_prims
from ws_prims import (Button as _Button, Columns, Emblem, Facts, Figure, Link, Notice, Page, PageForm, PageHeader, Plate, PlateColumns,
                      PlanHead, Prose, Row, RowGrid, Rows, RowsSlim, Section, Skeleton, Steps, Switch, Tag, assert_no_surface_styles,
                      control_bar, crest_img, icon, seg, sel, write, REGION)
from ws_lib import page as _page

HERE = Path(__file__).resolve().parent
ROOT = HERE.parents[1]
F = json.loads(Path(os.environ.get('FS_PREMIUM_FACTS') or HERE / 'data' / 'premium-facts.json').read_text())
P, CK, BB, PB, SRC, GC = F['plans'], F['checkout'], F['billing_block'], F['premium_block'], F['source'], F['guild_checkout']
AP, ACC, CL, DC, DL, DR, NC, CON = F['account_plan'], F['account'], F['characters_copy'], F['devices_copy'], F['discord_link'], F['discord_role'], F['notifications'], F['consent']
A = F['accounts']

# Strings the copy modules do not carry yet (spec 8.2). Each becomes a key in the module named beside it.
PROPOSED = {
    'guideLink': 'Talent builds in the class guides',                   # the API message page's second link (page.go)
    'askSender': 'Ask whoever sent it for a fresh link.',               # the API not-found page's line (ruling Q12, page.go)
}
# The 2026-11-08 rulings' strings now live in the copy modules; these are their keys in the facts.
YOUR_PLAN, FAQ_TITLE = F['plans']['yourPlanTag'], F['plans']['faqTitle']
STEPS = F['checkout']['steps']


def e(t) -> str:
    return escape(str(t), quote=False)


def r(t) -> str:
    return Link(f'{e(t)} &rsaquo;')


def Btn(label: str, primary: bool = False, small: bool = False, tone: str = '', off: bool = False) -> str:
    """The system Button as the build renders an action (a <button>); a navigation stays the <a> ws_prims draws."""
    b = _Button(label, primary, small, tone).replace('<a ', '<button type="button" ', 1).replace(' href="#"', '', 1).replace('</a>', '</button>')
    return b.replace('class="btn', 'aria-disabled="true" class="btn is-off', 1) if off else b


def LinkBtn(label: str, primary: bool = False, small: bool = False) -> str:
    return _Button(label, primary, small)


CSS = '''<style>
/* page-level layout only (no surfaces) */
section>.sec-h+*{margin-top:12px}
.stack{display:flex;flex-direction:column;gap:12px;min-width:0}.stack.tight{gap:6px}
.acts{display:flex;gap:10px 16px;flex-wrap:wrap;align-items:center}
.lede{font-size:15px;line-height:24px;max-width:72ch}.t-small{max-width:80ch}
.fd{display:inline-flex;gap:8px;align-items:center;flex-wrap:wrap}
.pcol .rprose.ruled li{padding-block:8px}.pcol .rprose.lane{font-size:15px;line-height:22px}
.pcol .btn{align-self:flex-start}
@media(max-width:767px){.pcol .btn{align-self:stretch}}
.ans-who{display:flex;align-items:center;gap:12px}.ans-who b{font:700 18px/22px var(--font-body)}.ans-links{display:flex;gap:16px;flex-wrap:wrap}
@media(min-width:900px){.head .ans{min-width:300px}.ans-who{justify-content:flex-end}.ans-links{justify-content:flex-end}}
.ans-plan{display:flex;flex-direction:column;gap:2px}
@media(min-width:900px){.ans-plan{align-items:flex-end}}
.coach .opening{font-size:15px;line-height:23px}
.coach .who{display:flex;gap:12px;align-items:center}.coach .who b{font:700 16px/20px var(--font-body);color:var(--strong)}
.alert{color:var(--strong)}.wipe-t{color:var(--wipe)}
.rowsel{display:flex;align-items:center;gap:8px}
.swt{display:flex;align-items:center;justify-content:space-between;gap:12px;min-height:44px}
.planfig{display:flex;align-items:flex-end;justify-content:space-between;gap:12px;flex-wrap:wrap}
.dev .row .sub{white-space:normal}
.msg{display:flex;flex-direction:column;gap:12px}
@media(max-width:1023px){.btn.sm{height:44px}}
.statesheet{display:grid;gap:40px}.st-grid{display:grid;gap:24px 40px;grid-template-columns:repeat(auto-fill,minmax(340px,1fr))}.st-cell{display:flex;flex-direction:column;gap:6px;min-width:0}
</style>'''


def render(title: str, body: str, nav: str, who, chrome_state: str, note: str, extra_css: str = '') -> str:
    assert_no_surface_styles(body)
    return _page(title, body + CSS + extra_css + f'<style>{ws_prims.PRIM_CSS}</style>', nav, who, note, chrome_state)


WHO = {'free': ('Felsnap', 'hunter', 'Level 34'), 'premium': ('Thandrel', 'warrior', 'Level 60')}
def who_of(key: str):
    if key in (None, 'out', 'new', 'loading', 'error'):
        return None
    return WHO['premium'] if key in ('premium', 'claim') else WHO['free']


def chrome_of(key):
    return {'new': 'new', 'loading': 'pending', 'error': 'pending'}.get(key, '')


# ================================================================================================ /premium
def ruled(lines: list[str], anchors: list = ()) -> str:
    lis = ''.join(f'<li{f" id={a}" if a else ""}>{e(t)}</li>' for t, a in zip(lines, list(anchors) + [''] * len(lines)))
    return Prose(f'<ul>{lis}</ul>', 'ruled', 'lane')


def price_key(plan: str, interval: str, founder: bool) -> dict:
    return F['priced'][f'{plan}:{interval}:{"founder" if founder else "list"}']


def premium_column(st: dict) -> tuple:
    pc = price_key('premium', st['interval'], st['founder'])
    c = pc['cell']
    lines = [(e(c['secondLine']), '')] + ([(e(c['founderLine']), 'gold')] if c['founderLine'] else [])
    owned = st.get('owned')
    # Ruling PR2: founder is said once (the notice and the price line); the column carries only `Your plan`.
    tag = Tag(e(YOUR_PLAN), 'kill') if owned else ''
    head = PlanHead(e(P['premium']['name']), e(P['premium']['audience']), e(c['amount']), e(c['suffix'].strip()), lines, tag, gold=True)
    if owned:
        act = f'<p class="lede strong">{e(SRC["premiumCard"][owned])}</p>{LinkBtn(e(SRC["buttons"]["seePlan"]))}'
    else:
        act = LinkBtn(e(pc['button']), primary=True)
    feats = ruled([l['text'] for l in F['premium_lines']], [l['anchor'] if l['first'] else '' for l in F['premium_lines']])
    body = f'<div class="stack">{head}{act}<p class="t-small">{e(P["premium"]["support"])}</p>{feats}<p class="t-small">{e(P["premium"]["scope"])}</p></div>'
    return ('', '', body, '')


def free_column(st: dict) -> tuple:
    mine = st.get('signed') and not st.get('owned')
    tag = Tag(e(YOUR_PLAN), 'kill') if mine else ''
    head = PlanHead(e(P['free']['name']), e(P['free']['audience']), e(P['free']['price']), '', [(e(P['free']['note']), '')], tag)
    foot = '' if st.get('signed') else r(F['my_reports']['signInAction'])
    body = f'<div class="stack">{head}{ruled(P["free"]["features"])}</div>'
    return ('', '', body, foot)


def guild_column(plan: str, st: dict) -> tuple:
    block = P['guild'] if plan == 'guild' else P['raidGuild']
    pc = price_key(plan, st['interval'], st['founder'])
    c = pc['cell']
    lines = [(e(c['secondLine']), '')] + ([(e(c['founderLine']), 'gold')] if c['founderLine'] else []) + [(e(pc['cost']), '')]
    on = st.get('guild_on') == plan
    tag = Tag('Olympus XXVII', 'kill') if on else ''
    head = PlanHead(e(block['name']), e(block['audience']), e(c['amount']), e(c['suffix'].strip()), lines, tag)
    if on:
        act = f'<p class="lede strong">{e(F["guild_on_line"])}</p>{Btn(e(GC["onPlan"]["manageBilling"]))}'
    else:
        act = LinkBtn(e(pc['button']))
    trial = f'<p class="t-small"><b class="strong">{e(F["trial"]["lead"])}</b>{e(F["trial"]["rest"])}</p>'
    body = f'<div class="stack">{head}{act}{trial}{ruled(block["features"])}</div>'
    return ('', '', body, '')


def coach_rail() -> str:
    c = F['coach']
    cp = P['coachSample']
    fix = c['fixFirst']
    then = ''.join(Row(e(t['title']), e(t['why']), lead=icon(t['icon'], 32) if t['icon'] else '', link=False) for t in c['then'])
    well = ''.join(Row(e(w['text']), lead=icon(w['icon'], 24) if w['icon'] else '', link=False, slim=True) for w in c['doingWell'])
    body = (f'<div class="stack coach" id="coach"><div class="who">{crest_img(c["classSlug"], 40)}<span><b>{e(c["title"])}</b><span class="t-small" style="display:block">{e(c["context"])}</span></span></div>'
            f'<p class="opening">{e(c["opening"])}</p><span class="t-label">{e(cp["fixFirst"])}</span>'
            + Rows([Row(e(fix['title']), e(fix['why']), lead=icon(fix['icon'], 40), link=False, selected=True)])
            + f'<span class="t-label">{e(cp["then"])}</span>{Rows([then], True)}<span class="t-label">{e(cp["doingWell"])}</span>{Rows([well], True)}'
            f'<p class="t-small">{e(cp["caption"])}</p></div>')
    return Section(e(cp['label']), body, meta=Tag(e(cp['tag']), 'muted'))


def premium_answer(key: str) -> str:
    """Signed in: the plan the account is on now and the way to it (tenet 14: the player first)."""
    if key in (None, 'out'):
        return ''
    a = A[key]
    h = a['plan']['head']
    when = f'{h["when"]["label"]} {h["when"]["date"]}' if h['when'] else ''
    sub = ' · '.join(x for x in [when, h['extra'] or ''] if x)
    tag = Tag(e(AP['founder']['pill'])) if a['plan']['founder'] else ''
    return (f'<div class="ans-plan">{Figure(e(h["name"]), e(ACC["planLabel"]), size="sm")}'
            + (f'<span class="t-small">{e(sub)} {tag}</span>' if sub or tag else '') + f'{r(SRC["buttons"]["seePlan"])}</div>')


PREMIUM_STATES = {
    'out': dict(interval='yearly', founder=True),
    'monthly': dict(interval='monthly', founder=True),
    'closed': dict(interval='yearly', founder=False),
    'free': dict(interval='yearly', founder=True, signed='free'),
    'own': dict(interval='yearly', founder=True, signed='premium', owned='own', guild_on='guild'),
    'seat': dict(interval='yearly', founder=True, signed='seat', owned='seat'),
    'grant': dict(interval='yearly', founder=True, signed='grant', owned='grant'),
    'canceled': dict(interval='yearly', founder=True, notice='canceled'),
    'pastdue': dict(interval='yearly', founder=True, signed='pastdue', owned='own', notice='pastdue'),
    'ended': dict(interval='yearly', founder=True, signed='ended', notice='ended'),
}


def premium_page(state: str) -> str:
    st = PREMIUM_STATES[state]
    key = st.get('signed')
    head = PageHeader(e(P['title']), answer=premium_answer(key), lead=e(P['band']['lede']))
    iv = seg([(e(P['interval']['labels']['yearly']), st['interval'] == 'yearly', None, ''), (e(P['interval']['labels']['monthly']), st['interval'] == 'monthly', None, '')], label=e(P['interval']['groupLabel']))
    end = f'<span class="t-small">{e(P["band"]["guildLead"])}</span>{Link(e(P["band"]["guildLink"]) + " &rsaquo;")}'
    bar = control_bar([(e(P['interval']['groupLabel']), iv)], f'{e(P["interval"]["labels"][st["interval"]])}', end, label=e(P['interval']['groupLabel']))
    notices = ''
    if st['founder']:
        notices += Notice(f'{Tag(e(P["founder"]["label"]))} {e(P["founder"]["open"])}', tone='neutral')
    n = st.get('notice')
    if n == 'canceled':
        notices += Notice(f'<span role="status">{e(CK["canceled"])}</span>', tone='neutral')
    elif n == 'pastdue':
        notices += Notice(f'{Tag(e(BB["pastDueTag"]), "wipe")} <span class="alert" role="alert">{e(BB["pastDueBanner"])}</span>', r(SRC['buttons']['manage']), tone='neutral')
    elif n == 'ended':
        en = F['seat_ended_line']
        notices += Notice(e(en['text']), r(en['action']), tone='neutral')
    notes = f'<div data-region="{REGION["notes"]}">{notices}</div>' if notices else ''
    plate = Plate(Columns([free_column(st), premium_column(st)], open_=False, second_first=True), 'Plans for a player')
    gr = P['guildRegion']
    guild = Section(e(gr['heading']), f'<div class="stack" id="guild-leaders"><p class="lede">{e(gr["lede"])}</p><p class="t-small">{e(gr["members"])}</p>'
                    + Columns([guild_column('guild', st), guild_column('raid_guild', st)]) + f'<p class="t-small">{e(P["historyPromise"])}</p></div>')
    rules = Section(e(P['howItWorks']), RowGrid([Row(e(x['term']), e(x['sentence']), link=False) for x in P['rules']], 2))
    faq = Section(e(FAQ_TITLE), RowGrid([Row(e(x['question']), e(x['answer']), link=False) for x in P['faq']], 2), meta=e(P['foot']))
    body = head + bar + notes + Page('split', plate + guild + rules + faq, coach_rail())
    note = f'/premium mock, {state}. Data: premium-facts.json (plans-page-copy.ts, components/premium/model.ts on plans.ts; accounts DERIVED from home fixtures at 2026-11-08T10:00Z).'
    return render(f'{P["title"]} · Forever Sixty', body, 'Premium', who_of(key), chrome_of(key), note)


# ================================================================================================ /premium/checkout
def band(plan: str, founder: bool = True, change: bool = True) -> str:
    """The checkout head's answer: the plan picked on /premium, its price and the way back (CheckoutBand's summary)."""
    if plan == 'premium':
        pc = price_key('premium', 'yearly', founder)
        return f'<div class="ans-plan"><span class="fd">{Figure(e(pc["cell"]["amount"]), unit=e(pc["cell"]["suffix"].strip()), tone="gold")}{Tag(e(GC["planSummary"]["founder"])) if founder else ""}</span><span class="t-small">{e(pc["cell"]["secondLine"])}</span>{r(GC["planSummary"]["changePlan"]) if change else ""}</div>'
    pc = price_key('guild', 'yearly', founder)
    return (f'<div class="ans-plan"><span class="fd">{Figure(e(pc["cell"]["amount"]), "Guild", e(pc["cell"]["suffix"].strip()), tone="gold")}{Tag(e(GC["planSummary"]["founder"]))}</span>'
            f'<span class="t-small">{e(GC["sizeLine"]["guild"])}</span><span class="t-small"><b class="strong">{e(GC["planSummary"]["trial"])}</b></span>{r(GC["planSummary"]["changePlan"]) if change else ""}</div>')


def checkout_steps(state: str) -> str:
    s1 = ('ok', 1, e(STEPS['signIn']), f'<span class="tx">{e(F["checkout_signed_in"])}</span>', '')
    pay, on, paid = e(STEPS['pay']), e(F['plan_table']['premium']['name'] + ' on'), e(CK['paid'])
    if state == 'redirecting':
        rows = [s1, ('busy', 2, pay, f'<span class="tx" role="status">{e(CK["redirecting"])}</span><br>{e(CK["stripeLine"])}', ''), ('todo', 3, on, '', '')]
    elif state == 'processing':
        rows = [s1, ('ok', 2, pay, paid, ''), ('busy', 3, on, f'<span class="tx" role="status">{e(CK["processing"])}</span>', '')]
    elif state == 'confirmed':
        rows = [s1, ('ok', 2, pay, paid, ''), ('ok', 3, on, f'<span class="tx" role="status">{e(CK["confirmed"])}</span>', '')]
    elif state == 'slow':
        rows = [s1, ('ok', 2, pay, paid, ''), ('warn', 3, on, f'<span class="warn" role="status">{e(CK["confirmationSlow"])}</span>', Btn(e(CK['checkAgain']), primary=True))]
    elif state == 'error':
        rows = [s1, ('miss', 2, pay, f'<span class="miss" role="alert">{e(CK["genericFailure"])}</span>', Btn(e(CK['retry']), primary=True)), ('todo', 3, on, '', '')]
    else:  # notopen
        rows = [s1, ('miss', 2, pay, f'<span class="miss" role="alert">{e(CK["notOpenYet"])}</span>', ''), ('todo', 3, on, '', '')]
    return Plate(Steps(rows), 'Checkout')


def checkout_page(state: str) -> str:
    who, chrome = WHO['premium'], ''
    after = ''
    if state in ('redirecting', 'processing', 'confirmed', 'slow', 'error', 'notopen'):
        head = PageHeader(e(F['plan_table']['premium']['name']), answer=band('premium', change=state in ('redirecting', 'error', 'notopen')))
        plate = checkout_steps(state)
        if state == 'confirmed':
            # Ruling K2: the renewal line above the three next steps.
            after = (f'<p class="t-small" style="padding-top:16px">{e(F["checkout_renewal"])}</p>'
                     f'<div class="acts" style="padding-top:8px">{r(CK["nextSim"])}{r(CK["nextCoach"])}{r(SRC["buttons"]["seePlan"])}</div>')
        body = head + PageForm(plate + after)
    elif state == 'trial':
        # Ruling K1: a guild checkout in its free trial took no payment; row 2 says the card was saved.
        head = PageHeader(e(GC['band']['title']), answer=band('guild', change=False))
        rows = [('ok', 1, e(STEPS['signIn']), f'<span class="tx">{e(F["checkout_signed_in"])}</span>', ''),
                ('ok', 2, e(STEPS['pay']), e(F['checkout_trial_saved']), ''),
                ('ok', 3, e(F['plan_table']['guild']['name'] + ' on'), f'<span class="tx" role="status">{e(CK["confirmed"])}</span>', '')]
        body = head + PageForm(Plate(Steps(rows), 'Checkout'))
    elif state == 'owned':
        o = A['seat']['owned']
        who = WHO['free']
        head = PageHeader(e(F['plan_table']['premium']['name']), answer=band('premium'))
        inner = (f'<div class="stack" style="padding:16px 20px 20px"><p class="t-h3">{e(o["title"])}</p><p class="lede strong">{e(o["line"])}</p>'
                 + (f'<p class="t-small">{e(o["line2"])}</p>' if o['line2'] else '') + f'<div class="acts">{LinkBtn(e(o["action"]["label"]))}</div></div>')
        body = head + PageForm(Plate(inner, 'Your plan'))
    elif state == 'guild':
        fn = F['guild_checkout_fns']
        head = PageHeader(e(GC['band']['title']), answer=band('guild'), lead=e(GC['lead']['someClaimed']))
        g1 = Row('Olympus XXVII', f'{e(fn["holdsClaim"])}<br>{e(fn["verified"])} · {e(fn["onRoster"])} · {e(fn["noPlan"])}', lead=Emblem('alliance', 40),
                 trail=Btn(e(fn['startTrial']), primary=True), link=False)
        g2 = Row('Pinefeather Lodge', f'{e(fn["askLeader"])}<br>{e(fn["verified4"])} · {e(fn["onRoster4"])} · {e(fn["noPlan"])}', lead=Emblem('alliance', 40),
                 trail=Btn(e(GC['copyField']['copy']), small=True), link=False)
        foot = f'<p class="t-small" style="padding-top:14px">{e(GC["foot"]["missingGuild"])} {r(GC["foot"]["getSetUp"])}</p><p class="t-small">{e(fn["footPremium"])} {r(GC["foot"]["getPremium"])}</p>'
        body = head + Page('ledger', Plate(Rows([g1, g2]), e(CK['whichGuild'])) + foot)
    else:  # invalid
        head = PageHeader(e(GC['band']['invalidTitle']))
        body = head + PageForm(f'<p class="lede" role="alert">{e(GC["invalidLink"]["text"])} {r(GC["invalidLink"]["button"])}</p>')
    note = f'/premium/checkout mock, {state}. Data: premium-facts.json (lib/billing/copy.ts checkoutCopy, checkout-guild-copy.ts, premium-owned.ts on the seat account).'
    return render('Checkout · Forever Sixty', body, 'Premium', who, chrome, note)


# ================================================================================================ /account
def account_answer(a: dict) -> str:
    x = a['answer']
    if not x:
        return ''
    cls = x['classSlug']
    links = r(x['open']['label']) + (r(x['bis']['label']) if x['bis'] else '')
    return (f'<span class="ans-who">{crest_img(cls, 40)}<span><b style="color:{CLASS_COLOR[cls]}">{e(x["name"])}</b>'
            f'<span class="t-small" style="display:block">{e(x["levelLine"])}</span></span></span><span class="ans-links">{links}</span>')


def characters_section(a: dict, refreshing: bool = False) -> str:
    rows = []
    for c in a['characters']:
        tags = (Tag(e(CL['current'])) if c['main'] else '') + (Tag(e(CL['main']), 'muted') if c['main'] else '')
        src = Tag(e(CL['addonSource'] if c['source'] == 'addon' else (CL['battlenetSource'] if c['source'] else CL['noBuildYet'])), 'muted')
        setmain = '' if c['main'] else Btn(e(CL['setAsMain']), small=True)
        sub = e(c['descriptor']) + (f'<br>{e(c["guild"])}' if c['guild'] else '')
        rows.append(Row(f'<span class="fd"><span style="color:{CLASS_COLOR.get(c["cls"], "#f2eee4")}">{e(c["name"])}</span>{tags}</span>', sub, lead=crest_img(c['cls'], 40),
                        trail=f'{src}{setmain}{r(CL['open'])}', link=False))
    refresh = f'<span class="t-small">{e(CL["refreshing"])}</span>' if refreshing else r(CL['refreshFromBattlenet'])
    if not rows:
        inner = (f'<div class="stack" style="padding:16px 20px 20px"><p class="lede warn-t" style="color:var(--gold)">{e(CL["empty"])}</p>'
                 f'<div class="acts">{LinkBtn(e(CL["getSetUp"]), primary=True)}{LinkBtn(e(CL["pasteAnExport"]))}</div></div>')
        return Section(e(CL['heading']), Plate(inner, e(CL['heading'])), action=refresh)
    foot = (f'<p class="t-small" style="padding-top:12px">{e(CL["introBattlenetLine"])} {Link(e(CL["introInstallAddonLink"]))} {e(CL["introAddonTail"])} '
            f'{Link(e(CL["introPasteLink"]))} {e(CL["introPasteTail"])}</p><p style="padding-top:8px">{r(ACC["reportsLink"])}</p>')
    return Section(e(CL['heading']), Plate(Rows(rows), e(CL['heading'])) + foot, action=refresh, meta=e(a['imported'] or ''))


def plan_section(a: dict, offer_state: str = '') -> str:
    p = a['plan']
    h = p['head']
    tag = f' {Tag(e(AP["founder"]["pill"]))}' if p['founder'] else ''
    when = f'<div class="ans-plan">{Figure(e(h["when"]["date"]), e(h["when"]["label"]), size="sm")}</div>' if h['when'] else ''
    top = f'<div class="planfig"><div class="fd">{Figure(e(h["name"]), size="sm")}{tag}</div>{when}</div>'
    lines = ''
    if p['ended']:
        lines += f'<p class="lede">{e(p["ended"]["text"])}{Link(e(p["ended"]["action"]))}</p>'
    else:
        if h['extra']:
            lines += f'<p class="t-small">{e(h["extra"])}</p>'
        if p['billingButton'] == 'manage' and h['when'] and h['when']['label'] == AP['head']['renews'] and a['source'] != 'own' or (p['billingButton'] == 'manage' and h['when'] and h['when']['label'] == 'Renews'):
            lines += f'<p class="t-small">{e(AP["head"]["cancelAnytime"])}</p>'
    pastdue = a['key'] == 'pastdue'
    if pastdue:
        lines += f'<p class="fd" role="alert">{Tag(e(BB["pastDueTag"]), "wipe")}<span class="alert">{e(BB["pastDueBanner"])}</span></p>'
    elif p['line2']:
        lines += f'<p class="t-small">{e(p["line2"])}</p>'
    offer = ''
    o = a['offer']
    if o:
        if offer_state:
            offer = f'<p class="t-small" role="status">{e(o[offer_state])}</p>'
        else:
            offer = Rows([Row(e(o['title']), f'{e(o["from"])}<br><span class="tx" style="color:var(--text)">{e(o["consequence"])}</span>', link=False, selected=True,
                              foot=f'{Btn(e(AP["offer"]["accept"]), primary=True)}{Btn(e(AP["offer"]["decline"]))}')])
    ent = a['entitlements'] or {}
    pr = PB['rows']
    def benefit(label, on, off, live, linked=True):
        v = on if live else off
        return Row(e(label), trail=(r(v) if live and linked else f'<span class="t-small">{e(v)}</span>'), link=False, slim=True)
    rows = [benefit(pr['serverSims']['label'], a['serverSimsOn'], pr['serverSims']['off'], ent.get('server_sims')),
            benefit(pr['history']['label'], pr['history']['onWithMain'], pr['history']['off'], ent.get('history'))]
    if a['premiumRow']:
        rows.append(benefit(pr['notifications']['label'], a['premiumRow']['value'], a['premiumRow']['value'], a['premiumRow']['linked']))
    rows.append(benefit(pr['supporter']['label'], pr['supporter']['on'], pr['supporter']['off'], ent.get('supporter_mark')))
    acts = []
    if p['billingButton'] == 'manage':
        acts.append(Btn(e(SRC['buttons']['manage']), primary=pastdue))
    elif p['billingButton'] == 'history':
        acts.append(Btn(e(SRC['buttons']['history'])))
    if a['source'] == 'none' and not p['ended']:
        acts.append(LinkBtn(e(BB['seePlans']), primary=True))
    body = f'<div class="stack" id="plan">{top}{lines}{offer}{RowsSlim(rows)}<div class="acts">{"".join(acts)}</div></div>'
    return Section(e(ACC['planLabel']), body)


def guilds_section(a: dict, asking: bool = False) -> str:
    if not a['guilds']:
        return ''
    blocks = []
    for i, g in enumerate(a['guilds']):
        tags = (Tag(e(g['tag']), 'gold' if g['planLine'] else 'muted') if g['tag'] else '') + (Tag(e(g['mark']), 'muted') if g['mark'] else '') \
            + (Tag(e(AP['guildRow']['claimPending'])) if g['claimLine'] else '')
        rows = [Row(f'<span class="fd">{e(g["name"])}{tags}</span>', e(g['rank'] or ''), lead=Emblem('alliance', 32), trail=r(CL['open']), link=False)]
        rows.append(Row(e(CON['shareLabel']), trail=sel(e({'roster': CON['roster'], 'gear': CON['gear'], 'gear_bags': CON['gearBags']}[g['consent']])), link=False, slim=True))
        if g['planLine']:
            rows.append(Row(e(F['guild_settings']), trail=r(g['planLine']), link=False, slim=True))
        if g['claimLine']:
            rows.append(Row(e(g['claimLine']), trail=Btn(e(GC['copyField']['copy']), small=True), link=False))
        if g['next']:
            rows.append(Row(e(g['next']['text']), trail=r(g['next']['action']), link=False, slim=True))
        if asking and i == 0:
            conf = f'<p class="lede">{e(g["leaveLine"])}</p>' + (f'<p class="lede">{e(g["leaveOfficer"])}</p>' if g['leaveOfficer'] else '')
            leave = f'<div class="stack">{conf}<div class="acts">{Btn(e(CON["leaveConfirmButton"]), tone="danger")}{Btn(e(CON["leaveCancel"]))}</div></div>'
        else:
            leave = f'<div class="acts">{Btn(e(CON["leave"]), small=True, tone="danger")}</div>'
        blocks.append(RowsSlim(rows, True) + leave)
    return Section(e(ACC['guildLabel']), f'<div class="stack">{"".join(blocks)}</div>')


def devices_section(dev: str, revoking: bool = False, pairing: bool = False) -> str:
    d = F['devices'][dev]
    rows = []
    for i, (name, line) in enumerate(zip(d['names'], d['lines'])):
        sub = line.split(' · ', 1)[1]
        if revoking and i == 0:
            rows.append(Row(e(d['revoke'][0]), trail=f'{Btn(e(DC["revoke"]), small=True, tone="danger")}{Btn(e(DC["cancel"]), small=True)}', link=False))
        else:
            rows.append(Row(e(name), e(sub), trail=Btn(e(DC['revoke']), small=True), link=False))
    if pairing:
        tail = f'<p class="t-small">{e(DC["codeFallbackLine"])}</p><code class="c big">4821-9930</code>'
    elif not rows:
        tail = f'<p class="t-small">{e(ACC["noDevices"])}</p><div class="acts">{LinkBtn(e(ACC["downloadCompanion"]))}{Link(e(d["entry"]) + " &rsaquo;")}</div>'
    else:
        tail = f'<div class="acts">{Link(e(d["entry"]) + " &rsaquo;")}</div>'
    return Section(e(ACC['devicesLabel']), f'<div class="stack dev">{Rows(rows, True) if rows else ""}{tail}</div>')


def discord_section(a: dict, unlinking: bool = False) -> str:
    link = a['discord']['link']
    if link and unlinking:
        top = Row(e(DL['unlinkQuestion']), trail=f'{Btn(e(DL["unlink"]), small=True, tone="danger")}{Btn(e(DL["cancel"]), small=True)}', link=False)
    elif link:
        top = Row(f'{e(DL["linkedLead"])} <b class="strong">{e(link["display_name"])}</b>', trail=Btn(e(DL['unlink']), small=True), link=False)
    else:
        top = Row(e(DL['lead']), trail=Btn(e(DL['link']), small=True), link=False)
    # Ruling A4 / MUST 3: an account that already has Premium is told the role turns on as soon as it links.
    lead = DR['leadOwned'] if a['source'] != 'none' else DR['lead']
    role = Row(e(DR['label']), e(lead), trail=Btn(e(DR['link']), small=True), link=False)
    st = a['discord']['state']
    if st == 'saved':
        nt = (f'<div class="stack tight"><span class="t-label">{e(NC["headLabel"])}</span><p class="lede">{e(NC["destinationChannel"])}</p>'
              f'{Switch(e(NC["eventSimFinished"]))}{Switch(e(NC["eventUploadParsed"]))}{Switch(e(NC["eventUploadPrivate"]), on=False)}'
              f'<div class="acts">{Btn(e(NC["test"]), small=True)}{Link(e(NC["reconnect"]) + " &rsaquo;")}</div></div>')
    else:
        nt = f'<div class="stack tight"><span class="t-label">{e(NC["headLabel"])}</span><p class="t-small">{e(NC["freeSentence"])}</p><p>{r(NC["seePlans"])}</p></div>'
    return Section(e(NC['label']), f'<div class="stack" id="discord">{Rows([top, role], True)}{nt}</div>')


def you_section(a: dict) -> str:
    sw = Switch(e(ACC['pseudonymLabel']), on=False)
    rows = [Row(e(ACC['deleteTitle']), e(ACC['deleteBody']), link=False)]
    return Section(e(ACC['youLabel']), f'<div class="stack">{sw}<p class="t-small">{e(ACC["pseudonymNote"])}</p>{Rows(rows, True)}</div>')


ACCOUNT_STATES = {
    'free': dict(dev='none'), 'several': dict(dev='several'), 'premium': dict(dev='one'), 'seat': dict(dev='one'),
    'offer': dict(dev='one'), 'ended': dict(dev='one'), 'pastdue': dict(dev='one'), 'claim': dict(dev='one'), 'new': dict(dev='none'),
    'confirm': dict(dev='one', acct='premium', confirm=True), 'grant': dict(dev='one'),
}


def account_page(state: str) -> str:
    if state == 'out':
        head = PageHeader(e(ACC['title']))
        inner = f'<div class="stack" style="padding:16px 20px 20px"><p class="lede">{e(F["signin"]["reason"])}</p><div class="acts">{LinkBtn(e(F['my_reports']['signInAction']), primary=True)}</div></div>'
        body = head + PageForm(Plate(inner, 'Sign in'))
        return render(f'{ACC["title"]} · Forever Sixty', body, '', None, '', '/account mock, signed out. Copy: signin-copy.ts.')
    if state in ('loading', 'error'):
        head = PageHeader(e(ACC['title']), meta=[Skeleton(220, 14)] if state == 'loading' else None)
        if state == 'loading':
            sk = lambda n: ''.join(Skeleton(560, 14) + '<br>' for _ in range(n))
            plate = Section(e(CL['heading']), Plate(f'<div class="stack" style="padding:16px 20px" role="status" aria-busy="true">{"".join(Skeleton(480, 56) for _ in range(2))}</div>', e(CL['heading'])))
            rail = ''.join(Section(t, f'<div class="stack">{Skeleton(300, 20)}{Skeleton(240, 20)}{Skeleton(160, 44)}</div>') for t in (ACC['planLabel'], ACC['guildLabel'], ACC['devicesLabel'], NC['label']))
            body = head + Page('split', plate, rail)
        else:
            inner = f'<div class="stack" style="padding:16px 20px 20px"><p class="alert" role="alert">{e(F["account_failed"])}</p><div class="acts">{Btn(e(CK['retry']), primary=True)}</div></div>'
            body = head + Page('split', Section(e(CL['heading']), Plate(inner, e(CL['heading']))), '')
        return render(f'{ACC["title"]} · Forever Sixty', body, '', None, 'pending', f'/account mock, {state}.')
    st = ACCOUNT_STATES[state]
    a = A[st.get('acct', state)]
    confirm = st.get('confirm', False)
    meta = [e(a['battletag']), e(ACC['signedInWithBattlenet']), Btn(e(ACC['signOut']), small=True)]
    head = PageHeader(e(ACC['title']), meta=meta, answer=account_answer(a))
    # Main: the characters (the Plate), then the plan and the guilds (rows that want the width); rail: the connections.
    main = characters_section(a) + plan_section(a) + guilds_section(a, asking=confirm)
    rail = devices_section(st['dev'], revoking=confirm) + discord_section(a, unlinking=confirm) + you_section(a)
    body = head + Page('split', main, rail)
    key = st.get('acct', state)
    note = f'/account mock, {state}. Data: premium-facts.json (account models on the {key} account, DERIVED from home fixtures; devices {st["dev"]}).'
    return render(f'{ACC["title"]} · Forever Sixty', body, '', who_of(key) if key != 'new' else None, chrome_of(key), note)


def account_states() -> str:
    cells = []
    for k, lab in [('free', 'Free'), ('monthly', 'Own Premium, monthly, founder price'), ('premium', 'Own Premium, yearly, founder price'), ('ending', 'Own Premium, cancelled: ends'),
                   ('grant', 'A grant from Forever Sixty'), ('seat', 'A seat from a guild'), ('ended', 'A seat ended (30 days)'), ('pastdue', 'The last payment failed')]:
        cells.append(f'<div class="st-cell"><span class="t-label">{lab}</span>{plan_section(A[k])}</div>')
    for o in ('accepted', 'declined'):
        cells.append(f'<div class="st-cell"><span class="t-label">Seat offer {o}</span>{plan_section(A["offer"], o)}</div>')
    cells.append(f'<div class="st-cell"><span class="t-label">Devices: pairing code shown</span>{devices_section("one", pairing=True)}</div>')
    cells.append(f'<div class="st-cell"><span class="t-label">Characters: refreshing from Battle.net</span>{characters_section(A["free"], refreshing=True)}</div>')
    cells.append(f'<div class="st-cell"><span class="t-label">Sign out asks first</span><p class="lede">{e(ACC["signOutAsk"])}</p><div class="acts">{Btn(e(ACC["signOut"]), tone="danger")}{Btn(e(DC["cancel"]))}</div></div>')
    body = '<div class="pc" style="padding-top:24px"><h1 class="t-title">Account: element states</h1></div>' + Page('ledger', f'<div class="st-grid">{"".join(cells)}</div>')
    return render('Account states', body, '', WHO['free'], '', 'Account states sheet; data as the account mocks.')


# ================================================================================================ /b/:id
def b_page(state: str) -> str:
    import ws_planner as WP
    if state in ('missing', 'error', 'unavailable'):
        handler = (ROOT / 'api/internal/site/handler.go').read_text()
        msgs = re.findall(r'writeMessage\(w, r, http\.Status\w+, "([^"]+)",\s*"([^"]+)"\)', handler)
        if state == 'missing':
            title, text = msgs[0]
        elif state == 'error':
            title, text = msgs[1]
        else:
            astro = (ROOT / 'web/src/pages/b-unavailable.astro').read_text()
            title = re.search(r'<h1[^>]*>([^<]+)</h1>', astro).group(1)
            text = ' '.join(re.sub(r'\s+', ' ', re.sub(r'<[^>]+>', '', p)).strip() for p in re.findall(r'<p[^>]*>([\s\S]*?)</p>', astro))
        planner = re.search(r'class="[^"]*">([^<]+)</a></p>', (ROOT / 'api/internal/site/page.go').read_text()).group(1)
        ask = f'<p class="t-small">{e(PROPOSED["askSender"])}</p>' if state == 'missing' else ''
        body = PageHeader(e(title), meta=[e(text)]) + PageForm(f'<div class="stack">{ask}<div class="acts">{LinkBtn(e(planner), primary=True)}{r(PROPOSED["guideLink"])}</div></div>')
        return render(f'{title} · Forever Sixty', body, '', None, '', f'/b/:id mock, {state}: the API message page (api/internal/site/handler.go, page.go messagePage).')
    cfg = {'shared': ('guide', 'Survival raid build', None, 41), 'you': ('guide', 'Survival raid build', 'others', 41),
           'holy': ('holy', 'Holy raid healing', 'other_class', 7), 'fire': ('fire', '', None, 12)}[state]
    key, title, send_key, views = cfg
    st = WP.S[key]
    heading = f'{st["race"]} {st["className"]}'
    h1 = title or heading
    score = WP.SAMPLE_SCORE.get(key)
    split = '/'.join(str(x) for x in st['split'])
    SH = F['shared']
    # Ruling Q10: `Shared by {name}` only when the share record names the sharer; the record carries no sharer today.
    meta = [e(SH['meta']), e(st['meta']), e(split), e(SH['views'][str(views)])]
    head = PageHeader(e(h1), CLASS_COLOR[st['cls']], crest_img(st['cls'], 56), meta, WP.answer(st, score))
    send = WP.F['send'][send_key] if send_key else None
    fork = f'{Btn(e(SH["fork"]))}'
    note = f'<div class="pc"><div data-region="{REGION["notes"]}"><p class="notice neutral">{e(SH["note"])} {fork}</p></div></div>'
    foot = WP.tree_foot({**st, 'spent': 0})
    plate = WP.TalentEditor(WP.trees(st, readonly=True), max(range(3), key=lambda i: st['split'][i]), foot)
    main = WP.Ordered(plate, 1, span=True) + WP.Ordered(WP.order(st), 6, span=True)
    # Ruling MUST 5: another class's build says so in one plain line above Copy for the game, inside the rail.
    rail = WP.Ordered(WP.in_game(st, send, plain_block=True), 2) + WP.Ordered(WP.compare(st, readonly=True, heading=SH['guideBuild']), 3)
    who = WHO['free'] if send_key else None
    body = head + note + Page('workbench', main, rail) + WP.PAGE_CSS
    return render(f'{h1} · {split} · Forever Sixty', body, 'Tools', who, '', f'/b/:id mock ({state}): the API page mounting the planner read-only. Data: planner-facts.json ({key}); SAMPLE live score.')


PAGES = {('web-premium' if s == 'out' else f'web-premium-{s}'): (lambda s=s: premium_page(s)) for s in PREMIUM_STATES}
PAGES |= {f'web-checkout-{s}': (lambda s=s: checkout_page(s)) for s in ('redirecting', 'processing', 'confirmed', 'slow', 'error', 'notopen', 'owned', 'guild', 'trial', 'invalid')}
PAGES |= {('web-account' if s == 'free' else f'web-account-{s}'): (lambda s=s: account_page(s)) for s in ['free', 'out', 'loading', 'error', *[k for k in ACCOUNT_STATES if k != 'free']]}
PAGES['web-account-states'] = account_states
PAGES |= {('web-b' if s == 'shared' else f'web-b-{s}'): (lambda s=s: b_page(s)) for s in ('shared', 'you', 'holy', 'fire', 'missing', 'error', 'unavailable')}



def compare_sheets(before: Path) -> None:
    import ws_compose
    from ws_setup_compare import pal
    ws_compose.BEFORE = before
    review = ROOT / 'design/reviews/2026-11-08-premium-account'
    pairs = [('premium', 'premium-out', 'web-premium', (390, 1440, 1920, 2000)), ('premium-own', 'premium-own', 'web-premium-own', (390, 1440)),
             ('checkout-success', 'checkout-success', 'web-checkout-confirmed', (390, 1440)), ('checkout-guild', 'checkout-guild', 'web-checkout-guild', (390, 1440)),
             ('account', 'account-34', 'web-account', (390, 1440, 1920, 2000)), ('account-premium', 'account-premium', 'web-account-premium', (390, 1440, 2000)),
             ('account-new', 'account-new', 'web-account-new', (390, 1440)), ('b', 'b-shared', 'web-b', (390, 1440, 1920, 2000)), ('b-missing', 'b-missing', 'web-b-missing', (390, 1440))]
    scale = {390: .8, 1440: .5, 1920: .42, 2000: .4}
    for name, live, mock, widths in pairs:
        for w in widths:
            out = ws_compose.compose(name, live, mock, w, scale[w], max_h=5000)
            pal(out, review / 'after' / out.name, 3000)
            print(out)


if __name__ == '__main__':
    if sys.argv[1:2] == ['--list']:
        print(' '.join(PAGES))
    elif sys.argv[1:2] == ['compare']:
        compare_sheets(Path(sys.argv[2]))
    else:
        only = sys.argv[1:]
        bad = [n for n in only if n not in PAGES]
        if bad:
            sys.exit(f'ws_premium.py: no such mock: {" ".join(bad)} (see --list)')
        for name, fn in PAGES.items():
            if not only or name in only:
                print(write(f'{name}.html', fn()))
