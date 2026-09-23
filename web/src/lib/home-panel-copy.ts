// web/src/lib/home-panel-copy.ts
// The home page's one account-aware panel (spec 2026-09-22 §3.2, restyled into the hero by
// spec 2026-09-23 §2): reference, not pitch -- one sentence, one button, signed out; the
// hero character, its descriptor, four links and (spec 2026-09-23) a rating figure when one
// exists, signed in.

/** The signed-out block's `id` in `index.astro`: `HomeAccountPanel.svelte` reaches outside
 *  its own root to find it (same cross-island DOM-reach pattern as `SIM_TAB_ATTR` /
 *  `syncTabHrefs` in `lib/sim/tabs.ts`) and mark it `inert`/`aria-hidden` once signed in, so
 *  the duplicate "Sign in with Battle.net" link the grid-overlay CLS trick leaves behind it
 *  is never focusable or announced. One id, exported once, so the producer (index.astro)
 *  and the consumer (HomeAccountPanel.svelte) can never drift apart. */
export const HOME_SIGNED_OUT_ID = 'home-signed-out';

export const homePanelCopy = {
  // Spec 2026-09-23 §2 item 2's exact hero sentence.
  signedOutLine: 'Sign in with Battle.net and your characters arrive with their gear, talents and guild.',
  signInButton: 'Sign in with Battle.net',
  openInPlanner: 'Open in planner',
  openInSimulator: 'Open in simulator',
  logs: 'Logs',
  yourCharacters: 'Your characters',
} as const;
