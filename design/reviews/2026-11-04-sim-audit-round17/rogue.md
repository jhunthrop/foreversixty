# Sim-accuracy round 17: rogue

Lane `rogue`. Site wt/r17-rogue (from int/r17 d7fc51fe7), fork lane/r17-rogue (from int/r17 f637ec32c). Client 1.60.1.70755 with the late hotfix re-layer. Specs: Assassination, Combat, Subtlety.

Method: spec-breakdown on the guide raid builds (`$R/rt`) and frozen BiS (`$R/bis-frozen`); 10000 iterations at seeds 1 and 50001 for every reported number, 3000 to 5000 to explore. Levels 10 to 19 (no BiS band) were run through a scratch request test on the ladder's gear. The lane's starting point reproduces `$R/before.tsv` for all twelve rogue rows.

Commits: site 600c52c0f (Backstab), 7cb1b23fc (Subtlety build), 080b12930 (notes and prose from the review), plus this report; fork 4fd849b92 and a037f143f (APL copy syncs, no engine code).

Totals: **2 CONFIRMED-BUGs fixed**: 1 `old` (Assassination and Subtlety never backstab before their talent builder), 1 `70755 3d4712981` (Subtlety's dead Improved Ambush point). No engine code changed.

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| K1 | Launch logs for Death Mark storage (white swings, poisons) and the Mutilate/Shank off-hand penalty | foreverlogs.gg, 2026-11-04 08:00 local: one report since round 16, #2814 ("Test", uploaded 2026-11-04 12:24 UTC, zero encounters); #2815 to #3000 do not exist. api.foreversixty.gg/v1/reports/recent lists only a 2026-10-20 report. Death Mark and Mutilate are learned at 40 and Shank at 24, so no launch-day log can carry them. Warcraft Logs stays blocked (gcloud login expired), not worked around | OPEN (no log exists) | - | - |
| K2 | Combat's Eviscerate at four: can the rotation test for a dagger? | **No APL value reads the weapon type** (fork `proto/apl.proto` APLValue: nothing on weapons; `spellCanCast(Backstab)` folds in energy, `spellIsKnown` is level-blind because the site rewrites ranks per level). **And the weapon is not the split anyway:** every published Combat set, bare 20 to 60 and raid, holds a main-hand dagger and builds with Backstab (37 to 41 Backstabs a fight at every band). Eviscerate at four, 10000 its, seeds 1/50001: raid 60 +0.46/+0.55%, bare 60 +0.79/+0.86%, bare 50 +0.05/+0.05%, bare 40 -0.24/-0.16%, bare 30 -0.43/-0.29%, bare 20 -0.48/-0.48%. Ruthlessness is not the driver either (raid without it: +0.47%). The split follows gear and level (Backstab crit 33% at bare 30 to 64% raid). Only `currentAttackPower` could proxy it, with a magic threshold; web consensus (warcrafttavern, icy-veins, mobalytics Forever combat guides, searched 2026-11-04) is Eviscerate at five. Not adopted | OPEN (owner taste call: an attack-power gate, or keep five) | - | - |
| K3 | Regression: Subtlety's Garrote after Ghostly Strike and the Hemorrhage hold | Raid 60 horde, 3000 its: Garrote 8.85 casts against 9.67 Ghostly Strikes, Hemorrhage 33.8, 1329326 uptime 21% | Aligned | - | - |
| K4 | Regression: Death Mark at four | Death Mark 5.32 casts at four points and 5.46 at five a fight, never held | Aligned | - | - |
| K5 | Rank rows level by level (`-level`, floor, +3, +6, +9, bands 20 to 50, all three specs) | No unresolved actions at any of the 48 levels; every rank's learn level in the engine (Eviscerate, Sinister Strike, Backstab, Rupture, Garrote, Slice and Dice, Ambush, Expose Armor, Shank, Mutilate, Death Mark, both poisons) matches `trainables/rogue.json` 70755. The rogue has no rank-choice rows. Eviscerate 31016 and Backstab 25300 at 60 are the pending level-60 book ranks (not reopened) | Aligned | - | - |

## Assassination

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | **No Backstab before Mutilate.** From 20 to 39 the rotation built with Sinister Strike only, though the published bare 20 and 30 sets hold a dagger in the main hand (Blackfang, Toxic Revenger) and the fight is from behind. Backstab is 150% weapon damage plus a flat bonus for 60 energy; Combat's rotation already backstabs at every level. Now Backstab builds until Mutilate whenever it can be cast, and Sinister Strike is the no-dagger fallback held to 60 energy (Combat's line) | Bare 20 39.0 to 42.8 (+9.7%), level 29 44.7 to 50.3 (+12.5%), bare 30 66.0 to 70.4 (+6.7%), level 39 77.2 to 81.5 (+5.6%), seeds 1/50001 alike. Test `TestDaggerLevelingBuildsWithBackstab` (sim/request/rogue_rotation_test.go): 0 Backstabs, 45 Sinister Strikes at 25 before; passes after. assassination.md:131 already told the player to Backstab from behind on mobs | CONFIRMED-BUG | old (the Sinister Strike-only gate from the 2026-10 curation; guide-accuracy and sim rounds checked level 60 and the sub-20 card, not this) | S 600c52c0f, 080b12930, E 4fd849b92, a037f143f |
| A2 | Re-audit: raid build point by point against the 70755 rank texts (Malice, Ruthlessness 20%/rank, Improved Slice and Dice 15%/rank, Vigor, Blood Splatter, Lethality 10%/rank on Mutilate, Vile Poisons 4%/rank, Cold Blood, Improved Poisons 2%/rank, Relentless Strikes 4%/rank, Venom, Seal Fate 33/66/100%, Mutilate, Coup de Grace 7/13/20%, Lightning Reflexes 2 (tier gate), Puncturing Wounds, Precision, Opportunity +20% Mutilate); poisons 20%/30% proc (spelltips 11340, 11356); raid breakdown (Mutilate 40.3 casts, both hands, Eviscerate at 4/5, Rupture with Blood Splatter, 1.43 Cold Bloods) | talents/rogue.json 70755; engine `client_values.go`, `talents.go`, `poisons.go` | Aligned | - | - |

## Combat

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | Eviscerate at four | See K2 | OPEN | - | - |
| C2 | Raid build dead points: Lightning Reflexes 4 | Lightning Reflexes 3 with Hack and Slash 5 is illegal (Flawless Execution needs 15 points above its row; the build has exactly 15 with LR 4), so the fourth point is a gate | Aligned | - | - |
| C3 | Re-audit: Puncturing Wounds (30% Backstab crit, 45% extra point), Hack and Slash, Flash of Steel (3 energy, 0.4 s a rank), Aggression, Dual Wield Specialization 25%/rank, Weapon Expertise, Composure, Flawless Execution, Coup de Grace, Blade Flurry, Adrenaline Rush against the 70755 rank texts; raid breakdown (Backstab 40.6, Eviscerate at five 9.9, Slice and Dice at 1-2, Shank 22.5) | talents/rogue.json 70755; engine | Aligned | - | - |

## Subtlety

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Improved Ambush held a point that does nothing.** The 70755 build (`035131031321001151`) put its last tier-two point in Improved Ambush (+15% Ambush crit). The rotation never casts Ambush (Garrote opener, then a Garrote after every Ghostly Strike), and the level 60 sets hold a sword and a mace, so Ambush cannot be pressed at all. Dirty Deeds' second point sits in the same tier (14082: "reduces the Energy cost of Cheap Shot and Garrote by 20"), keeps the 30 points that open Boot Knife and Death Mark, and takes 10 more energy off about nine Garrotes a fight. Build now `005302011/320003/035230031321001151` | Raid 60: horde 746.0 to 760.1 (+1.89/+1.89%), alliance 751.3 to 765.4 (+1.88/+2.01%); bare 60 +2.05/+2.08%, 50 +1.82/+1.87%, 40 +2.70/+2.63%, 30 +1.95/+1.71%, 20 +0.00/+0.24%. Passes even the 1% adoption bar, keeps every cast talent and the tree. The walk now takes Dirty Deeds at 18 and 19 and Camouflage from 24. Test `TestSubtletyBuildSpendsNoPointOnAmbush` failed first; `ladder_first_walk_test.go`'s level-30 want follows the build | CONFIRMED-BUG (build) | 70755 3d4712981 (the new-tree build; R16 f89d5bc98's Garrotes made the point worth more) | S 7cb1b23fc |
| S2 | **No Backstab before Hemorrhage** (same as A1). Subtlety's leveling build takes Opportunity (+20% Backstab) from 13 and the guide says to keep a dagger and Backstab until Hemorrhage, but the rotation cast Sinister Strike only | Levels 10 to 19 on the ladder's daggers: +7.5% (10), +14.9% (13), +9.6% (16), +8.8% (19). Bare 20 with a dagger main hand (Blackfang, Diamond Hammer off hand): +14.7% (20), +13.3% (25), +17.7% (29). The published band-20 set holds a mace in the main hand (Diamond Hammer, ranked under the old rotation), where the 60-energy fallback costs -0.2 to -0.6%; the dagger set beats it by +5 to +10%, so the re-rank should move it (Measure at refresh) | CONFIRMED-BUG (same fix as A1, counted once) | old | S 600c52c0f, 080b12930, E 4fd849b92, a037f143f |
| S3 | SIM-R11-1 (guide-accuracy rounds 11 to 13): "Subtlety below 20 never spends combo points" | It spends them all on Slice and Dice (11 casts a fight at 15). Below Rupture, Slice and Dice at three or more takes every set before a fifth point builds: an Eviscerate-at-five line placed above it never fires (levels 12, 16, 19 identical). Nothing to change | Aligned | - | - |
| S4 | Re-audit: Ghostly Strike, Hemorrhage, Death Mark, Premeditation, Preparation, Initiative, Serrated Blades, Boot Knife, Dirty Deeds, Opportunity against 70755; fillers Master of Deception, Camouflage, Elusiveness are tier gates (tiers 0-5 hold exactly 30) | talents/rogue.json 70755 | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Occult Poison I/II (458821, 1214171; 136 Nature over 12 s plus 4% magic damage taken) sit in `trainables/rogue.json` as `class_spell` at 54 | `class_spell` is the not-learnable bucket (data/pipeline/trainables.py:17-27); neither id has a SkillLineAbility row in 70755 (wago, fetched 2026-11-04). Likely a Season of Discovery leftover. An in-game check of the 54+ poison vendor or Wowhead Forever item search settles it | OPEN (low) | - | - |

## Before/after

Raid 60, guide raid builds, frozen BiS (relative %, seeds 1 / 50001; absolute at seed 1):

| spec | alliance | horde | seed 1 absolute |
|---|---|---|---|
| rogue-assassination | 0 / 0 | 0 / 0 | A 722.9, H 726.2 |
| rogue-combat | 0 / 0 | 0 / 0 | A 774.0, H 776.6 |
| rogue-subtlety | +1.88 / +2.01 | +1.89 / +1.89 | A 751.3 to 765.4, H 746.0 to 760.1 |

Bare bands (horde, seeds 1/50001): Assassination bare 20 +9.7%, bare 30 +6.7% (Backstab); Subtlety bare 60 +2.1%, 50 +1.8%, 40 +2.7%, 30 +1.8% (build), bare 20 -0.2/-0.5% on the published mace set (see S2).

No commit moves DPS for any spec outside the rogue class.

## Files to regenerate

- Ladder goldens `sim/request/testdata/ladder/rogue-assassination.golden.md` and `rogue-subtlety.golden.md` (Backstab before Mutilate/Hemorrhage; Subtlety's talent walk: Dirty Deeds 2 at 18-19, Camouflage from 24, no Improved Ambush).
- `addon-data.json` (Assassination and Subtlety rotation lines and notes; the band 10 to 30 cards), then Data.lua.
- `addon-builds.json` (Subtlety build), then Data.lua; coach starters if they carry the Subtlety build.
- `dps-per-point.json` for Subtlety (every band) and Assassination (bands 20, 30), then Data.lua.
- apl-sync: already hand-synced both sides (site `sim/request/apl/`, fork `ui/rogue/apls/forever_*.apl.json`).
- Fork goldens: none.

## Measure at refresh

- Subtlety band 20 weapons: with Backstab in the rotation a dagger main hand should replace Diamond Hammer (+5 to +10% at 20 to 29).
- Assassination bands 20 and 30 weights (Backstab builds now).
- Subtlety talents in every BiS file and its weights (Dirty Deeds 2, no Improved Ambush; Garrote cheaper).

## For other lanes

- None.

## For the guide lane

- Nothing left over: this lane wrote the text (assassination.md levels 20 to 60, subtlety.md talents paragraph, priority, levels 1 to 19 and 20 to 60, and the card notes). wow-player verdict below.

## Rendered review

wow-player on the rendered Assassination and Subtlety pages (astro dev on this worktree after `npm run sync`, 1440 px full page). First pass: both NOT SHIP. MUSTs: the Sinister Strike rule differed between prose ("when you are not behind") and card notes ("no dagger"); "held to 60 energy" read as a typo; the Backstab note left out "from behind"; Assassination's "your only builder" contradicted its own Sinister Strike sentence. All fixed (S 080b12930), with the SHOULDs: Dirty Deeds leads with the opener and Cheap Shot, "Improved Ambush is skipped", "Camouflage, from 24", "level 60 rotation" in the Overview, a paragraph break in Assassination's 20 to 60. Second pass: **Subtlety SHIP, Assassination SHIP.** Combat's own Sinister Strike note still says "held to 60 energy" (approved before, not touched here; the guide lane may align it). The rendered rotation cards show the old rows until addon-data is regenerated.

## Still open

- K1: Death Mark storage and the Mutilate/Shank off-hand penalty, until a level 24+ (Shank) or 40+ (Death Mark, Mutilate) Forever log exists.
- K2/C1: Combat Eviscerate at four (owner taste call: attack-power gate or keep five).
- X1: Occult Poison learnability.
- Carried: S4 of round 16 (Hemorrhage's dagger hand).
