// web/src/lib/site/pricing-redirect.test.ts
// /pricing is an edge 301 to /premium: the static build has no page for it, so the rule in
// public/_redirects is the only thing that answers the old URL.
import { readFileSync } from 'node:fs';
import { describe, expect, it } from 'vitest';
import { bandHash, bandSlug } from '../bis/band-path';

describe('public/_redirects', () => {
  it('sends /pricing to /premium with a 301', () => {
    const rules = readFileSync(new URL('../../../public/_redirects', import.meta.url), 'utf8')
      .split('\n')
      .map((line) => line.trim().split(/\s+/));
    expect(rules).toContainEqual(['/pricing', '/premium', '301']);
  });

  it('sends each folded BiS band to the spec page with its band hash, with a 301', () => {
    const rules = readFileSync(new URL('../../../public/_redirects', import.meta.url), 'utf8')
      .split('\n')
      .map((line) => line.trim().split(/\s+/));
    for (const band of [30, 40, 50]) {
      expect(rules).toContainEqual([`/bis/:class/:spec/${bandSlug(band)}`, `/bis/:class/:spec${bandHash(band)}`, '301']);
    }
    expect(rules.some((rule) => rule[0] === '/bis/:class/:spec/60')).toBe(false);
  });
});
