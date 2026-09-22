// @vitest-environment jsdom
// web/src/lib/islands/boot.test.ts
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import { scheduleBoot } from './boot';

describe('scheduleBoot', () => {
  beforeEach(() => vi.useFakeTimers());
  afterEach(() => vi.useRealTimers());

  it('does not run the boot in the task that evaluated the module', () => {
    const boot = vi.fn();
    scheduleBoot(boot, 'interactive');
    expect(boot).not.toHaveBeenCalled();
    vi.runAllTimers();
    expect(boot).toHaveBeenCalledTimes(1);
  });

  it('waits for the document to finish parsing, then still yields a task', () => {
    const boot = vi.fn();
    scheduleBoot(boot, 'loading');
    vi.runAllTimers();
    expect(boot).not.toHaveBeenCalled();
    document.dispatchEvent(new Event('DOMContentLoaded'));
    expect(boot).not.toHaveBeenCalled();
    vi.runAllTimers();
    expect(boot).toHaveBeenCalledTimes(1);
  });
});
