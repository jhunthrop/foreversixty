// web/src/lib/planner/fs1.ts
// The addon's export string, both directions.
//
//   FS1:<data-build>:<class-slug>:<race-slug>:<t1>/<t2>/<t3>:<slot>=<item-id>,...
//
// Each <tN> is that tree's talent ranks in tab order, one base-36 digit each, trailing
// zeros trimmed. The format is docs/superpowers/specs/2026-09-13-phase-2-addon-design.md's,
// and the addon decodes the same string, so neither side may drift.
//
// A decoder failure names its reason. The addon spec requires it and it is the right call
// anyway: "that code is FS2, this site reads FS1" tells someone what to do and "invalid
// code" does not.
//
// Contract 7, corrected by 10.5: everything before the first `|` is version 1 unchanged --
// every existing decoder keeps working -- and what follows is a fixed sequence of sections
// (bags, bank, sets, loadouts, professions), each optional, each ignored (and named) by the
// decoder when it does not recognise the section. 10.5 additionally lets a gear entry --
// on the head, in a bag, in the bank, or inside a named set -- carry `item_id[:enchant[:suffix]]`
// instead of a bare id; a bare id is still legal everywhere one was legal before, which is
// what keeps a version 1 string readable by this decoder and a version 2 string's plain
// entries readable by a version 1 one.
import { canAddPoint, withPoint, type TalentIndex } from './rules';
import { SLOTS, type Gear, type Slot } from './types';

export const FS1_PREFIX = 'FS1';
const TREES = 3;
/**
 * Contract 7 raises this from 2,048: a version 2 code carries bags, bank, named sets and
 * named loadouts, and a full bank is several thousand characters on its own. It is still a
 * finite bound checked before any splitting, which is what keeps a hostile query value
 * from doing real work.
 */
export const MAX_CODE_LENGTH = 16_384;

/** One bag or bank item: `item_id[:enchant[:suffix]]`, with no slot in the string. */
export interface FS1Item {
  itemId: number;
  enchant?: number;
  suffix?: number;
}

/** One gear entry: `<slot>=item_id[:enchant[:suffix]]` (contract 10.5). */
export interface FS1GearSlot {
  slot: Slot;
  itemId: number;
  enchant?: number;
  suffix?: number;
}

export interface FS1Set {
  name: string;
  gear: FS1GearSlot[];
}

export interface FS1Loadout {
  name: string;
  /** One array per tree, one entry per talent in tab order -- FS1Build.treeRanks's shape. */
  treeRanks: number[][];
}

export interface FS1Build {
  dataBuild: string;
  classSlug: string;
  raceSlug: string;
  treeRanks: number[][];
  /**
   * The planner's map of slot to item id. Lossy by design: the strip, the planner link and
   * `BuildDraft` all read it and none of them models an enchant or a suffix.
   */
  gear: Gear;
  /**
   * The whole truth about the gear, enchants and suffixes included (contract 10.5). The
   * decoder always fills it; it is optional only so a caller holding nothing but a `Gear`
   * map can still build an `FS1Build` without writing `gearSlotsFrom(gear)` by hand.
   */
  gearSlots?: FS1GearSlot[];
  /**
   * Version 2 (contract 7, corrected by 10.5), all optional (spec section 7's own words) --
   * `decodeFS1` always fills them (with `[]` when the code carries no such section), but an
   * `encodeFS1`/`encodeFS1V2` caller building a version 1 code has nothing to say about any
   * of them and should not have to write six empty arrays to say so.
   *
   * `SimCharacter`'s mirrors of these (`character.ts`) are deliberately NOT optional: that
   * interface is only ever the decoder's own output or another source's honest equivalent,
   * every source can fill it, and optional there would push `?? []` onto every reader
   * instead of onto the one place (`encodeFS1V2`) that actually needs a default.
   */
  bags?: FS1Item[];
  bank?: FS1Item[];
  sets?: FS1Set[];
  loadouts?: FS1Loadout[];
  professions?: string[];
  /**
   * The character's current guild, from `GetGuildInfo("player")` (section 1.1 of the
   * guild-membership design). Absent means unguilded -- never defaulted to `{}`, the
   * way `bags`/`bank`/`sets`/`loadouts`/`professions` default to `[]`, because "no
   * guild" and "an empty guild" are not the same fact.
   */
  guild?: { name: string; rankIndex: number };
  /** Section names the decoder did not recognise, reported rather than silently dropped. */
  ignored?: string[];
}

export type FS1Error = { ok: false; message: string };

/**
 * `decodeFS1`'s own return shape. `FS1Build` leaves `gearSlots` and the six version-2
 * fields optional for an *encoder* caller with nothing to add (contract section 7's "all
 * optional") -- but the decoder itself always fills every one of them, `[]` when the code
 * carries no such section. Every reader of a decoded build (`character.ts`, this file's
 * own tests) can therefore read them with no null check, which is the whole reason this
 * type exists rather than reusing `FS1Build` as `decodeFS1`'s own return shape too.
 */
export type DecodedFS1Build = FS1Build &
  Required<Pick<FS1Build, 'gearSlots' | 'bags' | 'bank' | 'sets' | 'loadouts' | 'professions' | 'ignored'>>;

export type FS1Result = { ok: true; build: DecodedFS1Build } | FS1Error;

function encodeTree(ranks: number[]): string {
  const digits = ranks.map((rank) => Math.max(0, Math.min(35, Math.round(rank))).toString(36));
  const trimmed = digits.join('').replace(/0+$/, '');
  return trimmed === '' ? '0' : trimmed;
}

/** The three tree fields, slash-joined. `encodeFS1` and `encodeFS1V2` share it. */
function encodeTrees(treeRanks: readonly number[][]): string {
  return Array.from({ length: TREES }, (_, index) => encodeTree(treeRanks[index] ?? [])).join('/');
}

/** Version 1: bare ids, in SLOTS order. `encodeFS1`'s output is unchanged by 10.5. */
function encodeGearList(gear: Gear): string {
  return SLOTS.filter((slot) => gear[slot] !== undefined)
    .map((slot) => `${slot}=${gear[slot]}`)
    .join(',');
}

export function encodeFS1(build: FS1Build): string {
  return [
    FS1_PREFIX,
    build.dataBuild,
    build.classSlug,
    build.raceSlug,
    encodeTrees(build.treeRanks),
    encodeGearList(build.gear),
  ].join(':');
}

type Parsed<T> = { ok: true; value: T } | FS1Error;

function parseTrees(field: string): Parsed<number[][]> {
  const treeStrings = field.split('/');
  if (treeStrings.length !== TREES) {
    return { ok: false, message: `That code has ${treeStrings.length} talent trees; a build has ${TREES}.` };
  }
  const treeRanks: number[][] = [];
  for (const tree of treeStrings) {
    const ranks: number[] = [];
    for (const digit of tree) {
      const rank = Number.parseInt(digit, 36);
      if (Number.isNaN(rank)) {
        return { ok: false, message: `That code has an unreadable talent rank: ${digit}.` };
      }
      ranks.push(rank);
    }
    treeRanks.push(ranks);
  }
  return { ok: true, value: treeRanks };
}

/**
 * `<slot>=item_id[:enchant[:suffix]]`, joined by commas (contract 10.5). A bare id is
 * still legal, which is what keeps a version-1 string readable by this and a version-2
 * string's plain entries readable by a version-1 decoder.
 */
function parseGearList(field: string): Parsed<FS1GearSlot[]> {
  const slots: FS1GearSlot[] = [];
  if (field === '') return { ok: true, value: slots };
  for (const entry of field.split(',')) {
    const eq = entry.split('=');
    // Exactly one `=`: `entry.split('=')` on "head=12640=99" is ["head","12640","99"], and
    // destructuring only the first two would silently discard the "=99" instead of refusing
    // an entry that is not this grammar's shape.
    if (eq.length !== 2) {
      return { ok: false, message: `That code has an unreadable gear entry: ${entry}.` };
    }
    const [slot, value] = eq;
    if (!(SLOTS as readonly string[]).includes(slot)) {
      return { ok: false, message: `That code names a slot this planner does not have: ${slot}.` };
    }
    // Digits only, never Number.parseInt on the whole field: parseInt stops at the first
    // non-digit and would silently turn "12640abc" into the item id 12640.
    const parts = value.split(':');
    if (parts.length > 3 || parts.some((part) => !/^\d+$/.test(part))) {
      return { ok: false, message: `That code has an unreadable gear entry: ${entry}.` };
    }
    const [itemId, enchant, suffix] = parts.map((part) => Number.parseInt(part, 10));
    slots.push({
      slot: slot as Slot,
      itemId,
      ...(enchant === undefined ? {} : { enchant }),
      ...(suffix === undefined ? {} : { suffix }),
    });
  }
  return { ok: true, value: slots };
}

/** The planner's lossy view of a gear list: slot to item id, enchants dropped. */
function gearMapOf(slots: readonly FS1GearSlot[]): Gear {
  const gear: Gear = {};
  for (const entry of slots) gear[entry.slot] = entry.itemId;
  return gear;
}

/** The inverse, for a caller holding only the planner's map. */
export function gearSlotsFrom(gear: Gear): FS1GearSlot[] {
  return SLOTS.filter((slot) => gear[slot] !== undefined).map((slot) => ({
    slot,
    itemId: gear[slot] as number,
  }));
}

/** `item_id[:enchant[:suffix]]`, joined by commas. Each number is digits only, as gear is. */
function parseItemList(field: string): Parsed<FS1Item[]> {
  if (field === '') return { ok: true, value: [] };
  const items: FS1Item[] = [];
  for (const entry of field.split(',')) {
    const parts = entry.split(':');
    if (parts.length > 3 || parts.some((part) => !/^\d+$/.test(part))) {
      return { ok: false, message: `That code has an unreadable item entry: ${entry}.` };
    }
    const [itemId, enchant, suffix] = parts.map((part) => Number.parseInt(part, 10));
    items.push({
      itemId,
      ...(enchant === undefined ? {} : { enchant }),
      ...(suffix === undefined ? {} : { suffix }),
    });
  }
  return { ok: true, value: items };
}

/**
 * `decodeURIComponent`, but a malformed escape (a bare `%`, or `%` not followed by two hex
 * digits) reads as itself rather than throwing. A `URIError` here would crash the caller --
 * `Planner.svelte` calls `decodeFS1` synchronously with no `try`, and `sources.ts`'s
 * `fromAddonExport` calls `characterFromFs1` outside its own -- over a set or loadout name
 * a third-party addon exported without encoding, which is an honest string the player
 * typed, not an attack to refuse.
 */
function decodeName(raw: string): string {
  try {
    return decodeURIComponent(raw);
  } catch {
    return raw;
  }
}

/** `<name>=<payload>;…`, the name URL-encoded so it may carry `;`, `=` and `|`. */
function parseNamed(field: string): { name: string; payload: string }[] {
  if (field === '') return [];
  return field.split(';').map((entry) => {
    const split = entry.indexOf('=');
    return split === -1
      ? { name: decodeName(entry), payload: '' }
      : { name: decodeName(entry.slice(0, split)), payload: entry.slice(split + 1) };
  });
}

/**
 * `<name>:<rank-index>`, split on the first colon -- a guild is one fact, not the
 * `name=payload;…` collection grammar `sets`/`loadouts` use. The name side never throws
 * (decodeName); the rank side must be digits-only or the whole code is refused, since a
 * malformed *known* section refuses the whole code (only an unrecognised section name is
 * forgiven).
 */
function parseGuild(field: string): Parsed<{ name: string; rankIndex: number }> {
  const at = field.indexOf(':');
  const namePart = at === -1 ? field : field.slice(0, at);
  const rankPart = at === -1 ? '' : field.slice(at + 1);
  if (!/^\d+$/.test(rankPart)) {
    return { ok: false, message: `That code has an unreadable guild rank: ${rankPart}.` };
  }
  return { ok: true, value: { name: decodeName(namePart), rankIndex: Number.parseInt(rankPart, 10) } };
}

export function decodeFS1(code: string): FS1Result {
  // Before any splitting or parsing: the cheapest possible check, and the one that keeps a
  // hostile multi-megabyte query value from doing any real work at all.
  if (code.length > MAX_CODE_LENGTH) {
    return { ok: false, message: 'That code is too long to read.' };
  }
  // Version 2 first, because everything before the first pipe is version 1 unchanged
  // (contract 7) -- so the version 1 parse below never has to know sections exist.
  const [head, ...sections] = code.trim().split('|');

  const parts = head.split(':');
  if (parts[0] !== FS1_PREFIX) {
    const named = parts[0] === undefined || parts[0] === '' ? 'unlabelled' : parts[0];
    return { ok: false, message: `That code is ${named}; this site reads ${FS1_PREFIX}.` };
  }
  if (parts.length < 6) return { ok: false, message: 'That code is missing its talent and gear fields.' };

  const [, dataBuild, classSlug, raceSlug, treeField, ...gearParts] = parts;
  const trees = parseTrees(treeField);
  if (!trees.ok) return trees;
  const gear = parseGearList(gearParts.join(':'));
  if (!gear.ok) return gear;

  // No `: FS1Build` annotation: the decoder always fills every version-2 field (with `[]`
  // when the code carries no such section), so this stays typed with them required rather
  // than inheriting `FS1Build`'s own optionality -- which exists for an *encoder* caller
  // with nothing to say, not for this function's own return value. A required field is
  // still assignable to `DecodedFS1Build`, below, where `FS1Result` carries it.
  const build = {
    dataBuild,
    classSlug,
    raceSlug,
    treeRanks: trees.value,
    gear: gearMapOf(gear.value),
    gearSlots: gear.value,
    bags: [] as FS1Item[],
    bank: [] as FS1Item[],
    sets: [] as FS1Set[],
    loadouts: [] as FS1Loadout[],
    professions: [] as string[],
    guild: undefined as { name: string; rankIndex: number } | undefined,
    ignored: [] as string[],
  };

  for (const section of sections) {
    const split = section.indexOf('=');
    const name = split === -1 ? section : section.slice(0, split);
    const field = split === -1 ? '' : section.slice(split + 1);
    if (name === 'bags' || name === 'bank') {
      const items = parseItemList(field);
      if (!items.ok) return items;
      build[name] = items.value;
    } else if (name === 'sets') {
      for (const { name: setName, payload } of parseNamed(field)) {
        const setGear = parseGearList(payload);
        if (!setGear.ok) return setGear;
        build.sets.push({ name: setName, gear: setGear.value });
      }
    } else if (name === 'loadouts') {
      for (const { name: loadoutName, payload } of parseNamed(field)) {
        const loadoutTrees = parseTrees(payload);
        if (!loadoutTrees.ok) return loadoutTrees;
        build.loadouts.push({ name: loadoutName, treeRanks: loadoutTrees.value });
      }
    } else if (name === 'professions') {
      // Slugs, unvalidated here: IDS.md's profession list is the vocabulary and
      // sim/request refuses one it cannot map, naming it. Silently dropping a slug this
      // decoder did not recognise would hide exactly that error -- but an empty entry from
      // a stray comma ("a,,b" or a trailing "a,") is not a slug at all, only a formatting
      // artifact, so those alone are filtered rather than forwarded to fail there instead.
      build.professions = field === '' ? [] : field.split(',').filter((slug) => slug !== '');
    } else if (name === 'guild') {
      const guild = parseGuild(field);
      if (!guild.ok) return guild;
      build.guild = guild.value;
    } else if (name !== '') {
      // Contract 7: unknown sections are ignored by the decoder and reported in its
      // result. An addon a version ahead of the site is a thing that will happen, and
      // refusing its whole string would make the site useless the day it ships.
      build.ignored.push(name);
    }
  }

  return { ok: true, build };
}

export interface OrderFromRanks {
  /** A legal point order producing the requested ranks, as far as they are reachable. */
  order: number[];
  /** Talent ids whose ranks could not be placed legally, so the caller can say so. */
  dropped: number[];
}

/**
 * An FS1 code carries a final tree, not the order it was spent in -- nothing in the game
 * records that. The planner needs an order, so this reconstructs a legal one by repeatedly
 * spending wherever the planner's own rules allow, lowest tier first, left to right. It
 * uses canAddPoint rather than reimplementing the gates, so an order it produces is one
 * the planner and the API both accept by construction.
 */
export function orderFromRanks(index: TalentIndex, treeRanks: number[][]): OrderFromRanks {
  const wanted = new Map<number, number>();
  index.trees.forEach((tree, treeIndex) => {
    const ranks = treeRanks[treeIndex] ?? [];
    tree.talents.forEach((talent, talentIndex) => {
      const rank = Math.min(ranks[talentIndex] ?? 0, talent.max_rank);
      if (rank > 0) wanted.set(talent.id, rank);
    });
  });

  const spent = new Map<number, number>();
  let order: number[] = [];
  let progress = true;
  while (progress) {
    progress = false;
    for (const tree of index.trees) {
      const byTier = [...tree.talents].sort((a, b) => a.tier - b.tier || a.column - b.column);
      for (const talent of byTier) {
        const target = wanted.get(talent.id) ?? 0;
        while ((spent.get(talent.id) ?? 0) < target && canAddPoint(index, order, talent.id).ok) {
          order = withPoint(order, talent.id);
          spent.set(talent.id, (spent.get(talent.id) ?? 0) + 1);
          progress = true;
        }
      }
    }
  }

  const dropped = [...wanted.entries()]
    .filter(([id, target]) => (spent.get(id) ?? 0) < target)
    .map(([id]) => id)
    .sort((a, b) => a - b);

  return { order, dropped };
}

function encodeItems(items: readonly FS1Item[]): string {
  return items
    .map((item) =>
      [item.itemId, item.enchant, item.suffix].filter((part): part is number => part !== undefined).join(':'),
    )
    .join(',');
}

/** Version 2: `<slot>=item_id[:enchant[:suffix]]`, in SLOTS order (contract 10.5). */
function encodeGearSlots(slots: readonly FS1GearSlot[]): string {
  const bySlot = new Map(slots.map((entry) => [entry.slot, entry]));
  return SLOTS.filter((slot) => bySlot.has(slot))
    .map((slot) => {
      const entry = bySlot.get(slot) as FS1GearSlot;
      return `${slot}=${[entry.itemId, entry.enchant, entry.suffix]
        .filter((part): part is number => part !== undefined)
        .join(':')}`;
    })
    .join(',');
}

/**
 * Version 2: the version 1 string, then the sections in the contract's fixed order, each
 * omitted when empty. Names are URL-encoded, so a set called "a;b=c|d" survives -- the
 * three characters the grammar itself uses are the three a player is most likely to type.
 *
 * `encodeFS1` is untouched and still produces a version 1 string: the planner's own share
 * links and the "Sim this build" URL are version 1 and there is nothing in them to carry.
 */
export function encodeFS1V2(build: FS1Build): string {
  const bags = build.bags ?? [];
  const bank = build.bank ?? [];
  const sets = build.sets ?? [];
  const loadouts = build.loadouts ?? [];
  const professions = build.professions ?? [];

  const sections: string[] = [];
  if (bags.length > 0) sections.push(`bags=${encodeItems(bags)}`);
  if (bank.length > 0) sections.push(`bank=${encodeItems(bank)}`);
  if (sets.length > 0) {
    sections.push(
      `sets=${sets.map((set) => `${encodeURIComponent(set.name)}=${encodeGearSlots(set.gear)}`).join(';')}`,
    );
  }
  if (loadouts.length > 0) {
    sections.push(
      `loadouts=${loadouts
        .map((loadout) => `${encodeURIComponent(loadout.name)}=${encodeTrees(loadout.treeRanks)}`)
        .join(';')}`,
    );
  }
  if (professions.length > 0) sections.push(`professions=${professions.join(',')}`);
  if (build.guild) sections.push(`guild=${encodeURIComponent(build.guild.name)}:${build.guild.rankIndex}`);

  // Always built from the slot list, never delegated to `encodeFS1` -- a caller can hold
  // `gearSlots` with nothing in `gear` (the doc comment on `FS1Build.gearSlots` invites
  // exactly that: "a caller holding nothing but a Gear map" implies the reverse is legal
  // too), and falling back to `encodeFS1(build)` for a "nothing enchanted" build used to
  // read `build.gear` instead and silently drop every item in that case. `encodeGearSlots`
  // produces byte-identical output to `encodeGearList` for a slot with no enchant or
  // suffix, so `encodeFS1V2(build) === encodeFS1(build)` still holds whenever `build.gear`
  // is the only source given.
  const slots = build.gearSlots ?? gearSlotsFrom(build.gear);
  const head = [
    FS1_PREFIX,
    build.dataBuild,
    build.classSlug,
    build.raceSlug,
    encodeTrees(build.treeRanks),
    encodeGearSlots(slots),
  ].join(':');

  return [head, ...sections].join('|');
}
