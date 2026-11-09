"""Guild mocks, the Settings tab and the guild settings page (GUILD-ALL lane, 2026-11-07). Composed only from ws_prims primitives; the data is
guild_settings_fixtures.py's, which guild_fixtures.py writes for the build. Registered in ws_guild2.all_pages().

  settings       the officer's Settings tab on the guild page: the plan (a founder Raid guild billed to Dunmaro, one seat of three given) and
                 Discord in the main column; Members and access, then Records and crest in the rail. No cards: open Sections, rows on hairlines.
  settings-bot   the same tab with the Forever Sixty bot installed in the guild's Discord server (no webhook): where raids post, the linked
                 members, Send test message, Change channel and Remove.
  settings-page  the standalone /settings route for the same guild before a plan: its own head (the guild's name in the faction colour, one
                 meta line), the plan chooser in the plan's place, Records without the crest (the crest lives on the guild page's tab).
"""
from datetime import datetime

import guild_settings_fixtures as S
from guild_data import ALLIANCE, C, GUILD
from ws_guild2 import O, _act, _view, hdr
from ws_prims import Button, Emblem, Link, Page, PageHeader, Row, Rows, Section, Segmented, Tag, build_page, crest_img

# The built copy (lib/billing/guild-billing-copy, lib/billing/seats-copy, lib/guild/copy, lib/guild/settings-view), verbatim.
WHO_PAYS_HELP = 'Hand the plan to another officer, for example when the guild leader changes. The founder price goes with it.'
SEATS_HELP = 'A seat gives one verified member Premium while the plan lasts. A seat can move once every 30 days.'
INVITE_HELP = (f'Anyone with this link joins {GUILD} as a verified member: no officer has to approve them, and they can sign up for raids and see the '
               'roster, loot and readiness at once. A character with an officer rank in game still waits for an officer.')
RANK_HELP = '0 is the guild master, 1 the rank under it, and so on. This rank and every rank above it can approve members and award loot.'
VISIBILITY_HELP = 'Public, Unlisted (link only) or Guild only. Applies to new reports.'
CREST_RULES = 'PNG, JPEG or WebP, up to 2 MB, at least 128 by 128; we crop to a square.'
TIERS = [('Guild', '$120', '$150', 'Up to 25 verified members &middot; 1 Premium seat'),
         ('Raid guild', '$200', '$250', 'Any size &middot; loot council and readiness &middot; 3 Premium seats')]

_CSS = ('<style>.sp-body{display:flex;flex-direction:column}.sp-btns{display:flex;flex-wrap:wrap;gap:12px;padding-top:12px}'
        '.sp-h3{display:flex;align-items:baseline;gap:12px;font:700 15px/20px var(--font-body);color:var(--strong);margin:0}'
        '.sp-seats{list-style:none;margin:0;padding:0;display:flex;flex-direction:column}'
        '.sp-seat{display:grid;grid-template-columns:minmax(220px,1fr) minmax(0,1.4fr) auto;align-items:center;gap:8px 16px;min-height:48px;padding:8px 0}'
        '.sp-who{display:flex;align-items:center;gap:12px;min-width:0}.sp-who b{font:700 14px/20px var(--font-body)}'
        '.sp-empty{display:inline-block;width:32px;height:32px;border-radius:50%;outline:1px dashed #ffffff33;outline-offset:-1px}'
        '.sp-col{display:flex;flex-direction:column;align-items:flex-start;gap:12px;padding-bottom:12px}'
        '.sp-in{box-sizing:border-box;height:44px;border-radius:4px;padding:0 12px;font:600 15px var(--font-body);color:var(--strong)}'
        '.sp-radio{display:flex;align-items:flex-start;gap:12px;padding:8px 0}.sp-radio input{margin-top:3px;width:16px;height:16px;accent-color:var(--gold)}'
        '.sp-body .row .sub{line-height:19.5px}.sp-radio{min-height:60px;box-sizing:border-box}.sp-act{display:inline-flex;align-items:center;min-height:44px}.sp-sr{position:absolute;width:1px;height:1px;padding:0;margin:-1px;overflow:hidden;clip:rect(0,0,0,0);white-space:nowrap}.sp-sel{height:44px;width:100%;box-sizing:border-box}.sp-file{position:absolute;width:1px;height:1px;opacity:0}</style>')


# The line under the page that says a save took or failed: empty here, its 21 px held under a 12 px gap.
NOTICE_LINE = '<div style="min-height:22px"></div>'


def body(html: str) -> str:
    """A Section body: 12 px under the head's hairline, a column (ui-section-body)."""
    return f'<div class="sp-body">{html}</div>'


def p(text: str, small: bool = True) -> str:
    return f'<p class="{"t-small" if small else ""}" style="font-size:{13 if small else 14}px;line-height:20px;margin:0">{text}</p>'


def day(iso: str) -> str:
    d = datetime.fromisoformat(iso.replace('Z', '+00:00'))
    return f'{d.day - (1 if d.hour == 0 else 0)} {d:%b}'


def renews_line() -> str:
    plan = S.billing()['billing']
    end = datetime.fromisoformat(plan['current_period_end'].replace('Z', '+00:00'))
    return f"Renews {end.day - 1} {end:%b %Y} &middot; ${plan['amount_cents'] // 100} a year"


def seat_rows() -> str:
    out = []
    for seat in S.billing()['seats']:
        if seat['holder'] is None:
            who = f'<span class="sp-empty" aria-hidden="true"></span><span class="t-small" style="font-size:14px">Not given</span>'
            give = _act('Give seat', small=False).replace('<button ', '<button aria-expanded="false" ', 1)
            out.append(f'<li class="sp-seat"><div class="sp-who">{who}</div><div></div><span>{give}</span></li>')
            continue
        h = seat['holder']
        who = (f'{crest_img(h["class"], 32).replace("<img ", "<img aria-hidden=\"true\" ", 1)}<div><b style="color:{C[h["class"]]}">{h["name"]}</b>'
               f'<div class="t-small" style="font-size:12px">Given {day(seat["given_at"])}</div></div>')
        move = _act('Move seat', small=False).replace('<button ', '<button aria-disabled="true" aria-expanded="false" ', 1)
        out.append(f'<li class="sp-seat"><div class="sp-who">{who}</div>{p(f"Can move from {day(seat["movable_at"])}")}<span>{move}</span></li>')
    return f'<ul class="sp-seats">{"".join(out)}</ul>'


def plan_section() -> str:
    view = S.billing()
    given = sum(1 for s in view['seats'] if s['holder'] is not None)
    rows = Rows([Row(renews_line(), trail=Tag('Founder'), link=False), Row('Billed to you.', link=False)], open_=True)
    who = Rows([Row('Who pays', WHO_PAYS_HELP, trail=_act('Transfer the plan'), link=False)], open_=True)
    seats = (f'<div style="display:flex;flex-direction:column;gap:8px;padding-top:20px"><h3 class="sp-h3">Premium seats '
             f'<span class="t-small" style="font-weight:400;font-size:13px">{given} of {len(view["seats"])} given</span></h3>{p(SEATS_HELP)}{seat_rows()}</div>')
    return Section('Guild plan', body(rows + f'<div class="sp-btns" style="padding-bottom:12px">{_act("Manage billing")}</div>' + who + seats), meta='Raid guild', panel=True, frame='framed')


def chooser_section() -> str:
    """No plan yet: the interval, the two plans as open radio rows, the count, the compare link, the trial button and its lines."""
    radios = ''.join(f'<label class="sp-radio"><input type="radio" name="tier"{" checked" if i == 0 else ""}><span><b class="strong">{name}</b> '
                     f'<b class="strong">{now}</b><span class="t-small"> a year</span> <span class="t-small">{later} after 6 Jan 2027</span>'
                     f'<span class="t-small" style="display:block">{line}</span></span></label>' for i, (name, now, later, line) in enumerate(TIERS))
    seg = Segmented([('Yearly', True), ('Monthly', False)]).replace('role="radiogroup"', 'role="radiogroup" aria-label="Billing interval"', 1)
    # The ready blocks sit 12 px under the section body's top (GuildBillingSection's lower block); the control is 46 px with its rule.
    inner = (f'<div style="display:flex;flex-direction:column;align-items:flex-start;gap:12px;max-width:720px;padding-top:12px">'
             f'<span style="display:flex;align-items:center;min-height:46px">{seg}</span>'
             f'<div role="radiogroup" aria-label="Plan" style="display:flex;flex-direction:column;gap:4px;width:100%">{radios}</div>'
             f'{p(f"{S.VERIFIED} verified members today.")}<span class="sp-act">{Link("Compare Guild and Raid guild &rsaquo;")}</span>{Button("Start the 7-day trial", primary=True)}'
             f'{p("First 7 days free, once per guild. Cancel before day 8 and nothing is charged; guild settings shows the date of the first charge.")}'
             f'{p("Founder price until 6 January 2027, kept while the plan stays active.")}</div>')
    return Section('Guild plan', body(inner), panel=True, frame='framed')


def discord_section() -> str:
    """A connected channel (webhook): GuildDiscordChannel's lines, its two actions and the reconnect link."""
    d = S.settings()['discord']
    short = f"{d['channel_id'][:4]}&hellip;{d['channel_id'][-4:]}"
    remove = _act('Remove').replace('class="btn', 'class="btn danger', 1)
    inner = (f'{p("Posts raids to <b class=\"strong\">a Discord channel</b>", small=False)}'
             f'<p class="t-small" style="font-size:13px;line-height:20px;margin:0">Channel <b class="strong">{short}</b> &middot; <a href="#" style="text-decoration:underline">Open in Discord</a></p>'
             f'<div class="sp-btns">{_act("Send test message")}{remove}</div>'
             f'<div style="padding-top:40px"><a href="#" style="font-size:13px;text-decoration:underline">Connect a different channel</a></div>'
             f'<div style="padding-top:12px">{p("Connecting again replaces this channel and removes its webhook from Discord.")}</div>')
    return Section('Discord', body(f'<div style="display:flex;flex-direction:column;gap:4px;min-height:176px">{inner}</div>'),
                   action=Tag('Connected', 'kill'), panel=True)


def discord_bot_section() -> str:
    """The Forever Sixty bot installed (GuildDiscordBot, installed): where raids post, the linked members, its three actions."""
    b = S.settings_bot()['discord']['bot']
    linked = f"{b['linked_members']} of {b['members']} members have linked Discord. The others get a link when they press a button."
    dest = (f'<div style="display:flex;flex-direction:column;gap:2px">'
            f'{p(f"Posts raids to <b class=\"strong\">#{b["channel_name"]}</b> in <b class=\"strong\">{b["server_name"]}</b>", small=False)}'
            f'{p("With sign-up buttons &middot; <a href=\"#\" style=\"text-decoration:underline\">Open in Discord</a>")}</div>')
    acts = f'<div class="sp-btns" style="padding-top:0">{_act("Send test message", small=False)}{_act("Change channel", small=False)}{_act("Remove", small=False)}</div>'
    # The status line under the actions holds its 21 px under an 8 px gap (DiscordStatusLine), empty at rest.
    inner = f'{dest}{p(linked)}<div style="display:flex;flex-direction:column">{acts}<div style="margin-top:8px;min-height:21px"></div></div>'
    return Section('Discord', body(f'<div style="display:flex;flex-direction:column;gap:12px">{inner}</div>'), action=Tag('Connected', 'kill'), panel=True)


def members_section() -> str:
    s = S.settings()
    holder = s['claimed_by']
    claimed = (f'<p style="font-size:13px;line-height:20px;margin:0">Claimed by {crest_img(holder["class"], 20).replace("<img ", "<img aria-hidden=\"true\" ", 1)} '
               f'<b style="color:{C[holder["class"]]}">{holder["name"]}</b></p>')
    invite = (Rows([Row('Invite link', INVITE_HELP, link=False)], open_=True)
              + f'<div class="sp-col"><label style="display:flex;flex-direction:column;gap:4px;width:100%"><span class="t-small" style="font-size:13px;font-weight:600">Invite link</span>'
              f'<input class="field sp-in" readonly value="https://foreversixty.gg{S.INVITE_URL}" style="height:29px"></label>{_act("Copy link", small=False)}'
              f'{p(f"Made {day(S.INVITE_MADE)} by <b style=\"color:{C[holder["class"]]}\">{holder["name"]}</b>.")}{_act("Make a new link")}</div>')
    rank = Row('Officer ranks', RANK_HELP, trail=f'<input class="field sp-in" type="number" aria-label="Officer ranks" value="{s["officer_max_rank_index"]}" style="width:72px">'
               + _act('Save').replace('<button ', '<button disabled ', 1), link=False)
    vis = Row('Default report visibility', VISIBILITY_HELP, trail=f'<span style="display:inline-block;width:160px">{_select("Guild only")}</span>', link=False)
    return Section('Members and access', body(claimed + invite + Rows([rank, vis], open_=True)), panel=True)


def _select(label: str) -> str:
    return f'<span role="combobox" aria-label="Default report visibility {label}" class="sel sp-sel">{label}</span>'


def records_section(crest: bool) -> str:
    rows = [Row('Download attendance (CSV)', trail='<span class="t-small">&rsaquo;</span>'), Row('Download loot (CSV)', trail='<span class="t-small">&rsaquo;</span>')]
    html = Rows(rows, open_=True)
    if crest:
        mark = Emblem('alliance', 40).replace('<img ', '<img aria-hidden="true" ', 1)
        choose = (f'<label class="btn sm" style="cursor:pointer">Choose an image<input class="sp-file" type="file" aria-label="Choose a crest image"></label>')
        html += (Rows([Row('Guild crest', 'Default: your faction&rsquo;s logo', lead=mark, trail=choose, link=False)], open_=True)
                 + f'<div class="sp-col" style="gap:12px"><div><span class="t-small" style="font-size:13px">No image chosen</span>'
                 f'<p class="t-small" style="font-size:12px;margin:0">{CREST_RULES}</p></div>{_act("Save").replace("<button ", "<button disabled ", 1)}</div>')
    return Section('Records and crest' if crest else 'Records', body(html), panel=True)


def settings_tab(bot: bool = False):
    main = plan_section() + (discord_bot_section() if bot else discord_section())
    rail = members_section() + records_section(crest=True)
    # The tab's own heading for screen readers (GuildSettings embedded): outside every region, as built.
    head = hdr(tabs_active='Settings') + '<h2 class="sp-sr">Guild settings</h2>'
    return build_page('Guild settings', _view(head + Page('split', main, rail, panels=True) + NOTICE_LINE + _CSS), 'Guild', ('Dunmaro', 'hunter', 'Beast Mastery Hunter &middot; 60'))


def settings_page():
    main = chooser_section() + discord_section()
    rail = members_section() + records_section(crest=False)
    head = f'<div style="padding-top:28px">{PageHeader(GUILD, ALLIANCE, meta=["Guild settings"])}</div>'
    return build_page('Guild settings page', _view(head + Page('split', main, rail, panels=True) + NOTICE_LINE + _CSS), 'Guild', ('Dunmaro', 'hunter', 'Beast Mastery Hunter &middot; 60'))


PAGES: dict = {'web-guild-settings': settings_tab, 'web-guild-settings-bot': lambda: settings_tab(bot=True), 'web-guild-settings-page': settings_page}
