// web/src/lib/sim/character.ts
// One character behind all four sources, convertible to the planner's BuildDraft and to the
// engine's CharacterSpec without losing a point or a slot.
//
// Two things here are not obvious and are load-bearing:
//
//   * `talent_level` is the level the talent spend implies, which is what the character
//     strip shows. It is NOT what the engine is sent. api.SimRequest.Validate refuses any
//     level but api.SimLevel (60) because the engine builds every character at
//     core.CharacterMaxLevel and proto.Player has no level field, so toCharacterSpec always
//     sends SIM_LEVEL.
//   * `buffs` and `consumables` are the engine's own ids in lower snake case, resolved off
//     the protobuf descriptors by sim/request and published as sim/request/IDS.md --
//     "battle_shout", "elixir_of_the_mongoose", "item:13452",
//     "off_hand_imbue:shadow_oil". An id the engine does not know
//     is ErrUnknownBuff or ErrUnknownConsume and fails the run, so nothing here invents one.
import { simCopy } from './copy';
import { pointsPerTree, ranksByTalent } from '../planner/derive';
import { decodeFS1, orderFromRanks, type FS1Item, type FS1Loadout, type FS1Set } from '../planner/fs1';
import { indexTalents, type TalentIndex } from '../planner/rules';
import type { PlannerStore } from '../planner/store.svelte';
import type { BuildDraft, ClassRow, Gear, RaceRow, Slot, TalentFile } from '../planner/types';
import { BASE_LEVEL, SLOTS } from '../planner/types';
import { SPECS } from './specs';
import type { CharacterSource, CharacterSpec, CooldownSpec, GearSlot } from './types';

export interface SimCharacter {
  name: string;
  spec: string;
  class_slug: string;
  race_slug: string;
  /** The level the talent spend implies. Shown in the strip; never sent to the engine. */
  talent_level: number;
  tree_version: string;
  point_order: number[];
  gear: Gear;
  /** The engine's own buff ids in lower snake case, not spell ids and not kebab case. */
  buffs: string[];
  consumables: string[];
  source: CharacterSource;
  /**
   * The gear as the engine takes it, enchants and suffixes included (contract 10.5).
   * `gear` above stays the planner's map of ids -- the strip, the planner link and
   * `BuildDraft` all read it and none of them models an enchant -- and this is the
   * authoritative list `toCharacterSpec` sends. The two always agree on item ids.
   */
  gear_slots: GearSlot[];
  /** From the export's `professions=` section (contract 10.5). Empty for every other source. */
  professions: string[];
  /**
   * Version 2 export sections (contract 7, corrected by 10.5). Empty for every other
   * source and for a version 1 export. These are the candidate lists `/sim/gear`,
   * `/sim/talents` and `/sim/drops` read; nothing on `/sim` itself renders them. They keep
   * the decoder's own shapes, so part B converts an `FS1Set` into the envelope's `GearSet`
   * once, where it builds the bulk request.
   */
  bags: FS1Item[];
  bank: FS1Item[];
  sets: FS1Set[];
  loadouts: FS1Loadout[];
}

/**
 * The only level the engine simulates. sim/request's own `engineCharacterLevel` is the same
 * number for the same reason, and its test asserts it against core.CharacterMaxLevel.
 */
export const SIM_LEVEL = 60;

/**
 * `race_slug` for a character whose source did not record a race. A combat log records
 * none at all, so `fromLoggedFight` sets this rather than substituting a race the player
 * does not play -- Forever's racials are two actives and two passives each, and guessing
 * one puts an orc's Blood Fury on a troll's sheet. The strip turns its race line into a
 * select while it is set, and the run control stays disabled: `toCharacterSpec` would
 * otherwise send an empty `race`, which `sim/request.ParseRace` refuses at the boundary.
 */
export const PENDING_RACE = '';

/** True while the character still needs a race from the player before it can be simmed. */
export function needsRace(character: SimCharacter): boolean {
  return character.race_slug === PENDING_RACE;
}

export function specForSplit(classSlug: string, split: readonly number[]): string {
  let best = 0;
  for (let i = 1; i < split.length; i += 1) {
    if (split[i] > split[best]) best = i;
  }
  const row = SPECS.find((entry) => entry.class_slug === classSlug && entry.tree_index === best);
  return row?.spec ?? `${classSlug}-${best}`;
}

export function specOf(index: TalentIndex, order: number[]): string {
  return specForSplit(index.file.class_slug, pointsPerTree(index, order));
}

/** The level a build of this many points belongs to. For the strip, not for the engine. */
export function talentLevel(order: number[]): number {
  return order.length === 0 ? BASE_LEVEL : Math.min(SIM_LEVEL, BASE_LEVEL + order.length);
}

export function toBuildDraft(
  character: SimCharacter,
  classes: readonly ClassRow[],
  races: readonly RaceRow[],
): BuildDraft | null {
  const classRow = classes.find((row) => row.slug === character.class_slug);
  const raceRow = races.find((row) => row.slug === character.race_slug);
  if (classRow === undefined || raceRow === undefined) return null;
  return {
    class_id: classRow.id,
    race_id: raceRow.id,
    tree_version: character.tree_version,
    point_order: [...character.point_order],
    gear: { ...character.gear },
  };
}

export interface CharacterExtras {
  name?: string;
  buffs?: string[];
  consumables?: string[];
  source: CharacterSource;
}

export function fromBuildDraft(
  draft: BuildDraft,
  talents: TalentFile,
  classes: readonly ClassRow[],
  races: readonly RaceRow[],
  extras: CharacterExtras,
): SimCharacter | null {
  const classRow = classes.find((row) => row.id === draft.class_id);
  const raceRow = races.find((row) => row.id === draft.race_id);
  if (classRow === undefined || raceRow === undefined) return null;
  if (talents.class_slug !== classRow.slug) return null;
  const index = indexTalents(talents);
  return {
    name: extras.name ?? classRow.name,
    spec: specOf(index, draft.point_order),
    class_slug: classRow.slug,
    race_slug: raceRow.slug,
    talent_level: talentLevel(draft.point_order),
    tree_version: draft.tree_version,
    point_order: [...draft.point_order],
    gear: { ...(draft.gear ?? {}) },
    buffs: [...(extras.buffs ?? [])],
    consumables: [...(extras.consumables ?? [])],
    source: extras.source,
    // A planner build has no enchants, so gear_slots is the id map converted -- the same
    // "the two always agree on item ids" promise every other source keeps.
    gear_slots: gearSlots(draft.gear ?? {}),
    professions: [],
    bags: [],
    bank: [],
    sets: [],
    loadouts: [],
  };
}

/**
 * The planner's live state as a character. It is `fromBuildDraft` over `store.toDraft()`,
 * so there is exactly one conversion in the codebase and the sim cannot disagree with the
 * planner about what a build is. Null while the planner's data is still loading, and null
 * again when `classRow` or `raceRow` cannot resolve: `store.toDraft()` throws for either
 * (`store.svelte.ts`'s own guard), and both can go unresolved on untrusted input this
 * function does not control -- an unvalidated `?class=`/`?race=` query string, or a decoded
 * FS1 code naming a class or race the reference data does not have. Treating that the same
 * as "still loading" is what keeps the automatic `$effect` in `Planner.svelte` from throwing
 * on a bad link instead of simply staying idle.
 */
export function characterFromPlanner(store: PlannerStore): SimCharacter | null {
  if (store.talents === null || store.classes.length === 0) return null;
  if (store.classRow === null || store.raceRow === null) return null;
  return fromBuildDraft(store.toDraft(), store.talents, store.classes, store.races, {
    name: store.classRow.name,
    source: { kind: 'manual', ref: '', captured_at: new Date().toISOString() },
  });
}

/**
 * The engine's talents string: one decimal digit per talent in tab order, trees joined by
 * "-", trailing zeros trimmed inside each tree. An empty tree is the empty string, which is
 * why a build with nothing in the last tree ends in a dash, exactly as the contract's
 * example "01102123133-12312312-" does. Ranks are clamped to one digit; vanilla's maximum
 * of five never approaches it, and a two-digit rank would silently shift every talent after
 * it in the string.
 */
export function talentsString(index: TalentIndex, order: number[]): string {
  const ranks = ranksByTalent(order);
  return index.trees
    .map((tree) =>
      tree.talents
        .map((talent) => Math.min(9, ranks.get(talent.id) ?? 0).toString(10))
        .join('')
        .replace(/0+$/, ''),
    )
    .join('-');
}

/** The planner's Gear map as the engine's list. An empty slot is left out, not sent as 0. */
export function gearSlots(gear: Gear): GearSlot[] {
  return SLOTS.flatMap((slot) => {
    const itemId = gear[slot as Slot];
    return itemId === undefined ? [] : [{ slot, item_id: itemId }];
  });
}

/** The inverse of gearSlots: the engine's list back as the planner's Gear map. */
export function gearFromSlots(gear: readonly GearSlot[]): Gear {
  const out: Gear = {};
  for (const { slot, item_id } of gear) out[slot as Slot] = item_id;
  return out;
}

/**
 * The inverse of talentsString: one decimal digit per talent in tab order, per tree,
 * dash-joined ("01102123133-12312312-"). Every rank talentsString ever wrote is a single
 * digit, so this is exact, unlike an FS1 code's own base-36 tree field, which loses
 * nothing either but is not what CharacterSpec carries.
 */
export function ranksFromTalentsString(talents: string): number[][] {
  return talents.split('-').map((tree) => Array.from(tree, (digit) => Number.parseInt(digit, 10) || 0));
}

/**
 * The character as the engine wants it: plain JSON, no protobuf. `sim/request` turns this
 * into a RaidSimRequest inside our own wasm, so this function is the whole of the web's
 * side of the conversion.
 *
 * `professions` is left unset rather than sent empty when the character has none recorded
 * (contract 10.5's `professions=` section is the only source that ever fills it): an empty
 * array would claim we had looked and found none.
 *
 * `cooldowns` is likewise omitted rather than sent as `[]`: an empty list would claim we
 * had scheduled something and found nothing, when the truth is "every cooldown on
 * cooldown", which is what an absent field means to the engine (contract 1.7).
 */
export function toCharacterSpec(
  character: SimCharacter,
  index: TalentIndex,
  buffs: string[],
  consumes: string[],
  cooldowns: readonly CooldownSpec[] = [],
): CharacterSpec {
  const spec: CharacterSpec = {
    name: character.name,
    race: character.race_slug,
    class: character.class_slug,
    // Always 60. api.SimRequest.Validate refuses anything else, so a level-33 request
    // would be rejected at the boundary rather than simmed as written.
    level: SIM_LEVEL,
    talents: talentsString(index, character.point_order),
    // The slot list when the source gave one, the id map otherwise. Never both, and never
    // a merge: one of the two is the truth about this character's gear and it is this one.
    gear:
      character.gear_slots.length > 0
        ? character.gear_slots.map((slot) => ({ ...slot }))
        : gearSlots(character.gear),
    buffs: [...buffs],
    consumes: [...consumes],
    // Still omitted when empty, for the reason the original comment gives: an empty list
    // would claim we had looked and found none.
    ...(character.professions.length === 0 ? {} : { professions: [...character.professions] }),
  };
  return cooldowns.length === 0
    ? spec
    : { ...spec, cooldowns: cooldowns.map((row) => ({ ...row, at_sec: [...row.at_sec] })) };
}

export function plannerHrefFor(character: SimCharacter): string {
  const params = new URLSearchParams({ class: character.class_slug, race: character.race_slug });
  return `/planner?${params.toString()}`;
}

export type CharacterResult = { ok: true; character: SimCharacter } | { ok: false; message: string };

/**
 * An addon export into a character. Decoding and the rank-to-order reconstruction are the
 * planner's (lib/planner/fs1.ts): the addon writes one string and both pillars read it with
 * the same code, so a format change cannot reach one and miss the other.
 */
export function characterFromFs1(
  code: string,
  talents: TalentFile,
  classes: readonly ClassRow[],
  races: readonly RaceRow[],
  source: CharacterSource,
  name?: string,
): CharacterResult {
  const decoded = decodeFS1(code);
  if (!decoded.ok) return { ok: false, message: decoded.message };
  if (decoded.build.classSlug !== talents.class_slug) {
    return { ok: false, message: simCopy.classMismatch(decoded.build.classSlug, talents.class_slug) };
  }

  const index = indexTalents(talents);
  const { order, dropped } = orderFromRanks(index, decoded.build.treeRanks);
  if (dropped.length > 0) {
    const names = dropped.map((id) => index.byId.get(id)?.name ?? String(id)).join(', ');
    return { ok: false, message: simCopy.unreachableTalents(names) };
  }

  const classRow = classes.find((row) => row.slug === decoded.build.classSlug);
  if (classRow === undefined) {
    return { ok: false, message: simCopy.classMismatch(decoded.build.classSlug, talents.class_slug) };
  }
  // No fallback race. The first draft substituted the first row in the file when the code's
  // race was unknown, which silently simmed an orc's racials for a troll. Forever's race
  // table has ten rows and two of them are new (high-order-skyborne, windshaper-skyborne),
  // so an unknown slug means the export is from another build, and saying so is the answer.
  const raceRow = races.find((row) => row.slug === decoded.build.raceSlug);
  if (raceRow === undefined) {
    return { ok: false, message: simCopy.unknownRace(decoded.build.raceSlug) };
  }

  return {
    ok: true,
    character: {
      name: name ?? classRow.name,
      spec: specOf(index, order),
      class_slug: classRow.slug,
      race_slug: raceRow.slug,
      talent_level: talentLevel(order),
      tree_version: decoded.build.dataBuild,
      point_order: order,
      gear: { ...decoded.build.gear },
      // decoded.build is DecodedFS1Build: the decoder always fills gearSlots (and every
      // other version 2 field below), so no `?? []` is needed here the way an
      // FS1Build-typed encoder caller would need one.
      gear_slots: decoded.build.gearSlots.map((entry) => ({
        slot: entry.slot,
        item_id: entry.itemId,
        ...(entry.enchant === undefined ? {} : { enchant: entry.enchant }),
        ...(entry.suffix === undefined ? {} : { suffix: entry.suffix }),
      })),
      professions: [...decoded.build.professions],
      buffs: [],
      consumables: [],
      source,
      bags: [...decoded.build.bags],
      bank: [...decoded.build.bank],
      sets: decoded.build.sets.map((set) => ({
        name: set.name,
        gear: set.gear.map((entry) => ({ ...entry })),
      })),
      loadouts: decoded.build.loadouts.map((row) => ({
        name: row.name,
        treeRanks: row.treeRanks.map((tree) => [...tree]),
      })),
    },
  };
}
