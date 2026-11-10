// web/tests/e2e/parity-b.spec.ts
// The shared build /b/:id against its mocks (design/parity.md, design/specs/2026-11-08-premium-account.md
// sections 8 and 14, gap G7): a stranger signed out on the Survival raid build (web-b) and an
// unknown id (web-b-missing), and the static build-unavailable page (web-b-unavailable), at 1440 and 390, each region (head, head-title, head-meta,
// head-answer, notes, plate, every Section) compared for its blocks and box (support/parity.ts).
//
// The API renders the page, so the build side is the API's own HTML: fixtures/shared-build/
// survival.html and missing.html, written by api/internal/site/page_test.go
// (`go test ./internal/site -run TestSharedBuildParityFixtures -update`) from the mock's record
// (api/internal/site/testdata/shared-survival.json: planner-facts.json `guide`, titled and viewed
// as web-b). It is served at /b/<id> with the island and its stylesheet from this build's dist.
// The API's chrome.html is a copy of an earlier build's header and footer (its scripts load by
// stable /_chrome/ path, scripts/chrome-aliases.mjs, but its markup can lag): the test puts this
// build's own chrome in its place (dist/about.html, the file chrome.html is copied from), so the
// header matches the site it is compared with.
// The mocks are rendered before the tests (support/parity-mocks.ts). A side-by-side per state
// and width lands in test-results/parity/b-<state>-<width>.png. Desktop project only: the 390 run
// is a 390 px desktop viewport, as the mock renders are.
import { readFileSync } from 'node:fs';
import path from 'node:path';
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
import { renderSharedBuildMocks } from './support/parity-mocks';
import { routePlannerState } from './support/planner-page';

const WEB_ROOT = path.join(import.meta.dirname, '..', '..');
const FIXTURES = path.join(import.meta.dirname, 'fixtures', 'shared-build');

/** The first site <header> and the last <footer> of a page (api/internal/site chrome_test.go's rule). */
function chromeOf(html: string): { header: string; footer: string } {
  const header = html.slice(html.indexOf('<header'), html.indexOf('</header>') + '</header>'.length);
  const footer = html.slice(html.lastIndexOf('<footer'), html.lastIndexOf('</footer>') + '</footer>'.length);
  return { header, footer };
}

/** The API page with this build's chrome in place of the copy it was rendered with. */
function apiPage(fixture: string): string {
  const page = readFileSync(path.join(FIXTURES, fixture), 'utf8');
  const served = chromeOf(page);
  const site = chromeOf(readFileSync(path.join(WEB_ROOT, 'dist', 'about.html'), 'utf8'));
  return page.replace(served.header, site.header).replace(served.footer, site.footer);
}

const RECORD = JSON.parse(
  readFileSync(path.join(WEB_ROOT, '..', 'api/internal/site/testdata/shared-survival.json'), 'utf8'),
) as { id: string };

async function openShared(page: Page, fixture: string, id: string): Promise<void> {
  await routePlannerState(page, 'out');
  const body = apiPage(fixture);
  await page.route(`**/b/${id}`, (route) => route.fulfill({ status: 200, contentType: 'text/html', body }));
  await page.goto(`/b/${id}`);
}

const STATES: readonly { id: string; mock: string; open: (page: Page) => Promise<void> }[] = [
  {
    id: 'b',
    mock: 'web-b',
    open: async (page) => {
      await openShared(page, 'survival.html', RECORD.id);
      await expect(page.getByTestId('planner-trees')).toBeVisible();
      await expect(page.getByTestId('band-compare-differs')).toBeVisible();
      await page.waitForLoadState('networkidle');
    },
  },
  {
    id: 'b-missing',
    mock: 'web-b-missing',
    open: async (page) => {
      await openShared(page, 'missing.html', 'nosuchbuild');
      await expect(page.getByTestId('page-title')).toHaveText('Build not found');
    },
  },
  {
    // The static page src/worker.ts serves with a 503 when the API cannot answer /b/:id.
    id: 'b-unavailable',
    mock: 'web-b-unavailable',
    open: async (page) => {
      await routePlannerState(page, 'out');
      await page.goto('/b-unavailable');
      await expect(page.getByTestId('page-title')).toHaveText('Build unavailable');
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
  'a primitive/mock-library question, the same drift parity-planner.spec.ts and parity-setup.spec.ts list';
const EXPECTED_DRIFT: Readonly<Record<string, { failures: readonly string[]; why: string }>> = {
  'b-390': { failures: ['head:box'], why: MOCK_HEAD_INSET },
  'b-missing-390': { failures: ['head:box'], why: MOCK_HEAD_INSET },
  'b-unavailable-390': { failures: ['head:box'], why: MOCK_HEAD_INSET },
};

test.describe('/b/:id mock parity', () => {
  test.skip(({ isMobile }) => isMobile, 'the parity mocks are desktop captures');

  let mocks = '';
  // eslint-disable-next-line no-empty-pattern -- Playwright reads a hook's fixtures from its pattern
  test.beforeAll(({}, testInfo) => {
    test.skip(testInfo.project.name !== 'desktop', 'the parity mocks are desktop captures');
    mocks = renderSharedBuildMocks(STATES.map((state) => state.mock));
  });

  for (const width of WIDTHS) {
    for (const { id, mock: mockName, open } of STATES) {
      test(`${id} at ${width} matches ${mockName}`, async ({ page, browser }) => {
        test.slow();
        const viewport = { width, height: HEIGHT };
        await page.setViewportSize(viewport);
        const mockPage = await browser.newPage({ viewport });
        await openMock(mockPage, mocks, `${mockName}.html`);
        const mock = await captureRegions(mockPage);

        await prepareParityPage(page);
        await open(page);
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
