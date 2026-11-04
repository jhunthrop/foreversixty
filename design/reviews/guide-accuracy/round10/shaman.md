# Guide accuracy round 10: Shaman

Auditor: shaman (class index, Elemental, Enhancement, Restoration). Audit tree origin/main 84460cec8, client 1.60.1.70755, engine pin 9f7f1f6e4 (`sim/enginever/version.go`; engine source read from the gr10/engine checkout at 9f7f1f6e4).
Files: `web/src/content/guides/shaman/{index,elemental,enhancement,restoration}.md`, card notes in `data/curated/apl/shaman-{elemental,enhancement,restoration}.json`, rank rows in `data/curated/rank-choices.json`. Read first: round1 to round9 `shaman.md` with every fix log, `round4/shaman-levels.md`, round2 to round5 `race-sims.md`, `ledger.md`.

## What moved since round 9 (645caa983 to 84460cec8)

Sim round 15 landed under the guides. For shaman it changed four things the pages depend on:

- `enhancement.md:6` leveling build, Elemental part `320033` to `323003` (485f1fe9a: Elemental Alacrity 3 traded for Elemental Devastation 3, "frontmatter build line only"). The prose was not touched.
- `addon-builds.json` Enhancement leveling and raid orders (round 15 regen d047d5359): Stormstrike moved from 32 to 25, and Elemental Alacrity 48-50 became Elemental Devastation 48-50. Elemental and Restoration orders are byte-identical to round 9.
- `shaman-restoration.json`: the Nature's Swiftness save casts Healing Wave rank 9 (cae4e3bb4), the tank Riptide line moved up to right after the save (d289877f9), and the tank filler Healing Wave names rank 9 (4faba89a9).
- `rank-choices.json`: Restoration Healing Wave (authored 10396) rows now cast rank 3 at 20-39 and rank 5 at 40-49 (before: the top rank learned at every band). Elemental Lightning Bolt rows now rank 3 at 30-39, 4 at 40-49, 6 at 50-59, 8 at bare 60.

Nothing changed in `talents/`, `trainables/`, `spelltips/`, `races.json`, `stat-weights.json` or `bis/` (the 13:00 UTC re-rank had not landed in this tree).

## Builds and shipped leveling order

Decoded `addon-builds.json` (entries [tree, tier, column], 1-based; index 0 = level 10) against `talents/shaman.json` (tier and column 0-based):

- Elemental (unchanged): Convection 10-14, Concussion 15-19, Elemental Devastation 20, Elemental Alacrity 21-23, Call of Flame 24-26, Elemental Focus 27, Elemental Warding 28-30, Call of Thunder 31, Lightning Overload 32-34, Elemental Fury 35-39, Lightning Mastery 40, Thundering Strikes 41-45, Tidal Focus 46-50, Mindfulness 51-53, Totemic Focus 54-58, Improved Reincarnation 59-60.
- Enhancement (new): Thundering Strikes 10-14, Ancestral Knowledge 15-19, Mental Dexterity 20-22, Elemental Weapons 23-24, **Stormstrike 25**, Elemental Weapons 26, Shamanistic Focus 27, Flurry 28-32, Improved Ghost Wolf 33-34, Mental Quickness 35-36, Maelstrom Weapon 37-39, Rage of the Farseer 40, Maelstrom Weapon 41-42, Convection 43-45, Concussion 46-47, **Elemental Devastation 48-50**, Call of Flame 51-53, Tidal Focus 54-58, Mindfulness 59-60. Stormstrike at 25 has 15 points above it (5+5+3+2); legal.
- Restoration (unchanged): Improved Healing Wave 10-14, Mindfulness 15-16, Tidal Focus 17-19, Water Shield 20, Tidal Focus 21-22, Mindfulness 23, Ancestral Healing 24, Mana Tide Totem 25, Ancestral Healing 26-27, Healing Focus 28-30, Tidal Mastery 31-35, Restorative Totems 36-40, Healing Way 41-43, Nature's Swiftness 44, Purification 45-49, Riptide 50, Thundering Strikes 51-55, Ancestral Knowledge 56-60.

Frontmatter strings: Enhancement leveling `323003/055030310051202051/052` = Convection 3, Concussion 2, Elemental Devastation 3, Call of Flame 3; Thundering Strikes 5, Ancestral Knowledge 5, Mental Dexterity 3, Elemental Weapons 3, Shamanistic Focus 1, Flurry 5, Stormstrike 1, Improved Ghost Wolf 2, Mental Quickness 2, Maelstrom Weapon 5, Rage of the Farseer 1; Tidal Focus 5, Mindfulness 2 (51). Raid `0533001/053030310051202051/053` = Concussion 5, Elemental Devastation 3, Reverberation 3, Elemental Focus 1; Enhancement as leveling but Ancestral Knowledge 3; Tidal Focus 5, Mindfulness 3 (51). Elemental and Restoration strings unchanged from round 9.

Stormstrike at 25 is round 1 Q1's researched answer (castable once talented, sixteenth point at 25); the shipped order now agrees with it.

## What checked out (no finding)

- Elemental and the class index: unchanged text, re-read line by line against the same 70755 data as round 9; every round 9 check stands (talent text at max rank, spell numbers, crit wording, races, stat reasons, bind facts, voice).
- Restoration card changes: card 5 note ("Your highest rank of Healing Wave right after Nature's Swiftness") and prose :33 and :74 agree with each other; the action casting rank 9 while rank 10 (25357, level 60) exists is a card action rank (known and accepted; engine `IncludeAQ = false` treats rank 10 as the Ahn'Qiraj book rank). Card 6 (Riptide on the tank under 90% when not ticking) now sits ahead of the tank's waves and matches prose item 1 and "your cheapest heal per mana" (230 mana; round 8 D1 and round 9 re-derivation stand).
- Round 9 D1 is closed by the engine: `sim/core/spell.go:822-831` at 9f7f1f6e4 implements the level-gap rule (17 levels of grace, then 5% a level, none from 37 levels past), the same rule restoration.md:76 and card 17 state. The page and the engine now agree.
- Enhancement Levels 1 to 19 (Thundering Strikes 10-14, Ancestral Knowledge 15-19) still match the order.
- Source links (fetched 2026-11-04): Blizzard Deep Dive recap 200, Talents Forever data.json 200, Wowhead Skyborne First Look 200, foreversixty.gg /sim/weights, /planner, /bis/shaman/{elemental,enhancement,restoration} 200; Icy Veins calculator 403 to automated fetches (unknown, not dead).

## Findings

| id | status | kind | file:line | text | evidence | fix |
|---|---|---|---|---|---|---|
| SHAMAN-R10-1 | CONFIRMED | talent | enhancement.md:52 | "The leveling build puts its Elemental points into Convection, Concussion, Elemental Alacrity and Call of Flame for cheaper, faster, harder Lightning Bolts and a stronger Searing Totem. The raid build moves them to Concussion, Reverberation, Elemental Devastation and Elemental Focus: Earth Shock hits harder and comes back sooner, its crits give your swings 9% crit for 10 seconds, and now and then a spell is free." | Leveling build `323003` (enhancement.md:6, commit 485f1fe9a) has Elemental Alacrity 0 and Elemental Devastation 3; `addon-builds.json` takes Elemental Devastation 48-50. So the leveling build has no "faster" bolts, and Elemental Devastation is not something the raid build "moves" to: both builds take it. | "The leveling build puts its Elemental points into Convection, Concussion, Elemental Devastation and Call of Flame: cheaper shocks and bolts, harder Lightning Bolts and Earth Shocks, 9% melee crit for 10 seconds after a spell crits, and a stronger Searing Totem. The raid build keeps Elemental Devastation, fills Concussion and moves the rest to Reverberation and Elemental Focus: Earth Shock hits harder and comes back sooner, and now and then a spell is free." |
| SHAMAN-R10-2 | CONFIRMED | leveling | enhancement.md:119 | "from 20 this build fills Mental Dexterity, Elemental Weapons, Shamanistic Focus and Flurry, and takes Stormstrike at 32, with Improved Ghost Wolf right after." | Shipped order (`addon-builds.json` shaman-enhancement leveling, index 15 = level 25 is [2,4,3] Stormstrike): Mental Dexterity 20-22, Elemental Weapons 23-24, Stormstrike 25, Elemental Weapons 26, Shamanistic Focus 27, Flurry 28-32, Improved Ghost Wolf 33-34. Stormstrike is at 25, before Shamanistic Focus and Flurry, and Improved Ghost Wolf follows Flurry, not Stormstrike. | "Stormstrike sits in the fourth row and you can press it the moment you learn it: from 20 this build fills Mental Dexterity, puts two points in Elemental Weapons and takes Stormstrike at 25, the first level the fourth row opens. Then it finishes Elemental Weapons, takes Shamanistic Focus and fills Flurry by 32, with Improved Ghost Wolf at 33 and 34." |
| SHAMAN-R10-3 | CONFIRMED | card note (rank) | `data/curated/apl/shaman-restoration.json` priorityList.17 | "While leveling, cast the highest rank you have." | `rank-choices.json` shaman-restoration, Healing Wave authored 10396: rank 3 at 20-39 and rank 5 at 40-49 (commit 4faba89a9 and `sim/request/healer_ranks_test.go:160-176`: +10.9 to +19.8% at 20 and 30, +8.4 to +8.6% at 40, each fight running dry). The same card line casts rank 3 or 5 at those bands while its note says to cast the top rank. Before round 15 the rows cast the top rank learned at every band, so the note was true; it is now false against its own card. | "Healing Wave when the tank drops below 60%. To save mana at 60 drop to rank 8, never lower: a rank loses your healing bonus once you are 17 levels past it, and ranks 1 to 4 keep none of it at 60. While leveling your gear adds little healing, so a cheaper low rank heals more before your mana runs out than your top rank does." |
| SHAMAN-R10-4 | LIKELY | leveling | elemental.md:116 | "Drop a rank or two on a long fight where your mana, not the mob, runs out first." | `rank-choices.json` shaman-elemental Lightning Bolt on the mana-bound leveling fight: rank 3 at 30-39 (top is 5 to 7, so 2 to 4 ranks down), rank 4 at 40-49 (top 7 or 8: 3 to 4 down), rank 6 at 50-59 (top 9 or 10: 3 to 4 down), rank 8 at bare 60 (2 down). The measured drop is mostly three or four ranks. LIKELY, not CONFIRMED: a quest pull is shorter than the measured fight, and "a rank or two" is still better than the top rank there. | "Drop well down the ranks on a long fight where your mana, not the mob, runs out first: a bolt three or four ranks below your best costs far less and does more in total before you run dry." |

Wording SHOULDs, carried and not findings: R9-S1 (Elemental Flame Shock card "nothing else wants the shock cooldown"), R9-S2 (Rockbiter "below 20" at enhancement.md:115), R9-S3 ("which behind a boss is never", enhancement.md:87), R9-S4 (Flame Shock "ahead of everything but your cooldowns", elemental.md:72, against the Mana Spring line).

## Decisions

- **D1. Stormstrike at 25 is right.** Round 1 Q1 (Method, Warcraft Tavern, Mobalytics Forever leveling guides: Stormstrike at 25, talent-taught) stands, and the shipped order now takes it at 25. The fix to SHAMAN-R10-2 keeps "you can press it the moment you learn it". The engine still gates it at 40 (sim finding below), which is the sim's error, not the page's.
- **D2. Healing Wave rank 10 in "highest rank".** restoration.md:33, :74 and card 5 say "highest rank" while the action casts rank 9. The engine treats 25357 as the Ahn'Qiraj book rank (`IncludeAQ = false`). "Highest rank" is true for whatever a player owns, so no prose change. The action rank is a known card action item.
- **D3. Elemental Devastation on the leveling build.** Commit 485f1fe9a measured it +1.0 to +1.2% over Elemental Alacrity (bare 50 and 60, both factions, two seeds), and Alacrity does nothing for a rotation that casts Lightning Bolt only when it is instant. The fix in SHAMAN-R10-1 states the effect and leaves out any number.

## Sims requested

1. Carried from rounds 6 to 9: **shaman-elemental, Earth Shock filler**, band 60, raid preset, sim/cmd/spec-breakdown, 10,000 iterations, seeds 1 and 50001, race dwarf, `-talents 5510331300001351-055000000000000000-0530020000000000`: (a) the published rotation; (b) the same with an Earth Shock rank 7 (10414) line directly after the Flame Shock line, gated on Flame Shock's remaining duration above 6 s and on mana percent above the share of the fight left. Report DPS and time out of mana.
2. New: **shaman-restoration, Nature's Swiftness save at 44 to 49**, bare preset, levels 45 and 49, heal profile bands, seeds 1 and 50001, race tauren and dwarf, the leveling build at those levels (Nature's Swiftness is taken at 44). The 40-49 rank 5 row on authored id 10396 also rewrites the Nature's Swiftness Healing Wave line (test comment `healer_ranks_test.go:165-166` assumes band builds below 50 have no Nature's Swiftness, but the shipped order takes it at 44). Compare (a) as shipped (save casts rank 5) against (b) the save at the top rank learned (7 at 44 to 47, 8 at 48 and 49). The page says the save is the biggest rank; if (b) wins or ties, the row should be scoped to the filler line only.
3. Provisional, waits for the 13:00 UTC best-in-slot re-rank: none of the shaman BiS prose names a row by value (bind facts only), so nothing to check after it lands beyond the stat-table order the prose no longer ranks.

## Sim findings

- New: the Restoration Healing Wave rank row (40-49 rank 5) applies to the Nature's Swiftness save as well as the filler, since both lines author 10396; a 44 to 49 shaman with Nature's Swiftness sees the save cast at rank 5 against the card note and restoration.md:74 (Sims requested 2).
- New: the Enhancement order now takes Stormstrike at 25, but `sim/shaman/stormstrike.go:41` at 9f7f1f6e4 still sets `RequiredLevel: 40`, so bands 25 to 39 measure the point as dead. Round 1 Q1's change (gate on the talent only) is still open; once it lands, the 25 to 39 bands and Enhancement weights there move.
- Closed: the healing low-rank curve (round 1 Q3 (a), round 9 D1); the engine now carries the level-gap rule the page states.
- Carried: haste shortening the global cooldown (SIM-R4-1); Grace of Air and Strength of Earth book ranks (SIM-R2-2); Lightning Overload ratio on the first level-60 log; Air totem raid context; Spellblasting potion on the Elemental pull; Totemic Call positioning (SIM-R5-1).
- Data lane (already filed): Restoration band 60 Horde rows run as tauren while `recommendedRaces` is [dwarf, orc]; Elemental raid off hand and Enhancement ranged slot.

## Sources

- Method, "WoW Forever Shaman Leveling Guide and Talents" (Stormstrike at 25) https://www.method.gg/wow-forever/wow-forever-shaman-leveling-guide-and-talents
- Warcraft Tavern, "Shaman Leveling Guide" (Forever; Stormstrike at 25) https://www.warcrafttavern.com/forever/guides/shaman-leveling/
- Blizzard, "WoW Forever Beta Development Notes" (Kaivax; 1 Oct low-rank note) https://us.forums.blizzard.com/en/wow/t/forever-beta-development-notes/2360696
- ForeverChanges, "WoW Forever downranking: the level penalty and which ranks to keep" https://foreverchanges.pro/downranking
- Blizzard, "World of Warcraft Forever Deep Dive panel recap" https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- Wowhead, "Skyborne First Look" https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
- Talents Forever, data export (beta client data files) https://talentsforever.com/data.json
- Icy Veins, "Start Planning Your WoW Forever Character With Our New Talent Calculator" (403 to automated fetch; unknown, not dead) https://www.icy-veins.com/wow-forever/news/start-planning-your-wow-forever-character-with-our-new-talent-calculator/

## Fix log (round 10)

- SHAMAN-R10-1: fixed, enhancement.md:52 (leveling build names Elemental Devastation; raid build keeps it, fills Concussion, moves the rest to Reverberation and Elemental Focus).
- SHAMAN-R10-2: fixed, enhancement.md:121 (Mental Dexterity, two points of Elemental Weapons, Stormstrike at 25, then Elemental Weapons, Shamanistic Focus, Flurry by 32, Improved Ghost Wolf 33 and 34). "Stays on for good" for Windfury Weapon became "stays on while you level" (see Rockbiter below).
- SHAMAN-R10-3: fixed, shaman-restoration.json priorityList.17 note: the leveling sentence now says a cheaper rank heals more below 50 (matches the 20-39 rank 3 and 40-49 rank 5 rows; top rank at 50 and 60).
- SHAMAN-R10-4: fixed (verified against rank-choices.json), elemental.md:116, the reviewer's wording.
- D1 applied (Stormstrike at 25, "press it the moment you learn it" kept). D2: no prose change, as decided. D3: effect stated, no number (enhancement.md:52).
- Sim-lane handoff, Elemental raid rank 10 (rank-choices raid-60 row 15207 to rank 10; bare-60 row rank 8): elemental.md:80 now says rank 10 the whole raid fight and, without raid buffs, pace and drop to a cheaper rank, rank 8 at 60. Card notes priorityList.7 ("In a raid you never need to") and .8 (rank 10 in a raid, rank 8 at 60 and lower while leveling without raid buffs) rewritten.
- Sim-lane handoff, Rockbiter under a Windfury Totem (shipped in sim/request/preset.go swapImbueUnderWindfuryTotem, from level 32 wherever the preset carries windfury_totem; client: Windfury Weapon "disables any benefit you personally receive from Windfury Totem", Rockbiter rank 7 554 melee AP plus extra threat, Elemental Weapons 3 +20% Rockbiter AP): enhancement.md:57 (Elemental Weapons Rockbiter part), :66 (before the pull), :79 new "Which imbue" paragraph (with Rockbiter's threat line from its tooltip), :81 (self-dropped Windfury Totem keeps Rockbiter on), :101 (consumables), :121; shaman-enhancement.json priorityList.2 note. No numbers quoted.
- Stat priority: all three spec pages already explain without ranking and point to the table; prose test clean for shaman. No other place ranks stats or gear by nightly value.
- Empty notes (Enhancement .0, Restoration .8-.15): not new lines (empty at 645caa983 as well; member repeats of carded lines), left.
- Carried SHOULDs R9-S1 to R9-S4: left, wording only and not findings.
- updated: already 2026-11-04 on elemental, enhancement, restoration and index; index.md not changed this round.
- SIM findings: none new beyond the two in this report (Nature's Swiftness save at 44-49 casts rank 5 through the shared 10396 row; Stormstrike RequiredLevel 40, which the sim lane reports only feeds the spell-level penalty and never gates the cast, so that one may be closed). New: the raid preset's Rockbiter swap is measured only with another shaman's totem; the self-dropped Windfury Totem plus Rockbiter case (+0.9% in the sim lane's probe) is not a shipped preset, so measure it at refresh.
