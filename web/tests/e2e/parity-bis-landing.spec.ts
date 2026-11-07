// web/tests/e2e/parity-bis-landing.spec.ts
// /bis against its mocks (design/parity.md, design/specs/2026-11-07-bis-landing.md section 6):
// signed out at every level range, Felsnap at 24 (the owner's report: 2 of 14 slots at 20 to 29),
// 34 and with no gear, Thandrel at 60, no character, a finished server sim, loading and error,
// each region (head, head-title, head-meta, head-answer when signed in, controls, summary, plate)
// compared for its blocks and box (support/parity.ts). The mocks are rendered before the tests
// from the landing's own models, this checkout's BiS files and the fixtures this spec serves
// (support/bis-landing-page.ts). A side-by-side per state and width lands in
// test-results/parity/bis-landing-<state>-<width>.png. Desktop project only.
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
import { openBisLanding, type BisLandingState } from './support/bis-landing-page';
import { renderBisLandingMocks } from './support/parity-mocks';

interface ParityState {
  id: string;
  path: string;
  state: BisLandingState;
  mock: string;
  widths: readonly number[];
}

const BOTH = [1440, 390] as const;
const DESKTOP = [1440] as const;
const STATES: readonly ParityState[] = [
  { id: 'out', path: '/bis', state: 'out', mock: 'web-bis-landing', widths: BOTH },
  ...[30, 40, 50, 60].map((band): ParityState => ({
    id: `out-${band}`,
    path: `/bis?band=${band}`,
    state: 'out',
    mock: `web-bis-landing-out-${band}`,
    widths: DESKTOP,
  })),
  { id: '24', path: '/bis', state: '24', mock: 'web-bis-landing-24', widths: BOTH },
  { id: '34', path: '/bis', state: '34', mock: 'web-bis-landing-34', widths: DESKTOP },
  { id: '60', path: '/bis', state: '60', mock: 'web-bis-landing-60', widths: DESKTOP },
  { id: 'new', path: '/bis', state: 'new', mock: 'web-bis-landing-new', widths: DESKTOP },
  { id: 'nogear', path: '/bis', state: 'nogear', mock: 'web-bis-landing-nogear', widths: DESKTOP },
  { id: 'sim', path: '/bis', state: 'sim', mock: 'web-bis-landing-sim', widths: DESKTOP },
  { id: 'loading', path: '/bis', state: 'loading', mock: 'web-bis-landing-loading', widths: DESKTOP },
  { id: 'error', path: '/bis', state: 'error', mock: 'web-bis-landing-error', widths: DESKTOP },
];
const HEIGHT = 1000;

/**
 * Known differences (design/parity.md EXPECTED_DRIFT), keyed `<id>-<width>`: each lists the
 * `region:kind` failures expected; every other region is still compared, and the test fails once
 * a listed failure is gone so its entry is deleted.
 */
const MOCK_HEAD_INSET =
  'below 768 the mock puts the 18 px page inset outside the head region (.pc), PageHeader inside it; ' +
  'a primitive/mock-library question, the same drift parity-guides.spec.ts lists (2026-11-08)';
const MOCK_PHONE_BAR =
  "below 768 the mock library's ControlBar line is a details/summary (ws_lib.control_bar); the built " +
  "ControlBar's is a button with aria-expanded, the same disclosure (every lane's bar shares it)";
const EXPECTED_DRIFT: Readonly<Record<string, { failures: readonly string[]; why: string }>> =
  Object.fromEntries(
    STATES.filter(({ widths }) => widths.includes(390)).map(({ id }) => [
      `${id}-390`,
      { failures: ['head:box', 'controls:inventory'], why: `${MOCK_HEAD_INSET}; ${MOCK_PHONE_BAR}` },
    ]),
  );

test.describe('/bis landing mock parity', () => {
  test.skip(({ isMobile }) => isMobile, 'the parity mocks are desktop captures');

  let mocks = '';
  // eslint-disable-next-line no-empty-pattern -- Playwright reads a hook's fixtures from its pattern
  test.beforeAll(({}, testInfo) => {
    test.skip(testInfo.project.name !== 'desktop', 'the parity mocks are desktop captures');
    mocks = renderBisLandingMocks(STATES.map((state) => state.mock));
  });

  for (const { id, path, state, mock: mockName, widths } of STATES) {
    for (const width of widths) {
      test(`${id} at ${width} matches ${mockName}`, async ({ page, browser }) => {
        test.slow();
        const viewport = { width, height: HEIGHT };
        await page.setViewportSize(viewport);
        const mockPage = await browser.newPage({ viewport });
        await openMock(mockPage, mocks, `${mockName}.html`);
        const mock = await captureRegions(mockPage);

        await prepareParityPage(page);
        await openBisLanding(page, path, state);
        await settleParityPage(page);
        const built = await captureRegions(page);

        await writeSideBySide(
          browser,
          `bis-landing-${id}-${width}`,
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
            new Set(expected.map((f) => `${f.region}:${f.kind}`)).size,
            `${id} at ${width}: the drift (${drift.why}) is gone; delete its entry`,
          ).toBe(drift.failures.length);
        }
      });
    }
  }
});
