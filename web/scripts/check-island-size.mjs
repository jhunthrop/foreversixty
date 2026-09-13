// web/scripts/check-island-size.mjs
// The spec caps the planner island at 60 KB gzipped. CI runs this straight after the
// island build so a dependency that quietly doubles the bundle fails the pull request
// instead of the mobile Lighthouse budget three steps later.
import { readFile } from 'node:fs/promises';
import path from 'node:path';
import process from 'node:process';
import { fileURLToPath } from 'node:url';
import { gzipSync } from 'node:zlib';

const LIMIT_BYTES = 60 * 1024;
const webRoot = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '..');
const bundle = path.join(webRoot, 'dist/planner-island.js');

const bytes = await readFile(bundle);
const gzipped = gzipSync(bytes, { level: 9 }).length;
console.log(`planner-island.js: ${bytes.length} bytes raw, ${gzipped} bytes gzipped`);

if (gzipped > LIMIT_BYTES) {
  console.error(
    `planner island is ${gzipped} bytes gzipped, over the ${LIMIT_BYTES} byte budget in ` +
      `docs/superpowers/specs/2026-09-13-phase-1-build-planner-design.md.`,
  );
  process.exit(1);
}
