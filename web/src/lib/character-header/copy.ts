// web/src/lib/character-header/copy.ts
// The strings the shared character header owns (design/specs/2026-10-11-character-header.md
// sections 2 and 3): the label, the one shared signed-out line, the stale tail, the retry link.
// A page supplies its own answer words (see types.ts); nothing here is page-specific.
export const characterHeaderCopy = {
  /** Always this, whatever the account holds: it must not change meaning when an alt is added. */
  label: 'Your character',
  /** The one shared signed-out band: no button, the nav selector is the only sign-in control. */
  signedOut: 'Sign in or paste an export to see your character here',
  /** A signed-in account whose characters the site has not seen yet. */
  noCharacters: (who: string): string =>
    `Signed in as ${who}. No characters yet: install the addon and the companion, then log out once in game.`,
  getSetUp: 'Get set up',
  getSetUpHref: '/setup',
  /** Only the words past the sync age go ember (nav selector spec 2.11 threshold). */
  staleTail: 'Log out in the game to sync',
  tryAgain: 'Try again',
  loading: 'Loading your character',
  levelWord: 'Level',
} as const;
