"""Guild mocks, Roster, Progression and Readiness tabs (GUILD-ALL lane, 2026-11-07). Composed only from ws_prims primitives; the data is
guild_tabs_fixtures.py's (the Roster's home with its waiting list) and the Overview's own answers (progression, readiness), which
guild_fixtures.py writes for the build. Registered in ws_guild2.all_pages().

Every tab is the guild head, the band line, then open Sections; the tab's primary list is its one Plate: the roster's rows, the
readiness board. Progression has no Plate: each named encounter is an open Section, the tier and the zones with no encounter list
sit in the rail.
"""
from datetime import date

from ws_prims import Page, Section, Plate, Rows, Row, Button, Figure, Meter, Tag, build_page, crest_img
import guild_overview_data as O
import guild_overview_fixtures as F
import guild_tabs_fixtures as T
from guild_data import C
from ws_guild2 import hdr, band_line, _view, pad

WHO = ('Dunmaro', 'hunter', 'Beast Mastery Hunter &middot; 60')
SPLIT_GAP = '<div style="display:flex;flex-direction:column;gap:12px">{}</div>'


def _tab(title, active, body):
    return build_page(title, _view(hdr(tabs_active=active, actions=Button('Invite link', small=True)) + band_line(O.OFFICER) + body), 'Guild', WHO)


def _col(*parts, gap=16):
    return f'<div style="display:flex;flex-direction:column;gap:{gap}px">{"".join(parts)}</div>'


def _p(text, muted=True):
    return f'<p class="t-small" style="font-size:14px;line-height:20px">{text}</p>' if muted else f'<p style="font-size:14px;line-height:20px">{text}</p>'


def _p8(text):
    """A muted note with the 8 px under it the built note keeps before its rows."""
    return f'<p class="t-small" style="font-size:14px;line-height:20px;padding-bottom:8px">{text}</p>'


def _spec(r):
    return f"{r['spec']} {r['class'].title()} on the site"


# ---------------------------------------------------------------- Roster
def _chips(labels, on=''):
    return ''.join(Button(x + (' &#9662;' if x == on else ''), small=True, act=True) for x in labels)


def roster():
    home = T.roster_home()
    rank = lambda r: Tag('Officer') if r['rank'] == 'officer' else '<span class="t-small">Member</span>'  # noqa: E731
    waiting = [Row(r['name'], _spec(r), lead=crest_img(r['class'], 36), title_color=C[r['class']], link=False,
                   trail=rank(r) + Button('Approve', small=True, act=True) + Button('Decline', small=True, tone='ghost', act=True)) for r in home['pending']]
    wait = Section('Waiting for approval', pad(_col(_p('Approving lets this character see the roster, loot and readiness; officers can also remove members and change settings.'),
                                                    Rows(waiting, open_=True), gap=12)),
                   action=Button('Approve all', small=True, act=True), meta=str(len(home['pending'])))
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
    body = _col(_p('Loot and readiness list members with gear on the site.'), controls,
                _p('No roster from the addon yet. Log in with the addon on your guild master and it sends one.'),
                Plate(Rows(rows), 'Roster'),
                _p('Raid nights each raider came to, of those they signed up for or came to, over this guild&rsquo;s last 8 and 30 logged raids.'))
    n = len(home['roster']) + len(home['pending'])
    sec = Section('Roster', pad(body), meta=f"{n} on the roster &middot; {len(home['pending'])} waiting for approval")
    return _tab('Guild roster', 'Roster', Page('ledger', _col(wait, sec, gap=48)))


# ---------------------------------------------------------------- Progression
def progression():
    page = F.progression()
    t = page['tier']
    main = ''
    for e in (x for x in page['encounters'] if x['pulls'] > 0):
        nights = Rows([Row(f"{date.fromisoformat(n['date']):%a} {date.fromisoformat(n['date']).day} {date.fromisoformat(n['date']):%b}",
                           f"{n['pulls']} pulls", trail=Tag('Kill', 'kill') if n['killed'] else '', link=False, slim=True) for n in e['pulls_by_night']],
                     open_=True, list_=True)
        main += Section(e['name'], pad(_col(_p('First kill Mon 7 Dec &middot; best kill 3:45', False), nights,
                                            _p(f"Deaths per pull: {e['deaths_per_pull']:.1f} on average."), gap=12)),
                        meta=f"{e['zone']} &middot; pulls to kill")
    tier = Section('This tier', pad(_col(Figure(f"{t['down']} of {t['named_encounters']}", '', size='sm').replace('</span></div>', '</span><span class="t-small" style="display:block">Named bosses down</span></div>', 1),
                                         Meter(t['down'] / t['named_encounters'], '', f"{t['down']} of {t['named_encounters']} named bosses down"), gap=8)),
                   meta=t['name'])
    zones = [z for z in page['unnamed'] if z['pulls'] > 0]
    unnamed = Section('Without an encounter list', pad(_p8('Barrow Deeps and Hyjal Summit have no encounter list yet. Their pulls count toward the tier total '
                                                          'but cannot be named or ranked per boss.')
                                                       + Rows([Row(z['zone'], f"{z['pulls']} pulls logged", link=False, slim=True) for z in zones], open_=True, list_=True)))
    return _tab('Guild progression', 'Progression', Page('split', main, tier + unnamed))


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
    toggle = Button('Next raid &middot; Thu 10 Dec', small=True, act=True)
    showing = _p(f'Showing the {len(order)} signed In or Late for <a class="link" href="#">Onyxia&rsquo;s Lair</a>, Thu 10 Dec, 8:00 pm.')
    body = _col(showing, Plate(head + f'<div class="ruled">{rows}</div>', 'Readiness board'),
                _p('Sorted by the most checks to fix, then the biggest gear gap. Nudge copies a message you can paste to them.'), gap=12)
    return _tab('Guild readiness', 'Readiness', Page('ledger', Section('Readiness', pad(body), action=toggle)))


PAGES: dict = {'web-guild-roster': roster, 'web-guild-progression': progression, 'web-guild-readiness': readiness}
