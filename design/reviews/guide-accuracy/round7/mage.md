# Guide accuracy round 7: Mage

Audited on the round 7 audit tree, land/guides-r6 head 394562f20 (all six rounds applied). Client 1.60.1.70755, engine pin d1d9f2775 (sim/enginever/version.go).
Files: web/src/content/guides/mage/index.md, arcane.md, fire.md, frost.md, read line by line, plus every priorityList note in data/curated/apl/mage-{arcane,fire,frost}.json.
Evidence: data/builds/1.60.1.70755/{talents,trainables,spelltips,spellconst}/mage.json, gametables/playerexpectedstat.csv, races.json, addon-builds.json; engine d1d9f2775 sim/core/base_stats.go, base_stats_levels_auto_gen.go, spell.go (levelPenalty), spell_resistances.go, sim/mage/ice_lance.go, talents.go; rounds 1 to 6 mage reports, round4/mage-levels.md, race-sims.md rounds 2 to 5, ledger.md; live fetches of every source link.

Diff since round 6's audit tree (367cbadb7): only the round 6 fixes in mage prose (arcane.md:86, :92, :96; fire.md:64, :88, :94, :98; frost.md:89, :99) and the `updated:` stamps. No change to the builds, addon-builds.json, talents, trainables, the mage BiS files or the mage APL files.

## Builds decoded (unchanged, re-checked)

Decoded by tier/column against talents/mage.json; all legal (51 points, tier locks, prerequisites).

- **Arcane leveling** `141005213100301531/4/05350003` (31/4/16): Wand Specialization 1, Improved Channeling 4, Arcane Resilience 1, Arcane Concentration 5, Missile Barrage 2, Arcane Geometry 1, Arcane Impact 3, Arcane Blast 1, Arcane Meditation 3, Presence of Mind 1, Arcane Mind 5, Arcane Instability 3, Arcane Power 1; Incineration 4; Improved Frostbolt 5, Precision 3, Ice Shards 5, Frost Channeling 3.
- **Arcane raid** `131205213100301531/3/05350003` (32/3/16): Improved Channeling 3, Arcane Subtlety 2, Incineration 3, rest as leveling.
- **Fire leveling** `2520150000002/51521001230133031/003` (17/31/3) and **raid** `2530050000002/50520201230133031/003` (raid: Arcane Resilience 3, no Magic Absorption, Accelerant or Impact, Burning Soul 2).
- **Frost (both)** `2300052/5102/253510130101301221` (12/8/31).

Shipped leveling order (addon-builds.json, index 0 = level 10) unchanged since round 5's decode. Every talent level in "Levels 1 to 19" and "Levels 20 to 60" of all three specs matches it, and the two blocks read in order without contradicting each other or the strip.

## What holds (checked this round)

- Round 6 fixes: arcane.md:92 (school power covers only its school's spells), arcane.md:96 / fire.md:98 (Skyborne named as paid, Read Leyline outside the measured claim), fire.md:64 ("mana and hit"), fire.md:94, frost.md:99 ("best free Alliance", haste only): all true. The stat paragraphs carry no ranking words, and every talent they name is in the build (Precision 3, Arcane Mind 5, Critical Mass 3, Combustion, Ice Shards 5, Shatter 3). Two of the reasons they give are false; see R7-1 and R7-2.
- Stat reasons checked against 70755: a missed Arcane Blast still adds a stack and raises the next cost (client text "Each time you cast", engine adds the stack outside the hit branch); a missed Frostbolt applies no chill, so no Fingers of Frost roll ("Gives your Chill effects a 15% chance"); a missed Scorch adds no Improved Scorch stack; Arcane power covers Blast and Missiles; Fire power is all of Fire's boss damage; Arcane Mind "+100% critical strike damage bonus of your Arcane spells"; Ice Shards "+100% ... of your Frost spells"; Shatter 50% on Frozen targets, which a Fingers of Frost charge counts as; Ice Lance takes spell power (engine ice_lance.go BonusCoefficient from the generator, the client column being 0).
- No ranking of stats or gear by a nightly value anywhere outside the stat table: "Greater Arcane Elixir is your elixir" and the Elixir of Frost Power line are school rules, Arcane Subtlety "worth less than the crit" is a talent trade (and in Classic-rule resist math boss level resistance cannot be penetrated, engine spell_resistances.go resistCoeff).
- Every cost, cast time, cooldown and level quoted (spellconst/spelltips at 60): Scorch r7 150 / 1.5 s, Fireball r11 395 / 3.5 s, Fire Blast r7 340 / 8 s, Pyroblast 6 s, Flamestrike 990 / 3 s / 8 s, Arcane Explosion 390, Blizzard 1,400 (1,190 with Frost Channeling), Cone of Cold 10 s, Ice Lance 160 / 8 s, Ice Barrier 30 s / 1 min, Arcane Power and Presence of Mind 3 min, Evocation 8 min +1500% 8 s, Combustion 3 min, Blink 20 yd / 15 s; trainer levels Remove Lesser Curse 18, Blink and Evocation 20, Counterspell 24, Ice Armor 30, Mage Armor 34, Citrine 48, Ruby 58, Arcane Brilliance 56.
- Talent text quoted in all three pages (Improved Channeling 56% for Blast and 80% for Missiles at 4/5, Incineration covers Fire Blast, Ice Lance, Arcane Blast and Scorch, Missile Barrage 40% / 20%, Winter's Chill 20% a rank on Frostbolt and Fingers of Frost Ice Lances, Burning Soul 47% / 20% at 2/3, Arcane Meditation a third at 2/3).
- Race text against races.json (Ruthless 1% crit with all spells, Blood Fury includes spell power, Expansive Mind 5% mana, Touch of the Grave at most once every 15 s, Will of the Forsaken removes fears) and round 3 race sims.
- Card notes other than R7-5: nothing false (Arcane Power and Combustion are off the global cooldown, gcd_ms 0; Cone of Cold +35% spell damage is 5 x 7%).
- No em or en dashes, no site/sim words in body prose, no quoted nightly numbers.

Not re-filed (fixed, sim round, accepted or with the data or component lane): MAGE-1 to 16, MAGE-R2-1 to R2-8, MAGE-R3-1 to R3-10, MAGE-R4-1 to R4-4, MAGE-R5-1 to R5-5, MAGE-R6-1 to R6-5, SIM-1 to SIM-6, SIM-R2-1 to R2-3, round 6 S1/S2 (race re-runs wait on the final re-rank), the Arcane and Frost ladder-first orders, the Accelerant/Impact toll points, the Frost Nova into Cone of Cold order (round 1 MAGE-10), component-rendered parts and the round 7 component list.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| MAGE-R7-1 | CONFIRMED | stats (round 6 edit) | arcane.md:86 | "so for you Intellect is damage as well as your mana bar, where a Fire or Frost mage only drinks it as mana." | Intellect gives every mage spell crit on 70755: gametables/playerexpectedstat.csv, ClassID 8 level 60, SpellCritPerIntellect 0.000168 (0.0168% crit a point, about 60 Intellect per 1%); the engine wires it the same way (sim/core/base_stats.go:213 AddStatDependency(Intellect, SpellCrit, SpellCritPerIntAtLevel), mage 60 = 0.0168). A Fire or Frost mage gets crit from Intellect, not only mana. | "Each Arcane Blast stack adds 12% of your Intellect as spell damage and Arcane Mind adds 10% more Intellect, so for you Intellect is spell damage on top of the spell crit and the mana bar every mage gets from it." |
| MAGE-R7-2 | CONFIRMED | stats (round 6 edit) | frost.md:89 | "Intellect is only your mana bar, and Frost is the spec that does not run out." | Same client row: Intellect adds spell crit at 60 (0.0168% a point), and crit is the stat the same paragraph says Frost's damage piles up on (Ice Shards, Shatter). | "Intellect is your mana bar and a little spell crit, and Frost is the spec that does not run out." |
| MAGE-R7-3 | LIKELY (low) | stats | fire.md:88 | "Intellect is your mana bar, and Fire drinks a Major Mana Potion on every boss, so a deeper bar keeps you casting to the end." | Not false, but it leaves out the reason that matters most to Fire: Intellect is spell crit (same client row), and the paragraph's own crit line says crit feeds Ignite, Heating Up and Master of Elements. The three spec pages should give Intellect the same lasting reasons. | "Intellect is your mana bar and adds spell crit, which Fire turns into Ignite and Heating Up; Fire drinks a Major Mana Potion on every boss, so a deeper bar keeps you casting to the end." |
| MAGE-R7-4 | CONFIRMED (low) | spell | fire.md:78 | "Never downrank Scorch; the top rank is already the cheapest damage you have." | spellconst/mage.json Scorch: r7 178 / 150 mana, r6 150 / 135, r5 116 / 115, all 1.5 s at coefficient 0.429. The engine keeps the full coefficient on every rank learned at 20 or later (sim/core/spell.go levelPenalty, levelPenaltyLevel 20; Classic rule). At 700 spell power that is 3.19 damage per mana for r7, 3.34 for r6, 3.62 for r5; Pyroblast (583 + 212 over 12 s, 440 mana) is cheaper damage still. The top rank is the most damage per cast, not the cheapest damage. Whether a lower rank is worth it for a Fire mage who runs dry is a sim question (S1). | Until S1: "Keep the top rank of Scorch: every rank is the same 1.5 second cast, so a lower rank saves a few mana and loses damage on every cast." If S1 finds a lower rank ahead beyond error, say which rank and when. |
| MAGE-R7-5 | CONFIRMED (low) | card note | data/curated/apl/mage-fire.json priorityList[6] | "At 150 mana for 1.5 seconds it does far more damage per mana than Fireball's 395 and its long cast" | Same tables: Scorch r7 (178 + 0.429 x SP) / 150 against Fireball r11 (451 + SP, plus 56 over 8 s) / 395. At 700 spell power, 3.19 against 3.06 a mana; at 400, 2.33 against 2.29. Scorch is a few percent ahead per mana, not far ahead. Its real edge is cast count: two Scorches in one Fireball's time, each a Heating Up and Ignite chance and a stack refresh (the measured 10 to 12% DPS gain the APL notes record). | "Scorch is the filler, not Fireball. Mana for damage is close to even, but you land two Scorches in the time of one Fireball, so you roll twice the crits for Heating Up and Ignite and keep your stacks rolling." |
| MAGE-R7-6 | LIKELY (low) | leveling | frost.md:117 | "After a Frost Nova or a Frostbite freeze, Ice Lance hits a mob for four times its damage, and once Shatter is in by 34 it crits." | Shatter 3/3 is "+50% critical strike chance ... against Frozen targets" (talents/mage.json), not a sure crit; at 34 a leveling Frost mage has no other crit talent (Incineration comes after 40) and a few percent of base and Intellect crit, so the shatter lands a crit a little more than half the time. frost.md:55 rightly says "most of the time" at 60. | "...Ice Lance hits a mob for four times its damage, and once Shatter is in by 34 it crits more often than not." |
| MAGE-R7-7 | LIKELY (low) | voice | index.md:31 | "...which is why every build, Arcane included, dips three points there." | Precision sits in the Frost tree; the Frost build is the Frost tree, so for it the three points are not a dip. True for Arcane and Fire only. | "...which is why the Arcane and Fire builds both dip three points there." |
| MAGE-R7-8 | NEEDS-SIM | rotation | arcane.md:76 (and frost.md:75 for leveling) | "Do not downrank Frostbolt to stretch mana; the cheaper ranks lose more damage than the mana buys." | Unmeasured and against the client numbers for a mana-bound spec: Frostbolt r10 386 / 260 mana, r6 134 / 130, r4 70 / 65 (2.6 s, coefficient 0.706, full value: spell.go levelPenalty spares ranks learned at 20+). At 700 spell power, 3.68 damage a mana for r10, 5.41 for r6, 8.68 for r4. The page says "You will run dry"; a cheap filler between Blast sets could be a gain or a loss, and only a sim settles it. The Frost line is true at 60 (Frost does not run out), so it stays. | Wait for S2. If rank 10 holds beyond error, keep the sentence. If a lower rank wins, rewrite arcane.md:76 and card notes mage-arcane.json [3], [4], [7] to name it. |

Counts: CONFIRMED 4 (R7-1, R7-2; R7-4, R7-5 low), LIKELY 3 (R7-3, R7-6, R7-7, all low), NEEDS-SIM 1 (R7-8), NEEDS-OWNER 0.

The three worst: R7-1 and R7-2 (round 6's new stat prose says Intellect is only mana for Fire and Frost, when the client gives every mage spell crit from it), R7-4 (the top Scorch rank called the cheapest damage, false on the client's own rank values).

No findings for: build legality and point counts; every leveling level and the order the two blocks read in; Levels 1 to 19 spells, talents, wands and quests; spell levels, ranks, costs, cast times and cooldowns; every talent description quoted; the Arcane loop, opener, Arcane Power windows, mana plan, threat and movement; Fire's opener, Heating Up, Combustion, Ignite, toll points and Frostfire Bolt line; Frost's Ice Barrier, Fingers of Frost, Ice Lance, Shatter at 60, Winter's Chill, Cone of Cold, Ice Block and threat; race paragraphs; professions; card notes other than R7-5.

## Links checked

- Blizzard, Deep Dive panel recap (all four pages): 200, "World of Warcraft: Forever Deep Dive Panel Recap". Label right.
- Wowhead, Skyborne First Look (index.md): 200, "Skyborne First Look: New Neutral Race in World of Warcraft: Forever". Label right.
- Power Up Gaming, Forever talent changes vs Classic (index.md): 200, "New WoW Forever Talents: Changes vs Classic Guide". Label right.
- Talents Forever, talent data from the beta client (all four pages): 200, data.json. Label right (round 6 read the `_readme`).
- Forever Sixty, stat weights (three specs): https://foreversixty.gg/sim/weights, 200. Damage specs only, per the round 4 rule.
- Forever Sixty, which boss drops are upgrades (three specs): https://foreversixty.gg/sim/drops, 200.
- /planner?class=mage (body, three specs): 200.
- All entries follow the round 4 source-link rule.

## Decisions

- **D1. Intellect's reason names crit on every mage page.** The 70755 gametable gives mages spell crit from Intellect at every level, so no page may call Intellect mana only. Arcane adds the Blast stacks' spell damage on top (R7-1 to R7-3).
- **D2. Rank advice states only what the client numbers show.** Lower Scorch and Frostbolt ranks keep their full coefficient in Forever (Classic's rule: only ranks below level 20 lose spell power), so they are cheaper damage per mana. The pages keep the top rank as the recommendation and give its true reason (more damage per cast) until S1 and S2 measure the trade (R7-4, R7-8).

## Sims requested

- **S1 (Fire Scorch rank).** mage-fire, raid preset, band 60, guide raidBuild, published gear, 10,000 iterations, seeds 1 and 50001: the APL's Scorch lines (stack build [3] and filler [6]) at rank 7 (10207) against rank 6 (10206) and rank 5 (10205), everything else unchanged. Also a 300 second fight, where mana binds harder. If a lower rank leads beyond error on either length, fire.md:78 names it; otherwise apply R7-4's text.
- **S2 (Arcane Frostbolt rank).** mage-arcane, same method: the Frostbolt in the filler ([7]) and Presence of Mind ([4]) lines at rank 10 (10181) against rank 6 (8407) and rank 4 (7322), on 180 and 300 second fights. If rank 10 holds, arcane.md:76 stands; if a lower rank wins beyond error, rewrite arcane.md:76 and the card notes.
- Carried unchanged: round 6 S1 (Skyborne Arcane with Read Leyline) and S2 (mage race re-runs on the re-ranked gear, after the final best in slot re-rank), SIM-R2-1 (Arcane Wrath), SIM-R2-2 (Arcane 32/16/3 Fire subspec), SIM-R2-3 (Fire leveling Scorch vs Fireball at 25 and 30), SIM-5/SIM-6 (Frost Accelerant point; Fire 60 toll), SIM-2 (Arcane Power at about +3 s), the Arcane and Frost ladder-first rows.

## Sim findings

- **SIM-R7-1 (low).** The engine models only the threat half of Arcane Subtlety (sim/mage/talents.go:710, SpellMod_Threat_Pct); the "reduces your target's resistance to all your spells by 15" half adds no SpellPenetration. It costs nothing on a boss with no resistance to your school (and level-based resistance cannot be penetrated, spell_resistances.go resistCoeff), but on a target with Arcane, Fire or Frost resistance the raid build's two points are under-valued. Add the 15 points of spell penetration when the talent is taken.

## Sources

- data/builds/1.60.1.70755/gametables/playerexpectedstat.csv (ClassID 8: SpellCritPerIntellect 0.000593 at 20, 0.000276 at 40, 0.000168 at 60)
- data/builds/1.60.1.70755/spellconst/mage.json and spelltips/mage.json (Scorch 2948 to 10207, Fireball 10151, Pyroblast 18809, Frostbolt 116 to 25304, Fire Blast 10199, Ice Lance 1240047, Ice Barrier 13033, Arcane Power 12042, Presence of Mind 12043, Combustion 11129, Evocation 12051, Blizzard 10187, Cone of Cold 10161, Arcane Brilliance 23028)
- data/builds/1.60.1.70755/talents/mage.json (Shatter, Ice Shards, Arcane Mind, Precision, Incineration, Improved Channeling, Winter's Chill, Fingers of Frost), trainables/mage.json, races.json, addon-builds.json
- data/curated/apl/mage-arcane.json, mage-fire.json, mage-frost.json (priorityList notes)
- Engine d1d9f2775: sim/core/base_stats.go:200-215 and base_stats_levels_auto_gen.go (spellCritPerIntByClassLevel), sim/core/spell.go:777-820 (levelPenalty), sim/core/spell_resistances.go (resistCoeff), sim/mage/ice_lance.go, sim/mage/talents.go:710
- Blizzard, World of Warcraft: Forever Deep Dive panel recap: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- Wowhead, Skyborne First Look: https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
- Power Up Gaming, New WoW Forever talents changes vs Classic: https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/
- Talents Forever data export: https://talentsforever.com/data.json
- Forever Sixty, stat weights https://foreversixty.gg/sim/weights and Droptimizer https://foreversixty.gg/sim/drops

## Fix log (round 7)

- Stat priority rule: arcane.md:86, fire.md:88, frost.md:89 already carried no ranking words (they point to the table and give lasting reasons); only the Intellect reasons changed. Prose test clean for mage (`-t mage`, 1 passed).
- MAGE-R7-1 fixed, arcane.md:86: Intellect is spell damage on top of the spell crit and mana bar every mage gets from it (D1).
- MAGE-R7-2 fixed, frost.md:89: "Intellect is your mana bar and a little spell crit" (D1).
- MAGE-R7-3 fixed, fire.md:88: Intellect is mana and spell crit, which Fire turns into Ignite and Heating Up; no ranking (D1).
- MAGE-R7-4 fixed, fire.md:78: "cheapest damage" and the downrank claim dropped; now the lasting reason only (two Scorches in one Fireball's time, each a Heating Up and Ignite roll). No rank claim either way (D2, integrator).
- MAGE-R7-5 fixed, data/curated/apl/mage-fire.json priorityList[6] notes: "far more damage per mana" dropped; two casts per Fireball for Heating Up and Ignite plus stack refresh. Action unchanged.
- MAGE-R7-6 fixed, frost.md:117: Shatter at 34 "crits more often than not" (+50% on Frozen, not a sure crit).
- MAGE-R7-7 fixed, index.md:31: "the Arcane and Fire builds both dip three points there"; index updated: 2026-11-04.
- MAGE-R7-8 fixed per integrator (no downrank claim either way): arcane.md:76 sentence removed; frost.md:75 "loses more damage than it saves at every level" replaced by "Drink before pulls so you start each one full." (same unmeasured claim). The rank-1 Frostbolt slow line stays. Arcane card notes [3], [4], [7] make no rank claim; unchanged.
- D1, D2 applied as above. Other pages: no stat or gear ranking by nightly value found outside the tables (race prose gives the measured order, allowed).
- SIM findings: S1 (Scorch rank) and S2 (Frostbolt rank, Arcane and Frost) still not run; prose now true whatever they find. SIM-R7-1 (Arcane Subtlety penetration not modeled) carried. No new SIM findings.
