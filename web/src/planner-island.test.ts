// web/src/planner-island.test.ts
// planner-island.css is the only stylesheet the API's server-rendered /b/{id} page links,
// and that page draws the site's header and footer around the planner with the site's own
// utility classes. Nothing in the island's module graph imports Header.astro or Footer.astro,
// so the only thing keeping their classes in the emitted file is Tailwind's source scan --
// which is a build-time heuristic, not a dependency. This test pins it: every class those two
// components render has to come out of the island build as a selector.
//
// The same "one stylesheet has to be enough" rule covers the faces below it. Both halves fail
// the same way -- a page on foreversixty.gg that does not look like foreversixty.gg -- and
// neither shows up in any other test, because the Astro pages get their CSS by a different
// route and would stay correct while this one rotted.
//
// It builds the island itself rather than reading dist/. A test that read the published file
// would pass silently on a tree that was never built, or against a stale one from a previous
// commit -- and `npm test` runs before `npm run build` in .github/workflows/web.yml. Building
// from vite.island.config.ts costs a fraction of a second and can be neither absent nor stale.
import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { mkdtemp, readFile, rm } from 'node:fs/promises';
import os from 'node:os';
import path from 'node:path';
import { afterAll, beforeAll, describe, expect, it } from 'vitest';
import { build } from 'vite';
import Footer from './components/Footer.astro';
import Header from './components/Header.astro';

const WEB_ROOT = path.resolve(import.meta.dirname, '..');

/**
 * Escapes a class name the way Tailwind does when it writes the selector: every character
 * outside `[A-Za-z0-9_-]` is backslashed, so `md:px-12` becomes `md\:px-12`, `px-[18px]`
 * becomes `px-\[18px\]` and `border-line/60` becomes `border-line\/60`.
 */
function escapeClassName(name: string): string {
  return name.replace(/[^A-Za-z0-9_-]/g, (character) => `\\${character}`);
}

/**
 * True when `css` carries a rule for `className`. The trailing look-ahead is what makes this
 * an assertion rather than a substring search: without it a missing `.flex` would be masked
 * by the `.flex-nowrap` two lines down.
 */
function hasSelector(css: string, className: string): boolean {
  const escaped = escapeClassName(className).replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
  return new RegExp(`\\.${escaped}(?![\\w-])`).test(css);
}

/** Every family `@font-face` declares, in the order a sorted list gives. */
function declaredFontFamilies(stylesheet: string): string[] {
  const families = new Set<string>();
  for (const [, block] of stylesheet.matchAll(/@font-face\s*\{([^}]*)\}/g)) {
    const declaration = /font-family:\s*(?:'([^']*)'|"([^"]*)"|([^;]+))/.exec(block);
    if (declaration) families.add((declaration[1] ?? declaration[2] ?? declaration[3]).trim());
  }
  return [...families].sort();
}

function classNamesIn(html: string): string[] {
  const names = new Set<string>();
  for (const [, value] of html.matchAll(/\sclass="([^"]*)"/g)) {
    for (const name of value.split(/\s+/)) if (name) names.add(name);
  }
  return [...names].sort();
}

/**
 * Header.astro's Reference disclosure (e082853) carries three class names that are not
 * Tailwind utilities, so this file's "Tailwind's source scan kept it" premise does not
 * apply to them: `nav-scroll-fade`'s rule lives in Header.astro's own scoped `<style>`
 * block, which this standalone build never sees (vite.island.config.ts runs the Tailwind
 * and Svelte plugins only -- no Astro plugin ever compiles Header.astro's `<style>` here,
 * on the main site that block reaches the page through Astro's own compiler, a separate
 * pipeline this file does not build); `reference-disclosure` and `reference-panel` are
 * bare semantic hooks with no CSS of their own anywhere in the codebase. Excluded here
 * rather than silently passing the "every class has a selector" check for them.
 */
const NON_TAILWIND_MARKER_CLASSES = new Set(['nav-scroll-fade', 'reference-disclosure', 'reference-panel']);

let css = '';
let outDir = '';

beforeAll(async () => {
  outDir = await mkdtemp(path.join(os.tmpdir(), 'planner-island-css-'));
  await build({
    configFile: path.join(WEB_ROOT, 'vite.island.config.ts'),
    root: WEB_ROOT,
    logLevel: 'silent',
    build: { outDir, emptyOutDir: true },
  });
  css = await readFile(path.join(outDir, 'planner-island.css'), 'utf8');
}, 120_000);

afterAll(async () => {
  if (outDir) await rm(outDir, { recursive: true, force: true });
});

describe('planner-island.css', () => {
  it('carries the design tokens and the base layer from global.css', () => {
    expect(css).toContain('--color-gold');
    expect(css).toContain('--font-display');
  });

  // The tokens above only name the families. Without the faces themselves the API page falls
  // back to Georgia and Arial, which is the same gap as a missing utility class one layer
  // down. An exact list rather than a `toContain` each, so a face that quietly stops being
  // emitted fails here and so does one nobody meant to ship.
  it('declares a face for every family tokens.css names', () => {
    expect(declaredFontFamilies(css)).toEqual(['Barlow', 'Barlow Fallback', 'Cinzel', 'JetBrains Mono']);
  });

  // A fallback face without the metric overrides is just Arial under another name, and the
  // CLS budget in lighthouserc.json depends on it matching Barlow's em-box.
  it('keeps the size-adjusted metrics on the Barlow fallback face', () => {
    expect(css).toMatch(/size-adjust:\s*95\.78%/);
    expect(css).toMatch(/ascent-override:\s*104\.41%/);
  });

  // Two regressions in one assertion. A relative `url()` would resolve against
  // https://foreversixty.gg/planner-island.css and 404, and a `data:` URL would mean the
  // build has gone back to inlining every face -- which is what `build.lib` does whatever
  // assetsInlineLimit says, and what turned this stylesheet into 307 kB of base64 once.
  it('references its font files at root-absolute URLs, not inlined and not relative', () => {
    const urls = [...css.matchAll(/url\(([^)]+)\)/g)].map(([, url]) => url.replace(/['"]/g, '').trim());
    expect(urls.length).toBeGreaterThanOrEqual(4);
    expect(urls.filter((url) => !url.startsWith('/'))).toEqual([]);
  });

  for (const [name, component] of [
    ['Header', Header],
    ['Footer', Footer],
  ] as const) {
    it(`covers every class ${name} renders`, async () => {
      const container = await AstroContainer.create();
      const names = classNamesIn(await container.renderToString(component));

      // Guards the extraction itself: were the attribute regex to stop matching, the filter
      // below would be vacuously empty and this would pass against an empty stylesheet.
      expect(names).toContain('px-[18px]');
      expect(
        names.filter(
          (className) => !NON_TAILWIND_MARKER_CLASSES.has(className) && !hasSelector(css, className),
        ),
      ).toEqual([]);
    });
  }
});
