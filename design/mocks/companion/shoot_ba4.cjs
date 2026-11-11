// design/mocks/companion/shoot_ba4.cjs: the companion surfaces shoot_ui.cjs does not cover, for lane BA4's review
// (design/reviews/2026-11-11-ba4): the setup steps (companion/ui/install.html in each mode, at the window's fixed
// 660 x 400), the tray menu as macOS and Windows draw it (its words read from the generated companion/ui/copy_gen.js,
// so the capture shows the keys the binary ships, and the tray dot from internal/icon/assets). The disk image's
// window is a real Finder capture of a create-dmg image built with package-mac-companion.sh's arguments (dmg-window.png).
//   node design/mocks/companion/shoot_ba4.cjs <node_modules with playwright> <out dir>
const path = require('node:path');
const fs = require('node:fs');
const { chromium } = require(path.join(process.argv[2], 'playwright'));

const REPO = path.resolve(__dirname, '../../..');
const UI = path.join(REPO, 'companion/ui');
const OUT = path.resolve(process.argv[3]);
const ORIGIN = 'http://127.0.0.1:47123/9f2c1a7e/';
const TYPES = { '.html': 'text/html', '.js': 'text/javascript', '.css': 'text/css', '.png': 'image/png', '.webp': 'image/webp', '.woff2': 'font/woff2', '.svg': 'image/svg+xml' };
const TEXT = require(path.join(UI, 'copy_gen.js'));
const dataUrl = (file) => 'data:image/png;base64,' + fs.readFileSync(file).toString('base64');

// The install window's states: the mode and what the Go side calls on it.
const INSTALL = [
  { name: 'setup', mode: 'setup' },
  { name: 'setup-upgrade', mode: 'setup', info: { upgrade: true, startup: true } },
  { name: 'setup-installing', mode: 'setup', act: async (p) => { await p.click('#move'); await p.evaluate(() => window.fsProgress(3, 5, 'shortcut')); } },
  { name: 'setup-done', mode: 'setup', act: async (p) => { await p.click('#move'); await p.evaluate(() => window.fsDone()); } },
  { name: 'setup-failed', mode: 'setup', act: async (p) => { await p.click('#move'); await p.evaluate(() => window.fsFailed()); } },
  { name: 'uninstall', mode: 'uninstall' },
  { name: 'move', mode: 'move', info: { loginItems: true } },
];

async function shootInstall(browser, s) {
  const context = await browser.newContext({ viewport: { width: 660, height: 400 }, deviceScaleFactor: 2, locale: 'en-US' });
  const page = await context.newPage();
  await page.addInitScript((info) => {
    window.fsSetupInfo = async () => info;
    window.fsInstallInfo = async () => info;
    window.fsSetupRun = () => {};
    window.fsInstallMove = () => new Promise(() => {});
  }, s.info || null);
  await page.route(ORIGIN + '**', (route) => {
    const rel = route.request().url().slice(ORIGIN.length).split('?')[0];
    const file = path.join(UI, rel);
    if (!file.startsWith(UI) || !fs.existsSync(file)) return route.fulfill({ status: 404, body: '' });
    return route.fulfill({ body: fs.readFileSync(file), contentType: TYPES[path.extname(file)] });
  });
  await page.goto(ORIGIN + 'install.html?mode=' + s.mode);
  await page.evaluate(() => document.fonts.ready);
  await page.waitForTimeout(300);
  if (s.act) await s.act(page);
  await page.waitForTimeout(500);
  const file = path.join(OUT, `install-${s.name}.png`);
  await page.screenshot({ path: file });
  await context.close();
  return file;
}

// The tray menu as each system draws a native menu; the words are the generated keys.
function trayHtml() {
  const icon = dataUrl(path.join(REPO, 'companion/internal/icon/assets/seal-tray-on.png'));
  const line = 'Logging';
  const items = (hover) => `<div class="dis">${line}</div><hr><div class="${hover === 0 ? 'hov' : ''}">${TEXT.trayOpen}</div><div class="${hover === 1 ? 'hov' : ''}">${TEXT.reportsOpen}</div><hr><div>${TEXT.trayQuit}</div>`;
  return `<!doctype html><html><head><meta charset="utf-8"><style>
body { margin: 0; padding: 28px; background: #2a2d33; font: 14px -apple-system, 'Segoe UI', sans-serif; color: #eee; display: flex; gap: 48px; }
h2 { font: 600 13px -apple-system, sans-serif; color: #bbb; margin: 0 0 10px; }
.win, .mac { width: 260px; padding: 4px; border-radius: 8px; }
.win { background: #2c2c2c; border: 1px solid #444; } .mac { background: rgba(40, 40, 44, .96); border: 1px solid #555; }
.win div, .mac div { padding: 6px 12px; border-radius: 4px; } .hov { background: #0a64d8; color: #fff; } .dis { color: #888; }
hr { border: 0; border-top: 1px solid #555; margin: 4px 8px; }
.tip { margin-top: 10px; display: inline-block; padding: 4px 8px; background: #2c2c2c; border: 1px solid #444; font-size: 12px; }
.bar { height: 24px; background: #1c1c1e; display: flex; justify-content: flex-end; align-items: center; padding: 0 10px; margin-bottom: 4px; border-radius: 4px; }
</style></head><body>
<div><h2>Windows 11</h2><div class="win">${items(1)}</div><div class="tip">${TEXT.trayTooltip.replace("{line}", line)}</div></div>
<div><h2>macOS</h2><div class="bar"><img src="${icon}" width="16" height="16" alt=""></div><div class="mac">${items(1)}</div></div>
</body></html>`;
}

async function shootHtml(browser, name, html, size) {
  const page = await browser.newPage({ viewport: size, deviceScaleFactor: 2 });
  await page.setContent(html);
  const file = path.join(OUT, name + '.png');
  await page.screenshot({ path: file, fullPage: true });
  await page.close();
  return file;
}

(async () => {
  fs.mkdirSync(OUT, { recursive: true });
  const browser = await chromium.launch();
  for (const s of INSTALL) console.log(path.relative(REPO, await shootInstall(browser, s)));
  console.log(path.relative(REPO, await shootHtml(browser, 'tray-menu', trayHtml(), { width: 680, height: 260 })));
  await browser.close();
})();
