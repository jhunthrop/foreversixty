// web/src/lib/guild/loot-icon.ts
// An awarded item's icon and its tooltip of last resort (attendance spec 3.4): the icon the API sends
// with the award (the Loot tab's loot table, the client's own art), or none, in which case the row
// draws the /bis neutral tile with the name's letters, never the client's red question mark; and,
// for an item no class file carries (a quest item such as Head of Onyxia), a tooltip model built
// from the award.
import type { RaidLootAward } from './attendance-contract';
import type { ItemTooltipModel } from '../items/tooltip';

/** The award's icon file name, or null when the loot table has none. */
export function lootIconOf(award: Pick<RaidLootAward, 'icon'>): string | null {
  return award.icon === '' ? null : award.icon;
}

/** The two letters the neutral tile shows when there is no icon to draw. */
export const lootInitials = (name: string): string => name.slice(0, 2);

/** The facts the award carries, and nothing it does not: no item level, slot or stats. */
export function awardTooltipModel(award: RaidLootAward): ItemTooltipModel {
  return {
    id: award.item_id,
    name: award.item_name,
    quality: award.quality,
    icon: award.icon,
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
