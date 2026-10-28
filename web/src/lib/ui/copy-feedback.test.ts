import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import { COPIED_VISIBLE_MS, createCopyFeedback, type CopyFeedbackState } from './copy-feedback';

describe('createCopyFeedback', () => {
  beforeEach(() => vi.useFakeTimers());
  afterEach(() => vi.useRealTimers());

  function setup(writeText: (text: string) => Promise<void>) {
    const states: CopyFeedbackState[] = [];
    const feedback = createCopyFeedback(writeText, (state) => states.push(state));
    return { feedback, states };
  }

  it('writes the url, reports copied, then idle after two seconds', async () => {
    const writeText = vi.fn().mockResolvedValue(undefined);
    const { feedback, states } = setup(writeText);
    await feedback.copy('https://example.test/join/abc');
    expect(writeText).toHaveBeenCalledWith('https://example.test/join/abc');
    expect(states).toEqual(['copied']);
    vi.advanceTimersByTime(COPIED_VISIBLE_MS - 1);
    expect(states).toEqual(['copied']);
    vi.advanceTimersByTime(1);
    expect(states).toEqual(['copied', 'idle']);
  });

  it('a second copy restarts the two seconds', async () => {
    const { feedback, states } = setup(vi.fn().mockResolvedValue(undefined));
    await feedback.copy('a');
    vi.advanceTimersByTime(1500);
    await feedback.copy('a');
    vi.advanceTimersByTime(1500);
    expect(states).toEqual(['copied', 'copied']);
    vi.advanceTimersByTime(500);
    expect(states).toEqual(['copied', 'copied', 'idle']);
  });

  it('reports failed when the clipboard rejects, and never goes idle on a timer', async () => {
    const { feedback, states } = setup(vi.fn().mockRejectedValue(new Error('denied')));
    await feedback.copy('a');
    vi.advanceTimersByTime(COPIED_VISIBLE_MS * 2);
    expect(states).toEqual(['failed']);
  });

  it('dispose cancels a pending reset', async () => {
    const { feedback, states } = setup(vi.fn().mockResolvedValue(undefined));
    await feedback.copy('a');
    feedback.dispose();
    vi.advanceTimersByTime(COPIED_VISIBLE_MS * 2);
    expect(states).toEqual(['copied']);
  });
});
