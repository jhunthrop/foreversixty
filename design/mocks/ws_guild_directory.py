"""Guild mocks, the guild landing (/guild), join and claim (GUILD-ALL lane, 2026-11-07). Composed only from ws_prims primitives; the data
is guild_directory_fixtures.py's, which guild_fixtures.py writes for the build. Registered in ws_guild2.all_pages().

- Landing (web layout spec 2026-11-05 section 6.1): PageHeader "Guilds" with the shipped lead, then Page split. Main: Open a guild (a
  FormRow: name Field, region and ruleset Selects, one primary Button) and Top guilds (the page's one Plate, ranked link Rows with a
  Meter against the leader). Rail: Your guilds (signed in with guilds: open link Rows with the rank Tag), then Run a guild? (signed
  out: the line and the sign-in Button).
- Join: the guild's PageHeader (faction crest, the faction-colour name, the realm line), then the form measure with one Section, Your
  invite: who invited you, the character you join as (an open Row), what joining does, Join and Not now.
- Claim (the holder's read): the back link, PageHeader "Olympus XXVII · your claim", Page split. Main: Claim (the state sentence and
  Release claim) and Next (the four Buttons, the plan-gate line). Rail: How claiming works.
"""
import guild_data as gd
import guild_directory_fixtures as D
from guild_data import ALLIANCE, C, GUILD
from ws_prims import (Button, Emblem, Meter, Page, PageForm, PageHeader, Plate, Row, Rows, Section, Tag, build_page, crest_img)

_OFFICER_LINE = 'Sign in and import your character, then open the guild page and choose Claim this guild.'
_PROSE = 'style="font-size:14px;line-height:22px"'
# guildClaimCopy.rules (web/src/lib/guild/copy.ts)
_CLAIM_RULES = ['Claiming needs a Battle.net-linked account.', 'You may attempt one claim every 30 days.',
                'You may hold only one claimed guild at a time.', 'A claim can be contested.']

# The built page's roles the ws_prims helpers do not draw: a <form> of labelled inputs and selects, an action <button>, a list.
_CSS = ('<style>.sr-only{position:absolute;width:1px;height:1px;padding:0;margin:-1px;overflow:hidden;clip:rect(0,0,0,0);white-space:nowrap}'
        'form.formrow input.field,form.formrow select.field{font:inherit;color:var(--strong)}form.formrow select.field{width:120px}'
        '.dir-list{margin:0;padding-left:20px;display:flex;flex-direction:column;gap:8px;font-size:14px;line-height:22px;list-style:disc}'
        '.dir-col{display:flex;flex-direction:column;align-items:flex-start;gap:16px}.dir-btns{display:flex;flex-wrap:wrap;gap:12px}'
        '.dir-trail{display:flex;align-items:center;gap:12px}.dir-trail .meter{width:96px}.dir-trail b{width:72px;text-align:right;font-weight:600;font-size:14px}'
        '.backlink{display:inline-flex;align-items:center;min-height:44px;font-size:13px;color:var(--muted);text-decoration:underline}'
        'button.backlink{background:none;border:0;padding:0;font-family:inherit;cursor:pointer}</style>')


def _act(label: str, primary: bool = False, tone: str = '', submit: bool = False) -> str:
    """The system Button as the build renders an action (a <button>; ws_prims draws an <a>)."""
    html = Button(label, primary, tone=tone)
    kind = 'submit' if submit else 'button'
    return html.replace('<a ', f'<button type="{kind}" ', 1).replace(' href="#"', '', 1).replace('</a>', '</button>')


def _body(html: str) -> str:
    """A Section body as the built Section spaces it: 12 px under the head row."""
    return f'<div style="padding-top:12px">{html}</div>'


def _view(html: str) -> str:
    """GuildShell keeps every guild view one screen tall (GUILD_VIEW_MIN); the guild route's main pads 40 px under it."""
    return f'<div style="min-height:calc(100vh - 81px);margin-bottom:40px">{html}</div>'


def _select(label: str, value: str, ident: str) -> str:
    return (f'<label class="f" for="{ident}"><span class="t-label">{label}</span>'
            f'<select class="field sel" id="{ident}"><option selected>{value}</option></select></label>')


def lookup(ruleset: str) -> str:
    name = '<label class="f grow"><span class="t-label">Guild name</span><input class="field grow" placeholder="Guild name"></label>'
    form = (f'<form class="formrow" aria-label="Open a guild">{name}{_select("Region", "US", "r")}{_select("Ruleset", ruleset, "s")}'
            f'<span>{_act("Open guild", primary=True, submit=True)}</span></form>')
    return Section('Open a guild', form, panel=True, frame='framed')


def top_guilds() -> str:
    lead = D.TOP[0][3]
    rows = [Row(n, f'{r.upper()} {D.RULESET_WORD[rs]}', rank=str(i + 1),
                trail=f'<span class="dir-trail">{Meter(v / lead, "", f"{n}: {v} bosses" if v != 1 else f"{n}: 1 boss")}<b>{v} {"boss" if v == 1 else "bosses"}</b></span>')
            for i, (n, r, rs, v) in enumerate(D.TOP)]
    return Section('Top guilds', Plate(Rows(rows), 'Top guilds'), boxed=True)


def landing(signed_in: bool):
    """Signed in, the lookup starts on the account's first guild's region and ruleset (US PvP)."""
    rail = ''
    if signed_in:
        rows = [Row(n, f'{r.upper()} {D.RULESET_WORD[rs]}', trail=Tag(rank.title(), 'muted')) for n, r, rs, rank in D.MINE]
        rail += Section('Your guilds', Rows(rows, open_=True), panel=True)
    sign_in = '' if signed_in else Button('Sign in with Battle.net')
    rail += Section('Run a guild?', f'<div class="dir-col"><p {_PROSE}>{_OFFICER_LINE}</p>{sign_in}</div>', panel=True)
    body = PageHeader('Guilds') + Page('split', lookup('PvP' if signed_in else 'Normal') + top_guilds(), rail, panels=True)
    who = (D.ALT[0], D.ALT[1], f'{D.ALT[1].title()} &middot; {D.ALT[2]}') if signed_in else None
    return build_page('Guilds', _view(body) + _CSS, 'Guild', who)


def join():
    name, cls, race, level = D.JOINER
    inviter = (f'<p style="display:inline-flex;align-items:center;gap:6px;font-size:15px">{crest_img(D._class_of(D.INVITER), 24)}'
               f'<b style="color:{C[D._class_of(D.INVITER)]}">{D.INVITER}</b> invited you to join.</p>')
    you = f'<div style="width:100%">{Rows([Row(f"You join as {name}", f"Level {level} {race} {cls.title()}", lead=crest_img(cls, 32), title_color=C[cls], link=False)], open_=True)}</div>'
    lines = ('<div style="display:flex;flex-direction:column;gap:4px"><p class="t-small">Joining by this link verifies you at once. No officer has to approve you.</p>'
             '<p class="t-small">Joining puts your characters on its roster and shares your gear with its officers; you can change that on your account.</p></div>')
    buttons = f'<div class="dir-btns">{_act(f"Join {GUILD}", primary=True)}{Button("Not now", tone="ghost")}</div>'
    section = Section('Your invite', f'<div class="dir-col">{inviter}{you}{lines}{buttons}</div>', panel=True, frame='framed')
    head = PageHeader(GUILD, ALLIANCE, Emblem(gd.FACTION, 64), meta=[f'PvP US &middot; {gd.RAIDERS} members on Forever Sixty'])
    return build_page('Join a guild', _view(head + PageForm(section)) + _CSS, 'Guild', (name, cls, f'{cls.title()} &middot; {level}'))


def claim():
    state = (f'<div class="dir-col"><p style="font-size:15px">You hold {GUILD}.</p>'
             f'{_act("Release claim").replace("class=\"btn\"", "class=\"backlink\"")}</div>')
    nxt = (f'<div class="dir-btns">{Button("Approve your officers and members", primary=True)}{Button("Make an invite link")}'
           f'{Button("Plan an event")}{Button("Add the Discord bot")}</div>'
           '<p class="t-small" style="padding-top:12px">Events and the Discord bot are part of the Guild and Raid guild plans.</p>')
    rules = '<ul class="dir-list">' + ''.join(f'<li>{r}</li>' for r in _CLAIM_RULES) + '</ul>'
    main = Section('Claim', state, panel=True, frame='framed') + Section('Next', nxt, panel=True)
    rail = Section('How claiming works', rules, panel=True)
    head = f'<div class="pc"><a class="backlink" href="#">&lsaquo; {GUILD}</a></div>' + PageHeader(f'{GUILD} &middot; your claim')
    return build_page('Claim', _view(head + Page('split', main, rail, panels=True)) + _CSS, 'Guild', (D.CLAIMANT, 'hunter', 'Beast Mastery Hunter &middot; 60'))


PAGES: dict = {'web-guild-landing-out': lambda: landing(False), 'web-guild-landing-in': lambda: landing(True),
               'web-guild-join': join, 'web-guild-claim': claim}
