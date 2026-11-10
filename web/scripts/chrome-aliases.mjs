// web/scripts/chrome-aliases.mjs
// The API's /b/{id} page wears a copy of the site's header (api/internal/site/chrome.html), and
// that copy is deployed separately from the site. Its scripts cannot point at the build's
// content-hashed /_astro/<name>.<hash>.js files: the next site deploy removes them. So every
// build also writes one stable module per header script, /_chrome/<name>.js, that re-exports the
// hashed file of THIS build, and the API's copy names the stable paths (chrome_test.go's
// extractChrome applies the same rule as stableChromePath here). Run from `postbuild`.
import { mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import path from 'node:path';

/** Where the stable modules live in dist/ and on the site. Not under /_astro/: no year-long cache. */
export const CHROME_DIR = '_chrome';

/** A built script's URL: /_astro/<name>.<8-character Vite hash>.js. */
const HASHED_SCRIPT = /\/_astro\/([^"'\s/]+)\.([A-Za-z0-9_-]{8})\.js/g;

/** Whether a built module has a default export (an island component or a renderer). */
const DEFAULT_EXPORT = /export\s*\{[^}]*\bas\s+default\b[^}]*\}|export\s+default\b/;

/** The stable path for one hashed script URL. */
export function stableChromePath(name) {
  return `/${CHROME_DIR}/${name}.js`;
}

/** Every hashed script a built page names, as { name, url }, one per name. */
export function hashedScripts(html) {
  const byName = new Map();
  for (const [url, name] of html.matchAll(HASHED_SCRIPT)) {
    const seen = byName.get(name);
    if (seen !== undefined && seen !== url) {
      throw new Error(
        `chrome-aliases: ${name} is built twice (${seen}, ${url}); its stable path would be ambiguous`,
      );
    }
    byName.set(name, url);
  }
  return [...byName].map(([name, url]) => ({ name, url }));
}

/** The stable module's source: everything the hashed module exports, its default included. */
export function aliasModule(url, source) {
  const lines = [`export * from '${url}';`];
  if (DEFAULT_EXPORT.test(source)) lines.push(`export { default } from '${url}';`);
  return lines.join('\n') + '\n';
}

/** Writes dist/_chrome/<name>.js for every hashed script the page names; returns the names. */
export function writeChromeAliases(dist, page = 'about.html') {
  const scripts = hashedScripts(readFileSync(path.join(dist, page), 'utf8'));
  if (scripts.length === 0) throw new Error(`chrome-aliases: ${page} names no /_astro/ script`);
  const outDir = path.join(dist, CHROME_DIR);
  mkdirSync(outDir, { recursive: true });
  for (const { name, url } of scripts) {
    const source = readFileSync(path.join(dist, url), 'utf8');
    writeFileSync(path.join(outDir, `${name}.js`), aliasModule(url, source));
  }
  return scripts.map(({ name }) => name);
}
