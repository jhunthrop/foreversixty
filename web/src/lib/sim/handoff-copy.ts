// web/src/lib/sim/handoff-copy.ts
// Copy for links that hand a bulk results row off somewhere else -- "Plan it" today
// (task 7), more later per this lane's own plan. A new module rather than another key on
// `bulkCopy` (`copy.ts`, 1,247 lines, already over this codebase's 800-line file ceiling):
// one short key would not by itself scatter `bulkCopy`'s vocabulary, but this lane keeps
// adding hand-off strings to it, and each one added here instead is one `copy.ts` never
// has to carry.
export const handoffCopy = {
  /**
   * Design 1: "'Plan it' on each Droptimizer and Top Gear upgrade row (opens the planner
   * with that item equipped)". Also the accessible name for ComboResults.svelte's action
   * column, which carries no visible header text of its own.
   */
  planIt: 'Plan it',
} as const;
