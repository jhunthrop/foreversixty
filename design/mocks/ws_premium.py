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
from ws_prims import (Button as _Button, Columns, Emblem, EmptyFill, Facts, Figure, Link, Notice, Page, PageForm, PageHeader, Plate, PlateColumns,
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
.stack{display:flex;flex-direction:column;gap:12px;min-width:0}.stack.tight{gap:6px}
.acts{display:flex;gap:10px 16px;flex-wrap:wrap;align-items:center}
.lede{font-size:15px;line-height:24px;max-width:72ch}.t-small{max-width:80ch}
.fd{display:inline-flex;gap:8px;align-items:center;flex-wrap:wrap}
.pcol .rprose.ruled li,.plan .rprose.ruled li{padding-block:8px}.pcol .rprose.lane,.plan .rprose.lane{font-size:15px;line-height:22px}
.pcol .btn{align-self:flex-start}
@media(max-width:767px){.pcol .btn{align-self:stretch}}
.ans-who{display:flex;align-items:center;gap:12px}.ans-who b{display:block;font:700 18px/22px var(--font-body)}.ans-links{display:flex;gap:16px;flex-wrap:wrap;font-size:16px}  /* the links take the page head's 16 px */
@media(min-width:900px){.head .ans:has(>.ans-plan),.head .ans:has(>.ans-who){min-width:300px}.ans-who{justify-content:flex-end;text-align:right}.ans-links{justify-content:flex-end}}
.ans-plan{display:flex;flex-direction:column;gap:2px}.ans-line{display:flex;flex-wrap:wrap;align-items:center;gap:4px 10px}
@media(min-width:900px){.ans-plan{align-items:flex-end}.ans-line{justify-content:flex-end}}
.coach .opening{font-size:15px;line-height:23px}
.coach .who{display:flex;gap:12px;align-items:center}.coach .who b{font:700 16px/20px var(--font-body);color:var(--strong)}
.alert{color:var(--strong)}.wipe-t{color:var(--wipe)}
.rowsel{display:flex;align-items:center;gap:8px}
.swt{display:flex;align-items:center;justify-content:space-between;gap:12px;min-height:44px}
.planfig{display:flex;align-items:flex-end;justify-content:space-between;gap:12px;flex-wrap:wrap}
.dev .row .sub{white-space:normal}
.msg{display:flex;flex-direction:column;gap:12px}
@media(max-width:1023px){.btn.sm{height:44px}}
@media(max-width:767px){.cbar .pc{padding-block:0;min-height:0}}   /* the phone bar is its 48 px line, as the built ControlBar */
.ck-on{display:flex;flex-direction:column;gap:4px;padding:20px 24px 12px}.ck-on .m0{margin:0}
.ck-on-t{margin:0;font:400 24px/30px var(--font-title);letter-spacing:-.015em;color:var(--strong)}.ck-next{display:flex;flex-direction:column;padding:0 8px 12px}
.statesheet{display:grid;gap:40px}.st-grid{display:grid;gap:24px 40px;grid-template-columns:repeat(auto-fill,minmax(340px,1fr))}.st-cell{display:flex;flex-direction:column;gap:6px;min-width:0}
</style>'''


# The account page's Sections open on a 12 px gap under the heading's hairline, as the built ui/Section (pt-3);
# /premium's and /b/:id's Sections open on a Plate, Columns or the planner's own blocks, which carry their own.
ACCOUNT_CSS = ('<style>section>.sec-h+*{margin-top:12px}.acc-foot{padding:0 20px 10px}'
               '@media(min-width:1024px){.plan-rows .row>.trail{grid-column:1/-1;justify-self:start}}'
               '@media(max-width:767px){.acc-foot{padding:0 14px 8px}}'
               # The Characters panel's inset around its EmptyFill (CharacterList.svelte px-5 pt-4 pb-5) and the Show all slot
               # under its rows (px-5 py-2, drawn even when the list already fits).
               '.es-fill-wrap{padding:16px 20px 20px}</style>')


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


# The four plans side by side (directions mock b4-premium; web pages/premium.astro): each a Plate holding one plan, Premium
# framed with the teal Subscribe. One column on a phone with Premium first and every buy button under its price; two from 768
# and four from 1280, every buy button closing its panel. Questions and the guild rules beside the coach sample from 1024.
PLANS_CSS = '''<style>
.pgrid{display:grid;grid-template-columns:minmax(0,1fr);gap:18px}.pgrid>:nth-child(2){order:-1}
@media(min-width:768px){.pgrid{grid-template-columns:repeat(2,minmax(0,1fr))}.pgrid>:nth-child(2){order:0}}
@media(min-width:1280px){.pgrid{grid-template-columns:repeat(4,minmax(0,1fr))}}
.plan{display:flex;flex-direction:column;gap:12px;height:100%;min-width:0;padding:22px 24px;box-sizing:border-box}
.plan>.plh{order:0}.plan>.pact{display:flex;flex-direction:column;align-items:stretch;gap:10px;order:1}
.plan>.rprose,.plan>.psm,.plan>.pfoot{order:2}.plan>.pfoot{margin-top:auto}
@media(min-width:768px){.plan>.pact{order:3;margin-top:auto;padding-top:4px}}
.plan .pact button.btn{width:100%}.plan .pact .btn{display:flex;width:100%;height:auto;min-height:44px;padding-block:10px;white-space:normal;text-align:center;text-wrap:balance;line-height:22px;box-sizing:border-box}
.pmore{display:grid;grid-template-columns:minmax(0,1fr);gap:22px}.pmore>.pmain{display:flex;flex-direction:column;gap:22px;min-width:0}
@media(min-width:1024px){.pmore{grid-template-columns:minmax(0,1fr) 340px;column-gap:28px;align-items:start}}
@media(min-width:1760px){.pmore{grid-template-columns:minmax(0,1fr) var(--rail-w)}}
.pc.plans{padding-top:28px}.pc.plans>.pstack{display:flex;flex-direction:column;gap:22px}
.prules{padding-top:20px}.prules-t{margin:0;padding-bottom:4px;font:400 20px/26px var(--font-title);letter-spacing:-.015em;color:var(--cream)}@media(max-width:1023px){.pmore>:last-child{order:-1}}.pmain .rgrid>.row{padding-inline:0}.phist{padding-top:12px}.pmain p.t-small.phist{max-width:none}.pmem{padding-bottom:4px}.pmain .rgrid>.row{padding-inline:0}
@media(max-width:767px){.plan{padding:18px 16px}}
.plan .plh .fig{margin-bottom:4px}   /* PlanPrice's figure keeps 4 px under it */
main{padding-bottom:40px}@media(min-width:1024px){main{padding-bottom:64px}}   /* Base's page body above the footer */
</style>'''


def plan_cell(inner: str, frame: str = 'plain') -> str:
    return Plate(inner, frame=frame)


def premium_column(st: dict) -> str:
    pc = price_key('premium', st['interval'], st['founder'])
    c = pc['cell']
    lines = [(e(c['secondLine']), '')] + ([(e(c['founderLine']), 'gold')] if c['founderLine'] else [])
    owned = st.get('owned')
    # Ruling PR2: founder is said once (the notice and the price line); the column carries only `Your plan`.
    tag = Tag(e(YOUR_PLAN), 'kill') if owned else ''
    head = PlanHead(e(P['premium']['name']), e(P['premium']['audience']), e(c['amount']), e(c['suffix'].strip()), lines, tag, gold=True, level=2)
    if owned:
        act = f'<p class="lede strong">{e(SRC["premiumCard"][owned])}</p>{LinkBtn(e(SRC["buttons"]["seePlan"]))}'
    else:
        act = _Button(e(pc['button']), tone='subscribe')
    feats = ruled([l['text'] for l in F['premium_lines']], [l['anchor'] if l['first'] else '' for l in F['premium_lines']])
    body = f'<div class="plan">{head}{feats}<p class="t-small psm">{e(P["premium"]["scope"])}</p><div class="pact">{act}</div></div>'
    return plan_cell(body, 'framed')


def free_column(st: dict) -> str:
    mine = st.get('signed') and not st.get('owned')
    tag = Tag(e(YOUR_PLAN), 'kill') if mine else ''
    head = PlanHead(e(P['free']['name']), e(P['free']['audience']), e(P['free']['price']), '', [(e(P['free']['note']), '')], tag, level=2)
    foot = '' if st.get('signed') else f'<span class="pfoot">{r(F["my_reports"]["signInAction"])}</span>'
    return plan_cell(f'<div class="plan">{head}{ruled(P["free"]["features"])}{foot}</div>')


def guild_column(plan: str, st: dict) -> str:
    block = P['guild'] if plan == 'guild' else P['raidGuild']
    pc = price_key(plan, st['interval'], st['founder'])
    c = pc['cell']
    lines = [(e(c['secondLine']), '')] + ([(e(c['founderLine']), 'gold')] if c['founderLine'] else []) + [(e(pc['cost']), '')]
    on = st.get('guild_on') == plan
    tag = Tag('Olympus XXVII', 'kill') if on else ''
    head = PlanHead(e(block['name']), e(block['audience']), e(c['amount']), e(c['suffix'].strip()), lines, tag, level=2)
    if on:
        act = f'<p class="lede strong">{e(F["guild_on_line"])}</p>{Btn(e(GC["onPlan"]["manageBilling"]))}'
    else:
        act = LinkBtn(e(pc['button']))
    trial = f'<p class="t-small psm"><b class="strong">{e(F["trial"]["lead"])}</b>{e(F["trial"]["rest"])}</p>'
    return plan_cell(f'<div class="plan">{head}{ruled(block["features"])}{trial}<div class="pact">{act}</div></div>')


def coach_rail() -> str:
    c = F['coach']
    cp = P['coachSample']
    fix = c['fixFirst']
    then = ''.join(Row(e(t['title']), e(t['why']), lead=icon(t['icon'], 32) if t['icon'] else '', link=False) for t in c['then'])
    well = ''.join(Row(e(w['text']), lead=icon(w['icon'], 24) if w['icon'] else '', link=False, slim=True) for w in c['doingWell'])
    body = (f'<div class="stack coach" id="coach"><div class="who">{crest_img(c["classSlug"], 40)}<span><b>{e(c["title"])}</b><span class="t-small" style="display:block">{e(c["context"])}</span></span></div>'
            f'<p class="opening">{e(c["opening"])}</p><span class="t-label">{e(cp["fixFirst"])}</span>'
            + Rows([Row(e(fix['title']), e(fix['why']), lead=icon(fix['icon'], 40), link=False)], True)
            + f'<span class="t-label">{e(cp["then"])}</span>{Rows([then], True)}<span class="t-label">{e(cp["doingWell"])}</span>{Rows([well], True)}'
            f'<p class="t-small">{e(cp["caption"])}</p></div>')
    return Section(e(cp['label']), body, action=Tag(e(cp['tag']), 'muted'), panel=True)


def premium_answer(key: str) -> str:
    """Signed in: the plan the account is on now and the way to it (tenet 14: the player first)."""
    if key in (None, 'out'):
        return ''
    a = A[key]
    h = a['plan']['head']
    when = f'{h["when"]["label"]} {h["when"]["date"]}' if h['when'] else ''
    sub = ' · '.join(x for x in [when, h['extra'] or ''] if x)
    tag = Tag(e(AP['founder']['pill'])) if a['plan']['founder'] else ''
    # Compact (batch F): the figure, then one line with the date or source, the Founder mark and the link.
    line = (f'<span>{e(sub)}</span>' if sub else '') + tag + r(SRC['buttons']['seePlan'])
    return f'<div class="ans-plan">{Figure(e(h["name"]), e(ACC["planLabel"]), size="sm")}<span class="t-small ans-line">{line}</span></div>'


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
    head = PageHeader(e(P['title']), answer=premium_answer(key))
    iv = seg([(e(P['interval']['labels']['yearly']), st['interval'] == 'yearly', None, ''), (e(P['interval']['labels']['monthly']), st['interval'] == 'monthly', None, '')], label=e(P['interval']['groupLabel']))
    bar = control_bar([(e(P['interval']['groupLabel']), iv)], f'{e(P["interval"]["labels"][st["interval"]])}', '', label=e(P['interval']['groupLabel']))
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
    # The region is the notes in the page column, as the planner's (the built Notice sits inside the column's inset):
    # one .pc around the region, the Notices' own .pc wrappers dropped.
    inner = notices.replace('<div class="pc">', '').replace('</p></div>', '</p>')
    notes = f'<div class="pc"><div data-region="{REGION["notes"]}">{inner}</div></div>' if notices else ''
    grid = (f'<div class="pgrid" role="group" aria-label="{e(P["title"])}">'
            + free_column(st) + premium_column(st) + guild_column('guild', st) + guild_column('raid_guild', st) + '</div>')
    # The guild rules under the questions on one panel (the redesign review: one prose panel on the page).
    rules = (f'<div class="prules" id="guild-leaders"><h3 class="prules-t">{e(P["howItWorks"])}</h3>'
             + RowGrid([Row(e(x['term']), e(x['sentence']), link=False) for x in P['rules']], 2)
             + f'<p class="t-small phist">{e(P["historyPromise"])}</p></div>')
    faq = Section(e(FAQ_TITLE), RowGrid([Row(e(x['question']), e(x['answer']), link=False) for x in P['faq']], 2) + rules,
                  meta=e(P['foot']), panel=True)
    more = f'<div class="pmore"><div class="pmain">{faq}</div>{coach_rail()}</div>'
    body = head + bar + notes + f'<div class="pc plans"><div class="pstack">{grid}{more}</div></div>' + PLANS_CSS
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
    return Plate(Steps(rows), 'Checkout', frame='framed')


def checkout_page(state: str) -> str:
    who, chrome = WHO['premium'], ''
    after = ''
    if state in ('redirecting', 'processing', 'confirmed', 'slow', 'error', 'notopen'):
        head = PageHeader(e(F['plan_table']['premium']['name']), answer=band('premium', change=state in ('redirecting', 'error', 'notopen')))
        plate = checkout_steps(state)
        if state == 'confirmed':
            # The redesign review (2026-11-09): no finished checklist; Premium is on, the renewal line, and what it turned
            # on as rows that go there (web checkout/CheckoutConfirmed.svelte).
            coach, sims = F['premium_lines'][0]['text'], F['premium_lines'][1]['text']
            nxt = Row(e(CK['nextCoach']), e(coach)) + Row(e(CK['nextSim']), e(sims)) + Row(e(SRC['buttons']['seePlan']), e(CK['seePlanSub']))
            plate = Plate(f'<div class="ck-on"><p class="ck-on-t" role="status">{e(F["checkout_plan_on"])}</p>'
                          f'<p class="t-small m0">{e(F["checkout_renewal"])}</p></div>'
                          f'<nav class="ck-next" aria-label="{e(CK["nextLabel"])}">{nxt}</nav>', 'Checkout', frame='framed')
        body = head + PageForm(plate + after)
    elif state == 'trial':
        # Ruling K1: a guild checkout in its free trial took no payment; row 2 says the card was saved.
        head = PageHeader(e(GC['band']['title']), answer=band('guild', change=False))
        rows = [('ok', 1, e(STEPS['signIn']), f'<span class="tx">{e(F["checkout_signed_in"])}</span>', ''),
                ('ok', 2, e(STEPS['pay']), e(F['checkout_trial_saved']), ''),
                ('ok', 3, e(F['plan_table']['guild']['name'] + ' on'), f'<span class="tx" role="status">{e(CK["confirmed"])}</span>', '')]
        body = head + PageForm(Plate(Steps(rows), 'Checkout', frame='framed'))
    elif state == 'owned':
        o = A['seat']['owned']
        who = WHO['free']
        # The account already has Premium: the head keeps the answer's box and draws nothing in it (CheckoutHead `owned`).
        head = PageHeader(e(F['plan_table']['premium']['name']), answer=band('premium')) + '<style>[data-region="head-answer"]{visibility:hidden}</style>'
        inner = (f'<div class="stack" style="padding:16px 20px 20px"><h2 class="t-h3">{e(o["title"])}</h2><p class="lede strong">{e(o["line"])}</p>'
                 + (f'<p class="t-small">{e(o["line2"])}</p>' if o['line2'] else '') + f'<div class="acts">{LinkBtn(e(o["action"]["label"]))}</div></div>')
        body = head + PageForm(Plate(inner, 'Your plan', frame='framed'))
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
def empty_fill_inset(lead: str, fill: str, action: str, secondary: str) -> str:
    """ws_prims.EmptyFill (ui/EmptyState `fill`) inside a flush panel (CharacterList.svelte): the panel's 20 px inset around
    the strong lead, the muted fill, the small primary button and the quieter gold link."""
    return f'<div class="es-fill-wrap">{EmptyFill(lead, fill, action, Link(f"{secondary} &rsaquo;"))}</div>'


def characters_section(a: dict, refreshing: bool = False) -> str:
    rows = []
    for c in a['characters']:
        tags = (Tag(e(CL['current'])) if c['main'] else '') + (Tag(e(CL['main']), 'muted') if c['main'] else '')
        src = Tag(e(CL['addonSource'] if c['source'] == 'addon' else (CL['battlenetSource'] if c['source'] else CL['noBuildYet'])), 'muted')
        setmain = '' if c['main'] else Btn(e(CL['setAsMain']), small=True)
        sub = e(c['descriptor']) + (f'<br>{e(c["guild"])}' if c['guild'] else '')
        ilvl = f'<span class="strong num">{e(c["itemLevel"])}</span>' if c.get('itemLevel') else ''
        rows.append(Row(f'<span class="fd"><span style="color:{CLASS_COLOR.get(c["cls"], "#f2eee4")}">{e(c["name"])}</span>{tags}</span>', sub, lead=crest_img(c['cls'], 40),
                        trail=f'{ilvl}{src}{setmain}{r(CL['open'])}', link=False))
    refresh = f'<span class="t-small">{e(CL["refreshing"])}</span>' if refreshing else r(CL['refreshFromBattlenet'])
    if not rows:
        inner = empty_fill_inset(e(CL['emptyLead']), e(CL['emptyKorea']), LinkBtn(e(CL['getSetUp']), primary=True, small=True), e(CL['pasteAnExport']))
        return Section(e(CL['heading']), inner, action=refresh, panel=True, frame='framed', flush=True)
    foot = (f'<div class="acc-foot"><p class="t-small" style="padding-top:12px">{e(CL["introBattlenetLine"])} {Link(e(CL["introInstallAddonLink"]))} {e(CL["introAddonTail"])} '
            f'{Link(e(CL["introPasteLink"]))} {e(CL["introPasteTail"])}</p><p style="padding-top:8px">{r(ACC["reportsLink"])}</p></div>')
    return Section(e(CL['heading']), Rows(rows) + foot, action=refresh, meta=e(a['imported'] or ''), panel=True, frame='framed', flush=True)


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
    body = f'<div class="stack" id="plan">{top}{lines}{offer}<div class="plan-rows">{RowsSlim(rows)}</div><div class="acts">{"".join(acts)}</div></div>'
    return Section(e(ACC['planLabel']), body, panel=True)


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
    return Section(e(ACC['guildLabel']), f'<div class="stack">{"".join(blocks)}</div>', panel=True)


def devices_section(dev: str, revoking: bool = False, pairing: bool = False) -> str:
    d = F['devices'][dev]
    rows = []
    for i, (name, line) in enumerate(zip(d['names'], d['lines'])):
        sub = line.split(' · ', 1)[1]
        if revoking and i == 0:
            rows.append(Row(e(d['revoke'][0]), trail=f'{Btn(e(DC["revoke"]), small=True, tone="danger")}{Btn(e(DC["cancel"]), small=True)}', link=False))
        else:
            rows.append(Row(e(name), e(sub), trail=Btn(e(DC['revoke']), small=True), link=False))
    entry = f'<button type="button" class="link">{e(d["entry"])} &rsaquo;</button>'
    if pairing:
        tail = f'<p class="t-small">{e(DC["codeFallbackLine"])}</p><code class="c big">4821-9930</code>'
    elif not rows:
        tail = f'<p class="t-small">{e(ACC["noDevices"])}</p><div class="acts">{LinkBtn(e(ACC["downloadCompanion"]))}{entry}</div>'
    else:
        tail = f'<div class="acts">{entry}</div>'
    return Section(e(ACC['devicesLabel']), f'<div class="stack dev">{Rows(rows, True) if rows else ""}{tail}</div>', panel=True)


def discord_section(a: dict, unlinking: bool = False) -> str:
    link = a['discord']['link']
    if link and unlinking:
        top = Row(e(DL['unlinkQuestion']), trail=f'{Btn(e(DL["unlink"]), small=True, tone="danger")}{Btn(e(DL["cancel"]), small=True)}', link=False)
    elif link:
        top = Row(f'{e(DL["linkedLead"])} <b class="strong">{e(link["display_name"])}</b>', trail=Btn(e(DL['unlink']), small=True), link=False)
    else:
        top = Row(e(DL['lead']), trail=LinkBtn(e(DL['link']), small=True), link=False)
    # Ruling A4 / MUST 3: an account that already has Premium is told the role turns on as soon as it links.
    lead = DR['leadOwned'] if a['source'] != 'none' else DR['lead']
    role = Row(e(DR['label']), e(lead), trail=LinkBtn(e(DR['link']), small=True), link=False)
    st = a['discord']['state']
    if st == 'saved':
        nt = (f'<div class="stack tight"><span class="t-label">{e(NC["headLabel"])}</span><p class="lede acts"><span>{e(NC["destinationChannel"])}</span>{r(NC["openInDiscord"])}</p>'
              f'{Switch(e(NC["eventSimFinished"]))}{Switch(e(NC["eventUploadParsed"]))}'
              f'{Switch(e(NC["eventUploadPrivate"]) + f"<span class=t-small style=display:block>{e(NC["eventsNote"])}</span>", on=False)}'
              f'<div class="acts">{Btn(e(NC["test"]), small=True)}{Btn(e(NC["remove"]), small=True, tone="danger")}</div>'
              f'<p class="acts">{r(NC["reconnect"])}<span class="t-small">{e(NC["reconnectAside"])}</span></p></div>')
    else:
        nt = f'<div class="stack tight"><span class="t-label">{e(NC["headLabel"])}</span><p class="t-small">{e(NC["freeSentence"])}</p><p>{r(NC["seePlans"])}</p></div>'
    return Section(e(NC['label']), f'<div class="stack" id="discord">{Rows([top, role], True)}{nt}</div>', panel=True)


def you_section(a: dict) -> str:
    """The rail's first panel (directions mock b4-account): who is signed in and how, Sign out, then the pseudonym."""
    who = Facts([(e(ACC['battletagLabel']), e(a['battletag'])), (e(ACC['signInLabel']), f'<span class="muted">{e(ACC["signInBattlenet"])}</span>')])
    sw = Switch(e(ACC['pseudonymLabel']), on=False)
    return Section(e(ACC['youLabel']), f'<div class="stack">{who}<div class="acts">{Btn(e(ACC["signOut"]), small=True)}</div>{sw}'
                   f'<p class="t-small">{e(ACC["pseudonymNote"])}</p></div>', panel=True)


def delete_section() -> str:
    """The rail's last Section (gap G6 closed 2026-11-07): one line and the danger Button that opens the confirm dialog."""
    ds = ACC['deleteSection']
    return Section(e(ds['title']), f'<div class="stack"><p class="t-small">{e(ds["line"])}</p><div class="acts">{Btn(e(ds["button"]), small=True, tone="danger")}</div></div>', panel=True)


def dialog_line(line: dict) -> str:
    """One dialog line, its linked words (the refund policy) as a link, as delete-confirm.ts lineParts splits it."""
    link = line.get('link')
    if not link or link['label'] not in line['text']:
        return f'<li>{e(line["text"])}</li>'
    before, after = line['text'].split(link['label'], 1)
    return f'<li>{e(before)}<a class="link" href="#">{e(link["label"])}</a>{e(after)}</li>'


def delete_dialog(a: dict) -> str:
    """The confirm dialog over the page: what is deleted, what is kept, the typed confirm (empty: the button is off)."""
    dd = ACC['deleteDialog']
    word = a['deleteWord']
    label = e(a['deleteConfirm']).replace(e(word), f'<b class="strong">{e(word)}</b>', 1)
    body = (f'<h3 class="wipe">{e(dd["deletedLabel"])}</h3><ul>{"".join(dialog_line(x) for x in dd["deleted"])}</ul>'
            f'<h3>{e(dd["keptLabel"])}</h3><ul>{"".join(dialog_line(x) for x in dd["kept"])}</ul>'
            f'<div class="dlg-f"><label for="account-delete-confirm">{label}</label>'
            f'<input id="account-delete-confirm" type="text" autocomplete="off" spellcheck="false" autocapitalize="off"></div>')
    actions = Btn(e(dd['confirm']), tone='danger', off=True) + Btn(e(dd['cancel']))
    return ws_prims.Dialog(e(dd['title']), e(dd['intro']), body, actions, 'account-delete-title')


ACCOUNT_STATES = {
    'free': dict(dev='none'), 'several': dict(dev='several'), 'premium': dict(dev='one'), 'seat': dict(dev='one'),
    'offer': dict(dev='one'), 'ended': dict(dev='one'), 'pastdue': dict(dev='one'), 'claim': dict(dev='one'), 'new': dict(dev='none'),
    'confirm': dict(dev='one', acct='premium', confirm=True), 'grant': dict(dev='one'),
    'delete': dict(dev='none', acct='free', dialog=True),
}


def account_page(state: str) -> str:
    if state == 'out':
        head = PageHeader(e(ACC['title']))
        inner = f'<div class="stack" style="padding:16px 20px 20px"><p class="lede">{e(F["signin"]["reason"])}</p><div class="acts">{LinkBtn(e(F['my_reports']['signInAction']), primary=True)}</div></div>'
        body = head + PageForm(Plate(inner, 'Sign in'))
        return render(f'{ACC["title"]} · Forever Sixty', body, '', None, '', '/account mock, signed out. Copy: signin-copy.ts.', extra_css=ACCOUNT_CSS)
    if state in ('loading', 'error'):
        head = PageHeader(e(ACC['title']))
        if state == 'loading':
            sk = lambda n: ''.join(Skeleton(560, 14) + '<br>' for _ in range(n))
            plate = Section(e(CL['heading']), f'<div class="stack" role="status" aria-busy="true">{"".join(Skeleton(480, 56) for _ in range(2))}</div>', panel=True, frame='framed')
            sk = f'<div class="stack">{Skeleton(300, 20)}{Skeleton(240, 20)}{Skeleton(160, 44)}</div>'
            plate += ''.join(Section(t, sk, panel=True) for t in (ACC['devicesLabel'], NC['label']))
            rail = ''.join(Section(t, sk, panel=True) for t in (ACC['youLabel'], ACC['planLabel'], ACC['deleteSection']['title']))
            body = head + Page('split', plate, rail, panels=True)
        else:
            inner = f'<div class="stack"><p class="alert" role="alert">{e(F["account_failed"])}</p><div class="acts">{Btn(e(CK['retry']), primary=True)}</div></div>'
            body = head + Page('split', Section(e(CL['heading']), inner, panel=True, frame='framed'), '', panels=True)
        return render(f'{ACC["title"]} · Forever Sixty', body, '', None, 'pending', f'/account mock, {state}.', extra_css=ACCOUNT_CSS)
    st = ACCOUNT_STATES[state]
    a = A[st.get('acct', state)]
    confirm = st.get('confirm', False)
    head = PageHeader(e(ACC['title']))
    # Main: the characters (the framed panel), the guilds, then the connections; rail: who you are, the plan, Delete
    # account last (directions mock b4-account).
    main = characters_section(a) + guilds_section(a, asking=confirm) + devices_section(st['dev'], revoking=confirm) + discord_section(a, unlinking=confirm)
    rail = you_section(a) + plan_section(a) + delete_section()
    body = head + Page('split', main, rail, panels=True) + (delete_dialog(a) if st.get('dialog') else '')
    key = st.get('acct', state)
    note = f'/account mock, {state}. Data: premium-facts.json (account models on the {key} account, DERIVED from home fixtures; devices {st["dev"]}).'
    return render(f'{ACC["title"]} · Forever Sixty', body, '', who_of(key) if key != 'new' else None, chrome_of(key), note, extra_css=ACCOUNT_CSS)


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
    return render('Account states', body, '', WHO['free'], '', 'Account states sheet; data as the account mocks.', extra_css=ACCOUNT_CSS)


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
            # The page's frontmatter holds its two strings (`const title`, `const sentence`, the
            # sentence possibly split over '...' + '...' parts).
            title = re.search(r"const title = '([^']+)'", astro).group(1)
            parts = re.search(r"const sentence =\s*((?:'[^']*'\s*\+?\s*)+);", astro).group(1)
            text = ''.join(re.findall(r"'([^']*)'", parts))
        planner = re.search(r'class="\{\{\.ButtonClass\}\}">([^<]+)</a>', (ROOT / 'api/internal/site/page.go').read_text()).group(1)
        # A missing build closes its sentence with the ask (ruling Q12); a failure on our side does not.
        ask = f' {e(PROPOSED["askSender"])}' if state == 'missing' else ''
        # Round 5 (directions mock b4-unavailable): the title band, then the sentence and the two ways on one framed
        # panel. The API's missing and error pages (api/internal/site/page.go messagePage) take the same shape.
        panel = (f'<div class="plate framed" data-region="{REGION["plate"]}" style="padding:22px 24px"><p style="margin:0">{e(text)}{ask}</p>'
                 f'<div class="acts" style="margin-top:16px">{LinkBtn(e(planner), primary=True)}{r(PROPOSED["guideLink"])}</div></div>')
        if state == 'unavailable':
            # The static page fills the frame (owner 2026-11-10, one width on every page): the panel the
            # frame's width, then the 404's Pick your class (ws_pages.class_picker, NotFoundClassPicker).
            import ws_pages as WPG
            # PAGE_CSS carries the 404's own main padding (its <main> has pb-(--section-gap)); this page's main has none.
            from ws_lib import main_foot_padding
            body = (PageHeader(e(title)) + Page('ledger', panel + f'<div class="section">{WPG.class_picker()}</div>') + WPG.PAGE_CSS
                    + f'<style>{main_foot_padding()}</style>')
        else:
            body = PageHeader(e(title)) + PageForm(panel)
        return render(f'{title} · Forever Sixty', body, '', None, '', f'/b/:id mock, {state}: the API message page (api/internal/site/handler.go, page.go messagePage).')
    cfg = {'shared': ('guide', 'Survival raid build', None, 41), 'you': ('guide', 'Survival raid build', 'others', 41),
           'holy': ('holy', 'Holy raid healing', 'other_class', 7), 'fire': ('fire', '', None, 12)}[state]
    key, title, send_key, views = cfg
    st = WP.S[key]
    heading = f'{st["race"]} {st["className"]}'
    h1 = title or heading
    split = '/'.join(str(x) for x in st['split'])
    SH = F['shared']
    # Ruling Q10: `Shared by {name}` only when the share record names the sharer; the record carries no sharer today.
    meta = [e(SH['meta']), e(st['meta']), e(split), e(SH['views'][str(views)])]
    head = PageHeader(e(h1), CLASS_COLOR[st['cls']], crest_img(st['cls'], 56), meta, WP.answer(st))
    send = WP.F['send'][send_key] if send_key else None
    fork = f'{Btn(e(SH["fork"]))}'
    # The shared-link line at the reading measure, as the built planner sets its note line.
    note = f'<div class="pc" style="padding-top:16px"><div data-region="{REGION["notes"]}" style="max-width:var(--measure-reading)"><p class="notice neutral">{e(SH["note"])} {fork}</p></div></div>'
    foot = WP.tree_foot({**st, 'spent': 0})
    plate = WP.TalentEditor(WP.trees(st, readonly=True), max(range(3), key=lambda i: st['split'][i]), foot)
    main = WP.Ordered(plate, 1, span=True)
    # Point order full width under the trees and the rail (the planner's `below`, h-audit 2026-11-09); on a phone after the trees.
    below = WP.Ordered(WP.order(st), 1, span=True)
    # Ruling MUST 5: another class's build says so in one plain line above Copy for the game, inside the rail. The planner's
    # own panels (lane P7): WP.Page sets the workbench as panels, WP.TalentEditor is the framed Talents panel.
    # The page's own rail order: the guide build first in the markup, In the game placed first by its order (Planner.svelte).
    rail = WP.Ordered(WP.compare(st, readonly=True, heading=SH['guideBuild']), 3) + WP.Ordered(WP.in_game(st, send, plain_block=True), 2)
    who = WHO['free'] if send_key else None
    # The planner's head answer is as wide as its points and link (PlannerAnswer), not the account answer's 300 px; the API
    # page's main holds no bottom pad (the planner page's own does).
    body = head + note + WP.Page('workbench', main, rail, below=below) + WP.PAGE_CSS + '<style>@media(min-width:900px){.head .ans[data-region]{min-width:0}}.pc:has(>.layout-wb){padding-bottom:0}</style>'
    return render(f'{h1} · {split} · Forever Sixty', body, 'Tools', who, '', f'/b/:id mock ({state}): the API page mounting the planner read-only. Data: planner-facts.json ({key}).')


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
