import { describe, expect, it } from 'vitest';
import type { ReportMeta } from '../report/types';
import { formatDate } from '../report/format';
import { SAMPLE_FALLBACK_ROW, SAMPLE_REPORT_ID, sampleDateIso, sampleRowFromMeta } from './sample-row';

const META = {
  id: SAMPLE_REPORT_ID,
  title: 'Sanguine Depths, sample log',
  zone: 'Sanguine Depths',
  created_at: '2026-09-15T12:55:13Z',
  fights: [
    { index: 1, kind: 'trash', kill: false, start: '2026-09-26T20:10:04Z' },
    { index: 2, kind: 'encounter', kill: true, start: '2026-09-26T20:12:00Z' },
    { index: 3, kind: 'encounter', kill: false, start: '2026-09-26T20:13:00Z' },
    { index: 4, kind: 'encounter', kill: true, start: '2026-09-26T20:14:00Z' },
  ],
} as unknown as ReportMeta;

describe('sampleRowFromMeta', () => {
  it('reads its title, zone and boss kills from the report itself', () => {
    expect(sampleRowFromMeta(META)).toEqual({
      id: SAMPLE_REPORT_ID,
      title: 'Sanguine Depths, sample log',
      createdAt: '2026-09-26T20:10:04Z',
      dateLabel: '26 Sept 2026',
      summary: 'Sanguine Depths · 2 bosses · 4 fights',
    });
  });

  it('says the date the report page says: the first fight, in the page format', () => {
    expect(sampleDateIso(META)).toBe('2026-09-26T20:10:04Z');
    expect(sampleRowFromMeta(META).dateLabel).toBe(formatDate(META.fights[0].start));
    // A meta with no fights falls back to its own created_at.
    expect(sampleRowFromMeta({ ...META, fights: [] } as unknown as ReportMeta).createdAt).toBe(
      '2026-09-15T12:55:13Z',
    );
  });

  it('counts what the page counts: the same bosses and fights as its header', () => {
    expect(sampleRowFromMeta(META).summary).toBe(`${META.zone} · 2 bosses · ${META.fights.length} fights`);
  });

  it('falls back to the zone when the report has no title', () => {
    expect(sampleRowFromMeta({ ...META, title: '' }).title).toBe('Sanguine Depths');
  });

  it('has a link-only fallback for a report that cannot be read', () => {
    expect(SAMPLE_FALLBACK_ROW.id).toBe(SAMPLE_REPORT_ID);
    expect(SAMPLE_FALLBACK_ROW.summary).toBe('');
  });
});
