"""Guild persona round: role-aware landing (member and officer), raid planner, loot council, raid report, recruitment, applications and a
member profile. Composed only from ws_prims primitives. Data shapes: api/cmd/seedguild (24 characters), api/internal/guilds (sign-up states
In / Late / Bench, loot candidates, readiness, attendance, raid report). Modules whose data source is not confirmed carry a Gap marker.
"""
from ws_prims import *  # noqa: F403
from ws_prims import (Page, PageHeader, Section, Plate, Rows, Row, Figures, Figure, Meter, Tag, Button, Link, TabStrip, FormRow, Field, Select, Notice, Segmented, Grid, Tally,
                      Checklist, Votes, Gap, build_page, crest_img, icon, write, Emblem)

from guild_data import ALLIANCE, GUILD, TABS, C, R, STATE_TAG  # noqa: F401


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
