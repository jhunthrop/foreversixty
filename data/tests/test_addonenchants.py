"""pipeline.addonenchants: the enchant each spec's BiS picks wear, packed for the Gear tab."""

import json
from pathlib import Path

import pytest

from pipeline.addonenchants import AddonEnchantError, build_enchant_plans

BUILD = "1.60.1.70374"

_ENCHANTS = [
    {"id": 1891, "name": "Enchant Chest - Greater Stats", "icon": "spell_holy_greaterheal",
     "slots": ["chest"], "spell_id": 20025, "stats": {"agility": 4, "stamina": 4}},
    {"id": 904, "name": "Enchant Boots - Agility", "icon": "i", "slots": ["feet"], "spell_id": 1,
     "stats": {"agility": 5}},
    {"id": 904, "name": "Enchant Necklace - Agility", "icon": "i", "slots": ["neck"], "spell_id": 2,
     "stats": {"agility": 5}},
    {"id": 664, "name": "Sniper Scope", "icon": "s", "slots": ["ranged"], "spell_id": 12460,
     "item_id": 10548, "stats": {"ranged_attack_power": 7}},
    {"id": 1, "name": "Healing kit", "icon": "h", "slots": ["legs"], "spell_id": 3,
     "stats": {"healing": 8, "bonus_armor": 5}},
]  # fmt: skip


def _write(root: Path, specs: dict, profiles: dict) -> tuple[Path, Path]:
    (root / BUILD).mkdir(parents=True)
    (root / BUILD / "enchants.json").write_text(json.dumps(_ENCHANTS), encoding="utf-8")
    (root / BUILD / "items.json").write_text(
        json.dumps([{"id": 10548, "name": "Sniper Scope", "required_level": 40}]), encoding="utf-8"
    )
    curated = root / "curated"
    curated.mkdir()
    (curated / "raid-enchants.json").write_text(
        json.dumps({"profiles": profiles, "specs": specs}), encoding="utf-8"
    )
    return root, curated


def test_a_plan_is_its_profile_overlaid_by_the_spec_with_shared_records(tmp_path):
    profiles = {
        "hunter": {
            "chest": {"id": 1891},
            "neck": {"id": 904},
            "feet": {"id": 904},
            "legs": {"id": 1},
        }
    }
    specs = {
        "hunter-marksmanship": {
            "profile": "hunter",
            "slots": {"ranged": {"id": 664}, "legs": None},
        },
        "hunter-survival": {"profile": "hunter"},
    }
    root, curated = _write(tmp_path, specs, profiles)
    plans, records = build_enchant_plans(
        root, BUILD, curated, ["hunter-marksmanship", "hunter-survival"]
    )
    assert records == [
        "1891|Enchant Chest - Greater Stats|spell_holy_greaterheal|20025|0|0|agility=4,stamina=4",
        "904|Enchant Boots - Agility|i|1|0|0|agility=5",
        "904|Enchant Necklace - Agility|i|2|0|0|agility=5",
        "664|Sniper Scope|s|12460|10548|40|ranged_attack_power=7",
        "1|Healing kit|h|3|0|0|armor=5,healing_power=8",
    ]
    assert plans["hunter-marksmanship"] == {"chest": 1, "feet": 2, "neck": 3, "ranged": 4}
    assert plans["hunter-survival"] == {"chest": 1, "feet": 2, "legs": 5, "neck": 3}


def test_an_enchant_that_does_not_fit_its_slot_refuses(tmp_path):
    root, curated = _write(
        tmp_path, {"mage-fire": {"profile": "caster"}}, {"caster": {"head": {"id": 1891}}}
    )
    with pytest.raises(AddonEnchantError, match="1891 for head"):
        build_enchant_plans(root, BUILD, curated, ["mage-fire"])


def test_the_tables_render_into_data_lua(tmp_path):
    from pipeline.addonlua import render_lua
    from pipeline.models import AddonData

    lua = render_lua(
        AddonData(
            build="x",
            classes={},
            weights={},
            enchant_plans={"hunter-marksmanship": {"ranged": 1, "chest": 2}},
            enchant_records=[
                "664|Sniper Scope|s|12460|10548|40|ranged_attack_power=7",
                "1891|x|i|1|0|0|",
            ],
        )
    )
    assert '\t\t["hunter-marksmanship"] = { chest = 2, ranged = 1 },' in lua
    assert '\t\t"664|Sniper Scope|s|12460|10548|40|ranged_attack_power=7",' in lua
