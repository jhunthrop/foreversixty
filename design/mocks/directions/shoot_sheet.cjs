// Captures one local study sheet full page: node design/mocks/directions/shoot_sheet.cjs <node_modules> <in.html> <out.png>
const path = require("node:path");
const { chromium } = require(path.join(process.argv[2], "playwright-core"));
const [, , , src, out] = process.argv;

(async () => {
  const browser = await chromium.launch();
  const pg = await browser.newPage({ viewport: { width: 1300, height: 900 } });
  await pg.goto(`file://${path.resolve(src)}`, { waitUntil: "load", timeout: 90000 });
  await pg.evaluate(() => document.fonts.ready);
  await pg.waitForTimeout(1500);
  await pg.screenshot({ path: path.resolve(out), fullPage: true });
  await browser.close();
  console.log("ok", out);
})();
