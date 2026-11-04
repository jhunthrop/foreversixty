# Sim-accuracy round 19: paladin

Lane paladin. Site wt/r19-paladin from int/r19 (f91e9d3f9), fork lane/r19-paladin from int/r19 (6ea675594, no engine change this round). Client 1.60.1.70755. Every sim went through `$R/sim.sh`. Measured on frozen BiS (`$R/bis-frozen`) and guide builds (`$R/rt`). Leveling bands use the band's own published build and the bare preset; healers use `-heal` on their band's heal profile.

**Result:** 2 CONFIRMED rank-row fixes, both measured level by level (H1: Holy at 40-49; R1: Retribution Consecration at 54-59). No engine bug. Every other seed item is Aligned or OPEN with what settles it.

**How a rank was judged:** I followed round 18 priest. I built a binary with no paladin rank rows (`sb-norows`, embedded copy edited and restored), then forced each candidate rank by authoring that rank's own id. For damage specs, every reference was swapped, conditions included. Holy's lines were swapped one by one, because the tank and group Flash of Light lines share ids across ranks. Then:
- coordinate descent at the floor, +3, +6 and +9 (3,000 iterations, seed 1);
- every level from 40 to 49 (Holy) or 50 to 59 (Retribution) at 10,000 iterations, seeds 1 and 50001, both factions;
- the adopted rows checked on the real binaries against `$R/sb-before`.

Before I rewrote the swap per line, a first Holy pass collided two families on 19941. Those numbers were thrown away.

## Holy

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | **Holy's 40-49 rows were wrong at several levels.** Under the MaxLevel + 12 rule, these ranks win over the old rows at both seeds and both factions:<br>- **Holy Light rank 4 at 40-42**, over the top rank: +0.5 to +0.9%. At 43 it splits by faction, and from 44 it ties.<br>- **Group Flash of Light rank 2 at 44-46**, over rank 1: +0.35 to +1.0%. Rank 1 (MaxLevel 25) keeps 15 to 25% of the bonus there; rank 2 (MaxLevel 31) keeps 85 to 100%. At 40-43 and 47-49 rank 2 or 3 is within +0.25%, so rank 1 stays.<br>- **Tank Flash of Light rank 3 at 49**, over rank 2: +0.8 to +0.95%. Rank 2 keeps 70% of the bonus at 49, rank 3 all of it. At 47-48 rank 3 loses 0.6 to 2%.<br>- **Holy Shock rank 1 at 40-46**, over rank 2: rank 1 keeps its whole bonus there for 160 mana against 225. Alone it gains +0.2 to +0.56%; with the rows above it adds +0.3 to +0.7% at 40-42 and +0.2 to +0.4% at 44-46. From 47 it loses.<br>**After, on the real binaries (effective HPS, seed 1 / 50001):**<br>- Alliance: 40: 177.2/177.3 to 179.8/179.7 (+1.47/+1.35%); 42: +1.17/+1.23%; 43: +0.24/+0.24%; 44: +0.59/+0.71%; 46: +1.05/+0.99%; 49: 157.5/157.4 to 159.0/158.9 (+0.95/+0.95%).<br>- Horde: 40: +1.22/+1.33%; 42: +0.90/+0.90%; 43: +0.06/+0.18%; 44: +0.75/+0.75%; 46: +1.17/+1.11%; 49: +0.85/+0.78%.<br>- Unchanged: 50, 60.<br>**Mana:** every 40-46 fight still runs dry (the mana lasts 165 to 263 s of 300). At 49 the old rows ended with mana unspent (364 s); the new row spends it (135 s) and heals more. | `scan.log`, `combo.log` and `ba.log` in `$R/paladin`. Tests first: `TestHolyPaladinFlashOfLightAndHolyLightKeepTheirHealingBonus` (49 tank rank 3, Holy Light rank 4 at 40 and 42) and the new `TestHolyPaladinGroupFlashOfLightAndHolyShockBelowFifty` both failed, then passed. `TestRoundTwelveHealerRowsBelowSixty` (45 is now [2]) and the addon pytest (band 40 exports Holy Light 1026 and Holy Shock 1311606) were updated, and both pass. | CONFIRMED | old. Round 16 W3 swept 20-59 at 3,000 iterations, one seed, one faction, on the round-16 engine; rounds 17 and 18 did not re-judge 40-49. | S 65be740e8 |
| H2 | **Holy 50-59 rows hold** (descent at 50, 53, 56 and 59, both factions): Holy Light rank 2, tank Flash of Light rank 4 then 5, group rank 4, and Holy Shock at its top rank all stay best, within +0.05%. | `descent.log` | Aligned | - | - |
| H3 | **Shared watch (priest P1):** Holy Light rank 1 at 47-49 (MaxLevel 5) and rank 2 at 50-59 (MaxLevel 11) keep **0%** of the healing bonus. They win only through Oath of Light's flat bonus, which skips the penalty (pending with the owner). If low ranks also proc Illumination or Infusion of Light less, those two rows are the ones exposed. The code is not changed. | client SpellLevels 70755 | OPEN (watch) | - | - |
| H4 | Talents: Precision 3 and Guardian's Favor 2 fill the Protection points for Holy. Every Holy talent that heals is already taken, and no other legal point in Protection rows 1-2 or Retribution rows 1-2 heals. So no swap gains, and the guide (holy.md:67) says so. | talents/paladin.json 70755 | Aligned | - | - |
| H5 | Oath of Light is pending with the owner and was not reopened. | - | pending | - | - |

## Retribution

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | **Consecration rank 1 at 54-59 was wrong.** Rank 1 (MaxLevel 25) keeps about 15% of its spell damage at 54 and almost none from 56. Rank 4 is the top rank learned and wins:<br>- Undead (horde, seed 1/50001): 54 +0.56/+0.71%, 56 +1.55/+1.70%, 57 +2.05/+2.20%, 59 +1.29/+1.50%.<br>- Human: 54 +0.05/0.00 (a tie), 55 +0.05/+0.05 (a tie), 56 +0.25/+0.25%, 58 +0.69/+0.74%, 59 +1.04/+1.09%.<br>Rank 1 still wins at 50-53: rank 4 loses 0.2 to 2.5% for Human, and Undead is mixed at 53.<br>**Fix:** the 50-59 rank-1 row now ends at 53, and from 54 the rotation casts the top rank learned (a row naming rank 4 is refused as a no-op).<br>**After, on the real binaries:**<br>- Alliance seed 1: 54: 198.9 to 199.0; 56: 201.7 to 202.2; 59: 201.5 to 203.6 (+1.04%).<br>- Horde seed 1: 54: 196.9 to 198.0 (+0.56%); 56: 200.0 to 203.1 (+1.55%); 59: 200.8 to 203.4 (+1.29%).<br>- Seed 50001 agrees. Bands 50 and 60 unchanged.<br>Hammer of Wrath rank 1 against 2 is mixed (±0.3%), so no row. | `dscan-ret.log` and `ba-ret.log`. Tests: `TestRetributionCastsTheReMeasuredConsecrationRanks` (53 rank 1; 54 and 59 rank 4) failed, then passed. `TestChosenRanksReadsTheCuratedRowsAtALevel` was updated (53 [1], 59 none). | CONFIRMED | old. Round 16 W3 had ranks 1 and 4 within 0.5% at 53-59 (one faction, 3,000 iterations); later rounds moved it. | S 2fba11123 |
| R2 | Retribution at 40, 43, 46 and 49: Consecration, Holy Strike, Seal of Righteousness and Hammer of Wrath ranks all hold (descent: no gain over the rows). | `ddescent.log` | Aligned | - | - |
| R3 | **Seal of Righteousness one-hand/two-hand spell damage split.**<br>- The engine uses the client's 0.1 with a one-hander and 0.11 with a two-hander. The 1.1 is "from testing in SoD" (sor.go), and no Forever client row has it.<br>- Forever's tooltip (SpellDescriptionVariables 1023: `$2h=${0.12*$SPH}`, `$1h=${0.0765*$SPH}`) is a hand-written range, fastest one-hander to slowest two-hander. It is not a coefficient table: Round 18 S2 showed that the same variable set prints a term on rank 8, which has a 0 coefficient.<br>- Measured on the published two-hander, Human, raid build, seed 1 / 50001:<br>&nbsp;&nbsp;- tooltip split (0.0765/0.12): raid 60 +0.15/+0.15%, bare 60 +0.64%, bare 40 +0.52/+0.52%;<br>&nbsp;&nbsp;- flat 0.1 (no split): raid 60 -0.15%, bare 60 -0.61%, bare 40 -0.60%.<br>- Web search (2026-11-04) found no Forever measurement. Classic community figures disagree (about 10%/12.5% and 9.2%/10.8% per speed).<br>- **Settles it:** a level-60 Forever log with Seal of Righteousness rank 7 procs at two spell damage values, one with a one-hander and one with a two-hander. No such log exists yet (launch is today). | `ret.log`; variant binaries `sb-sortip` and `sb-sorflat` (built and reverted, nothing committed); wago SpellDescriptionVariables 70755 | OPEN (under 1% either way) | - | - |
| R4 | **Retribution Dwarf race line** (retribution.md:111, "Dwarf matches Human on the raid build in a raid, and Human is ahead solo"). Raid 60, Human against Dwarf, guide raid build, seed 1 / 50001:<br>- raid: 785.0/784.5 against 787.4/786.8 (Dwarf +0.31/+0.29%);<br>- bare 60: 280.2/280.2 against 278.0/278.1 (Human +0.79/+0.75%);<br>- Beast target, raid: Dwarf +5.0% (Big Game Hunter).<br>The line is true: Dwarf is within the 1% race rule in a raid and Human leads solo. | `ret.log` | Aligned | - | - |
| R5 | Round 18's rank 7 from 58 still holds: Seal of Righteousness rank 7 is the descent's best at 59, and every lower rank loses 2 to 22%. The rank 7/8 decision itself is not reopened. | `ddescent.log` | Aligned | - | - |

## Protection

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | **Reckoning against Conviction 3 + Deflection 2** (guide round 15 PALADIN-R15-2), re-run on the round-19 base. Dwarf, alliance, 10,000 iterations, seed 1 / 50001:<br>- raid 60: threat 1151.9/1152.2 against 1126.6/1127.2 (Reckoning -2.2%), damage taken 367.0 against 382.2;<br>- bare 60: threat 484.3/483.9 against 469.0/469.3 (-3.1%), chance of death 26.7% against 39.3%.<br>These are the same numbers as round 18 P1. The guide's measured reason (protection.md:82) is true. | `reck.log` | Aligned | - | - |
| P2 | **Seal of Fury's aura coefficients below 60** (round 18 A2): the aura's effect 0 is not read.<br>- **The rows are a dormant copy of Seal of Righteousness.** Each Seal of Fury rank's aura effect 0 is 0.9 times the next Seal of Righteousness rank's weapon-speed value. Examples: 20163 486.9 = 0.9 x 541, with points per level 27.9 = 0.9 x 31; 20422 1266 = 0.9 x 1407; 20423 1607 = 0.9 x 1786. The coefficient is 0.9 x 0.1 = 0.09, and rank 7's is 0, like Seal of Righteousness rank 8. Rank 4's 0.9 is the same pattern with a slipped digit.<br>- **Era has none of these rows;** they are new in Forever.<br>- **The tooltips name only the proc's flat damage** (`$20415s1` and the like, 6 to 35 plus the proc's 0.1).<br>- **Every guide says the damage does not scale with weapon speed:** gfinityesports.com/article/world-of-warcraft-forever-protection-paladin-guide-build-rotation, mobalytics.gg/wow-forever/guides/paladin-class-overview and kami-labs.fr/en/wow-forever/paladin-wow-forever-deep-dive (all fetched 2026-11-04). If the server read the aura as it does for Seal of Righteousness, the damage would scale with weapon speed.<br>So the engine's flat proc plus 0.1 is right at every rank, rank 4 included. | wago SpellEffect 70755 and 1.15.9.70576; Spell.csv descriptions 70755 | Aligned (closes round 18 A2) | - | - |
| P3 | **Protection rows hold.** Consecration rank 1 at 44-49 against ranks 2 and 3 at 10,000 iterations: ranks 2 and 3 are -1.1 to -2.0% at 44, and at 47-49 ranks 1 to 3 are within ±0.4% (mixed). Holy Strike rank 1 at 40-42, Seal of Fury, Holy Shield and Hammer of Wrath at their top ranks all hold at 40, 43, 49, 50, 53, 56 and 59. | `ddescent.log` and `dscan-prot.log` | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **Integrator lead: a client coefficient of 0 hidden behind a generator default.** No paladin spell has this. Every coefficient the engine uses has a nonzero client row with the same value:<br>- Holy Light 0.714, Flash of Light 0.429, Holy Shock 0.429 (damage and heal), Exorcism 0.429, Holy Wrath 0.19, Hammer of Wrath 0.429;<br>- Holy Strike 0.429 (effect 121), Holy Shield 0.08;<br>- Seal of Command proc 0.29; Judgements of Command, Righteousness and Fury 0.429;<br>- Consecration 0.042 (second damage effect), Light's Vigil 0.143, Seal of Fury proc 0.1, Seal of Righteousness ranks 1-7 0.1.<br>The only 0 is Seal of Righteousness rank 8, which was decided in round 18. | spellconst/paladin.json 70755; engine sim/paladin/*.go | Aligned | - | - |
| X2 | **Unused dummy coefficients.** The Judgement of Righteousness and Judgement of Fury ids carry a second, dummy effect with a coefficient:<br>- Righteousness: 0.058, 0.125, 0.185, then 0.2 on ranks 4-8. In Era these are twice the Seal's old per-rank values, with the Seal's weapon-speed base.<br>- Fury: 0.1125, then 0.18 and 0.2.<br>Like Seal of Fury's aura, they look like the data the seal script reads, not a second damage term on the Judgement. The tooltips print only the Judgement's own damage. Not changed. | wago SpellEffect 70755 and Era | OPEN (watch; a level-60 Judgement of Righteousness log against spell damage settles it) | - | - |
| X3 | Every paladin MaxLevel exception in the 70755 SpellLevels is in the engine's `clientSpellMaxLevel` table: zero missing. Holy Shock's cast ids have MaxLevel 0 and key on learn level + 5 by the engine rule; the heal ids carry their own MaxLevel. | `spell_max_level_auto_gen.go` against SpellLevels.csv | Aligned | - | - |
| X4 | **Client build 1.60.1.70790** appeared on wago at 18:08 UTC. I flagged it to the integrator, who confirmed it is 70755 with the hotfixes baked in. My diff of SpellEffect, SpellLevels, SpellPower, SpellMisc and SpellCooldowns against every paladin spellconst id found no paladin row changed. | wago api/builds; `$R/paladin/c790` | Aligned (no paladin change) | - | - |

## Before/after

- **Raid 60:** unchanged for all three specs. Every change is a rank row below 60, and bare 60 and raid 60 cast as before. Measured: Holy bare 60 255.4/254.5 alliance and 243.0/242.3 horde, both before and after; Retribution bare 60 298.0 and 296.5, both before and after.
- **Holy, leveling (effective HPS):** see H1. +0.06% to +1.47% at 40-49, unchanged at 50-59.
- **Retribution, leveling:** see R1. +0.56% to +1.70% at 54-59 for Undead, a tie to +1.09% for Human; unchanged at 40-53.
- **Protection:** unchanged.

## Files to regenerate

- addon-data, then Data.lua (addonrotation): Holy's new rows (Holy Light 40-42, Flash of Light 44-46 and 49, Holy Shock 40-46) and Retribution's Consecration from 54.
- sim/request ladder golden testdata/ladder/paladin-holy.golden.md (band 40 now casts Holy Light rank 4 and Holy Shock rank 1). paladin-retribution needs it only if its golden samples levels 54-59.
- coach starters, if they quote leveling ranks.

## Measure at refresh

- Nothing new. No BiS-relevant engine change.

## For other lanes

- No shared code was touched, and no commit moves DPS outside paladin.

## For the guide lane

- retribution.md :95 and :133, Consecration ranks: rank 1 again at 50 and rank 4 from about 55. Shipped in S 2fba11123.
  - wow-player, first pass: NO-SHIP (the Rotation paragraph did not name rank 1 at 50; "mid 50s" was vague).
  - Second pass: **SHIP**. Its optional SHOULD was to say once that rank 3 never wins. The data shows rank 3 tying rank 4 for Undead at 54-55, so I left it out.
  - Captures: `$R/paladin/cap/ret-cons2.png` and `ret-leveling2.png`.
- Holy: no guide prose or card note names a 40-49 rank, so there is no text change. The rows reach players through the addon after the regen.

## Still open

- R3: the Seal of Righteousness one-hand/two-hand split (a level-60 log).
- X2: the Judgement dummy coefficients (a level-60 log).
- H3: Holy Light 0%-bonus rows if low ranks proc less, together with Oath of Light (owner).
- Live logs: api.foreversixty.gg/v1/reports/recent still has only the 2026-10-20 report (last looked at 18:37 UTC, before the 23:00 launch). Warcraft Logs is still blocked (gcloud).

## Sources

- wago.tools SpellEffect, SpellLevels, SpellPower, SpellMisc, SpellCooldowns, Spell, SpellDescriptionVariables and SpellXDescriptionVariables, builds 1.60.1.70755, 1.60.1.70790 and 1.15.9.70576, fetched 2026-11-04.
- data/builds/1.60.1.70755: spellconst/paladin.json, spellranks.json, talents/paladin.json.
- Seal of Fury guides (fetched 2026-11-04):
  - https://www.gfinityesports.com/article/world-of-warcraft-forever-protection-paladin-guide-build-rotation
  - https://mobalytics.gg/wow-forever/guides/paladin-class-overview
  - https://kami-labs.fr/en/wow-forever/paladin-wow-forever-deep-dive/
- Seal of Righteousness coefficients (Classic, conflicting):
  - https://us.forums.blizzard.com/en/wow/t/seal-of-righteousness/394206
  - https://warcraft.wiki.gg/wiki/Seal_of_Righteousness
- Lane logs in `$R/paladin` (`$R` = the round-19 scratchpad): descent.log, scan.log, combo.log, ba.log, ddescent.log, dscan-ret.log, dscan-prot.log, ba-ret.log, ret.log, reck.log.
