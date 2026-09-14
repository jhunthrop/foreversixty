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
import { canAddPoint, withPoint, type TalentIndex } from './rules';
import { SLOTS, type Gear, type Slot } from './types';

export const FS1_PREFIX = 'FS1';
const TREES = 3;
// A query string is attacker-controlled. A real code tops out around a few hundred characters
// (build id, two slugs, three short talent strings, at most 17 gear entries); this is a
// generous multiple of that, high enough to never clip a real code and low enough that nothing
// past it is worth splitting or scanning.
const MAX_CODE_LENGTH = 2048;

export interface FS1Build {
  dataBuild: string;
  classSlug: string;
  raceSlug: string;
  /** One array per tree, one entry per talent in tab order. */
  treeRanks: number[][];
  gear: Gear;
}

export type FS1Error = { ok: false; message: string };
export type FS1Result = { ok: true; build: FS1Build } | FS1Error;

function encodeTree(ranks: number[]): string {
  const digits = ranks.map((rank) => Math.max(0, Math.min(35, Math.round(rank))).toString(36));
  const trimmed = digits.join('').replace(/0+$/, '');
  return trimmed === '' ? '0' : trimmed;
}

export function encodeFS1(build: FS1Build): string {
  const trees = Array.from({ length: TREES }, (_, index) => encodeTree(build.treeRanks[index] ?? []));
  const gear = SLOTS.filter((slot) => build.gear[slot] !== undefined)
    .map((slot) => `${slot}=${build.gear[slot]}`)
    .join(',');
  return [FS1_PREFIX, build.dataBuild, build.classSlug, build.raceSlug, trees.join('/'), gear].join(':');
}

export function decodeFS1(code: string): FS1Result {
  // Before any splitting or parsing: the cheapest possible check, and the one that keeps a
  // hostile multi-megabyte query value from doing any real work at all.
  if (code.length > MAX_CODE_LENGTH) {
    return { ok: false, message: 'That code is too long to read.' };
  }
  const parts = code.trim().split(':');
  if (parts[0] !== FS1_PREFIX) {
    const named = parts[0] === undefined || parts[0] === '' ? 'unlabelled' : parts[0];
    return { ok: false, message: `That code is ${named}; this site reads ${FS1_PREFIX}.` };
  }
  if (parts.length < 6) return { ok: false, message: 'That code is missing its talent and gear fields.' };

  const [, dataBuild, classSlug, raceSlug, treeField, ...gearParts] = parts;
  const treeStrings = treeField.split('/');
  if (treeStrings.length !== TREES) {
    return { ok: false, message: `That code has ${treeStrings.length} talent trees; a build has ${TREES}.` };
  }

  const treeRanks: number[][] = [];
  for (const tree of treeStrings) {
    const ranks: number[] = [];
    for (const digit of tree) {
      const rank = Number.parseInt(digit, 36);
      if (Number.isNaN(rank))
        return { ok: false, message: `That code has an unreadable talent rank: ${digit}.` };
      ranks.push(rank);
    }
    treeRanks.push(ranks);
  }

  const gear: Gear = {};
  const gearField = gearParts.join(':');
  if (gearField !== '') {
    for (const entry of gearField.split(',')) {
      const [slot, value] = entry.split('=');
      if (!(SLOTS as readonly string[]).includes(slot)) {
        return { ok: false, message: `That code names a slot this planner does not have: ${slot}.` };
      }
      // A strict digits-only match rather than Number.parseInt: parseInt stops at the first
      // non-digit character and returns what came before it, so "12640abc" and "12640.5" would
      // otherwise silently become the item id 12640 instead of being refused.
      if (value === undefined || !/^\d+$/.test(value)) {
        return { ok: false, message: `That code has an unreadable gear entry: ${entry}.` };
      }
      gear[slot as Slot] = Number.parseInt(value, 10);
    }
  }

  return { ok: true, build: { dataBuild, classSlug, raceSlug, treeRanks, gear } };
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
