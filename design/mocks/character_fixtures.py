"""The data the character mocks and the character parity spec share (design/specs/2026-11-06-character.md section 3).

Writes web/tests/e2e/fixtures/character/*.json: the GET /v1/characters/... answers, the rating, the /v1/me answers of the
owner states, the search answer of the elsewhere state, and a trimmed BiS and items file for the level 60 priest. The
level 34 hunter reuses the home fixtures (fixtures/home: me-34, the band 30 BiS file and the hunter items). Every worn
item is the band's pick or a listed lower alternative, so every gain is the ranker's own sim number (lib/home/upgrades.ts
rule a) and the mock can read it straight from the file. ws_character.py renders the mocks from these files and
web/tests/e2e/parity-character.spec.ts routes the same files.

  python3 design/mocks/character_fixtures.py && python3 design/mocks/ws_character.py
"""
import json
from pathlib import Path

ROOT = Path(__file__).resolve().parents[2]
BUILD = json.loads((ROOT / 'web/src/data/active-build.json').read_text())['build']
DATA = ROOT / 'data/builds' / BUILD
HOME = ROOT / 'web/tests/e2e/fixtures/home'
OUT = ROOT / 'web/tests/e2e/fixtures/character'
SLOTS = ['head', 'neck', 'shoulder', 'back', 'chest', 'wrist', 'hands', 'waist', 'legs', 'feet', 'finger1', 'finger2',
         'trinket1', 'trinket2', 'main_hand', 'off_hand', 'ranged']
PRIEST_SPEC = 'priest-holy'
CAPTURED = '2026-11-04T19:30:00Z'

BOSSES = [  # Molten Core, best first; HPS for a healer
    ('Lucifron', 12118, 98.4, 1012, '2026-11-03T20:40:00Z'),
    ('Magmadar', 11982, 91.2, 948, '2026-11-03T21:05:00Z'),
    ('Gehennas', 12259, 76.0, 902, '2026-11-03T21:30:00Z'),
    ('Garr', 12057, 63.5, 866, '2026-10-27T21:10:00Z'),
    ('Baron Geddon', 12056, 41.8, 790, '2026-10-27T21:50:00Z'),
    ('Shazzrah', 12264, 37.0, 760, '2026-10-27T22:10:00Z'),
]


def read(path: Path) -> dict:
    return json.loads(path.read_text())


PRINT_WIDTH = 110  # web/.prettierrc


def _flat(value) -> str | None:
    """The value on one line as Prettier prints it, or None when Prettier always breaks it (an object)."""
    if isinstance(value, dict):
        return '{}' if not value else None
    if isinstance(value, list):
        parts = [_flat(item) for item in value]
        return None if None in parts else '[' + ', '.join(parts) + ']'
    return json.dumps(value)


def _is_number(value) -> bool:
    return isinstance(value, (int, float)) and not isinstance(value, bool)


def _pretty(value, indent: int, column: int, tail: int) -> str:
    """Prettier's JSON layout: objects broken one key per line; an array flat when it fits the print width
    (from `column`, with `tail` characters after it on the line), a number array filled line by line, any
    other array one item per line."""
    pad, inner = ' ' * indent, ' ' * (indent + 2)
    if isinstance(value, dict):
        if not value:
            return '{}'
        lines = []
        for index, (key, item) in enumerate(value.items()):
            comma = 1 if index < len(value) - 1 else 0
            head = f'{inner}{json.dumps(key)}: '
            lines.append(head + _pretty(item, indent + 2, len(head), comma) + (',' if comma else ''))
        return '{\n' + '\n'.join(lines) + '\n' + pad + '}'
    if isinstance(value, list):
        if not value:
            return '[]'
        flat = _flat(value)
        if flat is not None and column + len(flat) + tail <= PRINT_WIDTH:
            return flat
        if all(_is_number(item) for item in value):
            rows, row = [], ''
            for index, item in enumerate(value):
                word = json.dumps(item) + (',' if index < len(value) - 1 else '')
                if row and len(inner) + len(row) + 1 + len(word) > PRINT_WIDTH:
                    rows.append(row)
                    row = word
                else:
                    row = f'{row} {word}' if row else word
            rows.append(row)
            return '[\n' + '\n'.join(inner + r for r in rows) + '\n' + pad + ']'
        lines = []
        for index, item in enumerate(value):
            comma = 1 if index < len(value) - 1 else 0
            lines.append(inner + _pretty(item, indent + 2, len(inner), comma) + (',' if comma else ''))
        return '[\n' + '\n'.join(lines) + '\n' + pad + ']'
    return json.dumps(value)


def write(name: str, value) -> None:
    """Writes the fixture in the layout Prettier gives it (2-space indent, a trailing newline, short arrays
    on one line), so `npm run format:check` passes on a fresh regeneration."""
    OUT.mkdir(parents=True, exist_ok=True)
    (OUT / name).write_text(_pretty(value, 0, 0, 0) + '\n')


def chances() -> dict[int, float]:
    out: dict[int, float] = {}
    for src in read(DATA / 'loot.json')['sources']:
        for boss in src.get('bosses') or []:
            for item, chance in (boss.get('item_chances') or {}).items():
                out.setdefault(int(item), chance)
    return out


def worn_of(slot: dict) -> int:
    lower = [a for a in slot.get('alternatives') or [] if a.get('dps_delta', 0) < 0]
    return lower[0]['item_id'] if lower else slot['item_id']


def gear_from(band: dict, items: dict, drop: dict, worn_ids: dict | None = None) -> list[dict]:
    """The API's 17 gear slots: worn (from `worn_ids`, else the band's own rule), the pick, the gap in points."""
    picks = {s['slot']: s for s in band['slots'] if 'item_id' in s}
    rows = []
    for slot in SLOTS:
        pick = picks.get(slot)
        worn_id = (worn_ids or {}).get(slot) if worn_ids is not None else (worn_of(pick) if pick else None)
        it = items.get(worn_id) if worn_id else None
        worn = {'id': worn_id, 'name': it['name'], 'item_level': it.get('item_level')} if it else None
        p = None
        if pick:
            ch = round(drop[pick['item_id']] / 100, 4) if pick['item_id'] in drop and pick.get('source_kind') in ('dungeon', 'raid') else None
            p = {'id': pick['item_id'], 'name': pick['item_name'], 'source': pick['source'], 'chance': ch}
        gap = None if p is None else (0 if worn and worn['id'] == p['id'] else 12.5)
        rows.append({'slot': slot, 'worn': worn, 'pick': p, 'gap': gap})
    return rows


def fights() -> tuple[list[dict], list[dict]]:
    rows = [{'encounter': name, 'encounter_id': eid, 'difficulty': 9, 'metric': 'hps', 'value': hps, 'percentile': pct,
             'spec': 'Holy', 'fought_at': at, 'report_id': f'rpt{eid}', 'fight_index': i + 1, 'execution_score': 0.9 - i * 0.05}
            for i, (name, eid, pct, hps, at) in enumerate(BOSSES)]
    return rows, rows + [{**rows[0], 'value': 930, 'percentile': 88.0, 'fought_at': '2026-10-27T20:30:00Z', 'report_id': 'rpt0000older'}]


def elyra(gear: list[dict]) -> dict:
    best, history = fights()
    return {
        'supporter': True,
        'character': {'name': 'Elyra Duskvale', 'region': 'us', 'ruleset': 'hardcore', 'class': 'Priest',
                      'guild': {'name': 'Cx66 Olympus', 'rank': 'officer'}, 'level': 60, 'race': 'Night Elf', 'realm': 'Doomhowl'},
        'score': 1830, 'tier': 'epic', 'as_of': '2026-11-03',
        'components': {'gear': 410, 'experience': 520, 'parses': 470, 'attendance': 330, 'discipline': 100},
        'experience': None,
        'best': best, 'history': history,
        'builds_seen': [
            {'talent_split': '15/31/5', 'spec': 'Holy', 'first_seen': '2026-11-03T20:40:00Z'},
            {'talent_split': '21/30/0', 'spec': 'Holy', 'first_seen': '2026-10-27T20:30:00Z'},
            {'talent_split': '31/20/0', 'spec': 'Discipline', 'first_seen': '2026-10-20T20:00:00Z'},
        ],
        'score_history': [{'date': '2026-10-20', 'score': 1520, 'tier': 'epic'}, {'date': '2026-10-27', 'score': 1690, 'tier': 'epic'},
                          {'date': '2026-11-03', 'score': 1830, 'tier': 'epic'}],
        'gear': gear, 'gear_gap_total': 96.0, 'gear_spec': PRIEST_SPEC, 'gear_spec_source': 'talents',
    }


def felsnap(gear: list[dict]) -> dict:
    return {
        'supporter': False,
        'character': {'name': 'Felsnap', 'region': 'us', 'ruleset': 'normal', 'class': 'Hunter', 'guild': None, 'level': 34,
                      'race': 'Dwarf', 'realm': 'Skyborne'},
        'score': None, 'tier': None, 'as_of': None, 'components': None, 'experience': None,
        'best': [], 'history': [], 'builds_seen': [], 'score_history': [],
        'gear': gear, 'gear_gap_total': 40.0, 'gear_spec': 'hunter-beast-mastery', 'gear_spec_source': 'talents',
    }


def rating() -> dict:
    comps = [{'name': n, 'score': s, 'excluded': False} for n, s in
             [('activity', 72), ('mechanics', 58), ('cooldowns', 66), ('consumables', 80), ('positioning', 61), ('resources', 54)]]
    latest = {'player_key': 'us/hardcore/elyra-duskvale', 'player_name': 'Elyra Duskvale', 'class': 'priest', 'spec': 'holy',
              'role': 'healer', 'overall': 65, 'overall_uncapped': 65, 'overall_capped': False, 'coverage': 0.9, 'insufficient': False,
              'insufficient_reason': '', 'basis': 'percentile', 'components': comps}
    trend = [{'report_id': f'rpt{b[1]}', 'fight_index': i + 1, 'overall': 60 + i, 'fought_at': b[4]} for i, b in enumerate(BOSSES)]
    return {'player_key': 'us/hardcore/elyra-duskvale', 'sample_size': len(BOSSES), 'trend': trend, 'best_component': 'consumables',
            'worst_component': 'resources', 'latest': latest}


def owner_me(character: dict, premium: bool) -> dict:
    me = read(HOME / 'me-34.json')
    return {**me, 'characters': [character], 'main_character_key': character['key'],
            'entitlements': {'history': premium, 'log_coach': premium}, 'plan': 'premium' if premium else 'free'}


GUILD = {'id': 2, 'name': 'Cx66 Olympus', 'region': 'us', 'ruleset': 'hardcore', 'slug': 'cx66-olympus'}


def feed_row(i: int, kind: str, at: str, tag: str, segments: list[dict], href: str, crest: str | None = 'priest') -> dict:
    return {'id': f'e_{i}', 'kind': kind, 'at': at, 'tag': tag, 'icon': None if crest else 'inv_misc_book_09', 'crest': crest,
            'quality': None, 'href': href, 'date_only': False, 'segments': segments}


def elyra_feed(owner: bool) -> dict:
    name = {'text': 'Elyra Duskvale', 'style': 'name', 'class': 'priest'}
    rows = [
        feed_row(2, 'log_upload', '2026-11-06T16:10:00Z', 'logs', [name, {'text': ' uploaded Molten Core, 10 bosses in 92 min'}], '/reports/rpt12118'),
        feed_row(3, 'ranking_climb', '2026-11-05T21:00:00Z', 'rankings',
                 [{'text': 'Holy, 60: ', 'style': 'strong'}, name, {'text': ' climbed 6 places to '}, {'text': '#4', 'style': 'strong'}],
                 '/character/us/hardcore/elyra-duskvale'),
        feed_row(4, 'boss_kill', '2026-11-03T20:40:00Z', 'logs', [name, {'text': ' downed Lucifron in Molten Core, the first kill on Forever'}], '/reports/rpt12118'),
    ]
    if owner:
        rows.insert(0, feed_row(1, 'guild_sim', '2026-11-06T17:30:00Z', 'your_guild',
                                [name, {'text': ' simmed 912 HPS, #3 of 14 healers in Cx66 Olympus (Holy, 60 raid)'}], '/sim/s_1'))
    return {'rows': rows, 'next_cursor': None}


def home_answer(level: int) -> dict:
    if level < 60:
        return {'import_sim': {'state': 'done', 'dps': 26.2, 'simmed_at': '2026-11-06T12:00:00Z', 'sim_id': 's_2'},
                'spec_week': {'spec': 'hunter-beast-mastery', 'band': 30, 'faction': 'alliance', 'role': 'dps', 'grade': 'B',
                              'last_week_grade': None, 'place': 7, 'of': 23},
                'movers': None, 'guild_sims': None}
    ladder = [('Thalrin', 'paladin', 1001.2, False), ('Mirelle', 'druid', 955.0, False), ('Elyra Duskvale', 'priest', 912.4, True)]
    return {'import_sim': {'state': 'done', 'dps': 912.4, 'simmed_at': '2026-11-06T17:30:00Z', 'sim_id': 's_1'},
            'spec_week': {'spec': 'priest-holy', 'band': 60, 'faction': 'alliance', 'role': 'healer', 'grade': 'A',
                          'last_week_grade': None, 'place': 2, 'of': 9},
            'movers': None,
            'guild_sims': {'guild_id': GUILD['id'], 'guild_name': GUILD['name'], 'role': 'healer', 'band': 60, 'preset': 'raid', 'unit': 'HPS',
                           'of': 14, 'rows': [{'rank': i + 1, 'name': n, 'class': c, 'key': f"us/hardcore/{n.lower().replace(' ', '-')}",
                                               'value': v, 'delta': None, 'simmed_at': '2026-11-06T17:30:00Z', 'you': y}
                                              for i, (n, c, v, y) in enumerate(ladder)],
                           'you': None}}


def elyra_me(gear: list[dict]) -> dict:
    return {'key': 'us/hardcore/elyra-duskvale', 'region': 'us', 'ruleset': 'hardcore', 'name': 'Elyra Duskvale', 'class': 'Priest',
            'spec': 'Holy', 'race': 'Night Elf', 'realm': 'Doomhowl', 'level': 60, 'faction': 'alliance', 'item_level': 62, 'source': 'bnet',
            'guild': {'id': GUILD['id'], 'name': GUILD['name'], 'rank': 'Officer', 'verified': True, 'attendance': {'present': 3, 'nights': 4}},
            'build': {'source': 'addon', 'captured_at': CAPTURED, 'level': 60, 'data_build': BUILD,
                      'gear': {g['slot']: g['worn']['id'] for g in gear if g['worn']},
                      'talents': {'trees': ['0' * 16, '0' * 16, '0' * 16], 'points': [15, 31, 5]}}}


def coach_history() -> dict:
    """Elyra's coach history (GET .../coach/history): four reviews, newest first, one page. SAMPLE (typed): the titles are
    the coach's own Holy output (web/src/fixtures/coach/holy-model.json and holy-rules.json, the Onyxia pull) and its
    `historyNone` rows; the pulls are the fixture's own report."""
    title = read(ROOT / 'web/src/fixtures/coach/holy-model.json')['review']['fixFirst']['title']
    review = lambda fight, boss, kill, source, at, fix: {  # noqa: E731
        'report_id': 'fixture2abcd', 'fight_index': fight, 'player_key': 'us/hardcore/elyra-duskvale', 'spec': 'Holy',
        'boss': boss, 'kill': kill, 'source': source, 'generated_at': at, 'fix_first': {'title': fix} if fix else None}
    return {'reviews': [
        review(4, 'Onyxia', True, 'model', '2026-11-04T21:40:00Z', title),
        review(3, 'Onyxia', False, 'rules', '2026-11-04T21:12:00Z', title),
        review(2, 'Magmadar', True, 'model', '2026-11-03T21:05:00Z', None),
        review(1, 'Lucifron', True, 'rules', '2026-11-03T20:40:00Z', None),
    ], 'next_cursor': None}


def main() -> None:
    drop = chances()
    pbis = read(DATA / 'bis' / f'{PRIEST_SPEC}.json')
    pband = next(b for b in pbis['bands'] if b['band'] == 60 and b['faction'] == 'alliance' and b.get('preset') == 'raid')
    pitems_all = read(DATA / 'items' / 'priest.json')
    pitems = {i['id']: i for i in pitems_all['items']}
    pgear = gear_from(pband, pitems, drop)
    ids = {s['item_id'] for s in pband['slots'] if 'item_id' in s} | {a['item_id'] for s in pband['slots'] for a in s.get('alternatives') or []}
    write(f'bis-{PRIEST_SPEC}.json', {**pbis, 'bands': [pband]})
    write('items-priest.json', {**pitems_all, 'items': [i for i in pitems_all['items'] if i['id'] in ids]})
    write('character-60.json', elyra(pgear))
    write('rating-60.json', rating())
    me60 = owner_me(elyra_me(pgear), premium=True)
    write('me-60.json', {**me60, 'guilds': [{**GUILD, 'rank': 'Officer', 'verified': True, 'verified_count': 14, 'faction': 'alliance'}]})
    write('feed-60.json', elyra_feed(owner=False))
    write('feed-60-owner.json', elyra_feed(owner=True))
    write('coach-history-60.json', coach_history())
    write('home-60.json', home_answer(60))
    write('home-34.json', home_answer(34))

    hbis = read(HOME / 'bis-hunter-beast-mastery.json')
    hitems = {i['id']: i for i in read(HOME / 'items-hunter.json')['items']}
    me34 = read(HOME / 'me-34.json')
    worn34 = me34['characters'][0]['build']['gear']
    write('character-34.json', felsnap(gear_from(hbis['bands'][0], hitems, drop, worn34)))
    write('character-34-unread.json', felsnap(gear_from(hbis['bands'][0], hitems, drop, {})))
    fel = me34['characters'][0]
    write('me-34.json', owner_me(fel, premium=False))
    write('me-34-unread.json', owner_me({**fel, 'build': {**fel['build'], 'gear': {}, 'talents': {**fel['build']['talents'], 'points': [0, 0, 0]}}}, premium=False))
    write('search-thalrin.json', {'rows': [
        {'key': 'us/normal/thalrin', 'region': 'us', 'ruleset': 'normal', 'name': 'Thalrin', 'class': 'Warrior', 'spec': 'Fury', 'level': 60},
        {'key': 'eu/hardcore/thalrin', 'region': 'eu', 'ruleset': 'hardcore', 'name': 'Thalrin', 'class': 'Mage', 'spec': 'Frost', 'level': 42},
    ]})


if __name__ == '__main__':
    main()
