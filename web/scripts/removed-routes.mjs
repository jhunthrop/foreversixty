// web/scripts/removed-routes.mjs
// The check behind check-no-removed-routes.mjs, as a pure function of a directory so
// src/lib/site/removed-routes.test.ts can drive it (a plain .mjs: the script runs under node
// before vitest ever starts, the same reason rotation-apl.mjs is one).
//
// L8.4 fix round 3 item 5: the script used to take its dist path from
// `new URL('../dist/', import.meta.url).pathname`. On Windows that is "/C:/Users/.../dist/",
// and path.join made it "\C:\Users\...\dist", which never exists: every removed route read as
// gone and dist/setup.html read as missing, so `npm run build` exited 1 on Windows alone.
// fileURLToPath gives the platform's own path on Windows and Linux alike.
import { existsSync, readdirSync } from 'node:fs';
import { join } from 'node:path';
import { fileURLToPath } from 'node:url';

/** Pages spec 2026-09-25 section 3.1 removes; Astro's `file` format builds /x to dist/x.html. */
export const MUST_NOT_EXIST = [
  'classes.html',
  'zones.html',
  'dungeons.html',
  'search.html',
  'skyborne.html',
  'everything-we-know.html',
  'editions.html',
  'addon.html',
];

/** Dynamic sections that build one file per entry under their own directory. */
export const MUST_NOT_EXIST_DIRS = ['zones', 'dungeons'];

/** The `dist/` directory beside the `scripts/` directory a script URL sits in. */
export function distDirOf(scriptUrl) {
  return fileURLToPath(new URL('../dist/', scriptUrl));
}

/** Every way `dist` still ships a removed route or lacks /setup; empty when it is right. */
export function removedRouteFailures(dist) {
  const failures = [];
  for (const file of MUST_NOT_EXIST) {
    if (existsSync(join(dist, file))) failures.push(`dist/${file} still exists`);
  }
  for (const dir of MUST_NOT_EXIST_DIRS) {
    const path = join(dist, dir);
    if (existsSync(path) && readdirSync(path).length > 0) {
      failures.push(`dist/${dir}/ still has files: ${readdirSync(path).join(', ')}`);
    }
  }
  if (!existsSync(join(dist, 'setup.html'))) {
    failures.push('dist/setup.html is missing -- /setup should replace /addon');
  }
  return failures;
}
