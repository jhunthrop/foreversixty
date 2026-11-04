# Guide accuracy round 9: Mage

Audited on the round 9 audit tree, land/guides-r8 head 645caa983 (all eight rounds applied). Client 1.60.1.70755, engine pin d1d9f2775 (sim/enginever/version.go).
Files: web/src/content/guides/mage/index.md, arcane.md, fire.md, frost.md, read line by line (frontmatter, body), plus every priorityList note in data/curated/apl/mage-{arcane,fire,frost}.json.
Evidence: data/builds/1.60.1.70755/{talents,spelltips,spellranks,trainables}, spelltips/shared.json, races.json, addon-builds.json; engine d1d9f2775 sim/core/eureka.go, eureka_auto_gen.go (mage row 1259817), racials.go:432, spell_outcome.go:1080 (critMultiplierAs), sim/mage/talents.go (Arcane Mind, Ice Shards), heating_up.go; Wowhead Forever tooltip for 1259817; rounds 1 to 8 mage reports, round4/mage-levels.md, race-sims rounds 2 to 5, ledger.md; live fetches of every source link.

Diff since round 8's audit tree (09b98d9e4): for mages only fire.md:116 (the R8-1 fix). No change to builds, addon-builds.json, talents, trainables, BiS, race data or the mage APL files.

## Builds decoded (unchanged, re-checked)

Decoded by tier/column against talents/mage.json; all legal (51 points, tier locks, prerequisites).

- **Arcane leveling** `141005213100301531/4/05350003` (31/4/16) and **raid** `131205213100301531/3/05350003` (32/3/16: Improved Channeling 3, Arcane Subtlety 2, Incineration 3).
- **Fire leveling** `2520150000002/51521001230133031/003` (17/31/3: Accelerant 1, Impact 1) and **raid** `2530050000002/50520201230133031/003` (Arcane Resilience 3, no Magic Absorption, Burning Soul 2). Fire damage talents 28 + toll 2 = 30 above Combustion.
- **Frost (both)** `2300052/5102/253510130101301221` (12/8/31). Frost 28 + Winter's Chill 2 = 30 above Ice Barrier.

Shipped leveling order (addon-builds.json, entries [tree, tier, column], index 0 = level 10), re-decoded this round:
- Arcane: Wand Spec 10, Improved Channeling 11-14, Arcane Resilience 15, Arcane Concentration 16-20, Missile Barrage 21-22, Arcane Geometry 23, Arcane Impact 24-26, Arcane Blast 27, Arcane Meditation 28-30, Presence of Mind 31, Arcane Mind 32-36, Arcane Instability 37-39, Arcane Power 40, Incineration 41-44, Improved Frostbolt 45-49, Precision 50-52, Ice Shards 53-57, Frost Channeling 58-60.
- Fire: Incineration 10-14, Accelerant 15, Ignite 16-19 and 21, Pyroblast 20, Flame Throwing 22-23, Impact 24, Improved Fire Blast 25-26, Improved Scorch 27-29, Heating Up 30, Master of Elements 31-33, Critical Mass 34-36, Fire Power 37-39, Combustion 40, then Arcane and Precision 58-60.
- Frost: Frost Warding 10-11, Improved Frostbolt 12-16, Precision 17-19, Ice Shards 20-24, Frostbite 25, Permafrost 26, Frost Channeling 27-29, Ice Lance 30, Ice Block 31, Shatter 32-34, Cold Snap 35, Fingers of Frost 36-37, Winter's Chill 38-39, Ice Barrier 40.

Every level in "Levels 1 to 19" and "Levels 20 to 60" of all three specs matches it; the two blocks read in order and agree with the strip.

## Round 9 focus checks

- **Crit damage bonus.** The engine's spell crit is `1 + (1.5 - 1) x CritDamageBonus` (spell_outcome.go:1080); Arcane Mind 5/5 and Ice Shards 5/5 add 1.0 to the bonus multiplier (talents.go, 0.20 a rank, client "increases the critical strike damage bonus ... by 100%"), so a crit goes from 150% to 200% of a hit: the +50% bonus doubles. arcane.md:58 "Arcane Mind also doubles the bonus on your Arcane crits", arcane.md:86 "Arcane Mind doubles their crit bonus", frost.md:56 "doubles the bonus of your Frost crits" and frost.md:89 "Ice Shards doubles the crit bonus of your Frost spells" are all right. No page quotes a crit-damage percentage. arcane.md:59 "full crit damage" names no number.
- **Eureka!** The mage row is 1259817 (spellranks.json classes.mage; engine racials.go:433). Wowhead Forever tooltip: "Your damaging abilities have their Mana cost reduced by 10% and deal 10% more damage for 8 sec." The shared tip's "non-periodic damaging abilities ... Energy cost" is the rogue row. The engine's mage row (eureka_auto_gen.go) puts every mage damage spell in the damage and cost masks (Frostbolt 10181/25304, Ice Lance 1240047, Arcane Missiles 10212/25345, Arcane Blast 1239700, Scorch 10207, Fire Blast 10199, Fireball 10151, Pyroblast 18809, Flamestrike 10216, Blizzard 10187, Cone of Cold 10161, Arcane Explosion 10202, Frostfire Bolt 1237313) and leaves the periodic mask empty. No utility spell (Polymorph, Ice Barrier, Arcane Brilliance, Blink, Evocation, conjures) is in the cost mask. So "10% more damage" holds for every mage damage spell's hit. Pyroblast's, Fireball's and Flamestrike's burns do not get it, but Ignite does because it grows with the crit. "10% cheaper spells" overstates: only the damage spells get cheaper. See R9-1.

## Round 8 fix re-checked

- fire.md:116 "none of your Fire spells slow until Frostfire Bolt at 40, and Frostbolt is still the cheaper slow after that": Frostfire Bolt is Fire skill line, 40/50/60, 370 mana at 60, "slowing movement speed by 40%"; Frostbolt r11 290 mana. True.

## What holds (checked this round)

- Talent text on all four pages against talents/mage.json rank text (every line in the Talents and builds lists, the build paragraphs and the leveling prose), including Impact 3% at 1/3, Burning Soul 47% at 2/3, Arcane Meditation 33% at 2/3, Accelerant 0.1 s at 1/5, Arcane Geometry 3 yd at 1/2, Arcane Resilience 25% at 1/3, Frost Warding 30% at 2/2, Winter's Chill 40% at 2/5, Improved Channeling 56% Blast / 80% Missiles at 4/5, Heating Up 25% a stack (engine heating_up.go subtracts 0.25 a stack from the cast multiplier, so 3 stacks make a 6 s Pyroblast 1.5 s).
- Spell values at 60: Arcane Blast r5 1.5 s (so Presence of Mind on it saves nothing), Frostbolt r10 260 / r11 290 mana, Frostfire Bolt 292 + 57 over 9 s and 370 mana against Frostbolt r11's 475 (frost.md:85 "hits for less ... costs more" true; 2.9 s with Frost's one Accelerant point, "close to a 3 second cast"), Blizzard 1,400 (1,190 with Frost Channeling, "about 1,200"), Arcane Explosion 390 mana at 10 yd, Flamestrike 990 / 3 s / 8 s burn, Ice Barrier 30 s cooldown and no pushback while shielded, Ice Lance 8 s and +300% on Frozen, Arcane Power, Presence of Mind and Combustion 180 s, Evocation 480 s / +1500% / 8 s, Cold Snap 10 min, Blink 20 yd / 15 s, Arcane Brilliance "party and raid", Mage Armor 50% casting regeneration.
- Trainer levels: Fireball 1/6/12/18, Frostbolt 4/8/14/20, Arcane Missiles 8, Scorch 22, Cone of Cold 26, Frostfire Bolt 40/50/60, Arcane Missiles r7 at 56 with r8 the book (card note).
- Index page: Ice Block at 15 points with Cold Snap below it, Precision in Frost for every school, the new-talent list, Orc and Skyborne lines, race picks against round 3 race sims.
- Mana plan (gem and rune 2 minute shared cooldown, Robe 5 minutes, one minute lockout) per round 1 O4. Race, stat, enchant, profession paragraphs unchanged and true.
- Card notes in all three APL files: nothing false (Cone of Cold on 5 Winter's Chill stacks "double damage and 35% more of your spell damage", Heating Up feeders, Combustion "10% crit until three non-periodic crits", Missile Barrage 40% / 20%).
- No em or en dashes, no site/sim words in body prose, no quoted nightly numbers, no stat or gear ranking outside the tables.

Not re-filed (fixed, sim round, accepted or with the data or component lane): MAGE-1 to 16, MAGE-R2-1 to R2-8, MAGE-R3-1 to R3-10, MAGE-R4-1 to R4-4, MAGE-R5-1 to R5-5, MAGE-R6-1 to R6-5, MAGE-R7-1 to R7-8, MAGE-R8-1, SIM-1 to SIM-6, SIM-R2-1 to R2-3, SIM-R7-1, round 6 S1/S2 and round 7 S1/S2, the Arcane and Frost ladder-first orders, the Accelerant/Impact toll points, the Frost Nova into Cone of Cold order, component-rendered parts and the round 6 to 8 component lists.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| MAGE-R9-1 | CONFIRMED (low) | race | index.md:38, arcane.md:96, fire.md:98, frost.md:99 | index.md:38 "Eureka! makes your spells 10% stronger and 10% cheaper for 8 seconds"; arcane.md:96, fire.md:98, frost.md:99 "Eureka! is 10% more damage and 10% cheaper spells for 8 seconds" | Mage row 1259817 (Wowhead Forever tooltip): "Your damaging abilities have their Mana cost reduced by 10% and deal 10% more damage for 8 sec." Engine eureka_auto_gen.go mage row: the cost and damage masks hold only damage spells, and the periodic mask is empty. Polymorph, Ice Barrier, Arcane Brilliance and Blink do not get cheaper. On Fire the burns of Pyroblast, Fireball and Flamestrike do not get the 10%. | index.md:38: "Eureka! makes your damage spells hit 10% harder and cost 10% less for 8 seconds, every 2 minutes." arcane.md:96, fire.md:98, frost.md:99: "Eureka! makes your damage spells hit 10% harder and cost 10% less for 8 seconds on a 2 minute cooldown". |

Counts: CONFIRMED 1 (R9-1, low), LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

The three worst: R9-1 (Eureka! called 10% cheaper "spells" when the mage row reaches only damage spells, and only their hits) is the only finding. Every crit-damage claim is right on the engine's 150% spell crit base, and nothing else on the four pages or in the card notes is false on 70755.

No findings for: build legality and point counts; every leveling level and the order the two blocks read in; Levels 1 to 19 spells, talents, wands and quests; every spell level, rank, cost, cast time, cooldown and value at 60; every talent description quoted; crit-damage claims; the Arcane loop, opener, Arcane Power windows, mana plan, AoE, threat and movement; Fire's opener, loop, Heating Up, Combustion, Ignite, threat, mana, AoE, toll points and Frostfire Bolt line; Frost's Ice Barrier, Fingers of Frost, Ice Lance, Shatter, Winter's Chill, Cone of Cold, Ice Block, mana, AoE and Frostfire Bolt line; stat paragraphs; enchants and consumables; other race lines; professions; card notes.

## SHOULDs (wording, not findings)

- Carried from round 8, still unapplied: fire.md:78 "every cast is a Heating Up stack" would match the client ("every time it deals damage") and card note [5] better as "every hit". frost.md:42 "Cone of Cold and Winter's Chill want you in melee range": Winter's Chill stacks from 30 yards and Cone of Cold reaches 10, which is not melee range. arcane.md:76 and fire.md:78 "a Mana Ruby when you are 1,200 mana down and a Major Mana Potion" reads as if both go at 1,200 down. index.md:32 leaves Accelerant out of the new-talent list.
- fire.md:56 "Improved Scorch, Ignite and Pyroblast's burn are three of them": with Eureka! up, note that the burns are not boosted. Not needed. This is only context for R9-1.

## Links checked

- Blizzard, Deep Dive panel recap (all four pages): 200. Label right.
- Wowhead, Skyborne First Look (index.md): 200. Label right.
- Power Up Gaming, Forever talent changes vs Classic (index.md): 200. Label right.
- Talents Forever, talent data from the beta client (all four pages): 200, data.json. Label right.
- Forever Sixty, stat weights and which boss drops are upgrades (three specs): 200 each.
- /planner?class=mage (body, three specs): 200.
- All entries follow the round 4 source-link rule.

## Decisions

- **D1. Eureka! on mage pages follows the mage row, not the shared tip.** The shared tooltip is the rogue row ("non-periodic ... Energy"). The mage row says "damaging abilities ... Mana cost", and the engine masks reach only damage spells and leave out their periodic parts. Pages say "damage spells" and "hit harder" (R9-1).
- **D2. Crit-damage talents on spells double the +50% bonus.** Arcane Mind and Ice Shards at 5/5 take a spell crit from 150% to 200% (engine critMultiplierAs, magic base 1.5). "Doubles the bonus" stays on all pages. Do not apply the melee +100% base reading to mage spells.

## Sims requested

- None new. Carried unchanged: round 7 S1 (Fire Scorch rank 7 against 6 and 5, 180 and 300 s) and S2 (Arcane Frostbolt rank 10 against 6 and 4); round 6 S1 (Skyborne Arcane with Read Leyline cast) and S2 (mage race re-runs on the final re-ranked gear); SIM-R2-1 (Arcane Wrath), SIM-R2-2 (Arcane 32/16/3 Fire subspec), SIM-R2-3 (Fire leveling Scorch vs Fireball at 25 and 30), SIM-5/SIM-6 (Frost Accelerant point; Fire 60 toll), SIM-2 (Arcane Power at about +3 s), the Arcane and Frost ladder-first rows. The prose is true whatever these find.

## Sim findings

- None new. SIM-R7-1 (Arcane Subtlety's 15 spell penetration not modeled) carried.

## Sources

- data/builds/1.60.1.70755/talents/mage.json, spelltips/mage.json (10181, 25304, 1237313, 1239700, 12042, 12043, 11129, 10207, 10151, 18809, 10212, 25345, 10187, 10202, 10216, 13033, 12051, 12472, 1953, 23028, 22783, 1240047), spellranks.json (mage ranks; Eureka! class rows), spelltips/shared.json (Eureka! rogue-row text), races.json, addon-builds.json
- Engine d1d9f2775: sim/core/eureka.go, sim/core/eureka_auto_gen.go (mage row 1259817), sim/core/racials.go:432-434, sim/core/spell_outcome.go:1080-1086, sim/mage/talents.go:79-85, 593-597, 650-652, 726-731, sim/mage/heating_up.go:25-58
- data/curated/apl/mage-arcane.json, mage-fire.json, mage-frost.json (priorityList notes)
- Wowhead Forever, Eureka! (mage, spell 1259817) tooltip: https://nether.wowhead.com/forever/tooltip/spell/1259817
- Blizzard, World of Warcraft: Forever Deep Dive panel recap: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- Wowhead, Skyborne First Look: https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
- Power Up Gaming, New WoW Forever talents changes vs Classic: https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/
- Talents Forever data export: https://talentsforever.com/data.json
- Forever Sixty, stat weights https://foreversixty.gg/sim/weights and Droptimizer https://foreversixty.gg/sim/drops

## Fix log (round 9)

Applied by the integrator; see the round 9 ledger entry and the commit for the exact lines.
