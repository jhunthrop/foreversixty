// web/src/lib/items/tooltip.ts
// The item tooltip's pure model: every field ItemTooltip.svelte renders, computed once from
// an items/<class>.json row plus the build's loot and set files -- the same split TalentCell
// keeps between "what to show" and "how to show it" (tenet 2: an item is never just a name,
// so this exists to make the real item -- stats, level, source -- a plain data structure a
// component can render without reaching back into three different files itself).
//
// Client-exact where the published row says enough, and silent where it does not:
//
//   * the slot line reads as the client's: a weapon's inventory type on the left ("One-Hand",
//     "Main Hand", "Two-Hand"), its weapon type on the right ("Sword", "Mace") from the row's
//     class and subclass; a shield and an armour piece's material ("Plate") on the right from
//     `armor` and `armor_type`. Nothing is guessed from a name or an icon.
//   * `bindLabel` is the client's binding line ("Binds when picked up") from `bonding`, only
//     when the row states it.
//   * a flat stat prints "+<value> <label>"; a rating-family stat an on-equip spell states as a
//     literal percent (`percent_stats`) prints the client's own sentence ("Equip: Improves
//     your chance to hit by 2%."), after the flat stats, as the client orders them.
import { itemSourceRows, sourceRowLine } from './item-sources';
import type { Item, ItemSet } from '../planner/types';
import { SLOT_LABELS, STAT_KEYS, STAT_LABELS, type StatKey } from '../planner/types';
import { RELIC_SLOT_LABEL } from '../bis/slot-display-labels';
import type { LootFile } from '../sim/loot';
import { humanise } from '../sim/humanise';
import { displayDescription } from '../tooltips/game-tip';

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

/** The client's inventory type for a relic (libram, totem, idol): its tooltip reads "Relic". */
const RELIC_INVENTORY_TYPE = 28;

/** A relic's kind, on the slot line's right, by its armor subclass. */
const RELIC_KIND_LABELS: Readonly<Record<number, string>> = { 7: 'Libram', 8: 'Idol', 9: 'Totem' };

function isRelicItem(item: Item): boolean {
  return item.inventory_type === RELIC_INVENTORY_TYPE;
}

/** The client's weapon class (ItemClass 2). */
const WEAPON_CLASS = 2;

/**
 * The slot line's left for a weapon, shield or held item, by inventory type: GlobalStrings'
 * INVTYPE_WEAPON "One-Hand", INVTYPE_2HWEAPON "Two-Hand", INVTYPE_WEAPONMAINHAND "Main Hand",
 * INVTYPE_WEAPONOFFHAND and INVTYPE_SHIELD "Off Hand", INVTYPE_HOLDABLE "Held In Off-hand",
 * INVTYPE_RANGED and INVTYPE_RANGEDRIGHT "Ranged", INVTYPE_THROWN "Thrown" (1.60.1.70790). The
 * client prints where the item can go, never the slot a character happens to wear it in.
 */
/** INVTYPE_2HWEAPON: the slot line of anything that takes both hands. */
export const TWO_HAND_LABEL = 'Two-Hand';

const INVENTORY_TYPE_LABELS: Readonly<Record<number, string>> = {
  13: 'One-Hand',
  14: 'Off Hand',
  15: 'Ranged',
  17: TWO_HAND_LABEL,
  21: 'Main Hand',
  22: 'Off Hand',
  23: 'Held In Off-hand',
  25: 'Thrown',
  26: 'Ranged',
};

/** The slot line's right for a weapon: ItemSubClass.DisplayName_lang for class 2 (1.60.1.70790). */
const WEAPON_SUBCLASS_LABELS: Readonly<Record<number, string>> = {
  0: 'Axe',
  1: 'Axe',
  2: 'Bow',
  3: 'Gun',
  4: 'Mace',
  5: 'Mace',
  6: 'Polearm',
  7: 'Sword',
  8: 'Sword',
  10: 'Staff',
  13: 'Fist Weapon',
  14: 'Miscellaneous',
  15: 'Dagger',
  16: 'Thrown',
  18: 'Crossbow',
  19: 'Wand',
  20: 'Fishing Pole',
};

function slotLineLeft(item: Item): string {
  if (isRelicItem(item)) return RELIC_SLOT_LABEL;
  const byType = INVENTORY_TYPE_LABELS[item.inventory_type ?? -1];
  if (byType !== undefined) return byType;
  return item.two_hand === true ? TWO_HAND_LABEL : slotLabelFor(item.slot);
}

function typeLabelFor(item: Item): string | undefined {
  if (isRelicItem(item)) return RELIC_KIND_LABELS[item.subclass_id ?? -1];
  if (item.class_id === WEAPON_CLASS) return WEAPON_SUBCLASS_LABELS[item.subclass_id ?? -1];
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

/** The stats the Classic client prints only as a green "Equip:" sentence, whatever states
 *  them (an item stat or an on-equip spell): "Equip: Increases damage and healing done by
 *  magical spells and effects by up to 14.", "Equip: Restores 4 mana per 5 sec.", never a white
 *  "+14" line. The apps' ItemTooltipModel keeps the same list. */
const EQUIP_STAT_LINES: Partial<Record<StatKey, (value: number) => string>> = {
  spell_power: (value) =>
    `Equip: Increases damage and healing done by magical spells and effects by up to ${value}.`,
  healing: (value) => `Equip: Increases healing done by spells and effects by up to ${value}.`,
  attack_power: (value) => `Equip: +${value} Attack Power.`,
  ranged_attack_power: (value) => `Equip: +${value} ranged Attack Power.`,
  mp5: (value) => `Equip: Restores ${value} mana per 5 sec.`,
  spell_penetration: (value) => `Equip: Decreases the magical resistances of your spell targets by ${value}.`,
  fire_power: (value) => `Equip: Increases damage done by Fire spells and effects by up to ${value}.`,
  frost_power: (value) => `Equip: Increases damage done by Frost spells and effects by up to ${value}.`,
  shadow_power: (value) => `Equip: Increases damage done by Shadow spells and effects by up to ${value}.`,
  arcane_power: (value) => `Equip: Increases damage done by Arcane spells and effects by up to ${value}.`,
  nature_power: (value) => `Equip: Increases damage done by Nature spells and effects by up to ${value}.`,
  holy_power: (value) => `Equip: Increases damage done by Holy spells and effects by up to ${value}.`,
  block_value: (value) => `Equip: Increases the block value of your shield by ${value}.`,
};

/** The client's casing for a white stat line ("+10 Shadow Resistance"); the site's own labels
 *  (STAT_LABELS) are sentence case for its panels. */
const CLIENT_STAT_LABELS: Partial<Record<StatKey, string>> = {
  fire_res: 'Fire Resistance',
  frost_res: 'Frost Resistance',
  nature_res: 'Nature Resistance',
  shadow_res: 'Shadow Resistance',
  arcane_res: 'Arcane Resistance',
};

/**
 * Each Equip stat as the amount its sentence states. A single "Increase Spell Dam" spell gives
 * the same spell power and healing and is the one "damage and healing" sentence, so the healing
 * sentence states only the healing beyond the spell power.
 */
function equipShares(stats: Item['stats']): Partial<Record<StatKey, number>> {
  const spellPower = stats.spell_power ?? 0;
  return { ...stats, healing: Math.max(0, (stats.healing ?? 0) - spellPower) };
}

/** Feral attack power is a druid's, in forms only: the client states it as one sentence and never
 *  as a white "+N" line. Most items carry that sentence in their own effect text already, so the
 *  tooltip adds it only when the effect text does not. */
const FERAL_FORMS = 'in Cat, Bear, and Dire Bear forms only';
const feralLine = (value: number): string => `Equip: +${value} Attack Power ${FERAL_FORMS}.`;

function feralLineFor(item: Item): string[] {
  const value = item.stats.feral_attack_power ?? 0;
  return value > 0 && !(item.effect_text ?? '').includes(FERAL_FORMS) ? [feralLine(value)] : [];
}

/** "+10 Strength" / "-15 Parry" (the client's sign: never "+-15") for each white stat, then each
 *  percent stat and each Equip stat in the client's "Equip:" words, in STAT_KEYS order within
 *  each group. */
function statLines(item: Item): string[] {
  const keys = STAT_KEYS.filter(
    (key) => key !== 'armor' && key !== 'feral_attack_power' && !!item.stats[key],
  );
  const flat = keys
    .filter((key) => percentLineFor(item, key) === null && EQUIP_STAT_LINES[key] === undefined)
    .map((key) => {
      const value = item.stats[key] as number;
      return `${value < 0 ? '-' : '+'}${Math.abs(value)} ${CLIENT_STAT_LABELS[key] ?? STAT_LABELS[key]}`;
    });
  const percent = keys
    .map((key) => percentLineFor(item, key))
    .filter((line): line is string => line !== null);
  const shares = equipShares(item.stats);
  const equip = STAT_KEYS.flatMap((key) => {
    const value = shares[key] ?? 0;
    const line = EQUIP_STAT_LINES[key];
    return value > 0 && line ? [line(value)] : [];
  });
  return [...flat, ...percent, ...equip, ...feralLineFor(item)];
}

/** The item's equip or use line as the client prints it, with no raw token left in it. */
function effectTextFor(item: Item): string | null {
  const text = displayDescription(item.effect_text ?? '')
    .replace(/\s+/g, ' ')
    .trim();
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
    slotLabel: slotLineLeft(row),
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
