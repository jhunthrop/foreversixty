"""What is committed under builds/1.60.1.69893/ must satisfy the simulator contract.

Like tests/test_build_conformance.py and tests/test_beta_build.py, this reads
the real output rather than running the pipeline over fixtures. The Go half is
skipped unless an engine checkout and a Go toolchain are both available; CI
clones neither.
"""

import json
import os
import shutil
import subprocess
from functools import cache
from pathlib import Path

import pytest

from pipeline.simproto import pb

BUILD = "1.60.1.69893"
BUILD_DIR = Path("builds") / BUILD
GOCHECK = Path(__file__).parent / "gocheck"

EXPECTED_ITEMS = 4986
EXPECTED_ENCHANTS = 2216
EXPECTED_WEAPONS = 759
EXPECTED_IN_A_SET = 828
#: 334 InventoryType-{13,21} weapons split into two disjoint HandType buckets:
#: 216 InventoryType 13 (one-hand, either hand) and 118 InventoryType 21
#: (main-hand only). See pipeline/simdb/items.py's HAND_TYPE_BY_INVENTORY_TYPE
#: and tests/test_simdb_items.py's test_inventory_type_13_is_one_hand_not_main_hand,
#: which pins the mapping itself against the fixtures.
EXPECTED_ONE_HAND_WEAPONS = 216
EXPECTED_MAIN_HAND_WEAPONS = 118
#: 1,329 enchants carry a stat in the final database, counting every source:
#: the direct ITEM_MOD/resistance slots plus equip-spell auras. A narrower
#: count, 1,255, is the equip-spell path alone (40 more of that path's rows
#: carry only a weapon skill, which `pb.SimEnchant` cannot represent) -- see
#: pipeline/simdb/equip.py's module docstring, not this constant.
EXPECTED_ENCHANTS_WITH_STATS = 1329
#: The committed database's four contract-10.3 fields, measured against
#: the committed items.json that feeds two of them.
#: MaxCount == 1 over the 4,986 items simdb_item_rows keeps (measured by
#: running `pipeline loot` then `pipeline simdb` for 1.60.1.69893 and
#: reading builds/1.60.1.69893/simdb.bin -- not the same population as
#: items.json's, so this is not derivable from any earlier task's numbers).
EXPECTED_UNIQUE = 933
#: The whole-items.json figures from task 11's log line (loot: items.json:
#: 69 with suffix options, 819 faction-restricted). simdb keeps only 4,986
#: of the build's 19,171 items; both counts happen to survive that
#: narrowing unchanged -- measured, not assumed.
EXPECTED_FACTION_RESTRICTED = 819
EXPECTED_WITH_SUFFIX_OPTIONS = 69


@cache
def database() -> pb.SimDatabase:
    parsed = pb.SimDatabase()
    parsed.ParseFromString((BUILD_DIR / "simdb.bin").read_bytes())
    return parsed


def item(item_id: int) -> pb.SimItem:
    return next(row for row in database().items if row.id == item_id)


def test_the_committed_simdb_has_the_whole_item_universe():
    assert len(database().items) == EXPECTED_ITEMS
    assert len(database().enchants) == EXPECTED_ENCHANTS


def test_item_ids_are_sorted_and_unique():
    ids = [row.id for row in database().items]
    assert ids == sorted(ids)
    assert len(set(ids)) == len(ids)


def test_weapons_and_set_pieces_are_populated():
    weapons = [row for row in database().items if row.weapon_speed > 0]
    in_a_set = [row for row in database().items if row.set_id]
    assert len(weapons) == EXPECTED_WEAPONS
    assert len(in_a_set) == EXPECTED_IN_A_SET
    assert all(row.set_name for row in in_a_set)
    assert all(row.weapon_damage_max >= row.weapon_damage_min > 0 for row in weapons)


def test_one_hand_weapons_are_dual_wieldable_not_main_hand_locked():
    """InventoryType 13 (one-hand) must map to HandTypeOneHand, not
    HandTypeMainHand -- see pipeline/simdb/items.py's
    HAND_TYPE_BY_INVENTORY_TYPE. HandTypeMainHand is InventoryType 21's
    disjoint, main-hand-only bucket; collapsing 13 into it left the engine
    with zero HandTypeOneHand rows and no way to dual-wield a one-hand item
    typed InventoryType 13.

    This reads only the committed simdb.bin, which CI has -- the raw client
    tables (builds/<build>/raw/) are never committed (see .gitignore), so a
    row-level check against ItemSparse's InventoryType column cannot run
    here. tests/test_simdb_items.py's test_inventory_type_13_is_one_hand_not_main_hand
    pins the mapping itself against the fixtures instead; this test pins
    that the regenerated database actually reflects it: HandTypeOneHand is
    populated, and HandTypeMainHand holds only the true main-hand-only
    (InventoryType 21) rows rather than the pre-fix total of both.
    """
    one_hand = pb.HandType.Value("HandTypeOneHand")
    main_hand = pb.HandType.Value("HandTypeMainHand")
    assert sum(1 for row in database().items if row.hand_type == one_hand) == (
        EXPECTED_ONE_HAND_WEAPONS
    )
    assert sum(1 for row in database().items if row.hand_type == main_hand) == (
        EXPECTED_MAIN_HAND_WEAPONS
    )


def test_enchants_carry_the_stats_their_equip_spells_grant():
    assert sum(1 for row in database().enchants if row.stats) == EXPECTED_ENCHANTS_WITH_STATS


def test_random_suffixes_are_empty():
    """ItemRandomSuffix 404s on this build; an empty field is the honest answer."""
    assert list(database().random_suffixes) == []


def test_a_weapon_carries_the_damage_the_client_computes():
    reaper = item(12784)
    assert reaper.weapon_speed == pytest.approx(3.8)
    assert (reaper.weapon_damage_min, reaper.weapon_damage_max) == (153.0, 256.0)
    assert reaper.stats[pb.Stat.Value("StatAttackPower")] == 62.0


def test_a_set_piece_carries_its_armour_and_its_set():
    chest = item(22416)
    assert chest.armor_type == pb.ArmorType.Value("ArmorTypePlate")
    assert chest.stats[pb.Stat.Value("StatArmor")] == 1027.0
    assert (chest.set_id, chest.set_name) == (523, "Dreadnaught's Battlegear")


def test_a_known_on_equip_stat_survived_the_round_trip():
    """Rune of the Guard Captain's +42 attack power is in no ItemSparse column;
    its 7 hit rating is on the curve. Both must be in the same stat array.

    The engine lane has already collapsed MeleeHit/SpellHit into a single
    Hit stat, so this reads StatHit rather than the brief's StatMeleeHit.
    The client's 7 is a combat-rating amount (ItemSparse's
    StatModifier_bonusStat 31 is ITEM_MOD_HIT_RATING); the engine reads Hit
    as a flat percentage, so simdb divides it by this build's level-60 hit
    factor (10, from gametables/combatratings.txt) -- see
    pipeline/simdb/ratings.py.
    """
    rune = item(19120)
    assert rune.stats[pb.Stat.Value("StatAttackPower")] == 42.0
    assert rune.stats[pb.Stat.Value("StatRangedAttackPower")] == 42.0
    assert rune.stats[pb.Stat.Value("StatHit")] == pytest.approx(0.7)


def test_item_hit_and_crit_are_percentages_not_combat_rating_points():
    """Lionheart Helm (12640), Quick Strike Ring (18821) and Onyxia Tooth
    Pendant (18404) state hit/crit through StatModifier_bonusStat ids 31/32
    -- ITEM_MOD_HIT_RATING/ITEM_MOD_CRIT_RATING -- so ItemSparse's own
    amounts (20 hit / 28 crit, 14 crit, 10 hit / 14 crit respectively) are
    combat-rating points, not the flat percentage the engine's `StatHit` and
    `StatCrit` are. Divided by this build's level-60 factors (hit 10, crit
    14, from gametables/combatratings.txt -- see pipeline/simdb/ratings.py),
    they are 2% hit / 2% crit, 1% crit, and 1% hit / 1% crit. Before this
    conversion existed, every one of these was read as 10-28% hit or crit
    outright, hit-capping and near-100%-crit-ing any character wearing one.
    """
    hit = pb.Stat.Value("StatHit")
    crit = pb.Stat.Value("StatCrit")
    assert (item(12640).stats[hit], item(12640).stats[crit]) == (2.0, 2.0)
    assert item(18821).stats[crit] == 1.0
    assert (item(18404).stats[hit], item(18404).stats[crit]) == (1.0, 1.0)


def test_sim_items_carry_the_four_fields_contract_10_3_adds():
    items = database().items
    assert sum(1 for item in items if item.faction_restriction) == (
        EXPECTED_FACTION_RESTRICTED
    )
    assert sum(1 for item in items if item.random_suffix_options) == (
        EXPECTED_WITH_SUFFIX_OPTIONS
    )
    # Dreadnaught Breastplate (22416), also pinned in
    # test_a_set_piece_carries_its_armour_and_its_set -- `item.required_level
    # >= 0` would pass even if the field were never populated, since the
    # proto's int32 default is 0. This pins a real, known value instead.
    assert item(22416).required_level == 60
    assert sum(1 for item in items if item.unique) == EXPECTED_UNIQUE


def test_consumables_are_emitted_beside_the_protobuf():
    rows = json.loads((BUILD_DIR / "simconsumes.json").read_text())
    assert len(rows) == 1579
    assert all(row["spell_ids"] for row in rows)
    assert [row["id"] for row in rows] == sorted(row["id"] for row in rows)


def test_go_reads_back_what_python_wrote(tmp_path: Path):
    engine = os.environ.get("FOREVER_ENGINE_PATH")
    if not engine or not shutil.which("go"):
        pytest.skip("FOREVER_ENGINE_PATH or the Go toolchain is unavailable")
    shutil.copyfile(GOCHECK / "main.go", tmp_path / "main.go")
    template = (GOCHECK / "go.mod.tmpl").read_text()
    (tmp_path / "go.mod").write_text(template.replace("ENGINE_PATH", engine))
    # A fresh module has no go.sum; `go run` alone refuses to fetch checksums
    # under the default -mod=readonly. `go mod tidy` resolves the replaced
    # engine module's own (non-replaced) transitive dependencies from the
    # local module cache and writes it, the same way a developer's first
    # `go build` in a new module would.
    subprocess.run(["go", "mod", "tidy"], cwd=tmp_path, capture_output=True, text=True, check=True)
    result = subprocess.run(
        ["go", "run", ".", str(BUILD_DIR.resolve() / "simdb.bin")],
        cwd=tmp_path,
        capture_output=True,
        text=True,
        check=True,
    )
    report = json.loads(result.stdout)
    assert report["items"] == EXPECTED_ITEMS
    assert report["enchants"] == EXPECTED_ENCHANTS
    assert report["weapons"] == EXPECTED_WEAPONS
    assert report["in_a_set"] == EXPECTED_IN_A_SET
    assert report["enchants_with_stats"] == EXPECTED_ENCHANTS_WITH_STATS
    assert report["spot"]["reaper_attack_power"] == 62.0
    assert report["spot"]["reaper_stamina"] == 13.0
    assert report["spot"]["reaper_speed"] == pytest.approx(3.8)
    assert report["spot"]["reaper_damage_min"] == 153.0
    assert report["spot"]["reaper_damage_max"] == 256.0
    assert report["spot"]["dreadnaught_armor"] == 1027.0
    assert report["spot"]["dreadnaught_is_plate"] == 1.0
    assert report["spot"]["dreadnaught_in_a_set"] == 1.0
    assert report["spot"]["rune_attack_power"] == 42.0
    assert report["spot"]["rune_ranged_attack_power"] == 42.0
    assert report["spot"]["rune_hit"] == 7.0
    assert report["spot"]["player_database_items"] == EXPECTED_ITEMS
