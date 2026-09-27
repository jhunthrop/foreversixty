// web/src/lib/account/character-list-copy.ts
// Every visible string CharacterList.svelte and Account.svelte's ?refreshed=1 toast use.
// Voice, per design/DESIGN-SYSTEM.md: reference, not pitch. State the thing and stop.

export const characterListCopy = {
  heading: 'Characters',
  /** Blizzard serves no Forever character data yet (the API's import is off until it
   *  does), so the empty state sends people to the addon's paste box, not to Battle.net. */
  empty:
    'No characters yet. Blizzard does not serve Forever character data yet, so paste an export from the addon.',
  pasteAnExport: 'Paste an export',
  pasteHref: '/setup#paste',
  refreshFromBattlenet: 'Refresh from Battle.net',
  importedFrom: (relative: string): string => `Imported from Battle.net ${relative}.`,
  verified: 'Verified',
  /** The main character's pill, and the control on every alt. */
  main: 'Main',
  setAsMain: 'Set as main',
  setAsMainFailed: 'Could not set your main just now.',
  levelPrefix: (level: number): string => `Level ${level}`,
  /** spec 2026-09-22 §7.3, shown on /account after a Battle.net refresh redirect. */
  refreshedToast: 'Characters refreshed from Battle.net.',
  /** spec 2026-09-22 §3.3: the export "how" explained once, here, for every row. */
  introBattlenetLine:
    'Characters come from the addon for now: Blizzard serves no Forever character data yet.',
  introInstallAddonLink: 'Install the addon',
  introAddonTail: 'to export gear, talents, bags and bank, then',
  introPasteLink: 'paste an export',
  introPasteTail: 'to add a character here.',
  /** The account rows' and the sim landing rows' build-source pill (spec 2026-09-22 §3.1). */
  battlenetSource: 'Battle.net',
  addonSource: 'Addon',
  noBuildYet: 'No build yet',
} as const;
