// design/mocks/companion/shoot_ba4.cjs: the companion surfaces shoot_ui.cjs does not cover, for lane BA4's review
// (design/reviews/2026-11-11-ba4): the setup steps (companion/ui/install.html in each mode, at the window's fixed
// 660 x 400), the tray menu as macOS and Windows draw it (its words read from the generated companion/ui/copy_gen.js,
// so the capture shows the keys the binary ships, and the tray dot from internal/icon/assets), and the disk image's
// window (packaging/macos/dmg-background@2x.png with the two icons at dmg-layout.sh's places, labels in Finder's ink).
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

// The disk image's window: the background at its 2x, the app and Applications icons at dmg-layout.sh's centres.
function dmgHtml() {
  const layout = fs.readFileSync(path.join(REPO, 'companion/packaging/macos/dmg-layout.sh'), 'utf8');
  const nums = (key) => layout.match(new RegExp(`${key}=\\((\\d+) (\\d+)\\)`)).slice(1).map(Number);
  const [app, drop] = [nums('DMG_APP'), nums('DMG_DROP')];
  const bg = dataUrl(path.join(REPO, 'companion/packaging/macos/dmg-background@2x.png'));
  const seal = dataUrl(path.join(REPO, 'design/assets/logo/lx-seal-v4.png'));
  const icon = (at, img, label) => `<div class="icon" style="left:${at[0] - 64}px;top:${at[1] - 64}px">${img}<span>${label}</span></div>`;
  const folder = '<div class="folder"></div>';
  return `<!doctype html><html><head><meta charset="utf-8"><style>
body { margin: 0; background: #1e1e1e; font: 13px -apple-system, sans-serif; }
.title { height: 32px; background: #e8e6e3; display: flex; align-items: center; justify-content: center; color: #333; font-weight: 600; position: relative; }
.title i { position: absolute; left: 12px; top: 10px; width: 12px; height: 12px; border-radius: 50%; background: #ff5f57; box-shadow: 20px 0 #febc2e, 40px 0 #28c840; }
.win { position: relative; width: 660px; height: 400px; background: url(${bg}) 0 0 / 660px 400px; }
.icon { position: absolute; width: 128px; text-align: center; }
.icon img { width: 128px; height: 128px; display: block; }
.icon span { display: block; margin-top: 4px; color: #1d1d1f; font-size: 13px; }
.folder { width: 128px; height: 104px; margin-top: 12px; border-radius: 10px; background: linear-gradient(#64b5f6, #1e88e5); margin-bottom: 12px; }
</style></head><body><div class="title"><i></i>Forever Sixty</div><div class="win">
${icon(app, `<img src="${seal}" alt="">`, 'Forever Sixty')}${icon(drop, folder, 'Applications')}
</div></body></html>`;
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
  console.log(path.relative(REPO, await shootHtml(browser, 'dmg-window', dmgHtml(), { width: 660, height: 432 })));
  await browser.close();
})();
