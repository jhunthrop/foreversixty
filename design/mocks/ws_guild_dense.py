"""The dense guild pages (design/specs/2026-11-09-guild-dense.md): the officer and member Overviews, the Roster, Readiness and
Loot, drawn from the 40-raider world (guild_dense_data.py) with the shared primitives only. Round 2: the mock-r1 verdicts.

  python3 design/mocks/ws_guild_dense.py            -> design/mocks/web/web-guild-dense-*.html
"""
from __future__ import annotations

from datetime import datetime
from statistics import median
from zoneinfo import ZoneInfo

import ws_prims
from ws_prims import (Button, Figures, Link, Page, PageHeader, Row, Rows, Section, Segmented, SpellIcon, StatRow, StatTable, TabStrip,
                      Emblem, build_page, crest_img, icon, write)
from guild_data import C, TABS
import guild_data as gd
import guild_dense_data as D
import guild_dense_fixtures as F
from ws_guild2 import _view, _ROLE_CSS, band_facts

# Additive (web ui/StatTable `dense`, ui/Figure `caption`): 40-row tables at 40 px a row. The attendance strip and the filler
# chips are guild components (guild/AttendanceStrip, guild/FillerChip).
ws_prims.PRIM_CSS += '''
.st.dense{--st-name:minmax(160px,2fr);--st-bar:minmax(0,96px);--st-fig:minmax(max-content,1fr)}
.legend{white-space:nowrap;flex-wrap:wrap}.att i.none{background:repeating-linear-gradient(45deg,transparent 0 2px,var(--line) 2px 3px)}
.st-r .nm .sub .m .btn{display:inline-flex;margin-top:6px}.fig .cap{border:0;padding:0;background:none;font-size:12px;color:var(--muted)}
.fchip.done{border-color:transparent;color:var(--muted)}
.st.dense .st-r{min-height:40px;padding:4px 16px}.st.dense .st-h{padding:10px 16px 8px}.st.dense .st-r .f{font-size:14px}
.st.dense .st-r .nm b{font-size:14px;line-height:18px}.st.dense .st-r .nm .sub{font-size:12px;line-height:16px}
.st.dense .st-r.grp{min-height:34px;background:rgba(255,255,255,.03)}.st.dense .st-r.grp .nm b{font-size:13px;color:var(--muted)}
.st-r.sel{background:var(--hover);box-shadow:inset 3px 0 0 var(--gold)}
.st-r .f.kill,.st-r .f .kill{color:var(--kill)}.st-r .f.wipe,.st-r .f .wipe{color:var(--wipe)}.st-r .f .gold{color:var(--gold)}.st-r .f .mute{color:var(--muted);font-weight:400}
.st-r .f .p-grey{color:#9d9d9d}.st-r .f .p-green{color:#1eff00}.st-r .f .p-blue{color:#0070dd}.st-r .f .p-purple{color:#a335ee}.st-r .f .p-orange{color:#ff8000}
.att{display:inline-flex;gap:2px;vertical-align:middle}.att i{width:14px;height:14px;border-radius:2px;font:700 10px/14px var(--font-body);font-style:normal;text-align:center;color:#07090d}
.att i.present{background:var(--kill)}.att i.absent{background:var(--wipe)}.att i.bench{background:var(--gold)}.att i.none{background:transparent;box-shadow:inset 0 0 0 1px var(--line)}
.fchips{display:flex;flex-wrap:wrap;gap:6px;margin-top:4px}.fchip{display:inline-flex;align-items:center;gap:6px;height:26px;padding:0 8px;border:1px solid var(--line);border-radius:4px;font:600 12px/1 var(--font-body);color:var(--text);background:none}
.fchip b{font-weight:700}.fchip .act{color:var(--gold)}
.fig .cap{display:block}
.legend{display:inline-flex;gap:10px;align-items:center}.legend span{display:inline-flex;gap:4px;align-items:center}
'''

ZONE = ZoneInfo(D.ZONE)
VIEW_OFFICER = (D.OFFICER, 'hunter', 'Beast Mastery Hunter &middot; 60')
VIEW_MEMBER = (D.MEMBER, 'priest', 'Holy Priest &middot; 60')
CELL = {'P': ('present', '&#10003;', 'Present'), 'A': ('absent', '&times;', 'Absent'), 'B': ('bench', 'B', 'Bench'),
        'N': ('none', '', 'Before joining')}
GROUP_WORD = {'Tank': 'Tanks', 'Healer': 'Healers', 'Melee': 'Melee', 'Ranged': 'Ranged'}
SHORT_WORD = {'attunement': 'not attuned', 'flask_food': 'no flask or food', 'enchants': 'enchants', 'talents': 'talent points'}
SLOT_WORD = {'hands': 'Hands', 'feet': 'Feet', 'chest': 'Chest', 'finger1': 'Finger 1'}
PART_WORD = {'output': 'output', 'survival': 'survival', 'mechanics': 'mechanics', 'utility': 'utility', 'preparation': 'preparation',
             'activity': 'activity'}
FILLERS_SHOWN = 3
FAIRNESS_SHOWN = 12


def at(iso: str) -> datetime:
    return datetime.fromisoformat(iso.replace('Z', '+00:00')).astimezone(ZONE)


def day(iso: str) -> str:
    t = at(iso)
    return f"{t.strftime('%a')} {t.day} {t.strftime('%b')}"


def name_html(r: D.Raider) -> str:
    return f'<b style="color:{C[r.cls]}">{r.name}</b>'


def strip(cells: str) -> str:
    return ('<span class="att" role="img" aria-label="' + ', '.join(CELL[c][2] for c in cells) + '">'
            + ''.join(f'<i class="{CELL[c][0]}">{CELL[c][1]}</i>' for c in cells) + '</span>')


def legend() -> str:
    return '<span class="legend">' + ''.join(f'<span><span class="att"><i class="{CELL[c][0]}">{CELL[c][1]}</i></span>{CELL[c][2]}</span>'
                                              for c in 'PBAN') + '</span>'


def tone(value: str, t: str) -> str:
    return f'<span class="{t}">{value}</span>' if t else value


def parse_class(pct: int) -> str:
    return 'p-orange' if pct >= 95 else 'p-purple' if pct >= 75 else 'p-blue' if pct >= 50 else 'p-green' if pct >= 25 else 'p-grey'


def dense(table: str) -> str:
    return table.replace('class="st', 'class="st dense', 1)


def fig(value: str, label: str, cap: str = '', t: str = '') -> tuple:
    return (value + (f'<span class="cap">{cap}</span>' if cap else ''), label, '', t, '')


def countdown() -> str:
    left = at(D.NEXT_RAID_AT) - at(D.NOW)
    return f'in {gd.countdown_text(left)}'


def last_night_raiders() -> list[D.Raider]:
    return [x for x in D.ROSTER if x.cells[-1] == 'P']


def head(active: str, officer: bool = True, answer: str = '') -> str:
    """The guild's head (Guild.svelte): crest, the band line, the answer slot, the role's tabs."""
    shown = TABS if officer else [t for t in TABS if t != 'Settings']
    return PageHeader(gd.GUILD, '', Emblem('alliance', 64), meta=[band_facts(len(D.ROSTER))], answer=answer,
                      tabs=TabStrip(shown, active, tablist=True))


def officer_answer() -> str:
    gaps = [(g, D.COMP_TARGET[D.GROUP_TARGET[g]] - D.group_counts(g)['in'] - D.group_counts(g)['late']) for g in D.GROUPS]
    short = sum(n for _, n in gaps if n > 0)
    return (f'<span class="t-small">{day(D.NEXT_RAID_AT)} &middot; {countdown()}</span><br>'
            f'<b class="wipe" style="font-size:15px">{short} seats open</b>')


# ---- the officer's Wednesday
def ready_counts() -> tuple[int, int]:
    coming = [r for r in D.ROSTER if r.answer in ('in', 'late')]
    return sum(1 for r in coming if not D.short_for_raid(r)), len(coming)


# The state after an action, as the built chip shows it: Voidcall was pinged from this chip.
PINGED = {'Voidcall': '4:02 pm'}


def chip(x: D.Raider) -> str:
    if x.name in PINGED:
        return f'<span class="fchip done"><b style="color:{C[x.cls]}">{x.name}</b>Pinged {PINGED[x.name]}</span>'
    act = 'Ping' if x.answer is None else 'Move to In'
    return f'<button type="button" class="fchip"><b style="color:{C[x.cls]}">{x.name}</b><span class="act">{act}</span></button>'


def gap_of(g: str) -> int:
    c = D.group_counts(g)
    return D.COMP_TARGET[D.GROUP_TARGET[g]] - c['in'] - c['late']


def filler_rows(g: str) -> list[D.Raider]:
    rows = [r for r in D.ROSTER if r.group == g]
    return [r for r in rows if r.answer is None] + [r for r in rows if r.answer == 'bench']


def comp_row(g: str) -> str:
    c = D.group_counts(g)
    need = D.COMP_TARGET[D.GROUP_TARGET[g]]
    gap = gap_of(g)
    fill = filler_rows(g)
    sub = ''
    if gap > 0:
        more = len(fill) - FILLERS_SHOWN
        sub = '<div class="fchips">' + ''.join(chip(x) for x in fill[:FILLERS_SHOWN]) + (f'<span class="t-small">+{more} more</span>' if more > 0 else '') + '</div>'
        if len(fill) < gap:
            sub += f'<span class="wipe">Still {gap - len(fill)} short if all answer</span>'
    return StatRow('', f'<b>{GROUP_WORD[g]}</b>', sub, [str(need), tone(str(c['in']), 'kill' if c['in'] >= need else 'gold'),
                                                      str(c['late']) if c['late'] else '&ndash;',
                                                      tone(str(gap), 'wipe') if gap > 0 else tone('&#10003;', 'kill')], bar=False)


def names(state: str) -> str:
    return ', '.join(name_html(r) for r in D.answered(state))


def wednesday() -> str:
    ready, coming = ready_counts()
    c = {s: len(D.answered(s)) for s in ('in', 'late', 'bench', 'out', None)}
    figs = Figures([fig(str(c['in']), 'In', '', 'kill'), fig(str(c['late']), 'Late', '', 'gold'), fig(str(c['bench']), 'Bench'),
                    fig(str(c['out']), 'Out'), fig(str(c[None]), 'Not answered', '', 'wipe'), fig(f'{ready}', f'Ready of {coming} coming')])
    who = (f'<p class="t-small" style="padding:0 20px 12px">Late {names("late")} &middot; Bench {names("bench")} &middot; Out {names("out")}</p>')
    table = dense(StatTable([('', False), ('Role', False), ('Need', True), ('In', True), ('Late', True), ('Gap', True)],
                            [comp_row(g) for g in D.GROUPS], 4, bar=False))
    acts = ('<div style="display:flex;gap:10px;flex-wrap:wrap;align-items:center;padding:14px 20px 16px">'
            + Button(f'Ping {c[None]} not answered', True, True, act=True) + Link('Open the planner &rsaquo;')
            + f'<span class="t-small">Posted to #raids {day(D.NEXT_POSTED_AT)}</span></div>')
    raid = Row("Onyxia's Lair", f'{day(D.NEXT_RAID_AT)} &middot; 8 pm CST', lead=icon('inv_misc_head_dragon_black', 40, 'q5'), link=False,
               trail=f'<span class="gold" style="font-weight:700">{countdown()}</span>')
    return f'<div style="padding:0 20px">{raid}</div><div style="padding:14px 20px 8px">{figs}</div>' + who + table + acts


def word_rows() -> str:
    """Needs a word: who the officer talks to before Wednesday, the reason as figures, one action each (all secondary)."""
    def r(name: str, why: str, act: str) -> str:
        x = D.BY_NAME[name]
        return Row(name, why, lead=crest_img(x.cls, 32), title_color=C[x.cls], trail=Button(act, False, True, act=True), link=False)
    rows = [r('Skyarcana', 'Not answered &middot; missed 3 of the last 4 &middot; last log 9 d ago &middot; no Discord', 'Profile'),
            r('Voidcall', 'Not answered &middot; missed 2 of the last 4 &middot; last log 8 d ago', 'Ping'),
            r('Sunmantle', 'Late &middot; not attuned &middot; 3 talent points unspent', 'Nudge'),
            r('Ashblade', 'Hands, Feet not enchanted &middot; no flask or food', 'Nudge')]
    for n, (a, b) in sorted(D.RATING_MOVES.items(), key=lambda kv: kv[1][1] - kv[1][0])[:2]:
        part = weakest(D.BY_NAME[n])
        rows.append(r(n, f'Rating {a} to {b} on Onyxia &middot; {part[0]} {part[1]}', 'Report'))
    return Rows(rows, open_=True)


def weakest(x: D.Raider) -> tuple[str, int]:
    parts = F.rating(x)
    k = min(PART_WORD, key=lambda p: parts[p])
    return PART_WORD[k], parts[k]


def last_night() -> str:
    def row(x: D.Raider) -> str:
        a, b = x.ratings
        part = weakest(x)
        return StatRow(crest_img(x.cls, 28), name_html(x), f'{x.spec} {x.cls.title()} &middot; weakest {part[0]} {part[1]}',
                       [str(b), tone(f'{b - a:+d}', 'kill' if b > a else 'wipe'), f'{x.figures[1]:,.0f}'],
                       bar=False, primary_phone=f'{b - a:+d} vs 30 Dec')
    came = last_night_raiders()
    moved = sorted(came, key=lambda x: x.ratings[1] - x.ratings[0])
    up, down = list(reversed(moved[-3:])), moved[:3]
    rows = [StatRow('', '<b>Up</b>', '', ['', '', ''], bar=False).replace('st-r stc', 'st-r stc grp', 1)] + [row(x) for x in up]
    rows += [StatRow('', '<b>Down</b>', '', ['', '', ''], bar=False).replace('st-r stc', 'st-r stc grp', 1)] + [row(x) for x in down]
    return dense(StatTable([('', False), ('Raider', False), ('Rating', True), ('vs 30 Dec', True), ('DPS or HPS', True)],
                           rows, 3, bar=False))


def recruit_rail() -> str:
    rows = []
    for g in D.GROUPS:
        gap = gap_of(g)
        still = gap - len(filler_rows(g))
        if still <= 0:
            continue
        apps = [a for a in gd.APPLICANTS if D.APPLICANT_GROUP.get(a[1]) == g]
        sub = f'Short {still} Wednesday if all answer'
        if apps:
            a = apps[0]
            sub += f' &middot; <b style="color:{C[a[2]]}">{a[1]}</b> {a[3]} {a[2].title()}, {a[4]} ilvl, {a[5][2]}% {a[5][0].upper()}'
        rows.append(Row(GROUP_WORD[g], sub, trail=Button('Open application', False, True) if apps else Button('Post the role', False, True), link=False))
    return Rows(rows, open_=True)


def cover_rail() -> str:
    rows = []
    for c in F.cover()[:5]:
        signed = c['providers_in'] + c['providers_late']
        t = 'wipe' if signed < c['need'] else 'kill'
        bench = f" &middot; {c['providers_bench']} on the bench" if c['providers_bench'] else ''
        rows.append(StatRow(SpellIcon(c['name'], c['icon']), f"<b>{c['name']}</b>", f"need {c['need']}{bench}", [tone(str(signed), t)], bar=False))
    return dense(StatTable([('', False), ('Buff or debuff', False), ('In', True)], rows, 1, bar=False))


def slipping_rail() -> str:
    rows = [x for x in D.ROSTER if x.cells[-4:].count('A') >= 2 or x.cells.count('B') >= 2]
    rows.sort(key=lambda x: (-x.cells[-4:].count('A'), x.name))
    body = [StatRow(crest_img(x.cls, 28), name_html(x), f'{x.present} of {x.nights}', [strip(x.cells)], bar=False) for x in rows]
    return dense(StatTable([('', False), ('Raider', False), ('Last 5', True)], body, 1, bar=False)) + f'<p class="t-small" style="padding:8px 20px 14px">{legend()}</p>'


def loot_last_night() -> str:
    won = [a for a in D.AWARDS if a[0] == D.LAST_RAID_ID]
    rows = [Row(d[1], f'Open &middot; {name_html(D.BY_NAME[D.CANDIDATES[0][0]])} +{D.CANDIDATES[0][1]}' if d[0] == D.SELECTED_DROP else f'Open &middot; {d[3]}', lead=icon(d[2], 32, 'q4'), title_color='#a335ee', trail=Button('Council', False, True), link=False)
            for d in D.OPEN_DROPS]
    rows += [Row(a[2], f'{name_html(D.BY_NAME[a[5]])} &middot; {a[4]}', lead=icon(a[3], 32, 'q4'), title_color='#a335ee', link=False) for a in won]
    return Rows(rows, open_=True)


def action_line() -> str:
    return ('<div class="pc" style="padding-top:20px"><div style="display:flex;gap:10px;flex-wrap:wrap;align-items:center">'
            + Button('3 waiting for approval', False, True, act=True) + Button('2 applications', False, True, act=True)
            + Button('2 drops to award', False, True, act=True) + '</div></div>')


def officer_overview() -> str:
    night = next(n for n in D.NIGHTS if n[0] == D.LAST_RAID_ID)
    kill = f'{night[3] // 60000}:{night[3] // 1000 % 60:02d}'
    main = (Section('Wednesday', wednesday(), Link('Edit target &rsaquo;'), panel=True, frame='framed', flush=True)
            + Section('Needs a word', word_rows(), meta='6 raiders', panel=True)
            + Section('Last night', last_night(), Link('Report &rsaquo;'),
                      meta=f'Onyxia {kill}, best yet &middot; {night[4]} wipes &middot; {night[5]} deaths &middot; {len(last_night_raiders())} raiders',
                      panel=True, flush=True))
    rail = (Section('Recruit for the gaps', recruit_rail(), panel=True)
            + Section(f'Loot, {day(night[1])}', loot_last_night(), Link('Loot &rsaquo;'), panel=True)
            + Section('Cover, thinnest first', cover_rail(), Link('All &rsaquo;'), panel=True, flush=True)
            + Section('Slipping', slipping_rail(), Link('Roster &rsaquo;'), panel=True, flush=True))
    return build_page('Guild, officer, dense', _view(head('Overview', True, officer_answer()) + action_line() + Page('split', main, rail, panels=True))
                      + _ROLE_CSS, 'Guild', VIEW_OFFICER)


# ---- member
def healer_rank(me: D.Raider) -> tuple[int, int]:
    healers = sorted((x for x in last_night_raiders() if x.group == me.group), key=lambda x: -x.figures[1])
    return healers.index(me) + 1, len(healers)


def your_week() -> str:
    me = D.BY_NAME[D.MEMBER]
    st = F.standing(D.MEMBER)
    you = Row(me.name, f'{me.spec} {me.cls.title()} &middot; Healer', lead=crest_img(me.cls, 36), title_color=C[me.cls],
              trail=Segmented([(w, w == 'In') for w in ('In', 'Late', 'Bench', 'Out')], small=True), link=False)
    rank, of = healer_rank(me)
    won = D.items_won(me.name)
    figs = Figures([fig(str(me.ilvl), 'Item level', f"Holy Priest {['1st', '2nd', '3rd'][st['rank_by_item_level'] - 1]} of {st['same_spec_count']}"),
                    fig(f'{rank}', 'HPS rank', f'of {of} healers, {day(D.NIGHTS[-1][1])}'), fig(str(me.ratings[1]), 'Rating', f'{me.ratings[1] - me.ratings[0]:+d} vs 30 Dec'),
                    fig(f'{me.present}/{me.nights}', 'Nights', 'streak 5', 'kill'), fig(str(len(won)), 'Items', 'this tier')])
    raid = Row("Onyxia's Lair", f'{day(D.NEXT_RAID_AT)} &middot; 8 pm CST', lead=icon('inv_misc_head_dragon_black', 40, 'q5'), link=False,
               trail=f'<span class="gold" style="font-weight:700">{countdown()}</span>')
    return raid + you + f'<div style="padding:14px 0 4px">{figs}</div>'


def before_wednesday() -> str:
    me = D.BY_NAME[D.MEMBER]
    def r(mark: str, t: str, title: str, sub: str, link: bool) -> str:
        lead = f'<span class="ck {t}" role="img" aria-label="{"Done" if t == "ok" else "To do"}">{mark}</span>'
        return Row(title, sub, lead=lead, trail='&rsaquo;' if link else '', link=link)
    return Rows([r('!', 'warn', f'{me.upgrades} upgrades, +{me.gap:.0f} HPS', 'Hands, Finger 1', True),
                 r('!', 'warn', 'No flask or buff food in your bags', 'Flask of Distilled Wisdom, Nightfin Soup', False),
                 r('&#10003;', 'ok', 'Attuned', 'Drakefire Amulet', False), r('&#10003;', 'ok', 'Enchants', 'Every slot', False)], open_=True)


def your_onyxia() -> str:
    me = D.BY_NAME[D.MEMBER]
    rows = []
    for week, n in ((1, D.NIGHTS[-1]), (0, D.NIGHTS[-2])):
        ch = tone(f'{me.figures[1] - me.figures[0]:+,.0f}', 'kill' if me.figures[1] >= me.figures[0] else 'wipe') if week else '&ndash;'
        deaths = D.DEATHS.get(n[0], {}).get(me.name, 0)
        rows.append(StatRow(icon('inv_misc_head_dragon_black', 28, 'q5'), f'<b>{day(n[1])}</b>', f'kill {n[3] // 60000}:{n[3] // 1000 % 60:02d}',
                            [f'{me.figures[week]:,.0f}', str(me.ratings[week]), str(deaths), ch], bar=False))
    return dense(StatTable([('', False), ('Night', False), ('HPS', True), ('Rating', True), ('Deaths', True), ('HPS vs last', True)], rows, 4, bar=False))


def member_comp() -> str:
    rows = []
    for g in D.GROUPS:
        c = D.group_counts(g)
        need = D.COMP_TARGET[D.GROUP_TARGET[g]]
        gap = gap_of(g)
        label = f'<b>{GROUP_WORD[g]}</b>' + (' <span class="t-small">you</span>' if g == D.BY_NAME[D.MEMBER].group else '')
        rows.append(StatRow('', label, '', [str(need), str(c['in']), tone(str(gap), 'wipe') if gap > 0 else tone('&#10003;', 'kill')], bar=False))
    return dense(StatTable([('', False), ('Role', False), ('Need', True), ('In', True), ('Gap', True)], rows, 3, bar=False))


def member_overview() -> str:
    me = D.BY_NAME[D.MEMBER]
    won = D.items_won(me.name)
    per = len(won) / me.present
    guild_per = len(D.AWARDS) / sum(x.present for x in D.ROSTER)
    loot = Rows([Row(a[2], f'{day(D.night_at(a[0]))} &middot; {a[4]}', lead=icon(a[3], 32, 'q4'), title_color='#a335ee', link=False) for a in won],
                open_=True) + Figures([fig(f'{per:.2f}', 'Your items per night'), fig(f'{guild_per:.2f}', 'Guild per night')])
    night = D.NIGHTS[-1]
    guild_line = Rows([Row(f'Onyxia {night[3] // 60000}:{night[3] // 1000 % 60:02d}', f'{day(night[1])} &middot; best kill yet &middot; {night[5]} deaths',
                           lead=icon('inv_misc_head_dragon_black', 32, 'q5'), trail=Link('Report &rsaquo;'), link=False)], open_=True)
    main = (Section('Your week', your_week(), panel=True, frame='framed')
            + Section('Before Wednesday', before_wednesday(), meta='2 to do', panel=True)
            + Section('Your Onyxia', your_onyxia(), Link('Report &rsaquo;'), panel=True, flush=True))
    rail = (Section("Wednesday's comp", member_comp(), panel=True, flush=True) + Section('Your loot', loot, panel=True)
            + Section('Last night', guild_line, panel=True))
    answer = f'<span class="t-small">You are <b>In</b> &middot; {day(D.NEXT_RAID_AT)}</span><br><b class="gold" style="font-size:15px">{countdown()}</b>'
    return build_page('Guild, member, dense', _view(head('Overview', False, answer) + Page('split', main, rail, panels=True)) + _ROLE_CSS, 'Guild', VIEW_MEMBER)


# ---- roster
def roster_row(x: D.Raider) -> str:
    short = D.short_for_raid(x)
    gap = '&ndash;' if x.gap is None else f'+{x.gap:.0f}'
    d = x.ilvl - D.MEDIAN_ILVL
    seen = '' if x.seen_days < 2 else tone(f'{x.seen_days} d', 'wipe' if x.seen_days >= 7 else 'mute')
    rank = {'leader': ' &middot; Guild master', 'officer': ' &middot; Officer'}.get(x.rank, '')
    pct = round(100 * x.present / x.nights)
    phone = f'<span class="m"><br>{x.ilvl} ilvl &middot; {pct}% {strip(x.cells)}</span>'
    return StatRow(crest_img(x.cls, 28), name_html(x), f'{x.spec} {x.cls.title()}{rank}{phone}',
                   [str(x.ratings[1]), f'{x.ilvl} <span class="mute">{d:+d}</span>', gap, tone(str(x.best_pct), parse_class(x.best_pct)),
                    f'{pct}%', str(len(D.items_won(x.name))),
                    tone('&#10003;', 'kill') if not short else f'<a class="wipe" href="#">{len(short)} short</a>', seen],
                   meter=strip(x.cells), primary_phone='rating')


def roster_table() -> str:
    rows = []
    for g in D.GROUPS:
        members = sorted((x for x in D.ROSTER if x.group == g), key=lambda x: -x.ratings[1])
        need = D.COMP_TARGET[D.GROUP_TARGET[g]]
        rows.append(StatRow('', f'<b>{GROUP_WORD[g]} {len(members)}</b>', f'need {need} &middot; median item level {median(x.ilvl for x in members):.0f}',
                            [''] * 8).replace('st-r stc', 'st-r stc grp', 1))
        rows += [roster_row(x) for x in members]
    heads = [('', False), ('Raider', False), ('Last 5', False), ('Rating', True), ('Item level', True), ('To BiS', True), ('Best %', True),
             ('Nights', True), ('Loot', True), ('Ready', True), ('Seen', True)]
    return dense(StatTable(heads, rows, 8, sort='Rating', sortable=tuple(h for h, r in heads if r)))


def roster_controls() -> str:
    counts = {g: sum(1 for x in D.ROSTER if x.group == g) for g in D.GROUPS}
    seg = Segmented([(f'All {len(D.ROSTER)}', True)] + [(f'{GROUP_WORD[g]} {counts[g]}', False) for g in D.GROUPS], small=True)
    return (f'<div style="display:flex;gap:12px;flex-wrap:wrap;align-items:center;padding:0 20px 12px">{seg}'
            f'<input class="field" aria-label="Find a raider" placeholder="Find a raider" style="height:32px;width:200px">'
            f'<span class="t-small" style="margin-left:auto">{legend()}</span></div>')


def roster() -> str:
    approvals = ('<div class="pc" style="padding-top:20px"><p class="notice gold">3 waiting for approval: Brambleshot, Mirelocke, Oakhallow '
                 + Button('Review', False, True, act=True) + ' ' + Button('Approve all', True, True, act=True) + '</p></div>')
    meta = (f'{len(D.ROSTER)} raiders &middot; item level vs guild median {D.MEDIAN_ILVL} &middot; nights over the last 30 &middot; '
            f'{len(last_night_raiders())} at Onyxia {day(D.NIGHTS[-1][1])}')
    body = Section('Roster', roster_controls() + roster_table(), Button('Copy CSV', False, True, tone='ghost', act=True), meta=meta, panel=True, flush=True)
    return build_page('Guild roster, dense', _view(head('Roster', True, officer_answer()) + approvals + Page('ledger', body)) + _ROLE_CSS, 'Guild', VIEW_OFFICER)


# ---- readiness
STATE_WORD = {'bench': 'Bench', 'out': 'Out', None: 'No answer'}


def readiness_row(x: D.Raider) -> str:
    """A raider's checks for Wednesday. Raiders not coming read muted with their answer as the state, and a Ping when unanswered."""
    coming = x.answer in ('in', 'late')
    short = D.short_for_raid(x) if coming else []
    miss = D.MISSING_ENCHANTS.get(x.name, [])
    ok = tone('&#10003;', 'mute')
    nudge = Button('Nudge', False, True, act=True) if short else Button('Ping', False, True, act=True) if x.answer is None else ''
    why = ', '.join(SHORT_WORD[s] for s in short)
    phone = f'<span class="m"><br>{why} {nudge}</span>' if why or nudge else ''
    synced = tone(f'{x.seen_days} d', 'wipe') if x.seen_days >= 7 else tone('today' if x.seen_days == 0 else f'{x.seen_days} d', 'mute')
    state = (tone(f'{len(short)} short', 'wipe') if short else tone('Ready', 'mute')) if coming else tone(STATE_WORD[x.answer], 'mute')

    def check(fails: bool, word: str = '&times;') -> str:
        return tone(word, 'wipe' if coming else 'mute') if fails else ok
    return StatRow(crest_img(x.cls, 28), name_html(x), f'{x.spec} {x.cls.title()}{phone}',
                   [state, check(x.name in D.NOT_ATTUNED), f'{x.ilvl - D.MEDIAN_ILVL:+d}', '&ndash;' if x.gap is None else f'+{x.gap:.0f}',
                    check(bool(miss), ', '.join(SLOT_WORD[s] for s in miss)), check(x.name in D.NO_FLASK_FOOD),
                    check(x.name in D.UNSPENT_TALENTS, str(D.UNSPENT_TALENTS.get(x.name, ''))), synced, nudge], bar=False)


def readiness() -> str:
    coming = sorted((x for x in D.ROSTER if x.answer in ('in', 'late')), key=lambda x: (-len(D.short_for_raid(x)), x.name))
    rest = sorted((x for x in D.ROSTER if x.answer not in ('in', 'late')), key=lambda x: (x.answer is not None, x.name))
    rows = [readiness_row(x) for x in coming]
    rows.append(StatRow('', f'<b>Not coming {len(rest)}</b>', 'Bench, Out and no answer', [''] * 9, bar=False).replace('st-r stc', 'st-r stc grp', 1))
    rows += [readiness_row(x) for x in rest]
    heads = [('', False), ('Raider', False), ('State', True), ('Attuned', True), ('Item level', True), ('To BiS', True), ('Enchants', True),
             ('Flask, food', True), ('Talents', True), ('Synced', True), ('', True)]
    ready, n = ready_counts()
    table = dense(StatTable(heads, rows, 9, bar=False, sort='State', sortable=tuple(h for h, r in heads if r and h)))
    return build_page('Guild readiness, dense', _view(head('Readiness', True, officer_answer()) + Page('ledger', Section(
        f'{ready} of {n} ready for {day(D.NEXT_RAID_AT)}', table, Button(f'Nudge all {n - ready} short', True, True, act=True),
        meta=f'In and Late &middot; item level vs guild median {D.MEDIAN_ILVL}', panel=True, flush=True))) + _ROLE_CSS, 'Guild', VIEW_OFFICER)


# ---- loot
def drops_table() -> str:
    rows = []
    for d in D.OPEN_DROPS:
        selected = d[0] == D.SELECTED_DROP
        top = D.CANDIDATES[0] if selected else None
        row = StatRow(icon(d[2], 28, 'q4'), f'<b style="color:#a335ee">{d[1]}</b>', f'{d[3]} &middot; Onyxia',
                      [tone('Open', 'gold'), name_html(D.BY_NAME[top[0]]) if top else tone('no sim gain yet', 'mute'),
                       f'+{top[1]}' if top else '&ndash;', '&ndash;'], bar=False, primary_phone=f'+{top[1]} {top[0]}' if top else '')
        rows.append(row.replace('st-r stc', 'st-r stc sel', 1) if selected else row)
    for a in (x for x in D.AWARDS if x[0] == D.LAST_RAID_ID):
        rows.append(StatRow(icon(a[3], 28, 'q4'), f'<b style="color:#a335ee">{a[2]}</b>', f'{a[4]} &middot; Onyxia',
                            [tone('Awarded', 'mute'), name_html(D.BY_NAME[a[5]]), '&ndash;', f'{a[7]} of 3'], bar=False, primary_phone=a[5]))
    return dense(StatTable([('', False), ('Drop', False), ('State', True), ('Top candidate or winner', True), ('DPS', True), ('Votes', True)],
                           rows, 4, bar=False))


def council() -> str:
    rows = []
    for i, (name, gain, worn, bis, ms_os, yes) in enumerate(D.CANDIDATES):
        x = D.BY_NAME[name]
        last = D.last_award_raid(name)
        won = len(D.items_won(name))
        last_word = day(D.night_at(last)) if last else 'never'
        award = Button('Award', i == 0, True, act=True)
        phone = f'<span class="m"><br>{x.present}/{x.nights} nights &middot; {won} items &middot; {last_word} &middot; {yes} of 3 yes<br>{award}</span>'
        rows.append(StatRow(crest_img(x.cls, 28), name_html(x), f'over {worn[1]}' + (' &middot; BiS' if bis else '') + f' &middot; {ms_os.upper()}{phone}',
                            [f'+{gain}', f'{x.present}/{x.nights}', str(won), last_word, f'{yes} of 3', award], bar=False))
    return dense(StatTable([('', False), ('Candidate', False), ('DPS', True), ('Nights', True), ('Items', True), ('Last item', True), ('Yes', True), ('', True)],
                           rows, 6, bar=False, sort='DPS'))


def nights_since_item(x: D.Raider) -> int:
    """Nights the raider came since their last item this tier (every night they came, with none)."""
    last = D.last_award_raid(x.name)
    ids = [n[0] for n in D.NIGHTS]
    since = ids.index(last) + 1 if last else 0
    return x.cells[since:].count('P')


def fairness() -> str:
    """Most overdue first: nights attended since the last item, then fewest items. Raiders level on both share one row."""
    won = {x.name: len(D.items_won(x.name)) for x in D.ROSTER}
    groups: dict[tuple[int, int], list[D.Raider]] = {}
    for x in D.ROSTER:
        groups.setdefault((nights_since_item(x), won[x.name]), []).append(x)
    rows = []
    top = max(k[0] for k in groups)
    for (since, items), raiders in sorted(groups.items(), key=lambda kv: (-kv[0][0], kv[0][1])):
        raiders.sort(key=lambda x: x.name)
        meter = ws_prims.Meter(since / top)
        if len(raiders) >= 3:
            names = ', '.join(name_html(x) for x in raiders)
            rows.append(StatRow('', f'<b>{len(raiders)} raiders</b>', names,
                                [str(since), str(items), '&ndash;', 'never' if not items else '&ndash;'], meter=meter, primary_phone=f'{items} items'))
            continue
        for x in raiders:
            last = D.items_won(x.name)[-1] if items else None
            last_word = f'{last[2]}, {day(D.night_at(last[0]))}' if last else 'never'
            rows.append(StatRow(crest_img(x.cls, 28), name_html(x), f'{x.spec} {x.cls.title()}<span class="m"><br>{items} items &middot; {last_word}</span>',
                                [str(since), str(items), f'{x.present}/{x.nights}', last_word], meter=meter))
    return dense(StatTable([('', False), ('Raider', False), ('', False), ('Nights since an item', True), ('Items', True), ('Nights', True),
                            ('Last item', True)], rows, 4, sort='Nights since an item',
                           sortable=('Nights since an item', 'Items', 'Nights', 'Last item')))


def decisions() -> str:
    rows = [Row(f'{a[2]} to {a[5]}', f'{a[7]} yes, {3 - a[7]} no &middot; by {a[6]} &middot; {day(D.night_at(a[0]))}', lead=icon(a[3], 32, 'q4'), link=False)
            for a in reversed(D.AWARDS[-6:])]
    return Rows(rows, open_=True)


def loot() -> str:
    med = len(D.AWARDS) / sum(x.present for x in D.ROSTER)
    main = (Section(f'Onyxia, {day(D.NIGHTS[-1][1])}', drops_table(), meta='2 open', panel=True, flush=True)
            + Section("Vis'kag the Bloodletter", council(), meta='One-Hand Sword &middot; 4 candidates', panel=True, frame='framed', flush=True)
            + Section('Fairness, this tier', fairness(), Button('Export', False, True, tone='ghost', act=True),
                      meta=f'{len(D.AWARDS)} items over {len(D.NIGHTS)} nights &middot; guild {med:.2f} per night &middot; fewest first', panel=True, flush=True))
    rail = Section('Decision record', decisions(), Link('Export history'), panel=True)
    answer = f'<b class="gold" style="font-size:15px">2 drops to award</b><br><span class="t-small">Onyxia, {day(D.NIGHTS[-1][1])}</span>'
    return build_page('Guild loot, dense', _view(head('Loot', True, answer) + Page('split', main, rail, panels=True)) + _ROLE_CSS, 'Guild', VIEW_OFFICER)


PAGES = {'web-guild-dense-officer': officer_overview, 'web-guild-dense-member': member_overview, 'web-guild-dense-roster': roster,
         'web-guild-dense-readiness': readiness, 'web-guild-dense-loot': loot}

if __name__ == '__main__':
    import sys
    for n, f in PAGES.items():
        if not sys.argv[1:] or n in sys.argv[1:]:
            print(write(n + '.html', f()))
