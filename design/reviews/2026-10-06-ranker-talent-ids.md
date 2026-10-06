# Ranker talent ids — route every engine-facing talent string through enginetalents

Branch `ranker-talent-ids`. Built in `.worktrees/ranker-talent-ids` against `sim/`.

## What was wrong

The compiled engine (fork at `90f9325b0`) was last regenerated from client `1.60.1.69893`.
`sim/internal/enginetalents` (from the merged `talent-search` lane) already documents the
consequence: the engine's own `proto/paladin.proto` still carries two talents the active build
(`1.60.1.70009`) no longer has (Improved Holy Strike, Crusade), and its `proto/shaman.proto` has
Elemental Fury and Elemental Alacrity sitting in each other's compiled field. A talent string
written in the active build's own (tier, column) positional order is misread by the engine for
exactly these two classes — every point after the first shifted talent lands on the wrong field.

`sim/cmd/talent-search` already read every talent id-first through `enginetalents.Layout`
(`ForClass`/`Encode`/`FieldFor`). Two other call sites still built an `api.CharacterSpec.Talents`
string positionally and handed it to the engine unconverted:

- **`sim/cmd/leveling-bis`**: `bandCharacter`/`ladderCharacter` (`character.go`) are "the one
  constructor every plain-DPS site in this band's rank/verify passes... builds its
  `api.CharacterSpec` from" (that file's own doc) — but every one of those call sites
  (`rank.go`, `trinkets.go`, `sets.go`, `verify.go`, `faction_trinkets.go`) was forwarding
  `main.go`'s `talents := leveling.LadderTalentString(...)` straight through. Every BiS band ever
  published for `paladin-retribution` and `shaman-elemental` was ranked, verified and weighted
  against a character the engine itself misread.
- **`sim/request/ladder_test.go`** (`TestRotationLadder`, the rotation-accuracy harness): the same
  pattern — `talents` went straight into the `api.SimRequest` both `statsReq` and `simReq` were
  built from. The ladder's own cast/DPS telemetry for these two specs was measured against the
  same misread character.

`sim/request/request.go`'s production `BuildWith` was **not** touched: `api.CharacterSpec.Talents`
is documented (`api/envelope.go`) as "the engine's own talent string, positional against the
class's tree sizes" — `BuildWith` correctly trusts that contract; the two bugs above were call
sites inside `sim/` that *construct* a talent string from `data/builds/<build>/talents/<class>.json`
and a guide's digits themselves, then violate their own API's contract before it ever reaches
`BuildWith`.

A second, independent bug (already fixed on `main` by the `guide-codes-70009` lane before this
branch started, confirmed by the coordinator mid-lane) compounded this for the same two classes:
every guide's FS1 stamp said `1.60.1.69893` while its digits were always authored in
`1.60.1.70009`'s order, so `GuideTalentTargets` mapped those digits onto the wrong talent ids for
paladin and shaman specifically, before `LadderTalentString` ever ran. With that stamp corrected,
`GuideBuildTalents`/`LoadTalentTrees`/`GuideTalentTargets` now resolve the right ids — this lane
added `leveling.RequireGuideBuildMatchesActive` as a loud, explicit guard against that invariant
drifting again, called from `cmd/leveling-bis/main.go`'s `runSpec` right after `GuideBuildTalents`.

## Design

1. **`leveling.TalentRanksFromString`** (`sim/leveling/talents.go`, new): the positional decode
   every repositioning needs — `LadderTalentString`'s own inverse, by stable talent id.
   `sim/cmd/talent-search/build.go`'s `decodeActive` now delegates to it instead of carrying its
   own copy of the same loop (DRY).
2. **`enginetalents.Layout.Reposition`** (`sim/internal/enginetalents/enginetalents.go`, new):
   decode-then-`Encode` in one call — the one method every site holding an already-positional
   string (the ranker's per-band truncated build, the ladder's) needs before that string reaches
   the engine.
3. **`leveling.RequireGuideBuildMatchesActive`** (new): fails loudly if a guide's FS1 stamp ever
   drifts from the active build again, wired into `cmd/leveling-bis/main.go`'s `runSpec`.
4. **`cmd/leveling-bis`**: `runSpec` takes a new `talentLayoutResolver` dependency (same role as
   its existing `engineRunner`) and a `bandTalentStrings` helper returns both the **site** string
   (active-build-positional — unchanged, still what `talentPoints` and the published band
   `talents` JSON field read) and the **engine** string (repositioned). Every rank/verify/
   trinket/set call site in the band loop now passes `engineTalents`, not `talents`, into
   `bandCharacter`/`ladderCharacter`. `bandCharacter` itself is untouched — it stays a dumb
   constructor (its own existing tests pass it arbitrary strings and assert them back verbatim).
5. **`sim/request/ladder_test.go`**: resolves the same `enginetalents.Layout` once per spec and
   calls `Reposition` once per level; the character built for the engine uses `engineTalents`,
   the golden's own `Talents` column keeps reporting the site string unchanged.

## Call sites changed

- `sim/leveling/talents.go` — added `TalentRanksFromString`, `ErrGuideBuildMismatch`,
  `RequireGuideBuildMatchesActive`.
- `sim/internal/enginetalents/enginetalents.go` — added `Layout.Reposition`.
- `sim/cmd/talent-search/build.go` — `decodeActive` now delegates to
  `leveling.TalentRanksFromString` (no behavior change, DRY).
- `sim/cmd/leveling-bis/main.go` — added `talentLayout`/`talentLayoutResolver`/
  `resolveTalentLayout`/`bandTalentStrings`; `runSpec` gained a `resolveLayout` parameter and the
  guide-build guard; every `rankTrinketSlot`/`reconcileFactionTrinkets`/`rankSlotWithEffects`/
  `trySetCompletion`/`verifyBand`/`applySwaps`/`ladderCharacter` call in the band loop now passes
  `engineTalents`; `buildReport` and `talentPointsSpent` still read the site `talents` string.
- `sim/cmd/leveling-bis/testhelpers_test.go` — added `identityLayout`/`identityTalentLayout` (the
  test double every fixture-based `runSpec` test now passes).
- `sim/cmd/leveling-bis/main_test.go`, `swapnote_supersede_test.go` — every `runSpec(...)` call
  site now passes `identityTalentLayout`.
- `sim/request/ladder_test.go` — resolves `enginetalents.Layout` per spec, repositions per level,
  builds the engine-facing character from `engineTalents` instead of `talents`.
- `sim/request/testdata/ladder/paladin-retribution.golden.md`,
  `shaman-elemental.golden.md` — regenerated (`FOREVER_UPDATE_GOLDEN=1`, scoped to these two specs
  only); DPS/cast telemetry unchanged, only the reported site-layout `Talents` column differs,
  entirely attributable to the already-merged guide-stamp fix (see above) — this lane's own
  engine-layout fix is invisible in that column by design.

## Verification

`cd sim && go vet ./... && go test -p 1 ./...` — green, including `request.TestRotationLadder`
and every `cmd/leveling-bis` test.

New/updated tests:

- `sim/internal/enginetalents/reposition_test.go`:
  `TestRepositionPutsRankedTalentsOnTheEnginesOwnIndex` — table test, real guide + active-build
  data for `paladin-retribution` and `shaman-elemental`, decoded through the real compiled
  engine's `FillTalentsProto`. Confirms (by stable id, reading the engine's own field by name):
  - Paladin: `conviction` (unaffected position, still resolved by id), `instrument_of_law` and
    `twist_of_light` (both shifted one field by Crusade's compiled-but-removed slot), and
    `vengeance` (correctly left at 0).
  - Shaman: `elemental_fury` and `elemental_alacrity` land on their own engine fields, not each
    other's.
  `TestRepositionBypassGuardWouldHaveCaughtTheRealDrift` — fails if the site and engine strings
  for these two real classes were ever equal, i.e. if the known drift this fix depends on ever
  closes without anyone noticing (which would silently defang the guard above).
- `sim/leveling/talents_more_test.go`: unit tests for `TalentRanksFromString` (by-position decode,
  zero-dropping, and its three error paths) and `RequireGuideBuildMatchesActive` (accepts a match,
  rejects and wraps `ErrGuideBuildMismatch` on a drift).
- `sim/cmd/leveling-bis/band_talent_strings_test.go`:
  `TestBandTalentStringsRoutesPaladinAndShamanThroughTheEngineLayout` — real repo data, asserts
  `bandTalentStrings`' site and engine returns differ for these two real specs (the actual
  function `runSpec`'s band loop calls — if a future edit reverts it to returning the site string
  twice, this fails). `TestBandTalentStringsUsesTheGivenLayoutNotTheSiteOrder` — a synthetic,
  filesystem-free check that the "engine" return really comes from the given layout's
  `Reposition`, not from the site string computed alongside it.

## Confirmation: Ret Paladin band-60 engine-facing string

For the real `paladin-retribution` guide build at band 60, `bandTalentStrings` (and
`enginetalents.Layout.Reposition` under it) now puts:

- `conviction` = 3, `improved_judgement` = 2 (both already correctly placed — unaffected by the
  Crusade/Improved Holy Strike shift)
- `vindication` = 3, `sanctified_judgement` = 3, `seal_of_command` = 1, `pursuit_of_justice` = 2,
  `eye_for_an_eye` = 1, `sacred_arbiter` = 1 (all correctly placed)
- `instrument_of_law` = 2 and `twist_of_light` = 1 — both now on their own compiled engine field
  (`instrument_of_law`/`twist_of_light`, fields 51/52), one slot later than the active build's own
  position, exactly where Crusade's still-compiled-but-removed field pushes them
- `champion_of_the_light` and `vengeance` correctly left at 0

on the engine's own indices (`TestRepositionPutsRankedTalentsOnTheEnginesOwnIndex/retribution_paladin`),
rather than on whatever index those talents occupy in the active build's own positional order.
