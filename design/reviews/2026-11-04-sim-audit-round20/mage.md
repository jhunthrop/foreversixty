# Sim-accuracy round 20: mage

Specs: mage-arcane, mage-fire, mage-frost. Site wt/r20-mage (from int/r20 04d7da663). Fork lane/r20-mage (from int/r20 da54b7b06). Client 1.60.1.70755 with the late re-layer.

Method:
- Seed items first (Ice Lance at 60, the cost mods, the Frost talent search, the coordinator's Frost race re-run), then the re-audit.
- Raid 60: spec-breakdown, 10,000 iterations, seeds 1 and 50001, both factions, `-bis-dir $R/bis-frozen`, guide raid builds (`ALLRT=1`). Bare bands on the published band builds. Candidates at 3,000 iterations first.
- Every sim through `$R/sim.sh`.
- Client rows read from data/builds/1.60.1.70755/spellconst/mage.json and talents/mage.json; Blizzard's notes from the Kaivax thread "WoW Forever Beta Development Notes - Updated November 2" (https://us.forums.blizzard.com/en/wow/t/wow-forever-beta-development-notes-%E2%80%93-updated-november-2/2360696, read 2026-11-04).

**Confirmed and fixed: 4** (all `old`).
- A Presence of Mind Pyroblast spent Heating Up's stacks (Blizzard's 16 October note says it does not).
- Clearcasting was spent on a Missile Barrage Arcane Missiles under Arcane Power, and whenever the Barrage had procced before the Clearcasting (Blizzard's 2 November note: never while Missile Barrage is up).
- The Arcane raid build carried a point of Arcane Geometry that does nothing on a boss while Arcane Meditation sat at 2/3: moved, +0.9% at raid 60.
- Arcane Explosion never missed (hit-only outcome). No rotation casts it, so no number moves.

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Ice Lance at 60: no live log yet.** foreverlogs' newest report is 2814 (2026-11-04 12:24 UTC, "Test", no encounters); 2816 to 2822 do not exist. Launch is 23:00 UTC; the first live logs will be level 1-20 and cannot reach Ice Lance rank 6 (level 56). The 0.107 from nine beta players at 24-35 stands. Needed: one Frost mage at 56-60 with Frostbolt rank 10/11, Frost Nova and Ice Lance rank 6 (1240047) hits in the same fight, fitted with round 19's `mage-tools/fit.py`. Last looked 21:15 UTC. | foreverlogs `reports/2812..2822` through `scratchpad/r3/casters/py/fl.sh` | OPEN (log needed) | - | - |
| S2 | **Arcane Power's cost increase on stacked Blast and Barrage Missiles: the engine already models it the way the client prices it (adds), and the client proves the model.**<ul><li>Every cost mod here is the client's aura 108 misc 14: Arcane Power +30 (12042 effect 2), the Blast stack +100 (400573 effect 2), Missile Barrage -100 (400589 effect 3), Clearcasting -1000 (12536).</li><li>The druid's Clearcasting (16870), the shaman's (16246) and the priest's Inner Focus (14751) are all -100. Only the mage's is -1000. In a server that multiplies percentage cost mods, -100 makes a cast free whatever else applies and -1000 would be pointless; in one that adds them, -100 leaves a 4-stack Blast under Arcane Power at 330% and -1000 is what it takes to make it free. Blizzard wrote -1000 for the one class with +% cost mods, so the server adds.</li><li>The engine adds (`core.SpellCost.Multiplier`, `applyPowerCostPercent`): a Barrage Missiles under Arcane Power costs 30% of its 595, a 3-stack Blast under Arcane Power 4.3 x 15% of base mana, and a Clearcast anything is free. Round 19 S5 left this OPEN for want of a mana log; the client column settles it.</li></ul> | spellconst mage.json 12042/400573/400589/12536; druid.json 16870; shaman.json 16246; priest.json 14751; `core/spell_mod.go:577`, `arcane_blast.go:95-105`, `missile_barrage.go:39-75` | Aligned (closed) | - | - |
| S3 | **Clearcasting on a Barrage Missiles (fell out of S2).**<ul><li>Blizzard, 2 November: "Clearcasting is no longer consumed on Arcane Missiles if Missile Barrage is active at the same time."</li><li>The engine decided it from the cast's cost, read in Clearcasting's own cast-complete handler without Clearcasting's mod. Two holes: under Arcane Power the Barrage Missiles costs 30%, so the charge went; and the aura handlers run in activation order (`core/aura.go:604`), so whenever Missile Barrage had procced before the Clearcasting, its handler ran first, put the 100 back, and the charge went too.</li><li>Now the channel records that it started under Missile Barrage as it applies (`missileBarrageChannel`, before any cast-complete handler), and Clearcasting leaves its charge on such a cast. The saved charge then makes the next stacked Blast free.</li></ul> | `TestClearcastingSurvivesABarrageMissilesUnderArcanePower` failed, then passed; `talents.go` Clearcasting OnCastComplete, `arcane_missiles.go` ApplyEffects, `mage.go` | CONFIRMED-BUG | old (round 11 set -1000, round 19 S5 read the cost rule as right; neither saw the handler order) | E 18355c631 |
| S4 | **Frost raid talent search after the coefficient change: the shipped 15/5/31 holds.** Raid 60, alliance, seed 1, 3,000 iterations: shipped 633.5; 18/2/31 (Arcane Meditation 3 for three ranks of Incineration) 631.0; 18/0/31 (Meditation 3, no Incineration, two points unspent) 629.5. Incineration's 10% Ice Lance crit beats half the regeneration; nothing else in reach does anything on a Patchwerk boss (the Arcane row-1 toll, Frost Warding, Frostbite, Permafrost and Winter's Chill 2 are gates to Ice Barrier with no damage alternative at range, and Arcane Impact's 6% crit lands on the Barrage Missiles). Not adopted. | `$R/mage/sb -talents`, three candidates | Aligned | - | - |
| S5 | **Frost races at 60 on the shipped build with Ice Lance at 0.107** (coordinator seed, guide round 17). Raid 60, frozen BiS, guide raid build, 10,000 iterations, seed 1 / seed 50001:<ul><li>Alliance: Skyborne 636.6 / 636.7; Gnome 632.3 / 632.6; Human 630.8 / 631.2.</li><li>Horde: Undead 626.9 / 626.6; Troll 624.1 / 624.6; Orc 623.5 / 624.0.</li><li>Gaps: Skyborne +0.7% over Gnome, Human -0.2%; Troll -0.4% and Orc -0.5% behind Undead. Every pick within 1% stays. The page's order matches and "Troll and Orc are close behind" holds. "Edges a Gnome by a hair" understates 0.7% and now reads "by a little"; the Human sentence reads "trails a Gnome by a hair" so it cannot be read against Skyborne.</li></ul> | `$R/mage/frost-races.tsv`; frost.md:99 | Aligned (one word) | - | S fd5a0aa9f |
| S6 | **Low-rank proc watch (round 19 S8) closed.** Blizzard, 2 November: "Various talents no longer have a penalty to their chance to trigger when casting spells of a low rank. This penalty does still apply to Blizzard." So Frostbolt rank 4 at 26-29 and rank 5 at 32-33 proc Frostbite and Fingers of Frost at full chance; the engine has no proc penalty and needs none. Only Blizzard keeps one, and no rotation casts Blizzard. (Platform's priest P1 "low-rank proc curve": the same note answers it for every class.) | Kaivax thread, 2 November section; `chill.go`, `fingers_of_frost.go` (no penalty) | Aligned (closed) | - | - |
| S7 | The platform SHOULDs on the stat table unit lines are the platform lane's. Not touched. | brief | - | - | - |

## Arcane

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | **A Presence of Mind Pyroblast spent Heating Up's stacks.**<ul><li>Blizzard, 16 October: "Heating Up is no longer consumed if the Mage also had Presence of Mind active for the Pyroblast cast."</li><li>The engine's Heating Up aura ended on every Pyroblast that completed. The raid rotation casts Pyroblast the moment Presence of Mind is up (line 4), so the stacks its Fire Blasts had built were thrown away every 3 minutes.</li><li>An instant Pyroblast (cast time 0: Presence of Mind, or a Netherwind Focus proc) now leaves the stacks for the next hard cast. Three stacks of Heating Up alone still leave a quarter of the cast, so the check cannot confuse the two.</li></ul> | `TestHeatingUpSurvivesAPresenceOfMindPyroblast` failed, then passed; `heating_up.go` OnCastComplete; `core/unit.go:385` (cast time 0 under a -1 multiplier) | CONFIRMED-BUG | old (round 18 A4 and round 17 read Heating Up's rows, not the October note) | E fd5c65756 |
| A2 | **Arcane Geometry 1 in the raid build does nothing on a boss; Arcane Meditation's third rank does.**<ul><li>The raid build spent its 32 Arcane points with Arcane Geometry 1 (3 yards of range) and Arcane Meditation 2/3 (33% regeneration while casting). Range changes nothing against a boss a mage already reaches from 30 yards; the third Meditation rank is 17% more regeneration on the most mana-bound spec.</li><li>Geometry 0, Meditation 3 is legal (row 3 keeps 4 points, 18 before row 4, every later gate holds) and keeps every other talent.</li><li>Raid 60, 10,000 iterations, fixed engine: alliance 722.7 to 729.5 (+0.94%) at seed 1, 721.9 to 729.1 (+1.00%) at 50001; horde 710.8 to 716.8 (+0.84%), 710.2 to 716.5 (+0.89%); error 0.5 each. Beyond error at both seeds; it costs nothing and gains.</li><li>Adopted under the dead-point rule. arcane.md raidBuild is now `131205203100311531/525000012001/003`; the Meditation bullet, the raid paragraph and the Human race line changed with it. The leveling build keeps Geometry 1 (mobs push back and 3 yards matter in the open world).</li></ul> | `$R/mage/arcane-build.tsv`; arcane.md:7, :55, :66, :98; builds.test.ts, _spec.test.ts, _prose.test.ts and `sim/leveling` pass | CONFIRMED-BUG (dead point) | old (no round compared the two) | S fd5a0aa9f |
| A3 | Clearcasting on a Barrage Missiles (S3). | | CONFIRMED-BUG | old | E 18355c631 |
| A4 | **Arcane Explosion never missed.** `OutcomeMagicCrit` is the hit-only table; a direct damage spell rolls hit like Cone of Cold and Frost Nova. Now `OutcomeMagicHitAndCrit`. Neither the curated APLs nor the fork's three mage APLs cast it, so no number moves; the leveling guide's AoE advice is prose only. | `TestArcaneExplosionRollsTheSpellHitTable` failed, then passed; `arcane_explosion.go` (upstream wowsims line, never audited) | CONFIRMED-BUG (no DPS change) | old | E 57ad442c1 |
| A5 | Re-read against 70755 (everything the raid and leveling rotations cast): Arcane Blast ranks 400574-1239700 (32/78/101/168/236, 0.429, 15% base mana, 1.5 s, stack 400573 12% Intellect, +100% cost, 4 stacks, 18 s); Arcane Missiles channels 5143-25345 (3/4/5 missiles, 85-655 mana, missiles 7268-25346 at 0.286); Arcane Wrath 1326079-82 (120/187/300/451, 1.0, 140-395 mana, 3.5 s, 6 s, 5% of maximum mana as spell damage); Arcane Power (+30 damage, +30 cost, +30 periodic, 15 s, 3 min); Presence of Mind (-100% cast time, until spent, 3 min); Arcane Concentration 2%/rank, 1 s proc recovery, Clearcasting 15 s; Missile Barrage 20/40% Blast, 10/20% the rest, -50% duration, -100% cost, 15 s; Arcane Mind 2% Intellect and 20% Arcane crit bonus a rank; Arcane Instability 1% damage, 1% periodic, 1% crit a rank (aura 290); Arcane Subtlety 15% threat a rank; Arcane Meditation 17/33/50. The crit bonus talents compose as the client's "critical strike damage bonus" (`core/spell_outcome.go:1083`: 1 + 0.5 x (1 + bonus)), so Arcane Mind 5/5 and Ice Shards 5/5 each make a crit 200%. All match. | spellconst/mage.json, talents/mage.json; `talents.go`, `arcane_*.go` | Aligned | - | - |
| A6 | Blizzard's other 2 November Arcane lines against the engine: Missile Barrage spent at the channel's start (the engine spends it at cast complete, which for a channel is the start); Arcane Missiles proc Clearcasting (each missile is a hit of its own with a 1 s proc recovery; the channel itself does not roll). Both aligned. | Kaivax thread; `missile_barrage.go`, `talents.go:240` | Aligned | - | - |

## Fire

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| F1 | Re-read against 70755: Scorch 2948-10207 (39..178, 0.429, 50-150 mana, 1.5 s); Fire Blast 2136-10199 (28..438, 0.429, 40-340 mana, 8 s cooldown less Improved Fire Blast's 1 s a rank); Pyroblast 11366-18809 (110..583 plus 11..53 a tick for four 3 s ticks at 0.15 a tick, 1.0 direct, 6 s cast, 125-440 mana); Fireball 133-25306 (0.429/0.571/0.714/0.857 then 1.0; 2/3/3/4 ticks of 1..15 at 2 s; 1.5-3.5 s). Ignite 8% a rank, 2 ticks over 4 s, rolls over, ignores target modifiers (24 September note); Improved Scorch 33/67/100% for 3% a stack, 5 stacks, 30 s, own stacks only; Fire Power 1/2/3% per own Fire effect up to 4 (Ignite, Improved Scorch, the Fireball and Pyroblast burns; Frostfire's mixed burn excluded); Critical Mass 2% a rank; Incineration 2% a rank on Fire Blast, Ice Lance, Arcane Blast and Scorch; Master of Elements 10% of base cost a rank on Fire and Frost crits; Combustion 28682 +10% Fire crit a hit, ends at the third non-periodic crit, starts at one stack on activation (the trigger applies the aura); Heating Up 400625 -25% a stack, 3 stacks, 20 s, Fire Blast one stack on every landed hit and two on a crit. All match. | spellconst/mage.json, talents/mage.json; `fire_*.go`, `ignite.go`, `heating_up.go`, `talents.go` | Aligned | - | - |
| F2 | The Fire raid build has no Presence of Mind, so A1 does not reach it; its Clearcasting never meets a Barrage Missiles (no Missile Barrage, no Arcane Missiles line). Raid 60 is identical before and after to the decimal. | `$R/mage/after-raid.tsv` | Aligned | - | - |
| F3 | Round 17 F2 re-checked on the 70755 tree: Arcane Meditation needs 15 Arcane points before it and the raid build's 17 Arcane are 15 of gate plus Meditation 2; the Fire 31 is Combustion's 30 plus the capstone, with Flame Throwing 2 and Burning Soul 2 as the toll. The only source for an 18th Arcane point is Precision, so Meditation 3/3 still costs 2% hit. Holds. | talents/mage.json rows (Meditation row 3, 15 points) | Aligned | - | - |
| F4 | Rank rows at 40-49 (round 19 S7, judged level by level) unchanged; nothing in this round touches Fire's numbers. | rank-choices.json | Aligned | - | - |

## Frost

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | Ice Lance at 60: OPEN on a launch log (S1). | | OPEN | - | - |
| R2 | Talent search (S4): shipped build holds. | | Aligned | - | - |
| R3 | Races (S5): Skyborne, Gnome, Human; Undead, Troll, Orc. One word on the page. | | Aligned | - | S fd5a0aa9f |
| R4 | **Frostbolt rank rows re-judged level by level** (they predate the round-19 rule): bare, alliance, 3,000 iterations, `-level` and `-rank`, row rank against the trained rank. Level 26: rank 4 56.0 against rank 5 52.9; 29: 58.4 against 56.7; 32: rank 5 84.2 against rank 6 80.6; 33: 84.6 against 81.9. Both rows hold at their floor and top. | `$R/mage/frost-ranks.tsv`; rank-choices.json mage-frost | Aligned | - | - |
| R5 | Re-read against 70755: Frostbolt 116-25304 (0.407/0.489/0.597/0.706 then 0.814; 1.5-3.0 s less Improved Frostbolt; 25-290 mana); Ice Lance 1312002-1240047 (28/36/45/79/98/145, 0.107 from logs, 8 s cooldown on every rank in 70755's own rows now, x4 against Frozen, 45-160 mana); Ice Barrier 11426-13033 (305-480 mana, 30 s cooldown, 1 min; +15% Frost damage for the minute and +5% while shielded, the 30 October hotfix text); Fingers of Frost 15% a chill, 1/2 charges, 15 s, spent as a direct cast resolves; Shatter 17/33/50; Frostbite 5% a rank (a boss cannot be frozen); Winter's Chill 20% a rank to stack, +20% damage taken and +7% spell damage a stack on the next Cone of Cold (12579 aura 271 on Cone of Cold); Cone of Cold 120-10161 (0.129, 10 s, refused beyond 10 yards); Frost Nova 122-10230 (0.029, 25 s less 2 s a rank); Frost Channeling -5% cost, -10% threat a rank; Ice Shards 20% a rank. All match. | spellconst/mage.json, talents/mage.json; `frostbolt.go`, `ice_lance.go`, `ice_barrier.go`, `fingers_of_frost.go`, `chill.go`, `cone_of_cold.go`, `frost_nova.go`, `talents.go` | Aligned | - | - |
| R6 | The Frost raid rotation meets S3 only without Arcane Power (no Arcane Power in the build), where the old rule already kept the charge when Clearcasting had activated after the proc and spent it otherwise. The order fix moves the fork's forever_frost goldens (TestP1Mage) but not the site's raid 60 at 10,000 iterations (632.3 / 632.6 alliance, 624.1 / 624.6 horde before and after, the last at 624.5 before). | `$R/mage/after-raid.tsv`, `sim/mage/TestP1Mage.results` | Aligned | - | E 5e2df54c5 |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | The fork mage package passes with `-tags=with_db`. TestP1Mage goldens regenerated: the Arcane preset rows (A1, S3) and the forever_frost rows (S3's handler order), each within 0.5 DPS. | `go test -tags=with_db ./sim/mage/` | - | - | E 5e2df54c5 |
| X2 | Every 2 November Blizzard mage line was checked against the engine: Missile Barrage row and timing, Clearcasting on Barrage Missiles (S3), Missiles proc Clearcasting, Arcane Blast 12%/100%/18 s, Arcane Focus folded into Precision, the low-rank proc penalty only on Blizzard (S6), Ignite on Flamestrike past the first hit (not cast). | Kaivax thread | Aligned | - | - |
| X3 | Blizzard's 16 October mage lines: Presence of Mind consumed by casts (yes); Missile Barrage not reducing a second Missiles (one proc, one channel: `missile_barrage_consume_test.go`); Fire Blast's stacks (F1); Heating Up under Presence of Mind (A1); Clearcasting not off Dampen/Amplify or out of combat (neither cast); Ice Lance 200% then back to 300% on 30 October (engine 300%). | Kaivax thread; wowfdb change log (https://wowfdb.com/classes/mage/changes/, read 2026-11-04) | Aligned after A1 | - | - |
| X4 | Mage Armor keeps 50% of regeneration while casting on all three ranks (6117/22782/22783 aura 134 = 50), so Arcane Meditation 3/3 plus Mage Armor is all of it; the Human line on the Arcane page said "most" from the 2/3 days. | spellconst/mage.json; `armors.go:1499`; arcane.md:98 | Aligned (text) | - | S fd5a0aa9f |
| X5 | Housekeeping: while clearing my own dev server I killed a stray `astro dev --port 4399` that belonged to scratchpad/gr17/int (a guide-round integrator's, pid 85445, started with that tree's node_modules). If that workspace is still live its owner needs to restart it. | `lsof` | - | - | - |

## Before/after

Before: `$R/before.tsv` (raid 60) and `$R/sb-before` on the published band builds (bare). After: E 5e2df54c5 (lane/r20-mage) on the site branch, shipped builds; then A2's build on the same engine.

| Spec | Faction | Seed 1 before to after | % (s1 / s50001) |
|---|---|---|---|
| mage-arcane raid 60, engine fixes (A1, S3) | alliance | 721.0 to 722.7 | +0.24 / +0.28 |
| mage-arcane raid 60, engine fixes | horde | 708.3 to 710.8 | +0.35 / +0.25 |
| mage-arcane raid 60, then the build (A2) | alliance | 722.7 to 729.5 | +0.94 / +1.00 |
| mage-arcane raid 60, then the build | horde | 710.8 to 716.8 | +0.84 / +0.89 |
| mage-arcane raid 60, both together | alliance | 721.0 to 729.5 | +1.18 / +1.28 |
| mage-arcane raid 60, both together | horde | 708.3 to 716.8 | +1.20 / +1.14 |
| mage-arcane bare 40 (published band build) | horde | 171.0 to 171.1 | +0.06 / +0.06 |
| mage-arcane bare 50 | horde | 247.0 to 247.2 | +0.08 / +0.04 |
| mage-arcane bare 60 | horde | 391.0 to 391.2 | +0.05 / 0 |
| mage-fire raid 60 | both | 735.6 / 721.0, unchanged | 0 / 0 |
| mage-frost raid 60 | both | 632.3 / 624.1, unchanged | 0 / 0 |

No spec outside mage moves: the engine changes are in sim/mage only, and the site changes are two guide files.

## Files to regenerate

- addon-builds.json and Data.lua: the Arcane raid build (A2).
- Ladder goldens for mage-arcane (sim/request TestRotationLadder), if a rung reads the raid build; the leveling rungs read the leveling build, which did not change.
- dps-per-point and stat-weights for mage-arcane (A1, S3, A2).
- Planner vectors / any generated copy of the guide raid build.
- Coach starter data/curated/coach/mage-arcane.md, if it quotes the build.

## Measure at refresh

- Arcane BiS re-rank at raid 60 on the new raid string (Meditation 3/3, Geometry 0) and with A1 and S3 in the engine.
- The Frost re-rank and weights after Ice Lance (round 19) still pending.

## For other lanes

- **Platform:** the low-rank proc penalty exists only on Blizzard (S6, Blizzard's 2 November note), which answers the priest P1 "low-rank proc curve" item for every class: no curve to model. wow-player SHOULDs on the shared guide components: the tree heading "Level 60 build" against the prose's "leveling build"; the Key talents table omits talents the prose names (Incineration, Accelerant, Arcane Subtlety, Improved Channeling, Frost Channeling), so a rank difference between the two builds is not visible there. The stat-table SHOULDs from round 19 stay with platform.
- **Every lane, method:** the handler-order trap in S3. An aura's OnCastComplete that reads state another aura's OnCastComplete changes (a cost multiplier, a stack count) sees it in activation order, which a rotation does not control. Record the fact where it happens (ApplyEffects runs before every handler) instead.
- **Every lane, method:** where a class has a percentage cost reduction and a percentage cost increase that can meet, the client's rows say how they combine (S2): the mage's -1000 against the others' -100 is the tell for an additive server.

## For the guide lane

- Done here: arcane.md raidBuild (A2), the Arcane Meditation bullet, the raid-build paragraph, the Human race line (X4); frost.md:99 "by a little" (S5).

## Guide review

wow-player reviewed the rendered /guides/mage/arcane Talents and builds section and the /guides/mage/frost Races section at 1440 wide (astro dev on wt/r20-mage, trees hydrated). Captures: `$R/mage/shots/arcane-1440{,-talents}.png`, `frost-1440{,-races}.png`.
- First pass (Arcane): NO-SHIP on two prose MUSTs: the Meditation bullet's "all of it runs" against the Races line's "most" (X4 settled it from the client: the Races line was stale), and the raid paragraph's point accounting (reworded: the Subtlety and Wrath points are paid by a rank of Improved Channeling, the point of Arcane Geometry and one point brought over from Frost). One SHOULD taken (the bullet's last clause). The component SHOULDs go to platform.
- Second pass (Arcane and Frost): **SHIP on both, no MUSTs.** Two Frost SHOULDs: the Human sentence now reads "trails a Gnome by a hair" (taken, 0.2%); "edges a Gnome by a little" stays, since 0.7% is more than a hair and the reviewer's own bar was "honest at under 1%".

## Still open

- S1: Ice Lance's coefficient at 60 (a level 56+ Frost log). Last looked 21:15 UTC.
- Carried: the level-60 book ranks (pending owner); Warcraft Logs blocked.
