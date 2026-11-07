// web/tests/e2e/parity-guides.spec.ts
// The guide pages against their mocks (design/parity.md, design/specs/2026-11-08-guides.md section
// 10): the index signed out and as Felsnap, the hunter landing, Survival signed out and as a level
// 34 Survival hunter, Fire and Holy, at 1440 and 390, each region (head, head-title, head-meta,
// head-answer when signed in, controls, plate, every Section, footer-note) compared for its blocks
// and box (support/parity.ts). The mocks are rendered before the tests from the guide content, the
// guide libs and the fixtures this spec serves (support/guides-page.ts, fixtures/guides). A
// side-by-side per state and width lands in test-results/parity/guides-<state>-<width>.png.
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
import { openGuidesState, type GuidesState } from './support/guides-page';
import { renderGuidesMocks } from './support/parity-mocks';

const STATES: readonly { id: string; path: string; state: GuidesState; mock: string }[] = [
  { id: 'guides', path: '/guides', state: 'out', mock: 'web-guides' },
  { id: 'guides-you', path: '/guides', state: 'hunter34', mock: 'web-guides-you' },
  { id: 'guides-hunter', path: '/guides/hunter', state: 'out', mock: 'web-guides-hunter' },
  {
    id: 'guides-hunter-survival',
    path: '/guides/hunter/survival',
    state: 'out',
    mock: 'web-guides-hunter-survival',
  },
  {
    id: 'guides-hunter-survival-you',
    path: '/guides/hunter/survival',
    state: 'survival34',
    mock: 'web-guides-hunter-survival-you',
  },
  { id: 'guides-mage-fire', path: '/guides/mage/fire', state: 'out', mock: 'web-guides-mage-fire' },
  { id: 'guides-priest-holy', path: '/guides/priest/holy', state: 'out', mock: 'web-guides-priest-holy' },
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
/** The spec guides carry the ControlBar; /guides and the class landing have none. */
const HAS_BAR = new Set([
  'guides-hunter-survival',
  'guides-hunter-survival-you',
  'guides-mage-fire',
  'guides-priest-holy',
]);
const EXPECTED_DRIFT: Readonly<Record<string, { failures: readonly string[]; why: string }>> =
  Object.fromEntries(
    STATES.map(({ id }) => [
      `${id}-390`,
      HAS_BAR.has(id)
        ? {
            failures: ['head:box', 'controls:inventory'],
            why: `${MOCK_HEAD_INSET}; ${MOCK_PHONE_BAR}`,
          }
        : { failures: ['head:box'], why: MOCK_HEAD_INSET },
    ]),
  );

test.describe('/guides mock parity', () => {
  test.skip(({ isMobile }) => isMobile, 'the parity mocks are desktop captures');

  let mocks = '';
  // eslint-disable-next-line no-empty-pattern -- Playwright reads a hook's fixtures from its pattern
  test.beforeAll(({}, testInfo) => {
    test.skip(testInfo.project.name !== 'desktop', 'the parity mocks are desktop captures');
    mocks = renderGuidesMocks(STATES.map((state) => state.mock));
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
        await openGuidesState(page, path, state);
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
          // One entry per region:kind, however many inventory keys it spans.
          expect(
            new Set(expected.map((f) => `${f.region}:${f.kind}`)).size,
            `${id} at ${width}: the drift (${drift.why}) is gone; delete its entry`,
          ).toBe(drift.failures.length);
        }
      });
    }
  }
});
