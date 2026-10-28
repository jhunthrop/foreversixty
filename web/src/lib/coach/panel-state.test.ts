import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import {
  coachBusyReview,
  coachModelReview,
  coachRulesReview,
  coachTeaser,
  coachWritingReview,
} from '../../test-support/coach-fixtures';
import { CoachApiError, type fetchFightCoach } from './api';
import {
  MAX_POLLS,
  POLL_INTERVAL_MS,
  createCoachPanelState as create,
  type CoachPanelStateStore,
} from './panel-state.svelte';

const fetchMock = vi.fn<typeof fetchFightCoach>();
const createCoachPanelState = (): CoachPanelStateStore => create(undefined, fetchMock);

const flush = async (): Promise<void> => {
  await vi.advanceTimersByTimeAsync(0);
};

beforeEach(() => {
  vi.useFakeTimers();
  fetchMock.mockReset();
});
afterEach(() => vi.useRealTimers());

describe('createCoachPanelState', () => {
  it('starts loading, then settles on a ready review without polling', async () => {
    fetchMock.mockResolvedValue(coachModelReview());
    const store = createCoachPanelState();
    store.load('rep1', 3, 'k');
    expect(store.view.state).toBe('loading');
    await flush();
    expect(store.view.state).toBe('ready_model');
    expect(store.reserve).toBeNull();
    await vi.advanceTimersByTimeAsync(POLL_INTERVAL_MS * 3);
    expect(fetchMock).toHaveBeenCalledTimes(1);
  });

  it('writing: shows the rule-only review, polls every 4 s, swaps the model review in place', async () => {
    fetchMock.mockResolvedValueOnce(coachWritingReview()).mockResolvedValueOnce(coachModelReview());
    const store = createCoachPanelState();
    store.load('rep1', 3, 'k');
    await flush();
    expect(store.view.state).toBe('writing');
    expect(store.reserve).toEqual({ thenCount: 1, wellCount: 1, links: [false, false] });
    await vi.advanceTimersByTimeAsync(POLL_INTERVAL_MS);
    expect(store.view.state).toBe('ready_model');
    // The reserve is kept after the swap: nothing below moves.
    expect(store.reserve).not.toBeNull();
    await vi.advanceTimersByTimeAsync(POLL_INTERVAL_MS * 3);
    expect(fetchMock).toHaveBeenCalledTimes(2);
  });

  it('stops after 15 polls and reads late', async () => {
    fetchMock.mockResolvedValue(coachWritingReview());
    const store = createCoachPanelState();
    store.load('rep1', 3, 'k');
    await flush();
    await vi.advanceTimersByTimeAsync(POLL_INTERVAL_MS * (MAX_POLLS + 5));
    expect(fetchMock).toHaveBeenCalledTimes(1 + MAX_POLLS);
    expect(store.late).toBe(true);
    expect(store.view.state).toBe('writing');
  });

  it('busy does not poll', async () => {
    fetchMock.mockResolvedValue({ ...coachBusyReview(), status: 'writing' });
    const store = createCoachPanelState();
    store.load('rep1', 3, 'k');
    await flush();
    expect(store.view.state).toBe('busy');
    await vi.advanceTimersByTimeAsync(POLL_INTERVAL_MS * 2);
    expect(fetchMock).toHaveBeenCalledTimes(1);
  });

  it('teaser, 404 and failures map to their states', async () => {
    const store = createCoachPanelState();
    fetchMock.mockResolvedValueOnce(coachTeaser());
    store.load('rep1', 3, 'a');
    await flush();
    expect(store.view.state).toBe('teaser');
    fetchMock.mockResolvedValueOnce(null);
    store.load('rep1', 3, 'b');
    await flush();
    expect(store.view.state).toBe('not_available');
    fetchMock.mockRejectedValueOnce(new CoachApiError('x', 401));
    store.load('rep1', 3, 'c');
    await flush();
    expect(store.view.state).toBe('signed_out');
    fetchMock.mockRejectedValueOnce(new Error('boom'));
    store.load('rep1', 3, 'd');
    await flush();
    expect(store.view.state).toBe('error');
  });

  it('retry asks again, and a stale answer is dropped', async () => {
    let release: (value: ReturnType<typeof coachRulesReview>) => void = () => {};
    fetchMock.mockReturnValueOnce(new Promise((resolve) => (release = resolve)));
    const store = createCoachPanelState();
    store.load('rep1', 3, 'old');
    fetchMock.mockResolvedValueOnce(coachTeaser());
    store.load('rep1', 3, 'new');
    await flush();
    release(coachRulesReview());
    await flush();
    expect(store.view.state).toBe('teaser');
    fetchMock.mockResolvedValueOnce(coachModelReview());
    store.retry();
    await flush();
    expect(store.view.state).toBe('ready_model');
  });
});
