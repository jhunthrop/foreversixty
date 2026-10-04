// web/src/lib/home/upgrades-loader.test.ts
import { describe, expect, it } from 'vitest';
import type { MeCharacter } from '../account/api';
import { upgradesCacheKey } from './upgrades-loader';

type Subject = Pick<MeCharacter, 'key' | 'level' | 'faction' | 'build'>;
const EXPORT = { source: 'addon' as const, captured_at: '2026-10-04T00:00:00Z' };
const BASE: Subject = {
  key: 'us/pvp/thoradin',
  level: 60,
  faction: 'alliance',
  build: { ...EXPORT, gear: { head: 1, chest: 2 } },
};

describe('upgradesCacheKey', () => {
  it('names the build, the character, their band inputs and the exact worn set', () => {
    expect(upgradesCacheKey(BASE, 'b1')).toBe('fs.upgrades:b1:us/pvp/thoradin:60:alliance:chest=2,head=1');
  });

  it('changes when a worn item, the level or the faction changes', () => {
    const same = upgradesCacheKey(BASE, 'b1');
    expect(upgradesCacheKey({ ...BASE, build: { ...EXPORT, gear: { head: 1, chest: 3 } } }, 'b1')).not.toBe(
      same,
    );
    expect(upgradesCacheKey({ ...BASE, level: 59 }, 'b1')).not.toBe(same);
    expect(upgradesCacheKey({ ...BASE, faction: 'horde' }, 'b1')).not.toBe(same);
  });

  it('is stable across gear key order and tolerates a character with no export yet', () => {
    const reordered: Subject = { ...BASE, build: { ...EXPORT, gear: { chest: 2, head: 1 } } };
    expect(upgradesCacheKey(reordered, 'b1')).toBe(upgradesCacheKey(BASE, 'b1'));
    expect(upgradesCacheKey({ key: 'k' }, 'b1')).toBe('fs.upgrades:b1:k:::');
  });
});
