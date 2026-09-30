"""The classic-db supplement's own `SimItem` rows (simdb-supplement lane,
2026-09-30): cmangos/classic-db's 1.12 `item_template`, for the ids neither
the client's `ItemSparse` nor wowhead's Forever gear-planner scrape carry at
all (`pipeline.classicdb_items`'s own doc: Hand of Justice, Blackhand's
Breadth, Devilsaur Eye and eight more).

Three contracts this lane's own brief names:

1. A combat-rating stat off classic-db's `stat_type<n>`/`stat_value<n>`
   columns converts exactly like a client row's does (`test_a_classicdb_
   crit_rating_converts_like_a_client_row`).
2. A wowhead-only weapon still lands with its own damage and speed -- the
   already-shipped path this lane extends, not replaces
   (`test_a_wowhead_only_weapon_still_lands_with_its_damage`).
3. An id the client already states wins outright; classic-db's own row for
   the same id never reaches `simdb.bin` (`test_a_client_row_wins_over_a_
   classicdb_row_for_the_same_id`).
"""

import json
import shutil
from collections import Counter
from pathlib import Path

import pytest

from pipeline import wowhead_items as wh
from pipeline.classicdb_items import ClassicDbItem
from pipeline.csvio import read_csv
from pipeline.manifest import write_manifest
from pipeline.normalize import write_json
from pipeline.normalize.items import normalize_items
from pipeline.simdb import write_sim_database
from pipeline.simdb.items import build_classicdb_sim_items, build_wowhead_sim_items
from pipeline.simproto import pb

HERE = Path(__file__).parent
FIXTURES = HERE / "fixtures"
SIM = FIXTURES / "sim"

#: The committed build's own level-60 combatratings.txt row -- the same
#: RATING_FACTORS tests/test_simdb_items.py and tests/test_simdb_wowhead.py
#: use, so a classic-db crit converts through the identical divisor a
#: client row's crit would.
RATING_FACTORS = {
    "hit": 10.0,
    "crit": 14.0,
    "dodge": 12.0,
    "parry": 15.0,
    "block": 5.0,
    "defense": 1.0,
}


def _trinket(**overrides: object) -> ClassicDbItem:
    values: dict[object, object] = {
        "id": 900001,
        "name": "Trial Trinket",
        "quality": 3,
        "item_level": 40,
        "required_level": 40,
        "class_id": 4,  # ARMOR
        "subclass_id": 0,  # misc -- rings, necks, trinkets
        "inventory_type": 12,  # trinket
        "allowable_class": -1,
        "allowable_race": -1,
        "armor": 0,
        "raw_stats": {},
        "resistances": {},
        "damage_min": 0,
        "damage_max": 0,
        "delay": 0,
        "set_id": None,
        "unique": False,
        "spells": [],
    }
    values.update(overrides)
    return ClassicDbItem(**values)


def test_a_classicdb_crit_rating_converts_like_a_client_row():
    """14 crit rating at this fixture's level-60 crit factor (14) is 14/14
    = 1.0 -- the same `ratings.convert_rating_stats` division a client
    row's ItemSparse-column crit rating gets."""
    trinket = _trinket(raw_stats={32: 14})  # 32 == STAT_BY_MODIFIER_ID's crit id
    built = build_classicdb_sim_items([trinket], {}, RATING_FACTORS)
    assert len(built) == 1
    assert built[0].stats[pb.Stat.Value("StatCrit")] == pytest.approx(1.0)


def test_a_classicdb_item_with_no_simdb_inventory_type_is_dropped():
    """A defensive floor, not a real gap today: `simdb.items.
    ITEM_TYPE_BY_INVENTORY_TYPE` and `normalize/gear.py`'s own
    `SLOT_BY_INVENTORY_TYPE` are kept in sync (28, relics, was the one real
    divergence -- fixed alongside this lane), but `build_classicdb_sim_items`
    still refuses to guess at an inventory type this module has never mapped,
    counting it in `untracked` instead of raising or crashing."""
    unmapped = _trinket(id=900002, inventory_type=999)
    counter: Counter[str] = Counter()
    built = build_classicdb_sim_items([unmapped], {}, RATING_FACTORS, untracked=counter)
    assert built == []
    assert counter == {"inventory_type_999": 1}


def test_a_wowhead_only_weapon_still_lands_with_its_damage():
    """The already-shipped wowhead path this lane extends, not replaces:
    a synthetic one-hand sword wowhead states with no ItemSparse row still
    lands in simdb with its own damage and speed."""
    sword = wh.WowheadItem(
        id=800001,
        name="Trial Sword",
        quality=3,
        item_level=40,
        required_level=17,
        class_id=2,  # WEAPON
        subclass_id=7,  # one-hand sword
        inventory_type=13,
        icon="inv_sword_04",
        class_mask=None,
        stats={"agi": 2.0},
        set_id=None,
        unique=False,
        armor=0,
        damage_min=26,
        damage_max=50,
        speed=2.6,
        dps=14.6,
    )
    built = build_wowhead_sim_items([sword], RATING_FACTORS)
    assert len(built) == 1
    item = built[0]
    assert item.weapon_type == pb.WeaponType.Value("WeaponTypeSword")
    assert (item.weapon_damage_min, item.weapon_damage_max, item.weapon_speed) == (26.0, 50.0, 2.6)


#: Same raw tables tests/test_simdb.py's build_dir fixture assembles.
RAW_FIXTURES = (
    (FIXTURES / "ItemSparse_1_60.csv", "ItemSparse.csv"),
    (FIXTURES / "Item.csv", "Item.csv"),
    (FIXTURES / "ItemArmorTotal.csv", "ItemArmorTotal.csv"),
    (FIXTURES / "ItemArmorQuality.csv", "ItemArmorQuality.csv"),
    (FIXTURES / "ItemArmorShield.csv", "ItemArmorShield.csv"),
    (FIXTURES / "ArmorLocation.csv", "ArmorLocation.csv"),
    (FIXTURES / "RandPropPoints.csv", "RandPropPoints.csv"),
    (SIM / "ItemEffect.csv", "ItemEffect.csv"),
    (SIM / "ItemXItemEffect.csv", "ItemXItemEffect.csv"),
    (SIM / "SpellEffect.csv", "SpellEffect.csv"),
    (SIM / "SpellItemEnchantment.csv", "SpellItemEnchantment.csv"),
    (SIM / "ItemDamageOneHand.csv", "ItemDamageOneHand.csv"),
    (SIM / "ItemDamageTwoHand.csv", "ItemDamageTwoHand.csv"),
    (SIM / "ItemDamageRanged.csv", "ItemDamageRanged.csv"),
    (SIM / "ItemDamageWand.csv", "ItemDamageWand.csv"),
    (SIM / "ItemDamageThrown.csv", "ItemDamageThrown.csv"),
)

#: 12798 "Annihilator" -- a real client weapon both ItemSparse_1_60.csv and
#: Item.csv fixtures carry (tests/test_simdb.py's own 8-item baseline keeps
#: it), reused here as the id a classic-db extract row ALSO names.
CLIENT_ITEM_ID = 12798


def make_build_dir(tmp_path: Path, *, extract_items: list[dict]) -> Path:
    build = tmp_path / "9.9.9.9"
    (build / "raw").mkdir(parents=True)
    for source, name in RAW_FIXTURES:
        shutil.copyfile(source, build / "raw" / name)
    classicdb_dir = build / "raw" / "classicdb"
    classicdb_dir.mkdir(parents=True)
    (classicdb_dir / "item_template.json").write_text(
        json.dumps({"items": extract_items, "spells": []}), encoding="utf-8"
    )
    shutil.copyfile(SIM / "sets.json", build / "sets.json")
    write_json(
        normalize_items(
            read_csv(build / "raw" / "ItemSparse.csv"), read_csv(build / "raw" / "Item.csv")
        ),
        build / "items.json",
    )
    (build / "gametables").mkdir()
    shutil.copyfile(SIM / "combatratings.txt", build / "gametables" / "combatratings.txt")
    write_manifest(
        build, build="9.9.9.9", product="wow_classic_beta", fetched_at="2026-01-01T00:00:00Z"
    )
    return build


def _extract_row(item_id: int, **overrides: object) -> dict:
    row = _trinket(id=item_id).model_dump()
    row.update(overrides)
    return row


def parsed(path: Path) -> pb.SimDatabase:
    database = pb.SimDatabase()
    database.ParseFromString(path.read_bytes())
    return database


def test_a_client_row_wins_over_a_classicdb_row_for_the_same_id(tmp_path: Path):
    """A classic-db extract row for 12798 (the client's own real
    "Annihilator") must never reach simdb.bin: the client's own weapon
    stays exactly as tests/test_simdb.py's baseline pins it, and the
    classic-db row's fake +999 strength is dropped outright, not merged or
    duplicated."""
    build_dir = make_build_dir(
        tmp_path,
        extract_items=[
            _extract_row(CLIENT_ITEM_ID, name="Fake Duplicate", raw_stats={4: 999}),
            _extract_row(9990001, name="Real Classicdb Item"),
        ],
    )
    path = write_sim_database("9.9.9.9", root=build_dir.parent)
    database = parsed(path)
    ids = [item.id for item in database.items]
    assert ids.count(CLIENT_ITEM_ID) == 1
    client_item = next(item for item in database.items if item.id == CLIENT_ITEM_ID)
    # ItemSparse_1_60.csv's own Display_lang for 12798 is "Annihilator" --
    # the classic-db row's "Fake Duplicate" name never wins.
    assert client_item.name == "Annihilator"
    strength = pb.Stat.Value("StatStrength")
    assert strength >= len(client_item.stats) or client_item.stats[strength] != 999.0
    # The classic-db-only id (never in ItemSparse/Item at all) still lands.
    assert 9990001 in ids

    simitems = json.loads((build_dir / "simitems.json").read_text(encoding="utf-8"))
    assert simitems["sim_source"][str(CLIENT_ITEM_ID)] == "client"
    assert simitems["sim_source"]["9990001"] == "classic-db"
