// web/scripts/check-no-removed-routes.mjs
// Astro's static build format is `file` (astro.config.mjs), so a route at /classes builds to
// dist/classes.html and a dynamic route at /zones/[slug] builds one dist/zones/<slug>.html
// per entry. Cloudflare Pages' _redirects only takes effect once the site is deployed --
// nothing in this repo's own build or test pipeline exercises it -- so this script is the
// build-time proof that the routes spec 2026-09-25 section 3.1 removes are actually gone
// from what ships, run from `postbuild` alongside check-island-size.mjs.
import { existsSync, readdirSync } from 'node:fs';
import { join } from 'node:path';

const DIST = new URL('../dist/', import.meta.url).pathname;

const MUST_NOT_EXIST = [
  'classes.html',
  'zones.html',
  'dungeons.html',
  'search.html',
  'skyborne.html',
  'everything-we-know.html',
  'editions.html',
  'addon.html',
];

const MUST_NOT_EXIST_DIRS = ['zones', 'dungeons'];

const failures = [];

for (const file of MUST_NOT_EXIST) {
  if (existsSync(join(DIST, file))) failures.push(`dist/${file} still exists`);
}

for (const dir of MUST_NOT_EXIST_DIRS) {
  const path = join(DIST, dir);
  if (existsSync(path) && readdirSync(path).length > 0) {
    failures.push(`dist/${dir}/ still has files: ${readdirSync(path).join(', ')}`);
  }
}

if (!existsSync(join(DIST, 'setup.html'))) {
  failures.push('dist/setup.html is missing -- /setup should replace /addon');
}

if (failures.length > 0) {
  console.error(
    'check-no-removed-routes: the cut left routes behind:\n' + failures.map((f) => `  - ${f}`).join('\n'),
  );
  process.exit(1);
}

console.log('check-no-removed-routes: every removed route is gone from dist/, /setup is present.');
