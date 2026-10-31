// web/src/lib/chart/brush.ts
// Where a dragged brush edge lands (components/chart/OverviewBrush). Edges move in whole seconds,
// the tables' bucket, except at the fight's end: a 9.9 s pull's whole-fight window ends at
// 9.879 s, not at 9 s or 10 s.
import { BUCKET_MS } from '../report/window';

/** The nearest whole second to `ms`, or the fight's exact end when that second is at or past it. */
export function brushSnap(ms: number, domainMs: number): number {
  const snapped = Math.round(ms / BUCKET_MS) * BUCKET_MS;
  if (snapped >= domainMs) return domainMs;
  return Math.max(0, snapped);
}
