import { describe, expect, it } from 'vitest';
import catalog from '../../../../apps/shared/copy/shared.json';
import { sharedPlural, sharedText } from './shared';

describe('shared copy', () => {
  it('reads a plain string', () => {
    expect(sharedText('shared.cta.getAddonAndCompanion')).toBe('Get the addon and companion');
  });

  it('fills placeholders, typed entries included', () => {
    expect(sharedText('shared.account.delete.mismatch', { battletag: 'Felsnap#1234' })).toBe(
      'Type Felsnap#1234 to confirm.',
    );
    expect(sharedText('shared.lowLevel.show', { count: 3 })).toBe('Show low-level characters (3)');
  });

  it('picks the plural form by count', () => {
    expect(sharedPlural('shared.home.upgradesCount', 1)).toBe('1 upgrade');
    expect(sharedPlural('shared.home.upgradesCount', 3)).toBe('3 upgrades');
  });

  it('fails loudly on a missing value or the wrong kind of entry', () => {
    expect(() => sharedText('shared.home.simmed')).toThrow(/no value for \{date\}/);
    expect(() => sharedText('shared.home.upgradesCount')).toThrow(/not a text entry/);
    expect(() => sharedPlural('shared.home.simmed', 2)).toThrow(/not a plural entry/);
  });

  it('keeps every key under shared. and no em dash in any word', () => {
    for (const [key, value] of Object.entries(catalog.strings)) {
      expect(key.startsWith('shared.'), key).toBe(true);
      expect(JSON.stringify(value), key).not.toContain('—');
    }
  });
});
