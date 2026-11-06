// web/scripts/check-chunk-css.mjs
// Run from `postbuild`: fails the build if a chunk preloads a stylesheet dist/ does not hold
// (the check and why it exists: chunk-css.mjs).
import { fileURLToPath } from 'node:url';
import { missingChunkCss } from './chunk-css.mjs';

const failures = missingChunkCss(fileURLToPath(new URL('../dist/', import.meta.url)));

if (failures.length > 0) {
  console.error(
    'check-chunk-css: a lazily imported island has scoped CSS that pages inline away; move it to a global stylesheet:\n' +
      failures.map((f) => `  - ${f}`).join('\n'),
  );
  process.exit(1);
}

console.log('check-chunk-css: every stylesheet a chunk preloads is in dist/.');
