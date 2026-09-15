// web/src/lib/report/tree-sizes.ts
// How many deaths and combatant rows can name a build link's talent tree sizes, pulled out
// of ReportView.svelte because a Svelte component's own $effect cannot be unit tested here
// (vitest.config.ts hands mount() Svelte's server entry, which throws
// lifecycle_function_unavailable) -- the same reason planner-link.ts exists as a plain
// module. ReportView.svelte's effect is a thin wrapper: gather the classes the fight's
// roster names, call this, merge the answer into its treeSizes cache.
import { DataLoadError } from '../planner/load';
import type { TalentFile } from '../planner/types';

export function classSlugFromName(className: string): string {
  return className.toLowerCase().replace(/\s+/g, '-');
}

/**
 * Resolves tree sizes (talent count per tree) for every name in `wanted` that is not
 * already in `known`, using `loadTalents` -- injected rather than imported directly, so this
 * stays testable without a network.
 *
 * A class the build genuinely ships no talent file for (404) resolves to an empty list:
 * that answer is permanent, and the caller is meant to cache it, which is what makes a
 * death or combatant link for that class fall back to gear-only.
 *
 * Any other failure -- a network blip, a 5xx, a malformed file -- is not an answer about
 * the class at all, and is left out of the result rather than resolved to `[]`. Returning
 * it as "no data" here would let the caller cache a permanent wrong answer that a retry can
 * never overwrite, because the caller's own "already known" check would then treat the
 * class as settled forever. Leaving it out keeps the class absent from the cache, so the
 * next call (the next time ReportView.svelte's effect runs) tries again.
 */
export async function resolveTreeSizes(
  wanted: readonly string[],
  known: ReadonlySet<string>,
  loadTalents: (build: string, classSlug: string) => Promise<TalentFile>,
  build: string,
): Promise<Array<readonly [string, number[]]>> {
  const missing = wanted.filter((name) => !known.has(name));
  const results = await Promise.all(
    missing.map(async (name): Promise<readonly [string, number[]] | null> => {
      try {
        const file = await loadTalents(build, classSlugFromName(name));
        return [name, file.trees.map((tree) => tree.talents.length)] as const;
      } catch (error) {
        if (error instanceof DataLoadError && error.status === 404) return [name, []] as const;
        return null;
      }
    }),
  );
  return results.filter((entry): entry is readonly [string, number[]] => entry !== null);
}
