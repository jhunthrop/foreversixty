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
#: Controller ruling: 1,329, not the 1,334 the brief originally stated --
#: measured against the committed build, matching equip.py's own docstring.
EXPECTED_ENCHANTS_WITH_STATS = 1329


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
    its 7 hit is on the curve. Both must be in the same stat array.

    The engine lane has already collapsed MeleeHit/SpellHit into a single
    Hit stat, so this reads StatHit rather than the brief's StatMeleeHit.
    """
    rune = item(19120)
    assert rune.stats[pb.Stat.Value("StatAttackPower")] == 42.0
    assert rune.stats[pb.Stat.Value("StatRangedAttackPower")] == 42.0
    assert rune.stats[pb.Stat.Value("StatHit")] == 7.0


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
