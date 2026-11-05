// web/src/lib/guild/signup-strip.ts
// The sign-up strip's pure rules (raids spec 3.3 and 4 "Races", raid comp spec 3.2): row order,
// the pick a row starts on, which presses wait for a pick, the note counter and the line a failed
// save shows.
import type { GuildRaidRole, GuildRaidViewerCharacter, GuildSignupState, GuildZoneKind } from './api';
import { pickFor, type SpecPick } from './raid-comp';
import { endedLineOf, raidCopy } from './raid-copy';

/** The note's `maxlength`. */
export const NOTE_MAX = 140;
/** The counter shows from this many characters. */
export const NOTE_COUNTER_FROM = 120;
/** A pressed In, Late or Bench that waits for the spec pick is dropped after this long out of focus. */
export const SPEC_ASK_IDLE_MS = 60_000;
/** The sign-up states that carry a spec; Out never needs one. */
const SPEC_STATES: readonly GuildSignupState[] = ['in', 'late', 'bench'];
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

/** The pick a row starts on (raid comp spec 3.2): the saved raid spec, else the API's default
 *  spec, with the sign-up's role (a returning off-tank reads `Fury · Tank`); `null` shows the
 *  placeholder. */
export function defaultPick(character: GuildRaidViewerCharacter): SpecPick | null {
  return pickFor(character.class, character.raid_spec ?? character.default_spec ?? null, character.role);
}

/** Whether a press of this state carries the pick: In, Late and Bench do, Out never. */
export function stateTakesSpec(state: GuildSignupState): boolean {
  return SPEC_STATES.includes(state);
}

/** In, Late or Bench pressed with no spec picked waits for the pick (`Pick a spec first.`). */
export function pressWaitsForSpec(state: GuildSignupState, pick: SpecPick | null): boolean {
  return pick === null && stateTakesSpec(state);
}

/** The role a save sends: the pick's, else the character's own. */
export function roleOfPick(pick: SpecPick | null, character: GuildRaidViewerCharacter): GuildRaidRole {
  return pick?.role ?? character.role;
}

/** `20 left` from 120 characters, else `null`. */
export function noteCounter(note: string): string | null {
  return note.length >= NOTE_COUNTER_FROM
    ? raidCopy.stripStatus.left(Math.max(0, NOTE_MAX - note.length))
    : null;
}

/** The status line of a refused save: a race names what changed and takes focus. */
export function failureLine(code: string, kind: GuildZoneKind = 'raid'): { message: string; race: boolean } {
  const race = code === 'raid_ended' ? endedLineOf(kind) : RACE_LINES[code];
  return race === undefined
    ? { message: raidCopy.stripStatus.failed, race: false }
    : { message: race, race: true };
}
