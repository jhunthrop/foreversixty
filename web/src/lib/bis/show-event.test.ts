// @vitest-environment jsdom
import { describe, expect, it, vi } from 'vitest';
import { BIS_SHOW_EVENT, sendBisShow, takePendingShow } from './show-event';

describe('bis:show', () => {
  it('dispatches the event and leaves the request pending for a late island', () => {
    const root = document.createElement('div');
    const heard = vi.fn();
    root.addEventListener(BIS_SHOW_EVENT, (event) => heard((event as CustomEvent).detail));
    sendBisShow(root, { faction: 'horde', band: 40 });
    expect(heard).toHaveBeenCalledWith({ faction: 'horde', band: 40 });
    expect(takePendingShow(root)).toEqual({ faction: 'horde', band: 40 });
    expect(takePendingShow(root)).toBeUndefined();
  });

  it('ignores a malformed pending value', () => {
    const root = document.createElement('div');
    root.dataset.bisShow = 'neutral:x';
    expect(takePendingShow(root)).toBeUndefined();
  });
});
