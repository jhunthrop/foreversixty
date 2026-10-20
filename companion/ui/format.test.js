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

test('a report says who it is credited to, or why it is not', () => {
  const w = COPY.status.credit;
  assert.strictEqual(FORMAT.creditText({ state: 'credited', name: 'Bow Jackzon' }, w), 'credited to Bow Jackzon');
  assert.strictEqual(FORMAT.creditText({ state: 'unmatched', name: 'Bow',
    reason: 'the log names Bow, but no addon record carries that realm — log out once with the addon loaded' }, w),
  'not credited: the log names Bow, but no addon record carries that realm — log out once with the addon loaded');
  assert.strictEqual(FORMAT.creditText({ state: 'none' }, w), 'not credited');
  assert.strictEqual(FORMAT.creditText(undefined, w), 'not credited');
});

test('the stopped-responding note says where the tray icon is', () => {
  assert.strictEqual(COPY.localDown,
    'The companion stopped responding. Right-click its icon in the system tray (bottom right, under the ^ arrow), choose Quit, then open it again.');
});
