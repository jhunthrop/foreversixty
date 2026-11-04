// web/src/lib/account/character-list-copy.ts
// Every visible string CharacterList.svelte and Account.svelte's ?refreshed=1 toast use.
// Voice, per design/DESIGN-SYSTEM.md: reference, not pitch. State the thing and stop.

export const characterListCopy = {
  heading: 'Characters',
  /** Battle.net is the primary path (2026-11-03): a Battle.net sign-in imports every Forever
   *  character, so an empty list means no Forever character on that account yet, an email
   *  sign-in, or a region Blizzard does not serve (KR); the addon covers the last two. */
  empty:
    'No Forever characters on this Battle.net account yet. Make one in game, then refresh from Battle.net. Playing in Korea? Install the addon and the companion, or paste an export.',
  emptyByEmail:
    'No characters yet. Sign in with Battle.net to bring your Forever characters here, or install the addon and the companion.',
  getSetUp: 'Get set up',
  getSetUpHref: '/setup',
  pasteAnExport: 'Paste an export',
  pasteHref: '/setup#paste',
  refreshFromBattlenet: 'Refresh from Battle.net',
  importedFrom: (relative: string): string => `Imported from Battle.net ${relative}.`,
  verified: 'Verified',
  /** The main character's pill, and the control on every alt. */
  main: 'Main',
  current: 'Current',
  setAsMain: 'Set as main',
  setAsMainFailed: 'Could not set your main just now.',
  levelPrefix: (level: number): string => `Level ${level}`,
  /** spec 2026-09-22 §7.3, shown on /account after a Battle.net refresh redirect. */
  refreshedToast: 'Characters refreshed from Battle.net.',
  /** spec 2026-09-22 §3.3: the export "how" explained once, here, for every row. */
  introBattlenetLine: 'Your characters come from Battle.net each time you sign in.',
  introInstallAddonLink: 'Install the addon',
  introAddonTail: 'for bags, bank and live sync. In Korea, or Battle.net is down?',
  introPasteLink: 'Paste an export',
  introPasteTail: 'instead.',
  /** The account rows' and the sim landing rows' build-source pill (spec 2026-09-22 §3.1). */
  battlenetSource: 'Battle.net',
  addonSource: 'Addon',
  noBuildYet: 'No build yet',
} as const;
