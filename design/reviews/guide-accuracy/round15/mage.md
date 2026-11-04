# Guide accuracy round 15: Mage

Audited on the round 15 audit tree, main 4600b7d82 (fourteen guide rounds, sim round 17, launch state). Client 1.60.1.70755, engine pin b2c4a2d05 (sim/enginever/version.go).

Files read line by line:
- web/src/content/guides/mage/index.md, arcane.md, fire.md and frost.md.
- Every card note and action in data/curated/apl/mage-{arcane,fire,frost}.json, priorityList and prepullActions.
- The live BiS files bis/mage-{arcane,fire,frost}.json: band 60 bare and raid rows for both factions, including the talent strings each row was ranked on.

Read first: the mage reports for rounds 1 to 14 with their fix logs, round4/mage-levels.md, the race-sims files for rounds 2 to 5, ledger.md, and the sim round 17 mage report (design/reviews/2026-11-04-sim-audit-round17/mage.md).

## What changed since round 14

Main moved from d7fc51fe7 to 4600b7d82. The mage changes are:
- **Round 14 fixes:**
  - fire.md:102 and frost.md:103: the professions paragraph names the Robe, the Frozen Heart and the goggles' Engineering 270.
  - arcane.md:72: "if it crits, that is your first Ignite".
- **Sim round 17 (S 88cacd3af), restoring round 16's wording:**
  - arcane.md:56, :72 and :84 now read "with the leveling build", replacing "before you have the Fire points".
  - The same change is in Arcane card notes 4 and 6.
- **Sim round 17, new text:**
  - arcane.md:116 adds: "At 20 and 21, before Missile Barrage, Arcane Missiles still does more damage for its mana than Blast, so lead with Missiles; from 22 you level on Blast and free Missiles."
  - New Arcane card line 8: "Arcane Missiles before you have Missile Barrage, at 20 and 21: on a long fight it does more damage for its mana than Arcane Blast or Frostbolt. From 22 Missiles waits for the free proc and Blast takes over."
  - The Missile Barrage line (7) is now gated on the talent.

addon-builds.json, talents, trainables, spelltips and the mage BiS files have no mage changes. The addon-builds.json diff touches only rogue-subtlety and warlock-affliction.

## Builds decoded

These are unchanged from round 14 (talents/mage.json, tier/column order). All five are legal at 51 points, with every tier gate and prerequisite met at the level the shipped order takes it:
- Arcane Meditation after Arcane Concentration 5.
- Arcane Wrath after Arcane Blast.
- Arcane Power after Presence of Mind.
- Heating Up after Pyroblast.
- Combustion after Critical Mass 3.
- Fingers of Frost after Ice Lance.
- Ice Barrier after Cold Snap.

The builds:
- **Arcane leveling 31/4/16:**
  - Arcane: Wand Specialization 1, Improved Channeling 4, Arcane Resilience 1, Arcane Concentration 5, Missile Barrage 2, Arcane Geometry 1, Arcane Impact 3, Arcane Blast 1, Arcane Meditation 3, Presence of Mind 1, Arcane Mind 5, Arcane Instability 3, Arcane Power 1.
  - Fire: Incineration 4.
  - Frost: Improved Frostbolt 5, Precision 3, Ice Shards 5, Frost Channeling 3.
- **Arcane raid 32/16/3:**
  - Arcane: as the leveling build, with Improved Channeling 3, Arcane Meditation 2, plus Arcane Subtlety 2 and Arcane Wrath 1.
  - Fire: Incineration 5, Accelerant 2, Ignite 5, Pyroblast 1, Improved Fire Blast 2, Heating Up 1.
  - Frost: Precision 3.
- **Fire leveling 17/31/3:**
  - Arcane: Wand Specialization 2, Improved Channeling 5, Arcane Resilience 2, Magic Absorption 1, Arcane Concentration 5, Arcane Meditation 2.
  - Fire: Incineration 5, Accelerant 1, Ignite 5, Flame Throwing 2, Burning Soul 1, Pyroblast 1, Improved Fire Blast 2, Improved Scorch 3, Heating Up 1, Master of Elements 3, Critical Mass 3, Fire Power 3, Combustion 1.
  - Frost: Precision 3.
- **Fire raid 17/31/3:** as leveling, with Arcane Resilience 3 instead of Magic Absorption, and Burning Soul 2 instead of Accelerant 1.
- **Frost 15/5/31, both strings:**
  - Arcane: Wand Specialization 2, Improved Channeling 3, Arcane Concentration 5, Missile Barrage 2, Arcane Impact 3.
  - Fire: Incineration 5.
  - Frost: Frost Warding 2, Improved Frostbolt 5, Precision 3, Ice Shards 5, Frostbite 1, Permafrost 1, Frost Channeling 3, Ice Lance 1, Ice Block 1, Shatter 3, Cold Snap 1, Fingers of Frost 2, Winter's Chill 2, Ice Barrier 1.

Point counts the prose quotes, re-added:
- **Arcane:**
  - Raid 32 = leveling 31, minus a rank of Arcane Meditation and of Improved Channeling, plus Arcane Subtlety 2 and Arcane Wrath 1.
  - Fire 16 = 5/2/5/1/2/1.
  - The 9 raid points are Ignite, Pyroblast, Improved Fire Blast and Heating Up.
  - The 8 points are Arcane Mind and Arcane Instability.
  - The 10 leveling points are Improved Frostbolt and Ice Shards.
- **Fire:**
  - Damage talents plus Flame Throwing come to 28, plus the 2-point toll = 30, plus Combustion = 31.
  - 17 Arcane points.
- **Frost:**
  - 31 Frost, of which Winter's Chill 2 is the toll to 30 points above Ice Barrier.
  - 15 Arcane, 5 Fire.

## Leveling orders against the prose (addon-builds.json, index 0 = level 10)

- **Arcane leveling:**
  - Wand Specialization at 10, Improved Channeling 11 to 14, Arcane Concentration 15 to 19 (arcane.md:110).
  - Arcane Blast 20, Arcane Resilience 21, Missile Barrage 22 and 23, Arcane Power 40, every point Arcane to 40 (arcane.md:116).
  - Incineration 41 to 44, Improved Frostbolt 45 to 49, Precision 50 to 52, Ice Shards 53 to 57, Frost Channeling 58 to 60. That fits "Incineration, then Improved Frostbolt, Precision, Ice Shards and Frost Channeling", and Improved Frostbolt and Ice Shards are both in "by 57" (arcane.md:59).
- **Fire leveling:**
  - Incineration 10 to 14, Accelerant 15, Ignite 16 to 19 (fire.md:110).
  - Burning Soul 24, Improved Scorch 27 to 29, Heating Up 30 (fire.md:116).
- **Frost:**
  - Improved Frostbolt 10 to 14, Frost Warding 15 and 16, Precision 17 to 19 (frost.md:111).
  - Ice Shards 20 to 24, Frostbite 25, Permafrost 26, Frost Channeling 27 to 29, Ice Lance 30, Ice Block 31.
  - Shatter 32 to 34, Fingers of Frost 36 and 37, Ice Barrier 40, every point Frost to 40 (frost.md:117).
- Trainer levels (trainables/mage.json), all as the pages state:
  - Frostbolt 4/8/14/20, Fireball 1/6/12/18.
  - Conjure Water 4, Fire Blast 6, Arcane Missiles 8, Frost Nova 10, Arcane Explosion 14, Flamestrike 16.
  - Scorch 22, Cone of Cold 26, Ice Armor 30, Mage Armor 34, Frostfire Bolt 40/50/60, Teleport: Dalaran 50.

## Findings

| id | status | kind | file:line | claim | evidence | fix |
|---|---|---|---|---|---|---|
| MAGE-R15-1 | NEEDS-SIM (low) | leveling | arcane.md:116; mage-arcane.json line 8 note | "from 22 you level on Blast and free Missiles" / "From 22 Missiles waits for the free proc and Blast takes over." | Sim round 17 A1 measured Missiles-first against Blast at levels 20 and 21 (no Missile Barrage) and at 25, 30 and 40 (Missile Barrage 2/2). At 25 Missiles-first lost by 59.1 against 59.7, which is inside the run's spread. Level 22 has Missile Barrage 1/2 only: a 20% Blast proc, with rank 1 Blast at 32 to 38 damage for 15% of 415 base mana. That level was not measured. The sentence is plausible but unmeasured at the level it names. | None until measured. If Missiles-first still wins at 22, change "from 22" to "from 23, once Missile Barrage has both ranks" on the page and in the card note. |

Counts: CONFIRMED 0, LIKELY 0, NEEDS-SIM 1, NEEDS-OWNER 0.

## New text this round, checked

- **arcane.md:116, Missiles at 20 and 21.**
  - At 20 and 21 you have Arcane Missiles rank 2 (5144: 140 mana, 34 a second for 4 seconds, 136 damage, so 0.97 damage per mana) and rank 1 Arcane Blast (talent text 34 to 38; cost 15% of base mana, 371 at 20 from gametables/basemp.txt, so about 56 mana and 0.64 damage per mana for the first Blast, then double).
  - So Missiles does more damage for its mana than Blast, as the sentence says.
  - The card note's wider claim (also more than Frostbolt) holds with Arcane Concentration 5/5 in by 19. Each of the four missiles can roll Clearcasting, so about a third of Missiles casts are followed by a free spell, against 10% for a Frostbolt. Rank 4 Frostbolt is 65 mana for 70.
  - Sim round 17 measured Missiles-first at 48.0 against 33.1 to 39.0.
  - "Before Missile Barrage" matches the order (Missile Barrage at 22).
  - The sentence does not contradict Levels 1 to 19 ("From 8, open with Frostbolt ... channel Missiles as it closes").
- **"With the leveling build"** (arcane.md:56, :72, :84; card notes 4 and 6):
  - True. The leveling build has no Fire talent beyond Incineration 4, so no Pyroblast, and its Frostbolt is 2.5 seconds at 60 (Improved Frostbolt 5/5 from 49).
  - The prepull matches. Pyroblast r8 at -6 s is gated on Heating Up (400625), which only the raid build has. Frostbolt r10 at -2.5 s goes out otherwise.
  - Card line 5 is gated on Heating Up and line 6 on not knowing it, both on Presence of Mind being up.
- **arcane.md:72, "if it crits, that is your first Ignite":** right. Pyroblast has no guaranteed crit, and Ignite triggers on Fire crits.
- **fire.md:102 and frost.md:103, professions:**
  - Robe of the Archmage (14152) is the chest in every band 60 row.
  - Frozen Heart of the Mountain (249469) is in every band 60 raid row of Fire and Frost, on both factions.
  - Spellpower Goggles Xtreme Plus (15999) is the head in all four Fire and Frost raid rows.
  - "Your raid set" is right: the Fire bare rows carry no Heart.

## Round 14's other state, re-checked

- **Talent text** (talents/mage.json), every value the pages quote:
  - Arcane Blast 12% Intellect a stack, +100% cost, 4 stacks, 18 s.
  - Missile Barrage 40/20%.
  - Arcane Concentration 10%.
  - Arcane Meditation 17/33/50%.
  - Arcane Power 30%/30%, 3 min.
  - Arcane Mind 10% Intellect, +100% Arcane crit bonus.
  - Arcane Instability 3%/3%.
  - Arcane Wrath 3.5 s minus Accelerant 0.2, 6 s cooldown, 5% of maximum mana.
  - Arcane Subtlety 30% Arcane threat.
  - Improved Channeling 80% Missiles at 4 ranks.
  - Arcane Geometry 3 yards.
  - Incineration 10% on Fire Blast, Ice Lance, Arcane Blast and Scorch.
  - Accelerant 0.1 s a rank on Fireball, Frostfire Bolt and Arcane Wrath.
  - Ignite 40% over 4 s.
  - Burning Soul 23%/10% at one rank and 47%/20% at two.
  - Improved Scorch 3% a stack, 5 stacks, 30 s.
  - Heating Up 25% a stack, 3 stacks, Fire Blast hits add a stack.
  - Improved Fire Blast 2 s and 15 yards.
  - Master of Elements 30%.
  - Critical Mass 6%.
  - Fire Power 3% an effect to 12%, mixed schools excluded.
  - Combustion 10% a hit until 3 non-periodic crits, 3 min.
  - Flame Throwing 6 yards.
  - Frost Warding 30%.
  - Improved Frostbolt 0.5 s.
  - Precision 6%.
  - Ice Shards +100% bonus.
  - Frost Channeling 15%/30%.
  - Shatter 50%.
  - Ice Lance +300% against Frozen, 8 s.
  - Fingers of Frost 15%, 2 spells, 15 s.
  - Winter's Chill 20% a rank (40% at 2), 5 stacks double Cone of Cold and add 35% of spell damage.
  - Ice Barrier 15% plus 5%, no pushback, 30 s cooldown.
  - Ice Block 10 s, 5 min.
  - Arcane Impact 6%.
- **Spell values** (spelltips/mage.json):
  - Scorch r7 is 150 mana and 1.5 s.
  - Flamestrike r6 is 990 mana, a 3 s cast, an 8 s burn.
  - Arcane Explosion r6 is 390 mana, 10 yards.
  - Blizzard r6 is 1,400 mana, or 1,190 with Frost Channeling.
  - Pyroblast r8 is a 6 s cast.
  - Frostfire Bolt r3 is 370 mana and 3 s, against Frostbolt r11 at 290 and Scorch at 150.
  - Frostfire Bolt r1 at 40 is 205 mana, against Frostbolt r7 at 160 ("still the cheaper slow").
  - Evocation is 8 min.
  - Mana Ruby restores 1,100.
- **Card notes** (all three APL files): each matches its action and the client values above. The Missile Barrage line now checks the talent (sim round 17 A1).
- **Race lines:**
  - Every racial value is as round 14 listed it: Eureka!, Touch of the Grave 30% on the mage row, Blood Fury, Berserking, Ruthless, The Human Spirit, Expansive Mind, Wind Blessed, Read Leyline.
  - Every race order still matches round3/race-sims.md, the only measured mage race run, and recommendedRaces.
  - Skyborne is named only as the paid option.
- **Stat prose:** it explains and never ranks. Each reason holds against the band 60 raid scale factors as round 14 derived them (bis files unchanged).
- **Gear and consumables:** no prose names gear that contradicts the live rows on either faction (see the professions checks above).
- **Source links:** all seven fetched again this round, and all returned HTTP 200:
  - the Deep Dive recap
  - Wowhead Skyborne First Look
  - Power Up Gaming
  - talentsforever.com/data.json
  - /sim/weights and /sim/drops
  - /planner?class=mage
- **Voice:**
  - No em or en dashes.
  - No site, sim, APL or band words in body prose.
  - No quoted nightly numbers.

## SHOULDs (wording, not findings)

- **arcane.md:74:** "On the way to 60 ... Blast to 3 stacks, Frostbolt as the filler" now sits beside the new advice to lead with Missiles at 20 and 21. Suggested: "On the way to 60, from 22, there is no Arcane Wrath and no Pyroblast: ..."
- **Card line 8:** "at 20 and 21 ... From 22" describes the leveling build. The raid build's order takes Missile Barrage at 23 and 24, after Arcane Subtlety at 21 and 22. Harmless at 60.
- **Carried from round 14:**
  - arcane.md:102 professions reasons (the real reasons are the Robe, Frozen Heart and Brilliant Wand).
  - fire.md:94 "Greater Arcane Elixir is your elixir" while the table also lists Elixir of Fire Power.
  - arcane.md:72 Missile Barrage "before the fight starts".
  - frost.md:71 "Recast Ice Barrier the moment it breaks" (30 s cooldown).
  - fire.md:116 Burning Soul "keeps a mob's hits from pushing back every cast" (23%).
  - index.md:32 is missing Accelerant from the new-talent list.
- **frost.md:63 (sim round 17 R2):** Winter's Chill's 40% applies to the Fingers of Frost Ice Lance too. This is an omission, not an error.

## Decisions

- **D1.** The new Missiles sentence is not a finding. The per-mana claim is true on the client values at 20, and the sim measured the play. Only the "from 22" boundary is unmeasured (R15-1, NEEDS-SIM).
- **D2.** Round 14 D1 stands. Arcane's professions line is a SHOULD, because no Arcane row holds an Engineering item and the Tailoring and Enchanting pair is right.
- **D3.** The Frost BiS rows being ranked on an old talent string is a data-lane item, not a guide finding. No Frost prose depends on the Arcane Impact or Accelerant points.

## Sims requested

- **Arcane, level 22, leveling build** (Missile Barrage 1/2, ladder gear, bare, seeds 1 and 50001): the APL with line 8 (Missiles while Missile Barrage is not known) extended to "Missile Barrage below 2 ranks", against the shipped APL. Repeat at 23. This settles R15-1.
- **BiS (data lane, post-launch re-rank):**
  - **New:** Frost's band 60 rows (bare and raid, both factions) were ranked on 230005200000000000-51020000000000000-253510130101301221. That is the round 2 Frost build: 12/8/31 with Accelerant 1 and Flame Throwing 2 and no Arcane Impact. The shipped build is 15/5/31 (Arcane Impact 3, Incineration 5 only). Re-rank on 230005203/5/253510130101301221.
  - **Carried from round 14:**
    - Fire and Frost raid sets need three professions (Engineering 270 goggles, the Tailoring Robe and the Enchanting Heart).
    - Frozen Heart of the Mountain is BiS for Arcane and Fire.
    - Thunderbrew's Boot Flask is an Arcane Alliance trinket.
    - The Arcane raid rows are on 32/3/16, not 32/16/3.
- **Race re-runs at 60 on the shipped builds (carried):**
  - Arcane 32/16/3: Gnome, Human, Undead, Orc, Troll.
  - Frost 15/5/31: Gnome, Human, Skyborne, Undead, Troll, Orc.
- **Carried:**
  - the Arcane raid pre-pull on seeds 1 and 50001
  - the Presence of Mind target
  - Arcane Wrath at 3 stacks against 4
  - the gem-timing question (sim round 17 X2)
  - S1, S2, SIM-R2-1, SIM-R2-3, SIM-5, SIM-6 and SIM-2

## Sim findings

- None new. SIM-R7-1 is carried: Arcane Subtlety's 15 resistance cut is not modeled.

## Sources

- data/builds/1.60.1.70755:
  - talents/mage.json: every rank text and prerequisite.
  - trainables/mage.json.
  - spelltips/mage.json: 5144, 400574, 7322, 10207, 10216, 10202, 10187, 18809, 1237313, 401502, 8408, 10181, 25304, 12051, 10054, 13033, 11958, 1240047.
  - gametables/basemp.txt: mage at 20, 21 and 22.
  - addon-builds.json: mage-arcane, mage-fire, mage-frost.
  - bis/mage-{arcane,fire,frost}.json: band 60 rows and their talent strings.
  - raw/wowhead/items.json.
- data/curated/apl/mage-{arcane,fire,frost}.json, as diffed against d7fc51fe7.
- Engine b2c4a2d05: sim/mage/arcane_blast.go:180-195 (the stack is added whether or not the Blast lands).
- design/reviews/2026-11-04-sim-audit-round17/mage.md (A1, A2, S3, R2, X2).
- design/reviews/guide-accuracy/round1 to round14/mage.md, round2 to round5 race-sims.md, ledger.md.
- Source links fetched 2026-11-04, all HTTP 200:
  - Blizzard, Deep Dive panel recap: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
  - Wowhead, Skyborne First Look: https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
  - Power Up Gaming, Forever talent changes vs Classic: https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/
  - Talents Forever data export: https://talentsforever.com/data.json
  - https://foreversixty.gg/sim/weights, https://foreversixty.gg/sim/drops, https://foreversixty.gg/planner?class=mage
