// web/src/lib/sim/server-runs-copy.ts
// Every string of experience spec 2026-10-24 section 6 for "Your server runs" and the band's
// "done while away" line. Functions take the data and return the string; no number is typed
// into prose. The shared retry word is `uiCopy.retry`, not repeated here.
import { simCopy } from './copy';
import { estimateWords, isOverdue } from './server-runs';

function formatCount(value: number): string {
  return value.toLocaleString('en-US');
}

export const serverRunsCopy = {
  heading: 'Server runs',
  empty: `No server runs yet. Choose ${simCopy.runOnServers} when you start a sim, and it keeps going after you close this tab.`,
  loadingSr: 'Reading your server runs…',
  loadFailed: 'Your server runs did not load.',
  overflow: 'Older finished runs are in Your sims below.',
  unreadSr: 'Not opened yet',

  queued: 'Starting',
  running: 'Running',
  /** The muted words after "Running": the estimate, "taking longer than usual", or nothing. */
  runningLine: (estimateSeconds: number | undefined, elapsedSeconds: number): string => {
    if (estimateSeconds === undefined) return '';
    if (isOverdue(estimateSeconds, elapsedSeconds)) return ' · taking longer than usual';
    return ` · takes ${estimateWords(estimateSeconds)}`;
  },
  failedLine: (reason: string | undefined): string =>
    reason === undefined || reason === '' ? 'Failed' : `Failed: ${reason}`,

  startedLine: (age: string): string => `Started ${age}`,
  finishedLine: (age: string): string => `Finished ${age}`,
  endedLine: (age: string): string => `Ended ${age}`,

  runAgain: 'Run again',
  runAgainPending: 'Starting…',
  runAgainDone: 'Started',
  runAgainFailed: 'Could not start it again. Try again.',
  planLapsed: 'Server sims are part of Premium.',
  seePlans: 'See plans',

  /** The sentence in two parts, so the band can set the DPS figure apart in mono. */
  bandLineParts: (specLabel: string, dps: number): { prefix: string; figure: string } => ({
    prefix: `Your ${specLabel} sim finished: `,
    figure: `${formatCount(Math.round(dps))} DPS`,
  }),
  bandOpen: 'Open result →',
  bandOpenPhone: '→',
  moreLabel: (count: number): string => `and ${formatCount(count)} more`,
  dismissLabel: (specLabel: string): string => `Dismiss: your ${specLabel} sim finished`,
  dismissFailed: 'Could not dismiss that. Try again.',
};
