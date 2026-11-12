// web/src/lib/report/copy.ts
// User-visible strings this round of the layout loop adds to the report page. Not every
// string on the page lives here -- most of ReportView.svelte and ModeBar.svelte predate
// this file -- but every new one does, so a copy change is one diff in one file rather
// than a search through two components.
import { uiCopy } from '../ui/copy';

export const reportCopy = {
  /** The roster's show-all under its first rows. */
  showAllPlayers: (total: number): string => `Show all ${total} players`,
  /** The phone's collapsed chart strip, which expands the chart in place. */
  showChart: 'Show chart',
  /** The category tab row's overflow menu, and its label while none of its own tabs is active. */
  more: 'More',
  /** The phone table-switcher's visible label, so it reads next to "Source" rather than under it. */
  table: 'Table',
  /** The Source select (spec 2026-10-30-controls-charts §A10 row 1): its label, the two
   *  scopes and the role groups its players sit under. */
  source: 'Source',
  allFriendlies: 'All friendlies',
  /** The default scope when it narrows an open-world log to the logger's group (lib/report/group.ts). */
  yourGroup: 'Your group',
  yourRaid: 'Your raid',
  /** Every player the log saw, the group or not. */
  everyoneInRange: 'Everyone in range',
  allEnemies: 'All enemies',
  sourceGroups: { tank: 'Tanks', healer: 'Healers', dps: 'Damage' },
  /** The filter bar's two selects (§A10 rows 6, 7). */
  target: 'Target',
  ability: 'Ability',
  everyTarget: 'Every target',
  everyAbility: 'Every ability',
  /** Compare mode's selects (§A10 rows 8 to 11). */
  compareWith: 'Compare with',
  pickFight: 'Pick a fight',
  bossPulls: 'Boss pulls',
  trash: 'Trash',
  orAPlayer: 'Or a player',
  compareFightsInstead: 'Compare fights instead',
  metric: 'Metric',
  phase: 'Phase',
  wholePull: 'The whole pull',
  spec: 'Spec',
  everySpec: 'Every spec',
  /** The merged preset/slider row's select, which replaced the old six-button preset row. */
  window: 'Window',
  /** The control bar's first select (logs spec 2026-11-07 6.1). */
  fight: 'Fight',
  /** The tab strip's accessible name. */
  tabsLabel: 'Report tabs',
  /** The Plate's names on the state screens. */
  fightsLabel: 'Fights',
  onePullLabel: 'One pull only',
  /** The head's two actions, and the phone's one menu button that holds them. */
  copyLink: 'Copy link',
  /** The phone band's ··· menu holding Sim this fight and Copy link. */
  actionsLabel: 'Report actions',
  reportActions: 'Report actions',
  /** The Filters button and popover (logs spec 6.1), as the BiS Sources pattern. */
  filters: 'Filters',
  clearFilters: 'Clear filters',
  /** Every tab's empty table, and its chart. */
  nothingInWindow: 'Nothing in this window.',
  /** Section 6.3: a report that did not load, a missing one, a tab whose chunk failed. */
  reportFailed: 'Report data did not load.',
  reportMissing: 'No report with that id.',
  openYourReports: 'Open your reports',
  viewFailed: 'This view did not load.',
  tryAgain: uiCopy.retry,
  /** The night view's line on a tab that is one pull's (section 6.3). */
  onePullOnly: 'This tab is one pull’s. Pick a pull from the Fight list to see it.',
  /** A pull the logger's group was not in (the scope falls back to everyone in range). */
  groupMissing: 'Your group wasn’t in this pull.',
  ratingNight: 'Ratings are one pull’s. Pick a boss pull from the list to see one.',
  /** The Parse column's one line when the rankings did not answer (boss pulls only). */
  parsesUnknown: 'Parses didn’t load.',
  /** The Live tag while the companion logs. */
  liveTag: 'Live',
  loggingNow: 'Logging now',
  /** The header of a report whose title is not known yet. */
  yourUpload: 'Your upload',
  /** The fights strip of a report being parsed, and the hover of every tab that needs fights. */
  parsing: 'Parsing your log… this takes a minute or two for a long night',
  parseFailed: 'This log could not be parsed.',
  parseFailedLink: 'Check your logging setup',
  /** The fights strip of a live report whose first fight has not been uploaded yet. */
  waiting:
    'Logging · waiting for the first fight. The companion uploads a fight when it ends; this page updates by itself.',
  /** The waiting screen once the companion has not been heard from for a while; the minutes are rounded. */
  quiet: (minutes: number) =>
    `Not heard from the companion for ${minutes} minutes. Check that it is running.`,
  /** The badge that replaces "Live" while the companion is quiet. */
  quietBadge: 'Quiet',
  /** The badge where "Live" goes on the bundled sample report. */
  eyebrow: 'Report',
  sampleBadge: 'Sample report',
  /** The state tag beside the credit sentence of a report no character is credited for. */
  notCreditedTag: 'Not credited',
  /** The one muted line under the sample's subtitle: what it is, and that yours look like it. */
  sampleNote: 'A real dungeon night, names changed.',
} as const;
