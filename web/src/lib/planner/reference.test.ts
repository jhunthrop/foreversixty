// web/src/lib/planner/reference.test.ts
import { describe, expect, it } from 'vitest';
import { classRows, comboFor, comboRows, plannerHref, raceRows, racesForClass } from './reference';

describe('generated reference data', () => {
  it('loads the classes and races the sync wrote', () => {
    // Structural, not literal counts or slugs: the fixture and the real sync disagree on
    // both (nine races under the fixture's Era world, ten under the beta client's own
    // races), and this file has to be green after either one runs npm run sync.
    expect(classRows.length).toBeGreaterThan(0);
    expect(classRows.map((c) => c.slug)).toContain('paladin');
    expect(raceRows.length).toBeGreaterThan(0);
    expect(raceRows.some((r) => r.slug.includes('skyborne'))).toBe(true);
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
    // Derived from comboRows/raceRows rather than restated as a literal slug list, so this
    // does not pin the fixture's three-race Shaman roster against a real sync that gives it
    // five.
    const raceIDs = new Set(
      comboRows.filter((combo) => combo.class_id === shaman.id).map((combo) => combo.race_id),
    );
    const expected = raceRows.filter((race) => raceIDs.has(race.id)).map((r) => r.slug);
    expect(expected.length).toBeGreaterThan(0);
    expect(racesForClass(shaman.id).map((r) => r.slug)).toEqual(expected);
  });
});

describe('plannerHref', () => {
  it('deep links to a class, and to a class and race together', () => {
    expect(plannerHref('paladin')).toBe('/planner?class=paladin');
    expect(plannerHref('paladin', 'undead')).toBe('/planner?class=paladin&race=undead');
  });
});
