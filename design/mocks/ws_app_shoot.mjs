// Captures the native app mocks (design/mocks/app/<platform>-<screen>.html, written by ws_app.py) at the device sizes:
//   node design/mocks/ws_app_shoot.mjs [name-filter]
// -> design/reviews/2026-11-07-mobile/<platform>-<screen>.png: the whole scroll at 1x (393 wide for iOS, 412 for Android).
// Playwright is found through web/. Fonts are the repo's own (app_fonts.mjs), so captures match offline and in CI.
import { createRequire } from "node:module";
import { mkdirSync, readdirSync } from "node:fs";
import { routeFonts } from "./app_fonts.mjs";

// PLAYWRIGHT_FROM: another package.json whose node_modules carries @playwright/test (a checkout without web/node_modules).
const require = createRequire(process.env.PLAYWRIGHT_FROM ?? new URL("../../web/package.json", import.meta.url));
const { chromium } = require("@playwright/test");
const SRC = new URL("./app/", import.meta.url).pathname;
// SHOOT_OUT: another folder for the captures (a review round of its own, e.g. design/reviews/2026-11-08-app-flows).
const OUT = process.env.SHOOT_OUT ? `${process.env.SHOOT_OUT.replace(/\/$/, "")}/` : new URL("../reviews/2026-11-07-mobile/", import.meta.url).pathname;
const DEVICE = { ios: { width: 393, height: 852 }, android: { width: 412, height: 915 } };
const filter = process.argv[2] ?? "";
mkdirSync(OUT, { recursive: true });

const names = readdirSync(SRC).filter((f) => f.endsWith(".html") && f.includes(filter));
const browser = await chromium.launch();
const contexts = Object.fromEntries(
  await Promise.all(
    Object.entries(DEVICE).map(async ([p, viewport]) => [
      p,
      await (async () => {
        const c = await browser.newContext({ viewport, deviceScaleFactor: 1, reducedMotion: "reduce" });
        await routeFonts(c);
        return c;
      })(),
    ]),
  ),
);
for (const file of names) {
  const platform = file.startsWith("android-") ? "android" : "ios";
  const pg = await contexts[platform].newPage();
  await pg.goto(`file://${SRC}${file}`, { waitUntil: "networkidle" });
  await pg.evaluate(() => Promise.all([...document.images].map((i) => i.decode().catch(() => null))));
  await pg.waitForTimeout(150);
  await pg.screenshot({ path: `${OUT}${file.replace(/\.html$/, ".png")}`, fullPage: true });
  await pg.close();
}
await browser.close();
console.log(`${names.length} captures in ${OUT}`);
