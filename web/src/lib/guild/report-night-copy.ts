// web/src/lib/guild/report-night-copy.ts
// Every string of the done raid's report page (web layout spec 2026-11-05 §7.2 "Raid report"),
// in a raider's voice: constants where the words are fixed, functions where the data fills them.
// Output figures follow output-figure.ts: a raw DPS or HPS is never called a parse.
import type { AttendanceOverrideStatus, PullConsumablesReason, PullResult } from './attendance-contract';
import { SEP } from './raid-copy';

export const reportNightCopy = {
  figures: { kills: 'Kills', wipes: 'Wipes', attempts: 'Attempts', deaths: 'Deaths' },
  pulls: {
    head: 'Pulls',
    caption: 'Every pull of the night',
    trash: 'Trash',
    showDeaths: 'Show the deaths',
    notRead: 'These deaths are not read from the log yet.',
    noHit: 'The log shows no hit before this death.',
    killingBlow: 'Killing blow',
    lastHits: 'Last hits',
    none: 'No pulls in this log.',
    melee: 'Melee',
  },
  result: { kill: 'Kill', wipe: 'Wipe', clear: 'Clear' } satisfies Record<PullResult, string>,
  output: {
    head: 'DPS and HPS',
    meta: 'Boss kills',
    officerEmpty: 'No boss kill to read DPS and HPS from.',
    memberEmpty: 'None of your characters was in a boss kill.',
  },
  consumables: {
    head: 'Consumables by pull',
    caption: 'Flask, food and weapon enchant by boss pull',
    raider: 'Raider',
    legend: 'Each cell reads flask, food, weapon enchant: ✓ up at the pull, × missing, · the log cannot say.',
    notInPull: 'Not in this pull',
    none: 'No boss pulls to read.',
    reasons: {
      no_character_sheet: 'the log carried no character sheet, so nothing can be said',
      not_read: 'not read from the log yet',
    } satisfies Record<PullConsumablesReason, string>,
  },
  attendance: {
    head: 'Attendance',
    fixLabel: (name: string): string => `Fix ${name}'s attendance`,
    fixPlaceholder: 'Fix',
    backToLog: 'Back to the log',
    status: {
      present: 'Present',
      absent: 'Absent',
      bench: 'Bench',
      excused: 'Excused',
    } satisfies Record<AttendanceOverrideStatus, string>,
    out: 'Out',
  },
  loot: {
    head: 'Loot awarded',
    to: (boss: string, name: string): string => `${boss}${SEP}to ${name}`,
  },
  discord: {
    head: 'Discord summary',
    post: 'Post to Discord',
    postAgain: 'Post again',
    posting: 'Posting…',
    title: (title: string): string => `Report: ${title}`,
    noKills: 'No kills',
    line: (parts: readonly string[]): string => parts.join(SEP),
    top: (name: string, figure: string, unit: string, boss: string): string =>
      `${name} ${figure} ${unit} on ${boss}`,
    lootLine: (item: string, name: string): string => `${item} to ${name}`,
    more: (n: number): string => `+${n} more`,
  },
} as const;

/** `Onyxia, pull 2`; a trash pull keeps its own name. */
export function pullTitle(name: string, attempt: number | null): string {
  return attempt === null ? name : `${name}, pull ${attempt}`;
}

/** `1:42 in`. */
export const intoPull = (clock: string): string => `${clock} in`;

/** `4,210 (1,050 overkill)`; no overkill reads the amount alone. */
export function hitAmount(amount: string, overkill: string | null): string {
  return overkill === null ? amount : `${amount} (${overkill} overkill)`;
}

/** `Deep Breath ×3`; one reads the name alone. */
export const causeCount = (ability: string, n: number): string => (n === 1 ? ability : `${ability} ×${n}`);

/** `died to Deep Breath ×3, Tail Sweep`. */
export const diedTo = (causes: readonly string[]): string => `died to ${causes.join(', ')}`;

/** `rated 88 of 100`. */
export const ratedOf = (rating: number): string => `rated ${Math.round(rating)} of 100`;

/** `Onyxia, pull 1: not read from the log yet.` */
export const unknownPullLine = (title: string, reason: string): string => `${title}: ${reason}.`;

/** `Bench, set by Kraggor`. */
export const fixedBy = (status: string, by: string): string => `${status}, set by ${by}`;

/** `The log said Not in the log`. */
export const logSaid = (auto: string): string => `The log said ${auto}`;
