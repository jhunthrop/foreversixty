import { describe, expect, it } from 'vitest';
import { controlSummary, isClassSlug, rowElement, signedFigure, titleColor } from './primitives';

describe('titleColor', () => {
  it('paints a class page in its class colour, a guild in its faction colour, the rest strong', () => {
    expect(titleColor({ classSlug: 'hunter' })).toBe('var(--color-class-hunter)');
    expect(titleColor({ faction: 'horde' })).toBe('var(--color-horde)');
    expect(titleColor('strong')).toBe('var(--color-cream)');
  });

  it('knows the nine classes', () => {
    expect(isClassSlug('druid')).toBe(true);
    expect(isClassSlug('monk')).toBe(false);
  });
});

describe('controlSummary', () => {
  it("joins each group's chosen label with a middle dot", () => {
    const groups = [
      {
        label: 'Spec',
        param: 'spec',
        kind: 'segmented',
        value: 'bm',
        options: [{ value: 'bm', label: 'Beast Mastery' }],
      },
      {
        label: 'Level',
        param: 'band',
        kind: 'segmented',
        value: '20',
        options: [{ value: '20', label: '20-29' }],
      },
      {
        label: 'Race',
        param: 'race',
        kind: 'select',
        value: 'x',
        options: [{ value: 'dwarf', label: 'Dwarf' }],
      },
    ] as const;
    expect(controlSummary(groups)).toBe('Beast Mastery · 20-29');
  });
});

describe('rowElement', () => {
  it('is a link only with an href', () => {
    expect(rowElement('/x')).toBe('a');
    expect(rowElement(undefined)).toBe('div');
    expect(rowElement('')).toBe('div');
    expect(rowElement(undefined, true)).toBe('li');
    expect(rowElement('/x', true)).toBe('a');
  });
});

describe('signedFigure', () => {
  it('signs a gain with a plus and a loss with a true minus, and prints zero bare', () => {
    expect(signedFigure(1.84)).toBe('+1.8');
    expect(signedFigure(-1.84)).toBe('−1.8');
    expect(signedFigure(0.04)).toBe('0.0');
    expect(signedFigure(-0.04)).toBe('0.0');
  });
});
