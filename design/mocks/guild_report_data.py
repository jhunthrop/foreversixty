"""The raid night report's data (design/parity.md, "The guild states"; state `report`): Onyxia, Wed 9 Dec (the first night), read by an officer
(Dunmaro). One source for both sides: ws_guild2.raid_report draws the mock from `raid_view()` and guild_fixtures.py writes the
same raid view into web/tests/e2e/fixtures/guild-parity/report-raid.json, which web/tests/e2e/support/guild-parity/report.ts
routes. The viewer, guild and home reads are the raid planner's (guild_raid.py: the same officer on the same guild). Shapes:
web/src/lib/guild/api.ts (GuildRaidView), attendance-contract.ts (RaidReport). Roster and sign-up states come from guild_data.R.
"""
import guild_data as gd
from guild_data import R, ZONE
from guild_raid import ONYXIA_ENCOUNTER as ONYXIA, RAID_ZONE_KEY, RAID_ZONE_NAME as ZONE_NAME, VIEWER, counts, key_of as key, member

# The guild's first raid night, Wed 9 Dec, 8:00 pm CST (guild_data.LAST_RAID_*), the week before the planner's raid.
RAID_ID = gd.LAST_RAID_ID
TITLE = ZONE_NAME
STARTS_AT = gd.LAST_RAID_AT
UPLOADED_AT = '2026-12-10T05:10:00Z'
NIGHT_MS = 90 * 60_000
AWAY = gd.NIGHT_AWAY
SPELL = {'Flame Breath': 18435, 'Deep Breath': 17086, 'Tail Sweep': 15847}
PULLS = gd.NIGHT_PULLS
KILL_FIGHT = next(p[0] for p in PULLS if p[3] == 'kill')
# name, DPS on the kill, the fight's rating
PARSES = [('Frostquill', 812, 96), ('Doomwhisper', 744, 88), ('Dunmaro', 618, 61)]


def row(name: str) -> tuple:
    return next(r for r in R if r[0] == name)


def who(name: str) -> dict:
    r = row(name)
    return {'character_key': key(name), 'name': name, 'class': r[1]}


def signed_state(state: str) -> str | None:
    return None if state == 'none' else state


def death(name: str, ability: str, at_ms: int) -> dict:
    hit = {'at_ms': at_ms, 'spell_id': SPELL[ability], 'ability': ability, 'source': 'Onyxia', 'amount': 4210, 'overkill': 1050}
    return {**who(name), 'at_ms': at_ms, 'killing_blow': hit, 'last_hits': []}


def pulls() -> list[dict]:
    out = []
    for index, name, attempt, result, ms, deaths in PULLS:
        causes = [death(n, ability, 30_000 + 20_000 * i) for i, (n, ability) in enumerate(deaths)]
        out.append({'fight_index': index, 'encounter_id': None if attempt is None else ONYXIA, 'name': name, 'attempt': attempt,
                    'result': result, 'duration_ms': ms, 'deaths': len(deaths), 'death_causes': causes})
    return out


def headline() -> dict:
    boss = [p for p in PULLS if p[2] is not None]
    kills = sum(1 for p in boss if p[3] == 'kill')
    return {'kills': kills, 'wipes': len(boss) - kills, 'attempts': len(boss), 'deaths': sum(len(p[5]) for p in PULLS)}


def parses() -> list[dict]:
    r = [p for p in PULLS if p[0] == KILL_FIGHT][0]
    return [{**who(n), 'role': 'dps', 'encounter_id': ONYXIA, 'encounter': r[1], 'fight_index': KILL_FIGHT, 'metric': 'dps',
             'value': value, 'rating': rating} for n, value, rating in PARSES]


def death_rows() -> list[dict]:
    tally: dict[str, list[int]] = {}
    for _, _, _, result, _, deaths in PULLS:
        for n, _ in deaths:
            on = tally.setdefault(n, [0, 0])
            on[0 if result == 'wipe' else 1] += 1
    return [{**who(n), 'on_wipes': w, 'on_kills': k} for n, (w, k) in tally.items()]


def report() -> dict:
    boss_pulls = [p for p in PULLS if p[2] is not None]
    kill = [p for p in PULLS if p[0] == KILL_FIGHT][0]
    return {
        'headline': headline(),
        'pulls': pulls(),
        'parses': parses(),
        'kill_order': [{'encounter_id': ONYXIA, 'name': kill[1], 'fight_index': KILL_FIGHT, 'killed': True, 'kill_ms': kill[4],
                        'wipes_before': len(boss_pulls) - 1, 'deaths': len(kill[5])}],
        'unnamed': None,
        'deaths': death_rows(),
        # The log carried no character sheet: nothing can be said about flasks, food or enchants.
        'consumables': {'log_has_buffs': False, 'rows': [],
                        'by_pull': [{'fight_index': p[0], 'name': p[1], 'attempt': p[2], 'reason': 'no_character_sheet', 'rows': None}
                                    for p in boss_pulls]},
        'loot': [{'encounter_id': ONYXIA, 'boss': kill[1], 'awards': [
            {'item_id': d['item_id'], 'item_name': d['item_name'], 'quality': d['quality'], 'icon': d['icon'], **who(d['name']), 'by_name': d['awarded_by'],
             'at': d['at']} for d in awards()]}],
        'moves': {'previous': None, 'rows': []},
        'posted_at': None,
    }


def attendance() -> list[dict]:
    out = []
    for r in R:
        name, _, spec, _, _, state = r
        here = name not in AWAY
        outcome = 'present' if here else 'absent'
        out.append({**who(name), 'spec': spec, 'signed': signed_state(state), 'role': member(name)['role'], 'outcome': outcome, 'came_as': None,
                    'first_seen_offset_s': 0 if here else None, 'pulls_present': len(PULLS) if here else 0, 'pulls_total': len(PULLS),
                    'excused': False, 'own': name == VIEWER, 'auto_outcome': outcome, 'auto_excused': False, 'override': None})
    return out


def night() -> dict:
    h = headline()
    signed = {r[0] for r in R if r[5] in ('in', 'late')}
    here = {r[0] for r in R if r[0] not in AWAY}
    return {'signed': len(signed), 'came': len(signed & here), 'late': 0, 'not_in_log': len(signed - here), 'joined': len(here - signed),
            'excused': 0, 'kills': h['kills'], 'wipes': h['wipes'], 'deaths': h['deaths'], 'duration_ms': NIGHT_MS}


def signups() -> list[dict]:
    """The first night's sign-ups: the roster's states (guild_data.R), as the planner's next raid reads them."""
    out = []
    for r in R:
        if r[5] == 'none':
            continue
        m = member(r[0])
        out.append({**who(r[0]), 'spec': m['spec'], 'role': m['role'], 'state': m['state'], 'note': '', 'own': r[0] == VIEWER,
                    'group': m['group'], 'spec_key': m['spec_key']})
    return out


def raid_view() -> dict:
    rows = signups()
    v = member(VIEWER)
    return {
        'id': RAID_ID, 'zone': RAID_ZONE_KEY, 'zone_name': ZONE_NAME, 'zone_kind': 'raid', 'key_item': None, 'title': TITLE,
        'starts_at': STARTS_AT, 'duration_minutes': 180, 'status': 'done', 'size': 40, 'guild_timezone': ZONE, 'notes_html': '', 'notes': '',
        'counts': counts(), 'signups': rows,
        'not_signed': [{**who(m['name']), 'spec': m['spec'], 'group': m['group'], 'spec_key': m['spec_key']}
                       for m in map(member, (r[0] for r in R)) if m['state'] == 'none'],
        'viewer': {'characters': [{**who(VIEWER), 'spec': v['spec'], 'verified': True, 'state': v['state'], 'role': v['role'],
                                   'raid_spec': v['spec_key'], 'spec_key': v['spec_key'], 'default_spec': v['spec_key']}]},
        'discord': {'available': True, 'connected': True, 'post_default': True, 'remind_default': True, 'channel_name': 'raids'},
        'discord_post': None, 'cover': None, 'comp_target': None, 'assignments': None,
        'pre_done_status': 'locked',
        'match': {'report_id': gd.LAST_REPORT_ID, 'zone_name': ZONE_NAME, 'uploaded_at': UPLOADED_AT, 'uploaded_by_name': VIEWER,
                  'pulls': len(PULLS), 'kills': headline()['kills'], 'duration_ms': NIGHT_MS, 'how': 'auto'},
        'night': night(), 'attendance': attendance(), 'report': report(),
    }


def awards() -> list[dict]:
    """The loot council's awards that night (guild_overview_fixtures.decisions: one source)."""
    from guild_overview_fixtures import decisions
    return decisions()


def write_report_fixtures(write) -> list:
    return [write('report-raid.json', raid_view())]


# ------------------------------------------------ what the mock prints from the data (web/src/lib/guild/report-night-view.ts words)
SEP = ' &middot; '


def clock(ms: int) -> str:
    s = ms // 1000
    return f'{s // 60}:{s % 60:02d}'


def local(iso: str):
    from datetime import datetime
    from zoneinfo import ZoneInfo
    return datetime.fromisoformat(iso.replace('Z', '+00:00')).astimezone(ZoneInfo(ZONE))


def hour(t) -> str:
    return f"{t.hour % 12 or 12}:{t.minute:02d} {'am' if t.hour < 12 else 'pm'}"


def when_line(raid: dict) -> str:
    from datetime import timedelta
    start = local(raid['starts_at'])
    end = start + timedelta(minutes=raid['duration_minutes'])
    return f"{start:%A} {start.day} {start:%B}{SEP}{hour(start)} to {hour(end)} {start.tzname()}"


def pull_title(pull: dict) -> str:
    return pull['name'] if pull['attempt'] is None else f"{pull['name']}, pull {pull['attempt']}"


def pull_sub(pull: dict) -> str:
    parts = (['Trash'] if pull['attempt'] is None else []) + [clock(pull['duration_ms'])]
    if pull['deaths']:
        parts.append(f"{pull['deaths']} death{'s' if pull['deaths'] != 1 else ''}")
    tally: dict[str, int] = {}
    for d in pull['death_causes'] or []:
        tally[d['killing_blow']['ability']] = tally.get(d['killing_blow']['ability'], 0) + 1
    if tally:
        ordered = sorted(tally.items(), key=lambda kv: (-kv[1], kv[0]))
        parts.append('died to ' + ', '.join(a if n == 1 else f'{a} &times;{n}' for a, n in ordered))
    return SEP.join(parts)


def default_pull(rep: dict) -> dict | None:
    """The pull the side Inspector opens on: the first wipe with deaths, else the first with deaths."""
    open_ = [p for p in rep['pulls'] if p['death_causes']]
    return next((p for p in open_ if p['result'] == 'wipe'), open_[0] if open_ else None)


def attendance_summary(n: dict) -> str:
    parts = [f"{n['came']} of {n['signed']} came"] + ([f"{n['joined']} more joined"] if n['joined'] else [])
    return SEP.join(parts)


def discord_lines(raid: dict) -> list[str]:
    rep, n = raid['report'], raid['night']
    kills = ', '.join(f"{k['name']} {clock(k['kill_ms'])}" for k in rep['kill_order'] if k['killed']) or 'No kills'
    w = rep['headline']['wipes']
    facts = SEP.join([f"{w} wipe{'' if w == 1 else 's'}", f"{n['duration_ms'] // 3_600_000} h {n['duration_ms'] // 60_000 % 60} min",
                      f"{n['came']} of {n['signed']} came", f"{rep['headline']['deaths']} deaths"])
    top = [f"{p['name']} {p['value']} DPS on {p['encounter']}" for p in sorted(rep['parses'], key=lambda p: -p['value'])[:3]]
    loot = [f"{a['item_name']} to {a['name']}" for g in rep['loot'] for a in g['awards']][:5]
    return [f"Report: {raid['title']}", kills, facts, *top, *loot]
