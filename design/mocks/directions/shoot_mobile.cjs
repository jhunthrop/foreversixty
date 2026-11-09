// Renders the round-5 mobile mocks: node design/mocks/directions/shoot_mobile.cjs <node_modules with playwright-core> [name ...]
// -> design/mocks/directions/png/m-<platform>-<screen>.png at the device size (iPhone 15 393 x 852, Pixel 7 412 x 915);
//    the four tabs full length, the overlays (menu, sheets, lock screen) at one screen.
const path = require("node:path");
const fs = require("node:fs");
const { chromium } = require(path.join(process.argv[2], "playwright-core"));

const SRC = path.join(__dirname, "html");
const OUT = path.join(__dirname, "png");
const SIZE = { ios: [393, 852], android: [412, 915] };
const only = process.argv.slice(3);

(async () => {
  const names = fs.readdirSync(SRC).filter((f) => f.startsWith("m-") && f.endsWith(".html")).map((f) => f.slice(0, -5));
  const browser = await chromium.launch();
  for (const name of names.filter((n) => !only.length || only.includes(n))) {
    const [w, h] = SIZE[name.split("-")[1]];
    const pg = await browser.newPage({ viewport: { width: w, height: h } });
    await pg.goto(`file://${path.join(SRC, name + ".html")}`, { waitUntil: "load", timeout: 90000 });
    await pg.evaluate(() => document.fonts.ready);
    await pg.waitForTimeout(300);
    await pg.screenshot({ path: path.join(OUT, `${name}.png`), fullPage: true });
    await pg.close();
    console.log("ok", name);
  }
  await browser.close();
})();
