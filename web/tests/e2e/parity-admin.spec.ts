// web/tests/e2e/parity-admin.spec.ts
// /admin/links against its mock (design/parity.md): the owner (Thandrel, fixtures/home/me-60.json, role
// admin) with the three links admin-links.spec.ts serves, at 1440; every region compared for its
// blocks and box (support/parity.ts). The mock is rendered before the test from the copy modules
// (design/mocks/billing_facts.mjs, then ws_billing.py). The side-by-side lands in
// test-results/parity/admin-links-1440.png. Desktop project only.
import { execFileSync } from 'node:child_process';
import { mkdirSync, readFileSync, rmSync } from 'node:fs';
import path from 'node:path';
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

const MOCK = 'web-admin-links';
const VIEWPORT = { width: 1440, height: 1000 };

/** Renders the billing lane's mocks from this checkout's copy modules into test-results/parity/mocks/billing. */
function renderBillingMocks(): string {
  const dir = path.join(PARITY_OUT, 'mocks', 'billing');
  rmSync(dir, { recursive: true, force: true });
  mkdirSync(dir, { recursive: true });
  const facts = path.join(dir, 'billing-facts.json');
  const generators = path.join(REPO_ROOT, 'design/mocks');
  const env = { ...process.env, FS_BILLING_FACTS: facts, FS_MOCK_OUT: dir };
  execFileSync('node', [path.join(generators, 'billing_facts.mjs'), facts], { cwd: REPO_ROOT, env });
  execFileSync(process.env.PYTHON ?? 'python3', [path.join(generators, 'ws_billing.py'), MOCK], {
    cwd: REPO_ROOT,
    env,
  });
  return dir;
}

const envelope = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ ok: true, data, error: null, request_id: 'admin' }),
});

const link = (slug: string, target: string, label: string, f: readonly number[]) => ({
  slug,
  target,
  label,
  created_at: '2026-11-02T00:00:00Z',
  clicks: f[0],
  bot_clicks: f[1],
  sign_ins: f[2],
  linked: f[3],
  premium: f[4],
});

/** The links admin-links.spec.ts serves, and billing_facts.mjs draws. */
const LINKS = [
  link('reddit', '/', 'Reddit', [40, 6, 10, 5, 1]),
  link('reddit-sim', '/sim', 'Reddit: sim post', [120, 2, 12, 9, 3]),
  link('wago', '/', 'Wago', [0, 0, 0, 0, 0]),
];

test.describe('/admin/links mock parity', () => {
  test.skip(({ isMobile }) => isMobile, 'the parity mocks are desktop captures');

  let mocks = '';
  // eslint-disable-next-line no-empty-pattern -- Playwright reads a hook's fixtures from its pattern
  test.beforeAll(({}, testInfo) => {
    test.skip(testInfo.project.name !== 'desktop', 'the parity mocks are desktop captures');
    mocks = renderBillingMocks();
  });

  test(`admin-links at 1440 matches ${MOCK}`, async ({ page, browser }) => {
    test.slow();
    await page.setViewportSize(VIEWPORT);
    const mockPage = await browser.newPage({ viewport: VIEWPORT });
    await openMock(mockPage, mocks, `${MOCK}.html`);
    const mock = await captureRegions(mockPage);

    const owner = JSON.parse(
      readFileSync(path.join(REPO_ROOT, 'web/tests/e2e/fixtures/home/me-60.json'), 'utf8'),
    ) as { user: Record<string, unknown> };
    await prepareParityPage(page);
    await page.context().addCookies([{ name: 'fs_csrf', value: 'tok', domain: 'localhost', path: '/' }]);
    await page.route('**/v1/me', (route) =>
      route.fulfill(envelope({ ...owner, user: { ...owner.user, role: 'admin' } })),
    );
    await page.route('**/v1/admin/golinks', (route) => route.fulfill(envelope({ links: LINKS })));
    await page.goto('/admin/links');
    await expect(page.getByTestId('links-row-reddit-sim')).toBeVisible();
    await settleParityPage(page);
    const built = await captureRegions(page);

    await writeSideBySide(
      browser,
      'admin-links-1440',
      await captureOutlined(mockPage, mock),
      await captureOutlined(page, built),
    );
    await mockPage.close();
    const failures = compareRegions(mock, built);
    expect(failures, `admin-links differs from ${MOCK}:\n${formatFailures(failures)}`).toEqual([]);
  });
});
