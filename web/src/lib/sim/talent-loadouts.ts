// web/src/lib/sim/talent-loadouts.ts
// TalentCandidates.svelte's own pure decisions, pulled out the same reason drop-picks.ts
// was: the component owns `picked`/`saved`/`exported` state, these just compute the next
// value from it.
import { bulkCopy } from './copy';
import { poolQualityCopy } from './pool-quality-copy';
import type { TalentLoadout } from './types';

/**
 * dps D31/E2: a loadout added through ADD A BUILD (a paste via the inline planner's
 * ImportBox, or a hand-built tree) is ticked into `picked` the moment it is accepted
 * (`store.addLoadout`), but TalentCandidates.svelte used to render only three named lists
 * -- the character's own build, the signed-in player's saved builds, and the addon
 * export's in-game loadouts -- so a picked loadout claimed by none of them had no checkbox
 * anywhere: nothing confirmed it was added, and nothing let a player untick it. The round-2
 * review reproduced exactly this after pasting a real alt build: the ranked table still
 * only offered "Your current build," because the newly-picked build was invisible, not
 * because picking it failed to work.
 *
 * This is every `picked` entry none of `own` or `named`'s lists claims by name, in the
 * order `picked` itself carries them -- so a signed-out player can paste a second build,
 * see it ticked, and run it beside their own without an account (contract for E2: "the
 * ranked list must work from pasted builds alone").
 */
export function customLoadouts(
  picked: readonly TalentLoadout[],
  own: TalentLoadout | null,
  named: readonly (readonly TalentLoadout[])[],
): TalentLoadout[] {
  const claimed = new Set(named.flat().map((entry) => entry.name));
  if (own !== null) claimed.add(own.name);
  return picked.filter((entry) => !claimed.has(entry.name));
}

/**
 * TalentCandidates.svelte, dps D31: signed out, `fetchMyBuilds` 401s the way every
 * signed-out read does (`account/api.ts`'s own 401-means-signed-out convention), and the
 * page used to show "Your saved builds could not be read; the rest of the page still
 * works." for that alongside every OTHER failure -- worded like a bug report for what is
 * actually the expected, permanent state for a visitor with no account. `signedOut` is
 * whether the failure was a 401/403; any other failure keeps the original message, since a
 * real read failure is still worth "could not be read."
 */
export function savedBuildsMessage(signedOut: boolean): string {
  return signedOut ? poolQualityCopy.talentsSavedSignedOut : bulkCopy.talentsSavedUnavailable;
}
