// web/scripts/sync-report-fixture.mjs
// Copies the checked-in report fixture to public/logs-data/reports/<id>/ so the browser
// fetches it from exactly the path the Worker serves real reports from. Under
// FOREVER_DATA=real this does nothing: production has no fixture report, and
// src/pages/reports/[id].astro prerenders no fixture pages either.
//
// Mirrors scripts/sync-data.mjs: same selector, same "write into public/, gitignored"
// shape, run from the same pre* hooks.
import { cp, mkdir, rm } from 'node:fs/promises';
import path from 'node:path';
import process from 'node:process';
import { fileURLToPath } from 'node:url';

const FIXTURE_REPORT_ID = 'fixture2abcd';

const webRoot = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '..');
const target = path.join(webRoot, 'public/logs-data/reports', FIXTURE_REPORT_ID);

await rm(path.join(webRoot, 'public/logs-data'), { recursive: true, force: true });

if (process.env.FOREVER_DATA !== 'fixture') {
  console.log('report fixture: FOREVER_DATA is not fixture, nothing published');
} else {
  const source = path.join(webRoot, 'src/fixtures/report');
  await mkdir(target, { recursive: true });
  await cp(path.join(source, 'report.json'), path.join(target, 'report.json'));
  await cp(path.join(source, 'fights'), path.join(target, 'fights'), { recursive: true });
  console.log(`report fixture published to public/logs-data/reports/${FIXTURE_REPORT_ID}`);
}
