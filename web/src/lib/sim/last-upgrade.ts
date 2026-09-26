// web/src/lib/sim/last-upgrade.ts
// The most recent Droptimizer top-row upgrade a run produced, in localStorage -- read back
// by the plain /sim results card (spec 2026-09-25 §6's after-sim sentence) with no network
// call: "nothing fetched anew, nothing invented." Written by Droptimizer.svelte the moment
// a /sim/drops run finishes with a real upgrade. Same try/catch/JSON pattern as
// current-character.ts: a convenience, never a source of truth, safe to import during
// SSR/prerender since nothing here touches `window`/`localStorage` at import time.
export interface LastUpgrade {
  itemName: string;
  sourceName: string;
  /** Already formatted, e.g. "+14 ± 3" -- combos.ts's own deltaLabel(). */
  gain: string;
  savedAt: string; // ISO
}

const STORAGE_KEY = 'fs.lastDroptimizerUpgrade';

function storageOf(storage: Storage | undefined): Storage | null {
  if (storage !== undefined) return storage;
  try {
    return globalThis.localStorage ?? null;
  } catch {
    return null;
  }
}

function isLastUpgrade(value: unknown): value is LastUpgrade {
  if (typeof value !== 'object' || value === null) return false;
  const candidate = value as Record<string, unknown>;
  return (
    typeof candidate.itemName === 'string' &&
    typeof candidate.sourceName === 'string' &&
    typeof candidate.gain === 'string' &&
    typeof candidate.savedAt === 'string'
  );
}

export function readLastUpgrade(storage?: Storage): LastUpgrade | null {
  const target = storageOf(storage);
  if (target === null) return null;
  try {
    const raw = target.getItem(STORAGE_KEY);
    if (raw === null) return null;
    const parsed: unknown = JSON.parse(raw);
    return isLastUpgrade(parsed) ? parsed : null;
  } catch {
    return null;
  }
}

export function writeLastUpgrade(value: LastUpgrade, storage?: Storage): void {
  const target = storageOf(storage);
  if (target === null) return;
  try {
    target.setItem(STORAGE_KEY, JSON.stringify(value));
  } catch {
    // Private browsing, quota exceeded, or a disabled storage API: this is a convenience,
    // so a failed write is silently skipped, same as current-character.ts's writeCurrent.
  }
}
