// web/src/lib/layout-reserve.ts
// The stored rendered height of a box that loads late (the server-runs list, the character
// History section): written after paint, read back by the pre-paint script in Base.astro so
// the next visit reserves the box before data. A convenience, never a requirement: storage
// that is missing, full or tampered with only costs the fallback height.

/** A stored reserve above this is nonsense (tampered or from another layout) and is ignored. */
export const MAX_RESERVE_PX = 2000;

function storage(): Storage | null {
  try {
    return globalThis.localStorage ?? null;
  } catch {
    return null;
  }
}

/** Stores a box height in px under `key`; a non-positive, non-finite or oversized value is dropped. */
export function writeReserve(key: string, pixels: number, max: number = MAX_RESERVE_PX): void {
  if (!Number.isFinite(pixels) || pixels <= 0 || pixels > max) return;
  try {
    storage()?.setItem(key, String(Math.round(pixels)));
  } catch {
    // Private mode or quota: the hint is a convenience.
  }
}
