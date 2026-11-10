# AI-slop review round 5: character and planner (1440 and 390)

Not counted: handoffs (dashed Point order border, yellow dash, Client build footer, tier-gate gutter, paperdoll ilvl and Upgrades card), plated Activity and class-coloured lists (owner rule), the "Rating 65 of 100" row.

## Verdicts
- character-visitor-60: PASS
- character-owner-60-premium: PASS WITH FIXES
- planner: PASS WITH FIXES
- planner-you: PASS WITH FIXES
- planner-guide: PASS WITH FIXES

## Round 4 fixes confirmed
Owner banner no longer truncates at 390 ("Devout Skirt, Baron Rivendare"). The duplicate "Get the addon and companion" is gone from the top of the planner plate. Guide builds now sit under Point order at full width on all three planner pages. No glow, gradient, eyebrow, pill pile or emoji anywhere.

## Remaining
1. Tell 13, owner 1440 and 390: History plate still holds two collapsed rows ("Score 1,830 on 3 Nov", "Coach reviews") and nothing else. Make them plain rows under Best per encounter or open the first. (CharacterOwnerSections.svelte)
2. Tell 11 (layout), planner 1440: Guide builds moved left, so the rail now ends at ~800 px while the left column runs to ~1590, ~790 px of dead space right of Point order and Guide builds. planner-you is ~470 px, planner-guide ~390 px. Let Point order and Guide builds span full width below the Talents/rail row.
3. Tell 17/8 mild, planner-guide: Level 60 build plate holds a "Matches" chip and a button. Fold the chip into the heading line or drop it.
4. Tell 22 mild, planner 1440 and 390: "Gear for 60" / "Gear for 30 to 39" label beside "Best in slot list" restates the link. Keep the link only.
5. Minor, visitor and owner 390: Lucifron row wraps "+82 vs 27 Oct" onto a third line, sparkline floats beside the parse.

## Keep
Paperdoll and real icons, quality-coloured Upgrades with HPS, +128.9 HPS lead, parse-coloured Best per encounter, label/value Standing, bar rows in the score block, real trees, point-order strip, named guide builds with "points off" and "Missing X", Send this build to Felsnap.
