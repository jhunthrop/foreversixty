// web/src/directives/interaction.ts
//
// Custom `client:interaction` directive: hydrates an island the first time the user
// focuses into it, or presses "/" anywhere on the page (the app-wide search shortcut).
//
// Astro's built-in directives (`load`, `idle`, `visible`) all still start fetching and
// executing the island's JS during the browser's *initial* render pass -- fast enough that
// Lighthouse's LCP simulation counts the extra network requests against the page's LCP even
// though nothing on screen actually depends on them. Waiting for a real user gesture instead
// means the fetch never happens during the window Lighthouse measures, at the cost of a few
// milliseconds of hydration latency on that first interaction (imperceptible locally, and
// Playwright's default retrying assertions tolerate it fine -- see tests/e2e/search.spec.ts).
//
// Two things have to be handled carefully because there's a real (if small) window between
// "user starts interacting" and "the island's JS has finished loading and mounting":
//
// 1. The "/" case: that keypress happens *before* the component's own global-shortcut
//    listener (set up in its onMount effect) exists to catch it, so we focus the
//    server-rendered <input> ourselves.
// 2. Keystrokes typed during that window land in the plain, not-yet-hydrated <input> as
//    normal browser text input (no framework listening yet). When Svelte's `hydrate()` then
//    mounts onto that same DOM node, `bind:value={query}` resets the input back to the
//    component's initial internal state (empty, unless `initial` was passed as a prop) --
//    silently dropping whatever the user already typed. We buffer the input's value/caret
//    position for the duration of hydration and replay it (set the value, dispatch a real
//    `input` event so Svelte's `oninput` picks it up and runs the search, restore the caret)
//    once mounting is done.
import type { ClientDirective } from 'astro';

const interactionDirective: ClientDirective = (load, _options, el) => {
  let hydrating = false;

  const cleanup = () => {
    window.removeEventListener('focusin', onFocusIn, true);
    window.removeEventListener('keydown', onSlashKey, true);
  };

  const hydrateAndReplay = async (input: HTMLInputElement | null) => {
    if (hydrating) return;
    hydrating = true;
    cleanup();

    if (!input) {
      const run = await load();
      await run();
      return;
    }

    // Buffer whatever the user types into the not-yet-hydrated input while the component's
    // JS loads and mounts, since hydration otherwise resets the input to the component's
    // (empty) initial state and drops it -- see the file header for why.
    let bufferedValue = input.value;
    let bufferedStart = input.selectionStart;
    let bufferedEnd = input.selectionEnd;
    const captureBuffer = () => {
      bufferedValue = input.value;
      bufferedStart = input.selectionStart;
      bufferedEnd = input.selectionEnd;
    };
    input.addEventListener('input', captureBuffer);
    input.addEventListener('keydown', captureBuffer);

    const run = await load();
    await run();

    input.removeEventListener('input', captureBuffer);
    input.removeEventListener('keydown', captureBuffer);

    if (bufferedValue) {
      input.focus();
      input.value = bufferedValue;
      if (bufferedStart !== null && bufferedEnd !== null) {
        input.setSelectionRange(bufferedStart, bufferedEnd);
      }
      // A real `input` event (not just setting `.value`) so Svelte's `bind:value`/`oninput`
      // pick up the replayed text and run the search, exactly as if the user had typed it
      // after hydration finished.
      input.dispatchEvent(new Event('input', { bubbles: true }));
    }
  };

  const onFocusIn = (event: FocusEvent) => {
    if (el.contains(event.target as Node)) hydrateAndReplay(el.querySelector('input'));
  };

  const onSlashKey = (event: KeyboardEvent) => {
    const target = event.target;
    const alreadyTyping = target instanceof HTMLInputElement || target instanceof HTMLTextAreaElement;
    if (event.key === '/' && !alreadyTyping) {
      event.preventDefault();
      const input = el.querySelector('input');
      input?.focus();
      hydrateAndReplay(input);
    }
  };

  window.addEventListener('focusin', onFocusIn, true);
  window.addEventListener('keydown', onSlashKey, true);
};

export default interactionDirective;
