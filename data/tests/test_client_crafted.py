"""pipeline.client_crafted: the client's recipe chains for Forever-new crafts."""

from pathlib import Path

from pipeline.classicdb_crafted import ClassicDbCraftedRecipe
from pipeline.client_crafted import (
    CLIENT_CRAFTED,
    build_recipes,
    load_client_crafted,
    merge_recipes,
    write_client_crafted,
)

# Build 1.60.1.70790's own rows for Greenhammer (spell 1306535).
GREENHAMMER, PLANS, CRAFT_SPELL = 279261, 279310, 1306535
REAGENTS = [274030, 251290, 251291]


def _rows() -> dict[str, list[dict[str, str]]]:
    reagents = {f"Reagent_{i}": str(item_id) for i, item_id in enumerate(REAGENTS)}
    return {
        "spell_effects": [
            {"SpellID": str(CRAFT_SPELL), "Effect": "24", "EffectItemType": str(GREENHAMMER)},
            # A create-item spell with no reagents: a conjure, not a craft.
            {"SpellID": "5504", "Effect": "24", "EffectItemType": "5350"},
            {"SpellID": "133", "Effect": "2", "EffectItemType": "0"},
        ],
        "spell_reagents": [{"SpellID": str(CRAFT_SPELL), **reagents, "Reagent_3": "0"}],
        "item_effects": [
            {"ID": "236035", "TriggerType": "6", "SpellID": str(CRAFT_SPELL)},
            {"ID": "236007", "TriggerType": "0", "SpellID": "402265"},
        ],
        "item_x_effects": [
            {"ItemEffectID": "236035", "ItemID": str(PLANS)},
            {"ItemEffectID": "236007", "ItemID": str(PLANS)},
        ],
        "item_sparse": [{"ID": str(PLANS), "RequiredSkillRank": "300"}],
    }


def test_the_client_chain_names_the_plans_the_reagents_and_the_skill():
    assert build_recipes(**_rows()) == {
        GREENHAMMER: ClassicDbCraftedRecipe(
            recipe_item_ids=[PLANS], reagent_item_ids=sorted(REAGENTS), required_skill=300
        )
    }


def test_the_extract_round_trips_and_is_empty_before_the_first_run(tmp_path: Path):
    assert load_client_crafted(tmp_path) == {}
    recipes = build_recipes(**_rows())
    assert write_client_crafted(tmp_path, recipes) == tmp_path / CLIENT_CRAFTED
    assert load_client_crafted(tmp_path) == recipes


def test_classic_db_keeps_its_own_recipe_for_a_classic_craft():
    classic = {1: ClassicDbCraftedRecipe(recipe_item_ids=[2])}
    client = {1: ClassicDbCraftedRecipe(), 3: ClassicDbCraftedRecipe(reagent_item_ids=[4])}
    assert merge_recipes(classic, client) == {1: classic[1], 3: client[3]}
