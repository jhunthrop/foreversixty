# Sim-accuracy round 9: ranks lane (owner ruling 3)

Ruling: rotations may downrank. When a fight is mana-bound, the rotation casts the rank that gives the most damage or healing over the fight. Deliberate band-60 downranks are kept, never undone by a top-rank rule.

Site branch `wt/r9-ranks`. No engine change. Measured with spec-breakdown, horde, 10,000 iterations, seeds 1 and 50001, published preset per band (bare 20 to 60, raid 60), engine 40021374a. "Before" is origin/main 55a1212f7. Healers use `-heal` (effective HPS on data/curated/heal-profile.json); paladin-protection uses `-tank` (threat/s).

## Mechanism

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| M1 | **The rank rewrite undid the deliberate downranks below 60.** `sim/request/rotation_ranks.go` moved every ranked spell to the top rank learned at the band. So below 60 Holy and Discipline's Renew rank 3, Discipline's Power Word: Shield rank 4, and Holy paladin's party Flash of Light rank 4 and Holy Shock rank 1 were raised to the band's top rank. | Failing first: `TestRewriteKeepsTheAuthoredDownranksBelowSixty`. At band 50, Renew 6075 came out as 10928 and Shield 3747 as 10899. | CONFIRMED-BUG | S 5c63f6719 |
| M2 | **New default rule.** `spellranks.LearnedSpellID` (renamed from HighestLearnedSpellID) gives the highest rank learned that is **not above the authored rank**. A rotation written at the top rank behaves exactly as before. A deliberate lower rank keeps its rank, and below that rank's learn level the rank under it is used. Ids for an Ahn'Qiraj book rank cap at the top trainer rank. | `TestLearnedUpToKeepsAnAuthoredDownrank`. `sim/internal/spellranks/spellranks.go` (learnedIn, authoredTier). | - | S 5c63f6719 |
| M3 | **Per-band rank choice.** `data/curated/rank-choices.json` (spec, then authored id, then rows of `{from_level, to_level, rank}`) is copied to `sim/request/rank-choices.json` by `make rank-choices-sync`. The table is parsed once per process. The rewrite (`rankResolver`) takes a row's rank where one covers the level; otherwise M2 applies. The site's /sim, the ranker and the nightly (all through `request.rotation`), spec-breakdown `-rotation` and the ladder harness's zero-casts rule all call the same resolver. Level 60 still runs the curated rotation as written. | `TestRewriteHonoursAPerBandRankChoice`, `TestThePublishedRotationTakesTheCuratedRankChoices`, `TestCuratedRankChoicesAreCastableDownranks` (every row names a spec, an id its rotation casts, the right spell, a band below 60, a rank learned at from_level and a rank lower than the default), `TestEmbeddedRankChoicesMatchCurated`. | - | S 5c63f6719, 5f1cbfacf |
| M4 | **A probe can pin a rank.** spec-breakdown `-rotation` used to raise the top rank back before the probe ran, so a probe below 60 could not test a lower rank (warlock lane's report). Under M2, a probe that names a lower rank keeps it. | `sim/cmd/spec-breakdown/main.go` loadRotation | - | S 5c63f6719 |
| M5 | **The addon now casts the sim's ranks.** `data/pipeline/addonrotation.py` follows M2 and M3 below 60: it never raises above the list's own rank, and a band's row wins. It is built on the healers lane's a65bd8111, cherry-picked here as 547d3f60f. | `test_a_deliberate_lower_rank_is_never_lifted_below_sixty`, `test_a_band_casts_the_rank_choice_the_curated_table_names`. The test asserting Renew is lifted at 50 is replaced. | CONFIRMED-BUG (same as M1, player-facing) | S 134d5065f |

Why per-band data and not a search at run time: /sim stays fast (one map lookup per ranked id). Each choice is a reviewed row with evidence. The ladder and the addon read the same answer.

Flat or behind-the-pace? For Elemental at 50, rank 4 cast all fight gave 152.0. Rank 4 cast only while mana is behind the fight's pace gave 150.6 (+38.2% vs +39.4%). Pacing at x0.8 and x1.2 was no better. The casters lane found the same. A flat per-band rank needs no new rotation lines.

## Adoption rule used

- **Damage specs:** adopt a lower rank where the fight runs dry and the gain is at least 1% at both seeds, beyond the stated error.
- **Healers:** the same, with the healer rank rule's 2% bar (design/reviews/2026-10-08-healer-mana-audit.md). A healer row must not turn a fight that runs dry into one that ends with mana unspent (integrator's check).
- **Method:** every lower learned rank of every ranked spell a rotation names, swapped in alone at 3,000 iterations (`scratchpad r9/ranks/sweep.py`). Then lower ranks were combined, adding one at a time while each added at least 1% (`greedy.py`). Then the final set was confirmed at 10,000 iterations on both seeds.
- **Ranks learned below 20:** these were allowed, because the engine applies the sub-20 coefficient penalty (`sim/core/spell.go` levelPenalty).
- **Every healer row is measured on the level-60 heal profile.** The healers lane notes that this profile is the level-60 Phase 1 fight at every band. Below 60, nearly nothing overheals, so these rows are pure mana efficiency. The owner may want to rule on whether the profile should scale below 60.

## Adopted per-band choices (rank-choices.json)

| spec | levels | change | gain seed 1 / 50001 | absolute, seed 1 | mana, s (before to after) |
|---|---|---|---|---|---|
| shaman-elemental | 30-39 | Lightning Bolt rank 4 (915), Flame Shock rank 1 (8050) | +10.17% / +9.94% | 46.2 to 50.9 DPS | OOM 90.2 to 60.4 |
| shaman-elemental | 40-49 | Lightning Bolt rank 4, Flame Shock rank 1 | +18.51% / +18.80% | 83.2 to 98.6 | OOM 97.2 to 41.6 |
| shaman-elemental | 50-59 | Lightning Bolt rank 4, Lava Burst rank 2 (1238299), Flame Shock rank 3 (8053) | +46.37% / +46.69% | 108.9 to 159.4 | OOM 94.0 to 5.3 |
| priest-shadow | 30-39 | Shadow Word: Pain rank 3 (970), Devouring Plague rank 1 (2944) | +4.39% / +4.39% | 56.9 to 59.4 | OOM 91.9 to 92.8 |
| priest-shadow | 40-49 | Devouring Plague rank 1, Mind Blast rank 4 (8104) | +4.70% / +4.86% | 125.6 to 131.5 | OOM 70.8 to 58.7 |
| priest-shadow | 50-59 | Devouring Plague rank 1, Mind Blast rank 2 (8102), Shadow Word: Death rank 2 (1309633) | +7.45% / +7.39% | 189.2 to 203.3 | OOM 67.4 to 40.6 |
| druid-balance | 30-39 | Wrath rank 4 (5179), Moonfire rank 3 (8925), Insect Swarm rank 1 (5570) | +13.89% / +13.89% | 65.5 to 74.6 | OOM 48.0 to 9.7 |
| mage-fire | 40-49 | Scorch rank 1 (2948) | +5.71% / +5.71% | 147.0 to 155.4 | OOM 55.6 to 23.0 |
| mage-fire | 50-59 | Scorch rank 1 | +6.20% / +6.19% | 182.2 to 193.5 | OOM 58.2 to 19.4 |
| mage-arcane | 30-39 | Frostbolt rank 4 (7322) | +1.46% / +1.46% (on the casters lane's Evocation branch: +1.21% / +1.45%) | 75.6 to 76.7 | OOM 47.5 to 30.6 |
| paladin-holy | 30-39 | Flash of Light rank 1 (19750) on the tank | +22.13% / +22.13% | 70.5 to 86.1 HPS | lasts 28.4 to 28.7 |
| paladin-holy | 40-49 | Flash of Light rank 1 on the tank, Holy Light rank 3 (647) | +50.92% / +51.08% | 124.7 to 188.2 | lasts 36.2 to 118.5 |
| paladin-holy | 50-59 | Flash of Light rank 1 on the tank, Holy Light rank 3 | +60.25% / +60.22% | 152.2 to 243.9 | lasts 44.3 to 210.2 |
| shaman-restoration | 20-29 | Healing Wave rank 3 (547) | +8.60% / +8.60% | 46.5 to 50.5 | lasts 31.9 to 50.1 |
| shaman-restoration | 30-39 | Healing Wave rank 3 | +13.62% / +13.62% | 85.9 to 97.6 | lasts 44.3 to 125.8 |
| shaman-restoration | 40-49 | Healing Wave rank 3 | +19.46% / +19.46% | 144.4 to 172.5 | lasts 39.3 to 249.3 |
| shaman-restoration | 50-59 | Healing Wave rank 3 | +10.41% / +10.36% | 225.8 to 249.3 | lasts 48.6 to 62.4 |
| priest-discipline | 40-49 | Binding Heal rank 1 (401937), Renew rank 1 (139) | +13.09% / +13.02% (vs main; Renew's own part +2.85%) | 143.6 to 162.4 | lasts 47.6 to 57.8 |
| priest-discipline | 50-59 | Binding Heal rank 1 | +7.95% / +7.90% | 206.2 to 222.6 | lasts 49.6 to 62.5 |
| druid-restoration | 20-29 | Rejuvenation rank 1 (774) | +2.54% / +2.54% | 39.4 to 40.4 | lasts 253.4 to 252.8 |
| druid-restoration | 30-39 | Rejuvenation rank 2 (1058) | +10.26% / +10.16% | 98.4 to 108.5 | lasts 199.8 to 223.1 |

From the default rule alone (M2, no row):
- **Holy priest:** Renew rank 3 kept. +0.77/+0.62% at 30, +1.76/+1.68% at 40, +5.72/+5.60% at 50.
- **Discipline at 30:** Renew rank 3 and Shield rank 4 kept, +0.78/+0.78%.
- **Holy paladin's party Flash of Light rank 4:** also kept, and part of its rows' gains.

Not adopted:
- **Holy paladin, band 40, Holy Light rank 2:** +53.9%, but mana lasts 453 s of a 300 s fight, so it ends with mana unspent. Rank 3 together with the tank's Flash of Light rank 1 is the row above.
- **Discipline, band 30, Renew rank 1:** +1.8%, under the 2% bar.
- **Paladin-protection, Divine Protection or Divine Shield rank 1 (+1.9 to +3.7% threat):** a defensive cooldown, not a mana choice. See "For other lanes".
- **Paladin-protection and Retribution, Consecration rank 1 (+1.3 to +1.6%):** Retribution is out of mana for 1.3 s, so the fight is not mana-bound. See "For other lanes".

## Every spec and band (published rotation, before to after)

Specs without mana are not swept: rogues, warriors, Feral and Guardian druid. Their rotations name only top-rank ids and have no row, so the resolver returns exactly what it did before. Every band 60 is unchanged, because the rotation runs as written there. Destruction keeps `spellIsKnown(412758)` (Incinerate rank 1) as written below 60, and its numbers did not move.

Columns: published metric seed 1 (DPS; HPS for healers; threat/s for paladin-protection), gain at seed 1 / 50001, and mana seconds. Mana seconds are out-of-mana seconds for damage specs and "mana lasts" for healers. Paladin-protection reports none.

| spec | band | metric, seed 1 | change | mana, s |
|---|---|---|---|---|
| druid-balance | bare 20 | 25.8 to 25.8 | +0.00% / +0.00% | 20.8 to 20.8 |
| druid-balance | bare 30 | 65.5 to 74.6 | +13.89% / +13.89% | 48.0 to 9.7 |
| druid-balance | bare 40 | 156.1 to 156.1 | +0.00% / +0.00% | 5.8 to 5.8 |
| druid-balance | bare 50 | 198.8 to 198.8 | +0.00% / +0.00% | 7.9 to 7.9 |
| druid-balance | bare 60 | 289.2 to 289.2 | +0.00% / +0.00% | 0.5 to 0.5 |
| druid-balance | raid 60 | 579.5 to 579.5 | +0.00% / +0.00% | 0.0 to 0.0 |
| druid-restoration | bare 20 | 39.4 to 40.4 | +2.54% / +2.54% | 253.4 to 252.8 |
| druid-restoration | bare 30 | 98.4 to 108.5 | +10.26% / +10.16% | 199.8 to 223.1 |
| druid-restoration | bare 40 | 176.5 to 176.5 | +0.00% / +0.00% | 213.2 to 213.2 |
| druid-restoration | bare 50 | 271.0 to 271.0 | +0.00% / +0.00% | 214.6 to 214.6 |
| druid-restoration | bare 60 | 393.3 to 393.3 | +0.00% / +0.00% | 222.4 to 222.4 |
| druid-restoration | raid 60 | 777.3 to 777.3 | +0.00% / +0.00% | 287.9 to 287.9 |
| hunter-beast-mastery | bare 20 | 105.6 to 105.6 | +0.00% / +0.00% | 103.8 to 103.8 |
| hunter-beast-mastery | bare 30 | 136.3 to 136.3 | +0.00% / +0.00% | 117.5 to 117.5 |
| hunter-beast-mastery | bare 40 | 190.4 to 190.4 | +0.00% / +0.00% | 105.4 to 105.4 |
| hunter-beast-mastery | bare 50 | 242.5 to 242.5 | +0.00% / +0.00% | 110.5 to 110.5 |
| hunter-beast-mastery | bare 60 | 309.3 to 309.3 | +0.00% / +0.00% | 96.8 to 96.8 |
| hunter-beast-mastery | raid 60 | 661.8 to 661.8 | +0.00% / +0.00% | 0.0 to 0.0 |
| hunter-marksmanship | bare 20 | 103.4 to 103.4 | +0.00% / +0.00% | 25.8 to 25.8 |
| hunter-marksmanship | bare 30 | 138.4 to 138.4 | +0.00% / +0.00% | 15.7 to 15.7 |
| hunter-marksmanship | bare 40 | 184.0 to 184.0 | +0.00% / +0.00% | 35.7 to 35.7 |
| hunter-marksmanship | bare 50 | 242.8 to 242.8 | +0.00% / +0.00% | 37.7 to 37.7 |
| hunter-marksmanship | bare 60 | 305.2 to 305.2 | +0.00% / +0.00% | 46.3 to 46.3 |
| hunter-marksmanship | raid 60 | 653.1 to 653.1 | +0.00% / +0.00% | 0.0 to 0.0 |
| hunter-survival | bare 20 | 69.9 to 69.9 | +0.00% / +0.00% | 0.1 to 0.1 |
| hunter-survival | bare 30 | 103.9 to 103.9 | +0.00% / +0.00% | 0.0 to 0.0 |
| hunter-survival | bare 40 | 173.3 to 173.3 | +0.00% / +0.00% | 9.9 to 9.9 |
| hunter-survival | bare 50 | 222.8 to 222.8 | +0.00% / +0.00% | 0.0 to 0.0 |
| hunter-survival | bare 60 | 316.8 to 316.8 | +0.00% / +0.00% | 0.0 to 0.0 |
| hunter-survival | raid 60 | 759.5 to 759.5 | +0.00% / +0.00% | 0.0 to 0.0 |
| mage-arcane | bare 20 | 36.5 to 36.5 | +0.00% / +0.00% | 71.7 to 71.7 |
| mage-arcane | bare 30 | 75.6 to 76.7 | +1.46% / +1.46% | 47.5 to 30.6 |
| mage-arcane | bare 40 | 155.5 to 155.5 | +0.00% / +0.00% | 16.1 to 16.1 |
| mage-arcane | bare 50 | 196.5 to 196.5 | +0.00% / +0.00% | 24.5 to 24.5 |
| mage-arcane | bare 60 | 319.8 to 319.8 | +0.00% / +0.00% | 25.7 to 25.7 |
| mage-arcane | raid 60 | 601.4 to 601.4 | +0.00% / +0.00% | 1.0 to 1.0 |
| mage-fire | bare 20 | 35.2 to 35.2 | +0.00% / +0.00% | 68.4 to 68.4 |
| mage-fire | bare 30 | 70.9 to 70.9 | +0.00% / +0.00% | 75.9 to 75.9 |
| mage-fire | bare 40 | 147.0 to 155.4 | +5.71% / +5.71% | 55.6 to 23.0 |
| mage-fire | bare 50 | 182.2 to 193.5 | +6.20% / +6.19% | 58.2 to 19.4 |
| mage-fire | bare 60 | 342.2 to 342.2 | +0.00% / +0.00% | 21.1 to 21.1 |
| mage-fire | raid 60 | 670.8 to 670.8 | +0.00% / +0.00% | 0.0 to 0.0 |
| mage-frost | bare 20 | 36.7 to 36.7 | +0.00% / +0.00% | 99.6 to 99.6 |
| mage-frost | bare 30 | 62.7 to 62.7 | +0.00% / +0.00% | 75.2 to 75.2 |
| mage-frost | bare 40 | 144.8 to 144.8 | +0.00% / +0.00% | 60.3 to 60.3 |
| mage-frost | bare 50 | 193.8 to 193.8 | +0.00% / +0.00% | 64.4 to 64.4 |
| mage-frost | bare 60 | 299.3 to 299.3 | +0.00% / +0.00% | 50.2 to 50.2 |
| mage-frost | raid 60 | 601.5 to 601.5 | +0.00% / +0.00% | 0.0 to 0.0 |
| paladin-holy | bare 20 | 30.1 to 30.1 | +0.00% / +0.00% | 43.1 to 43.1 |
| paladin-holy | bare 30 | 70.5 to 86.1 | +22.13% / +22.13% | 28.4 to 28.7 |
| paladin-holy | bare 40 | 124.7 to 188.2 | +50.92% / +51.08% | 36.2 to 118.5 |
| paladin-holy | bare 50 | 152.2 to 243.9 | +60.25% / +60.22% | 44.3 to 210.2 |
| paladin-holy | bare 60 | 220.1 to 220.1 | +0.00% / +0.00% | 59.8 to 59.8 |
| paladin-holy | raid 60 | 648.2 to 648.2 | +0.00% / +0.00% | 258.4 to 258.4 |
| paladin-protection | bare 20 | 41.8 to 41.8 | +0.00% / +0.00% | None to None |
| paladin-protection | bare 30 | 68.9 to 68.9 | +0.00% / +0.00% | None to None |
| paladin-protection | bare 40 | 200.6 to 200.6 | +0.00% / +0.00% | None to None |
| paladin-protection | bare 50 | 292.8 to 292.8 | +0.00% / +0.00% | None to None |
| paladin-protection | bare 60 | 428.3 to 428.3 | +0.00% / +0.00% | None to None |
| paladin-protection | raid 60 | 1018.2 to 1018.2 | +0.00% / +0.00% | None to None |
| paladin-retribution | bare 20 | 41.1 to 41.1 | +0.00% / +0.00% | 0.1 to 0.1 |
| paladin-retribution | bare 30 | 73.3 to 73.3 | +0.00% / +0.00% | 0.0 to 0.0 |
| paladin-retribution | bare 40 | 108.2 to 108.2 | +0.00% / +0.00% | 0.0 to 0.0 |
| paladin-retribution | bare 50 | 179.3 to 179.3 | +0.00% / +0.00% | 1.3 to 1.3 |
| paladin-retribution | bare 60 | 259.8 to 259.8 | +0.00% / +0.00% | 0.0 to 0.0 |
| paladin-retribution | raid 60 | 658.8 to 658.8 | +0.00% / +0.00% | 0.0 to 0.0 |
| priest-discipline | bare 20 | 29.5 to 29.5 | +0.00% / +0.00% | 29.0 to 29.0 |
| priest-discipline | bare 30 | 89.3 to 90.0 | +0.78% / +0.78% | 43.5 to 44.5 |
| priest-discipline | bare 40 | 143.6 to 162.4 | +13.09% / +13.02% | 47.6 to 57.8 |
| priest-discipline | bare 50 | 206.2 to 222.6 | +7.95% / +7.90% | 49.6 to 62.5 |
| priest-discipline | bare 60 | 288.4 to 288.4 | +0.00% / +0.00% | 53.9 to 53.9 |
| priest-discipline | raid 60 | 620.0 to 620.0 | +0.00% / +0.00% | 112.7 to 112.7 |
| priest-holy | bare 20 | 30.5 to 30.5 | +0.00% / +0.00% | 24.4 to 24.4 |
| priest-holy | bare 30 | 64.7 to 65.2 | +0.77% / +0.62% | 26.1 to 26.3 |
| priest-holy | bare 40 | 119.3 to 121.4 | +1.76% / +1.68% | 29.4 to 30.5 |
| priest-holy | bare 50 | 173.2 to 183.1 | +5.72% / +5.60% | 38.5 to 40.3 |
| priest-holy | bare 60 | 309.2 to 309.2 | +0.00% / +0.00% | 50.9 to 50.9 |
| priest-holy | raid 60 | 674.1 to 674.1 | +0.00% / +0.00% | 135.0 to 135.0 |
| priest-shadow | bare 20 | 35.6 to 35.6 | +0.00% / +0.00% | 90.2 to 90.2 |
| priest-shadow | bare 30 | 56.9 to 59.4 | +4.39% / +4.39% | 91.9 to 92.8 |
| priest-shadow | bare 40 | 125.6 to 131.5 | +4.70% / +4.86% | 70.8 to 58.7 |
| priest-shadow | bare 50 | 189.2 to 203.3 | +7.45% / +7.39% | 67.4 to 40.6 |
| priest-shadow | bare 60 | 349.4 to 349.4 | +0.00% / +0.00% | 19.0 to 19.0 |
| priest-shadow | raid 60 | 724.2 to 724.2 | +0.00% / +0.00% | 0.0 to 0.0 |
| shaman-elemental | bare 20 | 22.5 to 22.5 | +0.00% / +0.00% | 92.6 to 92.6 |
| shaman-elemental | bare 30 | 46.2 to 50.9 | +10.17% / +9.94% | 90.2 to 60.4 |
| shaman-elemental | bare 40 | 83.2 to 98.6 | +18.51% / +18.80% | 97.2 to 41.6 |
| shaman-elemental | bare 50 | 108.9 to 159.4 | +46.37% / +46.69% | 94.0 to 5.3 |
| shaman-elemental | bare 60 | 143.8 to 143.8 | +0.00% / +0.00% | 83.7 to 83.7 |
| shaman-elemental | raid 60 | 476.8 to 476.8 | +0.00% / +0.00% | 3.9 to 3.9 |
| shaman-enhancement | bare 20 | 33.8 to 33.8 | +0.00% / +0.00% | 0.0 to 0.0 |
| shaman-enhancement | bare 30 | 69.0 to 69.0 | +0.00% / +0.00% | 0.0 to 0.0 |
| shaman-enhancement | bare 40 | 123.9 to 123.9 | +0.00% / +0.00% | 0.4 to 0.4 |
| shaman-enhancement | bare 50 | 160.3 to 160.3 | +0.00% / +0.00% | 0.0 to 0.0 |
| shaman-enhancement | bare 60 | 217.3 to 217.3 | +0.00% / +0.00% | 0.0 to 0.0 |
| shaman-enhancement | raid 60 | 594.0 to 594.0 | +0.00% / +0.00% | 0.0 to 0.0 |
| shaman-restoration | bare 20 | 46.5 to 50.5 | +8.60% / +8.60% | 31.9 to 50.1 |
| shaman-restoration | bare 30 | 85.9 to 97.6 | +13.62% / +13.62% | 44.3 to 125.8 |
| shaman-restoration | bare 40 | 144.4 to 172.5 | +19.46% / +19.46% | 39.3 to 249.3 |
| shaman-restoration | bare 50 | 225.8 to 249.3 | +10.41% / +10.36% | 48.6 to 62.4 |
| shaman-restoration | bare 60 | 337.6 to 337.6 | +0.00% / +0.00% | 76.2 to 76.2 |
| shaman-restoration | raid 60 | 838.9 to 838.9 | +0.00% / +0.00% | 263.3 to 263.3 |
| warlock-affliction | bare 20 | 54.5 to 54.5 | +0.00% / +0.00% | 0.0 to 0.0 |
| warlock-affliction | bare 30 | 88.2 to 88.2 | +0.00% / +0.00% | 0.0 to 0.0 |
| warlock-affliction | bare 40 | 135.1 to 135.1 | +0.00% / +0.00% | 0.0 to 0.0 |
| warlock-affliction | bare 50 | 196.6 to 196.6 | +0.00% / +0.00% | 0.0 to 0.0 |
| warlock-affliction | bare 60 | 291.4 to 291.4 | +0.00% / +0.00% | 0.7 to 0.7 |
| warlock-affliction | raid 60 | 585.2 to 585.2 | +0.00% / +0.00% | 0.0 to 0.0 |
| warlock-demonology | bare 20 | 48.2 to 48.2 | +0.00% / +0.00% | 0.0 to 0.0 |
| warlock-demonology | bare 30 | 80.5 to 80.5 | +0.00% / +0.00% | 0.0 to 0.0 |
| warlock-demonology | bare 40 | 142.0 to 142.0 | +0.00% / +0.00% | 0.0 to 0.0 |
| warlock-demonology | bare 50 | 208.8 to 208.8 | +0.00% / +0.00% | 1.8 to 1.8 |
| warlock-demonology | bare 60 | 326.1 to 326.1 | +0.00% / +0.00% | 3.9 to 3.9 |
| warlock-demonology | raid 60 | 654.9 to 654.9 | +0.00% / +0.00% | 0.7 to 0.7 |
| warlock-destruction | bare 20 | 53.0 to 53.0 | +0.00% / +0.00% | 3.1 to 3.1 |
| warlock-destruction | bare 30 | 89.3 to 89.3 | +0.00% / +0.00% | 2.8 to 2.8 |
| warlock-destruction | bare 40 | 141.2 to 141.2 | +0.00% / +0.00% | 3.5 to 3.5 |
| warlock-destruction | bare 50 | 239.0 to 239.0 | +0.00% / +0.00% | 2.0 to 2.0 |
| warlock-destruction | bare 60 | 358.5 to 358.5 | +0.00% / +0.00% | 4.5 to 4.5 |
| warlock-destruction | raid 60 | 649.7 to 649.7 | +0.00% / +0.00% | 0.6 to 0.6 |

Hunters (Beast Mastery and Marksmanship run dry for 15 to 117 s on bare bands) gain nothing from any lower rank: their damage is mostly auto shot and pet, and every lower Arcane Shot, Serpent Sting or Aimed Shot rank came within ±0.4%. Enhancement, Retribution, warlocks (Life Tap) and Frost mage gain nothing either. Frost mage's +4 to +11% at 30 to 40 disappears once the casters lane's Evocation line is in.

## Band 60 (the rotation as written; not changed here)

A flat lower rank at 60 changes both presets. These candidates are for the APL owners:

| spec | change | raid 60 | bare 60 | note |
|---|---|---|---|---|
| priest-holy | Flash Heal rank 2 (9472) | +1.63% / +1.53% | -8.44% / -8.44% | under the healer 2% bar; bare loses |
| priest-discipline | Heal rank 1 (2054) | +1.24% / +1.19% | -3.22% / -3.22% | under the bar; bare loses |
| paladin-holy | tank Flash of Light rank 1 | +1.40% / +1.39% | +35.21% / +35.27% | raid under the 2% bar |
| shaman-elemental | Lightning Bolt rank 9 all fight | -0.02% / -0.04% | +5.49% / +5.42% | casters lane: rank 9 **behind the pace** is +2.0 to +2.2% raid on their Clearcasting fix (E ff4bae666) |
| druid-restoration | party Rejuvenation rank 4 (2090) | +3.2 / +3.3% (healers lane's table) | - | needs the party's casts split from the tank's in the list |
| shaman-elemental, paladin-holy, shaman-restoration | Lightning Bolt rank 4 / Holy Light rank 3 / Healing Wave rank 3 | (raid not mana-bound) | +37% / +38% / +10% (3,000 iterations) | bare 60 only; would need a preset-aware choice or a pace line |

## For other lanes

- **Integrator:** every TestRotationLadder golden moves: the rules header text (resolver renamed) and the downranked bands of the specs above. Regenerate at the pin. 547d3f60f is the healers lane's a65bd8111 cherry-picked unchanged; 134d5065f builds on it.
- **Healers / integrator:** the notes of priest-holy, priest-discipline and paladin-holy still say "The site runs this list as written at band 60 and lifts every ranked spell to the highest rank learned below it". That is no longer true below 60. The notes are long single lines, so this lane left them alone to keep the merge clean. Suggested replacement: "Below 60 a ranked spell casts the highest rank learned that is not above the rank named here, or the band's row in data/curated/rank-choices.json."
- **Healers:** the band-60 raid rows above (Restoration druid +3.2% clears the 2% bar). Whether the heal profile should scale below 60 is a question for the owner.
- **Casters:** Elemental raid 60 Lightning Bolt rank 9 behind the pace (+2.0 to +2.2% on your engine) is a new rotation line at 60, yours to add if it holds after the pin.
- **Hybrids / tank:** paladin-protection gains threat when it casts Divine Protection or Divine Shield rank 1 instead of the top rank (bare 20 to 50, +1.9 to +3.7%). A shorter defensive should not raise threat; check what the longer one blocks. Consecration rank 1 beats the top rank on a single target for Retribution 50 (+1.5%) and Protection 40 to 60 (+1.3 to +3.1%) without a mana shortage; check the Consecration line's value on one target.

## Still open

- Band-60 rank choices that differ by preset (bare 60 versus raid 60) cannot be expressed per level. The headline at 60 is raid (`PublishedPreset`), so this lane left 60 as written.
- A choice keyed by authored id cannot split one spell's tank casts from its party casts when both lines name the same id. Today the rotations give them distinct ids where it matters (Holy paladin's Flash of Light).
