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
