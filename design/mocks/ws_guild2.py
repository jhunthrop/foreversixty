"""Guild persona round: role-aware landing (member and officer), raid planner, loot council, raid report, recruitment, applications and a
member profile. Composed only from ws_prims primitives. Data shapes: api/cmd/seedguild (24 characters), api/internal/guilds (sign-up states
In / Late / Bench, loot candidates, readiness, attendance, raid report). Modules whose data source is not confirmed carry a Gap marker.
"""
from ws_prims import *  # noqa: F403
from ws_prims import (Page, PageHeader, Section, Plate, Rows, Row, Figures, Figure, Meter, Tag, Button, Link, TabStrip, Select, Notice, Segmented, Grid, Tally,
                      Checklist, Votes, Gap, build_page, crest_img, icon, write, Emblem, RowGroup, Prose, control_bar, seg, SignupRow, ButtonRow,
                      SpellIcon, DisclosureButton, VoteButtons, Toggle, Entries, Inspector, Live, Split, RowsSlim, Disclosure, ItemRow)

from datetime import datetime
from zoneinfo import ZoneInfo

import guild_overview_data as O
import guild_overview_fixtures as F
from guild_data import ALLIANCE, GUILD, NOW, TABS, ZONE, C, R, STATE_TAG  # noqa: F401
import guild_data as gd
import guild_raid
import guild_report_data as rd

# Figures and rows under a Section head sit 12 px below its hairline, as the built Section body does (padding-top on a layout div).
# Lane A's words (the built page's own): the raid's icon (lib/first-raids.ts), the sign-up role words by roster group
# (raidCopy.roleLabels), the BiS slot words (readiness-view enchantSlotLabel) and the loot table's slot lines (slotLine).
RAID_ICON = 'inv_misc_head_dragon_black'
GROUP_ROLE_WORD = {'Tank': 'Tank', 'Healer': 'Healer', 'Melee': 'DPS', 'Ranged': 'DPS'}
SLOT_WORD = {'hands': 'Hands', 'finger1': 'Finger 1'}
SLOT_LINE = {(13, 7): 'One-Hand &middot; Sword', (13, 0): 'One-Hand &middot; Axe', (2, None): 'Neck', (11, None): 'Finger', (16, None): 'Back'}


def ch(name):
    return next(r for r in R if r[0] == name)


def lead(name, size=32):
    r = ch(name)
    return crest_img(r[1], size)


def band_facts(raiders: int | None = gd.RAIDERS) -> str:
    """The band's one line on every guild tab (batch 3 ruling; lib/guild/fact-line bandFacts): "US PvP · 24 raiders"."""
    return ' &middot; '.join(['US PvP'] + ([f'{raiders} raider{"" if raiders == 1 else "s"}'] if raiders else []))


def hdr(tabs_active='Overview', officer=True, tabs=None, raiders: int | None = gd.RAIDERS, sub: str = ''):
    """The guild's head as the built Guild.svelte draws it: the faction crest, the one band line, the role's tabs as a tablist
    (hash tabs on one URL); `sub` is the Roster's sub-view row (Members, Applications: link tabs), its current view."""
    shown = tabs or (TABS if officer else [t for t in TABS if t != 'Settings'])
    views = TabStrip(['Members', 'Applications'], sub).replace('strip-tabs ph-tabs', 'strip-tabs', 1) if sub else ''
    return PageHeader(GUILD, '', Emblem('alliance', 64), meta=[band_facts(raiders)], tabs=TabStrip(shown, tabs_active, tablist=True) + views)


def pad(html: str) -> str:
    """A Section body's 12 px under the head's hairline (the built Section body's padding)."""
    return f'<div style="padding-top:12px">{html}</div>'


def guild_time(iso: str) -> datetime:
    return datetime.fromisoformat(iso.replace('Z', '+00:00')).astimezone(ZoneInfo(ZONE))


def day_label(at: datetime) -> str:
    return f"{at.strftime('%a')} {at.day} {at.strftime('%b')}"


def clock_label(at: datetime) -> str:
    hour = at.hour % 12 or 12
    return f"{hour}{'' if at.minute == 0 else f':{at.minute:02d}'} {'am' if at.hour < 12 else 'pm'} {at.tzname()}"


def countdown_label() -> str:
    left = guild_time(O.NEXT_RAID_AT) - guild_time(NOW)
    days, hours = left.days, left.seconds // 3600
    return f'in {days} days {hours} h'


def next_raid(officer):
    """The Next raid panel's body: the raid's Row (icon, the guild's time, the countdown), the member's own Row with In / Late /
    Bench / Out (the built sign-up's four states), and for an officer the In, Late, Bench and Not signed figures."""
    at = guild_time(O.NEXT_RAID_AT)
    top = Row(O.ZONE_NAME, f'{day_label(at)} &middot; {clock_label(at)}', lead=icon(RAID_ICON, 40, 'q5'),
              trail=Figure(countdown_label(), '', tone='gold', size='sm'), link=False)
    you = ''
    if not officer:
        name = O.MEMBER
        _, cls, spec, role, _, state = ch(name)
        word = {'in': 'In', 'late': 'Late', 'bench': 'Bench'}[state]
        you = Row(name, f'{spec} {cls.title()} &middot; level 60 &middot; you are <strong>{word}</strong> as {GROUP_ROLE_WORD[role]}', lead=crest_img(cls, 36),
                  title_color=C[cls], trail=Segmented([(w, w.lower() == state) for w in ('In', 'Late', 'Bench', 'Out')], small=True), link=False)
    tally = ''
    if officer:
        view = F.raid_view(O.OFFICER, True)
        c = view['counts']
        tally = '<div style="padding-top:12px">' + Figures([(str(c['in']), 'In', '', 'kill', ''), (str(c['late']), 'Late', '', 'gold', ''), (str(c['bench']), 'Bench', '', '', ''),
                                                           (str(len(view['not_signed_raiders'])), 'Not signed', '', 'wipe', '')]) + '</div>'
    return top + you + tally


def comp_grid(counts: dict, target: dict | None = None) -> str:
    """Comp against the target (CompTally, overview-role-view compGridRows): per role the seats, who is In, how many short."""
    rows = []
    target = target or O.COMP_TARGET
    for role in ('tanks', 'healers', 'melee', 'ranged'):
        have, want, late = counts[role], target[role], counts['late_groups'][role]
        tone = 'ok' if have >= want else ('warn' if have + late >= want else 'miss')
        short = want - have
        rows.append((role.title(), f'{late} late' if late else '', [('none', str(want)), (tone, str(have)), (tone, str(short)) if short > 0 else ('none', '&mdash;')]))
    return Grid(['Need', 'In', 'Short'], rows, caption='Comp against the target', edge=True)


def last_raid(meta_line: str) -> str:
    counts = O.last_raid_counts()
    at = guild_time(O.LAST_RAID_AT)
    facts = [day_label(at), f'{counts["kills"]} boss ({O.BOSS})', f'{counts["kills"]} kill', f'{counts["wipes"]} wipe']
    row = Row(f'{O.ZONE_NAME} &middot; raid night', ' &middot; '.join(facts), trail='<span class="gold" style="font-weight:600">Report &rsaquo;</span>')
    return Section('Last raid', Rows([row], open_=True), meta=meta_line, panel=True)


def progression_line() -> str:
    t = F.tier()
    return f"{t['down']} of {t['named_encounters']} named bosses down: {O.BOSS}."


def checklist_rows() -> list[tuple]:
    """The member's checklist rows in the built words (overview-role-view checklistRows) for guild_overview_fixtures.checklist()."""
    rows = []
    for item in F.checklist()['items']:
        k, ok = item['key'], item['state'] == 'ok'
        if item['state'] == 'not_tracked':
            continue
        if k == 'level':
            rows.append(('ok', f"Level {item['required_level']}", f"You are level {item['level']}.", None))
        elif k == 'attunement':
            rows.append(('ok', f'Attuned to {O.ZONE_NAME}', f"Attunement chain done. Keep the {item['key_item']['name']} on you.", None))
        elif k == 'gear_gap':
            slots = [SLOT_WORD[x] for x in item['slots']]
            rows.append(('ok', 'Best in slot', 'Every slot is your best in slot.', None) if ok else
                        ('warn', f"{item['upgrades']} slots are not best in slot", f"{' and '.join(slots)}.", 'Fix in the BiS list &rsaquo;'))
        elif k == 'enchants':
            rows.append(('ok', 'Enchants', 'Every slot that takes one is enchanted.', None))
        elif k == 'flask_food':
            rows.append(('warn', 'No flask, elixir or buff food in your bags', 'Bring a flask or a raid elixir, and buff food.', None))
    return rows


def landing_member():
    """The Overview as a member. Since the persona round (summary.md): no "Officers asked you to" and no class-wanted Tally (no
    source); the rail carries the raid's comp; progression is the Last raid's meta line; a check with no source is not drawn."""
    rows = checklist_rows()
    open_ = sum(1 for r in rows if r[0] != 'ok')
    words = ['no', 'one', 'two', 'three', 'four', 'five']
    weekday = guild_time(O.NEXT_RAID_AT).strftime('%A')
    notice = Notice(f'Almost there: {words[open_]} things before {weekday}.', Link('See your checklist &rsaquo;'))
    me = next(r for r in F.home(O.MEMBER, False)['roster'] if r['name'] == O.MEMBER)
    nums = Figures([(f"{me['attendance']['present']} of {me['attendance']['nights']}", 'Raids attended', '', '', 'sm'), (str(me['item_level']), 'Your item level', '', '', 'sm')])
    main = (Section('Next raid', next_raid(False), panel=True) + Section('Before the raid', Checklist(rows), meta='Private to you &middot; Almost there', panel=True)
            + Section('Your numbers', nums, panel=True))
    rail = Section("The raid's comp", comp_grid(F.raid_view(O.MEMBER, False)['counts']), panel=True, flush=True) + last_raid(progression_line())
    return build_page('Guild, member', _view(hdr(officer=False) + notice + Page('split', main, rail, sticky_rail=True, panels=True)), 'Guild', ('Hollowmend', 'priest', 'Holy Priest &middot; 60'))


COVER_SHOWN = 6  # BuffCover: the gaps and the thinnest lines; the rest behind "All"


def cover_rows_thinnest() -> list[dict]:
    """The cover, thinnest first (cover-view coverThinnestFirst): nobody, then too few, then covered; the fewest signed first."""
    def tone(r):
        signed = r['providers_in'] + r['providers_late']
        return 0 if signed == 0 else (2 if signed - r['likely'] >= r['need'] else 1)
    return sorted(F.cover_rows(), key=lambda r: (tone(r), r['providers_in'] + r['providers_late']))


def cover_grid() -> str:
    rows = []
    for r in cover_rows_thinnest()[:COVER_SHOWN]:
        signed = r['providers_in'] + r['providers_late']
        tone = 'miss' if signed == 0 else ('ok' if signed - r['likely'] >= r['need'] else 'warn')
        sub = ' &middot; '.join(x for x in (f"{r['likely']} talent not confirmed" if r['likely'] else '',
                                             f"{r['providers_bench']} on the bench" if r['providers_bench'] else '') if x)
        rows.append((r['name'], sub, [('none', str(r['need'])), (tone, str(signed))], SpellIcon(r['name'], r['icon'])))
    return Grid(['Need', 'In'], rows, caption='Buff and debuff cover against signed-up raiders', edge=True)


def attendance_grid() -> str:
    cell = {'present': ('ok', '&#10003;'), 'absent': ('miss', '&times;'), 'bench': ('warn', 'B'), 'none': ('none', '&mdash;')}
    page = F.recent_attendance()
    cols = [guild_time(r['starts_at']).strftime('%a') for r in page['raids']]
    rows = sorted(page['rows'], key=lambda r: (not r['slipping'], r['name']))
    return Grid(cols, [(r['name'], 'slipping' if r['slipping'] else 'bench', [cell[c] for c in r['cells']]) for r in rows], caption='Attendance over the last raids', edge=True)


def action_rows() -> str:
    a = F.officer_actions()
    weekday = guild_time(O.NEXT_RAID_AT).strftime('%A')
    rows = [Row(f"{a['approvals']['count']} raiders waiting for approval", ', '.join(a['approvals']['names']), trail=Button('Review', small=True, act=True), link=False),
            Row(f"{a['applications']['count']} applications", ', '.join(a['applications']['names']), trail=Button('Open', small=True), link=False),
            Row(f"{a['not_signed']['count']} not signed up for {weekday}", ', '.join(a['not_signed']['names']), trail=Button('Ping on Discord', small=True, act=True), link=False),
            Row(f"{a['drops_to_award']['count']} drops to award", a['drops_to_award']['label'], trail=Button('Open loot', small=True, act=True), link=False)]
    return Rows(rows, open_=True)


def landing_officer():
    """The Overview as an officer (batch 3, b3-guild-overview): Action needed (framed), Next raid, Comp against the target;
    the rail's cover thinnest first, and attendance once there are two raids to compare."""
    main = (Section('Action needed', action_rows(), panel=True, frame='framed') + Section('Next raid', next_raid(True), Link('Open the planner &rsaquo;'), panel=True)
            + Section('Comp against the target', comp_grid(F.raid_view(O.MEMBER, False)['counts']), Link('Edit target &rsaquo;'), panel=True, flush=True))
    n = len(F.cover_rows())
    rail = Section('Cover, thinnest first', cover_grid(), Button(f'All {n} &rsaquo;', small=True, tone='ghost', act=True) if n > COVER_SHOWN else '', panel=True, flush=True)
    if len(O.RECENT_RAIDS) > 1:
        rail += Section(f'Attendance, last {len(O.RECENT_RAIDS)} raids', attendance_grid(), Link('Roster &rsaquo;'), panel=True, flush=True)
    return build_page('Guild, officer', _view(hdr() + Page('split', main, rail, sticky_rail=True, panels=True)), 'Guild', ('Dunmaro', 'hunter', 'Beast Mastery Hunter &middot; 60'))


ROLE_WORD = {'tank': 'Tank', 'healer': 'Healer', 'dps': 'DPS'}
CLASS_ORDER = ['warrior', 'paladin', 'hunter', 'rogue', 'priest', 'shaman', 'mage', 'warlock', 'druid']
SIGNUP_WORDS = ['In', 'Late', 'Bench', 'Out']


def planner_sub(m):
    """The planner row's line: `Holy Paladin · Healer` (raid-planner-view rowSub)."""
    return f"{m['spec']} {m['class'].title()} &middot; {ROLE_WORD[m['role']]}"


def signup_segments(state):
    return Segmented([(w, w.lower() == state) for w in SIGNUP_WORDS], small=True)


def planner_row(m, sub=''):
    """A sign-up Row as an officer reads it: crest, class-colour name, the in-row In / Late / Bench / Out."""
    return Row(m['name'], sub or planner_sub(m), lead=crest_img(m['class'], 28), title_color=C[m['class']], trail=signup_segments(m['state']),
               link=False, selected=m['name'] == guild_raid.VIEWER)


def roster_groups(raid):
    """Group by role, the default for a raid (raid-planner-view plannerGroups): In by role, then Late and Bench."""
    order = {'tank': 0, 'healer': 1, 'dps': 2}
    rows = sorted((guild_raid.member(s['name']) for s in raid['signups']), key=lambda m: (order[m['role']], CLASS_ORDER.index(m['class']), m['name']))
    on = [m for m in rows if m['state'] == 'in']
    groups = [(label, [m for m in on if m['group'] == g]) for label, g in [('Tanks', 'tank'), ('Healers', 'healer'), ('Melee', 'melee'), ('Ranged', 'ranged')]]
    groups += [(label, [m for m in rows if m['state'] == st]) for label, st in [('Late', 'late'), ('Bench', 'bench')]]
    return ''.join(RowGroup(label, len(ms), [planner_row(m) for m in ms]) for label, ms in groups if ms or label in ('Tanks', 'Healers', 'Melee', 'Ranged'))


def raid_when(raid):
    """The head's time line in the guild's zone (the viewer is in it: no "your time"), counted from guild_data.NOW."""
    from datetime import datetime, timedelta, timezone
    start = datetime.fromisoformat(raid['starts_at'].replace('Z', '+00:00'))
    cst = timezone(timedelta(hours=-6))
    a, b = start.astimezone(cst), (start + timedelta(minutes=raid['duration_minutes'])).astimezone(cst)
    left = start - datetime.fromisoformat(NOW.replace('Z', '+00:00'))
    clock = lambda t: t.strftime('%I:%M %p').lstrip('0').lower()  # noqa: E731
    return (f"{a.strftime('%A')} {a.day} {a.strftime('%B')} &middot; {clock(a)} to {clock(b)} CST &middot; "
            f'<span class="gold">in {left.days} days {left.seconds // 3600} h</span>')


def assignment_rows(encounter, edited=None):
    """Each slot's Row; the `edited` slot (an officer's Edit) reads Done and opens its picker under the Row (AssignmentSlotRow)."""
    rows = []
    for s in encounter['slots']:
        names = ''.join(Tag(n['name'], 'gold' if n['own'] else 'muted') for n in s['names']) or '<span class="t-small">Open</span>'
        own = any(n['own'] for n in s['names'])
        open_ = s['key'] == edited
        row = Row(s['label'], f"{len(s['names'])} of {s['count']} &middot; {s['purpose']}", trail=Button('Done' if open_ else 'Edit', small=True, tone='ghost', act=True),
                  foot=names, link=False, selected=own)
        if open_:
            row += f'<div style="display:flex;flex-direction:column;gap:8px;padding:8px 0 12px">{Select("Add a raider")}</div>'
        rows.append(row)
    return rows


def raid_planner(edited=None):
    raid = guild_raid.raid_view()
    n, by = raid['counts'], {m['name']: m for m in guild_raid.MEMBERS}
    you = by[guild_raid.VIEWER]
    signed = n['in'] + n['late']
    bar = control_bar([('Group by', seg([('Status', False, None, ''), ('Role', True, None, ''), ('Class', False, None, '')], label='Group by'))], 'Group by role',
                      label='Sign-up grouping')
    own = '<ul>' + SignupRow(crest_img(you['class'], 32), you['name'], C[you['class']], f"{you['spec']} {you['class'].title()}", Segmented([(w, w.lower() == you['state']) for w in SIGNUP_WORDS]),
                             Select(f"{you['spec']} &middot; {ROLE_WORD[you['role']]}")) + '</ul>'
    not_signed = [planner_row({**by[r['name']], 'state': None}, planner_sub(by[r['name']]) + ('' if r['discord_linked'] else ' &middot; No Discord linked'))
                  for r in raid['not_signed_raiders']]
    main = (Section('Your sign-up', own, named=True, panel=True)
            + Section('Sign-ups', f'<div role="region" aria-label="Sign-ups">{roster_groups(raid)}</div>', meta=f"{signed}/{raid['size']}", panel=True, frame='framed')
            + Section('Not signed up', Rows(not_signed, open_=True, list_=True), action=Button('Ping on Discord', small=True, act=True), meta=str(len(not_signed)), panel=True))
    t = raid['comp_target']
    late = n['late_groups']
    bench = [Row(m['name'], planner_sub(m), lead=crest_img(m['class'], 24), title_color=C[m['class']], trail=Button('Swap in', small=True, act=True), link=False, slim=True)
             for m in sorted((by[s['name']] for s in raid['signups'] if s['state'] == 'bench'), key=lambda m: (CLASS_ORDER.index(m['class']), m['name']))]
    enc = raid['assignments']['encounters'][0]
    yours = ', '.join(s['label'] for s in enc['slots'] if any(x['own'] for x in s['names']))
    assign = (Rows(assignment_rows(enc, edited), open_=True, list_=True)
              + ButtonRow(Button('Copy as raid warning', small=True, act=True) + Button('Post to Discord', small=True, act=True)))
    comp = comp_grid({**n, 'late_groups': late}, t) + f'<div style="padding:16px 20px 8px">{Button("Change the target", small=True, act=True)}</div>'
    rail = (Section('Comp against the target', comp, panel=True, flush=True)
            + Section('Bench and swaps', Rows(bench, open_=True, list_=True), meta=str(len(bench)), panel=True)
            + Section(f"{enc['name']} assignments", assign, meta=f'Your job: {yours}', named=True, panel=True)
            + Section('Notes', Prose(raid['notes_html'], 'note'), named=True, panel=True))
    actions = ''.join(Button(x, small=True, act=True) for x in ['Edit', 'Lock sign-ups', 'Copy the sign-up list']) + Button('Cancel raid', small=True, tone='ghost', act=True)
    # The way back is the head's crumb, on the band above the title (the one head contract).
    head = (PageHeader(raid['title'], ALLIANCE, meta=[Tag('Planned'), raid_when(raid), f"<b class=\"strong\">{signed} of {raid['size']} signed up</b> &middot; {n['in']} In &middot; {n['late']} Late &middot; {n['bench']} Bench"],
                         actions=actions, tabs=TabStrip(TABS, 'Raids'), crumb=f'<a href="#">&lsaquo; {GUILD} &middot; Raids</a>'))
    # The guild route's <main> ends 40 px under its view (pages/guild/[...path].astro pb-10).
    body = f'<div style="padding-bottom:40px">{head + bar + Page("split", main, rail, sticky_rail=True, panels=True)}</div>{GUILD_CSS}'
    return build_page('Raid planner, assignments' if edited else 'Raid planner', body, 'Guild', (guild_raid.VIEWER, you['class'], f"{you['spec']} {you['class'].title()} &middot; 60"))


def band_line(viewer: str) -> str:
    """The next raid's strip on the Roster only (batch 3 ruling; raids spec 3.4): on a panel under the head, the sentence with
    the state word, the four sign-up buttons, Open the raid."""
    nxt = F.next_raid(viewer)
    at = guild_time(O.NEXT_RAID_AT)
    _, _, _, role, _, state = ch(viewer)
    word = {'in': 'In', 'late': 'Late', 'bench': 'Bench'}[state]
    n = nxt['counts']['in'] + nxt['counts']['late']
    sentence = f"<p class=\"t-small\">{at.strftime('%A')} {clock_label(at)}, you are <strong>{word}</strong> as {GROUP_ROLE_WORD[role]}, {n} of {nxt['size']} signed up.</p>"
    segs = Segmented([(w, w.lower() == state) for w in ('In', 'Late', 'Bench', 'Out')], small=True).replace('<a href="#" role="radio"', '<button type="button" role="radio"').replace('</a>', '</button>')
    return (f'<div class="pc" style="padding-bottom:13px">' + Plate(f'<div style="display:flex;align-items:center;justify-content:space-between;gap:16px;min-height:60px;padding:8px 20px">'
            f'{sentence}<span style="display:flex;align-items:center;gap:16px">{segs}{Link("Open the raid &rsaquo;")}</span></div>') + '</div>')


def loot_council():
    """The Loot tab as an officer, Vis'kag picked. It lives on the guild page (hash tabs), so its head is the guild's; the night
    chooser is the page's ControlBar; on the current night a picker opens any drop of the loot table; the decision record's
    rows carry no Edit (the award is undone from the drop's own rail) and Export history is a disclosure."""
    nights = F.loot_nights()['nights']
    night = nights[0]
    open_n = sum(1 for d in night['drops'] if d['state'] == 'open')
    table = {i['item_id']: i for i in F.loot_table()['items']}
    drops = []
    for d in night['drops']:
        item = table[d['item_id']]
        users = len(item['candidates']) if d['kind'] == 'gear' else len(item['eligible'])
        what = SLOT_LINE.get((item['inventory_type'], item['item_subclass']), SLOT_LINE.get((item['inventory_type'], None), '')) if d['kind'] == 'gear' else 'Quest item'
        tail = f"Awarded to {d['award']['name']}" if d['award'] else f'{users} candidates'
        drops.append(Row(f'<span role="button" tabindex="0">{d["name"]}</span>', f'{what} &middot; {tail}', lead=icon(d['icon'], 40, f"q{d['quality']}"),
                         trail=Tag('Open', 'gold') if d['state'] == 'open' else Tag('Awarded', 'kill'), selected=d['item_id'] == O.SELECTED_DROP,
                         title_color='#b866f5', link=False))
    picker = (f'<div style="display:flex;flex-direction:column;gap:6px;max-width:360px;margin-top:12px"><span class="sr-only-mock">Open a drop from the loot table</span>'
              f'{Select("Pick a drop")}</div>')
    main = Section(f'{O.BOSS} drops', f'<div>{Rows(drops)}</div>' + picker, meta=f"{day_label(guild_time(night['at']))} &middot; {open_n} open", panel=True)
    drop = F.council_drop()
    cands = []
    for c in drop['candidates']:
        won = f"{c['items_won']} item{'' if c['items_won'] == 1 else 's'} this tier"
        sub = f"+{c['gain_dps']} DPS over {c['replaces']} &middot; {c['attendance']['present']} of {c['attendance']['nights']} raid nights &middot; {won}"
        foot = (Tag(c['ms_os'].upper(), 'gold') + (Tag('BiS for them', 'gold') if c['bis_for_them'] else '')
                + VoteButtons(c['votes']['yes'], c['votes']['no'], c['votes']['voters']) + f'<span style="margin-left:auto">{Button("Award", small=True, act=True)}</span>')
        cands.append(Row(c['name'], sub, lead=crest_img(c['class'], 32), title_color=C[c['class']], foot=foot, link=False))
    record = Rows([Row(f"{d['item_name']} to {d['name']}",
                       f"{d['votes_yes']} yes, {d['votes_no']} no of {d['votes_officers']} &middot; by {d['awarded_by']} &middot; {day_label(guild_time(d['at']))}",
                       lead=icon(d['icon'], 32, f"q{d['quality']}"), link=False) for d in F.decisions()], open_=True)
    rail = (Section(f"Candidates for {drop['name']}", Rows(cands, open_=True), panel=True, frame='framed')
            + Section('Decision record', record, DisclosureButton('Export history'), panel=True))
    bar = control_bar([('Raid night', seg([(day_label(guild_time(n['at'])), n is night, None, '') for n in nights]))], day_label(guild_time(night['at'])), label='Loot council choices')
    return build_page('Loot council', _view(hdr(tabs_active='Loot') + bar + Page('split', main, rail, sticky_rail=True, panels=True)), 'Guild', ('Dunmaro', 'hunter', 'Beast Mastery Hunter &middot; 60'))


def _section_body(html):
    """A Section body as the built Section draws it: 12 px under the head row."""
    return f'<div style="padding-top:12px">{html}</div>'


def raid_report():
    """The done raid as an officer reads it (state `report`), drawn from guild_report_data (the parity fixtures' own data)."""
    raid = rd.raid_view()
    rep, viewer = raid['report'], rd.row(rd.VIEWER)
    shown = rd.default_pull(rep)
    tone = {'kill': ('Kill', 'kill'), 'wipe': ('Wipe', 'wipe'), 'clear': ('Clear', 'muted')}
    rows = []
    for p in rep['pulls']:
        is_open = p is shown
        title = Toggle(rd.pull_title(p), is_open) if p['death_causes'] else rd.pull_title(p)
        rows.append(Row(title, rd.pull_sub(p), trail=Tag(*tone[p['result']]), link=False, selected=is_open))
    deaths = Entries([(f'<b style="color:{C[d["class"]]}">{d["name"]}</b> <span class="t-small">{rd.clock(d["at_ms"])} in</span>',
                       f'<b>{d["killing_blow"]["ability"]}</b> from {d["killing_blow"]["source"]}, {d["killing_blow"]["amount"]:,} '
                       f'({d["killing_blow"]["overkill"]:,} overkill)') for d in shown['death_causes']])
    insp = Inspector(rd.pull_title(shown), [('Deaths', deaths)], rd.pull_title(shown), rd.pull_sub(shown))
    pulls = f'<div role="group" aria-label="Every pull of the night">{Split(Rows(rows), insp)}</div>'
    output = Rows([Row(p['name'], f'DPS &middot; rated {p["rating"]} of 100', lead=crest_img(p['class'], 24), title_color=C[p['class']],
                       trail=Figure(str(p['value']), '', 'DPS', size='sm'), link=True, selected=p['name'] == rd.VIEWER, slim=True) for p in rep['parses']], open_=True)
    # One reason said once, naming every pull it covers (report-night-view unknownLines).
    titles = [rd.pull_title(b) for b in rep['consumables']['by_pull']]
    named = titles[0] if len(titles) == 1 else ', '.join(titles[:-1]) + ' and ' + titles[-1]
    unknown = f'<p class="t-small" style="margin-top:4px">{named}: the log carried no character sheet, so nothing can be said.</p>'
    # The officer's deaths by raider (ReportDeaths), in the night's order, never worst first.
    def _count(w, k):
        if w + k < 2:
            return ''
        parts = ([] if w == 0 else ['1 on a wipe' if w == 1 else f'{w} on wipes']) + ([] if k == 0 else ['1 on a kill' if k == 1 else f'{k} on kills'])
        return ' &middot; '.join(parts)

    def _lines(key):
        return '; '.join(f"{d['killing_blow']['ability']}, {rd.pull_title(p)} at {rd.clock(d['at_ms'])}"
                         for p in rep['pulls'] for d in p['death_causes'] if d['character_key'] == key)
    died = Rows([Row(r['name'], _lines(r['character_key']),
                         lead=crest_img(r['class'], 24), title_color=C[r['class']], link=False,
                         trail=(f'<span style="font-size:13px">{_count(r["on_wipes"], r["on_kills"])}</span>' if _count(r['on_wipes'], r['on_kills']) else ''),
                         slim=True) for r in rep['deaths']],
                 open_=True, list_=True).replace('<ul class="rows', '<ul style="display:grid;grid-template-columns:1fr 1fr;column-gap:32px" class="rows', 1)
    match = raid['match']
    logged = (f'<p class="t-small">Logged as a raid night. {Link("Open the night")}</p><div><p class="t-small">Log uploaded '
              f'{rd.local(match["uploaded_at"]):%a} {rd.local(match["uploaded_at"]).day} Dec, {rd.hour(rd.local(match["uploaded_at"]))} by {match["uploaded_by_name"]}'
              f' &middot; matched by time</p>'
              # RaidMatchLine's Not this log: a quiet underlined text button on the line under the upload.
              '<button type="button" class="t-small" style="all:unset;display:block;font-size:13px;line-height:18px;'
              'text-decoration:underline;cursor:pointer">Not this log</button></div>')
    fix = RowsSlim([Row(a['name'], 'Present' if a['outcome'] == 'present' else 'Out', lead=crest_img(a['class'], 24), title_color=C[a['class']],
                        trail=Select('Fix'), link=False, selected=a['own']) for a in raid['attendance']])
    # The officer's names under the summary (report-attendance-view nightNames): not in the log, late, joined, by name.
    signed_in = {'in', 'late'}
    groups = [('Signed up, not in the log', [x for x in raid['attendance'] if x['outcome'] == 'no_show']),
              ('Late', [x for x in raid['attendance'] if x['outcome'] == 'late']),
              ('Joined from the bench or without a sign-up',
               [x for x in raid['attendance'] if x['outcome'] in ('present', 'late') and x['signed'] not in signed_in])]
    names = ''.join(f'<p class="t-small" style="margin-top:4px">{label}: '
                    + ', '.join(f'<b style="color:{C[x["class"]]}">{x["name"]}</b>' for x in rows) + '</p>'
                    for label, rows in groups if rows)
    att = logged + f'<p>{rd.attendance_summary(raid["night"])}</p>' + names + Disclosure('Fix attendance', fix)
    lootr = ''.join(ItemRow(a['item_name'], a['quality'], icon(a['icon'], 32, f'q{a["quality"]}'), f'{g["boss"]} &middot; to {a["name"]}')
                    for g in rep['loot'] for a in g['awards'])
    lines = rd.discord_lines(raid)
    discord = (''.join(f'<p class="t-small">{"<b>" + x + "</b>" if i == 0 else x}</p>' for i, x in enumerate(lines))
               + f'<div style="display:flex;gap:12px;padding-top:12px">{Button("Post to Discord", primary=True, small=True, act=True)}{Button("Copy for Discord", small=True, act=True)}{Live()}</div>')
    # Round 5 (lane h-edges): every section on its panel, the pulls framed and edge to edge for the Inspector.
    main = (Section('Pulls', pulls, meta=f'{len(rep["pulls"])} fights', panel=True, flush=True)
            + Section('Deaths', died, meta=f'{len(rep["deaths"])} raiders', panel=True) + Section('DPS and HPS', output, meta='Boss kills', panel=True) + Section('Consumables by pull', unknown, panel=True))
    rail = Section('Attendance', att, panel=True) + Section('Loot awarded', lootr, panel=True) + Section('Discord summary', _section_body(discord))  # ReportDiscord: lane f-content's file, still open
    h = rep['headline']
    fig = Figures([(str(h['kills']), 'Kills', '', 'kill'), (str(h['wipes']), 'Wipes', '', 'wipe'), (str(h['attempts']), 'Attempts', '', ''), (str(h['deaths']), 'Deaths', '', '')])
    # One raid page: the crumb back and the guild tabs as the planner's (raid_planner).
    head = PageHeader(raid['title'], '', Emblem(gd.FACTION, 56), meta=[Tag('Done', 'muted'), rd.when_line(raid)], answer=fig,
                      tabs=TabStrip(TABS, 'Raids'), crumb=f'<a href="#">&lsaquo; {GUILD} &middot; Raids</a>')
    return build_page('Raid report', _view(head + Page('split', main, rail, sticky_rail=True, panels=True)), 'Guild',
                      (rd.VIEWER, viewer[1], f'{viewer[2]} {viewer[1].title()} &middot; 60'))


# ---- Recruitment, applications and the member profile (lane D). Drawn from guild_data (the values guild_fixtures.recruit_fixtures
# serves the build). The build renders actions as <button>, forms as <form>, the attendance Grid as a <table>: the helpers below
# give the ws_prims primitives the same roles, so the parity gate counts the same blocks.
def _act(label, primary=False, small=True, submit=False):
    """The system Button as the build renders an action (a <button>; ws_prims draws an <a>)."""
    kind = 'submit' if submit else 'button'
    return Button(label, primary, small).replace('<a ', f'<button type="{kind}" ', 1).replace(' href="#"', '', 1).replace('</a>', '</button>')


def _votes(yes, no, total):
    """Votes as an officer casts them: the yes and no Tags are the vote buttons (ui/Votes with onVote)."""
    return Votes(yes, no, total).replace('<span class="pill tag', '<button type="button" class="vote"><span class="pill tag').replace(' no</span>', ' no</span></button>', 1).replace(' yes</span>', ' yes</span></button>', 1)


def _form(label, fields, submit, floor=True):
    """FormRow as the build renders it: a <form> of labelled text inputs and one submit Button. `floor` False is a field
    that takes what the row leaves beside the button (OfficerNotes.svelte), not the Field's 220 px grow floor."""
    cls = 'f grow' if floor else 'f" style="flex:1;min-width:0'
    f = ''.join(f'<label class="{cls}"><span class="t-label">{lab}</span><input class="field grow" placeholder="{hint}"></label>' for lab, hint in fields)
    return f'<form class="formrow" aria-label="{label}">{f}<span>{_act(submit, submit=True)}</span></form>'


def _grid_table(caption, cols, rows):
    """Grid as the build renders it (ui/Grid, a <table>): the same cells, with table roles."""
    head = '<div role="row"><div role="columnheader" class="h"><span class="sr-only">' + caption + '</span></div>' + ''.join(f'<div role="columnheader" class="h" style="text-align:center">{c}</div>' for c in cols) + '</div>'
    body = ''.join(f'<div role="row"><div role="rowheader" class="l"><div class="ttl">{lab}</div></div>' + ''.join(f'<div role="cell" class="c {st}">{tx}</div>' for st, tx in cells) + '</div>' for lab, cells in rows)
    return (f'<div class="gridm" role="table" aria-label="{caption}" style="--n:{len(cols)}"><div role="caption" class="sr-only">{caption}</div>'
            f'<div role="rowgroup">{head}</div><div role="rowgroup">{body}</div></div>')


_ROLE_CSS = ('<style>.sr-only{position:absolute;width:1px;height:1px;padding:0;margin:-1px;overflow:hidden;clip:rect(0,0,0,0);white-space:nowrap}'
             '.gridm [role=rowgroup],.gridm [role=row]{display:contents}button.vote{all:unset;display:inline-flex;cursor:pointer}form.formrow input.field{font:inherit}</style>')


def _body(html):
    """A Section's body as the built Section spaces it (ui-section-body, 12 px under the head) where the body is not a Plate."""
    return f'<div style="padding-top:12px">{html}</div>'


def _view(html):
    """GuildShell keeps every guild view one screen tall (GUILD_VIEW_MIN), so the footer starts below the fold; the guild page's
    main pads 40 px under it (pages/guild/[...path].astro pb-10)."""
    return f'<div style="min-height:calc(100vh - 81px);margin-bottom:40px">{html}</div>{GUILD_CSS}'


# The built ghost Button (lib/ui/button.ts BUTTON_GHOST): a gold label, no fill, 10 px each side.
GUILD_CSS = '<style>.sr-only-mock{position:absolute;width:1px;height:1px;overflow:hidden;clip:rect(0 0 0 0)}html :is(a,button).btn.ghost{background:transparent;border-color:transparent;color:var(--gold);padding:0 10px;font-size:15px}</style>'


_PROSE = 'style="font-size:14px;line-height:22px"'


def _day(iso):
    """'Wed 16 Dec' in the guild's zone (raid-time dateLabel)."""
    d = datetime.fromisoformat(iso.replace('Z', '+00:00')).astimezone(ZoneInfo(gd.ZONE))
    return f'{d:%a} {d.day} {d:%b}'


def _clock(iso):
    d = datetime.fromisoformat(iso.replace('Z', '+00:00')).astimezone(ZoneInfo(gd.ZONE))
    return f'{d.hour % 12 or 12}:{d:%M} {"am" if d.hour < 12 else "pm"}'


def _spec_class(name):
    r = ch(name)
    return f'{r[2]} {r[1].title()}'


def _wanted_summary():
    return ', '.join(n for _c, _k, n, _note in gd.RECRUIT_WANTED)


def _top(metric, value):
    return f'Top {metric.upper()} {round(value)}'


def recruitment():
    """The public card (batch 3, b3-guild-recruit): Join and Recruiting on the band; Looking for framed, then Apply; the rail's
    guild facts (ui/Facts), why this guild and the officers."""
    wanted = Rows([Row(n, note, lead=crest_img(c, 30), title_color=C[c], link=False) for c, _k, n, note in gd.RECRUIT_WANTED], list_=True)
    # The viewer's clock as the build prints it (raid-time clockText): the minutes only when they are not :00.
    upcoming = ', '.join(f"{_day(at)}, {clock_label(guild_time(at)).rsplit(' ', 1)[0]}" for at in gd.RECRUIT_UPCOMING)
    t = F.tier()
    facts = [('Region', f'{gd.REGION.upper()} PvP'), ('Faction', gd.FACTION.title()), ('Members', str(gd.RAIDERS)), ('Raids', gd.RECRUIT_NIGHTS),
             ('Next raid nights', upcoming), ('Progress', f"{t['name']} &middot; {t['down']} of {t['named_encounters']} boss{'' if t['named_encounters'] == 1 else 'es'} down")]
    # The progress line sits on one line in the build at 1440 (its face sets a hair narrower than the mock's).
    card = '<dl class="gfacts">' + ''.join(f'<dt>{k}</dt><dd{" style=white-space:nowrap" if k == "Progress" else ""}>{v}</dd>' for k, v in facts) + '</dl>'
    # Signed out nobody can apply: the API takes one of the viewer's own characters, so the build asks them to sign in (SignInPrompt).
    sign_in = Button('Sign in with Battle.net to apply', primary=True)
    rank = {'leader': 'Guild master', 'officer': 'Officer'}  # recruit-copy card.leader / card.officer
    officers = Rows([Row(n, f'{_spec_class(n)} &middot; {rank[gd.rank_of(n)]}', lead=crest_img(ch(n)[1], 32), title_color=C[ch(n)[1]], link=False)
                     for n in gd.COUNCIL], open_=True)
    main = Section('Looking for', wanted, panel=True, frame='framed') + Section('Apply', sign_in, panel=True)
    rail = (Section(GUILD, card, panel=True) + Section('Why this guild', f'<p {_PROSE}>{gd.RECRUIT_PITCH}</p>', panel=True)
            + Section('Officers', officers, panel=True))
    head = PageHeader(f'Join {GUILD}', '', Emblem(gd.FACTION, 56), meta=[Tag('Recruiting', 'kill')])
    return build_page('Recruitment', _view(head + Page('split', main, rail, panels=True)) + _FACTS_CSS, 'Guild', None)


# ui/Facts: the key muted at the left, the value strong at the right, a hairline between pairs.
_FACTS_CSS = ('<style>.gfacts{display:grid;grid-template-columns:minmax(0,auto) minmax(0,1fr);margin:0}'
              '.gfacts dt,.gfacts dd{margin:0;padding:10px 0;border-bottom:1px solid var(--line-soft);font-size:14px;line-height:20px}'
              '.gfacts dt{color:var(--muted);padding-right:16px;font:400 14px/20px var(--font-body)}'
              '.gfacts dd{text-align:right;color:var(--strong);font:600 14px/20px var(--font-body);font-variant-numeric:tabular-nums}</style>')


def applications():
    """Applications (batch 3, b3-guild-applications): the guild's band with the Roster's sub-views, Open applications framed with
    Accept and a quiet Decline; the rail's public card editor and the recent decisions."""
    apps = []
    for _i, name, cls, spec, il, parse, msg, avail, y, n, _mine, _sent in gd.APPLICANTS:
        line = msg + (f' &middot; Can raid: {avail}' if avail else '')
        apps.append(Row(name, f'{spec} {cls.title()} &middot; Level 60 &middot; Item level {il} &middot; {_top(parse[0], parse[1])}<br>{line}', lead=crest_img(cls, 36), title_color=C[cls],
                        trail=_votes(y, n, gd.VOTERS) + _act('Accept') + Button('Decline', small=True, tone='ghost', act=True), link=False))
    main = Section('Open applications', f'<div role="region" aria-label="Open applications">{Rows(apps)}</div>', panel=True, frame='framed')
    card = Rows([Row('Looking for', _wanted_summary(), trail=_act('Edit'), link=False), Row('Raid nights', gd.RECRUIT_NIGHTS, trail=_act('Edit'), link=False)], open_=True)
    decided = Rows([Row(f'{name} accepted', f'{y} yes, 0 no of {gd.VOTERS} officers &middot; by {by}', link=False, slim=True)
                    for _i, name, _c, _s, y, by, _r, _d in gd.DECIDED], open_=True)
    rail = Section('Your public card', card, Link('See it as a visitor'), panel=True) + Section('Recent decisions', decided, panel=True)
    views = TabStrip(['Members', 'Applications'], 'Applications').replace('strip-tabs ph-tabs', 'strip-tabs', 1)
    head = PageHeader(GUILD, '', Emblem('alliance', 56), meta=[band_facts(None)], tabs=TabStrip(TABS, 'Roster') + views)
    return build_page('Applications', _view(head + Page('split', main, rail, panels=True)) + _ROLE_CSS, 'Guild', (gd.VIEWER, 'hunter', 'Beast Mastery Hunter &middot; 60'))


def _ready_facts(ilvl):
    """Raid ready (profile-view readyFacts): the gear gap in the BiS unit, enchants, bags, talents, attunement, item level against the
    guild's median and the last export, as ui/Facts draws them (the sub line muted under the value)."""
    gap = gd.PROFILE_GAP
    sub = lambda text: f'<small class="t-small" style="display:block;font-weight:400">{text}</small>'
    rows = [('Gear', f'{gap["upgrades"]} upgrades' + sub(f'+{gap["gain_dps"]:.1f} {gap["unit"].upper()}')),
            ('Enchants', 'All enchanted'), ('Bags', 'Flask and food'), ('Talents', 'All spent'), ('Attunement', 'Attuned'),
            # At the guild median, the item level reads the figure alone (readyFacts).
            ('Item level', f'{ilvl}'), ('Top fight', 'Onyxia' + sub(f'rated {gd.PROFILE_PARSE[2]} of 100')),
            ('Last logged', _day(gd.NOW))]
    return '<dl class="gfacts">' + ''.join(f'<dt>{k}</dt><dd>{v}</dd>' for k, v in rows) + '</dl>'


def member_profile():
    """The member profile as an officer reads it (round 5, lane h-edges): the raider in class colour over the guild's tabs (Roster),
    Raid ready framed, the last raids and loot this tier; the rail holds their characters, then Officer notes."""
    name, cls, spec, _role, ilvl, _st = ch(gd.PROFILE)
    from guild_fixtures import profile_loot
    won = [(d['item_id'], d['item_name'], d['icon'], d['quality'], d['at']) for d in profile_loot()]
    present = sum(1 for _at, o in gd.PROFILE_NIGHTS if o == 'present')
    figs = Figures([(f'{present} of {len(gd.PROFILE_NIGHTS)}', 'Attendance', '', ''), (str(round(gd.PROFILE_PARSE[1])), f'Top {gd.PROFILE_PARSE[0].upper()}', '', ''),
                    (str(ilvl), 'Item level', '', ''), (str(len(won)), 'Items this tier', '', '')])
    mark = {'present': ('ok', '&#10003;'), 'late': ('warn', 'L'), 'absent': ('miss', '&times;')}
    last = 'Last raid' if len(gd.PROFILE_NIGHTS) == 1 else f'Last {len(gd.PROFILE_NIGHTS)} raids'  # recruit-copy lastNights
    att = (_grid_table(last, [_day(at) for at, _o in gd.PROFILE_NIGHTS], [(last, [mark[o] for _at, o in gd.PROFILE_NIGHTS])])
           + f'<p style="padding-top:12px;font-size:14px;line-height:22px">{gd.PROFILE_LAST_30[0]} of {"1 raid" if gd.PROFILE_LAST_30[1] == 1 else f"the last {gd.PROFILE_LAST_30[1]} raids"}</p>')
    loot = '<ul>' + ''.join(f'<li>{ItemRow(n, q, icon(ic, 32, f"q{q}"), f"Awarded {_day(at)}")}</li>' for _i, n, ic, q, at in won) + '</ul>'
    # The notes panel says once that members never see them: no chip per note, no placeholder; Delete is the quiet danger button.
    delete = Button('Delete', small=True, tone='danger').replace('<a ', '<button type="button" ', 1).replace(' href="#"', '', 1).replace('</a>', '</button>')
    notes = Rows([Row(body, f'{who} &middot; {_day(at)}', trail=delete if mine else '', link=False) for _i, body, who, at, mine in gd.PROFILE_NOTES], open_=True)
    addn = f'<div style="padding-top:12px">{_form("Add a note", [("Add a note", "")], "Save note", floor=False)}</div>'
    characters = Rows([Row(n, lead=crest_img(c, 28), title_color=C[c], trail=Tag('Main', 'muted') if main else '', slim=True)
                       for n, c, main in [(name, cls, True), (*gd.PROFILE_ALT, False)]], list_=True)
    # Raid ready's roll-up tag shows only when a check fails (readyTag); Frostquill passes every one.
    ready = Section('Raid ready', _ready_facts(ilvl), Link('Open their character'), panel=True)
    main = (ready + Section('Attendance', att, panel=True)
            + Section('Loot this tier', loot, panel=True))
    rail = Section('Characters', characters, panel=True) + Section('Officer notes', notes + addn, panel=True)
    # No "member since" (the roster carries no join date): the rank instead.
    head = PageHeader(name, C[cls], crest_img(cls, 56), meta=[' &middot; '.join([f'{spec} {cls.title()}', 'Level 60', 'Member'])], answer=figs,
                      tabs=TabStrip(TABS, 'Roster'))
    return build_page('Member profile', _view(head + Page('split', main, rail, panels=True)) + _ROLE_CSS + _FACTS_CSS, 'Guild',
                      (gd.VIEWER, 'hunter', 'Beast Mastery Hunter &middot; 60'))


PAGES = {'web-guild-member': landing_member, 'web-guild-officer': landing_officer, 'web-guild-raid': raid_planner, 'web-guild-loot': loot_council, 'web-guild-report': raid_report,
         'web-guild-recruit': recruitment, 'web-guild-applications': applications, 'web-guild-profile': member_profile}

def all_pages() -> dict:
    """Every guild mock: the persona round's pages here and the GUILD-ALL lane's, one module per page group."""
    import ws_guild_directory
    import ws_guild_raids
    import ws_guild_settings
    import ws_guild_states
    import ws_guild_tabs
    return {**PAGES, **ws_guild_states.PAGES, **ws_guild_tabs.PAGES, **ws_guild_raids.PAGES, **ws_guild_settings.PAGES, **ws_guild_directory.PAGES}


if __name__ == '__main__':
    import sys
    for n, f in all_pages().items():
        if not sys.argv[1:] or n in sys.argv[1:]:
            print(write(n + '.html', f()))
