# Guide accuracy round 10: Mage

Audited on the round 10 audit tree, origin/main 84460cec8: all nine guide rounds plus sim round 15. Client 1.60.1.70755, engine pin 9f7f1f6e4 (sim/enginever/version.go). Engine source read at that pin.
Files: web/src/content/guides/mage/index.md, arcane.md, fire.md and frost.md, read line by line. Every priorityList and prepull note in data/curated/apl/mage-{arcane,fire,frost}.json was read too.
Evidence:
- data/builds/1.60.1.70755: talents, trainables, spelltips, spellconst, addon-builds.json and bis/mage-*.json.
- Engine 9f7f1f6e4: sim/mage/arcane_wrath.go.
- Sim round 15 commits: e01dc2dff, 071e7b59e, 79fb5536d and 57b90e032.
- Mage reports from rounds 1 to 9, and ledger.md.
- Warcraft Tavern's Forever Arcane guide.
- Live fetches of every source link.

## What changed since round 9 (645caa983..84460cec8, mage only)

Sim round 15 changed three of the five mage builds, the Arcane rotation and the shipped leveling orders. It only touched guide frontmatter, plus the R9-1 Eureka! lines. Most of the body prose still describes the old builds. That is where every finding this round comes from.

- **Arcane raid** (e01dc2dff) is now `131205213100211531/525000012001/003` (32/16/3):
  - Arcane: Wand Specialization 1, Improved Channeling 3, Arcane Resilience 1, Arcane Subtlety 2, Arcane Concentration 5, Missile Barrage 2, Arcane Geometry 1, Arcane Impact 3, Arcane Blast 1, Arcane Meditation 2, **Arcane Wrath 1**, Presence of Mind 1, Arcane Mind 5, Arcane Instability 3, Arcane Power 1.
  - Fire: **Incineration 5, Accelerant 2, Ignite 5, Pyroblast 1, Improved Fire Blast 2, Heating Up 1**.
  - Frost: Precision 3 only. No Improved Frostbolt, Ice Shards or Frost Channeling.
  - The rotation now casts Pyroblast under Presence of Mind and on 3 Heating Up stacks, Fire Blast on cooldown, Arcane Wrath at 3 or more Blast stacks and a Fireball filler. Each of these is gated on Heating Up (400625) being known, so the leveling build keeps Frostbolt.
  - Arcane Power and Presence of Mind now wait 3 s (57b90e032).
- **Arcane leveling** is unchanged at `141005213100301531/4/05350003` (31/4/16). Its shipped order changed: Wand Specialization 10, Improved Channeling 11-14, **Arcane Concentration 15-19, Arcane Blast 20, Arcane Resilience 21, Missile Barrage 22-23**, Arcane Geometry 24, Arcane Impact 25-27, Arcane Meditation 28-30, Presence of Mind 31, Arcane Mind 32-36, Arcane Instability 37-39, Arcane Power 40, Incineration 41-44, Improved Frostbolt 45-49, Precision 50-52, Ice Shards 53-57, Frost Channeling 58-60.
- **Fire leveling** (071e7b59e) is now `2520150000002/51520101230133031/003`, with **Burning Soul 1 in place of Impact 1** and Accelerant 1 kept. The order is unchanged except for Burning Soul at 24.
- **Fire raid** is unchanged: `2530050000002/50520201230133031/003`.
- **Frost (both builds)** (79fb5536d) is now `230005203/5/253510130101301221` (15/5/31):
  - Arcane: Wand Specialization 2, Improved Channeling 3, Arcane Concentration 5, Missile Barrage 2, **Arcane Impact 3**.
  - Fire: **Incineration 5 only**. Accelerant and Flame Throwing are gone.
  - Shipped order: **Improved Frostbolt 10-14, Frost Warding 15-16**, Precision 17-19, Ice Shards 20-24, Frostbite 25, Permafrost 26, Frost Channeling 27-29, Ice Lance 30, Ice Block 31, Shatter 32-34, Cold Snap 35, Fingers of Frost 36-37, Winter's Chill 38-39, Ice Barrier 40, Wand Specialization 41-42, Improved Channeling 43-45, Arcane Concentration 46-50, Missile Barrage 51-52, Arcane Impact 53-55, Incineration 56-60.

All five builds are legal on 70755: 51 points, tier locks hold and prerequisites are met. Every level above was decoded from addon-builds.json, where index 0 is level 10.

R9-1 fix re-checked: "Eureka! makes your damage spells hit 10% harder and cost 10% less" (index.md:38, arcane.md:96, fire.md:98, frost.md:99). This matches the mage row 1259817. True.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| MAGE-R10-1 | CONFIRMED (high) | talent | arcane.md:55, :56, :59, :62, :64, :66 | :64 "The raid build trades a rank of Improved Channeling and the fourth Incineration for both ranks of Arcane Subtlety"; :62 "Both builds ... then go to Frost for the filler"; :59 "Improved Frostbolt and Ice Shards (10 points): Frostbolt is half your casts"; :56 Presence of Mind "Spend it on Frostbolt"; :55 "Arcane Meditation (3 points)"; :66 "Fire Blast on the move is the only Fire spell you cast." | The raid build in the frontmatter (arcane.md:7) decodes to 32/16/3. It has 5 Incineration (more, not fewer), Accelerant 2, Ignite 5, Pyroblast, Improved Fire Blast 2, Heating Up, Arcane Wrath and Arcane Meditation 2. It has no Improved Frostbolt, Ice Shards or Frost Channeling. The rotation (mage-arcane.json priorityList [4], [7], [8], [9], [11]) casts Pyroblast under Presence of Mind and on 3 Heating Up stacks, Fire Blast on cooldown, Arcane Wrath at 3 stacks and a Fireball filler. Commit e01dc2dff changed only the frontmatter line. | :55 "**Arcane Meditation** (3 points, 2 in the raid build): half your mana regeneration keeps running while you cast, a third at two ranks. With Mage Armor on top, all of it runs at three ranks and most of it at two." :56 "**Presence of Mind** (1 point): your next spell under 10 seconds is instant. In a raid spend it on Pyroblast, your longest cast; on the way to 60, before you have the Fire points, spend it on Frostbolt. Never spend it on a 1.5 second Arcane Blast that the global cooldown already caps." :59 "**Improved Frostbolt** and **Ice Shards** (10 points, leveling build): Frostbolt is half your casts on the way to 60, and a 2.5 second Frostbolt with full crit damage is the filler that pays for the Blast stacks." Add after :59: "- **Arcane Wrath** (1 point, raid build): a 3.5 second Arcane bolt on a 6 second cooldown that adds 5% of your maximum mana as spell damage. Cast it every time it is ready while you hold 3 Blast stacks." and "- **Ignite**, **Pyroblast**, **Improved Fire Blast** and **Heating Up** (9 points, raid build): Fire Blast every 6 seconds, and every Fire Blast hit and every Fire Blast or Fireball crit cuts your next Pyroblast's cast by 25%, up to 3 stacks. Three stacks make Pyroblast a 1.5 second cast, and every Fire crit burns for another 40% through Ignite." :62 "The 60 build spends 31 points in Arcane for Arcane Power, then goes to Frost for the filler: **Frost Channeling** (3 points: 15% cheaper Frostbolts and 30% less threat from them), all five ranks of Ice Shards, and four ranks of **Incineration** in Fire for Arcane Blast crit. Improved Channeling stays at four ranks so mobs hitting you do not push back every Blast, and one point of Arcane Geometry adds 3 yards to Arcane Blast and Missiles, so you can blast from a little further out." :64 "The raid build puts 32 points in Arcane and takes its second tree from Fire. Both ranks of **Arcane Subtlety** cut 30% of the threat from your Arcane spells, which is what lets you press Arcane Power early in a pull, and Arcane Wrath costs a rank of Arcane Meditation and a rank of Improved Channeling. Bosses do not push back your casts the way packs of mobs do, so that lost rank costs nothing. Sixteen go into Fire: five ranks of Incineration, two of **Accelerant** to take 0.2 seconds off Fireball and Arcane Wrath, then Ignite, Pyroblast, Improved Fire Blast and Heating Up. Precision's three points are all it keeps in Frost." :66 "Neither build takes Impact: its stun does nothing to a boss. [Open the planner](/planner?class=mage) to change the build." |
| MAGE-R10-2 | CONFIRMED (high) | rotation | arcane.md:30, :37, :42, :70, :72, :80, :82 | :30 "fills with Frostbolt"; :70 "Start a Frostbolt 2.5 seconds before the pull ... put the instant Frostbolt out"; :72 "switch to Frostbolt until the 18 seconds run out"; :80 "Frost Channeling cuts 30% of the threat from your Frostbolts, and the raid build's Arcane Subtlety does the same"; :82 "hold Presence of Mind for an instant Frostbolt on the run"; :42 "the build already puts 16 points in Frost"; :37 "Frostbolt filler and Frost Channeling" | On the raid build, priorityList [11] Fireball comes ahead of [12] Frostbolt, so Frostbolt never fills. [4] Pyroblast takes Presence of Mind ahead of [5] Frostbolt. [9] Arcane Wrath at 3 or more stacks and [8] Fire Blast on cooldown both come ahead of [10] Arcane Blast. Without Improved Frostbolt, Frostbolt r10 is a 3.0 s cast (spelltips 10181, 3000 ms). The raid build has 3 Frost points and no Frost Channeling. | :30 "Arcane stacks Arcane Blast for spell damage, fires free Missiles on every Missile Barrage and saves its burst for Arcane Power. In a raid it fills with Arcane Wrath, Fire Blast, Heating Up Pyroblasts and Fireball; on the way to 60 it fills with Frostbolt." :37 "On the way to 60, Frostbolt filler and Frost Channeling keep you slowing adds and light on threat". :42 "No Ice Block: the leveling build spends its 16 Frost points on the Frostbolt filler and the raid build keeps only Precision in Frost". :70 replace the Frostbolt sentences with "Start a Frostbolt before the pull so it lands just after the tank's first hit: 2.5 seconds out with the leveling build's Improved Frostbolt, 3 seconds with the raid build. It can roll Missile Barrage before the fight starts. Blast twice while the tank puts a couple of Sunders up, then press Arcane Power, Presence of Mind, your trinkets and your racial together, put the instant Pyroblast out (an instant Frostbolt before you have the Fire points) and go back to blasting." :72 open with "**The loop in a raid.** Blast to 3 stacks. Missiles the moment Missile Barrage lights up, Pyroblast on 3 Heating Up stacks when it is a 1.5 second cast, Fire Blast every time it is ready, and Arcane Wrath every time it comes off its 6 second cooldown while you hold 3 stacks. Past 3, keep blasting only while your mana bar is fuller than the boss's health bar and you are above 20% mana; otherwise Fireball fills until the 18 seconds run out. On the way to 60 the loop is the same with Frostbolt as the filler and no Fire spells." Keep the stack-refresh and Clearcasting sentences. :80 "Frost Channeling cuts 30% of the threat from your Frostbolts on the way to 60, and the raid build's Arcane Subtlety does the same for your Arcane spells, but ...". :82 "hold Presence of Mind for an instant Pyroblast on the run, or an instant Frostbolt before you have the Fire points". |
| MAGE-R10-3 | CONFIRMED | stats | arcane.md:86, :92 | :92 "Arcane casts Arcane and Frost spells ... Frost power covers only the Frostbolts"; :86 "Spell power goes into every Blast, Missile and Frostbolt you cast. Arcane power covers your Blasts and Missiles but not your Frostbolts." | The raid build (the band 60 raid table) casts Arcane Blast, Arcane Wrath and Missiles (Arcane) and Fire Blast, Pyroblast and Fireball (Fire). Apart from the prepull bolt it casts no Frostbolt. Arcane power does not cover the Fire half. | :86 "Spell power goes into every spell you cast. Arcane power covers your Blasts, Arcane Wrath and Missiles but not your Fire spells in a raid or your Frostbolts on the way to 60." Replace the last sentence with "Crit hits harder on your Arcane spells, because Arcane Mind doubles their crit bonus, and in a raid every Fire crit is an Ignite and a Heating Up stack." :92 "Arcane casts spells of two schools, Arcane and Fire in a raid and Arcane and Frost on the way to 60, so spell power that names one school covers only part of your damage: Arcane power covers your Blasts, Wrath and Missiles, and Fire power only your Fire Blasts, Pyroblasts and Fireballs." |
| MAGE-R10-4 | CONFIRMED | leveling | arcane.md:108, :114 | :108 "Arcane Resilience at 15 turns a quarter of your Intellect into armor, and 16 to 19 go into Arcane Concentration"; :114 "Missile Barrage comes at 21 and Arcane Blast at 27" | The shipped Arcane leveling order (addon-builds.json mage-arcane.leveling, indices 5-13) is Arcane Concentration 15-19, Arcane Blast 20, Arcane Resilience 21, Missile Barrage 22-23. | :108 "At 10 the first point goes into Wand Specialization for 13% more wand damage, then 11 to 14 into Improved Channeling, so most hits no longer push back your Missiles, and 15 to 19 into Arcane Concentration for free casts." :114 second and third sentences: "Arcane Blast comes at 20, Arcane Resilience at 21 for armor from your Intellect, and Missile Barrage at 22 and 23, and from there you level on Blast and free Missiles." |
| MAGE-R10-5 | CONFIRMED | talent, leveling | fire.md:66, :116 | :66 "The 60 build pays that toll with a point in **Impact**, a 3% chance to stun a mob for 2 seconds while you farm, and a point in Accelerant. The raid build pays it with two ranks of Burning Soul instead"; :116 "and Impact now and then stops a caster or a runner for you" | The Fire leveling build (fire.md:7, 071e7b59e) has Burning Soul 1 and Impact 0, and the shipped order puts Burning Soul at 24. Burning Soul 1/3: "23% chance to not lose casting time ... threat ... by 10%". | :66 "The 60 build pays that toll with a point in **Burning Soul**, a 23% chance to keep a cast from being pushed back and 10% less Fire threat, and a point in Accelerant. The raid build pays it with two ranks of Burning Soul instead: 20% less threat from your Fire spells and a 47% chance to keep a cast from being pushed back. It also moves the point of Magic Absorption into a third rank of Arcane Resilience." :116 "Fire levels best from 30, once Improved Scorch and Heating Up are in, and the Burning Soul point from 24 keeps a mob's hits from pushing back every cast." |
| MAGE-R10-6 | CONFIRMED | talent | frost.md:63 | "Twelve go into Arcane for Arcane Concentration ... and Missile Barrage. The last eight go into Fire: five ranks of Incineration, which adds 10% crit to Ice Lance, two of Flame Throwing for 6 more yards on Fire Blast, and one in Accelerant." | frost.md:6-7 decode to 15 Arcane points (Wand Specialization 2, Improved Channeling 3, Arcane Concentration 5, Missile Barrage 2, Arcane Impact 3) and 5 Fire points (Incineration 5). There is no Flame Throwing or Accelerant (79fb5536d). Arcane Impact 3/3: "critical strike chance of your Arcane spells by 6%". | "Fifteen go into Arcane: **Arcane Concentration** (a free spell 10% of the time), Missile Barrage, and three ranks of **Arcane Impact** for 6% more crit on the Missiles it fires, with Wand Specialization and Improved Channeling as the toll to get there. The last five go into Fire for five ranks of **Incineration**, which adds 10% crit to Ice Lance." |
| MAGE-R10-7 | CONFIRMED | leveling | frost.md:111 | "At 10 and 11 the points go into Frost Warding, 30% more armor from Frost Armor, then Improved Frostbolt from 12 to 16 for a Frostbolt 0.5 sec faster, and Precision from 17 to 19" | The shipped Frost leveling order (addon-builds.json mage-frost.leveling, indices 0-9) is Improved Frostbolt 10-14, Frost Warding 15-16, Precision 17-19. | "From 10 to 14 the points go into Improved Frostbolt for a Frostbolt 0.5 sec faster, then Frost Warding at 15 and 16 for 30% more armor from Frost Armor, and Precision from 17 to 19 for 6% hit." |
| MAGE-R10-8 | CONFIRMED | rotation (card note) | data/curated/apl/mage-arcane.json priorityList [3], [12] | [3] Presence of Mind: "spend it on a 2.5 second Frostbolt"; [12] "Frostbolt fills the rest of the 18 seconds once you stop blasting" | On the raid build, row [4] (Pyroblast while Presence of Mind is up, gated on Heating Up) comes ahead of [5] Frostbolt, and row [11] Fireball comes ahead of [12] Frostbolt. That build's Frostbolt is also 3.0 s, because it has no Improved Frostbolt. Each note is contradicted by the row above it on the same card. | [3] "Presence of Mind on cooldown, with Arcane Power. Your next spell is instant: spend it on Pyroblast, your longest cast, or on Frostbolt before you have the Fire points. Never spend it on an Arcane Blast that the global cooldown already caps at 1.5 seconds." [12] "Frostbolt fills the rest of the 18 seconds once you stop blasting on the way to 60; in a raid Fireball takes its place. The stacks' spell damage still applies to it, it procs Missile Barrage 20% of the time, and it costs a fraction of a stacked Blast. When the stacks fall off, the next Blast starts a cheap new set." |

Counts: CONFIRMED 8 (two high), LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

The three worst:
1. R10-1 and R10-2: the Arcane page describes a Frost-filler raid build that no longer exists. The shipped raid build is 32/16/3 Fire with Arcane Wrath, and Presence of Mind now goes on Pyroblast.
2. R10-6: the Frost build paragraph lists Flame Throwing and Accelerant, which the build dropped for Arcane Impact.
3. R10-4 and R10-7: the leveling levels for Arcane (Arcane Blast at 20, not 27) and Frost (Improved Frostbolt at 10, not 12) contradict the shipped order.

No findings in these areas:
- index.md: Precision "Arcane and Fire builds both dip three points" still holds, and so does "Arcane builds have no spare point" for Ice Block.
- Every part of fire.md other than R10-5.
- Every part of frost.md other than R10-6 and R10-7, including Frostfire Bolt "close to a 3 second cast", which is now exactly 3.0 s.
- The Fire and Frost card notes.
- Spell levels and values.
- Race paragraphs.
- Professions.

## SHOULDs (wording, not findings)

- The new mage-arcane.json rows [4], [7], [8], [9] and [11] (Pyroblast, Fire Blast, Arcane Wrath, Fireball) carry no notes. Give each a note, for example:
  - [9] "Arcane Wrath every time it is ready while you hold 3 Blast stacks: the stacks' spell damage and 5% of your maximum mana both go into it."
  - [11] "Fireball fills in a raid once you stop blasting; it procs Missile Barrage 20% of the time and its crits feed Heating Up and Ignite."
- frost.md:85 "still close to a 3 second cast": with Accelerant gone it is exactly 3 seconds. "so it stays a 3 second cast" is exact.
- Carried from round 9: fire.md:78 "every cast is a Heating Up stack" should say "every hit"; frost.md:42 "melee range"; the Mana Ruby and Major Mana Potion sentence; Accelerant missing from index.md:32.

## Links checked

All return 200 and match their labels:
- Blizzard Deep Dive recap
- Wowhead Skyborne First Look
- Power Up Gaming talent changes
- Talents Forever data.json
- /sim/weights
- /sim/drops
- /planner?class=mage

Every link follows the round 4 source-link rule.

## Decisions

- **D1. Arcane prose follows the shipped 32/16/3 raid build and its rotation.** Warcraft Tavern's Forever Arcane guide plays the same Fire second tree (Incineration, Improved Fire Blast, Ignite). Its rotation is the same: Arcane Wrath at 3 or 4 Blast stacks, Fire Blast on cooldown, Fireball or Frostbolt as filler. The leveling build keeps the Frost filler. The prose names which build each line belongs to.
- **D2. Presence of Mind goes on Pyroblast in a raid.** Warcraft Tavern names Arcane Blast or Arcane Wrath. Arcane Blast saves nothing against the 1.5 s global cooldown, and Arcane Wrath saves 3.3 s while Pyroblast saves 6 s. The measured rotation (row [4]) uses Pyroblast. Rows that might change this are listed under Sims requested.

## Sims requested

- **PROVISIONAL (BiS re-rank).** bis/mage-arcane.json and bis/mage-frost.json were generated at d1d9f2775 (2026-11-04T03:13Z) on the old talent strings. The Arcane raid string is 131205213100301531-3-0535..., and Frost has the 8 Fire points.
  - Re-rank Arcane on the 32/16/3 build. Fire power now covers half its casts, but the Arcane statPriority frontmatter has no Fire power.
  - Re-rank Frost on the 15/5/31 build.
- **Race re-runs at 60 on the new builds.**
  - Arcane raid 32/16/3: Gnome, Human, Undead, Orc, Troll.
  - Frost 15/5/31: Gnome, Human, Undead, Troll, Orc.
  - The round 3 race picks were measured on the old builds. The prose order (Gnome then Human, Undead then Orc or Troll) holds until these report.
- **Arcane raid 60, Presence of Mind target.** Compare Pyroblast (current), Arcane Wrath and Arcane Blast. Also compare Arcane Wrath at 3 stacks (current) against 4 stacks. Use seeds 1 and 50001.
- Carried unchanged from round 9: S1 and S2 from rounds 6 and 7, SIM-R2-1 and SIM-R2-3, SIM-5 and SIM-6, and SIM-2. SIM-R2-2 (Arcane Fire second tree) is now adopted (e01dc2dff), so it is closed.

## Sim findings

- None new. SIM-R7-1 (Arcane Subtlety's 15 spell penetration not modeled) is carried. Arcane Wrath is modeled at 9f7f1f6e4 (sim/mage/arcane_wrath.go: 5% of maximum mana, client ranks, cooldown and cast).

## Sources

- data/builds/1.60.1.70755:
  - talents/mage.json: Arcane Wrath, Arcane Impact, Burning Soul, Accelerant, Arcane Meditation, Improved Channeling, Heating Up, Ignite, Missile Barrage, Frost Warding.
  - spelltips/mage.json: 1326079-1326082 Arcane Wrath (3.5 s, 6 s cooldown, 395 mana at 60), 10181, 25304, 10151, 18809, 10199.
  - trainables/mage.json.
  - addon-builds.json: mage-arcane, mage-fire and mage-frost.
  - bis/mage-arcane.json and bis/mage-frost.json (engine_version and talents).
- data/curated/apl/mage-arcane.json (priorityList [2]-[12], prepull), mage-fire.json and mage-frost.json.
- Commits e01dc2dff, 071e7b59e, 79fb5536d and 57b90e032 (sim round 15, mage).
- Engine 9f7f1f6e4: sim/mage/arcane_wrath.go.
- Warcraft Tavern, Arcane Mage Guide (Forever): https://www.warcrafttavern.com/forever/guides/arcane-mage/
- Mobalytics Forever Arcane Mage guide: https://mobalytics.gg/wow-forever/classes/arcane-mage-guide (403 on fetch, treated as unknown).
- Blizzard, Deep Dive panel recap: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- Wowhead, Skyborne First Look: https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
- Power Up Gaming, Forever talent changes vs Classic: https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/
- Talents Forever data export: https://talentsforever.com/data.json

## Fix log (round 10)

All three spec pages already carried `updated: 2026-11-04`; index.md is unchanged. The prose test (`-t mage`) is green. The Stat priority sections on arcane, fire and frost give reasons only and point to the table, with no ranking words.

- **MAGE-R10-1**: fixed in arcane.md.
  - :55 Arcane Meditation, 3 points or 2 in the raid build. Mage Armor is 50% on 70755 (spelltips 22783), so three ranks give all of your regeneration and two ranks give most of it.
  - :56 Presence of Mind goes on Pyroblast in a raid and on Frostbolt before the Fire points.
  - :59 Improved Frostbolt and Ice Shards are marked as the leveling build.
  - :60 new Arcane Wrath line (talents 113627, spelltips 1326082).
  - :61 new Ignite/Pyroblast/Improved Fire Blast/Heating Up line (talents 400624, Ignite 5/5 = 40%, Pyroblast 6000 ms).
  - :64 is now the 60-build paragraph.
  - :66 is the raid-build paragraph. Arcane Subtlety 2 plus Arcane Wrath 1 come from Improved Channeling 4→3, Arcane Meditation 3→2 and a 32nd Arcane point. The 16 Fire points and Precision are listed.
  - :68 now says neither build takes Impact.
- **MAGE-R10-2**: fixed in arcane.md.
  - :30 overview and :37 strength.
  - :42 weakness.
  - :72 opener. Frostbolt is 2.5 s with Improved Frostbolt and 3 s without it. The instant Pyroblast goes out, or a Frostbolt before you have the Fire points.
  - :74 the raid loop, with the leveling loop after it.
  - :82 threat. Added that the raid build's Fire spells get no threat cut, because it has no Burning Soul (sim-lane note).
  - :84 movement. Added Arcane Wrath and Fireball to the standing-still list, and Presence of Mind goes on Pyroblast.
- **MAGE-R10-3**: fixed at arcane.md:88 (Stat priority) and :94 (Enchants).
- **MAGE-R10-4**: fixed at arcane.md:110 and :116.
- **MAGE-R10-5**: fixed at fire.md:66 (Burning Soul 1/3 client text) and :116.
- **MAGE-R10-6**: fixed at frost.md:63 (Arcane Impact 3/3 client text).
- **MAGE-R10-7**: fixed at frost.md:111.
- **MAGE-R10-8**: fixed in mage-arcane.json.
  - Notes [3] and [12] rewritten.
  - Note [5] now says "before you have the Fire points".
  - Note [6] Missile Barrage now names Arcane Wrath and Fireball at 20%, from the talent rank 2 text.
- **SHOULD, new rows**: notes written for mage-arcane.json [4], [7], [8], [9] and [11]. Every action is unchanged; the diff touches notes only.
- **SHOULD, Frostfire Bolt**: frost.md:85 now reads "so it stays a 3 second cast".
- **SHOULD, Heating Up wording**: fire.md:78 now says "every hit is a Heating Up stack". The card note already said this.
- **SHOULD, Cone of Cold range**: frost.md:42 now says "within 10 yards" instead of melee range. Winter's Chill stacks from range, so it was dropped from that line.
- **SHOULD, potion timing**: arcane.md:78 and fire.md:78 now read "a Mana Ruby when you are 1,200 mana down, a Major Mana Potion the next time you are well down". Frost's mana line names no potion, so it is unchanged.
- **Left: index.md:32 Accelerant.** Accelerant is 1.12's Improved Fireball (0.1 s a rank, 5 ranks) with Frostfire Bolt and Arcane Wrath added. It is not a new talent the way the others on that list are.
- **D1 and D2**: applied as above.
- **SIM findings**: none new. The Sims requested list stands. The prose takes no position on the Presence of Mind target beyond what the shipped row casts, or on 3 against 4 stacks for Arcane Wrath. Both stay true whatever those runs find.
