// web/src/lib/bis/band-view.ts
// One band's computed view for the /bis page: everything a band panel body renders.
import type { BandInfo } from './panel-view';
import type { newAtBandLinesFor, sourceGroupsFor } from './footer-view';
import type { rotationLinesFor } from './rotation-view';

export interface BandView {
  band: number;
  bandIndex: number;
  bandLabel: string;
  info: BandInfo;
  simHref: string;
  plannerHref: string;
  generatedAtDate: string;
  rotationLines: ReturnType<typeof rotationLinesFor>;
  sourceGroups: ReturnType<typeof sourceGroupsFor>;
  newAtBand: ReturnType<typeof newAtBandLinesFor>;
  nextBandLabel: string | undefined;
  nextBand: number | undefined;
}
