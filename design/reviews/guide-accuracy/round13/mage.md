# Guide accuracy round 13: Mage

Audited on the round 13 audit tree, main 604b78a80: twelve guide rounds, sim round 15, guide-header-build. Client 1.60.1.70755, engine pin 9f7f1f6e4 (sim/enginever/version.go).

Files read line by line:
- web/src/content/guides/mage/index.md, arcane.md, fire.md and frost.md.
- Every note and action in data/curated/apl/mage-{arcane,fire,frost}.json, priorityList and prepullActions.

Read first: the mage reports for rounds 1 to 12 with their fix logs, round4/mage-levels.md, the race-sims files and ledger.md through round 12.

## What changed since round 12

Main moved from a8b147a6f to 604b78a80 with one commit, the round 12 landing. The only mage changes are in arcane.md:
- :37, the Frostbolt slow line
- :59, Improved Frostbolt and Ice Shards "by 57"
- :60, Arcane Wrath "3.3 second ... with Accelerant"
- :74, the leveling loop "no Arcane Wrath and no Pyroblast"
- :82, "The leveling build's Frost Channeling"

No commit touches mage APL files, addon-builds.json, talents, spelltips or bis/mage-*.json.

## Builds decoded (talents/mage.json, tier/column order; all five legal, 51 points)

- **Arcane leveling, 31/4/16.** Wand Specialization 1, Improved Channeling 4, Arcane Resilience 1, Arcane Concentration 5, Missile Barrage 2, Arcane Geometry 1, Arcane Impact 3, Arcane Blast 1, Arcane Meditation 3, Presence of Mind 1, Arcane Mind 5, Arcane Instability 3, Arcane Power 1. Incineration 4. Improved Frostbolt 5, Precision 3, Ice Shards 5, Frost Channeling 3.
- **Arcane raid, 32/16/3.** Wand Specialization 1, Improved Channeling 3, Arcane Resilience 1, Arcane Subtlety 2, Arcane Concentration 5, Missile Barrage 2, Arcane Geometry 1, Arcane Impact 3, Arcane Blast 1, Arcane Meditation 2, Arcane Wrath 1, Presence of Mind 1, Arcane Mind 5, Arcane Instability 3, Arcane Power 1. Incineration 5, Accelerant 2, Ignite 5, Pyroblast 1, Improved Fire Blast 2, Heating Up 1. Precision 3.
- **Fire leveling, 17/31/3.** Wand Specialization 2, Improved Channeling 5, Arcane Resilience 2, Magic Absorption 1, Arcane Concentration 5, Arcane Meditation 2. Incineration 5, Accelerant 1, Ignite 5, Flame Throwing 2, Burning Soul 1, Pyroblast 1, Improved Fire Blast 2, Improved Scorch 3, Heating Up 1, Master of Elements 3, Critical Mass 3, Fire Power 3, Combustion 1. Precision 3.
- **Fire raid, 17/31/3.** As leveling, except Arcane Resilience 3 and no Magic Absorption, and Burning Soul 2 with no Accelerant.
- **Frost, both builds, 15/5/31.** Wand Specialization 2, Improved Channeling 3, Arcane Concentration 5, Missile Barrage 2, Arcane Impact 3. Incineration 5. Frost Warding 2, Improved Frostbolt 5, Precision 3, Ice Shards 5, Frostbite 1, Permafrost 1, Frost Channeling 3, Ice Lance 1, Ice Block 1, Shatter 3, Cold Snap 1, Fingers of Frost 2, Winter's Chill 2, Ice Barrier 1.

Shipped orders (addon-builds.json, index 0 = level 10), re-decoded from [tree, tier, column]:
- **Arcane leveling:**
  - Wand Specialization 10, Improved Channeling 11-14, Arcane Concentration 15-19.
  - Arcane Blast 20, Arcane Resilience 21, Missile Barrage 22-23, Arcane Geometry 24, Arcane Impact 25-27.
  - Arcane Meditation 28-30, Presence of Mind 31, Arcane Mind 32-36, Arcane Instability 37-39, Arcane Power 40.
  - Incineration 41-44, Improved Frostbolt 45-49, Precision 50-52, Ice Shards 53-57, Frost Channeling 58-60.
- **Fire leveling:**
  - Incineration 10-14, Accelerant 15, Ignite 16-19 and 21, Pyroblast 20.
  - Flame Throwing 22-23, Burning Soul 24, Improved Fire Blast 25-26, Improved Scorch 27-29, Heating Up 30.
  - Master of Elements 31-33, Critical Mass 34-36, Fire Power 37-39, Combustion 40.
  - Arcane 41-57: Wand Specialization, Improved Channeling, Arcane Resilience, Magic Absorption, Arcane Concentration and Arcane Meditation. Then Precision 58-60.
- **Frost:**
  - Improved Frostbolt 10-14, Frost Warding 15-16, Precision 17-19, Ice Shards 20-24.
  - Frostbite 25, Permafrost 26, Frost Channeling 27-29, Ice Lance 30, Ice Block 31, Shatter 32-34.
  - Cold Snap 35, Fingers of Frost 36-37, Winter's Chill 38-39, Ice Barrier 40.
  - Arcane 41-55, then Incineration 56-60.

The Levels 1 to 19 and Levels 20 to 60 prose agree with these orders level for level. They do not contradict each other, and both read right in that order.

## Findings

None confirmed.

| id | status | kind | file:line | claim | evidence | fix |
|---|---|---|---|---|---|---|
| (none) | | | | | | |

Counts: CONFIRMED 0, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

## Round 12's own fixes, re-checked

- **arcane.md:37** "On the way to 60, the Frostbolt slow keeps adds off you while you blast". This is true from level 4. It no longer credits Frost Channeling before 58.
- **arcane.md:59** "once both are in, by 57". Improved Frostbolt comes at 45-49 and Ice Shards at 53-57. Frostbolt r10 (10181) is a 3000 ms cast, so it is 2.5 s with 5/5 Improved Frostbolt. Ice Shards 5/5 doubles the crit bonus.
- **arcane.md:60** "a 3.3 second Arcane bolt with Accelerant". Arcane Wrath r4 (1326082) is 3500 ms, and the raid build has Accelerant 2/5 at 0.1 s per rank.
- **arcane.md:74** "On the way to 60 there is no Arcane Wrath and no Pyroblast". The leveling string has neither talent, and its only Fire points are Incineration 4. The sentences that follow (stacks do not refresh, Missiles on Missile Barrage, Clearcasting) hold for both builds.
- **arcane.md:82** "The leveling build's Frost Channeling cuts 30% of the threat from your Frostbolts". Rank 3 cuts Frost threat by 30%. The raid build's Arcane Subtlety 2/2 cuts Arcane threat by 30%. The raid build has no Burning Soul, so "Your Fire spells get no cut" is true.

## No findings in these areas (re-derived this round)

- **Spell values** (spelltips/mage.json, at 60):
  - Arcane Blast r5: 1.5 s cast, 15% of base mana, +100% cost per stack, 4 stacks for 18 s from the first. A 4-stack Blast costs five times a fresh one.
  - Arcane Wrath r4: 3.5 s cast, 6 s cooldown.
  - Pyroblast r8: 6.0 s cast. Three Heating Up stacks at 25% each make it 1.5 s.
  - Fire Blast r7: 340 mana, 8 s cooldown, 20 yd.
  - Scorch r7: 150 mana, 1.5 s. Fireball r11: 3.5 s. Two Scorches fit inside one Fireball.
  - Frostbolt ranks 5 to 11: 3.0 s cast. Ranks 1 to 4 are shorter (1.5, 1.8, 2.2 and 2.6 s), so "0.5 sec faster" at 10-14 holds.
  - Frostfire Bolt r1, r2, r3: 205, 285 and 370 mana, 3.0 s, with a Frostfire DoT. r3 does 292 against Frostbolt r10's 386 and costs 370 against 260. "Hits for less, costs more" holds, and so does Fire's "Frostbolt is still the cheaper slow" (FFB r1 205 against Frostbolt r7 160, and so on up the ranks).
  - Ice Lance r6: instant, 8 s cooldown, +300% against Frozen targets.
  - Ice Barrier r4: 30 s cooldown, 15% for 1 min plus 5% while shielded, no pushback.
  - Ice Block: 10 s, 5 min. Cold Snap: 10 min.
  - Evocation: 8 min. Arcane Power, Presence of Mind and Combustion: 3 min.
  - Arcane Explosion r6: 390 mana, 10 yd. Flamestrike r6: 990 mana, 3 s cast, 8 s burn.
  - Blizzard r6: 1,400 mana, 8 s channel. With Frost Channeling that is 1,190 ("about 1,200").
  - Cone of Cold r5: 10 s cooldown.
  - Mana Ruby: 1,100. Mana Citrine: 850. Mage Armor: 50%.
  - Arcane Brilliance: "the target's party and raid", so index.md:22 "every caster in the raid" holds.
- **Talent text** (talents/mage.json): every rank value the pages quote. Some needed working out:
  - Arcane Meditation is 17/33/50. With Mage Armor it gives 100% at three ranks and 83% at two.
  - Improved Channeling 4/5 gives Missiles 80% and Blast or Wrath 56%.
  - Burning Soul is 23%/10% at 1 rank and 47%/20% at 2.
  - Winter's Chill 2/5 is 40%, at 20% damage taken and 7% spell damage per stack, so 5 stacks give double damage and 35%.
  - Missile Barrage is 40% from Blast and 20% from Arcane Wrath, Fireball, Frostbolt and Frostfire Bolt.
  - Heating Up is triggered by Fireball, Frostfire Bolt, Fire Blast and Scorch crits, plus every Fire Blast hit.
  - Accelerant affects Fireball, Frostfire Bolt and Arcane Wrath at 0.1 s per rank.
- **Tier gates:**
  - Combustion: the raid's 30 points above it are 28 for the damage talents and Flame Throwing, plus Burning Soul 2. The leveling build reaches the same 30 with Burning Soul 1 and Accelerant 1.
  - Arcane Meditation: 15 points in tiers 0-2 on both Fire builds.
  - Arcane Impact on Frost: 12 points in tiers 0-1.
  - Pyroblast: 12 Fire points in tiers 0-1 on the Arcane raid build. Heating Up: 15.
  - Ice Block at tier 3 needs 15 Frost points. Fire has 3, so "only Frost mages carry Ice Block" holds.
- **Trainer levels** (trainables/mage.json):
  - Fireball 1/6/12/18, Frostbolt 4/8/14/20, Conjure Water 4, Fire Blast 6, Arcane Missiles 8, Frost Nova 10, Arcane Explosion 14, Flamestrike 16.
  - Scorch 22, Cone of Cold 26, Ice Armor 30, Mage Armor 34.
  - Frostfire Bolt 40/50/60, Teleport: Dalaran 50.
  - Arcane Missiles r8 (25345) is Tome of Arcane Missiles VIII (21280, item level 60, quality 3), an Ahn'Qiraj book, as card note [6] says.
- **Racials** (spelltips/shared.json, races.json mage rows as settled in round 8): Eureka! 8 s on a 2 minute cooldown, Ruthless 1%, The Human Spirit 10%, Expansive Mind 5%, Blood Fury 10% for 15 s every 2 minutes, Berserking every 3 minutes, Wind Blessed 2%, Read Leyline 50% for 15 s every 2 minutes. Mage Skyborne are Alliance (High Order) only. No race line contradicts a measured result or the client.
- **Mana plan:** gems and Demonic Runes share one 2 minute cooldown, the Robe has its own 5, and there is a 1 minute shared lockout (round 1 O4, client ItemEffect and SpellCategories). The Robe of the Archmage (14152) is still in all three bis files.
- **Leveling 1 to 19 items:**
  - Lesser Magic Wand 11287 (req 5) and Greater Magic Wand 11288 (req 13) are crafted by Enchanting.
  - Elven Wand 5604 is the reward for Bashal'Aran (Darkshore, Alliance). Riptear's Spare Arm 286743 is the reward for Rear Guard Patrol (Horde).
  - The four class-tagged level 10 mage quests pay a staff or an orb.
- **Card notes:** Arcane [1]-[13], Fire [1]-[8] and Frost [1]-[7] agree with the prose and the client values above.
- **Header card:**
  - The Arcane rail shows the gnome raid build 32/16/3 and the Fire rail the human raid build. No prose contradicts either.
  - The Arcane key-talent lines carry the right build labels and point counts: 10 points for Improved Frostbolt and Ice Shards, 9 points for Ignite, Pyroblast, Improved Fire Blast and Heating Up, and 8 points for Arcane Mind and Arcane Instability.
- **Voice:**
  - No em or en dashes.
  - No site, sim, search, APL, pipeline, preset or band words in body prose.
  - No quoted nightly numbers. Stat paragraphs explain and do not rank.
- **Source links:** all seven fetched live this round, all HTTP 200, and each matches its label:
  - Blizzard Deep Dive recap
  - talentsforever.com/data.json
  - /sim/weights
  - /sim/drops
  - Wowhead Skyborne First Look
  - Power Up Gaming talent changes
  - /planner?class=mage

## SHOULDs (wording, not findings)

- **arcane.md:72** "it can roll Missile Barrage before the fight starts". The same sentence lands the bolt "just after the tank's first hit". Suggested: "and it can roll Missile Barrage before your first Blast."
- **frost.md:71** "Recast Ice Barrier the moment it breaks or runs out". Ice Barrier has a 30 second cooldown, so a shield that breaks early cannot be recast at once. The Cooldowns paragraph covers this with Cold Snap, and card note [2] gives the cooldown. Suggested: "Recast Ice Barrier the moment it is ready again after it breaks or runs out".
- **fire.md:116** "the Burning Soul point from 24 keeps a mob's hits from pushing back every cast". One rank is a 23% chance, so most hits still push back. Suggested: "the Burning Soul point from 24 saves the odd cast from pushback".
- **arcane.md:56 and :72** "before you have the Fire points". The leveling build gets Fire points (Incineration at 41-44) but never Pyroblast. Suggested: "(an instant Frostbolt on the leveling build, which has no Pyroblast)".
- Carried from round 10: Accelerant is missing from the index.md:32 new-talent list.

## Decisions

- **D1.** The arcane.md:72 Missile Barrage timing is a SHOULD, not a finding. The proc rolls when the cast completes, at about the moment the pull starts, so "before the fight starts" is loose, not false.
- **D2.** Fire's "two Flamestrikes eat a third of your bar" (fire.md:80) is not filed. 1,980 mana is a quarter to a third of a level 60 bar depending on gear, and nothing on 70755 sets the bar to a single value.
- **D3.** "Before you have the Fire points" stays a SHOULD, as rounds 10 to 12 accepted it. No leveling build can cast Pyroblast, so the instruction cannot be misapplied.

## Sims requested

- **PROVISIONAL (BiS re-rank).** bis/mage-*.json are still the d1d9f2775 runs (2026-11-04T03:13Z to 03:17Z), on the old strings:
  - The Arcane band 60 uses 31/4/16 and 32/3/16, not the shipped raid build 32/16/3.
  - The Fire band 60 carries Impact 1.
  - The Frost band 60 uses 12/8/31, not 15/5/31.
  - Re-rank all three on the shipped strings. Then check the Arcane statPriority frontmatter: it has no Fire power, though Fire spells are half the raid build's casts.
- **Race re-runs at 60 on the new builds:**
  - Arcane raid 32/16/3: Gnome, Human, Undead, Orc, Troll.
  - Frost 15/5/31: Gnome, Human, Skyborne, Undead, Troll, Orc.
- **Carried from round 12:**
  - The Arcane raid pre-pull: Frostbolt r10 against Pyroblast r8. Seeds 1 and 50001; adopt only on a gain of 1% or more beyond error.
  - The Presence of Mind target: Pyroblast against Arcane Wrath and Arcane Blast.
  - Arcane Wrath at 3 stacks against 4.
  - S1, S2, SIM-R2-1, SIM-R2-3, SIM-5, SIM-6 and SIM-2.

## Sim findings

- None new. SIM-R7-1 is carried: the engine does not model Arcane Subtlety's 15 spell penetration.

## Sources

- data/builds/1.60.1.70755:
  - talents/mage.json: every rank text.
  - trainables/mage.json.
  - spelltips/mage.json: 10181, 25304, 401502, 1237312, 1237313, 10151, 25306, 10207, 10199, 18809, 1239700, 1326082, 1240047, 13033, 11958, 12472, 12051, 12042, 12043, 11129, 10202, 10216, 10187, 10161, 10054, 10053, 22783, 10220, 23028, 25345 and 10212.
  - spelltips/shared.json: the racials.
  - races.json.
  - items.json: 21280, 21279, 21214 and 14152.
  - addon-builds.json: mage-arcane, mage-fire and mage-frost, both orders.
  - bis/mage-{arcane,fire,frost}.json: generated_at, engine_version and band 60 talents.
- data/curated/apl/mage-arcane.json, mage-fire.json and mage-frost.json.
- design/reviews/guide-accuracy/round1/mage.md, O4: the mana gem, rune and Robe cooldowns.
- design/reviews/guide-accuracy/round4/mage.md and round5/mage.md: wands and class quests.
- Commit 604b78a80: round 12 fixes.
- Blizzard, Deep Dive panel recap: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- Wowhead, Skyborne First Look: https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
- Power Up Gaming, Forever talent changes vs Classic: https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/
- Talents Forever data export: https://talentsforever.com/data.json
