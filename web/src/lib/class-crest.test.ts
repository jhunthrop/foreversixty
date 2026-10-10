// web/src/lib/class-crest.test.ts
import { readdirSync, readFileSync } from 'node:fs';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import { classCrestSrc } from './class-crest';
import { factionMarkSrc } from './faction-mark';

// The character selector's first frame (nav spec 2026-10-09 section 8): Base.astro's head
// script sets `data-pointer-class` from a hand-copied class list (an `is:inline` script cannot
// import), and nav.css maps each class to its crest file. These tests pin both against the
// crest files that actually ship and the path this module builds, so a new or renamed crest
// never silently goes stale in either copy.
const baseAstroPath = fileURLToPath(new URL('../layouts/Base.astro', import.meta.url));
const navCssPath = fileURLToPath(new URL('../styles/nav.css', import.meta.url));
const crestDir = fileURLToPath(new URL('../../public/icons/hd/crests/', import.meta.url));

function shippedCrestSlugs(): string[] {
  return readdirSync(crestDir)
    .filter((file) => file.endsWith('.webp'))
    .map((file) => file.replace('.webp', ''))
    .sort();
}

describe('classCrestSrc', () => {
  it('resolves the upscaled crest webp path for a slug', () => {
    expect(classCrestSrc('hunter')).toBe('/icons/hd/crests/hunter.webp');
  });

  it("Base.astro's pre-paint class list is exactly the crests that ship", () => {
    const base = readFileSync(baseAstroPath, 'utf8');
    const list = /const CLASSES = \[([^\]]*)\]/.exec(base)?.[1] ?? '';
    const listed = [...list.matchAll(/'([a-z]+)'/g)].map((match) => match[1]).sort();
    expect(listed).toEqual(shippedCrestSlugs());
  });

  it('nav.css paints each shipped crest from the path this module builds, and no other', () => {
    const css = readFileSync(navCssPath, 'utf8');
    const painted = [
      ...css.matchAll(/data-pointer-class='([a-z]+)'\] \{\s*--pointer-crest: url\('([^']+)'\)/g),
    ];
    expect(painted.map((match) => match[1]).sort()).toEqual(shippedCrestSlugs());
    for (const [, slug, url] of painted) expect(url).toBe(classCrestSrc(slug));
  });
});

describe('factionMarkSrc', () => {
  it('resolves the upscaled faction emblem webp path', () => {
    expect(factionMarkSrc('horde')).toBe('/icons/hd/faction/horde.webp');
  });
});
