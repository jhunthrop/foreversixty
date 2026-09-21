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
import {
  decodeFS1,
  encodeFS1V2,
  orderFromRanks,
  type FS1Item,
  type FS1Loadout,
  type FS1Set,
} from '../planner/fs1';
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

/**
 * A build's total spent points, straight from the engine's own talents STRING -- no
 * `point_order` (the click order that produced it) required. A saved sim's stored request
 * carries only that final string, never the order (CharacterStrip.svelte's own comment on
 * why its point count used to read blank there), but the string alone already says how many
 * points were spent: every character `ranksFromTalentsString` decodes is one talent's own
 * rank (0-9, `talentsString`'s own encoder caps it there), so the sum of every rank, across
 * every tree, is the point count -- the same number `point_order.length` would have given
 * had the order been available (2026-09-21 result-page review round 3, newcomer's own
 * finding).
 */
export function talentPointsFromString(talents: string): number {
  return ranksFromTalentsString(talents)
    .flat()
    .reduce((sum, rank) => sum + rank, 0);
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
 * The gear the inline planner (Planner.svelte's `gear` prop) should start from: `gear_slots`
 * converted back to the planner's id map when the character has one, the id map itself
 * otherwise. Mirrors `toCharacterSpec`'s own precedence above rather than inventing a second
 * one -- the two must never disagree about which of a character's two gear fields is the
 * truth. Always a fresh object; mutating the result never touches the character's own gear
 * (dps D39/D40: without a seeded gear, the inline build card sims an empty character while
 * the comparison table sims the real one).
 */
export function plannerGearFor(character: SimCharacter): Gear {
  return character.gear_slots.length > 0 ? gearFromSlots(character.gear_slots) : { ...character.gear };
}

/**
 * A `CharacterSpec` (the engine's own JSON shape) as an FS1 version 2 code -- the one place
 * this conversion is written. "Run this yourself" (SimView.svelte, a saved result's own
 * request) and the request drawer's Apply (store-request.ts, a pasted/edited request) both
 * need a code to feed the store's existing `?code=`/`fromPlannerCode` bootstrap, and both
 * used to build one by hand; the two copies drifted once already (final whole-branch
 * review, finding 1) when `encodeFS1` was upgraded to `encodeFS1V2` in only one of them.
 *
 * `encodeFS1V2`, not `encodeFS1`: a `CharacterSpec.gear` entry carries any enchant or suffix
 * (contract 10.5), and `encodeFS1` would drop exactly what `characterFromFs1` now keeps.
 * The six version-2 sections travel empty -- a `CharacterSpec` never carries bags, bank,
 * sets, loadouts, or a decoder's own `ignored` list; only `professions` round-trips.
 */
export function codeForCharacterSpec(spec: CharacterSpec, dataBuild: string): string {
  return encodeFS1V2({
    dataBuild,
    classSlug: spec.class,
    raceSlug: spec.race,
    treeRanks: ranksFromTalentsString(spec.talents),
    gear: gearFromSlots(spec.gear),
    gearSlots: spec.gear.map((slot) => ({
      slot: slot.slot as Slot,
      itemId: slot.item_id,
      ...(slot.enchant === undefined ? {} : { enchant: slot.enchant }),
      ...(slot.suffix === undefined ? {} : { suffix: slot.suffix }),
    })),
    bags: [],
    bank: [],
    sets: [],
    loadouts: [],
    professions: [...(spec.professions ?? [])],
    ignored: [],
  });
}

/**
 * `/planner?code=...` for a `CharacterSpec`, through `codeForCharacterSpec` -- the one URL
 * template every "Open in planner"/"Plan it" link builds. `plannerHrefFor` below reaches it
 * once it has turned a live character's `point_order` into a `CharacterSpec` via
 * `toCharacterSpec`; `combos.ts`'s `planItHref` reaches it straight from a bulk result's own
 * stored `CharacterSpec`, substitutions applied; a saved sim's "Open in planner" link
 * (SavedSim.svelte) reaches it the same way `planItHref` does -- the stored request's own
 * `CharacterSpec.talents` is already the ground truth there, with no `TalentIndex` or
 * `point_order` to reconstruct one through at all (defect fixed here: SavedSim.svelte used
 * to build its strip's `character` with `point_order: []` and hand it to `plannerHrefFor`,
 * which zeroed every talent through `toCharacterSpec`'s `talentsString(index, [])` -- a
 * saved sim's own point-order is genuinely unknowable, the honest-empty rule `sources.ts`
 * already follows for a combat log, but its final talent RANKS are not: they are sitting
 * in the stored request's own `character.talents` string, untouched).
 */
export function plannerHrefForSpec(spec: CharacterSpec, dataBuild: string): string {
  return `/planner?code=${encodeURIComponent(codeForCharacterSpec(spec, dataBuild))}`;
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
/** The two classes with a rage bar. A druid has one in bear form, which the engine models. */
const RAGE_CLASSES: ReadonlySet<string> = new Set(['warrior', 'druid']);
/** Consumables that do nothing but add rage. */
const RAGE_CONSUMABLES: ReadonlySet<string> = new Set(['mighty_rage_potion']);

/**
 * The consumables this class can actually use. In game a rogue who drinks a rage potion gets
 * nothing, so leaving it out of the request is what happened, not a rewrite of what the
 * player asked for -- and it has to be left out, because the engine adds the rage to a rage
 * bar that only a warrior or a druid has and dereferences nil for everyone else. Every
 * request is built here, so the Raid-buffed preset, a Custom panel tick and a saved sim's
 * re-run all pass through the one guard.
 */
export function usableConsumables(classSlug: string, consumes: readonly string[]): string[] {
  if (RAGE_CLASSES.has(classSlug)) return [...consumes];
  return consumes.filter((id) => !RAGE_CONSUMABLES.has(id));
}

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
    consumes: usableConsumables(character.class_slug, consumes),
    // Still omitted when empty, for the reason the original comment gives: an empty list
    // would claim we had looked and found none.
    ...(character.professions.length === 0 ? {} : { professions: [...character.professions] }),
  };
  return cooldowns.length === 0
    ? spec
    : { ...spec, cooldowns: cooldowns.map((row) => ({ ...row, at_sec: [...row.at_sec] })) };
}

/**
 * Where "Open in planner" goes. `index` is null while the character's talent file is still
 * streaming in (or never resolves): the link must render from first paint rather than stay
 * absent or disabled, so it starts here -- class and race only, exactly the URL this
 * function has always emitted -- and upgrades in place once the file loads.
 *
 * With an index, the character's talents, gear and professions ride along as an FS1 v2
 * code, through `codeForCharacterSpec` -- the same conversion `ComboResults.svelte`'s own
 * "Open in planner" link already uses, so the two links can never disagree about what a
 * code encodes (newcomer MAJOR, review.md:82-88: the link used to open a blank character).
 *
 * The spec itself is `toCharacterSpec`, not a second hand-built literal: a first version of
 * this function rebuilt the same fields inline and, in doing so, dropped `professions` --
 * `toCharacterSpec` omits that key only when the character truly has none (character.ts:319),
 * and the hand-built literal had no such key at all, so a combat-log character with
 * professions silently lost them from its own planner link (task-2 fix round 1's review,
 * Important). `character.buffs`/`character.consumables` ride along because `toCharacterSpec`
 * requires them, but `codeForCharacterSpec` never reads a spec's `buffs`/`consumes` fields
 * (character.ts:296-309 above), so neither ends up in the encoded code either way.
 */
export function plannerHrefFor(character: SimCharacter, index: TalentIndex | null): string {
  if (index === null) {
    const params = new URLSearchParams({ class: character.class_slug, race: character.race_slug });
    return `/planner?${params.toString()}`;
  }
  const spec = toCharacterSpec(character, index, character.buffs, character.consumables);
  return plannerHrefForSpec(spec, character.tree_version);
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
