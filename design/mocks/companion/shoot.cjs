// design/mocks/companion/shoot.cjs: screenshots of the companion window mock at its real size (800 x 620, 2x), and
// the icon sheet and tray menu (icon_sheet.py) as full pages at 1100 wide.
// node design/mocks/companion/shoot.cjs <web/node_modules>
const path = require('node:path');
const fs = require('node:fs');
const { chromium } = require(path.join(process.argv[2], 'playwright'));

(async () => {
  const html = path.join(__dirname, 'html');
  const out = path.join(__dirname, 'png');
  fs.mkdirSync(out, { recursive: true });
  const browser = await chromium.launch();
  const SHEETS = new Set(['icons.html', 'tray-menu.html']);
  for (const file of fs.readdirSync(html).filter((f) => f.endsWith('.html')).sort()) {
    const sheet = SHEETS.has(file);
    const page = await browser.newPage({ viewport: { width: sheet ? 1100 : 800, height: sheet ? 400 : 620 }, deviceScaleFactor: 2 });
    await page.goto('file://' + path.join(html, file));
    await page.evaluate(() => document.fonts.ready);
    await page.screenshot({ path: path.join(out, file.replace('.html', '.png')), fullPage: sheet });
    await page.close();
    console.log('png/' + file.replace('.html', '.png'));
  }
  await browser.close();
})();
