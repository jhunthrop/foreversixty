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
  /**
   * Task 9, spec section 1: "a one-line framing on /sim for visitors below 60" --
   * ScopeNote.astro's second line, below the damage-specs-only one, so a visitor who has
   * not levelled yet is told where to go rather than left to guess why the page wants a
   * character it cannot describe.
   */
  belowSixtyFraming:
    'The simulator models level 60 characters. Below 60, plan your build in the planner and come back.',
  /**
   * Task 9, spec section 1: "when a planner build with fewer than 51 points is sent to
   * the sim, the sim says it is simming it as a level 60 with those talents rather than
   * silently relabelling it". CharacterStrip.svelte's own `levelSuffix` appends this
   * beside the talent-point count for exactly that character.
   */
  simmedAtSixty: 'Simmed as a level 60 with these talents.',
} as const;
