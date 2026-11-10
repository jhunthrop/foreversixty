// web/tests/e2e/parity-account.spec.ts
// /account against its mocks (design/parity.md, design/specs/2026-11-08-premium-account.md section
// 14): a free account with one character, an account with none, and Thandrel on Premium with a
// guild, a device and Discord linked, at 1440 and 390, each region (head, head-title, head-meta,
// head-answer with a character, the Characters plate and every Section, Delete account last in the
// rail) compared for its blocks and
// box (support/parity.ts). The mocks are rendered before the tests from the copy modules and the
// fixtures this spec serves (support/account-page.ts), at the mocks' pinned clock. A side-by-side
// per state and width lands in test-results/parity/account-<state>-<width>.png. `account-delete`
// is the free account with the Delete account dialog open over it (region `dialog`). Desktop project
// only: the 390 run is a 390 px desktop viewport, as the mock renders are.
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
import type { Page } from '@playwright/test';
import {
  openAccountDeleteDialog,
  openAccountState,
  renderAccountMocks,
  type AccountPageState,
} from './support/account-page';

const STATES: readonly {
  id: string;
  state: AccountPageState;
  mock: string;
  /** What the viewer does on the page before it is measured (opens a dialog). */
  act?: (page: Page) => Promise<void>;
}[] = [
  { id: 'account', state: 'account', mock: 'web-account' },
  { id: 'account-new', state: 'account-new', mock: 'web-account-new' },
  { id: 'account-premium', state: 'account-premium', mock: 'web-account-premium' },
  { id: 'account-delete', state: 'account', mock: 'web-account-delete', act: openAccountDeleteDialog },
];
const WIDTHS = [1440, 390] as const;
const HEIGHT = 1000;

/**
 * Known differences (design/parity.md EXPECTED_DRIFT), keyed `<id>-<width>`: each lists the
 * `region:kind` failures expected; every other region is still compared, and the test fails once
 * a listed failure is gone so its entry is deleted. Below 768 the mock puts the 18 px page inset
 * outside the head region (.pc), PageHeader inside it; a primitive/mock-library question, the drift
 * parity-setup and parity-connect list too.
 */
const MOCK_HEAD_INSET = { failures: ['head:box'], why: "the mock's 18 px head inset below 768" };
const PREMIUM_ROWS = (regions: readonly string[]): string[] => regions.map((region) => `${region}:box`);
const EXPECTED_DRIFT: Readonly<Record<string, { failures: readonly string[]; why: string }>> = {
  ...Object.fromEntries(STATES.map(({ id }) => [`${id}-390`, MOCK_HEAD_INSET])),
  // A Premium account's site header has no Premium link; the mock's chrome draws one. Its saved Discord block and
  // the guild block above it are the shared settings rows (ui/Row, ui/Switch), each a few px taller than ws_prims
  // draws them; since the mock's Tag took the built 22 px (2026-11-10) only the panels further down move past the
  // 8 px tolerance.
  'account-premium-1440': {
    failures: ['site-header:inventory', ...PREMIUM_ROWS(['site-footer'])],
    why: 'the Premium header drops its Premium link; the settings rows run taller than the mock draws them',
  },
  'account-premium-390': {
    failures: [
      'head:box',
      ...PREMIUM_ROWS([
        'section-you',
        'plate#5',
        'section-plan',
        'plate#6',
        'section-delete-account',
        'plate#7',
        'site-footer',
      ]),
    ],
    why: "the mock's 18 px head inset below 768; the settings rows run taller than the mock draws them",
  },
};

test.describe('/account mock parity', () => {
  test.skip(({ isMobile }) => isMobile, 'the parity mocks are desktop captures');

  let mocks = '';
  // eslint-disable-next-line no-empty-pattern -- Playwright reads a hook's fixtures from its pattern
  test.beforeAll(({}, testInfo) => {
    test.skip(testInfo.project.name !== 'desktop', 'the parity mocks are desktop captures');
    mocks = renderAccountMocks(STATES.map((state) => state.mock));
  });

  for (const width of WIDTHS) {
    for (const { id, state, mock: mockName, act } of STATES) {
      test(`${id} at ${width} matches ${mockName}`, async ({ page, browser }) => {
        test.slow();
        const viewport = { width, height: HEIGHT };
        await page.setViewportSize(viewport);
        const mockPage = await browser.newPage({ viewport });
        await openMock(mockPage, mocks, `${mockName}.html`);
        const mock = await captureRegions(mockPage);

        await prepareParityPage(page);
        await openAccountState(page, state);
        await settleParityPage(page);
        await act?.(page);
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
          // An inventory drift can fail on several role chains: each listed key must still fail.
          const still = new Set(expected.map((f) => `${f.region}:${f.kind}`));
          expect(
            drift.failures.filter((key) => !still.has(key)),
            `${id} at ${width}: part of the drift (${drift.why}) is gone; delete it`,
          ).toEqual([]);
        }
      });
    }
  }
});
