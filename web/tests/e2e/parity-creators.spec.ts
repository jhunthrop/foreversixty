// web/tests/e2e/parity-creators.spec.ts
// The public /creators page against its mock (design/parity.md; ws_creators.py, the parity twin of the
// directions mock b4-creators), signed out, at 1440 and 390: every region (head, the lead, each section)
// compared for its blocks and box (support/parity.ts). The mock is rendered before the test from
// lib/creators/copy.ts (support/parity-mocks.ts renderCreatorsMocks). A side-by-side per width lands in
// test-results/parity/creators-<width>.png. Also: Apply is the support email, at the head and the end.
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
import { renderCreatorsMocks } from './support/parity-mocks';
import { openPagesState } from './support/pages-page';
import { CREATORS_APPLY_HREF } from '../../src/lib/creators/copy';

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

test('Apply at the head and at the end is the support email', async ({ page }) => {
  await openPagesState(page, '/creators', 'out');
  await expect(page.getByTestId('creators-apply-top')).toHaveAttribute('href', CREATORS_APPLY_HREF);
  await expect(page.getByTestId('creators-apply-end')).toHaveAttribute('href', CREATORS_APPLY_HREF);
  await expect(page.getByTestId('creators-terms')).toBeVisible();
});
