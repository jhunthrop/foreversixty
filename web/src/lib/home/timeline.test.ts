import { describe, expect, it } from 'vitest';
import { timelineDisplayRows, timelineRowText, timelineRows, type DateEntry } from './timeline';

const DATES: DateEntry[] = [
  { key: 'Sept 13', value: 'Deep Dive panel', iso: '2026-09-13' },
  { key: 'Sept 17', value: 'Beta opens', note: 'level cap 30', iso: '2026-09-17' },
  { key: 'Oct 27', value: 'Name reservation', iso: '2026-10-27' },
  { key: 'Nov 4', value: 'Launch', iso: '2026-11-04' },
  { key: 'Dec 9', value: 'First raids', iso: '2026-12-09' },
];

describe('timelineRows', () => {
  it('marks every date before now past, the first date at or after now next, and the rest future', () => {
    const rows = timelineRows(DATES, new Date('2026-09-24T12:00:00Z'));
    expect(rows.map((r) => r.status)).toEqual(['past', 'past', 'next', 'future', 'future']);
  });

  it('marks every date past when now is after the last one', () => {
    const rows = timelineRows(DATES, new Date('2027-01-01T00:00:00Z'));
    expect(rows.map((r) => r.status)).toEqual(['past', 'past', 'past', 'past', 'past']);
  });

  it('marks every date next-or-future when now is before the first one', () => {
    const rows = timelineRows(DATES, new Date('2026-01-01T00:00:00Z'));
    expect(rows.map((r) => r.status)).toEqual(['next', 'future', 'future', 'future', 'future']);
  });

  it('preserves every field of the source row', () => {
    const rows = timelineRows(DATES, new Date('2026-09-24T12:00:00Z'));
    expect(rows[1]).toEqual({
      key: 'Sept 17',
      value: 'Beta opens',
      note: 'level cap 30',
      iso: '2026-09-17',
      status: 'past',
    });
  });

  describe('a span (Beta, Sept 17 to Oct 21)', () => {
    const SPAN: DateEntry[] = [
      { key: 'Sept 13', value: 'Deep Dive panel', iso: '2026-09-13' },
      {
        key: 'Sept 17',
        value: 'Beta opens',
        note: 'level cap 30 · through Oct 21',
        iso: '2026-09-17',
        endIso: '2026-10-21',
        pastKey: 'Sept 17 to Oct 21',
        pastValue: 'Beta',
        pastNote: 'level cap 30',
      },
      { key: 'Oct 27', value: 'Name reservation', iso: '2026-10-27' },
      { key: 'Nov 4', value: 'Launch', iso: '2026-11-04' },
    ];

    it('is live while the beta runs, including all of its last day, with the next date after it', () => {
      for (const now of ['2026-09-17T12:00:00Z', '2026-10-21T23:00:00Z']) {
        const rows = timelineRows(SPAN, new Date(now));
        expect(rows.map((r) => r.status)).toEqual(['past', 'live', 'next', 'future']);
      }
    });

    it('is past once its last day is over, and then reads in the past tense', () => {
      const rows = timelineRows(SPAN, new Date('2026-10-22T08:00:00Z'));
      expect(rows.map((r) => r.status)).toEqual(['past', 'past', 'next', 'future']);
      expect(timelineRowText(rows[1])).toEqual({
        key: 'Sept 17 to Oct 21',
        value: 'Beta',
        note: 'level cap 30',
      });
    });

    it('reads in the present tense while live', () => {
      const rows = timelineRows(SPAN, new Date('2026-10-20T08:00:00Z'));
      expect(timelineRowText(rows[1]).value).toBe('Beta opens');
    });

    it('shows live, next and future first and the past dates last, each group in date order', () => {
      const rows = timelineRows(SPAN, new Date('2026-10-22T08:00:00Z'));
      expect(timelineDisplayRows(rows).map((r) => r.value)).toEqual([
        'Name reservation',
        'Launch',
        'Deep Dive panel',
        'Beta opens',
      ]);
      const live = timelineRows(SPAN, new Date('2026-10-20T08:00:00Z'));
      expect(timelineDisplayRows(live).map((r) => r.status)).toEqual(['live', 'next', 'future', 'past']);
    });
  });
});
