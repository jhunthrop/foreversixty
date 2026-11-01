// web/src/lib/app-data/guide-index.ts
// /data/{build}/guides/index.json (Wave 3 contract W3-C1): every spec guide the apps can open,
// grouped by class in classes.json's order and by spec in the talent trees' order. The Guides
// tab and the class page draw from this file alone; a guide's body is fetched when opened.
import classes from '../../data/classes.json';
import type { SpecCatalogEntry } from '../bis/types';
import type { AppGuideFile, GuideRole } from './guide-file';

export interface AppGuideIndexSpec {
  spec: string;
  spec_slug: string;
  spec_name: string;
  role: GuideRole;
  icon: string | null;
  updated: string;
  /** The guide's file, relative to `/data/{build}/`. */
  path: string;
}

export interface AppGuideIndexClass {
  class: string;
  class_name: string;
  specs: AppGuideIndexSpec[];
}

export interface AppGuideIndex {
  build: string;
  classes: AppGuideIndexClass[];
}

/** `guides/shaman-elemental.json`: where a spec's guide file sits under the build's directory. */
export function guideFilePath(spec: string): string {
  return `guides/${spec}.json`;
}

function treeOrder(catalog: readonly SpecCatalogEntry[], spec: string): number {
  return catalog.find((entry) => entry.spec === spec)?.tree_index ?? Number.MAX_SAFE_INTEGER;
}

/** Classes in classes.json's order, each class's specs in its talent trees' order. */
export function buildGuideIndex(
  build: string,
  guides: readonly AppGuideFile[],
  catalog: readonly SpecCatalogEntry[],
): AppGuideIndex {
  return {
    build,
    classes: classes.map((row) => ({
      class: row.slug,
      class_name: row.name,
      specs: guides
        .filter((guide) => guide.class === row.slug)
        .sort((a, b) => treeOrder(catalog, a.spec) - treeOrder(catalog, b.spec))
        .map((guide) => ({
          spec: guide.spec,
          spec_slug: guide.spec_slug,
          spec_name: guide.spec_name,
          role: guide.role,
          icon: guide.icon,
          updated: guide.updated,
          path: guideFilePath(guide.spec),
        })),
    })),
  };
}
