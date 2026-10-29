// web/src/lib/guild/signup-strip.ts
// The sign-up strip's pure rules (spec 3.3 and 4 "Races"): row order, the prefilled role, the
// note counter and the line a failed save shows.
import type { GuildRaidRole, GuildRaidViewerCharacter, GuildSignupState } from './api';
import { raidCopy } from './raid-copy';
import { roleForSpec } from './raid-roles';

/** The note's `maxlength`. */
export const NOTE_MAX = 140;
/** The counter shows from this many characters. */
export const NOTE_COUNTER_FROM = 120;
/** The role select's visually hidden label and accessible name (spec 3.3). */
export const ROLE_FIELD_LABEL = 'Role';
/** A pressed In, Late or Bench that waits for the spec pick is dropped after this long out of focus. */
export const SPEC_ASK_IDLE_MS = 60_000;
/** The spec block's height on an unsigned row, so pressing a state shifts nothing below it: the
 *  ask measured at 360 is 116 px (label, the 44 px select, a two-line lead) and 100 px from md (one-line lead), pinned by the e2e 'shifts nothing' test; the
 *  saved control is the 44 px button + 8 top. */
export const SPEC_RESERVE_ASK = 'min-h-[116px] md:min-h-[100px]';
export const SPEC_RESERVE_SAVED = 'min-h-[52px]';
/** How long `Saved.` stays. */
export const SAVED_LINE_MS = 4000;

const RACE_LINES: Readonly<Record<string, string>> = {
  raid_locked: raidCopy.races.locked,
  raid_cancelled: raidCopy.races.cancelled,
};

/** What one sign-up save sends. */
export interface SignupChoice {
  state: GuildSignupState;
  role: GuildRaidRole;
  note: string;
  /** A spec slug picked on this save; left out when the character's saved spec stands. */
  spec?: string;
}

/** The page's answer to one save; `moved` names the character the API moved to Bench (R3). */
export type SignupOutcome = { ok: true; moved: string | null } | { ok: false; code: string; message: string };

export type SignupHandler = (
  characterKey: string,
  state: GuildSignupState,
  role: GuildRaidRole,
  note: string,
  spec?: string,
) => Promise<SignupOutcome>;

/** The header selector's current character first, the rest in the API's order. */
export function orderedCharacters(
  characters: readonly GuildRaidViewerCharacter[],
  currentKey: string | null,
): GuildRaidViewerCharacter[] {
  const current = characters.filter((c) => c.character_key === currentKey);
  const rest = characters.filter((c) => c.character_key !== currentKey);
  return [...current, ...rest];
}

/** A signed-up character keeps its role; an unsigned one takes its spec's, else the API's. */
export function initialRole(character: GuildRaidViewerCharacter): GuildRaidRole {
  if (character.state !== null) return character.role;
  return roleForSpec(character.spec) ?? character.role;
}

/** `20 left` from 120 characters, else `null`. */
export function noteCounter(note: string): string | null {
  return note.length >= NOTE_COUNTER_FROM
    ? raidCopy.stripStatus.left(Math.max(0, NOTE_MAX - note.length))
    : null;
}

/** The status line of a refused save: a race names what changed and takes focus. */
export function failureLine(code: string): { message: string; race: boolean } {
  const race = RACE_LINES[code];
  return race === undefined
    ? { message: raidCopy.stripStatus.failed, race: false }
    : { message: race, race: true };
}
