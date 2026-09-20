import json
import shutil
from pathlib import Path

import pytest

from pipeline.csvio import read_csv
from pipeline.manifest import verify, write_manifest
from pipeline.normalize import write_json
from pipeline.normalize.items import normalize_items
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
    # items.json's `suffixes` and `faction_restriction` columns default empty
    # via the Item model -- the same shape a build normalized before `loot`
    # has run carries. _fork_columns only requires the keys to be present.
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
    assert len(database.items) == 8
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
    assert "simitems.json" in files


def test_simitems_json_is_the_kept_items_sorted_ids(build_dir: Path):
    """The web lane's own copy of simdb.bin's item universe (contract 10.1 A6's candidate
    filtering) -- every bulk candidate source there filters against it, so it has to name
    exactly the ids `simdb_item_rows` kept, not the planner's wider items/<class>.json set.
    """
    path = write_sim_database("9.9.9.9", root=build_dir.parent)
    database = parsed(path)
    simitems = json.loads((build_dir / "simitems.json").read_text(encoding="utf-8"))
    assert simitems == {
        "build": "9.9.9.9",
        "items": sorted(item.id for item in database.items),
    }


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


def test_a_build_missing_items_json_is_a_clear_error(build_dir: Path):
    (build_dir / "items.json").unlink()
    with pytest.raises(SystemExit, match="normalize"):
        write_sim_database("9.9.9.9", root=build_dir.parent)


def test_a_build_whose_loot_step_never_ran_is_a_clear_error(build_dir: Path):
    """items.json from before `pipeline loot` populated its two fork columns
    -- an older schema, or a hand-built build directory -- must not read as
    an honestly-empty build. The guard names the missing columns and tells
    the operator which command fills them."""
    rows = json.loads((build_dir / "items.json").read_text())
    for row in rows:
        del row["suffixes"]
        del row["faction_restriction"]
    (build_dir / "items.json").write_text(json.dumps(rows))
    with pytest.raises(SystemExit, match="loot"):
        write_sim_database("9.9.9.9", root=build_dir.parent)


def test_a_legitimately_empty_fork_column_is_not_mistaken_for_a_missing_one(
    build_dir: Path,
):
    """Every item in this fixture's items.json genuinely has no suffixes and
    no faction restriction -- the guard must not confuse that with the
    column being absent."""
    database = parsed(write_sim_database("9.9.9.9", root=build_dir.parent))
    assert all(not item.random_suffix_options for item in database.items)
    assert all(not item.faction_restriction for item in database.items)


def test_an_empty_items_json_is_a_clear_error(build_dir: Path):
    """A build directory with items.json present but empty must not read as
    an honestly-restriction-free build -- there is no build with zero items."""
    (build_dir / "items.json").write_text("[]")
    with pytest.raises(SystemExit, match="normalize"):
        write_sim_database("9.9.9.9", root=build_dir.parent)


def test_normalize_running_again_after_loot_is_a_clear_error(build_dir: Path):
    """`loot` writes suffixes.json alongside items.json's two fork columns
    (contract 10.8). If `normalize` then runs again, it overwrites items.json
    from the model defaults and blanks both columns back out, but leaves
    suffixes.json sitting in the build directory from the earlier `loot` run.
    That combination -- loot's own output present, every item unrestricted --
    is not a build that has never been looted; it is one that was looted and
    then had the columns wiped, so the guard must catch it too."""
    (build_dir / "suffixes.json").write_text("[]")
    with pytest.raises(SystemExit, match="loot"):
        write_sim_database("9.9.9.9", root=build_dir.parent)


def test_loot_json_alone_also_trips_the_guard(build_dir: Path):
    """Either of loot's own outputs is enough to prove loot ran; the guard
    does not require both files."""
    (build_dir / "loot.json").write_text("[]")
    with pytest.raises(SystemExit, match="loot"):
        write_sim_database("9.9.9.9", root=build_dir.parent)
