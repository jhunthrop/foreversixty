"""pipeline.loot.crafted_skill: each crafted item's recipe skill, from the
classic-db crafted extract onto loot.json's crafted sources."""

from pathlib import Path

import pytest

from conftest import ACTIVE_BUILD
from pipeline.classicdb_crafted import ClassicDbCraftedRecipe, load_extract
from pipeline.loot.crafted_skill import apply_crafted_skill
from pipeline.models import LootBoss, LootFile, LootSource

EMBOSSED_LEATHER_VEST = 2300
FOREVER_TWIN = 250482


def _crafted(items, **extra) -> LootSource:
    return LootSource(
        id="crafted:leatherworking",
        kind="crafted",
        name="Leatherworking",
        profession="leatherworking",
        items=items,
        **extra,
    )


def test_each_crafted_item_gets_its_recipe_skill_and_an_unknown_one_is_left_out():
    document = LootFile(sources=[_crafted([EMBOSSED_LEATHER_VEST, 2309, 2318])])
    recipes = {
        EMBOSSED_LEATHER_VEST: ClassicDbCraftedRecipe(required_skill=40),
        2309: ClassicDbCraftedRecipe(),
    }
    out = apply_crafted_skill(document, recipes)
    assert out.sources[0].required_skill == {"2300": 40}


def test_a_reitemised_forever_id_takes_its_classic_twins_skill():
    document = LootFile(
        sources=[
            _crafted([FOREVER_TWIN]),
            LootSource(
                id="raid:x",
                kind="raid",
                name="x",
                bosses=[
                    LootBoss(
                        id="raid:x:1",
                        name="b",
                        npc_id=1,
                        items=[FOREVER_TWIN],
                        reitemised_from={str(FOREVER_TWIN): EMBOSSED_LEATHER_VEST},
                    )
                ],
            ),
        ]
    )
    out = apply_crafted_skill(
        document, {EMBOSSED_LEATHER_VEST: ClassicDbCraftedRecipe(required_skill=40)}
    )
    assert out.sources[0].required_skill == {str(FOREVER_TWIN): 40}


def test_a_source_with_no_known_skill_carries_no_map_and_other_kinds_are_untouched():
    vendor = LootSource(id="vendor:1", kind="vendor", name="v", npc_id=1, items=[5])
    out = apply_crafted_skill(LootFile(sources=[_crafted([1]), vendor]), {})
    assert out.sources[0].required_skill is None
    assert out.sources[1] == vendor


def test_the_pinned_dump_says_embossed_leather_vest_needs_leatherworking_40():
    """cmangos/classic-db ec4f5961: npc_trainer teaches spell 2883 (which
    learns create spell 2160, Embossed Leather Vest) at reqskillvalue 40.
    Skipped until `fetch-classic-sources` has rewritten the committed
    extract with `required_skill`."""
    recipes = load_extract(Path("builds") / ACTIVE_BUILD)
    recipe = recipes.get(EMBOSSED_LEATHER_VEST)
    if recipe is None or all(r.required_skill is None for r in recipes.values()):
        pytest.skip("the committed crafted extract predates required_skill")
    assert recipe.required_skill == 40
