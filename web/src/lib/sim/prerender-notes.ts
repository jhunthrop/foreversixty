// web/src/lib/sim/prerender-notes.ts
// Worth knowing for a prerendered saved run (sim/[id].astro, the fixture ids Lighthouse and the
// e2e measure): the same warnings RunReport computes in the browser (lib/sim/warnings.ts), read at
// build time from the synced public/data files through the page's own loaders, so the lines paint
// with the page instead of arriving above it once the names load (a shift, and the late LCP).
// Build time only: it answers the loaders' /data/ reads from disk for the length of one call.
import { readFile } from 'node:fs/promises';
import path from 'node:path';
import { loadActionNames, withShardNames } from './action-names';
import { summaryActionKeys } from './sentence';
import { simSpellBundle } from './spell-bundle';
import { loadSpellInfo } from './spell-info';
import type { SimResult } from './types';
import { simWarnings, type SimWarning } from './warnings';

/** A warning with its leading icon resolved, ready to draw before any data loads. */
export interface PrerenderedNote extends SimWarning {
  iconUrl: string | null;
}

async function withDiskFetch<T>(publicDir: string, run: () => Promise<T>): Promise<T> {
  const original = globalThis.fetch;
  globalThis.fetch = (async (input: RequestInfo | URL) => {
    const url = typeof input === 'string' ? input : input instanceof URL ? input.pathname : input.url;
    // The class's spell facts are a build-time endpoint (pages/data/[build]/sim-spells), not a
    // synced file: answered from the same builder.
    const bundle = /^\/data\/([^/]+)\/sim-spells\/([a-z-]+)\.json$/.exec(url);
    if (bundle !== null) {
      return new Response(JSON.stringify(simSpellBundle(bundle[1], bundle[2])), { status: 200 });
    }
    try {
      const body = await readFile(path.join(publicDir, decodeURIComponent(url.split('?')[0])), 'utf8');
      return new Response(body, { status: 200, headers: { 'content-type': 'application/json' } });
    } catch {
      return new Response('', { status: 404 });
    }
  }) as typeof fetch;
  try {
    return await run();
  } finally {
    globalThis.fetch = original;
  }
}

export async function prerenderedNotes(
  result: SimResult,
  build: string,
  publicDir: string,
): Promise<PrerenderedNote[]> {
  const classSlug = result.request.character.class;
  return withDiskFetch(publicDir, async () => {
    // A file the build does not ship (a fixture build's names) reads as it does in the browser:
    // RunReport keeps going without it.
    const [loaded, spells] = await Promise.all([
      loadActionNames(build, classSlug).catch(() => null),
      loadSpellInfo(build, classSlug),
    ]);
    const names = loaded === null ? null : await withShardNames(loaded, summaryActionKeys(result.summary));
    return simWarnings({ result, names, fact: spells.fact }).map((note) => ({
      ...note,
      iconUrl: note.actionKey === undefined ? null : spells.iconUrl(note.actionKey),
    }));
  });
}
