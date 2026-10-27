// web/src/lib/site/pricing-redirect.test.ts
// /pricing is an edge 301 to /premium: the static build has no page for it, so the rule in
// public/_redirects is the only thing that answers the old URL.
import { readFileSync } from 'node:fs';
import { describe, expect, it } from 'vitest';

describe('public/_redirects', () => {
  it('sends /pricing to /premium with a 301', () => {
    const rules = readFileSync(new URL('../../../public/_redirects', import.meta.url), 'utf8')
      .split('\n')
      .map((line) => line.trim().split(/\s+/));
    expect(rules).toContainEqual(['/pricing', '/premium', '301']);
  });
});
