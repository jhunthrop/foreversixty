# Sim-accuracy lane brief

The standing brief for a sim-accuracy audit-and-fix lane. The 4-hourly sweep dispatches one lane per pick with this brief plus its scope. It lives here, not in a scratch folder, so it survives sessions ([[feedback-design-work-lives-in-repo]]).

## Why

The owner: "it's crucial we're accurate". A player called out the warlock sim. Rounds 1 to 29 are in `ledger.md`; round 29's open list is in `design/reviews/2026-11-07-sim-audit-round29/integration.md` ("Open, for round 30").

## Scope

The sweep names the specs. Audit only those, end to end:

1. **Engine against the client.** Read the fork's implementation of every spell, talent, aura, proc and set bonus the spec's rotation and builds touch (the fork is `/Users/jh/code/wowsims-forever`; the pinned sha is what `.github/actions/pin-engine` and `sim/go.mod` name). Check each number against the client build under `data/builds/<current>/` (spells, the hotfix push layer, gametables), then Blizzard's Forever notes. The client wins over Wowhead; the hotfix push wins over the client tables.
2. **Engine against real fights.** Compare the sim's cast counts, uptimes and damage shares with real Forever logs: our own uploaded reports in the API (the `reports`/logs tables and `api/internal/logs`) and public Warcraft Logs reports for the spec. A sim that disagrees with logs by more than its own error on a stable number is a bug until explained.
3. **Builds, rotations and rank rows** follow the talent adoption rule ([[forever-talent-adoption-rule]]): a change lands only with the gain beyond error, nothing the rotation casts dropped, and the tank death-chance rule respected.

## Every confirmed bug is fixed end to end

- A failing test first (fork unit test, or a targeted `sim/` fixture test: one spec, one band, low iterations).
- The fix in the fork on a branch, pushed; the site's engine pin moved through the engine-follow path (`make follow`, `engine.yml`, or a hand `engine: pin <sha>` commit with the goldens regenerated).
- Before/after DPS or HPS for every band the fix touches, as relative gains; prose never quotes nightly numbers ([[feedback-prose-never-quotes-nightly-numbers]]).
- Guide prose and card notes updated in the same landing, with a wow-player SHIP per class ([[feedback-sim-changes-ship-guide-text]]). Guide voice: no em dashes, no sim meta talk ([[feedback-guide-voice]]).
- No regens: no local or CI BiS rank, data run or simdb rebuild. Anything that needs a full re-rank goes on the ledger's "Measure at refresh" table.
- Gates before the push: lint, types, unit tests, the touched Playwright specs only, the guard tests. Rebase on origin/main, one push, no PR ([[feedback-push-to-main-no-prs]], [[feedback-personas-over-tests]]).

## Compute

Every sim run, Go test and golden regeneration runs on the Windows test box through `tools/remote-test.sh` (`--dir sim` or `--dir api`; engine pinning is on by default for `go` commands, see the script header). Never on the Mac: the laptop has hit load 70 under local sims. If the script cannot target an unpinned fork branch, add a small `--engine-sha <sha>` option to it and keep the change.

## Report

- Write `design/reviews/<date>-sim-accuracy/<scope>.md`: method, every finding with its source line and log evidence, Fixed/Held/Open, before/after, shas.
- Add the lane's row to `ledger.md` (date, scope, bugs fixed, origin, engine pin after).
- Tell the coordinator in three lines only when a bug was fixed or something needs the owner.

## Skip

Skip a spec if another lane already holds it (check `ListAgents` and the ledger's latest rows).
