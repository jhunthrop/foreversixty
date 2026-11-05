# Guide accuracy round 20: Shaman

Auditor: shaman (class index, Elemental, Enhancement, Restoration). Audit tree 4234643f5 (main after guide round 19 and sim round 20), engine pin d5c708834 (engine source read-only at `gr20/engine`), ACTIVE client 1.60.1.70790 under `data/builds/1.60.1.70790/`.
Files: `web/src/content/guides/shaman/{index,elemental,enhancement,restoration}.md`, every note in `data/curated/apl/shaman-{elemental,enhancement,restoration}.json`, every rendered row of `addon-data.json` (`rotations/shaman-elemental` 26 bands, `shaman-enhancement` 19, `shaman-restoration` 23), the first sentences in `addon/ForeverSixty/Data.lua`, all six orders in `addon-builds.json` (leveling AND raid), `talents/shaman.json` (every rank text), `spelltips/shaman.json`, `trainables/shaman.json`. Read first: round 19 `shaman.md`, sim round 20 `shaman.md` and `integration.md`, `ledger.md` through round 19, the brief with its round 20 notes. Engine read: `sim/core/spell.go` `SpellLevelPenalty` / `LevelPenaltyFromMaxLevel`. Site: `web/src/lib/guides/rotation-lists.ts` (`clauseOf`), `web/src/lib/sim/rotations.ts` (the drawer prints every full note, per spec, not per band).

## What moved since round 19 (6beb919c2 to 4234643f5)

- **Elemental talent order (sim round 20 E1).** Both Elemental orders in `data/builds/1.60.1.70790/addon-builds.json` changed: Elemental Devastation 1 and Elemental Warding 3 are replaced by Trial by Fire 2 and Elemental Warding 2. Frontmatter `build` / `raidBuild` moved to `5500331220001351`. elemental.md:66 to :68 rewritten.
- **Enhancement prose (sim round 20 N3, N7).** enhancement.md:55 (Improved Ghost Wolf points to Ancestral Knowledge for a raid that never moves), :87 ("so Earth Shock never needs to wait").
- **Round 19's fixes.** Elemental Mana Spring note (R19-1), Elemental Fire Nova note (R19-2), Enhancement Earth Shock note (R19-3), restoration.md:120 "from 44" (R19-4). `addon-data.json` and Data.lua regenerated: the only changed shaman rows are the Elemental Fire Nova and Enhancement Earth Shock first sentences; every other row is byte-identical.
- No shaman page or rendered note names 70755 (the one 70755 is the unrendered top-level changelog in `shaman-elemental.json`).

## Builds and shipped orders (decoded, index 0 = level 10)

- **Elemental leveling (NEW):** Convection 10-14, Concussion 15-19, Elemental Alacrity 20-22, Call of Flame 23-25, Elemental Focus 26, Elemental Warding 27-28, **Trial by Fire 29-30**, Call of Thunder 31, Lightning Overload 32-34, Elemental Fury 35-39, Lightning Mastery 40, Thundering Strikes 41-45, Tidal Focus 46-50, Mindfulness 51-53, Totemic Focus 54-58, Improved Reincarnation 59-60.
- **Elemental raid (NEW):** identical to 45, then Ancestral Knowledge 46-50, Tidal Focus 51-55, Mindfulness 56-58, Improved Reincarnation 59-60.
- **Enhancement leveling:** Thundering Strikes 10-14, Ancestral Knowledge 15-19, Mental Dexterity 20-22, Elemental Weapons 23-24, Stormstrike 25, Elemental Weapons 26, Shamanistic Focus 27, Flurry 28-32, Improved Ghost Wolf 33-34, Mental Quickness 35-36, Maelstrom Weapon 37-39, Rage of the Farseer 40, Maelstrom Weapon 41-42, Convection 43-45, Concussion 46-47, Elemental Devastation 48-50, Call of Flame 51-53, Tidal Focus 54-58, Mindfulness 59-60.
- **Enhancement raid:** Thundering Strikes 10-14, Ancestral Knowledge 15-17, Mental Dexterity 18-20, Elemental Weapons 21-23, Shamanistic Focus 24, Stormstrike 25, Flurry 26-30, Improved Ghost Wolf 31-32, Mental Quickness 33-34, **Maelstrom Weapon 35-39**, Rage of the Farseer 40, Concussion 41-45, Elemental Devastation 46-48, Reverberation 49-51, Elemental Focus 52, Tidal Focus 53-57, Mindfulness 58-60.
- **Restoration (both identical):** Improved Healing Wave 10-14, Mindfulness 15-16, Tidal Focus 17-19, Water Shield 20, Tidal Focus 21-22, Mindfulness 23, Ancestral Healing 24, Mana Tide 25, Ancestral Healing 26-27, Healing Focus 28-30, Tidal Mastery 31-35, Restorative Totems 36-40, Healing Way 41-43, Nature's Swiftness 44, Purification 45-49, Riptide 50, Thundering Strikes 51-55, Ancestral Knowledge 56-60.

Client values used (`talents/shaman.json`, `spelltips/shaman.json`): Trial by Fire (16086) rank 1 "Fire Nova by 10% and reduces its cooldown by 2 sec, ... Magma Totem by 10 sec", rank 2 "20% ... 4 sec ... 20 sec"; Magma Totem 8190 / 10585 / 10586 / 10587 "for 20 sec ... within 8 yards"; Fire Nova 408341 to 408345 cooldown 10,000 ms, "within 10 yd of your active Fire totem"; Elemental Warding rank 2 "7%". Trainer levels as round 19 (Magma Totem and Mana Spring 26 / 36 / 46 / 56, Chain Lightning 32 / 40 / 48 / 56, Grace of Air 42 / 56, Riptide 50 / 60 trained ranks).

## Findings

Each card fix needs `sim/request/apl/shaman-*.apl.json`, `addon-data.json` and Data.lua regenerated (notes only).

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| SHAMAN-R20-1 | CONFIRMED | card note (talent value), new with sim round 20 | `shaman-elemental.json` priorityList[4] (Magma Totem) note, first sentence; rendered on the Elemental Magma Totem row at bands 30, 32, 34, 36, 38, 39, 40, 42, 44, 45, 46, 48, 50, 51, 52, 56, 57, 60 (18 of its 20 bands), in the web card and in Data.lua (20 rows) | "Magma Totem on two or more targets: it burns everything standing next to it for 20 seconds." | Sim round 20 put Trial by Fire in both Elemental orders at 29 and 30. Rank 2 "increases the duration of your Magma Totem by 20 sec" (16086), so from 30 every rank of Magma Totem (20 sec base, `spelltips` 8190 to 10587) burns for 40 seconds, 30 at 29. The engine applies it (sim round 20 E2, `fire_totems.go`). The page's own talents paragraph says so: elemental.md:66 "keeps Magma Totem burning for 40 seconds". True only at 26 and 28. Both orders are identical here, so no order rescues it. | Replace "for 20 seconds." with "for 20 seconds, or 40 from 30 with both points of Trial by Fire." Whole note: "Magma Totem on two or more targets: it burns everything standing next to it for 20 seconds, or 40 from 30 with both points of Trial by Fire." |
| SHAMAN-R20-2 | CONFIRMED | page prose (talent value, self-contradiction), new with sim round 20 | elemental.md:84 (Packs) | "drop Magma Totem in the pack (20 seconds), cast Chain Lightning on its 6 second cooldown and Fire Nova on its 10 second cooldown" | The rotation section is the level 60 page, and both shipped builds carry Trial by Fire 2/2 from 30: Fire Nova's 10 s cooldown (408345 `cooldown_ms` 10000) minus 4 s is 6 s, and Magma Totem lasts 20 + 20 = 40 s (16086 rank 2). The same page states both at :66 ("Fire Nova hit 20% harder on a 6 second cooldown and keeps Magma Totem burning for 40 seconds"), so :84 contradicts :66 twice. Sim round 20 rewrote :66 and left :84. (enhancement.md:85's "10 second cooldown" is right: neither Enhancement build takes Trial by Fire.) | Replace "drop Magma Totem in the pack (20 seconds), cast Chain Lightning on its 6 second cooldown and Fire Nova on its 10 second cooldown" with "drop Magma Totem in the pack (40 seconds with Trial by Fire), and cast Chain Lightning and Fire Nova each on its 6 second cooldown". |
| SHAMAN-R20-3 | LIKELY | card note (level 60 figure on lower bands) | `shaman-elemental.json` priorityList[7] (Lightning Bolt at the top rank) note, second sentence; the row renders at all 26 bands 10 to 60 (full note in the sim drawer) | "Rank 10 costs 220 mana to rank 9's 190 for about 20% more base damage, so you spend the spare mana on it and drop to a cheaper rank as soon as you fall behind." | Rank 10 (15208) is trained at 56; the row casts rank 2 (529) at 10 and 12, rank 3 at 14 and 18, rank 4 at 20 and 22, rank 5 at 26 to 30, rank 6 at 32 to 36, rank 7 at 38 to 42, rank 8 at 44 to 48, rank 9 at 50 to 52 and rank 10 only at 56, 57 and 60. On 23 of 26 bands "spend the spare mana on it" points at a rank the player does not have, and the figures are the level 60 pair: the round 20 rule ("quote no level-60 figure ... on a band where a lower rank renders"). LIKELY, not CONFIRMED, because the sentence names its ranks, so every number in it is true as written (220 / 353 against 190 / 295, +19.7%); the reasoning (the top rank costs a little more for more damage) holds at every band. | Replace the sentence with "At 60, rank 10 costs 220 mana to rank 9's 190 for about 20% more base damage; at any level you spend spare mana on your top rank and drop to a cheaper rank as soon as you fall behind." |

Counts: CONFIRMED 2, LIKELY 1, NEEDS-SIM 0, NEEDS-OWNER 0.

## Shoulds (wording, not findings)

- **elemental.md:66 "you never swing a weapon"** against elemental.md:114 ("finish in melee to save mana" at 1 to 19). The talents paragraph is about the builds, whose toll points land at 27 to 30, so not filed; suggested "and an Elemental shaman never fights in melee once these points arrive, so skip it" or simply "Elemental Devastation is melee crit, which Elemental never uses, so skip it."
- **Restoration Healing Wave note** (priorityList[17]) "a rank loses your healing bonus once you are 17 levels past it": `LevelPenaltyFromMaxLevel` keeps the whole bonus at learn + 17 and cuts 5% from learn + 18 (gap = level - (learn + 5) - 12). restoration.md:76 ("keeps ... until you are 17 levels past ... then loses") and :80 ("learned at 40, starts losing your bonus at 58") read it right. Suggested "more than 17 levels past it". The instruction (rank 8, never lower) is right either way.
- **enhancement.md:55** "the two Improved Ghost Wolf points go here instead": the bullet covers both builds and the leveling build already has 5/5 Ancestral Knowledge; suggested "in the raid build, the two Improved Ghost Wolf points go here instead". Legal: tiers 1 to 4 still hold 23 points, Maelstrom Weapon's tier still opens at 25.
- **index.md:28** "Lightning Bolt and Chain Lightning are both 2.5 second casts before any talents": true from Lightning Bolt rank 3 (548, 2,500 ms); rank 2 (529, the top rank at 8 to 13) is 2,000 ms. Suggested "from rank 3".
- **Elemental Flame Shock note** (carried, round 19): Call of Flame is now 23 to 25 on both orders, so "with all three points of Call of Flame, from 25, it ticks for 21 seconds". Still never false on a rendered band (22 has none, 26 has all three).
- **Elemental Chain Lightning note** (carried, round 19): "for the same cast time" holds outside Lightning Mastery only.
- Carried as filed: R18-S1 (restoration.md:80), R14-S2, R15-S2, R12-S1, R9-S1 to S4.

## Checked and right this round

- **Sim round 20's Elemental build and prose (new work).** Both orders legal on the 70790 tree: Elemental Focus (tier 3) at 26 with 16 points above it; Trial by Fire (tier 4) at 29 with 19; Call of Thunder (tier 5) at 31 with 21; Elemental Fury (tier 6) at 35 with 25; Lightning Mastery at 40 as the 31st point. elemental.md:52 (every point in Elemental to 40, then Thundering Strikes, Tidal Focus, Mindfulness, Totemic Focus, Improved Reincarnation; raid trades Totemic Focus for Ancestral Knowledge), :54, :56 to :64 (Convection 10%, Concussion, Alacrity 2.5 to 2.0, Call of Flame 21 s and 15%, 21/6 = 3.5 cooldowns, Elemental Focus 10%, Call of Thunder 5%, Lightning Overload 30% half damage no threat, Elemental Fury +100% bonus), :66 (four toll points: 21 in the listed talents + 4 = 25; Trial by Fire 2/2 20%, 10 - 4 = 6 s, 20 + 20 = 40 s; Warding 2/3 = 7%; Devastation "melee attacks"), :68 (Improved Reincarnation 2/2: 4% health, 20 min off the hour, more health and mana). Levels 1 to 19 (:114: Convection 10 to 14 at 2% a point on shocks and Lightning Bolt, Concussion 15 to 19 at 1%) and :118 (Chain Lightning 32, Lightning Mastery 40, the bolt-rank list against the rendered filler rows) unchanged and still match the new order: no sentence names a level between 20 and 39.
- **Sim round 20's Enhancement prose (new work).** enhancement.md:55: moving Improved Ghost Wolf 2 into Ancestral Knowledge 3/5 keeps the raid tree legal (see Shoulds) and the sim measured it as a small gain (N3), so "a little more damage" holds. :87: "Earth Shock never needs to wait" matches N7 (the shipped gate and no gate are identical to the tenth at raid 60, the bar never sits under a quarter with a shock ready); the half-mana sentence after it is talent-relative and holds on both orders (Maelstrom Weapon 37 to 42 leveling, 35 to 39 raid).
- **Round 19's fixes.** R19-1 Mana Spring "At every rank it returns many times its own cost ... over its 5 minutes": rank 1 4 mana / 2 s for 40 (600 over 300 s), rank 4 10 for 100. R19-2 Fire Nova "(Searing Totem until 26, Magma Totem from then on)": Fire Nova 12, Magma Totem 26; the Searing fallback on two targets is the carried sim finding, not re-filed. R19-3 Earth Shock "once you have Stormstrike": Stormstrike 25 on both Enhancement orders, the row renders from 10. R19-4 restoration.md:120 "from 44": Nature's Swiftness at 44 on both orders.
- **Every other note on every band**: see the Note sweep appendix.
- **Prose, all four pages, line by line** (unchanged lines re-read against the new orders and 70790 files): index.md:22 to :36; elemental.md :28 to :118 (R20-2 aside); enhancement.md :28 to :121 (:52 raid 12 Elemental points against 11, Mindfulness 3 against 2, Ancestral Knowledge 3 against 5; Elemental Devastation 3/3 is 9%; :59 Stormstrike 6 s; :62 Rage of the Farseer 3 min; :85 Fire Nova 10 s with no Trial by Fire; :117 and :121 the leveling order exactly: Mental Dexterity 20 to 22, Elemental Weapons 23 and 24, Stormstrike 25 as the 16th point, Elemental Weapons 26, Shamanistic Focus 27, Flurry by 32, Improved Ghost Wolf 33 and 34); restoration.md :26 to :120 (41 Restoration points; :57 Mana Tide 5 min; :61 Riptide 12 s cooldown, 15 s; :76 and :80 against `LevelPenaltyFromMaxLevel`: Healing Wave 7 and Chain Heal 1 85%, Lesser Healing Wave 3 65%, Healing Wave 6 45%, learned at 23 or earlier 0%, Chain Heal 1 full to 57; :116 and :120 against the order).
- **Races, BiS, items, professions**: `races.json`, `bis/shaman-*.json` and the shaman `items` unchanged since round 19's verified state (no shaman diff under `data/builds/1.60.1.70790/` but `addon-builds.json`, `addon-data.json` and `dps-per-point.json`).
- **Voice**: no em or en dashes on any page or in any note; no nightly numbers; no site or sim talk.
- **Source links**: frontmatter sources unchanged since round 18's fetch.

## Decisions

- **D1. SHAMAN-R20-1 and R20-2 are CONFIRMED on client rank text** (16086) and the page's own :66; no sim is needed. The level-60 rotation prose (D1 of round 19) is held to the shipped build at 60, which now includes Trial by Fire.
- **D2. SHAMAN-R20-3 is LIKELY.** The round 20 rule names it (a level 60 figure on bands where a lower rank renders), but the sentence labels both ranks, so no number in it is false; the fix only scopes it.
- **D3. Not filed, per the brief and earlier rounds**: rows rendering before their talent (Enhancement Lightning Bolt and Fire Nova before Maelstrom Weapon and Elemental Focus; Restoration Mana Tide, Nature's Swiftness and its Healing Wave, Riptide 40 to 49); the autocast notes (Elemental Lightning Mastery, Restoration potions and Nature's Swiftness), which have no band row and describe the cooldown they cast; the Elemental card's missing Fire totem below 26 on two targets (round 19 Sim findings); BiS race captions; the leveling tree's rendering.

## Engine questions

- None new. (Sim round 20 E2 says `fire_totems.go` adds Trial by Fire's 20 s to Magma Totem and `fire_nova.go` takes its 4 s; R20-1 and R20-2 bring the text to the engine, not the other way round.)

## Sims requested

- None for this round's findings (client rank text and the page's own lines).
- Carried, already routed: Enhancement half reserve against a quarter at 37 to 41; Elemental band 20 / 30 talent strings at the refresh; Enhancement band 30 with Stormstrike, the relic slot; Restoration bands 40, 50 and bare 60 with the Chain Heal rows.

## Sim findings

- None new. Carried: the Elemental card drops no Fire totem below 26 on two or more targets; Stormstrike `RequiredLevel: 40` against the order's point at 25; Maelstrom Weapon proc rate (sim round 20 N2, open); Lightning Overload rank spells, Healing Wave rank 10 and the book ranks (pending owner).

## Sources

- Client data 1.60.1.70790 (`data/builds/1.60.1.70790/`): `talents/shaman.json` (Trial by Fire 16086 both ranks, Elemental Warding 28996, Elemental Devastation 30160, Call of Flame 16038, Elemental Alacrity 16578, Maelstrom Weapon 408498, Improved Ghost Wolf 16262, Mindfulness 1223033, every other rank text), `spelltips/shaman.json` (Magma Totem 8190 / 10585 / 10586 / 10587; Fire Nova 408341 to 408345; Flame Shock 8050 to 29228; Searing Totem 3599, 6363, 10438; Lightning Bolt 529 to 15208; Chain Lightning 421, 930, 2860, 10605; Earth Shock 8044, 10414; Stormstrike 17364; Mana Spring 5675, 10497; Mana Tide 16190, 17354, 17359, 1325256; Healing Stream 5394 to 10463; Healing Wave 332 to 25357; Lesser Healing Wave 8004, 10468; Chain Heal 1064, 10622, 10623; Riptide 408521, 1239242, 1239243), `trainables/shaman.json`, `addon-builds.json` (six orders), `addon-data.json` (`rotations/shaman-*`, every band), `races.json`, `bis/shaman-*.json`. `addon/ForeverSixty/Data.lua`.
- Curated: `data/curated/apl/shaman-{elemental,enhancement,restoration}.json`.
- Engine d5c708834: `sim/core/spell.go` (`SpellLevelPenalty`, `LevelPenaltyFromMaxLevel`, grace 12 past MaxLevel = learn + 5).
- Site code: `web/src/lib/guides/rotation-lists.ts`, `web/src/lib/sim/rotations.ts`.
- Reviews: `design/reviews/2026-11-04-sim-audit-round20/shaman.md` (E1, E2, N3, N7), `integration.md`; `guide-accuracy/round19/shaman.md`; `ledger.md` rounds 18 and 19.

## Note sweep (every row note, every band it renders)

Bands: Elemental 10 12 14 18 20 22 26 28 30 32 34 36 38 39 40 42 44 45 46 48 50 51 52 56 57 60; Enhancement 10 12 14 20 22 24 26 30 32 36 38 40 42 44 48 50 52 56 60; Restoration 10 12 18 20 24 26 28 30 32 36 37 38 40 42 44 47 48 50 52 56 58 59 60. Levels are the shipped order (L = leveling, R = raid) or the trainer.

**Elemental**
- [0] Mana Spring Totem, bands 26 to 60: Mana Spring trained 26 (rank 2 36, 3 46, 4 56); "at every rank ... many times its own cost ... 5 minutes" true at ranks 1 to 4; "two Mana Springs do not stack" (client: one Water totem per shaman, group aura). OK.
- [1] Lightning Mastery autocast, no band row (drawer): Lightning Mastery 40 L and R; 15 s, Lightning Bolt -0.5 s, Chain Lightning no cooldown (16166); Berserking after it, 1.5 s floor. OK (D3).
- [2] Flame Shock, bands 10 to 60: base 12 s every rank; "with Call of Flame ... 21 seconds" needs 3/3, which is 23 to 25 L and R: bands 10 to 22 have none, 26 up have all three. OK (Should).
- [3] Searing Totem, bands 10 to 60: trained 10; "strongest Fire totem" single target; "only reach 20 yards" every rank. OK.
- [4] Magma Totem, bands 26 to 60: trained 26; "20 seconds" true at 26 and 28, false at 30 to 60 (Trial by Fire 2/2 at 30 L and R, 40 s). **SHAMAN-R20-1.**
- [5] Chain Lightning, bands 32 to 60: trained 32; first hit 215 / 348 / 428 / 510 against the top bolt 148 / 194 / 240 / 295 / 353 at every band; 30% less a jump; 2.5 s base for both with Elemental Alacrity 3/3 from 22. OK (Should, Lightning Mastery).
- [6] Fire Nova, bands 12 to 60: trained 12; "Searing Totem until 26, Magma Totem from then on" (Magma 26); "within 10 yards of that totem" every rank; "Searing Totem stands 20 yards out". OK.
- [7] Lightning Bolt top rank, bands 10 to 60: first sentence OK at every band; second sentence's rank 10 / rank 9 pair renders at 10 to 52 where rank 2 to 9 is the top rank. **SHAMAN-R20-3.** "In a raid you never need to" OK.
- [8] Lightning Bolt filler, bands 10 to 60: rendered ranks rank 2 at 10 and 12, 3 at 14 and 18, 4 at 20 and 22, 3 at 26, 5 at 28, 3 at 30 and 32, 4 at 34 to 38, 5 at 39 to 44, 6 at 45 to 50, 7 at 51 to 56, 8 at 57, rank 10 at raid 60; the note's list (3 at 26 and 27, top at 28 and 29, 3 from 30 to 33, 4 until 38, 5 until 44, 6 until 50, 7 until 56, 8 from 57), "rank 8 at 60" without raid buffs, "rank 10 here too" in a raid. OK.

**Enhancement**
- [0] autocast, no note.
- [1] Strength of Earth, bands 10 to 60: trained 10; "whole raid within 30 yards" every rank. OK.
- [2] Grace of Air, bands 42 to 60: trained 42, rank 2 56 ("Learned at 42, rank 2 at 56"); Windfury Totem 32 and Rockbiter 1, both in hand from 42. OK.
- [3] Stormstrike, bands 40 to 60: talent 25 L and R; 150% weapon damage, 20% on the next Lightning Bolt, Chain Lightning or Earth Shock for 12 s, 6 s cooldown (17364). OK.
- [4] Lightning Bolt at five Maelstrom stacks, bands 10 to 60: "with all five points it is instant and free" conditional (5/5 at 42 L, 39 R; 4% a point a stack); "right after Stormstrike" can only fire from 37 L / 35 R, after Stormstrike at 25. OK.
- [5] Searing Totem, bands 10 to 60: trained 10; 20 s guard. OK.
- [6] Fire Nova on Clearcasting, bands 12 to 60: Fire Nova 12; "Searing Totem ... within its 10 yards" in melee; "Clearcasting comes from Elemental Focus in the raid build" (52 R, none L). OK.
- [7] Earth Shock, bands 10 to 60: "a quarter, or half from your first point of Maelstrom Weapon until the fifth" (37 to 41 L, 35 to 38 R); "once you have Stormstrike" (25 L and R); "Keep that reserve for your totems" (Searing and Strength of Earth from 10); "Below five points your five-stack bolts still cost mana" (4/5 = 80% off). OK on both orders.

**Restoration**
- [0] autocast (potions, runes, Nature's Swiftness), no band row (drawer). OK (D3).
- [1] Mana Tide Totem, bands 10 to 60: talent 25 (rows from 10, D3); 12 s (`${13 sec-1}`), group within 30 yards; Healing Stream from 20; "Rank 4 at 58" (1325256; ranks 2 and 3 at 38 and 48 match the rows). OK.
- [2] Water Shield, bands 20 to 60: talent 20; three globes, 2% of maximum mana, hit or a spell crit. OK.
- [3] Healing Stream, bands 20 to 60: trained 20; every 2 s for 5 min at every rank (5394 to 10463). OK.
- [4] Nature's Swiftness, bands 10 to 60: talent 44 (rows from 10, D3); "next Healing Wave becomes instant" (next Nature spell under 10 s). OK.
- [5] Healing Wave after Nature's Swiftness, bands 10 to 60: highest rank rendered at each band (rank 2 at 10, 3 at 12, 4 at 18 and 20, 5 at 24 to 30, 6 at 32 to 38, 7 at 40 to 47, 8 at 48 to 52, 9 at 56 to 60). OK.
- [6] Riptide, bands 40 to 60: talent 50 (rows from 40, D3); "cheapest heal per mana" (rank 1 445 / 150, rank 2 575 / 180, rank 3 805 / 230 against any Healing Wave, Lesser Healing Wave or Chain Heal); "Chain Heal ... 25% more". OK.
- [7] Lesser Healing Wave on the tank, bands 20 to 60: trained 20; 1.5 s every rank. OK.
- [8] to [15] party Lesser Healing Wave and Riptide rows, bands 20 to 60: no note.
- [16] Chain Heal, bands 40 to 60: trained 40, ranks 2 and 3 at 46 and 54; three targets, each jump 50%; party-only jumps (1064 tooltip); "rank 1 until 59" (rows 1064 at 40 to 59); 260 / 315 / 405 mana; "At 60 without raid buffs, cast rank 2". OK.
- [17] Healing Wave on the tank, bands 10 to 60: "at 60 drop to rank 8, never lower" (rank 8 learned 48 full; rank 7 85%); "ranks 1 to 4 keep none of it at 60" (rank 4 learned 18, zero from 55); leveling list from 20 matches every rendered filler row (rank 2 at 20 and 24, 3 at 26 to 36, 4 at 37 to 40, 5 at 42 and 44, 6 at 47 to 52, 7 at 56 and 58, 8 at 59). OK (Should: "17 levels past" is one level early).

## Counts

CONFIRMED 2 (SHAMAN-R20-1 Elemental Magma Totem note "20 seconds" on bands 30 to 60 after Trial by Fire; SHAMAN-R20-2 elemental.md:84 Magma Totem 20 s and Fire Nova 10 s against :66 and the client), LIKELY 1 (SHAMAN-R20-3 Elemental top-rank bolt note's rank 10 / 9 pair on bands 10 to 52), NEEDS-SIM 0, NEEDS-OWNER 0. Round 19's four fixes: 0 errors. Sim round 20's build and prose: 2 errors (both the Trial by Fire values it did not carry beyond :66). Shoulds 4 new, 2 updated, rest carried.
