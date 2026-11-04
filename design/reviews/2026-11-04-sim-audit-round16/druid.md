# Sim-accuracy round 16: Druid

Lane `druid`. Site wt/r16-druid, fork lane/r16-druid, both from int/r16 (site c62d77ab8, fork 9f7f1f6e4). Client 1.60.1.70755 with the late re-layer. Specs: Balance, Feral cat, Feral bear (threat), Restoration. Measured on guide builds (`$R/rt`) and the frozen BiS (`$R/bis-frozen`).

Summary: 2 confirmed bugs, both fixed (the cat clipped its own Rip; Genesis missed Lacerate). One adoption (Read Leyline for Alliance Skyborne Restoration). Round 15's druid changes all hold against the client.

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Cat Form ±20% spread vs the weapon's own spread.** The client's SpellShapeshiftForm rows 1/5/8 (Cat, Bear, Dire Bear) carry DamageVariance 0.4 and CombatRoundTime 1000/2500. `formPawWeapon` (`forms.go:50`) keeps the weapon's DPS and its own min/max ratio at the form speed. Both readings have the same mean, and every cat and bear ability that uses the paw (Claw 140%, Shred 200%, Ravage 435%, Maul, Primal Bite) is linear in it, so no published average moves either way. Only the hit-to-hit spread differs (The Unstoppable Force 152-255 is ±25%, the form table ±20%). | wago SpellShapeshiftForm 70755 (fetched 2026-11-04); `forms.go:37-64` | OPEN (no number moves; a Forever cat log's white-hit min/max at a known weapon settles it) | - | - |
| S2 | **Unarmed 54.8 DPS fallback.** It is reached only with no main hand. Every published band and preset of druid-feral and druid-feral-bear wears a weapon (checked all 24 rows of the frozen BiS). Blizzard's deep dive covers the equipped weapon only. | bis-frozen druid-feral.json, druid-feral-bear.json `slots.main_hand` | OPEN (no published case; a naked-paw log settles it) | - | - |
| S3 | **Read Leyline for Restoration (Alliance Skyborne).** The racial (1259705, Energized 1270842: +50% mana regeneration for 15 s, 2 min) was modelled and cast by no rotation. A druid runs dry on the raid healer fight (mana lasts about 240-250 s of 300). Lines tried on High Order Skyborne, raid 60: under 70% mana +1.0% (3000 it.); under 90% +1.50 / +1.50%; whenever ready +1.44 / +1.36% (10,000 it., seeds 1 / 50001). Bare 40 +1.8%, bare 60 +2.1%. Adopted as "whenever it is ready" (the guide's Races paragraph already says so; the 90% line is within noise of it). **Race gating checked:** only `Race_RaceHighOrderSkyborne` registers the spell (`racials.go:502-507`); Windshaper (Horde) Skyborne has Skysight. Night Elf (939.3) and Horde Skyborne runs are unchanged with the line, so published numbers do not move. Balance is not mana bound (out of mana 0.4 s of 180): the same line costs it 0.9-1.7%, so Balance gets none. | `rot/resto-ley*.json`, `rot/bal-ley*.json` (lane dir); smoke test allow-list; fork `TestDefaultRotationRunsWithoutWarnings` now runs all three druid races and failed red on the new line before the race table | Adopted (Skyborne only) | - | S 0e4f87a82, E 5741ac22d, E d8023e83e |
| S4 | **Round 15 late-push follow, re-checked every rank** against spellconst 70755 and live Wowhead Forever tooltips (2026-11-04): Claw 140% + 27/39/57/88/115; Shred 200% + 22/32/44/64/80; Ravage 435% + 42/62/78/98; Rake 27/39/56/76 up front, 20/26/32/42 a tick; Rip 4/5/8/10/14/19 a tick plus 5.5/9/10.5/15.8/22.5/31.9 a point (tooltip 9896: 305 at 1 point, 1071 at 5 = 6 x (19 + 31.9 x 5)); Ferocious Bite 17/27/40/62/87/102 with the client variance, 30/45/74/115/160/181 a point, 0.7..3.3 an energy (tooltip 22829: 855-919 at 5 points, 3.1 an energy); Starfire 113/154/201/277/372/471/533, per level 1.8..3.4. All match the engine. Ravage's tooltip says "plus 343" (98 x 3.5, the old multiplier in a hard-coded tooltip); the engine multiplies the flat bonus by 4.35, the same way Shred's beta-log check confirmed. | spellconst/druid.json; nether.wowhead.com/forever/tooltip/spell/{9850,9830,9867,9904,9896,22829} | Aligned | - | - |
| S5 | Rend and Tear 3% (15% at 5): unchanged; Nature's Reach 2% melee+spell hit a rank in both cat builds (raid `050022/...`, leveling `050022/...`, decoded); bear Furor 1 / Naturalist 5 (`0/5123033202310212551/055`); Ravage needs Prowl and not in front (`ravage.go` ExtraCastCondition); Restoration Improved Regrowth 1 (`05302001/0/555300515311321`). All still in place and legal. | decode of the four guide strings against talents/druid.json | Aligned | - | - |
| S6 | **Cat +20% at raid 60 (round 15) re-checked.** The move is the push follow (S4) plus Rend and Tear and the re-layer's gear; each value is the client's. Breakdown at raid 60: Shred 40%, white 26%, Rip 23%, Windfury 7.5%. Nothing out of line. | breakdown, seed 1 | Aligned | - | - |
| S7 | **Restoration -11.7% at raid 60 (round 15) re-checked.** Round 15's own wave 3 attributes about 8 points of it to the shared low-rank penalty (pending owner). The druid rows it reaches: party Rejuvenation rank 7 (8910, learned 40, 3 levels past the 17-level grace at 60: 85% of its bonus). Sweep of the party rank at raid 60 (3000 it.): r5 880.2, r6 915.4, **r7 939.3**, r8 933.0, r9 915.8, r10 897.0. Rank 7 still wins; the tank's Rejuvenation 10, Regrowth 9, Healing Touch 10 and Tranquility 4 are inside the grace. | `rot/resto-party*.json` | Aligned (penalty parameters pending owner) | - | - |

## Feral cat (re-audit)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | **The cat clipped its own Rip.**<br>**What was wrong:** the list recast Rip at 5 points whenever it had under 6 seconds left. A new Rip replaces the old one (the engine's `Dot.Apply` cancels it, as the game does), so the ticks it had left are lost: 76.97 ticks from 15.90 casts, 4.84 of 6 a cast.<br>**What changed:** Rip at 5 points when it is down or under 2 s (its last tick); Ferocious Bite at 5 points only while Rip has 5 s or more. Sweep (raid 60, 3000 it., seed 1): Rip/Bite thresholds 6/6 872.1, 0/0 874.5, 1/1 877.2, 2/2 880.7, 2/3 883.9, 2/4 885.6, **2/5 886.4**, 2/6 884.6, 3/3 879.7, 4/4 879.4. Rip ticks a cast rise to 6.0.<br>**R9-1 (guide round 9):** a 5-point Rip now lands about 2,270 over its duration and a 5-point Bite about 2,270 too; "far more than a Bite" is no longer true. Rip still comes first: it keeps Rend and Tear on and costs less. Prose and notes changed. | `TestCatLetsRipRunToItsLastTick` (sim/request) red on the old file ("refreshes Rip with 6s left"), green now. Raid 60, 10,000 it.: +1.66 / +1.61% (alliance, seeds 1 / 50001), +1.61 / +1.50% (horde). Bare 20/30 unchanged, bare 40 +0.1%, 50 +0.2%, 60 +0.6% (3000 it.) | CONFIRMED-BUG | old (5a8c0299c, 2026-09-28; round 15 C2 read the line and passed it) | S 2caa8fc90, E c5704945c |
| C2 | **S10 (guide round 4): Claw vs Shred before 24.** The Shred line waits for Ferocious Bite (learned 24); with an ungated Shred, bare 20: 52.3 to 52.7 (alliance) and 51.3 to 51.7 (horde), +0.76% at both seeds (10,000 it.); bare 30 unchanged. Under the 1% bar; the guide's "Claw before Shredding Attacks is full" is as good. | `rot/cat-shredfree.json` | Aligned (under the bar) | - | - |
| C3 | Shifting Power 55% of base mana, 40 energy, 8 s (spellconst 70755 now carries the 8000 ms the hotfix gave), 1 s global; Natural Shapeshifter 10% a rank. Faerie Fire 9907 rank 4, 505 armor. Omen 2 PPM, 10 s proc cooldown. Prowl 9913, Poised. All as earlier rounds read them. | spellconst; `shifting_power.go`, `talents.go:311-402` | Aligned | - | - |
| C4 | Raid BiS: The Unstoppable Force's 1 s stun deals no damage (unmodelled, correct). Set and trinket effects as round 15 C4. | bis-frozen | Aligned | - | - |

## Feral bear (re-audit)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| B1 | **Genesis missed Lacerate's bleed.**<br>**What was wrong:** Genesis ("periodic damage and healing done by your spells and abilities") reached Moonfire, Insect Swarm, Rake, Rip and Pounce but not Lacerate.<br>**What the client says:** Genesis 1223081 effect 0 (SpellEffect 1210794) is aura 108 with mask {0x220252, 0x4000100, 0x200000, 0x40000000}; Lacerate's SpellClassOptions mask is {0, 0x100, 0, 0} on every rank (1328617 to 1235827), inside mask_1.<br>**What changed:** `DruidSpellMaskPeriodicDamage` includes Lacerate. No published build has both (the bear takes no Genesis), so no published number moves. | `TestGenesisRaisesLaceratesBleed` (sim/druid/tank) red (1, want 1.1), green now; TestFeralTank golden unchanged | CONFIRMED-BUG | old (Lacerate became a talent in 70707; the mask predates it) | E b0ac55c84 |
| B2 | **Furor 1 in the bear build.** The bear never shifts in the tank fight, so the point never fires there. In play a bear shifts out to Innervate, battle rez or decurse and the point is a 20% chance of 10 rage on the way back: a real, unmodelled effect, so the adoption rule keeps it. The only modelled alternative is Genesis 1 (now on Lacerate, B1): threat +0.15% at raid 60, chance of death unchanged (both factions, both seeds). Not adopted (under 1%). Round 15's source (Warcraft Tavern Forever feral guide: Furor "usually not needed in endgame group content") does not separate one point from five. | spec-breakdown -tank, 10,000 it. | Aligned (keep) | - | - |
| B3 | APL and costs unchanged since round 15 B2: Maul r7 128, Swipe r5 83, Lacerate r5 15 a tick (15 rage less Shredding Attacks' 3), Primal Bite r3 100% + 77 on 6 s, Demoralizing Roar r5, Frenzied Regeneration, Barkskin, Enrage. | spellconst rows | Aligned | - | - |

## Balance (re-audit)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| L1 | Wrath r8 91 ±11.2%, 1.0 a level, 0.571; Moonfire r10 135 at 0.15 and 60 a 3 s tick at 0.13; Insect Swarm r5 31 a 2 s tick at 0.158; Starfire r6 (S4). Engine tables equal spellconst 70755 rank by rank. | `wrath.go:25-34`, `moonfire.go:19-26`, `insect_swarm.go:20-26` | Aligned | - | - |
| L2 | Read Leyline line: loses (S3). The guide's "press it when your mana starts to run low" is the right advice for a Balance druid. | S3 | Aligned | - | - |
| L3 | Raid BiS trinket 2 is Uther's Strength (11302): 6 spell power and 18 healing, its absorb proc does nothing for damage. A ranker pick, for the BiS refresh. | bis-frozen druid-balance.json | For other lanes | - | - |

## Restoration (re-audit)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | Read Leyline line adopted for Alliance Skyborne (S3). | S3 | Adopted | - | S 0e4f87a82 |
| R2 | Heal rows: Healing Touch r10 1920 ±16.5%, 6.6 a level, 1.0; Regrowth r9 1021 at 0.286 plus 142 a 3 s tick at 0.071; Rejuvenation r10 161 and r7 71 a tick at 0.2; Tranquility r4 285 a 2 s tick at 0.067; Swiftmend 20% base mana, 15 s; Innervate 6 min. All equal spellconst 70755. | `healing_tables.go`; spellconst | Aligned | - | - |
| R3 | Party Rejuvenation rank sweep at raid 60: S7. The guide's Ranks paragraph told players to cast rank 8 or 9 on the party; the rotation casts rank 7 and the sweep agrees (195 mana against the top rank's 335, 85% of its bonus at 60). Prose fixed, with the round-15-stale point count (40 Restoration, 11 Balance) and the Improved Moonfire sentence (now Improved Regrowth). | `restoration.md` Talents and Ranks paragraphs | Guide text (not a sim bug) | - | S 571428b11 |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Druid fork tests need `-tags=with_db` (the item database is behind that build tag); without it every package that equips an item panics "No item with id". Not a code fault. | `sim/core/database_load.go:1` | - | - | - |

## Before/after

Raid 60, 10,000 iterations, guide builds, frozen BiS. Before = `$R/before.tsv` (identical to this lane's own start run). After = fork b0ac55c84 + c5704945c on site 2caa8fc90.

| Spec | Faction | Seed | Before | After | % |
|---|---|---|---|---|---|
| druid-balance | alliance | 1 | 652.8 | 652.8 | +0.00% |
| druid-balance | alliance | 50001 | 653.5 | 653.5 | +0.00% |
| druid-balance | horde | 1 | 646.9 | 646.9 | +0.00% |
| druid-balance | horde | 50001 | 647.2 | 647.2 | +0.00% |
| druid-feral | alliance | 1 | 871.6 | 886.1 | +1.66% |
| druid-feral | alliance | 50001 | 871.6 | 885.6 | +1.61% |
| druid-feral | horde | 1 | 856.8 | 870.6 | +1.61% |
| druid-feral | horde | 50001 | 857.6 | 870.5 | +1.50% |
| druid-feral-bear | alliance | 1 | 1200.3 | 1200.3 | +0.00% |
| druid-feral-bear | alliance | 50001 | 1198.9 | 1198.9 | +0.00% |
| druid-feral-bear | horde | 1 | 1194.2 | 1194.2 | +0.00% |
| druid-feral-bear | horde | 50001 | 1195.0 | 1195.0 | +0.00% |
| druid-restoration | alliance | 1 | 939.5 | 939.5 | +0.00% |
| druid-restoration | alliance | 50001 | 939.5 | 939.5 | +0.00% |
| druid-restoration | horde | 1 | 935.3 | 935.3 | +0.00% |
| druid-restoration | horde | 50001 | 935.6 | 935.6 | +0.00% |

Restoration on High Order Skyborne (not a published race): +1.44 / +1.36% at raid 60. No spec outside the druid moves: the engine change is in `sim/druid`, the site changes are two druid APLs, a druid guide and tests.

## Files to regenerate

- `sim/request/testdata/ladder/druid-feral.golden.md` (TestRotationLadder; the Rip threshold reaches bands 40-60).
- `data/builds/1.60.1.70755/addon-data.json` and `addon/ForeverSixty/Data.lua`: the rendered Priority card reads its notes from addon-data, so the feral Rip/Bite notes and the Restoration Read Leyline line show only after this regen.
- BiS/weights/dps-per-point for druid-feral at the refresh (rotation moved +1.5 to 1.7%).
- Fork goldens: none regenerated.

## Measure at refresh

- druid-feral raid and bare 40-60 weights and picks on the new Rip timing.

## For other lanes

- **BiS lane:** Balance raid BiS trinket Uther's Strength (6 spell power) looks like a thin candidate pool for the trinket slot (L3).
- **Platform lane:** the feral page's Cooldowns card shows blank icons for Innervate and Rebirth (Restoration's card shows them). wow-player: not a text blocker, but a real defect.
- **Process note:** port 4391 is the priest lane's dev server; this lane used 4392.

## For the guide lane

Done here, in feral.md (S 2caa8fc90): the Finishers paragraph and Overview bullet (Rip timing, R9-1), and the round-15 build prose round 15 left stale (lines 55, 60, 65, 67, 81, 89: Nature's Reach, Heart of the Wild 5, Blood Frenzy 2, one leveling point of Natural Shapeshifter). Restoration (S 571428b11): the build's point count, the Improved Regrowth bullet, the party rank (rank 7) and the Mana paragraph's Read Leyline sentence.

wow-player SHIP on the rendered feral and Restoration pages (captured at 1440 from this branch). Its open SHOULDs, outside this change: Rend and Tear shown 3/5 under "15% at 5" in the Key talents card (leveling build), the hit-cap line not counting Nature's Reach, Furor described twice (Weaknesses vs talents), Restoration BiS leveling-era picks, and "the paid Skyborne Heroic Pack" wording on both pages. The Priority cards show the new notes and the Read Leyline line only after addon-data regenerates: check them at integration.

## Still open

- S1, S2 (form spread, unarmed paw): need a Forever log.
- Pending owner, not reopened: low-rank penalty parameters and the Rejuvenation rank 1 tooltip; Ferocious Bite rank 6 (31018) and Healing Touch / Rejuvenation 11 as level-60 book ranks.
- Warcraft Logs blocked (gcloud login); not used.
