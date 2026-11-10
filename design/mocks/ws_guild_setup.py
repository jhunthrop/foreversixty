"""Guild mocks, the officer's setup checklist after a claim (/guild/<region>/<ruleset>/<slug>/setup), the claim page that leads to it,
and the settings tab's way back to it. Composed only from ws_prims primitives; the guild is guild_data's. Registered in
ws_guild2.all_pages().

  setup                  right after the claim: the invite link made at claim time, every other step to do.
  setup-mixed            later: five steps done, the linked role skipped, recruiting and the raid schedule to do.
  claim-done             the claim page after a claim: the old Next buttons become a "Set up <guild>" panel and Start setup.
  settings-setup-row     the settings tab with one slim row at the top of the main column: Setup, 5 of 8 done, Continue setup.

Each step row: a status mark (done a teal check, to do a gold ring with the number, skipped a muted dash), the title and its Tag,
one line of live facts, and for a step to do its controls under the facts and a quiet Skip at the right. A done step is one line
with Change; a skipped step reads Skipped with Undo.
"""
from dataclasses import dataclass

import guild_data as gd
import guild_settings_fixtures as GS
import ws_guild_directory as GD
import ws_guild_settings as WS
from guild_data import ALLIANCE, GUILD
from ws_guild2 import _act, _view, hdr
from ws_prims import (Button, Code, Figure, Link, Meter, Page, PageHeader, Plate, Row, Rows, Section, Segmented, Select, Tag, build_page)

WHO = (gd.VIEWER, 'hunter', 'Beast Mastery Hunter &middot; 60')
MEMBERS = gd.RAIDERS
# The guild's ranks and their member counts (rank index, members); ranks 0 to OFFICER_MAX count as officers.
RANKS = [(0, 1), (1, 2), (2, 5), (3, 12), (4, 4)]
OFFICER_MAX = 1
OFFICERS = sum(n for r, n in RANKS if r <= OFFICER_MAX)
JOIN_LINK = 'foreversixty.gg/j/7Hk2Qa'
JOINED_BY_LINK = 14
SYNCED = 12
SYNC_DAYS = 14
RECRUIT_CHANNEL = '#recruitment'
NEXT_RAID = "Onyxia's Lair Wed 16 Dec, 8:00 pm"
PLANNED = 2
RENEWS = '6 Jan 2027'
INVITE_PASTE = f'Join {GUILD} on Forever Sixty and sign up for raids: {JOIN_LINK}'
ADDON_PASTE = 'Get the Forever Sixty addon and companion so your gear and the roster sync: foreversixty.gg/setup'
ROLE_PATH = 'Give Premium members a role in your server: Server Settings, Roles, the role, Links, Add requirement, Forever Sixty Premium.'
BOT_LINE = 'Raid posts with In, Late, Bench and Out buttons, /raid, /signup, /bis and /char. Free.'

FREE = ['The Discord bot and raid posts with sign-up buttons', 'Sign-ups and attendance from them', 'The recruiting channel',
        'Member commands /raid, /signup, /bis, /char, /help', 'The raid calendar']
GUILD_ADDS = ['Comp targets', 'Discord pings for non-signers', 'Boss assignments', 'Log matching and the raid-night report', '1 Premium seat',
              'Up to 25 members']
RAID_GUILD_ADDS = ['Loot council', 'The readiness board', '3 Premium seats', 'Any size']

DONE, TODO, SKIPPED = 'done', 'todo', 'skipped'
STATE_TAG = {DONE: ('Done', 'kill'), TODO: ('To do', 'gold'), SKIPPED: ('Skipped', 'muted')}
STATE_GLYPH = {DONE: '&#10003;', SKIPPED: '&#8211;'}

_CSS = '''<style>
.gs .row{align-items:start}.gs .row>span:first-child{padding-top:1px}
.gs .row .ttl{display:flex;flex-wrap:wrap;align-items:center;gap:6px 10px}
.gs .row .ttl .gs-opt{font:400 13px/20px var(--font-body);color:var(--muted)}
.gs .row .ttl .gs-fact{font:400 14px/20px var(--font-body);color:var(--muted)}
.gs .row .sub{font-size:14px;line-height:20px;margin-top:4px}
.gs .row .trail{align-self:start;min-height:28px}
@media(min-width:601px){.gs .row .foot{grid-column:2/-1;padding-top:10px}}
.gs .row.gs-one{min-height:52px;align-items:center}.gs .row.gs-one>span:first-child{padding-top:0}.gs .row.gs-one .trail{align-self:center;min-height:0}
.smk.gs-done{background:var(--teal);box-shadow:inset 0 0 0 1px var(--teal-edge);color:var(--strong)}
.smk.gs-todo{box-shadow:inset 0 0 0 1.5px var(--gold);color:var(--gold)}
.smk.gs-skip{box-shadow:inset 0 0 0 1.5px var(--line);color:var(--muted)}
.smk.gs-xs{width:16px;height:16px;font-size:10px;line-height:16px}
.gs-ranks{display:flex;flex-wrap:wrap;gap:4px 14px;font-variant-numeric:tabular-nums}.gs-ranks .on{color:var(--strong);font-weight:600}
.gs-subs{display:flex;flex-wrap:wrap;gap:6px 14px;margin-top:6px;font-size:13px;line-height:20px;color:var(--text)}
.gs-subs>span{display:inline-flex;align-items:center;gap:6px}
.gs-ctl{display:flex;flex-wrap:wrap;align-items:center;gap:8px 12px;width:100%}
.gs-ctl label{display:inline-flex;align-items:center;gap:8px;font-size:14px;line-height:20px;color:var(--text)}
.gs-ctl input.field.gs-num{width:64px;height:40px;color:var(--strong);font-weight:600}
.gs-ctl input.field.gs-link{width:240px;height:40px;color:var(--strong)}
.gs-ctl .link{display:inline-flex;align-items:center;min-height:40px}
.gs-paste{height:auto;min-height:40px;padding:8px 12px;flex:1 1 320px;max-width:560px;white-space:normal}.gs-paste code.c{font-size:13px;line-height:20px;font-weight:500}
.gs-paste code.c{display:block;white-space:pre-wrap;overflow-wrap:anywhere}
@media(max-width:600px){.gs .row{grid-template-columns:auto minmax(0,1fr) auto}.gs .row .trail{grid-column:auto;justify-self:end;margin-top:0}.gs-paste{flex-basis:100%;max-width:none}}
.gs-lbl{font-size:13px;line-height:20px;color:var(--muted);width:100%}
.gs-bar{display:flex;align-items:center;gap:12px;margin-top:6px;max-width:320px}.gs-bar .meter{flex:1}
a.link.gs-quiet{font:400 13px/20px var(--font-body);color:var(--muted);text-decoration:underline}
.gs-plan{display:flex;flex-direction:column;gap:16px}
.gs-plan h3{font:700 14px/20px var(--font-body);color:var(--strong);margin:0 0 4px}
.gs-plan ul{margin:0;padding-left:18px;display:flex;flex-direction:column;gap:2px;font-size:13px;line-height:20px;color:var(--text);list-style:disc}
.gs-plan .gs-now{font-size:14px;line-height:20px;color:var(--text)}.gs-plan .gs-now b{color:var(--strong)}
.gs-plan .link{display:inline-flex;align-items:center;min-height:44px}
.gs-ol{list-style:none;margin:0;padding:0;display:grid;gap:0 28px;grid-template-columns:1fr}
@media(min-width:768px){.gs-ol{grid-template-columns:1fr 1fr;grid-auto-flow:column;grid-template-rows:repeat(4,auto)}}
.gs-ol li{display:flex;align-items:center;gap:10px;min-height:40px;border-bottom:1px solid var(--line-soft);font-size:14px;line-height:20px;color:var(--strong)}
.gs-ol li .st{margin-left:auto;font-size:13px;color:var(--muted)}
.gs-start{display:flex;flex-wrap:wrap;align-items:center;gap:12px 20px;padding-top:16px}
.gs-slim .row .ttl .gs-fact{font-weight:400;color:var(--muted)}
</style>'''


@dataclass(frozen=True)
class Step:
    """One setup step: its title, its live facts and controls for each state it can be in."""
    n: int
    title: str
    state: str
    facts: str = ''
    controls: str = ''
    done_facts: str = ''
    optional: bool = False


def mark(n: int, state: str, xs: bool = False) -> str:
    cls = {DONE: 'gs-done', TODO: 'gs-todo', SKIPPED: 'gs-skip'}[state]
    word = STATE_TAG[state][0]
    glyph = STATE_GLYPH.get(state, '' if xs else str(n))
    label = f' role="img" aria-label="Step {n}: {word}"' if not xs else ' aria-hidden="true"'
    return f'<span class="smk {cls}{" gs-xs" if xs else ""}"{label}>{glyph}</span>'


def quiet(label: str) -> str:
    return Link(label).replace('class="link"', 'class="link gs-quiet"', 1)


def fact(text: str) -> str:
    return f'<span class="gs-fact">{text}</span>'


def subs(items: list[tuple[str, bool, str]]) -> str:
    """Sub-items as small inline checks: (label, done, tag html)."""
    out = ''.join(f'<span>{mark(0, DONE if ok else TODO, xs=True)}{label}{tag}</span>' for label, ok, tag in items)
    return f'<div class="gs-subs">{out}</div>'


def ctl(html: str) -> str:
    return f'<div class="gs-ctl">{html}</div>'


def step_row(s: Step) -> str:
    tag = Tag(*STATE_TAG[s.state])
    opt = '<span class="gs-opt">Optional</span>' if s.optional else ''
    if s.state == TODO:
        return Row(f'{s.title}{opt}{tag}', s.facts, lead=mark(s.n, s.state), trail=quiet('Skip'), link=False, foot=s.controls)
    line = fact(s.done_facts) if s.state == DONE and s.done_facts else ''
    row = Row(f'{s.title}{opt}{tag}{line}', lead=mark(s.n, s.state), trail=quiet('Change' if s.state == DONE else 'Undo'), link=False)
    return row.replace('class="row', 'class="row gs-one', 1)


def copy_box(text: str, label: str) -> str:
    return (f'<span class="gs-lbl">{label}</span><div class="field gs-paste">{Code(text)}</div>'
            f'{_act("Copy", small=False)}')


# ---- the eight steps ------------------------------------------------------------------------------------------------------------
def ranks_step(state: str) -> Step:
    chips = ''.join(f'<span class="{"on" if r <= OFFICER_MAX else ""}">Rank {r} &middot; {n}</span>' for r, n in RANKS)
    facts = f'<div class="gs-ranks">{chips}</div><div>Ranks 0 to {OFFICER_MAX} count as officers: {OFFICERS} members</div>'
    controls = ctl(f'<label>Officers are ranks 0 to <input class="field gs-num" type="number" min="0" max="9" value="{OFFICER_MAX}" '
                   f'aria-label="Highest officer rank"></label>{_act("Save", small=False)}{Button(f"Keep 0 to {OFFICER_MAX}", tone="ghost", act=True)}')
    return Step(1, 'Officer ranks', state, facts, controls, f'Ranks 0 to {OFFICER_MAX} &middot; {OFFICERS} officers')


def bot_step(state: str) -> Step:
    controls = ctl(f'{_act("Add the Discord bot", primary=True, small=False)}<span class="gs-lbl">{BOT_LINE}</span>')
    done = f'Posts in #{GS.BOT_CHANNEL} on {GUILD} &middot; {GS.BOT_LINKED} of {MEMBERS} members linked Discord'
    return Step(2, 'Discord bot', state, 'Not in your server yet', controls, done)


def recruit_step(state: str, set_up: bool) -> Step:
    wanted = ', '.join(spec for _, _, spec, _ in gd.RECRUIT_WANTED)
    facts = f'Open &middot; 1 night &middot; wanted: {wanted}' if set_up else 'Closed'
    done = [('Pitch', set_up), ('Nights', set_up), ('Wanted', set_up), ('Recruit role', False), ('Recruit channel', False)]
    controls = ctl(f'<label>Post applications in {Select(RECRUIT_CHANNEL)}</label>{_act("Save", small=False)}{Link("Edit the pitch &rsaquo;")}')
    return Step(3, 'Recruiting', state, facts + subs([(lab, ok, '') for lab, ok in done]), controls, facts)


def schedule_step(state: str, planned: bool) -> Step:
    facts = f'{gd.ZONE} &middot; next: {NEXT_RAID} &middot; {PLANNED} planned' if planned else f'{gd.ZONE} &middot; nothing planned'
    items = [('Timezone', True, ''), ('First raid', planned, ''), ('Comp targets', False, ' ' + Tag('Guild', 'muted'))]
    controls = ctl(f'{_act("Plan a raid", small=False)}<label>Timezone {Select(gd.ZONE)}</label>')
    return Step(4, 'Raid schedule', state, facts + subs(items), controls, facts)


def invite_step(state: str) -> Step:
    controls = ctl(f'<input class="field gs-link" readonly value="{JOIN_LINK}" aria-label="Guild invite link">{_act("Copy", small=False)}'
                   f'{copy_box(INVITE_PASTE, "Paste in guild chat")}{_act("I posted it", small=False)}')
    done = f'Made by {gd.VIEWER} &middot; {JOINED_BY_LINK} joined by it'
    return Step(5, 'Guild invite link', state, 'Made at claim &middot; nobody joined by it yet', controls, done)


def addon_step(state: str) -> Step:
    line = f'{SYNCED} of {MEMBERS} members synced in the last {SYNC_DAYS} days'
    bar = f'<div class="gs-bar">{Meter(SYNCED / MEMBERS, "", f"{SYNCED} of {MEMBERS} synced")}</div>'
    controls = ctl(f'{copy_box(ADDON_PASTE, "Paste in guild chat")}{_act("I posted it", small=False)}')
    return Step(6, 'Addon and companion', state, line + bar, controls, line)


def visibility_step(state: str) -> Step:
    seg = Segmented([('Public', True), ('Unlisted', False), ('Guild only', False)]).replace('role="radiogroup"', 'role="radiogroup" aria-label="Report visibility"', 1)
    controls = ctl(f'{seg}{Button("Keep Public", tone="ghost", act=True)}')
    return Step(7, 'Report visibility', state, 'Raid reports are Public', controls, 'Public')


def role_step(state: str) -> Step:
    return Step(8, 'Discord linked role', state, ROLE_PATH, ctl(_act('Done', small=False)), 'Premium role linked', optional=True)


def fresh_steps() -> list[Step]:
    return [ranks_step(TODO), bot_step(TODO), recruit_step(TODO, False), schedule_step(TODO, False), invite_step(DONE), addon_step(TODO),
            visibility_step(TODO), role_step(TODO)]


def mixed_steps() -> list[Step]:
    return [ranks_step(DONE), bot_step(DONE), recruit_step(TODO, True), schedule_step(TODO, True), invite_step(DONE), addon_step(DONE),
            visibility_step(DONE), role_step(SKIPPED)]


def count(steps: list[Step], state: str) -> int:
    return sum(1 for s in steps if s.state == state)


def done_line(steps: list[Step]) -> str:
    return f'{count(steps, DONE)} of {len(steps)} done'


# ---- the page -------------------------------------------------------------------------------------------------------------------
def setup_section(steps: list[Step]) -> str:
    todo, skipped = count(steps, TODO), count(steps, SKIPPED)
    meta = ' &middot; '.join([f'{todo} to do'] + ([f'{skipped} skipped'] if skipped else []))
    return Section('Setup', f'<div class="gs">{Rows([step_row(s) for s in steps], open_=True)}</div>', meta=meta, panel=True, frame='framed')


def bullets(title: str, items: list[str]) -> str:
    return f'<div><h3>{title}</h3><ul>{"".join(f"<li>{i}</li>" for i in items)}</ul></div>'


def plan_section(plan: str) -> str:
    now = f'<p class="gs-now" style="margin:0">{plan}</p>'
    body = (f'<div class="gs-plan">{now}{bullets("Free for every claimed guild", FREE)}{bullets("Guild adds", GUILD_ADDS)}'
            f'{bullets("Raid guild adds", RAID_GUILD_ADDS)}<span>{Link("Compare plans &rsaquo;")}</span></div>')
    return Section('Plan', body, panel=True)


def setup_page(title: str, steps: list[Step], plan: str):
    n = count(steps, DONE)
    answer = Figure(str(n), unit=f'of {len(steps)} done')
    head = PageHeader(GUILD, ALLIANCE, crumb=f'<a href="#">&lsaquo; {GUILD}</a>', meta=['Setup'], answer=answer)
    body = head + Page('split', setup_section(steps), plan_section(plan), panels=True)
    return build_page(title, _view(body) + _CSS, 'Guild', WHO)


def setup_fresh():
    return setup_page('Guild setup', fresh_steps(), '<b>No plan</b>')


def setup_mixed():
    return setup_page('Guild setup, part done', mixed_steps(), f'<b>Guild</b> &middot; renews {RENEWS}')


# ---- the claim page after a claim -----------------------------------------------------------------------------------------------
def setup_list(steps: list[Step]) -> str:
    items = ''.join(f'<li>{mark(s.n, s.state)}{s.title}<span class="st">{STATE_TAG[s.state][0]}</span></li>' for s in steps)
    # ClaimNextSteps: Start setup and the roster's approvals are links (Button href).
    start = (f'<div class="gs-start">{Button("Start setup", primary=True)}'
             f'{Button("Approve your officers and members", tone="ghost")}</div>')
    return Section(f'Set up {GUILD}', f'<ol class="gs-ol">{items}</ol>{start}', meta=done_line(steps), panel=True)


def claim_done():
    state = (f'<div class="dir-col"><p style="font-size:15px">You hold {GUILD}.</p>'
             f'{GD._act("Release claim").replace("class=\"btn\"", "class=\"backlink\"")}</div>')
    rules = '<ul class="dir-list">' + ''.join(f'<li>{r}</li>' for r in GD._CLAIM_RULES) + '</ul>'
    main = Section('Claim', state, panel=True, frame='framed') + setup_list(fresh_steps())
    rail = Section('How claiming works', rules, panel=True)
    head = PageHeader(f'{GUILD} &middot; your claim', crumb=f'<a href="#">&lsaquo; {GUILD}</a>')
    return build_page('Claim', GD._view(head + Page('split', main, rail, panels=True)) + GD._CSS + _CSS, 'Guild', WHO)


# ---- the settings tab with the way back to setup --------------------------------------------------------------------------------
def setup_row(steps: list[Step]) -> str:
    row = Row(f'Setup {fact("&middot; " + done_line(steps))}', trail=Link('Continue setup &rsaquo;'), link=False, slim=True)
    return f'<div class="gs-slim">{Plate(Rows([row]), "Setup")}</div>'


def settings_with_setup_row():
    main = setup_row(mixed_steps()) + WS.plan_section() + WS.discord_section()
    rail = WS.members_section() + WS.records_section(crest=True)
    head = hdr(tabs_active='Settings') + '<h2 class="sp-sr">Guild settings</h2>'
    body = _view(head + Page('split', main, rail, panels=True) + WS.NOTICE_LINE + WS._CSS)
    return build_page('Guild settings', body + _CSS, 'Guild', WHO)


PAGES: dict = {'web-guild-setup': setup_fresh, 'web-guild-setup-mixed': setup_mixed, 'web-guild-claim-done': claim_done,
               'web-guild-settings-setup-row': settings_with_setup_row}
