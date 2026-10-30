// web/src/lib/ui/copy.ts
// The strings the shared state components (components/ui/) show. Reference, not pitch.
export const uiCopy = {
  loading: 'Loading',
  retry: 'Try again',
  copyLink: 'Copy link',
  copiedLink: 'Copied',
  copyLinkFailed: 'Could not copy. Select the link and copy it.',
  shareLink: 'Share',
} as const;

/** The shared dropdown's copy (design/specs/2026-10-30-controls-charts.md §A8). */
export const selectCopy = {
  loading: 'Loading…',
  noMatch: (query: string): string => `Nothing matches “${query}”.`,
  matchCount: (n: number): string => (n === 1 ? '1 match' : `${n} matches`),
  clearFilter: 'Clear filter',
  you: 'You',
  filterPlayers: 'Filter players',
  filterAbilities: 'Filter abilities',
  filterTargets: 'Filter targets',
  filterSources: 'Filter sources',
  filterTimeZones: 'Filter time zones',
  filterEncounters: 'Filter encounters',
  noEnemies: 'No enemies in this window',
  noAbilities: 'No abilities in this window',
  noTargets: 'No targets in this window',
  noCompareFight: 'No other fight to compare',
  windowRange: (start: string, end: string): string => `${start} to ${end}`,
} as const;
