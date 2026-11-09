// web/tests/e2e/parity-creators.spec.ts
// The creator program's pages against their mocks (design/parity.md, design/specs/2026-11-08-creator-program.md):
// the owner (Thandrel, fixtures/home/me-60.json, role admin) on /admin/creators and its statements, the
// same account as a creator on /creators/stats and /creators/apply, and /creators signed out, at 1440.
// The mocks are rendered before the tests from the copy and model modules over fixtures/creators
// (design/mocks/creators_facts.mjs, then ws_creators.py). Side-by-sides land in
// test-results/parity/creators-*.png. Desktop project only.
import { execFileSync } from 'node:child_process';
import { mkdirSync, readFileSync, rmSync } from 'node:fs';
import path from 'node:path';
import type { Page } from '@playwright/test';
import { test, expect } from './support/test';
import {
  PARITY_OUT,
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

const VIEWPORT = { width: 1440, height: 1000 };
const FIX = path.join(REPO_ROOT, 'web/tests/e2e/fixtures/creators');
const read = (name: string): unknown => JSON.parse(readFileSync(path.join(FIX, name), 'utf8'));

function renderMocks(): string {
  const dir = path.join(PARITY_OUT, 'mocks', 'creators');
  rmSync(dir, { recursive: true, force: true });
  mkdirSync(dir, { recursive: true });
  const facts = path.join(dir, 'creators-facts.json');
  const generators = path.join(REPO_ROOT, 'design/mocks');
  const env = { ...process.env, FS_CREATORS_FACTS: facts, FS_MOCK_OUT: dir };
  execFileSync('node', [path.join(generators, 'creators_facts.mjs'), facts], { cwd: REPO_ROOT, env });
  execFileSync(process.env.PYTHON ?? 'python3', [path.join(generators, 'ws_creators.py')], { cwd: REPO_ROOT, env });
  return dir;
}

const envelope = (data: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify({ ok: status < 400, data: status < 400 ? data : null, error: status < 400 ? null : { code: 'x', message: 'x' }, request_id: 'p' }),
});

const owner = JSON.parse(readFileSync(path.join(REPO_ROOT, 'web/tests/e2e/fixtures/home/me-60.json'), 'utf8')) as {
  user: Record<string, unknown>;
};

async function signedInAsOwner(page: Page, extra: Record<string, unknown> = {}): Promise<void> {
  await page.context().addCookies([{ name: 'fs_csrf', value: 'tok', domain: 'localhost', path: '/' }]);
  await page.route('**/v1/me', (route) => route.fulfill(envelope({ ...owner, user: { ...owner.user, role: 'admin' }, ...extra })));
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
        route.fulfill(envelope({ creators: read('creators.json'), stripe_enabled: true, payouts_enabled: true })),
      );
      await page.route('**/v1/admin/creators/requests', (route) => route.fulfill(envelope({ requests: read('requests.json') })));
    },
  },
  {
    mock: 'web-admin-creators-payouts',
    url: '/admin/creators/payouts',
    ready: 'payout-row-1',
    setup: async (page) => {
      await signedInAsOwner(page);
      await page.route('**/v1/admin/creators/statements**', (route) => route.fulfill(envelope(read('statements.json'))));
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
      await page.route('**/v1/creators/application', (route) => route.fulfill(envelope({ status: null, created_at: null, creator: false })));
    },
  },
  {
    mock: 'web-creators',
    url: '/creators',
    ready: 'creators-section-apply',
    setup: async (page) => {
      await page.route('**/v1/me', (route) => route.fulfill(envelope(null, 401)));
    },
  },
];

test.describe('creator program mock parity', () => {
  test.skip(({ isMobile }) => isMobile, 'the parity mocks are desktop captures');

  let mocks = '';
  // eslint-disable-next-line no-empty-pattern -- Playwright reads a hook's fixtures from its pattern
  test.beforeAll(({}, testInfo) => {
    test.skip(testInfo.project.name !== 'desktop', 'the parity mocks are desktop captures');
    mocks = renderMocks();
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
