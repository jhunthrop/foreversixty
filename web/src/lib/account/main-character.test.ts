import { describe, expect, it } from 'vitest';
import { mainCharacter, pointerForCharacter } from './main-character';
import type { MeCharacter } from './api';

function character(overrides: Partial<MeCharacter>): MeCharacter {
  return {
    key: 'us/normal/test',
    region: 'us',
    ruleset: 'normal',
    name: 'Test',
    ...overrides,
  };
}

describe('mainCharacter', () => {
  it('is null with no characters', () => {
    expect(mainCharacter([])).toBeNull();
  });

  it('picks the only simmable character over a higher-level one with no build', () => {
    const simmable = character({
      key: 'us/normal/a',
      name: 'A',
      level: 10,
      build: { source: 'blizzard', captured_at: '2026-09-20T00:00:00Z' },
    });
    const noBuild = character({ key: 'us/normal/b', name: 'B', level: 60 });
    expect(mainCharacter([noBuild, simmable])).toBe(simmable);
  });

  it('picks the newest captured_at among simmable characters', () => {
    const older = character({
      key: 'us/normal/a',
      name: 'A',
      build: { source: 'addon', captured_at: '2026-09-19T00:00:00Z' },
    });
    const newer = character({
      key: 'us/normal/b',
      name: 'B',
      build: { source: 'blizzard', captured_at: '2026-09-21T00:00:00Z' },
    });
    expect(mainCharacter([older, newer])).toBe(newer);
  });

  it('breaks a captured_at tie by level', () => {
    const lower = character({
      key: 'us/normal/a',
      name: 'A',
      level: 40,
      build: { source: 'addon', captured_at: '2026-09-21T00:00:00Z' },
    });
    const higher = character({
      key: 'us/normal/b',
      name: 'B',
      level: 60,
      build: { source: 'addon', captured_at: '2026-09-21T00:00:00Z' },
    });
    expect(mainCharacter([lower, higher])).toBe(higher);
  });

  it('falls back to the highest level when nothing is simmable', () => {
    const low = character({ key: 'us/normal/a', name: 'A', level: 10 });
    const high = character({ key: 'us/normal/b', name: 'B', level: 45 });
    expect(mainCharacter([low, high])).toBe(high);
  });

  it('treats an undefined level as the lowest', () => {
    const withLevel = character({ key: 'us/normal/a', name: 'A', level: 5 });
    const noLevel = character({ key: 'us/normal/b', name: 'B' });
    expect(mainCharacter([noLevel, withLevel])).toBe(withLevel);
  });
});

describe('pointerForCharacter', () => {
  it('writes an armory-kind pointer keyed by the character', () => {
    const c = character({ key: 'us/normal/a', name: 'Aria', class: 'Mage' });
    const pointer = pointerForCharacter(c);
    expect(pointer.source).toBe('armory');
    expect(pointer.ref).toBe('us/normal/a');
    expect(pointer.classSlug).toBe('mage');
    expect(pointer.label).toBe('Aria · Mage');
  });

  it('falls back to the character name alone with no class on file', () => {
    const c = character({ key: 'us/normal/a', name: 'Aria' });
    const pointer = pointerForCharacter(c);
    expect(pointer.label).toBe('Aria');
    expect(pointer.classSlug).toBe('');
  });
});

describe('mainCharacter with a chosen main', () => {
  const low = { key: 'us/pvp/dottzz', region: 'us', ruleset: 'pvp', name: 'Dottzz', level: 12 };
  const high = { key: 'us/pvp/reloadd', region: 'us', ruleset: 'pvp', name: 'Reloadd', level: 25 };

  it("returns the chosen character over the site's own guess", () => {
    expect(mainCharacter([low, high], 'us/pvp/dottzz')?.key).toBe('us/pvp/dottzz');
  });

  it('falls back to the guess when the chosen key is no longer one of the characters', () => {
    expect(mainCharacter([low, high], 'us/pvp/gone')?.key).toBe('us/pvp/reloadd');
  });
});
