// web/src/lib/sim/source-rows.ts
// How the Droptimizer's source picker lays its rows out (SourcePicker.svelte): what a row says
// about its items, reputation folded to one row per faction, PvP cut to the character's own side,
// and a gated kind's sources named in one line instead of one line each.
import type { Faction } from '../bis/types';
import type { Item } from '../planner/types';
import { isKnownItem } from './sim-items';
import type { LootSource } from './loot';

export interface SourceItemStats {
  /** Items this class can try from the source. */
  count: number;
  /** The highest item level among them; 0 with none. */
  bestIlvl: number;
}

export function sourceItemStats(
  itemIds: readonly number[],
  items: ReadonlyMap<number, Item>,
  known: ReadonlySet<number> | null,
): SourceItemStats {
  let count = 0;
  let bestIlvl = 0;
  for (const id of itemIds) {
    const item = items.get(id);
    if (item === undefined || !isKnownItem(id, known)) continue;
    count += 1;
    bestIlvl = Math.max(bestIlvl, item.item_level);
  }
  return { count, bestIlvl };
}

export interface FactionGroup {
  /** The reputation faction's name. */
  name: string;
  /** One source per standing, in the file's order. */
  standings: LootSource[];
}

/** Reputation sources folded by faction name, first appearance first. */
export function groupByFaction(sources: readonly LootSource[]): FactionGroup[] {
  const groups = new Map<string, LootSource[]>();
  for (const source of sources) {
    const list = groups.get(source.name);
    if (list === undefined) groups.set(source.name, [source]);
    else list.push(source);
  }
  return [...groups].map(([name, standings]) => ({ name, standings }));
}

/** PvP ranks of the character's own side; every rank when the side is unknown. */
export function pvpForFaction(sources: readonly LootSource[], faction: Faction | null): LootSource[] {
  if (faction === null) return [...sources];
  return sources.filter((source) => source.faction === undefined || source.faction === faction);
}

export interface GatedLine {
  /** Sources that open on a known date, each with its label ("Onyxia's Lair: First raids, 9 December 2026"). */
  dated: string[];
  /** Names of sources with no date. */
  undated: string[];
}

/** A gated kind's sources split by whether they carry an opening date, for one summary line. */
export function gatedLine(
  sources: readonly LootSource[],
  gate: (source: LootSource) => string,
  noDate: string,
): GatedLine {
  const dated: string[] = [];
  const undated: string[] = [];
  for (const source of sources) {
    const label = gate(source);
    if (label === '' || label === noDate) undated.push(source.name);
    else dated.push(`${source.name}: ${label}`);
  }
  return { dated, undated };
}
