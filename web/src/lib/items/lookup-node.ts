// web/src/lib/items/lookup-node.ts
// The build-time half of lib/items/lookup.ts: an Astro page's frontmatter, in Node, reads
// the model for an item straight off data/builds/<build>/items/<classSlug>.json, the same
// straight-off-disk read lib/bis/load.ts uses for the /bis pages, and passes it as
// ItemHover's `model` prop so the browser never fetches anything for that page. This file
// imports node:fs, so nothing the browser loads may import it (see lookup.ts's header).
import { existsSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import type { Item, ItemSet } from '../planner/types';
import type { LootFile } from '../sim/loot';
import { itemTooltipModel, type ItemTooltipModel, type ItemTooltipSources } from './tooltip';

function repoRoot(): string {
  return path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../../..');
}

function cacheKey(build: string, classSlug: string): string {
  return `${build}::${classSlug}`;
}

function readJson<T>(file: string): T {
  return JSON.parse(readFileSync(file, 'utf8')) as T;
}

// Build-time reads, cached in-process: a class's items/<classSlug>.json runs 1.5-3.4 MB on
// the real build (warrior.json alone is over 3 MB), and a page like /bis renders a row per
// slot per band per faction -- a couple hundred `readItemTooltipModel` calls in one page,
// every written spec's own page besides. Re-reading and re-JSON.parse-ing multi-megabyte
// files that many times per page (and again for every other spec page in the same class)
// would cost real `astro build` minutes for no benefit: these files never change mid-build.
const buildTimeItemsCache = new Map<string, Item[]>();
const buildTimeSourcesCache = new Map<string, ItemTooltipSources>();

function readItemsForBuildTime(build: string, classSlug: string): Item[] | null {
  const key = cacheKey(build, classSlug);
  const cached = buildTimeItemsCache.get(key);
  if (cached !== undefined) return cached;
  const itemsFile = path.join(repoRoot(), 'data/builds', build, 'items', `${classSlug}.json`);
  if (!existsSync(itemsFile)) return null;
  const { items } = readJson<{ items: Item[] }>(itemsFile);
  buildTimeItemsCache.set(key, items);
  return items;
}

/** Build-time equivalent of loadSources above, reading straight off disk. Cached per build
 *  for the same reason readItemsForBuildTime is. */
function readSources(build: string): ItemTooltipSources {
  const cached = buildTimeSourcesCache.get(build);
  if (cached !== undefined) return cached;
  const root = repoRoot();
  const lootFile = path.join(root, 'data/builds', build, 'loot.json');
  const setsFile = path.join(root, 'data/builds', build, 'sets.json');
  const loot: LootFile = existsSync(lootFile) ? readJson<LootFile>(lootFile) : { sources: [] };
  const sets: ItemSet[] = existsSync(setsFile) ? readJson<ItemSet[]>(setsFile) : [];
  const sources: ItemTooltipSources = { loot, sets };
  buildTimeSourcesCache.set(build, sources);
  return sources;
}

/**
 * The model for `itemId`, read straight off `data/builds/<build>/items/<classSlug>.json` at
 * build time -- for a static page (like /bis) to pass as ItemHover's `model` prop, the same
 * way `lib/bis/load.ts` reads the BiS file itself. `null` when the build ships no item file
 * for the class, or the file does not carry `itemId` -- never a guess at what the item is.
 */
export function readItemTooltipModel(
  build: string,
  classSlug: string,
  itemId: number,
): ItemTooltipModel | null {
  const items = readItemsForBuildTime(build, classSlug);
  if (items === null) return null;
  const row = items.find((item) => item.id === itemId);
  return row === undefined ? null : itemTooltipModel(row, readSources(build));
}
