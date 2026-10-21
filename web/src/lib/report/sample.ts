// web/src/lib/report/sample.ts
// The site's one sample report lives in the site, not the database: its meta and data files
// are static assets under /sample-report/<id>/ (web/public/sample-report), so a database
// reset or migration cannot take the link away. The island reads the meta from the bundle
// instead of GET /v1/reports/{id}; the data files are read from data_base_url as usual.
import sampleReport from '../../data/sample-report.json';

export const SAMPLE_REPORT_ID: string = sampleReport.id;
/** Ids the sample answered to before it was bundled; the Worker redirects them to the page. */
export const SAMPLE_LEGACY_IDS: readonly string[] = sampleReport.legacyIds;
export const SAMPLE_REPORT_PATH = `/reports/${SAMPLE_REPORT_ID}`;
export const SAMPLE_META_URL = `/sample-report/${SAMPLE_REPORT_ID}/meta.json`;

export function isSampleReport(id: string): boolean {
  return id === SAMPLE_REPORT_ID;
}
