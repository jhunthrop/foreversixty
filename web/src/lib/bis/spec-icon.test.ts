// web/src/lib/bis/spec-icon.test.ts
import { existsSync, readdirSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import activeBuild from '../../data/active-build.json';
import { readSpecCatalog } from './load';
import { publishedIcon, specCardIcon } from './spec-icon';

const ICONS = path.resolve(
  path.dirname(fileURLToPath(import.meta.url)),
  '../../../public/data',
  activeBuild.build,
  'icons',
);

describe('publishedIcon', () => {
  it('passes nothing through as nothing', () => {
    expect(publishedIcon(activeBuild.build, undefined)).toBeUndefined();
  });

  it('drops an icon this build does not publish', () => {
    expect(publishedIcon(activeBuild.build, 'no_such_icon_anywhere')).toBeUndefined();
  });

  it.skipIf(!existsSync(ICONS))('keeps an icon this build publishes', () => {
    const stem = readdirSync(ICONS)
      .find((file) => file.endsWith('.webp'))
      ?.replace(/\.webp$/, '');
    if (stem === undefined) return;
    expect(publishedIcon(activeBuild.build, stem)).toBe(stem);
  });
});

describe('specCardIcon', () => {
  it('is the catalogue icon', () => {
    expect(specCardIcon({ spec: 'mage-fire', icon: 'spell_fire_firebolt02' })).toBe('spell_fire_firebolt02');
  });

  it('gives every spec of a class its own icon, so Feral and Feral Bear read apart', () => {
    const catalog = readSpecCatalog();
    for (const classSlug of new Set(catalog.map((s) => s.class_slug))) {
      const icons = catalog.filter((s) => s.class_slug === classSlug).map(specCardIcon);
      expect(new Set(icons).size, classSlug).toBe(icons.length);
    }
  });
});
