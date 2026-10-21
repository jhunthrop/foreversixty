// web/src/lib/logs/sample-row.ts
// The one sample report, as the plain row it is in "Recent public reports"
// (design/specs/2026-10-11-logs.md sections 2.1 and 9): the site's canonical public report,
// bundled with the site and read from its static meta (lib/report/sample.ts), shown once and labelled `A sample
// report` above its row. It never doubles as a feed entry (the feed filters it out by id).
import { formatDate } from '../report/format';
import { reportFacts } from '../report/heading';
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
  /** The date as the report page says it ("26 Sep 2026"), so the row and the page agree. */
  dateLabel: string;
  summary: string;
}

/** The report page dates a report by its first fight; the row reads the same field. */
export function sampleDateIso(meta: Pick<ReportMeta, 'fights' | 'created_at'>): string {
  return meta.fights[0]?.start ?? meta.created_at;
}

export function sampleRowFromMeta(meta: ReportMeta): SampleRow {
  const killCount = meta.fights.filter((fight) => fight.kind === 'encounter' && fight.kill).length;
  const createdAt = sampleDateIso(meta);
  return {
    id: meta.id,
    title: meta.title === '' ? meta.zone : meta.title,
    createdAt,
    dateLabel: reportFacts(meta).date || formatDate(createdAt),
    summary: recentReportSummary({ zone: meta.zone, killCount, fightCount: meta.fights.length }),
  };
}

/** The link alone: a sample that cannot be read is still one click from opening. */
export const SAMPLE_FALLBACK_ROW: SampleRow = {
  id: SAMPLE_REPORT_ID,
  title: FALLBACK_TITLE,
  createdAt: '',
  dateLabel: '',
  summary: '',
};

export async function loadSampleRow(): Promise<SampleRow> {
  try {
    return sampleRowFromMeta(await fetchReportMeta(SAMPLE_REPORT_ID));
  } catch {
    return SAMPLE_FALLBACK_ROW;
  }
}
