// web/src/lib/guides/race-table.test.ts
// Runs against whichever race data the sync published: the real build's (ten races, racial
// sentences) or the unit-test fixture's (a neutral `skyborne`, no racial sentences). Nothing here
// assumes one set's slugs or sentences; what the editorial table names is checked against the
// race text only where the data carries that text.
import { describe, expect, it } from 'vitest';
import { classRows, raceRows, racesForClass } from '../planner/reference';
import {
  RACIAL_NAMES,
  RACIAL_PICKS,
  racialEffects,
  raceNameFor,
  raceTableFor,
  recommendedRaceLines,
  specKindOf,
} from './race-table';

const mageRaces = racesForClass(classRows.find((row) => row.slug === 'mage')!.id);
const raceBy = (slug: string) => raceRows.find((race) => race.slug === slug);

describe('the editorial racial tables against the race data', () => {
  it('names only picks that are among the race’s own racials', () => {
    for (const [slug, picks] of Object.entries(RACIAL_PICKS)) {
      for (const pick of Object.values(picks).flat()) {
        expect(RACIAL_NAMES[slug], slug).toContain(pick);
      }
    }
  });

  it('finds, with an effect, every racial a pick names wherever the data describes that race’s racials', () => {
    for (const race of raceRows) {
      const effects = racialEffects(race);
      if (effects.size === 0) continue;
      for (const pick of new Set(Object.values(RACIAL_PICKS[race.slug] ?? {}).flat())) {
        if (race.forever_changes.some((change) => change.text.includes(pick))) {
          expect(effects.get(pick), `${race.slug} ${pick}`).toBeTruthy();
        }
      }
    }
  });
});

describe('racialEffects', () => {
  it('cuts each described racial’s effect out of the data sentence, without the next racial’s name', () => {
    for (const race of raceRows) {
      const effects = racialEffects(race);
      for (const [name, effect] of effects) {
        expect(effect, `${race.slug} ${name}`).not.toBe('');
        for (const other of RACIAL_NAMES[race.slug] ?? []) {
          if (other !== name && !name.includes(other))
            expect(effect, `${race.slug} ${name}`).not.toContain(other);
        }
      }
    }
  });

  it('reads a known racial’s effect in the data’s own words when the data has it', () => {
    const troll = raceBy('troll');
    const berserking = troll === undefined ? undefined : racialEffects(troll).get('Berserking');
    if (berserking !== undefined) expect(berserking).toMatch(/casting and attack speed/);
  });
});

describe('specKindOf', () => {
  it('sorts casters, melee, hunters, healers and tanks', () => {
    expect(specKindOf('mage', 'mage-frost', 'dps')).toBe('caster');
    expect(specKindOf('priest', 'priest-shadow', 'dps')).toBe('caster');
    expect(specKindOf('priest', 'priest-holy', 'healer')).toBe('healer');
    expect(specKindOf('warrior', 'warrior-protection', 'tank')).toBe('tank');
    expect(specKindOf('warrior', 'warrior-fury', 'dps')).toBe('melee');
    expect(specKindOf('hunter', 'hunter-marksmanship', 'dps')).toBe('ranged');
    expect(specKindOf('hunter', 'hunter-survival', 'dps')).toBe('melee');
  });
});

describe('raceTableFor', () => {
  it('has one row per race the class can be, in the data’s order, with the recommended ones marked', () => {
    const slugs = mageRaces.map((race) => race.slug);
    const recommended = slugs.filter((slug) => ['gnome', 'troll'].includes(slug));
    const rows = raceTableFor('mage', ['gnome', 'troll'], 'caster');
    expect(rows.map((row) => row.slug)).toEqual(slugs);
    expect(rows.filter((row) => row.recommended).map((row) => row.slug)).toEqual(recommended);
  });

  it('gives a row the racials the data describes, and none when it describes none', () => {
    for (const row of raceTableFor('mage', [], 'caster')) {
      const race = raceBy(row.slug)!;
      const effects = racialEffects(race);
      if (effects.size === 0) expect(row.racials, row.slug).toEqual([]);
      else {
        expect(row.racials.length, row.slug).toBeGreaterThan(0);
        for (const racial of row.racials)
          expect(effects.has(racial.name), `${row.slug} ${racial.name}`).toBe(true);
      }
      expect(raceNameFor(row.slug)).toBe(race.name);
    }
  });

  it('puts the editorial racials first when the data carries them', () => {
    const rows = raceTableFor('mage', [], 'caster');
    const gnome = rows.find((row) => row.slug === 'gnome');
    const gnomeData = raceBy('gnome');
    if (gnome === undefined || gnomeData === undefined || !racialEffects(gnomeData).has('Escape Artist'))
      return;
    expect(gnome.racials.map((racial) => racial.name)).toEqual(['Escape Artist', 'Expansive Mind']);
    expect(rows.flatMap((row) => row.racials.map((racial) => racial.name))).not.toContain('Eureka!');
  });

  it('falls back to the race’s first described racial when the editorial pick is not in the data', () => {
    const picks = RACIAL_PICKS.troll;
    const troll = raceBy('troll');
    if (picks === undefined || troll === undefined) return;
    const effects = racialEffects(troll);
    if (effects.size === 0) return;
    const saved = [...picks.default];
    picks.default.splice(0, picks.default.length, 'No Such Racial');
    try {
      const row = raceTableFor('mage', [], 'caster').find((r) => r.slug === 'troll');
      if (row === undefined) return;
      expect(row.racials).toHaveLength(1);
      expect(row.racials[0]!.name).toBe([...effects.keys()][0]);
    } finally {
      picks.default.splice(0, picks.default.length, ...saved);
    }
  });

  it('is empty for a class it does not know', () => {
    expect(raceTableFor('nonexistent', [], 'melee')).toEqual([]);
  });
});

describe('recommendedRaceLines', () => {
  it('lists each recommended race that has racials, with their names', () => {
    const rows = raceTableFor('mage', ['gnome', 'troll'], 'caster');
    const lines = recommendedRaceLines(rows);
    for (const line of lines) {
      const row = rows.find((r) => r.slug === line.slug)!;
      expect(row.recommended).toBe(true);
      expect(line.names).toEqual(row.racials.map((racial) => racial.name));
    }
    expect(lines.length).toBe(rows.filter((row) => row.recommended && row.racials.length > 0).length);
  });
});

describe('raceNameFor', () => {
  it('turns a slug into the race’s name, and passes an unknown slug through', () => {
    const [first] = raceRows;
    expect(raceNameFor(first!.slug)).toBe(first!.name);
    expect(raceNameFor('who')).toBe('who');
  });
});
