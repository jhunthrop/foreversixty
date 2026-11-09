// Captures the live site as the "Now" column of the directions gallery (design/specs/2026-11-09-directions.md).
//   node design/mocks/directions/shoot_now.cjs <a web/node_modules that has playwright-core>
// -> design/mocks/directions/now/<page>-<width>.png (first screen) and <page>-<width>-full.png
const path = require("node:path");
const fs = require("node:fs");
const { chromium } = require(path.join(process.argv[2], "playwright-core"));

const OUT = path.join(__dirname, "now");
const BASE = "https://foreversixty.gg";
const PAGES = { home: "/", bis: "/bis/hunter/marksmanship", tiers: "/tier-list", guide: "/guides/warrior/fury", sim: "/sim", planner: "/planner" };
const WIDTHS = [1440, 390];

async function shoot(ctx, name, route, width) {
  const pg = await ctx.newPage();
  try {
    await pg.goto(BASE + route, { waitUntil: "networkidle", timeout: 45000 });
    await pg.waitForTimeout(2500);
    await pg.screenshot({ path: path.join(OUT, `${name}-${width}.png`) });
    await pg.screenshot({ path: path.join(OUT, `${name}-${width}-full.png`), fullPage: true });
    console.log("ok", name, width, pg.url());
  } catch (e) {
    console.error("fail", name, width, e.message);
  } finally {
    await pg.close();
  }
}

(async () => {
  fs.mkdirSync(OUT, { recursive: true });
  const browser = await chromium.launch();
  for (const width of WIDTHS) {
    const ctx = await browser.newContext({ viewport: { width, height: width < 500 ? 844 : 900 } });
    for (const [name, route] of Object.entries(PAGES)) await shoot(ctx, name, route, width);
    await ctx.close();
  }
  await browser.close();
})();
