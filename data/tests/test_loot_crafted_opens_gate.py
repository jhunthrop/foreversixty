# data/tests/test_loot_crafted_opens_gate.py
"""data-followups-3 lane, 2026-09-30, item 1: `pipeline.loot.sources.
apply_crafted_opens_gate` gives a crafted item's `opens` the most
restrictive gate among its recipe item(s) and its reagents -- the eighth
wow-player sweep's own finding: Sulfuron Hammer is `crafted:blacksmithing`
with no `opens` even though its Plans are a Molten Core quest reward and
one reagent (Sulfuron Ingot) is a Molten Core boss drop, so a fresh 60
reads it as launch content.

Built by hand, like `test_loot_quest_opens_gate.py` beside this file --
`apply_crafted_opens_gate` only reads `LootFile.sources`/`.quests` and a
`item id -> ClassicDbCraftedRecipe` map, never a fork/classic-db/wowhead
merge, so no fixture engine dir is needed.
"""

from pipeline.classicdb_crafted import ClassicDbCraftedRecipe
from pipeline.loot.sources import apply_crafted_opens_gate, item_effective_gate
from pipeline.models import LootFile, LootSource, QuestSource

# Real ids (verified against the pinned classic-db dump and this build's
# own committed items.json/loot.json while building this lane's report):
SULFURON_HAMMER = 17193
PLANS_SULFURON_HAMMER = 18592
SULFURON_INGOT = 17203
FIERY_CORE = 17010
DARK_IRON_BAR = 11371  # a reagent with NO known loot.json source at all

DARK_IRON_PLATE = 20518
PLANS_DARK_IRON_PLATE = 20520  # BRD Thorium Brotherhood rep vendor, launch
COMMON_REAGENT = 3577  # an unrestricted, always-available reagent

LIONHEART_HELM = 16993
PLANS_LIONHEART_HELM = 16992  # world-drop plans, launch


def _quest_source(quest_id: int, opens: str | None = None) -> QuestSource:
    return QuestSource(
        quest_id=quest_id, name=f"quest {quest_id}", faction="both", min_level=60, level=60,
        level_source="classic-db", opens=opens,
    )  # fmt: skip


def test_sulfuron_hammer_gains_the_recipes_own_quest_gate_and_the_reagents_raid_gate():
    """Plans (quest-taught, opens "later") AND a raid-drop reagent (also
    "later") both gate -- the most restrictive of the two is still
    "later", the same value either alone would give, but the item must
    be gated at all (today it is not -- this is the defect)."""
    document = LootFile(
        sources=[
            LootSource(
                id="crafted:blacksmithing", kind="crafted", name="Blacksmithing",
                profession="blacksmithing", items=[SULFURON_HAMMER],
            ),
            LootSource(
                id="raid:molten-core", kind="raid", name="Molten Core", opens="later",
                items=[SULFURON_INGOT],
            ),
        ],
        quests={str(PLANS_SULFURON_HAMMER): [_quest_source(7604, opens="later")]},
    )
    classic_crafted = {
        SULFURON_HAMMER: ClassicDbCraftedRecipe(
            recipe_item_ids=[PLANS_SULFURON_HAMMER], reagent_item_ids=[SULFURON_INGOT]
        )
    }
    gated = apply_crafted_opens_gate(document, classic_crafted)
    gated_source = next(s for s in gated.sources if s.id == "crafted:blacksmithing:later")
    assert gated_source.kind == "crafted"
    assert gated_source.profession == "blacksmithing"
    assert gated_source.opens == "later"
    assert gated_source.items == [SULFURON_HAMMER]
    # The original bucket is now empty (its one item gated away) and
    # dropped entirely -- "a source the filter emptied is not a source",
    # this file's own module doc.
    assert not any(s.id == "crafted:blacksmithing" for s in gated.sources)


def test_a_reagent_with_no_known_source_at_all_is_never_invented_into_a_gate():
    """Dark Iron Bar (a reagent) names no `LootSource` at all in this
    document -- tenet 8's own "never invented": it must not gate the
    item, exactly `item_effective_gate`'s own "no source" case."""
    document = LootFile(
        sources=[
            LootSource(
                id="crafted:blacksmithing", kind="crafted", name="Blacksmithing",
                profession="blacksmithing", items=[SULFURON_HAMMER],
            ),
        ],
        quests={str(PLANS_SULFURON_HAMMER): [_quest_source(7604, opens=None)]},
    )
    classic_crafted = {
        SULFURON_HAMMER: ClassicDbCraftedRecipe(
            recipe_item_ids=[PLANS_SULFURON_HAMMER], reagent_item_ids=[DARK_IRON_BAR]
        )
    }
    gated = apply_crafted_opens_gate(document, classic_crafted)
    assert gated.sources == document.sources  # untouched: nothing gated


def test_a_reagent_with_even_one_ungated_way_does_not_gate_the_item():
    """Fiery Core's own real shape: a Molten Core drop AND an open-world
    boss (Firelord) kill -- the world boss way is un-gated, so the
    reagent never blocks the item even though ONE of its sources is a
    gated raid."""
    document = LootFile(
        sources=[
            LootSource(
                id="crafted:blacksmithing", kind="crafted", name="Blacksmithing",
                profession="blacksmithing", items=[SULFURON_HAMMER],
            ),
            LootSource(id="raid:molten-core", kind="raid", name="Molten Core", opens="later",
                       items=[FIERY_CORE]),
            LootSource(id="world:firelord", kind="world", name="Firelord", items=[FIERY_CORE]),
        ],
        quests={str(PLANS_SULFURON_HAMMER): [_quest_source(7604, opens=None)]},
    )  # fmt: skip
    classic_crafted = {
        SULFURON_HAMMER: ClassicDbCraftedRecipe(
            recipe_item_ids=[PLANS_SULFURON_HAMMER], reagent_item_ids=[FIERY_CORE]
        )
    }
    gated = apply_crafted_opens_gate(document, classic_crafted)
    assert gated.sources == document.sources


def test_trainer_taught_recipe_is_never_gated_by_the_recipe_side_at_all():
    """`recipe_item_ids` empty means no item stands between a character
    and the recipe (it is simply trained) -- only the reagents can gate
    such an item."""
    document = LootFile(
        sources=[
            LootSource(
                id="crafted:blacksmithing", kind="crafted", name="Blacksmithing",
                profession="blacksmithing", items=[SULFURON_HAMMER],
            ),
        ],
    )
    classic_crafted = {
        SULFURON_HAMMER: ClassicDbCraftedRecipe(recipe_item_ids=[], reagent_item_ids=[])
    }
    gated = apply_crafted_opens_gate(document, classic_crafted)
    assert gated.sources == document.sources


def test_dark_iron_plate_stays_open_launch_content_bg_thorium_brotherhood_plans():
    """Dark Iron Plate's own real shape: the Plans are a BRD Thorium
    Brotherhood reputation-vendor sale (launch content, no `opens` --
    `sim/cmd/leveling-bis/data.go`'s own `repFactionRaidPhaseOpens` names
    no gate for that faction) and its reagents are ordinary, unrestricted
    ones -- must NOT gain a gate."""
    document = LootFile(
        sources=[
            LootSource(
                id="crafted:blacksmithing", kind="crafted", name="Blacksmithing",
                profession="blacksmithing", items=[DARK_IRON_PLATE],
            ),
            LootSource(
                id="rep:thorium-brotherhood:honored", kind="rep", name="Thorium Brotherhood",
                items=[PLANS_DARK_IRON_PLATE],
            ),
            LootSource(id="world_drop:1-60", kind="world_drop", name="World Drop",
                       items=[COMMON_REAGENT]),
        ],
    )  # fmt: skip
    classic_crafted = {
        DARK_IRON_PLATE: ClassicDbCraftedRecipe(
            recipe_item_ids=[PLANS_DARK_IRON_PLATE], reagent_item_ids=[COMMON_REAGENT]
        )
    }
    gated = apply_crafted_opens_gate(document, classic_crafted)
    assert gated.sources == document.sources


def test_lionheart_helm_stays_open_world_drop_plans():
    """Lionheart Helm's own real shape: the Plans are a world drop
    (launch content) -- must NOT gain a gate."""
    document = LootFile(
        sources=[
            LootSource(
                id="crafted:blacksmithing", kind="crafted", name="Blacksmithing",
                profession="blacksmithing", items=[LIONHEART_HELM],
            ),
            LootSource(id="world_drop:55-60", kind="world_drop", name="World Drop",
                       items=[PLANS_LIONHEART_HELM, COMMON_REAGENT]),
        ],
    )  # fmt: skip
    classic_crafted = {
        LIONHEART_HELM: ClassicDbCraftedRecipe(
            recipe_item_ids=[PLANS_LIONHEART_HELM], reagent_item_ids=[COMMON_REAGENT]
        )
    }
    gated = apply_crafted_opens_gate(document, classic_crafted)
    assert gated.sources == document.sources


def test_an_alternative_ungated_recipe_source_keeps_the_item_open():
    """Two recipe items teach the same spell (a rare but real shape) --
    one is a raid drop, the other a world drop -- only ONE is needed, so
    the world-drop alternative keeps the item open."""
    plans_a, plans_b = 90001, 90002
    document = LootFile(
        sources=[
            LootSource(
                id="crafted:blacksmithing", kind="crafted", name="Blacksmithing",
                profession="blacksmithing", items=[SULFURON_HAMMER],
            ),
            LootSource(id="raid:molten-core", kind="raid", name="Molten Core", opens="later",
                       items=[plans_a]),
            LootSource(id="world_drop:55-60", kind="world_drop", name="World Drop",
                       items=[plans_b]),
        ],
    )  # fmt: skip
    classic_crafted = {
        SULFURON_HAMMER: ClassicDbCraftedRecipe(
            recipe_item_ids=[plans_a, plans_b], reagent_item_ids=[]
        )
    }
    gated = apply_crafted_opens_gate(document, classic_crafted)
    assert gated.sources == document.sources


def test_the_most_restrictive_of_two_different_reagent_gates_wins():
    """One reagent is Onyxia's Lair-gated ("raids-1", opens 9 December,
    the earliest of the two announced phases); another is a `"later"`
    raid with no announced date at all -- `"later"` is the more
    restrictive of the two known phases (`_CRAFTED_OPENS_ORDER`'s own
    doc) and must win."""
    reagent_early, reagent_late = 90101, 90102
    document = LootFile(
        sources=[
            LootSource(
                id="crafted:blacksmithing", kind="crafted", name="Blacksmithing",
                profession="blacksmithing", items=[SULFURON_HAMMER],
            ),
            LootSource(id="raid:onyxias-lair", kind="raid", name="Onyxia's Lair",
                       opens="raids-1", items=[reagent_early]),
            LootSource(id="raid:molten-core", kind="raid", name="Molten Core", opens="later",
                       items=[reagent_late]),
        ],
    )  # fmt: skip
    classic_crafted = {
        SULFURON_HAMMER: ClassicDbCraftedRecipe(
            recipe_item_ids=[], reagent_item_ids=[reagent_early, reagent_late]
        )
    }
    gated = apply_crafted_opens_gate(document, classic_crafted)
    gated_source = next(s for s in gated.sources if s.id == "crafted:blacksmithing:later")
    assert gated_source.opens == "later"


def test_a_forever_new_crafted_id_classic_crafted_does_not_cover_stays_ungated():
    """`classic_crafted` names no entry at all for this item (a
    Forever-new id, or a Classic id classic-db's own dump has no
    create-item spell for) -- left alone, never invented into a gate."""
    forever_new_item = 271234
    document = LootFile(
        sources=[
            LootSource(
                id="crafted:blacksmithing", kind="crafted", name="Blacksmithing",
                profession="blacksmithing", items=[forever_new_item],
            ),
        ],
    )
    gated = apply_crafted_opens_gate(document, classic_crafted={})
    assert gated.sources == document.sources


def test_item_effective_gate_reads_the_final_document_not_a_fixed_point_loop():
    """`item_effective_gate` (the helper `_crafted_item_gate` shares with
    itself, no recursion needed): an item present in two sources, one
    open, resolves to None; an item present in a single gated source
    resolves to that gate; an item absent entirely resolves to None
    (never invented)."""
    document = LootFile(
        sources=[
            LootSource(id="raid:molten-core", kind="raid", name="Molten Core", opens="later",
                       items=[1]),
            LootSource(id="world_drop:1-60", kind="world_drop", name="World Drop", items=[1, 2]),
            LootSource(id="raid:onyxias-lair", kind="raid", name="Onyxia's Lair",
                       opens="raids-1", items=[3]),
        ],
    )  # fmt: skip
    assert item_effective_gate(document, 1) is None  # also world-dropped: has an open way
    assert item_effective_gate(document, 2) is None  # only ever world-dropped
    assert item_effective_gate(document, 3) == "raids-1"
    assert item_effective_gate(document, 999) is None  # no source at all
