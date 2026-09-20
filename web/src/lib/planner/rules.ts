// web/src/lib/planner/rules.ts
// Client-side mirror of the API's build validation (rules 1-6 in
// docs/superpowers/specs/2026-09-13-phase-1-interfaces.md), so the planner refuses exactly
// the moves POST /v1/builds refuses, with the same wording.
//
// Rules 2-5 are expressed once, in validateOrder. canAddPoint and canRemovePoint build the
// candidate order and ask validateOrder about it, so an added point and a saved build can
// never disagree about what is legal.
import {
  MAX_POINTS,
  POINTS_PER_TIER,
  SLOTS,
  SLOT_ALIASES,
  SLOT_LABELS,
  type Combo,
  type Gear,
  type Item,
  type ItemFile,
  type Slot,
  type Talent,
  type TalentFile,
  type TalentTree,
} from './types';

export interface TalentIndex {
  file: TalentFile;
  /** Trees in left-to-right (`position`) order. */
  trees: TalentTree[];
  byId: Map<number, Talent>;
  treeOf: Map<number, TalentTree>;
}

export interface FieldError {
  /** `point_order[7]` or `gear.head`, matching the API's 400 response. */
  field: string;
  message: string;
}

export type Decision = { ok: true } | { ok: false; reason: string };

const ALLOWED: Decision = { ok: true };
const refuse = (reason: string): Decision => ({ ok: false, reason });

/** Every refusal string in the planner. Components and tests import these, never literals. */
export const messages = {
  tierLocked: (tier: number, treeName: string): string =>
    `Tier ${tier} of ${treeName} needs ${POINTS_PER_TIER * tier} points in ${treeName} first`,
  prereqMissing: (talentName: string, rank: number, prereqName: string): string =>
    `${talentName} needs ${rank} ${rank === 1 ? 'point' : 'points'} in ${prereqName} first`,
  maxRank: (talentName: string, maxRank: number): string =>
    `${talentName} is already at ${maxRank} of ${maxRank} points`,
  capReached: (): string => `A build spends at most ${MAX_POINTS} points`,
  unknownTalent: (id: number): string => `Talent ${id} is not in this class`,
  noPoints: (talentName: string): string => `${talentName} has no points to remove`,
  unknownSlot: (slot: string): string => `${slot} is not a gear slot`,
  unknownItem: (id: number): string => `Item ${id} is not in this class list`,
  wrongSlot: (itemName: string, slotLabel: string): string => `${itemName} cannot go in ${slotLabel}`,
  duplicateUnique: (itemName: string): string => `${itemName} is unique; equip it once`,
  illegalCombo: (raceName: string, className: string): string => `${raceName} cannot be a ${className}`,
  twoHandOffHand: (mainHandName: string): string =>
    `${mainHandName} is two-handed; there is no room for an off-hand item`,
} as const;

export function indexTalents(file: TalentFile): TalentIndex {
  const trees = [...file.trees].sort((a, b) => a.position - b.position);
  const byId = new Map<number, Talent>();
  const treeOf = new Map<number, TalentTree>();
  for (const tree of trees) {
    for (const talent of tree.talents) {
      byId.set(talent.id, talent);
      treeOf.set(talent.id, tree);
    }
  }
  return { file, trees, byId, treeOf };
}

export function indexItems(file: ItemFile): Map<number, Item> {
  return new Map(file.items.map((item) => [item.id, item]));
}

/** Rules 2 to 5, applied point by point. Returns one message per offending index. */
export function validateOrder(index: TalentIndex, order: number[]): FieldError[] {
  const errors: FieldError[] = [];
  const pointsInTree = new Map<number, number>();
  const ranks = new Map<number, number>();

  order.forEach((id, i) => {
    const field = `point_order[${i}]`;
    const talent = index.byId.get(id);
    if (!talent) {
      errors.push({ field, message: messages.unknownTalent(id) });
      return;
    }
    const tree = index.treeOf.get(id)!;
    const inTree = pointsInTree.get(tree.id) ?? 0;
    const rank = ranks.get(id) ?? 0;

    if (inTree < POINTS_PER_TIER * talent.tier) {
      errors.push({ field, message: messages.tierLocked(talent.tier, tree.name) });
    } else if (
      talent.prereq_talent_id !== null &&
      (ranks.get(talent.prereq_talent_id) ?? 0) < (talent.prereq_rank ?? 0)
    ) {
      const prereq = index.byId.get(talent.prereq_talent_id);
      errors.push({
        field,
        message: messages.prereqMissing(
          talent.name,
          talent.prereq_rank ?? 0,
          prereq?.name ?? String(talent.prereq_talent_id),
        ),
      });
    } else if (rank >= talent.max_rank) {
      errors.push({ field, message: messages.maxRank(talent.name, talent.max_rank) });
    } else if (i >= MAX_POINTS) {
      errors.push({ field, message: messages.capReached() });
    }

    // Both counters advance even when the point above was just refused: a refused
    // point still occupies a slot in point_order, and the only branch that skips this
    // update is the unknown-talent early return. This can make a later index's checks
    // pass (or fail differently) than they would if only legal points were counted, but
    // it never lets an illegal order through — any single FieldError already invalidates
    // the whole order, so the effect is limited to which secondary diagnostic is
    // reported. The API's validator must replay counts the same way, or the two would
    // disagree about which message to show for a given index.
    pointsInTree.set(tree.id, inTree + 1);
    ranks.set(id, rank + 1);
  });

  return errors;
}

/** A new order with one more point in `talentId`. Never mutates. */
export function withPoint(order: number[], talentId: number): number[] {
  return [...order, talentId];
}

/** A new order with the last point in `talentId` dropped. Never mutates. */
export function withoutLastPoint(order: number[], talentId: number): number[] {
  const last = order.lastIndexOf(talentId);
  if (last === -1) return [...order];
  return [...order.slice(0, last), ...order.slice(last + 1)];
}

export function canAddPoint(index: TalentIndex, order: number[], talentId: number): Decision {
  const candidate = withPoint(order, talentId);
  const field = `point_order[${candidate.length - 1}]`;
  const failure = validateOrder(index, candidate).find((error) => error.field === field);
  return failure ? refuse(failure.message) : ALLOWED;
}

/** The index a `point_order[n]` field name refers to. */
function fieldIndex(field: string): number {
  return Number(field.slice('point_order['.length, -1));
}

/**
 * A talent in `talentId`'s tree with a point spent on it in `order` that requires
 * more rank in `talentId` than `newRank` would leave it. `validateOrder` cannot surface
 * this on its own: rule 4 (prerequisite rank) is checked as an `else if` behind
 * rule 3 (tier lock) for the *dependant's* own point, so if that point also happens
 * to be tier-locked -- true whenever the candidate order is missing points
 * elsewhere in the tree, as a partial or synthetic order can be -- the tier-lock
 * message masks the prerequisite one entirely. Checking the dependant's rank
 * requirement directly, against the prerequisite's rank rather than replaying the
 * whole order, sidesteps that masking.
 */
function dependantNeedingMoreRank(
  index: TalentIndex,
  order: number[],
  talentId: number,
  newRank: number,
): Talent | undefined {
  const tree = index.treeOf.get(talentId);
  return tree?.talents.find(
    (talent) =>
      talent.prereq_talent_id === talentId &&
      order.includes(talent.id) &&
      (talent.prereq_rank ?? 0) > newRank,
  );
}

/**
 * Removing a point can only be refused for a violation the removal itself causes.
 * An order the caller hands in may already fail some other rule for reasons that
 * have nothing to do with the point being removed (the UI never lets that happen
 * in practice, but the check must not pretend a pre-existing problem is this
 * removal's fault). So, beyond the direct prerequisite-rank check above, every
 * other violation in the post-removal order is compared against the same point's
 * violation, if any, in the order before removal -- matched by carrying the shift
 * the removed slot introduces -- and only a violation that is new is grounds for
 * refusal.
 */
export function canRemovePoint(index: TalentIndex, order: number[], talentId: number): Decision {
  if (!order.includes(talentId)) {
    const talent = index.byId.get(talentId);
    return refuse(talent ? messages.noPoints(talent.name) : messages.unknownTalent(talentId));
  }
  const talent = index.byId.get(talentId);
  const newRank = order.filter((id) => id === talentId).length - 1;
  const dependant = talent && dependantNeedingMoreRank(index, order, talentId, newRank);
  if (dependant && talent) {
    return refuse(messages.prereqMissing(dependant.name, dependant.prereq_rank ?? 0, talent.name));
  }

  const removedAt = order.lastIndexOf(talentId);
  const before = new Map(
    validateOrder(index, order).map((error) => [fieldIndex(error.field), error.message]),
  );
  const after = validateOrder(index, withoutLastPoint(order, talentId));
  const newViolation = after.find((error) => {
    const beforeIndex =
      fieldIndex(error.field) < removedAt ? fieldIndex(error.field) : fieldIndex(error.field) + 1;
    return before.get(beforeIndex) !== error.message;
  });
  return newViolation ? refuse(newViolation.message) : ALLOWED;
}

/** Rule 1. */
export function comboIsLegal(combos: Combo[], raceId: number, classId: number): boolean {
  return combos.some((combo) => combo.race_id === raceId && combo.class_id === classId);
}

/** The planner slots an item can occupy; `finger`/`trinket` map to both numbered slots. */
export function slotsForItem(item: Item): Slot[] {
  const aliased = SLOT_ALIASES[item.slot];
  if (aliased) return aliased;
  return (SLOTS as readonly string[]).includes(item.slot) ? [item.slot as Slot] : [];
}

/** Rule 6, for one slot. */
export function canEquip(items: Map<number, Item>, gear: Gear, slot: Slot, itemId: number): Decision {
  if (!(SLOTS as readonly string[]).includes(slot)) return refuse(messages.unknownSlot(slot));
  const item = items.get(itemId);
  if (!item) return refuse(messages.unknownItem(itemId));
  if (!slotsForItem(item).includes(slot)) return refuse(messages.wrongSlot(item.name, SLOT_LABELS[slot]));
  if (item.unique) {
    const elsewhere = (Object.entries(gear) as [Slot, number][]).some(
      ([other, id]) => other !== slot && id === itemId,
    );
    if (elsewhere) return refuse(messages.duplicateUnique(item.name));
  }
  return ALLOWED;
}

/** Rule 6, for a whole gear map. */
export function validateGear(items: Map<number, Item>, gear: Gear): FieldError[] {
  const errors: FieldError[] = [];
  for (const [slot, itemId] of Object.entries(gear) as [Slot, number][]) {
    if (itemId === undefined) continue;
    const rest: Gear = { ...gear };
    delete rest[slot];
    const decision = canEquip(items, rest, slot, itemId);
    if (!decision.ok) errors.push({ field: `gear.${slot}`, message: decision.reason });
  }

  // Rule 6, the client mirror: an off-hand item with a two-handed main hand is refused.
  // The API states the same rule against the same `two_hand` field; this is the copy the
  // planner can enforce before a save round-trip.
  const mainHand = gear.main_hand === undefined ? undefined : items.get(gear.main_hand);
  if (mainHand?.two_hand === true && gear.off_hand !== undefined) {
    errors.push({ field: 'gear.off_hand', message: messages.twoHandOffHand(mainHand.name) });
  }

  return errors;
}
