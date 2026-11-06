# Sim-accuracy round 24: hunter lane

Specs: Beast Mastery, Marksmanship, Survival. Client 1.60.1.70790. Site branch wt/r24-hunter, fork branch lane/r24-hunter (both on int/r24 7f7b29fa5 / bacbf86df). Gear is the frozen BiS ($R/bis-frozen). Build mode: touched tests only, no wow-player gate.

Sources read: the round 23 hunter report and integration "Open, for round 24"; guide-accuracy round 26 hunter (R26-1..R26-4 and its optional measurements); Blizzard's Forever notes thread https://us.forums.blizzard.com/en/wow/t/2360696 (JSON, 5 Nov ~18:20 local: 10 posts, last 4 Nov 19:35, nothing new for hunters).

**Summary:** 4 confirmed bugs, all fixed:
- one engine bug: Mongoose Bite and Lacerate registered only their top rank;
- one method bug in the shared tool: the platform lane fixed it;
- the round 22 rank rows, judged on the wrong talents: two items, rows rewritten;
- one guide/card claim (R26-2, the sting note).

No raid-60 number moves. Every row change is bare bands only.

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **`spec-breakdown -level N` kept the band floor's talents.** A rank row judged above a band floor ran without the points the player has by that level. For example, Survival at 33 to 39 had no Resourcefulness, and Marksmanship and Beast Mastery lacked their mana talents. Every lane's rank rows above a floor were judged this way. Also, a forced `-rank` equal to the authored rank at bare 60 was overwritten by the bare-60 row, so "authored vs row" there compared the same rank twice. | `sim/cmd/spec-breakdown/main.go` rotationFor, `sim/request/force_rank.go` (as they were before the platform fix) | CONFIRMED-BUG (method) | old | platform lane (int/r24 second pass); the 193 suspect rows are in platform-suspect-rank-rows.tsv |
| X2 | **Mongoose Bite and Lacerate registered only their top rank,** so `-rank` on them dropped the line instead of casting the lower rank. Round 22 S1's "newest wins" for both compared the top rank against no cast. After the fix, the top rank still wins at 44-45 and 58-59 (Mongoose Bite) and at 40-41 and 50-51 (Lacerate). | `sim/hunter/melee_ranks_test.go` (red, then green) | CONFIRMED-BUG (engine) | old (round 19 S3 flagged the trap, round 22 S1 fell into it) | E ec4eafd62 (already on int/r24) |
| X3 | Round 23's change (the pet takes the scope crit once, E bc0291071): `core.Pet.Enable` sets the dynamic inheritance before OnPetEnable, so the override holds. | read | Aligned | - | - |
| X4 | Logs, 07:51 UTC 6 Nov: `api.foreversixty.gg/v1/reports/recent` lists only the 20 October report, and foreverlogs.gg serves no readable report list. No leveling hunter log exists to judge a rank row. | curl | no logs | - | - |

## Rank rows, re-judged on the talents at the level

**Method.** First, a coordinate descent over Arcane Shot, Serpent Sting and Aimed Shot (Survival: Raptor Strike, Mongoose Bite and Lacerate). It ran at every level from 20 to 60 with the platform tool's ladder talents: 10,000 iterations, seed 1, Troll (`$R/hunter/desc2.tsv`). Each win was then confirmed at seeds 1 and 50001, as Troll (Horde) and Dwarf (Alliance), the published races (`conf_mm.tsv`, `conf_mm2.tsv`, `conf_mm3.tsv`, `conf_bm.tsv`, `conf_sv.tsv`; tool `tools/confirm.py`). A row is adopted on any gain beyond the summed error at both seeds. Below 20 the bare sim has no ranged weapon at 12 to 19, so it was not judged there.

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| K1 | **Marksmanship Arcane Shot from 33 cast rank 3 (33-39) or rank 1 (40-59).** On the hunter's real talents (Rapid Recuperation from 33, the mana talents), his mana holds up. So the rows were wrong:<br>- the newest rank wins at 33 to 41 (no row now);<br>- rank 4 wins at 42 to 47, where Serpent Sting rank 6 arrives;<br>- rank 5 wins at 48 to 59.<br>Gains over the old row are 1.6 to 4.1% for the Troll and 0.6 to 3.8% for the Dwarf.<br>Race keys:<br>- at 48 the Troll takes rank 5 (+0.2% over rank 4) and the Dwarf rank 4 (+0.2%);<br>- at 54 to 56 the Dwarf takes rank 6 (+0.2 to 0.4%), while the Troll keeps rank 5 (a tie at 56);<br>- at 57, rank 6 wins for both;<br>- at bare 60, rank 4 beats rank 1 by 1.7% (Troll) and 2.0% (Dwarf), and beats ranks 3, 5, 6, 7 and 8.<br>Raid 60 keeps the top rank. | conf_mm*.tsv; `TestRound24MarksmanshipArcaneShotRankRows` | CONFIRMED-BUG (rank row) | R22 aae25e43b (judged at band floors on the floor's talents) | S 1b396ffb9 |
| K2 | **Beast Mastery Arcane Shot and Serpent Sting.**<br>- Rank 1 of both now starts at 20, where the old rows started at 26: +1.5 to 2.8% at 20 to 25.<br>- Arcane Shot rank 3 beats rank 1 at 34 to 39 (+1.2 to 1.8%).<br>- Serpent Sting rank 2 beats rank 1 at 38 and 39 (+0.1 to 0.3% on top).<br>- Arcane Shot rank 4 beats rank 3 at 54 to 57 (+0.2 to 0.6%), and at 53 for the Troll only (+0.1%; the Dwarf ties).<br>- At 58 rank 3 wins again (rank 4 -0.6 to -0.8%).<br>The 40 to 49 (rank 1) and 50 to 52 (rank 3) rows hold. So do bare 60 (rank 3) and the Aimed Shot rank 1 rows. | conf_bm.tsv; `TestRound24BeastMastery*RankRows` | CONFIRMED-BUG (rank row) | R22 aae25e43b | S 1b396ffb9 |
| K3 | **Survival Raptor Strike at 33 cast rank 5.** At 33 only the first point of Resourcefulness is in. Rank 1 beats rank 5 there by 1.38% for both races and both seeds. At 34 the second point is in, and rank 1 loses by 2.3%. At 40 the Troll gains 0.18 to 0.23% with rank 5 over rank 6, beyond error at both seeds. The Dwarf ties, and at 41 the Troll's gain is inside error. The round-21 rows (26-29, 46-49) hold for both published races. They were not re-run for the other six races. | conf_sv.tsv, desc_sv_all.tsv; `TestRound24SurvivalRaptorStrikeRankRows` | CONFIRMED-BUG (rank row) | old (round 21 judged at the band floors) | S 1b396ffb9 |
| K4 | **The Marksmanship and Beast Mastery sting note said the sting "does more for its mana than Arcane Shot".** At Beast Mastery 50 to 59, Arcane Shot rank 3 beats the sting per mana at every level: 4.7 to 6.1 damage per mana against 3.8 to 4.3 (`s2.txt`). The comparison is gone. The Multi-Shot comparison from 18 stays. At raid 60 the guide's "sting best for its mana, Arcane next" still holds: BM 6.17 / 4.40, MM 8.33 / 6.30 (`s1.txt`). | guide round 26 R26-2 | CONFIRMED-BUG (card text) | guides-handoff | S ceb5e00e8, E d8de466ba |
| K5 | The guide round 26 Aimed Shot claims hold, measured with the ladder talents (`s3.txt`):<br>- rank 1's hit as a share of the newest rank's: BM 28 88%, 36 78%, 44 70%, 52 62%;<br>- the same for MM: 52 64%.<br>So BM's "well over half" is true, and MM's "most" holds, loosely. | s3.txt | Aligned | - | - |

Guide text shipped with the rows:
- beast-mastery.md, the Arcane Shot and Serpent Sting paragraph in Leveling;
- marksmanship.md, the same paragraph;
- survival.md, a new 33 paragraph and the Troll at 40 in the 46 paragraph;
- the Survival Raptor Strike APL note, first sentence and the rank list (curated plus the sim/request copy). Its first sentence is shorter than before, so it still fits the panel.

## Before/after (bare bands; row vs new row, 10,000 iterations)

The figures are the gain of the new row over the old one, at seeds 1 / 50001.

| spec | levels | Troll | Dwarf |
|---|---|---|---|
| MM | 33-41 | +1.6 to +4.1% | +0.6 to +3.4% |
| MM | 42-47 | +2.5 to +2.9% | +2.0 to +2.7% |
| MM | 48-59 | +2.2 to +3.7% | +2.7 to +4.4% |
| MM | bare 60 | +1.65 / +1.65% | +1.96 / +1.96% |
| BM | 20-25 | +1.9 to +2.8% | +1.5 to +2.2% |
| BM | 34-39 | +1.2 to +2.0% | +1.4 to +2.2% |
| BM | 53-57 | +0.1 to +0.6% | tie to +0.5% |
| SV | 33 | +1.38% | +1.38% |
| SV | 40 | +0.23 / +0.18% | (no row) |

Absolute values at seed 1, Troll:
- MM 40: 221.1 to 230.0;
- BM 37: 164.7 to 167.4;
- SV 33: see conf_sv.tsv.

Raid 60 is unchanged for every spec.

## Files to regenerate
- The ladder goldens for hunter-marksmanship, hunter-beast-mastery and hunter-survival (sim/request TestRotationLadder), because the bare-band rows moved.
- addon-data / Data.lua, for the sting note and the Survival Raptor Strike note.
- The fork APL copy ui/hunter/apls/forever_survival.apl.json carries an older Raptor Strike note than curated (from before this round). The apl-sync copy should be refreshed by the integrator.

## For other lanes
- **Integrator:** on int/r24, `TestCuratedRankChoicesAreCastableDownranks` already fails, on the paladin-protection Consecration rows (rank 3 from 49, rank 4 from 58: "what the rotation casts anyway"). It is not hunter, and it fails the same without this lane's changes.

## Measure at refresh
- The Survival race-keyed Raptor Strike rows at 26 to 29 and 48 to 49 for the six races that are not published (Orc, Tauren, Windshaper Skyborne, Human, Night Elf, High Order Skyborne), on ladder talents.
- The Marksmanship Arcane Shot race splits (48, 54 to 56) for the other six races. The rows key only the Troll or Dwarf side and give the rest the other side's rank.

## Still open
- Bare sim below 20: the band-20 entry wears no ranged weapon at 12 to 19, so Multi-Shot deals NaN and there is no Auto Shot. Rank rows below 20 cannot be judged. The engine also lets shots cast with no ranged weapon (minor).
- Pending owner items unchanged (Arcane Shot tooltip with spell power; level-60 book ranks).
