// web/src/lib/ui/copy.ts
// The strings the shared state components (components/ui/) show, and the one home of every retry,
// copy and busy word (component audit 2026-11-10 clusters 10, 11, 13 and 16). A page never writes
// "Try again", "Retry" or "Copied" itself: it takes these, so the words read the same everywhere
// (state-guard.test.ts). Reference, not pitch.

/** The retry control's one word, on every error line and panel. */
const RETRY = 'Try again';

export const uiCopy = {
  /** The one loading line (Skeleton's status, a loading row). */
  loading: 'Loading…',
  retry: RETRY,
  copy: 'Copy',
  copyLink: 'Copy link',
  /** Shown for about two seconds after a copy (owner ruling 2026-11-10). */
  copied: 'Copied',
  copyFailed: 'Copy failed',
  copyLinkFailed: 'Could not copy. Select the link and copy it.',
  shareLink: 'Share',
  cancel: 'Cancel',
  showFewer: 'Show fewer',
  loadOlder: 'Load older',
  loadingOlder: 'Loading…',
  olderFailed: 'Older entries did not load.',
} as const;

/** The busy word a button reads while its request is in flight (owner ruling 2026-11-10: spinner
 *  plus a busy word). A caller with a more exact verb passes its own `busyLabel`. */
export const busyCopy = {
  saving: 'Saving…',
  sending: 'Sending…',
  posting: 'Posting…',
  working: 'Working…',
  loading: 'Loading…',
} as const;

/** A failure sentence that ends in the retry word: "That did not save. Try again." */
export function withRetry(sentence: string): string {
  return `${sentence} ${RETRY}.`;
}

/** "Copied" plus what to do next: "Copied. Paste it in game with /fs". */
export function copiedThen(next: string): string {
  return `${uiCopy.copied}. ${next}`;
}

/** The show-all control: "Show all 24". */
export function showAllLabel(total: number): string {
  return `Show all ${total}`;
}

/** The shared dropdown's copy (design/specs/2026-10-30-controls-charts.md §A8). */
export const selectCopy = {
  loading: uiCopy.loading,
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
