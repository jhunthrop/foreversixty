// web/tests/captures/reports-trash/capture.spec.ts
// Review screenshots for the reports-trash lane: the /reports page ("Your reports") signed out,
// empty, listed and failed at 1440 and 390; the fixture's Fights card with its trash entry folded
// and opened; the combined trash view; a raid night with boss pulls; and the themed scrollbar on
// the sample's 79-pull fight list. Not a test of behaviour: it writes PNGs to
// design/reviews/2026-10-31-reports-trash/.
import { writeFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { chromium, test, type Page } from '@playwright/test';

const HERE = path.dirname(fileURLToPath(import.meta.url));
const OUT_DIR = path.resolve(HERE, '../../../../design/reviews/2026-10-31-reports-trash');
const WIDTHS = [1440, 390] as const;
const HEIGHT = { 1440: 1000, 390: 844 } as const;
const REPORT = '/reports/fixture2abcd';
const SETTLE_MS = 400;

const fulfil = (body: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify({
    ok: status < 400,
    data: status < 400 ? body : null,
    error: status < 400 ? null : { message: 'oops' },
    request_id: 'r',
  }),
});

const ME = {
  user: { id: 1, battletag: 'Hunthrop#1894', email: null, role: 'user', anonymize: false },
  characters: [],
  guilds: [],
};

const HOUR_MS = 60 * 60 * 1000;
const ROWS = [
  ['newreport2ab', 'Tuesday raid', 'Sanguine Depths', 20, 12, 3],
  ['midreport2ab', 'Ashen Spire clear', 'Ashen Spire', 50, 31, 6],
  ['oldreport2ab', '', 'Zephras Isle', 170, 79, 0],
].map(([id, title, zone, hoursAgo, fights, kills]) => ({
  id,
  title,
  zone,
  status: 'ready',
  visibility: 'public',
  created_at: new Date(Date.now() - Number(hoursAgo) * HOUR_MS).toISOString(),
  fight_count: fights,
  kill_count: kills,
  boss_count: kills,
  logging_character: null,
}));

type ReportsState = 'signed-out' | 'empty' | 'list' | 'error';

async function stubReports(page: Page, state: ReportsState): Promise<void> {
  await page.route('**/v1/me', (route) =>
    route.fulfill(state === 'signed-out' ? fulfil(null, 401) : fulfil(ME)),
  );
  await page.route('**/v1/reports?mine=1**', (route) =>
    route.fulfill(
      state === 'error'
        ? fulfil(null, 500)
        : fulfil({ rows: state === 'list' ? ROWS : [], total: 0, page: 1, per_page: 100 }),
    ),
  );
}

async function shot(page: Page, name: string): Promise<void> {
  await page.waitForTimeout(SETTLE_MS);
  await page.screenshot({ path: path.join(OUT_DIR, `${name}.png`), fullPage: true });
}

for (const width of WIDTHS) {
  for (const state of ['signed-out', 'empty', 'list', 'error'] as const) {
    test(`reports page ${state} at ${width}`, async ({ page }) => {
      await page.setViewportSize({ width, height: HEIGHT[width] });
      await stubReports(page, state);
      await page.goto('/reports');
      await page.getByTestId('your-reports').waitFor();
      await shot(page, `reports-${state}-${width}`);
    });
  }
}

test('fights card: trash entry folded, then opened', async ({ page }) => {
  await page.setViewportSize({ width: 1440, height: 1000 });
  await page.goto(REPORT);
  const card = page.getByTestId('fight-selector');
  await page.getByTestId('fight-trash-1').waitFor();
  await page.waitForTimeout(SETTLE_MS);
  await card.screenshot({ path: path.join(OUT_DIR, 'fights-trash-collapsed-1440.png') });
  await page.getByTestId('fight-group-1').click();
  await page.waitForTimeout(SETTLE_MS);
  await card.screenshot({ path: path.join(OUT_DIR, 'fights-trash-expanded-1440.png') });
});

for (const width of WIDTHS) {
  test(`combined trash view at ${width}`, async ({ page }) => {
    await page.setViewportSize({ width, height: HEIGHT[width] });
    await page.goto(`${REPORT}?fight=trash-1`);
    await page.getByTestId('night-view').waitFor();
    await shot(page, `trash-combined-${width}`);
  });

  test(`raid night with boss pulls at ${width}`, async ({ page }) => {
    await page.setViewportSize({ width, height: HEIGHT[width] });
    await page.goto(REPORT);
    await page.getByTestId('summary-tab').waitFor();
    await shot(page, `raid-night-${width}`);
  });
}

test('trash combined view, damage done table', async ({ page }) => {
  await page.setViewportSize({ width: 1440, height: 1000 });
  await page.goto(`${REPORT}?fight=trash-1&tab=damage-done`);
  await page.getByTestId('actor-table').first().waitFor();
  await shot(page, 'trash-combined-damage-1440');
});

// Headless Chromium hides scrollbars, and a Mac on a trackpad draws overlay bars only while
// scrolling. This launch keeps them and hands the browser macOS's own per-process
// `-AppleShowScrollBars Always` (an argument-domain default, so no system setting changes): the
// bar a Mac with a mouse plugged in, or "Show scroll bars: Always", draws. Playwright refuses a
// bare "Always" argument, so a two-line wrapper passes it.
test('the sample’s fight list with its scrollbar showing', async () => {
  const wrapper = test.info().outputPath('chrome-classic-scrollbars.sh');
  writeFileSync(
    wrapper,
    `#!/bin/sh\nexec "${chromium.executablePath()}" "$@" -AppleShowScrollBars Always\n`,
    {
      mode: 0o755,
    },
  );
  const browser = await chromium.launch({
    executablePath: wrapper,
    ignoreDefaultArgs: ['--hide-scrollbars'],
  });
  const page = await browser.newPage({ viewport: { width: 1440, height: 1000 }, deviceScaleFactor: 2 });
  const port = process.env.CAPTURE_PORT ?? '4707';
  await page.goto(`http://localhost:${port}/reports/sample`);
  const card = page.getByTestId('fight-selector');
  await card.waitFor();
  await page.waitForTimeout(SETTLE_MS);
  await card.screenshot({ path: path.join(OUT_DIR, 'scrollbar-fight-list-1440.png') });
  await card.locator('ul').first().hover();
  await page.waitForTimeout(SETTLE_MS);
  await page.screenshot({ path: path.join(OUT_DIR, 'scrollbar-page-1440.png') });
  await browser.close();
});
