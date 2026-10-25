// web/src/lib/sim/mark-seen-on-open.ts
// Opening a saved sim's page marks that server run seen (experience spec 2026-10-24 section
// 3.6), once, silently. Only a run in the member's own unseen list (`/v1/me`'s
// `finished_sims`) is marked: any other sim (a visitor's, someone else's, an already-seen
// one) makes no call. Every failure is ignored (the band's line simply stays).
import { fetchMeOnce, type Me } from '../account/api';
import { markSimSeen } from './api';

export async function markSimSeenOnOpen(
  simId: string,
  readMe: () => Promise<Me | null> = fetchMeOnce,
  markSeen: (simId: string) => Promise<void> = markSimSeen,
): Promise<void> {
  if (simId === '') return;
  try {
    const me = await readMe();
    if (!me?.finished_sims?.some((run) => run.sim_id === simId)) return;
    await markSeen(simId);
  } catch {
    /* no UI: a failed read or seen call leaves the line for the member to dismiss */
  }
}
