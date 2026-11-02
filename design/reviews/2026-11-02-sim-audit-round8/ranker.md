# Sim audit round 8: ranker (2026-11-02)

This lane owns `sim/cmd/leveling-bis` and the BiS pipeline code. It read the ranker and its tests end to end and checked every rule it applies against the game and the owner's class rules. It ran the ranker for all 28 specs at every band (20, 30, 40, 50, 60; raid and bare at 60), both factions, and audited every pick for legality. It made no engine change.

**Setup:**
- Site: `wt/r8-ranker` from origin/main 170b7b84c. Engine: cd4dfdcc0 (the pin), untouched.
- Runs (`leveling-bis -bands 20,30,40,50,60`, every spec; output in `scratchpad/r8/ranker/`, not committed):
  - `before`: origin/main on origin/main's committed `simdb.bin`, as the brief asks.
  - `beforeheld`: origin/main on that simdb with S 1bd3920b5's held-item typing applied (`scratchpad/r7/ranker/simdb-held.bin`, the round-7 patch).
  - `after`: this branch on the held simdb.
- Why the held simdb: the committed `simdb.bin` predates 1bd3920b5, so an orb or tome still replaces the main hand in every sim. Round 7's one-hand fix therefore cannot show on it. The committed-simdb run keeps Hand of Edward the Odd with an empty off hand for Balance, Shadow and Elemental at 60, because every set wearing a held item measures a main hand short. The next data run fixes this by itself.
- Both re-rank runs read the published `bis/*.json` as incumbents, so the guard (`keepVerifyBest`) is live in both.

**In short: 5 CONFIRMED-BUGs, all fixed with tests first** (K1, K2, K3, K7 on today's data; K4 latent until the PvP relabel lands). Plus two robustness fixes for the relabel (K5, K6), with eligibility unchanged.

## Rules the ranker applies, checked

| # | rule | where | verdict |
|---|---|---|---|
| R1 | Class may equip the item: the per-class item file, plus a Go weapon-subclass safety net | `data.go`, `weapon_requirements.go` classWeaponSubclasses | **CONFIRMED-BUG**: paladin axes (K1) |
| R2 | Armor by level: mail from 40 (hunter, shaman), plate from 40 (warrior, paladin) | `eligible.go` armorAvailableLevel | Aligned with forever-class-rules. No sub-40 mail or plate in any run. |
| R3 | Required level, raised to a quest's or recipe's floor | `eligible.go`, `applyEffectiveRequiredLevels` | Aligned; no pick above the band in any run |
| R4 | Faction: the item's restriction beats a mined rep Side | `eligible.go`, `band.go` sourceObtainable | Aligned; no wrong-faction pick |
| R5 | Phase gates: raids, world bosses, Dire Maul, ZG, AQ, BG rep, Darkspear and the late recipes are all closed at launch; honor ranks are open | `content-phases.json`, `contentphases.go`, `sourceObtainable` (Opens) | Aligned; no raid, MC, Onyxia or legendary source anywhere. Raids open 2026-12-09, with phase raids-1 gated today. |
| R6 | Rep standing: friendly and honored while leveling, revered and exalted only at 60 | `repStandingObtainable` | Aligned |
| R7 | Rare spawns give way to any rival source | `withoutRareSpawnDrops` | Aligned |
| R8 | PvP rank above 10: an alternative only, never a default pick | `pvpRankCap`, `excludeAbovePvpRankCap` | Aligned; inclusion untouched (owner ruling pending). See PvP below. |
| R9 | Rings and trinkets: no same id or same name in both slots | `excludePaired` | Aligned. Stricter than the game: a non-unique ring can be worn twice. The `Unique` flag is read and never used. This is a design choice, not a bug. |
| R10 | Two-hander means an empty off hand, and a one-hander gets an off hand | `pick`, `enforceTwoHandOffHandInvariant`, `fillEmptyOffHand` | Aligned on the held simdb: no 2H plus off hand and no 1H beside an empty off hand. On the committed simdb, 9 caster bands show 1H with no off hand (the simdb artifact above). |
| R11 | Dual wield only for rogues, Fury, and hunters (from 20); shamans never | `leveling.DualWieldSpecs` | Aligned with the owner rule. No shaman off-hand weapon, no held item on a dual-wielder, no weapon on a non-dual-wielder. |
| R12 | Main-hand-only weapons never go in the off hand | `offHandWeapons`, `MainHandOnly` | Aligned |
| R13 | Daggers when the rotation casts Backstab, Ambush or Mutilate | `aplRotationRequiresDagger` | Assassination and Subtlety are aligned. Combat is caught too, because its APL casts Backstab with a Sinister Strike fallback, so the rule is broader than it needs to be. Measured at 60: daggers 210.4 ± 0.2, the best sword or mace main hand 202.0 ± 0.2. No pick changes. Left as is (Still open). |
| R14 | Ranged slot by class: casters wand; hunters bow, gun or crossbow (no thrown); warriors and rogues those plus thrown; paladin, druid and shaman a relic | `restrictRangedByProficiency` | Aligned; no wrong ranged type in any run |
| R15 | Shields only for warrior, paladin and shaman | per-class file | Aligned |
| R16 | Set bonuses: completion trials, cross-class sets excluded | `sets.go`, `crossClassSetItem` | Aligned on the picks |
| R17 | Weights: a sweep per band, a retry, a lower-band fallback, the primary anchor | `main.go` | **CONFIRMED-BUG**: measured on a naked character (K2) |
| R18 | Weapon DPS at 14 AP per DPS, in every weapon slot | `score.go` | The off hand's 50% damage is not modelled. Measured on Fury at 60: halving the off-hand DPS term does not change the pick (Greenhammer 270.4 against Mirah's Song 266.6). Aligned in outcome. |
| R19 | Trinkets: an engine tournament, then an empty slot when the gain is noise | `trinkets.go`, `report.go` trinketLowGain | **CONFIRMED-BUG**: emptied a stat trinket (K3) |
| R20 | The incumbent guard: the published set stays unless the new one beats it | `incumbent.go` | Works, and it is what exposed K2: 56 band rows kept the published set on origin/main (held simdb) |
| R21 | Survival raid build: the raid pass reads `raidBuild:` and falls back to the leveling build when that line is absent | `raid_build.go`, `leveling.GuideRaidBuildTalents` | Aligned. The physical lane's change takes effect with no ranker change (tested by `TestPassTargetsWithoutRaidBuildReuseTheLevelingBuild`). |

## Ranker

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| K1 | **Paladins could not wear axes.** In 2026-09-30, the ranker's static weapon table and the pipeline's proficiency table both took paladins' one- and two-hand axes away. The reason given was a persona claim that "paladins train maces, swords and polearms". Classic paladins do train both axes. | - Client 1.60.1.70707 `SkillLineAbility`: spell 196 (One-Handed Axes) has ClassMask 79 and spell 197 (Two-Handed Axes) has 71. Both include paladin (2). `SkillRaceClassInfo` rows 335 (mask 66) and 122 (mask 6) agree.<br>- Wowhead Classic trainer data for Buliwyf Stonehand (Ironforge, npc 11865, fetched 2026-11-01) teaches both to reqclass 71: warrior, paladin, hunter, shaman. The same data gives druids no polearm in practice and rogues no axe, so those two exclusions stay.<br>- Icy Veins' Classic Retribution pre-raid list and Wowhead's Classic Retribution guide put Arcanite Reaper, a two-hand axe, in the weapon slot.<br>- Tests first: `weapon_requirements_test.go` (paladin has axes; a druid polearm is still dropped) and `test_proficiency.py` (paladin can equip 0 and 1; rogue axe and druid polearm stay vestigial). Both failed, then passed. | CONFIRMED-BUG (three sources) | S 188718042 |
| K2 | **Every band ranked by weights measured on a naked character.** The weights sweep wore only a weapon. A bare character's small mana pool overvalues Intellect for every mana-using melee and hunter spec, and with no gear, hit and crit are priced off an empty stat sheet. So Survival band 40 picked Thinking Cap, Archmagister's Faceted Pendant and Windchaser Amice: cloth caster pieces on a melee hunter. The ranker's own final set then lost to the previously published set beyond error in many bands. The guard kept the old set each time; Survival band 40 was 170.9 ± 0.1 against 166.8 ± 0.1. | Swept on the published band sets at 400 iterations a direction (`zz_probe_test.go`, scratch only). Figures are attack power per point, bare then geared:<br>- Retribution 60: Intellect 6.71 → 3.36, hit 38.9 → 34.7<br>- Retribution 40: agility 1.69 → 1.37<br>- Survival 40: Intellect 3.30 → 1.37<br>- Enhancement 60: Intellect 4.83 → 2.29, hit 35.2 → not significant<br>- Fury 60: hit 38.2 → 21.5, crit 34.7 → 25.2, agility 1.73 → 1.26<br>- Arcane 60: haste not significant → 3.05<br>Fix: sweep the ladder character, pick a first-pass set with those weights, sweep again wearing it (`gearedBandWeights`), and rank and publish the geared weights. Two guards keep it honest:<br>- When the geared weights' first pass measures worse beyond error than the ladder's, the band keeps the ladder sweep (S e73208467). 24 band passes took this branch, among them Fire 30/50/60, Balance 60, Arcane 60, Beast Mastery 20/50, Fury 60, the healers' 50/60 and Destruction 30/60.<br>- The ladder weights' set, wearing the final trinkets, goes through the same verify swaps and set completion as the band's own set, and the guard measures it alongside (S c4c68b405, S 7a6027c84). It won 20 band rows, for example Arms band 60 Horde (250.0 against 246.2) and Demonology band 60.<br>Tests first in `geared_weights_test.go`: the band publishes the geared sweep, an untrusted geared sweep is ignored, a ladder set that measures better wins, and the contender wears the final trinkets. Each failed, then passed. | CONFIRMED-BUG | S 249b83fed, S e73208467, S c4c68b405, S 7a6027c84 |
| K3 | **A stat trinket was published as an empty slot.** Retribution raid band 60 (Horde) published trinket1 empty with `no_dps_value`. Its pick was Rune of the Guard Captain (+42 attack power, +7 hit, no effect), and it measured 4.6 DPS above Frozen Heart of the Mountain in the same set. The raid preset's noise (a gain of a few DPS against an error near 6) kept the gain under `trinketGainSignificant`, so the report gate emptied the slot. The gate is meant for a trinket worth nothing to the spec, such as Rune of Perfection on an Arms warrior. | `before/paladin-retribution.json` band 60 raid horde. The alternatives list shows Frozen Heart −4.62 and Uther's Strength −11.44 against the emptied pick. Fix: a pick whose score is above zero, so the band's significant weights value its stats, is never emptied for a noisy gain. Test first: `TestBuildReportKeepsAStatTrinketTheWeightsValueWhenItsGainIsNoisy` failed, then passed. Rune of Perfection's test still empties. | CONFIRMED-BUG | S 69a6adf3e |
| K7 | **The faction reconciliation put one trinket in both slots.** In this lane's re-rank, Discipline band 60 (Alliance) published Ankh of Life as trinket1 and trinket2. Horde's trinket1 pick was Ankh of Life, and Alliance already wore it in trinket2. `reconcileTrinketDirection` measured it as the better trinket1 and adopted it without looking at the pair mate. The rule (pick's excludePaired) forbids one item in both slots. | `after/priest-discipline.log`: "trinket1: alliance adopted Ankh of Life (id 1713) from horde". The existing bug was exposed by K2's different picks; no published file has it today. Fix: never adopt an item the target already wears in the other trinket slot (`wouldDuplicatePairMate`). Test first: `TestReconcileFactionTrinketsNeverAdoptsTheTargetsOtherTrinket` failed, then passed. The band-60 re-run is clean: Shard of the Splithooves and Uther's Strength, 282.4 HPS, as on origin/main. | CONFIRMED-BUG | S 39296e4e3 |
| K4 | **"PvP rank 0".** A pvp source with no rank was labelled "PvP rank 0 · Horde", and one with no faction got a dangling " · ". The owner's ruling makes Forever's PvP rewards rank-free, so the pipeline's relabelled rows would show this. | `band.go` pvpSourceLabel. Now "PvP reward · Horde", and empty parts are left out. Tests first in `band_test.go`. Eligibility is unchanged. | CONFIRMED-BUG (latent: no current row has rank 0) | S 9947f008d |
| K5 | **A relabelled PvP reward could vanish silently.** `dropNonRankRewards` strips the pvp source of every item classic-db gives no honor rank. An unranked Forever PvP reward that is also a classic-db item would lose its only source with no word. | `classicranks.go`. The ranker now names such items in a warning (`droppedToNoSource`). Eligibility is unchanged. Test first: `classicranks_test.go`. There are none on today's loot.json. | Robustness | S 9947f008d |
| K6 | **The rank cap must survive the relabel.** The pipeline relabels 491 Forever PvP items as unranked "PvP reward" and keeps the Era rank in `title_rank`. Without a change, every one of them would become a default pick. That is a policy change the owner has not made. | The integrator's request. `pvpRankExceedsCap` now reads max(rank, title_rank). A vendor row inherits title_rank, and dropNonRankRewards clears it. Test first: `TestAnUnrankedPvpSourceCarryingATitleRankKeepsTheRankCap` (title_rank 12 stays out of the defaults; title_rank 8 stays in). | Robustness (keeps eligibility) | S 19f27af30 |

## Legality audit of every pick

`scratchpad/r8/ranker/audit.py` checked every slot of every band, faction and preset in all three runs and in the published files. It looked for:
- the wrong inventory slot;
- an item above the band's level;
- the wrong faction;
- armor the class cannot wear, or sub-40 mail or plate;
- a weapon subclass the class cannot use;
- the wrong ranged type for the class;
- a raid source;
- a two-hander plus an off hand, or a one-hander with no off hand;
- a weapon on a non-dual-wielder, or a held item or shield on a dual-wielder;
- a main-hand-only weapon in the off hand, or an off-hand-only weapon in the main hand;
- a shield on a class without shields;
- a non-dagger for Assassination or Subtlety;
- a duplicate id or name in a pair.

Results:
- **Committed simdb (`before`): no violation** in 28 specs × 24 band rows, apart from the held-item artifact (9 caster bands at 60 with a one-hander and no off hand) and K3.
- **Held simdb (`beforeheld`, `after`): no violation**, except K7's duplicate trinket in the first `after` run, now fixed. Three Horde band-20 caster necks (Frost, Elemental, Affliction) are empty with `no_dps_value` on both sides: only 3 of 18 eligible necks are sourced, and Horde's have no caster stat.
- Empty slots that remain are all honest: trinkets at bands 20 to 40 where no trinket is sourced, a relic whose effect the engine does not model, and two-hander off hands.
- Two faction rows looked wrong at first and are fine: the Horde-only Rune of the Guard Captain, and Horde- and Alliance-only quest rings, sit on their own faction.

## PvP and battlegrounds (integrator's request; policy unchanged)

**How the ranker treats them today:**
- **PvP rank rewards.** These are loot.json `kind: "pvp"` sources, carrying `rank`, `title` and `faction`. The content-phases gate `honor-system` sits at launch, so they are open.
  - Rank 11 and up is never a default pick, only a labelled alternative (`excludeAbovePvpRankCap`).
  - Side comes from the source's faction.
  - A vendor row inherits the matching pvp source's rank (`vendorInheritsPvpRankGate`).
  - The pvp source is dropped for an item classic-db gives `required_honor_rank` 0 (`dropNonRankRewards`).
  - Label: "PvP rank N · Title · Faction".
- **Battleground items.** These are rep sources for factions 889/890 (WSG), 509/510 (AB) and 729/730 (AV), plus AV zone 2597, the BG quartermasters, Korrak's quests and Darkspear's Creeg Bothunk.
  - All of them sit behind the gates `battleground-reputation` and `darkspear-islands`, both at phase "later". No battleground item is eligible today, at any band, whatever loot.json emits.
  - An unknown new source kind (say "battleground" or "honor") would be ignored by `sourceFor`, leaving the item unsourced. The relabelled rows should keep kind `pvp`, `rep` or `vendor`.
- **Published PvP picks.** 69 slot picks, all rank 7 to 10:
  - Sergeant's Cape and Cloak;
  - Sergeant Major's Cape (16315, 16336, 16337);
  - Senior and Master Sergeant's Insignia;
  - the band-50 Stone Guard's and First Sergeant's plate, leather and mail pieces (220xxx ids);
  - Dragonhide Armsplints and Armguards for bear at 60.

  They sit mostly in tank and melee bands. The rank-10 bucket was lopsided: 21 Horde items against 1 Alliance item.
- **Stone Guard's after the pipeline's relabel (e1b402d61).** I re-ranked band 50 (Assassination, Arms, Survival) on the pipeline lane's scratch `full2` loot.json. Picks and set DPS are identical to the decimal. Only the source changed: "PvP rank 10 · Stone Guard · Horde" became "Lady Palanseer" (npc 12792, Horde, open under honor-system). So they won on stats; rank 10 was at the cap, never above it.

## Owner decision

- **Battleground rewards.** The owner says battlegrounds exist at launch. Flipping `battleground-reputation` and `darkspear-islands` in `data/curated/content-phases.json` to launch would make WSG, AB and AV reputation rewards eligible: friendly and honored while leveling, revered and exalted at 60. That would change BiS, mostly necks, rings, cloaks and some band 50 to 60 weapons. I have not done it; the owner rules.
- **PvP rank cap.** Forever's PvP gear is not rank-locked. If the owner wants it treated as plain vendor gear, the change is to stop reading `title_rank` in `pvpRankExceedsCap`. Today it is read, so nothing moves.

## Before/after

These are the ranker's published set scores, `beforeheld` against `after`: same held simdb, same incumbents, all 28 specs, every band. Set DPS is at the verify harness (300 iterations, seed 7); a kept set is at the guard harness. Healers are HPS and tanks are the tank score. Each row is the mean over the spec's 12 band rows, with the largest single moves.

| spec | mean | largest gain | largest loss |
|---|---|---|---|
| druid-balance | -0.09% | +1.25% (40 A, 130.8 → 132.4) | -1.92% (50 H, 157.2 → 154.2) |
| druid-feral | +0.00% | +0.03% (20 H) | - |
| druid-feral-bear (tank score) | +4.39% | +23.73% (30 H, 26.2 → 32.5) | -2.59% (60 A, 152.7 → 148.8) |
| druid-restoration (HPS) | +0.03% | +0.38% (30 A) | -0.32% (40 H) |
| hunter-beast-mastery | +0.01% | +0.34% (60 raid H, 688.2 → 690.5) | -0.18% (60 raid A) |
| hunter-marksmanship | -0.03% | +0.16% (20 H) | -0.28% (50 H, 243.4 → 242.8) |
| hunter-survival | +0.52% | +1.51% (50 H, 254.3 → 258.2) | -0.02% |
| mage-arcane | +0.22% | +1.51% (60 raid A, 600.7 → 609.7) | -0.98% (30 A, 77.5 → 76.8) |
| mage-fire | +0.03% | +0.84% (60 raid A, 664.7 → 670.3) | -1.13% (40 A, 150.8 → 149.1) |
| mage-frost | +0.60% | +2.99% (50 H, 194.1 → 199.9) | -0.63% (60 H, 299.5 → 297.6) |
| paladin-holy (HPS) | +0.46% | +1.46% (60 A, 229.6 → 233.0) | - |
| paladin-protection (tank score) | -0.53% | +4.02% (60 raid H, 425.9 → 443.0) | -6.08% (60 A, 168.3 → 158.0) |
| paladin-retribution | +0.09% | +0.61% (60 H, 261.9 → 263.5) | -0.34% (40 H) |
| priest-discipline (HPS; band 60 re-run after S 39296e4e3) | +0.00% | - | - |
| priest-holy (HPS) | +0.03% | +0.31% (40 A) | - |
| priest-shadow | +1.42% | +11.76% (40 A, 133.7 → 149.5) | -2.86% (60 H, 363.6 → 353.2) |
| rogue-assassination | +0.04% | +0.31% (60 raid H) | - |
| rogue-combat | +0.02% | +0.84% (60 H, 214.4 → 216.2) | -0.37% (60 A) |
| rogue-subtlety | +0.04% | +0.48% (60 H) | -0.07% |
| shaman-elemental | +0.92% | +8.10% (20 A, 22.7 → 24.6) | -1.35% (60 A, 152.0 → 149.9) |
| shaman-enhancement | +0.91% | +2.88% (60 H, 219.3 → 225.6) | -0.06% |
| shaman-restoration (HPS) | +0.06% | +0.37% (40 A) | - |
| warlock-affliction | -0.07% | +0.46% (30 H) | -1.15% (20 H, 57.1 → 56.5) |
| warlock-demonology | +0.20% | +1.78% (20 H, 51.3 → 52.2) | -0.22% (20 A) |
| warlock-destruction | +0.39% | +2.02% (40 H, 141.9 → 144.8) | - |
| warrior-arms | -0.30% | +0.03% | -1.89% (60 raid A, 740.0 → 726.0) |
| warrior-fury | +0.21% | +0.86% (20 A) | - |
| warrior-protection (tank score) | -0.71% | +1.01% (60 raid A, 375.0 → 378.8) | -3.23% (60 H, 187.6 → 181.5) |

Totals:
- Across the 300 DPS and healer band rows, 39 rose more than 0.5% and 13 fell more than 0.5%.
- The guard kept the previously published set in 56 band rows on origin/main and 45 with this branch.
- The largest gains come from K2: Shadow band 40 +11.8%, Elemental band 20 +8.1% and band 30 +4.1%, Frost band 50 +3.0%, Survival bands 30 and 50 +1.4 to +1.5%, Enhancement band 60 +2.9%.
- The losses are not removed. See Still open.

K1 (paladin axes) needs `items/paladin.json` to carry the axes, which only the next data run writes. I measured it on a scratch copy of paladin.json with the 257 warrior-and-shaman axes added. That run used no incumbent guard and this branch at 69a6adf3e:
- **Retribution:** band 60 bare +3.10% (Alliance), +2.41% (Horde), Forest Defender's Axe; band 20 Horde +1.54%, Forsaken Greataxe; every other band unchanged.
- **Protection paladin:** band 50 −3.2% and −1.9%. That is not the axe: its two runs' tank sweeps disagree beyond their own noise (parry significant in one and not in the other), and the armor picks follow the weights. See Still open.

## For other lanes

- **Pipeline / data:** `data/pipeline/proficiency.py` now gives paladins both axes (S 188718042, with its test). `items/paladin.json` gains about 257 axes at the next data run; that output is nightly-owned, so it is not regenerated here.
- **Pipeline:** keep `kind` on the relabelled PvP rows to `pvp`, `rep` or `vendor`. The ranker ignores any other kind. Field name `title_rank` is what S 19f27af30 reads.
- **Page (web/src/lib/items/source-labels.ts):** its `pvpSourceLabel(rank, faction)` will also print "PvP rank 0" for an unranked reward.
- **Integrator:** every ranker commit here moves BiS picks and published stat weights for specs across every group, once the owner's refresh runs. The weights rail shows the geared sweep. No sim DPS moves: these are ranker-only changes.

## Still open

- **Tank weights churn between runs.** Protection paladin band 50 measured parry significant in one run and not in the next, at 100 iterations a direction, and the armor picks followed. A higher iteration count for tank sweeps would settle it. It costs run time, so I have not changed it.
- **The dagger rule is broader than it needs to be.** Combat's APL casts Backstab with a Sinister Strike fallback, so the ranker forces daggers on Combat. Daggers win today by 4%; an APL value for "main hand is a dagger" (rogue curation note) would let the sim decide.
- **Non-unique items twice.** The ranker never recommends two copies of the same ring or weapon. That is a design choice; the owner can rule on it.
- **The committed simdb** still types held items as main hands. The next data run fixes it, and every caster and Feral bear re-rank depends on it.
- **K2 leaves some band rows lower than origin/main's re-rank would.** These are path effects that the two guards above do not catch:
  - Shadow band 60: −1.5% Alliance, −2.9% Horde.
  - Arms raid band 60: −1.9% Alliance, −1.1% Horde.
  - Balance band 50 Horde: −1.9%.
  - Elemental band 60: −1.1 to −1.4%.
  - The tank scores: Protection paladin band 60 Alliance −6.1%, Protection warrior band 60 −2.2 to −3.2%.

  None of these can publish below the current published set, because the incumbent guard still applies. A fuller fix is to run the whole per-faction pipeline under both sweeps and let the guard pick. That doubles the ranker's run time, so I left it for the integrator to weigh. Tanks were already swept on baseline armor (`tankLadderCharacter`); excluding them from the geared sweep is an option.
- **The off-hand weapon-DPS term** counts full DPS where the game deals half. It did not change Fury's pick at 60; other dual-wield bands were not checked one by one.
