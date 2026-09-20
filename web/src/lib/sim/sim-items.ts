// web/src/lib/sim/sim-items.ts
// data/builds/<build>/simitems.json: the item ids the engine's embedded SimDatabase
// actually carries (pipeline/simdb/items.py's simdb_item_rows). For a build whose
// items/<class>.json was computed from the SAME raw tables, that file's ids are a subset
// of this one's (build_class_items applies simdb_item_rows's own filter plus an extra "has
// gear value" clause -- see that module's docstring) -- but not every build's items/ came
// from the same tables: forever-prebeta's is copied wholesale from a different source
// build's own ItemSparse (data/pipeline/forever.py), so it can name an id (e.g. 16963,
// Helm of Wrath) this build's ItemSparse never had a row for at all, and therefore an id
// this file does not carry either. sim/bulk refuses any candidate it does not recognise
// outright: `{"error": "bulk: the build has no such item: <id>"}`. Every candidate source
// on the bulk pages (item search, bag, bank, named sets, Droptimizer picks) filters
// against this file so a player is never offered -- or silently sends -- a candidate the
// engine is about to refuse.
//
// Optional, the same way loot.json/enchants.json/suffixes.json are (planner/load.ts's
// loadOptional): a build the data lane has not regenerated this file for ships none, and
// an absent file means "nothing is known to filter against" rather than "everything is
// unknown" -- see knownItemIds below.
import { dataUrl, loadOptional } from '../planner/load';

export interface SimItemsFile {
  build: string;
  items: number[];
}

export function loadSimItems(build: string): Promise<SimItemsFile | null> {
  return loadOptional<SimItemsFile | null>(dataUrl(build, 'simitems.json'), null);
}

/**
 * `null` means the build ships no simitems.json -- there is nothing to filter against, so
 * every id is treated as known (the pre-fix behaviour, and the only sound default for a
 * build the data lane has not regenerated this file for yet). A `Set`, even an empty one,
 * means the file loaded and its ids are the whole truth.
 */
export function knownItemIds(file: SimItemsFile | null): Set<number> | null {
  return file === null ? null : new Set(file.items);
}

/**
 * Whether the engine's embedded database carries `itemId`. `known === null` means no
 * simitems.json loaded for this build -- nothing to filter against, so every id passes.
 */
export function isKnownItem(itemId: number, known: ReadonlySet<number> | null): boolean {
  return known === null || known.has(itemId);
}
