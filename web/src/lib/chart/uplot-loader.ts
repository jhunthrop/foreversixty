// web/src/lib/chart/uplot-loader.ts
// The chart chunk (controls and charts spec §B3, §B11): uPlot and the plot that drives it load as
// one lazy chunk, so report-island.js carries none of it. The report island starts the import
// when it mounts, in parallel with the summary fetch, so a chart never waits on it after its data
// arrives on a warm cache. A failed import is forgotten, so `Try again` re-imports.

export type PlotComponent = typeof import('../../components/chart/TimeSeriesPlot.svelte').default;

let pending: Promise<PlotComponent> | null = null;

export function loadChartChunk(): Promise<PlotComponent> {
  pending ??= import('../../components/chart/TimeSeriesPlot.svelte').then(
    (module) => module.default,
    (error: unknown) => {
      pending = null;
      throw error;
    },
  );
  return pending;
}
