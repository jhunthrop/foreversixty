// web/scripts/upload-duckdb-runtime.mjs
// Puts the two staged DuckDB engine modules into the LOGS bucket, where src/worker.ts
// serves them from. Once per @duckdb/duckdb-wasm version: the key carries the version, so
// an upload never overwrites bytes a browser has been told are immutable, and re-running
// it for a version already there is harmless.
//
// Needs an R2 write token (CLOUDFLARE_API_TOKEN plus CLOUDFLARE_ACCOUNT_ID, or a
// `wrangler login` session). That is user-owned setup -- see web/README.md -- so
// .github/workflows/web.yml skips this with a notice rather than failing the deploy when
// the token is not configured.
import { spawnSync } from 'node:child_process';
import { access } from 'node:fs/promises';
import path from 'node:path';
import process from 'node:process';
import {
  DUCKDB_WASM_VERSION,
  R2_BUCKET,
  R2_PREFIX,
  RUNTIME_MODULES,
  stagingDir,
  webRoot,
} from './duckdb-runtime.mjs';

for (const file of RUNTIME_MODULES) {
  const source = path.join(stagingDir, file);
  try {
    await access(source);
  } catch {
    throw new Error(`${source} is not staged. Run \`npm run sync:duckdb\` first.`);
  }

  const args = [
    'wrangler',
    'r2',
    'object',
    'put',
    `${R2_BUCKET}/${R2_PREFIX}/${file}`,
    '--file',
    source,
    // WebAssembly.instantiateStreaming refuses anything else; src/worker.ts sets the same
    // header on the way out, so this is belt and braces rather than the only guard.
    '--content-type',
    'application/wasm',
    // Without it wrangler writes into the local .wrangler/ simulator instead of the bucket.
    '--remote',
  ];
  console.log(`uploading ${file} to ${R2_BUCKET}/${R2_PREFIX}/`);
  const { status, error } = spawnSync('npx', args, { cwd: webRoot, stdio: 'inherit' });
  if (error !== undefined) throw error;
  if (status !== 0) process.exit(status ?? 1);
}

console.log(`duckdb ${DUCKDB_WASM_VERSION} runtime uploaded to ${R2_BUCKET}/${R2_PREFIX}/`);
