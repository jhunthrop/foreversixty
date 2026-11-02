# Sim audit, round 8: physical group (2026-11-02)

This lane covers Arms, Fury, Protection, Assassination, Combat, Subtlety, Beast Mastery, Marksmanship and Survival (melee), plus hunter pets, on client 1.60.1.70707.

- **Baseline:** fork pin cd4dfdcc0 with site origin/main 170b7b84c, each with its own `make simdb`.
- **Branches:** fork fixes are on `lane/r8-physical` (E), site changes on `wt/r8-physical` (S).
- **Measurements:** `sim/cmd/spec-breakdown`, 10,000 iterations, seeds 1 and 50001, both factions, on the published sets.
- **Client rows:** the layered 70707 raw tables (SpellEffect, SpellLevels, SkillLineAbility, ItemXItemEffect, ItemEffect, SpellMisc, SpellAuraOptions, SpellClassOptions), read 2026-11-01. The Era comparison uses data/builds/1.15.9.69722/raw.

## Owner ruling 1: Survival has one build at 60

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| O1 | **The separate Survival raid build is gone.** It traded Resourcefulness 2 for Rapid Killing 2 and measured +0.63 / +0.64% (alliance) and +0.75 / +0.76% (horde) over the leveling build on the band-60 raid preset, under the adoption rule's 1% bar. The guide drops its `raidBuild:` line. Everything that reads the raid build goes through `leveling.GuideRaidBuildTalents`, which falls back to `build:` when there is no `raidBuild:`. That covers the ranker's raid pass, talent-search, dps-per-point and addon-builds. So raids now use the leveling build. `addon-builds.json` is regenerated from the guide, and its `raid` entry for Survival now matches `leveling`. The guide's talent paragraph says the trade is too small to carry a second build. A vitest, "Survival has one build at 60", fails if a `raidBuild:` comes back. The raid keeps (`raid-keeps.json`) still guard the search. The talent-search report in design/reviews/talent-search is history and stays. | raid build `5…/0050052010…/…1022151` against leveling `5…/0050050010…/…1222151` (LadderTalentString of each guide line); `scratchpad/r8/physical/sv-ruling.txt` | done (ruling) | S 768c32817 |

## Cross-cutting

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| X1 | **Twenty-six chance-on-hit weapons used Era or Season of Discovery proc damage.** The engine typed each proc by hand. The client names each weapon's proc spell in ItemXItemEffect (trigger 2) and states its roll in SpellEffect (Effect 2: EffectBasePointsF, Variance). Forever retuned many of these and gave several of them new spells. On published sets: Nightblade 231 ±30% Shadow (was Era's 125 to 275; Arms 40), Coldrage Dagger's own 1293790 at 42 ±20% (was 13439's 20 to 30; rogues 40), Hammer of the Northern Wind 50 ±20% (was 25; Fury 50), Gut Ripper 105 ±12% (was 108; many band-40 sets). Off the published sets: Vilerend Slicer 1320808 at 225 (was 75), Deathbringer 1305398 at 140 ±25%, Vis'kag 1305394 at 236 ±17%, Linken's Sword 98 ±25% (was 60), Perdition's Blade 48 ±17%. Perdition's Blade was also mistyped as 40 to 96; Era is 40 to 56. The other 16 are listed in the test. The rolls now live in one table read through `clientdamage.Effect`. Proc rates stay the engine's: the client states none. | `sim/common/weapon_proc_rolls.go`; Era rows for comparison (18211: 124 + 1d151); test `TestWeaponProcsMatchTheClient` failed first on 34 rows | CONFIRMED-BUG (client rows; Forever tooltips agree, e.g. Nightblade "causing 231 Shadow damage") | E 3fcf9342d |
| X2 | **Two band-30 rogue weapons had no proc at all.** Claw of the Shadowmancer (16409, 35 Shadow) and Thornspike (1293183, 45 Physical) are worn by the published band-30 rogue sets. They now proc at the file's one-a-minute convention. Thornspike is missing from the fork's own with_db item table, so it registers only in the site's build. | ItemXItemEffect trigger 2; SpellEffect; same test | CONFIRMED-BUG | E 3fcf9342d |
| X3 | **Curse of Recklessness was Era's: -640 armor and +90 attack power.** Forever's rank 4 (11717) is aura 674 at -505 armor beside aura 92 (no fleeing). The attack power is gone. The Forever tooltip reads "reducing the target's armor by 505". The 70291 build diff records the change (EffectAura 22 to 674 on all four ranks). The raid preset carries the curse, so every physical attacker hit 135 armor too little: about -2% for every physical spec on the raid preset. | SpellEffect 11717 index 1; Era 1.15.9 row (-641, aura 99 +89); https://nether.wowhead.com/forever/tooltip/spell/11717 (read 2026-11-01); test `TestCurseOfRecklessnessIsForevers` failed first | CONFIRMED-BUG (client + server tooltip) | E a16b3446e |
| X4 | Sunder Armor (11597, 5 × 450), Faerie Fire (9907, 505) and Curse of Recklessness all moved to aura 674 in Forever (Era: aura 22). The engine stacks the three, as Era does. Whether aura 674 still stacks across the three has no second source. | SpellEffect 11597, 9907, 11717 | OPEN (a log with all three debuffs up) | - |
| X5 | **Rank-0 abilities above their learn level.** The engine never enforces `Spell.RequiredLevel`, and the site's rank rewrite skips abilities spellranks.json carries as a single rank-0 row. So every APL cast of such an ability runs at any level. In this group: Whirlwind (X6) and Berserker Stance (X7). Shield Wall (871, level 28) is in the Protection APL but is never cast below 28. Across all specs, the other APL casts of rank-0 abilities above level 20 are Barkskin, Frenzied Regeneration, Rapid Fire, Hammer of the Righteous, Bane of Doom and Recklessness, all already gated. One is not gated: **Innervate (29166, level 40) in druid-restoration** (For other lanes). | spellranks.json; SpellLevels BaseLevel; `internal/spellranks/spellranks.go` singleTierLevelOverrides | see X6, X7 | - |
| X6 | Side effects of round 7, re-checked. Rend crit with Impale: Impale's mask 0xEE604CEE covers Rend, Whirlwind, Bloodthirst, Mortal Strike, Overpower, Execute, Heroic Strike, Hamstring and Slam. Every one of them sets `impale()`, and Deep Wounds (12721) has no class options, so it is outside the mask. Bloodthrill is 5 s. Dual Wield Specialization reaches only the weapon part of Mutilate's off hand. Hemorrhage is normalized, and its debuff is the Rupture amplifier (aura 271, 15). Strength of Earth is 42 at 60, for both factions: the raid preset carries the totem for alliance too. The Deadmines addon spec reads the shipped Data. | client rows above; `sim/warrior/*.go`; `hemorrhage.go` | aligned | - |
| X7 | Item effects on every published band of the nine specs, checked against HasItemEffect and the client: Teebu's Blazing Longsword (1300753, 140 Fire), Arcanite Champion (16916: 120 Strength, heal 360 ±50%), Thrash Blade, Flurry Axe, Manual Crowd Pummeler, Bow of Searing Arrows and the Hearts of the Mountain are modelled as the client says. Every set worn (Devilsaur, Stormshroud, Gladiator, Defias, Viper, Scarlet, Magister, Shadowcraft, Beaststalker, Lightforge, Valor, the six PvP sets) is registered. Not modelled: Gloves of Holy Might (1325229, 112 Holy on hit; Survival 40 and 50 bare) needs a non-weapon melee proc path and a rate; Toxic Revenger's cloud (11790; rogues 30) needs a rate; Gut Ripper's and the Fatal Wound weapons' "double damage below 25% health" effect is unmodelled. Each is well under 1% of the set's DPS. | `scratchpad/r8/physical/bisitems.txt`, `sets.txt` | OPEN (no proc rate in the client) | - |
| X8 | Rage from damage taken by level (round 7 X5), Heroic Strike as a swing (W4), Battle Shout's rounding (W5) and Arcanite Champion's proc rate: no new evidence. | | OPEN (carried) | - |

## Warrior: Arms and Fury

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| W1 | **Fury cast Whirlwind at bands 20 and 30.** Whirlwind (1680) is SpellLevels BaseLevel 36, but it was a fifth of a band-20 Fury warrior's damage. The engine now refuses the cast below 36. spellranks' single-tier overrides (and the ladder's mirror) list it at 36, so the rotation drops the line below that. | SpellLevels 97881 (1680: 36); SkillLineAbility 4888; tests `TestWhirlwindIsLearnedAtThirtySix` and `TestSingleTierOverridesGateTheirIDs` failed first | CONFIRMED-BUG | E f9b58fcab; S 2624f7fcb |
| W2 | **A warrior below 30 used Berserker Stance.** Berserker Stance (2458) is BaseLevel 30, the level-30 quest. A band-20 Fury warrior opened in it, gaining 3% crit, and band-20 Arms swapped into it. A stance is now castable, and a fight's default stance chosen, only once it is learned. A Fury warrior below 30 opens in Battle Stance. | SpellLevels 98539 (2458: 30); test `TestBerserkerStanceIsLearnedAtThirty` failed first | CONFIRMED-BUG | E f9b58fcab; S 2624f7fcb |
| W3 | With W1 and W2, a band-20 Fury warrior is in Battle Stance with nothing to press but Heroic Strike, while its Overpower window sits unused. The Fury rotation gains a bare Overpower line. It fires only in Battle Stance, so from 30 on it falls through. At band 20 it is worth about +9% (3000 iterations), and it does not move bands 30 to 60. The fork's copy is synced. | `data/curated/apl/warrior-fury.json`; test_apl's name sidecar | rotation follow-up to W1/W2 | E f9b58fcab; S 2624f7fcb |
| W4 | Fury at 30 to 35 now has Berserker Stance and no Whirlwind or Bloodthirst: Heroic Strike is its only special. Staying in Battle Stance for Overpower until 36 might be better. That needs a rotation search with a stance condition. | | OPEN (rotation search) | - |
| W5 | Ranks per band, read from what each band's sim actually cast: Heroic Strike, Execute, Overpower, Rend, Mortal Strike, Slam, Hamstring, Battle Shout, Bloodrage and Cleave are each the highest rank learnable at 20, 30, 40, 50 and 60. Cleave registers its rank 5 at every level, but no single-target rotation casts it (0 DPS). The only higher ranks at 60 are the Ahn'Qiraj books (Heroic Strike 9, Battle Shout 7, Revenge 6), which stay out while IncludeAQ is off. | `scratchpad/r8/physical/ranks.py`, `ranks.out` | aligned | - |

## Warrior: Protection

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | Curse of Recklessness (X3) lowers Protection's raid threat by about 1.8%. Shield Wall is never cast below its level 28. Ranks are as in W5. Round 7's Shield Block, Revenge and Shield Slam rows stand. | measure60 | covered by X3 | E a16b3446e |
| P2 | Revenge, Shield Slam and Thunder Clap flat threat, and rage from damage taken. | | OPEN (carried) | - |

## Rogue: Assassination, Combat, Subtlety

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R1 | Band-30 and band-40 rogue sets move with X1 and X2: Claw of the Shadowmancer and Thornspike now proc; Coldrage Dagger's proc nearly doubles. | measure.out | covered by X1, X2 | E 3fcf9342d |
| R2 | Ranks per band: Sinister Strike, Backstab, Eviscerate, Rupture, Ambush, Slice and Dice, and Instant and Deadly Poison are the highest learnable at every band. The only higher ranks at 60 are the books (Eviscerate 9, Backstab 9). Stealth registers its rank 4 at every level (0 DPS). Instant Poison is 20% a hit and Deadly and Wound Poison 30% (SpellAuraOptions 8679, 11340, 2823, 11356, 13219). | ranks.out; SpellAuraOptions | aligned | - |
| R3 | Held: Assassination leveling +3.6%, which drops Murder and Venom. No engine change this round touches either talent, so the hold stands. | | HOLD | - |
| R4 | Eviscerate's and Rupture's attack power; Preparation's trigger. | | OPEN (carried) | - |

## Hunter: Beast Mastery, Marksmanship, Survival, pets

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| H1 | Ranks per band: Arcane Shot, Serpent Sting, Aimed Shot, Raptor Strike, Mongoose Bite, Wing Clip, Immolation Trap, Aspect of the Hawk and Aspect of the Beast are the highest learnable at every band. The only higher ranks are the books (Serpent Sting 9, Hawk 7). Multi-Shot is one rank in Forever: 14288 to 14290 have SkillLineAbility rows but no spell data, and 2643 is normalized weapon damage with no flat bonus. Explosive Trap's duration mismatch in the conformance report is the trap's arming life (0 DPS: no rotation lays it). | ranks.out; SpellEffect 2643; conformance hunter.golden.md | aligned | - |
| H2 | The Raptor Strike note in the Survival rotation said "gated at 20% mana", but the condition is mana >= 0%, which is always true. The guide already says the gate was a small loss and came out, so the note was what was stale. Only the note changed (0 DPS). | `data/curated/apl/hunter-survival.json` | stale note | S 3f8d91d77; E b2e96663e |
| H3 | Held: Marksmanship +0.9%. No change this round touches its build. | | HOLD | - |
| H4 | Aspect of the Hawk's proc mask 320, Serpent Sting's spell power, pet family, haste and expertise. | | OPEN (carried) | - |

## Before/after

**Before:** fork cd4dfdcc0 with site 170b7b84c. **After:** `lane/r8-physical` (a16b3446e; b2e96663e after it only edits an APL note) with `wt/r8-physical`. The figures are relative changes, seed 1 / seed 50001, then the absolute values at seed 1.

Band 60, raid preset (CoR is X3; Protection is in threat per second):

| Spec | alliance % | horde % | seed 1 alliance / horde |
|---|---|---|---|
| Arms | -2.15 / -2.15 | -2.15 / -2.15 | 718.0 → 702.6 / 729.9 → 714.2 |
| Fury | -2.21 / -2.21 | -2.23 / -2.23 | 764.0 → 747.2 / 767.8 → 750.7 |
| Protection (threat) | -1.82 / -1.76 | -1.76 / -1.72 | 1055.5 → 1036.3 / 1043.7 → 1025.3 |
| Assassination | -2.02 / -2.02 | -2.02 / -2.02 | 623.5 → 610.9 / 620.6 → 608.0 |
| Combat | -2.10 / -2.10 | -2.10 / -2.10 | 618.2 → 605.2 / 618.3 → 605.3 |
| Subtlety | -1.84 / -1.84 | -1.84 / -1.84 | 542.2 → 532.2 / 542.0 → 532.0 |
| Beast Mastery | -1.93 / -1.93 | -1.94 / -1.94 | 670.9 → 657.9 / 674.8 → 661.8 |
| Marksmanship | -2.07 / -2.07 | -2.07 / -2.07 | 662.7 → 649.0 / 666.7 → 652.9 |
| Survival (published talents) | -1.97 / -1.97 | -1.97 / -1.97 | 769.3 → 754.2 / 772.6 → 757.4 |

Survival's raid number at 60 under ruling 1: the old raid build before, the leveling build after:

| Spec | alliance % | horde % | seed 1 alliance / horde |
|---|---|---|---|
| Survival, raid build → leveling build (X3 included) | -2.61 / -2.62 | -2.73 / -2.73 | 842.0 → 820.1 / 848.5 → 825.3 |

The ruling alone, with the engine held fixed, is -0.63 / -0.64% (alliance) and -0.75 / -0.76% (horde); the rest is X3.

Leveling bands, bare (only bands that move):

| Spec, band | alliance % | horde % | seed 1 alliance / horde | cause |
|---|---|---|---|---|
| Fury 20 | -6.03 / -6.14 | -5.53 / -5.57 | 36.96 → 34.73 / 37.05 → 35.00 | W1, W2 (W3 recovers part) |
| Fury 30 | -6.10 / -5.98 | -5.98 / -6.03 | 65.38 → 61.39 / 65.71 → 61.78 | W1 |
| Fury 40 | -0.05 / -0.05 | -0.05 / -0.04 | 127.38 → 127.32 / 126.71 → 126.65 | Gut Ripper |
| Fury 50 | +1.06 / +1.06 | 0.00 / 0.00 | 191.69 → 193.73 / 194.69 | Hammer of the Northern Wind (alliance set) |
| Arms 20 | +4.69 / +4.65 | +4.55 / +4.63 | 25.37 → 26.56 / 25.69 → 26.86 | W2: no Berserker swaps that bled rage |
| Arms 40 | +0.64 / +0.64 | +0.64 / +0.65 | 96.01 → 96.62 / 95.03 → 95.64 | Nightblade |
| Assassination 30 | +2.05 / +2.07 | +0.53 / +0.64 | 57.00 → 58.17 / 56.57 → 56.87 | Claw, Thornspike |
| Assassination 40 | -0.04 / -0.03 | +0.50 / +0.51 | 97.74 → 97.70 / 97.58 → 98.07 | Gut Ripper; horde Coldrage |
| Combat 30 | +0.50 / +0.50 | +1.61 / +1.65 | 64.43 → 64.75 / 64.75 → 65.79 | Claw, Thornspike |
| Combat 40 | -0.04 / -0.03 | +0.35 / +0.35 | 109.71 → 109.67 / 109.10 → 109.48 | Gut Ripper; horde Coldrage |
| Subtlety 30 | +0.53 / +0.57 | +0.60 / +0.49 | 59.89 → 60.21 / 58.64 → 58.99 | Claw |
| Subtlety 40 | -0.04 / -0.03 | +0.38 / +0.38 | 94.76 → 94.72 / 94.21 → 94.57 | Gut Ripper; horde Coldrage |
| Subtlety 50, Survival 40 | -0.03 to -0.04 | -0.03 to -0.04 | | Gut Ripper |

Arms 30 and every band below 60 not listed: no change.

**Round 7's Survival drop, by cause.** On cd4dfdcc0, with Strength of Earth put back to Era's 61 and nothing else changed, Survival's band-60 raid number returns exactly to round 7's before: 775.0 / 774.7 (alliance) and 778.4 / 778.5 (horde). So the totem explains all of round 7's -0.7%, for both factions alike; the raid preset gives alliance the totem too.

## For other lanes

- **Hybrids, healers and the integrator:** X3, Curse of Recklessness at 505 armor with no attack power, moves every physical attacker on the raid preset by about -2%. That includes Enhancement, Retribution, Feral cat and bear, and Protection paladin threat, and fork goldens TestEnhancement, TestFeralTank and TestProtection (paladin). Re-measure them at the pin.
- **Hybrids, healers and the items lane:** X1 changes proc damage on 26 weapons. No published set outside this group wears one, but ranker alternatives across classes do.
- **Healers:** Innervate (29166) is BaseLevel 40, but the engine never enforces `RequiredLevel`, and spellranks carries it as a rank-0 row, so a band-20 or band-30 Restoration druid may cast it. Check the engine and add it to singleTierLevelOverrides if so (X5).
- **Integrator:** TestRotationLadder rows move for warrior-arms and warrior-fury (W1 to W3) and for every physical spec's raid row (X3). They were not regenerated here.

## Still open

- Whether aura 674 armor reductions stack (X4).
- Fury's stance at 30 to 35 (W4).
- Gloves of Holy Might, Toxic Revenger, and the below-25% weapon effects (X7).
- Carried rows: rage from damage taken, Heroic Strike as a swing, Battle Shout rounding, Eviscerate and Rupture attack power, Hawk's proc mask, Serpent Sting's spell power, the pet rows and threat values. These need launch logs.

Sources (read 2026-11-01): the layered 70707 raw tables listed above; data/builds/1.15.9.69722/raw/SpellEffect.csv; https://nether.wowhead.com/forever/tooltip/spell/11717 and /classic/tooltip/spell/11717.
