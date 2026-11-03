# Sim-accuracy round 14: platform lane

Site branch `wt/r14-platform` from int/r14 (72471fcb3). Engine branch `lane/r14-platform` from 5195a8889. Scope: stale guide build codes in fixtures; a regression pass over the landing's platform paths; the /sim page, wasm and sims API against spec-breakdown; what the sim pages and addon export show. Plus three items the coordinator passed on from other lanes (guides, regcast).

**Found: 6 CONFIRMED-BUG rows, all fixed.** By origin: R13 1 (F1), landing 1 (L1), old 4 (F2, D1, C1, C2). R12 and new-data: 0. Two of the six were found by other lanes: C1 by guides, C2 by regcast. Only D1 moves a published number, and only by ±0.03% (mage-arcane, warlock-affliction, warlock-demonology).

## Item 1: copies of guide build codes

Method: a scan (`$R/platform/tools/stale_codes.py`, `stale_bare.py`) collected every `build`/`raidBuild` code the guides have ever had (git log of `web/src/content/guides`, 218 distinct) and every current one (45). It then looked for stale codes, FS1 or bare 3-tree strings, across apps/, docs/contracts/, web/src, web/tests, addon/tests, api, sim and companion.

| # | finding | evidence and source | confirmed? | origin | fix sha |
| --- | --- | --- | --- | --- | --- |
| F1 | **The Elemental guide fixture held the 70707 leveling build, before Tidal Focus.** `docs/contracts/fixtures/app/cdn.guides.shaman-elemental.json` had Restoration `503352` (Improved Healing Wave) where the guide now has `053352` (Tidal Focus). The apps' planner golden (`apps/shared/fixtures/planner-rules.json`) replays these codes, so it tested the apps' planner on a build no guide recommends. Fix: only `build` and the two build objects come from the W3-C1 generator (`appGuides`, `web/src/lib/app-data/app-files.ts`). The prose stays, so the app guide screenshots do not move. The planner golden is regenerated with its generator (`node web/scripts/gen-planner-rules.mjs`; these are the fix's test vectors). | elemental.md:6 against the fixture; the casters round-13 note | CONFIRMED-BUG | R13 S 11f258619 (missed the fixture) | S d0391912d |
| F2 | **The Fire guide fixture held the old Fire raid build.** `cdn.guides.mage-fire.json` raid_build was `251015/54521001230133031/003` (totals 14/34/3). The guide is `2530050000002/50520201230133031/003` (Arcane Meditation, 17/31/3). Same fix. | fire.md:7 | CONFIRMED-BUG | old (S 65fa2ffb1, 2026-11-01, before round 12) | S d0391912d |
| F3 | The test that would have caught both: `app-files.test.ts` "carries each guide's current build codes in its fixture". It reads every `cdn.guides.<spec>` fixture and compares its codes, build id aside, with the generator's. Red first (mage-fire raid_build), green after. The build id is dropped on purpose: a client move that keeps the trees leaves the vectors valid, and `tools/activate-build.sh` does not refresh these fixtures. | vitest `src/lib/app-data`, `src/lib/mobile`, `band-compare` (71 pass, real and fixture sync); Go `api/internal/app`, `internal/contract` pass; Android `PlannerGoldenTest`, `PlannerEditorTest`, Guides tests and `verifyRoborazziDebug` for Guides (31) and Planner (12) pass, so no screenshot moved | test | - | S d0391912d |
| F4 | Every other copy is current or a parse vector. Current: warrior-protection fixture, `apps/shared/links.json` example (tree current, build id 70707), iOS `PlannerCaptureTests`, Android `PlannerEditorTest`/`TalentTestData`, the 28 ladder goldens, `sim/request/lone_wolf_test.go` (Marksmanship's Lone Wolf raid build), `fs1.test.ts` Shadow. Parse vectors only, not copies: `sim/leveling/talents_test.go` (70009), `guide-builds.test.ts` (70374). Fury's Raging Blows, Affliction's Improved Bane of Agony and Balance's leveling build have no fixture copy. Protection paladin band 20 (lane tanks): no fixture copies a band-20 paladin build. | scan output | Aligned | - | - |
| F5 | The BiS fixtures `cdn.bis.warrior-fury` and `cdn.bis.hunter-beast-mastery` are 70658 snapshots of nightly BiS bands. Their band talents trail even the published 70753 files (11-01). They are floors by contract, and refreshing them copies nightly output, so they wait for the BiS refresh. | `docs/contracts/app-v1.md` "Fixtures" | Aligned (listed) | - | - |

## Item 2: regression over the landing

| # | finding | evidence and source | confirmed? | origin | fix sha |
| --- | --- | --- | --- | --- | --- |
| L1 | **`tools/land.sh` could land a failed run.** `wait_for_checks` read each run's jobs inside a command substitution: `$(for id ...; do retry run_jobs_json "$id" \| failed_lines; done)`. If `gh run view` failed three times, nothing was printed and the run counted as green. A completed, failed run whose view failed would have fast-forwarded main. Now the loop stops with "could not read run <id>". | Test first: `tools/tests/test_land.sh` gets a case (a failed run with no readable view); 3 assertions failed (it landed), then 30 of 30 pass; `test_go_test_shard.sh` 9 pass; shellcheck clean | CONFIRMED-BUG | landing S 2030bef5e | S a3bdf59e6 |
| L2 | The CI changes in 2030bef5e and 3046f2a76, read only: land/** triggers keep each workflow's path filters, land runs are grouped by sha and never cancelled, deploys and the companion release stay main-only, and the companion diffs a new branch against its merge-base. The 300-file path-filter limit is a real risk for big batches; the runbook states it. No check job is under 20 minutes; `lhci-alert` (5 min, main only, no checkout) is not a check. No workflow edited. | `.github/workflows/*.yml`, `docs/runbooks/landing.md` | Aligned | - | - |
| L3 | The dps-per-point weapon term, every consumer (3e3a14ff9). The only consumer is the addon: `Gear.dpsFigure`, `Gear.weaponTerm` (slot `ranged` for a ranged-attack-power spec, else `main_hand`; weight = weaponDps / dps) and Progress. The web (BiS page, item pages, home upgrades) never reads it; it reads published scores, and a weapon slot is only ever priced off a sim-verified pick. The API does not read it either. Data.lua's 150 `dpsPerPoint` rows equal `dps-per-point.json` to 4 decimals. A non-zero cat term (lane hybrids deletes `formIgnoresWeapon`, S ffdb90ce4) needs no consumer change: the cat's reference is attack_power, so the term lands on the main hand. One stale doc line: `Gear.lua:500` still lists "feral druids" among the specs with no term (passed to hybrids). The ladder's `specWeaponDPSMatters` already credits a cat's weapon DPS. | compare script; addon busted `item_tooltip_spec`, `rotation_spec`, `gear_spec`, `progress_spec`, `tooltip_bis_spec`, `export_spec` (410 pass) | Aligned | - | - |
| L4 | The coach starters (72cecbabb): `tools/coach/test_persona_starter.py` passed (53) before this lane's changes; `data/tests/test_coachrules.py` drift test passed. | pytest | Aligned | - | - |
| L5 | The Android planner captures (72cecbabb): `verifyRoborazziDebug` for PlannerScreenshotTest (12) passes, before and after F1. | gradle, JDK 21 | Aligned | - | - |
| L6 | **Same request, three paths.** spec-breakdown's own `legalRequest` (raid 60, horde, seed 1, 10,000 iterations) for shaman-elemental, warrior-fury, warlock-affliction and druid-balance. Each request ran through spec-breakdown (inproc), forever-sim (the server lane) and the wasm under node through simSplit, simRun and simCombine as the worker runs them. Elemental 480.2776, Fury 758.2387 and Balance 574.9437 agree on all paths and with `before.tsv`. **Affliction did not**: 571.9998 concurrent, 572.0266 single-threaded, 572.0236 with 4 wasm shards, 572.0186 with 8. See D1. Talent legality: the API's `talentsLegal` and spec-breakdown's `legalRequest` call the same `sim/talents` Check; the page and wasm send only what `orderFromRanks` reaches. | `$R/platform/req/`, `run.mjs`; `go test ./internal/sims -run 'Legal\|Talent'`; `go vet ./...` on api | Aligned except D1 | - | - |
| D1 | **A harmful-spell proc's recovery outlived the iteration, so a run's mean depended on how its iterations were split.** Comparing the concurrent and single-threaded paths for all 28 specs (3,000 iterations) found real gaps only in mage-arcane (619.234 / 619.217), warlock-affliction and warlock-demonology. Every other spec differs only in float summation (1e-12). Taking the gear apart left one item that alone causes it: Kindling Stave (11750), whose Ignition is a `MakeHarmfulSpellProcAura`. Its `lastProc` map (`sim/core/harmful_spell_proc.go`) was never cleared. A new iteration's cast landing inside the 10 ms window the last iteration's proc had left open was held by it. So the /sim page (wasm shards), forever-sim's sample path and the published concurrent number each read a little differently on the same seed. The same helper serves Necropile Drain and the client damage procs. Fix: the aura's OnReset clears the map. After the fix, all three paths give 572.0181 for Affliction (1, 4 and 8 shards), and the three specs give the same number concurrent and single-threaded. | Test first: `TestHarmfulSpellProcRecoveryEndsWithTheIteration` failed (a cast 5 ms into a new iteration was held), then passed. `go test --tags=with_db ./sim/core ./sim/common/item_effects ./sim/common/item_sets ./sim/mage/... ./sim/warlock/...` pass, so no golden moves. The same dependence is there on fork 076a5a517 and 2606bc05e. | CONFIRMED-BUG | old (E 3026483e6, 2026-11-01) | E b08e7cf68 |

## Item 3: what the sim pages and the addon export show

| # | finding | evidence and source | confirmed? | origin | fix sha |
| --- | --- | --- | --- | --- | --- |
| S1 | Rank rows (raceSpellId). An in-memory `build_rotations` equals the committed `addon-data.json` for every spec, band, line, spell id and race key. Data.lua carries the same 6 race-keyed lines. | compare script; busted `rotation_spec` | Aligned | - | - |
| S2 | The weights panel. top per-point stat = 1.00 (`buildScaleRows` / `perPointDivisor`), haste and expertise per 1% outside the normalization (`PER_PERCENT_STATS`: melee_haste, spell_haste, expertise; no other haste stat is in any published band's weights), a weight inside its error reads "Not significant". | `scale-view.ts`, `weight-reading.ts`; vitest (175 pass) | Aligned | - | - |
| S3 | The rotation drawer shows the curated notes in priority order (`rotations.ts`, synced from data/curated/apl). It names no ranks, so the rank rewrite cannot disagree with it. The note prose belongs to lane guides. | read | Aligned | - | - |
| S4 | Talent strips: `CharacterStrip` and `character.ts` tests pass under the CI fixture sync. The strip spends points through `orderFromRanks`, the same order the planner and the apps use (round 13 P1). | vitest | Aligned | - | - |
| S5 | The addon export's band 60 skips every rank-choice row (`addonrotation.py`: `if level < MAX_LEVEL else {}`), so it shows the raid-60 ranks. But the addon's BiS (`addonbis.py`) and its dps-per-point at 60 are the bare band's. 18 bare-60 rows exist (for example Shadow Vampiric Embrace rank 2 by race, Elemental's three). It is the same character as the site's default raid preset at 60, so no number is wrong today. The question is which preset the addon's level-60 player stands for. | `data/curated/rank-choices.json`; `addonrotation.py:_band_ranks` docstring | OPEN (owner: is the addon's 60 the raid or the bare character?) | - | - |

## Coach (passed on by other lanes)

| # | finding | evidence and source | confirmed? | origin | fix sha |
| --- | --- | --- | --- | --- | --- |
| C1 | **The coach named Ahn'Qiraj book ranks.** Found by lane guides. When the catalogue generator picks the rank to stamp a rule with, only the downranks skipped the book ranks (`curated/book-ranks.json`). An aura's stamped rank is the one an uptime reading names (`f2_uptime.go` `uptimeFacts`: Spell rank). So hunters were told "Aspect of the Hawk (Rank 7)", a book, when a level-60 trainer teaches rank 6 (14322; 110 attack power after the 11-03 push). The same went for Serpent Sting 25295 (rank 9; trainer top 13555), Corruption 25311, Immolate 25309 and Battle Shout 25289. Every list now reads the context's books. | Test first: `test_no_rule_names_an_ahnqiraj_book_rank_a_trainer_rank_could_name` failed (hunter-beast-mastery auras Serpent Sting 25295), then passed. trainables/hunter.json ranks 1-6 at 10-58; book-ranks.json 25296 rank 7. | CONFIRMED-BUG | old (S c8e762ec1 added books for downranks only) | S f7c1c5773 |
| C2 | **The Shadow coach starter contradicted the guide.** Found by lane regcast; their patch, reviewed against `shadow.md:73-93` and `data/curated/apl/priest-shadow.json` (Inner Focus into Mind Blast, Mind Blast, Shadow Word: Death, Mind Flay, wand). It said Inner Focus into Devouring Plague, Mind Blast only when Mind Flay cannot be cast, and Shadow Word: Death only below 20%. Now: Inner Focus on Mind Blast; Mind Blast and Shadow Word: Death on cooldown, ahead of Mind Flay; Mind Flay with everything else, never clipped. I also fixed the GUIDE_ORDER comment, which still quoted the old guide. Also fixed: `GUIDE_LINES` named "priest-discipline" twice, so the first entry was dead. I removed it rather than merging it: it was vaguer than the rotation's own line ("under ninety percent yourself"), so Discipline's starter is unchanged. | Tests first: `test_shadow_priority_follows_the_launch_guide` and the updated `test_guide_order_statements_win` failed, then passed. `test_discipline_names_its_overrides_once` failed on the duplicate, then passed. In a scratch run only `priest-shadow.md` differs. | CONFIRMED-BUG | old (dcd34cf96, per regcast) | S 46c1579fb |

## Before/after

Only D1 moves a number. Raid 60, 10,000 iterations, before = `$R/before.tsv`, after = lane engine b08e7cf68:

| spec | faction | seed 1 | seed 50001 |
| --- | --- | --- | --- |
| mage-arcane | horde | 619.2 → 619.20 (0.00%) | 619.5 → 619.43 (-0.01%) |
| mage-arcane | alliance | 614.2 → 614.17 (-0.01%) | 613.8 → 613.82 (+0.00%) |
| warlock-affliction | horde | 572.0 → 572.02 (+0.00%) | 571.8 → 571.90 (+0.02%) |
| warlock-affliction | alliance | 569.9 → 569.91 (+0.00%) | 570.3 → 570.17 (-0.02%) |
| warlock-demonology | horde | 639.3 → 639.19 (-0.02%) | 639.9 → 639.92 (+0.00%) |
| warlock-demonology | alliance | 635.4 → 635.29 (-0.02%) | 635.1 → 635.27 (+0.03%) |

All are inside error (about 0.29 DPS). Any spec whose published gear carries a harmful-spell proc moves by about this much; the three above are the ones found at raid 60. What matters is that the page, the server and the published number now agree exactly.

## Files to regenerate

- `data/curated/coach/hunter-beast-mastery.json`, `hunter-marksmanship.json`, `warlock-affliction.json`, `warlock-demonology.json`, `warlock-destruction.json`, `warrior-arms.json`, `warrior-fury.json` (`coachrules.write_rule_files`, C1). Until then `test_committed_files_match_the_active_build` fails.
- `data/curated/coach/priest-shadow.md` (`python tools/coach/persona_starter.py --all`, C2). Until then `test_generator_is_deterministic_and_committed_files_match` fails.
- `docs/contracts/fixtures/app/cdn.bis.warrior-fury.json`, `cdn.bis.hunter-beast-mastery.json` after the BiS refresh (F5).
- Regenerated with the fix as test vectors: the three `cdn.guides.*` build objects and `apps/shared/fixtures/planner-rules.json` (F1, F2).

## Measure at refresh

- None from this lane.

## For other lanes

- **hybrids:** `addon/ForeverSixty/Gear.lua:499-500` still says the weapon term is nil for "feral druids"; drop that from the comment with your formIgnoresWeapon removal. No consumer needs a change for a non-zero cat term (L3).
- **Integrator:** D1 is in `sim/core/harmful_spell_proc.go` (the coordinator said take it here). No fork golden moves. Some tests fail on a plain checkout and pass with the db: `sim/common/item_effects` needs `--tags=with_db`.
- **Integrator:** a client move that changes the trees re-encodes the guides. After that, refresh the `cdn.guides.*` build objects and run gen-planner-rules, or F3 fails. `activate-build.sh` regenerates only planner-rules.

## Still open

- S5: which preset the addon's level-60 rotation stands for (owner).
- iOS: `GuidesSnapshotTests` and `PlannerGoldenTests` were not run (xcodebuild). The fixture prose is unchanged, the guide totals are unchanged for Elemental and Protection, and Android's identical screenshots did not move. CI's `ios.yml` covers it.
- A Gradle daemon (pid 48077, started 14:04, before this lane) is still running. It is not this lane's; I left it.

## Tests run

Web vitest: `src/lib/app-data`, `src/lib/mobile`, `src/lib/planner/band-compare`, `src/lib/bis/{scale-view,weight-reading,panel-view}`, `src/lib/sim/{character,weights}`, `CharacterStrip`. Go: `api/internal/app`, `api/internal/contract`, `api/internal/sims` (legality), `go vet ./...` on api; fork `sim/core`, `sim/common/item_effects`, `sim/common/item_sets`, `sim/mage/...`, `sim/warlock/...` (with_db). Android: `PlannerGoldenTest`, `PlannerEditorTest`, Guides tests, `verifyRoborazziDebug` (Guides, Planner). Addon busted: 6 specs (410). Python: `data/tests/test_coachrules.py`, `tools/coach/test_persona_starter.py`. Shell: `test_land.sh`, `test_go_test_shard.sh`.

## wow-player review

- F1/F2: in the app, the planner tests now replay the Elemental leveling build with Tidal Focus, and the app fixture's Fire raid build has Arcane Meditation, as in the guides.
- L1: a landing can no longer go to main past a check whose result it could not read.
- D1: Kindling Stave's Ignition could be held by the last fight's recovery for 10 ms at the start of a new fight. The /sim page and the published number now agree exactly for an Arcane mage or a warlock wearing the staff.
- C1: the coach names Aspect of the Hawk rank 6, Serpent Sting rank 8, Corruption rank 6, Immolate rank 7 and Battle Shout rank 6, the top ranks a trainer teaches at 60, never the Ahn'Qiraj book ranks.
- C2: the Shadow coach says what the guide says: Inner Focus on Mind Blast, Mind Blast and Shadow Word: Death on cooldown, Mind Flay with everything else, never clipped.
