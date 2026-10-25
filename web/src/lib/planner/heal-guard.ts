// web/src/lib/planner/heal-guard.ts
// When the session's pointer heal (a fresh session settling on the account's main) must leave an
// open planner alone: the heal moves a pristine planner to the main, but never replaces a build
// the visitor owns. A send in flight needs no case of its own -- sending requires spent points.

export interface PlannerWork {
  /** A link names the build (`?code=`, `?talents=`, `?class=`, `?spec=`). */
  linkNamesBuild: boolean;
  /** Talent points spent on screen. */
  pointsSpent: number;
  /** Worn items on screen. */
  gearCount: number;
}

export function healWouldDiscardPlannerWork(work: PlannerWork): boolean {
  return work.linkNamesBuild || work.pointsSpent > 0 || work.gearCount > 0;
}
