// web/src/lib/sim/published-weights.ts
// /sim/weights before a run: the weights the nightly ranker published for the character's own
// band (data/builds/<build>/bis/<spec>.json), at the character's level band, faction and preset,
// read as the run's table will be (lib/bis/band-scale-view.ts), so the reader has the published
// figures beside the run they are about to make. Null when the file has no such band.
import { bandScaleView } from '../bis/band-scale-view';
import { bandForLevel, bisPageHref } from '../bis/hover';
import { presetOf, selectBand } from '../bis/presets';
import type { ScaleView } from '../bis/scale-view';
import type { BisFile, BisPresetId } from '../bis/types';
import { factionOfRace } from '../race-faction';
import { bulkCopy } from './copy';

export interface PublishedWeights {
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

export function publishedWeights(file: BisFile, who: PublishedWeightsFor): PublishedWeights | null {
  const faction = factionOfRace(who.race);
  if (faction === null) return null;
  const level = bandForLevel(who.level);
  const band = selectBand(file, level, faction, who.preset);
  if (band === undefined || band.weights.length === 0) return null;
  const view = bandScaleView(band, who.spec, 'dps');
  if (view.rows.length === 0) return null;
  return {
    view,
    label: bulkCopy.weightsPublishedBand(level, presetOf(band)),
    href: bisPageHref(who.spec, faction, level),
  };
}
