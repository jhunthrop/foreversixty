"""Real data for the design-direction mocks (design/specs/2026-11-09-directions.md).

Everything here is read from the repo: the home parity fixtures (web/tests/e2e/fixtures/home), the BiS facts read from
the built page (design/mocks/data/bis-facts.json), the active build's item files and icons. Nothing is invented. The
one table typed by hand is ZONE_RANGES (the Classic world-map level ranges), named as a data gap in the spec.
"""
import base64
import json
from datetime import datetime, timezone
from functools import lru_cache
from pathlib import Path

ROOT = Path(__file__).resolve().parents[3]
BUILD = json.loads((ROOT / 'web/src/data/active-build.json').read_text())['build']
HOME_FIX = ROOT / 'web/tests/e2e/fixtures/home'
FACTS = json.loads((ROOT / 'design/mocks/data/bis-facts.json').read_text())
ICON_DIRS = [ROOT / 'data/builds' / BUILD / 'icons', *sorted((ROOT / 'web/public/data').glob('*/icons'))]
CREST_DIR = ROOT / 'design/community/discord/assets'
NOW = datetime(2026, 11, 6, 18, 30, tzinfo=timezone.utc)  # the mock clock: the fixtures' own day, so feed times read as they would live

# Game colours, exactly as the client draws them.
QUALITY = {'poor': '#9d9d9d', 'common': '#ffffff', 'uncommon': '#1eff00', 'rare': '#0070dd', 'epic': '#a335ee', 'legendary': '#ff8000'}
QUALITY_BY_NUM = {0: 'poor', 1: 'common', 2: 'uncommon', 3: 'rare', 4: 'epic', 5: 'legendary'}
CLASS_COLOR = {'warrior': '#c69b6d', 'paladin': '#f48cba', 'hunter': '#aad372', 'rogue': '#fff468', 'priest': '#ffffff',
               'shaman': '#0070dd', 'mage': '#3fc7eb', 'warlock': '#8788ee', 'druid': '#ff7c0a'}
RAIDS = [('Barrow Deeps', 'spell_shadow_raisedead'), ('Hyjal Summit', 'spell_nature_timestop'), ("Onyxia's Lair", 'inv_misc_head_dragon_black')]
RAIDS_OPEN = datetime(2026, 12, 9, 0, 0, tzinfo=timezone.utc)  # web/src/data/phases.json raids-1

# Classic world-map level ranges, Alliance road. DATA GAP: zones.json carries level_min/level_max as null; the pipeline
# should fill them from the client's UiMap levels before a build uses this table.
ZONE_RANGES = [('Elwynn Forest', 1, 10), ('Westfall', 10, 20), ('Loch Modan', 10, 20), ('Redridge Mountains', 15, 25),
               ('Duskwood', 18, 30), ('Wetlands', 20, 30), ('Ashenvale', 18, 30), ('Hillsbrad Foothills', 20, 30),
               ('Thousand Needles', 25, 35), ('Desolace', 30, 40), ('Arathi Highlands', 30, 40), ('Stranglethorn Vale', 30, 45),
               ('Dustwallow Marsh', 35, 45), ('Badlands', 35, 45), ('Feralas', 40, 50), ('Tanaris', 40, 50),
               ('The Hinterlands', 40, 50), ("Un'Goro Crater", 48, 55), ('Felwood', 48, 55), ('Burning Steppes', 50, 58),
               ('Western Plaguelands', 51, 58), ('Eastern Plaguelands', 53, 60), ('Winterspring', 55, 60)]


# One common Alliance levelling path through those zones (wow-player, round 3: a route by mid-level alone hops between
# continents). Typed by hand like ZONE_RANGES; the guide writers confirm it before a build uses it.
ALLIANCE_ROAD = ['Elwynn Forest', 'Westfall', 'Redridge Mountains', 'Duskwood', 'Wetlands', 'Ashenvale', 'Hillsbrad Foothills',
                 'Thousand Needles', 'Arathi Highlands', 'Desolace', 'Stranglethorn Vale', 'Dustwallow Marsh', 'Badlands', 'Tanaris',
                 'Feralas', 'The Hinterlands', "Un'Goro Crater", 'Felwood', 'Burning Steppes', 'Western Plaguelands',
                 'Eastern Plaguelands', 'Winterspring']


def _json(path: Path) -> dict:
    return json.loads(path.read_text())


def data_uri(path: Path, mime: str) -> str:
    return f'data:{mime};base64,' + base64.b64encode(path.read_bytes()).decode()


@lru_cache(maxsize=None)
def icon_uri(name: str) -> str:
    """A real game icon (the client's own art, HD webp from the build) as a data URI; '' when the build lacks it."""
    path = next((d / f'{name}.webp' for d in ICON_DIRS if (d / f'{name}.webp').exists()), None)
    return data_uri(path, 'image/webp') if path else ''


@lru_cache(maxsize=None)
def crest_uri(slug: str) -> str:
    return data_uri(CREST_DIR / f'crest-{slug}.png', 'image/png')


def ago(at: str) -> str:
    secs = (NOW - datetime.fromisoformat(at.replace('Z', '+00:00'))).total_seconds()
    if secs < 3600:
        return f'{int(secs // 60)} min ago'
    if secs < 86400:
        return f'{int(secs // 3600)} h ago'
    return f'{int(secs // 86400)} d ago'


def countdown() -> tuple[int, int, int]:
    secs = int((RAIDS_OPEN - NOW).total_seconds())
    return secs // 86400, secs % 86400 // 3600, secs % 3600 // 60


def _items(class_slug: str) -> dict[int, dict]:
    full = _json(ROOT / 'data/builds' / BUILD / 'items' / f'{class_slug}.json')['items']
    return {i['id']: i for i in full}


# Items the BiS fix gates to raids-1 (9 Dec), which the home fixtures predate: Greenhammer (Hyjal-tier crafted mace).
# A gated pick is not an upgrade yet, so the mock drops it, as post-fix data would. Remove once the fixtures regenerate.
GATED_RAIDS_1 = {279261}


def upgrades(sample: str, class_slug: str, spec_file: str) -> list[dict]:
    """The character's BiS gains, lib/home/upgrades.ts rule b: a worn listed alternative's dps_delta, sign flipped.
    Picks gated to a later phase (GATED_RAIDS_1) are left out."""
    me = _json(HOME_FIX / f'me-{sample}.json')['characters'][0]
    gear = me['build']['gear']
    bis = _json(HOME_FIX / f'bis-{spec_file}.json')
    band = bis['bands'][0]
    items = _items(class_slug)
    rows = []
    for slot in band['slots']:
        if slot['item_id'] in GATED_RAIDS_1:
            continue
        worn = gear.get(slot['slot'])
        alt = next((a for a in slot.get('alternatives') or [] if a['item_id'] == worn), None)
        if not alt or alt.get('dps_delta', 0) >= 0:
            continue
        it = items.get(slot['item_id'], {})
        rows.append({'id': slot['item_id'], 'slot': slot['slot'], 'name': slot['item_name'], 'source': slot.get('source', ''), 'kind': slot.get('source_kind', ''),
                     'gain': -alt['dps_delta'], 'icon': it.get('icon', ''), 'quality': QUALITY_BY_NUM.get(it.get('quality', 3), 'rare'),
                     'worn': alt['item_name']})
    return sorted(rows, key=lambda r: -r['gain'])


def home(sample: str) -> dict:
    """One signed-in home: the character, its import sim, the feed, the week, the guild."""
    me = _json(HOME_FIX / f'me-{sample}.json')
    char = me['characters'][0]
    slug = char['class'].lower()
    spec_file = f"{slug}-{char['spec'].lower().replace(' ', '-')}"
    feed = sorted(_json(HOME_FIX / f'feed-{sample}.json')['rows'], key=lambda r: r['at'], reverse=True)
    guild_path = HOME_FIX / f'guild-home-{sample}.json'
    return {'char': char, 'slug': slug, 'home': _json(HOME_FIX / f'home-{sample}.json'), 'feed': feed,
            'guild': _json(guild_path) if guild_path.exists() else None, 'upgrades': upgrades(sample, slug, spec_file)}


def bis(key: str) -> dict:
    return FACTS[key]


SLOT_NAME = {'head': 'Head', 'neck': 'Neck', 'shoulder': 'Shoulder', 'back': 'Back', 'chest': 'Chest', 'wrist': 'Wrist',
             'hands': 'Hands', 'waist': 'Waist', 'legs': 'Legs', 'feet': 'Feet', 'finger1': 'Finger', 'finger2': 'Finger',
             'trinket1': 'Trinket', 'trinket2': 'Trinket', 'main_hand': 'Main Hand', 'off_hand': 'Off Hand', 'ranged': 'Ranged'}
STAT_NAME = {'agility': 'Agility', 'strength': 'Strength', 'stamina': 'Stamina', 'intellect': 'Intellect', 'spirit': 'Spirit'}
BIND = {'pickup': 'Binds when picked up', 'equip': 'Binds when equipped'}


def tooltip_lines(class_slug: str, item_id: int) -> list[tuple[str, str, str]]:
    """An item's tooltip as the client prints it: (left text, right text, colour role) per line, from the item file."""
    it = _items(class_slug)[item_id]
    lines = [(it['name'], '', QUALITY_BY_NUM[it['quality']])]
    if it.get('bonding') in BIND:
        lines.append((BIND[it['bonding']], '', 'white'))
    lines.append((SLOT_NAME.get(it['slot'], it['slot']), (it.get('armor_type') or '').title(), 'white'))
    if it.get('armor'):
        lines.append((f"{it['armor']} Armor", '', 'white'))
    lines += [(f'+{v} {STAT_NAME.get(k, k.title())}', '', 'white') for k, v in it['stats'].items() if k in STAT_NAME]
    lines.append((f"Requires Level {it['required_level']}", '', 'white'))
    rep = it.get('reputation')
    if rep:
        lines.append((f"Requires {rep['faction_name']} - {rep['standing'].title()}", '', 'white'))
    if it.get('effect_text'):
        lines.append((it['effect_text'], '', 'green'))
    return lines
