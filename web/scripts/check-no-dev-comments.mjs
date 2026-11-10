// web/scripts/check-no-dev-comments.mjs
// Run from `postbuild`: fails the build if a built page still ships a developer comment or an
// internal path in its inline code or HTML (the `ship-minified` integration in
// astro.config.mjs strips them; the rule lives in inline-code.mjs, tested by
// src/lib/site/inline-code.test.ts).
import { fileURLToPath } from 'node:url';
import { distDevCommentFailures } from './inline-code.mjs';

const web = fileURLToPath(new URL('../', import.meta.url));
const failures = distDevCommentFailures(fileURLToPath(new URL('../dist/', import.meta.url)), web);

if (failures.length > 0) {
  console.error(
    `check-no-dev-comments: ${failures.length} developer comment(s) ship in dist/:\n` +
      failures.map((f) => `  - ${f}`).join('\n'),
  );
  process.exit(1);
}

console.log('check-no-dev-comments: no built page ships a developer comment.');
