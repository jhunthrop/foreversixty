"""Check H: `data/curated/specs.json` against the engine's own stat
vocabulary, and an apl file for every spec it names.

`pipeline.specs.load_specs` already enforces "reference_stat and every
weight_stats id is a member of the engine's own vendored `Stat` enum"
(`pipeline.simdb.statmap.STAT_IDS`, generated from `pipeline.simproto`'s
vendored bindings, not retyped) every time it runs -- that IS this
category's "consistent with the engine's stat names" brief, so this check
calls it rather than re-implementing the same comparison a second time.
"""

from __future__ import annotations

from pipeline.audit.context import AuditContext
from pipeline.audit.findings import CategoryResult

CODE = "H"
LABEL = "Curated specs vs engine"
PRIMARY_SOURCE = (
    "pipeline.specs.load_specs (reference_stat/weight_stats vs pipeline.simdb.statmap.STAT_IDS, "
    "the engine's own vendored Stat enum); data/curated/apl/<spec>.json existence"
)


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
    return result
