// web/scripts/make-report-fixture.mjs
// Regenerates src/fixtures/report/ by running the engine's CLI over
// src/fixtures/report/fixture.log. The CLI writes store.Keys paths under its -out
// directory (reports/<id>/…), so the output is moved up one level into the fixture
// directory, where the site's own /logs-data/reports/<id>/ layout starts at `fights/`.
//
// There is no go.work at the repository root yet (the api plan adds one), so the CLI runs
// from logs/ where its module lives.
import { execFileSync } from 'node:child_process';
import { cp, mkdtemp, rm } from 'node:fs/promises';
import { tmpdir } from 'node:os';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

export const FIXTURE_REPORT_ID = 'fixture2abcd';

const webRoot = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '..');
const repoRoot = path.resolve(webRoot, '..');
const fixtureDir = path.join(webRoot, 'src/fixtures/report');

const staging = await mkdtemp(path.join(tmpdir(), 'forever-report-fixture-'));
try {
  execFileSync(
    'go',
    [
      'run',
      './cmd/forever-logs',
      'parse',
      '-out',
      staging,
      '-report',
      FIXTURE_REPORT_ID,
      path.join(fixtureDir, 'fixture.log'),
    ],
    { cwd: path.join(repoRoot, 'logs'), stdio: 'inherit' },
  );
  const produced = path.join(staging, 'reports', FIXTURE_REPORT_ID);
  await rm(path.join(fixtureDir, 'fights'), { recursive: true, force: true });
  await cp(path.join(produced, 'fights'), path.join(fixtureDir, 'fights'), { recursive: true });
  await cp(path.join(produced, 'report.json'), path.join(fixtureDir, 'report.json'));
} finally {
  await rm(staging, { recursive: true, force: true });
}
console.log(`report fixture written to ${fixtureDir}`);
