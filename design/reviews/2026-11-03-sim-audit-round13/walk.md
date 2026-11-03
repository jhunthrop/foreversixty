# Sim-accuracy round 13: the leveling talent walk (lane walk)

Base: site `wt/r13-hybrids` (e0e7f1ed7), fork `lane/r13-hybrids` (50cb4ed67); no engine change in this lane.

## What the walk does, and why it left talents out

Below 60, a band's build is the guide's level-60 build cut down to level-9 points (`sim/leveling/talents.go`, `LadderTalentString`). The walk fills the spec's own tree top row down, then the other trees. One exception, the spine (78efc0d3e), takes the guide's deepest talent and its prerequisites first once the budget can hold them. A spec listed in `data/curated/ladder-first.json` reserves its own ranks first instead.

So a band whose points can just reach a talent on row N spends them on rows above it. Arms at 40 has 31 points: 30 fill the top six rows, and the 31st goes to Weaponmaster (row 6) before Mortal Strike (row 7). This is the walk's order, not a data error. The guide builds have not changed since the landing (8d0f78ea0 against now: the same own-tree digits for every spec below). Round 12's platform change (leveling strips in point-spend order) only shows `LadderPointOrder`, which follows the same walk. **Origin for every row: old.** The top-row-down truncation is 7e2ea0c7d (2026-09-28) and the spine is 78efc0d3e. Rounds 10 to 12 built the reservation for healers and Protection, and round 13 hybrids for the bear; none of them swept the other specs.

**One general rule ("reserve every talent the rotation casts as soon as it fits") is not correct.** Of 37 measured rows, 13 lose against the plain walk by more than their error. Examples: Swiftmend -4%, Conflagrate -4.7%, Mind Flay -3.7%, Arcane Blast -3.2%, Presence of Mind -2%. The fixes are therefore curated rows, each measured and each held to the adoption test.

## Method

- Probe (an uncommitted test in sim/leveling) built each candidate: the spec's curated list plus the talent and its prerequisite chain, put through the same `LadderTalentString`. It checked every tree with `treeIsLegal`: tier gates (5 points per row) and prerequisites at their listed rank.
- The tier gates line up with the learn levels: a row-t talent first fits at level 10+5t (row 2 at 20, row 3 at 25, row 4 at 30, row 6 at 40). Every reserved talent's client SpellLevels is at or below that: Conflagrate and Lacerate are learned at 25, Pyroblast at 20, Strider Kick, Mutilate, Hemorrhage and Siphon Life at 30, Mortal Strike at 40. So no reservation takes a talent before the client allows it.
- spec-breakdown on the bare preset at the band, both builds passed with `-talents`, 10,000 iterations, seeds 1 and 50001. Healers use `-heal` (effective HPS) and tanks `-tank` (threat/s). Horde for every row; both factions for the adopted rows.
- Adoption: at least 1% and beyond error at both seeds, both factions. The build must keep every talent the engine does not model, hold the tree, and drop nothing the rotation casts. A point that does nothing in the rotation may be replaced (the Shadow precedent).
- Band 38 is a ladder rung with no published BiS entry, so spec-breakdown cannot measure it. A 38 row follows its spec's 30/40 decision, and the same reservation applies.

## Adopted (CONFIRMED-BUG)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| W1 | **Arms band 40 had no Mortal Strike.** The rotation casts it; it is learned at 40, on the 30-point row behind Sweeping Strikes, and 31 points hold it. The walk spent the 31st point on Weaponmaster. Now Mortal Strike for Weaponmaster's point. | +20.3 / +20.4% horde, +20.5 / +20.4% alliance (seed 1 / 50001), error 0.17% | CONFIRMED-BUG | old (7e2ea0c7d) | S 340048375 |
| W2 | **Fury band 30 had no Death Wish** (20-point row, cast on cooldown). The walk spent the point on Enrage. Now Death Wish for Enrage's point. Bloodthirst, the walk's own capstone from 40, is listed after it so band 40 keeps it (first try dropped it, see X2). | +3.8 / +3.8% horde, +3.9 / +3.9% alliance | CONFIRMED-BUG | old | S 340048375 |
| W3 | **Survival band 30 had no Strider Kick** (20-point row, learned at 30, cast). The walk spent the point on Predator's Edge's fifth rank. Now Strider Kick for that rank. Lacerating Strikes is listed after it, so band 40 keeps it. | +9.8 / +9.7% horde, +10.6 / +10.4% alliance | CONFIRMED-BUG | old | S 340048375 |
| W4 | **Fire band 20 had no Pyroblast** (10-point row, learned at 20, the opener). The walk spent the point on Ignite's fifth rank. Now Pyroblast for that rank. | +2.0 / +1.7% horde, +2.1 / +2.0% alliance, error 0.16% | CONFIRMED-BUG | old | S 340048375 |
| W5 | **Assassination band 30 (and rung 38) had no Mutilate** (20-point row, learned at 30, the builder). The walk spent the point on Vile Poisons' second rank. Now Mutilate for that rank. Cold Blood as well measured +0.6% more (under 1%), so it is not reserved. Holding Seal Fate as well (the spine's talent at 36-39) cost 0.5% at band 40, so it is not reserved either; band 40 is unchanged. | +8.1 / +7.9% horde, +8.7 / +8.6% alliance; Seal Fate variant at 40: -0.6 / -0.5% | CONFIRMED-BUG | old | S 340048375 |
| W6 | **Subtlety bands 20 and 30 had no Ghostly Strike or Hemorrhage.** The rotation casts both: Ghostly Strike (10-point row), Premeditation (15-point row) and Hemorrhage (20-point row, learned at 30, behind Serrated Blades 3). The plain swap would drop Elusiveness, which the engine models only through Vanish, never cast, so a probe reads it as zero. It would also drop Premeditation, which the opener casts. The row holds Elusiveness 1, Ghostly Strike, Premeditation, Serrated Blades 3 and Hemorrhage. So band 20 takes Ghostly Strike for one Opportunity rank, and band 30 takes Hemorrhage for one Initiative rank. | Band 20: +2.6 / +2.6% horde, +2.7 / +2.6% alliance. Band 30: +6.7 / +6.8% horde, +6.5 / +6.6% alliance. (Plain Hemorrhage swap: +6.3%, but it drops Premeditation; Ghostly Strike for Elusiveness: +2.7%, but it drops an unmodeled talent.) | CONFIRMED-BUG | old | S 340048375 |
| W7 | **Affliction band 30 had no Siphon Life** (20-point row, learned at 30, kept up). The walk spent the point on Pandemic. Now Siphon Life for Pandemic's point. Wrack is listed after it, so band 40 keeps it. The guide already says "Siphon Life from 30". | +1.37 / +1.27% horde, +1.19 / +1.25% alliance, error 0.07% | CONFIRMED-BUG | old | S 340048375 |
| W8 | **Demonology band 30 had no Soul Link** (20-point row, behind Demonic Sacrifice; cast before the pull). The walk spent the point on Decimation's second rank, which does nothing before Soul Fire at 48. Now Soul Link for that rank. Only Soul Link is reserved. Reserving Demonic Sacrifice too would make band 20 take it for Demonic Aegis, which the engine does not model (`_ = warlock.Talents.DemonicAegis`). The rotation never casts the sacrifice without Demonic Pact, and the walk reaches it on its own by 30. | +2.4 / +2.3% horde, +2.3 / +2.3% alliance; band 20 Demonic Sacrifice for Aegis: 0.00% | CONFIRMED-BUG | old | S 340048375 |
| W9 | **Restoration shaman band 30 had no Mana Tide Totem** (15-point row, cast). Now Mana Tide Totem for Tidal Mastery's point; appended to the existing row (after Mindfulness 2 and Water Shield). | Effective HPS +1.93 / +1.93% horde, +1.92 / +2.03% alliance; mana lasts 179 to 200 s | CONFIRMED-BUG | old | S 340048375 |

Tests first: `TestBandBuildsTakeTheCastTalentsThatFit` and `TestDemonologyKeepsDemonicAegisAtTwenty` (sim/leveling/ladder_first_walk_test.go) failed on all ten band strings and now pass. `TestCuratedLadderFirstIsSound` now knows the eight new specs and checks one more thing: at a published band a reservation must not lose the walk's spine capstone. That check found X2 and now passes. `TestRestorationShamanHoldsWaterShieldAtTwenty` moves its band-30 string to the Mana Tide build. Level 60 is the guide build for every spec (the soundness test).

## Held (measured, not adopted)

| # | band row | build change | horde seed 1 / 50001 | why held |
|---|---|---|---|---|
| H1 | Balance 30 Eclipse | Vengeance 2 to 1 | -0.86 / -0.84% | loses |
| H2 | Restoration druid 30 / 40 Swiftmend | Gift of Nature 1 / Improved Rejuvenation 1 | -3.9 / -3.9%, -4.1 / -4.0% | loses (healers: see For other lanes) |
| H3 | Arcane 20 Arcane Blast | Arcane Concentration 5 to 4 | -3.2 / -3.2% | loses |
| H4 | Arcane 30 Presence of Mind | Arcane Meditation 3 to 2 | -2.0 / -2.1% | loses |
| H5 | Frost 20 Ice Lance | Ice Shards 1 | -0.54 / -0.54% | loses |
| H6 | Holy paladin 20 Divine Favor | Purifying Power 1 | 0.00 / 0.00% | cast twice a fight, but band 20's effective healing and mana are identical: its crits only overheal |
| H7 | Holy paladin 30 Holy Shock | Infusion of Light 1 | -0.38 / -0.38% | loses |
| H8 | Protection paladin 20 Templar's Bulwark | Sacred Duty 2 to 1 | (hybrids T2: threat -2%) | open with the owner (hybrids T2) |
| H9 | Retribution 20 Seal of Command | Sanctified Judgement 1 | -0.55 / -0.43% | loses |
| H10 | Shadow 20 Mind Flay | Spirit Tap 2 to 1 | -3.7 / -3.8% | loses (the guide's questing order already takes Mind Flay at 20; see For other lanes) |
| H11 | Shadow 20 to 50 Inner Focus | - | not measured | Inner Focus is in Discipline. At 20 to 40 it would cost Shadow points the guide's leveling order spends first ("every point in Shadow gets you to Shadowform at 40"). At 50 the 10 Discipline points sit on rows 0 and 1, so it needs an 11th Discipline point taken from Shadow. Not a walk defect. |
| H12 | Assassination 30 Cold Blood (alone) | Vile Poisons 2 to 1 | +0.09 / +0.10% | under 1% |
| H13 | Elemental 30 Lava Burst | Elemental Warding 3 to 2 | +0.47 / +0.57% | under 1% |
| H14 | Restoration shaman 30 Nature's Swiftness (alone / with Mana Tide) | Tidal Mastery / Healing Focus | -0.68%; with Mana Tide +1.8 / +1.7%, lower than Mana Tide alone | loses |
| H15 | Restoration shaman 40 Nature's Swiftness, Riptide, combinations | Mana Tide / Restorative Totems | -0.62%, +0.46%, -1.19%, -0.83%, -0.77% | loses or under 1% |
| H16 | Affliction 20 Amplify Curse | Improved Corruption 4 to 3 | -0.46 / -0.39% | loses |
| H17 | Demonology 20 Demonic Sacrifice | Demonic Aegis 1 | 0.00 / 0.00% | not cast without Pact; drops an unmodeled talent |
| H18 | Destruction 20 Shadowburn | Cataclysm 1 | +0.55 / +0.73% | under 1% |
| H19 | Destruction 30 Conflagrate (for Shadowburn / for Ruin 5 to 4) | | -4.8 / -4.7%, -4.1 / -4.1% | loses: at 30 it eats the Immolate it needs |
| H20 | Protection warrior 20 Last Stand | Anticipation 4 to 3 | threat +0.13 / +0.26%, damage taken +1.3% | under 1% |
| H21 | Bear 30 / 40, Protection paladin 30, Restoration shaman 20 Water Shield | - | - | already fixed (hybrids 7cbe92a23, round 12) |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **The reservation replaces the spine.** `LadderTalentString` takes the curated path instead of `spendOwnTree`. So a new row silently lost the capstone the plain walk had: Bloodthirst at Fury 40, Wrack at Affliction 40 to 43, Lacerating Strikes at Survival 40 to 43. `TestCuratedLadderFirstIsSound` now fails a row that drops the spine's capstone at a published band. The three rows list the capstone after their talent. Merging the spine into every reservation instead was tried, but it moved the round-12 Restoration shaman band 40 build (Riptide for Restorative Totems, unmeasured), so it was not kept. | soundness test | guard added (the defect never shipped) | - | S 340048375 |
| X2 | **Two older rows lose the capstone between published bands.** Discipline at 35 drops Divine Aegis. Restoration shaman at 46 to 49 drops Riptide, which the plain walk holds from 46. This affects only the addon's point order and the strips between bands, not a published band. | soundness check at every level | OPEN (healers lane: measure or accept) | R10/R11/R12 rows (the round 10 to 12 healprofile reservations) | - |
| X3 | **Trainables lane, item 1** (bear Lacerate in the addon order): with hybrids' bear row, `LadderPointOrder` now spends Lacerate at level 25 and Primal Bite at 40 (probe on this branch). The committed `addon-builds.json` and the addon-data default builds are stale until they are regenerated. | probe of `LadderPointOrder(..., ForSpec("druid-feral-bear"))` | Aligned once regenerated | - | - |
| X4 | **classes.json, Lacerate copy.** It read "Lacerate is in as a baseline Bear bleed learned after level 30". The client and Wowhead Forever (1328617, "Talent", "Requires level 25", fetched 2026-11-03) make it a Feral talent on the 15-point row, learned at 25. Fixed, with the Wowhead source added. The build copy needs a regen: `tests/test_curated.py::test_the_committed_build_matches_the_curated_facts` fails until then. | data/curated/classes.json | CONFIRMED-BUG | new-data (1e0a8ca39 transcribed the demo; 70707 moved it) | S 0bec04102 |

## Before/after (bare, the band's own fight, 10,000 iterations)

| spec, band | horde s1 | horde s50001 | alliance s1 | alliance s50001 | horde s1 absolute |
|---|---|---|---|---|---|
| warrior-arms 40 | +20.34% | +20.42% | +20.48% | +20.40% | 105.7 to 127.3 DPS |
| hunter-survival 30 | +9.75% | +9.66% | +10.55% | +10.42% | 80.4 to 88.2 |
| rogue-assassination 30 | +8.05% | +7.87% | +8.74% | +8.58% | 58.4 to 63.0 |
| rogue-subtlety 30 | +6.72% | +6.76% | +6.47% | +6.59% | 60.1 to 64.1 |
| warrior-fury 30 | +3.80% | +3.82% | +3.86% | +3.86% | 59.9 to 62.2 |
| rogue-subtlety 20 | +2.57% | +2.55% | +2.67% | +2.56% | 37.5 to 38.5 |
| warlock-demonology 30 | +2.40% | +2.28% | +2.32% | +2.33% | 82.0 to 83.9 |
| mage-fire 20 | +1.96% | +1.69% | +2.14% | +2.01% | 34.5 to 35.2 |
| shaman-restoration 30 (effective HPS) | +1.93% | +1.93% | +1.92% | +2.03% | 87.9 to 89.6 HPS |
| warlock-affliction 30 | +1.37% | +1.27% | +1.19% | +1.25% | 89.9 to 91.1 |

Every other band of every spec is unchanged: a dump of all 23 walked specs' builds at levels 10 to 60, before and after, differs only on the rows above and on the levels between them, including rungs 38 for Assassination and Restoration shaman. Raid 60 is unchanged by construction.

## Files to regenerate

- Ladder goldens (sim/request TestRotationLadder): hunter-survival, mage-fire, rogue-assassination, rogue-subtlety, shaman-restoration, warlock-affliction, warlock-demonology, warrior-arms, warrior-fury (plus hybrids' druid-feral-bear).
- `data/builds/1.60.1.70753/addon-builds.json` (`TestCommittedAddonBuildsFollowTheGuides` fails until then), and the addon-data default builds and Data.lua that follow it.
- BiS band talents at refresh (`data/builds/1.60.1.70753/bis/<spec>.json`/.md `talents` for the bands above), and dps-per-point for those bands.
- `data/builds/1.60.1.70753/classes.json` from data/curated/classes.json (X4).

## Measure at refresh

- The BiS ranker's band entries for the ten adopted rows: gear re-ranked on the new builds. These numbers are on the current published gear.

## For other lanes

- **Healers:** H2 (Swiftmend at 30/40 loses 4%: check how the rotation uses it, since it eats a HoT and the walk's Gift of Nature / Improved Rejuvenation point is worth more), H6 (Divine Favor at 20 only overheals), X2 (Discipline 35, Restoration shaman 46-49 capstone gap).
- **Warlock:** H19, Conflagrate at 30 loses 4.7% against Shadowburn or a Ruin rank. Check the Immolate consumption timing at low level.
- **Guides (prose that disagrees with the band builds the strips will show):**
  - shaman/elemental.md Leveling: "Lightning Bolt and Flame Shock carry you until Lava Burst around 30". Band 30's build does not take Lava Burst; the walk reaches it later (H13, under 1%).
  - paladin/holy.md Leveling: "Holy Shock comes at 30 ... if you level as Holy". Band 30's build does not take it (H7).
  - warlock/destruction.md Leveling: "Conflagrate can eat your Immolate" before Shadow and Flame 5, which reads as if Conflagrate were in the 30s build. It is not (H19).
  - priest/shadow.md already says its questing order differs from the strips (Mind Flay at 20). No change.
  - warlock/affliction.md ("Siphon Life from 30"), warrior/arms.md ("Mortal Strike" from 40), rogue/assassination.md ("Mutilate comes at 30") and rogue/subtlety.md ("Hemorrhage only arrives at 30") now match the band builds. Before this fix the band builds, not the prose, were wrong.
- **Coordinator:** H8 (hybrids T2) stays the owner's call.

## Still open

1. X2: the capstone gaps between bands in the Discipline and Restoration shaman rows.
2. Band 38 cannot be measured with spec-breakdown (no BiS entry); its builds follow the 30/40 decisions.
