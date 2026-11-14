VERDICT: SHIP

Reviewed on side-by-side.png, side-by-side-2.png, brand-me.png, tooltip.png against the site captures and the spec. The mock reads as the site: painted blue-teal band with the ringed crest, nav-bar strip with the seal and site icons, accent underline on the active tab, umber plates on a painted brown ground, Marcellus titles in sentence case, teal and blue buttons, one teal selected state. No BLOCKER or SHOULD findings.

1. NICE (side-by-side.png, Me and Builds): the brown ground only shows in the 8 px margin and gap, so it is thin next to the site where it dominates. Keep 8 px for density, but make sure ground.tga keeps the site's mottling visible at that width (forever.ground #67563a, no extra tint).
2. NICE (tooltip.png, right): the "Alt-click: open Gear" hint is gold. The spec reserves gold for the waiting dot, score and item level; use forever.accent or base.muted for hints, as bis-tooltip-1440.png does.
3. NICE (brand-me.png, Right now rows): the spec's chevron-right glyph for rows that open a detail is not visible in the mock; add it or drop it from ADDON_ICONS so the manifest matches.
4. NICE (brand-builds.png): the "Following" label is gold. If it is a state, use forever.accent or gain green to match the Pinned/Following treatment on the web.

Round 3 (d939533ef): VERDICT: SHIP holds. Re-checked side-by-side.png and side-by-side-2.png against the site captures. More earth shows through the 10 px margins and 6 px gaps, the plates at 0.84 still carry text cleanly, the fixed-width tabs and accent underline match the site nav, the Guild header dusk and cave glow echo the site's raid surfaces, and the accent Following label, ticks and chevrons close my round 1 NICE items 3 and 4. The colour key maps every region to a token. No BLOCKER or SHOULD. Remaining NICE: the tooltip hint line is still tooltipYellow; the site's bis tooltip keeps hints quiet, so consider forever.accent or base.muted.
