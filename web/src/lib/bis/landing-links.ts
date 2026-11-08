// web/src/lib/bis/landing-links.ts
// Where a spec link on the /bis landing (and a guide's BiS lane) opens: the spec's own page at a
// band, as that band's own URL (lib/bis/band-path.ts), so a crawler following it lands on the
// band's list. Signed out the link carries no faction (the spec page's default stands); with a
// character it carries theirs.
import { bisBandPath } from './band-path';
import { MIN_BAND } from './band-levels';
import type { Faction } from './types';

const DEFAULT_LINK_FACTION: Faction = 'alliance';

export interface SpecLinkTarget {
  classSlug: string;
  specSlug: string;
  band: number;
  faction?: Faction;
}

/** `/bis/warrior/fury` (the first band), `/bis/warrior/fury/40-49`, or `…?faction=horde` with Horde. */
export function specLinkHref({ classSlug, specSlug, band, faction }: SpecLinkTarget): string {
  const query = faction === undefined || faction === DEFAULT_LINK_FACTION ? '' : `?faction=${faction}`;
  return `${bisBandPath(classSlug, specSlug, band, MIN_BAND)}${query}`;
}
