// web/src/lib/addon/fsb1.ts
// FSB1, the site-to-addon build code.
//
//   FSB1:<data-build>:<class-slug>:<order>:<gear>
//
//   <order>  one triple per point spent -- tab, tier, column -- each a single base-36
//            character, all 1-based, concatenated with no separator. 1-based because
//            GetTalentInfo is, so the addon compares what the client hands it with no
//            arithmetic in two places; a leading 0 is therefore a malformed code rather
//            than a silently wrong tree.
//   <gear>   entries joined by ",", each <slot>=<item_id>[:<stat>=<value>[;<stat>=<value>]...]
//            The stat names are the engine's own vocabulary (parity contract 10.8), which
//            is what Data.lua's weights are keyed by, so a planned item and a weight meet
//            without a translation table.
//
// FS1 (planner/fs1.ts) stays the character export and is unchanged. This is the other
// direction, and it is a separate format because FS1 carries neither a point order nor an
// item's stats -- the two things Follow and Gear need.
//
// addon/ForeverSixty/Codec.lua implements the same grammar against the same fixture
// vectors (web/src/fixtures/addon/codec-vectors.json), so neither side may drift.
import { addonCopy } from './copy';
import { SLOTS, type Slot } from '../planner/types';

export const FSB1_PREFIX = 'FSB1';
/** The same bound `fs1.ts` uses; checked before any splitting. */
export const MAX_CODE_LENGTH = 16_384;

const DIGITS = '0123456789abcdefghijklmnopqrstuvwxyz';
const SLOT_SET = new Set<string>(SLOTS);
/** A stat value on the wire: an optional minus, then a decimal with no redundant zeros. */
const CANONICAL_INTEGER = /^(?:0|-?[1-9]\d*)$/;

export interface FSB1Point {
  tab: number;
  tier: number;
  column: number;
}

export interface FSB1GearSlot {
  slot: Slot;
  itemId: number;
  /** Contract 10.8 stat names to amounts. Empty when the planner knows none. */
  stats: Record<string, number>;
}

export interface FSB1Build {
  dataBuild: string;
  classSlug: string;
  order: FSB1Point[];
  gear: FSB1GearSlot[];
}

export type FSB1Result = { ok: true; build: FSB1Build } | { ok: false; message: string };

/**
 * One base-36 digit. The clamp is a last resort, not a range policy: the only caller is
 * `build-code.ts`'s `pointsOf`, whose tab/tier/column all come out of an indexed talent
 * file (at most 3 tabs, 9 tiers, 4 columns, every one of them 1-based), so nothing in the
 * app can reach it. It clamps rather than throws because `encodeFSB1` is called from a
 * `$derived` in the share panel, where a throw would take the whole planner down for a
 * data defect that costs at worst one wrong character. A clamped 0 is refused by
 * `parseOrder` on the far side; a clamped 35 is a cell no real tree has. If a future
 * caller can legitimately exceed the range, it must widen the encoding rather than lean
 * on this.
 */
function toBase36(value: number): string {
  return DIGITS[Math.min(35, Math.max(0, Math.round(value)))];
}

function fromBase36(char: string): number | undefined {
  const index = DIGITS.indexOf(char.toLowerCase());
  return index === -1 ? undefined : index;
}

function encodeStats(stats: Record<string, number>): string {
  // Sorted, so one build is one string: an object's key order is insertion order, and two
  // callers building the same set of stats in different orders would otherwise produce two
  // codes for one set.
  //
  // Sorted by name, not by contract 10.8's `PINNED_STATS` vocabulary order. Both were
  // defensible in isolation and this file used to do the latter; the addon lane owns the
  // shared fixture and its Codec.lua `encodeStats` is a plain `table.sort(names)`, so name
  // order is what the two sides actually have to agree on. It is also what the format's own
  // spec snippet (`Object.keys(stats).sort()`) always said. A vocabulary order would need a
  // table shared across a TypeScript bundle and a WoW addon zip and kept in step by hand;
  // name order needs nothing shared at all, which is the whole argument for it.
  //
  // Byte order, never `localeCompare`. With name order this is the only comparator rather
  // than a tie-break, so the difference is now load-bearing: collation folds case
  // (`'B'.localeCompare('a')` is 1 where `'B' < 'a'` is true) and is locale- and
  // environment-dependent, so two browsers could emit two codes for one build -- the exact
  // invariant this sort exists to hold. Lua's `table.sort` on strings compares bytes, so
  // byte order is also what the parallel Codec.lua encoder does.
  return Object.keys(stats)
    .sort((a, b) => (a < b ? -1 : a > b ? 1 : 0))
    .map((name) => `${name}=${Math.round(stats[name])}`)
    .join(';');
}

export function encodeFSB1(build: FSB1Build): string {
  const order = build.order
    .map((point) => `${toBase36(point.tab)}${toBase36(point.tier)}${toBase36(point.column)}`)
    .join('');
  const gear = build.gear
    .map((entry) => {
      const stats = encodeStats(entry.stats);
      return stats === '' ? `${entry.slot}=${entry.itemId}` : `${entry.slot}=${entry.itemId}:${stats}`;
    })
    .join(',');
  return [FSB1_PREFIX, build.dataBuild, build.classSlug, order, gear].join(':');
}

type Parsed<T> = { ok: true; value: T } | { ok: false; message: string };

function parseOrder(field: string): Parsed<FSB1Point[]> {
  if (field === '') return { ok: true, value: [] };
  if (field.length % 3 !== 0) return { ok: false, message: addonCopy.orderLength };
  const order: FSB1Point[] = [];
  for (let at = 0; at < field.length; at += 3) {
    const triple = field.slice(at, at + 3);
    const [tab, tier, column] = [...triple].map(fromBase36);
    if (
      tab === undefined ||
      tier === undefined ||
      column === undefined ||
      tab < 1 ||
      tier < 1 ||
      column < 1
    ) {
      return { ok: false, message: addonCopy.orderCell(triple) };
    }
    order.push({ tab, tier, column });
  }
  return { ok: true, value: order };
}

// The brief's plan signature also took the full gear `entry` for the error message, but
// every specified refusal names the stat pair (`addonCopy.statPair(pair)`), never the
// entry -- so the parameter was dead weight. Dropped rather than kept with `void entry;`;
// no specified message or behaviour changes.
function parseStats(field: string): Parsed<Record<string, number>> {
  const stats: Record<string, number> = {};
  if (field === '') return { ok: true, value: stats };
  for (const pair of field.split(';')) {
    const at = pair.indexOf('=');
    if (at === -1) return { ok: false, message: addonCopy.statPair(pair) };
    const name = pair.slice(0, at);
    const value = pair.slice(at + 1);
    // A canonical decimal integer, never Number.parseInt on the raw field: parseInt stops
    // at the first non-digit and would turn "lots" into NaN and "30x" into 30.
    //
    // The sign is part of the grammar because real items carry negative stats -- 43 values
    // across data/builds/<build>/items/*.json, Ring of Scorn's spirit -3 among them. The
    // encoder copies an item's stats through verbatim, so a decoder that refused a minus
    // sign would refuse this site's own "Copy addon code" output; and dropping the penalty
    // instead would make the addon score a cursed item as though it had none.
    //
    // Canonical, so one build is still one string: "-0", "+3", "007" and "--3" are all
    // refused, because `encodeStats` emits none of them and two spellings of one number
    // would be two codes for one build.
    if (name === '' || !CANONICAL_INTEGER.test(value)) {
      return { ok: false, message: addonCopy.statPair(pair) };
    }
    stats[name] = Number.parseInt(value, 10);
  }
  return { ok: true, value: stats };
}

function parseGear(field: string): Parsed<FSB1GearSlot[]> {
  if (field === '') return { ok: true, value: [] };
  const gear: FSB1GearSlot[] = [];
  for (const entry of field.split(',')) {
    const at = entry.indexOf('=');
    if (at === -1) return { ok: false, message: addonCopy.gearEntry(entry) };
    const slot = entry.slice(0, at);
    if (!SLOT_SET.has(slot)) return { ok: false, message: addonCopy.unknownSlot(slot) };
    const [id, ...rest] = entry.slice(at + 1).split(':');
    if (!/^\d+$/.test(id)) return { ok: false, message: addonCopy.gearEntry(entry) };
    const stats = parseStats(rest.join(':'));
    if (!stats.ok) return stats;
    gear.push({ slot: slot as Slot, itemId: Number.parseInt(id, 10), stats: stats.value });
  }
  return { ok: true, value: gear };
}

export function decodeFSB1(code: string): FSB1Result {
  if (code.length > MAX_CODE_LENGTH) return { ok: false, message: addonCopy.tooLong };
  const parts = code.trim().split(':');
  if (parts[0] !== FSB1_PREFIX) {
    const named = parts[0] === undefined || parts[0] === '' ? addonCopy.unlabelledCode : parts[0];
    return { ok: false, message: addonCopy.wrongPrefix(named, FSB1_PREFIX) };
  }
  if (parts.length < 5) return { ok: false, message: addonCopy.shortCode };
  // `parts.length < 5` catches a field that is missing; these two catch one that is present
  // but empty, which is the same defect with a colon in front of it. `order` and `gear` are
  // both legitimately empty (a fresh build has no points and no gear); the data build and
  // the class are not -- the addon can neither check staleness nor find a tree without them.
  if (parts[1] === '') return { ok: false, message: addonCopy.emptyField(addonCopy.dataBuildField) };
  if (parts[2] === '') return { ok: false, message: addonCopy.emptyField(addonCopy.classField) };
  const order = parseOrder(parts[3]);
  if (!order.ok) return order;
  const gear = parseGear(parts.slice(4).join(':'));
  if (!gear.ok) return gear;
  return {
    ok: true,
    build: { dataBuild: parts[1], classSlug: parts[2], order: order.value, gear: gear.value },
  };
}
