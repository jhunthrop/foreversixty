# ux-designer review, sim pages, h-audit round 1 (2026-11-09)

Reviewed: design/reviews/2026-11-09-h-audit/first/{sim-landing-out, sim-landing-in, sim-result, sim-saved, sim-gear, sim-talents, sim-drops, sim-weights, sim-specs}-{1440,390}.png
Against: design/specs/2026-11-09-directions.md (round 5 corners, .framed, one lead panel), design/specs/2026-11-06-sim.md, docs/design-system.md, BRIEF owner rules.
Method: captures only. Fixture artefacts (blank spell icons, "Not yet" status from a missing API row, the saved-builds read error) are flagged "check" when they look like fixture state.

## Verdicts
| Page | Verdict |
|---|---|
| sim-landing-out | SHIP WITH FIXES |
| sim-landing-in | SHIP WITH FIXES |
| sim-result | SHIP WITH FIXES |
| sim-saved | SHIP WITH FIXES |
| sim-gear (Top Gear) | NOT SHIPPABLE |
| sim-talents | NOT SHIPPABLE |
| sim-drops | NOT SHIPPABLE |
| sim-weights | SHIP WITH FIXES |
| sim-specs | NOT SHIPPABLE |

Works: compact title band with tool tabs, plate panels, one framed lead panel per page, class-coloured names on plates, quality-coloured item names, DPS with its unit, no em dashes seen, result and saved pages dense and well structured.

## sim-landing-out
Question: how do I get my character in? Paste box and Run are above the fold at both widths. Framed lead panel is correct.
- polish: 390 "Sim a best-in-slot set": each class row is 87 px tall and spec links wrap leaving orphans ("Survival ›", "Destruction ›"); crest floats mid-row. One wrapping line beside the crest, row 56 px. SimLanding.svelte / SpecGrid.svelte.
- polish: the subline "Damage specs only. Nothing is uploaded until you save." and on other tabs "The simulator runs damage specs only. Healing and tanking specs are not simulated yet." repeat one fact; one wording (SimToolHead.svelte, lib/sim/copy.ts).

## sim-landing-in
Run sim is at 1440 y=252 and in the 390 run bar. Character head with crest, spec, ilvl correct.
- must fix (1440): left column ends at y=1330 (Gear, Talents) while the Rotation rail runs to y=1880; a 550 px void under Talents. Move "Your sims" under Talents or shorten the rail. SimView.svelte / SimLandingPage.svelte.
- must fix (1440+390): Rotation step 1 "Other cooldowns as they come up" has no icon and is indented differently from steps 2 to 12. Give it the ActionIcon box. SimRotationSection.svelte / RotationDisclosure.svelte.
- polish: "Before the pull" rows are 58 px with 20 px gaps; tighten to 44 px.
- polish: Gear panel shows enchant lines here, "Item level 61" on the result page; one subtitle rule in SimGearSection.svelte.
- polish: "No sim yet / Press Run sim..." two lines of right-aligned grey in the band; the instruction should be cream.

## sim-result
Answer (904.4 DPS, unit, error, Level 60 raid BiS comparison, Save this sim, Find upgrades) is in the band. Correct.
- must fix (390): 12,209 px tall, every section open. Collapse Rage, Buffs and debuffs, This run and Rotation behind Disclosure.svelte on phone; keep Damage by ability, DPS spread, Gear open. SimView.svelte, LazySection.svelte.
- must fix (1440+390): Rage row "Refunded on a miss, dodge or parry" has an empty icon gap where every sibling has an icon. Supply an icon. SimResourceSection.svelte.
- must fix (1440): Timeline "Off / On" toggle beside "Swings and ticks" has no label of what it toggles. Label it. SimTimeline.svelte / TimelineBrush.svelte.
- polish: Copy for the game / Copy JSON / Download JSON stack in two rows of uneven width; one row of three.
- polish: "Up all fight" buff chips are 7 loose rows; two columns on one plate.

## sim-saved
Same body as result; band "Fury Warrior", "Saved 4 Nov", Open in the sim, Copy link. Correct for a share.
- must fix: same three as sim-result (Rage icon gap, Timeline toggle label, 390 collapse at 11,821 px).
- must fix: title is the spec only; show the character name when the sim carries one.
- polish (390): keep DPS plus the two actions inside the first 420 px.

## sim-gear (Top Gear)
- blocker (1440+390): Candidates plate lists only Head and Main hand; 1,450 px of empty page under it while the right rail stacks four panels. Render all 16 slots as the two-column Gear list from sim-landing-in, each with an add-candidate control. CandidateRows.svelte / SlotGrid.svelte / TopGear.svelte.
- must fix: "Copy and modify" is a full-width cyan block per row; make it a secondary text button so Run is the one primary.
- must fix: strip "Nothing to try yet. ... Three stages: everything at 100 iterations, the survivors at 1,000, the finalists at 3,000." is bold on tan and mixes empty state with engine talk. Empty line only; stages in a disclosure (lib/sim/copy.ts combinationsNone).
- must fix: Find an item Add buttons (50 px cyan blocks) outweigh the item names; use a 28 px icon button. ItemSearch.svelte.
- must fix: "Your saved builds could not be read" in default state (check: fixture); needs Retry if real.
- polish: Whole sets and Talent builds into Disclosure.

## sim-talents
- blocker (1440+390): one 260 px plate with two checkboxes; no talent preview, no point split, no icons. Each option shows the three-tree mini grid used on sim-landing-in. TalentCandidates.svelte.
- must fix: 500 px empty tan under the plate at 1440.
- must fix: saved-builds error in default state (check: fixture).
- polish: "Add a build" cyan primary beside a disabled Run; make it secondary.

## sim-drops
- blocker (1440+390): 7,331 px (1440) and 7,751 px (390) of 11 px grey checkbox rows: 20 dungeons each with "Bosses n" and "Trash and chests", 6 professions, 36 reputation, 38 PvP ("Rank 10 (Alliance)", counts 0 to 2). No icons, no filter, nothing collapsed, no answer. Reference: Adventure Guide / Wowhead loot table. Collapse each category behind a Disclosure showing "n of m ticked", give dungeons the instance icon, hide zero-count rows, add a filter field. SourcePicker.svelte, DropResults.svelte, Droptimizer.svelte.
- must fix: Raids block has the checkbox "Show unreleased content" and, below the raid list, a button with the same label; keep one.
- must fix: "No source ticked yet" at 11 px below the whole list; move beside Run.
- must fix: same "Nothing to try yet ... Three stages" copy problem as Top Gear.
- polish: ticked-row state not shown in capture; verify selected tint at 6 px radius.

## sim-weights
Closest to standard. Hit rating 1.00 first, bars, per-point DPS and error with units, framed table.
- must fix (1440): every row has a 40 px empty left gutter where a stat icon would go. Add stat icons or drop the gutter. WeightsScaleTable.svelte.
- must fix (1440): "Stats to weigh" sits alone in the right rail, 220 px tall, 400 px shorter than the table; it is a setup control after the result. Make it a chip strip above the table. StatWeights.svelte.
- must fix (390): character strip wraps badly ("Entered by hand" alone, "Open in planner ›" right-aligned above a full-width Change source). One row for the two actions under the identity. ToolCharacterLine.svelte.
- polish: "3,120 engine iterations in your browser: a baseline pass plus a low and a high pass..." is engine talk; use "Runs in your browser."
- polish: Haste note in grey 12 px in the plate header; make it a table foot row.

## sim-specs
- blocker (1440+390): all 18 damage specs read "Not yet" including Fury Warrior, which sims on every other page (check: fixture lacks status; capture not reviewable if so). Each row also carries a "What it does" arrow explaining nothing visible. SpecGrid.svelte / SpecCard.svelte / UnsimulatedSpecCard.svelte, mergeSpecRows.
- blocker (1440+390): no crests anywhere; 26 text rows at 65 px. One-crest rule: circular ringed spec crest per row, status pill, one line of what is modeled. Healers and tanks rows get the crest too.
- must fix: "Guide >" chevron differs from the site's "›".
- must fix: grey "Not yet" pill 11 px, about 3.4:1; raise to 12 px, cream 70%.
- polish: two columns by class at 1440.

## Checklist
- [ ] No empty icon squares or gutters: result Rage row, landing-in Rotation step 1, weights table.
- [ ] No thin pages: talents, top gear, spec support.
- [ ] One crest language: spec support.
- [ ] Answer without scrolling at 390: pass landing, result, weights; fail drops, top gear.
- [ ] Phone collapse of secondary sections: result, saved fail.
- [ ] One primary per region: top gear, talents.
- [ ] No engine meta talk: top gear, talents, drops, weights.
- [x] No em dashes.
- [x] Number takes its unit: result, weights.
- [x] Class and quality colours on plates.
