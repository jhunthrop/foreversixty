// web/src/lib/report/percentile.ts
// Parse percentiles for the rows the report shows. One API call per row is a lot of calls
// for a twenty-five-player raid, so: at most six at a time, cached for the life of the
// page by everything that changes the answer, and asked for only where the number means
// something -- an encounter fight, at the whole-fight window, on the tables that carry a
// role metric. ReportView enforces those three conditions; this module enforces the rest.
//
// A percentile that does not arrive leaves its row without one. It is a nice-to-have
// beside a number that is already correct, and a failed ranking service must never be the
// reason a table does not render.
import { API_BASE_URL } from '../planner/config';

export interface PercentileQuery {
  encounterId: number;
  difficulty: number;
  spec: string;
  phase: string;
  metric: string;
  value: number;
}

export function percentileKey(query: PercentileQuery): string {
  return [query.encounterId, query.difficulty, query.spec, query.phase, query.metric, query.value].join('|');
}

export interface PercentileLoader {
  load(queries: PercentileQuery[]): Promise<Map<string, number>>;
}

const CONCURRENCY = 6;

export function createPercentileLoader(apiBase: string = API_BASE_URL): PercentileLoader {
  // null remembers "nothing is ranked in that bracket": the API's 404 is as final as a
  // number for the life of the page, and re-asking on every tab switch tripped the limiter.
  const cache = new Map<string, number | null>();

  async function one(query: PercentileQuery): Promise<void> {
    const key = percentileKey(query);
    if (cache.has(key)) return;
    const params = new URLSearchParams({
      encounter: String(query.encounterId),
      difficulty: String(query.difficulty),
      spec: query.spec,
      phase: query.phase,
      metric: query.metric,
      value: String(query.value),
    });
    try {
      // A Request, not a bare URL string: the test asserts on `.url`, and passing a
      // Request here is no different for a real fetch implementation.
      const response = await fetch(new Request(`${apiBase}/v1/rankings/percentile?${params.toString()}`));
      if (response.status === 404) {
        cache.set(key, null);
        return;
      }
      if (!response.ok) return;
      const envelope = (await response.json()) as { ok: boolean; data: { percentile?: number } | null };
      const percentile = envelope.data?.percentile;
      if (envelope.ok && typeof percentile === 'number') cache.set(key, percentile);
    } catch {
      /* No percentile for this row. The row still renders. */
    }
  }

  return {
    async load(queries: PercentileQuery[]): Promise<Map<string, number>> {
      // Deduplicated by key, not just filtered against the cache: two identical queries in
      // the same call (a tank and a healer both parsed at the same rounded dps, say) would
      // otherwise both pass the "not cached yet" check before either request lands, and
      // the second would count as a second call.
      const seen = new Set<string>();
      const pending: PercentileQuery[] = [];
      for (const query of queries) {
        const key = percentileKey(query);
        if (cache.has(key) || seen.has(key)) continue;
        seen.add(key);
        pending.push(query);
      }
      // A fixed pool of workers pulling off one queue: simpler than batching, and it keeps
      // exactly CONCURRENCY requests in flight rather than CONCURRENCY per batch.
      const queue = [...pending];
      const workers = Array.from({ length: Math.min(CONCURRENCY, queue.length) }, async () => {
        for (let next = queue.shift(); next !== undefined; next = queue.shift()) await one(next);
      });
      await Promise.all(workers);

      const answers = new Map<string, number>();
      for (const query of queries) {
        const key = percentileKey(query);
        const found = cache.get(key);
        if (found !== undefined && found !== null) answers.set(key, found);
      }
      return answers;
    },
  };
}
