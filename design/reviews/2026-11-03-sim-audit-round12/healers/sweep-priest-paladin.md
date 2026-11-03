# Round 12 healers: client-row sweep of the priest (Holy, Discipline) and Holy paladin kits

Engine: `scratchpad/r12/healers/engine` at 2606bc05e. Client: `scratchpad/r11/c70753/builds/1.60.1.70753/raw` (70753 with the push layered in). The 70753 raw set has no SpellCategories.csv, so DefenseType and Category come from `scratchpad/r12/itemscrit/wago/SpellCategories.csv` (fetched 2026-11-03, current build). Site data: worktree r12-healers.

Scripts, all in this folder:
- `lib.py`: the loaders.
- `ranks.py <file.go>`: parses every `healRank` row and checks it against the client. It compares SpellLevels SpellLevel/MaxLevel, SpellPower ManaCost (and flags any PowerCostPct), the SpellMisc cast time, and the heal or aura effect's base points, variance, per-level growth and coefficient.
- `tal.py <class> <id>...`: prints a talent's effects, its trait curves, and which kit spells its class masks hit, decoded against each kit spell's SpellClassOptions.
- `talents.py <spec> <band> bare horde`: decodes the published band builds.

Builds decoded from `bis/<spec>.json`:
- Holy priest 60: `025003031330000000-03505003030121431`.
- Discipline 60: `025005031330011530-03505003030000000`.
- Holy paladin 60: `50320215012101251-50322-0503`.
- Bands 20 to 50 as printed by talents.py. Holy priest's bis rows still carry gnome/troll; round 11 moved the guide to human/undead.

Every BaseLevel equals SpellLevel on every kit row, so growth from SpellLevel is the client's growth.

## 1. Rows checked

E = engine, C = client. Match: Y (yes), N (no), O (open, carried from an earlier round).

### Priest ranks (ranks.py, every row of heal_ranks.go and binding_heal.go)

| spell | ids | fields | E | C | match |
|---|---|---|---|---|---|
| Lesser Heal 1-3 | 2050, 2052, 2053 | level, mana, cast, coef, base, variance, per-level, max level | heal_ranks.go:37-44 | SpellLevels / SpellPower / SpellMisc / SpellEffect 681298, 681300, 681320 | Y |
| Heal 1-4 | 2054, 2055, 6063, 6064 | same | heal_ranks.go:46-55 | 681114, 681115, 683551, 683232 | Y |
| Flash Heal 1-7 | 2061 ... 10917 | same | :57-72 | 679669 ... 689213 | Y |
| Greater Heal 1-5 | 2060 ... 25314 | same | :74-85 | 679668 ... 701697 | Y |
| Desperate Prayer 1-7 | 13908 ... 19243 | same | :87-102 | 691280 ... 695985 | Y (dwarf only) |
| Prayer of Healing 1-5 | 596 ... 25316 | same; targets 132/31, radius 23 (40 yd); text "the target and their party" | :104-115, direct_heals.go:58-68 | 679131 ... 702104; tooltip 10961 | Y |
| Renew 1-10 | 139 ... 25315 | same; aura 8, 3 s period, 15 s, so 5 ticks; no Attributes_5 0x200, so no tick on application; PERIODIC_CAN_CRIT is in periodic_crit_auto_gen | :117-138, renew.go:10-15 | 678932 ... 702021; SpellMisc DurationIndex 8; Attributes_8 0x1000200 | Y |
| Power Word: Shield 1-10 | 17 ... 10901 | same; aura 69; 0.1; category 56, 4 s category recovery; 30 s | :140-161, power_word_shield.go:13-14 | 687563 ... 688352; SpellCooldowns CategoryRecoveryTime 4000 | Y |
| Holy Nova damage and heal 1-6 | 15237 ... 27805 | same | :163-191 | as printed | Y (not cast) |
| Penance 1-4 | 402174 / 402289 ... 1316995 / 1316991 | cast cost; pulse base, per-level, max level, 0.19 | :193-202 | 1048638, 1239259, 1239260, 1350036 | Y |
| Penance channel | 402277, 1240732, 1240733, 1316992 | 2 s, 1 s period, Attributes_5 0x600 (tick on application) gives 3 pulses; 12 s category 2414; GCD 1500 | penance.go:10-15, 43-46 | SpellMisc, SpellEffect e1 aura 23 | Y |
| Prayer of Mending 1-3 | 401859, 1240826, 1240827 | 172/298/413 at 0.429; 5 charges (e1 bp 5, CumulativeAura 5); 10 s category 1181; holder 30 s; ProcTypeMask 0x800AAAA8 | heal_ranks.go:204-211, prayer_of_mending.go:9-16 | 1047961, 1239406, 1239408; 401877 / 1240848 / 1240849 | Y, but see M3 |
| Binding Heal 1-6 | 401937 ... 1240774 | same as above, both halves; learn 25/32/38/44/50/56 | binding_heal.go:15-28 | 1048084 ... 1239341 (e0 and e1) | Y. Rank 2 self half grows 2.2, not 2.5; round 10 trivia, rank 2 not cast. |
| Inner Fire 1-6 | 588 ... 10952 | mana, armor, level | heal_ranks.go:215-222 | 679452 ... 688399 | Y (not cast) |
| Lightwell 1-3 | 724, 27870, 27871 | 225/295/365 mana; 1.5 s cast; GCD; category 1145, 600 s; 180 s well | lightwell.go:26-54 | SpellPower, SpellMisc, SpellCooldowns | Y |
| Lightwell Renew 1-3 | 7001, 27873, 27874 | 160/233/320 every 2 s for 10 s (5 ticks); no coefficient; no class mask; no tick on application | lightwell.go:32-33, 50-54, 96-134 | SpellEffect 704051 etc.; DurationIndex 10000 | Y. Crit is O (round 10 P3: DefenseType 0 with PERIODIC_CAN_CRIT). |
| Dark Sacrifice 1-5 | 1277324 ... 1277328 | 80 ... 320 health and mana every 3 s for 15 s; no first tick; 1 a level growth, which is 0 at every band (each band casts the rank learned at that level); 600 s; GCD 1500; free | mana_abilities.go:39-43, 91-121; constants_auto_gen.go:729-739 | SpellEffect 1289628 ... 1289637; SpellLevels | Y on amounts and timing. See M2. |
| Divine Grace 1-7 (human) | 1277370 ... 1277378 | base, variance, per-level, max level, 0.429, 600 s, below 50% | racials.go:20-46 | as printed | Y |
| Weakened Soul | 6788 | 15 s | power_word_shield.go:16-17 | DurationIndex 8 (15000) | Y |
| Inner Focus | 14751 | -100% cost (e0), +25 crit (e1), 1 charge, 180 s RecoveryTime, no GCD, no cost, never expires | talents.go:406-484 | SpellEffect e0/e1, SpellAuraOptions 127176, SpellCooldowns 54951 | Y. Scope notes below. |
| Power Infusion | 10060 | 20% of base mana; 15 s; 180 s; no GCD; aura 136 and 79 at 20, misc 126 | power_infusion.go:9-15 | SpellPower PowerCostPct 20, SpellEffect 688356/7 | Y on values. Mask: M1 (unreachable). |

### Priest talents (tal.py; curve value at the taken rank, and mask against the engine's mask)

| talent | id | E value (file:line) | C curve / mask | match |
|---|---|---|---|---|
| Twin Disciplines | 1225132 | 1%/rank, PriestSpellMaskTwinDisciplines (talents.go:139-160, 251; priest.go:96-109) | curve 83143 1-5; e0 hits Penance pulse/channel, PoM heal, Holy Nova, Divine Grace; e1 (op 22) hits Renew | Y. PoM is named only through its heal; the engine heals through the cast. |
| Mental Strength | 18551 | 2%/rank Intellect (healing_talents.go:22; talents.go:35-37) | 83132 2/4/6/8/10, aura 137 stat 8 | Y |
| Mental Agility | 14520 | 3/7/10 (healing_talents.go:45); mask priest.go:116-119 | 83137 -3/-7/-10; mask hits Renew, PW:S, PoM cast, Inner Fire, Smite, Holy Fire, Holy Nova, **Power Infusion** | Y for every cast spell. PI: M1. |
| Inner Focus scope | 14751 | cost: every SpellFlagPriest spell with a cost; crit: not Renew, PureDot or Channeled (talents.go:418-421, 430-455) | e0 names every cast here, PI included; e1 names Lesser/Heal/Flash/Greater/PoH/Binding/PoM cast and heal/Penance pulse and channel/Lightwell, and not Renew or PW:S | Y, except that Penance pulses get no crit (O, round 10 carried) and PI (M1) |
| Improved PW:S | 14748 | 7/14/20 (healing_talents.go:44) | 83140 7/14/20; op 3, mask 0x1 | Y |
| Meditation | 14521 | 0/.17/.33/.5 (talents.go:33) | 83138 17/33/50, aura 134 | Y |
| Soul Warding | 402000 | -15% cost, -4 s cooldown (healing_talents.go:27-28, 96-101) | e0 aura 107 op 11 -4000; e1 op 14 -15; mask 0x1 | Y |
| Renewed Hope | 425280 | 2/rank crit, 1 s/rank cut; mask Flash/Lesser/Heal/Greater/Penance (heals.go:13-14; healing_talents.go:33-34) | 83131 2-10; 83130 -1000 to -5000; aura 4 (script) | Y. Per-pulse cut is O (round 10 P4). |
| Divine Aegis | 431622 | 5/rank, cap 20 × level, 12 s (healing_talents.go:29-32) | 83134 5/10/15; e1 bp 20; 431624 DurationIndex 12000; ProcTypeMask 0x4400 | Y. Cap at bands 40/50: "(20 * 60)" text could be a flat 1,200 (low; O). |
| Improved Renew | 14908 | 5%/rank (healing_talents.go:24) | 83171 5/10/15; op 22, mask 0x40 | Y |
| Holy Specialization | 14889 | 1%/rank on Holy SpellFlagPriest spells and Divine Grace (talents.go:262-275) | 83170 1-5; mask hits Lesser/Heal/Flash/Greater/PoH/Renew/Penance/PoM cast/Binding/Divine Grace; not PoM heal | Y. PoM heal is O (round 9 H2). |
| Divine Fury | 18530 | -100 ms/rank on Smite/Holy Fire/Heal/Greater Heal (healing_talents.go:26, 83-89) | 83167 -100 to -500; mask 0x101480 = those four | Y |
| Inspiration | 14892 | 8/17/25 armor, 15 s (healing_talents.go:36, 47) | 83165 8/17/25; 14893 15 s; ProcTypeMask 0x4400 | Y |
| Improved Healing | 14912 | -5%/rank on Lesser/Heal/Greater/Penance/PoM (healing_talents.go:25, 68-74) | 83163 -5/-10/-15; mask hits exactly those | Y |
| Holy Binding | 1328569 | 50%, Flash Heal only, not self (healing_talents.go:39-40, 250-275) | bp 50, ProcTypeMask 0x4000 | Y |
| Litany of Light | 1317006 | 5%/rank of base cost (healing_talents.go:35) | 83161 5/10 | Y. Spell-code compare is O (round 10 P6). |
| Spiritual Guidance | 14901 | healing 5%/rank Spirit; damage 1/3/5/6/8 (healing_talents.go:19, 23; talents.go:71-79) | 87617 5-25 (aura 175); 83160 1/3/5/6/8 (aura 174) | Y |
| Spiritual Healing | 14898 | 3/7/10 on PriestSpellMaskSpiritualHealing, healing proc mask (healing_talents.go:46, 90-95) | 83159 / 87616 3/7/10; e0 heals incl. PoM cast; e1 (op 22) Renew | Y (PoM heal O, H2) |

### Paladin ranks and spells

| spell | ids | fields | E | C | match |
|---|---|---|---|---|---|
| Holy Light 1-9 | 635 ... 25292 | level, mana, base, variance, per-level, max level; 2500 ms; 0.714; GCD 1500; DT 1 | holy_light.go:9-29 | SpellEffect (one heal effect each), SpellLevels, SpellPower, SpellMisc | Y |
| Flash of Light 1-6 | 19750 ... 19943 | same; 1500 ms; 0.429 | flash_of_light.go:9-25 | as above | Y |
| Holy Shock cast 1-4 | 1311606, 20473, 20929, 20930 | 160/225/275/325 mana; 10 s category 892; GCD 1500; learn 30/40/48/56 | holy_shock.go:33-42, 58-62 | SpellPower, SpellCooldowns | Y |
| Holy Shock heal 1-4 | 1311605, 25914, 25913, 25903 | 114/156/230/320; variance .0755/.0755/.0759/.0789; no growth; 0.429; DT 1 | holy_shock.go:34-41, 60-61 | SpellEffect e0 effect 10 | Y |
| Holy Shock damage 1-4 | 1311604, 25912, 25911, 25902 | 134/182/258/348 | holy_shock.go:50-56 | SpellEffect e0 effect 2 | Y |
| Divine Favor | 20216 | +100 crit, 1 charge, mask HL/FoL/HS; 4% base mana; 120 s RecoveryTime, no GCD; cooldown starts on spend | divine_favor.go:21-76 | aura 107 op 7 bp 100, mask 0x80206000; SpellPower PowerCostPct 4; StartRecovery 0 | Y |
| Lay on Hands 1-3 | 633, 2800, 10310 | learn 10/30/50; 0/250/550 mana back; 20 min category; GCD 1500 | lay_on_hands.go:12-55 | SpellEffect e1 effect 30; SpellCooldowns 1200000 | Y |
| Oath of Light | 1328548 | 6.6667/1.9167/3.8333 × level; 8% base mana; 5 min; GCD | oath_of_light.go:23-30 | e0-e2 per-level; PowerCostPct 8; DurationIndex 300000 | Y. The coefficient scaling is O (round 7/10 P7). |
| Light's Vigil 1-3 | 1310911 / 1310912 ... 1311595 / 1311596 | 730/1000/1340 mana; 1.5 s; 6 s; heal 324/487/704, .0562, 1.2/1.5/1.8, 0.143 | lights_vigil.go:22-36 | as printed | Y (not cast) |
| Illumination buff | 20272 | +50% casting regen for 15 s | healing_talents.go:32-34 | e1 aura 134 bp 50; 15000 | Y |
| Infusion of Light buff | 437063 | -500 ms/rank on Holy Light, 15 s, spent by Holy Light | healing_talents.go:25-27, 84-124 | aura 107 op 10 -1000, mask 0x80000000, 1 charge, 15 s | Y |

### Paladin talents

| talent | id | E value | C curve / mask | match |
|---|---|---|---|---|
| Spiritual Focus | 20205 | 14%/rank pushback on HL/FoL/LV cast (healing_talents.go:20, 77-82) | 82976 14-70; op 9, mask 0xC0000000 / word3 0x4000000 | Y |
| Healing Light | 20237 | 4%/rank on HL/FoL/HS heal (healing_talents.go:14) | 82974 4/8/12; mask 0x6000 / 0x10000 | Y |
| Reverence | 1310899 | 15%/rank (talents.go:213, 298-299) | 110321 15/30, aura 134 | Y |
| Illumination | 20210 | 20%/rank, 50% of the flat cost (healing_talents.go:35-36) | 82971 20-100; e1 curve 85182 10-50; e2 50; **e3 aura 4 bp 10** | Y at 5/5. e3: M4 (open). Holy Shock rank 1 refund is O (round 10 P3). |
| Divine Precision | 1310904 | no heal effect | e0 hit 4/8; e1 silence | Y |
| Infusion of Light | 426065 | -500 ms/rank (healing_talents.go:27) | 110323 -500/-1000; ProcTypeMask 0x14000 | Y |
| Boundless Light | 1328550 | 15%/rank, off the sworn ally (healing_talents.go:50) | 127943 15/30 | Y |
| Holy Power | 5923 | +0.10/rank CritDamageBonus on HL/FoL/HS heal/LV heal (healing_talents.go:44, 185-195) | 82962 10-50, aura 50 misc 2 (Holy) | Y on values; the crit size is O (round 9/10 P1) |
| Divine Intellect | 20257 | 2%/rank (talents.go:26) | 82973 2-10 | Y |
| Benediction | 20101 | 2%/rank on Holy Shock only (talents.go:56-67; holy_shock.go:109) | 82990 -2 to -10; mask includes HS cast, not HL/FoL/LoH; DF and Oath have no mask | Y |
| Conviction | 20117 | 2%/rank to stats.Crit (talents.go:79-83) | 82985 2/4/6, aura 290 | Y |

### Cross-cutting

| check | E | C | match |
|---|---|---|---|
| Sub-20 coefficient exemption on every heal path | spell.go:803-811 exempts any config whose ProcMask matches ProcMaskSpellHealing | (round 11 C1) | Y. Every priest heal (direct, Renew, PoH, Penance, PoM, Binding, Lesser Heal r3, Heal r1, Renew r1 at bands 20-50) goes through healSpellConfig, ProcMaskSpellHealing (heals.go:63). Divine Grace and Contingency use racials.go:92/162/175. Paladin Holy Light r1 (bands 20-49) uses healing.go:68. Holy Shock's heal has no cost or GCD, so it is not a cast rank. PW:S reads entry.coefficient directly (power_word_shield.go:61), never BonusCoefficient. Lightwell Renew, Dark Sacrifice and Lay on Hands carry no coefficient. |
| Heal crit | spell_result.go:281-283: spell table + BonusCritRating + school bonus; ticks spell_outcome.go:78-86; multiplier 1.5 (spell_outcome.go:1082-1083) | - | Y (the Holy Power size is O) |
| HoT tick crit flags | dot.go:437-441, periodic_crit_auto_gen.go | Attributes_8 0x200 on 139 ... 25315, 7001/27873/27874; not on Dark Sacrifice (0x1000) | Y |
| Band rank resolution | sim/request/rotation_ranks.go:183-195, spellranks.json (70753) | trainables / SkillLineAbility learn levels | Y. Every chain level equals SpellLevels. The Renew NPC copies 425268... are not chosen (sameBandOrFirst). |
| Cooldown categories shared across ranks | PW:S, Penance, PoM, Lightwell, Holy Shock: one timer each | categories 56, 2414, 1181, 1145, 892 | Y. Dark Sacrifice: M2. |

## 2. Mismatches

No mismatch found reaches a published spec at any band. The client-row sweep matches the engine on every heal, HoT, shield, mana tool, rank and talent value that the three lists cast or the published builds take. What differs:

**M1. Power Infusion: Mental Agility does not reach it, and casting it does not spend Inner Focus.** Unreachable.
- Engine: `sim/priest/power_infusion.go:43-47`. No ClassSpellMask, so PriestSpellMaskMentalAgility (`priest.go:116-119`) cannot name it. The flags include `SpellFlagNoOnCastComplete`, so the Inner Focus OnCastComplete (`talents.go:452-458`) never fires on it. Inner Focus's cost cut still applies (`talents.go:435-437`, SpellFlagPriest with a cost).
- Client: SpellClassOptions 10060 mask_0 0x80000000. Mental Agility SpellEffect 687728 (14520 e0, op 14, -10) has mask_0 0x8251c1eb, and 0x80000000 is set. Inner Focus 14751 e0 (op 14, -100) mask_0 0xd4f7fffb also has 0x80000000, with ProcCharges 1.
- So the client prices PI at 20% × 0.9 of base mana, and a PI cast would use up Inner Focus.
- Reach: none. No published Discipline build takes Power Infusion (the decoded strings at 20 to 60), and Holy cannot.
- Confidence: high on the rows, zero on impact.

**M2. Dark Sacrifice: each rank has its own cooldown timer; the client shares category 1144.** Unreachable, plus two trivia.
- Engine: `sim/priest/mana_abilities.go:95, 103-106`. `darkSacrificeConfig` makes a new `priest.NewTimer()` for every rank.
- Client: SpellCategories 1277324 and 1277328 both have Category 1144, with CategoryRecoveryTime 600000 on every rank (SpellCooldowns 98122 ... 98126).
- Reach: none. A sim casts only the one rank its band resolves to.
- Trivia, no effect:
  - The engine spell has no SpellSchool (Physical). The client SchoolMask is 32 (Shadow).
  - The engine adds Spirit/5 to each tick from the Spirit at that tick. The text "${$o2+$SPI}" reads Spirit once.
- Confidence: high on the rows.

**M3. Prayer of Mending's holder aura has a variance; the engine heals a flat amount.**
- Engine: `sim/priest/prayer_of_mending.go:89` heals `pom.entry.effect.Center(...)`, and heal_ranks.go:204-211 has Variance 0.
- Client: SpellEffect 1047993 (401877 e0, aura 4) has Variance 0.10935857. The rank 2 and 3 holders 1240848 and 1240849 also have e0 Variance 0.1093, and both have bp 1902. The cast effect 1047961 / 1239408 has no variance. The heals 401863 / 1313379 / 1313381 have bp 1 (the amount is scripted).
- If the server rolls the charge from the holder's variance, a jump heals ±5.5% around the same mean.
- Reach: Holy at 40, 50 and 60 (PoM 401859 / 1240826 / 1240827). The mean is unchanged, so only the overheal distribution moves.
- Confidence: low. It is a dummy aura's field, and nothing says the script reads it. A log of PoM jump amounts at one healing power settles it.

**M4. Illumination's fourth effect is not accounted for.**
- Client: 20210 has e3 aura 4 (dummy) bp 10, besides e0 (proc, curve 82971 20-100), e1 (bp 50, curve 85182 10-50) and e2 (bp 50).
- Engine: `sim/paladin/healing_talents.go:29-36, 141-182` models e0 as the chance and the 50% refund only.
- The server tooltip (read 2026-11-03) says only "...a 0% chance to gain Mana equal to 50% of the base cost of the spell". The chance prints 0 because the curve is unresolved. Nothing in it reads a 10.
- Reach: Holy paladin at 30 to 60 (Illumination 5/5), only if e3 is live.
- Confidence: low. It is probably a leftover, like round 10's priest P7 trait curves. A beta log of Illumination energizes settles it; round 11 found that energize events cannot be fetched.

Carried open, re-read with no new evidence (not re-filed):
- Lightwell Renew crit with DefenseType 0 (r10 P3).
- Renewed Hope cut per Penance pulse (r10 P4).
- Sudden Insight on DefenseType-0 casts (r10 P5).
- Litany of Light by spell code (r10 P6).
- Spiritual Healing and Holy Specialization on PoM's heal (r9 H2).
- Inner Focus crit on Penance pulses (r10). The engine excludes channels. The 14751 e1 mask names 1316991 and 1316992 but not the cast 1316995, which spends the charge.
- Divine Aegis cap at bands 40/50: 20 × level, against a possible flat 1,200.
- Holy Power crit size, Oath of Light scaling, and Illumination on Holy Shock rank 1 (paladin r10 P1/P3/P7).
