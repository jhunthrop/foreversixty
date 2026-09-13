import { describe, expect, it } from 'vitest';
import { levelColorClass, levelRange, levelRangeLong } from './levels';

describe('levelRange', () => {
  it('joins a complete range and falls back to an em dash', () => {
    expect(levelRange(30, 40)).toBe('30–40');
    expect(levelRange(undefined, 40)).toBe('—');
    expect(levelRange(30, undefined)).toBe('—');
    expect(levelRange()).toBe('—');
  });
});

describe('levelRangeLong', () => {
  it('spells out a complete range and says so when there is none', () => {
    expect(levelRangeLong(52, 60)).toBe('Level 52–60');
    expect(levelRangeLong()).toBe('Level range not yet known');
  });
});

describe('levelColorClass', () => {
  it('switches rarity colour at the bracket boundaries', () => {
    expect(levelColorClass(29)).toBe('text-rarity-uncommon');
    expect(levelColorClass(30)).toBe('text-rarity-rare');
    expect(levelColorClass(54)).toBe('text-rarity-rare');
    expect(levelColorClass(55)).toBe('text-rarity-epic');
  });

  it('is muted when the lower bound is unknown', () => {
    expect(levelColorClass()).toBe('text-muted');
  });
});
