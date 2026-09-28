import json
from pathlib import Path

from pipeline.forkdb import ForkDatabase, load_fork_database
from pipeline.loot.weapons import apply_fork_weapon_damage, fork_weapon_damage
from pipeline.models import ClassItems, GearItem

HERE = Path(__file__).parent
ENGINE = HERE / "fixtures/loot"


def fork():
    return load_fork_database(ENGINE)


def synthetic_fork(*items: dict) -> ForkDatabase:
    """A `ForkDatabase` with only `items` populated, for a case the shared
    `fixtures/loot` fork database (reused by test_loot_gear.py, test_loot.py,
    test_loot_sources.py and test_loot_buffs.py) has no reason to carry."""
    return ForkDatabase(
        items=items,
        enchants=(),
        random_suffixes=(),
        zones={},
        npcs={},
        factions={},
        item_icons={},
        spell_icons={},
        spell_icon_rows=(),
        item_icon_rows=(),
    )


def gear_item(item_id: int, **overrides) -> GearItem:
    base = dict(
        id=item_id,
        name=f"Item {item_id}",
        icon="inv_x",
        slot="main_hand",
        quality=3,
        required_level=40,
        item_level=50,
        armor=0,
        stats={},
        set_id=None,
        unique=False,
    )
    base.update(overrides)
    return GearItem(**base)


# --- fork_weapon_damage --------------------------------------------------


def test_fork_weapon_damage_reads_the_fixture_sword():
    """Suffixed Sword (110) on the shared loot fixture carries the fork's
    own damage, added alongside its suffix options."""
    assert fork_weapon_damage(fork()) == {110: (38, 71, 2.6)}


def test_a_weapon_type_with_no_damage_is_left_out():
    """A fork row the fork itself has not itemised yet (weaponType set, no
    weaponDamageMin/Max/Speed) is exactly the gap the curve derivation
    exists to fill -- it must not be mapped to (0, 0, 0.0)."""
    undamaged = synthetic_fork({"id": 500, "name": "Unitemised Axe", "weaponType": 1})
    assert fork_weapon_damage(undamaged) == {}


def test_a_partial_row_is_left_out_too():
    """Every one of min/max/speed must be present; a row missing one is not
    trustworthy data to overwrite `normalize`'s output with."""
    partial = synthetic_fork(
        {"id": 501, "weaponDamageMin": 10, "weaponDamageMax": 20},  # no weaponSpeed
        {"id": 502, "weaponSpeed": 2.0},  # no min/max
    )
    assert fork_weapon_damage(partial) == {}


def test_a_non_weapon_row_is_naturally_left_out():
    assert fork_weapon_damage(synthetic_fork({"id": 600, "name": "A Ring"})) == {}


# --- apply_fork_weapon_damage --------------------------------------------


def prepared(tmp_path: Path) -> Path:
    build_dir = tmp_path / "1.60.1.69893"
    items_dir = build_dir / "items"
    items_dir.mkdir(parents=True)
    warrior = ClassItems(
        build="1.60.1.69893",
        class_slug="warrior",
        items=[
            # The fork itemises 110 -- normalize left it at the honest zero
            # a curve-only build with no ItemDamage* fetch would report.
            gear_item(110, damage_min=0, damage_max=0, speed=2.6, dps=0.0),
            # The fork does not know 999 -- normalize's own curve-derived
            # value must survive untouched.
            gear_item(999, damage_min=44, damage_max=82, speed=2.9, dps=21.72),
        ],
    )
    (items_dir / "warrior.json").write_text(
        json.dumps(warrior.model_dump(), indent=2) + "\n", encoding="utf-8"
    )
    return build_dir


def test_the_forks_numbers_win_where_it_itemises_a_weapon(tmp_path):
    build_dir = prepared(tmp_path)
    won = apply_fork_weapon_damage(build_dir, fork_weapon_damage(fork()))
    assert won == 1
    rows = json.loads((build_dir / "items" / "warrior.json").read_text(encoding="utf-8"))
    by_id = {row["id"]: row for row in rows["items"]}
    assert (by_id[110]["damage_min"], by_id[110]["damage_max"], by_id[110]["speed"]) == (
        38,
        71,
        2.6,
    )
    assert by_id[110]["dps"] == round((38 + 71) / 2 / 2.6, 2)


def test_normalizes_own_derived_value_survives_when_the_fork_lacks_the_id(tmp_path):
    build_dir = prepared(tmp_path)
    apply_fork_weapon_damage(build_dir, fork_weapon_damage(fork()))
    rows = json.loads((build_dir / "items" / "warrior.json").read_text(encoding="utf-8"))
    by_id = {row["id"]: row for row in rows["items"]}
    assert (by_id[999]["damage_min"], by_id[999]["damage_max"], by_id[999]["speed"]) == (
        44,
        82,
        2.9,
    )
    assert by_id[999]["dps"] == 21.72


def test_is_idempotent(tmp_path):
    build_dir = prepared(tmp_path)
    damage = fork_weapon_damage(fork())
    first = apply_fork_weapon_damage(build_dir, damage)
    second_bytes = (build_dir / "items" / "warrior.json").read_text(encoding="utf-8")
    second = apply_fork_weapon_damage(build_dir, damage)
    assert (first, second) == (1, 1)
    assert (build_dir / "items" / "warrior.json").read_text(encoding="utf-8") == second_bytes


def test_a_build_with_no_items_directory_wins_nothing(tmp_path):
    build_dir = tmp_path / "1.60.1.69893"
    build_dir.mkdir()
    assert apply_fork_weapon_damage(build_dir, fork_weapon_damage(fork())) == 0
