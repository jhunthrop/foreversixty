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
  /** That item, the row's icon and (for a PvP reward) its name; null with none. */
  best: Item | null;
}

export function sourceItemStats(
  itemIds: readonly number[],
  items: ReadonlyMap<number, Item>,
  known: ReadonlySet<number> | null,
): SourceItemStats {
  let count = 0;
  let best: Item | null = null;
  for (const id of itemIds) {
    const item = items.get(id);
    if (item === undefined || !isKnownItem(id, known)) continue;
    count += 1;
    if (best === null || item.item_level > best.item_level) best = item;
  }
  return { count, bestIlvl: best?.item_level ?? 0, best };
}

/** Kinds whose sources the picker leaves out when this class can use nothing from them. */
const HIDE_EMPTY: ReadonlySet<string> = new Set([
  'dungeon',
  'crafted',
  'rep',
  'pvp',
  'world_drop',
  'zone',
  'vendor',
]);

/** Whether a source's row is worth drawing: anything with a usable item, and every raid, world boss and quest. */
export function showsSource(kind: string, count: number): boolean {
  return count > 0 || !HIDE_EMPTY.has(kind);
}

/** How many of a kind's sources have a tick on them, whole or by boss (picks are `<source id>|<boss id>`). */
export function tickedSources(sources: readonly LootSource[], picked: readonly string[]): number {
  const ids = new Set(sources.map((source) => source.id));
  return new Set(picked.map((key) => key.slice(0, key.indexOf('|'))).filter((id) => ids.has(id))).size;
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
  /** One entry per opening date, its sources named once each: "First raids, 9 December 2026: Blacksmithing, Leatherworking". */
  dated: string[];
  /** Names of sources with no date, each once. */
  undated: string[];
}

/** A gated kind's sources grouped by their opening date, every name once, for one summary line. */
export function gatedLine(
  sources: readonly LootSource[],
  gate: (source: LootSource) => string,
  noDate: string,
): GatedLine {
  const byDate = new Map<string, Set<string>>();
  const undated = new Set<string>();
  for (const source of sources) {
    const label = gate(source);
    if (label === '' || label === noDate) {
      undated.add(source.name);
      continue;
    }
    const names = byDate.get(label) ?? new Set<string>();
    names.add(source.name);
    byDate.set(label, names);
  }
  return {
    dated: [...byDate].map(([label, names]) => `${label}: ${[...names].join(', ')}`),
    undated: [...undated],
  };
}
