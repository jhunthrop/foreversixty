# Round 15, lane warrior (Arms, Fury, Protection)

Site `wt/r15-warrior` on int/r15 fd320b60e; engine `lane/r15-warrior` on fork int/r15 d1d9f2775. Client 1.60.1.70755.
Client rows are the push-layered 70755 tables (`scratchpad/c755/fetch/data/builds/1.60.1.70755/raw/`); talent rank values come from `data/builds/1.60.1.70755/talents/warrior.json` and TraitDefinitionEffectPoints / CurvePoint (`c755/others-tools/curvediff.py`).

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| 1 | Known red: "Fury rotation expected-names test missing Reckless Swings". It is `data/tests/test_apl.py::test_every_named_spell_matches_its_own_label_by_name`. The rotation is right. 70755 made Reckless Swings (1329555) a Fury talent ability and the curated Fury list casts it, but the test's `EXPECTED_ABILITY_NAMES` sidecar was never updated. **The test was wrong, not the engine or the rotation.** Saw it fail, added the row, saw it pass. `spells.json` names 1329555 "Reckless Swings". The sim smoke test `TestEveryWrittenRotationRunsInTheEngine/warrior-fury` was already green (it declares 1329555 unresolved under ForeverFuryTalents). | test_apl.py:308 | CONFIRMED (test) | 70755 d009d90a3 | S a45b11f30 |
| 2a | **Guides S2, Fury.** The build spent 1 point in Lingering Rage (rage decays later "after leaving combat") and 1 in Piercing Howl (a 10 yd daze). Neither does anything inside a fight. Moved both to Improved Berserker Rage 2 (node 105978, tier 4: "will instantly generate 10 Rage"). The rotation now casts Berserker Rage below 80 rage, right after Reckless Swings (18499: 30 s cooldown, 1.5 s global, Berserker Stance). The build stays legal: tier 3 keeps Cruelty 5 + Unbridled Wrath 5 + Furious Precision 3 + Enrage 2 = 15. Placing the line after Bloodrage instead lost 0.4%. Gain is +0.2 to +0.5% at both seeds and both factions: under 1%, so this is adopted under the dead-point rule, not the gain rule. Web check: one Forever guide (wowtbc.gg, undated, 2026) presses Berserker Rage "to generate rage". Warcraft Tavern's page has only a level-30 beta build. Icy Veins and Mobalytics returned 403. | talents/warrior.json; SpellCooldowns 55298; SpellEffect 697806 | CONFIRMED-BUG | old (in the build since cb1325223 / 70aff3eaa), guides-handoff | S c328e4192, E 88b1d57dc |
| 2b | **Guides S3 and S4, Arms raid.** A raid boss never swings at an Arms warrior, so Deflection 4 (parry) and Blood Craze 2 (a heal after a crit taken) did nothing in the raid build. The 70755 carry turned Deflection 2 into 4 when it absorbed the Improved Tactical Mastery and Bloodthrill cuts. New raid build `30303213132215201/05050005012/1`: <br>- Arms 32: Improved Heroic Strike 3, Improved Rend 3, Improved Tactical Mastery 3, Improved Overpower 2, Anger Management 1, Deep Wounds 3, Spearing Strike 1, Two-Handed Spec 3, Impale 2, Bloodthrill 2, Sweeping Strikes 1, Weaponmaster 5, Improved Slam 2, Mortal Strike 1.<br>- Fury 18: Cruelty 5, Unbridled Wrath 5, Enrage 5, Reckless Swings 1, Improved Execute 2.<br>- Protection 1: Improved Bloodrage 1.<br>Dropped: Deflection 4, Blood Craze 2, Improved Charge 2 (6 rage on one pre-pull Charge) and one Improved Bloodrage point. The rotation casts Reckless Swings after Mortal Strike and Spearing Strike after Slam. Enrage is the Fury gate filler: it does nothing in the sim, but it procs on raid damage in a real raid. S4: the third Improved Heroic Strike point is now a tier-3 gate point (tiers 0-2 hold exactly 15), so it stays. Gains: +5.2 to +5.3% at both seeds and both factions. Spearing Strike alone is +1.9% on the default target and +4.5% on Dragonkin (3000 iterations). Reckless Swings plus Improved Execute is +3.4%. The adoption rule passes: the build keeps Mortal Strike and every unmodeled talent, stays in the Arms tree, and drops nothing the rotation casts. | talents/warrior.json; Spearing Strike SpellEffect 1339573 (75), SpellCooldowns 101004 (20 s), SpellPower 314497 (15 rage) | CONFIRMED-BUG | 70755 3d4712981 (Deflection 2→4) and old (Deflection 2, Blood Craze 2), guides-handoff | S 9ba13b069, E 88b1d57dc |
| 2c | WARRIOR-13, Protection: Improved Thunder Clap 3→1 into Improved Shield Wall 2. Held for the owner by the guide ledger (NEEDS-OWNER). Thunder Clap costs 4 more rage, and Shield Wall is cast only below 40% health, which never happens in the tank sim (0 casts per fight, 0% chance of death). Under the threat metric it cannot win. | tank breakdown below | OPEN (owner) | guides-handoff | - |
| 2d | Guides: "the bis weights carry expertise for all three warrior specs". The BiS files carry `expertise_to_cap`, written by the shared ranker (`sim/cmd/leveling-bis/expertisecap.go`). Weapon skill from gear goes to PseudoStats weapon skills (`core/character.go:284`), not to Expertise. No Forever item has expertise, so the row is inert for warriors. | | OPEN, for other lanes (ranker) | guides-handoff | - |
| 2e | Fury, Spearing Strike from Improved Bloodrage 2→1 (`30323013102/050530025122100231/1`). +0.7% at seed 1 (3000 iterations), under 1%, and the dropped point does something. Not adopted. | | Aligned (held) | - | - |
| 3 | The 70755 warrior push, every row: Slam category cooldown 18→15 s on all 11 ids (including the Improved Slam overrides 1310196-1310200, which carry the same base points 16/32/43/68/87); Improved Slam e2 -2500/-5000; Improved Whirlwind e0 -1000/-2000 ms, e1 15/30%; Reckless Swings (15 s, 1 s global, 4 charges, 0x4 white swings only, 5 rage each, 25% base health); Flurry 10/20/30; Improved Tactical Mastery 5/10/15; Bloodthrill 10/20; Two-Handed Spec 2/4/6; Spearing Strike 75 + 75; Booming Voice reaching Piercing Howl. A full 70753→70755 row diff of every SpellClassSet 4 spell found nothing else (only the Attributes_14 bit 64 on Sweeping Strikes and Concussion Blow). The engine follows every row: slam.go:33-58, whirlwind.go:39-46, reckless_swings.go, talents.go:146-167, spearing_strike.go, stances.go:34. Guide builds are legal on 70755. Dead points: Fury fixed in 2a, Arms raid in 2b. Protection has none: Improved Rend 3 gates Deep Wounds, Concussion Blow gates Shield Slam. Fury's Improved Charge 2 is the gate to Impale; every alternative point at that depth (Deflection, Improved Overpower, which Berserker Stance cannot cast) does nothing either. | curvediff output; 70753/70755 table diff | Aligned | - | - |

## Arms

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | **Spearing Strike works in any stance in the engine; the client allows Battle Stance only.** SpellShapeshift row 58595 for 1310222 is ShapeshiftMask 65536 (Battle Stance). An Arms warrior who swaps to Berserker Stance for the execute phase cannot press it there. No DPS moved: after the swap Execute always outranks it. | SpellShapeshift.csv 58595 | CONFIRMED-BUG | old (de47da919, 2026-10-06; no round checked stance masks) | E 091fa0e56 |
| A2 | Every other warrior stance mask matches the client: Execute and Hamstring in Battle/Berserker, Rend in Battle/Defensive, Overpower, Charge, Retaliation and Sweeping Strikes in Battle, Whirlwind, Pummel, Recklessness and Berserker Rage in Berserker, Revenge and Shield Block in Defensive, Thunder Clap per stanceMask. Mortal Strike, Slam, Bloodthirst and Reckless Swings have no row (any stance). | SpellShapeshift.csv | Aligned | - | - |
| A3 | Checked against client rows: Mortal Strike r4 (160 weapon bonus, 30 rage, 6 s), Overpower r4 (+35, 5 rage, 5 s), Rend r7 (21 per 3 s tick, 10 rage), Heroic Strike r8 (+138, 15 rage), Slam ranks (above), Execute r5 (600 + 15 per rage, 15 rage, Improved Execute 3/5), Recklessness (100% crit, +20% damage taken, 30 min), Deep Wounds 20/40/60% over 12 s, Improved Rend 12/23/35%, Impale 10/20%, Weaponmaster 1/3/1 per point, Improved Overpower 25/50%, Anger Management 1 rage per 3 s, Improved Charge 3/6. | q.py rows above | Aligned | - | - |

## Fury

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| F1 | Bloodthirst r4: 48 + 45% AP (e0 48, e1 dummy 45), 30 rage, 6 s. Whirlwind 25 rage and 10 s; the off-hand strike 462891 is under Improved Whirlwind's multiplier on the closest target only. Death Wish: 10 rage, 3 min, 1.5 s global, +20% physical, +5% damage taken. Unbridled Wrath 12% a point, white swings only (ProcTypeMask 4). Dual Wield Spec 5%/10% a point. Furious Precision 4/7/10. Enrage 30% for 2% a point, 12 s. Flurry consumed by white swings. Reckless Swings: 10.6 casts per fight, 19.8 rage a cast. Bloodrage plus Improved Bloodrage 2: 29.6 rage a cast. | client rows; Fury breakdown, horde, seed 1 | Aligned | - | - |
| F2 | Berserker Rage gives rage when hit (2 a melee hit). Checked because the client's 18499 has only two aura-77 effects. The Forever tooltip says "generating extra rage when taking damage" (https://nether.wowhead.com/forever/tooltip/spell/18499, fetched 2026-11-03), as vanilla did server-side. No change. Moot for DPS: a DPS warrior is not hit. | Wowhead tooltip | Aligned | - | - |
| F3 | Gore Drinker's comment (talents.go:420) says "needs Enrage x5". The client prerequisite is Blood Craze 3 (`prereq 105934/3`). Comment only; the talent is not modeled. | talents/warrior.json | Aligned (comment nit) | - | - |

## Protection

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | Tank breakdown (alliance, seed 1, 3000 iterations): 1053.7 threat/s, Shield Slam every 6 s, Revenge 30 per fight, Shield Block 80% uptime, Demoralizing Shout kept up (O1), Thunder Clap refreshed under 3 s, Shield Wall 0 casts. Values checked: Improved Revenge 20/40/60, Improved Thunder Clap 2 a point, Defiance 5 a point, Focused Rage 1 a point, Bastion 2 a point, Anticipation 4 defense a point, Shield Specialization, Master of Defense. Sunder threat (effect 63) and Demoralizing Shout threat were re-verified in round 14. 70755 changed nothing Protection casts (the Protection APL has no Slam). | tank breakdown | Aligned | - | - |
| P2 | The boss's 268 AP and Demoralizing Shout ruling: no side effect found; Demoralizing Shout uptime holds. | | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | The fork reference build ForeverFuryTalents takes no Improved Berserker Rage. The fork's Fury list now casts Berserker Rage, so TestP1DPSWarrior moves -0.3% (Average-Default 1589.6 to 1585.5). Regenerated in E 88b1d57dc. Retuning the reference build (Lingering Rage 3→1 for Improved Berserker Rage 2) is optional. | | note | - | E 88b1d57dc |

## Before/after (raid 60, 10000 iterations; Arms and Fury on the guide raid builds via `rt`)

| Spec | Faction | Seed 1 | Seed 50001 |
|---|---|---|---|
| warrior-arms | alliance | 750.7 → 790.6 (+5.32%) | 750.7 → 790.3 (+5.28%) |
| warrior-arms | horde | 763.3 → 803.5 (+5.27%) | 763.2 → 803.1 (+5.23%) |
| warrior-fury | alliance | 830.1 → 832.6 (+0.30%) | 829.1 → 833.0 (+0.47%) |
| warrior-fury | horde | 831.0 → 833.6 (+0.31%) | 831.6 → 833.5 (+0.23%) |
| warrior-protection | both | unchanged (1053.1 / 1042.7 threat/s) | unchanged |

Bare 60, horde, on the same builds: Fury 313.6 → 314.3 / 313.3 → 314.9; Arms 274.7 → 290.9 / 274.5 → 290.7. A45b11f30 and E 091fa0e56 move nothing. **No commit moves a spec outside the warrior class.**

## Files to regenerate

- `make apl-sync` is a no-op check: the fork copies (ui/warrior/apls/forever_{arms,fury}.apl.json) and the site copies (sim/request/apl/warrior-{arms,fury}.apl.json) were synced by hand, and `make apl-check` passes.
- addon-builds.json: the Fury `build` changed (leveling point order).
- sim/request TestRotationLadder goldens: warrior-fury (every band), warrior-arms (raid 60 and any band on the new list).
- addon-data.json / Data.lua: both rotations.
- dps-per-point.json: warrior-fury, warrior-arms.
- Coach: pipeline.coachrules and persona starters for warrior-arms and warrior-fury.
- Fork golden: TestP1DPSWarrior, already regenerated in E 88b1d57dc.

## Measure at refresh

- MR: warrior-arms picks, weights and tiers on the new raid build. Reckless Swings and Spearing Strike change rage and the value of hit and weapon damage.
- MR: warrior-fury picks and weights on the new build (small).
- The six-spec note in the brief stands: the BiS refresh must carry the new Arms raidBuild and Fury build, not the 70753 strings.

## For other lanes

- Ranker/core: the `expertise_to_cap` block and expertise weight in bis/warrior-*.json. No Forever item has expertise (row 2d).

## For the guide lane (land/guides-r1 text)

- fury.md:78. Replace "The build also spends a point in **Piercing Howl**, which dazes everything within 10 yards for 6 seconds to stop runners and kite on trash, a point in **Lingering Rage**, which holds your rage 10 seconds longer after a fight so it carries into the next pull, and two in **Improved Charge**," with "The build puts two points in **Improved Berserker Rage**, so every Berserker Rage brings 10 rage, and two in **Improved Charge**,". The rest of the sentence stays.
- fury.md:90. Replace "Berserker Rage breaks fear and keeps you immune to it for 10 seconds: hold it for Magmadar's Panic, Onyxia's Bellowing Roar and Nefarian's fear, and Death Wish covers you the same way." with "Berserker Rage breaks fear and keeps you immune to it for 10 seconds. With Improved Berserker Rage it is also 10 rage, so press it on cooldown whenever you are under 80 rage, except on Magmadar, Onyxia and Nefarian, where you hold it for the fear. Death Wish covers you the same way."
- Fury rotation steps: add Berserker Rage after Reckless Swings. Suggested APL note for the new line: "Berserker Rage under 80 rage: Improved Berserker Rage makes it 10 rage. Hold it for the fear on Magmadar, Onyxia and Nefarian."
- arms.md:71. "**Anger Management** and **Improved Charge** (3 points)" is now true of the level 60 build only. Suggest adding: "The raid build skips Improved Charge."
- arms.md:79. Replace the whole paragraph with: "The raid build drops Iron Will, Deflection and Improved Charge. Two more points of **Improved Heroic Strike** make your rage dump cost 3 less. **Spearing Strike** is a second weapon strike on a 20 second cooldown that hits twice as hard on dragons and giants. Fifteen points in Fury reach **Reckless Swings**, four swings that bring 5 more rage each, and **Improved Execute**; **Enrage** is the way there, and it pays whenever raid damage lands on you."
- Arms rotation steps: Reckless Swings right after Mortal Strike; Spearing Strike after Slam, in Battle Stance only. Suggested APL notes: "Reckless Swings on cooldown after Mortal Strike: the next four swings each bring 5 more rage." and "Spearing Strike when Mortal Strike and Overpower are down; it hits twice as hard on Dragonkin and Giants."

## Still open

- WARRIOR-13 (Improved Shield Wall) for the owner (2c).
- Warcraft Logs blocked (gcloud login), not worked around. No level-60 warrior log was used.

## Wave 2

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| W2-1 | **Death Wish breaks fear.** True in Forever: 12328 e1 is aura 77 with mechanic 5 (fear immunity), and Attributes_1 has 0x8000, the same dispel-on-immunity bit Berserker Rage carries. Since 1.6.0 it is usable while feared and removes the fear (Vanilla WoW Wiki, https://vanilla-wow-archive.fandom.com/wiki/Death_Wish). **No sim effect:** no encounter casts a fear, the engine's Death Wish aura (talents.go:752) only applies +20% physical damage and +5% damage taken, and no rotation gates Death Wish on a fear. | SpellEffect 690204; SpellMisc 323746 | Aligned | - | - |
| W2-2 | **A crit swing gives double rage.** Kaivax, "Warrior Updates in Today's Beta Build" (Blizzard forums, 2026-10-02, https://us.forums.blizzard.com/en/wow/t/2369360, re-fetched 2026-11-03): "Additional Rage generated from landing Critical Strikes increased to 100% increased Rage (was 75%)." The engine matches: `sim/core/rage.go` normalizedSwingRage multiplies a landed white swing by `normalizedRageCritMultiplier = 2.0`, on main-hand and off-hand auto attacks only. Specials generate no swing rage, as in vanilla, and the crit check is `result.DidCrit()`. The rage is still per swing by weapon speed with no damage term, so the "crit size unverified" note in memory forever-rage-normalization-gap is settled. The 70755 client has no table row for this server rule; the blue post is the one source, and it is Blizzard's own. The off-hand 0.5 factor stays unverified (no public number). | Blizzard forums post; rage.go:120-170 | Aligned | - | - |

No code changed in wave 2, so no measurement was needed: the three warrior specs stand at the before/after table above.

## Wave 3

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| W3-1 | `TestBandBuildsTakeTheCastTalentsThatFit` warrior-fury level 30 went red after S c328e4192. The test's want had the old build written into it: Lingering Rage 1, Piercing Howl 1, Dual Wield Spec 3. The new walk still takes Death Wish; it puts those two points in Dual Wield Specialization instead (`-050530025000100000-`, tier 4 gate 5+5+3+2+5 = 20). **The test was stale, not the walk or the reservation.** Band 30 bare, horde, 10000 iterations: 60.1 → 62.2 at seed 1 and 60.0 → 62.1 at seed 50001 (+3.5%). Want updated, with a comment. The whole sim/leveling package is green. | ladder_first_walk_test.go:38 | CONFIRMED (test) | R15 c328e4192 | S (this commit) |
| W3-2 | The live 70755 push (wago.tools, fetched 2026-11-03 ~22:30, search totals about 3x the committed record) was layered over the shipped tables and compared with our layered copy for SpellEffect, SpellCooldowns, SpellPower, SpellShapeshift, SpellAuraOptions, SpellLevels, SpellMisc, TraitDefinitionEffectPoints, CurvePoint and TraitNodeEntry. **No warrior row differs.** The only differing curve is CurvePoint 251541-251545 (curve 85357, 2/4/6/8/10 → 3/6/9/12/15), which is druid **Rend and Tear** (1223246, TraitDefinition 134414). That is for the druid lane. 52 non-warrior SpellEffect rows also differ. | hotfix_push.fetch_push + layer_rows | Aligned (warrior); for other lanes | new-data | - |

## Wave 6

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| W6-1 | **Elune's Light outside Recklessness** (guide round 3 WARRIOR-R3-1, sim request 1). Client and live push agree; the push was re-fetched from wago.tools and layered, and it matches our copy row for row. Elune's Light 1259799: e0 aura 290, 10% crit, 15 s, 3 min. Recklessness 1719: aura 290, 100% crit, +20% damage taken, 15 s, 30 min. The engine autocasts Elune's Light on cooldown (`core/racials.go:313`). Recklessness adds 100% crit (`recklessness.go`), and a crit chance past 100% is just an always-crit, so overlapping the two wastes Elune's Light. Night Elf, Alliance, guide raid builds, raid 60, 10000 iterations, Elune's Light on cooldown against held for Recklessness: Arms 792.7 / 792.4 vs 784.8 / 784.4 (holding loses 1.0%); Fury 833.7 / 834.8 vs 825.1 / 825.9 (holding loses 1.0 to 1.1%). The guide's round-3 fix (on cooldown, never inside Recklessness) is right, and the engine already plays it that way. No sim change. | SpellEffect 1269038, SpellCooldowns 96706; SpellEffect 687178-80 | Aligned | - | - |
| W6-2 | `recklessness.go:10` described a 50%, 12 s, 5 min Recklessness; the code and client are 100%, 15 s, 30 min. Comment only, no behaviour change. | | Fixed (comment) | old | E 9ddc99572 |
| W6-3 | Stoneform for Dwarf Protection: left to the paladin lane, as instructed. | | - | - | - |

## Wave 7

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| W7-1 | **"The engine lets haste shorten the global cooldown"** (guide round 4, shaman.md SIM-R4-1, citing `sim/core/unit.go:374` SpellGCD and `constants.go:25` GCDMin). **Not what the engine does.** `Unit.SpellGCD` (max(1 s, 1.5 s × cast speed)) was called only by `Spell.EffectiveCastTime`, and nothing called that, in the fork or the site. The cast path that sets the global cooldown (`cast.go:255-301`) hastes CastTime alone; its own comment reads "Vanilla has no natural GCD reduction", and it holds the global for `Cast.EffectiveTime()` = max(GCD, CastTime), floored at 1 s only for a GCD a talent has already cut. No class code touches `cast.GCD` with a haste term; the only GCD writes are fixed values (Faerie Fire (Feral), forms, 1 s totems, Reckless Swings 1 s, Nature's Grace's -150 ms spell mod). Melee haste never reaches the GCD at all. A hasted Lightning Bolt under 1.5 s therefore already waits the full 1.5 s, so Berserking and Wind Blessed are not over-valued through the global. | cast.go:57-61, 255-301; grep of every `GCD` write | Aligned (claim refuted) | - | - |
| W7-2 | Client and Forever rule: on 70755 with the live push (wago SpellCategories `useHotfixes=1`), Lightning Bolt 15208, Frostbolt 10181, Whirlwind 1680, Bloodthirst 23894 and Slam 11605 are StartRecoveryCategory 133 with StartRecoveryTime 1500. The client has no "haste affects GCD" attribute; in Classic that is a server rule. Blizzard's Forever notes thread 2360696 (read 2026-11-04, latest update Oct 27) never mentions haste and the global cooldown. Blizzard's Hunter and Druid deep dive (2026-09-30, quoted by rankedboost.com/world-of-warcraft/forever/stats/haste/, updated 2026-11-02) names Nature's Grace as granting haste **and** "reduces your global cooldown". That is a named exception, which means baseline haste does not do it; the engine models it (druid/talents.go:116, 10% = 150 ms). Totems' 1 s GCD is the other named exception (finalboss.io). Verdict: Forever keeps Classic's rule, and the engine follows it. | wago SpellCategories; forum thread; deep dive | Aligned | - | - |
| W7-3 | The dead helpers misled the audit, so both are removed (`Unit.SpellGCD`, `Spell.EffectiveCastTime`), and `TestHasteNeverShortensTheGlobalCooldown` (sim/core/gcd_haste_test.go) pins the rule. **This touches shared core** but changes no behaviour: the fork builds clean (every package but sim/web, which needs binary_dist and was already broken) and the site sim builds. With no simulated number able to move, no 28-spec measure was run. | | Cleanup | - | E fbc5dbddf |

For the guide lane: the round-4 shaman edit ("Berserking inside Lightning Mastery buys nothing") stands. A cast hasted under 1.5 s gains nothing in the engine or in Forever, since the global holds 1.5 s. The SIM-R4-1 "engine bug" note in round4/shaman.md:74 should be withdrawn.
