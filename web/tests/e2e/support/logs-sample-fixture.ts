// web/tests/e2e/support/logs-hero-fixture.ts
// The canonical sample report's meta, which /logs reads on every visit
// to draw the one sample row under "Recent public reports" (lib/logs/sample-row.ts). A spec that
// needs the row to resolve to its real title, date and kills -- rather than the link-only fallback
// an unmocked fetch gives -- routes through here, the same "one fixture, several callers"
// convention active-build.ts and character-fixture.ts already use.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import type { Page } from '@playwright/test';
import type { ReportMeta } from '../../../src/lib/report/types';

const WEB_ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../..');

export const SAMPLE_REPORT_ID: string = (
  JSON.parse(readFileSync(path.join(WEB_ROOT, 'src/data/sample-report.json'), 'utf8')) as { id: string }
).id;

/** The bundled sample's own meta, the one source the page, the /logs row and the landing card read. */
export const SAMPLE_REPORT_META = JSON.parse(
  readFileSync(path.join(WEB_ROOT, `public/sample-report/${SAMPLE_REPORT_ID}/meta.json`), 'utf8'),
) as ReportMeta;

/** Routes the bundled sample's meta, a plain JSON file beside its data (lib/report/sample.ts). */
export async function routeSampleReportMeta(page: Page): Promise<void> {
  await page.route(`**/sample-report/${SAMPLE_REPORT_ID}/meta.json`, (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(SAMPLE_REPORT_META) }),
  );
}
