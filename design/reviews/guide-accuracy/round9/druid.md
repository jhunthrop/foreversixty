# Guide accuracy round 9: Druid

Auditor: druid persona. Tree: gr9/audit at 645caa983 (land/guides-r8, all eight rounds applied; READ-ONLY), client 1.60.1.70755, engine pin d1d9f2775 (sim/enginever/version.go).

Specs audited: class page (`index.md`), Balance (`balance.md`), Feral cat (`feral.md`), Feral bear (`feral-bear.md`), Restoration (`restoration.md`). I read the round 1 to 8 druid reports with their fix logs, round4/druid-levels.md and ledger.md before starting. Between round 8's audit head (09b98d9e4) and 645caa983, only one druid line changed: the round 8 fix at balance.md:105. No commit touched addon-builds.json, talents/druid.json, trainables/druid.json, spelltips/druid.json or the druid APLs. I still re-read every line of all five pages and every rendered card note from scratch.

## Decoded builds (unchanged since round 1, all legal, 51 points)

- Balance `5232220115101151/0/555301`, raid `5232220115501351/0/505003`.
- Feral `01/50212032123012123001/535033`, raid `05002/50230032023011125001/05503`.
- Bear `0/5523033202310212551/051`.
- Restoration `05312001/0/55530051531132`.

The shipped leveling orders (addon-builds.json, unchanged) are as decoded in round5/druid.md. Every Levels 1 to 19 paragraph and every 20 to 60 talent claim still matches them. The feral 60 respec sentence (feral.md:55) still matches the decoded diff.

## Findings

No CONFIRMED findings this round.

| id | status | kind | file:line | claim | evidence | action |
|---|---|---|---|---|---|---|
| DRUID-R9-1 | NEEDS-SIM | rotation | feral.md:73; druid-feral.json priorityList[6] note | "Rip at 5 points does far more over its duration than a Bite, so it always wins." | The 70755 push raised both finishers. Rip r6 tick base goes from 15 to 19 (spellconst 9896), and Ferocious Bite r5 goes from 70/250 to 87/310 (22829). Engine d1d9f2775 still has the 70753 values. A rough 60 hand count (1200 AP, Genesis 5, Rend and Tear 5, Predatory Discipline with about 50 spare energy) puts a 5-point Bite near the size of a full Rip before crits. The ordering itself, Rip first and Bite only while Rip runs, is almost certainly still right, because Rip costs 5 less energy and keeps Rend and Tear on. But "far more" now needs a measurement. | No prose change until Sims requested 2 runs. If Rip still clearly wins, keep the line. If the two are close, write: "Rip at 5 points comes first, because it is the bleed that keeps Rend and Tear on." |

Checked and right (no finding):

**Crit-damage and Eureka! lines (round 9 focus)**
- Vengeance (balance.md:50, :69): "Increases the critical strike damage bonus of your Arcane and Nature spells by 100%" (talents/druid.json). A spell crit's bonus is +50%, so 100% more of that bonus doubles it, to 200% of a normal hit. "Doubled" is exact.
- Natural Instinct ("bigger crits", feral.md:60, feral-bear.md:61): +20% to the melee abilities' crit bonus, which gives 220% of a normal hit. No page states a number, so nothing is wrong.
- No druid page mentions Eureka!. No druid race can take it.

**Spells and passives (70755 text, spelltips/druid.json and trainables/druid.json)**
- Omen of Clarity 16864 is baseline at 20 and "not consumed by Wrath or by spells or abilities that cost no resources". So the cat's free Faerie Fire never eats a proc, and feral.md:77 ("nearly always lands on a Shred") holds.
- Poised 1328829 is learned at 32: "While at maximum Energy, and out of combat, reduce the Energy cost of your next ability by 10."
- Shifting Power 1322605: 55% base mana for 40 energy, 8 s cooldown, cheaper with shapeshift cost cuts.
- Thick Hide talent: the armor is "further increased by multipliers from those forms". The baseline Thick Hide passive 16929 is learned at 1. So feral-bear.md:35 and :53 hold.
- Ferocious Bite rank 1 is learned at 24. Rake trains at 24, Claw, Rip, Shred and Prowl at 20, Ravage at 32, Pounce at 36 ("Must be prowling and behind the target"), Cower at 28, Frenzied Regeneration at 36, Barkskin at 44, Starfire at 20, Hurricane at 40 with no cooldown, Moonkin Form and Innervate at 40, Tranquility at 30, and Gift of the Wild at 50 and 60.
- Wrath tops out at rank 8 (54). Wrath 120 mana becomes 60 with Improved Wrath, against Starfire's 315 or 340, and 1.5 s against 3.0 s, so balance.md:95 holds.
- Enrage: 10 rage up front plus 20 over 10 s. It costs 16% base armor in Dire Bear Form and puts you in combat. Demoralizing Roar r5 is 193 at its spell level 52 for 30 s, which is 204 at 60. Swipe hits 3. Challenging Roar lasts 6 s. Frenzied Regeneration converts 1% health a rage. Barkskin cuts physical damage by 20% for 15 s.
- Innervate: 400% and 100% while casting, 20 s. Swiftmend: 20% base mana, 15 s. Wild Growth: 3.0 s, 6 s cooldown, 1050 mana, heals the target's party over 7 s. Tranquility: party within 20 yards.
- Rejuvenation, Healing Touch and Starfire rank 11/11/7 are book ranks (spellranks.json `"book": true`). So "your last three ranks (8, 9 and 10)" on restoration.md:80 still means the trainer ranks.

**Talents (every named one, re-read)**
- Every rank text quoted on the pages matches: Improved Wrath, Genesis, Moonglow, Improved Moonfire (damage and crit of Moonfire), Nature's Majesty (spells and melee), Nature's Reach, Nature's Splendor, Eclipse (2 Starfires, 4 charges), Nature's Grace (non-periodic spell crits), Moonfury, Moonkin Form, Ferocity, Furor, Feral Swiftness, Feral Instinct, Brutal Impact, Shredding Attacks, Savage Fury, Sharpened Claws, Predatory Strikes, Lacerate, Predatory Discipline, Blood Frenzy, Leader of the Pack, Rend and Tear (melee abilities), Natural Reaction, Primal Bite, Nature's Focus, Heart of the Wild, Naturalist, Reflection, Natural Shapeshifter, Subtlety (30% threat), Gift of Nature, Bloom, Tranquil Spirit, Improved Rejuvenation, Swiftmend, Nature's Swiftness, Living Spirit, Improved Tranquility (60% cooldown, no threat), Improved Regrowth and Wild Growth.

**Round 8 edit re-checked**
- balance.md:105 "Nature damage feeds Wrath, Insect Swarm and Hurricane, never Starfire or Moonfire". By school_mask, Wrath, Insect Swarm and Hurricane are 8 (Nature), and Starfire and Moonfire are 64 (Arcane). Right.

**Card notes**
- I re-read every rendered druid note (priorityList notes in the four APLs). None is false on 70755.
- The Claw note's "does more for its energy" at 20 and 21 still holds on the pushed numbers. Rank 1 Claw is 140% plus 27 for 40 energy, and rank 1 Shred is 200% plus 22 for 54 or 48 energy. For a level 20 paw hit around 25 to 35, Claw leads per energy at 20 and about ties at 21.

**Sources**
- Every frontmatter link returns 200 on 2026-11-04 to a browser user agent: the Blizzard deep dive, the Blizzard panel recap, Wowhead Skyborne First Look, Talents Forever data.json, /sim/weights, /sim/drops, /bis/druid/feral-bear and /bis/druid/restoration. Each says what its label says, and each follows the round 4 source-link rule.
- No em or en dashes. No site, sim, APL, preset or band words in prose. No quoted nightly numbers. No stat ranking inside Stat priority.

## Shoulds (wording, not findings)

- balance.md:91 "it always lands on a Starfire or a DoT": a Faerie Fire in Moonkin Form (:97) or a Hurricane (:99) costs mana and takes the proc. "Almost always" would be exact. This is carried from round 8.
- feral.md:73 "so do not pool energy for it": with Predatory Discipline, spare energy adds Bite damage and is not spent. That is an argument for holding energy, not against it. See Sims requested 3. The advice matches the rotation, so this is wording until it is measured.

## Decisions

1. **The 70755 push to cat abilities and Starfire is a sim finding, not a guide finding.** No druid sentence quotes a number the push moved. Every prose comparison that rests on those numbers (Rake not worth its energy, Shred over Claw, Wrath plus two Starfires over plain Starfires) still holds on the pushed values by hand count. The one that has narrowed, Rip "far more" than a Bite, is filed NEEDS-SIM (R9-1).
2. **Rounds 1 to 8 decisions stand:** Swiftmend keeps the HoT, Gift of Nature misses Tranquility, Moonkin casts Rebirth, Innervate and Remove Curse, Pounce needs you behind the target, a failed Shred refunds 80%, and Restoration's Improved Moonfire point is named as dead with Improved Regrowth as the option.

## Sims requested

1. **Engine catch-up first (blocks 2 and 3):** carry the 70755 druid push into the engine (see Sim findings S13). Then re-run druid-feral and druid-balance at band 60, raid and bare presets, seeds 1 and 50001, and the feral and balance ladders.
2. **druid-feral, Rip against Bite at 5 points:** on the caught-up engine, compare the shipped list with a variant that Bites at 5 points whenever Rip has 6 s or more left (current), against one that also Bites in place of a Rip refresh. Report the DPS of each and the share each finisher takes. This settles R9-1.
3. **druid-feral, Bite energy pooling under Predatory Discipline:** the shipped list, against a variant that holds the 5-point Bite until energy is 80 or more (Shifting Power line unchanged). It settles the "do not pool" SHOULD.
4. **Carried from round 4 (S10):** feral leveling, the switch from Claw to Shred. Gate the Shred line on Shredding Attacks 3/3 at 22 and compare it with the Ferocious Bite rank 5 gate, ladder levels 20 to 25, re-run on the pushed numbers.
5. **Carried from round 3:** druid-restoration with Improved Moonfire 1 moved into Improved Regrowth 1 (`05302001/0/555300515311321`), heal profile, band 60 raid, with a ladder-first row that keeps Improved Regrowth.

## Sim findings

- **S13 (new, large): engine d1d9f2775 misses the 70755 druid damage push.** The site's 70755 spellconst/druid.json and the live Wowhead Forever tooltips agree. The engine still carries the 70753 numbers:

| Spell | 70753 / engine | 70755 / Wowhead | Engine line |
|---|---|---|---|
| Shred | 155% weapon damage | 200% ("causing 200% damage plus 80", 9830) | `shredWeaponDamageMultiplier = 1.55` |
| Claw | 110% | 140% (9850) | `clawWeaponDamageMultiplier = 1.10` |
| Ravage | 350% | 435% (9867) | `ravageDamageMultiplier = 3.5` |
| Starfire r6 | 337 base, 2.3 a level | 471 base, 3.2 a level ("451 to 529" at 60, 9876) | starfire.go table |
| Starfire r7 | 381 base | 533 base | starfire.go table |
| Rip ticks | r1 3 to r6 15 | r1 4 to r6 19 | `RipBaseDamage` |
| Rake r4 | 61 plus 34 a tick | 76 plus 42 a tick | `RakeBaseDamage` |
| Ferocious Bite (all ranks) | r5 70 / 250 per energy, r6 82 / 270 | r5 87 / 310, r6 102 / 330, r1 17 / 70 | `FerociousBiteBaseDamage`, `ferociousBiteDamagePerEnergy` |

  Because constants_auto_gen.go is labelled "from build 1.60.1.70755" but holds 70753 values, its generator does not layer the hotfix push. Every feral and balance number on the site is low: the stat weights, BiS, dps-per-point and race sims. The bear's Lacerate and Primal Bite did not change. The fix is to regenerate the constants with the push layered and replace the four hardcoded multipliers. After that, re-run the weights and BiS at the refresh, and re-check the measured Rake, Bite and Claw rotation choices.
- **S1 (carried, still open):** `rendAndTearDamagePerRank = 0.02` against the client's 3% a rank.
- **S7, S8, S9, S11, S12 and SIM-D1 (carried):** Nature's Reach melee hit for the feral raid re-run, Feral Charge (Cat) unmodeled, the Pounce behind-target check, expertise in the bear weights, bear Defense noise, and the feral leveling Nature's Focus and Subtlety points.

## Sources

- Client 1.60.1.70755: talents/druid.json, trainables/druid.json, spelltips/druid.json (Omen of Clarity 16864, Poised 1328829, Shifting Power 1322605, Thick Hide 16929, Shred 9830, Claw 9850, Ravage 9867, Rake 9904, Rip 9896, Ferocious Bite 1328618, Starfire 25298, Wrath 9912, Enrage 5229, Demoralizing Roar 9898, Innervate 29166, Swiftmend 18562, Wild Growth 1238215, Tranquility 9863), spellconst/druid.json diffed against 1.60.1.70753/spellconst/druid.json, spellranks.json (book ranks), addon-builds.json; APLs data/curated/apl/druid-*.json.
- Engine d1d9f2775: sim/druid/shred.go:27, claw.go:22, ravage.go:22, starfire.go:17-26, rip.go, ferocious_bite.go, constants_auto_gen.go (RakeBaseDamage, RipBaseDamage, FerociousBiteBaseDamage).
- Wowhead Forever tooltips (fetched 2026-11-04): https://nether.wowhead.com/forever/tooltip/spell/9830 (Shred 200% plus 80), https://nether.wowhead.com/forever/tooltip/spell/9850 (Claw 140% plus 115), https://nether.wowhead.com/forever/tooltip/spell/9867 (Ravage 435% plus 343), https://nether.wowhead.com/forever/tooltip/spell/9876 (Starfire 451 to 529).
- Blizzard, "World of Warcraft: Forever Class Deep Dives: Hunter and Druid", https://news.blizzard.com/en-us/article/24301515/world-of-warcraft-forever-class-deep-dives-hunter-and-druid (200).
- Blizzard, "World of Warcraft: Forever Deep Dive Panel Recap", https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap (200).
- Wowhead, "Skyborne First Look", https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829 (200).
- Talents Forever, https://talentsforever.com/data.json (200).
- Forever Sixty: https://foreversixty.gg/sim/weights, https://foreversixty.gg/sim/drops, https://foreversixty.gg/bis/druid/feral-bear, https://foreversixty.gg/bis/druid/restoration (all 200).
