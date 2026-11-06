// web/src/lib/report/aura-names.test.ts
import { describe, expect, it } from 'vitest';
import { isUnnamed, nameAuras } from './aura-names';
import type { AuraTrack } from './types';

const track = (spell_id: number, name: string): AuraTrack => ({ spell_id, name }) as unknown as AuraTrack;

describe('aura names (B2)', () => {
  it('knows the engine placeholder', () => {
    expect(isUnnamed('Spell #17')).toBe(true);
    expect(isUnnamed('')).toBe(true);
    expect(isUnnamed('Power Word: Shield')).toBe(false);
  });

  it('names a placeholder the build table knows, and leaves the rest', () => {
    const named = nameAuras(
      [track(17, 'Spell #17'), track(871, 'Spell #871'), track(5, 'Spell #5'), track(1, 'Fort')],
      {
        '17': 'Power Word: Shield',
        '871': 'Shield Wall',
      },
    );
    expect(named.map((t) => t.name)).toEqual(['Power Word: Shield', 'Shield Wall', 'Spell #5', 'Fort']);
  });
});
