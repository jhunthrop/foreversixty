// web/src/pages/_headers.test.ts
// Underscore-prefixed so Astro does not build it as a route (see _logs.test.ts).
// The browser caching for the crests and page art (owner bug 2026-11-09): without a rule they
// took Cloudflare's max-age=0, so every navigation revalidated every crest and Safari drew the
// header's character crest empty for that round trip.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { describe, expect, it } from 'vitest';

const headers = readFileSync(path.join(import.meta.dirname, '../../public/_headers'), 'utf8');

/** The header lines of one rule: the indented lines under its path line. */
function ruleBlock(rulePath: string): string {
  const start = headers.indexOf(`\n${rulePath}\n`);
  expect(start, `${rulePath} has a rule`).toBeGreaterThanOrEqual(0);
  const rest = headers.slice(start + rulePath.length + 2);
  return rest.split(/\n(?=\S)/)[0] ?? rest;
}

describe('public/_headers', () => {
  for (const rulePath of ['/icons/*', '/art/*']) {
    it(`caches ${rulePath} for a day with a week of stale, never as immutable`, () => {
      const block = ruleBlock(rulePath);
      expect(block).toContain('Cache-Control: public, max-age=86400, stale-while-revalidate=604800');
      expect(block).not.toContain('immutable');
    });
  }

  it('keeps the workers.dev preview host out of the index', () => {
    expect(ruleBlock('https://foreversixty.jhunthrop.workers.dev/*')).toContain('X-Robots-Tag: noindex');
  });
});
