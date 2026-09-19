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
  simSplit(requestJSON: string, n: number): string[];
  simCombine(resultsJSON: string[]): string;
  simAbort(callbackId: string): void;
  onProgress(handler: ProgressHandler): void;
}

interface GoGlue {
  new (): { importObject: WebAssembly.Imports; run(instance: WebAssembly.Instance): Promise<void> };
}

type WasmGlobals = {
  Go?: GoGlue;
  simRun?: (requestJSON: string, callbackId: string) => Promise<string>;
  simSplit?: (requestJSON: string, n: number) => string[];
  simCombine?: (resultsJSON: string[]) => string;
  simAbort?: (callbackId: string) => void;
  simProgress?: ProgressHandler;
};

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
  void go.run(instance);
  if (typeof globals.simRun !== 'function') throw new Error('sim.wasm did not export simRun');
  return {
    simRun: (requestJSON, callbackId) => globals.simRun!(requestJSON, callbackId),
    simSplit: (requestJSON, n) => globals.simSplit!(requestJSON, n),
    simCombine: (resultsJSON) => globals.simCombine!(resultsJSON),
    simAbort: (callbackId) => globals.simAbort!(callbackId),
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
