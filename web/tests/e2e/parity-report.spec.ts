// web/tests/e2e/parity-report.spec.ts
// The report page (/reports/sample on the bundled real log) against its mocks (design/parity.md,
// design/specs/2026-11-07-logs-reports.md section 14): every state at 1440 x 1000, each region's
// blocks and box compared (support/parity.ts). The mocks are rendered before the tests
// (support/parity-mocks.ts renderReportMocks: logs_facts.mjs and ws_report.py) from the same
// log this spec serves. A side-by-side per state lands in test-results/parity/report-<state>.png.
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
import { renderReportMocks } from './support/parity-mocks';
import { openReportState, type ReportParityState } from './support/report-parity';

const CONFIG = readParityConfig('report-parity.json');

test.describe('Report mock parity at 1440', () => {
  test.skip(({ isMobile }) => isMobile, 'the parity mocks are desktop captures');
  test.use({ viewport: CONFIG.viewport, timezoneId: 'UTC' });

  let mocks = '';
  // eslint-disable-next-line no-empty-pattern -- Playwright reads a hook's fixtures from its pattern
  test.beforeAll(({}, testInfo) => {
    test.skip(testInfo.project.name !== 'desktop', 'the parity mocks are desktop captures');
    mocks = renderReportMocks(CONFIG.states.map((state) => state.mock.replace(/\.html$/, '')));
  });

  for (const state of CONFIG.states) {
    test(`${state.id} matches ${state.mock}`, async ({ page, browser }) => {
      test.slow();
      const mockPage = await browser.newPage({ viewport: CONFIG.viewport });
      await openMock(mockPage, mocks, state.mock);
      const mock = await captureRegions(mockPage);

      await prepareParityPage(page);
      await openReportState(page, state.id as ReportParityState, state.path);
      await settleParityPage(page);
      const built = await captureRegions(page);

      await writeSideBySide(
        browser,
        `report-${state.id}`,
        await captureOutlined(mockPage, mock),
        await captureOutlined(page, built),
      );
      await mockPage.close();

      const failures = compareRegions(mock, built);
      expect(failures, `${state.id} differs from ${state.mock}:\n${formatFailures(failures)}`).toEqual([]);
    });
  }
});
