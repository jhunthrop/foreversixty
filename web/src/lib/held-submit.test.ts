// @vitest-environment jsdom
// web/src/lib/held-submit.test.ts
// A submit pressed before the lookup form's handler is wired is held, not sent, and replayed
// once the handler is in place; a form nobody submitted is left alone on release.
import { describe, expect, it, vi } from 'vitest';
import { holdSubmitUntilWired } from './held-submit';

function form(): HTMLFormElement {
  document.body.innerHTML = '<form><input name="name" /><button type="submit">Look up</button></form>';
  return document.querySelector('form')!;
}

describe('holdSubmitUntilWired', () => {
  it('cancels an early submit and replays it once released', () => {
    const lookup = form();
    const release = holdSubmitUntilWired(lookup);
    const early = new Event('submit', { cancelable: true });
    lookup.dispatchEvent(early);
    expect(early.defaultPrevented).toBe(true);

    const wired = vi.fn((event: Event) => event.preventDefault());
    lookup.addEventListener('submit', wired);
    release();
    expect(wired).toHaveBeenCalledTimes(1);
  });

  it('replays nothing when no submit came in before release', () => {
    const lookup = form();
    const release = holdSubmitUntilWired(lookup);
    const wired = vi.fn((event: Event) => event.preventDefault());
    lookup.addEventListener('submit', wired);
    release();
    expect(wired).not.toHaveBeenCalled();
  });
});
