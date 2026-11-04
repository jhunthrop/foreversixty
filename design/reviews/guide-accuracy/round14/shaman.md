# Guide accuracy round 14: Shaman

Auditor: shaman (class index, Elemental, Enhancement, Restoration). Audit tree main d7fc51fe7, client 1.60.1.70755, engine pin f637ec32c (`sim/enginever/version.go`; source read from the gr14 engine checkout at that sha).
Files: `web/src/content/guides/shaman/{index,elemental,enhancement,restoration}.md`, card notes in `data/curated/apl/shaman-{elemental,enhancement,restoration}.json`, rank rows in `data/curated/rank-choices.json`, live BiS `data/builds/1.60.1.70755/bis/shaman-*.json` (de044125c picks). Read first: rounds 1 to 13 `shaman.md`, `round4/shaman-levels.md`, `round2..5/race-sims.md`, `ledger.md` through round 13, the round 1, 2 and 14 briefs.

## What moved since round 13 (604b78a80 to d7fc51fe7)

- `elemental.md:80`: "drop to rank 8 the moment it is not. Rank 8 is the cheapest bolt that still takes all of your spell damage at 60."
- `elemental.md:116`: the "three or four ranks below" line became a measured ladder (rank 3 from 30 to 33, 4 until 38, 5 until 44, 6 until 50, 7 until 56, 8 from 57).
- `restoration.md:120`: a new Healing Wave ladder for dungeons (rank 2 until 25, 3 until 36, 4 at 37, 5 until 43, 6 until 56, 7 at 57 and 58, 8 at 59; top rank for Nature's Swiftness).
- Cards: Elemental .8 now lists the leveling ranks; the Elemental prepull and lines 7/8 cast rank 10 (15208); Restoration .17 lists the leveling ranks; Restoration .5 (Nature's Swiftness Healing Wave) gained `auraIsKnown 16188`.
- `rank-choices.json`: Elemental Lightning Bolt rows 30 to 60 (lines [8], prepull [1]); Restoration Healing Wave rows 20 to 59 (lines [17]).
- Engine: the level-gap rule now keys on the client `SpellLevels.MaxLevel` (`sim/core/spell.go` `SpellLevelPenalty`, grace 12 past MaxLevel, 5% a level). No shaman Lightning Bolt, Healing Wave (except rank 1), Chain Heal or Lesser Healing Wave rank is in `clientSpellMaxLevel`, so every shaman rank still keys on learn level + 17, as the pages say.
- BiS: no change to the files, but the round 14 rule makes them binding (no provisional carve-out).
- `addon-builds.json`, `talents/`, `trainables/`, `races.json`: nothing shaman.

## Builds and shipped leveling order

Unchanged from round 13 and re-decoded against `talents/shaman.json` and `addon-builds.json`: all legal, 51 points, prerequisites met. Every "Levels 1 to 19" and "Levels 20 to 60" paragraph matches the order level for level, and nothing contradicts the header card.

## What checked out (no finding)

- **Elemental bolt ladder** (elemental.md:116, card .8 bands): matches `rank-choices.json` exactly (30-33 r3, 34-38 r4, 39-44 r5, 45-50 r6, 51-56 r7, 57-59 r8, 60 bare r8). "Still keeps most of your spell damage" holds on the engine rule with client MaxLevel (learn + 5): rank 3 (MaxLevel 19) keeps 90% at 33, ranks 4 to 7 keep 95% at the last level of their band, rank 8 (MaxLevel 49) keeps 100% to 60. "Costs far less" holds at every band (45 vs 85, 60 vs 110/135, 85 vs 135/160, 110 vs 160/190, 135 vs 190/220, 160 vs 220).
- **"Rank 8 is the cheapest bolt that still takes all of your spell damage at 60"** (elemental.md:80): rank 8 gap 60 - 49 - 12 = -1, full; rank 7 (MaxLevel 43) keeps 75%. Correct.
- **Restoration Healing Wave ladder** (restoration.md:120, card .17): matches `rank-choices.json` exactly; every listed rank is cheaper than the top rank at its levels; Nature's Swiftness is learned at 44 so "keep your top rank for Nature's Swiftness" never names a spell you do not have. Restoration card .5 casts the top learned rank now that the rank row is scoped to `lines: [17]` (default rule: highest rank learned not above rank 9), so round 13's NS request is settled (Decisions D1).
- **Low ranks** (restoration.md:76) re-derived on the new MaxLevel engine rule: Healing Wave 7 and Chain Heal 1 85%, Lesser Healing Wave 3 65%, Healing Wave 6 45%, rank 5 (24) 5%, ranks learned at 23 or earlier 0. Unchanged and exact.
- **Bind facts against the live BiS rows, both factions.** Elemental raid (dwarf and orc): Frozen Heart of the Mountain and Totem of Thunder are the only bind-on-pickup crafts (`items/shaman.json` bonding "pickup"); every other craft (Blue Dragonscale Helm, Ironfeather Shoulders/Breastplate/Leggings, Gloves of the Greatfather, Mantle of the Timbermaw) is "equip". Restoration raid (dwarf and tauren): Dormant Heart is the only BoP craft; Totem of Urgency (Enchanting) binds when equipped, so "every other craft ... can be bought" holds; Glowing Brightwood Staff is the main hand for both factions ("at 60 the staff usually wins" holds). Enhancement Alliance raid: Frozen and Molten Heart are the two BoP crafts; every other craft is "equip"; main hand is a two-hander for both factions.
- **Consumables** against `sim/request/raid-consumes.json`: Elemental Brilliant Wizard Oil, Flask of Supreme Power, Greater Arcane Elixir, Major Mana Potion, Demonic Rune; Enhancement Mongoose, Juju Power, Juju Might, R.O.I.D.S., Smoked Desert Dumpling, Rockbiter under the totem; Restoration Brilliant Mana Oil, Major Mana Potion, Demonic Rune. All present.
- **Races** against `races.json` and the measured rows (round3 race-sims: Elemental orc 630.5 > troll 624.8 ≈ tauren 624.6, Skyborne 633.4 paid; Enhancement orc 795.6 > troll 792.8 > tauren 787.4; round5: Restoration orc 723.0 > troll 721.2 > tauren 718.8). Every race line agrees; Skyborne is named only as the paid option.
- **Stat paragraphs** against the band 60 raid tables: no ranking; each reason true on 70755 (spell hit cap 16 = baseline 7 + 9 for both factions; Enhancement hit insignificant at baseline 11 over the 9% melee cap, matching "matters up to the cap").
- **Spell and talent numbers** re-read: Flame Shock 410 mana, Earth Shock 450 ("about twice" a 220 bolt), Mana Spring 100 mana, 10 every 2 s for 5 min; Mana Tide ranks 38/48/58; Riptide ranks 40/50/60; Grace of Air 42/56; Strength of Earth 10; Tremor 18; every talent value as round 13.
- **Source links** (fetched 2026-11-04): Blizzard Deep Dive recap 200, Talents Forever data.json 200, Wowhead Skyborne First Look 200, foreversixty.gg /sim/weights, /planner, /bis/shaman/{elemental,enhancement,restoration} 200; Icy Veins calculator 403 (unknown, not dead).
- **Voice**: no em or en dashes, no site, sim, APL, preset or band talk, no nightly numbers in the new text.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| SHAMAN-R14-1 | CONFIRMED | gear | enhancement.md:111 | "Frozen Heart of the Mountain and Molten Heart of the Mountain are Enchanting trinkets that bind on pickup, so nobody can make them for you." | Live Horde (orc) band 60 raid row: trinket1 Frozen Heart of the Mountain, trinket2 **Rune of the Guard Captain** (19120, quest "Job Opening: Guard Captain of Revantusk Village", bonding pickup), with Molten Heart only as a tied alternative (dps_delta 0). Molten Heart is in the Alliance set only. Round 13 held this as provisional (its Sims requested 4); the re-rank failed and these rows ship. | "Take Enchanting. Frozen Heart of the Mountain is an Enchanting trinket that binds on pickup, and so is Molten Heart of the Mountain, which an Alliance set wears beside it, so nobody can make them for you. Every other craft in your set binds when equipped, so buy them or ask a guildmate; your second profession can be a gathering one that pays for them." |
| SHAMAN-R14-2 | CONFIRMED | gear | enhancement.md:97 | "The Heart of the Mountain trinkets are Enchanting crafts that bind on pickup, so you make those yourself" | Same rows: a Horde set carries one Heart of the Mountain (Frozen) plus Rune of the Guard Captain, so "the trinkets" (plural) is false for half the page's readers. | "Your raid set leans on crafts. Frozen Heart of the Mountain, and on the Alliance Molten Heart of the Mountain beside it, are Enchanting crafts that bind on pickup, so you make those yourself; every other craft in it binds when equipped, so you can buy it." |
| SHAMAN-R14-3 | CONFIRMED | rotation (card note) | `shaman-elemental.json` priorityList.8 | "while leveling the rank for your level: 3 until 33, 4 until 38, ..." | The first `rank-choices.json` row starts at 30; from 20 to 29 there is no row and the line casts the highest rank learned (4 at 20 to 25, 5 at 26 to 29). The card tells a level 20 to 29 player to cast rank 3, which contradicts the data and the page's own "rank 3 from 30 to 33" (elemental.md:116). | "Without raid buffs a cheaper rank, rank 8 at 60, and while leveling the rank for your level: 3 from 30 to 33, 4 until 38, 5 until 44, 6 until 50 and 7 until 56." |
| SHAMAN-R14-4 | CONFIRMED | rotation (card note) | `shaman-restoration.json` priorityList.17 | "While leveling a cheaper rank heals more ...: rank 2 until 25, ..." | The first row starts at 20; below 20 the line casts the top rank learned (3 at 12 to 17, 4 at 18 and 19). The card's "until 25" with no floor tells a level 12 to 19 healer to cast rank 2. The page prose puts the ladder under "Levels 20 to 60", so only the card is wrong. | "... While leveling from 20 a cheaper rank heals more before your mana runs out than your top rank: rank 2 until 25, 3 until 36, 4 at 37, 5 until 43, 6 until 56, 7 at 57 and 58, and 8 at 59." |
| SHAMAN-R14-5 | LIKELY | gear | enhancement.md:97 | "Shamans wear leather until 40 and mail from 40." | In the best-in-slot section this reads as "wear mail at 60", but both raid rows carry leather: Alliance chest Timbermaw Tunic and hands Timbermaw Brawlers; Horde chest Timbermaw Tunic, hands Devilsaur Gauntlets, legs Devilsaur Leggings, waist Defiler's Lizardhide Girdle (`items/shaman.json` armor_type leather). The class rule (mail is learned at 40) is right; the instruction is not. | "Shamans learn mail at 40, but a leather piece with better stats still wins its slot." |

Wording SHOULDs, not findings:

| id | status | kind | file:line | text | note | suggested wording |
|---|---|---|---|---|---|---|
| SHAMAN-R14-S1 | SHOULD | wording | index.md:26 | "they wear leather until 40 and mail from 40" | Same reading as R14-5 on the class page, where it is a "what changed" rule rather than gear advice; Elemental's and Restoration's raid sets are cloth and leather. | "they wear leather until 40 and can wear mail from 40" |
| SHAMAN-R14-S2 | SHOULD | wording | enhancement.md:52 | "9% melee crit for 10 seconds after a spell crits" | Elemental Devastation reads "critical strikes with direct damage spells"; a Flame Shock tick or a Searing Totem crit does not fire it. | "9% melee crit for 10 seconds after a direct damage spell crits" |
| SHAMAN-R13-S1 | SHOULD (carried) | wording | enhancement.md:93 | "past the cap it is wasted" | As round 13. | As round 13. |
| SHAMAN-R12-S1, R9-S1 to S4 | SHOULD (carried) | wording | as filed | | Unchanged. | As filed. |

## Decisions

- **D1. Round 13's Sims requested 2 (Nature's Swiftness save at 44 to 49) is closed.** The Healing Wave row is now scoped to `lines: [17]`, so card .5 falls to the default rule (highest rank learned not above 9): rank 7 at 44 to 47, 8 at 48 and 49. The new `auraIsKnown 16188` gate keeps the line off before 44. The prose ("keep your top rank for Nature's Swiftness", "Nature's Swiftness into your highest rank") now matches what the rotation casts.
- **D2. Round 13's Sims requested 3 (staff) is closed**: both factions' band 60 raid rows take Glowing Brightwood Staff.
- **D3. Round 13's Sims requested 4 becomes R14-1 and R14-2** under the round 14 rule (live BiS ships).
- **D4. Round 13's D3 ("three or four ranks below") is moot**; the prose now quotes the measured rows and they hold on the MaxLevel engine rule.
- **D5. The 7% spell hit baseline for an Orc** (no Artisan's Precision) is the "built in" hit baseline already with the data lane; elemental.md:88 names only Tidal Focus and Artisan's Precision, which is true for a Dwarf and adds up for an Orc with gear. Not re-filed.

## Sims requested

1. Carried (rounds 6 to 13): **shaman-elemental, Earth Shock filler**, band 60, raid preset, sim/cmd/spec-breakdown, 10,000 iterations, seeds 1 and 50001, race dwarf, `-talents 5510331300001351-055000000000000000-0530020000000000`: (a) published rotation; (b) plus an Earth Shock rank 7 (10414) line after Flame Shock, gated on Flame Shock remaining above 6 s and mana percent above the share of the fight left. Report DPS and time out of mana.
2. New, optional (settles which way R14-3 is fixed): **shaman-elemental, Lightning Bolt rank at 20 to 29**, bare preset, bands 20 and 25, leveling build, races dwarf and orc, seeds 1 and 50001, 10,000 iterations: line 8 at (a) the top rank learned against (b) rank 3. If (b) wins beyond error, add a 20 to 29 row and change the prose to "rank 3 from 20 to 33"; if not, apply R14-3's fix as written.

## Sim findings

- Carried: Enhancement band 60 raid hit weight 0 at baseline 11 (spell side between 9% and 16% may be uncounted); Chain Heal party-only jump rule absent from `chain_heal.go`; Stormstrike `RequiredLevel: 40`; haste and the global cooldown (SIM-R4-1); Grace of Air (rank 3 at 60) and Strength of Earth (rank 5 at 60) book ranks; Lightning Overload ratio on the first level-60 log; Air totem raid context; Spellblasting potion on the Elemental pull; Totemic Call positioning (SIM-R5-1); self-dropped Windfury Totem plus Rockbiter not a shipped preset.
- Data lane (already filed): Restoration band 60 Horde rows run as tauren while `recommendedRaces` is [dwarf, orc]; Elemental raid off hand and Enhancement ranged slot; BiS race captions.

## Sources

- Client data 1.60.1.70755: `bis/shaman-enhancement.json` (band 60 raid, orc: Rune of the Guard Captain 19120), `items/shaman.json` (bonding, armor_type), `spellconst/shaman.json` (Lightning Bolt and Healing Wave max_level, cost, base), `trainables/shaman.json`.
- Engine f637ec32c `sim/core/spell.go` `SpellLevelPenalty` and `spell_max_level_auto_gen.go` (no shaman rank but Healing Wave 1 listed).
- ForeverChanges, "WoW Forever downranking: the level penalty and which ranks to keep" https://foreverchanges.pro/downranking
- Blizzard, "World of Warcraft Forever Deep Dive panel recap" https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- Wowhead, "Skyborne First Look" https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
- Talents Forever, data export (beta client data files) https://talentsforever.com/data.json
- Icy Veins, "Start Planning Your WoW Forever Character With Our New Talent Calculator" (403 to automated fetch; unknown, not dead) https://www.icy-veins.com/wow-forever/news/start-planning-your-wow-forever-character-with-our-new-talent-calculator/

## Fix log (round 14)

- SHAMAN-R14-1: fixed, enhancement.md:111. Frozen Heart is on both factions, plus Molten Heart "which an Alliance raid set wears beside it". The Alliance bare row wears Frozen only.
- SHAMAN-R14-2: fixed, enhancement.md:97, same wording.
- SHAMAN-R14-3: fixed, shaman-elemental.json Lightning Bolt note: "3 from 30 to 33 ... 7 until 56 and 8 from 57", which matches the rows.
- SHAMAN-R14-4: fixed, shaman-restoration.json Healing Wave note: "While leveling from 20".
- SHAMAN-R14-5 (LIKELY, verified): fixed, enhancement.md:97: "Shamans learn mail at 40, but a leather piece with better stats still wins its slot." Both raid rows wear leather pieces.
- SHAMAN-R14-S1: applied, index.md:26, "can wear mail from 40". Other SHOULDs: not applied.

All files were already at updated: 2026-11-04. The prose and sections tests pass (145/145).
