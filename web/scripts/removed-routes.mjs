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
  // /pricing is an edge 301 to /premium (public/_redirects), not a built page.
  'pricing.html',
  // /roadmap folded into the changelog (pages spec 2026-11-08 ruling Q7): an edge 301 to /changelog.
  'roadmap.html',
];

/** Dynamic sections that build one file per entry under their own directory. */
export const MUST_NOT_EXIST_DIRS = ['zones', 'dungeons'];

/** The BiS mid-level band pages folded into the spec page on 2026-11-15
 *  (src/lib/bis/band-path.ts): edge 301s to /bis/<class>/<spec>#levels-<slug> (public/_redirects),
 *  so dist/bis/<class>/<spec>/<slug>.html must never ship. */
export const FOLDED_BIS_BAND_SLUGS = ['30-39', '40-49', '50-59'];

/** Every folded BiS band page under dist/bis, as dist-relative paths. */
function foldedBisBandPages(dist) {
  const bis = join(dist, 'bis');
  if (!existsSync(bis)) return [];
  const dirs = (path) =>
    readdirSync(path, { withFileTypes: true })
      .filter((entry) => entry.isDirectory())
      .map((entry) => entry.name);
  return dirs(bis).flatMap((classSlug) =>
    dirs(join(bis, classSlug)).flatMap((specSlug) =>
      FOLDED_BIS_BAND_SLUGS.map((slug) => `bis/${classSlug}/${specSlug}/${slug}.html`).filter((file) =>
        existsSync(join(dist, file)),
      ),
    ),
  );
}

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
  for (const file of foldedBisBandPages(dist)) failures.push(`dist/${file} still exists`);
  if (!existsSync(join(dist, 'setup.html'))) {
    failures.push('dist/setup.html is missing -- /setup should replace /addon');
  }
  return failures;
}
