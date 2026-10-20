// web/src/lib/planner/band-title.ts
// The name a band build is saved under. "Load the 20 to 29 build" used to leave the draft
// untitled, so the inbox entry and the share card described it by race and class. The title is
// the band's own label, with the spec when the loaded build has one: "20 to 29 · Beast Mastery".
const SEPARATOR = ' · ';

/** "20 to 29", or "20 to 29 · Beast Mastery" once the loaded build names a spec. */
export function bandBuildTitle(bandLabel: string, specName: string | null): string {
  return specName === null || specName === '' ? bandLabel : `${bandLabel}${SEPARATOR}${specName}`;
}

/**
 * The title after a band load. The draft's own title is replaced only when it is empty or is
 * still the title the previous band load gave it; a title the player typed wins.
 */
export function titleAfterBandLoad(current: string, previousBandTitle: string, bandTitle: string): string {
  return current === '' || current === previousBandTitle ? bandTitle : current;
}
