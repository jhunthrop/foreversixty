// web/src/lib/current-character-copy.ts
// Strings for the current-character pointer and the chip (/setup's step card), plus the two the
// simulator's own surfaces reuse, so the wording can never drift between pages that describe
// the same pointer. The header selector's restored line and Forget live in
// character-selector/copy.ts with the rest of the selector's words.
export const currentCharacterCopy = {
  forget: 'Forget',
  openInPlanner: 'Open in planner',
  openInSimulator: 'Open in simulator',
  noCharacterLine: 'No character loaded. Paste an addon export in the planner or the simulator.',
  getTheAddon: "Don't have an export? Get the addon.",
  /** The sim landing list's current-row marker. */
  switchCurrentMarker: 'Current',
} as const;
