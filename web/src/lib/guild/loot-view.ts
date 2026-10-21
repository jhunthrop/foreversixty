// web/src/lib/guild/loot-view.ts
// Guild control-centre spec §4.F: the loot council helper. The API (contract's
// GET .../loot) already ranks candidates and names the awardee; this module only formats
// what it sent -- never re-ranks or re-decides an award client-side.
import { dayWord } from '../dates';
import { SLOT_LABELS, type Slot } from '../planner/types';
import type { GuildLootCandidate, GuildLootEncounter, GuildLootItem } from './api';

export const LOOT_FOOTNOTE =
  'Candidates are ranked by gain over what they wear now. A drop stays unawarded until an officer awards it.';
export const NOT_SIMMED_NOTE = 'Gain here is item level. The simulator has not checked these.';
export const NO_CANDIDATE = 'No one on the roster can use this.';
export const ALREADY_AWARDED = 'Already awarded';
export const UNDO_AWARD = 'Undo';

/** "Award Stormrage Cover to Rootbound?", what the inline confirm asks before an award is recorded. */
export function awardQuestion(itemName: string, candidateName: string): string {
  return `Award ${itemName} to ${candidateName}?`;
}

/** The boss picker's status word: a boss that is down reads "on farm", one that is not "not
 *  killed yet". Never the filter's own jargon. */
export function bossPickerLabel(selected: GuildLootEncounter | undefined): string {
  if (selected === undefined) return 'No bosses killed yet this tier';
  return selected.killed ? `${selected.name} · on farm` : `${selected.name} · not killed yet`;
}

/** "Head", "Main hand", "Finger": the slot a drop goes in, as the planner names it. */
function slotWord(slot: string): string {
  if (Object.prototype.hasOwnProperty.call(SLOT_LABELS, slot)) return SLOT_LABELS[slot as Slot];
  const plain = slot.replace(/_/g, ' ');
  return plain === '' ? plain : plain[0].toUpperCase() + plain.slice(1);
}

/** Whether any candidate on the page was ranked by item level rather than by a simulation. */
export function anyNotSimmed(items: readonly GuildLootItem[]): boolean {
  return items.some((item) => item.candidates.some((candidate) => candidate.not_sim_checked));
}

/** "Rootbound · +76 item levels over their current Head · 1 of 1 raid nights", the tail of a row. */
export function attendanceLabel(present: number, nights: number): string {
  return `${present} of ${nights} raid night${nights === 1 ? '' : 's'}`;
}

/** "Awarded to Rootbound · today · by Bow Jackzon": said once, on the item. */
export function awardedLine(
  awarded: { name: string; at: string; by_name: string },
  now: Date = new Date(),
): string {
  const day = dayWord(new Date(awarded.at), now);
  return [`Awarded to ${awarded.name}`, day, awarded.by_name === '' ? '' : `by ${awarded.by_name}`]
    .filter((part) => part !== '')
    .join(' · ');
}

/** The first encounter in the list the API has not marked killed -- the boss picker's own
 *  default (spec §4.F); `undefined` when every listed encounter is already killed, same
 *  meaning as `bossPickerLabel`'s "none, farm" branch. */
export function nextUnkilledEncounter(
  encounters: readonly GuildLootEncounter[],
): GuildLootEncounter | undefined {
  return encounters.find((encounter) => !encounter.killed);
}

/**
 * A candidate row's own gain figure (defect 5's own live-fix round): a tier-0 (BiS-
 * matched) candidate carries `gain_dps`, a tier-1 (fallback) one carries `ilvl_delta`
 * instead and `gain_dps: null` -- the contract's own two-tier shape
 * (docs/contracts/2026-10-04-guild-centre-api.md's loot section). Every real candidate on
 * a live roster is tier-1 today (no BiS file in this build names a raid-tier item yet),
 * so this is the common case, not an edge case: GuildLoot.svelte must never call a number
 * method on `gain_dps` directly, the bug that broke the whole tab on real data while the
 * fixture (predating the tier-1 fallback) never once exercised a null `gain_dps`.
 */
export function candidateGainLabel(candidate: GuildLootCandidate, slot?: string): string {
  if (candidate.gain_dps !== null) return `+${Math.round(candidate.gain_dps)} DPS`;
  if (candidate.ilvl_delta !== null) {
    const delta = candidate.ilvl_delta;
    const sign = delta >= 0 ? '+' : '';
    const figure = `${sign}${delta} item level${Math.abs(delta) === 1 ? '' : 's'}`;
    return slot === undefined || slot === '' ? figure : `${figure} over their current ${slotWord(slot)}`;
  }
  return '—';
}

export type LootCandidateAction =
  | { kind: 'awarded' }
  | { kind: 'awarded-to-other'; label: string }
  | { kind: 'equivalent' }
  | { kind: 'award' };

/**
 * One candidate row's own action cell (spec §12.1's round-2 fix): the awardee -- the
 * ranking's own top candidate, read off `item.awarded_to` -- carries the green "Awarded"
 * tag with no button; every other candidate on an awarded item shows a muted "Awarded to
 * {name}" note and no button; an item with no `already_equivalent` candidate and not yet
 * awarded keeps the `Award` button.
 */
export function candidateAction(item: GuildLootItem, candidateKey: string): LootCandidateAction {
  if (item.awarded_to !== null) {
    return item.awarded_to.character_key === candidateKey
      ? { kind: 'awarded' }
      : { kind: 'awarded-to-other', label: ALREADY_AWARDED };
  }
  const candidate = item.candidates.find((row) => row.character_key === candidateKey);
  if (candidate?.already_equivalent === true) return { kind: 'equivalent' };
  return { kind: 'award' };
}
