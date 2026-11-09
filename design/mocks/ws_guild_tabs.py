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
    return f"{r['spec']} {r['class'].title()}"


# ---------------------------------------------------------------- Roster
def _chips(labels, on=''):
    return ''.join(Button(x + (' &#9662;' if x == on else ''), small=True, act=True) for x in labels)


def roster():
    """The Roster (batch 3, b3-guild-roster): Waiting for approval framed with Approve all; the Roster panel with its filters and
    sort on one line, the figure heads, then the rows: a rank tag only for an officer or the guild master, the figures bare in their
    columns (item level, raids)."""
    home = T.roster_home()
    words = {'leader': 'Guild master', 'officer': 'Officer'}  # copy.ts guildRosterModelCopy.ranks
    waiting = [Row(r['name'], _spec(r), lead=crest_img(r['class'], 36), title_color=C[r['class']], link=False,
                   trail=(Tag(words[r['rank']]) if r['rank'] in words else '')
                   + Button('Approve', small=True, act=True) + Button('Decline', small=True, tone='ghost', act=True)) for r in home['pending']]
    wait = Section('Waiting for approval', Rows(waiting, open_=True), action=Button('Approve all', primary=True, small=True, act=True), panel=True, frame='framed')
    verified = sorted(home['roster'], key=lambda r: (r['name'] != O.OFFICER, -r['item_level']))
    cols = lambda r: (f'<span class="rcol" style="width:88px">{r["item_level"]}</span>'  # noqa: E731
                      f'<span class="rcol" style="width:132px">{(str(r["attendance"]["present"]) + " of " + str(r["attendance"]["nights"])) if r["attendance"]["nights"] else ""}</span>')
    # The viewer's own name links to their member profile.
    # The rank rides the spec line ("Beast Mastery Hunter · Officer"), a member's has none.
    sub = lambda r: _spec(r) + (f' &middot; {words[r["rank"]]}' if r['rank'] in words else '')  # noqa: E731
    rows = [Row(f'<a href="#">{r["name"]} (you)</a>' if r['name'] == O.OFFICER else r['name'], sub(r), lead=crest_img(r['class'], 36), title_color=C[r['class']], link=False,
                selected=r['name'] == O.OFFICER,
                trail=f'<span style="display:flex;align-items:center;min-height:39px">{cols(r)}</span>')
            for r in verified]
    controls = (f'<div style="display:flex;flex-wrap:wrap;gap:8px 24px;align-items:center">'
                f'<div style="display:flex;flex-wrap:wrap;gap:8px;align-items:center"><span class="t-small">Filter</span>'
                f'{_chips(["Role: All", "Class: All", "Verified only", "No account", "Below rating floor"])}</div>'
                f'<div role="group" aria-label="Sort" style="display:flex;flex-wrap:wrap;gap:8px;align-items:center"><span class="t-small">Sort</span>'
                f'{_chips(["Rank", "Item level", "Attendance", "Last seen"], "Item level")}</div></div>')
    heads = ('<div aria-hidden="true" style="display:flex;justify-content:flex-end;gap:12px;padding-bottom:8px;font:700 12px/16px var(--font-body);color:var(--muted)">'
             '<span class="rcol" style="width:88px">Item level</span><span class="rcol" style="width:132px">Raids</span></div>')
    body = _col(controls, f'<div role="region" aria-label="Roster">{heads}{Rows(rows)}</div>')
    n = len(home['roster']) + len(home['pending'])
    sec = Section('Roster', body, meta=f"{n} on the roster &middot; {len(home['pending'])} waiting for approval", panel=True)
    css = '<style>.rcol{display:inline-block;text-align:right;font-size:14px;font-variant-numeric:tabular-nums}</style>'
    return _tab('Guild roster', 'Roster', Page('ledger', _col(wait, sec, gap=22)).replace('padding-top:20px', 'padding-top:28px', 1) + css)


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
def shortfalls(r: dict, raid: dict) -> list[str]:
    """What a raider is short of (readiness-view readinessShortfalls), worst first."""
    if r['no_export']:
        return ['no export yet']
    out = []
    n = (r['gear_gap'] or {}).get('upgrades') or 0
    if n:
        out.append(f"{n} upgrade{'' if n == 1 else 's'}")
    miss = len(r['enchants']['missing_slots']) if r['enchants']['checked'] else 0
    if miss:
        out.append(f"{miss} enchant{'' if miss == 1 else 's'} missing")
    ff = r.get('flask_food')
    lacks = ([] if ff is None else (['no flask or raid elixir'] if not ff['flask_or_elixir'] else []) + (['no buff food'] if not ff['food'] else []))
    if lacks:
        out += lacks
    elif r['consumables']['state'] == 'short':
        out.append('consumables short')
    if r['talent_points_unspent']:
        out.append(f"{r['talent_points_unspent']} talent point{'' if r['talent_points_unspent'] == 1 else 's'} unspent")
    if raid.get('attunement') and r.get('attuned') is False:
        out.append('not attuned')
    return out


def readiness():
    """Readiness by exception (batch 3, b3-guild-readiness): who is short of what first, the tanks' fire resistance on its own
    panel (a raid that asks for it), then who is ready."""
    page = F.readiness(True)
    raid = page['raid']
    rows = sorted(page['rows'], key=lambda r: -r['failing'])  # worst first; the board's sort holds ties in the API's order
    short = [(r, shortfalls(r, raid)) for r in rows if shortfalls(r, raid)]
    ready = [r for r in rows if not shortfalls(r, raid)]
    at = guild_time(raid['starts_at'])
    listed = Rows([Row(r['name'], f"{r['spec']} {r['class'].title()} &middot; item level {r['item_level']}", lead=crest_img(r['class'], 30), title_color=C[r['class']],
                       foot=f'<span class="gold" style="font-size:13px;line-height:20px">{" &middot; ".join(why)}</span>', trail=Button('Nudge', small=True, act=True), link=False)
                   for r, why in short], list_=True)
    toggle = Button('Next raid only', small=True, act=True)
    body = Section('Short of something', listed, toggle, meta=f'In or Late for {day_label(at)}', panel=True, frame='framed')
    if raid.get('fire_resistance'):
        signups = F.raid_view(O.OFFICER, True)['signups']
        tanks = [x for x in signups if x['role'] == 'tank' and x['state'] in ('in', 'late')]
        fire = {r['name']: (r.get('resistances') or {}).get('fire') for r in page['rows']}
        trows = Rows([Row(t['name'], title_color=C[t['class']], link=False, slim=True,
                          trail=f'<span style="font-size:14px;font-weight:600">{fire[t["name"]]} fire resistance</span>' if fire.get(t['name']) is not None
                          else '<span>Fire resistance not shared</span>') for t in tanks], list_=True)
        body += Section('Fire resistance on tanks', trows, panel=True)
    names = '<span class="t-small">, </span>'.join(f'<span style="color:{C[r["class"]]}">{r["name"]}</span>' for r in ready)
    body += Section(f'{len(ready)} ready', f'<p style="font-size:14px;line-height:23px">{names}</p>',
                    meta='Best in slot or no upgrades, enchanted, stocked, every point spent', panel=True)
    return _tab('Guild readiness', 'Readiness', Page('ledger', _col(body, gap=22)).replace('padding-top:20px', 'padding-top:28px', 1))


PAGES: dict = {'web-guild-roster': roster, 'web-guild-progression': progression, 'web-guild-readiness': readiness}
