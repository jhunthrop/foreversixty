# Guide accuracy round 23: Druid

Audited on the round 23 audit tree (`scratchpad/gr23/audit`, head 24ee457d0: main after guide round 22 landed; no sim round 21 commit in the log, so there is no sim round 21 druid report). Engine pin d5c708834 (`sim/enginever/version.go`), source read from `scratchpad/gr23/engine`. Active client 1.60.1.70790.

Files: `web/src/content/guides/druid/{index,balance,feral,feral-bear,restoration}.md`, every row note in `data/curated/apl/druid-{balance,feral,feral-bear,restoration}.json`, and every rendered band of `addon-data.json` `rotations/druid-*` (balance 22 bands, feral 24, bear 20, restoration 28), against both shipped orders in `addon-builds.json`, `trainables/druid.json`, `spelltips/druid.json`, `spellconst/druid.json`, `talents/druid.json`, `items/druid.json` and `bis/druid-*.json`.

Read first: the round 1, 2 and 23 briefs, `round22/druid.md`, the ledger's Rounds 19 to 22.

## Changes since round 22's tree (987d08a7c)

Only the three round 22 druid lines changed (commit 42b002c21). No druid APL file, `addon-builds.json`, BiS file or talent file moved. The four `rotations/druid-*` arrays in `addon-data.json` compare equal to the round 22 tree. The frontmatter strings were decoded again and match the orders below.

Totals (Balance/Feral/Restoration): Balance leveling 32/0/19, Balance raid 38/0/13, Feral leveling 9/31/11, Feral raid 9/32/10, Bear 0/41/10, Restoration 11/0/40.

## Round 22's fixes, rechecked

- **feral.md:119, "From 25, Feral Charge's Cat Form leap".** Holds. Feral Charge is index 15 (level 25) on the leveling order. The raid order has no Feral Charge, and this paragraph is the leveling one and ends on the respec. The leap is spell 414924; Blizzard's deep dive confirms the Cat Form version with a 30 second cooldown. Shred renders before 25, so the leap's purpose is live.
- **balance.md:135, "Insect Swarm, your talent at 27, stays at rank 1 through 38".** Holds. Index 17 (27) on both orders. Rows: rank 1 on bands 20 to 38, rank 2 at band 39, rank 3 at 40. Rank 2 trains at 30, so 30 to 38 is a true downrank.
- **restoration.md:91, "Healing power is every heal but Bloom".** The Bloom half holds (1328622 coefficient 0; `bloom.go`). The new "every ... but" now names one exception and leaves out the one the page gives ten lines up. See DRUID-R23-2.

## Findings

### DRUID-R23-1: feral-bear.md:107 "Nothing in your raid set binds on pickup" (CONFIRMED, kind: professions vs BiS)

- **Line:** "Nothing in your raid set binds on pickup, so no profession is needed for it: the Leatherworking and Tailoring pieces bind when equipped and can be bought."
- **Evidence:** `bis/druid-feral-bear.json`, band 60 raid, read against `items/druid.json` `bonding`. All eleven Alliance pieces outside the four crafts are "pickup" (Sentinel's Medallion, Highlander's Leather Shoulders, Forest Stalker's Bracers, Theramore Gloves, Sentinel's Leather Pants, Highlander's Leather Boots, Theramore Signet, Militant's Band, Talisman of Arathor, Stormpike Insignia Rank 6, The Unstoppable Force), and so are the Horde set's eleven. What the sentence means is true: all four crafts (Outlaw's Collar, Fel Cape, Warbear Harness, Hivethrasher's Girdle) bind when equipped.
- **Why it counts:** round 16 read this as a statement about crafts and passed it. As written it is false about the set, and a reader looking at their table sees rep and PvP pieces that bind on pickup. The colon clause does not rescue the subject of the sentence.
- **Fix:** "No craft in your raid set binds on pickup, so no profession is needed for it: the Leatherworking and Tailoring pieces bind when equipped and can be bought."

### DRUID-R23-2: restoration.md:91 "every heal but Bloom" against restoration.md:81 (CONFIRMED, kind: stats, page contradiction, round 22 side effect)

- **Line :91:** "Healing power is every heal but Bloom, and your HoTs take a big share of it per tick."
- **Line :81:** "at 60, ranks 8, 9 and 10 keep all of it, rank 7 most of it, and rank 1 none".
- **Evidence:** `sim/core/spell.go` level rule (grace 12, 5% a level past the rank's max level). Rejuvenation rank 1 (774) has `max_level` 9 in `spellconst/druid.json`, so it keeps no healing power from level 41. The card casts it on the party on bands 43 to 59, and :81 tells a 60 to cast it in a long five-man. So on those rows the heal takes no healing power, as the page itself says.
- **Shape:** the round 22 fix narrowed an "every" claim and left a second exception standing. This is the "every" shape from the round 22 notes, and the check on a fix's own side effects from round 23's.
- **Fix:** "Healing power is every heal except Bloom and a rank far below your level, and your HoTs take a big share of it per tick."

Counts: CONFIRMED 2, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0. Both are low severity: neither moves a button, a rank or a number on the card. No card note has a finding.

## Round 22 and 23 shapes, checked on every page and note

- **A spell or talent named in a leveling paragraph before the player has it.** All four "1 to 19" and "20 to 60" blocks read against the trainer and both orders.
  - Balance: Wrath 1 (6, 14), Moonfire 4, Entangling Roots 8, Nature's Grasp 10, Faerie Fire 18, Improved Wrath 10-14, Genesis 15-16, Moonglow 17-19, Moonkin 40, Insect Swarm 27. OK.
  - Feral: Moonfire 4, Bear Form 10, Bash 14, Swipe 16, Ferocity 10-14, Feral Swiftness 15-16, Feral Instinct 17, Brutal Impact 18-19, Cat Form, Claw, Prowl, Rip 20, Shredding Attacks 20-22, Rake and Bite 24, Feral Charge 25, Ravage 32, Pounce 36. OK.
  - Bear: Maul, Growl, Roar 10, Enrage 12, Bash 14, Swipe 16, Maul 2 at 18, Furor 15, Feral Swiftness 16-17, Feral Instinct 18-19, Lacerate 25, Frenzied Regeneration 36, Primal Bite and Dire Bear 40. OK.
  - Restoration: Healing Touch 1, Rejuvenation 4, Regrowth 12, Revive 12, Cure Poison 14, Heart of the Wild 10-14, Reflection 15-17, Nature's Focus 18-19 (28%, "more than a quarter"). OK.
- **A later note clause naming a spell or rank on bands before it is trained.** Every clause is scoped by level: Balance [2], [5], [6], [7]; Feral [6], [8], [9]; Bear [4], [7]; Restoration [7], [8], [9]. See the sweep.
- **"Cheap" or "best" list containing its own top.** None. Restoration :81 puts rank 7, not the top, on the party. Balance :135-137 lists only ranks below the newest; each matches the rendered rows, including band 59's Insect Swarm rank 4 (the newest there, so "rank 3 from 54 to 58" stops correctly).
- **True only among one spell's ranks.** Balance :135 and Restoration :85 scope themselves to one spell's ranks. Restoration :81 rank 1 "most healing your mana buys" re-derived as in round 22 (rank 1 Rejuvenation with Bloom beats every other spell for its mana). OK.
- **"Every" claims a talent covers only in part.** Moonglow "every damage spell" (client "damaging spells"); Genesis "every HoT tick"; Blood Frenzy "every builder crit" (non-periodic, Cat abilities that build points); Bear "5 rage on every crit" (in form); Nature's Grace "every non-periodic crit" (a moonkin's crits are spells); Naturalist "every kind". OK. Restoration :91 is R23-2. Balance :105 "every cast" stays a Should.
- **Rounded costs stated as exact.** Every quoted cost is the exact client value less the build's discount: Shred 42 (60 - 18), 54 and 48 at 20 and 21; Claw 40 (45 - 5); Rake 35 (40 - 5); Primal Bite 15 (20 - 5); Maul 10 by 14 (15 - 5); Shifting Power 55% of base mana (no Natural Shapeshifter in the raid build); Swiftmend 20% of base mana; Moonglow 25%; Improved Wrath half. OK.
- **Both sides of every comparison across all bands (round 23).**
  - Feral [9], Claw 40 against Shred 54 at 20 and 48 at 21: Ferocity is full by 14 and Shredding Attacks takes 20, 21, 22 on both orders. OK.
  - Feral [8], "harder at every rank": Shred 200% + flat against Claw 140% + flat. The tightest pairing is bands 28-29, Claw 2 (+39) against Shred 1 (+22), which needs a normal hit above 28. A level 28 cat with a fair weapon clears that (weapon DPS plus AP/14 at a 1.0 swing, about 30 to 35), and a very weak weapon makes it a tie. Bands 48-53 (+88 against +64) and 58-60 (+115 against +80) clear it easily. OK; see Shoulds.
  - Balance [6] and :95, Wrath against Starfire. Mana: a fraction at every band (band 20: 40 at half against 95; band 60: 120 at half against 315). Cast: 1.5 against 3.5 before Improved Starfire, 3.4 on the leveling build, 3.0 on the raid build at 37; "half" is exact at the raid 60 the rotation section means and an understatement below it. OK.
  - Feral :75, 35 energy of Rake against Shred at 42: both sides fixed from 24. OK.
  - Restoration :81, rank 7 against rank 10 at raid 60: mana 195 against 335 (42% less). With the raid set's 498 healing and 195 spell power, rank 10 is about 644 + 554 = 1198 and rank 7 about 284 + 0.85 x 554 = 755, so 63% of the heal. "Most of the heal for far less mana" holds; both sides scale alike with the talents.
  - Restoration :83 and :85 rank ladders against every band. OK.
  - Bear [7], 15 against Maul's cost: 14 at band 10, 10 by 14. OK.
  - Balance :56, "moving costs you more than it costs a mage": a mage's fillers are Scorch and Arcane Blast at 1.5 s and Frostbolt at 2.5 s; a Fireball build's 3.0 s matches plain Starfire. Defensible as written, no finding.
- **Other superlatives and absolutes.** Grep over all five pages plus every note. Dodge "the only avoidance"; Shifting Power "the only energy button" (a talent button; Thistle Tea is a consumable, as in round 22); "no cleave" (Swipe is a bear button); Wolfshead "druid-only" (Wowhead Forever tooltip: Classes: Druid); Enchanting "the one profession your raid set asks for" (both Heart trinkets, every other craft binds when equipped); Hurricane no cooldown; Tranquility 5 to 2 minutes; Wild Growth 2.5 s on 6 s; Frenzied Regeneration whole health bar at 100 rage (client and Blizzard). No new error beyond R23-1 and R23-2.
- **Gear and profession prose against the tables.** Feral :105, Restoration :105 and Bear :107's second sentence all re-read against the band rows and `bonding`. The Defender's, Brawler's and Warden's pieces in Bear bands 20 to 50 are "pickup". Living Crown, Idol of Swiftness, Mooncloth and Timbermaw pieces are "equip". OK.
- **Source links.** All eight URLs in the druid frontmatter returned 200 on this tree. The Blizzard deep dive still says what the pages cite it for (Swiftmend, Feral Charge, Faerie Fire, Thistle Tea, Tiger's Fury).

## Routed items, not re-filed

- Row-before-talent rendering (Shifting Power 10-39, Insect Swarm 20-26, Nature's Swiftness 10-29, Swiftmend 10-42, Feral "Below level 20" rows past 20), the Feral Shred row's Ferocious Bite gate, Claw against Shred at 21 to 23, Balance Eclipse Wrath on the leveling order.
- `balance.md:105` "a resisted Wrath banks no Eclipse charges" against the engine's cast-complete trigger stays open for a launch log (carried since round 13). The client proc mask is still not in the tree.
- "Biggest threat hit" and the bear threat model (ledger Round 21).

## Shoulds (wording, not findings)

- **New: restoration.md:73** "whenever he is under 75%" and "Cast it on cooldown whenever there is damage to heal" leave out the card's mana gate ("your mana is keeping pace with the fight"). Suggested: end with "as long as your mana is keeping pace".
- **New: Feral [8]** "at every rank": on a very weak weapon bands 28-29 are close to a tie. Suggested: leave it, or "Shred hits harder than Claw".
- **New: Balance [6] and :95** "half the cast": less than half below the raid 60. No change needed unless the line is touched.
- **Carried, none applied:** balance :105 "every cast"; balance :91 "always lands on a Starfire or a DoT"; balance :119 could name the Frozen Heart; feral-bear :76 and bear [3] "biggest threat hit" (routed); restoration :121 add the tauren 21 to 23; Balance [6] "With Improved Wrath" at band 10; restoration :85 "Below 20, from 21"; feral :67 "The raid build skips Furor"; Feral Bite "With Predatory Discipline" and Balance Moonfire "With Improved Moonfire and Nature's Splendor", both conditional; Bear Demoralizing Roar "204 ... at 60" on rank 1 bands; Feral and bear Faerie Fire "free" on band 18; "for good" at balance :115 and restoration :101; restoration :57 "next Healing Touch" should be "next Nature spell"; feral :77 "(55 with Wolfshead Helm)".

## Decisions

- **D1. R23-1 is CONFIRMED, not a Should.** Round 16 passed the line on its meaning. Round 23's bar is what the sentence states, and it states something false about eleven of the set's fifteen filled slots. Decided on the data.
- **D2. R23-2 is CONFIRMED, not a Should.** Round 22 filed the same shape ("every heal") for one heal and the fix's own wording now reads as a complete list. Rank 1 Rejuvenation is on the card at 43 to 59, and :81 says it keeps none.
- **D3. Rank 7 "most of the heal" stays.** 63% in the raid set; see the comparison above.
- **D4. Balance :56 stays.** See the comparison above.

## Sims requested

None new. Carried and already routed: Feral Claw against Shred at 21 to 23; Balance leveling Eclipse Wrath on and off at 36, 40 and 50.

## Sim findings

None new.

## Engine questions

None new. Carried and held by the coordinator: `BloomSpellCoeff` 0.2143 fill against client 0; Eclipse on a resisted Wrath; the bear threat model; Rend and Tear on any bleeding target.

## Sources

- **Client 1.60.1.70790** (`data/builds/1.60.1.70790/`): `talents/druid.json`, `trainables/druid.json`, `spelltips/druid.json` (all 218 rows), `spellconst/druid.json` (774, 8910, 9841, 9858, 408248, 408255, 441260), `addon-builds.json`, `addon-data.json`, `bis/druid-*.json`, `items/druid.json`, `races.json`, `spells.json`.
- **Curated:** `data/curated/apl/druid-{balance,feral,feral-bear,restoration}.json`.
- **Engine d5c708834:** `sim/core/spell.go` (level rule), `sim/druid/talents.go` (Eclipse, Natural Instinct), `sim/druid/shred.go` (0.8 refund), `sim/druid/faerie_fire.go`.
- **Reports:** `round22/druid.md`, `round16/druid.md` (the passed bear line), `round14/druid.md`, `round13/druid.md`, `ledger.md` Rounds 19 to 22.
- **Web:**
  - Blizzard, "World of Warcraft Forever class deep dives: Hunter and Druid": https://news.blizzard.com/en-us/article/24301515/world-of-warcraft-forever-class-deep-dives-hunter-and-druid
  - Wowhead Forever tooltip, Wolfshead Helm: https://nether.wowhead.com/forever/tooltip/item/8345

## Shipped orders used

- **Balance leveling:** Improved Wrath 10-14, Genesis 15-16, Moonglow 17-19, Improved Moonfire 20-21, Nature's Majesty 22-23, Nature's Reach 24-25, Nature's Splendor 26, Insect Swarm 27, Vengeance 28-32, Improved Starfire 33, Nature's Grace 34, Eclipse 35, Moonfury 36-39 and 41, Moonkin 40, Nature's Focus 42-46, Heart of the Wild 47-51, Naturalist 52-56, Reflection 57-59, Subtlety 60.
- **Balance raid:** same to 32, Improved Starfire 33-37, Nature's Grace 38, Eclipse 39, 41, 42, Moonkin 40, Moonfury 43-47, Nature's Focus 48-52, Naturalist 53-57, Subtlety 58-60.
- **Feral leveling:** Ferocity 10-14, Feral Swiftness 15-16, Feral Instinct 17, Brutal Impact 18-19, Shredding Attacks 20-22, Savage Fury 23-24, Feral Charge 25, Sharpened Claws 26-27, Predatory Strikes 28-30, Predatory Discipline 31, Blood Frenzy 32-33, Leader of the Pack 34, Natural Instinct 35-36, Rend and Tear 37-39, Shifting Power 40, Genesis 41-45, Nature's Majesty 46-47, Nature's Reach 48-49, Heart of the Wild 50-54, Naturalist 55-59, Natural Shapeshifter 60.
- **Feral raid:** Feral Instinct 17-19, Shredding Attacks 20-22, Savage Fury 23-24, Sharpened Claws 25-26, Predatory Strikes 27-29, Predatory Discipline 30, Blood Frenzy 31-32, Leader of the Pack 33, Natural Instinct 34-35, Rend and Tear 36-39 and 41, Shifting Power 40, Genesis 42-46, Nature's Majesty 47-48, Nature's Reach 49-50, Heart of the Wild 51-55, Naturalist 56-60.
- **Bear (both):** Ferocity 10-14, Furor 15, Feral Swiftness 16-17, Feral Instinct 18-19 and 22, Savage Fury 20-21, Thick Hide 23-24 and 26, Lacerate 25, Shredding Attacks 27-29, Leader of the Pack 30, Sharpened Claws 31-32, Predatory Strikes 33-35, Blood Frenzy 36-37, Natural Instinct 38-39, Primal Bite 40, Rend and Tear 41-45, Natural Reaction 46-50, Heart of the Wild 51-55, Naturalist 56-60.
- **Restoration (both):** Heart of the Wild 10-14, Reflection 15-17, Nature's Focus 18-19 and 21-23, Bloom 20, Naturalist 24-28, Gift of Nature 29 and 31-34, Nature's Swiftness 30, Tranquil Spirit 35-39, Improved Rejuvenation 40-42, Swiftmend 43, Living Spirit 44-46, Improved Tranquility 47-48, Improved Regrowth 49, Genesis 50-54, Moonglow 55-57, Nature's Majesty 58-59, Nature's Splendor 60.

## Note sweep

One line per card row note: spec, row and action, the bands it renders on, each named talent, spell, rank or value with the level the player has it (leveling / raid order where they differ), and the verdict.

- **Balance [1] Moonfire, end of fight. Bands 10-60, ranks 2 to 10.** Wrath's cast time; "best instant": Insect Swarm has no hit (27), Faerie Fire no damage (18), Hurricane is channelled (40). OK.
- **Balance [2] Wrath, end of fight. Bands 10-60.** Starfire at 20, scoped; before 20 "the last few seconds". OK.
- **Balance [3] Moonfire upkeep. Bands 10-60.** Improved Moonfire 20-21 and Nature's Splendor 26 on both orders, as a conditional. OK (Should carried).
- **Balance [4] Insect Swarm. Bands 20-60.** Talent at 27 on both orders (whole-row issue, routed); one global a refresh. OK.
- **Balance [5] Eclipse Wrath. Bands 10-60.** Eclipse 35 / 39, scoped; two charges; 0.17 s at one point (client rank 1 text); half a second at three (raid 42); "in the raid build". OK.
- **Balance [6] Wrath while behind on mana. Bands 10-60.** Innervate 40, scoped; Improved Wrath 10-14; Starfire 20, scoped; both sides of the cost and cast comparison above. OK (band 10 Should carried).
- **Balance [7] Starfire. Bands 20-60, ranks 1 to 6.** Eclipse 35 / 39, scoped. OK.
- **Feral [1] Ravage. Bands 32-60, ranks 1 to 4.** Prowl 20, Poised 32 (10 off 60). OK.
- **Feral [2] Faerie Fire. Bands 18-60.** Cat Form 20, conditional; free on a 6 s cooldown. OK (band 18-19 Should carried).
- **Feral [3] Moonfire, below 20. Bands 10-60.** Cat Form 20. OK (component row issue, routed).
- **Feral [4] Wrath, below 20. Bands 10-60.** As [3]. OK.
- **Feral [5] Shifting Power. Bands 10-60.** Talent at 40 on both orders (whole-row issue, routed); 40 energy, costs mana, 8 s. OK.
- **Feral [6] Rip. Bands 20-60, ranks 1 to 6.** 5 points, under 2 s; Rend and Tear 37 / 36, scoped. OK.
- **Feral [7] Ferocious Bite. Bands 24-60, ranks 1 to 5.** Shred live from 20; Predatory Discipline 31 / 30, conditional. OK (Should carried).
- **Feral [8] Shred. Bands 20-60, ranks 1 to 5.** "From level 22" and Shredding Attacks 20-22 on both orders; 8 to 14 more energy (48 - 40, 54 - 40); "harder at every rank" as above; the row's gate is routed. OK.
- **Feral [9] Claw. Bands 20-60, ranks 1 to 5.** 40 against 54 at 20 and 48 at 21 on both orders; "more for its energy" at 21 is in the routed sim. OK.
- **Bear [0] Cooldowns. Page only.** Mighty Rage Potion in form. OK.
- **Bear [1] Barkskin. Bands 44-60.** Trainer 44; 20%, 15 s, 60 s (22812). OK.
- **Bear [2] Frenzied Regeneration. Bands 36-60.** Trainer 36; half health, 20 rage. OK.
- **Bear [3] Primal Bite. Bands 40-60, ranks 1 to 3.** Talent 40; 6 s cooldown; "biggest threat hit" routed. Not re-filed.
- **Bear [4] Lacerate. Bands 25-60, ranks 1 to 5.** Talent 25; five stacks; 3 s; Rend and Tear 41, scoped. OK.
- **Bear [5] Faerie Fire. Bands 18-60.** Bear Form 10; free on a 6 s cooldown; carries threat. OK (band 18 Should carried).
- **Bear [6] Demoralizing Roar. Bands 10-60, ranks 1 to 5.** 204 at 60, scoped (193 at 52, 1.4 a level); the last 3 s; does not stack with Demoralizing Shout. OK (rank 1 bands Should carried).
- **Bear [7] Maul. Bands 10-60, ranks 1 to 7.** 15 rage, 10 by 14 (Ferocity 10-14); Lacerate 25, Primal Bite 40, scoped. OK.
- **Bear [8] Swipe. Bands 16-60, ranks 1 to 5.** Trainer 16; above 40 rage; 3 targets (client text). OK.
- **Restoration [1] Innervate. Bands 40-60.** Trainer 40; under 30%; +400% regeneration. OK.
- **Restoration [2] Read Leyline. Alliance Skyborne, no addon row.** 50% for 15 s; not under Innervate. OK.
- **Restoration [3] Nature's Swiftness. Bands 10-60.** Talent 30 (whole-row issue, routed); tank under 35%. OK.
- **Restoration [4] Healing Touch save. Bands 10-60.** Newest trainer rank at every band (5186 to 9889; rank 11 is a book). OK.
- **Restoration [5] Swiftmend. Bands 10-60.** Talent 43 (whole-row issue, routed); under 75%; keeps the HoT (Blizzard); 15 s; 20% of base mana, no discount in the build. OK.
- **Restoration [6] Tranquility. Bands 30-60, ranks 1 to 4.** Trainer 30; two of the party under 70%; every 2 s. OK.
- **Restoration [7] Regrowth, tank. Bands 12-60, ranks 1 to 9.** The top rank at every band; Bloom 20 (index 10), scoped. OK.
- **Restoration [8] Rejuvenation, tank. Bands 10-60.** Rank 10 at 60; Bloom 20, scoped; ranks match the rows: 2 at 20-39 (band 29's default rank 5 is overridden for all four druid races: night elf rank 3, the others rank 2), 4 at 40-43, 5 at 44-48, 6 at 49-56, 8 at 57-59. OK.
- **Restoration [9] Healing Touch filler. Bands 10-60.** Rank 2 at 20 (tauren 21-23 via the band 21 override), 3 at 30-33, 4 at 34-39, 5 at 40-45, 6 at 46-49, the top rank elsewhere from 21. Matches the rows. OK.
- **Restoration [10] Rejuvenation, party. Bands 10-60.** Under 70%; no mana-pace gate on the row. Rank 1, rank 2 at 27-29 and 40-42, rank 7 at 60. OK. Rows [11] to [14] carry no note.
