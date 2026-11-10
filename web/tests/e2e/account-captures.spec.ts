// web/tests/e2e/account-captures.spec.ts
// Not a test: the account lane's review captures (redesign plan, Wave 2, P11 and P13: "capture at 390,
// 1024, 1440 and 2000 beside the mock"). Skipped unless ACCOUNT_CAPTURE=1:
//
//   ACCOUNT_CAPTURE=1 npx playwright test tests/e2e/account-captures.spec.ts --project=desktop --workers=1
//
// Each view opens the built page as its parity spec does (the parity fixtures and clocks) and its batch 4
// directions mock (design/mocks/directions/html), full page at every width, into
// design/reviews/2026-11-09-f-account/{build,mock}-<view>-<width>.png.
import { mkdirSync } from 'node:fs';
import path from 'node:path';
import type { Page } from '@playwright/test';
import { test } from './support/test';
import { openAccountState, renderAccountMocks, type AccountPageState } from './support/account-page';
import { openConnectState, type ConnectPageState } from './support/connect-page';
import { REPO_ROOT, prepareParityPage, settleParityPage } from './support/parity';
import { openSetupState, type SetupPageState } from './support/setup-page';

const OUT = path.join(REPO_ROOT, 'design/reviews/2026-11-09-f-account');
const MOCKS = path.join(REPO_ROOT, 'design/mocks/directions/html');
const WIDTHS = [390, 1024, 1440, 2000] as const;

const signedOut = async (page: Page, pathname: string): Promise<void> => {
  await page.route('**/v1/me', (route) => route.fulfill({ status: 401, body: '' }));
  await page.goto(pathname);
  await page.waitForLoadState('networkidle');
};
const account = (state: AccountPageState) => (page: Page) => openAccountState(page, state);
const setup = (state: SetupPageState) => (page: Page) => openSetupState(page, state);
const connect = (state: ConnectPageState) => (page: Page) => openConnectState(page, state);

/** Each view: how to open it and its directions mock (none: the build alone). */
const VIEWS: ReadonlyArray<{ view: string; open: (page: Page) => Promise<void>; mock?: string }> = [
  { view: 'account', open: account('account-premium'), mock: 'b4-account' },
  { view: 'account-free', open: account('account') },
  { view: 'account-new', open: account('account-new') },
  { view: 'account-out', open: (page) => signedOut(page, '/account') },
  { view: 'login', open: (page) => signedOut(page, '/login'), mock: 'b4-login' },
  { view: 'setup', open: setup('out'), mock: 'b4-setup' },
  { view: 'setup-done', open: setup('done') },
  { view: 'connect', open: connect('pending'), mock: 'b4-connect' },
  { view: 'connect-out', open: connect('signed-out') },
  { view: 'connect-success', open: connect('success') },
  {
    view: 'app-auth',
    open: (page) => signedOut(page, '/app/auth?code=kq47tm&state=s1'),
    mock: 'b4-app-auth',
  },
  { view: 'app', open: (page) => signedOut(page, '/app') },
];

const only = process.env.ACCOUNT_CAPTURE_VIEWS?.split(',');

/** A full-page capture never scrolls, so lazy captures below the fold would shoot empty: load them all first. */
async function loadLazyImages(page: Page): Promise<void> {
  await page.evaluate(async () => {
    const images = [...document.querySelectorAll('img')];
    for (const image of images) image.loading = 'eager';
    await Promise.all(
      images.map((image) =>
        image.complete
          ? null
          : new Promise((resolve) => {
              image.addEventListener('load', resolve, { once: true });
              image.addEventListener('error', resolve, { once: true });
            }),
      ),
    );
    // Loaded is not painted: decode each before the shot.
    await Promise.all(images.map((image) => image.decode().catch(() => null)));
  });
}

test.describe('Account lane review captures', () => {
  test.skip(() => process.env.ACCOUNT_CAPTURE !== '1', 'a capture tool, run on demand');
  test.skip(({ isMobile }) => isMobile, 'the widths are set per capture');

  // eslint-disable-next-line no-empty-pattern -- Playwright reads a hook's fixtures from its pattern
  test.beforeAll(({}) => {
    // The account states serve the bodies the facts file holds (support/account-page.ts).
    renderAccountMocks(['web-account']);
  });

  for (const { view, open, mock } of VIEWS.filter((v) => only === undefined || only.includes(v.view))) {
    test(view, async ({ page, browser }) => {
      test.slow();
      mkdirSync(OUT, { recursive: true });
      for (const width of WIDTHS) {
        await page.setViewportSize({ width, height: 1000 });
        await prepareParityPage(page);
        await open(page);
        await settleParityPage(page);
        await loadLazyImages(page);
        await page.screenshot({ path: path.join(OUT, `build-${view}-${width}.png`), fullPage: true });
        await page.unrouteAll({ behavior: 'ignoreErrors' });

        if (mock === undefined) continue;
        const mockPage = await browser.newPage({ viewport: { width, height: 1000 } });
        await mockPage.goto(`file://${path.join(MOCKS, `${mock}.html`)}`);
        await mockPage.waitForLoadState('networkidle');
        await mockPage.screenshot({ path: path.join(OUT, `mock-${view}-${width}.png`), fullPage: true });
        await mockPage.close();
      }
    });
  }
});
