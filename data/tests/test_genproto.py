import os
from pathlib import Path

import pytest

from pipeline.genproto import PROTO_FILES, fix_imports
from pipeline.simproto import ENGINE_SHA, apl, pb

PROTO_DIR = Path("proto")


def test_fix_imports_rewrites_sibling_generated_modules():
    source = "import common_pb2 as common__pb2\nimport shaman_pb2 as shaman__pb2\n"
    assert fix_imports(source, "pipeline.simproto") == (
        "from pipeline.simproto import common_pb2 as common__pb2\n"
        "from pipeline.simproto import shaman_pb2 as shaman__pb2\n"
    )


def test_fix_imports_leaves_the_protobuf_runtime_alone():
    source = "from google.protobuf import descriptor as _descriptor\n"
    assert fix_imports(source, "pipeline.simproto") == source


def test_every_vendored_proto_is_present():
    for name in PROTO_FILES:
        assert (PROTO_DIR / name).is_file(), f"{name} is not vendored"
    assert (PROTO_DIR / "ENGINE_SHA").read_text().strip()


def test_the_engine_messages_this_lane_depends_on_exist():
    assert set(pb.SimDatabase.DESCRIPTOR.fields_by_name) == {
        "items",
        "enchants",
        "random_suffixes",
    }
    assert set(pb.SimItem.DESCRIPTOR.fields_by_name) == {
        "id",
        "class_allowlist",
        "name",
        "type",
        "armor_type",
        "weapon_type",
        "hand_type",
        "ranged_weapon_type",
        "stats",
        "weapon_damage_min",
        "weapon_damage_max",
        "weapon_speed",
        "bonus_physical_damage",
        "set_name",
        "set_id",
        "weapon_skills",
    }
    assert set(pb.SimEnchant.DESCRIPTOR.fields_by_name) == {"effect_id", "stats"}
    assert set(pb.ItemSpec.DESCRIPTOR.fields_by_name) == {"id", "random_suffix", "enchant"}
    assert {"type", "prepull_actions", "priority_list"} <= set(
        apl.APLRotation.DESCRIPTOR.fields_by_name
    )


def test_the_stat_enums_are_the_size_the_stat_map_assumes():
    # ENGINE_SHA 0900ba8b8 carries the engine's unified Hit/Crit collapse
    # (StatHit/StatCrit replace the separate melee/spell variants), so these
    # counts are lower than the pre-unification 44/16 the brief quotes.
    assert len(pb.Stat.keys()) == 41
    assert len(pb.WeaponSkill.keys()) == 16


def test_engine_sha_is_recorded_and_short():
    assert ENGINE_SHA.isalnum() and 7 <= len(ENGINE_SHA) <= 12


def test_vendored_protos_match_the_engine_checkout():
    """Skipped unless the engine checkout is available.

    CI does not clone the engine; a developer regenerating after an engine
    change gets a red test here until they re-run `python -m pipeline simproto`.
    """
    engine = os.environ.get("FOREVER_ENGINE_PATH")
    if not engine:
        pytest.skip("FOREVER_ENGINE_PATH is not set")
    for name in PROTO_FILES:
        expected = (Path(engine) / "proto" / name).read_bytes()
        assert (PROTO_DIR / name).read_bytes() == expected, (
            f"{name} drifted; rerun `python -m pipeline simproto --engine $FOREVER_ENGINE_PATH`"
        )
