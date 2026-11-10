// web/src/lib/guides/gear-extras.ts
// The guide's two small tables under "Enchants and consumables".
//
// Enchants: for each slot, the enchant in the build's `enchants.json` that adds the most of the
// stats this spec's own level 60 weights value (stat-weighted, flat stats only; an enchant that
// grants crit, hit or haste is not scored because those weights are per rating point and the
// enchant data does not say which unit it is in). Enchants of a later phase and class-locked
// ones are left out, and an enchant is offered only for the level 60 pick it can go on, by the
// client's own rule (`enchantFits`: its spell's SpellEquippedItems class, subclass and
// inventory-type masks): no weapon enchant on a held-in-off-hand item, no two-hand enchant on a
// one-hander, a shield enchant only on a shield.
//
// Consumables: the ones the sim's raid preset gives this spec at level 60, read from the source
// of truth the sim itself runs (sim/request/raid-consumes.json, data/curated/presets.json
// resolved per spec with the class kit, held current by a Go test), never from the published
// BiS file, which trails a curated change until the next refresh. Labels come from
// data/curated/presets.json, else the build's `simbuffs.json`, which also gives the icon; a
// short "what it gives" line comes from CONSUMABLE_GIVES, which a test holds to the preset. A
// Classic flask stacks with the battle and guardian elixirs, so all of them are listed together.
import { existsSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { repoRoot } from '../data/build-files';
import { bandEntry } from '../bis/load';
import { slotDisplayLabel } from '../bis/slot-display-labels';
import { enchantFits, type EnchantRule, type EnchantTarget } from '../gear/enchant-fit';
import { buildFile, buildItemCatalog, readBuildJson } from './build-files';

export { enchantFits, type EnchantTarget };
import { factionForFirstRace, rankedBisFile } from './class-dps';
import { guidesCopy } from './copy';
import { isRealIcon } from '../icons/placeholder';
import type { BisBand, BisStatWeight } from '../bis/types';
import type { GuideBis } from './guide-bis';

/** An `enchants.json` row; its `EnchantRule` says what it goes on (`enchantFits`). */
export interface EnchantRecord extends EnchantRule {
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
  /** The `enchants.json` id, the key of the enchant's spell tooltip. */
  id: number;
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
  /** A flask and the elixirs are all used together (Classic flasks stack with elixirs). */
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
  /** The level 60 pick in each slot; a slot with none gets no enchant row. */
  items: Readonly<Partial<Record<string, EnchantTarget>>>;
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
  iconExists: (icon: string) => boolean = () => true,
): EnchantRow[] {
  const byStat = new Map(weights.filter((w) => !(w.insignificant ?? false)).map((w) => [w.stat, w.weight]));
  const best = new Map<string, { score: number; record: EnchantRecord }>();
  for (const record of records) {
    if (record.phase > LATEST_PHASE || record.item_types.includes('kit')) continue;
    if (record.classes.length > 0 && !record.classes.includes(context.classSlug)) continue;
    const score = enchantScore(record.stats, byStat);
    if (score <= 0) continue;
    for (const slot of record.slots) {
      const item = context.items[slot];
      if (item === undefined || !enchantFits(record, item)) continue;
      const current = best.get(slot);
      if (current === undefined || score > current.score) best.set(slot, { score, record });
    }
  }
  return SLOT_ORDER.filter((slot) => best.has(slot)).map((slot) => {
    const { record } = best.get(slot)!;
    return {
      id: record.id,
      slot,
      slotLabel: slot === 'main_hand' ? 'Weapon' : slotDisplayLabel(slot, context.classSlug),
      name: record.name,
      // The enchant's own spell icon (the client gives most enchanting spells the same one, and
      // that is what the game shows): MUST 6, the same icon rows as the consumables.
      icon: isRealIcon(record.icon) && iconExists(record.icon) ? record.icon : undefined,
      gives: enchantGives(record.stats),
    };
  });
}

/** What each consumable the raid-ready preset uses gives. A new id needs a line here. */
export const CONSUMABLE_GIVES: Readonly<Record<string, string>> = {
  alcohol_rumsey_rum_black_label: 'Stamina',
  cerebral_cortex_compound: 'Intellect',
  conjured_demonic_rune: 'Mana, used on cooldown',
  elixir_of_firepower: 'Fire spell damage',
  elixir_of_fortitude: 'Health',
  elixir_of_frost_power: 'Frost spell damage',
  elixir_of_greater_defense: 'Armor',
  elixir_of_holy_power: 'Holy spell damage',
  elixir_of_sages: 'Spirit and spell critical strike',
  elixir_of_shadow_power: 'Shadow spell damage',
  elixir_of_the_mongoose: 'Agility and critical strike',
  flask_of_distilled_wisdom: 'A large pool of mana for the whole fight',
  flask_of_natural_aggression: 'Stamina and critical strike',
  flask_of_natural_precision: 'Stamina, and fewer of your attacks dodged or parried',
  flask_of_supreme_power: 'Spell power',
  flask_of_the_titans: '1200 health for two hours',
  food_baked_salmon: 'Spell damage',
  food_grilled_squid: 'Critical strike',
  food_nightfin_soup: 'Spell damage',
  food_poached_sunscale_salmon: 'Attack power and ranged attack power',
  food_runn_tum_tuber_surprise: 'Intellect',
  food_sages_tea: 'Healing power',
  food_smoked_desert_dumpling: 'Strength',
  food_tender_wolf_steak: 'Agility',
  gizzard_gum: 'Spirit',
  greater_arcane_elixir: 'Spell power',
  greater_clerics_elixir: 'Healing power',
  greater_mageblood_elixir: 'Mana per 5 sec',
  frenzy_potion: 'Attack power for 30 sec, on the potion cooldown',
  greater_stoneshield_potion: 'Armor for two minutes',
  ground_scorpok_assay: 'Agility',
  juju_might: 'Attack power',
  juju_power: 'Strength',
  conjured_rogue_thistle_tea: 'Energy, used when it runs low',
  'main_hand_imbue:brilliant_mana_oil': 'Mana per 5 sec and healing power, on the main hand',
  'main_hand_imbue:brilliant_wizard_oil': 'Spell power, on the main hand',
  'main_hand_imbue:elemental_sharpening_stone': 'Critical strike, on the main hand',
  'main_hand_stone:elemental_sharpening_stone': 'Critical strike, on the main hand',
  'main_hand_imbue:dense_sharpening_stone':
    'Weapon damage, on the main hand, until the Elemental stone at 50',
  major_frenzy_potion: 'Attack power for 30 sec, on the potion cooldown',
  major_mana_potion: 'Mana, used on cooldown',
  major_spellblasting_potion: 'Spell damage for 30 sec, on the potion cooldown',
  mighty_rage_potion: 'Rage',
  'off_hand_imbue:dense_sharpening_stone': 'Weapon damage, on the off hand',
  'off_hand_stone:dense_sharpening_stone': 'Weapon damage, on the off hand',
  r_o_i_d_s: 'Strength',
};

/** A row of the consumables table, or a group heading between rows. */
export type ConsumableLine = { header: string } | { row: ConsumableRow };

/** The table's lines: the flask and the elixirs first, all used together, then the food, oil
 *  and potions under their own heading. A flask never replaces an elixir in Classic. */
/** A consumable row's line under its name: only what the item's own tooltip does not say, the
 *  slot a weapon stone or oil goes on ("Critical strike, on the main hand" is "On the main hand");
 *  none for a plain stat, which restates the tooltip (h-audit 2026-11-09). */
export function consumableNote(gives: string): string | undefined {
  const comma = gives.indexOf(', ');
  if (comma === -1) return undefined;
  const note = gives.slice(comma + 2);
  return note.charAt(0).toUpperCase() + note.slice(1);
}

export function consumableLines(rows: readonly ConsumableRow[]): ConsumableLine[] {
  const kindOrder: Record<ConsumableKind, number> = { flask: 0, elixir: 1, other: 2 };
  const ordered = [...rows].sort((a, b) => kindOrder[a.kind] - kindOrder[b.kind]);
  const others = ordered.filter((row) => row.kind === 'other');
  const headed = others.length > 0 && others.length < ordered.length;
  return ordered.flatMap((row, index): ConsumableLine[] =>
    headed && row.kind === 'other' && ordered[index - 1]?.kind !== 'other'
      ? [{ header: guidesCopy.alsoUsed }, { row }]
      : [{ row }],
  );
}

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

/** Class kit imbues the guide's Rotation and class sections cover, not bought consumables. */
export const KIT_IMBUES: ReadonlySet<string> = new Set([
  'main_hand_imbue:deadly_poison',
  'main_hand_imbue:instant_poison',
  'off_hand_imbue:deadly_poison',
  'off_hand_imbue:instant_poison',
  'main_hand_imbue:windfury_weapon',
  'main_hand_imbue:rockbiter_weapon',
]);

interface RaidConsumesFile {
  specs: Record<string, { from_level: number; consumes: string[] }[]>;
}

/** The consumable ids the sim's raid preset sends `spec` at `level`, in the sim's order. */
export function raidConsumeIds(file: RaidConsumesFile, spec: string, level: number): string[] {
  const rows = (file.specs[spec] ?? []).filter((row) => row.from_level <= level);
  return (
    rows.reduce<{ from_level: number; consumes: string[] } | undefined>(
      (best, row) => (best === undefined || row.from_level > best.from_level ? row : best),
      undefined,
    )?.consumes ?? []
  );
}

/** Every `{ id, label }` the curated presets file names, at any depth. */
export function curatedLabels(node: unknown, into: Map<string, string> = new Map()): Map<string, string> {
  if (Array.isArray(node)) {
    for (const item of node) curatedLabels(item, into);
  } else if (typeof node === 'object' && node !== null) {
    const record = node as Record<string, unknown>;
    if (typeof record.id === 'string' && typeof record.label === 'string') into.set(record.id, record.label);
    for (const value of Object.values(record)) curatedLabels(value, into);
  }
  return into;
}

function readRepoJson<T>(relativePath: string): T {
  return JSON.parse(readFileSync(path.join(repoRoot(), relativePath), 'utf8')) as T;
}

/** The level 60 raid consumables the sim runs for `spec`, labelled, kit imbues left out. */
export function simRaidConsumes(
  spec: string,
  simBuffs: Readonly<Record<string, SimBuffEntry>>,
): { id: string; label: string }[] {
  const file = readRepoJson<RaidConsumesFile>('sim/request/raid-consumes.json');
  const labels = curatedLabels(readRepoJson<unknown>('data/curated/presets.json'));
  return raidConsumeIds(file, spec, 60)
    .filter((id) => !KIT_IMBUES.has(id))
    .map((id) => ({
      id,
      label: labels.get(id) ?? (simBuffs[id] ?? simBuffs[id.split(':').pop() ?? ''])?.name ?? id,
    }));
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
  items: EnchantContext['items'],
): GearExtras | undefined {
  const file = rankedBisFile(build, spec);
  const faction = factionForFirstRace(classSlug, recommendedRaces);
  if (file === undefined || faction === undefined) return undefined;
  const band: BisBand | undefined = bandEntry(file, 60, faction);
  if (band === undefined) return undefined;
  const records = readBuildJson<EnchantRecord[]>(build, 'enchants.json') ?? [];
  const iconExists = (icon: string): boolean => existsSync(buildFile(build, 'icons', `${icon}.webp`));
  const simBuffs =
    readBuildJson<{ entries: Record<string, SimBuffEntry> }>(build, 'simbuffs.json')?.entries ?? {};
  return {
    enchants: pickEnchants(records, band.weights, { classSlug, items }, iconExists),
    consumables: consumableRows(simRaidConsumes(spec, simBuffs), simBuffs, iconExists),
  };
}

/** The guide's tables, for the items its level 60 picks are: each slot's enchant is one the
 *  client lets go on that item (a two-hander takes a two-hand enchant, a shield a shield
 *  enchant, a held-in-off-hand item no weapon enchant, an empty slot none). */
export function guideGearExtras(
  build: string,
  classSlug: string,
  spec: string,
  recommendedRaces: readonly string[],
  bis: GuideBis | undefined,
): GearExtras | undefined {
  const top = bis?.panels.find((p) => p.faction === bis.defaultFaction && p.band === bis.defaultBand);
  const catalog = buildItemCatalog(build);
  const items = Object.fromEntries(
    (top?.rows ?? []).flatMap((row) => {
      const item = row.empty || row.itemId === undefined ? undefined : catalog.get(row.itemId);
      return item === undefined ? [] : [[row.slot, item] as const];
    }),
  );
  return loadGearExtras(build, classSlug, spec, recommendedRaces, items);
}
