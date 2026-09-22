// web/src/lib/reports/my-reports-copy.ts
// Every visible string MyReports.svelte ("Your reports") uses -- split from
// lib/reports/copy.ts's recentReportsCopy, which is RecentReports.svelte's own strings for
// the public feed, a different component with a different failure and empty case.
export const myReportsCopy = {
  failed: 'Your reports did not load.',
  empty: 'No reports yet. Upload a log or run the desktop companion.',
  uploadAction: 'Upload a log',
} as const;
