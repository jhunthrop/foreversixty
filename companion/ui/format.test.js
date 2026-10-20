// node --test companion/ui/format.test.js
'use strict';
const test = require('node:test');
const assert = require('node:assert');
const FORMAT = require('./format.js');
const COPY = require('./copy.js');

test('a long path keeps its drive and its flavour folder', () => {
  const p = 'C:\\Program Files (x86)\\World of Warcraft\\_classic_era_';
  const out = FORMAT.midTruncate(p, 32);
  assert.strictEqual(out.length, 32);
  assert.ok(out.startsWith('C:\\Program'));
  assert.ok(out.endsWith('_classic_era_'));
  assert.ok(out.includes('\u2026'));
  assert.strictEqual(FORMAT.midTruncate('short', 32), 'short');
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
  assert.strictEqual(FORMAT.creditText({ state: 'none' }, w), 'not credited yet');
  assert.strictEqual(FORMAT.creditText(undefined, w), 'not credited yet');
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
    credited: true, help: '' });

  assert.strictEqual(notCredited.headline.sub, 'Uploading your fights as they end · character not matched yet');
  assert.deepStrictEqual(row(notCredited), { lead: 'In progress', separator: ' · ', credit: 'not credited yet',
    credited: false, help: 'Log out once in game with the addon loaded. The companion credits this report then.' });

  assert.strictEqual(noName.headline.sub, 'Uploading your fights as they end');
  assert.strictEqual(row(noName).credit, 'not credited yet');
  assert.strictEqual(row(noName).credited, false);

  for (const s of [credited, notCredited, noName]) {
    assert.strictEqual(s.headline.tone, 'on', 'the dot stays green');
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

test('the stopped-responding note says where the tray icon is', () => {
  assert.strictEqual(COPY.localDown,
    'The companion stopped responding. Right-click its icon in the system tray (bottom right, under the ^ arrow), choose Quit, then open it again.');
});
