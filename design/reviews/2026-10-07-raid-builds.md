# Raid builds: a level-60 build for the raid-ready preset

2026-10-07. Lane `raid-builds`; site branch `raid-builds`, fork branch `raid-builds` (the fork side is the seven synced rotation copies under `ui/`, nothing else).

## Decision

Owner, 2026-10-07: "make the best decision for the accuracy of the sim". The raid-ready preset exists to show the outcome when the playing field is level, so its level-60 entry is simmed on a **raid build**, not the leveling guide build. The leveling bands (20 to 50) and the bare band 60 keep the guide's leveling build.

### Data model

- A spec guide carries an optional `raidBuild:` FS1 line beside `build:` (`web/src/content.config.ts`). A guide without one has no separate raid build: the leveling build is the raid build.
- `leveling.GuideRaidBuildTalents` reads it (falling back to `build:`). The BiS ranker (`sim/cmd/leveling-bis`, `raid_build.go`) spends the raid pass at band 60 from it and every other pass from the leveling build; the raid entry's published `talents` is therefore the raid build's code. Tests: the leveling reader, the pass selection, and a stale-stamp refusal.
- The guide's Talents section shows both builds under "Leveling build" and "Raid build" headings, each with its tree and its planner and sim links, or one tree under "Leveling and raid build" when they are the same (`GuideBuilds.astro`, `guide-builds.ts`). `/bis` keeps only its planner link, which already reads the published `talents`; it states no talents.
- `builds.test.ts` runs every legality check on a raid build too, and the rotation-signature table applies to it.

## The rule

Applied the same way to all twenty DPS specs. The search is `talent-search -spec <spec> -preset raid` from the current guide build, on the committed level-60 raid-preset gear. The winner becomes the raid build when:

1. its own tree holds more points than either other tree;
2. it keeps every talent the curated rotation casts or gates on (the builds test's signature table, plus the ids the APL names);
3. it keeps every talent that matters in a raid but that the engine cannot model: threat reduction, the raid-wide auras and debuffs the preset assumes from the spec's own tree (Trueshot Aura, Leader of the Pack, Moonkin Form, Improved Scorch, Winter's Chill), survivability cooldowns and passive mitigation;
4. it MAY drop utility the engine cannot model that a raid DPS does not use (movement, stealth, crowd control, PvP, out-of-combat, solo-leveling), each classified below;
5. it gains at least 1% and is beyond the search's combined error (the 2026-10-06 adoption rule's own wording, "gain beyond error and at least 1%").

Rules 2 to 4 are enforced inside the search, not by filtering its output: `talent-search` gained `-keep` (a list of talent names no candidate may take a point from; deep archetypes start from the guide stripped to them), and the per-spec lists are in `design/reviews/talent-search/raid-keeps.json`. A modeled talent needs no protection: the sim prices every candidate whole. An unmodeled talent that is not in rule 4's classes stays. If nothing qualifies the raid build is the guide build.

The reports are `design/reviews/talent-search/<spec>-raid.md`. For the five specs with no qualifying winner the search was rerun with `-cap 150 -top 25` to be sure the finalist list was not hiding one.

## Per spec

Builds are the digits of the FS1 code (tree 1 / tree 2 / tree 3). "modeled" in the dropped list means the engine reads the talent and the sim prices it, so the trade is measured; the bracket otherwise is the utility class of rule 4. Gain is the search's, against the guide build, on the raid preset, with its combined error in percentage points.

| Spec | Guide build | Raid build | Moved (added / dropped) | Gain | Rotation change |
|---|---|---|---|---|---|
| druid-balance | `5232221115400051/05/503301` | `4132220115501051/05/5053` | +1 Improved Starfire; +1 Nature's Grace; +2 Naturalist / -1 Improved Wrath [modeled]; -1 Genesis [modeled]; -1 Improved Entangling Roots [crowd control]; -1 Reflection [modeled] | +8.1% (±0.4) | none |
| druid-feral | `01/55232032121032012001/50532` | `05002/45211031021032212001/5053` | +4 Genesis; +2 Nature's Majesty; +2 Improved Shifting Power / -1 Ferocity [modeled]; -2 Feral Instinct [stealth (and a Swipe bonus the rotation never casts)]; -1 Brutal Impact [crowd control]; -1 Savage Fury [modeled]; -1 Feral Charge [crowd control]; -2 Natural Shapeshifter [modeled] | +10.2% (±0.2) | Rake dropped (verified, see Rotation) |
| hunter-beast-mastery | same | same | none: the guide is the best qualifying build (best other finalist -0.6%) | - | none |
| hunter-marksmanship | same | same | none: best qualifying finalist +0.6%, under the 1% bar | - | none |
| hunter-survival | same | same | none: best finalist +0.2% and outside the spec's own tree (0/16/35) | - | Wing Clip added, gated on the Windfury Totem |
| mage-arcane | `153005113100011531/032023/055` | `153005113100011531/03/05450003` | +5 Ice Shards; +3 Piercing Ice / -2 Improved Fireball [modeled]; -2 Flame Throwing [solo-leveling (Fire range on a mage who casts Arcane and Frost)]; -3 Impact [crowd control]; -1 Elemental Precision [modeled] | +1.5% (±0.3) | none |
| mage-fire | `2050151/23552100130103051/005` | `2050151/23252100130133051/005` | +3 Master of Elements / -3 Improved Fireball [modeled] | +1.1% (±0.4) | none |
| mage-frost | `203005/113023/253510130100030025` | `203005/13102/255510032100030025` | +2 Incineration; +2 Elemental Precision; +2 Frost Channeling / -2 Improved Fireball [modeled]; -3 Impact [crowd control]; -1 Frostbite [crowd control] | +3.4% (±0.3) | none |
| paladin-retribution | `52003003/052/05025331001330311` | `54003/053/0520533100133032` | +2 Divine Intellect; +1 Precision; +2 Improved Judgement; +1 Instrument of Law / -3 Reverence [modeled]; -2 Holy Conduit [modeled]; -1 Twist of Light [modeled] | +1.8% (±0.3) | none |
| priest-shadow | `5240110313/0/543120301201300051` | `3250010313/0/523120501201300251` | +1 Twin Disciplines; +2 Improved Mind Blast; +2 Early Demise / -2 Power in Light [solo-leveling (Holy damage before Shadowform)]; -1 Holy Precision [solo-leveling (Holy hit before Shadowform)]; -2 Blackout [crowd control] | +1.6% (±0.1) | none |
| rogue-assassination | `32502110551501001/302303/512` | `01532310421501/315303000015/002` | +3 Ruthlessness; +2 Improved Slice and Dice; +1 Improved Sinister Strike; +3 Lightning Reflexes; +1 Flawless Execution; +5 Dual Wield Specialization / -3 Improved Gouge [crowd control]; -1 Remorseless Attacks [solo-leveling (kill-chain bonus)]; -1 Lethality [modeled]; -3 Vile Poisons [modeled]; -1 Venom [modeled]; -5 Camouflage [stealth]; -1 Master of Deception [stealth] | +8.9% (±0.2) | Slice and Dice ahead of Eviscerate; Venom dropped (verified) |
| rogue-combat | `32531/32530300001515201/51` | `005320105/31530300001515231/002` | +1 Murder; +1 Relentless Strikes; +5 Lethality; +3 Aggression; +2 Opportunity / -3 Improved Gouge [crowd control]; -2 Remorseless Attacks [solo-leveling (kill-chain bonus)]; -1 Improved Sinister Strike [modeled]; -5 Camouflage [stealth]; -1 Master of Deception [stealth] | +15.9% (±0.2) | Backstab ahead of Adrenaline Rush |
| rogue-subtlety | same | same | none: the only finalist beyond error (+2.3%) is deep Assassination, 29/0/22, outside the spec's own tree | - | none |
| shaman-elemental | `553031130010305/03/553302` | `5530311300103051/052/053302` | +1 Lava Burst; +2 Thundering Strikes; +2 Ancestral Knowledge / -5 Improved Healing Wave [solo-leveling (self-heal)] | +3.1% (±0.4) | none |
| shaman-enhancement | same | same | none: the guide is the best build found | - | none |
| warlock-affliction | `25552300120201351/200522/003` | `25550300100201351/00052/0550001` | +5 Improved Shadow Bolt; +2 Bane; +1 Ruin / -2 Soul Harvest [solo-leveling (mana on kills)]; -2 Pandemic [modeled]; -2 Improved Health Funnel [solo-leveling (pet upkeep)]; -2 Improved Voidwalker [solo-leveling (pet tanking)] | +11.4% (±0.2) | Shoot below 50% mana (was 40%) |
| warlock-demonology | `25532/233211310122000135/004` | `055/233511310102000135/055` | +3 Unholy Power; +5 Improved Shadow Bolt; +1 Bane / -2 Improved Life Tap [modeled]; -3 Malediction [modeled]; -2 Soul Harvest [solo-leveling (mana on kills)]; -2 Master Summoner [modeled] | +9.7% (±0.2) | Soul Fire added |
| warlock-destruction | `25532/0/2053225103101351` | `255/0005/2053045103101351` | +5 Unholy Power; +2 Aftermath / -3 Malediction [modeled]; -2 Soul Harvest [solo-leveling (mana on kills)]; -2 Cataclysm [modeled] | +2.9% (±0.2) | none |
| warrior-arms | `03325213032515001/0505/005` | `02305213032515001/55050000001/2` | +5 Booming Voice; +1 Improved Execute; +2 Improved Bloodrage / -1 Deflection [modeled]; -2 Improved Charge [modeled]; -5 Iron Will [PvP (stun and fear duration)] | +1.3% (±0.3) | Hamstring added, gated on the Windfury Totem |
| warrior-fury | same | same | none: the guide is the best build found | - | none |

Three things the table shows. The big winners (Feral, Assassination, Combat, Affliction, Demonology, about 9 to 16%) are specs whose guide build still spends points the engine measures at nothing; the leveling build is a reasonable build for a character that levels, and the raid search does not argue with that. The classified drops are all low-tier fillers: crowd control and stealth rows and solo-leveling returns. And Warlock Demonology and Destruction lost their Suppression drops to the keep list (the first pass dropped five points of it, which is threat reduction), at a cost of about 5 points of gain for Demonology.

## Rotation

The single curated rotation serves both preset entries and the ladder, so a line measured at the raid build is adopted only when it also leaves the bare guide build level (`rotation-search`, bare, guide build, before and after) and, for any ability added, dropped or reordered, when the engine's model of it matches the client tables. `rotation-search -preset raid -build-code <raid build>` was run for the fourteen specs whose raid build differs, and for Survival at its guide build.

### Venom (Assassination)

Client rows for spell 1310703 against `sim/rogue/venom.go` (client tables in `data/builds/1.60.1.70009/raw`):

| Quantity | Client | Engine | |
|---|---|---|---|
| Cost | SpellPower: 25 energy; and one combo point minimum (power type 4, cost 1) | EnergyCost 25; needs a combo point | match |
| Global cooldown | SpellCooldowns StartRecoveryTime 1000 | GCD 1 s, ignores haste | match |
| Duration | DurationIndex 185: 6000 ms base + 3000 per combo point, max 21000 | 9, 12, 15, 18, 21 s for 1 to 5 points | match |
| Poison damage | Spell.csv "$m2%"; effect 1: aura 108 (percent modifier), +30, damage, class mask 0x2000 (every Instant Poison rank) | x1.30 on Instant Poison | match |
| Poison periodic damage | effect 2: aura 108, +30, misc 22 (periodic), mask 0x10000 (every Deadly Poison rank) | x1.30 on the Deadly Poison tick | match |
| Poison apply chance | effect 3: aura 107 (flat modifier), +10, misc 18 (chance of success), mask 0x1001E000 (Instant, Deadly, Crippling, Mind-numbing, Wound) | +0.10 added to Instant, Deadly and Wound | match for the poisons the kit uses |
| Proc | SpellAuraOptions: no row | none | match |
| Talent tree | `talents/rogue.json` node: single rank, Mutilate prerequisite, tier 6 | `Talents.Venom`, gated registration | match |

Two deviations, both documented here for the engine lane. (1) Vile Poisons (spell 16513) carries the same percent-modifier ops and class masks, and the client adds percent modifiers on one op, so Instant and Deadly damage with Vile Poisons 5/5 and Venom is x(1 + 0.20 + 0.30) = 1.50; the engine multiplies, 1.20 x 1.30 = 1.56. Venom's marginal poison damage is therefore modeled at +30% where the client gives +25%: the sim flatters Venom by about a fifth of its damage term. (2) The engine also applies x1.30 to Wound Poison, which deals no damage in the client; the rogue kit carries Instant and Deadly, so it has no effect.

The model matches the client on every number that decides whether to cast it, and its one deviation favours Venom, so the sim's verdict against the line is a floor. Measured: with the Venom line removed, the bare guide build gains 1.3% (the earlier lane's 1.6% held), and at the raid build under the raid preset the baseline never reaches it (zero casts: Eviscerate takes every set of four combo points). **Adopted**: the Venom line is removed from `data/curated/apl/rogue-assassination.json`, together with the search's one winning mutation at the raid build, Slice and Dice ahead of Eviscerate (+7.0 DPS of 637, beyond error). With the line gone the talent is a spare point, so the second talent search ran with Venom unprotected and the raid build drops it (rotation signature table updated). The leveling build keeps the point; whether the guide's leveling build should drop it too is the guide owner's call.

### Other drops and additions

| Spec | Line | Raid gain | Bare guide build | Model check | Decision |
|---|---|---|---|---|---|
| Feral | remove Rake | +8.9% | +5.1% | client 9904: 61 up front, 34 x 3 ticks, 40 energy, one combo point, 9 s; engine reads the same generated table and tick row | **adopted** (drop, verified) |
| Survival | add Wing Clip | +9.8% (found at the first-pass raid build) | -6.0% ungated | client 14268: 50 damage, 80 mana, 1.5 s global; engine matches, and `TestWingClipDamageMatchesClient` pins it | **adopted**, gated on the Windfury Totem (see below) |
| Arms | add Hamstring | +7.3% (first-pass raid build) | -5.3% ungated | client 7373: 45 damage, 10 rage (the 27584 row is a free reissue); engine matches, `TestHamstringDamageMatchesClient` | **adopted**, gated on the Windfury Totem |
| Demonology | add Soul Fire; Bane of Doom after Corruption | +10.3% (first-pass raid build) | +6.1% | client 17924: 335 mana, 6 s cast, 60 s cooldown, 431 +1.9 a level, coefficient 1.0; Decimation 2/2 in both builds cuts the cooldown 90% | **adopted** |
| Assassination | Slice and Dice ahead of Eviscerate; Venom | +1.3% | +1.6% | see above | **adopted** |
| Combat | Backstab ahead of Adrenaline Rush | +0.4% | +0.2% | order only | **adopted** (beyond error, level on bare) |
| Affliction | Shoot below 50% mana | +0.6% | +0.3% | cost-free wand | **adopted** (beyond error, level on the fixed-gear bare search) |
| Retribution | add Consecration | +1.7% | level | the client states no Consecration damage (periodic dummy, server-side script); the engine uses Classic's table | **held**: unverified damage |
| Shadow | Shadow Word: Death without its execute gate | +2.7% | level | client states a backlash of 10% of maximum health (effect 3); the engine removes the damage dealt and the sim has no health to price either | **held**: unpriced self-damage, and the engine's backlash differs from the client's |
| Shadow | remove the Inner Focus and Devouring Plague sequence | +0.3% | level | removing it alone measures -0.1 +- 0.5 | **held**: no effect (the line is dead for a gnome) |
| Destruction | replace Bane of Doom with Curse of the Elements | +0.6% | - | client 603: 1742, coefficient 4.0, 300 mana, 60 s; engine matches (`TestBaneOfDoomDeclaresTheLearnableClientRow`) | **held**: the same run's action probe prices Bane of Doom at +51 DPS removal-loss, so the gain is the raid preset's own Curse of the Elements colliding with the cast one, not a better rotation |
| Survival | Lacerate refresh below 3 s | +0.4% | -0.5% | order of refresh only | **held**: wins the raid, loses the bare build |

The Wing Clip and Hamstring gains are Windfury Totem procs: the engine rolls the totem's 20% extra attack off every main-hand hit including specials (`sim/core/buffs.go`), so a cheap special is a proc source. Added ungated they lose 5 to 6% on the bare character, which has no totem and little mana or rage. The lines are gated on `auraIsKnown(10610)`, the totem's attack power buff, which the engine registers only for a character the totem reaches, so the bare character skips them and the raid character plays them; both rotations measure level on the bare build and keep the raid gain. The rotation-search probe, the ladder harness and the data lane's spell check each needed to learn about the id (`inert`, `expected_idle`, `ENGINE_AURA_IDS`). If the real client does not roll Windfury off specials, these two lines are wrong, and so is the engine's totem model.

## BiS ranker, before and after

Band 60, one `leveling-bis -spec <spec> -bands 60` run per spec per side into a scratch directory; before is `main`, after is this branch. Set DPS of the published picks. The raid columns move by the build and rotation changes together and by the gear the ranker re-picks for them; the bare column is the guard.

| Spec | Raid A before | Raid A after | Δ | Raid H before | Raid H after | Δ | Bare A before | Bare A after | Δ |
|---|---|---|---|---|---|---|---|---|---|
| druid-balance | 477.7 | 507.6 | +6.3% | 477.7 | 512.1 | +7.2% | 220.5 | 220.5 | +0.0% |
| druid-feral | 510.2 | 617.9 | +21.1% | 513.1 | 620.7 | +21.0% | 286.6 | 300.1 | +4.7% |
| hunter-survival | 620.3 | 686.4 | +10.7% | 618.1 | 686.8 | +11.1% | 224.6 | 224.6 | +0.0% |
| mage-arcane | 698.9 | 709.3 | +1.5% | 683.4 | 693.6 | +1.5% | 484.6 | 484.6 | +0.0% |
| mage-fire | 737.7 | 743.7 | +0.8% | 724.4 | 737.6 | +1.8% | 403.5 | 403.5 | +0.0% |
| mage-frost | 548.5 | 563.5 | +2.7% | 545.5 | 563.1 | +3.2% | 323.6 | 323.6 | +0.0% |
| paladin-retribution | 515.7 | 527.0 | +2.2% | 527.5 | 543.5 | +3.0% | 247.0 | 247.0 | +0.0% |
| priest-shadow | 581.9 | 591.4 | +1.6% | 593.7 | 602.9 | +1.5% | 299.4 | 299.4 | +0.0% |
| rogue-assassination | 574.9 | 646.4 | +12.4% | 569.1 | 643.8 | +13.1% | 268.4 | 275.2 | +2.5% |
| rogue-combat | 549.2 | 623.1 | +13.4% | 547.9 | 618.6 | +12.9% | 253.3 | 257.3 | +1.6% |
| shaman-elemental | 411.5 | 423.0 | +2.8% | 412.4 | 425.5 | +3.2% | 140.2 | 140.2 | +0.0% |
| warlock-affliction | 801.6 | 894.5 | +11.6% | 795.1 | 887.7 | +11.7% | 481.4 | 477.9 | -0.7% |
| warlock-demonology | 678.0 | 793.0 | +17.0% | 670.1 | 785.5 | +17.2% | 239.2 | 425.1 | +77.7% |
| warlock-destruction | 784.8 | 795.5 | +1.4% | 771.6 | 787.7 | +2.1% | 419.4 | 419.4 | +0.0% |
| warrior-arms | 624.8 | 695.2 | +11.3% | 614.7 | 698.2 | +13.6% | 259.3 | 259.3 | +0.0% |

Notes: Demonology's bare "before" (239) is a defect of the old rotation, not a gain to bank: its weights sweep measured spell power at -0.04 DPS a point, so the ranker picked gear on garbage weights; with Soul Fire the sweep is sane. Affliction's bare entry reads -0.7% in the ranker against +0.3% in the fixed-gear rotation search; the ranker re-picks gear on a noisy sweep, and the difference is inside that noise.

## Held, and follow-ups

- Held rotation lines are in the table above. Consecration needs Forever's own damage measured in a log. Shadow Word: Death's backlash should be modeled as the client states it (10% of maximum health) before the gate is reconsidered.
- Engine, for the engine lane: Venom and Vile Poisons add in the client and multiply in the engine; Wound Poison gets the Venom multiplier though it deals no damage; whether specials roll Windfury Totem is the load-bearing assumption under two adopted lines.
- The Assassination leveling build keeps the Venom point; `builds.test.ts`' signature table no longer requires it.
- The raid builds of the six unchanged specs are their guide builds; the guides say so.
- Nightly: `addon-data.json` and the two `Data.lua` copies drift against the new rotation notes (Rake's line text is gone and four lines are new), which fails `test_the_cli_check_passes_on_the_committed_file` and `test_clean_build_has_no_drift_findings` until `bis.yml` regenerates them; they were not regenerated here by instruction. `test_simbuffs_names_every_id_the_engine_lets_a_request_send` fails on `main` already (`windfury_totem`).
- Fork: `go test` for warrior, warlock, hunter, druid, paladin and priest passes; `sim/rogue` has a failing test on `main` as well (`TestImprovedEviscerateFollowsTheClientRanks`, item 16707 missing from the item database).

## Checks run

`go test ./...` under `sim/` passes (ladder goldens regenerated for the seven rotations that changed; `make apl-check` against the fork worktree is clean). `cd data && uv run pytest tests -q`: 1413 passed, 3 failed (the three above). Web: `npx vitest run src/lib/bis src/lib/guides src/content src/pages/guides src/components/guides` passes, `npm run check` has 0 errors, `npm run format:check` is clean.
