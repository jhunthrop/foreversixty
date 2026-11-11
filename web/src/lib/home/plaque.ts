// web/src/lib/home/plaque.ts
// Home's raid plaque (design/specs/2026-11-09-directions.md, Home): the guild's next raid with its
// clock, sign-ups, role counts and Before the pull (the sign-up is the guild's SignupSegments); with no guild raid,
// the first raids opening. Pure: components/home/HomeGuildRaid.svelte and HomeFirstRaids.svelte
// read it, the unit tests pin it.
import type { GuildRaidRole, NextRaidView } from '../guild/api';
import { roleForSpec, rolesForClass } from '../guild/raid-roles';
import { clock, longDay, viewerLocale } from '../dates';
import { clockParts } from '../ui/countdown';

/** One big figure on the plaque: `32` over `days`. */
export interface PlaqueFigure {
  value: string;
  label: string;
}

export const plaqueCopy = {
  beforePull: 'Before the pull',
  in: 'in',
  late: 'late',
  firstRaids: 'First raids open',
  raidsOpen: 'Raids open now',
  roles: { tanks: 'tanks', healers: 'healers', melee: 'melee', ranged: 'ranged' },
} as const;

/** The one countdown (lib/ui/countdown.ts) as the plaque's big figures: `32` over `days`; empty
 *  once passed. Read aloud through `countdownText`. */
export function countdownFigures(target: Date, now: Date): PlaqueFigure[] {
  return clockParts(target, now).map((part) => ({ value: part.value, label: part.unit }));
}

/** Signed up and late. */
export function signupFigures(counts: NextRaidView['counts']): PlaqueFigure[] {
  return [
    { value: String(counts.in), label: plaqueCopy.in },
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

/** The raid's start on the viewer's own calendar and clock, with no zone label:
 *  `Wednesday 9 December · 8 pm` (`· 20:00` for a 24-hour locale). */
export function raidWhen(iso: string, timeZone: string, locale: string = viewerLocale()): string {
  const at = new Date(iso);
  return `${longDay(at, timeZone)} · ${clock(at, timeZone, { locale })}`;
}
