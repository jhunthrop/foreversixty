# Guide accuracy round 4: Mage

Audited on land/guides-r4 752c02e43 (audit tree), client 1.60.1.70755, engine pin d1d9f2775 (sim/enginever/version.go).
Files: web/src/content/guides/mage/index.md, arcane.md, fire.md, frost.md, read line by line, with the new "### Levels 1 to 19" subsections checked against round4/mage-levels.md.
Evidence: data/builds/1.60.1.70755/{talents,trainables,spelltips}/mage.json, addon-builds.json, items.json, loot.json (sources + quests), combos.json, races.json; round1 to round3 mage reports, round2/round3 race-sims.md, ledger.md; Wowhead Forever quest tooltips 93797 and 99156; live fetches of every frontmatter and body link.

## Builds decoded (unchanged since round 3, re-checked)

Decoded by tier/column against talents/mage.json; all legal (51 points, tier locks, prerequisites).

- **Arcane leveling** `141005213100301531/4/05350003`: Wand Specialization 1, Improved Channeling 4, Arcane Resilience 1, Arcane Concentration 5, Missile Barrage 2, Arcane Geometry 1, Arcane Impact 3, Arcane Blast 1, Arcane Meditation 3, Presence of Mind 1, Arcane Mind 5, Arcane Instability 3, Arcane Power 1; Incineration 4; Improved Frostbolt 5, Precision 3, Ice Shards 5, Frost Channeling 3.
- **Arcane raid** `131205213100301531/3/05350003`: as leveling but Improved Channeling 3, Arcane Subtlety 2, Incineration 3.
- **Fire leveling** `2520150000002/51521001230133031/003`: Wand Specialization 2, Improved Channeling 5, Arcane Resilience 2, Magic Absorption 1, Arcane Concentration 5, Arcane Meditation 2; Incineration 5, Accelerant 1, Ignite 5, Flame Throwing 2, Impact 1, Pyroblast 1, Improved Fire Blast 2, Improved Scorch 3, Heating Up 1, Master of Elements 3, Critical Mass 3, Fire Power 3, Combustion 1; Precision 3.
- **Fire raid** `2530050000002/50520201230133031/003`: as leveling but Arcane Resilience 3, no Magic Absorption, Accelerant or Impact, Burning Soul 2.
- **Frost (both)** `2300052/5102/253510130101301221`: Wand Specialization 2, Improved Channeling 3, Arcane Concentration 5, Missile Barrage 2; Incineration 5, Accelerant 1, Flame Throwing 2; Frost Warding 2, Improved Frostbolt 5, Precision 3, Ice Shards 5, Frostbite 1, Permafrost 1, Frost Channeling 3, Ice Lance 1, Ice Block 1, Shatter 3, Cold Snap 1, Fingers of Frost 2, Winter's Chill 2, Ice Barrier 1.

Shipped leveling order (addon-builds.json, index 0 = level 10), re-decoded:
- Arcane: 10 Wand Specialization, 11 to 14 Improved Channeling 4, 15 Arcane Resilience, 16 to 20 Arcane Concentration 5, 21 to 22 Missile Barrage, 23 Arcane Geometry, 24 to 26 Arcane Impact, 27 Arcane Blast, 28 to 30 Arcane Meditation, 31 Presence of Mind, 32 to 36 Arcane Mind, 37 to 39 Arcane Instability, 40 Arcane Power, 41 to 44 Incineration, 45 on Frost.
- Fire: 10 to 14 Incineration, 15 Accelerant, 16 to 19 Ignite 4, 20 Pyroblast, 21 Ignite 5, 22 to 23 Flame Throwing, 24 Impact, 25 to 26 Improved Fire Blast, 27 to 29 Improved Scorch, 30 Heating Up, 31 to 39 Master of Elements, Critical Mass, Fire Power, 40 Combustion.
- Frost: 10 to 11 Frost Warding, 12 to 16 Improved Frostbolt, 17 to 19 Precision, 20 to 24 Ice Shards, 25 Frostbite, 26 Permafrost, 27 to 29 Frost Channeling, 30 Ice Lance, 31 Ice Block, 32 to 34 Shatter, 35 Cold Snap, 36 to 37 Fingers of Frost, 38 to 39 Winter's Chill, 40 Ice Barrier.

Every Levels 1 to 19 talent level and every Levels 20 to 60 talent level matches this order; nothing in the new subsections contradicts the text below them.

Levels 1 to 19 data checks that hold (trainables/mage.json): Fireball 1/6/12/18, Frostbolt 4/8/14/20, Conjure Water 4, Fire Blast 6, Arcane Missiles 8, Frost Nova 10, Arcane Explosion 14, Flamestrike 16. Talent text: Wand Specialization 13%, Arcane Resilience 25% of Intellect, Incineration 10% at 5/5 (Fire Blast is the only listed spell trained before 20), Accelerant 0.1 sec, Frost Warding 30% Frost Armor armor at 2/2, Improved Frostbolt 0.5 sec, Precision 6%. Items: Lesser Magic Wand 11287 (req 5) and Greater Magic Wand 11288 (req 13) are loot.json `crafted:enchanting`; Elven Wand 5604 from Bashal'Aran (957, Alliance, level 13); Riptear's Spare Arm 286743 from Rear Guard Patrol (99156, Horde, level 13; Wowhead Forever tooltip: hand-in to Deathguard Linnea). Level 10 mage class quests tagged `classes: [mage]` in loot.json: Mirror Lake and Ju-Ju Heaps (Ley Staff 9513 or Ley Orb 7508), Mage-tastic Gizmonitor and The Balnir Farmstead (Arcane Staff 9514 or Arcane Orb 7507).

Round 3 fixes re-checked: R3-1 to R3-10 hold. The round 3 rendered-review lines (arcane.md:45, :83; index.md:29) are true: the Arcane builds put 16 points in Frost, enough for the tier that holds Ice Block, but every one of them is spent. Round 4's own edits other than those filed below hold.

Not re-filed (fixed, sim round or accepted): MAGE-1 to 16, MAGE-R2-1 to R2-8, MAGE-R3-1 to R3-10, SIM-1 to SIM-6, SIM-R2-1 (Arcane Wrath), SIM-R2-2 (Arcane Fire subspec), SIM-R2-3 (Fire leveling filler, fire.md:117), the Arcane and Frost ladder-first orders (Arcane Blast at 20; Improved Frostbolt before Frost Warding, which the new frost.md:113 now repeats from the shipped order), Accelerant/Impact toll points, component-rendered parts.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| MAGE-R4-1 | CONFIRMED | source | arcane.md:20-28, fire.md:20-28, frost.md:20-28 | `'Forever Sixty, stat weights'`, `'Forever Sixty, rotation'`, `'Forever Sixty, raid loot tables'`, all with `url: https://foreversixty.gg/sources` | The link resolves (200), but the page is "How sources are classified": its headings are The four pill kinds, What "single-source" means, When we get something wrong, Primary sources this site checks, Panel recaps and transcripts, Developer interviews, Community data we credit and reuse. It says nothing about stat weights, a rotation or raid loot, so none of the three labels says what its link holds. The pages that do exist and resolve: https://foreversixty.gg/sim/weights ("What one point of each stat is worth for your damage spec"), https://foreversixty.gg/bis/mage/arcane, /fire, /frost (200). Hunter and warrior already link /sim/weights. | In each spec file: stat weights entry `url: https://foreversixty.gg/sim/weights`, label `'Forever Sixty, what each stat is worth'`. Raid loot entry `url: https://foreversixty.gg/bis/mage/<spec>`, label `'Forever Sixty, best in slot and where it drops'`. Delete the `'Forever Sixty, rotation'` entry: no page shows the rotation outside the guide itself (see D1). |
| MAGE-R4-2 | CONFIRMED (low) | leveling | arcane.md:110 | "then 11 to 14 into Improved Channeling so taking hits stops pushing back your Missiles" | talents/mage.json Improved Channeling: 20% per rank for Arcane Missiles; four ranks at 14 is an 80% chance, not a stop (100% needs the fifth rank, which the build never takes). | "then 11 to 14 into Improved Channeling, so most hits no longer push back your Missiles." |
| MAGE-R4-3 | CONFIRMED (low) | leveling | fire.md:112 | "At 10 put five points into Incineration, 10% more crit on Fire Blast, then Accelerant at 15" | You have one talent point at 10; addon-builds mage-fire takes Incineration one rank a level, 10 to 14. The other two spec bullets state ranges ("11 to 14", "12 to 16"). | "From 10 to 14 put five points into Incineration, 10% more crit on Fire Blast, then Accelerant at 15 to cut 0.1 sec off Fireball, and Ignite from 16 to 19 so every Fire crit burns for more." |
| MAGE-R4-4 | LIKELY (low) | leveling/gear | arcane.md:111, fire.md:113, frost.md:114 | "Your trainer's level 10 class quest pays a staff, an orb or a wand" | The four level 10 quests loot.json tags `classes: [mage]` (Mirror Lake, Ju-Ju Heaps, Mage-tastic Gizmonitor, The Balnir Farmstead) all pay a staff or an orb. The wand comes from Boughs in the Wind (93797), which loot.json does not class-tag; its Wowhead Forever tooltip is a hand-in to Belann Windwood in Valanaar with nothing marking it as a mage or trainer quest. Every playable mage start that can be confirmed gets a staff or an orb. | "Your trainer's level 10 class quest pays a staff or an orb; take the staff and retire the starter one." |

Counts: CONFIRMED 3 (R4-1; R4-2 and R4-3 low), LIKELY 1 (R4-4, low), NEEDS-SIM 0, NEEDS-OWNER 0.

No findings for: talent legality and every point count; every trainer level in the Levels 1 to 19 bullets; Levels 1 to 19 vs Levels 20 to 60 (no contradiction; Arcane Concentration's fifth rank at 20 falls between the two headings and neither text says otherwise); wands and quest names; Fire Blast as the low-level finisher; Frost Armor armor from Frost Warding; Ice Armor's chill (every Frost and Ice Armor rank: "may have their movement slowed by 30% and the time between their attacks increased by 25%"); every spell cost, cast time and cooldown quoted in the spec text (Presence of Mind 180000 ms, Evocation 480000 ms +1500% for 8 s, Flamestrike 990 / 3 s / 8 s, Arcane Explosion 390 / 10 yd, Fire Blast 8 s, Frost Nova 25 s, Mage Armor 50%); Heating Up, Combustion, Burning Soul 47%, Arcane Meditation a third at 2/3, Winter's Chill 5 x 20%; race text vs round 3 sims; Skyborne mage only Alliance (combos.json: race 95 only); Orc as the third Horde mage race (combos.json: Undead, Troll, Orc); em/en dashes; site/sim words in body prose; quoted nightly numbers.

Links checked (all frontmatter and body):
- Blizzard, Deep Dive panel recap: 200; the recap of the Forever Deep Dive, lists Orc Mage among the six new combinations. Label right.
- Wowhead, Skyborne First Look (index.md): 200; title matches. The fetch did not return the class list, so the Alliance-only mage claim is unknown from this page, but combos.json settles it.
- Power Up Gaming, Forever talent changes vs Classic (index.md): 200; a Forever-vs-Classic talent change guide with a mage section. Label right.
- Talents Forever (demo transcription): 200; the export's `_readme` says the site "began as frame-by-frame reads of the BlizzCon 2026 demo" and now reads the beta client files. The label is out of date but not false; no fix filed.
- foreversixty.gg/sources (three spec labels): see MAGE-R4-1.
- /planner?class=mage (body, all three specs): 200.

## Decisions

- **D1. A site source must link the page that holds it.** For mage: stat weights go to /sim/weights, raid loot to the spec's /bis page, and the rotation entry goes, because the guide's own Rotation section is the only place the rotation is shown and a source pointing at the page you are reading is no source. Do not reuse /sim/specs for the rotation (warrior's pattern): that page covers which specs are modelled, not how one plays. (Warrior fury/arms/protection also point "where the raid gear drops" at /sources; flag for that class's auditor.)
- **D2. Level 10 quest reward wording follows the class-tagged quests.** Until a Forever source shows Boughs in the Wind is a mage trainer quest, the guide names only staff and orb (R4-4).

## Sims requested

- None new. Carried unchanged: SIM-R2-1 (Arcane Wrath in the Arcane builds and APL), SIM-R2-2 (Arcane 32/16/3 Fire subspec vs 32/3/16, 60 raid band), SIM-R2-3 (Fire leveling Scorch vs Fireball filler at ladder levels 25 and 30, shipped Fire order), SIM-5/SIM-6 (Frost Accelerant point; Fire 60 toll Impact + Accelerant vs Burning Soul 2), SIM-2 (Arcane Power at about +3 s), and the Arcane (Arcane Blast at 20) and Frost (Improved Frostbolt at 10) ladder-first rows.

## Sim findings

- None new.

## Sources

- Forever Sixty, How sources are classified: https://foreversixty.gg/sources (fetched; headings quoted in R4-1)
- Forever Sixty, stat weights page: https://foreversixty.gg/sim/weights; best in slot pages https://foreversixty.gg/bis/mage/arcane (and /fire, /frost)
- Blizzard, World of Warcraft: Forever Deep Dive panel recap: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- Wowhead, Skyborne First Look: https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
- Power Up Gaming, New WoW Forever talents changes vs Classic: https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/
- Talents Forever data export: https://talentsforever.com/data.json
- Wowhead Forever quest tooltips: Boughs in the Wind https://nether.wowhead.com/forever/tooltip/quest/93797, Rear Guard Patrol https://nether.wowhead.com/forever/tooltip/quest/99156
- data/builds/1.60.1.70755: talents/mage.json (Improved Channeling, Incineration), trainables/mage.json, addon-builds.json (mage-arcane, mage-fire, mage-frost leveling orders), items.json, loot.json (crafted:enchanting; quests 957, 1861, 1880, 1882, 1884, 93797, 99156), combos.json, spelltips/mage.json

## Fix log (round 4)

- MAGE-R4-1 fixed per the fix brief's round 4 source rules, which override D1 where they differ: arcane.md, fire.md, frost.md:20-28 now link "Forever Sixty, stat weights" to /sim/weights, "Forever Sixty, how <Spec> plays" to /sim/specs (relabelled from "rotation", kept rather than deleted per the brief), and "Forever Sixty, raid loot tables" to /sim/drops (not /bis, per the brief). No guide source links /sources any more.
- Talents Forever relabelled "Talents Forever, talent data from the beta client" on index.md:17, arcane.md:17, fire.md:17, frost.md:17.
- MAGE-R4-2 fixed: arcane.md:110 "so most hits no longer push back your Missiles".
- MAGE-R4-3 fixed: fire.md:112 "From 10 to 14 put five points into Incineration".
- MAGE-R4-4 and D2 fixed: arcane.md:111, fire.md:113, frost.md:114 "pays a staff or an orb".
- No wago.tools links and no level-scaled values in mage pages needing restatement; APL card notes re-read, nothing false found, unchanged. No race or build changes.
- New SIM findings: none.
