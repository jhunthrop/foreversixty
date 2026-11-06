"""The data the home mocks and the home parity spec share (design/specs/2026-11-06-home.md section 0).

Writes web/tests/e2e/fixtures/home/*.json: the API answers (/v1/me for each signed-in state, sims, reports, the guild
home) and trimmed BiS and item files read from the active build. ws_home.py renders the mocks from these files and
web/tests/e2e/parity-home.spec.ts routes the same files, so both sides of the gate show the same data. Each sample
character wears its band's pick in every slot except where the pick lists a lower alternative: there it wears the
first such alternative, so every gain is the ranker's own sim-verified number (lib/home/upgrades.ts rule b).

  python3 design/mocks/home_fixtures.py
"""
import json
from pathlib import Path

ROOT = Path(__file__).resolve().parents[2]
BUILD = json.loads((ROOT / 'web/src/data/active-build.json').read_text())['build']
DATA = ROOT / 'data/builds' / BUILD
OUT = ROOT / 'web/tests/e2e/fixtures/home'
WEAPONS = {'main_hand', 'off_hand', 'ranged'}

# The two sample characters (spec section 1, H1): a level 34 leveler with no guild and a level 60 raider in one.
SAMPLES = {
    '34': {'name': 'Felsnap', 'class': 'Hunter', 'class_slug': 'hunter', 'spec': 'Beast Mastery', 'spec_key': 'hunter-beast-mastery',
           'level': 34, 'faction': 'alliance', 'race': 'Dwarf', 'item_level': 31, 'points': [23, 0, 0], 'preset': 'bare'},
    '60': {'name': 'Thandrel', 'class': 'Warrior', 'class_slug': 'warrior', 'spec': 'Fury', 'spec_key': 'warrior-fury',
           'level': 60, 'faction': 'alliance', 'race': 'Human', 'item_level': 61, 'points': [17, 34, 0], 'preset': 'raid'},
}
GUILD = {'id': 2, 'name': 'Olympus XXVII', 'slug': 'olympus-xxvii', 'region': 'us', 'ruleset': 'normal'}
CAPTURED = '2026-11-04T19:30:00Z'
IMPORTED = '2026-11-06T09:00:00Z'


def band_of(bis: dict, band: int, faction: str, preset: str) -> dict:
    return next(b for b in bis['bands'] if b['band'] == band and b['faction'] == faction and b.get('preset') == preset)


def worn_gear(band: dict) -> dict:
    gear = {}
    for slot in band['slots']:
        if 'item_id' not in slot:
            continue
        lower = [a for a in slot.get('alternatives') or [] if a.get('dps_delta', 0) < 0]
        gear[slot['slot']] = lower[0]['item_id'] if lower else slot['item_id']
    return gear


def trimmed_items(class_slug: str, ids: set[int]) -> dict:
    items = json.loads((DATA / 'items' / f'{class_slug}.json').read_text())
    return {**items, 'items': [i for i in items['items'] if i['id'] in ids]}


def character(sample: dict, gear: dict, guild: bool) -> dict:
    key = f"us/normal/{sample['name'].lower()}"
    talents_points = sample['points']  # the spec is read from the tree with the most points (lib/character-spec.ts)
    row = {
        'key': key, 'region': 'us', 'ruleset': 'normal', 'name': sample['name'], 'class': sample['class'], 'spec': sample['spec'],
        'race': sample['race'], 'realm': 'Skyborne', 'level': sample['level'], 'faction': sample['faction'],
        'item_level': sample['item_level'], 'source': 'bnet',
        'build': {'source': 'addon', 'captured_at': CAPTURED, 'level': sample['level'], 'data_build': BUILD, 'gear': gear,
                  'talents': {'trees': ['0' * 16, '0' * 17, '0' * 18], 'points': talents_points}},
    }
    if guild:
        row['guild'] = {'id': GUILD['id'], 'name': GUILD['name'], 'rank': 'Officer', 'verified': True, 'attendance': {'present': 3, 'nights': 4}}
    return row


def me(characters: list[dict], guilds: list[dict]) -> dict:
    user = {'id': 7, 'battletag': 'Hunthrop#1894', 'email': None, 'role': 'user', 'anonymize': False}
    return {'user': user, 'characters': characters, 'guilds': guilds, 'bnet_imported_at': IMPORTED,
            'main_character_key': characters[0]['key'] if characters else None}


def write(name: str, value: dict) -> None:
    OUT.mkdir(parents=True, exist_ok=True)
    (OUT / name).write_text(json.dumps(value, indent=1) + '\n')


def main() -> None:
    for state, sample in SAMPLES.items():
        bis = json.loads((DATA / 'bis' / f"{sample['spec_key']}.json").read_text())
        band = band_of(bis, 60 if sample['level'] >= 60 else sample['level'] // 10 * 10, sample['faction'], sample['preset'])
        gear = worn_gear(band)
        ids = set(gear.values()) | {s['item_id'] for s in band['slots'] if 'item_id' in s}
        write(f"bis-{sample['spec_key']}.json", {**bis, 'bands': [band]})
        write(f"items-{sample['class_slug']}.json", trimmed_items(sample['class_slug'], ids))
        guilds = [{'id': GUILD['id'], 'region': 'us', 'ruleset': 'normal', 'name': GUILD['name'], 'slug': GUILD['slug'], 'rank': 'Officer',
                   'verified': True, 'verified_count': 18, 'faction': 'alliance'}] if state == '60' else []
        write(f'me-{state}.json', me([character(sample, gear, state == '60')], guilds))
    write('me-new.json', me([], []))
    write('sims-34.json', {'rows': [{'sim_id': 'k2h9f', 'spec': 'hunter-beast-mastery', 'dps': 61.2, 'engine_version': 'fixture',
                                     'created_at': '2026-11-04T20:10:00Z', 'title': 'Beast Mastery Hunter, 30 to 39', 'kind': 'run',
                                     'band': 30, 'faction': 'alliance', 'preset': 'bare'}], 'total': 1, 'page': 1, 'per_page': 20})
    write('sims-60.json', {'rows': [{'sim_id': 'p8d3q', 'spec': 'warrior-fury', 'dps': 948.6, 'engine_version': 'fixture',
                                     'created_at': '2026-11-05T22:40:00Z', 'title': 'Fury Warrior, raid-ready', 'kind': 'run',
                                     'band': 60, 'faction': 'alliance', 'preset': 'raid'}], 'total': 1, 'page': 1, 'per_page': 20})
    write('reports-34.json', {'rows': [], 'total': 0, 'page': 1, 'per_page': 100})
    report = {'status': 'done', 'visibility': 'guild', 'logging_character': {'key': 'us/normal/thandrel', 'name': 'Thandrel'}}
    write('reports-60.json', {'rows': [
        {**report, 'id': 'r-3', 'title': 'Upper Blackrock Spire', 'zone': 'Blackrock Spire', 'created_at': '2026-11-05T21:00:00Z', 'fight_count': 14, 'kill_count': 5, 'boss_count': 5},
        {**report, 'id': 'r-2', 'title': 'Scholomance', 'zone': 'Scholomance', 'created_at': '2026-11-03T20:30:00Z', 'fight_count': 11, 'kill_count': 6, 'boss_count': 6},
        {**report, 'id': 'r-1', 'title': 'Stratholme, living side', 'zone': 'Stratholme', 'created_at': '2026-11-02T19:00:00Z', 'fight_count': 9, 'kill_count': 4, 'boss_count': 4},
    ], 'total': 3, 'page': 1, 'per_page': 100})
    write('guild-home-60.json', {
        'guild': {'id': GUILD['id'], 'name': GUILD['name'], 'region': 'us', 'ruleset': 'normal', 'slug': GUILD['slug'], 'faction': 'alliance'},
        'claim': {'state': 'claimed'}, 'reports': [], 'roster': [], 'pending': [],
        'standing': {'spec': 'Fury', 'class': 'Warrior', 'same_spec_count': 4, 'rank_by_item_level': 2, 'item_level': 61,
                     'needs_before_next_raid': ['Enchant your chest and boots', 'Bring a Flask of the Titans or Elixir of the Mongoose'],
                     'needs_raid': {'title': "Onyxia's Lair", 'starts_at': '2026-12-10T02:00:00Z', 'character_key': 'us/normal/thandrel'}},
        'next_raid': {'id': 41, 'zone': 'onyxia', 'zone_name': "Onyxia's Lair", 'zone_kind': 'raid', 'title': "Onyxia's Lair",
                      'starts_at': '2026-12-10T02:00:00Z', 'status': 'planned', 'size': 40, 'guild_timezone': 'America/Chicago',
                      'counts': {'in': 18, 'late': 3, 'tanks': 3, 'healers': 6, 'dps': 12, 'melee': 7, 'ranged': 5},
                      'own': {'character_key': 'us/normal/thandrel', 'state': 'late', 'role': 'dps'}},
    })


if __name__ == '__main__':
    main()
