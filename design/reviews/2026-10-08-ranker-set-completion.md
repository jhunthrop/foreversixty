# Ranker set completion, 2026-10-08

Lane set-completion. Branch `set-completion`. Closes the gap named in design/reviews/2026-10-08-phase1-set-bonuses.md section 4: the BiS ranker could not find a set bonus. The engine models every Phase 1 bonus; the ranker now picks pieces for them when the sim says the bonus pays.

## 1. Why a bonus could not enter a pick

How a band is built (`sim/cmd/leveling-bis/main.go`, per band and faction):

1. `pick()` takes, per slot, the highest `score()` candidate. `score()` sums one item's own stats times the band's stat weights (score.go). A set bonus is not on any item, so it is invisible to it, the same way a proc is.
2. Trinkets are re-ranked by `rankTrinketSlot` and slots with an implemented proc by `rankSlotWithEffects` (rank.go). Both swap one slot at a time and only among that slot's own candidates.
3. `trySetCompletion` (sets.go, before this lane) built its candidate list with `bestSetPieces`, which read `bySlot[slot][0]` for each slot: only the top-scored piece of each slot, and only when that piece belongs to an implemented set. A set was tried only when two or more slots had a top-scored piece of the same set.
4. `verifyBand` and `applySwaps` (verify.go) test each slot's runner-up alone against the pick and promote it at +1% beyond the sim error. A runner-up completes a bonus only when the other pieces are already worn.

Why a 2-, 3- or 5-piece bonus cannot enter a pick:

- **Step 3 needs the pieces to win their slots on stats first.** A set piece wins a slot on stats only when it also happens to be the best loose piece there. The piece that is worth wearing for a bonus is, by construction, one that loses on stats (otherwise nobody needs the bonus to justify it). A 2-piece bonus needs two such wins in different slots at once, a 3-piece three, a 5-piece five. The Phase 1 review records that it never fired, and the 28-spec band-60 run of the old binary used as the baseline here (no set-completion adoption note in 112 passes; the only two `adopted` lines are faction trinket reconciliations).
- **Step 4 can only add the last piece.** A runner-up swap changes one slot, so it completes a bonus only when all but one of the pieces are already worn. The first and second piece of a 3-piece bonus, or the first four of a 5-piece, are each measured alone, where they carry only their stat cost.
- **Step 4 also hides the first piece.** Each runner-up is measured against the current set. With one piece worn the bonus is not on, so the piece measures as its stat loss and is rejected one slot at a time.

The Phase 1 review proposed trying the best piece of each set per slot (about 43 extra runs per band and faction). That is what this lane does, with a bounded ladder over thresholds and a cheaper screen.

## 2. Design

`sim/cmd/leveling-bis/sets.go` (the old `bestSetPieces`, `setAlreadyFullyEquipped` and the independent-trial `trySetCompletion` are replaced, not kept beside it).

1. **Candidates.** `bestSetPiecePerSlot`: for each implemented set with a piece in at least two distinct slots of the band's pool, the best-scored piece of that set in every slot (trinkets excluded, a dual-wielder never offered a two-hander, one physical ring never counted as two pieces: the same three defences as before).
2. **Trial.** For each bonus threshold above the pieces already worn (from `sets.json`), whose piece count the pool can reach: wear the cheapest missing pieces (smallest score loss against the current pick, slot order breaking ties) up to the threshold, keep the per-slot picks elsewhere. Thresholds climb in order, so a weak 2-piece bonus does not hide a strong 3-piece one; each adoption becomes the baseline for the next threshold and the next set (sets run in id order).
3. **Verification harness.** The same one as the per-slot picks: `plainRequest`, `verifySeed` (7), `verifyIterations` (300), the band's talents and race, the engine runner (the healer engine for healers). Adopt when the trial beats the working picks by `swapMargin` (1%) and by more than the two runs' combined standard error (the rule verify.go's `Significant` and trinkets.go already use).
4. **Cost bound.** A screening run at `trinketRankIterations` (100) comes first; only a trial that beats its own screening baseline by half the margin earns the 300-iteration run. Same seed in both.
5. **Order.** Set completion now runs after `verifyBand` and `applySwaps`, against the swapped set. Before this the trial was compared with a set that the later swap pass improves by several percent (the first version of this lane lost 0.9% on shaman-elemental that way). When anything is adopted the completed picks go through `verifyBand`/`applySwaps` again (an adopted slot keeps the pick it replaced as its runner-up, so the sim can still undo it) and the completed band is published only if its verified set DPS is above the per-slot band's. Otherwise the per-slot band stands and the log says "reverted". This makes a DPS loss from set completion impossible by construction.
6. **Published.** `slots[].set_bonus: {set, pieces, bonus}` on each adopted slot only (the slots whose piece changed), `pieces` the threshold the trial was built for and `bonus` that tier's text from `sets.json`. Absent from every other row. The note is cleared from any slot whose set bonus a later swap broke (fewer pieces worn than the note says). Adopted pieces carry the winning trial's `sim_dps` like every other sim-decided pick.

Extra runs and wall time (28-spec band 60 run, raid and bare, both factions, 112 band-passes, standalone, 18 cores): an average of 89 extra runs per band, faction and preset for the 20 DPS specs (range 44 to 163 across specs; most are 100-iteration screens, the rest 300-iteration baselines and confirmations), 7.0 s per pass and 28 s per DPS spec at band 60 (563 s across the 20 DPS specs for four passes each). Whole ranker run, 28 specs, band 60: 1836 s. The old binary's run was measured at 948 s while competing with the first draft of this one for the same cores, so there is no clean standalone figure for it. At the nightly's five bands the set pass is about 140 s per DPS spec.

## 3. Tests

`sim/cmd/leveling-bis/sets_test.go`:

- two mediocre pieces of a set with a strong 2-piece bonus beat two better loose pieces and are adopted, with the published note, runner-up and run count asserted (`TestTrySetCompletionAdoptsAStrongTwoPieceBonus`);
- a pool where the bonus does not pay changes nothing and stops at the screen after two runs;
- a gain inside the sim error is not adopted;
- the ladder climbs past a losing 2-piece trial to a winning 3-piece one;
- a fully worn set and an empty pool cost no runs;
- `completeSets` keeps a completed band that verifies higher and reverts one that verifies lower;
- the `set_bonus` JSON shape (present, and absent from an ordinary row) and `loadSetCatalog`.

Web: `panel-view.test.ts` (wording, absent case), `GearRow.test.ts` (rendered note, title, glyph, absent case), the `priest-holy` fixture carries one adopted slot, and `tests/e2e/bis-healer.spec.ts` checks the note's shape. The e2e asserts shape, not count: a published `priest-holy.json` wins over the fixture (as that spec's header says), so the count depends on the nightly.

## 4. Web

`GearRow.astro`: one muted line under the source line, in the style of the evidence line (11px, muted, a stroked 12px glyph, no new colour), reading "Worn for the <Set> <N>-piece bonus" with the bonus text as its hover title. The glyph is a stroked two-link mark added to `GlyphSheet.astro` (`bis-set-bonus`) and drawn like the source-kind glyphs. The item icon on the row is unchanged. Guides untouched.

## 5. Results: band 60, all 28 specs, raid and bare, both factions

Before is the old binary (main at 17ace40b), after is this branch, same engine, data and seeds, scratch `-out`; the figure is the published `set_dps` (HPS for healers, tank score for tanks). 112 rows, 57 changed, 55 identical. Specs with no change at all: druid-feral-bear, shaman-restoration, warlock-affliction, warlock-demonology, warlock-destruction, warrior-protection. Pick changes without a set note are runner-up swaps that fell out of the second verification pass (a set piece in the neighbouring slot changed what the sim prefers).

**No DPS spec or tank lost anything: 0 of 92 DPS and tank rows are lower.** One healer row is lower: druid-restoration bare horde, published HPS 320.3 to 297.7 (-7.1%). The ranker's healer objective is the mana-guarded score (score_heal.go: HPS times the share of the fight the mana lasts, squared), not the published unguarded HPS. On that row the guarded score rose from 186.5 to 193.8 (+3.9%) because the set lasts 241 s of mana against 229 s, while raw HPS fell. It is the existing objective doing what it says, the same trade the per-slot picks already make; it is flagged here because the instruction was no loss, and it is a decision for the owner whether a healer set may trade headline HPS for mana. The other 19 healer rows are equal or higher.

| Spec | Preset | Faction | Before | After | Change | Set pieces adopted for a bonus | Other pick changes |
|---|---|---|---|---|---|---|---|
| druid-balance | bare | alliance | 220.0 | 225.7 | +2.60% | Feralheart Raiment 6pc (wrist: Dryad's Wrist Bindings -> Feralheart Wraps) | - |
| druid-balance | raid | alliance | 504.6 | 516.5 | +2.36% | Lieutenant Commander's Wildhide 2pc (chest: Feralheart Vest -> Knight-Captain's Dragonhide Armor, feet: Feralheart Galoshes -> Knight-Lieutenant's Dragonhide Boots) | - |
| druid-feral | bare | alliance | 296.7 | 303.1 | +2.15% | Feralheart Raiment 4pc (shoulder: Darkspear Pauldrons -> Feralheart Epaulets, wrist: Forest Stalker's Bracers -> Feralheart Bands, feet: Drudge Boots -> Feralheart Walkers) | - |
| druid-feral | bare | horde | 299.0 | 305.2 | +2.07% | Feralheart Raiment 4pc (shoulder: Darkspear Pauldrons -> Feralheart Epaulets, wrist: Forest Stalker's Bracers -> Feralheart Bands, feet: Drudge Boots -> Feralheart Walkers) | - |
| druid-feral | raid | alliance | 610.5 | 628.3 | +2.91% | Feralheart Raiment 6pc (shoulder: Darkspear Pauldrons -> Feralheart Epaulets, wrist: Forest Stalker's Bracers -> Feralheart Bands, hands: Raider Gloves -> Feralheart Fists, legs: Warbear Woolies -> Feralheart Trousers, feet: Drudge Boots -> Feralheart Walkers) | - |
| druid-feral | raid | horde | 614.6 | 629.6 | +2.45% | Feralheart Raiment 6pc (shoulder: Darkspear Pauldrons -> Feralheart Epaulets, wrist: Forest Stalker's Bracers -> Feralheart Bands, hands: Raider Gloves -> Feralheart Fists, legs: Warbear Woolies -> Feralheart Trousers, feet: Drudge Boots -> Feralheart Walkers) | - |
| druid-restoration | bare | horde | 320.3 | 297.7 | -7.07% | The Postmaster 2pc (head: Living Crown -> The Postmaster's Band); Magister's Regalia 2pc (wrist: Bracers of Hope -> Magister's Bindings, hands: Feralheart Gauntlets -> Magister's Gloves) | - |
| hunter-beast-mastery | bare | alliance | 237.8 | 244.4 | +2.80% | Beastmaster Armor 4pc (wrist: Bracers of the Eclipse -> Beastmaster's Bindings, hands: Raider Gloves -> Beastmaster's Gauntlets, feet: Windreaver Greaves -> Beastmaster's Treads) | - |
| hunter-beast-mastery | bare | horde | 240.9 | 247.9 | +2.89% | Beastmaster Armor 4pc (wrist: Bracers of the Eclipse -> Beastmaster's Bindings, waist: Marksman's Girdle -> Beastmaster's Belt, feet: Windreaver Greaves -> Beastmaster's Treads) | hands: Raider Gloves -> Beaststalker's Gloves |
| hunter-beast-mastery | raid | alliance | 648.8 | 669.3 | +3.16% | Beastmaster Armor 4pc (wrist: Slashclaw Bracers -> Beastmaster's Bindings, waist: Marksman's Girdle -> Beastmaster's Belt, feet: Windreaver Greaves -> Beastmaster's Treads) | - |
| hunter-beast-mastery | raid | horde | 660.1 | 680.5 | +3.08% | Beastmaster Armor 4pc (wrist: Slashclaw Bracers -> Beastmaster's Bindings, waist: Marksman's Girdle -> Beastmaster's Belt, feet: Windreaver Greaves -> Beastmaster's Treads) | - |
| hunter-marksmanship | bare | alliance | 228.9 | 238.2 | +4.06% | Beastmaster Armor 4pc (head: Outlaw's Collar -> Beastmaster's Cap, wrist: Bracers of the Eclipse -> Beastmaster's Bindings, waist: Marksman's Girdle -> Beastmaster's Belt, feet: Windreaver Greaves -> Beastmaster's Treads) | hands: Voone's Vice Grips -> Beaststalker's Gloves |
| hunter-marksmanship | bare | horde | 235.9 | 243.6 | +3.29% | Beastmaster Armor 4pc (head: Outlaw's Collar -> Beastmaster's Cap, wrist: Bracers of the Eclipse -> Beastmaster's Bindings, waist: Marksman's Girdle -> Beastmaster's Belt, feet: Windreaver Greaves -> Beastmaster's Treads) | hands: Raider Gloves -> Beaststalker's Gloves |
| hunter-marksmanship | raid | alliance | 621.8 | 647.6 | +4.15% | Beastmaster Armor 4pc (wrist: Slashclaw Bracers -> Beastmaster's Bindings, waist: Marksman's Girdle -> Beastmaster's Belt, feet: Windreaver Greaves -> Beastmaster's Treads) | - |
| hunter-marksmanship | raid | horde | 638.3 | 659.8 | +3.37% | Beastmaster Armor 4pc (wrist: Slashclaw Bracers -> Beastmaster's Bindings, waist: Marksman's Girdle -> Beastmaster's Belt, feet: Windreaver Greaves -> Beastmaster's Treads) | - |
| hunter-survival | bare | alliance | 224.6 | 243.9 | +8.58% | Beastmaster Armor 4pc (head: Outlaw's Collar -> Beastmaster's Cap, wrist: Forest Stalker's Bracers -> Beastmaster's Bindings, waist: Belt of Preserved Heads -> Beastmaster's Belt, feet: Bloodmail Boots -> Beastmaster's Treads); Beastmaster Armor 6pc (chest: Dawn Armor -> Beastmaster's Tunic, hands: Raider Gloves -> Beastmaster's Gauntlets) | - |
| hunter-survival | bare | horde | 225.8 | 244.7 | +8.39% | Beastmaster Armor 4pc (head: Outlaw's Collar -> Beastmaster's Cap, wrist: Forest Stalker's Bracers -> Beastmaster's Bindings, waist: Belt of Preserved Heads -> Beastmaster's Belt, feet: Bloodmail Boots -> Beastmaster's Treads); Beastmaster Armor 6pc (chest: Dawn Armor -> Beastmaster's Tunic, hands: Raider Gloves -> Beastmaster's Gauntlets) | - |
| hunter-survival | raid | alliance | 686.4 | 694.9 | +1.23% | Beastmaster Armor 4pc (head: Mask of the Unforgiven -> Beastmaster's Cap, wrist: Slashclaw Bracers -> Beastmaster's Bindings, waist: Belt of Preserved Heads -> Beastmaster's Belt, feet: Windreaver Greaves -> Beastmaster's Treads) | - |
| hunter-survival | raid | horde | 686.8 | 694.4 | +1.10% | Bloodmail Regalia 3pc (hands: Voone's Vice Grips -> Bloodmail Gauntlets, waist: Belt of Preserved Heads -> Bloodmail Belt, feet: Windreaver Greaves -> Bloodmail Boots) | - |
| mage-arcane | bare | horde | 470.6 | 470.7 | +0.03% | - | hands: Sorcerer's Gauntlets -> Sorcerer's Gloves |
| mage-fire | bare | alliance | 403.5 | 411.6 | +2.01% | Sorcerer's Regalia 4pc (hands: Sorcerer's Gloves -> Sorcerer's Gauntlets, legs: Sentinel's Silk Leggings -> Sorcerer's Leggings, feet: Sorcerer's Boots -> Sorcerer's Sandals) | - |
| mage-frost | bare | alliance | 323.6 | 329.4 | +1.81% | Sorcerer's Regalia 4pc (hands: Sorcerer's Gloves -> Sorcerer's Gauntlets, legs: Sentinel's Silk Leggings -> Sorcerer's Leggings, feet: Sorcerer's Boots -> Sorcerer's Sandals) | - |
| mage-frost | bare | horde | 311.0 | 320.9 | +3.19% | Sorcerer's Regalia 3pc (hands: Sorcerer's Gloves -> Sorcerer's Gauntlets, feet: Sorcerer's Boots -> Sorcerer's Sandals); Sorcerer's Regalia 4pc (legs: Sentinel's Silk Leggings -> Sorcerer's Leggings) | - |
| paladin-holy | bare | alliance | 168.6 | 173.1 | +2.65% | Magister's Regalia 2pc (shoulder: Darkspear Shoulderpads -> Magister's Mantle, chest: Knight-Captain's Lamellar Chestplate -> Magister's Robes); Magister's Regalia 5pc (wrist: Gallant's Wristguards -> Magister's Bindings, waist: Belt of Tiny Heads -> Magister's Belt, feet: Soulforge Treads -> Magister's Boots) | - |
| paladin-protection | bare | alliance | 160.9 | 186.4 | +15.84% | Soulforge Armor 3pc (chest: Ornate Adamantium Breastplate -> Soulforge Chestguards, wrist: Vigorsteel Vambraces -> Soulforge Wristguards) | waist: Redeemer's Waistcord -> Deathbone Girdle |
| paladin-protection | bare | horde | 177.1 | 212.9 | +20.20% | Soulforge Armor 4pc (chest: Ornate Adamantium Breastplate -> Soulforge Chestguards, hands: Heavy Thorium Gauntlets -> Soulforge Handguards); Soulforge Armor 2pc (feet: Boots of Avoidance -> Soulforge Sabatons) | - |
| paladin-protection | raid | alliance | 382.1 | 390.0 | +2.08% | Soulforge Armor 3pc (head: Helm of Awareness -> Soulforge Faceguard, wrist: Sentinel's Wristguards -> Soulforge Wristguards) | - |
| paladin-retribution | bare | alliance | 254.5 | 268.8 | +5.60% | Soulforge Armor 6pc (shoulder: Darkspear Epaulets -> Soulforge Pauldrons, chest: Timbermaw Tunic -> Soulforge Breastplate, wrist: Battleborn Armbraces -> Soulforge Bracers, hands: Voone's Vice Grips -> Soulforge Gauntlets, waist: Radiant Girdle of the Dawn -> Soulforge Belt, feet: Windreaver Greaves -> Soulforge Sabatons) | - |
| paladin-retribution | bare | horde | 266.5 | 269.6 | +1.14% | The Gladiator 3pc (chest: Timbermaw Tunic -> Savage Gladiator Chain, hands: Voone's Vice Grips -> Savage Gladiator Grips, feet: Windreaver Greaves -> Savage Gladiator Greaves) | - |
| paladin-retribution | raid | alliance | 538.5 | 556.7 | +3.37% | Soulforge Armor 6pc (shoulder: Darkspear Pauldrons -> Soulforge Pauldrons, chest: Timbermaw Tunic -> Soulforge Breastplate, wrist: Battleborn Armbraces -> Soulforge Bracers, hands: Voone's Vice Grips -> Soulforge Gauntlets, waist: Radiant Girdle of the Dawn -> Soulforge Belt, feet: Bloodmail Boots -> Soulforge Sabatons) | - |
| paladin-retribution | raid | horde | 551.5 | 558.5 | +1.27% | Bloodmail Regalia 2pc (hands: Voone's Vice Grips -> Bloodmail Gauntlets) | - |
| priest-discipline | raid | alliance | 578.3 | 583.9 | +0.98% | Vestments of the Virtuous 4pc (waist: Wisdom of the Timbermaw -> Virtuous Belt) | trinket2: Darkspear Voodoo Seal -> Serenity Field |
| priest-holy | bare | alliance | 330.3 | 331.2 | +0.27% | Vestments of the Virtuous 4pc (head: Mooncloth Circlet -> Virtuous Crown, shoulder: Mooncloth Shoulders -> Virtuous Mantle, wrist: Bracers of Hope -> Virtuous Bracers, feet: Mooncloth Boots -> Virtuous Sandals) | chest: Truefaith Vestments -> Robes of the Exalted |
| priest-holy | bare | horde | 312.3 | 323.2 | +3.47% | Vestments of the Virtuous 4pc (shoulder: Mooncloth Shoulders -> Virtuous Mantle, wrist: Bracers of Hope -> Virtuous Bracers, feet: Mooncloth Boots -> Virtuous Sandals) | trinket2: Mindtap Talisman -> Serenity Field |
| priest-shadow | bare | alliance | 321.8 | 326.2 | +1.38% | Vestments of the Virtuous 6pc (legs: Sentinel's Silk Leggings -> Virtuous Leggings) | - |
| priest-shadow | bare | horde | 324.4 | 329.2 | +1.46% | Vestments of the Virtuous 6pc (legs: Outrider's Silk Leggings -> Virtuous Leggings) | - |
| rogue-assassination | bare | alliance | 275.2 | 288.4 | +4.81% | Stormshroud Armor 3pc (shoulder: Truestrike Shoulders -> Stormshroud Shoulders, legs: Sentinel's Leather Pants -> Stormshroud Pants) | - |
| rogue-assassination | bare | horde | 274.5 | 290.1 | +5.68% | Stormshroud Armor 3pc (shoulder: Truestrike Shoulders -> Stormshroud Shoulders, legs: Sentinel's Leather Pants -> Stormshroud Pants) | - |
| rogue-assassination | raid | alliance | 646.4 | 682.3 | +5.55% | Stormshroud Armor 3pc (shoulder: Truestrike Shoulders -> Stormshroud Shoulders, legs: Sentinel's Leather Pants -> Stormshroud Pants) | - |
| rogue-assassination | raid | horde | 643.8 | 676.1 | +5.03% | Stormshroud Armor 3pc (shoulder: Truestrike Shoulders -> Stormshroud Shoulders, legs: Sentinel's Leather Pants -> Stormshroud Pants) | - |
| rogue-combat | bare | alliance | 257.3 | 265.1 | +3.03% | Stormshroud Armor 3pc (shoulder: Truestrike Shoulders -> Stormshroud Shoulders, hands: Raider Gloves -> Stormshroud Gloves, legs: Sentinel's Leather Pants -> Stormshroud Pants) | - |
| rogue-combat | bare | horde | 253.8 | 264.7 | +4.27% | Stormshroud Armor 3pc (shoulder: Truestrike Shoulders -> Stormshroud Shoulders, hands: Raider Gloves -> Stormshroud Gloves, legs: Sentinel's Leather Pants -> Stormshroud Pants) | off_hand: The Lobotomizer -> Felstriker |
| rogue-combat | raid | alliance | 623.1 | 659.4 | +5.82% | Stormshroud Armor 3pc (shoulder: Truestrike Shoulders -> Stormshroud Shoulders, legs: Sentinel's Leather Pants -> Stormshroud Pants) | off_hand: The Lobotomizer -> Felstriker |
| rogue-combat | raid | horde | 618.6 | 654.8 | +5.85% | Stormshroud Armor 3pc (shoulder: Truestrike Shoulders -> Stormshroud Shoulders, legs: Sentinel's Leather Pants -> Stormshroud Pants) | off_hand: The Lobotomizer -> Felstriker |
| rogue-subtlety | bare | alliance | 248.6 | 256.4 | +3.15% | Stormshroud Armor 3pc (shoulder: Truestrike Shoulders -> Stormshroud Shoulders, hands: Raider Gloves -> Stormshroud Gloves, legs: Sentinel's Leather Pants -> Stormshroud Pants) | - |
| rogue-subtlety | bare | horde | 250.1 | 257.5 | +2.98% | Stormshroud Armor 3pc (shoulder: Truestrike Shoulders -> Stormshroud Shoulders, hands: Raider Gloves -> Stormshroud Gloves, legs: Sentinel's Leather Pants -> Stormshroud Pants) | - |
| rogue-subtlety | raid | alliance | 528.5 | 547.5 | +3.60% | Stormshroud Armor 3pc (shoulder: Truestrike Shoulders -> Stormshroud Shoulders, hands: Raider Gloves -> Stormshroud Gloves, legs: Sentinel's Leather Pants -> Stormshroud Pants) | - |
| rogue-subtlety | raid | horde | 527.2 | 547.3 | +3.82% | Stormshroud Armor 3pc (shoulder: Truestrike Shoulders -> Stormshroud Shoulders, hands: Raider Gloves -> Stormshroud Gloves, legs: Sentinel's Leather Pants -> Stormshroud Pants) | - |
| shaman-elemental | bare | alliance | 157.5 | 162.2 | +3.01% | Ironfeather Armor 2pc (shoulder: Rugged Mantle of the Timbermaw -> Ironfeather Shoulders, chest: Vest of Elements -> Ironfeather Breastplate) | - |
| shaman-elemental | bare | horde | 156.0 | 163.6 | +4.87% | The Five Thunders 3pc (chest: Vest of Elements -> Vest of The Five Thunders); The Five Thunders 4pc (hands: Raider Handguards -> Gauntlets of The Five Thunders) | - |
| shaman-enhancement | bare | alliance | 226.0 | 234.8 | +3.87% | Bloodmail Regalia 3pc (hands: Voone's Vice Grips -> Bloodmail Gauntlets, waist: Ferocity of the Timbermaw -> Bloodmail Belt, feet: Windreaver Greaves -> Bloodmail Boots) | - |
| shaman-enhancement | bare | horde | 223.5 | 226.6 | +1.41% | The Gladiator 2pc (chest: Dawn Armor -> Savage Gladiator Chain, feet: Windreaver Greaves -> Savage Gladiator Greaves) | - |
| warrior-arms | bare | alliance | 259.3 | 267.0 | +2.98% | The Highlander's Resolution 3pc (shoulder: Truestrike Shoulders -> Highlander's Plate Spaulders, waist: Radiant Girdle of the Dawn -> Highlander's Plate Girdle, feet: Boots of Heroism -> Highlander's Plate Greaves) | hands: Voone's Vice Grips -> Bloodmail Gauntlets |
| warrior-arms | bare | horde | 257.6 | 264.7 | +2.75% | The Defiler's Resolution 3pc (shoulder: Truestrike Shoulders -> Defiler's Plate Spaulders, waist: Radiant Girdle of the Dawn -> Defiler's Plate Girdle, feet: Boots of Heroism -> Defiler's Plate Greaves) | hands: Voone's Vice Grips -> Bloodmail Gauntlets |
| warrior-arms | raid | alliance | 683.6 | 703.6 | +2.92% | Battlegear of Heroism 6pc (shoulder: Truestrike Shoulders -> Spaulders of Heroism, chest: Dawn Armor -> Breastplate of Heroism, wrist: Battleborn Armbraces -> Bracers of Heroism, hands: Voone's Vice Grips -> Gauntlets of Heroism, waist: Radiant Girdle of the Dawn -> Belt of Heroism, feet: Boots of Heroism -> Battleboots of Heroism) | - |
| warrior-arms | raid | horde | 693.6 | 701.5 | +1.14% | The Gladiator 3pc (chest: Dawn Armor -> Savage Gladiator Chain, hands: Voone's Vice Grips -> Savage Gladiator Grips, feet: Boots of Heroism -> Savage Gladiator Greaves) | - |
| warrior-fury | raid | alliance | 782.6 | 792.4 | +1.25% | Battlegear of Heroism 2pc (hands: Voone's Vice Grips -> Gauntlets of Heroism, feet: Windreaver Greaves -> Battleboots of Heroism) | - |

## 6. Caveats

- Paladin-protection bare gains +16% and +20% (tank score); the raid rows gain 2.1%, so it is a bare-preset effect.
- Hunter-survival bare +8.6% comes from the Beastmaster 5- and 6-piece bonuses the engine now models; the raid rows gain 1.1 to 1.2%.
- The `bonus` text is the client's description of the tier named by `pieces`; when a trial wears pieces for a threshold that also unlocks lower tiers, the note names the highest tier only.
- Not run: the full web vitest and Playwright suites (shared layout was not touched; GearRow gained a line).
