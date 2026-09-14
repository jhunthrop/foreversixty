// web/scripts/check-island-size.mjs
// The spec caps the planner island at 60 KB gzipped. CI runs this straight after the
// island build so a dependency that quietly doubles the bundle fails the pull request
// instead of the mobile Lighthouse budget three steps later.
import { readFile } from 'node:fs/promises';
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
if (failed) process.exit(1);
