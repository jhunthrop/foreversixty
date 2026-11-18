// web/src/lib/bis/leveling-fold.ts
// The spec page's Leveling section (2026-11-15, the SEO report of 2026-11-10): the leveling
// bands (lib/bis/band-path.ts isFoldedBand: 20-29 to 50-59) server-rendered on
// /bis/<class>/<spec>, both factions, from the same ranked file and the same view builder the
// guide's lane and the doll read, so the three can never name different items.
import { raceNameFor } from '../guides/race-table';
import { guideBisDeps, guideBisFrom, type GuideBisPanel } from '../guides/guide-bis';
import { isFoldedBand, specPageBand } from './band-path';
import { loadBisFile } from './load';
import type { BisFile, Faction } from './types';

/** The faction the section opens on: the doll's own default (page-state.ts). */
const DEFAULT_FACTION: Faction = 'alliance';

/** The folded bands' panels from a ranked file, each band's Alliance then Horde, ascending. */
export function foldedPanelsFrom(
  file: BisFile,
  deps: Parameters<typeof guideBisFrom>[2],
  raceName: (slug: string) => string,
): GuideBisPanel[] {
  const bis = guideBisFrom(file, DEFAULT_FACTION, deps, raceName);
  if (bis === undefined) return [];
  // The band the doll opens on is never repeated below it (a spec ranked short of 60 opens on its first).
  const home = specPageBand(bis.bands.map((entry) => entry.band));
  return bis.panels
    .filter((panel) => isFoldedBand(panel.band) && panel.band !== home)
    .sort((a, b) => a.band - b.band || bis.factions.indexOf(a.faction) - bis.factions.indexOf(b.faction));
}

/** The spec's folded panels on the build; empty when the spec has no ranked file. */
export function loadFoldedPanels(build: string, classSlug: string, spec: string): GuideBisPanel[] {
  const file = loadBisFile(spec, build);
  if (file === null) return [];
  return foldedPanelsFrom(file, guideBisDeps(build, classSlug, spec), (slug) => raceNameFor(build, slug));
}
