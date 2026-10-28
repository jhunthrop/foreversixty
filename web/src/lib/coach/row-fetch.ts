// web/src/lib/coach/row-fetch.ts
// Whether a report row asks the coach API for its line, and what it asks (Ruling 6): only the
// viewer's own characters, only boss pulls long enough to review, never a whole night. Anyone
// else costs no request.
import { reviewableFight, type CoachRowReport } from './report-context';
import { ownedKey, type CoachViewer } from './viewer';

export interface LineRequest {
  reportId: string;
  fightIndex: number;
  playerKey: string;
}

export function lineRequestFor(
  viewer: CoachViewer,
  actorName: string,
  report: CoachRowReport,
): LineRequest | null {
  if (!viewer.signedIn) return null;
  const fight = reviewableFight(report);
  const playerKey = ownedKey(viewer, actorName);
  if (fight === null || playerKey === null) return null;
  return { reportId: report.reportId, fightIndex: fight.index, playerKey };
}

/** An officer gets the on-demand button on a row that is not their own. */
export function officerButtonFor(viewer: CoachViewer, actorName: string, report: CoachRowReport): boolean {
  return viewer.officer && reviewableFight(report) !== null && ownedKey(viewer, actorName) === null;
}
