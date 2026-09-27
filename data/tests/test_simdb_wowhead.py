"""The wowhead supplement's own SimItem rows (Part 1 of the wowhead item
supplement design, docs/superpowers/specs/2026-09-27-wowhead-item-supplement-
design.md): ids the client's ItemSparse lacks entirely, added to simdb.bin on
top of the client universe.

Same fixture as tests/test_wowhead_items.py: a cloak the client lacks, a
one-hand sword with damage, a set id and a class mask, a quality-1 ring the
planner's own gates drop, a belt the client already carries, a QA-named epic
and a level-70 chest.
"""

import json
import shutil
from collections import Counter
from pathlib import Path

import pytest

from pipeline import wowhead_items as wh
from pipeline.csvio import read_csv
from pipeline.manifest import write_manifest
from pipeline.normalize import write_json
from pipeline.normalize.items import normalize_items
from pipeline.simdb import write_sim_database
from pipeline.simdb.items import build_wowhead_sim_items
from pipeline.simdb.statmap import stat_keys
from pipeline.simproto import pb

HERE = Path(__file__).parent
FIXTURES = HERE / "fixtures"
SIM = FIXTURES / "sim"
FIXTURE = FIXTURES / "wowhead-gear-planner.js"

#: The same raw tables tests/test_simdb.py's build_dir fixture assembles --
#: this module needs the identical shape to exercise write_sim_database, so
#: it is not a new fixture, only a second build directory to drop the
#: wowhead payload into.
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


def make_build_dir(tmp_path: Path, *, with_wowhead_payload: bool) -> Path:
    build = tmp_path / "9.9.9.9"
    (build / "raw").mkdir(parents=True)
    for source, name in RAW_FIXTURES:
        shutil.copyfile(source, build / "raw" / name)
    if with_wowhead_payload:
        shutil.copyfile(FIXTURE, build / "raw" / wh.RAW_FILE)
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

#: The committed build's own level-60 combatratings.txt row, the same
#: RATING_FACTORS tests/test_simdb_items.py uses, so a wowhead crit converts
#: through the identical divisor a client row's crit would.
RATING_FACTORS = {
    "hit": 10.0,
    "crit": 14.0,
    "dodge": 12.0,
    "parry": 15.0,
    "block": 5.0,
    "defense": 1.0,
}

#: The belt (client-owned) and the heirloom ring (quality 1) stand in for
#: "the client already has this id" -- the same client_ids
#: tests/test_wowhead_items.py's test_supplement_is_the_planner_gear_the_client_lacks
#: passes, so both suites agree on what "the client" means for this fixture.
CLIENT_IDS = {11726, 264908}


def supplement_items() -> list[wh.WowheadItem]:
    return wh.supplement(wh.load_items(FIXTURE), CLIENT_IDS)


def built(untracked: Counter[str] | None = None, set_names=None) -> dict[int, pb.SimItem]:
    items = build_wowhead_sim_items(
        supplement_items(), RATING_FACTORS, set_names=set_names, untracked=untracked
    )
    return {item.id: item for item in items}


def test_only_the_planner_gear_the_client_lacks_is_built():
    """The ring is quality 1, the QA robe is junk-named, the chest needs level
    70, and the belt is the client's -- only the cloak and the sword remain."""
    assert set(built()) == {271218, 279865}


def test_the_sword_is_a_one_hand_sword_with_its_own_damage_and_speed():
    sword = built()[271218]
    assert sword.type == pb.ItemType.Value("ItemTypeWeapon")
    assert sword.weapon_type == pb.WeaponType.Value("WeaponTypeSword")
    assert sword.hand_type == pb.HandType.Value("HandTypeOneHand")
    assert (sword.weapon_damage_min, sword.weapon_damage_max, sword.weapon_speed) == (
        26.0,
        50.0,
        2.6,
    )


def test_the_sword_s_crit_is_converted_like_a_client_row():
    """3 crit rating at this fixture's level-60 crit factor (14) is 3/14%, the
    same conversion ratings.convert_rating_stats applies to a client row's
    combat-rating-denominated stat."""
    sword = built()[271218]
    assert sword.stats[pb.Stat.Value("StatCrit")] == pytest.approx(3.0 / 14.0)
    assert sword.stats[pb.Stat.Value("StatAgility")] == 2.0
    assert sword.stats[pb.Stat.Value("StatStrength")] == 4.0


def test_the_sword_s_class_allowlist_comes_from_its_mask():
    """classMask 1029 is bits 0, 2 and 10: warrior, hunter and druid."""
    sword = built()[271218]
    names = {pb.Class.Name(value) for value in sword.class_allowlist}
    assert names == {"ClassWarrior", "ClassHunter", "ClassDruid"}


def test_the_sword_carries_its_set_id():
    sword = built(set_names={9001: "Vileblood's Battlegear"})[271218]
    assert (sword.set_id, sword.set_name) == (9001, "Vileblood's Battlegear")


def test_the_sword_is_unique_and_has_its_required_level():
    sword = built()[271218]
    assert sword.unique is True
    assert sword.required_level == 17


def test_the_sword_has_no_on_equip_fields():
    """No on-equip effects: wowhead states them as prose, not data."""
    sword = built()[271218]
    assert list(sword.random_suffix_options) == []
    assert sword.faction_restriction == pb.SimItem.FactionRestriction.Value(
        "FACTION_RESTRICTION_UNSPECIFIED"
    )
    assert list(sword.weapon_skills) == []
    assert sword.bonus_physical_damage == 0.0


def test_the_cloak_has_its_armour_and_no_allowlist():
    cloak = built()[279865]
    assert cloak.type == pb.ItemType.Value("ItemTypeBack")
    assert cloak.armor_type == pb.ArmorType.Value("ArmorTypeCloth")
    assert cloak.stats[pb.Stat.Value("StatArmor")] == 20.0
    assert cloak.stats[pb.Stat.Value("StatStamina")] == 5.0
    assert cloak.stats[pb.Stat.Value("StatStrength")] == 3.0
    assert list(cloak.class_allowlist) == []


def test_the_cloak_has_no_weapon_fields():
    cloak = built()[279865]
    assert (cloak.weapon_damage_min, cloak.weapon_damage_max, cloak.weapon_speed) == (
        0.0,
        0.0,
        0.0,
    )
    assert cloak.weapon_type == pb.WeaponType.Value("WeaponTypeUnknown")


def test_the_cloak_is_not_unique_and_carries_no_set():
    cloak = built()[279865]
    assert cloak.unique is False
    assert cloak.set_id == 0
    assert cloak.set_name == ""


def test_an_untracked_wowhead_stat_is_counted_not_raised():
    """hastertng (haste rating) has no engine-bound planner stat
    (wowhead_items.STAT_KEYS maps it to None); it is dropped from the
    sword's stats and counted rather than raising."""
    untracked: Counter[str] = Counter()
    items = built(untracked=untracked)
    assert untracked == {"hastertng": 1}
    sword = items[271218]
    populated = stat_keys(sword.stats)
    assert populated.pop("crit") == pytest.approx(3.0 / 14.0)
    assert populated == {"agility": 2.0, "strength": 4.0}


def test_no_set_names_defaults_to_an_empty_name():
    sword = built()[271218]  # no set_names argument passed
    assert sword.set_id == 9001
    assert sword.set_name == ""


def test_an_item_with_no_class_mask_has_no_class_restriction():
    """The fixture models this on the cloak; a mask value of exactly -1 (the
    client's own 'every class' sentinel) must not be confused with wowhead's
    None -- both mean unrestricted but through different vocabularies, and
    only wowhead's classMask column speaks the mask form at all."""
    cloak = next(item for item in supplement_items() if item.id == 279865)
    assert cloak.class_mask is None


def parsed(path: Path) -> pb.SimDatabase:
    database = pb.SimDatabase()
    database.ParseFromString(path.read_bytes())
    return database


def test_a_build_with_no_wowhead_payload_is_unchanged(tmp_path: Path):
    """Part 1's own requirement: a build that never fetched wowhead's payload
    (every build before fetch-wowhead became a workflow step) must come out
    exactly as write_sim_database produced it before this feature existed --
    the same 8 items tests/test_simdb.py's identical fixture set pins."""
    build_dir = make_build_dir(tmp_path, with_wowhead_payload=False)
    database = parsed(write_sim_database("9.9.9.9", root=build_dir.parent))
    assert len(database.items) == 8
    assert {271218, 279865, 11726} & {item.id for item in database.items} == set()


def test_a_build_with_the_payload_adds_the_supplement_on_top(tmp_path: Path):
    """None of this fixture's ItemSparse ids (12798, 14152, 16866, 30001-30006,
    40001) collide with the wowhead fixture's, so the sword, the cloak and the
    belt are all "the client lacks this id" here -- the ring, the QA robe and
    the chest still fail the planner's own gates."""
    build_dir = make_build_dir(tmp_path, with_wowhead_payload=True)
    path = write_sim_database("9.9.9.9", root=build_dir.parent)
    database = parsed(path)
    ids = [item.id for item in database.items]
    assert ids == sorted(ids)
    assert len(database.items) == 11
    added = {271218, 279865, 11726} & set(ids)
    assert added == {271218, 279865, 11726}
    assert {264908, 777777, 888888} & set(ids) == set()

    sword = next(item for item in database.items if item.id == 271218)
    assert sword.weapon_type == pb.WeaponType.Value("WeaponTypeSword")
    assert sword.hand_type == pb.HandType.Value("HandTypeOneHand")
    assert (sword.weapon_damage_min, sword.weapon_damage_max, sword.weapon_speed) == (
        26.0,
        50.0,
        2.6,
    )


def test_simitems_json_lists_the_supplement_ids_too(tmp_path: Path):
    build_dir = make_build_dir(tmp_path, with_wowhead_payload=True)
    write_sim_database("9.9.9.9", root=build_dir.parent)
    simitems = json.loads((build_dir / "simitems.json").read_text(encoding="utf-8"))
    assert {271218, 279865, 11726} <= set(simitems["items"])
    assert simitems["items"] == sorted(simitems["items"])
