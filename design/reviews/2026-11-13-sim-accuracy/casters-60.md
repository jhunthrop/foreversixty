# Sim accuracy 2026-11-13: Balance and Elemental at 50 and 60

Lane brief: `design/reviews/sim-accuracy/brief.md`. Lane `sim-casters60`. Site from origin/main 9b2facbc7 (client 1.60.1.70938 active since a4bce48ae), fork at the pin b0a43f004 (no engine follow moved it during the lane). Scope: Balance druid and Elemental shaman at 50 and 60, bare and the raid 60 preset. They sit at the bottom of the raid 60 DPS list, about 33% behind the top, and the question is whether that gap comes from the game or from the model.

**Totals: 0 engine bugs, 0 site changes. The 70938 client and its hotfix push change nothing that either rotation, build, proc or talent touches. The engine reads every row the way the client has it, and neither spec runs out of mana in the raid fight. The gap is real: it comes from the raid preset's physical-only buffs and armor debuffs, not from a caster defect. No readable log holds a Balance druid or an Elemental shaman above level 34. Engine pin b0a43f004, unchanged.**

## Method

1. **Client.** wago.tools 1.60.1.70888 and 1.60.1.70938 with `useHotfixes=1`, fifteen spell tables. The diff covers every SpellClassSet 7 (730 druid spells) and SpellClassSet 11 (705 shaman spells) spell, every talent rank and two hops of trigger spells (`casters-60/client_diff.py`, the 2026-11-12 casters lane's script; output in `client-diff-druid.txt` and `client-diff-shaman.txt`). The lane also compared the vendored `data/builds/1.60.1.70938/spellconst` and `talents` against 70790, and read the 70938 `raw/wago-push` entries (pushes 113270 to 114036 are new since 70790).
2. **Blizzard.** "WoW: Forever Hotfixes" (t/2403040), posts 7 to 9 (12 Nov):
   - Post 8 fixes Flametongue Totem rank 2. That is Enhancement's totem; the 70938 row 8250 now carries its trigger 8248.
   - Post 9 restates the 10 Nov shaman changes, which the engine already carries.
   - There is nothing for Balance, nothing new for Elemental, and nothing on spell hit, spell crit, mana or raid buffs.
3. **Engine.** The lane read the fork's druid, shaman and core changes between the last caster pins (4f8bb748d for Balance, 4c595ee6b for Elemental) and b0a43f004: Hurricane's tick crit, Concentration Aura's pushback, `NeverCrits` for Deep Wounds and Ignite, and the relic weapon skill. None of them reaches either card. The two lanes before this one read every Balance and Elemental spell, talent and proc against 70790 and 70888 (`2026-11-12-sim-accuracy/casters.md`, `shamans.md`, `shamans-60.md`). This lane re-read the hit table, crit, Curse of the Elements, Judgement of Wisdom and Lightning Mastery at the pin.
4. **Sims.** All on the Windows box (`tools/remote-test.sh --dir sim --pin-engine`), one job at a time, `nice -n 15`, GOMAXPROCS=5, 10,000 iterations, seed 1:
   - `casters-60/baseline.sh`: both specs, raid 60, bare 60 and bare 50, both factions (`baseline.txt`).
   - `casters-60/variants.sh`: Elemental's Chain Lightning on a single target (`variants.txt`).
5. **Logs.**
   - Our seven backfilled reports (1,178 fights), read only through the Cloud SQL proxy and the downloaded fight files (`casters-60/logs/scan_own.py`, `casters_own.py`, `ele_own.py`). The scripts match by spell id, since three reports come from a Portuguese client.
   - foreverlogs.gg (`/api/reports/public`, 13 Nov): still 26 reports, up to 2854. Report 2855 and later answer 403.
   - Warcraft Logs: no Forever partition on any host (the rogues and hunters lanes checked on 13 Nov).

## Engine against the client

| item | 70938 (hotfixed) against 70888 | engine at b0a43f004 | verdict |
|---|---|---|---|
| Starfire, Wrath, Moonfire, Insect Swarm, Hurricane, every druid talent and its trigger | float formatting only (the CSV now rounds to 11 places); no value, cost, cooldown, duration, mask or level moves | the 2026-11-12 reads stand | aligned |
| Nature's Grace 16886 | aura 65 +10 (cast speed) and aura 108 misc 21 -10 (global cooldown), 3 s | 10% cast speed and a 150 ms global for 3 s on a non-periodic crit; 18.3 procs and 27% uptime at raid 60 | aligned |
| Eclipse 408248, Moonkin Form 24858, Moonkin Aura 24907, Omen of Clarity, Vengeance, Moonfury, Improved Starfire, Nature's Reach | unchanged | as read on 2026-11-12 | aligned |
| Wild Growth (408120, 1238214, 1238215) | coefficient 0.033 to 0.066, InterruptFlags 8 to 15 | the healers' hand row (MR58) | not Balance |
| Lightning Bolt and its overload rows, Chain Lightning and its overload rows, Earth, Flame and Frost Shock, Lava Burst 1238300 | float formatting only; the vendored spellconst now carries the overload rows the engine had by hand (147.5 + 0.8 and 176.5 + 1.05 at 0.357, Chain Lightning 255 + 1.65) | the same | aligned; MR46's overload rows can now read the generated tables |
| Lightning Overload, Lightning Mastery 16166, Call of Thunder, Elemental Fury, Concussion, Elemental Alacrity, Tidal Focus, Elemental Focus | unchanged | Lightning Mastery takes 0.5 s off Lightning Bolt and removes Chain Lightning's cooldown for 15 s, every 180 s; the rest as read on 2026-11-12 | aligned |
| Lava Burst, Elemental Mastery | not in the tree | not castable (the rotation search's insert has no effect) | aligned |
| Flametongue Totem rank 2 (8250) | effect 0 now triggers 8248 (was 0), the 11 Nov fix | Enhancement's totem; no Elemental card drops it | out of scope, noted |
| Spell hit at +3 | base miss 17%, 1% floor | `BaseSpellMissChance` 0.17 at a 3-level gap, `minSpellMissChance` 0.01; Forever merges melee and spell hit into one stat | aligned |
| Spell crit at +3 | no suppression for spells in Classic | none applied (melee carries the 4.8%) | aligned |
| Curse of the Elements (raid) | Forever's curse: aura 87 on misc 126, every magic school | +10% on arcane, fire, frost, holy, nature and shadow, 75 resistance | aligned; Starfire and Lightning Bolt both get it |
| Judgement of Wisdom | procs on spell hits | any direct hit that lands | aligned |

## Mana on the 3 minute raid fight

| cell | DPS | time out of mana | where the mana goes |
|---|---|---|---|
| Balance raid 60 Alliance | 737.5 | 0.1 s | spent: Starfire 9,235, Moonfire 3,507, Insect Swarm 1,356, Wrath 792. Gained: regen 5,409, Judgement of Wisdom 2,590, Demonic Rune 2,334, Innervate once |
| Balance raid 60 Horde | 722.1 | 0.1 s | |
| Elemental raid 60 Alliance | 748.0 | 1.6 s | spent: Lightning Bolt 15,122, Flame Shock 2,874, Searing Totem 623. Gained: regen 5,227, Judgement of Wisdom 3,446, Major Mana Potion 2,789, Demonic Rune 2,389 |
| Elemental raid 60 Horde | 749.6 | 1.4 s | |
| Balance bare 60 / 50 (Alliance) | 376.0 / 247.3 | 0.7 / 2.9 s | |
| Elemental bare 60 / 50 (Alliance) | 301.1 / 186.8 | 40.2 / 43.9 s | the card drops to the rank 8 bolt once mana falls behind the fight, and the bar still runs dry for about a quarter of the fight |

The breakdowns reproduce the published `set_dps` (raid: 737.9, 722.9, 748.7, 750.4). Neither spec runs out of mana in the raid, so mana is not what holds them down. Elemental without raid buffs does run dry, and that is why its bare figures sit last; this matches the guide's mana section.

**Logged casters at the highest levels:**

- **Balance druid, level 32:** the highest one casting in a fight. 16 Moonfire, 16 Insect Swarm, 1 Starfire, 5 Nature's Grace. Mana goes as low as 3% between drinks. Moonfire ticks crit 12 times in 43 and Insect Swarm ticks 6 in 58, as the DoT crit rule says.
- **Elemental shaman, level 34:** the highest one. 44 Lightning Bolts, 5 crits, no overload: Lightning Overload needs 31 talent points and he has not taken it. Mana never drops below 40%.

Both are dungeon pulls with drinking in between, so neither says anything about a 3 minute raid fight. No Balance druid or Elemental shaman at 35 or above appears in any log.

## Hit, crit and the raid buffs

- **Hit.** Elemental's raid set is at the cap: Lightning Bolt misses 1.0% (Tidal Focus gives 5%, gear the rest). Balance's set is not: Starfire misses 7.3% and Wrath 7.5%, with Nature's Reach giving 4% and gear about 5.7% against the 16% cap. The published weights price this correctly. A point of hit is worth 0.65 spell power to Balance, because about a third of its damage is Moonfire and Insect Swarm ticks, which never miss. The weights' hit-to-cap line says 7% to go.
- **Crit.** Starfire 20.9%, Wrath 20.3%, Moonfire 32.0% (Improved Moonfire), Insect Swarm ticks 22.4%, Lightning Bolt 29.6% (Call of Thunder and Thundering Strikes), Flame Shock 23.7%. Moonkin Aura and Leader of the Pack give 3% once between them.
- **Raid buffs a caster receives.** Buffs: Arcane Intellect, Mark of the Wild, Fortitude, Divine Spirit, Blessing of Wisdom, Heart of the Lion (10% to every stat), Moonkin Aura, improved Mana Spring. On the boss: Curse of the Elements and Judgement of Wisdom.
- **Raid buffs only physical damage receives.** Everything else in the preset: Blessing of Might, Battle Shout, Trueshot Aura, Heart of the Lion's attack power, Windfury Totem, Strength of Earth, Sunder Armor, Faerie Fire, Curse of Recklessness, Hunter's Mark, and Judgement of the Crusader's Holy damage.

## Elemental's Chain Lightning on one target (checked, holds)

The rotation search only inserts spells the card does not already hold. The card has Chain Lightning for packs, so the search never tried it on a single target, where its first hit is bigger than a Lightning Bolt for the same 2.0 s cast. Measured (`variants.txt`):

| variant | raid 60 Alliance | raid 60 Horde | bare 60 Horde | bare 50 Horde |
|---|---|---|---|---|
| the card | 748.0 | 749.6 | 299.6 | 186.2 |
| Chain Lightning on its cooldown | -13.6% (30 s out of mana) | -13.6% | -17.7% | -18.6% |
| only inside Lightning Mastery | -3.5% | -3.5% | -9.3% | -14.7% |
| while mana is ahead of the fight | -2.8% | -2.7% | -2.3% | 0.0% |

Each Chain Lightning lands about 20% more than a bolt, overloads included, but costs 495 mana to the bolt's 198. Under Lightning Mastery the bolt is also the faster cast. The card is right, and nothing changes.

## Is the gap real

Yes. The raid 60 Alliance figures:

- Marksmanship 1102 and Feral 1090 at the top.
- The casters: Destruction 887, Fire 821, Demonology 815, Arcane 801, Shadow 798, Affliction 760, Elemental 749, Balance 738, Frost 696.

Balance and Elemental sit inside the caster group, not apart from it: 10% under Fire, 16 to 17% under Destruction, and above Frost. What separates the casters from the top is the raid preset, which reflects the game's raid:

- **Raid over bare (60, Alliance).** The physical specs gain 2.0 to 3.0 times from raid buffs: Arms 2.97, Combat 2.88, Fury 2.75, Feral 2.34, Marksmanship 2.28. The casters gain 1.6 to 2.05 times: Frost 1.64, Destruction 1.79, Shadow 1.89, Balance 1.96, Fire 2.05. Elemental's 2.48 is mostly mana relief, since it runs dry for 40 s without raid buffs.
- **Bare 60, with no buffs in play.** Balance is mid-pack at 376: 24% under Destruction, and ahead of Fury, Retribution, the rogues and Arms. Elemental is last at 301, held down by mana.
- **Nothing is missing.** The preset has every buff a Phase 1 raid gives a caster. Forever has no world buffs, and Curse of the Elements already covers Arcane and Nature. Nothing the casters receive is under-read by the engine.

What is specific to the two specs:

- **Balance** spends about a fifth of its mana and a quarter of its casts keeping Moonfire and Insect Swarm up. The searches keep both, because dropping either costs more. Its Phase 1 gear also leaves it 7% short of the hit cap.
- **Elemental** casts almost nothing but Lightning Bolt and pays for it in mana, so adding Chain Lightning, even inside Lightning Mastery, loses damage.

Neither is a model gap.

## Open

- A Balance druid or Elemental shaman at 50 or 60 in any log, to check the cast mix, Nature's Grace and overload rates, and mana on a long fight. The earliest source is the level 60 dungeons; raids open 9 December.
- MR46: Lightning Overload's rows are now in the vendored 70938 spellconst, so the hand values can move to the generated tables (tests unchanged).

## Sources

- wago.tools db2 CSVs, builds 1.60.1.70888 and 1.60.1.70938, `useHotfixes=1` (fetched 13 Nov).
- `data/builds/1.60.1.70938/{spellconst,talents,raw/wago-push}`, `data/curated/presets.json`, `data/curated/apl/{druid-balance,shaman-elemental}.json`.
- Blizzard, "WoW: Forever Hotfixes" t/2403040, posts 7 to 9.
- Our reports (7 reports, 1,178 fights; private reports in aggregate only), the foreverlogs.gg public list (13 Nov), Warcraft Logs (no Forever partition).
