// web/tests/e2e/parity-report.spec.ts
// The report page (/reports/sample on the bundled real log) against its mocks (design/parity.md,
// design/specs/2026-11-07-logs-reports.md section 14): every state at 1440 x 1000, each region's
// blocks and box compared (support/parity.ts). The mocks are rendered before the tests
// (support/parity-mocks.ts renderReportMocks: logs_facts.mjs and ws_report.py) from the same
// log this spec serves. A side-by-side per state lands in test-results/parity/report-<state>.png.
import { test, expect } from './support/test';
import {
  type ParityFailure,
  type RegionSnapshot,
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

/**
 * Known differences (design/parity.md EXPECTED_DRIFT). The chart Section: the mock library's
 * LineChart draws the plot as a static picture (a legend list, one image, an axis), while the
 * shipped chart under it is the interactive one the spec keeps (logs spec 2026-11-07 6.1: "the
 * shipped canvas and brush stay inside"): legend chips that toggle a line, the plot's death
 * marks, the brush's two sliders, the Players control and the log scale in the Section head.
 * Its blocks and its height differ; the regions below it are compared after moving them by the
 * height it adds. Since the panels (directions b1) the Section draws its head and body in one Plate, and
 * that Plate carries the same drift (chartPlates). Lane: REPORT-FRAME, 2026-11-07, until LineChart in ws_prims.py draws the
 * shipped chart's controls.
 */
const CHART_DRIFT = /^section-(damage|healing|damage-taken)-over-the-fight$/;

/** The build's regions below the chart Section, moved up by the height it adds over the mock's. */
function withoutChartHeight(mock: RegionSnapshot[], build: RegionSnapshot[]): RegionSnapshot[] {
  const m = mock.find((r) => CHART_DRIFT.test(r.name));
  const b = build.find((r) => CHART_DRIFT.test(r.name));
  if (m === undefined || b === undefined) return build;
  // Measured at the bottom edge: the regions below land where the mock's chart ends.
  const extra = b.box.y + b.box.height - (m.box.y + m.box.height);
  const bottom = b.box.y + b.box.height;
  // Everything under it in the main column, both halves of the Summary's ability grid among them.
  const inColumn = (r: RegionSnapshot): boolean => r.box.x >= b.box.x && r.box.x < b.box.x + b.box.width;
  return build.map((r) =>
    r.box.y >= bottom && inColumn(r) ? { ...r, box: { ...r.box, y: r.box.y - extra } } : r,
  );
}

/** The chart Section's own panel (Section `panel` draws a Plate around its head and body), on
 *  either side: the same drift as its Section. */
function chartPlates(...sides: RegionSnapshot[][]): Set<string> {
  const plates = new Set<string>();
  for (const side of sides) {
    const chart = side.find((r) => CHART_DRIFT.test(r.name));
    if (chart === undefined) continue;
    for (const r of side)
      if (/^plate(#\d+)?$/.test(r.name) && r.box.x === chart.box.x && r.box.y === chart.box.y)
        plates.add(r.name);
  }
  return plates;
}

function unexpected(failures: ParityFailure[], plates: ReadonlySet<string>): ParityFailure[] {
  return failures.filter(
    (f) =>
      !CHART_DRIFT.test(f.region) &&
      !plates.has(f.region) &&
      !(f.region === 'site-footer' && f.kind === 'box'),
  );
}

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

      const failures = unexpected(
        compareRegions(mock, withoutChartHeight(mock, built)),
        chartPlates(mock, built),
      );
      expect(failures, `${state.id} differs from ${state.mock}:\n${formatFailures(failures)}`).toEqual([]);
    });
  }
});
