// web/src/components/surface-guard.test.ts
// Only the ui/ primitives draw a surface (web layout spec 2026-11-05 §3A.3, owner decision 7). A
// card border, a raised ground, the panel radius, the card gradient or shadow, outside
// web/src/components/ui/, is a page inventing its own box: use Plate, Section, Row or Overlay.
// It scans every .astro and .svelte file under src/ and every styles/*.css file.
//
// It shipped as a ratchet (an allowlist of files still drawing, 168 on 2026-11-05) and every page
// lane deleted its files as it adopted the primitives; the restyle-rest round emptied the list on
// 2026-11-07, so the allowance is gone and any surface outside the primitives is a hard fail.
//
// The old type is the guard's second half (2026-11-07, after /bis kept its eyebrow, caps title, caps
// button and caps tag through the rollout: the old styling lived in an .astro page, a primitive the
// surface scan exempts, and global CSS classes it never looked at). It scans every .astro, .svelte
// and .css file under src/ (pages, layouts, the primitives included) and the class strings in lib/
// for the retired recipes (spec 3.3: one Marcellus title, section heads in the body face in normal
// case, sentence-case buttons and tags, nothing under 12 px): the caps heading and card utilities,
// the 11 px `label` micro-caps, the PageHead eyebrow and the Card primitives, uppercase type under
// the 12 px Label size, and uppercase display type (a heading face, a heading size or a button's
// tracking). The 12 px caps Label role (group labels, column heads) is the one caps type allowed.
// ui/Card, CardHead and IconTile are deleted and any import of them fails, as does a class
// constant in lib/ (outside the primitives' own lib/ui/ recipes) that rebuilds the card look.
import { readdirSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';

const SRC = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '..');
/** The primitives themselves, the one place a surface may be drawn. */
const PRIMITIVES = 'components/ui/';

/** Surface classes: the card, the panel radius, the line borders, the raised and hover grounds. */
const SURFACE_CLASS =
  /(?<![\w-])(?:card-surface(?:-quiet)?|rounded-panel|rounded-plate|(?:border|divide)(?:-[trblxy])?-line(?:-soft|-warm(?:-strong)?|-card)?|bg-raised|bg-card-top|bg-tile|bg-hover|bg-selected|shadow-card|shadow-overlay|border-\[var\(--color-line[\w-]*\))(?![\w-])/g;
/** A CSS background or border declaration that reads a surface token. */
const SURFACE_DECLARATION =
  /(?:^|[\s;{])(?:background|border)[\w-]*\s*:[^;}]*var\(--(?:color-(?:raised|card-top|line[\w-]*)|surface-card)\)/g;

function scannedFiles(dir: string): string[] {
  return readdirSync(dir, { withFileTypes: true }).flatMap((entry) => {
    const full = path.join(dir, entry.name);
    if (entry.isDirectory())
      return entry.name === 'data' || entry.name === 'fixtures' ? [] : scannedFiles(full);
    if (entry.name.includes('.test.')) return [];
    const rel = relative(full);
    if (rel.startsWith(PRIMITIVES)) return [];
    if (/\.(astro|svelte)$/.test(entry.name)) return [full];
    return rel.startsWith('styles/') && entry.name.endsWith('.css') ? [full] : [];
  });
}

function relative(file: string): string {
  return path.relative(SRC, file).split(path.sep).join('/');
}

function withoutComments(text: string): string {
  return text
    .replace(/\/\*[\s\S]*?\*\//g, '')
    .replace(/<!--[\s\S]*?-->/g, '')
    .replace(/^\s*\/\/.*$/gm, '');
}

/** Surface draws per file, keyed relative to src/. */
export function surfaceUses(): Record<string, number> {
  const uses: Record<string, number> = {};
  for (const file of scannedFiles(SRC)) {
    const text = withoutComments(readFileSync(file, 'utf8'));
    const count = (text.match(SURFACE_CLASS)?.length ?? 0) + (text.match(SURFACE_DECLARATION)?.length ?? 0);
    if (count > 0) uses[relative(file)] = count;
  }
  return uses;
}

/** Retired recipes by name: the caps heading and card utilities and the 11 px micro-caps label. */
const RETIRED_UTILITY =
  /(?<![\w-])(?:heading-hero|heading-page|heading-section|section-title|heading-plan|gold-text|surface-band|card-surface(?:-quiet)?)(?![\w-])/g;
/** The retired primitives: PageHead (the eyebrow), the gold-edged Card family and the IconTile chip. */
const RETIRED_PRIMITIVE =
  /(?:ui|\.)\/(?:PageHead|Card|CardHead|IconTile)\.(?:astro|svelte)['"]|lib\/ui\/(?:page-head|card)['"]/g;
/** A class string in lib/ that rebuilds the card look: a card utility or token, or a raised
 *  ground with an edge. lib/ui/ holds the primitives' own recipes (PLATE_SURFACE, the field). */
const CARD_LOOK_TOKEN =
  /(?<![\w-])(?:card-surface(?:-quiet)?|shadow-card|bg-tile|rounded-panel|rounded-plate)(?![\w-])|--surface-card/;
const RAISED_GROUND = /(?<![\w-])(?:bg-raised|bg-card-top)(?![\w-])/;
const EDGE = /(?<![\w-])border(?:-[trblxy])?(?:-line[\w-]*)?(?![\w-])/;
/** An element's class attribute (static, template, list or quoted expression). */
const CLASS_ATTRIBUTE = /\bclass(?::list)?=(?:"([^"]*)"|\{`([^`]*)`\}|\{\[([^\]]*)\]\}|\{'([^']*)'\})/g;
/** A CSS rule body that sets uppercase. */
const UPPERCASE_RULE = /\{[^{}]*text-transform:\s*uppercase[^{}]*\}/g;
/** A quoted class string that sets uppercase. */
const UPPERCASE_CLASSES = /["'`]([^"'`\n]*\buppercase\b[^"'`\n]*)["'`]/g;
/** The Label size, the only size caps type may take. */
const LABEL_PX = 12;
/** What makes caps a display heading or a button: the heading face, a heading or hero size, the button tracking. */
const DISPLAY_TOKENS = /--font-heading|--text-h\d|--text-hero|--text-card-title|--tracking-button/;
function typeScannedFiles(dir: string): string[] {
  return readdirSync(dir, { withFileTypes: true }).flatMap((entry) => {
    const full = path.join(dir, entry.name);
    if (entry.isDirectory())
      return entry.name === 'data' || entry.name === 'fixtures' ? [] : typeScannedFiles(full);
    if (entry.name.includes('.test.')) return [];
    const rel = relative(full);
    if (/\.(astro|svelte|css)$/.test(entry.name)) return [full];
    return rel.startsWith('lib/') && entry.name.endsWith('.ts') ? [full] : [];
  });
}

/** Quoted class strings in a lib/ module (outside lib/ui/) that draw a card. */
export function cardLookOffences(text: string): string[] {
  return [...text.matchAll(/["'`]([^"'`\n]+)["'`]/g)]
    .map(([, classes]) => classes!)
    .filter((classes) => CARD_LOOK_TOKEN.test(classes) || (RAISED_GROUND.test(classes) && EDGE.test(classes)))
    .map((classes) => `card look "${classes.slice(0, 48)}"`);
}

function uppercaseOffences(text: string): string[] {
  const found: string[] = [];
  for (const [rule] of text.matchAll(UPPERCASE_RULE)) {
    const px = /font-size:\s*(\d+)px/.exec(rule)?.[1];
    if (px !== undefined && Number(px) !== LABEL_PX) found.push(`uppercase at ${px}px`);
    else if (DISPLAY_TOKENS.test(rule)) found.push('uppercase display type');
  }
  for (const [, classes] of text.matchAll(UPPERCASE_CLASSES)) {
    const px = /(?:^|\s)(?:\w+:)?text-\[(\d+)px\]/.exec(classes!)?.[1];
    if (px !== undefined && Number(px) !== LABEL_PX) found.push(`uppercase at ${px}px`);
    else if (/(?:^|\s)text-(?:base|lg|[2-9]?xl)(?:\s|$)/.test(classes!)) found.push('uppercase display type');
  }
  return found;
}

/** Old-type offences per file, keyed relative to src/. */
export function oldTypeUses(): Record<string, string[]> {
  const uses: Record<string, string[]> = {};
  for (const file of typeScannedFiles(SRC)) {
    const text = withoutComments(readFileSync(file, 'utf8'));
    const found = [
      ...[...text.matchAll(RETIRED_UTILITY)].map(([name]) => name),
      ...[...text.matchAll(RETIRED_PRIMITIVE)].map(([name]) => `import ${name}`),
      ...uppercaseOffences(text),
      ...(/^lib\/(?!ui\/)/.test(relative(file)) ? cardLookOffences(text) : []),
    ];
    if (/\.(astro|svelte)$/.test(file)) {
      for (const match of text.matchAll(CLASS_ATTRIBUTE)) {
        const value = match.slice(1).find((part) => part !== undefined) ?? '';
        if (value.split(/[\s'",{}?:]+/).includes('label')) found.push('class label');
      }
    }
    if (found.length > 0) uses[relative(file)] = found;
  }
  return uses;
}

describe('the surface guard', () => {
  const uses = surfaceUses();

  it('no file outside the primitives draws a surface', () => {
    const drawn = Object.entries(uses).map(([file, count]) => `${file}: ${count}`);
    expect(drawn, 'draw it with a ui/ primitive (Plate, Section, Row, Overlay) instead').toEqual([]);
  });

  it('no file sets the old type: caps headings, eyebrows, micro-caps, caps buttons', () => {
    const old = Object.entries(oldTypeUses()).map(([file, found]) => `${file}: ${found.join(', ')}`);
    expect(
      old,
      'use PageHeader (one Marcellus title), Section (body-face heads), Button and Tag (sentence case), the 12 px Label role',
    ).toEqual([]);
  });

  it('catches the old type it was written for', () => {
    expect(uppercaseOffences('.x { font-size: 11px; text-transform: uppercase; }')).toEqual([
      'uppercase at 11px',
    ]);
    expect(uppercaseOffences('.x { font-family: var(--font-heading); text-transform: uppercase; }')).toEqual([
      'uppercase display type',
    ]);
    expect(uppercaseOffences("'text-[11px] font-bold uppercase'")).toEqual(['uppercase at 11px']);
    expect(uppercaseOffences("'text-label tracking-label font-bold uppercase'")).toEqual([]);
    expect('<h1 class="heading-page">'.match(RETIRED_UTILITY)).toEqual(['heading-page']);
    expect("import PageHead from '../ui/PageHead.astro';".match(RETIRED_PRIMITIVE)).toHaveLength(1);
    expect("import IconTile from '../ui/IconTile.svelte';".match(RETIRED_PRIMITIVE)).toHaveLength(1);
    expect(cardLookOffences("const BOX = 'bg-raised border-line rounded-control border';")).toHaveLength(1);
    expect(cardLookOffences("const ROW = 'border-b border-line-soft';")).toEqual([]);
  });
});
