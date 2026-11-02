"""ItemSparse stat 45 is spell power: it raises damage AND healing.

The 1.60 client states three spell bonuses (sim audit round 8, items lane):

* 42, "Increases damage done by magical spells and effects by up to N"
  (Moss Cinch 6911): damage only.
* 41, healing only (Death Speaker Scepter 2816).
* 45, "Increases damage and healing done by magical spells and effects by
  up to N" (Dreamweave Gloves 10019, Gnoll Casting Gloves 892): both.

Wowhead Forever's tooltips add stat 45 into the healing total wherever it
sits beside 41 or 42: Atiesh 22631 (45 = 120, 41 = 180) reads "healing done
by up to 300 and damage done by up to 120"; Field Marshal's Satin Hood
231622 (45 = 21, 41 = 41) reads "healing 62 and damage 21"; Earthen Silk
Slippers 254013 (45 = 18, 42 = 6) reads "healing 18 and damage 24". The
engine heals with the healing stat alone (`Spell.HealingPower`), so stat 45
mapped to spell power only gave 1,569 client rows no healing at all.
Wowhead's feed names stat 45 `splpwr`; on a record with no `splheal` it is
the same damage-and-healing bonus.
"""

from pathlib import Path

from test_normalize_gear import fixture_curves, fixture_icons
from test_simdb_items import RATING_FACTORS, curves, weapon_curves

from pipeline import wowhead_items as wh
from pipeline.csvio import read_csv
from pipeline.normalize.gear import build_class_items
from pipeline.simdb.items import build_sim_items, build_wowhead_sim_items, simdb_item_rows
from pipeline.simproto import pb

HERE = Path(__file__).parent
FIXTURES = HERE / "fixtures"

#: Helm of Might in tests/fixtures/ItemSparse_1_60.csv: budget 47, stamina in
#: slot 0 at editor 7440 (35), strength in slot 1 at 3200 (15).
HELM = "16866"


def _helm_rows(second_stat: str) -> list[dict[str, str]]:
    rows = read_csv(FIXTURES / "ItemSparse_1_60.csv")
    for row in rows:
        if row["ID"] == HELM:
            row["StatModifier_bonusStat_1"] = second_stat
    return rows


def _planner_helm(second_stat: str) -> dict[str, int]:
    records = build_class_items(
        _helm_rows(second_stat),
        read_csv(FIXTURES / "Item.csv"),
        read_csv(FIXTURES / "ChrClasses.csv"),
        fixture_icons(),
        "1.60.1.69893",
        fixture_curves(),
    )
    return next(i for r in records for i in r.items if i.id == int(HELM)).stats


def _sim_helm(second_stat: str) -> pb.SimItem:
    pairs = simdb_item_rows(_helm_rows(second_stat), read_csv(FIXTURES / "Item.csv"))
    items = build_sim_items(pairs, {}, {}, curves(), weapon_curves(), RATING_FACTORS, {})
    return next(item for item in items if item.id == int(HELM))


def _stat(item: pb.SimItem, name: str) -> float:
    index = pb.Stat.Value(name)
    return item.stats[index] if index < len(item.stats) else 0.0


def test_stat_45_is_spell_power_and_healing_for_the_planner() -> None:
    assert _planner_helm("45") == {"stamina": 35, "spell_power": 15, "healing": 15}


def test_stat_42_stays_damage_only() -> None:
    assert _planner_helm("42") == {"stamina": 35, "spell_power": 15}


def test_stat_45_reaches_the_engine_as_spell_power_and_healing() -> None:
    helm = _sim_helm("45")
    assert _stat(helm, "StatSpellPower") == 15.0
    assert _stat(helm, "StatHealingPower") == 15.0


def test_stat_45_adds_to_a_stated_healing_bonus() -> None:
    """Field Marshal's Satin Hood's shape: 41 and 45 on one row; the healing
    the tooltip shows is the two summed."""
    rows = _helm_rows("45")
    for row in rows:
        if row["ID"] == HELM:
            row["StatModifier_bonusStat_0"] = "41"
    stats = build_class_items(
        rows,
        read_csv(FIXTURES / "Item.csv"),
        read_csv(FIXTURES / "ChrClasses.csv"),
        fixture_icons(),
        "1.60.1.69893",
        fixture_curves(),
    )
    helm = next(i for r in stats for i in r.items if i.id == int(HELM))
    assert helm.stats == {"healing": 35 + 15, "spell_power": 15}


def _wowhead_cloak(stats: dict[str, float]) -> wh.WowheadItem:
    return wh.WowheadItem(
        id=990002,
        name="Trial Cape",
        quality=3,
        item_level=50,
        required_level=45,
        class_id=4,
        subclass_id=1,
        inventory_type=16,
        icon="inv_misc_cape_01",
        class_mask=None,
        stats=stats,
        set_id=None,
        unique=False,
        armor=0,
        damage_min=0,
        damage_max=0,
        speed=0.0,
        dps=0.0,
    )


def test_a_wowhead_splpwr_alone_is_spell_power_and_healing() -> None:
    """Spritecaster Cape 11623: `splpwr` 14, no `splheal`."""
    assert wh.planner_stats(_wowhead_cloak({"splpwr": 14})) == {"spell_power": 14, "healing": 14}


def test_a_wowhead_splheal_is_the_whole_healing_bonus() -> None:
    """Gemburst Circlet 10751: `splheal` 9 beside `splpwr` 4 reads "healing 9
    and damage 4"; the healing is not 13."""
    assert wh.planner_stats(_wowhead_cloak({"splpwr": 4, "splheal": 9})) == {
        "spell_power": 4,
        "healing": 9,
    }


def test_a_wowhead_splpwr_alone_reaches_the_engine_as_healing() -> None:
    (cloak,) = build_wowhead_sim_items([_wowhead_cloak({"splpwr": 14})], RATING_FACTORS)
    assert _stat(cloak, "StatSpellPower") == 14.0
    assert _stat(cloak, "StatHealingPower") == 14.0
