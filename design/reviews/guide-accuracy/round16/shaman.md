# Guide accuracy round 16: Shaman

Auditor: shaman (class index, Elemental, Enhancement, Restoration). Audit tree land/guides-r15 2c3135a2e (guide rounds 1 to 15 on launch-state main), client 1.60.1.70755, engine pin b2c4a2d05 (`sim/enginever/version.go`).
Files: `web/src/content/guides/shaman/{index,elemental,enhancement,restoration}.md`, card notes in `data/curated/apl/shaman-{elemental,enhancement,restoration}.json`, rank rows in `data/curated/rank-choices.json`, live BiS `data/builds/1.60.1.70755/bis/shaman-*.json`. Read first: rounds 1 to 15 `shaman.md` (round 14 and 15 with fix logs), `round4/shaman-levels.md`, `ledger.md` through round 15, the round 1, 2 and 16 briefs.

## What moved since round 15 (4600b7d82 to 2c3135a2e)

- `enhancement.md:93` (R15-1 fix, new work): "Hit covers your swings, your Stormstrike and your shocks at once. Your swings stop missing a raid boss at 9% hit and your spells at 16%, so past 9% hit only helps your shocks, totems and free bolts, a small share of your damage."
- `restoration.md:120` and `shaman-restoration.json` priorityList[17] (R15-S1): "rank 5 from 42 to 46".
- Nothing else shaman: `addon-builds.json`, `talents/`, `trainables/`, `spellconst/`, `spelltips/`, `races.json`, `rank-choices.json`, `bis/shaman-*` and the other two APLs are unchanged.

## Builds and shipped leveling order

Re-decoded against `talents/shaman.json` (tab order) and `addon-builds.json` (index 0 = level 10; coordinates are tree, tier+1, column+1). All three builds legal at 51 points with prerequisites met; the talent lists are as round 15 records them.

- Elemental leveling: Convection 10-14, Concussion 15-19, Elemental Devastation 20, Elemental Alacrity 21-23, Call of Flame 24-26, Elemental Focus 27, Elemental Warding 28-30, Call of Thunder 31, Lightning Overload 32-34, Elemental Fury 35-39, Lightning Mastery 40, Thundering Strikes 41-45, Tidal Focus 46-50, Mindfulness 51-53, Totemic Focus 54-58, Improved Reincarnation 59-60.
- Enhancement leveling: Thundering Strikes 10-14, Ancestral Knowledge 15-19, Mental Dexterity 20-22, Elemental Weapons 23-24, Stormstrike 25, Elemental Weapons 26, Shamanistic Focus 27, Flurry 28-32, Improved Ghost Wolf 33-34, Mental Quickness 35-36, Maelstrom Weapon 37-39, Rage of the Farseer 40, Maelstrom Weapon 41-42, then Convection, Concussion, Elemental Devastation, Call of Flame, Tidal Focus, Mindfulness.
- Restoration: Improved Healing Wave 10-14, Mindfulness 15-16, Tidal Focus 17-19, Water Shield 20, Tidal Focus 21-22, Mindfulness 23, Ancestral Healing 24, Mana Tide 25, Ancestral Healing 26-27, Healing Focus 28-30, Tidal Mastery 31-35, Restorative Totems 36-40, Healing Way 41-43, Nature's Swiftness 44, Purification 45-49, Riptide 50, Thundering Strikes 51-55, Ancestral Knowledge 56-60.

Every "Levels 1 to 19" and "Levels 20 to 60" paragraph and every talents paragraph matches these orders level for level, and nothing contradicts the header card or the block above or below it.

## What checked out (no finding)

- **R15-1 fix (enhancement.md:93), new work.** The 9% figure is the site's own melee cap (`web/src/lib/bis/types.ts` `BisMeleeHitToCap.specials`, "the 9% cap against a level-63 target"; 300 skill against 315 defense, no dual wield, so white and yellow share it). 16% is the spell cap the Elemental page states (17% base miss, 1% always left). The Enhancement raid rows sit at baseline 11, past the melee cap and short of the spell cap, so "past 9% hit only helps your shocks, totems and free bolts" agrees with sim round 17 N1 and with the table, which marks hit insignificant on the frozen weights. Tidal Focus's 70755 text ("improves your chance to hit by 5%", no school or weapon limit) supports "covers your swings ... and your shocks at once". The wording on Fire Nova is a SHOULD (S1 below), not a finding.
- **R15-S1 fix**: the "rank 5 from 42 to 46" wording on the page and the card matches `rank-choices.json` shaman-restoration (42-46 r5) and no longer reads as "from 24".
- **Spell numbers re-derived from 70755** (`spellconst`, `spelltips`, `trainables`): Lightning Bolt r10 220 mana and 353, r9 190 and 295 (16.4% less, card's "about 20% more" = 1.197), ranks 3 to 10 all 2.5 s with a 0.714 coefficient; Chain Lightning r4 510, 30% less a jump, 6 s; Earth Shock r7 450 mana ("about twice" 220); Flame Shock r6 292 + 320 over 12 s; Searing r6 55 s, 20 yd; Magma 20 s, 8 yd; Fire Nova 10 yd of the Fire totem, 10 s, learned 12; Mana Spring r4 10 every 2 s, 5 min, 100 mana; Healing Stream every 2 s, 5 min; Mana Tide 88/133/197/290 at 1/38/48/58, 300 s cooldown; Riptide 445 over 15 s (805 at r3), 12 s cooldown, ranks 40/50/60, nothing up front; Chain Heal r1 251/260, r2 335/315, r3 490/405, 50% a jump, party-only rule; Totemic Call 22, 25% base mana, 30 s; Ghost Wolf 20, 3 s, 40%, outdoors; Reincarnation 60 min (40 with Improved Reincarnation); Rockbiter r7 554 attack power plus threat; Windfury Weapon r4 disables Windfury Totem for you; Strength of Earth, Grace of Air, Windfury Totem "party or raid members within 30 yards"; Tremor removes Sleep; cleansing totems every 5 s, party; Stormstrike 6 s.
- **Low-rank arithmetic** (restoration.md:76, :80, card .17) on MaxLevel = learn + 5, full to MaxLevel + 12, 5% a level after: Healing Wave 7 and Chain Heal 1 85% at 60, Lesser Healing Wave 3 65%, Healing Wave 6 45%, rank 5 (24) 5%, anything learned at 23 or earlier 0; ranks 1 to 4 none. Chain Heal per mana: rank 3 leads with no gear, rank 2 leads past about 290 healing power ("with a raid's healing gear"); rank 1 at 85% of the bonus only passes rank 2 per mana past about 1,400 healing power, so "never better than rank 2" holds for every set the game offers.
- **Elemental bolt ladder** (elemental.md:116, card .8): matches the rows (30-33 r3 ... 57-59 r8, 60 bare r8).
- **Talent text**: every value the three talent lists and the leveling paragraphs quote matches `talents/shaman.json`.
- **Gear and professions against the live BiS rows, both factions**: Elemental Frozen Heart and Totem of Thunder the only bind-on-pickup crafts (Enchanting), every other craft "equip". Enhancement Frozen Heart on both factions, Molten Heart on the Alliance raid row only (a tied alternative on the Horde), every other craft "equip", leather Timbermaw Tunic, Brawlers and Devilsaur pieces in the raid rows, two-handers in the main hand. Restoration Dormant Heart the only bind-on-pickup craft, Totem of Urgency and the tailoring and leatherworking pieces "equip", Glowing Brightwood Staff both factions.
- **Races** against `races.json` and the measured rows: Blood Fury 10% attack power and spell power (covers heals, aura 317), Berserking 10% casting and attack speed, 15 s, 3 min, Stoneform 10% physical for 8 s plus bleed, poison and disease, Endurance 5%, Artisan's Precision 1% spells and attacks, Wind Blessed 2% haste; Skyborne named only as the paid option; Alliance Skyborne cannot be shamans.
- **Consumables and enchants**: unchanged since round 15 and still in the raid consumes list; the imbue-as-temporary-enchant and "imbues only proc on melee swings" lines hold on the 70755 imbue tooltips.
- **Card notes**: all 9 Elemental, 7 Enhancement and 10 Restoration notes re-read against the client text, the APL conditions and the prose; no contradiction.
- **Source links** (fetched 2026-11-04): Blizzard Deep Dive recap 200, Talents Forever data.json 200, Wowhead Skyborne First Look 200, foreversixty.gg /sim/weights, /planner, /bis/shaman/{elemental,enhancement,restoration} 200; Icy Veins calculator 403 (unknown, not dead). Labels match, and the source-link rule holds.
- **Voice**: no em or en dashes in the pages or card notes; no site, sim, APL, preset or band talk; no nightly numbers.

## Findings

None. Zero CONFIRMED, zero LIKELY.

Wording SHOULDs, not findings:

| id | status | kind | file:line | text | note | suggested wording |
|---|---|---|---|---|---|---|
| SHAMAN-R16-S1 | SHOULD | stats wording | enhancement.md:93 | "past 9% hit only helps your shocks, totems and free bolts" | Fire Nova is a spell that bursts from the totem, not totem damage, and the raid rotation casts it on Clearcasting (card .6), so a strict reader finds it missing from an "only" list. Chain Lightning is left out too, but only the pack advice casts it, never a raid boss. Not false in the boss context this sentence sets. | "past 9% hit only helps your shocks, Fire Nova, totems and free bolts, a small share of your damage." |
| SHAMAN-R14-S2 | SHOULD (carried) | wording | enhancement.md:52 | "9% melee crit for 10 seconds after a spell crits" | Elemental Devastation reads "critical strikes with direct damage spells". | "after a direct damage spell crits" |
| SHAMAN-R15-S2 | SHOULD (carried) | wording | restoration.md:76 and card .17 | two phrasings of the same zero-bonus fact | Both true. | One phrasing on both. |
| SHAMAN-R12-S1, R9-S1 to S4 | SHOULD (carried) | wording | as filed | | Unchanged. | As filed. |

## Decisions

- **D1. R15-1's fix stands.** It names both caps, matches the table at the shipped baseline, and stays true after the refresh regenerates the hit weight with the spell-cap window (it says hit past 9% helps a small share, not that it is worth a lot or nothing).
- **D2. Grace of Air rank 3 and Strength of Earth rank 5** (Enhancement card .2 "Learned at 42, rank 2 at 56"; trainables carry 25359 and 25361 at 60) stay with the sim lane and the owner (sim round 17 N5). Not re-filed.
- **D3. Stormstrike's `RequiredLevel: 40`** against "you can press it the moment you learn it" (enhancement.md:121) stays a carried sim finding.
- **D4. Restoration's statPriority leaves out spell haste** though the band 60 table weights it: no item carries haste in any shaman band (`haste_on_items` false), and every healer page does the same. Not a finding.

## Sims requested

1. Carried, optional: **shaman-elemental, Lightning Bolt rank at 20 to 29**, bare preset, bands 20 and 25, leveling build, races dwarf and orc, seeds 1 and 50001, 10,000 iterations: line 8 at (a) the top rank learned against (b) rank 3. If (b) wins beyond error, add a 20 to 29 row and change elemental.md:116 to "rank 3 from 20 to 33"; if not, no change.
2. Carried, at the refresh: regenerate the Enhancement band 60 hit weight with E 7c8f710d0. Not a page edit.

## Sim findings

- Carried: Chain Heal party-only jump rule absent from `chain_heal.go`; Stormstrike `RequiredLevel: 40`; haste and the global cooldown (SIM-R4-1); Grace of Air rank 3 and Strength of Earth rank 5 book ranks (N5); Lightning Overload ratio on the first level-60 log; Spellblasting potion on the Elemental pull; Totemic Call positioning (SIM-R5-1); self-dropped Windfury Totem plus Rockbiter not a shipped preset; Healing Wave rank 10 treated as a book rank by the Nature's Swiftness save.
- Data lane (already filed): Restoration band 60 Horde rows run as tauren while `recommendedRaces` is [dwarf, orc]; BiS race captions; the 7% spell-hit baseline for an Orc; `spellranks.json` still carries Lava Burst; Enhancement hit weight.

## Sources

- Client data 1.60.1.70755: `talents/shaman.json`, `trainables/shaman.json`, `spellconst/shaman.json` (Lightning Bolt, Chain Lightning, Healing Wave, Lesser Healing Wave, Chain Heal, Riptide, totem rows), `spelltips/shaman.json` (Rockbiter 8018/8019/16316, Flametongue 8024/8027/8030/16342, Windfury Weapon 16362, Windfury Totem 10614, Grace of Air 25359, Strength of Earth 25361, Mana Tide 16190/17354/17359/1325256, Chain Heal 1064/10623, Riptide 408521/1239243, Totemic Call 437009, Lightning Shield 10432, Tremor 8143), `races.json`, `addon-builds.json`, `bis/shaman-*.json` (band 60 rows, hit_to_cap, haste_on_items), `items/shaman.json` (bonding, armor_type).
- `data/curated/rank-choices.json` (shaman-elemental line 8, shaman-restoration line 17); `web/src/lib/bis/types.ts` and `hit-cap.ts` (melee 9% and spell 16% caps).
- Blizzard, "World of Warcraft Forever Deep Dive panel recap" https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- Wowhead, "Skyborne First Look" https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
- Talents Forever, data export (beta client data files) https://talentsforever.com/data.json
- Icy Veins, "Start Planning Your WoW Forever Character With Our New Talent Calculator" (403 to automated fetch; unknown, not dead) https://www.icy-veins.com/wow-forever/news/start-planning-your-wow-forever-character-with-our-new-talent-calculator/

## Fix log (round 16)

- SHAMAN-R16-S1: fixed, enhancement.md:93: Fire Nova added to the list.
