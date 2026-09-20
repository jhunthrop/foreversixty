// web/src/lib/sim/combos.ts
// Reading a ranked bulk result. Nothing here is a statistic: `group` already says which
// runs are within error of the leader, `delta` already carries its own error, and the
// ordering is the planner's. This turns those into rows, labels and a winning gear list.
import type { BulkResult, Combo, Substitution } from './bulk-types';
import { bulkCopy } from './copy';
import type { Estimate, GearSlot } from './types';
import type { Item, ItemSet } from '../planner/types';

export interface ComboRow {
  /** 1-based, and shared by every member of a within-error group (design 3.3). */
  rank: number;
  combo: Combo;
  /** True for the leader's group, which is the one the page draws a rule under. */
  withinError: boolean;
  percent: number;
}

export function percentOf(delta: number, equippedMean: number): number {
  return equippedMean === 0 ? 0 : (delta / equippedMean) * 100;
}

const MINUS = '−';

/**
 * "+41 ± 11": the gain, rounded, and its 95% band. A minus sign, not a hyphen -- the
 * design system's rule for a negative figure in a table.
 */
export function deltaLabel(delta: Estimate): string {
  const mean = Math.round(delta.mean);
  const band = Math.round(1.96 * delta.error);
  const sign = mean < 0 ? MINUS : '+';
  return `${sign}${Math.abs(mean).toLocaleString('en-US')} ± ${band.toLocaleString('en-US')}`;
}

/**
 * Rank shared across a group, per design 3.3: "Rows in the leader's within-error group
 * carry the same rank." A group's rank is the 1-based position of its first member, so the
 * run after a two-member tie is rank 3.
 */
export function comboRows(result: BulkResult): ComboRow[] {
  const firstOfGroup = new Map<number, number>();
  return result.combos.map((combo, index) => {
    if (!firstOfGroup.has(combo.group)) firstOfGroup.set(combo.group, index + 1);
    return {
      rank: firstOfGroup.get(combo.group)!,
      combo,
      withinError: combo.group === 0,
      percent: percentOf(combo.delta.mean, result.equipped.mean),
    };
  });
}

/** The leader's substitutions written over the base character's gear. */
export function winningGear(result: BulkResult): GearSlot[] {
  const gear = result.request.character.gear.map((slot) => ({ ...slot }));
  for (const sub of result.combos[0]?.substitutions ?? []) {
    if (sub.kind !== 'item' || sub.slot === undefined || sub.item_id === undefined) continue;
    const next: GearSlot = { slot: sub.slot, item_id: sub.item_id };
    if (sub.enchant !== undefined && sub.enchant > 0) next.enchant = sub.enchant;
    if (sub.suffix !== undefined && sub.suffix > 0) next.suffix = sub.suffix;
    const at = gear.findIndex((slot) => slot.slot === sub.slot);
    if (at >= 0) gear[at] = next;
    else gear.push(next);
  }
  return gear;
}

export interface SlotSummaryRow {
  slot: string;
  item_id: number;
  /** The item's own name, off the substitution (contract 10.1 A6). Never a re-join. */
  name: string;
  /**
   * What this slot contributed on its own, from the combination that changed only this
   * slot. Null when no such combination exists -- a Droptimizer run has one per slot by
   * construction, a Top Gear run has one whenever the slot's item was also tried alone, and
   * a run where it was not is honest about not knowing rather than apportioning the total.
   */
  gain: number | null;
}

/**
 * Engine-lane rule 5's sentinel: a two-hander replacing a main-plus-off-hand pair emits a
 * second substitution `{kind:"item", slot:"off_hand", item_id:0, name:"<item removed>"}`.
 * `item_id: 0` is never a real item (simdb has no id 0) -- the same check
 * `SubstitutionChips.svelte` makes before its own generic chip branch would look it up.
 */
function isEmptiedOffHand(sub: Substitution): boolean {
  return sub.kind === 'item' && sub.slot === 'off_hand' && sub.item_id === 0;
}

/** Design 3.3's per-slot summary: what the winner uses, and what that slot was worth. */
export function slotSummary(result: BulkResult): SlotSummaryRow[] {
  const winner = result.combos[0];
  if (winner === undefined) return [];
  const alone = new Map<string, number>();
  for (const combo of result.combos) {
    if (combo.substitutions.length !== 1) continue;
    const only = combo.substitutions[0];
    if (only.kind !== 'item' || only.slot === undefined) continue;
    alone.set(`${only.slot}:${only.item_id}`, combo.delta.mean);
  }
  return winner.substitutions
    .filter(
      (sub): sub is Substitution & { slot: string; item_id: number } =>
        sub.kind === 'item' && sub.slot !== undefined && sub.item_id !== undefined,
    )
    .map((sub) => ({
      slot: sub.slot,
      item_id: sub.item_id,
      // Engine-lane rule 5: the emptied off-hand is not an item, and "<item removed>" is
      // not a name a player should ever read verbatim -- the same fix SubstitutionChips
      // already applies, so every consumer of slotSummary (this "By slot" panel included)
      // inherits it from the one place the sentinel is read.
      name: isEmptiedOffHand(sub) ? bulkCopy.offHandEmptied : substitutionLabel(sub),
      gain: alone.get(`${sub.slot}:${sub.item_id}`) ?? null,
    }));
}

/**
 * Design 4.4's results filter: "only combos keeping 4-piece". Tier bonuses are counted by
 * the engine from the gear itself, so this is a filter over what a combination would be
 * wearing, never an input to the run.
 */
export function keepsSetBonus(
  combo: Combo,
  result: BulkResult,
  items: ReadonlyMap<number, Item>,
  sets: readonly ItemSet[],
  pieces: number,
): boolean {
  const gear = new Map(result.request.character.gear.map((slot) => [slot.slot, slot.item_id]));
  for (const sub of combo.substitutions) {
    if (sub.kind !== 'item' || sub.slot === undefined || sub.item_id === undefined) continue;
    gear.set(sub.slot, sub.item_id);
  }
  const counts = new Map<number, number>();
  for (const itemId of gear.values()) {
    const setId = items.get(itemId)?.set_id ?? null;
    if (setId === null) continue;
    counts.set(setId, (counts.get(setId) ?? 0) + 1);
  }
  return sets.some((set) => (counts.get(set.id) ?? 0) >= pieces);
}

/**
 * "Helm of Wrath", "Deep Fury", "AQ set" -- whatever this substitution actually changed.
 *
 * It reads `name`, which contract 10.1 A6 fills for items too, from simdb. There is no
 * item-id join here and no item map parameter: the result already carries every name it
 * needs, and re-deriving one from a per-class item file the reader may not have loaded is
 * how a saved page ends up showing "Item 16963" for something the engine named.
 */
export function substitutionLabel(sub: Substitution): string {
  // Contract 10.8: a `consumes` substitution's name is the ids joined by ", ", which is a
  // list and not a phrase, so it is the one kind that goes through copy on its way out.
  if (sub.kind === 'consumes') return bulkCopy.consumesChip(sub.name ?? '');
  if (sub.name !== undefined && sub.name !== '') return sub.name;
  return sub.kind === 'item' ? `Item ${sub.item_id ?? 0}` : '';
}

/** Where a Droptimizer combination's one substitution came from, in words. */
export function sourceNameOfCombo(combo: Combo): string {
  return combo.substitutions[0]?.source_name ?? '';
}

/**
 * "+41 DPS from Helm of Wrath", for the page's own display. The leader's first substitution
 * is the one named -- the planner orders a combination's substitutions by the slot's own
 * contribution, so the first is the one that carried it.
 *
 * The *stored* headline on a saved sim is not this: contract 10.6 has the API compose it
 * at save time, with its own rules for several substitutions ("… and 2 more") and for an
 * empty result ("no combinations"). `SimListRow.headline` is read, never recomputed.
 */
export function headlineFor(result: BulkResult): string {
  const winner = result.combos[0];
  if (winner === undefined) return '';
  const gain = deltaLabel(winner.delta).split(' ')[0];
  const what = substitutionLabel(winner.substitutions[0]);
  return what === '' ? `${gain} DPS` : `${gain} DPS from ${what}`;
}
