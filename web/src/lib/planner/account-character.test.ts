import { describe, expect, it } from 'vitest';
import type { MeCharacter } from '../account/api';
import { decodeFS1 } from './fs1';
import { exportCodeFromAccount, raceSlugOf } from './account-character';

const BOW: MeCharacter = {
  key: 'us/pvp/bow-jackzon',
  region: 'us',
  ruleset: 'pvp',
  name: 'Bow Jackzon',
  realm: 'Realm',
  class: 'hunter',
  race: 'Dwarf',
  level: 60,
  build: {
    source: 'addon',
    captured_at: '2026-10-21T00:00:00Z',
    gear: { head: 12640 },
    talents: { trees: ['0000', '503200', '1'], points: [0, 10, 1] },
    level: 60,
    data_build: '1.60.1.70450',
  },
};

describe('raceSlugOf', () => {
  it('slugs the race name the account carries', () => {
    expect(raceSlugOf('Dwarf')).toBe('dwarf');
    expect(raceSlugOf('Night Elf')).toBe('night-elf');
    expect(raceSlugOf(undefined)).toBe('');
  });
});

describe('exportCodeFromAccount', () => {
  it('rebuilds the character as an FS1 code the planner reads: ranks, gear, race, level, name', () => {
    const code = exportCodeFromAccount(BOW, '1.60.1.70450');
    expect(code).not.toBeNull();
    const decoded = decodeFS1(code ?? '');
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;
    expect(decoded.build).toMatchObject({
      classSlug: 'hunter',
      raceSlug: 'dwarf',
      dataBuild: '1.60.1.70450',
      level: 60,
      character: { name: 'Bow Jackzon', realm: 'Realm' },
    });
    expect(decoded.build.treeRanks).toEqual([[0], [5, 0, 3, 2], [1]]);
    expect(decoded.build.gear.head).toBe(12640);
  });

  it('is null when the account holds no ranks, or only zeros', () => {
    expect(exportCodeFromAccount({ ...BOW, build: undefined }, 'b')).toBeNull();
    const zeros = {
      ...BOW,
      build: { ...BOW.build!, talents: { trees: ['0', '0', '0'], points: [0, 0, 0] } },
    };
    expect(exportCodeFromAccount(zeros, 'b')).toBeNull();
  });
});
