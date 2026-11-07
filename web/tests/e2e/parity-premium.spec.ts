// web/tests/e2e/parity-premium.spec.ts
// /premium and /premium/checkout against their mocks (design/parity.md, design/specs/2026-11-08-premium-account.md
// section 14): /premium signed out, signed in on a free account and holding a guild seat, and the
// checkout back from Stripe with Premium on, at 1440 and 390; every region compared for its blocks
// and box (support/parity.ts). The mocks are rendered before the tests from the copy modules and
// the facts reader's own accounts (support/premium-page.ts), at the mocks' pinned clock. A
// side-by-side per state and width lands in test-results/parity/premium-<state>-<width>.png.
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
import { openPremiumState, renderPremiumMocks, type PremiumPageState } from './support/premium-page';

const STATES: readonly { id: string; state: PremiumPageState; mock: string }[] = [
  { id: 'premium', state: 'out', mock: 'web-premium' },
  { id: 'premium-free', state: 'free', mock: 'web-premium-free' },
  { id: 'premium-seat', state: 'seat', mock: 'web-premium-seat' },
  { id: 'checkout-confirmed', state: 'checkout-confirmed', mock: 'web-checkout-confirmed' },
  { id: 'checkout-owned', state: 'checkout-owned', mock: 'web-checkout-owned' },
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
  'the same mock-library drift parity-setup.spec.ts lists (2026-11-07)';
const PREMIUM_DOOR =
  'the header hides its Premium door for an account that has Premium (Header.astro, data-premium-nav); ' +
  'the mock chrome always draws it';
const MOCK_PHONE_BAR =
  "below 768 the mock library's ControlBar line is a details/summary (ws_lib.control_bar); the built " +
  "ControlBar's is a button with aria-expanded, the same disclosure (every lane's bar shares it)";
const PAID = new Set(['premium-seat', 'checkout-confirmed', 'checkout-owned']);
const HAS_BAR = (id: string) => id.startsWith('premium');
const SECTION_META =
  "the built Section's head with a meta line (Questions) is 12 px taller than the mock library's sec-h, " +
  'which moves the footer by the same 12 px';
const PHONE_ROWS =
  "at 354 px the built open Rows of the rules wrap a line fewer than the mock's, and the rail's coach " +
  'sample a line more; every block and its order match, only the heights below the Plate move';
const drift = (id: string, width: number): { failures: string[]; why: string[] } => {
  const failures: string[] = [];
  const why: string[] = [];
  const add = (reason: string, ...items: string[]) => {
    failures.push(...items);
    why.push(reason);
  };
  if (width === 390) add(MOCK_HEAD_INSET, 'head:box');
  if (width === 390 && HAS_BAR(id)) {
    add(MOCK_PHONE_BAR, 'controls:inventory');
    add(
      PHONE_ROWS,
      'section-how-the-guild-plans-work:box',
      'section-questions:box',
      'section-what-the-coach-says-after-a-raid:box',
      'site-footer:box',
    );
    if (id !== 'premium') failures.push('plate:box');
  }
  if (width === 1440 && HAS_BAR(id)) add(SECTION_META, 'section-questions:box', 'site-footer:box');
  if (width === 1440 && PAID.has(id)) add(PREMIUM_DOOR, 'site-header:inventory');
  return { failures, why };
};
const EXPECTED_DRIFT: Readonly<Record<string, { failures: readonly string[]; why: string }>> =
  Object.fromEntries(
    STATES.flatMap(({ id }) =>
      WIDTHS.map((width) => {
        const { failures, why } = drift(id, width);
        return [`${id}-${width}`, { failures, why: why.join('; ') }] as const;
      }).filter(([, d]) => d.failures.length > 0),
    ),
  );

test.describe('/premium and checkout mock parity', () => {
  test.skip(({ isMobile }) => isMobile, 'the parity mocks are desktop captures');

  let mocks = '';
  // eslint-disable-next-line no-empty-pattern -- Playwright reads a hook's fixtures from its pattern
  test.beforeAll(({}, testInfo) => {
    test.skip(testInfo.project.name !== 'desktop', 'the parity mocks are desktop captures');
    mocks = renderPremiumMocks(STATES.map((state) => state.mock));
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
        await openPremiumState(page, state, mocks);
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
          // A box drift below the Plate depends on line wraps that move by a few pixels between
          // machines, so an entry may hold in part; once none of it holds, it goes.
          expect(
            expected.length,
            `${id} at ${width}: the drift (${drift.why}) is gone; delete its entry`,
          ).toBeGreaterThan(0);
        }
      });
    }
  }
});
