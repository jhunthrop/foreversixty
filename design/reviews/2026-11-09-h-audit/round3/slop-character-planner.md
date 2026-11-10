# AI-slop review round 3: character and planner (1440 and 390)

Handoffs (Client build footer, yellow dash, tier-gate gutter, paperdoll item level and Upgrades card) and the owner-requested "Rating 65 of 100" row are not counted.

## Verdicts
- character-visitor-60: PASS WITH FIXES
- character-owner-60-premium: PASS WITH FIXES
- planner: PASS WITH FIXES
- planner-you: PASS WITH FIXES
- planner-guide: PASS WITH FIXES

## Round 2 fixes confirmed
Standing block now label/value; "Synced 1 day ago"; "Show DPS"; "51 points left"; "Felsnap is 34. Then Beast Mastery guide order."; planner blank no longer shows an empty Point order plate or dim Share/Copy; Score block shows no Tier row (Epic sits beside the score).

## Remaining
1. Tell 8, all five pages: every block is the same dark bordered plate with serif heading. Planner stacks Talents, Point order, Share, Import, Level build, In the game, Guide builds; owner character 390 stacks eight. Files: components/planner/Planner.svelte, components/character/CharacterOwnerSections.svelte. Fix: Share and Import in one plate with a hairline; Builds seen and Activity as plain lists without plate chrome.
2. Tell 11, dead space at 1440: planner rail ends ~880 vs left ~1095; planner-you ~1005 vs ~1835; planner-guide ~990 vs ~1795; character visitor ~1200 vs ~1415. Fix: Share and Import side by side as a full-width pair under Point order, or move Guide builds below Point order.
3. Tell 13 (mild), planner-you 1440 and 390: Point order plate still has a dashed border, unlike planner-guide. Fix: solid like the others.
4. Tell 22, owner 390: Upgrade banner truncates to "Upgrade: Devout Skirt, Stratholme: B...", cutting the source. Fix: drop "Upgrade:" or truncate the source so "Devout Skirt, Baron Rivendare" fits (CharacterOwnerSections.svelte).
5. Tell 22/17, both character pages, both widths: band meta "Supporter" is plain grey text with no state. Fix: crest/badge or drop.

## Keep
Paperdoll and real icons, quality-coloured Upgrades with HPS, +128.9 HPS lead, parse-coloured Best per encounter, real trees, point-order strip, named guide builds with "points off", Send this build to Felsnap, no glow/gradient/eyebrow.
