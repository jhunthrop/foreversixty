// web/src/lib/items/item-sources.ts
// Every place loot.json says an item comes from, as structured rows: what the place is (a
// boss, a vendor, a quest, a recipe, a world drop), where it is, its kind and the drop chance
// when the data has one. The tooltip's one-line sources (`tooltip.ts`) and the apps' item page
// (`/data/{build}/item-uses/{shard}.json`, Wave 3 spec F12.4) are both read from these rows,
// so the two never word one source two ways.
import { craftedSourceLabel, pvpSourceLabel } from './source-labels';
import {
  bossName,
  itemsOfBoss,
  itemsOfSource,
  sourceLabel,
  type LootFile,
  type LootKind,
  type LootSource,
} from '../sim/loot';
import { humanise } from '../sim/humanise';
import { dropPercentLabel } from './drop-chance';

/** `world` today, `world_drop` once the loot pipeline's own migration lands -- both name the
 *  same thing and collapse to one "World drop" line. */
const WORLD_DROP_KINDS: ReadonlySet<string> = new Set(['world', 'world_drop']);

export interface ItemSourceRow {
  /** loot.json's source kind; `quest` for a quest reward. */
  kind: LootKind;
  /** The thing a player looks for: the boss, the vendor, the quest, the recipe's profession,
   *  the faction and standing, `World drop`. */
  title: string;
  /** Where it is: the instance of a boss drop, the faction side of a quest. Absent when the
   *  data names no place. */
  place?: string;
  /** Percent (0-100) when loot.json records one for this item at this source; absent when
   *  unknown, never 0 for unknown. */
  chance?: number;
  /** The quest's minimum level, quests only; absent when unknown. */
  min_level?: number;
  /** True for a named boss kill or a quest reward -- the row a player goes and does. */
  primary: boolean;
  /** The one-line form the site tooltip prints ("Blackrock Depths — Houndmaster Grebmar"). */
  text: string;
}

function worldDropTitle(source: LootSource): string {
  const { level_min: levelMin, level_max: levelMax } = source;
  if (typeof levelMin === 'number' && typeof levelMax === 'number') {
    return `World drop (BoE) · levels ${levelMin}-${levelMax}`;
  }
  return 'World drop';
}

function pvpTitle(source: LootSource): string {
  return source.rank !== undefined && (source.faction === 'alliance' || source.faction === 'horde')
    ? pvpSourceLabel(source.rank, source.faction)
    : source.name;
}

/** A rank quartermaster's vendor row repeats its pvp source's items verbatim; the pvp row
 *  already names the rank and title, so the bare vendor name is left out. */
function isPvpVendorDuplicate(source: LootSource, itemId: number, loot: LootFile): boolean {
  return (
    source.kind === 'vendor' &&
    loot.sources.some((other) => other.kind === 'pvp' && itemsOfSource(other).includes(itemId))
  );
}

/** One row per boss of the source that drops the item (Spritecaster Cape drops from two
 *  Blackrock Depths bosses); a boss drop never also prints the instance's own line. */
function bossRows(source: LootSource, itemId: number): ItemSourceRow[] {
  return (source.bosses ?? [])
    .filter((boss) => itemsOfBoss(source, boss.id).includes(itemId))
    .map((boss) => {
      const title = bossName(source, boss).trim();
      const chance = boss.item_chances?.[String(itemId)];
      return {
        kind: source.kind,
        title,
        place: source.name,
        ...(chance === undefined ? {} : { chance }),
        primary: true,
        text: `${source.name} — ${title}`,
      };
    });
}

function rowsFromSource(source: LootSource, itemId: number, loot: LootFile): ItemSourceRow[] {
  if (WORLD_DROP_KINDS.has(source.kind)) {
    if (!(source.items ?? []).includes(itemId)) return [];
    const title = worldDropTitle(source);
    return [{ kind: source.kind, title, primary: false, text: title }];
  }
  const bosses = bossRows(source, itemId);
  if (bosses.length > 0) return bosses;
  const row = placeRow(source, itemId, loot);
  return row === null ? [] : [row];
}

function placeRow(source: LootSource, itemId: number, loot: LootFile): ItemSourceRow | null {
  if (!itemsOfSource(source).includes(itemId) || isPvpVendorDuplicate(source, itemId, loot)) return null;
  const chance = source.item_chances?.[String(itemId)];
  const titled = (title: string, withChance: boolean): ItemSourceRow => ({
    kind: source.kind,
    title,
    ...(withChance && chance !== undefined ? { chance } : {}),
    primary: false,
    text: title,
  });
  switch (source.kind) {
    case 'crafted':
      return titled(craftedSourceLabel(source.name, source.profession), false);
    case 'rep':
      return titled(sourceLabel(source), false);
    case 'pvp':
      return titled(pvpTitle(source), false);
    default:
      return titled(source.name, true);
  }
}

function questRows(itemId: number, loot: LootFile): ItemSourceRow[] {
  return (loot.quests?.[String(itemId)] ?? []).map((quest) => {
    const place = humanise(quest.faction);
    return {
      kind: 'quest' as const,
      title: quest.name,
      place,
      ...(quest.min_level > 0 ? { min_level: quest.min_level } : {}),
      primary: true,
      text: `${quest.name} (${place})`,
    };
  });
}

/** The row's line with its chance, the tooltip's own form ("… (20%)"). */
export function sourceRowLine(row: ItemSourceRow): string {
  return row.chance === undefined ? row.text : `${row.text} (${dropPercentLabel(row.chance)})`;
}

/**
 * Every source of `itemId`, ranked: boss kills and quest rewards first, a known drop chance
 * descending within its group, an unknown chance last, ties in loot.json's own order. A row
 * whose line repeats an earlier one (a quest given to both sides, two world pools with the
 * same words) is dropped, so the list never names one place twice.
 */
export function itemSourceRows(itemId: number, loot: LootFile): ItemSourceRow[] {
  const candidates = [
    ...loot.sources.flatMap((source) => rowsFromSource(source, itemId, loot)),
    ...questRows(itemId, loot),
  ];
  const ranked = candidates
    .map((row, order) => ({ row, order }))
    .sort((a, b) => {
      if (a.row.primary !== b.row.primary) return a.row.primary ? -1 : 1;
      if (a.row.chance !== b.row.chance) {
        if (a.row.chance === undefined) return 1;
        if (b.row.chance === undefined) return -1;
        return b.row.chance - a.row.chance;
      }
      return a.order - b.order;
    })
    .map(({ row }) => row);
  const seen = new Set<string>();
  return ranked.filter((row) => {
    const line = sourceRowLine(row);
    if (seen.has(line)) return false;
    seen.add(line);
    return true;
  });
}
