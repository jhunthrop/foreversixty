"""Guild persona round: role-aware landing (member and officer), raid planner, loot council, raid report, recruitment, applications and a
member profile. Composed only from ws_prims primitives. Data shapes: api/cmd/seedguild (24 characters), api/internal/guilds (sign-up states
In / Late / Bench, loot candidates, readiness, attendance, raid report). Modules whose data source is not confirmed carry a Gap marker.
"""
from ws_prims import *  # noqa: F403
from ws_prims import (Page, PageHeader, Section, Plate, Rows, Row, Figures, Figure, Tag, Button, Link, TabStrip, FormRow, Field, Select, Notice, Segmented, Grid, Tally,
                      Checklist, Votes, Gap, build_page, crest_img, icon, write, Emblem)

from datetime import datetime
from zoneinfo import ZoneInfo

import guild_overview_data as O
import guild_overview_fixtures as F
from guild_data import ALLIANCE, GUILD, NOW, TABS, ZONE, C, R, STATE_TAG  # noqa: F401
from ws_prims import SpellIcon, control_bar, seg

# Figures under a Section head sit 12 px below its hairline, as the built Section body does (padding-top on a layout div).
# Lane A's words (the built page's own): the raid's icon (lib/first-raids.ts), the sign-up role words (raidCopy.roleLabels), the
# BiS slot words (readiness-view enchantSlotLabel) and the loot table's slot lines (slotLine) for the Onyxia drops shown.
RAID_ICON = 'inv_misc_head_dragon_black'
ROLE_WORD = {'Tank': 'Tank', 'Healer': 'Healer', 'Melee': 'DPS', 'Ranged': 'DPS'}
SLOT_WORD = {'hands': 'Hands', 'finger1': 'Finger 1'}
SLOT_LINE = {(13, 7): 'One-Hand &middot; Sword', (13, 0): 'One-Hand &middot; Axe', (2, None): 'Neck', (11, None): 'Finger', (16, None): 'Back'}


def ch(name):
    return next(r for r in R if r[0] == name)


def lead(name, size=32):
    r = ch(name)
    return crest_img(r[1], size)


def hdr(title_suffix='', tabs_active='Overview', actions='', notice='', officer=True):
    """The guild's head as the built Guild.svelte draws it: the faction crest (64 px, guild page spec 12.2), the facts line
    (ruleset and region, the tier's counts, Updated, the next raid; a member also reads who holds the claim), the role's tabs
    as a tablist (hash tabs on one URL), an officer's actions."""
    claim = [] if officer else [f'Claimed by {crest_img(ch(O.OFFICER)[1], 20)} <b style="color:{C[ch(O.OFFICER)[1]]}">{O.OFFICER}</b>']
    tabs = TABS if officer else [t for t in TABS if t != 'Settings']
    return PageHeader(GUILD, ALLIANCE, Emblem('alliance', 64), meta=[' &middot; '.join(head_facts())] + claim, actions=actions,
                      tabs=TabStrip(tabs, tabs_active, tablist=True)) + notice


def head_facts() -> list[str]:
    t = F.tier()
    nxt = guild_time(O.NEXT_RAID_AT)
    return ['PvP US', f"{t['down']} boss down", f"{t['pulls_total']} pulls", 'Updated today', f"Next raid: {O.ZONE_NAME}, {day_label(nxt)}"]


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
    """The next-raid Plate: the raid's Row (icon, the guild's time, the countdown), the viewer's own Row with In / Late /
    Bench / Out (the built sign-up's four states), and for an officer the In, Late, Bench and Not signed figures."""
    at = guild_time(O.NEXT_RAID_AT)
    top = Row(O.ZONE_NAME, f'{day_label(at)} &middot; {clock_label(at)}', lead=icon(RAID_ICON, 40, 'q5'),
              trail=Figure(countdown_label(), 'Starts in', tone='gold', size='sm'), link=False)
    name = O.OFFICER if officer else O.MEMBER
    _, cls, spec, role, _, state = ch(name)
    word = {'in': 'In', 'late': 'Late', 'bench': 'Bench'}[state]
    you = Row(name, f'{spec} {cls.title()} &middot; level 60 &middot; you are <strong>{word}</strong> as {ROLE_WORD[role]}', lead=crest_img(cls, 36),
              title_color=C[cls], trail=Segmented([(w, w.lower() == state) for w in ('In', 'Late', 'Bench', 'Out')], small=True), link=False)
    tally = ''
    if officer:
        view = F.raid_view(name, True)
        c = view['counts']
        tally = '<div style="padding:12px 16px">' + Figures([(str(c['in']), 'In', '', 'kill', 'sm'), (str(c['late']), 'Late', '', 'gold', 'sm'), (str(c['bench']), 'Bench', '', '', 'sm'),
                                                             (str(len(view['not_signed_raiders'])), 'Not signed', '', 'wipe', 'sm')]) + '</div>'
    return Plate(Rows([top, you]) + tally)


def comp_tally() -> str:
    view = F.raid_view(O.MEMBER, False)
    c, late = view['counts'], view['counts']['late_groups']
    have = {'tanks': c['tanks'], 'healers': c['healers'], 'melee': c['melee'], 'ranged': c['ranged']}
    return Tally([(role.title(), have[role], O.COMP_TARGET[role], f'{late[role]} late' if late[role] else '') for role in ('tanks', 'healers', 'melee', 'ranged')])


def comp_meta() -> str:
    return f'{O.ZONE_NAME} default: {sum(O.COMP_TARGET.values())} players'


def last_raid(meta_line: str) -> str:
    counts = O.last_raid_counts()
    at = guild_time(O.LAST_RAID_AT)
    facts = [day_label(at), f'{counts["kills"]} boss ({O.BOSS})', f'{counts["kills"]} kill', f'{counts["wipes"]} wipe']
    row = Row(f'{O.ZONE_NAME} &middot; raid night', ' &middot; '.join(facts), trail='<span class="gold" style="font-weight:600">Report &rsaquo;</span>')
    return Section('Last raid', Rows([row], open_=True), meta=meta_line)


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
    main = Section('Next raid', next_raid(False)) + Section('Before the raid', Checklist(rows), meta='Private to you &middot; Almost there') + Section('Your numbers', f'<div style="padding-top:12px">{nums}</div>')
    rail = Section("The raid's comp", comp_tally(), meta=comp_meta()) + last_raid(progression_line())
    return build_page('Guild, member', hdr(officer=False) + notice + Page('split', main, rail, sticky_rail=True), 'Guild', ('Hollowmend', 'priest', 'Holy Priest &middot; 60'))


def cover_grids() -> str:
    per = {'raid': '', 'melee_group': 'per melee group', 'caster_group': 'per caster group'}
    out = ''
    for kind, word in (('buff', 'Buffs'), ('debuff', 'Debuffs')):
        rows = []
        for r in (x for x in F.cover_rows() if x['kind'] == kind):
            signed = r['providers_in'] + r['providers_late']
            tone = 'miss' if signed == 0 else ('ok' if signed - r['likely'] >= r['need'] else 'warn')
            sub = ' &middot; '.join(x for x in (per[r['per']], f"{r['likely']} on a talent we cannot see" if r['likely'] else '',
                                                 f"{r['providers_bench']} on the bench" if r['providers_bench'] else '') if x)
            rows.append((r['name'], sub, [('none', str(r['need'])), (tone, str(signed))], SpellIcon(r['name'], r['icon'])))
        out += Grid(['Need', 'In'], rows, caption=f'{word}: need against signed-up raiders', label_header=word)
    return out


def raid_ready() -> str:
    rows = F.readiness(True)['rows']
    att = sum(1 for r in rows if r['attuned'])
    ff = sum(1 for r in rows if r['flask_food']['flask_or_elixir'] and r['flask_food']['food'])
    tanks = [r for r in F.raid_view(O.OFFICER, True)['signups'] if r['role'] == 'tank' and r['state'] in ('in', 'late')]
    figs = Figures([(f'{att} of {len(rows)}', 'Attuned', '', '', 'sm'), (f'{ff} of {len(rows)}', 'Flask and food', '', '', 'sm')])
    label = (f'<div style="display:flex;gap:12px;align-items:baseline;padding-top:8px"><span class="t-label">Fire resistance on tanks</span>'
             f'<span class="t-small">{len(tanks)} tank signed</span></div>')
    trows = Rows([Row(t['name'], trail=f"<b>{O.TANK_FIRE.get(t['name'], 0)} fire resistance</b>", title_color=C[t['class']], slim=True, link=False) for t in tanks], open_=True)
    return f'<div style="padding-top:12px">{figs}</div>' + label + trows


def attendance_grid() -> str:
    cell = {'present': ('ok', '&#10003;'), 'absent': ('miss', '&times;'), 'bench': ('warn', 'B'), 'none': ('none', '&mdash;')}
    page = F.recent_attendance()
    cols = [guild_time(r['starts_at']).strftime('%a') for r in page['raids']]
    rows = sorted(page['rows'], key=lambda r: (not r['slipping'], r['name']))
    return Grid(cols, [(r['name'], 'slipping' if r['slipping'] else 'bench', [cell[c] for c in r['cells']]) for r in rows], caption='Attendance over the last raids')


def action_rows() -> str:
    a = F.officer_actions()
    weekday = guild_time(O.NEXT_RAID_AT).strftime('%A')
    rows = [Row(f"{a['approvals']['count']} raiders are waiting for approval", ', '.join(a['approvals']['names']), trail=Button('Review', small=True, action=True), link=False),
            Row(f"{a['applications']['count']} applications", ', '.join(a['applications']['names']), trail=Button('Open', small=True), link=False),
            Row(f"{a['not_signed']['count']} have not signed up for {weekday}", ', '.join(a['not_signed']['names']), trail=Button('Ping on Discord', small=True, action=True), link=False),
            Row(f"{a['drops_to_award']['count']} drops to award", a['drops_to_award']['label'], trail=Button('Open loot', small=True, action=True), link=False)]
    return Rows(rows, open_=True)


def landing_officer():
    """The Overview as an officer. Since the persona round: the cover is the curated catalogue (D15) in two Grids, Buffs and
    Debuffs, every entry; Raid-ready counts attunement and flask and food and lists each signed tank's fire resistance; the
    rail closes with Recruitment (slice 5)."""
    rec = Rows([Row('Taking applications', 'Applications, votes and your public card')], open_=True)
    main = (Section('Action needed', action_rows()) + Section('Next raid', next_raid(True), Link('Open the planner &rsaquo;'))
            + Section('Comp against the target', comp_tally(), Link('Edit target &rsaquo;'), meta=comp_meta()))
    rail = (Section('Buff and debuff cover', cover_grids(), meta='Signed-up raiders') + Section('Raid-ready', raid_ready(), meta='Signed-up raiders')
            + Section(f'Attendance, last {len(O.RECENT_RAIDS)} raids', attendance_grid(), Link('Roster &rsaquo;')) + last_raid(progression_line())
            + Section('Recruitment', rec))
    return build_page('Guild, officer', hdr(actions=Button('Invite link', small=True)) + Page('split', main, rail, sticky_rail=True), 'Guild', ('Dunmaro', 'hunter', 'Beast Mastery Hunter &middot; 60'))


def raid_planner():
    bar = control_bar([('Group by', seg([('Status', True, None, ''), ('Role', False, None, ''), ('Class', False, None, '')]))], 'Group by status', end=Button('Copy sign-up list', small=True))
    rows = []
    for name, cls, spec, role, il, st in R:
        tg, tone = STATE_TAG[st]
        rows.append(Row(name, f'{spec} {cls.title()} &middot; {role.lower()} &middot; ilvl {il}', lead=crest_img(cls, 32), title_color=C[cls], trail=Tag(tg, tone) + Segmented([('In', st == 'in'), ('Late', st == 'late'), ('Bench', st == 'bench')], small=True), link=False))
    roster = Plate(Rows(rows))
    comp = Tally([(crest_img('warrior', 22), 'Tanks', 1, 2), (crest_img('paladin', 22), 'Healers', 6, 8), (crest_img('rogue', 22), 'Melee', 5, 10), (crest_img('mage', 22), 'Ranged', 11, 20)])
    cls_tally = Figures([('5', 'Warrior', '', ''), ('3', 'Paladin', '', ''), ('3', 'Hunter', '', '')])
    cover = Grid(['Need', 'In'], [('Windfury Totem', '', [('ok', '2'), ('ok', '3')]), ('Blessing of Kings', '', [('ok', '1'), ('warn', '1')]), ('Curse of the Elements', '', [('ok', '1'), ('miss', '0')]), ('Prayer of Fortitude', '', [('ok', '1'), ('ok', '4')])])
    bench = Rows([Row('Cinderglass', 'Fire Mage &middot; benched', lead=crest_img('mage', 28), trail=Button('Swap in', small=True), link=False), Row('Nightlatch', 'Subtlety Rogue &middot; benched', lead=crest_img('rogue', 28), trail=Button('Swap in', small=True), link=False)], open_=True)
    assign = Rows([Row('Tanks', 'Thornwicke on Onyxia; off-tank needed', trail=Link('Assign'), link=False), Row('Interrupts', 'Not assigned', trail=Link('Assign'), link=False), Row('Heal groups', 'Not assigned', trail=Link('Assign'), link=False)], open_=True)
    main = Section('Sign-ups', roster, meta='19 in &middot; 2 late &middot; 2 bench &middot; 3 not signed')
    rail = Section('Comp against the target', comp, meta='Onyxia default, editable') + Section('Buff cover', cover) + Section('Bench and swaps', bench) + Section('Assignments', assign, meta='Tanks, interrupts, heal groups')
    head = PageHeader("Onyxia's Lair", ALLIANCE, icon('inv_misc_head_dragon_01', 56, 'q5'), meta=['Thu 10 Dec &middot; 8:00 pm CST', '9:00 pm EST your time', '23 of 40 signed up'], actions=Button('Post to Discord', small=True))
    return build_page('Raid planner', head + bar + Page('split', main, rail, sticky_rail=True), 'Guild', ('Dunmaro', 'hunter', 'Beast Mastery Hunter &middot; 60'))


def band_line(viewer: str) -> str:
    """The guild band's raid line under the head on every tab but the Overview (raids spec 3.4): the sentence with the state
    word, the four sign-up buttons, the raid's link."""
    nxt = F.next_raid(viewer)
    at = guild_time(O.NEXT_RAID_AT)
    _, _, _, role, _, state = ch(viewer)
    word = {'in': 'In', 'late': 'Late', 'bench': 'Bench'}[state]
    n = nxt['counts']['in'] + nxt['counts']['late']
    sentence = f"<p class=\"t-small\">{at.strftime('%A')} {clock_label(at)}, you are <strong>{word}</strong> as {ROLE_WORD[role]}, {n} of {nxt['size']} signed up.</p>"
    segs = Segmented([(w, w.lower() == state) for w in ('In', 'Late', 'Bench', 'Out')], small=True).replace('<a href="#" role="radio"', '<button type="button" role="radio"').replace('</a>', '</button>')
    return (f'<div class="pc"><div style="display:flex;align-items:center;justify-content:space-between;gap:16px;min-height:44px;padding-top:12px">'
            f'{sentence}<span style="display:flex;align-items:center;gap:16px">{segs}{Link(f"{O.ZONE_NAME} &rarr;")}</span></div></div>')


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
    picker = (f'<div style="display:flex;flex-direction:column;gap:6px;max-width:360px;margin-top:12px"><span class="t-label">Open a drop from the loot table</span>'
              f'{Select("Pick a drop")}</div>')
    main = Section(f'{O.BOSS} drops', Plate(Rows(drops)) + picker, meta=f"{day_label(guild_time(night['at']))} &middot; {len(night['drops'])} drops, {open_n} open &middot; Recorded by the addon")
    drop = F.council_drop()
    cands = []
    for c in drop['candidates']:
        won = f"{c['items_won']} item{'' if c['items_won'] == 1 else 's'} this tier"
        sub = f"+{c['gain_dps']} DPS over {c['replaces']} &middot; {c['attendance']['present']} of {c['attendance']['nights']} raid nights &middot; {won}"
        foot = (Tag(c['ms_os'].upper(), 'gold') + (Tag('BiS for them', 'gold') if c['bis_for_them'] else '')
                + Button(f"{c['votes']['yes']} yes", small=True, action=True) + Button(f"{c['votes']['no']} no", small=True, action=True)
                + f'<span class="t-small">of {c["votes"]["voters"]} officers</span>' + Button('Award', small=True, action=True))
        cands.append(Row(c['name'], sub, lead=crest_img(c['class'], 32), title_color=C[c['class']], foot=foot, link=False))
    record = Rows([Row(f"{d['item_name']} to {d['name']}",
                       f"Main spec &middot; {d['votes_yes']} yes, {d['votes_no']} no of {d['votes_officers']} &middot; by {d['awarded_by']} &middot; {day_label(guild_time(d['at']))}",
                       lead=icon(d['icon'], 32, f"q{d['quality']}"), link=False) for d in F.decisions()], open_=True)
    rail = (Section(f"Candidates for {drop['name']}", Rows(cands, open_=True), meta='Gain as on the BiS page')
            + Section('Decision record', record, Button('Export history', small=True, action=True, expanded=False)))
    bar = control_bar([('Raid night', seg([(day_label(guild_time(n['at'])), n is night, None, '') for n in nights]))], day_label(guild_time(night['at'])), label='Loot council choices')
    return build_page('Loot council', hdr(tabs_active='Loot', actions=Button('Invite link', small=True)) + band_line(O.OFFICER) + bar + Page('split', main, rail, sticky_rail=True), 'Guild', ('Dunmaro', 'hunter', 'Beast Mastery Hunter &middot; 60'))


def raid_report():
    pulls = Plate(Rows([Row('Onyxian Whelp pack', 'Trash &middot; 1:00', trail=Tag('Clear', 'kill'), link=False), Row('Onyxia, pull 1', '4:00 &middot; 6 deaths &middot; dead to Deep Breath ×3, Tail Sweep ×2, Flame Breath', trail=Tag('Wipe', 'wipe'), link=False),
                        Row('Onyxian Whelp pack', 'Trash &middot; 0:55', trail=Tag('Clear', 'kill'), link=False), Row('Onyxia, pull 2', '3:45 &middot; 1 death &middot; Flame Breath', trail=Tag('Kill', 'kill'), link=False)]))
    fig = Figures([('1', 'Kills', '', 'kill'), ('1', 'Wipes', '', 'wipe'), ('2', 'Attempts', '', ''), ('7', 'Deaths', '', '')])
    parse = Rows([Row('Frostquill', 'Frost Mage &middot; 3:45', lead=crest_img('mage', 32), title_color=C['mage'], trail=Tag('96', 'gold') + Figure('812', '', 'DPS', size='sm'), link=False),
                  Row('Doomwhisper', 'Affliction Warlock', lead=crest_img('warlock', 32), title_color=C['warlock'], trail=Tag('88', 'gold') + Figure('744', '', 'DPS', size='sm'), link=False),
                  Row('Dunmaro', 'Beast Mastery Hunter', lead=crest_img('hunter', 32), title_color=C['hunter'], trail=Tag('61', 'muted') + Figure('618', '', 'DPS', size='sm'), link=False)], open_=True)
    cons = Grid(['Pull 1', 'Pull 2'], [('Flask', '', [('none', '?'), ('none', '?')]), ('Food', '', [('none', '?'), ('none', '?')]), ('Potion', '', [('none', '?'), ('none', '?')])]) + f'<div style="padding-top:8px">{Gap("The log carries no buff events yet: needs a buff-event parser")}</div>'
    att = Rows([Row('22 of 24 on the roster were there', 'Read from the log. Voidcall and Starnettle were not.', trail=Link('Fix attendance'), link=False)], open_=True)
    lootr = Rows([Row('Ring of Binding to Doomwhisper', 'Awarded 8 Dec', lead=icon('inv_jewelry_ring_13', 36, 'q4'), link=False)], open_=True)
    discord = Rows([Row('Olympus XXVII: Onyxia down', 'Onyxia killed on the second pull, 3:45. 22 of 24 there. Ring of Binding to Doomwhisper.', trail=Button('Post to Discord', small=True), link=False)], open_=True)
    main = Section('Pulls', pulls, meta='4 fights') + Section('Parses', parse, meta='Best of the night') + Section('Consumables by pull', cons)
    rail = Section('Attendance', att) + Section('Loot awarded', lootr) + Section('Discord summary', discord)
    head = PageHeader('Onyxia, Tue 8 Dec', ALLIANCE, icon('inv_misc_head_dragon_01', 56, 'q5'), meta=["Onyxia's Lair", '1 boss &middot; 1 kill &middot; 2 trash packs &middot; 1 wipe'], answer=fig)
    return build_page('Raid report', head + Page('split', main, rail, sticky_rail=True), 'Guild', ('Dunmaro', 'hunter', 'Beast Mastery Hunter &middot; 60'))


def recruitment():
    card = Plate(Rows([Row(GUILD, 'US &middot; PvP &middot; Alliance &middot; 24 raiders', lead=Emblem('alliance', 40), trail=Tag('Recruiting', 'kill'), link=False),
                       Row('Raids', 'Tue and Thu, 8:00 pm to 11:00 pm CST', link=False), Row('Progress', '1 of 1 named encounters down', link=False), Row('Looking for', '1 tank, 2 healers', link=False)]))
    form = FormRow([('Character', Field('Character name', grow=True), True), ('Spec', Select('Pick a spec'), False), ('Raid experience', Field('Previous guilds and tiers', grow=True), True)], Button('Apply', primary=True))
    why = Section('Why this guild', '<p class="t-prose" style="padding-top:12px">We raid two nights a week, bring flasks, and split loot by council with the numbers shown. New to Forever is fine.</p>')
    main = Section('Olympus XXVII is recruiting', card) + Section('Apply', form)
    rail = why + Section('Officers', Rows([Row('Ironbrand', 'Holy Paladin &middot; officer', lead=crest_img('paladin', 32), title_color=C['paladin'], link=False), Row('Shiverknife', 'Assassination Rogue &middot; officer', lead=crest_img('rogue', 32), title_color=C['rogue'], link=False)], open_=True))
    return build_page('Recruitment', PageHeader('Join Olympus XXVII', ALLIANCE, Emblem('alliance', 56), meta=['Public guild card']) + Page('split', main, rail), 'Guild', None)


def applications():
    apps = []
    for name, cls, spec, il, parse, msg, y, n in [('Ashenveil', 'paladin', 'Holy', 63, 'Best parse 91', 'Two tiers in a guild that disbanded; can raid Tue and Thu.', 2, 0), ('Quillbright', 'warlock', 'Destruction', 60, 'Best parse 74', 'New to Forever, level 60 since launch.', 1, 1)]:
        apps.append(Row(name, f'{spec} {cls.title()} &middot; ilvl {il} &middot; {parse}<br>{msg}', lead=crest_img(cls, 36), title_color=C[cls], trail=Votes(y, n, 3) + Button('Accept', primary=True, small=True) + Button('Decline', small=True), link=False))
    main = Section('Open applications', Plate(Rows(apps)), meta='Accepting grants the Discord role')
    rail = Section('Your public card', Rows([Row('Looking for', '1 tank, 2 healers', trail=Link('Edit'), link=False), Row('Raids', 'Tue and Thu, 8:00 pm CST', trail=Link('Edit'), link=False)], open_=True)) + Section('Recent decisions', Rows([Row('Hexmarrow accepted', '3 of 3 votes &middot; Discord role granted', link=False)], open_=True))
    return build_page('Applications', hdr(tabs_active='Roster') + Page('split', main, rail), 'Guild', ('Dunmaro', 'hunter', 'Beast Mastery Hunter &middot; 60'))


def member_profile():
    r = ch('Frostquill')
    figs = Figures([('3 of 4', 'Attendance', '', ''), ('96', 'Best parse', '', 'gold'), ('68', 'Item level', '', ''), ('1', 'Items this tier', '', '')])
    gear = Rows([Row('Gloves', 'Frozen Gauntlets &rarr; Gloves of the Greatfather', lead=icon('inv_gauntlets_18', 36, 'q3'), trail=Figure('+9.8', '', 'DPS', 'gold', 'sm'), link=False),
                 Row('Ring', 'Band of the Elder &rarr; Seal of the Archmagus', lead=icon('inv_jewelry_ring_14', 36, 'q3'), trail=Figure('+6.1', '', 'DPS', 'gold', 'sm'), link=False)], open_=True)
    att = Grid(['Mon', 'Tue', 'Thu', 'Sat'], [('Last 4 raids', '', [('ok', '&#10003;'), ('ok', '&#10003;'), ('miss', '&times;'), ('ok', '&#10003;')])])
    notes = Rows([Row('Always brings flasks.', 'Ironbrand &middot; 2 Dec', trail=Tag('Officers only', 'muted'), link=False), Row('Ask about the Thursday conflict.', 'Dunmaro &middot; 5 Dec', trail=Tag('Officers only', 'muted'), link=False)], open_=True)
    addn = FormRow([('Add a note', Field('Only officers see notes', grow=True), True)], Button('Save note', small=True))
    main = Section('Gear gap', gear, Link('Open in the planner &rsaquo;'), meta='Same units as the BiS page') + Section('Attendance', att) + Section('Loot this tier', Rows([Row('Ring of Binding', 'Awarded 8 Dec &middot; main spec', lead=icon('inv_jewelry_ring_13', 36, 'q4'), link=False)], open_=True))
    rail = Section('Officer notes', notes + addn, meta='Private: members never see these')
    head = PageHeader(r[0], C[r[1]], crest_img(r[1], 56), meta=[f'{r[2]} {r[1].title()}', 'Level 60', 'Guild member since 12 Nov'], answer=figs, actions='')
    return build_page('Member profile', head + Page('split', main, rail), 'Guild', ('Dunmaro', 'hunter', 'Beast Mastery Hunter &middot; 60'))


PAGES = {'web-guild-member': landing_member, 'web-guild-officer': landing_officer, 'web-guild-raid': raid_planner, 'web-guild-loot': loot_council, 'web-guild-report': raid_report,
         'web-guild-recruit': recruitment, 'web-guild-applications': applications, 'web-guild-profile': member_profile}

if __name__ == '__main__':
    import sys
    for n, f in PAGES.items():
        if not sys.argv[1:] or n in sys.argv[1:]:
            print(write(n + '.html', f()))
