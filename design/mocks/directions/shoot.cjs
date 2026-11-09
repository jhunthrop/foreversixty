// Renders the direction mocks: node design/mocks/directions/shoot.cjs <node_modules with playwright-core> [name ...]
// -> design/mocks/directions/png/<name>-1440.png and -390.png (full page), plus -2000.png for the home mocks.
const path = require("node:path");
const fs = require("node:fs");
const { chromium } = require(path.join(process.argv[2], "playwright-core"));

const SRC = path.join(__dirname, "html");
const OUT = path.join(__dirname, "png");
const only = process.argv.slice(3);

function widthsFor(name) {
  return name.endsWith("-home") ? [1440, 390, 2000] : [1440, 390];
}

(async () => {
  fs.mkdirSync(OUT, { recursive: true });
  // og-home.html is the 1200 x 630 link preview, captured on its own by shoot_sheet.cjs (art_crops.py)
  const names = fs.readdirSync(SRC).filter((f) => f.endsWith(".html") && !f.startsWith("og-")).map((f) => f.slice(0, -5));
  const browser = await chromium.launch();
  for (const name of names.filter((n) => !only.length || only.includes(n))) {
    for (const width of widthsFor(name)) {
      const pg = await browser.newPage({ viewport: { width, height: width < 500 ? 844 : 900 } });
      await pg.goto(`file://${path.join(SRC, name + ".html")}`, { waitUntil: "load", timeout: 90000 });
      await pg.evaluate(() => document.fonts.ready);
      await pg.waitForTimeout(300);
      await pg.screenshot({ path: path.join(OUT, `${name}-${width}.png`), fullPage: true });
      await pg.close();
      console.log("ok", name, width);
    }
  }
  await browser.close();
})();
