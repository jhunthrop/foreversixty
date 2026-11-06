import { describe, expect, it } from 'vitest';
import activeBuild from '../../data/active-build.json';
import { loadBisSetRail } from './bis-set-rail';

describe('the best-in-slot set rail', () => {
  const rail = loadBisSetRail(activeBuild.build);

  it('lists classes in the client order, each with only its damage specs that link to a ranking sim', () => {
    expect(rail.length).toBeGreaterThan(0);
    for (const klass of rail) {
      expect(klass.specs.length).toBeGreaterThan(0);
      for (const spec of klass.specs) expect(spec.href).toMatch(/^\/sim\?.*run=1/);
    }
    expect(rail.some((klass) => klass.slug === 'warrior')).toBe(true);
  });
});
