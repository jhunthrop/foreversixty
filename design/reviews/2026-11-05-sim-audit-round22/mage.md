# Sim-accuracy round 22: mage

Specs: mage-arcane, mage-fire, mage-frost. Site wt/r22-mage (from int/r22 591df0a19). Fork lane/r22-mage (from int/r22). Client 1.60.1.70790.

Method:
- Baseline on this branch, not round 21's before.tsv. int/r22 already carries the druid lane's crit-aura fix (Leader of the Pack and Moonkin Aura no longer stack, 3% crit in the raid preset), so every raid-60 number here starts lower. Lane baseline `$R/mage/base.tsv`, from frozen BiS (`-bis-dir $R/bis-frozen`) and the guide raid builds (`$R/rt`), at 10,000 iterations.
- Leads, in this order: the MythicSim Fire lead, the platform P5 wording, round 21's open items, the guide round 22 sims requested, and Blizzard's Forever notes line by line. Then the round 24 guide lead (Improved Scorch at 27 to 29 on short fights), and a re-check of near ties at 3% crit.
- Every sim went through `$R/sim.sh`. I used 3,000 iterations to explore, and 10,000 at seeds 1 and 50001 for every reported figure.

**Confirmed and fixed: 1** (`old`): the Fire raid build and rotation (F1). One guide-text fix (F2, guides-handoff) is not a sim bug. Arcane and Frost: aligned.

## Fire

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| F1 | **The Fire raid build was a Scorch-filler build. A Fireball build with Accelerant 5/5, Missile Barrage and Arcane Blast does about 2% more.** (MythicSim lead, verified on our engine.)<ul><li>**What was wrong.** The raid build spent 9 points on talents that do nothing on a boss: Wand Specialization 2, Arcane Resilience 3, Flame Throwing 2, plus Arcane Meditation 2 and a single Accelerant point. It used Scorch as its filler. Round 16 measured a Fireball filler only on that Scorch build's talents, where Fireball is 3.5 s and loses. Round 21 F3 judged that "no point in reach does damage on a boss" and missed Accelerant 5 and Missile Barrage.</li><li>**New raid build.** `0501052001000/55500201230133031/003`: Arcane 14 (Improved Channeling 5, Arcane Subtlety 1, Arcane Concentration 5, Missile Barrage 2, Arcane Blast 1), Fire 34 (Accelerant 5, Burning Soul 2, and the damage talents to Combustion), Frost 3 (Precision).</li><li>**Adoption rule.** The spec tree holds. Nothing the old rotation casts is dropped. Burning Soul, the unmodeled threat talent, stays. The points that go do nothing in the rotation: Wand Specialization, Arcane Resilience and Flame Throwing (the Shadow precedent). Arcane Meditation is modeled, and the sim prices it.</li><li>**Other variants tried, horde seed 1 at 3,000 iterations, with the competitor's rotation:**<ul><li>Their build: 715.0. It drops Burning Soul.</li><li>Keeping Wand Specialization instead of Subtlety: 715.0, the same. Subtlety does nothing in the engine.</li><li>Accelerant 1 with Flame Throwing and Burning Soul kept: 688.6.</li><li>Accelerant 3: 701.6.</li><li>No Arcane Blast (stillsixty's build): 708.9.</li><li>Burning Soul 2 with Arcane Blast (the adopted build): 712.1.</li></ul>Accelerant 5 drives the gain. Keeping Burning Soul costs about 0.4%.</li><li>**Rotation, tuned in our engine.** It scores 713.4 against 712.1 for the competitor's rotation on the adopted build.<ul><li>Fireball is the filler.</li><li>Scorch fills the gap when Fire Blast returns within a Scorch cast. Without this line the rotation scores 670.9.</li><li>Arcane Missiles on every Missile Barrage proc.</li><li>Arcane Blast once whenever its buff has dropped. Recasting it at under 1.5 s left was 1 DPS worse, and a pre-Pyroblast Blast line was 1 to 2 DPS worse.</li><li>The Improved Scorch refresh moves from 1.5 s left to 5 s left (701 to 711): a 3 s Fireball skipped past the old window.</li></ul></li><li>**Gates.** Every new line, the Fireball filler included, is gated on `spellIsKnown(Arcane Blast)`, a talent only the raid build takes. The Scorch filler is gated on Heating Up and not Arcane Blast. The leveling build and its cards are unchanged: bare 30, 40, 50 and 60 give 76.0, 159.4, 198.4 and 342.0, before and after. The card shows the four new rows marked "Raid build only".</li><li>**Forever sources agree on a Fireball filler with Accelerant 5 and Missile Barrage:**<ul><li>stillsixty raid build (Accelerant 5/5, Burning Soul 2/3, Missile Barrage, "Fireball as your filler"): https://www.stillsixty.com/classes/mage/fire</li><li>wowf.io (Accelerant 5, Missile Barrage, Arcane Blast, "Fill the gaps with Fireball"): https://wowf.io/en/leveling/mage/fire</li><li>neverff 20/31/0: https://www.neverff.com/wow-forever/specs/mage/fire</li><li>wowforeverbuilds ("Cast Fireball as your filler", Burning Soul 2/3): https://wowforeverbuilds.com/guide/fire-mage-pve-guide</li></ul>All read 5 Nov 2026.</li><li>**Races hold** at 3,000 iterations: Human 725.6 against Gnome 725.1; Undead 713.7, Orc 713.4, Troll 707.5.</li></ul> | runs `$R/mage/fire-lead-10k.txt`, `$R/mage/after.tsv`; tests `TestFireScorchFillerWaitsForHeatingUp` (rewritten) and `TestFireRaidBuildFillsWithFireballAndKeepsArcaneBlastUp` (sim/request), and vitest `Fire raids with a Fireball build` (builds.test.ts): all red, then green; `make apl-check` and `data/tests/test_apl.py` pass | CONFIRMED-BUG (build and rotation defect, under the adoption rule) | old (round 16 F3 and round 21 F3 judged the build on Scorch-build talents) | S 4a280289a, d3a8eb25b; E f52ae0d41, df85569ee |
| F2 | **Improved Scorch at 27 to 29 on short fights** (guide-accuracy round 24 lead).<ul><li>Bare band 20, the leveling order truncated to the level, Alliance, 10,000 iterations at seeds 1 and 50001, stack line against no stack line.</li><li>**At 27 (one point, a stack on one Scorch in three), stacking loses on short fights:** 20 s 67.9 against 73.0 (-7%), 30 s 65.2 against 69.4 (-6%), 45 s 65.7 against 66.2 (-0.7%). It wins from 60 s: 64.6 against 60.5, 90 s 58.4 against 57.5, 180 s 48.9 against 45.4.</li><li>**At 28 and 29 it wins at every length:** 20 s +5% and +8%, 60 s +11% and +13%.</li><li>The card line is right for the 180 s fight it runs. The leveling paragraph said "when a fight is long enough to pay for them", with no number. It now gives the 27 and 28 rule, and the card note's later sentence says the same.</li></ul> | `$R/mage/fire-iscorch-short-10k.txt`; fire.md leveling paragraph | Guide text (not a sim bug) | guides-handoff (round 24) | S d3a8eb25b |
| F3 | Platform P5's wording SHOULD ("From Heating Up at 30, Scorch is the filler" read as leveling advice on a 60 card). This is moot now: that row is the level 60 build's only, and its note reads "On the leveling build Scorch is the filler once you have Heating Up at 30, not Fireball." The Fireball below-30 row keeps `not auraIsKnown(400625)`. | card capture `$R/mage/shots/fire-card-0.png` | Done | - | S 4a280289a |
| F4 | MAGE-R22-1 (Fire Blast "most expensive" claim) was already fixed on int/r22. Main's newer Fire Blast note (main 2026-11-05) is not merged here, and that is left to the integrator. | `git diff HEAD origin/main` | Aligned | - | - |

## Arcane

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | Guide round 22 "Sims requested": (1) Missiles first against Frostbolt first at 10 to 22. Round 21 A2 ran this at every level, 10,000 iterations at both seeds: Missiles first holds everywhere except level 14. (2) Presence of Mind on Pyroblast against Frostbolt at 54 to 56. Round 21 A5 ran this (+0.6 to +0.75%) and adopted it. Neither moved this round: no Arcane APL or engine change, and the bare runs have no crit auras. | round 21 report | Aligned (already run) | - | - |
| A2 | Round 21's changes re-read for regressions. Both PoM lines are gated on `spellIsKnown(18809)`. Evocation registers from 20. Counterspell registers from 24 (platform P2), and no mage APL names 2139. | curated APLs; `evocation_level_test.go`, `counterspell_level_test.go` | Aligned | - | - |
| A3 | Near ties at 3% crit, races at 3,000 iterations: Gnome 712.5 against Human 708.4; Undead 701.8, Orc 699.0, Troll 693.7. The picks (Gnome, Undead) hold. | runs | Aligned | - | - |
| A4 | Arcane Concentration and Missiles: the Clearcasting proc skips the channel spell, but the missiles (`SpellCode_MageArcaneMissilesTick`, SpellFlagMage) proc it. That matches the 3 Nov note "Fixed an issue where Arcane Missiles did not proc Clearcasting". Two more 3 Nov rules are already modeled from round 20 S3: Missile Barrage is spent at channel start, and Clearcasting is kept on a Barrage Missiles. | `talents.go:186-240`, `arcane_missiles.go:142-155` | Aligned | - | - |

## Frost

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | Ice Lance's coefficient at 60 still has no live log. No source lists a Forever report (see Logs). | | OPEN (log needed) | - | - |
| R2 | Ice Lance on a Frozen target: the 16 Oct note says +200%, and the 70790 talent text says "300% increased damage". The client text is later and wins, and the engine is x4 (`TestIceLanceFrozenMultiplierIsFour`). | talents/mage.json; Blizzard thread 2360696 (16 Oct) | Aligned (client beats the note) | - | - |
| R3 | Races at 3% crit: Gnome 620.4 against Human 618.9; Undead 613.5, Orc 611.1, Troll 611.2. The picks hold. | runs | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **Blizzard's Forever notes, every mage line** (thread 2360696, all three chunks read 5 Nov). Each line checked against the engine:<ul><li>Ignite no longer double dips (`ignite.go:59-88`, two tests).</li><li>Combustion has 3 charges.</li><li>Improved Scorch's debuff makes one roll, on the Scorch's own landed result (`scorch.go:68-72`). Winter's Chill takes no resist roll either.</li><li>Heating Up is kept under Presence of Mind (round 20).</li><li>Fire Power is 1/2/3%. Incineration has 5 ranks. Accelerant includes Arcane Wrath. Critical Mass has 3 ranks.</li><li>Arcane Blast's cost is +100% and its spell damage is 12% of Intellect.</li><li>Ice Barrier adds 15% Frost damage. Ice Lance shares an 8 s cooldown with Fire Blast. Winter's Chill adds 20% to Cone of Cold.</li><li>The low-rank talent penalty is gone except on Blizzard (round 20). Precision gives 2/4/6%.</li></ul>Every line is aligned. The spellconst, talents, trainables and hotfix-push files for mage are unchanged since round 21, and hotfix-push still lists no rows. | engine files above; `git log` on data/builds/1.60.1.70790 | Aligned | - | - |
| X2 | The guide round 22 Should "fire.md:84 takes twice as long" is gone: the Frostfire Bolt line now compares it with Fireball ("hits for much less"). | fire.md | Done | - | - |

## Logs

Last looked 07:59 UTC 5 Nov.
- api.foreversixty.gg `/v1/reports/recent` still lists only `z2opt2x2gs5u` (20 Oct).
- foreverlogs.gg `/api/reports/public` answers 403.
- Warcraft Logs v1 `/v1/zones` lists no Forever zone on www (66 zones), classic (57) or fresh (21).

None of the mage log items can be settled.

## Before/after

Raid 60, frozen BiS, guide raid builds, 10,000 iterations. Before is this lane's baseline on int/r22 (already 3% crit). After is wt/r22-mage plus lane/r22-mage.

| Spec | Faction | Seed 1 before to after | % (s1 / s50001) |
|---|---|---|---|
| mage-fire | alliance | 712.5 to 725.0 | +1.75 / +1.74 |
| mage-fire | horde | 698.2 to 712.4 | +2.03 / +2.06 |
| mage-fire bare 30, 40, 50, 60 | horde | 76.0, 159.4, 198.4, 342.0, unchanged | 0 |
| mage-arcane | both | unchanged (711.5 / 698.8) | 0 |
| mage-frost | both | unchanged (619.2 / 610.9) | 0 |

No spec outside mage moves. The changes are the mage-fire APL, its two copies, the Fire guide and the mage-fire line of raid-keeps.json.

## Files to regenerate

- addon-builds.json (mage-fire raid order) and addon-data.json plus Data.lua: the four new mage-fire rows and their notes, the changed notes, and the Scorch refresh at 5 s. They were regenerated locally only to preview the card, then restored.
- sim/request ladder golden mage-fire, only if a rung's casts read the changed lines. The ladder walks the leveling build, so the Arcane Blast-gated rows never fire. Expect at most the Unresolved column for 400589 or 1239700, both declared inert or expected idle.
- coach starter data/curated/coach/mage-fire.md, if it is regenerated from the APL.

## Measure at refresh

- mage-fire raid BiS and weights on the new raid build. Fireball is now the filler, and Arcane Blast and Missiles now carry about 13% of the damage, so plain spell power gains on Fire-only power.
- Carried from round 21: the Arcane re-rank on its raid string, and the Frost re-rank after Ice Lance's 0.107.

## For other lanes

- **Platform (component), wow-player SHOULD:** the card legend "Raid build only: needs a talent the level 60 build does not take. On the level 60 build, skip that row." reads backwards to a raider. Name the non-raid build "Leveling build", or say "levels 1 to 59", in the legend, header and chips.
- **Every lane, method:** spec-breakdown `-level N` on the band-20 entry does not truncate the talents to the level. Pass the leveling order truncated by `-talents`, or talent-gated lines read as never firing. The first F2 runs showed identical numbers with and without the stack line for this reason.

## For the guide lane

- Done here: fire.md description, overview, strengths, talents (raid build), loop, opener, Combustion, mana, Frostfire Bolt, stat priority, enchants and the leveling Improved Scorch rule. This builds on main's opener wording ("unless 3 Heating Up stacks have a Pyroblast waiting").
- wow-player SHOULDs not taken:
  - "Improved Scorch now only boosts your own Fire damage" (fire.md:48): "now" is patch-history talk.
  - The hit cap is not given in gear terms. Forever hit is a percentage, so name the gear percentage to aim for.
  - "Fire drinks a Major Mana Potion on every boss" appears three times.
- main has a newer Fire Blast note than int/r22. Integrate it on top of this branch's notes: the Fire Blast row is untouched here.

## Guide review

wow-player read the rendered /guides/mage/fire at 1440 (astro dev on wt/r22-mage, with addon data regenerated locally and then restored). Captures: `$R/mage/shots/fire-card-0.png`, `fire-talents.png`, `fire-talents2.png`, `fire-1440.png`.
- **First pass: NO-SHIP on three MUSTs.**
  - Legend naming: my prose said "leveling build" where the page says "level 60 build". The prose now matches the page, and the legend wording goes to platform.
  - Skeleton trees: capture timing. They hydrate, as `fire-talents.png` shows.
  - Gem and rune shared cooldown: the Forever client's SpellCategory 1153, with an engine test, so the text stands.
  - SHOULDs taken: the Missile Barrage wording, the Strengths line, Incineration on Arcane Blast, and a shorter gap-Scorch note.
- **Second pass: SHIP**, no MUSTs. The leveling Improved Scorch paragraph was confirmed in the rendered HTML.

## Still open

- **Fire leveling build at 60.** The raid build does +2.3% at bare 60 (350.0 against 342.0, 3,000 iterations). Adopting it for the level-60 build means re-walking the leveling order, and Accelerant early is what the Forever leveling guides do. Lead for round 23.
- Ice Lance's coefficient at 60 needs a level 56+ Frost log.
- Arcane Power's cost needs a mana log.
- Carried: the level-60 book ranks (pending owner).
