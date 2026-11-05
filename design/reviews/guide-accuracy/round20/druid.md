# Guide accuracy round 20: Druid

Audited on the round 20 audit tree (4234643f5: main after guide round 19 and sim round 20; engine pin d5c708834 in `sim/enginever/version.go`, source read from `scratchpad/gr20/engine`), against the ACTIVE client 1.60.1.70790 under `data/builds/1.60.1.70790/`. Files: `web/src/content/guides/druid/{index,balance,feral,feral-bear,restoration}.md`, every card note in `data/curated/apl/druid-{balance,feral,feral-bear,restoration}.json`, and every rendered row of `addon-data.json` (`rotations/druid-balance` 22 bands, `druid-feral` 24, `druid-feral-bear` 20, `druid-restoration` 28), read band by band against BOTH shipped point orders in `addon-builds.json` (leveling and raid; index 0 = level 10) and the trainer levels in `trainables/druid.json`. Read first: round 19 druid report, sim round 20 `druid.md` and `integration.md`, the ledger's Round 19 section, and the round 1 and 2 briefs.

Main lens: every sentence of every card note, at every band its row renders on. Not filed (as round 19): rows that render before their own talent exists (Feral Shifting Power 10 to 39, Balance Insect Swarm 20 to 26, Restoration Nature's Swiftness 10 to 29 and Swiftmend 10 to 42, the Feral "Below level 20" rows past 20) are the component lane's whole-row issue.

## Changes since round 19's tree (6beb919c2)

- Druid card notes: round 19's seven fixes (Balance Wrath x3 and Starfire; Feral Rip; bear Lacerate and Maul; Restoration Regrowth and tank Rejuvenation) and sim round 20's Restoration tank Rejuvenation and Healing Touch filler notes plus the file note.
- Prose: `balance.md:135` (Moonfire and Wrath rank sentence, sim L1/L2); `restoration.md:83`, `:85`, `:121` (sim R2, R3a to R3c).
- `rank-choices.json`: Balance Wrath rank 3 to 23, Moonfire rank 2 to 28 and rank 3 from 29; Restoration tank Rejuvenation rank 2 to 39 with the night elf rank 3 at 29, party rank 2 at 27 to 29, Healing Touch filler rank 2 at 20 and tauren 21 to 23. `addon-data.json` regenerated to match (checked row by row below).
- `addon-builds.json`: no druid change. No druid page or note names 70755; no em or en dashes.

## Shipped orders (decoded from `addon-builds.json`)

- Balance leveling: Improved Wrath 10-14, Genesis 15-16, Moonglow 17-19, Improved Moonfire 20-21, Nature's Majesty 22-23, Nature's Reach 24-25, Nature's Splendor 26, Insect Swarm 27, Vengeance 28-32, Improved Starfire 33, Nature's Grace 34, **Eclipse 35**, Moonfury 36-39 and 41, Moonkin Form 40, Nature's Focus 42-46, Heart of the Wild 47-51, Naturalist 52-56, Reflection 57-59, Subtlety 60.
- **Balance raid: identical to 32, then Improved Starfire 33-37, Nature's Grace 38, Eclipse 39 (one point), Moonkin Form 40, Eclipse 41-42 (three)**, Moonfury 43-47, Nature's Focus 48-52, Naturalist 53-57, Subtlety 58-60.
- Feral leveling: Ferocity 10-14, Feral Swiftness 15-16, Feral Instinct 17, Brutal Impact 18-19, Shredding Attacks 20-22, Savage Fury 23-24, Feral Charge 25, Sharpened Claws 26-27, Predatory Strikes 28-30, Predatory Discipline 31, Blood Frenzy 32-33, Leader of the Pack 34, Natural Instinct 35-36, **Rend and Tear 37-39**, Shifting Power 40, then Genesis, Nature's Majesty, Nature's Reach, Heart of the Wild, Naturalist, Natural Shapeshifter.
- **Feral raid: Ferocity 10-14, Feral Swiftness 15-16, Feral Instinct 17-19, Shredding Attacks 20-22, Savage Fury 23-24, Sharpened Claws 25-26, Predatory Strikes 27-29, Predatory Discipline 30, Blood Frenzy 31-32, Leader of the Pack 33, Natural Instinct 34-35, Rend and Tear 36-39 and 41**, Shifting Power 40, Genesis 42-46, Nature's Majesty 47-48, Nature's Reach 49-50, Heart of the Wild 51-55, Naturalist 56-60.
- Bear (leveling = raid): Ferocity 10-14, Furor 15, Feral Swiftness 16-17, Feral Instinct 18-19 and 22, Savage Fury 20-21, Thick Hide 23-24 and 26, Lacerate 25, Shredding Attacks 27-29, Leader of the Pack 30, Sharpened Claws 31-32, Predatory Strikes 33-35, Blood Frenzy 36-37, Natural Instinct 38-39, Primal Bite 40, Rend and Tear 41-45, Natural Reaction 46-50, Heart of the Wild 51-55, Naturalist 56-60.
- Restoration (leveling = raid): Heart of the Wild 10-14, Reflection 15-17, Nature's Focus 18-19 and 21-23, Bloom 20, Naturalist 24-28, Gift of Nature 29 and 31-34, Nature's Swiftness 30, Tranquil Spirit 35-39, Improved Rejuvenation 40-42, Swiftmend 43, Living Spirit 44-46, Improved Tranquility 47-48, Improved Regrowth 49, Genesis 50-54, Moonglow 55-57, Nature's Majesty 58-59, Nature's Splendor 60.

Balance and Feral carry a separate `raidBuild`, and the addon offers both order slots (`Follow.SLOTS`, `/fs build raid`); round 19 already worded Arcane's Missile Barrage note "on the leveling order and ... on the raid order". Round 19's druid fixes read only the leveling order, which is where R20-1 and R20-3 come from.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| DRUID-R20-1 | CONFIRMED | card note (talent before the raid order) | `druid-balance.json` Wrath note (priorityList[5], :177), first sentence, rendered on bands 10 to 60; Starfire note (priorityList[7], :251), rendered 20 to 60 | Wrath: "From 35, whenever no Eclipse charge is banked." Starfire: "From 35 it spends the Eclipse charges Wrath banked." | Eclipse is the leveling order's point at 35 but the raid order's point at 39 (`[1,5,3]` at index 29; Improved Starfire 33-37, Nature's Grace 38). The Balance page ships both orders (`build` and `raidBuild`) and the addon offers both slots. A moonkin on the raid order reads "From 35" on bands 36 and 38 with no Eclipse; the card gates the Wrath on `auraIsKnown(408248)`, registered only with the talent (`sim/druid/talents.go` applyEclipse), so the row never fires for him until 39. The first sentence is all the addon row shows. Round 19's own fix (R19-1) put the level in on the leveling order alone. The rest of the note holds on both orders: one point at 35 to 60 (leveling) and at 39 to 40 (raid), 0.17 s a charge, three points by 42 on the raid order, 0.50 s (talents/druid.json Eclipse ranks 1 and 3). | Wrath: "From 35 on the leveling order and 39 on the raid order, whenever no Eclipse charge is banked. Each Wrath banks two charges, each taking time off one Starfire: 0.17 seconds a charge with one point, half a second with all three, so in the raid build one Wrath and two quick Starfires beat three plain ones." Starfire: "Everything else. Once you have Eclipse (35 leveling, 39 raid order) it spends the charges Wrath banked." |
| DRUID-R20-2 | CONFIRMED | card note (contradicts its sibling note and its own reason) | `druid-feral.json` Shred note (priorityList[8], :289), first sentence, rendered on bands 20 to 60 | "Your builder from level 20, always from behind the target." | On band 20 (levels 20 and 21) the same card's Claw note says "Your builder at 20 and 21, before Shredding Attacks is full: Claw costs 40 to Shred's 54 at 20 and 48 at 21, and does more for its energy", and the Shred note's own second sentence agrees ("until Shredding Attacks is full at 22 ... Claw does more for the energy"), as does feral.md:119 ("Shred from behind once Shredding Attacks has it at 42 energy; before that, Claw"). Shredding Attacks is 20-22 on both orders. The addon row shows only the first sentence, so a level 20 cat sees two rows each calling itself "your builder", Shred on top. The card itself never casts Shred before 24 (row 8 is gated on `spellIsKnown(Ferocious Bite)`, the 24 gate; see Sim findings). | "Your builder from level 22, always from behind the target. Shred hits harder than Claw at every rank, but until Shredding Attacks is full at 22 it costs 8 to 14 more energy, and Claw does more for the energy." |
| DRUID-R20-3 | CONFIRMED | card note (talent level on one order only) | `druid-feral.json` Rip note (priorityList[6], :193), third sentence, rendered on bands 20 to 60 | "Rip comes first, and from 37 it is also the bleed that keeps Rend and Tear on." | Rend and Tear is the leveling order's points at 37 to 39 but the raid order's at 36 to 39 and 41 (`[2,6,1]` at index 26). Band 36 renders (Rip rank 3, 9493); a cat on the raid order has two points of Rend and Tear there (three at 37), so "from 37" is a level late for him. The smallest finding: one band, an understatement. | "Rip comes first, and once you have Rend and Tear (37 leveling, 36 raid order) it is also the bleed that keeps it on." |

Counts: CONFIRMED 3, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

## Shoulds (wording, not findings)

- **restoration.md:85** "From 21 (24 for a tauren) to 29 and from 50, cast the top rank." Below 20 the filler is also the top rank (rows 10 to 18: ranks 2 and 3, the newest each band) and the sentence no longer says so since sim round 20 replaced "Below 30". Suggested: "Below 20, from 21 (24 for a tauren) to 29 and from 50, cast the top rank."
- **feral-bear.md:70** Before the pull: sim round 20's integrator gave the bear his own Thorns from 6 (S eba3eb7cc, +0.7 to +1.0% threat at raid 60) because every bear buffs it before shifting; the page never says to. Suggested: "Thorns on yourself, then Dire Bear Form."
- **feral.md:67** "The raid build skips Furor." Neither build takes Furor (leveling order has no Furor point either); "Both builds skip Furor" reads truer.
- Carried from round 19 (none applied): Feral Ferocious Bite "With Predatory Discipline" (from 31 leveling, 30 raid), Balance Moonfire "With Improved Moonfire and Nature's Splendor" (from 26), bear Demoralizing Roar "204 ... at 60" on rank 1 bands, Feral and bear Faerie Fire "In Cat/Bear Form it is free" (Cat Form row at 18), and the 14 carried before that (balance.md:135 "until 39" to "through 38", balance.md:115 and restoration.md:101 "for good", restoration.md:57 "your next Healing Touch" to "your next Nature spell", restoration.md:83 old "(rank 3 is not worth it)" is gone, so drop that one; feral.md:89, restoration.md:105, feral.md:67, feral.md:73, balance.md:119, balance.md:105 twice, feral.md:77, feral.md:85, balance.md:91).

## Checked and right this round

- **Sim round 20's druid edits.** `balance.md:135`: Moonfire rank 2 through 28 (rows 20, 24, 26 cast 8924), rank 3 29 to 35 (8925 at 29, 30, 34), rank 4 36 to 39 (8926); Wrath rank 3 at 22 and 23 (rank 4 trains at 22; band 20 casts 5178, the newest), rank 4 from 24 through 39 (5179 at 24 to 39 while rank 5 trains at 30 and rank 6 at 38). `restoration.md:83`: party rank 2 at 27 to 29 and 40 to 42 (1058 on those bands, 774 elsewhere below 60); tank rank 2 20 to 39, rank 3 at 29 for a night elf (band 29 race key 1430; tauren and both Skyborne 1058), rank 4 40 to 43, 5 44 to 48, 6 49 to 56, 8 57 to 59 (2090, 2091, 3627, 9839), "below the newest rank you know all the way to 59, rank 10 included once you learn it at 58" (rank 10 9841 at 58). `:85`: Healing Touch filler rank 2 at 20 (5186 on band 20), tauren to 23 (band 21 race key 5186), rank 3 30 to 33, rank 4 34 to 39, rank 5 40 to 45, rank 6 46 to 49, top rank 24 to 29 and from 50. `:121` matches `:83` and `:85`. Restoration file note matches the rows.
- **Round 19's seven fixes**, re-read band by band: R19-1 and R19-3 hold on the leveling order (R19-1 fails on the raid order: R20-1); R19-2 (Innervate 40, Starfire 20); R19-4 holds on the leveling order (R20-3 for raid); R19-5 (Rend and Tear 41, bear has one order); R19-6 (Ferocity 1 a point 10 to 14, Lacerate 25, Primal Bite 40); R19-7 (Bloom 20).
- **Client values re-read this round**: Bear Form (Passive) 1178 and Dire Bear Form (Passive) 9635 carry the Faerie Fire -100% cost (aura 108, misc 14) and +6000 ms cooldown (aura 107, misc 11) on mask 1024, as Cat Form (Passive) 3025 does, so "free on a 6 second cooldown" in Bear Form (bear note, feral-bear.md:72, index.md:28) is right (70790 SpellEffect). Engine Shred 2.00 x (weapon + 22/32/44/64/80) against Claw 1.40 x (weapon + 27/39/57/88/115): "Shred hits harder than Claw at every rank" holds. Eclipse 0.17 / 0.50 s, Rend and Tear 3% a point on bleeding targets, Predatory Discipline, Shifting Power (prereq Predatory Discipline), Primal Bite (prereq Leader of the Pack), Nature's Swiftness (prereq Naturalist 5), Wild Growth (prereq Living Spirit 3) as the pages say.
- **Page prose, all five pages, line by line**: balance.md:45 to :137, feral.md:40 to :119, feral-bear.md:30 to :121, restoration.md:40 to :121, index.md. Every trainer level (Wrath 6/14, Moonfire 4, Roots 8, Grasp 10, Faerie Fire 18, Cat Form 20, Bash 14, Swipe 16, Maul 10 and 18, Growl 10, Enrage 12, Rake and Bite 24, Ravage and Poised 32, Pounce 36, Frenzied Regeneration 36, Barkskin 44, Dire Bear 40, Innervate 40, Regrowth 12, Revive 12, Cure Poison 14, Rejuvenation 4) and every leveling talent sentence (balance.md:127, feral.md:113, feral-bear.md:115, restoration.md:113) matches the trainables and the orders. Build sums: bear 41 Feral and 10 Restoration; Restoration 40 and 11; Balance raid move (Heart of the Wild 5 and Reflection 3 into Improved Starfire +4, Eclipse +2, Subtlety +2); Feral respec (Brutal Impact 2, Feral Charge 1, Natural Shapeshifter 1 into Rend and Tear +2, Feral Instinct +2).
- **Races, sources, voice**: unchanged since round 19 and as checked there; Skyborne only as the paid option; no em or en dashes; no nightly numbers; no site or sim talk.

## Decisions

- **D1. Notes name the level on both orders where the page ships both.** Balance (Eclipse 35 / 39) and Feral (Rend and Tear 37 / 36) carry a raid build whose order the addon offers as its own slot; the round 19 Arcane precedent words both. The bear and Restoration orders are one order each, so their notes need one level.
- **D2. R20-2 fixes the note, not the card.** Notes, the Claw note and feral.md:119 agree that Shred takes over when Shredding Attacks is full at 22; the card's Shred gate sits on Ferocious Bite (24), a card action the guide lane does not own (known and accepted: card actions the research moved away from). The card gate goes to the sim lane below.
- **D3. Unmet conditionals stay Shoulds** (round 19 D3).
- **D4. Sim round 20's druid report says "Eclipse is a level-40+ talent" (S1).** It is 35 on the leveling order and 39 on the raid order; the open log question (a resisted Wrath followed by Starfire) can be read off any moonkin log from 35, not only 40+. Report text only, nothing on a page.

## Sims requested

- **Feral builder at 21, 22 and 23, guide leveling build, behind the target** (Claw against Shred; seeds 1 and 50001, night elf and tauren). On engine numbers a 5-point Rip cycle at 22 (Shred 42 energy at 2.00 x weapon + 22, Claw 40 at 1.40 x weapon + 27) favours Shred by roughly a tenth; at 21 (Shred 48) they are close. Decides the card's Shred gate (now `spellIsKnown(Ferocious Bite)`, 24) and whether the Claw note's "and 48 at 21, and does more for its energy" holds.
- Carried, not re-filed: Balance leveling Eclipse Wrath on and off at 36, 40 and 50 (round 19, routed); the Balance card's empty filler at 10 to 16 (round 19 sim finding).

## Sim findings

- **Feral card keeps Claw as the builder at 22 and 23.** Row 9 (Claw) is gated `not spellIsKnown(22829)` and row 8 (Shred) `spellIsKnown(22829)`, so the switch is Ferocious Bite's trainer level, 24, not Shredding Attacks' third point at 22 that the notes and feral.md:119 name. The `expected_idle` reason ("rank 1 Shred does less there") dates from the 155% Shred; the engine now has 200%. Gate the switch on Shredding Attacks 3 (or level 22) if the sim above agrees.

## Engine questions

- **Rend and Tear reads "Bleeding targets"** (talents/druid.json: "Increases damage done by your melee abilities on Bleeding targets"); the engine counts only the druid's own Lacerate, Rip, Rake and Pounce unless `AssumeBleedActive` is set (`sim/druid/feral_talents.go` targetIsBleeding), and no preset sets it. In a raid a warrior's Deep Wounds or Rend, or a rogue's Rupture or Garrote, would turn it on before the cat's first Rip or the bear's first Lacerate. Small (opener seconds only, Rip and Lacerate stay up), but the client text is wider than the engine. Not a guide finding: feral.md:59 and :73 say Rip keeps it on, which is true.

## Sources

- **Client 1.60.1.70790** (`data/builds/1.60.1.70790/`): `addon-builds.json` (druid leveling and raid orders), `addon-data.json` (`rotations/druid-*`, every band and race key), `talents/druid.json` (Eclipse, Rend and Tear, Shredding Attacks, Ferocity, Predatory Discipline, Bloom, Furor, prerequisites), `trainables/druid.json`; raw `SpellEffect.csv` (70790) rows for 1178, 9635, 3025, 21178.
- **Curated**: `data/curated/apl/druid-{balance,feral,feral-bear,restoration}.json`, `data/curated/rank-choices.json` (round 20 diff).
- **Engine d5c708834**: `sim/druid/shred.go`, `claw.go` (weapon multipliers and flats), `feral_talents.go` (RendAndTearMultiplier, targetIsBleeding), `talents.go` (applyEclipse).
- **Site and addon code**: `web/src/lib/guides/rotation-lists.ts` (clauseOf), `web/src/lib/bis/rotation-view.ts`, `apl-level-gate.ts` (Shred "from level 24" gate), `addon/ForeverSixty/Follow.lua` (order slots).
- **Reports**: `round19/druid.md`, `2026-11-04-sim-audit-round20/druid.md` and `integration.md` (Thorns S eba3eb7cc), `ledger.md` Round 19, `round19` Arcane wording precedent via the ledger.
- No web research was needed: every call this round is settled by the client rows, the orders and the engine.

## Note sweep

One line per card row note: spec [index] row, bands it renders, each named talent, spell, rank or value with the level the player has it, verdict.

- Balance [1] Moonfire end-of-fight, bands 10-60: Moonfire (4), Wrath cast time; "best instant" true at every band (Insect Swarm has no hit). OK.
- Balance [2] Wrath end-of-fight, 10-60: Starfire (20), "before you have Starfire at 20 ... last few seconds" (the 3.5 s branch), Wrath lands. OK.
- Balance [3] Moonfire upkeep, 10-60: Improved Moonfire (20-21), Nature's Splendor (26), conditional. OK (Should, carried).
- Balance [4] Insect Swarm, 20-60: talent 27 (whole-row, not filed), one global. OK.
- Balance [5] Eclipse Wrath, 10-60: Eclipse 35 leveling / 39 raid, 0.17 s at one point, 0.50 s at three (raid 42). **DRUID-R20-1.**
- Balance [6] mana Wrath, 10-60: Innervate (40), Improved Wrath (10-14), Starfire (20), half the cast. OK.
- Balance [7] Starfire, 20-60: Eclipse 35 / 39. **DRUID-R20-1.**
- Feral [1] Ravage, 32-60: Prowl (20), Poised (32). OK.
- Feral [2] Faerie Fire, 18-60: Cat Form (20), free on 6 s (3025). OK (band 18 conditional, Should carried).
- Feral [3] Moonfire below 20, 10-60: Cat Form 20. OK (row past 20 is component).
- Feral [4] Wrath below 20, 10-60: as [3]. OK.
- Feral [5] Shifting Power, 10-60: talent 40 (whole-row), 40 energy, mana cost, 8 s. OK.
- Feral [6] Rip, 20-60: Rip (20), 5 points, under 2 s; Rend and Tear 37 leveling / 36 raid. **DRUID-R20-3.**
- Feral [7] Ferocious Bite, 24-60: Bite (24), 5 s and 2 s windows, Shred builder from 24 on the card; Predatory Discipline 31 / 30 conditional. OK (Should, carried).
- Feral [8] Shred, 20-60: Shred (20), Shredding Attacks 20-22 both orders, 54 / 48 / 42 against Claw 40 (Ferocity full at 14). **DRUID-R20-2** (first sentence); card gate in Sim findings.
- Feral [9] Claw, 20-60: Claw (20), Ferocity 5 by 14, Shredding Attacks 20-22; "does more for its energy" at 21 in Sims requested. OK pending sim.
- Bear [0] cooldown pass, page only: Mighty Rage Potion in form. OK.
- Bear [1] Barkskin, 44-60: Barkskin (44), 20%, 15 s, 60 s, cast in form. OK.
- Bear [2] Frenzied Regeneration, 36-60: trainer 36, under 50%, 20 rage. OK.
- Bear [3] Primal Bite, 40-60: talent 40, 6 s. OK.
- Bear [4] Lacerate, 25-60: talent 25, five stacks, 3 s, Rend and Tear 41. OK.
- Bear [5] Faerie Fire, 18-60: Bear Form (10), free on 6 s (1178 / 9635 auras 108 and 107), threat. OK.
- Bear [6] Demoralizing Roar, 10-60: rank 1 at 10, 204 named at 60, last 3 s, no stack with Demoralizing Shout. OK (Should, carried).
- Bear [7] Maul, 10-60: 15 rage, Ferocity to 10 by 14, Lacerate 25, Primal Bite 40. OK.
- Bear [8] Swipe, 16-60: trainer 16, 40 rage, three targets. OK.
- Restoration [1] Innervate, 40-60: trainer 40, 30% mana, Spirit regeneration. OK.
- Restoration [2] Read Leyline, Alliance Skyborne only (no addon row): 50% mana regeneration, 15 s, Innervate interaction. OK.
- Restoration [3] Nature's Swiftness, 10-60: talent 30 (whole-row), under 35%. OK.
- Restoration [4] Healing Touch save, 10-60: top rank every band (5186 to 9889). OK.
- Restoration [5] Swiftmend, 10-60: talent 43 (whole-row), 75%, keeps the HoT, 15 s, 20% base mana. OK.
- Restoration [6] Tranquility, 30-60: trainer 30, two under 70%, every 2 s. OK.
- Restoration [7] Regrowth (tank), 12-60: trainer 12, Bloom 20. OK.
- Restoration [8] Rejuvenation (tank), 10-60: rank 10 at 58, Bloom 20, rank rows 20-39 rank 2 (night elf rank 3 at 29), 40-43 rank 4, 44-48 rank 5, 49-56 rank 6, 57-59 rank 8, all matching the rows. OK.
- Restoration [9] Healing Touch filler, 10-60: 55%, mana pace; rank 2 at 20 (tauren to 23), 3 at 30-33, 4 at 34-39, 5 at 40-45, 6 at 46-49, all matching the rows. OK.
- Restoration [10] Rejuvenation (party), 10-60: under 70%, no mana gate. OK (rows 11 to 14 carry no note).
