// web/src/lib/guild/applicant-figures.ts
// The figures an open application carries on the officers' list (guild dense spec, h-guild-tabs):
// how long it has waited, the applicant's best boss kill as a percentile and a figure, and the open
// seat their spec fills for the next raid (the raid view's comp gaps against the applicant's group,
// which the API derives from their spec). Pure.
import type { GuildRaidGroup } from './api';
import type { CompGapRow } from './dense/week-board';
import { raidCopy } from './raid-copy';
import type { ApplicantCharacter } from './recruit-api';

const DAY_MS = 86_400_000;

export const GROUP_WORD: Readonly<Record<GuildRaidGroup, string>> = {
  tank: raidCopy.subheads.tanks,
  healer: raidCopy.subheads.healers,
  melee: raidCopy.subheads.melee,
  ranged: raidCopy.subheads.ranged,
};

/** `today`, `1 d`, `6 d`. */
export function waitedWords(createdAt: string, now: Date): string {
  const days = Math.floor((now.getTime() - new Date(createdAt).getTime()) / DAY_MS);
  return days <= 0 ? 'today' : `${days} d`;
}

export interface SeatFill {
  /** `Healers, 2 short` or `Healers, full`. */
  text: string;
  short: boolean;
}

/** The seat the applicant's group fills against the next raid's comp; null without either. */
export function seatFill(character: ApplicantCharacter, gaps: readonly CompGapRow[] | null): SeatFill | null {
  const group = character.group ?? null;
  if (group === null || gaps === null) return null;
  const row = gaps.find((gap) => gap.group === group);
  if (row === undefined) return null;
  return row.gap > 0
    ? { text: `${GROUP_WORD[group]}, ${row.gap} short`, short: true }
    : { text: `${GROUP_WORD[group]}, full`, short: false };
}
