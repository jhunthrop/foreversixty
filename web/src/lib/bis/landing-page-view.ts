// web/src/lib/bis/landing-page-view.ts
// What the /bis landing shows outside the lead, as one pure decision: which band pill is selected,
// which carries the "you" mark, which faction the spec links carry, which spec row is "Yours" and
// which heading the grid has. Re-derived whenever the viewer picks a pill or the header selector
// changes character.
import { characterBand, selectedLandingBand } from './landing-bands';
import { landingCopy } from './landing-copy';
import type { Faction } from './types';

/** The part of the selected character the page below the lead needs. */
export interface LandingCharacter {
  level: number | undefined;
  faction: Faction | undefined;
  /** Undefined while the character has no spec the site knows. */
  specKey: string | undefined;
}

export interface LandingPageView {
  selectedBand: number;
  /** The character's own band (the "you" mark); undefined without a character or under level 20. */
  youBand: number | undefined;
  /** Undefined signed out: the links then carry none and the spec page's default stands. */
  faction: Faction | undefined;
  yoursSpecKey: string | undefined;
  gridHeading: string;
}

export function landingPageView(
  explicitBand: number | null,
  character: LandingCharacter | null,
): LandingPageView {
  const youBand = character === null ? undefined : characterBand(character.level);
  return {
    selectedBand: selectedLandingBand(explicitBand, youBand),
    youBand,
    faction: character?.faction,
    yoursSpecKey: character?.specKey,
    gridHeading: character === null ? landingCopy.gridHeadingSignedOut : landingCopy.gridHeadingSignedIn,
  };
}
