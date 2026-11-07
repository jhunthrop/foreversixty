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
 * a listed failure is gone so its entry is deleted. Each is a mock against the build's rule, raised
 * with the coordinator (premium round, 2026-11-08):
 *  - Devices: `Use a pairing code instead` / `Pair another computer` asks the API for a code in
 *    place, an action, so the build draws a <button> (the sim ruling: a <button> for an action);
 *    the mock draws a link.
 *  - Discord: `Link Discord` and `Link Premium role` leave for Discord, navigations, so the build
 *    keeps them links (middle-click, long-press; DiscordConnectLink); the mock draws buttons.
 *  - Discord with a channel (Premium): the shipped block keeps `Open in Discord ›`, the private
 *    logs note, Remove beside Send test message and the reconnect note, which the mock leaves out;
 *    the section is taller, and You, Delete account and the footer under it move down.
 *  - At 1440 the site header of a Premium account has no Premium link; the mock's chrome draws one.
 */
const DEVICES_ACTION = 'section-devices:inventory';
const DISCORD_LINKS = 'section-discord:inventory';
const DISCORD_SAVED = [
  'section-discord:inventory',
  'section-discord:box',
  'section-you:box',
  'section-delete-account:box',
  'site-footer:box',
];
const WHY =
  'the pairing entry is an action (a button) and the Discord links navigate (links); the mock draws them ' +
  'the other way round. Premium: the saved Discord block keeps Open in Discord, the private note, Remove ' +
  'and the reconnect note, and the header drops its Premium link';
/** Below 768 the mock puts the 18 px page inset outside the head region (.pc), PageHeader inside it;
 *  a primitive/mock-library question, the drift parity-setup and parity-connect list too. */
const MOCK_HEAD_INSET = 'head:box';
const drift = (failures: readonly string[], width: number) => ({
  failures: width === 390 ? [MOCK_HEAD_INSET, ...failures] : failures,
  why: width === 390 ? `${WHY}; and the mock's 18 px head inset below 768` : WHY,
});
const EXPECTED_DRIFT: Readonly<Record<string, { failures: readonly string[]; why: string }>> =
  Object.fromEntries(
    WIDTHS.flatMap((width) => [
      [`account-${width}`, drift([DEVICES_ACTION, DISCORD_LINKS], width)],
      [`account-new-${width}`, drift([DEVICES_ACTION, DISCORD_LINKS], width)],
      [`account-delete-${width}`, drift([DEVICES_ACTION, DISCORD_LINKS], width)],
      [
        `account-premium-${width}`,
        drift(
          [...(width === 1440 ? ['site-header:inventory'] : []), DEVICES_ACTION, ...DISCORD_SAVED],
          width,
        ),
      ],
    ]),
  );

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
