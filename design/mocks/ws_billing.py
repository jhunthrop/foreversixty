"""Web-layout reference mocks for the billing lane (RESTYLE-REST, 2026-11-07), composed ONLY from ws_prims:
  - web-guild-billing-states: the guild settings Billing tab blocks (the plan chooser, Premium seats with the Give
    picker open, the Move to Raid guild confirm, the Who pays transfer confirm). The settings page frame belongs to the
    guild lane (ws_guild_settings.py), which can import the block functions below; this sheet is a reference, not a
    parity state;
  - web-account-discord-paste: the account's Discord paste form (the notifications Section when connect is off);
  - web-admin-links: the owner's /admin/links (PageHeader, a split Page: the links' Plate of Rows, the New link rail).
Data and copy are read at render time from design/mocks/data/billing-facts.json (node design/mocks/billing_facts.mjs),
never retyped:

  node design/mocks/billing_facts.mjs && python3 design/mocks/ws_billing.py
FS_BILLING_FACTS and FS_MOCK_OUT point the generator elsewhere.
"""
import json
import os
import sys
from html import escape
from pathlib import Path

import ws_prims
from ws_prims import (Button as _Button, EmptyCrest, Figure, Link, Page, PageHeader, Plate, Row, Rows, Section, Segmented, Select, Switch, Tag,
                      assert_no_surface_styles, crest_img, write)
from ws_lib import page as _page

HERE = Path(__file__).resolve().parent
F = json.loads(Path(os.environ.get('FS_BILLING_FACTS') or HERE / 'data' / 'billing-facts.json').read_text())
G, D, A = F['guild'], F['discord'], F['admin']
WHO = ('Thandrel', 'warrior', 'Level 60')

CSS = '''<style>
/* page-level layout only (no surfaces) */
.stack{display:flex;flex-direction:column;gap:12px;min-width:0}.acts{display:flex;gap:10px 12px;flex-wrap:wrap;align-items:center}
.st-grid{display:grid;gap:32px 40px;grid-template-columns:repeat(auto-fill,minmax(420px,1fr))}.st-cell{display:flex;flex-direction:column;gap:6px;min-width:0}
.radio{width:18px;height:18px;accent-color:var(--gold)}.figs4{display:flex;flex-wrap:wrap;gap:8px 20px}.figs4>span{width:128px;display:flex;flex-direction:column}.spacer{width:76px}
.f{display:flex;flex-direction:column;gap:6px;min-width:0}.f .t-small{font-size:12px;line-height:16px}
.status{min-height:22px}.page-foot{padding-bottom:48px}
.steps{margin:0;padding-left:20px;display:flex;flex-direction:column;gap:8px;font-size:13px;color:var(--muted)}.steps li::marker{color:var(--gold);font-weight:700}
section>.sec-h+*{margin-top:12px}
</style>'''


def e(t) -> str:
    return escape(str(t), quote=False)


def Input(placeholder: str) -> str:
    """The built Field's input (ui/Field: a real textbox), at the mock library's field look."""
    return f'<input class="field" type="text" placeholder="{placeholder}">'


def Btn(label: str, primary: bool = False, small: bool = True) -> str:
    """The system Button as the build renders an action (a <button>)."""
    return _Button(label, primary, small).replace('<a ', '<button type="button" ', 1).replace(' href="#"', '', 1).replace('</a>', '</button>')


def render(title: str, body: str, nav: str, note: str) -> str:
    assert_no_surface_styles(body)
    return _page(title, body + CSS + f'<style>{ws_prims.PRIM_CSS}</style>', nav, WHO, note, '')


# ================================================================================================ the guild Billing tab blocks
def plan_chooser() -> str:
    """GuildPlanChooser: the interval Segmented, one Row per plan (radio lead, the picked one selected), the count, the
    compare link, the one gold button, the trial and founder lines."""
    c = G['chooser']
    rows = [Row(e(t['name']), f'{e(t["line"])}<br>{e(t["list"])}', lead='<input class="radio" type="radio" name="tier"' + (' checked' if k == 'guild' else '') + '>',
                trail=Figure(e(t['price']), unit=e(t['unit']), size='sm'), link=False, selected=k == 'guild') for k, t in G['tiers'].items()]
    return (f'<div class="stack">{Segmented([(e(G["interval"]["labels"]["yearly"]), True), (e(G["interval"]["labels"]["monthly"]), False)])}'
            f'{Rows(rows)}<p class="t-small">{e(G["verified"])}</p>{Link(e(c["compare"]) + " &rsaquo;")}'
            f'<div class="acts">{_Button(e(G["subscribe"]), primary=True)}</div><p class="t-small">{e(c["trialLine"])}</p><p class="t-small">{e(c["founderLine"])}</p></div>')


def seats(picking: bool = True) -> str:
    """GuildSeats: a level 3 Section with its count, the helper, an open Row per seat (crest, class-coloured name, the
    seat's lines, the action); Give seat opens the picker in its Row's foot."""
    s = G['seats']
    held = Row('Kraggor', f'{e(s["held"]["given"])}<br>{e(s["held"]["lock"])}', lead=crest_img('warrior', 32), title_color='var(--class-warrior)',
               trail=Btn(e(s['move'])), link=False)
    offered = Row('Mirelle', f'{Tag(e(s["waiting"]), "muted")}<br>{e(s["offered"]["note"])}', lead=crest_img('priest', 32), title_color='var(--class-priest)',
                  trail=Btn(e(s['withdraw'])), link=False)
    p = s['picker']
    foot = (f'<div class="stack"><label class="f"><span class="t-label">{e(p["title"])}</span>{Select(e(p["option"]))}</label>'
            f'<div class="acts">{Btn(e(p["give"]), primary=True)}{Btn(e(p["cancel"]))}</div></div>') if picking else ''
    empty = Row(f'<span class="t-small">{e(s["notGiven"])}</span>', lead=EmptyCrest(32), trail=Btn(e(s['give'])), link=False, foot=foot)
    return Section(e(s['heading']), f'<p class="t-small">{e(s["explainer"])}</p>{Rows([held, offered, empty], open_=True)}', meta=e(s['count']))


def upgrade_confirm() -> str:
    """GuildUpgradeConfirm: a selected Row, the prorated sentence, then Move to Raid guild and Keep Guild in its foot."""
    u = G['upgrade']
    return Rows([Row(e(u['sentence']), link=False, selected=True, foot=f'{Btn(e(u["commit"]), primary=True)}{Btn(e(u["keep"]))}')])


def who_pays() -> str:
    """GuildWhoPays: the Who pays row, then the transfer confirm as a selected Row (the officer select, what it does)."""
    w = G['whoPays']
    top = Row(e(w['heading']), e(w['help']), trail=Btn(e(w['transfer'])), link=False)
    confirm = Row(e(w['transferTo']), link=False, selected=True,
                  foot=(f'<div class="stack">{Select("Mirelle")}<p>{e(w["consequence"])}</p><p class="t-small">{e(w["stays"])}</p>'
                        f'<div class="acts">{Btn(e(w["send"]), primary=True)}{Btn(e(w["cancel"]))}</div></div>'))
    return Rows([top, confirm])


def guild_billing_states() -> str:
    cells = [('The plan chooser (no plan, 12 verified)', plan_chooser()), ('Premium seats, Give seat open', seats()),
             ('Move to Raid guild, confirm', upgrade_confirm()), ('Who pays, the transfer confirm', who_pays())]
    grid = ''.join(f'<div class="st-cell"><span class="t-label">{lab}</span>{html}</div>' for lab, html in cells)
    body = '<div class="pc" style="padding-top:24px"><h1 class="t-title">Guild billing: element states</h1></div>' + Page('ledger', f'<div class="st-grid">{grid}</div>')
    return render('Guild billing states', body, 'Guilds', 'Guild settings Billing tab blocks; data billing-facts.json (guild-billing-copy.ts, seats-copy.ts, plans).')


# ================================================================================================ the Discord paste form
def discord_paste() -> str:
    """DiscordPasteForm in the account's notifications Section: the steps (an ordered list), the Field, the trust helper,
    the event switches, Save."""
    steps = ''.join(f'<li>{e(s)}</li>' for s in D['howSteps'])
    form = (f'<div class="stack"><p>{e(D["intro"])}</p><ol class="steps">{steps}</ol>'
            f'<label class="f"><span class="t-label">{e(D["fieldLabel"])}</span>{Input(e(D["placeholder"]))}</label>'
            f'<p class="t-small">{e(D["trustHelper"])}</p>'
            f'{Switch(e(D["eventSimFinished"]))}{Switch(e(D["eventUploadParsed"]))}{Switch(e(D["eventUploadPrivate"]), on=False)}'
            f'<div class="acts">{Btn(e(D["save"]), primary=True, small=False)}</div></div>')
    body = '<div class="pc" style="padding-top:24px"><h1 class="t-title">Account: the Discord paste form</h1></div>' + Page('ledger', f'<div style="max-width:420px">{Section(e(D["label"]), form)}</div>')
    return render('Discord paste form', body, '', 'Account notifications Section, connect off; data billing-facts.json (notifications-copy.ts).')


# ================================================================================================ /admin/links
def figures(steps: list[dict]) -> str:
    return '<span class="figs4">' + ''.join(f'<span>{Figure(e(s["value"]), e(s["label"]), size="sm")}<span class="t-small">{e(s["note"])}</span></span>' for s in steps) + '</span>'


def admin_links() -> str:
    c = A['copy']
    rows = [Row(e(r['label']), f'/go/{e(r["slug"])}' + (f' &rarr; {e(r["target"])}' if r['target'] != '/' else ''),
                trail=figures(r['steps']) + Btn(e(c['copy'])), link=False) for r in A['rows']]
    rows.append(Row(e(c['totals']), trail=figures(A['totals']) + '<span class="spacer"></span>', link=False))
    main = Section(e(c['tableTitle']), Plate(Rows(rows), e(c['tableTitle'])), meta=e(A['meta']))
    form = (f'<div class="stack"><p class="t-small">{e(c["formJob"])}</p>'
            + ''.join(f'<label class="f"><span class="t-label">{e(c[k + "Label"])}</span>{Input(e(ph))}<span class="t-small">{e(c[k + "Hint"])}</span></label>'
                      for k, ph in (('slug', 'reddit-classicwow'), ('target', '/'), ('label', 'Reddit: r/classicwow')))
            + f'<div class="acts">{Btn(e(c["create"]), primary=True, small=False)}</div><div class="status" aria-live="polite"></div></div>')
    # The page's foot: the built <main> ends a section gap (48 px) above the site footer.
    body = PageHeader(e(c['pageTitle']), lead=e(c['job'])) + f'<div class="page-foot">{Page("split", main, Section(e(c["formTitle"]), form))}</div>'
    return render('Links · Forever Sixty', body, '', '/admin/links mock (the owner); data billing-facts.json (admin-copy.ts, funnel.ts over the e2e links).')


PAGES = {'web-guild-billing-states': guild_billing_states, 'web-account-discord-paste': discord_paste, 'web-admin-links': admin_links}

if __name__ == '__main__':
    if sys.argv[1:2] == ['--list']:
        print(' '.join(PAGES))
    else:
        only = sys.argv[1:]
        for name, fn in PAGES.items():
            if not only or name in only:
                print(write(f'{name}.html', fn()))
