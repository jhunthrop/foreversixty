import { describe, expect, it } from 'vitest';
import { LOW_LEVEL_BELOW, isLowLevel, lowLevelCopy, splitByLevel } from './level-filter';

const at = (key: string, level?: number) => ({ key, level });

describe('the low-level rule', () => {
  it('folds below level 10 and never an unknown level', () => {
    expect(LOW_LEVEL_BELOW).toBe(10);
    expect(isLowLevel(at('a', 9))).toBe(true);
    expect(isLowLevel(at('a', 10))).toBe(false);
    expect(isLowLevel(at('a'))).toBe(false);
  });

  it('shows level 10 and up by default, in order, and counts the rest', () => {
    const items = [at('main', 60), at('alt', 4), at('mid', 10), at('new'), at('baby', 1)];
    const split = splitByLevel(items, false);
    expect(split.shown.map((item) => item.key)).toEqual(['main', 'mid', 'new']);
    expect(split.hidden).toBe(2);
    expect(split.lowCount).toBe(2);
  });

  it('shows everything once revealed', () => {
    const items = [at('main', 60), at('alt', 4)];
    expect(splitByLevel(items, true)).toEqual({ shown: items, hidden: 0, lowCount: 1 });
  });

  it('always keeps the current character even when it is low level', () => {
    const split = splitByLevel([at('main', 60), at('current', 3), at('alt', 2)], false, ['current']);
    expect(split.shown.map((item) => item.key)).toEqual(['main', 'current']);
    expect(split.hidden).toBe(1);
  });

  it('names the count on the toggle', () => {
    expect(lowLevelCopy.show(3)).toBe('Show low-level characters (3)');
    expect(lowLevelCopy.hide).toBe('Hide low-level characters');
  });
});
