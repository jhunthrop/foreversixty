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
  /** The spine bar's Switch popover (`CharacterSwitchList.svelte`, spec 2026-09-25 §4.1). */
  switchCurrentMarker: 'Current',
  switchAction: 'Switch',
  /** Home rebuild spec §3.B.4: the Switch character panel's own trailing stat for a
   *  non-current row, once its own upgrade count is known (`HomeSwitchCharacterPanel.svelte`,
   *  `lib/home/upgrades.ts`) -- replaces the plain "Switch" label with the one more useful
   *  fact the mock shows there; absent (no stat at all) for a row that cannot be computed. */
  switchUpgradesStat: (upgradeCount: number): string =>
    `${upgradeCount} upgrade${upgradeCount === 1 ? '' : 's'}`,
  /** The spine bar's signed-out, no-pointer line -- spec 2026-09-25 section 4.1, distinct
   *  from `noCharacterLine` above (that one names the planner/simulator paste boxes
   *  specifically; this one names both entry points the bar itself offers). */
  barSignedOutLine: 'Sign in with Battle.net or paste an export to point the site at your character.',
  barSignIn: 'Sign in with Battle.net',
  barPasteExport: 'Paste an export',
  barSwitch: 'Switch',
} as const;
