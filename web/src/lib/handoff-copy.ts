// web/src/lib/handoff-copy.ts
// Every visible string CharacterHandoffLinks.svelte uses. Two products are named here and
// they are never blurred: the addon runs in the game client; the companion is the desktop
// app that pairs with an account and sends what the addon wrote.

export const handoffCopy = {
  openInSimulator: 'Open in simulator',
  openInPlanner: 'Open in planner',
  /** Shown when the site holds no export for the character: the two ways to get one. */
  needsExportLead: 'No export yet. In game, type /fs export and',
  needsExportPasteLink: 'paste it here',
  needsExportTail: ', or run the desktop companion to send it for you.',
  pasteHref: '/addon#paste',
} as const;

/** The plain-text form of the needs-export line, for tests and accessible names. */
export const NEEDS_EXPORT_TEXT = `${handoffCopy.needsExportLead} ${handoffCopy.needsExportPasteLink}${handoffCopy.needsExportTail}`;
