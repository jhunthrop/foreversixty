// web/src/lib/sim/loot.ts
// data/builds/<build>/loot.json: where every item comes from, contract 6.1 as corrected by
// 10.4 -- AreaTable zone ids, `<source>:<npc-id>` boss ids, `opens: "later"` for an
// unknown date, and an empty boss name where neither database had one.
//
// The file is per build and item-id keyed through its sources, so the page never asks the
// API where an item drops -- it is the same fetch-and-cache path items.json already takes.
// Drop chances are in neither the fork's database nor the curated overlay, so nothing here
// computes or exposes a probability.
//
// Contract 10.4 also warns that the re-itemised raid tier is thin: 1,809 of the fork's
// sourced item ids do not exist in the 1.60 client, loot.json lists only items the build
// has, and the first curated overlay records the gap per raid in its notes. The page shows
// what the file contains and counts what it does not -- it never fabricates a drop.
import { dataUrl, loadOptional } from '../planner/load';
import { bulkCopy } from './copy';
import { hasOpened, type PhaseRow } from './phase';
import { poolQualityCopy } from './pool-quality-copy';

export const LOOT_KINDS = ['raid', 'dungeon', 'world', 'crafted', 'rep', 'pvp', 'quest'] as const;
export type LootKind = (typeof LOOT_KINDS)[number];

export const SOURCE_KIND_LABELS: Record<LootKind, string> = {
  raid: bulkCopy.sourcesRaids,
  dungeon: bulkCopy.sourcesDungeons,
  world: bulkCopy.sourcesWorld,
  crafted: bulkCopy.sourcesCrafted,
  rep: bulkCopy.sourcesRep,
  pvp: bulkCopy.sourcesPvp,
  quest: bulkCopy.sourcesQuests,
};

/** Off unless the player asks: a quest reward is a one-time source (design 6.1). */
export const DEFAULT_OFF_KINDS: readonly LootKind[] = ['quest'];

export interface LootBoss {
  /** `<source id>:<npc-id>` (contract 10.4). */
  id: string;
  /** Empty when neither database names the boss; the picker shows the id in that case. */
  name: string;
  npc_id?: number;
  items: number[];
}

export interface LootSource {
  /** `raid:<zone-slug>`, `dungeon:<zone-slug>`, `pvp:rank-<n>`, … (contract 10.4). */
  id: string;
  kind: LootKind;
  name: string;
  /** An AreaTable id -- Molten Core is 2717, not 409 (contract 10.4). A `world` source
   *  carries none at all. */
  zone_id?: number;
  /**
   * A phase name from the phase table; the literal `"later"` for a source whose date is
   * unknown, which the page shows as unreleased without a date; absent means open from
   * launch (contract 6.1 and 10.4).
   */
  opens?: string;
  bosses?: LootBoss[];
  trash?: number[];
  items?: number[];
  profession?: string;
  faction_id?: number;
  standing?: string;
  rank?: number;
}

export interface LootFile {
  sources: LootSource[];
}

const EMPTY: LootFile = { sources: [] };

/**
 * The build's loot table. A build the data lane has not regenerated ships none, and a 404
 * resolves to an empty file rather than throwing -- the Droptimizer then says it has no
 * sources, which is true, instead of showing an error for a file that was never promised.
 * Every other failure is a broken build and is rethrown.
 */
export async function loadLoot(build: string): Promise<LootFile> {
  return loadOptional<LootFile>(dataUrl(build, 'loot.json'), EMPTY);
}

/** Every item a source yields: its own list, every boss's, and its trash. */
export function itemsOfSource(source: LootSource): number[] {
  return [
    ...(source.items ?? []),
    ...(source.bosses ?? []).flatMap((boss) => boss.items),
    ...(source.trash ?? []),
  ];
}

export function itemsOfBoss(source: LootSource, bossId: string): number[] {
  return (source.bosses ?? []).find((boss) => boss.id === bossId)?.items ?? [];
}

/** item id -> the source ids it drops from. Built once per loaded file. */
export function sourcesByItem(file: LootFile): Map<number, string[]> {
  const index = new Map<number, string[]>();
  for (const source of file.sources) {
    for (const item of new Set(itemsOfSource(source))) {
      const existing = index.get(item);
      if (existing === undefined) index.set(item, [source.id]);
      else if (!existing.includes(source.id)) existing.push(source.id);
    }
  }
  return index;
}

export interface SourceGroup {
  kind: LootKind;
  label: string;
  sources: LootSource[];
}

/** The picker's groups, in the design's own order, skipping kinds this build has none of. */
export function groupSources(file: LootFile): SourceGroup[] {
  return LOOT_KINDS.flatMap((kind) => {
    const sources = file.sources.filter((source) => source.kind === kind);
    return sources.length === 0 ? [] : [{ kind, label: SOURCE_KIND_LABELS[kind], sources }];
  });
}

export function isOpen(phases: readonly PhaseRow[], source: LootSource, when: Date): boolean {
  return hasOpened(phases, source.opens, when);
}

/**
 * A boss's own name, honest about the gap: contract 10.4 lets a boss neither database
 * names come through with an empty `name` (`data/pipeline/loot/sources.py`'s own header,
 * "a boss the fork does not name is emitted with an empty name"). The picker used to fall
 * back to the raw `<source id>:<npc id>` key in that case (dps D35: a
 * `dungeon:blackrock-spire:175245` row where a boss name belongs) -- a key is not a name a
 * player should ever read. This names the one thing both databases DO agree on instead --
 * which zone it is under -- rather than guessing at what kind of thing an unresolved id is:
 * the fork carries no game-object table this pipeline reads, so "Chest" or any other kind
 * would be invented, not resolved.
 */
export function bossName(source: LootSource, boss: LootBoss): string {
  return boss.name === '' ? poolQualityCopy.unnamedSourceIn(source.name) : boss.name;
}

/**
 * A source or boss id as words -- "Molten Core", "Ragnaros" -- for `Candidate.SourceName`
 * (contract 10.1 A6). The page fills it once, when it builds a drops request; every later
 * read is off the result, never a second join.
 */
export function sourceNameOf(file: LootFile, id: string): string {
  for (const source of file.sources) {
    if (source.id === id) return source.name;
    for (const boss of source.bosses ?? []) if (boss.id === id) return bossName(source, boss);
  }
  return '';
}

/**
 * `source.name`, with its reputation standing appended when it has one (dps D30: the
 * source picker listed the same faction four times, once per standing tier -- "Brood of
 * Nozdormu" x4 -- with nothing on the row to tell them apart, though `standing` already
 * travels on every `rep` source). Every other kind is one row per source already and reads
 * exactly as named.
 */
export function sourceLabel(source: LootSource): string {
  if (source.kind !== 'rep' || source.standing === undefined) return source.name;
  const standingLabel = poolQualityCopy.standingLabel[source.standing] ?? source.standing;
  return poolQualityCopy.sourceWithStanding(source.name, standingLabel);
}

/**
 * Crafted sources split into the character's own professions and the rest. Nothing records
 * a character's professions today (`CharacterSpec.professions` is optional and every source
 * leaves it unset), so an absent list puts everything in `other` and the page says why
 * rather than claiming the character has none.
 */
export function professionSplit(
  crafted: readonly LootSource[],
  professions: readonly string[] | undefined,
): { mine: LootSource[]; other: LootSource[] } {
  const mine = new Set(professions ?? []);
  return {
    mine: crafted.filter((source) => source.profession !== undefined && mine.has(source.profession)),
    other: crafted.filter((source) => source.profession === undefined || !mine.has(source.profession)),
  };
}
