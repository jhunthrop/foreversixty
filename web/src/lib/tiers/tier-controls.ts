// web/src/lib/tiers/tier-controls.ts
// The tier list's two page controls as pure functions: which faction a URL asks for and where
// the arrow keys move on the role tabs. The page's inline pre-paint script and its module
// script share these names, so the URL parameter and the attribute the CSS reads are spelled
// once.
import type { Faction } from '../bis/types';

export const FACTION_PARAM = 'faction';
/** The `<html>` attribute the stylesheet reads to show one faction's lists. */
export const FACTION_ATTRIBUTE = 'data-tier-faction';
export const DEFAULT_FACTION: Faction = 'alliance';
export const TIER_FACTION_ORDER: readonly Faction[] = ['alliance', 'horde'];

/** The faction a query string asks for; anything else is the default. */
export function parseFaction(search: string): Faction {
  const asked = new URLSearchParams(search).get(FACTION_PARAM);
  return TIER_FACTION_ORDER.find((faction) => faction === asked) ?? DEFAULT_FACTION;
}

/** `path` with the faction carried; the default faction keeps the bare path. */
export function hrefWithFaction(path: string, faction: Faction): string {
  return faction === DEFAULT_FACTION ? path : `${path}?${FACTION_PARAM}=${faction}`;
}

/** The index the tab strip's focus moves to for a key, or `null` for a key it ignores. */
export function nextTabIndex(key: string, index: number, count: number): number | null {
  switch (key) {
    case 'ArrowRight':
      return (index + 1) % count;
    case 'ArrowLeft':
      return (index - 1 + count) % count;
    case 'Home':
      return 0;
    case 'End':
      return count - 1;
    default:
      return null;
  }
}
