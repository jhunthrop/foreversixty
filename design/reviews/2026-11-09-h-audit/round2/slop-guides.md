# AI-slop review: guides, round 2 (2026-11-10)

Captures: round2-captures/guides, guides-you, guides-class, guides-spec at 1440 and 390.
Tell 23 = definitional blurb ("X is the Y: a, b and c") explaining a class, spec or feature to a veteran.

## First impression
Reads as a player tool. Index is plain plates with spec icons; class page now states changes, not definitions; spec page is dense and authored. Remaining giveaways: the spec lede, blank icon squares, and tooltip text pasted into racials.

## /guides (signed out): PASS WITH FIXES
- Round 1 item 13 (blank spec icons) still present: Beast Mastery, Marksmanship, Warrior/Paladin Protection, Retribution, Priest Holy, Fire, Affliction, Demonology (1440 and 390). Fixture artefact if the build lacks them; if production lacks them, fall back to the spec crest. ClassSpecPicker.astro.
- No tell 23, no blurbs. Keep it that way.

## /guides signed in (level 34 hunter): PASS WITH FIXES
- 17 "Your spec" pill on Beast Mastery (1440 and 390) beside the yellow left rule. Drop the pill, keep the rule. GuideSpecRow.astro (wrap at 390 is f-selected, handed off).
- 22 "Level 34 · Dwarf Beast Mastery Hunter" (both widths): use "34 Dwarf Beast Mastery". GuideAnswer.astro.
- 22 "Your guide: Beast Mastery Hunter ›": prefix is the qualifier. lib/guides/copy.ts signedInCallout (f-content, handed off).
- 13 Same blank icons, including the marked row.

## /guides/hunter (class): PASS WITH FIXES (was FAIL)
Tell 23 is gone: "What changed from 1.12" now leads with the change ("Aimed Shot: trained, 2 second cast, one 6 second cooldown shared with Multi-Shot", "Mail from 40."), Heart of the Lion and the Aspect line are trimmed. Races prose cut to two sentences.
- 23 residual, Races rail (1440 under list, 390 below grid): "Human hunters are new in Forever. Skyborne hunt on both sides, as the High Order and the Windshapers." Restates the list above. Delete; mark Human "new" on its row. content/guides/hunter/index.md, GuideRaceList.
- 13 Blank Beast Mastery and Marksmanship icons on spec plates; plates are name + "Read the guide" only, three equal plates (1440, 390). Show point split or top stat on the plate. pages/guides/[class]/index.astro.
- 21 "Pets are reworked." bullet still two sentences plus a clause; trim to the numbers. hunter/index.md.
- 8 Bullets, Sources and Races all same boxed plate; make Sources a plain list. Source pills (Blizzard/Community/Datamined) on every row: keep, rule allows a Sources footer.
- Races icons are empty coloured squares (1440 and 390).

## /guides/hunter/survival (spec): PASS WITH FIXES
Remaining items are mostly f-content's (handoffs.md), noted not failed on:
- 23 Lede "Survival is a melee spec in Forever: two weapons..." (1440 under tab row, 390 first paragraph). content/guides/hunter/survival.md. f-content owns it; fix: "Two weapons, Raptor Strike on cooldown, Mongoose Bite when it lights up, no pet in a raid."
- 22 "DPS · Updated 7 Nov" in the band (both widths). lib/guides/copy.ts headerMeta (f-content).
- 22 Racials column (1440, Races): tooltip text pasted ("Grants 10% critical strike chance for 15 seconds on a 3 minute cooldown", "Increases hit chance with spells and attacks by 1%") beside prose saying the same. Trim to "10% crit, 15 s / 3 min". race-table.ts (f-content).
- 17 "Recommended" pills on Night Elf and Orc beside coloured names. Keep one marker. GuideRacesLane.
- 13 Race icons, enchant and consumable rows have no icons (empty gutter at left of every row, 1440). f-content lane.
- 8 Dead plate under the Enchants prose (about 500 px empty, 1440). f-content.
- 21 Race prose: "Humans can be Hunters for the first time in Forever" and "Skyborne can be Hunters on either faction" are definitional restatements of the rail. Cut. survival.md.
- Em dashes: none seen.

## Authentic, keep
Talent trees with ranks, rotation with conditions, normalised weights with hit cap breakdown, "1 point of agility = 0.63 DPS at level 60", class-coloured plates, compact title band, uneven plate heights, spell tooltips in prose, Sources with provenance.

## Verdicts
- /guides: PASS WITH FIXES
- /guides-you: PASS WITH FIXES
- /guides/hunter: PASS WITH FIXES (tell 23 mostly cleared; Races rail sentence remains and should go)
- /guides/hunter/survival: PASS WITH FIXES (lede tell 23 is f-content's, noted not failed on)
