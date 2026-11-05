// web/src/lib/sim/result-sections.ts
// The saved run's sections in page order (spec 2026-11-04-sim-results 3.1, 3.2, 3.4, 4.3, 4.5):
// each card's anchor and its nav label. A section with nothing to draw drops its link.
import { primaryPowerTrack } from './primary-power';
import { navCopy } from './result-copy';
import type { SimResult } from './types';

export type ResultSectionId =
  'damage' | 'timeline' | 'gear' | 'talents' | 'resource' | 'buffs' | 'rotation' | 'settings';

export interface ResultSection {
  id: ResultSectionId;
  label: string;
}

/** The word for the player's power, from the class (primary-power.ts); null when the run has no track for it. */
export function resourceLabel(result: SimResult): string | null {
  const found = primaryPowerTrack(result);
  return found === null ? null : navCopy.resource[found.power];
}

export function resultSections(result: SimResult): ResultSection[] {
  const hasSample = (result.sample?.length ?? 0) > 0;
  const hasGear = result.request.character.gear.length > 0;
  const resource = resourceLabel(result);
  return [
    { id: 'damage', label: navCopy.damage },
    ...(hasSample ? [{ id: 'timeline' as const, label: navCopy.timeline }] : []),
    ...(hasGear ? [{ id: 'gear' as const, label: navCopy.gear }] : []),
    { id: 'talents', label: navCopy.talents },
    ...(resource === null ? [] : [{ id: 'resource' as const, label: resource }]),
    { id: 'buffs', label: navCopy.buffs },
    { id: 'rotation', label: navCopy.rotation },
    { id: 'settings', label: navCopy.settings },
  ];
}
