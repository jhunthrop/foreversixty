# ai-slop review, round 12 (1440 and 390)

Haste-per-1% caption excluded by owner rule. Healer/tank rows without numbers accepted (needs-data: no healer/tank ladder). No visual tell 1, 2, 5, 8, 12 or 15 on any page.

## sim-landing-out: PASS
1440 and 390 clean. Spec links no longer wrap or collide. Nit only: 390 class rows are tall (about 2,000 px for nine classes); two class columns would halve it.

## sim-talents: PASS WITH FIXES
- Tell 21: "Pick builds to compare." under the tab band explains the obvious. Delete.
- Tell 13 (mild): "Your saved builds / No saved builds" over an indented "Add a build" is an empty state with a heading. Show only "Add a build" until a build exists.
- Nit, 1440: "Your current build" row has an empty left column under the checkbox; the Gains/Drops list runs far past the tree.
- 390 trees are now plate width and legible.

## sim-drops: PASS WITH FIXES
- Tell 21: "Tick a source." explains the obvious. Delete.
- Fixed: Raids is a row ("0 of 7 open"). Crafted rows have profession icons, item counts and ilvl.
- Nit: the "Show" lead-in before the checkboxes can go.

## sim-specs: PASS WITH FIXES
- 390: Healers and tanks is a two-column grid whose names wrap to two lines ("Restoration / Druid", "Protection / Paladin") with the grey "Healer"/"Tank" word floating beside it; "Healer" touches the plate edge. One row per spec at 390, or group under "Healers" and "Tanks" headings and drop the per-row grey word at both widths.
- Nit: "#n" repeats the sort order next to the bar.

## guides-spec (Survival Hunter): FAIL
Fixed since r11: rail title repeating "Stat priority" gone; "Mail at 40" said once; 390 band is "Survival Hunter / Jump to" only.
Still present:
1. Empty icon squares (owner rule, never): every Enchants row (7), every Consumables row (left list and rail), every Races rail entry (9 empty bordered squares), plus Strider Kick, Rapid Fire, Aspect of the Beast, Lightning Reflexes, Resourcefulness. Fix: real icons for enchants and consumables; race crests or drop the square on the Races rail.
2. Tell 21: Rotation is six paragraphs of 5 to 9 lines that restate the Opener/Priority/Cooldowns list beside them (Hunter's Mark, Raptor Strike, Mongoose Bite, Lacerate, Strider Kick, Immolation Trap, Rapid Fire all twice). Keep only what the list cannot say (pet dismissed, Falcon vs Beast, Expose Prey needs a mark, threat, trash, movement); cut to about a third.
3. Tell 21: Stat priority prose still explains the obvious ("Strength is attack power directly", "Haste is rare on gear but strong when you find it") and keeps an "Attack speed: Not on any item" row. Drop those.
4. Repetition: Talents prose re-lists Key talents that are in the rail and Overview Strengths; Levels 20 to 60 gives the rank-swap paragraph twice (26 and 46) in near-identical wording; Races prose restates the rail.
5. Dead space: Races ~500 px empty under the prose; Stat priority ~200 px; Professions and Leveling set prose ~740 px wide with the right third of the plate empty. Put the slider/tree in the rail or run the prose full width; fold Races into one race / racial / effect table.
6. Tell 6, weak: grey labels "Jump to", "Racials", "Alliance", "Horde". Drop "Racials" (section is Races).
7. Tell 17, mild: four pill kinds across ten Sources rows. Allowed by owner rule; group under plain headings instead.
8. 390: the "Survival Hunter" sticky bar repeats the title band right above it. Show "Jump to" alone.
Authentic: Rank-tagged spell rows with timings, "20 s before the pull", item source lines ("1% from Warchief Rend Blackhand"), paperdoll with tooltip, hit cap 9% / 28% / 16%, level slider, Overview Role/Brings/What the raid adds.

## Summary
PASS: sim-landing-out. PASS WITH FIXES: sim-talents, sim-drops, sim-specs. FAIL: guides-spec.
