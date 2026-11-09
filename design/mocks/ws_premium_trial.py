"""Web-layout mocks for the free trials (design/specs/2026-11-08-premium-trial.md), composed ONLY from ws_prims and the
premium mocks' own blocks (ws_premium: the free column, the checkout band, render and the account CSS):
  - web-trial-card: /premium's Premium column with the trial line and Start 7 free days;
  - web-trial-chooser, -code, -code-unknown, -used: the checkout chooser (the step before Stripe), with no code, a valid
    creator code (30 days), an unknown code, and an account that already had its trial;
  - web-trial-confirmed: back from Stripe, the trial on, its end date, price and Cancel;
  - web-trial-account, -confirm, -canceled: the account page's plan in trial, the cancel confirm, and after cancelling;
  - web-trial-guild: guild settings' plan line in trial;
  - web-trial-email-reminder, -email-guild, -email-receipt: the plain-text emails as the inbox shows them.
Every string and date is read from design/mocks/data/premium-trial-facts.json (node design/mocks/premium_trial_facts.mjs:
trial-copy.ts on plans.ts at 2026-11-08T10:00Z); the premium page's own copy from premium-facts.json.

  node design/mocks/premium_trial_facts.mjs && python3 design/mocks/ws_premium_trial.py
FS_TRIAL_FACTS and FS_MOCK_OUT point the generator elsewhere.
"""
import json
import os
import sys
from pathlib import Path

import ws_premium as wp
from ws_premium import ACCOUNT_CSS, Btn, LinkBtn, WHO, band, e, free_column, r, render, ruled
from ws_prims import (Columns, Dialog, Facts, Field, Figure, Page, PageForm, PageHeader, Plate, PlanHead, Row, Rows, RowsSlim, Section,
                      Steps, Tag, write)
from ws_lib import seg

HERE = Path(__file__).resolve().parent
T = json.loads(Path(os.environ.get('FS_TRIAL_FACTS') or HERE / 'data' / 'premium-trial-facts.json').read_text())
C = T['copy']
P = wp.P

CSS = '''<style>
.trl{font-size:15px;line-height:23px;color:var(--strong);max-width:62ch}
.codebox{display:flex;gap:10px;align-items:flex-end;flex-wrap:wrap}.codebox .f{flex:1 1 220px}
.f{display:flex;flex-direction:column;gap:6px;min-width:0}
.ok-t{color:var(--kill)}.miss-t{color:var(--wipe)}
.mail{font:14px/22px ui-monospace,Menlo,monospace;white-space:pre-wrap;padding:16px 20px 20px;color:var(--text)}
.mailhead{padding:16px 20px 0}
section>.sec-h+*{margin-top:12px}
</style>'''


def page(title: str, body: str, note: str, who=WHO['premium']) -> str:
    return render(f'{title} · Forever Sixty', body, 'Premium', who, '', note, CSS + ACCOUNT_CSS)


def code_input(value: str = '') -> str:
    v = f' value="{e(value)}"' if value else ''
    return f'<input class="field" type="text" placeholder="{e(C["chooser"]["codePlaceholder"])}"{v}>'


# ================================================================================================ /premium's Premium column
def trial_column() -> tuple:
    pc = wp.price_key('premium', 'monthly', T['founder'])
    c = pc['cell']
    head = PlanHead(e(P['premium']['name']), e(P['premium']['audience']), e(c['amount']), e(c['suffix'].strip()),
                    [(e(c['secondLine']), '')], Tag(e(C['cardTag'])), gold=True, level=2)
    body = (f'<div class="stack">{head}<p class="trl">{e(T["lines"]["monthly"])}</p>{LinkBtn(e(T["chooser"]["start"]), primary=True)}'
            f'<p class="t-small">{e(P["premium"]["support"])}</p>'
            + ruled([l['text'] for l in wp.F['premium_lines']]) + '</div>')
    return ('', '', body, '')


def card() -> str:
    head = PageHeader(e(P['title']), lead=e(P['band']['lede']))
    plate = Plate(Columns([free_column({'interval': 'monthly', 'founder': True}), trial_column()], open_=False, second_first=True), 'Plans for a player')
    return page(P['title'], head + Page('split', plate, ''), '/premium mock: the Premium column with the 7-day trial line (trial-copy.ts trialLine).', who=None)


# ================================================================================================ the checkout chooser
def chooser(state: str) -> str:
    ch = C['chooser']
    interval = seg([(e(P['interval']['labels']['yearly']), True, None, ''), (e(P['interval']['labels']['monthly']), False, None, '')],
                   label=e(ch['intervalLabel']))
    line = {'plain': T['lines']['yearly'], 'code': T['lines']['codeYearly'], 'unknown': T['lines']['yearly'], 'used': ''}[state]
    status = {
        'plain': '',
        'code': f'<p class="ok-t" role="status">{e(T["chooser"]["applied"])} {wp.Link(e(ch["remove"]))}</p>',
        'unknown': f'<p class="miss-t" role="alert">{e(ch["codeUnknown"])}</p>',
        'used': f'<p class="miss-t" role="alert">{e(ch["trialUsed"])}</p>',
    }[state]
    value = {'plain': '', 'code': 'GRIMBREW', 'unknown': 'GRIMBREWW', 'used': 'GRIMBREW'}[state]
    action = {'plain': T['chooser']['start'], 'code': T['chooser']['startCode'], 'unknown': T['chooser']['start'], 'used': T['chooser']['subscribeYearly']}[state]
    code = (f'<div class="codebox"><label class="f"><span class="t-label">{e(ch["codeLabel"])}</span>{code_input(value)}'
            f'<span class="t-small">{e(ch["codeHint"])}</span></label>{Btn(e(ch["apply"]))}</div>{status}')
    inner = (f'<div class="stack" style="padding:16px 20px 20px"><div class="f"><span class="t-label">{e(ch["intervalLabel"])}</span>{interval}</div>'
             + (f'<p class="trl">{e(line)}</p>' if line else '') + code
             + f'<p class="t-small">{e(ch["cardNote"])}</p><div class="acts">{Btn(e(action), primary=True)}</div></div>')
    head = PageHeader(e(ch['title']), answer=band('premium'))
    return page(ch['title'], head + PageForm(Plate(inner, e(ch['title']))), f'/premium/checkout chooser mock, {state}.')


# ================================================================================================ the confirmation
def confirmed() -> str:
    cf = C['confirmation']
    steps = wp.STEPS
    rows = [('ok', 1, e(steps['signIn']), f'<span class="tx">{e(wp.F["checkout_signed_in"])}</span>', ''),
            ('ok', 2, e(steps['pay']), e(wp.F['checkout_trial_saved']), ''),
            ('ok', 3, e(T['premium']['name'] + ' on'), f'<span class="tx" role="status">{e(T["status"]["premium"])}</span>', '')]
    after = (f'<div class="stack" style="padding-top:16px"><p class="trl">{e(T["lines"]["monthly"])}</p><p class="t-small">{e(cf["keep"])}</p>'
             f'<div class="acts">{Btn(e(cf["cancel"]))}</div></div>')
    head = PageHeader(e(cf['title']), answer=band('premium', change=False))
    return page(cf['title'], head + PageForm(Plate(Steps(rows), 'Checkout') + after), '/premium/checkout mock: the trial confirmed.')


# ================================================================================================ the account page
def plan(state: str) -> str:
    a = C['account']
    status = T['status']['premiumCanceled'] if state == 'canceled' else T['status']['premium']
    top = f'<div class="planfig"><div class="fd">{Figure(e(T["premium"]["name"]), size="sm")} {Tag(e(C["cardTag"]))}</div></div>'
    act = Btn(e(a['resume'])) if state == 'canceled' else Btn(e(a['cancel']))
    body = f'<div class="stack" id="plan">{top}<p class="trl">{e(status)}</p><div class="acts">{act}</div></div>'
    return Section(e(wp.ACC['planLabel']), body)


def account(state: str) -> str:
    head = PageHeader(e(wp.ACC.get('title', 'Your account')))
    body = head + Page('split', plan(state), '')
    if state == 'confirm':
        cf = C['confirm']
        body += Dialog(e(cf['title']), e(T['confirm']['body']), '', Btn(e(cf['no'])) + Btn(e(cf['yes']), primary=True), 'trial-cancel-title')
    return page('Account', body, f'/account mock, plan in trial ({state}).')


# ================================================================================================ guild settings
def guild() -> str:
    g = T['guild']
    row = Row(f'{e(g["plan"])} · {e(g["name"])}', e(T['status']['guild']), link=False, trail=Btn(e(C['account']['cancel']), small=True))
    body = PageHeader(e(g['name'])) + Page('split', Section('Plan', Plate(Rows([row]), 'Plan') + f'<p class="t-small" style="padding-top:12px">{e(T["lines"]["guild"])}</p>'), '')
    return page(g['name'], body, '/guild settings mock: the plan line in trial.')


# ================================================================================================ the emails
def email(kind: str) -> str:
    m = {'reminder': T['email']['reminder'], 'guild': T['email']['guildReminder'], 'receipt': T['email']['receipt']}[kind]
    facts = Facts([('From', 'Forever Sixty &lt;hello@foreversixty.gg&gt;'), ('To', 'the Stripe customer’s email'), ('Subject', e(m['subject']))])
    inner = f'<div class="mailhead">{facts}</div><div class="mail">{e(m["text"])}</div>'
    when = f'Sent {T["email"]["remindOn"]}, {T["email"]["lead"]} days before the first charge.' if kind != 'receipt' else 'Sent when the first invoice is paid.'
    body = PageHeader(e(m['subject']), lead=e(when)) + PageForm(Plate(inner, 'Email'))
    return page('Email', body, f'Email mock ({kind}): the plain-text message Resend sends.', who=None)


PAGES = {
    'web-trial-card': card,
    'web-trial-chooser': lambda: chooser('plain'),
    'web-trial-code': lambda: chooser('code'),
    'web-trial-code-unknown': lambda: chooser('unknown'),
    'web-trial-used': lambda: chooser('used'),
    'web-trial-confirmed': confirmed,
    'web-trial-account': lambda: account('trial'),
    'web-trial-confirm': lambda: account('confirm'),
    'web-trial-canceled': lambda: account('canceled'),
    'web-trial-guild': guild,
    'web-trial-email-reminder': lambda: email('reminder'),
    'web-trial-email-guild': lambda: email('guild'),
    'web-trial-email-receipt': lambda: email('receipt'),
}

if __name__ == '__main__':
    if sys.argv[1:2] == ['--list']:
        print(' '.join(PAGES))
    else:
        only = sys.argv[1:]
        for name, fn in PAGES.items():
            if not only or name in only:
                print(write(f'{name}.html', fn()))
