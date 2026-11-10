"""Web-layout mock: home (/), round 5, every state, composed ONLY from ws_prims primitives.
Spec: design/specs/2026-11-09-directions.md "Home" (the look: design/mocks/directions/pages.py home(), png/g-home-*), on the facts and
states of design/specs/2026-11-07-home-b.md. Structure: the title band on the camp-road hero; the raid plaque (the ornate Frame on the
cave: the guild's next raid, else the first raids); Latest on Forever with Your upgrades (a Frame) beside it; then Guild sims in two
columns (or Your spec this week). Signed out: How to start in the rail and Pick your class below.

States (design/mocks/web/home-parity.json): out (signed out), 34 (a level 34 leveler, no guild), 60 (a level 60 raider in a
guild with a raid planned and the guild's sims), new (signed in, no character), guild-none (the level 60 raider's guild with no raid planned), feed-empty (the feed floored with real site
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
from ws_prims import (Button, FeedRow, FeedSkeleton, Frame, GradeTile, ItemRow, Link, MeterClass, Page, PageHeader, Plate, Row, RowGrid, Rows,
                      Section, Skeleton, build_page, crest_img, icon, write)
from ws_lib import REGION, QCOL, region_slug

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
FLOOR_MIN, FEED_ROWS, UPGRADES_SHOWN, GUILD_SHOWN = 5, 7, 5, 5
FRAME_LINE = 'rgba(243,238,226,.42)'  # web --color-frame-line
GUILD_SIM_TAIL = re.compile(r', #(\d+) of (\d+) DPS in [^,(]+ \([^)]*\)$')  # web lib/home/feed.ts homeRows
HOME_CSS = f'''
.ph-meta .t-small+.t-small::before{{content:none}}
.pad16{{padding:16px}}
.hans{{display:flex;flex-direction:column;align-items:inherit;gap:2px}}.hans .link{{font-size:13px}}
.hgrid{{display:grid;grid-template-columns:minmax(0,1fr);grid-template-areas:"lead" "rail" "feed" "foot";align-items:start;gap:var(--gap);padding-top:28px}}
.hgrid>div{{min-width:0}}.hg-lead{{grid-area:lead}}.hg-feed{{grid-area:feed}}.hg-rail{{grid-area:rail;display:flex;flex-direction:column;gap:var(--gap)}}.hg-foot{{grid-area:foot}}
@media(min-width:1100px){{.hgrid{{grid-template-columns:minmax(0,1fr) 400px;grid-template-areas:"lead lead" "feed rail" "foot foot"}}}}
.pq{{display:grid;gap:16px 34px}}
@media(min-width:900px){{.pq{{grid-template-columns:minmax(0,1fr) minmax(0,1fr)}}.pq .prep{{border-top:0;border-left:1px solid {FRAME_LINE};padding:0 0 0 28px}}}}
.pt{{font:400 30px/34px var(--font-title);margin:0}}.psub{{margin-top:4px;font-size:15px;line-height:22px}}
.bigf{{display:flex;flex-wrap:wrap;gap:12px 26px;margin-top:16px}}.bigf .g{{display:flex;gap:26px}}.bigf .rule{{width:1px;background:{FRAME_LINE}}}
.bigf .f{{display:flex;flex-direction:column}}.bigf b{{font-weight:600;font-size:44px;line-height:1;color:var(--strong)}}
.bigf .f span{{margin-top:4px;font-size:13px;line-height:18px;color:var(--muted)}}
.roles{{display:flex;flex-wrap:wrap;gap:6px 18px;margin-top:14px;font-size:14px;line-height:20px}}
.acts{{display:flex;flex-wrap:wrap;align-items:center;gap:10px;margin-top:18px}}
.prep{{border-top:1px solid {FRAME_LINE};padding-top:14px}}.prep h3{{font:400 19px/26px var(--font-title)}}
.prep ul{{margin:8px 0 0;padding:0;list-style:none}}.prep li{{display:flex;gap:10px;align-items:baseline;padding:3px 0;font-size:15px;line-height:22px}}
.prep li::before{{content:"";flex:none;width:11px;height:11px;border:1px solid {FRAME_LINE}}}
.frb{{container-type:inline-size}}.fr{{display:grid;gap:20px 34px;align-items:center}}@container (min-width:640px){{.fr{{grid-template-columns:minmax(0,1fr) auto}}}}
.fr .foot{{margin-top:18px;font-size:15px;line-height:22px}}
.fr ul{{display:flex;gap:18px;margin:0;padding:0;list-style:none}}
.fr li{{display:flex;flex-direction:column;align-items:center;gap:8px;width:96px;font:600 14px/18px var(--font-body);text-align:center}}
.fhd{{display:flex;align-items:baseline;justify-content:space-between;gap:12px;padding-bottom:6px}}.fhd h2{{font:400 22px/28px var(--font-title);margin:0}}
.lad>*{{break-inside:avoid}}
.simline{{padding:12px 0 8px;font-size:15px;line-height:22px}}
.facts{{display:flex;flex-wrap:wrap;gap:6px 18px;margin:0;font-size:14px;line-height:20px}}.facts b{{font-weight:600;color:var(--strong)}}
.upfoot{{padding-top:12px}}
.irow .tr{{display:flex;align-items:center;gap:12px}}
@media(max-width:767px){{.irow .worn{{display:none}}}}
.lbox{{container-type:inline-size}}@container (min-width:900px){{.lad{{columns:2;column-gap:48px}}}}
.strail{{display:flex;align-items:center;gap:14px}}.strail .meter{{width:120px}}.strail .sfig{{display:flex;justify-content:flex-end;min-width:104px}}.strail .gap{{width:96px;font-size:13px;line-height:18px;text-align:right;color:var(--muted)}}
@media(max-width:1099px){{.strail .meter{{width:56px}}.strail .gap{{display:none}}}}
@media(max-width:767px){{.strail .meter{{display:none}}}}
.fpair{{display:grid;grid-template-columns:minmax(0,1fr);align-items:start;gap:var(--gap)}}@media(min-width:1100px){{.fpair{{grid-template-columns:minmax(0,1fr) 400px}}}}
.sp h3{{margin:0;padding:8px 0 2px;font:600 13px/18px var(--font-body);color:var(--muted)}}
.err{{display:flex;flex-wrap:wrap;align-items:center;gap:12px;min-height:44px;font-size:14px;color:var(--muted)}}
.empty-b{{padding-block:12px;font-size:15px;line-height:24px;color:var(--muted)}}
.gap-sec{{height:var(--section-gap)}}
.lead-sk{{display:block;height:276px;border-radius:10px}}
@media(max-width:767px){{.hgrid{{padding-top:20px}}.pt{{font-size:25px;line-height:30px}}.bigf,.bigf .g{{gap:10px 16px}}.bigf b{{font-size:26px}}}}
'''


def J(name: str) -> dict:
    return json.loads((FIX / name).read_text())


def iso(s: str) -> datetime:
    return datetime.fromisoformat(s.replace('Z', '+00:00'))


def long_day(s: str) -> str:
    d = iso(s)
    return f'{d.strftime("%A")} {d.day} {d.strftime("%B")}'


def band_words(band: int) -> str:
    return '60' if band >= 60 else f'{band} to {band + 9}'


def spec_name(key: str) -> str:
    s = next(x for x in SPECS if f"{x['class_slug']}-{x['spec_slug']}" == key or x.get('spec') == key)
    cls = next(c['name'] for c in CLASSES if c['slug'] == s['class_slug'])
    return f"{s['name']} {cls}"


def sec(title: str, body: str, action: str = '', meta: str = '', head: str = '') -> str:
    """A boxed Section (round 5: every list sits on the panel)."""
    act = (head + (Link(f'{action} &rsaquo;') if action else '')) if (head or action) else ''
    return Section(title, body, act, meta, boxed=True)


def raids_open() -> str | None:
    raid = next((d for d in DATES if d.get('role') == 'raids'), None)
    return f"{raid['iso']}T00:00:00Z" if raid else None


def clock_to(target: str) -> list[tuple[str, str]]:
    """The plaque's clock (web lib/home/plaque.ts plaqueClock): days and hours, minutes in the final hour."""
    secs = int((iso(target) - NOW).total_seconds())
    if secs <= 0:
        return []
    if secs < 3600:
        return [(str(-(-secs // 60)), 'min')]
    return [(str(secs // 86400), 'days'), (str(secs % 86400 // 3600), 'hours')]


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


# ------------------------------------------------------------------------------------------------ the plaque
def big_figures(groups: list[list[tuple[str, str]]], timer: bool) -> str:
    """web ui/BigFigures: groups of a figure over its label, a hairline between; the first group a timer when `timer`."""
    out = []
    for i, group in enumerate(g for g in groups if g):
        if i:
            out.append('<span class="rule" aria-hidden="true"></span>')
        clock = i == 0 and timer
        attrs = f' role="timer" aria-label="{", ".join(f"{v} {u}" for v, u in group)}"' if clock else ''
        hide = ' aria-hidden="true"' if clock else ''
        out.append(f'<div class="g"{attrs}>' + ''.join(f'<span class="f"{hide}><b>{v}</b><span>{u}</span></span>' for v, u in group) + '</div>')
    return f'<div class="bigf">{"".join(out)}</div>'


def plaque_title(title: str, sub: str = '') -> str:
    return f'<h2 class="pt">{title}</h2>' + (f'<p class="psub">{sub}</p>' if sub else '')


def facts_line(facts: list[tuple[str, str]]) -> str:
    """web components/home/HomeFacts: one line of facts, a strong figure and its words."""
    return f'<p class="facts">{"".join(f"<span><b>{n}</b> {t}</span>" for n, t in facts)}</p>' if facts else ''


def standing_facts(char: dict, guild_home: dict) -> list[tuple[str, str]]:
    """web lib/home/home-facts.ts standingFacts: your item level place among your spec, the raids attended."""
    st, att = guild_home.get('standing'), char.get('guild', {}).get('attendance')
    out = []
    if st and st['same_spec_count'] > 1:
        out.append((f"#{st['rank_by_item_level']}", f"of {st['same_spec_count']} guild {st['spec']} {st['class']}s by item level"))
    if att and att['nights'] > 0:
        unit = 'raids' if guild_home.get('attendance_source') == 'raids' else 'raid nights'
        out.append((str(att['present']), f"of {att['nights']} {unit} attended"))
    return out


def summary_facts(guild_home: dict) -> list[tuple[str, str]]:
    """web lib/home/home-facts.ts guildSummaryFacts: the raiders, and the bosses down once the tier is pulled."""
    sm = guild_home.get('summary')
    if not sm:
        return []
    out = [(str(sm['raiders']), 'raider' if sm['raiders'] == 1 else 'raiders')]
    if sm['tier']['pulls_total'] > 0:
        out.append((str(sm['tier']['down']), f"of {sm['tier']['named_encounters']} bosses down"))
    return out


def guild_plaque(char: dict, guild_home: dict) -> str:
    r, needs = guild_home['next_raid'], guild_home['standing']['needs_before_next_raid']
    c = r['counts']
    standing = standing_facts(char, guild_home)
    d = iso(r['starts_at']) - timedelta(hours=6)
    when = f"{d.strftime('%A')} {d.day} {d.strftime('%B')} &middot; 8 pm"  # web lib/home/plaque.ts raidWhen, the viewer's time
    roles = ''.join(f'<span><b>{c[k]}</b> {k}</span>' for k in ('tanks', 'healers', 'melee', 'ranged'))
    own = r['own']['state'] if r['own'] else None
    buttons = ''.join(f'<button type="button" class="btn{" primary" if k == "in" else ""} sm" aria-pressed="{str(own == k).lower()}">{label}</button>'
                      for k, label in (('in', "I'm in"), ('late', 'Late'), ('out', "Can't come")))
    prep = (f'<div class="prep"><h3>Before the pull</h3><ul>{"".join(f"<li>{n}</li>" for n in needs)}</ul></div>') if needs else ''
    main = (f'<div style="min-width:0">{plaque_title(r["zone_name"], f"{char["guild"]["name"]} &middot; {when}")}'
            f'{big_figures([clock_to(r["starts_at"]), [(str(c["in"]), "signed up"), (str(c["late"]), "late")]], True)}'
            f'<p class="roles">{roles}</p>'
            + (f'<div style="margin-top:6px">{facts_line(standing)}</div>' if standing else '')
            + f'<div class="acts">{buttons}</div></div>')
    return Frame(f'<div class="pq">{main}{prep}</div>', ornate=True, region=REGION['countdowns'])


def first_raids_plaque(foot: str = '') -> str:
    at = raids_open()
    clock = clock_to(at)
    head = plaque_title('First raids open', long_day(at)) + big_figures([clock], True) if clock else plaque_title('Raids open now')
    f = f'<div class="foot">{foot}</div>' if foot else ''
    raids = '<ul aria-label="First raids">' + ''.join(f'<li>{icon(i, 56)}<span>{n}</span></li>' for n, i in RAIDS) + '</ul>'
    return Frame(f'<div class="frb"><div class="fr"><div style="min-width:0">{head}{f}</div>{raids}</div></div>', ornate=True, region=REGION['countdowns'])


def grid(lead: str, feed: str, rail: str = '', foot: str = '') -> str:
    """web components/home/HomeGrid: lead full width, the feed with the rail beside it, then the full-width foot."""
    r = f'<div class="hg-rail">{rail}</div>' if rail else ''
    f = f'<div class="hg-foot">{foot}</div>' if foot else ''
    return f'<div class="pc"><div class="hgrid"><div class="hg-lead">{lead}</div><div class="hg-feed">{feed}</div>{r}{f}</div></div>'


# ------------------------------------------------------------------------------------------------ the feed
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


def home_rows(rows: list[dict], raid_on_plaque: bool) -> list[dict]:
    """web lib/home/feed.ts homeRows: newest first; a guild sim reads `#2 of 14`; the raid post is left to the plaque."""
    out = []
    for r in sorted(rows, key=lambda x: iso(x['at']), reverse=True):
        if raid_on_plaque and r['kind'] == 'guild_raid':
            continue
        if r['kind'] == 'guild_sim':
            r = {**r, 'segments': [dict(s, text=GUILD_SIM_TAIL.sub(r', #\1 of \2', s['text'])) if s.get('style') is None else s
                                   for s in r['segments']]}
        out.append(r)
    return out


def feed_lines(rows: list[dict]) -> str:
    out = ''
    for r in rows[:FEED_ROWS]:
        lead = crest_img(r['crest'], 32) if r.get('crest') else icon(r.get('icon') or '', 32)
        when = day_text(r['at']) if r.get('date_only') else ago(r['at'])
        out += FeedRow(lead, sentence(r['segments']), when, '')
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


def feed_section(rows: list[dict] | None, state: str = 'ready') -> str:
    """Latest on Forever, with no filters (round 5)."""
    if state == 'loading':
        body = FeedSkeleton(FLOOR_MIN)
    elif state == 'error':
        body = ('<div class="err" role="alert"><span>We could not load what is happening on Forever.</span>'
                '<button type="button" class="btn sm">Try again</button></div>')
    else:
        shown = rows[:FEED_ROWS] if len(rows) >= FLOOR_MIN else rows + floor_rows()[:FLOOR_MIN - len(rows)]
        body = feed_lines(shown)
    return sec('Latest on Forever', body)


def frame_section(title: str, body: str, action: str = '') -> str:
    """A featured rail panel with its title inside (Your upgrades, How to start), carrying the section's region."""
    a = Link(f'{action} &rsaquo;') if action else ''
    return Frame(f'<div class="fhd"><h2>{title}</h2>{a}</div>{body}', region=region_slug(title))


# ------------------------------------------------------------------------------------------------ signed out
def class_picker(headless: bool = False) -> str:
    """web HomeClassPicker: boxed under its heading, or `headless` on its own plate under the title that names it."""
    rows = []
    for c in CLASSES:
        specs = ', '.join(s['name'] for s in SPECS if s['class_slug'] == c['slug'])
        rows.append(Row(c['name'], specs, lead=crest_img(c['slug'], 44), title_color=c['color']))
    if headless:
        return Plate(f'<div class="pad16">{RowGrid(rows, 3)}</div>', 'Pick your class')
    return sec('Pick your class', RowGrid(rows, 3))


def how_to_start() -> str:
    return frame_section('How to start', Rows([
        Row('Sign in with Battle.net', rank='1'),
        Row('Install the addon and the companion', rank='2'),
    ], open_=True))


def state_out() -> str:
    head = PageHeader('Pick your class', actions=Button('Sign in with Battle.net', True))
    rows = home_rows(J('feed-out.json')['rows'], False)
    return head + grid(class_picker(headless=True), feed_section(rows), how_to_start() + first_raids_plaque())


# ------------------------------------------------------------------------------------------------ signed in
def spec_key(char: dict) -> str:
    return f"{char['class'].lower()}-{char['spec'].lower().replace(' ', '-')}"


def upgrades_of(char: dict) -> tuple[list[tuple], dict]:
    cls = char['class'].lower()
    band = J(f'bis-{spec_key(char)}.json')['bands'][0]
    items = {i['id']: i for i in J(f'items-{cls}.json')['items']}
    ups = []
    for s in band['slots']:
        worn = char['build']['gear'].get(s['slot'])
        alt = next((a for a in s.get('alternatives') or [] if a['item_id'] == worn and a.get('dps_delta', 0) < 0), None)
        if alt:
            ups.append((-alt['dps_delta'], s))
    ups.sort(key=lambda u: -u[0])
    return ups, items


def who_line(char: dict) -> str:
    """web lib/home/home-view.ts whoLine, then the guild."""
    who = f"Level {char['level']} {char['race']} {char['spec']} {char['class']}"
    who += f" &middot; &lt;{char['guild']['name']}&gt;" if 'guild' in char else ''
    return who + (f" &middot; Item level {char['item_level']}" if char.get('item_level') else '')


def head_in(char: dict, home: dict) -> str:
    cls = char['class'].lower()
    sim = home['import_sim']
    hours = int((NOW - datetime.fromisoformat(sim['simmed_at'].replace('Z', '+00:00'))).total_seconds() // 3600)
    line = f'<a class="link" href="/sim/{sim["sim_id"]}">Sim {hours} h ago &rsaquo;</a>' if sim.get('sim_id') else ''
    answer = f'<div class="hans"><span class="fig"><span class="v">{sim["dps"]:,.1f}<small>DPS</small></span></span>{line}</div>'
    return PageHeader(char['name'], CLASS_COLOR[cls], crest_img(cls, 48), [who_line(char)], answer)


def plaque_in(char: dict, guild_home: dict | None) -> str:
    if guild_home is not None and 'next_raid' in guild_home:
        return guild_plaque(char, guild_home)
    if guild_home is not None:
        facts = summary_facts(guild_home) + standing_facts(char, guild_home)
        line = f'<div style="margin-top:8px">{facts_line(facts)}</div>' if facts else ''
        return first_raids_plaque(f"{char['guild']['name']}: no raid planned &middot; "
                                  f'<a class="link" href="#">Raid page &rsaquo;</a>{line}')
    return first_raids_plaque('Raid with a guild that plans its nights here. <a class="link" href="#">Guilds &rsaquo;</a>')


def upgrades_section(char: dict, ups: list, items: dict, bis_slots: int) -> str:
    """Your upgrades: each pick with the piece worn beside its gain, then the whole gap and the slots already best in slot."""
    rows = []
    for gain, s in ups[:UPGRADES_SHOWN]:
        it = items.get(s['item_id']) or {}
        q = it.get('quality', 1)
        ic = icon(it.get('icon', ''), 36, {2: 'q2', 3: 'q3', 4: 'q4', 5: 'q5'}.get(q, ''))
        # web lib/home/home-view.ts upgradeLine: the slot (a ring is a Finger, whichever of the two) and the pick's source
        slot = re.sub(r' \d$', '', SLOT_LABELS[s['slot']])
        row = ItemRow(s['item_name'], q, ic, f"{slot} &middot; {s['source']}", gain=(f'+{gain:,.1f}', 'DPS', ''))
        worn = items.get(char['build']['gear'].get(s['slot'])) or {}
        if worn:
            wq = {2: 'q2', 3: 'q3', 4: 'q4', 5: 'q5'}.get(worn.get('quality', 1), '')
            row = row.replace('<span class="tr">', f'<span class="tr"><span class="worn">{icon(worn.get("icon", ""), 28, wq)}</span>', 1)
        rows.append(row)
    facts = []
    total = sum(g for g, _ in ups)
    if total > 0:
        facts.append((f'+{total:,.1f}', f"DPS in {len(ups)} {'upgrade' if len(ups) == 1 else 'upgrades'}"))
    if bis_slots > 0:
        facts.append((str(bis_slots), 'best in slot'))
    foot = f'<div class="upfoot">{facts_line(facts)}</div>' if facts else ''
    return frame_section('Your upgrades', f'<div>{"".join(rows)}</div>{foot}', 'Best in slot')


def ladder_row(r: dict, unit: str, top: float, you: dict | None) -> str:
    """A guild sim: the bar in the class colour against the top of the ladder, the gap to your sim, the figure."""
    gap = '' if you is None or r['you'] else f"{abs(r['value'] - you['value']):,.1f} {'ahead' if r['value'] >= you['value'] else 'behind'}"
    trail = (f'<span class="strail">{MeterClass(r["value"] / top if top else 0, CLASS_COLOR[r["class"]])}<span class="gap">{gap}</span>'
             f'<span class="sfig"><span class="fig sm"><span class="v">{r["value"]:,.1f}<small>{unit}</small></span></span></span></span>')
    return Row(r['name'], '', lead=crest_img(r['class'], 28), trail=trail, title_color=CLASS_COLOR[r['class']], rank=str(r['rank']),
               selected=r['you'], slim=True)


def guild_sims_section(g: dict) -> str:
    top = g['rows'][0]['value'] if g['rows'] else 0
    you = next((r for r in g['rows'] if r['you']), None) or g['you']
    rows = [ladder_row(r, g['unit'], top, you) for r in g['rows'][:GUILD_SHOWN]]
    if g['you'] is not None and not any(r['you'] for r in g['rows'][:GUILD_SHOWN]):
        rows.append(ladder_row(g['you'], g['unit'], top, you))
    return sec('Guild sims', f'<div class="lbox"><div class="lad">{"".join(rows)}</div></div>')


def movers_block(m: dict | None) -> str:
    """The spec's board this week (web lib/home/home-facts.ts): the climbers with a baseline, else the top three."""
    if not m or not m['rows']:
        return ''
    rows = []
    for r in m['rows']:
        trail = (f'<span class="fig sm kill"><span class="v">+{r["delta"]}<small>places</small></span></span>'
                 if r['delta'] and r['delta'] > 0 else '')
        rows.append(Row(r['name'], '', lead=crest_img(r['class'], 28), trail=trail, title_color=CLASS_COLOR[r['class']],
                        rank=str(r['rank']), slim=True))
    title = 'Climbing this week' if m['baseline'] else 'Top this week'
    return f'<div><h3>{title}</h3>{"".join(rows)}</div>'


def spec_week_section(w: dict, movers: dict | None) -> str:
    place = f"#{w['place']} of {w['of']} damage specs"
    order = 'SABCD'
    if w['last_week_grade']:
        a, b = order.index(w['grade']), order.index(w['last_week_grade'])
        move = 'Same as last week' if a == b else (f"Up from {w['last_week_grade']} last week" if a < b else f"Down from {w['last_week_grade']} last week")
        sub = f'{move}, {place}'
    else:
        sub = place
    row = Row(f"{spec_name(w['spec'])} at {band_words(w['band'])}", sub, lead=GradeTile(w['grade']), link=False)
    body = f'<div class="sp">{Rows([row], open_=True)}{movers_block(movers)}</div>'
    return sec('Your spec this week', body, 'Tier list')


def state_in(level: str, feed: str = 'ready', feed_file: str = '', guild_file: str = 'guild-home-60.json') -> str:
    me = J(f'me-{level}.json')
    char = me['characters'][0]
    home = J(f'home-{level}.json')
    ups, items = upgrades_of(char)
    bis_slots = sum(1 for s in J(f'bis-{spec_key(char)}.json')['bands'][0]['slots']
                    if 'item_id' in s and char['build']['gear'].get(s['slot']) == s['item_id'])
    guild_home = J(guild_file) if 'guild' in char else None
    raid_on_plaque = guild_home is not None and 'next_raid' in guild_home
    rows = home_rows(J(feed_file or f'feed-{level}.json')['rows'], raid_on_plaque) if feed == 'ready' else None
    week = spec_week_section(home['spec_week'], home['movers']) if home['spec_week'] else ''
    foot = (f'<div class="fpair">{guild_sims_section(home["guild_sims"])}{week}</div>' if home['guild_sims'] else week)
    return head_in(char, home) + grid(plaque_in(char, guild_home), feed_section(rows, feed), upgrades_section(char, ups, items, bis_slots), foot)


def state_new() -> str:
    head = PageHeader('Get your character in')
    rows = Rows([
        Row('No character yet', 'Battle.net lists a character after you play it and log out.',
            trail='<button type="button" class="btn primary sm">Check again</button>', link=False),
        Row('Paste an addon export', trail='<a class="btn sm" href="#">Paste export</a>', link=False),
        Row('Get the addon and companion', trail='<a class="btn sm" href="#">Get set up</a>', link=False),
    ], open_=True)
    return head + Page('ledger', Plate(f'<div class="pad16">{rows}</div>', 'Get your character in') + '<div class="gap-sec"></div>' + class_picker())


def state_loading() -> str:
    # The loading head is the PageHeader, pending on the hero (HomeSignedIn.svelte): the crest skeleton, the title's skeleton
    # (the h1 a screen reader's: lib/home/home-view.ts loadingTitle), the whole head one status.
    head = (f'<div role="status" style="display:contents">'
            f'{PageHeader("Loading your character", crest=Skeleton(48, 48, True), pending=True)}</div>')
    return head + grid('<span class="skel lead-sk" aria-hidden="true"></span>', sec('Latest on Forever', FeedSkeleton(FLOOR_MIN)))


STATES = {
    'out': (state_out, None),
    '34': (lambda: state_in('34'), ('Felsnap', 'hunter', 'Level 34')),
    '60': (lambda: state_in('60'), ('Thandrel', 'warrior', 'Level 60')),
    'guild-none': (lambda: state_in('60', feed_file='feed-guild-none.json', guild_file='guild-home-none.json'), ('Thandrel', 'warrior', 'Level 60')),
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
