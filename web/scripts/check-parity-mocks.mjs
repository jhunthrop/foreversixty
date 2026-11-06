// web/scripts/check-parity-mocks.mjs
// After the parity specs: are the committed reference mocks (design/mocks/web) behind what the
// generators rendered from the served build? Exit 1 with the stale files and the regen commands,
// 0 when they match. CI runs it as a report, not a gate (design/parity.md, "The reference mocks"):
// on GitHub Actions it also writes a warning annotation and the job summary.
//   npx playwright test tests/e2e/parity-*.spec.ts --project=desktop && node scripts/check-parity-mocks.mjs
import { appendFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { formatStaleReport, PARITY_PAGES, staleMocks } from './parity-mocks-stale.mjs';

const WEB_ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '..');
const RENDERED = path.join(WEB_ROOT, 'test-results/parity/mocks');
const COMMITTED = path.join(WEB_ROOT, '..', 'design/mocks/web');

const stale = Object.fromEntries(
  Object.entries(PARITY_PAGES).map(([page, { prefix }]) => [
    page,
    staleMocks(path.join(RENDERED, page), COMMITTED, prefix),
  ]),
);
const report = formatStaleReport(stale);

if (report === '') {
  console.log('check-parity-mocks: the committed mocks match the build');
  process.exit(0);
}
console.error(`check-parity-mocks:\n${report}`);
if (process.env.GITHUB_ACTIONS === 'true') {
  const files = Object.values(stale)
    .flat()
    .map((f) => f.file)
    .join(', ');
  console.log(
    `::warning title=Parity reference mocks are stale::${files} differ from the build; the gate used test-time mocks. Regenerate (job summary has the commands).`,
  );
  if (process.env.GITHUB_STEP_SUMMARY) {
    appendFileSync(
      process.env.GITHUB_STEP_SUMMARY,
      `### Parity reference mocks are stale\n\n\`\`\`\n${report}\n\`\`\`\n`,
    );
  }
}
process.exit(1);
