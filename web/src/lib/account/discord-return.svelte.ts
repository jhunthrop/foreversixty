// web/src/lib/account/discord-return.svelte.ts
// The Discord block's two page-level behaviours that are not a flow: moving focus to a control
// that may not exist yet, and reading the outcome Discord's callback left in the address bar.
import { tick } from 'svelte';
import { refreshAfterConnect, returnOutcome, takeReturnFlag, type ReturnOutcome } from './discord-actions';

export interface FocusQueue {
  /** Focus the first element matching `selector` once it exists (null cancels). */
  request: (selector: string | null) => void;
}

/**
 * A save or remove changes the block's state when the patched `me` arrives, so the target may
 * not exist yet: the request is retried each time `watch` (read inside an effect) changes. Call
 * during component setup.
 */
export function createFocusQueue(watch: () => unknown): FocusQueue {
  let pending = $state<string | null>(null);
  $effect(() => {
    const selector = pending;
    if (selector === null) return;
    void watch();
    void tick().then(() => {
      const target = document.querySelector<HTMLElement>(selector);
      if (target === null) return;
      target.focus();
      if (pending === selector) pending = null;
    });
  });
  return {
    request: (selector) => {
      pending = selector;
    },
  };
}

/** Reads and clears `?discord=`; a `connected` outcome also re-reads `/v1/me` once. Null when absent. */
export function consumeReturn(
  onRefreshFailure: () => void,
  location: Location = window.location,
  history: History = window.history,
): ReturnOutcome | null {
  const flag = takeReturnFlag(location, history);
  if (flag === null) return null;
  const outcome = returnOutcome(flag);
  if (outcome.refresh) refreshAfterConnect(onRefreshFailure);
  return outcome;
}
