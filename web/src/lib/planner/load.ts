// web/src/lib/planner/load.ts
// Every planner fetch. The files live under /data/<build>/ so the browser caches them
// per build id and a data swap never serves stale trees.
import type { ReferenceData } from './store.svelte';
import type { ClassRow, Combo, ItemFile, ItemSet, RaceRow, TalentFile } from './types';

/** The one message the planner shows when data cannot be read; the UI adds a retry. */
export const DATA_LOAD_FAILED = 'Talent data did not load';

export function dataUrl(build: string, file: string): string {
  return `/data/${build}/${file}`;
}

export async function fetchJson<T>(url: string): Promise<T> {
  let response: Response;
  try {
    response = await fetch(url, { headers: { accept: 'application/json' } });
  } catch (cause) {
    throw new Error(DATA_LOAD_FAILED, { cause });
  }
  if (!response.ok) throw new Error(`${DATA_LOAD_FAILED} (${response.status} for ${url})`);
  return (await response.json()) as T;
}

export async function loadTalents(build: string, classSlug: string): Promise<TalentFile> {
  return fetchJson<TalentFile>(dataUrl(build, `talents/${classSlug}.json`));
}

export async function loadItems(build: string, classSlug: string): Promise<ItemFile> {
  return fetchJson<ItemFile>(dataUrl(build, `items/${classSlug}.json`));
}

/** Sets are optional: a build without normalized items ships no sets.json. */
export async function loadSets(build: string): Promise<ItemSet[]> {
  try {
    return await fetchJson<ItemSet[]>(dataUrl(build, 'sets.json'));
  } catch {
    return [];
  }
}

export async function loadReference(build: string): Promise<ReferenceData> {
  const [classes, races, combos] = await Promise.all([
    fetchJson<ClassRow[]>(dataUrl(build, 'classes.json')),
    fetchJson<RaceRow[]>(dataUrl(build, 'races.json')),
    fetchJson<Combo[]>(dataUrl(build, 'combos.json')),
  ]);
  return { classes, races, combos };
}
