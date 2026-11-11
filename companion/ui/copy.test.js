// node --test companion/ui/copy.test.js
// The window's copy guard (the companion's surface guard): every word the window shows comes from
// apps/shared/copy/companion.json through copy_gen.js (generated) and copy.js (lookups). A sentence
// written inline in a script fails here with its file and line; add it to companion.json and run
// `node web/scripts/gen-mobile-copy.mjs`. The Go side's guard is internal/copygen/guard_test.go.
'use strict';
const test = require('node:test');
const assert = require('node:assert');
const fs = require('node:fs');
const path = require('node:path');

// The scripts the window runs, less the generated ones (copy_gen.js, icons.js) and the tests.
const GUARDED = ['app.js', 'install.js', 'copy.js', 'format.js'];

// literals is every string literal in a script, with its line, comments skipped.
function literals(src) {
  const out = [];
  let line = 1;
  for (let i = 0; i < src.length; i++) {
    const c = src[i];
    if (c === '\n') line++;
    else if (c === '/' && src[i + 1] === '/') {
      while (i < src.length && src[i] !== '\n') i++;
      line++;
    } else if (c === '/' && src[i + 1] === '*') {
      const end = src.indexOf('*/', i + 2);
      line += (src.slice(i, end).match(/\n/g) || []).length;
      i = end + 1;
    } else if (c === "'" || c === '"' || c === '`') {
      let j = i + 1;
      let text = '';
      while (j < src.length && src[j] !== c) {
        if (src[j] === '\\') text += src[++j];
        else text += src[j];
        j++;
      }
      out.push({ line, text });
      line += (text.match(/\n/g) || []).length;
      i = j;
    }
  }
  return out;
}

const TWO_WORDS = /\p{L}{2,}[ ,’']+\p{L}{2,}/u;
const SENTENCE = /\p{L}{2,}[ ,]+\p{L}+[ ,]+\p{L}+.*[.?:…]$/u;
// markup is an element's attributes or a CSS selector list, which name things, never say them.
const MARKUP = /^[a-z0-9-]+( [a-z0-9-]+)*$|^[#.[]|\[[a-z-]+[=\]]/;

// isProse reports whether a literal reads as words for a player: two or more words, and a
// capital first letter, a {placeholder} or a sentence's end. Class lists and selectors are not.
function isProse(s) {
  const t = s.trim();
  if (!TWO_WORDS.test(t) || !/\p{Ll}/u.test(t) || MARKUP.test(t)) return false;
  const first = t.match(/\p{L}/u)[0];
  return first === first.toUpperCase() || t.includes('{') || SENTENCE.test(t);
}

const inlineCopy = (src) => literals(src).filter((l) => isProse(l.text)).map((l) => l.line + ': ' + l.text);

test('the window\'s words come from the shared copy, never a literal in a script', () => {
  for (const file of GUARDED) {
    const hits = inlineCopy(fs.readFileSync(path.join(__dirname, file), 'utf8'));
    assert.deepStrictEqual(hits, [], file + ': words for a player belong in apps/shared/copy/companion.json');
  }
});

test('the guard finds a sentence, a label and a placeholder line, and passes the rest', () => {
  const src = [
    "// A comment: 'Not this one.'",
    "const a = 'Fights upload as they end.';",
    "const b = { label: 'Open your reports' };",
    "const c = `{n} fights are waiting`;",
    "/* 'Nor this.' */ const d = 'ghost small outline';",
    "const e = ' · ' + 'just now' + 'https://foreversixty.gg';",
    "const f = 'main section';",
  ].join('\n');
  assert.deepStrictEqual(inlineCopy(src), ['2: Fights upload as they end.', '3: Open your reports', '4: {n} fights are waiting']);
});

test('copy.js reads every word from copy_gen.js: no leaf is missing', () => {
  const COPY = require('./copy.js');
  const walk = (o, at) => {
    for (const [k, v] of Object.entries(o)) {
      assert.notStrictEqual(v, undefined, at + '.' + k + ' names a key copy_gen.js does not have');
      if (typeof v === 'object') walk(v, at + '.' + k);
    }
  };
  walk(COPY, 'COPY');
});

test('the reports job has one name: the tray, the Reports page and the Device page say the site\'s words', () => {
  const COPY = require('./copy.js');
  const TEXT = require('./copy_gen.js');
  assert.strictEqual(TEXT.reportsOpen, 'Open your reports');
  assert.strictEqual(COPY.reports.all, TEXT.reportsOpen);
  assert.strictEqual(COPY.device.openReports, TEXT.reportsOpen);
  assert.strictEqual(COPY.pages.reports, 'Your reports');
});
