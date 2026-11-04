// web/src/lib/leveling/plans-file.ts
// /data/leveling-plans.json's loader, apart from leveling-plan.ts so the planner island can read
// the ladder orders without the home page's character helpers in its bundle.
import type { LevelingPlans } from './leveling-plan';

/** Where the plans are published. */
export const LEVELING_PLANS_URL = '/data/leveling-plans.json';

let plansPromise: Promise<LevelingPlans> | null = null;

/** The published plans, fetched once per page; an empty map when they cannot be read. */
export function loadLevelingPlans(fetcher: typeof fetch = fetch): Promise<LevelingPlans> {
  plansPromise ??= fetcher(LEVELING_PLANS_URL)
    .then((response) => (response.ok ? (response.json() as Promise<LevelingPlans>) : {}))
    .catch(() => {
      plansPromise = null;
      return {};
    });
  return plansPromise;
}
