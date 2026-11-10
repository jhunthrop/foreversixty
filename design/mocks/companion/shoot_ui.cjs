// design/mocks/companion/shoot_ui.cjs: captures of the BUILT companion window (companion/ui) in every state of
// design/reviews/2026-11-09-companion/inventory.md, at the window's starting size (800 x 620) and, for the main
// pages, at a larger 1100 x 800 and at the smallest the window goes, 560 x 480. The page is served from companion/ui exactly as the binary
// embeds it; api/status answers with a snapshot from companion/ui/testdata (the same snapshots the Go tests pin),
// adjusted per state where the fixtures have none. The clock is frozen 4 minutes after the fixtures' report start.
//   node design/mocks/companion/shoot_ui.cjs <node_modules with playwright> <out dir> [state-name-filter]
const path = require('node:path');
const fs = require('node:fs');
const { chromium } = require(path.join(process.argv[2], 'playwright'));

const REPO = path.resolve(__dirname, '../../..');
const UI = path.join(REPO, 'companion/ui');
const OUT = path.resolve(process.argv[3]);
const FILTER = process.argv[4] ? new RegExp(process.argv[4]) : null;
const ORIGIN = 'http://127.0.0.1:47123/9f2c1a7e/';
const NOW = new Date('2026-10-20T17:44:00Z');
const TYPES = { '.html': 'text/html', '.js': 'text/javascript', '.css': 'text/css', '.png': 'image/png', '.webp': 'image/webp', '.woff2': 'font/woff2' };

const fixture = (name) => JSON.parse(fs.readFileSync(path.join(UI, 'testdata', name + '.json'), 'utf8')).data;
const ok = (data) => ({ ok: true, data, error: null });
const fail = (code, message) => ({ ok: false, data: null, error: { code, message } });
const merge = (base, patch) => Object.assign(fixture(base), patch);

const ADDON_LINES = {
  enabled: true,
  lines: [{ text: 'Addon 2.0.16 · up to date', tone: '' }, { text: 'Data 2026.11.09 · up to date', tone: '' }],
  offers: [],
};
// One raid night across every page: the Reports list's reports, the open one (Molten Core, 3 fights) the
// report Status shows.
const logging = (patch) => {
  const s = merge('status-credited', Object.assign({ addon_update: ADDON_LINES, reports: fixture('reports-rows').reports }, patch));
  s.pipeline = Object.assign(s.pipeline, { report_key: 'kMolten Core', report_id: 'r_Molten Core', fights: 3, last_append: '2026-10-20T17:43:48Z' }, (patch || {}).pipeline || {});
  s.addon_seen_at = '2026-10-20T17:36:00Z';
  return s;
};

// Each state: the snapshot, the page, what to do once it shows, and the window size.
const STATES = [
  // first run, sign-in and pairing
  { name: 'first-run-device', snap: () => fixture('status-not-connected'), page: 'device' },
  { name: 'first-run-status', snap: () => fixture('status-not-connected'), page: 'status' },
  { name: 'pair-paste-code', snap: () => fixture('status-not-connected'), page: 'device', act: async (p) => p.click('[data-fid="fallback-toggle"]') },
  { name: 'pair-paste-code-bad', snap: () => fixture('status-not-connected'), page: 'device', post: { 'api/pair': fail('bad_code', 'That code is not valid. A code works once and expires after 10 minutes. Get a new one on the site.') },
    act: async (p) => { await p.click('[data-fid="fallback-toggle"]'); await p.fill('#pair-code', 'QQQ-QQQ'); await p.click('#pair'); } },
  { name: 'sign-in-waiting', snap: () => fixture('device-waiting'), page: 'device', small: true },
  { name: 'sign-in-last-minute', snap: () => fixture('device-waiting-last-minute'), page: 'device' },
  { name: 'sign-in-trouble', snap: () => fixture('device-waiting-trouble'), page: 'device' },
  { name: 'sign-in-browser-blocked', snap: () => fixture('device-blocked'), page: 'device' },
  { name: 'sign-in-failed', snap: () => fixture('device-failed'), page: 'device' },
  { name: 'sign-in-connecting-status', snap: () => fixture('status-connecting'), page: 'status' },
  { name: 'signed-out-by-site', snap: () => fixture('status-signed-out'), page: 'status' },
  { name: 'signed-out-device', snap: () => fixture('status-signed-out'), page: 'device' },
  { name: 'connected-flash', snap: () => fixture('status-not-connected'), page: 'device', then: () => fixture('status-credited'), wait: 2600 },
  { name: 'device-connected', snap: () => logging(), page: 'device' },
  { name: 'device-forget-confirm', snap: () => logging(), page: 'device', act: async (p) => p.click('[data-fid="forget"]') },
  { name: 'device-forgotten-local', snap: () => logging(), page: 'device', post: { 'api/unpair': ok(Object.assign(fixture('status-not-connected'), { revoked: false })) },
    act: async (p) => { await p.click('[data-fid="forget"]'); await p.click('#forget-yes'); }, then: () => fixture('status-not-connected') },
  // choosing the WoW folder
  { name: 'no-folder-status', snap: () => fixture('status-no-folder'), page: 'status' },
  { name: 'no-folder-settings', snap: () => fixture('status-no-folder'), page: 'settings' },
  { name: 'folder-added', snap: () => fixture('status-no-folder'), page: 'settings', post: { 'api/folders/pick': ok(fixture('status-ready')) },
    act: async (p) => { await p.click('#add-folder'); await p.evaluate(() => { const m = document.getElementById('main'); m.scrollTop = m.scrollHeight; }); }, after: () => fixture('status-ready') },
  { name: 'folder-not-wow', snap: () => fixture('status-ready'), page: 'settings', post: { 'api/folders/pick': fail('not_wow', 'That folder is not a World of Warcraft folder. Forever Sixty runs on the Classic beta; pick its _classic_beta_ folder.') }, act: async (p) => p.click('#add-folder') },
  // watching
  { name: 'ready', snap: () => merge('status-ready', { addon_update: ADDON_LINES }), page: 'status' },
  { name: 'advanced-logging-off', snap: () => fixture('status-advanced-off'), page: 'status' },
  { name: 'logging-credited', snap: () => logging(), page: 'status', big: true, small: true },
  { name: 'logging-no-name', snap: () => fixture('status-no-name'), page: 'status' },
  { name: 'logging-not-credited', snap: () => fixture('status-not-credited'), page: 'status' },
  // uploading a fight, done and failed
  { name: 'uploading', snap: () => fixture('status-uploading-some'), page: 'status' },
  { name: 'uploading-first', snap: () => fixture('status-uploading'), page: 'status' },
  { name: 'uploading-held', snap: () => fixture('status-upload-held'), page: 'status' },
  { name: 'reports', snap: () => fixture('reports-rows'), page: 'reports', big: true },
  { name: 'reports-copied', snap: () => fixture('reports-rows'), page: 'reports', act: async (p) => p.click('[data-fid="copy-kMolten Core"]') },
  { name: 'reports-empty', snap: () => fixture('reports-empty'), page: 'reports' },
  { name: 'reports-signed-out', snap: () => Object.assign(fixture('status-not-connected'), { reports: [] }), page: 'reports' },
  // the addon
  { name: 'addon-missing', snap: () => fixture('status-addon-missing'), page: 'status' },
  { name: 'addon-installing', snap: () => fixture('status-addon-missing'), page: 'status', post: { 'api/addon/install': 'hang' }, act: async (p) => p.click('[data-fid^="install-addon-"]') },
  { name: 'addon-install-game-running', snap: () => fixture('status-addon-missing'), page: 'status', post: { 'api/addon/install': fail('game_running', 'Close World of Warcraft first, then try again. The companion does not change addon files while the game is running.') }, act: async (p) => p.click('[data-fid^="install-addon-"]') },
  { name: 'addon-installed', snap: () => fixture('status-addon-installed'), page: 'status' },
  { name: 'addon-ready', snap: () => fixture('status-addon-ready'), page: 'status' },
  { name: 'addon-current', snap: () => fixture('status-addon-current'), page: 'status' },
  { name: 'addon-build-waiting', snap: () => fixture('status-addon-build-waiting'), page: 'status' },
  { name: 'addon-no-character', snap: () => fixture('status-addon-installed-no-character'), page: 'status' },
  { name: 'addon-stale', snap: () => fixture('status-addon-stale'), page: 'status' },
  { name: 'addon-never', snap: () => fixture('status-addon-never'), page: 'status' },
  { name: 'addon-checking-builds', snap: () => logging(), page: 'status', post: { 'api/builds/check': 'hang' }, act: async (p) => p.click('[data-fid="check-builds"]') },
  // settings
  { name: 'settings', snap: () => logging(), page: 'settings', big: true },
  { name: 'settings-unsaved', snap: () => logging(), page: 'settings', act: async (p) => p.selectOption('#visibility', 'guild'), small: true },
  { name: 'reports-560', snap: () => fixture('reports-rows'), page: 'reports', only: 'small' },
  { name: 'settings-saved', snap: () => logging(), page: 'settings', post: { 'api/settings': ok(Object.assign(logging(), { visibility: 'guild' })) },
    act: async (p) => { await p.selectOption('#visibility', 'guild'); await p.click('#save'); } },
  // errors and offline
  { name: 'offline-status', snap: () => logging(), page: 'status', then: 'down', wait: 2600 },
  { name: 'offline-first-launch', snap: 'down', page: null, wait: 1800 },
  { name: 'tray-notice', snap: () => logging(), page: 'status', act: async (p) => p.evaluate(() => window.trayNotice()) },
  // the updater
  { name: 'update-ready', snap: () => fixture('update-staged'), page: 'status' },
  { name: 'just-updated', snap: () => fixture('update-arrived'), page: 'status' },
];

async function shoot(browser, state, size, suffix) {
  const context = await browser.newContext({ viewport: size, deviceScaleFactor: 2, timezoneId: 'Europe/London', locale: 'en-GB' });
  const page = await context.newPage();
  await page.clock.install({ time: NOW });
  let current = state.snap;
  await page.route(ORIGIN + '**', async (route) => {
    const req = route.request();
    const rel = req.url().slice(ORIGIN.length).split('?')[0] || 'index.html';
    if (rel === 'api/status') {
      if (current === 'down') return route.abort();
      return route.fulfill({ json: ok(current()) });
    }
    if (rel.startsWith('api/')) {
      const answer = (state.post || {})[rel];
      if (answer === 'hang') return; // never answers: the busy state stays on screen
      return route.fulfill({ json: answer || ok(current === 'down' ? {} : current()) });
    }
    const file = path.join(UI, rel);
    if (!file.startsWith(UI) || !fs.existsSync(file)) return route.fulfill({ status: 404, body: '' });
    return route.fulfill({ body: fs.readFileSync(file), contentType: TYPES[path.extname(file)] });
  });
  await page.goto(ORIGIN);
  await page.evaluate(() => document.fonts.ready);
  await page.clock.runFor(300);
  if (state.page) {
    await page.click(`nav button[data-page="${state.page}"]`);
    await page.clock.runFor(300);
  }
  if (state.act) {
    if (state.after) current = state.after;
    await state.act(page);
    await page.clock.runFor(400);
  }
  if (state.then) {
    current = state.then;
    await page.clock.runFor(state.wait || 2200);
  } else if (state.wait) {
    await page.clock.runFor(state.wait);
  }
  // Let the 160 ms entrances settle; the live dot's ring is caught wherever it is.
  await page.clock.runFor(400);
  await page.waitForTimeout(800); // CSS animations run on the real clock
  const file = path.join(OUT, `${state.name}${suffix}.png`);
  await page.screenshot({ path: file });
  await context.close();
  return file;
}

(async () => {
  fs.mkdirSync(OUT, { recursive: true });
  const browser = await chromium.launch();
  for (const state of STATES) {
    if (FILTER && !FILTER.test(state.name)) continue;
    if (state.only !== 'small') console.log(path.relative(REPO, await shoot(browser, state, { width: 800, height: 620 }, '')));
    if (state.small || state.only === 'small') console.log(path.relative(REPO, await shoot(browser, state, { width: 560, height: 480 }, state.only ? '' : '-560')));
    if (state.big) console.log(path.relative(REPO, await shoot(browser, state, { width: 1100, height: 800 }, '-1100')));
  }
  await browser.close();
})();
