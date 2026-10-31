// web/src/lib/report/copy.ts
// User-visible strings this round of the layout loop adds to the report page. Not every
// string on the page lives here -- most of ReportView.svelte and ModeBar.svelte predate
// this file -- but every new one does, so a copy change is one diff in one file rather
// than a search through two components.

export const reportCopy = {
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
  sampleNote:
    'A real beta night with the names changed. Level 10 to 30 players, so the numbers are small. Your own reports look like this.',
} as const;
