# data/tests/test_classicdb_crafted.py
"""`pipeline.classicdb_crafted` -- the recipe item / reagent chain
behind a crafted item's `SPELL_EFFECT_CREATE_ITEM` spell, from a small
synthetic mysqldump snippet shaped after the real Sulfuron Hammer chain
this lane's own report verifies against the pinned dump: item 18592
"Plans: Sulfuron Hammer" has `spellid_1` 23007 ("Learn: Plans..."),
spell 23007 is `Effect1` 36 (`SPELL_EFFECT_LEARN_SPELL`) with
`EffectTriggerSpell1` 21161, and spell 21161 is `Effect1` 24
(`SPELL_EFFECT_CREATE_ITEM`) `EffectItemType1` 17193 (Sulfuron Hammer
itself) with `Reagent1`/`ReagentCount1` 17203/8 (Sulfuron Ingot x8).

`ClassicDbDump.LEARN_SPELL_EFFECT`/`SPELL_EFFECT_CREATE_ITEM`'s own
values are reused directly (not hardcoded here) so a real schema
constant change fails this file first.
"""

from pipeline.audit.dumpdb import SPELL_EFFECT_CREATE_ITEM, ClassicDbDump
from pipeline.classicdb_crafted import ClassicDbCraftedRecipe, extract_records

SULFURON_HAMMER = 17193
PLANS_SULFURON_HAMMER = 18592
LEARN_PLANS_SPELL = 23007
CREATE_HAMMER_SPELL = 21161
SULFURON_INGOT = 17203
DARK_IRON_BAR = 11371

# A trainer-taught recipe with no item at all: item 900 is created
# directly by spell 8002 (no "learn" spell/item precedes it).
TRAINER_ITEM = 900
TRAINER_CREATE_SPELL = 8002

# A recipe taught by TWO alternative items (a rare but real shape --
# `pipeline.classicdb_crafted`'s own module doc: "almost always exactly
# one").
DUAL_RECIPE_ITEM = 950
DUAL_CREATE_SPELL = 8102
DUAL_LEARN_SPELL = 8101
DUAL_RECIPE_A = 951
DUAL_RECIPE_B = 952

_ITEM_TEMPLATE_COLUMNS = ["entry", "spellid_1", "spellid_2", "spellid_3", "spellid_4", "spellid_5"]
_SPELL_TEMPLATE_COLUMNS = [
    "Id", "ProcChance",
    "Effect1", "Effect2", "Effect3",
    "EffectApplyAuraName1", "EffectApplyAuraName2", "EffectApplyAuraName3",
    "EffectBasePoints1", "EffectBasePoints2", "EffectBasePoints3",
    "EffectDieSides1", "EffectDieSides2", "EffectDieSides3",
    "EffectMiscValue1", "EffectMiscValue2", "EffectMiscValue3",
    "EffectTriggerSpell1", "EffectTriggerSpell2", "EffectTriggerSpell3",
    "EffectItemType1", "EffectItemType2", "EffectItemType3",
    "Reagent1", "Reagent2", "Reagent3", "Reagent4",
    "Reagent5", "Reagent6", "Reagent7", "Reagent8",
    "ReagentCount1", "ReagentCount2", "ReagentCount3", "ReagentCount4",
    "ReagentCount5", "ReagentCount6", "ReagentCount7", "ReagentCount8",
]  # fmt: skip


def _item_row(entry: int, spellid_1: int = 0) -> str:
    return f"({entry},{spellid_1},0,0,0,0)"


def _spell_row(
    spell_id: int,
    effect1: int = 0,
    effect_item_type1: int = 0,
    trigger_spell1: int = 0,
    reagent1: int = 0,
    reagent_count1: int = 0,
) -> str:
    return (
        f"({spell_id},101,"
        f"{effect1},0,0,"
        f"0,0,0,"
        f"0,0,0,"
        f"1,1,1,"
        f"0,0,0,"
        f"{trigger_spell1},0,0,"
        f"{effect_item_type1},0,0,"
        f"{reagent1},0,0,0,0,0,0,0,"
        f"{reagent_count1},0,0,0,0,0,0,0)"
    )


_ITEM_ROWS = ",\n  ".join(
    [
        _item_row(PLANS_SULFURON_HAMMER, LEARN_PLANS_SPELL),
        _item_row(DUAL_RECIPE_A, DUAL_LEARN_SPELL),
        _item_row(DUAL_RECIPE_B, DUAL_LEARN_SPELL),
    ]
)
_SPELL_ROWS = ",\n  ".join(
    [
        _spell_row(
            LEARN_PLANS_SPELL,
            effect1=ClassicDbDump.LEARN_SPELL_EFFECT,
            trigger_spell1=CREATE_HAMMER_SPELL,
        ),
        _spell_row(
            CREATE_HAMMER_SPELL,
            effect1=SPELL_EFFECT_CREATE_ITEM,
            effect_item_type1=SULFURON_HAMMER,
            reagent1=SULFURON_INGOT,
            reagent_count1=8,
        ),
        _spell_row(
            TRAINER_CREATE_SPELL,
            effect1=SPELL_EFFECT_CREATE_ITEM,
            effect_item_type1=TRAINER_ITEM,
            reagent1=DARK_IRON_BAR,
            reagent_count1=20,
        ),
        _spell_row(
            DUAL_LEARN_SPELL,
            effect1=ClassicDbDump.LEARN_SPELL_EFFECT,
            trigger_spell1=DUAL_CREATE_SPELL,
        ),
        _spell_row(
            DUAL_CREATE_SPELL,
            effect1=SPELL_EFFECT_CREATE_ITEM,
            effect_item_type1=DUAL_RECIPE_ITEM,
        ),
    ]
)
_ITEM_TEMPLATE_SCHEMA = ",\n  ".join(f"`{c}`" for c in _ITEM_TEMPLATE_COLUMNS)
_SPELL_TEMPLATE_SCHEMA = ",\n  ".join(f"`{c}`" for c in _SPELL_TEMPLATE_COLUMNS)

SAMPLE_SQL = f"""
CREATE TABLE `item_template` (
  {_ITEM_TEMPLATE_SCHEMA}
) ENGINE=MyISAM;
INSERT INTO `item_template` VALUES
  {_ITEM_ROWS};

CREATE TABLE `spell_template` (
  {_SPELL_TEMPLATE_SCHEMA}
) ENGINE=MyISAM;
INSERT INTO `spell_template` VALUES
  {_SPELL_ROWS};
"""


def test_sulfuron_hammer_shaped_chain_resolves_recipe_item_and_reagent():
    records = extract_records(SAMPLE_SQL)
    assert records[SULFURON_HAMMER] == ClassicDbCraftedRecipe(
        recipe_item_ids=[PLANS_SULFURON_HAMMER], reagent_item_ids=[SULFURON_INGOT]
    )


def test_trainer_taught_recipe_has_no_recipe_item_at_all():
    """No item's own `spellid_1..5` names `TRAINER_CREATE_SPELL`'s
    "learn" spell (there isn't one) -- `recipe_item_ids` is empty, the
    trainer-taught contract `ClassicDbCraftedRecipe`'s own doc states."""
    records = extract_records(SAMPLE_SQL)
    assert records[TRAINER_ITEM] == ClassicDbCraftedRecipe(
        recipe_item_ids=[], reagent_item_ids=[DARK_IRON_BAR]
    )


def test_two_alternative_recipe_items_both_resolve():
    """Two items (951, 952) both teach the SAME learn spell -- both
    belong in `recipe_item_ids` (an OR: either one obtained is enough,
    `apply_crafted_opens_gate`'s own doc)."""
    records = extract_records(SAMPLE_SQL)
    assert records[DUAL_RECIPE_ITEM].recipe_item_ids == [DUAL_RECIPE_A, DUAL_RECIPE_B]


def test_an_item_with_no_create_item_spell_at_all_is_absent():
    records = extract_records(SAMPLE_SQL)
    assert 999999 not in records


def test_write_and_load_extract_round_trip(tmp_path):
    from pipeline.classicdb_crafted import load_extract, write_extract

    records = extract_records(SAMPLE_SQL)
    build_dir = tmp_path / "1.60.1.99999"
    path = write_extract(build_dir, records, source_commit="deadbeef")
    assert path.exists()
    loaded = load_extract(build_dir)
    assert loaded == records


def test_load_extract_returns_empty_when_nothing_was_fetched_yet(tmp_path):
    from pipeline.classicdb_crafted import load_extract

    assert load_extract(tmp_path / "1.60.1.99999") == {}
