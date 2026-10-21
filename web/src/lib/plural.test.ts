import { describe, expect, it } from 'vitest';
import { counted, noun } from './plural';

describe('counted', () => {
  it('uses the singular for exactly one', () => {
    expect(counted(1, 'raider')).toBe('1 raider');
    expect(counted(1, 'boss', 'bosses')).toBe('1 boss');
    expect(counted(1, 'named kill')).toBe('1 named kill');
  });

  it('uses the plural for zero and for more than one', () => {
    expect(counted(0, 'wipe')).toBe('0 wipes');
    expect(counted(2, 'boss', 'bosses')).toBe('2 bosses');
    expect(counted(4, 'trash pack')).toBe('4 trash packs');
  });
});

describe('noun', () => {
  it('returns the noun alone', () => {
    expect(noun(1, 'kill')).toBe('kill');
    expect(noun(2, 'kill')).toBe('kills');
    expect(noun(1, 'boss', 'bosses')).toBe('boss');
  });
});
