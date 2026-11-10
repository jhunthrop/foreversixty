// web/src/lib/sim/published-weights.ts
// /sim/weights before a run: the weights the nightly ranker published for the character's own
// band (data/builds/<build>/bis/<spec>.json), at the character's level band, faction and preset,
// read as the run's table will be (lib/bis/band-scale-view.ts), so the reader has the published
// figures beside the run they are about to make. Null when the file has no such band.
//
// With no character loaded the page shows one spec's level 60 figures (DEFAULT_WEIGHTS_SPEC),
// computed at build time by sim/weights.astro and handed to the island on its mount, so the
// server render and the island show the same table.
import { MAX_BAND } from '../bis/band-levels';
import { bandScaleView } from '../bis/band-scale-view';
import { bandForLevel, bisPageHref } from '../bis/hover';
import { presetOf, selectBand } from '../bis/presets';
import type { ScaleView } from '../bis/scale-view';
import type { BisFile, BisPresetId, Faction } from '../bis/types';
import { factionOfRace } from '../race-faction';
import { bulkCopy } from './copy';

export interface PublishedWeights {
  /** The spec the figures belong to ("warrior-fury"). */
  spec: string;
  view: ScaleView;
  /** "Level 60 raid set": the band the figures belong to, the link's label. */
  label: string;
  href: string;
}

export interface PublishedWeightsFor {
  spec: string;
  race: string;
  level: number;
  /** The run's preset; undefined (Custom) reads the band's headline preset. */
  preset: BisPresetId | undefined;
}

interface PublishedWeightsAt {
  spec: string;
  faction: Faction;
  level: number;
  preset: BisPresetId | undefined;
}

/** The spec /sim/weights shows before a character is loaded: the planner's opening class. */
export const DEFAULT_WEIGHTS_SPEC = 'warrior-fury';

/** The faction read first for the default table; the other one when the file has no such band.
 *  The order is a tie-break, not a ranking of the factions: the link carries whichever faction
 *  the figures were read from. */
const DEFAULT_FACTIONS: readonly Faction[] = ['alliance', 'horde'];

function publishedWeightsAt(file: BisFile, at: PublishedWeightsAt): PublishedWeights | null {
  const level = bandForLevel(at.level);
  const band = selectBand(file, level, at.faction, at.preset);
  if (band === undefined || band.weights.length === 0) return null;
  const view = bandScaleView(band, at.spec, 'dps');
  if (view.rows.length === 0) return null;
  return {
    spec: at.spec,
    view,
    label: bulkCopy.weightsPublishedBand(level, presetOf(band)),
    href: bisPageHref(at.spec, at.faction, level),
  };
}

export function publishedWeights(file: BisFile, who: PublishedWeightsFor): PublishedWeights | null {
  const faction = factionOfRace(who.race);
  if (faction === null) return null;
  return publishedWeightsAt(file, { spec: who.spec, faction, level: who.level, preset: who.preset });
}

/** The default spec's level 60 headline figures, or null when its file publishes none. */
export function defaultPublishedWeights(file: BisFile): PublishedWeights | null {
  for (const faction of DEFAULT_FACTIONS) {
    const found = publishedWeightsAt(file, {
      spec: DEFAULT_WEIGHTS_SPEC,
      faction,
      level: MAX_BAND,
      preset: undefined,
    });
    if (found !== null) return found;
  }
  return null;
}

function isPublishedWeights(value: unknown): value is PublishedWeights {
  if (typeof value !== 'object' || value === null) return false;
  const candidate = value as Partial<PublishedWeights>;
  return (
    typeof candidate.spec === 'string' &&
    typeof candidate.label === 'string' &&
    typeof candidate.href === 'string' &&
    typeof candidate.view === 'object' &&
    candidate.view !== null &&
    Array.isArray(candidate.view.rows) &&
    typeof candidate.view.noteLine === 'string'
  );
}

/** The mount's `data-default-weights`, read back; null when absent or not the shape written. */
export function parseDefaultWeights(raw: string | undefined): PublishedWeights | null {
  if (raw === undefined || raw === '') return null;
  try {
    const value: unknown = JSON.parse(raw);
    return isPublishedWeights(value) ? value : null;
  } catch {
    return null;
  }
}
