// web/src/lib/guild/loot-icon.ts
// An awarded item's icon and its tooltip of last resort (attendance spec 3.4): the icon the API sends
// with the award (the Loot tab's loot table, the client's own art), or none, in which case the row
// draws the /bis neutral tile with the name's letters, never the client's red question mark; and,
// for an item no class file carries (a quest item such as Head of Onyxia), a tooltip model built
// from the award.
import type { GuildKeyItem } from './api';
import type { RaidLootAward } from './attendance-contract';
import type { ItemTooltipModel } from '../items/tooltip';

/** The award's icon file name, or null when the loot table has none. */
export function lootIconOf(award: Pick<RaidLootAward, 'icon'>): string | null {
  return award.icon === '' ? null : award.icon;
}

/** Where the item stands against its class's item file: still loading, a row there, or none. */
export type ClassFileRow = 'pending' | 'found' | 'none';

/**
 * The icon to request, or null for the neutral tile. The build ships an item's icon only when a
 * class item file carries the item (data/pipeline/icons.py `_referenced_names`), so a drop no class
 * file has (the bag, a quest item, a drop with no class to read) would only 404 (guild round 3:
 * `inv_misc_bag_22`); it draws the tile instead. Nothing is requested while the file loads.
 */
export function shippedLootIcon(award: Pick<RaidLootAward, 'icon'>, row: ClassFileRow): string | null {
  return row === 'found' ? lootIconOf(award) : null;
}

/** The two letters the neutral tile shows when there is no icon to draw. */
export const lootInitials = (name: string): string => name.slice(0, 2);

/** A tooltip of an item's name, quality and icon, and nothing else. */
function bareTooltipModel(id: number, name: string, quality: number, icon: string): ItemTooltipModel {
  return {
    id,
    name,
    quality,
    icon,
    slotLabel: '',
    typeLabel: undefined,
    bindLabel: null,
    itemLevel: 0,
    requiredLevel: 0,
    armor: null,
    weapon: null,
    stats: [],
    effectText: null,
    setName: null,
    sourceLines: [],
    unique: false,
    clientUnconfirmed: false,
  };
}

/** The facts the award carries, and nothing it does not: no item level, slot or stats. */
export function awardTooltipModel(award: RaidLootAward): ItemTooltipModel {
  return bareTooltipModel(award.item_id, award.item_name, award.quality, award.icon);
}

/** A run's key: the API names it and its quality, and no class file carries it. */
export function keyItemTooltipModel(item: GuildKeyItem): ItemTooltipModel {
  return bareTooltipModel(item.id, item.name, item.quality, '');
}
