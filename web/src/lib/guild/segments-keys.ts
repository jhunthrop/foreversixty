// web/src/lib/guild/segments-keys.ts
// The sign-up segments' states and their roving tab stop (spec 3.3 and 4); the keyboard model
// itself is the one segmented control's (lib/ui/segmented.ts).
import type { GuildSignupState } from './api';

export const SIGNUP_STATES = ['in', 'late', 'bench', 'out'] as const satisfies readonly GuildSignupState[];

/** The roving tab stop: the checked segment, else the first. */
export function tabbableSegment(value: GuildSignupState | null): number {
  return value === null ? 0 : SIGNUP_STATES.indexOf(value);
}
