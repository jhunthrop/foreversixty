# Sim-accuracy round 19: mage

Specs: mage-arcane, mage-fire, mage-frost. Site wt/r19-mage (from int/r19 f91e9d3f9). Fork lane/r19-mage (from int/r19 6ea675594). Client 1.60.1.70755 with the late re-layer.

Method:
- Seed items first, then the re-audit.
- Raid 60: spec-breakdown, 10,000 iterations, seeds 1 and 50001, both factions, `-bis-dir $R/bis-frozen`, guide raid builds (`ALLRT=1`).
- Leveling levels: the shipped leveling order turned into a talent string per level (`mage-tools/` here and `$R/mage/tools/ladder.py`), bare preset, `-band`/`-level`.
- Every sim went through `$R/sim.sh`.
- Logs: foreverlogs beta reports, damage events through the public API, fetched and fitted with the scripts in `mage-tools/`.

**Confirmed and fixed: 2.**
- Ice Lance's spell power coefficient, tagged old (rounds 4, 6-15 carried it, blocked on logs). Frost raid 60 drops 9.1%.
- The guide stat table's unit line ("1 point of hit"), tagged old. Text only, every class page.

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Missiles/Blast boundary (guide round 15 MAGE-R15-1): it holds at 22 / 23.**<ul><li>Level 22 (Missile Barrage 1/2): Missiles-first 49.0 against 41.4 for Blast-first (alliance), 46.6 against 39.3 (horde). The shipped APL leads with Missiles.</li><li>Level 23: Blast-first 51.0 against 50.2 (alliance); 48.4/48.5 against 48.3 (horde, inside error).</li><li>Level 24: Blast-first 59.2 against 57.9/58.0 (alliance), 56.3/56.5 against 55.6/55.7 (horde).</li><li>The figures are identical at seeds 1 and 50001.</li><li>arcane.md:116 and the line-7 card note already say "20 to 22" and "from 23". Nothing changes.</li></ul> | ladder talents per level, bare band 20, `-rotation` with line 7 forced on or removed; `mage-tools/mb-boundary.tsv` (cur, Missiles-first, Blast-first) | Aligned | - | - |
| S2 | **Ice Lance's coefficient is about 0.107, not 0.4286.**<ul><li>**Client rows (SpellEffect, wago, builds 1.60.1.70755 and 1.60.1.70790, identical):** every effect of every rank has EffectBonusCoefficient 0, BonusCoefficientFromAP 0 and Coefficient 0. That covers the damage effect (Effect 2) of 1312002 (row 1342605, bp 28), 400640 (1045851, bp 36), 1240044 (1238139, bp 45), 1240045 (1238143, bp 79), 1240046 (1238147, bp 98) and 1240047 (1238151, bp 145), each with variance 0.1667. The three dummies (Effect 3, bp 300/20/5) are 0 too, as is the learn row 401732.</li><li>So 0.4286 was never a client value: it is the generator's instant-spell default (1.5/3.5). The Forever tooltip (nether 1240047, "Deals 140 to 164 Frost damage... 300% increased damage to Frozen targets") has no spell power term either. The server sets the value.</li><li>**Logs: 9 players in 7 reports** (2699, 2727, 2730, 2747, 2762, 2763, 2786), levels 24-35, ranks 1312002 and 400640.</li><li>Each player's spell power comes from their own Frostbolt (0.814), Cone of Cold (0.129) and Frost Nova (0.029) non-crit hits in the same fights, at the client's coefficients. The three agree within about 10% for every player with all three. A flat damage multiplier would break that agreement, because Frost Nova's 0.029 makes its estimate very sensitive to one.</li><li>Ice Lance's own non-crit hits then give, per player (unfrozen / frozen ÷ 4): Alte .113/.126, Barry .114/.087, Greatwizardx .107/.109, Jaced .119/.112, Mana .101/.099, Murph -/.089, Schmexy .107/.115, Shortus .102/.110 and .112/.101, Slampiece .096/.104. Hit-weighted mean 0.105 over 334 hits.</li><li>At 0.4286, Greatwizardx's unfrozen Lances would sit near 91; they sit at 50 (n=38).</li><li>The frozen hits are 4.0 times the unfrozen ones, which again confirms round 6's ruling that the ×4 applies to the whole hit.</li><li>Fix: `iceLanceSpellCoeff = 0.107` on every rank. The low-rank penalty still applies on top of it.</li></ul> | `ice_lance.go`; test `TestIceLanceSpellPowerCoefficientFollowsTheLogs` failed (0.4286), then passed; SpellEffect CSVs at `$R/mage/il/client/` (also the integrator's `$R/icelance/se-707{55,90}.csv`); `mage-tools/fit.py`, `fit-final.txt` | CONFIRMED-BUG (nine independent players agree; the client column is 0, so the value is server-side and comes from logs) | old (round 4 R3, round 6 R2/R6, carried to round 18) | E c718869a2 |
| S3 | **Uncertainty on S2.**<ul><li>The logs are beta (2026-10-01 to 10-19). Launch could carry a different server value; the client rows did not change from 70755 to 70790.</li><li>The per-player spread is 0.09 to 0.13. 0.107 against the mean of 0.105 is worth under 0.3% of Frost.</li><li>No web source publishes the value. The SpellCoeffsForever addon computes from the cast-time formula and does not cover Ice Lance (curseforge.com/wow/addons/spellcoeffsforever, read 2026-11-04).</li><li>Re-check against a level-60 launch log of a Frost mage once one exists.</li></ul> | as S2 | recorded | - | - |
| S4 | **Arcane Subtlety's −15 resistance: no effect, aligned.** The published target has 0 resistance, and the level-based part is not penetrable (`core/spell_resistances.go:146`). Round 14 M6 said the same. | engine read | Aligned | - | - |
| S5 | **Cost mods (Arcane Power on stacked Blast and on Barrage Missiles): still OPEN.**<ul><li>The engine adds percent cost mods: Missile Barrage −100 + Arcane Power +30 leaves a Barrage Missiles at 30% cost, and a 3-stack Blast under Arcane Power costs 15% × 4.3.</li><li>Original 1.12 (vmangos) adds them. The modern server code that Classic Era runs on multiplies them, which would make a Barrage Missiles under Arcane Power free and a 3-stack Blast cost 15% × 5.2.</li><li>Foreverlogs records no mana per cast, so the logs cannot settle it. Web searches turned up no Forever or Classic statement either.</li><li>It is settled in game at level 40 or above: cast a Missile Barrage Missiles inside Arcane Power and watch the mana bar.</li></ul> | `arcane_blast.go:94-103`, `missile_barrage.go:39-75`; foreverlogs openapi (no resource fields on events) | OPEN | - | - |
| S6 | **Hit-weight unit wording.**<ul><li>The line under every guide's Stat weights table read "1 point of hit = 0.48 DPS at level 60".</li><li>Hit and crit weights are per rating point (10 hit rating, 14 crit rating = 1%, `rating_factor` in the BiS files), so a Classic player reads "1 point of hit" as 1% and overrates it tenfold.</li><li>It now reads "1 hit rating = 0.48 DPS at level 60." A second line names every rating stat the table prices: "Hit and Critical strike are per rating point: 10 hit rating is 1% hit and 14 critical strike rating is 1% critical strike."</li><li>The change touches every class page whose table carries hit or crit. Text only.</li></ul> | rendered /guides/mage/{arcane,fire,frost} (`$R/mage/shots/*-stats.png`); tests in `copy.test.ts`, `rail-stats.test.ts` and `_spec.test.ts` failed, then passed | CONFIRMED-BUG (text) | old | S 023f89ce8, S 86323e7c5 |
| S7 | **40-59 rank rows, judged level by level.** Fire is the only mage spec with rows there. Ladder talents, bare band 40, alliance, 3,000 iterations, DPS:<ul><li>L40: row 163.6 against FB5 160.5 and Scorch 3/4 159.7/155.2.</li><li>L43: 169.1 against 167.4 and 165.4/163.3.</li><li>L46: 169.3 against FB5/6 167.6/166.9 and Scorch 2/4/5 168.1/167.1/166.1.</li><li>L49: 174.8 against FB5 174.8 (tie) and Scorch 2/4/5 167.1/173.3/173.5.</li><li>Every row holds at floor, +3, +6 and +9. At 49, Fire Blast 4 ties 5, so the row stays.</li></ul> | `mage-tools/fire-ranks.tsv`; rank-choices.json:197-237 | Aligned | - | - |
| S8 | **Low-rank proc watch (Frostbite), flag only.** Frost's rows cast Frostbolt rank 4 at 26-29 and rank 5 at 32-33. The leveling build takes Frostbite at 25. If a low rank procs Frostbite less (Blizzard names Frostbite), those rows lose some freezes they are credited with now. No code change; no curve is known. | rank-choices.json:239-258; addon-builds leveling order | Flagged | - | - |
| S9 | Offensive Cold Snap is decided (keep it for Ice Block). The Arcane card SHOULDs go to the platform lane. Not reopened. | brief | - | - | - |

## Arcane

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | Round 18's own changes hold. The rank-tagged Missile Barrage gate gives Missiles at 22 and Blast from 23 (S1). Raid 60 matches before.tsv on all 4 rows. | `mage-tools/after-il.tsv` | Aligned | - | - |
| A2 | Log check on the Arcane spells. Arcane Explosion's spell power matches the frost spells for Dainer (147 against 103-126) and Chocolat. Arcane Missiles reads 1.5 to 1.7 times higher for all three Arcane mages. That fits Arcane Blast's stacks (12% of Intellect as spell damage), which feed every Missiles cast after a Blast, so the logs cannot separate the Missiles coefficient from the stack bonus. No evidence against 0.286 (client column). | `mage-tools/fit2.py` (reports 2721, 2726, 2809) | Aligned (not separable) | - | - |

## Fire

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| F1 | **Fire coefficients check against the logs.** Matilda (report 2749, 21 Fire points) reads 84-103 spell power from Cone of Cold, Frost Nova and Arcane Explosion. Her Fire Blast, Scorch and Blast Wave read 116-128, which comes down to 82-88 once Fire Power's 10% is applied. Fire Blast 0.429, Scorch 0.429 and Blast Wave 0.129 (client columns) fit. | `mage-tools/fit2.py` | Aligned | - | - |
| F2 | Rank rows at 40-49 (S7). | | Aligned | - | - |

## Frost

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | Ice Lance's coefficient (S2). | | CONFIRMED-BUG | old | E c718869a2 |
| R2 | **The rotation still holds after R1.** Dropping the Fingers of Frost Ice Lance line loses 2.5% (633.5 against 618.2, alliance, seed 1, 3,000 iterations): the ×4 against a Fingers of Frost "frozen" boss still beats another Frostbolt. The APL, the card notes and frost.md:54/71/117 make no spell-power claim about Ice Lance, so no text changes. | `$R/mage/rot-frost-nolance.json` | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Ice Lance is the only mage damage spell with a zero client coefficient. Every other cast spell carries its own column (Frostbolt 0.814, Fireball 1.0, Fire Blast/Scorch/Arcane Blast 0.429, Missiles 0.286 a missile, Cone of Cold 0.129, Frost Nova 0.029, Arcane Explosion 0.143, Pyroblast 1.0, Arcane Wrath 1.0). | spellconst/mage.json | Aligned | - | - |
| X2 | Wago lists 1.60.1.70790 (product wowf). Integrator ruling: it is the live client, 70755 with its hotfixes baked in; the integrator activates it after launch. The Ice Lance rows do not change. | wago.tools/api/builds, 18:35 UTC | Recorded | - | - |
| X3 | The fork mage package passes with `-tags=with_db`. TestP1Mage goldens regenerated: the Frost rows and the item-set rows that run the Frost rotation (dps 832.0 → 764.9 on the first Frost row). | `go test -tags=with_db ./sim/mage/` | - | - | E c718869a2 |

## Before/after

Before: `$R/before.tsv` and `$R/sb-before` (same flags; Frost re-run with `-talents` from rt and `-bis-dir`: 696.0). After: E c718869a2 on the site branch.

| Spec | Faction | Seed 1 before → after | % (s1 / s50001) |
|---|---|---|---|
| mage-frost raid 60 | alliance | 696.0 → 632.3 | −9.15 / −9.12 |
| mage-frost raid 60 | horde | 686.9 → 624.1 | −9.14 / −9.12 |
| mage-frost bare 40 | horde | 191.4 → 174.2 | −9.0 / −9.0 |
| mage-frost bare 50 | horde | 267.0 → 247.3 | −7.4 / −7.4 |
| mage-frost bare 60 | horde | 412.0 → 382.0 | −7.3 / −7.2 |
| mage-frost bare 30 | horde | 80.1, unchanged | 0 / 0 |
| mage-arcane, mage-fire raid 60 | both | unchanged (721.0 / 708.3; 735.6 / 721.0) | 0 / 0 |

No spec outside mage moves. The engine change is in sim/mage only; the web change is text.

## Files to regenerate

- Ladder goldens for mage-frost (sim/request TestRotationLadder): every rung with Ice Lance (30+).
- dps-per-point and the Frost stat weights (spell power is worth less to Frost).
- addon-data.json / Data.lua for any Frost weights or DPS they carry.
- Coach starter data/curated/coach/mage-frost.md, if it quotes a Frost figure.

## Measure at refresh

- Frost BiS re-rank and weights after S2: spell power and Frost power drop for Frost.
- Frost raid talent search after S2 (Arcane Impact and the Shatter/Fingers of Frost points), with a handful of candidates. Not run here; Ice Lance still earns its line (R2).

## For other lanes

- **Platform:**
  - The wow-player SHOULDs on the stat table: the "DPS per point" header covers both points and rating points ("DPS per point or rating"); the two unit lines want more contrast than the haste notes.
  - The Spell haste row reads "Not significant" on Arcane but "Not on any item" on Fire and Frost.
- **Every lane:** the stat table's unit line and the new rating line change on every class page. Text only, from S6.
- **Every lane, method:** foreverlogs damage events, plus a spell with a known client coefficient in the same fight, give a player's spell power. That settles server-side coefficients (`mage-tools/fit.py`). The same may apply to Seal of Righteousness rank 8's 0 column.

## For the guide lane

- Nothing in mage prose depends on Ice Lance's coefficient. The guide text did not change.

## Guide review

wow-player reviewed the rendered stat tables on /guides/mage/arcane, /fire and /frost at 1440 wide (astro dev on wt/r19-mage). Captures: `$R/mage/shots/*-stats.png`.
- First pass: NO-SHIP. The captures missed the table.
- Second pass: **SHIP, no MUSTs.** The crit-rating SHOULD was taken; the rest are listed under For other lanes.

## Still open

- S5: cost-mod stacking (Arcane Power with Missile Barrage and Blast stacks). Settled in game at 40 or above.
- S3: re-check Ice Lance against a level-60 launch log.
- Carried: the level-60 book ranks (pending owner); Warcraft Logs blocked. Foreverlogs was last checked at 18:35 UTC, before launch, so no live logs yet.
