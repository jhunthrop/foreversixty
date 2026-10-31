// web/src/lib/guides/gear-extras.ts
// The guide's two small tables under "Enchants and consumables".
//
// Enchants: for each slot, the enchant in the build's `enchants.json` that adds the most of the
// stats this spec's own level 60 weights value (stat-weighted, flat stats only; an enchant that
// grants crit, hit or haste is not scored because those weights are per rating point and the
// enchant data does not say which unit it is in). Enchants of a later phase and class-locked
// ones are left out.
//
// Consumables: the ones the BiS ranker's raid-ready preset gives this spec, with the icon from
// the build's `simbuffs.json` and a short "what it gives" line from CONSUMABLE_GIVES, which a
// test holds to the preset: a consumable the preset adds without a line here fails the test.
import { existsSync } from 'node:fs';
import { bandEntry } from '../bis/load';
import { SLOT_DISPLAY_LABELS } from '../bis/slot-display-labels';
import type { Slot } from '../planner/types';
import { buildFile, readBuildJson } from './build-files';
import { factionForFirstRace, rankedBisFile } from './class-dps';
import type { BisBand, BisStatWeight } from '../bis/types';

export interface EnchantRecord {
  id: number;
  name: string;
  icon: string;
  slots: string[];
  item_types: string[];
  classes: string[];
  stats: Record<string, number>;
  phase: number;
}

export interface EnchantRow {
  slot: string;
  slotLabel: string;
  name: string;
  icon: string | undefined;
  /** "+30 spell power". */
  gives: string;
}

export type ConsumableKind = 'flask' | 'elixir' | 'other';

export interface ConsumableRow {
  id: string;
  /** A flask takes both elixir slots; an elixir is what a player without a flask uses. */
  kind: ConsumableKind;
  name: string;
  icon: string | undefined;
  gives: string;
}

const LATEST_PHASE = 1;

/** Enchant stat ids the weights call something else. */
const STAT_ALIASES: Readonly<Record<string, string>> = {
  spell_damage: 'spell_power',
  healing: 'healing_power',
  bonus_armor: 'armor',
};

/** Rating-family and haste stats: not scored (see the header). */
const UNSCORED = new Set(['crit', 'hit', 'dodge', 'parry', 'block', 'defense', 'melee_haste', 'spell_haste']);

/** An icon shared by this many enchants is the data's stand-in, not the enchant's own. */
const STAND_IN_ICON_USES = 4;

export function consumableKind(id: string): ConsumableKind {
  if (id.startsWith('flask_')) return 'flask';
  if (id.includes('elixir')) return 'elixir';
  return 'other';
}

const STAT_WORDS: Readonly<Record<string, string>> = {
  spell_damage: 'spell power',
  healing: 'healing power',
  bonus_armor: 'armor',
  mp5: 'mana per 5 sec',
  attack_power: 'attack power',
  ranged_attack_power: 'ranged attack power',
  frost_power: 'frost damage',
  fire_power: 'fire damage',
  shadow_power: 'shadow damage',
  nature_power: 'nature damage',
  arcane_power: 'arcane damage',
};

function statWord(stat: string): string {
  return STAT_WORDS[stat] ?? stat.replace(/_/g, ' ');
}

/** "+7 intellect", "+4 strength, agility, stamina, intellect and spirit" collapses to the first
 *  four stats' worth: an all-stats enchant reads "+4 to all stats". */
export function enchantGives(stats: Readonly<Record<string, number>>): string {
  const entries = Object.entries(stats).filter(([stat, value]) => value > 0 && !UNSCORED.has(stat));
  if (entries.length >= 5 && new Set(entries.map(([, value]) => value)).size === 1) {
    return `+${entries[0]![1]} to all stats`;
  }
  return entries.map(([stat, value]) => `+${value} ${statWord(stat)}`).join(', ');
}

export function enchantScore(
  stats: Readonly<Record<string, number>>,
  weights: ReadonlyMap<string, number>,
): number {
  let score = 0;
  for (const [stat, value] of Object.entries(stats)) {
    if (UNSCORED.has(stat)) continue;
    score += value * (weights.get(STAT_ALIASES[stat] ?? stat) ?? 0);
  }
  return score;
}

export interface EnchantContext {
  classSlug: string;
  /** The level 60 main hand is two-handed. */
  twoHanded: boolean;
  /** The level 60 off hand is a shield. */
  shield: boolean;
  /** The level 60 off hand holds an item. */
  hasOffHand: boolean;
}

/** The order the table lists slots in. */
const SLOT_ORDER: readonly string[] = [
  'head',
  'shoulder',
  'back',
  'chest',
  'wrist',
  'hands',
  'legs',
  'feet',
  'main_hand',
  'off_hand',
];

export function pickEnchants(
  records: readonly EnchantRecord[],
  weights: readonly BisStatWeight[],
  context: EnchantContext,
): EnchantRow[] {
  const byStat = new Map(weights.filter((w) => !(w.insignificant ?? false)).map((w) => [w.stat, w.weight]));
  const iconUses = new Map<string, number>();
  for (const record of records) iconUses.set(record.icon, (iconUses.get(record.icon) ?? 0) + 1);
  const best = new Map<string, { score: number; record: EnchantRecord }>();
  for (const record of records) {
    if (record.phase > LATEST_PHASE || record.item_types.includes('kit')) continue;
    if (record.classes.length > 0 && !record.classes.includes(context.classSlug)) continue;
    const types = record.item_types;
    const forWeapon = types.includes('normal') || (types.includes('two_hand') && context.twoHanded);
    const forShield = types.includes('shield') && context.shield;
    const score = enchantScore(record.stats, byStat);
    if (score <= 0) continue;
    for (const slot of record.slots) {
      if (slot === 'main_hand' && !forWeapon) continue;
      if (slot === 'off_hand') {
        if (!context.hasOffHand || !(forShield || (types.includes('normal') && !context.shield))) continue;
      }
      if (slot !== 'main_hand' && slot !== 'off_hand' && !types.includes('normal')) continue;
      const current = best.get(slot);
      if (current === undefined || score > current.score) best.set(slot, { score, record });
    }
  }
  return SLOT_ORDER.filter((slot) => best.has(slot)).map((slot) => {
    const { record } = best.get(slot)!;
    return {
      slot,
      slotLabel: slot === 'main_hand' ? 'Weapon' : (SLOT_DISPLAY_LABELS[slot as Slot] ?? slot),
      name: record.name,
      icon: (iconUses.get(record.icon) ?? 0) >= STAND_IN_ICON_USES ? undefined : record.icon,
      gives: enchantGives(record.stats),
    };
  });
}

/** What each consumable the raid-ready preset uses gives. A new id needs a line here. */
export const CONSUMABLE_GIVES: Readonly<Record<string, string>> = {
  alcohol_rumsey_rum_black_label: 'Stamina',
  conjured_demonic_rune: 'Mana, used on cooldown',
  elixir_of_firepower: 'Fire spell damage',
  elixir_of_fortitude: 'Health',
  elixir_of_frost_power: 'Frost spell damage',
  elixir_of_greater_defense: 'Armor',
  elixir_of_sages: 'Spirit and spell critical strike',
  elixir_of_shadow_power: 'Shadow spell damage',
  elixir_of_the_mongoose: 'Agility and critical strike',
  flask_of_distilled_wisdom: 'A large pool of mana for the whole fight',
  flask_of_supreme_power: 'Spell power',
  flask_of_the_titans: '1200 health for two hours',
  food_grilled_squid: 'Agility',
  food_runn_tum_tuber_surprise: 'Intellect (well fed)',
  food_sages_tea: 'Healing power',
  food_tender_wolf_steak: 'Stamina',
  greater_arcane_elixir: 'Spell power',
  greater_clerics_elixir: 'Healing power',
  greater_mageblood_elixir: 'Mana per 5 sec',
  greater_stoneshield_potion: 'Armor for two minutes',
  juju_might: 'Attack power',
  juju_power: 'Strength',
  conjured_rogue_thistle_tea: 'Energy, used when it runs low',
  'main_hand_imbue:brilliant_mana_oil': 'Mana per 5 sec and healing power, on the main hand',
  'main_hand_imbue:brilliant_wizard_oil': 'Spell power, on the main hand',
  'main_hand_imbue:elemental_sharpening_stone': 'Critical strike, on the main hand',
  major_mana_potion: 'Mana, used on cooldown',
  mighty_rage_potion: 'Rage',
  'off_hand_imbue:dense_sharpening_stone': 'Weapon damage, on the off hand',
  r_o_i_d_s: 'Strength',
};

export interface SimBuffEntry {
  name: string;
  icon?: string;
}

export function consumableRows(
  consumes: readonly { id: string; label: string }[],
  simBuffs: Readonly<Record<string, SimBuffEntry>>,
  iconExists: (icon: string) => boolean,
): ConsumableRow[] {
  const seen = new Set<string>();
  return consumes.flatMap((consume): ConsumableRow[] => {
    const gives = CONSUMABLE_GIVES[consume.id];
    // A preset can name one consumable twice (a class potion and the role's): one row.
    if (gives === undefined || seen.has(consume.id)) return [];
    seen.add(consume.id);
    const icon = (simBuffs[consume.id] ?? simBuffs[consume.id.split(':').pop() ?? ''])?.icon;
    return [
      {
        id: consume.id,
        kind: consumableKind(consume.id),
        name: consume.label,
        icon: icon !== undefined && iconExists(icon) ? icon : undefined,
        gives,
      },
    ];
  });
}

export interface GearExtras {
  enchants: EnchantRow[];
  consumables: ConsumableRow[];
}

/** Both tables for a spec, from its level 60 band; `undefined` when it has no ranked file. */
export function loadGearExtras(
  build: string,
  classSlug: string,
  spec: string,
  recommendedRaces: readonly string[],
  hands: { twoHanded?: boolean; shield?: boolean; hasOffHand?: boolean },
): GearExtras | undefined {
  const file = rankedBisFile(build, spec);
  const faction = factionForFirstRace(classSlug, recommendedRaces);
  if (file === undefined || faction === undefined) return undefined;
  const band: BisBand | undefined = bandEntry(file, 60, faction);
  if (band === undefined) return undefined;
  const records = readBuildJson<EnchantRecord[]>(build, 'enchants.json') ?? [];
  const simBuffs =
    readBuildJson<{ entries: Record<string, SimBuffEntry> }>(build, 'simbuffs.json')?.entries ?? {};
  const raid = file.presets?.raid;
  const consumes = band.preset === 'raid' && raid !== undefined ? raid.consumes : [];
  return {
    enchants: pickEnchants(records, band.weights, {
      classSlug,
      twoHanded: hands.twoHanded ?? false,
      shield: hands.shield ?? false,
      hasOffHand: hands.hasOffHand ?? false,
    }),
    consumables: consumableRows(
      consumes.map((consume) => ({ id: String(consume.id), label: consume.label })),
      simBuffs,
      (icon) => existsSync(buildFile(build, 'icons', `${icon}.webp`)),
    ),
  };
}
