# Guide accuracy round 8: Mage

Audited on the round 8 audit tree, land/guides-r7 head 09b98d9e4 (all seven rounds applied). Client 1.60.1.70755, engine pin d1d9f2775 (sim/enginever/version.go).
Files: web/src/content/guides/mage/index.md, arcane.md, fire.md, frost.md, read line by line (frontmatter, body, every card note in data/curated/apl/mage-{arcane,fire,frost}.json priorityList).
Evidence: data/builds/1.60.1.70755/{talents,trainables,spelltips,spellconst}/mage.json, spelltips/shared.json (racials), races.json, items.json, simconsumes.json, bis/mage-{arcane,fire,frost}.json (band 60 weights and hit_to_cap), addon-builds.json; engine d1d9f2775 sim/core/racials.go (Touch of the Grave rows), sim/mage/talents.go and armors.go (casting regen is additive); rounds 1 to 7 mage reports, round4/mage-levels.md, race-sims.md rounds 2 to 5, ledger.md; live fetches of every source link.

Diff since round 7's audit tree (394562f20): only the round 7 fixes (arcane.md:35, :76, :86; fire.md:34, :78, :88; frost.md:36, :75, :79, :81, :89, :117; index.md:5, :31; mage-fire.json priorityList[6] notes) and the Skyborne web fixes (715359838, no guide or data change for mages). No change to builds, addon-builds.json, talents, trainables, BiS or race data.

## Builds decoded (unchanged, re-checked)

Decoded by tier/column against talents/mage.json; all legal (51 points, tier locks, prerequisites).

- **Arcane leveling** `141005213100301531/4/05350003` (31/4/16): Wand Specialization 1, Improved Channeling 4, Arcane Resilience 1, Arcane Concentration 5, Missile Barrage 2, Arcane Geometry 1, Arcane Impact 3, Arcane Blast 1, Arcane Meditation 3, Presence of Mind 1, Arcane Mind 5, Arcane Instability 3, Arcane Power 1; Incineration 4; Improved Frostbolt 5, Precision 3, Ice Shards 5, Frost Channeling 3.
- **Arcane raid** `131205213100301531/3/05350003` (32/3/16): Improved Channeling 3, Arcane Subtlety 2, Incineration 3, rest as leveling.
- **Fire leveling** `2520150000002/51521001230133031/003` (17/31/3) and **raid** `2530050000002/50520201230133031/003` (raid: Arcane Resilience 3, no Magic Absorption, Accelerant or Impact, Burning Soul 2). Fire damage talents 28 + toll 2 = 30 above Combustion; Arcane 15 to Arcane Meditation's tier, 17 in all.
- **Frost (both)** `2300052/5102/253510130101301221` (12/8/31). Frost 28 + Winter's Chill 2 = 30 above Ice Barrier.

Shipped leveling order (addon-builds.json, index 0 = level 10) unchanged since round 5's decode. Every talent level in "Levels 1 to 19" and "Levels 20 to 60" of all three specs matches it, each tier lock holds at the level the prose names (Ice Block 31, Fingers of Frost 36, Winter's Chill 38, Ice Barrier 40; Arcane Power 40; Heating Up 30, Combustion 40), and the two blocks read in order without contradicting each other or the strip.

## Round 7 fixes re-checked (all true on 70755)

- arcane.md:35 "12% of your Intellect per stack ... up to 48% at four stacks": talent text, 12% per cast, 4 stacks.
- arcane.md:86, fire.md:88, frost.md:89 Intellect reasons: playerexpectedstat.csv gives ClassID 8 spell crit per Intellect at 60, so "spell crit and the mana bar every mage gets", "mana bar and adds spell crit", "mana bar and a little spell crit" are right.
- arcane.md:76 (downrank sentence removed) and frost.md:75 "Drink before pulls so you start each one full": no rank claim either way, nothing false.
- fire.md:34 "Scorch is 150 mana for a 1.5 second cast, so you land two for every Fireball": Scorch 10207 150 / 1.5 s, Fireball 3.5 s (3.4 s with the leveling build's Accelerant point).
- fire.md:78 "Scorch stays your filler when you are short too: you land two of them in the time of one Fireball": true, no rank claim.
- mage-fire.json priorityList[6]: "two Scorches in the time of one Fireball ... twice the crits for Heating Up and Ignite": true.
- frost.md:36 and :81 "15% cheaper and 30% lower in threat" / "cuts 30% of the threat from your Frostbolts": Frost Channeling 3/3 "mana cost ... by 15% ... threat ... by 30%".
- frost.md:79 "(12 with Arctic Reach)" dropped: the build has no Arctic Reach; the card note keeps the rule (10 yards, 20% more = 12), still true.
- frost.md:117 "crits more often than not": Shatter 3/3 +50% on Frozen, plus base and Intellect crit at 34.
- index.md:31 "the Arcane and Fire builds both dip three points there": Precision is a Frost talent; both off-tree builds carry 3.

## What holds (checked this round)

- Every talent line on all four pages against talents/mage.json rank text: Arcane Blast, Missile Barrage 40% Blast / 20% Frostbolt (Fireball and Frostfire Bolt also 20%), Arcane Concentration 10% "after any damage spell hits", Arcane Meditation 50% at 3/3 and 33% at 2/3, Presence of Mind "casting time less than 10 sec", Arcane Power 30% / 30% / 15 s, Arcane Mind 10% Intellect and +100% Arcane crit bonus, Arcane Instability 3% / 3%, Arcane Subtlety 15 resistance and 30% Arcane threat at 2/2, Improved Channeling 56% Blast / 80% Missiles at 4/5, Arcane Geometry 3 yd, Arcane Resilience 25% / 75% of Intellect; Incineration (Fire Blast, Ice Lance, Arcane Blast, Scorch) 10%, Accelerant 0.1 s a rank (Fireball, Frostfire Bolt, Arcane Wrath), Ignite 40% / 4 s, Flame Throwing 6 yd, Impact 3% / 2 s, Burning Soul 47% and 20% at 2/3, Improved Fire Blast 2 s / 15 yd, Improved Scorch 3% / 5 stacks / 30 s, Heating Up 25% a stack, 3 stacks, Fire Blast stack on every damaging hit, Master of Elements 30% of base cost, Critical Mass 6%, Fire Power 3% per Fire effect to 12% with mixed schools excluded, Combustion 10% per hit until 3 non-periodic crits; Frost Warding 30% Frost Armor / Ice Armor armor, Improved Frostbolt 0.5 s, Precision 6% all spells, Ice Shards +100% Frost crit bonus, Frostbite 5% at 1/3, Frost Channeling 15% / 30%, Ice Lance +300% on Frozen, Ice Block 10 s, Shatter 50%, Cold Snap, Fingers of Frost 15% / 2 spells / 15 s, Winter's Chill 20% a rank, +20% Cone of Cold damage taken and +7% spell damage a stack to 5, Ice Barrier 15% for 1 min, +5% while shielded, no pushback.
- Spell values at 60 (spelltips/spellconst): Scorch 150 / 1.5 s, Fire Blast 340 / 8 s, Pyroblast 440 / 6 s, Flamestrike 990 / 3 s / 8 s burn, Arcane Explosion 390 / 10 yd instant, Blizzard 1,400 / 8 s (1,190 with Frost Channeling), Cone of Cold 10 s, Frost Nova 25 s / 8 s root, Ice Lance 160 / 8 s, Ice Barrier 30 s, Ice Block 5 min, Cold Snap 10 min, Evocation 8 min / +1500% / 8 s, Arcane Power, Presence of Mind and Combustion 3 min, Blink 20 yd / 15 s, Mana Ruby 1,100 (so "1,200 mana down" wastes nothing), Mana Citrine 850, Frostfire Bolt 370 / 3.0 s with a 40% slow and a Frostfire burn, Frostbolt 2.5 s with Improved Frostbolt, Mage Armor 50% casting regen at every rank (additive with Arcane Meditation in the engine, armors.go:118 and talents.go:77, so "all of it" and "most of it" hold), Ice Armor chills attackers, Arcane Brilliance party and raid. Trainer levels in prose: Fireball 1/6/12/18, Frostbolt 4/8/14/20, Conjure Water 4, Fire Blast 6, Arcane Missiles 8, Frost Nova 10, Arcane Explosion 14 (rank 2 at 22), Flamestrike 16, Blizzard 20, Scorch 22, Cone of Cold 26, Ice Armor 30, Mage Armor 34, Citrine 48, Teleport: Dalaran 50, Arcane Brilliance 56, Ruby 58, Frostfire Bolt 40/50/60.
- Racials (spelltips/shared.json and engine racials.go:376): Ruthless 1% crit with all spells, The Human Spirit 10% Spirit, Eureka! 10% / 10% / 8 s / 2 min, Expansive Mind 5% mana, Blood Fury 10% spell power 15 s / 2 min, Berserking 10% 15 s / 3 min, Touch of the Grave 30% on the mage row (1260201; the shared tip's 20% is the warrior/paladin/rogue row 1260189), 5% of maximum health, 15 s internal cooldown, Will of the Forsaken removes fears, Read Leyline 50% / 15 s / 2 min, Wind Blessed 2% haste. Race order against round 3 race sims (re-runs still wait on the final re-rank).
- Stat paragraphs: no ranking words; every reason true (missed Arcane Blast still stacks and costs, missed Frostbolt gives no chill and so no Fingers of Frost or Winter's Chill roll, missed Scorch adds no Improved Scorch stack; school power coverage; Arcane Mind and Ice Shards crit bonus; Shatter on a Fingers of Frost charge). "Precision gives you 6%, your gear finds the rest" agrees with bis hit_to_cap (baseline 6 to 8, spell remainder).
- No ranking of stats or gear by a nightly value outside the tables. Race prose gives the measured order (allowed). "Greater Arcane Elixir is your elixir" is a rule: Elixir of Greater Firepower is not in Forever (round 1 O1; items.json has only its recipe, 21546 is Elixir of Holy Power).
- Card notes, all three APL files: nothing false (Fire Blast the worst damage per mana of Fire's boss spells, Arcane Power off the global cooldown, Ice Barrier 30 s, Cone of Cold 10 or 12 yards, Frostbolt a fraction of a stacked Blast).
- No em or en dashes, no site/sim words in body prose, no quoted nightly numbers.

Not re-filed (fixed, sim round, accepted or with the data or component lane): MAGE-1 to 16, MAGE-R2-1 to R2-8, MAGE-R3-1 to R3-10, MAGE-R4-1 to R4-4, MAGE-R5-1 to R5-5, MAGE-R6-1 to R6-5, MAGE-R7-1 to R7-8, SIM-1 to SIM-6, SIM-R2-1 to R2-3, SIM-R7-1, round 6 S1/S2 and round 7 S1/S2, the Arcane and Frost ladder-first orders, the Accelerant/Impact toll points, the Frost Nova into Cone of Cold order, component-rendered parts and the round 7 and round 8 component lists.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| MAGE-R8-1 | LIKELY (low) | leveling | fire.md:116 | "Keep Frost Nova and Blink for anything that reaches you, and Frostbolt for kiting: Fire has no slow of its own." | From 40 a Fire mage trains Frostfire Bolt, a Fire skill-line spell (trainables/mage.json: skill_line "Fire", ranks 40/50/60) whose client text is "slowing movement speed by 40%" (spelltips 1237313), the same slow as Frostbolt, and Heating Up names it ("Non-periodic critical strikes with Fireball, Frostfire Bolt, Fire Blast, and Scorch"). "No slow of its own" stops being true at 40. Frostbolt stays the cheaper kiting spell, so the advice holds; only the reason is wrong. | "Keep Frost Nova and Blink for anything that reaches you, and Frostbolt for kiting: none of your Fire spells slow until Frostfire Bolt at 40, and Frostbolt is still the cheaper slow after that." |

Counts: CONFIRMED 0, LIKELY 1 (R8-1, low), NEEDS-SIM 0, NEEDS-OWNER 0.

The three worst: R8-1 (Fire "has no slow of its own" while Frostfire Bolt, trained at 40 in the Fire line, slows 40%) is the only finding. Nothing else on the four pages or the card notes is false on 70755.

No findings for: build legality and point counts; every leveling level and the order the two blocks read in; Levels 1 to 19 spells, talents, wands and quests; every spell level, rank, cost, cast time, cooldown and value at 60; every talent description quoted; the Arcane loop, opener, Arcane Power windows, mana plan, AoE, threat and movement; Fire's opener, loop, Heating Up, Combustion, Ignite, threat, mana, AoE, toll points and Frostfire Bolt line; Frost's Ice Barrier, Fingers of Frost, Ice Lance, Shatter, Winter's Chill, Cone of Cold, Ice Block, mana, AoE and Frostfire Bolt line; stat paragraphs; enchants and consumables; race paragraphs; professions; the index page; card notes.

## SHOULDs (wording, not findings)

- fire.md:78 "Keep Fire Blast on cooldown even when you are short: every cast is a Heating Up stack": Heating Up adds the stack "every time it deals damage", so a missed Fire Blast adds none. "every hit is a Heating Up stack" matches the client text and the card note (mage-fire.json [5] "every hit").
- frost.md:42 "Cone of Cold and Winter's Chill want you in melee range": Winter's Chill stacks from Frostbolt at 30 yards; only its payoff, Cone of Cold, needs 10 yards. Clearer: "Cone of Cold, and with it Winter's Chill, wants you within 10 yards, where most bosses punish you."
- arcane.md:76 and fire.md:78 "a Mana Ruby when you are 1,200 mana down and a Major Mana Potion": reads as if both go at 1,200 down. A potion restores far more than the Ruby's 1,100, so naming "and a Major Mana Potion once you are down by more than it gives back" would stop a reader wasting it.
- index.md:32 "New spells ... are new talents" omits Accelerant (new Fire tier 1 talent, Fireball, Frostfire Bolt and Arcane Wrath cast time). The list does not claim to be complete, so nothing is false.

## Links checked

- Blizzard, Deep Dive panel recap (all four pages): 200, "World of Warcraft: Forever Deep Dive Panel Recap". Label right.
- Wowhead, Skyborne First Look (index.md): 200, "Skyborne First Look: New Neutral Race in World of Warcraft: Forever". Label right.
- Power Up Gaming, Forever talent changes vs Classic (index.md): 200, "New WoW Forever Talents: Changes vs Classic Guide". Label right.
- Talents Forever, talent data from the beta client (all four pages): 200, data.json. Label right.
- Forever Sixty, stat weights (three specs): https://foreversixty.gg/sim/weights, 200. Damage specs only.
- Forever Sixty, which boss drops are upgrades (three specs): https://foreversixty.gg/sim/drops, 200.
- /planner?class=mage (body, three specs): 200.
- All entries follow the round 4 source-link rule.

## Decisions

- **D1. A "no X of its own" claim counts every spell the spec trains, not only its tree.** Frostfire Bolt is in the Fire skill line and Heating Up names it, so Fire has a slow from 40; the leveling line keeps Frostbolt as the kiting spell for its cost, not for being the only slow (R8-1).
- **D2. Touch of the Grave stays at 30% on mage pages.** The shared racial tooltip shows the 20% row (1260189, warrior, paladin, rogue); the mage row 1260201 is 30% (engine racials.go:377, Wowhead Forever tooltip cited in round 3). No change.

## Sims requested

- None new. Carried unchanged: round 7 S1 (Fire Scorch rank 7 against 6 and 5, 180 and 300 s) and S2 (Arcane Frostbolt rank 10 against 6 and 4); round 6 S1 (Skyborne Arcane with Read Leyline cast) and S2 (mage race re-runs on the final re-ranked gear); SIM-R2-1 (Arcane Wrath), SIM-R2-2 (Arcane 32/16/3 Fire subspec), SIM-R2-3 (Fire leveling Scorch vs Fireball at 25 and 30), SIM-5/SIM-6 (Frost Accelerant point; Fire 60 toll), SIM-2 (Arcane Power at about +3 s), the Arcane and Frost ladder-first rows. The prose is true whatever these find.

## Sim findings

- None new. SIM-R7-1 (Arcane Subtlety's 15 spell penetration not modeled) carried.

## Sources

- data/builds/1.60.1.70755/trainables/mage.json (Frostfire Bolt skill line Fire, ranks 401502/1237312/1237313 at 40/50/60; every trainer level above), spelltips/mage.json (1237313 Frostfire Bolt, 10054 Mana Ruby 1,100, 10053 Citrine 850, 10230 Frost Nova, 13033 Ice Barrier, 10187 Blizzard, 10202 Arcane Explosion, 10216 Flamestrike, 10207 Scorch, 10199 Fire Blast, 18809 Pyroblast, 22783 Mage Armor, 10220 Ice Armor, 12051 Evocation, 11958 Ice Block, 12472 Cold Snap), spellconst/mage.json (Frostbolt ranks 116 to 25304, Frostfire Bolt ranks), talents/mage.json, spelltips/shared.json (racials), races.json, items.json, simconsumes.json, bis/mage-*.json (band 60 weights, hit_to_cap), addon-builds.json
- Engine d1d9f2775: sim/core/racials.go:376-379 (Touch of the Grave rows), sim/mage/armors.go:118 and sim/mage/talents.go:77 (casting regen additive)
- data/curated/apl/mage-arcane.json, mage-fire.json, mage-frost.json (priorityList notes)
- Blizzard, World of Warcraft: Forever Deep Dive panel recap: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- Wowhead, Skyborne First Look: https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
- Power Up Gaming, New WoW Forever talents changes vs Classic: https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/
- Talents Forever data export: https://talentsforever.com/data.json
- Forever Sixty, stat weights https://foreversixty.gg/sim/weights and Droptimizer https://foreversixty.gg/sim/drops

## Fix log (round 8)

- MAGE-R8-1 fixed, fire.md:116: "none of your Fire spells slow until Frostfire Bolt at 40, and Frostbolt is still the cheaper slow after that." Verified: trainables/mage.json Frostfire Bolt skill_line Fire, levels 40/50/60, 370 mana; spelltips 401502 to 1237313 "slowing movement speed by 40%".
- SHOULDs left: wording only, outside the assigned finding. Prose check: mage clean. No new SIM findings.
