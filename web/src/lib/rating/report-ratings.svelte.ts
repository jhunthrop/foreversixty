// web/src/lib/rating/report-ratings.svelte.ts
// A tiny reusable fetch-state container for GET /v1/reports/{id}/fights/{n}/ratings,
// shared by RatingPanel.svelte (the Summary dashboard row) and RatingTab.svelte (the full
// card) rather than each reimplementing the same load/loading/error/stale-response guard --
// see docs/superpowers/plans/2026-09-21-rating-web.md's Ruling 10 on why this is a shared
// hook and not a shared cross-component cache. Modeled on
// web/src/lib/report/lazy-component.svelte.ts's plain-state-container style, and on
// Character.svelte's own `resolved !== requested` stale-response guard.
import { RankingsError, fetchReportRatings } from '../rankings/api';
import { isSampleReport } from '../report/sample';
import type { ReportRatings } from './types';

/** `none`: the API answered 404 "no ratings for that fight", which is by design for trash: ratings
 *  are made for boss encounters only. Not a failure and not styled as one.
 *  `flagged`: the API answered 404 with error code "flagged", a flagged report that has no
 *  ratings to show; not a failure (L8.6 fix round 6 item 5). Any other error is `failed`. */
export type ReportRatingsStatus = 'idle' | 'loading' | 'ready' | 'flagged' | 'none' | 'failed';

export interface ReportRatingsFetch {
  readonly data: ReportRatings | null;
  readonly status: ReportRatingsStatus;
  readonly error: string;
  /** Starts a fetch for this (reportId, fightIndex) pair unless one already succeeded for
   *  the same pair; safe to call on every render/effect run. */
  load(reportId: string, fightIndex: number): void;
}

/** The API's own message for a fight that has no ratings (api/internal/rating/handler.go). */
const NO_RATINGS_MESSAGE = 'no ratings for that fight';

export function createReportRatingsFetch(apiBase?: string): ReportRatingsFetch {
  let data = $state<ReportRatings | null>(null);
  let status = $state<ReportRatingsStatus>('idle');
  let error = $state('');
  let loadedKey = '';
  let requestedKey = '';

  return {
    get data() {
      return data;
    },
    get status() {
      return status;
    },
    get error() {
      return error;
    },
    load(reportId: string, fightIndex: number): void {
      const key = `${reportId}:${fightIndex}`;
      // The bundled sample has no database row, so there is nothing for the API to rate.
      if (isSampleReport(reportId)) {
        requestedKey = key;
        data = null;
        status = 'none';
        return;
      }
      if (key === loadedKey || (key === requestedKey && status === 'loading')) return;
      requestedKey = key;
      status = 'loading';
      error = '';
      void fetchReportRatings(reportId, fightIndex, apiBase)
        .then((result) => {
          if (requestedKey !== key) return;
          data = result;
          loadedKey = key;
          status = 'ready';
        })
        .catch((thrown: unknown) => {
          if (requestedKey !== key) return;
          data = null;
          if (thrown instanceof RankingsError && thrown.status === 404 && thrown.code === 'flagged') {
            status = 'flagged';
            return;
          }
          if (
            thrown instanceof RankingsError &&
            thrown.status === 404 &&
            thrown.code === 'not_found' &&
            thrown.message === NO_RATINGS_MESSAGE
          ) {
            status = 'none';
            return;
          }
          status = 'failed';
          error = thrown instanceof Error ? thrown.message : 'Ratings did not load.';
        });
    },
  };
}
