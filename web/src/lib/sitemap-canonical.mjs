// web/src/lib/sitemap-canonical.mjs
// The sitemap lists only pages that are their own canonical. A page whose <link rel="canonical">
// names another URL (/bis/<class>/<spec>/60/pre-raid while the level 60 list is already the
// pre-raid list) is an alias: listing it asks Google to index a URL the page itself says not to,
// which Search Console reports as "Alternate page with proper canonical tag" and which spends
// crawl on duplicates. astro.config.mjs's sitemap filter reads the built page through this.
import { readFileSync } from 'node:fs';
import { join } from 'node:path';

const CANONICAL = /<link\b[^>]*\brel="canonical"[^>]*>/i;
const HREF = /\bhref="([^"]*)"/i;

/**
 * The canonical URL a page's HTML declares, or null when it declares none.
 * @param {string} html
 * @returns {string | null}
 */
export function canonicalOf(html) {
  const tag = CANONICAL.exec(html)?.[0];
  return tag === undefined ? null : (HREF.exec(tag)?.[1] ?? null);
}

/**
 * The built file for a page URL under `build.format: 'file'`: /sim is dist/sim.html, / is
 * dist/index.html.
 * @param {string} distDir
 * @param {string} pageUrl
 * @returns {string}
 */
export function builtFileOf(distDir, pageUrl) {
  const path = new URL(pageUrl).pathname.replace(/\/+$/, '');
  return join(distDir, path === '' ? 'index.html' : `${path}.html`);
}

/**
 * Whether the page at `pageUrl` is its own canonical (or names none). A page with no built file
 * is kept: the sitemap integration only lists pages the build wrote, so a miss is a path shape
 * this reader does not know, and dropping a real page from the sitemap is the worse failure.
 * @param {string} distDir
 * @param {string} pageUrl
 * @returns {boolean}
 */
export function isSelfCanonical(distDir, pageUrl) {
  let html;
  try {
    html = readFileSync(builtFileOf(distDir, pageUrl), 'utf8');
  } catch {
    return true;
  }
  const canonical = canonicalOf(html);
  return canonical === null || canonical.replace(/\/+$/, '') === pageUrl.replace(/\/+$/, '');
}
