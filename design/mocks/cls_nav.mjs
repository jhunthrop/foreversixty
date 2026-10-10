// Layout-shift sources on a built page, via the PerformanceObserver (debugging aid for the nav).
//   node design/mocks/cls_nav.mjs /planner
import { createRequire } from 'node:module';

const require = createRequire(new URL('../../web/package.json', import.meta.url));
const { chromium } = require('@playwright/test');
const base = process.env.NAV_BASE ?? 'http://localhost:4325';
const browser = await chromium.launch();
const page = await (await browser.newContext({ viewport: { width: 1350, height: 940 } })).newPage();
await page.addInitScript(() => {
  window.__shifts = [];
  new PerformanceObserver((list) => {
    for (const e of list.getEntries()) {
      window.__shifts.push({
        value: e.value,
        sources: e.sources.map((s) => ({
          node: s.node ? (s.node.className || s.node.tagName) : null,
          prev: JSON.stringify(s.previousRect),
          cur: JSON.stringify(s.currentRect),
        })),
      });
    }
  }).observe({ type: 'layout-shift', buffered: true });
});
await page.goto(base + (process.argv[2] ?? '/'));
await page.waitForTimeout(3500);
console.log(JSON.stringify(await page.evaluate(() => window.__shifts), null, 1));
await browser.close();
