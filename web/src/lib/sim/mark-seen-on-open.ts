// web/src/lib/sim/mark-seen-on-open.ts
// Opening a saved sim's page marks that server run seen (experience spec 2026-10-24 section
// 3.6), once, silently. The page cannot tell whose sim it is, so it asks only members the
// seen call can apply to -- signed in with server sims -- and the API answers 404 for a run
// that is not theirs; every failure is ignored (the band's line simply stays).
import { effectiveServerSims, fetchMeOnce, type Me } from '../account/api';
import { markSimSeen } from './api';

export async function markSimSeenOnOpen(
  simId: string,
  readMe: () => Promise<Me | null> = fetchMeOnce,
  markSeen: (simId: string) => Promise<void> = markSimSeen,
): Promise<void> {
  if (simId === '') return;
  try {
    if (!effectiveServerSims(await readMe())) return;
    await markSeen(simId);
  } catch {
    /* no UI: a failed read or seen call leaves the line for the member to dismiss */
  }
}
