// The Guides door's panel data (spec 2026-10-15-grouped-header.md section 2): the nine
// classes in classes.json order, each with its spec names in talent-tree order. Pure, so the
// build-time catalogue read stays in Header.astro and this is testable on its own.
import type { SpecCatalogEntry } from './bis/types';

export interface GuideClass {
  readonly slug: string;
  readonly name: string;
  readonly color: string;
  readonly specs: readonly string[];
}

interface ClassIdentity {
  readonly slug: string;
  readonly name: string;
  readonly color: string;
}

/**
 * Spec names for one class, in `tree_index` order. A role variant sharing a tree with a
 * spec already listed (Feral Bear beside Feral) is one talent tree, so it is listed once.
 */
export function treeSpecNames(catalog: readonly SpecCatalogEntry[], classSlug: string): string[] {
  const seenTrees = new Set<number>();
  return catalog
    .filter((entry) => entry.class_slug === classSlug)
    .sort((a, b) => a.tree_index - b.tree_index)
    .filter((entry) => {
      if (seenTrees.has(entry.tree_index)) return false;
      seenTrees.add(entry.tree_index);
      return true;
    })
    .map((entry) => entry.name);
}

export function guideClasses(
  classes: readonly ClassIdentity[],
  catalog: readonly SpecCatalogEntry[],
): GuideClass[] {
  return classes.map((c) => ({ ...c, specs: treeSpecNames(catalog, c.slug) }));
}
