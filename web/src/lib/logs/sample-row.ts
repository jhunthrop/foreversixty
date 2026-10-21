// web/src/lib/logs/sample-row.ts
// The one sample report, as the plain row it is in "Recent public reports"
// (design/specs/2026-10-11-logs.md sections 2.1 and 9): the site's canonical public report,
// bundled with the site and read from its static meta (lib/report/sample.ts), shown once and labelled `A sample
// report` above its row. It never doubles as a feed entry (the feed filters it out by id).
import { fetchReportMeta } from '../report/load';
import { SAMPLE_REPORT_ID } from '../report/sample';
import type { ReportMeta } from '../report/types';
import { recentReportSummary } from './report-summary';

export { SAMPLE_REPORT_ID };
export const SAMPLE_LABEL = 'A sample report';
const FALLBACK_TITLE = 'Open the sample report';

export interface SampleRow {
  id: string;
  title: string;
  /** Empty when the report could not be read: the row is then just the link. */
  createdAt: string;
  summary: string;
}

export function sampleRowFromMeta(meta: ReportMeta): SampleRow {
  const killCount = meta.fights.filter((fight) => fight.kind === 'encounter' && fight.kill).length;
  return {
    id: meta.id,
    title: meta.title === '' ? meta.zone : meta.title,
    createdAt: meta.created_at,
    summary: recentReportSummary({ zone: meta.zone, killCount, fightCount: meta.fights.length }),
  };
}

/** The link alone: a sample that cannot be read is still one click from opening. */
export const SAMPLE_FALLBACK_ROW: SampleRow = {
  id: SAMPLE_REPORT_ID,
  title: FALLBACK_TITLE,
  createdAt: '',
  summary: '',
};

export async function loadSampleRow(): Promise<SampleRow> {
  try {
    return sampleRowFromMeta(await fetchReportMeta(SAMPLE_REPORT_ID));
  } catch {
    return SAMPLE_FALLBACK_ROW;
  }
}
