# Sim-accuracy round 23: rogue

Lane: site `wt/r23-rogue` (from int/r23 1da2b60c4), fork `lane/r23-rogue` (from int/r23 84ce92bdd). **One confirmed bug, origin R22:** with round 22's 31/20/0 Assassination raid build, the card's Rupture line now loses 1.5% at raid 60; it is gone. Round 22's raid-build adoption itself still holds on the corrected parser path. Combat and Subtlety: aligned.

Measurements: spec-breakdown, frozen BiS (`$R/bis-frozen`), guide builds via `$R/rt` (the int/r23 helper prints Assassination's raid build `00530311251051031-30230302020122-` with the empty third tree), 10000 iterations at seeds 1 and 50001 for reported numbers, 3000 to explore. Variants and outputs: `$R/rogue/a-*.json`, `$R/rogue/bd-*.txt`; tools `$R/rogue/tools/`.

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix |
|---|---|---|---|---|---|
| K1 | **Assassination raid 60 on the corrected raid-build path (S d8c453fdf): round 22's adoption holds.** | Raid 60, frozen BiS, current card: raid build 31/20/0 780.6 / 780.6 (Alliance), 778.8 / 778.4 (Horde); round 22's old raid build `00530311251551031-302302-005` 759.4 / 759.2 and 762.1 / 762.3 (+2.8% / +2.1% for the new build, above the 1% bar at both seeds); the leveling build every raid-build reader used to get (`32530200251551031-302303-011`) 725.7 / 725.7 and 725.5 / 724.7 (-7%). Note: spec-breakdown without `-talents` reads the BiS file's talents, which still carry the old raid build (759.8 at 3000): the re-rank fixes that. | Aligned | - | - |
| K2 | **Rupture in the Assassination raid rotation loses since the 31/20/0 build.** The card cast Rupture at four or more points whenever it was not ticking, gated on Blood Splatter so only the raid build reached it. Coup de Grace 3, Aggression 2 and Flawless Execution make Eviscerate a 20-energy finisher that is 30% harder than the old build's; Rupture at 60 is `(35 + 4.73 x cp)` a tick (client 11275, Wowhead Forever tooltip "58.65 x 8 at five points", plus 24% AP over the run), so a four-point Eviscerate out-damages a whole four-point Rupture up front. | Raid 60, raid build, 3000: card 780.8 / 781.0, no Rupture 793.3 / 792.9, Rupture at five only 783.7 / 782.7, Rupture without the Slice and Dice gate = card. Old raid build: card 759.8 / 759.9, no Rupture 754.0 / 754.5 (Rupture gained there, so round 22's build swap created the loss). With Rupture gone, the three-point Slice and Dice still gains (787.0 / 788.0 without it) and Eviscerate at four still beats five (768.7 / 768.4 at five). Vigor 2 in place of Blood Splatter ties (792.4 / 792.4), so Blood Splatter stays for the opener Garrote. Bare 40 and 60 identical before/after (the leveling build has no Blood Splatter). `TestAssassinationCastsNoRupture` red on the old card, green after; curated card, request copy, fork copy, top note; guide: talent paragraph (Blood Splatter now the Garrote only) and priority (Rupture is not on the list, and why). | CONFIRMED | R22 a0a2818ed (raid build) / 31d269ec3 | S ed9369755, E f242077f4 |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix |
|---|---|---|---|---|---|
| X1 | Blizzard's Forever notes since round 22 | Thread 2360696 JSON fetched 5 Nov 12:54 UTC: last post still #10 (4 Nov 19:35), no edit after 4 Nov 19:35. No new rogue line; round 22 X1 stands | Aligned | - | - |
| X2 | Server hotfixes | `hotfix-push.json` (70790): 0 rows added/changed in every table | Aligned | - | - |
| X3 | Live logs | 5 Nov 12:54 UTC: api.foreversixty.gg `/v1/reports/recent` only the 20 Oct report; foreverlogs.gg `/api/reports/public` 403; Warcraft Logs v1 `/v1/zones` 401 on www/classic/fresh/vanilla/sod | OPEN | - | - |
| X4 | Guide-accuracy round 24 rogue "Sims requested" | Subtlety Hemorrhage at five: done in round 22 (S1). Rest carried to the re-rank | Aligned | - | - |
| X5 | Assassination poison hands on the new raid build (it dropped Improved Poisons) | Raid 60, 10000: Instant main / Deadly off 780.6 / 780.6 and 778.8 / 778.4; swapped 771.6 / 772.4 and 769.8 / 769.4. The guide's swap at 60 holds | Aligned | - | - |

## Assassination

| # | finding | evidence and source | confirmed? | origin | fix |
|---|---|---|---|---|---|
| A1 | Rupture line | K2 | CONFIRMED | R22 | S ed9369755, E f242077f4 |
| A2 | Raid-build talents against the client rank text (`talents/rogue.json` 70790) | Coup de Grace 7/13/20% (`client_values.go` coupDeGraceBonus), Aggression 5%/rank, Flawless Execution -15, Composure -5/-10 on Poised (`poised.go`), Puncturing Wounds Mutilate +5%/rank, Lethality +10%/rank, Vigor +5/rank, Weapon Expertise 1%/rank, Seal Fate 33/66/100, Ruthlessness 20%/rank, Relentless Strikes 4% per point per rank, Venom 30%/20% additive with Vile Poisons. Breakdown cross-checks: Seal Fate 28.9 points from 43.7 Mutilates at 42.4% crit a strike (1 - 0.576^2 = 66.8%), Ruthlessness 17.5 from 29.2 finishers, Relentless Strikes 25.2 procs from 125.7 points spent, Mutilate 60 energy, Garrote 30 with Poised and Composure | Aligned | - | - |
| A3 | Stormshroud 3-piece (published BiS) | Revitalize 23863 ProcChance 4, ProcTypeMask 20 in the client; 12.8 procs a fight on about 330 melee hits matches | Aligned | - | - |

## Combat

| # | finding | evidence and source | confirmed? | origin | fix |
|---|---|---|---|---|---|
| C1 | Raid build points | Lightning Reflexes 4 and Composure 2 are the toll to Flawless Execution, Weapon Expertise and Aggression (15 points in rows 1-3); every other option in those rows (Efficient Strikes, Deflection, Improved Sprint, Endurance, Riposte) does nothing for a Backstab rogue either, and Hack and Slash 5 cannot take a toll point without breaking the row requirement | Aligned | - | - |
| C2 | Breakdown | Puncturing Wounds extra point on 45% of 38.8 Backstab hits (17.5), Backstab crit 60.7%, Blade Flurry 2 casts and Adrenaline Rush 1 in 180 s, glancing 40% at 63 | Aligned | - | - |

## Subtlety

| # | finding | evidence and source | confirmed? | origin | fix |
|---|---|---|---|---|---|
| S1 | Round 22's Hemorrhage gate | Card still gates Hemorrhage at four or fewer; costs check (Hemorrhage 30 with Efficient Strikes, Ghostly Strike 40, Garrote 30 with Dirty Deeds, 20 with Poised) | Aligned | - | - |
| S2 | Toll points (Master of Deception 3, Camouflage 3, Elusiveness 2) | Sub needs 25 points in rows 1-5 for Boot Knife; without them the tree reaches 22. The alternatives in those rows (Setup, Improved Ambush on a sword, Dirty Tricks, Heightened Senses, Quick Getaway) do nothing on a boss either | Aligned | - | - |
| S3 | Preparation and Premeditation | Preparation once a fight resets Premeditation (3 casts), never while a Death Mark stores (`preparation.go`) | Aligned | - | - |

## Before/after

Raid 60, guide builds, frozen BiS, 10000 iterations, seeds 1 / 50001, against `$R/before.tsv`:

| spec | faction | before | after | change |
|---|---|---|---|---|
| rogue-assassination | alliance | 780.6 / 780.6 | 792.8 / 792.8 | +1.6 / +1.6% |
| rogue-assassination | horde | 778.8 / 778.4 | 790.2 / 790.6 | +1.5 / +1.6% |
| rogue-combat, rogue-subtlety | both | as before.tsv | unchanged | 0 |

Bare bands unchanged (bare 40 129.3 / 129.3, bare 60 319.5 / 319.5, Horde seed 1, 3000). No spec outside the rogue moves: the only engine commit is the rogue APL copy.

## Files to regenerate

- `sim/request/testdata/ladder/rogue-assassination.golden.md` (`TestRotationLadder/rogue-assassination` differs after the change; matched on int/r23 before it).
- `addon-data.json` and `Data.lua`: Assassination raid card (no Rupture row). Together with platform's `addon-builds.json` (Assassination raid order).
- Assassination raid `dps-per-point` / stat weights, and coach starters if they read the rotation.

## Measure at refresh

- Assassination raid 60 on the re-ranked BiS (the BiS file still carries the old raid talents); confirm no Rupture still wins there.
- Carried from round 22: rogue picks with one-hand axes, Subtlety Hemorrhage gate and 50-energy Shank on the re-ranked gear, rogue races, Subtlety band 20 weapon.

## For other lanes

- **Data tests:** `data/tests/test_apl.py` fails 13 tests on int/r23 with or without this change: the data venv's protobuf has no `manaIncomePerSecond` (platform's new APL value, used by the paladin card). Not a rogue file.

## For the guide lane

- Done here: assassination.md talent paragraph (Blood Splatter keeps the opener Garrote longer; no Rupture claim) and the priority (Rupture dropped with its reason). `_prose.test.ts` and the guide content tests pass.

## Guide text not done

- None.

## Still open

- Death Mark storage and the Mutilate / Shank off-hand penalty (40+ log); the doubled poison coefficients and Backstab's flat bonus on a live log. Combat Eviscerate at five and the level-60 book ranks stay with the owner.

## Sources

- Client 1.60.1.70790: `talents/rogue.json`, `spellconst/rogue.json` (11275, 11300, 23863), `spelltips/rogue.json` (11275), `hotfix-push.json`, `core/client_sets_gen.go` (set 142).
- Wowhead Forever tooltips, 5 Nov about 13:00 UTC: https://nether.wowhead.com/forever/tooltip/spell/11275, /spell/11300.
- Blizzard, WoW Forever Beta Development Notes: https://us.forums.blizzard.com/en/wow/t/2360696 (10 posts, fetched 5 Nov 12:54 UTC).
- Round 22 rogue, platform round 23 (E1, E2) and guide-accuracy round 24 rogue reports.
