// web/src/lib/sim/planner-copy.ts
// The sim strings the planner island itself shows. They live apart from `copy.ts` so the
// island ships these few lines rather than the whole sim vocabulary; `simCopy` spreads them
// back in, so every other reader is unchanged.
export const plannerSimCopy = {
  // --- Task 20: live DPS in the planner. ---
  /** The summary bar's fourth figure, beside Level, Split and Points. */
  plannerDpsLabel: 'DPS',
  /** Under the figure while points are unspent: the live estimate waits for a whole build. */
  plannerDpsPointsToGo: (left: number): string => (left === 1 ? '1 point to go' : `${left} points to go`),
  /** On a phone or a data-saver connection the estimate is asked for, never assumed. */
  plannerDpsShow: 'Show DPS',
  plannerDpsShowNote: 'Runs on this device',
  /** The link that opens the full results for the build on the page. */
  simThisBuild: 'Sim this build',

  // --- Task 21: build unfurls carrying simmed DPS. ---
  includeSimOnCard: 'Include a simmed DPS on the card',
  buildSimUnavailable: 'Needs a DPS estimate first; change a talent or a slot to get one.',
  buildSimRunning: 'Simming this build for the card, a few seconds…',
  buildSimDone: (dps: string, version: string): string =>
    `The card will show ${dps} DPS on engine ${version}.`,
  buildSimSkipped: 'The card will not show a DPS figure; the sim did not finish.',
} as const;
