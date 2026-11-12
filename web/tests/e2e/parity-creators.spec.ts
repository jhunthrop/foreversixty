// web/tests/e2e/parity-creators.spec.ts
// The creator program's pages against their mocks (design/parity.md; ws_creators.py, drawn from lib/creators/copy.ts and
// model.ts over fixtures/creators by support/parity-mocks.ts renderCreatorsMocks): the public /creators page signed out at
// 1440 and 390 (the parity twin of the directions mock b4-creators), then at 1440 the owner (Thandrel,
// fixtures/home/me-60.json, role admin) on /admin/creators and its statements, and the same account as a creator on
// /creators/stats and /creators/apply. Every region (head, the lead, each section) is compared for its blocks and box
// (support/parity.ts); a side-by-side per page lands in test-results/parity/creators-*.png. Desktop project only: the
// 390 run is a 390 px desktop viewport, as the mock renders are. Also: Apply at the head and the end goes to /creators/apply.
import { test, expect } from './support/test';
import {
  REPO_ROOT,
  captureOutlined,
  captureRegions,
  compareRegions,
  formatFailures,
  openMock,
  prepareParityPage,
  settleParityPage,
  writeSideBySide,
} from './support/parity';
import { renderCreatorsMocks } from './support/parity-mocks';
import { openPagesState } from './support/pages-page';
import { readFileSync } from 'node:fs';
import path from 'node:path';
import type { Page } from '@playwright/test';
import { CREATORS_APPLY_PATH } from '../../src/lib/creators/copy';

const WIDTHS = [1440, 390] as const;
const HEIGHT = 1000;
const MOCK_HEAD_INSET =
  'below 768 the mock puts the 18 px page inset outside the head region (.pc), PageHeader inside it; ' +
  'the same mock-library drift parity-pages.spec.ts lists';
const EXPECTED_DRIFT: Readonly<Record<string, { failures: readonly string[]; why: string }>> = {
  'creators-390': { failures: ['head:box'], why: MOCK_HEAD_INSET },
};

test.describe('/creators mock parity', () => {
  test.skip(({ isMobile }) => isMobile, 'the parity mocks are desktop captures');

  let mocks = '';
  // eslint-disable-next-line no-empty-pattern -- Playwright reads a hook's fixtures from its pattern
  test.beforeAll(({}, testInfo) => {
    test.skip(testInfo.project.name !== 'desktop', 'the parity mocks are desktop captures');
    mocks = renderCreatorsMocks();
  });

  for (const width of WIDTHS) {
    test(`creators at ${width} matches web-creators`, async ({ page, browser }) => {
      test.slow();
      const viewport = { width, height: HEIGHT };
      await page.setViewportSize(viewport);
      const mockPage = await browser.newPage({ viewport });
      await openMock(mockPage, mocks, 'web-creators.html');
      const mock = await captureRegions(mockPage);

      await prepareParityPage(page);
      await openPagesState(page, '/creators', 'out');
      await settleParityPage(page);
      const built = await captureRegions(page);

      await writeSideBySide(
        browser,
        `creators-${width}`,
        await captureOutlined(mockPage, mock),
        await captureOutlined(page, built),
      );
      await mockPage.close();

      const drift = EXPECTED_DRIFT[`creators-${width}`];
      const failures = compareRegions(mock, built);
      const expected = failures.filter((f) => drift?.failures.includes(`${f.region}:${f.kind}`));
      const unexpected = failures.filter((f) => !expected.includes(f));
      expect(
        unexpected,
        `creators at ${width} differs from web-creators:\n${formatFailures(unexpected)}`,
      ).toEqual([]);
      if (drift !== undefined) {
        test.info().annotations.push({ type: 'expected drift', description: formatFailures(expected) });
        expect(
          expected.length,
          `creators at ${width}: the drift (${drift.why}) is gone; delete its entry`,
        ).toBeGreaterThan(0);
      }
    });
  }
});

test('Apply at the head and at the end goes to the application form', async ({ page }) => {
  await openPagesState(page, '/creators', 'out');
  await expect(page.getByTestId('creators-apply-top')).toHaveAttribute('href', CREATORS_APPLY_PATH);
  await expect(page.getByTestId('creators-apply-end')).toHaveAttribute('href', CREATORS_APPLY_PATH);
  await expect(page.getByTestId('creators-figures-what-you-get')).toBeVisible();
});

const VIEWPORT = { width: 1440, height: HEIGHT };
const FIX = path.join(REPO_ROOT, 'web/tests/e2e/fixtures/creators');
const read = (name: string): unknown => JSON.parse(readFileSync(path.join(FIX, name), 'utf8'));

const envelope = (data: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify({
    ok: status < 400,
    data: status < 400 ? data : null,
    error: status < 400 ? null : { code: 'x', message: 'x' },
    request_id: 'p',
  }),
});

const owner = JSON.parse(
  readFileSync(path.join(REPO_ROOT, 'web/tests/e2e/fixtures/home/me-60.json'), 'utf8'),
) as {
  user: Record<string, unknown>;
};

async function signedInAsOwner(page: Page, extra: Record<string, unknown> = {}): Promise<void> {
  await page.context().addCookies([{ name: 'fs_csrf', value: 'tok', domain: 'localhost', path: '/' }]);
  await page.route('**/v1/me', (route) =>
    route.fulfill(envelope({ ...owner, user: { ...owner.user, role: 'admin' }, ...extra })),
  );
}

interface State {
  mock: string;
  url: string;
  ready: string;
  setup: (page: Page) => Promise<void>;
}

const STATES: State[] = [
  {
    mock: 'web-admin-creators',
    url: '/admin/creators',
    ready: 'creator-row-grimbrew',
    setup: async (page) => {
      await signedInAsOwner(page);
      await page.route('**/v1/admin/creators', (route) =>
        route.fulfill(
          envelope({ creators: read('creators.json'), stripe_enabled: true, payouts_enabled: true }),
        ),
      );
      await page.route('**/v1/admin/creators/requests', (route) =>
        route.fulfill(envelope({ requests: read('requests.json') })),
      );
    },
  },
  {
    mock: 'web-admin-creators-payouts',
    url: '/admin/creators/payouts',
    ready: 'payout-row-1',
    setup: async (page) => {
      await signedInAsOwner(page);
      await page.route('**/v1/admin/creators/statements**', (route) =>
        route.fulfill(envelope(read('statements.json'))),
      );
    },
  },
  {
    mock: 'web-creators-stats',
    url: '/creators/stats',
    ready: 'stats-figures',
    setup: async (page) => {
      await signedInAsOwner(page, { creator: true });
      await page.route('**/v1/creators/me', (route) => route.fulfill(envelope(read('me.json'))));
    },
  },
  {
    mock: 'web-creators-apply',
    url: '/creators/apply',
    ready: 'apply-form',
    setup: async (page) => {
      await signedInAsOwner(page);
      await page.route('**/v1/creators/application', (route) =>
        route.fulfill(envelope({ status: null, created_at: null, creator: false })),
      );
    },
  },
];

test.describe('the program pages mock parity', () => {
  test.skip(({ isMobile }) => isMobile, 'the parity mocks are desktop captures');

  let mocks = '';
  // eslint-disable-next-line no-empty-pattern -- Playwright reads a hook's fixtures from its pattern
  test.beforeAll(({}, testInfo) => {
    test.skip(testInfo.project.name !== 'desktop', 'the parity mocks are desktop captures');
    mocks = renderCreatorsMocks();
  });

  for (const state of STATES) {
    test(`${state.url} at 1440 matches ${state.mock}`, async ({ page, browser }) => {
      test.slow();
      await page.setViewportSize(VIEWPORT);
      const mockPage = await browser.newPage({ viewport: VIEWPORT });
      await openMock(mockPage, mocks, `${state.mock}.html`);
      const mock = await captureRegions(mockPage);

      await prepareParityPage(page);
      await state.setup(page);
      await page.goto(state.url);
      await expect(page.getByTestId(state.ready)).toBeVisible();
      await settleParityPage(page);
      const built = await captureRegions(page);

      await writeSideBySide(
        browser,
        `creators-${state.mock}-1440`,
        await captureOutlined(mockPage, mock),
        await captureOutlined(page, built),
      );
      await mockPage.close();
      const failures = compareRegions(mock, built);
      expect(failures, `${state.url} differs from ${state.mock}:\n${formatFailures(failures)}`).toEqual([]);
    });
  }
});
