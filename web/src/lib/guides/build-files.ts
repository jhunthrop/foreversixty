// web/src/lib/guides/build-files.ts
// The guide pages read the active build's tables straight off data/builds/<build>/ at build
// time, so a guide says what the build says under either FOREVER_DATA: the planner fixture
// only swaps what src/data/generated/ and public/data/ carry, never these files. Every guide
// loader that touches a build's files resolves them here.
import { existsSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

/** The repository root, from this file's own location. */
export function repoRoot(): string {
  return path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../../..');
}

/** The path of one of `build`'s files under data/builds/. */
export function buildFile(build: string, ...parts: string[]): string {
  return path.join(repoRoot(), 'data/builds', build, ...parts);
}

/** A build's JSON file, parsed; `undefined` when the build has none. */
export function readBuildJson<T>(build: string, file: string): T | undefined {
  const full = buildFile(build, file);
  return existsSync(full) ? (JSON.parse(readFileSync(full, 'utf8')) as T) : undefined;
}

/** A build's JSON file the pipeline publishes for every build; a missing one is a broken
 *  checkout, named rather than rendered as an empty table. */
export function requireBuildJson<T>(build: string, file: string): T {
  const value = readBuildJson<T>(build, file);
  if (value === undefined) {
    throw new Error(`data/builds/${build}/${file} is missing; the pipeline publishes it for every build`);
  }
  return value;
}

/** One row of a build's `items.json`: what the client says an item is. */
export interface CatalogItem {
  id: number;
  name: string;
  quality: number;
  item_level: number;
  required_level: number;
  /** ItemClass: 2 weapon, 4 armor. */
  class_id: number;
  subclass_id: number;
  /** InventoryType: 13 one-hand, 14 shield, 17 two-hand, 23 held in off hand. */
  inventory_type: number;
}

const catalogCache = new Map<string, Map<number, CatalogItem>>();

/** `build`'s `items.json`, by item id, read once per build. */
export function buildItemCatalog(build: string): ReadonlyMap<number, CatalogItem> {
  const cached = catalogCache.get(build);
  if (cached !== undefined) return cached;
  const rows = readBuildJson<CatalogItem[]>(build, 'items.json') ?? [];
  const indexed = new Map(rows.map((row) => [row.id, row]));
  catalogCache.set(build, indexed);
  return indexed;
}
