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
} as const;

/** The /reports page ("Your reports"): its head, its card, and the words its three states add. */
export const yourReportsPageCopy = {
  eyebrow: 'Logs',
  title: 'Your reports',
  job: 'Every fight you logged, newest first.',
  cardTitle: 'Reports',
  empty: 'No logs yet.',
  signedOut: 'Sign in to see every report you logged.',
  sessionFailed: 'Your account did not load.',
  sampleAction: 'Open the sample report',
} as const;
