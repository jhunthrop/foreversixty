// web/scripts/check-sim-engine-artifact.mjs
// The last-mile guard on web/public/_sim/README.md's contract: when PUBLIC_SIM_ENGINE=wasm
// built this dist/, web/src/lib/sim/engine.ts's loadWasmEngine fetches
// /_sim/<ENGINE_VERSION>/sim.wasm and sim.js at runtime, on a visitor's machine, not at
// build time -- a stale or missing publish would ship silently and only fail in the
// browser. This runs in `npm run postbuild`, after `astro build` has copied public/ into
// dist/, so a pin bump (web/src/lib/sim/version.ts) that outran `make publish-wasm` (or a
// CI step reordered ahead of the publish) fails the build instead of the page.
//
// A build with PUBLIC_SIM_ENGINE unset or 'fake' (every fixture build: `verify`'s test job,
// Lighthouse, `npm run dev`) never reads these files, so this is a deliberate no-op there --
// see web/src/lib/sim/engine.ts's own ENGINE_MODE default.
import { existsSync, readFileSync } from 'node:fs';
import path from 'node:path';
import process from 'node:process';
import { fileURLToPath } from 'node:url';

const webRoot = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '..');

if (process.env.PUBLIC_SIM_ENGINE !== 'wasm') {
  console.log('check-sim-engine-artifact: PUBLIC_SIM_ENGINE is not "wasm"; skipping (fake-engine build)');
  process.exit(0);
}

// Parsed, not imported: this script runs on plain Node against version.ts's source text, the
// same way web/src/lib/sim/version.test.ts parses sim/enginever/version.go rather than
// giving the web build a Go dependency.
const versionSource = readFileSync(path.join(webRoot, 'src/lib/sim/version.ts'), 'utf8');
const match = /ENGINE_VERSION = '([0-9a-f]{7,12})'/.exec(versionSource);
if (match === null) {
  console.error('check-sim-engine-artifact: could not parse ENGINE_VERSION out of src/lib/sim/version.ts');
  process.exit(1);
}
const engineVersion = match[1];

const distSimDir = path.join(webRoot, 'dist/_sim', engineVersion);
const missing = ['sim.wasm', 'sim.js'].filter((file) => !existsSync(path.join(distSimDir, file)));
if (missing.length > 0) {
  console.error(
    `check-sim-engine-artifact: PUBLIC_SIM_ENGINE=wasm but dist/_sim/${engineVersion}/ is missing ` +
      `${missing.join(', ')}. A page built this way links to an engine that does not exist at the ` +
      'path it names. Run `make simdb && make artifacts && make publish-wasm` from the repository ' +
      'root before `npm run build`, and confirm the published sha matches ENGINE_VERSION.',
  );
  process.exit(1);
}
console.log(`check-sim-engine-artifact: dist/_sim/${engineVersion}/ carries sim.wasm and sim.js`);
