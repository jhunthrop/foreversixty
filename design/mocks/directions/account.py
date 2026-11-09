"""Mock batch 4a: Premium and plans, checkout, the account and sign-in in the directions system (plan 2026-11-09-redesign).

Data and copy, never invented: design/mocks/data/premium-facts.json (node design/mocks/premium_facts.mjs: the plans copy,
the plan cards' own price model, the checkout, account, sign-in and Discord copy), setup-facts.json (the paired device's
status line) and web/tests/e2e/fixtures/home/me-60.json (the signed-in account: Hunthrop#1894 and Thandrel). Drawn by
prims.py; the plan cards are panels, the Premium feature list a row list.
"""
import json
from html import escape
from pathlib import Path

import data as D
from common import write
from prims import PRIMS_CSS, facts, panel, player_cell, row_list, tag
from style import footer, header, medallion, page, titlebar

DATA = Path(__file__).resolve().parent.parent / 'data'
P = json.loads((DATA / 'premium-facts.json').read_text())
SETUP = json.loads((DATA / 'setup-facts.json').read_text())
ME = json.loads((D.ROOT / 'web/tests/e2e/fixtures/home/me-60.json').read_text())
PL = P['plans']
VIEWER = ('Thandrel', 'warrior', 60)

ACC_CSS = PRIMS_CSS + '''
.ac{padding-top:6px}.plans{display:grid;grid-template-columns:repeat(4,minmax(0,1fr));gap:18px;margin-top:22px;align-items:start}
.lede{margin-top:22px;max-width:80ch}.price s{font-size:18px;color:var(--muted);font-weight:400;margin-left:8px}
.spin{width:28px;height:28px;border:3px solid rgba(255,255,255,.18);border-top-color:var(--gold);border-radius:50%;margin-bottom:14px}
.plans .card{display:flex;flex-direction:column;margin-top:0}.plans .bd{display:flex;flex-direction:column;gap:10px;flex:1}
.price{font:600 34px var(--body);color:#fff}.price small{font-size:15px;color:var(--muted);font-weight:400}.aud{font-size:13.5px;color:var(--muted)}
.feat li{padding:5px 0 5px 20px;position:relative;font-size:14px;line-height:1.5}.feat li::before{content:"";position:absolute;left:3px;top:13px;width:7px;height:7px;border-radius:50%;background:var(--gain)}
.plans .go{margin-top:auto;padding-top:8px}.plans .go .btn,.plans .go .btn2{display:block;text-align:center}
.faq dt{color:#fff;font-weight:600;margin-top:14px}.faq dt:first-child{margin-top:0}.faq dd{margin:4px 0 0;font-size:14.5px}
.msg{max-width:560px;margin:22px 0 0}.msg h2{font-size:28px;margin-bottom:10px}.msg p{margin-top:8px}
@media(max-width:1100px){.plans{grid-template-columns:repeat(2,minmax(0,1fr))}}
@media(max-width:700px){.plans{grid-template-columns:minmax(0,1fr)}.plans>.card:nth-child(2){order:-1}}
'''


def plan_card(name: str, aud: str, price: str, per: str, features: list[str], button: str, framed: bool = False, sub: str = '',
              was: str = '') -> str:
    feats = '<ul class="feat">' + ''.join(f'<li>{escape(f)}</li>' for f in features) + '</ul>'
    body = (f'<p class="aud">{escape(aud)}</p><p class="price num">{price}<small>{per}</small>{f"<s>{was}</s>" if was else ""}</p>'
            + (f'<p class="aud">{escape(sub)}</p>' if sub else '') + feats + f'<div class="go">{button}</div>')
    return panel(name, body, framed=framed)


def premium_page() -> str:
    pr = P['priced']
    prem = pr['premium:yearly:founder']
    gld = pr.get('guild:yearly:founder') or {}
    raid = pr.get('raid_guild:yearly:founder') or {}
    free = PL['free']
    cards = (plan_card(free['name'], free['audience'], free['price'], '', free['features'], '<span class="muted">No card</span>')
             + plan_card(PL['premium']['name'], PL['premium']['audience'], prem['cell']['amount'], prem['cell']['suffix'],
                         [f['text'] for f in PL['premium']['features']], f'<a class="btn" href="#">{escape(prem["button"])}</a>', True, prem['cell']['secondLine'],
                         f'${P["plan_table"]["premium"]["yearly"]}')
             + plan_card(PL['guild']['name'], PL['guild']['audience'], gld.get('cell', {}).get('amount', '$120'), gld.get('cell', {}).get('suffix', ' /year'),
                         PL['guild']['features'], f'<a class="btn2" href="#">{escape(gld.get("button", "Subscribe"))}</a>', was=f'${P["plan_table"]["guild"]["yearly"]}')
             + plan_card(PL['raidGuild']['name'], PL['raidGuild']['audience'], raid.get('cell', {}).get('amount', '$200'), raid.get('cell', {}).get('suffix', ' /year'),
                         PL['raidGuild']['features'], f'<a class="btn2" href="#">{escape(raid.get("button", "Subscribe"))}</a>', was=f'${P["plan_table"]["raid_guild"]["yearly"]}'))
    faq = '<dl class="faq">' + ''.join(f'<dt>{escape(q["question"])}</dt><dd>{escape(q["answer"])}</dd>' for q in PL['faq']) + '</dl>'
    band = titlebar(None, PL['title'], '', ctl=[[(PL['interval']['labels']['yearly'], True), (PL['interval']['labels']['monthly'], False)]],
                    top=header(*VIEWER, ''))
    body = band + f'<main class="w ac"><p class="lede">{escape(PL["founder"]["open"])}</p><div class="plans">{cards}</div>{panel(PL["faqTitle"], faq)}</main>' + footer()
    return page('Forever Sixty, plans', ACC_CSS, body)


def checkout_page() -> str:
    c = P['checkout']
    body = (titlebar(None, 'Checkout', '', top=header(*VIEWER, ''))
            + f'<main class="w ac"><div class="msg card pad framed"><div class="spin" role="progressbar" aria-label="Working"></div><h2>{escape(c["processing"])}</h2><p class="muted">{escape(c["confirmationSlow"])}</p></div></main>' + footer())
    return page('Forever Sixty, checkout', ACC_CSS, body)


def account_page() -> str:
    a = P['account']
    user = ME['user']
    chars = row_list([(medallion(c['class'].lower(), 34), f'<span style="color:{D.CLASS_COLOR[c["class"].lower()]}">{escape(c["name"])}</span>',
                       f'Level {c["level"]} {escape(c["race"])} {escape(c["spec"])} {escape(c["class"])}',
                       f'item level {c["item_level"]}') for c in ME['characters']])
    you = facts([('BattleTag', f'<b class="strong">{escape(user["battletag"])}</b>'), ('Sign-in', 'Battle.net')])
    you += f'<p style="margin-top:12px"><a class="btn2" href="#">{escape(a["signOut"])}</a></p>'
    guild = ME['characters'][0].get('guild') or {}
    g0 = ME['guilds'][0]
    plan = facts([('Now', escape(P['account_plan']['head']['names']['free'])), (a['guildLabel'], f'{escape(g0["name"])} &middot; {escape(g0["rank"])}')])
    devices = row_list([('', escape(SETUP['devices'][0]['name']), escape(SETUP['device_status'][0].split(' · ', 1)[1]), '<a href="#" class="muted">Remove</a>')])
    discord = (f'<p>{escape(P["discord_link"]["linkedLead"])} <b class="strong">hunthrop</b></p>'
               f'<p style="margin-top:10px"><a class="btn2" href="#">{escape(P["discord_link"]["unlink"])}</a></p>')
    ds = a['deleteSection']
    delete = (f'<p style="font-size:14px">{escape(ds["line"])}</p><p style="margin-top:12px">'
              f'<a class="btn2" href="#" style="background:transparent;border-color:#c0392b;color:#f0736b">{escape(ds["button"])}</a></p>')
    main = panel('Characters', chars, framed=True) + panel(a['devicesLabel'], devices, action=f'<a href="#">{escape(a["downloadCompanion"])} &rsaquo;</a>') + panel('Discord', discord)
    rail = panel(a['youLabel'], you) + panel(a['planLabel'], plan, action='<a href="#">See plans &rsaquo;</a>') + panel(ds['title'], delete)
    band = titlebar(None, a['title'], '', top=header(*VIEWER, ''))
    body = band + f'<main class="w ac"><div class="twocol"><div>{main}</div><aside>{rail}</aside></div></main>' + footer()
    return page('Forever Sixty, account', ACC_CSS, body)


def login_page() -> str:
    s = P['signin']
    body = (titlebar(None, s['title'], '', top=header('', 'warrior', None, '').replace('<span class="nm" style="color:#C69B6D"></span>', ''))
            + f'<main class="w ac"><div class="msg card pad framed">'
            f'<p><a class="btn" href="#">Sign in with Battle.net</a></p><p class="muted" style="font-size:13.5px;margin-top:14px">{escape(s["privacy"])}</p></div></main>'
            + footer())
    return page('Forever Sixty, sign in', ACC_CSS, body)


def build() -> None:
    write('b4-premium', premium_page())
    write('b4-checkout', checkout_page())
    write('b4-account', account_page())
    write('b4-login', login_page())


if __name__ == '__main__':
    build()
