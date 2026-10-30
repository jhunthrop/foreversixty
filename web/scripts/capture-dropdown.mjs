// web/scripts/capture-dropdown.mjs
// The dropdown lane's acceptance captures (design/specs/2026-10-30-controls-charts.md §D), taken
// from a running fixture build: every changed surface at 1440, 2000 and 390, with the list open
// on a fine pointer and the native field on the phone.
// Usage: node scripts/capture-dropdown.mjs [base-url]   (default http://localhost:4701)
import { mkdir, readFile } from 'node:fs/promises';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { chromium, devices } from 'playwright';

const BASE = process.argv[2] ?? 'http://localhost:4701';
const OUT = path.resolve(
  path.dirname(fileURLToPath(import.meta.url)),
  '../../design/reviews/2026-10-30-controls-charts',
);
const WIDTHS = [1440, 2000, 390];
const { build } = JSON.parse(
  await readFile(
    path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../src/data/active-build.json'),
    'utf8',
  ),
);
/** The Fury export sim-settings.spec.ts loads: the sim's settings draw once a character runs. */
const FURY = `FS1:${build}:warrior:orc:0/5530515/0:head=12640,main_hand=11726`;

async function loadFury(page) {
  await page.getByTestId('sim-cold-paste-input').fill(FURY);
  await page.getByTestId('sim-cold-paste-run').click();
  await page.getByTestId('sim-run-button').filter({ hasText: 'Run again' }).waitFor({ timeout: 20_000 });
}

/** Opens the field whose select carries `testid`, once the listbox has enhanced it. */
async function open(page, testid) {
  const trigger = page.locator(`[data-testid="${testid}-trigger"]`);
  await trigger.waitFor();
  await trigger.scrollIntoViewIfNeeded();
  await page.evaluate(() => window.scrollBy(0, -160));
  await trigger.click();
  await page.locator(`.sel:has(> select[data-testid="${testid}"]) > .sel-pop`).waitFor();
}

const SURFACES = [
  // The phone captures show each surface's closed native fields; the list opens on a fine pointer.
  { name: 'sheet', url: '/design-sheet', ready: 'sheet-fight-style', full: true },
  {
    name: 'report-source',
    url: '/reports/fixture2abcd?fight=3&tab=threat',
    ready: 'source-scope',
    act: (page) => open(page, 'source-scope'),
  },
  {
    name: 'report-more',
    url: '/reports/fixture2abcd?fight=3&tab=threat',
    ready: 'tab-more',
    act: async (page) => {
      await page.getByTestId('tab-more').click();
      await page.getByTestId('tab-more-menu').waitFor();
    },
  },
  {
    name: 'sim-style',
    url: '/sim',
    ready: 'sim-cold-paste-input',
    prepare: loadFury,
    act: (page) => open(page, 'sim-style'),
  },
  { name: 'rankings', url: '/rankings', ready: 'filter-region', act: (page) => open(page, 'filter-region') },
  { name: 'home-search', url: '/', ready: 'home-character-search-submit' },
  {
    name: 'planner',
    url: '/planner',
    ready: 'planner-race-select',
    act: (page) => open(page, 'planner-race-select'),
  },
];

await mkdir(OUT, { recursive: true });
const browser = await chromium.launch();
for (const width of WIDTHS) {
  const phone = width < 768;
  const context = await browser.newContext(
    phone
      ? { ...devices['Pixel 7'], viewport: { width, height: 844 } }
      : { viewport: { width, height: 1000 } },
  );
  for (const surface of SURFACES) {
    const page = await context.newPage();
    await page.goto(`${BASE}${surface.url}`, { waitUntil: 'networkidle' });
    await page.getByTestId(surface.ready).first().waitFor({ state: 'attached' });
    if (surface.prepare !== undefined) await surface.prepare(page);
    if (!phone && surface.act !== undefined) await surface.act(page);
    await page.waitForTimeout(250);
    const file = path.join(OUT, `build-dropdown-${surface.name}-${width}.png`);
    await page.screenshot({ path: file, fullPage: surface.full === true });
    console.log(file);
    await page.close();
  }
  await context.close();
}
await browser.close();
