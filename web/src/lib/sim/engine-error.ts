// web/src/lib/sim/engine-error.ts
// The engine's own words, translated at exactly one shape: an unsupported-spec refusal.
//
// Task 3 (healer review MAJOR, lines 133-136): loading a healer or tank character used to
// let the player set up and start a run, which then printed sim/request's raw
// `combine: part 0 failed: request: … unsupported spec: "druid-restoration"` -- the only
// honest sentence on the page, and one nobody but this codebase can read. Task 3's real fix
// is upstream of this file: `spec-label.ts`'s `isSimulatedSpec` disables the run path for
// every spec the web already knows is not dps (RunControl.svelte, BulkRunBar.svelte), so
// the engine should never see one of those specs in a request again.
//
// This function is the narrower, defensive half: the case where the web thinks a spec is
// dps and the engine still refuses it (a spec `isSimulatedSpec` has not caught up to, or a
// build the engine has not yet added support for). `detail` is shown verbatim everywhere
// else -- RunControl's `sim-detail`, BulkRunBar's own detail line -- run.ts's own doc
// comment says an unknown buff or consumable id is worth showing exactly as the engine
// wrote it, because the id is the only thing that says what to change. An unsupported-spec
// refusal is the one shape that is not: it names the engine's own spec id, not a player
// fact, so this is the one translation this lane makes.
import { SimApiError } from './api';
import { simCopy } from './copy';
import { specLabel } from './spec-label';

/** sim/request's own wording (ErrUnknownSpec), wherever it lands inside a wrapped failure
 *  like `combine: part 0 failed: request: …`. The id is always double-quoted. */
const UNSUPPORTED_SPEC = /unsupported spec: "([^"]*)"/;

/** The adapter's own wrapper around a recovered panic, or the dump that follows it. */
const ENGINE_PANIC = /runtime error:|Stack Trace:|goroutine \d+ \[/;

/**
 * The engine's raw failure text, or the honest sentence when it names an unsupported spec.
 * Everything else -- an unknown buff id, a duplicate summary row, a cancel, a plain network
 * failure -- is returned exactly as it arrived: this is not a second place that decides
 * what counts as worth paraphrasing, only the one shape the engine's own words are unfit
 * for a player to read.
 */
export function humaniseEngineError(detail: string): string {
  // A Go panic is not the engine's words to a player, it is a goroutine dump with the build
  // machine's own file paths in it. It is the second shape that is never shown verbatim.
  if (ENGINE_PANIC.test(detail)) return simCopy.engineCrashed;
  const match = UNSUPPORTED_SPEC.exec(detail);
  return match === null ? detail : simCopy.engineUnsupportedSpec(specLabel(match[1]));
}

/**
 * Fix round 1: `store.svelte.ts`'s `runOnServer()` never went through `humaniseEngineError`
 * -- it read `SimApiError.message` straight onto `message`, so the server lane was the one
 * place the guarantee above did not hold. `SimApiError.message` is never the raw wasm text
 * today (api.ts's `asSimError` only ever hands it fixed copy), but this closes the seam
 * structurally rather than leaving it open for whenever the API starts proxying one
 * through. Lives here, not in store.svelte.ts (already at its own 800-line ceiling before
 * this fix), since it is one more caller of `humaniseEngineError`, not new logic.
 */
export function humaniseServerFailure(error: unknown, fallback: string): string {
  return error instanceof SimApiError ? humaniseEngineError(error.message) : fallback;
}
