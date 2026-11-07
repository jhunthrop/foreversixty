"""Web-layout mock: a report (/reports/<id>, every tab and state, and /reports/sample), composed ONLY from ws_prims.

Spec: design/specs/2026-11-07-logs-reports.md section 6. Every figure, name, time and series is read at render time from
REAL Forever data the build and its parity spec serve (decision D1, ruled 2026-11-07):
  - web/public/sample-report/sample: the 17 Sept 2026 beta night on Zephras Isle (client 1.60.1, 79 trash pulls,
    anonymised names; meta.json, fights/<n>/summary.json, fights/<n>/events.parquet, which the Events, Timeline and Queries
    tabs read through DuckDB exactly as the shipped page does in the browser);
  - api/internal/coach/copy.go (the coach's own line for a trash pull);
  - data/builds/<active build>/ spells.json and spellicons.json (spell names and icons, as the build resolves them).
The log has no boss pull, so the boss-only filled states (your DPS and parse in the head, Rating, the Coach review, a
Mechanics table, interrupts, dispels, taunts, gear at the pull) are not mocked: each tab shows what the real pull shows.
SAMPLE (typed, marked): the Coach tab's signed-in viewer owns Soren Everhart.

  python3 design/mocks/ws_report.py && for p in $(python3 design/mocks/ws_report.py --list); do node design/mocks/ws_shoot.mjs $p 390,1440,1920,2000; done
"""
import sys
from html import escape

from ws_report_lib import *  # noqa: F401,F403  (data, names, icons, head, bar, chart, this_pull)
from ws_report_lib import _Button  # noqa: F401
from ws_lib import REGION


# ------------------------------------------------------------------------------------------------ tabs
def player_sub(r: dict, extra: str = '') -> str:
    bits = [b for b in (f'{r.get("spec", "")} {r.get("class", "")}'.strip(), f'item level {r["item_level"]}' if r.get('item_level') else '', extra) if b]
    return ' &middot; '.join(escape(b) if '<' not in b else b for b in bits)


def plink(r: dict | None, name: str = '') -> str:
    """A player's name in a table row: the link that opens their view (Source = that player, spec 9 "StatTable row")."""
    return f'<a href="#" class="nm-l">{pname(r, name)}</a>'


# Copy CSV, the tables' one quiet text action in the Section head (coordinator ruling 2026-11-07: a restyle keeps
# every working feature; components/report/CopyCsv.svelte `inHead`).
COPY_CSV = '<button type="button" class="link">Copy CSV</button>'


def mitigated_line(s: dict, fr: set[str]) -> str:
    """Damage taken's Mitigated line (ActorTable.svelte): absorbed, blocked and the hits avoided by type, alphabetical."""
    absorbed = blocked = 0
    misses: dict[str, int] = {}
    for a in s['damage_taken']:
        if a['guid'] in fr and a['effective'] > 0:
            for ab in a['abilities']:
                absorbed += ab.get('absorbed', 0)
                blocked += ab.get('blocked', 0)
                for k, v in (ab.get('misses') or {}).items():
                    misses[k] = misses.get(k, 0) + v
    hits = sum(misses.values())
    if not (absorbed or blocked or hits):
        return ''
    by = ', '.join(f'{v} {k.lower()}' for k, v in sorted(misses.items()))
    parts = ([f'{n0(absorbed)} absorbed'] if absorbed else []) + ([f'{n0(blocked)} blocked'] if blocked else []) + ([f'{hits} hits avoided ({by})'] if hits else [])
    return 'Mitigated: ' + ' &middot; '.join(parts)


def melee(sid: int, logged: str) -> str:
    return 'Melee' if sid == 0 and not logged else spell(sid, logged)


# Summary (MUST 1 as ruled): the amount beside its rate, as Details! and WCL print it; a pair whose rate is 0.0 on every row
# is hidden; Parse is a boss pull's column (a trash pull has no parse).
SUMMARY_PAIRS = (('Damage', 'damage_done', 'DPS', 'dps'), ('Healing', 'healing_done', 'HPS', 'hps'), ('Taken', 'damage_taken', 'Taken / s', 'dtps'))


def tab_summary(R, fi, source):
    s = R.summary(fi)
    f = R.fight(fi)
    rs = sorted(in_scope(R, s, source), key=lambda r: (-r['damage_done'], -r['healing_done'], r['name']))
    fr = friendlies(R, s, source)
    pairs = [pr for pr in SUMMARY_PAIRS if any(n1(r[pr[3]]) != '0.0' for r in rs)]
    boss = f['kind'] != 'trash'
    heads = [('', False), ('Player', False)] + ([('Parse', True)] if boss else []) + [h for a, _, rt, _ in pairs for h in ((a, True), (rt, True))] + [('Active', True), ('Deaths', True)]
    rows = []
    for r in rs:
        figs = (['&ndash;'] if boss else []) + [x for _, ak, _, rk in pairs for x in (n0(r[ak]), rate(r[rk]))]
        figs += [f'{r["activity_pct"]:.1f}%', f'<span class="{"wipe" if r["deaths"] else ""}">{r["deaths"]}</span>']
        rows.append(StatRow(crest(r, 28), plink(r), player_sub(r) or '&nbsp;', figs,
                            primary_phone=f'damage &middot; {rate(r["dps"])} DPS', bar=False))
    table = StatTable(heads, rows, len(heads) - 2, bar=False, sort='Damage', sortable=tuple(h for h, _ in heads[2:]))
    nd = sum(1 for d in s['deaths'] if d['guid'] in fr)
    plate = Section('Players', Plate(table, 'Players'), meta=f'{plural(len(rs), "player")} &middot; {plural(nd, "death")} &middot; {mmss(s["duration_ms"])}')
    abil: dict[str, list] = {}
    for a in s['damage_taken']:
        if a['guid'] in fr:
            for ab in a['abilities']:
                abil.setdefault(ab['name'] or 'Melee', [ab.get('spell_id', 0), 0])[1] += ab.get('effective', ab['total'])
    # Share is of everything every player took in this window, whatever the source scope shows (the shipped caption)
    tot = sum(ab.get('effective', ab['total']) for a in s['damage_taken'] if a['guid'] in roster(s) for ab in a['abilities']) or 1
    dur = s['duration_ms'] / 1000
    trows = [StatRow(sicon(sid, 24), f'<b>{escape(melee(sid, k))}</b>', '', [f'{v / tot * 100:.1f}%', n0(v), rate(v / dur)], Meter(v / tot), primary_phone=f'{v / tot * 100:.1f}%')
             for k, (sid, v) in sorted(abil.items(), key=lambda kv: -kv[1][1])]
    taken = Section('Damage taken by ability', StatTable([('', False), ('Ability', False), ('', False), ('Share', True), ('Amount', True), ('Per sec', True)], trows, 3, open_=True)
                    + '<p class="t-small" style="padding-top:8px">Share is of everything every player took in this window, whatever the source scope shows.</p>'
                    if trows else '<p class="empty">Nobody took damage in this window.</p>', action=Link('Damage taken &rsaquo;'))
    main = plate + chart(R, fi, 'damage', source) + taken   # R4: the table first, the chart under it
    rail = this_pull(R, fi, source) + deaths_rail(R, fi, source) + rating_rail(R, fi)
    return main, rail, ([PARSES_MISSING] if boss else [])


def deaths_rail(R, fi, source):
    s = R.summary(fi)
    ros = roster(s)
    fr = friendlies(R, s, source)
    dd = [d for d in s['deaths'] if d['guid'] in fr]
    if not dd:
        return Section('Deaths', '<p class="empty">Nobody died.</p>')
    rows = [Row(pname(ros.get(d['guid']), d['name']), f'{escape(melee(d["killing_blow"]["spell_id"], d["killing_blow"]["spell_name"]))} &middot; {escape(d["killing_blow"]["source_name"])} &middot; {n0(d["killing_blow"]["amount"])}',
                lead=f'<span class="num muted">{at(d["at_ms"])}</span>{crest(ros.get(d["guid"]), 28)}') for d in dd]
    return Section('Deaths', Rows(rows, open_=True), action=Link('Deaths &rsaquo;'))


def rating_rail(R, fi):
    rc = COPY['ratingCopy']
    return Section(rc['panelHeading'], f'<p class="empty">{escape(rc["noRatingForTrash"] if R.fight(fi)["kind"] == "trash" else rc["panelEmpty"])}</p>')


def side_rows(pairs: dict[str, float], cap: int = 8) -> str:
    tot = sum(pairs.values()) or 1
    ranked = sorted(pairs.items(), key=lambda kv: -kv[1])
    rows = [Row(f'<b>{escape(split_name(k)[0])}</b>', f'{v / tot * 100:.1f}% &middot; {n0(v)}', trail=f'<span style="width:96px;display:block">{Meter(v / tot)}</span>', link=False)
            for k, v in ranked[:cap]]
    more = f'<p class="t-small rail-p">and {plural(len(ranked) - cap, "more")}</p>' if len(ranked) > cap else ''
    return Rows(rows, open_=True) + more


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
        extra = f'{n0(a["overheal"])} overheal ({a["overheal"] / a["total"] * 100:.0f}%)' if key == 'healing' and a.get('overheal') else ''
        figs = [n0(a['effective']), rate(a['effective'] / dur), f'{a["effective"] / tot * 100:.1f}%', f'{a["active_ms"] / s["duration_ms"] * 100:.1f}%']
        rows.append(StatRow(crest(r, 28), plink(r, a['name']), player_sub(r, extra) or '&nbsp;', figs, MeterClass(a['effective'] / top, colour(r)),
                            primary_phone=f'{figs[1]} {unit}'))
    if acts:
        rows.append(StatRow('', f'Total, {plural(len(acts), "player")}', mitigated_line(s, fr) if key == 'damage_taken' else '', [n0(sum(a['effective'] for a in acts)), rate(sum(a['effective'] for a in acts) / dur), '100%', ''], total=True))
    heads = [('', False), ('Player', False), ('', False), (head_amt, True), (unit, True), ('Share', True), ('Active', True)]
    return StatTable(heads, rows, 4, sort=head_amt, sortable=(head_amt, unit, 'Share', 'Active')) if acts else '<p class="empty">Nothing in this window.</p>'


def tab_actor(R, fi, source, which):
    s = R.summary(fi)
    key, unit, amt, metric, title = {'Damage done': ('damage_done', 'DPS', 'Amount', 'damage', 'Damage done'),
                                     'Healing': ('healing', 'HPS', 'Effective', 'healing', 'Healing'),
                                     'Damage taken': ('damage_taken', 'Taken / s', 'Amount', 'taken', 'Damage taken')}[which]
    has_rows = any(a['guid'] in friendlies(R, s, source) and a['effective'] > 0 for a in s[key])
    plate = Section(title, Plate(actor_table(R, fi, source, key, unit, amt), title), meta='Click a player to see their abilities',
                    action=COPY_CSV if has_rows else '')
    fr = friendlies(R, s, source)
    side: dict[str, float] = {}
    for a in s[key]:
        if a['guid'] in fr:
            for t in a.get('targets', []):
                side[t['name']] = side.get(t['name'], 0) + t['total']
    rail_title = {'Damage done': 'Targets', 'Healing': 'Healed', 'Damage taken': 'Sources'}[which]
    rail = this_pull(R, fi, source) + Section(rail_title, side_rows(side) if side else '<p class="empty">Nothing in this window.</p>')
    return plate + chart(R, fi, metric, source), rail, []


NO_GEAR = 'No gear in this pull. The game writes each player&rsquo;s gear into the log when a boss pull starts.'


def tab_player(R, fi, guid):
    """Source = one player, Damage done: their abilities, then the gear they wore at the pull (Paperdoll, report mode)."""
    s = R.summary(fi)
    a = next(x for x in s['damage_done'] if x['guid'] == guid)
    r = roster(s)[guid]
    dur = s['duration_ms'] / 1000
    top = max(ab['effective'] for ab in a['abilities'])
    def notes(ab):
        out = [f'{n0(ab["absorbed"])} absorbed'] if ab.get('absorbed') else []
        out += [f'{n0(ab["blocked"])} blocked'] if ab.get('blocked') else []
        out += [f'{v} {k.lower()}' for k, v in sorted((ab.get('misses') or {}).items())]
        return ''.join(f' &middot; {x}' for x in out)
    toggle = '<button type="button" class="link chart-t" aria-pressed="false">On the chart</button>'   # PlayerView.svelte: one ability's line on the chart
    rows = [StatRow(sicon(ab['spell_id']), f'<b>{escape(melee(ab["spell_id"], ab["name"]))}</b>{toggle}',
                    f'<span>{plural(ab["hits"], "hit")}</span> &middot; <span>{plural(ab["crits"], "crit")}</span> &middot; <span>{n0(ab["min"])}&ndash;{n0(ab["max"])}</span>'
                    + (f' &middot; {n0(ab["effective"] / ab["hits"])} each' if ab['hits'] else '') + notes(ab),
                    [n0(ab['effective']), rate(ab['effective'] / dur), f'{ab["effective"] / a["effective"] * 100:.1f}%', str(ab['hits']), f'{ab["crits"] / max(ab["hits"], 1) * 100:.0f}%'],
                    Meter(ab['effective'] / top), primary_phone=f'{rate(ab["effective"] / dur)} DPS') for ab in sorted(a['abilities'], key=lambda x: -x['effective'])]
    table = StatTable([('', False), ('Ability', False), ('', False), ('Amount', True), ('DPS', True), ('Share', True), ('Hits', True), ('Crit', True)], rows, 5,
                      sort='Amount', sortable=('Amount', 'DPS', 'Share', 'Hits', 'Crit'))
    plate = Section(f'{split_name(r["name"])[0]}&rsquo;s damage', Plate(table, 'Abilities'), meta=f'{n0(a["effective"])} damage &middot; {rate(a["effective"] / dur)} DPS',
                    action=COPY_CSV)
    comb = next((c for c in s.get('combatants', []) if c['guid'] == guid), None)
    doll = Section('Gear at the pull', f'<p class="empty">{NO_GEAR}</p>' if not comb else '', meta=escape(player_sub(r)))
    agg: dict[str, float] = {}
    for k, v in ((t['name'], t['total']) for t in a.get('targets', [])):
        agg[k] = agg.get(k, 0) + v
    rail = this_pull(R, fi, 'Your group') + Section('Targets', side_rows(agg))
    return plate + chart(R, fi, 'damage', 'Your group') + doll, rail, []


def tab_threat(R, fi, source):
    s = R.summary(fi)
    ros = roster(s)
    fr = friendlies(R, s, source)
    dur = s['duration_ms'] / 1000
    by_enemy: dict[str, float] = {}
    for t in s['threat_by_target']:
        if t['guid'] in fr:
            by_enemy[t['target_guid']] = by_enemy.get(t['target_guid'], 0) + t['threat']
    eg = max(by_enemy, key=by_enemy.get) if by_enemy else ''
    tt = sorted([t for t in s['threat_by_target'] if t['guid'] in fr and t['target_guid'] == eg], key=lambda t: -t['threat'])
    enemy = tt[0]['target_name'] if tt else ''
    series = [(split_name(t['name'])[0], colour(ros.get(t['guid'])), t['series'], f'<span class="num">{n0(t["threat"])}</span>') for t in tt]
    marks = [(x['at_ms'] / 1000, '#e5b955') for x in s['taunts']] + [(d['at_ms'] / 1000, '#f0736b') for d in s['deaths'] if d['guid'] in fr]
    cav = f'<p class="t-small">{escape(FACTS_CHART["threatCaveat"])} <a class="link" href="#">{FACTS_CHART["threatCaveatLink"]}</a></p>'
    ch = Section(f'Estimated threat on {escape(enemy)}', LineChart(series, dur, 180, [], marks, label='Estimated threat', step_ms=1000) + cav, meta='Whole fight')
    top = max((t['threat'] for t in tt), default=1)
    tot = sum(t['threat'] for t in tt) or 1
    rows = [StatRow(crest(ros.get(t['guid']), 28), pname(ros.get(t['guid'])), player_sub(ros.get(t['guid'], {})) or '&nbsp;',
                    [n0(t['threat']), f'{t["threat"] / tot * 100:.1f}%', f'{t["threat"] / top * 100:.0f}%'], MeterClass(t['threat'] / top, colour(ros.get(t['guid']))),
                    primary_phone=f'{t["threat"] / tot * 100:.1f}% of the threat') for t in tt]
    plate = Section('Estimated threat', Plate(StatTable([('', False), ('Player', False), ('', False), ('Threat', True), ('Share', True), ('Of the top', True)], rows, 3,
                                                        sort='Threat', sortable=('Threat', 'Share')), 'Estimated threat'), meta=f'On {escape(enemy)} at the end of the fight')
    trows = [Row(f'{pname(ros.get(x["source_guid"]), x["source_name"])} taunted {escape(x["target_name"])}', escape(x['spell_name']),
                 lead=f'<span class="num muted">{at(x["at_ms"])}</span>{sicon(x["spell_id"], 24)}', link=False) for x in s['taunts']]
    rail = this_pull(R, fi, source) + Section('Taunts', Rows(trows, open_=True) if trows else '<p class="empty">Nobody taunted.</p>')
    return plate + ch, rail, []


def tab_auras(R, fi, kind, source):
    s = R.summary(fi)
    ros = roster(s)
    fr = friendlies(R, s, source)
    want = 'BUFF' if kind == 'Buffs' else 'DEBUFF'
    aus = sorted([a for a in s['auras'] if a['type'] == want and a['target_guid'] in fr], key=lambda a: -a['uptime_ms'])
    dur = s['duration_ms']
    rows = []
    for a in aus:
        by = ', '.join(split_name(ros[g]['name'])[0] if g in ros else 'an enemy' for g in a['appliers'])
        sub = f'on {escape(split_name(a["target_name"])[0])} &middot; from {escape(by)}' + (' &middot; up at the pull' if a.get('seeded') else '')
        rows.append(StatRow(sicon(a['spell_id']), f'<b>{escape(spell(a["spell_id"], a["name"]))}</b>', sub,
                            [f'{a["uptime_ms"] / dur * 100:.1f}%', str(a['applications'])], Meter(a['uptime_ms'] / dur, 'muted' if kind == 'Debuffs' else ''),
                            primary_phone=f'up {a["uptime_ms"] / dur * 100:.1f}%'))
    plate = Section(kind, Plate(StatTable([('', False), ('Aura', False), ('Uptime', False), ('Up', True), ('Applied', True)], rows, 2, sort='Up', sortable=('Up', 'Applied')), kind)
                    if rows else Plate(f'<p class="empty">No {kind.lower()} in this window.</p>', kind), meta=f'{len(rows)} on players &middot; uptime over the whole fight')
    return plate, this_pull(R, fi, source), []


def tab_deaths(R, fi, source):
    """The death recap (SHOULDs as ruled): last hits and heals received in time order (heals as green rows), health as numbers
    ("74 / 114"), what the player held. Defensives ready and unused need data the log does not carry (gap G13)."""
    s = R.summary(fi)
    ros = roster(s)
    fr = friendlies(R, s, source)
    blocks = []
    for d in [d for d in s['deaths'] if d['guid'] in fr]:
        r = ros.get(d['guid'])
        kb = d['killing_blow']
        head_row = Row(pname(r, d['name']), f'Killed by {escape(kb["source_name"])} &middot; {escape(melee(kb["spell_id"], kb["spell_name"]))} &middot; {n0(kb["amount"])} ({n0(kb.get("overkill", 0))} overkill) &middot; {n0(kb["max_hp"])} health',
                       lead=f'<span class="num muted">{at(d["at_ms"])}</span>{crest(r, 32)}', trail='<button type="button" class="link">The 20 s before it &rsaquo;</button>', link=False)
        events = [('hit', h) for h in d['last']] + [('heal', h) for h in d.get('heals', [])]
        lines = []
        for kind, h in sorted(events, key=lambda e: e[1]['at_ms']):
            mx = h.get('max_hp') or kb['max_hp']
            hp = f'{n0(h["hp_after"])} / {n0(mx)}' if 'hp_after' in h else 'dead'
            if kind == 'heal':
                lines.append(StatRow(sicon(h['spell_id'], 24), f'<b class="kill">{escape(spell(h["spell_id"], h["spell_name"]))}</b>', f'<span>{at(h["at_ms"])}</span> &middot; from {escape(split_name(h["source_name"])[0])}',
                                     [f'<span class="kill">+{n0(h["amount"])}</span>', hp], Meter(h.get('hp_after', 0) / mx, 'kill'), primary_phone=f'+{n0(h["amount"])}'))
            else:
                lines.append(StatRow(sicon(h['spell_id'], 24), f'<b>{escape(melee(h["spell_id"], h["spell_name"]))}</b>', f'<span>{at(h["at_ms"])}</span> &middot; from {escape(h["source_name"])}',
                                     [f'&minus;{n0(h["amount"])}', hp], Meter(h.get('hp_after', 0) / mx, 'kill'), primary_phone=f'&minus;{n0(h["amount"])} &middot; {hp}'))
        held = ', '.join(escape(x['name']) for x in d['auras_held'])
        healed = any(h['at_ms'] >= d['at_ms'] - 10000 for h in d.get('heals', []))
        foot = ' '.join(x for x in (f'Holding: {held}.' if held else '', '' if healed else 'No healing landed in the last 10 seconds.') if x)
        body = StatTable([('', False), ('Last hits and heals', False), ('Health after', False), ('Amount', True), ('Health', True)], lines, 2, sort='', sortable=())
        blocks.append(f'<div class="dblock">{Rows([head_row])}{body}' + (f'<p class="t-small dfoot">{foot}</p>' if foot else '') + '</div>')
    plate = Section('Deaths', Plate(''.join(blocks), 'Deaths') if blocks else Plate('<p class="empty">Nobody died.</p>', 'Deaths'), meta=f'{len(blocks)} in this fight',
                    action=COPY_CSV if blocks else '')
    return plate + chart(R, fi, 'taken', source), this_pull(R, fi, source), []


def tab_kicks(R, fi, which, source):
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
    return plate, this_pull(R, fi, source), []


def tab_resources(R, fi, source):
    s = R.summary(fi)
    ros = roster(s)
    fr = friendlies(R, s, source)
    rows = []
    for x in s['resources']:
        if x['guid'] not in fr or x['guid'] not in ros:
            continue
        r = ros[x['guid']]
        rows.append(StatRow(crest(r, 28), pname(r), POWER.get(x['power_type'], 'Power'), [n0(x['gained']), n0(x['spent']), n0(x['wasted']), mmss(x['zero_ms'])],
                            Spark(x['series'], colour(r), total=x['max'] or None), primary_phone=f'{n0(x["spent"])} spent'))
    plate = Section('Resources', Plate(StatTable([('', False), ('Player', False), ('Over the fight', False), ('Gained', True), ('Spent', True), ('Wasted', True), ('Time empty', True)], rows, 4), 'Resources')
                    if rows else Plate('<p class="empty">No resource changes in this window.</p>', 'Resources'), meta='Mana, rage, energy and focus, whole fight')
    return plate, this_pull(R, fi, source), []


def tab_casts(R, fi, source):
    s = R.summary(fi)
    ros = roster(s)
    fr = friendlies(R, s, source)
    dur = s['duration_ms']
    rows = []
    for c in sorted([c for c in s['casts'] if c['guid'] in fr or c['owner_guid'] in fr], key=lambda c: -(c['succeeded'] or c['started'])):
        r = ros.get(c['guid']) or ros.get(c['owner_guid'])
        who = split_name(c['name'])[0] + (f' ({split_name(ros[c["owner_guid"]]["name"])[0]}&rsquo;s pet)' if c['guid'] != c['owner_guid'] and c['owner_guid'] in ros else '')
        n = c['succeeded'] or c['started']
        rows.append(StatRow(sicon(c['spell_id']), f'<b>{escape(spell(c["spell_id"], c["spell_name"]))}</b>', f'<span style="color:{colour(r)}">{who}</span>',
                            [str(n), f'{n / (dur / 60000):.1f}', str(c['failed'])], Spark([], colour(r) if r else '#9a9484', marks=[t / dur for t in c['sequence']]),
                            primary_phone=plural(n, 'cast')))
    plate = Section('Casts', Plate(StatTable([('', False), ('Spell', False), ('When', False), ('Casts', True), ('Per min', True), ('Failed', True)], rows, 3,
                                             sort='Casts', sortable=('Casts', 'Per min', 'Failed')), 'Casts') if rows else Plate('<p class="empty">No casts in this window.</p>', 'Casts'),
                    meta=f'Every caster, {mmss(dur)}')
    return plate, this_pull(R, fi, source), []


def tab_rating(R, fi, source):
    rc = COPY['ratingCopy']
    trash = R.fight(fi)['kind'] == 'trash'
    plate = Section(rc['panelHeading'], Plate(f'<p class="empty">{escape(rc["noRatingForTrash"] if trash else rc["panelEmpty"])}</p>', 'Rating'),
                    action=Link(escape(rc['explainLink']) + ' &rsaquo;'))
    rail = this_pull(R, fi, source) + Section('The cap', f'<p class="t-small rail-p">{escape(rc["cappedRule"])}</p>')
    return plate, rail, []


def tab_coach(R, fi, source):
    """The coach answers a trash pull with its own line (api/internal/coach/copy.go NotReviewableTrash). The filled review
    needs a real boss pull (D1)."""
    me = roster(R.summary(fi))[VIEWER]
    plate = Section('Coach', Plate(f'<p class="empty">{escape(COACH_TRASH)}</p>', 'Coach review'), meta=f'{pname(me)} &middot; {escape(player_sub(me))}')
    return plate, this_pull(R, fi, source), []


MECH_TRASH = 'Mechanics are judged per boss, and this is a trash segment: pick a boss pull, or the whole night.'   # MechanicsMode.svelte, verbatim


def tab_mechanics(R, fi, source):
    s = R.summary(fi)
    if R.fight(fi)['kind'] == 'trash' and not s['mechanics']['table_found']:
        plate = Section('Mechanics', Plate(f'<p class="empty">{MECH_TRASH}</p>', 'Mechanics'), meta='Problems first, then what went right')
        return plate, this_pull(R, fi, source), []
    raise NotImplementedError('a boss pull with a mechanics table: no real log has one yet (D1)')


def events(R, fi, names: set[str] | None = None):
    import duckdb
    p = R.base / f'fights/{fi}/events.parquet'
    q = (f"select (time_unix_nano - (select min(time_unix_nano) from '{p}'))/1e6 as ms, kind, source_name, dest_name, spell_id, spell_name, amount "
         f"from '{p}' where kind not in ('combatant_info','encounter_start','encounter_end')")
    rows = duckdb.sql(q + ' order by 1').fetchall()
    # the page's Source scope (EventsView.svelte): a line stays when anyone it involves is in scope
    return [r for r in rows if names is None or r[2] in names or r[3] in names]


def ev_sentence(ros_by_name, k, src, dst, sid, sp, amt):
    n = lambda x: f'<span style="color:{colour(ros_by_name.get(x))}">{escape(split_name(x)[0])}</span>' if x else ''
    spn = escape(spell(sid, sp)) if sp else 'melee'
    return {
        'cast_success': f'{n(src)} casts {spn}', 'cast_start': f'{n(src)} begins casting {spn}', 'cast_failed': f'{n(src)} fails to cast {spn}',
        'damage': f'{n(src)} hits {n(dst)} with {spn}', 'damage_landed': f'{n(src)} hits {n(dst)} in melee', 'missed': f'{n(src)} misses {n(dst)}',
        'heal': f'{n(src)} heals {n(dst)} with {spn}', 'death': f'{n(dst)} dies', 'party_kill': f'{n(src)} kills {n(dst)}',
        'aura_applied': f'{spn} on {n(dst)} (from {n(src)})', 'aura_removed': f'{spn} fades from {n(dst)}', 'aura_refresh': f'{spn} refreshed on {n(dst)}',
        'energize': f'{n(src)} gains power from {spn}', 'instakill': f'{n(src)} uses {spn} on {n(dst)}', 'enchant': f'{n(src)} applies {spn}',
    }.get(k, f'{k} {n(src)} {n(dst)}')


EVENTS_PAGE = 200   # EventsView.svelte PAGE


def tab_events(R, fi, source):
    s = R.summary(fi)
    scope = in_scope(R, s, source)
    by_name = {r['name']: r for r in s['roster']}
    evs = events(R, fi, {r['name'] for r in scope})
    shown = evs[:EVENTS_PAGE]
    rows = [StatRow(sicon(sid, 20) if sid else '', f'<span class="ev">{ev_sentence(by_name, k, src, dst, sid, sp, amt)}</span>', '',
                    [at(ms), n0(amt) if amt else '&ndash;'], primary_phone=at(ms), bar=False) for ms, k, src, dst, sid, sp, amt in shown]
    more = len(evs) - len(shown)
    foot = f'<div class="pager">{Button(f"Show {min(EVENTS_PAGE, more)} more &middot; {more:,} not yet shown", small=True)}</div>' if more > 0 else ''
    plate = Section('Events', Plate(StatTable([('', False), ('Event', False), ('Time', True), ('Amount', True)], rows, 2, bar=False) + foot, 'Events'),
                    meta=f'{len(evs):,} events in this window, in order')
    rail = this_pull(R, fi, source) + Section('Show', '<div class="chips">' + ''.join(f'<label class="chip"><input type="checkbox" checked> {t}</label>' for t in ('Casts', 'Damage', 'Heals', 'Auras', 'Deaths', 'Power')) + '</div>')
    return plate, rail, []


QUERY = ("SELECT source_name, spell_name, count(*) AS hits, sum(amount) AS total\nFROM read_parquet('events.parquet')\n"
         "WHERE kind = 'damage' AND (time_unix_nano - (SELECT min(time_unix_nano) FROM read_parquet('events.parquet'))) / 1000000 BETWEEN 0 AND {end}\n"
         "GROUP BY 1, 2\nORDER BY total DESC")
QUERY_ROWS = 200   # lib/report/query.ts MAX_ROWS


def tab_queries(R, fi, source):
    import duckdb
    s = R.summary(fi)
    p = R.base / f'fights/{fi}/events.parquet'
    sql = QUERY.format(end=s['duration_ms'])
    res = duckdb.sql(sql.replace("read_parquet('events.parquet')", f"read_parquet('{p}')")).fetchall()
    by_name = {r['name']: r for r in s['roster']}
    rows = [StatRow('', f'<b style="color:{colour(by_name.get(a))}">{escape(split_name(a)[0])}</b>', '', [escape(melee(0, b or "")), str(c), n0(d)], primary_phone=n0(d), bar=False)
            for a, b, c, d in res[:QUERY_ROWS]]
    starts = ''.join(f'<a href="#" class="{"on" if i == 0 else ""}" role="radio" aria-checked="{"true" if i == 0 else "false"}">{t}</a>' for i, t in enumerate(
        ['Damage by ability, exactly, in this window', 'Pet damage, by pet', 'Every miss, by type', 'Healing and overhealing by spell', 'Healing by target, with overhealing']))
    body = (f'<div class="qwrap"><span class="t-label">Starting points</span><div class="seg qstarts" role="radiogroup" aria-label="Starting points">{starts}</div>'
            f'<label class="t-label" for="q">Query</label><pre class="code" id="q">{escape(sql)}</pre><div class="qrun">{Button("Run", primary=True)}<span class="t-small">{plural(len(res), "row")}</span></div></div>'
            + StatTable([('', False), ('source_name', False), ('spell_name', True), ('hits', True), ('total', True)], rows, 3, bar=False))
    plate = Section('Queries', Plate(body, 'Query'), meta='SQL over this fight&rsquo;s events, run in your browser')
    rail = this_pull(R, fi, source) + Section('Columns', '<p class="t-small rail-p">time_unix_nano, kind, event, source_name, dest_name, spell_id, spell_name, amount, overkill, overheal, absorbed, critical, miss_type, aura_type, power_type and 50 more.</p>')
    return plate, rail, []


def tab_timeline(R, fi, source):
    s = R.summary(fi)
    ros = roster(s)
    scope = sorted(in_scope(R, s, source), key=lambda r: -r['damage_done'])
    evs = events(R, fi, {r['name'] for r in scope})
    dur = s['duration_ms']
    W, LH = 1000, 34
    actors = list(scope) + ([{'guid': 'boss', 'name': R.fight(fi)['name']}] if R.fight(fi)['kind'] != 'trash' else [])
    lanes, labels = '', ''
    for i, r in enumerate(actors):
        y = i * LH
        lanes += f'<rect x="0" y="{y + 15}" width="{W}" height="4" fill="#1c2230"/>'
        nm = r['name']
        for a in [a for a in s['auras'] if a['target_guid'] == r['guid']]:
            for sg in a['segments']:
                lanes += f'<rect x="{sg["start_ms"] / dur * W:.1f}" y="{y + 22}" width="{max(1, (sg["end_ms"] - sg["start_ms"]) / dur * W):.1f}" height="5" rx="2" fill="{"#7fd48a99" if a["type"] == "BUFF" else "#f0736b99"}"/>'
        for ms, k, src, dst, sid, sp, amt in evs:
            if k == 'cast_success' and src == nm:
                lanes += f'<rect x="{ms / dur * W - 1:.1f}" y="{y + 6}" width="2" height="14" fill="#e5b955"/>'
            elif k in ('damage', 'damage_landed') and src == nm:
                lanes += f'<rect x="{ms / dur * W - 1:.1f}" y="{y + 9}" width="2" height="8" fill="{colour(ros.get(r["guid"]))}"/>'
            elif k == 'death' and dst == nm:
                lanes += f'<rect x="{ms / dur * W - 1.5:.1f}" y="{y + 2}" width="3" height="28" fill="#f0736b"/>'
        labels += f'<div class="tl-l">{crest(r, 20) if r["guid"] != "boss" else EmptyCrest(20)}<span style="color:{colour(r) if r["guid"] != "boss" else "var(--text)"}">{escape(split_name(nm)[0])}</span></div>'
    svg = f'<svg viewBox="0 0 {W} {len(actors) * LH}" preserveAspectRatio="none" style="height:{len(actors) * LH}px" role="img" aria-label="Every player&rsquo;s casts, hits, auras and death on one time axis">{lanes}</svg>'
    ticks = ''.join(f'<span>{mmss(dur * q)}</span>' for q in (0, .25, .5, .75, 1))
    legend = LanesLegend([('#e5b955', 'Cast'), ('#3fc7eb', 'Hit, in the class colour'), ('#7fd48a99', 'Buff up'), ('#f0736b99', 'Debuff up'), ('#f0736b', 'Death')])
    body = f'{legend}<div class="tlg"><div class="tl-ls">{labels}</div><div class="tl-p">{svg}<div class="ax">{ticks}</div></div></div>'
    plate = Section('Timeline', Plate(Lanes(body), 'Timeline'), meta=f'Every player, {mmss(dur)} &middot; hover for the event')
    return plate, this_pull(R, fi, source), []


def tab_compare(R, fi, source, other):
    a, b = R.summary(fi), R.summary(other)
    ra, rb = roster(a), roster(b)
    rows = []
    for r in sorted(in_scope(R, a, source), key=lambda r: -r['dps']):
        o = rb.get(r['guid'])
        diff = (r['dps'] - o['dps']) if o else None
        dtxt = '&ndash;' if diff is None else f'<span class="{"kill" if diff > 0 else "wipe" if diff < 0 else ""}">{"+" if diff > 0 else ("&minus;" if diff < 0 else "")}{rate(abs(diff))}</span>'
        rows.append(StatRow(crest(r, 28), pname(r), player_sub(r) or '&nbsp;', [rate(r['dps']), rate(o['dps']) if o else '&ndash;', dtxt],
                            primary_phone=f'{rate(r["dps"])} vs {rate(o["dps"]) if o else "&ndash;"}', bar=False))
    fa, fb = R.fight(fi), R.fight(other)
    plate = Section('Compare', Plate(StatTable([('', False), ('Player', False), ('This pull', True), (fight_label(fb), True), ('Difference', True)], rows, 3, bar=False), 'Compare'),
                    meta=f'DPS, {fight_label(fa)} against {fight_label(fb)}')
    return plate, this_pull(R, fi, source), ['Players missing from the other pull show a dash.']


# ------------------------------------------------------------------------------------------------ pages
FACTS_CHART = {'threatCaveat': 'An estimate: 1 threat per point of damage and 0.5 per point of healing, split across the enemies in combat. No stance, form, talent or taunt modifiers yet.',
               'threatCaveatLink': 'What it leaves out'}   # lib/chart/copy.ts, verbatim


PAGE_CSS = '''<style>
button.sel{background:none;cursor:pointer}
.gloss>summary{list-style:none;cursor:pointer;min-height:36px;display:inline-flex;align-items:center;font-size:14px;font-weight:600}.gloss>summary::-webkit-details-marker{display:none}
.chart-t{margin-left:8px;font-size:13px;font-weight:600}
.ans-act{display:flex;gap:8px;margin-top:6px}@media(min-width:900px){.ans-act{justify-content:flex-end}}
.ans-more{display:none}
.head .ans{min-width:0}@media(min-width:900px){.head .ans{min-width:360px}}
.dblock+.dblock{border-top:1px solid var(--line)}.dblock .st-h{padding-top:6px}.dfoot{padding:8px 16px 12px}
b.kill,.f .kill{color:var(--kill)}
.rail-l{display:flex;flex-direction:column;gap:8px;padding-top:10px;font-size:14px;line-height:20px}.rail-p{padding-top:10px}
.qwrap{display:flex;flex-direction:column;gap:8px;padding:14px 16px;border-bottom:1px solid var(--line-soft)}
.qstarts{flex-wrap:wrap}.qstarts a{height:32px;font-size:13px}
pre.code{margin:0;padding:10px 12px;white-space:pre-wrap;font:13px/20px ui-monospace,Menlo,monospace;color:var(--text);border:1px solid var(--line);border-radius:4px}
.qrun{display:flex;gap:12px;align-items:center}
.chips{display:flex;flex-wrap:wrap;gap:8px 16px;padding-top:10px;font-size:14px}
.pager{display:flex;gap:16px;align-items:center;justify-content:flex-start;padding:12px 16px;border-top:1px solid var(--line-soft)}
.tlg{display:grid;grid-template-columns:160px minmax(0,1fr);padding:4px 16px 12px}.tl-ls{display:flex;flex-direction:column}.tl-l{height:34px;display:flex;align-items:center;gap:8px;font:600 14px/16px var(--font-body);white-space:nowrap;overflow:hidden}
.tl-p svg{display:block;width:100%}.tl-p .ax{display:flex;justify-content:space-between;font:600 12px/16px var(--font-body);color:var(--muted);padding-top:4px}
.lanes .lg{padding-inline:16px}
.ev{font-size:14px}
.menu-at{position:relative}.menu-at .menu{top:4px;left:0;max-height:70vh;overflow:auto}
.mi-empty{padding:8px 16px 12px;font-size:14px;color:var(--muted)}
.mi-toggle{display:flex;gap:10px;align-items:center;min-height:44px;padding:4px 16px;border-top:1px solid var(--line-soft);font-size:14px;color:var(--text)}
/* phone (ruled SHOULD): Sim and Copy in one actions menu, the sample note under the meta, ONE tab line that scrolls sideways */
@media(max-width:767px){
 .ans-act{display:none}.ans-more{display:block;position:absolute;top:0;right:0}.head .ans{position:relative}
 .ph-note,.ph-fights{display:none}
 .ph-tabs.strip-tabs{flex-wrap:nowrap;overflow-x:auto;scrollbar-width:none;gap:0;margin-inline:-16px;padding-inline:8px}
 .ph-tabs.strip-tabs::-webkit-scrollbar{display:none}.ph-tabs.strip-tabs a{flex:none;padding:12px 12px}
 .tlg{grid-template-columns:96px minmax(560px,1fr);overflow-x:auto}}
</style>'''


def body_for(R, fi, tab, source, viewer=''):
    return {
        'Summary': lambda: tab_summary(R, fi, source),
        'Damage done': lambda: tab_actor(R, fi, source, 'Damage done'),
        'Healing': lambda: tab_actor(R, fi, source, 'Healing'),
        'Damage taken': lambda: tab_actor(R, fi, source, 'Damage taken'),
        'Threat': lambda: tab_threat(R, fi, source),
        'Buffs': lambda: tab_auras(R, fi, 'Buffs', source),
        'Debuffs': lambda: tab_auras(R, fi, 'Debuffs', source),
        'Deaths': lambda: tab_deaths(R, fi, source),
        'Interrupts': lambda: tab_kicks(R, fi, 'Interrupts', source),
        'Dispels': lambda: tab_kicks(R, fi, 'Dispels', source),
        'Resources': lambda: tab_resources(R, fi, source),
        'Casts': lambda: tab_casts(R, fi, source),
        'Rating': lambda: tab_rating(R, fi, source),
        'Coach': lambda: tab_coach(R, fi, source),
        'Mechanics': lambda: tab_mechanics(R, fi, source),
        'Timeline': lambda: tab_timeline(R, fi, source),
        'Events': lambda: tab_events(R, fi, source),
        'Queries': lambda: tab_queries(R, fi, source),
        'Compare': lambda: tab_compare(R, fi, source, COMPARE_PULL),
    }[tab]()


DATA_NOTE = 'web/public/sample-report/sample (the real 17 Sept 2026 beta night, Zephras Isle, client 1.60.1)'


def report_page(tab: str, fi: int = DEFAULT_PULL, source: str = 'Your group', viewer: str = '', sample: bool = False,
                open_m: bool = False, menu: bool = False, player: str = '', extra_notices: list[str] | None = None) -> str:
    R = Report()
    bar_source = source
    if player:   # the answer keeps the group's scope; the bar's Source names the player
        main, rail, nts = tab_player(R, fi, player)
        bar_source = split_name(roster(R.summary(fi))[player]['name'])[0]
    else:
        main, rail, nts = body_for(R, fi, tab, source, viewer)
    nts = (extra_notices or []) + nts
    main += GLOSSARY
    compare = fight_label(R.fight(COMPARE_PULL)) if tab == 'Compare' else ''
    menu_html = fight_menu(R, fi) if menu else ''
    who = None
    if viewer:
        me = roster(R.summary(fi))[viewer]
        who = (split_name(me['name'])[0], me['class'].lower(), escape(me['class']))
    body = (head(R, fi, tab, source, viewer, sample) + bar(R, fi, tab, bar_source, open_m, compare) + menu_html + notices(nts)
            + Page('split', main, rail) + PAGE_CSS)
    note = (f'Report mock (design/specs/2026-11-07-logs-reports.md), built only from ws_prims. Data: {DATA_NOTE}, pull {fi}, source {escape(source)}'
            + ('; SAMPLE viewer: an owner of Soren Everhart' if viewer else '') + f'; build {BUILD} names and icons.')
    return build_page(f'{R.title}: {tab}', body, 'Logs', who, note)


# The report's help, last in the main column (coordinator ruling 2026-11-07: the glossary stays): a quiet "What these
# mean ›" line opening the words the tables use in place; the terms are lib/report/glossary.ts's, folded here.
GLOSSARY = (f'<div data-region="{REGION["footer_note"]}"><details class="gloss"><summary class="link">What these mean &rsaquo;</summary>'
            '<dl><dt>Parse</dt><dd>Where a kill sits among every ranked kill of the same boss by the same spec: the shipped glossary, '
            'lib/report/glossary.ts.</dd></dl></details></div>')


def sample_page() -> str:
    """/reports/sample (MUST 2 as ruled): the sample opens on its default trash pull, and the head leads with "177 mobs · 12:43"."""
    return report_page('Summary', sample=True)


def skeleton_rows(n: int) -> str:
    return ''.join(f'<div class="st-r stc" style="--n:4">{Skeleton(28, 28, True)}<span>{Skeleton(180, 16)}</span><span></span>{Skeleton(48, 16)}{Skeleton(48, 16)}{Skeleton(48, 16)}{Skeleton(48, 16)}</div>' for _ in range(n))


def state_page(state: str) -> str:
    R = Report()
    rc = COPY['reportCopy']
    if state == 'loading':
        # The sample's shell knows its title (the parity render, FS_REPORT_SAMPLE): the heading is the text, not a block.
        head_html = PageHeader(escape(R.title) if AS_SAMPLE else Skeleton(520, 40), meta=[Skeleton(320, 16)], answer=f'{Skeleton(200, 16)}{Skeleton(160, 40)}{Skeleton(220, 16)}',
                               tabs=TabStripMore(['Summary'], '', 1, 1).replace('>Summary<', ' aria-disabled="true">Summary<'))
        main = Section('Players', Plate(f'<div class="st" style="--n:4">{skeleton_rows(6)}</div>', 'Players'))
        rail = Section('This pull', ''.join(f'<p>{Skeleton(300, 16)}</p>' for _ in range(5)))
        body = head_html + bar(R, DEFAULT_PULL, 'Summary', 'Your group') + Page('split', main, rail)
    elif state in ('error', 'missing'):
        msg = 'Report data did not load' if state == 'error' else 'No report with that id'
        act = Button('Try again', small=True) if state == 'error' else Link('Open your reports &rsaquo;')
        body = PageHeader('Report') + Page('ledger', Plate(f'<p class="empty" role="alert">{msg}. {act}</p>', 'Report'))
    elif state == 'empty':
        # Pull 13 ran without the logger's group: the scope falls back to everyone in range (lib/report/group.ts) and
        # one muted line says so (coordinator ruling 2026-11-07).
        return report_page('Damage done', fi=EMPTY_PULL, source='Everyone in range', extra_notices=['Your group wasn&rsquo;t in this pull.'])
    elif state in ('processing', 'waiting', 'failed'):
        line = {'processing': rc['parsing'], 'waiting': rc['waiting'], 'failed': f'{rc["parseFailed"]} <a class="link" href="#">{rc["parseFailedLink"]}</a>'}[state]
        title = rc['yourUpload'] if state != 'waiting' else escape(R.zone)
        meta = ['Uploaded 2 minutes ago'] if state != 'waiting' else [Tag('Live', 'kill'), 'Logging now']
        body = PageHeader(escape(title), meta=meta, tabs=TabStripMore(['Summary'], '', 1, 1).replace('>Summary<', ' aria-disabled="true">Summary<')) + Page('ledger', Plate(f'<p class="empty" role="status">{line}</p>', 'Fights'))
    return build_page(f'Report: {state}', body + PAGE_CSS, 'Logs', None, f'Report state mock: {state}. Data: {DATA_NOTE}.')


PAGES = {
    'web-report-summary': lambda: report_page('Summary'),
    'web-report-damage-done': lambda: report_page('Damage done'),
    'web-report-healing': lambda: report_page('Healing'),
    'web-report-damage-taken': lambda: report_page('Damage taken'),
    'web-report-deaths': lambda: report_page('Deaths', source='Everyone in range'),   # the group had no death on pull 6; two players in range did
    'web-report-mechanics': lambda: report_page('Mechanics'),
    'web-report-timeline': lambda: report_page('Timeline'),
    'web-report-threat': lambda: report_page('Threat'),
    'web-report-casts': lambda: report_page('Casts'),
    'web-report-buffs': lambda: report_page('Buffs'),
    'web-report-debuffs': lambda: report_page('Debuffs'),
    'web-report-rating': lambda: report_page('Rating'),
    'web-report-coach': lambda: report_page('Coach', viewer=VIEWER),
    'web-report-interrupts': lambda: report_page('Interrupts'),
    'web-report-dispels': lambda: report_page('Dispels'),
    'web-report-resources': lambda: report_page('Resources'),
    'web-report-events': lambda: report_page('Events'),
    'web-report-queries': lambda: report_page('Queries'),
    'web-report-compare': lambda: report_page('Compare'),
    'web-report-player': lambda: report_page('Damage done', player=VIEWER),
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
