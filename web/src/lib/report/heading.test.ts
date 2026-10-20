import { describe, expect, it } from 'vitest';
import { reportHeading } from './heading';

const fights = (count: number) => Array.from({ length: count }, () => ({ start: '2026-09-17T20:10:04Z' }));

describe('reportHeading', () => {
  it('is the upload’s title when it has one, then the zone', () => {
    expect(
      reportHeading({ title: 'Friday raid', zone: 'Sanguine Depths', created_at: '', fights: fights(2) }),
    ).toBe('Friday raid');
    expect(reportHeading({ title: '', zone: 'Sanguine Depths', created_at: '', fights: fights(2) })).toBe(
      'Sanguine Depths',
    );
  });
  it('is the log’s date and fight count when there is neither', () => {
    const heading = reportHeading({
      title: '',
      zone: '',
      created_at: '2026-10-19T00:00:00Z',
      fights: fights(42),
    });
    expect(heading).toMatch(/^17 Sept? 2026 · 42 fights$/);
  });
  it('falls back to the upload date for a report with no fights yet', () => {
    expect(reportHeading({ title: '', zone: '', created_at: '2026-10-19T00:00:00Z', fights: [] })).toMatch(
      /^19 Oct 2026 · 0 fights$/,
    );
  });
  it('leads with the logging character when the report is credited, and only then', () => {
    const meta = { title: '', zone: '', created_at: '2026-10-20T09:00:00Z', fights: [] };
    expect(reportHeading(meta, 'Bow Jackzon')).toMatch(/^Bow Jackzon · 20 Oct 2026 · 0 fights$/);
    expect(reportHeading(meta)).toMatch(/^20 Oct 2026 · 0 fights$/);
    expect(reportHeading({ ...meta, title: 'Friday raid' }, 'Bow Jackzon')).toBe('Friday raid');
  });
});
