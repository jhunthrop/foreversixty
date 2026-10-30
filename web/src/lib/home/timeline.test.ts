import { describe, expect, it } from 'vitest';
import { datesLineItems, timelineRows, type DateEntry } from './timeline';

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

    it('is past once its last day is over', () => {
      const rows = timelineRows(SPAN, new Date('2026-10-22T08:00:00Z'));
      expect(rows.map((r) => r.status)).toEqual(['past', 'past', 'next', 'future']);
    });
  });

  describe('a single day', () => {
    it('stays live through its own day and drops off the day after', () => {
      const day = [{ key: 'Nov 4', value: 'Launch', iso: '2026-11-04' }];
      expect(timelineRows(day, new Date('2026-11-04T22:00:00Z'))[0].status).toBe('live');
      expect(timelineRows(day, new Date('2026-11-05T00:30:00Z'))[0].status).toBe('past');
    });
  });
});

describe('datesLineItems', () => {
  const LINE: DateEntry[] = [
    { key: 'Sept 17', value: 'Beta opens', iso: '2026-09-17', endIso: '2026-10-21' },
    { key: 'Oct 27', value: 'Name reservation', iso: '2026-10-27', endIso: '2026-11-03' },
    { key: 'Nov 4', value: 'Launch', iso: '2026-11-04', time: '15:00 PST', at: '2026-11-04T23:00:00Z' },
    { key: 'Dec 9', value: 'First raids', iso: '2026-12-09', note: 'Barrow Deeps' },
  ];
  const itemsAt = (iso: string) => {
    const now = new Date(iso);
    return datesLineItems(timelineRows(LINE, now), now);
  };

  it('lists only what is still to come, the first as the next, with the time where there is one', () => {
    expect(itemsAt('2026-10-29T12:00:00Z')).toEqual([
      { text: 'Launch Nov 4 · 15:00 PST', next: true },
      { text: 'First raids Dec 9', next: false },
    ]);
  });

  it('leads with a span that has not started yet', () => {
    expect(itemsAt('2026-10-23T12:00:00Z').map((i) => i.text)).toEqual([
      'Name reservation Oct 27',
      'Launch Nov 4 · 15:00 PST',
      'First raids Dec 9',
    ]);
  });

  it('keeps the launch until its own hour, then drops it', () => {
    expect(itemsAt('2026-11-04T22:59:00Z')[0]?.text).toBe('Launch Nov 4 · 15:00 PST');
    expect(itemsAt('2026-11-04T23:01:00Z')).toEqual([{ text: 'First raids Dec 9', next: true }]);
  });

  it('is empty once every date has passed', () => {
    expect(itemsAt('2026-12-10T12:00:00Z')).toEqual([]);
  });
});
