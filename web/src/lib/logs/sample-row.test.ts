import { describe, expect, it } from 'vitest';
import type { ReportMeta } from '../report/types';
import { SAMPLE_FALLBACK_ROW, SAMPLE_REPORT_ID, sampleRowFromMeta } from './sample-row';

const META = {
  id: SAMPLE_REPORT_ID,
  title: 'Sanguine Depths, sample log',
  zone: 'Sanguine Depths',
  created_at: '2026-09-15T12:55:13Z',
  fights: [
    { index: 1, kind: 'trash', kill: false },
    { index: 2, kind: 'encounter', kill: true },
    { index: 3, kind: 'encounter', kill: false },
    { index: 4, kind: 'encounter', kill: true },
  ],
} as unknown as ReportMeta;

describe('sampleRowFromMeta', () => {
  it('reads its title, zone and boss kills from the report itself', () => {
    expect(sampleRowFromMeta(META)).toEqual({
      id: SAMPLE_REPORT_ID,
      title: 'Sanguine Depths, sample log',
      createdAt: '2026-09-15T12:55:13Z',
      summary: 'Sanguine Depths · 2 bosses · 4 fights',
    });
  });

  it('falls back to the zone when the report has no title', () => {
    expect(sampleRowFromMeta({ ...META, title: '' }).title).toBe('Sanguine Depths');
  });

  it('has a link-only fallback for a report that cannot be read', () => {
    expect(SAMPLE_FALLBACK_ROW.id).toBe(SAMPLE_REPORT_ID);
    expect(SAMPLE_FALLBACK_ROW.summary).toBe('');
  });
});
