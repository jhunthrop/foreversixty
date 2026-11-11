// web/src/lib/class-crest.test.ts
import { readdirSync, readFileSync } from 'node:fs';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import { classCrestSrc, crestImage, crestSlug } from './class-crest';
import { CLASS_SLUGS, type ClassSlug } from './ui/primitives';
import { factionMarkSrc } from './faction-mark';

// The character selector's first frame (nav spec 2026-10-09 section 8): Base.astro's head
// script sets `data-pointer-class` from a hand-copied class list (an `is:inline` script cannot
// import), and crest.css maps each class to its crest file. These tests pin both against the
// crest files that actually ship and the path this module builds, so a new or renamed crest
// never silently goes stale in either copy.
const baseAstroPath = fileURLToPath(new URL('../layouts/Base.astro', import.meta.url));
const crestCssPath = fileURLToPath(new URL('../styles/crest.css', import.meta.url));
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

  it("Base.astro's pre-paint class list is the one class list, and a crest ships for each", () => {
    const base = readFileSync(baseAstroPath, 'utf8');
    expect(base).toContain("data-class-slugs={CLASS_SLUGS.join(' ')}");
    expect([...CLASS_SLUGS].sort()).toEqual(shippedCrestSlugs());
  });

  it('crest.css paints each shipped crest from the path this module builds, and no other', () => {
    const css = readFileSync(crestCssPath, 'utf8');
    const painted = [
      ...css.matchAll(/data-pointer-class='([a-z]+)'\] \{\s*--pointer-crest: url\('([^']+)'\)/g),
    ];
    expect(painted.map((match) => match[1]).sort()).toEqual(shippedCrestSlugs());
    for (const [, slug, url] of painted) expect(url).toBe(classCrestSrc(slug as ClassSlug));
  });
});

describe('crestImage', () => {
  it('takes a class name or slug and rings it in the class colour', () => {
    expect(crestSlug('Warrior')).toBe('warrior');
    expect(crestImage('warrior', 36)).toEqual({
      src: '/icons/hd/crests/warrior.webp',
      ring: 'var(--color-class-warrior)',
      known: true,
    });
  });

  it('draws the LX seal in a muted ring for no class, never a person glyph or a broken request', () => {
    for (const cls of ['', null, undefined, 'Unknown']) {
      const image = crestImage(cls, 36);
      expect(image.known).toBe(false);
      expect(image.src).toMatch(/seal/);
      expect(image.ring).toBe('var(--color-muted)');
    }
  });
});

describe('factionMarkSrc', () => {
  it('resolves the upscaled faction emblem webp path', () => {
    expect(factionMarkSrc('horde')).toBe('/icons/hd/faction/horde.webp');
  });
});
