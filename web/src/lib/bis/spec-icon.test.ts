// web/src/lib/bis/spec-icon.test.ts
import { describe, expect, it } from 'vitest';
import { readSpecCatalog } from './load';
import { specCardIcon } from './spec-icon';

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
