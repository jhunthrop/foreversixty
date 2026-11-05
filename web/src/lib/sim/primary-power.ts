// web/src/lib/sim/primary-power.ts
// The power a saved run's Resource card, nav and cap note name (spec 2026-11-04-sim-results 3.10):
// the class's own, never a track's number. Warriors and bear druids run on rage, rogues and cat
// druids on energy, everyone else on mana. The track's `power_type` only picks which track to read.
//
// Results saved before 2026-09-21 (sim/adapter's engineResourceTypeToLogPowerType) stored the
// engine's own resource enum, where rage is 3 and energy 2, the combat log's energy and focus. A
// warrior's rage track on those reads 3, so a run is matched on the log's number first and on the
// engine's old number second, and the name always comes from the class.
import type { ResourceTrack } from '../report/types';
import type { SimResult } from './types';

export type PrimaryPower = 'rage' | 'energy' | 'mana';

const CLASS_POWER: Readonly<Record<string, PrimaryPower>> = { warrior: 'rage', rogue: 'energy' };
const SPEC_POWER: Readonly<Record<string, PrimaryPower>> = {
  'druid-feral': 'energy',
  'druid-feral-bear': 'rage',
};

/** The combat log's power type (what a result stamps today). */
const LOG_POWER_TYPE: Readonly<Record<PrimaryPower, number>> = { mana: 0, rage: 1, energy: 3 };
/** The engine's own resource enum, which results saved before 2026-09-21 carry. */
const OLD_ENGINE_POWER_TYPE: Readonly<Record<PrimaryPower, number>> = { mana: 1, rage: 3, energy: 2 };

/** The power a spec fights on: by spec where the class has forms, else by class. */
export function primaryPower(spec: string): PrimaryPower {
  const classSlug = spec.split('-')[0] ?? '';
  return SPEC_POWER[spec] ?? CLASS_POWER[classSlug] ?? 'mana';
}

export interface PowerTrack {
  power: PrimaryPower;
  track: ResourceTrack;
}

/** The run's track for its class's power; null when the run reports none it can be matched to. */
export function primaryPowerTrack(result: SimResult): PowerTrack | null {
  const power = primaryPower(result.request.spec);
  const tracks = result.summary.resources;
  const track =
    tracks.find((row) => row.power_type === LOG_POWER_TYPE[power]) ??
    tracks.find((row) => row.power_type === OLD_ENGINE_POWER_TYPE[power]) ??
    (tracks.length === 1 ? tracks[0] : undefined);
  return track === undefined ? null : { power, track };
}
