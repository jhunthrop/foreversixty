# F5 Sims mock, player verdict (2026-10-31)

Verdict: HOLD (one fixture fix, then SHIP; the layout is sound).

## Would I use it between pulls?
Yes. Newest run on top, state words match the site ("Running · takes about 4 minutes", "Failed: the engine stopped"), the bar says "40 of 64 combinations", and Run again only shows on failures, as on the web. Kind names match web KIND_TITLES (Top Gear, Talent compare, Droptimizer, Stat weights; "Sim" for a plain run). DPS is mono and has its unit. The result shows 412.3 DPS, "± 1.8 · 10,000 iterations · engine 1f542def3" in the sim-tool convention. Icons are real item icons with epic purple and rare blue borders. Lapsed state is quiet and honest.

## Blockers
1. Fixture lie: "+41 DPS FROM VIS'KAG" on an Elemental Shaman, in a mock labelled Phase 1. Vis'kag the Bloodletter is a one-handed sword from Blackwing Lair. Shamans cannot wield swords in Forever (axes, maces, fists, daggers, staves only), and BWL is not Phase 1. A caster weapon with no spell power giving +41 DPS is not believable. Swap in a plausible Top Gear gain (a Molten Core or dungeon caster item). Same for the saved row "Elemental +41 DPS from Vis'kag" (F5.1, F5.2, F5.8).

## Should fix
- Result grid: 17 bare icons, no slot names or item names visible. I cannot tell which icon is the Vis'kag swap. Tapping opens the tooltip, but for Top Gear the changed item should be marked (spec item 6 does not say).
- Failed rows: "Failed: the engine stopped" tells me nothing actionable. The text is the site's, so it is fine for parity.
- Row title "Tuesday" with meta "Elemental · Sim" looks like a date. Make the fixture a user-style name.
- XL.5 clips "Started 3 min ago" at the right edge (state/age should stack, per spec). XL.6 header and nav overlap lower content; check on build.
- Grimtusk 88 DPS row: confirm the class and spec shown are legal for the race.
- F5.1 iOS board: tab bar overlaps the last row; there is no bottom padding.

## Re-check (fix round)
Boards seen: f5-sims-ios, f5-result-states-ios, f5-rerun-ios, xl-ios (XL.5/XL.6). Verdict: SHIP once one number is fixed.

Fixed: Vis'kag is gone. Titles read like a player's own ("MC full buffs", "Council pull", "Fury 40 to 49"). The gold corner and the "Gold corner: the slot this result changes" legend sit on the back slot (slot 4, correct order: head, neck, shoulder, back). The tab bar no longer covers the footer. XL.5 stacks state and age, and Run again grows. XL.6 holds the 6-up grid. Rerun states R1 to R6 read right.

Spritecaster Cape check (data/builds/1.60.1.70425/items.json and bis/shaman-elemental.md): item 11623, blue, item level 52, requires level 47, cloak, drops from High Interrogator Gerstahn in Blackrock Depths. A dungeon cloak a level-60 caster can have at launch, so the pick is legal and the Elemental Shaman can use it.

Remaining must-fix: "+19 DPS from Spritecaster Cape" is about ten times too big. The repo's own sim rates the cape's 18.8 spell power at 1.91 DPS at best, and at 60 it ranks below the BiS cloak (-0.42 DPS in the 50 table, -3.69 in the 60 table). It reads like the spell-power points were typed as DPS. Use "+2 DPS from Spritecaster Cape" (F5.1, F5.2, F5.2b, F5.3b, F5.8) so a shaman who knows the item is not misled.

## Re-check 2 (headline number, Android)
Verdict: SHIP.
"+2 DPS from Spritecaster Cape" now reads right on f5-sims-ios and f5-sims-and (F5.1, F5.2, F5.2b, F5.3b) and on f5-states-and (F5.8, F5.8b, F5.9, F5.9b). It matches the repo's 1.91 DPS for the cape. Android boards: gold corner and legend on the back slot, pinned footer, row states, offline strip with Run again disabled, spinner row and "Your sims did not load." with Try again all read fine. No new defects.
