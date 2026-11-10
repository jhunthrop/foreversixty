# AI-slop review round 6: character and planner (1440 and 390)

Not counted: handoffs (dashed Point order border, yellow dash, Client build footer, tier-gate gutter, paperdoll ilvl and Upgrades card), plates on class-coloured lists, the "Rating 65 of 100" row.

## Verdicts
- character-visitor-60: PASS
- character-owner-60-premium: PASS WITH FIXES
- planner: PASS WITH FIXES
- planner-you: PASS WITH FIXES
- planner-guide: PASS

## Round 5 fixes confirmed
History is no longer its own plate; its two rows sit under Best per encounter. Point order and Guide builds span full width on all three planner pages. The Matches chip and the "Gear for N" label are gone. No glow, gradient, eyebrow, pill pile, emoji or icon-per-bullet grid.

## Remaining
1. Tell 13 mild, owner 1440 and 390: "Score 1,830 on 3 Nov  7 fights" and "Coach reviews" are collapsed rows that show nothing until opened. Open the first, or drop "Coach reviews" when it holds nothing. (CharacterOwnerSections.svelte)
2. Tell 11 (dead space), planner-you 1440: the rail (Level 30 to 39 build, In the game with Send, Copy, Share, Show DPS, Import) runs to ~1190 px while Talents end at ~840, leaving a ~370 px empty block under the trees. Put Share and Import side by side, or move Import below the full-width row.
3. Tell 19/22 mild, planner-you: "Past it: Endurance Training 4/5, Bestial Swiftness, Intimidation and Bestial Discipline 2/2" is grey explainer prose under "21 of 25 points". Keep as a plain list or cut.
4. Tell 13 mild, planner 1440: "In the game" plate has an empty ~40 px row under the heading before Import from addon. Drop the row when Send and Share are absent.
5. Minor, owner 390 (and visitor): Lucifron row wraps "+82 vs 27 Oct" to a third line and the sparkline floats beside the parse.

## Keep
Paperdoll with real icons, quality-coloured Upgrades with HPS, +128.9 HPS lead, parse-coloured Best per encounter, label/value Standing, bar rows in the Score block, real trees, point-order strip, named guide builds with "points off" and "Missing X", "Loaded" state on the matched build, Send this build to Felsnap.
