// web/src/lib/sim/bis-benchmark.ts
// The run's DPS beside its spec's published best-in-slot set (the DPS spread's "Best-in-slot set"
// row). The table is build-time data (bis-benchmark-load.ts reads every damage spec's
// data/builds/<build>/bis/<spec>.json); this file only matches a run to one of its rows. The two
// figures are only comparable when the run fought the set's own fight: the spec, the level a band
// was ranked at, the faction, the preset, and the file's run block (length, variation, targets,
// execute share, target level). Anything else answers null and the row is not drawn.
import { factionOfRace } from '../race-faction';
import type { Faction } from '../bis/types';
import { confidenceBand } from './estimate';
import { defaultTargetLevel } from './settings';
import type { SimResult } from './types';

/** The ranker's run conditions (BisFile.run), the ones a comparable run must share. */
export interface BenchmarkRun {
  duration_sec: number;
  variation: number;
  targets: number;
  execute_ratio: number;
  target_level_offset: number;
}

/** One published band: [band level, faction, preset, set DPS, its standard error]. */
export type BenchmarkBand = readonly [number, Faction, 'raid' | 'bare', number, number];

export interface SpecBenchmarks {
  run: BenchmarkRun;
  bands: readonly BenchmarkBand[];
}

/** Spec key to its published bands. */
export type BisBenchmarks = Readonly<Record<string, SpecBenchmarks>>;

export interface BenchmarkMatch {
  band: number;
  faction: Faction;
  preset: 'raid' | 'bare';
  /** The set's DPS. */
  dps: number;
  /** The run's mean less the set's. */
  gap: number;
  /** The gap's 95% band: both runs' errors combined. */
  margin: number;
  /** True when the gap sits inside its margin. */
  withinNoise: boolean;
}

/** Fight fields compared as floats (variation, execute share). */
const EPSILON = 1e-9;

function sameFight(result: SimResult, run: BenchmarkRun): boolean {
  const { encounter, character } = result.request;
  // Absent, the engine fights its default: three above the character.
  const targetLevel = encounter.target_level ?? defaultTargetLevel(character.level);
  return (
    encounter.duration_sec === run.duration_sec &&
    Math.abs(encounter.variation - run.variation) < EPSILON &&
    encounter.targets === run.targets &&
    Math.abs(encounter.execute_ratio - run.execute_ratio) < EPSILON &&
    targetLevel === character.level + run.target_level_offset &&
    encounter.dummy !== true &&
    encounter.movement === undefined &&
    (encounter.targets_over_time ?? []).length === 0 &&
    encounter.target_armor === undefined
  );
}

/** The published band this run can be read against, or null when none is comparable. */
export function benchmarkFor(result: SimResult, table: BisBenchmarks): BenchmarkMatch | null {
  const spec = table[result.request.spec];
  const { character, preset } = result.request;
  const faction = factionOfRace(character.race);
  if (spec === undefined || faction === null || preset === undefined) return null;
  if (!sameFight(result, spec.run)) return null;
  // A raid band is buffed and fed; a run with none of either is not that fight.
  if (preset === 'raid' && (character.buffs.length === 0 || character.consumes.length === 0)) return null;
  const row = spec.bands.find(
    ([band, side, kind]) => band === character.level && side === faction && kind === preset,
  );
  if (row === undefined) return null;
  const [band, , , dps, error] = row;
  const gap = result.dps.mean - dps;
  const margin = confidenceBand({ ...result.dps, error: Math.hypot(result.dps.error, error) });
  return { band, faction, preset, dps, gap, margin, withinNoise: Math.abs(gap) <= margin };
}
