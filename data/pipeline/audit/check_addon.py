"""Check G: `addon-data.json` and the addon's `Data.lua` against the same
drift check CI already runs (`python -m pipeline addon-data --check`) --
`pipeline.addondata.check_addon_data`/`pipeline.addonlua.lua_has_drifted`
regenerate both from the published `bis/*.json`/`talents/*.json`/curated
inputs and diff. A pass here is what "addon-data.json and Data.lua equal
the published bis/*.json picks" (this category's own brief) means in
practice: both files ARE `build_addon_data`'s output, byte- (Lua) or value-
(JSON) identical. Rotation-band completeness is the same function's own
`AddonRotationError` (it does not tolerate a missing curated/apl entry, per
that module's doc) plus a direct count against `pipeline.addonrotation.
LEVEL_BANDS`.
"""

from __future__ import annotations

from pipeline.audit.context import AuditContext
from pipeline.audit.findings import CategoryResult

CODE = "G"
LABEL = "Addon parity"
PRIMARY_SOURCE = (
    "pipeline.addondata.check_addon_data + pipeline.addonlua.lua_has_drifted "
    "(build_addon_data regenerated from bis/*.json + talents/*.json + curated, diffed against "
    "the committed addon-data.json/Data.lua)"
)


def check(ctx: AuditContext) -> CategoryResult:
    result = CategoryResult(code=CODE, label=LABEL, primary_source=PRIMARY_SOURCE)
    from pipeline.addonbis import AddonBisError
    from pipeline.addondata import AddonDataError, check_addon_data
    from pipeline.addonlua import LUA_PATH, lua_has_drifted
    from pipeline.addonrotation import LEVEL_BANDS, AddonRotationError

    result.checked += 1
    try:
        if check_addon_data(ctx.build, root=ctx.root, curated_dir=ctx.curated_dir):
            result.add(
                "blocker",
                ctx.build,
                "addon-data.json has drifted from build_addon_data's regenerated output",
            )
    except (AddonDataError, AddonBisError, AddonRotationError) as error:
        result.add(
            "blocker", ctx.build, f"addon-data.json could not even be regenerated: {error}"
        )
        return result

    result.checked += 1
    try:
        if lua_has_drifted(ctx.build, root=ctx.root, curated_dir=ctx.curated_dir):
            result.add(
                "blocker",
                ctx.build,
                f"{LUA_PATH} has drifted from build_addon_data's regenerated Lua",
            )
    except (AddonDataError, AddonBisError, AddonRotationError) as error:
        result.add("blocker", ctx.build, f"Data.lua could not even be regenerated: {error}")

    addon_data = ctx.addon_data
    if addon_data is not None:
        rotations = addon_data.get("rotations", {})
        for spec in ctx.bis_specs:
            bands = rotations.get(spec)
            result.checked += 1
            if not bands:
                result.add(
                    "major", spec, f"addon-data.json has no rotation bands at all for {spec}"
                )
            elif len(bands) != len(LEVEL_BANDS):
                result.add(
                    "minor",
                    spec,
                    f"{spec} has {len(bands)} rotation bands, LEVEL_BANDS names {len(LEVEL_BANDS)}",
                    ours=str(len(bands)),
                    theirs=str(len(LEVEL_BANDS)),
                )
    return result
