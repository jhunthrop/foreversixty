// web/src/lib/sim/copy.ts
// Every user-visible string on the simulator lane. Components import from here and tests
// assert against the same constants, the way web/src/lib/planner/rules.ts holds `messages`:
// a copy change is then one diff in one file, and no test asserts on a literal that a
// component could quietly stop rendering.
//
// Voice, per design/DESIGN-SYSTEM.md: reference, not pitch. State the number and stop.

export const simCopy = {
  /** Network and API failures. */
  saveFailed: 'The sim could not be saved; try again.',
  loadFailed: 'That sim did not load.',
  notFound: 'No sim with that id.',
  premiumRequired: 'Running on our servers is a premium feature. The browser lane is free and unlimited.',
  specsFailed: 'Spec support could not be read.',
  characterFailed: 'That character could not be read.',

  /**
   * The run loop's two outcomes. They start here, in the base task, rather than in
   * `run.ts`: Task 5 raises them, Task 11's store stores them and Task 13's e2e asserts
   * them, and a copy of either in `run.ts` is how the button and the alert end up saying
   * two different things. `run.ts` declares no message constants of its own.
   */
  stopped: 'Stopped.',
  failed: 'The engine could not run this character.',

  // Used by the sim page (group D) and the planner's live estimate (group F), which run in
  // parallel. Both need the same sentence, so it starts here rather than in either.
  specUnsupportedLead: 'The simulator does not model this spec yet.',
} as const;
