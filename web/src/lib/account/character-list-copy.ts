// web/src/lib/account/character-list-copy.ts
// Every visible string CharacterList.svelte and Account.svelte's ?refreshed=1 toast use.
// Voice, per design/DESIGN-SYSTEM.md: reference, not pitch. State the thing and stop.

export const characterListCopy = {
  heading: 'Characters',
  empty: 'No characters linked yet. Sign in with Battle.net to link them.',
  refreshFromBattlenet: 'Refresh from Battle.net',
  importedFrom: (relative: string): string => `Imported from Battle.net ${relative}.`,
  verified: 'Verified',
  levelPrefix: (level: number): string => `Level ${level}`,
  /** spec 2026-09-22 §7.3, shown on /account after a Battle.net refresh redirect. */
  refreshedToast: 'Characters refreshed from Battle.net.',
  /** spec 2026-09-22 §3.3: the export "how" explained once, here, for every row. */
  introLead:
    'Gear and talents come from Battle.net and refresh nightly. Install the addon to include bags and bank and to update right after a session; paste an export for a character Battle.net has no data for.',
  introInstallAddonLink: 'Install the addon',
  introPasteLink: 'paste an export',
  /** The account rows' and the sim landing rows' build-source pill (spec 2026-09-22 §3.1). */
  battlenetSource: 'Battle.net',
  addonSource: 'Addon',
  noBuildYet: 'No build yet',
} as const;
