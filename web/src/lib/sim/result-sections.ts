// web/src/lib/sim/result-sections.ts
// The saved run's sections in page order (spec 2026-11-04-sim-results 3.1, 3.2, 3.4, 4.3, 4.5):
// each card's anchor and its nav label. A section with nothing to draw drops its link.
import { navCopy } from './result-copy';
import type { SimResult } from './types';

export type ResultSectionId =
  | 'damage'
  | 'timeline'
  | 'gear'
  | 'talents'
  | 'resource'
  | 'buffs'
  | 'rotation'
  | 'settings';

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

/**
 * Each lazy card's ready height, phone then from 768 (measured off the approved mocks; lanes B
 * and C report their measured heights and the integrator reconciles these). Literal classes,
 * because Tailwind reads them from this file.
 */
export const SECTION_RESERVE = {
  gear: 'min-h-[1719px] md:min-h-[889px]',
  talents: 'min-h-[661px] md:min-h-[593px]',
  resource: 'min-h-[711px] md:min-h-[448px]',
  buffs: 'min-h-[894px] md:min-h-[626px]',
  rotation: 'min-h-[1800px] md:min-h-[1177px]',
  settings: 'min-h-[1250px] md:min-h-[703px]',
} as const;

/** Whole-fight lanes on the timeline, for its skeleton's ready height (spec 3.7). */
export function timelineSkeletonHeight(result: SimResult): number {
  const TOOLBAR = 56;
  const AXIS = 30;
  const GROUP = 30;
  const LANE = 24;
  const RESOURCE = 72;
  const BRUSH = 44;
  const CAST_FOOT = 52;
  const CAST_ROW = 40;
  const COLLAPSED_ROWS = 5;
  // A periodic tick is not a cast: the ticks' lanes are off until the switch asks for them.
  const sample = (result.sample ?? []).filter((row) => row.outcome !== 'tick' && row.outcome !== 'tick_crit');
  const castLanes = new Set(
    sample
      .map((row) => row.action)
      .filter((key) => !key.startsWith('other:'))
      .map((key) => key.split('/')[0]),
  ).size;
  const playerAuras = new Set(
    (result.sample_auras ?? []).filter((row) => row.on === 'player').map((row) => row.action),
  ).size;
  const targetAuras = new Set(
    (result.sample_auras ?? []).filter((row) => row.on === 'target').map((row) => row.action),
  ).size;
  const groups = [castLanes, playerAuras, targetAuras].filter((count) => count > 0).length;
  return (
    TOOLBAR +
    AXIS +
    GROUP * groups +
    LANE * (castLanes + playerAuras + targetAuras) +
    RESOURCE +
    BRUSH +
    CAST_FOOT +
    CAST_ROW * COLLAPSED_ROWS
  );
}
