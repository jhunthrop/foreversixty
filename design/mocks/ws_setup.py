"""Web-layout mocks: /setup (the addon, the companion, signing in) and /connect (the companion's Battle.net hand-off),
composed ONLY from ws_prims.

Spec: design/specs/2026-11-07-setup.md. Data and copy, read at render time, never retyped:
  - design/mocks/data/setup-facts.json (node design/mocks/setup_facts.mjs): lib/addon/copy.ts, cta-copy.ts,
    lib/account/connect-copy.ts, character-list-copy.ts, lib/reports/copy.ts and my-reports-copy.ts; the page's own line
    builders (lib/setup/progress.ts, lib/reports/device-status.ts, lib/account/connect-flow.ts) run on the fixtures at the
    pinned clock 2026-11-06T10:00Z; web/src/data/companion.json and active-build.json;
  - the fixtures those lines read: web/tests/e2e/fixtures/home/me-34.json (Felsnap), me-new.json (no character) and
    web/tests/e2e/fixtures/setup/*.json (the companion JHUNT-PC, the connect lookup, a pairing code).
The plate's rows are lib/setup/steps.ts run on those fixtures (FACTS['rows']): the page's island draws the same rows.

  node design/mocks/setup_facts.mjs && python3 design/mocks/ws_setup.py
  for p in $(python3 design/mocks/ws_setup.py --list); do node design/mocks/ws_shoot.mjs $p 390,1440,1920,2000; done
"""
import json
import os
import re
import sys
from html import escape
from pathlib import Path

from mocklib import CLASS_COLOR
import ws_prims
from ws_prims import (Button as _Button, Capture, Code, Facts, Field, FormRow, Link, PageForm, PageHeader, Page, Plate, Row, RowGrid,
                      Rows, Section, Skeleton, Steps, Tag, assert_no_surface_styles, crest_img, write)
from ws_lib import OUT, page as _page

ROOT = Path(__file__).resolve().parents[2]
FACTS = json.loads(Path(os.environ.get('FS_SETUP_FACTS') or Path(__file__).resolve().parent / 'data' / 'setup-facts.json').read_text())
A, C, CH, PC = FACTS['addon'], FACTS['connect'], FACTS['characters'], FACTS['pairing']
PUBLIC = os.path.relpath(ROOT / 'web' / 'public', OUT)
ME = FACTS['me']['addon']['characters'][0]
WHO = (ME['name'], ME['class'].lower(), str(ME['level']))

def LinkButton(label: str, primary: bool = False, small: bool = False) -> str:
    return _Button(label, primary, small)


def Button(label: str, primary: bool = False, small: bool = False, off: bool = False) -> str:
    b = _Button(label, primary, small).replace('<a ', '<button type="button" ', 1).replace(' href="#"', '', 1).replace('</a>', '</button>')
    return b.replace('class="btn', 'aria-disabled="true" class="btn is-off', 1) if off else b


def arrow(label: str) -> str:
    return Link(f'{escape(label)} &rsaquo;')


def disclosure_summary(label: str) -> str:
    """The built ui/Disclosure's `link` summary: the label and a gold caret 8 px apart (the caret is its own span, so
    the phone's wrap point is the build's)."""
    return f'<summary class="link">{escape(label)}<span aria-hidden="true">&rsaquo;</span></summary>'


def sentences(text: str) -> list[str]:
    """A copy paragraph split at its full stops, each sentence verbatim (the companion's three steps)."""
    return [s.strip() for s in re.split(r'(?<=\.)\s+(?=[A-Z])', text) if s.strip()]


def img(src: str) -> str:
    return f'{PUBLIC}{src}'


CSS = '''<style>
.ph-only{display:block}@media(min-width:768px){.ph-only{display:none}}
section>.sec-h+*{margin-top:12px}
.acts{display:flex;gap:12px;align-items:center;flex-wrap:wrap}
.body{display:flex;flex-direction:column;gap:14px}.body p{font-size:15px;line-height:24px;max-width:72ch;text-wrap:pretty}
.body .t-small{max-width:72ch}
.hand{display:flex;flex-direction:column;gap:6px;padding-top:12px;border-top:1px solid var(--line-soft)}
.hand .acts{gap:20px}
.ans-who{display:flex;align-items:center;gap:12px}.ans-who b{display:block;font:700 18px/22px var(--font-body)}.ans-links{display:flex;gap:16px;flex-wrap:wrap;font-size:16px}  /* the links take the page head's 16 px */
@media(min-width:900px){.head .ans:has(>.ans-who){min-width:300px}.ans-who{justify-content:flex-end;text-align:right}.ans-links{justify-content:flex-end}}
.pair{display:flex;flex-direction:column;gap:8px;align-items:flex-start;padding-top:12px;border-top:1px solid var(--line-soft)}
.pair-h{font:700 15px/1.5 var(--font-body);color:var(--strong)}  /* PairingBlock's label (h-audit): 15 px bold, no caps label */
.paste textarea{width:100%;min-height:72px;resize:vertical;font:600 14px/20px ui-monospace,'SF Mono',Menlo,monospace}
.kv code.c{font-size:14px}.kv dd .t-small{font-weight:400}
.surf .row{align-items:start;align-content:start}.surf .row .foot{padding-top:12px}.surf .cap{--mh:240px!important}
.cmds .kv dt{color:var(--text);padding-right:12px}
.cc{display:flex;flex-direction:column;gap:8px}.sr-only{position:absolute;width:1px;height:1px;overflow:hidden;clip:rect(0,0,0,0)}.cmds .kv{grid-template-columns:minmax(0,1fr) auto}.cmds .kv dd{white-space:nowrap}
.alert{color:var(--strong)}.steps .row .sub .tx{color:var(--text)}
.surf-fold{display:none}
@media(max-width:767px){.surf-fold{display:inline-flex}.surf{display:none}}
@media(max-width:600px){.steps .row{align-items:start}}
/* the built sm Button is a 44 px target below 1024 (web/src/lib/ui/button.ts BUTTON_SM) */@media(max-width:1023px){.btn.sm{height:44px}}
.tm{display:inline-flex;align-items:center;gap:12px}
.dl{display:flex;flex-direction:column;gap:4px}.dl .acts{column-gap:16px;row-gap:8px}.dl-meta{min-height:20px}
.dl-copy{gap:8px;align-items:flex-start}
@media(min-width:900px){.head .ph-actions .dl{align-items:flex-end}.head .ph-actions .dl .acts{justify-content:flex-end}}
.dl-other summary{cursor:pointer;color:var(--gold);display:inline-flex;align-items:center;gap:8px}@media(max-width:767px){.dl-other summary{min-height:44px}}.dl-other ul{margin-top:8px;display:flex;flex-direction:column;gap:4px;list-style:none;padding:0}
.mark-ok{width:40px;height:40px;border-radius:50%;display:inline-flex;align-items:center;justify-content:center;font:700 20px/40px var(--font-body);color:#07090d;background:var(--kill);flex:none}
</style>'''


def render(title: str, body: str, nav: str, who, chrome_state: str, note: str) -> str:
    assert_no_surface_styles(body)
    return _page(title, body + CSS + f'<style>{ws_prims.PRIM_CSS}</style>', nav, who, note, chrome_state)


# ------------------------------------------------------------------------------------------------ /setup: the plate
def _line(l: dict) -> str:
    if l.get('skeleton'):
        # phoneOnly (web lib/setup/steps.ts): the second line an answer wraps to on a phone, reserved below md
        return f'<span class="ph-only">{Skeleton(240, 14)}</span>' if l.get('phoneOnly') else Skeleton(240, 14)
    tone = {'muted': '', 'text': ' class="tx"', 'ok': ' class="ok"', 'warn': ' class="warn"', 'miss': ' class="miss"'}[l['tone']]
    role = ' role="alert"' if l.get('alert') else ''
    code = (' ' + Code(l['code'])) if l.get('code') else ''
    return f'<span{tone}{role}>{escape(l["text"])}</span>{code}'


PLATFORM_OF = {'mac': 'mac', 'mobile': 'mobile', 'unknown': 'unknown'}


def platform_of(state: str) -> str:
    """The visitor's system each state is drawn for (web/tests/e2e/support/setup-page.ts): Windows unless named."""
    return PLATFORM_OF.get(state, 'windows')


def download_block(state: str, primary: bool) -> str:
    """The companion's download (components/setup/CompanionDownload.svelte): the own build's button, its version and size,
    and Other platforms; for a system the page cannot name, every build side by side with the version alone; on a phone,
    "runs on your PC or Mac" and Copy link (row 3 draws its phone row from steps.ts, so only the head and the Section
    reach that branch)."""
    if platform_of(state) == 'mobile':
        return f'<div class="dl dl-copy"><span class="t-small">{escape(A["setupCompanionMobile"])}</span>{Button(A["setupCopyLink"])}</div>'
    d = FACTS['download_for'][platform_of(state)]
    meta = lambda text: f'<span class="t-small dl-meta">{escape(text or "")}</span>'
    if d['own'] is None:
        buttons = ''.join(LinkButton(b['label']) for b in d['others'])
        return f'<div class="dl" role="group" aria-label="{escape(A["setupDownloadCompanion"])}"><div class="acts">{buttons}</div>{meta(FACTS["version_line"])}</div>'
    own = d['own']
    other = ''.join(f'<li class="t-small">{arrow(b["label"])} <span>{escape(b["meta"] or "")}</span></li>' for b in d['others'])
    fold = f'<details class="dl-other">{disclosure_summary(A["setupOtherPlatforms"])}<ul>{other}</ul></details>' if other else ''
    return f'<div class="dl"><div class="acts">{LinkButton(own["label"], primary=primary)}{fold}</div>{meta(own["meta"])}</div>'


def _action(a: dict | None, state: str = 'out') -> str:
    if a is None:
        return ''
    if a['id'] == 'download':
        return download_block(state, a['primary'])
    if a['id'] == 'install':
        # The addon's store group (SetupStepList.svelte): CurseForge, gold when it is the next step, then Wago.
        return f'<div class="acts" role="group" aria-label="{escape(a["label"])}">{LinkButton(A["installCurseForge"], primary=a["primary"])}{LinkButton(A["installWago"])}</div>'
    if a['id'] == 'update':
        # A ticked step's update: quiet gold links.
        return f'<div class="acts" role="group" aria-label="{escape(a["label"])}">{arrow(A["setupUpdateCurseForge"])}{arrow(A["setupUpdateWago"])}</div>'
    if a['kind'] == 'link':
        return arrow(a['label'])
    if a['kind'] == 'button':
        return Button(a['label'], primary=a['primary'])
    return LinkButton(a['label'], primary=a['primary'])


def step_rows(state: str) -> list[tuple]:
    """The three steps for a state: lib/setup/steps.ts's rows on the fixtures (FACTS['rows'])."""
    return [(r['mark'], r['n'], r['title'], '<br>'.join(_line(l) for l in r['lines']), _action(r['action'], state)) for r in FACTS['rows'][state]]


def download_actions(state: str) -> str:
    """The head's actions: the companion's download in every state (owner 2026-11-08: "a download link at the top"),
    never gold (the plate keeps the page's one primary). The signed-in character is the header's and step 1's."""
    return f'<div class="dl-head">{download_block(state, primary=False)}</div>'


def row_done(state: str, step: str) -> bool:
    """Whether the plate's step (lib/setup/steps.ts StepId: signin, companion, addon) is ticked."""
    return next(r for r in FACTS['rows'][state] if r['id'] == step)['mark'] == 'ok'


# ------------------------------------------------------------------------------------------------ /setup: the sections
def addon_section(state: str) -> str:
    hand = (f'<div class="hand"><span class="t-label">{escape(A["setupHandInstall"])}</span>'
            f'<p class="t-small">{escape(FACTS["hand_install"]["mac" if platform_of(state) == "mac" else "other"])}</p>'
            f'<div class="acts">{arrow(A["installZip"])}{arrow(A["installDataRelease"])}</div><p class="t-small">{escape(FACTS["data_build"])}</p></div>')
    # U4 (rulings): once the addon step is ticked its buttons are gone; both stores stay here for a second PC.
    again = f'<div class="acts">{arrow(A["installCurseForge"])}{arrow(A["installWago"])}</div>' if row_done(state, 'addon') else ''
    body = (f'<div class="body" id="addon"><p class="t-small">{escape(A["installDataHint"])}</p>{again}{hand}</div>')
    return Section(A['setupAddonTitle'], body, panel=True)


def companion_section(state: str) -> str:
    # The visitor's own build (lib/setup/downloads.ts): None on a phone or an unknown system.
    comp = FACTS['download_for'][platform_of(state)]
    steps = sentences(A['setupCompanionBody'])   # [what it is, the setting, /combatlog each time you log in]
    lead_sentence, rest = steps[0], steps[1:]
    howto = [A['setupCompanionPair']] + rest
    def withcode(s: str) -> str:
        return escape(s).replace('/combatlog', Code('/combatlog'))
    rows = Rows([Row(withcode(s), rank=str(i + 1), link=False) for i, s in enumerate(howto)], open_=True)
    cap = A['setupCompanionCapture']
    narrow = {**cap['narrow'], 'src': img(cap['narrow']['src'])}
    pic = Capture(img(cap['src']), cap['width'], cap['height'], cap['alt'], tag=cap['exampleLabel'], lazy=False,
                  narrow=narrow)
    pair = ''
    if state not in ('out', 'mac', 'mobile', 'unknown', 'export', 'error', 'loading', 'signin-failed'):
        lc = FACTS['logs_companion']
        if state == 'done':
            status = ''.join(f'<p class="t-small" style="color:var(--text)">{escape(x)}</p>' for x in FACTS['device_status'])
            pair = f'<div class="pair">{status}<span class="pair-h">{escape(lc["pairAnotherLabel"])}</span><p class="t-small">{escape(PC["companionFirst"])}</p><div class="acts">{Button(PC["action"], small=True)}</div></div>'
        elif state == 'pairing':
            pair = (f'<div class="pair"><span class="pair-h">{escape(lc["pairLabel"])}</span>{Code(FACTS["pair_code"]["code"], big=True)}'
                    f'<p class="t-small">{escape(PC["code_line"])}</p></div>')
        else:
            pair = f'<div class="pair"><span class="pair-h">{escape(lc["pairLabel"])}</span><p class="t-small">{escape(PC["companionFirst"])}</p><div class="acts">{Button(PC["action"], small=True)}</div></div>'
    # The signing and first-run line (U7): the own build's; every build's, named, for an unknown system; none on a phone.
    # It belongs to the button: 8 px under it, in body ink (installers review 2026-11-10; SetupCompanion.svelte).
    first_run = '<p class="t-small" style="color:var(--text)">{}</p>'
    if comp['own'] is not None:
        plat = first_run.format(escape(comp["own"]["firstRun"]))
    elif platform_of(state) == 'unknown':
        plat = ''.join(first_run.format(f'{escape(b["platform"])}: {escape(b["firstRun"])}') for b in comp['others'])
    else:
        plat = ''
    # Owner 2026-11-08: the download right under the heading, in every state (it replaces U4's quiet "another PC" link).
    download = f'<div style="display:flex;flex-direction:column;gap:8px">{download_block(state, primary=False)}{plat}</div>'
    body = (f'<div class="body" id="companion">{download}<p>{escape(lead_sentence)}</p>'
            f'{rows}{pic}{pair}</div>')
    return Section(A['setupCompanionTitle'], body, meta=escape(A['setupCompanionNote']), panel=True)


def surfaces_section() -> str:
    rows = []
    for s in A['surfaces']:
        caps = ''.join(Capture(img(i['src']), i['width'], i['height'], i['alt'], lazy=False) for i in s['images'])
        rows.append(Row(escape(s['title']), escape(s['body']), link=False, foot=caps))
    # U5 (rulings): open from 768, folded on a phone behind one button (SetupSurfaces.svelte: aria-expanded, no details).
    fold = Button(A['surfacesShow']).replace('class="btn', 'class="btn surf-fold', 1).replace('<button ', '<button aria-expanded="false" ', 1)
    return Section(A['surfacesTitle'], f'{fold}<div class="surf">{RowGrid(rows, 2)}</div>', panel=True)


def paste_section(state: str) -> str:
    field = '<textarea class="field grow" aria-label="{0}" placeholder="{1}">{2}</textarea>'
    filled = {'export': 'FS1:1.60.1.70790:hunter:dwarf:…', 'paste-error': 'FS2:nope'}.get(state, '')
    form = FormRow([(escape(A['pasteFieldLabel']), field.format(escape(A['pasteTitle']), escape(A['pastePlaceholder']), filled), True)], Button(A['pasteAction']))
    # The built FormRow is a named <form> (ui/FormRow.svelte): the mock's row is that form too.
    form = form.replace('<div class="formrow">', f'<form class="formrow" aria-label="{escape(A["setupTryExport"])}">', 1)[:-len('</div>')] + '</form>'
    after = ''
    if state == 'paste-error':
        after = f'<p class="alert" role="alert">{escape(FACTS["wrong_prefix"])}</p>'
    if state == 'export':
        after = f'<div class="acts">{arrow(A["pasteOpenPlanner"])}{arrow(A["pasteOpenSim"])}</div><p class="t-small">{escape(A["pasteSignInHint"])}</p>'
    body = f'<div class="body paste" id="paste"><p class="t-small">{escape(A["setupTryExportBody"])}</p>{form}{after}</div>'
    return Section(A['setupTryExport'], body, panel=True)


def commands_section() -> str:
    return Section(A['commandsTitle'], f'<div class="cmds">{Facts([(escape(c["description"]), Code(c["command"])) for c in A["commands"]])}</div>', panel=True)


def handshake_section() -> str:
    paras = ''.join(f'<p class="t-small">{escape(s)}</p>' for s in A['handshake'])
    return Section(A['handshakeTitle'], f'<div class="body">{paras}</div>', panel=True)


def flows_section() -> str:
    rows = [Row(escape(A['flowOutTitle']), escape(A['flowOutBody']), link=False), Row(escape(A['flowInTitle']), escape(A['flowInBody']), link=False)]
    return Section(A['setupHowItWorks'], Rows(rows, open_=True), panel=True)


SIGNED_IN = ('none', 'importing', 'failed', 'bnet', 'addon', 'pairing', 'done', 'loading')


def setup_page(state: str) -> str:
    paste_state = state if state in ('export', 'paste-error') else ''
    plate_state = 'out' if state == 'paste-error' else state
    head = PageHeader(escape(A['setupPageTitle']), actions=download_actions(state))
    main = Plate(Steps(step_rows(plate_state)), 'Your setup').replace('class="plate"', 'class="plate framed"', 1) + companion_section(plate_state) + addon_section(plate_state) + surfaces_section()
    rail = paste_section(paste_state) + commands_section() + handshake_section() + flows_section()
    body = head + Page('split', main, rail, panels=True)
    who = WHO if state in ('bnet', 'addon', 'pairing', 'done') else None
    chrome = {'none': 'new', 'importing': 'new', 'failed': 'new', 'loading': 'pending'}.get(state, '')
    note = f'/setup mock, {state}. Data: design/mocks/data/setup-facts.json (copy modules, progress.ts lines on fixtures home/me-34, me-new, setup/devices at 2026-11-06T10:00Z).'
    return render(A['setupPageTitle'], body, 'Get set up', who, chrome, note)


# ------------------------------------------------------------------------------------------------ /connect
def connect_page(state: str) -> str:
    """The framed plate (directions mock b4-connect): the check line and the code at display size, Device and Platform
    under it, the consent line, the actions. The code never moves between loading, signed out and pending (MUST 3)."""
    L = FACTS['lookup']
    big = f'<span class="sr-only">{escape(C["codeRowLabel"])}</span>{Code(FACTS["code"], big=True)}'
    after = f'<span class="t-small">{escape(C["afterSignIn"])}</span>'
    device = [(escape(C['deviceLabel']), escape(L['name'])), (escape(C['platformLabel']), escape(L['platform']))]
    title, meta, crest = C['title'], '', ''
    check, facts, consent, acts, tag = '', device, False, '', ''
    if state == 'loading':
        meta, check = C['checking'], C['pendingCodeCheck']
        big = '<span class="skeleton" style="display:block;width:150px;height:30px"></span>'
        facts = [(Skeleton(64, 14), Skeleton(120, 14))] * 2
        acts = f'<div class="acts">{Skeleton(200, 44)}{Skeleton(96, 44)}</div>'
    elif state == 'out':
        meta, check, consent = C['signedOutLead'], C['codeCheck'], True
        facts = [(escape(C['deviceLabel']), after), (escape(C['platformLabel']), after)]
        acts = f'<div class="acts">{LinkButton(C["battlenet"], primary=True)}</div>'
    elif state in ('pending', 'connecting'):
        title, meta = FACTS['connect_fns']['pendingTitle'], C['pendingBody']
        check, consent = C['pendingCodeCheck'], True
        busy = state == 'connecting'
        acts = (f'<div class="acts">{Button(C["connecting"] if busy else C["connect"], primary=True, off=busy)}{Button(C["cancel"], off=busy)}</div>'
                f'<p class="t-small">{escape(C["notYou"])} {Link(escape(C["signOut"]) + ".")}</p>')
    elif state == 'success':
        title, meta = C['successTitle'], C['successBody']
        crest = '<span class="mark-ok ok" role="img" aria-label="Done">&#10003;</span>'
        tag = Tag(C['connectedTag'], 'kill')
        acts = f'<div class="acts">{LinkButton(C["getSetUp"], primary=True)}{LinkButton(C["yourAccount"])}</div>'
    elif state == 'cancelled':
        title, meta = C['cancelledTitle'], C['cancelledBody']
        acts = f'<div class="acts">{LinkButton(C["close"])}</div>'
    elif state == 'expired':
        title, meta, facts = C['expiredTitle'], C['expiredBody'], []
        acts = f'<div class="acts">{LinkButton(C["backHome"])}</div>'
    elif state == 'used':
        title, meta = C['usedTitle'], C['usedBody']
        acts = f'<div class="acts">{LinkButton(C["openReports"], primary=True)}</div>'
    elif state == 'failed':
        title, meta = C['failedTitle'], C['failedBody']
        acts = f'<div class="acts">{Button(C["tryAgain"], primary=True)}</div>'
    line = f'<p>{escape(check)}</p>' if check else ''
    code = f'<div class="cc">{line}<div class="acts">{big}{tag}</div></div>'
    rows = Facts(facts) if facts else ''
    agree = f'<p class="t-small">{escape(C["consent"])}</p>' if consent else ''
    inner = f'<div class="body" style="padding:18px 20px 20px">{code}{rows}{agree}{acts}</div>'
    head = PageHeader(escape(title), crest=crest, meta=[escape(meta)])
    if state == 'failed':
        head = head.replace('<span class="t-small">', '<span class="t-small" role="alert">', 1)
    body = head + PageForm(Plate(inner, 'The companion').replace('class="plate"', 'class="plate framed"', 1))
    signed = state not in ('out', 'expired', 'loading')
    chrome = 'new' if signed else ('pending' if state == 'loading' else '')
    note = f'/connect mock, {state}. Data: design/mocks/data/setup-facts.json (lib/account/connect-copy.ts, fixtures setup/connect-device.json, home/me-new.json).'
    return render(C['title'], body, '', None, chrome, note)


SETUP_STATES = ['out', 'mac', 'mobile', 'unknown', 'export', 'paste-error', 'signin-failed', 'error', 'loading', 'none', 'importing', 'failed', 'bnet', 'addon', 'pairing', 'done']
CONNECT_STATES = ['out', 'loading', 'pending', 'connecting', 'success', 'cancelled', 'expired', 'used', 'failed']
PAGES = {('web-setup' if s == 'out' else f'web-setup-{s}'): (lambda s=s: setup_page(s)) for s in SETUP_STATES}
PAGES |= {('web-connect' if s == 'out' else f'web-connect-{s}'): (lambda s=s: connect_page(s)) for s in CONNECT_STATES}

if __name__ == '__main__':
    if sys.argv[1:] == ['--list']:
        print(' '.join(PAGES))
    else:
        only = sys.argv[1:]
        for name, fn in PAGES.items():
            if not only or name in only:
                print(write(f'{name}.html', fn()))
