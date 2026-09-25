import { describe, expect, it } from 'vitest';
import { timelineRows, type DateEntry } from './timeline';

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
    expect(rows[1]).toEqual({ key: 'Sept 17', value: 'Beta opens', note: 'level cap 30', iso: '2026-09-17', status: 'past' });
  });
});
