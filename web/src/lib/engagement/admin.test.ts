import { describe, expect, it } from 'vitest';
import type { EngagementPeriod, EngagementReport } from './api';
import {
  adminActivityCopy as copy,
  cohortRows,
  dayRows,
  floorLabel,
  headlineFigures,
  weekRows,
} from './admin';

const period = (start: string, figures: number[], partial = false): EngagementPeriod => ({
  start,
  active: figures[0],
  new: figures[1],
  returning: figures[0] - figures[1],
  site: figures[2],
  addon: figures[3],
  app: figures[4],
  partial,
});

const REPORT: EngagementReport = {
  today: '2026-11-13',
  live_since: '2026-11-13',
  internal_accounts: 6,
  days: [period('2026-11-13', [12, 3, 10, 4, 1]), period('2026-11-12', [9, 2, 8, 3, 0], true)],
  weeks: [period('2026-11-09', [40, 11, 33, 15, 2], true), period('2026-11-02', [52, 30, 44, 20, 0], true)],
  cohorts: [
    {
      week: '2026-10-26',
      joined: 20,
      weeks: [
        { active: 9, partial: true },
        { active: 6, partial: true },
      ],
    },
    { week: '2026-11-02', joined: 30, weeks: [{ active: 12, partial: true }] },
    { week: '2026-11-09', joined: 11, weeks: [] },
  ],
};

describe("the owner's activity page", () => {
  it('shows a figure from before the live record as a floor', () => {
    expect(floorLabel(1240, false)).toBe('1,240');
    expect(floorLabel(9, true)).toBe('≥\u00a09');
    expect(floorLabel(0, true)).toBe('0');
  });

  it('draws a day row per day, newest first, today named', () => {
    const rows = dayRows(REPORT.days);
    expect(rows[0]).toEqual({
      key: '2026-11-13',
      label: '13 Nov',
      sub: 'today',
      cells: ['12', '3', '9', '10', '4', '1'].map((text) => ({ tone: 'none', text })),
    });
    expect(rows[1].sub).toBeUndefined();
    expect(rows[1].cells[0].text).toBe('≥\u00a09');
  });

  it('names the weeks by their Monday and marks this one as running', () => {
    const rows = weekRows(REPORT.weeks);
    expect(rows[0].label).toBe('Week of 9 Nov');
    expect(rows[0].sub).toBe(copy.thisWeek);
    expect(rows[1].cells.map((c) => c.text)).toEqual([
      '≥\u00a052',
      '≥\u00a030',
      '≥\u00a022',
      '≥\u00a044',
      '≥\u00a020',
      '0',
    ]);
  });

  it('lays retention out newest signup week first, a share of the joined per later week', () => {
    const rows = cohortRows(REPORT.cohorts);
    expect(rows.map((r) => r.label)).toEqual(['9 Nov', '2 Nov', '26 Oct']);
    expect(rows[2].sub).toBe('20 joined');
    expect(rows[2].cells).toHaveLength(copy.cohortColumns.length);
    expect(rows[2].cells[0]).toEqual({ tone: 'none', text: '≥\u00a045%', label: '9 of 20, ≥\u00a045%' });
    expect(rows[2].cells[2].text).toBe('');
    expect(rows[0].cells.every((c) => c.text === '')).toBe(true);
  });

  it('heads the page with today, the week split new and returning, and last week back', () => {
    expect(headlineFigures(REPORT)).toEqual([
      [
        { value: '12', label: 'active today' },
        { value: '≥\u00a040', label: 'active this week' },
      ],
      [
        { value: '≥\u00a011', label: 'new this week' },
        { value: '≥\u00a029', label: 'returning this week' },
      ],
      [{ value: '40%', label: 'of last week back' }],
    ]);
    expect(headlineFigures({ ...REPORT, days: [], weeks: [] })).toEqual([]);
  });

  it('says what is left out and since when the record is live', () => {
    expect(copy.internal(6)).toBe('Leaves out 6 internal accounts: yours and every granted Premium.');
    expect(copy.internal(1)).toContain('1 internal account:');
    expect(copy.floor('2026-11-13')).toMatch(/^Recorded live since 13 Nov\./);
    expect(copy.floor(null)).toMatch(/^Every figure is rebuilt/);
  });
});
