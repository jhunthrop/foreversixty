// web/src/lib/chart/legend-state.svelte.ts
// A chart's legend state for one page view (controls and charts spec §B6.4): which lines are
// toggled off, which one is highlighted, the phone's sheet and a touch scrub's second. Toggles
// are not written to the url; the window, the log scale and the Players choice are.
// The hidden set is replaced, never mutated: every toggle builds a new Set and assigns it, which is
// what makes the $state field react. A SvelteSet would add mutation tracking nothing here uses.
/* eslint-disable svelte/prefer-svelte-reactivity */
import type { ChartSeries } from './types';

export class LegendState {
  hidden = $state<ReadonlySet<string>>(new Set());
  highlightId = $state<string | null>(null);
  sheetOpen = $state(false);
  scrubMs = $state<number | null>(null);

  toggle(id: string): void {
    const next = new Set(this.hidden);
    if (next.has(id)) next.delete(id);
    else next.add(id);
    this.hidden = next;
  }

  /** That series alone; the always-drawn players and the medians (the reference) stay. */
  only(id: string, series: readonly ChartSeries[]): void {
    this.hidden = new Set(
      series
        .filter((line) => line.id !== id && !line.pinned && line.kind !== 'median')
        .map((line) => line.id),
    );
  }

  highlight(id: string | null): void {
    this.highlightId = id;
  }
}
