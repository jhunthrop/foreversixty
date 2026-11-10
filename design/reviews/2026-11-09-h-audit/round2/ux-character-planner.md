# UX review: character and planner, round 2 (captures in round2-captures/, 1440 and 390)

Reference: in-game character pane (paperdoll), client talent frame. Mocks: b2-character, b2-planner.
Handoff items (Paperdoll Upgrades card, item level/enchant, spec median per boss, PointOrder dash marker, "Client build" footer, tier-gate gutter) are noted, not failed.

## Verdicts
- character-visitor-60: SHIP
- character-owner-60-premium: SHIP WITH FIXES
- planner: SHIP
- planner-you: SHIP
- planner-guide: SHIP

## Round 1 findings, status
- 1 Nudge row: PARTLY FIXED. Four links are now one ("Upgrade: Devout Skirt, Stratholme: Baron Rivendare >"), but it is still a bare yellow line on the band gradient, no plate, no icon; at 390 it wraps to two lines. STILL MUST FIX.
- 2 History/Coach thin plates: PARTLY. Lead sentences gone; now two collapsed rows ("Score and parses 7 fights", "Coach reviews") on one plate with a tiny caret and nothing previewed. POLISH (disclosure, content one tap away).
- 3 Right-column void at 1440: FIXED. Standing, Builds seen and Activity now balance the doll column on owner; visitor void is ~170 px.
- 4 Owner 390 order: NOT FIXED. Forever Score sits after History. POLISH.
- 5, 6: FIXED/OK. Spec in class colour and guild in yellow give one step of emphasis.
- 7, 8: unchanged (fixture, no avatar_url/render_url). Not judged.
- 9 Point order empty state: 60 empty squares gone, replaced by a one-line plate "No points spent yet." POLISH: still thin.
- 10 Default Level 9: NOT FIXED. Signed-out warrior still reads "Human . Level 9", "0 of 51 points". POLISH.
- 11 Tan title contrast: NOT FIXED. "Warrior talents" in class tan on blue band. POLISH.
- 12 Disabled toggle: reworded "Show DPS on the card", still disabled with no reason at 0 points. POLISH.
- 13 "Talents loaded from a character" loose line: FIXED (gone).
- 14 Dash marker: handoff (ui/PointOrder.svelte).
- 15 Share under Point order: FIXED. Right column now ends early, 800 px void at 1440 on planner-guide. POLISH.
- 16 "Matches": FIXED, now a chip like "Loaded".

## Per page

### character-visitor-60 (SHIP)
1440: compact band (crest, display-type name, one facts row, "Sim this character", "+128.9 HPS from upgrades" unit right). Ornate doll with quality rings and greyed empty slots, Upgrades card over the centre (handoff h-doll). Forever Score with bars, Builds seen, Activity with crest avatars, Best per encounter with parse colour and HPS unit. Columns end within 170 px.
390: doll swaps to a two-column icon slot list; Upgrades one plate; no horizontal scroll.

### character-owner-60-premium (SHIP WITH FIXES)
- MUST FIX. Upgrade nudge, 1440 y 192 and 390 y 233, unplated. Source components/character/CharacterOwnerSections.svelte (band in CharacterIdentity.svelte). Fix: one plate with the Devout Skirt icon in its quality ring, 44 px target, single line with truncation at 390.
- POLISH. 390: move Forever Score and Standing directly under Gear.
- POLISH. History plate: preview the last parse chip or sparkline in the closed rows (CharacterHistory.svelte).
- POLISH. Standing line "#3 of 14 healers in Cx66 Olympus, 60 raid" orphans "raid" at 1440.

### planner (SHIP)
Band compact with "0 of 51 points", class strip with crests, three trees with rank pills, Level 60 build, In the game, Guide builds with spec icons. 390: tab per tree with counts, tree fits, plates stack. Polish only: items 9, 10, 11, 12.

### planner-you (SHIP)
"25 of 25 points", "Sim this build >", Levels 30 to 39 build, Send this build to Felsnap, Point order with real icons (dash marker is handoff). 390 ok. POLISH: Point order plate has a dashed border that no other plate has; use the solid plate. Right column ends at y~1000 against a 1830 left column.

### planner-guide (SHIP)
"51 of 51 points", "Matches" chip, "Loaded" chip on the Survival row, Point order full, Share and Import in the left column. 390 ok. POLISH: 800 px right-column void at 1440; make the right column sticky (top 16 px).

## Checklist
- [x] Title band compact, crest, no big header
- [x] Paperdoll ornate, quality rings, client order
- [x] Unit follows number
- [x] No em dashes in copy
- [x] 390: no horizontal scroll, tree fits
- [ ] Lists sit on plates (owner nudge fails)
- [x] No empty icon squares
- [x] One job per page
