// web/src/lib/items/lookup.ts
// Two ways to get an ItemTooltipModel, for the two places ItemHover.svelte runs:
//
//   * fetchItemTooltipModel -- a planner/sim island, in the browser. Fetches the class's
//     item file and the build's loot/set files once per build (the same fetch-and-cache
//     rule sim-items.ts and loot.ts already follow), and reuses that cache for every
//     ItemHover the island opens rather than re-fetching per hover.
//   * readItemTooltipModel -- an Astro page's frontmatter, in Node at build time: it lives
//     in lookup-node.ts, beside this file, and nothing the browser loads imports it.
//
// ItemHover.svelte imports fetchItemTooltipModel from this module for every island
// (planner, sim, sim-tools), so this file DOES reach the browser. The build-time half used
// to live here too, behind a lazily-called `repoRoot()`, which a production bundle
// tree-shook away -- but the dev server serves modules unbundled, and Vite's browser shim
// for node:fs throws on the first property read, so every sim and planner island crashed
// on `npm run dev` the moment ItemHover loaded (21 Oct 2026). The split keeps node:fs out
// of anything the browser can reach.
import { loadItems, loadSets } from '../planner/load';
import type { Item } from '../planner/types';
import { loadLoot } from '../sim/loot';
import { itemTooltipModel, type ItemTooltipModel, type ItemTooltipSources } from './tooltip';

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
