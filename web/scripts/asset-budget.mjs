// web/scripts/asset-budget.mjs
// The per-file size budget over dist/, run by check-island-size.mjs in `npm run postbuild`.
import { readdir, stat } from 'node:fs/promises';
import path from 'node:path';

/**
 * Cloudflare Workers static assets refuse a single file over 25 MiB, on every plan, and
 * the refusal is the whole deploy rather than that one file. Nothing else here catches it:
 * vitest, Playwright and Lighthouse all serve dist/ off the disk, where a 34 MB file is
 * merely a 34 MB file. This is the class of defect that put DuckDB's two engine modules --
 * 32.7 and 37.5 MiB -- into dist/ (they are served from R2 now; see scripts/duckdb-runtime.mjs)
 * and that left the raw sim.wasm at 24.7 MB, 94% of the limit, on 2026-11-01 (it is
 * published gzipped now; see the root Makefile's publish-wasm).
 */
export const CLOUDFLARE_ASSET_LIMIT_BYTES = 25 * 1024 * 1024;

/**
 * The budget every dist/ file is held to: 20 MiB, so a file growing toward the limit fails
 * the build with a fifth of the limit still in hand, not the deploy at the limit itself.
 */
export const ASSET_BUDGET_BYTES = 20 * 1024 * 1024;

/** Every file under `dir`, depth first, as paths relative to it. */
async function* filesUnder(dir, prefix = '') {
  for (const entry of await readdir(dir, { withFileTypes: true })) {
    const relative = prefix === '' ? entry.name : `${prefix}/${entry.name}`;
    if (entry.isDirectory()) yield* filesUnder(path.join(dir, entry.name), relative);
    else yield relative;
  }
}

/**
 * The largest file under `root` and every file over `budgetBytes`, largest first.
 * @param {string} root
 * @param {number} budgetBytes
 * @returns {Promise<{ largest: { file: string, bytes: number }, overBudget: { file: string, bytes: number }[] }>}
 */
export async function measureAssets(root, budgetBytes) {
  if (!Number.isFinite(budgetBytes) || budgetBytes <= 0) {
    throw new Error(`asset budget must be a positive byte count, got ${budgetBytes}`);
  }
  let largest = { file: '', bytes: 0 };
  const overBudget = [];
  for await (const file of filesUnder(root)) {
    const { size } = await stat(path.join(root, file));
    if (size > largest.bytes) largest = { file, bytes: size };
    if (size > budgetBytes) overBudget.push({ file, bytes: size });
  }
  return { largest, overBudget: overBudget.toSorted((a, b) => b.bytes - a.bytes) };
}
