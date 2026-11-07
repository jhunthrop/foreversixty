// web/tests/e2e/parity-pages.spec.ts
// The changelog, a legal page, the 404 and sign in against their mocks (design/parity.md,
// design/specs/2026-11-08-pages.md section 10): the changelog (All and Game), privacy, the 404
// signed out and as Felsnap, and /login signed out, at 1440 and 390, each region (head, head-title, head-meta,
// head-answer signed in, controls, plate, every section and month) compared for its blocks and
// box (support/parity.ts). The mocks are rendered before the tests from the content collections,
// the copy modules, the changelog model and the fixture this spec serves (support/pages-page.ts).
// A side-by-side per state and width lands in test-results/parity/pages-<state>-<width>.png.
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
import { openPagesState, type PagesState } from './support/pages-page';
import { renderPagesMocks } from './support/parity-mocks';

const STATES: readonly { id: string; path: string; state: PagesState; mock: string }[] = [
  { id: 'changelog', path: '/changelog', state: 'out', mock: 'web-changelog' },
  { id: 'changelog-game', path: '/changelog?show=game', state: 'out', mock: 'web-changelog-game' },
  { id: 'legal-privacy', path: '/privacy', state: 'out', mock: 'web-legal-privacy' },
  { id: '404', path: '/this-page-does-not-exist', state: 'out', mock: 'web-404' },
  { id: '404-in', path: '/this-page-does-not-exist', state: 'in', mock: 'web-404-in' },
  { id: 'login', path: '/login', state: 'out', mock: 'web-login' },
];
const WIDTHS = [1440, 390] as const;
const HEIGHT = 1000;

/**
 * Known differences (design/parity.md EXPECTED_DRIFT), keyed `<id>-<width>`: each lists the
 * `region:kind` failures expected; every other region is still compared, and the test fails once
 * a listed failure is gone so its entry is deleted.
 */
const MOCK_HEAD_INSET =
  'below 768 the mock puts the 18 px page inset outside the head region (.pc), PageHeader inside it; ' +
  'a primitive/mock-library question, the same drift parity-setup.spec.ts lists (2026-11-08)';
const MOCK_PHONE_BAR =
  "below 768 the mock library's ControlBar line is a details/summary (ws_lib.control_bar); the built " +
  "ControlBar's is a button with aria-expanded, the same disclosure (every lane's bar shares it)";
const HAS_BAR = new Set(['changelog', 'changelog-game']);
const EXPECTED_DRIFT: Readonly<Record<string, { failures: readonly string[]; why: string }>> =
  Object.fromEntries(
    STATES.map(({ id }) => [
      `${id}-390`,
      HAS_BAR.has(id)
        ? { failures: ['head:box', 'controls:inventory'], why: `${MOCK_HEAD_INSET}; ${MOCK_PHONE_BAR}` }
        : { failures: ['head:box'], why: MOCK_HEAD_INSET },
    ]),
  );

test.describe('changelog, legal and 404 mock parity', () => {
  test.skip(({ isMobile }) => isMobile, 'the parity mocks are desktop captures');

  let mocks = '';
  // eslint-disable-next-line no-empty-pattern -- Playwright reads a hook's fixtures from its pattern
  test.beforeAll(({}, testInfo) => {
    test.skip(testInfo.project.name !== 'desktop', 'the parity mocks are desktop captures');
    mocks = renderPagesMocks(STATES.map((state) => state.mock));
  });

  for (const width of WIDTHS) {
    for (const { id, path, state, mock: mockName } of STATES) {
      test(`${id} at ${width} matches ${mockName}`, async ({ page, browser }) => {
        test.slow();
        const viewport = { width, height: HEIGHT };
        await page.setViewportSize(viewport);
        const mockPage = await browser.newPage({ viewport });
        await openMock(mockPage, mocks, `${mockName}.html`);
        const mock = await captureRegions(mockPage);

        await prepareParityPage(page);
        await openPagesState(page, path, state);
        await settleParityPage(page);
        const built = await captureRegions(page);

        await writeSideBySide(
          browser,
          `pages-${id}-${width}`,
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
