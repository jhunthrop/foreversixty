// web/src/lib/sim/bulk-server-run.ts
// The premium lane's poll loop, pulled out of bulk-store.svelte.ts (fix round 1) so that
// file stays under the 800-line cap with room for Task 17 still to add to it. This module
// owns no state of its own -- every decision that depends on the store's own fields
// (whether this call is still the current one, whether a prior result already exists) comes
// in as a callback, and every observation goes back out as one `ServerRunUpdate` the store
// applies to its own `$state`.
import { dispatchServerSim, fetchBulkProgress, fetchSim, SimApiError } from './api';
import type { BulkProgress } from './bulk-run';
import type { BulkRequest, WeightsRequest } from './bulk-types';
import { bulkCopy } from './copy';
import type { SimResult } from './types';

/**
 * The poll's own ceiling, in attempts rather than elapsed time so a test can reach it
 * deterministically. ~15 minutes at the default 2-second interval -- the job's own timeout
 * (Task 5's report: "a 20,000-combination fast run cannot finish inside the job's 15-minute
 * timeout"). Without this, an API that never answers `done`/`error` polls forever and
 * `serverRunning` stays true indefinitely (fix round 1, Important 1).
 */
export const MAX_SERVER_POLLS = 450;

export interface ServerRunOptions {
  apiBase?: string;
  pollMs: number;
  pollLimit: number;
  /** False once a newer `run()`/`runOnServer()`/`loadAddon()` etc. has superseded this call. */
  isCurrent(): boolean;
  /** Whether the store already holds a completed result, for the done-vs-error landing. */
  hasPriorResult(): boolean;
}

export type ServerRunUpdate =
  | { kind: 'progress'; progress: BulkProgress }
  | { kind: 'message'; message: string }
  | { kind: 'failed'; message: string; phase: 'done' | 'error' }
  | { kind: 'done'; result: SimResult };

const delay = (ms: number): Promise<void> => new Promise((resolve) => setTimeout(resolve, ms));

/** A landing when the poll cannot go on: keeps a partial result on screen rather than losing it. */
function landing(hasPriorResult: () => boolean): 'done' | 'error' {
  return hasPriorResult() ? 'done' : 'error';
}

/**
 * Dispatches the envelope to `POST /v1/sims/run`, then polls `GET /v1/sims/<id>/progress`
 * until it answers `done` (one more read, `GET /v1/sims/<id>`, for the full result) or
 * `error`, or until `options.pollLimit` is reached. Every observation is reported through
 * `onUpdate`, never written directly -- this function holds no `$state` of its own.
 */
export async function runServerJob(
  request: BulkRequest | WeightsRequest,
  options: ServerRunOptions,
  onUpdate: (update: ServerRunUpdate) => void,
): Promise<void> {
  let simId: string;
  try {
    simId = await dispatchServerSim(request, options.apiBase);
  } catch (error) {
    if (options.isCurrent()) {
      onUpdate({
        kind: 'message',
        message: error instanceof SimApiError ? error.message : bulkCopy.bulkFailed,
      });
    }
    return;
  }

  // The API's own `SimProgress` carries no total stage count (only the current `stage`), so
  // this is carried forward locally rather than read back off the store's own `progress` --
  // decoupling this loop from the store's state entirely.
  let stages = 1;

  for (let attempt = 0; ; attempt += 1) {
    if (attempt >= options.pollLimit) {
      if (options.isCurrent()) {
        onUpdate({ kind: 'failed', message: bulkCopy.bulkFailed, phase: landing(options.hasPriorResult) });
      }
      return;
    }
    await delay(options.pollMs);
    if (!options.isCurrent()) return;

    let row;
    try {
      row = await fetchBulkProgress(simId, options.apiBase);
    } catch (error) {
      if (options.isCurrent()) {
        onUpdate({
          kind: 'failed',
          message: error instanceof SimApiError ? error.message : bulkCopy.bulkFailed,
          phase: landing(options.hasPriorResult),
        });
      }
      return;
    }
    if (!options.isCurrent()) return;

    if (row.stage !== undefined && row.combos_total !== undefined) {
      if (row.combos_total === 0) stages = 1;
      onUpdate({
        kind: 'progress',
        progress: {
          stage: row.stage,
          stages,
          combosDone: row.combos_done ?? 0,
          combosTotal: row.combos_total,
        },
      });
    }

    if (row.state === 'error') {
      if (options.isCurrent()) {
        onUpdate({ kind: 'failed', message: bulkCopy.bulkFailed, phase: landing(options.hasPriorResult) });
      }
      return;
    }

    if (row.state === 'done') {
      try {
        const finished = await fetchSim(simId, options.apiBase);
        if (options.isCurrent()) onUpdate({ kind: 'done', result: finished });
      } catch (error) {
        if (options.isCurrent()) {
          onUpdate({
            kind: 'failed',
            message: error instanceof SimApiError ? error.message : bulkCopy.bulkFailed,
            phase: landing(options.hasPriorResult),
          });
        }
      }
      return;
    }
  }
}
