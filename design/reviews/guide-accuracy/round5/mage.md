# Guide accuracy round 5: Mage

Audited on the round 5 audit tree, head 42f10dda9 (land/guides-r4: 752c02e43 Levels 1 to 19, 77bd7c12b round 4 fixes, a901bae80 rendered-review fixes, 42f10dda9 ledger). Client 1.60.1.70755, engine pin d1d9f2775 (sim/enginever/version.go).
Files: web/src/content/guides/mage/index.md, arcane.md, fire.md, frost.md, read line by line, plus the card notes in data/curated/apl/mage-{arcane,fire,frost}.json.
Evidence: data/builds/1.60.1.70755/{talents,trainables,spelltips}/mage.json, addon-builds.json, items.json, loot.json, races.json, combos.json, bis/mage-*.json (60 band weights), stat-weights.json; design/reviews/talent-search/mage-arcane-raid.md; sim/request/testdata/ladder/mage-*.golden.md; engine sim/mage/{missile_barrage,fire_power}.go at d1d9f2775; rounds 1 to 4 mage reports, round2/3/4 race-sims.md, ledger.md; live fetches of every source link; web research under Sources.

## Builds decoded (unchanged since round 4, re-checked)

Decoded by tier/column against talents/mage.json; all legal (51 points, tier locks, prerequisites).

- **Arcane leveling** `141005213100301531/4/05350003` (31/4/16): Wand Specialization 1, Improved Channeling 4, Arcane Resilience 1, Arcane Concentration 5, Missile Barrage 2, Arcane Geometry 1, Arcane Impact 3, Arcane Blast 1, Arcane Meditation 3, Presence of Mind 1, Arcane Mind 5, Arcane Instability 3, Arcane Power 1; Incineration 4; Improved Frostbolt 5, Precision 3, Ice Shards 5, Frost Channeling 3.
- **Arcane raid** `131205213100301531/3/05350003` (32/3/16): as leveling but Improved Channeling 3, Arcane Subtlety 2, Incineration 3.
- **Fire leveling** `2520150000002/51521001230133031/003` (17/31/3): Wand Specialization 2, Improved Channeling 5, Arcane Resilience 2, Magic Absorption 1, Arcane Concentration 5, Arcane Meditation 2; Incineration 5, Accelerant 1, Ignite 5, Flame Throwing 2, Impact 1, Pyroblast 1, Improved Fire Blast 2, Improved Scorch 3, Heating Up 1, Master of Elements 3, Critical Mass 3, Fire Power 3, Combustion 1; Precision 3.
- **Fire raid** `2530050000002/50520201230133031/003`: as leveling but Arcane Resilience 3, no Magic Absorption, Accelerant or Impact, Burning Soul 2.
- **Frost (both)** `2300052/5102/253510130101301221` (12/8/31): Wand Specialization 2, Improved Channeling 3, Arcane Concentration 5, Missile Barrage 2; Incineration 5, Accelerant 1, Flame Throwing 2; Frost Warding 2, Improved Frostbolt 5, Precision 3, Ice Shards 5, Frostbite 1, Permafrost 1, Frost Channeling 3, Ice Lance 1, Ice Block 1, Shatter 3, Cold Snap 1, Fingers of Frost 2, Winter's Chill 2, Ice Barrier 1.

Shipped leveling order (addon-builds.json, index 0 = level 10; entries are [tree, tier, column], 1-based), re-decoded:
- Arcane: 10 Wand Specialization, 11 to 14 Improved Channeling, 15 Arcane Resilience, 16 to 20 Arcane Concentration, 21 to 22 Missile Barrage, 23 Arcane Geometry, 24 to 26 Arcane Impact, 27 Arcane Blast, 28 to 30 Arcane Meditation, 31 Presence of Mind, 32 to 36 Arcane Mind, 37 to 39 Arcane Instability, 40 Arcane Power, 41 to 44 Incineration, 45 to 49 Improved Frostbolt, 50 to 52 Precision, 53 to 57 Ice Shards, 58 to 60 Frost Channeling.
- Fire: 10 to 14 Incineration, 15 Accelerant, 16 to 19 Ignite, 20 Pyroblast, 21 Ignite, 22 to 23 Flame Throwing, 24 Impact, 25 to 26 Improved Fire Blast, 27 to 29 Improved Scorch, 30 Heating Up, 31 to 33 Master of Elements, 34 to 36 Critical Mass, 37 to 39 Fire Power, 40 Combustion, then Arcane and Precision.
- Frost: 10 to 11 Frost Warding, 12 to 16 Improved Frostbolt, 17 to 19 Precision, 20 to 24 Ice Shards, 25 Frostbite, 26 Permafrost, 27 to 29 Frost Channeling, 30 Ice Lance, 31 Ice Block, 32 to 34 Shatter, 35 Cold Snap, 36 to 37 Fingers of Frost, 38 to 39 Winter's Chill, 40 Ice Barrier, then Arcane and Fire.

Every talent level in the Levels 1 to 19 paragraphs and in Levels 20 to 60 matches this order; the two subsections do not contradict each other or the strip.

## What holds (checked this round)

- Levels 1 to 19 (all three specs): trainer levels (Fireball 1/6/12/18, Frostbolt 4/8/14/20, Conjure Water 4, Fire Blast 6, Arcane Missiles 8, Frost Nova 10, Arcane Explosion 14, Flamestrike 16) match trainables; talent text (Wand Specialization 13%, Improved Channeling 80% for Missiles at 4 ranks, Arcane Resilience 25% of Intellect, Incineration 10% at 5/5, Accelerant 0.1 sec, Ignite, Frost Warding 30% at 2/2, Improved Frostbolt 0.5 sec, Precision 6%); wands (Lesser Magic Wand 11287 req 5, Greater Magic Wand 11288 req 13, Elven Wand 5604 from Bashal'Aran 957 Alliance, Riptear's Spare Arm 286743 from Rear Guard Patrol 99156 Horde); the four class-tagged level 10 quests pay a staff or an orb (9513/7508, 9514/7507).
- Every cost, cast time, cooldown and level-scaled value is stated at 60 and matches: Scorch rank 7 150 mana 1.5 s, Arcane Explosion rank 6 390 mana 10 yd, Flamestrike rank 6 990 / 3 s, Blizzard rank 6 1,400 (1,190 with Frost Channeling), Arcane Blast 15% of base mana +100% a stack (five times at 4 stacks), Pyroblast 6 s (1.5 s at 3 Heating Up), Fire Blast 8 s (6 s improved), Ice Lance 8 s, Ice Block 5 min, Cold Snap 10 min, Ice Barrier 30 s, Evocation 8 min +1500% 8 s, Arcane Power / Presence of Mind / Combustion 3 min, Mage Armor 50% at every rank, Mana Ruby 1,000 to 1,200, Frostfire Bolt 370 mana 3.0 s vs Frostbolt rank 10 260 mana. Trainer levels in prose: Scorch 22, Cone of Cold 26, Mage Armor 34, Frostfire Bolt 40/50/60, Teleport: Dalaran 50, Arcane Explosion rank 2 at 22.
- statPriority frontmatter vs bis/mage-*.json 60 raid weights (haste per 1% outside the normalization): Arcane SP 1 > hit 0.96 > Int 0.67 > Arcane power 0.63 > crit 0.48; Fire hit 1.11 > SP = Fire power 1 > crit 0.83 > Int 0.55; Frost SP 1 > hit 0.97 > Frost power 0.82 > crit 0.42 > Int 0.17. All three orders and every stat paragraph agree (Arcane Intellect a bit above Fire, far above Frost).
- "Its 15 points of resistance reduction is worth less than the crit" (arcane.md:64): talent search scores Arcane Subtlety "modeled, no damage" and an Incineration rank +1.93 (talent-search/mage-arcane-raid.md:48, :63).
- Fire Power counts Improved Scorch, Ignite and the Pyroblast burn, mixed schools excluded (engine fire_power.go, client text).
- Races vs round 3 sims (raid and bare) and races.json: Gnome / Undead (Arcane, Frost), Human / Undead (Fire), Orc ahead of Troll for Arcane and Fire, Troll and Orc tied behind Undead for Frost, Skyborne paid and Alliance-only (combos.json), Will of the Forsaken removes fears (races.json: no longer an immunity; the guide says "breaks fears").
- Round 4 fixes R4-1 to R4-4 hold; frontmatter sources follow the round 4 rule (stat weights to /sim/weights, "which boss drops are upgrades" to /sim/drops, no /sources link, no "how plays" entry, Talents Forever labelled beta client data).
- Card notes (mage-arcane, mage-fire, mage-frost priorityList notes): nothing false found.
- No em or en dashes, no site/sim words in body prose, no quoted nightly numbers.

Not re-filed (fixed, sim round or accepted): MAGE-1 to 16, MAGE-R2-1 to R2-8, MAGE-R3-1 to R3-10, MAGE-R4-1 to R4-4, SIM-1 to SIM-6, SIM-R2-1 to R2-3, the Arcane (Arcane Blast at 20) and Frost (Improved Frostbolt at 10) ladder-first orders, Accelerant/Impact toll points, component-rendered parts.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| MAGE-R5-1 | CONFIRMED (low) | stale/voice | index.md:29 | Under "What changed from 1.12": "**Ice Block is a talent.** It sits in the Frost tree at 15 points." | Ice Block was already a Frost talent in 1.12: tier 5, 20 points (Warcraft Wiki, Ice Block: Tier 5 Frost talent in Classic; it became trainable only in BC). Forever moved it to tier 3 (talents/mage.json: Ice Block tier index 3, 15 points) and pushed Cold Snap down to tier 4 (20 points). Being a talent is not the change; the tier is. | "- **Ice Block comes earlier.** It is still a Frost talent, but it sits at 15 points instead of 20, with Cold Snap below it now. Fire mages never get there, and the Arcane builds have no spare point for it, so only Frost mages carry Ice Block." |
| MAGE-R5-2 | CONFIRMED (low) | stale/voice | index.md:30 | Under "What changed from 1.12": "**No Molten Armor and no Invisibility.**" | Neither spell existed in 1.12: Molten Armor and Invisibility arrived with The Burning Crusade (trained at 62 and 68). Listing their absence as a change from 1.12 tells a 1.12 mage something was taken away. trainables/mage.json has neither, so the facts in the rest of the bullet hold. | "- **Armor and threat.** Mage Armor is the raid armor and now keeps half your mana regeneration running while you cast; Ice Armor is for when things hit you. There is no Molten Armor or Invisibility, so threat is handled with Arcane Subtlety, Burning Soul and Frost Channeling, and by not out-casting your tank." |
| MAGE-R5-3 | CONFIRMED (low) | stale/voice | index.md:31 | "**One hit stat.** Hit is a single stat for every school, and Precision in the Frost tree adds 6% hit to every spell" | Spell hit in 1.12 already covered every school; what Forever merged is spell, melee and ranged hit into one stat (method.gg, How does hit chance work in WoW Forever; Warcraft Tavern, Forever stats guide). The real school change is the talent: 1.12's Elemental Precision gave 6% hit to Fire and Frost only, Forever's Precision gives it to all spells (talents/mage.json Precision: "all spells"), which is what makes Arcane take it too. | "- **One hit stat.** Hit now counts for spells and attacks alike, and Precision in the Frost tree adds 6% hit to every school where 1.12's Elemental Precision covered only Fire and Frost, which is why every build, Arcane included, dips three points there." |
| MAGE-R5-4 | LIKELY (low) | talent | arcane.md:62 | "and Arcane Geometry gives you the extra range for pulling." | The one rank is 3 yards on Arcane spells only (talents/mage.json rank 1; Missiles, Blast, Explosion). The guide pulls with Frostbolt at every level (arcane.md:106 "open with Frostbolt", :114 "Frostbolt's slow is your kiting tool"), which Geometry does not touch. It is also not a tier key: without it the build still has 17 Arcane points before Arcane Meditation's 15 (addon-builds order). | "and one point of Arcane Geometry adds 3 yards to Arcane Blast and Missiles, so you can blast from a little further out." |
| MAGE-R5-5 | LIKELY (low) | stats | frost.md:89 | "a missed Frostbolt is a missed Fingers of Frost and Missile Barrage chance." | Fingers of Frost rolls off a Chill effect, which only a landed Frostbolt applies. Missile Barrage does not ask for a hit: the client text is "Gives your ... Frostbolt ... spells a 10/20% chance" (no "when it hits", unlike Arcane Concentration's "after any damage spell hits"), and the engine rolls it on cast complete (sim/mage/missile_barrage.go, "Missile Barrage Trigger" OnCastComplete). | "a missed Frostbolt is a missed Fingers of Frost chance and no chill on the boss." |

Counts: CONFIRMED 3 (R5-1 to R5-3, all low), LIKELY 2 (R5-4, R5-5, low), NEEDS-SIM 0, NEEDS-OWNER 0.

No findings for: talent legality and every point count; every leveling level vs the shipped order; every spell level, cost, cast time and cooldown; Arcane loop, opener, Arcane Power windows, mana plan and gem/rune/Robe text; Fire Heating Up, Combustion, Improved Scorch, Fire Power, Burning Soul 47%/20%, toll points; Frost Ice Barrier, Fingers of Frost, Ice Lance, Shatter, Winter's Chill 40% at 2/5 and 5 x 20% + 5 x 7%, Ice Block threat, Cone of Cold 10 yd (12 with Arctic Reach); stat paragraphs; race paragraphs; professions; card notes.

## Links checked

- Blizzard, Deep Dive panel recap (all four pages): 200; "World of Warcraft: Forever Deep Dive Panel Recap", lists Orc Mage among the six new pairs. Label right.
- Wowhead, Skyborne First Look (index.md): 200; title matches. The fetch did not return the class list; combos.json settles the Alliance-only mage line.
- Power Up Gaming, Forever talent changes vs Classic (index.md): 200; "New WoW Forever Talents: Changes vs Classic Guide", has a mage section (Arcane Blast new, 6 new / 32 changed talents). It calls the Pyroblast talent "Hot Streak"; the guide rightly uses the client's name, Heating Up. Label right.
- Talents Forever, talent data from the beta client (all four pages): 200, data.json; `_readme` says the data comes from the beta client's files. Label right.
- Forever Sixty, stat weights (three specs): https://foreversixty.gg/sim/weights, 200, "Stat weights", what one point of each stat is worth for a damage spec. Label right, damage specs only.
- Forever Sixty, which boss drops are upgrades (three specs): https://foreversixty.gg/sim/drops, 200, "Droptimizer", simulates each boss drop against your gear. Label right.
- /planner?class=mage (body, three specs): planner link, unchanged.

## Decisions

- **D1. "What changed from 1.12" lists only real changes from 1.12.** A spell that was never in 1.12 (TBC's Molten Armor, Invisibility, Icy Veins) can be named as missing, but not as removed; a talent that moved is described by where it moved. Applied in R5-1 to R5-3.

## Sims requested

- None new. Carried unchanged: SIM-R2-1 (Arcane Wrath in the Arcane builds and APL), SIM-R2-2 (Arcane 32/16/3 Fire subspec vs 32/3/16, 60 raid band), SIM-R2-3 (Fire leveling Scorch vs Fireball filler at ladder levels 25 and 30), SIM-5/SIM-6 (Frost Accelerant point; Fire 60 toll Impact + Accelerant vs Burning Soul 2), SIM-2 (Arcane Power at about +3 s), the Arcane (Arcane Blast at 20) and Frost (Improved Frostbolt at 10) ladder-first rows.

## Sim findings

- None new. The engine's on-cast Missile Barrage roll (R5-5) matches the client text, which does not require a hit.

## Sources

- Warcraft Wiki, Ice Block (Classic: Tier 5 Frost talent; trainable from BC): https://warcraft.wiki.gg/wiki/Ice_Block
- method.gg, How Does Hit Chance Work in WoW Forever: https://www.method.gg/wow-forever/how-does-hit-chance-work-in-wow-forever
- Warcraft Tavern, Forever Stats and Attributes Guide: https://www.warcrafttavern.com/forever/guides/stats/
- Blizzard, World of Warcraft: Forever Deep Dive panel recap: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- Wowhead, Skyborne First Look: https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
- Power Up Gaming, New WoW Forever talents changes vs Classic: https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/
- Talents Forever data export: https://talentsforever.com/data.json
- Forever Sixty, stat weights https://foreversixty.gg/sim/weights and Droptimizer https://foreversixty.gg/sim/drops
- data/builds/1.60.1.70755: talents/mage.json (Ice Block tier, Precision, Arcane Geometry, Missile Barrage), trainables/mage.json, spelltips/mage.json, addon-builds.json, bis/mage-*.json, races.json, combos.json, items.json, loot.json
- Engine d1d9f2775: sim/mage/missile_barrage.go (on-cast trigger), sim/mage/fire_power.go

## Fix log (round 5)

- MAGE-R5-1: fixed, web/src/content/guides/mage/index.md:29. "Ice Block comes earlier": still a Frost talent, now at 15 points instead of 20, Cold Snap below it (talents/mage.json: Ice Block tier index 3, Cold Snap tier index 4).
- MAGE-R5-2: fixed, index.md:30. Bullet retitled "Armor and threat"; Molten Armor and Invisibility named as absent, not removed.
- MAGE-R5-3: fixed, index.md:31. Hit merged across spells and attacks; Precision covers every school where Elemental Precision covered Fire and Frost (talents/mage.json Precision "all spells").
- MAGE-R5-4: fixed (verified: talents/mage.json Arcane Geometry rank 1 "range of your Arcane spells by 3 yards"), web/src/content/guides/mage/arcane.md:62. No pulling claim; 3 yards on Arcane Blast and Missiles.
- MAGE-R5-5: fixed (verified: client text for Missile Barrage has no hit condition; engine rolls on cast complete), web/src/content/guides/mage/frost.md:89. Missile Barrage dropped from the missed-Frostbolt line; the miss costs the Fingers of Frost chance and the chill.
- D1: applied through R5-1 to R5-3; the rest of the "What changed from 1.12" list re-read (Personal debuffs, New spells): only real changes remain.
- Card notes: grep of mage-*.json notes for Molten, Invisibility, Geometry and Missile Barrage hit wording: nothing to change. Build strings, point orders, card actions untouched. No em or en dashes.
- New SIM findings: none.
