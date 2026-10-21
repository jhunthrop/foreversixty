// web/src/lib/planner/items-ready.ts
// The planner's item catalogue is not part of the page's boot (the gear panel left the page), but
// two things need it: the draft that goes to the API carries only gear the class catalogue lists,
// and the character card names what it left out. Both ask here; the file loads once per class.
import { loadItems } from './load';
import type { PlannerStore } from './store.svelte';
import type { ItemFile } from './types';

const inFlight = new Map<string, Promise<void>>();

/**
 * Loads the current class's item file into `store` unless it is already there. Never throws: a
 * failed read leaves the store without a catalogue, and a draft then carries the gear as worn
 * (the API's own answer is the next word).
 */
export function ensureItems(
  store: Pick<PlannerStore, 'itemIndex' | 'classSlug' | 'treeVersion' | 'setItems'>,
  load: (build: string, classSlug: string) => Promise<ItemFile> = loadItems,
): Promise<void> {
  if (store.itemIndex.size > 0) return Promise.resolve();
  const build = store.treeVersion;
  const slug = store.classSlug;
  const key = `${build}|${slug}`;
  const pending = inFlight.get(key);
  if (pending !== undefined) return pending;
  const run = load(build, slug)
    .then((file) => {
      // The class may have changed while the file was on its way.
      if (store.classSlug === slug && store.treeVersion === build) store.setItems(file);
    })
    .catch(() => undefined)
    .finally(() => inFlight.delete(key));
  inFlight.set(key, run);
  return run;
}
