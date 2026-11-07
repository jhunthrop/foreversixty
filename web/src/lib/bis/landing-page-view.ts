// web/src/lib/bis/landing-page-view.ts
// What the /bis landing shows outside the head's answer, as one pure decision: which level range
// the ControlBar has selected, the character's own range (the summary line's "Your range" link
// while another is selected), which faction the spec links carry and which spec row is "Yours".
// Re-derived whenever the viewer picks a range or the header selector changes character.
import { characterBand, selectedLandingBand } from './landing-bands';
import type { Faction } from './types';

/** The part of the selected character the page below the head needs. */
export interface LandingCharacter {
  level: number | undefined;
  faction: Faction | undefined;
  /** Undefined while the character has no spec the site knows. */
  specKey: string | undefined;
}

/** Where the viewer's own range choice came from. */
export interface BandChoice {
  /** `?band=` in the address bar (the ControlBar writes it). */
  url: number | null;
  /** The last pick kept in storage. */
  stored: number | null;
}

export interface LandingPageView {
  selectedBand: number;
  /** The character's own band; undefined without a character or under level 20. */
  youBand: number | undefined;
  /** Undefined signed out: the links then carry none and the spec page's default stands. */
  faction: Faction | undefined;
  yoursSpecKey: string | undefined;
}

export function landingPageView(choice: BandChoice, character: LandingCharacter | null): LandingPageView {
  const youBand = character === null ? undefined : characterBand(character.level);
  return {
    selectedBand: selectedLandingBand(choice.url, youBand, choice.stored),
    youBand,
    faction: character?.faction,
    yoursSpecKey: character?.specKey,
  };
}
