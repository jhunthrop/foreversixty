// web/src/lib/gear/qualities.ts
// The item qualities the Gear section colours its names with (Task L8.6 ruling 1). The API's
// gear panel carries ids and names, not qualities; the build's items/<class>.json carries them,
// the same file the planner, the home upgrades panel and the item tooltip read (planner/load.ts's
// loadItems, cached by URL). One read per class for the life of the page; a class the build has
// no file for, or a failed read, gives an empty map, and every name keeps the common colour.
import { loadItems } from '../planner/load';

const cache = new Map<string, Promise<ReadonlyMap<number, number>>>();

export function itemQualities(build: string, classSlug: string): Promise<ReadonlyMap<number, number>> {
  const key = `${build}/${classSlug}`;
  const cached = cache.get(key);
  if (cached !== undefined) return cached;
  const promise = loadItems(build, classSlug)
    .then((file) => new Map(file.items.map((item) => [item.id, item.quality] as const)))
    .catch(() => new Map<number, number>());
  cache.set(key, promise);
  return promise;
}
