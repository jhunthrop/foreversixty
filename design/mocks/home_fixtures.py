"""The data the home mocks and the home parity spec share (design/specs/2026-11-07-home-b.md section 0).

Writes web/tests/e2e/fixtures/home/*.json: the API answers (/v1/me for each signed-in state, /v1/feed, /v1/home, the
guild home) and trimmed BiS and item files read from the active build. ws_home.py renders the mocks from these files and
web/tests/e2e/parity-home.spec.ts routes the same files, so both sides of the gate show the same data. Each sample
character wears its band's pick in every slot except where the pick lists a lower alternative: there it wears the
first such alternative, so every gain is the ranker's own sim-verified number (lib/home/upgrades.ts rule b).

  python3 design/mocks/home_fixtures.py
"""
import json
from datetime import datetime, timedelta, timezone
from pathlib import Path

ROOT = Path(__file__).resolve().parents[2]
BUILD = json.loads((ROOT / 'web/src/data/active-build.json').read_text())['build']
DATA = ROOT / 'data/builds' / BUILD
OUT = ROOT / 'web/tests/e2e/fixtures/home'
WEAPONS = {'main_hand', 'off_hand', 'ranged'}

# The sample characters (spec section 1, H1): a level 34 leveler with no guild and a level 60 raider in one; and a
# level 41 Protection warrior (2026-11-11), whose upgrades count in the ranker's tank score, never DPS.
SAMPLES = {
    '34': {'name': 'Felsnap', 'class': 'Hunter', 'class_slug': 'hunter', 'spec': 'Beast Mastery', 'spec_key': 'hunter-beast-mastery',
           'level': 34, 'faction': 'alliance', 'race': 'Dwarf', 'item_level': 31, 'points': [23, 0, 0], 'preset': 'bare'},
    '60': {'name': 'Thandrel', 'class': 'Warrior', 'class_slug': 'warrior', 'spec': 'Fury', 'spec_key': 'warrior-fury',
           'level': 60, 'faction': 'alliance', 'race': 'Human', 'item_level': 61, 'points': [17, 34, 0], 'preset': 'raid'},
    '41': {'name': 'Bulwark', 'class': 'Warrior', 'class_slug': 'warrior', 'spec': 'Protection', 'spec_key': 'warrior-protection',
           'level': 41, 'faction': 'alliance', 'race': 'Human', 'item_level': 38, 'points': [0, 0, 32], 'preset': 'bare'},
}
GUILD = {'id': 2, 'name': 'Olympus XXVII', 'slug': 'olympus-xxvii', 'region': 'us', 'ruleset': 'normal'}
CAPTURED = '2026-11-04T19:30:00Z'
# Slots a sample wears its band's own pick in: the leveler's first waist alternative is Highlander's Chain Girdle, leather in
# this client but a mail name a level 34 hunter would never read as legal (wow-player, 2026-11-09).
WORN_PICK = {'34': {'waist': 252521}}
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


def character(sample: dict, gear: dict, guild: bool, attendance: bool = True) -> dict:
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
        row['guild'] = {'id': GUILD['id'], 'name': GUILD['name'], 'rank': 'Officer', 'verified': True}
        if attendance:
            row['guild']['attendance'] = {'present': 3, 'nights': 4}
    return row


def me(characters: list[dict], guilds: list[dict]) -> dict:
    user = {'id': 7, 'battletag': 'Hunthrop#1894', 'email': None, 'role': 'user', 'anonymize': False}
    return {'user': user, 'characters': characters, 'guilds': guilds, 'bnet_imported_at': IMPORTED,
            'main_character_key': characters[0]['key'] if characters else None}


def guild_entry() -> dict:
    return {'id': GUILD['id'], 'region': 'us', 'ruleset': 'normal', 'name': GUILD['name'], 'slug': GUILD['slug'], 'rank': 'Officer',
            'verified': True, 'verified_count': 18, 'faction': 'alliance'}


def write(name: str, value: dict) -> None:
    OUT.mkdir(parents=True, exist_ok=True)
    (OUT / name).write_text(json.dumps(value, indent=1) + '\n')


def main() -> None:
    # One item file per class, carrying every sample of that class (the two warriors share items-warrior.json).
    item_ids: dict[str, set[int]] = {}
    for state, sample in SAMPLES.items():
        bis = json.loads((DATA / 'bis' / f"{sample['spec_key']}.json").read_text())
        band = band_of(bis, 60 if sample['level'] >= 60 else sample['level'] // 10 * 10, sample['faction'], sample['preset'])
        gear = {**worn_gear(band), **WORN_PICK.get(state, {})}
        ids = set(gear.values()) | {s['item_id'] for s in band['slots'] if 'item_id' in s}
        item_ids[sample['class_slug']] = item_ids.get(sample['class_slug'], set()) | ids
        write(f"bis-{sample['spec_key']}.json", {**bis, 'bands': [band]})
        guilds = [guild_entry()] if state == '60' else []
        write(f'me-{state}.json', me([character(sample, gear, state == '60')], guilds))
        if state == '34':
            # The same leveler in the guild (the owner's own state, 2026-11-09): no raid planned yet, so no
            # raid nights attended, and several sims.
            write('me-34-guild.json', me([character(sample, gear, True, attendance=False)], [guild_entry()]))
    for class_slug, ids in item_ids.items():
        write(f'items-{class_slug}.json', trimmed_items(class_slug, ids))
    write('me-new.json', me([], []))
    guild_home = {
        'guild': {'id': GUILD['id'], 'name': GUILD['name'], 'region': 'us', 'ruleset': 'normal', 'slug': GUILD['slug'], 'faction': 'alliance'},
        'claim': {'state': 'claimed'}, 'reports': [], 'roster': [], 'pending': [],
        'summary': {'raiders': 14, 'waiting_for_approval': 0, 'below_rating_floor': 0, 'named_encounters_down': 0, 'pulls_this_tier': 0,
                    'tier': {'name': 'Tier 1', 'raids': ['Barrow Deeps', 'Hyjal Summit', "Onyxia's Lair"], 'named_encounters': 9, 'down': 0,
                             'pulls_total': 0},
                    'updated_at': '2026-11-06T08:00:00Z'},
        'standing': {'spec': 'Fury', 'class': 'Warrior', 'same_spec_count': 4, 'rank_by_item_level': 2, 'item_level': 61,
                     'needs_before_next_raid': ['Enchant your chest and boots', 'Bring a Flask of the Titans or Elixir of the Mongoose'],
                     'needs_raid': {'title': "Onyxia's Lair", 'starts_at': '2026-12-10T02:00:00Z', 'character_key': 'us/normal/thandrel'}},
        'next_raid': {'id': 41, 'zone': 'onyxia', 'zone_name': "Onyxia's Lair", 'zone_kind': 'raid', 'title': "Onyxia's Lair",
                      'starts_at': '2026-12-10T02:00:00Z', 'status': 'planned', 'size': 40, 'guild_timezone': 'America/Chicago',
                      'counts': {'in': 18, 'late': 3, 'tanks': 3, 'healers': 6, 'dps': 12, 'melee': 7, 'ranged': 5},
                      'own': {'character_key': 'us/normal/thandrel', 'state': 'late', 'role': 'dps'}},
    }
    write('guild-home-60.json', guild_home)
    # The same guild with no raid planned (the plaque falls back to the first raids).
    write('guild-home-none.json', {k: v for k, v in guild_home.items() if k != 'next_raid'})
    standing_34 = {'spec': 'Beast Mastery', 'class': 'Hunter', 'same_spec_count': 3, 'rank_by_item_level': 2, 'item_level': 31,
                   'needs_before_next_raid': [], 'needs_raid': None}
    write('guild-home-34.json', {**{k: v for k, v in guild_home.items() if k != 'next_raid'}, 'standing': standing_34})
    write_feeds()
    write('home-34.json', {
        'import_sim': {'state': 'done', 'dps': 58.4, 'simmed_at': '2026-11-06T09:04:00Z', 'sim_id': 'imp34'},
        'spec_week': {'spec': 'hunter-beast-mastery', 'band': 30, 'faction': 'alliance', 'role': 'dps', 'grade': 'A',
                      'last_week_grade': 'B', 'place': 4, 'of': 23},
        'movers': {'spec': 'hunter-beast-mastery', 'band': 30, 'baseline': True, 'rows': [
            mover(3, 'Pawsworth', 'hunter', 6), mover(5, 'Ellowen', 'hunter', 4), mover(9, 'Gritfang', 'hunter', 11)]},
        'guild_sims': None,
    })
    # The tank's home: a sim (its figure the tank's own), no board place yet, no guild.
    write('home-41.json', {
        'import_sim': {'state': 'done', 'dps': 83.5, 'simmed_at': '2026-11-06T09:04:00Z', 'sim_id': 'imp41'},
        'spec_week': None, 'movers': None, 'guild_sims': None,
    })
    write('home-34-guild.json', {
        'import_sim': {'state': 'done', 'dps': 58.4, 'simmed_at': '2026-11-06T09:04:00Z', 'sim_id': 'imp34'},
        'spec_week': {'spec': 'hunter-beast-mastery', 'band': 30, 'faction': 'alliance', 'role': 'dps', 'grade': 'A',
                      'last_week_grade': 'B', 'place': 4, 'of': 23},
        'movers': {'spec': 'hunter-beast-mastery', 'band': 30, 'baseline': True, 'rows': [
            mover(3, 'Pawsworth', 'hunter', 6), mover(5, 'Ellowen', 'hunter', 4), mover(9, 'Gritfang', 'hunter', 11)]},
        'guild_sims': {'guild_id': GUILD['id'], 'guild_name': GUILD['name'], 'role': 'dps', 'band': 30, 'preset': 'bare',
                       'unit': 'DPS', 'of': 6, 'rows': [
                           guild_sim(1, 'Vessa', 'rogue', 71.2), guild_sim(2, 'Kaelra', 'mage', 66.9),
                           guild_sim(3, 'Morrigan', 'warlock', 61.5), guild_sim(4, 'Felsnap', 'hunter', 58.4, you=True),
                           guild_sim(5, 'Brindle', 'warrior', 55.0)],
                       'you': guild_sim(4, 'Felsnap', 'hunter', 58.4, you=True)},
    })
    write('home-60.json', {
        'import_sim': {'state': 'done', 'dps': 948.6, 'simmed_at': '2026-11-06T09:04:00Z', 'sim_id': 'imp60'},
        'spec_week': {'spec': 'warrior-fury', 'band': 60, 'faction': 'alliance', 'role': 'dps', 'grade': 'S',
                      'last_week_grade': None, 'place': 2, 'of': 23},
        'movers': {'spec': 'warrior-fury', 'band': 60, 'baseline': False, 'rows': [
            mover(1, 'Brakka', 'warrior', None), mover(2, 'Thalrin', 'warrior', None), mover(3, 'Zugzug', 'warrior', None)]},
        'guild_sims': {'guild_id': GUILD['id'], 'guild_name': GUILD['name'], 'role': 'dps', 'band': 60, 'preset': 'raid',
                       'unit': 'DPS', 'of': 14, 'rows': [
                           guild_sim(1, 'Kaelra', 'mage', 1012.3), guild_sim(2, 'Thalrin', 'warrior', 987.1),
                           guild_sim(3, 'Vessa', 'rogue', 975.4), guild_sim(4, 'Morrigan', 'warlock', 961.0),
                           guild_sim(5, 'Arrowind', 'hunter', 955.8)],
                       'you': {**guild_sim(7, 'Thandrel', 'warrior', 948.6), 'you': True}},
    })


def mover(rank: int, name: str, cls: str, delta: int | None) -> dict:
    return {'rank': rank, 'name': name, 'class': cls, 'key': f'us/normal/{name.lower()}', 'delta': delta}


def guild_sim(rank: int, name: str, cls: str, value: float, you: bool = False) -> dict:
    return {'rank': rank, 'name': name, 'class': cls, 'key': f'us/normal/{name.lower()}', 'value': value, 'delta': None,
            'simmed_at': '2026-11-06T08:00:00Z', 'you': you}


def sim_row(i: str, minutes: int, name: str, cls: str, figure: str, spec: str, band: str, place: str = '') -> dict:
    """A guild sim in the API's own words (api/internal/activity guildSimEvent): the guild is its own strong segment."""
    return feed_row(i, 'guild_sim', minutes, 'your_guild', [seg(name, 'name', **{'class': cls}), seg(f' simmed {figure} DPS{place} in '),
                    seg(GUILD['name'], 'strong'), seg(f' ({spec}, {band})')], crest=cls, href=f'/sim/{i}')


def seg(text: str, style: str = '', **extra) -> dict:
    return {'text': text, **({'style': style} if style else {}), **extra}


def feed_row(i: str, kind: str, minutes: int, tag: str, segments: list[dict], icon: str | None = None, crest: str | None = None,
             quality: int | None = None, href: str = '/reports/r1') -> dict:
    at = (datetime(2026, 11, 6, 18, tzinfo=timezone.utc) - timedelta(minutes=minutes)).strftime('%Y-%m-%dT%H:%M:%SZ')
    return {'id': i, 'kind': kind, 'at': at, 'tag': tag, 'icon': icon, 'crest': crest, 'quality': quality, 'href': href,
            'date_only': False, 'segments': segments}


# The public rows every state shares (P1, P2, P4, P5): uploads, a ranking climb, a BiS change.
PUBLIC = [
    feed_row('e1', 'log_upload', 4, 'logs', [seg('Nightfall', 'name', **{'class': 'warlock'}), seg(' uploaded a full Shadowfang Keep, 8 bosses in 41 min')], crest='warlock'),
    feed_row('e2', 'ranking_climb', 17, 'rankings', [seg('Fury Warrior, 20 to 29:', 'strong'), seg(' '), seg('Grukk', 'name', **{'class': 'warrior'}), seg(' climbed 6 places to #3')], crest='warrior', href='/rankings'),
    feed_row('e3', 'bis_change', 62, 'best_in_slot', [seg('Best in slot changed', 'strong'), seg(' for Fury Warrior 20 to 29: '), seg('Blackened Defias Armor', 'item', quality=3, item_id=10399), seg(' is the new chest')], icon='inv_chest_leather_08', quality=3, href='/bis/warrior/fury'),
    feed_row('e4', 'log_upload', 95, 'logs', [seg('Ellowen', 'name', **{'class': 'hunter'}), seg(' uploaded Razorfen Kraul, 5 bosses in 52 min')], crest='hunter'),
    feed_row('e5', 'ranking_top', 140, 'rankings', [seg('Brakka', 'name', **{'class': 'warrior'}), seg(' took #1 for Arms Warrior, 30 to 39')], crest='warrior', href='/rankings'),
    feed_row('e6', 'log_upload', 205, 'logs', [seg('Saelis', 'name', **{'class': 'priest'}), seg(' uploaded The Deadmines, 6 bosses in 38 min')], crest='priest'),
    feed_row('e7', 'bis_change', 300, 'best_in_slot', [seg('Best in slot refreshed', 'strong'), seg(' for Fire Mage 40 to 49: 3 new picks')], icon='spell_fire_flamebolt', href='/bis/mage/fire'),
]


def write_feeds() -> None:
    page = lambda rows: {'rows': rows, 'next_cursor': None}
    write('feed-out.json', page(PUBLIC))
    hunter = [
        feed_row('h1', 'ranking_climb', 9, 'your_spec', [seg('Beast Mastery Hunter, 30 to 39:', 'strong'), seg(' '), seg('Pawsworth', 'name', **{'class': 'hunter'}), seg(' climbed 6 places to #3')], crest='hunter', href='/rankings'),
        # One changed pick names its item, as the API writes it (api/internal/activity bisdiff.go).
        feed_row('h2', 'bis_change', 48, 'your_spec', [seg('Best in slot changed', 'strong'), seg(' for Beast Mastery Hunter 30 to 39: '), seg('Golemsight Long Gun', 'item', quality=3, item_id=273029), seg(' is the new ranged')], icon='inv_musket_02', quality=3, href='/bis/hunter/beast-mastery'),
    ]
    write('feed-34.json', page(hunter + PUBLIC[:5]))
    guild = [
        sim_row('g1', 12, 'Thalrin', 'warrior', '987', 'Fury Warrior', '60 raid', ', #2 of 14 DPS'),
        feed_row('g2', 'guild_raid', 70, 'your_guild', [seg('Olympus XXVII', 'strong'), seg(" posted Onyxia's Lair, Wed 9 Dec: 18 in, 3 late")], icon='inv_misc_head_dragon_black', href='/guild/us/normal/olympus-xxvii/raids/41'),
        sim_row('g3', 180, 'Kaelra', 'mage', '1012', 'Fire Mage', '60 raid', ', #1 of 14 DPS'),
    ]
    write('feed-60.json', page(guild + PUBLIC[:4]))
    # The guild with no raid planned posts no raid.
    write('feed-guild-none.json', page([g for g in guild if g['kind'] != 'guild_raid'] + PUBLIC[:5]))
    write('feed-empty.json', page([]))
    # The owner's 2026-11-09 feed: the API puts the guild's own rows first within each hour, so a busy guild fills the
    # page with sims, the same member twice.
    sims = [
        sim_row('s1', 3, 'Vessa', 'rogue', '71.2', 'Combat Rogue', '30 to 39'),
        sim_row('s2', 6, 'Kaelra', 'mage', '66.9', 'Frost Mage', '30 to 39'),
        sim_row('s3', 22, 'Vessa', 'rogue', '68.0', 'Combat Rogue', '30 to 39'),
        sim_row('s4', 31, 'Morrigan', 'warlock', '61.5', 'Affliction Warlock', '30 to 39'),
        sim_row('s5', 44, 'Felsnap', 'hunter', '58.4', 'Beast Mastery Hunter', '30 to 39'),
        sim_row('s6', 52, 'Kaelra', 'mage', '64.1', 'Frost Mage', '30 to 39'),
        sim_row('s7', 58, 'Brindle', 'warrior', '55.0', 'Arms Warrior', '30 to 39'),
    ]
    write('feed-34-guild.json', page(sims + hunter + PUBLIC))


if __name__ == '__main__':
    main()
