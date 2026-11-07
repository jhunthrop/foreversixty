"""Guild persona round: role-aware landing (member and officer), raid planner, loot council, raid report, recruitment, applications and a
member profile. Composed only from ws_prims primitives. Data shapes: api/cmd/seedguild (24 characters), api/internal/guilds (sign-up states
In / Late / Bench, loot candidates, readiness, attendance, raid report). Modules whose data source is not confirmed carry a Gap marker.
"""
from ws_prims import *  # noqa: F403
from ws_prims import (Page, PageHeader, Section, Plate, Rows, Row, Figures, Figure, Meter, Tag, Button, Link, TabStrip, FormRow, Field, Select, Notice, Segmented, Grid, Tally,
                      Checklist, Votes, Gap, build_page, crest_img, icon, write, Emblem, RowGroup, Prose, control_bar, seg, BackLink, SignupRow, ButtonRow,
                      Toggle, Entries, Inspector, Live, Split, RowsSlim, Disclosure, ItemRow)

import guild_report_data as rd

from guild_data import ALLIANCE, GUILD, NOW, TABS, C, R, STATE_TAG  # noqa: F401
import guild_raid


def ch(name):
    return next(r for r in R if r[0] == name)


def lead(name, size=32):
    r = ch(name)
    return crest_img(r[1], size)


def hdr(title_suffix='', tabs_active='Overview', actions='', notice=''):
    return PageHeader(GUILD, ALLIANCE, Emblem('alliance', 56), meta=['<b class="strong">US</b> &middot; PvP', '24 raiders', 'Next raid: Onyxia&rsquo;s Lair, Thu 10 Dec'], actions=actions, tabs=TabStrip(TABS, tabs_active)) + notice


def next_raid(officer):
    tally = Figures([('19', 'In', '', 'kill'), ('2', 'Late', '', 'gold'), ('2', 'Bench', '', ''), ('3', 'Not signed', '', 'wipe')]) if officer else ''
    you = ('<div style="padding:16px"></div>')
    top = Row("Onyxia's Lair", 'Thu 10 Dec &middot; 8:00 pm CST &middot; 9:00 pm EST your time', lead=icon('inv_misc_head_dragon_01', 40, 'q5'), trail=Figure('2d 4h', 'Starts in', size='sm'), link=False)
    status = Row('Hollowmend' if not officer else 'Dunmaro', ('Holy Priest &middot; level 60 &middot; you are <b class="strong">In</b> as a healer' if not officer else 'Beast Mastery Hunter &middot; level 60 &middot; you are <b class="strong">In</b> as ranged'),
                 lead=crest_img('priest' if not officer else 'hunter', 36), trail=Segmented([('In', True), ('Late', False), ('Bench', False)], small=True), link=False)
    return Plate(Rows([top, status]) + (f'<div style="padding:12px 16px">{tally}</div>' if officer else ''))


def landing_member():
    notice = Notice('Almost there: two things before Thursday.', Link('See your checklist below &darr;'))
    check = Checklist([('ok', 'Level 60', 'You are level 60.', None), ('ok', 'Attuned to Onyxia&rsquo;s Lair', 'Drakefire Amulet in your bags. Read from your addon export.', None),
                       ('warn', '2 slots are not best in slot', 'Your gloves and ring. The planner shows the swap.', 'Fix in the BiS list &rsaquo;'), ('ok', 'Enchants', 'Every slot that takes one is enchanted.', None),
                       ('miss', 'Flask or food', 'None in your bags. Read from your addon export, before you go.', 'What to bring &rsaquo;'), ('gap', 'Fire resistance', 'Not tracked yet: needs the resistance total in the addon export.', None)])
    nums = Figures([('3 of 4', 'Raids attended', '', ''), ('1', 'Items this tier', '', ''), ('62', 'Your item level', '', '')])
    asked = Rows([Row('Update your addon export', 'Officers need your gear to plan loot. Last read 9 days ago.', trail=Button('How', small=True), link=False)], open_=True)
    cls = Tally([(crest_img('priest', 24), 'Priests', 2, 3), (crest_img('paladin', 24), 'Healers in the raid', 6, 8)])
    main = (Section('Next raid', next_raid(False)) + Section('Before the raid', check, meta='Private to you &middot; Almost there')
            + Section('Your numbers', nums) + Section('Officers asked you to', asked))
    rail = (Section('Your class in the guild', cls, meta='Two priests; the guild wants three')
            + Section('Last raid', Rows([Row('Onyxia', '1 kill &middot; 1 wipe &middot; you were there', trail=Link('Report &rsaquo;'), link=False)], open_=True))
            + Section('Progression', '<p class="t-small" style="padding-top:12px">1 of 1 named encounters down.</p>'))
    return build_page('Guild, member', hdr(notice='') + notice + Page('split', main, rail, sticky_rail=True), 'Guild', ('Hollowmend', 'priest', 'Holy Priest &middot; 60'))


def landing_officer():
    action = Rows([Row('3 raiders are waiting for approval', 'Nightlatch, Voidcall, Skyarcana', trail=Button('Review', small=True), link=False),
                   Row('2 applications', 'Both level 60, both with a parse', trail=Button('Open', small=True), link=False),
                   Row('3 have not signed up for Thursday', 'Voidcall, Skyarcana, Starnettle', trail=Button('Ping on Discord', small=True), link=False),
                   Row('2 drops to award', 'Onyxia, Tue 8 Dec', trail=Button('Open loot', small=True), link=False)], open_=True)
    comp = Tally([(crest_img('warrior', 22), 'Tanks', 1, 2), (crest_img('paladin', 22), 'Healers', 6, 8), (crest_img('rogue', 22), 'Melee', 5, 10), (crest_img('mage', 22), 'Ranged', 11, 20)])
    cover = Grid(['Need', 'In'], [('Windfury Totem', '2 groups', [('ok', '2'), ('ok', '3')]), ('Blessing of Might', '', [('ok', '1'), ('ok', '3')]), ('Blessing of Kings', '', [('ok', '1'), ('warn', '1')]),
                                  ('Blessing of Salvation', '', [('ok', '1'), ('ok', '2')]), ('Prayer of Fortitude', '', [('ok', '1'), ('ok', '4')]), ('Arcane Intellect', '', [('ok', '1'), ('ok', '3')]),
                                  ('Mark of the Wild', '', [('ok', '1'), ('ok', '1')]), ('Curse of the Elements', '', [('ok', '1'), ('miss', '0')]), ('Sunder Armor', '', [('ok', '1'), ('ok', '1')])])
    ready = Figures([('18 of 24', 'Attuned', '', ''), ('&mdash;', 'Fire resist on tanks', '', '')]) + f'<div style="padding-top:8px">{Gap("Needs resistance totals and consumables in the addon export")}</div>'
    att = Grid(['Mon', 'Tue', 'Thu', 'Sat'], [('Pinefeather', 'slipping', [('ok', '&#10003;'), ('ok', '&#10003;'), ('miss', '&times;'), ('miss', '&times;')]), ('Skyarcana', 'slipping', [('ok', '&#10003;'), ('miss', '&times;'), ('miss', '&times;'), ('none', '&mdash;')]),
                                              ('Cinderglass', 'bench', [('warn', 'B'), ('warn', 'B'), ('ok', '&#10003;'), ('warn', 'B')])])
    main = (Section('Action needed', action) + Section('Next raid', next_raid(True), Link('Open the planner &rsaquo;'))
            + Section('Comp against the target', comp, Link('Edit target &rsaquo;'), meta='Onyxia default: 40 players, editable'))
    rail = (Section('Buff and debuff cover', cover, meta='Signed-up raiders') + Section('Raid-ready', ready) + Section('Attendance, last 4 raids', att, Link('Roster &rsaquo;'))
            + Section('Last raid', Rows([Row('Onyxia', '1 kill &middot; 1 wipe &middot; 22 of 24 there', trail=Link('Report &rsaquo;'), link=False)], open_=True), meta='1 of 1 named encounters down'))
    return build_page('Guild, officer', hdr(actions=Button('Invite link', small=True)) + Page('split', main, rail, sticky_rail=True), 'Guild', ('Dunmaro', 'hunter', 'Beast Mastery Hunter &middot; 60'))


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


def cover_rows(raid, kind):
    per = {'raid': '', 'melee_group': 'per melee group', 'caster_group': 'per caster group'}
    tone = {'ok': 'ok', 'short': 'warn', 'missing': 'miss'}
    out = []
    for r in (x for x in raid['cover'] if x['kind'] == kind):
        sub = ' &middot; '.join(x for x in [per[r['per']], f"{r['likely']} on a talent we cannot see" if r['likely'] else '',
                                              f"{r['providers_bench']} on the bench" if r['providers_bench'] else ''] if x)
        out.append((r['name'], sub, [('none', str(r['need'])), (tone[r['state']], str(r['providers_in'] + r['providers_late']))], icon(r['icon'], 24, '').replace('<img ', f'<img title="{r["name"]}" ', 1)))
    return out


def assignment_rows(encounter):
    rows = []
    for s in encounter['slots']:
        names = ''.join(Tag(n['name'], 'gold' if n['own'] else 'muted') for n in s['names']) or '<span class="t-small">Open</span>'
        own = any(n['own'] for n in s['names'])
        rows.append(Row(s['label'], f"{len(s['names'])} of {s['count']} &middot; {s['purpose']}", trail=Button('Edit', small=True, tone='ghost', act=True), foot=names,
                        link=False, selected=own))
    return rows


def raid_planner():
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
    main = (Section('Your sign-up', own, named=True) + Section('Sign-ups', Plate(roster_groups(raid), 'Sign-ups'), meta=f"{signed}/{raid['size']}")
            + Section('Not signed up', Rows(not_signed, open_=True, list_=True), action=Button('Ping on Discord', small=True, act=True), meta=str(len(not_signed))))
    t = raid['comp_target']
    late = n['late_groups']
    comp = Tally([('', label, have, t[key], f'{late[key]} late' if late[key] else '') for key, label, have in
                  [('tanks', 'Tanks', n['tanks']), ('healers', 'Healers', n['healers']), ('melee', 'Melee', n['melee']), ('ranged', 'Ranged', n['ranged'])]])
    bench = [Row(m['name'], planner_sub(m), lead=crest_img(m['class'], 24), title_color=C[m['class']], trail=Button('Swap in', small=True, act=True), link=False, slim=True)
             for m in sorted((by[s['name']] for s in raid['signups'] if s['state'] == 'bench'), key=lambda m: (CLASS_ORDER.index(m['class']), m['name']))]
    enc = raid['assignments']['encounters'][0]
    yours = ', '.join(s['label'] for s in enc['slots'] if any(x['own'] for x in s['names']))
    assign = (Rows(assignment_rows(enc), open_=True, list_=True)
              + ButtonRow(Button('Copy as raid warning', small=True, act=True) + Button('Post to Discord', small=True, act=True)))
    cover = (Grid(['Need', 'In'], cover_rows(raid, 'buff'), 'Buffs: need against signed-up raiders', 'Buffs')
             + Grid(['Need', 'In'], cover_rows(raid, 'debuff'), 'Debuffs: need against signed-up raiders', 'Debuffs'))
    rail = (Section('Comp against the target', comp + f'<div style="padding-top:16px">{Button("Change the target", small=True, act=True)}</div>', meta=f"{raid['zone_name']} default: {sum(t[k] for k in ('tanks', 'healers', 'melee', 'ranged'))} players")
            + Section('Bench and swaps', Rows(bench, open_=True, list_=True), meta=str(len(bench)))
            + Section(f"{enc['name']} assignments", assign, meta=f'Your job: {yours}', named=True)
            + Section('Buff and debuff cover', cover, meta='Signed-up raiders')
            + Section('Notes', Prose(raid['notes_html'], 'note'), named=True))
    actions = ''.join(Button(x, small=True, act=True) for x in ['Edit', 'Lock sign-ups', 'Copy the sign-up list', 'Cancel raid'])
    head = (BackLink(f'&lsaquo; {GUILD} &middot; Raids')
            + PageHeader(raid['title'], ALLIANCE, meta=[Tag('Planned'), raid_when(raid), f"<b class=\"strong\">{signed} of {raid['size']} signed up</b> &middot; {n['in']} In &middot; {n['late']} Late &middot; {n['bench']} Bench"],
                         actions=actions, tabs=TabStrip(TABS, 'Raids')))
    # The guild route's <main> ends 40 px under its view (pages/guild/[...path].astro pb-10).
    body = f'<div style="padding-bottom:40px">{head + bar + Page("split", main, rail, sticky_rail=True)}</div>'
    return build_page('Raid planner', body, 'Guild', (guild_raid.VIEWER, you['class'], f"{you['spec']} {you['class'].title()} &middot; 60"))


def loot_council():
    drops = []
    for nm, ic, q, sl, n in [('Vis&rsquo;kag the Bloodletter', 'inv_sword_18', 'q4', 'One-hand sword', 5), ('Deathbringer', 'inv_axe_09', 'q4', 'One-hand axe', 4), ('Eskhandar&rsquo;s Collar', 'inv_belt_12', 'q4', 'Neck', 6),
                             ('Ring of Binding', 'inv_jewelry_ring_13', 'q4', 'Finger', 7), ('Sapphiron Drape', 'inv_misc_cape_16', 'q4', 'Back', 5), ('Head of Onyxia', 'inv_misc_head_dragon_01', 'q4', 'Quest item', 2)]:
        drops.append(Row(nm, f'{sl} &middot; {n} candidates', lead=icon(ic, 40, q), trail=Tag('Open', 'gold') if n != 2 else Tag('Awarded', 'kill'), selected=(nm.startswith('Vis')), title_color='#b866f5'))
    plate = Plate(Rows(drops))
    cands = []
    for name, gain, worn, bis, att, got, ms in [('Seraphine', '+14.2', 'Mirah&rsquo;s Song', 'Yes', '4 of 4', 0, 'MS'), ('Shiverknife', '+11.6', 'Dal&rsquo;Rend&rsquo;s Sacred Charge', 'No', '3 of 4', 1, 'MS'), ('Groundshock', '+9.1', 'Windfury Blade', 'No', '4 of 4', 0, 'OS'), ('Blacktide', '+4.0', 'Skullforge Reaver', 'No', '2 of 4', 2, 'OS')]:
        r = ch(name)
        cands.append(Row(name, f'<b class="gold">{gain} DPS</b> vs your {worn} &middot; BiS for them: {bis} &middot; attended {att} &middot; {got} this tier', lead=crest_img(r[1], 32), title_color=C[r[1]],
                         trail=Tag(ms, 'gold' if ms == 'MS' else 'muted') + Votes(3 if name == 'Seraphine' else 1, 0, 4) + Button('Award', small=True), link=False))
    cand = Rows(cands, open_=True)
    record = Rows([Row('Eskhandar&rsquo;s Collar to Hollowmend', 'Awarded by Ironbrand &middot; main spec &middot; 3 votes of 4', lead=icon('inv_belt_12', 36, 'q4'), trail=Link('Edit'), link=False),
                   Row('Ring of Binding to Doomwhisper', 'Awarded by Dunmaro &middot; main spec &middot; 4 votes of 4', lead=icon('inv_jewelry_ring_13', 36, 'q4'), trail=Link('Edit'), link=False)], open_=True)
    main = Section('Onyxia drops', plate, meta='Grouped by raid night: Tue 8 Dec')
    rail = Section('Candidates for Vis&rsquo;kag the Bloodletter', cand, meta='Gain in DPS, as on the BiS page') + Section('Decision record', record, Button('Export history', small=True))
    head = PageHeader('Loot council', ALLIANCE, Emblem('alliance', 56), meta=['Onyxia', '6 drops, 4 open', 'Next unkilled: none, farm'], tabs=TabStrip(TABS, 'Loot'))
    bar = control_bar([('Raid night', seg([('Tue 8 Dec', True, None, ''), ('Mon 15 Dec', False, None, '')]))], 'Tue 8 Dec')
    return build_page('Loot council', head + bar + Page('split', main, rail, sticky_rail=True), 'Guild', ('Dunmaro', 'hunter', 'Beast Mastery Hunter &middot; 60'))


def body(html):
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
                       f'<b>{d["killing_blow"]["ability"]}</b> from {d["killing_blow"]["source"]}, 4,210 (1,050 overkill)') for d in shown['death_causes']])
    insp = Inspector(rd.pull_title(shown), [('Deaths', deaths)], rd.pull_title(shown), rd.pull_sub(shown))
    pulls = Plate(Split(Rows(rows), insp), 'Every pull of the night')
    output = Rows([Row(p['name'], f'DPS &middot; rated {p["rating"]} of 100', lead=crest_img(p['class'], 24), title_color=C[p['class']],
                       trail=Figure(str(p['value']), '', 'DPS', size='sm'), link=False, selected=p['name'] == rd.VIEWER, slim=True) for p in rep['parses']], open_=True)
    unknown = ''.join(f'<p class="t-small" style="margin-top:4px">{rd.pull_title(b)}: the log carried no character sheet, so nothing can be said.</p>' for b in rep['consumables']['by_pull'])
    match = raid['match']
    logged = (f'<p class="t-small">Logged as a raid night. {Link("Open the night")}</p><div><p class="t-small">Log uploaded '
              f'{rd.local(match["uploaded_at"]):%a} {rd.local(match["uploaded_at"]).day} Dec, {rd.hour(rd.local(match["uploaded_at"]))} by {match["uploaded_by_name"]}'
              f' &middot; matched by time</p>{Button("Not this log", small=True, act=True)}</div>')
    fix = RowsSlim([Row(a['name'], 'Present' if a['outcome'] == 'present' else 'Out', lead=crest_img(a['class'], 24), title_color=C[a['class']],
                        trail=Select('Fix'), link=False, selected=a['own']) for a in raid['attendance']])
    att = logged + f'<p>{rd.attendance_summary(raid["night"])}</p>' + Disclosure('Fix attendance', fix)
    lootr = ''.join(ItemRow(a['item_name'], a['quality'], icon(a['icon'], 32, f'q{a["quality"]}'), f'{g["boss"]} &middot; to {a["name"]}')
                    for g in rep['loot'] for a in g['awards'])
    lines = rd.discord_lines(raid)
    discord = (''.join(f'<p class="t-small">{"<b>" + x + "</b>" if i == 0 else x}</p>' for i, x in enumerate(lines))
               + f'<div style="display:flex;gap:12px;padding-top:12px">{Button("Post to Discord", primary=True, small=True, act=True)}{Button("Copy for Discord", small=True, act=True)}{Live()}</div>')
    main = Section('Pulls', pulls, meta=f'{len(rep["pulls"])} fights') + Section('DPS and HPS', body(output), meta='Boss kills') + Section('Consumables by pull', body(unknown))
    rail = Section('Attendance', att) + Section('Loot awarded', body(lootr)) + Section('Discord summary', body(discord))
    h = rep['headline']
    fig = Figures([(str(h['kills']), 'Kills', '', 'kill'), (str(h['wipes']), 'Wipes', '', 'wipe'), (str(h['attempts']), 'Attempts', '', ''), (str(h['deaths']), 'Deaths', '', '')])
    head = PageHeader(raid['title'], ALLIANCE, meta=[Tag('Done', 'muted'), rd.when_line(raid)], answer=fig)
    return build_page('Raid report', head + Page('split', main, rail, sticky_rail=True), 'Guild', (rd.VIEWER, viewer[1], f'{viewer[2]} {viewer[1].title()} &middot; 60'))


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
