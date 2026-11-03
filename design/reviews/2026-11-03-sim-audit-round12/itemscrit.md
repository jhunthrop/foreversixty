# Sim audit round 12: itemscrit lane

This lane owns the engine's crit plumbing and item effects. This round covers:
- a fresh crit trace on the landing pin 2606bc05e;
- every item effect on the published BiS sets, checked against 70753;
- the open round-11 item checks;
- the landing's consumable changes.

Branches: engine `lane/r12-itemscrit` (fork, from 2606bc05e); site `wt/r12-itemscrit` (from c6bdc855c).

Client data:
- 1.60.1.70753, fetched fresh from wago.tools on 2026-11-03, shipped and `useHotfixes=1`, 26 tables. The two copies are byte-identical: 70753 has no hotfix push yet.
- `python -m pipeline push-check --build 1.60.1.70753` gave `moved=false` on 2026-11-03.
- The item layer (`raw/hotfixes/`) is seeded from 70707's push.
- Server text from nether.wowhead.com/forever/tooltip, read 2026-11-03.

Scratch files, none committed, all under `scratchpad/r12/itemscrit/`:
- `tools/zzcritdump`: prints each published entry's crit on every table, and every registered spell's physical, magic and heal crit chance after reset and pre-pull.
- `dump/`: that tool's output for every spec at raid 60 and bare 60.
- `py/`: `sp.py` (client spell rows), `census.py` (item effects on published sets), `itemdiff.py` (70707 against 70753), `cons.py` (consumables).

## 1. Crit trace (pin 2606bc05e, raid and bare 60, Horde)

Method: for every spec, the Crit / MeleeCrit / SpellCrit stats after reset, split by source. The sources are:
- item stat 32;
- buffs (Leader of the Pack 3 and Moonkin Aura 3, both aura 290 after the push);
- consumables (Mongoose 2, Sages 2, Squid 1, all aura 290; Wizard Oil 1, aura 57);
- talents and racials;
- Agility and Intellect at the level rate.

The dump also gives each spell's chance on its client defense type. Per-spell bonuses were read the same way: the stone's 2% on melee-defense weapon spells only, and school and talent crit.

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | Every spec's Crit stat sums exactly from its sources at raid 60 and bare 60. A few examples. Retribution: 4 items + 3 + 3 + 2 Mongoose + 6 Conviction = 18. Rogues: items + 6 + 2 + 5 Malice. Feral cat: items + 6 + 2 + 6 Sharpened Claws + 4 Nature's Majesty. Arcane: items + 6 + 3 Arcane Instability. Fury: Berserker Stance's 3 is in Crit. MeleeCrit and SpellCrit are base plus Agility or Intellect at the level rate, plus Cruelty or weapon specialization, which sit on the melee table only. | `dump/*.60.txt`; client: Malice 14138, Recklessness 1719, Berserker Stance Passive 7381 and Arcane Instability 15058 are aura 290; Cruelty 12320 is aura 52 with a weapon requirement | Aligned | - | - |
| C2 | Round 11's Conviction (aura 290) has no side effect on other specs. Holy paladin's heals read 25.03% on the spell table with no leftover heal-only bonus. Retribution's Consecration and Exorcism roll the spell table, and its seals and Judgements roll the melee table plus the stone. Every published paladin band holds a melee weapon (the weapon gate). | `dump/paladin-*.txt` | Aligned | - | - |
| C3 | Stones and oils land where they should. The Elemental stone's 2% (enchant 2506 -> 22755, aura 52) is a per-spell bonus on melee-defense weapon spells and autos only: never Consecration, poisons, spell procs, pets or ranged. Brilliant Wizard Oil (2628 -> 25113: aura 13 36, aura 135 36, aura 57 1) is SpellCrit, so it reaches heals. | dumps; SpellItemEnchantment and SpellEffect 70753 | Aligned | - | - |
| C4 | Round 11's Atiesh, Blackhand's Breadth and Eye of the Beast have no side effects. No published band wears them; the engine's use effects are on their own item ids. | `bisitems.json` | Aligned | - | - |
| C5 | Hunter pets inherit the hunter's every-table crit (Forever's pet rework). They do not also get the raid's Leader of the Pack or Moonkin Aura, because the inherited 15% already includes them. | `hunter/pet.go:209`; dump | Aligned (as in round 10) | - | - |
| C6 | Base crit columns (chancetomeleecrit and chancetospellcrit, CASC 3999262 to 3999265) are still empty for 70753 (re-checked 2026-11-03). | wago.tools CASC API | OPEN (round 11 P2) | - | - |

## 2. Item effects on the published sets (70753)

The 70753 BiS files carry the same 719 items, in the same bands, as 70707's (`bisitems.json`).

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| I1 | **70753 against 70707 (with its push).** No ItemSparse, ItemEffect, ItemXItemEffect, ItemSet, ItemSetSpell, SpellItemEnchantment, SpellAuraOptions or SpellCooldowns row of a published item or its spells changes.<br>- Against 70707 shipped, 70753 adds 40 ItemSparse rows: recipes, plans and patterns, Scorched Rocket Boots, Pouch of Violet Dust, Enchant Boots - Speed and its formula. It changes 4 rows: three prices, and Toughened Leather Armor 2314, which the 70707 push already had and no set carries.<br>- The 469 rows missing against 70707-layered are push rows that 70753 has no push for yet. The item layer still carries them.<br>- The SpellEffect changes are Holy Shield, Orc Command, the removed blessings, Boiling Scream and float printing. None is an item spell. | `py/itemdiff.py` | Aligned | - | - |
| I2 | **Census of non-flat client effects on published items: 59 items.** Each one is either modelled (37 engine item effects plus the client set rows) or has nothing to model. The things with nothing to model:<br>- self-heals and absorbs;<br>- movement speed (Staff of Westfall 1292011, Forsaken Weapons 1309369);<br>- snare reduction, fear resist, Slow Fall, Safe Fall;<br>- cooking speed (Cookie's Stirring Rod 1322312, aura 4 "Cooking speed 25%");<br>- an experience buff (Archmagister's 1292997);<br>- fire damage taken (Petrolspill 1293003);<br>- Fade's cooldown (Truefaith 18388).<br>New this round: **Arena Master 1317285 and Arena Grand Master 1318318** are aura 87, -5% and -6% damage taken. Unmodelled, they would cut a tank's rage from damage taken. But the text limits them to "the Battle Ring and The Maul", so they are right to be out. | `census.txt`; SpellEffect 1350310 and 1351731; Spell text | Aligned | - | - |
| I3 | The relics on published sets match the server text: Idol of the Moon (+33 Moonfire scaling power), Totem of Thunder (+1% Lightning Bolt crit), Libram of Invocation (-5% Seal cost), Idol of the Dream (Rip +2 s), Talons of Wrath (50%, 35 mana). | tooltips 23197, 28854, 228176, 461295, 249442, 1249005 (2026-11-03); client rows | Aligned | - | - |
| I4 | Set bonuses: ItemSet and ItemSetSpell have no change from 70707. Round 11's S1 and S2 (proc flags, amounts) stand. | `itemdiff.py` | Aligned | - | - |

## 3. Round 11's open checks

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| O1 | **Thornspike and poisons.** Three beta-log rogues carry Thornspike (6681): 30169 in reports 2707 and 2708, and 168741 in 2805, 2808 and 2809. None has a poison on: their temporary enchants are 0 or Sharpened +6 (483). So no log can show whether Thornspike's ranged strike (DefenseType 3) sets off a poison. | roster snapshots in `scratchpad/fl/rep`, 119 reports | OPEN (needs a poisoned Thornspike log) | - | - |
| O2 | **One-a-minute rates.** 70753's SpellProcsPerMinute is the same 11 rows. One SpellAuraOptions row points at one: Holy Smite 1326019. Weapon and enchant procs carry no rate column. The beta logs show no proc enchant on any weapon: the most common are +damage, oils, stones and Forever's Revelation, Insight and Recovery. | wago SpellProcsPerMinute / SpellAuraOptions 70753; roster enchants | OPEN (needs level-60 logs) | - | - |
| O3 | **Mindtap Talisman (18371) and Force of Will (11810).** Neither has an ItemSparse row (shipped or item layer), a Wowhead Forever page (404 "Entity not found") or a loot source. The ranker cannot pick them. Mindtap's client link is the use only (1302205), and simdb keeps Era's 11 mp5. Force of Will's rate on being struck is unstated. | ItemXItemEffect 70753; tooltips | OPEN (unreachable) | - | - |
| O4 | **Weapon enchants.** No published set carries a permanent enchant: the ranker sends none. Crusader, Fiery Weapon and Lifestealing match the client rows as in round 10's N1. | SpellItemEnchantment 1900, 803, 1898 | Aligned | - | - |
| O5 | **Stats 45 and 50.** Stat 45 counts as spell power and also grants healing, as the 41/42 pairs show (Glowing Brightwood Staff 812: spell power 50, healing 150). Stat 50 is bonus armor, which the committed simdb still lacks (MR9). Toughness scales item Armor and not BonusArmor (`paladin/tank_talents.go:114`), which is round 9's open D4. | `normalize/gear.py:156-221`; simdb dump | Aligned / OPEN (D4) | - | - |

## 4. Consumables and the landing

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| K1 | **Beast Mastery and Marksmanship lost Juju Might.** The landing moved Juju Might to Survival only, on the reading that it is "melee attack power, so a bow hunter gains nothing". The 70753 client's Juju Might (16329) has two effects: aura 99 40 (attack power) and aura 124 40 (ranged attack power). The engine has always applied both (`core/consumes.go:500`). So a bow hunter's shots lose 40 ranged attack power without it. The landing's Juju Power move is right, since Strength gives a hunter no ranged attack power. Juju Might is back in the ranged group for every hunter, and Survival's own list keeps Juju Power. It recovers the landing's whole -1.0% (BM) and -1.2% (MM). | SpellEffect 693538 and 870621; `TestEveryHunterCarriesJujuMightAndOnlySurvivalJujuPower` failed first (BM and MM had no juju_might), then passed; the sim/request preset and raid-consumes tests pass | CONFIRMED-BUG | landing 252455d15 | S 0206582a6 |
| K2 | **The Lesser Mana Potion restored nothing.** Its item id 3385 had no row in the engine's roll table, so it gave 0 mana. The client's 438 is 320 with variance 0.25, so 280 to 360. In the same table, the Minor Recombobulator carried Era's 150 to 250 mana; the client's 18805 is 128 with variance 0.5, and the server text reads "96 to 160 Mana". The healthstones healed Era's 500, 800 and 1200; the client's 5720, 5723 and 11732 are flat 600, 960 and 1440 (server text "Instantly restores 1440 life"). No preset or kit carries any of them, so only /sim custom picks reach them. | `TestManaConsumablesRestoreTheClientsAmount` and `TestHealthstonesHealTheClientsAmount` failed first (Lesser 0 to 280, Recombobulator 150 to 250, stones 500/800/1200), then passed. The core package passes. | CONFIRMED-BUG (custom only) | old (upstream fdfff4571; no round checked these rows) | E d347b5094 |
| K3 | Every other raid consumable matches the client's amount, duration and table: R.O.I.D.S., Scorpok, Cortex and Gizzard (25, 60 min); Mongoose; Sages; the flasks (1200 health, 2000 mana, 150 spell damage); Greater Arcane 35; Shadow 40; Frost 15; Fire Power 10; Greater Defense 450 bonus armor; Fortitude 200; Mageblood 12 and Greater 20 mp5; Greater Cleric's 40; the Forever foods (Nightfin 22, Baked Salmon 28 spell damage, Sage's Tea 44 healing, Salmon 40 AP and RAP, Dumplings 20 Strength, Squid 1% every table); the Brilliant oils; the stones; the potions (Major Mana 1350 to 2250, Mighty Rage 45 to 75 and 60 Strength for 20 s, the Frenzy and Spellblasting 30-second potions); and the Demonic Rune (900 to 1500, 800 health). | `cons.txt`; `core/consumes.go` | Aligned | - | - |
| K4 | **Elixir of Holy Power (21546), a consumable the engine does not have.** Forever turned Era's Elixir of Greater Firepower into "Elixir of Holy Power": spell 1310077, +40 Holy spell damage for 30 min, level 40, made from recipe 21547. The engine has no Holy power consumable, so the preset cannot offer it. Retribution's spell power weight is 0.21 DPS a point at raid 60 (bis weights), so it is worth roughly 1% for Retribution and more threat for Protection. Adding it needs a new proto field (`Consumes`) in the fork, its vocabulary row on the site and a preset row. protoc-gen-go is not installed here. | ItemSparse 21546; SpellEffect 1339308 (aura 13, 40, misc 2); tooltips item/21546 and spell/1310077 (2026-11-03) | OPEN (feature; routed) | - | - |
| K6 | **An item's haste counted twice** (found by the casters lane, X2, and routed here). `applyEquipment` added an item's and its enchant's MeleeHaste and SpellHaste to the unit's stats, which the swing and cast speeds read. It then also multiplied the speed multipliers by the same amounts. So a 15% item (Shard of the Defiler 17142, "attack speed and casting speed by 15%") swung and cast 32.25% faster. Now it counts once. No published band carries haste (`haste_on_items` false; no published item has a haste stat in simdb), so no published number moves. On custom gear, Combat with Shard of the Defiler in the main hand goes 799.1 -> 735.4 at raid 60 (-8.0%, Horde, seed 1, 3,000 iterations). The fork's Enhancement Phase 3 and Phase 5 gear (Enchant Gloves - Minor Haste 931, Arcanum of Rapidity 2543) moves 2 to 3%, and its golden is regenerated. | `TestAnItemsHasteCountsOnce` failed first (swing and cast speed 1.3225, want 1.15), then passed. The core, shaman (with_db) and feral suites pass; only `sim/shaman/enhancement/TestEnhancement.results` moves (its Phase 3 and 5 rows). | CONFIRMED-BUG (custom only) | old (upstream bda658a24) | E 0a879f7f6 |
| K5 | The landing's other preset rows (the Dense stone from 35, the improved Mana Spring, the bear's notes) match the client: Dense Sharpening Stone 12404 and Weightstone 12643 need level 35; enchant 1643 is Sharpened +8. | ItemSparse; SpellItemEnchantment 1643 | Aligned | - | - |

## Per spec

Crit tables (section 1) are aligned for all 28 specs at raid and bare 60. Item effects (section 2) are aligned for all. Beast Mastery and Marksmanship have one confirmed bug each, K1. K2 and K6 are engine-wide and reach no published spec.

## Before/after

Engine `lane/r12-itemscrit` d347b5094 with site `wt/r12-itemscrit` 0206582a6, against `$R/before.tsv`. Raid 60, 10,000 iterations.

| spec | alliance s1 | alliance s50001 | horde s1 | horde s50001 | absolute (alliance s1) |
|---|---|---|---|---|---|
| hunter-beast-mastery | +1.01% | +1.01% | +1.02% | +1.02% | 623.0 -> 629.3 |
| hunter-marksmanship | +1.15% | +1.15% | +1.15% | +1.14% | 610.7 -> 617.7 |
| hunter-survival | 0.00% | 0.00% | 0.00% | 0.00% | 741.1 -> 741.1 |

These are the round-11 pin's numbers again: the landing's -1.0% and -1.2% came only from Juju Might. Bare bands do not move, because the raid preset is the only one with Juju. No other spec moves: K2 touches no preset, kit or published set.

## Files to regenerate

- Fork golden regenerated in E 0a879f7f6: `sim/shaman/enhancement/TestEnhancement.results` (only its Phase 3 and 5 rows, the haste-enchanted gear).

- `sim/request/testdata/ladder/hunter-beast-mastery.golden.md` and `hunter-marksmanship.golden.md` (the raid rung moves about +1%).
- `sim/request/raid-consumes.json` is already regenerated in S 0206582a6 (`go run ./internal/genraidconsumes`). Its committed-copy test requires it. `published-presets.json` did not change.
- The web copy (`web/scripts/sync-raid-preset.mjs`) syncs at build.

## Measure at refresh

- Beast Mastery and Marksmanship at raid 60, every slot: their stat weights and picks were ranked without Juju Might's 40 ranged attack power.

## For other lanes

- **Integrator:** S 0206582a6 moves BM and MM at raid 60 (table above). E d347b5094 and E 0a879f7f6 move nothing published. E 0a879f7f6 moves the fork's TestEnhancement golden, which is already regenerated.
- **Fork maintainers / platform:** K4, Elixir of Holy Power for Retribution and Protection paladins. It needs a new `Consumes` proto field (fork), a vocabulary id and a preset row (site). Measure it once it exists.
- **Auras:** none from this lane. The warlock lane's Orc Command note went to auras directly.

## Still open

- C6: the base crit columns, still empty in the 70753 CASC.
- O1: Thornspike and poisons (no poisoned Thornspike in the beta logs).
- O2: the one-a-minute rate (no PPM data, no level-60 logs).
- O3: Mindtap Talisman and Force of Will (unreachable: no item row, page or source).
- O5: Toughness on bonus armor (round 9's D4); it matters once the refresh adds stat-50 armor.
- K4: Elixir of Holy Power (a feature).
- Crit suppression on rating crit (round 11's open item) is unchanged.
