# UX review: character and planner, round 1 (captures in first/, 1440 and 390)

Spec: design/specs/2026-11-09-directions.md (paperdoll, round 5), 2026-11-06-character.md, 2026-11-08-planner.md. Mocks: b2-character, b2-planner.
Reference: in-game character pane (paperdoll) and the client talent frame.

## Verdicts
- character-visitor-60: SHIP WITH FIXES
- character-owner-60-premium: SHIP WITH FIXES
- planner (empty warrior): SHIP WITH FIXES
- planner-you: SHIP
- planner-guide: SHIP

## What matches the mocks
Character: compact title band with crest, name in display type, one-row facts, HPS figure with unit right; ornate paperdoll with item icons in quality rings, client order, empty glyph slots, greyed two-hander, Upgrades tooltip standing over the centre; Forever Score and Builds seen in the right column; fight table with parse colour; Activity with crest avatars. Mobile swaps the doll for a two-column slot list and keeps Upgrades as one plate.
Planner: title band with class crest and "N of 51 points", class strip with crests, three trees with ranks and gold paths, point order with real icons and rank pips, Level N build / In the game / Guide builds / Share plates, the "Loaded" chip, mobile tab per tree.

## Findings

### character-owner-60-premium
1. MUST FIX. Nudge row under the title band (1440 y~190; 390 four stacked lines). Four yellow bare links ("Upgrade: Devout Skirt, Stratholme: Baron Rivendare >" etc.) float on the band gradient with no plate, no icon, no separation. At 390 they read as a ragged yellow list and push the doll 160 px down. Spec: lists sit on a plate. Fix: one plate (or a single horizontal scroll strip of chips with icon + label, 44 px targets) and give the first (upgrade) its item icon with quality colour. Source: components/character/CharacterOwnerSections.svelte (and CharacterIdentity.svelte for the band).
2. MUST FIX. History and Coach history (1440 y~1800-2050; 390 y~1640-1800) are two near-empty plates, each a heading plus one line ("Score over time, then parse per boss", "Every review the coach wrote for this character") with a tiny caret. The line explains the obvious and reads as a thin page. Fix: render the score sparkline and the last parse chips inside the closed History plate; Coach history shows the latest review title and date; drop the lead sentences. Copy lives in lib/character/page-view.ts lines 57-62. Source: CharacterHistory.svelte.
3. MUST FIX. Right column dead zone at 1440 (below Builds seen, ~1100 px tall on the visitor view, ~1300 on the owner view). The mock has the same void but with a shorter page; here History and Coach sit full-width under it. Fix: move Activity (or History) into the right column at 1440 so the two columns balance. Source: CharacterOwnerSections.svelte layout grid.
4. POLISH. Owner 390: the Score and Builds plates sit after History/Coach, so the character's headline numbers are 2,000 px down. Move Forever Score directly under Gear at 390.

### character-visitor-60
5. POLISH. Same right-column void at 1440 below Builds seen (y 920-1780).
6. POLISH. Header facts at 1440 use five equal-weight grey runs plus "Supporter" with no separator; the mock groups "Level 60 Night Elf . Holy Priest . <Guild>". Add the middle-dot grouping or weight the spec and guild in one step brighter. Source: CharacterIdentity.svelte.
7. POLISH. Medallion shows the LX seal; the spec allows this only with no portrait. Confirm the fixture has no avatar_url; if it has, the portrait is not wired.
8. POLISH. Owner and visitor views show no Battle.net render in the doll centre; spec says signed-in owner sees it. Fixture likely lacks render_url; verify with a character that has one.

### planner (empty warrior)
9. MUST FIX. Point order with nothing spent (1440 y 970-1270; 390 y 1450-1810) is 60 empty dashed squares. Brief: never an empty icon square. Fix: until a point is spent, collapse to one line plate "Spend a point and the order fills in here" with the Load the level 60 build button; show the grid only after the first point. Copy to planner-copy; source components/planner/Planner.svelte and lib/planner/copy.ts ("No points spent yet.").
10. POLISH. Default race/level reads "Human . Level 9", "0 of 51 points": a first-visit warrior sees a level 9 build with 51 points still to take on the right. Default the band to Level 60 when signed out, as the mock does.
11. POLISH. "Warrior talents" title is class tan on the blue band; contrast is about 4.5:1. Use white display with the class crest carrying colour, or raise the tan 10 percent.
12. POLISH. Toggle "Include a simmed DPS on the card" is shown disabled with no reason at 0 points. Dim the row and say why in the title attribute, or hide until points are spent.

### planner-you
13. POLISH. "Talents loaded from a character." (1440 y 268, 390 y 255) is a loose grey line on the band fade with nothing around it. Put it in the Talents panel header as a quiet right-aligned note. Copy lib/planner/copy.ts line 60.
14. POLISH. Marker line in Point order "Felsnap is level 34. Then the Beast Mastery guide's order." begins with a yellow dash glyph. No em dash in copy, but the glyph reads as one; use a small chevron or a rule. Source lib/planner/copy.ts line 95 and the row in Planner.svelte.

### planner-guide
15. POLISH. Void under Point order at 1440 (y 1320-1720, left column). Move Share this build under Point order in the left column so the columns balance.
16. POLISH. "Matches" in the Level 60 build plate is a bare green word; make it a chip like "Loaded".

## Checklist
- [x] Title band compact, crest, no big header (all five)
- [x] Paperdoll ornate, quality rings, client order (character)
- [x] Unit follows number (HPS, points)
- [x] No em dashes in copy
- [x] 390 has no horizontal scroll, tree fits
- [ ] Lists sit on plates (owner nudges fail)
- [ ] No empty icon squares (planner empty state fails)
- [ ] No thin plates (History, Coach history fail)
- [x] One job per page: planner has no gear panel, only a link to BiS
