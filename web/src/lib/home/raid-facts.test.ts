// The raid plaque's facts (lib/home/raid-facts.ts).
import { describe, expect, it } from 'vitest';
import { raidFacts, raidLine } from './raid-facts';

const ONYXIA = { kind: 'raid', name: "Onyxia's Lair", bosses: [{ items: [1, 2, 3, 3] }] };
const CLASSES = new Map([
  ['hunter', new Set([1, 3, 9])],
  ['mage', new Set([2])],
]);

describe('raid facts', () => {
  it('counts each first raid bosses, distinct drops and each class share', () => {
    const [barrow, hyjal, onyxia] = raidFacts([ONYXIA, { kind: 'dungeon', name: 'Barrow Deeps' }], CLASSES);
    expect(barrow).toMatchObject({ name: 'Barrow Deeps', bosses: null, drops: null });
    expect(hyjal.bosses).toBeNull();
    expect(onyxia).toMatchObject({ bosses: 1, drops: 3, usable: { hunter: 2, mage: 1 } });
  });

  it('words a raid for a class, for anyone, and with no loot rows', () => {
    const [barrow, , onyxia] = raidFacts([ONYXIA], CLASSES);
    expect(raidLine(onyxia, { slug: 'hunter', name: 'Hunter' })).toBe('1 boss · 2 Hunter drops');
    expect(raidLine(onyxia)).toBe('1 boss · 3 drops');
    expect(raidLine(barrow, { slug: 'hunter', name: 'Hunter' })).toBe('Loot not out yet');
  });
});
