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
import { PINNED_STATS } from '../sim/stats';

export const FSB1_PREFIX = 'FSB1';
/** The same bound `fs1.ts` uses; checked before any splitting. */
export const MAX_CODE_LENGTH = 16_384;

const DIGITS = '0123456789abcdefghijklmnopqrstuvwxyz';
const SLOT_SET = new Set<string>(SLOTS);
/** Contract 10.8's stat vocabulary order, verbatim (sim/stats.ts's `PINNED_STATS`, not the
 *  generated `SIM_STATS`, which is alphabetised for the weights page and not the wire
 *  order): `stamina` sorts before `spell_power` because the pinned enum does, and the
 *  shared fixture vectors -- and the addon's own encoder -- are generated in that order. */
const STAT_ORDER = new Map(PINNED_STATS.map((name, index) => [name, index]));

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

function toBase36(value: number): string {
  return DIGITS[Math.min(35, Math.max(0, Math.round(value)))];
}

function fromBase36(char: string): number | undefined {
  const index = DIGITS.indexOf(char.toLowerCase());
  return index === -1 ? undefined : index;
}

/** `PINNED_STATS`'s index for a known name; every unrecognised name shares this rank, so
 *  two of them fall through to the alphabetical tie-break below rather than comparing
 *  `Infinity - Infinity` (`NaN`, which is neither `< 0`, `> 0` nor `0` and leaves `sort`'s
 *  ordering for that pair undefined). */
function statRank(name: string): number {
  return STAT_ORDER.get(name) ?? Number.POSITIVE_INFINITY;
}

function encodeStats(stats: Record<string, number>): string {
  // Sorted, so one build is one string: an object's key order is insertion order and two
  // callers building the same set of stats in different orders would otherwise produce
  // two codes for one set. Sorted by the pinned vocabulary's own order (STAT_ORDER), with
  // an unrecognised name (outside contract 10.8) falling after every known one, tied
  // alphabetically among themselves.
  return Object.keys(stats)
    .sort((a, b) => {
      const [rankA, rankB] = [statRank(a), statRank(b)];
      return rankA !== rankB ? rankA - rankB : a.localeCompare(b);
    })
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
    // Digits only, never Number.parseInt on the field: parseInt stops at the first
    // non-digit and would turn "lots" into NaN and "30x" into 30.
    if (name === '' || !/^\d+$/.test(value)) return { ok: false, message: addonCopy.statPair(pair) };
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
  const order = parseOrder(parts[3]);
  if (!order.ok) return order;
  const gear = parseGear(parts.slice(4).join(':'));
  if (!gear.ok) return gear;
  return {
    ok: true,
    build: { dataBuild: parts[1], classSlug: parts[2], order: order.value, gear: gear.value },
  };
}
