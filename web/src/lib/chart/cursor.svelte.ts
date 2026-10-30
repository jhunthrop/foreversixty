// web/src/lib/chart/cursor.svelte.ts
// The page's shared chart cursor (controls and charts spec §B6.3): the second under the pointer
// on any chart or lane view, keyed by fight, so every view on the page draws its crosshair at the
// same ms. Only the view that set it (`owner`) shows a tooltip.

export interface ChartCursor {
  syncKey: string;
  ms: number;
  /** The view under the pointer; the others draw the crosshair only. */
  owner: string;
}

const shared = $state<{ value: ChartCursor | null }>({ value: null });

export function chartCursor(): ChartCursor | null {
  return shared.value;
}

export function setChartCursor(next: ChartCursor | null): void {
  const current = shared.value;
  if (next === null && current === null) return;
  if (
    next !== null &&
    current !== null &&
    next.ms === current.ms &&
    next.owner === current.owner &&
    next.syncKey === current.syncKey
  )
    return;
  shared.value = next;
}

/** Clears the cursor only when this view owns it, so leaving one chart never wipes another's. */
export function clearChartCursor(owner: string): void {
  if (shared.value?.owner === owner) shared.value = null;
}

/** The cursor's ms for a view on `syncKey`, or null when the cursor is elsewhere. */
export function cursorMsFor(syncKey: string): number | null {
  const value = shared.value;
  return value !== null && value.syncKey === syncKey ? value.ms : null;
}
