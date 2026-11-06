"""Web-layout mock: a report (/reports/<id>, every tab and state, and /reports/sample), composed ONLY from ws_prims.

Spec: design/specs/2026-11-07-logs-reports.md section 6. Every figure, name, time and series is read at render time from
the files the build and its parity spec serve:
  - web/src/fixtures/report (fixture2abcd, the e2e report: meta.json, fights/<n>/summary.json, fights/<n>/events.parquet,
    the Events and Queries tabs read the parquet through DuckDB exactly as the shipped page does in the browser);
  - web/src/fixtures/coach/fury-model.json (the Coach tab; the e2e re-reads fight 3 as the Onyxia pull, as coach-stubs.ts does);
  - design/mocks/data/logs-rating-fight3.json (the Rating tab: report-rating.spec.ts's stub, gap G7);
  - web/public/sample-report/sample (the bundled sample report);
  - data/builds/<active build>/ spells.json and spellicons.json (spell names and icons, as the build resolves them).
SAMPLE (typed, marked): the signed-in viewer for the Coach tab is the owner of Baelgrim (coach-stubs.ts's premium viewer).

  python3 design/mocks/ws_report.py && for p in $(python3 design/mocks/ws_report.py --list); do node design/mocks/ws_shoot.mjs $p 390,1440,1920,2000; done
"""
import json
import statistics
import sys
from html import escape

from mocklib import CLASS_COLOR, PARSE
from ws_prims import (Button as _Button, Disclosure, EmptyCrest, Facts, Figure, Lanes, LanesLegend, LineChart, Link, Menu, Meter,
                      MeterClass, Notice, Page, PageHeader, Paperdoll, Plate, Row, Rows, Section, Skeleton, Spark, StatRow,
                      StatTable, TabStripMore, Tag, build_page, control_bar, crest_img, seg, sel, write)


from ws_report_lib import *  # noqa: F401,F403  (data, names, icons, head, bar, chart, this_pull)
from ws_report_lib import _Button  # noqa: F401


# ------------------------------------------------------------------------------------------------ tabs
def player_sub(r: dict, extra: str = '') -> str:
    bits = [b for b in (f'{r.get("spec", "")} {r.get("class", "")}'.strip(), f'item level {r["item_level"]}' if r.get('item_level') else '', extra) if b]
    return ' &middot; '.join(escape(b) if '<' not in b else b for b in bits)


def tab_summary(R, fi, source):
    s = R.summary(fi)
    fr = friendlies(R, s, source)
    rs = sorted([r for r in s['roster'] if r['guid'] in fr], key=lambda r: (-r['damage_done'], -r['healing_done'], r['name']))
    home = statistics.mode([split_name(r['name'])[1] for r in rs]) if rs else ''
    rows = []
    for r in rs:
        tagr = realm_tag(r['name'], home)
        nm = pname(r) + (f' <span class="t-small">{escape(tagr)}</span>' if tagr else '')
        sub = player_sub(r) or '&nbsp;'
        d = '&ndash;'
        rows.append(StatRow(crest(r, 28), nm, sub, ['&ndash;', n1(r['dps']), n1(r['hps']), n1(r['dtps']), f'{r["activity_pct"]:.1f}%',
                                                  f'<span class="{"wipe" if r["deaths"] else ""}">{r["deaths"]}</span>'],
                            primary_phone=f'{n1(r["dps"])} DPS', bar=False))
    table = StatTable([('', False), ('Player', False), ('Parse', True), ('DPS', True), ('HPS', True), ('Taken / s', True), ('Active', True), ('Deaths', True)],
                      rows, 6, bar=False)
    nd = sum(1 for d in s['deaths'] if d['guid'] in fr)
    plate = Section('Players', Plate(table, 'Players'), meta=f'{len(rs)} players &middot; {nd} death{"" if nd == 1 else "s"} &middot; {mmss(s["duration_ms"])}')
    # damage taken by ability, every friendly
    abil: dict[str, list] = {}
    for a in s['damage_taken']:
        if a['guid'] in fr:
            for ab in a['abilities']:
                k = ab['name']
                abil.setdefault(k, [ab.get('spell_id', 0), 0])[1] += ab.get('effective', ab['total'])
    # Share is of everything every player took in this window, whatever the source scope shows (the shipped caption)
    tot = sum(ab.get('effective', ab['total']) for a in s['damage_taken'] if a['guid'] in roster(s) for ab in a['abilities']) or 1
    dur = s['duration_ms'] / 1000
    trows = [StatRow(sicon(sid, 24), f'<b>{escape(spell(sid, k))}</b>', '', [f'{v / tot * 100:.1f}%', n0(v), n1(v / dur)], Meter(v / tot), primary_phone=f'{v / tot * 100:.1f}%')
             for k, (sid, v) in sorted(abil.items(), key=lambda kv: -kv[1][1])]
    taken = Section('Damage taken by ability', StatTable([('', False), ('Ability', False), ('', False), ('Share', True), ('Amount', True), ('Per sec', True)], trows, 3, open_=True)
                    + '<p class="t-small" style="padding-top:8px">Share is of everything every player took in this window, whatever the source scope shows.</p>'
                    if trows else '<p class="empty">Nobody took damage in this window.</p>', action=Link('Damage taken &rsaquo;'))
    main = plate + chart(R, fi, 'damage', source) + taken   # R4: the table first, the chart under it
    rail = this_pull(R, fi, source) + deaths_rail(R, fi, source) + rating_rail(R, fi) + at_pull(R, fi)
    return main, rail, ([PARSES_MISSING] if R.fight(fi)['kind'] == 'encounter' else [])


def deaths_rail(R, fi, source='All friendlies'):
    s = R.summary(fi)
    ros = roster(s)
    fr = friendlies(R, s, source)
    dd = [d for d in s['deaths'] if d['guid'] in fr]
    if not dd:
        return Section('Deaths', '<p class="empty">Nobody died.</p>')
    rows = [Row(pname(ros.get(d['guid']), d['name']), f'{escape(d["killing_blow"]["spell_name"])} &middot; {escape(d["killing_blow"]["source_name"])} &middot; {n0(d["killing_blow"]["amount"])}',
                lead=f'<span class="num muted">{at(d["at_ms"])}</span>{crest(ros.get(d["guid"]), 28)}') for d in dd]
    return Section('Deaths', Rows(rows, open_=True), action=Link('Deaths &rsaquo;'))


def rating_rail(R, fi):
    if fi != RATING['fight_index'] or R.base != FIX:
        return Section(COPY['ratingCopy']['panelHeading'], f'<p class="empty">{escape(COPY["ratingCopy"]["noRatingForTrash"] if R.fight(fi)["kind"] == "trash" else COPY["ratingCopy"]["panelEmpty"])}</p>')
    ros = {split_name(r['name'])[0]: r for r in R.summary(fi)['roster']}
    rows = [Row(pname({'class': p['class']}, p['player_name']), f'{p["spec"]} {p["class"]}', lead=crest({'class': p['class']}, 28),
                trail=Figure(str(p['overall']), unit='of 100', size='sm')) for p in RATING['players']]
    return Section(COPY['ratingCopy']['panelHeading'], Rows(rows, open_=True), action=Link(COPY['ratingCopy']['panelMore'] + ' &rsaquo;'))


def at_pull(R, fi):
    s = R.summary(fi)
    ros = roster(s)
    if not s.get('combatants'):
        return ''
    rows = [Row(pname(ros.get(c['guid']), c['name']), f'{escape(c["spec"])} &middot; item level {c["item_level"]} &middot; {sum(1 for g in c["gear"] if g["ID"])} items',
                lead=crest(ros.get(c['guid']), 28), trail='<span class="link">Gear &rsaquo;</span>') for c in s['combatants']]
    return Section('At pull', Rows(rows, open_=True))


def actor_table(R, fi, source, key, unit, head_amt):
    s = R.summary(fi)
    fr = friendlies(R, s, source)
    ros = roster(s)
    dur = s['duration_ms'] / 1000
    acts = sorted([a for a in s[key] if a['guid'] in fr and a['effective'] > 0], key=lambda a: -a['effective'])
    tot = sum(a['effective'] for a in acts) or 1
    top = acts[0]['effective'] if acts else 1
    rows = []
    for a in acts:
        r = ros.get(a['guid'], {'name': a['name'], 'class': a.get('class')})
        extra = f'{n0(a["overheal"])} overheal ({a["overheal"] / a["total"] * 100:.0f}%)' if key == 'healing' else ''
        figs = [n0(a['effective']), n1(a['effective'] / dur), f'{a["effective"] / tot * 100:.1f}%',
                f'{a["active_ms"] / s["duration_ms"] * 100:.1f}%']
        rows.append(StatRow(crest(r, 28), pname(r, a['name']), player_sub(r, extra) or '&nbsp;', figs, MeterClass(a['effective'] / top, colour(r)), primary_phone=f'{figs[1]} {unit}'))
    if acts:
        rows.append(StatRow('', f'Total, {len(acts)} player{"s" if len(acts) != 1 else ""}', '', [n0(sum(a['effective'] for a in acts)), n1(sum(a['effective'] for a in acts) / dur), '100%', ''], total=True))
    heads = [('', False), ('Player', False), ('', False), (head_amt, True), (unit, True), ('Share', True), ('Active', True)]
    return StatTable(heads, rows, 4, sort=head_amt, sortable=(head_amt, unit, 'Share', 'Active')) if acts else f'<p class="empty">Nothing in this window.</p>'


def tab_actor(R, fi, source, which):
    s = R.summary(fi)
    key, unit, amt, metric, title = {'Damage done': ('damage_done', 'DPS', 'Amount', 'damage', 'Damage done'),
                                     'Healing': ('healing', 'HPS', 'Effective', 'healing', 'Healing'),
                                     'Damage taken': ('damage_taken', 'Taken / s', 'Amount', 'taken', 'Damage taken')}[which]
    plate = Section(title, Plate(actor_table(R, fi, source, key, unit, amt), title), meta='Click a player to see their abilities')
    fr = friendlies(R, s, source)
    side = {}
    pairs = 'targets'
    for a in s[key]:
        if a['guid'] in fr:
            for t in a.get(pairs, []):
                side[t['name']] = side.get(t['name'], 0) + t['total']
    tot = sum(side.values()) or 1
    ros = roster(s)
    rail_title = {'Damage done': 'Targets', 'Healing': 'Healed', 'Damage taken': 'Sources'}[which]
    rrows = [Row(pname(ros.get(next((g for g, r in ros.items() if r['name'] == k), ''), {'name': k}), k), f'{v / tot * 100:.1f}% &middot; {n0(v)}',
                 trail=f'<span style="width:96px;display:block">{Meter(v / tot)}</span>', link=False) for k, v in sorted(side.items(), key=lambda kv: -kv[1])]
    rail = this_pull(R, fi) + Section(rail_title, Rows(rrows, open_=True) if rrows else '<p class="empty">Nothing in this window.</p>')
    return plate + chart(R, fi, metric, source), rail, []


def tab_player(R, fi, guid):
    """Source = one player, Damage done: their abilities, then the gear they wore at the pull (Paperdoll, report mode)."""
    s = R.summary(fi)
    a = next(x for x in s['damage_done'] if x['guid'] == guid)
    r = roster(s)[guid]
    dur = s['duration_ms'] / 1000
    top = max(ab['total'] for ab in a['abilities'])
    rows = [StatRow(sicon(ab['spell_id']), f'<b>{escape(spell(ab["spell_id"], ab["name"]))}</b>',
                    f'<span>{ab["hits"]} hits</span> &middot; <span>{ab["crits"]} crit{"" if ab["crits"] == 1 else "s"}</span> &middot; <span>{n0(ab["min"])}&ndash;{n0(ab["max"])}</span>',
                    [n0(ab['total']), n1(ab['total'] / dur), f'{ab["total"] / a["total"] * 100:.1f}%', str(ab['hits']), f'{ab["crits"] / max(ab["hits"], 1) * 100:.0f}%'],
                    Meter(ab['total'] / top), primary_phone=f'{n1(ab["total"] / dur)} DPS') for ab in a['abilities']]
    table = StatTable([('', False), ('Ability', False), ('', False), ('Amount', True), ('DPS', True), ('Share', True), ('Hits', True), ('Crit', True)], rows, 5,
                      sort='Amount', sortable=('Amount', 'DPS', 'Share', 'Hits', 'Crit'))
    plate = Section(f'{split_name(r["name"])[0]}&rsquo;s damage', Plate(table, 'Abilities'), meta=f'{n0(a["total"])} damage &middot; {n1(a["total"] / dur)} DPS')
    comb = next((c for c in s['combatants'] if c['guid'] == guid), None)
    doll = ''
    if comb:
        slots_l = ['Head', 'Neck', 'Shoulder', 'Back', 'Chest', 'Wrist', 'Main hand', 'Off hand', 'Ranged']
        slots_r = ['Hands', 'Waist', 'Legs', 'Feet', 'Ring', 'Ring', 'Trinket', 'Trinket']
        order = ['Head', 'Neck', 'Shoulder']   # the log's gear list is in the game's slot order; the fixture logs three
        logged = {order[i]: g for i, g in enumerate(comb['gear'][:3])}
        def cell(lbl):
            g = logged.get(lbl)
            if g is None or not g['ID']:
                return {'label': lbl, 'flag': None, 'state': 'empty', 'name': 'Nothing logged', 'line2': 'The log has no item in this slot.'}
            return {'label': lbl, 'flag': None, 'state': 'ok', 'name': f'Item {g["ID"]}', 'quality': 'poor', 'icon': None, 'stats': '',
                    'line2': f'Item level {g["ItemLevel"]} &middot; Not in our item data yet'.replace('&middot;', '·'), 'gain': None, 'basis': None}
        doll = Section('Gear at the pull', Paperdoll('Gear at the pull', 'slot', ([cell(x) for x in slots_l], [cell(x) for x in slots_r])),
                       meta=f'{escape(comb["spec"])} &middot; item level {comb["item_level"]}', action=Link('Open in the planner &rsaquo;'))
    rail = this_pull(R, fi) + Section('Targets', Rows([Row(escape(t['name']), n0(t['total']), link=False) for t in a['targets']], open_=True))
    return plate + chart(R, fi, 'damage', 'All friendlies') + doll, rail, []


def tab_threat(R, fi, source):
    s = R.summary(fi)
    ros = roster(s)
    dur = s['duration_ms'] / 1000
    tt = [t for t in s['threat_by_target'] if t['guid'] in ros]
    enemy = tt[0]['target_name'] if tt else ''
    series = [(split_name(t['name'])[0], colour(ros.get(t['guid'])), t['series'], f'<span class="num">{n0(t["threat"])}</span>') for t in sorted(tt, key=lambda t: -t['threat'])]
    marks = [(x['at_ms'] / 1000, '#e5b955') for x in s['taunts']] + [(d['at_ms'] / 1000, '#f0736b') for d in s['deaths']]
    phases = [(p['name'], p['start_ms'] / 1000, p['end_ms'] / 1000) for p in s.get('phases', [])]
    cav = f'<p class="t-small">{escape(FACTS_CHART["threatCaveat"])} <a class="link" href="#">{FACTS_CHART["threatCaveatLink"]}</a></p>'
    ch = Section(f'Estimated threat on {escape(enemy)}', LineChart(series, dur, 180, phases, marks, label='Estimated threat', step_ms=1000) + cav, meta='Whole fight')
    top = max((t['threat'] for t in tt), default=1)
    tot = sum(t['threat'] for t in tt) or 1
    rows = [StatRow(crest(ros.get(t['guid']), 28), pname(ros.get(t['guid'])), player_sub(ros.get(t['guid'], {})) or '&nbsp;',
                    [n0(t['threat']), f'{t["threat"] / tot * 100:.1f}%', f'{t["threat"] / top * 100:.0f}%'], MeterClass(t['threat'] / top, colour(ros.get(t['guid']))),
                    primary_phone=f'{t["threat"] / tot * 100:.1f}% of the threat') for t in sorted(tt, key=lambda t: -t['threat'])]
    plate = Section('Estimated threat', Plate(StatTable([('', False), ('Player', False), ('', False), ('Threat', True), ('Share', True), ('Of the top', True)], rows, 3,
                                                        sort='Threat', sortable=('Threat', 'Share')), 'Estimated threat'), meta=f'On {escape(enemy)} at the end of the fight')
    trows = [Row(f'{pname(ros.get(x["source_guid"]), x["source_name"])} taunted {escape(x["target_name"])}', escape(x['spell_name']),
                 lead=f'<span class="num muted">{at(x["at_ms"])}</span>{sicon(x["spell_id"], 24)}', link=False) for x in s['taunts']]
    rail = this_pull(R, fi) + Section('Taunts', Rows(trows, open_=True) if trows else '<p class="empty">Nobody taunted.</p>')
    return plate + ch, rail, []


def tab_auras(R, fi, kind):
    s = R.summary(fi)
    ros = roster(s)
    want = 'BUFF' if kind == 'Buffs' else 'DEBUFF'
    aus = sorted([a for a in s['auras'] if a['type'] == want], key=lambda a: -a['uptime_ms'])
    dur = s['duration_ms']
    rows = []
    for a in aus:
        on = ros.get(a['target_guid'])
        by = ', '.join(split_name((ros.get(g) or {'name': g}).get('name', g))[0] if g in ros else 'an enemy' for g in a['appliers'])
        sub = f'on {split_name(a["target_name"])[0]} &middot; from {escape(by)}' + (' &middot; up at the pull' if a.get('seeded') else '')
        rows.append(StatRow(sicon(a['spell_id']), f'<b>{escape(spell(a["spell_id"], a["name"]))}</b>', sub,
                            [f'{a["uptime_ms"] / dur * 100:.1f}%', str(a['applications'])], Meter(a['uptime_ms'] / dur, 'muted' if kind == 'Debuffs' else ''),
                            primary_phone=f'up {a["uptime_ms"] / dur * 100:.1f}%'))
    plate = Section(kind, Plate(StatTable([('', False), ('Aura', False), ('Uptime', False), ('Up', True), ('Applied', True)], rows, 2, sort='Up', sortable=('Up', 'Applied')), kind)
                    if rows else Plate(f'<p class="empty">No {kind.lower()} in this window.</p>', kind), meta=f'{len(rows)} on players &middot; uptime over the whole fight')
    return plate, this_pull(R, fi), []


def tab_deaths(R, fi):
    s = R.summary(fi)
    ros = roster(s)
    blocks = []
    for d in s['deaths']:
        r = ros.get(d['guid'])
        kb = d['killing_blow']
        head_row = Row(pname(r, d['name']), f'Killed by {escape(kb["source_name"])} &middot; {escape(kb["spell_name"])} &middot; {n0(kb["amount"])} ({n0(kb["overkill"])} overkill) &middot; {n0(kb["max_hp"])} health',
                       lead=f'<span class="num muted">{at(d["at_ms"])}</span>{crest(r, 32)}', trail='<span class="link">The 20 s before it &rsaquo;</span>', link=False)
        hits = [StatRow(sicon(h['spell_id'], 24), f'<b>{escape(h["spell_name"])}</b>', f'<span>{at(h["at_ms"])}</span> &middot; from {escape(h["source_name"])}',
                        [n0(h['amount']), f'{h.get("hp_after", 0) / h["max_hp"] * 100:.0f}%' if 'hp_after' in h else 'dead'],
                        Meter(h.get('hp_after', 0) / h['max_hp'], 'kill'), primary_phone=f'{n0(h["amount"])}') for h in d['last']]
        held = ', '.join(escape(x['name']) for x in d['auras_held'])
        body = StatTable([('', False), ('Last hits', False), ('Health after', False), ('Amount', True), ('Health', True)], hits, 2, sort='', sortable=())
        foot = f'<p class="t-small dfoot">Holding: {held}. No healing landed in the last 10 seconds.</p>' if held else ''
        blocks.append(f'<div class="dblock">{Rows([head_row])}{body}{foot}</div>')
    plate = Section('Deaths', Plate(''.join(blocks), 'Deaths') if blocks else Plate('<p class="empty">Nobody died.</p>', 'Deaths'), meta=f'{len(blocks)} in this fight')
    return plate + chart(R, fi, 'damage', 'All friendlies'), this_pull(R, fi), []


def tab_kicks(R, fi, which):
    s = R.summary(fi)
    ros = roster(s)
    key = 'interrupts' if which == 'Interrupts' else 'dispels'
    verb = 'Interrupted' if key == 'interrupts' else 'Removed'
    rows = [StatRow(sicon(x['spell_id']), f'{pname(ros.get(x["source_guid"]), x["source_name"])} &middot; {escape(spell(x["spell_id"], x["spell_name"]))}',
                    f'{verb} {escape(x["extra_spell_name"])} on {escape(split_name(x["target_name"])[0])}', [str(x['count'])], primary_phone=f'{x["count"]}&times;', bar=False)
            for x in s[key]]
    empty = 'Nothing interrupted.' if key == 'interrupts' else 'Nothing dispelled.'
    plate = Section(which, Plate(StatTable([('', False), ('Who, with what', False), ('Count', True)], rows, 1, bar=False) if rows else f'<p class="empty">{empty}</p>', which),
                    meta=f'{sum(x["count"] for x in s[key])} in this fight')
    return plate, this_pull(R, fi), []


def tab_resources(R, fi):
    s = R.summary(fi)
    ros = roster(s)
    rows = []
    for x in s['resources']:
        if x['guid'] not in ros:
            continue
        r = ros[x['guid']]
        rows.append(StatRow(crest(r, 28), pname(r), POWER.get(x['power_type'], 'Power'), [n0(x['gained']), n0(x['spent']), n0(x['wasted']), mmss(x['zero_ms'])],
                            Spark(x['series'], colour(r), total=x['max'] or None), primary_phone=f'{n0(x["gained"])} gained'))
    plate = Section('Resources', Plate(StatTable([('', False), ('Player', False), ('Over the fight', False), ('Gained', True), ('Spent', True), ('Wasted', True), ('Time empty', True)], rows, 4), 'Resources'),
                    meta='Mana, rage, energy and focus, whole fight')
    return plate, this_pull(R, fi), []


def tab_casts(R, fi):
    s = R.summary(fi)
    ros = roster(s)
    dur = s['duration_ms']
    rows = []
    for c in sorted(s['casts'], key=lambda c: -(c['succeeded'] + c['started'])):
        r = ros.get(c['guid']) or ros.get(c['owner_guid'])
        who = split_name(c['name'])[0] + (f' ({split_name(ros[c["owner_guid"]]["name"])[0]}&rsquo;s pet)' if c['guid'] != c['owner_guid'] and c['owner_guid'] in ros else '')
        n = c['succeeded'] or c['started']
        rows.append(StatRow(sicon(c['spell_id']), f'<b>{escape(spell(c["spell_id"], c["spell_name"]))}</b>', f'<span style="color:{colour(r)}">{who}</span>',
                            [str(n), f'{n / (dur / 60000):.1f}', str(c['failed'])], Spark([], colour(r) if r else '#9a9484', marks=[t / dur for t in c['sequence']]),
                            primary_phone=f'{n} cast{"s" if n != 1 else ""}'))
    plate = Section('Casts', Plate(StatTable([('', False), ('Spell', False), ('When', False), ('Casts', True), ('Per min', True), ('Failed', True)], rows, 3,
                                             sort='Casts', sortable=('Casts', 'Per min', 'Failed')), 'Casts'), meta=f'Every caster, {mmss(dur)}')
    return plate, this_pull(R, fi), []


def tab_rating(R, fi):
    p = RATING['players'][0]
    ros = {split_name(r['name'])[0]: r for r in R.summary(fi)['roster']}
    r = {**(ros.get(p['player_name']) or {'name': p['player_name']}), 'class': p['class']}   # the rating carries the class the log lacks
    rc, fx = COPY['ratingCopy'], FACTS['fns']['rating']
    basis_line = escape(fx['percentileBasisHolyPriest'].replace('Priestss', 'Priests'))   # shipped copy bug B1 (spec 9), mocked as fixed
    headrow = Row(f'<span style="color:{colour(r)}">{escape(fx["headingElyra"])}</span>', basis_line, lead=crest(r, 40),
                  trail=Figure(str(p['overall']), unit=rc['outOf100'], tone='gold', size='big'), link=False)
    rows = []
    for c in p['components']:
        lab = fx['labels'][c['name']]
        if c['excluded']:
            rows.append(StatRow('', f'<b>{lab}</b>', escape(fx['excludedMechanics']), ['&ndash;', '&ndash;', '&ndash;'], '', bar=True))
            continue
        basis = f'Percentile of {c["bracket_n"]}' if c['basis'] == 'percentile' else 'Absolute'
        rows.append(StatRow('', f'<b>{lab}</b>', escape(fx['definitions'][c['name']].split(': ', 1)[1]), [str(c['score']), f'{c["weight"]}%', basis],
                            Meter(c['score'] / 100, 'kill' if c['score'] >= 75 else ''), primary_phone=f'{c["score"]} of 100'))
    table = StatTable([('', False), ('Part', False), ('Score', False), ('Of 100', True), ('Weight', True), ('Basis', True)], rows, 3)
    plate = Section(escape(fx['headingElyra']), Plate(Rows([headrow]) + table, 'Rating'), meta=f'{p["spec"]} {p["class"]} &middot; {RATING["kill_time_band"]} kill time',
                    action=Link(escape(rc['explainLink']) + ' &rsaquo;'))
    rail = this_pull(R, fi) + rating_rail(R, fi) + Section('The cap', f'<p class="t-small rail-p">{escape(rc["cappedRule"])}</p>')
    return plate, rail, []


def tab_coach(R, fi):
    rv = COACH['review']
    ros = roster(R.summary(fi))
    me = ros[VIEWER]
    anchors = COACH['anchors']
    def anc(ids):
        out = []
        for i in ids:
            a = anchors[i]
            if a['kind'] == 'time':
                out.append(Tag(f'{mmss(a["at_ms"])}&ndash;{mmss(a["to_ms"])}', 'muted'))
            elif a['kind'] == 'cast':
                out.append(f'<span class="anc">{sicon(a["spell_id"], 20)}{escape(a["spell_name"])}</span>')
            elif a['kind'] == 'resource':
                out.append(Tag(f'Rage {mmss(a["at_ms"])}&ndash;{mmss(a["to_ms"])}', 'muted'))
        return '<div class="ancs">' + ''.join(out) + '</div>' if out else ''
    ff = rv['fixFirst']
    body = (f'<div class="coach"><p class="lead">{escape(rv["opening"])}</p>'
            f'<div class="cb"><span class="t-label gold">Fix first</span><h3 class="t-h3">{escape(ff["title"])}</h3><p>{escape(ff["why"])}</p>{anc(ff["anchors"])}'
            f'<p class="drill"><span class="t-label">The drill</span> {escape(ff["drillText"])}</p></div>'
            + ''.join(f'<div class="cb"><span class="t-label">Then fix</span><h3 class="t-h3">{escape(t["title"])}</h3><p>{escape(t["why"])}</p>{anc(t["anchors"])}</div>' for t in rv['thenFix'])
            + '<div class="cb"><span class="t-label">Doing well</span><ul class="ok">' + ''.join(f'<li><span class="ck ok">&#10003;</span>{escape(d["text"])}</li>' for d in rv['doingWell']) + '</ul></div></div>')
    plate = Section(escape(COACH['persona']['label']), Plate(body, 'Coach review'), meta=f'{pname(me)} &middot; {escape(me["spec"])} {escape(me["class"])} &middot; reviewed 25 Oct')
    rail = (this_pull(R, fi) + Section('Measured against', '<ul class="rail-l">' + ''.join(f'<li>{escape(b)}</li>' for b in COACH['basis_labels']) + '</ul>')
            + Section('What the coach could not check', '<ul class="rail-l">' + ''.join(f'<li>{escape(c)}</li>' for c in rv['couldntCheck']) + '</ul>'))
    return plate, rail, []


def tab_mechanics(R, fi):
    s = R.summary(fi)
    ros = roster(s)
    rows = []
    kinds = {'avoidable': 'Avoidable', 'interrupt': 'Interrupt', 'dispel': 'Dispel', 'unavoidable': 'Unavoidable'}
    for m in s['mechanics']['rows']:
        if m['kind'] == 'avoidable' and m.get('players'):
            who = ', '.join(f'{pname(ros.get(p["guid"]))} {p["hits"]} hits{" (died)" if p.get("killed") else ""}' for p in m['players'])
            figs, st = [str(sum(p['hits'] for p in m['players'])), n0(sum(p['damage'] for p in m['players']))], 'wipe'
        elif m['kind'] == 'interrupt':
            done = sum(x['count'] for x in s['interrupts'] if x['extra_spell_id'] == m['spell_id'])
            who, figs, st = f'{m["casts"]} cast, {done} interrupted', [str(m['casts'] - done), '&ndash;'], 'wipe' if m['casts'] - done else 'kill'
        elif m['kind'] == 'dispel':
            done = sum(x['count'] for x in s['dispels'] if x['extra_spell_id'] == m['spell_id'])
            who, figs, st = f'{m["applied"]} applied, {done} dispelled', [str(m['applied'] - done), '&ndash;'], 'wipe' if m['applied'] - done else 'kill'
        else:
            who, figs, st = 'Nothing the table lists went wrong.' if m['kind'] == 'avoidable' else 'Lands on whoever holds the boss.', ['0', '&ndash;'], 'kill'
        rows.append(StatRow(sicon(m['spell_id']), f'<b>{escape(m["name"])}</b> {Tag(kinds[m["kind"]], "wipe" if st == "wipe" else "muted")}',
                            f'{who} &middot; {escape(m["note"])}', figs, primary_phone=figs[0] + ' failed', bar=False))
    plate = Section('Mechanics', Plate(StatTable([('', False), ('Mechanic', False), ('Failed', True), ('Damage', True)], rows, 2, bar=False), 'Mechanics'),
                    meta='Problems first, then what went right')
    clean = [r for g, r in ros.items() if not any(p['guid'] == g for m in s['mechanics']['rows'] for p in m.get('players', []))]
    rail = this_pull(R, fi) + Section('Clean', Rows([Row(pname(r), 'No avoidable damage', lead=crest(r, 28), link=False) for r in clean], open_=True))
    return plate, rail, []


def events(R, fi):
    import duckdb
    p = R.base / f'fights/{fi}/events.parquet'
    q = (f"select (time_unix_nano - (select min(time_unix_nano) from '{p}'))/1e6 as ms, kind, source_name, dest_name, spell_id, spell_name, amount "
         f"from '{p}' where kind not in ('combatant_info','encounter_start','encounter_end') order by 1")
    return duckdb.sql(q).fetchall()


def ev_sentence(ros_by_name, k, src, dst, sid, sp, amt):
    n = lambda x: f'<span style="color:{colour(ros_by_name.get(x))}">{escape(split_name(x)[0])}</span>' if x else ''
    spn = escape(spell(sid, sp)) if sp else 'melee'
    return {
        'cast_success': f'{n(src)} casts {spn}', 'cast_start': f'{n(src)} begins casting {spn}', 'damage': f'{n(src)} hits {n(dst)} with {spn}',
        'damage_landed': f'{n(src)} hits {n(dst)} in melee', 'heal': f'{n(src)} heals {n(dst)} with {spn}', 'death': f'{n(dst)} dies',
        'aura_applied': f'{spn} on {n(dst)} (from {n(src)})', 'aura_removed': f'{spn} fades from {n(dst)}', 'aura_refresh': f'{spn} refreshed on {n(dst)}',
        'energize': f'{n(src)} gains power from {spn}', 'instakill': f'{n(src)} uses {spn} on {n(dst)}', 'enchant': f'{n(src)} applies {spn}',
    }.get(k, f'{k} {n(src)} {n(dst)}')


def tab_events(R, fi):
    s = R.summary(fi)
    by_name = {r['name']: r for r in s['roster']}
    evs = events(R, fi)
    rows = [StatRow(sicon(sid, 20) if sid else '', f'<span class="ev">{ev_sentence(by_name, k, src, dst, sid, sp, amt)}</span>', '',
                    [at(ms), n0(amt) if amt else '&ndash;'], primary_phone=at(ms), bar=False) for ms, k, src, dst, sid, sp, amt in evs]
    plate = Section('Events', Plate(StatTable([('', False), ('Event', False), ('Time', True), ('Amount', True)], rows, 2, bar=False), 'Events'),
                    meta=f'{len(evs)} events in this window, in order')
    rail = this_pull(R, fi) + Section('Show', '<div class="chips">' + ''.join(f'<label class="chip"><input type="checkbox" checked> {t}</label>' for t in ('Casts', 'Damage', 'Heals', 'Auras', 'Deaths', 'Power')) + '</div>')
    return plate, rail, []


QUERY = ("SELECT source_name, spell_name, count(*) AS hits, sum(amount) AS total\nFROM read_parquet('events.parquet')\n"
         "WHERE kind = 'damage' AND (time_unix_nano - (SELECT min(time_unix_nano) FROM read_parquet('events.parquet'))) / 1000000 BETWEEN 0 AND {end}\n"
         "GROUP BY 1, 2\nORDER BY total DESC")


def tab_queries(R, fi):
    import duckdb
    s = R.summary(fi)
    p = R.base / f'fights/{fi}/events.parquet'
    sql = QUERY.format(end=s['duration_ms'])
    res = duckdb.sql(sql.replace("read_parquet('events.parquet')", f"read_parquet('{p}')")).fetchall()
    cols = ['source_name', 'spell_name', 'hits', 'total']
    rows = [StatRow('', f'<b>{escape(split_name(a)[0])}</b>', '', [escape(b), str(c), n0(d)], primary_phone=n0(d), bar=False) for a, b, c, d in res]
    starts = ''.join(f'<a href="#" class="{"on" if i == 0 else ""}" role="radio" aria-checked="{"true" if i == 0 else "false"}">{t}</a>' for i, t in enumerate(
        ['Damage by ability, exactly, in this window', 'Pet damage, by pet', 'Every miss, by type', 'Healing and overhealing by spell', 'Healing by target, with overhealing']))
    body = (f'<div class="qwrap"><span class="t-label">Starting points</span><div class="seg qstarts" role="radiogroup" aria-label="Starting points">{starts}</div>'
            f'<label class="t-label" for="q">Query</label><pre class="code" id="q">{escape(sql)}</pre><div class="qrun">{Button("Run", primary=True)}<span class="t-small">{len(res)} rows</span></div></div>'
            + StatTable([('', False), (cols[0], False), (cols[1], True), (cols[2], True), (cols[3], True)], rows, 3, bar=False))
    plate = Section('Queries', Plate(body, 'Query'), meta='SQL over this fight&rsquo;s events, run in your browser')
    rail = this_pull(R, fi) + Section('Columns', '<p class="t-small rail-p">time_unix_nano, kind, event, source_name, dest_name, spell_id, spell_name, amount, overkill, overheal, absorbed, critical, miss_type, aura_type, power_type and 50 more.</p>')
    return plate, rail, []


def tab_timeline(R, fi):
    s = R.summary(fi)
    ros = roster(s)
    evs = events(R, fi)
    dur = s['duration_ms']
    W, LH = 1000, 34
    actors = [r for r in s['roster']] + [{'guid': 'boss', 'name': R.fight(fi)['name']}]
    lanes, labels = '', ''
    for i, r in enumerate(actors):
        y = i * LH
        lanes += f'<rect x="0" y="{y + 15}" width="{W}" height="4" fill="#1c2230"/>'
        nm = r['name']
        for a in [a for a in s['auras'] if a['target_guid'] == r['guid']]:
            for sg in a['segments']:
                lanes += f'<rect x="{sg["start_ms"] / dur * W:.1f}" y="{y + 22}" width="{(sg["end_ms"] - sg["start_ms"]) / dur * W:.1f}" height="5" rx="2" fill="{"#7fd48a99" if a["type"] == "BUFF" else "#f0736b99"}"/>'
        for ms, k, src, dst, sid, sp, amt in evs:
            if k == 'cast_success' and src == nm:
                lanes += f'<rect x="{ms / dur * W - 2:.1f}" y="{y + 6}" width="4" height="14" rx="1" fill="#e5b955"/>'
            elif k in ('damage', 'damage_landed') and src == nm:
                lanes += f'<rect x="{ms / dur * W - 1.5:.1f}" y="{y + 9}" width="3" height="8" fill="{colour(ros.get(r["guid"]))}"/>'
            elif k == 'death' and dst == nm:
                lanes += f'<rect x="{ms / dur * W - 1.5:.1f}" y="{y + 2}" width="3" height="28" fill="#f0736b"/>'
        labels += f'<div class="tl-l">{crest(r, 20) if r["guid"] != "boss" else EmptyCrest(20)}<span style="color:{colour(r) if r["guid"] != "boss" else "var(--text)"}">{escape(split_name(nm)[0])}</span></div>'
    svg = f'<svg viewBox="0 0 {W} {len(actors) * LH}" preserveAspectRatio="none" style="height:{len(actors) * LH}px" role="img" aria-label="Every player&rsquo;s casts, hits, auras and death on one time axis">{lanes}</svg>'
    ticks = ''.join(f'<span>{mmss(dur * q)}</span>' for q in (0, .25, .5, .75, 1))
    legend = LanesLegend([('#e5b955', 'Cast'), ('#3fc7eb', 'Hit, in the class colour'), ('#7fd48a99', 'Buff up'), ('#f0736b99', 'Debuff up'), ('#f0736b', 'Death')])
    body = f'{legend}<div class="tlg"><div class="tl-ls">{labels}</div><div class="tl-p">{svg}<div class="ax">{ticks}</div></div></div>'
    plate = Section('Timeline', Plate(Lanes(body), 'Timeline'), meta=f'Every player, {mmss(dur)} &middot; hover for the event')
    return plate, this_pull(R, fi), []


def tab_compare(R, fi, other):
    a, b = R.summary(fi), R.summary(other)
    ra, rb = roster(a), roster(b)
    rows = []
    for g, r in sorted(ra.items(), key=lambda kv: -kv[1]['dps']):
        o = rb.get(g)
        diff = (r['dps'] - o['dps']) if o else None
        dtxt = '&ndash;' if diff is None else f'<span class="{"kill" if diff > 0 else "wipe" if diff < 0 else ""}">{"+" if diff > 0 else ("&minus;" if diff < 0 else "")}{n1(abs(diff))}</span>'
        rows.append(StatRow(crest(r, 28), pname(r), player_sub(r) or '&nbsp;', [n1(r['dps']), n1(o['dps']) if o else '&ndash;', dtxt], primary_phone=f'{n1(r["dps"])} vs {n1(o["dps"]) if o else "&ndash;"}', bar=False))
    fa, fb = R.fight(fi), R.fight(other)
    plate = Section('Compare', Plate(StatTable([('', False), ('Player', False), ('This pull', True), (escape(fb['name']), True), ('Difference', True)], rows, 3, bar=False), 'Compare'),
                    meta=f'DPS, {escape(fa["name"])} {mmss(fa["duration_ms"])} against {escape(fb["name"])} {mmss(fb["duration_ms"])}')
    return plate, this_pull(R, fi), ['Players missing from the other pull show a dash.']


# ------------------------------------------------------------------------------------------------ pages
FACTS_CHART = {'threatCaveat': 'An estimate: 1 threat per point of damage and 0.5 per point of healing, split across the enemies in combat. No stance, form, talent or taunt modifiers yet.',
               'threatCaveatLink': 'What it leaves out'}   # lib/chart/copy.ts, verbatim


PAGE_CSS = '''<style>
.ans-act{display:flex;gap:8px;margin-top:6px}@media(min-width:900px){.ans-act{justify-content:flex-end}}
.head .ans{min-width:0}@media(min-width:900px){.head .ans{min-width:360px}}
.dblock+.dblock{border-top:1px solid var(--line)}.dblock .st-h{padding-top:6px}.dfoot{padding:8px 16px 12px}
.coach{padding:18px 20px;display:flex;flex-direction:column;gap:18px;max-width:880px}.coach .lead{font-size:17px;line-height:26px;color:var(--strong)}
.coach .cb{display:flex;flex-direction:column;gap:6px}.coach .cb p{font-size:15px;line-height:24px}.coach .drill{color:var(--text)}.coach .t-label.gold{color:var(--gold)}
.ancs{display:flex;flex-wrap:wrap;gap:8px;align-items:center}.anc{display:inline-flex;align-items:center;gap:6px;font-size:13px;color:var(--text)}
.coach ul.ok{display:flex;flex-direction:column;gap:6px}.coach ul.ok li{display:flex;gap:10px;align-items:center}
.rail-l{display:flex;flex-direction:column;gap:8px;padding-top:10px;font-size:14px;line-height:20px}.rail-p{padding-top:10px}
.qwrap{display:flex;flex-direction:column;gap:8px;padding:14px 16px;border-bottom:1px solid var(--line-soft)}
.qstarts{flex-wrap:wrap}.qstarts a{height:32px;font-size:13px}
pre.code{margin:0;padding:10px 12px;white-space:pre-wrap;font:13px/20px ui-monospace,Menlo,monospace;color:var(--text);border:1px solid var(--line);border-radius:4px}
.qrun{display:flex;gap:12px;align-items:center}
.chips{display:flex;flex-wrap:wrap;gap:8px 16px;padding-top:10px;font-size:14px}
.tlg{display:grid;grid-template-columns:160px minmax(0,1fr);padding:4px 16px 12px}.tl-ls{display:flex;flex-direction:column}.tl-l{height:34px;display:flex;align-items:center;gap:8px;font:600 14px/16px var(--font-body);white-space:nowrap;overflow:hidden}
.tl-p svg{display:block;width:100%}.tl-p .ax{display:flex;justify-content:space-between;font:600 12px/16px var(--font-body);color:var(--muted);padding-top:4px}
.lanes .lg{padding-inline:16px}
@media(max-width:767px){.tlg{grid-template-columns:96px minmax(560px,1fr);overflow-x:auto}}
.ev{font-size:14px}
.menu-at{position:relative}.menu-at .menu{top:4px;left:0}
.statesheet{display:grid;gap:40px}
</style>'''


def body_for(R, fi, tab, source, compare_with=4):
    return {
        'Summary': lambda: tab_summary(R, fi, source),
        'Damage done': lambda: tab_actor(R, fi, source, 'Damage done'),
        'Healing': lambda: tab_actor(R, fi, source, 'Healing'),
        'Damage taken': lambda: tab_actor(R, fi, source, 'Damage taken'),
        'Threat': lambda: tab_threat(R, fi, source),
        'Buffs': lambda: tab_auras(R, fi, 'Buffs'),
        'Debuffs': lambda: tab_auras(R, fi, 'Debuffs'),
        'Deaths': lambda: tab_deaths(R, fi),
        'Interrupts': lambda: tab_kicks(R, fi, 'Interrupts'),
        'Dispels': lambda: tab_kicks(R, fi, 'Dispels'),
        'Resources': lambda: tab_resources(R, fi),
        'Casts': lambda: tab_casts(R, fi),
        'Rating': lambda: tab_rating(R, fi),
        'Coach': lambda: tab_coach(R, fi),
        'Mechanics': lambda: tab_mechanics(R, fi),
        'Timeline': lambda: tab_timeline(R, fi),
        'Events': lambda: tab_events(R, fi),
        'Queries': lambda: tab_queries(R, fi),
        'Compare': lambda: tab_compare(R, fi, compare_with),
    }[tab]()


def report_page(tab: str, fi: int = 3, source: str = 'All friendlies', viewer: bool = False, sample: bool = False, onyxia: bool = False,
                open_m: bool = False, menu: bool = False, player: str = '') -> str:
    R = Report(SAMPLE_DIR if sample else FIX, onyxia)
    if player:
        main, rail, nts = tab_player(R, fi, player)
        source = split_name(roster(R.summary(fi))[player]['name'])[0]
    else:
        main, rail, nts = body_for(R, fi, tab, source)
    compare = fight_label(R.fight(4)) if tab == 'Compare' else ''
    menu_html = ''
    if menu:
        fights = R.meta['fights']
        bosses = [(fight_label(f), f'{clock(f["start"])} &middot; {f.get("deaths", 0)} death{"" if f.get("deaths", 0) == 1 else "s"}', '', '', f['index'] == fi) for f in fights if f['kind'] == 'encounter']
        trash = [(fight_label(f), f'{clock(f["start"])} &middot; {f.get("mob_deaths", 0)} mobs died', '', '', f['index'] == fi) for f in fights if f['kind'] == 'trash']
        menu_html = '<div class="pc menu-at">' + Menu('Fights', [('Boss pulls', [('All boss pulls', f'{len(bosses)} pulls', '', '', False)] + bosses), ('Trash', trash)]) + '</div>'
    who = ('Baelgrim', 'warrior', 'Fury Warrior &middot; 60') if viewer else None
    body = (head(R, fi, tab, source, viewer, sample) + bar(R, fi, tab, source, open_m, compare) + menu_html + notices(nts)
            + Page('split', main, rail) + PAGE_CSS)
    note = (f'Report mock (design/specs/2026-11-07-logs-reports.md), built only from ws_prims. Data: '
            + ('web/public/sample-report/sample' if sample else 'web/src/fixtures/report (fixture2abcd)') + f', fight {fi}'
            + ('; the fight re-read as the Onyxia pull as coach-stubs.ts serves it; SAMPLE viewer: the owner of Baelgrim' if onyxia else '') + f'; build {BUILD} names and icons.')
    return build_page(f'{R.title}: {tab}', body, 'Logs', who, note)


def sample_page() -> str:
    R = Report(SAMPLE_DIR)
    fi = 6   # the sample's default pull as the live page opens it (Trash 12:43 at 21:07:51)
    main, rail, nts = tab_summary(R, fi, 'Your group')
    body = head(R, fi, 'Summary', 'Your group', sample=True) + bar(R, fi, 'Summary', 'Your group') + notices([]) + Page('split', main, rail) + PAGE_CSS
    return build_page(R.title, body, 'Logs', None, 'Sample report mock; data web/public/sample-report/sample, fight 6, source Your group (meta.group).')


def skeleton_rows(n: int) -> str:
    return ''.join(f'<div class="st-r stc" style="--n:4">{Skeleton(28, 28, True)}<span>{Skeleton(180, 16)}</span><span></span>{Skeleton(48, 16)}{Skeleton(48, 16)}{Skeleton(48, 16)}{Skeleton(48, 16)}</div>' for _ in range(n))


def state_page(state: str) -> str:
    R = Report(FIX)
    rc = COPY['reportCopy']
    if state == 'loading':
        head_html = PageHeader(Skeleton(520, 40), meta=[Skeleton(320, 16)], answer=f'{Skeleton(200, 16)}{Skeleton(160, 40)}{Skeleton(220, 16)}',
                               tabs=TabStripMore(['Summary'], '', 1, 1).replace('>Summary<', ' aria-disabled="true">Summary<'))
        main = Section('Players', Plate(f'<div class="st" style="--n:4">{skeleton_rows(6)}</div>', 'Players'))
        rail = Section('This pull', ''.join(f'<p>{Skeleton(300, 16)}</p>' for _ in range(5)))
        body = head_html + bar(R, 3, 'Summary', 'All friendlies') + Page('split', main, rail)
    elif state in ('error', 'missing'):
        msg = 'Report data did not load' if state == 'error' else 'No report with that id'
        act = Button('Try again', small=True) if state == 'error' else Link('Open your reports &rsaquo;')
        body = PageHeader('Report') + Page('ledger', Plate(f'<p class="empty" role="alert">{msg}. {act}</p>', 'Report'))
    elif state == 'empty':
        return report_page('Damage done', fi=2)
    elif state in ('processing', 'waiting', 'failed'):
        line = {'processing': rc['parsing'], 'waiting': rc['waiting'], 'failed': f'{rc["parseFailed"]} <a class="link" href="#">{rc["parseFailedLink"]}</a>'}[state]
        title = rc['yourUpload'] if state != 'waiting' else R.title
        meta = ['Uploaded 2 minutes ago'] if state != 'waiting' else [Tag('Live', 'kill'), 'Logging now']
        body = PageHeader(escape(title), meta=meta, tabs=TabStripMore(['Summary'], '', 1, 1).replace('>Summary<', ' aria-disabled="true">Summary<')) + Page('ledger', Plate(f'<p class="empty" role="status">{line}</p>', 'Fights'))
    return build_page(f'Report: {state}', body + PAGE_CSS, 'Logs', None, f'Report state mock: {state}.')


PAGES = {
    'web-report-summary': lambda: report_page('Summary'),
    'web-report-damage-done': lambda: report_page('Damage done'),
    'web-report-healing': lambda: report_page('Healing'),
    'web-report-damage-taken': lambda: report_page('Damage taken'),
    'web-report-deaths': lambda: report_page('Deaths'),
    'web-report-mechanics': lambda: report_page('Mechanics'),
    'web-report-timeline': lambda: report_page('Timeline'),
    'web-report-threat': lambda: report_page('Threat'),
    'web-report-casts': lambda: report_page('Casts'),
    'web-report-buffs': lambda: report_page('Buffs'),
    'web-report-debuffs': lambda: report_page('Debuffs'),
    'web-report-rating': lambda: report_page('Rating'),
    'web-report-coach': lambda: report_page('Coach', viewer=True, onyxia=True),
    'web-report-interrupts': lambda: report_page('Interrupts', fi=1),
    'web-report-dispels': lambda: report_page('Dispels', fi=1),
    'web-report-resources': lambda: report_page('Resources'),
    'web-report-events': lambda: report_page('Events'),
    'web-report-queries': lambda: report_page('Queries'),
    'web-report-compare': lambda: report_page('Compare'),
    'web-report-player': lambda: report_page('Damage done', player='Player-4184-000000A1'),
    'web-report-fights': lambda: report_page('Summary', menu=True),
    'web-report-bar': lambda: report_page('Damage done', open_m=True),
    'web-report-loading': lambda: state_page('loading'),
    'web-report-error': lambda: state_page('error'),
    'web-report-missing': lambda: state_page('missing'),
    'web-report-empty': lambda: state_page('empty'),
    'web-report-processing': lambda: state_page('processing'),
    'web-report-waiting': lambda: state_page('waiting'),
    'web-report-failed': lambda: state_page('failed'),
    'web-report-sample': sample_page,
}

if __name__ == '__main__':
    if sys.argv[1:] == ['--list']:
        print(' '.join(PAGES))
    else:
        only = sys.argv[1:]
        for name, fn in PAGES.items():
            if not only or name in only:
                print(write(f'{name}.html', fn()))
