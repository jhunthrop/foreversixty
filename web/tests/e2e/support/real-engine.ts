// web/tests/e2e/support/real-engine.ts
// The real-wasm gate every *-real-engine.spec.ts file needs before it can run: instantiates
// the published artifact in an ephemeral page and reports which exports the running Go
// program actually registered on `window`, the same handshake engine.ts's own
// loadWasmEngine performs (Go, instantiateStreaming, race against wasmready).
//
// Duplicated out of sim-gear-real-engine.spec.ts (its own comment explains why this cannot
// import a browser-only module into a Node-side gate check) rather than left a second time
// inline, once a second real-engine spec (sim-weights-real-engine.spec.ts) needed the exact
// same probe.
import { existsSync } from 'node:fs';
import path from 'node:path';
import type { Browser } from '@playwright/test';
import { engineAssetUrl } from '../../../src/lib/sim/version';

export function artifactPublished(webRoot: string, version: string): boolean {
  return existsSync(path.join(webRoot, 'public/_sim', version, 'sim.wasm'));
}

export function realEngineSkipReason(what: string): string {
  return (
    `real-engine ${what}; publish web/public/_sim/<ENGINE_VERSION>/ first ` +
    '(make simdb && make artifacts && make publish-wasm) and run with npm run test:e2e:real-engine'
  );
}

/**
 * Loads `sim.js`/`sim.wasm` in a scratch page (closed before returning, so a negative
 * result leaves nothing running behind it) and answers whether the running program
 * registered every export in `need` as a function on `window`.
 */
export async function hasWasmExports(
  browser: Browser,
  baseURL: string,
  version: string,
  need: readonly string[],
): Promise<boolean> {
  const page = await browser.newPage();
  try {
    await page.goto(baseURL);
    const glueUrl = engineAssetUrl('sim.js', version);
    const wasmUrl = engineAssetUrl('sim.wasm', version);
    await page.addScriptTag({ url: glueUrl });
    return await page.evaluate(
      async ({ wasm, exportNames }) => {
        type GoGlue = new () => {
          importObject: WebAssembly.Imports;
          run(instance: WebAssembly.Instance): Promise<void>;
        };
        const w = window as unknown as Record<string, unknown> & { Go?: GoGlue; wasmready?: () => void };
        try {
          if (typeof w.Go !== 'function') return false;
          const go = new w.Go();
          const { instance } = await WebAssembly.instantiateStreaming(fetch(wasm), go.importObject);
          const ready = new Promise<'ready'>((resolve) => {
            w.wasmready = () => resolve('ready');
          });
          const exited: Promise<'exited'> = go.run(instance).then(
            () => 'exited',
            () => 'exited',
          );
          const outcome = await Promise.race([ready, exited]);
          if (outcome !== 'ready') return false;
          return exportNames.every((name) => typeof w[name] === 'function');
        } catch {
          return false;
        }
      },
      { wasm: wasmUrl, exportNames: [...need] },
    );
  } catch {
    return false;
  } finally {
    await page.close();
  }
}
