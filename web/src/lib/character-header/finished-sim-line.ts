// web/src/lib/character-header/finished-sim-line.ts
// What the band's "your server sim finished" line shows (experience spec 2026-10-24 section
// 3.4): who sees it, which runs it names, and the seen-call fan-out behind its dismiss.
import type { FinishedSim, Me } from '../account/api';

/** How long the "Could not dismiss that" text replaces the sentence before it returns. */
export const DISMISS_FAILURE_MS = 4_000;

/**
 * The unseen finished runs the line names, newest first; null when the line is absent: no
 * server-sim entitlement, no `finished_sims` (an older API sends none) or an empty list.
 */
export function finishedRunsFor(me: Me | null): readonly [FinishedSim, ...FinishedSim[]] | null {
  if (me?.entitlements?.server_sims !== true) return null;
  const [newest, ...older] = me.finished_sims ?? [];
  return newest === undefined ? null : [newest, ...older];
}

/** Posts the seen call for every run the line names; true only when all of them answered. */
export async function markAllSeen(
  runs: readonly FinishedSim[],
  markSeen: (simId: string) => Promise<void>,
): Promise<boolean> {
  const outcomes = await Promise.allSettled(runs.map((run) => markSeen(run.sim_id)));
  return outcomes.every((outcome) => outcome.status === 'fulfilled');
}
