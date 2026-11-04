# Guide accuracy round 11: Mage

Audited on the round 11 audit tree, main d54de48aa: ten guide rounds, sim round 15, leveling-order. Client 1.60.1.70755, engine pin 9f7f1f6e4 (sim/enginever/version.go).

Files read line by line:
- web/src/content/guides/mage/index.md, arcane.md, fire.md and frost.md.
- Every note and action in data/curated/apl/mage-{arcane,fire,frost}.json, priorityList and prepull.

Read first: mage reports for rounds 1 to 10 (with fix logs), round4/mage-levels.md and ledger.md.

## What changed since round 10

Since the round 10 fix commit (eabba8f84), main has four commits: the leveling-order lane (ec60e6abf, 344b90a86) and two merges. None of them touch mage guide prose, mage APL files, addon-builds.json, talents or bis/*.json. Every mage line on the page is therefore round 10's text, and this audit re-checks round 10's rewrites against the client.

## Builds decoded (talents/mage.json, tier/column order; all five legal: 51 points, tier locks and prerequisites hold)

- **Arcane leveling** `141005213100301531/4/05350003` (31/4/16)
  - Arcane: Wand Specialization 1, Improved Channeling 4, Arcane Resilience 1, Arcane Concentration 5, Missile Barrage 2, Arcane Geometry 1, Arcane Impact 3, Arcane Blast 1, Arcane Meditation 3, Presence of Mind 1, Arcane Mind 5, Arcane Instability 3, Arcane Power 1.
  - Fire: Incineration 4.
  - Frost: Improved Frostbolt 5, Precision 3, Ice Shards 5, Frost Channeling 3.
- **Arcane raid** `131205213100211531/525000012001/003` (32/16/3)
  - Arcane: Wand Specialization 1, Improved Channeling 3, Arcane Resilience 1, Arcane Subtlety 2, Arcane Concentration 5, Missile Barrage 2, Arcane Geometry 1, Arcane Impact 3, Arcane Blast 1, Arcane Meditation 2, Arcane Wrath 1, Presence of Mind 1, Arcane Mind 5, Arcane Instability 3, Arcane Power 1.
  - Fire: Incineration 5, Accelerant 2, Ignite 5, Pyroblast 1, Improved Fire Blast 2, Heating Up 1.
  - Frost: Precision 3.
- **Fire leveling** `2520150000002/51520101230133031/003` (17/31/3)
  - Arcane: Wand Specialization 2, Improved Channeling 5, Arcane Resilience 2, Magic Absorption 1, Arcane Concentration 5, Arcane Meditation 2.
  - Fire: Incineration 5, Accelerant 1, Ignite 5, Flame Throwing 2, Burning Soul 1, Pyroblast 1, Improved Fire Blast 2, Improved Scorch 3, Heating Up 1, Master of Elements 3, Critical Mass 3, Fire Power 3, Combustion 1.
  - Frost: Precision 3.
- **Fire raid** `2530050000002/50520201230133031/003` (17/31/3)
  - Arcane: Arcane Resilience 3 and no Magic Absorption.
  - Fire: Burning Soul 2 and no Accelerant.
  - Otherwise as the leveling build.
- **Frost (both)** `230005203/5/253510130101301221` (15/5/31)
  - Arcane: Wand Specialization 2, Improved Channeling 3, Arcane Concentration 5, Missile Barrage 2, Arcane Impact 3.
  - Fire: Incineration 5.
  - Frost: Frost Warding 2, Improved Frostbolt 5, Precision 3, Ice Shards 5, Frostbite 1, Permafrost 1, Frost Channeling 3, Ice Lance 1, Ice Block 1, Shatter 3, Cold Snap 1, Fingers of Frost 2, Winter's Chill 2, Ice Barrier 1.

Shipped leveling orders (addon-builds.json, index 0 = level 10):
- **Arcane**: Wand Specialization 10, Improved Channeling 11-14, Arcane Concentration 15-19, Arcane Blast 20, Arcane Resilience 21, Missile Barrage 22-23, Arcane Geometry 24, Arcane Impact 25-27, Arcane Meditation 28-30, Presence of Mind 31, Arcane Mind 32-36, Arcane Instability 37-39, Arcane Power 40, Incineration 41-44, Improved Frostbolt 45-49, Precision 50-52, Ice Shards 53-57, Frost Channeling 58-60.
- **Fire**: Incineration 10-14, Accelerant 15, Ignite 16-19, Pyroblast 20, Ignite 21, Flame Throwing 22-23, Burning Soul 24, Improved Fire Blast 25-26, Improved Scorch 27-29, Heating Up 30, Master of Elements 31-33, Critical Mass 34-36, Fire Power 37-39, Combustion 40, then the Arcane points and Precision 58-60.
- **Frost**: Improved Frostbolt 10-14, Frost Warding 15-16, Precision 17-19, Ice Shards 20-24, Frostbite 25, Permafrost 26, Frost Channeling 27-29, Ice Lance 30, Ice Block 31, Shatter 32-34, Cold Snap 35, Fingers of Frost 36-37, Winter's Chill 38-39, Ice Barrier 40, then Arcane 41-55 and Incineration 56-60.

The three Levels 1 to 19 paragraphs and the three Levels 20 to 60 paragraphs match these orders level for level.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| MAGE-R11-1 | CONFIRMED (low) | consumable/race, stale build | arcane.md:98 | "**Human** is close behind with 1% crit and 10% Spirit, which runs all the time thanks to Arcane Meditation and Mage Armor." | The Races section compares raid picks. Since sim round 15 (e01dc2dff) the raid build (arcane.md:7) has **2/3 Arcane Meditation**: "Allows 33% of your Mana regeneration to continue while casting" (talents/mage.json 18462 rank 2). Mage Armor adds 50% (spelltips 22783), so 83% runs while casting, not all of it. The page says so itself at arcane.md:55: "all of it runs at three ranks and most of it at two". The sentence was true in round 3, when the raid build had 3 ranks; the build change made it false and round 10 missed it. | "**Human** is close behind with 1% crit and 10% Spirit, and Arcane Meditation and Mage Armor keep most of that Spirit's regeneration running while you cast." |

Counts: CONFIRMED 1 (low), LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

The three worst:
1. MAGE-R11-1. Arcane's Human line says Spirit regeneration "runs all the time", but the raid build's two ranks of Arcane Meditation plus Mage Armor keep 83% of it running. The page's own talent line says "most of it at two".
2. None else confirmed.
3. None else confirmed.

## Round 10's own fixes, re-checked (all hold)

- **arcane.md:55, Arcane Meditation.** 17/33/50% per rank, and Mage Armor 50%, so all of it at three ranks and most of it at two. True.
- **arcane.md:56 and APL [3]/[4], Presence of Mind on Pyroblast.**
  - Presence of Mind affects spells "with a casting time less than 10 sec" (12043).
  - Pyroblast r8 is 6.0 s (18809), the longest such cast the rotation has.
- **arcane.md:60, Arcane Wrath.** 3.5 s cast, 6 s cooldown, 395 mana, plus 5% of maximum mana as spell damage (1326082). True. The talent line gives the base cast; arcane.md:66 adds Accelerant's 0.2 s.
- **arcane.md:61.**
  - Ignite 5 + Pyroblast 1 + Improved Fire Blast 2 + Heating Up 1 = 9 points.
  - Fire Blast's 8 s cooldown minus 2 is 6 s (10199, 11078 rank 2).
  - Heating Up: crits with Fireball, Frostfire Bolt, Fire Blast or Scorch, plus one stack per Fire Blast hit; 3 x 25% makes a 6 s Pyroblast 1.5 s.
  - Ignite 5/5 is 40%.
- **arcane.md:64 (leveling build) and :66 (raid build).**
  - The raid build's Arcane Subtlety 2 and Arcane Wrath 1 come from Improved Channeling 4 to 3, Arcane Meditation 3 to 2 and a 32nd Arcane point.
  - Its Fire points total 16: 5 + 2 + 5 + 1 + 2 + 1.
  - Accelerant rank 2 takes 0.2 s off Fireball and Arcane Wrath.
  - Improved Channeling 4/5 gives Blast 56% and Missiles 80%.
  - Arcane Geometry 1/2 adds 3 yards.
- **arcane.md:72-84 against priorityList.**
  - Missiles [6], Pyroblast on 3 Heating Up stacks [7], Fire Blast [8], Arcane Wrath at 3 or more stacks [9], the mana and boss-health Blast gate [10], Fireball [11] and Frostbolt [12].
  - The threat line is true: the raid build has no Burning Soul.
- **arcane.md:88 and :94.** The schools match the raid and leveling casts.
- **arcane.md:110 and :116.** Match the shipped Arcane order.
- **fire.md:66.** Burning Soul ranks 1 and 2 are 23%/10% and 47%/20%. The Accelerant and Magic Absorption to Burning Soul and Arcane Resilience trade checks out.
- **fire.md:116.** The Burning Soul point comes at 24.
- **frost.md:63.**
  - Arcane points total 15: 2 + 3 + 5 + 2 + 3.
  - Arcane Impact 3/3 is 6%.
  - Incineration covers Ice Lance.
  - Winter's Chill 2/5 is 40%.
- **frost.md:85.** Frostfire Bolt r3 is a 3.0 s cast, 370 mana and 292 damage, against Frostbolt r11's 290 mana and 475 damage.
- **frost.md:111.** Matches the shipped Frost order.
- **APL card notes [3], [4], [6], [7], [8], [9], [11] and [12] (mage-arcane.json).** Each agrees with its action and with the rows above it.

## No findings in these areas

- **Spell levels, ranks, costs, casts and cooldowns** (trainables/mage.json, spelltips/mage.json):
  - Frostbolt 4/8/14/20.
  - Fireball 1/6/12/18.
  - Fire Blast 6, Arcane Missiles 8, Frost Nova 10, Arcane Explosion 14, Flamestrike 16, Arcane Blast 20.
  - Scorch 22 (150 mana, 1.5 s), Cone of Cold 26, Mage Armor 34.
  - Frostfire Bolt 40/50/60, Teleport: Dalaran 50.
  - Arcane Explosion 390 mana, 10 yards.
  - Flamestrike 990 mana, 3 s cast, 8 s burn.
  - Blizzard 1,400 mana (1,190 with Frost Channeling).
  - Evocation 8 min. Arcane Power, Presence of Mind and Combustion 3 min.
  - Ice Block 5 min. Ice Barrier 30 s cooldown and 1 min buff. Ice Lance 8 s.
  - Mana Ruby restores 1,100.
- **Talent text**:
  - Arcane Blast 12% Intellect a stack, 4 stacks, 18 s, +100% cost a stack.
  - Missile Barrage 40% and 20%.
  - Arcane Concentration 10%.
  - Arcane Mind doubles the crit bonus.
  - Arcane Instability 3% and 3%.
  - Improved Scorch 3% x 5 for 30 s.
  - Fire Power 3% per effect to 12%, mixed schools excluded.
  - Master of Elements 30%.
  - Critical Mass 6%.
  - Combustion.
  - Ice Barrier 15% + 5%, no pushback.
  - Fingers of Frost 15% for 2 charges.
  - Shatter 50%.
  - Ice Shards.
  - Frost Channeling 15% and 30%.
  - Winter's Chill 5 x 20% (double damage) and 5 x 7% (35% spell damage).
- **The 1 to 19 paragraphs**: wands (11287 req 5, 11288 req 13, 286743 and 5604 as early quest wands) and the class quest staff, as checked in rounds 4 and 5.
- **index.md**:
  - Ice Block tier 3 (15 points) with Cold Snap below.
  - Precision 6% on all schools.
  - Mage Armor 50%.
  - The new-talent list.
  - Orc mage.
  - The Skyborne Alliance-only line.
- **Race lines** other than R11-1. Race picks on the new builds are still unmeasured and nothing contradicts a measured result.
- **Voice**: no em or en dashes, no site, sim or pipeline words, and no quoted nightly numbers.
- **Source links**, all fetched live with HTTP 200 and matching their labels under the round 4 source-link rule:
  - Blizzard Deep Dive recap
  - Talents Forever data.json
  - /sim/weights
  - /sim/drops
  - Wowhead Skyborne First Look
  - Power Up Gaming talent changes
  - /planner?class=mage

## SHOULDs (wording, not findings)

- arcane.md:42 "the leveling build spends its 16 Frost points on the Frostbolt filler". Three of the 16 are Precision, which is hit for every school. Suggested: "the leveling build spends its 16 Frost points on Precision and the Frostbolt filler".
- arcane.md:64 and fire.md:66 call the leveling build "the 60 build", while arcane.md:59-61 call it "leveling build". Both builds are level 60 builds, so "the leveling build" in both places reads cleaner.
- arcane.md:74 "On the way to 60 the loop is the same with Frostbolt as the filler and no Fire spells". The 1 to 19 paragraph (arcane.md:108) and the Movement line (arcane.md:84) both use Fire Blast. Suggested: "...with Frostbolt as the filler and Fire Blast only to finish a mob or on the move."
- arcane.md:84 lists the spells that need you standing still. A 1.5 s Pyroblast on 3 Heating Up stacks belongs on that list too.
- Carried from round 10: Accelerant missing from the index.md:32 new-talent list (round 10 left it, with a reason).

## Decisions

- **D1.** The Races section is about raid picks, so its regeneration claims follow the raid build: 2/3 Arcane Meditation plus Mage Armor. That makes R11-1 a finding and not a wording preference.
- **D2.** Arcane Missiles rank 8 (25345) appears in trainables at level 56 under skill_line_ability, next to Fireball r12 and Frostbolt r11. items.json carries "Tome of Arcane Missiles VIII", "Tome of Fireball XII" and "Tome of Frostbolt XI". The data does not settle whether Forever trains rank 8 or teaches it from the book, so the card note [6] "rank 8 is a book from Ahn'Qiraj" is not filed. It is a card rank question the engine note already tracks (core.IncludeAQ).

## Sims requested

- **PROVISIONAL (BiS re-rank).** bis/mage-arcane.json and bis/mage-frost.json are still the d1d9f2775 runs (2026-11-04T03:13Z/03:17Z) on the pre-round-15 strings: Arcane raid `131205213100301531-3-...` and Frost with 8 Fire points.
  - Re-rank Arcane on 32/16/3. Fire power now covers half its casts, but the Arcane statPriority frontmatter has no Fire power.
  - Re-rank Frost on 15/5/31.
- **Race re-runs at 60 on the new builds.**
  - Arcane raid 32/16/3: Gnome, Human, Undead, Orc, Troll.
  - Frost 15/5/31: Gnome, Human, Undead, Troll, Orc.
  - The prose order holds until these report.
- **Arcane raid 60, Presence of Mind target.** Pyroblast (current) against Arcane Wrath and Arcane Blast. Also Arcane Wrath at 3 stacks (current) against 4 stacks. Seeds 1 and 50001.
- Carried: S1, S2, SIM-R2-1, SIM-R2-3, SIM-5, SIM-6 and SIM-2, as listed in round 10.

## Sim findings

- None new. SIM-R7-1 (Arcane Subtlety's 15 spell penetration not modeled) is carried.

## Sources

- data/builds/1.60.1.70755:
  - talents/mage.json: every rank text quoted above, including Arcane Meditation 18462 ranks 1-3, Burning Soul, Accelerant, Improved Channeling, Winter's Chill and Heating Up 400624.
  - trainables/mage.json.
  - spelltips/mage.json: 22783, 10181, 25304, 10151, 25306, 18809, 10199, 10207, 12051, 10202, 10216, 10187, 1326079-1326082, 1239700, 10212, 25345, 13033, 1240047, 10161, 12043, 12042, 11129, 11958, 12472, 10054, 10053, 401502, 1237312 and 1237313.
  - addon-builds.json: mage-arcane, mage-fire and mage-frost, leveling and raid.
  - races.json.
  - items.json: wands and tomes.
  - bis/mage-{arcane,fire,frost}.json: engine_version, generated_at and band 60 talents.
- data/curated/apl/mage-arcane.json, mage-fire.json and mage-frost.json (priorityList and prepullActions).
- Commits eabba8f84 (round 10 fixes), e01dc2dff (Arcane raid build), ec60e6abf and 344b90a86 (leveling-order, no mage prose).
- Blizzard, Deep Dive panel recap: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- Wowhead, Skyborne First Look: https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
- Power Up Gaming, Forever talent changes vs Classic: https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/
- Talents Forever data export: https://talentsforever.com/data.json

## Fix log (round 11)

- MAGE-R11-1: fixed, arcane.md:98 with the report's text (talents/mage.json 18462 rank 2 is 33%; Mage Armor 50%).
- SHOULD arcane.md:42: applied, "on Precision and the Frostbolt filler" (leveling Frost 5+3+5+3 = 16).
- SHOULD arcane.md:64 and fire.md:66: applied, "The leveling build".
- SHOULD arcane.md:74: applied, "Fire Blast only to finish a mob or on the move".
- SHOULD arcane.md:84: applied, "and a 1.5 second Pyroblast" added to the stand-still list.
- Carried Accelerant index SHOULD: left, as round 10.
