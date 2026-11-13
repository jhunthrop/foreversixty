// web/src/lib/og-pages.test.ts
// Base.astro gives every page og:image = /og<path>.png, so every static page needs a card or
// its link preview is a 404 (the tier lists, the sim and Premium all were, at launch).
import { readdirSync } from 'node:fs';
import { join, relative } from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import { STATIC_OG_PAGES } from './og-pages';

const pagesRoot = fileURLToPath(new URL('../pages', import.meta.url));

/** Pages nobody shares: never in the sitemap, plumbing, or the owner's admin. */
const NO_CARD = new Set([
  '404',
  'admin/links',
  'admin/creators',
  'admin/creators/payouts',
  'admin/testers',
  'admin/activity',
  'app/auth',
  'creators/apply',
  'creators/onboarding',
  'creators/stats',
  'design-sheet',
  'guild-shell',
  'report-shell',
  'reports',
]);

/** The card path of every static (non-dynamic) .astro page under src/pages. */
function staticPagePaths(dir: string): string[] {
  return readdirSync(dir, { withFileTypes: true }).flatMap((entry) => {
    const full = join(dir, entry.name);
    if (entry.isDirectory()) return entry.name.startsWith('[') ? [] : staticPagePaths(full);
    if (!entry.name.endsWith('.astro') || entry.name.startsWith('[')) return [];
    const route = relative(pagesRoot, full)
      .replace(/\.astro$/, '')
      .replace(/\/index$/, '');
    return [route];
  });
}

describe('STATIC_OG_PAGES', () => {
  it('has a share card for every static page that can be shared', () => {
    const missing = staticPagePaths(pagesRoot).filter(
      (path) => !NO_CARD.has(path) && !(path in STATIC_OG_PAGES),
    );
    expect(missing).toEqual([]);
  });

  it('covers the three tier lists, each with its own title', () => {
    const titles = ['tiers', 'tiers/tank', 'tiers/healer'].map((path) => STATIC_OG_PAGES[path]?.title);
    expect(titles).toEqual(['DPS tier list', 'Tank tier list', 'Healer tier list']);
  });
});
