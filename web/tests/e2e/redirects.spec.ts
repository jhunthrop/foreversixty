// web/tests/e2e/redirects.spec.ts
// Astro preview (the server this suite runs against) does not read web/public/_redirects --
// only Cloudflare Pages does. This spec proves the file's shape and contents directly rather
// than exercising a redirect through the dev server; the companion build-time check
// (scripts/check-no-removed-routes.mjs, run from `npm run build`) proves the routes it names
// are actually gone from the built site.
import { readFileSync } from 'node:fs';
import { fileURLToPath } from 'node:url';
import { expect, test } from '@playwright/test';

const REDIRECTS_PATH = fileURLToPath(new URL('../../public/_redirects', import.meta.url));

const EXPECTED_ROWS: ReadonlyArray<readonly [string, string]> = [
  ['/classes', '/guides'],
  ['/zones', '/guides'],
  ['/zones/*', '/guides'],
  ['/dungeons', '/guides'],
  ['/dungeons/*', '/guides'],
  ['/search', '/guides'],
  ['/skyborne', '/changelog'],
  ['/everything-we-know', '/changelog'],
  ['/editions', '/changelog'],
  ['/addon', '/setup'],
];

test('every removed route in spec 3.2 has a 301 row in _redirects', () => {
  const text = readFileSync(REDIRECTS_PATH, 'utf8');
  const rows = text
    .split('\n')
    .map((line) => line.trim())
    .filter((line) => line !== '')
    .map((line) => line.split(/\s+/));

  for (const [from, to] of EXPECTED_ROWS) {
    const row = rows.find((r) => r[0] === from);
    expect(row, `no _redirects row for ${from}`).toBeDefined();
    expect(row?.[1]).toBe(to);
    expect(row?.[2]).toBe('301');
  }
});

test('every row is exactly three fields: from, to, 301', () => {
  const text = readFileSync(REDIRECTS_PATH, 'utf8');
  const rows = text
    .split('\n')
    .map((line) => line.trim())
    .filter((line) => line !== '');
  for (const row of rows) {
    expect(row.split(/\s+/)).toHaveLength(3);
  }
});
