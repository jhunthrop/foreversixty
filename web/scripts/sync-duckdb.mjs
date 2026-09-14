// web/scripts/sync-duckdb.mjs
// Publishes everything DuckDB-WASM loads at runtime under public/duckdb/, so the Queries
// view gets it from foreversixty.gg and no third-party request is ever made. The
// package's own getJsDelivrBundles() helper exists and is deliberately not used.
//
// Two sets of files.
//
// The runtime, out of node_modules: the exception-handling build and the MVP build, each
// a .wasm and a worker. The browser picks between them (src/lib/report/query.ts), because
// Safari versions still in use lack WebAssembly exceptions.
//
// The parquet extension, out of vendor/duckdb-extensions/. Since DuckDB 1.4 the Parquet
// reader is no longer compiled into the engine: the first read_parquet() in a session
// makes DuckDB fetch it from https://extensions.duckdb.org, which is exactly the
// third-party request this whole arrangement exists to prevent. npm publishes no package
// carrying it, so the signed bytes are vendored and served from our own repository path,
// which src/lib/report/query.ts points DuckDB at with custom_extension_repository. The
// layout under vendor/ mirrors the official repository exactly -- <version>/<platform>/
// <name>.duckdb_extension.wasm -- because DuckDB builds that path itself.
//
// To refresh after bumping @duckdb/duckdb-wasm, read the new DuckDB version out of the
// engine (`SELECT version()`; it is v1.4.3 for 1.32.0), then, from web/:
//   V=v1.4.3
//   for p in wasm_eh wasm_mvp; do
//     mkdir -p vendor/duckdb-extensions/$V/$p
//     curl -sSLo vendor/duckdb-extensions/$V/$p/parquet.duckdb_extension.wasm \
//       https://extensions.duckdb.org/$V/$p/parquet.duckdb_extension.wasm
//   done
// and update DUCKDB_VERSION below. The e2e (tests/e2e/report-queries.spec.ts) fails if
// the version and the vendored directory ever drift apart, because the query then reaches
// for a file this script never published.
import { copyFile, mkdir, readdir, stat } from 'node:fs/promises';
import { createRequire } from 'node:module';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

/** The DuckDB the pinned @duckdb/duckdb-wasm carries; also the vendored directory name. */
export const DUCKDB_VERSION = 'v1.4.3';

const RUNTIME_FILES = [
  'duckdb-eh.wasm',
  'duckdb-browser-eh.worker.js',
  'duckdb-mvp.wasm',
  'duckdb-browser-mvp.worker.js',
];

const require = createRequire(import.meta.url);
const webRoot = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '..');
const dist = path.dirname(require.resolve('@duckdb/duckdb-wasm/dist/duckdb-browser.mjs'));
const vendor = path.join(webRoot, 'vendor/duckdb-extensions', DUCKDB_VERSION);
const target = path.join(webRoot, 'public/duckdb');

/** Null rather than a throw: a file that is not there is the ordinary first-run case. */
async function sizeOf(file) {
  try {
    return (await stat(file)).size;
  } catch {
    return null;
  }
}

/**
 * Copies one file only when the destination differs in size. Every `npm run sync` runs
 * this, and the two engine builds are ~35 MB each: re-copying 80 MB on each predev,
 * pretest, precheck and prebuild is pure latency. Size is fingerprint enough, because the
 * source only changes when the pinned package version does.
 */
async function publish(from, to) {
  const [source, existing] = await Promise.all([sizeOf(from), sizeOf(to)]);
  if (source === null) throw new Error(`duckdb runtime file is missing: ${from}`);
  if (source === existing) return 0;
  await mkdir(path.dirname(to), { recursive: true });
  await copyFile(from, to);
  return 1;
}

let copied = 0;
for (const file of RUNTIME_FILES) {
  copied += await publish(path.join(dist, file), path.join(target, file));
}

let extensions = 0;
for (const platform of ['wasm_eh', 'wasm_mvp']) {
  const from = path.join(vendor, platform);
  let names;
  try {
    names = await readdir(from);
  } catch {
    throw new Error(
      `No vendored DuckDB extensions for ${DUCKDB_VERSION}/${platform}. ` +
        'See the refresh commands at the top of scripts/sync-duckdb.mjs.',
    );
  }
  for (const name of names.filter((entry) => entry.endsWith('.duckdb_extension.wasm'))) {
    extensions += 1;
    copied += await publish(
      path.join(from, name),
      path.join(target, 'extensions', DUCKDB_VERSION, platform, name),
    );
  }
}
if (extensions === 0) throw new Error(`No .duckdb_extension.wasm vendored for ${DUCKDB_VERSION}`);

console.log(
  `duckdb published to public/duckdb (${RUNTIME_FILES.length} runtime files, ` +
    `${extensions} ${DUCKDB_VERSION} extensions, ${copied} copied)`,
);
