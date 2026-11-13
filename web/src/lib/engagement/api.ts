// web/src/lib/engagement/api.ts
// The owner's activity figures (GET /v1/admin/engagement, api/internal/engagement): daily and
// weekly active accounts, new against returning, the surface split and retention by signup week.
import { requestEnvelope } from '../account/api';
import { API_BASE_URL } from '../planner/config';

export const ADMIN_ENGAGEMENT_PATH = '/v1/admin/engagement';

/** One UTC day or one Monday-to-Sunday week. */
export interface EngagementPeriod {
  start: string;
  active: number;
  new: number;
  returning: number;
  site: number;
  addon: number;
  app: number;
  /** The period begins before the live record: rebuilt from history, a floor. */
  partial: boolean;
}

export interface CohortWeek {
  active: number;
  partial: boolean;
}

/** The accounts that signed up in one week; weeks[k] is week N+k+1. */
export interface EngagementCohort {
  week: string;
  joined: number;
  weeks: CohortWeek[];
}

export interface EngagementReport {
  today: string;
  live_since: string | null;
  internal_accounts: number;
  /** Newest first. */
  days: EngagementPeriod[];
  /** Newest first. */
  weeks: EngagementPeriod[];
  /** Oldest first. */
  cohorts: EngagementCohort[];
}

/** The figures (the owner only: 401 and 403 throw). */
export async function fetchEngagement(apiBase: string = API_BASE_URL): Promise<EngagementReport | null> {
  const { data } = await requestEnvelope<EngagementReport>(ADMIN_ENGAGEMENT_PATH, apiBase);
  return data ?? null;
}
