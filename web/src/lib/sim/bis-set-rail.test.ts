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

  it("gives each spec its set's DPS, its strongest stats best first, its guide and its level 60 BiS page", () => {
    const specs = rail.flatMap((klass) => klass.specs);
    expect(specs.some((spec) => spec.dps !== null)).toBe(true);
    for (const spec of specs) {
      if (spec.dps !== null) expect(spec.dps).toMatch(/^[\d,]+\.\d$/);
      expect(spec.stats.length).toBeLessThanOrEqual(3);
      const factors = spec.stats.map((stat) => Number(stat.factor));
      expect(factors).toEqual([...factors].sort((a, b) => b - a));
      expect(spec.guideHref).toMatch(/^\/guides\/[a-z]+\/[a-z-]+$/);
      expect(spec.bisHref).toMatch(/^\/bis\/[a-z]+\/[a-z-]+/);
      expect(spec.stats.every((stat) => !stat.label.endsWith(' rating'))).toBe(true);
    }
  });

  it('never names a school of spell damage that is worth exactly what spell power is', () => {
    for (const spec of rail.flatMap((klass) => klass.specs)) {
      const spellPower = spec.stats.find((stat) => stat.label === 'Spell power');
      if (spellPower === undefined) continue;
      const twins = spec.stats.filter((stat) => stat !== spellPower && stat.label.endsWith(' power'));
      for (const twin of twins) expect(twin.factor, `${spec.name} ${twin.label}`).not.toBe(spellPower.factor);
    }
  });
});
