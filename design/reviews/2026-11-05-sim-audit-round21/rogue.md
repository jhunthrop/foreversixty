# Sim-accuracy round 21: rogue

Lane: fork `lane/r21-rogue` (from int/r21 90f97679c), site `wt/r21-rogue` (from int/r21 f07babdd9). No engine bug. Two Subtlety card fixes: Shank's combo-point hold now lifts below 50 energy (a measured gain on every band from 30 up), and the two finishers carry a branch for a rogue who has not trained Rupture (a card line no rogue below 20 could meet, carried since round 11). Assassination's poison hands below 60 measured level by level; the guide line now names 30 to 59.

Measurements: spec-breakdown, 10000 iterations, seeds 1 and 50001, frozen BiS (`$R/bis-frozen`), guide builds (`$R/rt`) at raid 60, the leveling order truncated by level on bare bands. Before = `$R/before.tsv` (int/r21 gives the same figures for every rogue spec). Outputs in `$R/rogue/`: `shank.tsv`, `shank2.tsv`, `shank3.tsv`, `hands.tsv`.

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix |
|---|---|---|---|---|---|
| K1 | **Subtlety Shank held at four or five combo points cost energy on every leveling band** (guide r19 seed: the gate before Boot Knife) | Shank (1329283 / 1329549) awards no combo point of its own; only Boot Knife (1329286, 20% a point) adds two. The card held it at four or five from 24. Measured, horde bare, the hold lifted entirely: 24 and 28 identical (Shank never meets the gate there: 18.1 casts a fight, its cooldown cap), 30 +0.2%, 34 +0.3%, 40 +0.2%, 50 +0.8%, 55 +0.9%, 59 +1.1%, but bare 60 -0.6% and raid 60 -0.6% (799.6 vs 804.3). The level-60 swing is Death Mark rank 4 (1325758, cap 156 + 245 a point against rank 3's 138 + 217): at 60 with rank 3 forced, the lifted hold wins again (+0.15%); on band 50 gear at level 60 the hold still wins, so it is the rank, not the gear. No APL value tells 59 from 60 (a ranked `spellIsKnown` resolves to the level's own rank, so `spellIsKnown(1325758 r4)` is true from 40). The real-play rule that captures both: hold Shank at four or five only with 50 energy or more. `cp <= 3 or energy < 50` against the old `cp <= 3`, seeds 1 / 50001: bare 24 45.3 = 45.3; 30 88.9/89.0 -> 89.1/89.1; 34 90.0/90.1 -> 90.2/90.2; 40 148.4 -> 148.7; 44 166.3/166.1 -> 166.5/166.5; 50 219.0/218.9 -> 221.5/221.5 (+1.2%); 56 247.0/246.8 -> 250.2/250.2 (+1.3%); 59 247.0/246.8 -> 250.2/250.3 (+1.3 to +1.4%); bare 60 374.0/374.1 -> 374.3/374.1; raid 60 alliance 809.9/810.2 -> 810.5/811.1, horde 804.3/804.3 -> 805.3/805.1 (+0.1%). Thresholds 30/40/60/70 tried at 5000 iterations; 70 loses at raid 60, 30 to 60 tie there, 50 is best on the 50s bands | CONFIRMED (rotation; gain beyond error at both seeds, no band loses) | old (the hold dates from round 16; guide r19 flagged 24-32), guides-handoff | S 1039da44a, E b0e54a7a6 |
| K2 | **Subtlety Eviscerate and Slice and Dice were gated on Rupture ticking at 10 to 19** (SIM-R11-1, carried by guide rounds 11-21) | Rupture trains at 20 (`trainables/rogue.json` 1943). The engine nils a condition on an unlearned spell, so it already cast Slice and Dice at three points below 20; the card told a level 10-19 player to wait for a Rupture he cannot have. Both lines now read `dotIsActive(Rupture) or not spellIsKnown(Rupture)` (Eviscerate: also "while Slice and Dice is up", which is what the engine does: Slice and Dice at three falls off before three more points are built, so every set goes into it). Engine casts identical at 12/14/16/18 (band-10 test entry with level-14 weapons 1936/14151, `$R/rogue/bis10`) and from 20 up (the branch is false once Rupture is known). Tests `TestSubtletyFinishersBeforeRupture`, `TestSubtletyShanksWhenShortOnEnergy` (sim/request) fail on the old file, pass on the new | CONFIRMED (card) | old, guides-handoff | S 1039da44a, E b0e54a7a6 |
| K3 | Assassination poison hands at 34 and 38 (guide r20 seed) | Horde bare, Deadly MH / Instant OH against Instant MH / Deadly OH, seeds 1 / 50001: 30 75.2/75.2 vs 73.1/73.0; 34 76.7/76.6 vs 74.5/74.4 (+2.9%); 38 86.8/86.7 vs 83.8/83.8 (+3.5%); 44 146.5/146.5 vs 145.7/145.5 (+0.6%); 48 164.5/164.4 vs 163.7/163.6 (+0.5%); 54 and 58 tie. The rule holds; assassination.md:88 now says "From 30 to 59" (guide r20 S2, no Deadly before 30) and that the gap closes into a tie in the 50s | Aligned (guide text sharpened) | - | S 1039da44a |
| K4 | Do Vile Poisons, Venom and Malice scale the poisons' new attack-power share? (guide r20 engine question) | Yes, as they should. The AP share is added to the base roll before any modifier (`poisons.go`:199 Deadly per stack, :278 Instant), so Vile Poisons (spell DamageMultiplier, :178/:264), Venom (`scalePoisonDamage`, venom.go:57) and the outcome roll all apply to the whole hit, the way the server applies percent modifiers to base plus coefficient. Malice is aura 290 (14138, 70790 SpellEffect 691045), and `stats.Crit` feeds every table (`spell_result.go`:130-136), so it reaches the poisons' spell-table crit. Raid 60 Instant Poison crits 20.2%, Deadly ticks 21.8% | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix |
|---|---|---|---|---|---|
| X1 | Every rogue line in Blizzard's Forever notes, 24 Sep to 4 Nov (thread posts #1-#10) | Read from the cached thread text (`$R/paladin/notes/raw.txt`). 2 Oct: Expose Armor miss, Setup target (dead talents, X3 of round 20). 16 Oct: poisons vs attackers, Sap, spec swap energy, Cold Blood consumed by a Mutilate crit (Mutilate's two strikes carry SpellFlagColdBlooded, one charge). 3/4 Nov: Poised 18, Shank, the poison AP coefficients (round 20), Lethality, Vigor, Blood Splatter 6/9 s, Relentless Strikes 25 energy at 4% a point a rank, Venom 30%/20%, Seal Fate, Mutilate 90%, Coup de Grace 7/13/20%, Efficient Strikes 3/5, Composure, Flawless Execution 15, Aggression 5/10/15% on Sinister Strike and Eviscerate, Dual Wield Specialization 25/50%, Flash of Steel 3 energy and 0.4 s a rank, Hemorrhage +15% Rupture and Garrote, Ghostly Strike 7 s, Opportunity 4% a rank on Backstab/Garrote/Ambush/Mutilate, Boot Knife, Death Mark. Each matches `client_values.go`, `talents.go`, `shank.go`, `death_mark.go`. Nothing in the notes is missing from the engine | Aligned | - | - |
| X2 | Cached 70790 rows against what is live (the druid lane's lead) | Fetched Wowhead's Forever tooltip for all 172 rogue spell and talent ids (`$R/rogue/wh`, diff tool `$R/rogue/tools/whdiff.py`). Every trainable spell's numbers match the client text (ranges against midpoints, Mutilate 39.6 shown as 40). Talent tooltips print the talent spell's own base points, not the trait curve the client rank text and the engine use (Blood Splatter 4/6 against the curve's 6/9, Flash of Steel 2.5 s against 2.0, Lethality 6% against 10/20%, Seal Fate and Relentless Strikes 0%): the notes agree with the curves every time, so the curves stay. No Forever-only change like Bear Form's 115% exists for the rogue | Aligned | - | - |
| X3 | Live logs | Last looked 04:31 UTC 5 Nov: api.foreversixty.gg `/v1/reports/recent` still only the 20 October report; foreverlogs.gg answers "API access not allowed" to a plain request (platform's 03:02 check with browser headers: empty); platform found no Forever zone on Warcraft Logs. Nothing for Death Mark, Mutilate or the doubled poison coefficients | OPEN | - | - |

## Assassination

| # | finding | evidence and source | confirmed? | origin | fix |
|---|---|---|---|---|---|
| A1 | Poison hands below 60 | K3 | Aligned | - | S 1039da44a (prose) |
| A2 | Spell-by-spell | Mutilate (1241584 and ranks, 90% + 39.6/56.7/79.2, +20% poisoned), Shank, Venom, Cold Blood, Seal Fate, Relentless Strikes, Eviscerate with Coup de Grace, Rupture with Blood Splatter, Slice and Dice, Vile/Improved Poisons, Malice, Lethality, Opportunity, Puncturing Wounds (30/15/45): client text = Wowhead live = engine (X1, X2); poison scaling K4 | Aligned | - | - |

## Combat

| # | finding | evidence and source | confirmed? | origin | fix |
|---|---|---|---|---|---|
| C1 | Spell-by-spell | Backstab, Sinister Strike, Eviscerate (five, pending owner), Slice and Dice, Shank with Flash of Steel (curve, X2), Blade Flurry (25 energy, 2 min, +20% attack speed: talents.go:396-401), Adrenaline Rush (100%, 15 s, 5 min), Hack and Slash 1% a rank, Aggression, Dual Wield Specialization, Efficient Strikes, Flawless Execution: as X1/X2. Combat's Shank has no combo-point hold, so K1 does not apply | Aligned | - | - |

## Subtlety

| # | finding | evidence and source | confirmed? | origin | fix |
|---|---|---|---|---|---|
| S1 | Shank hold | K1 | CONFIRMED | old, guides-handoff | S 1039da44a, E b0e54a7a6 |
| S2 | Finishers below 20 | K2 | CONFIRMED (card) | old, guides-handoff | S 1039da44a, E b0e54a7a6 |
| S3 | Spell-by-spell | Hemorrhage 100/145% + 15% bleeds, Ghostly Strike 125/180% and 7 s window, Garrote, Rupture, Death Mark ranks 40/48/56/60 and caps, Boot Knife (2 points, trigger row 1329285 against the tooltip's 1), Premeditation, Preparation, Serrated Blades 3% a rank, Initiative: as X1/X2. Death Mark storage stays OPEN on a 40+ log | Aligned (storage OPEN) | - | - |

## Before/after

Raid 60, guide builds, frozen BiS (seeds 1 / 50001): Subtlety alliance 809.9 / 810.2 -> 810.5 / 811.1 (+0.07 / +0.11%), horde 804.3 / 804.3 -> 805.3 / 805.1 (+0.12 / +0.10%). Assassination and Combat unchanged (777.3 / 776.8, 780.7 / 781.4; 798.1 / 798.5, 800.9 / 800.9). Bare bands, Subtlety horde: see K1 (+0.2% at 30 up to +1.4% at 59, 0 at 24 to 28 and at 60). No spec outside rogue moves (the only engine-side change is the fork's copy of the Subtlety rotation).

## Files to regenerate

- Ladder golden `sim/request/testdata/ladder/rogue-subtlety.golden.md` (Shank line; DPS at 30 to 50 rungs).
- addon-data / addon-builds / Data.lua (Subtlety rotation rows: the Shank, Eviscerate and Slice and Dice notes and conditions; the guide card renders from addon-data, so the page shows the old Shank card text until then).
- Subtlety stat weights / dps-per-point if they read the rotation.

## Measure at refresh

- K1 on the re-ranked BiS: the 50-energy Shank threshold at raid 60 and bare 50/60.
- Carried: rogue BiS and races at the re-rank; Assassination Instant main hand at 60 on the new BiS; Subtlety band 20 weapon (K4 of round 20).

## For other lanes

- Any lane gating on a level-60-only rank: a ranked `spellIsKnown(<id>, rank N)` resolves to the level's learned rank, so it is true from the first rank's level, not from rank N's. It cannot tell 59 from 60.
- Wowhead's Forever tooltips print a talent spell's base points, not its trait curve: they disagree with the curves (and with Blizzard's notes) on many talents, so they are not evidence against a curve value.

## For the guide lane

- Done here: subtlety.md priority paragraph and a new Levels 1 to 19 line (finishers before Rupture), index.md Shank bullet, assassination.md Poisons paragraph. `_prose.test.ts` passes.
- index.md:27 "for Assassination until 60" (guide r21 S1) could read "from 30 until 60", matching assassination.md now.

## Rendered review

wow-player on the rendered pages (astro dev on this worktree after `npm run sync`, 1440 px, captures in `$R/rogue/shots`) plus the curated card notes: **SHIP**, no MUST. Shoulds taken: the before-Rupture sentence moved from "Levels 20 to 60" to "Levels 1 to 19" and names "Eviscerate at three or more while Slice and Dice is up"; the class-page bullet says "overflow its combo points"; the Shank card note reads "at three combo points or fewer, or under 50 energy" and gives Boot Knife from 35, 5/5 by 39. Not taken: rewording the `expected_idle` reasons for 11300/6774 (engineering text, not on any page; they still hold from 40 and on the 20s rungs).

## Still open

- Death Mark storage and the Mutilate / Shank off-hand penalty (40+ log); the doubled poison coefficients on a live log with a known attack power. Combat Eviscerate at five stays with the owner.

## Sources

- Blizzard, WoW Forever Beta Development Notes thread, posts 24 Sep to 4 Nov 2026 (Kaivax): https://us.forums.blizzard.com/en/wow/t/wow-forever-beta-development-notes-%E2%80%93-updated-november-2/2360696 (cached text read 2026-11-05)
- Wowhead Forever tooltips: https://nether.wowhead.com/forever/tooltip/spell/<id> (172 ids, fetched 2026-11-05 04:00 UTC)
- Client 1.60.1.70790: `spelltips/rogue.json`, `talents/rogue.json`, `trainables/rogue.json`, wago SpellEffect (`scratchpad/b70790/70790`)
