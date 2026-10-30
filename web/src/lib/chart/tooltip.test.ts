// web/src/lib/chart/tooltip.test.ts
import { describe, expect, it } from 'vitest';
import { tooltipLines } from './tooltip';
import type { ChartSeries } from './types';

const line = (id: string, value: number | null, diedAtMs?: number): ChartSeries => ({
  id,
  label: id,
  colour: 'var(--color-class-mage)',
  kind: 'player',
  values: [value, value, value],
  figure: value ?? 0,
  diedAtMs,
});

const format = (value: number): string => String(Math.round(value));

describe('tooltipLines', () => {
  it('sorts highest first and puts the dead last as Died m:ss', () => {
    const { lines } = tooltipLines([line('a', 10), line('b', null, 1000), line('c', 30)], 2000, {
      hidden: new Set(),
      highlightId: null,
      format,
    });
    expect(lines.map((entry) => entry.id)).toEqual(['c', 'a', 'b']);
    expect(lines[2].valueText).toBe('Died 0:01');
    expect(lines[2].dead).toBe(true);
  });

  it('caps at 12 rows, counts the rest, and always keeps the hovered series', () => {
    const many = Array.from({ length: 20 }, (_, index) => line(`p${index}`, 100 - index));
    const { lines, more } = tooltipLines(many, 0, { hidden: new Set(), highlightId: 'p19', format });
    expect(lines).toHaveLength(12);
    expect(more).toBe(8);
    expect(lines.some((entry) => entry.id === 'p19' && entry.highlighted)).toBe(true);
  });

  it('leaves out a series the legend turned off', () => {
    const { lines } = tooltipLines([line('a', 1), line('b', 2)], 0, {
      hidden: new Set(['b']),
      highlightId: null,
      format,
    });
    expect(lines.map((entry) => entry.id)).toEqual(['a']);
  });
});
