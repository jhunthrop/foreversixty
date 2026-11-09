// web/src/lib/bis/pre-raid.ts
// The level 60 list's Pre-raid toggle (owner 2026-11-08: "a simple toggle that just sets no raid in
// the filter"): a preset over the page's source filter (source-filter.ts) with Raid and PvP off
// (high-rank PvP gear is not pre-raid), the same kinds the ranker measures as `filtered_set_dps`
// `pre_raid` (sim/cmd/leveling-bis/filtered_sets.go). It needs no code change when raids open: it
// shows on any level 60 band with a pick had only from a raid or PvP (those slots re-pick from
// their `by_kind` entries), and its own page, /bis/<class>/<spec>/60/pre-raid, is canonical to /60
// for as long as the plain list has no such pick.
import { MAX_BAND } from './band-levels';
import { SOURCE_FILTER_KINDS, pickKinds, type FilterableSlot } from './source-filter';
import type { SourceFilterKind } from './types';

/** The kinds the toggle turns off. */
export const PRE_RAID_OFF: readonly SourceFilterKind[] = ['raid', 'pvp'];

/** The kinds on under the toggle, in the filter row's order. */
export const PRE_RAID_KINDS: readonly SourceFilterKind[] = SOURCE_FILTER_KINDS.filter(
  (kind) => !PRE_RAID_OFF.includes(kind),
);

/** The path segment after /60: /bis/warrior/fury/60/pre-raid. */
export const PRE_RAID_SLUG = 'pre-raid';

/** The toggle's words: a link on the summary line, where a lower band shows its next band. */
export const preRaidCopy = {
  turnOn: 'Pre-raid only',
  turnOff: 'Include raid and PvP gear',
} as const;

/** The ranker's `filtered_set_dps` key for the same kinds. */
export const PRE_RAID_SET_KEY = 'pre_raid';

/** Whether the kinds on are exactly the pre-raid kinds. */
export function isPreRaidKinds(on: readonly SourceFilterKind[]): boolean {
  return on.length === PRE_RAID_KINDS.length && PRE_RAID_KINDS.every((kind) => on.includes(kind));
}

/** Whether a band and its kinds are the pre-raid list (level 60, raid and PvP off). */
export function isPreRaidView(band: number, on: readonly SourceFilterKind[]): boolean {
  return band >= MAX_BAND && isPreRaidKinds(on);
}

/** Whether a pick turns over under the toggle: an item had only from a raid or PvP. */
function offList(slot: Pick<FilterableSlot, 'item_id' | 'kinds' | 'source_kind'>): boolean {
  if (typeof slot.item_id !== 'number' || slot.item_id === 0) return false;
  const kinds = pickKinds(slot);
  return kinds.length > 0 && kinds.every((kind) => PRE_RAID_OFF.includes(kind));
}

/** Whether a level 60 list is already pre-raid: no pick is had only from a raid or PvP. */
export function listIsPreRaid(
  slots: readonly Pick<FilterableSlot, 'item_id' | 'kinds' | 'source_kind'>[],
): boolean {
  return !slots.some(offList);
}

/** Whether a band offers the toggle: level 60, and the toggle would change a pick (a slot whose
 *  pick is a raid or PvP item re-picks from its `by_kind` entries). On a list that is already
 *  pre-raid it would change nothing, so it shows by itself the night raid gear enters the data. */
export function preRaidOffered(
  band: number,
  slots: readonly Pick<FilterableSlot, 'item_id' | 'kinds' | 'source_kind'>[],
): boolean {
  return band >= MAX_BAND && !listIsPreRaid(slots);
}
