// web/tests/e2e/parity-app-beta.spec.ts
// /app and /admin/testers against their mocks (design/parity.md; ws_app_beta.py): /app in every
// signup state (idle, signed up, already on the list, a refused address, a failure) at 1440 and
// 390, where the QR rail is gone on both sides; /admin/testers for the owner at 1440. Every
// region compared for its blocks and box, plus what lies outside them. The mocks are rendered
// before the tests from this checkout's copy modules and QR (app_beta_facts.mjs).
import { readFileSync } from 'node:fs';
import path from 'node:path';
import type { Page } from '@playwright/test';
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
  type RegionSnapshot,
} from './support/parity';
import { renderAppBetaMocks } from './support/parity-mocks';

const envelope = (status: number, data: unknown, fields: Record<string, string> | null = null) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify({
    ok: status < 400,
    data: status < 400 ? data : null,
    error: status < 400 ? null : { code: 'invalid', message: 'refused', fields },
    request_id: 'parity',
  }),
});

/** How the build reaches each state: the API's answer and the address typed, or nothing. */
const STATES = [
  { id: 'idle', mock: 'web-app', send: null },
  { id: 'done', mock: 'web-app-done', send: { status: 201, email: 'tester@example.com' } },
  { id: 'exists', mock: 'web-app-exists', send: { status: 200, email: 'tester@example.com' } },
  { id: 'invalid', mock: 'web-app-invalid', send: { status: 201, email: 'not-an-email' } },
  { id: 'failed', mock: 'web-app-failed', send: { status: 500, email: 'tester@example.com' } },
] as const;
const WIDTHS = [1440, 390] as const;
const HEIGHT = 1000;
const MOCK_HEAD_INSET =
  'below 768 the mock puts the 18 px page inset outside the head region (.pc), PageHeader inside it; ' +
  'the same drift parity-app-auth.spec.ts lists';
const PHONE_DRIFT = ['head:box'];

async function compareWith(
  page: Page,
  mockPage: Page,
  mock: RegionSnapshot[],
  name: string,
  drift: readonly string[],
): Promise<void> {
  const built = await captureRegions(page);
  await writeSideBySide(
    page.context().browser()!,
    name,
    await captureOutlined(mockPage, mock),
    await captureOutlined(page, built),
  );
  const failures = compareRegions(mock, built);
  const unexpected = failures.filter((f) => !drift.includes(`${f.region}:${f.kind}`));
  expect(unexpected, `${name} differs from its mock:\n${formatFailures(unexpected)}`).toEqual([]);
  expect(
    new Set(
      failures.filter((f) => drift.includes(`${f.region}:${f.kind}`)).map((f) => `${f.region}:${f.kind}`),
    ).size,
    `${name}: the drift (${MOCK_HEAD_INSET}) is gone; delete its entry`,
  ).toBe(drift.length);
}

test.describe('/app and /admin/testers mock parity', () => {
  test.skip(({ isMobile }) => isMobile, 'the parity mocks are desktop captures');

  let mocks = '';
  // eslint-disable-next-line no-empty-pattern -- Playwright reads a hook's fixtures from its pattern
  test.beforeAll(({}, testInfo) => {
    test.skip(testInfo.project.name !== 'desktop', 'the parity mocks are desktop captures');
    mocks = renderAppBetaMocks();
  });

  for (const width of WIDTHS) {
    for (const { id, mock: mockName, send } of STATES) {
      test(`app-${id} at ${width} matches ${mockName}`, async ({ page, browser }) => {
        const viewport = { width, height: HEIGHT };
        await page.setViewportSize(viewport);
        const mockPage = await browser.newPage({ viewport });
        await openMock(mockPage, mocks, `${mockName}.html`);
        const mock = await captureRegions(mockPage);

        await prepareParityPage(page);
        await page.route('**/v1/me', (route) => route.fulfill({ status: 401, body: '' }));
        await page.route('**/v1/app-testers', (route) =>
          route.fulfill(
            envelope(send?.status ?? 201, { status: send?.status === 200 ? 'exists' : 'created' }),
          ),
        );
        await page.goto('/app');
        await settleParityPage(page);
        if (send !== null) {
          await page.getByTestId('app-beta-email').fill(send.email);
          await page.getByTestId('app-beta-submit').click();
          const shown =
            id === 'failed' ? 'app-beta-failed' : id === 'invalid' ? 'app-beta-email' : 'app-beta-done';
          if (id === 'invalid') await expect(page.getByTestId(shown)).toHaveAttribute('aria-invalid', 'true');
          else await expect(page.getByTestId(shown)).toBeVisible();
          await page.getByTestId('app-beta-email').or(page.getByTestId('app-beta-done')).first().blur();
        }
        await compareWith(page, mockPage, mock, `app-${id}-${width}`, width < 768 ? PHONE_DRIFT : []);
        await mockPage.close();
      });
    }
  }

  test('admin-testers at 1440 matches web-admin-testers', async ({ page, browser }) => {
    const viewport = { width: 1440, height: HEIGHT };
    await page.setViewportSize(viewport);
    const mockPage = await browser.newPage({ viewport });
    await openMock(mockPage, mocks, 'web-admin-testers.html');
    const mock = await captureRegions(mockPage);

    const owner = JSON.parse(
      readFileSync(path.join(REPO_ROOT, 'web/tests/e2e/fixtures/home/me-60.json'), 'utf8'),
    ) as {
      user: Record<string, unknown>;
    };
    const testers = JSON.parse(
      readFileSync(path.join(REPO_ROOT, 'web/tests/e2e/fixtures/app-beta/testers.json'), 'utf8'),
    ) as unknown;
    await prepareParityPage(page);
    await page.context().addCookies([{ name: 'fs_csrf', value: 'tok', domain: 'localhost', path: '/' }]);
    await page.route('**/v1/me', (route) =>
      route.fulfill(envelope(200, { ...owner, user: { ...owner.user, role: 'admin' } })),
    );
    await page.route('**/v1/admin/app-testers', (route) => route.fulfill(envelope(200, { testers })));
    await page.goto('/admin/testers');
    await expect(page.getByTestId('testers-grid')).toBeVisible();
    await settleParityPage(page);
    await compareWith(page, mockPage, mock, 'admin-testers-1440', []);
    await mockPage.close();
  });
});
