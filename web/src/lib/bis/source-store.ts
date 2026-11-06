// web/src/lib/bis/source-store.ts
// The viewer's own Sources choice (web layout spec 2026-11-05 §5.11), remembered per browser in
// two buckets, leveling (below 60) and 60, because their defaults differ. A URL's ?sources= always
// wins; storage that throws or is empty simply means the band's default.
import { MAX_BAND } from './page-state';
import { SOURCE_FILTER_KINDS } from './source-filter';
import type { SourceFilterKind } from './types';

const KEY_PREFIX = 'fs.bis.sources.';

function keyFor(band: number): string {
  return `${KEY_PREFIX}${band < MAX_BAND ? 'leveling' : 'max'}`;
}

/** Both buckets' keys (the head script checks them before the island loads). */
export const SOURCES_STORAGE_KEYS: readonly string[] = [keyFor(0), keyFor(MAX_BAND)];

export function readStoredSources(
  band: number,
  storage: Pick<Storage, 'getItem'> | undefined = globalThis.localStorage,
): SourceFilterKind[] | undefined {
  try {
    const raw = storage?.getItem(keyFor(band));
    if (raw === null || raw === undefined) return undefined;
    const kinds = raw.split(',');
    return SOURCE_FILTER_KINDS.filter((kind) => kinds.includes(kind));
  } catch {
    return undefined;
  }
}

export function storeSources(
  band: number,
  sources: readonly SourceFilterKind[],
  storage: Pick<Storage, 'setItem'> | undefined = globalThis.localStorage,
): void {
  try {
    storage?.setItem(keyFor(band), sources.join(','));
  } catch {
    // Private mode or blocked storage: the choice still lives in the URL.
  }
}
