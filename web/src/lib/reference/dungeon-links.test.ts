// web/src/lib/reference/dungeon-links.test.ts
import { describe, expect, it } from 'vitest';
import { dungeonHasLoot, zoneEntryForDungeon, type LootFile, type ZoneEntry } from './dungeon-links';

const ZONES: ZoneEntry[] = [
  { id: 'mount-hyjal', title: 'Mount Hyjal' },
  { id: 'riverglades', title: 'Riverglades' },
];

describe('zoneEntryForDungeon', () => {
  it('matches a zone by its exact title', () => {
    expect(zoneEntryForDungeon(ZONES, 'Riverglades')).toEqual({ id: 'riverglades', title: 'Riverglades' });
  });

  it('returns null for an old-world zone this site has no page for', () => {
    expect(zoneEntryForDungeon(ZONES, 'Dun Morogh')).toBeNull();
  });

  it('returns null for "Not yet known"', () => {
    expect(zoneEntryForDungeon(ZONES, 'Not yet known')).toBeNull();
  });
});

describe('dungeonHasLoot', () => {
  const loot: LootFile = {
    sources: [
      { id: 'dungeon:hall-of-thanes', kind: 'dungeon' },
      { id: 'raid:molten-core', kind: 'raid' },
    ],
  };

  it('is true when the loot file records a dungeon source at this slug', () => {
    expect(dungeonHasLoot(loot, 'hall-of-thanes')).toBe(true);
  });

  it('is false for a slug the loot file has no dungeon source for', () => {
    expect(dungeonHasLoot(loot, 'alcaz-prison')).toBe(false);
  });

  it('is false for a source of a different kind at the same-looking id', () => {
    expect(dungeonHasLoot(loot, 'molten-core')).toBe(false);
  });
});
