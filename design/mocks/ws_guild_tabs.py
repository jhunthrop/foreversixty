"""Guild mocks, Roster, Progression and Readiness tabs (GUILD-ALL lane, 2026-11-07). Composed only from ws_prims primitives; the data is
guild_tabs_fixtures.py's (the Roster's home with its waiting list) and the Overview's own answers (progression, readiness), which
guild_fixtures.py writes for the build. Registered in ws_guild2.all_pages().

Every tab is the guild head, the band line, then open Sections; the tab's primary list is its one Plate: the roster's rows, the
readiness board. Progression has no Plate: each named encounter is an open Section, the tier and the zones with no encounter list
sit in the rail.
"""

from ws_prims import Page, Section, Plate, Rows, Row, Button, Meter, Tag, build_page, crest_img, icon
import guild_overview_data as O
import guild_overview_fixtures as F
import guild_tabs_fixtures as T
from guild_data import C
from ws_guild2 import RAID_ICON, hdr, band_line, _view, pad, guild_time, day_label

WHO = ('Dunmaro', 'hunter', 'Beast Mastery Hunter &middot; 60')
SPLIT_GAP = '<div style="display:flex;flex-direction:column;gap:12px">{}</div>'


def _tab(title, active, body):
    strip = band_line(O.OFFICER) if active == 'Roster' else ''  # the next raid's strip on the Roster only (batch 3)
    head = hdr(tabs_active=active, sub='Members' if active == 'Roster' else '')
    return build_page(title, _view(head + strip + body), 'Guild', WHO)


def _col(*parts, gap=16):
    return f'<div style="display:flex;flex-direction:column;gap:{gap}px">{"".join(parts)}</div>'


def _p(text, muted=True):
    return f'<p class="t-small" style="font-size:14px;line-height:20px">{text}</p>' if muted else f'<p style="font-size:14px;line-height:20px">{text}</p>'


def _p8(text):
    """A muted note with the 8 px under it the built note keeps before its rows."""
    return f'<p class="t-small" style="font-size:14px;line-height:20px;padding-bottom:8px">{text}</p>'


def _hour(at):
    return f"{at.hour % 12 or 12}:{at.minute:02d} {'am' if at.hour < 12 else 'pm'}"


def _spec(r):
    return f"{r['spec']} {r['class'].title()} on the site"


# ---------------------------------------------------------------- Roster
def _chips(labels, on=''):
    return ''.join(Button(x + (' &#9662;' if x == on else ''), small=True, act=True) for x in labels)


def roster():
    """The Roster (batch 3, b3-guild-roster): Waiting for approval framed with Approve all, then the Roster panel with its
    filters, sort, the addon line and the rows. No explaining lines."""
    home = T.roster_home()
    words = {'leader': 'Leader', 'officer': 'Officer'}  # copy.ts guildRosterModelCopy.ranks
    rank = lambda r: Tag(words[r['rank']]) if r['rank'] in words else '<span class="t-small">Member</span>'  # noqa: E731
    waiting = [Row(r['name'], _spec(r), lead=crest_img(r['class'], 36), title_color=C[r['class']], link=False,
                   trail=rank(r) + Button('Approve', small=True, act=True) + Button('Decline', small=True, tone='ghost', act=True)) for r in home['pending']]
    wait = Section('Waiting for approval', Rows(waiting, open_=True), action=Button('Approve all', primary=True, small=True, act=True), panel=True, frame='framed')
    verified = sorted(home['roster'], key=lambda r: (r['name'] != O.OFFICER, -r['item_level']))
    # The viewer's own name links to their member profile.
    rows = [Row(f'<a href="#">{r["name"]} (you)</a>' if r['name'] == O.OFFICER else r['name'], _spec(r), lead=crest_img(r['class'], 36), title_color=C[r['class']], link=False,
                selected=r['name'] == O.OFFICER,
                trail=f"<span style=\"display:flex;gap:12px;align-items:center;min-height:39px\">{rank(r)}</span><span class=\"t-small\">Item level {r['item_level']} &middot; {r['attendance']['present']} of {r['attendance']['nights']} raid nights</span>")
            for r in verified]
    controls = (f'<div style="display:flex;flex-wrap:wrap;gap:8px;align-items:center"><span class="t-small">Filter</span>'
                f'{_chips(["Role: All", "Class: All", "Verified only", "No account", "Below rating floor"])}</div>'
                f'<div role="group" aria-label="Sort" style="display:flex;flex-wrap:wrap;gap:8px;align-items:center"><span class="t-small">Sort</span>'
                f'{_chips(["Rank", "Item level", "Attendance", "Last seen"], "Item level")}</div>')
    body = _col(controls, _p('No roster from the addon yet. Log in with the addon on your guild master and it sends one.'),
                f'<div role="region" aria-label="Roster">{Rows(rows)}</div>')
    n = len(home['roster']) + len(home['pending'])
    sec = Section('Roster', body, meta=f"{n} on the roster &middot; {len(home['pending'])} waiting for approval", panel=True)
    return _tab('Guild roster', 'Roster', Page('ledger', _col(wait, sec, gap=22)).replace('padding-top:20px', 'padding-top:28px', 1))


# ---------------------------------------------------------------- Progression
def progression():
    """Progression (batch 3, b3-guild-progression): one framed panel a zone, its bosses down over a meter, a Row per boss pulled."""
    page = F.progression()
    zones: dict[str, list] = {}
    for e in page['encounters']:
        zones.setdefault(e['zone'], []).append(e)
    main = ''
    for zone, encs in zones.items():
        pulled = [e for e in encs if e['pulls'] > 0]
        if not pulled:
            continue
        down = sum(1 for e in encs if e['kills'] > 0)
        rows = []
        for e in pulled:
            first = f" &middot; first kill {day_label(guild_time(e['first_kill_at']))}" if e.get('first_kill_at') else ''
            wipes = e['pulls'] - e['kills']
            sub = (f"{e['pulls']} pull{'' if e['pulls'] == 1 else 's'} &middot; {wipes} wipe{'' if wipes == 1 else 's'} &middot; "
                   f"{e['deaths_per_pull']:.1f} deaths a pull{first}")
            rows.append(Row(e['name'], sub, lead=icon(RAID_ICON, 36, 'q5'), trail=Tag('Killed', 'kill') if e['kills'] else Tag('Progress'), link=False))
        body = _col(Meter(down / len(encs), 'kill', f'{down} of {len(encs)} named bosses down'), Rows(rows, list_=True), gap=8)
        main += Section(zone, body, meta=f"{down} of {len(encs)} boss{'' if len(encs) == 1 else 'es'} down", panel=True, frame='framed')
    zones_unnamed = [z for z in page['unnamed'] if z['pulls'] > 0]
    if zones_unnamed:
        main += Section('Without an encounter list', Rows([Row(z['zone'], f"{z['pulls']} pulls logged", link=False, slim=True) for z in zones_unnamed], list_=True), panel=True)
    return _tab('Guild progression', 'Progression', Page('ledger', _col(main, gap=22)).replace('padding-top:20px', 'padding-top:28px', 1))


# ---------------------------------------------------------------- Readiness
def readiness():
    page = F.readiness(True)
    order = page['rows']  # worst first; no raider has a failing check, so the roster's order holds
    heads = ['Raider', 'Gear gap', 'Enchants', 'Consumables', 'Talent points', 'Avg item level', 'Last synced', '']
    cols = 'grid-template-columns:minmax(120px,170px) minmax(120px,170px) minmax(120px,170px) 100px 100px 140px 128px 84px'
    head = f'<div style="display:grid;{cols};gap:8px;padding:12px 16px 8px">' + ''.join(f'<span class="t-small">{h}</span>' for h in heads) + '</div>'
    rows = ''.join(
        f'<div style="display:grid;{cols};gap:8px;align-items:center;min-height:56px;padding:0 12px;margin:0 4px" class="rrow">'
        f'<span style="display:flex;gap:8px;align-items:center">{crest_img(r["class"], 26)}<b style="color:{C[r["class"]]}">{r["name"]}{" (you)" if r["name"] == O.OFFICER else ""}</b></span>'
        f'<span class="t-small">No upgrades</span><span class="kill">All enchanted</span><span class="{"kill" if r["consumables"]["state"] == "stocked" else "wipe"}">'
        f'{"Stocked" if r["consumables"]["state"] == "stocked" else "Short"}</span><span class="t-small">None unspent</span><span>{r["item_level"]}<span class="t-small" style="display:block">at the raid median</span></span>'
        f'<span class="gold">Synced today</span>{Button("Nudge", small=True, act=True)}</div>' for r in order)
    at = guild_time(page['raid']['starts_at'])
    toggle = Button(f'Next raid &middot; {day_label(at)}', small=True, act=True)
    showing = _p(f'Showing the {len(order)} signed In or Late for <a class="link" href="#">Onyxia&rsquo;s Lair</a>, {day_label(at)}, {_hour(at)}.')
    body = _col(showing, Plate(head + f'<div class="ruled">{rows}</div>', 'Readiness board'),
                _p('Sorted by the most checks to fix, then the biggest gear gap. Nudge copies a message you can paste to them.'), gap=12)
    return _tab('Guild readiness', 'Readiness', Page('ledger', Section('Readiness', pad(body), action=toggle)))


PAGES: dict = {'web-guild-roster': roster, 'web-guild-progression': progression, 'web-guild-readiness': readiness}
