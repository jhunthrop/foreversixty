// web/src/lib/report/threat-taunts.ts
// The Threat tab's Taunts list (logs spec 2026-11-07 section 6.2, the rail's `Taunts`): the taunts
// in pull order, narrowed to the picked enemy, and the line that says what emptied the list. The
// table and the rail read the same rule, so both name the same enemy.
import { splitUnitName } from '../characters';
import type { Taunt, ThreatPair } from './types';

/** The url's `target` as an enemy name: a pair on that GUID (over a night, that name) names it. */
export function pickedEnemyName(pairs: readonly ThreatPair[], target: string): string | undefined {
  if (target === '') return undefined;
  const pair = pairs.find((candidate) => candidate.target_guid === target);
  return pair === undefined ? undefined : splitUnitName(pair.target_name).name;
}

/** The taunts in pull order; with an enemy picked, only the ones on it. */
export function orderedTaunts(taunts: readonly Taunt[], enemy: string | undefined): Taunt[] {
  return taunts
    .filter((taunt) => enemy === undefined || splitUnitName(taunt.target_name).name === enemy)
    .sort((a, b) => a.at_ms - b.at_ms);
}

/** The empty line names what emptied it: the source scope, the picked enemy, or the brush. */
export function noTauntsLine(options: {
  sourceName: string | undefined;
  enemy: string | undefined;
  windowed: boolean;
  scopeNoun: 'pull' | 'night';
}): string {
  const by = options.sourceName === undefined ? '' : ` by ${splitUnitName(options.sourceName).name}`;
  const on = options.enemy === undefined ? '' : ` on ${options.enemy}`;
  return `No taunts${by}${on} in this ${options.windowed ? 'window' : options.scopeNoun}.`;
}
