// Renders each mock HTML at its body width, full page, to ../<name>.png. Run by gen.py from web/.
import path from 'node:path';
import { createRequire } from 'node:module';
import { pathToFileURL } from 'node:url';

// Resolved from the cwd (web/), where node_modules lives, not from this file's folder.
const { chromium } = createRequire(path.join(process.cwd(), 'package.json'))('playwright');
const browser = await chromium.launch();
for (const file of process.argv.slice(2)) {
  const width = Number((/-(\d+)\.html$/.exec(file) ?? [, 1280])[1]);
  // A tooltip board is a viewport capture (its placement depends on the viewport); the rest are full pages.
  const tip = path.basename(file).startsWith('tip-');
  const height = tip && width < 640 ? 844 : 800;
  const page = await browser.newPage({ viewport: { width, height }, deviceScaleFactor: 2 });
  await page.goto(pathToFileURL(file).href);
  await page.evaluate(() => document.fonts.ready);
  await page.waitForTimeout(150);
  const out = path.join(path.dirname(path.dirname(file)), path.basename(file, '.html') + '.png');
  await page.screenshot({ path: out, fullPage: !tip });
  await page.close();
  console.log(out);
}
await browser.close();
