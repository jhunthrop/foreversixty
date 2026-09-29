// web/src/lib/items/lookup.ts
// Two ways to get an ItemTooltipModel, for the two places ItemHover.svelte runs:
//
//   * fetchItemTooltipModel -- a planner/sim island, in the browser. Fetches the class's
//     item file and the build's loot/set files once per build (the same fetch-and-cache
//     rule sim-items.ts and loot.ts already follow), and reuses that cache for every
//     ItemHover the island opens rather than re-fetching per hover.
//   * readItemTooltipModel -- an Astro page's frontmatter, in Node at build time, the same
//     straight-off-disk read lib/bis/load.ts already uses for the /bis pages. A page calls
//     this once per item it names and passes the result as ItemHover's `model` prop, so the
//     browser never fetches anything for a page that already has the model server-side.
//
// Neither function is reachable from the other's caller: no island imports
// readItemTooltipModel (its node:fs import would have nothing to resolve to in a browser
// bundle), and no Astro page needs fetchItemTooltipModel's cache, which only pays for
// itself across many hovers in one page session.
import { existsSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { loadItems, loadSets } from '../planner/load';
import type { Item, ItemSet } from '../planner/types';
import { loadLoot, type LootFile } from '../sim/loot';
import { itemTooltipModel, type ItemTooltipModel, type ItemTooltipSources } from './tooltip';

const REPO_ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../../..');

function cacheKey(build: string, classSlug: string): string {
  return `${build}::${classSlug}`;
}

const itemMapCache = new Map<string, Promise<Map<number, Item>>>();
const sourcesCache = new Map<string, Promise<ItemTooltipSources>>();

function loadItemMap(build: string, classSlug: string): Promise<Map<number, Item>> {
  const key = cacheKey(build, classSlug);
  const cached = itemMapCache.get(key);
  if (cached !== undefined) return cached;
  const promise = loadItems(build, classSlug).then(
    (file) => new Map(file.items.map((item) => [item.id, item])),
  );
  itemMapCache.set(key, promise);
  return promise;
}

function loadSources(build: string): Promise<ItemTooltipSources> {
  const cached = sourcesCache.get(build);
  if (cached !== undefined) return cached;
  const promise = Promise.all([loadLoot(build), loadSets(build)]).then(
    ([loot, sets]): ItemTooltipSources => ({ loot, sets }),
  );
  sourcesCache.set(build, promise);
  return promise;
}

/**
 * The model for `itemId` in `classSlug`'s item file for `build`, fetched (and cached) at
 * runtime. `null` when the build's item file does not carry the id -- the same "not simmed"
 * gap wornItemLabel already names elsewhere -- so ItemHover can show a plain, tooltip-less
 * pill instead of stalling on a fetch that will never resolve to a row.
 */
export async function fetchItemTooltipModel(
  build: string,
  classSlug: string,
  itemId: number,
): Promise<ItemTooltipModel | null> {
  const [items, sources] = await Promise.all([loadItemMap(build, classSlug), loadSources(build)]);
  const row = items.get(itemId);
  return row === undefined ? null : itemTooltipModel(row, sources);
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
  const itemsFile = path.join(REPO_ROOT, 'data/builds', build, 'items', `${classSlug}.json`);
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
  const lootFile = path.join(REPO_ROOT, 'data/builds', build, 'loot.json');
  const setsFile = path.join(REPO_ROOT, 'data/builds', build, 'sets.json');
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
