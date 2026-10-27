// web/src/lib/guild/night-copy.ts
// The sentence templates of night.ts: the `Your night` outcome sentences (spec 6.1 "Your night
// sentences") and the copied Discord text's labels (spec 4 "Copy for Discord"). Kept beside
// attendance-copy.ts, which is at the file-size cap; night-copy.test.ts reads the spec for both.

/** `11 of 11 pulls`. */
const pulls = (n: number, total: number): string => `${n} of ${total} pulls`;

/** ` as DPS`, or nothing when the row carries no role. */
const asRole = (role: string | null): string => (role === null ? '' : ` as ${role}`);

export const yourNightText = {
  present: (state: string, role: string | null, n: number, total: number): string =>
    `Signed ${state}${asRole(role)}. Present from the first pull, ${pulls(n, total)}.`,
  late: (state: string, role: string | null, minutes: number, n: number, total: number): string =>
    `Signed ${state}${asRole(role)}. First seen ${minutes} min after the first pull, ${pulls(n, total)}.`,
  cameAs: (state: string, name: string, alt: string, n: number, total: number): string =>
    `Signed ${state} on ${name}; came as ${alt}, ${pulls(n, total)}.`,
  notInLog: (state: string, role: string | null): string =>
    `Signed ${state}${asRole(role)}. You are not in this log. If you were there under another name, tell an officer.`,
  unsigned: (n: number, total: number): string => `Not signed up. In the log for ${pulls(n, total)}.`,
  joined: (state: string, n: number, total: number): string =>
    `Signed ${state}. Joined for ${pulls(n, total)}.`,
  /** A bench or out character who did not come: the spec writes no sentence for it. */
  stayedOut: (state: string): string => `Signed ${state}.`,
} as const;

export const discordText = {
  kills: 'Kills:',
  mostDeaths: 'Most deaths:',
  loot: 'Loot:',
  changes: 'Biggest changes:',
  report: 'Report:',
  /** The line that stands for loot lines the 12-line cap leaves out. */
  more: (n: number): string => `+${n} more`,
} as const;

/** The Excused checkbox's accessible name: the visible word, then whose row it is. */
export const excusedLabel = (name: string): string => `Excused, ${name}`;
