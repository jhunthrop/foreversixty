// web/src/lib/home/plaque.ts
// Home's raid plaque (design/specs/2026-11-09-directions.md, Home): the guild's next raid with its
// clock, sign-ups, role counts, the three sign-up buttons and Before the pull; with no guild raid,
// the first raids opening. Pure: components/home/HomeGuildRaid.svelte and HomeFirstRaids.svelte
// read it, the unit tests pin it.
import type { GuildRaidRole, GuildSignupState, NextRaidView } from '../guild/api';
import { roleForSpec, rolesForClass } from '../guild/raid-roles';
import type { SpecWeek } from './home-answer';

const MINUTE = 60_000;
const HOUR = 60 * MINUTE;
const DAY = 24 * HOUR;

/** One big figure on the plaque: `32` over `days`. */
export interface PlaqueFigure {
  value: string;
  label: string;
}

/** The sign-up choices the plaque offers, in order: I'm in (the primary), Late, Can't come. */
export const PLAQUE_CHOICES = ['in', 'late', 'out'] as const satisfies readonly GuildSignupState[];
export type PlaqueChoice = (typeof PLAQUE_CHOICES)[number];

export const plaqueCopy = {
  choices: { in: "I'm in", late: 'Late', out: "Can't come" } satisfies Record<PlaqueChoice, string>,
  beforePull: 'Before the pull',
  signedUp: 'signed up',
  late: 'late',
  days: 'days',
  hours: 'hours',
  min: 'min',
  firstRaids: 'First raids open',
  raidsOpen: 'Raids open now',
  signUpFailed: 'Could not save your sign-up. Try again.',
  roles: { tanks: 'tanks', healers: 'healers', melee: 'melee', ranged: 'ranged' },
  tierLine: (grade: string, place: number, of: number) => `Tier ${grade} · #${place} of ${of}`,
} as const;

/** The clock as the plaque reads it: days and hours, minutes in the final hour; empty once passed. */
export function plaqueClock(target: Date, now: Date): PlaqueFigure[] {
  const left = target.getTime() - now.getTime();
  if (left <= 0) return [];
  if (left < HOUR) return [{ value: String(Math.ceil(left / MINUTE)), label: plaqueCopy.min }];
  return [
    { value: String(Math.floor(left / DAY)), label: plaqueCopy.days },
    { value: String(Math.floor((left % DAY) / HOUR)), label: plaqueCopy.hours },
  ];
}

/** The clock read aloud: `32 days, 5 hours`. */
export function plaqueClockLabel(parts: readonly PlaqueFigure[]): string {
  return parts.map((part) => `${part.value} ${part.label}`).join(', ');
}

/** Signed up and late. */
export function signupFigures(counts: NextRaidView['counts']): PlaqueFigure[] {
  return [
    { value: String(counts.in), label: plaqueCopy.signedUp },
    { value: String(counts.late), label: plaqueCopy.late },
  ];
}

/** The role line: tanks, healers, melee and ranged (In only). An API before melee and ranged
 *  sends DPS alone, which the line then leaves out rather than guess the split. */
export function roleCounts(counts: NextRaidView['counts']): { n: number; word: string }[] {
  const rows = [
    { n: counts.tanks, word: plaqueCopy.roles.tanks },
    { n: counts.healers, word: plaqueCopy.roles.healers },
  ];
  if (counts.melee === undefined || counts.ranged === undefined) return rows;
  return [
    ...rows,
    { n: counts.melee, word: plaqueCopy.roles.melee },
    { n: counts.ranged, word: plaqueCopy.roles.ranged },
  ];
}

/** The role a plaque sign-up carries: the one already saved, else the spec's, else the class's
 *  only role; null when it cannot be told (the plaque then links to the raid page instead). */
export function signupRole(
  own: NextRaidView['own'],
  character: { class?: string; spec?: string },
): GuildRaidRole | null {
  if (own !== null) return own.role;
  const spec = character.spec === undefined ? null : roleForSpec(character.spec);
  if (spec !== null) return spec;
  const roles = character.class === undefined ? [] : rolesForClass(character.class);
  return roles.length === 1 ? roles[0] : null;
}

/** The plaque's choice for a saved state: Bench reads as none of the three. */
export function choiceOf(state: GuildSignupState | null): PlaqueChoice | null {
  return state !== null && (PLAQUE_CHOICES as readonly string[]).includes(state)
    ? (state as PlaqueChoice)
    : null;
}

/** `Wednesday 9 December`, a date-only instant read as that calendar day. */
export function longDay(iso: string, locale = 'en-GB'): string {
  return new Date(iso).toLocaleDateString(locale, {
    weekday: 'long',
    day: 'numeric',
    month: 'long',
    timeZone: 'UTC',
  });
}

/** The tier line under the guild sims: `Tier A · #2 of 20`; null with no grade. */
export function tierLine(week: SpecWeek | null): string | null {
  if (week === null || week.grade === null) return null;
  return plaqueCopy.tierLine(week.grade, week.place, week.of);
}
