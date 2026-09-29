"""Check E: every crafted item (loot.json's `crafted:<profession>` buckets)
against cmangos/classic-db's own recipe data -- a spell whose
`SPELL_EFFECT_CREATE_ITEM` effect (`spell_template.EffectItemType`) names
the item.

The pinned SQL dump carries no `skill_line_ability` table at all -- that is
client-side `SkillLineAbility.dbc` data, not part of a cmangos world
database -- so a recipe's own required skill level cannot be read from this
source and is not reported as a fact; only whether a recipe spell exists at
all is checked here.
"""

from __future__ import annotations

from pipeline.audit.context import AuditContext
from pipeline.audit.findings import CategoryResult
from pipeline.loot.wowhead import FOREVER_NEW_ID_THRESHOLD

CODE = "E"
LABEL = "Crafted"
PRIMARY_SOURCE = (
    "--classicdb-dump item_template.spellid_1-5 + spell_template's own SPELL_EFFECT_CREATE_ITEM "
    "(EffectItemType) -- the dump has no skill_line_ability table, so a recipe's required skill "
    "level is not available from this source and is not reported"
)


def check(ctx: AuditContext) -> CategoryResult:
    result = CategoryResult(code=CODE, label=LABEL, primary_source=PRIMARY_SOURCE)
    if ctx.dump is None:
        result.skipped = "no --classicdb-dump given -- crafted-item recipe verification skipped"
        return result
    no_recipe = 0
    forever_new_no_recipe = 0
    for source in ctx.loot.get("sources", []):
        if source.get("kind") != "crafted":
            continue
        profession = source.get("profession", source.get("name", "?"))
        for item_id in source.get("items") or []:
            result.checked += 1
            spells = ctx.dump.created_item_to_spells.get(item_id, [])
            if spells:
                continue
            if item_id >= FOREVER_NEW_ID_THRESHOLD:
                forever_new_no_recipe += 1
                continue
            no_recipe += 1
            result.add(
                "major",
                item_id,
                f"item {item_id} is in the {profession} crafted bucket but no classic-db "
                "spell's SPELL_EFFECT_CREATE_ITEM effect names it",
            )
    if no_recipe:
        result.add(
            "minor",
            "*",
            f"{no_recipe} Classic-id crafted item(s) have no resolvable classic-db recipe spell",
        )
    if forever_new_no_recipe:
        result.add(
            "minor",
            "*",
            f"{forever_new_no_recipe} Forever-new crafted item(s) (id >= "
            f"{FOREVER_NEW_ID_THRESHOLD}) have no classic-db recipe, as expected -- not "
            "counted against the Classic-id total above",
        )
    return result
