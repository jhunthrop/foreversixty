# UX review: character and planner, round 6 (round6-captures/, 1440 and 390)

Regression check after the round-5 layout change (planner: Talents beside the rail, Point order and Guide builds full width beneath; character: History folded into Best per encounter).

## Verdicts
- character-visitor-60: SHIP
- character-owner-60-premium: SHIP
- planner: SHIP
- planner-you: SHIP WITH FIXES
- planner-guide: SHIP

## Blockers: none. Must fix: planner-you 1440 void
- planner-you 1440: the rail (Level build, In the game with Send, Copy, Share, Show DPS, Import) runs to ~y1190; Talents ends ~y845. That leaves ~345 px of empty tan under the talent plate, then Point order. This is the r4 void again, mirrored. Fix: at 1440 move Share this build and Show DPS out of the rail (a row under Talents, or the Talents plate footer beside Reset), or drop "In the game" beneath the full-width plates, so the rail is not taller than Talents by more than ~60 px.

## Regression check
- 390, all five: no horizontal scroll; tree fits with a tab per tree; no empty icon squares; no em dashes; Point order and Guide builds stack cleanly full width.
- planner 1440: columns balance (rail ends ~y800, Talents ~y845). The r5 ~800 px void is gone.
- planner-guide 1440: balanced (rail ~y860, Talents ~y845).
- character-owner 1440: Gear left ~y1130 then Best per encounter to ~y1580; rail ends ~y1555. Balanced. History is now two closed disclosures inside Best per encounter ("Score 1,830 on 3 Nov", "Coach reviews"); nothing previewed, but they read as a footer, not a gap.
- character-visitor 1440: Best per encounter ends ~y1415; rail ends ~y1178 (~235 px void). Unchanged polish.

## Polish
- planner-you 390: the Level 30 to 39 build card with Load sits below Point order (~y1520); primary action two screens down. Move above Point order. planner-guide 390: In the game (Copy for the game) is at the very bottom, below Guide builds.
- planner-you/guide 390: "Show DPS" toggle sits ~8 px past plate content padding.
- owner 1440: upgrade nudge plate is 720 of 1248 px wide with an empty right half; span full width.
- owner/visitor 390: Lucifron row wraps "+82 vs 27 Oct" onto its own line.
- planner 1440 signed-out: tan "Warrior talents" title on the blue band reads low contrast.
- Handed off, not failed: dashed Point order border, "Client build" footer, PointOrder dash marker, item level/enchant.

## Checklist
- [x] Compact title band, crest
- [x] Paperdoll ornate, quality rings
- [x] Unit follows number
- [x] No em dashes
- [x] 390 no horizontal scroll, tree fits, tab per tree
- [x] Lists sit on plates
- [x] No empty icon squares
- [x] One job per page
- [x] Columns balance at 1440: planner, planner-guide, owner
- [ ] Columns balance at 1440: planner-you (345 px void), character-visitor (235 px)
- [ ] One plate style (dashed Point order, handed off)
