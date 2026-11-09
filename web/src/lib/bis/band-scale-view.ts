// web/src/lib/bis/band-scale-view.ts
// One published band's weights as the SimulationCraft-form table (scale-view.ts): the Best in slot
// page's weights panel (panel-view.ts) and /sim/weights' published table read a band the same way.
import { weightsForGainedStats } from './expertise-cap';
import { scaleViewFor, UNMEASURED_SCALE_VIEW, type ScaleView } from './scale-view';
import type { SlotScoreUnit } from './tank-view';
import type { BisBand } from './types';

export function bandScaleView(band: BisBand, spec: string, unit: SlotScoreUnit): ScaleView {
  if (band.weights_reason !== undefined && band.weights_reason !== null) return UNMEASURED_SCALE_VIEW;
  return scaleViewFor({
    weights: weightsForGainedStats(band.weights, band.expertise_to_cap),
    spec,
    unit,
    referenceDpsPerPoint: band.reference_dps_per_point ?? null,
    publishedAnchorStat: band.scale_reference_stat ?? null,
    publishedHasteScaleFactor: band.haste_scale_factor ?? null,
    hasteOnItems: band.haste_on_items ?? true,
  });
}
