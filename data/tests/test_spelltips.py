import json
from pathlib import Path

from pipeline.spelltips import _Tables, consumable_tips, racial_tips, write_spell_tips


def _csv(path: Path, header: str, rows: list[str]) -> None:
    path.write_text("\n".join([header, *rows]) + "\n", encoding="utf-8")


def _raw(raw: Path) -> None:
    """Mortal Strike rank 4, a channeled Blizzard, Blood Fury on the orc racial line and a
    flask's on-use spell, with only the columns the tooltip reads."""
    raw.mkdir(parents=True)
    _csv(
        raw / "SpellName.csv",
        "ID,Name_lang",
        ["21553,Mortal Strike", "10,Blizzard", "20572,Blood Fury", "17626,Flask of the Titans"],
    )
    _csv(
        raw / "Spell.csv",
        "ID,NameSubtext_lang,Description_lang",
        [
            '21553,Rank 4,"A vicious strike that deals weapon damage plus $s1 and wounds the '
            'target for $d."',
            '10,Rank 1,"Ice shards pelt the target area within $a1 yards."',
            '20572,Racial,"Increases Attack Power by $s1% for $d."',
            '17626,,"Increases the player\'s maximum health by $s1 for $d."',
        ],
    )
    _csv(
        raw / "SpellMisc.csv",
        "ID,Attributes_0,Attributes_1,DifficultyID,CastingTimeIndex,DurationIndex,RangeIndex,"
        "SpellIconFileDataID,SpellID",
        [
            "1,0,0,0,1,1,2,132355,21553",
            "2,0,64,0,1,2,3,0,10",
            "3,0,0,0,1,3,1,0,20572",
            "4,0,0,0,1,4,1,0,17626",
        ],
    )
    _csv(
        raw / "SpellEffect.csv",
        "ID,DifficultyID,SpellID,EffectIndex,EffectBasePoints,EffectDieSides,"
        "EffectAuraPeriod,EffectRadiusIndex_0,EffectRadiusIndex_1,EffectChainTargets",
        [
            "1,0,21553,0,159,1,0,0,0,0",
            # The radius rides in index 1 with index 0 unset, as Classic's rows often do.
            "2,0,10,0,24,1,0,0,13,0",
            "3,0,20572,0,24,1,0,0,0,0",
            "4,0,17626,0,1199,1,0,0,0,0",
        ],
    )
    _csv(
        raw / "ManifestInterfaceData.csv",
        "ID,FilePath,FileName",
        ["132355,Interface/Icons/,Ability_Warrior_SavageBlow.blp"],
    )
    _csv(raw / "SpellRadius.csv", "ID,Radius,RadiusPerLevel,RadiusMin,RadiusMax", ["13,10,0,0,10"])
    _csv(
        raw / "SpellRange.csv",
        "ID,DisplayName_lang,RangeMin_0,RangeMax_0",
        ["1,Self,0,0", "2,Combat,0,5", "3,Long,0,30"],
    )
    _csv(
        raw / "SpellDuration.csv",
        "ID,Duration,MaxDuration",
        ["1,10000,10000", "2,8000,8000", "3,15000,15000", "4,7200000,7200000"],
    )
    _csv(raw / "SpellCastTimes.csv", "ID,Base,Minimum", ["1,0,0"])
    _csv(
        raw / "SpellCooldowns.csv",
        "ID,DifficultyID,CategoryRecoveryTime,RecoveryTime,SpellID",
        ["1,0,6000,0,21553", "2,0,0,120000,20572"],
    )
    _csv(
        raw / "SpellPower.csv",
        "ID,ManaCost,PowerType,PowerCostPct,DifficultyID,SpellID",
        ["1,300,1,0,0,21553", "2,320,0,0,0,10"],
    )
    _csv(
        raw / "SkillLine.csv",
        "DisplayName_lang,ID,CategoryID",
        ["Orc Racial,125,9", "Cooking,185,9"],
    )
    _csv(raw / "SkillLineAbility.csv", "ID,SkillLine,Spell", ["1,125,20572", "2,185,21553"])
    _csv(
        raw / "ItemEffect.csv",
        "ID,LegacySlotIndex,TriggerType,Charges,CoolDownMSec,CategoryCoolDownMSec,"
        "SpellCategoryID,SpellID,ChrSpecializationID,PlayerConditionID",
        ["50,0,0,-1,-1,-1,0,17626,0,0", "51,0,1,0,0,0,0,21553,0,0"],
    )
    _csv(raw / "ItemXItemEffect.csv", "ID,ItemEffectID,ItemID", ["1,50,13510", "2,51,999"])


def test_a_class_spell_reads_as_the_client_tooltip(tmp_path):
    raw = tmp_path / "raw"
    _raw(raw)
    tip = _Tables(raw).tip(21553)
    assert tip is not None
    assert tip.model_dump() == {
        "name": "Mortal Strike",
        "rank": 4,
        "cost": 300,
        "cost_type": 1,
        "cost_pct": 0.0,
        "cast_time_ms": 0,
        "channeled": False,
        "passive": False,
        "cooldown_ms": 6000,
        "range_yd": 5.0,
        "description": "A vicious strike that deals weapon damage plus 160 and wounds the "
        "target for 10 sec.",
        "icon": "ability_warrior_savageblow",
    }


def test_a_channel_is_marked_and_its_index_1_radius_resolves(tmp_path):
    raw = tmp_path / "raw"
    _raw(raw)
    tip = _Tables(raw).tip(10)
    assert tip is not None
    assert tip.channeled is True
    assert tip.range_yd == 30.0
    assert tip.description == "Ice shards pelt the target area within 10 yards."
    # No icon file in the client: its placeholder art, never an empty name.
    assert tip.icon != ""


def test_racials_come_only_from_the_racial_skill_lines(tmp_path):
    raw = tmp_path / "raw"
    _raw(raw)
    tips = racial_tips(raw, _Tables(raw))
    assert list(tips) == ["Blood Fury"]
    assert tips["Blood Fury"].description == "Increases Attack Power by 25% for 15 sec."
    assert tips["Blood Fury"].cooldown_ms == 120000


def test_a_consumable_is_its_one_named_item_and_on_use_spell(tmp_path):
    raw = tmp_path / "raw"
    _raw(raw)
    simbuffs = {
        "flask_of_the_titans": {"name": "Flask of the Titans"},
        "no_such_item": {"name": "Nowhere Brew"},
    }
    items = [{"id": 13510, "name": "Flask of the Titans"}, {"id": 999, "name": "Mortal Strike"}]
    tips = consumable_tips(raw, simbuffs, items, _Tables(raw))
    assert {k: v.model_dump() for k, v in tips.items()} == {
        "flask_of_the_titans": {
            "item_id": 13510,
            "spell_id": 17626,
            "use": "Increases the player's maximum health by 1200 for 120 min.",
        }
    }


def test_write_spell_tips_writes_one_file_per_class_and_the_shared_file(tmp_path):
    build_dir = tmp_path / "1.0.0.1"
    _raw(build_dir / "raw")
    (build_dir / "trainables").mkdir()
    (build_dir / "trainables" / "warrior.json").write_text(
        json.dumps(
            {
                "build": "1.0.0.1",
                "class_slug": "warrior",
                "trainables": [
                    {"name": "Mortal Strike", "ranks": [{"id": 21553, "rank": 4, "level": 60}]}
                ],
            }
        ),
        encoding="utf-8",
    )
    (build_dir / "enchants.json").write_text(
        json.dumps([{"id": 1900, "spell_id": 10}, {"id": 1, "spell_id": 0}]), encoding="utf-8"
    )
    (build_dir / "items.json").write_text(json.dumps([]), encoding="utf-8")
    out = write_spell_tips("1.0.0.1", root=tmp_path)
    warrior = json.loads((out / "warrior.json").read_text(encoding="utf-8"))
    assert list(warrior["spells"]) == ["21553"]
    shared = json.loads((out / "shared.json").read_text(encoding="utf-8"))
    assert list(shared["enchants"]) == ["1900"]
    assert shared["consumables"] == {}
    assert list(shared["racials"]) == ["Blood Fury"]
