"""cmangos/classic-db's `item_template`/`spell_template` as the recipe
item and reagent items behind a crafted item's `SPELL_EFFECT_CREATE_ITEM`
spell -- the fact `loot.json`'s own `crafted:<profession>` buckets do not
carry at all (the engine fork's own `assets/database/db.json` states only
which PROFESSION crafts an item, never which Plans/Pattern/Schematic/
Formula/Recipe/Manual item teaches it or which reagents it consumes).

data-followups-3 lane, 2026-09-30, item 1: Sulfuron Hammer (17193) is
`crafted:blacksmithing` with no `opens` -- a fresh level-60 reads that as
launch content, but its recipe (Plans, 18592/227727) is a Molten Core
quest reward and one of its reagents (Sulfuron Ingot, 17203) is a Molten
Core boss drop, both gated `"later"`. `pipeline.loot.sources.
apply_crafted_opens_gate` is what actually computes and publishes the
gate, from the `ClassicDbCraftedRecipe` this module extracts.

Same two-stage shape as `pipeline.classicdb_items`/
`pipeline.classic_quest_levels`: `extract_records` (occasional, needs the
full pinned mysqldump -- `pipeline.audit.dumpdb.ClassicDbDump`, the SAME
dump `fetch-classic-sources` already downloads once) builds the table in
memory; `write_extract`/`load_extract` are the committed, offline-in-CI
`raw/classicdb/crafted-recipes.json` cache `pipeline.loot.sources` reads
at every `loot`/`loot-merge` run, same as `raw/classicdb/sources.json`
and `raw/classicdb/item_template.json` beside it.

The chain `extract_records` walks, per crafted item id (verified against
Sulfuron Hammer's own real ids while building this lane's report):

1. `dump.created_item_to_spells[item_id]` -- every crafting spell whose
   own `SPELL_EFFECT_CREATE_ITEM` effect names this item (Sulfuron
   Hammer 17193 -> spell 21161 "Sulfuron Hammer").
2. For each such spell, every OTHER spell whose own effect is
   `SPELL_EFFECT_LEARN_SPELL` (`ClassicDbDump.LEARN_SPELL_EFFECT`) and
   whose `EffectTriggerSpell` names it (spell 21161 <- spell 23007
   "Plans: Sulfuron Hammer", a recipe item's own on-use spell).
3. For each such "learn" spell, every item whose own `spellid_1..5`
   names it (`dump.item_template_spell_ids`, reversed) -- the recipe
   ITEM the player actually holds/uses (18592 "Plans: Sulfuron Hammer",
   and its Forever-new re-itemised twin 227727). Empty when no such item
   exists at all -- a TRAINER-taught recipe, which needs no item and so
   is never a gate on its own (`recipe_item_ids: []`,
   `apply_crafted_opens_gate`'s own doc for how that reads).
4. `dump.spell_reagents[spell_id]` (only the CREATE-ITEM spell from step
   1, never the "learn" spell) -- every reagent item id the recipe
   consumes (17193's own: Sulfuron Ingot 17203, Dark Iron Bar 11371,
   Arcanite Bar 12360, Essence of Fire 7078, Blood of the Mountain
   11382, Lava Core 17011, Fiery Core 17010).

A crafted item id classic-db's own dump names NO create-item spell for
at all (a Forever-new id, id >= `pipeline.loot.wowhead.
FOREVER_NEW_ID_THRESHOLD`, or a Classic-id one `pipeline.audit.
check_crafted`'s own "major" finding already flags) is simply absent
from the result -- never invented (tenet 8): `apply_crafted_opens_gate`
leaves an item this dict does not cover ungated, the same "unverifiable
stays unlabelled" rule this pipeline follows everywhere else.
"""

from __future__ import annotations

import json
import logging
from collections import defaultdict
from datetime import UTC, datetime
from pathlib import Path

from pydantic import BaseModel

logger = logging.getLogger(__name__)

FILE_NAME = "crafted-recipes.json"


class ClassicDbCraftedRecipe(BaseModel):
    """One crafted item's recipe chain, `pipeline.loot.sources.
    apply_crafted_opens_gate`'s own input.

    `recipe_item_ids`: every item (Plans/Pattern/Schematic/Formula/
    Recipe/Manual) whose own on-use spell teaches this item's crafting
    spell -- ANY ONE obtained is enough (an "OR": a recipe sometimes has
    more than one drop/vendor/quest source, each its own item row in
    rare cases, though almost always exactly one). Empty means
    TRAINER-taught: no item stands between a character and the recipe,
    so this side of the gate is always open.

    `reagent_item_ids`: every reagent the crafting spell consumes --
    ALL of them are needed (an "AND"), each checked against `loot.json`
    the same way any other item's own effective gate is.
    """

    recipe_item_ids: list[int] = []
    reagent_item_ids: list[int] = []


def extract_records(sql_text: str) -> dict[int, ClassicDbCraftedRecipe]:
    """Every item id classic-db's own dump resolves a create-item spell
    for (`ClassicDbDump.created_item_to_spells`, ~1,669 on the pinned
    commit) -> its `ClassicDbCraftedRecipe`. NOT narrowed to ids a
    particular build's `loot.json` currently buckets under `crafted:
    <profession>` -- this extract, like `item_template.json` beside it,
    is a fact about classic-db's OWN universe, read fresh by
    `apply_crafted_opens_gate` for whichever crafted item ids that
    build's own `crafted:<profession>` sources happen to name (a
    Forever-new id, or a Classic id classic-db itself has no recipe for,
    simply gets no entry here and stays ungated by this fact alone --
    `pipeline.audit.check_crafted`'s own "major"/"minor" findings already
    report the latter case)."""
    # Local import: `pipeline.audit.dumpdb` lives under the `pipeline.audit`
    # PACKAGE, whose own `__init__` eagerly imports every check module
    # (`check_crafted` among them) -- one of which imports `pipeline.loot`,
    # which imports THIS module at its own top level (`pipeline.loot.
    # sources`' own `apply_crafted_opens_gate`). A module-level import here
    # would be circular; a local one, same convention `pipeline.
    # classic_sources.fetch_and_write_classic_sources` already uses for
    # this module's sibling extract (`pipeline.classicdb_items`), is not.
    from pipeline.audit.dumpdb import ClassicDbDump

    dump = ClassicDbDump.from_text(sql_text)
    created_item_to_spells = dump.created_item_to_spells
    spell_effects = dump.spell_effects
    item_template_spell_ids = dump.item_template_spell_ids
    spell_reagents = dump.spell_reagents

    # spell id -> every OTHER spell id that LEARNS it (SPELL_EFFECT_
    # LEARN_SPELL's own EffectTriggerSpell naming it) -- built once,
    # rather than rescanning all ~22k spell_template rows per crafted
    # item.
    learn_spell_for: dict[int, list[int]] = defaultdict(list)
    for spell_id, effects in spell_effects.items():
        for effect in effects["effects"]:
            if effect["effect"] == ClassicDbDump.LEARN_SPELL_EFFECT and effect["trigger_spell"]:
                learn_spell_for[effect["trigger_spell"]].append(spell_id)

    # spell id -> every item whose own spellid_1..5 names it -- the
    # SAME reversal, for item_template this time.
    items_with_spell: dict[int, list[int]] = defaultdict(list)
    for item_id, spell_ids in item_template_spell_ids.items():
        for spell_id in spell_ids:
            items_with_spell[spell_id].append(item_id)

    out: dict[int, ClassicDbCraftedRecipe] = {}
    for item_id, craft_spells in sorted(created_item_to_spells.items()):
        recipe_item_ids: set[int] = set()
        reagent_item_ids: set[int] = set()
        for craft_spell_id in craft_spells:
            for learn_spell_id in learn_spell_for.get(craft_spell_id, []):
                recipe_item_ids.update(items_with_spell.get(learn_spell_id, []))
            for reagent_item_id, _count in spell_reagents.get(craft_spell_id, []):
                reagent_item_ids.add(reagent_item_id)
        out[item_id] = ClassicDbCraftedRecipe(
            recipe_item_ids=sorted(recipe_item_ids), reagent_item_ids=sorted(reagent_item_ids)
        )
    return out


def raw_path(build_dir: Path) -> Path:
    """`data/builds/<build>/raw/classicdb/crafted-recipes.json` --
    committed despite living inside the otherwise-gitignored `raw/` tree,
    the same exception `.gitignore` already carries for
    `raw/classicdb/sources.json`/`item_template.json`."""
    return build_dir / "raw" / "classicdb" / FILE_NAME


def write_extract(
    build_dir: Path, records: dict[int, ClassicDbCraftedRecipe], *, source_commit: str
) -> Path:
    path = raw_path(build_dir)
    path.parent.mkdir(parents=True, exist_ok=True)
    document = {
        "generated_at": datetime.now(UTC).isoformat(),
        "source": {
            "classic-db": {
                "repo": "cmangos/classic-db",
                "commit": source_commit,
                "license": "GPL-3.0",
                "note": "item_template.spellid_1-5 + spell_template's own SPELL_EFFECT_"
                "LEARN_SPELL/SPELL_EFFECT_CREATE_ITEM effects and Reagent1-8/"
                "ReagentCount1-8 columns -- see pipeline.classicdb_crafted's own doc.",
            },
        },
        "items": {
            str(item_id): record.model_dump() for item_id, record in sorted(records.items())
        },
    }
    path.write_text(json.dumps(document, indent=1) + "\n", encoding="utf-8")
    return path


def load_extract(build_dir: Path) -> dict[int, ClassicDbCraftedRecipe]:
    """The committed cache for one build, or `{}` (with a warning) when
    nothing has been extracted for it yet -- `loot`/`loot-merge` still
    run, every crafted item simply stays ungated by this fact (its
    profession bucket alone) until then, same graceful degradation every
    other optional classic-db cache already gets."""
    path = raw_path(build_dir)
    if not path.exists():
        logger.warning(
            "classicdb-crafted: no %s; run `python -m pipeline fetch-classic-sources` first "
            "-- every crafted item's recipe/reagent chain stays unresolved until then",
            path,
        )
        return {}
    document = json.loads(path.read_text(encoding="utf-8"))
    return {
        int(item_id): ClassicDbCraftedRecipe(**fields)
        for item_id, fields in document["items"].items()
    }
