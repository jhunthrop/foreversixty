# Ranker stat-weight anchor — SimulationCraft convention

Branch `ranker-weights-anchor`. Built in `.worktrees/ranker-weights-anchor` against
`sim/cmd/leveling-bis`.

## Before (today's committed `data/builds/*/bis/*.json`)

`checkWeightsAnchorCoverage` (new, `sim/cmd/leveling-bis/check_weights_test.go`) walked every
committed band and flagged one whose `scale_reference_stat` was not the spec's primary stat, or
whose primary-stat row published `insignificant` on an otherwise-trustworthy band. Run via
`go test ./cmd/leveling-bis/... -run TestCheckWeightsAnchorReportsTodaysCommittedFiles -v`:

```
88 total issue(s) across 12 spec(s):
  druid-balance: 4 band+faction row(s)
  druid-feral: 10 band+faction row(s)
  mage-arcane: 10 band+faction row(s)
  mage-fire: 8 band+faction row(s)
  mage-frost: 10 band+faction row(s)
  priest-shadow: 6 band+faction row(s)
  shaman-elemental: 4 band+faction row(s)
  warlock-affliction: 10 band+faction row(s)
  warlock-demonology: 12 band+faction row(s)
  warlock-destruction: 6 band+faction row(s)
  warrior-arms: 4 band+faction row(s)
  warrior-fury: 4 band+faction row(s)
```

Fury band 60's own `attack_power` anchor (cited in the brief) is one instance of this
`warrior-fury` count — Strength scored lower than a secondary stat and the old rule anchored on
the secondary instead.

## Design

1. **Primary-stat table** (`sim/cmd/leveling-bis/primary_stat.go`): `primaryStatBySpec`, a single
   `map[string]string` keyed by `data/curated/specs.json`'s `spec` id, canonical values
   `strength`/`agility`/`intellect` only. `primaryAnchorStat(spec)` turns that into the actual
   `weight_stats` row id to anchor on, falling back to `healing_power` then `spell_power` when a
   spec's own `weight_stats` carries no literal `intellect` row (not exercised by any spec in
   today's data — every Intellect-primary spec already lists `intellect` — but the fallback is
   unit-tested). `TestPrimaryStatCoversEverySpec` (`primary_stat_test.go`) walks the real
   `specs.json` and fails if any of the 27 specs is missing or maps to something other than the
   three canonical stats.

2. **Anchoring** (`normalizeScaleFactors`, `weights.go`): now takes `primaryStat` and anchors
   there first (ignoring a merely-larger secondary), refusing a haste id defensively; falls back
   to the pre-existing largest-significant-weight rule only when the primary row is absent, has a
   non-positive weight, or is itself a haste id. `main.go` passes `primaryAnchorStat(specInfo)`'s
   result through the whole per-band/faction pipeline.

3. **Significance**: `primaryStatSignificanceCheck` (`weights.go`) re-runs the sweep once at
   `primaryStatRetryIterationsFactor` (2x) when the primary row is insignificant, splicing only
   that row back into the result — independent of the pre-existing reference-stat
   retry/fallback (4x), which protects a different row (hunter's reference stat is
   `ranged_attack_power`; its primary is `agility`) and is skipped entirely once the whole band is
   already untrustworthy. `forceAnchorRowSignificant` then clears `Insignificant` on exactly the
   anchor row on an otherwise-trustworthy band — never on one `weights_reason` already marked
   untrustworthy — so "no primary stat is ever published as not significant." A still-insignificant
   retry sets the new band field `weights_low_confidence` (omitted/false in the ordinary case).

## Timing (item 4)

`hunter-marksmanship -bands 60`, weights sweep only (per-band `weights (Ns)` log line), three runs
each:

| `-weights-iterations` | seconds |
|---|---|
| 100 (today's default) | 0.9–1.0s |
| 200 | 1.0–1.1s |

The sweep is a small fraction of a band's own ~13s total (trinket-rank/effect-rank/verify don't
scale with this flag). Raised `bis.yml`'s nightly run to 200 via `BIS_ARGS` (doubling adds at most
a few seconds across the whole `-all` run against the 30-minute budget) — fewer bands should ever
need the per-band retry guards at all now.

## Flag

`weights_low_confidence` (`bandReport`, `report.go`) — band-level, `omitempty`; web shows nothing
new yet per the brief.

## Tests

`cd sim && go vet ./... && go test ./cmd/leveling-bis/...` → `ok`. New/changed: `primary_stat.go`
+ `primary_stat_test.go` (coverage + fallback), `weights.go` (`normalizeScaleFactors` signature,
`primaryStatSignificanceCheck`, `forceAnchorRowSignificant`) + unit tests in `weights_test.go`
(primary-over-larger-secondary, haste never anchors, insignificant-primary-retries-once,
retry-still-insignificant-sets-flag, force/never-force-significant), two `main_test.go`
integration tests exercising `runSpec` end to end with a fake engine, and
`check_weights_test.go`'s own non-failing "before" report.

## `bis.yml`

Changed: `BIS_ARGS: '-all -weights-iterations 200'` on the ranking step, with an updated comment
documenting the measurement above.

No `data/builds/*/bis/*.json`, `web/public/data`, `loot.json`, `manifest.json`, `Data.lua`, or
`simdb.bin` was committed — the nightly `bis.yml` regenerates them; trigger it after merge.
