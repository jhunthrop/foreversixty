// web/src/lib/guides/race-table.test.ts
import { describe, expect, it } from 'vitest';
import { raceRows } from '../planner/reference';
import {
  RACIAL_NAMES,
  RACIAL_PICKS,
  racialEffects,
  raceNameFor,
  raceTableFor,
  specKindOf,
} from './race-table';

describe('the racial tables against the race data', () => {
  it('has an entry for every race the build ships', () => {
    for (const race of raceRows) {
      expect(RACIAL_NAMES[race.slug], `no racial names for ${race.slug}`).toBeDefined();
      expect(RACIAL_PICKS[race.slug], `no racial picks for ${race.slug}`).toBeDefined();
    }
  });

  it('names only picks that are among the race’s own racials', () => {
    for (const [slug, picks] of Object.entries(RACIAL_PICKS)) {
      for (const pick of Object.values(picks)) {
        expect(RACIAL_NAMES[slug], slug).toContain(pick);
      }
    }
  });

  it('finds every racial a pick names in the race’s own text, with an effect', () => {
    for (const race of raceRows) {
      const effects = racialEffects(race);
      for (const pick of new Set(Object.values(RACIAL_PICKS[race.slug] ?? {}))) {
        expect(effects.get(pick), `${race.slug} ${pick}`).toBeTruthy();
      }
    }
  });
});

describe('racialEffects', () => {
  const troll = raceRows.find((race) => race.slug === 'troll')!;

  it('cuts each racial’s effect out of the data sentence', () => {
    const effects = racialEffects(troll);
    expect(effects.get('Berserking')).toBe('grants 10% casting and attack speed on a 3 minute cooldown');
    expect(effects.get('Regeneration')).toBe('keeps 10% of Health regeneration running in combat');
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
  it('has one row per race the class can be, with the recommended ones marked', () => {
    const rows = raceTableFor('mage', ['gnome', 'troll'], 'caster');
    expect(rows.map((row) => row.slug)).toEqual([
      'human',
      'orc',
      'undead',
      'gnome',
      'troll',
      'high-order-skyborne',
    ]);
    expect(rows.filter((row) => row.recommended).map((row) => row.slug)).toEqual(['gnome', 'troll']);
  });

  it('gives each row the racial that matters for the kind of spec, and what it does', () => {
    const mage = raceTableFor('mage', [], 'caster');
    expect(mage.find((row) => row.slug === 'troll')).toMatchObject({
      racial: 'Berserking',
      faction: 'horde',
      effect: 'Grants 10% casting and attack speed on a 3 minute cooldown',
    });
    expect(mage.find((row) => row.slug === 'human')?.racial).toBe('The Human Spirit');
    expect(raceTableFor('warrior', [], 'melee').find((row) => row.slug === 'human')?.racial).toBe(
      'Sword Specialization',
    );
  });

  it('is empty for a class it does not know', () => {
    expect(raceTableFor('nonexistent', [], 'melee')).toEqual([]);
  });
});

describe('raceNameFor', () => {
  it('turns a slug into the race’s name, and passes an unknown slug through', () => {
    expect(raceNameFor('gnome')).toBe('Gnome');
    expect(raceNameFor('who')).toBe('who');
  });
});
