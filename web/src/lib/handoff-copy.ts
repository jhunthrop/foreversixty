// web/src/lib/handoff-copy.ts
// Every visible string CharacterHandoffLinks.svelte uses. Two products are named here and
// they are never blurred: the addon runs in the game client; the companion is the desktop
// app that pairs with an account and sends what the addon wrote.

export const handoffCopy = {
  openInSimulator: 'Open in simulator',
  openInPlanner: 'Open in planner',
  /** Shown when the site holds no export for the character. Spec 2026-09-22 §2.3: the "how"
   *  (addon vs. companion) is explained once, in the Characters section's own intro line
   *  (character-list-copy.ts) -- this per-row line only points at the one action. */
  needsExportLead: 'No export yet ·',
  needsExportPasteLink: 'paste it here',
  /** A Battle.net-imported character with no gear or talents yet. Blizzard publishes a character's
   *  profile on its own schedule (launch night, 2026-11-04: every new character's profile answered
   *  404), so this says when they come, never that they are coming soon; the nightly refresh
   *  picks them up the day Blizzard publishes. */
  battlenetPendingLead: 'Gear and talents appear once Battle.net publishes this character ·',
  battlenetPendingPasteLink: 'or paste an export',
  pasteHref: '/setup#paste',
} as const;

/** The plain-text form of the needs-export line, for tests and accessible names. */
export const NEEDS_EXPORT_TEXT = `${handoffCopy.needsExportLead} ${handoffCopy.needsExportPasteLink}`;
