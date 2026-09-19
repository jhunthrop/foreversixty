// web/src/lib/sim/engine.ts
// The four functions our sim.wasm exports, and the two implementations of them: the
// checked-in fake, and the real one loaded from /_sim/<ENGINE_VERSION>/.
//
// All four take and return JSON strings, never bytes. The contract's rule is that no
// protobuf crosses a lane boundary: we build sim.wasm ourselves from the site's sim/ Go
// module, which imports the engine as a library, so sim/request and sim/adapter both run
// INSIDE the wasm. The browser hands it a SimRequest and gets back a SimResult whose
// summary.Summary was built by the same Go code the server lane runs. There is therefore no
// protobuf toolchain in web/, no request encoder, and no TypeScript copy of the adapter.
//
// The engine's own thirteen js.Global().Set entrypoints sit behind these four and are not
// ours to call; this interface does not name them, which is the enforcement.
//
// PUBLIC_SIM_ENGINE is read through import.meta.env so both the Astro build and the
// standalone island build inline it (vite.island.config.ts allow-lists the PUBLIC_ prefix).
// It defaults to 'fake' because sim.wasm does not exist yet; web.yml sets it to 'wasm' once
// CI builds an artifact for the pinned sha.
import { ENGINE_VERSION, engineAssetUrl } from './version';

export type EngineMode = 'fake' | 'wasm';

export const ENGINE_MODE: EngineMode =
  (import.meta.env.PUBLIC_SIM_ENGINE as EngineMode | undefined) ?? 'fake';

export type ProgressHandler = (callbackId: string, progressJSON: string) => void;

export interface EngineModule {
  simRun(requestJSON: string, callbackId: string): Promise<string>;
  /** Returns ONE JSON string encoding a SimRequest[] (main.go's simSplit), never a JS array. */
  simSplit(requestJSON: string, n: number): string;
  /** Takes ONE JSON string encoding a SimResult[] (main.go's simCombine); returns one SimResult JSON. */
  simCombine(resultsJSON: string): string;
  /** Returns `{"aborted": boolean}` JSON -- false means no run was registered under that id. */
  simAbort(callbackId: string): string;
  onProgress(handler: ProgressHandler): void;
  /**
   * Whether a target-error run has another step to do. The decision is the engine's, not
   * the page's: the page never computes a stopping rule over an estimate it did not pool.
   * Returns `{"needs_more": boolean}`, or the `{"error": …}` envelope.
   *
   * Contract 10.2. The Go original is `api.NeedsMoreIterations`.
   */
  simNeedsMore(resultJSON: string, requestJSON: string): string;
  /**
   * `api.SimRequest.Validate`, for the request drawer (contract 10.2). Returns
   * `{"ok": boolean, "errors": [{"field", "message"}]}`, or the `{"error": …}` envelope for
   * a string that is not a request at all.
   */
  simValidate(requestJSON: string): string;
  /**
   * How many combinations a bulk request expands to, without allocating the requests
   * (contract 10.2). Part B's live combination count and cap notice. A breach answers
   * `{"error":"cap_exceeded","cap":n,"combinations":n}`, which is an answer and not a
   * failure, so this one is NOT passed through `unwrapOrThrow`.
   */
  simCount(requestJSON: string): string;
}

export interface RequestValidationError {
  /** The request's own JSON path, e.g. "iterations" or "encounter.targets". */
  field: string;
  /** The engine's own sentence, shown verbatim. */
  message: string;
}

export interface RequestValidation {
  ok: boolean;
  errors: RequestValidationError[];
}

/** `simCount`'s two answers, both of them ordinary (contract 10.2). */
export type CountAnswer =
  { ok: true; combinations: number } | { ok: false; cap: number; combinations: number };

interface GoGlue {
  new (): { importObject: WebAssembly.Imports; run(instance: WebAssembly.Instance): Promise<void> };
}

type WasmGlobals = {
  Go?: GoGlue;
  simRun?: (requestJSON: string, callbackId: string) => Promise<string>;
  simSplit?: (requestJSON: string, n: number) => string;
  simCombine?: (resultsJSON: string) => string;
  simAbort?: (callbackId: string) => string;
  simNeedsMore?: (resultJSON: string, requestJSON: string) => string;
  simValidate?: (requestJSON: string) => string;
  simCount?: (requestJSON: string) => string;
  simProgress?: ProgressHandler;
  /** Called by sim/cmd/wasm/main.go the moment the exports above and simEngineVersion are
   * set -- see loadWasmEngine below for why the page must define this before go.run(). */
  wasmready?: () => void;
};

/**
 * simSplit, simCombine and simAbort all fail the same way: `{"error": "..."}` JSON instead
 * of their success shape (main.go's errorJSON). simRun's failures are a full SimResult JSON
 * with `.error` set instead (main.go's fail()), which the caller already reads as a normal
 * result, so this check does not apply there.
 */
export function unwrapOrThrow(json: string): string {
  const parsed: unknown = JSON.parse(json);
  const error =
    typeof parsed === 'object' && parsed !== null && !Array.isArray(parsed)
      ? (parsed as { error?: unknown }).error
      : undefined;
  if (typeof error === 'string') throw new Error(error);
  return json;
}

async function loadWasmEngine(version: string): Promise<EngineModule> {
  const glueUrl = engineAssetUrl('sim.js', version);
  await import(/* @vite-ignore */ glueUrl);
  const globals = globalThis as unknown as WasmGlobals;
  if (globals.Go === undefined) throw new Error(`sim.js at ${glueUrl} did not define Go`);
  const go = new globals.Go();
  const { instance } = await WebAssembly.instantiateStreaming(
    fetch(engineAssetUrl('sim.wasm', version)),
    go.importObject,
  );
  // main.go's main() calls `js.Global().Call("wasmready")` the instant the four exports and
  // simEngineVersion are set, and blocks forever after (`select{}`) -- it never returns, so
  // go.run()'s own promise cannot be the ready signal. Without this, main.go finds no
  // `wasmready` function, panics on the call, and the Go runtime exits: every export it had
  // just registered stops working, and every later call to it fails with "Go program has
  // already exited". This must be set before go.run() starts running the program.
  const ready = new Promise<void>((resolve) => {
    globals.wasmready = resolve;
  });
  // A program that returns or throws before signalling wasmready never will: race the two so a
  // panic during start-up surfaces as an error instead of a page that waits forever. The exit
  // outcome is folded into a value (never a rejection) because after wasmready fires the run
  // promise stays pending for the life of the page and a late exit must not go unhandled.
  const exited = go.run(instance).then(
    () => new Error('sim.wasm exited before signalling wasmready'),
    (cause: unknown) => new Error(`sim.wasm failed before signalling wasmready: ${String(cause)}`),
  );
  const outcome = await Promise.race([ready, exited]);
  if (outcome instanceof Error) throw outcome;
  if (typeof globals.simRun !== 'function') throw new Error('sim.wasm did not export simRun');
  return {
    simRun: (requestJSON, callbackId) => globals.simRun!(requestJSON, callbackId),
    simSplit: (requestJSON, n) => unwrapOrThrow(globals.simSplit!(requestJSON, n)),
    simCombine: (resultsJSON) => unwrapOrThrow(globals.simCombine!(resultsJSON)),
    simAbort: (callbackId) => unwrapOrThrow(globals.simAbort!(callbackId)),
    simNeedsMore: (resultJSON, requestJSON) => unwrapOrThrow(globals.simNeedsMore!(resultJSON, requestJSON)),
    simValidate: (requestJSON) => unwrapOrThrow(globals.simValidate!(requestJSON)),
    // No unwrapOrThrow: `cap_exceeded` carries two numbers the page renders.
    simCount: (requestJSON) => globals.simCount!(requestJSON),
    onProgress: (handler) => {
      globals.simProgress = handler;
    },
  };
}

export async function loadEngine(version: string = ENGINE_VERSION): Promise<EngineModule> {
  if (ENGINE_MODE === 'fake') {
    const { createFakeEngine } = await import('../../fixtures/sim/engine-fake');
    return createFakeEngine();
  }
  return loadWasmEngine(version);
}
