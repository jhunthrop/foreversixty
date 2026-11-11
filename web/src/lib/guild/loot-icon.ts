// web/src/lib/guild/loot-icon.ts
// An awarded item's tooltip of last resort (attendance spec 3.4): for an item no class file carries (a
// quest item such as Head of Onyxia), a tooltip model built from the award. Its icon is the award's
// own, drawn by ui/ItemIcon (the dark socket for none, never the client's red question mark).
import type { GuildKeyItem } from './api';
import type { RaidLootAward } from './attendance-contract';
import type { ItemTooltipModel } from '../items/tooltip';

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
