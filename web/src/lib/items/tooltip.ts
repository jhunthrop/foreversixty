// web/src/lib/items/tooltip.ts
// The item tooltip's pure model: every field ItemTooltip.svelte renders, computed once from
// an items/<class>.json row plus the build's loot and set files -- the same split TalentCell
// keeps between "what to show" and "how to show it" (tenet 2: an item is never just a name,
// so this exists to make the real item -- stats, level, source -- a plain data structure a
// component can render without reaching back into three different files itself).
//
// Client-exact where the published row says enough, and silent where it does not:
//
//   * `typeLabel` names a weapon's handedness ("Two-Handed Weapon") and a shield from
//     `two_hand` and `armor`, and an armour piece's material ("Plate") from the row's
//     `armor_type` when the build publishes it. Nothing is guessed from a name or an icon.
//   * `bindLabel` is the client's binding line ("Binds when picked up") from `bonding`, only
//     when the row states it.
//   * a flat stat prints "+<value> <label>"; a rating-family stat an on-equip spell states as a
//     literal percent (`percent_stats`) prints the client's own sentence ("Equip: Improves
//     your chance to hit by 2%."), after the flat stats, as the client orders them.
import { itemSourceRows, sourceRowLine } from './item-sources';
import type { Item, ItemSet } from '../planner/types';
import { SLOT_LABELS, STAT_KEYS, STAT_LABELS, type StatKey } from '../planner/types';
import type { LootFile } from '../sim/loot';
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
  /** "Binds when picked up", …; null when the row states no binding. */
  bindLabel: string | null;
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
  /** `Item.client_unconfirmed ?? false` -- one field, computed once here, read by every
   *  surface (character-panel row, alternatives row, this tooltip) off the same model,
   *  never three separate lookups back into the item row (spec addendum §B). */
  clientUnconfirmed: boolean;
}

export interface ItemTooltipSources {
  loot: LootFile;
  sets: readonly ItemSet[];
}

/** The source block never shows more than this many lines before collapsing the rest into
 *  "and N more" (owner screenshot 2026-09-29, tooltip-polish brief item 4) -- a raid boss
 *  with a dozen possible drops used to print a dozen lines, dwarfing the rest of the panel. */
const MAX_SOURCE_LINES = 3;

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
  return item.armor_type == null ? undefined : ARMOR_LABELS[item.armor_type];
}

/** `armor_type` -> the client's word on the slot line's right. */
const ARMOR_LABELS: Readonly<Record<NonNullable<Item['armor_type']>, string>> = {
  cloth: 'Cloth',
  leather: 'Leather',
  mail: 'Mail',
  plate: 'Plate',
  shield: 'Shield',
};

/** `bonding` -> the client's binding line. */
const BIND_LABELS: Readonly<Record<NonNullable<Item['bonding']>, string>> = {
  pickup: 'Binds when picked up',
  equip: 'Binds when equipped',
  use: 'Binds when used',
  quest: 'Quest Item',
};

function bindLabelFor(item: Item): string | null {
  return item.bonding == null ? null : BIND_LABELS[item.bonding];
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

/** The client's own sentence for a rating-family stat stated as a percent. */
const PERCENT_STAT_LINES: Partial<Record<StatKey, (value: number) => string>> = {
  crit: (value) => `Equip: Improves your chance to get a critical strike by ${value}%.`,
  hit: (value) => `Equip: Improves your chance to hit by ${value}%.`,
  defense: (value) => `Equip: Increased Defense +${value}.`,
  dodge: (value) => `Equip: Increases your chance to dodge an attack by ${value}%.`,
  parry: (value) => `Equip: Increases your chance to parry an attack by ${value}%.`,
  block: (value) => `Equip: Increases your chance to block attacks with a shield by ${value}%.`,
};

function percentLineFor(item: Item, key: StatKey): string | null {
  const percent = item.percent_stats?.[key];
  const line = PERCENT_STAT_LINES[key];
  return percent && line && percent === item.stats[key] ? line(percent) : null;
}

/** "+10 Strength" / "-15 Parry" (the client's sign: never "+-15"), then each percent stat in
 *  the client's "Equip:" words, in STAT_KEYS order within each group. */
function statLines(item: Item): string[] {
  const keys = STAT_KEYS.filter((key) => key !== 'armor' && !!item.stats[key]);
  const flat = keys
    .filter((key) => percentLineFor(item, key) === null)
    .map((key) => {
      const value = item.stats[key] as number;
      return `${value < 0 ? '-' : '+'}${Math.abs(value)} ${STAT_LABELS[key]}`;
    });
  const percent = keys
    .map((key) => percentLineFor(item, key))
    .filter((line): line is string => line !== null);
  return [...flat, ...percent];
}

function effectTextFor(item: Item): string | null {
  const text = (item.effect_text ?? '').replace(/\s+/g, ' ').trim();
  return text === '' ? null : text;
}

function setNameFor(item: Item, sets: readonly ItemSet[]): string | null {
  if (item.set_id === null) return null;
  return sets.find((set) => set.id === item.set_id)?.name ?? null;
}

/** The tooltip's source lines: `itemSourceRows`' ranked, deduped rows in one-line form, at
 *  most `MAX_SOURCE_LINES` of them, the rest collapsed into one "and N more" line instead of
 *  dwarfing the rest of the panel. */
function sourceLinesFor(itemId: number, loot: LootFile): string[] {
  const lines = itemSourceRows(itemId, loot).map(sourceRowLine);
  if (lines.length <= MAX_SOURCE_LINES) return lines;
  return [...lines.slice(0, MAX_SOURCE_LINES), `and ${lines.length - MAX_SOURCE_LINES} more`];
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
    bindLabel: bindLabelFor(row),
    itemLevel: row.item_level,
    requiredLevel: row.required_level,
    armor: row.armor > 0 ? row.armor : null,
    weapon: weaponLineFor(row),
    stats: statLines(row),
    effectText: effectTextFor(row),
    setName: setNameFor(row, sources.sets),
    sourceLines: sourceLinesFor(row.id, sources.loot),
    unique: row.unique,
    clientUnconfirmed: row.client_unconfirmed ?? false,
  };
}
