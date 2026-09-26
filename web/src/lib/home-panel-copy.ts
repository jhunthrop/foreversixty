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
  logs: 'Logs',
  yourCharacters: 'Your characters',
  planTalents: 'Plan talents',
  // Review round 1 fix item 1: "Get the build" retired everywhere on the page -- the hero's
  // one primary action reads "Open the planner" when the character has no build yet (it is
  // the same next step the planner door itself points to), never a claim that a build
  // already exists to retrieve.
  openThePlanner: 'Open the planner',
  simCharacter: (name: string): string => `Sim ${name}`,
  /** The chip row under the hero: every other character, one click to make it current. */
  switchTo: (name: string): string => `Switch to ${name}`,
  moreCharacters: (count: number): string => `+${count} more`,
  noSimYet: 'No sim yet.',
  runAction: 'Run',
  noBuildYet: 'No build yet.',
  continuePlanning: 'Continue planning',
  /** The planner door card's no-build state (review round 1 fix item 1): one link, not a
   *  muted status line plus a separate "Get the build" action, and it sends the visitor to
   *  the fastest way to get one (pasting an export) rather than the account page. */
  pasteAnExport: 'Paste an export',
  noLogsYet: 'No logs yet.',
  uploadALog: 'Upload a log',
  openAction: 'Open',
  notRatedYet: 'Not rated yet.',
  rankingsForClass: (klass: string): string => `Rankings for ${klass}`,
  /** Guides carries no per-character state (review round 1 fix item 2): its status line is
   *  the site's own fixed fact, the same figure `home-landing-copy.ts`'s Guides sentence
   *  names, so the signed-in card's live-status shape still applies to it. */
  guidesStatus: '27 spec guides',
} as const;

/** How many other characters the hub's chip row shows before "+N more" links to /account. */
export const HOME_CHIP_LIMIT = 6;
