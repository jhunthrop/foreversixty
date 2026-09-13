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
// The "/" case needs one extra step: the keypress that triggers hydration happens before the
// component's own global shortcut listener (set up in its onMount effect) exists to catch it,
// so we replay it by focusing the island's input ourselves once hydration finishes.
import type { ClientDirective } from 'astro';

const interactionDirective: ClientDirective = (load, _options, el) => {
  let hydrating = false;

  const cleanup = () => {
    window.removeEventListener('focusin', onFocusIn, true);
    window.removeEventListener('keydown', onSlashKey, true);
  };

  const hydrate = async (afterHydrate?: () => void) => {
    if (hydrating) return;
    hydrating = true;
    cleanup();
    const run = await load();
    await run();
    afterHydrate?.();
  };

  const onFocusIn = (event: FocusEvent) => {
    if (el.contains(event.target as Node)) hydrate();
  };

  const onSlashKey = (event: KeyboardEvent) => {
    const target = event.target;
    const alreadyTyping = target instanceof HTMLInputElement || target instanceof HTMLTextAreaElement;
    if (event.key === '/' && !alreadyTyping) {
      hydrate(() => el.querySelector('input')?.focus());
    }
  };

  window.addEventListener('focusin', onFocusIn, true);
  window.addEventListener('keydown', onSlashKey, true);
};

export default interactionDirective;
