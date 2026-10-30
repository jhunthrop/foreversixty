// web/src/lib/chart/copy.test.ts
import { describe, expect, it } from 'vitest';
import { chartCopy } from './copy';

describe('chartCopy.playersField (fix round 2: never a limit above the roster)', () => {
  it('reads All N when N is 8 or fewer, whatever the preset', () => {
    expect(chartCopy.playersField(3, 3, 'split')).toBe('All 3');
    expect(chartCopy.playersField(8, 8, 'top')).toBe('All 8');
  });

  it('names the preset over 8, and the count once the reader picks', () => {
    expect(chartCopy.playersField(8, 40, 'split')).toBe('Top 4, bottom 4 of 40');
    expect(chartCopy.playersField(8, 40, 'top')).toBe('Top 8 of 40');
    expect(chartCopy.playersField(8, 40, 'bottom')).toBe('Bottom 8 of 40');
    expect(chartCopy.playersField(10, 40, 'custom')).toBe('10 of 40');
    expect(chartCopy.playersField(2, 3, 'custom')).toBe('2 of 3');
  });
});

describe('the threat title and the Enemy field agree (fix round 2)', () => {
  it('names every enemy or the picked one, at every width', () => {
    expect(chartCopy.threatTitle(null)).toBe('Estimated threat, every enemy');
    expect(chartCopy.threatTitle('Ragnaros')).toBe('Estimated threat on Ragnaros');
    expect(chartCopy.tooltipRight('Ragnaros', chartCopy.tooltipThreat(chartCopy.threatSubject(null)))).toBe(
      'Ragnaros phase · estimated threat, every enemy',
    );
  });
});

describe('chartCopy.tooltipRate', () => {
  it('names the average the points are', () => {
    expect(chartCopy.tooltipRate('damage-done', 5)).toBe('DPS, 5 s average');
    expect(chartCopy.tooltipRate('healing', 2)).toBe('HPS, 2 s average');
    expect(chartCopy.tooltipRate('damage-done', 1)).toBe('DPS, each second');
  });
});
