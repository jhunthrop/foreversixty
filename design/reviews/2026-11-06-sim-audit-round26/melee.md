# Sim-accuracy round 26: melee (fresh audit of Subtlety and Arms)

Site `wt/r26-melee` from int/r26 86d239f85. Fork `lane/r26-melee` from int/r26 f871a65ae. Client 1.60.1.70790. BUILD MODE.

**Result: 3 confirmed bugs, all fixed (origin: old).** One engine bug: Ghostly Strike ignored the normalized-weapon effect the 70755 client added to it. Two leveling-order bugs: Subtlety took dead toll points (Camouflage, Elusiveness) at 24 to 28 while Initiative waited, and Arms took Improved Tactical Mastery at 19 to 21, which does nothing before 50, while Improved Overpower waited. Everything else I read in both specs, raid and leveling, is aligned or stays open on a Forever log.

Measurements use frozen BiS (`$R/bis-frozen`) and guide builds (`$R/rt`). Tools and outputs are in `$R/melee/{tools,rot,out}/`. Subtlety has one guide build (`005302011-320003-035230031321001151`). Arms has a level 60 build (`13323213032215201-0505-205`) and a raid build (`30303213132215201-05050005012-1`).

## Subtlety

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Ghostly Strike used the weapon's own speed.** The 70755 client gave Ghostly Strike 14278 a new effect 121 (normalized weapon damage, SpellEffect 1368764, amount 0). It sits beside its effect 31 at 125% and the dagger dummy at 180%, the same shape as Hemorrhage (693307 effect 121 under 1231002 effect 31), and the engine already normalizes Hemorrhage. The same patch added the opener trigger (1368766, effect 64 to 1329326), which the engine models. The engine read `MHWeaponDamage`, so on Subtlety's 2.9 s Teebu's Blazing Longsword every Ghostly Strike counted attack power at 2.9 s instead of 2.4 s (and at 2.1 s instead of 2.4 on the band 50 mace). Blizzard's 70755 notes (post 9) also changed Ghostly Strike and Hemorrhage damage ("dealing much more damage than intended when using a Dagger"). Rounds 7, 14 and 23 read Ghostly Strike as "effect 31 only, no 121", which was true on the 70009 to 70532 tables (cached `raw/SpellEffect.csv`) but not since 70755. `TestGhostlyStrikeUsesNormalizedWeaponDamage` failed first (332.9 against the normalized 304.3), then passed. | `b70790/70790/SpellEffect.csv` rows 1368764, 1368766; `b70790/spelldiff.txt`:61-63 ("ADDED"); `spellconst/rogue.json` 14278; `ghostly_strike.go`:80; https://us.forums.blizzard.com/en/wow/t/2360696 post 9 | CONFIRMED-BUG | old (70755 follow a57b0221b; rounds 14 and 23 read the old rows) | E 414da219b |
| S2 | **Leveling order: Camouflage and Elusiveness before Initiative.** The walk spent 24 on Camouflage, 25 on Elusiveness and 27 and 28 on Camouflage. Neither does anything in a fight (the engine models neither, and the card never casts Vanish). They only pay the 15- and 20-point row tolls. Initiative (a combo point on every Garrote, which the card casts after every Ghostly Strike from 26) waited until 29, 31 and 32. Initiative is now held from 24: two points before Ghostly Strike, which still comes at 26, and the third at 27. Hemorrhage still comes at 30. Bare, 10,000 iterations, seeds 1 / 50001: level 25 +0.41% / +0.41% (both factions); level 30 (same engine, old against new talents) Horde 90.6 → 91.4 / 90.4 → 91.4, Alliance 90.0 → 90.9 / 89.9 → 90.8 (+0.9 to +1.1%). No other checkpoint moves. Red first: `TestSubtletyTakesInitiativeBeforeTheTollPoints`. | `out/ladder-old.txt`, `out/ladder-new.txt`, `out/ladder-measure.txt`; `data/curated/ladder-first.json` | CONFIRMED-BUG | old (round 13 walk lane reservation) | S cf3be3516 |
| S3 | Raid-build points re-read on the 70790 text. Assassination 12: Malice 5, Ruthlessness 3 (60%), Improved Slice and Dice 2 (+30%), Blood Splatter (Rupture +6 s, Garrote +9 s, `client_values.go`:110-111), Lethality 1 (+10% crit bonus on Hemorrhage and Ghostly Strike, `CritDamageBonus` 1.1). Combat 8: Coup de Grace 3 (Death Mark +20% on cap and storage rate, 1369147), Efficient Strikes 2 (Hemorrhage 30 energy: 1108.6 / 36.95), Precision 3. Subtlety: Opportunity 20%, Dirty Deeds -20 (Garrote 30, 20 under Poised), Initiative 100%, Serrated Blades 30% Rupture and 9% armor, Boot Knife 2 points per Shank (37.4 a fight on 18.7 Shanks), Death Mark rank text (156 + 245 x cp, 0.04 AP a point, crit, coup). Swapping Lethality for the third Improved Slice and Dice point gains +0.09% (799.0 against 798.3 / 798.5, 3,000 iterations): under the 1% bar, and Lethality does something. Vigor in its place loses. | `talents/rogue.json`; `out/sub-raid-h1.txt`; talent variants | Aligned | - | - |
| S4 | Poison hands. Deadly main hand / Instant off hand 798.3 / 798.5 against swapped 782.0 / 780.1 (-2%). The kit's stone beside the poison is the Forever rule (`stone_beside_poison_test.go`, Deep Dive 7 Oct). | 3,000 iterations, seeds 1 / 50001 | Aligned | - | - |
| S5 | Ghostly Strike line at 26 to 29 on the band 20 table (a mace, before Hemorrhage), after S1: card 57.7 / 61.0 at 26 / 28, without Ghostly Strike 48.4 / 51.8. The Garrote it lets in keeps the line far ahead. | `rot/sub-nogs.json` | Aligned | - | - |
| S6 | Card notes after S1. "From 30 harder than Hemorrhage with any weapon": both now normalize, so Ghostly Strike's 125% (180%) beats Hemorrhage's 100% (145%) on the same weapon at any speed. Round 24 guide S4's 1.92 s caveat is gone. "A slow sword's bigger hits make up for it" (subtlety.md:105, index.md:26) still holds: with the sword at 1.25 x (137 + 2.4 AP/14) against the dagger at 1.8 x (76.5 + 1.7 AP/14), the sword stays ahead to about AP 7800. | arithmetic on `items/rogue.json` | Aligned | - | - |
| S7 | Spells and ranks the card casts per band. Bare breakdowns at 30 and 40 show the trainer ranks (Garrote 8632/8633, Rupture 8639/8640, Slice and Dice 5171, Deadly Poison 2823/2824, Shank 1329283) and Death Mark rank 1 at 40. Round 24's 297-run cast-id sweep stands for the spells S1 does not touch. | `out/`; round 24 R3 | Aligned | - | - |

## Arms

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | **Leveling order: Improved Tactical Mastery before Improved Overpower.** The walk spent 19 to 21 on Improved Tactical Mastery. It only keeps rage through a stance change, and the Arms card changes stance only for Recklessness, from 50, so below 50 its points do nothing. Improved Overpower (Overpower +25% crit a point, cast from 12) waited until 22 and 23. Improved Overpower is now held first, at 15 and 16 (Improved Rend finishes at 17 and 18, Improved Charge at 19 and 20). Tactical Mastery still comes at 21 to 23, as Anger Management's prerequisite at 24. Bare level 20, 10,000 iterations, seeds 1 / 50001: Alliance 30.1 → 31.1, Horde 30.6 → 31.6 (+3.3% both). Checkpoints 10 and 25 onward do not move. Red first: `TestArmsTakesImprovedOverpowerBeforeTacticalMastery`. | `out/ladder-*.txt`, `out/bare-before-after.txt`; Blizzard notes post 9 (Improved Tactical Mastery 5/10/15) | CONFIRMED-BUG | old (walk order since the 70755 trees) | S cf3be3516 |
| A2 | Improved Tactical Mastery in the raid build. It is Anger Management's prerequisite (the planner refuses `30320213...`: "Anger Management needs 3 points in Improved Tactical Mastery"), and the execute swap keeps 25 rage with it. | spec-breakdown error text; `stances.go`:28-47 | Aligned (toll) | - | - |
| A3 | Raid-build points on the 70790 text. Improved Heroic Strike 3 (12 rage: 87.5 / 7.29), Improved Rend 1.35, Improved Overpower +50% (Overpower 76.2% crit against 26.5% on other specials), Deep Wounds 60% weapon average, rolled over, no attack power (Blizzard #234, notes post 5), Two-Handed Weapon Specialization 6%, Impale +20% crit bonus, Bloodthrill 20% (trait curve 109366; ProcTypeMask 20, auto and special, "includes Cleave and Heroic Strike" per post 1), Weaponmaster mace 15% armor (The Unstoppable Force), Improved Slam (cast and global -0.5 s, cooldown 15 → 10 s; 12862 effects 107 misc 10/21/11), Reckless Swings (4 white-swing charges, ProcTypeMask 4, 15 s, 1329555), Improved Execute -5, Cruelty 5, Unbridled Wrath 60% on white swings. Enrage 5 never procs on a DPS warrior (ProcTypeMask 139944, taken hits) and is the 15-point Fury toll: every other point in Fury rows 1 to 3 is just as idle (round 24 W5). | `spellconst/warrior.json`; `b70790/70790/SpellAuraOptions.csv`, `SpellEffect.csv`; `out/arms-raid-h1.txt` | Aligned | - | - |
| A4 | Card order. Spearing Strike ahead of Slam loses 0.5% (842.6 / 842.9 against 846.7 / 846.4, 3,000 iterations). Slam (effect 17, unnormalized) and Spearing Strike (effect 121, normalized 75%, Battle Stance only, 20 s) follow the client. | `rot/arms-spear-first.json` | Aligned | - | - |
| A5 | Overpower: a window from any dodge (5 s) or Bloodthrill (1282733, 5 s), cannot be dodged, 5 s category cooldown, normalized + 35. An Overpower's own Bloodthrill proc is cleared by the cast that made it (`overpower.go`: Deactivate after the hit). The new window would close as the 5 s cooldown ends, so it is worth almost nothing. | `overpower.go`; client 11585, 1282733 | OPEN (low; a log with an Overpower right after an Overpower) | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Blizzard's Forever notes thread, read 6 Nov 21:44 UTC: 10 posts, the last is 4 Nov 19:35 UTC (no new post). I re-read every Subtlety and Arms line (posts 1, 4, 5, 6, 9) against the engine. Only S1 differed. | https://us.forums.blizzard.com/en/wow/t/2360696.json | Aligned | - | - |
| X2 | 70790 server hotfixes: no `hotfix-push.json` row for any spell or talent either spec casts or takes. | `hotfix-push.json` | Aligned | - | - |
| X3 | Logs, last looked 6 Nov 21:44 UTC: api.foreversixty.gg `/v1/reports/recent` lists only the 20 Oct report. | curl | OPEN | - | - |
| X4 | The published BiS races (Arms orc / human, Subtlety troll / night elf) differ from the guide races (night elf / undead). The race lane judged on the guide races (round 25 race.md); the BiS files take the race at the re-rank. | `bis-frozen/*.json` | Aligned (re-rank) | - | - |

## Before/after

Raid 60, guide builds, frozen BiS, 10,000 iterations, against `$R/before.tsv` (seed 1 / 50001):

| spec | faction | before | after | change |
|---|---|---|---|---|
| rogue-subtlety | alliance | 808.7 / 809.4 | 804.1 / 804.8 | -0.57 / -0.57% |
| rogue-subtlety | horde | 802.9 / 802.8 | 798.4 / 798.3 | -0.56 / -0.56% |
| warrior-arms, rogue-assassination, rogue-combat | both | as before.tsv | unchanged | 0 |

Bare (before = `$R/sb-before`, after = this lane), seeds 1 / 50001, both factions agree within 0.1 point:

| spec | level | Horde before → after (seed 1) | change |
|---|---|---|---|
| rogue-subtlety | 25 | 48.2 → 48.4 | +0.41% (S2) |
| rogue-subtlety | 30 | 90.8 → 91.4 | +0.66 to +0.78% (S2 and S1) |
| rogue-subtlety | 40 | 152.8 → 152.6 | -0.06 to -0.13% (S1) |
| rogue-subtlety | 50 | 224.2 → 224.7 | +0.22 to +0.27% (S1, a 2.1 s mace) |
| rogue-subtlety | 60 | 383.3 → 381.8 | -0.37 to -0.39% (S1) |
| warrior-arms | 20 | 30.6 → 31.6 | +3.27 to +3.32% (A1) |

No spec outside rogue and warrior moves. The engine commit touches only `sim/rogue/ghostly_strike.go`, and the ladder rows touch only these two specs.

## Files to regenerate

- `sim/request/testdata/ladder/rogue-subtlety.golden.md` and `warrior-arms.golden.md` (`TestRotationLadder`). Until then `web/src/lib/guides/leveling-builds.test.ts` fails on `warrior-arms` level 20 and `rogue-subtlety` 25 / 30, because it compares the ladder with those goldens.
- `addon-builds.json` (the Subtlety and Arms leveling orders), then `Data.lua`.
- Subtlety `dps-per-point` / stat weights (S1), and the Subtlety 25 / 30 and Arms 20 rows (S2, A1).
- Coach starters (`persona_starter.py --all`) once `addon-builds.json` is regenerated, if they read talent levels.
- Hand copy, not a regen: `sim/leveling/ladder-first.json` is a byte copy of `data/curated/ladder-first.json` (`TestEmbeddedLadderFirstMatchesCurated`).

## Re-rank needed

- **rogue-subtlety** (every band). S1 changes how much Ghostly Strike gets from weapon speed, so main-hand picks can move: slow weapons lose a little and weapons under 2.4 s gain. S2 also changes the band 30 build.
- **warrior-arms** band 20: A1 changes its build.

## Measure at refresh

- Subtlety raid 60 and bands on the re-ranked BiS (S1). The carried items are one-hand axes for rogues and the Subtlety Hemorrhage gate and 50-energy Shank on re-ranked gear.
- Arms band 20 on its new build.

## For other lanes

- **Ranks / platform (judge on the new below-20 gear):** Subtlety's walk puts Master of Deception (stealth detection only) at 10 to 12 before Opportunity (+4% a point to Backstab and Garrote) at 13 to 17. Taking Opportunity first changes checkpoint 10 only, so it is a ladder-first row that needs a measured gain at level 10. I left it because the below-20 gear is changing. Arms takes Deflection at 11 to 13 before Improved Rend. That moves no checkpoint, so it would be a `ladder-order.json` row decided by leveling research.
- **Coach:** `tools/coach/starter_overrides.py`:489's comment says "Initiative full only from thirty-two". It is now 27. I did not touch it (coach tools are off limits).

## For the guide lane

- Done here: arms.md:125 and :131 (Improved Overpower at 15 and 16, Improved Rend 17 and 18, Improved Charge 19 and 20; the "Improved Overpower comes in at 22" clause dropped) and subtlety.md:135 (Initiative at 24, 25 and 27; Camouflage from 28). `_prose.test.ts` and the guide tests pass.
- No card note changes: S6.

## Guide text not done

- None.

## Still open

- Rogue: Death Mark storage (which hits it stores), the Shank and Mutilate off-hand penalty, poison coefficients and Backstab's flat bonus on a live log. Warrior: Defiance 7/13/20, Revenge, Shield Slam and Thunder Clap threat, Deep Wounds' double attacker multiplier, Rend and Thunder Clap attack power (N2), A5. Owner: level-60 book ranks, WARRIOR-13, Combat Eviscerate at five.

## Sources

- Client 1.60.1.70790: `spellconst/{rogue,warrior}.json`, `spelltips/rogue.json` (1325758, 14278), `talents/{rogue,warrior}.json`, `hotfix-push.json`. Wago tables cached at `scratchpad/b70790/70790/{SpellEffect,SpellAuraOptions}.csv`, and the 70755 diff `scratchpad/b70790/spelldiff.txt`. Earlier tables: `data/builds/1.60.1.{70009,70532}/raw/SpellEffect.csv`.
- Blizzard, WoW Forever Beta Development Notes, posts 1, 4, 5, 6, 9: https://us.forums.blizzard.com/en/wow/t/2360696 (read 6 Nov 21:44 UTC).
- Round 24 melee, round 23 rogue and warrior, guide-accuracy round 26 rogue and warrior.
