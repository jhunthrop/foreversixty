import { describe, expect, it } from 'vitest';
import catalog from '../../../../apps/shared/copy/shared.json';
import { fillShared, pluralShared } from './shared';
import { sharedAccountDeleteMismatch } from './shared/account.generated';
import { sharedCtaGetAddonAndCompanion } from './shared/cta.generated';
import { sharedHomeSimmed, sharedHomeUpgradesCount } from './shared/home.generated';
import { sharedLowLevelShow } from './shared/lowLevel.generated';

describe('shared copy', () => {
  it('exports each word on its own, one module per group', () => {
    expect(sharedCtaGetAddonAndCompanion).toBe('Get the addon and companion');
  });

  it('fills placeholders, typed entries included', () => {
    expect(fillShared(sharedAccountDeleteMismatch, { battletag: 'Felsnap#1234' })).toBe(
      'Type Felsnap#1234 to confirm.',
    );
    expect(fillShared(sharedLowLevelShow, { count: 3 })).toBe('Show low-level characters (3)');
  });

  it('picks the plural form by count', () => {
    expect(pluralShared(sharedHomeUpgradesCount, 1)).toBe('1 upgrade');
    expect(pluralShared(sharedHomeUpgradesCount, 3)).toBe('3 upgrades');
  });

  it('fails loudly on a missing value', () => {
    expect(() => fillShared(sharedHomeSimmed, {})).toThrow(/no value for \{date\}/);
  });

  it('keeps every key under shared. and no em dash in any word', () => {
    for (const [key, value] of Object.entries(catalog.strings)) {
      expect(key.startsWith('shared.'), key).toBe(true);
      expect(JSON.stringify(value), key).not.toContain('—');
    }
  });
});
