// web/src/lib/bis/source-cell.ts
// The BiS row's source cell, resolved to the real thing the tenets ask for -- a quest name
// and level, an instance and boss, a profession, an npc, a faction and standing, a zone --
// rather than the generic per-kind badge (`load.ts`'s `sourceBadgeLabel`, still used for the
// small badge word itself) or a bare dash. Reads `loot.json`'s `sources` (this build's own
// copy of `sim/loot.ts`'s `LootFile`, loaded off disk at build time the way `load.ts` reads
// every other `/bis` input) and its `quests` map (`LootQuestsFile`, `./types.ts`), never a
// fetch -- this runs in Astro's Node build step, the same boundary `load.ts`'s header
// documents for the rest of this page's data.
import { bossName, itemsOfSource, type LootFile, type LootSource } from '../sim/loot';
import { bisCopy } from './copy';
import type { BisSlot, Faction, LootQuestOption, LootQuestsFile } from './types';

/**
 * Whether a slot row carries an actual pick. The pipeline writes a slot the ranking found no
 * source for as `{ slot, verified: false }` alone (see `BisSlot`'s own doc comment) even
 * though the type says `item_id` is required -- this is the one place that trust gap is
 * checked, so every caller below it (and the page) can trust `item_id` once this is true.
 */
export function hasKnownSource(row: BisSlot): boolean {
  return typeof row.item_id === 'number';
}

export interface QuestSourceCell {
  kind: 'quest';
  questName: string;
  faction: Faction;
  level: number;
}

export interface PlaceSourceCell {
  /** `dungeon` or `raid` -- the two kinds `loot.json` gives a boss list. */
  kind: 'dungeon' | 'raid';
  instance: string;
  /** Undefined when the item is the instance's own trash or unbossed loot list, not any
   *  named boss's own drop. */
  boss?: string;
  /** Percent drop chance (0-100), from cmangos/classic-db (src-classicdb lane,
   *  2026-09-29). Undefined when no source contributing to this bucket names one -- the
   *  fork database and wowhead's scrape never do. */
  dropChance?: number;
}

export interface CraftedSourceCell {
  kind: 'crafted';
  profession: string;
}

export interface VendorSourceCell {
  kind: 'vendor';
  npc: string;
}

export interface RepSourceCell {
  kind: 'rep';
  faction: string;
  standing?: string;
}

export interface ZoneSourceCell {
  /** `zone`, `world` or `pvp` -- every other kind `loot.json` names as one flat place. */
  kind: 'zone' | 'world' | 'pvp';
  place: string;
  /** See `PlaceSourceCell.dropChance`'s own doc. */
  dropChance?: number;
}

export interface WorldDropSourceCell {
  kind: 'world_drop';
  /** See `LootSource.level_min`/`level_max`'s own doc -- absent when the pinned dump names
   *  no level range for this pool at all. */
  levelMin?: number;
  levelMax?: number;
}

/** A source_kind `loot.json` has nothing for this item under (a data gap, not a defect in
 *  this resolver -- `sim/loot.ts`'s own header names 1,809 sourced ids the 1.60 client does
 *  not carry), or a kind this resolver does not special-case. Carries the pipeline's own
 *  `source` string so the row still says something true rather than nothing. */
export interface FallbackSourceCell {
  kind: 'unknown';
  label: string;
}

export type SourceCell =
  | QuestSourceCell
  | PlaceSourceCell
  | CraftedSourceCell
  | VendorSourceCell
  | RepSourceCell
  | ZoneSourceCell
  | WorldDropSourceCell
  | FallbackSourceCell;

function findSource(sources: readonly LootSource[], kind: string, itemId: number): LootSource | undefined {
  return sources.find((source) => source.kind === kind && itemsOfSource(source).includes(itemId));
}

/**
 * Among every `kind`-matching source naming `itemId`, the one whose own drop chance for it
 * is highest -- `findSource`'s plain "first array match" is not the same choice
 * `sim/cmd/leveling-bis/band.go`'s `bestBoss` makes when an item drops from bosses in more
 * than one dungeon/raid zone. Mirrors that choice (wowhead-world-drops lane, 2026-09-29) so
 * the row never shows an arbitrary trash-mob source when a higher-chance boss source exists
 * elsewhere in the array (tenet 7). A source with no chance data (`findChance` undefined,
 * treated as -1) never beats one that states a chance, and ties keep the first array match,
 * same fallback `findSource` always had.
 */
function findBestChanceSource(
  sources: readonly LootSource[],
  kind: string,
  itemId: number,
): LootSource | undefined {
  let best: LootSource | undefined;
  let bestChance = -1;
  for (const candidate of sources) {
    if (candidate.kind !== kind || !itemsOfSource(candidate).includes(itemId)) continue;
    const chance = findChance(candidate, itemId) ?? -1;
    if (best === undefined || chance > bestChance) {
      best = candidate;
      bestChance = chance;
    }
  }
  return best;
}

function findBossEntry(source: LootSource, itemId: number) {
  return (source.bosses ?? []).find((entry) => entry.items.includes(itemId));
}

function findBoss(source: LootSource, itemId: number): string | undefined {
  const boss = findBossEntry(source, itemId);
  return boss === undefined ? undefined : bossName(source, boss);
}

/** `source`'s (or, for a boss drop, that boss's own) `item_chances` entry for `itemId` --
 *  `PlaceSourceCell.dropChance`/`ZoneSourceCell.dropChance`'s own doc. */
function findChance(source: LootSource, itemId: number): number | undefined {
  const boss = findBossEntry(source, itemId);
  const chances = boss?.item_chances ?? source.item_chances;
  return chances?.[String(itemId)];
}

/** The quest matching this band's own faction when `loot.json` lists more than one (a
 *  faction-specific quest chain for the same item); falls back to the one entry a neutral
 *  quest has rather than showing nothing for a faction it never distinguished. */
function questFor(options: readonly LootQuestOption[], faction: Faction): LootQuestOption | undefined {
  return options.find((option) => option.faction === faction) ?? options[0];
}

/**
 * The row's real source, honestly falling back to `slot.source_kind`'s generic label
 * (`fallbackLabel`) when `loot.json` has nothing keyed to this exact item -- never a guess,
 * and never a throw for an item this build's loot table does not carry.
 */
export function resolveSourceCell(
  slot: BisSlot,
  faction: Faction,
  loot: LootFile & Partial<LootQuestsFile>,
  fallbackLabel: string,
): SourceCell {
  const itemId = slot.item_id;
  const fallback: FallbackSourceCell = { kind: 'unknown', label: fallbackLabel };

  if (slot.source_kind === 'quest') {
    const options = loot.quests?.[String(itemId)];
    const option = options === undefined ? undefined : questFor(options, faction);
    // The badge reads the band's own faction, not the quest's -- a `'both'` (neutral) quest
    // still shows "Alliance quest" on the Alliance panel and "Horde quest" on Horde's,
    // matching `sourceBadgeLabel`'s existing rule for the generic per-kind badge.
    return option === undefined
      ? fallback
      : { kind: 'quest', questName: option.name, faction, level: option.level };
  }

  if (slot.source_kind === 'dungeon' || slot.source_kind === 'raid') {
    const source = findBestChanceSource(loot.sources, slot.source_kind, itemId);
    return source === undefined
      ? fallback
      : {
          kind: slot.source_kind,
          instance: source.name,
          boss: findBoss(source, itemId),
          dropChance: findChance(source, itemId),
        };
  }

  if (slot.source_kind === 'crafted') {
    const source = findSource(loot.sources, 'crafted', itemId);
    return source === undefined ? fallback : { kind: 'crafted', profession: source.name };
  }

  if (slot.source_kind === 'vendor') {
    const source = findSource(loot.sources, 'vendor', itemId);
    return source === undefined ? fallback : { kind: 'vendor', npc: source.name };
  }

  if (slot.source_kind === 'rep') {
    const source = findSource(loot.sources, 'rep', itemId);
    return source === undefined ? fallback : { kind: 'rep', faction: source.name, standing: source.standing };
  }

  if (slot.source_kind === 'zone' || slot.source_kind === 'world' || slot.source_kind === 'pvp') {
    const source = findSource(loot.sources, slot.source_kind, itemId);
    return source === undefined
      ? fallback
      : { kind: slot.source_kind, place: source.name, dropChance: findChance(source, itemId) };
  }

  if (slot.source_kind === 'world_drop') {
    const source = findSource(loot.sources, 'world_drop', itemId);
    return source === undefined
      ? fallback
      : { kind: 'world_drop', levelMin: source.level_min, levelMax: source.level_max };
  }

  return fallback;
}

/** The specific line a source cell shows beneath the row's generic kind badge (`load.ts`'s
 *  `sourceBadgeLabel`): "Quest: Pulsating Crystalline Shard · Level 50", "Molten Core ·
 *  Ragnaros", "Crafted: Blacksmithing", and so on -- one line per kind, the real thing tenet
 *  2 asks a source cell to be. */
export function describeSourceCell(cell: SourceCell): string {
  switch (cell.kind) {
    case 'quest':
      return cell.level > 0
        ? `${bisCopy.questSourceLabel(cell.questName)} · ${bisCopy.questLevelLabel(cell.level)}`
        : bisCopy.questSourceLabel(cell.questName);
    case 'dungeon':
    case 'raid':
      return cell.dropChance === undefined
        ? bisCopy.dungeonSourceLabel(cell.instance, cell.boss)
        : bisCopy.dropChanceLabel(cell.dropChance, cell.boss ?? `${cell.instance} trash`);
    case 'crafted':
      return bisCopy.craftedSourceLabel(cell.profession);
    case 'vendor':
      return bisCopy.vendorSourceLabel(cell.npc);
    case 'rep':
      return bisCopy.repSourceLabel(cell.faction, cell.standing);
    case 'zone':
    case 'world':
    case 'pvp':
      return cell.dropChance === undefined
        ? bisCopy.placeSourceLabel(cell.place)
        : bisCopy.dropChanceLabel(cell.dropChance, cell.place);
    case 'world_drop':
      return bisCopy.worldDropSourceLabel(cell.levelMin, cell.levelMax);
    case 'unknown':
      return cell.label;
  }
}

export interface ParsedSwapNote {
  runnerUpName: string;
  runnerUpItemId: number;
  higherDps: number;
  lowerDps: number;
}

const SWAP_NOTE_PATTERN =
  /^runner-up (.+) \(id (\d+)\) measured higher: ([\d.]+) vs ([\d.]+) set DPS(?: - swapped in)?$/;

/**
 * `swap_note`'s free text, structured: every unverified slot in every band file this lane
 * has read carries exactly this shape (`runner-up <name> (id <id>) measured higher: <a> vs
 * <b> set DPS[ - swapped in]`), one sentence the nightly's own writer composes. Returns null
 * for a note that does not match rather than throwing -- a future pipeline change to the
 * sentence should degrade to the plain "Unverified" badge, not break the page.
 */
export function parseSwapNote(note: string): ParsedSwapNote | null {
  const match = SWAP_NOTE_PATTERN.exec(note);
  if (match === null) return null;
  return {
    runnerUpName: match[1],
    runnerUpItemId: Number(match[2]),
    higherDps: Number(match[3]),
    lowerDps: Number(match[4]),
  };
}
