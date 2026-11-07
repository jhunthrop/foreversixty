"""The raid night report's data (design/parity.md, "The guild states"; state `report`): Onyxia, Tue 8 Dec, read by an officer
(Dunmaro). One source for both sides: ws_guild2.raid_report draws the mock from `raid_view()` and guild_fixtures.py writes the
same answers into web/tests/e2e/fixtures/guild-parity/report-*.json, which web/tests/e2e/support/guild-parity/report.ts routes.
Shapes: web/src/lib/guild/api.ts (GuildRaidView, GuildHome), attendance-contract.ts (RaidReport), rankings/api.ts (GuildPage),
account/api.ts (Me). Roster and sign-up states come from guild_data.R.
"""
from guild_data import GUILD, GUILD_ID, R, REGION, RULESET, SLUG, ZONE

VIEWER = 'Dunmaro'
RAID_ID = 41
ZONE_NAME = "Onyxia's Lair"
TITLE = ZONE_NAME
# Tue 8 Dec, 8:00 pm CST (the guild raids Tue and Thu at 8:00 pm).
STARTS_AT = '2026-12-09T02:00:00Z'
UPLOADED_AT = '2026-12-09T05:10:00Z'
AWARDED_AT = '2026-12-09T04:30:00Z'
NIGHT_MS = 90 * 60_000
ONYXIA = 1084
# Not in the log: everyone else on the roster came.
AWAY = ('Voidcall', 'Starnettle')
ROLE = {'Tank': 'tank', 'Healer': 'healer', 'Melee': 'dps', 'Ranged': 'dps'}
GROUP = {'Tank': 'tank', 'Healer': 'healer', 'Melee': 'melee', 'Ranged': 'ranged'}
SPELL = {'Flame Breath': 18435, 'Deep Breath': 17086, 'Tail Sweep': 15847}

# fight index, name, boss pull number (None on trash), result, duration ms, deaths as (who, killing ability)
PULLS = [
    (1, 'Onyxian Whelp pack', None, 'clear', 60_000, []),
    (2, 'Onyxia', 1, 'wipe', 240_000, [('Emberfall', 'Deep Breath'), ('Skytalon', 'Deep Breath'), ('Hexmarrow', 'Deep Breath'),
                                       ('Shiverknife', 'Tail Sweep'), ('Blacktide', 'Tail Sweep'), ('Seraphine', 'Flame Breath')]),
    (3, 'Onyxian Whelp pack', None, 'clear', 55_000, []),
    (4, 'Onyxia', 2, 'kill', 225_000, [('Groundshock', 'Flame Breath')]),
]
KILL_FIGHT = 4
# name, DPS on the kill, the fight's rating
PARSES = [('Frostquill', 812, 96), ('Doomwhisper', 744, 88), ('Dunmaro', 618, 61)]
# item id, name, quality, icon, winner
LOOT = [(18813, 'Ring of Binding', 4, 'inv_jewelry_ring_13', 'Doomwhisper')]


def key(name: str) -> str:
    return f'{REGION}/{RULESET}/{name.lower()}'


def row(name: str) -> tuple:
    return next(r for r in R if r[0] == name)


def who(name: str) -> dict:
    r = row(name)
    return {'character_key': key(name), 'name': name, 'class': r[1]}


def signed_state(state: str) -> str | None:
    return None if state == 'none' else state


def spec_key(r: tuple) -> str:
    return f"{r[1]}-{r[2].lower().replace(' ', '-')}"


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
            {'item_id': i, 'item_name': n, 'quality': q, 'icon': ic, **who(w), 'by_name': VIEWER, 'at': AWARDED_AT} for i, n, q, ic, w in LOOT]}],
        'moves': {'previous': None, 'rows': []},
        'posted_at': None,
    }


def attendance() -> list[dict]:
    out = []
    for r in R:
        name, _, spec, role, _, state = r
        here = name not in AWAY
        outcome = 'present' if here else 'absent'
        out.append({**who(name), 'spec': spec, 'signed': signed_state(state), 'role': ROLE[role], 'outcome': outcome, 'came_as': None,
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
    return [{**who(r[0]), 'spec': r[2], 'role': ROLE[r[3]], 'state': r[5], 'note': '', 'own': r[0] == VIEWER, 'group': GROUP[r[3]],
             'spec_key': spec_key(r)} for r in R if r[5] != 'none']


def counts(rows: list[dict]) -> dict:
    live = [s for s in rows if s['state'] in ('in', 'late')]
    classes: dict[str, int] = {}
    for s in live:
        classes[s['class']] = classes.get(s['class'], 0) + 1
    by = lambda field, value: sum(1 for s in live if s[field] == value)  # noqa: E731
    return {'in': sum(1 for s in rows if s['state'] == 'in'), 'late': sum(1 for s in rows if s['state'] == 'late'),
            'bench': sum(1 for s in rows if s['state'] == 'bench'), 'out': 0, 'tanks': by('role', 'tank'), 'healers': by('role', 'healer'),
            'dps': by('role', 'dps'), 'melee': by('group', 'melee'), 'ranged': by('group', 'ranged'), 'classes': classes}


def raid_view() -> dict:
    rows = signups()
    viewer = row(VIEWER)
    return {
        'id': RAID_ID, 'zone': 'raid:onyxias-lair', 'zone_name': ZONE_NAME, 'zone_kind': 'raid', 'key_item': None, 'title': TITLE,
        'starts_at': STARTS_AT, 'duration_minutes': 180, 'status': 'done', 'size': 40, 'guild_timezone': ZONE, 'notes_html': '', 'notes': '',
        'counts': counts(rows), 'signups': rows,
        'not_signed': [{**who(r[0]), 'spec': r[2], 'group': GROUP[r[3]], 'spec_key': spec_key(r)} for r in R if r[5] == 'none'],
        'viewer': {'characters': [{**who(VIEWER), 'spec': viewer[2], 'verified': True, 'state': viewer[5], 'role': ROLE[viewer[3]],
                                   'raid_spec': spec_key(viewer), 'spec_key': spec_key(viewer), 'default_spec': spec_key(viewer)}]},
        'discord': {'available': True, 'connected': True, 'post_default': True, 'remind_default': True, 'channel_name': 'raids'},
        'discord_post': None, 'cover': None, 'comp_target': None, 'assignments': None,
        'pre_done_status': 'locked',
        'match': {'report_id': 'onyxia8dec', 'zone_name': ZONE_NAME, 'uploaded_at': UPLOADED_AT, 'uploaded_by_name': VIEWER,
                  'pulls': len(PULLS), 'kills': headline()['kills'], 'duration_ms': NIGHT_MS, 'how': 'auto'},
        'night': night(), 'attendance': attendance(), 'report': report(),
    }


def guild_summary() -> dict:
    return {'id': GUILD_ID, 'name': GUILD, 'region': REGION, 'ruleset': RULESET, 'faction': 'alliance', 'crest_url': None}


def guild_page() -> dict:
    """GET /v1/guilds/{region}/{ruleset}/{slug}: the public read that resolves the guild's id."""
    return {'guild': guild_summary(), 'progression': [], 'roster_best': [], 'reports': []}


def guild_home() -> dict:
    """GET /v1/guilds/{id}/home as the officer reads it; the raid page reads its viewer and claim."""
    return {'guild': guild_summary(),
            'viewer': {'role': 'officer', 'character_key': key(VIEWER), 'verified': True, 'officer_tools': True, 'raid_guild_tools': True},
            'claim': {'state': 'claimed', 'since': '2026-11-12T00:00:00Z', 'frozen': False, 'yours': True,
                      'claimed_by': {'name': VIEWER, 'class': row(VIEWER)[1], 'character_key': key(VIEWER)}},
            'reports': [], 'roster': [], 'pending': []}


def me() -> dict:
    """GET /v1/me: the officer signed in with Dunmaro, the mock's `who`."""
    r = row(VIEWER)
    return {
        'user': {'id': 7, 'battletag': 'Fixture#1894', 'email': None, 'role': 'user', 'anonymize': False},
        'characters': [{'key': key(VIEWER), 'region': REGION, 'ruleset': RULESET, 'name': VIEWER, 'class': r[1].title(), 'spec': r[2],
                        'level': 60, 'realm': 'Skyborne', 'faction': 'alliance',
                        'guild': {'id': GUILD_ID, 'name': GUILD, 'rank': 'Officer', 'verified': True}}],
        'guilds': [{'id': GUILD_ID, 'region': REGION, 'ruleset': RULESET, 'name': GUILD, 'slug': SLUG, 'rank': 'officer',
                    'verified': True, 'can': ['approve', 'remove', 'settings', 'award'], 'faction': 'alliance'}],
        'main_character_key': key(VIEWER),
    }


FIXTURES = {'report-raid.json': raid_view, 'report-guild.json': guild_page, 'report-home.json': guild_home, 'report-me.json': me}


def write_report_fixtures(write) -> list:
    return [write(name, build()) for name, build in FIXTURES.items()]


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
