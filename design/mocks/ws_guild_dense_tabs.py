"""The guild tabs batch H redraws in the dense language (design/specs/2026-11-09-guild-dense.md, section 5.6 to 5.11):
Raids, the raid planner, Progression, Applications, Recruit and Settings, drawn from the 40-raider world
(guild_dense_data.py) with the shared primitives only.

  python3 design/mocks/ws_guild_dense_tabs.py        -> design/mocks/web/web-guild-dense-{raids,planner,...}.html
"""
from __future__ import annotations



from ws_prims import (Button, Figures, Link, Page, PageHeader, Row, Rows, Section, Segmented, StatRow, StatTable, Tag, TabStrip,
                      build_page, crest_img, icon, write)
import ws_prims
from guild_data import C
import guild_data as gd
import guild_dense_data as D
from ws_guild2 import _view, _ROLE_CSS
from ws_guild_dense import (GROUP_WORD, VIEW_MEMBER, VIEW_OFFICER, countdown, day, dense, fig, gap_of, head, legend, name_html,
                            officer_answer, parse_class, ready_counts, strip, tone)
from guild_dense_tabs_fixtures import APPLICATIONS, WANTED

ONY_ICON = 'inv_misc_head_dragon_black'
VOTERS = gd.VOTERS


def clock(ms: int) -> str:
    return f'{ms // 60000}:{ms // 1000 % 60:02d}'


def grp(label: str, sub: str, n: int) -> str:
    """A group head row inside a dense table (the roster's role heads)."""
    return StatRow('', f'<b>{label}</b>', sub, [''] * n, bar=False).replace('st-r stc', 'st-r stc grp', 1)


# ---- Raids
def planned_table() -> str:
    ready, coming = ready_counts()
    short = sum(max(0, gap_of(g)) for g in D.GROUPS)
    n_in, n_late, n_bench = (len(D.answered(s)) for s in ('in', 'late', 'bench'))
    rows = [StatRow(icon(ONY_ICON, 28, 'q5'), "<b>Onyxia's Lair</b>", f'{day(D.NEXT_RAID_AT)} &middot; 8 pm &middot; '
                    f'<span class="gold">{countdown()}</span> &middot; posted to #raid-signups',
                    [str(n_in), str(n_late), str(n_bench), tone(str(len(D.answered(None))), 'wipe'), tone(f'{short} short', 'wipe'),
                     f'{ready}/{coming}'], bar=False, primary_phone=f'{n_late} late &middot; {short} short'),
            StatRow(icon(ONY_ICON, 28, 'q5'), "<b>Onyxia's Lair</b>", f'{day(D.SECOND_RAID_AT)} &middot; 8 pm &middot; in 11 d 4 h &middot; posts Thu',
                    ['11', '0', '0', tone('29', 'mute'), tone('&ndash;', 'mute'), tone('&ndash;', 'mute')], bar=False)]
    return dense(StatTable([('', False), ('Raid', False), ('In', True), ('Late', True), ('Bench', True), ('No answer', True),
                            ('Comp', True), ('Ready', True)], rows, 6, bar=False))


def nights_table() -> str:
    worst = max(n[3] for n in D.NIGHTS)
    rows = []
    prev = None
    for n in D.NIGHTS:
        rid, at_, rep, kill, wipes, deaths = n
        came = [r for r in D.ROSTER if r.cells[D.NIGHTS.index(n)] == 'P']
        top = max((r for r in came if r.group not in ('Tank', 'Healer')), key=lambda r: r.figures[1])
        heal = max((r for r in came if r.group == 'Healer'), key=lambda r: r.figures[1])
        loot = sum(1 for a in D.AWARDS if a[0] == rid)
        delta = '&ndash;' if prev is None else tone(f'{(kill - prev) / 1000:+.0f} s', 'kill' if kill < prev else 'wipe')
        best = kill == min(x[3] for x in D.NIGHTS)
        rows.append(StatRow(icon(ONY_ICON, 28, 'q5'), f'<b>{day(at_)}</b>', f"Onyxia's Lair &middot; {3 + wipes} pulls" + (' &middot; <span class="gold">best kill</span>' if best else ''),
                            [clock(kill), delta, str(wipes), str(deaths), f'{len(came)}', f'{name_html(top)} <span class="mute">{top.figures[1]:.0f}</span>',
                             f'{name_html(heal)} <span class="mute">{heal.figures[1]:.0f}</span>', str(loot)],
                            meter=ws_prims.Meter(kill / worst), primary_phone=f'{wipes} wipes &middot; {deaths} deaths'))
        prev = kill
    rows.reverse()
    return dense(StatTable([('', False), ('Night', False), ('Onyxia kill', False), ('Kill', True), ('vs last', True), ('Wipes', True),
                            ('Deaths', True), ('Raiders', True), ('Top DPS', True), ('Top HPS', True), ('Loot', True)], rows, 8))


def raids() -> str:
    kills = [n[3] for n in D.NIGHTS]
    meta = f'{len(D.NIGHTS)} nights &middot; {len(D.NIGHTS)} kills &middot; best {clock(min(kills))} &middot; {sum(n[4] for n in D.NIGHTS)} wipe'
    main = (Section('Planned', planned_table(), Button('Plan an event', True, True, act=True), meta='2', panel=True, flush=True)
            + Section('Raid nights', nights_table(), Link('Loot &rsaquo;'), meta=meta, panel=True, flush=True))
    return build_page('Guild raids, dense', _view(head('Raids', True, officer_answer()) + Page('ledger', main)) + _ROLE_CSS, 'Guild', VIEW_OFFICER)


# ---- the raid planner (raid 45)
STATE_WORD = {'in': 'In', 'late': 'Late', 'bench': 'Bench', 'out': 'Out'}


def signup_row(x: D.Raider, officer: bool) -> str:
    short = D.short_for_raid(x) if x.answer in ('in', 'late') else []
    ready = tone('&#10003;', 'kill') if not short else tone(f'{len(short)} short', 'wipe')
    seg = Segmented([(w, x.answer == k) for k, w in STATE_WORD.items()], small=True) if officer else tone(STATE_WORD.get(x.answer, 'No answer'), 'mute')
    pct = round(100 * x.present / x.nights)
    return StatRow(crest_img(x.cls, 28), name_html(x), f'{x.spec} {x.cls.title()}' + (' &middot; late, work till 8:30' if x.name == 'Pinefeather' else ''),
                   [f'{x.ilvl} <span class="mute">{x.ilvl - D.MEDIAN_ILVL:+d}</span>', ready, str(x.ratings[1]), f'{pct}%', seg],
                   meter=strip(x.cells), primary_phone=STATE_WORD.get(x.answer, ''))


def signups_table(officer: bool) -> str:
    rows = []
    for g in D.GROUPS:
        members = sorted((x for x in D.ROSTER if x.group == g and x.answer in ('in', 'late')), key=lambda x: (x.answer != 'in', -x.ratings[1]))
        need = D.COMP_TARGET[D.GROUP_TARGET[g]]
        gap = gap_of(g)
        rows.append(grp(f'{GROUP_WORD[g]} {len(members)} of {need}', tone(f'{gap} short', 'wipe') if gap > 0 else tone('full', 'kill'), 5))
        rows += [signup_row(x, officer) for x in members]
    rest = [x for x in D.ROSTER if x.answer in ('bench', 'out')]
    rows.append(grp(f'Bench and Out {len(rest)}', '', 5))
    rows += [signup_row(x, officer) for x in rest]
    return dense(StatTable([('', False), ('Raider', False), ('Last 5', False), ('Item level', True), ('Ready', True), ('Rating', True),
                            ('30 nights', True), ('', True)], rows, 5))


def not_answered() -> str:
    rows = []
    for x in D.answered(None):
        why = f'{x.present} of {x.nights} nights' + (f' &middot; last log {x.seen_days} d' if x.seen_days >= 2 else '')
        rows.append(StatRow(crest_img(x.cls, 28), name_html(x), f'{x.spec} {x.cls.title()} &middot; {why}',
                            [GROUP_WORD[x.group], Button('Move to In', False, True, tone='ghost', act=True)], meter=strip(x.cells), primary_phone=GROUP_WORD[x.group]))
    return dense(StatTable([('', False), ('Raider', False), ('Last 5', False), ('Role', True), ('', True)], rows, 2))


def comp_rail() -> str:
    rows = []
    for g in D.GROUPS:
        c = D.group_counts(g)
        need = D.COMP_TARGET[D.GROUP_TARGET[g]]
        gap = gap_of(g)
        fill = [x for x in D.ROSTER if x.group == g and x.answer in (None, 'bench')]
        sub = ', '.join(f'<span style="color:{C[x.cls]}">{x.name}</span> {"bench" if x.answer else "no answer"}' for x in fill[:2])
        rows.append(StatRow('', f'<b>{GROUP_WORD[g]}</b>', sub, [str(need), str(c['in']), str(c['late']),
                            tone(str(gap), 'wipe') if gap > 0 else tone('&#10003;', 'kill')], bar=False))
    return dense(StatTable([('', False), ('Role', False), ('Need', True), ('In', True), ('Late', True), ('Short', True)], rows, 4, bar=False))


def planner(officer: bool = True) -> str:
    n_in, n_late, n_bench = (len(D.answered(s)) for s in ('in', 'late', 'bench'))
    ready, coming = ready_counts()
    meta = [Tag('Planned', 'gold'), f'Wednesday 13 January &middot; 8:00 pm to 11:00 pm CST &middot; <span class="gold">{countdown()}</span>',
            f'<b>{n_in + n_late} of 40</b> &middot; {n_in} In &middot; {n_late} Late &middot; {n_bench} Bench &middot; {len(D.answered(None))} no answer &middot; {ready} of {coming} ready']
    actions = (Button('Edit', False, True, act=True) + ' ' + Button('Lock sign-ups', False, True, act=True) + ' '
               + Button('Copy the list', False, True, act=True)) if officer else ''
    hdr = PageHeader("Onyxia's Lair", '#3fc7eb', meta=meta, actions=actions, crumb='&lsaquo; Olympus XXVII &middot; Raids',
                     tabs=TabStrip(gd.TABS if officer else [t for t in gd.TABS if t != 'Settings'], 'Raids'))
    you = D.BY_NAME[D.OFFICER if officer else D.MEMBER]
    yours = (f'<div style="display:flex;gap:12px;align-items:center;padding:12px 20px">{crest_img(you.cls, 28)}{name_html(you)}'
             f'<span class="t-small">{you.spec} {you.cls.title()}</span><span style="margin-left:auto">'
             + Segmented([(w, you.answer == k) for k, w in STATE_WORD.items()], small=True) + '</span></div>')
    main = (Section('Your sign-up', yours, panel=True, flush=True)
            + Section('Sign-ups', signups_table(officer), Link('Readiness &rsaquo;'), meta=f'{n_in + n_late}/40 &middot; {legend()}', panel=True, frame='framed', flush=True))
    if officer:
        main += Section('No answer', not_answered(), Button(f'Ping {len(D.answered(None))} on Discord', True, True, act=True),
                        meta='Posted to #raid-signups Thu 7 Jan, 12 pm', panel=True, flush=True)
    rail = (Section('Comp against the target', comp_rail(), Link('Change') if officer else '', panel=True, flush=True)
            + Section('Notes', '<p class="t-small" style="padding:0 20px 16px">Fire resist on tanks. Flask and food at the door, 7:50 pm invites.</p>', panel=True))
    page = Page('split', main, rail, panels=True)
    who = VIEW_OFFICER if officer else VIEW_MEMBER
    return build_page('Raid planner, dense', _view(hdr + page) + _ROLE_CSS, 'Guild', who)


# ---- Progression
def nights_by_boss() -> str:
    worst = max(n[3] for n in D.NIGHTS)
    rows = []
    for i, n in enumerate(D.NIGHTS):
        rid, at_, rep, kill, wipes, deaths = n
        came = sum(1 for r in D.ROSTER if r.cells[i] == 'P')
        prev = D.NIGHTS[i - 1][3] if i else None
        delta = '&ndash;' if prev is None else tone(f'{(kill - prev) / 1000:+.0f} s', 'kill' if kill < prev else 'wipe')
        rows.append(StatRow('', f'<b>{day(at_)}</b>', f'{came} raiders' + (' &middot; first kill' if i == 0 else ''),
                            [clock(kill), delta, str(1 + wipes), str(wipes), str(deaths)], meter=ws_prims.Meter(kill / worst)))
    rows.reverse()
    return dense(StatTable([('', False), ('Night', False), ('Kill time', False), ('Kill', True), ('vs last', True), ('Pulls', True),
                            ('Wipes', True), ('Deaths', True)], rows, 5))


def best_by_role() -> str:
    picks = []
    for group, label in (('Tank', 'Tank'), ('Healer', 'Healer'), ('Melee', 'Damage')):
        pool = [r for r in D.ROSTER if r.group == group] if group != 'Melee' else [r for r in D.ROSTER if r.group in ('Melee', 'Ranged')]
        best = max(pool, key=lambda r: max(r.figures))
        picks.append(StatRow(crest_img(best.cls, 28), name_html(best), f'{label} &middot; {best.spec} {best.cls.title()}',
                             [f'{max(best.figures):.0f} <span class="mute">{best.metric.upper()}</span>', tone(str(best.best_pct), parse_class(best.best_pct))], bar=False))
    return dense(StatTable([('', False), ('Best on Onyxia', False), ('Figure', True), ('%', True)], picks, 2, bar=False))


def roster_bests() -> str:
    top = sorted(D.ROSTER, key=lambda r: -r.best_pct)[:8]
    rows = [StatRow(crest_img(r.cls, 28), name_html(r), f'{r.spec} {r.cls.title()}',
                    [tone(str(r.best_pct), parse_class(r.best_pct)), f'{max(r.figures):.0f}'], bar=False) for r in top]
    return dense(StatTable([('', False), ('Raider', False), ('Best %', True), ('Figure', True)], rows, 2, bar=False))


def progression() -> str:
    kills = [n[3] for n in D.NIGHTS]
    pulls = sum(1 + n[4] for n in D.NIGHTS)
    deaths = sum(n[5] for n in D.NIGHTS)
    tier = Figures([fig('1/1', 'Bosses down'), fig(str(len(D.NIGHTS)), 'Kills'), fig(clock(min(kills)), 'Best kill', day(D.NIGHTS[-1][1])),
                    fig(f'{(kills[0] - kills[-1]) / 1000:.0f} s', 'Faster since first', t='kill'), fig(f'{deaths / pulls:.1f}', 'Deaths a pull'),
                    fig(day(D.NIGHTS[0][1]), 'First kill')])
    unnamed = Rows([Row('Barrow Deeps', 'no named encounters published &middot; 0 pulls', link=False, slim=True),
                    Row('Hyjal Summit', 'no named encounters published &middot; 0 pulls', link=False, slim=True)], open_=True)
    main = (Section('Onyxia', f'<div style="padding:0 20px 12px">{tier}</div>' + nights_by_boss(), Link('Last report &rsaquo;'),
                    meta=f"Onyxia's Lair &middot; {pulls} pulls &middot; {len(D.NIGHTS)} nights", panel=True, frame='framed', flush=True)
            + Section('Best by role', best_by_role(), panel=True, flush=True))
    rail = (Section('First tier', unnamed, meta='1 of 1 named boss down', panel=True)
            + Section('Roster bests', roster_bests(), meta='Onyxia percentile', panel=True, flush=True))
    answer = f'<b class="gold" style="font-size:15px">1 of 1 down</b><br><span class="t-small">best {clock(min(kills))}</span>'
    return build_page('Guild progression, dense', _view(head('Progression', True, answer) + Page('split', main, rail, panels=True)) + _ROLE_CSS, 'Guild', VIEW_OFFICER)


# ---- Applications
GROUP_OF_SPEC = {'paladin-holy': 'Healer', 'warlock-destruction': 'Ranged', 'warrior-protection': 'Tank', 'druid-restoration': 'Healer',
                 'hunter-marksmanship': 'Ranged', 'rogue-combat': 'Melee'}


def applicant_row(a: tuple) -> str:
    (aid, name, cls, spec, spec_name, ilvl, metric, value, pct, days, avail, message, yes, no, mine, status, decided) = a
    g = GROUP_OF_SPEC[spec]
    gap = gap_of(g)
    fills = tone(f'{GROUP_WORD[g]}, {gap} short', 'gold') if gap > 0 else tone(f'{GROUP_WORD[g]}, full', 'mute')
    actions = Button('Accept', True, True, act=True) + ' ' + Button('Decline', False, True, tone='ghost', act=True)
    sub = spec_name + (f' &middot; {message}' if message else '') + (f' &middot; {avail}' if avail else '')
    return StatRow(crest_img(cls, 28), f'<b style="color:{C[cls]}">{name}</b>', sub,
                   [f'{ilvl} <span class="mute">{ilvl - D.MEDIAN_ILVL:+d}</span>', tone(str(pct), parse_class(pct)) + f' <span class="mute">{value:.0f}</span>',
                    fills, f'{yes}&ndash;{no} <span class="mute">of {VOTERS}</span>', f'{days} d', actions], bar=False, primary_phone=f'{pct} &middot; {days} d')


def applications() -> str:
    open_ = [a for a in APPLICATIONS if a[15] == 'pending']
    done = [a for a in APPLICATIONS if a[15] != 'pending']
    table = dense(StatTable([('', False), ('Applicant', False), ('Item level', True), ('Best %', True), ('Fills', True), ('Votes', True),
                             ('Waiting', True), ('', True)], [applicant_row(a) for a in open_], 6, bar=False))
    rec = Rows([Row(f'{a[1]} {"accepted" if a[15] == "accepted" else "declined"}', f'{a[4]} &middot; {a[12]} yes, {a[13]} no &middot; by Dunmaro &middot; {a[9] - a[16]} d ago',
                    lead=crest_img(a[2], 28), title_color=C[a[2]], link=False) for a in done], open_=True)
    gaps = Rows([Row(GROUP_WORD[g], f'{gap_of(g)} short for Wednesday &middot; {sum(1 for a in open_ if GROUP_OF_SPEC[a[3]] == g)} applying',
                     trail=tone(str(gap_of(g)), 'wipe'), link=False, slim=True) for g in D.GROUPS if gap_of(g) > 0], open_=True)
    main = Section('Open applications', table, meta=f'{len(open_)} &middot; item level vs guild median {D.MEDIAN_ILVL}', panel=True, frame='framed', flush=True)
    rail = (Section('Short for Wednesday', gaps, panel=True)
            + Section('Your public card', Rows([Row('Looking for', ', '.join(w[2] for w in WANTED), trail=Link('Edit'), link=False),
                                                Row('Raid nights', 'Wed, 8:00 pm to 11:00 pm CST', trail=Link('Edit'), link=False)], open_=True),
                      Link('See it as a visitor &rsaquo;'), panel=True)
            + Section('Decided', rec, panel=True))
    hdr = head('Roster', True, officer_answer()).replace('</nav>', '</nav>' + TabStrip(['Members', 'Applications'], 'Applications'), 1)
    return build_page('Guild applications, dense', _view(hdr + Page('split', main, rail, panels=True)) + _ROLE_CSS, 'Guild', VIEW_OFFICER)


# ---- Recruit (the public card, read by its officer)
def recruit() -> str:
    open_ = [a for a in APPLICATIONS if a[15] == 'pending']
    rows = []
    for cls, spec, spec_name, note in WANTED:
        g = GROUP_OF_SPEC.get(spec, 'Ranged')
        rows.append(StatRow(crest_img(cls, 28), f'<b style="color:{C[cls]}">{spec_name}</b>', note or GROUP_WORD[g],
                            [tone(f'{max(0, gap_of(g))} short', 'gold') if gap_of(g) > 0 else tone('bench spot', 'mute'),
                             str(sum(1 for a in open_ if a[3] == spec))], bar=False))
    wanted = dense(StatTable([('', False), ('Looking for', False), ('Wednesday', True), ('Applying', True)], rows, 2, bar=False))
    kills = [n[3] for n in D.NIGHTS]
    facts = Figures([fig('40', 'Raiders'), fig('1/1', 'Onyxia'), fig(clock(min(kills)), 'Best kill'), fig(str(len(D.NIGHTS)), 'Nights logged')])
    week = Rows([Row('Raid nights', 'Wed, 8:00 pm to 11:00 pm CST', link=False, slim=True),
                 Row('Next', f'{day(D.NEXT_RAID_AT)}, 8 pm &middot; {day(D.SECOND_RAID_AT)}, 8 pm', link=False, slim=True),
                 Row('Loot', 'Council of 3, the numbers shown', link=False, slim=True)], open_=True)
    officers = Rows([Row(n, f'{s} &middot; {r}', lead=crest_img(c, 28), title_color=C[c], link=False)
                     for n, c, s, r in (('Ironbrand', 'paladin', 'Holy Paladin', 'Guild master'), ('Dunmaro', 'hunter', 'Beast Mastery Hunter', 'Officer'),
                                        ('Shiverknife', 'rogue', 'Assassination Rogue', 'Officer'))], open_=True)
    main = (Section('Looking for', wanted, Link('Edit'), panel=True, frame='framed', flush=True)
            + Section('Why this guild', '<p style="padding:0 20px 16px">' + 'We raid Wednesday nights, bring flasks, and split loot by council with the numbers shown. New to Forever is fine.</p>', panel=True))
    rail = (Section('Olympus XXVII', f'<div style="padding:0 20px 12px">{facts}</div>' + week, panel=True, flush=True)
            + Section('Officers', officers, panel=True))
    notice = f'<div class="pc"><p class="notice gold">{len(open_)} open applications {Button("Review", True, True, act=True)}</p></div>'
    hdr = PageHeader('Join Olympus XXVII', '', ws_prims.Emblem('alliance', 64), meta=[Tag('Recruiting', 'kill'), 'US PvP &middot; Alliance'],
                     tabs=TabStrip(gd.TABS, 'Roster'))
    return build_page('Guild recruit, dense', _view(hdr + notice + Page('split', main, rail, panels=True)) + _ROLE_CSS, 'Guild', VIEW_OFFICER)


# ---- Settings
def settings() -> str:
    plan = Figures([fig('Raid guild', 'Plan', 'founder'), fig('$200', 'A year'), fig('11 Nov 2027', 'Renews'), fig('1 of 3', 'Seats given')])
    seats = dense(StatTable([('', False), ('Premium seat', False), ('Given', True), ('Can move', True), ('', True)], [
        StatRow(crest_img('mage', 28), '<b style="color:#3fc7eb">Frostquill</b>', 'Frost Mage', ['19 Nov', 'now', Button('Move', False, True, tone='ghost', act=True)], bar=False),
        StatRow('', '<b class="mute">Open seat</b>', '', ['&ndash;', '&ndash;', Button('Give', False, True, tone='ghost', act=True)], bar=False),
        StatRow('', '<b class="mute">Open seat</b>', '', ['&ndash;', '&ndash;', Button('Give', False, True, tone='ghost', act=True)], bar=False)], 3, bar=False))
    bill = (f'<div style="padding:0 20px 12px">{plan}</div>' + seats
            + f'<div style="display:flex;gap:8px;padding:12px 20px 16px">{Button("Manage billing", False, True, act=True)}{Button("Transfer the plan", False, True, tone="ghost", act=True)}</div>')
    discord = Rows([Row('Server', 'Olympus XXVII &middot; bot installed', trail=Tag('Connected', 'kill'), link=False, slim=True),
                    Row('Raid posts', '#raid-signups &middot; last post Thu 7 Jan, 12 pm', trail=Link('Change'), link=False, slim=True),
                    Row('Recruiting', '#recruitment', trail=Link('Change'), link=False, slim=True),
                    Row('Linked members', '31 of 40 &middot; 9 without Discord', trail=Link('List'), link=False, slim=True),
                    Row('Linked role', 'not set', trail=Link('Set'), link=False, slim=True)], open_=True)
    access = Rows([Row('Claimed by', '<span style="color:#aad372">Dunmaro</span> &middot; since 12 Nov', link=False, slim=True),
                   Row('Invite link', '39 joined &middot; made 1 Dec by Dunmaro', trail=Link('Copy') + ' ' + Link('New link'), link=False, slim=True),
                   Row('Officer ranks', 'Rank 0 and 1 &middot; Ironbrand, Dunmaro, Shiverknife', trail=Link('Change'), link=False, slim=True),
                   Row('Report visibility', 'Guild only', trail=Link('Change'), link=False, slim=True),
                   Row('Addon sync', '35 of 40 in 14 days &middot; Skyarcana, Voidcall, Felbrand', link=False, slim=True)], open_=True)
    records = Rows([Row('Attendance', '5 nights &middot; 40 raiders', trail=Link('CSV'), link=False, slim=True),
                    Row('Loot', '17 awards &middot; 5 nights', trail=Link('CSV'), link=False, slim=True),
                    Row('Guild crest', 'faction crest', lead=ws_prims.Emblem('alliance', 28), trail=Link('Choose an image'), link=False, slim=True)], open_=True)
    notice = f'<div class="pc"><p class="notice gold">Setup 7 of 8 &middot; linked Discord role left {Link("Finish &rsaquo;")}</p></div>'
    main = Section('Guild plan', bill, panel=True, frame='framed', flush=True) + Section('Discord', discord, panel=True)
    rail = Section('Members and access', access, panel=True) + Section('Records and crest', records, panel=True)
    return build_page('Guild settings, dense', _view(head('Settings', True, officer_answer()) + notice + Page('split', main, rail, panels=True)) + _ROLE_CSS, 'Guild', VIEW_OFFICER)


PAGES = {'web-guild-dense-raids': raids, 'web-guild-dense-planner': planner, 'web-guild-dense-raid-member': lambda: planner(False),
         'web-guild-dense-progression': progression, 'web-guild-dense-applications': applications, 'web-guild-dense-recruit': recruit,
         'web-guild-dense-settings': settings}

if __name__ == '__main__':
    import sys

    for n, f in PAGES.items():
        if not sys.argv[1:] or n in sys.argv[1:]:
            print(write(n + '.html', f()))
