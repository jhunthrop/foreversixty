import { describe, expect, it } from 'vitest';
import { levelColorClass } from './levels';

describe('levelColorClass', () => {
  it('switches rarity colour at the bracket boundaries', () => {
    expect(levelColorClass(29)).toBe('text-rarity-uncommon');
    expect(levelColorClass(30)).toBe('text-rarity-rare-text');
    expect(levelColorClass(54)).toBe('text-rarity-rare-text');
    expect(levelColorClass(55)).toBe('text-rarity-epic-text');
  });

  it('is muted when the lower bound is unknown', () => {
    expect(levelColorClass()).toBe('text-muted');
  });
});
