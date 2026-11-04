# Guide accuracy round 11: Shaman

Auditor: shaman (class index, Elemental, Enhancement, Restoration). Audit tree main d54de48aa, client 1.60.1.70755, engine pin 9f7f1f6e4 (`sim/enginever/version.go`).
Files: `web/src/content/guides/shaman/{index,elemental,enhancement,restoration}.md`, card notes in `data/curated/apl/shaman-{elemental,enhancement,restoration}.json`, rank rows in `data/curated/rank-choices.json`. Read first: round1 to round10 `shaman.md` with every fix log, `round4/shaman-levels.md`, round2 to round5 `race-sims.md`, `ledger.md` (round 10 section included).

## What moved since round 10 (84460cec8 to d54de48aa)

Only commit eabba8f84 (round 10 landing) touched shaman files. It applied SHAMAN-R10-1 to R10-4, the Rockbiter-under-Windfury-Totem prose (enhancement.md:57, :66, :79, :81, :101, :121; card priorityList.2), the Elemental raid rank 10 prose (elemental.md:80; cards .7 and .8), and one change no shaman report lists: restoration.md:72 gained "Cast on someone in your party, it only jumps to your party, so ask to be grouped with the tank." The ledger names it as a rendered-review MUST ("Chain Heal's party-only jumps"). `addon-builds.json`, `talents/`, `trainables/`, `spelltips/`, `races.json`, `stat-weights.json`, `bis/` (still de044125c) and `rank-choices.json` did not change.

## Builds and shipped leveling order

Re-decoded `addon-builds.json` (index 0 = level 10) against `talents/shaman.json`; identical to round 10:

- Elemental: Convection 10-14, Concussion 15-19, Elemental Devastation 20, Elemental Alacrity 21-23, Call of Flame 24-26, Elemental Focus 27, Elemental Warding 28-30, Call of Thunder 31, Lightning Overload 32-34, Elemental Fury 35-39, Lightning Mastery 40, Thundering Strikes 41-45, Tidal Focus 46-50, Mindfulness 51-53, Totemic Focus 54-58, Improved Reincarnation 59-60.
- Enhancement: Thundering Strikes 10-14, Ancestral Knowledge 15-19, Mental Dexterity 20-22, Elemental Weapons 23-24, Stormstrike 25, Elemental Weapons 26, Shamanistic Focus 27, Flurry 28-32, Improved Ghost Wolf 33-34, Mental Quickness 35-36, Maelstrom Weapon 37-39, Rage of the Farseer 40, Maelstrom Weapon 41-42, Convection 43-45, Concussion 46-47, Elemental Devastation 48-50, Call of Flame 51-53, Tidal Focus 54-58, Mindfulness 59-60.
- Restoration: Improved Healing Wave 10-14, Mindfulness 15-16, Tidal Focus 17-19, Water Shield 20, Tidal Focus 21-22, Mindfulness 23, Ancestral Healing 24, Mana Tide Totem 25, Ancestral Healing 26-27, Healing Focus 28-30, Tidal Mastery 31-35, Restorative Totems 36-40, Healing Way 41-43, Nature's Swiftness 44, Purification 45-49, Riptide 50, Thundering Strikes 51-55, Ancestral Knowledge 56-60.

Frontmatter strings: Elemental `5510331300001351/05/053052` and `/055/053002` (31 Elemental each, totals 51); Enhancement `323003/055030310051202051/052` (11/33/7) and `0533001/053030310051202051/053` (12/31/8); Restoration `0/055/5530003315513151` (0/10/41). Every Leveling paragraph (1 to 19, then 20 to 60) matches these orders, and the round 10 rewrite of enhancement.md:121 is exact (two points of Elemental Weapons, Stormstrike 25, Elemental Weapons, Shamanistic Focus 27, Flurry by 32, Improved Ghost Wolf 33 and 34). enhancement.md:52's leveling and raid Elemental points, the twelfth Elemental point and the Ancestral Knowledge trade all decode correctly.

## What checked out (no finding)

- **Round 10 Rockbiter prose** against 70755: Rockbiter r7 (16316) "increasing melee attack power by 554 and allowing melee attacks to cause additional threat"; Elemental Weapons 3/3 "Increases the melee attack power bonus of your Rockbiter Weapon by 20%, your Windfury Weapon effect by 40%"; Windfury Weapon r4 "When applied to main hand, disables any benefit you personally receive from Windfury Totem"; Windfury Totem "party or raid members within 30 yards". Raid preset (`sim/request/raid-consumes.json` shaman-enhancement) carries `windfury_totem` and `main_hand_imbue:rockbiter_weapon` from 32 on, Windfury Weapon only at 30-31. enhancement.md:57, :66, :79, :81, :101 and card .2 all agree with this and with each other. Rockbiter is a temporary weapon enchant, so "no room for a sharpening stone" still holds. Spirit Weapons (the Rockbiter threat talent) is not in either build, so the pages correctly say nothing about it.
- **Round 10 Elemental ranks**: `rank-choices.json` shaman-elemental 15207 rows (rank 3 at 30-39, 4 at 40-49, 6 at 50-59, 8 bare 60, 10 raid 60) match elemental.md:80 and cards .7/.8. Lightning Bolt r8/r9/r10 160/190/220 mana, 240/295/353 base: "about 16% less", "about 20% more" both exact. Rank 8 (learned 44) keeps its whole bonus at 60.
- **Round 10 Restoration card .17**: "Below 50 a cheaper rank" matches the 10396 rows (rank 3 at 20-39, rank 5 at 40-49). Healing Wave trainer levels 1/6/12/18/24/32/40/48/56/60, so "ranks 1 to 4 keep none of it at 60" and restoration.md:76's arithmetic (Lesser Healing Wave 3 at 36 keeps 65%, Healing Wave 6 at 32 keeps 45%, Chain Heal 1 at 40 keeps 85%) all hold.
- **Chain Heal ranks** (restoration.md:80): r1/r2/r3 260/315/405 mana, 251/335/490. With rank 1 at 85% of the bonus, rank 2 stays ahead per mana up to roughly 1,450 healing power, so "never better than rank 2" holds for any Forever raid set.
- **Trainer rows** for every spell the pages name (`trainables/shaman.json`): Rockbiter 1 (r4 24, r7 54), Earth Shock 4, Lightning Shield 8, Flame Shock, Searing, Strength of Earth, Flametongue 10, Fire Nova 12, Tremor 18, Ghost Wolf, Call of the Elements, Lesser Healing Wave, Healing Stream 20, Totemic Call, Poison Cleansing, Cure Disease 22, Mana Spring and Magma 26, Windfury Weapon 30, Chain Lightning and Windfury Totem 32, Disease Cleansing 38, Chain Heal 40/46/54, Grace of Air 42/56, Mana Tide 38/48/58, Riptide 40/50/60, Lightning Bolt 9/10 at 50/56, Earth Shock 7 at 60.
- **Talent text at max rank** re-read for every talent the pages describe; unchanged from round 10 and every quoted value holds.
- **Races**: Blood Fury "Attack Power, Ranged Attack Power and Spell Power by 10% for 15 seconds on a 2 minute cooldown"; Berserking 10% casting and attack speed, 15 s / 3 min; Stoneform 10% physical, Bleed, Poison and Disease, 8 s; Artisan's Precision 1%; Endurance 5%; Windshaper Skyborne named only as the paid option. No race line contradicts a measured result.
- **Consumables** against the raid preset: Elemental Flask of Supreme Power, Greater Arcane Elixir, Brilliant Wizard Oil, Major Mana Potion, Demonic Rune; Enhancement Mongoose, Juju Power, Juju Might, R.O.I.D.S., Smoked Desert Dumpling; Restoration Brilliant Mana Oil, Major Mana Potion, Demonic Rune.
- **Voice**: no em or en dashes, no site, sim, APL, preset or band talk, no nightly numbers, no stat ranking in prose.
- **Source links** (fetched 2026-11-04): Blizzard Deep Dive recap 200, Talents Forever data.json 200, Wowhead Skyborne First Look 200, foreversixty.gg /sim/weights, /planner and the three /bis/shaman/<spec> links 200; Icy Veins calculator 403 to automated fetches (unknown, not dead).

## Findings

| id | status | kind | file:line | text | evidence | fix |
|---|---|---|---|---|---|---|
| SHAMAN-R11-1 | CONFIRMED | card note | `data/curated/apl/shaman-restoration.json` priorityList.16 | "it heals three targets, each jump half the last, and the jumps reach any raid member near the target." | The 70755 client description of every Chain Heal rank (1064, 10622, 10623, `spelltips/shaman.json`): "If cast on a party member, the heal will only jump to other party members." Patch 1.12.0 (Warcraft Wiki) only made a raid-member target prefer raid targets over non-raid ones; it did not lift the party clause, which Engadget's 2009 Raid Rx still states as live behaviour. So when the target is in your party, the jumps cannot reach "any raid member". The card also contradicts restoration.md:72 ("Cast on someone in your party, it only jumps to your party"). Round 1 Q2 and round 2 Q2 read the 1.12.0 note as removing the party rule; per the round 2 lesson the client description wins. | "Chain Heal on the tank as soon as anyone near it is under 90%, while your mana is ahead of the fight: it heals three targets, each jump half the last. Cast on someone in your own party it only jumps inside your party; cast on anyone else it can jump to any raid member near them." |
| SHAMAN-R11-2 | LIKELY | rotation | restoration.md:72 | "Cast on someone in your party, it only jumps to your party, so ask to be grouped with the tank." | The rule is right (see R11-1), but the conclusion is backwards for Chain Heal: with the tank in your party, every Chain Heal you cast on the tank can only jump to your own four party members; with the tank in another group, the same cast can jump to anyone in the raid near it. Grouping with the tank is still right for Healing Stream and Mana Tide, which restoration.md:82 already says. LIKELY, not CONFIRMED: the raid-wide half rests on the 1.12.0 note plus the tooltip, not a Forever log. | "Riptide's 25% applies to the target you cast it on, and the two jumps go to the most injured people near it, each for half the heal before. Cast on someone in your own party, it only jumps inside your party, so with the tank grouped with you your Chain Heals on it stay in your group; cast on a raid member outside your party, it can jump to anyone near them." |

Wording SHOULDs, not findings:

| id | status | kind | file:line | text | note | suggested wording |
|---|---|---|---|---|---|---|
| SHAMAN-R11-S1 | SHOULD | wording | index.md:26 | "Windfury Weapon goes on that weapon." | Since round 10 Enhancement wears Rockbiter in a raid under a Windfury Totem. The line is not false (Windfury Weapon does go on the one weapon when you use it). | "Enhancement fights with a two-hander, or a one-hander with a shield or a held off-hand item, and your weapon imbue goes on that one weapon." |
| SHAMAN-R11-S2 | SHOULD | wording | enhancement.md:30 | "one weapon with Windfury on it" | Same as S1: in a raid the weapon carries Rockbiter and the Windfury swings come from a totem. | "one weapon carrying Windfury, from your own imbue or a raid's totem, Stormstrike on cooldown, ..." |
| R9-S1 to R9-S4 | SHOULD (carried) | wording | as round 9 | Elemental Flame Shock card "nothing else wants the shock cooldown"; Rockbiter "below 20" (enhancement.md:117); "which behind a boss is never" (enhancement.md:89); Flame Shock "ahead of everything but your cooldowns" (elemental.md:72). | Unchanged. | As round 9. |

## Decisions

- **D1. Chain Heal's party clause is live, and round 1 Q2 / round 2 Q2 / round 8 D2 over-read the 1.12.0 note.** The note says a raid-member target "will look for valid raid targets to jump to rather than non-raid targets as a priority": it ranks raid over non-raid, and says nothing about the party clause. The 70755 tooltip on all three ranks keeps "If cast on a party member, the heal will only jump to other party members", and Engadget (2009) still teaches it as current behaviour two expansions later. Client description beats an inferred patch reading, so the rule is: own-party target, party jumps; any other raid member, raid jumps. The rendered-review MUST that added it to restoration.md:72 was right on the fact; R11-2 is only about the conclusion drawn from it.
- **D2. Group placement advice stays.** restoration.md:82 (be grouped with the tank and melee for Healing Stream, Mana Spring and Mana Tide) is still the right call: those totems never leave your party, while Chain Heal on the tank still reaches your melee group. R11-2's fix states the trade without telling the player to leave the tank's group.
- **D3. Rockbiter under a Windfury Totem** (round 10 sim-lane handoff) is consistent with the 70755 tooltips and the preset; nothing to change. Spirit Weapons is not taken, so no threat-talent caveat is needed.
- **D4. Chain Lightning card "most of another bolt"** (Elemental .5): on two targets the jump's base (357) equals a rank 10 bolt (353), but its spell power share is 0.5 against the bolt's 0.857, so in raid gear the jump is most of a bolt. Kept.

## Sims requested

1. Carried (rounds 6 to 10): **shaman-elemental, Earth Shock filler**, band 60, raid preset, sim/cmd/spec-breakdown, 10,000 iterations, seeds 1 and 50001, race dwarf, `-talents 5510331300001351-055000000000000000-0530020000000000`: (a) published rotation; (b) plus an Earth Shock rank 7 (10414) line right after Flame Shock, gated on Flame Shock remaining above 6 s and mana percent above the share of the fight left. Report DPS and time out of mana.
2. Carried (round 10): **shaman-restoration, Nature's Swiftness save at 44 to 49**, bare preset, levels 45 and 49, heal profile, seeds 1 and 50001, races tauren and dwarf, leveling build: (a) as shipped (save casts rank 5 through the shared 10396 row) against (b) the save at the top rank learned (7 at 44 to 47, 8 at 48 and 49). If (b) wins or ties, scope the row to the filler line.
3. None new. Race lines on the round 15 builds remain unmeasured and none contradicts a measured result.

## Sim findings

- New: the engine's Chain Heal (`chain_heal.go`, sim audit round 10 R5) has no party-only jump rule. The heal profile puts the healer in raid slot 1 and the tank in slot 5 (another party), so the shipped numbers are what the game does for that layout; the rule only matters if a profile ever groups the healer with the tank.
- Carried: Stormstrike `RequiredLevel: 40` (only feeds the level penalty per the sim lane; may be closed); haste and the global cooldown (SIM-R4-1); Grace of Air and Strength of Earth book ranks; Lightning Overload ratio on the first level-60 log; Air totem raid context; Spellblasting potion on the Elemental pull; Totemic Call positioning (SIM-R5-1); self-dropped Windfury Totem plus Rockbiter not a shipped preset.
- Data lane (already filed): Restoration band 60 Horde rows run as tauren while `recommendedRaces` is [dwarf, orc]; Elemental raid off hand and Enhancement ranged slot.

## Sources

- Wowhead Forever / client, Chain Heal tooltips 1064, 10622, 10623 (text as in `spelltips/shaman.json`) https://nether.wowhead.com/forever/tooltip/spell/10623
- Warcraft Wiki, "Chain Heal" (Classic tooltip; patch 1.12.0 raid-target priority) https://warcraft.wiki.gg/wiki/Chain_Heal
- Engadget, "Raid Rx: Importance of group assigned healing" (2009; "If the Shaman casts it on a party member, it will only jump to other party members") https://www.engadget.com/2009-07-01-raid-rx-importance-of-group-assigned-healing.html
- Wowhead Forever, Rockbiter Weapon r7 and Windfury Weapon r4 tooltips https://nether.wowhead.com/forever/tooltip/spell/16316 , https://nether.wowhead.com/forever/tooltip/spell/16362
- Blizzard, "World of Warcraft Forever Deep Dive panel recap" https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- Wowhead, "Skyborne First Look" https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
- Talents Forever, data export (beta client data files) https://talentsforever.com/data.json
- Icy Veins, "Start Planning Your WoW Forever Character With Our New Talent Calculator" (403 to automated fetch; unknown, not dead) https://www.icy-veins.com/wow-forever/news/start-planning-your-wow-forever-character-with-our-new-talent-calculator/

## Fix log (round 11)

- SHAMAN-R11-1: fixed, shaman-restoration.json Chain Heal note (notes only), using the report's text. Party clause checked against spelltips/shaman.json ("If cast on a party member, the heal will only jump to other party members").
- SHAMAN-R11-2: fixed, restoration.md:72: "Cast on someone in your own party, it only jumps to other members of your party, so with the tank in your party your Chain Heals on it stay in your party." Worded from the client rule only; no claim about which grouping is better, and no raid-wide jump claim in the prose. restoration.md:82 (totem grouping) kept.
- SHAMAN-R11-S1: applied, index.md:26 "your weapon imbue goes on that one weapon".
- SHAMAN-R11-S2: left. The longer enhancement.md:30 summary breaks the 30-word summary limit in _sections.test.ts; the current line is not false.
