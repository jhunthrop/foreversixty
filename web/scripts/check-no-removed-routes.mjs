// web/scripts/check-no-removed-routes.mjs
// Astro's static build format is `file` (astro.config.mjs), so a route at /classes builds to
// dist/classes.html and a dynamic route at /zones/[slug] builds one dist/zones/<slug>.html
// per entry. Cloudflare Pages' _redirects only takes effect once the site is deployed --
// nothing in this repo's own build or test pipeline exercises it -- so this script is the
// build-time proof that the routes spec 2026-09-25 section 3.1 removes are actually gone
// from what ships, run from `postbuild` alongside check-island-size.mjs. The check itself
// lives in removed-routes.mjs (tested by src/lib/site/removed-routes.test.ts).
import { distDirOf, removedRouteFailures } from './removed-routes.mjs';

const failures = removedRouteFailures(distDirOf(import.meta.url));

if (failures.length > 0) {
  console.error(
    'check-no-removed-routes: the cut left routes behind:\n' + failures.map((f) => `  - ${f}`).join('\n'),
  );
  process.exit(1);
}

console.log('check-no-removed-routes: every removed route is gone from dist/, /setup is present.');
