// web/src/lib/search/item-uses.ts
// `/data/{build}/item-uses/{shard}.json` (Wave 3 contract W3-C8): for each item, where it comes
// from (`item-sources.ts`' rows, the tooltip's own words) and which specs' best-in-slot lists use
// it. Sharded by id (`id % ITEM_USES_SHARDS`) so the apps' item page (spec F12.4) fetches one
// small file on demand rather than every spec's list. Pure.
import { itemSourceRows, type ItemSourceRow } from '../items/item-sources';
import { itemsOfSource, type LootFile, type LootSource } from '../sim/loot';
import type { BisBand, BisFile, SpecCatalogEntry } from '../bis/types';
import { ITEM_USES_SHARDS } from './search-index';

export type ItemUseSource = Omit<ItemSourceRow, 'primary' | 'text'>;

export interface ItemUseSpec {
  /** The BiS file's spec key (`shaman-elemental`). */
  spec: string;
  class: string;
  spec_name: string;
  /** The level band of the spec's best placement of the item. */
  band: number;
  preset: string;
  /** Null when both factions' lists place the item the same way. */
  faction: 'alliance' | 'horde' | null;
  /** 0 is the pick; n is runner-up n. */
  rank: number;
}

export interface ItemUses {
  sources: ItemUseSource[];
  /** One row per spec: picks before runners-up, then class and spec in the catalog's order. */
  specs: ItemUseSpec[];
}

export interface ItemUsesShard {
  build: string;
  shard: number;
  shards: number;
  /** Item id (string key) -> its uses; an item with no source and no list is absent. */
  items: Record<string, ItemUses>;
}

export interface ItemUsesInput {
  build: string;
  /** The items the search index carries; only these get an entry. */
  itemIds: ReadonlySet<number>;
  loot: LootFile;
  bisFiles: readonly BisFile[];
  specs: readonly SpecCatalogEntry[];
}

interface Placement {
  spec: string;
  band: number;
  preset: string;
  faction: 'alliance' | 'horde';
  rank: number;
}

function bandPlacements(file: BisFile, band: BisBand): [number, Placement][] {
  const base = {
    spec: file.spec,
    band: band.band,
    preset: band.preset ?? 'bare',
    faction: band.faction,
  };
  return band.slots.flatMap((slot) => {
    const out: [number, Placement][] = [];
    if (typeof slot.item_id === 'number' && slot.item_id > 0) out.push([slot.item_id, { ...base, rank: 0 }]);
    (slot.alternatives ?? []).forEach((alternative, index) => {
      out.push([alternative.item_id, { ...base, rank: index + 1 }]);
    });
    return out;
  });
}

/** Lower rank first, then the higher band, then the raid preset. */
function comparePlacements(a: Placement, b: Placement): number {
  if (a.rank !== b.rank) return a.rank - b.rank;
  if (a.band !== b.band) return b.band - a.band;
  if (a.preset !== b.preset) return a.preset === 'raid' ? -1 : 1;
  return 0;
}

function bestPlacement(
  placements: readonly Placement[],
): Omit<Placement, 'faction'> & { faction: Placement['faction'] | null } {
  const [best, ...rest] = [...placements].sort(comparePlacements);
  const sameElsewhere = rest.some(
    (other) => comparePlacements(best, other) === 0 && other.faction !== best.faction,
  );
  return { ...best, faction: sameElsewhere ? null : best.faction };
}

function placementsByItem(bisFiles: readonly BisFile[]): Map<number, Map<string, Placement[]>> {
  const byItem = new Map<number, Map<string, Placement[]>>();
  for (const file of bisFiles) {
    for (const band of file.bands) {
      for (const [itemId, placement] of bandPlacements(file, band)) {
        const bySpec = byItem.get(itemId) ?? new Map<string, Placement[]>();
        bySpec.set(placement.spec, [...(bySpec.get(placement.spec) ?? []), placement]);
        byItem.set(itemId, bySpec);
      }
    }
  }
  return byItem;
}

function specRows(
  bySpec: Map<string, Placement[]> | undefined,
  specs: readonly SpecCatalogEntry[],
): ItemUseSpec[] {
  if (!bySpec) return [];
  const order = new Map(specs.map((spec, index) => [spec.spec, index]));
  return [...bySpec.entries()]
    .flatMap(([specKey, placements]) => {
      const catalog = specs.find((spec) => spec.spec === specKey);
      if (!catalog) return [];
      return [{ ...bestPlacement(placements), class: catalog.class_slug, spec_name: catalog.name }];
    })
    .sort((a, b) => {
      if ((a.rank === 0) !== (b.rank === 0)) return a.rank === 0 ? -1 : 1;
      return (order.get(a.spec) ?? 0) - (order.get(b.spec) ?? 0);
    })
    .map(({ spec, class: classSlug, spec_name, band, preset, faction, rank }) => ({
      spec,
      class: classSlug,
      spec_name,
      band,
      preset,
      faction,
      rank,
    }));
}

/** item id -> the loot sources that name it, so each item's rows are read from its own few
 *  sources rather than the whole file (11,000 items over 3,000 sources). */
function sourcesByItem(loot: LootFile): Map<number, LootSource[]> {
  const index = new Map<number, LootSource[]>();
  for (const source of loot.sources) {
    for (const itemId of new Set(itemsOfSource(source))) {
      index.set(itemId, [...(index.get(itemId) ?? []), source]);
    }
  }
  return index;
}

function useSources(itemId: number, sources: readonly LootSource[], loot: LootFile): ItemUseSource[] {
  return itemSourceRows(itemId, { ...loot, sources: [...sources] }).map(
    ({ primary: _primary, text: _text, ...row }) => row,
  );
}

export function buildItemUses(input: ItemUsesInput): ItemUsesShard[] {
  const placements = placementsByItem(input.bisFiles);
  const lootByItem = sourcesByItem(input.loot);
  const shards: ItemUsesShard[] = Array.from({ length: ITEM_USES_SHARDS }, (_, shard) => ({
    build: input.build,
    shard,
    shards: ITEM_USES_SHARDS,
    items: {},
  }));
  const ids = [...input.itemIds].sort((a, b) => a - b);
  for (const itemId of ids) {
    const sources = useSources(itemId, lootByItem.get(itemId) ?? [], input.loot);
    const specs = specRows(placements.get(itemId), input.specs);
    if (sources.length === 0 && specs.length === 0) continue;
    const shard = shards[itemId % ITEM_USES_SHARDS];
    shards[itemId % ITEM_USES_SHARDS] = {
      ...shard,
      items: { ...shard.items, [String(itemId)]: { sources, specs } },
    };
  }
  return shards;
}
