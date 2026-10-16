"""Check H: `data/curated/specs.json` against the engine's own stat
vocabulary, an apl file for every spec it names, and every published
leveling-BiS pick against the engine's own `SimDatabase` item universe.

`pipeline.specs.load_specs` already enforces "reference_stat and every
weight_stats id is a member of the engine's own vendored `Stat` enum"
(`pipeline.simdb.statmap.STAT_IDS`, generated from `pipeline.simproto`'s
vendored bindings, not retyped) every time it runs -- that IS this
category's "consistent with the engine's stat names" brief, so this check
calls it rather than re-implementing the same comparison a second time.

simdb-supplement lane, 2026-09-30: `sim/internal/simdb.Attach`'s own
`UnequipUnknown` silently strips any item id `simitems.json` does not carry
from every tournament and verify character before it sims a `bis/*.json`
pick -- a set containing such an item is simmed WITHOUT it, so that pick's
`verified`/`score`/`set_dps` (and every trinket/effect tournament run
against it) is sim-checked evidence for the WRONG set. `pipeline/simdb/
items.py`'s client-row filter kept only ids present in both raw
`ItemSparse.csv` and `Item.csv`, so every classic-db-only id (the client's
`ItemSparse` never had a row for it at all) used to fail this silently;
this is the blocker-level check that closes it for good -- both a pick's
primary choice and its `alternatives` (tenet 7: a pick shows its
alternatives, which must be exactly as sim-checked as the pick itself).
"""

from __future__ import annotations

from pipeline.audit.context import AuditContext
from pipeline.audit.findings import CategoryResult

CODE = "H"
LABEL = "Curated specs vs engine"
PRIMARY_SOURCE = (
    "pipeline.specs.load_specs (reference_stat/weight_stats vs pipeline.simdb.statmap.STAT_IDS, "
    "the engine's own vendored Stat enum); data/curated/apl/<spec>.json existence; "
    "bis/<spec>.json picks + alternatives vs simitems.json (the engine's own SimDatabase "
    "item universe, pipeline.simdb.write_sim_items)"
)


def _picked_ids(band: dict) -> set[tuple[int, str]]:
    """Every `(item_id, item_name)` a band's own picks and alternatives
    name -- a pick's primary choice and every one of its `alternatives`
    (tenet 7) are equally published, so both are checked. A slot with no
    real pick at all (`empty_reason`: no source cleared the band's DPS
    floor, a two-hander already fills off_hand, an effect this build does
    not model yet, ...) simply has no `item_id` and names nothing here."""
    picked: set[tuple[int, str]] = set()
    for slot in band.get("slots", []):
        if "item_id" in slot:
            picked.add((slot["item_id"], slot.get("item_name", "")))
        for alt in slot.get("alternatives", []):
            if "item_id" in alt:
                picked.add((alt["item_id"], alt.get("item_name", "")))
    return picked


def check(ctx: AuditContext) -> CategoryResult:
    result = CategoryResult(code=CODE, label=LABEL, primary_source=PRIMARY_SOURCE)
    from pipeline.specs import SpecError, load_specs

    result.checked += 1
    try:
        specs = load_specs(ctx.curated_dir)
    except SpecError as error:
        result.add("blocker", "*", f"data/curated/specs.json failed its own validation: {error}")
        return result

    for spec in specs:
        result.checked += 1
        apl_path = ctx.apl_dir / f"{spec.spec}.json"
        if not apl_path.exists():
            result.add(
                "major", spec.spec, f"no apl file for spec {spec.spec} at {apl_path}"
            )

    simitems = ctx.simitems
    if simitems is None:
        result.skipped = (
            f"{ctx.build}: no simitems.json -- run `python -m pipeline simdb` for this build "
            "first; published picks were not checked against the engine's SimDatabase"
        )
    else:
        for bis_spec, doc in ctx.bis_by_spec.items():
            seen: set[int] = set()
            for band in doc.get("bands", []):
                for item_id, item_name in _picked_ids(band):
                    if item_id in seen:
                        continue
                    seen.add(item_id)
                    result.checked += 1
                    if item_id not in simitems:
                        result.add(
                            "blocker",
                            bis_spec,
                            f"published pick {item_id} ({item_name}) is not in simitems.json -- "
                            "the engine's own SimDatabase does not carry it, so "
                            "simdb.Attach's UnequipUnknown strips it from every tournament and "
                            "verify character; every sim-checked number for a set containing it "
                            "is wrong",
                        )
    return result
