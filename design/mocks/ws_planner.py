"""Web-layout mock: the talent planner (/planner), composed ONLY from ws_prims primitives.

Spec: design/specs/2026-11-08-planner.md. Every talent, tree, rank, description, arrow, point order, compare line and
string is read from design/mocks/data/planner-facts.json, which design/mocks/planner_facts.mjs writes by running the
planner's own models (lib/planner/*, its copy modules) on build 1.60.1.70790. Typed in: the strings this round proposes
(PROPOSED, spec section 6.2), the live estimate's figures (SAMPLE: it runs the engine in the browser), and the layout.

  node design/mocks/planner_facts.mjs && python3 design/mocks/ws_planner.py
  python3 design/mocks/ws_planner.py [mock ...]                 (every mock, or just those named)
  node design/mocks/ws_shoot.mjs <mock> 390,1440,1920,2000     (each name in `python3 design/mocks/ws_planner.py --list`)
  python3 design/mocks/ws_planner.py compare <before dir>        (before/after sheets; before: shoot_planner_before.mjs)
FS_PLANNER_FACTS and FS_MOCK_OUT point the generator elsewhere (a parity spec's test-time render).
"""
import base64
import json
import os
import sys
from html import escape
from pathlib import Path

from mocklib import CLASS_COLOR, ROOT
from ws_prims import (Button, Field, ClassPicker, CodeField, Figure, Link, Notice, OrderSlot, Ordered, Page, PageHeader, PointOrder,
                      Section, Skeleton, Switch, Tag, TalentCell2, TalentEditor, TalentTip, build_page, control_bar, crest_img,
                      icon, sel, REGION)

HERE = Path(__file__).resolve().parent
F = json.loads(Path(os.environ.get('FS_PLANNER_FACTS') or HERE / 'data' / 'planner-facts.json').read_text())
S = F['states']
C = F['copy']
WHO = ('Felsnap', 'hunter', 'Level 34')

# The strings this round proposes (spec section 6.2); the build adds them to lib/planner/copy.ts.
PROPOSED = {
    'inGame': 'In the game',
    'pointsLeft': 'Points left',
    'spentOf': lambda spent, avail: f'{spent} of {avail} spent',
    'spentAtLevel': lambda spent, avail, level: f'{spent} of {avail} spent at level {level}',
    'linkedNote': 'Opened from a link. The game does not record the order points were spent in, so this is one way to level into it: the main tree first, each talent as soon as its row opens.',
    'levelLine': lambda name, level: f'{name} is level {level}: points after level {level} are planned ahead.',
    'decade': lambda lo: 'Level 60' if lo == 60 else f'Levels {lo} to {lo + 9}',
    'nextPoint': lambda level, name: f'Next point at {level}',
    'resetAsk': 'Clear every point in this build?',   # today a literal in PlannerToolbar.svelte (gap G3)
    'resetYes': 'Clear all points', 'resetNo': 'Keep the build', 'reset': 'Reset…',
    'scoreBasis': lambda band: f'with the {band} gear set',   # band: bandPhrase(label).toLowerCase()
    'copyLink': 'Copy link', 'copied': 'Copied', 'retry': 'Retry', 'title': 'Title', 'optional': 'Optional',
}
# SAMPLE: the live estimate (PlannerDps runs the engine in the browser; no file carries a figure).
SAMPLE_SCORE = {'guide': ('612', '± 4'), 'holy': ('418.6', '± 1.2'), 'fire': ('655', '± 5'), 'you_cap': ('171', '± 2')}


def e(t) -> str:
    return escape(str(t), quote=False)


def r(t) -> str:
    return f'{e(t)} &rsaquo;'


_art: dict[str, str] = {}


def art(bg: str) -> str:
    """The client's tree background (data/builds/<build>/trees), as the planner draws it behind the grid."""
    if bg not in _art:
        p = ROOT / 'data/builds' / F['build'] / 'trees' / f'{bg}.webp'
        _art[bg] = 'data:image/webp;base64,' + base64.b64encode(p.read_bytes()).decode() if p.exists() else ''
    return _art[bg]


STATE = {'available': 'avail', 'filled': 'part', 'maxed': 'max', 'locked': 'lock'}


def trees(st: dict, compare: bool = True, cap: bool = False, hover: int | None = None, focus: int | None = None) -> list[dict]:
    """The facts' trees as TalentEditor columns. `cap`: the character's points are all spent (decision P3: the unspent cells
    read locked, as the game's frame does with no points left)."""
    diff = set(st['compare']['differs']) if compare and st.get('compare') and st['spent'] else set()
    out = []
    for t in st['trees']:
        cells = []
        for c in t['cells']:
            state = STATE[c['state']]
            if cap and state == 'avail':
                state = 'lock'
            extra = 'hov' if c['id'] == hover else ('fv' if c['id'] == focus else '')
            cells.append((c['tier'], c['column'], TalentCell2(icon(c['icon'], 48), c['name'], c['rank'], c['max'], state, c['id'] in diff, extra)))
        out.append({'name': t['name'], 'icon_html': icon(t['icon'], 28), 'points': t['points'], 'art_src': art(t['background']),
                    'cols': t['columns'], 'tiers': t['tiers'], 'viewbox': t['viewBox'], 'connectors': [(x['d'], x['met']) for x in t['connectors']], 'cells': cells})
    return out


def tree_foot(st: dict, status: str = '', status_tone: str = '', reset_open: bool = False) -> str:
    s = f'<p class="st2{" " + status_tone if status_tone else ""}" role="status" aria-live="polite">{e(status)}</p>'
    if reset_open:
        rs = f'<span class="t-small">{PROPOSED["resetAsk"]}</span>{Button(PROPOSED["resetYes"], small=True)}{Button(PROPOSED["resetNo"], small=True)}'
    else:
        rs = Link(PROPOSED['reset']) if st['spent'] else ''
    return s + f'<span class="t-small">{e(C["tree_source"])}</span><span class="rs">{rs}</span>'


def answer(st: dict, score: tuple | None, loading: bool = False) -> str:
    """The head answer: Points left (gold) and the role's live score, then Sim this build."""
    if loading:
        ans = f'<div class="ans-box pl-ans">{Skeleton(150, 64)}{Skeleton(150, 64)}</div>'
    else:
        cap = PROPOSED['spentAtLevel'](st['spent'], st['available'], st['char_level']) if st['char_level'] else PROPOSED['spentOf'](st['spent'], st['available'])
        left = f'<div class="pl-fig">{Figure(str(st["left"]), PROPOSED["pointsLeft"], tone="gold", size="big")}<span class="t-small">{cap}</span></div>'
        if score is None:
            sc = f'<div class="pl-fig">{Figure("&mdash;", e(st["unit_label"]), size="big")}<span class="t-small">{e(C["score"]["waiting"])}</span></div>'
        else:
            v, err = score
            band = ('level 60' if st['band']['label'] == '60' else f"levels {st['band']['label']}")
            sc = f'<div class="pl-fig">{Figure(v, e(st["unit_label"]), size="big")}<span class="t-small">{err} &middot; {PROPOSED["scoreBasis"](band)}</span></div>'
        sim = Link(r(C['sim']['simThisBuild'])) if st['spent'] else ''
        ans = f'<div class="ans-box pl-ans"><div class="pl-figs">{left}{sc}</div>{sim}</div>'
    return ans


def head(st: dict, score: tuple | None, signed: bool, origin: str = '', loading: bool = False) -> str:
    """PageHeader: crest, "{Class} talents" in the class colour, the meta line; the answer: Points left and the role's score."""
    cls = st['cls']
    ans = answer(st, score, loading)
    meta = [e(st['meta'])] + ([f'<span class="gold">{e(origin)}</span>'] if origin else [])
    return PageHeader(e(st['title']), CLASS_COLOR[cls], crest_img(cls, 56), meta, ans)


def bar(st: dict, open_m: bool = False) -> str:
    picks = ClassPicker([(crest_img(c['slug'], 22), c['name'], c['slug'] == st['cls']) for c in F['classes']])
    race = sel(e(st['race']))
    end = f'<span class="t-small">{e(C["header"]["factionFollowsRace"])}</span>'
    return control_bar([('Class', picks), ('Race', race)], f'{e(st["className"])} &middot; {e(st["race"])}', end, open_m=open_m, label='Planner choices')


def notes(text: str, lead: str = '') -> str:
    if not text and not lead:
        return ''
    l = f'<b class="strong">{e(lead)}</b> ' if lead else ''
    return f'<div data-region="{REGION["notes"]}" class="pl-notes">{Notice(l + e(text), tone="neutral")}</div>'


def compare(st: dict, confirm: bool = False, loaded: str = '', readonly: bool = False) -> str:
    b, cmp = st['band'], st['compare']
    if cmp is None:
        return Section(e(b['heading']), f'<p class="t-small pl-pad">{e(b["none"])}</p>')
    s = cmp['status']
    if s['kind'] == 'matches':
        line = f'<p class="pl-line"><b class="kill">{e(cmp["line"])}</b></p>'
    else:
        line = f'<p class="pl-line"><b class="gold num">{s["count"]}</b> {e(cmp["line"])}</p>'
    beyond = f'<p class="t-small">{e(cmp["beyond"])}</p>' if cmp['beyond'] else ''
    if confirm:
        acts = f'<p class="t-small">{e(C["band"]["loadConfirmMessage"])}</p><div class="pl-acts">{Button(e(C["band"]["loadConfirmProceed"]))}{Button(e(C["band"]["loadConfirmCancel"]))}</div>'
    else:
        see = Button(e(C['band']['seeWhatDiffers'])) if s['kind'] != 'matches' else ''
        acts = f'<div class="pl-acts">{"" if readonly else Button(e(b["load"]))}{see}</div>'
    ld = f'<p class="t-small">{e(loaded)}</p>' if loaded else ''
    gear = f'<p class="pl-gear"><span class="t-small">{e(b["gear"])}</span>{Link(r(C["band"]["bestInSlotList"]))}</p>'
    return Section(e(b['heading']), f'<div class="pl-sec">{line}{beyond}{acts}{ld}{gear}</div>')


def in_game(st: dict, send: dict | None, copy_state: str = 'idle') -> str:
    """The addon handoff. send: a sendActionFor result (signed in) or None (signed out: Copy for the game is the primary)."""
    copy_label = {'idle': C['addon']['copyForGame'], 'copied': C['addon']['copiedForGame'], 'failed': C['addon']['copyForGame']}[copy_state]
    send_live = send is not None and not send.get('disabled')
    copy_btn = Button(e(copy_label), primary=not send_live).replace('class="btn', 'class="btn is-off' if st['spent'] == 0 else 'class="btn', 1)
    failed = ''
    if copy_state == 'failed':
        failed = f'<p class="t-small" role="status">{e(C["addon"]["copyFailed"])}</p>{CodeField("FSB1:hunter:5/005005201/500230230050222151", label=C["addon"]["buildCodeLabel"], short=True)}'
    body = ''
    if send is not None:
        btn = Button(e(send['label']), primary=send_live).replace('class="btn', 'class="btn is-off' if send.get('disabled') else 'class="btn', 1)
        sec = send.get('secondary') or {}
        sec_btn = Button(e(sec['label'])) if sec.get('testid') == 'planner-send-cancel' else ''
        note = ''
        if send.get('note'):
            n = send['note']
            tone = {'muted': 'muted', 'ok': 'kill', 'error': 'wipe'}[n['tone']]
            link = f' {Link(r(n["link"]["label"]))}' if n.get('link') else ''
            role = ' role="alert"' if n['tone'] == 'error' else ' role="status"'
            note = f'<p class="pl-note {tone}"{role}>{e(n["text"])}{link}</p>' + (f'<p class="t-small">{e(n["detail"])}</p>' if n.get('detail') else '')
        body += f'<div class="pl-acts">{btn}{sec_btn}</div>{note}'
        body += f'<div class="pl-acts">{copy_btn}</div>{failed}<p class="t-small">{e(C["addon"]["addonCodeHint"])}</p>'
    else:
        body += f'<div class="pl-acts">{copy_btn}</div>{failed}<p class="t-small">{e(C["addon"]["addonCodeHint"])}</p>'
        body += f'<p>{Link(r(C["cta"]["getAddonAndCompanion"]))}</p>'
    return Section(PROPOSED['inGame'], f'<div class="pl-sec">{body}</div>')


def share(st: dict, mode: str = 'idle') -> str:
    """Share this build: title, Share, Copy as text, the sim on the card. mode: idle | confirm | saved | failed | ratelimit."""
    P = C['planner']
    title = f'<label class="pl-field"><span class="t-label">{PROPOSED["title"]}</span>{Field(PROPOSED["optional"])}</label>'
    off = st['spent'] == 0
    share_btn = Button(e(P['share'])).replace('class="btn', 'class="btn is-off' if off else 'class="btn', 1)
    body = f'<div class="pl-acts pl-end">{title}{share_btn}</div>'
    body += f'<div class="pl-acts"><span class="t-small">{e(st["text_summary"])}</span>{Button(e(P["copyAsText"]), small=True)}</div>'
    body += Switch(e(C['sim']['includeSimOnCard']), on=mode != 'nosim', disabled=off)
    if mode == 'confirm':
        items = ''.join(f'<li>{e(x)}</li>' for x in st['share_fields'] + [P['shareFieldSim']])
        body += (f'<div class="pl-confirm" role="group" aria-label="{e(P["shareConfirmTitle"])}"><p class="strong">{e(P["shareConfirmTitle"])}</p><p class="t-small">{e(P["shareConfirmIntro"])}</p>'
                 f'<ul class="pl-list t-small">{items}</ul><div class="pl-acts">{Button(e(P["shareConfirmProceed"]))}{Button(e(P["shareConfirmCancel"]))}{Button(e(P["shareConfirmCopyUnsaved"]))}</div></div>')
    elif mode == 'saved':
        body += (f'<div class="pl-confirm"><div class="pl-acts"><span class="pl-url">foreversixty.gg/b/k7Qm2xa</span>{Button(PROPOSED["copyLink"], small=True)}</div>'
                 f'<p class="t-small">{e(C["sim_done"])}</p></div>')
    elif mode in ('failed', 'ratelimit'):
        msg = C['save_failed'] if mode == 'failed' else C['rate_limit']
        body += f'<div class="pl-confirm" role="alert"><p class="wipe">{e(msg)}</p><div class="pl-acts">{Button(PROPOSED["retry"], small=True)}</div></div>'
    return Section(e(P['shareTitle']), f'<div class="pl-sec">{body}</div>')


def importer(st: dict, mode: str = 'idle') -> str:
    A = C['addon']
    text = {'idle': A['importPlaceholder'], 'bad': 'FS9:nope', 'switch': 'FS1:1.60.1.70790:priest:human:00500503103/…'}[mode]
    field = CodeField(text, placeholder=mode == 'idle', bad=mode == 'bad', label=A['importTitle'])
    msg = ''
    if mode == 'bad':
        msg = f'<p class="pl-note wipe" role="alert">{e(F["notes"]["invalid"])}</p>'
    elif mode == 'switch':
        msg = f'<p class="pl-note muted" role="status">{e(A["importSwitching"])}</p>'
    imp = Button(e(A['importAction'])).replace('class="btn', 'class="btn is-off' if mode == 'idle' else 'class="btn', 1)
    body = f'{field}<div class="pl-acts">{imp}</div>{msg}<p class="t-small">{e(A["importNoExport"])} {Link(r(A["importGetAddon"]))}</p>'
    return Section(e(A['importTitle']), f'<div class="pl-sec">{body}</div>')


def pointsAt(level: int) -> int:
    return max(0, min(51, level - 9))


def order(st: dict) -> str:
    """Point order: rows of ten levels (10 to 19 ... 60), a socket per level, the next point's socket gold; past a character's
    level the points read as planned ahead."""
    pts = st['points']
    cap = st['char_level']
    top = 60 if cap is None else 60
    rows, cap_after = [], -1
    for lo in range(10, top + 1, 10):
        slots = []
        for lvl in range(lo, min(lo + 9, 60) + 1):
            i = lvl - 10
            if i < len(pts):
                p = pts[i]
                kind = 'pl' if cap is not None and lvl > cap else ''
                slots.append(OrderSlot(icon(p['icon'], 36), f'Level {lvl}: {p["name"]}', f'{p["rank"]}', kind))
            elif i == len(pts) and (cap is None or lvl <= cap):
                slots.append(OrderSlot(name=f'Level {lvl}: next point', kind='nx'))
            else:
                slots.append(OrderSlot(name=f'Level {lvl}', kind='em'))
        rows.append((PROPOSED['decade'](lo), slots))
        if cap is not None and lo <= cap <= lo + 9:
            cap_after = len(rows) - 1
    # the character's level line: shown once the build reaches the level (points past it are planned ahead, decision P4)
    cap_text = PROPOSED['levelLine'](F['felsnap']['name'], cap) if cap is not None and st['spent'] >= pointsAt(cap) else ''
    body = PointOrder(rows, cap_after, cap_text) if pts else f'<p class="t-small pl-pad">No points spent yet.</p>' + PointOrder(rows)
    return Section(e(C['planner']['pointOrderTitle']), body, meta=e(st['point_meta']))


def planner(key: str, *, signed: str = 'out', send_key: str | None = None, note: str = '', note_lead: str = '', origin: bool = False,
            share_mode: str = 'idle', import_mode: str = 'idle', copy_state: str = 'idle', reset_open: bool = False, confirm: bool = False,
            loaded: str = '', open_m: bool = False, cap: bool = False, status: str = '', status_tone: str = '', active: int | None = None,
            title_note: str = '') -> str:
    st = S[key]
    score = SAMPLE_SCORE.get(key) if st['left'] == 0 or key in SAMPLE_SCORE else None
    if key == 'you_cap':
        score = SAMPLE_SCORE['you_cap']
    lead_tree = active if active is not None else max(range(len(st['split'])), key=lambda i: st['split'][i]) if st['spent'] else 0
    plate = TalentEditor(trees(st, cap=cap), lead_tree, tree_foot(st, status, status_tone, reset_open))
    send = F['send'][send_key] if send_key else None
    main = Ordered(plate, 1, span=True) + Ordered(order(st), 6, span=True)
    rail = (Ordered(compare(st, confirm, loaded), 2) + Ordered(in_game(st, send, copy_state), 3) + Ordered(share(st, share_mode), 7)
            + Ordered(importer(st, import_mode), 8))
    body = (head(st, score, signed != 'out', st.get('origin', '') if origin else '') + bar(st, open_m) + notes(note, note_lead)
            + Page('workbench', main, rail) + PAGE_CSS)
    who = WHO if signed == 'in' else None
    return build_page(f'{st["title"]} · Forever Sixty', body, 'Tools', who,
                      f'Planner mock ({key}). Data: planner-facts.json (lib/planner/* on build {F["build"]}). SAMPLE: the live {st["unit_label"]} figure. {title_note} Spec: design/specs/2026-11-08-planner.md.')


def loading_page(failed: bool = False) -> str:
    st = S['empty']
    if failed:
        inner = (f'<div class="pl-err" role="alert"><p class="strong">{e(C["load_failed"])}</p><p class="t-small">{e(C["load_failed_detail"])}</p>'
                 f'<div class="pl-acts">{Button(PROPOSED["retry"])}</div></div>')
    else:
        inner = '<div class="pl-skel" role="status" aria-busy="true" aria-label="Loading talent data">' + ''.join(
            f'<div class="pl-skt">{Skeleton(180, 24)}{Skeleton(216, 384)}</div>' for _ in range(3)) + '</div>'
    from ws_prims import Plate
    plate = Plate(inner, 'Talent trees')
    rail_sk = lambda t: Section(t, f'<div class="pl-sec">{Skeleton(260, 20)}{Skeleton(200, 44)}</div>')
    main = Ordered(plate, 1, span=True)
    rail = Ordered(rail_sk(e(st['band']['heading'])), 2) + Ordered(rail_sk(PROPOSED['inGame']), 3)
    body = head(st, None, False, loading=not failed) + bar(st) + Page('workbench', main, rail) + PAGE_CSS
    return build_page('Warrior talents · Forever Sixty', body, 'Tools', None, f'Planner mock ({"error" if failed else "loading"}). Spec: design/specs/2026-11-08-planner.md.')


def states_page() -> str:
    g = S['guide']
    blocks = []
    def block(t, html):
        blocks.append(f'<section class="st-b"><h2 class="t-h3">{t}</h2>{html}</section>')
    t2 = g['trees'][2]
    pick = lambda name: next(c for t in g['trees'] for c in t['cells'] if c['name'] == name)
    prog = S['progress']
    pc = {c['state']: c for t in prog['trees'] for c in t['cells']}
    cells = ''.join(f'<span class="st-c">{TalentCell2(icon(c["icon"], 48), c["name"], c["rank"], c["max"], STATE[k], extra=x)}<span class="t-small">{lab}</span></span>'
                    for k, x, lab, c in [('available', '', 'Can spend (0 of n)', pc['available']), ('filled', '', 'Partial', pc['filled']), ('maxed', '', 'Maxed', pc['maxed']),
                                         ('locked', '', 'Locked', pc['locked']), ('available', 'hov', 'Hover', pc['available']), ('available', 'fv', 'Focus-visible', pc['available']),
                                         ('available', 'act', 'Active (pressed)', pc['available'])])
    cells += f'<span class="st-c">{TalentCell2(icon(pc["filled"]["icon"], 48), pc["filled"]["name"], pc["filled"]["rank"], pc["filled"]["max"], "part", True)}<span class="t-small">Differs from the band build</span></span>'
    cells += f'<span class="st-c">{TalentCell2("<span class=\"ic\"></span>", "Missing icon", 0, 3, "avail")}<span class="t-small">Icon missing: the dark socket</span></span>'
    block('Talent cell: every state (the game&rsquo;s colours, decision P2; 48 px, 44 below 768, 56 from 1760)', f'<div class="st-row tedit">{cells}</div>')
    lk = pc['locked']
    av = pc['available']
    block('Talent tooltip: hover or focus (150 ms); a locked talent names what it needs; after a band load, the band&rsquo;s rank',
          '<div class="st-row">' + TalentTip(av['name'], av['rank'], av['max'], av['now'], av['next']) + TalentTip(lk['name'], 0, lk['max'], '', lk['next'], lk['refusal'])
          + TalentTip(pc['filled']['name'], pc['filled']['rank'], pc['filled']['max'], pc['filled']['now'], pc['filled']['next'], band=F['band_tip']) + '</div>')
    block('Status line under the trees (role="status"): a refusal, then empty', '<div class="st-stack">' + ''.join(
        f'<p class="pl-note wipe">{e(x)}</p>' for x in F['refusals'].values()) + '</div>')
    sends = ''
    for k, lab in [('signed_out', 'Signed out'), ('empty', 'Signed in, nothing spent'), ('idle', 'Ready'), ('sending', 'Sending'), ('sent', 'Sent'), ('others', 'Another player&rsquo;s build'),
                   ('confirming', 'Asking first'), ('other_class', 'Another class&rsquo;s build'), ('nodevice', 'No companion yet'), ('signin', 'Session expired (401)'),
                   ('failed', 'Failed'), ('refused', 'Refused by the site')]:
        s = F['send'][k]
        sends += f'<div class="st-cell"><span class="t-label">{lab}</span>{in_game(g, None if k == "signed_out" else s)}</div>'
    block('In the game: every send state (sendActionFor on Felsnap), and Copy for the game', f'<div class="st-grid">{sends}</div>')
    block('Copy for the game: copied (3 s), then the clipboard refused', '<div class="st-grid">' + f'<div class="st-cell">{in_game(g, None, "copied")}</div><div class="st-cell">{in_game(g, None, "failed")}</div></div>')
    block('Level build: the load confirm (a build in progress), the loaded line, matches, no build for this level', '<div class="st-grid">'
          + f'<div class="st-cell">{compare(prog, confirm=True)}</div><div class="st-cell">{compare(S["you"], loaded=F["felsnap"]["loaded_line"])}</div>'
          + f'<div class="st-cell">{compare(S["holy"])}</div><div class="st-cell">{compare({**S["empty"], "compare": None, "band": {**S["empty"]["band"], "none": F["no_level_example"]}})}</div></div>')
    block('Share: the confirm, saved, failed, rate limited; nothing spent (Share off)', '<div class="st-grid">' + ''.join(f'<div class="st-cell">{share(g, m)}</div>' for m in ('confirm', 'saved', 'failed', 'ratelimit'))
          + f'<div class="st-cell">{share(S["empty"])}</div></div>')
    block('Import from addon: empty, a code for another class (the planner switches), an unreadable code', '<div class="st-grid">' + ''.join(f'<div class="st-cell">{importer(g, m)}</div>' for m in ('idle', 'switch', 'bad')) + '</div>')
    block('Head answer: unfinished build, finished, loading (reserved), on a phone the estimate is asked for, the estimate failed', '<div class="st-grid">'
          + f'<div class="st-cell">{answer(S["progress"], None)}</div><div class="st-cell">{answer(g, SAMPLE_SCORE["guide"])}</div><div class="st-cell">{answer(g, None, loading=True)}</div>'
          + f'<div class="st-cell"><div class="pl-fig">{Figure("&mdash;", "DPS", size="big")}{Button(e(C["sim"]["plannerDpsShow"]), small=True)}<span class="t-small">{e(C["sim"]["plannerDpsShowNote"])}</span></div></div>'
          + f'<div class="st-cell"><div class="pl-fig">{Figure("&mdash;", "DPS", size="big")}<span class="t-small">{e(g["unit_failed"])}</span></div></div></div>')
    block('Plate foot: Reset asks first', f'<div class="st-foot">{tree_foot(g, reset_open=True)}</div>')
    block('Phone tree tabs (below 768): Survival open', TalentEditor(trees(g), 2, '')[:0] + f'<div class="st-tabs">{TalentEditor(trees(g), 2)}</div>')
    body = '<div class="pc" style="padding-top:24px"><h1 class="t-title">Planner: element states</h1></div>' + Page('ledger', f'<div class="statesheet">{"".join(blocks)}</div>') + PAGE_CSS
    return build_page('Planner states', body, 'Tools', WHO, 'Planner states sheet; data as the planner mocks.')


PAGE_CSS = '''<style>
/* page-level layout only (no surfaces): the answer's figures, rail section spacing, rows of actions */
.pl-ans{gap:10px}.pl-figs{display:grid;grid-template-columns:repeat(2,minmax(0,1fr));gap:24px}.pl-fig{display:flex;flex-direction:column;gap:2px;min-width:0}
@media(min-width:900px){.pl-figs{display:flex;gap:32px}.pl-fig{max-width:220px}}
@media(min-width:900px){.pl-figs{justify-content:flex-end}.pl-fig{align-items:flex-end}}
.pl-notes .notice{padding-top:4px;color:var(--text)}
.pl-sec{display:flex;flex-direction:column;gap:12px;padding-top:14px}.pl-pad{padding-top:14px}
.pl-acts{display:flex;flex-wrap:wrap;gap:10px;align-items:center}.pl-acts.pl-end{align-items:flex-end}
.pl-line{font-size:15px;line-height:22px}.pl-note{font-size:14px;line-height:20px}.pl-note.muted{color:var(--muted)}
.pl-gear{display:flex;justify-content:space-between;gap:12px;align-items:baseline;padding-top:10px;border-top:1px solid var(--line-soft)}
.pl-field{display:flex;flex-direction:column;gap:6px;flex:1;min-width:180px}.pl-field .field{height:44px}
.pl-confirm{display:flex;flex-direction:column;gap:10px;padding:12px 0;border-top:1px solid var(--line-soft)}
.pl-list{list-style:disc;padding-left:20px}.pl-url{font:600 14px/20px var(--font-body);color:var(--gold);overflow-wrap:anywhere}
.pl-skel{display:grid;grid-template-columns:repeat(3,minmax(0,1fr));gap:24px;padding:16px 20px;min-height:520px}.pl-skt{display:flex;flex-direction:column;align-items:center;gap:14px}
@media(max-width:767px){.pl-skel{grid-template-columns:minmax(0,1fr)}.pl-skt+.pl-skt{display:none}}
.pl-err{display:flex;flex-direction:column;gap:8px;padding:20px;min-height:520px}
@media(max-width:767px){.cbar .pc{padding-block:0;min-height:0}}
.statesheet{display:grid;gap:40px}.st-b>h2{margin-bottom:12px}.st-row{display:flex;flex-wrap:wrap;gap:24px;align-items:flex-start}
.st-c{display:flex;flex-direction:column;align-items:center;gap:10px;width:120px;text-align:center}
.st-grid{display:grid;gap:24px 40px;grid-template-columns:repeat(auto-fill,minmax(320px,1fr))}.st-cell{display:flex;flex-direction:column;gap:6px;min-width:0}
.st-stack{display:flex;flex-direction:column;gap:6px}.st-foot{max-width:900px}.st-foot .tedit-f{border-top:0;padding-inline:0}
.st-tabs{max-width:390px}.st-tabs .tedit-tabs{display:flex}.st-tabs .tedit-tree:not(.on){display:none}.st-tabs .tedit-cols{grid-template-columns:minmax(0,1fr)}.st-tabs .tedit-tree{border-left:0!important}
</style>'''


PAGES = {
    'web-planner': lambda: planner('empty'),
    'web-planner-progress': lambda: planner('progress'),
    'web-planner-guide': lambda: planner('guide', note=PROPOSED['linkedNote'], origin=True),
    'web-planner-holy': lambda: planner('holy', note=PROPOSED['linkedNote'], origin=True),
    'web-planner-fire': lambda: planner('fire', note=PROPOSED['linkedNote'], origin=True),
    'web-planner-you': lambda: planner('you', signed='in', send_key='idle', note=F['felsnap']['loaded_line']),
    'web-planner-you-cap': lambda: planner('you_cap', signed='in', send_key='idle', cap=True),
    'web-planner-you-empty': lambda: planner('you_empty', signed='in', send_key='empty'),
    'web-planner-other': lambda: planner('holy', signed='in', send_key='other_class', note=PROPOSED['linkedNote'], origin=True),
    'web-planner-invalid': lambda: planner('invalid', note=F['notes']['invalid']),
    'web-planner-share': lambda: planner('guide', share_mode='confirm', note=PROPOSED['linkedNote'], origin=True),
    'web-planner-loading': lambda: loading_page(),
    'web-planner-error': lambda: loading_page(True),
    'web-planner-bar': lambda: planner('progress', open_m=True),
    'web-planner-states': states_page,
}


def compare_sheets(before: Path) -> None:
    import ws_compose
    ws_compose.BEFORE = before
    for name, live, mock in [('planner', 'planner-out', 'web-planner'), ('planner-guide', 'planner-guide', 'web-planner-guide'), ('planner-you', 'planner-in', 'web-planner-you')]:
        for w, sc in ((1440, .5), (1920, .42), (2000, .4), (390, .8)):
            if (before / f'{live}-{w}.png').exists():
                print(ws_compose.compose(name, live, mock, w, sc, max_h=6000))


if __name__ == '__main__':
    if sys.argv[1:2] == ['--list']:
        print('\n'.join(PAGES))
    elif sys.argv[1:2] == ['compare']:
        compare_sheets(Path(sys.argv[2]))
    else:
        only = sys.argv[1:]
        bad = [n for n in only if n not in PAGES]
        if bad:
            sys.exit(f'ws_planner.py: no such mock: {" ".join(bad)} (see --list)')
        from ws_lib import write
        for name, fn in PAGES.items():
            if not only or name in only:
                print(write(f'{name}.html', fn()))
