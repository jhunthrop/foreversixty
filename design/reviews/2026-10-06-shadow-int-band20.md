# Shadow Priest band 20 negative Intellect weight

Lane: sim-shadow-int. After the primary-stat anchoring fix, priest-shadow
band 20 (both factions) was the last bad weight anchor: Intellect
published at scale factor -2.60 (error 0.044, significant), so
`normalizeScaleFactors` fell back to Spell power as the anchor instead of
the spec's own primary stat.

## Reproduction

`data/builds/1.60.1.70009/bis/priest-shadow.json` band 20:

| stat | weight | error | insignificant |
|---|---|---|---|
| spell_power | 1 | 0.0015 | false |
| intellect | **-2.6047** | 0.0435 | false |
| shadow_power | 1 | 0.0015 | false |

Bands 30/40/50/60 (alliance) all publish Intellect positive and growing
with gear: 0.61, 1.00, 1.11, 1.22. The sign flip is unique to band 20.

## Root cause: the APL, not the engine or the sweep

`data/curated/apl/priest-shadow.json`'s rotation, in priority order, is:
SW:Pain upkeep → SW:Death execute (unlearned below 32, dropped) → Mind
Blast execute fallback → Inner Focus+Devouring Plague (`strictSequence`)
→ Mind Blast on cooldown → Mind Flay filler → **Shoot (wand), gated to
`currentManaPercent < 20%`**.

Band 20's own published talent build (`...-443000000000000000`, 11
points) spends every point on the first three Shadow tier-1 talents and
has **not** taken Mind Flay (bit 8 of the Shadow tree segment is `0`).
`sim/priest/mind_flay.go`'s `registerMindFlay` is gated by
`priest.Talents.MindFlay` and returns immediately when it is false, so
Mind Flay is never registered at band 20 - not a rank problem, a missing
talent point. Confirmed directly: a bare level-20 ladder character's
sample cast log shows only `spell:5019` (Shoot), `spell:8102` (Mind
Blast r2), `spell:970` (SW:Pain r3) - no Mind Flay, no Devouring Plague
(Inner Focus's sequence-mate DP rank 1 is learned at exactly 20, but
Inner Focus itself never fired in the sample either; out of scope here).

With Mind Flay unavailable, the rotation's only mana-spending filler is
gone. Once Mind Blast is on its 8 s cooldown and SW:Pain is up, the
priority list has **nothing** to cast until mana drops under the wand's
20% gate - every such global stood idle instead of wanding. Confirmed by
running the real ladder character (bare, rank-correct, talents as
published) through the engine directly with Intellect bonus-stat `-20`
(matching the weights sweep's own `defaultStatMod*20` nudge,
`wowsims-forever` `sim/core/statweight.go` lines ~178-186) vs `0` vs
`+20`:

| Intellect mod | DPS | Shoot casts (one 180s sample) |
|---|---|---|
| -20 | 19.09 | 84 |
| 0 | 18.60 | 65 |
| +20 | 16.39 | 26 |

More Intellect enlarges the mana pool, which shrinks the time spent
under the 20% wand gate and so **replaces wand casts with idle time**,
not with real damage - a genuine, reproducible negative marginal effect
at this exact gear/talent combination, which is why it measured
significant rather than noisy. This is hypothesis (a): the rotation
has an action that effectively loses to wanding (idling loses to
wanding), confirmed by a direct cast-count/DPS sweep. Hypothesis (b)
(engine sign/scale bug) is ruled out - Intellect's mana and crit
conversions are ordinary and the effect is confined to one band/talent
combination, not a systemic engine error. Hypothesis (c) (the sweep's
own ±20 perturbation) is a contributing amplifier, not the cause: ±20 is
a large fraction of a level-20 character's own base Intellect, which is
exactly why it is big enough to visibly cross the idle/wand threshold,
but the underlying defect - a missing filler - is real at ±1 point too,
just far noisier to measure there.

Bands 30-60 anchor correctly because their own talent builds all carry
the Mind Flay point (bit 8 of tree3 = `1` at 30/40/50/60), so Mind Flay
is registered and genuinely fills the idle gaps; Intellect's marginal
value is positive and ordinary there.

## Fix

`data/curated/apl/priest-shadow.json`: removed the `currentManaPercent
< 20%` condition from the Shoot action, making it the unconditional,
unconditional-priority-last filler. Mind Flay and Mind Blast still sit
above it and win the priority race whenever either is actually
available, so a band where Mind Flay IS talented (30-60) is unaffected
in practice - wand now fires only in the same gaps it already won
before. Where Mind Flay is not talented (band 20), wand now fills every
otherwise-idle global instead of standing still.

Verified with the same direct sweep after the fix:

| Intellect mod | DPS | Shoot casts |
|---|---|---|
| -20 | 22.35 | 121 |
| 0 | 23.31 | 119 |
| +20 | 23.72 | 118 |

Sign is now correct and monotonic. Re-running the real ranker
(`leveling-bis -spec priest-shadow -bands 20`) publishes Intellect at
**+0.302** (error 0.015, significant) at band 20, consistent in sign and
magnitude trend with bands 30-60 - `primaryAnchorStat` now anchors on
Intellect directly, no fallback needed.

The rotation ladder's golden (`sim/request/testdata/ladder/priest-
shadow.golden.md`) improved DPS at every rung (10 through 60), not only
20, since the old mana gate was leaving idle globals at every level
whenever Mind Blast was on cooldown and mana was above 20% - the fix is
a strict, level-independent improvement, not a band-20-only patch.
Regenerated with `FOREVER_UPDATE_GOLDEN=1`; only priest-shadow's golden
changed.

Also re-synced `sim/request/apl/priest-shadow.apl.json` (the embedded
copy `go test` and the real engine actually read) from the curated file.
`make apl-sync` found the fork's and this repo's embedded copies already
stale for ~18 other specs (unrelated notes-text drift); those were
reverted in both this repo and `/Users/jh/code/wowsims-forever` so only
priest-shadow's own change ships here. **Follow-up for the owner**: the
fork's `ui/shadow_priest/apls/forever_shadow.apl.json` mirror still needs
`make apl-sync`/`apl-check` run and committed there separately - not
done in this lane, since it touches a different repository outside this
worktree's scope.

## Tests

`cd sim && go vet ./... && go test -p 1 ./cmd/leveling-bis/... ./request/...`
- both green. `TestRotationLadder` required `FOREVER_UPDATE_GOLDEN=1`
once to accept the (improved) priest-shadow golden; no other spec's
golden changed. Pre-existing, unrelated ladder violations (hunter
Arcane Shot zero-casts at 50, mage-frost and warrior-fury dps_regression
at 40) are untouched by this change and were already present before it.
