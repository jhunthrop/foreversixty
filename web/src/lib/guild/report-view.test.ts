// web/src/lib/guild/report-view.test.ts
// The copy state machine and the report URL.
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import { attendanceCopy } from './attendance-copy';
import { CopyState } from '../ui/copy-state.svelte';
import { COPIED_VISIBLE_MS } from '../ui/copy-feedback';
import { reportUrlOf } from './report-view';

describe('reportUrlOf', () => {
  it('is the raid page with the report anchor', () => {
    expect(reportUrlOf({ origin: 'https://foreversixty.gg', pathname: '/guild/a/b/c/raids/7' })).toBe(
      'https://foreversixty.gg/guild/a/b/c/raids/7#report',
    );
  });
});

describe('Copy for Discord (CopyState)', () => {
  beforeEach(() => vi.useFakeTimers());
  afterEach(() => vi.useRealTimers());

  it('says Copied, then goes quiet after the visible time', async () => {
    const writeText = vi.fn(() => Promise.resolve());
    const copy = new CopyState(writeText);
    await copy.copy('text');
    expect(copy.state).toBe('copied');
    expect(writeText).toHaveBeenCalledWith('text');
    vi.advanceTimersByTime(COPIED_VISIBLE_MS);
    expect(copy.state).toBe('idle');
  });

  it('keeps the text for the fallback when the clipboard refuses', async () => {
    const copy = new CopyState(() => Promise.reject(new Error('no')));
    await copy.copy('text');
    expect(copy.state).toBe('failed');
    expect(copy.text).toBe('text');
    expect(attendanceCopy.errors.copyFailed).toContain('Could not copy');
  });
});
