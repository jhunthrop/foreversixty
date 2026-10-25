// web/src/lib/sim/server-runs-rerun.ts
// What a `Run again` call comes to, as the status the failed row shows: started, lapsed
// (the plan no longer includes server sims, HTTP 402) or failed (anything else).
import { PREMIUM_REQUIRED_STATUS, SimApiError } from './api';
import type { RerunStatus } from './server-runs-poll';

export type RerunOutcome = Extract<RerunStatus, 'started' | 'failed' | 'lapsed'>;

export async function rerunOutcome(start: () => Promise<unknown>): Promise<RerunOutcome> {
  try {
    await start();
  } catch (thrown) {
    const lapsed = thrown instanceof SimApiError && thrown.status === PREMIUM_REQUIRED_STATUS;
    return lapsed ? 'lapsed' : 'failed';
  }
  return 'started';
}
