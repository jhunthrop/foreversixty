// web/src/planner-island.test.ts
// planner-island.css is the only stylesheet the API's server-rendered /b/{id} page links,
// and that page draws the site's header and footer around the planner with the site's own
// utility classes. Nothing in the island's module graph imports Header.astro or Footer.astro,
// so the only thing keeping their classes in the emitted file is Tailwind's source scan --
// which is a build-time heuristic, not a dependency. This test pins it: every class those two
// components render has to come out of the island build as a selector.
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

function classNamesIn(html: string): string[] {
  const names = new Set<string>();
  for (const [, value] of html.matchAll(/\sclass="([^"]*)"/g)) {
    for (const name of value.split(/\s+/)) if (name) names.add(name);
  }
  return [...names].sort();
}

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
      expect(names.filter((className) => !hasSelector(css, className))).toEqual([]);
    });
  }
});
