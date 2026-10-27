// web/src/lib/guild/loot-icon.ts
// An awarded item's icon and its tooltip of last resort (attendance spec 3.4): the icon the API sends
// with the award (the Loot tab's loot table), else the client's question mark; and, for an item no
// class file carries (a quest item such as Head of Onyxia), a tooltip model built from the award.
import type { RaidLootAward } from './attendance-contract';
import type { ItemTooltipModel } from '../items/tooltip';

/** The client's own icon for an item it has no art for. */
export const QUESTION_MARK_ICON = 'inv_misc_questionmark';

export function lootIconOf(award: Pick<RaidLootAward, 'icon'>): string {
  return award.icon === '' ? QUESTION_MARK_ICON : award.icon;
}

/** The facts the award carries, and nothing it does not: no item level, slot or stats. */
export function awardTooltipModel(award: RaidLootAward): ItemTooltipModel {
  return {
    id: award.item_id,
    name: award.item_name,
    quality: award.quality,
    icon: lootIconOf(award),
    slotLabel: '',
    typeLabel: undefined,
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
