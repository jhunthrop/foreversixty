// web/src/lib/guild/assignments-view.ts
// The raid planner's Assignments (web layout spec 2026-11-05 §7.2), as the rail reads them: each
// slot's names as Tags (the viewer's own in gold, an off-role pick flagged), the officer's picker
// with the raiders who fit the slot first, the slot's next names after an add or a remove, and
// the raid warning text a leader pastes into /rw. Pure: the component and the actions read it.
import type { SelectOption } from '../ui/listbox-model';
import type {
  AssignedName,
  AssignmentRaider,
  AssignmentSlot,
  EncounterAssignments,
  GuildRaidAssignments,
} from './assignments-api';
import { uiCopy, withRetry } from '../ui/copy';

/** A picked name's Tag tone (a subset of ui/Tag's). */
export type NameTone = 'gold' | 'warn' | 'muted';

export const assignmentsCopy = {
  title: 'Assignments',
  encounterTitle: (boss: string) => `${boss} assignments`,
  notInClient: 'Not in the game client yet',
  notInClientLine: (raid: string) =>
    `${raid} has no bosses in the game client yet. Its jobs show here once the client carries them.`,
  open: 'Open',
  edit: 'Edit',
  done: 'Done',
  add: 'Add a raider',
  addLabel: (slot: string) => `Add a raider to ${slot}`,
  full: 'Slot full',
  noOneSigned: 'No one is signed up yet',
  remove: (name: string) => `Remove ${name}`,
  clear: 'Clear',
  copyWarning: 'Copy as raid warning',
  copied: uiCopy.copied,
  copyFailed: 'Your browser did not let us copy. Select the lines below and copy them.',
  post: 'Post to Discord',
  posting: 'Posting',
  posted: 'Posted under the raid post on Discord.',
  postFailed: withRetry('The post did not go out.'),
  saveFailed: withRetry('That slot did not save.'),
  nothingYet: 'No one is picked yet.',
  offRole: (name: string) => `${name} is not signed up for this role tonight.`,
  fits: 'Fits the slot',
  others: 'Others',
  yours: 'Your job',
  late: 'Late',
  bench: 'Bench',
} as const;

/** The block's encounters, empty when the raid has none or is not a raid. */
export function encountersOf(block: GuildRaidAssignments | null | undefined): EncounterAssignments[] {
  return block?.encounters ?? [];
}

/** The raids whose bosses the client does not carry yet. */
export function notInClientOf(block: GuildRaidAssignments | null | undefined): string[] {
  return block?.not_in_client ?? [];
}

/** True when the viewer holds the slot with one of their characters. */
export function holdsSlot(slot: AssignmentSlot): boolean {
  return slot.names.some((n) => n.own);
}

/** A picked name's Tag: the viewer's own in gold, an off-role pick to look at, the rest muted. */
export function nameTone(name: AssignedName): NameTone {
  if (name.own) return 'gold';
  return name.off_role ? 'warn' : 'muted';
}

/** "2 of 4": how full the slot is. */
export function filledLine(slot: AssignmentSlot): string {
  return `${slot.names.length} of ${slot.count}`;
}

/** The slot's sub line: how full it is, then what it does. */
export function slotSub(slot: AssignmentSlot): string {
  return `${filledLine(slot)} · ${slot.purpose}`;
}

export function isFull(slot: AssignmentSlot): boolean {
  return slot.names.length >= slot.count;
}

function fitsSlot(slot: AssignmentSlot, raider: AssignmentRaider): boolean {
  return slot.role === 'any' || slot.role === raider.group;
}

function stateMeta(raider: AssignmentRaider): string | undefined {
  if (raider.state === 'late') return assignmentsCopy.late;
  return raider.state === 'bench' ? assignmentsCopy.bench : undefined;
}

/** The picker's options: every signed raider not in the slot, the ones who fit it first. */
export function pickerOptions(slot: AssignmentSlot, candidates: readonly AssignmentRaider[]): SelectOption[] {
  const picked = new Set(slot.names.map((n) => n.character_key));
  const open = candidates.filter((c) => !picked.has(c.character_key));
  const option = (c: AssignmentRaider, group?: string): SelectOption => ({
    value: c.character_key,
    label: c.name,
    colour: `--color-class-${c.class.toLowerCase()}`,
    you: c.own,
    ...(group === undefined ? {} : { group }),
    ...(stateMeta(c) === undefined ? {} : { meta: stateMeta(c) }),
  });
  if (slot.role === 'any') return open.map((c) => option(c));
  return [
    ...open.filter((c) => fitsSlot(slot, c)).map((c) => option(c, assignmentsCopy.fits)),
    ...open.filter((c) => !fitsSlot(slot, c)).map((c) => option(c, assignmentsCopy.others)),
  ];
}

/** The slot's keys with `key` added at the end (unchanged when already there). */
export function withAdded(slot: AssignmentSlot, key: string): string[] {
  const keys = slot.names.map((n) => n.character_key);
  return keys.includes(key) ? keys : [...keys, key];
}

/** The slot's keys without `key`. */
export function withRemoved(slot: AssignmentSlot, key: string): string[] {
  return slot.names.map((n) => n.character_key).filter((k) => k !== key);
}

/** The /rw lines, one per filled slot, ready to paste. */
export function raidWarningText(encounter: EncounterAssignments): string {
  return encounter.raid_warning.join('\n');
}

/** True when any slot of the encounter holds a name (Post and Copy have something to say). */
export function anyPicked(encounter: EncounterAssignments): boolean {
  return encounter.slots.some((s) => s.names.length > 0);
}

/** The slot the viewer holds, by its label, for the member's "Your job" line. */
export function ownJobs(encounter: EncounterAssignments): string[] {
  return encounter.slots.filter(holdsSlot).map((s) => s.label);
}
