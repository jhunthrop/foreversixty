// web/src/lib/chart/palette.test.ts
import { describe, expect, it } from 'vitest';
import { entityStyle, withAlpha } from './palette';

describe('entityStyle', () => {
  it('takes the palette in order and dashes the seventh line', () => {
    expect(entityStyle(0)).toEqual({ colour: 'var(--color-gold)' });
    expect(entityStyle(1).colour).toBe('var(--color-school-frost)');
    expect(entityStyle(6)).toEqual({ colour: 'var(--color-gold)', dash: [4, 3] });
  });
});

describe('withAlpha', () => {
  it('turns a hex or rgb colour into rgba', () => {
    expect(withAlpha('#e5b955', 0.2)).toBe('rgba(229, 185, 85, 0.2)');
    expect(withAlpha('rgb(1, 2, 3)', 0.5)).toBe('rgba(1, 2, 3, 0.5)');
    expect(withAlpha('red', 0.5)).toBe('red');
  });
});
