// web/src/lib/guild/loot-council-view.ts
// The loot council's words (web layout spec 2026-11-05 §7.2): "This dropped. Who gets it, and
// why?" A night's drops are what a raider's addon recorded dropping, plus what the council awarded,
// voted on or flagged; the current night also offers the whole loot table (GET .../loot) to open a
// drop on a night no addon recorded.
// This module only formats what the API sent; it never re-ranks or re-decides an award.
import { counted } from '../plural';
import { slotLine } from '../items/slot-words';
import { raidsAttended } from './attendance-copy';
import type { GuildLootItem, GuildLootRaid } from './api';
import type { SelectOption } from '../ui/listbox-model';
import type {
  AwardRecord,
  DropKind,
  LootCouncilCandidate,
  LootDecision,
  LootNight,
  LootNightDrop,
  MsOs,
} from './loot-council-api';
import { candidateGainLabel } from './loot-view';

/** A decision's reason, as the API caps it. */
export const REASON_MAX = 200;
export const NO_DROPS_OFFICER = 'Nothing opened on this night yet. Open a drop from the loot table.';
export const NO_DROPS_MEMBER = 'Nothing on this night yet.';
export const NO_DROPS_PAST = 'Nothing was recorded, awarded, voted on or flagged on this night.';
export const PICK_PROMPT = 'Pick a drop to see who wants it.';
export const NO_CANDIDATES = 'Nobody here can use it yet.';
export const MEMBER_FLAG_HELP = 'Tell the officers how you want it. Only officers see it.';
export const MEMBER_FLAG_NONE = 'None of your characters can use this drop.';
export const CLOSED_LINE = 'Awarded. Voting on this drop is closed.';

const MONTHS = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'];
const WEEKDAYS = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'];

/** "Tue 8 Dec" from the night's YYYY-MM-DD (already the guild's raid day, so read as UTC). */
export function nightDay(date: string): string {
  const d = new Date(`${date}T00:00:00Z`);
  if (Number.isNaN(d.getTime())) return date;
  return `${WEEKDAYS[d.getUTCDay()]} ${d.getUTCDate()} ${MONTHS[d.getUTCMonth()]}`;
}

/** The night chooser: one option per night, newest first. */
export function nightOptions(nights: readonly LootNight[]): { value: string; label: string }[] {
  return nights.map((night) => ({ value: night.night, label: nightDay(night.date) }));
}

/** The night to open: the current one, else the newest. */
export function defaultNight(nights: readonly LootNight[]): LootNight | null {
  return nights.find((night) => night.current) ?? nights[0] ?? null;
}

/** "3 open" for a night (batch 3: the date and how many wait for an award). */
export function nightSummary(night: LootNight): string {
  if (night.drops.length === 0) return 'No drops yet';
  return `${night.drops.filter((drop) => drop.state === 'open').length} open`;
}

/** One row of the drops Plate. */
export interface DropRow {
  itemId: number;
  name: string;
  icon: string;
  quality: number;
  kind: DropKind;
  encounterId: number;
  state: 'open' | 'awarded';
  award: LootDecision | null;
  /** The muted line: slot and type, then who can use it or who got it. */
  sub: string;
}

function kindWord(kind: DropKind): string {
  if (kind === 'quest') return 'Quest item';
  if (kind === 'bag') return 'Bag';
  return 'Skinned';
}

/** What the drop is: the game's slot words, or its kind. */
export function dropWhat(kind: DropKind, table: GuildLootItem | undefined): string {
  if (kind !== 'gear') return kindWord(kind);
  return table === undefined ? '' : slotLine(table.inventory_type, table.item_class, table.item_subclass);
}

function canUse(table: GuildLootItem | undefined): number | null {
  if (table === undefined) return null;
  return table.kind === 'gear' ? table.candidates.length : table.eligible.length;
}

type SubDrop = Pick<LootNightDrop, 'kind' | 'award'> & Partial<Pick<LootNightDrop, 'count' | 'looted_by'>>;

function dropSub(drop: SubDrop, table: GuildLootItem | undefined): string {
  const parts = [dropWhat(drop.kind, table)];
  if ((drop.count ?? 0) > 1) parts.push(`${drop.count} dropped`);
  if (drop.award !== null) parts.push(`Awarded to ${drop.award.name}`);
  else {
    const count = canUse(table);
    if (count !== null) parts.push(counted(count, 'candidate'));
  }
  // The loot window's own winner, beside the council's award: who actually took it in game.
  const looted = drop.looted_by ?? [];
  if (looted.length > 0) parts.push(`Looted by ${looted.join(', ')}`);
  return parts.filter((part) => part !== '').join(' · ');
}

const tableIndex = (table: readonly GuildLootItem[]): Map<number, GuildLootItem> =>
  new Map(table.map((item) => [item.item_id, item]));

/** The night's drops as Plate rows, in the loot table's order (the API's). */
export function dropRows(night: LootNight, table: readonly GuildLootItem[]): DropRow[] {
  const byId = tableIndex(table);
  return night.drops.map((drop) => ({
    itemId: drop.item_id,
    name: drop.name,
    icon: drop.icon,
    quality: drop.quality,
    kind: drop.kind,
    encounterId: drop.encounter_id,
    state: drop.state,
    award: drop.award,
    sub: dropSub(drop, byId.get(drop.item_id)),
  }));
}

/** The viewer's own characters who can use a drop (the loot table's candidates and eligible). */
export function ownUsers(
  item: GuildLootItem | undefined,
  ownKeys: readonly string[],
): { key: string; name: string }[] {
  if (item === undefined) return [];
  const rows = [
    ...item.candidates.map((c) => ({ key: c.character_key, name: c.name })),
    ...item.eligible.map((e) => ({ key: e.character_key, name: e.name })),
  ];
  const out = new Map<string, { key: string; name: string }>();
  for (const row of rows) if (ownKeys.includes(row.key) && !out.has(row.key)) out.set(row.key, row);
  return [...out.values()];
}

/**
 * The loot table's drops the current night can open, not already on it: an officer any drop
 * someone can take; a member only those their characters can use. The skinned scale is never a pick.
 */
export function openable(
  night: LootNight,
  table: readonly GuildLootItem[],
  ownKeys: readonly string[] | null,
): GuildLootItem[] {
  if (!night.current) return [];
  const onNight = new Set(night.drops.map((drop) => drop.item_id));
  return table.filter(
    (item) =>
      item.kind !== 'skinned' &&
      !onNight.has(item.item_id) &&
      (canUse(item) ?? 0) > 0 &&
      (ownKeys === null || ownUsers(item, ownKeys).length > 0),
  );
}

/** The picker's line for a raid the game client does not carry yet. */
export const NOT_IN_CLIENT = 'Not in the game client yet';

/**
 * The drop picker's options: a prompt, then the openable drops grouped by raid then boss in the
 * table's order ("Onyxia's Lair · Onyxia"). A raid the game client does not carry yet keeps its
 * group with one disabled "Not in the game client yet" line, never hidden.
 */
export function pickOptions(open: readonly GuildLootItem[], raids: readonly GuildLootRaid[]): SelectOption[] {
  const out: SelectOption[] = [{ value: '', label: 'Pick a drop' }];
  for (const raid of raids) {
    if (!raid.in_client) {
      out.push({ value: `raid:${raid.key}`, label: NOT_IN_CLIENT, group: raid.name, disabled: true });
      continue;
    }
    for (const encounter of raid.encounters) {
      for (const item of open.filter((row) => row.encounter_id === encounter)) {
        out.push({ value: String(item.item_id), label: item.name, group: `${raid.name} · ${item.boss}` });
      }
    }
  }
  return out;
}

/** A loot table item as a drop row, before anything happened to it on the night. */
export function tableDrop(item: GuildLootItem): DropRow {
  return {
    itemId: item.item_id,
    name: item.name,
    icon: item.icon,
    quality: item.quality,
    kind: item.kind,
    encounterId: item.encounter_id,
    state: 'open',
    award: null,
    sub: dropSub({ kind: item.kind, award: null }, item),
  };
}

/** "+14.2 DPS over Mirah's Song", "+12 item levels, empty slot", "On their BiS list". */
export function candidateGain(candidate: LootCouncilCandidate): string {
  const gain = candidateGainLabel(candidate);
  if (candidate.fills_empty_slot) return gain;
  if (candidate.replaces === '') return gain;
  return `${gain} over ${candidate.replaces}`;
}

/** The candidate's facts after the gain: attendance and loot this tier. */
export function candidateFacts(candidate: LootCouncilCandidate): string[] {
  const facts: string[] = [];
  if (candidate.attendance.nights > 0)
    facts.push(raidsAttended(candidate.attendance.present, candidate.attendance.nights));
  facts.push(`${counted(candidate.items_won, 'item')} this tier`);
  return facts;
}

/** The candidate's whole sub line. */
export function candidateSub(candidate: LootCouncilCandidate): string {
  return [candidateGain(candidate), ...candidateFacts(candidate)].join(' · ');
}

export interface MsOsTag {
  label: 'MS' | 'OS';
  declared: boolean;
  title: string;
}

const MS_OS_WORD: Readonly<Record<MsOs, string>> = { ms: 'main spec', os: 'off spec' };

/** The MS or OS Tag: what they flagged, else what their BiS list says; none for a quest drop. */
export function msOsTag(candidate: LootCouncilCandidate): MsOsTag | null {
  if (candidate.ms_os === null) return null;
  const word = MS_OS_WORD[candidate.ms_os];
  const declared = candidate.ms_os_source === 'declared';
  return {
    label: candidate.ms_os === 'ms' ? 'MS' : 'OS',
    declared,
    title: declared ? `They flagged it ${word}` : `Not flagged. ${capitalized(word)} by their BiS list`,
  };
}

function capitalized(text: string): string {
  return text.charAt(0).toUpperCase() + text.slice(1);
}

/** "Eskhandar's Collar to Hollowmend". */
export function decisionTitle(decision: LootDecision): string {
  return `${decision.item_name} to ${decision.name}`;
}

/** The record's own parts: "Off spec" when it went off spec (main spec is the rule, so it is not
 *  named), then "3 yes, 0 no of 4"; a part the award never kept is left out. */
export function recordParts(record: AwardRecord): string[] {
  const parts: string[] = [];
  if (record.ms_os === 'os') parts.push(capitalized(MS_OS_WORD[record.ms_os]));
  if (record.votes_yes !== null && record.votes_no !== null) {
    const of = record.votes_officers === null ? '' : ` of ${record.votes_officers}`;
    parts.push(`${record.votes_yes} yes, ${record.votes_no} no${of}`);
  }
  return parts;
}

/** The votes, short: "3 of 3 yes" when nobody said no, "3 yes, 1 no of 4" when someone did
 *  ('' when the award kept none). */
function votesShort(record: AwardRecord): string {
  const { votes_yes: yes, votes_no: no, votes_officers: officers } = record;
  if (yes === null || no === null) return '';
  const of = officers === null ? '' : ` of ${officers}`;
  return no === 0 ? `${yes}${of} yes` : `${yes} yes, ${no} no${of}`;
}

/** Why a drop went where it did, in one short line: "Off spec · 3 of 3 yes · Best in slot" (off
 *  spec only when it went off spec, the votes, the officer's reason); '' for an award made
 *  before the council kept records. */
export function recordLine(record: AwardRecord): string {
  const reason = record.reason?.trim() ?? '';
  return [record.ms_os === 'os' ? capitalized(MS_OS_WORD.os) : '', votesShort(record), reason]
    .filter((part) => part !== '')
    .join(' · ');
}

/** "3 yes, 0 no of 4 · by Ironbrand · Tue 8 Dec" ("Off spec · ..." when it went off spec); parts the award never kept are left out. */
export function decisionSub(decision: LootDecision, timeZone?: string): string {
  const parts = recordParts(decision);
  if (decision.awarded_by !== '') parts.push(`by ${decision.awarded_by}`);
  const at = new Date(decision.at);
  if (!Number.isNaN(at.getTime()))
    parts.push(
      at
        .toLocaleDateString('en-GB', { weekday: 'short', day: 'numeric', month: 'short', timeZone })
        .replace(',', ''),
    );
  return parts.join(' · ');
}

/** The reason as the award sends it: trimmed, null when empty. */
export function awardReason(text: string): string | null {
  const trimmed = text.trim();
  return trimmed === '' ? null : trimmed;
}

/** The reason field's error, empty when it fits. */
export function reasonError(text: string): string {
  return text.trim().length > REASON_MAX ? `Keep the reason to ${REASON_MAX} characters.` : '';
}

/** "Award Vis'kag the Bloodletter to Seraphine?" */
export function awardAsk(itemName: string, name: string): string {
  return `Award ${itemName} to ${name}?`;
}
