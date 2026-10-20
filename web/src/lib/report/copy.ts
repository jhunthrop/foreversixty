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
} as const;
