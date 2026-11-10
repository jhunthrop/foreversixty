"""The dense guild pages (design/specs/2026-11-09-guild-dense.md): the officer and member Overviews, the Roster, Readiness and
Loot, drawn from the 40-raider world (guild_dense_data.py) with the shared primitives only.

  python3 design/mocks/ws_guild_dense.py            -> design/mocks/web/web-guild-dense-*.html
"""
from __future__ import annotations

from datetime import datetime
from statistics import median
from zoneinfo import ZoneInfo

import ws_prims
from ws_prims import (Button, Figures, Grid, Link, Page, Row, Rows, Section, Segmented, SpellIcon, StatRow, StatTable, Checklist,
                      build_page, crest_img, icon, write)
from guild_data import C
import guild_dense_data as D
import guild_dense_fixtures as F
from ws_guild2 import hdr, _view, _ROLE_CSS

# Additive (web ui/StatTable `dense`): 40-row tables at 40 px a row; a guild attendance strip (guild/AttendanceStrip).
ws_prims.PRIM_CSS += '''
.st.dense .st-r{min-height:40px;padding:4px 16px}.st.dense .st-h{padding:10px 16px 8px}.st.dense .st-r .f{font-size:14px}
.st.dense .st-r .nm b{font-size:14px;line-height:18px}.st.dense .st-r .nm .sub{font-size:12px;line-height:16px}
.st.dense .st-r.grp{min-height:34px;background:rgba(255,255,255,.03)}.st.dense .st-r.grp .nm b{font-size:13px;color:var(--muted);letter-spacing:.04em}
.st-r .f.kill{color:var(--kill)}.st-r .f.wipe{color:var(--wipe)}.st-r .f.gold{color:var(--gold)}
.att{display:inline-flex;gap:3px;vertical-align:middle}.att i{width:12px;height:12px;border-radius:2px;background:var(--line)}
.att i.present{background:var(--kill)}.att i.absent{background:var(--wipe)}.att i.bench{background:var(--gold)}.att i.none{background:transparent;border:1px solid var(--line)}
'''

ZONE = ZoneInfo(D.ZONE)
VIEW_OFFICER = (D.OFFICER, 'hunter', 'Beast Mastery Hunter &middot; 60')
VIEW_MEMBER = (D.MEMBER, 'priest', 'Holy Priest &middot; 60')
CELL = {'P': 'present', 'A': 'absent', 'B': 'bench', 'N': 'none'}
GROUP_WORD = {'Tank': 'Tanks', 'Healer': 'Healers', 'Melee': 'Melee', 'Ranged': 'Ranged'}
SHORT_WORD = {'attunement': 'not attuned', 'flask_food': 'no flask or food', 'enchants': 'enchants', 'talents': 'talent points'}
SLOT_WORD = {'hands': 'Hands', 'feet': 'Feet', 'chest': 'Chest', 'finger1': 'Finger 1'}


def at(iso: str) -> datetime:
    return datetime.fromisoformat(iso.replace('Z', '+00:00')).astimezone(ZONE)


def day(iso: str) -> str:
    t = at(iso)
    return f"{t.strftime('%a')} {t.day} {t.strftime('%b')}"


def name_html(r: D.Raider) -> str:
    return f'<b style="color:{C[r.cls]}">{r.name}</b>'


def strip(cells: str) -> str:
    return '<span class="att" role="img" aria-label="Last 5 nights">' + ''.join(f'<i class="{CELL[c]}"></i>' for c in cells) + '</span>'


def tone_fig(value: str, tone: str) -> str:
    return f'<span class="{tone}">{value}</span>' if tone else value


def dense(table: str) -> str:
    return table.replace('class="st', 'class="st dense', 1)


def countdown() -> str:
    left = at(D.NEXT_RAID_AT) - at(D.NOW)
    return f'in {left.days} d {left.seconds // 3600} h'


def raid_row() -> str:
    t = at(D.NEXT_RAID_AT)
    return Row("Onyxia's Lair", f'{day(D.NEXT_RAID_AT)} &middot; 8 pm CST', lead=icon('inv_misc_head_dragon_black', 40, 'q5'),
               trail=f'<span class="gold" style="font-weight:700">{countdown()}</span>', link=False)


def ready_counts() -> tuple[int, int]:
    coming = [r for r in D.ROSTER if r.answer in ('in', 'late')]
    return sum(1 for r in coming if not D.short_for_raid(r)), len(coming)


def fillers(group: str) -> str:
    rows = [r for r in D.ROSTER if r.group == group]
    names = [f'{name_html(r)} <span class="t-small">not answered</span>' for r in rows if r.answer is None]
    names += [f'{name_html(r)} <span class="t-small">bench</span>' for r in rows if r.answer == 'bench']
    return ', '.join(names)


def comp_rows() -> list[tuple]:
    out = []
    for g in D.GROUPS:
        c = D.group_counts(g)
        need = D.COMP_TARGET[D.GROUP_TARGET[g]]
        gap = need - c['in'] - c['late']
        tone = 'ok' if gap <= 0 else 'miss'
        out.append((GROUP_WORD[g], fillers(g) if gap > 0 else '', [('none', str(need)), ('ok' if c['in'] >= need else 'warn', str(c['in'])),
                                                                   ('none', str(c['late']) if c['late'] else '&ndash;'),
                                                                   (tone, str(gap) if gap > 0 else '&ndash;')]))
    return out


def wednesday(officer: bool) -> str:
    ready, coming = ready_counts()
    c = {s: len(D.answered(s)) for s in ('in', 'late', 'bench', 'out', None)}
    figs = Figures([(str(c['in']), 'In', '', 'kill', ''), (str(c['late']), 'Late', '', 'gold', ''), (str(c['bench']), 'Bench', '', '', ''),
                    (str(c['out']), 'Out', '', '', ''), (str(c[None]), 'Not answered', '', 'wipe', ''), (f'{ready}/{coming}', 'Ready', '', '', '')])
    grid = Grid(['Need', 'In', 'Late', 'Gap'], comp_rows(), caption='Comp against the target', edge=True)
    acts = ('<div style="display:flex;gap:10px;flex-wrap:wrap;padding:14px 20px 4px">' + Button(f'Ping {c[None]} not answered', True, True, act=True)
            + Button('Open the planner', False, True) + '<span class="t-small" style="align-self:center">Posted to #raids Thu 7 Jan</span></div>')
    return raid_row() + f'<div style="padding:14px 20px">{figs}</div>' + grid + (acts if officer else '')


def word_rows() -> str:
    """Needs a word: who the officer talks to before Wednesday, the reason as figures, one action each."""
    def r(name: str, why: str, act: str) -> str:
        x = D.BY_NAME[name]
        return Row(name, why, lead=crest_img(x.cls, 32), title_color=C[x.cls], trail=Button(act, False, True, act=True), link=False)
    moves = sorted(D.RATING_MOVES.items(), key=lambda kv: kv[1][1] - kv[1][0])[:2]
    rows = [r('Skyarcana', 'Not answered &middot; missed 3 of the last 4 &middot; last log 9 days ago &middot; no Discord', 'Profile'),
            r('Voidcall', 'Not answered &middot; missed 2 of the last 4 &middot; last log 8 days ago', 'Ping'),
            r('Sunmantle', 'Late for Wednesday &middot; not attuned &middot; 3 talent points unspent', 'Nudge'),
            r('Ashblade', 'Hands and Feet not enchanted &middot; no flask or food', 'Nudge')]
    rows += [r(n, f'Rating {a} to {b} on Onyxia &middot; {day(D.night_at(D.LAST_RAID_ID))}', 'Report') for n, (a, b) in moves]
    return Rows(rows, open_=True)


def last_night() -> str:
    def row(name: str) -> str:
        x = D.BY_NAME[name]
        a, b = x.ratings
        ch = b - a
        return StatRow(crest_img(x.cls, 28), name_html(x), f'{x.spec} {x.cls.title()}',
                       [str(b), tone_fig(f'{ch:+d}', 'kill' if ch > 0 else 'wipe'), f'{x.figures[1]:,.0f} {x.metric.upper()}'], bar=False)
    moved = sorted(D.ROSTER, key=lambda x: x.ratings[1] - x.ratings[0])
    names = [x.name for x in reversed(moved[-3:])] + [x.name for x in moved[:3]]
    return dense(StatTable([('', False), ('Raider', False), ('Rating', True), ('vs 30 Dec', True), ('Onyxia', True)],
                           [row(n) for n in names], 3, bar=False))


def cover_grid() -> str:
    rows = []
    for c in F.cover()[:6]:
        signed = c['providers_in'] + c['providers_late']
        tone = 'miss' if signed < c['need'] else 'ok'
        sub = f"{c['providers_bench']} on the bench" if c['providers_bench'] else ''
        rows.append((c['name'], sub, [('none', str(c['need'])), (tone, str(signed))], SpellIcon(c['name'], c['icon'])))
    return Grid(['Need', 'In'], rows, caption='Buff and debuff cover', edge=True)


def attendance_grid() -> str:
    cell = {'P': ('ok', '&#10003;'), 'A': ('miss', '&times;'), 'B': ('warn', 'B'), 'N': ('none', '&ndash;')}
    rows = [x for x in D.ROSTER if 'A' in x.cells[-4:] or 'B' in x.cells]
    rows.sort(key=lambda x: (-x.cells[-4:].count('A'), x.name))
    cols = [day(n[1]).split(' ', 1)[1] for n in D.NIGHTS]
    return Grid(cols, [(x.name, f'{x.present} of {x.nights}', [cell[c] for c in x.cells]) for x in rows[:8]], caption='Attendance', edge=True)


def loot_last_night() -> str:
    won = [a for a in D.AWARDS if a[0] == D.LAST_RAID_ID]
    rows = [Row(a[2], f'{name_html(D.BY_NAME[a[5]])} &middot; {a[4]}', lead=icon(a[3], 32, 'q4'), title_color='#a335ee', link=False) for a in won]
    rows += [Row(d[1], f'Open &middot; {d[3]}', lead=icon(d[2], 32, 'q4'), title_color='#a335ee', trail=Button('Award', False, True, act=True), link=False)
             for d in D.OPEN_DROPS]
    return Rows(rows, open_=True)


def action_line() -> str:
    return ('<div class="pc" style="padding-top:20px"><div style="display:flex;gap:10px;flex-wrap:wrap;align-items:center">'
            + Button('3 waiting for approval', False, True, act=True) + Button('2 applications', False, True, act=True)
            + Button('2 drops to award', False, True, act=True) + '</div></div>')


def officer_overview() -> str:
    night = next(n for n in D.NIGHTS if n[0] == D.LAST_RAID_ID)
    kill = f'{night[3] // 60000}:{night[3] // 1000 % 60:02d}'
    main = (Section('Wednesday', wednesday(True), Link('Edit target &rsaquo;'), panel=True, frame='framed', flush=True)
            + Section('Needs a word', word_rows(), meta=f'{6} raiders', panel=True)
            + Section('Last night', last_night(), Link('Report &rsaquo;'),
                      meta=f'Onyxia {kill} (best) &middot; {night[4]} wipes &middot; {night[5]} deaths &middot; 37 raiders', panel=True, flush=True))
    rail = (Section('Cover', cover_grid(), Link('All &rsaquo;'), panel=True, flush=True)
            + Section('Attendance', attendance_grid(), Link('Roster &rsaquo;'), panel=True, flush=True)
            + Section(f'Loot, {day(night[1])}', loot_last_night(), Link('Loot &rsaquo;'), panel=True))
    return build_page('Guild, officer, dense', _view(hdr(raiders=40) + action_line() + Page('split', main, rail, panels=True)) + _ROLE_CSS, 'Guild', VIEW_OFFICER)


# ---- member
def your_week() -> str:
    me = D.BY_NAME[D.MEMBER]
    st = F.standing(D.MEMBER)
    you = Row(me.name, f'{me.spec} {me.cls.title()} &middot; you are <strong>In</strong> as Healer', lead=crest_img(me.cls, 36), title_color=C[me.cls],
              trail=Segmented([(w, w == 'In') for w in ('In', 'Late', 'Bench', 'Out')], small=True), link=False)
    won = D.items_won(me.name)
    figs = Figures([(str(me.ilvl), f"Item level, {st['rank_by_item_level']} of {st['same_spec_count']} Holy Priests", '', '', ''),
                    (f'{me.present}/{me.nights}', 'Nights, streak 5', '', 'kill', ''), (f'{me.best_pct}', 'Best parse %', '', '', ''),
                    (str(len(won)), 'Items this tier', '', '', '')])
    return raid_row() + you + f'<div style="padding:12px 0 4px">{figs}</div>'


def before_wednesday() -> str:
    me = D.BY_NAME[D.MEMBER]
    return Checklist([('ok', 'Attuned', 'Drakefire Amulet', None),
                      ('warn', f'{me.upgrades} upgrades, +{me.gap} HPS', 'Hands, Finger 1', 'BiS list &rsaquo;'),
                      ('ok', 'Enchants', 'Every slot', None), ('warn', 'No flask or buff food in your bags', 'Flask of Distilled Wisdom, Nightfin Soup', None)])


def your_onyxia() -> str:
    me = D.BY_NAME[D.MEMBER]
    healers = sorted((x for x in D.ROSTER if x.group == 'Healer' and x.cells[-1] == 'P'), key=lambda x: -x.figures[1])
    rows = []
    for i, n in enumerate(D.NIGHTS[-2:][::-1]):
        week = 1 - i
        prev = me.figures[0] if week == 1 else None
        ch = '' if prev is None else tone_fig(f'{me.figures[1] - prev:+,.0f}', 'kill' if me.figures[1] >= prev else 'wipe')
        rows.append(StatRow(icon('inv_misc_head_dragon_black', 28, 'q5'), f'<b>{day(n[1])}</b>', 'Onyxia kill',
                            [f'{me.figures[week]:,.0f}', str(me.ratings[week]), ch or '&ndash;'], bar=False))
    table = dense(StatTable([('', False), ('Night', False), ('HPS', True), ('Rating', True), ('Change', True)], rows, 3, bar=False))
    return table + f'<p class="t-small" style="padding:10px 20px 14px">{healers.index(me) + 1} of {len(healers)} healers by HPS on {day(D.NIGHTS[-1][1])}</p>'


def member_overview() -> str:
    me = D.BY_NAME[D.MEMBER]
    won = D.items_won(me.name)
    same = sum(1 for x in D.ROSTER if len(D.items_won(x.name)) == len(won))
    loot = Rows([Row(a[2], f'{day(D.night_at(a[0]))} &middot; {a[4]}', lead=icon(a[3], 32, 'q4'), title_color='#a335ee', link=False) for a in won],
                open_=True) + f'<p class="t-small" style="padding-top:10px">{len(D.AWARDS)} items over {len(D.NIGHTS)} nights &middot; {same} raiders on {len(won)}</p>'
    main = (Section('Your week', your_week(), panel=True, frame='framed')
            + Section('Before Wednesday', before_wednesday(), panel=True)
            + Section('Your Onyxia', your_onyxia(), Link('Report &rsaquo;'), panel=True, flush=True))
    rail = (Section("Wednesday's comp", Grid(['Need', 'In', 'Late', 'Gap'], [(r[0], '', r[2]) for r in comp_rows()], caption='Comp', edge=True), panel=True, flush=True)
            + Section('Your loot', loot, panel=True))
    return build_page('Guild, member, dense', _view(hdr(officer=False, raiders=40) + Page('split', main, rail, panels=True)) + _ROLE_CSS, 'Guild', VIEW_MEMBER)


# ---- roster
ROSTER_HEADS = [('', False), ('Raider', False), ('Last 5', False), ('Item level', True), ('Gap', True), ('Rating', True), ('Best', True),
                ('Nights', True), ('Loot', True), ('Ready', True), ('Seen', True)]


def roster_row(x: D.Raider) -> str:
    short = D.short_for_raid(x)
    gap = '&ndash;' if x.gap is None else f'+{x.gap:.0f}'
    d = x.ilvl - D.MEDIAN_ILVL
    seen = 'today' if x.seen_days == 0 else f'{x.seen_days} d'
    rank = {'leader': ' &middot; Guild master', 'officer': ' &middot; Officer'}.get(x.rank, '')
    phone = f'<span class="m"><br>{x.ilvl} ilvl &middot; {x.present}/{x.nights} nights {strip(x.cells)}</span>'
    return StatRow(crest_img(x.cls, 28), name_html(x), f'{x.spec} {x.cls.title()}{rank}{phone}',
                   [f'{x.ratings[1]}', f'{x.ilvl} <span class="t-small">{d:+d}</span>', gap, f'{x.best_pct}',
                    f'{x.present}/{x.nights}', str(len(D.items_won(x.name))) or '0',
                    tone_fig('&#10003;' if not short else f'{len(short)} short', 'kill' if not short else 'wipe'),
                    tone_fig(seen, 'wipe' if x.seen_days >= 7 else '')], meter=strip(x.cells))


def roster_table() -> str:
    rows = []
    for g in D.GROUPS:
        members = sorted((x for x in D.ROSTER if x.group == g), key=lambda x: -x.ratings[1])
        need = D.COMP_TARGET[D.GROUP_TARGET[g]]
        rows.append(StatRow('', f'<b>{GROUP_WORD[g]} {len(members)}</b>', f'need {need} &middot; median item level {median(x.ilvl for x in members):.0f}',
                            [''] * 8, bar=True).replace('st-r stc', 'st-r stc grp', 1))
        rows += [roster_row(x) for x in members]
    heads = [ROSTER_HEADS[0], ROSTER_HEADS[1], ROSTER_HEADS[2], ('Rating', True), ('Item level', True), ('Gap', True), ('Best %', True),
             ('Nights', True), ('Loot', True), ('Ready', True), ('Seen', True)]
    return dense(StatTable(heads, rows, 8, sort='Rating', sortable=tuple(h for h, r in heads if r)))


def roster() -> str:
    approvals = ('<div class="pc" style="padding-top:20px"><p class="notice gold">3 waiting for approval: Brambleshot, Mirelocke, Oakhallow '
                 + Button('Review', False, True, act=True) + ' ' + Button('Approve all', True, True, act=True) + '</p></div>')
    meta = f'40 raiders &middot; median item level {D.MEDIAN_ILVL} &middot; {sum(1 for x in D.ROSTER if x.cells[-1] == "P")} at Onyxia {day(D.NIGHTS[-1][1])}'
    body = Section('Roster', roster_table(), Button('Copy CSV', False, True, tone='ghost', act=True), meta=meta, panel=True, flush=True)
    return build_page('Guild roster, dense', _view(hdr('Roster', raiders=40, sub='Members') + approvals + Page('ledger', body)) + _ROLE_CSS, 'Guild', VIEW_OFFICER)


# ---- readiness
def readiness() -> str:
    coming = [x for x in D.ROSTER if x.answer in ('in', 'late')]
    coming.sort(key=lambda x: (-len(D.short_for_raid(x)), x.name))

    def ok(flag: bool, word: str = '&#10003;') -> str:
        return tone_fig(word, '') if flag else ''

    rows = []
    for x in coming:
        short = D.short_for_raid(x)
        miss = D.MISSING_ENCHANTS.get(x.name, [])
        fr = F.readiness_row(x)['resistances']['fire']
        why = ', '.join(SHORT_WORD[s] for s in short)
        phone = f'<span class="m"><br>{why}</span>' if short else ''
        rows.append(StatRow(crest_img(x.cls, 28), name_html(x), f'{x.spec} {x.cls.title()}{phone}',
                            [tone_fig(f'{len(short)} short', 'wipe') if short else tone_fig('Ready', 'kill'),
                             tone_fig('&times;', 'wipe') if x.name in D.NOT_ATTUNED else '&#10003;',
                             f'{x.ilvl - D.MEDIAN_ILVL:+d}', '&ndash;' if x.gap is None else f'+{x.gap:.0f}',
                             tone_fig(', '.join(SLOT_WORD[s] for s in miss), 'wipe') if miss else '&#10003;',
                             tone_fig('&times;', 'wipe') if x.name in D.NO_FLASK_FOOD else '&#10003;',
                             tone_fig(str(D.UNSPENT_TALENTS[x.name]), 'wipe') if x.name in D.UNSPENT_TALENTS else '&#10003;',
                             str(fr) if x.group == 'Tank' else '&ndash;',
                             Button('Nudge', False, True, act=True) if short else ''], bar=False))
    heads = [('', False), ('Raider', False), ('State', True), ('Attuned', True), ('Item level', True), ('Gap', True), ('Enchants', True),
             ('Flask, food', True), ('Talents', True), ('Fire res', True), ('', True)]
    ready, n = ready_counts()
    table = dense(StatTable(heads, rows, 9, bar=False, sort='State'))
    return build_page('Guild readiness, dense', _view(hdr('Readiness', raiders=40) + Page('ledger', Section(
        f'{ready} of {n} ready for {day(D.NEXT_RAID_AT)}', table, Button(f'Nudge all {n - ready} short', True, True, act=True),
        meta='In and Late', panel=True, flush=True))) + _ROLE_CSS, 'Guild', VIEW_OFFICER)


# ---- loot
def drops_table() -> str:
    rows = []
    for d in D.OPEN_DROPS:
        top = D.CANDIDATES[0] if d[0] == D.SELECTED_DROP else ('Ironbrand', None)
        rows.append(StatRow(icon(d[2], 28, 'q4'), f'<b style="color:#a335ee">{d[1]}</b>', f'{d[3]} &middot; Onyxia',
                            [tone_fig('Open', 'gold'), name_html(D.BY_NAME[top[0]]), f'+{top[1]} DPS' if top[1] else '&ndash;'], bar=False))
    for a in (x for x in D.AWARDS if x[0] == D.LAST_RAID_ID):
        rows.append(StatRow(icon(a[3], 28, 'q4'), f'<b style="color:#a335ee">{a[2]}</b>', f'{a[4]} &middot; Onyxia',
                            ['Awarded', name_html(D.BY_NAME[a[5]]), f'{a[7]} of 3'], bar=False))
    return dense(StatTable([('', False), ('Drop', False), ('State', True), ('Top or winner', True), ('Gain or votes', True)], rows, 3, bar=False))


def council() -> str:
    rows = []
    for name, gain, worn, bis, ms_os, yes in D.CANDIDATES:
        x = D.BY_NAME[name]
        last = D.last_award_raid(name)
        rows.append(StatRow(crest_img(x.cls, 28), name_html(x), f'over {worn[1]}' + (' &middot; BiS' if bis else '') + f' &middot; {ms_os.upper()}',
                            [f'+{gain}', f'{x.present}/{x.nights}', str(len(D.items_won(name))), day(D.night_at(last)) if last else 'never',
                             f'{yes} of 3', Button('Award', yes == 2, True, act=True)], bar=False))
    return dense(StatTable([('', False), ('Candidate', False), ('DPS', True), ('Nights', True), ('Items', True), ('Last item', True), ('Yes', True), ('', True)],
                           rows, 6, bar=False, sort='DPS'))


def fairness_bands() -> list[tuple[str, list[D.Raider]]]:
    """Every raider by items won this tier, then by nights attended: the council's who-is-due read."""
    won = {x.name: len(D.items_won(x.name)) for x in D.ROSTER}
    bands = []
    for n in sorted(set(won.values())):
        for full in (True, False):
            rows = [x for x in D.ROSTER if won[x.name] == n and (x.present == x.nights) == full]
            if rows:
                bands.append((n, full, sorted(rows, key=lambda x: (-x.present, x.name))))
    return bands


def fairness() -> str:
    rows = []
    for n, full, raiders in fairness_bands():
        label = f'{n} item{"" if n == 1 else "s"}' + (', every night' if full else ', missed a night')
        names = ', '.join(f'{name_html(x)} <span class="t-small">{x.present}/{x.nights}</span>' for x in raiders)
        rows.append((label, names, [('none', str(len(raiders)))]))
    return Grid(['Raiders'], rows, caption='Items won this tier by raider', edge=True)


def decisions() -> str:
    rows = [Row(f'{a[2]} to {a[5]}', f'{a[7]} yes, {3 - a[7]} no &middot; by {a[6]} &middot; {day(D.night_at(a[0]))}', lead=icon(a[3], 32, 'q4'), link=False)
            for a in reversed(D.AWARDS[-6:])]
    return Rows(rows, open_=True)


def loot() -> str:
    main = (Section(f'Onyxia, {day(D.NIGHTS[-1][1])}', drops_table(), meta='2 open', panel=True, flush=True)
            + Section("Vis'kag the Bloodletter", council(), meta='One-Hand Sword &middot; 4 candidates', panel=True, frame='framed', flush=True)
            + Section('Fairness, this tier', fairness(), Button('Export', False, True, tone='ghost', act=True),
                      meta=f'{len(D.AWARDS)} items over {len(D.NIGHTS)} nights', panel=True, flush=True))
    rail = Section('Decision record', decisions(), Link('Export history'), panel=True)
    return build_page('Guild loot, dense', _view(hdr('Loot', raiders=40) + Page('split', main, rail, panels=True)) + _ROLE_CSS, 'Guild', VIEW_OFFICER)


PAGES = {'web-guild-dense-officer': officer_overview, 'web-guild-dense-member': member_overview, 'web-guild-dense-roster': roster,
         'web-guild-dense-readiness': readiness, 'web-guild-dense-loot': loot}

if __name__ == '__main__':
    import sys
    for n, f in PAGES.items():
        if not sys.argv[1:] or n in sys.argv[1:]:
            print(write(n + '.html', f()))
