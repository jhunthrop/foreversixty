// web/src/lib/rating/report-ratings.svelte.ts
// A tiny reusable fetch-state container for GET /v1/reports/{id}/fights/{n}/ratings,
// shared by RatingPanel.svelte (the Summary dashboard row) and RatingTab.svelte (the full
// card) rather than each reimplementing the same load/loading/error/stale-response guard --
// see docs/superpowers/plans/2026-09-21-rating-web.md's Ruling 10 on why this is a shared
// hook and not a shared cross-component cache. Modeled on
// web/src/lib/report/lazy-component.svelte.ts's plain-state-container style, and on
// Character.svelte's own `resolved !== requested` stale-response guard.
import { fetchReportRatings } from '../rankings/api';
import type { ReportRatings } from './types';

export type ReportRatingsStatus = 'idle' | 'loading' | 'ready' | 'failed';

export interface ReportRatingsFetch {
  readonly data: ReportRatings | null;
  readonly status: ReportRatingsStatus;
  readonly error: string;
  /** Starts a fetch for this (reportId, fightIndex) pair unless one already succeeded for
   *  the same pair; safe to call on every render/effect run. */
  load(reportId: string, fightIndex: number): void;
}

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
          status = 'failed';
          error = thrown instanceof Error ? thrown.message : 'Ratings did not load.';
        });
    },
  };
}
