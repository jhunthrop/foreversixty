// web/src/lib/bis/landing-links.ts
// Where a spec link on the /bis landing opens: the spec's own page at a band. Signed out the link
// carries no faction (the spec page's default stands); with a character it carries theirs.
import type { Faction } from './types';

const DEFAULT_LINK_FACTION: Faction = 'alliance';

export interface SpecLinkTarget {
  classSlug: string;
  specSlug: string;
  band: number;
  faction?: Faction;
}

/** `/bis/warrior/fury#band-alliance-20`, or `…?faction=horde#band-horde-40` with a faction. */
export function specLinkHref({ classSlug, specSlug, band, faction }: SpecLinkTarget): string {
  const query = faction === undefined ? '' : `?faction=${faction}`;
  return `/bis/${classSlug}/${specSlug}${query}#band-${faction ?? DEFAULT_LINK_FACTION}-${band}`;
}
