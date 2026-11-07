// web/tests/e2e/parity-app-auth.spec.ts
// /app/auth against its mocks (design/parity.md, design/specs/2026-11-07-bis-landing.md section
// 7): with no code in the address bar (the start-again line) and with one (the Open button), at
// 1440 and 390, each region (head, head-title, head-meta) compared for its blocks and box, plus
// what lies outside them. Mocks rendered before the tests by ws_app_auth.py from the copy module.
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
import { renderAppAuthMocks } from './support/parity-mocks';

const STATES = [
  { id: 'app-auth', path: '/app/auth', mock: 'web-app-auth' },
  { id: 'app-auth-code', path: '/app/auth?code=kq47tm&state=s1', mock: 'web-app-auth-code' },
] as const;
const WIDTHS = [1440, 390] as const;
const HEIGHT = 1000;
const MOCK_HEAD_INSET =
  'below 768 the mock puts the 18 px page inset outside the head region (.pc), PageHeader inside it; ' +
  'the same drift parity-b.spec.ts lists';
const EXPECTED_DRIFT: Readonly<Record<string, readonly string[]>> = {
  'app-auth-390': ['head:box'],
  'app-auth-code-390': ['head:box'],
};

test.describe('/app/auth mock parity', () => {
  test.skip(({ isMobile }) => isMobile, 'the parity mocks are desktop captures');

  let mocks = '';
  // eslint-disable-next-line no-empty-pattern -- Playwright reads a hook's fixtures from its pattern
  test.beforeAll(({}, testInfo) => {
    test.skip(testInfo.project.name !== 'desktop', 'the parity mocks are desktop captures');
    mocks = renderAppAuthMocks();
  });

  for (const width of WIDTHS) {
    for (const { id, path, mock: mockName } of STATES) {
      test(`${id} at ${width} matches ${mockName}`, async ({ page, browser }) => {
        const viewport = { width, height: HEIGHT };
        await page.setViewportSize(viewport);
        const mockPage = await browser.newPage({ viewport });
        await openMock(mockPage, mocks, `${mockName}.html`);
        const mock = await captureRegions(mockPage);

        await prepareParityPage(page);
        await page.route('**/v1/me', (route) => route.fulfill({ status: 401, body: '' }));
        await page.goto(path);
        await settleParityPage(page);
        const built = await captureRegions(page);

        await writeSideBySide(
          browser,
          `${id}-${width}`,
          await captureOutlined(mockPage, mock),
          await captureOutlined(page, built),
        );
        await mockPage.close();

        const drift = EXPECTED_DRIFT[`${id}-${width}`] ?? [];
        const failures = compareRegions(mock, built);
        const unexpected = failures.filter((f) => !drift.includes(`${f.region}:${f.kind}`));
        expect(
          unexpected,
          `${id} at ${width} differs from ${mockName}:\n${formatFailures(unexpected)}`,
        ).toEqual([]);
        expect(
          new Set(
            failures
              .filter((f) => drift.includes(`${f.region}:${f.kind}`))
              .map((f) => `${f.region}:${f.kind}`),
          ).size,
          `${id} at ${width}: the drift (${MOCK_HEAD_INSET}) is gone; delete its entry`,
        ).toBe(drift.length);
      });
    }
  }
});
