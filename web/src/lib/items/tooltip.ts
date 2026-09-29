// web/src/lib/items/tooltip.ts
// The item tooltip's pure model: every field ItemTooltip.svelte renders, computed once from
// an items/<class>.json row plus the build's loot and set files -- the same split TalentCell
// keeps between "what to show" and "how to show it" (tenet 2: an item is never just a name,
// so this exists to make the real item -- stats, level, source -- a plain data structure a
// component can render without reaching back into three different files itself).
//
// Deliberately NOT client-tooltip-exact on two points, both decided rather than guessed:
//
//   * `typeLabel` names a weapon's handedness ("Two-Handed Weapon") and a shield, because
//     those are exact from fields the published row already carries (`two_hand`, `armor` on
//     an off-hand). It does not name a weapon's subclass (sword, axe, mace) or an armor
//     piece's material (cloth, leather, mail, plate): items/<class>.json carries no
//     class_id/subclass_id (normalize/items.py's own Item model has them; build_class_items
//     never publishes them), and guessing either from a name or an icon string would be
//     exactly the fabrication tenet 5 rules out. `typeLabel` is `undefined` rather than a
//     invented "Armor" for anything else -- a follow-up for the data lane to publish those
//     two ids, not a job for this lane's pure function.
//   * every stat prints as "+<value> <label>", the same convention GearPanel's own totals
//     and the addon's stat weights already use, rather than the client's percent-converted
//     "Equip: Improves your chance to get a critical strike by 1%." wording for crit/hit/
//     dodge/parry/block. The published `stats.crit` etc. are the engine's own Stat values,
//     with no documented, verified per-point-to-percent table in this repo -- inventing one
//     here would put an unverified number on screen, which tenet 5 puts ahead of matching
//     the client's exact prose.
import type { Item, ItemSet } from '../planner/types';
import { SLOT_LABELS, STAT_KEYS, STAT_LABELS } from '../planner/types';
import { bossName, itemsOfBoss, itemsOfSource, sourceLabel, type LootFile } from '../sim/loot';
import { humanise } from '../sim/humanise';

export interface WeaponLine {
  damageRange: string;
  speed: string;
  dps: string;
}

export interface ItemTooltipModel {
  id: number;
  name: string;
  /** Client value: 0 poor … 5 legendary -- ItemTooltip.svelte's own rarityClassFor call. */
  quality: number;
  icon: string;
  slotLabel: string;
  /** See the module note above: present only where it can be exact, absent otherwise. */
  typeLabel: string | undefined;
  itemLevel: number;
  requiredLevel: number;
  /** `null` rather than 0 -- an accessory with no armor value shows no armor row at all,
   *  the same way the client's own tooltip omits it. */
  armor: number | null;
  weapon: WeaponLine | null;
  /** "+18 Strength", "+14 Crit", … in STAT_KEYS' own order (armor excluded -- it has its
   *  own row above). Only the stats the item actually carries. */
  stats: string[];
  /** The weapon's on-hit/on-use/on-equip line the data pipeline already renders as prose
   *  (`data/pipeline/normalize/effects.py`'s `EffectIndex.text`); `null` when the item has
   *  none. */
  effectText: string | null;
  setName: string | null;
  /** "Molten Core — Ragnaros", "Sweet Amber (Alliance)", "Argent Dawn — Exalted", … one
   *  line per place the item is known to come from; empty when loot.json has none. */
  sourceLines: string[];
  unique: boolean;
}

export interface ItemTooltipSources {
  loot: LootFile;
  sets: readonly ItemSet[];
}

const ALIAS_SLOT_LABELS: Record<string, string> = { finger: 'Finger', trinket: 'Trinket' };

function slotLabelFor(slot: string): string {
  if (Object.prototype.hasOwnProperty.call(SLOT_LABELS, slot)) {
    return SLOT_LABELS[slot as keyof typeof SLOT_LABELS];
  }
  return ALIAS_SLOT_LABELS[slot] ?? humanise(slot);
}

const WEAPON_SLOTS: ReadonlySet<string> = new Set(['main_hand', 'off_hand', 'ranged']);

function typeLabelFor(item: Item): string | undefined {
  const isWeapon = WEAPON_SLOTS.has(item.slot) && (item.damage_max ?? 0) > 0;
  if (isWeapon) {
    if (item.two_hand === true) return 'Two-Handed Weapon';
    if (item.slot === 'ranged') return 'Ranged Weapon';
    return 'One-Handed Weapon';
  }
  if (item.slot === 'off_hand' && item.armor > 0) return 'Shield';
  return undefined;
}

function weaponLineFor(item: Item): WeaponLine | null {
  if (item.damage_min === undefined || item.damage_max === undefined || item.speed === undefined) {
    return null;
  }
  if (item.damage_max <= 0) return null;
  return {
    damageRange: `${Math.round(item.damage_min)} - ${Math.round(item.damage_max)} Damage`,
    speed: `Speed ${item.speed.toFixed(2)}`,
    dps: `(${(item.dps ?? 0).toFixed(1)} damage per second)`,
  };
}

function statLines(item: Item): string[] {
  return STAT_KEYS.filter((key) => key !== 'armor' && !!item.stats[key]).map(
    (key) => `+${item.stats[key]} ${STAT_LABELS[key]}`,
  );
}

function effectTextFor(item: Item): string | null {
  const text = (item.effect_text ?? '').replace(/\s+/g, ' ').trim();
  return text === '' ? null : text;
}

function setNameFor(item: Item, sets: readonly ItemSet[]): string | null {
  if (item.set_id === null) return null;
  return sets.find((set) => set.id === item.set_id)?.name ?? null;
}

/**
 * Every place loot.json says `itemId` comes from: a specific boss when one of the source's
 * bosses names it (checked first, so a raid/dungeon item never also prints its zone-level
 * source line for the same drop), else the source itself worded per kind, plus one line per
 * quest that rewards it. `sourceLabel` and `bossName` are loot.ts's own -- reused rather than
 * re-worded here, so a rep source with a standing or an unnamed boss never drifts from how
 * the source picker already renders the identical source.
 */
function sourceLinesFor(itemId: number, loot: LootFile): string[] {
  const lines: string[] = [];
  for (const source of loot.sources) {
    const boss = (source.bosses ?? []).find((candidate) =>
      itemsOfBoss(source, candidate.id).includes(itemId),
    );
    if (boss !== undefined) {
      lines.push(`${source.name} — ${bossName(source, boss)}`);
      continue;
    }
    if (!itemsOfSource(source).includes(itemId)) continue;
    if (source.kind === 'crafted') {
      lines.push(source.profession !== undefined ? `${source.name} (${source.profession})` : source.name);
    } else if (source.kind === 'rep') {
      lines.push(sourceLabel(source));
    } else if (source.kind === 'pvp') {
      lines.push(source.rank !== undefined ? `${source.name}, rank ${source.rank}` : source.name);
    } else {
      lines.push(source.name);
    }
  }
  for (const quest of loot.quests?.[String(itemId)] ?? []) {
    lines.push(`${quest.name} (${humanise(quest.faction)})`);
  }
  return lines;
}

/** The one place every ItemHover/ItemTooltip in the site builds its model -- a runtime
 *  fetch (lookup.ts's fetchItemTooltipModel) and a build-time read (lookup.ts's
 *  readItemTooltipModel) both end here, so an SSR-rendered model and a client-fetched one
 *  are never two different shapes. */
export function itemTooltipModel(row: Item, sources: ItemTooltipSources): ItemTooltipModel {
  return {
    id: row.id,
    name: row.name,
    quality: row.quality,
    icon: row.icon,
    slotLabel: slotLabelFor(row.slot),
    typeLabel: typeLabelFor(row),
    itemLevel: row.item_level,
    requiredLevel: row.required_level,
    armor: row.armor > 0 ? row.armor : null,
    weapon: weaponLineFor(row),
    stats: statLines(row),
    effectText: effectTextFor(row),
    setName: setNameFor(row, sources.sets),
    sourceLines: sourceLinesFor(row.id, sources.loot),
    unique: row.unique,
  };
}
