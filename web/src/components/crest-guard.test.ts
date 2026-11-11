// web/src/components/crest-guard.test.ts
// One crest, one sign-up save, states as Tag (component audit 2026-11-10 clusters 2, 5 and 17; lane
// D). It scans every .astro, .svelte, .ts and .css file under src/ (tests, fixtures and generated
// data aside) and fails on:
//   - a crest path built outside ui/Crest and lib/class-crest.ts (draw ui/Crest, Crest.astro or
//     paintCrest);
//   - a round crest rule outside styles/crest.css (the `.crest` class owns the circle and the ring);
//   - `putSignup(` outside the shared save helper (lib/guild/signup-save.ts);
//   - the provenance `.pill` used outside feed rows (a state is a ui/Tag).
// Each rule's self-check below proves it catches the pattern it was written for.
import { readdirSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';

const SRC = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '..');

interface Rule {
  /** What the offence is and what to use instead. */
  readonly name: string;
  readonly pattern: RegExp;
  /** Files (relative to src/) the pattern is allowed in, each for the reason beside it. */
  readonly allowed: (file: string) => boolean;
  /** Which files the rule reads. */
  readonly scans: (file: string) => boolean;
}

const ANY = (): boolean => true;
const NOT_CSS = (file: string): boolean => !file.endsWith('.css');
const STYLED = (file: string): boolean => /\.(astro|svelte|css)$/.test(file);

/** The one crest module and its stylesheet. */
const CREST_HOME = new Set(['lib/class-crest.ts', 'styles/crest.css']);
/** The sign-up save: its API call's definition and the one helper every control saves through. */
const SIGNUP_HOME = new Set(['lib/guild/raids-api.ts', 'lib/guild/signup-save.ts']);
/** The provenance pill: its definition and the feed rows' source pill. */
const PILL_HOME = new Set(['styles/global.css', 'components/SourcePill.astro']);

const RULES: readonly Rule[] = [
  {
    name: 'a crest path built by hand: draw ui/Crest, ui/Crest.astro or paintCrest (lib/class-crest.ts)',
    pattern: /\/crests\/|\bclassCrestSrc\(/g,
    allowed: (file) => CREST_HOME.has(file),
    scans: ANY,
  },
  {
    name: 'a round crest rule outside styles/crest.css: the .crest class owns the circle and the ring',
    pattern:
      /[^{};]*crest[^{};]*\{[^}]*border-radius\s*:|<(?:img|Crest)\b[^>]*\bcrest[^>]*\brounded-full\b|<(?:img|Crest)\b[^>]*\brounded-full\b[^>]*\bcrest/gis,
    allowed: (file) =>
      CREST_HOME.has(file) ||
      // A guild's faction emblem or uploaded crest, not a class crest: its own faction ring.
      file === 'components/character/FactionCrest.svelte',
    scans: STYLED,
  },
  {
    name: 'a sign-up saved by hand: save through lib/guild/signup-save.ts',
    pattern: /\bputSignup\(/g,
    allowed: (file) => SIGNUP_HOME.has(file),
    scans: NOT_CSS,
  },
  {
    name: 'the provenance .pill outside feed rows: a state is a ui/Tag',
    pattern:
      /class(?::list)?=["'{`][^"'}`>]*(?<![\w-])pill(?![\w-])|^\s*\.pill(?![\w-])|['"`]pill(?:-(?:site|sample|blizzard|community|datamined))?\s+pill-/gm,
    allowed: (file) => PILL_HOME.has(file),
    scans: ANY,
  },
];

function relative(file: string): string {
  return path.relative(SRC, file).split(path.sep).join('/');
}

function scannedFiles(dir: string): string[] {
  return readdirSync(dir, { withFileTypes: true }).flatMap((entry) => {
    const full = path.join(dir, entry.name);
    if (entry.isDirectory())
      return entry.name === 'data' || entry.name === 'fixtures' || entry.name === 'test-support'
        ? []
        : scannedFiles(full);
    if (entry.name.includes('.test.')) return [];
    return /\.(astro|svelte|ts|css)$/.test(entry.name) ? [full] : [];
  });
}

function withoutComments(text: string): string {
  return text
    .replace(/\/\*[\s\S]*?\*\//g, '')
    .replace(/<!--[\s\S]*?-->/g, '')
    .replace(/^\s*\/\/.*$/gm, '');
}

/** The offences one file's text holds, by rule name. */
export function offencesIn(file: string, text: string): string[] {
  const code = withoutComments(text);
  return RULES.filter((rule) => rule.scans(file) && !rule.allowed(file))
    .filter((rule) => (code.match(rule.pattern)?.length ?? 0) > 0)
    .map((rule) => rule.name);
}

describe('the crest guard', () => {
  const files = scannedFiles(SRC).map((full) => ({ file: relative(full), full }));

  it('no file builds its own crest, rounds a crest, saves a sign-up by hand or draws a state as a pill', () => {
    const found = files.flatMap(({ file, full }) =>
      offencesIn(file, readFileSync(full, 'utf8')).map((name) => `${file}: ${name}`),
    );
    expect(found).toEqual([]);
  });

  it('catches every pattern it was written for', () => {
    const page = 'components/guild/Example.svelte';
    const lib = 'lib/guild/example.ts';
    const css = 'styles/example.css';
    expect(offencesIn(lib, 'img.src = `/icons/hd/crests/${slug}.webp`;')).toHaveLength(1);
    expect(offencesIn(page, '<img src={classCrestSrc(slug)} alt="" />')).toHaveLength(1);
    expect(offencesIn(css, "  --pointer-crest: url('/icons/hd/crests/mage.webp');")).toHaveLength(1);
    expect(offencesIn(css, '.row-crest {\n  width: 28px;\n  border-radius: 999px;\n}')).toHaveLength(1);
    expect(offencesIn(page, '<style>\n  .x-crest { border-radius: 50%; }\n</style>')).toHaveLength(1);
    expect(offencesIn(page, '<img class="typeahead-crest rounded-full" src={src} alt="" />')).toHaveLength(1);
    expect(offencesIn(page, '<Crest cls={c} size={28} class="rounded-full" testid="crest" />')).toHaveLength(1);
    expect(offencesIn(page, 'await putSignup(guildId, raid.id, key, body);')).toHaveLength(1);
    expect(offencesIn(lib, 'this.deps.putSignup(id, raidId, key, body)')).toHaveLength(1);
    expect(offencesIn(page, '<span class="pill pill-sample">{copy.waiting}</span>')).toHaveLength(1);
    expect(offencesIn(page, '<span class="pill {chip.recipe} shrink-0">{chip.text}</span>')).toHaveLength(1);
    expect(offencesIn(css, '.pill {\n  padding: 0 6px;\n}')).toHaveLength(1);
    expect(offencesIn(lib, "return ok ? 'pill pill-site' : 'pill pill-sample';")).toHaveLength(1);
    // What the shared modules write, and the look-alikes, are allowed.
    expect(offencesIn('lib/class-crest.ts', 'return `/icons/hd/crests/${slug}.webp`;')).toEqual([]);
    expect(offencesIn('styles/crest.css', '.crest {\n  border-radius: 999px;\n}')).toEqual([]);
    expect(offencesIn('lib/guild/signup-save.ts', 'await putSignup(g, r, k, b);')).toEqual([]);
    expect(offencesIn(page, '<Crest cls={c} size={28} class="csel-crest" />')).toEqual([]);
    expect(offencesIn(page, '<span class="skeleton-block size-12 rounded-full"></span>')).toEqual([]);
    expect(offencesIn(css, '.csel-crest {\n  width: 32px;\n  height: 32px;\n}')).toEqual([]);
    expect(offencesIn(page, '<Tag tone="waiting">{copy.waiting}</Tag>')).toEqual([]);
    expect(offencesIn(page, 'const pill = pills.find((entry) => entry.id === id);')).toEqual([]);
    expect(offencesIn(page, '<span class="time-pill numeric">{t}</span>')).toEqual([]);
  });
});
