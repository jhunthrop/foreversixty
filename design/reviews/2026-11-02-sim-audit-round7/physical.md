# Sim audit, round 7: physical group (2026-11-02)

This lane covers Arms, Fury, Protection, Assassination, Combat, Subtlety, Beast Mastery, Marksmanship and Survival (melee), plus hunter pets, on client 1.60.1.70707.

- **Baseline:** engine pin b57765fb9 with the committed simdb.bin (c0bd0bd3d).
- **Branches:** fork fixes are on `lane/r7-physical` (E) and site fixes on `wt/r7-physical` (S).
- **Measurements:** `sim/cmd/spec-breakdown`, 10000 iterations, seeds 1 and 50001, band 60, raid preset, published sets.
- **Fresh re-audit:** the warrior and rogue passes were run by two independent read-only reviewers. Every row they marked confirmed was re-checked against the client rows here, and got a failing test before its fix.
- **bis run 41363895819:** cancelled by the owner. In its place the ranker ran locally in a scratch directory (not committed); see X3.

## Cross-cutting

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| X1 | **Ledger: the "24 aura-99-only spells" are all Attack Power - Feral.** All 24 item spells that have aura 99 but no aura 124 are "+N Attack Power in Cat, Bear, and Dire Bear forms only." simdb read them as generic attack power, mirrored into ranged attack power, for every class. In /sim a warrior or hunter wielding High Warlord's Pig Sticker gained 358 attack power: Arms 740.4 → 894.7 (+21%), Marksmanship 666.1 → 714.6. They now map to the engine's FeralAttackPower, which only Cat and Bear form turn into attack power. A feral druid with Mace of Unending Life measures the same before and after (579.0 / 578.7). Only 6 simdb rows change: 21407, 22632 and the four rank-18 PvP polearms. None of them is a published pick. The four polearms do appear as alternatives in the warrior, hunter and Retribution sets. | SpellShapeshift 1.60.1.70707 gives 26153, 460339 and the others ShapeshiftMask 145 (Cat 1, Bear 16, Dire Bear 128). Wowhead Forever tooltips for 234547 and 21407 (read 2026-11-01) say "+358 / +140 Attack Power in Cat, Bear, and Dire Bear forms only." Tests `test_feral_form_spells_are_named_by_the_client_shapeshift_mask` and `test_a_feral_form_attack_power_spell_is_feral_attack_power` failed first. The fetch now takes SpellShapeshift. | CONFIRMED-BUG (client + server tooltip) | S c9f0c1bb1 |
| X1a | X1, what remains: the planner's and ranker's `items/<class>.json` (normalize path) still list these as `attack_power`. The ranker's score therefore credits a non-druid with the 358, and the planner shows it as generic attack power. No pick moves today. The fix needs a planner stat key or a class-aware rule, which touches druid ranking (`feral_attack_power` is only in Feral cat's weights, not bear's). | `pipeline/normalize/effects.py`; `data/tests/test_specs.py:341` | OPEN (pipeline/ranker) | - |
| X2 | **Ledger: the warrior Deadmines addon spec hardcoded its pick.** Five dungeon-line specs pinned Smite's Mighty Hammer, its boss, "4 more" and six printed lines. They now read the picks, bosses, chances and count from the shipped Data, as 2c5b85218 did for the hunter. | `addon/tests/reward_marks_spec.lua`; busted 43/43 | fixed (test) | S 8b293152f |
| X3 | **Item 2, checked locally** (the bis run was cancelled). `leveling-bis` on this branch's committed sources: Survival takes Deadly Aspects 5/5 at bands 50 and 60, with neither Improved Concussive Shot nor Hawk Eye. Raid 60 is `5/0050052010/…022151`, set DPS 853.7 horde. The Beast Mastery raid build publishes (Mortal Shots 4, Efficiency 3). The horde hunter sets wear Rune of the Guard Captain, which simdb now states once: 14 attack power and 14 ranged attack power. Horde Marksmanship measures 666.1 and Beast Mastery 673.4, the round-6 values after the dedup. Horde Arms and Fury carry 14 AP (740.4 / 775.0). | scratch output in `scratchpad/r7/physical/bisout/` | aligned | - |
| X4 | **Mark of the Chosen 17774 used vanilla numbers.** The engine gave +25 to every stat and had no proc cooldown. The client gives 21 stats and a 120 s cooldown on the proc. It procs only when struck, so DPS sets are unaffected. Tanks wearing it move: bear 60 raid horde -0.49 / -0.46% threat, Protection paladin 60 bare -1.0 / -0.9% (alliance) and -0.84 / -0.84% (horde). | 21970 is aura 29 at 21; SpellAuraOptions 132914 (ProcChance 2, ProcCategoryRecovery 120000); test `TestMarkOfTheChosenMatchesTheClient` | CONFIRMED-BUG | E 5680c5426 |
| X5 | Rage from damage taken uses the attacker's level (63) in `core/rage.go:175`. cmangos `Unit::RewardRage` and Kalgan's formula use the warrior's own level, which would mean about 5% more rage from hits. Upstream wowsims chose the attacker's level, and the only other source is an emulator. | `core/rage.go:175` | OPEN (a tank rage log) | - |

## Hunter: Survival

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| S1 | **Owner ruling: Hawk Eye's points go to Deadly Aspects.** The leveling build is now `5/005005001/500230231051222151`, with Deadly Aspects 5/5 in place of Hawk Eye 3/3. Hawk Eye leaves the Survival raid keeps. The addon's default builds follow the guide. The bands derive from the guide: band 50 becomes `5…-001…`, and bands 20 to 40 do not change. | guide frontmatter; `raid-keeps.json`; `addon-builds.json`; leveling and talent-search tests pass | done | S 123c7b26f |
| S2 | **Question for the integrator.** The raid build already had Deadly Aspects 5/5, so its three Hawk Eye points needed somewhere else to go. Ten placements were measured (horde raid, seed 1 / 50001). The best is Lightning Reflexes 3→5 plus Rapid Killing 1→2, at 854.6 / 855.4 against 841.5 / 842.6 (+1.6%), and it is the one committed. Next best: Rapid Killing plus Lightning Reflexes plus Resourcefulness, 852.1. Resourcefulness 2 with Rapid Killing, 848.8. Survivalist and Improved Stings add nothing. Rapid Killing's second point is worth about +0.8% because the encounter's duration varies: a fight past 180 s gets a second Rapid Fire. **Note:** the committed raid build is now only +0.7% over the new leveling build in raid (854.6 against 848.5), under the adoption rule's 1% bar. Please decide whether Survival keeps a separate raid build. | `scratchpad/r7/physical/` runs | ASK | S 123c7b26f |
| S3 | Sweep for other range-only talents in the Survival builds: none remain. Every point left is melee damage, crit, hit, agility, a melee-cast cooldown, the Counterattack path (Deterrence), traps the rotation uses, or Survivalist (health, a path point). | tooltips in `talents/hunter.json` | aligned | - |
| S4 | Fresh re-audit: no new findings. Checked: Raptor Strike (+70 at rank 8, 6 s), Mongoose Bite (weapon +57), Counterattack (50% weapon +110), Strider Kick (100% weapon, 8 s, 5.81% base mana), Lacerate (58 a tick, 7 ticks, snapshot bleed), Wing Clip (50, 80 mana), Immolation Trap (5 × 138, 30 s shared trap cooldown, no AP), Rapid Fire (40% melee and ranged haste), Aspect of the Beast's proc mask 4 (melee auto only, matching the text), and Heart of the Lion. | `spellconst/hunter.json`; SpellAuraOptions 246416 | aligned | - |

## Hunter: Beast Mastery and Marksmanship

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| B1 | **Round 6's B4, Serpent Sting's raid tick, is explained.** Bare, a tick averages 93.2 before crit. That is 83 × 1.13 (Improved Stings 2/3) at Ranged-defense crits (×2). Raid adds two things. Forever's Curse of the Elements covers every magic school, Nature included (1311680, misc 126), for ×1.10. The raid set's Ironfeather Leggings add 38 spell power, × 0.2 a tick. Together they predict 102.5 + 8.6, against a measured 111.5. | per-buff breakdowns at 3000 iterations; spellconst 1311680, 13555 | aligned (B4's tick part closed) | - |
| B2 | B4's other half is still open. Aspect of the Hawk's proc row is mask 320 (ranged auto and ability), while the text and the engine use Auto Shot only. Deadly Aspects adds its 10% to that row (aura 107, misc 18). | SpellAuraOptions 126992; SpellEffect 697480 | OPEN (a log) | - |
| B3 | Serpent Sting's 0.2 spell power a tick has a single source (the client row states none). After B1 it is also what makes Ironfeather Leggings the raid pick. | | OPEN (carried) | - |
| B4 | Marksmanship's +0.9% search winner stays held under the adoption rule. No engine change touches hunters. | | HOLD | - |

## Hunter pets

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| PET1 | Pet family, haste and expertise are carried open. Nothing new this round. | | OPEN (carried) | - |

## Warrior: Arms and Fury

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| W1 | **Whirlwind's off-hand hit missed Dual Wield Specialization (fresh).** 1680 strikes "with both melee weapons". The off-hand hit is dealt through the main-hand spell, so the talent's ProcMask filter (`talents.go`) never reached it: at 5/5, Fury lost +25% on that hit. The off-hand weapon damage now takes the talent. | 23584 is aura 122 (off-hand damage %, +5 a rank), which scales every off-hand weapon strike. Test `TestWhirlwindOffHandTakesDualWieldSpecialization` failed first. | CONFIRMED-BUG | E a3b0a2611 |
| W2 | **Rend's crits did not take Impale (fresh).** Impale's family mask includes Rend, and Rend crits under the owner's bleed ruling, but Rend set no crit damage bonus. | Impale 16493: SpellEffect 693622 (aura 108, misc 15, mask 0xee604cee, bit 32 is Rend). Test `TestRendCritsTakeImpale` failed first. | CONFIRMED-BUG | E a3b0a2611 |
| W3 | **Bloodthrill's Overpower window was 6 s; the client says 5 s (fresh).** The Arms guide repeated the 6 s and is corrected. | 1282733 has DurationIndex 28, which SpellDuration gives as 5000 ms. Test `TestBloodthrillWindowIsFiveSeconds` failed first. | CONFIRMED-BUG | E a3b0a2611; S 4acb9f45e |
| W4 | Heroic Strike carries the main-hand auto-attack proc flag, so it procs Unbridled Wrath and uses up Flurry charges. The client's rows proc on white swings only (mask 4). The two effects pull DPS in opposite directions, each about 0.3 to 0.5% for Fury. | `heroic_strike_cleave.go:84`; AO for 12322 and 12966 | OPEN (a log: does a queued Heroic Strike count as a swing) | - |
| W5 | Battle Shout is 115 attack power at level 60 (111 + 0.6 × 8, floored). The Forever tooltip for 11551 says 116. Round 6's "aligned at 116" was wrong as written. Server rounding is unsettled; about 0.03%. | `core/buff_ranks.go:46` | OPEN | - |
| W6 | Also unsettled: rage on dodged and parried swings; Overpower's own hit can never proc Bloodthrill (tiny); Arcanite Champion's proc rate (carried). | | OPEN | - |
| W7 | Fresh re-audit, aligned. **Stances and masks:** stance passives (Berserker 7381 +3% crit and -20% threat, Battle -20% threat), the 1 s stance cooldown, Tactical Mastery, and every rotation spell's SpellShapeshift mask. **Abilities:** Bloodrage (10 rage, then 1 a second for 10 s), Thunder Clap (magic defense type, 103 damage, 6 s), Hamstring, Whirlwind, Improved Slam, Sweeping Strikes, Death Wish, Recklessness. **Shared cooldowns:** category 971 (Bloodthirst, Mortal Strike, Shield Slam) and category 65 (Overpower, Revenge). **Talents:** Deep Wounds, Anger Management, Flurry, Unbridled Wrath, Enrage. **Items:** set bonuses and items on the band-60 sets. Execute's 15 a point of rage is confirmed by EffectChainAmplitude 1.5. | SpellCategories 70707 (fetched to `scratchpad/r7w/dl/`) | aligned | - |

## Warrior: Protection

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | Checked and aligned: Shield Block (+75%, 2 charges, 7 s / 5 s, no global cooldown), Shield Wall, Last Stand, Shield Specialization and Master of Defense rage, Defiance (Defensive Stance only), Revenge (109 to 133 + 0.25 attack power), and Shield Slam plus block value. No engine change moves Protection threat (1067.7 / 1066.5 alliance, 1055.9 / 1055.0 horde). | | aligned | - |
| P2 | Revenge, Shield Slam and Thunder Clap flat threat, and Sunder's attack power threat: no new evidence. | | OPEN (launch threat log) | - |
| P3 | Rage from damage taken (X5). | | OPEN | - |

## Rogue: Assassination, Combat, Subtlety

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R1 | **Hemorrhage used the weapon's own speed (fresh).** The client normalizes it, as it does Backstab, Sinister Strike, Ambush and Mutilate (all Effect 121). Ghostly Strike has no Effect 121, so it stays unnormalized. On a 1.9 s Heartseeker, Subtlety was about 2.1% high. | 16511: SpellEffect 693307 is Effect 121 (normalized weapon damage), with 1231002 Effect 31 at 100%. Test `TestHemorrhageUsesNormalizedWeaponDamage` failed first (332.9 against the normalized 304.3). | CONFIRMED-BUG | E a67c63d0a |
| R2 | **Mutilate's off-hand hit missed Dual Wield Specialization (fresh).** The talent was applied to the white off-hand swing only. At 2/5, Assassination was about 1.0% low. The talent now scales the weapon part only, not the flat 67, as aura 122 raises the off-hand weapon's damage. | 13715 is aura 122 (off-hand damage %). Test `TestMutilateOffHandTakesDualWieldSpecialization` failed first (1.00 against 1.10). | CONFIRMED-BUG | E a67c63d0a |
| R3 | Preparation never fires, because no rotation casts Vanish. Its reset list also omits cooldowns the client text resets (e.g. Ghostly Strike). It moves no DPS: Subtlety's Cold Blood waits on its rare Eviscerates. | `preparation.go`; 14185 | OPEN (0 DPS) | - |
| R4 | Eviscerate adds 3% of attack power per combo point (`eviscerate.go:80`). The client row 689090 has BonusCoefficientFromAP 0, and the tooltip formula has no attack power term. This is the Eviscerate twin of Rupture (round 6 R4). Cold Blood's charge spending on Mutilate's two hits and Hack and Slash's crit on poisons are near zero. | | OPEN (level-60 log) | - |
| R5 | Venom and Rupture attack power: no new evidence. | | OPEN (carried) | - |
| R6 | Fresh re-audit, aligned. **Finishers:** Slice and Dice (20/30% haste, 9 to 21 s, Improved 15/30/45%), Eviscerate's damage and cost (with Flawless Execution), Rupture's per-point ticks against every tooltip figure. **Builders and openers:** Backstab, Ambush, Ghostly Strike, Sinister Strike costs, Mutilate. **Cooldowns:** Cold Blood, Premeditation, Adrenaline Rush, Blade Flurry. **Talents:** Lethality's mask (excludes Ambush and Eviscerate), Malice, Precision, Seal Fate's 500 ms, Relentless Strikes, Ruthlessness, Thousand Cuts, Quietus. **Poisons and energy:** Improved Poisons and Vile Poisons masks, energy regen. | 70707 rows, TraitDefinitionEffectPoints, Forever tooltips | aligned | - |

## Before/after (band 60, raid preset, published sets and talents, 10000 iterations)

**Before:** b57765fb9 with the committed simdb.

**After:** `lane/r7-physical` at 5680c5426 plus `wt/r7-physical`. The feral-AP simdb change (X1) touches no published item, so the after simdb is identical for these sets.

| Spec | alliance % (seed 1 / 50001) | horde % (seed 1 / 50001) | seed 1 before → after (alliance / horde) |
|---|---|---|---|
| Arms | -0.14 / -0.20 | -0.14 / -0.19 | 728.3 → 727.3 / 740.4 → 739.4 |
| Fury | +0.48 / +0.48 | +0.48 / +0.49 | 771.3 → 775.0 / 775.0 → 778.7 |
| Protection (threat) | 0.00 / 0.00 | 0.00 / 0.00 | 1067.7 / 1055.9 |
| Assassination | +1.00 / +1.00 | +0.99 / +0.99 | 621.9 → 628.1 / 619.0 → 625.1 |
| Combat | 0.00 / 0.00 | 0.00 / 0.00 | 622.7 / 622.8 |
| Subtlety | -2.08 / -2.08 | -2.08 / -2.07 | 558.4 → 546.8 / 558.3 → 546.7 |
| Beast Mastery | 0.00 / 0.00 | 0.00 / 0.00 | 669.3 / 673.4 |
| Marksmanship | 0.00 / 0.00 | 0.00 / 0.00 | 661.7 / 666.1 |
| Survival | 0.00 / 0.00 | 0.00 / 0.00 | 775.0 / 778.4 |

**What drives each move:**
- Arms: Bloodthrill's shorter window, partly offset by Impale on Rend.
- Fury: Whirlwind's off hand.
- Assassination: Mutilate's off hand.
- Subtlety: Hemorrhage normalized.

**Survival's new builds** (they publish at the next re-rank; seed 1 / 50001):
- Raid build, Hawk Eye → Lightning Reflexes and Rapid Killing: horde +1.56 / +1.53% (841.5 → 854.6), alliance +1.38 / +1.38% (837.1 → 848.7).
- Leveling build, Hawk Eye → Deadly Aspects:
  - Band 60 bare: horde +4.22 / +4.30% (329.8 → 343.7), alliance +4.37 / +4.32%.
  - Band 60 raid: horde +4.48 / +4.63%.
  - Band 50 bare: horde +3.20 / +3.35%, alliance +3.23 / +3.24%.

**Outside the group:** Mark of the Chosen (X4) moves bear 60 raid horde threat -0.49 / -0.46%, and Protection paladin 60 bare threat -1.0 / -0.9% (alliance) and -0.84 / -0.84% (horde). DPS sets that wear the Mark do not move, because the target never swings.

## Still open

- **Integrator:**
  - Whether Survival keeps a separate raid build at +0.7% (S2).
  - TestRotationLadder rows move for Survival (bands 50 and 60), Fury, Arms, Assassination and Subtlety.
  - Assassination's talent search should be re-run now that Dual Wield Specialization reaches Mutilate.
  - `data/tests/test_apl.py::test_every_spell_the_rotations_name_exists_with_that_rank` fails on origin/main: druid-feral-bear names 1235827 at rank 3, the build says 5. This predates the lane.
- **Pipeline or ranker:** feral attack power in `items/<class>.json` (X1a).
- **Need a log:**
  - Rage from damage taken, by level (X5).
  - Heroic Strike as a swing (W4).
  - Battle Shout's rounding (W5).
  - Eviscerate and Rupture attack power (R4, R5).
  - Hawk's proc mask (B2).
  - Serpent Sting's spell power (B3).
  - Revenge, Shield Slam and Thunder Clap threat (P2).
  - Arcanite Champion's proc rate.
- **Zero DPS:** Preparation's trigger and reset list (R3). Whirlwind's off-hand hit also rolls on the main-hand table and main-hand procs (W1, not changed).

Sources (read 2026-11-01): https://wago.tools/db2/SpellShapeshift/csv?build=1.60.1.70707 and /SpellCategories; the layered 70707 raw tables (SpellEffect, SpellAuraOptions, SpellMisc, SpellDuration, Spell); https://nether.wowhead.com/forever/tooltip/item/234547 and /21407.
