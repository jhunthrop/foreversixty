// web/src/lib/guides/race-pills.ts
// A spec guide's Races section renders a row of name pills (no race portraits exist --
// design/DESIGN-SYSTEM.md and the current-character spec both say so) rather than a second
// prose enumeration of the class's legal races; this is the one place that list is derived,
// from the build's own class, race and combo tables (build-races.ts), the same under either
// FOREVER_DATA.
import { buildRacesForClass } from './build-races';

export interface GuideRacePill {
  slug: string;
  name: string;
  faction: string;
  recommended: boolean;
}

export function racePillsFor(
  build: string,
  classSlug: string,
  recommendedRaces: readonly string[],
): GuideRacePill[] {
  return buildRacesForClass(build, classSlug).map((race) => ({
    slug: race.slug,
    name: race.name,
    faction: race.faction,
    recommended: recommendedRaces.includes(race.slug),
  }));
}

const FACTION_COLORS: Record<string, string> = {
  alliance: 'var(--color-alliance)',
  horde: 'var(--color-horde)',
};

export function factionColorVar(faction: string): string {
  return FACTION_COLORS[faction] ?? 'var(--color-text)';
}
