// web/src/lib/sim/bis-benchmark-load.ts
// Build time (node fs): every damage spec's published bands, cut to what bis-benchmark.ts
// matches a run against. The pages that draw a run (sim.astro, sim/[id].astro) hand it to the
// island as `data-bis-benchmarks`. A spec with no file or no run block is left out: without the
// run block nobody can say which fight its figures came from.
import { loadBisFile, readSpecCatalog } from '../bis/load';
import { presetOf } from '../bis/presets';
import type { BisFile, SpecCatalogEntry } from '../bis/types';
import type { BenchmarkBand, BisBenchmarks, SpecBenchmarks } from './bis-benchmark';

/** Figures are kept to a hundredth: the page shows one decimal. */
const KEPT = 100;
const keep = (value: number): number => Math.round(value * KEPT) / KEPT;

export function specBenchmarks(file: BisFile): SpecBenchmarks | null {
  if (file.run === undefined) return null;
  const { duration_sec, variation, targets, execute_ratio, target_level_offset } = file.run;
  const bands = file.bands
    .filter((band) => (band.role ?? 'dps') === 'dps' && band.set_dps_partial !== true && band.set_dps > 0)
    .map((band): BenchmarkBand => [
      band.band,
      band.faction,
      presetOf(band),
      keep(band.set_dps),
      keep(band.set_dps_error ?? 0),
    ]);
  return bands.length === 0
    ? null
    : { run: { duration_sec, variation, targets, execute_ratio, target_level_offset }, bands };
}

export function buildBisBenchmarks(
  catalog: readonly SpecCatalogEntry[],
  loadFile: (spec: string) => BisFile | null,
): BisBenchmarks {
  return Object.fromEntries(
    catalog
      .filter((entry) => entry.role === 'dps')
      .flatMap((entry) => {
        const file = loadFile(entry.spec);
        const benchmarks = file === null ? null : specBenchmarks(file);
        return benchmarks === null ? [] : [[entry.spec, benchmarks] as const];
      }),
  );
}

export function loadBisBenchmarks(build: string): BisBenchmarks {
  return buildBisBenchmarks(readSpecCatalog(), (spec) => loadBisFile(spec, build));
}
