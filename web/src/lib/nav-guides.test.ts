import { describe, expect, it } from 'vitest';
import classes from '../data/classes.json';
import { readSpecCatalog } from './bis/load';
import { guideClasses, treeSpecNames } from './nav-guides';

describe('treeSpecNames', () => {
  const catalog = readSpecCatalog();

  it('lists a class in talent-tree order', () => {
    expect(treeSpecNames(catalog, 'warrior')).toEqual(['Arms', 'Fury', 'Protection']);
    expect(treeSpecNames(catalog, 'hunter')).toEqual(['Beast Mastery', 'Marksmanship', 'Survival']);
  });

  it('lists a role variant that shares a tree once', () => {
    expect(treeSpecNames(catalog, 'druid')).toEqual(['Balance', 'Feral', 'Restoration']);
  });
});

describe('guideClasses', () => {
  it('keeps classes.json order and carries the colour', () => {
    const entries = guideClasses(classes, readSpecCatalog());
    expect(entries.map((entry) => entry.slug)).toEqual(classes.map((c) => c.slug));
    expect(entries[0].color).toBe(classes[0].color);
    expect(entries.every((entry) => entry.specs.length === 3)).toBe(true);
  });
});
