// web/src/lib/reports/my-reports-copy.ts
// Every visible string MyReports.svelte ("Your reports") uses -- split from
// lib/reports/copy.ts's recentReportsCopy, which is RecentReports.svelte's own strings for
// the public feed, a different component with a different failure and empty case.
export const myReportsCopy = {
  failed: 'Your reports did not load.',
  empty: 'No reports yet. Upload a log, or let the companion send your next raid.',
  signedOut: 'Sign in to see your reports.',
  signInAction: 'Sign in with Battle.net',
  uploadAction: 'Upload a log',
  older: 'Older reports',
} as const;

/** The /reports page ("Your reports"): its head and the words its states add. */
export const yourReportsPageCopy = {
  title: 'Your reports',
  empty: 'No logs yet. Upload a log, or let the companion send your next raid.',
  signedOut: 'Sign in to see every report you logged.',
  sessionFailed: 'Your account did not load.',
  sampleAction: 'Open the sample report',
} as const;
