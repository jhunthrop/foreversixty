// web/tests/e2e/parity-planner.spec.ts
// /planner against its mocks (design/parity.md, design/specs/2026-11-08-planner.md section 10):
// bare (signed out, Warrior), a class chosen in the bar (Hunter), arrived from the Survival guide's
// Load this build (its own levelling order), a healer's build (Holy, HPS), and Felsnap signed in at
// the level cap with two points planned past it, at 1440 and 390, each region (head, head-title,
// head-meta, head-answer, controls, notes, plate, every Section) compared for its blocks and box
// (support/parity.ts). The mocks are rendered before the tests from the data/builds files, the
// planner libs and the fixtures this spec serves (support/planner-page.ts, fixtures/planner). A
// side-by-side per state and width lands in test-results/parity/planner-<state>-<width>.png.
// Desktop project only: the 390 run is a 390 px desktop viewport, as the mock renders are.
import type { Page } from '@playwright/test';
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
import { renderPlannerMocks } from './support/parity-mocks';
import {
  addTalent,
  openPlanner,
  plannerAccount,
  PLANNER_CODES,
  type PlannerState,
} from './support/planner-page';

const felsnapKey = plannerAccount().characters[0].key;
const codeUrl = (code: string, spec?: string) =>
  `/planner?code=${encodeURIComponent(code)}${spec === undefined ? '' : `&spec=${spec}`}`;

const STATES: readonly {
  id: string;
  url: string;
  state: PlannerState;
  mock: string;
  then?: (page: Page) => Promise<void>;
}[] = [
  { id: 'planner', url: '/planner', state: 'out', mock: 'web-planner' },
  {
    id: 'planner-class',
    url: '/planner',
    state: 'out',
    mock: 'web-planner-class',
    then: async (page) => {
      const hunter = page.getByTestId('control-class-hunter').first();
      if (!(await hunter.isVisible())) await page.locator('[data-control-bar-toggle]').click();
      await hunter.click();
      await expect(page.getByTestId('page-title')).toHaveText('Hunter talents');
      await expect(page.getByTestId('band-compare-differs')).toBeVisible();
    },
  },
  {
    id: 'planner-guide',
    url: codeUrl(PLANNER_CODES.survival, 'hunter-survival'),
    state: 'out',
    mock: 'web-planner-guide',
  },
  { id: 'planner-holy', url: codeUrl(PLANNER_CODES.holy), state: 'out', mock: 'web-planner-holy' },
  {
    id: 'planner-you-cap',
    url: `/planner?character=${encodeURIComponent(felsnapKey)}`,
    state: 'felsnap',
    mock: 'web-planner-you-cap',
    then: async (page) => {
      await expect(page.getByTestId('planner-points-left')).toContainText('0');
      // Decision P4: two clicks past the cap plan ahead.
      await addTalent(page, 'Frenzy');
      await addTalent(page, 'Frenzy');
      await expect(page.getByTestId('planner-spent-line')).toContainText('2 planned ahead');
      // The live estimate sims the character's level (gap G7): wait for it to answer.
      await expect(page.getByTestId('planner-dps-note')).toContainText('±');
      await page.mouse.move(0, 0);
    },
  },
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
  'a primitive/mock-library question, the same drift parity-setup.spec.ts and parity-guides.spec.ts list';
const MOCK_PHONE_BAR =
  "below 768 the mock library's ControlBar line is a details/summary (ws_lib.control_bar); the built " +
  "ControlBar's is a button with aria-expanded, the same disclosure (every lane's bar shares it)";
const SAMPLE_SCORE =
  "the live score's figure and its error are the mock's SAMPLE (171 ± 2) and the fixture build's " +
  'test engine figure in the build; the right-aligned answer is as wide as its caption';
const EXPECTED_DRIFT: Readonly<Record<string, { failures: readonly string[]; why: string }>> = {
  ...Object.fromEntries(
    STATES.map(({ id }) => [
      `${id}-390`,
      { failures: ['head:box', 'controls:inventory'], why: `${MOCK_HEAD_INSET}; ${MOCK_PHONE_BAR}` },
    ]),
  ),
  'planner-you-cap-1440': { failures: ['head-answer:box'], why: SAMPLE_SCORE },
};

test.describe('/planner mock parity', () => {
  test.skip(({ isMobile }) => isMobile, 'the parity mocks are desktop captures');

  let mocks = '';
  // eslint-disable-next-line no-empty-pattern -- Playwright reads a hook's fixtures from its pattern
  test.beforeAll(({}, testInfo) => {
    test.skip(testInfo.project.name !== 'desktop', 'the parity mocks are desktop captures');
    mocks = renderPlannerMocks(STATES.map((state) => state.mock));
  });

  for (const width of WIDTHS) {
    for (const { id, url, state, mock: mockName, then } of STATES) {
      test(`${id} at ${width} matches ${mockName}`, async ({ page, browser }) => {
        test.slow();
        const viewport = { width, height: HEIGHT };
        await page.setViewportSize(viewport);
        const mockPage = await browser.newPage({ viewport });
        await openMock(mockPage, mocks, `${mockName}.html`);
        const mock = await captureRegions(mockPage);

        await prepareParityPage(page);
        await openPlanner(page, url, state);
        if (then !== undefined) await then(page);
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
            new Set(expected.map((f) => `${f.region}:${f.kind}`)).size,
            `${id} at ${width}: the drift (${drift.why}) is gone; delete its entry`,
          ).toBe(drift.failures.length);
        }
      });
    }
  }
});
