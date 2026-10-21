// node --test companion/ui/format.test.js
'use strict';
const test = require('node:test');
const assert = require('node:assert');
const FORMAT = require('./format.js');
const COPY = require('./copy.js');

test('a long path is cut at separators only, keeping the drive and the last two folders', () => {
  const p = 'C:\\Program Files (x86)\\World of Warcraft\\_classic_era_';
  assert.strictEqual(p.length, 54);
  // The UX review's case: budget 38 at the 560 px window.
  assert.strictEqual(FORMAT.pathTruncate(p, 38), 'C:\\\u2026\\World of Warcraft\\_classic_era_');
  // 56 at the 800 px window holds the whole path.
  assert.strictEqual(FORMAT.pathTruncate(p, 56), p);
  assert.strictEqual(FORMAT.pathTruncate('short', 38), 'short');
  // Too small for two folders: the drive and the last one.
  assert.strictEqual(FORMAT.pathTruncate(p, 30), 'C:\\\u2026\\_classic_era_');
  // One name alone too long: the ellipsis ends that name, never a stray bracket or half a word in the middle.
  assert.strictEqual(FORMAT.pathTruncate(p, 12), 'C:\\\u2026\\_class\u2026');
  // Never a cut inside a name while a separator cut fits.
  for (let max = 20; max < 54; max++) {
    const out = FORMAT.pathTruncate(p, max);
    assert.ok(out.length <= max, out);
    assert.ok(out.endsWith('_classic_era_'), out);
    assert.ok(!out.includes('Files'), out);
    assert.ok(!out.includes('('), out);
  }
  // Another flavour folder and a POSIX path.
  assert.strictEqual(FORMAT.pathTruncate('D:\\Games\\Blizzard\\World of Warcraft\\_anniversary_', 38), 'D:\\\u2026\\World of Warcraft\\_anniversary_');
  assert.strictEqual(FORMAT.pathTruncate('/home/jhunt/Games/World of Warcraft/_classic_era_', 38), '/\u2026/World of Warcraft/_classic_era_');
});

test('a time is today, yesterday or a short date', () => {
  const now = new Date(2026, 9, 15, 21, 0);
  assert.strictEqual(FORMAT.when(new Date(2026, 9, 15, 20, 14).toISOString(), now, COPY.reports), 'Today 20:14');
  assert.strictEqual(FORMAT.when(new Date(2026, 9, 14, 20, 14).toISOString(), now, COPY.reports), 'Yesterday 20:14');
  assert.strictEqual(FORMAT.when(new Date(2026, 9, 7, 9, 5).toISOString(), now, COPY.reports), 'Wed 7 Oct, 09:05');
  assert.strictEqual(FORMAT.when('not a date', now, COPY.reports), '');
});

test('Automatic names who the log says is logging once it knows', () => {
  assert.strictEqual(FORMAT.automaticLabel('', COPY.settings), 'Automatic (read from the log)');
  assert.strictEqual(FORMAT.automaticLabel(undefined, COPY.settings), 'Automatic (read from the log)');
  assert.strictEqual(FORMAT.automaticLabel('Bow Jackzon', COPY.settings), 'Automatic (currently Bow Jackzon)');
  assert.strictEqual(COPY.settings.characterHelp,
    'Automatic reads who is logging from the combat log itself. Pick a character only if that is wrong.');
});

test('a character is named with its ruleset in words', () => {
  const rs = COPY.settings.rulesets;
  assert.strictEqual(FORMAT.characterLabel({ name: 'Thrall', ruleset: 'hardcore', region: 'us' }, rs), 'Thrall \u00b7 Hardcore \u00b7 US');
  assert.strictEqual(FORMAT.characterLabel({ name: 'Obnoxious', ruleset: 'weird', region: 'eu' }, rs), 'Obnoxious \u00b7 weird \u00b7 EU');
  assert.strictEqual(FORMAT.characterValue({ region: 'us', ruleset: 'normal', name: 'A' }), 'us|normal|A');
});

test('the countdown never goes below zero', () => {
  const now = new Date('2026-10-15T20:00:00Z');
  assert.deepStrictEqual(FORMAT.countdown('2026-10-15T20:09:41Z', now), { m: 9, s: 41 });
  assert.deepStrictEqual(FORMAT.countdown('2026-10-15T19:00:00Z', now), { m: 0, s: 0 });
});

test('every sentence in copy.js is a string or a function that returns one', () => {
  const walk = (o, path) => {
    for (const [k, v] of Object.entries(o)) {
      if (typeof v === 'object') walk(v, path + '.' + k);
      else if (typeof v === 'function') assert.strictEqual(typeof v('x', 1), 'string', path + '.' + k);
      else assert.strictEqual(typeof v, 'string', path + '.' + k);
    }
  };
  walk(COPY, 'COPY');
});

test('last played is a clock today, a short date otherwise, and honest when unknown', () => {
  const now = new Date(2026, 9, 15, 21, 0);
  assert.strictEqual(FORMAT.lastPlayed(new Date(2026, 9, 15, 19, 42).toISOString(), now, COPY.status), 'Last played 19:42');
  assert.strictEqual(FORMAT.lastPlayed(new Date(2026, 9, 14, 9, 5).toISOString(), now, COPY.status), 'Last played 14 Oct 09:05');
  assert.strictEqual(FORMAT.lastPlayed(undefined, now, COPY.status), 'Not played yet');
  assert.strictEqual(FORMAT.lastPlayed('', now, COPY.status), 'Not played yet');
  assert.strictEqual(FORMAT.lastPlayed('garbage', now, COPY.status), 'Not played yet');
});

test('the Device page says whose sign-in the site stopped accepting', () => {
  assert.strictEqual(COPY.device.signedOutBar('Thrall#1234'),
    'Your sign-in as Thrall#1234 is no longer valid. Sign in again to keep uploading.');
  assert.strictEqual(COPY.device.signedOutBar(''),
    'Your sign-in is no longer valid. Sign in again to keep uploading.');
});

test('how long ago says minutes, then the clock, then days', () => {
  const now = new Date(2026, 9, 19, 15, 0, 30);
  const at = (d, h, m, s) => new Date(2026, 9, d, h, m, s || 0).toISOString();
  const w = COPY.status.ago;
  assert.strictEqual(FORMAT.ago(at(19, 14, 59, 50), now, w), 'less than a minute ago');
  assert.strictEqual(FORMAT.ago(at(19, 14, 58, 40), now, w), '1 minute ago');
  assert.strictEqual(FORMAT.ago(at(19, 14, 58), now, w), '2 minutes ago');
  assert.strictEqual(FORMAT.ago(at(19, 9, 5), now, w), 'today at 09:05');
  assert.strictEqual(FORMAT.ago(at(18, 23, 50), now, w), 'yesterday');
  assert.strictEqual(FORMAT.ago(at(1, 12, 0), now, w), '18 days ago');
  assert.strictEqual(FORMAT.ago('garbage', now, w), '');
});

test('the Last sync row counts characters and says how long ago', () => {
  assert.strictEqual(COPY.status.lastSyncSent(10, '2 minutes ago'), '10 characters, 2 minutes ago');
  assert.strictEqual(COPY.status.lastSyncSent(1, 'yesterday'), '1 character, yesterday');
});

test('a report says who it is credited to, or that it is not yet', () => {
  const w = COPY.status.credit;
  assert.strictEqual(FORMAT.creditText({ state: 'credited', name: 'Bow Jackzon' }, w), 'credited to Bow Jackzon');
  assert.strictEqual(FORMAT.creditText({ state: 'credited', name: 'Bow-Jackzon' }, w), 'credited to Bow Jackzon');
  assert.strictEqual(FORMAT.creditText({ state: 'unmatched', name: 'Bow Jackzon' }, w), 'not credited yet');
  // A log that has named nobody yet is not a fault; once the report is done it is not credited.
  assert.strictEqual(FORMAT.creditText({ state: 'none' }, w), 'name not read yet');
  assert.strictEqual(FORMAT.creditText({ state: 'none' }, w, true), 'not credited yet');
  assert.strictEqual(FORMAT.creditText(undefined, w, true), 'not credited yet');
});

// The three /api/status bodies the companion sends (written by a Go test from the real headline
// builder: internal/app TestTheStatusFixturesAreWhatTheCompanionSends).
const fs = require('node:fs');
const path = require('node:path');
const fixture = (name) => JSON.parse(fs.readFileSync(path.join(__dirname, 'testdata', name), 'utf8')).data;

test('the Current report row and the headline agree in each of the three credit states', () => {
  const credited = fixture('status-credited.json');
  const notCredited = fixture('status-not-credited.json');
  const noName = fixture('status-no-name.json');
  const report = (s) => ({ zone: '', credit: s.credit });
  const row = (s) => FORMAT.currentReport(report(s), COPY.status);

  assert.strictEqual(credited.headline.sub, 'Uploading your fights as they end · logging as Bow Jackzon');
  assert.deepStrictEqual(row(credited), { lead: 'In progress', separator: ' · ', credit: 'credited to Bow Jackzon',
    credited: true, noName: false, help: '' });

  // Not credited: the card says what to do once, the row keeps its amber phrase and its help line.
  assert.strictEqual(notCredited.headline.sub,
    'Uploading your fights as they end. Log out once in game with the addon loaded and it is credited to you.');
  assert.deepStrictEqual(row(notCredited), { lead: 'In progress', separator: ' · ', credit: 'not credited yet',
    credited: false, noName: false, help: 'Log out once in game with the addon loaded. The companion credits this report then.' });
  assert.ok(!notCredited.headline.sub.includes(COPY.status.credit.notYetHelp), 'the help sentence is never on screen three times');
  assert.ok(!/matched/.test(notCredited.headline.sub), 'matched is the project\u2019s word');

  // No name read yet is not the unmatched state: no not-credited title, no amber, and the row agrees.
  assert.strictEqual(noName.headline.title, 'Logging');
  assert.strictEqual(noName.headline.sub, 'Uploading your fights as they end. Your name is read from the first fight.');
  assert.deepStrictEqual(row(noName), { lead: 'In progress', separator: ' · ', credit: 'name not read yet',
    credited: false, noName: true, help: '' });

  assert.strictEqual(credited.headline.tone, 'on');
  assert.strictEqual(noName.headline.tone, 'on');
  assert.strictEqual(notCredited.headline.tone, 'warn', 'the one Logging state that is amber');
  for (const s of [credited, notCredited, noName]) {
    assert.ok(!/-/.test(s.headline.sub.replace(/ · /g, '')), 'no hyphenated name');
    assert.ok(!('reason' in s.credit), 'the window reads two states and no reason');
  }
  // A known zone leads the row in place of "In progress".
  assert.strictEqual(FORMAT.currentReport({ zone: 'Molten Core', credit: credited.credit }, COPY.status).lead, 'Molten Core');
});

test('only the watched install shouts that advanced logging is off', () => {
  const s = fixture('status-credited.json');
  const [watched, other] = s.installs;
  assert.strictEqual(watched.watching, true);
  assert.strictEqual(other.watching, false);

  const off = (i, extra) => Object.assign({}, i, { advanced_logging: false, advanced_logging_known: true }, extra);
  // The unwatched install: muted "off", no instruction, whatever the payload carries.
  assert.deepStrictEqual(FORMAT.advancedView(off(other, { advanced_logging_help: 'x' }), 'logging', COPY.status),
    { kind: 'muted', text: 'Advanced logging: off', help: '' });
  assert.deepStrictEqual(FORMAT.advancedView(other, 'logging', COPY.status),
    { kind: 'muted', text: 'Advanced logging: off', help: '' });
  // The watched install: the warning and the instruction, the instruction once only.
  const help = 'Turn it on in Options > Network > Advanced Combat Logging.';
  assert.deepStrictEqual(FORMAT.advancedView(off(watched, { advanced_logging_help: help }), 'logging', COPY.status),
    { kind: 'warn', text: 'Advanced logging: OFF', help });
  assert.strictEqual(FORMAT.advancedView(off(watched, { advanced_logging_help: help }), 'advanced_off', COPY.status).help, '');
  // On and unknown are as they were.
  assert.strictEqual(FORMAT.advancedView(watched, 'logging', COPY.status).kind, 'on');
  assert.strictEqual(FORMAT.advancedView({ watching: true }, 'logging', COPY.status).kind, 'unknown');
  // The fixture's own payload: the unwatched install carries no instruction field.
  assert.ok(!('advanced_logging_help' in other));
});

test('Check for builds now asks once at a time and once every half minute', async () => {
  let now = 1000000;
  let calls = 0;
  let release;
  const gate = new Promise((r) => { release = r; });
  const c = FORMAT.checker(async () => { calls += 1; await gate; return 'done'; }, () => now);

  assert.strictEqual(c.canPress(), true);
  const first = c.press();
  assert.strictEqual(c.running, true);
  assert.strictEqual(await c.press(), 'ignored', 'a press while one is running does nothing');
  release();
  assert.strictEqual(await first, 'done');
  assert.strictEqual(c.running, false);
  assert.strictEqual(calls, 1);

  now += 29999;
  assert.strictEqual(c.canPress(), false);
  assert.strictEqual(await c.press(), 'ignored');
  now += 1;
  assert.strictEqual(c.canPress(), true);
  assert.strictEqual(await c.press(), 'done');
  assert.strictEqual(calls, 2);

  // A request that throws still frees the button, and still counts as the press.
  const bad = FORMAT.checker(async () => { throw new Error('x'); }, () => now);
  await assert.rejects(bad.press());
  assert.strictEqual(bad.running, false);
  assert.strictEqual(bad.canPress(), false);
});

test('the Addon row says when it checked and what it found', () => {
  const now = new Date('2026-10-20T17:45:00Z');
  const at = (secs) => new Date(now.getTime() - secs * 1000).toISOString();
  const line = (checked) => FORMAT.checkedLine(checked, now, COPY.status, COPY.status.ago);
  assert.strictEqual(line({ at: at(5), new_builds: 1, for: ['Bow Jackzon'] }), 'Checked just now · 1 build for Bow Jackzon');
  assert.strictEqual(line({ at: at(5), new_builds: 0 }), 'Checked just now · nothing new');
  assert.strictEqual(line({ at: at(5), new_builds: 2, for: ['Bow Jackzon', 'Booty'] }), 'Checked just now · 2 builds for Bow Jackzon, Booty');
  assert.strictEqual(line({ at: at(5), new_builds: 3 }), 'Checked just now · 3 builds');
  assert.strictEqual(line({ at: at(125), new_builds: 0 }), 'Checked 2 minutes ago · nothing new');
  assert.strictEqual(COPY.status.checkBuilds, 'Check for builds now');
  assert.strictEqual(COPY.status.checking, 'Checking…');
});

test('the stopped-responding note says where the tray icon is, and that uploads may have stopped', () => {
  assert.strictEqual(COPY.localDown,
    'The companion stopped responding, so your fights may not be uploading. Right-click its icon in the system tray (bottom right, under the ^ arrow), choose Quit, then open it again.');
});

test('the expiry bar is the share of the ten minutes left, and the last minute is amber', () => {
  const now = new Date('2026-10-15T20:00:00Z');
  assert.strictEqual(FORMAT.expiryShare('2026-10-15T20:10:00Z', now), 1);
  assert.strictEqual(FORMAT.expiryShare('2026-10-15T20:08:00Z', now), 0.8);
  assert.strictEqual(FORMAT.expiryShare('2026-10-15T20:00:30Z', now), 0.05);
  assert.strictEqual(FORMAT.expiryShare('2026-10-15T19:59:00Z', now), 0);
  assert.strictEqual(FORMAT.expiryShare('2026-10-15T20:30:00Z', now), 1);
  assert.strictEqual(FORMAT.expiryShare('', now), 0);
  assert.strictEqual(FORMAT.lastMinute('2026-10-15T20:01:00Z', now), true);
  assert.strictEqual(FORMAT.lastMinute('2026-10-15T20:01:01Z', now), false);
  assert.strictEqual(FORMAT.lastMinute('2026-10-15T19:00:00Z', now), true);
});

test('the Settings path budget is 56 characters at the 800 px window and 38 at 560', () => {
  const p = 'C:\\Program Files (x86)\\World of Warcraft\\_classic_era_';
  assert.strictEqual(FORMAT.pathBudget(800), 56);
  assert.strictEqual(FORMAT.pathBudget(560), 38);
  for (const w of [800, 560]) {
    const out = FORMAT.pathTruncate(p, FORMAT.pathBudget(w));
    assert.ok(out.length <= FORMAT.pathBudget(w));
    assert.ok(out.endsWith('_classic_era_'), 'the tail names the flavour: ' + out);
  }
  assert.strictEqual(FORMAT.pathTruncate(p, FORMAT.pathBudget(800)), p);
  const anniv = 'C:\\Program Files (x86)\\World of Warcraft\\_anniversary_';
  assert.strictEqual(FORMAT.pathTruncate(anniv, FORMAT.pathBudget(560)), 'C:\\\u2026\\World of Warcraft\\_anniversary_');
});

test('the Save strip says No changes (Save disabled) or Unsaved changes (Save enabled)', () => {
  assert.strictEqual(COPY.settings.noChanges, 'No changes');
  assert.strictEqual(COPY.settings.unsaved, 'Unsaved changes');
  assert.deepStrictEqual(FORMAT.saveStrip(false, false, COPY.settings), { text: 'No changes', enabled: false });
  assert.deepStrictEqual(FORMAT.saveStrip(true, false, COPY.settings), { text: 'Unsaved changes', enabled: true });
  assert.deepStrictEqual(FORMAT.saveStrip(true, true, COPY.settings), { text: 'Unsaved changes', enabled: false }, 'not pressable mid-save');
  const html = read('index.html');
  assert.match(html, /id="save-state"[^>]*>No changes</);
  assert.ok(read('app.js').includes('FORMAT.saveStrip('));
});

test('Settings has a page heading, and Add a folder sits in the Game folders label row', () => {
  const html = read('index.html');
  assert.match(html, /<section id="settings"[^>]*>\s*<h2 tabindex="-1">Settings<\/h2>/);
  const row = html.slice(html.indexOf('class="label-row"'));
  assert.ok(row.indexOf('id="folders-label"') < row.indexOf('id="add-folder"'));
  assert.ok(row.indexOf('id="add-folder"') < row.indexOf('id="folder-list"'), 'the button is above the list, never under the fade');
  assert.match(html, /<h2 tabindex="-1">This device<\/h2>/, 'the Device page shows its heading again');
});

test('the sign-in screen says the browser has opened, and the expiry once', () => {
  assert.strictEqual(COPY.device.opened, 'Your browser has opened. Sign in with Battle.net there, then come back.');
  assert.strictEqual(COPY.device.countdown(7, 31), 'Expires in 7:31');
  assert.strictEqual(COPY.device.lasts, undefined);
  assert.ok(COPY.device.blocked.startsWith('Could not open your browser'));
});

test('the update bar says how to install and what it costs, and the footer keeps the version', () => {
  assert.strictEqual(COPY.footer.updateReady,
    'Update ready. Quit the companion from the tray icon and start it again to install it. Uploads pause while it restarts, so pick a quiet moment.');
  assert.strictEqual(COPY.footer.version('0.2.13') + COPY.footer.joint + COPY.footer.justUpdated, 'Companion 0.2.13 \u00b7 just updated');
  assert.strictEqual(COPY.footer.updatedTo, undefined);
  const js = read('app.js');
  assert.ok(js.includes("class: 'updated', text: COPY.footer.justUpdated"), 'green on just updated only');
  // The footer says Companion at first paint, before the first snapshot.
  assert.match(read('index.html'), /id="foot-version">Companion</);
});

test('Status unavailable greys the last numbers and says they are last known', () => {
  assert.strictEqual(COPY.status.lastKnown, 'Last known');
  assert.strictEqual(COPY.status.unavailableTitle, 'Status unavailable');
});

const read = (f) => fs.readFileSync(path.join(__dirname, f), 'utf8');

test('a log that is not credited yet says so in the Status title, and the dot is amber', () => {
  const h = fixture('status-not-credited.json').headline;
  assert.strictEqual(h.title, 'Logging \u00b7 not credited to you yet');
  assert.strictEqual(h.tone, 'warn');
  assert.strictEqual(fixture('status-credited.json').headline.title, 'Logging');
  assert.strictEqual(fixture('status-no-name.json').headline.title, 'Logging');
});

test('Settings lists visibility, then the credited character, then the game folders, then Save', () => {
  const html = read('index.html');
  const at = (id) => html.indexOf('id="' + id + '"');
  assert.ok(at('visibility') > 0);
  assert.ok(at('visibility') < at('logging-character'));
  assert.ok(at('logging-character') < at('folder-list'));
  assert.ok(at('folder-list') < at('save'));
});

test('the Save strip is the last thing in Settings, with a fade above it', () => {
  const html = read('index.html');
  const settings = html.slice(html.indexOf('<section id="settings"'), html.indexOf('</section>', html.indexOf('<section id="settings"')));
  assert.ok(settings.trimEnd().endsWith('</div>'));
  assert.ok(settings.lastIndexOf('class="save-row"') > settings.lastIndexOf('class="panel"'), 'the strip sits after the panel, not inside it');
  const css = read('styles.css');
  assert.match(css, /\.save-row::before\s*\{[^}]*height: 24px/);
  assert.match(css, /--save-h: 64px/);
});

test('the nav opens with Status current, so no tab changes at first paint', () => {
  assert.match(read('index.html'), /<button data-page="status" aria-current="page">/);
});

test('the connected token line is muted, never a warning', () => {
  const js = read('app.js');
  const row = js.slice(js.indexOf("class: 'row device-saved'"), js.indexOf("class: 'row device-saved'") + 400);
  assert.ok(!row.includes('amber'));
});

test('Reports: the hint over the list, and a row the site has not got says so in muted words', () => {
  const w = COPY.status.credit;
  assert.strictEqual(COPY.reports.notCreditedHint, 'Not credited yet? Log out once in game with the addon loaded.');
  const row = (extra) => Object.assign({ key: 'k', url: 'https://x/r/1', done: true, credit: { state: 'unmatched', name: 'A' } }, extra);
  assert.deepStrictEqual(FORMAT.reportCredit(row(), w), { kind: 'notyet', text: 'not credited yet' });
  assert.strictEqual(FORMAT.reportCredit(row({ credit: { state: 'credited', name: 'Bow-Jackzon' } }), w).text, 'credited to Bow Jackzon');
  assert.strictEqual(FORMAT.reportCredit(row({ done: false, credit: { state: 'none' } }), w).kind, 'noname');
  assert.strictEqual(FORMAT.reportCredit(row({ done: true, credit: { state: 'none' } }), w).kind, 'notyet');
  // No URL: the site has no such report, so there is nothing to credit and no orange word.
  assert.deepStrictEqual(FORMAT.reportCredit(row({ url: '' }), w), { kind: 'absent', text: '' });
  assert.strictEqual(COPY.reports.notOnSite + w.separator + COPY.reports.notOnSiteWhen, 'not on the site yet \u00b7 uploads when the site is reachable');
  assert.strictEqual(FORMAT.anyNotCredited([row({ url: '' }), row({ credit: { state: 'credited', name: 'A' } })], w), false);
  assert.strictEqual(FORMAT.anyNotCredited([row({ url: '' }), row()], w), true);
  // The fixture list: a not-credited row on the site, one never created.
  const rows = fixture('reports-rows.json').reports;
  assert.strictEqual(FORMAT.anyNotCredited(rows, w), true);
  assert.deepStrictEqual(rows.filter((r) => !r.url).map((r) => FORMAT.reportCredit(r, w).kind), ['absent']);
});

test('/combatlog is a toggle: the words never tell the player to type it blind', () => {
  assert.strictEqual(COPY.reports.emptyBody,
    'In game, type /combatlog once to start logging. It switches off if you type it again, so leave it if it is already on.');
  const ready = fixture('status-ready.json').headline;
  assert.strictEqual(ready.sub, 'Waiting for a combat log. In game, type /combatlog once to start logging. It switches off if you type it again, so leave it if it is already on.');
  assert.strictEqual(COPY.status.noReport, 'None yet');
  assert.strictEqual(COPY.status.rows.lastSync, 'Characters sent');
});

test('every count is in the body face, the poll failure disables Check for builds, bars start with the text', () => {
  const js = read('app.js');
  assert.ok(!/class: 'v mono', text: String\(p\./.test(js), 'no mono count');
  assert.ok(js.includes('rows.checking || !rows.canCheck || rows.failed'));
  const css = read('styles.css');
  assert.match(css, /\.headline \.sub[^}]*text-wrap: pretty/);
  assert.match(css, /\.expiry \{ width: 190px/);
  assert.match(css, /\.bar\.update \{\s*align-items: flex-start/);
  assert.ok(!/\.bar\.update \{\s*justify-content: flex-start/.test(css));
  assert.match(css, /\.tag \{[^}]*var\(--line-strong\)/);
  assert.match(css, /footer a \{ color: var\(--muted\)/);
});

test('the Fights row never shows a bare 0 on an open report', () => {
  assert.strictEqual(COPY.status.noFightsYet, '0 so far');
  assert.strictEqual(COPY.status.noFightsHelp, 'The first fight uploads when it ends.');
});

test('the fight rows never count one fight as both uploaded and queued', () => {
  const w = COPY.status;
  const open = (o) => Object.assign({ report_key: 'k', report_id: 'r_1', logging: true, fights: 0, queued: 0, queued_fights: 0 }, o);
  // No report open: no Fights row at all.
  assert.deepStrictEqual(FORMAT.fightRows({ fights: 0, queued: 0 }, w), { show: false });
  // Everything still waiting: the one row says so, with no Queued row beside it.
  const first = FORMAT.fightRows(open({ fights: 3, queued: 5, queued_fights: 3 }), w);
  assert.strictEqual(first.text, 'Uploading 3 fights…');
  assert.strictEqual(first.queued, 0);
  assert.strictEqual(FORMAT.fightRows(open({ fights: 1, queued: 2, queued_fights: 1 }), w).text, 'Uploading 1 fight…');
  // Some up, some waiting: uploaded is the site's, queued is the rest.
  const some = FORMAT.fightRows(open({ fights: 5, queued: 4, queued_fights: 2 }), w);
  assert.strictEqual(some.text, '3');
  assert.strictEqual(some.queued, 2);
  // Nothing waiting: only the uploaded count, even with raw chunks queued.
  const clear = FORMAT.fightRows(open({ fights: 4, queued: 2, queued_fights: 0 }), w);
  assert.strictEqual(clear.text, '4');
  assert.strictEqual(clear.queued, 0);
  // An open report with no fight yet.
  assert.strictEqual(FORMAT.fightRows(open(), w).text, w.noFightsYet);
  assert.strictEqual(FORMAT.fightRows(open({ report_id: '' }), w).text, w.noReport);
});

test('the fixtures: a fresh report reads Uploading, Ready has no fights row', () => {
  const first = fixture('status-uploading.json');
  assert.strictEqual(FORMAT.fightRows(first.pipeline, COPY.status).text, 'Uploading 3 fights…');
  const some = FORMAT.fightRows(fixture('status-uploading-some.json').pipeline, COPY.status);
  assert.deepStrictEqual([some.text, some.queued], ['3', 2]);
  assert.strictEqual(FORMAT.fightRows(fixture('status-ready.json').pipeline, COPY.status).show, false);
});

test('the Device page keeps its heading visible and only Status hides its heading', () => {
  const html = fs.readFileSync(path.join(__dirname, 'index.html'), 'utf8');
  const heads = [...html.matchAll(/<section id="(\w+)"[^>]*>\s*(?:<div class="h2line">\s*)?<h2 ([^>]*)>/g)];
  const hidden = heads.filter((m) => /class="sr"/.test(m[2])).map((m) => m[1]);
  assert.deepStrictEqual(hidden, ['status']);
  assert.strictEqual(heads.length, 4);
});

test('the Settings help sentences collapse to one line in a short window and the bar balances its text', () => {
  const css = fs.readFileSync(path.join(__dirname, 'styles.css'), 'utf8');
  const short = css.match(/@media \(max-height: 520px\) \{[\s\S]*?\n\}/)[0];
  assert.match(short, /#settings \.help \{[^}]*white-space: nowrap[^}]*text-overflow: ellipsis/);
  assert.match(css, /\.bar \{[^}]*text-wrap: pretty/);
  assert.match(css, /@media \(max-width: 520px\)[\s\S]*\.row \.v \.check-builds \{ margin: 6px 0 0 -11px; \}/);
});
