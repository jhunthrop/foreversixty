// web/src/lib/planner/reference.test.ts
import { describe, expect, it } from 'vitest';
import { classRows, comboFor, comboRows, plannerHref, raceRows, racesForClass } from './reference';

describe('generated reference data', () => {
  it('loads the nine classes and nine races the sync wrote', () => {
    expect(classRows.map((c) => c.slug)).toContain('paladin');
    expect(raceRows.map((r) => r.slug)).toContain('skyborne');
    expect(comboRows.length).toBeGreaterThan(0);
  });
});

describe('comboFor', () => {
  it('finds the Forever addition and returns null for an illegal pair', () => {
    const undead = raceRows.find((r) => r.slug === 'undead')!;
    const paladin = classRows.find((c) => c.slug === 'paladin')!;
    const shaman = classRows.find((c) => c.slug === 'shaman')!;
    expect(comboFor(undead.id, paladin.id)?.new_in_forever).toBe(true);
    expect(comboFor(undead.id, shaman.id)).toBeNull();
  });
});

describe('racesForClass', () => {
  it('lists only the races that can be that class, in race order', () => {
    const shaman = classRows.find((c) => c.slug === 'shaman')!;
    expect(racesForClass(shaman.id).map((r) => r.slug)).toEqual(['orc', 'tauren', 'troll']);
  });
});

describe('plannerHref', () => {
  it('deep links to a class, and to a class and race together', () => {
    expect(plannerHref('paladin')).toBe('/planner?class=paladin');
    expect(plannerHref('paladin', 'undead')).toBe('/planner?class=paladin&race=undead');
  });
});
