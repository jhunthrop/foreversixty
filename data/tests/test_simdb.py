import json
import shutil
from pathlib import Path

import pytest

from pipeline.csvio import read_csv
from pipeline.manifest import verify, write_manifest
from pipeline.simdb import build_consumables, write_sim_database
from pipeline.simproto import pb

HERE = Path(__file__).parent
FIXTURES = HERE / "fixtures"
SIM = FIXTURES / "sim"

#: (source path, name in builds/<build>/raw/). The item and curve tables come
#: from the shared fixtures; the simulator's own tables from fixtures/sim/.
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


@pytest.fixture
def build_dir(tmp_path: Path) -> Path:
    build = tmp_path / "9.9.9.9"
    (build / "raw").mkdir(parents=True)
    for source, name in RAW_FIXTURES:
        shutil.copyfile(source, build / "raw" / name)
    shutil.copyfile(SIM / "sets.json", build / "sets.json")
    (build / "gametables").mkdir()
    shutil.copyfile(SIM / "combatratings.txt", build / "gametables" / "combatratings.txt")
    write_manifest(
        build, build="9.9.9.9", product="wow_classic_beta", fetched_at="2026-01-01T00:00:00Z"
    )
    return build


def parsed(path: Path) -> pb.SimDatabase:
    database = pb.SimDatabase()
    database.ParseFromString(path.read_bytes())
    return database


def test_simdb_is_written_and_parses_back(build_dir: Path):
    path = write_sim_database("9.9.9.9", root=build_dir.parent)
    assert path == build_dir / "simdb.bin"
    database = parsed(path)
    ids = [item.id for item in database.items]
    assert ids == sorted(ids)
    assert len(database.items) == 7
    # 934 "Sword Skill +3" grants only a weapon skill through its equip spell,
    # which SimEnchant has no field for -- pipeline.simdb.enchants still
    # emits it, with empty stats, and warns (see test_simdb_enchants.py's
    # test_a_weapon_skill_only_equip_spell_is_dropped_with_a_warning).
    assert [row.effect_id for row in database.enchants] == [352, 803, 930, 931, 934, 1900, 2504]


def test_the_set_name_comes_from_the_normalized_sets_file(build_dir: Path):
    helm = next(item for item in parsed(write_sim_database("9.9.9.9", root=build_dir.parent)).items
                if item.id == 16866)
    assert (helm.set_id, helm.set_name) == (209, "Battlegear of Might")


def test_random_suffixes_are_deliberately_empty(build_dir: Path):
    assert list(parsed(write_sim_database("9.9.9.9", root=build_dir.parent)).random_suffixes) == []


def test_two_runs_are_byte_identical(build_dir: Path):
    first = write_sim_database("9.9.9.9", root=build_dir.parent).read_bytes()
    second = write_sim_database("9.9.9.9", root=build_dir.parent).read_bytes()
    assert first == second


def test_the_manifest_covers_the_new_files(build_dir: Path):
    write_sim_database("9.9.9.9", root=build_dir.parent)
    assert verify(build_dir) == []
    files = json.loads((build_dir / "manifest.json").read_text())["files"]
    assert "simdb.bin" in files
    assert "simconsumes.json" in files


def test_consumables_are_the_sidecar_the_engine_lane_needs(build_dir: Path):
    """SimDatabase has no consumable field -- the engine's consumes.go is an enum
    table. The pipeline emits the raw material for regenerating it instead."""
    raw = build_dir / "raw"
    records = build_consumables(
        read_csv(raw / "ItemSparse.csv"),
        read_csv(raw / "Item.csv"),
        read_csv(raw / "ItemEffect.csv"),
        read_csv(raw / "ItemXItemEffect.csv"),
    )
    assert [record.model_dump() for record in records] == [
        {
            "id": 40001,
            "name": "Greater Healing Potion",
            "quality": 1,
            "required_level": 35,
            "spell_ids": [900009],
        }
    ]


def test_a_missing_raw_directory_is_a_clear_error(tmp_path: Path):
    (tmp_path / "1.0.0.0").mkdir()
    with pytest.raises(SystemExit, match="fetch"):
        write_sim_database("1.0.0.0", root=tmp_path)


def test_a_build_that_was_never_normalized_is_a_clear_error(build_dir: Path):
    (build_dir / "sets.json").unlink()
    with pytest.raises(SystemExit, match="normalize"):
        write_sim_database("9.9.9.9", root=build_dir.parent)
