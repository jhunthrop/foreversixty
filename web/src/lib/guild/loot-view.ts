// web/src/lib/guild/loot-view.ts
// Guild control-centre spec §4.F: the loot council helper. The API (contract's
// GET .../loot) already ranks candidates and names the awardee; this module only formats
// what it sent -- never re-ranks or re-decides an award client-side.
import { dayWord } from '../dates';
import { raidsAttended } from './attendance-copy';
import { specClassName } from './spec-name';
import type { GuildLootCandidate, GuildLootEncounter, GuildLootItem, GuildLootPage } from './api';

export const LOOT_FOOTNOTE =
  'Candidates are characters who can wear the item at their level. Each is ranked by how much it adds over what they wear now. A drop stays open until an officer awards it.';
export const TIER_ONE_LABEL = 'Ranked by item level only. Not on their best-in-slot list.';
/** The tier-1 group's label on every card after the first that says it in full (guild round 2). */
export const TIER_ONE_SHORT = 'Item level only';
export const NOBODY_AT_60 = "Onyxia's drops need level 60. Candidates show here as your raiders reach it.";
export const NO_CANDIDATES_AT_60 = 'Nobody at 60 can use these drops yet.';
export const LOOT_LOAD_ERROR = 'Loot did not load.';
export const EMPTY_SLOT = 'Empty slot';
/** A tier-1 gain cell when the candidate wears nothing there: the item-level gap would only be the
 *  item's own level (guild round 3). */
export const FILLS_EMPTY_SLOT = 'Fills an empty slot';
export const ON_BIS_LIST = 'On their BiS list';
export const ALREADY_AWARDED = 'Already awarded';
export const UNDO_AWARD = 'Undo';

/** A no-break space: "by Bow Jackzon" and "+16 item levels" never break inside. */
const NBSP = '\u00a0';

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

/** The drops the council ranks: a slotted drop with at least one candidate, in the boss's order. */
export function rankedDrops(items: readonly GuildLootItem[]): GuildLootItem[] {
  return items.filter((item) => item.kind === 'gear' && item.candidates.length > 0);
}

/** Everything else: quest drops, the bag, the skinned scale and a slotted drop nobody can use. */
export function unrankedDrops(items: readonly GuildLootItem[]): GuildLootItem[] {
  return items.filter((item) => item.kind !== 'gear' || item.candidates.length === 0);
}

/** The collapsed line's words. */
export function unrankedLine(count: number): string {
  return `${count} more ${count === 1 ? 'drop' : 'drops'} nobody here ranks for`;
}

/** What an open collapsed row says about its drop. */
export function unrankedLabel(item: GuildLootItem): string {
  switch (item.kind) {
    case 'quest':
      return item.quest_label;
    case 'bag':
      return 'Bag. Anyone can use it.';
    case 'skinned':
      return 'Skinned from Onyxia. Not a drop.';
    default:
      return 'Nobody here can use it yet.';
  }
}

/** Attendance shows once the guild has a counted raid night: no row with one, no cell at all. */
export function showsAttendance(items: readonly GuildLootItem[]): boolean {
  return items.some((item) => item.candidates.some((candidate) => candidate.attendance.nights > 0));
}

/** "3 of 4 raid nights", the attendance cell: the one attendance figure (attendance-copy). */
export const attendanceLabel = raidsAttended;

/** "Awarded to Rootbound · today · by Bow Jackzon": said once, on the item. */
export function awardedLine(
  awarded: { name: string; at: string; by_name: string },
  now: Date = new Date(),
): string {
  const day = dayWord(new Date(awarded.at), now);
  return [`Awarded to ${awarded.name}`, day, awarded.by_name === '' ? '' : `by${NBSP}${awarded.by_name}`]
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
 * A candidate's gain cell. Tier 0 reads the sim's gain in the BiS file's own unit (`+12.4 DPS`, a
 * healer's `HPS`, a tank's `TPS`), or `On their BiS list` when no figure was measured. Tier 1 reads
 * the item-level gap (`+14 item levels`), or `Fills an empty slot` when the candidate wears nothing
 * there (the gap would only be the item's own level). `gain_dps` is null on every tier-1 row, so nothing may
 * call a number method on it without checking.
 */
export function candidateGainLabel(candidate: GuildLootCandidate): string {
  if (candidate.tier === 0) {
    if (candidate.gain_dps === null || candidate.gain_unit === null) return ON_BIS_LIST;
    return `+${candidate.gain_dps.toFixed(1)}${NBSP}${candidate.gain_unit.toUpperCase()}`;
  }
  if (candidate.fills_empty_slot === true) return FILLS_EMPTY_SLOT;
  const delta = candidate.ilvl_delta ?? 0;
  return `+${delta}${NBSP}item${NBSP}level${delta === 1 ? '' : 's'}`;
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

/** "1 item won this tier", "3 items won this tier"; empty for none. */
export function itemsWonLabel(won: number | undefined): string {
  if (won === undefined || won <= 0) return '';
  return `${won} item${won === 1 ? '' : 's'} won this tier`;
}

/** Under the name: the spec with its class, then the tier count when the guild has awarded them something. */
export function candidateSubline(candidate: GuildLootCandidate): string {
  return [specClassName(candidate.spec, candidate.class), itemsWonLabel(candidate.items_won)]
    .filter((part) => part !== '')
    .join(' · ');
}

/** One row of the overview's Loot card: the item and who it is for. */
export interface LootPick {
  item: GuildLootItem;
  candidate: GuildLootCandidate;
}

function outranks(a: GuildLootCandidate, b: GuildLootCandidate): boolean {
  if (a.tier !== b.tier) return a.tier < b.tier;
  if (a.tier === 0) return (a.gain_dps ?? -Infinity) > (b.gain_dps ?? -Infinity);
  return (a.ilvl_delta ?? 0) > (b.ilvl_delta ?? 0);
}

/** The best open pick on the page: the highest tier-0 gain across every slotted drop, else the
 *  highest tier-1 row. An awarded drop is not open, and the skinned scale has no candidates. */
export function topPick(items: readonly GuildLootItem[]): LootPick | null {
  let best: LootPick | null = null;
  for (const item of rankedDrops(items)) {
    if (item.awarded_to !== null) continue;
    const candidate = item.candidates.find((row) => !row.already_equivalent);
    if (candidate === undefined) continue;
    if (best === null || outranks(candidate, best.candidate)) best = { item, candidate };
  }
  return best;
}

/** The overview's Loot line, as words around an item link and a name. */
export type OverviewLootLine =
  { kind: 'text'; text: string } | { kind: 'pick'; lead: string; pick: LootPick };

export function overviewLootLine(page: GuildLootPage): OverviewLootLine {
  const boss = page.encounters.find((encounter) => encounter.encounter_id === page.selected);
  const killed = boss?.killed === true;
  const pick = topPick(page.items);
  if (!killed && page.at_level_cap === 0) {
    return { kind: 'text', text: "Onyxia's drops need level 60. Nobody here is 60 yet." };
  }
  if (pick === null) return { kind: 'text', text: bossPickerLabel(boss) };
  return {
    kind: 'pick',
    lead: killed ? 'Next open drop: ' : 'Onyxia not killed yet. Top pick today: ',
    pick,
  };
}
