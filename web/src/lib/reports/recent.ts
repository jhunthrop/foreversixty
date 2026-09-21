// web/src/lib/reports/recent.ts
// GET /v1/reports/recent: the public "Recent public reports" feed on /logs, and the
// compact panel wherever RecentReports.svelte is mounted. The route is public and
// unauthenticated -- it never answers for a signed-in visitor any differently than for
// a stranger -- so this goes through the shared transport (account/api.ts's
// requestEnvelope) with `credentials: 'omit'`, exactly the way lib/rankings/api.ts's
// public reads do, rather than account/api.ts's own session-carrying default. The
// AccountError-to-module-error wrapping below mirrors rankings/api.ts's own `get()`.
import { AccountError, requestEnvelope } from '../account/api';
import { API_BASE_URL } from '../planner/config';

export const RECENT_REPORTS_FAILED = 'Recent reports did not load';

export class RecentReportsError extends Error {
  constructor(
    message: string,
    readonly status: number,
  ) {
    super(message);
    this.name = 'RecentReportsError';
  }
}

/** One row of the public feed. The API already resolves title-or-zone server side. */
export interface RecentReport {
  id: string;
  title: string;
  created_at: string;
  fight_count: number;
  kill_count: number;
  /** Absent when the report has no guild. */
  guild_name?: string;
}

export interface RecentReportsPage {
  rows: RecentReport[];
  /** Present only when a further page may exist. */
  next_cursor?: string;
}

/**
 * Reads one page of the feed. cursor is the previous page's own next_cursor, passed back
 * verbatim; omitted, this reads the first page.
 */
export async function fetchRecentReports(
  cursor?: string,
  apiBase: string = API_BASE_URL,
): Promise<RecentReportsPage> {
  const query = cursor === undefined || cursor === '' ? '' : `?cursor=${encodeURIComponent(cursor)}`;
  let result;
  try {
    result = await requestEnvelope<RecentReportsPage>(`/v1/reports/recent${query}`, apiBase, {
      credentials: 'omit',
      failureMessage: RECENT_REPORTS_FAILED,
    });
  } catch (error) {
    if (error instanceof AccountError) throw new RecentReportsError(error.message, error.status);
    throw new RecentReportsError(RECENT_REPORTS_FAILED, 0);
  }
  return result.data ?? { rows: [] };
}
