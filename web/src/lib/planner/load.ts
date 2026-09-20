// web/src/lib/planner/load.ts
// Every planner fetch. The files live under /data/<build>/ so the browser caches them
// per build id and a data swap never serves stale trees.
import type { ReferenceData } from './store.svelte';
import type { ClassRow, Combo, ItemFile, ItemSet, RaceRow, TalentFile } from './types';

/** The one message the planner shows when data cannot be read; the UI adds a retry. */
export const DATA_LOAD_FAILED = 'Talent data did not load';

/**
 * What every planner fetch throws. `status` is the HTTP status when the server answered at
 * all, and undefined when the request never got a response, so a caller can tell "this build
 * ships no such file" (404) apart from "this build is broken" (5xx, offline).
 */
export class DataLoadError extends Error {
  readonly status: number | undefined;

  constructor(message: string, options: { status?: number; cause?: unknown } = {}) {
    super(message, { cause: options.cause });
    this.name = 'DataLoadError';
    this.status = options.status;
  }
}

export function dataUrl(build: string, file: string): string {
  return `/data/${build}/${file}`;
}

export async function fetchJson<T>(url: string): Promise<T> {
  let response: Response;
  try {
    response = await fetch(url, { headers: { accept: 'application/json' } });
  } catch (cause) {
    throw new DataLoadError(DATA_LOAD_FAILED, { cause });
  }
  if (!response.ok) {
    throw new DataLoadError(`${DATA_LOAD_FAILED} (${response.status} for ${url})`, {
      status: response.status,
    });
  }
  return (await response.json()) as T;
}

export async function loadTalents(build: string, classSlug: string): Promise<TalentFile> {
  return fetchJson<TalentFile>(dataUrl(build, `talents/${classSlug}.json`));
}

export async function loadItems(build: string, classSlug: string): Promise<ItemFile> {
  return fetchJson<ItemFile>(dataUrl(build, `items/${classSlug}.json`));
}

/**
 * The shared shape of every optional build file: sets, loot, enchants, suffixes and
 * simbuffs are all "a build the data lane has not regenerated ships none of this," where
 * only a 404 means "absent" and every other failure (5xx, an unreachable network, a
 * malformed file) is a broken build, not an absent one, and is rethrown rather than
 * rendered as "nothing here."
 */
export async function loadOptional<T>(url: string, empty: T): Promise<T> {
  try {
    return await fetchJson<T>(url);
  } catch (error) {
    if (error instanceof DataLoadError && error.status === 404) return empty;
    throw error;
  }
}

/** Sets are optional: a build without normalized items ships no sets.json. */
export async function loadSets(build: string): Promise<ItemSet[]> {
  return loadOptional<ItemSet[]>(dataUrl(build, 'sets.json'), []);
}

export async function loadReference(build: string): Promise<ReferenceData> {
  const [classes, races, combos] = await Promise.all([
    fetchJson<ClassRow[]>(dataUrl(build, 'classes.json')),
    fetchJson<RaceRow[]>(dataUrl(build, 'races.json')),
    fetchJson<Combo[]>(dataUrl(build, 'combos.json')),
  ]);
  return { classes, races, combos };
}
