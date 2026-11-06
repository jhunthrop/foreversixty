// web/scripts/check-no-owner-links.mjs
// Run from `postbuild`: fails the build if any file the site serves (dist/) links to the
// owner's code host account. The check lives in owner-links.mjs (tested by
// src/lib/site/owner-links.test.ts).
import { fileURLToPath } from 'node:url';
import { OWNER_LINK, ownerLinkFailures } from './owner-links.mjs';

const web = fileURLToPath(new URL('../', import.meta.url));
const failures = ownerLinkFailures([fileURLToPath(new URL('../dist/', import.meta.url))], web);

if (failures.length > 0) {
  console.error(
    `check-no-owner-links: ${OWNER_LINK} is in what the site serves:\n` +
      failures.map((f) => `  - ${f}`).join('\n'),
  );
  process.exit(1);
}

console.log(`check-no-owner-links: nothing in dist/ contains ${OWNER_LINK}.`);
