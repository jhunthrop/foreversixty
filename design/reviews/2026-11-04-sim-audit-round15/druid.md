# Sim-accuracy round 15: Druid

Lane `druid`. Site wt/r15-druid, fork lane/r15-druid, both from int/r15 (site fd320b60e, fork d1d9f2775). Client 1.60.1.70755. Specs: Balance, Feral cat, Feral bear (threat, boss AP 268, keeps Demoralizing Roar), Restoration (heal scale 2.3; druid rows of rank-choices.json and its APL).

Summary: 3 confirmed bugs, all fixed: one engine value the server changed (Rend and Tear) and two dead talent points in guide builds (cat leveling, Restoration). One owner question (bear Furor). The guide-accuracy round-1 sim leads are settled below.

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Rend and Tear was 2% a rank; Forever's server makes it 3%.**<br>**What was wrong:** cat and bear melee abilities against a bleeding target dealt 10% more at 5/5.<br>**What the client says:** the client's own curve for the talent (TraitDefinitionEffectPoints 23779, curve 85357) still reads 2/4/6/8/10 on 70755, and the hotfix push does not touch it. The server's value is in Wowhead's live Forever talent payload (fetched 2026-11-03): "by 3% / 6% / 9% / 12% / 15%". The site's talents/druid.json carries the same text (its 70753 copy read 2%..10%). As the talent-hotfix memory note says, talent hotfixes live in that payload, not in wago tables. Note: the talent has 5 ranks, so 15% is 5/5, not 3/3.<br>**What it touches (checked):** the direct damage of Claw, Shred, the first hit of Rake, Ravage, Ferocious Bite, Primal Bite, Maul, Swipe and Lacerate (`feral_talents.go:99`, callers in claw/shred/rake/ravage/ferocious_bite/primal_bite/maul/swipe/lacerate.go). It never touches a bleed's own ticks or white swings. "Bleeding" means the druid's own Lacerate, Rip, Rake or Pounce, or the raid's assumed bleed.<br>**What changed:** `rendAndTearDamagePerRank` 0.02 to 0.03. | Tests changed first and failed: `TestRendAndTearNeedsABleed` (1.1, want 1.15), `TestRendAndTearRaisesShredOnABleedingTarget`, `TestRendAndTearCountsTheCatsOwnBleeds`. All three pass now. Fork golden `sim/druid/tank/TestFeralTank.results` regenerated. talents/druid.json lines 784-800; nether.wowhead.com/forever/data/talents-classic?dv=100 | CONFIRMED-BUG | new-data (70755 server value; client curve unchanged), guides-handoff (S1) | E e45a540db |
| S2 | **The cat leveling build spent 8 points on caster talents.**<br>**What was wrong:** `01/50212032123012123001/535033` took Nature's Focus 5 (pushback on Arcane and Nature spells) and Subtlety 3 (threat from Nature and Arcane spells). The cat's rotation casts neither school after 20, so neither does anything for it. Meanwhile Heart of the Wild sat at 3/5.<br>**What changed:** the build is now `05002/50212032123012123001/05503`, guide-accuracy DRUID-3's proposal. It keeps the questing Feral 31 (Brutal Impact, Feral Charge) and takes the raid build's Balance 7 (Genesis 5, Nature's Majesty 2) and Restoration 13 (Heart of the Wild 5, Naturalist 5, Natural Shapeshifter 3). It is legal: 51 points, the tier gates are met by Genesis 5 and Heart of the Wild 5, and json order equals the (tier, column) order the ladder reads.<br>**Bands:** the ladder walks Feral, then Balance, then Restoration, so bands 20-40 are unchanged. Band 50 swaps Genesis 1, Nature's Focus 5, Heart of the Wild 3 and Naturalist 1 for Genesis 5, Nature's Majesty 2 and Heart of the Wild 3. Bare 60 is the whole build. | `TestGuideBuildsSpendNothingOnDeadTalents/feral` and `TestCatLevelSixtyBuildFinishesHeartOfTheWild` (sim/leveling) failed on the old string and pass. Bare 50: 161.1 to 168.7 (alliance) and 157.0 to 164.5 (horde) at seed 1, +4.7% at both seeds. Bare 60: 275.6 to 297.1 and 272.0 to 293.7, +7.8 / +8.0% at both seeds (`lvl.tsv`) | CONFIRMED-BUG (owner rule: a provably dead point is a bug by itself) | old (round 12, 454ae42b9: the talent search held both talents as unmodeled keeps), guides-handoff (DRUID-3 / SIM-D1) | S 6b9b19145 |
| S3 | Guide-accuracy round-1 sim rows, each re-checked:<br>- **S1 Rend and Tear:** item S1 above.<br>- **S2 Demoralizing Roar in the bear APL:** present (9898, priority 6, refreshed in its last 3 s).<br>- **S3 Frenzied Regeneration:** the engine reads 1% of max health a rage point, up to 10 rage a second for 10 s (`frenzied_regeneration.go:23`, round 14 B1). Wowhead tooltip 22842 says the same: "Converts up to 10 Rage per second into health for 10 sec. Each point of Rage is converted into 1% health." So at full rage it can heal your whole health bar. The guide's original wording was right, and the guide lane softened it (see "For the guide lane").<br>- **S4 expertise in the bear weights:** no Forever item for any class carries expertise (items/*.json). The weight comes from the shared weights tool (`sim/cmd/leveling-bis/expertisecap.go`); it never moves a pick. Sent to "For other lanes" and "Measure at refresh".<br>- **S5 Eclipse banked on a resisted Wrath:** round 14 casters B1 settled this by the identical proc-mask analogue (408248 = Missile Barrage's 0x10000). No new data. Aligned (analogue).<br>- **S6 feral ladder goldens:** listed under Files to regenerate.<br>- **DRUID-19 Rejuvenation rank 1 tooltip:** pending owner, not reopened. | see each line | S1 confirmed; S2, S3, S5 aligned; S4 other lane; DRUID-19 pending | - | - |

## Feral cat (re-audit)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | **Rake upkeep re-measured** after Rend and Tear's 15% and the bleed-crit ruling (last measured in round 3 F5, -4.1%). Raid 60, 3000 iterations, seed 1: base 675.1 / 670.7; Rake whenever it is down, ahead of Shred: 648.5 / 644.6 (-3.9%); Rake only before the first Rip (to open Rend and Tear early): 673.4 / 668.3 (-0.3%). The list stays without Rake. | `rot/rake*.json` (lane dir) | Aligned | - | - |
| C2 | APL read cast by cast: Prowl pre-pull, Ravage 9867 r4, Faerie Fire 9907 in form (free, 6 s cd; round 14 C3), the sub-20 Moonfire and Wrath lines gated on Claw, Shifting Power at 60 energy or less, Rip at 5 CP under 6 s, Bite at 5 CP with Rip at 6 s or more, then Shred, Claw before Bite is learned. Ranks are the trainable maxima (Rip 9896 r6 at 60, Rake 9904 r4 at 54). Costs match trainables (Shred 60, less Shredding Attacks' 18 = 42; Claw 45; Rake 40; Rip 30). | trainables/druid.json; `claw.go:47`, `lacerate.go:42` | Aligned | - | - |
| C3 | Raid build, talent by talent: every point is modeled or a path point. Ferocity, Feral Swiftness and Feral Instinct get you down the tree; Shredding Attacks, Savage Fury, Sharpened Claws, Predatory Strikes, Predatory Discipline, Blood Frenzy 1, Leader of the Pack, Natural Instinct, Rend and Tear 5 and Shifting Power are all modeled; Genesis feeds Rip; Nature's Majesty 4% crit; Heart of the Wild, Naturalist and Natural Shapeshifter (Shifting Power's cost). No dead point. | talents/druid.json; round 14 hybrids C1 | Aligned | - | - |
| C4 | Raid BiS gear (night elf): Shadowcraft 5 pieces give Rogue Armor Energize (27787 to 27788, 20 energy) through the shared client-set model at the unstated proc rate. Molten Heart of the Mountain is Molten Fury (1249113, 55 AP for 15 s). Frozen Heart of the Mountain is worn for its 9 hit; its Frost/Shadow use does nothing for a cat. Idol of the Dream adds 2 s to Rip. Hammer of the Titans' stun does no damage. All are modeled as the item text says. The proc rate of the unstated set chance is shared core work that earlier rounds own. | items/druid.json effect text; `sim/common/item_effects.go:3674-3730`; `client_sets_gen.go:26,406-407` | Aligned | - | - |

## Feral bear (re-audit)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| B1 | Rend and Tear (S1) is 5/5 in the bear build: threat rises 3.8% at raid 60. | after.tsv | (S1) | - | E e45a540db |
| B2 | APL and costs: Enrage pre-pull, Barkskin on cooldown (castable in form, round 11 G3), Frenzied Regeneration under 50% health with 20+ rage, Primal Bite r3, Lacerate r5 (15 rage less Shredding Attacks' 3), Faerie Fire, Demoralizing Roar r5, Maul r7 at 15+ rage (less Ferocity's 5), Swipe r5 at 40+. Every value matches round 14 B3 and the client rows; 70755 changed none of them. | spellconst druid.json; `maul.go:50`, `lacerate.go:42-45`, `bear_talents_consts.go` | Aligned | - | - |
| B3 | **Furor 5 never fires in the tank fight** (the bear starts in Dire Bear Form and never shifts; round 6 G0). In game it is 10 rage each time a bear shifts back in (Innervate, Rebirth, a decurse), which the engine does not model, so the adoption rule keeps it. Probe: Furor 0, Naturalist 1 to 5 and Feral Charge 1 raise threat 993.1 to 1070.0 (+7.7%, raid 60, horde, 3000 iterations). The one Forever source (Warcraft Tavern's Forever feral guide, undated, read 2026-11-03) says Furor is "usually not needed in endgame group content". Its bear build is 9/42/0, which predates Heart of the Wild's move to Restoration. Mobalytics and Icy Veins Forever pages refused the fetch (403). | warcrafttavern.com/forever/guides/feral-druid/ | OPEN (owner: is a bear's Furor an unmodeled keep, or a point that does nothing in a raid?) | - | - |
| B4 | Leveling ladder and talents: Lacerate at 30 and Primal Bite at 40 (ladder-first row); Thick Hide, Natural Reaction, Feral Swiftness dodge, Blood Frenzy rage on crits, Heart of the Wild Stamina, Naturalist 1. All modeled. | ladder-first.json; `feral_talents.go` | Aligned | - | - |

## Balance (re-audit)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| L1 | APL cast by cast: Wrath r8 pre-pull, Moonfire as the last-seconds filler, Starfire r6 9876 (rank 7 is a book, round 14 B5), Moonfire r10 and Insect Swarm r5 upkeep, Wrath to bank Eclipse, Wrath when mana-bound (downranking allowed), Starfire. 70755 changed no druid spell row (round 14 build.md X7). | APL; round 14 casters B2 | Aligned | - | - |
| L2 | Builds, talent by talent: every Balance point is modeled or a path point. Improved Moonfire 2 is cast and is Vengeance's prerequisite. Nature's Reach gives 4% hit. Improved Starfire 5 is 0.1 s a rank. Restoration's Nature's Focus 5 and Subtlety 3 do nothing in the sim but are real for a caster (pushback, threat): they are on the raid keep list (raid-keeps.json), so the adoption rule keeps them. Round 14 B3's +1.6% swap stays with the owner. | raid-keeps.json; round 14 casters B3 | Aligned (keeps) | - | - |
| L3 | Rank rows: Moonfire r3, Wrath r4 and Insect Swarm r1 at 30-39 are all trainable by 30 (learned at 16, 22 and 20; Insect Swarm's 15-point row opens at 24). | trainables/druid.json | Aligned | - | - |

## Restoration (re-audit)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | **Improved Moonfire 1 did nothing for a Restoration druid.**<br>**What was wrong:** the build `05312001/0/55530051531132` spent a point on "damage and critical strike chance of your Moonfire spell". The healer never casts Moonfire. Improved Moonfire is the prerequisite only of Vengeance, which the build does not take. Nature's Splendor's 10-point gate is already met by Genesis 5, Moonglow 3 and Nature's Majesty 2. Moonglow is itself damage-only in Forever (text "damaging spells"; engine mask `DruidSpellMaskBalanceDamage`), but it is the cheapest path to Nature's Splendor.<br>**What changed:** the point moves to Improved Regrowth 1 (+10% crit on Regrowth, which the tank line casts; prerequisite Improved Rejuvenation 3 is held): `05302001/0/555300515311321`. Bands 20-40 are unchanged. Band 50 swaps Genesis 2 for Genesis 1 and Improved Regrowth 1. | `TestGuideBuildsSpendNothingOnDeadTalents/restoration` failed on the old string and passes. Raid 60: 1014.5 to 1019.2 (alliance) and 1013.6 to 1018.5 (horde), +0.46 to +0.48%. Bare 60: +0.79 to +0.81%. Bare 50: +0.18 to +0.23%. Both seeds (`resto-imf.tsv`) | CONFIRMED-BUG (dead point) | old (Improved Moonfire in the build since e43958ce5, 2026-10-31; round 14 healers R2 read the spells, not the build) | S 9b931dc3d |
| R2 | **Side effect in a shared test:** `TestCuratedLadderFirstIsSound` takes the build's deepest talent as the "capstone" a reservation must not drop at a published band. Restoration's deepest talent is now Improved Regrowth (row 5). The plain walk's spine would chase it at 40 by dropping Nature's Swiftness and Swiftmend, and a reservation cannot hold it at 40. The test gains a one-entry map naming this filler as not defining. The ladder output is unchanged (checked at 20/30/40/50/60). | `sim/leveling/ladder_first_curated_test.go` | - (test policy) | - | S 9b931dc3d |
| R3 | APL: Innervate under 30% mana, Nature's Swiftness into Healing Touch r10 on a tank under 35%, Swiftmend on the tank under 75% with mana to spare, Tranquility r4 on group damage, Regrowth r9 and Rejuvenation r10 on the tank, Healing Touch r10 under 55%, Rejuvenation r1 on the party (pending owner, O2). Spell rows, Swiftmend (20% base mana, 15 s), Innervate and Bloom are as round 14 R2 records. | APL; round 14 healers R2 | Aligned | - | - |
| R4 | Rank rows: Rejuvenation r4 at 28-49 and bare 60, r1 at 50-59. Each rank is trainable at its band (r4 at 22, r1 at 4). The r1 row rests on the pending tooltip. | rank-choices.json; trainables | Aligned (r1 pending owner) | - | - |
| R5 | Build, rest of it: Nature's Focus 5 (heals are Nature, so pushback protection is real but unmodeled), Heart of the Wild, Naturalist (Healing Touch cast time), Reflection, Gift of Nature, Bloom, Tranquil Spirit, Improved Rejuvenation, Swiftmend, Nature's Swiftness, Living Spirit, Improved Tranquility (Tranquility's cooldown -60%, which the list uses) are all modeled or real. | talents/druid.json | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Racials of the published races (night elf, tauren): unchanged since round 12; Elune's Light, Quickness, Endurance and War Stomp match races.json (guide-accuracy "checked and right"). | races.json | Aligned | - | - |
| X2 | Talent-text provenance: on 70755 the client curve and the server value disagree for Rend and Tear, and the pipeline's talents/druid.json follows the server. I diffed every druid talent between the two builds' talents json: Rend and Tear is the only change. | `diff talents/druid.json` | Aligned (others) | - | - |

## Before/after

Raid 60, 10,000 iterations, both factions, seeds 1 and 50001. Before = `$R/before.tsv` (site fd320b60e, engine d1d9f2775). After = lane engine e45a540db with site wt/r15-druid. spec-breakdown sims the published BiS talents, which predate the guide builds. So the raid table shows Rend and Tear. The build fixes are measured on the guide builds in S2 and R1.

| Spec | Faction | Seed | Before | After | % |
|---|---|---|---|---|---|
| druid-balance | alliance | 1 | 569.5 | 569.5 | +0.00% |
| druid-balance | alliance | 50001 | 570.3 | 570.3 | +0.00% |
| druid-balance | horde | 1 | 574.9 | 574.9 | +0.00% |
| druid-balance | horde | 50001 | 575.1 | 575.1 | +0.00% |
| druid-feral-bear | alliance | 1 | 964.2 | 1000.6 | +3.78% |
| druid-feral-bear | alliance | 50001 | 965.2 | 1001.6 | +3.77% |
| druid-feral-bear | horde | 1 | 955.7 | 992.0 | +3.80% |
| druid-feral-bear | horde | 50001 | 956.3 | 992.6 | +3.80% |
| druid-feral | alliance | 1 | 664.9 | 675.1 | +1.53% |
| druid-feral | alliance | 50001 | 664.8 | 675.1 | +1.55% |
| druid-feral | horde | 1 | 660.4 | 670.6 | +1.54% |
| druid-feral | horde | 50001 | 660.2 | 670.3 | +1.53% |
| druid-restoration | alliance | 1 | 1006.4 | 1006.4 | +0.00% |
| druid-restoration | alliance | 50001 | 1006.7 | 1006.7 | +0.00% |
| druid-restoration | horde | 1 | 1005.3 | 1005.3 | +0.00% |
| druid-restoration | horde | 50001 | 1005.2 | 1005.2 | +0.00% |

Healers are in effective HPS, the bear in threat per second. Balance and Restoration do not move on published talents (the Restoration build change is R1).

Bare bands (horde, published band talents, Rend and Tear only):

| Spec | Band | Seed | Before | After | % |
|---|---|---|---|---|---|
| druid-feral | 20 | 1 | 45.0 | 45.0 | +0.00% |
| druid-feral | 20 | 50001 | 45.0 | 45.0 | +0.00% |
| druid-feral | 30 | 1 | 83.3 | 83.3 | +0.00% |
| druid-feral | 30 | 50001 | 83.2 | 83.2 | +0.00% |
| druid-feral | 40 | 1 | 115.5 | 116.2 | +0.61% |
| druid-feral | 40 | 50001 | 115.4 | 116.2 | +0.69% |
| druid-feral | 50 | 1 | 153.6 | 154.6 | +0.65% |
| druid-feral | 50 | 50001 | 153.6 | 154.6 | +0.65% |
| druid-feral | 60 | 1 | 266.8 | 268.8 | +0.75% |
| druid-feral | 60 | 50001 | 267.0 | 269.0 | +0.75% |
| druid-feral-bear | 20 | 1 | 57.2 | 57.2 | +0.00% |
| druid-feral-bear | 20 | 50001 | 57.1 | 57.1 | +0.00% |
| druid-feral-bear | 30 | 1 | 75.1 | 75.1 | +0.00% |
| druid-feral-bear | 30 | 50001 | 75.0 | 75.0 | +0.00% |
| druid-feral-bear | 40 | 1 | 170.1 | 170.1 | +0.00% |
| druid-feral-bear | 40 | 50001 | 170.0 | 170.0 | +0.00% |
| druid-feral-bear | 50 | 1 | 302.1 | 312.0 | +3.28% |
| druid-feral-bear | 50 | 50001 | 301.9 | 311.8 | +3.28% |
| druid-feral-bear | 60 | 1 | 416.7 | 430.9 | +3.41% |
| druid-feral-bear | 60 | 50001 | 416.0 | 430.1 | +3.39% |

Bands 20-30 take no Rend and Tear; the cat build reaches it at 40 (3 ranks), the bear at 50.

No spec outside the druid moves: the engine change is in `sim/druid`, and the site changes are two druid guide build lines and sim/leveling tests.

## Files to regenerate

- `sim/request/testdata/ladder/druid-feral.golden.md`, `druid-feral-bear.golden.md` (Rend and Tear) and `druid-restoration.golden.md` (the band 50/60 build): TestRotationLadder.
- BiS, weights and band talents (at the refresh): `bis/druid-feral.json` (band 50/60 bare talents, raid Rend and Tear), `bis/druid-feral-bear.json`, `bis/druid-restoration.json` (band 50/60 talents).
- `addon-builds.json`, `dps-per-point.json`, `addon-data.json` / Data.lua, and the coach starters that carry the cat leveling or Restoration build strings.
- Fork golden regenerated here: `sim/druid/tank/TestFeralTank.results` (E e45a540db).

## Measure at refresh

- The cat bare-50/60 and Restoration band-50/60 gains on the re-picked gear (measured here on published gear with guide-build talents).
- S4: drop expertise from the bear weights table (no item carries it).
- Restoration coverage at raid 60 after the Improved Regrowth point (+0.5% healing; round 14 A1's 83% coverage figure moves slightly).

## For other lanes

- **Weights tooling (shared):** bear band weights publish an expertise scale factor (7.06 raid, 8.32 bare) although no Forever item has expertise. Hide it from the rendered table, or drop the stat from the bear weights run.
- **sim/leveling (shared test):** `TestCuratedLadderFirstIsSound` gained `spineCapstoneNotDefining` (one entry, druid-restoration). No other spec is affected.

## For the guide lane

On land/guides-r1:
- `feral.md:58`, replace "From 41 take Genesis, then Nature's Focus and three points of Heart of the Wild to open Restoration, then Naturalist, Natural Shapeshifter and Subtlety last. Nature's Focus keeps a mob's hits from pushing back the Healing Touch or Regrowth you cast between forms; Subtlety and Nature's Focus both move at the 60 respec." with "From 41 take Genesis and Nature's Majesty in Balance, then Heart of the Wild, Naturalist and Natural Shapeshifter in Restoration: the same points as the raid build." In the next sentence, drop "and Heart of the Wild" from "it finishes Rend and Tear and Heart of the Wild" and drop "and takes Genesis and Nature's Majesty" (the leveling build already has both).
- `restoration.md:54`, replace "The build is 39 points of Restoration and 12 of Balance." with "The build is 40 points of Restoration and 11 of Balance."
- `restoration.md:65`, replace "Moonglow (3 points) and Improved Moonfire (1 point) are the price of Nature's Splendor" with "Moonglow (3 points) is the price of Nature's Splendor". Add to the Restoration talent list: "- **Improved Regrowth** (1 point): more crits on the Regrowth you keep on the tank."
- `feral-bear.md:88`: Frenzied Regeneration's tooltip (22842) is 1% of your health per rage point, up to 10 rage a second for 10 seconds. "Frenzied Regeneration turns rage into health over 10 seconds" can become "Frenzied Regeneration turns up to 10 rage a second into 1% of your health a point for 10 seconds, so a full bar of rage heals you to full; use it under half health with rage to spare, not at the pull."
- (Already on the guide lane) Rend and Tear's 3% / 15% text (DRUID-1, DRUID-2) is now what the engine deals.

## Still open

- B3: bear Furor (owner).
- Pending owner, not reopened: Rejuvenation rank 1 tooltip (DRUID-19, O2); Moonkin Rebirth (DRUID-20); Eclipse on a resisted Wrath (DRUID-16, aligned by analogue).
- Warcraft Logs is blocked (gcloud login expired); not used.

## Wave 2

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| W1 | **Frenzied Regeneration "heals the full bar" (guide-accuracy Q1, which calls for a sim change from 0.3% to 1% a rage).** The guide lane worked on engine 2eaae03c6, which still had 0.3%. Round 14 B1 (E 809b1fc9b, inside the int/r15 base d1d9f2775) already made it 1%. The 70755 client agrees on every value:<br>- SpellEffect 700422 is aura 23, period 1000 ms, base points 1.<br>- The Spell text is "Converts up to 10 Rage per second into health for $d. Each point of Rage is converted into $s1% health."<br>- SpellDuration 1 is 10,000 ms; SpellCooldowns 55743 is a 180,000 ms category with a 1.5 s global.<br>The engine (`frenzied_regeneration.go:17-24`) has 10 s, 3 min, 10 rage a second and 0.01 a rage. At full rage that is 10 ticks of 10% = the whole bar. The guide claim is right and the engine already deals it. No change. | `TestFrenziedRegenerationIsOnePercentAPointOfRage` passes on the lane engine; client rows above | Aligned (fixed in round 14) | - | - |

Bear raid 60 on the lane (10,000 iterations): threat 1000.6 / 1001.6 (alliance, seeds 1 / 50001) and 992.0 / 992.6 (horde); chance of death 10.9 / 9.9% and 10.1 / 9.8%; damage taken 465.8 / 465.3 and 473.5 / 472.9. Round 14 recorded 10.1% for horde at seed 1, so Rend and Tear (S1) moves threat, not survival. The bear's "under half health with 20 rage" line was already measured on the 1% conversion in round 14 (B1, O4).

For the guide lane: the original "can heal you for your whole health bar given the rage" stands; see the feral-bear.md:88 wording above.

## Wave 3

The shared low-rank penalty arrived with int/r15 (fork 9c511a2cd, site 1e96d8df2; priest lane E 8d9ca6aae, 974614b8b). Under it a rank keeps its full bonus until the caster is 17 levels past the rank's learn level, then loses 5% a level. I merged both lane branches onto those heads; both merges were fast-forwards, and my wave-1/2 commits are inside them. The sweep ran one spec-breakdown binary per variant, built from an edited copy of the embedded rank-choices.json and restored afterwards. 10,000 iterations, both factions, seeds 1 and 50001. Lane dir: `w3/w3.tsv`.

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| W3-1 | **Balance's three 30-39 rows all still win** at bare 30 (each rank is under 17 levels past its learn level there: Moonfire 3 learned at 16, Wrath 4 at 22, Insect Swarm 1 at 20). Dropping all three costs 16.7 to 17.6%; Moonfire alone 5.6 to 5.8%, Wrath alone 6.5 to 6.7%, Insect Swarm alone 2.4 to 2.7%, at both seeds and factions. Kept. The APL's fixed ranks are all the maxima (Wrath 8, Starfire 6, Moonfire 10, Insect Swarm 5), so raid 60 does not move. | table below | Aligned (rows kept) | - | - |
| W3-2 | **Restoration's rows (the priest lane's re-pick) do not lose.** Party Rejuvenation rank 1 (row 8910, 10-59 and bare 60) beats the authored rank 7 by 13 to 18% at bare 20-50 and 9.2 to 9.3% at bare 60. Tank Rejuvenation rank 4 (row 9841, 28-49) beats max rank by 4.4 to 4.8% at bare 30 and 40. Kept. | table below | Aligned (rows kept) | - | - |
| W3-3 | **Cat and bear: no rank-sensitive cast.** Their abilities carry no spell-power or healing coefficient. Every APL rank is the maximum (Shred 5, Claw 5, Rip 6, Bite 5, Ravage 4, Faerie Fire 4, Prowl 3; Maul 7, Swipe 5, Lacerate 5, Primal Bite 3, Demoralizing Roar 5). The cat's Moonfire and Wrath lines run only below 20, where no band is published. | APLs | Aligned | - | - |

What the engine change cost (round start to now, rows kept) and what the rows hold back. "Row dropped" is the loss against keeping the row:

| Spec | Preset/band | Faction | Seed | Before (round start) | Now (rows kept) | Engine cost | Row dropped |
|---|---|---|---|---|---|---|---|
| druid-balance | bare 30 | alliance | 1 | 73.9 | 75.0 | +1.49% | none -16.7%; noMF -5.6%; noW -6.7%; noIS -2.7% |
| druid-balance | bare 30 | alliance | 50001 | 73.9 | 74.9 | +1.35% | none -16.7%; noMF -5.7%; noW -6.5%; noIS -2.7% |
| druid-balance | bare 30 | horde | 1 | 73.3 | 74.3 | +1.36% | none -17.5%; noMF -5.7%; noW -6.6%; noIS -2.4% |
| druid-balance | bare 30 | horde | 50001 | 73.2 | 74.3 | +1.50% | none -17.6%; noMF -5.8%; noW -6.6%; noIS -2.6% |
| druid-balance | raid 60 | alliance | 1 | 569.5 | 569.5 | +0.00% | - |
| druid-balance | raid 60 | alliance | 50001 | 570.3 | 570.3 | +0.00% | - |
| druid-balance | raid 60 | horde | 1 | 574.9 | 574.9 | +0.00% | - |
| druid-balance | raid 60 | horde | 50001 | 575.1 | 575.1 | +0.00% | - |
| druid-restoration | bare 20 | alliance | 1 | 44.1 | 44.1 | +0.00% | no8910 -13.6% |
| druid-restoration | bare 20 | alliance | 50001 | 44.1 | 44.1 | +0.00% | no8910 -13.6% |
| druid-restoration | bare 20 | horde | 1 | 40.6 | 40.6 | +0.00% | no8910 -13.1% |
| druid-restoration | bare 20 | horde | 50001 | 40.6 | 40.6 | +0.00% | no8910 -13.3% |
| druid-restoration | bare 30 | alliance | 1 | 118.1 | 113.2 | -4.15% | no8910 -17.3%; no9841 -4.5% |
| druid-restoration | bare 30 | alliance | 50001 | 118.2 | 113.3 | -4.15% | no8910 -17.2%; no9841 -4.6% |
| druid-restoration | bare 30 | horde | 1 | 116.0 | 111.4 | -3.97% | no8910 -17.5%; no9841 -4.5% |
| druid-restoration | bare 30 | horde | 50001 | 116.1 | 111.5 | -3.96% | no8910 -17.5%; no9841 -4.6% |
| druid-restoration | bare 40 | alliance | 1 | 228.0 | 213.7 | -6.27% | no8910 -14.0%; no9841 -4.5% |
| druid-restoration | bare 40 | alliance | 50001 | 228.1 | 213.8 | -6.27% | no8910 -14.1%; no9841 -4.4% |
| druid-restoration | bare 40 | horde | 1 | 228.0 | 214.4 | -5.96% | no8910 -13.6%; no9841 -4.8% |
| druid-restoration | bare 40 | horde | 50001 | 228.0 | 214.3 | -6.01% | no8910 -13.7%; no9841 -4.6% |
| druid-restoration | bare 50 | alliance | 1 | 392.5 | 325.9 | -16.97% | no8910 -15.8% |
| druid-restoration | bare 50 | alliance | 50001 | 392.7 | 326.0 | -16.98% | no8910 -15.9% |
| druid-restoration | bare 50 | horde | 1 | 389.3 | 323.4 | -16.93% | no8910 -16.0% |
| druid-restoration | bare 50 | horde | 50001 | 389.6 | 323.5 | -16.97% | no8910 -16.0% |
| druid-restoration | bare 60 | alliance | 1 | 597.9 | 463.0 | -22.56% | no8910 -9.2% |
| druid-restoration | bare 60 | alliance | 50001 | 598.3 | 463.0 | -22.61% | no8910 -9.3% |
| druid-restoration | bare 60 | horde | 1 | 594.2 | 459.1 | -22.74% | no8910 -9.3% |
| druid-restoration | bare 60 | horde | 50001 | 594.4 | 458.9 | -22.80% | no8910 -9.2% |
| druid-restoration | raid 60 | alliance | 1 | 1006.4 | 925.8 | -8.01% | - |
| druid-restoration | raid 60 | alliance | 50001 | 1006.7 | 925.8 | -8.04% | - |
| druid-restoration | raid 60 | horde | 1 | 1005.3 | 923.3 | -8.16% | - |
| druid-restoration | raid 60 | horde | 50001 | 1005.2 | 923.8 | -8.10% | - |

Wave 1 left Balance and Restoration at these bands as they were at the round start (the Restoration build change reaches bands 50 and 60, and only on the guide build). So the "Before" column is also the wave-1 figure. Restoration's loss (bare 30 to 60: -4 to -23%; raid 60: -8.0 to -8.2%) is the shared penalty itself, not a druid row. Balance bare 30 rises 1.3 to 1.5% on the merged tree.

No row changed, so there is nothing to sync (sim/rankchoices copy, sim/request APLs). Files to regenerate: the healer ladder goldens are the priest lane's list; druid-balance needs nothing new.

## Wave 4

I merged int/r15 into both lane branches (site 0d9241e0c with guides-r1 landed; fork 86fe4bf34), rebuilt spec-breakdown, and read guide-accuracy round 2 (`origin/land/guides-r2:design/reviews/guide-accuracy/round2/druid.md`). The druid's sim seeds there are S1 (Rend and Tear, fixed in wave 1, E e45a540db), S7 (Nature's Reach) and S8 (Feral Charge, Cat). No card-order item is assigned to the druid.

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| W4-1 | **The cat builds left Nature's Reach out, though it gives melee hit.**<br>**What was wrong:** the raid and leveling builds take Genesis 5 and Nature's Majesty 2 but not Nature's Reach. Its 70755 row 16819 carries melee hit (aura 54) next to spell hit (aura 55), 4 points with no class mask. The text is unscoped ("improves your chance to hit by 4%"). The engine has added melee hit since round 6 (`applyNaturesReach`, `talents.go:610`). The raid talent search that set the build listed it as "no damage" before that. The cat has 9 hit of headroom at raid 60 (`bis/druid-feral.json` hit_to_cap).<br>**What the probes said** (raid 60, 3000 iterations, then 10,000):<br>- Paying with Nature's Majesty loses 0.7% (4% crit is worth more than 4% hit).<br>- No Feral point can pay: the 31 are exactly Shifting Power's 30-point row plus itself.<br>- Genesis 5 is Balance row 1's gate.<br>- That leaves Natural Shapeshifter. It only cheapens Shifting Power, and mana never runs short at 1 or 0 points (out of mana 0.0 to 0.3 s of 180).<br>**What changed:**<br>- raidBuild `050022/50230032023012125001/055`: Nature's Reach 2, Blood Frenzy 1 to 2, Natural Shapeshifter 3 to 0. Raid 60 685.9 to 706.9 (alliance) and 681.5 to 701.3 (horde) at seed 1; +2.91 to +3.15% at both seeds. Natural Shapeshifter 1 with Blood Frenzy 1 was +2.64 to +2.81%.<br>- build `050022/50212032123012123001/05501`: Nature's Reach 2, Natural Shapeshifter 3 to 1. The ladder gives band 50 Nature's Reach 2 in place of two Heart of the Wild points. Bare 50 +1.66 to +1.88%, bare 60 +0.40 to +0.51%, bands 20 to 40 unchanged.<br>The adoption rule holds: the keep-list talents (Shifting Power, Leader of the Pack) stay, the Feral tree holds, nothing the rotation casts is dropped, and the gain is ≥1% at both seeds. | `TestCatBuildsTakeNaturesReach` (sim/leveling) failed ("raidBuild takes Nature's Reach 0, want 2") and passes. Client SpellEffect rows for 16819; `w4.tsv` (lane dir) | CONFIRMED-BUG (the build misses a modeled 4% hit) | old (raid build from the 2026-10-07 talent search, before E 9af6d73ac gave Nature's Reach melee hit; no round re-ran it), guides-handoff (round 2 S7) | S 0eeab841a |
| W4-2 | **Bear and Nature's Reach:** taking it costs Genesis 5 plus Nature's Reach 2, which can only come from Heart of the Wild 5, Naturalist 1 and Feral Swiftness 1. Probe (raid 60, horde, 3000 iterations): threat 993.1 to 1011.4 (+1.8%), but chance of death 9.3% to 20.1% without Heart of the Wild's 20% Stamina. Not adopted. | probe | OPEN (owner: threat +1.8% against double the deaths) | - | - |
| W4-3 | S8: Feral Charge (Cat) 414924 is not modeled. It is a leap with no damage, and no raid build takes it. | round 2 S8 | Aligned | - | - |

**Files to regenerate (wave 4):** the druid-feral ladder golden (bands 50 and 60), `bis/druid-feral.json` talents at band 50, bare 60 and raid 60, addon-builds, dps-per-point, Data.lua, and the cat coach starter if it carries the build.

**For other lanes:** `TestBandBuildsTakeTheCastTalentsThatFit` (sim/leveling) fails on int/r15 0d9241e0c without any druid change. warrior-fury level 30 is missing Death Wish (`ladder_first_walk_test.go:48`). That belongs to the warrior lane.

**For the guide lane (feral.md on main after guides-r1):**
- Line 58, replace "From 41 take Genesis, then Nature's Focus and three points of Heart of the Wild to open Restoration, then Naturalist, Natural Shapeshifter and Subtlety last. Nature's Focus keeps a mob's hits from pushing back the Healing Touch or Regrowth you cast between forms; Subtlety and Nature's Focus both move at the 60 respec." with "From 41 take Genesis, Nature's Majesty and Nature's Reach in Balance, then Heart of the Wild, Naturalist and a point of Natural Shapeshifter in Restoration." This supersedes the wave-1 replacement for the same sentence.
- Line 63, replace "(every time with both points while leveling, a 50% chance with the raid build's single point)" with "(every time with both points, in both builds)".
- Line 70, replace "The raid build's Balance points are Genesis, 10% more periodic damage (a bigger Rip), and Nature's Majesty, 4% crit. Natural Shapeshifter (3 points in both builds) makes Shifting Power cheaper, which is mana you spend every 8 seconds." with "The raid build's Balance points are Genesis, 10% more periodic damage (a bigger Rip), Nature's Majesty, 4% crit, and Nature's Reach, 4% hit on your attacks as well as your spells. Shifting Power's mana never runs short in a raid, so Natural Shapeshifter is not worth its points there."

### Live hotfix check (2026-11-03, after the later 70755 pushes)

I fetched wago.tools' live 70755 tables (`https://wago.tools/db2/<Table>/csv?build=1.60.1.70755`: SpellEffect, Spell, SpellMisc, SpellCooldowns, SpellDuration, CurvePoint, TraitDefinitionEffectPoints) and Wowhead's live Forever talent payload (`https://nether.wowhead.com/forever/data/talents-classic?dv=100`). I compared them against our layered raw. Lane dir: `live/`.

- **Every spell this lane fixed or relies on is unchanged live:** Rend and Tear 1223246, Nature's Reach 16819, Frenzied Regeneration 22842, Improved Regrowth 17074, Improved Moonfire 16821, Moonglow 16845, Genesis 1223081, Nature's Majesty 1223082, Heart of the Wild 17003, Naturalist 17069, Natural Shapeshifter 16833, Shifting Power 1322605, Regrowth 9858, Rejuvenation 9841 / 8910 / 774, Moonfire 9835, Wrath 9912, Insect Swarm 24977, Rip 9896, Shred 9830, Lacerate 1235827, Primal Bite 1238070, Maul 9881 and Swipe 9908. No row differs in SpellEffect, SpellMisc or SpellCooldowns.
- **Rend and Tear:** the live server text is still 3 / 6 / 9 / 12 / 15%. The client curve 85357 is still 2..10 live, so the engine's 3% (E e45a540db) stands.
- **Nature's Reach:** the live text is "improves your chance to hit by 2%" at rank 1 and 4% at rank 2, matching `applyNaturesReach` (2% a rank).
- Nothing to record for the BiS lane's re-layer.

## Hotfix diff (stopped)

The all-class diff was stopped for wave 8; the re-layer (site 43dc0f074) supersedes it. Two notes from the attempt:
- `druid/hfx/diff.json` is not a live diff. wago's db2 CSV is the shipped client without hotfixes: every one of its 1,371 differing rows is "live == shipped". It is not to be used (the hunter lane saw false "reversions" in it).
- My wave-4 "live check" used that same CSV, so it only showed that our own push layer left those druid spells alone. Wave 8 below re-checks the druid spells per record on the live push. That check found the late push did change seven of them (113403).

## Wave 8: the engine follows the late 70755 push

Site merged int/r15 (68ddc6dc1, which carries the re-layer 43dc0f074). Fork merged int/r15 (fbc5dbddf). `sim/druid/constants_auto_gen.go` was regenerated from the site's `spellconst/druid.json` (`go run ./sim/core/spellconst/gen`, the `make spellconst` step for druid). The conformance copy `sim/core/testdata/conformance/client/druid.json` was byte-identical to the pre-re-layer site file and now matches the new one. Every value was then confirmed by a per-record query on wago.tools, newest push per record (all ranks: push 113403, one entry each; `druid/hfx/perrecord.json`).

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| W8-1 | **The late push retuned the cat and Starfire; the hand-coded values had not followed.** Old to new, every rank:<br>- Claw weapon damage 110% to 140%.<br>- Shred 155% to 200%.<br>- Ravage 350% to 435%.<br>- Rake initial 23/31/45/61 to 27/39/56/76 (generated); tick 16/21/26/34 to 20/26/32/42.<br>- Rip tick base 3/4/6/8/11/15 to 4/5/8/10/14/19 (generated). Per combo point 4.4/7.2/8.5/12.7/18.2/25.5 to 5.5/9/10.5/15.8/22.5/31.9 (EffectPointsPerResource, not in spellconst).<br>- Ferocious Bite base 14/22/32/50/70/82 to 17/27/40/62/87/102, ±35.7% (generated). Per combo point 24/36/59/92/128/147 to 30/45/74/115/160/181. Per excess energy 0.6/1.0/1.5/2.0/2.5/2.7 to 0.7/1.2/1.8/2.5/3.1/3.3.<br>- Starfire, ranks 1-7: base 81/110/144/198/266/337/381 to 113/154/201/277/372/471/533; per level 1.3/1.5/1.6/1.8/2.2/2.3/2.4 to 1.8/2.1/2.2/2.5/3.0/3.2/3.4. Rank 6 at 471 is round 9's finding.<br>No other druid spell changed in the re-layer's spellconst diff. | Red first, on the refreshed conformance copy, every rank: `TestStarfireDamageMatchesClient`, `TestRakeDamageMatchesClient`, `TestRipDamageMatchesClient`, `TestFerociousBiteDamagePerEnergyMatchesClient`, `TestCatWeaponSpecialsMatchTheClient`. Red by hand from the push: `TestRipComboPointStepsMatchClient`, `TestFerociousBiteDamagePerComboPointIsThePushs`. All pass. | CONFIRMED-BUG | new-data (70755 server push 113403) | E 230a079c5 |
| W8-2 | Fork goldens: none moves. TestP1Feral and TestP1Balance skip (`SkipAwaitingForeverTalentRewrite`), and the bear's spells (Maul, Swipe, Lacerate, Primal Bite) are not in the push; TestFeralTank passes unchanged. | druid package tests | - | - | - |

Before/after (published BiS talents; 10,000 iterations). "Round start" is `$R/sb-before`; "Before follow" is engine int/r15 47e9c9e62 on site 68ddc6dc1; "After" is E 230a079c5 on the same site. The bear is in threat per second, with chance of death in the last column.

| Spec | Preset band | Faction | Seed | Round start | Before follow | After | vs round start | vs before follow | CoD start / before / after |
|---|---|---|---|---|---|---|---|---|---|
| druid-balance | raid 60 | alliance | 1 | 569.5 | 609.2 | 652.8 | +14.63% | +7.16% | - |
| druid-balance | raid 60 | alliance | 50001 | 570.3 | 609.8 | 653.5 | +14.59% | +7.17% | - |
| druid-balance | raid 60 | horde | 1 | 574.9 | 603.8 | 646.9 | +12.52% | +7.14% | - |
| druid-balance | raid 60 | horde | 50001 | 575.1 | 604.2 | 647.2 | +12.54% | +7.12% | - |
| druid-balance | bare 40 | alliance | 1 | 152.8 | 163.1 | 169.6 | +10.99% | +3.99% | - |
| druid-balance | bare 40 | alliance | 50001 | 152.8 | 163.1 | 169.6 | +10.99% | +3.99% | - |
| druid-balance | bare 40 | horde | 1 | 151.9 | 160.8 | 167.1 | +10.01% | +3.92% | - |
| druid-balance | bare 40 | horde | 50001 | 151.9 | 160.8 | 167.1 | +10.01% | +3.92% | - |
| druid-balance | bare 50 | alliance | 1 | 197.4 | 210.5 | 221.0 | +11.96% | +4.99% | - |
| druid-balance | bare 50 | alliance | 50001 | 197.4 | 210.3 | 220.8 | +11.85% | +4.99% | - |
| druid-balance | bare 50 | horde | 1 | 192.6 | 206.6 | 217.2 | +12.77% | +5.13% | - |
| druid-balance | bare 50 | horde | 50001 | 192.8 | 206.4 | 217.0 | +12.55% | +5.14% | - |
| druid-balance | bare 60 | alliance | 1 | 286.2 | 334.1 | 352.3 | +23.10% | +5.45% | - |
| druid-balance | bare 60 | alliance | 50001 | 286.2 | 334.2 | 352.5 | +23.17% | +5.48% | - |
| druid-balance | bare 60 | horde | 1 | 282.0 | 330.4 | 348.1 | +23.44% | +5.36% | - |
| druid-balance | bare 60 | horde | 50001 | 282.2 | 330.6 | 348.4 | +23.46% | +5.38% | - |
| druid-feral | raid 60 | alliance | 1 | 664.9 | 735.7 | 846.4 | +27.30% | +15.05% | - |
| druid-feral | raid 60 | alliance | 50001 | 664.8 | 735.8 | 846.4 | +27.32% | +15.03% | - |
| druid-feral | raid 60 | horde | 1 | 660.4 | 726.5 | 834.8 | +26.41% | +14.91% | - |
| druid-feral | raid 60 | horde | 50001 | 660.2 | 726.5 | 834.8 | +26.45% | +14.91% | - |
| druid-feral | bare 40 | alliance | 1 | 118.6 | 125.9 | 144.9 | +22.18% | +15.09% | - |
| druid-feral | bare 40 | alliance | 50001 | 118.4 | 126.0 | 144.9 | +22.38% | +15.00% | - |
| druid-feral | bare 40 | horde | 1 | 115.5 | 123.1 | 141.7 | +22.68% | +15.11% | - |
| druid-feral | bare 40 | horde | 50001 | 115.4 | 123.2 | 141.8 | +22.88% | +15.10% | - |
| druid-feral | bare 50 | alliance | 1 | 157.6 | 178.6 | 205.9 | +30.65% | +15.29% | - |
| druid-feral | bare 50 | alliance | 50001 | 157.5 | 178.6 | 205.9 | +30.73% | +15.29% | - |
| druid-feral | bare 50 | horde | 1 | 153.6 | 174.5 | 201.1 | +30.92% | +15.24% | - |
| druid-feral | bare 50 | horde | 50001 | 153.6 | 174.4 | 201.0 | +30.86% | +15.25% | - |
| druid-feral | bare 60 | alliance | 1 | 269.9 | 297.8 | 348.3 | +29.05% | +16.96% | - |
| druid-feral | bare 60 | alliance | 50001 | 270.2 | 297.9 | 348.4 | +28.94% | +16.95% | - |
| druid-feral | bare 60 | horde | 1 | 266.8 | 291.4 | 340.5 | +27.62% | +16.85% | - |
| druid-feral | bare 60 | horde | 50001 | 267.0 | 291.7 | 340.9 | +27.68% | +16.87% | - |
| druid-feral-bear | raid 60 | alliance | 1 | 964.2 | 1114.1 | 1114.1 | +15.55% | +0.00% | 10.9 / 0.1 / 0.1 |
| druid-feral-bear | raid 60 | alliance | 50001 | 965.2 | 1112.8 | 1112.8 | +15.29% | +0.00% | 9.9 / 0.1 / 0.1 |
| druid-feral-bear | raid 60 | horde | 1 | 955.7 | 1108.4 | 1108.4 | +15.98% | +0.00% | 10.1 / 0.1 / 0.1 |
| druid-feral-bear | raid 60 | horde | 50001 | 956.3 | 1109.1 | 1109.1 | +15.98% | +0.00% | 9.8 / 0.2 / 0.2 |
| druid-feral-bear | bare 40 | alliance | 1 | 172.5 | 202.0 | 202.0 | +17.10% | +0.00% | 87.2 / 80.2 / 80.2 |
| druid-feral-bear | bare 40 | alliance | 50001 | 172.4 | 202.0 | 202.0 | +17.17% | +0.00% | 87.9 / 80.1 / 80.1 |
| druid-feral-bear | bare 40 | horde | 1 | 170.1 | 200.2 | 200.2 | +17.70% | +0.00% | 89.3 / 82.1 / 82.1 |
| druid-feral-bear | bare 40 | horde | 50001 | 170.0 | 199.9 | 199.9 | +17.59% | +0.00% | 88.9 / 82.8 / 82.8 |
| druid-feral-bear | bare 50 | alliance | 1 | 306.4 | 336.1 | 336.1 | +9.69% | +0.00% | 56.1 / 29.2 / 29.2 |
| druid-feral-bear | bare 50 | alliance | 50001 | 306.3 | 335.8 | 335.8 | +9.63% | +0.00% | 56.3 / 30.1 / 30.1 |
| druid-feral-bear | bare 50 | horde | 1 | 302.1 | 312.1 | 312.1 | +3.31% | +0.00% | 56.7 / 31.8 / 31.8 |
| druid-feral-bear | bare 50 | horde | 50001 | 301.9 | 311.4 | 311.4 | +3.15% | +0.00% | 56.4 / 32.2 / 32.2 |
| druid-feral-bear | bare 60 | alliance | 1 | 417.2 | 477.3 | 477.3 | +14.41% | +0.00% | 77.7 / 8.4 / 8.4 |
| druid-feral-bear | bare 60 | alliance | 50001 | 416.6 | 477.0 | 477.0 | +14.50% | +0.00% | 77.6 / 8.6 / 8.6 |
| druid-feral-bear | bare 60 | horde | 1 | 416.7 | 473.6 | 473.6 | +13.65% | +0.00% | 87.0 / 15.0 / 15.0 |
| druid-feral-bear | bare 60 | horde | 50001 | 416.0 | 473.1 | 473.1 | +13.73% | +0.00% | 86.7 / 15.2 / 15.2 |

The round-start-to-before-follow move is the rest of int/r15: the re-layer's items and sets, other lanes' core fixes, and my wave-1 Rend and Tear. The follow's own effect is the "vs before follow" column: cat +14.9 to +17.0%, Balance +3.9 to +7.2%, bear 0.

**Files to regenerate (wave 8):** ladder goldens for druid-feral and druid-balance; BiS, weights and dps-per-point for druid-feral and druid-balance (Shred, Ravage and Starfire move a long way); Data.lua.

## Calls

| # | call | decision and why | sources | fix sha |
|---|---|---|---|---|
| K1 | Bear Furor vs Naturalist | **Furor 5 to 1, Naturalist 1 to 5** (`0/5123033202310212551/055`). A tank never shifts into Bear Form mid-fight, so Furor's rage never lands. Naturalist is threat. Raid and bare 60: threat +7.6 to +7.7%, chance of death and damage taken identical (both factions, seeds 1 and 50001). Bands 30-50: threat +1.9 to +5.9%, chance of death equal or lower (band 50 56% to 38%). Band 20: threat -0.5%, but chance of death 68-76% to 41-49%, because the freed points reach Feral Swiftness's dodge early. Survival over threat at a low band follows the Protection paladin band 20 precedent. Test: `TestBearBuildFinishesNaturalistOverFuror`, red then green. | `bear-calls.tsv`, `bear-bands.tsv` (lane dir); Warcraft Tavern Forever feral guide ("usually not needed in endgame group content", read 2026-11-03); round 6 G0 (the tank never shifts) | S 6de6cbf98 |
| K2 | Bear Nature's Reach | **Rejected.** It costs Genesis 5 plus Nature's Reach 2, paid by Heart of the Wild 5, Naturalist 1 and Feral Swiftness 1. Threat +1.4 to +2.0%, but chance of death about doubles at raid 60 (9.8-10.9% to 19.8-21.6%) and rises at bare 60 (87% to 95% horde). K1 gains four times the threat at no survival cost. | `bear-calls.tsv` | - |
| K3 | Pounce behind the target (guide round 3, S9) | **Aligned:** `pounce.go:75` already requires Prowl and not in front. **But Ravage was missing the same check:** its 70755 text (Spell 6785-9867) is "Must be prowling and behind the target", and `ravage.go` checked Prowl only. Fixed test-first (`TestRavageNeedsToBeBehindTheTarget`). No published fight stands in front, so no number moves. Origin old (758141eb6, 2026-09-29). | client Spell.csv descriptions; Wowhead Forever 9827 | E 725a50f11 |
| K4 | Feral Charge (Cat) leap (guide round 4) | **Aligned.** 414924 is "Leap behind an enemy.": no damage, no immobilize, no Prowl or opener clause. The cat leveling build takes Feral Charge from band 30 (it is also the bear Charge); no raid build takes it. The sim's cat always stands behind the target, so a leap behind adds nothing to model. The point is a real repositioning tool, not a provably dead one. | nether.wowhead.com/forever/tooltip/spell/414924 | - |

**For the guide lane:** feral-bear.md's talent prose for Furor and Naturalist needs to say the bear build keeps one Furor point and finishes Naturalist (10% more damage, so threat). The exact sentence is the guide lane's to place; I have not edited prose.

**Files to regenerate (calls):** druid-feral-bear ladder golden (every band), `bis/druid-feral-bear.json` band talents, addon-builds.
