// web/src/lib/sim/result-sections.ts
// The saved run's sections in page order (spec 2026-11-04-sim-results 3.1, 3.2, 3.4, 4.3, 4.5):
// each card's anchor and its nav label. A section with nothing to draw drops its link.
import { navCopy } from './result-copy';
import type { SimResult } from './types';

export type ResultSectionId =
  'damage' | 'timeline' | 'gear' | 'talents' | 'resource' | 'buffs' | 'rotation' | 'settings';

export interface ResultSection {
  id: ResultSectionId;
  label: string;
}

/** The nav word for the player's power: 1 rage, 3 energy, everything else mana. */
export function resourceLabel(powerType: number | undefined): string {
  if (powerType === 1) return navCopy.resource.rage;
  if (powerType === 3) return navCopy.resource.energy;
  return navCopy.resource.mana;
}

export function resultSections(result: SimResult): ResultSection[] {
  const hasSample = (result.sample?.length ?? 0) > 0;
  const hasGear = result.request.character.gear.length > 0;
  const resource = result.summary.resources[0];
  return [
    { id: 'damage', label: navCopy.damage },
    ...(hasSample ? [{ id: 'timeline' as const, label: navCopy.timeline }] : []),
    ...(hasGear ? [{ id: 'gear' as const, label: navCopy.gear }] : []),
    { id: 'talents', label: navCopy.talents },
    ...(resource === undefined
      ? []
      : [{ id: 'resource' as const, label: resourceLabel(resource.power_type) }]),
    { id: 'buffs', label: navCopy.buffs },
    { id: 'rotation', label: navCopy.rotation },
    { id: 'settings', label: navCopy.settings },
  ];
}
