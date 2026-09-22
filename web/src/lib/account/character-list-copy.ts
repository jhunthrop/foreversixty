// web/src/lib/account/character-list-copy.ts
// Every visible string CharacterList.svelte and Account.svelte's ?refreshed=1 toast use.
// Voice, per design/DESIGN-SYSTEM.md: reference, not pitch. State the thing and stop.

export const characterListCopy = {
  heading: 'Characters',
  empty: 'No characters linked yet. Sign in with Battle.net to link them.',
  refreshFromBattlenet: 'Refresh from Battle.net',
  pasteAnExport: 'Paste an export',
  importedFrom: (relative: string): string => `Imported from Battle.net ${relative}.`,
  verified: 'Verified',
  levelPrefix: (level: number): string => `Level ${level}`,
  itemLevelPrefix: (itemLevel: number): string => `ilvl ${itemLevel}`,
  /** spec 2026-09-22 §7.3, shown on /account after a Battle.net refresh redirect. */
  refreshedToast: 'Characters refreshed from Battle.net.',
} as const;
