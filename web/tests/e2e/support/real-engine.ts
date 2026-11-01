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
import { existsSync, readFileSync } from 'node:fs';
import path from 'node:path';
import type { Browser } from '@playwright/test';
import { ENGINE_VERSION, engineAssetUrl, WASM_FILE } from '../../../src/lib/sim/version';

/**
 * The artifact directory the page was built against: `make publish-wasm` names it
 * `<engine sha>-<wasm digest>` and the build reads it from PUBLIC_SIM_ARTIFACT, so the gate
 * reads the same variable (a hand-copied `public/_sim/<engine sha>/` still works without it).
 */
export const E2E_ARTIFACT: string = process.env.PUBLIC_SIM_ARTIFACT ?? ENGINE_VERSION;

/**
 * A spec's published talent string at a band, in FS1's slash form, read from the active
 * build's BiS file: a hardcoded string goes illegal the next time a tree moves (round 6:
 * three real-engine specs failed "not a legal build" on strings written before 70707).
 */
export function publishedTalents(webRoot: string, build: string, spec: string, band = 60): string {
  const file = path.join(webRoot, '..', 'data', 'builds', build, 'bis', `${spec}.json`);
  const bis = JSON.parse(readFileSync(file, 'utf8')) as { bands: { band: number; talents: string }[] };
  const entry = bis.bands.find((row) => row.band === band);
  if (entry === undefined) throw new Error(`${file} publishes no band ${band}`);
  return entry.talents.replaceAll('-', '/');
}

export function artifactPublished(webRoot: string, version: string): boolean {
  return existsSync(path.join(webRoot, 'public/_sim', version, WASM_FILE));
}

export function realEngineSkipReason(what: string): string {
  return (
    `real-engine ${what}; publish the wasm first (make simdb && make publish-wasm), then run ` +
    'npm run test:e2e:real-engine with PUBLIC_SIM_ARTIFACT set to artifacts/ARTIFACT_ID'
  );
}

/**
 * Loads `sim.js`/`sim.wasm.gz` in a scratch page (closed before returning, so a negative
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
    const wasmUrl = engineAssetUrl(WASM_FILE, version);
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
          // The published module is gzipped (version.ts's WASM_FILE). engine.ts's wasmResponse
          // cannot be imported into this serialized callback, so the same sniff is inline: a
          // host that sends Content-Encoding: gzip (astro preview does) has fetch inflate it
          // already, and only a body that still starts with gzip's magic is inflated here.
          const fetched = await fetch(wasm);
          if (!fetched.ok) return false;
          let bytes = new Uint8Array(await fetched.arrayBuffer());
          if (bytes[0] === 0x1f && bytes[1] === 0x8b) {
            const inflated = new Blob([bytes]).stream().pipeThrough(new DecompressionStream('gzip'));
            bytes = new Uint8Array(await new Response(inflated).arrayBuffer());
          }
          const { instance } = await WebAssembly.instantiate(bytes, go.importObject);
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
