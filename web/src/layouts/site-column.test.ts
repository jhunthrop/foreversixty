// The owner's rule: one centred site column for the whole site. Base.astro wraps every page in
// `.page-column` once, a full-bleed band breaks out with `.full-bleed` and puts its content back
// with `.page-column`, and no page or component declares a page-level container of its own.
import { readdirSync, readFileSync, statSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import { createPageContainer } from '../test-support/page-container';
import Base from './Base.astro';

const SRC = path.join(path.dirname(fileURLToPath(import.meta.url)), '..');
const SCANNED_DIRS = ['pages', 'layouts', 'components'];
const SOURCE_EXTENSIONS = /\.(astro|svelte|css)$/;
/** Layouts a page may render through; each one reaches Base.astro. */
const LAYOUT_IMPORT = /(layouts\/(Base|Content)|tiers\/TierPage)\.astro/;
/** Pages that render no HTML of their own, so they have no column to be in. */
const PAGES_WITHOUT_LAYOUT = new Set(['pricing.astro']);

/** Patterns that declare a page-level container, with the reason each is banned. */
const FORBIDDEN: ReadonlyArray<{ pattern: RegExp; why: string }> = [
  { pattern: /\bmx-auto\b/, why: 'centres a container; use .page-column or a .measure-* class' },
  { pattern: /margin-inline:\s*auto|margin:\s*0 auto/, why: 'centres a container in CSS' },
  { pattern: /\b100vw\b|\bw-screen\b/, why: 'a viewport-wide breakout; use .full-bleed' },
  {
    pattern: /max-w-\[(?:[89]\d\d|\d{4,})px\]|^\s*max-width:\s*(?:[89]\d\d|\d{4,})px/,
    why: 'a container-sized max-width (800px and up)',
  },
];

/**
 * Component-internal widths that are not a page container, file -> the exact lines allowed.
 * `100vw` inside a calc() or clamp() sizing a popover or a width query is not a breakout;
 * everything else must go through the shared column classes.
 */
const ALLOWED: Readonly<Record<string, readonly RegExp[]>> = {
  // The tier page's two cards cap their own width at the reading measure, inside the column.
  'components/tiers/TierPage.astro': [/max-width:\s*820px/],
  // A grid of tree panels centred as a fit-content group, not a page container.
  'components/planner/TreeGrid.svelte': [/mx-auto w-fit/],
  // The item tooltip sizes itself to the viewport, minus a margin.
  'components/ItemTooltip.svelte': [/calc\(100vw_-_32px\)/],
  // The header bar's phone name clamp reads the viewport, not a page container.
  'styles/nav.css': [/calc\(100vw - 460px\)/],
  // The breakout itself, and the clip that keeps it from adding a scrollbar.
  'styles/global.css': [/100vw/, /margin-inline: auto/, /margin-inline: calc/],
};

function sourceFiles(dir: string): string[] {
  return readdirSync(dir).flatMap((name) => {
    const full = path.join(dir, name);
    if (statSync(full).isDirectory()) return sourceFiles(full);
    return SOURCE_EXTENSIONS.test(name) && !name.includes('.test.') ? [full] : [];
  });
}

function relative(file: string): string {
  return path.relative(SRC, file).split(path.sep).join('/');
}

function scannedFiles(): string[] {
  const dirs = [...SCANNED_DIRS.map((d) => path.join(SRC, d)), path.join(SRC, 'styles')];
  return dirs.flatMap(sourceFiles);
}

describe('the single site column', () => {
  it('Base.astro wraps the page slot in .page-column', async () => {
    const container = await createPageContainer();
    const html = await container.renderToString(Base, {
      props: { title: 'Column', description: 'x', path: '/column' },
      slots: { default: '<main data-probe>body</main>' },
    });
    expect(html).toMatch(/<div class="page-column"[^>]*>\s*<main data-probe>body<\/main>\s*<\/div>/);
  });

  it('every page renders through a layout that reaches Base.astro', () => {
    const pagesDir = path.join(SRC, 'pages');
    const pages = sourceFiles(pagesDir).filter((f) => f.endsWith('.astro'));
    expect(pages.length).toBeGreaterThan(20);
    const outside = pages.filter(
      (file) =>
        !PAGES_WITHOUT_LAYOUT.has(path.basename(file)) && !LAYOUT_IMPORT.test(readFileSync(file, 'utf8')),
    );
    expect(outside.map(relative)).toEqual([]);
  });

  it('only Base.astro opens the column around the page slot', () => {
    // A band component may use .page-column for its own inner content (ArtPanel does, inside
    // .full-bleed); what only the layout may do is wrap the page's slot.
    const wrappers = scannedFiles()
      .filter((file) => /^(pages|layouts)\//.test(relative(file)))
      .filter((file) => /<div class="page-column">\s*<slot\s*\/>/.test(readFileSync(file, 'utf8')));
    expect(wrappers.map(relative)).toEqual(['layouts/Base.astro']);
  });

  it('defines the column once, in global.css, from the --page-max token', () => {
    const css = readFileSync(path.join(SRC, 'styles/global.css'), 'utf8');
    expect(css.match(/^\.page-column \{/gm)).toHaveLength(1);
    expect(css).toMatch(/\.page-column \{[^}]*max-width: var\(--page-max\)/);
    expect(css).toMatch(/\.full-bleed \{/);
  });

  it('no page, layout or component declares a page-level container of its own', () => {
    const violations = scannedFiles().flatMap((file) => {
      const allowed = ALLOWED[relative(file)] ?? [];
      return readFileSync(file, 'utf8')
        .split('\n')
        .flatMap((line, index) =>
          FORBIDDEN.filter(({ pattern }) => pattern.test(line))
            .filter(() => !allowed.some((ok) => ok.test(line)))
            .map(({ why }) => `${relative(file)}:${index + 1} ${why}: ${line.trim()}`),
        );
    });
    expect(violations).toEqual([]);
  });
});
