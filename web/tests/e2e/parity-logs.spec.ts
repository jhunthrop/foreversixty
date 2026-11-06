// web/tests/e2e/parity-logs.spec.ts
// /logs, /reports, /rankings, /rankings/<slug> and /ratings against their mocks (design/parity.md,
// design/specs/2026-11-07-logs-reports.md section 14): every state at 1440 x 1000, each region's
// blocks and box compared (support/parity.ts). The mocks are rendered before the tests
// (support/parity-mocks.ts: logs_facts.mjs and ws_logs.py) from the files this spec serves
// (support/logs-parity.ts): web/src/fixtures/reports, the rankings stubs and the bundled sample.
// A side-by-side per state lands in test-results/parity/logs-<state>.png. Desktop only.
import { test, expect } from './support/test';
import {
  captureOutlined,
  captureRegions,
  compareRegions,
  formatFailures,
  openMock,
  prepareParityPage,
  readParityConfig,
  settleParityPage,
  writeSideBySide,
} from './support/parity';
import { renderLogsMocks } from './support/parity-mocks';
import { openLogsState, type LogsParityState } from './support/logs-parity';

const CONFIG = readParityConfig('logs-parity.json');

test.describe('Logs, reports list, rankings and ratings mock parity at 1440', () => {
  test.skip(({ isMobile }) => isMobile, 'the parity mocks are desktop captures');
  test.use({ viewport: CONFIG.viewport });

  let mocks = '';
  // eslint-disable-next-line no-empty-pattern -- Playwright reads a hook's fixtures from its pattern
  test.beforeAll(({}, testInfo) => {
    test.skip(testInfo.project.name !== 'desktop', 'the parity mocks are desktop captures');
    mocks = renderLogsMocks();
  });

  for (const state of CONFIG.states) {
    test(`${state.id} matches ${state.mock}`, async ({ page, browser }) => {
      test.slow();
      const mockPage = await browser.newPage({ viewport: CONFIG.viewport });
      await openMock(mockPage, mocks, state.mock);
      const mock = await captureRegions(mockPage);

      await prepareParityPage(page);
      await openLogsState(page, state.id as LogsParityState);
      await settleParityPage(page);
      const built = await captureRegions(page);

      await writeSideBySide(
        browser,
        `logs-${state.id}`,
        await captureOutlined(mockPage, mock),
        await captureOutlined(page, built),
      );
      await mockPage.close();

      const failures = compareRegions(mock, built);
      expect(failures, `${state.id} differs from ${state.mock}:\n${formatFailures(failures)}`).toEqual([]);
    });
  }
});
