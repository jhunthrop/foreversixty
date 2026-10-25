// web/src/lib/sim/server-runs-controller.svelte.ts
// The reactive side of "Your server runs": reads the list, re-reads every ten seconds while a
// run is queued or running and the tab is visible, marks a run seen at once when it finishes
// under the member's eyes, and keeps the two localStorage hints current (spec 2026-10-24
// sections 3.5 and 3.7). SimView owns one; ServerRuns.svelte only draws what it holds.
import { invalidate } from '../data/query';
import { API_BASE_URL } from '../planner/config';
import { listMySims, markSimSeen } from './api';
import {
  POLL_INTERVAL_MS,
  RUN_START_REFRESH_MS,
  hasInFlight,
  markLive,
  newlyFinished,
  withSeen,
} from './server-runs-poll';
import type { SimListRow } from './types';

/** Every `mine=1` list shares this prefix; the same one saveSim clears after a save. */
const MY_SIMS_PREFIX = `${API_BASE_URL}/v1/sims?mine=1`;
const SERVER_LANE = 'server';

function isVisible(): boolean {
  return typeof document === 'undefined' || document.visibilityState === 'visible';
}

export function createServerRunsController() {
  let rows = $state<SimListRow[] | null>(null);
  let failed = $state(false);
  let nowMs = $state(Date.now());
  let timer: ReturnType<typeof setTimeout> | undefined;
  let started = false;

  function clearTimer(): void {
    if (timer !== undefined) clearTimeout(timer);
    timer = undefined;
  }

  function schedule(delayMs: number): void {
    clearTimer();
    timer = setTimeout(() => void load(), delayMs);
  }

  /** Watched finishes are seen: the member saw it happen, so no dot and no band line follow. */
  async function markWatched(finished: SimListRow[]): Promise<void> {
    if (finished.length === 0 || !isVisible()) return;
    const results = await Promise.allSettled(finished.map((row) => markSimSeen(row.sim_id)));
    const done = finished
      .filter((_, index) => results[index]?.status === 'fulfilled')
      .map((row) => row.sim_id);
    if (done.length === 0) return;
    invalidate(MY_SIMS_PREFIX);
    if (rows !== null) rows = withSeen(rows, done);
  }

  async function load(): Promise<void> {
    clearTimer();
    invalidate(MY_SIMS_PREFIX);
    try {
      const page = await listMySims(1, undefined, 'all', SERVER_LANE);
      const before = rows;
      rows = page.rows;
      failed = false;
      nowMs = Date.now();
      if (hasInFlight(page.rows)) markLive();
      await markWatched(newlyFinished(before, page.rows));
    } catch {
      if (rows === null) failed = true;
    }
    if (started && hasInFlight(rows) && isVisible()) schedule(POLL_INTERVAL_MS);
  }

  function onVisibility(): void {
    if (isVisible() && hasInFlight(rows)) void load();
    else clearTimer();
  }

  return {
    get rows() {
      return rows;
    },
    /** The first read failed and there is nothing to show; a later failed poll keeps the rows. */
    get failed() {
      return failed;
    },
    /** Epoch ms of the last read: the clock the rows' ages and overdue test are measured on. */
    get nowMs() {
      return nowMs;
    },
    /** First read plus the visibility listener; returns the stop function for cleanup. */
    start(): () => void {
      started = true;
      document.addEventListener('visibilitychange', onVisibility);
      void load();
      return () => {
        started = false;
        clearTimer();
        document.removeEventListener('visibilitychange', onVisibility);
      };
    },
    reload: load,
    /** This page started a server run: remember it and look for its row shortly. */
    runStarted(): void {
      markLive();
      schedule(RUN_START_REFRESH_MS);
    },
    /** The run this page started has ended: its row is final now. */
    runEnded(): void {
      void load();
    },
  };
}

export type ServerRunsController = ReturnType<typeof createServerRunsController>;
