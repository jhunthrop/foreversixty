# Warrior Arms + Fury engine vs live client 1.60.1.70888h (read-only audit, 2026-11-12)

Engine: /Users/jh/code/wowsims-forever-sim-melee @ 5f1b39450 (paths below relative to sim/warrior unless noted).
Client: scratchpad/melee/wago/*_70888h.csv (hotfix-layered). Talent rank values: talents/warrior.json rank text (70790; 70790->70888 diff has no warrior row).
Helper scripts: scratchpad/melee/wa/sp.py (dump a spell), ranks.py (all ranks by name).
Builds decoded (tree order = tier/column order of talents/warrior.json; checked equal to file order):
- Arms 60 `13323213032215201/0505/205`: IHS1 Defl3 IRend3 ICharge2 ITM3 IOP2 AM1 DW3 2HWS3 Impale2 Bloodthrill2 SS1 WM5 ISlam2 MS1 | Cruelty5 UW5 | IBloodrage2 IronWill5 (51 pts)
- Arms raid `30303213132215201/05050005012/1`: IHS3 IRend3 ITM3 IOP2 AM1 DW3 Spearing1 2HWS3 Impale2 Bloodthrill2 SS1 WM5 ISlam2 MS1 | Cruelty5 UW5 Enrage5 RS1 IExec2 | IBloodrage1
- Fury `30323013002/050530025122100231/2`: IHS3 IRend3 ICharge2 ITM3 AM1 DW3 Impale2 | Cruelty5 UW5 FP3 Enrage2 DWS5 RS1 IExec2 IBR2 DeathWish1 IWW2 Flurry3 BT1 | IBloodrage2

## Abilities
| item | engine | client (70888h) | verdict |
|---|---|---|---|
| Mortal Strike ranks | client_damage.go:33-39 85/110/135/160 @40/48/54/60; mortal_strike.go:22 cost 30; CD 6 s on primaryTimer | 12294/21551/21552/21553 eff1 121 (normalized) base 85/110/135/160; SpellPower 300; Cat 971 CatRecovery 6000; GCD 1500; no stance row | aligned |
| MS/BT shared CD | warrior.go:180 primaryTimer shared | Category 971 on both 21553 and 23881 | aligned |
| MS normalized weapon | mortal_strike.go:88 MHNormalizedWeaponDamage | Effect 121 | aligned |
| Slam ranks | client_damage.go:51 16/32/43/68/87 @20/30/38/46/54; slam.go ids 1240193,1464,8820,11604,11605 | SpellEffect eff 17 base 16/32/43/68/87; SpellLevels 20/30/38/46/54 | aligned |
| Slam cost/cast/CD/GCD | 15 rage, 1.5 s cast, 15 s cat CD, 1.5 GCD | SpellPower 150, CastTimes idx16 1500, CatRecovery 15000, Cat 2412, StartRecovery 1500 | aligned |
| Slam weapon (non-normalized) | slam.go MHWeaponDamage | Effect 17 | aligned |
| Improved Slam 1/2 | slam.go:33 cast/GCD -250/-500 ms, CD -2.5/-5 s, keeps swing at any rank | 12862 aura107 misc10/21/11 on Slam mask 0x200000; rank text 2.5/5.0 s, 0.25/0.50 s; override spells 1310196-1310200 identical numbers + Attr2 0x20000 | aligned |
| Heroic Strike ranks | HS 11/21/32/44/58/80/111/138 @1..56, rank 9 (157 @60) held as AQ book | eff 17 base 11..157; cost 150 every rank; no GCD | aligned (rank 9 = settled book ruling, round 3 P1 / round 6) |
| HS queue / OH hit | no HS-queue OH hit bonus in core | Blizzard 1.60 notes: "Queueing Heroic Strike will no longer increase off-hand hit chance" | aligned |
| HS on-next-swing gives no swing rage | rage.go:150 only pure MH/OH auto | - | aligned |
| Rend ranks | RendBaseDamage 5/7/9/11/14/18/21 per 3 s; rend.go:15 ticks 3/4/5/6/7/7/7 | eff 3 aura3 period 3000 base 5..21; tooltips 15/9s,28/12s,45/15s,66/18s,98/21s,126/21s,147/21s; 11574 dur idx86 21000 | aligned |
| Rend cost/stance/AP | 10 rage, Battle|Defensive, no AP coeff | SpellPower 100; Shapeshift 196608 (Battle+Def); BonusCoefficient 0, FromAP 0 | aligned |
| Improved Rend | talents.go:150 x1.12/1.23/1.35 | 12286 aura108 misc22 mask 0x20; rank text 12/23/35% | aligned |
| Overpower ranks | +5/15/25/35 normalized @12/28/44/60, 5 rage, 5 s cat CD, Battle | 7384/7887/11584/11585 eff 121 base 5..35; cost 50; Cat 65 CatRecovery 5000; Shapeshift 65536 | aligned |
| Overpower window | overpower.go:51 5 s | 1282733 DurationIndex 28 = 5000 | aligned |
| Improved Overpower | overpower.go:84 +25%/rank crit | 12290 aura107 misc7 mask 0x4; text 25/50% | aligned |
| Bloodthrill | bloodthrill.go 10/20%, 5 s, MH incl. HS/Cleave, Rend required | 1289682 proc mask 20; text 10/20%; notes 70755 "10/20%", "Main Hand ... includes Cleave and Heroic Strike" | aligned |
| Execute ranks | ExecuteBaseDamage 125/200/325/450/600; execute.go:12 3/6/9/12/15 per rage | eff 3 base 125..600; rank tooltips "3/6/9/12/15 additional damage" | aligned |
| Execute cost/stance | 15 rage, Battle|Berserker, GCD | SpellPower 150; Shapeshift 327680; StartRecovery 1500 | aligned |
| Improved Execute | talents.go:149 -3/-5 | 20502 aura107 misc14 mask 0x20000000 (-50 tenths max); text 3/5 | aligned |
| Execute stub 20647 below 24 | rank 0 castable only if named | not in spellranks.json chain, so the ladder never names it | aligned |
| Bloodthirst ranks | 30/37/43/48 + 45% AP @40/48/54/60 | 23881..23894 eff2 base 30/37/43/48, eff3 45; notes "45% (was 35%)" | aligned |
| Bloodthirst cost/CD | 30 rage, 6 s | SpellPower 300; CatRecovery 6000 Cat 971 | aligned |
| Whirlwind | whirlwind.go 25 rage, 10 s, Berserker, normalized MH + OH (462891) each target up to 4 | 1680 cost 250, Cat 891 10000, Shapeshift 262144, eff 121; 462891 OH flag Attr3 0x01000000, eff 121; tooltip "up to 4 enemies ... both melee weapons" | aligned |
| Improved Whirlwind | -1/-2 s, +15/30% primary target | 1310315 aura107 misc11 -2000 on WW mask_1 0x4; text 15/30% | aligned |
| Spearing Strike | 15 rage, 20 s, Battle, 75% normalized, +75% vs Giant/Dragonkin | 1310222 cost 150, RecoveryTime 20000, Shapeshift 65536, eff121 + eff31 75; notes "requires Battle Stance", no 2H requirement | aligned (mounted clause unmodelled, no effect on bosses) |
| Reckless Swings | 15 s CD, 1 s GCD, 4 charges, 8 s, +5 rage per landed white swing, any stance | 1329555 Recovery 15000, StartRecovery 1000, ProcCharges 4, ProcTypeMask 4, dur 8000, trigger 1329554 energize 50; no Shapeshift row | aligned |
| Hamstring ranks | 5/18/45 @8/32/54, 10 rage, Battle|Berserker | 1715/7372/7373 eff2 5/18/45; cost 100; Shapeshift 327680 | aligned |
| Charge | 9/12/15 rage @4/26/46, 15 s, Battle, pre-pull only, no GCD | 100/6178/11578 energize 90/120/150; CatRecovery 15000; Shapeshift 65536; no StartRecovery | aligned |
| Improved Charge | +3/+6 | text 3/6 | aligned |
| Bloodrage | 10 now + 1/s x10, 60 s, no GCD, any stance, level 10 | 2687 energize 100, Recovery 60000, no StartRecovery; 29131 aura24 period 1000 base 10 dur 10000 | aligned |
| Improved Bloodrage | x1.25/x1.5 all | 12301 aura108 misc8 mask 0x100 (+50 max); text 25/50% | aligned |
| Berserker Rage | level 30, 30 s, GCD, Berserker, IBR 5/10 instant | 18499 BaseLevel 30, Recovery 30000, StartRecovery 1500, Shapeshift 262144; 20500 text 5/10 | aligned |
| Recklessness | level 50, 30 min, 15 s, +100% crit, +20% taken, Berserker, GCD | 1719 Recovery 1800000, dur 15000, aura290 100, aura87 20, Shapeshift 262144 | aligned |
| Death Wish | 10 rage, 3 min, 30 s, +20% physical, +5% taken, GCD, any stance | 12328 cost 100, Recovery 180000, dur 30000, aura79 misc1 20, aura87 5 | aligned |
| Battle Shout ranks | core/buff_ranks.go:248 9/21/33/51/78/111 +0.3/0.6 per level to MaxLevel 11/21/31/41/51/61, floor; rank 7 book | 6673..11551 aura99 base + RealPointsPerLevel, SpellLevels MaxLevel; 25289 book | aligned (at 60: 115 AP) |
| Battle Shout cost/duration/GCD | 10 rage, 3 min, GCD | cost 100, dur idx25 180000, StartRecovery 1500 | aligned |
| Sweeping Strikes | registered, Battle, 30 rage | 12292 cost 300, Recovery 30000, Shapeshift 65536 | aligned (never cast single target) |

## Stances and baseline
| item | engine | client | verdict |
|---|---|---|---|
| Battle Stance | stances.go: threat x0.8 | 21156 aura10 -20 | aligned |
| Berserker Stance | +3% crit (all tables), threat x0.8, +10% taken | 7381 aura290 3, aura10 -20, aura87 10 | aligned |
| Stance CD / GCD | 1 s shared, no GCD | Cat 47 CatRecovery 1000, no StartRecovery | aligned |
| Stance levels | 1/10/30 | 2457 L1, 2458 L30 | aligned |
| Tactical Mastery | 10 rage from 14 | 1310185 L14, base 10 | aligned |
| Improved Tactical Mastery | +5/rank | 12295 base 15 max; text 5/10/15 | aligned |
| Anger Management | 1 rage / 3 s | text "1 Rage every 3 sec"; 12296 e0 aura85 misc1 base 15 | aligned to text (question Q2) |
| Normalized rage | core/rage.go:96-99 1H 3.46/speed, 2H 4.50, OH x0.5, crit x2 | notes: crit +100%; 1H/crit log-confirmed (physical-logs L1) | 1H aligned; 2H/OH open (known, unlogged) |
| Unbridled Wrath | 12%/rank, 1 rage, white swings only, landed | 12322 ProcTypeMask 4, trigger 12964 energize 10; text 12..60% | aligned |
| Dual Wield Spec | +5% OH dmg/rank, +10% OH rage/rank | 23584 aura122 base 5; text 5..25% / 10..50% | aligned |
| Furious Precision | OH hit 4/7/10 | 1323963 aura54 OH-weapon restricted; text 4/7/10 | aligned |
| Cruelty | +1% melee crit/rank | 12320 aura52 base 5 | aligned |
| Impale | +10%/rank of the crit bonus (2.2x) on every warrior ability incl. Rend | 16493 aura108 misc15 mask 0xEE604CEE/0x8140/0x1/0x10000400; covers HS, Cleave, Rend, OP, Hamstring, MS, BT, Slam, Execute, WW (+462891), Spearing | aligned |
| Two-Handed Weapon Spec | physical x1.02/1.04/1.06 when MH is 2H | 12163 aura79 misc1, 2H equip mask; text 2/4/6; notes 70755 | aligned |
| Weaponmaster | axe/polearm +1% crit, mace/staff 3% armor ignore, sword 1% extra attack per rank | 1290261 e0 5 / e1 15 / e2 5; text per rank | aligned (sword ICD 200 ms unconfirmed, no client value) |
| Deep Wounds | 20%/rank avg weapon dmg over 12 s (4 ticks), rolling, no AP | text 20/40/60% over 12 s; Blizzard #234 + notes "does not scale with AP, rolling" | aligned |
| Flurry | 10/20/30% speed, 3 charges, white consumes | 12319 trigger mask 87380; 12966 charges 3, mask 4, aura319 30 | aligned |
| Enrage | 30% on damage taken, +2%/rank physical, 12 s | 12317 ProcChance 30, taken mask; text 2..10% | aligned (no procs in a raid DPS seat) |
| Death Wish talent | see above | | aligned |
| Improved Heroic Strike | -1 rage/rank | 12282 mask 0x40; text 1/2/3 | aligned |
| Improved Berserker Rage | 5/10 rage | text | aligned |
| Improved Execute | see above | | aligned |
| Booming Voice, Lingering Rage, Gore Drinker, Imp Hamstring, Iron Will, Deflection | not taken or no DPS effect | | aligned |
| Talent string -> proto | proto/warrior.proto field order = tree order; core.FillTalentsProto | talents/warrior.json tier/column order | aligned |
| GCDs (MS, BT, Slam, OP, Rend, Exec, Ham, WW, Spearing, BShout, DW, Reck, BRage = 1.5; RS 1.0; Bloodrage, Charge, stances, HS none) | per files | SpellCooldowns StartRecoveryTime | aligned |
| Blizzard notes, warrior lines (2360696, 2403040; 2404092 has none) | crit rage +100%, UW flat 1, DWS no hit, FP, Booming Voice, BT 45%, Berserker Rage 30, Slam 15 s, ITM 3 pts, 2HWS 2/4/6, Bloodthrill 10/20 MH incl HS, ISlam 2.5/5, Flurry 3 pts, IWW, RS, DW rolling no AP | | aligned |
| Set bonuses | T1 2102/2103 via client set model; no set piece in any Arms/Fury BiS slot (only 1-piece hits of sets 161/163/1618/1619) | | not exercised by BiS |

## Findings
B1 (BUG, small): rage refund ignores cost discounts. core/rage.go:362 `Refund: options.Refund * options.Cost` fixes the refund at 80% of the BASE cost at registration; IssueRefund (rage.go:387) pays that fixed amount. The energy bar (energy.go:513) refunds `ec.Refund*spell.CurCast.Cost` (80% of what was paid). With the builds' discounts: Heroic Strike with Improved Heroic Strike 3 (Fury, Arms raid; 12282 -3) refunds 12 of 12 paid (should be 9.6); with IHS 1 (Arms 60) 12 of 14 (should be 11.2); Execute with Improved Execute 2 (Fury, Arms raid; 20502 -5) refunds 12 of 10 paid, a net rage GAIN on a dodge/parry/miss (should be 8). Gnome Eureka! (10% percent cut) is also ignored. Size: Fury's raid breakdown shows "Refund 78" rage a fight; the excess is a few rage a fight (order 0.1-0.3% DPS), larger for low-hit bare bands.

Q1: Execute on a miss/dodge/parry spends all extra rage (execute.go:97 SpendRage before the outcome roll). Whether Forever's server keeps the extra rage on a failed Execute is server-side (the client only shows eff 3 dummy + eff 64 trigger 26651 self-target 1 s dummy). Needs a log.
Q2: Anger Management's working effect is aura 85 (mod power regen, rage) base 15 tenths; the tooltip says 1 rage per 3 s. If Forever's server keeps the classic "amount x 3/5 per tick" handling, 15 gives 0.9 rage per 3 s, not 1.0 (vanilla data was 17 -> 1.02). About 6 rage a 3-minute fight. Needs a log (AM tick size).
Q3: 2H normalized rage 4.50/speed and OH factor 0.5 remain unlogged (already open in physical-logs).
Q4: Reckless Swings and Unbridled Wrath require a landed swing (engine); the client proc mask (4) gives no hit mask, the server default (hit or crit) is assumed.
Q5: Bloodthrill can be procced by Rend's own landed cast and by Overpower itself (both MH specials on a Rend'd target); the client mask 20 includes "done melee ability", so plausible, unverified.
Q6: Weaponmaster sword extra-attack ICD 200 ms (no client value). Arms raid wields Blackblade of Shahram (2H sword), so it is live at raid 60.

Count: 64 rows aligned (tables above), 1 bug, 6 questions.
