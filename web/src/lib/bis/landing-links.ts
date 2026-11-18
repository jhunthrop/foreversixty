// web/src/lib/bis/landing-links.ts
// Where a spec link on the /bis landing (and a guide's BiS lane) opens: the page that holds the
// band (lib/bis/band-path.ts), a mid-level band as the spec page's anchor, so a crawler following
// it lands on the band's list. Signed out the link carries no faction (the spec page's default stands); with a
// character it carries theirs.
import { bisBandHref, bisSpecPath } from './band-path';
import { MAX_BAND, MIN_BAND } from './band-levels';
import type { Faction } from './types';

const DEFAULT_LINK_FACTION: Faction = 'alliance';

function factionQuery(faction: Faction | undefined): string {
  return faction === undefined || faction === DEFAULT_LINK_FACTION ? '' : `?faction=${faction}`;
}

export interface SpecLinkTarget {
  classSlug: string;
  specSlug: string;
  band: number;
  faction?: Faction;
}

/** `/bis/warrior/fury` (the first band), `/bis/warrior/fury#levels-40-49`, `/bis/warrior/fury/60`,
 *  with `?faction=horde` before any anchor for Horde. */
export function specLinkHref({ classSlug, specSlug, band, faction }: SpecLinkTarget): string {
  return bisBandHref(classSlug, specSlug, band, MIN_BAND, factionQuery(faction));
}

/** A /bis landing spec link: as specLinkHref, but level 60 opens the spec page itself, which shows
 *  level 60 first (owner 2026-11-17), so the landing keeps linking the spec hubs. */
export function landingSpecHref(target: SpecLinkTarget): string {
  if (target.band < MAX_BAND) return specLinkHref(target);
  return `${bisSpecPath(target.classSlug, target.specSlug)}${factionQuery(target.faction)}`;
}
