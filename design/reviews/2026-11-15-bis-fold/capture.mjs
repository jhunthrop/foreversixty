// Captures the BiS spec page with its folded Leveling section (2026-11-15) for the persona reviews:
//   cd web && npm run build && npx astro preview --port 4399
//   node design/reviews/2026-11-15-bis-fold/capture.mjs [http://localhost:4399]
// -> design/reviews/2026-11-15-bis-fold/build/<name>-<width>.png, full page, signed out, after the
// island has hydrated and the full data is in. Playwright is found through web/.
import { createRequire } from "node:module";
import { mkdirSync } from "node:fs";
import { dirname, join } from "node:path";
import { fileURLToPath } from "node:url";

const HERE = dirname(fileURLToPath(import.meta.url));
const ROOT = join(HERE, "../../..");
const require = createRequire(join(ROOT, "web/package.json"));
const { chromium } = require("@playwright/test");
const BASE = process.argv[2] ?? "http://localhost:4399";
const OUT = join(HERE, "build");
mkdirSync(OUT, { recursive: true });

const SHOTS = [
  { name: "fury", url: "/bis/warrior/fury", widths: [1440, 390] },
  { name: "frost", url: "/bis/mage/frost", widths: [1440, 390] },
  { name: "fury-horde", url: "/bis/warrior/fury?faction=horde", widths: [1440] },
  { name: "fury-hash-40", url: "/bis/warrior/fury#levels-40-49", widths: [1440] },
];

const browser = await chromium.launch();
for (const shot of SHOTS) {
  for (const width of shot.widths) {
    const page = await browser.newPage({ viewport: { width, height: 900 }, deviceScaleFactor: 1 });
    await page.route("**/v1/**", (route) => route.fulfill({ status: 401, body: "{}" }));
    await page.goto(`${BASE}${shot.url}`, { waitUntil: "networkidle" });
    await page.waitForFunction(() => !document.documentElement.hasAttribute("data-bis-pending"));
    // Walk the page so every lazy icon loads before the full-page shot, then back to the top.
    await page.evaluate(async () => {
      for (let y = 0; y < document.body.scrollHeight; y += 600) {
        window.scrollTo(0, y);
        await new Promise((resolve) => setTimeout(resolve, 60));
      }
      window.scrollTo(0, 0);
    });
    await page.waitForLoadState("networkidle");
    await page.waitForTimeout(800);
    const file = join(OUT, `${shot.name}-${width}.png`);
    await page.screenshot({ path: file, fullPage: true });
    console.log(`captured ${file} (${page.url()})`);
    await page.close();
  }
}
await browser.close();
