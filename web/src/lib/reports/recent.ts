// web/src/lib/reports/recent.ts
// GET /v1/reports/recent: the public "Recent public reports" feed on /logs, and the
// compact panel wherever RecentReports.svelte is mounted. The route is public and
// unauthenticated -- it never answers for a signed-in visitor any differently than for
// a stranger -- so this goes through the shared transport (account/api.ts's
// requestEnvelope) with `credentials: 'omit'`, exactly the way lib/rankings/api.ts's
// public reads do, rather than account/api.ts's own session-carrying default. The
// AccountError-to-module-error wrapping below mirrors rankings/api.ts's own `get()`.
import { AccountError, requestEnvelope } from '../account/api';
import { query } from '../data/query';
import { API_BASE_URL } from '../planner/config';

export const RECENT_REPORTS_FAILED = 'Recent reports did not load';

/** Spec §0/§3.2's "reports feed" class: public, 60 seconds. */
const RECENT_REPORTS_TTL_MS = 60_000;

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
  /** Distinct bosses killed; absent from an API before it. */
  boss_count?: number;
  /** Absent when the report has no guild. */
  guild_name?: string;
  /** The zone the report was logged in; the API sends it once the list carries it (A.6). */
  zone?: string;
  /** The credited character, as the single-report meta has it; null or absent when none. */
  logging_character?: string | { key?: string; name?: string } | null;
}

/**
 * Whether a report belongs on the public list: it names a zone or has a boss kill. A log with
 * neither (a test upload, a walk through a city) says nothing a visitor can read, and was the
 * first thing /logs showed under "Recent public reports": rows of "today · 0 kills".
 */
export function isPublicListable(report: Pick<RecentReport, 'zone' | 'kill_count'>): boolean {
  return report.kill_count > 0 || (report.zone ?? '').trim() !== '';
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
  const path = `/v1/reports/recent${cursor === undefined || cursor === '' ? '' : `?cursor=${encodeURIComponent(cursor)}`}`;
  return query<RecentReportsPage>(
    `${apiBase}${path}`,
    async () => {
      let result;
      try {
        result = await requestEnvelope<RecentReportsPage>(path, apiBase, {
          credentials: 'omit',
          failureMessage: RECENT_REPORTS_FAILED,
        });
      } catch (error) {
        if (error instanceof AccountError) throw new RecentReportsError(error.message, error.status);
        throw new RecentReportsError(RECENT_REPORTS_FAILED, 0);
      }
      return result.data ?? { rows: [] };
    },
    { scope: 'public', ttlMs: RECENT_REPORTS_TTL_MS },
  );
}

/**
 * How far below the viewport the panel starts loading. It sits under the fold, so waiting
 * for it to be fully on screen shows every visitor a "Loading" line first; one screen of
 * lead time means it is usually filled by the time it is scrolled to.
 */
export const RECENT_REPORTS_PRELOAD = '900px';
