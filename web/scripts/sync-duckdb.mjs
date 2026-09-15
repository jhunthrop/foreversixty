// web/scripts/sync-duckdb.mjs
// Publishes everything DuckDB-WASM loads at runtime from foreversixty.gg, so the Queries
// view makes no third-party request. The package's own getJsDelivrBundles() helper exists
// and is deliberately not used.
//
// Three sets of files.
//
// The two engine modules, out of node_modules. These do NOT go into public/: at 32.7 and
// 37.5 MiB they are over Cloudflare's 25 MiB static-asset cap, so a dist/ carrying them
// fails `wrangler deploy` for the whole site. They are staged into
// build/duckdb-runtime/<version>/ instead, uploaded to the LOGS bucket once per version
// with `npm run upload:duckdb`, and served by src/worker.ts at /duckdb-runtime/<version>/.
// scripts/duckdb-runtime.mjs owns that path arithmetic.
//
// The two workers, also out of node_modules, published under public/duckdb/<version>/:
// well under the cap, and version-pinned so public/_headers can cache them for a year. The
// browser picks between the EH and the MVP pair (src/lib/report/query.ts), because Safari
// versions still in use lack WebAssembly exceptions.
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
import { createHash } from 'node:crypto';
import { copyFile, mkdir, readFile, readdir, rm, stat } from 'node:fs/promises';
import path from 'node:path';
import {
  DUCKDB_WASM_VERSION,
  RUNTIME_MODULES,
  RUNTIME_WORKERS,
  packageDist,
  stagingDir,
  uploadCommand,
  webRoot,
} from './duckdb-runtime.mjs';

/** The DuckDB the pinned @duckdb/duckdb-wasm carries; also the vendored directory name. */
export const DUCKDB_VERSION = 'v1.4.3';

/**
 * The vendored extensions, pinned. The runtime files come from node_modules and are
 * already pinned by package-lock.json's integrity hashes; these have no npm source, so
 * this is their lockfile. A swapped, truncated or half-downloaded file fails the build
 * here rather than the browser, and refreshing the extension means updating the hash in
 * the same commit as the bytes. `shasum -a 256 <file>` prints these.
 */
export const EXTENSION_SHA256 = {
  'wasm_eh/parquet.duckdb_extension.wasm': '22765c8f7dc741cda2b571a66ac7bb355295d7d69a6c37e5315b265672984f55',
  'wasm_mvp/parquet.duckdb_extension.wasm':
    '0785c6c95d003eff4faa7b3b4b660f02c9c92f6d68d135ddf330d42e3a650600',
};

const vendor = path.join(webRoot, 'vendor/duckdb-extensions', DUCKDB_VERSION);
/** Everything under public/duckdb/ is version-pinned, so all of it can be cached forever. */
const target = path.join(webRoot, 'public/duckdb', DUCKDB_WASM_VERSION);

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

/**
 * Anything under public/duckdb/ that is not this version's directory. The path is
 * version-pinned so old bytes are never overwritten, which means nothing would ever
 * remove them either: they would go on being published, and a layout this script has
 * since moved away from would go on being deployed with it.
 */
async function pruneOldVersions() {
  let entries;
  try {
    entries = await readdir(path.dirname(target));
  } catch {
    return 0;
  }
  const stale = entries.filter((entry) => entry !== DUCKDB_WASM_VERSION);
  for (const entry of stale) await rm(path.join(path.dirname(target), entry), { recursive: true });
  return stale.length;
}

let copied = 0;
for (const file of RUNTIME_WORKERS) {
  copied += await publish(path.join(packageDist, file), path.join(target, file));
}
for (const file of RUNTIME_MODULES) {
  copied += await publish(path.join(packageDist, file), path.join(stagingDir, file));
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
    const key = `${platform}/${name}`;
    const expected = EXTENSION_SHA256[key];
    if (expected === undefined) {
      throw new Error(`No SHA-256 pinned for vendored extension ${key} in scripts/sync-duckdb.mjs`);
    }
    const source = path.join(from, name);
    const actual = createHash('sha256')
      .update(await readFile(source))
      .digest('hex');
    if (actual !== expected) {
      throw new Error(`${key} is ${actual}, not the pinned ${expected}. Refusing to publish it.`);
    }
    extensions += 1;
    copied += await publish(source, path.join(target, 'extensions', DUCKDB_VERSION, platform, name));
  }
}
if (extensions === 0) throw new Error(`No .duckdb_extension.wasm vendored for ${DUCKDB_VERSION}`);

const pruned = await pruneOldVersions();

console.log(
  `duckdb published to public/duckdb/${DUCKDB_WASM_VERSION} (${RUNTIME_WORKERS.length} workers, ` +
    `${extensions} ${DUCKDB_VERSION} extensions, ${copied} copied, ${pruned} stale entries removed)`,
);
// Printed every run rather than only when the bytes change: the upload is idempotent, it
// is the one part of this script CI cannot do without an R2 token, and a deploy that skips
// it leaves the Queries view with nothing to instantiate.
console.log(
  `duckdb engine modules staged in build/duckdb-runtime/${DUCKDB_WASM_VERSION}, ` +
    'too large to be static assets. Upload them once per version with `npm run upload:duckdb`, or:',
);
for (const file of RUNTIME_MODULES) console.log(`  ${uploadCommand(file)}`);
