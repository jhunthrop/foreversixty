// web/src/lib/report/quiet.ts
// Whether the companion has gone quiet on a live report. GET /v1/reports/{id} carries
// `last_upload_at` (the newest fight upload, else creation) and `heartbeat_at` (the
// companion's last heartbeat while tailing); the newer of the two is the last time anyone
// heard from it. Older than QUIET_AFTER_MS on the page's clock, and "Logging" is no longer
// true. The clock is the caller's, re-read on every poll tick, so the sentence and the
// badge move without a reload.
import type { ReportMeta } from './types';

export const QUIET_AFTER_MS = 15 * 60 * 1000;

type Heard = Pick<ReportMeta, 'last_upload_at' | 'heartbeat_at'>;

/** The newest of the two instants, in ms; null when neither is set (an older API) or neither parses. */
export function lastHeardMs(meta: Heard | null): number | null {
  let newest: number | null = null;
  for (const stamp of [meta?.last_upload_at, meta?.heartbeat_at]) {
    if (!stamp) continue;
    const ms = Date.parse(stamp);
    if (Number.isNaN(ms)) continue;
    if (newest === null || ms > newest) newest = ms;
  }
  return newest;
}

/** Whole minutes since the companion was last heard, once that is past the threshold; null while it is fresh or unknown. */
export function quietMinutes(meta: Heard | null, nowMs: number): number | null {
  const heard = lastHeardMs(meta);
  if (heard === null) return null;
  const age = nowMs - heard;
  return age > QUIET_AFTER_MS ? Math.round(age / 60_000) : null;
}
