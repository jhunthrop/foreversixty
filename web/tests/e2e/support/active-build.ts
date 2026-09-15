// web/tests/e2e/support/active-build.ts
// The build the site is currently published from. Specs that stub a shared build, or
// assert the version the planner sends to the API, have to name the build the island
// will actually fetch: naming a different one makes the island ask for a build whose
// talent ids it does not have, and the failure looks like a broken planner rather than
// a stale constant. Reading it here means flipping src/data/active-build.json is a
// one-line change and not a sweep through the suite.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

const WEB_ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../..');

export const ACTIVE_BUILD: string = (
  JSON.parse(readFileSync(path.join(WEB_ROOT, 'src/data/active-build.json'), 'utf8')) as {
    build: string;
  }
).build;
