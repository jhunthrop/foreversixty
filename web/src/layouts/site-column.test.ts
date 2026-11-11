// The owner's rule: one centred page frame for the whole site, one width on every page at a given
// viewport (2026-11-10: "consistent width for every experience"). `.page-column` (global.css) is
// the frame: Base.astro puts the header bar, the page body and the footer in it, a full-bleed band
// breaks out with `.full-bleed` and puts its content back with `.page-column`, and no page or
// component declares a page-level container, a frame width or a frame gutter of its own. A narrower
// measure inside the frame is named: the `.measure-*` classes, or ui/Page's `reading` and `form`.
import { readdirSync, readFileSync, statSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import { createPageContainer } from '../test-support/page-container';
import Base from './Base.astro';

const SRC = path.join(path.dirname(fileURLToPath(import.meta.url)), '..');
const SCANNED_DIRS = ['pages', 'layouts', 'components'];
const SOURCE_EXTENSIONS = /\.(astro|svelte|css)$/;
/** Layouts a page may render through; each one reaches Base.astro (the route components import it). */
const LAYOUT_IMPORT = /(layouts\/(Base|Content)|tiers\/TierPage|bis\/BisSpecRoute|guild\/GuildRoute)\.astro/;
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
  {
    pattern: /--page-(?:max|gutter)\b/,
    why: 'reads the frame width or gutter; only .page-column (global.css) may, from tokens.css',
  },
  { pattern: /data-column/, why: 'a per-page column width; every page has one frame' },
  {
    pattern: /px-\[18px\][^"']*md:px-0|md:px-0[^"']*px-\[18px\]/,
    why: 'a hand-written phone inset; use px-(--page-inset)',
  },
];

/**
 * Component-internal widths that are not a page container, file -> the exact lines allowed.
 * `100vw` inside a calc() or clamp() sizing a popover or a width query is not a breakout;
 * everything else must go through the shared column classes.
 */
const ALLOWED: Readonly<Record<string, readonly RegExp[]>> = {
  // A grid of tree panels centred as a fit-content group, not a page container.
  'components/planner/TreeGrid.svelte': [/mx-auto w-fit/],
  // The item tooltip sizes itself to the viewport, minus a margin.
  'components/ItemTooltip.svelte': [/calc\(100vw_-_32px\)/],
  'components/ui/TalentTip.svelte': [/calc\(100vw - 32px\)/], // the talent tooltip's phone cap, a popover
  // The header bar's phone name clamp reads the viewport, not a page container.
  'styles/nav.css': [/calc\(100vw - 460px\)/],
  // The frame's own width and gutter (web layout spec 2026-11-05 §3.2): the tokens, not a container.
  'styles/tokens.css': [/clamp\(1600px, 100vw - 160px, 1792px\)/, /--page-(?:max|gutter):/],
  // The frame itself, the breakout, and the clip that keeps it from adding a scrollbar.
  'styles/global.css': [
    /100vw/,
    /margin-inline: auto/,
    /margin-inline: calc/,
    /var\(--page-(?:max|gutter)\)/,
  ],
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

  it('Base.astro puts the header bar in the same frame, and gives no page a column of its own', async () => {
    const container = await createPageContainer();
    const html = await container.renderToString(Base, {
      props: { title: 'Column', description: 'x', path: '/column' },
      slots: { default: '<main data-probe>body</main>' },
    });
    expect(html).toMatch(
      /<div class="site-top[^"]*"[^>]*>\s*<div class="page-column">\s*<header class="site-header"/,
    );
    expect(html).not.toMatch(/data-column/);
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
    // A band component may use .page-column for its own inner content (inside a .full-bleed
    // band); what only the layout may do is wrap the page's slot.
    const wrappers = scannedFiles()
      .filter((file) => /^(pages|layouts)\//.test(relative(file)))
      .filter((file) => /<div class="page-column"[^>]*>\s*<slot\s*\/>/.test(readFileSync(file, 'utf8')));
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
