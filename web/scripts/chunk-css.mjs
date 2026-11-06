// web/scripts/chunk-css.mjs
// Every stylesheet a built chunk preloads must exist. Pages inline their CSS
// (astro.config.mjs `inlineStylesheets: 'always'`), so the build emits no file for an island's
// scoped CSS, yet a page that imports that island lazily (/logs imports Account.svelte for its
// pairing block) still preloads the file; the preload 404s, Vite's helper rejects the import and
// the lazily loaded block never draws. A pure function of the dist directory, so
// src/lib/site/chunk-css.test.ts can drive it (plain .mjs: postbuild runs it under node).
import { existsSync, readdirSync, readFileSync } from 'node:fs';
import { join } from 'node:path';

/** The `_astro/<name>.css` paths a chunk names in its preload dependency list. */
const CSS_DEP = /["'`](_astro\/[^"'`]+\.css)["'`]/g;

/** One line per stylesheet a chunk under `dist/_astro` preloads that `dist` does not hold. */
export function missingChunkCss(dist) {
  const assets = join(dist, '_astro');
  if (!existsSync(assets)) return [];
  const failures = [];
  for (const name of readdirSync(assets)
    .filter((file) => file.endsWith('.js'))
    .sort()) {
    for (const [, css] of readFileSync(join(assets, name), 'utf8').matchAll(CSS_DEP)) {
      if (!existsSync(join(dist, css)))
        failures.push(`_astro/${name} preloads ${css}, which the build did not emit`);
    }
  }
  return [...new Set(failures)];
}
