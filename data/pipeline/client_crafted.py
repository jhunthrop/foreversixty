"""The client's own recipe chain for crafted items, in the same
`ClassicDbCraftedRecipe` shape `pipeline.classicdb_crafted` extracts from
cmangos' classic-db.

classic-db knows no Forever-new craft, so `apply_crafted_opens_gate` left
every one of them ungated. Feedback 2026-11-09: the level-60 Protection
Paladin list picked Greenhammer (279261) and Evergreen Shield (279262),
Forever-new Blacksmithing epics whose recipe (client spells 1306535 and
1306536) consumes Malleable Essence of Nature (274030), which the client
itself describes as "Acquired from disenchanting items from Hyjal Summit or
Barrow Deeps", raids that open on 9 December. Every Forever-new craft that
needs it is closed until then.

The chain, all from the client's DB2 tables (wago.tools CSVs under `raw/`):

1. `SpellEffect` rows with `Effect` 24 (create item) name the crafted item
   in `EffectItemType` and the craft spell in `SpellID`.
2. `SpellReagents` lists the craft spell's reagents (`Reagent_0..7`).
3. A recipe item (Plans, Pattern, Formula, ...) is an item whose
   `ItemEffect` row (through `ItemXItemEffect`) has `TriggerType` 6 (learn)
   and names the craft spell; its `ItemSparse.RequiredSkillRank` is the
   profession skill the recipe needs.

`loot` (the full run, with raw tables) writes the extract to
`raw/client/crafted-recipes.json`; `loot` and `loot-merge` both read it, so
the two commands gate crafts from the same facts.
"""

from __future__ import annotations

import json
import logging
from collections import defaultdict
from pathlib import Path

from pipeline.classicdb_crafted import ClassicDbCraftedRecipe
from pipeline.csvio import read_csv

logger = logging.getLogger(__name__)

#: Under `raw/` so the build manifest leaves it out, beside
#: `raw/client/loot-classification.json`.
CLIENT_CRAFTED = Path("raw/client/crafted-recipes.json")

#: `SpellEffect.Effect` for SPELL_EFFECT_CREATE_ITEM.
CREATE_ITEM_EFFECT = "24"
#: `ItemEffect.TriggerType` for an on-use "learn" (a recipe item).
LEARN_TRIGGER = "6"
REAGENT_COLUMNS = tuple(f"Reagent_{i}" for i in range(8))


def _positive(value: str | None) -> int:
    number = int(value or 0)
    return number if number > 0 else 0


def _created_items(spell_effects: list[dict[str, str]]) -> dict[int, set[int]]:
    """crafted item id -> every spell that creates it."""
    out: dict[int, set[int]] = defaultdict(set)
    for row in spell_effects:
        item_id = _positive(row.get("EffectItemType"))
        if row.get("Effect") == CREATE_ITEM_EFFECT and item_id:
            out[item_id].add(int(row["SpellID"]))
    return out


def _reagents(spell_reagents: list[dict[str, str]]) -> dict[int, set[int]]:
    """craft spell id -> its reagent item ids."""
    out: dict[int, set[int]] = defaultdict(set)
    for row in spell_reagents:
        out[int(row["SpellID"])].update(
            item_id for column in REAGENT_COLUMNS if (item_id := _positive(row.get(column)))
        )
    return out


def _recipe_items(
    item_effects: list[dict[str, str]], item_x_effects: list[dict[str, str]]
) -> dict[int, set[int]]:
    """craft spell id -> every recipe item whose on-use learns it."""
    learned_by_effect = {
        int(row["ID"]): int(row["SpellID"])
        for row in item_effects
        if row.get("TriggerType") == LEARN_TRIGGER and _positive(row.get("SpellID"))
    }
    out: dict[int, set[int]] = defaultdict(set)
    for row in item_x_effects:
        spell_id = learned_by_effect.get(int(row["ItemEffectID"]))
        if spell_id is not None:
            out[spell_id].add(int(row["ItemID"]))
    return out


def build_recipes(
    spell_effects: list[dict[str, str]],
    spell_reagents: list[dict[str, str]],
    item_effects: list[dict[str, str]],
    item_x_effects: list[dict[str, str]],
    item_sparse: list[dict[str, str]],
) -> dict[int, ClassicDbCraftedRecipe]:
    """Every client-crafted item with at least one reagent -> its recipe."""
    reagents = _reagents(spell_reagents)
    recipe_items = _recipe_items(item_effects, item_x_effects)
    skill_rank = {int(row["ID"]): _positive(row.get("RequiredSkillRank")) for row in item_sparse}
    out: dict[int, ClassicDbCraftedRecipe] = {}
    for item_id, spells in sorted(_created_items(spell_effects).items()):
        reagent_ids = set().union(*(reagents.get(spell, set()) for spell in spells))
        if not reagent_ids:
            continue  # a scripted conjure or quest spell, not a profession craft
        recipe_ids = set().union(*(recipe_items.get(spell, set()) for spell in spells))
        skills = [skill_rank[i] for i in recipe_ids if skill_rank.get(i)]
        out[item_id] = ClassicDbCraftedRecipe(
            recipe_item_ids=sorted(recipe_ids),
            reagent_item_ids=sorted(reagent_ids),
            required_skill=min(skills, default=None),
        )
    return out


def recipes_from_raw(raw: Path) -> dict[int, ClassicDbCraftedRecipe]:
    return build_recipes(
        read_csv(raw / "SpellEffect.csv"),
        read_csv(raw / "SpellReagents.csv"),
        read_csv(raw / "ItemEffect.csv"),
        read_csv(raw / "ItemXItemEffect.csv"),
        read_csv(raw / "ItemSparse.csv"),
    )


def write_client_crafted(build_dir: Path, recipes: dict[int, ClassicDbCraftedRecipe]) -> Path:
    path = build_dir / CLIENT_CRAFTED
    path.parent.mkdir(parents=True, exist_ok=True)
    document = {str(item_id): recipe.model_dump() for item_id, recipe in sorted(recipes.items())}
    path.write_text(json.dumps(document, indent=1) + "\n", encoding="utf-8")
    return path


def load_client_crafted(build_dir: Path) -> dict[int, ClassicDbCraftedRecipe]:
    """The extract, or `{}` with a warning before the first full `loot` run
    wrote it (every Forever-new craft then stays ungated, as before)."""
    path = build_dir / CLIENT_CRAFTED
    if not path.exists():
        logger.warning(
            "client-crafted: no %s; run `python -m pipeline loot --build %s` with the raw "
            "client tables to write it -- Forever-new crafts stay ungated until then",
            path,
            build_dir.name,
        )
        return {}
    document = json.loads(path.read_text(encoding="utf-8"))
    return {int(item_id): ClassicDbCraftedRecipe(**fields) for item_id, fields in document.items()}


def merge_recipes(
    classic: dict[int, ClassicDbCraftedRecipe], client: dict[int, ClassicDbCraftedRecipe]
) -> dict[int, ClassicDbCraftedRecipe]:
    """classic-db's recipes, plus the client's for every item classic-db
    does not know. classic-db stays first for a Classic craft: its recipe
    items carry the drop and vendor sources loot.json gates on."""
    return {**client, **classic}
