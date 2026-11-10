# ux-designer review, sim pages, h-audit round 2 (2026-11-10)

Reviewed: round2-captures/{sim-landing-out, sim-landing-in, sim-result, sim-saved, sim-gear, sim-talents, sim-drops, sim-weights, sim-specs}-{1440,390}.png
Against: round1/ux-sim.md, design/specs/2026-11-09-directions.md, design/specs/2026-11-06-sim.md, BRIEF owner rules. Handoff items (timeline Off/On label = f-content; SimGearSection stat strip = h-doll) noted, not failed.
Captures only; fixture artefacts (saved builds "did not load", talent "current build 0/20/0", two-slot candidates) are marked "check".

## Verdicts
| Page | R1 | R2 |
|---|---|---|
| sim-landing-out | SWF | SHIP |
| sim-landing-in | SWF | SHIP WITH FIXES |
| sim-result | SWF | SHIP WITH FIXES |
| sim-saved | SWF | SHIP WITH FIXES |
| sim-gear | NOT SHIPPABLE | SHIP WITH FIXES |
| sim-talents | NOT SHIPPABLE | SHIP WITH FIXES |
| sim-drops | NOT SHIPPABLE | NOT SHIPPABLE |
| sim-weights | SWF | SHIP |
| sim-specs | NOT SHIPPABLE | SHIP WITH FIXES |

## sim-landing-out: SHIP
Fixed: one subline ("Damage specs only."), paste box and Run above the fold at both widths, class names on a plate, crests ringed.
- polish (390): Hunter and Warlock rows still wrap one spec onto a second line (Survival, Destruction), rows 87-120 px against 87 px for the rest. SimLanding.svelte / SpecGrid.svelte.

## sim-landing-in: SHIP WITH FIXES
Fixed: 1440 columns balance (left ends y=1330, rail y=1375), "Your sims" moved under the rail, Before-the-pull rows tightened.
- must fix (1440+390): Rotation step 1 "Other cooldowns as they come up" still has no icon and a different indent than steps 2 to 12. Give it the ActionIcon box. SimRotationSection.svelte / RotationDisclosure.svelte.
- polish: "No sim yet" is 12 px grey at the band's right edge; make it cream and pair with Run sim.
- handoff: Gear panel enchant lines versus "Item level n" (h-doll, SimGearSection).

## sim-result: SHIP WITH FIXES
Answer (904.4 DPS, error, vs BiS, Save this sim, Find upgrades) is in the band at both widths.
- must fix (390): page is 12,189 px, every section open (was 12,209). Collapse Rage, Buffs and debuffs, This run, Rotation behind Disclosure.svelte on phone; keep Damage by ability, DPS spread, Gear open. SimView.svelte / LazySection.svelte.
- must fix (1440+390): Rage row "Refunded on a miss, dodge or parry" still has an empty icon gap (every sibling has an icon). SimResourceSection.svelte.
- polish: Copy for the game / Copy JSON / Download JSON still two uneven rows; one row of three.
- polish: "Up all fight" chips still 7 loose rows at 390.
- handoff: Timeline Off / On unlabeled (SimTimeline.svelte, f-content).

## sim-saved: SHIP WITH FIXES
Share band correct (Open in the sim, Copy link, Saved 4 Nov).
- must fix (390): 11,801 px, same collapse as sim-result.
- must fix (1440+390): Rage icon gap, same as sim-result.
- must fix: title is "Fury Warrior" only; show the character name when the sim carries one (check: fixture may not carry it).

## sim-gear (Top Gear): SHIP WITH FIXES
Fixed: right rail no longer leaves a void (page 1,697 px), Add is a quiet text action, Copy and modify is a text link, "Three stages" engine copy gone.
- must fix (1440+390): Candidates lists only Head and Main hand (check: fixture gear). With full gear it must render all 16 slots; confirm with a full-gear capture. CandidateRows.svelte / SlotGrid.svelte.
- must fix: "Add a build" and "Add set" are cyan primaries beside a disabled Run; make secondary. TopGear.svelte.
- polish: Whole sets and Talent builds into Disclosure; "Nothing ticked yet." sits detached under the band, put it beside Run.
- polish: "Saved builds did not load." needs Retry if real (check: fixture).

## sim-talents: SHIP WITH FIXES
Fixed: every option shows the three-tree mini grid with point split (Arms 17, Fury 32, Protection 2); plate fills the page; 390 stacks cleanly.
- must fix (1440): each option row leaves 560 px empty right of the grid (grid ends x=760). Add a point-split chip and "Open in planner" per row, or a diff against the current build. TalentCandidates.svelte.
- must fix: "Add a build" cyan primary beside disabled Run; secondary.
- polish: "Your current build 0/20/0" shows Fury 20 only (check: fixture).

## sim-drops: NOT SHIPPABLE
Better: 3 columns at 1440 (2,578 px, was 7,331), bold names, Bosses disclosure, one "Show unreleased content". Still short of standard.
- blocker (390): 5,570 px of 44 px checkbox rows, a second row "Bosses n" under every dungeon (88 px per dungeon); no category collapse. Disclosure per category with "n of m ticked". DropResults.svelte / SourcePicker.svelte.
- blocker (1440+390): no instance, profession or faction icon anywhere; text-only rows. SourcePicker.svelte.
- must fix (1440+390): top chips tick Raids, Dungeons, Crafted, Reputation, PvP while "No source ticked yet" shows above; contradictory. Droptimizer.svelte.
- must fix: zero-count rows still shown ("0 items" Rank 15 and Rank 5, "Friendly 0"); hide.
- must fix: Raids prose ("Nothing here has opened yet. Onyxia's Lair: ... No date yet: ...") is an 11 px grey run-on; one line plus chips.
- polish: 36 reputation tickboxes as loose text rows; add a filter field.

## sim-weights: SHIP
Fixed: stat icon gutter gone, Stats to weigh is a strip above the table, 390 table is two lines per row with bar, identity strip tidy.
- polish: "3,120 iterations" engine talk above the strip; drop or "Runs in your browser."
- polish: Haste note still a grey 12 px line in the plate header; make it a table foot row.

## sim-specs: SHIP WITH FIXES
Fixed: all 18 damage specs supported with Raid BiS DPS and rank, crests on every row, healers and tanks listed.
- must fix (1440+390): "What it does" disclosure on every row (18 identical triggers). One expand or drop. SpecGrid.svelte / SpecCard.svelte.
- must fix (1440): ~1,000 px dead gap between spec name and rank/DPS at 57 px rows. Add a DPS bar or two columns. SpecGrid.svelte.
- must fix (390): rows 85 px because the name wraps and the disclosure takes a second line. One-line name, 56 px.
- must fix: same class crest for all specs of a class at 22 px; spec crest per row at 28 px.
- must fix: healer and tank "Guide >" chevron; site uses "›".
- polish: sort by DPS rank, not by class.

## Checklist
- [ ] No empty icon squares or gutters: result Rage row, landing-in Rotation step 1, drops (no icons at all).
- [x] No thin pages: talents, specs, top gear (pending full-gear check).
- [ ] One crest language: specs uses class crest per spec, small.
- [x] Answer without scrolling at 390: landing, result, weights, specs. Drops: answer after 5,570 px of lists.
- [ ] Phone collapse of secondary sections: result, saved, drops.
- [ ] One primary per region: top gear, talents (Add a build cyan).
- [x] No engine meta talk: top gear, talents; remaining on weights (polish).
- [x] No em dashes.
- [x] Number takes its unit.
- [x] Class and quality colours on plates.
