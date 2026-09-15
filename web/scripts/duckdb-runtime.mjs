// web/scripts/duckdb-runtime.mjs
// Where the two DuckDB engine modules live, for the three build-time places that have to
// agree about it.
//
// They are the one part of the runtime that is not a static asset. Cloudflare Workers cap
// a single static asset at 25 MiB on every plan, and duckdb-eh.wasm is 32.7 MiB and
// duckdb-mvp.wasm 37.5 MiB, so publishing them under public/ makes `wrangler deploy` --
// the whole site's deploy, not just the Queries view -- fail. They go into the LOGS bucket
// instead, under runtime/duckdb/<version>/, and src/worker.ts serves them back at
// /duckdb-runtime/<version>/<file>: same origin, same "no third-party bytes" rule, no size
// limit. The worker JS files and the vendored extension are well under the cap and stay
// static assets.
//
// sync-duckdb.mjs stages them, upload-duckdb-runtime.mjs uploads them, and
// astro.config.mjs serves them straight off the staging directory in `astro dev`.
// src/lib/report/duckdb-runtime.ts is the browser and Worker half of the same contract;
// its own test holds the version in the two halves together.
import { readFileSync } from 'node:fs';
import { createRequire } from 'node:module';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

const require = createRequire(import.meta.url);

export const webRoot = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '..');

/** The directory the pinned @duckdb/duckdb-wasm ships its browser build in. */
export const packageDist = path.dirname(require.resolve('@duckdb/duckdb-wasm/dist/duckdb-browser.mjs'));

/**
 * The installed package's own version, not the range in package.json: this is the version
 * the staged bytes actually are, and it is what both the bucket key and the URL carry, so
 * a bump publishes to a new path rather than overwriting an immutable one. The package
 * does not export ./package.json, so it is read beside the dist directory.
 */
export const DUCKDB_WASM_VERSION = JSON.parse(
  readFileSync(path.join(packageDist, '..', 'package.json'), 'utf8'),
).version;

/** Over the static-asset cap: bucket, then Worker. */
export const RUNTIME_MODULES = ['duckdb-eh.wasm', 'duckdb-mvp.wasm'];

/** Under it: ordinary static assets, version-pinned so they can be cached immutably. */
export const RUNTIME_WORKERS = ['duckdb-browser-eh.worker.js', 'duckdb-browser-mvp.worker.js'];

/** The bucket src/worker.ts reads through its LOGS binding; see wrangler.jsonc. */
export const R2_BUCKET = 'foreversixty-logs';

export const R2_PREFIX = `runtime/duckdb/${DUCKDB_WASM_VERSION}`;

/** Where sync-duckdb.mjs puts the modules: outside public/, so they never reach dist/. */
export const stagingDir = path.join(webRoot, 'build/duckdb-runtime', DUCKDB_WASM_VERSION);

/** What `npm run upload:duckdb` runs, and what the README prints for a manual upload. */
export function uploadCommand(file) {
  return [
    'npx wrangler r2 object put',
    `${R2_BUCKET}/${R2_PREFIX}/${file}`,
    `--file build/duckdb-runtime/${DUCKDB_WASM_VERSION}/${file}`,
    '--content-type application/wasm --remote',
  ].join(' ');
}
