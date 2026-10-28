// web/src/lib/coach/panel-state.svelte.ts
// The Coach tab's fetch state. `load` asks for one player's review of one fight; while the answer
// is `writing` it asks again every 4 s, at most 15 times, and swaps the model's review in place
// (the panel keeps its reserved slots, see reserves.ts). A request that fails settles the state
// on `error` or `signed_out`; `retry` asks again. A late answer for an older request is dropped.
import { fetchFightCoach } from './api';
import { mergeReserve, reserveFor, type Reserve } from './reserves';
import { viewForAnswer, viewForFailure, viewOf, type CoachPanelView } from './states';
import type { CoachFull, CoachTeaser } from './types';

export const POLL_INTERVAL_MS = 4000;
export const MAX_POLLS = 15;

export type CoachAnswer = CoachFull | CoachTeaser;

export interface CoachPanelStateStore {
  readonly data: CoachAnswer | null;
  readonly view: CoachPanelView;
  /** True once the poll budget is spent and the answer is still `writing`. */
  readonly late: boolean;
  /** The slots a writing review is held to for this visit; null for a review that arrived ready. */
  readonly reserve: Reserve | null;
  load(reportId: string, fightIndex: number, playerKey: string): void;
  retry(): void;
  /** One poll step, exposed so tests drive the sequence without the clock. */
  poll(): Promise<void>;
  dispose(): void;
}

interface Request {
  reportId: string;
  fightIndex: number;
  playerKey: string;
}

export type CoachFetcher = typeof fetchFightCoach;

export function createCoachPanelState(
  apiBase?: string,
  fetcher: CoachFetcher = fetchFightCoach,
): CoachPanelStateStore {
  let data = $state<CoachAnswer | null>(null);
  let view = $state<CoachPanelView>(viewOf('loading'));
  let late = $state(false);
  let reserve = $state<Reserve | null>(null);
  let request: Request | null = null;
  let polls = 0;
  let timer: ReturnType<typeof setTimeout> | null = null;

  const stop = (): void => {
    if (timer !== null) clearTimeout(timer);
    timer = null;
  };

  const settle = (answer: CoachAnswer | null): void => {
    data = answer;
    view = viewForAnswer(answer);
    if (answer !== null && answer.access !== 'teaser') {
      if (answer.status === 'writing' || reserve !== null) {
        reserve = mergeReserve(reserve, reserveFor(answer.review));
      }
    }
    if (answer !== null && answer.access !== 'teaser' && answer.status === 'writing' && !answer.busy) {
      schedule();
    }
  };

  const schedule = (): void => {
    stop();
    if (polls >= MAX_POLLS) {
      late = true;
      return;
    }
    timer = setTimeout(() => void poll(), POLL_INTERVAL_MS);
  };

  async function poll(): Promise<void> {
    const current = request;
    if (current === null) return;
    polls += 1;
    try {
      const answer = await fetcher(current.reportId, current.fightIndex, current.playerKey, apiBase);
      if (request !== current) return;
      settle(answer);
    } catch {
      // A failed poll keeps the rule-only review on screen and tries again within the budget.
      if (request === current) schedule();
    }
  }

  async function fetchFirst(current: Request): Promise<void> {
    try {
      const answer = await fetcher(current.reportId, current.fightIndex, current.playerKey, apiBase);
      if (request === current) settle(answer);
    } catch (error: unknown) {
      if (request !== current) return;
      data = null;
      view = viewForFailure(error);
    }
  }

  function start(next: Request): void {
    stop();
    request = next;
    polls = 0;
    late = false;
    reserve = null;
    data = null;
    view = viewOf('loading');
    void fetchFirst(next);
  }

  return {
    get data() {
      return data;
    },
    get view() {
      return view;
    },
    get late() {
      return late;
    },
    get reserve() {
      return reserve;
    },
    load(reportId, fightIndex, playerKey) {
      const same =
        request?.reportId === reportId &&
        request.fightIndex === fightIndex &&
        request.playerKey === playerKey;
      if (!same) start({ reportId, fightIndex, playerKey });
    },
    retry() {
      if (request !== null) start(request);
    },
    poll,
    dispose: stop,
  };
}
