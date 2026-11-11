// design/mocks/2026-11-10-app-brand-polish/capture-web.mjs
// The website side of the app-brand-polish review: the live site's page heads at 390 px (iPhone
// width, 3x), to sit beside the app captures. Run from the repo root:
//   node design/mocks/2026-11-10-app-brand-polish/capture-web.mjs
import { createRequire } from 'node:module';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

const repo = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '..', '..', '..');
const { chromium } = createRequire(path.join(repo, 'web', 'package.json'))('playwright');
const OUT = path.join(repo, 'design/reviews/2026-11-10-app-brand-polish');
const SITE = 'https://foreversixty.gg';
const PAGES = { home: '/', bis: '/bis/warrior/fury', guides: '/guides' };

const browser = await chromium.launch();
const page = await browser.newPage({ viewport: { width: 390, height: 844 }, deviceScaleFactor: 3, colorScheme: 'dark' });
for (const [name, route] of Object.entries(PAGES)) {
  await page.goto(`${SITE}${route}`, { waitUntil: 'networkidle' });
  await page.waitForTimeout(1500);
  await page.screenshot({ path: path.join(OUT, `web-${name}-390.png`) });
}
await browser.close();
