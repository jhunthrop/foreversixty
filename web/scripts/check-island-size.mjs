// web/scripts/check-island-size.mjs
// Two size checks over dist/, both run by `npm run postbuild` so a pull request fails here
// rather than three steps later: the island budgets the spec sets, and the hard per-file
// limit Cloudflare enforces at deploy time.
import { readdir, readFile, stat } from 'node:fs/promises';
import path from 'node:path';
import process from 'node:process';
import { fileURLToPath } from 'node:url';
import { gzipSync } from 'node:zlib';

// Two standalone islands, two budgets. The planner's 60 KB comes from the Phase 1 spec.
// The report island is larger by design -- twelve tabs, a canvas chart and a filter bar --
// and 140 KB gzipped is the ceiling that keeps /reports/<id> inside its Lighthouse
// performance budget of 0.90 on a throttled phone. DuckDB-WASM is not counted: it is
// loaded lazily from separate files and never on page load.
const BUDGETS = [
  { file: 'dist/planner-island.js', limitBytes: 60 * 1024 },
  { file: 'dist/report-island.js', limitBytes: 140 * 1024 },
  // The sim island is the planner's gear grid plus the report's tables plus a run control.
  // 90 KB gzipped is roughly twice what those parts weigh today and well under the report's
  // ceiling; sim.wasm is not counted, since it is a separate immutable asset fetched after
  // first paint and never on the LCP path.
  { file: 'dist/sim-island.js', limitBytes: 90 * 1024 },
  // The four combination tools: a candidate grid, an item search, a source picker and two
  // results views, each a lazily-imported chunk off one entry. 70 KB gzipped for the entry
  // is roughly twice what the shared parts (the character strip, the source switcher, the
  // run bar) weigh; the per-tool chunks are not counted here because none of them is on any
  // page's LCP path -- the shell paints its own skeleton first.
  { file: 'dist/sim-tools-island.js', limitBytes: 70 * 1024 },
];

// Rankings, Character and GuildShell are ordinary `client:load` Astro islands, not
// standalone Vite builds, so Astro emits them under dist/_astro/ with a content hash in
// the filename (e.g. Rankings.xBviRjNL.js) rather than the fixed name the two budgets
// above match on. GuildShell is the client-side router guild.astro and
// guild/[...path].astro mount (Task 11): it statically imports Guild, GuildClaim,
// GuildSettings and GuildJoin, so its chunk carries all four views' weight, not just the
// one plain-guild page's -- still ~7 KB gzipped today, comfortably under budget. None of
// their pages are in lighthouserc.json's collect.url, so nothing else in CI catches a
// regression here. 16 KB gzipped is enough headroom that legitimate UI growth will not
// flap the build, tight enough that an accidental heavy import (a chart library, a
// duplicated data module) still trips it well before it could threaten a Lighthouse score
// no test here measures directly.
const PAGE_ISLAND_BUDGETS = [
  { component: 'Rankings', limitBytes: 16 * 1024 },
  { component: 'Character', limitBytes: 16 * 1024 },
  { component: 'GuildShell', limitBytes: 16 * 1024 },
];

/**
 * Cloudflare Workers static assets refuse a single file over 25 MiB, on every plan, and
 * the refusal is the whole deploy rather than that one file. Nothing else here catches it:
 * vitest, Playwright and Lighthouse all serve dist/ off the disk, where a 34 MB file is
 * merely a 34 MB file. This is what stops the class of defect that put DuckDB's two
 * engine modules -- 32.7 and 37.5 MiB -- into dist/ in the first place (they are served
 * from R2 now; see scripts/duckdb-runtime.mjs).
 */
const CLOUDFLARE_ASSET_LIMIT_BYTES = 25 * 1024 * 1024;

const webRoot = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '..');

let failed = false;

/** Every file under `dir`, depth first, as paths relative to dist/. */
async function* filesUnder(dir, prefix = '') {
  for (const entry of await readdir(dir, { withFileTypes: true })) {
    const relative = prefix === '' ? entry.name : `${prefix}/${entry.name}`;
    if (entry.isDirectory()) yield* filesUnder(path.join(dir, entry.name), relative);
    else yield relative;
  }
}

const distRoot = path.join(webRoot, 'dist');
let largest = { file: '', bytes: 0 };
for await (const file of filesUnder(distRoot)) {
  const { size } = await stat(path.join(distRoot, file));
  if (size > largest.bytes) largest = { file, bytes: size };
  if (size > CLOUDFLARE_ASSET_LIMIT_BYTES) {
    console.error(
      `dist/${file} is ${size} bytes, over Cloudflare's ${CLOUDFLARE_ASSET_LIMIT_BYTES} byte ` +
        'static-asset limit. `wrangler deploy` would refuse the whole site. Serve it from R2 ' +
        'through src/worker.ts instead, the way the DuckDB engine modules are.',
    );
    failed = true;
  }
}
console.log(`dist/: largest asset is dist/${largest.file} at ${largest.bytes} bytes`);

for (const budget of BUDGETS) {
  const bytes = await readFile(path.join(webRoot, budget.file));
  const gzipped = gzipSync(bytes, { level: 9 }).length;
  console.log(`${budget.file}: ${bytes.length} bytes raw, ${gzipped} bytes gzipped`);
  if (gzipped > budget.limitBytes) {
    console.error(`${budget.file} is ${gzipped} bytes gzipped, over the ${budget.limitBytes} byte budget.`);
    failed = true;
  }
}

const astroChunksDir = path.join(webRoot, 'dist/_astro');
const astroChunks = await readdir(astroChunksDir);
for (const budget of PAGE_ISLAND_BUDGETS) {
  const pattern = new RegExp(`^${budget.component}\\.[\\w-]+\\.js$`);
  const matches = astroChunks.filter((name) => pattern.test(name));
  if (matches.length !== 1) {
    console.error(
      `Expected exactly one dist/_astro/${budget.component}.<hash>.js chunk, found ${matches.length}` +
        (matches.length ? `: ${matches.join(', ')}` : '.'),
    );
    failed = true;
    continue;
  }
  const chunkPath = path.join(astroChunksDir, matches[0]);
  const bytes = await readFile(chunkPath);
  const gzipped = gzipSync(bytes, { level: 9 }).length;
  console.log(`dist/_astro/${matches[0]}: ${bytes.length} bytes raw, ${gzipped} bytes gzipped`);
  if (gzipped > budget.limitBytes) {
    console.error(
      `dist/_astro/${matches[0]} is ${gzipped} bytes gzipped, over the ${budget.limitBytes} byte budget.`,
    );
    failed = true;
  }
}

if (failed) process.exit(1);
