"""Web-layout mock: home (/), layout B "the world is moving", every state, composed ONLY from ws_prims primitives.
Spec: design/specs/2026-11-07-home-b.md (direction mock design/mocks/web/web-home-alive-b.html).

States (design/mocks/web/home-parity.json): out (signed out), 34 (a level 34 leveler, no guild), 60 (a level 60 raider in a
guild with a raid planned and the guild's sims), new (signed in, no character), feed-empty (the feed floored with real site
events), feed-loading, loading (session hinted, /v1/me in flight), error (the feed failed). Every figure is read from
web/tests/e2e/fixtures/home/ (home_fixtures.py writes them), the same files the parity spec routes; the floor rows are read
from the same sources the build reads (the raid date, the site changelog, the BiS files).

  python3 design/mocks/home_fixtures.py && python3 design/mocks/ws_home.py
  node design/mocks/ws_shoot.mjs web-home-60 1440,2000
"""
import json
import re
from datetime import datetime, timedelta, timezone
from pathlib import Path

from mocklib import CLASS_COLOR
from ws_prims import (Button, Countdown, FeedRow, FeedSkeleton, GradeTile, ItemRow, Link, Page, PageHeader, Row, RowGrid, Rows,
                      Section, Segmented, Skeleton, build_page, crest_img, icon, write)
from ws_lib import REGION, QCOL

ROOT = Path(__file__).resolve().parents[2]
BUILD = json.loads((ROOT / 'web/src/data/active-build.json').read_text())['build']
FIX = ROOT / 'web/tests/e2e/fixtures/home'
CLASSES = json.loads((ROOT / 'web/src/data/classes.json').read_text())
SPECS = json.loads((ROOT / 'data/curated/specs.json').read_text())
DATES = json.loads((ROOT / 'web/src/data/dates.json').read_text())
NOW = datetime(2026, 11, 6, 18, tzinfo=timezone.utc)
SLOT_LABELS = {'head': 'Head', 'neck': 'Neck', 'shoulder': 'Shoulder', 'back': 'Back', 'chest': 'Chest', 'wrist': 'Wrist', 'hands': 'Hands',
               'waist': 'Waist', 'legs': 'Legs', 'feet': 'Feet', 'finger1': 'Finger 1', 'finger2': 'Finger 2', 'trinket1': 'Trinket 1',
               'trinket2': 'Trinket 2', 'main_hand': 'Main hand', 'off_hand': 'Off hand', 'ranged': 'Ranged'}
RAIDS = [('Barrow Deeps', 'spell_shadow_raisedead'), ('Hyjal Summit', 'spell_nature_timestop'), ("Onyxia's Lair", 'inv_misc_head_dragon_black')]
TAGS = {'your_spec': 'Your spec', 'logs': 'Logs', 'rankings': 'Rankings', 'best_in_slot': 'Best in slot', 'guides': 'Guides', 'your_guild': 'Your guild', 'patch': 'Patch'}
FLOOR_MIN, FEED_ROWS, UPGRADES_SHOWN, GUILD_SHOWN = 5, 7, 3, 5
HOME_CSS = '''
.sec-b{padding-top:12px}
.ph-meta .t-small+.t-small::before{content:none}
.head-sk{display:flex;flex-direction:column;gap:12px;padding-block:20px 16px;min-height:118px;justify-content:flex-end}
.figs-b{display:flex;flex-wrap:wrap;justify-content:flex-end;gap:12px 36px}
.figs-b .f{display:flex;flex-direction:column;align-items:flex-end;gap:2px;text-align:right}
.figs-b .f .v{font:700 28px/32px var(--font-body);font-variant-numeric:tabular-nums;color:var(--strong)}
.figs-b .f .v small{font-size:13px;font-weight:600;color:var(--muted);margin-left:4px}
.figs-b .f .v.kill{color:var(--kill)}.figs-b .f .c{font-size:12px;line-height:16px;color:var(--muted)}
.cds{display:grid;gap:var(--gap);grid-template-columns:minmax(0,1.3fr) minmax(0,1fr);padding-top:28px}
.cds.solo{grid-template-columns:minmax(0,1fr)}
.raids{display:flex;gap:18px;margin:0;padding:0;list-style:none}
.raids li{display:flex;flex-direction:column;align-items:center;gap:6px;width:88px;font:600 13px/16px var(--font-body);text-align:center;color:var(--nav)}
.gw{margin:0;text-align:right;font-size:14px;line-height:22px;color:var(--nav)}.gw b{display:block;font-size:18px;line-height:24px;color:var(--strong)}
.dl{color:var(--kill);font-size:13px;font-weight:700}
.err{display:flex;flex-wrap:wrap;align-items:center;gap:12px;min-height:44px;font-size:14px;color:var(--muted)}
.empty-b{padding-block:12px;font-size:15px;line-height:24px;color:var(--muted)}
.gap-sec{height:var(--section-gap)}
'''


def J(name: str) -> dict:
    return json.loads((FIX / name).read_text())


def iso(s: str) -> datetime:
    return datetime.fromisoformat(s.replace('Z', '+00:00'))


def short_date(s: str) -> str:
    d = iso(s)
    return f'{d.day} {d.strftime("%b")}'


def band_words(band: int) -> str:
    return '60' if band >= 60 else f'{band} to {band + 9}'


def spec_name(key: str) -> str:
    s = next(x for x in SPECS if f"{x['class_slug']}-{x['spec_slug']}" == key or x.get('spec') == key)
    cls = next(c['name'] for c in CLASSES if c['slug'] == s['class_slug'])
    return f"{s['name']} {cls}"


def sec(title: str, body: str, action: str = '', meta: str = '', head: str = '') -> str:
    act = (head + (Link(f'{action} &rsaquo;') if action else '')) if (head or action) else ''
    return Section(title, f'<div class="sec-b">{body}</div>', act, meta)


def raids_open() -> str | None:
    raid = next((d for d in DATES if d.get('role') == 'raids'), None)
    return f"{raid['iso']}T00:00:00Z" if raid else None


def clock_to(target: str) -> list[tuple[str, str]]:
    left = iso(target) - NOW
    secs = int(left.total_seconds())
    if secs < 3600:
        return [(f'{secs // 60:02d}', 'min'), (f'{secs % 60:02d}', 'sec')]
    return [(f'{secs // 86400:02d}', 'days'), (f'{secs % 86400 // 3600:02d}', 'hours'), (f'{secs % 3600 // 60:02d}', 'min')]


def day_text(s: str) -> str:
    d = iso(s)
    return f'{d.strftime("%a")} {d.day} {d.strftime("%b")}'


def ago(s: str) -> str:
    m = int((NOW - iso(s)).total_seconds() // 60)
    if m < 1:
        return 'just now'
    if m < 60:
        return f'{m} min ago'
    return f'{m // 60} h ago' if m < 1440 else f'{m // 1440} d ago'


# ------------------------------------------------------------------------------------------------ shared blocks
def raid_panel() -> str:
    at = raids_open()
    raids = '<ul class="raids" aria-label="First raids">' + ''.join(f'<li>{icon(i, 40)}<span>{n}</span></li>' for n, i in RAIDS) + '</ul>'
    return Countdown(f'First raids open &middot; {day_text(at)}', clock_to(at), raids, 'raid', live=True)


def countdowns(guild_panel: str = '') -> str:
    solo = '' if guild_panel else ' solo'
    return f'<div class="pc"><div class="cds{solo}" data-region="{REGION["countdowns"]}">{raid_panel()}{guild_panel}</div></div>'


def sentence(segments: list[dict]) -> str:
    out = ''
    for s in segments:
        st = s.get('style')
        if st == 'name':
            out += f'<b style="color:{CLASS_COLOR[s["class"]]}">{s["text"]}</b>'
        elif st == 'item':
            out += f'<b style="color:{QCOL.get(s.get("quality", 1), "#fff")}">{s["text"]}</b>'
        elif st == 'strong':
            out += f'<b style="color:var(--strong)">{s["text"]}</b>'
        else:
            out += s['text']
    return out


def feed_lines(rows: list[dict]) -> str:
    out = ''
    for r in rows[:FEED_ROWS]:
        lead = crest_img(r['crest'], 40) if r.get('crest') else icon(r.get('icon') or '', 40)
        when = day_text(r['at']) if r.get('date_only') else ago(r['at'])
        out += FeedRow(lead, sentence(r['segments']), when, TAGS[r['tag']])
    return f'<div>{out}</div>'


def floor_rows() -> list[dict]:
    """The floor the build bakes (web/src/lib/home/floor.ts): the raid opening, then the newest site events."""
    rows = []
    at = raids_open()
    if at and iso(at) > NOW:
        rows.append({'icon': RAIDS[2][1], 'at': at, 'date_only': True, 'tag': 'patch',
                     'segments': [{'text': 'First raids open', 'style': 'strong'}, {'text': ': ' + ', '.join(n for n, _ in RAIDS)}]})
    events = []
    for f in (ROOT / 'web/src/content/changelog').glob('*.md'):
        head = f.read_text().split('---')[1]
        title = re.search(r"^title: '?(.*?)'?$", head, re.M).group(1)
        date = re.search(r'^date: (\S+)', head, re.M).group(1)
        events.append({'icon': 'inv_misc_book_09', 'at': f'{date}T00:00:00Z', 'tag': 'patch', 'segments': [{'text': title}]})
    events.sort(key=lambda e: e['at'], reverse=True)
    return rows + events


def feed_section(rows: list[dict] | None, filters: list[str] | None, state: str = 'ready') -> str:
    head = Segmented([(f, i == 0) for i, f in enumerate(filters)], small=True) if filters else ''
    if state == 'loading':
        body = FeedSkeleton(FLOOR_MIN)
    elif state == 'error':
        body = ('<div class="err" role="alert"><span>We could not load what is happening on Forever.</span>'
                '<button type="button" class="btn sm">Try again</button></div>')
    else:
        shown = rows[:FEED_ROWS] if len(rows) >= FLOOR_MIN else rows + floor_rows()[:FLOOR_MIN - len(rows)]
        body = feed_lines(shown)
    return sec('Forever right now', body, head=head)


# ------------------------------------------------------------------------------------------------ signed out
def class_picker() -> str:
    rows = []
    for c in CLASSES:
        specs = ', '.join(s['name'] for s in SPECS if s['class_slug'] == c['slug'])
        rows.append(Row(c['name'], specs, lead=crest_img(c['slug'], 44), title_color=c['color']))
    return sec('Pick your class', RowGrid(rows, 3))


def how_to_start() -> str:
    return sec('How to start', Rows([
        Row('Sign in with Battle.net', 'Imports your characters with their level, gear and item level.', rank='1'),
        Row('Install the addon and the companion', 'The addon rates every item in your bags; the companion uploads your logs as you play.', rank='2'),
    ], open_=True))


def state_out() -> str:
    head = PageHeader('Play your class better.', meta=['Pick your class.', f'First raids {short_date(raids_open())}'],
                      actions=Button('Sign in with Battle.net', True))
    return (head + Page('ledger', class_picker()) + countdowns()
            + Page('split', feed_section(J('feed-out.json')['rows'], None), how_to_start()))


# ------------------------------------------------------------------------------------------------ signed in
def spec_key(char: dict) -> str:
    return f"{char['class'].lower()}-{char['spec'].lower().replace(' ', '-')}"


def upgrades_of(char: dict) -> tuple[list[tuple], dict, int]:
    cls = char['class'].lower()
    bis = J(f'bis-{spec_key(char)}.json')
    band = bis['bands'][0]
    items = {i['id']: i for i in J(f'items-{cls}.json')['items']}
    ups, on_bis = [], 0
    for s in band['slots']:
        worn = char['build']['gear'].get(s['slot'])
        if 'item_id' in s and worn == s['item_id']:
            on_bis += 1
        alt = next((a for a in s.get('alternatives') or [] if a['item_id'] == worn and a.get('dps_delta', 0) < 0), None)
        if alt:
            ups.append((-alt['dps_delta'], s))
    ups.sort(key=lambda u: -u[0])
    return ups, items, on_bis


def figure(value: str, label: str, unit: str = '', caption: str = '', kill: bool = False) -> str:
    u = f'<small>{unit}</small>' if unit else ''
    c = f'<span class="c">{caption}</span>' if caption else ''
    return f'<span class="f"><span class="t-label">{label}</span><span class="v{" kill" if kill else ""}">{value}{u}</span>{c}</span>'


def head_in(char: dict, home: dict, ups: list, on_bis: int) -> str:
    cls = char['class'].lower()
    meta = [f"Level {char['level']} {char['spec']} {char['class']}", f"{char['faction'].capitalize()} {char['race']}", f"Item level {char['item_level']}"]
    if 'guild' in char:
        meta.append(f"&lt;{char['guild']['name']}&gt;")
    sim = home['import_sim']
    figs = figure(f"{sim['dps']:,.1f}", 'Simmed on import', 'DPS', f"Simmed {short_date(sim['simmed_at'])}")
    if ups:
        figs += figure(f'+{sum(u[0] for u in ups):,.1f}', f'{len(ups)} upgrades', 'DPS', kill=True)
    level = char['level']
    figs += figure(f'{on_bis} of 17', 'On best in slot') if level >= 60 else figure(str((level // 10 + 1) * 10 - level), f'To band {(level // 10 + 1) * 10}', 'levels')
    return PageHeader(char['name'], CLASS_COLOR[cls], crest_img(cls, 56), meta, f'<div class="figs-b">{figs}</div>')


def guild_panel(char: dict, guild_home: dict | None) -> str:
    if 'guild' not in char:
        return Countdown('Find a guild', None, '<p class="gw">Raid with a guild that plans its nights here.<br><a class="link" href="#">Guilds &rsaquo;</a></p>', 'guild')
    r = guild_home['next_raid']
    d = iso(r['starts_at']) - timedelta(hours=6)
    when = f"{d.strftime('%a')} {d.day} {d.strftime('%b')} &middot; 8:00 pm CST"
    state = r['own']['state'] if r['own'] else 'not signed'
    what = (f"<p class=\"gw\"><b>{r['title']}</b>{when} &middot; {r['counts']['in']} in &middot; you: {state}<br>"
            f"<a class=\"link\" href=\"#\">Sign up on the raid page &rsaquo;</a></p>")
    return Countdown(f"{char['guild']['name']} &middot; next guild event", clock_to(r['starts_at']), what, 'guild')


def upgrades_section(ups: list, items: dict) -> str:
    rows = []
    for gain, s in ups[:UPGRADES_SHOWN]:
        it = items.get(s['item_id']) or {}
        q = it.get('quality', 1)
        ic = icon(it.get('icon', ''), 32, {2: 'q2', 3: 'q3', 4: 'q4', 5: 'q5'}.get(q, ''))
        rows.append(ItemRow(s['item_name'], q, ic, f"{SLOT_LABELS[s['slot']]} &middot; {s['source']}", gain=(f'+{gain:,.1f}', 'DPS', '')))
    total = sum(u[0] for u in ups)
    return sec('Your upgrades', ''.join(rows), 'Best in slot', f'+{total:,.1f} DPS')


def ladder_row(r: dict, unit: str) -> str:
    trail = f'<span class="fig sm"><span class="v">{r["value"]:,.1f}<small>{unit}</small></span></span>'
    return Row(r['name'], '', lead=crest_img(r['class'], 28), trail=trail, title_color=CLASS_COLOR[r['class']], rank=f"#{r['rank']}",
               selected=r['you'], slim=True)


def guild_sims_section(g: dict) -> str:
    rows = [ladder_row(r, g['unit']) for r in g['rows'][:GUILD_SHOWN]]
    if g['you'] is not None and not any(r['you'] for r in g['rows'][:GUILD_SHOWN]):
        rows.append(ladder_row(g['you'], g['unit']))
    role = {'dps': 'DPS', 'tank': 'Tanks', 'healer': 'Healers'}[g['role']]
    return sec('Guild sims', Rows(rows, open_=True), 'Guild page', f"{role}, {band_words(g['band'])} {g['preset']}")


def spec_week_section(w: dict) -> str:
    place = f"#{w['place']} of {w['of']} damage specs"
    order = 'SABCD'
    if w['last_week_grade']:
        a, b = order.index(w['grade']), order.index(w['last_week_grade'])
        move = 'Same as last week' if a == b else (f"Up from {w['last_week_grade']} last week" if a < b else f"Down from {w['last_week_grade']} last week")
        sub = f'{move}, {place}'
    else:
        sub = place
    row = Row(f"{spec_name(w['spec'])} at {band_words(w['band'])}", sub, lead=GradeTile(w['grade']), link=False)
    return sec('Your spec this week', Rows([row], open_=True), 'Tier list')


def movers_section(m: dict) -> str:
    rows = [Row(r['name'], '', lead=crest_img(r['class'], 28), title_color=CLASS_COLOR[r['class']], rank=f"#{r['rank']}", slim=True,
                trail=(f'<span class="dl">&#9650; {r["delta"]}</span>' if r['delta'] else '')) for r in m['rows']]
    spec = next(x['name'] for x in SPECS if f"{x['class_slug']}-{x['spec_slug']}" == m['spec'])
    if m['baseline']:
        return sec('Top movers', Rows(rows, open_=True), 'Rankings', f"{spec}, {band_words(m['band'])}")
    return sec(f"Top 3 for {spec_name(m['spec'])}, {band_words(m['band'])}", Rows(rows, open_=True), 'Rankings')


def state_in(level: str, feed: str = 'ready', feed_file: str = '') -> str:
    me = J(f'me-{level}.json')
    char = me['characters'][0]
    home = J(f'home-{level}.json')
    ups, items, on_bis = upgrades_of(char)
    guild_home = J('guild-home-60.json') if 'guild' in char else None
    filters = ['All', 'Your spec', 'Your guild', 'Logs'] if 'guild' in char else ['All', 'Your spec', 'Logs']
    rows = J(feed_file or f'feed-{level}.json')['rows'] if feed == 'ready' else None
    main = feed_section(rows, filters, feed)
    rail = upgrades_section(ups, items)
    if home['guild_sims']:
        rail += guild_sims_section(home['guild_sims'])
    rail += spec_week_section(home['spec_week']) + movers_section(home['movers'])
    return head_in(char, home, ups, on_bis) + countdowns(guild_panel(char, guild_home)) + Page('split', main, rail)


def state_new() -> str:
    head = PageHeader('Play your class better.', meta=['Signed in as Hunthrop#1894'])
    rows = Rows([
        Row('Check again', 'Battle.net lists a character after you play it and log out.', trail='<button type="button" class="btn primary sm">Check again</button>', link=False),
        Row('Paste an addon export', 'Your gear and talents, straight from the game.'),
        Row('Get the addon and companion', 'The addon rates every item in your bags; the companion uploads your logs as you play.'),
    ], open_=True)
    note = f'<div class="pc"><div data-region="{REGION["summary"]}"><p class="notice gold">No character yet. Battle.net lists a character after you play it and log out.</p></div></div>'
    return head + note + Page('ledger', sec('Get your character in', rows) + '<div class="gap-sec"></div>' + class_picker())


def state_loading() -> str:
    head = f'<div class="pc"><div data-region="{REGION["head"]}" class="head-sk" role="status" aria-busy="true">{Skeleton(56, 56, True)}{Skeleton(280, 28)}</div></div>'
    return head + Page('split', sec('Forever right now', FeedSkeleton(FLOOR_MIN)), '')


STATES = {
    'out': (state_out, None),
    '34': (lambda: state_in('34'), ('Felsnap', 'hunter', 'Level 34')),
    '60': (lambda: state_in('60'), ('Thandrel', 'warrior', 'Level 60')),
    'new': (state_new, ('Hunthrop', 'warrior', 'No character')),
    'feed-empty': (lambda: state_in('34', feed_file='feed-empty.json'), ('Felsnap', 'hunter', 'Level 34')),
    'feed-loading': (lambda: state_in('34', 'loading'), ('Felsnap', 'hunter', 'Level 34')),
    'loading': (state_loading, ('Felsnap', 'hunter', 'Level 34')),
    'error': (lambda: state_in('34', 'error'), ('Felsnap', 'hunter', 'Level 34')),
}


def main() -> None:
    for state, (render, who) in STATES.items():
        body = render() + f'<style>{HOME_CSS}</style>'
        write(f'web-home-{state}.html', build_page('Forever Sixty', body, who=who))


if __name__ == '__main__':
    main()
