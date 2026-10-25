// web/src/lib/account/opening-watchdog.test.ts
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import { createOpeningWatchdog, OPENING_TIMEOUT_MS } from './opening-watchdog';

describe('createOpeningWatchdog', () => {
  beforeEach(() => vi.useFakeTimers());
  afterEach(() => vi.useRealTimers());

  it('waits ten seconds', () => {
    expect(OPENING_TIMEOUT_MS).toBe(10_000);
  });

  it('fires once when the page is still here after the timeout', () => {
    const onStuck = vi.fn();
    createOpeningWatchdog(onStuck).start();
    vi.advanceTimersByTime(OPENING_TIMEOUT_MS - 1);
    expect(onStuck).not.toHaveBeenCalled();
    vi.advanceTimersByTime(1);
    expect(onStuck).toHaveBeenCalledTimes(1);
    vi.advanceTimersByTime(OPENING_TIMEOUT_MS * 2);
    expect(onStuck).toHaveBeenCalledTimes(1);
  });

  it('does not fire after cancel (the page left or came back)', () => {
    const onStuck = vi.fn();
    const watchdog = createOpeningWatchdog(onStuck);
    watchdog.start();
    watchdog.cancel();
    vi.advanceTimersByTime(OPENING_TIMEOUT_MS * 2);
    expect(onStuck).not.toHaveBeenCalled();
  });

  it('restarts the countdown when started again', () => {
    const onStuck = vi.fn();
    const watchdog = createOpeningWatchdog(onStuck);
    watchdog.start();
    vi.advanceTimersByTime(OPENING_TIMEOUT_MS - 1);
    watchdog.start();
    vi.advanceTimersByTime(OPENING_TIMEOUT_MS - 1);
    expect(onStuck).not.toHaveBeenCalled();
    vi.advanceTimersByTime(1);
    expect(onStuck).toHaveBeenCalledTimes(1);
  });
});
