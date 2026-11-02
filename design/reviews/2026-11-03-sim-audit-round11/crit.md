# Sim audit round 11: crit lane

The crit lane owns the engine's crit plumbing. This round: owner ruling 2 (Conviction on paladin heals follows the client), round 10's late find (Blackhand's Breadth and Eye of the Beast), a sweep of every published-BiS item and every enchant for the same mistake, and a fresh look at base crit, Agility and Intellect ratios, crit suppression and crit multipliers.

Engine branch `lane/r11-crit` (fork, from cfcb76e80). Site branch `wt/r11-crit`. Client 1.60.1.70707 with the hotfix push layered in (raw tables from round 10's crit lane; `SpellShapeshift` fetched from wago.tools 2026-11-02 to rebuild simdb locally; wago.tools' hotfix search read 2026-11-02 for `SpellEquippedItems`, `SpellClassOptions`, `PlayerExpectedStat` and record 698173). Local simdb only; no simdb.bin, items.json or other nightly output is committed.

## 1. Conviction on paladin heals (owner ruling 2)

### The client evidence

| table | row | column | value |
|---|---|---|---|
| Talent | 1411 (node 105703 in the trait tree) | SpellRank_0..4 | 20117, 20118, 20119, 20120, 20121. Forever's trait tree uses 20117 for all three ranks (talents/paladin.json: 2%, 4%, 6%) |
| SpellEffect (shipped 70707) | 698173 (spell 20117, index 0) | Effect / EffectAura | 6 (apply aura) / **52** (melee and ranged crit) |
| SpellEffect (hotfix push 113139, first seen by wago.tools 2026-10-31 10:59) | 698173 | Effect / EffectAura | 6 / **290** (crit on every table: melee, ranged, spell, heal) |
| SpellEffect 698173 (both) | | EffectMiscValue_0 | 0 (no school mask) |
| SpellEffect 698173 (both) | | EffectSpellClassMask_0..3 | 0, 0, 0, 0 (no class mask) |
| SpellEffect 698173 (both) | | EffectBasePointsF | 5 (the tooltip's raw value; the trait points give 2/4/6) |
| SpellClassOptions | none for 20117 | | no SpellClassSet, no mask |
| SpellEquippedItems | 180908 | EquippedItemClass / Subclass | 2 (weapon) / 173555: axes, maces, swords (one and two hand), polearm, staff, fist, dagger, spear. No ranged weapon. Not touched by the push (all 21 of 70707's pushed SpellEquippedItems entries checked) |
| SpellMisc | 330259 | Attributes_0 | 464 (passive) |
| Wowhead Forever tooltip, spell 20117 (read 2026-11-02) | | text | "Requires Melee Weapon. Improves your chance to get a critical strike with all spells and attacks by 5%." |

The paladin heals, for the record (SpellClassOptions, SpellClassSet 10 on every row; all school 2 holy, DefenseType 1):

| heal | spell ids | SpellClassOptions rows | SpellClassMask_0..3 |
|---|---|---|---|
| Holy Light 1-9 | 635, 639, 647, 1026, 1042, 3472, 10328, 10329, 25292 | 43461, 43478, 43512, 43504, 43518, 43746, 44346, 44383, 47101 | -2147467264, 0, 1024, 0 |
| Flash of Light 1-6 | 19750, 19939-19943 | 46112, 46212, 46213, 46038, 46039, 46040 | 1073750016, 0, 0, 0 |
| Holy Shock heal 1-4 | 1311605, 25914, 25913, 25903 | 84401, 47232, 47231, 47245 | 2097152, 65536, 0, 0 |
| Lay on Hands 1-3 | 633, 2800, 10310 | 43437, 43704, 44373 | 32768, 0, 0, 0 |
| Light's Vigil heal (Forever) | 1310912, 1311591, 1311596 | 84357, 84391, 84396 | 0, 0, 0, 67108864 |
| Seal of Light heal | 20167, 20333, 20334, 20340 (and Wisdom's 20168, 20350, 20351) | 46209, 46289, 46599, 46595 (46306, 46395, 46446) | 262144, 0, 0, 0 |

Conviction carries no class mask and no school mask, so no bit-by-bit check applies: it is a flat crit aura. Aura 290 is every-table crit (round 9 law: "aura-290 every table"), and every heal rolls the spell table, so **the client covers the heals: the heal crit stays.** The weapon row only asks that a melee weapon be worn; every published Holy paladin band wields a mace or sword (bis/paladin-holy.md: Verigan's Fist, Death Speaker Scepter, Hand of Edward the Odd, Sword of the Fallen).

### The side effect: Conviction reaches the Holy damage spells too

The same row that keeps the heals says Conviction reaches Consecration, Holy Shock, Holy Wrath and every other spell. Round 10's C1 took it off them on beta logs (138k Consecration ticks). Those logs were the shipped aura 52: every cached log round 10 read starts 2026-09-18 and ends 2026-10-23, and the push that turned 52 into 290 was first seen 2026-10-31. It also explains round 10's open puzzle ("why Thundering Strikes reaches spells and Conviction does not"): Thundering Strikes was aura 290 in the shipped table all along, Conviction was not.

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| CR1 | **Conviction is crit on every table.** The engine gave it to melee and ranged (MeleeCrit) plus a heal-only bonus, so Consecration, Holy Shock damage, Holy Wrath and Exorcism's roll did not get it. The client's pushed row is aura 290 with no class or school mask: crit with every attack, spell and heal. It is now plain crit, like Thundering Strikes. Heals keep the same 6% (the heal number is unchanged); the Holy damage spells gain it. | SpellEffect 698173 shipped aura 52, push 113139 aura 290 (table above); Wowhead Forever tooltip "all spells and attacks"; log dates against the push date. `TestConvictionIsCritOnEveryTable` failed first (spell crit +0, want +6); `TestConvictionAddsTwoPercentCritARank` now reads every-table crit. TestProtection's golden moves (dps 575.96 to 577.60 on the first row). | CONFIRMED-BUG | R10 E fc8c92a4e (C1; the logs it rested on predate the push) | E d955cc903 |

Exorcism stays at no crit (the hybrids lane's round-10 fix, 0 of 641 logged): that is a separate rule about the spell, not about Conviction.

## 2. Blackhand's Breadth, Eye of the Beast and the sweep

Neither trinket has an ItemSparse row in 70707 (shipped or hotfix), so the simdb's classic-db supplement supplies them from cmangos 1.12: Era's 2% crit (7598, aura 52; 18382, aura 71 all schools), which the supplement maps to every-table crit. The client links both to Forever's own rows:

| item | ItemXItemEffect | ItemEffect (TriggerType, CoolDownMSec, CategoryCoolDownMSec) | spell | client effect |
|---|---|---|---|---|
| Blackhand's Breadth 13965 | 163702 | 237976 (1 equip) | 1318954 "Increased Critical 1 - Melee and Ranged" | SpellEffect 1352661: aura 52, 1 |
| | 170099 | 231535 (0 use, 300000, 20000) | 1318944 "Blackhand's Breadth" | 1352648: aura 306, 5; 20 s (SpellDuration 18). Tooltip: "Marks your target, increasing your Critical Strike chance against them by 5% for 20 sec. This effect is doubled against Orcs and Ogres." |
| | 170100 | 237977 (1) | 1318945 | aura 4 (dummy) |
| Eye of the Beast 13968 | 163703 | 237965 (1 equip) | 1318956 "Increased Critical 1 - Spell" | 1352663: aura 57, 1 |
| | 170088 | 231536 (0 use, 300000, 20000) | 1287840 "Eye of the Beast" | 1305615: aura 55, 7; 20 s. Tooltip: "Increases your Spell Hit chance by 7% for 20 sec." |
| | 170101 | 237978 (1) | 1318846 | aura 4 (dummy) |

Wowhead Forever's spell tooltips (read 2026-11-02) say 1.0% melee and ranged, 1.0% spell. Neither item has a Forever item page (Wowhead: entity not found).

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| CR2 | **Blackhand's Breadth and Eye of the Beast carried Era's 2% crit on every table.** Where the client links any crit to a classic-db item, the client's crit now replaces Era's: Blackhand's Breadth 1% melee and ranged, Eye of the Beast 1% spell. A local simdb rebuild changes exactly these two items. | Table above. `test_the_clients_own_crit_link_replaces_classic_dbs_crit` failed first. | CONFIRMED-BUG (neither is picked) | old (the classic-db supplement, S 5224e9de3; round 10 found it late) | S eff02da4c |
| CR3 | **Their use effects were not registered.** Blackhand's Breadth: +5% crit against the target for 20 s (the type-neutral target is no Orc or Ogre, so 5%); Eye of the Beast: +7% spell hit for 20 s on the magic schools. 5 minute cooldown, 20 s trinket category. | Table above. `TestBlackrockCritTrinketsUseEffects` failed first (no use effect). | CONFIRMED-BUG (neither is picked) | old | E e91c17c49 |
| CR4 | **Atiesh's spell crit read as a rating.** Wowhead's Forever planner states Atiesh (22589) as `splcritstrkpct` 2, its one such item (Era's equip spell is aura 57, spells only). The key sat with the rating keys, so the engine got 2 / 14 = 0.14% every-table crit. It is now 2% spell crit; the planner shows 2% crit. Only Atiesh moves. | `wowhead_items.STAT_KEYS`; simdb rebuild diff. `test_spell_crit_percent_is_spell_crit_undivided` failed first. | CONFIRMED-BUG (raid item, not in a band) | old (S 451441601) | S 3ff4491a4 |
| CR5 | **Forever's scopes gave their bonus to every table, and Biznicks counted twice.** SAF-T Ultra Precision Scope (enchant 8720 -> 1310308 "Ranged Crit Bonus +02", aura 290, 2) and Poultry Precision Scope (8057 -> 22780 "Ranged Hit Bonus +3", aura 54, 3) carry SpellEquippedItems class 2, mask 262156 (bow, gun, crossbow): the bonus is the scoped weapon's. simdb gave them as every-table crit and hit. Biznicks 247x128 Accurascope (2523, also 22780) had 3% every-table hit in simdb and the engine's own 3% on ranged attacks. simdb now leaves a ranged-weapon-only equip spell's hit and crit to the engine (and fetches `SpellEquippedItems` from now on); the engine applies 8720 and 8057 to ranged attacks. A local rebuild changes exactly 2523, 8057 and 8720. | SpellItemEnchantment 8720, 8057, 2523; SpellEquippedItems 256731 (1310308) and 181192 (22780). Tests `test_a_ranged_weapon_enchants_hit_and_crit_are_left_to_the_engine` and `TestForeverScopesReachRangedAttacksOnly` failed first. | CONFIRMED-BUG (no published preset carries a scope) | old | S 7bac21279, E e183994e8 |

### The sweep: every published-BiS item and every enchant

- **Published BiS items (719 picked items, every spec and band, alternatives excluded).** 42 carry crit in simdb. 35 are ItemSparse crit rating (stat 32, every table: round 9 law). The other 7 have no ItemSparse row and come from Wowhead's Forever planner as `critstrkpct` (Savage Gladiator Chain 11726, Kindling Stave 11750, Band of the Penitent 13217, Vambraces of the Sadist 13400, Mask of the Unforgiven 13404, Eye of Theradras 17715, Shroud of the Nathrezim 18720). Their Era rows are aura 52 or 71. **Aligned**: of the 335 items both Wowhead's Forever planner and the 70707 client carry with percent crit, every one is crit rating (stat 32) in the client, whether Era's row was aura 52 (222), 71 (104) or both (9). Forever re-itemised Era's percent crit as every-table crit rating. The same rule keeps classic-db's other crit items on every table (CR2 changes only the two the client links).
- **Every simdb item with crit (1,162 in a local rebuild).** After CR2 and CR4, 1,159 are every-table crit, Blackhand's Breadth is melee and ranged only, Eye of the Beast and Atiesh spell only. No item carries a table split it should not.
- **Every simdb enchant with a crit aura (31).** Aura 52 rows (Critical Hit +1% to +5%, 195-198 and 436; 2717) are melee crit; aura 57 rows (Firestones 1803, 1823-1825; Brilliant Wizard Oil 2628; 2721; 7644; 7650; 7825) are spell crit; rows with both (7664, 7708, 7746, 7791, 7824, 7883, 7884) are every-table crit for the shared part; stat 32 (8718 Precision) every table. All aligned with the client, except 8720 (CR5). The Elemental Sharpening Stone's 2506 is melee crit in simdb, but the engine applies the stone through its consumable path (round 10's per-spell bonus), not this row. Fanaticism (6967, 7088) is per-school spell crit, dropped as the engine models it: aligned (a Season of Discovery rune, never published).

## 3. Crit plumbing sweep (sim/core)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | Agility and Intellect crit per point, every class at every level 1 to 60, against the client's PlayerExpectedStat (ContentSetID 0): 0 mismatches. At 20/30/40/50/60, % per point (Agility; Intellect): warrior 0.128/0.096/0.076/0.061/0.050; paladin 0.123/0.094/0.074/0.061/0.051 (0.042/0.031/0.025/0.020/0.017); hunter 0.066/0.042/0.030/0.023/0.019 (0.040/0.030/0.024/0.020/0.017); rogue 0.124/0.077/0.056/0.043/0.035; priest 0.080/0.071/0.065/0.057/0.050 (0.058/0.037/0.027/0.021/0.017); shaman 0.107/0.087/0.072/0.060/0.051 (0.052/0.036/0.026/0.021/0.017); mage 0.075/0.069/0.062/0.056/0.051 (0.059/0.037/0.028/0.021/0.017); warlock 0.103/0.083/0.070/0.059/0.050 (0.054/0.036/0.026/0.020/0.017); druid 0.114/0.089/0.072/0.059/0.050 (0.053/0.035/0.026/0.021/0.017). The site's and the fork's copies of the table are byte-identical. No hotfix push row for PlayerExpectedStat or ChrClasses at 70707. | `py/gtcheck.py`; wago hotfix search "70707 playerexpectedstat" (0 entries) | Aligned | - | - |
| P2 | Base crit columns. chancetomeleecrit(base) and chancetospellcrit(base) (CASC 3999262 to 3999265) still answer 200 with an empty body for 70707 (re-checked 2026-11-02). They are GameTables, not DB2, so no hotfix can fill them, and no other 70707 table carries them. Era's columns stay (ClassBaseCrit, level-independent as Era's are). | wago.tools CASC API | OPEN (unchanged from round 10) | - | - |
| P3 | Melee crit suppression: (target defense - min(skill, level x 5)) x 0.2% plus a flat 1.8% at +3, so 4.8% against a level-63 boss with any weapon skill. Spells and heals take none. This is magey's Era attack table (github.com/magey/classic-warrior/wiki/Attack-table, read 2026-11-02). Forever states no change. | `sim/core/target.go:475-494`, `attack_table_test.go` | Aligned (Era; single source for Forever) | - | - |
| P4 | Crit multipliers: physical (melee and ranged defense types) 2x, magic and heals 1.5x, each scaled by the spell's CritDamageBonus (the talents' aura 108 misc 15) and the attack table's creature-type multiplier. | `sim/core/spell_outcome.go:1080` | Aligned | - | - |
| P5 | Round 10's core crit changes re-read: the stone's per-spell melee crit (`addMeleeWeaponCrit`) reaches melee-defense spells and physical no-defense spells, never ranged or pets; two stones give 4% to both hands, as a weapon-dependent crit does when the other hand also fits the mask; the feral stone; heals taking their school's crit (`healingCritRating`). | `sim/core/consumes.go:1351`, `spell_result.go:277` | Aligned | - | - |
| P6 | Consumable and buff crit against the layered client: Elixir of the Mongoose 17538 (aura 290, 2), Grilled Squid's Well Fed 1249523 (290, 1), Elixir of Sages 1250941 (290, 2), Brilliant Wizard Oil 25113 (57, 1), Firestones 23480-23483 (57: 1, 1, 2, 2), Leader of the Pack and Moonkin Aura (290, 3 after push 113283), Ruthless 20597 (290, 1 after push 113276). Engine amounts and tables match. | `py/critdiff.py` (every crit aura that differs between shipped and pushed: 6 rows, all checked) | Aligned | - | - |

## Before/after

Engine lane/r11-crit (head e183994e8) with site wt/r11-crit, against `$R/sb-before` from `$R/site-before/sim`. Headline metric: DPS, threat per second for Protection, effective HPS for Holy. 10,000 iterations.

| spec | alliance s1 | alliance s50001 | horde s1 | horde s50001 | horde s1 absolute |
|---|---|---|---|---|---|
| paladin-holy 60 raid (eff. HPS) | 0.00% | 0.00% | 0.00% | 0.00% | 641.4 -> 641.4 |
| paladin-retribution 60 raid | +0.22% | +0.22% | +0.22% | +0.20% | 685.7 -> 687.2 |
| paladin-protection 60 raid (threat/s) | +0.81% | +0.81% | +0.80% | +0.80% | 1013.5 -> 1021.6 |

Bands 20 to 50, published bare preset, horde, seed 1 / seed 50001:

| spec | 20 | 30 | 40 | 50 |
|---|---|---|---|---|
| paladin-holy (eff. HPS) | 0 / 0 (32.2) | 0 / 0 (53.9) | 0 / 0 (63.5) | 0 / 0 (190.5) |
| paladin-retribution | 0 / +0.26% | +0.14% / +0.14% | +0.28% / +0.28% | +0.30% / +0.24% |
| paladin-protection (threat/s) | 0 / 0 | 0 / 0 | 0 / 0 | 0 / 0 |

Holy paladin's healing does not move: its heals had the 6% as a heal bonus and now have it as crit. Retribution and Protection win back about what round 10's C1 took (C1 was -0.08% and -0.8%). The bare Protection bands carry no Conviction. CR2 to CR5 move no published number: no published band wears either trinket, Atiesh or a scope, and simdb is not regenerated here.

## For other lanes

- **Integrator:** E d955cc903 moves paladin-retribution and paladin-protection only (table above) and TestProtection's golden; the site's TestRotationLadder goldens for those two will move at the pin. No other spec moves: CR3 and CR5's engine effects are on items and enchants no published set carries.
- **Items / refresh:** at the owner's refresh, simdb changes exactly Blackhand's Breadth, Eye of the Beast, Atiesh and enchants 2523, 8057, 8720 from this lane, and `SpellEquippedItems` is a new fetched (and push-layered) table. Until the next fetch the ranged-scope rule reads no rows, so the committed simdb is unaffected either way.
- **Items:** three non-crit cousins found in the same sweep, for the items lane to judge.
  - Twelve classic-db items have client equip links that state other stats than Era's (Trueaim Gauntlets 13255 hit 1 to the client's 3, ranged only; Royal Seal of Eldre'Thalas 18467 to 18472 and Counterattack Lodestone 18537 re-itemised). This lane changed only their crit.
  - Enchant 7098 "Hit +2%" is an equip spell restricted to melee weapons (mask 42483) that simdb gives as every-table hit, so it reaches spells.
  - Shard of the Defiler (17142) is in simdb with 15% crit (a GM item; check it has no source).
- **Hybrids:** none needed. The paladin guides already say Conviction adds crit "to every spell, heals included" (holy.md:62, protection.md:84, retribution.md:61). The engine now agrees.
- **Physical:** Hack and Slash (13960) and Weaponmaster (1290261) are aura 4 dummies in Forever, so round 10's text-scoped fixes are not affected by the Conviction reading.

## Still open

- **Base crit columns** (P2): an unbuffed, ungeared level-60 character sheet per class would settle them.
- **Crit suppression in Forever** (P3): Era's 1.8% applies to crit "from auras"; Forever's item crit is a rating (stat 32), which may not be suppressed. A level-60 boss log with known gear would settle it.
- **A weapon-restricted aura 290 on spells.** For Conviction and Thundering Strikes the spell part follows the client's aura and the melee-weapon gate is met by every published set. The engine does not model the gate itself (a paladin or shaman with no melee weapon would keep the crit); no published set is affected.
- **Blood Siphon** (round 10's open item) rolls the physical outcome on the spell table and so also takes the melee crit suppression at +3. Its table is still open; no band-60 spec carries it.
