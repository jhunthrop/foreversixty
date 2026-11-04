# Guide accuracy round 12: Shaman

Auditor: shaman (class index, Elemental, Enhancement, Restoration). Audit tree main a8b147a6f, client 1.60.1.70755, engine pin 9f7f1f6e4 (`sim/enginever/version.go`).
Files: `web/src/content/guides/shaman/{index,elemental,enhancement,restoration}.md`, card notes in `data/curated/apl/shaman-{elemental,enhancement,restoration}.json`, rank rows in `data/curated/rank-choices.json`. Read first: round1 to round11 `shaman.md` with every fix log, `round4/shaman-levels.md`, `ledger.md`, the round 2 and round 1 briefs.

## What moved since round 11 (d54de48aa to a8b147a6f)

- 91da94f1e (round 11 landing): restoration.md:72 (Chain Heal party clause, R11-2), `shaman-restoration.json` priorityList.16 note (R11-1), index.md:26 (R11-S1, "your weapon imbue goes on that one weapon"), and two Enhancement edits the round 11 fix log does not list: enhancement.md:10 description ("one weapon and your imbue") and enhancement.md:30 overview ("one weapon and its imbue ... whenever Maelstrom Weapon reaches five stacks"). The fix log says R11-S2 was left; the landed text applies a shorter version of it. Both read true on 70755.
- 3bd24d3a2 (guide-header-build): the rail build card, key-talent chips and rotation card now follow `raidBuild` where one exists (Elemental, Enhancement); Restoration has no `raidBuild` and is unchanged.
- No change to `addon-builds.json`, `talents/`, `trainables/`, `spelltips/`, `races.json`, `stat-weights.json`, `rank-choices.json` or `bis/shaman-*` (still de044125c).

## Builds and shipped leveling order

Re-decoded both strings per page against `talents/shaman.json` (tier gates and prerequisites checked: Lightning Overload needs Elemental Alacrity 3, Elemental Fury needs Lightning Overload 3, Lightning Mastery needs Call of Thunder, Flurry needs Mental Dexterity 3, Rage of the Farseer needs Mental Quickness 2, Riptide needs Healing Way 3; all legal, 51 points each):

- Elemental leveling `5510331300001351/05/053052`: Convection 5, Concussion 5, Elemental Devastation 1, Elemental Alacrity 3, Call of Flame 3, Elemental Focus 1, Elemental Warding 3, Call of Thunder 1, Lightning Overload 3, Elemental Fury 5, Lightning Mastery 1; Thundering Strikes 5; Tidal Focus 5, Mindfulness 3, Totemic Focus 5, Improved Reincarnation 2. Raid `/055/053002`: same Elemental, Thundering Strikes 5, Ancestral Knowledge 5; Tidal Focus 5, Mindfulness 3, Improved Reincarnation 2.
- Enhancement leveling `323003/055030310051202051/052` (11/33/7) and raid `0533001/053030310051202051/053` (12/31/8): as round 10 and 11.
- Restoration `0/055/5530003315513151` (0/10/41): as round 11.

`addon-builds.json` orders are byte-identical to round 11 (Elemental Convection 10-14 ... Improved Reincarnation 59-60; Enhancement Thundering Strikes 10-14, Ancestral Knowledge 15-19, Mental Dexterity 20-22, Elemental Weapons 23-24, Stormstrike 25, ...; Restoration Improved Healing Wave 10-14, Mindfulness 15-16, Tidal Focus 17-19, Water Shield 20, ...). Every "Levels 1 to 19" and "Levels 20 to 60" paragraph matches them level for level, and none contradicts the other.

## What checked out (no finding)

- **Round 11 fixes.** restoration.md:72 now states the client rule exactly ("If cast on a party member, the heal will only jump to other party members", Chain Heal 1064/10622/10623 in `spelltips/shaman.json`) and draws no wrong conclusion; it agrees with :37, :38 and :82 (group with the tank and melee for the party-only totems). Card .16 matches the tooltip and the prose. index.md:26 and the two new Enhancement summary lines are true for both imbues.
- **Header card vs prose (guide-header-build).** Elemental and Enhancement now show the raid build in the rail. Every prose line that names a talent says which build has it (elemental.md:52, :66, :88 Ancestral Knowledge; enhancement.md:52, :55 "3 in the raid build", :83 "In the raid build, Elemental Focus"); the leveling paragraphs say "this build" of the leveling order, which `addon-builds.json` carries for both. Nothing contradicts the raid card. Restoration is unchanged (no `raidBuild`).
- **Talent text at max rank** re-read for every talent the pages describe: Convection 10%, Concussion 5% (LB, CL, Earth Shock), Elemental Devastation 9% melee crit 10 s, Reverberation 1.0 s at 5 (0.6 s at the raid build's 3), Elemental Alacrity 0.50 s, Call of Flame +15% and +9 s, Elemental Focus 10%, Elemental Warding 10%, Call of Thunder 5%, Lightning Overload 30% half damage no threat, Elemental Fury +100% of the crit bonus (spell crits at 200%, so "crit for double" and "doubles the crit bonus" hold under the round 9 rule), Lightning Mastery 0.5 s and no Chain Lightning cooldown for 15 s on a 3 min cooldown; Thundering Strikes 5%, Ancestral Knowledge 10%, Mental Dexterity 150% of Intellect, Elemental Weapons 20% Rockbiter / 40% Windfury, Shamanistic Focus 45% (shocks and Lightning Shield), Flurry 30% for 3 swings, Stormstrike 150% / +20% / 6 s, Improved Ghost Wolf 3.0 s and indoors, Mental Quickness 30% of Intellect, Maelstrom Weapon 20% a stack, 5 stacks, 30 s, Rage of the Farseer 30% for 25 s on 3 min; Improved Healing Wave 0.5 s, Tidal Focus 5%, Mindfulness 50%, Totemic Focus 25%, Improved Reincarnation -20 min (60 to 40), +4% health, Ancestral Healing 25% armor 15 s, Healing Focus 70%, Water Shield 3 globes 2% max mana, Tidal Mastery 5%, Restorative Totems 25%/50%, Mana Tide 12 s on 5 min, Healing Way 25%, Nature's Swiftness 3 min, Purification 10%, Riptide 15 s HoT, 12 s cooldown, +25% Chain Heal on its target.
- **Spell numbers** (`spelltips/shaman.json`): Flame Shock 12 s base (21 s with Call of Flame, 3.5 shock cooldowns of 6 s); Earth Shock 7 450 mana against Lightning Bolt 10 220 ("about twice"); Lightning Bolt 8/9/10 160/190/220 mana, 240/295/353 base ("about 16% less", card .7 "about 20% more"); Chain Lightning 4 510 base on a 6 s cooldown; Searing Totem 6 20 yards, 55 s; Magma Totem 20 s, 8 yards; Fire Nova 10 yards of the Fire totem, 10 s; Mana Spring 4 10 mana every 2 s for 5 min to group members; Healing Stream 5 min, group members; Mana Tide 12 s, group members; Totemic Call 25% of base mana, 30 s cooldown; Ghost Wolf 3 s cast, 40%; Reincarnation 60 min; Chain Heal 1/2/3 260/315/405 mana, 251/335/490; Riptide 3 230 mana.
- **Trainer levels** (`trainables/shaman.json`) for every spell named: Rockbiter 1, Earth Shock 4, Lightning Shield 8, Flame Shock, Searing, Strength of Earth, Flametongue 10, Fire Nova 12, Cure Poison 16, Tremor 18, Healing Wave 6/12/18, Ghost Wolf, Call of the Elements, Lesser Healing Wave, Healing Stream 20, Totemic Call, Cure Disease, Poison Cleansing 22, Mana Spring and Magma 26, Windfury Weapon 30, Chain Lightning and Windfury Totem 32, Disease Cleansing 38, Chain Heal 40, Grace of Air 42/56, Mana Tide 38/48/58, Riptide 40/50/60, Lightning Bolt 8 at 44.
- **Level-gap arithmetic** (restoration.md:76): learned at 23 or earlier keeps nothing at 60; Healing Wave 7 / Chain Heal 1 (40) keep 85%, Lesser Healing Wave 3 (36) 65%, Healing Wave 6 (32) 45%. Correct.
- **Card notes** (all three specs) re-read against the actions, conditions and client values: Elemental .0 to .8, Enhancement .1 to .7 (Grace of Air "Learned at 42, rank 2 at 56"), Restoration .0 to .7, .16, .17 (Riptide under 90% on the tank; everyone else Riptide under 50% and Lesser Healing Wave under 25%, both gated on mana above half, as restoration.md:78 says).
- **Best in slot bind facts** against `items/shaman.json` and the band 60 raid rows: Frozen Heart of the Mountain, Molten Heart of the Mountain, Dormant Heart of the Mountain and Totem of Thunder are Enchanting crafts with `bonding: pickup`; every other craft in the three raid sets binds on equip, including Restoration's Totem of Urgency (Enchanting, `bonding: equip`), so "every other craft in it binds when equipped" holds on all three pages.
- **Races** (`races.json`, `combos.json`): shaman combos are Orc, Tauren, Troll, Dwarf (new) and Windshaper Skyborne (new, Horde); Blood Fury "Attack Power and Spell Power by 10% for 15 sec" (Wowhead Forever 20572), 2 min; Berserking 10% casting and attack speed, 15 s, 3 min; Stoneform 10% physical, Bleed, Poison, Disease, 8 s; Artisan's Precision 1% with spells and attacks; Endurance 5%. No race line contradicts a measured result.
- **Consumables** against `sim/request/raid-consumes.json`: as round 11.
- **Voice**: no em or en dashes, no site, sim, APL, preset or band talk, no nightly numbers, no stat ranking in prose.
- **Source links** (fetched 2026-11-04): Blizzard Deep Dive recap 200, Talents Forever data.json 200, Wowhead Skyborne First Look 200, foreversixty.gg /sim/weights, /planner, /bis/shaman/{elemental,enhancement,restoration} 200; Icy Veins calculator 403 to automated fetches (unknown, not dead). Each label matches the page it opens.

## Findings

None CONFIRMED, none LIKELY.

Wording SHOULDs, not findings:

| id | status | kind | file:line | text | note | suggested wording |
|---|---|---|---|---|---|---|
| SHAMAN-R12-S1 | SHOULD | wording | `data/curated/apl/shaman-elemental.json` priorityList.6 | "on a single boss your Searing Totem stands 20 yards out" | Card .3 and elemental.md:76/:78 say within (up to) 20 yards; "stands 20 yards out" reads as a fixed distance. | "on a single boss your Searing Totem can stand up to 20 yards out, and Lightning Bolt is the better cast." |
| SHAMAN-R12-S2 | SHOULD | wording | enhancement.md:68 | "Rage of the Farseer, trinkets, Blood Fury and your potion go together on the pull" | The Enchants and consumables paragraph (:101) names no potion for Enhancement, and the raid set carries none; the line is not false, but the reader is left asking which potion. | "Rage of the Farseer, trinkets and Blood Fury go together on the pull and again on cooldown. Then Stormstrike." |
| R9-S1 to R9-S4 | SHOULD (carried) | wording | as round 9 | Elemental Flame Shock card "nothing else wants the shock cooldown"; Rockbiter "below 20" (enhancement.md:117); "which behind a boss is never" (enhancement.md:89); Flame Shock "ahead of everything but your cooldowns" (elemental.md:72). | Unchanged. | As round 9. |

## Decisions

- **D1. Round 11's two unlisted Enhancement summary edits stand.** "one weapon and its imbue" and "your imbue" are true under both Rockbiter (raid with a Windfury Totem) and Windfury Weapon (everywhere else), and agree with enhancement.md:79 and :101.
- **D2. Blood Fury's spell power on heals (restoration.md:106).** The 70755 tooltip says "Spell Power", the term the Forever client uses for damage and healing alike; no client text or Forever source limits it to damage. Kept.
- **D3. Fire Nova Totem and Earth Shield in `trainables/shaman.json`** carry `source: class_spell`, not `skill_line_ability`; per the trainables rule (SkillLineAbility decides) neither is learnable, so index.md:29 ("not a totem of its own") holds.

## Sims requested

1. Carried (rounds 6 to 11): **shaman-elemental, Earth Shock filler**, band 60, raid preset, sim/cmd/spec-breakdown, 10,000 iterations, seeds 1 and 50001, race dwarf, `-talents 5510331300001351-055000000000000000-0530020000000000`: (a) published rotation; (b) plus an Earth Shock rank 7 (10414) line right after Flame Shock, gated on Flame Shock remaining above 6 s and mana percent above the share of the fight left. Report DPS and time out of mana.
2. Carried (rounds 10 and 11): **shaman-restoration, Nature's Swiftness save at 44 to 49**, bare preset, levels 45 and 49, heal profile, seeds 1 and 50001, races tauren and dwarf, leveling build: (a) as shipped (save casts rank 5 through the shared 10396 row) against (b) the save at the top rank learned (7 at 44 to 47, 8 at 48 and 49). If (b) wins or ties, scope the row to the filler line.
3. Provisional (waits for the final best-in-slot re-rank): restoration.md:96 "at 60 the staff usually wins" agrees with the current band 60 rows (Glowing Brightwood Staff); recheck after the re-rank lands.

## Sim findings

- Carried: Chain Heal party-only jump rule absent from `chain_heal.go` (matters only if a profile groups the healer with the tank); Stormstrike `RequiredLevel: 40`; haste and the global cooldown (SIM-R4-1); Grace of Air and Strength of Earth book ranks; Lightning Overload ratio on the first level-60 log; Air totem raid context; Spellblasting potion on the Elemental pull; Totemic Call positioning (SIM-R5-1); self-dropped Windfury Totem plus Rockbiter not a shipped preset.
- Data lane (already filed): Restoration band 60 Horde rows run as tauren while `recommendedRaces` is [dwarf, orc]; Elemental raid off hand and Enhancement ranged slot.

## Sources

- Wowhead Forever / client, Blood Fury tooltip https://nether.wowhead.com/forever/tooltip/spell/20572
- Wowhead Forever / client, Chain Heal tooltip (text as in `spelltips/shaman.json`) https://nether.wowhead.com/forever/tooltip/spell/10623
- Blizzard, "World of Warcraft Forever Deep Dive panel recap" https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- Wowhead, "Skyborne First Look" https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
- Talents Forever, data export (beta client data files) https://talentsforever.com/data.json
- Icy Veins, "Start Planning Your WoW Forever Character With Our New Talent Calculator" (403 to automated fetch; unknown, not dead) https://www.icy-veins.com/wow-forever/news/start-planning-your-wow-forever-character-with-our-new-talent-calculator/
