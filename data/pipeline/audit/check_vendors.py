"""Check D: every vendor item against cmangos/classic-db's own npc_vendor
rows, and the rep gate on a vendor slot against what the dump actually
carries for one.

What the pinned dump carries for a rep gate (checked once, here, rather
than repeated in every finding): `conditions` type 5
(`CONDITION_REPUTATION_RANK`) on an `npc_vendor`/`npc_vendor_template` row's
own `condition_id` (faction + minimum standing), and `item_template`'s own
`RequiredReputationFaction`/`RequiredReputationRank` columns (one faction +
rank per item, independent of which vendor sells it). It carries no
`ExtendedCost` column at all -- a vendor item's token/currency cost is not
in this dump, so this check cannot verify one and does not try.
"""

from __future__ import annotations

from pipeline.audit.context import AuditContext
from pipeline.audit.findings import CategoryResult
from pipeline.loot.wowhead import FOREVER_NEW_ID_THRESHOLD

CODE = "D"
LABEL = "Vendors and rep"
PRIMARY_SOURCE = (
    "raw/classicdb/sources.json vendor records (npc_vendor(+template) + conditions type 5 "
    "CONDITION_REPUTATION_RANK); item_template.RequiredReputationFaction/RequiredReputationRank "
    "via --classicdb-dump. The dump carries no ExtendedCost column -- vendor token/currency "
    "costs are not checked here."
)


def check(ctx: AuditContext) -> CategoryResult:
    result = CategoryResult(code=CODE, label=LABEL, primary_source=PRIMARY_SOURCE)
    classic_sources = ctx.classic_sources_by_item
    if not classic_sources:
        result.skipped = "raw/classicdb/sources.json has no entries -- vendor checks skipped"
        return result
    for source in ctx.loot.get("sources", []):
        if source.get("kind") != "vendor":
            continue
        npc_id = source.get("npc_id")
        for item_id in source.get("items") or []:
            result.checked += 1
            records = [
                r
                for r in classic_sources.get(item_id, [])
                if r.kind == "vendor" and r.npc_id == npc_id
            ]
            if not records:
                severity = "minor" if item_id >= FOREVER_NEW_ID_THRESHOLD else "major"
                result.add(
                    severity,
                    item_id,
                    f"item {item_id} is sold by npc {npc_id} in loot.json but classic-db's "
                    "npc_vendor(+template) has no such row"
                    + (" (Forever-new item id)" if severity == "minor" else ""),
                )
                continue
            conditions = [r.condition for r in records if r.condition is not None]
            if conditions:
                condition = conditions[0]
                if source.get("faction_id") != condition.faction_id or source.get(
                    "standing"
                ) != condition.standing:
                    result.add(
                        "blocker",
                        item_id,
                        f"item {item_id} on npc {npc_id} is rep-gated in classic-db "
                        f"(faction {condition.faction_id}, {condition.standing}) but loot.json's "
                        "own source states a different (or no) gate",
                        ours=(
                            f"faction_id={source.get('faction_id')} "
                            f"standing={source.get('standing')}"
                        ),
                        theirs=f"faction_id={condition.faction_id} standing={condition.standing}",
                    )
            if ctx.dump is not None:
                template = ctx.dump.item_template.get(item_id)
                if template and template["required_reputation_faction"]:
                    if source.get("faction_id") != template["required_reputation_faction"]:
                        result.add(
                            "minor",
                            item_id,
                            f"item {item_id}'s own item_template states "
                            f"RequiredReputationFaction={template['required_reputation_faction']} "
                            f"rank={template['required_reputation_rank']}, not reflected in "
                            "loot.json's source-level faction_id/standing",
                            theirs=f"faction={template['required_reputation_faction']} "
                            f"rank={template['required_reputation_rank']}",
                        )
    return result
