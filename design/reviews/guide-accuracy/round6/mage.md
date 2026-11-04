# Guide accuracy round 6: Mage

Audited on the round 6 audit tree, land/guides-r5 head 367cbadb7 (all five rounds applied). Client 1.60.1.70755, engine pin d1d9f2775 (sim/enginever/version.go).
Files: web/src/content/guides/mage/index.md, arcane.md, fire.md, frost.md, read line by line, plus the rendered card notes (priorityList notes) in data/curated/apl/mage-{arcane,fire,frost}.json.
Evidence: data/builds/1.60.1.70755/{talents,trainables,spelltips}/mage.json, addon-builds.json, items.json, races.json, bis/mage-{arcane,fire,frost}.json (generated 2026-11-04T03:1x, nightly de044125c; band 60 raid rows, the rows GuideStatTable renders through band60Weights and railStatRows), web/src/lib/guides/rail-stats.ts and class-dps.ts (how the page's stat table orders rows: per-point scale factor, descending); rounds 1 to 5 mage reports, round4/mage-levels.md, race-sims.md for rounds 2 to 5, ledger.md; live fetches of every source link.

Diff since round 5's audit tree (42f10dda9): only the five round 5 fixes in mage prose (index.md:29-31, arcane.md:62, frost.md:89); no change to the builds, addon-builds.json, talents, trainables or mage APL files. The BiS files (and so the stat tables) are the nightly re-rank de044125c.

## Builds decoded (unchanged, re-checked)

Decoded by tier/column against talents/mage.json; all legal (51 points, tier locks, prerequisites).

- **Arcane leveling** `141005213100301531/4/05350003` (31/4/16): Wand Specialization 1, Improved Channeling 4, Arcane Resilience 1, Arcane Concentration 5, Missile Barrage 2, Arcane Geometry 1, Arcane Impact 3, Arcane Blast 1, Arcane Meditation 3, Presence of Mind 1, Arcane Mind 5, Arcane Instability 3, Arcane Power 1; Incineration 4; Improved Frostbolt 5, Precision 3, Ice Shards 5, Frost Channeling 3.
- **Arcane raid** `131205213100301531/3/05350003` (32/3/16): Improved Channeling 3, Arcane Subtlety 2, Incineration 3, rest as leveling.
- **Fire leveling** `2520150000002/51521001230133031/003` (17/31/3) and **raid** `2530050000002/50520201230133031/003`: as round 5 (raid: Arcane Resilience 3, no Magic Absorption, Accelerant or Impact, Burning Soul 2).
- **Frost (both)** `2300052/5102/253510130101301221` (12/8/31).

Shipped leveling order (addon-builds.json, index 0 = level 10) unchanged from round 5's decode; every talent level in "Levels 1 to 19" and "Levels 20 to 60" of all three specs matches it, and the two blocks read in order without contradiction (Arcane: Arcane Concentration 16 to 19, then "every point into Arcane up to 40"; Fire: Ignite 16 to 19, then Scorch at 22 and Heating Up at 30; Frost: Precision 17 to 19, then Ice Shards in the early 20s).

## Stat tables now (band 60 raid, per-point scale factors as the page's table sorts them)

From bis/mage-*.json, alliance row (the first recommended race's faction), insignificant rows (Spirit, mp5) last, haste in its per 1% caption:
- **Arcane:** Hit rating 1.58, Spell power 1.52, Arcane power 1.04, Intellect 1.00, Crit rating 0.77. Normalized to the top row: Hit 1.00, SP 0.96, Arcane power 0.66, Int 0.63, crit 0.49.
- **Fire:** Hit rating 4.44, Spell power 2.69 = Fire power 2.69, Crit rating 2.33, Intellect 1.00. Normalized: Hit 1.00, SP 0.61, Fire power 0.61, crit 0.52, Int 0.23.
- **Frost:** Hit rating 7.50, Spell power 5.12, Frost power 4.37, Crit rating 2.88, Intellect 1.00. Normalized: Hit 1.00, SP 0.68, Frost power 0.58, crit 0.38, Int 0.13.

Every spec: Hit and spell power top the table, which all three stat paragraphs say. The frontmatter lists are only the set of rows (GuideStatTable re-sorts them), and StatPriorityPills, the only component that would show the raw frontmatter order, is not used by any page.

## What holds (checked this round)

- Round 5 fixes: R5-1 (Ice Block was tier 5, 20 points, in 1.12 with Cold Snap at tier 3 above it; Forever tier 3 / tier 4, so "15 points instead of 20, with Cold Snap below it now" is right), R5-2 (1.12 Mage Armor kept 30% of regen, Forever 50% at every rank: spelltips 6117, 22782, 22783; "now keeps half" is right), R5-3 (method.gg: "Hit Chance was changed to affect everything, so Hit Chance will affect both spells and attacks"; 1.12 Elemental Precision was Fire and Frost only), R5-4 (Arcane Geometry rank 1 is 3 yards on Arcane spells), R5-5 (Missile Barrage has no hit condition) all hold.
- Every talent line in the three spec pages against talents/mage.json (Arcane Blast 12% Int / +100% cost / 4 stacks / 18 s; Missile Barrage 40% / 20%; Arcane Concentration 10%; Arcane Meditation 50% at 3/3, a third at 2/3; Arcane Mind 10% Int and +100% Arcane crit bonus; Arcane Instability 3%/3%; Improved Scorch, Heating Up, Ignite 40% / 4 s, Fire Power 3% per effect to 12%, Critical Mass 6%, Master of Elements 30%, Improved Fire Blast 2 s / 15 yd, Combustion, Burning Soul 47% / 20%; Ice Barrier 15% + 5% and no pushback, Fingers of Frost 15% / 2 spells / 15 s, Ice Lance +300%, Shatter 50%, Ice Shards, Frost Channeling 15% / 30%, Winter's Chill 40% at 2/5, Precision 6% all spells).
- Point arithmetic: Fire 28 damage-and-range points plus a 2 point toll to the 30 Combustion asks; Frost 30 points before Ice Barrier with Winter's Chill 2 as the toll; Arcane 31 and 32 in Arcane; Arcane raid trade (-1 Improved Channeling, -1 Incineration, +2 Arcane Subtlety).
- Every cost, cast time, cooldown and level-scaled value at 60 (Scorch 150 / 1.5 s, Arcane Explosion rank 6 390 / 10 yd, Flamestrike 990 / 3 s, Blizzard rank 6 1,400 and about 1,200 with Frost Channeling, Fire Blast 8 s, Frost Nova 25 s, Presence of Mind 180 s, Evocation 480 s / +1500% / 8 s, Ice Block 300 s / 10 s, Ice Lance 8 s, Frostfire Bolt 370 / 3.0 s and Frostbolt 3.0 s base, Arcane Blast 1.5 s), trainer levels (Scorch 22, Cone of Cold 26, Mage Armor 34, Frostfire Bolt 40/50/60, Teleport: Dalaran 50, Arcane Brilliance 56, every Levels 1 to 19 level).
- Arcane Brilliance "for every caster in the raid": spelltip 23028 "Infuses the target's party and raid with brilliance".
- Ice Armor's chill (every rank), Frost Warding 30% at 2/2, Wand Specialization 13% at 1/2, Arcane Resilience 25% of Intellect at 1/3.
- Mana plan: gem and rune share one 2 minute cooldown, Robe its own 5, a 1 minute category lockout (round 1 O4, client ItemEffect/SpellCategories). The Robe of the Archmage is still in all three 60 sets (bis/mage-*.json), so naming its cooldown is not stale gear prose.
- Race paragraphs vs round 3 race sims and races.json; Skyborne named only as the paid option; Human for Fire; Undead first on the Horde; Orc the third Horde mage race; Blood Fury 10% spell power 15 s / 2 min; Berserking 10% / 15 s / 3 min; Eureka! 10% / 10% / 8 s / 2 min; Will of the Forsaken removes fears.
- Card notes (priorityList notes, all three files): nothing false. Fire's "Fireball's 395" is rank 11 (10151), the Phase 1 top rank; Arcane Missiles rank 7 as top trainer rank, rank 8 an Ahn'Qiraj book.
- No em or en dashes, no site/sim words in body prose, no quoted nightly numbers. Gear named in prose: only wands and the level 10 quest reward (quest and crafted sources, stable), the Robe for its use cooldown.

Not re-filed (fixed, sim round, accepted or with the data lane): MAGE-1 to 16, MAGE-R2-1 to R2-8, MAGE-R3-1 to R3-10, MAGE-R4-1 to R4-4, MAGE-R5-1 to R5-5, SIM-1 to SIM-6, SIM-R2-1 to R2-3, the Arcane (Arcane Blast at 20) and Frost (Improved Frostbolt at 10) ladder-first orders, the Accelerant/Impact toll points, component-rendered parts, and the round 6 data-lane list (Frostwolf/Stormpike rows, "built in" hit baseline, BiS race captions).

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| MAGE-R6-1 | CONFIRMED (low) | stats | arcane.md:86 | "Intellect is close to a damage stat for this spec: a bit better than for Fire, which wants it for mana, and far better than for Frost." | Stat tables now (band 60 raid, normalized to each page's top row): Intellect is 0.63 on Arcane, 0.23 on Fire, 0.13 on Frost. Arcane's Intellect is close to three times Fire's, a bigger gap than Fire to Frost, so "a bit better than for Fire ... far better than for Frost" reads the wrong way round against both tables. The comparison is also a nightly-moving one; the stable reason (Arcane Blast turns Intellect into spell damage, Fire and Frost only drink it as mana) is enough. | "Arcane Blast gives spell damage from Intellect, so Intellect is close to a damage stat for this spec and worth far more to you than to a Fire or Frost mage, who only want it for mana." (rest of the paragraph unchanged) |
| MAGE-R6-2 | LIKELY (low) | stats | fire.md:88 | "Intellect is your mana bar, and Fire needs more of it than Frost. Crit feeds Ignite, Heating Up and Master of Elements at once." | The paragraph walks the stats in the order hit and spell power, Intellect, crit; the page's table now shows Hit, Spell power, Fire power, Crit rating, Intellect (crit 2.33 against Intellect 1.00 per point). The order a reader takes from the prose is not the table's. | "Hit and spell power lead for Fire until the hit cap: every missed Scorch is a missed stack, a missed Ignite and a missed Heating Up chance. Crit comes next, because it feeds Ignite, Heating Up and Master of Elements at once. Intellect is your mana bar, and Fire needs more of it than Frost." |
| MAGE-R6-3 | LIKELY (low) | consumable | arcane.md:92 | "Arcane casts Arcane and Frost spells, so any spell power enchant or elixir that names one school loses half its value." | The Arcane table puts Arcane power at 0.68 of spell power (bis/mage-arcane.json band 60 raid: 1.04 against 1.52): Arcane-only power keeps about two thirds of its value and Frost-only power about a third. Neither loses half. A share that moves with gear should not be stated as a fraction anyway. | "Arcane casts Arcane and Frost spells, so spell power that names one school is only worth part of its value: Arcane power covers your Blasts and Missiles but not your Frostbolts, and Frost power covers only the Frostbolts. Bring Major Mana Potions and Demonic Runes to every boss; you will use both." |
| MAGE-R6-4 | LIKELY (low) | race | arcane.md:96 | "Read Leyline (50% more health and mana regeneration for 15 seconds, 2 minute cooldown) and Wind Blessed's 2% haste still do less damage than a Gnome or a Human." | races.json (High Order Skyborne): "Read Leyline is spell 1259705 ... The sim does not cast it on its own". The race sims behind this line (round3/race-sims.md, Skyborne 607.5 against Gnome 614.0 and Human 611.8) measured Skyborne without Read Leyline, so the sentence claims a measurement that does not include half of what it names. For the most mana-hungry spec, 15 seconds of extra regeneration every 2 minutes is not obviously nothing. | Until the sim below runs: "**Skyborne** is a paid unlock and only Alliance for mages; Wind Blessed's 2% haste alone does less damage than a Gnome or a Human, and Read Leyline (50% more health and mana regeneration for 15 seconds, 2 minute cooldown) is a mana button, not a damage one." If the sim keeps Skyborne behind both with Read Leyline cast, the current sentence may stay. |
| MAGE-R6-5 | LIKELY (low) | voice | fire.md:64 | "The other 20 points buy mana. Seventeen in Arcane reach Arcane Concentration ... Three go into Precision in Frost for 6% hit." | The paragraph contradicts its own first sentence: three of the twenty are Precision (hit), and of the seventeen Arcane points Wand Specialization, Improved Channeling and Arcane Resilience (9 or 10 of them) are row fillers on the way to Arcane Concentration, not mana. | "The other 20 points buy mana and hit. Seventeen in Arcane reach **Arcane Concentration** (5 points, a free spell 10% of the time) and two ranks of **Arcane Meditation** (a third of your mana regeneration keeps running while you cast). With Mage Armor on top, most of your regeneration runs through the whole fight. Three go into **Precision** in Frost for 6% hit." |

Counts: CONFIRMED 1 (R6-1, low), LIKELY 4 (R6-2 to R6-5, all low), NEEDS-SIM 0 as findings (R6-4's measurement is under Sims requested), NEEDS-OWNER 0.

No findings for: build legality and point counts; every leveling level against the shipped order and the order the two leveling blocks read in; every spell level, rank, cost, cast time and cooldown; every talent description quoted; the Arcane loop, opener, Arcane Power windows, mana plan, threat and movement text; Fire's opener, Heating Up, Combustion, Ignite and threat, toll points, Frostfire Bolt line; Frost's Ice Barrier, Fingers of Frost, Ice Lance, Shatter, Winter's Chill, Cone of Cold reach, Ice Block threat, Blizzard cost; the Frost stat paragraph and the hit/spell power lead in all three; the index page; race paragraphs other than R6-4; professions; card notes.

## Links checked

- Blizzard, Deep Dive panel recap (all four pages): 200, "World of Warcraft: Forever Deep Dive Panel Recap", lists "Gnome Priest, Human Hunter, Dwarf Shaman, Orc Mage, Troll Warlock, and Undead Paladin". Label right.
- Wowhead, Skyborne First Look (index.md): 200, title "Skyborne First Look: New Neutral Race in World of Warcraft: Forever". The fetch returns no article body, so the class list is unknown from the page; combos.json and races.json settle the Alliance-only mage line. Label right.
- Power Up Gaming, Forever talent changes vs Classic (index.md): 200, "New WoW Forever Talents: Changes vs Classic Guide", with a mage section (Arcane Blast new, the Pyroblast stacking talent, Improved Counterspell changed). Label right.
- Talents Forever, talent data from the beta client (all four pages): 200; `_readme` says the talent data comes from the WoW Forever beta client's own data files. Label right.
- Forever Sixty, stat weights (three specs): https://foreversixty.gg/sim/weights, 200, "Stat weights", damage specs only. Label right.
- Forever Sixty, which boss drops are upgrades (three specs): https://foreversixty.gg/sim/drops, 200, "Droptimizer", simulates each boss drop against your gear. Label right.
- /planner?class=mage (body, three specs): planner link, unchanged.
- All entries follow the round 4 source-link rule (no /sources, weights to /sim/weights on damage specs, drops to /sim/drops, no "how plays" entry, Talents Forever labelled beta client data).

## Decisions

- **D1. Stat prose compares stats only by stable reasons.** A cross-spec comparison of how much a stat is worth ("a bit better than Fire") or a fraction of value ("loses half") moves with the nightly tables, so the prose names the mechanism (Arcane Blast converts Intellect; school power covers only the spells of that school) and lets the table carry the size. Applied in R6-1 and R6-3.
- **D2. A stat paragraph that walks several stats walks them in the table's order.** Applied in R6-2.
- **D3. A race claim names only what was measured.** Read Leyline is not cast by the race sims, so it cannot be part of a "does less damage" line until measured (R6-4).

## Sims requested

- **S1 (Skyborne Arcane with Read Leyline).** mage-arcane, raid preset, band 60, guide raidBuild, published gear, 10,000 iterations, seeds 1 and 50001: `-race high-order-skyborne` with a custom rotation line casting Read Leyline (1259705) on cooldown, ideally paired with Evocation or the low-mana phase, against `-race gnome` and `-race human` on the same gear. If Skyborne stays behind both beyond error, the current arcane.md:96 sentence may stand; otherwise apply R6-4's text (Skyborne is still only named as the paid option either way).
- **S2 (race picks on the re-ranked gear).** The mage race decisions rest on round 3's gear; the nightly re-rank de044125c changed the 60 sets (round 5 re-measured only Retribution). Re-run round 3's mage rows (raid and bare, all six mage races, guide raid builds) on the current bis/mage-*.json gear. Change recommendedRaces or the race paragraphs only if an order flips beyond error (Gnome over Human for Arcane and Frost, Human over Gnome for Fire, Undead first on the Horde).
- Carried unchanged: SIM-R2-1 (Arcane Wrath), SIM-R2-2 (Arcane 32/16/3 Fire subspec), SIM-R2-3 (Fire leveling Scorch vs Fireball at 25 and 30), SIM-5/SIM-6 (Frost Accelerant point; Fire 60 toll), SIM-2 (Arcane Power at about +3 s), the Arcane and Frost ladder-first rows.

## Sim findings

- None new. The race sims leave Read Leyline uncast (races.json says so), which is an unmodeled racial rather than an engine error; S1 measures it with a custom line.

## Sources

- method.gg, How Does Hit Chance Work in WoW Forever: https://www.method.gg/wow-forever/how-does-hit-chance-work-in-wow-forever (quote under "What holds")
- Blizzard, World of Warcraft: Forever Deep Dive panel recap: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- Wowhead, Skyborne First Look: https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
- Power Up Gaming, New WoW Forever talents changes vs Classic: https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/
- Talents Forever data export: https://talentsforever.com/data.json
- Forever Sixty, stat weights https://foreversixty.gg/sim/weights and Droptimizer https://foreversixty.gg/sim/drops
- data/builds/1.60.1.70755: bis/mage-arcane.json, mage-fire.json, mage-frost.json (band 60 raid weights), talents/mage.json, trainables/mage.json, spelltips/mage.json (6117, 22782, 22783, 23028, 11958, 12043, 12051, 10187, 10202), races.json (High Order Skyborne: Read Leyline not cast by the sim), addon-builds.json, items.json
- web/src/lib/guides/rail-stats.ts (railStatRows sorts by per-point scale factor), web/src/lib/guides/class-dps.ts (band60Weights), web/src/components/guides/GuideStatTable.astro
- design/reviews/guide-accuracy/round3/race-sims.md (mage rows)

## Fix log (round 6)

New stat rule applied on all three spec pages: each `## Stat priority` paragraph now opens "The table above gives the order for your gear; this is why each stat counts." and gives only lasting reasons per stat (talent or mechanic, hit-cap rule "once your gear reaches the cap more hit adds nothing"), no ranking words and no fractions. Every talent named was checked against talents/mage.json and the decoded builds (Precision 3 in all builds, Arcane Mind 5 in Arcane, Critical Mass 3 and Combustion in both Fire builds, Ice Shards 5 and Shatter 3 in Frost). `updated:` bumped to 2026-11-04 on arcane.md, fire.md and frost.md (index.md unchanged: it has no stat ranking). Prose test `-t mage` passes.

- **MAGE-R6-1** fixed by the stat rule, arcane.md:86. The cross-spec "a bit better than Fire, far better than Frost" is gone; Intellect's reason is the mechanism (each Arcane Blast stack adds 12% of Intellect as spell damage, Arcane Mind adds 10% Intellect; Fire and Frost only drink it as mana). Crit's reason added (Arcane Mind doubles the Arcane crit bonus).
- **MAGE-R6-2** fixed by the stat rule, fire.md:88. No order in the prose now; reasons for hit (Scorch stacks, Ignite, Heating Up), spell power = Fire power, crit (Ignite, Heating Up, Master of Elements, Critical Mass, Combustion), Intellect (mana bar, a Major Mana Potion every boss).
- **MAGE-R6-3** fixed, arcane.md:92. No fraction: "spell power that names one school covers only part of your damage: Arcane power covers your Blasts and Missiles but not your Frostbolts, and Frost power covers only the Frostbolts."
- **MAGE-R6-4** fixed per integrator note, arcane.md:96: "Played without Read Leyline (...), Wind Blessed's 2% haste alone does less damage than a Gnome or a Human." Same D3 wording carried to fire.md:96 ("played without Read Leyline it does the least Fire damage of the lot"). Frost's Skyborne line already names only Wind Blessed's haste, left.
- **MAGE-R6-5** fixed, fire.md:64: "The other 20 points buy mana and hit."
- Frost stat paragraph (frost.md:89) rewritten under the rule: the "two best stats" and "needs less Intellect than Arcane or Fire" comparisons removed; crit's reason added (Ice Shards doubles the Frost crit bonus, Shatter adds 50% crit on a Fingers of Frost charge).
- **D1, D2, D3** applied as above (D2 is superseded by the new rule: the prose no longer walks an order at all).
- **S1, S2** left: race re-runs wait for the final best in slot re-rank (integrator). recommendedRaces and race order unchanged.
- Other pages checked for nightly stat or gear ranking outside Stat priority: none left (arcane.md:64's Arcane Subtlety trade is a talent choice, not a nightly stat value).
- New SIM findings: none.
