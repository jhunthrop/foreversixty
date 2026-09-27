// @vitest-environment jsdom
// web/src/lib/islands/boot.test.ts
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import { FONT_WAIT_MS, scheduleBoot } from './boot';

const FAKED = ['setTimeout', 'clearTimeout', 'requestAnimationFrame', 'cancelAnimationFrame'] as const;

/** A `document.fonts` whose `ready` the test settles by hand. */
function fakeFonts(): { resolve: () => void } {
  let resolve = (): void => {};
  const ready = new Promise<void>((done) => {
    resolve = done;
  });
  Object.defineProperty(document, 'fonts', { value: { ready }, configurable: true });
  return { resolve };
}

describe('scheduleBoot', () => {
  // Animation frames are not in vitest's default fake list; the boot schedules two of them.
  beforeEach(() => vi.useFakeTimers({ toFake: FAKED }));
  afterEach(() => {
    vi.restoreAllMocks();
    vi.useRealTimers();
    Reflect.deleteProperty(document, 'fonts');
  });

  it('does not run the boot in the task that evaluated the module', async () => {
    const boot = vi.fn();
    scheduleBoot(boot, 'interactive');
    expect(boot).not.toHaveBeenCalled();
    await vi.runAllTimersAsync();
    expect(boot).toHaveBeenCalledTimes(1);
  });

  it('waits for the document to finish parsing, then still yields', async () => {
    const boot = vi.fn();
    scheduleBoot(boot, 'loading');
    await vi.runAllTimersAsync();
    expect(boot).not.toHaveBeenCalled();
    document.dispatchEvent(new Event('DOMContentLoaded'));
    expect(boot).not.toHaveBeenCalled();
    await vi.runAllTimersAsync();
    expect(boot).toHaveBeenCalledTimes(1);
  });

  it('lets a frame paint before the boot: two animation frames come first', async () => {
    const frames: FrameRequestCallback[] = [];
    vi.spyOn(window, 'requestAnimationFrame').mockImplementation((callback) => {
      frames.push(callback);
      return frames.length;
    });
    const boot = vi.fn();
    scheduleBoot(boot, 'interactive');
    await vi.advanceTimersByTimeAsync(0);
    // One frame is not enough: the callback runs before its own frame is painted.
    expect(frames).toHaveLength(1);
    frames[0](0);
    expect(frames).toHaveLength(2);
    frames[1](0);
    // Still on its own task after the second frame.
    expect(boot).not.toHaveBeenCalled();
    await vi.advanceTimersByTimeAsync(0);
    expect(boot).toHaveBeenCalledTimes(1);
  });

  it('waits for the fonts already loading, so the swap paints on the shell, not the island', async () => {
    const fonts = fakeFonts();
    const boot = vi.fn();
    scheduleBoot(boot, 'interactive');
    await vi.advanceTimersByTimeAsync(FONT_WAIT_MS - 1);
    expect(boot).not.toHaveBeenCalled();
    fonts.resolve();
    await vi.runAllTimersAsync();
    expect(boot).toHaveBeenCalledTimes(1);
  });

  it('does not let a slow font hold the island past the cap', async () => {
    fakeFonts();
    const boot = vi.fn();
    scheduleBoot(boot, 'interactive');
    await vi.advanceTimersByTimeAsync(FONT_WAIT_MS - 1);
    expect(boot).not.toHaveBeenCalled();
    await vi.runAllTimersAsync();
    expect(boot).toHaveBeenCalledTimes(1);
  });
});
