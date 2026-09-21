// web/src/lib/current-character-copy.ts
// Strings for the current-character pointer, the chip, and the bare-load restore banner.
// Shared across every page that mounts CurrentCharacterChip.svelte -- sim, planner, and
// (Lane A) /account, /character/<key>, /addon -- so the wording can never drift between
// pages that describe the same pointer.
export const currentCharacterCopy = {
  restoredNote: 'Restored your last character.',
  forget: 'Forget',
  openInPlanner: 'Open in planner',
  openInSimulator: 'Open in simulator',
  copyAddonCode: 'Copy addon code',
  copiedAddonCode: 'Copied',
  noCharacterLine: 'No character loaded. Paste an addon export in the planner or the simulator.',
  getTheAddon: "Don't have an export? Get the addon.",
} as const;
