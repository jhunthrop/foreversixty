# Sim accuracy 2026-11-12: Elemental and Enhancement against the client and real fights

Lane brief: `design/reviews/sim-accuracy/brief.md`. Site from origin/main 563388a64, fork from master 8dacbe217 (the pin). Client on the site: 1.60.1.70790; live client 1.60.1.70888 with its hotfix push (wago.tools, `useHotfixes=1`). The 2026-11-11 "hotfixes, other classes" lane moved Stormstrike, Mental Dexterity, Maelstrom Weapon, the imbue and totem rule and Lightning Overload to Blizzard's 9 and 10 November values from the notes alone; this lane is the first shaman pass against logs since.

**Totals: 2 confirmed engine bugs, both old, both fixed (Lightning Shield's orbs, Flametongue Weapon by level). Neither is on a published card, so no published number moves; the fork's Enhancement and Elemental suite goldens are byte-identical. 23 checks aligned against the client and the logs, among them the first log confirmation of the 1.5 s Lightning Shield cadence, the 3.3 normalization under Stormstrike and the 20% Windfury Weapon proc. The Enhancement card's Earth Shock reserve was re-measured on the hotfixed Stormstrike (the quarter now costs 3.1% at bare 60 and 1.3% at bare 40 against shocking down to 15%, which never runs dry; the card, its notes, the coach note and the guide carry 15%. Two Earth Shock rank rows (rank 3 at 25 and 26) and three races' Lightning Bolt rank 4 row at 39 are gone on the hotfixed numbers; no published raid or bare band number moves on the engine fix alone). Held: the Flametongue over Windfury totem order (one fight), the Rockbiter bleed (one shaman, an item). Open: every 10 November number still waits for a level 60 or post-deploy Stormstrike log. Engine pin 8dacbe217 to 03b13a6a5.**

## Method

1. **Client.** For every spell, talent, aura, proc, totem and imbue the two specs' rotations and builds touch, the live client's hotfixed rows (SpellEffect, SpellMisc, SpellCooldowns, SpellPower, SpellAuraOptions, SpellCategories, SpellClassOptions, SpellLevels, SpellDuration, SpellItemEnchantment, SpellName; 1.60.1.70888, `useHotfixes=1`, one spell a request) against the fork at the pin, and the site's vendored 70790 `spellconst/shaman.json`, `spelltips/shaman.json` and `talents/shaman.json` where they differ. Blizzard's hotfix thread (t/2403040) read through post 8 (11 November, posted 12 November 00:48 UTC) and "Class Tuning Incoming" (t/2404092).
2. **Our API.** Report gxrgvz4yzsi2 (8 November, advanced logging, 104 fights): two shamans in the logging group, an Enhancement shaman at 28 to 31 (Twilight Maul, 3.3) with Stormstrike, Earth Shock rank 4, Flame Shock rank 3, Lightning Shield rank 3, Fire Nova rank 2 and his own Flametongue Totem, and a Restoration shaman (Evocator's Blade, 1.6) who drops Windfury Totem. 44 fights' `events.parquet` with the advanced fields: on a cast the shaman's own power, cost, attack power; on a damage event the target's. `base_amount` is before armor and before the target's +20%, so the Stormstrike debuff shows as amount over base. The report is from before the 10 November hotfix (Stormstrike 6 s).
3. **foreverlogs.gg.** All 25 public reports (2815 to 2853, 5 to 12 November): every damage, cast, buff, debuff, avoidance and resource event from each shaman (`meta_events`, `sourceCharacterIds`, paged; 38,407 events from 18 reports) and their armory captures (weapon, speed, talent points). Fourteen Enhancement shamans with Stormstrike (20 to 47 by their ranks), four with Windfury Weapon, nine with Flametongue Weapon, none at 35 or more with Maelstrom Weapon taken and no Elemental at 31 or more casting Lightning Bolt, so Maelstrom Weapon and Lightning Overload stay unchecked by logs. Every Stormstrike user's cooldown is 6 s: no public report yet starts after the 10 November deploy with a Stormstrike user, so the 100% and 8 s stay on Blizzard's word.
4. **Engine.** The fork's `sim/shaman` (stormstrike, windfury_weapon, flametongue_weapon, rockbiter_weapon, lightning_shield, lightning_bolt, chain_lightning, earth_shock, flame_shock, fire_totems, flametongue_totem, air_totems, talents, items), `sim/core/buffs.go` (Windfury Totem) and `sim/core/flametongue_totem.go`.
5. **Sims** on the Windows box (`tools/remote-test.sh --dir . --engine-sha`, `make simdb` then `spec-breakdown`, one at a time, `nice`, GOMAXPROCS=6, 10,000 iterations, seed 1): the Enhancement raid 60 breakdown at the pin and on the fixed engine; the Earth Shock reserve at 0, 15, 25, 35 and 50% at raid 60, bare 60 and bare 40 (`shamans/enhancement-earth-shock-*.json`, the card with that one condition changed); Rockbiter against Flametongue Weapon at 10 to 29; the Earth Shock rank rows at 25 to 27 and the Lightning Bolt rows at 33, 34, 39 and 40. Scripts in `shamans/group_*.sh`.

## Confirmed and fixed

| # | finding | source | fix (fork) |
|---|---|---|---|
| F1 | **Lightning Shield's orbs carried vanilla's 13/29/51/80/114/154/198.** The client's orb spells (26364, 26365, 26366, 26367, 26369, 26370, 26363) are 13, 24, 40, 64, 96, 134 and 178 at a 0.267 coefficient, the amounts the rank tooltips name. Rank 2 to 7 were 11 to 27% too big. | 70790 and 70888 SpellEffect; `spelltips` ("struck for 40 Nature damage" at rank 3); our 8 Nov report: rank 3 orbs land for 43 to 45 against 12 spell power (vanilla's 51 would be 54); foreverlogs: rank 2 orbs 23 to 26 (three shamans), rank 3 38 to 46 (four), rank 4 62 to 80 (two), rank 5 118 to 134 (one) | `lightning_shield.go` |
| F2 | **Flametongue Weapon dealt a flat vanilla cap over four per rank** (18/26/42/57/85/112), 4.5 a second of weapon speed at level 10 and 14.25 at rank 4's cap. The client's proc dummies (8026, 8028, 8029, 10445, 16343, 16344) grow a level to each rank's MaxLevel: 326 + 19 to 16, 479 + 29 to 24, 716 + 42 to 34, 1144 + 73 to 44, 1876 + 62 to 54, 2498 + 78 to 64, over weapon speeds 1.3 to 4.0 ("$m1/77 to $M1/25"), so a hit adds amount x speed / 100: 3.26 a second at 10, 17.28 at rank 4's cap, 28.10 at 60. | 70790 and 70888 SpellEffect and SpellLevels; the enchant (SpellItemEnchantment 5, 4, 3, 523, 1665, 1666) triggers 436519 at ProcChance 100; foreverlogs: rank 2 on a 3.0 axe lands 17 at about 20 and 20 at its cap (two shamans), rank 2 on a 3.4 hammer 23 (two shamans, cap 22.2), rank 3 on a 3.3 axe 26 at about 28, on a 2.7 two-hander 33 with Mental Quickness spell power | `flametongue_weapon.go` |

Neither spell is on a published card: the Enhancement card casts no Lightning Shield, the leveling kit carries Rockbiter Weapon to 29 and Windfury Weapon from 30, and the Elemental kit an oil. The fork's `TestEnhancement.results` and `TestElemental.results` are unchanged on the fixed engine, and the Enhancement raid 60 breakdown is identical (857.4 +/- 0.7 on both, every action identical).

## Checked, nothing to change

| item | engine | client or log | verdict |
|---|---|---|---|
| Stormstrike 100% on 8 s, Mental Dexterity 33/66/100, Maelstrom Weapon 20% a stack and 6 to 10 a minute | hand values (2026-11-11 hotfix lane) | 70888 and its pushes still carry 150, 6000 and 50/100/150; 408498 now reads 20 and 408505 -20% cast and cost on Lightning Bolt, so the stack size is in the client; the rates stay server side | aligned, logs pending (no post-deploy Stormstrike user yet) |
| Stormstrike normalization | 150% (then) of normalized weapon damage, 3.3 for a two-hander | our 8 Nov report, 121 Stormstrikes against 189 white hits on a 3.3 maul: 1.58 of the white median; foreverlogs, nine shamans with their weapon from the armory and attack power solved from the upper half of white hits: Stormstrike lands 0.3 to 2.4% above 1.5 x (weapon + attack power x 3.3 / 14) on 3.0 to 3.4 weapons and 4 to 5% below it on the 3.8 Corpsemaker, the sign 3.3 normalization predicts | aligned |
| Stormstrike's debuff | +20% on the next Lightning Bolt, Chain Lightning or Earth Shock, one charge, 12 s | 17364 effect 1 aura 271, mask 0x100003, DurationIndex 29; our report: Earth Shock amount 1.19 of its pre-debuff base, Flame Shock ticks 1.00 | aligned |
| Stormstrike's cost | 8% of base mana | our report: 48 to 50 at 1,593 to 1,938 max mana, 25% for Totemic Call (150 to 158) | aligned |
| Windfury Weapon | 20% a main-hand hit, 1.5 s between procs, two strikes with the rank's attack power, special hits under the spell's own name | SpellItemEnchantment 283/284/525/1669 trigger 439431 (ProcChance 20, ProcCategoryRecovery 1500, ProcTypeMask 0x14); the 70790 tooltip "2 extra attacks with 104 extra melee attack power"; foreverlogs logs the strikes as damage under "Windfury Weapon" (439440), 1.7 to 1.8 landed hits a proc and 0.16 to 0.19 procs a landed main-hand hit for three shamans, never as extra attacks (Windfury Totem's are) | aligned; whether the strike glances is undetermined (target armor variation hides the glancing tail even under Stormstrike) |
| Windfury Weapon attack power by level | 46 + 7.2 to 38, 119 + 12.8 to 48, 249 + 8.3 to 58, 333 + 12.5 to 68 | 8233/8236/10484/16361 aura 99, unchanged in 70888 | aligned |
| Windfury Totem | one extra white attack, +95/179/246 attack power, 20%, area aura with no duration | 8516/10608/10610 (aura 99 and Effect 19 with 1 point, DurationIndex 36 = 1 s), 8515/10609/10612 area auras; our report: extra attacks 0.22 to 0.24 a landed main-hand hit for the two players who stood in range, 5 to 6% for the rest; the extra swing lands +9 to +26 over a 2.2 axe's normal swing (95 x 2.2 / 14 = 15) | aligned |
| Weapon imbue and the matching totem since 10 Nov | both work together | every public report with a Windfury Weapon shaman is from before the deploy (his own Windfury Totem never procs for him while Flametongue Totem does, the old rule) | aligned with the old rule in the old logs; the new rule is Blizzard's word |
| Flametongue Totem | amount x speed / 100 a main-hand hit; 548/781/1061/1363 | 8253 and kin; our report: rank 1 lands 18 on a 3.3 maul and 9 on a 1.6 dagger; foreverlogs 17 on a 3.3 staff and 19 on a 3.8 axe | aligned |
| Rockbiter Weapon | attack power by rank and level, threat | 10400/15567/15568/15569/16311/16312/16313 aura 99 and aura 10, unchanged in 70888 | aligned (see Held for a bleed) |
| Earth Shock, Flame Shock, Fire Nova, Searing Totem rows | client rolls, coefficients, cooldowns, costs | 8046 130 + 1.1 at 0.386 (shock category 19), 29228 292 + 3.8 at 0.214 with 80 a 3 s tick at 0.1, 408423/408424 33 + 0.4 and 55 + 0.6 at 0.143, 10436 47 at 0.017; our report: Earth Shock rank 4 at 145 base against the 134 roll plus spell power, Flame Shock rank 3 direct 103 and ticks 26, Fire Nova rank 2 57 to 65 on a 10.0 s cooldown at 120 mana, Searing Totem attacks a median 2.43 s apart (the engine's 2.437) | aligned |
| Flame Shock ticks crit | at spell crit, 1.5x | our report: 3 crits in 86 ticks | aligned |
| Lightning Shield cadence and charges | one orb every 3.5 s at most, three charges | foreverlogs: 98 gaps between a shaman's orbs inside a fight, the shortest 3.51 s, the tenth percentile 3.98; our report: 3.7 s shortest, stacks 3 to 2 to 1 to gone | aligned (the engine's "TODO: does vanilla have an ICD?" is answered) |
| Flurry | 6% a rank for three swings, 15 s, a stack spent a swing | 16256 trait curve, 16257 DurationIndex 8; our report: swings 6.2% faster under Flurry (1/5 by the leveling order at 28), 65 stacks spent over 34 applications | aligned |
| Improved Stormstrike | 50% regeneration while casting for 15 s | 1238931 aura 134, DurationIndex 8; our report: eleven buffs of 14.99 to 15.02 s | aligned |
| Lightning Overload rows | half the bolt, 0.357 | 408477 176.5 + 1.05, 408484 255 + 1.65, both 0.357 in 70888 | aligned (hand rows until the client move, MR46) |
| Lightning Mastery, Call of Thunder, Elemental Fury, Concussion, Convection, Reverberation, Call of Flame, Trial by Fire, Elemental Alacrity, Elemental Focus, Tidal Focus, Ancestral Knowledge, Elemental Weapons, Mental Quickness, Rage of the Farseer, Clearcasting | as modelled | 16166 (-500 ms, -100% Chain Lightning cooldown, 180 s), 16120 +5 on mask 3, 16089 +100 misc 15, 16035 mask 0x100003, 16039 -2, 16040 -200 ms, 16038 +15 and +9000 ms, 16086 +20/-2000 ms/+20000 ms, 16578 -200 ms, 16164 10% to 16246, 16179 +5 melee and spell hit, 17485 aura 137 stat mask 8 (Intellect), 16266 7/5/13 on the three masks, 30812 30/30, 425336 aura 342 30 for 25 s on 180 s, 16246 -100% cost for 15 s | aligned |
| Durations | Stormstrike 12 s, Flurry 15, Elemental Devastation 10, Maelstrom 30, Farseer 25, Lightning Mastery 15, Clearcasting 15, Windfury Totem buff 1, Windfury Weapon buff 1.5 | SpellDuration 29, 8, 1, 9, 63, 8, 8, 36, 65 | aligned |
| Flametongue Totem rank 2 "was not working" (11 Nov hotfix) | rank 2 works | the engine never modelled the bug | nothing to change |
| Lava Burst, Elemental Mastery | not in the tree; Lightning Mastery took their place in 70755 | 70790 and 70888 trees | nothing to change |

## Held

- **Flametongue Totem over Windfury Totem.** The engine lets Windfury Totem outrank Flametongue Totem when both stand (`meleeTotemCategory`). In our 8 Nov report two shamans stood both: in fight 16 Windfury procs stop the moment Flametongue Totem drops (13.6 s, totem at 13.9) and resume the moment Magma Totem replaces it (last Flametongue proc 32.7, Magma at 32.9, Windfury again at 38.4); fight 25 does the same (Windfury to 20.9, Flametongue dropped 21.7, Flametongue procs to 30.6, Searing at 30.6, Windfury at 32.4); fight 13 has no Windfury proc in the 24 s both stand. So Flametongue wins when both stand, at least when it is the newer. One report, one pair of totems; no raid preset drops both, and the Enhancement guide already sends a second shaman's Air slot to Grace of Air. Held for a log with both totems known to stand through a fight; the engine's order is then either Flametongue first or newest first.
- **"Bigger Rock to the Face".** One shaman at 44 or more (foreverlogs 2842) with Rockbiter Weapon rank 6 lands 141 ticks of a 14-a-2 s bleed, spell 1306816, under 41 debuff applications. The client's Rockbiter rows (16314 to 16316 and their enchants) carry no trigger and the spell has no client trigger; its name and its sibling "Rock to the Face" (1309967) read like an item. The armory holds no weapon for him. Held for his gear or a second Rockbiter shaman.

## Open

- **Every 10 November number** (Stormstrike 100% and 8 s, Mental Dexterity 33/66/100, Maelstrom Weapon's 6 to 10 a minute and 20% a stack, Windfury Weapon under Windfury Totem) needs a Stormstrike user logged after the deploy; the three public reports since (2849, 2850, 2853) have none. The first level 60 Enhancement log settles the glancing question on Windfury Weapon's strikes too (a 40% glancing tail on whites against a boss would show).
- **Lightning Overload** (30%, half) and **Maelstrom Weapon** stack rates: no Elemental at 31 or more and no Enhancement at 35 or more with the talent has logged.
- **The advanced log's attack power.** For all five melee players in our report (shaman, warrior, rogue, paladin, healer shaman) a white hit's base is 0.69 to 0.82 of weapon + logged attack power x speed / 14; the field is not the attack power that feeds the swing (or includes something that does not). Ratios inside one log are unaffected; the physical lane's rage and Shank checks used ratios.

## Sims

All on the box, 10,000 iterations, seed 1, Horde unless named (the card's Horde entry is an Orc; `-race` swaps the race on it), engine 5872d4cba (the fixed one; 03b13a6a5 adds only the rotation copy). Errors are the sim's own (0.7 at raid 60, 0.1 to 0.2 below).

**No move.** Enhancement raid 60 on the pin and on the fixed engine: identical to the decimal (every action's DPS, casts, hits and crit share the same). Lightning Shield is not cast and the kit carries Windfury Weapon.

**Earth Shock reserve** (`enhancement-earth-shock-*.json`: the card with the Earth Shock line's mana condition at the named share, the rest untouched), relative to the card's quarter:

| reserve | raid 60 | bare 60 | bare 40 |
|---|---|---|---|
| none | 0.0% (never dry) | +5.4% (dry 9.4 s of 180) | +2.1% (dry 55.4 s of 180) |
| 15% | 0.0% | +3.1% (never dry) | +1.3% (dry 1.9 s) |
| 25% (the card) | 0 | 0 (never dry) | 0 (dry 0.7 s) |
| 35% | 0.0% | -3.6% | -1.8% |
| 50% | 0.0% | -7.6% | -5.1% |

In a raid the reserve is never reached (Judgement of Wisdom and Mana Spring). Without raid buffs the quarter now throws shocks away: Stormstrike on 8 s spends 10% less mana than it did on 6 s (2,462 against the old card's 2,699 at bare 60) and the hotfixed Maelstrom bolts are free. Shocking to empty measures best but leaves a level-40 shaman with nothing for Stormstrike or a totem for a third of a long fight; 15% keeps both paid for, never runs dry, and takes 3.1% at 60 and 1.3% at 40 over the quarter. **Adopted: 15%.** The card, its notes, the coach note, the guide's mana paragraph and `TestEnhancementShocksDownToFifteenPercentOfItsMana` carry it; `make apl-sync` moved the site and fork copies.

**Rockbiter against Flametongue Weapon** at 10 to 29 on the fixed Flametongue (bare band 20, `-consumes` the one imbue, the kit's rule is Rockbiter to 29):

| level | Rockbiter | Flametongue | Flametongue over Rockbiter |
|---|---|---|---|
| 10 | 17.2 | 17.6 | +2.3% |
| 14 | 22.5 | 22.8 | +1.3% |
| 16 | 25.5 | 25.9 | +1.6% |
| 18 | 33.7 | 33.8 | +0.3% |
| 19 | 35.1 | 35.3 | +0.6% |
| 24 | 43.3 | 43.8 | +1.2% |
| 29 | 59.9 | 59.9 | 0.0% |

Flametongue is never worse and never more than 2.3% better. The kit keeps Rockbiter to 29 (a kit change re-ranks band 20; MR55), and the guide's "below 20 Rockbiter's attack power is worth more than Flametongue Weapon" is gone: the two are a percent or two apart, so the guide now says to wear whichever rank is newest.

**Earth Shock rank rows at 25 to 27** (the card's rows: rank 3 at 25 for a dwarf, tauren, troll or Windshaper Skyborne, rank 3 at 26 for every race; `-rank 10414=3` against `=4`):

| level | race | rank 3 | rank 4 | rank 4 over rank 3 |
|---|---|---|---|---|
| 25 | orc | 50.3 | 51.6 | +2.6% |
| 25 | dwarf | 49.8 | 51.0 | +2.4% |
| 26 | orc | 51.6 | 52.9 | +2.5% |
| 26 | dwarf | 51.0 | 52.2 | +2.4% |
| 27 | orc | 54.9 | 57.7 | +5.1% |
| 27 | dwarf | 54.5 | 57.1 | +4.8% |
| 25 | tauren | 49.8 | 50.9 | +2.2% |
| 25 | troll | 50.0 | 51.3 | +2.6% |
| 25 | Windshaper Skyborne | 50.6 | 51.7 | +2.2% |
| 26 | tauren | 51.0 | 52.2 | +2.4% |
| 26 | troll | 51.3 | 52.5 | +2.3% |
| 26 | Windshaper Skyborne | 51.8 | 53.2 | +2.7% |

On the hotfixed Stormstrike (less mana spent) the newest rank wins from the level it arrives. **Both Earth Shock rows are gone.**

**Lightning Bolt rows at the boundaries the overload could have moved** (`-rank 15208=N`):

| level | race | lower rank | newer rank | newer over lower | row |
|---|---|---|---|---|---|
| 33 | orc | 3: 68.0 | 4: 67.3 | -1.0% | rank 3 to 33 holds |
| 34 | orc | 3: 66.8 | 4: 68.0 | +1.8% | rank 4 from 34 holds |
| 39 | orc | 4: 75.2 | 5: 76.3 | +1.5% | orc rank 5 at 39 holds |
| 39 | dwarf | 4: 75.5 | 5: 77.0 | +2.0% | the rank 4 row at 39 is gone |
| 40 | orc | 4: 103.7 | 5: 106.4 | +2.6% | rank 5 from 40 holds |
| 39 | tauren | 4: 74.2 | 5: 74.8 | +0.8% | rank 5 |
| 39 | troll | 4: 74.5 | 5: 75.1 | +0.8% | rank 5 |
| 39 | Windshaper Skyborne | 4: 72.5 | 5: 72.2 | -0.4% | its rank 4 row stays, alone (seed 50001 reproduced seed 1 to the decimal for every race, which the tooling should not do and the integrator may want to look at: `spec-breakdown -seed` gave the same numbers twice) |

The overload is half the bolt at every rank, so the rows only moved where a race's margin was thin: the dwarf's +0.3% for rank 4 at 39 is +2.0% for rank 5 now. **Rank 5 from 39 for every race.**

## Guide and card

wow-player read the Enhancement guide, the card and the coach note in full: **NO-SHIP, then SHIP** on all three. Two MUSTs were taken: the imbue line said Flametongue "adds" to Rockbiter and told the player to wear whichever rank is newest, against the measurement (Flametongue wins at every level from 10 to 24 whichever rank is newer), and the 20 to 60 section still said Windfury replaces Rockbiter. Five SHOULDs were taken: "about 15% of your mana" became "about a sixth" everywhere (a bar has no 15% mark); "a quarter of a long fight" for 55 s of 180 became "about a third"; the mana paragraph lost its patch-note phrasing and tells a raider to shock on cooldown; the card note no longer says a cooldown spends mana; the imbue line sends a fire-immune pull back to Rockbiter. One optional SHOULD taken on the re-read ("not worth a trip to town", not "never"). Left alone: a "--" in an older note on the card's line 4, outside this lane's lines.

Pages and cards changed:

- **Enhancement (`enhancement.md`):** the mana paragraph (a sixth, why the quarter is gone, what shocking to empty costs at 40, the raid); "Levels 1 to 19" (Flametongue from 10, Rockbiter for fire-immune mobs, Windfury at 30); "Levels 20 to 60" (Windfury replaces whatever imbue). Card: the Earth Shock line's condition (15%), its note and the summary entry. Coach note: the Earth Shock line.
- **Elemental:** nothing in the prose. The Lightning Bolt rank list under "Levels 20 to 60" names the orc exception at 39 ("39 or 40 to 44: rank 5, from 39 for an orc and from 40 for every other race"); it now reads "34 to 38: rank 4, and at 39 too for a Windshaper Skyborne" and "39 to 44: rank 5, from 40 for a Windshaper Skyborne"; the card note says the same. No reviewer round for a rank list that restates the rows.

## Measure at refresh

- **MR55, the leveling kit's imbue below 30.** Flametongue Weapon beats Rockbiter by 0.3 to 2.3% at 10 to 24 and ties at 29 on the fixed rows (its 0.1 spell coefficient is the engine's, the client states none); the kit keeps Rockbiter to 29 until the band-20 re-rank can carry a switch. The guide already tells a player to run Flametongue.
- MR42 and MR43 stand; the Earth Shock reserve, the two Earth Shock rows and the Lightning Bolt row at 39 are settled here and need no refresh.

## Sim changes in this landing

- **Fork `lane/sim-shamans`, fast-forward of master 8dacbe217:**
  - 2b5600c35: tests, red on the box (Lightning Shield orbs, Flametongue Weapon by level).
  - 5872d4cba: the fixes.
- **Fork:** 03b13a6a5 adds the synced Enhancement rotation (Earth Shock at 15%). Master fast-forwarded to it.
- **Site pin:** 03b13a6a5.
- **Site:** the card and its two copies, the coach note, the two guides, `rank-choices.json` and its embedded copy, `shaman_earth_shock_gate_test.go`, `rank_choices_shaman_r24_test.go`, `elemental_ranks_test.go` (rank 5 at 39 for a troll), the Enhancement ladder golden regenerated on the box (the walk's Earth Shock share rises at every checkpoint; its low-iteration DPS moves within noise, +1.1% at 30 to -0.7% at 60), the review's run scripts and rotation variants.
- **Shared code:** none. Lightning Shield and Flametongue Weapon are shaman-only.

## Sources

- Blizzard, WoW: Forever Hotfixes, https://us.forums.blizzard.com/en/wow/t/2403040 (posts 5, 6 and 8).
- Blizzard, Class Tuning Incoming, https://us.forums.blizzard.com/en/wow/t/2404092.
- wago.tools DB2 CSVs for 1.60.1.70888 with its hotfix push (SpellEffect, SpellMisc, SpellCooldowns, SpellPower, SpellAuraOptions, SpellCategories, SpellClassOptions, SpellLevels, SpellDuration, SpellItemEnchantment, SpellName, Spell), and 70790 for Ancestral Knowledge.
- foreversixty API report gxrgvz4yzsi2.
- foreverlogs.gg reports 2815 to 2853 and the armory captures of their shamans.
