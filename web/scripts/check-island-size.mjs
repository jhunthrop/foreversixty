// web/scripts/check-island-size.mjs
// The spec caps the planner island at 60 KB gzipped. CI runs this straight after the
// island build so a dependency that quietly doubles the bundle fails the pull request
// instead of the mobile Lighthouse budget three steps later.
import { readdir, readFile } from 'node:fs/promises';
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
];

// Rankings, Character and Guild are ordinary `client:load` Astro islands, not standalone
// Vite builds, so Astro emits them under dist/_astro/ with a content hash in the filename
// (e.g. Rankings.xBviRjNL.js) rather than the fixed name the two budgets above match on.
// None of their pages are in lighthouserc.json's collect.url, so nothing else in CI catches
// a regression here. 16 KB gzipped is roughly 3x the heaviest of the three today (Rankings,
// ~4.1-4.8 KB depending on gzip level) -- enough headroom that legitimate UI growth will not
// flap the build, tight enough that an accidental heavy import (a chart library, a duplicated
// data module) still trips it well before it could threaten a Lighthouse score no test here
// measures directly.
const PAGE_ISLAND_BUDGETS = [
  { component: 'Rankings', limitBytes: 16 * 1024 },
  { component: 'Character', limitBytes: 16 * 1024 },
  { component: 'Guild', limitBytes: 16 * 1024 },
];

const webRoot = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '..');

let failed = false;

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
