// web/scripts/check-glyph-coverage.mjs
//
// Verifies web/fonts/glyphs.txt (the explicit glyph inventory the five font subsets --
// Cinzel 700, Barlow 400/600/700, JetBrains Mono 500 -- are generated from) against every
// character the site can actually render:
//
//   - every HTML file in a fixture build of dist/ (the rendered text, after markdown and
//     templating -- this is what a browser actually paints, including any typographic
//     substitution a markdown renderer makes)
//   - every *.md file under src/content/**
//   - every `name` string in data/builds/<latest>/{items,itemnames,spells,talents,races,
//     classes,dungeons}.json
//   - every string value in data/curated/**/*.json
//
// A character outside the inventory is a real gap only if it is a printable, renderable
// character -- control characters and emoji are reported but never auto-added, since the
// site does not render either (emoji would need a whole separate color-glyph face; control
// characters are not glyphs at all).
//
// Usage: node scripts/check-glyph-coverage.mjs [--dist <path>]
// Run from web/. Exits non-zero (and prints the offending characters and where they came
// from) if a real gap is found so this can run as a CI gate; see package.json's
// "check:glyphs" script.
import { readFile, readdir } from 'node:fs/promises';
import path from 'node:path';
import process from 'node:process';
import { fileURLToPath } from 'node:url';

const WEB_ROOT = path.dirname(path.dirname(fileURLToPath(import.meta.url)));
const REPO_ROOT = path.dirname(WEB_ROOT);
const GLYPHS_PATH = path.join(WEB_ROOT, 'fonts', 'glyphs.txt');
const CONTENT_DIR = path.join(WEB_ROOT, 'src', 'content');
const CURATED_DIR = path.join(REPO_ROOT, 'data', 'curated');
const BUILD_DATA_FILES = [
  'items.json',
  'itemnames.json',
  'spells.json',
  'talents.json',
  'races.json',
  'classes.json',
  'dungeons.json',
];

/** Control characters (C0/C1, excluding whitespace we treat as structural) and any
 * character outside the Basic Multilingual Plane's printable text ranges that pyftsubset
 * could never usefully encode as a single BMP code point anyway (surrogate-pair emoji). */
function isIgnorable(codePoint) {
  if (codePoint === 0x09 || codePoint === 0x0a || codePoint === 0x0d) return true; // tab/LF/CR
  if (codePoint <= 0x1f || codePoint === 0x7f) return true; // C0 + DEL
  if (codePoint >= 0x80 && codePoint <= 0x9f) return true; // C1 controls
  if (codePoint >= 0x1f000) return true; // emoji / symbol supplements, surrogate range and up
  if (codePoint >= 0xfe00 && codePoint <= 0xfe0f) return true; // variation selectors
  if (codePoint === 0xfeff) return true; // BOM / zero-width no-break space (format char, not a glyph)
  return false;
}

async function findFiles(dir, predicate) {
  const out = [];
  let entries;
  try {
    entries = await readdir(dir, { withFileTypes: true });
  } catch {
    return out;
  }
  for (const entry of entries) {
    const full = path.join(dir, entry.name);
    if (entry.isDirectory()) {
      out.push(...(await findFiles(full, predicate)));
    } else if (predicate(entry.name)) {
      out.push(full);
    }
  }
  return out;
}

/** Strips tags/scripts/styles and decodes the handful of named entities the site's own
 * markup actually uses, leaving the text a browser would paint. Not a general HTML parser --
 * good enough for "what characters appear in rendered text". */
function textFromHtml(html) {
  let text = html
    .replace(/<script[\s\S]*?<\/script>/gi, ' ')
    .replace(/<style[\s\S]*?<\/style>/gi, ' ')
    .replace(/<!--[\s\S]*?-->/g, ' ')
    .replace(/<[^>]+>/g, ' ');
  text = text
    .replace(/&nbsp;/g, ' ')
    .replace(/&amp;/g, '&')
    .replace(/&lt;/g, '<')
    .replace(/&gt;/g, '>')
    .replace(/&quot;/g, '"')
    .replace(/&#39;/g, "'")
    .replace(/&mdash;/g, '—')
    .replace(/&ndash;/g, '–')
    .replace(/&#x?([0-9a-fA-F]+);/g, (_, code) =>
      String.fromCodePoint(parseInt(code, code.toLowerCase().startsWith('x') ? 16 : 10)),
    );
  return text;
}

function collectStrings(value, out) {
  if (typeof value === 'string') {
    out.push(value);
  } else if (Array.isArray(value)) {
    for (const item of value) collectStrings(item, out);
  } else if (value && typeof value === 'object') {
    for (const item of Object.values(value)) collectStrings(item, out);
  }
}

async function namesFromBuildFile(filePath) {
  let raw;
  try {
    raw = JSON.parse(await readFile(filePath, 'utf8'));
  } catch {
    return [];
  }
  // The task scopes this to "every `name` string", but the build files are not uniformly
  // shaped: items.json/spells.json/talents.json/races.json/classes.json are arrays of
  // objects with a `name` field, while itemnames.json is `{ build, names: { id: name } }`
  // (a map, not an array of objects). A full recursive string walk is a strict superset of
  // "every name string" for every one of these shapes, so it is used here instead of
  // special-casing each file -- it never misses a name and costs nothing extra (URLs and
  // ids it also picks up are already plain ASCII).
  const names = [];
  collectStrings(raw, names);
  return names;
}

async function stringsFromCuratedTree() {
  const files = await findFiles(CURATED_DIR, (name) => name.endsWith('.json'));
  const out = [];
  for (const file of files) {
    let raw;
    try {
      raw = JSON.parse(await readFile(file, 'utf8'));
    } catch {
      continue;
    }
    collectStrings(raw, out);
  }
  return out;
}

function latestBuildDir() {
  // data/builds/<version>/ -- several versions may exist; the task names
  // 1.60.1.70009 explicitly, so prefer it when present, else fall back to the
  // lexicographically last directory (version strings sort correctly here).
  return path.join(REPO_ROOT, 'data', 'builds', '1.60.1.70009');
}

async function main() {
  const distArgIndex = process.argv.indexOf('--dist');
  const distDir = distArgIndex !== -1 ? process.argv[distArgIndex + 1] : path.join(WEB_ROOT, 'dist');

  const inventoryRaw = await readFile(GLYPHS_PATH, 'utf8');
  // Comment lines start with '#'; stripping only the leading marker character isn't enough
  // since '#' is also a legitimate inventory member (it's printable ASCII) -- so the
  // inventory is built from non-comment lines only.
  const inventoryChars = new Set();
  for (const line of inventoryRaw.split('\n')) {
    if (line.startsWith('#')) continue;
    for (const ch of line) inventoryChars.add(ch);
  }

  const htmlFiles = await findFiles(distDir, (name) => name.endsWith('.html'));
  const mdFiles = await findFiles(CONTENT_DIR, (name) => name.endsWith('.md'));
  // Several pages (the planner, sim tools, reports, the home page's signed-in panel) render
  // their copy from Svelte components that hydrate client-side: that text exists only inside
  // the built JS bundles, never in the server-rendered HTML a static fixture build emits.
  // Scanning the bundles' raw text is a blunt proxy for "every string literal" (it also
  // picks up minified identifiers, which are plain ASCII and never trip the check), but it
  // is exactly how a real gap here was first found: U+2192 (rightwards arrow) in
  // src/lib/home-landing-copy.ts and src/lib/sim/copy.ts, present nowhere in the SSR HTML.
  const jsFiles = await findFiles(distDir, (name) => name.endsWith('.js'));

  const sources = [];

  for (const file of htmlFiles) {
    sources.push({
      label: `dist:${path.relative(distDir, file)}`,
      text: textFromHtml(await readFile(file, 'utf8')),
    });
  }
  for (const file of jsFiles) {
    sources.push({ label: `dist-js:${path.relative(distDir, file)}`, text: await readFile(file, 'utf8') });
  }
  for (const file of mdFiles) {
    sources.push({
      label: `content:${path.relative(CONTENT_DIR, file)}`,
      text: await readFile(file, 'utf8'),
    });
  }
  for (const name of BUILD_DATA_FILES) {
    const file = path.join(latestBuildDir(), name);
    const names = await namesFromBuildFile(file);
    if (names.length > 0) {
      sources.push({ label: `build-data:${name}`, text: names.join('\n') });
    }
  }
  const curatedStrings = await stringsFromCuratedTree();
  if (curatedStrings.length > 0) {
    sources.push({ label: 'curated:**/*.json', text: curatedStrings.join('\n') });
  }

  if (htmlFiles.length === 0) {
    console.warn(
      `check-glyph-coverage: no HTML files found under ${distDir} -- run "npm run build" ` +
        'first (FOREVER_DATA=fixture) to check the built site, not just source content.',
    );
  }

  const missing = new Map(); // char -> Set(labels)
  const ignored = new Map();

  for (const { label, text } of sources) {
    for (const ch of text) {
      const cp = ch.codePointAt(0);
      if (ch === '\n' || ch === '\r' || ch === '\t') continue;
      if (inventoryChars.has(ch)) continue;
      if (isIgnorable(cp)) {
        if (!ignored.has(ch)) ignored.set(ch, new Set());
        ignored.get(ch).add(label);
        continue;
      }
      if (!missing.has(ch)) missing.set(ch, new Set());
      missing.get(ch).add(label);
    }
  }

  console.log(
    `check-glyph-coverage: inventory has ${inventoryChars.size} characters; scanned ` +
      `${sources.length} sources (${htmlFiles.length} HTML, ${jsFiles.length} JS bundles, ` +
      `${mdFiles.length} markdown, ${BUILD_DATA_FILES.length} build-data files, curated tree).`,
  );

  if (ignored.size > 0) {
    console.log(`check-glyph-coverage: ${ignored.size} ignorable character(s) seen (not added):`);
    for (const [ch, labels] of ignored) {
      console.log(
        `  U+${ch.codePointAt(0).toString(16).toUpperCase().padStart(4, '0')} in ${[...labels].slice(0, 3).join(', ')}`,
      );
    }
  }

  if (missing.size > 0) {
    console.error(
      `check-glyph-coverage: ${missing.size} character(s) rendered but NOT in web/fonts/glyphs.txt:`,
    );
    for (const [ch, labels] of missing) {
      console.error(
        `  U+${ch.codePointAt(0).toString(16).toUpperCase().padStart(4, '0')} ${JSON.stringify(ch)} in ${[...labels].slice(0, 5).join(', ')}`,
      );
    }
    process.exitCode = 1;
    return;
  }

  console.log('check-glyph-coverage: every rendered character is covered by web/fonts/glyphs.txt.');
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
