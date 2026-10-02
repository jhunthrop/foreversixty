// web/src/lib/class-crest.test.ts
import { readFileSync } from 'node:fs';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import { classCrestSrc } from './class-crest';
import { factionMarkSrc } from './faction-mark';

// Owner-reported defect, 2026-10-01 (the account chip "refreshes on every page load"):
// Base.astro's pre-paint account-chip script fills the crest/faction images directly, from
// localStorage, before AccountMenu.svelte hydrates -- an `is:inline` script cannot import
// this module or faction-mark.ts, so it hand-duplicates their exact template literal. These
// two tests pin that literal against the one this module (and faction-mark.ts) actually
// builds, so a path change to either source never silently goes stale in the other.
const baseAstroPath = fileURLToPath(new URL('../layouts/Base.astro', import.meta.url));

describe('classCrestSrc', () => {
  it('resolves the upscaled crest webp path for a slug', () => {
    expect(classCrestSrc('hunter')).toBe('/icons/hd/crests/hunter.webp');
  });

  it("Base.astro's pre-paint script builds the identical path pattern", () => {
    const base = readFileSync(baseAstroPath, 'utf8');
    const prefix = classCrestSrc('__SLUG__').replace('__SLUG__.webp', '');
    expect(base).toContain(`\`${prefix}\${slug}.webp\``);
  });
});

describe('factionMarkSrc', () => {
  it('resolves the upscaled faction emblem webp path', () => {
    expect(factionMarkSrc('horde')).toBe('/icons/hd/faction/horde.webp');
  });

  it("Base.astro's pre-paint script builds the identical path pattern", () => {
    const base = readFileSync(baseAstroPath, 'utf8');
    const prefix = factionMarkSrc('horde').replace('horde.webp', '');
    expect(base).toContain(`\`${prefix}\${main.faction}.webp\``);
  });
});
