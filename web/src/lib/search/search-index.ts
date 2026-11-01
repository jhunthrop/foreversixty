// web/src/lib/search/search-index.ts
// `/data/{build}/search.json` (Wave 3 contract W3-C8): the one file the apps' search reads to
// find items, talents and guides on the phone, offline. Names, quality, slot, levels, icon and
// the classes that may use each item the class files carry, every talent of every class, and
// every written guide. Pure: the endpoint (`pages/data/[build]/search.json.ts`) reads the
// build's files and hands them here.
import { gzipSync } from 'node:zlib';
import type { Item, TalentFile } from '../planner/types';
import type { SpecCatalogEntry } from '../bis/types';

/** How many `item-uses/{shard}.json` files the build splits its item uses into; an item's
 *  shard is `id % ITEM_USES_SHARDS`. Published in search.json so the apps never hardcode it. */
export const ITEM_USES_SHARDS = 64;

/** The spec's ceiling on search.json's gzipped size (spec §5.4 Data). */
export const SEARCH_INDEX_MAX_GZIP_BYTES = 300 * 1024;

export interface SearchItem {
  id: number;
  name: string;
  icon: string;
  quality: number;
  slot: string;
  item_level: number;
  required_level: number;
  /** Class slugs whose items file carries the item, in the class files' order. */
  classes: string[];
}

export interface SearchTalent {
  id: number;
  name: string;
  icon: string;
  class: string;
  /** The talent tree's own name ("Elemental", "Feral Combat"). */
  tree: string;
  /** The spec the tree is played as ("Feral" for Feral Combat); the tree's name when no
   *  written spec names it. */
  spec_name: string;
  max_rank: number;
}

export interface SearchGuide {
  class: string;
  /** The spec's slug inside its class (`elemental`), the guide's path segment. */
  spec: string;
  spec_name: string;
  role: string;
  icon: string;
}

export interface SearchIndex {
  build: string;
  item_uses_shards: number;
  items: SearchItem[];
  talents: SearchTalent[];
  guides: SearchGuide[];
}

export interface ClassItemsInput {
  classSlug: string;
  items: readonly Pick<Item, 'id' | 'name' | 'icon' | 'quality' | 'slot' | 'item_level' | 'required_level'>[];
}

export interface SearchIndexInput {
  build: string;
  /** One entry per class items file, in class order. */
  classItems: readonly ClassItemsInput[];
  talents: readonly TalentFile[];
  /** The written specs (`data/curated/specs.json`) with the extra `icon` stem it carries. */
  specs: readonly (SpecCatalogEntry & { icon?: string })[];
  /** `{class}/{spec}` of every spec with a written guide. */
  guidePaths: ReadonlySet<string>;
}

function searchItems(classItems: readonly ClassItemsInput[]): SearchItem[] {
  const byId = new Map<number, SearchItem>();
  for (const { classSlug, items } of classItems) {
    for (const item of items) {
      const known = byId.get(item.id);
      if (known) {
        byId.set(item.id, { ...known, classes: [...known.classes, classSlug] });
        continue;
      }
      byId.set(item.id, {
        id: item.id,
        name: item.name,
        icon: item.icon,
        quality: item.quality,
        slot: item.slot,
        item_level: item.item_level,
        required_level: item.required_level,
        classes: [classSlug],
      });
    }
  }
  return [...byId.values()].sort((a, b) => a.id - b.id);
}

function treeSpecName(
  specs: SearchIndexInput['specs'],
  classSlug: string,
  position: number,
  tree: string,
): string {
  return specs.find((spec) => spec.class_slug === classSlug && spec.tree_index === position)?.name ?? tree;
}

function searchTalents(talents: readonly TalentFile[], specs: SearchIndexInput['specs']): SearchTalent[] {
  return talents.flatMap((file) =>
    file.trees.flatMap((tree) => {
      const specName = treeSpecName(specs, file.class_slug, tree.position, tree.name);
      return tree.talents.map((talent) => ({
        id: talent.id,
        name: talent.name,
        icon: talent.icon,
        class: file.class_slug,
        tree: tree.name,
        spec_name: specName,
        max_rank: talent.max_rank,
      }));
    }),
  );
}

function searchGuides(specs: SearchIndexInput['specs'], guidePaths: ReadonlySet<string>): SearchGuide[] {
  return specs
    .filter((spec) => guidePaths.has(`${spec.class_slug}/${spec.spec_slug}`))
    .map((spec) => ({
      class: spec.class_slug,
      spec: spec.spec_slug,
      spec_name: spec.name,
      role: spec.role,
      icon: spec.icon ?? '',
    }));
}

export function buildSearchIndex(input: SearchIndexInput): SearchIndex {
  return {
    build: input.build,
    item_uses_shards: ITEM_USES_SHARDS,
    items: searchItems(input.classItems),
    talents: searchTalents(input.talents, input.specs),
    guides: searchGuides(input.specs, input.guidePaths),
  };
}

/** The file's JSON; throws when it gzips over `maxGzipBytes`, so the site build fails rather
 *  than ship an index the phones should not download (spec §5.4 Data). */
export function encodeSearchIndex(
  index: SearchIndex,
  maxGzipBytes: number = SEARCH_INDEX_MAX_GZIP_BYTES,
): string {
  const json = JSON.stringify(index);
  const size = gzipSync(json).length;
  if (size > maxGzipBytes) {
    throw new Error(
      `search.json for ${index.build} is ${size} bytes gzipped, over the ${maxGzipBytes}-byte ceiling`,
    );
  }
  return json;
}
