# Guide accuracy round 15: Shaman

Auditor: shaman (class index, Elemental, Enhancement, Restoration). Audit tree main 4600b7d82 (guide rounds 1 to 14, sim round 17), client 1.60.1.70755, engine pin b2c4a2d05 (`sim/enginever/version.go`; source read from the gr15 engine checkout at that sha).
Files: `web/src/content/guides/shaman/{index,elemental,enhancement,restoration}.md`, card notes in `data/curated/apl/shaman-{elemental,enhancement,restoration}.json`, rank rows in `data/curated/rank-choices.json`, live BiS `data/builds/1.60.1.70755/bis/shaman-*.json`. Read first: rounds 1 to 14 `shaman.md` (with fix logs), `round4/shaman-levels.md`, `round2..5/race-sims.md`, `ledger.md`, the round 1, 2 and 15 briefs, and `design/reviews/2026-11-04-sim-audit-round17/shaman.md`.

## What moved since round 14 (d7fc51fe7 to 4600b7d82)

- Round 14 fixes: enhancement.md:97 and :111 (Frozen Heart on both factions, Molten Heart on the Alliance raid set; "Shamans learn mail at 40, but a leather piece with better stats still wins its slot"), index.md:26 ("can wear mail from 40"), Elemental card .8 ("3 from 30 to 33 ... 8 from 57"), Restoration card .17 ("While leveling from 20").
- Sim round 17 (new work): Restoration Healing Wave rows on line 17 moved to rank 4 at 37 to 41 and rank 5 at 42 to 46 (R1, d153f2bdf); restoration.md:120 and card .17 now read "rank 4 until 41, rank 5 until 46". Engine: the stat-weight hit window for a physical caster now runs to the spell cap (N1, E 7c8f710d0); published DPS unchanged, the frozen BiS weights not yet regenerated.
- `addon-builds.json`, `talents/`, `trainables/`, `races.json`, `bis/shaman-*`: nothing shaman.

## Builds and shipped leveling order

Re-decoded against `talents/shaman.json` (tab order, as `fs1.ts`) and `addon-builds.json` (index 0 = level 10). All legal, 51 points, prerequisites met (Lightning Mastery <- Call of Thunder, Elemental Fury <- Lightning Overload 3 <- Elemental Alacrity 3, Flurry <- Mental Dexterity 3, Rage of the Farseer <- Mental Quickness 2, Riptide <- Healing Way 3).

- Elemental leveling: Convection 5, Concussion 5, Elemental Devastation 1, Elemental Alacrity 3, Call of Flame 3, Elemental Focus 1, Elemental Warding 3, Call of Thunder 1, Lightning Overload 3, Elemental Fury 5, Lightning Mastery 1 / Thundering Strikes 5 / Tidal Focus 5, Mindfulness 3, Totemic Focus 5, Improved Reincarnation 2. Raid: same 31, Thundering Strikes 5, Ancestral Knowledge 5 / Tidal Focus 5, Mindfulness 3, Improved Reincarnation 2.
- Enhancement leveling: Convection 3, Concussion 2, Elemental Devastation 3, Call of Flame 3 / Thundering Strikes 5, Ancestral Knowledge 5, Mental Dexterity 3, Elemental Weapons 3, Shamanistic Focus 1, Flurry 5, Stormstrike 1, Improved Ghost Wolf 2, Mental Quickness 2, Maelstrom Weapon 5, Rage of the Farseer 1 / Tidal Focus 5, Mindfulness 2. Raid: Concussion 5, Elemental Devastation 3, Reverberation 3, Elemental Focus 1 / same 33 with Ancestral Knowledge 3 / Tidal Focus 5, Mindfulness 3.
- Restoration: Thundering Strikes 5, Ancestral Knowledge 5 / Improved Healing Wave 5, Tidal Focus 5, Mindfulness 3, Ancestral Healing 3, Healing Focus 3, Water Shield 1, Tidal Mastery 5, Restorative Totems 5, Mana Tide Totem 1, Healing Way 3, Nature's Swiftness 1, Purification 5, Riptide 1.
- Orders: Elemental Convection 10-14, Concussion 15-19, ... Lightning Mastery 40, Thundering Strikes 41-45, Tidal Focus 46-50, Mindfulness 51-53, Totemic Focus 54-58, Improved Reincarnation 59-60. Enhancement Thundering Strikes 10-14, Ancestral Knowledge 15-19, Mental Dexterity 20-22, Elemental Weapons 23-24, Stormstrike 25, Elemental Weapons 26, Shamanistic Focus 27, Flurry 28-32, Improved Ghost Wolf 33-34. Restoration Improved Healing Wave 10-14, Mindfulness 15-16, Tidal Focus 17-19, Water Shield 20, Tidal Focus 21-22, Mindfulness 23, Ancestral Healing 24, Mana Tide 25, Ancestral Healing 26-27, Healing Focus 28-30, ... Riptide 50, Thundering Strikes 51-55, Ancestral Knowledge 56-60.

Every "Levels 1 to 19" and "Levels 20 to 60" paragraph and every talents paragraph matches these level for level; none contradicts the header card or the other block.

## What checked out (no finding)

- **Restoration Healing Wave ladder** (restoration.md:120, card .17), new this round: matches `rank-choices.json` line 17 exactly (20-25 r2, 26-36 r3, 37-41 r4, 42-46 r5, 47-56 r6, 57-58 r7, 59 r8). Every listed rank costs less than the top rank at its levels (45/80/155/200/265/340/440 against 155 to 560). The reason given ("a cheaper one heals more before your mana runs out, because your gear adds little healing yet") is the sim round 17 R1 finding (every leveling fight runs dry). Line 73's "never below rank 8" and line 76's low-rank shares are unchanged and still exact on the MaxLevel + 12 rule (every Healing Wave, Lesser Healing Wave and Chain Heal rank has MaxLevel = learn + 5).
- **Elemental bolt ladder** (elemental.md:116, card .8): matches the line 8 rows (30-33 r3 ... 57-59 r8, bare 60 r8); sim round 17 S2b re-probed every level 30 to 59 and every row holds. Card .7 "about 20% more base damage" (353 against 295) and prose :80 "about 16% less" agree.
- **Spell numbers against 70755** (`spellconst`, `spelltips`, `trainables`): Lightning Bolt and Chain Lightning 2.5 s; Chain Lightning r4 510, 30% less a jump, 3 targets, 6 s; Lightning Mastery 15 s every 180 s; Flame Shock r6 292 + 320 over 12 s, 410 mana; Earth Shock r7 531, 450 mana; Searing r6 55 s, 20 yd; Magma r4 20 s, 8 yd; Fire Nova 10 yd of the Fire totem, 10 s; Mana Spring r4 10 mana every 2 s, 5 min, 100 mana, group only; Healing Stream 5 min, every 2 s; Mana Tide 12 s (`${13 sec-1}`), 300 s, 88/133/197/290 at 1/38/48/58; Riptide 445 over 15 s at r1, 12 s cooldown, ranks 40/50/60; Chain Heal r2 315/335, r3 405/490, 50% a jump, "If cast on a party member, the heal will only jump to other party members"; Totemic Call 22, 30 s, 25% base mana, 1.5 s global; Call of the Elements 20; Ghost Wolf 20, 3 s, 40%; Rockbiter r7 554 attack power plus threat; Windfury Weapon r4 disables Windfury Totem for you; Strength of Earth, Grace of Air and Windfury Totem "party or raid members within 30 yards"; Tremor removes Sleep; Cleansing totems every 5 s, party.
- **Talent text**: every value the three talent lists quote matches `talents/shaman.json` (Elemental Fury +100% of the crit bonus, so "crit for double" is right; Healing Way flat 25%; Water Shield 3 globes, 2%, 10 min; Improved Reincarnation 60 to 40 min, +4% health).
- **Races** (`races.json`, `combos.json`): Dwarf is the only Alliance shaman (Orc, Tauren, Troll, Windshaper Skyborne on the Horde; Skyborne shaman Horde only); Artisan's Precision 1% hit with spells and attacks; Stoneform 10% physical, poison, disease and bleed; Blood Fury 10% attack power and spell power, 15 s, 2 min; Berserking 10% casting and attack speed, 15 s, 3 min; Endurance 5% health; Wind Blessed 2% haste. Every race line agrees with the measured rows (round 3 and round 5 race sims) and names Skyborne only as the paid option.
- **Gear and professions against the live BiS rows, both factions**: unchanged from round 14 and still true after its fixes. Elemental: Frozen Heart and Totem of Thunder the only bind-on-pickup crafts. Enhancement: Frozen Heart on both factions, Molten Heart on the Alliance raid row only, every other craft "equip", both main hands two-handers, leather pieces in both raid rows. Restoration: Dormant Heart the only bind-on-pickup craft, Totem of Urgency "equip", Glowing Brightwood Staff on both factions.
- **Consumables** against `sim/request/raid-consumes.json`: every item the three pages name is in its spec's list (Enhancement's pull potion at 55+ is Major Frenzy Potion).
- **Card notes**: all 17 Restoration, 8 Elemental and 8 Enhancement notes re-read against the client text and the prose; no contradiction. Restoration card .5 "Your highest rank" with a rank 9 action is round 10 D2 (card action rank, accepted).
- **Source links** (fetched 2026-11-04): Blizzard Deep Dive recap 200, Talents Forever data.json 200, Wowhead Skyborne First Look 200, foreversixty.gg /sim/weights, /planner, /bis/shaman/{elemental,enhancement,restoration} 200; Icy Veins talent calculator 403 (unknown, not dead). Labels match; the source-link rule holds (weights on the two damage pages only, talent sources open the planner).
- **Voice**: no em or en dashes in the pages or card notes; no site, sim, APL, preset or band talk; no nightly numbers.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| SHAMAN-R15-1 | CONFIRMED | stats | enhancement.md:93 | "Hit covers your swings, your Stormstrike and your shocks at once, and it matters up to the cap; past the cap it is wasted." | An Enhancement shaman has two caps, and the page names one. Sim round 17 N1 (CONFIRMED, E 7c8f710d0): at the raid band's 11% hit, past the 9% weapon cap, Earth Shock, the Maelstrom Lightning Bolts, Searing Totem and Fire Nova (about 22% of the damage) still miss 7.2 to 7.7% of the time, because spells cap at 16% (elemental.md:88, `bis/shaman-elemental.json` hit_to_cap baseline 7 + spell 9). Read as the melee cap, "past the cap it is wasted" is false on 70755; read as the spell cap, "it matters up to the cap" says hit still matters at 11%, which the page's own table (hit insignificant, from the frozen weights) does not show. Round 13 S1 carried this as wording; N1 makes it a fact. | "Hit covers your swings, your Stormstrike and your shocks at once. Your swings stop missing a raid boss at 9% hit and your spells at 16%, so past 9% hit only helps your shocks, totems and free bolts, a small share of your damage." |

Wording SHOULDs, not findings:

| id | status | kind | file:line | text | note | suggested wording |
|---|---|---|---|---|---|---|
| SHAMAN-R15-S1 | SHOULD | leveling wording | restoration.md:120 | "rank 4 until 41, rank 5 until 46" | The sim round 17 reviewer's note: rank 5 is learned at 24 and skipped until 42, so "until 46" can read as "from 24 to 46". | "rank 4 until 41, rank 5 from 42 to 46" (and the same in card .17) |
| SHAMAN-R15-S2 | SHOULD | wording | restoration.md:76 and card .17 | "a rank learned at 23 or earlier gets none of your gear" / "ranks 1 to 4 keep none of it at 60" | Same fact, two phrasings (sim round 17 reviewer). Both true. | Use one phrasing on both. |
| SHAMAN-R14-S2 | SHOULD (carried) | wording | enhancement.md:52 | "9% melee crit for 10 seconds after a spell crits" | Elemental Devastation reads "critical strikes with direct damage spells". | "after a direct damage spell crits" |
| SHAMAN-R13-S1 | superseded | | enhancement.md:93 | | Now SHAMAN-R15-1. | |
| SHAMAN-R12-S1, R9-S1 to S4 | SHOULD (carried) | wording | as filed | | Unchanged. | As filed. |

## Decisions

- **D1. The carried "Sims requested 1" (Earth Shock filler for Elemental, rounds 6 to 14) is closed** by sim round 17 E1: Earth Shock rank 7 after Flame Shock loses 1.5% gated and 5.4% ungated. elemental.md:72 ("Earth Shock stays on your bar for interrupts ... never let it delay a Flame Shock") already says what the measurement says. No change.
- **D2. Restoration ladder change accepted as written.** The new rows reproduce at 10,000 iterations on both seeds and factions (sim round 17 R1); the prose and card list them exactly. Only the "from 42" wording is a SHOULD (R15-S1).
- **D3. Enhancement's empty relic slot (sim round 17 N4)** is not a guide finding: no Enhancement prose names a relic or says the slot is empty.
- **D4. The fix for R15-1 keeps the table and the prose consistent**: it does not say hit is worth a lot past 9%, only that it is not wasted, which holds both before and after the refresh regenerates the hit weight with the spell-cap window.
- **D5. Healing Wave rank 10 vs "highest rank"** stays as round 10 D2.

## Sims requested

1. Carried from round 14 (optional, not taken up by sim round 17, whose S2b probe started at 30): **shaman-elemental, Lightning Bolt rank at 20 to 29**, bare preset, bands 20 and 25, leveling build, races dwarf and orc, seeds 1 and 50001, 10,000 iterations: line 8 at (a) the top rank learned against (b) rank 3. If (b) wins beyond error, add a 20 to 29 row and change elemental.md:116 to "rank 3 from 20 to 33"; if not, no change.
2. At the refresh (sim round 17 "Measure at refresh"): regenerate the Enhancement band 60 hit weight with E 7c8f710d0, so the page's table stops showing hit as worthless at 11%. Not a page edit.

## Sim findings

- Carried: Chain Heal party-only jump rule absent from `chain_heal.go` (the 70755 description states it); Stormstrike `RequiredLevel: 40`; haste and the global cooldown (SIM-R4-1); Grace of Air rank 3 and Strength of Earth rank 5 book ranks (pending owner, sim round 17 N5); Lightning Overload ratio on the first level-60 log; Spellblasting potion on the Elemental pull; Totemic Call positioning (SIM-R5-1); self-dropped Windfury Totem plus Rockbiter not a shipped preset; Healing Wave rank 10 treated as a book rank by the NS save.
- Closed by sim round 17: Enhancement hit weight 0 past the weapon cap (N1, engine side; the weights regenerate at refresh).
- Data lane (already filed): Restoration band 60 Horde rows run as tauren while `recommendedRaces` is [dwarf, orc]; BiS race captions; `spellranks.json` still carries Lava Burst (sim round 17 X2).

## Sources

- Client data 1.60.1.70755: `talents/shaman.json`, `trainables/shaman.json`, `spellconst/shaman.json` (Healing Wave, Lesser Healing Wave and Chain Heal max_level; Lightning Bolt and Chain Lightning rows), `spelltips/shaman.json` (Chain Heal 10623, Chain Lightning 10605, Rockbiter 16316, Windfury Weapon 16362, Mana Tide 1325256, Totemic Call 437009), `races.json`, `combos.json`, `bis/shaman-*.json` (hit_to_cap, weights, slots), `addon-builds.json`.
- `data/curated/rank-choices.json` (shaman-elemental line 8, shaman-restoration line 17); `sim/request/raid-consumes.json`.
- Sim-accuracy round 17 shaman report, `design/reviews/2026-11-04-sim-audit-round17/shaman.md` (R1, N1, N4, E1, S2b, S2c).
- Blizzard, "World of Warcraft Forever Deep Dive panel recap" https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- Wowhead, "Skyborne First Look" https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
- Talents Forever, data export (beta client data files) https://talentsforever.com/data.json
- Icy Veins, "Start Planning Your WoW Forever Character With Our New Talent Calculator" (403 to automated fetch; unknown, not dead) https://www.icy-veins.com/wow-forever/news/start-planning-your-wow-forever-character-with-our-new-talent-calculator/

## Fix log (round 15)

- SHAMAN-R15-1 fixed, enhancement.md:93: names both caps, 9% for melee and 16% for spells.
- SHAMAN-R15-S1 fixed: restoration.md:120 and the shaman-restoration.json line 17 note now say "rank 5 from 42 to 46". Checked against rank-choices.json.
- Other SHOULDs left.
