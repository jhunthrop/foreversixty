// web/tests/e2e/parity-connect.spec.ts
// /connect against its mocks (design/parity.md, design/specs/2026-11-07-setup.md section 12):
// signed out, pending and success at 1440 and 390, each region (head, head-title, head-meta,
// plate) compared for its blocks and box (support/parity.ts). The mocks are rendered before the
// tests from the copy modules and the fixtures this spec serves (support/connect-page.ts).
// A side-by-side per state and width lands in test-results/parity/connect-<state>-<width>.png.
// Desktop project only: the 390 run is a 390 px desktop viewport, as the mock renders are.
import { test, expect } from './support/test';
import {
  captureOutlined,
  captureRegions,
  compareRegions,
  formatFailures,
  openMock,
  prepareParityPage,
  settleParityPage,
  writeSideBySide,
} from './support/parity';
import { openConnectState, renderConnectMocks, type ConnectPageState } from './support/connect-page';

const STATES: readonly { id: string; state: ConnectPageState; mock: string }[] = [
  { id: 'connect', state: 'signed-out', mock: 'web-connect' },
  { id: 'connect-pending', state: 'pending', mock: 'web-connect-pending' },
  { id: 'connect-success', state: 'success', mock: 'web-connect-success' },
];
const WIDTHS = [1440, 390] as const;
const HEIGHT = 1000;

/**
 * Known differences the setup lane owns (design/parity.md EXPECTED_DRIFT), keyed `<id>-<width>`:
 * each lists the `region:kind` failures expected; every other region is still compared, and the
 * test fails once a listed failure is gone so its entry is deleted.
 */
const MOCK_HEAD_INSET =
  'below 768 the mock puts the 18 px page inset outside the head region (.pc), PageHeader inside it; ' +
  'a primitive/mock-library question for the setup lane (2026-11-07)';
const EXPECTED_DRIFT: Readonly<Record<string, { failures: readonly string[]; why: string }>> = {
  'connect-390': { failures: ['head:box'], why: MOCK_HEAD_INSET },
  'connect-pending-390': { failures: ['head:box'], why: MOCK_HEAD_INSET },
  'connect-success-390': { failures: ['head:box'], why: MOCK_HEAD_INSET },
};

test.describe('/connect mock parity', () => {
  test.skip(({ isMobile }) => isMobile, 'the parity mocks are desktop captures');

  let mocks = '';
  // eslint-disable-next-line no-empty-pattern -- Playwright reads a hook's fixtures from its pattern
  test.beforeAll(({}, testInfo) => {
    test.skip(testInfo.project.name !== 'desktop', 'the parity mocks are desktop captures');
    mocks = renderConnectMocks(STATES.map((state) => state.mock));
  });

  for (const width of WIDTHS) {
    for (const { id, state, mock: mockName } of STATES) {
      test(`${id} at ${width} matches ${mockName}`, async ({ page, browser }) => {
        test.slow();
        const viewport = { width, height: HEIGHT };
        await page.setViewportSize(viewport);
        const mockPage = await browser.newPage({ viewport });
        await openMock(mockPage, mocks, `${mockName}.html`);
        const mock = await captureRegions(mockPage);

        await prepareParityPage(page);
        await openConnectState(page, state);
        await settleParityPage(page);
        const built = await captureRegions(page);

        await writeSideBySide(
          browser,
          `${id}-${width}`,
          await captureOutlined(mockPage, mock),
          await captureOutlined(page, built),
        );
        await mockPage.close();

        const drift = EXPECTED_DRIFT[`${id}-${width}`];
        const failures = compareRegions(mock, built);
        const expected = failures.filter((f) => drift?.failures.includes(`${f.region}:${f.kind}`));
        const unexpected = failures.filter((f) => !expected.includes(f));
        expect(
          unexpected,
          `${id} at ${width} differs from ${mockName}:\n${formatFailures(unexpected)}`,
        ).toEqual([]);
        if (drift !== undefined) {
          test.info().annotations.push({ type: 'expected drift', description: formatFailures(expected) });
          expect(
            expected.length,
            `${id} at ${width}: the drift (${drift.why}) is gone; delete its entry`,
          ).toBe(drift.failures.length);
        }
      });
    }
  }
});
